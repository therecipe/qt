package sensors

//#include "qdistancesensor.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QDistanceSensor struct {
	QSensor
}

type QDistanceSensorITF interface {
	QSensorITF
	QDistanceSensorPTR() *QDistanceSensor
}

func PointerFromQDistanceSensor(ptr QDistanceSensorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDistanceSensorPTR().Pointer()
	}
	return nil
}

func QDistanceSensorFromPointer(ptr unsafe.Pointer) *QDistanceSensor {
	var n = new(QDistanceSensor)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QDistanceSensor_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QDistanceSensor) QDistanceSensorPTR() *QDistanceSensor {
	return ptr
}

func (ptr *QDistanceSensor) Reading() *QDistanceReading {
	if ptr.Pointer() != nil {
		return QDistanceReadingFromPointer(unsafe.Pointer(C.QDistanceSensor_Reading(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func NewQDistanceSensor(parent core.QObjectITF) *QDistanceSensor {
	return QDistanceSensorFromPointer(unsafe.Pointer(C.QDistanceSensor_NewQDistanceSensor(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QDistanceSensor) DestroyQDistanceSensor() {
	if ptr.Pointer() != nil {
		C.QDistanceSensor_DestroyQDistanceSensor(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
