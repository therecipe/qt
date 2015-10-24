package qml

//#include "qqmlscriptstring.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QQmlScriptString struct {
	ptr unsafe.Pointer
}

type QQmlScriptStringITF interface {
	QQmlScriptStringPTR() *QQmlScriptString
}

func (p *QQmlScriptString) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QQmlScriptString) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQQmlScriptString(ptr QQmlScriptStringITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlScriptStringPTR().Pointer()
	}
	return nil
}

func QQmlScriptStringFromPointer(ptr unsafe.Pointer) *QQmlScriptString {
	var n = new(QQmlScriptString)
	n.SetPointer(ptr)
	return n
}

func (ptr *QQmlScriptString) QQmlScriptStringPTR() *QQmlScriptString {
	return ptr
}

func NewQQmlScriptString() *QQmlScriptString {
	return QQmlScriptStringFromPointer(unsafe.Pointer(C.QQmlScriptString_NewQQmlScriptString()))
}

func NewQQmlScriptString2(other QQmlScriptStringITF) *QQmlScriptString {
	return QQmlScriptStringFromPointer(unsafe.Pointer(C.QQmlScriptString_NewQQmlScriptString2(C.QtObjectPtr(PointerFromQQmlScriptString(other)))))
}

func (ptr *QQmlScriptString) BooleanLiteral(ok bool) bool {
	if ptr.Pointer() != nil {
		return C.QQmlScriptString_BooleanLiteral(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(ok))) != 0
	}
	return false
}

func (ptr *QQmlScriptString) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QQmlScriptString_IsEmpty(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlScriptString) IsNullLiteral() bool {
	if ptr.Pointer() != nil {
		return C.QQmlScriptString_IsNullLiteral(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlScriptString) IsUndefinedLiteral() bool {
	if ptr.Pointer() != nil {
		return C.QQmlScriptString_IsUndefinedLiteral(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlScriptString) StringLiteral() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlScriptString_StringLiteral(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}
