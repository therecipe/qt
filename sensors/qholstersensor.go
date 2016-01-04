package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QHolsterSensor struct {
	QSensor
}

type QHolsterSensor_ITF interface {
	QSensor_ITF
	QHolsterSensor_PTR() *QHolsterSensor
}

func PointerFromQHolsterSensor(ptr QHolsterSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHolsterSensor_PTR().Pointer()
	}
	return nil
}

func NewQHolsterSensorFromPointer(ptr unsafe.Pointer) *QHolsterSensor {
	var n = new(QHolsterSensor)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QHolsterSensor_") {
		n.SetObjectName("QHolsterSensor_" + qt.Identifier())
	}
	return n
}

func (ptr *QHolsterSensor) QHolsterSensor_PTR() *QHolsterSensor {
	return ptr
}

func (ptr *QHolsterSensor) Reading() *QHolsterReading {
	defer qt.Recovering("QHolsterSensor::reading")

	if ptr.Pointer() != nil {
		return NewQHolsterReadingFromPointer(C.QHolsterSensor_Reading(ptr.Pointer()))
	}
	return nil
}

func NewQHolsterSensor(parent core.QObject_ITF) *QHolsterSensor {
	defer qt.Recovering("QHolsterSensor::QHolsterSensor")

	return NewQHolsterSensorFromPointer(C.QHolsterSensor_NewQHolsterSensor(core.PointerFromQObject(parent)))
}

func (ptr *QHolsterSensor) DestroyQHolsterSensor() {
	defer qt.Recovering("QHolsterSensor::~QHolsterSensor")

	if ptr.Pointer() != nil {
		C.QHolsterSensor_DestroyQHolsterSensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QHolsterSensor) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QHolsterSensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QHolsterSensor) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QHolsterSensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQHolsterSensorTimerEvent
func callbackQHolsterSensorTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHolsterSensor::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHolsterSensorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHolsterSensor) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHolsterSensor::timerEvent")

	if ptr.Pointer() != nil {
		C.QHolsterSensor_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHolsterSensor) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHolsterSensor::timerEvent")

	if ptr.Pointer() != nil {
		C.QHolsterSensor_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHolsterSensor) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QHolsterSensor::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QHolsterSensor) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QHolsterSensor::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQHolsterSensorChildEvent
func callbackQHolsterSensorChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHolsterSensor::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHolsterSensorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHolsterSensor) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHolsterSensor::childEvent")

	if ptr.Pointer() != nil {
		C.QHolsterSensor_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHolsterSensor) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHolsterSensor::childEvent")

	if ptr.Pointer() != nil {
		C.QHolsterSensor_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHolsterSensor) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHolsterSensor::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QHolsterSensor) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QHolsterSensor::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQHolsterSensorCustomEvent
func callbackQHolsterSensorCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHolsterSensor::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHolsterSensorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHolsterSensor) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHolsterSensor::customEvent")

	if ptr.Pointer() != nil {
		C.QHolsterSensor_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHolsterSensor) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHolsterSensor::customEvent")

	if ptr.Pointer() != nil {
		C.QHolsterSensor_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
