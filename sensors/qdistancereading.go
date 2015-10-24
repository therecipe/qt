package sensors

//#include "qdistancereading.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QDistanceReading struct {
	QSensorReading
}

type QDistanceReadingITF interface {
	QSensorReadingITF
	QDistanceReadingPTR() *QDistanceReading
}

func PointerFromQDistanceReading(ptr QDistanceReadingITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDistanceReadingPTR().Pointer()
	}
	return nil
}

func QDistanceReadingFromPointer(ptr unsafe.Pointer) *QDistanceReading {
	var n = new(QDistanceReading)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QDistanceReading_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QDistanceReading) QDistanceReadingPTR() *QDistanceReading {
	return ptr
}
