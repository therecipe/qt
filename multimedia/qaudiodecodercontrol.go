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

type QAudioDecoderControl_ITF interface {
	QMediaControl_ITF
	QAudioDecoderControl_PTR() *QAudioDecoderControl
}

func PointerFromQAudioDecoderControl(ptr QAudioDecoderControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAudioDecoderControl_PTR().Pointer()
	}
	return nil
}

func NewQAudioDecoderControlFromPointer(ptr unsafe.Pointer) *QAudioDecoderControl {
	var n = new(QAudioDecoderControl)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QAudioDecoderControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAudioDecoderControl) QAudioDecoderControl_PTR() *QAudioDecoderControl {
	return ptr
}

func (ptr *QAudioDecoderControl) BufferAvailable() bool {
	if ptr.Pointer() != nil {
		return C.QAudioDecoderControl_BufferAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAudioDecoderControl) ConnectBufferAvailableChanged(f func(available bool)) {
	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_ConnectBufferAvailableChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "bufferAvailableChanged", f)
	}
}

func (ptr *QAudioDecoderControl) DisconnectBufferAvailableChanged() {
	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_DisconnectBufferAvailableChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "bufferAvailableChanged")
	}
}

//export callbackQAudioDecoderControlBufferAvailableChanged
func callbackQAudioDecoderControlBufferAvailableChanged(ptrName *C.char, available C.int) {
	qt.GetSignal(C.GoString(ptrName), "bufferAvailableChanged").(func(bool))(int(available) != 0)
}

func (ptr *QAudioDecoderControl) ConnectBufferReady(f func()) {
	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_ConnectBufferReady(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "bufferReady", f)
	}
}

func (ptr *QAudioDecoderControl) DisconnectBufferReady() {
	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_DisconnectBufferReady(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "bufferReady")
	}
}

//export callbackQAudioDecoderControlBufferReady
func callbackQAudioDecoderControlBufferReady(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "bufferReady").(func())()
}

func (ptr *QAudioDecoderControl) ConnectError(f func(error int, errorString string)) {
	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_ConnectError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error", f)
	}
}

func (ptr *QAudioDecoderControl) DisconnectError() {
	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_DisconnectError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error")
	}
}

//export callbackQAudioDecoderControlError
func callbackQAudioDecoderControlError(ptrName *C.char, error C.int, errorString *C.char) {
	qt.GetSignal(C.GoString(ptrName), "error").(func(int, string))(int(error), C.GoString(errorString))
}

func (ptr *QAudioDecoderControl) ConnectFinished(f func()) {
	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QAudioDecoderControl) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQAudioDecoderControlFinished
func callbackQAudioDecoderControlFinished(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "finished").(func())()
}

func (ptr *QAudioDecoderControl) SetAudioFormat(format QAudioFormat_ITF) {
	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_SetAudioFormat(ptr.Pointer(), PointerFromQAudioFormat(format))
	}
}

func (ptr *QAudioDecoderControl) SetSourceDevice(device core.QIODevice_ITF) {
	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_SetSourceDevice(ptr.Pointer(), core.PointerFromQIODevice(device))
	}
}

func (ptr *QAudioDecoderControl) SetSourceFilename(fileName string) {
	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_SetSourceFilename(ptr.Pointer(), C.CString(fileName))
	}
}

func (ptr *QAudioDecoderControl) ConnectSourceChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_ConnectSourceChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sourceChanged", f)
	}
}

func (ptr *QAudioDecoderControl) DisconnectSourceChanged() {
	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_DisconnectSourceChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sourceChanged")
	}
}

//export callbackQAudioDecoderControlSourceChanged
func callbackQAudioDecoderControlSourceChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "sourceChanged").(func())()
}

func (ptr *QAudioDecoderControl) SourceDevice() *core.QIODevice {
	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QAudioDecoderControl_SourceDevice(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAudioDecoderControl) SourceFilename() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioDecoderControl_SourceFilename(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAudioDecoderControl) Start() {
	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_Start(ptr.Pointer())
	}
}

func (ptr *QAudioDecoderControl) ConnectStateChanged(f func(state QAudioDecoder__State)) {
	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QAudioDecoderControl) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQAudioDecoderControlStateChanged
func callbackQAudioDecoderControlStateChanged(ptrName *C.char, state C.int) {
	qt.GetSignal(C.GoString(ptrName), "stateChanged").(func(QAudioDecoder__State))(QAudioDecoder__State(state))
}

func (ptr *QAudioDecoderControl) Stop() {
	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_Stop(ptr.Pointer())
	}
}

func (ptr *QAudioDecoderControl) DestroyQAudioDecoderControl() {
	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_DestroyQAudioDecoderControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
