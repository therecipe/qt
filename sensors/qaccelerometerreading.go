package sensors

//#include "qaccelerometerreading.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QAccelerometerReading struct {
	QSensorReading
}

type QAccelerometerReadingITF interface {
	QSensorReadingITF
	QAccelerometerReadingPTR() *QAccelerometerReading
}

func PointerFromQAccelerometerReading(ptr QAccelerometerReadingITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccelerometerReadingPTR().Pointer()
	}
	return nil
}

func QAccelerometerReadingFromPointer(ptr unsafe.Pointer) *QAccelerometerReading {
	var n = new(QAccelerometerReading)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QAccelerometerReading_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAccelerometerReading) QAccelerometerReadingPTR() *QAccelerometerReading {
	return ptr
}
