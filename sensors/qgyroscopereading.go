package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QGyroscopeReading struct {
	QSensorReading
}

type QGyroscopeReading_ITF interface {
	QSensorReading_ITF
	QGyroscopeReading_PTR() *QGyroscopeReading
}

func PointerFromQGyroscopeReading(ptr QGyroscopeReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGyroscopeReading_PTR().Pointer()
	}
	return nil
}

func NewQGyroscopeReadingFromPointer(ptr unsafe.Pointer) *QGyroscopeReading {
	var n = new(QGyroscopeReading)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QGyroscopeReading_") {
		n.SetObjectName("QGyroscopeReading_" + qt.Identifier())
	}
	return n
}

func (ptr *QGyroscopeReading) QGyroscopeReading_PTR() *QGyroscopeReading {
	return ptr
}

func (ptr *QGyroscopeReading) X() float64 {
	defer qt.Recovering("QGyroscopeReading::x")

	if ptr.Pointer() != nil {
		return float64(C.QGyroscopeReading_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGyroscopeReading) Y() float64 {
	defer qt.Recovering("QGyroscopeReading::y")

	if ptr.Pointer() != nil {
		return float64(C.QGyroscopeReading_Y(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGyroscopeReading) Z() float64 {
	defer qt.Recovering("QGyroscopeReading::z")

	if ptr.Pointer() != nil {
		return float64(C.QGyroscopeReading_Z(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGyroscopeReading) SetX(x float64) {
	defer qt.Recovering("QGyroscopeReading::setX")

	if ptr.Pointer() != nil {
		C.QGyroscopeReading_SetX(ptr.Pointer(), C.double(x))
	}
}

func (ptr *QGyroscopeReading) SetY(y float64) {
	defer qt.Recovering("QGyroscopeReading::setY")

	if ptr.Pointer() != nil {
		C.QGyroscopeReading_SetY(ptr.Pointer(), C.double(y))
	}
}

func (ptr *QGyroscopeReading) SetZ(z float64) {
	defer qt.Recovering("QGyroscopeReading::setZ")

	if ptr.Pointer() != nil {
		C.QGyroscopeReading_SetZ(ptr.Pointer(), C.double(z))
	}
}

func (ptr *QGyroscopeReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QGyroscopeReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QGyroscopeReading) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QGyroscopeReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQGyroscopeReadingTimerEvent
func callbackQGyroscopeReadingTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGyroscopeReading::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QGyroscopeReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QGyroscopeReading::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QGyroscopeReading) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QGyroscopeReading::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQGyroscopeReadingChildEvent
func callbackQGyroscopeReadingChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGyroscopeReading::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QGyroscopeReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGyroscopeReading::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QGyroscopeReading) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QGyroscopeReading::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQGyroscopeReadingCustomEvent
func callbackQGyroscopeReadingCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGyroscopeReading::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
