package dbus

//#include "qdbusservicewatcher.h"
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

type QDBusServiceWatcherITF interface {
	core.QObjectITF
	QDBusServiceWatcherPTR() *QDBusServiceWatcher
}

func PointerFromQDBusServiceWatcher(ptr QDBusServiceWatcherITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusServiceWatcherPTR().Pointer()
	}
	return nil
}

func QDBusServiceWatcherFromPointer(ptr unsafe.Pointer) *QDBusServiceWatcher {
	var n = new(QDBusServiceWatcher)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QDBusServiceWatcher_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QDBusServiceWatcher) QDBusServiceWatcherPTR() *QDBusServiceWatcher {
	return ptr
}

//QDBusServiceWatcher::WatchModeFlag
type QDBusServiceWatcher__WatchModeFlag int

var (
	QDBusServiceWatcher__WatchForRegistration   = QDBusServiceWatcher__WatchModeFlag(0x01)
	QDBusServiceWatcher__WatchForUnregistration = QDBusServiceWatcher__WatchModeFlag(0x02)
	QDBusServiceWatcher__WatchForOwnerChange    = QDBusServiceWatcher__WatchModeFlag(0x03)
)

func (ptr *QDBusServiceWatcher) SetWatchMode(mode QDBusServiceWatcher__WatchModeFlag) {
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_SetWatchMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QDBusServiceWatcher) WatchMode() QDBusServiceWatcher__WatchModeFlag {
	if ptr.Pointer() != nil {
		return QDBusServiceWatcher__WatchModeFlag(C.QDBusServiceWatcher_WatchMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func NewQDBusServiceWatcher(parent core.QObjectITF) *QDBusServiceWatcher {
	return QDBusServiceWatcherFromPointer(unsafe.Pointer(C.QDBusServiceWatcher_NewQDBusServiceWatcher(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func NewQDBusServiceWatcher2(service string, connection QDBusConnectionITF, watchMode QDBusServiceWatcher__WatchModeFlag, parent core.QObjectITF) *QDBusServiceWatcher {
	return QDBusServiceWatcherFromPointer(unsafe.Pointer(C.QDBusServiceWatcher_NewQDBusServiceWatcher2(C.CString(service), C.QtObjectPtr(PointerFromQDBusConnection(connection)), C.int(watchMode), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QDBusServiceWatcher) AddWatchedService(newService string) {
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_AddWatchedService(C.QtObjectPtr(ptr.Pointer()), C.CString(newService))
	}
}

func (ptr *QDBusServiceWatcher) RemoveWatchedService(service string) bool {
	if ptr.Pointer() != nil {
		return C.QDBusServiceWatcher_RemoveWatchedService(C.QtObjectPtr(ptr.Pointer()), C.CString(service)) != 0
	}
	return false
}

func (ptr *QDBusServiceWatcher) ConnectServiceOwnerChanged(f func(serviceName string, oldOwner string, newOwner string)) {
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_ConnectServiceOwnerChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "serviceOwnerChanged", f)
	}
}

func (ptr *QDBusServiceWatcher) DisconnectServiceOwnerChanged() {
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_DisconnectServiceOwnerChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "serviceOwnerChanged")
	}
}

//export callbackQDBusServiceWatcherServiceOwnerChanged
func callbackQDBusServiceWatcherServiceOwnerChanged(ptrName *C.char, serviceName *C.char, oldOwner *C.char, newOwner *C.char) {
	qt.GetSignal(C.GoString(ptrName), "serviceOwnerChanged").(func(string, string, string))(C.GoString(serviceName), C.GoString(oldOwner), C.GoString(newOwner))
}

func (ptr *QDBusServiceWatcher) ConnectServiceRegistered(f func(serviceName string)) {
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_ConnectServiceRegistered(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "serviceRegistered", f)
	}
}

func (ptr *QDBusServiceWatcher) DisconnectServiceRegistered() {
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_DisconnectServiceRegistered(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "serviceRegistered")
	}
}

//export callbackQDBusServiceWatcherServiceRegistered
func callbackQDBusServiceWatcherServiceRegistered(ptrName *C.char, serviceName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "serviceRegistered").(func(string))(C.GoString(serviceName))
}

func (ptr *QDBusServiceWatcher) ConnectServiceUnregistered(f func(serviceName string)) {
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_ConnectServiceUnregistered(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "serviceUnregistered", f)
	}
}

func (ptr *QDBusServiceWatcher) DisconnectServiceUnregistered() {
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_DisconnectServiceUnregistered(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "serviceUnregistered")
	}
}

//export callbackQDBusServiceWatcherServiceUnregistered
func callbackQDBusServiceWatcherServiceUnregistered(ptrName *C.char, serviceName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "serviceUnregistered").(func(string))(C.GoString(serviceName))
}

func (ptr *QDBusServiceWatcher) SetConnection(connection QDBusConnectionITF) {
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_SetConnection(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQDBusConnection(connection)))
	}
}

func (ptr *QDBusServiceWatcher) SetWatchedServices(services []string) {
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_SetWatchedServices(C.QtObjectPtr(ptr.Pointer()), C.CString(strings.Join(services, "|")))
	}
}

func (ptr *QDBusServiceWatcher) WatchedServices() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QDBusServiceWatcher_WatchedServices(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QDBusServiceWatcher) DestroyQDBusServiceWatcher() {
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_DestroyQDBusServiceWatcher(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
