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

type QNearFieldShareTarget_ITF interface {
	core.QObject_ITF
	QNearFieldShareTarget_PTR() *QNearFieldShareTarget
}

func PointerFromQNearFieldShareTarget(ptr QNearFieldShareTarget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNearFieldShareTarget_PTR().Pointer()
	}
	return nil
}

func NewQNearFieldShareTargetFromPointer(ptr unsafe.Pointer) *QNearFieldShareTarget {
	var n = new(QNearFieldShareTarget)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QNearFieldShareTarget_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QNearFieldShareTarget) QNearFieldShareTarget_PTR() *QNearFieldShareTarget {
	return ptr
}

func (ptr *QNearFieldShareTarget) Cancel() {
	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_Cancel(ptr.Pointer())
	}
}

func (ptr *QNearFieldShareTarget) ConnectError(f func(error QNearFieldShareManager__ShareError)) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_ConnectError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error", f)
	}
}

func (ptr *QNearFieldShareTarget) DisconnectError() {
	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_DisconnectError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error")
	}
}

//export callbackQNearFieldShareTargetError
func callbackQNearFieldShareTargetError(ptrName *C.char, error C.int) {
	qt.GetSignal(C.GoString(ptrName), "error").(func(QNearFieldShareManager__ShareError))(QNearFieldShareManager__ShareError(error))
}

func (ptr *QNearFieldShareTarget) IsShareInProgress() bool {
	if ptr.Pointer() != nil {
		return C.QNearFieldShareTarget_IsShareInProgress(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNearFieldShareTarget) Share(message QNdefMessage_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNearFieldShareTarget_Share(ptr.Pointer(), PointerFromQNdefMessage(message)) != 0
	}
	return false
}

func (ptr *QNearFieldShareTarget) ShareError() QNearFieldShareManager__ShareError {
	if ptr.Pointer() != nil {
		return QNearFieldShareManager__ShareError(C.QNearFieldShareTarget_ShareError(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNearFieldShareTarget) ConnectShareFinished(f func()) {
	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_ConnectShareFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "shareFinished", f)
	}
}

func (ptr *QNearFieldShareTarget) DisconnectShareFinished() {
	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_DisconnectShareFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "shareFinished")
	}
}

//export callbackQNearFieldShareTargetShareFinished
func callbackQNearFieldShareTargetShareFinished(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "shareFinished").(func())()
}

func (ptr *QNearFieldShareTarget) ShareModes() QNearFieldShareManager__ShareMode {
	if ptr.Pointer() != nil {
		return QNearFieldShareManager__ShareMode(C.QNearFieldShareTarget_ShareModes(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNearFieldShareTarget) DestroyQNearFieldShareTarget() {
	if ptr.Pointer() != nil {
		C.QNearFieldShareTarget_DestroyQNearFieldShareTarget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
