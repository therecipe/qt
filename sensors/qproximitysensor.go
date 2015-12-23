package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QProximitySensor struct {
	QSensor
}

type QProximitySensor_ITF interface {
	QSensor_ITF
	QProximitySensor_PTR() *QProximitySensor
}

func PointerFromQProximitySensor(ptr QProximitySensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QProximitySensor_PTR().Pointer()
	}
	return nil
}

func NewQProximitySensorFromPointer(ptr unsafe.Pointer) *QProximitySensor {
	var n = new(QProximitySensor)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QProximitySensor_") {
		n.SetObjectName("QProximitySensor_" + qt.Identifier())
	}
	return n
}

func (ptr *QProximitySensor) QProximitySensor_PTR() *QProximitySensor {
	return ptr
}

func (ptr *QProximitySensor) Reading() *QProximityReading {
	defer qt.Recovering("QProximitySensor::reading")

	if ptr.Pointer() != nil {
		return NewQProximityReadingFromPointer(C.QProximitySensor_Reading(ptr.Pointer()))
	}
	return nil
}

func NewQProximitySensor(parent core.QObject_ITF) *QProximitySensor {
	defer qt.Recovering("QProximitySensor::QProximitySensor")

	return NewQProximitySensorFromPointer(C.QProximitySensor_NewQProximitySensor(core.PointerFromQObject(parent)))
}

func (ptr *QProximitySensor) DestroyQProximitySensor() {
	defer qt.Recovering("QProximitySensor::~QProximitySensor")

	if ptr.Pointer() != nil {
		C.QProximitySensor_DestroyQProximitySensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QProximitySensor) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QProximitySensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QProximitySensor) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QProximitySensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQProximitySensorTimerEvent
func callbackQProximitySensorTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QProximitySensor::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QProximitySensor) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QProximitySensor::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QProximitySensor) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QProximitySensor::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQProximitySensorChildEvent
func callbackQProximitySensorChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QProximitySensor::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QProximitySensor) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QProximitySensor::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QProximitySensor) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QProximitySensor::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQProximitySensorCustomEvent
func callbackQProximitySensorCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QProximitySensor::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
