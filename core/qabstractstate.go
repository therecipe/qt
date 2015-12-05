package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"unsafe"
)

type QAbstractState struct {
	QObject
}

type QAbstractState_ITF interface {
	QObject_ITF
	QAbstractState_PTR() *QAbstractState
}

func PointerFromQAbstractState(ptr QAbstractState_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractState_PTR().Pointer()
	}
	return nil
}

func NewQAbstractStateFromPointer(ptr unsafe.Pointer) *QAbstractState {
	var n = new(QAbstractState)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAbstractState_") {
		n.SetObjectName("QAbstractState_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAbstractState) QAbstractState_PTR() *QAbstractState {
	return ptr
}

func (ptr *QAbstractState) Active() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractState::active")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAbstractState_Active(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractState) ConnectActiveChanged(f func(active bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractState::activeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractState_ConnectActiveChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "activeChanged", f)
	}
}

func (ptr *QAbstractState) DisconnectActiveChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractState::activeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractState_DisconnectActiveChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "activeChanged")
	}
}

//export callbackQAbstractStateActiveChanged
func callbackQAbstractStateActiveChanged(ptrName *C.char, active C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractState::activeChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "activeChanged").(func(bool))(int(active) != 0)
}

func (ptr *QAbstractState) ConnectEntered(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractState::entered")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractState_ConnectEntered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "entered", f)
	}
}

func (ptr *QAbstractState) DisconnectEntered() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractState::entered")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractState_DisconnectEntered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "entered")
	}
}

//export callbackQAbstractStateEntered
func callbackQAbstractStateEntered(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractState::entered")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "entered").(func())()
}

func (ptr *QAbstractState) ConnectExited(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractState::exited")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractState_ConnectExited(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "exited", f)
	}
}

func (ptr *QAbstractState) DisconnectExited() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractState::exited")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractState_DisconnectExited(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "exited")
	}
}

//export callbackQAbstractStateExited
func callbackQAbstractStateExited(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractState::exited")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "exited").(func())()
}

func (ptr *QAbstractState) Machine() *QStateMachine {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractState::machine")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQStateMachineFromPointer(C.QAbstractState_Machine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractState) ParentState() *QState {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractState::parentState")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQStateFromPointer(C.QAbstractState_ParentState(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractState) DestroyQAbstractState() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractState::~QAbstractState")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractState_DestroyQAbstractState(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
