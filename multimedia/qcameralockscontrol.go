package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
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
		n.SetObjectName("QCameraLocksControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QCameraLocksControl) QCameraLocksControl_PTR() *QCameraLocksControl {
	return ptr
}

func (ptr *QCameraLocksControl) LockStatus(lock QCamera__LockType) QCamera__LockStatus {
	defer qt.Recovering("QCameraLocksControl::lockStatus")

	if ptr.Pointer() != nil {
		return QCamera__LockStatus(C.QCameraLocksControl_LockStatus(ptr.Pointer(), C.int(lock)))
	}
	return 0
}

func (ptr *QCameraLocksControl) ConnectLockStatusChanged(f func(lock QCamera__LockType, status QCamera__LockStatus, reason QCamera__LockChangeReason)) {
	defer qt.Recovering("connect QCameraLocksControl::lockStatusChanged")

	if ptr.Pointer() != nil {
		C.QCameraLocksControl_ConnectLockStatusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "lockStatusChanged", f)
	}
}

func (ptr *QCameraLocksControl) DisconnectLockStatusChanged() {
	defer qt.Recovering("disconnect QCameraLocksControl::lockStatusChanged")

	if ptr.Pointer() != nil {
		C.QCameraLocksControl_DisconnectLockStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "lockStatusChanged")
	}
}

//export callbackQCameraLocksControlLockStatusChanged
func callbackQCameraLocksControlLockStatusChanged(ptrName *C.char, lock C.int, status C.int, reason C.int) {
	defer qt.Recovering("callback QCameraLocksControl::lockStatusChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "lockStatusChanged")
	if signal != nil {
		signal.(func(QCamera__LockType, QCamera__LockStatus, QCamera__LockChangeReason))(QCamera__LockType(lock), QCamera__LockStatus(status), QCamera__LockChangeReason(reason))
	}

}

func (ptr *QCameraLocksControl) SearchAndLock(locks QCamera__LockType) {
	defer qt.Recovering("QCameraLocksControl::searchAndLock")

	if ptr.Pointer() != nil {
		C.QCameraLocksControl_SearchAndLock(ptr.Pointer(), C.int(locks))
	}
}

func (ptr *QCameraLocksControl) SupportedLocks() QCamera__LockType {
	defer qt.Recovering("QCameraLocksControl::supportedLocks")

	if ptr.Pointer() != nil {
		return QCamera__LockType(C.QCameraLocksControl_SupportedLocks(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraLocksControl) Unlock(locks QCamera__LockType) {
	defer qt.Recovering("QCameraLocksControl::unlock")

	if ptr.Pointer() != nil {
		C.QCameraLocksControl_Unlock(ptr.Pointer(), C.int(locks))
	}
}

func (ptr *QCameraLocksControl) DestroyQCameraLocksControl() {
	defer qt.Recovering("QCameraLocksControl::~QCameraLocksControl")

	if ptr.Pointer() != nil {
		C.QCameraLocksControl_DestroyQCameraLocksControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
