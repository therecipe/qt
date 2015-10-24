package sensors

//#include "qambienttemperaturesensor.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAmbientTemperatureSensor struct {
	QSensor
}

type QAmbientTemperatureSensorITF interface {
	QSensorITF
	QAmbientTemperatureSensorPTR() *QAmbientTemperatureSensor
}

func PointerFromQAmbientTemperatureSensor(ptr QAmbientTemperatureSensorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAmbientTemperatureSensorPTR().Pointer()
	}
	return nil
}

func QAmbientTemperatureSensorFromPointer(ptr unsafe.Pointer) *QAmbientTemperatureSensor {
	var n = new(QAmbientTemperatureSensor)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QAmbientTemperatureSensor_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAmbientTemperatureSensor) QAmbientTemperatureSensorPTR() *QAmbientTemperatureSensor {
	return ptr
}

func (ptr *QAmbientTemperatureSensor) Reading() *QAmbientTemperatureReading {
	if ptr.Pointer() != nil {
		return QAmbientTemperatureReadingFromPointer(unsafe.Pointer(C.QAmbientTemperatureSensor_Reading(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func NewQAmbientTemperatureSensor(parent core.QObjectITF) *QAmbientTemperatureSensor {
	return QAmbientTemperatureSensorFromPointer(unsafe.Pointer(C.QAmbientTemperatureSensor_NewQAmbientTemperatureSensor(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QAmbientTemperatureSensor) DestroyQAmbientTemperatureSensor() {
	if ptr.Pointer() != nil {
		C.QAmbientTemperatureSensor_DestroyQAmbientTemperatureSensor(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
