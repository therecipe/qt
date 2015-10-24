package sensors

//#include "qorientationsensor.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QOrientationSensor struct {
	QSensor
}

type QOrientationSensorITF interface {
	QSensorITF
	QOrientationSensorPTR() *QOrientationSensor
}

func PointerFromQOrientationSensor(ptr QOrientationSensorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOrientationSensorPTR().Pointer()
	}
	return nil
}

func QOrientationSensorFromPointer(ptr unsafe.Pointer) *QOrientationSensor {
	var n = new(QOrientationSensor)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QOrientationSensor_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QOrientationSensor) QOrientationSensorPTR() *QOrientationSensor {
	return ptr
}

func (ptr *QOrientationSensor) Reading() *QOrientationReading {
	if ptr.Pointer() != nil {
		return QOrientationReadingFromPointer(unsafe.Pointer(C.QOrientationSensor_Reading(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func NewQOrientationSensor(parent core.QObjectITF) *QOrientationSensor {
	return QOrientationSensorFromPointer(unsafe.Pointer(C.QOrientationSensor_NewQOrientationSensor(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QOrientationSensor) DestroyQOrientationSensor() {
	if ptr.Pointer() != nil {
		C.QOrientationSensor_DestroyQOrientationSensor(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
