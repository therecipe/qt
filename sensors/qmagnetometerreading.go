package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QMagnetometerReading struct {
	QSensorReading
}

type QMagnetometerReading_ITF interface {
	QSensorReading_ITF
	QMagnetometerReading_PTR() *QMagnetometerReading
}

func PointerFromQMagnetometerReading(ptr QMagnetometerReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMagnetometerReading_PTR().Pointer()
	}
	return nil
}

func NewQMagnetometerReadingFromPointer(ptr unsafe.Pointer) *QMagnetometerReading {
	var n = new(QMagnetometerReading)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMagnetometerReading_") {
		n.SetObjectName("QMagnetometerReading_" + qt.Identifier())
	}
	return n
}

func (ptr *QMagnetometerReading) QMagnetometerReading_PTR() *QMagnetometerReading {
	return ptr
}

func (ptr *QMagnetometerReading) CalibrationLevel() float64 {
	defer qt.Recovering("QMagnetometerReading::calibrationLevel")

	if ptr.Pointer() != nil {
		return float64(C.QMagnetometerReading_CalibrationLevel(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMagnetometerReading) X() float64 {
	defer qt.Recovering("QMagnetometerReading::x")

	if ptr.Pointer() != nil {
		return float64(C.QMagnetometerReading_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMagnetometerReading) Y() float64 {
	defer qt.Recovering("QMagnetometerReading::y")

	if ptr.Pointer() != nil {
		return float64(C.QMagnetometerReading_Y(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMagnetometerReading) Z() float64 {
	defer qt.Recovering("QMagnetometerReading::z")

	if ptr.Pointer() != nil {
		return float64(C.QMagnetometerReading_Z(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMagnetometerReading) SetCalibrationLevel(calibrationLevel float64) {
	defer qt.Recovering("QMagnetometerReading::setCalibrationLevel")

	if ptr.Pointer() != nil {
		C.QMagnetometerReading_SetCalibrationLevel(ptr.Pointer(), C.double(calibrationLevel))
	}
}

func (ptr *QMagnetometerReading) SetX(x float64) {
	defer qt.Recovering("QMagnetometerReading::setX")

	if ptr.Pointer() != nil {
		C.QMagnetometerReading_SetX(ptr.Pointer(), C.double(x))
	}
}

func (ptr *QMagnetometerReading) SetY(y float64) {
	defer qt.Recovering("QMagnetometerReading::setY")

	if ptr.Pointer() != nil {
		C.QMagnetometerReading_SetY(ptr.Pointer(), C.double(y))
	}
}

func (ptr *QMagnetometerReading) SetZ(z float64) {
	defer qt.Recovering("QMagnetometerReading::setZ")

	if ptr.Pointer() != nil {
		C.QMagnetometerReading_SetZ(ptr.Pointer(), C.double(z))
	}
}

func (ptr *QMagnetometerReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMagnetometerReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMagnetometerReading) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMagnetometerReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMagnetometerReadingTimerEvent
func callbackQMagnetometerReadingTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMagnetometerReading::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMagnetometerReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMagnetometerReading) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMagnetometerReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QMagnetometerReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMagnetometerReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMagnetometerReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QMagnetometerReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMagnetometerReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMagnetometerReading::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMagnetometerReading) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMagnetometerReading::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMagnetometerReadingChildEvent
func callbackQMagnetometerReadingChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMagnetometerReading::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMagnetometerReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMagnetometerReading) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMagnetometerReading::childEvent")

	if ptr.Pointer() != nil {
		C.QMagnetometerReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMagnetometerReading) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMagnetometerReading::childEvent")

	if ptr.Pointer() != nil {
		C.QMagnetometerReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMagnetometerReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMagnetometerReading::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMagnetometerReading) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMagnetometerReading::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMagnetometerReadingCustomEvent
func callbackQMagnetometerReadingCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMagnetometerReading::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMagnetometerReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMagnetometerReading) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMagnetometerReading::customEvent")

	if ptr.Pointer() != nil {
		C.QMagnetometerReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMagnetometerReading) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMagnetometerReading::customEvent")

	if ptr.Pointer() != nil {
		C.QMagnetometerReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
