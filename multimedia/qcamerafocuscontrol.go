package multimedia

//#include "qcamerafocuscontrol.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QCameraFocusControl struct {
	QMediaControl
}

type QCameraFocusControlITF interface {
	QMediaControlITF
	QCameraFocusControlPTR() *QCameraFocusControl
}

func PointerFromQCameraFocusControl(ptr QCameraFocusControlITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraFocusControlPTR().Pointer()
	}
	return nil
}

func QCameraFocusControlFromPointer(ptr unsafe.Pointer) *QCameraFocusControl {
	var n = new(QCameraFocusControl)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QCameraFocusControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QCameraFocusControl) QCameraFocusControlPTR() *QCameraFocusControl {
	return ptr
}

func (ptr *QCameraFocusControl) FocusMode() QCameraFocus__FocusMode {
	if ptr.Pointer() != nil {
		return QCameraFocus__FocusMode(C.QCameraFocusControl_FocusMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCameraFocusControl) ConnectFocusModeChanged(f func(mode QCameraFocus__FocusMode)) {
	if ptr.Pointer() != nil {
		C.QCameraFocusControl_ConnectFocusModeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "focusModeChanged", f)
	}
}

func (ptr *QCameraFocusControl) DisconnectFocusModeChanged() {
	if ptr.Pointer() != nil {
		C.QCameraFocusControl_DisconnectFocusModeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "focusModeChanged")
	}
}

//export callbackQCameraFocusControlFocusModeChanged
func callbackQCameraFocusControlFocusModeChanged(ptrName *C.char, mode C.int) {
	qt.GetSignal(C.GoString(ptrName), "focusModeChanged").(func(QCameraFocus__FocusMode))(QCameraFocus__FocusMode(mode))
}

func (ptr *QCameraFocusControl) FocusPointMode() QCameraFocus__FocusPointMode {
	if ptr.Pointer() != nil {
		return QCameraFocus__FocusPointMode(C.QCameraFocusControl_FocusPointMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCameraFocusControl) ConnectFocusPointModeChanged(f func(mode QCameraFocus__FocusPointMode)) {
	if ptr.Pointer() != nil {
		C.QCameraFocusControl_ConnectFocusPointModeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "focusPointModeChanged", f)
	}
}

func (ptr *QCameraFocusControl) DisconnectFocusPointModeChanged() {
	if ptr.Pointer() != nil {
		C.QCameraFocusControl_DisconnectFocusPointModeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "focusPointModeChanged")
	}
}

//export callbackQCameraFocusControlFocusPointModeChanged
func callbackQCameraFocusControlFocusPointModeChanged(ptrName *C.char, mode C.int) {
	qt.GetSignal(C.GoString(ptrName), "focusPointModeChanged").(func(QCameraFocus__FocusPointMode))(QCameraFocus__FocusPointMode(mode))
}

func (ptr *QCameraFocusControl) ConnectFocusZonesChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QCameraFocusControl_ConnectFocusZonesChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "focusZonesChanged", f)
	}
}

func (ptr *QCameraFocusControl) DisconnectFocusZonesChanged() {
	if ptr.Pointer() != nil {
		C.QCameraFocusControl_DisconnectFocusZonesChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "focusZonesChanged")
	}
}

//export callbackQCameraFocusControlFocusZonesChanged
func callbackQCameraFocusControlFocusZonesChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "focusZonesChanged").(func())()
}

func (ptr *QCameraFocusControl) IsFocusModeSupported(mode QCameraFocus__FocusMode) bool {
	if ptr.Pointer() != nil {
		return C.QCameraFocusControl_IsFocusModeSupported(C.QtObjectPtr(ptr.Pointer()), C.int(mode)) != 0
	}
	return false
}

func (ptr *QCameraFocusControl) IsFocusPointModeSupported(mode QCameraFocus__FocusPointMode) bool {
	if ptr.Pointer() != nil {
		return C.QCameraFocusControl_IsFocusPointModeSupported(C.QtObjectPtr(ptr.Pointer()), C.int(mode)) != 0
	}
	return false
}

func (ptr *QCameraFocusControl) SetCustomFocusPoint(point core.QPointFITF) {
	if ptr.Pointer() != nil {
		C.QCameraFocusControl_SetCustomFocusPoint(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(point)))
	}
}

func (ptr *QCameraFocusControl) SetFocusMode(mode QCameraFocus__FocusMode) {
	if ptr.Pointer() != nil {
		C.QCameraFocusControl_SetFocusMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QCameraFocusControl) SetFocusPointMode(mode QCameraFocus__FocusPointMode) {
	if ptr.Pointer() != nil {
		C.QCameraFocusControl_SetFocusPointMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QCameraFocusControl) DestroyQCameraFocusControl() {
	if ptr.Pointer() != nil {
		C.QCameraFocusControl_DestroyQCameraFocusControl(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
