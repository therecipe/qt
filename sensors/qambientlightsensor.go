package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAmbientLightSensor struct {
	QSensor
}

type QAmbientLightSensor_ITF interface {
	QSensor_ITF
	QAmbientLightSensor_PTR() *QAmbientLightSensor
}

func PointerFromQAmbientLightSensor(ptr QAmbientLightSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAmbientLightSensor_PTR().Pointer()
	}
	return nil
}

func NewQAmbientLightSensorFromPointer(ptr unsafe.Pointer) *QAmbientLightSensor {
	var n = new(QAmbientLightSensor)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAmbientLightSensor_") {
		n.SetObjectName("QAmbientLightSensor_" + qt.Identifier())
	}
	return n
}

func (ptr *QAmbientLightSensor) QAmbientLightSensor_PTR() *QAmbientLightSensor {
	return ptr
}

func (ptr *QAmbientLightSensor) Reading() *QAmbientLightReading {
	defer qt.Recovering("QAmbientLightSensor::reading")

	if ptr.Pointer() != nil {
		return NewQAmbientLightReadingFromPointer(C.QAmbientLightSensor_Reading(ptr.Pointer()))
	}
	return nil
}

func NewQAmbientLightSensor(parent core.QObject_ITF) *QAmbientLightSensor {
	defer qt.Recovering("QAmbientLightSensor::QAmbientLightSensor")

	return NewQAmbientLightSensorFromPointer(C.QAmbientLightSensor_NewQAmbientLightSensor(core.PointerFromQObject(parent)))
}

func (ptr *QAmbientLightSensor) DestroyQAmbientLightSensor() {
	defer qt.Recovering("QAmbientLightSensor::~QAmbientLightSensor")

	if ptr.Pointer() != nil {
		C.QAmbientLightSensor_DestroyQAmbientLightSensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
