package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAmbientLightSensor struct {
	QSensor
}

type QAmbientLightSensor_ITF interface {
	QSensor_ITF
	QAmbientLightSensor_PTR() *QAmbientLightSensor
}

func PointerFromQAmbientLightSensor(ptr QAmbientLightSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAmbientLightSensor_PTR().Pointer()
	}
	return nil
}

func NewQAmbientLightSensorFromPointer(ptr unsafe.Pointer) *QAmbientLightSensor {
	var n = new(QAmbientLightSensor)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAmbientLightSensor_") {
		n.SetObjectName("QAmbientLightSensor_" + qt.Identifier())
	}
	return n
}

func (ptr *QAmbientLightSensor) QAmbientLightSensor_PTR() *QAmbientLightSensor {
	return ptr
}

func (ptr *QAmbientLightSensor) Reading() *QAmbientLightReading {
	defer qt.Recovering("QAmbientLightSensor::reading")

	if ptr.Pointer() != nil {
		return NewQAmbientLightReadingFromPointer(C.QAmbientLightSensor_Reading(ptr.Pointer()))
	}
	return nil
}

func NewQAmbientLightSensor(parent core.QObject_ITF) *QAmbientLightSensor {
	defer qt.Recovering("QAmbientLightSensor::QAmbientLightSensor")

	return NewQAmbientLightSensorFromPointer(C.QAmbientLightSensor_NewQAmbientLightSensor(core.PointerFromQObject(parent)))
}

func (ptr *QAmbientLightSensor) DestroyQAmbientLightSensor() {
	defer qt.Recovering("QAmbientLightSensor::~QAmbientLightSensor")

	if ptr.Pointer() != nil {
		C.QAmbientLightSensor_DestroyQAmbientLightSensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAmbientLightSensor) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAmbientLightSensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAmbientLightSensor) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAmbientLightSensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAmbientLightSensorTimerEvent
func callbackQAmbientLightSensorTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAmbientLightSensor::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAmbientLightSensorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAmbientLightSensor) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAmbientLightSensor::timerEvent")

	if ptr.Pointer() != nil {
		C.QAmbientLightSensor_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAmbientLightSensor) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAmbientLightSensor::timerEvent")

	if ptr.Pointer() != nil {
		C.QAmbientLightSensor_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAmbientLightSensor) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAmbientLightSensor::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAmbientLightSensor) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAmbientLightSensor::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAmbientLightSensorChildEvent
func callbackQAmbientLightSensorChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAmbientLightSensor::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAmbientLightSensorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAmbientLightSensor) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAmbientLightSensor::childEvent")

	if ptr.Pointer() != nil {
		C.QAmbientLightSensor_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAmbientLightSensor) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAmbientLightSensor::childEvent")

	if ptr.Pointer() != nil {
		C.QAmbientLightSensor_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAmbientLightSensor) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAmbientLightSensor::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAmbientLightSensor) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAmbientLightSensor::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAmbientLightSensorCustomEvent
func callbackQAmbientLightSensorCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAmbientLightSensor::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAmbientLightSensorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAmbientLightSensor) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAmbientLightSensor::customEvent")

	if ptr.Pointer() != nil {
		C.QAmbientLightSensor_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAmbientLightSensor) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAmbientLightSensor::customEvent")

	if ptr.Pointer() != nil {
		C.QAmbientLightSensor_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
