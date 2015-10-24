package script

//#include "qscriptstring.h"
import "C"
import (
	"unsafe"
)

type QScriptString struct {
	ptr unsafe.Pointer
}

type QScriptStringITF interface {
	QScriptStringPTR() *QScriptString
}

func (p *QScriptString) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QScriptString) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQScriptString(ptr QScriptStringITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptStringPTR().Pointer()
	}
	return nil
}

func QScriptStringFromPointer(ptr unsafe.Pointer) *QScriptString {
	var n = new(QScriptString)
	n.SetPointer(ptr)
	return n
}

func (ptr *QScriptString) QScriptStringPTR() *QScriptString {
	return ptr
}

func NewQScriptString() *QScriptString {
	return QScriptStringFromPointer(unsafe.Pointer(C.QScriptString_NewQScriptString()))
}

func NewQScriptString2(other QScriptStringITF) *QScriptString {
	return QScriptStringFromPointer(unsafe.Pointer(C.QScriptString_NewQScriptString2(C.QtObjectPtr(PointerFromQScriptString(other)))))
}

func (ptr *QScriptString) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QScriptString_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QScriptString) ToString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QScriptString_ToString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QScriptString) DestroyQScriptString() {
	if ptr.Pointer() != nil {
		C.QScriptString_DestroyQScriptString(C.QtObjectPtr(ptr.Pointer()))
	}
}
