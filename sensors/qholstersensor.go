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

type QHolsterSensor_ITF interface {
	QSensor_ITF
	QHolsterSensor_PTR() *QHolsterSensor
}

func PointerFromQHolsterSensor(ptr QHolsterSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHolsterSensor_PTR().Pointer()
	}
	return nil
}

func NewQHolsterSensorFromPointer(ptr unsafe.Pointer) *QHolsterSensor {
	var n = new(QHolsterSensor)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QHolsterSensor_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QHolsterSensor) QHolsterSensor_PTR() *QHolsterSensor {
	return ptr
}

func (ptr *QHolsterSensor) Reading() *QHolsterReading {
	if ptr.Pointer() != nil {
		return NewQHolsterReadingFromPointer(C.QHolsterSensor_Reading(ptr.Pointer()))
	}
	return nil
}

func NewQHolsterSensor(parent core.QObject_ITF) *QHolsterSensor {
	return NewQHolsterSensorFromPointer(C.QHolsterSensor_NewQHolsterSensor(core.PointerFromQObject(parent)))
}

func (ptr *QHolsterSensor) DestroyQHolsterSensor() {
	if ptr.Pointer() != nil {
		C.QHolsterSensor_DestroyQHolsterSensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
