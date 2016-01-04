package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QHistoryState struct {
	QAbstractState
}

type QHistoryState_ITF interface {
	QAbstractState_ITF
	QHistoryState_PTR() *QHistoryState
}

func PointerFromQHistoryState(ptr QHistoryState_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHistoryState_PTR().Pointer()
	}
	return nil
}

func NewQHistoryStateFromPointer(ptr unsafe.Pointer) *QHistoryState {
	var n = new(QHistoryState)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QHistoryState_") {
		n.SetObjectName("QHistoryState_" + qt.Identifier())
	}
	return n
}

func (ptr *QHistoryState) QHistoryState_PTR() *QHistoryState {
	return ptr
}

//QHistoryState::HistoryType
type QHistoryState__HistoryType int64

const (
	QHistoryState__ShallowHistory = QHistoryState__HistoryType(0)
	QHistoryState__DeepHistory    = QHistoryState__HistoryType(1)
)

func NewQHistoryState2(ty QHistoryState__HistoryType, parent QState_ITF) *QHistoryState {
	defer qt.Recovering("QHistoryState::QHistoryState")

	return NewQHistoryStateFromPointer(C.QHistoryState_NewQHistoryState2(C.int(ty), PointerFromQState(parent)))
}

func NewQHistoryState(parent QState_ITF) *QHistoryState {
	defer qt.Recovering("QHistoryState::QHistoryState")

	return NewQHistoryStateFromPointer(C.QHistoryState_NewQHistoryState(PointerFromQState(parent)))
}

func (ptr *QHistoryState) DefaultState() *QAbstractState {
	defer qt.Recovering("QHistoryState::defaultState")

	if ptr.Pointer() != nil {
		return NewQAbstractStateFromPointer(C.QHistoryState_DefaultState(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHistoryState) ConnectDefaultStateChanged(f func()) {
	defer qt.Recovering("connect QHistoryState::defaultStateChanged")

	if ptr.Pointer() != nil {
		C.QHistoryState_ConnectDefaultStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "defaultStateChanged", f)
	}
}

func (ptr *QHistoryState) DisconnectDefaultStateChanged() {
	defer qt.Recovering("disconnect QHistoryState::defaultStateChanged")

	if ptr.Pointer() != nil {
		C.QHistoryState_DisconnectDefaultStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "defaultStateChanged")
	}
}

//export callbackQHistoryStateDefaultStateChanged
func callbackQHistoryStateDefaultStateChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QHistoryState::defaultStateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "defaultStateChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHistoryState) Event(e QEvent_ITF) bool {
	defer qt.Recovering("QHistoryState::event")

	if ptr.Pointer() != nil {
		return C.QHistoryState_Event(ptr.Pointer(), PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QHistoryState) HistoryType() QHistoryState__HistoryType {
	defer qt.Recovering("QHistoryState::historyType")

	if ptr.Pointer() != nil {
		return QHistoryState__HistoryType(C.QHistoryState_HistoryType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHistoryState) ConnectHistoryTypeChanged(f func()) {
	defer qt.Recovering("connect QHistoryState::historyTypeChanged")

	if ptr.Pointer() != nil {
		C.QHistoryState_ConnectHistoryTypeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "historyTypeChanged", f)
	}
}

func (ptr *QHistoryState) DisconnectHistoryTypeChanged() {
	defer qt.Recovering("disconnect QHistoryState::historyTypeChanged")

	if ptr.Pointer() != nil {
		C.QHistoryState_DisconnectHistoryTypeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "historyTypeChanged")
	}
}

//export callbackQHistoryStateHistoryTypeChanged
func callbackQHistoryStateHistoryTypeChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QHistoryState::historyTypeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "historyTypeChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHistoryState) ConnectOnEntry(f func(event *QEvent)) {
	defer qt.Recovering("connect QHistoryState::onEntry")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "onEntry", f)
	}
}

func (ptr *QHistoryState) DisconnectOnEntry() {
	defer qt.Recovering("disconnect QHistoryState::onEntry")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "onEntry")
	}
}

//export callbackQHistoryStateOnEntry
func callbackQHistoryStateOnEntry(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHistoryState::onEntry")

	if signal := qt.GetSignal(C.GoString(ptrName), "onEntry"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
	}

}

func (ptr *QHistoryState) OnEntry(event QEvent_ITF) {
	defer qt.Recovering("QHistoryState::onEntry")

	if ptr.Pointer() != nil {
		C.QHistoryState_OnEntry(ptr.Pointer(), PointerFromQEvent(event))
	}
}

func (ptr *QHistoryState) OnEntryDefault(event QEvent_ITF) {
	defer qt.Recovering("QHistoryState::onEntry")

	if ptr.Pointer() != nil {
		C.QHistoryState_OnEntryDefault(ptr.Pointer(), PointerFromQEvent(event))
	}
}

func (ptr *QHistoryState) ConnectOnExit(f func(event *QEvent)) {
	defer qt.Recovering("connect QHistoryState::onExit")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "onExit", f)
	}
}

