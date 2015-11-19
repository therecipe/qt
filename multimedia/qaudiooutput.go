package multimedia

//#include "qaudiooutput.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAudioOutput struct {
	core.QObject
}

type QAudioOutput_ITF interface {
	core.QObject_ITF
	QAudioOutput_PTR() *QAudioOutput
}

func PointerFromQAudioOutput(ptr QAudioOutput_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAudioOutput_PTR().Pointer()
	}
	return nil
}

func NewQAudioOutputFromPointer(ptr unsafe.Pointer) *QAudioOutput {
	var n = new(QAudioOutput)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QAudioOutput_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAudioOutput) QAudioOutput_PTR() *QAudioOutput {
	return ptr
}

func NewQAudioOutput2(audioDevice QAudioDeviceInfo_ITF, format QAudioFormat_ITF, parent core.QObject_ITF) *QAudioOutput {
	return NewQAudioOutputFromPointer(C.QAudioOutput_NewQAudioOutput2(PointerFromQAudioDeviceInfo(audioDevice), PointerFromQAudioFormat(format), core.PointerFromQObject(parent)))
}

func NewQAudioOutput(format QAudioFormat_ITF, parent core.QObject_ITF) *QAudioOutput {
	return NewQAudioOutputFromPointer(C.QAudioOutput_NewQAudioOutput(PointerFromQAudioFormat(format), core.PointerFromQObject(parent)))
}

func (ptr *QAudioOutput) BufferSize() int {
	if ptr.Pointer() != nil {
		return int(C.QAudioOutput_BufferSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioOutput) BytesFree() int {
	if ptr.Pointer() != nil {
		return int(C.QAudioOutput_BytesFree(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioOutput) Category() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioOutput_Category(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAudioOutput) ConnectNotify(f func()) {
	if ptr.Pointer() != nil {
		C.QAudioOutput_ConnectNotify(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "notify", f)
	}
}

func (ptr *QAudioOutput) DisconnectNotify() {
	if ptr.Pointer() != nil {
		C.QAudioOutput_DisconnectNotify(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "notify")
	}
}

//export callbackQAudioOutputNotify
func callbackQAudioOutputNotify(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "notify").(func())()
}

func (ptr *QAudioOutput) NotifyInterval() int {
	if ptr.Pointer() != nil {
		return int(C.QAudioOutput_NotifyInterval(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioOutput) PeriodSize() int {
	if ptr.Pointer() != nil {
		return int(C.QAudioOutput_PeriodSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioOutput) Reset() {
	if ptr.Pointer() != nil {
		C.QAudioOutput_Reset(ptr.Pointer())
	}
}

func (ptr *QAudioOutput) Resume() {
	if ptr.Pointer() != nil {
		C.QAudioOutput_Resume(ptr.Pointer())
	}
}

func (ptr *QAudioOutput) SetBufferSize(value int) {
	if ptr.Pointer() != nil {
		C.QAudioOutput_SetBufferSize(ptr.Pointer(), C.int(value))
	}
}

func (ptr *QAudioOutput) SetCategory(category string) {
	if ptr.Pointer() != nil {
		C.QAudioOutput_SetCategory(ptr.Pointer(), C.CString(category))
	}
}

func (ptr *QAudioOutput) SetNotifyInterval(ms int) {
	if ptr.Pointer() != nil {
		C.QAudioOutput_SetNotifyInterval(ptr.Pointer(), C.int(ms))
	}
}

func (ptr *QAudioOutput) SetVolume(volume float64) {
	if ptr.Pointer() != nil {
		C.QAudioOutput_SetVolume(ptr.Pointer(), C.double(volume))
	}
}

func (ptr *QAudioOutput) Start2() *core.QIODevice {
	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QAudioOutput_Start2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAudioOutput) Start(device core.QIODevice_ITF) {
	if ptr.Pointer() != nil {
		C.QAudioOutput_Start(ptr.Pointer(), core.PointerFromQIODevice(device))
	}
}

func (ptr *QAudioOutput) Stop() {
	if ptr.Pointer() != nil {
		C.QAudioOutput_Stop(ptr.Pointer())
	}
}

func (ptr *QAudioOutput) Suspend() {
	if ptr.Pointer() != nil {
		C.QAudioOutput_Suspend(ptr.Pointer())
	}
}

func (ptr *QAudioOutput) Volume() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QAudioOutput_Volume(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioOutput) DestroyQAudioOutput() {
	if ptr.Pointer() != nil {
		C.QAudioOutput_DestroyQAudioOutput(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
