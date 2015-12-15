package qml

//#include "qml.h"
import "C"
import (
	"github.com/therecipe/qt"
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
	for len(n.ObjectNameAbs()) < len("QQmlIncubationController_") {
		n.SetObjectNameAbs("QQmlIncubationController_" + qt.Identifier())
	}
	return n
}

func (ptr *QQmlIncubationController) QQmlIncubationController_PTR() *QQmlIncubationController {
	return ptr
}

func NewQQmlIncubationController() *QQmlIncubationController {
	defer qt.Recovering("QQmlIncubationController::QQmlIncubationController")

	return NewQQmlIncubationControllerFromPointer(C.QQmlIncubationController_NewQQmlIncubationController())
}

func (ptr *QQmlIncubationController) Engine() *QQmlEngine {
	defer qt.Recovering("QQmlIncubationController::engine")

	if ptr.Pointer() != nil {
		return NewQQmlEngineFromPointer(C.QQmlIncubationController_Engine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlIncubationController) IncubateFor(msecs int) {
	defer qt.Recovering("QQmlIncubationController::incubateFor")

	if ptr.Pointer() != nil {
		C.QQmlIncubationController_IncubateFor(ptr.Pointer(), C.int(msecs))
	}
}

func (ptr *QQmlIncubationController) IncubatingObjectCount() int {
	defer qt.Recovering("QQmlIncubationController::incubatingObjectCount")

	if ptr.Pointer() != nil {
		return int(C.QQmlIncubationController_IncubatingObjectCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlIncubationController) ConnectIncubatingObjectCountChanged(f func(incubatingObjectCount int)) {
	defer qt.Recovering("connect QQmlIncubationController::incubatingObjectCountChanged")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "incubatingObjectCountChanged", f)
	}
}

func (ptr *QQmlIncubationController) DisconnectIncubatingObjectCountChanged() {
	defer qt.Recovering("disconnect QQmlIncubationController::incubatingObjectCountChanged")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "incubatingObjectCountChanged")
	}
}

//export callbackQQmlIncubationControllerIncubatingObjectCountChanged
func callbackQQmlIncubationControllerIncubatingObjectCountChanged(ptrName *C.char, incubatingObjectCount C.int) bool {
	defer qt.Recovering("callback QQmlIncubationController::incubatingObjectCountChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "incubatingObjectCountChanged")
	if signal != nil {
		defer signal.(func(int))(int(incubatingObjectCount))
		return true
	}
	return false

}

func (ptr *QQmlIncubationController) ObjectNameAbs() string {
	defer qt.Recovering("QQmlIncubationController::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlIncubationController_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQmlIncubationController) SetObjectNameAbs(name string) {
	defer qt.Recovering("QQmlIncubationController::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QQmlIncubationController_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
