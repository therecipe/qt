package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QDistanceSensor struct {
	QSensor
}

type QDistanceSensor_ITF interface {
	QSensor_ITF
	QDistanceSensor_PTR() *QDistanceSensor
}

func PointerFromQDistanceSensor(ptr QDistanceSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDistanceSensor_PTR().Pointer()
	}
	return nil
}

func NewQDistanceSensorFromPointer(ptr unsafe.Pointer) *QDistanceSensor {
	var n = new(QDistanceSensor)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QDistanceSensor_") {
		n.SetObjectName("QDistanceSensor_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QDistanceSensor) QDistanceSensor_PTR() *QDistanceSensor {
	return ptr
}

func (ptr *QDistanceSensor) Reading() *QDistanceReading {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDistanceSensor::reading")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQDistanceReadingFromPointer(C.QDistanceSensor_Reading(ptr.Pointer()))
	}
	return nil
}

func NewQDistanceSensor(parent core.QObject_ITF) *QDistanceSensor {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDistanceSensor::QDistanceSensor")
		}
	}()

	return NewQDistanceSensorFromPointer(C.QDistanceSensor_NewQDistanceSensor(core.PointerFromQObject(parent)))
}

func (ptr *QDistanceSensor) DestroyQDistanceSensor() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDistanceSensor::~QDistanceSensor")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDistanceSensor_DestroyQDistanceSensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
