package core

//#include "qstatemachine.h"
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
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QStateMachine_" + qt.RandomIdentifier())
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
	return NewQStateMachineFromPointer(C.QStateMachine_NewQStateMachine(PointerFromQObject(parent)))
}

func NewQStateMachine2(childMode QState__ChildMode, parent QObject_ITF) *QStateMachine {
	return NewQStateMachineFromPointer(C.QStateMachine_NewQStateMachine2(C.int(childMode), PointerFromQObject(parent)))
}

func (ptr *QStateMachine) AddDefaultAnimation(animation QAbstractAnimation_ITF) {
	if ptr.Pointer() != nil {
		C.QStateMachine_AddDefaultAnimation(ptr.Pointer(), PointerFromQAbstractAnimation(animation))
	}
}

func (ptr *QStateMachine) AddState(state QAbstractState_ITF) {
	if ptr.Pointer() != nil {
		C.QStateMachine_AddState(ptr.Pointer(), PointerFromQAbstractState(state))
	}
}

func (ptr *QStateMachine) ClearError() {
	if ptr.Pointer() != nil {
		C.QStateMachine_ClearError(ptr.Pointer())
	}
}

func (ptr *QStateMachine) CancelDelayedEvent(id int) bool {
	if ptr.Pointer() != nil {
		return C.QStateMachine_CancelDelayedEvent(ptr.Pointer(), C.int(id)) != 0
	}
	return false
}

func (ptr *QStateMachine) Error() QStateMachine__Error {
	if ptr.Pointer() != nil {
		return QStateMachine__Error(C.QStateMachine_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStateMachine) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QStateMachine_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QStateMachine) EventFilter(watched QObject_ITF, event QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QStateMachine_EventFilter(ptr.Pointer(), PointerFromQObject(watched), PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QStateMachine) GlobalRestorePolicy() QState__RestorePolicy {
	if ptr.Pointer() != nil {
		return QState__RestorePolicy(C.QStateMachine_GlobalRestorePolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStateMachine) IsAnimated() bool {
	if ptr.Pointer() != nil {
		return C.QStateMachine_IsAnimated(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStateMachine) IsRunning() bool {
	if ptr.Pointer() != nil {
		return C.QStateMachine_IsRunning(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStateMachine) PostDelayedEvent(event QEvent_ITF, delay int) int {
	if ptr.Pointer() != nil {
		return int(C.QStateMachine_PostDelayedEvent(ptr.Pointer(), PointerFromQEvent(event), C.int(delay)))
	}
	return 0
}

func (ptr *QStateMachine) PostEvent(event QEvent_ITF, priority QStateMachine__EventPriority) {
	if ptr.Pointer() != nil {
		C.QStateMachine_PostEvent(ptr.Pointer(), PointerFromQEvent(event), C.int(priority))
	}
}

func (ptr *QStateMachine) RemoveDefaultAnimation(animation QAbstractAnimation_ITF) {
	if ptr.Pointer() != nil {
		C.QStateMachine_RemoveDefaultAnimation(ptr.Pointer(), PointerFromQAbstractAnimation(animation))
	}
}

func (ptr *QStateMachine) RemoveState(state QAbstractState_ITF) {
	if ptr.Pointer() != nil {
		C.QStateMachine_RemoveState(ptr.Pointer(), PointerFromQAbstractState(state))
	}
}

func (ptr *QStateMachine) ConnectRunningChanged(f func(running bool)) {
	if ptr.Pointer() != nil {
		C.QStateMachine_ConnectRunningChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "runningChanged", f)
	}
}

func (ptr *QStateMachine) DisconnectRunningChanged() {
	if ptr.Pointer() != nil {
		C.QStateMachine_DisconnectRunningChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "runningChanged")
	}
}

//export callbackQStateMachineRunningChanged
func callbackQStateMachineRunningChanged(ptrName *C.char, running C.int) {
	qt.GetSignal(C.GoString(ptrName), "runningChanged").(func(bool))(int(running) != 0)
}

func (ptr *QStateMachine) SetAnimated(enabled bool) {
	if ptr.Pointer() != nil {
		C.QStateMachine_SetAnimated(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QStateMachine) SetGlobalRestorePolicy(restorePolicy QState__RestorePolicy) {
	if ptr.Pointer() != nil {
		C.QStateMachine_SetGlobalRestorePolicy(ptr.Pointer(), C.int(restorePolicy))
	}
}

func (ptr *QStateMachine) SetRunning(running bool) {
	if ptr.Pointer() != nil {
		C.QStateMachine_SetRunning(ptr.Pointer(), C.int(qt.GoBoolToInt(running)))
	}
}

func (ptr *QStateMachine) Start() {
	if ptr.Pointer() != nil {
		C.QStateMachine_Start(ptr.Pointer())
	}
}

func (ptr *QStateMachine) ConnectStarted(f func()) {
	if ptr.Pointer() != nil {
		C.QStateMachine_ConnectStarted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "started", f)
	}
}

func (ptr *QStateMachine) DisconnectStarted() {
	if ptr.Pointer() != nil {
		C.QStateMachine_DisconnectStarted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "started")
	}
}

//export callbackQStateMachineStarted
func callbackQStateMachineStarted(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "started").(func())()
}

func (ptr *QStateMachine) Stop() {
	if ptr.Pointer() != nil {
		C.QStateMachine_Stop(ptr.Pointer())
	}
}

func (ptr *QStateMachine) ConnectStopped(f func()) {
	if ptr.Pointer() != nil {
		C.QStateMachine_ConnectStopped(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stopped", f)
	}
}

func (ptr *QStateMachine) DisconnectStopped() {
	if ptr.Pointer() != nil {
		C.QStateMachine_DisconnectStopped(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stopped")
	}
}

//export callbackQStateMachineStopped
func callbackQStateMachineStopped(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "stopped").(func())()
}

func (ptr *QStateMachine) DestroyQStateMachine() {
	if ptr.Pointer() != nil {
		C.QStateMachine_DestroyQStateMachine(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
