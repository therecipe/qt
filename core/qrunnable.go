package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"unsafe"
)

type QRunnable struct {
	ptr unsafe.Pointer
}

type QRunnable_ITF interface {
	QRunnable_PTR() *QRunnable
}

func (p *QRunnable) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QRunnable) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQRunnable(ptr QRunnable_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRunnable_PTR().Pointer()
	}
	return nil
}

func NewQRunnableFromPointer(ptr unsafe.Pointer) *QRunnable {
	var n = new(QRunnable)
	n.SetPointer(ptr)
	return n
}

func (ptr *QRunnable) QRunnable_PTR() *QRunnable {
	return ptr
}

func (ptr *QRunnable) AutoDelete() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRunnable::autoDelete")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QRunnable_AutoDelete(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QRunnable) Run() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRunnable::run")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRunnable_Run(ptr.Pointer())
	}
}

func (ptr *QRunnable) SetAutoDelete(autoDelete bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRunnable::setAutoDelete")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRunnable_SetAutoDelete(ptr.Pointer(), C.int(qt.GoBoolToInt(autoDelete)))
	}
}

func (ptr *QRunnable) DestroyQRunnable() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRunnable::~QRunnable")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRunnable_DestroyQRunnable(ptr.Pointer())
	}
}
