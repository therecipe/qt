package dbus

//#include "dbus.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QDBusConnectionInterface struct {
	QDBusAbstractInterface
}

type QDBusConnectionInterface_ITF interface {
	QDBusAbstractInterface_ITF
	QDBusConnectionInterface_PTR() *QDBusConnectionInterface
}

func PointerFromQDBusConnectionInterface(ptr QDBusConnectionInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusConnectionInterface_PTR().Pointer()
	}
	return nil
}

func NewQDBusConnectionInterfaceFromPointer(ptr unsafe.Pointer) *QDBusConnectionInterface {
	var n = new(QDBusConnectionInterface)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QDBusConnectionInterface_") {
		n.SetObjectName("QDBusConnectionInterface_" + qt.Identifier())
	}
	return n
}

func (ptr *QDBusConnectionInterface) QDBusConnectionInterface_PTR() *QDBusConnectionInterface {
	return ptr
}

//QDBusConnectionInterface::RegisterServiceReply
type QDBusConnectionInterface__RegisterServiceReply int64

const (
	QDBusConnectionInterface__ServiceNotRegistered = QDBusConnectionInterface__RegisterServiceReply(0)
	QDBusConnectionInterface__ServiceRegistered    = QDBusConnectionInterface__RegisterServiceReply(1)
	QDBusConnectionInterface__ServiceQueued        = QDBusConnectionInterface__RegisterServiceReply(2)
)

//QDBusConnectionInterface::ServiceQueueOptions
type QDBusConnectionInterface__ServiceQueueOptions int64

const (
	QDBusConnectionInterface__DontQueueService       = QDBusConnectionInterface__ServiceQueueOptions(0)
	QDBusConnectionInterface__QueueService           = QDBusConnectionInterface__ServiceQueueOptions(1)
	QDBusConnectionInterface__ReplaceExistingService = QDBusConnectionInterface__ServiceQueueOptions(2)
)

//QDBusConnectionInterface::ServiceReplacementOptions
type QDBusConnectionInterface__ServiceReplacementOptions int64

const (
	QDBusConnectionInterface__DontAllowReplacement = QDBusConnectionInterface__ServiceReplacementOptions(0)
	QDBusConnectionInterface__AllowReplacement     = QDBusConnectionInterface__ServiceReplacementOptions(1)
)

func (ptr *QDBusConnectionInterface) ConnectServiceRegistered(f func(serviceName string)) {
	defer qt.Recovering("connect QDBusConnectionInterface::serviceRegistered")

	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_ConnectServiceRegistered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "serviceRegistered", f)
	}
}

func (ptr *QDBusConnectionInterface) DisconnectServiceRegistered() {
	defer qt.Recovering("disconnect QDBusConnectionInterface::serviceRegistered")

	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_DisconnectServiceRegistered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "serviceRegistered")
	}
}

//export callbackQDBusConnectionInterfaceServiceRegistered
func callbackQDBusConnectionInterfaceServiceRegistered(ptr unsafe.Pointer, ptrName *C.char, serviceName *C.char) {
	defer qt.Recovering("callback QDBusConnectionInterface::serviceRegistered")

	if signal := qt.GetSignal(C.GoString(ptrName), "serviceRegistered"); signal != nil {
		signal.(func(string))(C.GoString(serviceName))
	}

}

func (ptr *QDBusConnectionInterface) ServiceRegistered(serviceName string) {
	defer qt.Recovering("QDBusConnectionInterface::serviceRegistered")

	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_ServiceRegistered(ptr.Pointer(), C.CString(serviceName))
	}
}

func (ptr *QDBusConnectionInterface) ConnectServiceUnregistered(f func(serviceName string)) {
	defer qt.Recovering("connect QDBusConnectionInterface::serviceUnregistered")

	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_ConnectServiceUnregistered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "serviceUnregistered", f)
	}
}

func (ptr *QDBusConnectionInterface) DisconnectServiceUnregistered() {
	defer qt.Recovering("disconnect QDBusConnectionInterface::serviceUnregistered")

	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_DisconnectServiceUnregistered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "serviceUnregistered")
	}
}

//export callbackQDBusConnectionInterfaceServiceUnregistered
func callbackQDBusConnectionInterfaceServiceUnregistered(ptr unsafe.Pointer, ptrName *C.char, serviceName *C.char) {
	defer qt.Recovering("callback QDBusConnectionInterface::serviceUnregistered")

	if signal := qt.GetSignal(C.GoString(ptrName), "serviceUnregistered"); signal != nil {
		signal.(func(string))(C.GoString(serviceName))
	}

}

func (ptr *QDBusConnectionInterface) ServiceUnregistered(serviceName string) {
	defer qt.Recovering("QDBusConnectionInterface::serviceUnregistered")

	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_ServiceUnregistered(ptr.Pointer(), C.CString(serviceName))
	}
}

func (ptr *QDBusConnectionInterface) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDBusConnectionInterface::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QDBusConnectionInterface) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDBusConnectionInterface::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQDBusConnectionInterfaceTimerEvent
func callbackQDBusConnectionInterfaceTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusConnectionInterface::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDBusConnectionInterfaceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDBusConnectionInterface) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDBusConnectionInterface::timerEvent")

	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDBusConnectionInterface) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDBusConnectionInterface::timerEvent")

	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDBusConnectionInterface) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDBusConnectionInterface::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QDBusConnectionInterface) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDBusConnectionInterface::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQDBusConnectionInterfaceChildEvent
func callbackQDBusConnectionInterfaceChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusConnectionInterface::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDBusConnectionInterfaceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDBusConnectionInterface) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDBusConnectionInterface::childEvent")

	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDBusConnectionInterface) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDBusConnectionInterface::childEvent")

	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDBusConnectionInterface) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDBusConnectionInterface::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QDBusConnectionInterface) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDBusConnectionInterface::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQDBusConnectionInterfaceCustomEvent
func callbackQDBusConnectionInterfaceCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusConnectionInterface::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDBusConnectionInterfaceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDBusConnectionInterface) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDBusConnectionInterface::customEvent")

	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDBusConnectionInterface) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDBusConnectionInterface::customEvent")

	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
