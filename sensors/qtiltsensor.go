package sensors

//#include "qtiltsensor.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QTiltSensor struct {
	QSensor
}

type QTiltSensor_ITF interface {
	QSensor_ITF
	QTiltSensor_PTR() *QTiltSensor
}

func PointerFromQTiltSensor(ptr QTiltSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTiltSensor_PTR().Pointer()
	}
	return nil
}

func NewQTiltSensorFromPointer(ptr unsafe.Pointer) *QTiltSensor {
	var n = new(QTiltSensor)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QTiltSensor_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTiltSensor) QTiltSensor_PTR() *QTiltSensor {
	return ptr
}

func NewQTiltSensor(parent core.QObject_ITF) *QTiltSensor {
	return NewQTiltSensorFromPointer(C.QTiltSensor_NewQTiltSensor(core.PointerFromQObject(parent)))
}

func (ptr *QTiltSensor) Reading() *QTiltReading {
	if ptr.Pointer() != nil {
		return NewQTiltReadingFromPointer(C.QTiltSensor_Reading(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTiltSensor) DestroyQTiltSensor() {
	if ptr.Pointer() != nil {
		C.QTiltSensor_DestroyQTiltSensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QTiltSensor) Calibrate() {
	if ptr.Pointer() != nil {
		C.QTiltSensor_Calibrate(ptr.Pointer())
	}
}
