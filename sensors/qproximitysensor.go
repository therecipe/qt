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

type QProximitySensor_ITF interface {
	QSensor_ITF
	QProximitySensor_PTR() *QProximitySensor
}

func PointerFromQProximitySensor(ptr QProximitySensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QProximitySensor_PTR().Pointer()
	}
	return nil
}

func NewQProximitySensorFromPointer(ptr unsafe.Pointer) *QProximitySensor {
	var n = new(QProximitySensor)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QProximitySensor_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QProximitySensor) QProximitySensor_PTR() *QProximitySensor {
	return ptr
}

func (ptr *QProximitySensor) Reading() *QProximityReading {
	if ptr.Pointer() != nil {
		return NewQProximityReadingFromPointer(C.QProximitySensor_Reading(ptr.Pointer()))
	}
	return nil
}

func NewQProximitySensor(parent core.QObject_ITF) *QProximitySensor {
	return NewQProximitySensorFromPointer(C.QProximitySensor_NewQProximitySensor(core.PointerFromQObject(parent)))
}

func (ptr *QProximitySensor) DestroyQProximitySensor() {
	if ptr.Pointer() != nil {
		C.QProximitySensor_DestroyQProximitySensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
