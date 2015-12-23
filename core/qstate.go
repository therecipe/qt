package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QState struct {
	QAbstractState
}

type QState_ITF interface {
	QAbstractState_ITF
	QState_PTR() *QState
}

func PointerFromQState(ptr QState_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QState_PTR().Pointer()
	}
	return nil
}

func NewQStateFromPointer(ptr unsafe.Pointer) *QState {
	var n = new(QState)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QState_") {
		n.SetObjectName("QState_" + qt.Identifier())
	}
	return n
}

func (ptr *QState) QState_PTR() *QState {
	return ptr
}

//QState::ChildMode
type QState__ChildMode int64

const (
	QState__ExclusiveStates = QState__ChildMode(0)
	QState__ParallelStates  = QState__ChildMode(1)
)

//QState::RestorePolicy
type QState__RestorePolicy int64

const (
	QState__DontRestoreProperties = QState__RestorePolicy(0)
	QState__RestoreProperties     = QState__RestorePolicy(1)
)

func NewQState2(childMode QState__ChildMode, parent QState_ITF) *QState {
	defer qt.Recovering("QState::QState")

	return NewQStateFromPointer(C.QState_NewQState2(C.int(childMode), PointerFromQState(parent)))
}

func NewQState(parent QState_ITF) *QState {
	defer qt.Recovering("QState::QState")

	return NewQStateFromPointer(C.QState_NewQState(PointerFromQState(parent)))
}

func (ptr *QState) AddTransition3(target QAbstractState_ITF) *QAbstractTransition {
	defer qt.Recovering("QState::addTransition")

	if ptr.Pointer() != nil {
		return NewQAbstractTransitionFromPointer(C.QState_AddTransition3(ptr.Pointer(), PointerFromQAbstractState(target)))
	}
	return nil
}

func (ptr *QState) AddTransition2(sender QObject_ITF, signal string, target QAbstractState_ITF) *QSignalTransition {
	defer qt.Recovering("QState::addTransition")

	if ptr.Pointer() != nil {
		return NewQSignalTransitionFromPointer(C.QState_AddTransition2(ptr.Pointer(), PointerFromQObject(sender), C.CString(signal), PointerFromQAbstractState(target)))
	}
	return nil
}

func (ptr *QState) AddTransition(transition QAbstractTransition_ITF) {
	defer qt.Recovering("QState::addTransition")

	if ptr.Pointer() != nil {
		C.QState_AddTransition(ptr.Pointer(), PointerFromQAbstractTransition(transition))
	}
}

func (ptr *QState) AssignProperty(object QObject_ITF, name string, value QVariant_ITF) {
	defer qt.Recovering("QState::assignProperty")

	if ptr.Pointer() != nil {
		C.QState_AssignProperty(ptr.Pointer(), PointerFromQObject(object), C.CString(name), PointerFromQVariant(value))
	}
}

func (ptr *QState) ChildMode() QState__ChildMode {
	defer qt.Recovering("QState::childMode")

	if ptr.Pointer() != nil {
		return QState__ChildMode(C.QState_ChildMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QState) ConnectChildModeChanged(f func()) {
	defer qt.Recovering("connect QState::childModeChanged")

	if ptr.Pointer() != nil {
		C.QState_ConnectChildModeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "childModeChanged", f)
	}
}

func (ptr *QState) DisconnectChildModeChanged() {
	defer qt.Recovering("disconnect QState::childModeChanged")

	if ptr.Pointer() != nil {
		C.QState_DisconnectChildModeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "childModeChanged")
	}
}

//export callbackQStateChildModeChanged
func callbackQStateChildModeChanged(ptrName *C.char) {
	defer qt.Recovering("callback QState::childModeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "childModeChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QState) ErrorState() *QAbstractState {
	defer qt.Recovering("QState::errorState")

	if ptr.Pointer() != nil {
		return NewQAbstractStateFromPointer(C.QState_ErrorState(ptr.Pointer()))
	}
	return nil
}

func (ptr *QState) ConnectErrorStateChanged(f func()) {
	defer qt.Recovering("connect QState::errorStateChanged")

	if ptr.Pointer() != nil {
		C.QState_ConnectErrorStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "errorStateChanged", f)
	}
}

func (ptr *QState) DisconnectErrorStateChanged() {
	defer qt.Recovering("disconnect QState::errorStateChanged")

	if ptr.Pointer() != nil {
		C.QState_DisconnectErrorStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "errorStateChanged")
	}
}

