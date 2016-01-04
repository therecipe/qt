package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QSensorReading struct {
	core.QObject
}

type QSensorReading_ITF interface {
	core.QObject_ITF
	QSensorReading_PTR() *QSensorReading
}

func PointerFromQSensorReading(ptr QSensorReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorReading_PTR().Pointer()
	}
	return nil
}

func NewQSensorReadingFromPointer(ptr unsafe.Pointer) *QSensorReading {
	var n = new(QSensorReading)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSensorReading_") {
		n.SetObjectName("QSensorReading_" + qt.Identifier())
	}
	return n
}

func (ptr *QSensorReading) QSensorReading_PTR() *QSensorReading {
	return ptr
}

func (ptr *QSensorReading) Value(index int) *core.QVariant {
	defer qt.Recovering("QSensorReading::value")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSensorReading_Value(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QSensorReading) ValueCount() int {
	defer qt.Recovering("QSensorReading::valueCount")

	if ptr.Pointer() != nil {
		return int(C.QSensorReading_ValueCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSensorReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSensorReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSensorReading) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSensorReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSensorReadingTimerEvent
func callbackQSensorReadingTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSensorReading::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSensorReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSensorReading) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSensorReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QSensorReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSensorReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSensorReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QSensorReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSensorReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSensorReading::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSensorReading) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSensorReading::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSensorReadingChildEvent
func callbackQSensorReadingChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSensorReading::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSensorReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSensorReading) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSensorReading::childEvent")

	if ptr.Pointer() != nil {
		C.QSensorReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSensorReading) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSensorReading::childEvent")

	if ptr.Pointer() != nil {
		C.QSensorReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSensorReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSensorReading::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSensorReading) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSensorReading::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSensorReadingCustomEvent
func callbackQSensorReadingCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSensorReading::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSensorReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSensorReading) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSensorReading::customEvent")

	if ptr.Pointer() != nil {
		C.QSensorReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSensorReading) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSensorReading::customEvent")

	if ptr.Pointer() != nil {
		C.QSensorReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
