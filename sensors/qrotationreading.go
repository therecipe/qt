package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QRotationReading struct {
	QSensorReading
}

type QRotationReading_ITF interface {
	QSensorReading_ITF
	QRotationReading_PTR() *QRotationReading
}

func PointerFromQRotationReading(ptr QRotationReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRotationReading_PTR().Pointer()
	}
	return nil
}

func NewQRotationReadingFromPointer(ptr unsafe.Pointer) *QRotationReading {
	var n = new(QRotationReading)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QRotationReading_") {
		n.SetObjectName("QRotationReading_" + qt.Identifier())
	}
	return n
}

func (ptr *QRotationReading) QRotationReading_PTR() *QRotationReading {
	return ptr
}

func (ptr *QRotationReading) X() float64 {
	defer qt.Recovering("QRotationReading::x")

	if ptr.Pointer() != nil {
		return float64(C.QRotationReading_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRotationReading) Y() float64 {
	defer qt.Recovering("QRotationReading::y")

	if ptr.Pointer() != nil {
		return float64(C.QRotationReading_Y(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRotationReading) Z() float64 {
	defer qt.Recovering("QRotationReading::z")

	if ptr.Pointer() != nil {
		return float64(C.QRotationReading_Z(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRotationReading) SetFromEuler(x float64, y float64, z float64) {
	defer qt.Recovering("QRotationReading::setFromEuler")

	if ptr.Pointer() != nil {
		C.QRotationReading_SetFromEuler(ptr.Pointer(), C.double(x), C.double(y), C.double(z))
	}
}

func (ptr *QRotationReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QRotationReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QRotationReading) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QRotationReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQRotationReadingTimerEvent
func callbackQRotationReadingTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRotationReading::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQRotationReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QRotationReading) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QRotationReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QRotationReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QRotationReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QRotationReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QRotationReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QRotationReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QRotationReading::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QRotationReading) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QRotationReading::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQRotationReadingChildEvent
func callbackQRotationReadingChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRotationReading::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQRotationReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QRotationReading) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QRotationReading::childEvent")

	if ptr.Pointer() != nil {
		C.QRotationReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QRotationReading) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QRotationReading::childEvent")

	if ptr.Pointer() != nil {
		C.QRotationReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QRotationReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QRotationReading::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QRotationReading) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QRotationReading::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQRotationReadingCustomEvent
func callbackQRotationReadingCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRotationReading::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQRotationReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QRotationReading) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QRotationReading::customEvent")

	if ptr.Pointer() != nil {
		C.QRotationReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QRotationReading) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QRotationReading::customEvent")

	if ptr.Pointer() != nil {
		C.QRotationReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
