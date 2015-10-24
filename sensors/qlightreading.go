package sensors

//#include "qlightreading.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QLightReading struct {
	QSensorReading
}

type QLightReadingITF interface {
	QSensorReadingITF
	QLightReadingPTR() *QLightReading
}

func PointerFromQLightReading(ptr QLightReadingITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLightReadingPTR().Pointer()
	}
	return nil
}

func QLightReadingFromPointer(ptr unsafe.Pointer) *QLightReading {
	var n = new(QLightReading)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QLightReading_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QLightReading) QLightReadingPTR() *QLightReading {
	return ptr
}
