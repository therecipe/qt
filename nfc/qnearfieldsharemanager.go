package nfc

//#include "qnearfieldsharemanager.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QNearFieldShareManager struct {
	core.QObject
}

type QNearFieldShareManagerITF interface {
	core.QObjectITF
	QNearFieldShareManagerPTR() *QNearFieldShareManager
}

func PointerFromQNearFieldShareManager(ptr QNearFieldShareManagerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNearFieldShareManagerPTR().Pointer()
	}
	return nil
}

func QNearFieldShareManagerFromPointer(ptr unsafe.Pointer) *QNearFieldShareManager {
	var n = new(QNearFieldShareManager)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QNearFieldShareManager_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QNearFieldShareManager) QNearFieldShareManagerPTR() *QNearFieldShareManager {
	return ptr
}

//QNearFieldShareManager::ShareError
type QNearFieldShareManager__ShareError int

var (
	QNearFieldShareManager__NoError                     = QNearFieldShareManager__ShareError(0)
	QNearFieldShareManager__UnknownError                = QNearFieldShareManager__ShareError(1)
	QNearFieldShareManager__InvalidShareContentError    = QNearFieldShareManager__ShareError(2)
	QNearFieldShareManager__ShareCanceledError          = QNearFieldShareManager__ShareError(3)
	QNearFieldShareManager__ShareInterruptedError       = QNearFieldShareManager__ShareError(4)
	QNearFieldShareManager__ShareRejectedError          = QNearFieldShareManager__ShareError(5)
	QNearFieldShareManager__UnsupportedShareModeError   = QNearFieldShareManager__ShareError(6)
	QNearFieldShareManager__ShareAlreadyInProgressError = QNearFieldShareManager__ShareError(7)
	QNearFieldShareManager__SharePermissionDeniedError  = QNearFieldShareManager__ShareError(8)
)

//QNearFieldShareManager::ShareMode
type QNearFieldShareManager__ShareMode int

var (
	QNearFieldShareManager__NoShare   = QNearFieldShareManager__ShareMode(0x00)
	QNearFieldShareManager__NdefShare = QNearFieldShareManager__ShareMode(0x01)
	QNearFieldShareManager__FileShare = QNearFieldShareManager__ShareMode(0x02)
)

func NewQNearFieldShareManager(parent core.QObjectITF) *QNearFieldShareManager {
	return QNearFieldShareManagerFromPointer(unsafe.Pointer(C.QNearFieldShareManager_NewQNearFieldShareManager(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QNearFieldShareManager) ConnectError(f func(error QNearFieldShareManager__ShareError)) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_ConnectError(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "error", f)
	}
}

func (ptr *QNearFieldShareManager) DisconnectError() {
	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_DisconnectError(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "error")
	}
}

//export callbackQNearFieldShareManagerError
func callbackQNearFieldShareManagerError(ptrName *C.char, error C.int) {
	qt.GetSignal(C.GoString(ptrName), "error").(func(QNearFieldShareManager__ShareError))(QNearFieldShareManager__ShareError(error))
}

func (ptr *QNearFieldShareManager) SetShareModes(mode QNearFieldShareManager__ShareMode) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_SetShareModes(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QNearFieldShareManager) ShareError() QNearFieldShareManager__ShareError {
	if ptr.Pointer() != nil {
		return QNearFieldShareManager__ShareError(C.QNearFieldShareManager_ShareError(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNearFieldShareManager) ShareModes() QNearFieldShareManager__ShareMode {
	if ptr.Pointer() != nil {
		return QNearFieldShareManager__ShareMode(C.QNearFieldShareManager_ShareModes(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNearFieldShareManager) ConnectShareModesChanged(f func(modes QNearFieldShareManager__ShareMode)) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_ConnectShareModesChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "shareModesChanged", f)
	}
}

func (ptr *QNearFieldShareManager) DisconnectShareModesChanged() {
	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_DisconnectShareModesChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "shareModesChanged")
	}
}

//export callbackQNearFieldShareManagerShareModesChanged
func callbackQNearFieldShareManagerShareModesChanged(ptrName *C.char, modes C.int) {
	qt.GetSignal(C.GoString(ptrName), "shareModesChanged").(func(QNearFieldShareManager__ShareMode))(QNearFieldShareManager__ShareMode(modes))
}

func QNearFieldShareManager_SupportedShareModes() QNearFieldShareManager__ShareMode {
	return QNearFieldShareManager__ShareMode(C.QNearFieldShareManager_QNearFieldShareManager_SupportedShareModes())
}

func (ptr *QNearFieldShareManager) ConnectTargetDetected(f func(shareTarget QNearFieldShareTargetITF)) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_ConnectTargetDetected(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "targetDetected", f)
	}
}

func (ptr *QNearFieldShareManager) DisconnectTargetDetected() {
	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_DisconnectTargetDetected(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "targetDetected")
	}
}

//export callbackQNearFieldShareManagerTargetDetected
func callbackQNearFieldShareManagerTargetDetected(ptrName *C.char, shareTarget unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "targetDetected").(func(*QNearFieldShareTarget))(QNearFieldShareTargetFromPointer(shareTarget))
}

func (ptr *QNearFieldShareManager) DestroyQNearFieldShareManager() {
	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_DestroyQNearFieldShareManager(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
