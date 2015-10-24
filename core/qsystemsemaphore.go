package core

//#include "qsystemsemaphore.h"
import "C"
import (
	"unsafe"
)

type QSystemSemaphore struct {
	ptr unsafe.Pointer
}

type QSystemSemaphoreITF interface {
	QSystemSemaphorePTR() *QSystemSemaphore
}

func (p *QSystemSemaphore) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSystemSemaphore) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSystemSemaphore(ptr QSystemSemaphoreITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSystemSemaphorePTR().Pointer()
	}
	return nil
}

func QSystemSemaphoreFromPointer(ptr unsafe.Pointer) *QSystemSemaphore {
	var n = new(QSystemSemaphore)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSystemSemaphore) QSystemSemaphorePTR() *QSystemSemaphore {
	return ptr
}

//QSystemSemaphore::AccessMode
type QSystemSemaphore__AccessMode int

var (
	QSystemSemaphore__Open   = QSystemSemaphore__AccessMode(0)
	QSystemSemaphore__Create = QSystemSemaphore__AccessMode(1)
)

//QSystemSemaphore::SystemSemaphoreError
type QSystemSemaphore__SystemSemaphoreError int

var (
	QSystemSemaphore__NoError          = QSystemSemaphore__SystemSemaphoreError(0)
	QSystemSemaphore__PermissionDenied = QSystemSemaphore__SystemSemaphoreError(1)
	QSystemSemaphore__KeyError         = QSystemSemaphore__SystemSemaphoreError(2)
	QSystemSemaphore__AlreadyExists    = QSystemSemaphore__SystemSemaphoreError(3)
	QSystemSemaphore__NotFound         = QSystemSemaphore__SystemSemaphoreError(4)
	QSystemSemaphore__OutOfResources   = QSystemSemaphore__SystemSemaphoreError(5)
	QSystemSemaphore__UnknownError     = QSystemSemaphore__SystemSemaphoreError(6)
)

func NewQSystemSemaphore(key string, initialValue int, mode QSystemSemaphore__AccessMode) *QSystemSemaphore {
	return QSystemSemaphoreFromPointer(unsafe.Pointer(C.QSystemSemaphore_NewQSystemSemaphore(C.CString(key), C.int(initialValue), C.int(mode))))
}

func (ptr *QSystemSemaphore) Acquire() bool {
	if ptr.Pointer() != nil {
		return C.QSystemSemaphore_Acquire(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSystemSemaphore) Error() QSystemSemaphore__SystemSemaphoreError {
	if ptr.Pointer() != nil {
		return QSystemSemaphore__SystemSemaphoreError(C.QSystemSemaphore_Error(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSystemSemaphore) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSystemSemaphore_ErrorString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSystemSemaphore) Key() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSystemSemaphore_Key(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSystemSemaphore) Release(n int) bool {
	if ptr.Pointer() != nil {
		return C.QSystemSemaphore_Release(C.QtObjectPtr(ptr.Pointer()), C.int(n)) != 0
	}
	return false
}

func (ptr *QSystemSemaphore) SetKey(key string, initialValue int, mode QSystemSemaphore__AccessMode) {
	if ptr.Pointer() != nil {
		C.QSystemSemaphore_SetKey(C.QtObjectPtr(ptr.Pointer()), C.CString(key), C.int(initialValue), C.int(mode))
	}
}

func (ptr *QSystemSemaphore) DestroyQSystemSemaphore() {
	if ptr.Pointer() != nil {
		C.QSystemSemaphore_DestroyQSystemSemaphore(C.QtObjectPtr(ptr.Pointer()))
	}
}
