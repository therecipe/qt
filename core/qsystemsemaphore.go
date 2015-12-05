package core

//#include "core.h"
import "C"
import (
	"log"
	"unsafe"
)

type QSystemSemaphore struct {
	ptr unsafe.Pointer
}

type QSystemSemaphore_ITF interface {
	QSystemSemaphore_PTR() *QSystemSemaphore
}

func (p *QSystemSemaphore) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSystemSemaphore) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSystemSemaphore(ptr QSystemSemaphore_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSystemSemaphore_PTR().Pointer()
	}
	return nil
}

func NewQSystemSemaphoreFromPointer(ptr unsafe.Pointer) *QSystemSemaphore {
	var n = new(QSystemSemaphore)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSystemSemaphore) QSystemSemaphore_PTR() *QSystemSemaphore {
	return ptr
}

//QSystemSemaphore::AccessMode
type QSystemSemaphore__AccessMode int64

const (
	QSystemSemaphore__Open   = QSystemSemaphore__AccessMode(0)
	QSystemSemaphore__Create = QSystemSemaphore__AccessMode(1)
)

//QSystemSemaphore::SystemSemaphoreError
type QSystemSemaphore__SystemSemaphoreError int64

const (
	QSystemSemaphore__NoError          = QSystemSemaphore__SystemSemaphoreError(0)
	QSystemSemaphore__PermissionDenied = QSystemSemaphore__SystemSemaphoreError(1)
	QSystemSemaphore__KeyError         = QSystemSemaphore__SystemSemaphoreError(2)
	QSystemSemaphore__AlreadyExists    = QSystemSemaphore__SystemSemaphoreError(3)
	QSystemSemaphore__NotFound         = QSystemSemaphore__SystemSemaphoreError(4)
	QSystemSemaphore__OutOfResources   = QSystemSemaphore__SystemSemaphoreError(5)
	QSystemSemaphore__UnknownError     = QSystemSemaphore__SystemSemaphoreError(6)
)

func NewQSystemSemaphore(key string, initialValue int, mode QSystemSemaphore__AccessMode) *QSystemSemaphore {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSystemSemaphore::QSystemSemaphore")
		}
	}()

	return NewQSystemSemaphoreFromPointer(C.QSystemSemaphore_NewQSystemSemaphore(C.CString(key), C.int(initialValue), C.int(mode)))
}

func (ptr *QSystemSemaphore) Acquire() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSystemSemaphore::acquire")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSystemSemaphore_Acquire(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSystemSemaphore) Error() QSystemSemaphore__SystemSemaphoreError {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSystemSemaphore::error")
		}
	}()

	if ptr.Pointer() != nil {
		return QSystemSemaphore__SystemSemaphoreError(C.QSystemSemaphore_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSystemSemaphore) ErrorString() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSystemSemaphore::errorString")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QSystemSemaphore_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSystemSemaphore) Key() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSystemSemaphore::key")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QSystemSemaphore_Key(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSystemSemaphore) Release(n int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSystemSemaphore::release")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSystemSemaphore_Release(ptr.Pointer(), C.int(n)) != 0
	}
	return false
}

func (ptr *QSystemSemaphore) SetKey(key string, initialValue int, mode QSystemSemaphore__AccessMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSystemSemaphore::setKey")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSystemSemaphore_SetKey(ptr.Pointer(), C.CString(key), C.int(initialValue), C.int(mode))
	}
}

func (ptr *QSystemSemaphore) DestroyQSystemSemaphore() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSystemSemaphore::~QSystemSemaphore")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSystemSemaphore_DestroyQSystemSemaphore(ptr.Pointer())
	}
}
