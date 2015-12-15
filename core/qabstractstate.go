package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
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
		n.SetObjectName("QAbstractState_" + qt.Identifier())
	}
	return n
}

func (ptr *QAbstractState) QAbstractState_PTR() *QAbstractState {
	return ptr
}

func (ptr *QAbstractState) Active() bool {
	defer qt.Recovering("QAbstractState::active")

	if ptr.Pointer() != nil {
		return C.QAbstractState_Active(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractState) ConnectActiveChanged(f func(active bool)) {
	defer qt.Recovering("connect QAbstractState::activeChanged")

	if ptr.Pointer() != nil {
		C.QAbstractState_ConnectActiveChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "activeChanged", f)
	}
}

func (ptr *QAbstractState) DisconnectActiveChanged() {
	defer qt.Recovering("disconnect QAbstractState::activeChanged")

	if ptr.Pointer() != nil {
		C.QAbstractState_DisconnectActiveChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "activeChanged")
	}
}

//export callbackQAbstractStateActiveChanged
func callbackQAbstractStateActiveChanged(ptrName *C.char, active C.int) {
	defer qt.Recovering("callback QAbstractState::activeChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "activeChanged")
	if signal != nil {
		signal.(func(bool))(int(active) != 0)
	}

}

func (ptr *QAbstractState) ConnectEntered(f func()) {
	defer qt.Recovering("connect QAbstractState::entered")

	if ptr.Pointer() != nil {
		C.QAbstractState_ConnectEntered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "entered", f)
	}
}

func (ptr *QAbstractState) DisconnectEntered() {
	defer qt.Recovering("disconnect QAbstractState::entered")

	if ptr.Pointer() != nil {
		C.QAbstractState_DisconnectEntered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "entered")
	}
}

//export callbackQAbstractStateEntered
func callbackQAbstractStateEntered(ptrName *C.char) {
	defer qt.Recovering("callback QAbstractState::entered")

	var signal = qt.GetSignal(C.GoString(ptrName), "entered")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractState) ConnectExited(f func()) {
	defer qt.Recovering("connect QAbstractState::exited")

	if ptr.Pointer() != nil {
		C.QAbstractState_ConnectExited(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "exited", f)
	}
}

func (ptr *QAbstractState) DisconnectExited() {
	defer qt.Recovering("disconnect QAbstractState::exited")

	if ptr.Pointer() != nil {
		C.QAbstractState_DisconnectExited(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "exited")
	}
}

//export callbackQAbstractStateExited
func callbackQAbstractStateExited(ptrName *C.char) {
	defer qt.Recovering("callback QAbstractState::exited")

	var signal = qt.GetSignal(C.GoString(ptrName), "exited")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractState) Machine() *QStateMachine {
	defer qt.Recovering("QAbstractState::machine")

	if ptr.Pointer() != nil {
		return NewQStateMachineFromPointer(C.QAbstractState_Machine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractState) ParentState() *QState {
	defer qt.Recovering("QAbstractState::parentState")

	if ptr.Pointer() != nil {
		return NewQStateFromPointer(C.QAbstractState_ParentState(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractState) DestroyQAbstractState() {
	defer qt.Recovering("QAbstractState::~QAbstractState")

	if ptr.Pointer() != nil {
		C.QAbstractState_DestroyQAbstractState(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
