package multimedia

//#include "qcamerafocus.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QCameraFocus struct {
	core.QObject
}

type QCameraFocusITF interface {
	core.QObjectITF
	QCameraFocusPTR() *QCameraFocus
}

func PointerFromQCameraFocus(ptr QCameraFocusITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraFocusPTR().Pointer()
	}
	return nil
}

func QCameraFocusFromPointer(ptr unsafe.Pointer) *QCameraFocus {
	var n = new(QCameraFocus)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QCameraFocus_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QCameraFocus) QCameraFocusPTR() *QCameraFocus {
	return ptr
}

//QCameraFocus::FocusMode
type QCameraFocus__FocusMode int

var (
	QCameraFocus__ManualFocus     = QCameraFocus__FocusMode(0x1)
	QCameraFocus__HyperfocalFocus = QCameraFocus__FocusMode(0x02)
	QCameraFocus__InfinityFocus   = QCameraFocus__FocusMode(0x04)
	QCameraFocus__AutoFocus       = QCameraFocus__FocusMode(0x8)
	QCameraFocus__ContinuousFocus = QCameraFocus__FocusMode(0x10)
	QCameraFocus__MacroFocus      = QCameraFocus__FocusMode(0x20)
)

//QCameraFocus::FocusPointMode
type QCameraFocus__FocusPointMode int

var (
	QCameraFocus__FocusPointAuto          = QCameraFocus__FocusPointMode(0)
	QCameraFocus__FocusPointCenter        = QCameraFocus__FocusPointMode(1)
	QCameraFocus__FocusPointFaceDetection = QCameraFocus__FocusPointMode(2)
	QCameraFocus__FocusPointCustom        = QCameraFocus__FocusPointMode(3)
)

func (ptr *QCameraFocus) FocusMode() QCameraFocus__FocusMode {
	if ptr.Pointer() != nil {
		return QCameraFocus__FocusMode(C.QCameraFocus_FocusMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCameraFocus) FocusPointMode() QCameraFocus__FocusPointMode {
	if ptr.Pointer() != nil {
		return QCameraFocus__FocusPointMode(C.QCameraFocus_FocusPointMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCameraFocus) SetCustomFocusPoint(point core.QPointFITF) {
	if ptr.Pointer() != nil {
		C.QCameraFocus_SetCustomFocusPoint(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(point)))
	}
}

func (ptr *QCameraFocus) SetFocusMode(mode QCameraFocus__FocusMode) {
	if ptr.Pointer() != nil {
		C.QCameraFocus_SetFocusMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QCameraFocus) SetFocusPointMode(mode QCameraFocus__FocusPointMode) {
	if ptr.Pointer() != nil {
		C.QCameraFocus_SetFocusPointMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QCameraFocus) ConnectFocusZonesChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QCameraFocus_ConnectFocusZonesChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "focusZonesChanged", f)
	}
}

func (ptr *QCameraFocus) DisconnectFocusZonesChanged() {
	if ptr.Pointer() != nil {
		C.QCameraFocus_DisconnectFocusZonesChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "focusZonesChanged")
	}
}

//export callbackQCameraFocusFocusZonesChanged
func callbackQCameraFocusFocusZonesChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "focusZonesChanged").(func())()
}

func (ptr *QCameraFocus) IsAvailable() bool {
	if ptr.Pointer() != nil {
		return C.QCameraFocus_IsAvailable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QCameraFocus) IsFocusModeSupported(mode QCameraFocus__FocusMode) bool {
	if ptr.Pointer() != nil {
		return C.QCameraFocus_IsFocusModeSupported(C.QtObjectPtr(ptr.Pointer()), C.int(mode)) != 0
	}
	return false
}

func (ptr *QCameraFocus) IsFocusPointModeSupported(mode QCameraFocus__FocusPointMode) bool {
	if ptr.Pointer() != nil {
		return C.QCameraFocus_IsFocusPointModeSupported(C.QtObjectPtr(ptr.Pointer()), C.int(mode)) != 0
	}
	return false
}
