package dbus

//#include "dbus.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QDBusAbstractInterface struct {
	core.QObject
}

type QDBusAbstractInterface_ITF interface {
	core.QObject_ITF
	QDBusAbstractInterface_PTR() *QDBusAbstractInterface
}

func PointerFromQDBusAbstractInterface(ptr QDBusAbstractInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusAbstractInterface_PTR().Pointer()
	}
	return nil
}

func NewQDBusAbstractInterfaceFromPointer(ptr unsafe.Pointer) *QDBusAbstractInterface {
	var n = new(QDBusAbstractInterface)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QDBusAbstractInterface_") {
		n.SetObjectName("QDBusAbstractInterface_" + qt.Identifier())
	}
	return n
}

func (ptr *QDBusAbstractInterface) QDBusAbstractInterface_PTR() *QDBusAbstractInterface {
	return ptr
}

func (ptr *QDBusAbstractInterface) Interface() string {
	defer qt.Recovering("QDBusAbstractInterface::interface")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusAbstractInterface_Interface(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusAbstractInterface) IsValid() bool {
	defer qt.Recovering("QDBusAbstractInterface::isValid")

	if ptr.Pointer() != nil {
		return C.QDBusAbstractInterface_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusAbstractInterface) Path() string {
	defer qt.Recovering("QDBusAbstractInterface::path")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusAbstractInterface_Path(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusAbstractInterface) Service() string {
	defer qt.Recovering("QDBusAbstractInterface::service")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusAbstractInterface_Service(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusAbstractInterface) SetTimeout(timeout int) {
	defer qt.Recovering("QDBusAbstractInterface::setTimeout")

	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_SetTimeout(ptr.Pointer(), C.int(timeout))
	}
}

func (ptr *QDBusAbstractInterface) Timeout() int {
	defer qt.Recovering("QDBusAbstractInterface::timeout")

	if ptr.Pointer() != nil {
		return int(C.QDBusAbstractInterface_Timeout(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDBusAbstractInterface) DestroyQDBusAbstractInterface() {
	defer qt.Recovering("QDBusAbstractInterface::~QDBusAbstractInterface")

	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_DestroyQDBusAbstractInterface(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDBusAbstractInterface) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDBusAbstractInterface::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QDBusAbstractInterface) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDBusAbstractInterface::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQDBusAbstractInterfaceTimerEvent
func callbackQDBusAbstractInterfaceTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusAbstractInterface::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDBusAbstractInterfaceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDBusAbstractInterface) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDBusAbstractInterface::timerEvent")

	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDBusAbstractInterface) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDBusAbstractInterface::timerEvent")

	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDBusAbstractInterface) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDBusAbstractInterface::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QDBusAbstractInterface) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDBusAbstractInterface::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQDBusAbstractInterfaceChildEvent
func callbackQDBusAbstractInterfaceChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusAbstractInterface::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDBusAbstractInterfaceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDBusAbstractInterface) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDBusAbstractInterface::childEvent")

	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDBusAbstractInterface) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDBusAbstractInterface::childEvent")

	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDBusAbstractInterface) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDBusAbstractInterface::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QDBusAbstractInterface) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDBusAbstractInterface::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQDBusAbstractInterfaceCustomEvent
func callbackQDBusAbstractInterfaceCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusAbstractInterface::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDBusAbstractInterfaceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDBusAbstractInterface) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDBusAbstractInterface::customEvent")

	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDBusAbstractInterface) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDBusAbstractInterface::customEvent")

	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
