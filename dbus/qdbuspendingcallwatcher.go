package dbus

//#include "dbus.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QDBusPendingCallWatcher struct {
	core.QObject
	QDBusPendingCall
}

type QDBusPendingCallWatcher_ITF interface {
	core.QObject_ITF
	QDBusPendingCall_ITF
	QDBusPendingCallWatcher_PTR() *QDBusPendingCallWatcher
}

func (p *QDBusPendingCallWatcher) Pointer() unsafe.Pointer {
	return p.QObject_PTR().Pointer()
}

func (p *QDBusPendingCallWatcher) SetPointer(ptr unsafe.Pointer) {
	p.QObject_PTR().SetPointer(ptr)
	p.QDBusPendingCall_PTR().SetPointer(ptr)
}

func PointerFromQDBusPendingCallWatcher(ptr QDBusPendingCallWatcher_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusPendingCallWatcher_PTR().Pointer()
	}
	return nil
}

func NewQDBusPendingCallWatcherFromPointer(ptr unsafe.Pointer) *QDBusPendingCallWatcher {
	var n = new(QDBusPendingCallWatcher)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QDBusPendingCallWatcher_") {
		n.SetObjectName("QDBusPendingCallWatcher_" + qt.Identifier())
	}
	return n
}

func (ptr *QDBusPendingCallWatcher) QDBusPendingCallWatcher_PTR() *QDBusPendingCallWatcher {
	return ptr
}

func (ptr *QDBusPendingCallWatcher) WaitForFinished() {
	defer qt.Recovering("QDBusPendingCallWatcher::waitForFinished")

	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_WaitForFinished(ptr.Pointer())
	}
}

func NewQDBusPendingCallWatcher(call QDBusPendingCall_ITF, parent core.QObject_ITF) *QDBusPendingCallWatcher {
	defer qt.Recovering("QDBusPendingCallWatcher::QDBusPendingCallWatcher")

	return NewQDBusPendingCallWatcherFromPointer(C.QDBusPendingCallWatcher_NewQDBusPendingCallWatcher(PointerFromQDBusPendingCall(call), core.PointerFromQObject(parent)))
}

func (ptr *QDBusPendingCallWatcher) ConnectFinished(f func(self *QDBusPendingCallWatcher)) {
	defer qt.Recovering("connect QDBusPendingCallWatcher::finished")

	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QDBusPendingCallWatcher) DisconnectFinished() {
	defer qt.Recovering("disconnect QDBusPendingCallWatcher::finished")

	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQDBusPendingCallWatcherFinished
func callbackQDBusPendingCallWatcherFinished(ptr unsafe.Pointer, ptrName *C.char, self unsafe.Pointer) {
	defer qt.Recovering("callback QDBusPendingCallWatcher::finished")

	if signal := qt.GetSignal(C.GoString(ptrName), "finished"); signal != nil {
		signal.(func(*QDBusPendingCallWatcher))(NewQDBusPendingCallWatcherFromPointer(self))
	}

}

func (ptr *QDBusPendingCallWatcher) Finished(self QDBusPendingCallWatcher_ITF) {
	defer qt.Recovering("QDBusPendingCallWatcher::finished")

	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_Finished(ptr.Pointer(), PointerFromQDBusPendingCallWatcher(self))
	}
}

func (ptr *QDBusPendingCallWatcher) IsFinished() bool {
	defer qt.Recovering("QDBusPendingCallWatcher::isFinished")

	if ptr.Pointer() != nil {
		return C.QDBusPendingCallWatcher_IsFinished(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusPendingCallWatcher) DestroyQDBusPendingCallWatcher() {
	defer qt.Recovering("QDBusPendingCallWatcher::~QDBusPendingCallWatcher")

	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_DestroyQDBusPendingCallWatcher(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDBusPendingCallWatcher) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDBusPendingCallWatcher::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QDBusPendingCallWatcher) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDBusPendingCallWatcher::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQDBusPendingCallWatcherTimerEvent
func callbackQDBusPendingCallWatcherTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusPendingCallWatcher::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDBusPendingCallWatcherFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDBusPendingCallWatcher) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDBusPendingCallWatcher::timerEvent")

	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDBusPendingCallWatcher) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDBusPendingCallWatcher::timerEvent")

	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDBusPendingCallWatcher) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDBusPendingCallWatcher::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QDBusPendingCallWatcher) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDBusPendingCallWatcher::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQDBusPendingCallWatcherChildEvent
func callbackQDBusPendingCallWatcherChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusPendingCallWatcher::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDBusPendingCallWatcherFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDBusPendingCallWatcher) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDBusPendingCallWatcher::childEvent")

	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDBusPendingCallWatcher) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDBusPendingCallWatcher::childEvent")

	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDBusPendingCallWatcher) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDBusPendingCallWatcher::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QDBusPendingCallWatcher) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDBusPendingCallWatcher::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQDBusPendingCallWatcherCustomEvent
func callbackQDBusPendingCallWatcherCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusPendingCallWatcher::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDBusPendingCallWatcherFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDBusPendingCallWatcher) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDBusPendingCallWatcher::customEvent")

	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDBusPendingCallWatcher) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDBusPendingCallWatcher::customEvent")

	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
