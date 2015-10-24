package sensors

//#include "qsensormanager.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QSensorManager struct {
	ptr unsafe.Pointer
}

type QSensorManagerITF interface {
	QSensorManagerPTR() *QSensorManager
}

func (p *QSensorManager) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSensorManager) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSensorManager(ptr QSensorManagerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorManagerPTR().Pointer()
	}
	return nil
}

func QSensorManagerFromPointer(ptr unsafe.Pointer) *QSensorManager {
	var n = new(QSensorManager)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSensorManager) QSensorManagerPTR() *QSensorManager {
	return ptr
}

func QSensorManager_CreateBackend(sensor QSensorITF) *QSensorBackend {
	return QSensorBackendFromPointer(unsafe.Pointer(C.QSensorManager_QSensorManager_CreateBackend(C.QtObjectPtr(PointerFromQSensor(sensor)))))
}

func QSensorManager_IsBackendRegistered(ty core.QByteArrayITF, identifier core.QByteArrayITF) bool {
	return C.QSensorManager_QSensorManager_IsBackendRegistered(C.QtObjectPtr(core.PointerFromQByteArray(ty)), C.QtObjectPtr(core.PointerFromQByteArray(identifier))) != 0
}

func QSensorManager_RegisterBackend(ty core.QByteArrayITF, identifier core.QByteArrayITF, factory QSensorBackendFactoryITF) {
	C.QSensorManager_QSensorManager_RegisterBackend(C.QtObjectPtr(core.PointerFromQByteArray(ty)), C.QtObjectPtr(core.PointerFromQByteArray(identifier)), C.QtObjectPtr(PointerFromQSensorBackendFactory(factory)))
}

func QSensorManager_SetDefaultBackend(ty core.QByteArrayITF, identifier core.QByteArrayITF) {
	C.QSensorManager_QSensorManager_SetDefaultBackend(C.QtObjectPtr(core.PointerFromQByteArray(ty)), C.QtObjectPtr(core.PointerFromQByteArray(identifier)))
}

func QSensorManager_UnregisterBackend(ty core.QByteArrayITF, identifier core.QByteArrayITF) {
	C.QSensorManager_QSensorManager_UnregisterBackend(C.QtObjectPtr(core.PointerFromQByteArray(ty)), C.QtObjectPtr(core.PointerFromQByteArray(identifier)))
}
