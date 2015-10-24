package core

//#include "qsemaphore.h"
import "C"
import (
	"unsafe"
)

type QSemaphore struct {
	ptr unsafe.Pointer
}

type QSemaphoreITF interface {
	QSemaphorePTR() *QSemaphore
}

func (p *QSemaphore) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSemaphore) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSemaphore(ptr QSemaphoreITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSemaphorePTR().Pointer()
	}
	return nil
}

func QSemaphoreFromPointer(ptr unsafe.Pointer) *QSemaphore {
	var n = new(QSemaphore)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSemaphore) QSemaphorePTR() *QSemaphore {
	return ptr
}

func NewQSemaphore(n int) *QSemaphore {
	return QSemaphoreFromPointer(unsafe.Pointer(C.QSemaphore_NewQSemaphore(C.int(n))))
}

func (ptr *QSemaphore) Acquire(n int) {
	if ptr.Pointer() != nil {
		C.QSemaphore_Acquire(C.QtObjectPtr(ptr.Pointer()), C.int(n))
	}
}

func (ptr *QSemaphore) Available() int {
	if ptr.Pointer() != nil {
		return int(C.QSemaphore_Available(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSemaphore) Release(n int) {
	if ptr.Pointer() != nil {
		C.QSemaphore_Release(C.QtObjectPtr(ptr.Pointer()), C.int(n))
	}
}

func (ptr *QSemaphore) TryAcquire(n int) bool {
	if ptr.Pointer() != nil {
		return C.QSemaphore_TryAcquire(C.QtObjectPtr(ptr.Pointer()), C.int(n)) != 0
	}
	return false
}

func (ptr *QSemaphore) TryAcquire2(n int, timeout int) bool {
	if ptr.Pointer() != nil {
		return C.QSemaphore_TryAcquire2(C.QtObjectPtr(ptr.Pointer()), C.int(n), C.int(timeout)) != 0
	}
	return false
}

func (ptr *QSemaphore) DestroyQSemaphore() {
	if ptr.Pointer() != nil {
		C.QSemaphore_DestroyQSemaphore(C.QtObjectPtr(ptr.Pointer()))
	}
}
