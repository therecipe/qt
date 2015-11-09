package multimedia

//#include "qmediarecordercontrol.h"
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
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QMediaRecorderControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMediaRecorderControl) QMediaRecorderControl_PTR() *QMediaRecorderControl {
	return ptr
}

func (ptr *QMediaRecorderControl) ApplySettings() {
	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_ApplySettings(ptr.Pointer())
	}
}

func (ptr *QMediaRecorderControl) ConnectError(f func(error int, errorString string)) {
	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_ConnectError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error", f)
	}
}

func (ptr *QMediaRecorderControl) DisconnectError() {
	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_DisconnectError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error")
	}
}

//export callbackQMediaRecorderControlError
func callbackQMediaRecorderControlError(ptrName *C.char, error C.int, errorString *C.char) {
	qt.GetSignal(C.GoString(ptrName), "error").(func(int, string))(int(error), C.GoString(errorString))
}

func (ptr *QMediaRecorderControl) IsMuted() bool {
	if ptr.Pointer() != nil {
		return C.QMediaRecorderControl_IsMuted(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaRecorderControl) ConnectMutedChanged(f func(muted bool)) {
	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_ConnectMutedChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mutedChanged", f)
	}
}

func (ptr *QMediaRecorderControl) DisconnectMutedChanged() {
	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_DisconnectMutedChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mutedChanged")
	}
}

//export callbackQMediaRecorderControlMutedChanged
func callbackQMediaRecorderControlMutedChanged(ptrName *C.char, muted C.int) {
	qt.GetSignal(C.GoString(ptrName), "mutedChanged").(func(bool))(int(muted) != 0)
}

func (ptr *QMediaRecorderControl) SetMuted(muted bool) {
	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_SetMuted(ptr.Pointer(), C.int(qt.GoBoolToInt(muted)))
	}
}

func (ptr *QMediaRecorderControl) SetOutputLocation(location core.QUrl_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QMediaRecorderControl_SetOutputLocation(ptr.Pointer(), core.PointerFromQUrl(location)) != 0
	}
	return false
}

func (ptr *QMediaRecorderControl) SetState(state QMediaRecorder__State) {
	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_SetState(ptr.Pointer(), C.int(state))
	}
}

func (ptr *QMediaRecorderControl) SetVolume(gain float64) {
	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_SetVolume(ptr.Pointer(), C.double(gain))
	}
}

func (ptr *QMediaRecorderControl) ConnectStateChanged(f func(state QMediaRecorder__State)) {
	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QMediaRecorderControl) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQMediaRecorderControlStateChanged
func callbackQMediaRecorderControlStateChanged(ptrName *C.char, state C.int) {
	qt.GetSignal(C.GoString(ptrName), "stateChanged").(func(QMediaRecorder__State))(QMediaRecorder__State(state))
}

func (ptr *QMediaRecorderControl) Status() QMediaRecorder__Status {
	if ptr.Pointer() != nil {
		return QMediaRecorder__Status(C.QMediaRecorderControl_Status(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaRecorderControl) ConnectStatusChanged(f func(status QMediaRecorder__Status)) {
	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_ConnectStatusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "statusChanged", f)
	}
}

func (ptr *QMediaRecorderControl) DisconnectStatusChanged() {
	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_DisconnectStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "statusChanged")
	}
}

//export callbackQMediaRecorderControlStatusChanged
func callbackQMediaRecorderControlStatusChanged(ptrName *C.char, status C.int) {
	qt.GetSignal(C.GoString(ptrName), "statusChanged").(func(QMediaRecorder__Status))(QMediaRecorder__Status(status))
}

func (ptr *QMediaRecorderControl) Volume() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QMediaRecorderControl_Volume(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaRecorderControl) DestroyQMediaRecorderControl() {
	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_DestroyQMediaRecorderControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
