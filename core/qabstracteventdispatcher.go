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

type QAbstractEventDispatcherITF interface {
	QObjectITF
	QAbstractEventDispatcherPTR() *QAbstractEventDispatcher
}

func PointerFromQAbstractEventDispatcher(ptr QAbstractEventDispatcherITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractEventDispatcherPTR().Pointer()
	}
	return nil
}

func QAbstractEventDispatcherFromPointer(ptr unsafe.Pointer) *QAbstractEventDispatcher {
	var n = new(QAbstractEventDispatcher)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QAbstractEventDispatcher_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAbstractEventDispatcher) QAbstractEventDispatcherPTR() *QAbstractEventDispatcher {
	return ptr
}

func (ptr *QAbstractEventDispatcher) ConnectAboutToBlock(f func()) {
	if ptr.Pointer() != nil {
		C.QAbstractEventDispatcher_ConnectAboutToBlock(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "aboutToBlock", f)
	}
}

func (ptr *QAbstractEventDispatcher) DisconnectAboutToBlock() {
	if ptr.Pointer() != nil {
		C.QAbstractEventDispatcher_DisconnectAboutToBlock(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "aboutToBlock")
	}
}

//export callbackQAbstractEventDispatcherAboutToBlock
func callbackQAbstractEventDispatcherAboutToBlock(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "aboutToBlock").(func())()
}

func (ptr *QAbstractEventDispatcher) ConnectAwake(f func()) {
	if ptr.Pointer() != nil {
		C.QAbstractEventDispatcher_ConnectAwake(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "awake", f)
	}
}

func (ptr *QAbstractEventDispatcher) DisconnectAwake() {
	if ptr.Pointer() != nil {
		C.QAbstractEventDispatcher_DisconnectAwake(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "awake")
	}
}

//export callbackQAbstractEventDispatcherAwake
func callbackQAbstractEventDispatcherAwake(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "awake").(func())()
}

func (ptr *QAbstractEventDispatcher) Flush() {
	if ptr.Pointer() != nil {
		C.QAbstractEventDispatcher_Flush(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QAbstractEventDispatcher) InstallNativeEventFilter(filterObj QAbstractNativeEventFilterITF) {
	if ptr.Pointer() != nil {
		C.QAbstractEventDispatcher_InstallNativeEventFilter(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAbstractNativeEventFilter(filterObj)))
	}
}

func QAbstractEventDispatcher_Instance(thread QThreadITF) *QAbstractEventDispatcher {
	return QAbstractEventDispatcherFromPointer(unsafe.Pointer(C.QAbstractEventDispatcher_QAbstractEventDispatcher_Instance(C.QtObjectPtr(PointerFromQThread(thread)))))
}

func (ptr *QAbstractEventDispatcher) Interrupt() {
	if ptr.Pointer() != nil {
		C.QAbstractEventDispatcher_Interrupt(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QAbstractEventDispatcher) ProcessEvents(flags QEventLoop__ProcessEventsFlag) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractEventDispatcher_ProcessEvents(C.QtObjectPtr(ptr.Pointer()), C.int(flags)) != 0
	}
	return false
}

func (ptr *QAbstractEventDispatcher) RegisterSocketNotifier(notifier QSocketNotifierITF) {
	if ptr.Pointer() != nil {
		C.QAbstractEventDispatcher_RegisterSocketNotifier(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSocketNotifier(notifier)))
	}
}

func (ptr *QAbstractEventDispatcher) RemainingTime(timerId int) int {
	if ptr.Pointer() != nil {
		return int(C.QAbstractEventDispatcher_RemainingTime(C.QtObjectPtr(ptr.Pointer()), C.int(timerId)))
	}
	return 0
}

func (ptr *QAbstractEventDispatcher) RemoveNativeEventFilter(filter QAbstractNativeEventFilterITF) {
	if ptr.Pointer() != nil {
		C.QAbstractEventDispatcher_RemoveNativeEventFilter(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAbstractNativeEventFilter(filter)))
	}
}

func (ptr *QAbstractEventDispatcher) UnregisterSocketNotifier(notifier QSocketNotifierITF) {
	if ptr.Pointer() != nil {
		C.QAbstractEventDispatcher_UnregisterSocketNotifier(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSocketNotifier(notifier)))
	}
}

func (ptr *QAbstractEventDispatcher) UnregisterTimer(timerId int) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractEventDispatcher_UnregisterTimer(C.QtObjectPtr(ptr.Pointer()), C.int(timerId)) != 0
	}
	return false
}

func (ptr *QAbstractEventDispatcher) UnregisterTimers(object QObjectITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractEventDispatcher_UnregisterTimers(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQObject(object))) != 0
	}
	return false
}

func (ptr *QAbstractEventDispatcher) WakeUp() {
	if ptr.Pointer() != nil {
		C.QAbstractEventDispatcher_WakeUp(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QAbstractEventDispatcher) DestroyQAbstractEventDispatcher() {
	if ptr.Pointer() != nil {
		C.QAbstractEventDispatcher_DestroyQAbstractEventDispatcher(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
