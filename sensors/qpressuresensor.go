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

type QPressureSensor_ITF interface {
	QSensor_ITF
	QPressureSensor_PTR() *QPressureSensor
}

func PointerFromQPressureSensor(ptr QPressureSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPressureSensor_PTR().Pointer()
	}
	return nil
}

func NewQPressureSensorFromPointer(ptr unsafe.Pointer) *QPressureSensor {
	var n = new(QPressureSensor)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QPressureSensor_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QPressureSensor) QPressureSensor_PTR() *QPressureSensor {
	return ptr
}

func (ptr *QPressureSensor) Reading() *QPressureReading {
	if ptr.Pointer() != nil {
		return NewQPressureReadingFromPointer(C.QPressureSensor_Reading(ptr.Pointer()))
	}
	return nil
}

func NewQPressureSensor(parent core.QObject_ITF) *QPressureSensor {
	return NewQPressureSensorFromPointer(C.QPressureSensor_NewQPressureSensor(core.PointerFromQObject(parent)))
}

func (ptr *QPressureSensor) DestroyQPressureSensor() {
	if ptr.Pointer() != nil {
		C.QPressureSensor_DestroyQPressureSensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
