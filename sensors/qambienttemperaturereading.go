package sensors

//#include "qambienttemperaturereading.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QAmbientTemperatureReading struct {
	QSensorReading
}

type QAmbientTemperatureReading_ITF interface {
	QSensorReading_ITF
	QAmbientTemperatureReading_PTR() *QAmbientTemperatureReading
}

func PointerFromQAmbientTemperatureReading(ptr QAmbientTemperatureReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAmbientTemperatureReading_PTR().Pointer()
	}
	return nil
}

func NewQAmbientTemperatureReadingFromPointer(ptr unsafe.Pointer) *QAmbientTemperatureReading {
	var n = new(QAmbientTemperatureReading)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QAmbientTemperatureReading_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAmbientTemperatureReading) QAmbientTemperatureReading_PTR() *QAmbientTemperatureReading {
	return ptr
}

func (ptr *QAmbientTemperatureReading) Temperature() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QAmbientTemperatureReading_Temperature(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAmbientTemperatureReading) SetTemperature(temperature float64) {
	if ptr.Pointer() != nil {
		C.QAmbientTemperatureReading_SetTemperature(ptr.Pointer(), C.double(temperature))
	}
}
