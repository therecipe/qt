package dbus

//#include "dbus.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QDBusAbstractAdaptor struct {
	core.QObject
}

type QDBusAbstractAdaptor_ITF interface {
	core.QObject_ITF
	QDBusAbstractAdaptor_PTR() *QDBusAbstractAdaptor
}

func PointerFromQDBusAbstractAdaptor(ptr QDBusAbstractAdaptor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusAbstractAdaptor_PTR().Pointer()
	}
	return nil
}

func NewQDBusAbstractAdaptorFromPointer(ptr unsafe.Pointer) *QDBusAbstractAdaptor {
	var n = new(QDBusAbstractAdaptor)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QDBusAbstractAdaptor_") {
		n.SetObjectName("QDBusAbstractAdaptor_" + qt.Identifier())
	}
	return n
}

func (ptr *QDBusAbstractAdaptor) QDBusAbstractAdaptor_PTR() *QDBusAbstractAdaptor {
	return ptr
}

func (ptr *QDBusAbstractAdaptor) DestroyQDBusAbstractAdaptor() {
	defer qt.Recovering("QDBusAbstractAdaptor::~QDBusAbstractAdaptor")

	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_DestroyQDBusAbstractAdaptor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDBusAbstractAdaptor) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDBusAbstractAdaptor::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QDBusAbstractAdaptor) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDBusAbstractAdaptor::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQDBusAbstractAdaptorTimerEvent
func callbackQDBusAbstractAdaptorTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusAbstractAdaptor::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDBusAbstractAdaptorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDBusAbstractAdaptor) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDBusAbstractAdaptor::timerEvent")

	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDBusAbstractAdaptor) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDBusAbstractAdaptor::timerEvent")

	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDBusAbstractAdaptor) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDBusAbstractAdaptor::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QDBusAbstractAdaptor) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDBusAbstractAdaptor::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQDBusAbstractAdaptorChildEvent
func callbackQDBusAbstractAdaptorChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusAbstractAdaptor::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDBusAbstractAdaptorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDBusAbstractAdaptor) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDBusAbstractAdaptor::childEvent")

	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDBusAbstractAdaptor) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDBusAbstractAdaptor::childEvent")

	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDBusAbstractAdaptor) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDBusAbstractAdaptor::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QDBusAbstractAdaptor) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDBusAbstractAdaptor::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQDBusAbstractAdaptorCustomEvent
func callbackQDBusAbstractAdaptorCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusAbstractAdaptor::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDBusAbstractAdaptorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDBusAbstractAdaptor) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDBusAbstractAdaptor::customEvent")

	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDBusAbstractAdaptor) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDBusAbstractAdaptor::customEvent")

	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
