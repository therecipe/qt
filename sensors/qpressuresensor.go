package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QPressureSensor struct {
	QSensor
}

type QPressureSensor_ITF interface {
	QSensor_ITF
	QPressureSensor_PTR() *QPressureSensor
}

func PointerFromQPressureSensor(ptr QPressureSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPressureSensor_PTR().Pointer()
	}
	return nil
}

func NewQPressureSensorFromPointer(ptr unsafe.Pointer) *QPressureSensor {
	var n = new(QPressureSensor)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QPressureSensor_") {
		n.SetObjectName("QPressureSensor_" + qt.Identifier())
	}
	return n
}

func (ptr *QPressureSensor) QPressureSensor_PTR() *QPressureSensor {
	return ptr
}

func (ptr *QPressureSensor) Reading() *QPressureReading {
	defer qt.Recovering("QPressureSensor::reading")

	if ptr.Pointer() != nil {
		return NewQPressureReadingFromPointer(C.QPressureSensor_Reading(ptr.Pointer()))
	}
	return nil
}

func NewQPressureSensor(parent core.QObject_ITF) *QPressureSensor {
	defer qt.Recovering("QPressureSensor::QPressureSensor")

	return NewQPressureSensorFromPointer(C.QPressureSensor_NewQPressureSensor(core.PointerFromQObject(parent)))
}

func (ptr *QPressureSensor) DestroyQPressureSensor() {
	defer qt.Recovering("QPressureSensor::~QPressureSensor")

	if ptr.Pointer() != nil {
		C.QPressureSensor_DestroyQPressureSensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QPressureSensor) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QPressureSensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QPressureSensor) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QPressureSensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQPressureSensorTimerEvent
func callbackQPressureSensorTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QPressureSensor::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQPressureSensorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QPressureSensor) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QPressureSensor::timerEvent")

	if ptr.Pointer() != nil {
		C.QPressureSensor_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QPressureSensor) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QPressureSensor::timerEvent")

	if ptr.Pointer() != nil {
		C.QPressureSensor_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QPressureSensor) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QPressureSensor::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QPressureSensor) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QPressureSensor::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQPressureSensorChildEvent
func callbackQPressureSensorChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QPressureSensor::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQPressureSensorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QPressureSensor) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QPressureSensor::childEvent")

	if ptr.Pointer() != nil {
		C.QPressureSensor_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QPressureSensor) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QPressureSensor::childEvent")

	if ptr.Pointer() != nil {
		C.QPressureSensor_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QPressureSensor) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QPressureSensor::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QPressureSensor) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QPressureSensor::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQPressureSensorCustomEvent
func callbackQPressureSensorCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QPressureSensor::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQPressureSensorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QPressureSensor) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QPressureSensor::customEvent")

	if ptr.Pointer() != nil {
		C.QPressureSensor_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QPressureSensor) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QPressureSensor::customEvent")

	if ptr.Pointer() != nil {
		C.QPressureSensor_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
