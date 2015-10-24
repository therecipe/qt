package multimedia

//#include "qaudiorecorder.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QAudioRecorder struct {
	QMediaRecorder
}

type QAudioRecorderITF interface {
	QMediaRecorderITF
	QAudioRecorderPTR() *QAudioRecorder
}

func PointerFromQAudioRecorder(ptr QAudioRecorderITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAudioRecorderPTR().Pointer()
	}
	return nil
}

func QAudioRecorderFromPointer(ptr unsafe.Pointer) *QAudioRecorder {
	var n = new(QAudioRecorder)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QAudioRecorder_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAudioRecorder) QAudioRecorderPTR() *QAudioRecorder {
	return ptr
}

func NewQAudioRecorder(parent core.QObjectITF) *QAudioRecorder {
	return QAudioRecorderFromPointer(unsafe.Pointer(C.QAudioRecorder_NewQAudioRecorder(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QAudioRecorder) AudioInput() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioRecorder_AudioInput(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QAudioRecorder) ConnectAudioInputChanged(f func(name string)) {
	if ptr.Pointer() != nil {
		C.QAudioRecorder_ConnectAudioInputChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "audioInputChanged", f)
	}
}

func (ptr *QAudioRecorder) DisconnectAudioInputChanged() {
	if ptr.Pointer() != nil {
		C.QAudioRecorder_DisconnectAudioInputChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "audioInputChanged")
	}
}

//export callbackQAudioRecorderAudioInputChanged
func callbackQAudioRecorderAudioInputChanged(ptrName *C.char, name *C.char) {
	qt.GetSignal(C.GoString(ptrName), "audioInputChanged").(func(string))(C.GoString(name))
}

func (ptr *QAudioRecorder) AudioInputDescription(name string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioRecorder_AudioInputDescription(C.QtObjectPtr(ptr.Pointer()), C.CString(name)))
	}
	return ""
}

func (ptr *QAudioRecorder) AudioInputs() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QAudioRecorder_AudioInputs(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QAudioRecorder) ConnectAvailableAudioInputsChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QAudioRecorder_ConnectAvailableAudioInputsChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "availableAudioInputsChanged", f)
	}
}

func (ptr *QAudioRecorder) DisconnectAvailableAudioInputsChanged() {
	if ptr.Pointer() != nil {
		C.QAudioRecorder_DisconnectAvailableAudioInputsChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "availableAudioInputsChanged")
	}
}

//export callbackQAudioRecorderAvailableAudioInputsChanged
func callbackQAudioRecorderAvailableAudioInputsChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "availableAudioInputsChanged").(func())()
}

func (ptr *QAudioRecorder) DefaultAudioInput() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioRecorder_DefaultAudioInput(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QAudioRecorder) SetAudioInput(name string) {
	if ptr.Pointer() != nil {
		C.QAudioRecorder_SetAudioInput(C.QtObjectPtr(ptr.Pointer()), C.CString(name))
	}
}

func (ptr *QAudioRecorder) DestroyQAudioRecorder() {
	if ptr.Pointer() != nil {
		C.QAudioRecorder_DestroyQAudioRecorder(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
