package dbus

//#include "dbus.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QDBusVirtualObject struct {
	core.QObject
}

type QDBusVirtualObject_ITF interface {
	core.QObject_ITF
	QDBusVirtualObject_PTR() *QDBusVirtualObject
}

func PointerFromQDBusVirtualObject(ptr QDBusVirtualObject_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusVirtualObject_PTR().Pointer()
	}
	return nil
}

func NewQDBusVirtualObjectFromPointer(ptr unsafe.Pointer) *QDBusVirtualObject {
	var n = new(QDBusVirtualObject)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QDBusVirtualObject_") {
		n.SetObjectName("QDBusVirtualObject_" + qt.Identifier())
	}
	return n
}

func (ptr *QDBusVirtualObject) QDBusVirtualObject_PTR() *QDBusVirtualObject {
	return ptr
}

func (ptr *QDBusVirtualObject) HandleMessage(message QDBusMessage_ITF, connection QDBusConnection_ITF) bool {
	defer qt.Recovering("QDBusVirtualObject::handleMessage")

	if ptr.Pointer() != nil {
		return C.QDBusVirtualObject_HandleMessage(ptr.Pointer(), PointerFromQDBusMessage(message), PointerFromQDBusConnection(connection)) != 0
	}
	return false
}

func (ptr *QDBusVirtualObject) Introspect(path string) string {
	defer qt.Recovering("QDBusVirtualObject::introspect")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusVirtualObject_Introspect(ptr.Pointer(), C.CString(path)))
	}
	return ""
}

func (ptr *QDBusVirtualObject) DestroyQDBusVirtualObject() {
	defer qt.Recovering("QDBusVirtualObject::~QDBusVirtualObject")

	if ptr.Pointer() != nil {
		C.QDBusVirtualObject_DestroyQDBusVirtualObject(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDBusVirtualObject) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDBusVirtualObject::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QDBusVirtualObject) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDBusVirtualObject::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQDBusVirtualObjectTimerEvent
func callbackQDBusVirtualObjectTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusVirtualObject::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDBusVirtualObjectFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDBusVirtualObject) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDBusVirtualObject::timerEvent")

	if ptr.Pointer() != nil {
		C.QDBusVirtualObject_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDBusVirtualObject) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDBusVirtualObject::timerEvent")

	if ptr.Pointer() != nil {
		C.QDBusVirtualObject_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDBusVirtualObject) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDBusVirtualObject::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QDBusVirtualObject) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDBusVirtualObject::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQDBusVirtualObjectChildEvent
func callbackQDBusVirtualObjectChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusVirtualObject::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDBusVirtualObjectFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDBusVirtualObject) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDBusVirtualObject::childEvent")

	if ptr.Pointer() != nil {
		C.QDBusVirtualObject_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDBusVirtualObject) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDBusVirtualObject::childEvent")

	if ptr.Pointer() != nil {
		C.QDBusVirtualObject_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDBusVirtualObject) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDBusVirtualObject::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QDBusVirtualObject) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDBusVirtualObject::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQDBusVirtualObjectCustomEvent
func callbackQDBusVirtualObjectCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusVirtualObject::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDBusVirtualObjectFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDBusVirtualObject) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDBusVirtualObject::customEvent")

	if ptr.Pointer() != nil {
		C.QDBusVirtualObject_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDBusVirtualObject) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDBusVirtualObject::customEvent")

	if ptr.Pointer() != nil {
		C.QDBusVirtualObject_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
