package multimedia

//#include "qaudioinputselectorcontrol.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QAudioInputSelectorControl struct {
	QMediaControl
}

type QAudioInputSelectorControl_ITF interface {
	QMediaControl_ITF
	QAudioInputSelectorControl_PTR() *QAudioInputSelectorControl
}

func PointerFromQAudioInputSelectorControl(ptr QAudioInputSelectorControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAudioInputSelectorControl_PTR().Pointer()
	}
	return nil
}

func NewQAudioInputSelectorControlFromPointer(ptr unsafe.Pointer) *QAudioInputSelectorControl {
	var n = new(QAudioInputSelectorControl)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QAudioInputSelectorControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAudioInputSelectorControl) QAudioInputSelectorControl_PTR() *QAudioInputSelectorControl {
	return ptr
}

func (ptr *QAudioInputSelectorControl) ActiveInput() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioInputSelectorControl_ActiveInput(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAudioInputSelectorControl) ConnectActiveInputChanged(f func(name string)) {
	if ptr.Pointer() != nil {
		C.QAudioInputSelectorControl_ConnectActiveInputChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "activeInputChanged", f)
	}
}

func (ptr *QAudioInputSelectorControl) DisconnectActiveInputChanged() {
	if ptr.Pointer() != nil {
		C.QAudioInputSelectorControl_DisconnectActiveInputChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "activeInputChanged")
	}
}

//export callbackQAudioInputSelectorControlActiveInputChanged
func callbackQAudioInputSelectorControlActiveInputChanged(ptrName *C.char, name *C.char) {
	qt.GetSignal(C.GoString(ptrName), "activeInputChanged").(func(string))(C.GoString(name))
}

func (ptr *QAudioInputSelectorControl) ConnectAvailableInputsChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QAudioInputSelectorControl_ConnectAvailableInputsChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "availableInputsChanged", f)
	}
}

func (ptr *QAudioInputSelectorControl) DisconnectAvailableInputsChanged() {
	if ptr.Pointer() != nil {
		C.QAudioInputSelectorControl_DisconnectAvailableInputsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "availableInputsChanged")
	}
}

//export callbackQAudioInputSelectorControlAvailableInputsChanged
func callbackQAudioInputSelectorControlAvailableInputsChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "availableInputsChanged").(func())()
}

func (ptr *QAudioInputSelectorControl) DefaultInput() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioInputSelectorControl_DefaultInput(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAudioInputSelectorControl) InputDescription(name string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioInputSelectorControl_InputDescription(ptr.Pointer(), C.CString(name)))
	}
	return ""
}

func (ptr *QAudioInputSelectorControl) SetActiveInput(name string) {
	if ptr.Pointer() != nil {
		C.QAudioInputSelectorControl_SetActiveInput(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QAudioInputSelectorControl) DestroyQAudioInputSelectorControl() {
	if ptr.Pointer() != nil {
		C.QAudioInputSelectorControl_DestroyQAudioInputSelectorControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
