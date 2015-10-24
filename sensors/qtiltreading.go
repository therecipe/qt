package sensors

//#include "qtiltreading.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QTiltReading struct {
	QSensorReading
}

type QTiltReadingITF interface {
	QSensorReadingITF
	QTiltReadingPTR() *QTiltReading
}

func PointerFromQTiltReading(ptr QTiltReadingITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTiltReadingPTR().Pointer()
	}
	return nil
}

func QTiltReadingFromPointer(ptr unsafe.Pointer) *QTiltReading {
	var n = new(QTiltReading)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QTiltReading_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTiltReading) QTiltReadingPTR() *QTiltReading {
	return ptr
}
