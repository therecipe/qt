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

type QAudioInputITF interface {
	core.QObjectITF
	QAudioInputPTR() *QAudioInput
}

func PointerFromQAudioInput(ptr QAudioInputITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAudioInputPTR().Pointer()
	}
	return nil
}

func QAudioInputFromPointer(ptr unsafe.Pointer) *QAudioInput {
	var n = new(QAudioInput)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QAudioInput_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAudioInput) QAudioInputPTR() *QAudioInput {
	return ptr
}

func NewQAudioInput2(audioDevice QAudioDeviceInfoITF, format QAudioFormatITF, parent core.QObjectITF) *QAudioInput {
	return QAudioInputFromPointer(unsafe.Pointer(C.QAudioInput_NewQAudioInput2(C.QtObjectPtr(PointerFromQAudioDeviceInfo(audioDevice)), C.QtObjectPtr(PointerFromQAudioFormat(format)), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func NewQAudioInput(format QAudioFormatITF, parent core.QObjectITF) *QAudioInput {
	return QAudioInputFromPointer(unsafe.Pointer(C.QAudioInput_NewQAudioInput(C.QtObjectPtr(PointerFromQAudioFormat(format)), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QAudioInput) BufferSize() int {
	if ptr.Pointer() != nil {
		return int(C.QAudioInput_BufferSize(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAudioInput) BytesReady() int {
	if ptr.Pointer() != nil {
		return int(C.QAudioInput_BytesReady(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAudioInput) ConnectNotify(f func()) {
	if ptr.Pointer() != nil {
		C.QAudioInput_ConnectNotify(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "notify", f)
	}
}

func (ptr *QAudioInput) DisconnectNotify() {
	if ptr.Pointer() != nil {
		C.QAudioInput_DisconnectNotify(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "notify")
	}
}

//export callbackQAudioInputNotify
func callbackQAudioInputNotify(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "notify").(func())()
}

func (ptr *QAudioInput) NotifyInterval() int {
	if ptr.Pointer() != nil {
		return int(C.QAudioInput_NotifyInterval(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAudioInput) PeriodSize() int {
	if ptr.Pointer() != nil {
		return int(C.QAudioInput_PeriodSize(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAudioInput) Reset() {
	if ptr.Pointer() != nil {
		C.QAudioInput_Reset(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QAudioInput) Resume() {
	if ptr.Pointer() != nil {
		C.QAudioInput_Resume(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QAudioInput) SetBufferSize(value int) {
	if ptr.Pointer() != nil {
		C.QAudioInput_SetBufferSize(C.QtObjectPtr(ptr.Pointer()), C.int(value))
	}
}

func (ptr *QAudioInput) SetNotifyInterval(ms int) {
	if ptr.Pointer() != nil {
		C.QAudioInput_SetNotifyInterval(C.QtObjectPtr(ptr.Pointer()), C.int(ms))
	}
}

func (ptr *QAudioInput) Start2() *core.QIODevice {
	if ptr.Pointer() != nil {
		return core.QIODeviceFromPointer(unsafe.Pointer(C.QAudioInput_Start2(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QAudioInput) Start(device core.QIODeviceITF) {
	if ptr.Pointer() != nil {
		C.QAudioInput_Start(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQIODevice(device)))
	}
}

func (ptr *QAudioInput) Stop() {
	if ptr.Pointer() != nil {
		C.QAudioInput_Stop(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QAudioInput) Suspend() {
	if ptr.Pointer() != nil {
		C.QAudioInput_Suspend(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QAudioInput) DestroyQAudioInput() {
	if ptr.Pointer() != nil {
		C.QAudioInput_DestroyQAudioInput(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
