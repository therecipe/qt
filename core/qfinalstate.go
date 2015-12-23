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

	if signal := qt.GetSignal(C.GoString(ptrName), "onEntry"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
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

	if signal := qt.GetSignal(C.GoString(ptrName), "onExit"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
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

func (ptr *QFinalState) ConnectTimerEvent(f func(event *QTimerEvent)) {
	defer qt.Recovering("connect QFinalState::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QFinalState) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QFinalState::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQFinalStateTimerEvent
func callbackQFinalStateTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFinalState::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*QTimerEvent))(NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFinalState) ConnectChildEvent(f func(event *QChildEvent)) {
	defer qt.Recovering("connect QFinalState::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QFinalState) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QFinalState::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQFinalStateChildEvent
func callbackQFinalStateChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFinalState::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*QChildEvent))(NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFinalState) ConnectCustomEvent(f func(event *QEvent)) {
	defer qt.Recovering("connect QFinalState::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QFinalState) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QFinalState::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQFinalStateCustomEvent
func callbackQFinalStateCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFinalState::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
		return true
	}
	return false

}
