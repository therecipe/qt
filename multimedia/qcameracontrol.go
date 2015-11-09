package multimedia

//#include "qcameracontrol.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QCameraControl struct {
	QMediaControl
}

type QCameraControl_ITF interface {
	QMediaControl_ITF
	QCameraControl_PTR() *QCameraControl
}

func PointerFromQCameraControl(ptr QCameraControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraControl_PTR().Pointer()
	}
	return nil
}

func NewQCameraControlFromPointer(ptr unsafe.Pointer) *QCameraControl {
	var n = new(QCameraControl)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QCameraControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QCameraControl) QCameraControl_PTR() *QCameraControl {
	return ptr
}

//QCameraControl::PropertyChangeType
type QCameraControl__PropertyChangeType int64

const (
	QCameraControl__CaptureMode           = QCameraControl__PropertyChangeType(1)
	QCameraControl__ImageEncodingSettings = QCameraControl__PropertyChangeType(2)
	QCameraControl__VideoEncodingSettings = QCameraControl__PropertyChangeType(3)
	QCameraControl__Viewfinder            = QCameraControl__PropertyChangeType(4)
	QCameraControl__ViewfinderSettings    = QCameraControl__PropertyChangeType(5)
)

func (ptr *QCameraControl) CanChangeProperty(changeType QCameraControl__PropertyChangeType, status QCamera__Status) bool {
	if ptr.Pointer() != nil {
		return C.QCameraControl_CanChangeProperty(ptr.Pointer(), C.int(changeType), C.int(status)) != 0
	}
	return false
}

func (ptr *QCameraControl) CaptureMode() QCamera__CaptureMode {
	if ptr.Pointer() != nil {
		return QCamera__CaptureMode(C.QCameraControl_CaptureMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraControl) ConnectCaptureModeChanged(f func(mode QCamera__CaptureMode)) {
	if ptr.Pointer() != nil {
		C.QCameraControl_ConnectCaptureModeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "captureModeChanged", f)
	}
}

func (ptr *QCameraControl) DisconnectCaptureModeChanged() {
	if ptr.Pointer() != nil {
		C.QCameraControl_DisconnectCaptureModeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "captureModeChanged")
	}
}

//export callbackQCameraControlCaptureModeChanged
func callbackQCameraControlCaptureModeChanged(ptrName *C.char, mode C.int) {
	qt.GetSignal(C.GoString(ptrName), "captureModeChanged").(func(QCamera__CaptureMode))(QCamera__CaptureMode(mode))
}

func (ptr *QCameraControl) ConnectError(f func(error int, errorString string)) {
	if ptr.Pointer() != nil {
		C.QCameraControl_ConnectError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error", f)
	}
}

func (ptr *QCameraControl) DisconnectError() {
	if ptr.Pointer() != nil {
		C.QCameraControl_DisconnectError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error")
	}
}

//export callbackQCameraControlError
func callbackQCameraControlError(ptrName *C.char, error C.int, errorString *C.char) {
	qt.GetSignal(C.GoString(ptrName), "error").(func(int, string))(int(error), C.GoString(errorString))
}

func (ptr *QCameraControl) IsCaptureModeSupported(mode QCamera__CaptureMode) bool {
	if ptr.Pointer() != nil {
		return C.QCameraControl_IsCaptureModeSupported(ptr.Pointer(), C.int(mode)) != 0
	}
	return false
}

func (ptr *QCameraControl) SetCaptureMode(mode QCamera__CaptureMode) {
	if ptr.Pointer() != nil {
		C.QCameraControl_SetCaptureMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QCameraControl) SetState(state QCamera__State) {
	if ptr.Pointer() != nil {
		C.QCameraControl_SetState(ptr.Pointer(), C.int(state))
	}
}

func (ptr *QCameraControl) ConnectStateChanged(f func(state QCamera__State)) {
	if ptr.Pointer() != nil {
		C.QCameraControl_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QCameraControl) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QCameraControl_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQCameraControlStateChanged
func callbackQCameraControlStateChanged(ptrName *C.char, state C.int) {
	qt.GetSignal(C.GoString(ptrName), "stateChanged").(func(QCamera__State))(QCamera__State(state))
}

func (ptr *QCameraControl) Status() QCamera__Status {
	if ptr.Pointer() != nil {
		return QCamera__Status(C.QCameraControl_Status(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraControl) ConnectStatusChanged(f func(status QCamera__Status)) {
	if ptr.Pointer() != nil {
		C.QCameraControl_ConnectStatusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "statusChanged", f)
	}
}

func (ptr *QCameraControl) DisconnectStatusChanged() {
	if ptr.Pointer() != nil {
		C.QCameraControl_DisconnectStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "statusChanged")
	}
}

//export callbackQCameraControlStatusChanged
func callbackQCameraControlStatusChanged(ptrName *C.char, status C.int) {
	qt.GetSignal(C.GoString(ptrName), "statusChanged").(func(QCamera__Status))(QCamera__Status(status))
}

func (ptr *QCameraControl) DestroyQCameraControl() {
	if ptr.Pointer() != nil {
		C.QCameraControl_DestroyQCameraControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
