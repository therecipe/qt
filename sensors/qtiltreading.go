package sensors

//#include "qtiltreading.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QTiltReading struct {
	QSensorReading
}

type QTiltReading_ITF interface {
	QSensorReading_ITF
	QTiltReading_PTR() *QTiltReading
}

func PointerFromQTiltReading(ptr QTiltReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTiltReading_PTR().Pointer()
	}
	return nil
}

func NewQTiltReadingFromPointer(ptr unsafe.Pointer) *QTiltReading {
	var n = new(QTiltReading)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QTiltReading_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTiltReading) QTiltReading_PTR() *QTiltReading {
	return ptr
}

func (ptr *QTiltReading) XRotation() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QTiltReading_XRotation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTiltReading) YRotation() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QTiltReading_YRotation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTiltReading) SetXRotation(x float64) {
	if ptr.Pointer() != nil {
		C.QTiltReading_SetXRotation(ptr.Pointer(), C.double(x))
	}
}

func (ptr *QTiltReading) SetYRotation(y float64) {
	if ptr.Pointer() != nil {
		C.QTiltReading_SetYRotation(ptr.Pointer(), C.double(y))
	}
}
