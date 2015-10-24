package sensors

//#include "qlightsensor.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QLightSensor struct {
	QSensor
}

type QLightSensorITF interface {
	QSensorITF
	QLightSensorPTR() *QLightSensor
}

func PointerFromQLightSensor(ptr QLightSensorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLightSensorPTR().Pointer()
	}
	return nil
}

func QLightSensorFromPointer(ptr unsafe.Pointer) *QLightSensor {
	var n = new(QLightSensor)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QLightSensor_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QLightSensor) QLightSensorPTR() *QLightSensor {
	return ptr
}

func (ptr *QLightSensor) Reading() *QLightReading {
	if ptr.Pointer() != nil {
		return QLightReadingFromPointer(unsafe.Pointer(C.QLightSensor_Reading(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func NewQLightSensor(parent core.QObjectITF) *QLightSensor {
	return QLightSensorFromPointer(unsafe.Pointer(C.QLightSensor_NewQLightSensor(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QLightSensor) DestroyQLightSensor() {
	if ptr.Pointer() != nil {
		C.QLightSensor_DestroyQLightSensor(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
