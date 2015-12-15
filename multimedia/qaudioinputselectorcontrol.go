package multimedia

//#include "multimedia.h"
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
	for len(n.ObjectName()) < len("QAudioInputSelectorControl_") {
		n.SetObjectName("QAudioInputSelectorControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QAudioInputSelectorControl) QAudioInputSelectorControl_PTR() *QAudioInputSelectorControl {
	return ptr
}

func (ptr *QAudioInputSelectorControl) ActiveInput() string {
	defer qt.Recovering("QAudioInputSelectorControl::activeInput")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioInputSelectorControl_ActiveInput(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAudioInputSelectorControl) ConnectActiveInputChanged(f func(name string)) {
	defer qt.Recovering("connect QAudioInputSelectorControl::activeInputChanged")

	if ptr.Pointer() != nil {
		C.QAudioInputSelectorControl_ConnectActiveInputChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "activeInputChanged", f)
	}
}

func (ptr *QAudioInputSelectorControl) DisconnectActiveInputChanged() {
	defer qt.Recovering("disconnect QAudioInputSelectorControl::activeInputChanged")

	if ptr.Pointer() != nil {
		C.QAudioInputSelectorControl_DisconnectActiveInputChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "activeInputChanged")
	}
}

//export callbackQAudioInputSelectorControlActiveInputChanged
func callbackQAudioInputSelectorControlActiveInputChanged(ptrName *C.char, name *C.char) {
	defer qt.Recovering("callback QAudioInputSelectorControl::activeInputChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "activeInputChanged")
	if signal != nil {
		signal.(func(string))(C.GoString(name))
	}

}

func (ptr *QAudioInputSelectorControl) ConnectAvailableInputsChanged(f func()) {
	defer qt.Recovering("connect QAudioInputSelectorControl::availableInputsChanged")

	if ptr.Pointer() != nil {
		C.QAudioInputSelectorControl_ConnectAvailableInputsChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "availableInputsChanged", f)
	}
}

func (ptr *QAudioInputSelectorControl) DisconnectAvailableInputsChanged() {
	defer qt.Recovering("disconnect QAudioInputSelectorControl::availableInputsChanged")

	if ptr.Pointer() != nil {
		C.QAudioInputSelectorControl_DisconnectAvailableInputsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "availableInputsChanged")
	}
}

//export callbackQAudioInputSelectorControlAvailableInputsChanged
func callbackQAudioInputSelectorControlAvailableInputsChanged(ptrName *C.char) {
	defer qt.Recovering("callback QAudioInputSelectorControl::availableInputsChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "availableInputsChanged")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QAudioInputSelectorControl) DefaultInput() string {
	defer qt.Recovering("QAudioInputSelectorControl::defaultInput")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioInputSelectorControl_DefaultInput(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAudioInputSelectorControl) InputDescription(name string) string {
	defer qt.Recovering("QAudioInputSelectorControl::inputDescription")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioInputSelectorControl_InputDescription(ptr.Pointer(), C.CString(name)))
	}
	return ""
}

func (ptr *QAudioInputSelectorControl) SetActiveInput(name string) {
	defer qt.Recovering("QAudioInputSelectorControl::setActiveInput")

	if ptr.Pointer() != nil {
		C.QAudioInputSelectorControl_SetActiveInput(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QAudioInputSelectorControl) DestroyQAudioInputSelectorControl() {
	defer qt.Recovering("QAudioInputSelectorControl::~QAudioInputSelectorControl")

	if ptr.Pointer() != nil {
		C.QAudioInputSelectorControl_DestroyQAudioInputSelectorControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
