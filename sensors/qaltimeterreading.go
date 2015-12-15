package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QAltimeterReading struct {
	QSensorReading
}

type QAltimeterReading_ITF interface {
	QSensorReading_ITF
	QAltimeterReading_PTR() *QAltimeterReading
}

func PointerFromQAltimeterReading(ptr QAltimeterReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAltimeterReading_PTR().Pointer()
	}
	return nil
}

func NewQAltimeterReadingFromPointer(ptr unsafe.Pointer) *QAltimeterReading {
	var n = new(QAltimeterReading)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAltimeterReading_") {
		n.SetObjectName("QAltimeterReading_" + qt.Identifier())
	}
	return n
}

func (ptr *QAltimeterReading) QAltimeterReading_PTR() *QAltimeterReading {
	return ptr
}

func (ptr *QAltimeterReading) Altitude() float64 {
	defer qt.Recovering("QAltimeterReading::altitude")

	if ptr.Pointer() != nil {
		return float64(C.QAltimeterReading_Altitude(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAltimeterReading) SetAltitude(altitude float64) {
	defer qt.Recovering("QAltimeterReading::setAltitude")

	if ptr.Pointer() != nil {
		C.QAltimeterReading_SetAltitude(ptr.Pointer(), C.double(altitude))
	}
}
