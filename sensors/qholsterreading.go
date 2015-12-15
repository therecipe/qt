package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QHolsterReading struct {
	QSensorReading
}

type QHolsterReading_ITF interface {
	QSensorReading_ITF
	QHolsterReading_PTR() *QHolsterReading
}

func PointerFromQHolsterReading(ptr QHolsterReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHolsterReading_PTR().Pointer()
	}
	return nil
}

func NewQHolsterReadingFromPointer(ptr unsafe.Pointer) *QHolsterReading {
	var n = new(QHolsterReading)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QHolsterReading_") {
		n.SetObjectName("QHolsterReading_" + qt.Identifier())
	}
	return n
}

func (ptr *QHolsterReading) QHolsterReading_PTR() *QHolsterReading {
	return ptr
}

func (ptr *QHolsterReading) Holstered() bool {
	defer qt.Recovering("QHolsterReading::holstered")

	if ptr.Pointer() != nil {
		return C.QHolsterReading_Holstered(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHolsterReading) SetHolstered(holstered bool) {
	defer qt.Recovering("QHolsterReading::setHolstered")

	if ptr.Pointer() != nil {
		C.QHolsterReading_SetHolstered(ptr.Pointer(), C.int(qt.GoBoolToInt(holstered)))
	}
}
