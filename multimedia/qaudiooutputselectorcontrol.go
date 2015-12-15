package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QAudioOutputSelectorControl struct {
	QMediaControl
}

type QAudioOutputSelectorControl_ITF interface {
	QMediaControl_ITF
	QAudioOutputSelectorControl_PTR() *QAudioOutputSelectorControl
}

func PointerFromQAudioOutputSelectorControl(ptr QAudioOutputSelectorControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAudioOutputSelectorControl_PTR().Pointer()
	}
	return nil
}

func NewQAudioOutputSelectorControlFromPointer(ptr unsafe.Pointer) *QAudioOutputSelectorControl {
	var n = new(QAudioOutputSelectorControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAudioOutputSelectorControl_") {
		n.SetObjectName("QAudioOutputSelectorControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QAudioOutputSelectorControl) QAudioOutputSelectorControl_PTR() *QAudioOutputSelectorControl {
	return ptr
}

func (ptr *QAudioOutputSelectorControl) ActiveOutput() string {
	defer qt.Recovering("QAudioOutputSelectorControl::activeOutput")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioOutputSelectorControl_ActiveOutput(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAudioOutputSelectorControl) ConnectActiveOutputChanged(f func(name string)) {
	defer qt.Recovering("connect QAudioOutputSelectorControl::activeOutputChanged")

	if ptr.Pointer() != nil {
		C.QAudioOutputSelectorControl_ConnectActiveOutputChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "activeOutputChanged", f)
	}
}

func (ptr *QAudioOutputSelectorControl) DisconnectActiveOutputChanged() {
	defer qt.Recovering("disconnect QAudioOutputSelectorControl::activeOutputChanged")

	if ptr.Pointer() != nil {
		C.QAudioOutputSelectorControl_DisconnectActiveOutputChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "activeOutputChanged")
	}
}

//export callbackQAudioOutputSelectorControlActiveOutputChanged
func callbackQAudioOutputSelectorControlActiveOutputChanged(ptrName *C.char, name *C.char) {
	defer qt.Recovering("callback QAudioOutputSelectorControl::activeOutputChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "activeOutputChanged")
	if signal != nil {
		signal.(func(string))(C.GoString(name))
	}

}

func (ptr *QAudioOutputSelectorControl) ConnectAvailableOutputsChanged(f func()) {
	defer qt.Recovering("connect QAudioOutputSelectorControl::availableOutputsChanged")

	if ptr.Pointer() != nil {
		C.QAudioOutputSelectorControl_ConnectAvailableOutputsChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "availableOutputsChanged", f)
	}
}

func (ptr *QAudioOutputSelectorControl) DisconnectAvailableOutputsChanged() {
	defer qt.Recovering("disconnect QAudioOutputSelectorControl::availableOutputsChanged")

	if ptr.Pointer() != nil {
		C.QAudioOutputSelectorControl_DisconnectAvailableOutputsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "availableOutputsChanged")
	}
}

//export callbackQAudioOutputSelectorControlAvailableOutputsChanged
func callbackQAudioOutputSelectorControlAvailableOutputsChanged(ptrName *C.char) {
	defer qt.Recovering("callback QAudioOutputSelectorControl::availableOutputsChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "availableOutputsChanged")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QAudioOutputSelectorControl) DefaultOutput() string {
	defer qt.Recovering("QAudioOutputSelectorControl::defaultOutput")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioOutputSelectorControl_DefaultOutput(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAudioOutputSelectorControl) OutputDescription(name string) string {
	defer qt.Recovering("QAudioOutputSelectorControl::outputDescription")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioOutputSelectorControl_OutputDescription(ptr.Pointer(), C.CString(name)))
	}
	return ""
}

func (ptr *QAudioOutputSelectorControl) SetActiveOutput(name string) {
	defer qt.Recovering("QAudioOutputSelectorControl::setActiveOutput")

	if ptr.Pointer() != nil {
		C.QAudioOutputSelectorControl_SetActiveOutput(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QAudioOutputSelectorControl) DestroyQAudioOutputSelectorControl() {
	defer qt.Recovering("QAudioOutputSelectorControl::~QAudioOutputSelectorControl")

	if ptr.Pointer() != nil {
		C.QAudioOutputSelectorControl_DestroyQAudioOutputSelectorControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
