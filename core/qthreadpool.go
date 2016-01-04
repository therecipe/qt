package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QThreadPool struct {
	QObject
}

type QThreadPool_ITF interface {
	QObject_ITF
	QThreadPool_PTR() *QThreadPool
}

func PointerFromQThreadPool(ptr QThreadPool_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QThreadPool_PTR().Pointer()
	}
	return nil
}

func NewQThreadPoolFromPointer(ptr unsafe.Pointer) *QThreadPool {
	var n = new(QThreadPool)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QThreadPool_") {
		n.SetObjectName("QThreadPool_" + qt.Identifier())
	}
	return n
}

func (ptr *QThreadPool) QThreadPool_PTR() *QThreadPool {
	return ptr
}

func (ptr *QThreadPool) ActiveThreadCount() int {
	defer qt.Recovering("QThreadPool::activeThreadCount")

	if ptr.Pointer() != nil {
		return int(C.QThreadPool_ActiveThreadCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QThreadPool) ExpiryTimeout() int {
	defer qt.Recovering("QThreadPool::expiryTimeout")

	if ptr.Pointer() != nil {
		return int(C.QThreadPool_ExpiryTimeout(ptr.Pointer()))
	}
	return 0
}

func (ptr *QThreadPool) MaxThreadCount() int {
	defer qt.Recovering("QThreadPool::maxThreadCount")

	if ptr.Pointer() != nil {
		return int(C.QThreadPool_MaxThreadCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QThreadPool) SetExpiryTimeout(expiryTimeout int) {
	defer qt.Recovering("QThreadPool::setExpiryTimeout")

	if ptr.Pointer() != nil {
		C.QThreadPool_SetExpiryTimeout(ptr.Pointer(), C.int(expiryTimeout))
	}
}

func (ptr *QThreadPool) SetMaxThreadCount(maxThreadCount int) {
	defer qt.Recovering("QThreadPool::setMaxThreadCount")

	if ptr.Pointer() != nil {
		C.QThreadPool_SetMaxThreadCount(ptr.Pointer(), C.int(maxThreadCount))
	}
}

func NewQThreadPool(parent QObject_ITF) *QThreadPool {
	defer qt.Recovering("QThreadPool::QThreadPool")

	return NewQThreadPoolFromPointer(C.QThreadPool_NewQThreadPool(PointerFromQObject(parent)))
}

func (ptr *QThreadPool) Cancel(runnable QRunnable_ITF) {
	defer qt.Recovering("QThreadPool::cancel")

	if ptr.Pointer() != nil {
		C.QThreadPool_Cancel(ptr.Pointer(), PointerFromQRunnable(runnable))
	}
}

func (ptr *QThreadPool) Clear() {
	defer qt.Recovering("QThreadPool::clear")

	if ptr.Pointer() != nil {
		C.QThreadPool_Clear(ptr.Pointer())
	}
}

func QThreadPool_GlobalInstance() *QThreadPool {
	defer qt.Recovering("QThreadPool::globalInstance")

	return NewQThreadPoolFromPointer(C.QThreadPool_QThreadPool_GlobalInstance())
}

func (ptr *QThreadPool) ReleaseThread() {
	defer qt.Recovering("QThreadPool::releaseThread")

	if ptr.Pointer() != nil {
		C.QThreadPool_ReleaseThread(ptr.Pointer())
	}
}

func (ptr *QThreadPool) ReserveThread() {
	defer qt.Recovering("QThreadPool::reserveThread")

	if ptr.Pointer() != nil {
		C.QThreadPool_ReserveThread(ptr.Pointer())
	}
}

func (ptr *QThreadPool) Start(runnable QRunnable_ITF, priority int) {
	defer qt.Recovering("QThreadPool::start")

	if ptr.Pointer() != nil {
		C.QThreadPool_Start(ptr.Pointer(), PointerFromQRunnable(runnable), C.int(priority))
	}
}

func (ptr *QThreadPool) TryStart(runnable QRunnable_ITF) bool {
	defer qt.Recovering("QThreadPool::tryStart")

	if ptr.Pointer() != nil {
		return C.QThreadPool_TryStart(ptr.Pointer(), PointerFromQRunnable(runnable)) != 0
	}
	return false
}

func (ptr *QThreadPool) WaitForDone(msecs int) bool {
	defer qt.Recovering("QThreadPool::waitForDone")

	if ptr.Pointer() != nil {
		return C.QThreadPool_WaitForDone(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QThreadPool) DestroyQThreadPool() {
	defer qt.Recovering("QThreadPool::~QThreadPool")

	if ptr.Pointer() != nil {
		C.QThreadPool_DestroyQThreadPool(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QThreadPool) ConnectTimerEvent(f func(event *QTimerEvent)) {
	defer qt.Recovering("connect QThreadPool::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QThreadPool) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QThreadPool::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQThreadPoolTimerEvent
func callbackQThreadPoolTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QThreadPool::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*QTimerEvent))(NewQTimerEventFromPointer(event))
	} else {
		NewQThreadPoolFromPointer(ptr).TimerEventDefault(NewQTimerEventFromPointer(event))
	}
}

func (ptr *QThreadPool) TimerEvent(event QTimerEvent_ITF) {
	defer qt.Recovering("QThreadPool::timerEvent")

	if ptr.Pointer() != nil {
		C.QThreadPool_TimerEvent(ptr.Pointer(), PointerFromQTimerEvent(event))
	}
}

func (ptr *QThreadPool) TimerEventDefault(event QTimerEvent_ITF) {
	defer qt.Recovering("QThreadPool::timerEvent")

	if ptr.Pointer() != nil {
		C.QThreadPool_TimerEventDefault(ptr.Pointer(), PointerFromQTimerEvent(event))
	}
}

func (ptr *QThreadPool) ConnectChildEvent(f func(event *QChildEvent)) {
	defer qt.Recovering("connect QThreadPool::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QThreadPool) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QThreadPool::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQThreadPoolChildEvent
func callbackQThreadPoolChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QThreadPool::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*QChildEvent))(NewQChildEventFromPointer(event))
	} else {
		NewQThreadPoolFromPointer(ptr).ChildEventDefault(NewQChildEventFromPointer(event))
	}
}

func (ptr *QThreadPool) ChildEvent(event QChildEvent_ITF) {
	defer qt.Recovering("QThreadPool::childEvent")

	if ptr.Pointer() != nil {
		C.QThreadPool_ChildEvent(ptr.Pointer(), PointerFromQChildEvent(event))
	}
}

func (ptr *QThreadPool) ChildEventDefault(event QChildEvent_ITF) {
	defer qt.Recovering("QThreadPool::childEvent")

	if ptr.Pointer() != nil {
		C.QThreadPool_ChildEventDefault(ptr.Pointer(), PointerFromQChildEvent(event))
	}
}

func (ptr *QThreadPool) ConnectCustomEvent(f func(event *QEvent)) {
	defer qt.Recovering("connect QThreadPool::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QThreadPool) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QThreadPool::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQThreadPoolCustomEvent
func callbackQThreadPoolCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QThreadPool::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
	} else {
		NewQThreadPoolFromPointer(ptr).CustomEventDefault(NewQEventFromPointer(event))
	}
}

func (ptr *QThreadPool) CustomEvent(event QEvent_ITF) {
	defer qt.Recovering("QThreadPool::customEvent")

	if ptr.Pointer() != nil {
		C.QThreadPool_CustomEvent(ptr.Pointer(), PointerFromQEvent(event))
	}
}

func (ptr *QThreadPool) CustomEventDefault(event QEvent_ITF) {
	defer qt.Recovering("QThreadPool::customEvent")

	if ptr.Pointer() != nil {
		C.QThreadPool_CustomEventDefault(ptr.Pointer(), PointerFromQEvent(event))
	}
}
