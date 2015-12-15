package multimedia

//#include "multimedia.h"
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
	for len(n.ObjectName()) < len("QAudioDecoderControl_") {
		n.SetObjectName("QAudioDecoderControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QAudioDecoderControl) QAudioDecoderControl_PTR() *QAudioDecoderControl {
	return ptr
}

func (ptr *QAudioDecoderControl) BufferAvailable() bool {
	defer qt.Recovering("QAudioDecoderControl::bufferAvailable")

	if ptr.Pointer() != nil {
		return C.QAudioDecoderControl_BufferAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAudioDecoderControl) ConnectBufferAvailableChanged(f func(available bool)) {
	defer qt.Recovering("connect QAudioDecoderControl::bufferAvailableChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_ConnectBufferAvailableChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "bufferAvailableChanged", f)
	}
}

func (ptr *QAudioDecoderControl) DisconnectBufferAvailableChanged() {
	defer qt.Recovering("disconnect QAudioDecoderControl::bufferAvailableChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_DisconnectBufferAvailableChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "bufferAvailableChanged")
	}
}

//export callbackQAudioDecoderControlBufferAvailableChanged
func callbackQAudioDecoderControlBufferAvailableChanged(ptrName *C.char, available C.int) {
	defer qt.Recovering("callback QAudioDecoderControl::bufferAvailableChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "bufferAvailableChanged")
	if signal != nil {
		signal.(func(bool))(int(available) != 0)
	}

}

func (ptr *QAudioDecoderControl) ConnectBufferReady(f func()) {
	defer qt.Recovering("connect QAudioDecoderControl::bufferReady")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_ConnectBufferReady(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "bufferReady", f)
	}
}

func (ptr *QAudioDecoderControl) DisconnectBufferReady() {
	defer qt.Recovering("disconnect QAudioDecoderControl::bufferReady")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_DisconnectBufferReady(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "bufferReady")
	}
}

//export callbackQAudioDecoderControlBufferReady
func callbackQAudioDecoderControlBufferReady(ptrName *C.char) {
	defer qt.Recovering("callback QAudioDecoderControl::bufferReady")

	var signal = qt.GetSignal(C.GoString(ptrName), "bufferReady")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QAudioDecoderControl) ConnectError(f func(error int, errorString string)) {
	defer qt.Recovering("connect QAudioDecoderControl::error")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_ConnectError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error", f)
	}
}

func (ptr *QAudioDecoderControl) DisconnectError() {
	defer qt.Recovering("disconnect QAudioDecoderControl::error")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_DisconnectError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error")
	}
}

//export callbackQAudioDecoderControlError
func callbackQAudioDecoderControlError(ptrName *C.char, error C.int, errorString *C.char) {
	defer qt.Recovering("callback QAudioDecoderControl::error")

	var signal = qt.GetSignal(C.GoString(ptrName), "error")
	if signal != nil {
		signal.(func(int, string))(int(error), C.GoString(errorString))
	}

}

func (ptr *QAudioDecoderControl) ConnectFinished(f func()) {
	defer qt.Recovering("connect QAudioDecoderControl::finished")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QAudioDecoderControl) DisconnectFinished() {
	defer qt.Recovering("disconnect QAudioDecoderControl::finished")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQAudioDecoderControlFinished
func callbackQAudioDecoderControlFinished(ptrName *C.char) {
	defer qt.Recovering("callback QAudioDecoderControl::finished")

	var signal = qt.GetSignal(C.GoString(ptrName), "finished")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QAudioDecoderControl) SetAudioFormat(format QAudioFormat_ITF) {
	defer qt.Recovering("QAudioDecoderControl::setAudioFormat")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_SetAudioFormat(ptr.Pointer(), PointerFromQAudioFormat(format))
	}
}

func (ptr *QAudioDecoderControl) SetSourceDevice(device core.QIODevice_ITF) {
	defer qt.Recovering("QAudioDecoderControl::setSourceDevice")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_SetSourceDevice(ptr.Pointer(), core.PointerFromQIODevice(device))
	}
}

func (ptr *QAudioDecoderControl) SetSourceFilename(fileName string) {
	defer qt.Recovering("QAudioDecoderControl::setSourceFilename")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_SetSourceFilename(ptr.Pointer(), C.CString(fileName))
	}
}

func (ptr *QAudioDecoderControl) ConnectSourceChanged(f func()) {
	defer qt.Recovering("connect QAudioDecoderControl::sourceChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_ConnectSourceChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sourceChanged", f)
	}
}

func (ptr *QAudioDecoderControl) DisconnectSourceChanged() {
	defer qt.Recovering("disconnect QAudioDecoderControl::sourceChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_DisconnectSourceChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sourceChanged")
	}
}

//export callbackQAudioDecoderControlSourceChanged
func callbackQAudioDecoderControlSourceChanged(ptrName *C.char) {
	defer qt.Recovering("callback QAudioDecoderControl::sourceChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "sourceChanged")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QAudioDecoderControl) SourceDevice() *core.QIODevice {
	defer qt.Recovering("QAudioDecoderControl::sourceDevice")

	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QAudioDecoderControl_SourceDevice(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAudioDecoderControl) SourceFilename() string {
	defer qt.Recovering("QAudioDecoderControl::sourceFilename")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioDecoderControl_SourceFilename(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAudioDecoderControl) Start() {
	defer qt.Recovering("QAudioDecoderControl::start")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_Start(ptr.Pointer())
	}
}

func (ptr *QAudioDecoderControl) ConnectStateChanged(f func(state QAudioDecoder__State)) {
	defer qt.Recovering("connect QAudioDecoderControl::stateChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QAudioDecoderControl) DisconnectStateChanged() {
	defer qt.Recovering("disconnect QAudioDecoderControl::stateChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQAudioDecoderControlStateChanged
func callbackQAudioDecoderControlStateChanged(ptrName *C.char, state C.int) {
	defer qt.Recovering("callback QAudioDecoderControl::stateChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "stateChanged")
	if signal != nil {
		signal.(func(QAudioDecoder__State))(QAudioDecoder__State(state))
	}

}

func (ptr *QAudioDecoderControl) Stop() {
	defer qt.Recovering("QAudioDecoderControl::stop")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_Stop(ptr.Pointer())
	}
}

func (ptr *QAudioDecoderControl) DestroyQAudioDecoderControl() {
	defer qt.Recovering("QAudioDecoderControl::~QAudioDecoderControl")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_DestroyQAudioDecoderControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
