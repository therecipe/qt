package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QMagnetometer struct {
	QSensor
}

type QMagnetometer_ITF interface {
	QSensor_ITF
	QMagnetometer_PTR() *QMagnetometer
}

func PointerFromQMagnetometer(ptr QMagnetometer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMagnetometer_PTR().Pointer()
	}
	return nil
}

func NewQMagnetometerFromPointer(ptr unsafe.Pointer) *QMagnetometer {
	var n = new(QMagnetometer)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMagnetometer_") {
		n.SetObjectName("QMagnetometer_" + qt.Identifier())
	}
	return n
}

func (ptr *QMagnetometer) QMagnetometer_PTR() *QMagnetometer {
	return ptr
}

func (ptr *QMagnetometer) Reading() *QMagnetometerReading {
	defer qt.Recovering("QMagnetometer::reading")

	if ptr.Pointer() != nil {
		return NewQMagnetometerReadingFromPointer(C.QMagnetometer_Reading(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMagnetometer) ReturnGeoValues() bool {
	defer qt.Recovering("QMagnetometer::returnGeoValues")

	if ptr.Pointer() != nil {
		return C.QMagnetometer_ReturnGeoValues(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMagnetometer) SetReturnGeoValues(returnGeoValues bool) {
	defer qt.Recovering("QMagnetometer::setReturnGeoValues")

	if ptr.Pointer() != nil {
		C.QMagnetometer_SetReturnGeoValues(ptr.Pointer(), C.int(qt.GoBoolToInt(returnGeoValues)))
	}
}

func NewQMagnetometer(parent core.QObject_ITF) *QMagnetometer {
	defer qt.Recovering("QMagnetometer::QMagnetometer")

	return NewQMagnetometerFromPointer(C.QMagnetometer_NewQMagnetometer(core.PointerFromQObject(parent)))
}

func (ptr *QMagnetometer) ConnectReturnGeoValuesChanged(f func(returnGeoValues bool)) {
	defer qt.Recovering("connect QMagnetometer::returnGeoValuesChanged")

	if ptr.Pointer() != nil {
		C.QMagnetometer_ConnectReturnGeoValuesChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "returnGeoValuesChanged", f)
	}
}

func (ptr *QMagnetometer) DisconnectReturnGeoValuesChanged() {
	defer qt.Recovering("disconnect QMagnetometer::returnGeoValuesChanged")

	if ptr.Pointer() != nil {
		C.QMagnetometer_DisconnectReturnGeoValuesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "returnGeoValuesChanged")
	}
}

//export callbackQMagnetometerReturnGeoValuesChanged
func callbackQMagnetometerReturnGeoValuesChanged(ptr unsafe.Pointer, ptrName *C.char, returnGeoValues C.int) {
	defer qt.Recovering("callback QMagnetometer::returnGeoValuesChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "returnGeoValuesChanged"); signal != nil {
		signal.(func(bool))(int(returnGeoValues) != 0)
	}

}

func (ptr *QMagnetometer) ReturnGeoValuesChanged(returnGeoValues bool) {
	defer qt.Recovering("QMagnetometer::returnGeoValuesChanged")

	if ptr.Pointer() != nil {
		C.QMagnetometer_ReturnGeoValuesChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(returnGeoValues)))
	}
}

func (ptr *QMagnetometer) DestroyQMagnetometer() {
	defer qt.Recovering("QMagnetometer::~QMagnetometer")

	if ptr.Pointer() != nil {
		C.QMagnetometer_DestroyQMagnetometer(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMagnetometer) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMagnetometer::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMagnetometer) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMagnetometer::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMagnetometerTimerEvent
func callbackQMagnetometerTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMagnetometer::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMagnetometerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMagnetometer) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMagnetometer::timerEvent")

	if ptr.Pointer() != nil {
		C.QMagnetometer_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMagnetometer) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMagnetometer::timerEvent")

	if ptr.Pointer() != nil {
		C.QMagnetometer_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMagnetometer) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMagnetometer::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMagnetometer) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMagnetometer::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMagnetometerChildEvent
func callbackQMagnetometerChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMagnetometer::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMagnetometerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMagnetometer) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMagnetometer::childEvent")

	if ptr.Pointer() != nil {
		C.QMagnetometer_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMagnetometer) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMagnetometer::childEvent")

	if ptr.Pointer() != nil {
		C.QMagnetometer_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMagnetometer) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMagnetometer::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMagnetometer) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMagnetometer::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMagnetometerCustomEvent
func callbackQMagnetometerCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMagnetometer::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMagnetometerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMagnetometer) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMagnetometer::customEvent")

	if ptr.Pointer() != nil {
		C.QMagnetometer_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMagnetometer) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMagnetometer::customEvent")

	if ptr.Pointer() != nil {
		C.QMagnetometer_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
