package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QStateMachine struct {
	QState
}

type QStateMachine_ITF interface {
	QState_ITF
	QStateMachine_PTR() *QStateMachine
}

func PointerFromQStateMachine(ptr QStateMachine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStateMachine_PTR().Pointer()
	}
	return nil
}

func NewQStateMachineFromPointer(ptr unsafe.Pointer) *QStateMachine {
	var n = new(QStateMachine)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QStateMachine_") {
		n.SetObjectName("QStateMachine_" + qt.Identifier())
	}
	return n
}

func (ptr *QStateMachine) QStateMachine_PTR() *QStateMachine {
	return ptr
}

//QStateMachine::Error
type QStateMachine__Error int64

const (
	QStateMachine__NoError                            = QStateMachine__Error(0)
	QStateMachine__NoInitialStateError                = QStateMachine__Error(1)
	QStateMachine__NoDefaultStateInHistoryStateError  = QStateMachine__Error(2)
	QStateMachine__NoCommonAncestorForTransitionError = QStateMachine__Error(3)
)

//QStateMachine::EventPriority
type QStateMachine__EventPriority int64

const (
	QStateMachine__NormalPriority = QStateMachine__EventPriority(0)
	QStateMachine__HighPriority   = QStateMachine__EventPriority(1)
)

func NewQStateMachine(parent QObject_ITF) *QStateMachine {
	defer qt.Recovering("QStateMachine::QStateMachine")

	return NewQStateMachineFromPointer(C.QStateMachine_NewQStateMachine(PointerFromQObject(parent)))
}

func NewQStateMachine2(childMode QState__ChildMode, parent QObject_ITF) *QStateMachine {
	defer qt.Recovering("QStateMachine::QStateMachine")

	return NewQStateMachineFromPointer(C.QStateMachine_NewQStateMachine2(C.int(childMode), PointerFromQObject(parent)))
}

func (ptr *QStateMachine) AddDefaultAnimation(animation QAbstractAnimation_ITF) {
	defer qt.Recovering("QStateMachine::addDefaultAnimation")

	if ptr.Pointer() != nil {
		C.QStateMachine_AddDefaultAnimation(ptr.Pointer(), PointerFromQAbstractAnimation(animation))
	}
}

func (ptr *QStateMachine) AddState(state QAbstractState_ITF) {
	defer qt.Recovering("QStateMachine::addState")

	if ptr.Pointer() != nil {
		C.QStateMachine_AddState(ptr.Pointer(), PointerFromQAbstractState(state))
	}
}

func (ptr *QStateMachine) ClearError() {
	defer qt.Recovering("QStateMachine::clearError")

	if ptr.Pointer() != nil {
		C.QStateMachine_ClearError(ptr.Pointer())
	}
}

func (ptr *QStateMachine) CancelDelayedEvent(id int) bool {
	defer qt.Recovering("QStateMachine::cancelDelayedEvent")

	if ptr.Pointer() != nil {
		return C.QStateMachine_CancelDelayedEvent(ptr.Pointer(), C.int(id)) != 0
	}
	return false
}

