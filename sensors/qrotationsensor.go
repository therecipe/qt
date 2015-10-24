package sensors

//#include "qrotationsensor.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QRotationSensor struct {
	QSensor
}

type QRotationSensorITF interface {
	QSensorITF
	QRotationSensorPTR() *QRotationSensor
}

func PointerFromQRotationSensor(ptr QRotationSensorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRotationSensorPTR().Pointer()
	}
	return nil
}

func QRotationSensorFromPointer(ptr unsafe.Pointer) *QRotationSensor {
	var n = new(QRotationSensor)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QRotationSensor_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QRotationSensor) QRotationSensorPTR() *QRotationSensor {
	return ptr
}

func (ptr *QRotationSensor) HasZ() bool {
	if ptr.Pointer() != nil {
		return C.QRotationSensor_HasZ(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QRotationSensor) Reading() *QRotationReading {
	if ptr.Pointer() != nil {
		return QRotationReadingFromPointer(unsafe.Pointer(C.QRotationSensor_Reading(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func NewQRotationSensor(parent core.QObjectITF) *QRotationSensor {
	return QRotationSensorFromPointer(unsafe.Pointer(C.QRotationSensor_NewQRotationSensor(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QRotationSensor) ConnectHasZChanged(f func(hasZ bool)) {
	if ptr.Pointer() != nil {
		C.QRotationSensor_ConnectHasZChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "hasZChanged", f)
	}
}

func (ptr *QRotationSensor) DisconnectHasZChanged() {
	if ptr.Pointer() != nil {
		C.QRotationSensor_DisconnectHasZChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "hasZChanged")
	}
}

//export callbackQRotationSensorHasZChanged
func callbackQRotationSensorHasZChanged(ptrName *C.char, hasZ C.int) {
	qt.GetSignal(C.GoString(ptrName), "hasZChanged").(func(bool))(int(hasZ) != 0)
}

func (ptr *QRotationSensor) SetHasZ(hasZ bool) {
	if ptr.Pointer() != nil {
		C.QRotationSensor_SetHasZ(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(hasZ)))
	}
}

func (ptr *QRotationSensor) DestroyQRotationSensor() {
	if ptr.Pointer() != nil {
		C.QRotationSensor_DestroyQRotationSensor(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
