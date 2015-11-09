package core

//#include "qmutexlocker.h"
import "C"
import (
	"unsafe"
)

type QMutexLocker struct {
	ptr unsafe.Pointer
}

type QMutexLocker_ITF interface {
	QMutexLocker_PTR() *QMutexLocker
}

func (p *QMutexLocker) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMutexLocker) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMutexLocker(ptr QMutexLocker_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMutexLocker_PTR().Pointer()
	}
	return nil
}

func NewQMutexLockerFromPointer(ptr unsafe.Pointer) *QMutexLocker {
	var n = new(QMutexLocker)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMutexLocker) QMutexLocker_PTR() *QMutexLocker {
	return ptr
}

func NewQMutexLocker(mutex QMutex_ITF) *QMutexLocker {
	return NewQMutexLockerFromPointer(C.QMutexLocker_NewQMutexLocker(PointerFromQMutex(mutex)))
}

func (ptr *QMutexLocker) Mutex() *QMutex {
	if ptr.Pointer() != nil {
		return NewQMutexFromPointer(C.QMutexLocker_Mutex(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMutexLocker) Relock() {
	if ptr.Pointer() != nil {
		C.QMutexLocker_Relock(ptr.Pointer())
	}
}

func (ptr *QMutexLocker) Unlock() {
	if ptr.Pointer() != nil {
		C.QMutexLocker_Unlock(ptr.Pointer())
	}
}

func (ptr *QMutexLocker) DestroyQMutexLocker() {
	if ptr.Pointer() != nil {
		C.QMutexLocker_DestroyQMutexLocker(ptr.Pointer())
	}
}
