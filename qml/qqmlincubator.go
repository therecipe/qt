package qml

//#include "qqmlincubator.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QQmlIncubator struct {
	ptr unsafe.Pointer
}

type QQmlIncubatorITF interface {
	QQmlIncubatorPTR() *QQmlIncubator
}

func (p *QQmlIncubator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QQmlIncubator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQQmlIncubator(ptr QQmlIncubatorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlIncubatorPTR().Pointer()
	}
	return nil
}

func QQmlIncubatorFromPointer(ptr unsafe.Pointer) *QQmlIncubator {
	var n = new(QQmlIncubator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QQmlIncubator) QQmlIncubatorPTR() *QQmlIncubator {
	return ptr
}

//QQmlIncubator::IncubationMode
type QQmlIncubator__IncubationMode int

var (
	QQmlIncubator__Asynchronous         = QQmlIncubator__IncubationMode(0)
	QQmlIncubator__AsynchronousIfNested = QQmlIncubator__IncubationMode(1)
	QQmlIncubator__Synchronous          = QQmlIncubator__IncubationMode(2)
)

//QQmlIncubator::Status
type QQmlIncubator__Status int

var (
	QQmlIncubator__Null    = QQmlIncubator__Status(0)
	QQmlIncubator__Ready   = QQmlIncubator__Status(1)
	QQmlIncubator__Loading = QQmlIncubator__Status(2)
	QQmlIncubator__Error   = QQmlIncubator__Status(3)
)

func NewQQmlIncubator(mode QQmlIncubator__IncubationMode) *QQmlIncubator {
	return QQmlIncubatorFromPointer(unsafe.Pointer(C.QQmlIncubator_NewQQmlIncubator(C.int(mode))))
}

func (ptr *QQmlIncubator) Clear() {
	if ptr.Pointer() != nil {
		C.QQmlIncubator_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QQmlIncubator) ForceCompletion() {
	if ptr.Pointer() != nil {
		C.QQmlIncubator_ForceCompletion(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QQmlIncubator) IncubationMode() QQmlIncubator__IncubationMode {
	if ptr.Pointer() != nil {
		return QQmlIncubator__IncubationMode(C.QQmlIncubator_IncubationMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QQmlIncubator) IsError() bool {
	if ptr.Pointer() != nil {
		return C.QQmlIncubator_IsError(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlIncubator) IsLoading() bool {
	if ptr.Pointer() != nil {
		return C.QQmlIncubator_IsLoading(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlIncubator) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QQmlIncubator_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlIncubator) IsReady() bool {
	if ptr.Pointer() != nil {
		return C.QQmlIncubator_IsReady(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlIncubator) Object() *core.QObject {
	if ptr.Pointer() != nil {
		return core.QObjectFromPointer(unsafe.Pointer(C.QQmlIncubator_Object(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QQmlIncubator) Status() QQmlIncubator__Status {
	if ptr.Pointer() != nil {
		return QQmlIncubator__Status(C.QQmlIncubator_Status(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}