func (ptr *QHistoryState) DisconnectOnExit() {
	defer qt.Recovering("disconnect QHistoryState::onExit")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "onExit")
	}
}

//export callbackQHistoryStateOnExit
func callbackQHistoryStateOnExit(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHistoryState::onExit")

	if signal := qt.GetSignal(C.GoString(ptrName), "onExit"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
	}

}

func (ptr *QHistoryState) OnExit(event QEvent_ITF) {
	defer qt.Recovering("QHistoryState::onExit")

	if ptr.Pointer() != nil {
		C.QHistoryState_OnExit(ptr.Pointer(), PointerFromQEvent(event))
	}
}

func (ptr *QHistoryState) OnExitDefault(event QEvent_ITF) {
	defer qt.Recovering("QHistoryState::onExit")

	if ptr.Pointer() != nil {
		C.QHistoryState_OnExitDefault(ptr.Pointer(), PointerFromQEvent(event))
	}
}

func (ptr *QHistoryState) SetDefaultState(state QAbstractState_ITF) {
	defer qt.Recovering("QHistoryState::setDefaultState")

	if ptr.Pointer() != nil {
		C.QHistoryState_SetDefaultState(ptr.Pointer(), PointerFromQAbstractState(state))
	}
}

func (ptr *QHistoryState) SetHistoryType(ty QHistoryState__HistoryType) {
	defer qt.Recovering("QHistoryState::setHistoryType")

	if ptr.Pointer() != nil {
		C.QHistoryState_SetHistoryType(ptr.Pointer(), C.int(ty))
	}
}

func (ptr *QHistoryState) DestroyQHistoryState() {
	defer qt.Recovering("QHistoryState::~QHistoryState")

	if ptr.Pointer() != nil {
		C.QHistoryState_DestroyQHistoryState(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QHistoryState) ConnectTimerEvent(f func(event *QTimerEvent)) {
	defer qt.Recovering("connect QHistoryState::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QHistoryState) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QHistoryState::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQHistoryStateTimerEvent
func callbackQHistoryStateTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHistoryState::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*QTimerEvent))(NewQTimerEventFromPointer(event))
	} else {
		NewQHistoryStateFromPointer(ptr).TimerEventDefault(NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHistoryState) TimerEvent(event QTimerEvent_ITF) {
	defer qt.Recovering("QHistoryState::timerEvent")

	if ptr.Pointer() != nil {
		C.QHistoryState_TimerEvent(ptr.Pointer(), PointerFromQTimerEvent(event))
	}
}

func (ptr *QHistoryState) TimerEventDefault(event QTimerEvent_ITF) {
	defer qt.Recovering("QHistoryState::timerEvent")

	if ptr.Pointer() != nil {
		C.QHistoryState_TimerEventDefault(ptr.Pointer(), PointerFromQTimerEvent(event))
	}
}

func (ptr *QHistoryState) ConnectChildEvent(f func(event *QChildEvent)) {
	defer qt.Recovering("connect QHistoryState::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QHistoryState) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QHistoryState::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQHistoryStateChildEvent
func callbackQHistoryStateChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHistoryState::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*QChildEvent))(NewQChildEventFromPointer(event))
	} else {
		NewQHistoryStateFromPointer(ptr).ChildEventDefault(NewQChildEventFromPointer(event))
	}
}

func (ptr *QHistoryState) ChildEvent(event QChildEvent_ITF) {
	defer qt.Recovering("QHistoryState::childEvent")

	if ptr.Pointer() != nil {
		C.QHistoryState_ChildEvent(ptr.Pointer(), PointerFromQChildEvent(event))
	}
}

func (ptr *QHistoryState) ChildEventDefault(event QChildEvent_ITF) {
	defer qt.Recovering("QHistoryState::childEvent")

	if ptr.Pointer() != nil {
		C.QHistoryState_ChildEventDefault(ptr.Pointer(), PointerFromQChildEvent(event))
	}
}

func (ptr *QHistoryState) ConnectCustomEvent(f func(event *QEvent)) {
	defer qt.Recovering("connect QHistoryState::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QHistoryState) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QHistoryState::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQHistoryStateCustomEvent
func callbackQHistoryStateCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHistoryState::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
	} else {
		NewQHistoryStateFromPointer(ptr).CustomEventDefault(NewQEventFromPointer(event))
	}
}

func (ptr *QHistoryState) CustomEvent(event QEvent_ITF) {
	defer qt.Recovering("QHistoryState::customEvent")

	if ptr.Pointer() != nil {
		C.QHistoryState_CustomEvent(ptr.Pointer(), PointerFromQEvent(event))
	}
}

func (ptr *QHistoryState) CustomEventDefault(event QEvent_ITF) {
	defer qt.Recovering("QHistoryState::customEvent")

	if ptr.Pointer() != nil {
		C.QHistoryState_CustomEventDefault(ptr.Pointer(), PointerFromQEvent(event))
	}
}
