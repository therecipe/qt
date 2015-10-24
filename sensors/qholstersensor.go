package sensors

//#include "qholstersensor.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QHolsterSensor struct {
	QSensor
}

type QHolsterSensorITF interface {
	QSensorITF
	QHolsterSensorPTR() *QHolsterSensor
}

func PointerFromQHolsterSensor(ptr QHolsterSensorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHolsterSensorPTR().Pointer()
	}
	return nil
}

func QHolsterSensorFromPointer(ptr unsafe.Pointer) *QHolsterSensor {
	var n = new(QHolsterSensor)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QHolsterSensor_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QHolsterSensor) QHolsterSensorPTR() *QHolsterSensor {
	return ptr
}

func (ptr *QHolsterSensor) Reading() *QHolsterReading {
	if ptr.Pointer() != nil {
		return QHolsterReadingFromPointer(unsafe.Pointer(C.QHolsterSensor_Reading(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func NewQHolsterSensor(parent core.QObjectITF) *QHolsterSensor {
	return QHolsterSensorFromPointer(unsafe.Pointer(C.QHolsterSensor_NewQHolsterSensor(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QHolsterSensor) DestroyQHolsterSensor() {
	if ptr.Pointer() != nil {
		C.QHolsterSensor_DestroyQHolsterSensor(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
