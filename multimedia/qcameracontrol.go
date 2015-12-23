package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
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
	for len(n.ObjectName()) < len("QCameraControl_") {
		n.SetObjectName("QCameraControl_" + qt.Identifier())
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
	defer qt.Recovering("QCameraControl::canChangeProperty")

	if ptr.Pointer() != nil {
		return C.QCameraControl_CanChangeProperty(ptr.Pointer(), C.int(changeType), C.int(status)) != 0
	}
	return false
}

func (ptr *QCameraControl) CaptureMode() QCamera__CaptureMode {
	defer qt.Recovering("QCameraControl::captureMode")

	if ptr.Pointer() != nil {
		return QCamera__CaptureMode(C.QCameraControl_CaptureMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraControl) ConnectCaptureModeChanged(f func(mode QCamera__CaptureMode)) {
	defer qt.Recovering("connect QCameraControl::captureModeChanged")

	if ptr.Pointer() != nil {
		C.QCameraControl_ConnectCaptureModeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "captureModeChanged", f)
	}
}

func (ptr *QCameraControl) DisconnectCaptureModeChanged() {
	defer qt.Recovering("disconnect QCameraControl::captureModeChanged")

	if ptr.Pointer() != nil {
		C.QCameraControl_DisconnectCaptureModeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "captureModeChanged")
	}
}

//export callbackQCameraControlCaptureModeChanged
func callbackQCameraControlCaptureModeChanged(ptrName *C.char, mode C.int) {
	defer qt.Recovering("callback QCameraControl::captureModeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "captureModeChanged"); signal != nil {
		signal.(func(QCamera__CaptureMode))(QCamera__CaptureMode(mode))
	}

}

func (ptr *QCameraControl) ConnectError(f func(error int, errorString string)) {
	defer qt.Recovering("connect QCameraControl::error")

	if ptr.Pointer() != nil {
		C.QCameraControl_ConnectError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error", f)
	}
}

func (ptr *QCameraControl) DisconnectError() {
	defer qt.Recovering("disconnect QCameraControl::error")

	if ptr.Pointer() != nil {
		C.QCameraControl_DisconnectError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error")
	}
}

//export callbackQCameraControlError
func callbackQCameraControlError(ptrName *C.char, error C.int, errorString *C.char) {
	defer qt.Recovering("callback QCameraControl::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error"); signal != nil {
		signal.(func(int, string))(int(error), C.GoString(errorString))
	}

}

func (ptr *QCameraControl) IsCaptureModeSupported(mode QCamera__CaptureMode) bool {
	defer qt.Recovering("QCameraControl::isCaptureModeSupported")

	if ptr.Pointer() != nil {
		return C.QCameraControl_IsCaptureModeSupported(ptr.Pointer(), C.int(mode)) != 0
	}
	return false
}

func (ptr *QCameraControl) SetCaptureMode(mode QCamera__CaptureMode) {
	defer qt.Recovering("QCameraControl::setCaptureMode")

	if ptr.Pointer() != nil {
		C.QCameraControl_SetCaptureMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QCameraControl) SetState(state QCamera__State) {
	defer qt.Recovering("QCameraControl::setState")

	if ptr.Pointer() != nil {
		C.QCameraControl_SetState(ptr.Pointer(), C.int(state))
	}
}

func (ptr *QCameraControl) State() QCamera__State {
	defer qt.Recovering("QCameraControl::state")

	if ptr.Pointer() != nil {
		return QCamera__State(C.QCameraControl_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraControl) ConnectStateChanged(f func(state QCamera__State)) {
	defer qt.Recovering("connect QCameraControl::stateChanged")

	if ptr.Pointer() != nil {
		C.QCameraControl_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QCameraControl) DisconnectStateChanged() {
	defer qt.Recovering("disconnect QCameraControl::stateChanged")

	if ptr.Pointer() != nil {
		C.QCameraControl_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQCameraControlStateChanged
func callbackQCameraControlStateChanged(ptrName *C.char, state C.int) {
	defer qt.Recovering("callback QCameraControl::stateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stateChanged"); signal != nil {
		signal.(func(QCamera__State))(QCamera__State(state))
	}

}

func (ptr *QCameraControl) Status() QCamera__Status {
	defer qt.Recovering("QCameraControl::status")

	if ptr.Pointer() != nil {
		return QCamera__Status(C.QCameraControl_Status(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraControl) ConnectStatusChanged(f func(status QCamera__Status)) {
	defer qt.Recovering("connect QCameraControl::statusChanged")

	if ptr.Pointer() != nil {
		C.QCameraControl_ConnectStatusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "statusChanged", f)
	}
}

func (ptr *QCameraControl) DisconnectStatusChanged() {
	defer qt.Recovering("disconnect QCameraControl::statusChanged")

	if ptr.Pointer() != nil {
		C.QCameraControl_DisconnectStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "statusChanged")
	}
}

//export callbackQCameraControlStatusChanged
func callbackQCameraControlStatusChanged(ptrName *C.char, status C.int) {
	defer qt.Recovering("callback QCameraControl::statusChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "statusChanged"); signal != nil {
		signal.(func(QCamera__Status))(QCamera__Status(status))
	}

}

func (ptr *QCameraControl) DestroyQCameraControl() {
	defer qt.Recovering("QCameraControl::~QCameraControl")

	if ptr.Pointer() != nil {
		C.QCameraControl_DestroyQCameraControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCameraControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QCameraControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCameraControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCameraControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQCameraControlTimerEvent
func callbackQCameraControlTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCameraControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCameraControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCameraControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCameraControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCameraControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQCameraControlChildEvent
func callbackQCameraControlChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCameraControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCameraControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCameraControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCameraControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCameraControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQCameraControlCustomEvent
func callbackQCameraControlCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCameraControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
