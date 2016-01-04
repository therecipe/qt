package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QTiltReading struct {
	QSensorReading
}

type QTiltReading_ITF interface {
	QSensorReading_ITF
	QTiltReading_PTR() *QTiltReading
}

func PointerFromQTiltReading(ptr QTiltReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTiltReading_PTR().Pointer()
	}
	return nil
}

func NewQTiltReadingFromPointer(ptr unsafe.Pointer) *QTiltReading {
	var n = new(QTiltReading)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QTiltReading_") {
		n.SetObjectName("QTiltReading_" + qt.Identifier())
	}
	return n
}

func (ptr *QTiltReading) QTiltReading_PTR() *QTiltReading {
	return ptr
}

func (ptr *QTiltReading) XRotation() float64 {
	defer qt.Recovering("QTiltReading::xRotation")

	if ptr.Pointer() != nil {
		return float64(C.QTiltReading_XRotation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTiltReading) YRotation() float64 {
	defer qt.Recovering("QTiltReading::yRotation")

	if ptr.Pointer() != nil {
		return float64(C.QTiltReading_YRotation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTiltReading) SetXRotation(x float64) {
	defer qt.Recovering("QTiltReading::setXRotation")

	if ptr.Pointer() != nil {
		C.QTiltReading_SetXRotation(ptr.Pointer(), C.double(x))
	}
}

func (ptr *QTiltReading) SetYRotation(y float64) {
	defer qt.Recovering("QTiltReading::setYRotation")

	if ptr.Pointer() != nil {
		C.QTiltReading_SetYRotation(ptr.Pointer(), C.double(y))
	}
}

func (ptr *QTiltReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QTiltReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QTiltReading) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QTiltReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQTiltReadingTimerEvent
func callbackQTiltReadingTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTiltReading::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQTiltReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QTiltReading) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QTiltReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QTiltReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QTiltReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QTiltReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QTiltReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QTiltReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QTiltReading::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QTiltReading) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QTiltReading::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQTiltReadingChildEvent
func callbackQTiltReadingChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTiltReading::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQTiltReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QTiltReading) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QTiltReading::childEvent")

	if ptr.Pointer() != nil {
		C.QTiltReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QTiltReading) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QTiltReading::childEvent")

	if ptr.Pointer() != nil {
		C.QTiltReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QTiltReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QTiltReading::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QTiltReading) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QTiltReading::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQTiltReadingCustomEvent
func callbackQTiltReadingCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTiltReading::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQTiltReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QTiltReading) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QTiltReading::customEvent")

	if ptr.Pointer() != nil {
		C.QTiltReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QTiltReading) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QTiltReading::customEvent")

	if ptr.Pointer() != nil {
		C.QTiltReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
