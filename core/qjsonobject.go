package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"strings"
	"unsafe"
)

type QJsonObject struct {
	ptr unsafe.Pointer
}

type QJsonObject_ITF interface {
	QJsonObject_PTR() *QJsonObject
}

func (p *QJsonObject) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QJsonObject) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQJsonObject(ptr QJsonObject_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QJsonObject_PTR().Pointer()
	}
	return nil
}

func NewQJsonObjectFromPointer(ptr unsafe.Pointer) *QJsonObject {
	var n = new(QJsonObject)
	n.SetPointer(ptr)
	return n
}

func (ptr *QJsonObject) QJsonObject_PTR() *QJsonObject {
	return ptr
}

func (ptr *QJsonObject) Contains(key string) bool {
	defer qt.Recovering("QJsonObject::contains")

	if ptr.Pointer() != nil {
		return C.QJsonObject_Contains(ptr.Pointer(), C.CString(key)) != 0
	}
	return false
}

func (ptr *QJsonObject) Count() int {
	defer qt.Recovering("QJsonObject::count")

	if ptr.Pointer() != nil {
		return int(C.QJsonObject_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QJsonObject) Empty() bool {
	defer qt.Recovering("QJsonObject::empty")

	if ptr.Pointer() != nil {
		return C.QJsonObject_Empty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJsonObject) IsEmpty() bool {
	defer qt.Recovering("QJsonObject::isEmpty")

	if ptr.Pointer() != nil {
		return C.QJsonObject_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJsonObject) Keys() []string {
	defer qt.Recovering("QJsonObject::keys")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QJsonObject_Keys(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QJsonObject) Length() int {
	defer qt.Recovering("QJsonObject::length")

	if ptr.Pointer() != nil {
		return int(C.QJsonObject_Length(ptr.Pointer()))
	}
	return 0
}

func (ptr *QJsonObject) Size() int {
	defer qt.Recovering("QJsonObject::size")

	if ptr.Pointer() != nil {
		return int(C.QJsonObject_Size(ptr.Pointer()))
	}
	return 0
}

func (ptr *QJsonObject) DestroyQJsonObject() {
	defer qt.Recovering("QJsonObject::~QJsonObject")

	if ptr.Pointer() != nil {
		C.QJsonObject_DestroyQJsonObject(ptr.Pointer())
	}
}
