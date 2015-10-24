package sensors

//#include "qirproximityreading.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QIRProximityReading struct {
	QSensorReading
}

type QIRProximityReadingITF interface {
	QSensorReadingITF
	QIRProximityReadingPTR() *QIRProximityReading
}

func PointerFromQIRProximityReading(ptr QIRProximityReadingITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QIRProximityReadingPTR().Pointer()
	}
	return nil
}

func QIRProximityReadingFromPointer(ptr unsafe.Pointer) *QIRProximityReading {
	var n = new(QIRProximityReading)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QIRProximityReading_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QIRProximityReading) QIRProximityReadingPTR() *QIRProximityReading {
	return ptr
}
