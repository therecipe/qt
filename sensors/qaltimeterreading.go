package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAltimeterReading struct {
	QSensorReading
}

type QAltimeterReading_ITF interface {
	QSensorReading_ITF
	QAltimeterReading_PTR() *QAltimeterReading
}

func PointerFromQAltimeterReading(ptr QAltimeterReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAltimeterReading_PTR().Pointer()
	}
	return nil
}

func NewQAltimeterReadingFromPointer(ptr unsafe.Pointer) *QAltimeterReading {
	var n = new(QAltimeterReading)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAltimeterReading_") {
		n.SetObjectName("QAltimeterReading_" + qt.Identifier())
	}
	return n
}

func (ptr *QAltimeterReading) QAltimeterReading_PTR() *QAltimeterReading {
	return ptr
}

func (ptr *QAltimeterReading) Altitude() float64 {
	defer qt.Recovering("QAltimeterReading::altitude")

	if ptr.Pointer() != nil {
		return float64(C.QAltimeterReading_Altitude(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAltimeterReading) SetAltitude(altitude float64) {
	defer qt.Recovering("QAltimeterReading::setAltitude")

	if ptr.Pointer() != nil {
		C.QAltimeterReading_SetAltitude(ptr.Pointer(), C.double(altitude))
	}
}

func (ptr *QAltimeterReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAltimeterReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAltimeterReading) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAltimeterReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAltimeterReadingTimerEvent
func callbackQAltimeterReadingTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAltimeterReading::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAltimeterReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAltimeterReading) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAltimeterReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QAltimeterReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAltimeterReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAltimeterReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QAltimeterReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAltimeterReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAltimeterReading::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAltimeterReading) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAltimeterReading::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAltimeterReadingChildEvent
func callbackQAltimeterReadingChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAltimeterReading::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAltimeterReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAltimeterReading) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAltimeterReading::childEvent")

	if ptr.Pointer() != nil {
		C.QAltimeterReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAltimeterReading) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAltimeterReading::childEvent")

	if ptr.Pointer() != nil {
		C.QAltimeterReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAltimeterReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAltimeterReading::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAltimeterReading) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAltimeterReading::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAltimeterReadingCustomEvent
func callbackQAltimeterReadingCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAltimeterReading::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAltimeterReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAltimeterReading) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAltimeterReading::customEvent")

	if ptr.Pointer() != nil {
		C.QAltimeterReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAltimeterReading) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAltimeterReading::customEvent")

	if ptr.Pointer() != nil {
		C.QAltimeterReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
