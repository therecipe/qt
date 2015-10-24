package qml

//#include "qqmlincubationcontroller.h"
import "C"
import (
	"unsafe"
)

type QQmlIncubationController struct {
	ptr unsafe.Pointer
}

type QQmlIncubationControllerITF interface {
	QQmlIncubationControllerPTR() *QQmlIncubationController
}

func (p *QQmlIncubationController) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QQmlIncubationController) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQQmlIncubationController(ptr QQmlIncubationControllerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlIncubationControllerPTR().Pointer()
	}
	return nil
}

func QQmlIncubationControllerFromPointer(ptr unsafe.Pointer) *QQmlIncubationController {
	var n = new(QQmlIncubationController)
	n.SetPointer(ptr)
	return n
}

func (ptr *QQmlIncubationController) QQmlIncubationControllerPTR() *QQmlIncubationController {
	return ptr
}

func NewQQmlIncubationController() *QQmlIncubationController {
	return QQmlIncubationControllerFromPointer(unsafe.Pointer(C.QQmlIncubationController_NewQQmlIncubationController()))
}

func (ptr *QQmlIncubationController) Engine() *QQmlEngine {
	if ptr.Pointer() != nil {
		return QQmlEngineFromPointer(unsafe.Pointer(C.QQmlIncubationController_Engine(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QQmlIncubationController) IncubateFor(msecs int) {
	if ptr.Pointer() != nil {
		C.QQmlIncubationController_IncubateFor(C.QtObjectPtr(ptr.Pointer()), C.int(msecs))
	}
}

func (ptr *QQmlIncubationController) IncubatingObjectCount() int {
	if ptr.Pointer() != nil {
		return int(C.QQmlIncubationController_IncubatingObjectCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}
