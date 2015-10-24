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

type QTiltSensorITF interface {
	QSensorITF
	QTiltSensorPTR() *QTiltSensor
}

func PointerFromQTiltSensor(ptr QTiltSensorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTiltSensorPTR().Pointer()
	}
	return nil
}

func QTiltSensorFromPointer(ptr unsafe.Pointer) *QTiltSensor {
	var n = new(QTiltSensor)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QTiltSensor_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTiltSensor) QTiltSensorPTR() *QTiltSensor {
	return ptr
}

func NewQTiltSensor(parent core.QObjectITF) *QTiltSensor {
	return QTiltSensorFromPointer(unsafe.Pointer(C.QTiltSensor_NewQTiltSensor(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QTiltSensor) Reading() *QTiltReading {
	if ptr.Pointer() != nil {
		return QTiltReadingFromPointer(unsafe.Pointer(C.QTiltSensor_Reading(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QTiltSensor) DestroyQTiltSensor() {
	if ptr.Pointer() != nil {
		C.QTiltSensor_DestroyQTiltSensor(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QTiltSensor) Calibrate() {
	if ptr.Pointer() != nil {
		C.QTiltSensor_Calibrate(C.QtObjectPtr(ptr.Pointer()))
	}
}
