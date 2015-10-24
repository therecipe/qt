package multimedia

//#include "qaudiodecodercontrol.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAudioDecoderControl struct {
	QMediaControl
}

type QAudioDecoderControlITF interface {
	QMediaControlITF
	QAudioDecoderControlPTR() *QAudioDecoderControl
}

func PointerFromQAudioDecoderControl(ptr QAudioDecoderControlITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAudioDecoderControlPTR().Pointer()
	}
	return nil
}

func QAudioDecoderControlFromPointer(ptr unsafe.Pointer) *QAudioDecoderControl {
	var n = new(QAudioDecoderControl)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QAudioDecoderControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAudioDecoderControl) QAudioDecoderControlPTR() *QAudioDecoderControl {
	return ptr
}

func (ptr *QAudioDecoderControl) BufferAvailable() bool {
	if ptr.Pointer() != nil {
		return C.QAudioDecoderControl_BufferAvailable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAudioDecoderControl) ConnectBufferAvailableChanged(f func(available bool)) {
	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_ConnectBufferAvailableChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "bufferAvailableChanged", f)
	}
}

func (ptr *QAudioDecoderControl) DisconnectBufferAvailableChanged() {
	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_DisconnectBufferAvailableChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "bufferAvailableChanged")
	}
}

//export callbackQAudioDecoderControlBufferAvailableChanged
func callbackQAudioDecoderControlBufferAvailableChanged(ptrName *C.char, available C.int) {
	qt.GetSignal(C.GoString(ptrName), "bufferAvailableChanged").(func(bool))(int(available) != 0)
}

func (ptr *QAudioDecoderControl) ConnectBufferReady(f func()) {
	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_ConnectBufferReady(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "bufferReady", f)
	}
}

func (ptr *QAudioDecoderControl) DisconnectBufferReady() {
	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_DisconnectBufferReady(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "bufferReady")
	}
}

//export callbackQAudioDecoderControlBufferReady
func callbackQAudioDecoderControlBufferReady(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "bufferReady").(func())()
}

func (ptr *QAudioDecoderControl) ConnectError(f func(error int, errorString string)) {
	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_ConnectError(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "error", f)
	}
}

func (ptr *QAudioDecoderControl) DisconnectError() {
	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_DisconnectError(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "error")
	}
}

//export callbackQAudioDecoderControlError
func callbackQAudioDecoderControlError(ptrName *C.char, error C.int, errorString *C.char) {
	qt.GetSignal(C.GoString(ptrName), "error").(func(int, string))(int(error), C.GoString(errorString))
}

func (ptr *QAudioDecoderControl) ConnectFinished(f func()) {
	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_ConnectFinished(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QAudioDecoderControl) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_DisconnectFinished(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQAudioDecoderControlFinished
func callbackQAudioDecoderControlFinished(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "finished").(func())()
}

func (ptr *QAudioDecoderControl) SetAudioFormat(format QAudioFormatITF) {
	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_SetAudioFormat(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAudioFormat(format)))
	}
}

func (ptr *QAudioDecoderControl) SetSourceDevice(device core.QIODeviceITF) {
	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_SetSourceDevice(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQIODevice(device)))
	}
}

func (ptr *QAudioDecoderControl) SetSourceFilename(fileName string) {
	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_SetSourceFilename(C.QtObjectPtr(ptr.Pointer()), C.CString(fileName))
	}
}

func (ptr *QAudioDecoderControl) ConnectSourceChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_ConnectSourceChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "sourceChanged", f)
	}
}

func (ptr *QAudioDecoderControl) DisconnectSourceChanged() {
	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_DisconnectSourceChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "sourceChanged")
	}
}

//export callbackQAudioDecoderControlSourceChanged
func callbackQAudioDecoderControlSourceChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "sourceChanged").(func())()
}

func (ptr *QAudioDecoderControl) SourceDevice() *core.QIODevice {
	if ptr.Pointer() != nil {
		return core.QIODeviceFromPointer(unsafe.Pointer(C.QAudioDecoderControl_SourceDevice(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QAudioDecoderControl) SourceFilename() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioDecoderControl_SourceFilename(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QAudioDecoderControl) Start() {
	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_Start(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QAudioDecoderControl) ConnectStateChanged(f func(state QAudioDecoder__State)) {
	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_ConnectStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QAudioDecoderControl) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_DisconnectStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQAudioDecoderControlStateChanged
func callbackQAudioDecoderControlStateChanged(ptrName *C.char, state C.int) {
	qt.GetSignal(C.GoString(ptrName), "stateChanged").(func(QAudioDecoder__State))(QAudioDecoder__State(state))
}

func (ptr *QAudioDecoderControl) Stop() {
	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_Stop(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QAudioDecoderControl) DestroyQAudioDecoderControl() {
	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_DestroyQAudioDecoderControl(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
