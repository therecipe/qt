package multimedia

//#include "qmediarecordercontrol.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QMediaRecorderControl struct {
	QMediaControl
}

type QMediaRecorderControlITF interface {
	QMediaControlITF
	QMediaRecorderControlPTR() *QMediaRecorderControl
}

func PointerFromQMediaRecorderControl(ptr QMediaRecorderControlITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaRecorderControlPTR().Pointer()
	}
	return nil
}

func QMediaRecorderControlFromPointer(ptr unsafe.Pointer) *QMediaRecorderControl {
	var n = new(QMediaRecorderControl)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QMediaRecorderControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMediaRecorderControl) QMediaRecorderControlPTR() *QMediaRecorderControl {
	return ptr
}

func (ptr *QMediaRecorderControl) ConnectActualLocationChanged(f func(location string)) {
	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_ConnectActualLocationChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "actualLocationChanged", f)
	}
}

func (ptr *QMediaRecorderControl) DisconnectActualLocationChanged() {
	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_DisconnectActualLocationChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "actualLocationChanged")
	}
}

//export callbackQMediaRecorderControlActualLocationChanged
func callbackQMediaRecorderControlActualLocationChanged(ptrName *C.char, location *C.char) {
	qt.GetSignal(C.GoString(ptrName), "actualLocationChanged").(func(string))(C.GoString(location))
}

func (ptr *QMediaRecorderControl) ApplySettings() {
	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_ApplySettings(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QMediaRecorderControl) ConnectError(f func(error int, errorString string)) {
	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_ConnectError(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "error", f)
	}
}

func (ptr *QMediaRecorderControl) DisconnectError() {
	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_DisconnectError(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "error")
	}
}

//export callbackQMediaRecorderControlError
func callbackQMediaRecorderControlError(ptrName *C.char, error C.int, errorString *C.char) {
	qt.GetSignal(C.GoString(ptrName), "error").(func(int, string))(int(error), C.GoString(errorString))
}

func (ptr *QMediaRecorderControl) IsMuted() bool {
	if ptr.Pointer() != nil {
		return C.QMediaRecorderControl_IsMuted(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMediaRecorderControl) ConnectMutedChanged(f func(muted bool)) {
	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_ConnectMutedChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "mutedChanged", f)
	}
}

func (ptr *QMediaRecorderControl) DisconnectMutedChanged() {
	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_DisconnectMutedChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "mutedChanged")
	}
}

//export callbackQMediaRecorderControlMutedChanged
func callbackQMediaRecorderControlMutedChanged(ptrName *C.char, muted C.int) {
	qt.GetSignal(C.GoString(ptrName), "mutedChanged").(func(bool))(int(muted) != 0)
}

func (ptr *QMediaRecorderControl) OutputLocation() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaRecorderControl_OutputLocation(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QMediaRecorderControl) SetMuted(muted bool) {
	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_SetMuted(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(muted)))
	}
}

func (ptr *QMediaRecorderControl) SetOutputLocation(location string) bool {
	if ptr.Pointer() != nil {
		return C.QMediaRecorderControl_SetOutputLocation(C.QtObjectPtr(ptr.Pointer()), C.CString(location)) != 0
	}
	return false
}

func (ptr *QMediaRecorderControl) SetState(state QMediaRecorder__State) {
	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_SetState(C.QtObjectPtr(ptr.Pointer()), C.int(state))
	}
}

func (ptr *QMediaRecorderControl) ConnectStateChanged(f func(state QMediaRecorder__State)) {
	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_ConnectStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QMediaRecorderControl) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_DisconnectStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQMediaRecorderControlStateChanged
func callbackQMediaRecorderControlStateChanged(ptrName *C.char, state C.int) {
	qt.GetSignal(C.GoString(ptrName), "stateChanged").(func(QMediaRecorder__State))(QMediaRecorder__State(state))
}

func (ptr *QMediaRecorderControl) Status() QMediaRecorder__Status {
	if ptr.Pointer() != nil {
		return QMediaRecorder__Status(C.QMediaRecorderControl_Status(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMediaRecorderControl) ConnectStatusChanged(f func(status QMediaRecorder__Status)) {
	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_ConnectStatusChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "statusChanged", f)
	}
}

func (ptr *QMediaRecorderControl) DisconnectStatusChanged() {
	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_DisconnectStatusChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "statusChanged")
	}
}

//export callbackQMediaRecorderControlStatusChanged
func callbackQMediaRecorderControlStatusChanged(ptrName *C.char, status C.int) {
	qt.GetSignal(C.GoString(ptrName), "statusChanged").(func(QMediaRecorder__Status))(QMediaRecorder__Status(status))
}

func (ptr *QMediaRecorderControl) DestroyQMediaRecorderControl() {
	if ptr.Pointer() != nil {
		C.QMediaRecorderControl_DestroyQMediaRecorderControl(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
