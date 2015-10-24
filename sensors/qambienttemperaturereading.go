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

type QAmbientTemperatureReadingITF interface {
	QSensorReadingITF
	QAmbientTemperatureReadingPTR() *QAmbientTemperatureReading
}

func PointerFromQAmbientTemperatureReading(ptr QAmbientTemperatureReadingITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAmbientTemperatureReadingPTR().Pointer()
	}
	return nil
}

func QAmbientTemperatureReadingFromPointer(ptr unsafe.Pointer) *QAmbientTemperatureReading {
	var n = new(QAmbientTemperatureReading)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QAmbientTemperatureReading_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAmbientTemperatureReading) QAmbientTemperatureReadingPTR() *QAmbientTemperatureReading {
	return ptr
}
