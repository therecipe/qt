package nfc

//#include "qnearfieldsharetarget.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QNearFieldShareTarget struct {
	core.QObject
}

type QNearFieldShareTargetITF interface {
	core.QObjectITF
	QNearFieldShareTargetPTR() *QNearFieldShareTarget
}

func PointerFromQNearFieldShareTarget(ptr QNearFieldShareTargetITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNearFieldShareTargetPTR().Pointer()
	}
	return nil
}

func QNearFieldShareTargetFromPointer(ptr unsafe.Pointer) *QNearFieldShareTarget {
	var n = new(QNearFieldShareTarget)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QNearFieldShareTarget_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QNearFieldShareTarget) QNearFieldShareTargetPTR() *QNearFieldShareTarget {
	return ptr
}

func (ptr *QNearFieldShareTarget) Cancel() {
	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_Cancel(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QNearFieldShareTarget) ConnectError(f func(error QNearFieldShareManager__ShareError)) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_ConnectError(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "error", f)
	}
}

func (ptr *QNearFieldShareTarget) DisconnectError() {
	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_DisconnectError(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "error")
	}
}

//export callbackQNearFieldShareTargetError
func callbackQNearFieldShareTargetError(ptrName *C.char, error C.int) {
	qt.GetSignal(C.GoString(ptrName), "error").(func(QNearFieldShareManager__ShareError))(QNearFieldShareManager__ShareError(error))
}

func (ptr *QNearFieldShareTarget) IsShareInProgress() bool {
	if ptr.Pointer() != nil {
		return C.QNearFieldShareTarget_IsShareInProgress(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNearFieldShareTarget) Share(message QNdefMessageITF) bool {
	if ptr.Pointer() != nil {
		return C.QNearFieldShareTarget_Share(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQNdefMessage(message))) != 0
	}
	return false
}

func (ptr *QNearFieldShareTarget) ShareError() QNearFieldShareManager__ShareError {
	if ptr.Pointer() != nil {
		return QNearFieldShareManager__ShareError(C.QNearFieldShareTarget_ShareError(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNearFieldShareTarget) ConnectShareFinished(f func()) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_ConnectShareFinished(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "shareFinished", f)
	}
}

func (ptr *QNearFieldShareTarget) DisconnectShareFinished() {
	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_DisconnectShareFinished(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "shareFinished")
	}
}

//export callbackQNearFieldShareTargetShareFinished
func callbackQNearFieldShareTargetShareFinished(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "shareFinished").(func())()
}

func (ptr *QNearFieldShareTarget) ShareModes() QNearFieldShareManager__ShareMode {
	if ptr.Pointer() != nil {
		return QNearFieldShareManager__ShareMode(C.QNearFieldShareTarget_ShareModes(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNearFieldShareTarget) DestroyQNearFieldShareTarget() {
	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_DestroyQNearFieldShareTarget(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
