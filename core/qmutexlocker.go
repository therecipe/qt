package core

//#include "core.h"
import "C"
import (
	"log"
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMutexLocker::QMutexLocker")
		}
	}()

	return NewQMutexLockerFromPointer(C.QMutexLocker_NewQMutexLocker(PointerFromQMutex(mutex)))
}

func (ptr *QMutexLocker) Mutex() *QMutex {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMutexLocker::mutex")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQMutexFromPointer(C.QMutexLocker_Mutex(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMutexLocker) Relock() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMutexLocker::relock")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMutexLocker_Relock(ptr.Pointer())
	}
}

func (ptr *QMutexLocker) Unlock() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMutexLocker::unlock")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMutexLocker_Unlock(ptr.Pointer())
	}
}

func (ptr *QMutexLocker) DestroyQMutexLocker() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMutexLocker::~QMutexLocker")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMutexLocker_DestroyQMutexLocker(ptr.Pointer())
	}
}
