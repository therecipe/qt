package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QAmbientLightSensor struct {
	QSensor
}

type QAmbientLightSensor_ITF interface {
	QSensor_ITF
	QAmbientLightSensor_PTR() *QAmbientLightSensor
}

func PointerFromQAmbientLightSensor(ptr QAmbientLightSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAmbientLightSensor_PTR().Pointer()
	}
	return nil
}

func NewQAmbientLightSensorFromPointer(ptr unsafe.Pointer) *QAmbientLightSensor {
	var n = new(QAmbientLightSensor)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAmbientLightSensor_") {
		n.SetObjectName("QAmbientLightSensor_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAmbientLightSensor) QAmbientLightSensor_PTR() *QAmbientLightSensor {
	return ptr
}

func (ptr *QAmbientLightSensor) Reading() *QAmbientLightReading {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAmbientLightSensor::reading")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQAmbientLightReadingFromPointer(C.QAmbientLightSensor_Reading(ptr.Pointer()))
	}
	return nil
}

func NewQAmbientLightSensor(parent core.QObject_ITF) *QAmbientLightSensor {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAmbientLightSensor::QAmbientLightSensor")
		}
	}()

	return NewQAmbientLightSensorFromPointer(C.QAmbientLightSensor_NewQAmbientLightSensor(core.PointerFromQObject(parent)))
}

func (ptr *QAmbientLightSensor) DestroyQAmbientLightSensor() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAmbientLightSensor::~QAmbientLightSensor")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAmbientLightSensor_DestroyQAmbientLightSensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
