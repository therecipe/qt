package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QOrientationSensor struct {
	QSensor
}

type QOrientationSensor_ITF interface {
	QSensor_ITF
	QOrientationSensor_PTR() *QOrientationSensor
}

func PointerFromQOrientationSensor(ptr QOrientationSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOrientationSensor_PTR().Pointer()
	}
	return nil
}

func NewQOrientationSensorFromPointer(ptr unsafe.Pointer) *QOrientationSensor {
	var n = new(QOrientationSensor)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QOrientationSensor_") {
		n.SetObjectName("QOrientationSensor_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QOrientationSensor) QOrientationSensor_PTR() *QOrientationSensor {
	return ptr
}

func (ptr *QOrientationSensor) Reading() *QOrientationReading {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QOrientationSensor::reading")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQOrientationReadingFromPointer(C.QOrientationSensor_Reading(ptr.Pointer()))
	}
	return nil
}

func NewQOrientationSensor(parent core.QObject_ITF) *QOrientationSensor {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QOrientationSensor::QOrientationSensor")
		}
	}()

	return NewQOrientationSensorFromPointer(C.QOrientationSensor_NewQOrientationSensor(core.PointerFromQObject(parent)))
}

func (ptr *QOrientationSensor) DestroyQOrientationSensor() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QOrientationSensor::~QOrientationSensor")
		}
	}()

	if ptr.Pointer() != nil {
		C.QOrientationSensor_DestroyQOrientationSensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
