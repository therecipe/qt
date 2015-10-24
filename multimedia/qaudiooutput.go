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

type QAudioOutputITF interface {
	core.QObjectITF
	QAudioOutputPTR() *QAudioOutput
}

func PointerFromQAudioOutput(ptr QAudioOutputITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAudioOutputPTR().Pointer()
	}
	return nil
}

func QAudioOutputFromPointer(ptr unsafe.Pointer) *QAudioOutput {
	var n = new(QAudioOutput)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QAudioOutput_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAudioOutput) QAudioOutputPTR() *QAudioOutput {
	return ptr
}

func NewQAudioOutput2(audioDevice QAudioDeviceInfoITF, format QAudioFormatITF, parent core.QObjectITF) *QAudioOutput {
	return QAudioOutputFromPointer(unsafe.Pointer(C.QAudioOutput_NewQAudioOutput2(C.QtObjectPtr(PointerFromQAudioDeviceInfo(audioDevice)), C.QtObjectPtr(PointerFromQAudioFormat(format)), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func NewQAudioOutput(format QAudioFormatITF, parent core.QObjectITF) *QAudioOutput {
	return QAudioOutputFromPointer(unsafe.Pointer(C.QAudioOutput_NewQAudioOutput(C.QtObjectPtr(PointerFromQAudioFormat(format)), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QAudioOutput) BufferSize() int {
	if ptr.Pointer() != nil {
		return int(C.QAudioOutput_BufferSize(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAudioOutput) BytesFree() int {
	if ptr.Pointer() != nil {
		return int(C.QAudioOutput_BytesFree(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAudioOutput) Category() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioOutput_Category(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QAudioOutput) ConnectNotify(f func()) {
	if ptr.Pointer() != nil {
		C.QAudioOutput_ConnectNotify(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "notify", f)
	}
}

func (ptr *QAudioOutput) DisconnectNotify() {
	if ptr.Pointer() != nil {
		C.QAudioOutput_DisconnectNotify(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "notify")
	}
}

//export callbackQAudioOutputNotify
func callbackQAudioOutputNotify(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "notify").(func())()
}

func (ptr *QAudioOutput) NotifyInterval() int {
	if ptr.Pointer() != nil {
		return int(C.QAudioOutput_NotifyInterval(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAudioOutput) PeriodSize() int {
	if ptr.Pointer() != nil {
		return int(C.QAudioOutput_PeriodSize(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAudioOutput) Reset() {
	if ptr.Pointer() != nil {
		C.QAudioOutput_Reset(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QAudioOutput) Resume() {
	if ptr.Pointer() != nil {
		C.QAudioOutput_Resume(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QAudioOutput) SetBufferSize(value int) {
	if ptr.Pointer() != nil {
		C.QAudioOutput_SetBufferSize(C.QtObjectPtr(ptr.Pointer()), C.int(value))
	}
}

func (ptr *QAudioOutput) SetCategory(category string) {
	if ptr.Pointer() != nil {
		C.QAudioOutput_SetCategory(C.QtObjectPtr(ptr.Pointer()), C.CString(category))
	}
}

func (ptr *QAudioOutput) SetNotifyInterval(ms int) {
	if ptr.Pointer() != nil {
		C.QAudioOutput_SetNotifyInterval(C.QtObjectPtr(ptr.Pointer()), C.int(ms))
	}
}

func (ptr *QAudioOutput) Start2() *core.QIODevice {
	if ptr.Pointer() != nil {
		return core.QIODeviceFromPointer(unsafe.Pointer(C.QAudioOutput_Start2(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QAudioOutput) Start(device core.QIODeviceITF) {
	if ptr.Pointer() != nil {
		C.QAudioOutput_Start(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQIODevice(device)))
	}
}

func (ptr *QAudioOutput) Stop() {
	if ptr.Pointer() != nil {
		C.QAudioOutput_Stop(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QAudioOutput) Suspend() {
	if ptr.Pointer() != nil {
		C.QAudioOutput_Suspend(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QAudioOutput) DestroyQAudioOutput() {
	if ptr.Pointer() != nil {
		C.QAudioOutput_DestroyQAudioOutput(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
