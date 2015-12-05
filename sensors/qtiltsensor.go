package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QTiltSensor struct {
	QSensor
}

type QTiltSensor_ITF interface {
	QSensor_ITF
	QTiltSensor_PTR() *QTiltSensor
}

func PointerFromQTiltSensor(ptr QTiltSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTiltSensor_PTR().Pointer()
	}
	return nil
}

func NewQTiltSensorFromPointer(ptr unsafe.Pointer) *QTiltSensor {
	var n = new(QTiltSensor)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QTiltSensor_") {
		n.SetObjectName("QTiltSensor_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTiltSensor) QTiltSensor_PTR() *QTiltSensor {
	return ptr
}

func NewQTiltSensor(parent core.QObject_ITF) *QTiltSensor {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTiltSensor::QTiltSensor")
		}
	}()

	return NewQTiltSensorFromPointer(C.QTiltSensor_NewQTiltSensor(core.PointerFromQObject(parent)))
}

func (ptr *QTiltSensor) Reading() *QTiltReading {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTiltSensor::reading")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQTiltReadingFromPointer(C.QTiltSensor_Reading(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTiltSensor) DestroyQTiltSensor() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTiltSensor::~QTiltSensor")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTiltSensor_DestroyQTiltSensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QTiltSensor) Calibrate() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTiltSensor::calibrate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTiltSensor_Calibrate(ptr.Pointer())
	}
}
