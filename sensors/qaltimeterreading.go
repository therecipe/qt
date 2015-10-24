package sensors

//#include "qaltimeterreading.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QAltimeterReading struct {
	QSensorReading
}

type QAltimeterReadingITF interface {
	QSensorReadingITF
	QAltimeterReadingPTR() *QAltimeterReading
}

func PointerFromQAltimeterReading(ptr QAltimeterReadingITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAltimeterReadingPTR().Pointer()
	}
	return nil
}

func QAltimeterReadingFromPointer(ptr unsafe.Pointer) *QAltimeterReading {
	var n = new(QAltimeterReading)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QAltimeterReading_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAltimeterReading) QAltimeterReadingPTR() *QAltimeterReading {
	return ptr
}
