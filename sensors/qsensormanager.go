package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QSensorManager struct {
	ptr unsafe.Pointer
}

type QSensorManager_ITF interface {
	QSensorManager_PTR() *QSensorManager
}

func (p *QSensorManager) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSensorManager) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSensorManager(ptr QSensorManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorManager_PTR().Pointer()
	}
	return nil
}

func NewQSensorManagerFromPointer(ptr unsafe.Pointer) *QSensorManager {
	var n = new(QSensorManager)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSensorManager) QSensorManager_PTR() *QSensorManager {
	return ptr
}

func QSensorManager_CreateBackend(sensor QSensor_ITF) *QSensorBackend {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensorManager::createBackend")
		}
	}()

	return NewQSensorBackendFromPointer(C.QSensorManager_QSensorManager_CreateBackend(PointerFromQSensor(sensor)))
}

func QSensorManager_IsBackendRegistered(ty core.QByteArray_ITF, identifier core.QByteArray_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensorManager::isBackendRegistered")
		}
	}()

	return C.QSensorManager_QSensorManager_IsBackendRegistered(core.PointerFromQByteArray(ty), core.PointerFromQByteArray(identifier)) != 0
}

func QSensorManager_RegisterBackend(ty core.QByteArray_ITF, identifier core.QByteArray_ITF, factory QSensorBackendFactory_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensorManager::registerBackend")
		}
	}()

	C.QSensorManager_QSensorManager_RegisterBackend(core.PointerFromQByteArray(ty), core.PointerFromQByteArray(identifier), PointerFromQSensorBackendFactory(factory))
}

func QSensorManager_SetDefaultBackend(ty core.QByteArray_ITF, identifier core.QByteArray_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensorManager::setDefaultBackend")
		}
	}()

	C.QSensorManager_QSensorManager_SetDefaultBackend(core.PointerFromQByteArray(ty), core.PointerFromQByteArray(identifier))
}

func QSensorManager_UnregisterBackend(ty core.QByteArray_ITF, identifier core.QByteArray_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensorManager::unregisterBackend")
		}
	}()

	C.QSensorManager_QSensorManager_UnregisterBackend(core.PointerFromQByteArray(ty), core.PointerFromQByteArray(identifier))
}
