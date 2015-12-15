package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QProximityReading struct {
	QSensorReading
}

type QProximityReading_ITF interface {
	QSensorReading_ITF
	QProximityReading_PTR() *QProximityReading
}

func PointerFromQProximityReading(ptr QProximityReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QProximityReading_PTR().Pointer()
	}
	return nil
}

func NewQProximityReadingFromPointer(ptr unsafe.Pointer) *QProximityReading {
	var n = new(QProximityReading)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QProximityReading_") {
		n.SetObjectName("QProximityReading_" + qt.Identifier())
	}
	return n
}

func (ptr *QProximityReading) QProximityReading_PTR() *QProximityReading {
	return ptr
}

func (ptr *QProximityReading) Close() bool {
	defer qt.Recovering("QProximityReading::close")

	if ptr.Pointer() != nil {
		return C.QProximityReading_Close(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QProximityReading) SetClose(close bool) {
	defer qt.Recovering("QProximityReading::setClose")

	if ptr.Pointer() != nil {
		C.QProximityReading_SetClose(ptr.Pointer(), C.int(qt.GoBoolToInt(close)))
	}
}
