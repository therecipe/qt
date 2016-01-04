package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QGyroscope struct {
	QSensor
}

type QGyroscope_ITF interface {
	QSensor_ITF
	QGyroscope_PTR() *QGyroscope
}

func PointerFromQGyroscope(ptr QGyroscope_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGyroscope_PTR().Pointer()
	}
	return nil
}

func NewQGyroscopeFromPointer(ptr unsafe.Pointer) *QGyroscope {
	var n = new(QGyroscope)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QGyroscope_") {
		n.SetObjectName("QGyroscope_" + qt.Identifier())
	}
	return n
}

func (ptr *QGyroscope) QGyroscope_PTR() *QGyroscope {
	return ptr
}

func (ptr *QGyroscope) Reading() *QGyroscopeReading {
	defer qt.Recovering("QGyroscope::reading")

	if ptr.Pointer() != nil {
		return NewQGyroscopeReadingFromPointer(C.QGyroscope_Reading(ptr.Pointer()))
	}
	return nil
}

func NewQGyroscope(parent core.QObject_ITF) *QGyroscope {
	defer qt.Recovering("QGyroscope::QGyroscope")

	return NewQGyroscopeFromPointer(C.QGyroscope_NewQGyroscope(core.PointerFromQObject(parent)))
}

func (ptr *QGyroscope) DestroyQGyroscope() {
	defer qt.Recovering("QGyroscope::~QGyroscope")

	if ptr.Pointer() != nil {
		C.QGyroscope_DestroyQGyroscope(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGyroscope) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QGyroscope::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QGyroscope) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QGyroscope::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQGyroscopeTimerEvent
func callbackQGyroscopeTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGyroscope::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGyroscopeFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGyroscope) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGyroscope::timerEvent")

	if ptr.Pointer() != nil {
		C.QGyroscope_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGyroscope) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGyroscope::timerEvent")

	if ptr.Pointer() != nil {
		C.QGyroscope_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGyroscope) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QGyroscope::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QGyroscope) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QGyroscope::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQGyroscopeChildEvent
func callbackQGyroscopeChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGyroscope::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQGyroscopeFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QGyroscope) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QGyroscope::childEvent")

	if ptr.Pointer() != nil {
		C.QGyroscope_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGyroscope) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QGyroscope::childEvent")

	if ptr.Pointer() != nil {
		C.QGyroscope_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGyroscope) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGyroscope::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QGyroscope) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QGyroscope::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQGyroscopeCustomEvent
func callbackQGyroscopeCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGyroscope::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGyroscopeFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGyroscope) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QGyroscope::customEvent")

	if ptr.Pointer() != nil {
		C.QGyroscope_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGyroscope) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QGyroscope::customEvent")

	if ptr.Pointer() != nil {
		C.QGyroscope_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
