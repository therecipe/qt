package dbus

//#include "dbus.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QDBusServiceWatcher struct {
	core.QObject
}

type QDBusServiceWatcher_ITF interface {
	core.QObject_ITF
	QDBusServiceWatcher_PTR() *QDBusServiceWatcher
}

func PointerFromQDBusServiceWatcher(ptr QDBusServiceWatcher_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusServiceWatcher_PTR().Pointer()
	}
	return nil
}

func NewQDBusServiceWatcherFromPointer(ptr unsafe.Pointer) *QDBusServiceWatcher {
	var n = new(QDBusServiceWatcher)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QDBusServiceWatcher_") {
		n.SetObjectName("QDBusServiceWatcher_" + qt.Identifier())
	}
	return n
}

func (ptr *QDBusServiceWatcher) QDBusServiceWatcher_PTR() *QDBusServiceWatcher {
	return ptr
}

//QDBusServiceWatcher::WatchModeFlag
type QDBusServiceWatcher__WatchModeFlag int64

const (
	QDBusServiceWatcher__WatchForRegistration   = QDBusServiceWatcher__WatchModeFlag(0x01)
	QDBusServiceWatcher__WatchForUnregistration = QDBusServiceWatcher__WatchModeFlag(0x02)
	QDBusServiceWatcher__WatchForOwnerChange    = QDBusServiceWatcher__WatchModeFlag(0x03)
)

func (ptr *QDBusServiceWatcher) SetWatchMode(mode QDBusServiceWatcher__WatchModeFlag) {
	defer qt.Recovering("QDBusServiceWatcher::setWatchMode")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_SetWatchMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QDBusServiceWatcher) WatchMode() QDBusServiceWatcher__WatchModeFlag {
	defer qt.Recovering("QDBusServiceWatcher::watchMode")

	if ptr.Pointer() != nil {
		return QDBusServiceWatcher__WatchModeFlag(C.QDBusServiceWatcher_WatchMode(ptr.Pointer()))
	}
	return 0
}

func NewQDBusServiceWatcher(parent core.QObject_ITF) *QDBusServiceWatcher {
	defer qt.Recovering("QDBusServiceWatcher::QDBusServiceWatcher")

	return NewQDBusServiceWatcherFromPointer(C.QDBusServiceWatcher_NewQDBusServiceWatcher(core.PointerFromQObject(parent)))
}

func NewQDBusServiceWatcher2(service string, connection QDBusConnection_ITF, watchMode QDBusServiceWatcher__WatchModeFlag, parent core.QObject_ITF) *QDBusServiceWatcher {
	defer qt.Recovering("QDBusServiceWatcher::QDBusServiceWatcher")

	return NewQDBusServiceWatcherFromPointer(C.QDBusServiceWatcher_NewQDBusServiceWatcher2(C.CString(service), PointerFromQDBusConnection(connection), C.int(watchMode), core.PointerFromQObject(parent)))
}

func (ptr *QDBusServiceWatcher) AddWatchedService(newService string) {
	defer qt.Recovering("QDBusServiceWatcher::addWatchedService")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_AddWatchedService(ptr.Pointer(), C.CString(newService))
	}
}

func (ptr *QDBusServiceWatcher) RemoveWatchedService(service string) bool {
	defer qt.Recovering("QDBusServiceWatcher::removeWatchedService")

	if ptr.Pointer() != nil {
		return C.QDBusServiceWatcher_RemoveWatchedService(ptr.Pointer(), C.CString(service)) != 0
	}
	return false
}

func (ptr *QDBusServiceWatcher) ConnectServiceOwnerChanged(f func(serviceName string, oldOwner string, newOwner string)) {
	defer qt.Recovering("connect QDBusServiceWatcher::serviceOwnerChanged")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_ConnectServiceOwnerChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "serviceOwnerChanged", f)
	}
}

func (ptr *QDBusServiceWatcher) DisconnectServiceOwnerChanged() {
	defer qt.Recovering("disconnect QDBusServiceWatcher::serviceOwnerChanged")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_DisconnectServiceOwnerChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "serviceOwnerChanged")
	}
}

//export callbackQDBusServiceWatcherServiceOwnerChanged
func callbackQDBusServiceWatcherServiceOwnerChanged(ptr unsafe.Pointer, ptrName *C.char, serviceName *C.char, oldOwner *C.char, newOwner *C.char) {
	defer qt.Recovering("callback QDBusServiceWatcher::serviceOwnerChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "serviceOwnerChanged"); signal != nil {
		signal.(func(string, string, string))(C.GoString(serviceName), C.GoString(oldOwner), C.GoString(newOwner))
	}

}

func (ptr *QDBusServiceWatcher) ServiceOwnerChanged(serviceName string, oldOwner string, newOwner string) {
	defer qt.Recovering("QDBusServiceWatcher::serviceOwnerChanged")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_ServiceOwnerChanged(ptr.Pointer(), C.CString(serviceName), C.CString(oldOwner), C.CString(newOwner))
	}
}

