package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QCompassReading struct {
	QSensorReading
}

type QCompassReading_ITF interface {
	QSensorReading_ITF
	QCompassReading_PTR() *QCompassReading
}

func PointerFromQCompassReading(ptr QCompassReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCompassReading_PTR().Pointer()
	}
	return nil
}

func NewQCompassReadingFromPointer(ptr unsafe.Pointer) *QCompassReading {
	var n = new(QCompassReading)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QCompassReading_") {
		n.SetObjectName("QCompassReading_" + qt.Identifier())
	}
	return n
}

func (ptr *QCompassReading) QCompassReading_PTR() *QCompassReading {
	return ptr
}

func (ptr *QCompassReading) Azimuth() float64 {
	defer qt.Recovering("QCompassReading::azimuth")

	if ptr.Pointer() != nil {
		return float64(C.QCompassReading_Azimuth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCompassReading) CalibrationLevel() float64 {
	defer qt.Recovering("QCompassReading::calibrationLevel")

	if ptr.Pointer() != nil {
		return float64(C.QCompassReading_CalibrationLevel(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCompassReading) SetAzimuth(azimuth float64) {
	defer qt.Recovering("QCompassReading::setAzimuth")

	if ptr.Pointer() != nil {
		C.QCompassReading_SetAzimuth(ptr.Pointer(), C.double(azimuth))
	}
}

func (ptr *QCompassReading) SetCalibrationLevel(calibrationLevel float64) {
	defer qt.Recovering("QCompassReading::setCalibrationLevel")

	if ptr.Pointer() != nil {
		C.QCompassReading_SetCalibrationLevel(ptr.Pointer(), C.double(calibrationLevel))
	}
}

func (ptr *QCompassReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QCompassReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCompassReading) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCompassReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQCompassReadingTimerEvent
func callbackQCompassReadingTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCompassReading::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQCompassReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QCompassReading) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCompassReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QCompassReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCompassReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCompassReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QCompassReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCompassReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCompassReading::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCompassReading) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCompassReading::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQCompassReadingChildEvent
func callbackQCompassReadingChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCompassReading::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCompassReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCompassReading) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCompassReading::childEvent")

	if ptr.Pointer() != nil {
		C.QCompassReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCompassReading) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCompassReading::childEvent")

	if ptr.Pointer() != nil {
		C.QCompassReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCompassReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCompassReading::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCompassReading) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCompassReading::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQCompassReadingCustomEvent
func callbackQCompassReadingCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCompassReading::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCompassReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCompassReading) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCompassReading::customEvent")

	if ptr.Pointer() != nil {
		C.QCompassReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCompassReading) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCompassReading::customEvent")

	if ptr.Pointer() != nil {
		C.QCompassReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
