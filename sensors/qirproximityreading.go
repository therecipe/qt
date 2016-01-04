package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QIRProximityReading struct {
	QSensorReading
}

type QIRProximityReading_ITF interface {
	QSensorReading_ITF
	QIRProximityReading_PTR() *QIRProximityReading
}

func PointerFromQIRProximityReading(ptr QIRProximityReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QIRProximityReading_PTR().Pointer()
	}
	return nil
}

func NewQIRProximityReadingFromPointer(ptr unsafe.Pointer) *QIRProximityReading {
	var n = new(QIRProximityReading)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QIRProximityReading_") {
		n.SetObjectName("QIRProximityReading_" + qt.Identifier())
	}
	return n
}

func (ptr *QIRProximityReading) QIRProximityReading_PTR() *QIRProximityReading {
	return ptr
}

func (ptr *QIRProximityReading) Reflectance() float64 {
	defer qt.Recovering("QIRProximityReading::reflectance")

	if ptr.Pointer() != nil {
		return float64(C.QIRProximityReading_Reflectance(ptr.Pointer()))
	}
	return 0
}

func (ptr *QIRProximityReading) SetReflectance(reflectance float64) {
	defer qt.Recovering("QIRProximityReading::setReflectance")

	if ptr.Pointer() != nil {
		C.QIRProximityReading_SetReflectance(ptr.Pointer(), C.double(reflectance))
	}
}

func (ptr *QIRProximityReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QIRProximityReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QIRProximityReading) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QIRProximityReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQIRProximityReadingTimerEvent
func callbackQIRProximityReadingTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QIRProximityReading::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQIRProximityReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QIRProximityReading) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QIRProximityReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QIRProximityReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QIRProximityReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QIRProximityReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QIRProximityReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QIRProximityReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QIRProximityReading::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QIRProximityReading) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QIRProximityReading::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQIRProximityReadingChildEvent
func callbackQIRProximityReadingChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QIRProximityReading::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQIRProximityReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QIRProximityReading) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QIRProximityReading::childEvent")

	if ptr.Pointer() != nil {
		C.QIRProximityReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QIRProximityReading) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QIRProximityReading::childEvent")

	if ptr.Pointer() != nil {
		C.QIRProximityReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QIRProximityReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QIRProximityReading::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QIRProximityReading) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QIRProximityReading::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQIRProximityReadingCustomEvent
func callbackQIRProximityReadingCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QIRProximityReading::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQIRProximityReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QIRProximityReading) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QIRProximityReading::customEvent")

	if ptr.Pointer() != nil {
		C.QIRProximityReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QIRProximityReading) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QIRProximityReading::customEvent")

	if ptr.Pointer() != nil {
		C.QIRProximityReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
