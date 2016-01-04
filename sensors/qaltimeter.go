package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAltimeter struct {
	QSensor
}

type QAltimeter_ITF interface {
	QSensor_ITF
	QAltimeter_PTR() *QAltimeter
}

func PointerFromQAltimeter(ptr QAltimeter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAltimeter_PTR().Pointer()
	}
	return nil
}

func NewQAltimeterFromPointer(ptr unsafe.Pointer) *QAltimeter {
	var n = new(QAltimeter)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAltimeter_") {
		n.SetObjectName("QAltimeter_" + qt.Identifier())
	}
	return n
}

func (ptr *QAltimeter) QAltimeter_PTR() *QAltimeter {
	return ptr
}

func (ptr *QAltimeter) Reading() *QAltimeterReading {
	defer qt.Recovering("QAltimeter::reading")

	if ptr.Pointer() != nil {
		return NewQAltimeterReadingFromPointer(C.QAltimeter_Reading(ptr.Pointer()))
	}
	return nil
}

func NewQAltimeter(parent core.QObject_ITF) *QAltimeter {
	defer qt.Recovering("QAltimeter::QAltimeter")

	return NewQAltimeterFromPointer(C.QAltimeter_NewQAltimeter(core.PointerFromQObject(parent)))
}

func (ptr *QAltimeter) DestroyQAltimeter() {
	defer qt.Recovering("QAltimeter::~QAltimeter")

	if ptr.Pointer() != nil {
		C.QAltimeter_DestroyQAltimeter(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAltimeter) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAltimeter::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAltimeter) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAltimeter::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAltimeterTimerEvent
func callbackQAltimeterTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAltimeter::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAltimeterFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAltimeter) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAltimeter::timerEvent")

	if ptr.Pointer() != nil {
		C.QAltimeter_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAltimeter) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAltimeter::timerEvent")

	if ptr.Pointer() != nil {
		C.QAltimeter_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAltimeter) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAltimeter::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAltimeter) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAltimeter::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAltimeterChildEvent
func callbackQAltimeterChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAltimeter::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAltimeterFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAltimeter) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAltimeter::childEvent")

	if ptr.Pointer() != nil {
		C.QAltimeter_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAltimeter) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAltimeter::childEvent")

	if ptr.Pointer() != nil {
		C.QAltimeter_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAltimeter) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAltimeter::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAltimeter) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAltimeter::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAltimeterCustomEvent
func callbackQAltimeterCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAltimeter::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAltimeterFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAltimeter) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAltimeter::customEvent")

	if ptr.Pointer() != nil {
		C.QAltimeter_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAltimeter) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAltimeter::customEvent")

	if ptr.Pointer() != nil {
		C.QAltimeter_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
