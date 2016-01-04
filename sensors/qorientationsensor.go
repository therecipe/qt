package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QOrientationSensor struct {
	QSensor
}

type QOrientationSensor_ITF interface {
	QSensor_ITF
	QOrientationSensor_PTR() *QOrientationSensor
}

func PointerFromQOrientationSensor(ptr QOrientationSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOrientationSensor_PTR().Pointer()
	}
	return nil
}

func NewQOrientationSensorFromPointer(ptr unsafe.Pointer) *QOrientationSensor {
	var n = new(QOrientationSensor)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QOrientationSensor_") {
		n.SetObjectName("QOrientationSensor_" + qt.Identifier())
	}
	return n
}

func (ptr *QOrientationSensor) QOrientationSensor_PTR() *QOrientationSensor {
	return ptr
}

func (ptr *QOrientationSensor) Reading() *QOrientationReading {
	defer qt.Recovering("QOrientationSensor::reading")

	if ptr.Pointer() != nil {
		return NewQOrientationReadingFromPointer(C.QOrientationSensor_Reading(ptr.Pointer()))
	}
	return nil
}

func NewQOrientationSensor(parent core.QObject_ITF) *QOrientationSensor {
	defer qt.Recovering("QOrientationSensor::QOrientationSensor")

	return NewQOrientationSensorFromPointer(C.QOrientationSensor_NewQOrientationSensor(core.PointerFromQObject(parent)))
}

func (ptr *QOrientationSensor) DestroyQOrientationSensor() {
	defer qt.Recovering("QOrientationSensor::~QOrientationSensor")

	if ptr.Pointer() != nil {
		C.QOrientationSensor_DestroyQOrientationSensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QOrientationSensor) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QOrientationSensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QOrientationSensor) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QOrientationSensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQOrientationSensorTimerEvent
func callbackQOrientationSensorTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QOrientationSensor::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQOrientationSensorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QOrientationSensor) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QOrientationSensor::timerEvent")

	if ptr.Pointer() != nil {
		C.QOrientationSensor_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QOrientationSensor) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QOrientationSensor::timerEvent")

	if ptr.Pointer() != nil {
		C.QOrientationSensor_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QOrientationSensor) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QOrientationSensor::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QOrientationSensor) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QOrientationSensor::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQOrientationSensorChildEvent
func callbackQOrientationSensorChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QOrientationSensor::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQOrientationSensorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QOrientationSensor) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QOrientationSensor::childEvent")

	if ptr.Pointer() != nil {
		C.QOrientationSensor_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QOrientationSensor) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QOrientationSensor::childEvent")

	if ptr.Pointer() != nil {
		C.QOrientationSensor_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QOrientationSensor) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QOrientationSensor::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QOrientationSensor) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QOrientationSensor::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQOrientationSensorCustomEvent
func callbackQOrientationSensorCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QOrientationSensor::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQOrientationSensorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QOrientationSensor) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QOrientationSensor::customEvent")

	if ptr.Pointer() != nil {
		C.QOrientationSensor_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QOrientationSensor) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QOrientationSensor::customEvent")

	if ptr.Pointer() != nil {
		C.QOrientationSensor_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
