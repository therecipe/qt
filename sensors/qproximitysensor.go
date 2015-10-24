package sensors

//#include "qproximitysensor.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QProximitySensor struct {
	QSensor
}

type QProximitySensorITF interface {
	QSensorITF
	QProximitySensorPTR() *QProximitySensor
}

func PointerFromQProximitySensor(ptr QProximitySensorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QProximitySensorPTR().Pointer()
	}
	return nil
}

func QProximitySensorFromPointer(ptr unsafe.Pointer) *QProximitySensor {
	var n = new(QProximitySensor)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QProximitySensor_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QProximitySensor) QProximitySensorPTR() *QProximitySensor {
	return ptr
}

func (ptr *QProximitySensor) Reading() *QProximityReading {
	if ptr.Pointer() != nil {
		return QProximityReadingFromPointer(unsafe.Pointer(C.QProximitySensor_Reading(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func NewQProximitySensor(parent core.QObjectITF) *QProximitySensor {
	return QProximitySensorFromPointer(unsafe.Pointer(C.QProximitySensor_NewQProximitySensor(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QProximitySensor) DestroyQProximitySensor() {
	if ptr.Pointer() != nil {
		C.QProximitySensor_DestroyQProximitySensor(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
