package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QMediaRecorderControl struct {
	QMediaControl
}

type QMediaRecorderControl_ITF interface {
	QMediaControl_ITF
	QMediaRecorderControl_PTR() *QMediaRecorderControl
}

func PointerFromQMediaRecorderControl(ptr QMediaRecorderControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaRecorderControl_PTR().Pointer()
	}
	return nil
}

func NewQMediaRecorderControlFromPointer(ptr unsafe.Pointer) *QMediaRecorderControl {
	var n = new(QMediaRecorderControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMediaRecorderControl_") {
		n.SetObjectName("QMediaRecorderControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QMediaRecorderControl) QMediaRecorderControl_PTR() *QMediaRecorderControl {
	return ptr
}

func (ptr *QMediaRecorderControl) ConnectActualLocationChanged(f func(location *core.QUrl)) {
	defer qt.Recovering("connect QMediaRecorderControl::actualLocationChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_ConnectActualLocationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "actualLocationChanged", f)
	}
}

func (ptr *QMediaRecorderControl) DisconnectActualLocationChanged() {
	defer qt.Recovering("disconnect QMediaRecorderControl::actualLocationChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_DisconnectActualLocationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "actualLocationChanged")
	}
}

//export callbackQMediaRecorderControlActualLocationChanged
func callbackQMediaRecorderControlActualLocationChanged(ptrName *C.char, location unsafe.Pointer) {
	defer qt.Recovering("callback QMediaRecorderControl::actualLocationChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "actualLocationChanged"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(location))
	}

}

func (ptr *QMediaRecorderControl) ApplySettings() {
	defer qt.Recovering("QMediaRecorderControl::applySettings")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_ApplySettings(ptr.Pointer())
	}
}

func (ptr *QMediaRecorderControl) Duration() int64 {
	defer qt.Recovering("QMediaRecorderControl::duration")

	if ptr.Pointer() != nil {
		return int64(C.QMediaRecorderControl_Duration(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaRecorderControl) ConnectDurationChanged(f func(duration int64)) {
	defer qt.Recovering("connect QMediaRecorderControl::durationChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_ConnectDurationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "durationChanged", f)
	}
}

func (ptr *QMediaRecorderControl) DisconnectDurationChanged() {
	defer qt.Recovering("disconnect QMediaRecorderControl::durationChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_DisconnectDurationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "durationChanged")
	}
}

//export callbackQMediaRecorderControlDurationChanged
func callbackQMediaRecorderControlDurationChanged(ptrName *C.char, duration C.longlong) {
	defer qt.Recovering("callback QMediaRecorderControl::durationChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "durationChanged"); signal != nil {
		signal.(func(int64))(int64(duration))
	}

}

func (ptr *QMediaRecorderControl) ConnectError(f func(error int, errorString string)) {
	defer qt.Recovering("connect QMediaRecorderControl::error")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_ConnectError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error", f)
	}
}

func (ptr *QMediaRecorderControl) DisconnectError() {
	defer qt.Recovering("disconnect QMediaRecorderControl::error")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_DisconnectError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error")
	}
}

//export callbackQMediaRecorderControlError
func callbackQMediaRecorderControlError(ptrName *C.char, error C.int, errorString *C.char) {
	defer qt.Recovering("callback QMediaRecorderControl::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error"); signal != nil {
		signal.(func(int, string))(int(error), C.GoString(errorString))
	}

}

func (ptr *QMediaRecorderControl) IsMuted() bool {
	defer qt.Recovering("QMediaRecorderControl::isMuted")

	if ptr.Pointer() != nil {
		return C.QMediaRecorderControl_IsMuted(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaRecorderControl) ConnectMutedChanged(f func(muted bool)) {
	defer qt.Recovering("connect QMediaRecorderControl::mutedChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_ConnectMutedChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mutedChanged", f)
	}
}

func (ptr *QMediaRecorderControl) DisconnectMutedChanged() {
	defer qt.Recovering("disconnect QMediaRecorderControl::mutedChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_DisconnectMutedChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mutedChanged")
	}
}

//export callbackQMediaRecorderControlMutedChanged
func callbackQMediaRecorderControlMutedChanged(ptrName *C.char, muted C.int) {
	defer qt.Recovering("callback QMediaRecorderControl::mutedChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "mutedChanged"); signal != nil {
		signal.(func(bool))(int(muted) != 0)
	}

}

func (ptr *QMediaRecorderControl) OutputLocation() *core.QUrl {
	defer qt.Recovering("QMediaRecorderControl::outputLocation")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QMediaRecorderControl_OutputLocation(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMediaRecorderControl) SetMuted(muted bool) {
	defer qt.Recovering("QMediaRecorderControl::setMuted")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_SetMuted(ptr.Pointer(), C.int(qt.GoBoolToInt(muted)))
	}
}

func (ptr *QMediaRecorderControl) SetOutputLocation(location core.QUrl_ITF) bool {
	defer qt.Recovering("QMediaRecorderControl::setOutputLocation")

	if ptr.Pointer() != nil {
		return C.QMediaRecorderControl_SetOutputLocation(ptr.Pointer(), core.PointerFromQUrl(location)) != 0
	}
	return false
}

func (ptr *QMediaRecorderControl) SetState(state QMediaRecorder__State) {
	defer qt.Recovering("QMediaRecorderControl::setState")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_SetState(ptr.Pointer(), C.int(state))
	}
}

func (ptr *QMediaRecorderControl) SetVolume(gain float64) {
	defer qt.Recovering("QMediaRecorderControl::setVolume")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_SetVolume(ptr.Pointer(), C.double(gain))
	}
}

