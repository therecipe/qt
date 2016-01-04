package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QTapReading struct {
	QSensorReading
}

type QTapReading_ITF interface {
	QSensorReading_ITF
	QTapReading_PTR() *QTapReading
}

func PointerFromQTapReading(ptr QTapReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTapReading_PTR().Pointer()
	}
	return nil
}

func NewQTapReadingFromPointer(ptr unsafe.Pointer) *QTapReading {
	var n = new(QTapReading)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QTapReading_") {
		n.SetObjectName("QTapReading_" + qt.Identifier())
	}
	return n
}

func (ptr *QTapReading) QTapReading_PTR() *QTapReading {
	return ptr
}

//QTapReading::TapDirection
type QTapReading__TapDirection int64

const (
	QTapReading__Undefined = QTapReading__TapDirection(0)
	QTapReading__X         = QTapReading__TapDirection(0x0001)
	QTapReading__Y         = QTapReading__TapDirection(0x0002)
	QTapReading__Z         = QTapReading__TapDirection(0x0004)
	QTapReading__X_Pos     = QTapReading__TapDirection(0x0011)
	QTapReading__Y_Pos     = QTapReading__TapDirection(0x0022)
	QTapReading__Z_Pos     = QTapReading__TapDirection(0x0044)
	QTapReading__X_Neg     = QTapReading__TapDirection(0x0101)
	QTapReading__Y_Neg     = QTapReading__TapDirection(0x0202)
	QTapReading__Z_Neg     = QTapReading__TapDirection(0x0404)
	QTapReading__X_Both    = QTapReading__TapDirection(0x0111)
	QTapReading__Y_Both    = QTapReading__TapDirection(0x0222)
	QTapReading__Z_Both    = QTapReading__TapDirection(0x0444)
)

func (ptr *QTapReading) IsDoubleTap() bool {
	defer qt.Recovering("QTapReading::isDoubleTap")

	if ptr.Pointer() != nil {
		return C.QTapReading_IsDoubleTap(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTapReading) TapDirection() QTapReading__TapDirection {
	defer qt.Recovering("QTapReading::tapDirection")

	if ptr.Pointer() != nil {
		return QTapReading__TapDirection(C.QTapReading_TapDirection(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTapReading) SetDoubleTap(doubleTap bool) {
	defer qt.Recovering("QTapReading::setDoubleTap")

	if ptr.Pointer() != nil {
		C.QTapReading_SetDoubleTap(ptr.Pointer(), C.int(qt.GoBoolToInt(doubleTap)))
	}
}

func (ptr *QTapReading) SetTapDirection(tapDirection QTapReading__TapDirection) {
	defer qt.Recovering("QTapReading::setTapDirection")

	if ptr.Pointer() != nil {
		C.QTapReading_SetTapDirection(ptr.Pointer(), C.int(tapDirection))
	}
}

func (ptr *QTapReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QTapReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QTapReading) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QTapReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQTapReadingTimerEvent
func callbackQTapReadingTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTapReading::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQTapReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QTapReading) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QTapReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QTapReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QTapReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QTapReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QTapReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QTapReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QTapReading::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QTapReading) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QTapReading::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQTapReadingChildEvent
func callbackQTapReadingChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTapReading::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQTapReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QTapReading) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QTapReading::childEvent")

	if ptr.Pointer() != nil {
		C.QTapReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QTapReading) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QTapReading::childEvent")

	if ptr.Pointer() != nil {
		C.QTapReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QTapReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QTapReading::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QTapReading) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QTapReading::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQTapReadingCustomEvent
func callbackQTapReadingCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTapReading::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQTapReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QTapReading) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QTapReading::customEvent")

	if ptr.Pointer() != nil {
		C.QTapReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QTapReading) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QTapReading::customEvent")

	if ptr.Pointer() != nil {
		C.QTapReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
