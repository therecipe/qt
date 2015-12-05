package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
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
	for len(n.ObjectName()) < len("QCameraFocusControl_") {
		n.SetObjectName("QCameraFocusControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QCameraFocusControl) QCameraFocusControl_PTR() *QCameraFocusControl {
	return ptr
}

func (ptr *QCameraFocusControl) FocusMode() QCameraFocus__FocusMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraFocusControl::focusMode")
		}
	}()

	if ptr.Pointer() != nil {
		return QCameraFocus__FocusMode(C.QCameraFocusControl_FocusMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraFocusControl) ConnectFocusModeChanged(f func(mode QCameraFocus__FocusMode)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraFocusControl::focusModeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraFocusControl_ConnectFocusModeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "focusModeChanged", f)
	}
}

func (ptr *QCameraFocusControl) DisconnectFocusModeChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraFocusControl::focusModeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraFocusControl_DisconnectFocusModeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "focusModeChanged")
	}
}

//export callbackQCameraFocusControlFocusModeChanged
func callbackQCameraFocusControlFocusModeChanged(ptrName *C.char, mode C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraFocusControl::focusModeChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "focusModeChanged").(func(QCameraFocus__FocusMode))(QCameraFocus__FocusMode(mode))
}

func (ptr *QCameraFocusControl) FocusPointMode() QCameraFocus__FocusPointMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraFocusControl::focusPointMode")
		}
	}()

	if ptr.Pointer() != nil {
		return QCameraFocus__FocusPointMode(C.QCameraFocusControl_FocusPointMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraFocusControl) ConnectFocusPointModeChanged(f func(mode QCameraFocus__FocusPointMode)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraFocusControl::focusPointModeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraFocusControl_ConnectFocusPointModeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "focusPointModeChanged", f)
	}
}

func (ptr *QCameraFocusControl) DisconnectFocusPointModeChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraFocusControl::focusPointModeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraFocusControl_DisconnectFocusPointModeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "focusPointModeChanged")
	}
}

//export callbackQCameraFocusControlFocusPointModeChanged
func callbackQCameraFocusControlFocusPointModeChanged(ptrName *C.char, mode C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraFocusControl::focusPointModeChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "focusPointModeChanged").(func(QCameraFocus__FocusPointMode))(QCameraFocus__FocusPointMode(mode))
}

func (ptr *QCameraFocusControl) ConnectFocusZonesChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraFocusControl::focusZonesChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraFocusControl_ConnectFocusZonesChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "focusZonesChanged", f)
	}
}

func (ptr *QCameraFocusControl) DisconnectFocusZonesChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraFocusControl::focusZonesChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraFocusControl_DisconnectFocusZonesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "focusZonesChanged")
	}
}

//export callbackQCameraFocusControlFocusZonesChanged
func callbackQCameraFocusControlFocusZonesChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraFocusControl::focusZonesChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "focusZonesChanged").(func())()
}

func (ptr *QCameraFocusControl) IsFocusModeSupported(mode QCameraFocus__FocusMode) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraFocusControl::isFocusModeSupported")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QCameraFocusControl_IsFocusModeSupported(ptr.Pointer(), C.int(mode)) != 0
	}
	return false
}

func (ptr *QCameraFocusControl) IsFocusPointModeSupported(mode QCameraFocus__FocusPointMode) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraFocusControl::isFocusPointModeSupported")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QCameraFocusControl_IsFocusPointModeSupported(ptr.Pointer(), C.int(mode)) != 0
	}
	return false
}

func (ptr *QCameraFocusControl) SetCustomFocusPoint(point core.QPointF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraFocusControl::setCustomFocusPoint")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraFocusControl_SetCustomFocusPoint(ptr.Pointer(), core.PointerFromQPointF(point))
	}
}

func (ptr *QCameraFocusControl) SetFocusMode(mode QCameraFocus__FocusMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraFocusControl::setFocusMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraFocusControl_SetFocusMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QCameraFocusControl) SetFocusPointMode(mode QCameraFocus__FocusPointMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraFocusControl::setFocusPointMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraFocusControl_SetFocusPointMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QCameraFocusControl) DestroyQCameraFocusControl() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraFocusControl::~QCameraFocusControl")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraFocusControl_DestroyQCameraFocusControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
