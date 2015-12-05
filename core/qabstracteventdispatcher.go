package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
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
		n.SetObjectName("QAbstractEventDispatcher_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAbstractEventDispatcher) QAbstractEventDispatcher_PTR() *QAbstractEventDispatcher {
	return ptr
}

func (ptr *QAbstractEventDispatcher) ConnectAboutToBlock(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractEventDispatcher::aboutToBlock")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractEventDispatcher_ConnectAboutToBlock(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "aboutToBlock", f)
	}
}

func (ptr *QAbstractEventDispatcher) DisconnectAboutToBlock() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractEventDispatcher::aboutToBlock")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractEventDispatcher_DisconnectAboutToBlock(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "aboutToBlock")
	}
}

//export callbackQAbstractEventDispatcherAboutToBlock
func callbackQAbstractEventDispatcherAboutToBlock(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractEventDispatcher::aboutToBlock")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "aboutToBlock").(func())()
}

func (ptr *QAbstractEventDispatcher) ConnectAwake(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractEventDispatcher::awake")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractEventDispatcher_ConnectAwake(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "awake", f)
	}
}

func (ptr *QAbstractEventDispatcher) DisconnectAwake() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractEventDispatcher::awake")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractEventDispatcher_DisconnectAwake(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "awake")
	}
}

//export callbackQAbstractEventDispatcherAwake
func callbackQAbstractEventDispatcherAwake(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractEventDispatcher::awake")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "awake").(func())()
}

func (ptr *QAbstractEventDispatcher) Flush() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractEventDispatcher::flush")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractEventDispatcher_Flush(ptr.Pointer())
	}
}

func (ptr *QAbstractEventDispatcher) InstallNativeEventFilter(filterObj QAbstractNativeEventFilter_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractEventDispatcher::installNativeEventFilter")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractEventDispatcher_InstallNativeEventFilter(ptr.Pointer(), PointerFromQAbstractNativeEventFilter(filterObj))
	}
}

func QAbstractEventDispatcher_Instance(thread QThread_ITF) *QAbstractEventDispatcher {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractEventDispatcher::instance")
		}
	}()

	return NewQAbstractEventDispatcherFromPointer(C.QAbstractEventDispatcher_QAbstractEventDispatcher_Instance(PointerFromQThread(thread)))
}

func (ptr *QAbstractEventDispatcher) Interrupt() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractEventDispatcher::interrupt")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractEventDispatcher_Interrupt(ptr.Pointer())
	}
}

func (ptr *QAbstractEventDispatcher) ProcessEvents(flags QEventLoop__ProcessEventsFlag) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractEventDispatcher::processEvents")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAbstractEventDispatcher_ProcessEvents(ptr.Pointer(), C.int(flags)) != 0
	}
	return false
}

func (ptr *QAbstractEventDispatcher) RegisterSocketNotifier(notifier QSocketNotifier_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractEventDispatcher::registerSocketNotifier")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractEventDispatcher_RegisterSocketNotifier(ptr.Pointer(), PointerFromQSocketNotifier(notifier))
	}
}

func (ptr *QAbstractEventDispatcher) RemainingTime(timerId int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractEventDispatcher::remainingTime")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QAbstractEventDispatcher_RemainingTime(ptr.Pointer(), C.int(timerId)))
	}
	return 0
}

func (ptr *QAbstractEventDispatcher) RemoveNativeEventFilter(filter QAbstractNativeEventFilter_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractEventDispatcher::removeNativeEventFilter")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractEventDispatcher_RemoveNativeEventFilter(ptr.Pointer(), PointerFromQAbstractNativeEventFilter(filter))
	}
}

func (ptr *QAbstractEventDispatcher) UnregisterSocketNotifier(notifier QSocketNotifier_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractEventDispatcher::unregisterSocketNotifier")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractEventDispatcher_UnregisterSocketNotifier(ptr.Pointer(), PointerFromQSocketNotifier(notifier))
	}
}

func (ptr *QAbstractEventDispatcher) UnregisterTimer(timerId int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractEventDispatcher::unregisterTimer")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAbstractEventDispatcher_UnregisterTimer(ptr.Pointer(), C.int(timerId)) != 0
	}
	return false
}

func (ptr *QAbstractEventDispatcher) UnregisterTimers(object QObject_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractEventDispatcher::unregisterTimers")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAbstractEventDispatcher_UnregisterTimers(ptr.Pointer(), PointerFromQObject(object)) != 0
	}
	return false
}

func (ptr *QAbstractEventDispatcher) WakeUp() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractEventDispatcher::wakeUp")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractEventDispatcher_WakeUp(ptr.Pointer())
	}
}

func (ptr *QAbstractEventDispatcher) DestroyQAbstractEventDispatcher() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractEventDispatcher::~QAbstractEventDispatcher")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractEventDispatcher_DestroyQAbstractEventDispatcher(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
