package multimedia

//#include "qaudioinput.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAudioInput struct {
	core.QObject
}

type QAudioInput_ITF interface {
	core.QObject_ITF
	QAudioInput_PTR() *QAudioInput
}

func PointerFromQAudioInput(ptr QAudioInput_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAudioInput_PTR().Pointer()
	}
	return nil
}

func NewQAudioInputFromPointer(ptr unsafe.Pointer) *QAudioInput {
	var n = new(QAudioInput)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QAudioInput_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAudioInput) QAudioInput_PTR() *QAudioInput {
	return ptr
}

func NewQAudioInput2(audioDevice QAudioDeviceInfo_ITF, format QAudioFormat_ITF, parent core.QObject_ITF) *QAudioInput {
	return NewQAudioInputFromPointer(C.QAudioInput_NewQAudioInput2(PointerFromQAudioDeviceInfo(audioDevice), PointerFromQAudioFormat(format), core.PointerFromQObject(parent)))
}

func NewQAudioInput(format QAudioFormat_ITF, parent core.QObject_ITF) *QAudioInput {
	return NewQAudioInputFromPointer(C.QAudioInput_NewQAudioInput(PointerFromQAudioFormat(format), core.PointerFromQObject(parent)))
}

func (ptr *QAudioInput) BufferSize() int {
	if ptr.Pointer() != nil {
		return int(C.QAudioInput_BufferSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioInput) BytesReady() int {
	if ptr.Pointer() != nil {
		return int(C.QAudioInput_BytesReady(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioInput) ConnectNotify(f func()) {
	if ptr.Pointer() != nil {
		C.QAudioInput_ConnectNotify(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "notify", f)
	}
}

func (ptr *QAudioInput) DisconnectNotify() {
	if ptr.Pointer() != nil {
		C.QAudioInput_DisconnectNotify(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "notify")
	}
}

//export callbackQAudioInputNotify
func callbackQAudioInputNotify(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "notify").(func())()
}

func (ptr *QAudioInput) NotifyInterval() int {
	if ptr.Pointer() != nil {
		return int(C.QAudioInput_NotifyInterval(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioInput) PeriodSize() int {
	if ptr.Pointer() != nil {
		return int(C.QAudioInput_PeriodSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioInput) Reset() {
	if ptr.Pointer() != nil {
		C.QAudioInput_Reset(ptr.Pointer())
	}
}

func (ptr *QAudioInput) Resume() {
	if ptr.Pointer() != nil {
		C.QAudioInput_Resume(ptr.Pointer())
	}
}

func (ptr *QAudioInput) SetBufferSize(value int) {
	if ptr.Pointer() != nil {
		C.QAudioInput_SetBufferSize(ptr.Pointer(), C.int(value))
	}
}

func (ptr *QAudioInput) SetNotifyInterval(ms int) {
	if ptr.Pointer() != nil {
		C.QAudioInput_SetNotifyInterval(ptr.Pointer(), C.int(ms))
	}
}

func (ptr *QAudioInput) SetVolume(volume float64) {
	if ptr.Pointer() != nil {
		C.QAudioInput_SetVolume(ptr.Pointer(), C.double(volume))
	}
}

func (ptr *QAudioInput) Start2() *core.QIODevice {
	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QAudioInput_Start2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAudioInput) Start(device core.QIODevice_ITF) {
	if ptr.Pointer() != nil {
		C.QAudioInput_Start(ptr.Pointer(), core.PointerFromQIODevice(device))
	}
}

func (ptr *QAudioInput) Stop() {
	if ptr.Pointer() != nil {
		C.QAudioInput_Stop(ptr.Pointer())
	}
}

func (ptr *QAudioInput) Suspend() {
	if ptr.Pointer() != nil {
		C.QAudioInput_Suspend(ptr.Pointer())
	}
}

func (ptr *QAudioInput) Volume() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QAudioInput_Volume(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioInput) DestroyQAudioInput() {
	if ptr.Pointer() != nil {
		C.QAudioInput_DestroyQAudioInput(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
