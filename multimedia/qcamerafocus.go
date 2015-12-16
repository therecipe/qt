package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QCameraFocus struct {
	core.QObject
}

type QCameraFocus_ITF interface {
	core.QObject_ITF
	QCameraFocus_PTR() *QCameraFocus
}

func PointerFromQCameraFocus(ptr QCameraFocus_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraFocus_PTR().Pointer()
	}
	return nil
}

func NewQCameraFocusFromPointer(ptr unsafe.Pointer) *QCameraFocus {
	var n = new(QCameraFocus)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QCameraFocus_") {
		n.SetObjectName("QCameraFocus_" + qt.Identifier())
	}
	return n
}

func (ptr *QCameraFocus) QCameraFocus_PTR() *QCameraFocus {
	return ptr
}

//QCameraFocus::FocusMode
type QCameraFocus__FocusMode int64

const (
	QCameraFocus__ManualFocus     = QCameraFocus__FocusMode(0x1)
	QCameraFocus__HyperfocalFocus = QCameraFocus__FocusMode(0x02)
	QCameraFocus__InfinityFocus   = QCameraFocus__FocusMode(0x04)
	QCameraFocus__AutoFocus       = QCameraFocus__FocusMode(0x8)
	QCameraFocus__ContinuousFocus = QCameraFocus__FocusMode(0x10)
	QCameraFocus__MacroFocus      = QCameraFocus__FocusMode(0x20)
)

//QCameraFocus::FocusPointMode
type QCameraFocus__FocusPointMode int64

const (
	QCameraFocus__FocusPointAuto          = QCameraFocus__FocusPointMode(0)
	QCameraFocus__FocusPointCenter        = QCameraFocus__FocusPointMode(1)
	QCameraFocus__FocusPointFaceDetection = QCameraFocus__FocusPointMode(2)
	QCameraFocus__FocusPointCustom        = QCameraFocus__FocusPointMode(3)
)

