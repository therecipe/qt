package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAmbientTemperatureSensor struct {
	QSensor
}

type QAmbientTemperatureSensor_ITF interface {
	QSensor_ITF
	QAmbientTemperatureSensor_PTR() *QAmbientTemperatureSensor
}

func PointerFromQAmbientTemperatureSensor(ptr QAmbientTemperatureSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAmbientTemperatureSensor_PTR().Pointer()
	}
	return nil
}

func NewQAmbientTemperatureSensorFromPointer(ptr unsafe.Pointer) *QAmbientTemperatureSensor {
	var n = new(QAmbientTemperatureSensor)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAmbientTemperatureSensor_") {
		n.SetObjectName("QAmbientTemperatureSensor_" + qt.Identifier())
	}
	return n
}

func (ptr *QAmbientTemperatureSensor) QAmbientTemperatureSensor_PTR() *QAmbientTemperatureSensor {
	return ptr
}

func (ptr *QAmbientTemperatureSensor) Reading() *QAmbientTemperatureReading {
	defer qt.Recovering("QAmbientTemperatureSensor::reading")

	if ptr.Pointer() != nil {
		return NewQAmbientTemperatureReadingFromPointer(C.QAmbientTemperatureSensor_Reading(ptr.Pointer()))
	}
	return nil
}

func NewQAmbientTemperatureSensor(parent core.QObject_ITF) *QAmbientTemperatureSensor {
	defer qt.Recovering("QAmbientTemperatureSensor::QAmbientTemperatureSensor")

	return NewQAmbientTemperatureSensorFromPointer(C.QAmbientTemperatureSensor_NewQAmbientTemperatureSensor(core.PointerFromQObject(parent)))
}

func (ptr *QAmbientTemperatureSensor) DestroyQAmbientTemperatureSensor() {
	defer qt.Recovering("QAmbientTemperatureSensor::~QAmbientTemperatureSensor")

	if ptr.Pointer() != nil {
		C.QAmbientTemperatureSensor_DestroyQAmbientTemperatureSensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAmbientTemperatureSensor) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAmbientTemperatureSensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAmbientTemperatureSensor) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAmbientTemperatureSensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAmbientTemperatureSensorTimerEvent
func callbackQAmbientTemperatureSensorTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAmbientTemperatureSensor::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAmbientTemperatureSensorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAmbientTemperatureSensor) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAmbientTemperatureSensor::timerEvent")

	if ptr.Pointer() != nil {
		C.QAmbientTemperatureSensor_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAmbientTemperatureSensor) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAmbientTemperatureSensor::timerEvent")

	if ptr.Pointer() != nil {
		C.QAmbientTemperatureSensor_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAmbientTemperatureSensor) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAmbientTemperatureSensor::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAmbientTemperatureSensor) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAmbientTemperatureSensor::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAmbientTemperatureSensorChildEvent
func callbackQAmbientTemperatureSensorChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAmbientTemperatureSensor::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAmbientTemperatureSensorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAmbientTemperatureSensor) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAmbientTemperatureSensor::childEvent")

	if ptr.Pointer() != nil {
		C.QAmbientTemperatureSensor_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAmbientTemperatureSensor) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAmbientTemperatureSensor::childEvent")

	if ptr.Pointer() != nil {
		C.QAmbientTemperatureSensor_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAmbientTemperatureSensor) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAmbientTemperatureSensor::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAmbientTemperatureSensor) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAmbientTemperatureSensor::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAmbientTemperatureSensorCustomEvent
func callbackQAmbientTemperatureSensorCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAmbientTemperatureSensor::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAmbientTemperatureSensorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAmbientTemperatureSensor) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAmbientTemperatureSensor::customEvent")

	if ptr.Pointer() != nil {
		C.QAmbientTemperatureSensor_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAmbientTemperatureSensor) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAmbientTemperatureSensor::customEvent")

	if ptr.Pointer() != nil {
		C.QAmbientTemperatureSensor_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
