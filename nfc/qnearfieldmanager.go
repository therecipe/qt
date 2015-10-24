package nfc

//#include "qnearfieldmanager.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QNearFieldManager struct {
	core.QObject
}

type QNearFieldManagerITF interface {
	core.QObjectITF
	QNearFieldManagerPTR() *QNearFieldManager
}

func PointerFromQNearFieldManager(ptr QNearFieldManagerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNearFieldManagerPTR().Pointer()
	}
	return nil
}

func QNearFieldManagerFromPointer(ptr unsafe.Pointer) *QNearFieldManager {
	var n = new(QNearFieldManager)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QNearFieldManager_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QNearFieldManager) QNearFieldManagerPTR() *QNearFieldManager {
	return ptr
}

//QNearFieldManager::TargetAccessMode
type QNearFieldManager__TargetAccessMode int

var (
	QNearFieldManager__NoTargetAccess              = QNearFieldManager__TargetAccessMode(0x00)
	QNearFieldManager__NdefReadTargetAccess        = QNearFieldManager__TargetAccessMode(0x01)
	QNearFieldManager__NdefWriteTargetAccess       = QNearFieldManager__TargetAccessMode(0x02)
	QNearFieldManager__TagTypeSpecificTargetAccess = QNearFieldManager__TargetAccessMode(0x04)
)

func (ptr *QNearFieldManager) RegisterNdefMessageHandler(object core.QObjectITF, method string) int {
	if ptr.Pointer() != nil {
		return int(C.QNearFieldManager_RegisterNdefMessageHandler(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQObject(object)), C.CString(method)))
	}
	return 0
}

func (ptr *QNearFieldManager) StartTargetDetection() bool {
	if ptr.Pointer() != nil {
		return C.QNearFieldManager_StartTargetDetection(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func NewQNearFieldManager(parent core.QObjectITF) *QNearFieldManager {
	return QNearFieldManagerFromPointer(unsafe.Pointer(C.QNearFieldManager_NewQNearFieldManager(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QNearFieldManager) IsAvailable() bool {
	if ptr.Pointer() != nil {
		return C.QNearFieldManager_IsAvailable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNearFieldManager) RegisterNdefMessageHandler2(typeNameFormat QNdefRecord__TypeNameFormat, ty core.QByteArrayITF, object core.QObjectITF, method string) int {
	if ptr.Pointer() != nil {
		return int(C.QNearFieldManager_RegisterNdefMessageHandler2(C.QtObjectPtr(ptr.Pointer()), C.int(typeNameFormat), C.QtObjectPtr(core.PointerFromQByteArray(ty)), C.QtObjectPtr(core.PointerFromQObject(object)), C.CString(method)))
	}
	return 0
}

func (ptr *QNearFieldManager) RegisterNdefMessageHandler3(filter QNdefFilterITF, object core.QObjectITF, method string) int {
	if ptr.Pointer() != nil {
		return int(C.QNearFieldManager_RegisterNdefMessageHandler3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQNdefFilter(filter)), C.QtObjectPtr(core.PointerFromQObject(object)), C.CString(method)))
	}
	return 0
}

func (ptr *QNearFieldManager) SetTargetAccessModes(accessModes QNearFieldManager__TargetAccessMode) {
	if ptr.Pointer() != nil {
		C.QNearFieldManager_SetTargetAccessModes(C.QtObjectPtr(ptr.Pointer()), C.int(accessModes))
	}
}

func (ptr *QNearFieldManager) StopTargetDetection() {
	if ptr.Pointer() != nil {
		C.QNearFieldManager_StopTargetDetection(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QNearFieldManager) TargetAccessModes() QNearFieldManager__TargetAccessMode {
	if ptr.Pointer() != nil {
		return QNearFieldManager__TargetAccessMode(C.QNearFieldManager_TargetAccessModes(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNearFieldManager) ConnectTargetDetected(f func(target QNearFieldTargetITF)) {
	if ptr.Pointer() != nil {
		C.QNearFieldManager_ConnectTargetDetected(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "targetDetected", f)
	}
}

func (ptr *QNearFieldManager) DisconnectTargetDetected() {
	if ptr.Pointer() != nil {
		C.QNearFieldManager_DisconnectTargetDetected(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "targetDetected")
	}
}

//export callbackQNearFieldManagerTargetDetected
func callbackQNearFieldManagerTargetDetected(ptrName *C.char, target unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "targetDetected").(func(*QNearFieldTarget))(QNearFieldTargetFromPointer(target))
}

func (ptr *QNearFieldManager) ConnectTargetLost(f func(target QNearFieldTargetITF)) {
	if ptr.Pointer() != nil {
		C.QNearFieldManager_ConnectTargetLost(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "targetLost", f)
	}
}

func (ptr *QNearFieldManager) DisconnectTargetLost() {
	if ptr.Pointer() != nil {
		C.QNearFieldManager_DisconnectTargetLost(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "targetLost")
	}
}

//export callbackQNearFieldManagerTargetLost
func callbackQNearFieldManagerTargetLost(ptrName *C.char, target unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "targetLost").(func(*QNearFieldTarget))(QNearFieldTargetFromPointer(target))
}

func (ptr *QNearFieldManager) UnregisterNdefMessageHandler(handlerId int) bool {
	if ptr.Pointer() != nil {
		return C.QNearFieldManager_UnregisterNdefMessageHandler(C.QtObjectPtr(ptr.Pointer()), C.int(handlerId)) != 0
	}
	return false
}

func (ptr *QNearFieldManager) DestroyQNearFieldManager() {
	if ptr.Pointer() != nil {
		C.QNearFieldManager_DestroyQNearFieldManager(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
