package core

//#include "qreadlocker.h"
import "C"
import (
	"unsafe"
)

type QReadLocker struct {
	ptr unsafe.Pointer
}

type QReadLockerITF interface {
	QReadLockerPTR() *QReadLocker
}

func (p *QReadLocker) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QReadLocker) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQReadLocker(ptr QReadLockerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QReadLockerPTR().Pointer()
	}
	return nil
}

func QReadLockerFromPointer(ptr unsafe.Pointer) *QReadLocker {
	var n = new(QReadLocker)
	n.SetPointer(ptr)
	return n
}

func (ptr *QReadLocker) QReadLockerPTR() *QReadLocker {
	return ptr
}

func NewQReadLocker(lock QReadWriteLockITF) *QReadLocker {
	return QReadLockerFromPointer(unsafe.Pointer(C.QReadLocker_NewQReadLocker(C.QtObjectPtr(PointerFromQReadWriteLock(lock)))))
}

func (ptr *QReadLocker) ReadWriteLock() *QReadWriteLock {
	if ptr.Pointer() != nil {
		return QReadWriteLockFromPointer(unsafe.Pointer(C.QReadLocker_ReadWriteLock(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QReadLocker) Relock() {
	if ptr.Pointer() != nil {
		C.QReadLocker_Relock(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QReadLocker) Unlock() {
	if ptr.Pointer() != nil {
		C.QReadLocker_Unlock(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QReadLocker) DestroyQReadLocker() {
	if ptr.Pointer() != nil {
		C.QReadLocker_DestroyQReadLocker(C.QtObjectPtr(ptr.Pointer()))
	}
}
