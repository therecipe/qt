package sensors

//#include "qrotationreading.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QRotationReading struct {
	QSensorReading
}

type QRotationReadingITF interface {
	QSensorReadingITF
	QRotationReadingPTR() *QRotationReading
}

func PointerFromQRotationReading(ptr QRotationReadingITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRotationReadingPTR().Pointer()
	}
	return nil
}

func QRotationReadingFromPointer(ptr unsafe.Pointer) *QRotationReading {
	var n = new(QRotationReading)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QRotationReading_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QRotationReading) QRotationReadingPTR() *QRotationReading {
	return ptr
}
