package sensors

//#include "qpressurereading.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QPressureReading struct {
	QSensorReading
}

type QPressureReadingITF interface {
	QSensorReadingITF
	QPressureReadingPTR() *QPressureReading
}

func PointerFromQPressureReading(ptr QPressureReadingITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPressureReadingPTR().Pointer()
	}
	return nil
}

func QPressureReadingFromPointer(ptr unsafe.Pointer) *QPressureReading {
	var n = new(QPressureReading)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QPressureReading_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QPressureReading) QPressureReadingPTR() *QPressureReading {
	return ptr
}
