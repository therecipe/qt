package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QFinalState struct {
	QAbstractState
}

type QFinalState_ITF interface {
	QAbstractState_ITF
	QFinalState_PTR() *QFinalState
}

func PointerFromQFinalState(ptr QFinalState_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFinalState_PTR().Pointer()
	}
	return nil
}

func NewQFinalStateFromPointer(ptr unsafe.Pointer) *QFinalState {
	var n = new(QFinalState)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QFinalState_") {
		n.SetObjectName("QFinalState_" + qt.Identifier())
	}
	return n
}

func (ptr *QFinalState) QFinalState_PTR() *QFinalState {
	return ptr
}

func NewQFinalState(parent QState_ITF) *QFinalState {
	defer qt.Recovering("QFinalState::QFinalState")

	return NewQFinalStateFromPointer(C.QFinalState_NewQFinalState(PointerFromQState(parent)))
}

func (ptr *QFinalState) ConnectOnEntry(f func(event *QEvent)) {
	defer qt.Recovering("connect QFinalState::onEntry")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "onEntry", f)
	}
}

func (ptr *QFinalState) DisconnectOnEntry() {
	defer qt.Recovering("disconnect QFinalState::onEntry")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "onEntry")
	}
}

//export callbackQFinalStateOnEntry
func callbackQFinalStateOnEntry(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFinalState::onEntry")

	var signal = qt.GetSignal(C.GoString(ptrName), "onEntry")
	if signal != nil {
		defer signal.(func(*QEvent))(NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFinalState) ConnectOnExit(f func(event *QEvent)) {
	defer qt.Recovering("connect QFinalState::onExit")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "onExit", f)
	}
}

func (ptr *QFinalState) DisconnectOnExit() {
	defer qt.Recovering("disconnect QFinalState::onExit")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "onExit")
	}
}

//export callbackQFinalStateOnExit
func callbackQFinalStateOnExit(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFinalState::onExit")

	var signal = qt.GetSignal(C.GoString(ptrName), "onExit")
	if signal != nil {
		defer signal.(func(*QEvent))(NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFinalState) DestroyQFinalState() {
	defer qt.Recovering("QFinalState::~QFinalState")

	if ptr.Pointer() != nil {
		C.QFinalState_DestroyQFinalState(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
