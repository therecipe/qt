package sensors

//#include "qpressuresensor.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QPressureSensor struct {
	QSensor
}

type QPressureSensorITF interface {
	QSensorITF
	QPressureSensorPTR() *QPressureSensor
}

func PointerFromQPressureSensor(ptr QPressureSensorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPressureSensorPTR().Pointer()
	}
	return nil
}

func QPressureSensorFromPointer(ptr unsafe.Pointer) *QPressureSensor {
	var n = new(QPressureSensor)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QPressureSensor_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QPressureSensor) QPressureSensorPTR() *QPressureSensor {
	return ptr
}

func (ptr *QPressureSensor) Reading() *QPressureReading {
	if ptr.Pointer() != nil {
		return QPressureReadingFromPointer(unsafe.Pointer(C.QPressureSensor_Reading(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func NewQPressureSensor(parent core.QObjectITF) *QPressureSensor {
	return QPressureSensorFromPointer(unsafe.Pointer(C.QPressureSensor_NewQPressureSensor(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QPressureSensor) DestroyQPressureSensor() {
	if ptr.Pointer() != nil {
		C.QPressureSensor_DestroyQPressureSensor(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
