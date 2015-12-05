package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
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
	for len(n.ObjectName()) < len("QPressureSensor_") {
		n.SetObjectName("QPressureSensor_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QPressureSensor) QPressureSensor_PTR() *QPressureSensor {
	return ptr
}

func (ptr *QPressureSensor) Reading() *QPressureReading {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPressureSensor::reading")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQPressureReadingFromPointer(C.QPressureSensor_Reading(ptr.Pointer()))
	}
	return nil
}

func NewQPressureSensor(parent core.QObject_ITF) *QPressureSensor {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPressureSensor::QPressureSensor")
		}
	}()

	return NewQPressureSensorFromPointer(C.QPressureSensor_NewQPressureSensor(core.PointerFromQObject(parent)))
}

func (ptr *QPressureSensor) DestroyQPressureSensor() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPressureSensor::~QPressureSensor")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPressureSensor_DestroyQPressureSensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
