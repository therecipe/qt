package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QAbstractEventDispatcher struct {
	QObject
}

type QAbstractEventDispatcher_ITF interface {
	QObject_ITF
	QAbstractEventDispatcher_PTR() *QAbstractEventDispatcher
}

func PointerFromQAbstractEventDispatcher(ptr QAbstractEventDispatcher_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractEventDispatcher_PTR().Pointer()
	}
	return nil
}

func NewQAbstractEventDispatcherFromPointer(ptr unsafe.Pointer) *QAbstractEventDispatcher {
	var n = new(QAbstractEventDispatcher)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAbstractEventDispatcher_") {
		n.SetObjectName("QAbstractEventDispatcher_" + qt.Identifier())
	}
	return n
}

func (ptr *QAbstractEventDispatcher) QAbstractEventDispatcher_PTR() *QAbstractEventDispatcher {
	return ptr
}

func (ptr *QAbstractEventDispatcher) ConnectAboutToBlock(f func()) {
	defer qt.Recovering("connect QAbstractEventDispatcher::aboutToBlock")

	if ptr.Pointer() != nil {
		C.QAbstractEventDispatcher_ConnectAboutToBlock(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "aboutToBlock", f)
	}
}

func (ptr *QAbstractEventDispatcher) DisconnectAboutToBlock() {
	defer qt.Recovering("disconnect QAbstractEventDispatcher::aboutToBlock")

	if ptr.Pointer() != nil {
		C.QAbstractEventDispatcher_DisconnectAboutToBlock(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "aboutToBlock")
	}
}

//export callbackQAbstractEventDispatcherAboutToBlock
func callbackQAbstractEventDispatcherAboutToBlock(ptrName *C.char) {
	defer qt.Recovering("callback QAbstractEventDispatcher::aboutToBlock")

	var signal = qt.GetSignal(C.GoString(ptrName), "aboutToBlock")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractEventDispatcher) ConnectAwake(f func()) {
	defer qt.Recovering("connect QAbstractEventDispatcher::awake")

	if ptr.Pointer() != nil {
		C.QAbstractEventDispatcher_ConnectAwake(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "awake", f)
	}
}

func (ptr *QAbstractEventDispatcher) DisconnectAwake() {
	defer qt.Recovering("disconnect QAbstractEventDispatcher::awake")

	if ptr.Pointer() != nil {
		C.QAbstractEventDispatcher_DisconnectAwake(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "awake")
	}
}

//export callbackQAbstractEventDispatcherAwake
func callbackQAbstractEventDispatcherAwake(ptrName *C.char) {
	defer qt.Recovering("callback QAbstractEventDispatcher::awake")

	var signal = qt.GetSignal(C.GoString(ptrName), "awake")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractEventDispatcher) Flush() {
	defer qt.Recovering("QAbstractEventDispatcher::flush")

	if ptr.Pointer() != nil {
		C.QAbstractEventDispatcher_Flush(ptr.Pointer())
	}
}

func (ptr *QAbstractEventDispatcher) InstallNativeEventFilter(filterObj QAbstractNativeEventFilter_ITF) {
	defer qt.Recovering("QAbstractEventDispatcher::installNativeEventFilter")

	if ptr.Pointer() != nil {
		C.QAbstractEventDispatcher_InstallNativeEventFilter(ptr.Pointer(), PointerFromQAbstractNativeEventFilter(filterObj))
	}
}

func QAbstractEventDispatcher_Instance(thread QThread_ITF) *QAbstractEventDispatcher {
	defer qt.Recovering("QAbstractEventDispatcher::instance")

	return NewQAbstractEventDispatcherFromPointer(C.QAbstractEventDispatcher_QAbstractEventDispatcher_Instance(PointerFromQThread(thread)))
}

func (ptr *QAbstractEventDispatcher) Interrupt() {
	defer qt.Recovering("QAbstractEventDispatcher::interrupt")

	if ptr.Pointer() != nil {
		C.QAbstractEventDispatcher_Interrupt(ptr.Pointer())
	}
}

func (ptr *QAbstractEventDispatcher) ProcessEvents(flags QEventLoop__ProcessEventsFlag) bool {
	defer qt.Recovering("QAbstractEventDispatcher::processEvents")

	if ptr.Pointer() != nil {
		return C.QAbstractEventDispatcher_ProcessEvents(ptr.Pointer(), C.int(flags)) != 0
	}
	return false
}

func (ptr *QAbstractEventDispatcher) RegisterSocketNotifier(notifier QSocketNotifier_ITF) {
	defer qt.Recovering("QAbstractEventDispatcher::registerSocketNotifier")

	if ptr.Pointer() != nil {
		C.QAbstractEventDispatcher_RegisterSocketNotifier(ptr.Pointer(), PointerFromQSocketNotifier(notifier))
	}
}

func (ptr *QAbstractEventDispatcher) RemainingTime(timerId int) int {
	defer qt.Recovering("QAbstractEventDispatcher::remainingTime")

	if ptr.Pointer() != nil {
		return int(C.QAbstractEventDispatcher_RemainingTime(ptr.Pointer(), C.int(timerId)))
	}
	return 0
}

func (ptr *QAbstractEventDispatcher) RemoveNativeEventFilter(filter QAbstractNativeEventFilter_ITF) {
	defer qt.Recovering("QAbstractEventDispatcher::removeNativeEventFilter")

	if ptr.Pointer() != nil {
		C.QAbstractEventDispatcher_RemoveNativeEventFilter(ptr.Pointer(), PointerFromQAbstractNativeEventFilter(filter))
	}
}

func (ptr *QAbstractEventDispatcher) UnregisterSocketNotifier(notifier QSocketNotifier_ITF) {
	defer qt.Recovering("QAbstractEventDispatcher::unregisterSocketNotifier")

	if ptr.Pointer() != nil {
		C.QAbstractEventDispatcher_UnregisterSocketNotifier(ptr.Pointer(), PointerFromQSocketNotifier(notifier))
	}
}

func (ptr *QAbstractEventDispatcher) UnregisterTimer(timerId int) bool {
	defer qt.Recovering("QAbstractEventDispatcher::unregisterTimer")

	if ptr.Pointer() != nil {
		return C.QAbstractEventDispatcher_UnregisterTimer(ptr.Pointer(), C.int(timerId)) != 0
	}
	return false
}

func (ptr *QAbstractEventDispatcher) UnregisterTimers(object QObject_ITF) bool {
	defer qt.Recovering("QAbstractEventDispatcher::unregisterTimers")

	if ptr.Pointer() != nil {
		return C.QAbstractEventDispatcher_UnregisterTimers(ptr.Pointer(), PointerFromQObject(object)) != 0
	}
	return false
}

func (ptr *QAbstractEventDispatcher) WakeUp() {
	defer qt.Recovering("QAbstractEventDispatcher::wakeUp")

	if ptr.Pointer() != nil {
		C.QAbstractEventDispatcher_WakeUp(ptr.Pointer())
	}
}

func (ptr *QAbstractEventDispatcher) DestroyQAbstractEventDispatcher() {
	defer qt.Recovering("QAbstractEventDispatcher::~QAbstractEventDispatcher")

	if ptr.Pointer() != nil {
		C.QAbstractEventDispatcher_DestroyQAbstractEventDispatcher(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
