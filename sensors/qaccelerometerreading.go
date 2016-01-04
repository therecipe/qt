package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAccelerometerReading struct {
	QSensorReading
}

type QAccelerometerReading_ITF interface {
	QSensorReading_ITF
	QAccelerometerReading_PTR() *QAccelerometerReading
}

func PointerFromQAccelerometerReading(ptr QAccelerometerReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccelerometerReading_PTR().Pointer()
	}
	return nil
}

func NewQAccelerometerReadingFromPointer(ptr unsafe.Pointer) *QAccelerometerReading {
	var n = new(QAccelerometerReading)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAccelerometerReading_") {
		n.SetObjectName("QAccelerometerReading_" + qt.Identifier())
	}
	return n
}

func (ptr *QAccelerometerReading) QAccelerometerReading_PTR() *QAccelerometerReading {
	return ptr
}

func (ptr *QAccelerometerReading) X() float64 {
	defer qt.Recovering("QAccelerometerReading::x")

	if ptr.Pointer() != nil {
		return float64(C.QAccelerometerReading_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccelerometerReading) Y() float64 {
	defer qt.Recovering("QAccelerometerReading::y")

	if ptr.Pointer() != nil {
		return float64(C.QAccelerometerReading_Y(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccelerometerReading) Z() float64 {
	defer qt.Recovering("QAccelerometerReading::z")

	if ptr.Pointer() != nil {
		return float64(C.QAccelerometerReading_Z(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccelerometerReading) SetX(x float64) {
	defer qt.Recovering("QAccelerometerReading::setX")

	if ptr.Pointer() != nil {
		C.QAccelerometerReading_SetX(ptr.Pointer(), C.double(x))
	}
}

func (ptr *QAccelerometerReading) SetY(y float64) {
	defer qt.Recovering("QAccelerometerReading::setY")

	if ptr.Pointer() != nil {
		C.QAccelerometerReading_SetY(ptr.Pointer(), C.double(y))
	}
}

func (ptr *QAccelerometerReading) SetZ(z float64) {
	defer qt.Recovering("QAccelerometerReading::setZ")

	if ptr.Pointer() != nil {
		C.QAccelerometerReading_SetZ(ptr.Pointer(), C.double(z))
	}
}

func (ptr *QAccelerometerReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAccelerometerReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAccelerometerReading) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAccelerometerReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAccelerometerReadingTimerEvent
func callbackQAccelerometerReadingTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAccelerometerReading::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAccelerometerReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAccelerometerReading) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAccelerometerReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QAccelerometerReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAccelerometerReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAccelerometerReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QAccelerometerReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAccelerometerReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAccelerometerReading::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAccelerometerReading) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAccelerometerReading::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAccelerometerReadingChildEvent
func callbackQAccelerometerReadingChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAccelerometerReading::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAccelerometerReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAccelerometerReading) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAccelerometerReading::childEvent")

	if ptr.Pointer() != nil {
		C.QAccelerometerReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAccelerometerReading) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAccelerometerReading::childEvent")

	if ptr.Pointer() != nil {
		C.QAccelerometerReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAccelerometerReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAccelerometerReading::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAccelerometerReading) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAccelerometerReading::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAccelerometerReadingCustomEvent
func callbackQAccelerometerReadingCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAccelerometerReading::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAccelerometerReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAccelerometerReading) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAccelerometerReading::customEvent")

	if ptr.Pointer() != nil {
		C.QAccelerometerReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAccelerometerReading) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAccelerometerReading::customEvent")

	if ptr.Pointer() != nil {
		C.QAccelerometerReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
