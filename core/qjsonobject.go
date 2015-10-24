package core

//#include "qjsonobject.h"
import "C"
import (
	"strings"
	"unsafe"
)

type QJsonObject struct {
	ptr unsafe.Pointer
}

type QJsonObjectITF interface {
	QJsonObjectPTR() *QJsonObject
}

func (p *QJsonObject) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QJsonObject) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQJsonObject(ptr QJsonObjectITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QJsonObjectPTR().Pointer()
	}
	return nil
}

func QJsonObjectFromPointer(ptr unsafe.Pointer) *QJsonObject {
	var n = new(QJsonObject)
	n.SetPointer(ptr)
	return n
}

func (ptr *QJsonObject) QJsonObjectPTR() *QJsonObject {
	return ptr
}

func (ptr *QJsonObject) Contains(key string) bool {
	if ptr.Pointer() != nil {
		return C.QJsonObject_Contains(C.QtObjectPtr(ptr.Pointer()), C.CString(key)) != 0
	}
	return false
}

func (ptr *QJsonObject) Count() int {
	if ptr.Pointer() != nil {
		return int(C.QJsonObject_Count(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QJsonObject) Empty() bool {
	if ptr.Pointer() != nil {
		return C.QJsonObject_Empty(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QJsonObject) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QJsonObject_IsEmpty(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QJsonObject) Keys() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QJsonObject_Keys(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QJsonObject) Length() int {
	if ptr.Pointer() != nil {
		return int(C.QJsonObject_Length(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QJsonObject) Size() int {
	if ptr.Pointer() != nil {
		return int(C.QJsonObject_Size(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QJsonObject) DestroyQJsonObject() {
	if ptr.Pointer() != nil {
		C.QJsonObject_DestroyQJsonObject(C.QtObjectPtr(ptr.Pointer()))
	}
}
