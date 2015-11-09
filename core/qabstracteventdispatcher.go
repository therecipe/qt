package core

//#include "qabstracteventdispatcher.h"
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
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QAbstractEventDispatcher_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAbstractEventDispatcher) QAbstractEventDispatcher_PTR() *QAbstractEventDispatcher {
	return ptr
}

func (ptr *QAbstractEventDispatcher) ConnectAboutToBlock(f func()) {
	if ptr.Pointer() != nil {
		C.QAbstractEventDispatcher_ConnectAboutToBlock(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "aboutToBlock", f)
	}
}

func (ptr *QAbstractEventDispatcher) DisconnectAboutToBlock() {
	if ptr.Pointer() != nil {
		C.QAbstractEventDispatcher_DisconnectAboutToBlock(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "aboutToBlock")
	}
}

//export callbackQAbstractEventDispatcherAboutToBlock
func callbackQAbstractEventDispatcherAboutToBlock(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "aboutToBlock").(func())()
}

func (ptr *QAbstractEventDispatcher) ConnectAwake(f func()) {
	if ptr.Pointer() != nil {
		C.QAbstractEventDispatcher_ConnectAwake(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "awake", f)
	}
}

func (ptr *QAbstractEventDispatcher) DisconnectAwake() {
	if ptr.Pointer() != nil {
		C.QAbstractEventDispatcher_DisconnectAwake(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "awake")
	}
}

//export callbackQAbstractEventDispatcherAwake
func callbackQAbstractEventDispatcherAwake(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "awake").(func())()
}

func (ptr *QAbstractEventDispatcher) Flush() {
	if ptr.Pointer() != nil {
		C.QAbstractEventDispatcher_Flush(ptr.Pointer())
	}
}

func (ptr *QAbstractEventDispatcher) InstallNativeEventFilter(filterObj QAbstractNativeEventFilter_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractEventDispatcher_InstallNativeEventFilter(ptr.Pointer(), PointerFromQAbstractNativeEventFilter(filterObj))
	}
}

func QAbstractEventDispatcher_Instance(thread QThread_ITF) *QAbstractEventDispatcher {
	return NewQAbstractEventDispatcherFromPointer(C.QAbstractEventDispatcher_QAbstractEventDispatcher_Instance(PointerFromQThread(thread)))
}

func (ptr *QAbstractEventDispatcher) Interrupt() {
	if ptr.Pointer() != nil {
		C.QAbstractEventDispatcher_Interrupt(ptr.Pointer())
	}
}

func (ptr *QAbstractEventDispatcher) ProcessEvents(flags QEventLoop__ProcessEventsFlag) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractEventDispatcher_ProcessEvents(ptr.Pointer(), C.int(flags)) != 0
	}
	return false
}

func (ptr *QAbstractEventDispatcher) RegisterSocketNotifier(notifier QSocketNotifier_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractEventDispatcher_RegisterSocketNotifier(ptr.Pointer(), PointerFromQSocketNotifier(notifier))
	}
}

func (ptr *QAbstractEventDispatcher) RemainingTime(timerId int) int {
	if ptr.Pointer() != nil {
		return int(C.QAbstractEventDispatcher_RemainingTime(ptr.Pointer(), C.int(timerId)))
	}
	return 0
}

func (ptr *QAbstractEventDispatcher) RemoveNativeEventFilter(filter QAbstractNativeEventFilter_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractEventDispatcher_RemoveNativeEventFilter(ptr.Pointer(), PointerFromQAbstractNativeEventFilter(filter))
	}
}

func (ptr *QAbstractEventDispatcher) UnregisterSocketNotifier(notifier QSocketNotifier_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractEventDispatcher_UnregisterSocketNotifier(ptr.Pointer(), PointerFromQSocketNotifier(notifier))
	}
}

func (ptr *QAbstractEventDispatcher) UnregisterTimer(timerId int) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractEventDispatcher_UnregisterTimer(ptr.Pointer(), C.int(timerId)) != 0
	}
	return false
}

func (ptr *QAbstractEventDispatcher) UnregisterTimers(object QObject_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractEventDispatcher_UnregisterTimers(ptr.Pointer(), PointerFromQObject(object)) != 0
	}
	return false
}

func (ptr *QAbstractEventDispatcher) WakeUp() {
	if ptr.Pointer() != nil {
		C.QAbstractEventDispatcher_WakeUp(ptr.Pointer())
	}
}

func (ptr *QAbstractEventDispatcher) DestroyQAbstractEventDispatcher() {
	if ptr.Pointer() != nil {
		C.QAbstractEventDispatcher_DestroyQAbstractEventDispatcher(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
