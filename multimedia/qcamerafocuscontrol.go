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

type QCameraFocusControl_ITF interface {
	QMediaControl_ITF
	QCameraFocusControl_PTR() *QCameraFocusControl
}

func PointerFromQCameraFocusControl(ptr QCameraFocusControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraFocusControl_PTR().Pointer()
	}
	return nil
}

func NewQCameraFocusControlFromPointer(ptr unsafe.Pointer) *QCameraFocusControl {
	var n = new(QCameraFocusControl)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QCameraFocusControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QCameraFocusControl) QCameraFocusControl_PTR() *QCameraFocusControl {
	return ptr
}

func (ptr *QCameraFocusControl) FocusMode() QCameraFocus__FocusMode {
	if ptr.Pointer() != nil {
		return QCameraFocus__FocusMode(C.QCameraFocusControl_FocusMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraFocusControl) ConnectFocusModeChanged(f func(mode QCameraFocus__FocusMode)) {
	if ptr.Pointer() != nil {
		C.QCameraFocusControl_ConnectFocusModeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "focusModeChanged", f)
	}
}

func (ptr *QCameraFocusControl) DisconnectFocusModeChanged() {
	if ptr.Pointer() != nil {
		C.QCameraFocusControl_DisconnectFocusModeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "focusModeChanged")
	}
}

//export callbackQCameraFocusControlFocusModeChanged
func callbackQCameraFocusControlFocusModeChanged(ptrName *C.char, mode C.int) {
	qt.GetSignal(C.GoString(ptrName), "focusModeChanged").(func(QCameraFocus__FocusMode))(QCameraFocus__FocusMode(mode))
}

func (ptr *QCameraFocusControl) FocusPointMode() QCameraFocus__FocusPointMode {
	if ptr.Pointer() != nil {
		return QCameraFocus__FocusPointMode(C.QCameraFocusControl_FocusPointMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraFocusControl) ConnectFocusPointModeChanged(f func(mode QCameraFocus__FocusPointMode)) {
	if ptr.Pointer() != nil {
		C.QCameraFocusControl_ConnectFocusPointModeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "focusPointModeChanged", f)
	}
}

func (ptr *QCameraFocusControl) DisconnectFocusPointModeChanged() {
	if ptr.Pointer() != nil {
		C.QCameraFocusControl_DisconnectFocusPointModeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "focusPointModeChanged")
	}
}

//export callbackQCameraFocusControlFocusPointModeChanged
func callbackQCameraFocusControlFocusPointModeChanged(ptrName *C.char, mode C.int) {
	qt.GetSignal(C.GoString(ptrName), "focusPointModeChanged").(func(QCameraFocus__FocusPointMode))(QCameraFocus__FocusPointMode(mode))
}

func (ptr *QCameraFocusControl) ConnectFocusZonesChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QCameraFocusControl_ConnectFocusZonesChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "focusZonesChanged", f)
	}
}

func (ptr *QCameraFocusControl) DisconnectFocusZonesChanged() {
	if ptr.Pointer() != nil {
		C.QCameraFocusControl_DisconnectFocusZonesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "focusZonesChanged")
	}
}

//export callbackQCameraFocusControlFocusZonesChanged
func callbackQCameraFocusControlFocusZonesChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "focusZonesChanged").(func())()
}

func (ptr *QCameraFocusControl) IsFocusModeSupported(mode QCameraFocus__FocusMode) bool {
	if ptr.Pointer() != nil {
		return C.QCameraFocusControl_IsFocusModeSupported(ptr.Pointer(), C.int(mode)) != 0
	}
	return false
}

func (ptr *QCameraFocusControl) IsFocusPointModeSupported(mode QCameraFocus__FocusPointMode) bool {
	if ptr.Pointer() != nil {
		return C.QCameraFocusControl_IsFocusPointModeSupported(ptr.Pointer(), C.int(mode)) != 0
	}
	return false
}

func (ptr *QCameraFocusControl) SetCustomFocusPoint(point core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QCameraFocusControl_SetCustomFocusPoint(ptr.Pointer(), core.PointerFromQPointF(point))
	}
}

func (ptr *QCameraFocusControl) SetFocusMode(mode QCameraFocus__FocusMode) {
	if ptr.Pointer() != nil {
		C.QCameraFocusControl_SetFocusMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QCameraFocusControl) SetFocusPointMode(mode QCameraFocus__FocusPointMode) {
	if ptr.Pointer() != nil {
		C.QCameraFocusControl_SetFocusPointMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QCameraFocusControl) DestroyQCameraFocusControl() {
	if ptr.Pointer() != nil {
		C.QCameraFocusControl_DestroyQCameraFocusControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
