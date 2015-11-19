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

type QDistanceReading_ITF interface {
	QSensorReading_ITF
	QDistanceReading_PTR() *QDistanceReading
}

func PointerFromQDistanceReading(ptr QDistanceReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDistanceReading_PTR().Pointer()
	}
	return nil
}

func NewQDistanceReadingFromPointer(ptr unsafe.Pointer) *QDistanceReading {
	var n = new(QDistanceReading)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QDistanceReading_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QDistanceReading) QDistanceReading_PTR() *QDistanceReading {
	return ptr
}

func (ptr *QDistanceReading) Distance() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QDistanceReading_Distance(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDistanceReading) SetDistance(distance float64) {
	if ptr.Pointer() != nil {
		C.QDistanceReading_SetDistance(ptr.Pointer(), C.double(distance))
	}
}
