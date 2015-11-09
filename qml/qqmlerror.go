package qml

//#include "qqmlerror.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QQmlError struct {
	ptr unsafe.Pointer
}

type QQmlError_ITF interface {
	QQmlError_PTR() *QQmlError
}

func (p *QQmlError) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QQmlError) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQQmlError(ptr QQmlError_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlError_PTR().Pointer()
	}
	return nil
}

func NewQQmlErrorFromPointer(ptr unsafe.Pointer) *QQmlError {
	var n = new(QQmlError)
	n.SetPointer(ptr)
	return n
}

func (ptr *QQmlError) QQmlError_PTR() *QQmlError {
	return ptr
}

func NewQQmlError() *QQmlError {
	return NewQQmlErrorFromPointer(C.QQmlError_NewQQmlError())
}

func NewQQmlError2(other QQmlError_ITF) *QQmlError {
	return NewQQmlErrorFromPointer(C.QQmlError_NewQQmlError2(PointerFromQQmlError(other)))
}

func (ptr *QQmlError) Column() int {
	if ptr.Pointer() != nil {
		return int(C.QQmlError_Column(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlError) Description() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlError_Description(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQmlError) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QQmlError_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlError) Line() int {
	if ptr.Pointer() != nil {
		return int(C.QQmlError_Line(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlError) Object() *core.QObject {
	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QQmlError_Object(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlError) SetColumn(column int) {
	if ptr.Pointer() != nil {
		C.QQmlError_SetColumn(ptr.Pointer(), C.int(column))
	}
}

func (ptr *QQmlError) SetDescription(description string) {
	if ptr.Pointer() != nil {
		C.QQmlError_SetDescription(ptr.Pointer(), C.CString(description))
	}
}

func (ptr *QQmlError) SetLine(line int) {
	if ptr.Pointer() != nil {
		C.QQmlError_SetLine(ptr.Pointer(), C.int(line))
	}
}

func (ptr *QQmlError) SetObject(object core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlError_SetObject(ptr.Pointer(), core.PointerFromQObject(object))
	}
}

func (ptr *QQmlError) SetUrl(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlError_SetUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QQmlError) ToString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlError_ToString(ptr.Pointer()))
	}
	return ""
}
