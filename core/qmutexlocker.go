package core

//#include "qmutexlocker.h"
import "C"
import (
	"unsafe"
)

type QMutexLocker struct {
	ptr unsafe.Pointer
}

type QMutexLockerITF interface {
	QMutexLockerPTR() *QMutexLocker
}

func (p *QMutexLocker) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMutexLocker) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMutexLocker(ptr QMutexLockerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMutexLockerPTR().Pointer()
	}
	return nil
}

func QMutexLockerFromPointer(ptr unsafe.Pointer) *QMutexLocker {
	var n = new(QMutexLocker)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMutexLocker) QMutexLockerPTR() *QMutexLocker {
	return ptr
}

func NewQMutexLocker(mutex QMutexITF) *QMutexLocker {
	return QMutexLockerFromPointer(unsafe.Pointer(C.QMutexLocker_NewQMutexLocker(C.QtObjectPtr(PointerFromQMutex(mutex)))))
}

func (ptr *QMutexLocker) Mutex() *QMutex {
	if ptr.Pointer() != nil {
		return QMutexFromPointer(unsafe.Pointer(C.QMutexLocker_Mutex(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QMutexLocker) Relock() {
	if ptr.Pointer() != nil {
		C.QMutexLocker_Relock(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QMutexLocker) Unlock() {
	if ptr.Pointer() != nil {
		C.QMutexLocker_Unlock(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QMutexLocker) DestroyQMutexLocker() {
	if ptr.Pointer() != nil {
		C.QMutexLocker_DestroyQMutexLocker(C.QtObjectPtr(ptr.Pointer()))
	}
}