func (ptr *QMediaRecorderControl) State() QMediaRecorder__State {
	defer qt.Recovering("QMediaRecorderControl::state")

	if ptr.Pointer() != nil {
		return QMediaRecorder__State(C.QMediaRecorderControl_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaRecorderControl) ConnectStateChanged(f func(state QMediaRecorder__State)) {
	defer qt.Recovering("connect QMediaRecorderControl::stateChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QMediaRecorderControl) DisconnectStateChanged() {
	defer qt.Recovering("disconnect QMediaRecorderControl::stateChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQMediaRecorderControlStateChanged
func callbackQMediaRecorderControlStateChanged(ptrName *C.char, state C.int) {
	defer qt.Recovering("callback QMediaRecorderControl::stateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stateChanged"); signal != nil {
		signal.(func(QMediaRecorder__State))(QMediaRecorder__State(state))
	}

}

func (ptr *QMediaRecorderControl) Status() QMediaRecorder__Status {
	defer qt.Recovering("QMediaRecorderControl::status")

	if ptr.Pointer() != nil {
		return QMediaRecorder__Status(C.QMediaRecorderControl_Status(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaRecorderControl) ConnectStatusChanged(f func(status QMediaRecorder__Status)) {
	defer qt.Recovering("connect QMediaRecorderControl::statusChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_ConnectStatusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "statusChanged", f)
	}
}

func (ptr *QMediaRecorderControl) DisconnectStatusChanged() {
	defer qt.Recovering("disconnect QMediaRecorderControl::statusChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_DisconnectStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "statusChanged")
	}
}

//export callbackQMediaRecorderControlStatusChanged
func callbackQMediaRecorderControlStatusChanged(ptrName *C.char, status C.int) {
	defer qt.Recovering("callback QMediaRecorderControl::statusChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "statusChanged"); signal != nil {
		signal.(func(QMediaRecorder__Status))(QMediaRecorder__Status(status))
	}

}

func (ptr *QMediaRecorderControl) Volume() float64 {
	defer qt.Recovering("QMediaRecorderControl::volume")

	if ptr.Pointer() != nil {
		return float64(C.QMediaRecorderControl_Volume(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaRecorderControl) ConnectVolumeChanged(f func(gain float64)) {
	defer qt.Recovering("connect QMediaRecorderControl::volumeChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_ConnectVolumeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "volumeChanged", f)
	}
}

func (ptr *QMediaRecorderControl) DisconnectVolumeChanged() {
	defer qt.Recovering("disconnect QMediaRecorderControl::volumeChanged")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_DisconnectVolumeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "volumeChanged")
	}
}

//export callbackQMediaRecorderControlVolumeChanged
func callbackQMediaRecorderControlVolumeChanged(ptrName *C.char, gain C.double) {
	defer qt.Recovering("callback QMediaRecorderControl::volumeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "volumeChanged"); signal != nil {
		signal.(func(float64))(float64(gain))
	}

}

func (ptr *QMediaRecorderControl) DestroyQMediaRecorderControl() {
	defer qt.Recovering("QMediaRecorderControl::~QMediaRecorderControl")

	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_DestroyQMediaRecorderControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
