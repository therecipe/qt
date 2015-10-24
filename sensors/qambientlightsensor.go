package sensors

//#include "qambientlightsensor.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAmbientLightSensor struct {
	QSensor
}

type QAmbientLightSensorITF interface {
	QSensorITF
	QAmbientLightSensorPTR() *QAmbientLightSensor
}

func PointerFromQAmbientLightSensor(ptr QAmbientLightSensorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAmbientLightSensorPTR().Pointer()
	}
	return nil
}

func QAmbientLightSensorFromPointer(ptr unsafe.Pointer) *QAmbientLightSensor {
	var n = new(QAmbientLightSensor)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QAmbientLightSensor_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAmbientLightSensor) QAmbientLightSensorPTR() *QAmbientLightSensor {
	return ptr
}

func (ptr *QAmbientLightSensor) Reading() *QAmbientLightReading {
	if ptr.Pointer() != nil {
		return QAmbientLightReadingFromPointer(unsafe.Pointer(C.QAmbientLightSensor_Reading(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func NewQAmbientLightSensor(parent core.QObjectITF) *QAmbientLightSensor {
	return QAmbientLightSensorFromPointer(unsafe.Pointer(C.QAmbientLightSensor_NewQAmbientLightSensor(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QAmbientLightSensor) DestroyQAmbientLightSensor() {
	if ptr.Pointer() != nil {
		C.QAmbientLightSensor_DestroyQAmbientLightSensor(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
