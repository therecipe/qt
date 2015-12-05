package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"unsafe"
)

type QCameraLocksControl struct {
	QMediaControl
}

type QCameraLocksControl_ITF interface {
	QMediaControl_ITF
	QCameraLocksControl_PTR() *QCameraLocksControl
}

func PointerFromQCameraLocksControl(ptr QCameraLocksControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraLocksControl_PTR().Pointer()
	}
	return nil
}

func NewQCameraLocksControlFromPointer(ptr unsafe.Pointer) *QCameraLocksControl {
	var n = new(QCameraLocksControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QCameraLocksControl_") {
		n.SetObjectName("QCameraLocksControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QCameraLocksControl) QCameraLocksControl_PTR() *QCameraLocksControl {
	return ptr
}

func (ptr *QCameraLocksControl) LockStatus(lock QCamera__LockType) QCamera__LockStatus {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraLocksControl::lockStatus")
		}
	}()

	if ptr.Pointer() != nil {
		return QCamera__LockStatus(C.QCameraLocksControl_LockStatus(ptr.Pointer(), C.int(lock)))
	}
	return 0
}

func (ptr *QCameraLocksControl) ConnectLockStatusChanged(f func(lock QCamera__LockType, status QCamera__LockStatus, reason QCamera__LockChangeReason)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraLocksControl::lockStatusChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraLocksControl_ConnectLockStatusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "lockStatusChanged", f)
	}
}

func (ptr *QCameraLocksControl) DisconnectLockStatusChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraLocksControl::lockStatusChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraLocksControl_DisconnectLockStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "lockStatusChanged")
	}
}

//export callbackQCameraLocksControlLockStatusChanged
func callbackQCameraLocksControlLockStatusChanged(ptrName *C.char, lock C.int, status C.int, reason C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraLocksControl::lockStatusChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "lockStatusChanged").(func(QCamera__LockType, QCamera__LockStatus, QCamera__LockChangeReason))(QCamera__LockType(lock), QCamera__LockStatus(status), QCamera__LockChangeReason(reason))
}

func (ptr *QCameraLocksControl) SearchAndLock(locks QCamera__LockType) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraLocksControl::searchAndLock")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraLocksControl_SearchAndLock(ptr.Pointer(), C.int(locks))
	}
}

func (ptr *QCameraLocksControl) SupportedLocks() QCamera__LockType {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraLocksControl::supportedLocks")
		}
	}()

	if ptr.Pointer() != nil {
		return QCamera__LockType(C.QCameraLocksControl_SupportedLocks(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraLocksControl) Unlock(locks QCamera__LockType) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraLocksControl::unlock")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraLocksControl_Unlock(ptr.Pointer(), C.int(locks))
	}
}

func (ptr *QCameraLocksControl) DestroyQCameraLocksControl() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraLocksControl::~QCameraLocksControl")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraLocksControl_DestroyQCameraLocksControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
