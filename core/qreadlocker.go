package core

//#include "core.h"
import "C"
import (
	"log"
	"unsafe"
)

type QReadLocker struct {
	ptr unsafe.Pointer
}

type QReadLocker_ITF interface {
	QReadLocker_PTR() *QReadLocker
}

func (p *QReadLocker) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QReadLocker) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQReadLocker(ptr QReadLocker_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QReadLocker_PTR().Pointer()
	}
	return nil
}

func NewQReadLockerFromPointer(ptr unsafe.Pointer) *QReadLocker {
	var n = new(QReadLocker)
	n.SetPointer(ptr)
	return n
}

func (ptr *QReadLocker) QReadLocker_PTR() *QReadLocker {
	return ptr
}

func NewQReadLocker(lock QReadWriteLock_ITF) *QReadLocker {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QReadLocker::QReadLocker")
		}
	}()

	return NewQReadLockerFromPointer(C.QReadLocker_NewQReadLocker(PointerFromQReadWriteLock(lock)))
}

func (ptr *QReadLocker) ReadWriteLock() *QReadWriteLock {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QReadLocker::readWriteLock")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQReadWriteLockFromPointer(C.QReadLocker_ReadWriteLock(ptr.Pointer()))
	}
	return nil
}

func (ptr *QReadLocker) Relock() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QReadLocker::relock")
		}
	}()

	if ptr.Pointer() != nil {
		C.QReadLocker_Relock(ptr.Pointer())
	}
}

func (ptr *QReadLocker) Unlock() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QReadLocker::unlock")
		}
	}()

	if ptr.Pointer() != nil {
		C.QReadLocker_Unlock(ptr.Pointer())
	}
}

func (ptr *QReadLocker) DestroyQReadLocker() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QReadLocker::~QReadLocker")
		}
	}()

	if ptr.Pointer() != nil {
		C.QReadLocker_DestroyQReadLocker(ptr.Pointer())
	}
}
