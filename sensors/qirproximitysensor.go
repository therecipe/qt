package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QIRProximitySensor struct {
	QSensor
}

type QIRProximitySensor_ITF interface {
	QSensor_ITF
	QIRProximitySensor_PTR() *QIRProximitySensor
}

func PointerFromQIRProximitySensor(ptr QIRProximitySensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QIRProximitySensor_PTR().Pointer()
	}
	return nil
}

func NewQIRProximitySensorFromPointer(ptr unsafe.Pointer) *QIRProximitySensor {
	var n = new(QIRProximitySensor)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QIRProximitySensor_") {
		n.SetObjectName("QIRProximitySensor_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QIRProximitySensor) QIRProximitySensor_PTR() *QIRProximitySensor {
	return ptr
}

func (ptr *QIRProximitySensor) Reading() *QIRProximityReading {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QIRProximitySensor::reading")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQIRProximityReadingFromPointer(C.QIRProximitySensor_Reading(ptr.Pointer()))
	}
	return nil
}

func NewQIRProximitySensor(parent core.QObject_ITF) *QIRProximitySensor {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QIRProximitySensor::QIRProximitySensor")
		}
	}()

	return NewQIRProximitySensorFromPointer(C.QIRProximitySensor_NewQIRProximitySensor(core.PointerFromQObject(parent)))
}

func (ptr *QIRProximitySensor) DestroyQIRProximitySensor() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QIRProximitySensor::~QIRProximitySensor")
		}
	}()

	if ptr.Pointer() != nil {
		C.QIRProximitySensor_DestroyQIRProximitySensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
