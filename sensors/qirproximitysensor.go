package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QIRProximitySensor struct {
	QSensor
}

type QIRProximitySensor_ITF interface {
	QSensor_ITF
	QIRProximitySensor_PTR() *QIRProximitySensor
}

func PointerFromQIRProximitySensor(ptr QIRProximitySensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QIRProximitySensor_PTR().Pointer()
	}
	return nil
}

func NewQIRProximitySensorFromPointer(ptr unsafe.Pointer) *QIRProximitySensor {
	var n = new(QIRProximitySensor)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QIRProximitySensor_") {
		n.SetObjectName("QIRProximitySensor_" + qt.Identifier())
	}
	return n
}

func (ptr *QIRProximitySensor) QIRProximitySensor_PTR() *QIRProximitySensor {
	return ptr
}

func (ptr *QIRProximitySensor) Reading() *QIRProximityReading {
	defer qt.Recovering("QIRProximitySensor::reading")

	if ptr.Pointer() != nil {
		return NewQIRProximityReadingFromPointer(C.QIRProximitySensor_Reading(ptr.Pointer()))
	}
	return nil
}

func NewQIRProximitySensor(parent core.QObject_ITF) *QIRProximitySensor {
	defer qt.Recovering("QIRProximitySensor::QIRProximitySensor")

	return NewQIRProximitySensorFromPointer(C.QIRProximitySensor_NewQIRProximitySensor(core.PointerFromQObject(parent)))
}

func (ptr *QIRProximitySensor) DestroyQIRProximitySensor() {
	defer qt.Recovering("QIRProximitySensor::~QIRProximitySensor")

	if ptr.Pointer() != nil {
		C.QIRProximitySensor_DestroyQIRProximitySensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QIRProximitySensor) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QIRProximitySensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QIRProximitySensor) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QIRProximitySensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQIRProximitySensorTimerEvent
func callbackQIRProximitySensorTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QIRProximitySensor::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQIRProximitySensorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QIRProximitySensor) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QIRProximitySensor::timerEvent")

	if ptr.Pointer() != nil {
		C.QIRProximitySensor_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QIRProximitySensor) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QIRProximitySensor::timerEvent")

	if ptr.Pointer() != nil {
		C.QIRProximitySensor_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QIRProximitySensor) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QIRProximitySensor::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QIRProximitySensor) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QIRProximitySensor::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQIRProximitySensorChildEvent
func callbackQIRProximitySensorChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QIRProximitySensor::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQIRProximitySensorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QIRProximitySensor) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QIRProximitySensor::childEvent")

	if ptr.Pointer() != nil {
		C.QIRProximitySensor_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QIRProximitySensor) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QIRProximitySensor::childEvent")

	if ptr.Pointer() != nil {
		C.QIRProximitySensor_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QIRProximitySensor) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QIRProximitySensor::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QIRProximitySensor) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QIRProximitySensor::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQIRProximitySensorCustomEvent
func callbackQIRProximitySensorCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QIRProximitySensor::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQIRProximitySensorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QIRProximitySensor) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QIRProximitySensor::customEvent")

	if ptr.Pointer() != nil {
		C.QIRProximitySensor_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QIRProximitySensor) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QIRProximitySensor::customEvent")

	if ptr.Pointer() != nil {
		C.QIRProximitySensor_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
