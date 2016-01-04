package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QDistanceReading struct {
	QSensorReading
}

type QDistanceReading_ITF interface {
	QSensorReading_ITF
	QDistanceReading_PTR() *QDistanceReading
}

func PointerFromQDistanceReading(ptr QDistanceReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDistanceReading_PTR().Pointer()
	}
	return nil
}

func NewQDistanceReadingFromPointer(ptr unsafe.Pointer) *QDistanceReading {
	var n = new(QDistanceReading)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QDistanceReading_") {
		n.SetObjectName("QDistanceReading_" + qt.Identifier())
	}
	return n
}

func (ptr *QDistanceReading) QDistanceReading_PTR() *QDistanceReading {
	return ptr
}

func (ptr *QDistanceReading) Distance() float64 {
	defer qt.Recovering("QDistanceReading::distance")

	if ptr.Pointer() != nil {
		return float64(C.QDistanceReading_Distance(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDistanceReading) SetDistance(distance float64) {
	defer qt.Recovering("QDistanceReading::setDistance")

	if ptr.Pointer() != nil {
		C.QDistanceReading_SetDistance(ptr.Pointer(), C.double(distance))
	}
}

func (ptr *QDistanceReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDistanceReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QDistanceReading) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDistanceReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQDistanceReadingTimerEvent
func callbackQDistanceReadingTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDistanceReading::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDistanceReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDistanceReading) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDistanceReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QDistanceReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDistanceReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDistanceReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QDistanceReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDistanceReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDistanceReading::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QDistanceReading) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDistanceReading::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQDistanceReadingChildEvent
func callbackQDistanceReadingChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDistanceReading::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDistanceReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDistanceReading) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDistanceReading::childEvent")

	if ptr.Pointer() != nil {
		C.QDistanceReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDistanceReading) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDistanceReading::childEvent")

	if ptr.Pointer() != nil {
		C.QDistanceReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDistanceReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDistanceReading::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QDistanceReading) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDistanceReading::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQDistanceReadingCustomEvent
func callbackQDistanceReadingCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDistanceReading::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDistanceReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDistanceReading) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDistanceReading::customEvent")

	if ptr.Pointer() != nil {
		C.QDistanceReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDistanceReading) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDistanceReading::customEvent")

	if ptr.Pointer() != nil {
		C.QDistanceReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
