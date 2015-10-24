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

type QStateMachineITF interface {
	QStateITF
	QStateMachinePTR() *QStateMachine
}

func PointerFromQStateMachine(ptr QStateMachineITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStateMachinePTR().Pointer()
	}
	return nil
}

func QStateMachineFromPointer(ptr unsafe.Pointer) *QStateMachine {
	var n = new(QStateMachine)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QStateMachine_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QStateMachine) QStateMachinePTR() *QStateMachine {
	return ptr
}

//QStateMachine::Error
type QStateMachine__Error int

var (
	QStateMachine__NoError                            = QStateMachine__Error(0)
	QStateMachine__NoInitialStateError                = QStateMachine__Error(1)
	QStateMachine__NoDefaultStateInHistoryStateError  = QStateMachine__Error(2)
	QStateMachine__NoCommonAncestorForTransitionError = QStateMachine__Error(3)
)

//QStateMachine::EventPriority
type QStateMachine__EventPriority int

var (
	QStateMachine__NormalPriority = QStateMachine__EventPriority(0)
	QStateMachine__HighPriority   = QStateMachine__EventPriority(1)
)

func NewQStateMachine(parent QObjectITF) *QStateMachine {
	return QStateMachineFromPointer(unsafe.Pointer(C.QStateMachine_NewQStateMachine(C.QtObjectPtr(PointerFromQObject(parent)))))
}

func NewQStateMachine2(childMode QState__ChildMode, parent QObjectITF) *QStateMachine {
	return QStateMachineFromPointer(unsafe.Pointer(C.QStateMachine_NewQStateMachine2(C.int(childMode), C.QtObjectPtr(PointerFromQObject(parent)))))
}

func (ptr *QStateMachine) AddDefaultAnimation(animation QAbstractAnimationITF) {
	if ptr.Pointer() != nil {
		C.QStateMachine_AddDefaultAnimation(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAbstractAnimation(animation)))
	}
}

func (ptr *QStateMachine) AddState(state QAbstractStateITF) {
	if ptr.Pointer() != nil {
		C.QStateMachine_AddState(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAbstractState(state)))
	}
}

func (ptr *QStateMachine) ClearError() {
	if ptr.Pointer() != nil {
		C.QStateMachine_ClearError(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QStateMachine) CancelDelayedEvent(id int) bool {
	if ptr.Pointer() != nil {
		return C.QStateMachine_CancelDelayedEvent(C.QtObjectPtr(ptr.Pointer()), C.int(id)) != 0
	}
	return false
}

func (ptr *QStateMachine) Error() QStateMachine__Error {
	if ptr.Pointer() != nil {
		return QStateMachine__Error(C.QStateMachine_Error(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QStateMachine) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QStateMachine_ErrorString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QStateMachine) EventFilter(watched QObjectITF, event QEventITF) bool {
	if ptr.Pointer() != nil {
		return C.QStateMachine_EventFilter(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQObject(watched)), C.QtObjectPtr(PointerFromQEvent(event))) != 0
	}
	return false
}

func (ptr *QStateMachine) GlobalRestorePolicy() QState__RestorePolicy {
	if ptr.Pointer() != nil {
		return QState__RestorePolicy(C.QStateMachine_GlobalRestorePolicy(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QStateMachine) IsAnimated() bool {
	if ptr.Pointer() != nil {
		return C.QStateMachine_IsAnimated(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QStateMachine) IsRunning() bool {
	if ptr.Pointer() != nil {
		return C.QStateMachine_IsRunning(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QStateMachine) PostDelayedEvent(event QEventITF, delay int) int {
	if ptr.Pointer() != nil {
		return int(C.QStateMachine_PostDelayedEvent(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQEvent(event)), C.int(delay)))
	}
	return 0
}

func (ptr *QStateMachine) PostEvent(event QEventITF, priority QStateMachine__EventPriority) {
	if ptr.Pointer() != nil {
		C.QStateMachine_PostEvent(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQEvent(event)), C.int(priority))
	}
}

func (ptr *QStateMachine) RemoveDefaultAnimation(animation QAbstractAnimationITF) {
	if ptr.Pointer() != nil {
		C.QStateMachine_RemoveDefaultAnimation(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAbstractAnimation(animation)))
	}
}

func (ptr *QStateMachine) RemoveState(state QAbstractStateITF) {
	if ptr.Pointer() != nil {
		C.QStateMachine_RemoveState(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAbstractState(state)))
	}
}

func (ptr *QStateMachine) ConnectRunningChanged(f func(running bool)) {
	if ptr.Pointer() != nil {
		C.QStateMachine_ConnectRunningChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "runningChanged", f)
	}
}

func (ptr *QStateMachine) DisconnectRunningChanged() {
	if ptr.Pointer() != nil {
		C.QStateMachine_DisconnectRunningChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "runningChanged")
	}
}

//export callbackQStateMachineRunningChanged
func callbackQStateMachineRunningChanged(ptrName *C.char, running C.int) {
	qt.GetSignal(C.GoString(ptrName), "runningChanged").(func(bool))(int(running) != 0)
}

func (ptr *QStateMachine) SetAnimated(enabled bool) {
	if ptr.Pointer() != nil {
		C.QStateMachine_SetAnimated(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QStateMachine) SetGlobalRestorePolicy(restorePolicy QState__RestorePolicy) {
	if ptr.Pointer() != nil {
		C.QStateMachine_SetGlobalRestorePolicy(C.QtObjectPtr(ptr.Pointer()), C.int(restorePolicy))
	}
}

func (ptr *QStateMachine) SetRunning(running bool) {
	if ptr.Pointer() != nil {
		C.QStateMachine_SetRunning(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(running)))
	}
}

func (ptr *QStateMachine) Start() {
	if ptr.Pointer() != nil {
		C.QStateMachine_Start(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QStateMachine) ConnectStarted(f func()) {
	if ptr.Pointer() != nil {
		C.QStateMachine_ConnectStarted(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "started", f)
	}
}

func (ptr *QStateMachine) DisconnectStarted() {
	if ptr.Pointer() != nil {
		C.QStateMachine_DisconnectStarted(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "started")
	}
}

//export callbackQStateMachineStarted
func callbackQStateMachineStarted(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "started").(func())()
}

func (ptr *QStateMachine) Stop() {
	if ptr.Pointer() != nil {
		C.QStateMachine_Stop(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QStateMachine) ConnectStopped(f func()) {
	if ptr.Pointer() != nil {
		C.QStateMachine_ConnectStopped(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "stopped", f)
	}
}

func (ptr *QStateMachine) DisconnectStopped() {
	if ptr.Pointer() != nil {
		C.QStateMachine_DisconnectStopped(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "stopped")
	}
}

//export callbackQStateMachineStopped
func callbackQStateMachineStopped(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "stopped").(func())()
}

func (ptr *QStateMachine) DestroyQStateMachine() {
	if ptr.Pointer() != nil {
		C.QStateMachine_DestroyQStateMachine(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
