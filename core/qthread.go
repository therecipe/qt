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

type QThreadITF interface {
	QObjectITF
	QThreadPTR() *QThread
}

func PointerFromQThread(ptr QThreadITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QThreadPTR().Pointer()
	}
	return nil
}

func QThreadFromPointer(ptr unsafe.Pointer) *QThread {
	var n = new(QThread)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QThread_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QThread) QThreadPTR() *QThread {
	return ptr
}

//QThread::Priority
type QThread__Priority int

var (
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
		C.QThread_SetPriority(C.QtObjectPtr(ptr.Pointer()), C.int(priority))
	}
}

func NewQThread(parent QObjectITF) *QThread {
	return QThreadFromPointer(unsafe.Pointer(C.QThread_NewQThread(C.QtObjectPtr(PointerFromQObject(parent)))))
}

func QThread_CurrentThread() *QThread {
	return QThreadFromPointer(unsafe.Pointer(C.QThread_QThread_CurrentThread()))
}

func (ptr *QThread) Event(event QEventITF) bool {
	if ptr.Pointer() != nil {
		return C.QThread_Event(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQEvent(event))) != 0
	}
	return false
}

func (ptr *QThread) EventDispatcher() *QAbstractEventDispatcher {
	if ptr.Pointer() != nil {
		return QAbstractEventDispatcherFromPointer(unsafe.Pointer(C.QThread_EventDispatcher(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QThread) Exit(returnCode int) {
	if ptr.Pointer() != nil {
		C.QThread_Exit(C.QtObjectPtr(ptr.Pointer()), C.int(returnCode))
	}
}

func (ptr *QThread) ConnectFinished(f func()) {
	if ptr.Pointer() != nil {
		C.QThread_ConnectFinished(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QThread) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QThread_DisconnectFinished(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQThreadFinished
func callbackQThreadFinished(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "finished").(func())()
}

func (ptr *QThread) IsFinished() bool {
	if ptr.Pointer() != nil {
		return C.QThread_IsFinished(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QThread) IsInterruptionRequested() bool {
	if ptr.Pointer() != nil {
		return C.QThread_IsInterruptionRequested(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QThread) IsRunning() bool {
	if ptr.Pointer() != nil {
		return C.QThread_IsRunning(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QThread) LoopLevel() int {
	if ptr.Pointer() != nil {
		return int(C.QThread_LoopLevel(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QThread) Priority() QThread__Priority {
	if ptr.Pointer() != nil {
		return QThread__Priority(C.QThread_Priority(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QThread) Quit() {
	if ptr.Pointer() != nil {
		C.QThread_Quit(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QThread) RequestInterruption() {
	if ptr.Pointer() != nil {
		C.QThread_RequestInterruption(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QThread) SetEventDispatcher(eventDispatcher QAbstractEventDispatcherITF) {
	if ptr.Pointer() != nil {
		C.QThread_SetEventDispatcher(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAbstractEventDispatcher(eventDispatcher)))
	}
}

func (ptr *QThread) ConnectStarted(f func()) {
	if ptr.Pointer() != nil {
		C.QThread_ConnectStarted(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "started", f)
	}
}

func (ptr *QThread) DisconnectStarted() {
	if ptr.Pointer() != nil {
		C.QThread_DisconnectStarted(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "started")
	}
}

//export callbackQThreadStarted
func callbackQThreadStarted(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "started").(func())()
}

func (ptr *QThread) DestroyQThread() {
	if ptr.Pointer() != nil {
		C.QThread_DestroyQThread(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func QThread_IdealThreadCount() int {
	return int(C.QThread_QThread_IdealThreadCount())
}

func (ptr *QThread) Start(priority QThread__Priority) {
	if ptr.Pointer() != nil {
		C.QThread_Start(C.QtObjectPtr(ptr.Pointer()), C.int(priority))
	}
}

func (ptr *QThread) Terminate() {
	if ptr.Pointer() != nil {
		C.QThread_Terminate(C.QtObjectPtr(ptr.Pointer()))
	}
}

func QThread_YieldCurrentThread() {
	C.QThread_QThread_YieldCurrentThread()
}
