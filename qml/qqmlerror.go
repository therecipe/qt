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

type QQmlErrorITF interface {
	QQmlErrorPTR() *QQmlError
}

func (p *QQmlError) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QQmlError) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQQmlError(ptr QQmlErrorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlErrorPTR().Pointer()
	}
	return nil
}

func QQmlErrorFromPointer(ptr unsafe.Pointer) *QQmlError {
	var n = new(QQmlError)
	n.SetPointer(ptr)
	return n
}

func (ptr *QQmlError) QQmlErrorPTR() *QQmlError {
	return ptr
}

func NewQQmlError() *QQmlError {
	return QQmlErrorFromPointer(unsafe.Pointer(C.QQmlError_NewQQmlError()))
}

func NewQQmlError2(other QQmlErrorITF) *QQmlError {
	return QQmlErrorFromPointer(unsafe.Pointer(C.QQmlError_NewQQmlError2(C.QtObjectPtr(PointerFromQQmlError(other)))))
}

func (ptr *QQmlError) Column() int {
	if ptr.Pointer() != nil {
		return int(C.QQmlError_Column(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QQmlError) Description() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlError_Description(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QQmlError) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QQmlError_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlError) Line() int {
	if ptr.Pointer() != nil {
		return int(C.QQmlError_Line(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QQmlError) Object() *core.QObject {
	if ptr.Pointer() != nil {
		return core.QObjectFromPointer(unsafe.Pointer(C.QQmlError_Object(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QQmlError) SetColumn(column int) {
	if ptr.Pointer() != nil {
		C.QQmlError_SetColumn(C.QtObjectPtr(ptr.Pointer()), C.int(column))
	}
}

func (ptr *QQmlError) SetDescription(description string) {
	if ptr.Pointer() != nil {
		C.QQmlError_SetDescription(C.QtObjectPtr(ptr.Pointer()), C.CString(description))
	}
}

func (ptr *QQmlError) SetLine(line int) {
	if ptr.Pointer() != nil {
		C.QQmlError_SetLine(C.QtObjectPtr(ptr.Pointer()), C.int(line))
	}
}

func (ptr *QQmlError) SetObject(object core.QObjectITF) {
	if ptr.Pointer() != nil {
		C.QQmlError_SetObject(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQObject(object)))
	}
}

func (ptr *QQmlError) SetUrl(url string) {
	if ptr.Pointer() != nil {
		C.QQmlError_SetUrl(C.QtObjectPtr(ptr.Pointer()), C.CString(url))
	}
}

func (ptr *QQmlError) ToString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlError_ToString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QQmlError) Url() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlError_Url(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}
