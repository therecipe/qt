package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QAudioEncoderSettingsControl struct {
	QMediaControl
}

type QAudioEncoderSettingsControl_ITF interface {
	QMediaControl_ITF
	QAudioEncoderSettingsControl_PTR() *QAudioEncoderSettingsControl
}

func PointerFromQAudioEncoderSettingsControl(ptr QAudioEncoderSettingsControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAudioEncoderSettingsControl_PTR().Pointer()
	}
	return nil
}

func NewQAudioEncoderSettingsControlFromPointer(ptr unsafe.Pointer) *QAudioEncoderSettingsControl {
	var n = new(QAudioEncoderSettingsControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAudioEncoderSettingsControl_") {
		n.SetObjectName("QAudioEncoderSettingsControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QAudioEncoderSettingsControl) QAudioEncoderSettingsControl_PTR() *QAudioEncoderSettingsControl {
	return ptr
}

func (ptr *QAudioEncoderSettingsControl) CodecDescription(codec string) string {
	defer qt.Recovering("QAudioEncoderSettingsControl::codecDescription")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioEncoderSettingsControl_CodecDescription(ptr.Pointer(), C.CString(codec)))
	}
	return ""
}

func (ptr *QAudioEncoderSettingsControl) SetAudioSettings(settings QAudioEncoderSettings_ITF) {
	defer qt.Recovering("QAudioEncoderSettingsControl::setAudioSettings")

	if ptr.Pointer() != nil {
		C.QAudioEncoderSettingsControl_SetAudioSettings(ptr.Pointer(), PointerFromQAudioEncoderSettings(settings))
	}
}

func (ptr *QAudioEncoderSettingsControl) SupportedAudioCodecs() []string {
	defer qt.Recovering("QAudioEncoderSettingsControl::supportedAudioCodecs")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QAudioEncoderSettingsControl_SupportedAudioCodecs(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QAudioEncoderSettingsControl) DestroyQAudioEncoderSettingsControl() {
	defer qt.Recovering("QAudioEncoderSettingsControl::~QAudioEncoderSettingsControl")

	if ptr.Pointer() != nil {
		C.QAudioEncoderSettingsControl_DestroyQAudioEncoderSettingsControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAudioEncoderSettingsControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAudioEncoderSettingsControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAudioEncoderSettingsControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAudioEncoderSettingsControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAudioEncoderSettingsControlTimerEvent
func callbackQAudioEncoderSettingsControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAudioEncoderSettingsControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAudioEncoderSettingsControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAudioEncoderSettingsControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAudioEncoderSettingsControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QAudioEncoderSettingsControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAudioEncoderSettingsControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAudioEncoderSettingsControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QAudioEncoderSettingsControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAudioEncoderSettingsControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAudioEncoderSettingsControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAudioEncoderSettingsControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAudioEncoderSettingsControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAudioEncoderSettingsControlChildEvent
func callbackQAudioEncoderSettingsControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAudioEncoderSettingsControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAudioEncoderSettingsControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAudioEncoderSettingsControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAudioEncoderSettingsControl::childEvent")

	if ptr.Pointer() != nil {
		C.QAudioEncoderSettingsControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAudioEncoderSettingsControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAudioEncoderSettingsControl::childEvent")

	if ptr.Pointer() != nil {
		C.QAudioEncoderSettingsControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAudioEncoderSettingsControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAudioEncoderSettingsControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAudioEncoderSettingsControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAudioEncoderSettingsControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAudioEncoderSettingsControlCustomEvent
func callbackQAudioEncoderSettingsControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAudioEncoderSettingsControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAudioEncoderSettingsControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAudioEncoderSettingsControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAudioEncoderSettingsControl::customEvent")

	if ptr.Pointer() != nil {
		C.QAudioEncoderSettingsControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAudioEncoderSettingsControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAudioEncoderSettingsControl::customEvent")

	if ptr.Pointer() != nil {
		C.QAudioEncoderSettingsControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
