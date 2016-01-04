package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QTiltSensor struct {
	QSensor
}

type QTiltSensor_ITF interface {
	QSensor_ITF
	QTiltSensor_PTR() *QTiltSensor
}

func PointerFromQTiltSensor(ptr QTiltSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTiltSensor_PTR().Pointer()
	}
	return nil
}

func NewQTiltSensorFromPointer(ptr unsafe.Pointer) *QTiltSensor {
	var n = new(QTiltSensor)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QTiltSensor_") {
		n.SetObjectName("QTiltSensor_" + qt.Identifier())
	}
	return n
}

func (ptr *QTiltSensor) QTiltSensor_PTR() *QTiltSensor {
	return ptr
}

func NewQTiltSensor(parent core.QObject_ITF) *QTiltSensor {
	defer qt.Recovering("QTiltSensor::QTiltSensor")

	return NewQTiltSensorFromPointer(C.QTiltSensor_NewQTiltSensor(core.PointerFromQObject(parent)))
}

func (ptr *QTiltSensor) Reading() *QTiltReading {
	defer qt.Recovering("QTiltSensor::reading")

	if ptr.Pointer() != nil {
		return NewQTiltReadingFromPointer(C.QTiltSensor_Reading(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTiltSensor) DestroyQTiltSensor() {
	defer qt.Recovering("QTiltSensor::~QTiltSensor")

	if ptr.Pointer() != nil {
		C.QTiltSensor_DestroyQTiltSensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QTiltSensor) Calibrate() {
	defer qt.Recovering("QTiltSensor::calibrate")

	if ptr.Pointer() != nil {
		C.QTiltSensor_Calibrate(ptr.Pointer())
	}
}

func (ptr *QTiltSensor) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QTiltSensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QTiltSensor) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QTiltSensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQTiltSensorTimerEvent
func callbackQTiltSensorTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTiltSensor::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQTiltSensorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QTiltSensor) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QTiltSensor::timerEvent")

	if ptr.Pointer() != nil {
		C.QTiltSensor_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QTiltSensor) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QTiltSensor::timerEvent")

	if ptr.Pointer() != nil {
		C.QTiltSensor_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QTiltSensor) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QTiltSensor::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QTiltSensor) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QTiltSensor::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQTiltSensorChildEvent
func callbackQTiltSensorChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTiltSensor::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQTiltSensorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QTiltSensor) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QTiltSensor::childEvent")

	if ptr.Pointer() != nil {
		C.QTiltSensor_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QTiltSensor) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QTiltSensor::childEvent")

	if ptr.Pointer() != nil {
		C.QTiltSensor_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QTiltSensor) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QTiltSensor::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QTiltSensor) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QTiltSensor::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQTiltSensorCustomEvent
func callbackQTiltSensorCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTiltSensor::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQTiltSensorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QTiltSensor) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QTiltSensor::customEvent")

	if ptr.Pointer() != nil {
		C.QTiltSensor_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QTiltSensor) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QTiltSensor::customEvent")

	if ptr.Pointer() != nil {
		C.QTiltSensor_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
