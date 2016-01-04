package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QLightSensor struct {
	QSensor
}

type QLightSensor_ITF interface {
	QSensor_ITF
	QLightSensor_PTR() *QLightSensor
}

func PointerFromQLightSensor(ptr QLightSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLightSensor_PTR().Pointer()
	}
	return nil
}

func NewQLightSensorFromPointer(ptr unsafe.Pointer) *QLightSensor {
	var n = new(QLightSensor)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QLightSensor_") {
		n.SetObjectName("QLightSensor_" + qt.Identifier())
	}
	return n
}

func (ptr *QLightSensor) QLightSensor_PTR() *QLightSensor {
	return ptr
}

func (ptr *QLightSensor) FieldOfView() float64 {
	defer qt.Recovering("QLightSensor::fieldOfView")

	if ptr.Pointer() != nil {
		return float64(C.QLightSensor_FieldOfView(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLightSensor) Reading() *QLightReading {
	defer qt.Recovering("QLightSensor::reading")

	if ptr.Pointer() != nil {
		return NewQLightReadingFromPointer(C.QLightSensor_Reading(ptr.Pointer()))
	}
	return nil
}

func NewQLightSensor(parent core.QObject_ITF) *QLightSensor {
	defer qt.Recovering("QLightSensor::QLightSensor")

	return NewQLightSensorFromPointer(C.QLightSensor_NewQLightSensor(core.PointerFromQObject(parent)))
}

func (ptr *QLightSensor) ConnectFieldOfViewChanged(f func(fieldOfView float64)) {
	defer qt.Recovering("connect QLightSensor::fieldOfViewChanged")

	if ptr.Pointer() != nil {
		C.QLightSensor_ConnectFieldOfViewChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "fieldOfViewChanged", f)
	}
}

func (ptr *QLightSensor) DisconnectFieldOfViewChanged() {
	defer qt.Recovering("disconnect QLightSensor::fieldOfViewChanged")

	if ptr.Pointer() != nil {
		C.QLightSensor_DisconnectFieldOfViewChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "fieldOfViewChanged")
	}
}

//export callbackQLightSensorFieldOfViewChanged
func callbackQLightSensorFieldOfViewChanged(ptr unsafe.Pointer, ptrName *C.char, fieldOfView C.double) {
	defer qt.Recovering("callback QLightSensor::fieldOfViewChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "fieldOfViewChanged"); signal != nil {
		signal.(func(float64))(float64(fieldOfView))
	}

}

func (ptr *QLightSensor) FieldOfViewChanged(fieldOfView float64) {
	defer qt.Recovering("QLightSensor::fieldOfViewChanged")

	if ptr.Pointer() != nil {
		C.QLightSensor_FieldOfViewChanged(ptr.Pointer(), C.double(fieldOfView))
	}
}

func (ptr *QLightSensor) SetFieldOfView(fieldOfView float64) {
	defer qt.Recovering("QLightSensor::setFieldOfView")

	if ptr.Pointer() != nil {
		C.QLightSensor_SetFieldOfView(ptr.Pointer(), C.double(fieldOfView))
	}
}

func (ptr *QLightSensor) DestroyQLightSensor() {
	defer qt.Recovering("QLightSensor::~QLightSensor")

	if ptr.Pointer() != nil {
		C.QLightSensor_DestroyQLightSensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QLightSensor) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QLightSensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QLightSensor) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QLightSensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQLightSensorTimerEvent
func callbackQLightSensorTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLightSensor::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQLightSensorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QLightSensor) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QLightSensor::timerEvent")

	if ptr.Pointer() != nil {
		C.QLightSensor_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QLightSensor) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QLightSensor::timerEvent")

	if ptr.Pointer() != nil {
		C.QLightSensor_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QLightSensor) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QLightSensor::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QLightSensor) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QLightSensor::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQLightSensorChildEvent
func callbackQLightSensorChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLightSensor::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQLightSensorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QLightSensor) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QLightSensor::childEvent")

	if ptr.Pointer() != nil {
		C.QLightSensor_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QLightSensor) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QLightSensor::childEvent")

	if ptr.Pointer() != nil {
		C.QLightSensor_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QLightSensor) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QLightSensor::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QLightSensor) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QLightSensor::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQLightSensorCustomEvent
func callbackQLightSensorCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLightSensor::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQLightSensorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QLightSensor) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QLightSensor::customEvent")

	if ptr.Pointer() != nil {
		C.QLightSensor_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QLightSensor) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QLightSensor::customEvent")

	if ptr.Pointer() != nil {
		C.QLightSensor_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
