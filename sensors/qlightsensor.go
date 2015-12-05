package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QLightSensor struct {
	QSensor
}

type QLightSensor_ITF interface {
	QSensor_ITF
	QLightSensor_PTR() *QLightSensor
}

func PointerFromQLightSensor(ptr QLightSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLightSensor_PTR().Pointer()
	}
	return nil
}

func NewQLightSensorFromPointer(ptr unsafe.Pointer) *QLightSensor {
	var n = new(QLightSensor)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QLightSensor_") {
		n.SetObjectName("QLightSensor_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QLightSensor) QLightSensor_PTR() *QLightSensor {
	return ptr
}

func (ptr *QLightSensor) FieldOfView() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLightSensor::fieldOfView")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QLightSensor_FieldOfView(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLightSensor) Reading() *QLightReading {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLightSensor::reading")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQLightReadingFromPointer(C.QLightSensor_Reading(ptr.Pointer()))
	}
	return nil
}

func NewQLightSensor(parent core.QObject_ITF) *QLightSensor {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLightSensor::QLightSensor")
		}
	}()

	return NewQLightSensorFromPointer(C.QLightSensor_NewQLightSensor(core.PointerFromQObject(parent)))
}

func (ptr *QLightSensor) SetFieldOfView(fieldOfView float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLightSensor::setFieldOfView")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLightSensor_SetFieldOfView(ptr.Pointer(), C.double(fieldOfView))
	}
}

func (ptr *QLightSensor) DestroyQLightSensor() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLightSensor::~QLightSensor")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLightSensor_DestroyQLightSensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
