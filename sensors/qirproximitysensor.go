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
func callbackQIRProximitySensorTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QIRProximitySensor::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

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
func callbackQIRProximitySensorChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QIRProximitySensor::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

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
func callbackQIRProximitySensorCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QIRProximitySensor::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
