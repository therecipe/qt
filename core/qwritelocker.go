package core

//#include "qwritelocker.h"
import "C"
import (
	"unsafe"
)

type QWriteLocker struct {
	ptr unsafe.Pointer
}

type QWriteLockerITF interface {
	QWriteLockerPTR() *QWriteLocker
}

func (p *QWriteLocker) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QWriteLocker) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQWriteLocker(ptr QWriteLockerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWriteLockerPTR().Pointer()
	}
	return nil
}

func QWriteLockerFromPointer(ptr unsafe.Pointer) *QWriteLocker {
	var n = new(QWriteLocker)
	n.SetPointer(ptr)
	return n
}

func (ptr *QWriteLocker) QWriteLockerPTR() *QWriteLocker {
	return ptr
}

func NewQWriteLocker(lock QReadWriteLockITF) *QWriteLocker {
	return QWriteLockerFromPointer(unsafe.Pointer(C.QWriteLocker_NewQWriteLocker(C.QtObjectPtr(PointerFromQReadWriteLock(lock)))))
}

func (ptr *QWriteLocker) ReadWriteLock() *QReadWriteLock {
	if ptr.Pointer() != nil {
		return QReadWriteLockFromPointer(unsafe.Pointer(C.QWriteLocker_ReadWriteLock(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QWriteLocker) Relock() {
	if ptr.Pointer() != nil {
		C.QWriteLocker_Relock(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWriteLocker) Unlock() {
	if ptr.Pointer() != nil {
		C.QWriteLocker_Unlock(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWriteLocker) DestroyQWriteLocker() {
	if ptr.Pointer() != nil {
		C.QWriteLocker_DestroyQWriteLocker(C.QtObjectPtr(ptr.Pointer()))
	}
}
