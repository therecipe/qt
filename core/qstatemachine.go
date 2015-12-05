package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStateMachine::QStateMachine")
		}
	}()

	return NewQStateMachineFromPointer(C.QStateMachine_NewQStateMachine(PointerFromQObject(parent)))
}

func NewQStateMachine2(childMode QState__ChildMode, parent QObject_ITF) *QStateMachine {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStateMachine::QStateMachine")
		}
	}()

	return NewQStateMachineFromPointer(C.QStateMachine_NewQStateMachine2(C.int(childMode), PointerFromQObject(parent)))
}

func (ptr *QStateMachine) AddDefaultAnimation(animation QAbstractAnimation_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStateMachine::addDefaultAnimation")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStateMachine_AddDefaultAnimation(ptr.Pointer(), PointerFromQAbstractAnimation(animation))
	}
}

func (ptr *QStateMachine) AddState(state QAbstractState_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStateMachine::addState")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStateMachine_AddState(ptr.Pointer(), PointerFromQAbstractState(state))
	}
}

func (ptr *QStateMachine) ClearError() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStateMachine::clearError")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStateMachine_ClearError(ptr.Pointer())
	}
}

func (ptr *QStateMachine) CancelDelayedEvent(id int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStateMachine::cancelDelayedEvent")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QStateMachine_CancelDelayedEvent(ptr.Pointer(), C.int(id)) != 0
	}
	return false
}

func (ptr *QStateMachine) Error() QStateMachine__Error {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStateMachine::error")
		}
	}()

	if ptr.Pointer() != nil {
		return QStateMachine__Error(C.QStateMachine_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStateMachine) ErrorString() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStateMachine::errorString")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QStateMachine_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QStateMachine) EventFilter(watched QObject_ITF, event QEvent_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStateMachine::eventFilter")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QStateMachine_EventFilter(ptr.Pointer(), PointerFromQObject(watched), PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QStateMachine) GlobalRestorePolicy() QState__RestorePolicy {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStateMachine::globalRestorePolicy")
		}
	}()

	if ptr.Pointer() != nil {
		return QState__RestorePolicy(C.QStateMachine_GlobalRestorePolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStateMachine) IsAnimated() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStateMachine::isAnimated")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QStateMachine_IsAnimated(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStateMachine) IsRunning() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStateMachine::isRunning")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QStateMachine_IsRunning(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStateMachine) PostDelayedEvent(event QEvent_ITF, delay int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStateMachine::postDelayedEvent")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStateMachine_PostDelayedEvent(ptr.Pointer(), PointerFromQEvent(event), C.int(delay)))
	}
	return 0
}

func (ptr *QStateMachine) PostEvent(event QEvent_ITF, priority QStateMachine__EventPriority) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStateMachine::postEvent")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStateMachine_PostEvent(ptr.Pointer(), PointerFromQEvent(event), C.int(priority))
	}
}

func (ptr *QStateMachine) RemoveDefaultAnimation(animation QAbstractAnimation_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStateMachine::removeDefaultAnimation")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStateMachine_RemoveDefaultAnimation(ptr.Pointer(), PointerFromQAbstractAnimation(animation))
	}
}

func (ptr *QStateMachine) RemoveState(state QAbstractState_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStateMachine::removeState")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStateMachine_RemoveState(ptr.Pointer(), PointerFromQAbstractState(state))
	}
}

func (ptr *QStateMachine) ConnectRunningChanged(f func(running bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStateMachine::runningChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStateMachine_ConnectRunningChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "runningChanged", f)
	}
}

func (ptr *QStateMachine) DisconnectRunningChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStateMachine::runningChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStateMachine_DisconnectRunningChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "runningChanged")
	}
}

//export callbackQStateMachineRunningChanged
func callbackQStateMachineRunningChanged(ptrName *C.char, running C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStateMachine::runningChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "runningChanged").(func(bool))(int(running) != 0)
}

func (ptr *QStateMachine) SetAnimated(enabled bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStateMachine::setAnimated")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStateMachine_SetAnimated(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QStateMachine) SetGlobalRestorePolicy(restorePolicy QState__RestorePolicy) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStateMachine::setGlobalRestorePolicy")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStateMachine_SetGlobalRestorePolicy(ptr.Pointer(), C.int(restorePolicy))
	}
}

func (ptr *QStateMachine) SetRunning(running bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStateMachine::setRunning")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStateMachine_SetRunning(ptr.Pointer(), C.int(qt.GoBoolToInt(running)))
	}
}

func (ptr *QStateMachine) Start() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStateMachine::start")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStateMachine_Start(ptr.Pointer())
	}
}

func (ptr *QStateMachine) ConnectStarted(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStateMachine::started")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStateMachine_ConnectStarted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "started", f)
	}
}

func (ptr *QStateMachine) DisconnectStarted() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStateMachine::started")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStateMachine_DisconnectStarted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "started")
	}
}

//export callbackQStateMachineStarted
func callbackQStateMachineStarted(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStateMachine::started")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "started").(func())()
}

func (ptr *QStateMachine) Stop() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStateMachine::stop")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStateMachine_Stop(ptr.Pointer())
	}
}

func (ptr *QStateMachine) ConnectStopped(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStateMachine::stopped")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStateMachine_ConnectStopped(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stopped", f)
	}
}

func (ptr *QStateMachine) DisconnectStopped() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStateMachine::stopped")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStateMachine_DisconnectStopped(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stopped")
	}
}

//export callbackQStateMachineStopped
func callbackQStateMachineStopped(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStateMachine::stopped")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "stopped").(func())()
}

func (ptr *QStateMachine) DestroyQStateMachine() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStateMachine::~QStateMachine")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStateMachine_DestroyQStateMachine(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
