package nfc

//#include "nfc.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QNearFieldShareManager struct {
	core.QObject
}

type QNearFieldShareManager_ITF interface {
	core.QObject_ITF
	QNearFieldShareManager_PTR() *QNearFieldShareManager
}

func PointerFromQNearFieldShareManager(ptr QNearFieldShareManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNearFieldShareManager_PTR().Pointer()
	}
	return nil
}

func NewQNearFieldShareManagerFromPointer(ptr unsafe.Pointer) *QNearFieldShareManager {
	var n = new(QNearFieldShareManager)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QNearFieldShareManager_") {
		n.SetObjectName("QNearFieldShareManager_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QNearFieldShareManager) QNearFieldShareManager_PTR() *QNearFieldShareManager {
	return ptr
}

//QNearFieldShareManager::ShareError
type QNearFieldShareManager__ShareError int64

const (
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
type QNearFieldShareManager__ShareMode int64

const (
	QNearFieldShareManager__NoShare   = QNearFieldShareManager__ShareMode(0x00)
	QNearFieldShareManager__NdefShare = QNearFieldShareManager__ShareMode(0x01)
	QNearFieldShareManager__FileShare = QNearFieldShareManager__ShareMode(0x02)
)

func NewQNearFieldShareManager(parent core.QObject_ITF) *QNearFieldShareManager {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNearFieldShareManager::QNearFieldShareManager")
		}
	}()

	return NewQNearFieldShareManagerFromPointer(C.QNearFieldShareManager_NewQNearFieldShareManager(core.PointerFromQObject(parent)))
}

func (ptr *QNearFieldShareManager) ConnectError(f func(error QNearFieldShareManager__ShareError)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNearFieldShareManager::error")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_ConnectError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error", f)
	}
}

func (ptr *QNearFieldShareManager) DisconnectError() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNearFieldShareManager::error")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_DisconnectError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error")
	}
}

//export callbackQNearFieldShareManagerError
func callbackQNearFieldShareManagerError(ptrName *C.char, error C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNearFieldShareManager::error")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "error").(func(QNearFieldShareManager__ShareError))(QNearFieldShareManager__ShareError(error))
}

func (ptr *QNearFieldShareManager) SetShareModes(mode QNearFieldShareManager__ShareMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNearFieldShareManager::setShareModes")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_SetShareModes(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QNearFieldShareManager) ShareError() QNearFieldShareManager__ShareError {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNearFieldShareManager::shareError")
		}
	}()

	if ptr.Pointer() != nil {
		return QNearFieldShareManager__ShareError(C.QNearFieldShareManager_ShareError(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNearFieldShareManager) ShareModes() QNearFieldShareManager__ShareMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNearFieldShareManager::shareModes")
		}
	}()

	if ptr.Pointer() != nil {
		return QNearFieldShareManager__ShareMode(C.QNearFieldShareManager_ShareModes(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNearFieldShareManager) ConnectShareModesChanged(f func(modes QNearFieldShareManager__ShareMode)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNearFieldShareManager::shareModesChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_ConnectShareModesChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "shareModesChanged", f)
	}
}

func (ptr *QNearFieldShareManager) DisconnectShareModesChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNearFieldShareManager::shareModesChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_DisconnectShareModesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "shareModesChanged")
	}
}

//export callbackQNearFieldShareManagerShareModesChanged
func callbackQNearFieldShareManagerShareModesChanged(ptrName *C.char, modes C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNearFieldShareManager::shareModesChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "shareModesChanged").(func(QNearFieldShareManager__ShareMode))(QNearFieldShareManager__ShareMode(modes))
}

func QNearFieldShareManager_SupportedShareModes() QNearFieldShareManager__ShareMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNearFieldShareManager::supportedShareModes")
		}
	}()

	return QNearFieldShareManager__ShareMode(C.QNearFieldShareManager_QNearFieldShareManager_SupportedShareModes())
}

func (ptr *QNearFieldShareManager) ConnectTargetDetected(f func(shareTarget *QNearFieldShareTarget)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNearFieldShareManager::targetDetected")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_ConnectTargetDetected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "targetDetected", f)
	}
}

func (ptr *QNearFieldShareManager) DisconnectTargetDetected() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNearFieldShareManager::targetDetected")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_DisconnectTargetDetected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "targetDetected")
	}
}

//export callbackQNearFieldShareManagerTargetDetected
func callbackQNearFieldShareManagerTargetDetected(ptrName *C.char, shareTarget unsafe.Pointer) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNearFieldShareManager::targetDetected")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "targetDetected").(func(*QNearFieldShareTarget))(NewQNearFieldShareTargetFromPointer(shareTarget))
}

func (ptr *QNearFieldShareManager) DestroyQNearFieldShareManager() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNearFieldShareManager::~QNearFieldShareManager")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_DestroyQNearFieldShareManager(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
