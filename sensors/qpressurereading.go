package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QPressureReading struct {
	QSensorReading
}

type QPressureReading_ITF interface {
	QSensorReading_ITF
	QPressureReading_PTR() *QPressureReading
}

func PointerFromQPressureReading(ptr QPressureReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPressureReading_PTR().Pointer()
	}
	return nil
}

func NewQPressureReadingFromPointer(ptr unsafe.Pointer) *QPressureReading {
	var n = new(QPressureReading)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QPressureReading_") {
		n.SetObjectName("QPressureReading_" + qt.Identifier())
	}
	return n
}

func (ptr *QPressureReading) QPressureReading_PTR() *QPressureReading {
	return ptr
}

func (ptr *QPressureReading) Pressure() float64 {
	defer qt.Recovering("QPressureReading::pressure")

	if ptr.Pointer() != nil {
		return float64(C.QPressureReading_Pressure(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPressureReading) Temperature() float64 {
	defer qt.Recovering("QPressureReading::temperature")

	if ptr.Pointer() != nil {
		return float64(C.QPressureReading_Temperature(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPressureReading) SetPressure(pressure float64) {
	defer qt.Recovering("QPressureReading::setPressure")

	if ptr.Pointer() != nil {
		C.QPressureReading_SetPressure(ptr.Pointer(), C.double(pressure))
	}
}

func (ptr *QPressureReading) SetTemperature(temperature float64) {
	defer qt.Recovering("QPressureReading::setTemperature")

	if ptr.Pointer() != nil {
		C.QPressureReading_SetTemperature(ptr.Pointer(), C.double(temperature))
	}
}

func (ptr *QPressureReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QPressureReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QPressureReading) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QPressureReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQPressureReadingTimerEvent
func callbackQPressureReadingTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QPressureReading::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQPressureReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QPressureReading) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QPressureReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QPressureReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QPressureReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QPressureReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QPressureReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QPressureReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QPressureReading::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QPressureReading) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QPressureReading::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQPressureReadingChildEvent
func callbackQPressureReadingChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QPressureReading::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQPressureReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QPressureReading) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QPressureReading::childEvent")

	if ptr.Pointer() != nil {
		C.QPressureReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QPressureReading) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QPressureReading::childEvent")

	if ptr.Pointer() != nil {
		C.QPressureReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QPressureReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QPressureReading::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QPressureReading) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QPressureReading::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQPressureReadingCustomEvent
func callbackQPressureReadingCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QPressureReading::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQPressureReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QPressureReading) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QPressureReading::customEvent")

	if ptr.Pointer() != nil {
		C.QPressureReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QPressureReading) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QPressureReading::customEvent")

	if ptr.Pointer() != nil {
		C.QPressureReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
