package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QProximityReading struct {
	QSensorReading
}

type QProximityReading_ITF interface {
	QSensorReading_ITF
	QProximityReading_PTR() *QProximityReading
}

func PointerFromQProximityReading(ptr QProximityReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QProximityReading_PTR().Pointer()
	}
	return nil
}

func NewQProximityReadingFromPointer(ptr unsafe.Pointer) *QProximityReading {
	var n = new(QProximityReading)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QProximityReading_") {
		n.SetObjectName("QProximityReading_" + qt.Identifier())
	}
	return n
}

func (ptr *QProximityReading) QProximityReading_PTR() *QProximityReading {
	return ptr
}

func (ptr *QProximityReading) Close() bool {
	defer qt.Recovering("QProximityReading::close")

	if ptr.Pointer() != nil {
		return C.QProximityReading_Close(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QProximityReading) SetClose(close bool) {
	defer qt.Recovering("QProximityReading::setClose")

	if ptr.Pointer() != nil {
		C.QProximityReading_SetClose(ptr.Pointer(), C.int(qt.GoBoolToInt(close)))
	}
}

func (ptr *QProximityReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QProximityReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QProximityReading) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QProximityReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQProximityReadingTimerEvent
func callbackQProximityReadingTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProximityReading::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQProximityReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QProximityReading) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QProximityReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QProximityReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QProximityReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QProximityReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QProximityReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QProximityReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QProximityReading::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QProximityReading) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QProximityReading::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQProximityReadingChildEvent
func callbackQProximityReadingChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProximityReading::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQProximityReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QProximityReading) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QProximityReading::childEvent")

	if ptr.Pointer() != nil {
		C.QProximityReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QProximityReading) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QProximityReading::childEvent")

	if ptr.Pointer() != nil {
		C.QProximityReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QProximityReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QProximityReading::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QProximityReading) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QProximityReading::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQProximityReadingCustomEvent
func callbackQProximityReadingCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QProximityReading::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQProximityReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QProximityReading) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QProximityReading::customEvent")

	if ptr.Pointer() != nil {
		C.QProximityReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QProximityReading) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QProximityReading::customEvent")

	if ptr.Pointer() != nil {
		C.QProximityReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
