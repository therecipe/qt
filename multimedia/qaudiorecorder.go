package multimedia

//#include "multimedia.h"
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

type QAudioRecorder_ITF interface {
	QMediaRecorder_ITF
	QAudioRecorder_PTR() *QAudioRecorder
}

func PointerFromQAudioRecorder(ptr QAudioRecorder_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAudioRecorder_PTR().Pointer()
	}
	return nil
}

func NewQAudioRecorderFromPointer(ptr unsafe.Pointer) *QAudioRecorder {
	var n = new(QAudioRecorder)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAudioRecorder_") {
		n.SetObjectName("QAudioRecorder_" + qt.Identifier())
	}
	return n
}

func (ptr *QAudioRecorder) QAudioRecorder_PTR() *QAudioRecorder {
	return ptr
}

func NewQAudioRecorder(parent core.QObject_ITF) *QAudioRecorder {
	defer qt.Recovering("QAudioRecorder::QAudioRecorder")

	return NewQAudioRecorderFromPointer(C.QAudioRecorder_NewQAudioRecorder(core.PointerFromQObject(parent)))
}

func (ptr *QAudioRecorder) AudioInput() string {
	defer qt.Recovering("QAudioRecorder::audioInput")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioRecorder_AudioInput(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAudioRecorder) ConnectAudioInputChanged(f func(name string)) {
	defer qt.Recovering("connect QAudioRecorder::audioInputChanged")

	if ptr.Pointer() != nil {
		C.QAudioRecorder_ConnectAudioInputChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "audioInputChanged", f)
	}
}

func (ptr *QAudioRecorder) DisconnectAudioInputChanged() {
	defer qt.Recovering("disconnect QAudioRecorder::audioInputChanged")

	if ptr.Pointer() != nil {
		C.QAudioRecorder_DisconnectAudioInputChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "audioInputChanged")
	}
}

//export callbackQAudioRecorderAudioInputChanged
func callbackQAudioRecorderAudioInputChanged(ptr unsafe.Pointer, ptrName *C.char, name *C.char) {
	defer qt.Recovering("callback QAudioRecorder::audioInputChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "audioInputChanged"); signal != nil {
		signal.(func(string))(C.GoString(name))
	}

}

func (ptr *QAudioRecorder) AudioInputChanged(name string) {
	defer qt.Recovering("QAudioRecorder::audioInputChanged")

	if ptr.Pointer() != nil {
		C.QAudioRecorder_AudioInputChanged(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QAudioRecorder) AudioInputDescription(name string) string {
	defer qt.Recovering("QAudioRecorder::audioInputDescription")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioRecorder_AudioInputDescription(ptr.Pointer(), C.CString(name)))
	}
	return ""
}

func (ptr *QAudioRecorder) AudioInputs() []string {
	defer qt.Recovering("QAudioRecorder::audioInputs")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QAudioRecorder_AudioInputs(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QAudioRecorder) ConnectAvailableAudioInputsChanged(f func()) {
	defer qt.Recovering("connect QAudioRecorder::availableAudioInputsChanged")

	if ptr.Pointer() != nil {
		C.QAudioRecorder_ConnectAvailableAudioInputsChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "availableAudioInputsChanged", f)
	}
}

func (ptr *QAudioRecorder) DisconnectAvailableAudioInputsChanged() {
	defer qt.Recovering("disconnect QAudioRecorder::availableAudioInputsChanged")

	if ptr.Pointer() != nil {
		C.QAudioRecorder_DisconnectAvailableAudioInputsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "availableAudioInputsChanged")
	}
}

//export callbackQAudioRecorderAvailableAudioInputsChanged
func callbackQAudioRecorderAvailableAudioInputsChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAudioRecorder::availableAudioInputsChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "availableAudioInputsChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAudioRecorder) AvailableAudioInputsChanged() {
	defer qt.Recovering("QAudioRecorder::availableAudioInputsChanged")

	if ptr.Pointer() != nil {
		C.QAudioRecorder_AvailableAudioInputsChanged(ptr.Pointer())
	}
}

func (ptr *QAudioRecorder) DefaultAudioInput() string {
	defer qt.Recovering("QAudioRecorder::defaultAudioInput")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioRecorder_DefaultAudioInput(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAudioRecorder) SetAudioInput(name string) {
	defer qt.Recovering("QAudioRecorder::setAudioInput")

	if ptr.Pointer() != nil {
		C.QAudioRecorder_SetAudioInput(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QAudioRecorder) DestroyQAudioRecorder() {
	defer qt.Recovering("QAudioRecorder::~QAudioRecorder")

	if ptr.Pointer() != nil {
		C.QAudioRecorder_DestroyQAudioRecorder(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAudioRecorder) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAudioRecorder::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAudioRecorder) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAudioRecorder::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAudioRecorderTimerEvent
func callbackQAudioRecorderTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAudioRecorder::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAudioRecorderFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAudioRecorder) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAudioRecorder::timerEvent")

	if ptr.Pointer() != nil {
		C.QAudioRecorder_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAudioRecorder) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAudioRecorder::timerEvent")

	if ptr.Pointer() != nil {
		C.QAudioRecorder_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAudioRecorder) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAudioRecorder::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAudioRecorder) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAudioRecorder::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAudioRecorderChildEvent
func callbackQAudioRecorderChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAudioRecorder::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAudioRecorderFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAudioRecorder) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAudioRecorder::childEvent")

	if ptr.Pointer() != nil {
		C.QAudioRecorder_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAudioRecorder) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAudioRecorder::childEvent")

	if ptr.Pointer() != nil {
		C.QAudioRecorder_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAudioRecorder) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAudioRecorder::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAudioRecorder) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAudioRecorder::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAudioRecorderCustomEvent
func callbackQAudioRecorderCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAudioRecorder::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAudioRecorderFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAudioRecorder) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAudioRecorder::customEvent")

	if ptr.Pointer() != nil {
		C.QAudioRecorder_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAudioRecorder) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAudioRecorder::customEvent")

	if ptr.Pointer() != nil {
		C.QAudioRecorder_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