//export callbackQStateErrorStateChanged
func callbackQStateErrorStateChanged(ptrName *C.char) {
	defer qt.Recovering("callback QState::errorStateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "errorStateChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QState) ConnectFinished(f func()) {
	defer qt.Recovering("connect QState::finished")

	if ptr.Pointer() != nil {
		C.QState_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QState) DisconnectFinished() {
	defer qt.Recovering("disconnect QState::finished")

	if ptr.Pointer() != nil {
		C.QState_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQStateFinished
func callbackQStateFinished(ptrName *C.char) {
	defer qt.Recovering("callback QState::finished")

	if signal := qt.GetSignal(C.GoString(ptrName), "finished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QState) InitialState() *QAbstractState {
	defer qt.Recovering("QState::initialState")

	if ptr.Pointer() != nil {
		return NewQAbstractStateFromPointer(C.QState_InitialState(ptr.Pointer()))
	}
	return nil
}

func (ptr *QState) ConnectInitialStateChanged(f func()) {
	defer qt.Recovering("connect QState::initialStateChanged")

	if ptr.Pointer() != nil {
		C.QState_ConnectInitialStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "initialStateChanged", f)
	}
}

func (ptr *QState) DisconnectInitialStateChanged() {
	defer qt.Recovering("disconnect QState::initialStateChanged")

	if ptr.Pointer() != nil {
		C.QState_DisconnectInitialStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "initialStateChanged")
	}
}

//export callbackQStateInitialStateChanged
func callbackQStateInitialStateChanged(ptrName *C.char) {
	defer qt.Recovering("callback QState::initialStateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "initialStateChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QState) ConnectOnEntry(f func(event *QEvent)) {
	defer qt.Recovering("connect QState::onEntry")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "onEntry", f)
	}
}

func (ptr *QState) DisconnectOnEntry() {
	defer qt.Recovering("disconnect QState::onEntry")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "onEntry")
	}
}

//export callbackQStateOnEntry
func callbackQStateOnEntry(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QState::onEntry")

	if signal := qt.GetSignal(C.GoString(ptrName), "onEntry"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QState) ConnectOnExit(f func(event *QEvent)) {
	defer qt.Recovering("connect QState::onExit")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "onExit", f)
	}
}

func (ptr *QState) DisconnectOnExit() {
	defer qt.Recovering("disconnect QState::onExit")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "onExit")
	}
}

//export callbackQStateOnExit
func callbackQStateOnExit(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QState::onExit")

	if signal := qt.GetSignal(C.GoString(ptrName), "onExit"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QState) ConnectPropertiesAssigned(f func()) {
	defer qt.Recovering("connect QState::propertiesAssigned")

	if ptr.Pointer() != nil {
		C.QState_ConnectPropertiesAssigned(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "propertiesAssigned", f)
	}
}

func (ptr *QState) DisconnectPropertiesAssigned() {
	defer qt.Recovering("disconnect QState::propertiesAssigned")

	if ptr.Pointer() != nil {
		C.QState_DisconnectPropertiesAssigned(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "propertiesAssigned")
	}
}

//export callbackQStatePropertiesAssigned
func callbackQStatePropertiesAssigned(ptrName *C.char) {
	defer qt.Recovering("callback QState::propertiesAssigned")

	if signal := qt.GetSignal(C.GoString(ptrName), "propertiesAssigned"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QState) RemoveTransition(transition QAbstractTransition_ITF) {
	defer qt.Recovering("QState::removeTransition")

	if ptr.Pointer() != nil {
		C.QState_RemoveTransition(ptr.Pointer(), PointerFromQAbstractTransition(transition))
	}
}

func (ptr *QState) SetChildMode(mode QState__ChildMode) {
	defer qt.Recovering("QState::setChildMode")

	if ptr.Pointer() != nil {
		C.QState_SetChildMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QState) SetErrorState(state QAbstractState_ITF) {
	defer qt.Recovering("QState::setErrorState")

	if ptr.Pointer() != nil {
		C.QState_SetErrorState(ptr.Pointer(), PointerFromQAbstractState(state))
	}
}

func (ptr *QState) SetInitialState(state QAbstractState_ITF) {
	defer qt.Recovering("QState::setInitialState")

	if ptr.Pointer() != nil {
		C.QState_SetInitialState(ptr.Pointer(), PointerFromQAbstractState(state))
	}
}

func (ptr *QState) DestroyQState() {
	defer qt.Recovering("QState::~QState")

	if ptr.Pointer() != nil {
		C.QState_DestroyQState(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QState) ConnectTimerEvent(f func(event *QTimerEvent)) {
	defer qt.Recovering("connect QState::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QState) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QState::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQStateTimerEvent
func callbackQStateTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QState::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*QTimerEvent))(NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QState) ConnectChildEvent(f func(event *QChildEvent)) {
	defer qt.Recovering("connect QState::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QState) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QState::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQStateChildEvent
func callbackQStateChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QState::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*QChildEvent))(NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QState) ConnectCustomEvent(f func(event *QEvent)) {
	defer qt.Recovering("connect QState::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QState) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QState::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQStateCustomEvent
func callbackQStateCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QState::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
		return true
	}
	return false

}
