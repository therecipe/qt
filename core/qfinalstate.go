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

func (ptr *QFinalState) Event(e QEvent_ITF) bool {
	defer qt.Recovering("QFinalState::event")

	if ptr.Pointer() != nil {
		return C.QFinalState_Event(ptr.Pointer(), PointerFromQEvent(e)) != 0
	}
	return false
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
func callbackQFinalStateOnEntry(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFinalState::onEntry")

	if signal := qt.GetSignal(C.GoString(ptrName), "onEntry"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
	} else {
		NewQFinalStateFromPointer(ptr).OnEntryDefault(NewQEventFromPointer(event))
	}
}

func (ptr *QFinalState) OnEntry(event QEvent_ITF) {
	defer qt.Recovering("QFinalState::onEntry")

	if ptr.Pointer() != nil {
		C.QFinalState_OnEntry(ptr.Pointer(), PointerFromQEvent(event))
	}
}

func (ptr *QFinalState) OnEntryDefault(event QEvent_ITF) {
	defer qt.Recovering("QFinalState::onEntry")

	if ptr.Pointer() != nil {
		C.QFinalState_OnEntryDefault(ptr.Pointer(), PointerFromQEvent(event))
	}
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
func callbackQFinalStateOnExit(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFinalState::onExit")

	if signal := qt.GetSignal(C.GoString(ptrName), "onExit"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
	} else {
		NewQFinalStateFromPointer(ptr).OnExitDefault(NewQEventFromPointer(event))
	}
}

func (ptr *QFinalState) OnExit(event QEvent_ITF) {
	defer qt.Recovering("QFinalState::onExit")

	if ptr.Pointer() != nil {
		C.QFinalState_OnExit(ptr.Pointer(), PointerFromQEvent(event))
	}
}

func (ptr *QFinalState) OnExitDefault(event QEvent_ITF) {
	defer qt.Recovering("QFinalState::onExit")

	if ptr.Pointer() != nil {
		C.QFinalState_OnExitDefault(ptr.Pointer(), PointerFromQEvent(event))
	}
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
func callbackQFinalStateTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFinalState::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*QTimerEvent))(NewQTimerEventFromPointer(event))
	} else {
		NewQFinalStateFromPointer(ptr).TimerEventDefault(NewQTimerEventFromPointer(event))
	}
}

func (ptr *QFinalState) TimerEvent(event QTimerEvent_ITF) {
	defer qt.Recovering("QFinalState::timerEvent")

	if ptr.Pointer() != nil {
		C.QFinalState_TimerEvent(ptr.Pointer(), PointerFromQTimerEvent(event))
	}
}

func (ptr *QFinalState) TimerEventDefault(event QTimerEvent_ITF) {
	defer qt.Recovering("QFinalState::timerEvent")

	if ptr.Pointer() != nil {
		C.QFinalState_TimerEventDefault(ptr.Pointer(), PointerFromQTimerEvent(event))
	}
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
func callbackQFinalStateChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFinalState::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*QChildEvent))(NewQChildEventFromPointer(event))
	} else {
		NewQFinalStateFromPointer(ptr).ChildEventDefault(NewQChildEventFromPointer(event))
	}
}

func (ptr *QFinalState) ChildEvent(event QChildEvent_ITF) {
	defer qt.Recovering("QFinalState::childEvent")

	if ptr.Pointer() != nil {
		C.QFinalState_ChildEvent(ptr.Pointer(), PointerFromQChildEvent(event))
	}
}

func (ptr *QFinalState) ChildEventDefault(event QChildEvent_ITF) {
	defer qt.Recovering("QFinalState::childEvent")

	if ptr.Pointer() != nil {
		C.QFinalState_ChildEventDefault(ptr.Pointer(), PointerFromQChildEvent(event))
	}
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
func callbackQFinalStateCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFinalState::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
	} else {
		NewQFinalStateFromPointer(ptr).CustomEventDefault(NewQEventFromPointer(event))
	}
}

func (ptr *QFinalState) CustomEvent(event QEvent_ITF) {
	defer qt.Recovering("QFinalState::customEvent")

	if ptr.Pointer() != nil {
		C.QFinalState_CustomEvent(ptr.Pointer(), PointerFromQEvent(event))
	}
}

func (ptr *QFinalState) CustomEventDefault(event QEvent_ITF) {
	defer qt.Recovering("QFinalState::customEvent")

	if ptr.Pointer() != nil {
		C.QFinalState_CustomEventDefault(ptr.Pointer(), PointerFromQEvent(event))
	}
}
