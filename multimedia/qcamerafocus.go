package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
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
		n.SetObjectName("QCameraFocus_" + qt.RandomIdentifier())
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraFocus::digitalZoom")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QCameraFocus_DigitalZoom(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraFocus) FocusMode() QCameraFocus__FocusMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraFocus::focusMode")
		}
	}()

	if ptr.Pointer() != nil {
		return QCameraFocus__FocusMode(C.QCameraFocus_FocusMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraFocus) FocusPointMode() QCameraFocus__FocusPointMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraFocus::focusPointMode")
		}
	}()

	if ptr.Pointer() != nil {
		return QCameraFocus__FocusPointMode(C.QCameraFocus_FocusPointMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraFocus) OpticalZoom() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraFocus::opticalZoom")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QCameraFocus_OpticalZoom(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraFocus) SetCustomFocusPoint(point core.QPointF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraFocus::setCustomFocusPoint")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraFocus_SetCustomFocusPoint(ptr.Pointer(), core.PointerFromQPointF(point))
	}
}

func (ptr *QCameraFocus) SetFocusMode(mode QCameraFocus__FocusMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraFocus::setFocusMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraFocus_SetFocusMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QCameraFocus) SetFocusPointMode(mode QCameraFocus__FocusPointMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraFocus::setFocusPointMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraFocus_SetFocusPointMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QCameraFocus) ConnectFocusZonesChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraFocus::focusZonesChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraFocus_ConnectFocusZonesChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "focusZonesChanged", f)
	}
}

func (ptr *QCameraFocus) DisconnectFocusZonesChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraFocus::focusZonesChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraFocus_DisconnectFocusZonesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "focusZonesChanged")
	}
}

//export callbackQCameraFocusFocusZonesChanged
func callbackQCameraFocusFocusZonesChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraFocus::focusZonesChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "focusZonesChanged").(func())()
}

func (ptr *QCameraFocus) IsAvailable() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraFocus::isAvailable")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QCameraFocus_IsAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCameraFocus) IsFocusModeSupported(mode QCameraFocus__FocusMode) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraFocus::isFocusModeSupported")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QCameraFocus_IsFocusModeSupported(ptr.Pointer(), C.int(mode)) != 0
	}
	return false
}

func (ptr *QCameraFocus) IsFocusPointModeSupported(mode QCameraFocus__FocusPointMode) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraFocus::isFocusPointModeSupported")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QCameraFocus_IsFocusPointModeSupported(ptr.Pointer(), C.int(mode)) != 0
	}
	return false
}

func (ptr *QCameraFocus) MaximumDigitalZoom() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraFocus::maximumDigitalZoom")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QCameraFocus_MaximumDigitalZoom(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraFocus) MaximumOpticalZoom() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraFocus::maximumOpticalZoom")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QCameraFocus_MaximumOpticalZoom(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraFocus) ZoomTo(optical float64, digital float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraFocus::zoomTo")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraFocus_ZoomTo(ptr.Pointer(), C.double(optical), C.double(digital))
	}
}
