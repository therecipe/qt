package core

//#include "qthread.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QThread struct {
	QObject
}

type QThread_ITF interface {
	QObject_ITF
	QThread_PTR() *QThread
}

func PointerFromQThread(ptr QThread_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QThread_PTR().Pointer()
	}
	return nil
}

func NewQThreadFromPointer(ptr unsafe.Pointer) *QThread {
	var n = new(QThread)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QThread_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QThread) QThread_PTR() *QThread {
	return ptr
}

//QThread::Priority
type QThread__Priority int64

const (
	QThread__IdlePriority         = QThread__Priority(0)
	QThread__LowestPriority       = QThread__Priority(1)
	QThread__LowPriority          = QThread__Priority(2)
	QThread__NormalPriority       = QThread__Priority(3)
	QThread__HighPriority         = QThread__Priority(4)
	QThread__HighestPriority      = QThread__Priority(5)
	QThread__TimeCriticalPriority = QThread__Priority(6)
	QThread__InheritPriority      = QThread__Priority(7)
)

func (ptr *QThread) SetPriority(priority QThread__Priority) {
	if ptr.Pointer() != nil {
		C.QThread_SetPriority(ptr.Pointer(), C.int(priority))
	}
}

func NewQThread(parent QObject_ITF) *QThread {
	return NewQThreadFromPointer(C.QThread_NewQThread(PointerFromQObject(parent)))
}

func QThread_CurrentThread() *QThread {
	return NewQThreadFromPointer(C.QThread_QThread_CurrentThread())
}

func (ptr *QThread) Event(event QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QThread_Event(ptr.Pointer(), PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QThread) EventDispatcher() *QAbstractEventDispatcher {
	if ptr.Pointer() != nil {
		return NewQAbstractEventDispatcherFromPointer(C.QThread_EventDispatcher(ptr.Pointer()))
	}
	return nil
}

func (ptr *QThread) Exit(returnCode int) {
	if ptr.Pointer() != nil {
		C.QThread_Exit(ptr.Pointer(), C.int(returnCode))
	}
}

func (ptr *QThread) ConnectFinished(f func()) {
	if ptr.Pointer() != nil {
		C.QThread_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QThread) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QThread_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQThreadFinished
func callbackQThreadFinished(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "finished").(func())()
}

func (ptr *QThread) IsFinished() bool {
	if ptr.Pointer() != nil {
		return C.QThread_IsFinished(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QThread) IsInterruptionRequested() bool {
	if ptr.Pointer() != nil {
		return C.QThread_IsInterruptionRequested(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QThread) IsRunning() bool {
	if ptr.Pointer() != nil {
		return C.QThread_IsRunning(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QThread) LoopLevel() int {
	if ptr.Pointer() != nil {
		return int(C.QThread_LoopLevel(ptr.Pointer()))
	}
	return 0
}

func (ptr *QThread) Priority() QThread__Priority {
	if ptr.Pointer() != nil {
		return QThread__Priority(C.QThread_Priority(ptr.Pointer()))
	}
	return 0
}

func (ptr *QThread) Quit() {
	if ptr.Pointer() != nil {
		C.QThread_Quit(ptr.Pointer())
	}
}

func (ptr *QThread) RequestInterruption() {
	if ptr.Pointer() != nil {
		C.QThread_RequestInterruption(ptr.Pointer())
	}
}

func (ptr *QThread) SetEventDispatcher(eventDispatcher QAbstractEventDispatcher_ITF) {
	if ptr.Pointer() != nil {
		C.QThread_SetEventDispatcher(ptr.Pointer(), PointerFromQAbstractEventDispatcher(eventDispatcher))
	}
}

func (ptr *QThread) ConnectStarted(f func()) {
	if ptr.Pointer() != nil {
		C.QThread_ConnectStarted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "started", f)
	}
}

func (ptr *QThread) DisconnectStarted() {
	if ptr.Pointer() != nil {
		C.QThread_DisconnectStarted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "started")
	}
}

//export callbackQThreadStarted
func callbackQThreadStarted(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "started").(func())()
}

func (ptr *QThread) DestroyQThread() {
	if ptr.Pointer() != nil {
		C.QThread_DestroyQThread(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func QThread_IdealThreadCount() int {
	return int(C.QThread_QThread_IdealThreadCount())
}

func (ptr *QThread) Start(priority QThread__Priority) {
	if ptr.Pointer() != nil {
		C.QThread_Start(ptr.Pointer(), C.int(priority))
	}
}

func (ptr *QThread) Terminate() {
	if ptr.Pointer() != nil {
		C.QThread_Terminate(ptr.Pointer())
	}
}

func QThread_YieldCurrentThread() {
	C.QThread_QThread_YieldCurrentThread()
}
