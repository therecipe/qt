package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QTapSensor struct {
	QSensor
}

type QTapSensor_ITF interface {
	QSensor_ITF
	QTapSensor_PTR() *QTapSensor
}

func PointerFromQTapSensor(ptr QTapSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTapSensor_PTR().Pointer()
	}
	return nil
}

func NewQTapSensorFromPointer(ptr unsafe.Pointer) *QTapSensor {
	var n = new(QTapSensor)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QTapSensor_") {
		n.SetObjectName("QTapSensor_" + qt.Identifier())
	}
	return n
}

func (ptr *QTapSensor) QTapSensor_PTR() *QTapSensor {
	return ptr
}

func (ptr *QTapSensor) Reading() *QTapReading {
	defer qt.Recovering("QTapSensor::reading")

	if ptr.Pointer() != nil {
		return NewQTapReadingFromPointer(C.QTapSensor_Reading(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTapSensor) ReturnDoubleTapEvents() bool {
	defer qt.Recovering("QTapSensor::returnDoubleTapEvents")

	if ptr.Pointer() != nil {
		return C.QTapSensor_ReturnDoubleTapEvents(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTapSensor) SetReturnDoubleTapEvents(returnDoubleTapEvents bool) {
	defer qt.Recovering("QTapSensor::setReturnDoubleTapEvents")

	if ptr.Pointer() != nil {
		C.QTapSensor_SetReturnDoubleTapEvents(ptr.Pointer(), C.int(qt.GoBoolToInt(returnDoubleTapEvents)))
	}
}

func NewQTapSensor(parent core.QObject_ITF) *QTapSensor {
	defer qt.Recovering("QTapSensor::QTapSensor")

	return NewQTapSensorFromPointer(C.QTapSensor_NewQTapSensor(core.PointerFromQObject(parent)))
}

func (ptr *QTapSensor) ConnectReturnDoubleTapEventsChanged(f func(returnDoubleTapEvents bool)) {
	defer qt.Recovering("connect QTapSensor::returnDoubleTapEventsChanged")

	if ptr.Pointer() != nil {
		C.QTapSensor_ConnectReturnDoubleTapEventsChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "returnDoubleTapEventsChanged", f)
	}
}

func (ptr *QTapSensor) DisconnectReturnDoubleTapEventsChanged() {
	defer qt.Recovering("disconnect QTapSensor::returnDoubleTapEventsChanged")

	if ptr.Pointer() != nil {
		C.QTapSensor_DisconnectReturnDoubleTapEventsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "returnDoubleTapEventsChanged")
	}
}

//export callbackQTapSensorReturnDoubleTapEventsChanged
func callbackQTapSensorReturnDoubleTapEventsChanged(ptr unsafe.Pointer, ptrName *C.char, returnDoubleTapEvents C.int) {
	defer qt.Recovering("callback QTapSensor::returnDoubleTapEventsChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "returnDoubleTapEventsChanged"); signal != nil {
		signal.(func(bool))(int(returnDoubleTapEvents) != 0)
	}

}

func (ptr *QTapSensor) ReturnDoubleTapEventsChanged(returnDoubleTapEvents bool) {
	defer qt.Recovering("QTapSensor::returnDoubleTapEventsChanged")

	if ptr.Pointer() != nil {
		C.QTapSensor_ReturnDoubleTapEventsChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(returnDoubleTapEvents)))
	}
}

func (ptr *QTapSensor) DestroyQTapSensor() {
	defer qt.Recovering("QTapSensor::~QTapSensor")

	if ptr.Pointer() != nil {
		C.QTapSensor_DestroyQTapSensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QTapSensor) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QTapSensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QTapSensor) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QTapSensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQTapSensorTimerEvent
func callbackQTapSensorTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTapSensor::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQTapSensorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QTapSensor) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QTapSensor::timerEvent")

	if ptr.Pointer() != nil {
		C.QTapSensor_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QTapSensor) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QTapSensor::timerEvent")

	if ptr.Pointer() != nil {
		C.QTapSensor_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QTapSensor) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QTapSensor::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QTapSensor) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QTapSensor::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQTapSensorChildEvent
func callbackQTapSensorChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTapSensor::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQTapSensorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QTapSensor) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QTapSensor::childEvent")

	if ptr.Pointer() != nil {
		C.QTapSensor_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QTapSensor) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QTapSensor::childEvent")

	if ptr.Pointer() != nil {
		C.QTapSensor_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QTapSensor) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QTapSensor::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QTapSensor) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QTapSensor::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQTapSensorCustomEvent
func callbackQTapSensorCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTapSensor::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQTapSensorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QTapSensor) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QTapSensor::customEvent")

	if ptr.Pointer() != nil {
		C.QTapSensor_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QTapSensor) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QTapSensor::customEvent")

	if ptr.Pointer() != nil {
		C.QTapSensor_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
