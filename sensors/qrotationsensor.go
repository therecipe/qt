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

	var signal = qt.GetSignal(C.GoString(ptrName), "hasZChanged")
	if signal != nil {
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
