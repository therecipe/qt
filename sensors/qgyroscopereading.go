package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QGyroscopeReading struct {
	QSensorReading
}

type QGyroscopeReading_ITF interface {
	QSensorReading_ITF
	QGyroscopeReading_PTR() *QGyroscopeReading
}

func PointerFromQGyroscopeReading(ptr QGyroscopeReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGyroscopeReading_PTR().Pointer()
	}
	return nil
}

func NewQGyroscopeReadingFromPointer(ptr unsafe.Pointer) *QGyroscopeReading {
	var n = new(QGyroscopeReading)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QGyroscopeReading_") {
		n.SetObjectName("QGyroscopeReading_" + qt.Identifier())
	}
	return n
}

func (ptr *QGyroscopeReading) QGyroscopeReading_PTR() *QGyroscopeReading {
	return ptr
}

func (ptr *QGyroscopeReading) X() float64 {
	defer qt.Recovering("QGyroscopeReading::x")

	if ptr.Pointer() != nil {
		return float64(C.QGyroscopeReading_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGyroscopeReading) Y() float64 {
	defer qt.Recovering("QGyroscopeReading::y")

	if ptr.Pointer() != nil {
		return float64(C.QGyroscopeReading_Y(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGyroscopeReading) Z() float64 {
	defer qt.Recovering("QGyroscopeReading::z")

	if ptr.Pointer() != nil {
		return float64(C.QGyroscopeReading_Z(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGyroscopeReading) SetX(x float64) {
	defer qt.Recovering("QGyroscopeReading::setX")

	if ptr.Pointer() != nil {
		C.QGyroscopeReading_SetX(ptr.Pointer(), C.double(x))
	}
}

func (ptr *QGyroscopeReading) SetY(y float64) {
	defer qt.Recovering("QGyroscopeReading::setY")

	if ptr.Pointer() != nil {
		C.QGyroscopeReading_SetY(ptr.Pointer(), C.double(y))
	}
}

func (ptr *QGyroscopeReading) SetZ(z float64) {
	defer qt.Recovering("QGyroscopeReading::setZ")

	if ptr.Pointer() != nil {
		C.QGyroscopeReading_SetZ(ptr.Pointer(), C.double(z))
	}
}
