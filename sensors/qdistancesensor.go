package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QDistanceSensor struct {
	QSensor
}

type QDistanceSensor_ITF interface {
	QSensor_ITF
	QDistanceSensor_PTR() *QDistanceSensor
}

func PointerFromQDistanceSensor(ptr QDistanceSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDistanceSensor_PTR().Pointer()
	}
	return nil
}

func NewQDistanceSensorFromPointer(ptr unsafe.Pointer) *QDistanceSensor {
	var n = new(QDistanceSensor)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QDistanceSensor_") {
		n.SetObjectName("QDistanceSensor_" + qt.Identifier())
	}
	return n
}

func (ptr *QDistanceSensor) QDistanceSensor_PTR() *QDistanceSensor {
	return ptr
}

func (ptr *QDistanceSensor) Reading() *QDistanceReading {
	defer qt.Recovering("QDistanceSensor::reading")

	if ptr.Pointer() != nil {
		return NewQDistanceReadingFromPointer(C.QDistanceSensor_Reading(ptr.Pointer()))
	}
	return nil
}

func NewQDistanceSensor(parent core.QObject_ITF) *QDistanceSensor {
	defer qt.Recovering("QDistanceSensor::QDistanceSensor")

	return NewQDistanceSensorFromPointer(C.QDistanceSensor_NewQDistanceSensor(core.PointerFromQObject(parent)))
}

func (ptr *QDistanceSensor) DestroyQDistanceSensor() {
	defer qt.Recovering("QDistanceSensor::~QDistanceSensor")

	if ptr.Pointer() != nil {
		C.QDistanceSensor_DestroyQDistanceSensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDistanceSensor) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDistanceSensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QDistanceSensor) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDistanceSensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQDistanceSensorTimerEvent
func callbackQDistanceSensorTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDistanceSensor::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDistanceSensorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDistanceSensor) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDistanceSensor::timerEvent")

	if ptr.Pointer() != nil {
		C.QDistanceSensor_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDistanceSensor) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDistanceSensor::timerEvent")

	if ptr.Pointer() != nil {
		C.QDistanceSensor_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDistanceSensor) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDistanceSensor::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QDistanceSensor) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDistanceSensor::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQDistanceSensorChildEvent
func callbackQDistanceSensorChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDistanceSensor::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDistanceSensorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDistanceSensor) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDistanceSensor::childEvent")

	if ptr.Pointer() != nil {
		C.QDistanceSensor_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDistanceSensor) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDistanceSensor::childEvent")

	if ptr.Pointer() != nil {
		C.QDistanceSensor_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDistanceSensor) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDistanceSensor::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QDistanceSensor) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDistanceSensor::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQDistanceSensorCustomEvent
func callbackQDistanceSensorCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDistanceSensor::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDistanceSensorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDistanceSensor) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDistanceSensor::customEvent")

	if ptr.Pointer() != nil {
		C.QDistanceSensor_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDistanceSensor) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDistanceSensor::customEvent")

	if ptr.Pointer() != nil {
		C.QDistanceSensor_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
