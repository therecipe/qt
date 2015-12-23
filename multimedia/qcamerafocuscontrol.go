package multimedia

//#include "multimedia.h"
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
	for len(n.ObjectName()) < len("QCameraFocusControl_") {
		n.SetObjectName("QCameraFocusControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QCameraFocusControl) QCameraFocusControl_PTR() *QCameraFocusControl {
	return ptr
}

func (ptr *QCameraFocusControl) FocusMode() QCameraFocus__FocusMode {
	defer qt.Recovering("QCameraFocusControl::focusMode")

	if ptr.Pointer() != nil {
		return QCameraFocus__FocusMode(C.QCameraFocusControl_FocusMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraFocusControl) ConnectFocusModeChanged(f func(mode QCameraFocus__FocusMode)) {
	defer qt.Recovering("connect QCameraFocusControl::focusModeChanged")

	if ptr.Pointer() != nil {
		C.QCameraFocusControl_ConnectFocusModeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "focusModeChanged", f)
	}
}

func (ptr *QCameraFocusControl) DisconnectFocusModeChanged() {
	defer qt.Recovering("disconnect QCameraFocusControl::focusModeChanged")

	if ptr.Pointer() != nil {
		C.QCameraFocusControl_DisconnectFocusModeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "focusModeChanged")
	}
}

//export callbackQCameraFocusControlFocusModeChanged
func callbackQCameraFocusControlFocusModeChanged(ptrName *C.char, mode C.int) {
	defer qt.Recovering("callback QCameraFocusControl::focusModeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusModeChanged"); signal != nil {
		signal.(func(QCameraFocus__FocusMode))(QCameraFocus__FocusMode(mode))
	}

}

func (ptr *QCameraFocusControl) FocusPointMode() QCameraFocus__FocusPointMode {
	defer qt.Recovering("QCameraFocusControl::focusPointMode")

	if ptr.Pointer() != nil {
		return QCameraFocus__FocusPointMode(C.QCameraFocusControl_FocusPointMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraFocusControl) ConnectFocusPointModeChanged(f func(mode QCameraFocus__FocusPointMode)) {
	defer qt.Recovering("connect QCameraFocusControl::focusPointModeChanged")

	if ptr.Pointer() != nil {
		C.QCameraFocusControl_ConnectFocusPointModeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "focusPointModeChanged", f)
	}
}

func (ptr *QCameraFocusControl) DisconnectFocusPointModeChanged() {
	defer qt.Recovering("disconnect QCameraFocusControl::focusPointModeChanged")

	if ptr.Pointer() != nil {
		C.QCameraFocusControl_DisconnectFocusPointModeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "focusPointModeChanged")
	}
}

//export callbackQCameraFocusControlFocusPointModeChanged
func callbackQCameraFocusControlFocusPointModeChanged(ptrName *C.char, mode C.int) {
	defer qt.Recovering("callback QCameraFocusControl::focusPointModeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusPointModeChanged"); signal != nil {
		signal.(func(QCameraFocus__FocusPointMode))(QCameraFocus__FocusPointMode(mode))
	}

}

func (ptr *QCameraFocusControl) ConnectFocusZonesChanged(f func()) {
	defer qt.Recovering("connect QCameraFocusControl::focusZonesChanged")

	if ptr.Pointer() != nil {
		C.QCameraFocusControl_ConnectFocusZonesChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "focusZonesChanged", f)
	}
}

func (ptr *QCameraFocusControl) DisconnectFocusZonesChanged() {
	defer qt.Recovering("disconnect QCameraFocusControl::focusZonesChanged")

	if ptr.Pointer() != nil {
		C.QCameraFocusControl_DisconnectFocusZonesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "focusZonesChanged")
	}
}

//export callbackQCameraFocusControlFocusZonesChanged
func callbackQCameraFocusControlFocusZonesChanged(ptrName *C.char) {
	defer qt.Recovering("callback QCameraFocusControl::focusZonesChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusZonesChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QCameraFocusControl) IsFocusModeSupported(mode QCameraFocus__FocusMode) bool {
	defer qt.Recovering("QCameraFocusControl::isFocusModeSupported")

	if ptr.Pointer() != nil {
		return C.QCameraFocusControl_IsFocusModeSupported(ptr.Pointer(), C.int(mode)) != 0
	}
	return false
}

func (ptr *QCameraFocusControl) IsFocusPointModeSupported(mode QCameraFocus__FocusPointMode) bool {
	defer qt.Recovering("QCameraFocusControl::isFocusPointModeSupported")

	if ptr.Pointer() != nil {
		return C.QCameraFocusControl_IsFocusPointModeSupported(ptr.Pointer(), C.int(mode)) != 0
	}
	return false
}

func (ptr *QCameraFocusControl) SetCustomFocusPoint(point core.QPointF_ITF) {
	defer qt.Recovering("QCameraFocusControl::setCustomFocusPoint")

	if ptr.Pointer() != nil {
		C.QCameraFocusControl_SetCustomFocusPoint(ptr.Pointer(), core.PointerFromQPointF(point))
	}
}

func (ptr *QCameraFocusControl) SetFocusMode(mode QCameraFocus__FocusMode) {
	defer qt.Recovering("QCameraFocusControl::setFocusMode")

	if ptr.Pointer() != nil {
		C.QCameraFocusControl_SetFocusMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QCameraFocusControl) SetFocusPointMode(mode QCameraFocus__FocusPointMode) {
	defer qt.Recovering("QCameraFocusControl::setFocusPointMode")

	if ptr.Pointer() != nil {
		C.QCameraFocusControl_SetFocusPointMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QCameraFocusControl) DestroyQCameraFocusControl() {
	defer qt.Recovering("QCameraFocusControl::~QCameraFocusControl")

	if ptr.Pointer() != nil {
		C.QCameraFocusControl_DestroyQCameraFocusControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCameraFocusControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QCameraFocusControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCameraFocusControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCameraFocusControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQCameraFocusControlTimerEvent
func callbackQCameraFocusControlTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCameraFocusControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCameraFocusControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCameraFocusControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCameraFocusControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCameraFocusControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQCameraFocusControlChildEvent
func callbackQCameraFocusControlChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCameraFocusControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCameraFocusControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCameraFocusControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCameraFocusControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCameraFocusControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQCameraFocusControlCustomEvent
func callbackQCameraFocusControlCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCameraFocusControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
