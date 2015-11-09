package sensors

//#include "qsensorbackendfactory.h"
import "C"
import (
	"unsafe"
)

type QSensorBackendFactory struct {
	ptr unsafe.Pointer
}

type QSensorBackendFactory_ITF interface {
	QSensorBackendFactory_PTR() *QSensorBackendFactory
}

func (p *QSensorBackendFactory) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSensorBackendFactory) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSensorBackendFactory(ptr QSensorBackendFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorBackendFactory_PTR().Pointer()
	}
	return nil
}

func NewQSensorBackendFactoryFromPointer(ptr unsafe.Pointer) *QSensorBackendFactory {
	var n = new(QSensorBackendFactory)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSensorBackendFactory) QSensorBackendFactory_PTR() *QSensorBackendFactory {
	return ptr
}

func (ptr *QSensorBackendFactory) CreateBackend(sensor QSensor_ITF) *QSensorBackend {
	if ptr.Pointer() != nil {
		return NewQSensorBackendFromPointer(C.QSensorBackendFactory_CreateBackend(ptr.Pointer(), PointerFromQSensor(sensor)))
	}
	return nil
}
