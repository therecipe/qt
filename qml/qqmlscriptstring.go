package qml

//#include "qml.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"unsafe"
)

type QQmlScriptString struct {
	ptr unsafe.Pointer
}

type QQmlScriptString_ITF interface {
	QQmlScriptString_PTR() *QQmlScriptString
}

func (p *QQmlScriptString) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QQmlScriptString) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQQmlScriptString(ptr QQmlScriptString_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlScriptString_PTR().Pointer()
	}
	return nil
}

func NewQQmlScriptStringFromPointer(ptr unsafe.Pointer) *QQmlScriptString {
	var n = new(QQmlScriptString)
	n.SetPointer(ptr)
	return n
}

func (ptr *QQmlScriptString) QQmlScriptString_PTR() *QQmlScriptString {
	return ptr
}

func NewQQmlScriptString() *QQmlScriptString {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlScriptString::QQmlScriptString")
		}
	}()

	return NewQQmlScriptStringFromPointer(C.QQmlScriptString_NewQQmlScriptString())
}

func NewQQmlScriptString2(other QQmlScriptString_ITF) *QQmlScriptString {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlScriptString::QQmlScriptString")
		}
	}()

	return NewQQmlScriptStringFromPointer(C.QQmlScriptString_NewQQmlScriptString2(PointerFromQQmlScriptString(other)))
}

func (ptr *QQmlScriptString) BooleanLiteral(ok bool) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlScriptString::booleanLiteral")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QQmlScriptString_BooleanLiteral(ptr.Pointer(), C.int(qt.GoBoolToInt(ok))) != 0
	}
	return false
}

func (ptr *QQmlScriptString) IsEmpty() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlScriptString::isEmpty")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QQmlScriptString_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlScriptString) IsNullLiteral() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlScriptString::isNullLiteral")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QQmlScriptString_IsNullLiteral(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlScriptString) IsUndefinedLiteral() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlScriptString::isUndefinedLiteral")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QQmlScriptString_IsUndefinedLiteral(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlScriptString) NumberLiteral(ok bool) float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlScriptString::numberLiteral")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QQmlScriptString_NumberLiteral(ptr.Pointer(), C.int(qt.GoBoolToInt(ok))))
	}
	return 0
}

func (ptr *QQmlScriptString) StringLiteral() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlScriptString::stringLiteral")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlScriptString_StringLiteral(ptr.Pointer()))
	}
	return ""
}
