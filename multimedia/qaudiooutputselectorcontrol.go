package multimedia

//#include "qaudiooutputselectorcontrol.h"
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
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QAudioOutputSelectorControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAudioOutputSelectorControl) QAudioOutputSelectorControl_PTR() *QAudioOutputSelectorControl {
	return ptr
}

func (ptr *QAudioOutputSelectorControl) ActiveOutput() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioOutputSelectorControl_ActiveOutput(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAudioOutputSelectorControl) ConnectActiveOutputChanged(f func(name string)) {
	if ptr.Pointer() != nil {
		C.QAudioOutputSelectorControl_ConnectActiveOutputChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "activeOutputChanged", f)
	}
}

func (ptr *QAudioOutputSelectorControl) DisconnectActiveOutputChanged() {
	if ptr.Pointer() != nil {
		C.QAudioOutputSelectorControl_DisconnectActiveOutputChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "activeOutputChanged")
	}
}

//export callbackQAudioOutputSelectorControlActiveOutputChanged
func callbackQAudioOutputSelectorControlActiveOutputChanged(ptrName *C.char, name *C.char) {
	qt.GetSignal(C.GoString(ptrName), "activeOutputChanged").(func(string))(C.GoString(name))
}

func (ptr *QAudioOutputSelectorControl) ConnectAvailableOutputsChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QAudioOutputSelectorControl_ConnectAvailableOutputsChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "availableOutputsChanged", f)
	}
}

func (ptr *QAudioOutputSelectorControl) DisconnectAvailableOutputsChanged() {
	if ptr.Pointer() != nil {
		C.QAudioOutputSelectorControl_DisconnectAvailableOutputsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "availableOutputsChanged")
	}
}

//export callbackQAudioOutputSelectorControlAvailableOutputsChanged
func callbackQAudioOutputSelectorControlAvailableOutputsChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "availableOutputsChanged").(func())()
}

func (ptr *QAudioOutputSelectorControl) DefaultOutput() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioOutputSelectorControl_DefaultOutput(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAudioOutputSelectorControl) OutputDescription(name string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioOutputSelectorControl_OutputDescription(ptr.Pointer(), C.CString(name)))
	}
	return ""
}

func (ptr *QAudioOutputSelectorControl) SetActiveOutput(name string) {
	if ptr.Pointer() != nil {
		C.QAudioOutputSelectorControl_SetActiveOutput(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QAudioOutputSelectorControl) DestroyQAudioOutputSelectorControl() {
	if ptr.Pointer() != nil {
		C.QAudioOutputSelectorControl_DestroyQAudioOutputSelectorControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
