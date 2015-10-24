package sensors

//#include "qmagnetometerreading.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QMagnetometerReading struct {
	QSensorReading
}

type QMagnetometerReadingITF interface {
	QSensorReadingITF
	QMagnetometerReadingPTR() *QMagnetometerReading
}

func PointerFromQMagnetometerReading(ptr QMagnetometerReadingITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMagnetometerReadingPTR().Pointer()
	}
	return nil
}

func QMagnetometerReadingFromPointer(ptr unsafe.Pointer) *QMagnetometerReading {
	var n = new(QMagnetometerReading)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QMagnetometerReading_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMagnetometerReading) QMagnetometerReadingPTR() *QMagnetometerReading {
	return ptr
}
