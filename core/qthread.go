package core

//#include "core.h"
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
	for len(n.ObjectName()) < len("QThread_") {
		n.SetObjectName("QThread_" + qt.Identifier())
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
	defer qt.Recovering("QThread::setPriority")

	if ptr.Pointer() != nil {
		C.QThread_SetPriority(ptr.Pointer(), C.int(priority))
	}
}

func NewQThread(parent QObject_ITF) *QThread {
	defer qt.Recovering("QThread::QThread")

	return NewQThreadFromPointer(C.QThread_NewQThread(PointerFromQObject(parent)))
}

func QThread_CurrentThread() *QThread {
	defer qt.Recovering("QThread::currentThread")

	return NewQThreadFromPointer(C.QThread_QThread_CurrentThread())
}

func (ptr *QThread) Event(event QEvent_ITF) bool {
	defer qt.Recovering("QThread::event")

	if ptr.Pointer() != nil {
		return C.QThread_Event(ptr.Pointer(), PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QThread) EventDispatcher() *QAbstractEventDispatcher {
	defer qt.Recovering("QThread::eventDispatcher")

	if ptr.Pointer() != nil {
		return NewQAbstractEventDispatcherFromPointer(C.QThread_EventDispatcher(ptr.Pointer()))
	}
	return nil
}

func (ptr *QThread) Exit(returnCode int) {
	defer qt.Recovering("QThread::exit")

	if ptr.Pointer() != nil {
		C.QThread_Exit(ptr.Pointer(), C.int(returnCode))
	}
}

func (ptr *QThread) ConnectFinished(f func()) {
	defer qt.Recovering("connect QThread::finished")

	if ptr.Pointer() != nil {
		C.QThread_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QThread) DisconnectFinished() {
	defer qt.Recovering("disconnect QThread::finished")

	if ptr.Pointer() != nil {
		C.QThread_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQThreadFinished
func callbackQThreadFinished(ptrName *C.char) {
	defer qt.Recovering("callback QThread::finished")

	if signal := qt.GetSignal(C.GoString(ptrName), "finished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QThread) IsFinished() bool {
	defer qt.Recovering("QThread::isFinished")

	if ptr.Pointer() != nil {
		return C.QThread_IsFinished(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QThread) IsInterruptionRequested() bool {
	defer qt.Recovering("QThread::isInterruptionRequested")

	if ptr.Pointer() != nil {
		return C.QThread_IsInterruptionRequested(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QThread) IsRunning() bool {
	defer qt.Recovering("QThread::isRunning")

	if ptr.Pointer() != nil {
		return C.QThread_IsRunning(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QThread) LoopLevel() int {
	defer qt.Recovering("QThread::loopLevel")

	if ptr.Pointer() != nil {
		return int(C.QThread_LoopLevel(ptr.Pointer()))
	}
	return 0
}

func (ptr *QThread) Priority() QThread__Priority {
	defer qt.Recovering("QThread::priority")

	if ptr.Pointer() != nil {
		return QThread__Priority(C.QThread_Priority(ptr.Pointer()))
	}
	return 0
}

func (ptr *QThread) Quit() {
	defer qt.Recovering("QThread::quit")

	if ptr.Pointer() != nil {
		C.QThread_Quit(ptr.Pointer())
	}
}

func (ptr *QThread) RequestInterruption() {
	defer qt.Recovering("QThread::requestInterruption")

	if ptr.Pointer() != nil {
		C.QThread_RequestInterruption(ptr.Pointer())
	}
}

func (ptr *QThread) ConnectRun(f func()) {
	defer qt.Recovering("connect QThread::run")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "run", f)
	}
}

func (ptr *QThread) DisconnectRun() {
	defer qt.Recovering("disconnect QThread::run")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "run")
	}
}

//export callbackQThreadRun
func callbackQThreadRun(ptrName *C.char) bool {
	defer qt.Recovering("callback QThread::run")

	if signal := qt.GetSignal(C.GoString(ptrName), "run"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QThread) SetEventDispatcher(eventDispatcher QAbstractEventDispatcher_ITF) {
	defer qt.Recovering("QThread::setEventDispatcher")

	if ptr.Pointer() != nil {
		C.QThread_SetEventDispatcher(ptr.Pointer(), PointerFromQAbstractEventDispatcher(eventDispatcher))
	}
}

func (ptr *QThread) ConnectStarted(f func()) {
	defer qt.Recovering("connect QThread::started")

	if ptr.Pointer() != nil {
		C.QThread_ConnectStarted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "started", f)
	}
}

func (ptr *QThread) DisconnectStarted() {
	defer qt.Recovering("disconnect QThread::started")

	if ptr.Pointer() != nil {
		C.QThread_DisconnectStarted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "started")
	}
}

//export callbackQThreadStarted
func callbackQThreadStarted(ptrName *C.char) {
	defer qt.Recovering("callback QThread::started")

	if signal := qt.GetSignal(C.GoString(ptrName), "started"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QThread) DestroyQThread() {
	defer qt.Recovering("QThread::~QThread")

	if ptr.Pointer() != nil {
		C.QThread_DestroyQThread(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func QThread_IdealThreadCount() int {
	defer qt.Recovering("QThread::idealThreadCount")

	return int(C.QThread_QThread_IdealThreadCount())
}

func (ptr *QThread) Start(priority QThread__Priority) {
	defer qt.Recovering("QThread::start")

	if ptr.Pointer() != nil {
		C.QThread_Start(ptr.Pointer(), C.int(priority))
	}
}

func (ptr *QThread) Terminate() {
	defer qt.Recovering("QThread::terminate")

	if ptr.Pointer() != nil {
		C.QThread_Terminate(ptr.Pointer())
	}
}

func QThread_YieldCurrentThread() {
	defer qt.Recovering("QThread::yieldCurrentThread")

	C.QThread_QThread_YieldCurrentThread()
}

func (ptr *QThread) ConnectTimerEvent(f func(event *QTimerEvent)) {
	defer qt.Recovering("connect QThread::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QThread) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QThread::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQThreadTimerEvent
func callbackQThreadTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QThread::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*QTimerEvent))(NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QThread) ConnectChildEvent(f func(event *QChildEvent)) {
	defer qt.Recovering("connect QThread::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QThread) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QThread::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQThreadChildEvent
func callbackQThreadChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QThread::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*QChildEvent))(NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QThread) ConnectCustomEvent(f func(event *QEvent)) {
	defer qt.Recovering("connect QThread::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QThread) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QThread::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQThreadCustomEvent
func callbackQThreadCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QThread::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
		return true
	}
	return false

}
