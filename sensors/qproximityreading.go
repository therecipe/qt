package sensors

//#include "qproximityreading.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QProximityReading struct {
	QSensorReading
}

type QProximityReadingITF interface {
	QSensorReadingITF
	QProximityReadingPTR() *QProximityReading
}

func PointerFromQProximityReading(ptr QProximityReadingITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QProximityReadingPTR().Pointer()
	}
	return nil
}

func QProximityReadingFromPointer(ptr unsafe.Pointer) *QProximityReading {
	var n = new(QProximityReading)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QProximityReading_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QProximityReading) QProximityReadingPTR() *QProximityReading {
	return ptr
}

func (ptr *QProximityReading) Close() bool {
	if ptr.Pointer() != nil {
		return C.QProximityReading_Close(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QProximityReading) SetClose(close bool) {
	if ptr.Pointer() != nil {
		C.QProximityReading_SetClose(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(close)))
	}
}
