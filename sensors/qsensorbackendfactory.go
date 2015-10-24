package sensors

//#include "qsensorbackendfactory.h"
import "C"
import (
	"unsafe"
)

type QSensorBackendFactory struct {
	ptr unsafe.Pointer
}

type QSensorBackendFactoryITF interface {
	QSensorBackendFactoryPTR() *QSensorBackendFactory
}

func (p *QSensorBackendFactory) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSensorBackendFactory) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSensorBackendFactory(ptr QSensorBackendFactoryITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorBackendFactoryPTR().Pointer()
	}
	return nil
}

func QSensorBackendFactoryFromPointer(ptr unsafe.Pointer) *QSensorBackendFactory {
	var n = new(QSensorBackendFactory)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSensorBackendFactory) QSensorBackendFactoryPTR() *QSensorBackendFactory {
	return ptr
}

func (ptr *QSensorBackendFactory) CreateBackend(sensor QSensorITF) *QSensorBackend {
	if ptr.Pointer() != nil {
		return QSensorBackendFromPointer(unsafe.Pointer(C.QSensorBackendFactory_CreateBackend(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSensor(sensor)))))
	}
	return nil
}
