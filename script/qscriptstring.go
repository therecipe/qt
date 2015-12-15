package script

//#include "script.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QScriptString struct {
	ptr unsafe.Pointer
}

type QScriptString_ITF interface {
	QScriptString_PTR() *QScriptString
}

func (p *QScriptString) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QScriptString) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQScriptString(ptr QScriptString_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptString_PTR().Pointer()
	}
	return nil
}

func NewQScriptStringFromPointer(ptr unsafe.Pointer) *QScriptString {
	var n = new(QScriptString)
	n.SetPointer(ptr)
	return n
}

func (ptr *QScriptString) QScriptString_PTR() *QScriptString {
	return ptr
}

func NewQScriptString() *QScriptString {
	defer qt.Recovering("QScriptString::QScriptString")

	return NewQScriptStringFromPointer(C.QScriptString_NewQScriptString())
}

func NewQScriptString2(other QScriptString_ITF) *QScriptString {
	defer qt.Recovering("QScriptString::QScriptString")

	return NewQScriptStringFromPointer(C.QScriptString_NewQScriptString2(PointerFromQScriptString(other)))
}

func (ptr *QScriptString) IsValid() bool {
	defer qt.Recovering("QScriptString::isValid")

	if ptr.Pointer() != nil {
		return C.QScriptString_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptString) ToString() string {
	defer qt.Recovering("QScriptString::toString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QScriptString_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScriptString) DestroyQScriptString() {
	defer qt.Recovering("QScriptString::~QScriptString")

	if ptr.Pointer() != nil {
		C.QScriptString_DestroyQScriptString(ptr.Pointer())
	}
}
