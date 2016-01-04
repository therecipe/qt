package dbus

//#include "dbus.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QDBusInterface struct {
	QDBusAbstractInterface
}

type QDBusInterface_ITF interface {
	QDBusAbstractInterface_ITF
	QDBusInterface_PTR() *QDBusInterface
}

func PointerFromQDBusInterface(ptr QDBusInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusInterface_PTR().Pointer()
	}
	return nil
}

func NewQDBusInterfaceFromPointer(ptr unsafe.Pointer) *QDBusInterface {
	var n = new(QDBusInterface)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QDBusInterface_") {
		n.SetObjectName("QDBusInterface_" + qt.Identifier())
	}
	return n
}

func (ptr *QDBusInterface) QDBusInterface_PTR() *QDBusInterface {
	return ptr
}

func NewQDBusInterface(service string, path string, interfa string, connection QDBusConnection_ITF, parent core.QObject_ITF) *QDBusInterface {
	defer qt.Recovering("QDBusInterface::QDBusInterface")

	return NewQDBusInterfaceFromPointer(C.QDBusInterface_NewQDBusInterface(C.CString(service), C.CString(path), C.CString(interfa), PointerFromQDBusConnection(connection), core.PointerFromQObject(parent)))
}

func (ptr *QDBusInterface) DestroyQDBusInterface() {
	defer qt.Recovering("QDBusInterface::~QDBusInterface")

	if ptr.Pointer() != nil {
		C.QDBusInterface_DestroyQDBusInterface(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDBusInterface) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDBusInterface::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QDBusInterface) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDBusInterface::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQDBusInterfaceTimerEvent
func callbackQDBusInterfaceTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusInterface::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDBusInterfaceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDBusInterface) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDBusInterface::timerEvent")

	if ptr.Pointer() != nil {
		C.QDBusInterface_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDBusInterface) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDBusInterface::timerEvent")

	if ptr.Pointer() != nil {
		C.QDBusInterface_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDBusInterface) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDBusInterface::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QDBusInterface) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDBusInterface::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQDBusInterfaceChildEvent
func callbackQDBusInterfaceChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusInterface::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDBusInterfaceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDBusInterface) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDBusInterface::childEvent")

	if ptr.Pointer() != nil {
		C.QDBusInterface_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDBusInterface) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDBusInterface::childEvent")

	if ptr.Pointer() != nil {
		C.QDBusInterface_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDBusInterface) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDBusInterface::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QDBusInterface) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDBusInterface::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQDBusInterfaceCustomEvent
func callbackQDBusInterfaceCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusInterface::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDBusInterfaceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDBusInterface) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDBusInterface::customEvent")

	if ptr.Pointer() != nil {
		C.QDBusInterface_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDBusInterface) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDBusInterface::customEvent")

	if ptr.Pointer() != nil {
		C.QDBusInterface_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