func (ptr *QStateMachine) Error() QStateMachine__Error {
	defer qt.Recovering("QStateMachine::error")

	if ptr.Pointer() != nil {
		return QStateMachine__Error(C.QStateMachine_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStateMachine) ErrorString() string {
	defer qt.Recovering("QStateMachine::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QStateMachine_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QStateMachine) EventFilter(watched QObject_ITF, event QEvent_ITF) bool {
	defer qt.Recovering("QStateMachine::eventFilter")

	if ptr.Pointer() != nil {
		return C.QStateMachine_EventFilter(ptr.Pointer(), PointerFromQObject(watched), PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QStateMachine) GlobalRestorePolicy() QState__RestorePolicy {
	defer qt.Recovering("QStateMachine::globalRestorePolicy")

	if ptr.Pointer() != nil {
		return QState__RestorePolicy(C.QStateMachine_GlobalRestorePolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStateMachine) IsAnimated() bool {
	defer qt.Recovering("QStateMachine::isAnimated")

	if ptr.Pointer() != nil {
		return C.QStateMachine_IsAnimated(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStateMachine) IsRunning() bool {
	defer qt.Recovering("QStateMachine::isRunning")

	if ptr.Pointer() != nil {
		return C.QStateMachine_IsRunning(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStateMachine) ConnectOnEntry(f func(event *QEvent)) {
	defer qt.Recovering("connect QStateMachine::onEntry")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "onEntry", f)
	}
}

func (ptr *QStateMachine) DisconnectOnEntry() {
	defer qt.Recovering("disconnect QStateMachine::onEntry")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "onEntry")
	}
}

//export callbackQStateMachineOnEntry
func callbackQStateMachineOnEntry(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QStateMachine::onEntry")

	var signal = qt.GetSignal(C.GoString(ptrName), "onEntry")
	if signal != nil {
		defer signal.(func(*QEvent))(NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QStateMachine) ConnectOnExit(f func(event *QEvent)) {
	defer qt.Recovering("connect QStateMachine::onExit")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "onExit", f)
	}
}

func (ptr *QStateMachine) DisconnectOnExit() {
	defer qt.Recovering("disconnect QStateMachine::onExit")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "onExit")
	}
}

//export callbackQStateMachineOnExit
func callbackQStateMachineOnExit(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QStateMachine::onExit")

	var signal = qt.GetSignal(C.GoString(ptrName), "onExit")
	if signal != nil {
		defer signal.(func(*QEvent))(NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QStateMachine) PostDelayedEvent(event QEvent_ITF, delay int) int {
	defer qt.Recovering("QStateMachine::postDelayedEvent")

	if ptr.Pointer() != nil {
		return int(C.QStateMachine_PostDelayedEvent(ptr.Pointer(), PointerFromQEvent(event), C.int(delay)))
	}
	return 0
}

func (ptr *QStateMachine) PostEvent(event QEvent_ITF, priority QStateMachine__EventPriority) {
	defer qt.Recovering("QStateMachine::postEvent")

	if ptr.Pointer() != nil {
		C.QStateMachine_PostEvent(ptr.Pointer(), PointerFromQEvent(event), C.int(priority))
	}
}

func (ptr *QStateMachine) RemoveDefaultAnimation(animation QAbstractAnimation_ITF) {
	defer qt.Recovering("QStateMachine::removeDefaultAnimation")

	if ptr.Pointer() != nil {
		C.QStateMachine_RemoveDefaultAnimation(ptr.Pointer(), PointerFromQAbstractAnimation(animation))
	}
}

func (ptr *QStateMachine) RemoveState(state QAbstractState_ITF) {
	defer qt.Recovering("QStateMachine::removeState")

	if ptr.Pointer() != nil {
		C.QStateMachine_RemoveState(ptr.Pointer(), PointerFromQAbstractState(state))
	}
}

func (ptr *QStateMachine) ConnectRunningChanged(f func(running bool)) {
	defer qt.Recovering("connect QStateMachine::runningChanged")

	if ptr.Pointer() != nil {
		C.QStateMachine_ConnectRunningChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "runningChanged", f)
	}
}

func (ptr *QStateMachine) DisconnectRunningChanged() {
	defer qt.Recovering("disconnect QStateMachine::runningChanged")

	if ptr.Pointer() != nil {
		C.QStateMachine_DisconnectRunningChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "runningChanged")
	}
}

//export callbackQStateMachineRunningChanged
func callbackQStateMachineRunningChanged(ptrName *C.char, running C.int) {
	defer qt.Recovering("callback QStateMachine::runningChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "runningChanged")
	if signal != nil {
		signal.(func(bool))(int(running) != 0)
	}

}

func (ptr *QStateMachine) SetAnimated(enabled bool) {
	defer qt.Recovering("QStateMachine::setAnimated")

	if ptr.Pointer() != nil {
		C.QStateMachine_SetAnimated(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QStateMachine) SetGlobalRestorePolicy(restorePolicy QState__RestorePolicy) {
	defer qt.Recovering("QStateMachine::setGlobalRestorePolicy")

	if ptr.Pointer() != nil {
		C.QStateMachine_SetGlobalRestorePolicy(ptr.Pointer(), C.int(restorePolicy))
	}
}

func (ptr *QStateMachine) SetRunning(running bool) {
	defer qt.Recovering("QStateMachine::setRunning")

	if ptr.Pointer() != nil {
		C.QStateMachine_SetRunning(ptr.Pointer(), C.int(qt.GoBoolToInt(running)))
	}
}

func (ptr *QStateMachine) Start() {
	defer qt.Recovering("QStateMachine::start")

	if ptr.Pointer() != nil {
		C.QStateMachine_Start(ptr.Pointer())
	}
}

func (ptr *QStateMachine) ConnectStarted(f func()) {
	defer qt.Recovering("connect QStateMachine::started")

	if ptr.Pointer() != nil {
		C.QStateMachine_ConnectStarted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "started", f)
	}
}

func (ptr *QStateMachine) DisconnectStarted() {
	defer qt.Recovering("disconnect QStateMachine::started")

	if ptr.Pointer() != nil {
		C.QStateMachine_DisconnectStarted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "started")
	}
}

//export callbackQStateMachineStarted
func callbackQStateMachineStarted(ptrName *C.char) {
	defer qt.Recovering("callback QStateMachine::started")

	var signal = qt.GetSignal(C.GoString(ptrName), "started")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QStateMachine) Stop() {
	defer qt.Recovering("QStateMachine::stop")

	if ptr.Pointer() != nil {
		C.QStateMachine_Stop(ptr.Pointer())
	}
}

func (ptr *QStateMachine) ConnectStopped(f func()) {
	defer qt.Recovering("connect QStateMachine::stopped")

	if ptr.Pointer() != nil {
		C.QStateMachine_ConnectStopped(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stopped", f)
	}
}

func (ptr *QStateMachine) DisconnectStopped() {
	defer qt.Recovering("disconnect QStateMachine::stopped")

	if ptr.Pointer() != nil {
		C.QStateMachine_DisconnectStopped(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stopped")
	}
}

//export callbackQStateMachineStopped
func callbackQStateMachineStopped(ptrName *C.char) {
	defer qt.Recovering("callback QStateMachine::stopped")

	var signal = qt.GetSignal(C.GoString(ptrName), "stopped")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QStateMachine) DestroyQStateMachine() {
	defer qt.Recovering("QStateMachine::~QStateMachine")

	if ptr.Pointer() != nil {
		C.QStateMachine_DestroyQStateMachine(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
