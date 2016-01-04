package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QLightReading struct {
	QSensorReading
}

type QLightReading_ITF interface {
	QSensorReading_ITF
	QLightReading_PTR() *QLightReading
}

func PointerFromQLightReading(ptr QLightReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLightReading_PTR().Pointer()
	}
	return nil
}

func NewQLightReadingFromPointer(ptr unsafe.Pointer) *QLightReading {
	var n = new(QLightReading)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QLightReading_") {
		n.SetObjectName("QLightReading_" + qt.Identifier())
	}
	return n
}

func (ptr *QLightReading) QLightReading_PTR() *QLightReading {
	return ptr
}

func (ptr *QLightReading) Lux() float64 {
	defer qt.Recovering("QLightReading::lux")

	if ptr.Pointer() != nil {
		return float64(C.QLightReading_Lux(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLightReading) SetLux(lux float64) {
	defer qt.Recovering("QLightReading::setLux")

	if ptr.Pointer() != nil {
		C.QLightReading_SetLux(ptr.Pointer(), C.double(lux))
	}
}

func (ptr *QLightReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QLightReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QLightReading) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QLightReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQLightReadingTimerEvent
func callbackQLightReadingTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLightReading::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQLightReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QLightReading) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QLightReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QLightReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QLightReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QLightReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QLightReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QLightReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QLightReading::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QLightReading) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QLightReading::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQLightReadingChildEvent
func callbackQLightReadingChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLightReading::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQLightReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QLightReading) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QLightReading::childEvent")

	if ptr.Pointer() != nil {
		C.QLightReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QLightReading) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QLightReading::childEvent")

	if ptr.Pointer() != nil {
		C.QLightReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QLightReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QLightReading::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QLightReading) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QLightReading::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQLightReadingCustomEvent
func callbackQLightReadingCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLightReading::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQLightReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QLightReading) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QLightReading::customEvent")

	if ptr.Pointer() != nil {
		C.QLightReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QLightReading) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QLightReading::customEvent")

	if ptr.Pointer() != nil {
		C.QLightReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
