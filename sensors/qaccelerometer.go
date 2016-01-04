package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAccelerometer struct {
	QSensor
}

type QAccelerometer_ITF interface {
	QSensor_ITF
	QAccelerometer_PTR() *QAccelerometer
}

func PointerFromQAccelerometer(ptr QAccelerometer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccelerometer_PTR().Pointer()
	}
	return nil
}

func NewQAccelerometerFromPointer(ptr unsafe.Pointer) *QAccelerometer {
	var n = new(QAccelerometer)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAccelerometer_") {
		n.SetObjectName("QAccelerometer_" + qt.Identifier())
	}
	return n
}

func (ptr *QAccelerometer) QAccelerometer_PTR() *QAccelerometer {
	return ptr
}

//QAccelerometer::AccelerationMode
type QAccelerometer__AccelerationMode int64

const (
	QAccelerometer__Combined = QAccelerometer__AccelerationMode(0)
	QAccelerometer__Gravity  = QAccelerometer__AccelerationMode(1)
	QAccelerometer__User     = QAccelerometer__AccelerationMode(2)
)

func (ptr *QAccelerometer) AccelerationMode() QAccelerometer__AccelerationMode {
	defer qt.Recovering("QAccelerometer::accelerationMode")

	if ptr.Pointer() != nil {
		return QAccelerometer__AccelerationMode(C.QAccelerometer_AccelerationMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccelerometer) Reading() *QAccelerometerReading {
	defer qt.Recovering("QAccelerometer::reading")

	if ptr.Pointer() != nil {
		return NewQAccelerometerReadingFromPointer(C.QAccelerometer_Reading(ptr.Pointer()))
	}
	return nil
}

func NewQAccelerometer(parent core.QObject_ITF) *QAccelerometer {
	defer qt.Recovering("QAccelerometer::QAccelerometer")

	return NewQAccelerometerFromPointer(C.QAccelerometer_NewQAccelerometer(core.PointerFromQObject(parent)))
}

func (ptr *QAccelerometer) ConnectAccelerationModeChanged(f func(accelerationMode QAccelerometer__AccelerationMode)) {
	defer qt.Recovering("connect QAccelerometer::accelerationModeChanged")

	if ptr.Pointer() != nil {
		C.QAccelerometer_ConnectAccelerationModeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "accelerationModeChanged", f)
	}
}

func (ptr *QAccelerometer) DisconnectAccelerationModeChanged() {
	defer qt.Recovering("disconnect QAccelerometer::accelerationModeChanged")

	if ptr.Pointer() != nil {
		C.QAccelerometer_DisconnectAccelerationModeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "accelerationModeChanged")
	}
}

//export callbackQAccelerometerAccelerationModeChanged
func callbackQAccelerometerAccelerationModeChanged(ptr unsafe.Pointer, ptrName *C.char, accelerationMode C.int) {
	defer qt.Recovering("callback QAccelerometer::accelerationModeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "accelerationModeChanged"); signal != nil {
		signal.(func(QAccelerometer__AccelerationMode))(QAccelerometer__AccelerationMode(accelerationMode))
	}

}

func (ptr *QAccelerometer) AccelerationModeChanged(accelerationMode QAccelerometer__AccelerationMode) {
	defer qt.Recovering("QAccelerometer::accelerationModeChanged")

	if ptr.Pointer() != nil {
		C.QAccelerometer_AccelerationModeChanged(ptr.Pointer(), C.int(accelerationMode))
	}
}

func (ptr *QAccelerometer) SetAccelerationMode(accelerationMode QAccelerometer__AccelerationMode) {
	defer qt.Recovering("QAccelerometer::setAccelerationMode")

	if ptr.Pointer() != nil {
		C.QAccelerometer_SetAccelerationMode(ptr.Pointer(), C.int(accelerationMode))
	}
}

func (ptr *QAccelerometer) DestroyQAccelerometer() {
	defer qt.Recovering("QAccelerometer::~QAccelerometer")

	if ptr.Pointer() != nil {
		C.QAccelerometer_DestroyQAccelerometer(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAccelerometer) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAccelerometer::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAccelerometer) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAccelerometer::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAccelerometerTimerEvent
func callbackQAccelerometerTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAccelerometer::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAccelerometerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAccelerometer) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAccelerometer::timerEvent")

	if ptr.Pointer() != nil {
		C.QAccelerometer_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAccelerometer) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAccelerometer::timerEvent")

	if ptr.Pointer() != nil {
		C.QAccelerometer_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAccelerometer) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAccelerometer::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAccelerometer) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAccelerometer::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAccelerometerChildEvent
func callbackQAccelerometerChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAccelerometer::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAccelerometerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAccelerometer) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAccelerometer::childEvent")

	if ptr.Pointer() != nil {
		C.QAccelerometer_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAccelerometer) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAccelerometer::childEvent")

	if ptr.Pointer() != nil {
		C.QAccelerometer_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAccelerometer) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAccelerometer::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAccelerometer) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAccelerometer::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAccelerometerCustomEvent
func callbackQAccelerometerCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAccelerometer::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAccelerometerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAccelerometer) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAccelerometer::customEvent")

	if ptr.Pointer() != nil {
		C.QAccelerometer_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAccelerometer) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAccelerometer::customEvent")

	if ptr.Pointer() != nil {
		C.QAccelerometer_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
