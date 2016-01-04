package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
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
func callbackQAudioOutputSelectorControlActiveOutputChanged(ptr unsafe.Pointer, ptrName *C.char, name *C.char) {
	defer qt.Recovering("callback QAudioOutputSelectorControl::activeOutputChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "activeOutputChanged"); signal != nil {
		signal.(func(string))(C.GoString(name))
	}

}

func (ptr *QAudioOutputSelectorControl) ActiveOutputChanged(name string) {
	defer qt.Recovering("QAudioOutputSelectorControl::activeOutputChanged")

	if ptr.Pointer() != nil {
		C.QAudioOutputSelectorControl_ActiveOutputChanged(ptr.Pointer(), C.CString(name))
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
func callbackQAudioOutputSelectorControlAvailableOutputsChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAudioOutputSelectorControl::availableOutputsChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "availableOutputsChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAudioOutputSelectorControl) AvailableOutputsChanged() {
	defer qt.Recovering("QAudioOutputSelectorControl::availableOutputsChanged")

	if ptr.Pointer() != nil {
		C.QAudioOutputSelectorControl_AvailableOutputsChanged(ptr.Pointer())
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

func (ptr *QAudioOutputSelectorControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAudioOutputSelectorControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAudioOutputSelectorControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAudioOutputSelectorControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAudioOutputSelectorControlTimerEvent
func callbackQAudioOutputSelectorControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAudioOutputSelectorControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAudioOutputSelectorControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAudioOutputSelectorControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAudioOutputSelectorControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QAudioOutputSelectorControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAudioOutputSelectorControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAudioOutputSelectorControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QAudioOutputSelectorControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAudioOutputSelectorControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAudioOutputSelectorControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAudioOutputSelectorControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAudioOutputSelectorControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAudioOutputSelectorControlChildEvent
func callbackQAudioOutputSelectorControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAudioOutputSelectorControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAudioOutputSelectorControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAudioOutputSelectorControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAudioOutputSelectorControl::childEvent")

	if ptr.Pointer() != nil {
		C.QAudioOutputSelectorControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAudioOutputSelectorControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAudioOutputSelectorControl::childEvent")

	if ptr.Pointer() != nil {
		C.QAudioOutputSelectorControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAudioOutputSelectorControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAudioOutputSelectorControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAudioOutputSelectorControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAudioOutputSelectorControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAudioOutputSelectorControlCustomEvent
func callbackQAudioOutputSelectorControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAudioOutputSelectorControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAudioOutputSelectorControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAudioOutputSelectorControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAudioOutputSelectorControl::customEvent")

	if ptr.Pointer() != nil {
		C.QAudioOutputSelectorControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAudioOutputSelectorControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAudioOutputSelectorControl::customEvent")

	if ptr.Pointer() != nil {
		C.QAudioOutputSelectorControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
