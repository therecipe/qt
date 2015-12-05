package qml

//#include "qml.h"
import "C"
import (
	"log"
	"unsafe"
)

type QQmlIncubationController struct {
	ptr unsafe.Pointer
}

type QQmlIncubationController_ITF interface {
	QQmlIncubationController_PTR() *QQmlIncubationController
}

func (p *QQmlIncubationController) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QQmlIncubationController) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQQmlIncubationController(ptr QQmlIncubationController_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlIncubationController_PTR().Pointer()
	}
	return nil
}

func NewQQmlIncubationControllerFromPointer(ptr unsafe.Pointer) *QQmlIncubationController {
	var n = new(QQmlIncubationController)
	n.SetPointer(ptr)
	return n
}

func (ptr *QQmlIncubationController) QQmlIncubationController_PTR() *QQmlIncubationController {
	return ptr
}

func NewQQmlIncubationController() *QQmlIncubationController {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlIncubationController::QQmlIncubationController")
		}
	}()

	return NewQQmlIncubationControllerFromPointer(C.QQmlIncubationController_NewQQmlIncubationController())
}

func (ptr *QQmlIncubationController) Engine() *QQmlEngine {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlIncubationController::engine")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQQmlEngineFromPointer(C.QQmlIncubationController_Engine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlIncubationController) IncubateFor(msecs int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlIncubationController::incubateFor")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQmlIncubationController_IncubateFor(ptr.Pointer(), C.int(msecs))
	}
}

func (ptr *QQmlIncubationController) IncubatingObjectCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlIncubationController::incubatingObjectCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QQmlIncubationController_IncubatingObjectCount(ptr.Pointer()))
	}
	return 0
}
