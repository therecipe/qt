package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QRotationSensor struct {
	QSensor
}

type QRotationSensor_ITF interface {
	QSensor_ITF
	QRotationSensor_PTR() *QRotationSensor
}

func PointerFromQRotationSensor(ptr QRotationSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRotationSensor_PTR().Pointer()
	}
	return nil
}

func NewQRotationSensorFromPointer(ptr unsafe.Pointer) *QRotationSensor {
	var n = new(QRotationSensor)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QRotationSensor_") {
		n.SetObjectName("QRotationSensor_" + qt.Identifier())
	}
	return n
}

func (ptr *QRotationSensor) QRotationSensor_PTR() *QRotationSensor {
	return ptr
}

func (ptr *QRotationSensor) HasZ() bool {
	defer qt.Recovering("QRotationSensor::hasZ")

	if ptr.Pointer() != nil {
		return C.QRotationSensor_HasZ(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QRotationSensor) Reading() *QRotationReading {
	defer qt.Recovering("QRotationSensor::reading")

	if ptr.Pointer() != nil {
		return NewQRotationReadingFromPointer(C.QRotationSensor_Reading(ptr.Pointer()))
	}
	return nil
}

func NewQRotationSensor(parent core.QObject_ITF) *QRotationSensor {
	defer qt.Recovering("QRotationSensor::QRotationSensor")

	return NewQRotationSensorFromPointer(C.QRotationSensor_NewQRotationSensor(core.PointerFromQObject(parent)))
}

func (ptr *QRotationSensor) ConnectHasZChanged(f func(hasZ bool)) {
	defer qt.Recovering("connect QRotationSensor::hasZChanged")

	if ptr.Pointer() != nil {
		C.QRotationSensor_ConnectHasZChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "hasZChanged", f)
	}
}

func (ptr *QRotationSensor) DisconnectHasZChanged() {
	defer qt.Recovering("disconnect QRotationSensor::hasZChanged")

	if ptr.Pointer() != nil {
		C.QRotationSensor_DisconnectHasZChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "hasZChanged")
	}
}

//export callbackQRotationSensorHasZChanged
func callbackQRotationSensorHasZChanged(ptrName *C.char, hasZ C.int) {
	defer qt.Recovering("callback QRotationSensor::hasZChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "hasZChanged"); signal != nil {
		signal.(func(bool))(int(hasZ) != 0)
	}

}

func (ptr *QRotationSensor) SetHasZ(hasZ bool) {
	defer qt.Recovering("QRotationSensor::setHasZ")

	if ptr.Pointer() != nil {
		C.QRotationSensor_SetHasZ(ptr.Pointer(), C.int(qt.GoBoolToInt(hasZ)))
	}
}

func (ptr *QRotationSensor) DestroyQRotationSensor() {
	defer qt.Recovering("QRotationSensor::~QRotationSensor")

	if ptr.Pointer() != nil {
		C.QRotationSensor_DestroyQRotationSensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QRotationSensor) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QRotationSensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QRotationSensor) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QRotationSensor::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQRotationSensorTimerEvent
func callbackQRotationSensorTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRotationSensor::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QRotationSensor) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QRotationSensor::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QRotationSensor) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QRotationSensor::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQRotationSensorChildEvent
func callbackQRotationSensorChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRotationSensor::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QRotationSensor) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QRotationSensor::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QRotationSensor) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QRotationSensor::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQRotationSensorCustomEvent
func callbackQRotationSensorCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRotationSensor::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
