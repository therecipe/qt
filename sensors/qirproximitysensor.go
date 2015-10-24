package sensors

//#include "qirproximitysensor.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QIRProximitySensor struct {
	QSensor
}

type QIRProximitySensorITF interface {
	QSensorITF
	QIRProximitySensorPTR() *QIRProximitySensor
}

func PointerFromQIRProximitySensor(ptr QIRProximitySensorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QIRProximitySensorPTR().Pointer()
	}
	return nil
}

func QIRProximitySensorFromPointer(ptr unsafe.Pointer) *QIRProximitySensor {
	var n = new(QIRProximitySensor)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QIRProximitySensor_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QIRProximitySensor) QIRProximitySensorPTR() *QIRProximitySensor {
	return ptr
}

func (ptr *QIRProximitySensor) Reading() *QIRProximityReading {
	if ptr.Pointer() != nil {
		return QIRProximityReadingFromPointer(unsafe.Pointer(C.QIRProximitySensor_Reading(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func NewQIRProximitySensor(parent core.QObjectITF) *QIRProximitySensor {
	return QIRProximitySensorFromPointer(unsafe.Pointer(C.QIRProximitySensor_NewQIRProximitySensor(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QIRProximitySensor) DestroyQIRProximitySensor() {
	if ptr.Pointer() != nil {
		C.QIRProximitySensor_DestroyQIRProximitySensor(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