func (ptr *QDBusServiceWatcher) ConnectServiceRegistered(f func(serviceName string)) {
	defer qt.Recovering("connect QDBusServiceWatcher::serviceRegistered")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_ConnectServiceRegistered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "serviceRegistered", f)
	}
}

func (ptr *QDBusServiceWatcher) DisconnectServiceRegistered() {
	defer qt.Recovering("disconnect QDBusServiceWatcher::serviceRegistered")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_DisconnectServiceRegistered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "serviceRegistered")
	}
}

//export callbackQDBusServiceWatcherServiceRegistered
func callbackQDBusServiceWatcherServiceRegistered(ptr unsafe.Pointer, ptrName *C.char, serviceName *C.char) {
	defer qt.Recovering("callback QDBusServiceWatcher::serviceRegistered")

	if signal := qt.GetSignal(C.GoString(ptrName), "serviceRegistered"); signal != nil {
		signal.(func(string))(C.GoString(serviceName))
	}

}

func (ptr *QDBusServiceWatcher) ServiceRegistered(serviceName string) {
	defer qt.Recovering("QDBusServiceWatcher::serviceRegistered")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_ServiceRegistered(ptr.Pointer(), C.CString(serviceName))
	}
}

func (ptr *QDBusServiceWatcher) ConnectServiceUnregistered(f func(serviceName string)) {
	defer qt.Recovering("connect QDBusServiceWatcher::serviceUnregistered")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_ConnectServiceUnregistered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "serviceUnregistered", f)
	}
}

func (ptr *QDBusServiceWatcher) DisconnectServiceUnregistered() {
	defer qt.Recovering("disconnect QDBusServiceWatcher::serviceUnregistered")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_DisconnectServiceUnregistered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "serviceUnregistered")
	}
}

//export callbackQDBusServiceWatcherServiceUnregistered
func callbackQDBusServiceWatcherServiceUnregistered(ptr unsafe.Pointer, ptrName *C.char, serviceName *C.char) {
	defer qt.Recovering("callback QDBusServiceWatcher::serviceUnregistered")

	if signal := qt.GetSignal(C.GoString(ptrName), "serviceUnregistered"); signal != nil {
		signal.(func(string))(C.GoString(serviceName))
	}

}

func (ptr *QDBusServiceWatcher) ServiceUnregistered(serviceName string) {
	defer qt.Recovering("QDBusServiceWatcher::serviceUnregistered")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_ServiceUnregistered(ptr.Pointer(), C.CString(serviceName))
	}
}

func (ptr *QDBusServiceWatcher) SetConnection(connection QDBusConnection_ITF) {
	defer qt.Recovering("QDBusServiceWatcher::setConnection")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_SetConnection(ptr.Pointer(), PointerFromQDBusConnection(connection))
	}
}

func (ptr *QDBusServiceWatcher) SetWatchedServices(services []string) {
	defer qt.Recovering("QDBusServiceWatcher::setWatchedServices")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_SetWatchedServices(ptr.Pointer(), C.CString(strings.Join(services, ",,,")))
	}
}

func (ptr *QDBusServiceWatcher) WatchedServices() []string {
	defer qt.Recovering("QDBusServiceWatcher::watchedServices")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QDBusServiceWatcher_WatchedServices(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QDBusServiceWatcher) DestroyQDBusServiceWatcher() {
	defer qt.Recovering("QDBusServiceWatcher::~QDBusServiceWatcher")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_DestroyQDBusServiceWatcher(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDBusServiceWatcher) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDBusServiceWatcher::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QDBusServiceWatcher) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDBusServiceWatcher::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQDBusServiceWatcherTimerEvent
func callbackQDBusServiceWatcherTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusServiceWatcher::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDBusServiceWatcherFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDBusServiceWatcher) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDBusServiceWatcher::timerEvent")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDBusServiceWatcher) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDBusServiceWatcher::timerEvent")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDBusServiceWatcher) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDBusServiceWatcher::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QDBusServiceWatcher) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDBusServiceWatcher::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQDBusServiceWatcherChildEvent
func callbackQDBusServiceWatcherChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusServiceWatcher::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDBusServiceWatcherFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDBusServiceWatcher) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDBusServiceWatcher::childEvent")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDBusServiceWatcher) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDBusServiceWatcher::childEvent")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDBusServiceWatcher) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDBusServiceWatcher::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QDBusServiceWatcher) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDBusServiceWatcher::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQDBusServiceWatcherCustomEvent
func callbackQDBusServiceWatcherCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusServiceWatcher::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDBusServiceWatcherFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDBusServiceWatcher) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDBusServiceWatcher::customEvent")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDBusServiceWatcher) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDBusServiceWatcher::customEvent")

	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
