package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QAccelerometerReading struct {
	QSensorReading
}

type QAccelerometerReading_ITF interface {
	QSensorReading_ITF
	QAccelerometerReading_PTR() *QAccelerometerReading
}

func PointerFromQAccelerometerReading(ptr QAccelerometerReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccelerometerReading_PTR().Pointer()
	}
	return nil
}

func NewQAccelerometerReadingFromPointer(ptr unsafe.Pointer) *QAccelerometerReading {
	var n = new(QAccelerometerReading)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAccelerometerReading_") {
		n.SetObjectName("QAccelerometerReading_" + qt.Identifier())
	}
	return n
}

func (ptr *QAccelerometerReading) QAccelerometerReading_PTR() *QAccelerometerReading {
	return ptr
}

func (ptr *QAccelerometerReading) X() float64 {
	defer qt.Recovering("QAccelerometerReading::x")

	if ptr.Pointer() != nil {
		return float64(C.QAccelerometerReading_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccelerometerReading) Y() float64 {
	defer qt.Recovering("QAccelerometerReading::y")

	if ptr.Pointer() != nil {
		return float64(C.QAccelerometerReading_Y(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccelerometerReading) Z() float64 {
	defer qt.Recovering("QAccelerometerReading::z")

	if ptr.Pointer() != nil {
		return float64(C.QAccelerometerReading_Z(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccelerometerReading) SetX(x float64) {
	defer qt.Recovering("QAccelerometerReading::setX")

	if ptr.Pointer() != nil {
		C.QAccelerometerReading_SetX(ptr.Pointer(), C.double(x))
	}
}

func (ptr *QAccelerometerReading) SetY(y float64) {
	defer qt.Recovering("QAccelerometerReading::setY")

	if ptr.Pointer() != nil {
		C.QAccelerometerReading_SetY(ptr.Pointer(), C.double(y))
	}
}

func (ptr *QAccelerometerReading) SetZ(z float64) {
	defer qt.Recovering("QAccelerometerReading::setZ")

	if ptr.Pointer() != nil {
		C.QAccelerometerReading_SetZ(ptr.Pointer(), C.double(z))
	}
}