func (ptr *QCameraFocus) DigitalZoom() float64 {
	defer qt.Recovering("QCameraFocus::digitalZoom")

	if ptr.Pointer() != nil {
		return float64(C.QCameraFocus_DigitalZoom(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraFocus) FocusMode() QCameraFocus__FocusMode {
	defer qt.Recovering("QCameraFocus::focusMode")

	if ptr.Pointer() != nil {
		return QCameraFocus__FocusMode(C.QCameraFocus_FocusMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraFocus) FocusPointMode() QCameraFocus__FocusPointMode {
	defer qt.Recovering("QCameraFocus::focusPointMode")

	if ptr.Pointer() != nil {
		return QCameraFocus__FocusPointMode(C.QCameraFocus_FocusPointMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraFocus) OpticalZoom() float64 {
	defer qt.Recovering("QCameraFocus::opticalZoom")

	if ptr.Pointer() != nil {
		return float64(C.QCameraFocus_OpticalZoom(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraFocus) SetCustomFocusPoint(point core.QPointF_ITF) {
	defer qt.Recovering("QCameraFocus::setCustomFocusPoint")

	if ptr.Pointer() != nil {
		C.QCameraFocus_SetCustomFocusPoint(ptr.Pointer(), core.PointerFromQPointF(point))
	}
}

func (ptr *QCameraFocus) SetFocusMode(mode QCameraFocus__FocusMode) {
	defer qt.Recovering("QCameraFocus::setFocusMode")

	if ptr.Pointer() != nil {
		C.QCameraFocus_SetFocusMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QCameraFocus) SetFocusPointMode(mode QCameraFocus__FocusPointMode) {
	defer qt.Recovering("QCameraFocus::setFocusPointMode")

	if ptr.Pointer() != nil {
		C.QCameraFocus_SetFocusPointMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QCameraFocus) ConnectDigitalZoomChanged(f func(value float64)) {
	defer qt.Recovering("connect QCameraFocus::digitalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraFocus_ConnectDigitalZoomChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "digitalZoomChanged", f)
	}
}

func (ptr *QCameraFocus) DisconnectDigitalZoomChanged() {
	defer qt.Recovering("disconnect QCameraFocus::digitalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraFocus_DisconnectDigitalZoomChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "digitalZoomChanged")
	}
}

//export callbackQCameraFocusDigitalZoomChanged
func callbackQCameraFocusDigitalZoomChanged(ptrName *C.char, value C.double) {
	defer qt.Recovering("callback QCameraFocus::digitalZoomChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "digitalZoomChanged")
	if signal != nil {
		signal.(func(float64))(float64(value))
	}

}

func (ptr *QCameraFocus) ConnectFocusZonesChanged(f func()) {
	defer qt.Recovering("connect QCameraFocus::focusZonesChanged")

	if ptr.Pointer() != nil {
		C.QCameraFocus_ConnectFocusZonesChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "focusZonesChanged", f)
	}
}

func (ptr *QCameraFocus) DisconnectFocusZonesChanged() {
	defer qt.Recovering("disconnect QCameraFocus::focusZonesChanged")

	if ptr.Pointer() != nil {
		C.QCameraFocus_DisconnectFocusZonesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "focusZonesChanged")
	}
}

//export callbackQCameraFocusFocusZonesChanged
func callbackQCameraFocusFocusZonesChanged(ptrName *C.char) {
	defer qt.Recovering("callback QCameraFocus::focusZonesChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "focusZonesChanged")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QCameraFocus) IsAvailable() bool {
	defer qt.Recovering("QCameraFocus::isAvailable")

	if ptr.Pointer() != nil {
		return C.QCameraFocus_IsAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCameraFocus) IsFocusModeSupported(mode QCameraFocus__FocusMode) bool {
	defer qt.Recovering("QCameraFocus::isFocusModeSupported")

	if ptr.Pointer() != nil {
		return C.QCameraFocus_IsFocusModeSupported(ptr.Pointer(), C.int(mode)) != 0
	}
	return false
}

func (ptr *QCameraFocus) IsFocusPointModeSupported(mode QCameraFocus__FocusPointMode) bool {
	defer qt.Recovering("QCameraFocus::isFocusPointModeSupported")

	if ptr.Pointer() != nil {
		return C.QCameraFocus_IsFocusPointModeSupported(ptr.Pointer(), C.int(mode)) != 0
	}
	return false
}

func (ptr *QCameraFocus) MaximumDigitalZoom() float64 {
	defer qt.Recovering("QCameraFocus::maximumDigitalZoom")

	if ptr.Pointer() != nil {
		return float64(C.QCameraFocus_MaximumDigitalZoom(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraFocus) ConnectMaximumDigitalZoomChanged(f func(zoom float64)) {
	defer qt.Recovering("connect QCameraFocus::maximumDigitalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraFocus_ConnectMaximumDigitalZoomChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "maximumDigitalZoomChanged", f)
	}
}

func (ptr *QCameraFocus) DisconnectMaximumDigitalZoomChanged() {
	defer qt.Recovering("disconnect QCameraFocus::maximumDigitalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraFocus_DisconnectMaximumDigitalZoomChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "maximumDigitalZoomChanged")
	}
}

//export callbackQCameraFocusMaximumDigitalZoomChanged
func callbackQCameraFocusMaximumDigitalZoomChanged(ptrName *C.char, zoom C.double) {
	defer qt.Recovering("callback QCameraFocus::maximumDigitalZoomChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "maximumDigitalZoomChanged")
	if signal != nil {
		signal.(func(float64))(float64(zoom))
	}

}

func (ptr *QCameraFocus) MaximumOpticalZoom() float64 {
	defer qt.Recovering("QCameraFocus::maximumOpticalZoom")

	if ptr.Pointer() != nil {
		return float64(C.QCameraFocus_MaximumOpticalZoom(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraFocus) ConnectMaximumOpticalZoomChanged(f func(zoom float64)) {
	defer qt.Recovering("connect QCameraFocus::maximumOpticalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraFocus_ConnectMaximumOpticalZoomChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "maximumOpticalZoomChanged", f)
	}
}

func (ptr *QCameraFocus) DisconnectMaximumOpticalZoomChanged() {
	defer qt.Recovering("disconnect QCameraFocus::maximumOpticalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraFocus_DisconnectMaximumOpticalZoomChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "maximumOpticalZoomChanged")
	}
}

//export callbackQCameraFocusMaximumOpticalZoomChanged
func callbackQCameraFocusMaximumOpticalZoomChanged(ptrName *C.char, zoom C.double) {
	defer qt.Recovering("callback QCameraFocus::maximumOpticalZoomChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "maximumOpticalZoomChanged")
	if signal != nil {
		signal.(func(float64))(float64(zoom))
	}

}

func (ptr *QCameraFocus) ConnectOpticalZoomChanged(f func(value float64)) {
	defer qt.Recovering("connect QCameraFocus::opticalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraFocus_ConnectOpticalZoomChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "opticalZoomChanged", f)
	}
}

func (ptr *QCameraFocus) DisconnectOpticalZoomChanged() {
	defer qt.Recovering("disconnect QCameraFocus::opticalZoomChanged")

	if ptr.Pointer() != nil {
		C.QCameraFocus_DisconnectOpticalZoomChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "opticalZoomChanged")
	}
}

//export callbackQCameraFocusOpticalZoomChanged
func callbackQCameraFocusOpticalZoomChanged(ptrName *C.char, value C.double) {
	defer qt.Recovering("callback QCameraFocus::opticalZoomChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "opticalZoomChanged")
	if signal != nil {
		signal.(func(float64))(float64(value))
	}

}

func (ptr *QCameraFocus) ZoomTo(optical float64, digital float64) {
	defer qt.Recovering("QCameraFocus::zoomTo")

	if ptr.Pointer() != nil {
		C.QCameraFocus_ZoomTo(ptr.Pointer(), C.double(optical), C.double(digital))
	}
}
