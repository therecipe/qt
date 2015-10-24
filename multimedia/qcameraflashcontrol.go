package multimedia

//#include "qcameraflashcontrol.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QCameraFlashControl struct {
	QMediaControl
}

type QCameraFlashControlITF interface {
	QMediaControlITF
	QCameraFlashControlPTR() *QCameraFlashControl
}

func PointerFromQCameraFlashControl(ptr QCameraFlashControlITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraFlashControlPTR().Pointer()
	}
	return nil
}

func QCameraFlashControlFromPointer(ptr unsafe.Pointer) *QCameraFlashControl {
	var n = new(QCameraFlashControl)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QCameraFlashControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QCameraFlashControl) QCameraFlashControlPTR() *QCameraFlashControl {
	return ptr
}

func (ptr *QCameraFlashControl) FlashMode() QCameraExposure__FlashMode {
	if ptr.Pointer() != nil {
		return QCameraExposure__FlashMode(C.QCameraFlashControl_FlashMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCameraFlashControl) ConnectFlashReady(f func(ready bool)) {
	if ptr.Pointer() != nil {
		C.QCameraFlashControl_ConnectFlashReady(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "flashReady", f)
	}
}

func (ptr *QCameraFlashControl) DisconnectFlashReady() {
	if ptr.Pointer() != nil {
		C.QCameraFlashControl_DisconnectFlashReady(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "flashReady")
	}
}

//export callbackQCameraFlashControlFlashReady
func callbackQCameraFlashControlFlashReady(ptrName *C.char, ready C.int) {
	qt.GetSignal(C.GoString(ptrName), "flashReady").(func(bool))(int(ready) != 0)
}

func (ptr *QCameraFlashControl) IsFlashModeSupported(mode QCameraExposure__FlashMode) bool {
	if ptr.Pointer() != nil {
		return C.QCameraFlashControl_IsFlashModeSupported(C.QtObjectPtr(ptr.Pointer()), C.int(mode)) != 0
	}
	return false
}

func (ptr *QCameraFlashControl) IsFlashReady() bool {
	if ptr.Pointer() != nil {
		return C.QCameraFlashControl_IsFlashReady(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QCameraFlashControl) SetFlashMode(mode QCameraExposure__FlashMode) {
	if ptr.Pointer() != nil {
		C.QCameraFlashControl_SetFlashMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QCameraFlashControl) DestroyQCameraFlashControl() {
	if ptr.Pointer() != nil {
		C.QCameraFlashControl_DestroyQCameraFlashControl(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
