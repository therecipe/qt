package sensors

//#include "qsensorbackend.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QSensorBackend struct {
	core.QObject
}

type QSensorBackendITF interface {
	core.QObjectITF
	QSensorBackendPTR() *QSensorBackend
}

func PointerFromQSensorBackend(ptr QSensorBackendITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorBackendPTR().Pointer()
	}
	return nil
}

func QSensorBackendFromPointer(ptr unsafe.Pointer) *QSensorBackend {
	var n = new(QSensorBackend)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QSensorBackend_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSensorBackend) QSensorBackendPTR() *QSensorBackend {
	return ptr
}

func (ptr *QSensorBackend) IsFeatureSupported(feature QSensor__Feature) bool {
	if ptr.Pointer() != nil {
		return C.QSensorBackend_IsFeatureSupported(C.QtObjectPtr(ptr.Pointer()), C.int(feature)) != 0
	}
	return false
}

func (ptr *QSensorBackend) SensorBusy() {
	if ptr.Pointer() != nil {
		C.QSensorBackend_SensorBusy(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSensorBackend) SensorError(error int) {
	if ptr.Pointer() != nil {
		C.QSensorBackend_SensorError(C.QtObjectPtr(ptr.Pointer()), C.int(error))
	}
}

func (ptr *QSensorBackend) NewReadingAvailable() {
	if ptr.Pointer() != nil {
		C.QSensorBackend_NewReadingAvailable(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSensorBackend) Reading() *QSensorReading {
	if ptr.Pointer() != nil {
		return QSensorReadingFromPointer(unsafe.Pointer(C.QSensorBackend_Reading(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QSensorBackend) Sensor() *QSensor {
	if ptr.Pointer() != nil {
		return QSensorFromPointer(unsafe.Pointer(C.QSensorBackend_Sensor(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QSensorBackend) SensorStopped() {
	if ptr.Pointer() != nil {
		C.QSensorBackend_SensorStopped(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSensorBackend) SetDataRates(otherSensor QSensorITF) {
	if ptr.Pointer() != nil {
		C.QSensorBackend_SetDataRates(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSensor(otherSensor)))
	}
}

func (ptr *QSensorBackend) SetDescription(description string) {
	if ptr.Pointer() != nil {
		C.QSensorBackend_SetDescription(C.QtObjectPtr(ptr.Pointer()), C.CString(description))
	}
}

func (ptr *QSensorBackend) Start() {
	if ptr.Pointer() != nil {
		C.QSensorBackend_Start(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSensorBackend) Stop() {
	if ptr.Pointer() != nil {
		C.QSensorBackend_Stop(C.QtObjectPtr(ptr.Pointer()))
	}
}
