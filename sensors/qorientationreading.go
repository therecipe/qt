package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QOrientationReading struct {
	QSensorReading
}

type QOrientationReading_ITF interface {
	QSensorReading_ITF
	QOrientationReading_PTR() *QOrientationReading
}

func PointerFromQOrientationReading(ptr QOrientationReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOrientationReading_PTR().Pointer()
	}
	return nil
}

func NewQOrientationReadingFromPointer(ptr unsafe.Pointer) *QOrientationReading {
	var n = new(QOrientationReading)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QOrientationReading_") {
		n.SetObjectName("QOrientationReading_" + qt.Identifier())
	}
	return n
}

func (ptr *QOrientationReading) QOrientationReading_PTR() *QOrientationReading {
	return ptr
}

//QOrientationReading::Orientation
type QOrientationReading__Orientation int64

const (
	QOrientationReading__Undefined = QOrientationReading__Orientation(0)
	QOrientationReading__TopUp     = QOrientationReading__Orientation(1)
	QOrientationReading__TopDown   = QOrientationReading__Orientation(2)
	QOrientationReading__LeftUp    = QOrientationReading__Orientation(3)
	QOrientationReading__RightUp   = QOrientationReading__Orientation(4)
	QOrientationReading__FaceUp    = QOrientationReading__Orientation(5)
	QOrientationReading__FaceDown  = QOrientationReading__Orientation(6)
)

func (ptr *QOrientationReading) Orientation() QOrientationReading__Orientation {
	defer qt.Recovering("QOrientationReading::orientation")

	if ptr.Pointer() != nil {
		return QOrientationReading__Orientation(C.QOrientationReading_Orientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QOrientationReading) SetOrientation(orientation QOrientationReading__Orientation) {
	defer qt.Recovering("QOrientationReading::setOrientation")

	if ptr.Pointer() != nil {
		C.QOrientationReading_SetOrientation(ptr.Pointer(), C.int(orientation))
	}
}

func (ptr *QOrientationReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QOrientationReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QOrientationReading) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QOrientationReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQOrientationReadingTimerEvent
func callbackQOrientationReadingTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QOrientationReading::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQOrientationReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QOrientationReading) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QOrientationReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QOrientationReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QOrientationReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QOrientationReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QOrientationReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QOrientationReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QOrientationReading::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QOrientationReading) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QOrientationReading::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQOrientationReadingChildEvent
func callbackQOrientationReadingChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QOrientationReading::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQOrientationReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QOrientationReading) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QOrientationReading::childEvent")

	if ptr.Pointer() != nil {
		C.QOrientationReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QOrientationReading) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QOrientationReading::childEvent")

	if ptr.Pointer() != nil {
		C.QOrientationReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QOrientationReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QOrientationReading::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QOrientationReading) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QOrientationReading::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQOrientationReadingCustomEvent
func callbackQOrientationReadingCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QOrientationReading::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQOrientationReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QOrientationReading) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QOrientationReading::customEvent")

	if ptr.Pointer() != nil {
		C.QOrientationReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QOrientationReading) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QOrientationReading::customEvent")

	if ptr.Pointer() != nil {
		C.QOrientationReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
