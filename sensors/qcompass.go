package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QCompass struct {
	QSensor
}

type QCompass_ITF interface {
	QSensor_ITF
	QCompass_PTR() *QCompass
}

func PointerFromQCompass(ptr QCompass_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCompass_PTR().Pointer()
	}
	return nil
}

func NewQCompassFromPointer(ptr unsafe.Pointer) *QCompass {
	var n = new(QCompass)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QCompass_") {
		n.SetObjectName("QCompass_" + qt.Identifier())
	}
	return n
}

func (ptr *QCompass) QCompass_PTR() *QCompass {
	return ptr
}

func (ptr *QCompass) Reading() *QCompassReading {
	defer qt.Recovering("QCompass::reading")

	if ptr.Pointer() != nil {
		return NewQCompassReadingFromPointer(C.QCompass_Reading(ptr.Pointer()))
	}
	return nil
}

func NewQCompass(parent core.QObject_ITF) *QCompass {
	defer qt.Recovering("QCompass::QCompass")

	return NewQCompassFromPointer(C.QCompass_NewQCompass(core.PointerFromQObject(parent)))
}

func (ptr *QCompass) DestroyQCompass() {
	defer qt.Recovering("QCompass::~QCompass")

	if ptr.Pointer() != nil {
		C.QCompass_DestroyQCompass(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCompass) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QCompass::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCompass) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCompass::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQCompassTimerEvent
func callbackQCompassTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCompass::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQCompassFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QCompass) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCompass::timerEvent")

	if ptr.Pointer() != nil {
		C.QCompass_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCompass) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCompass::timerEvent")

	if ptr.Pointer() != nil {
		C.QCompass_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCompass) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCompass::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCompass) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCompass::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQCompassChildEvent
func callbackQCompassChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCompass::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCompassFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCompass) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCompass::childEvent")

	if ptr.Pointer() != nil {
		C.QCompass_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCompass) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCompass::childEvent")

	if ptr.Pointer() != nil {
		C.QCompass_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCompass) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCompass::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCompass) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCompass::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQCompassCustomEvent
func callbackQCompassCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCompass::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCompassFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCompass) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCompass::customEvent")

	if ptr.Pointer() != nil {
		C.QCompass_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCompass) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCompass::customEvent")

	if ptr.Pointer() != nil {
		C.QCompass_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
