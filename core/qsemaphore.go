package core

//#include "core.h"
import "C"
import (
	"log"
	"unsafe"
)

type QSemaphore struct {
	ptr unsafe.Pointer
}

type QSemaphore_ITF interface {
	QSemaphore_PTR() *QSemaphore
}

func (p *QSemaphore) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSemaphore) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSemaphore(ptr QSemaphore_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSemaphore_PTR().Pointer()
	}
	return nil
}

func NewQSemaphoreFromPointer(ptr unsafe.Pointer) *QSemaphore {
	var n = new(QSemaphore)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSemaphore) QSemaphore_PTR() *QSemaphore {
	return ptr
}

func NewQSemaphore(n int) *QSemaphore {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSemaphore::QSemaphore")
		}
	}()

	return NewQSemaphoreFromPointer(C.QSemaphore_NewQSemaphore(C.int(n)))
}

func (ptr *QSemaphore) Acquire(n int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSemaphore::acquire")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSemaphore_Acquire(ptr.Pointer(), C.int(n))
	}
}

func (ptr *QSemaphore) Available() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSemaphore::available")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QSemaphore_Available(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSemaphore) Release(n int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSemaphore::release")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSemaphore_Release(ptr.Pointer(), C.int(n))
	}
}

func (ptr *QSemaphore) TryAcquire(n int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSemaphore::tryAcquire")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSemaphore_TryAcquire(ptr.Pointer(), C.int(n)) != 0
	}
	return false
}

func (ptr *QSemaphore) TryAcquire2(n int, timeout int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSemaphore::tryAcquire")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSemaphore_TryAcquire2(ptr.Pointer(), C.int(n), C.int(timeout)) != 0
	}
	return false
}

func (ptr *QSemaphore) DestroyQSemaphore() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSemaphore::~QSemaphore")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSemaphore_DestroyQSemaphore(ptr.Pointer())
	}
}
