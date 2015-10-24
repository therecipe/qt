package multimedia

//#include "qcameralockscontrol.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QCameraLocksControl struct {
	QMediaControl
}

type QCameraLocksControlITF interface {
	QMediaControlITF
	QCameraLocksControlPTR() *QCameraLocksControl
}

func PointerFromQCameraLocksControl(ptr QCameraLocksControlITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraLocksControlPTR().Pointer()
	}
	return nil
}

func QCameraLocksControlFromPointer(ptr unsafe.Pointer) *QCameraLocksControl {
	var n = new(QCameraLocksControl)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QCameraLocksControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QCameraLocksControl) QCameraLocksControlPTR() *QCameraLocksControl {
	return ptr
}

func (ptr *QCameraLocksControl) LockStatus(lock QCamera__LockType) QCamera__LockStatus {
	if ptr.Pointer() != nil {
		return QCamera__LockStatus(C.QCameraLocksControl_LockStatus(C.QtObjectPtr(ptr.Pointer()), C.int(lock)))
	}
	return 0
}

func (ptr *QCameraLocksControl) ConnectLockStatusChanged(f func(lock QCamera__LockType, status QCamera__LockStatus, reason QCamera__LockChangeReason)) {
	if ptr.Pointer() != nil {
		C.QCameraLocksControl_ConnectLockStatusChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "lockStatusChanged", f)
	}
}

func (ptr *QCameraLocksControl) DisconnectLockStatusChanged() {
	if ptr.Pointer() != nil {
		C.QCameraLocksControl_DisconnectLockStatusChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "lockStatusChanged")
	}
}

//export callbackQCameraLocksControlLockStatusChanged
func callbackQCameraLocksControlLockStatusChanged(ptrName *C.char, lock C.int, status C.int, reason C.int) {
	qt.GetSignal(C.GoString(ptrName), "lockStatusChanged").(func(QCamera__LockType, QCamera__LockStatus, QCamera__LockChangeReason))(QCamera__LockType(lock), QCamera__LockStatus(status), QCamera__LockChangeReason(reason))
}

func (ptr *QCameraLocksControl) SearchAndLock(locks QCamera__LockType) {
	if ptr.Pointer() != nil {
		C.QCameraLocksControl_SearchAndLock(C.QtObjectPtr(ptr.Pointer()), C.int(locks))
	}
}

func (ptr *QCameraLocksControl) SupportedLocks() QCamera__LockType {
	if ptr.Pointer() != nil {
		return QCamera__LockType(C.QCameraLocksControl_SupportedLocks(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCameraLocksControl) Unlock(locks QCamera__LockType) {
	if ptr.Pointer() != nil {
		C.QCameraLocksControl_Unlock(C.QtObjectPtr(ptr.Pointer()), C.int(locks))
	}
}

func (ptr *QCameraLocksControl) DestroyQCameraLocksControl() {
	if ptr.Pointer() != nil {
		C.QCameraLocksControl_DestroyQCameraLocksControl(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
