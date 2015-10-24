package multimedia

//#include "qaudiodecoder.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAudioDecoder struct {
	QMediaObject
}

type QAudioDecoderITF interface {
	QMediaObjectITF
	QAudioDecoderPTR() *QAudioDecoder
}

func PointerFromQAudioDecoder(ptr QAudioDecoderITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAudioDecoderPTR().Pointer()
	}
	return nil
}

func QAudioDecoderFromPointer(ptr unsafe.Pointer) *QAudioDecoder {
	var n = new(QAudioDecoder)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QAudioDecoder_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAudioDecoder) QAudioDecoderPTR() *QAudioDecoder {
	return ptr
}

//QAudioDecoder::Error
type QAudioDecoder__Error int

var (
	QAudioDecoder__NoError             = QAudioDecoder__Error(0)
	QAudioDecoder__ResourceError       = QAudioDecoder__Error(1)
	QAudioDecoder__FormatError         = QAudioDecoder__Error(2)
	QAudioDecoder__AccessDeniedError   = QAudioDecoder__Error(3)
	QAudioDecoder__ServiceMissingError = QAudioDecoder__Error(4)
)

//QAudioDecoder::State
type QAudioDecoder__State int

var (
	QAudioDecoder__StoppedState  = QAudioDecoder__State(0)
	QAudioDecoder__DecodingState = QAudioDecoder__State(1)
)

func (ptr *QAudioDecoder) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioDecoder_ErrorString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func NewQAudioDecoder(parent core.QObjectITF) *QAudioDecoder {
	return QAudioDecoderFromPointer(unsafe.Pointer(C.QAudioDecoder_NewQAudioDecoder(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QAudioDecoder) BufferAvailable() bool {
	if ptr.Pointer() != nil {
		return C.QAudioDecoder_BufferAvailable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAudioDecoder) ConnectBufferAvailableChanged(f func(available bool)) {
	if ptr.Pointer() != nil {
		C.QAudioDecoder_ConnectBufferAvailableChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "bufferAvailableChanged", f)
	}
}

func (ptr *QAudioDecoder) DisconnectBufferAvailableChanged() {
	if ptr.Pointer() != nil {
		C.QAudioDecoder_DisconnectBufferAvailableChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "bufferAvailableChanged")
	}
}

//export callbackQAudioDecoderBufferAvailableChanged
func callbackQAudioDecoderBufferAvailableChanged(ptrName *C.char, available C.int) {
	qt.GetSignal(C.GoString(ptrName), "bufferAvailableChanged").(func(bool))(int(available) != 0)
}

func (ptr *QAudioDecoder) ConnectBufferReady(f func()) {
	if ptr.Pointer() != nil {
		C.QAudioDecoder_ConnectBufferReady(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "bufferReady", f)
	}
}

func (ptr *QAudioDecoder) DisconnectBufferReady() {
	if ptr.Pointer() != nil {
		C.QAudioDecoder_DisconnectBufferReady(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "bufferReady")
	}
}

//export callbackQAudioDecoderBufferReady
func callbackQAudioDecoderBufferReady(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "bufferReady").(func())()
}

func (ptr *QAudioDecoder) Error() QAudioDecoder__Error {
	if ptr.Pointer() != nil {
		return QAudioDecoder__Error(C.QAudioDecoder_Error(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAudioDecoder) ConnectFinished(f func()) {
	if ptr.Pointer() != nil {
		C.QAudioDecoder_ConnectFinished(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QAudioDecoder) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QAudioDecoder_DisconnectFinished(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQAudioDecoderFinished
func callbackQAudioDecoderFinished(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "finished").(func())()
}

func (ptr *QAudioDecoder) SetAudioFormat(format QAudioFormatITF) {
	if ptr.Pointer() != nil {
		C.QAudioDecoder_SetAudioFormat(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAudioFormat(format)))
	}
}

func (ptr *QAudioDecoder) SetSourceDevice(device core.QIODeviceITF) {
	if ptr.Pointer() != nil {
		C.QAudioDecoder_SetSourceDevice(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQIODevice(device)))
	}
}

func (ptr *QAudioDecoder) SetSourceFilename(fileName string) {
	if ptr.Pointer() != nil {
		C.QAudioDecoder_SetSourceFilename(C.QtObjectPtr(ptr.Pointer()), C.CString(fileName))
	}
}

func (ptr *QAudioDecoder) ConnectSourceChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QAudioDecoder_ConnectSourceChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "sourceChanged", f)
	}
}

func (ptr *QAudioDecoder) DisconnectSourceChanged() {
	if ptr.Pointer() != nil {
		C.QAudioDecoder_DisconnectSourceChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "sourceChanged")
	}
}

//export callbackQAudioDecoderSourceChanged
func callbackQAudioDecoderSourceChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "sourceChanged").(func())()
}

func (ptr *QAudioDecoder) SourceDevice() *core.QIODevice {
	if ptr.Pointer() != nil {
		return core.QIODeviceFromPointer(unsafe.Pointer(C.QAudioDecoder_SourceDevice(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QAudioDecoder) SourceFilename() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioDecoder_SourceFilename(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QAudioDecoder) Start() {
	if ptr.Pointer() != nil {
		C.QAudioDecoder_Start(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QAudioDecoder) ConnectStateChanged(f func(state QAudioDecoder__State)) {
	if ptr.Pointer() != nil {
		C.QAudioDecoder_ConnectStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QAudioDecoder) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QAudioDecoder_DisconnectStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQAudioDecoderStateChanged
func callbackQAudioDecoderStateChanged(ptrName *C.char, state C.int) {
	qt.GetSignal(C.GoString(ptrName), "stateChanged").(func(QAudioDecoder__State))(QAudioDecoder__State(state))
}

func (ptr *QAudioDecoder) Stop() {
	if ptr.Pointer() != nil {
		C.QAudioDecoder_Stop(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QAudioDecoder) DestroyQAudioDecoder() {
	if ptr.Pointer() != nil {
		C.QAudioDecoder_DestroyQAudioDecoder(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
