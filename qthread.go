package qt

//#include "qthread.h"
import "C"

type qthread struct {
	qobject
}

type QThread interface {
	QObject
	Exit_Int(returnCode int)
	IsFinished() bool
	IsInterruptionRequested() bool
	IsRunning() bool
	Priority() Priority
	RequestInterruption()
	SetPriority_Priority(priority Priority)
	ConnectSlotQuit()
	DisconnectSlotQuit()
	SlotQuit()
	ConnectSlotTerminate()
	DisconnectSlotTerminate()
	SlotTerminate()
	ConnectSignalFinished(f func())
	DisconnectSignalFinished()
	SignalFinished() func()
	ConnectSignalStarted(f func())
	DisconnectSignalStarted()
	SignalStarted() func()
}

func (p *qthread) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qthread) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

//Priority
type Priority int

var (
	IDLEPRIORITY         = Priority(C.QThread_IdlePriority())
	LOWESTPRIORITY       = Priority(C.QThread_LowestPriority())
	LOWPRIORITY          = Priority(C.QThread_LowPriority())
	NORMALPRIORITY       = Priority(C.QThread_NormalPriority())
	HIGHPRIORITY         = Priority(C.QThread_HighPriority())
	HIGHESTPRIORITY      = Priority(C.QThread_HighestPriority())
	TIMECRITICALPRIORITY = Priority(C.QThread_TimeCriticalPriority())
	INHERITPRIORITY      = Priority(C.QThread_InheritPriority())
)

func NewQThread_QObject(parent QObject) QThread {
	var parentPtr C.QtObjectPtr = nil
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qthread = new(qthread)
	qthread.SetPointer(C.QThread_New_QObject(parentPtr))
	qthread.SetObjectName_String("QThread_" + randomIdentifier())
	return qthread
}

func (p *qthread) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QThread_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qthread) Exit_Int(returnCode int) {
	if p.Pointer() != nil {
		C.QThread_Exit_Int(p.Pointer(), C.int(returnCode))
	}
}

func (p *qthread) IsFinished() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QThread_IsFinished(p.Pointer()) != 0
	}
}

func (p *qthread) IsInterruptionRequested() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QThread_IsInterruptionRequested(p.Pointer()) != 0
	}
}

func (p *qthread) IsRunning() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QThread_IsRunning(p.Pointer()) != 0
	}
}

func (p *qthread) Priority() Priority {
	if p.Pointer() == nil {
		return 0
	} else {
		return Priority(C.QThread_Priority(p.Pointer()))
	}
}

func (p *qthread) RequestInterruption() {
	if p.Pointer() != nil {
		C.QThread_RequestInterruption(p.Pointer())
	}
}

func (p *qthread) SetPriority_Priority(priority Priority) {
	if p.Pointer() != nil {
		C.QThread_SetPriority_Priority(p.Pointer(), C.int(priority))
	}
}

func (p *qthread) ConnectSlotQuit() {
	C.QThread_ConnectSlotQuit(p.Pointer())
}

func (p *qthread) DisconnectSlotQuit() {
	C.QThread_DisconnectSlotQuit(p.Pointer())
}

func (p *qthread) SlotQuit() {
	if p.Pointer() != nil {
		C.QThread_Quit(p.Pointer())
	}
}

func (p *qthread) ConnectSlotTerminate() {
	C.QThread_ConnectSlotTerminate(p.Pointer())
}

func (p *qthread) DisconnectSlotTerminate() {
	C.QThread_DisconnectSlotTerminate(p.Pointer())
}

func (p *qthread) SlotTerminate() {
	if p.Pointer() != nil {
		C.QThread_Terminate(p.Pointer())
	}
}

func (p *qthread) ConnectSignalFinished(f func()) {
	C.QThread_ConnectSignalFinished(p.Pointer())
	connectSignal(p.ObjectName(), "finished", f)
}

func (p *qthread) DisconnectSignalFinished() {
	C.QThread_DisconnectSignalFinished(p.Pointer())
	disconnectSignal(p.ObjectName(), "finished")
}

func (p *qthread) SignalFinished() func() {
	return getSignal(p.ObjectName(), "finished")
}

func (p *qthread) ConnectSignalStarted(f func()) {
	C.QThread_ConnectSignalStarted(p.Pointer())
	connectSignal(p.ObjectName(), "started", f)
}

func (p *qthread) DisconnectSignalStarted() {
	C.QThread_DisconnectSignalStarted(p.Pointer())
	disconnectSignal(p.ObjectName(), "started")
}

func (p *qthread) SignalStarted() func() {
	return getSignal(p.ObjectName(), "started")
}

func QThread_CurrentThread() QThread {
	var qthread = new(qthread)
	qthread.SetPointer(C.QThread_CurrentThread())
	if qthread.ObjectName() == "" {
		qthread.SetObjectName_String("QThread_" + randomIdentifier())
	}
	return qthread
}

func QThread_IdealThreadCount() int {
	return int(C.QThread_IdealThreadCount())
}

//export callbackQThread
func callbackQThread(ptr C.QtObjectPtr, msg *C.char) {
	var qthread = new(qthread)
	qthread.SetPointer(ptr)
	getSignal(qthread.ObjectName(), C.GoString(msg))()
}
