package core

//#include "qreadwritelock.h"
import "C"
import (
	"unsafe"
)

type QReadWriteLock struct {
	ptr unsafe.Pointer
}

type QReadWriteLockITF interface {
	QReadWriteLockPTR() *QReadWriteLock
}

func (p *QReadWriteLock) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QReadWriteLock) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQReadWriteLock(ptr QReadWriteLockITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QReadWriteLockPTR().Pointer()
	}
	return nil
}

func QReadWriteLockFromPointer(ptr unsafe.Pointer) *QReadWriteLock {
	var n = new(QReadWriteLock)
	n.SetPointer(ptr)
	return n
}

func (ptr *QReadWriteLock) QReadWriteLockPTR() *QReadWriteLock {
	return ptr
}

//QReadWriteLock::RecursionMode
type QReadWriteLock__RecursionMode int

var (
	QReadWriteLock__NonRecursive = QReadWriteLock__RecursionMode(0)
	QReadWriteLock__Recursive    = QReadWriteLock__RecursionMode(1)
)

func NewQReadWriteLock(recursionMode QReadWriteLock__RecursionMode) *QReadWriteLock {
	return QReadWriteLockFromPointer(unsafe.Pointer(C.QReadWriteLock_NewQReadWriteLock(C.int(recursionMode))))
}

func (ptr *QReadWriteLock) LockForRead() {
	if ptr.Pointer() != nil {
		C.QReadWriteLock_LockForRead(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QReadWriteLock) LockForWrite() {
	if ptr.Pointer() != nil {
		C.QReadWriteLock_LockForWrite(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QReadWriteLock) TryLockForRead() bool {
	if ptr.Pointer() != nil {
		return C.QReadWriteLock_TryLockForRead(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QReadWriteLock) TryLockForRead2(timeout int) bool {
	if ptr.Pointer() != nil {
		return C.QReadWriteLock_TryLockForRead2(C.QtObjectPtr(ptr.Pointer()), C.int(timeout)) != 0
	}
	return false
}

func (ptr *QReadWriteLock) TryLockForWrite() bool {
	if ptr.Pointer() != nil {
		return C.QReadWriteLock_TryLockForWrite(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QReadWriteLock) TryLockForWrite2(timeout int) bool {
	if ptr.Pointer() != nil {
		return C.QReadWriteLock_TryLockForWrite2(C.QtObjectPtr(ptr.Pointer()), C.int(timeout)) != 0
	}
	return false
}

func (ptr *QReadWriteLock) Unlock() {
	if ptr.Pointer() != nil {
		C.QReadWriteLock_Unlock(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QReadWriteLock) DestroyQReadWriteLock() {
	if ptr.Pointer() != nil {
		C.QReadWriteLock_DestroyQReadWriteLock(C.QtObjectPtr(ptr.Pointer()))
	}
}
