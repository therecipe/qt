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

type QLightReading_ITF interface {
	QSensorReading_ITF
	QLightReading_PTR() *QLightReading
}

func PointerFromQLightReading(ptr QLightReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLightReading_PTR().Pointer()
	}
	return nil
}

func NewQLightReadingFromPointer(ptr unsafe.Pointer) *QLightReading {
	var n = new(QLightReading)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QLightReading_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QLightReading) QLightReading_PTR() *QLightReading {
	return ptr
}

func (ptr *QLightReading) Lux() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QLightReading_Lux(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLightReading) SetLux(lux float64) {
	if ptr.Pointer() != nil {
		C.QLightReading_SetLux(ptr.Pointer(), C.double(lux))
	}
}
