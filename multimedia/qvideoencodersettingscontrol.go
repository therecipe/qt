package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QVideoEncoderSettingsControl struct {
	QMediaControl
}

type QVideoEncoderSettingsControl_ITF interface {
	QMediaControl_ITF
	QVideoEncoderSettingsControl_PTR() *QVideoEncoderSettingsControl
}

func PointerFromQVideoEncoderSettingsControl(ptr QVideoEncoderSettingsControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVideoEncoderSettingsControl_PTR().Pointer()
	}
	return nil
}

func NewQVideoEncoderSettingsControlFromPointer(ptr unsafe.Pointer) *QVideoEncoderSettingsControl {
	var n = new(QVideoEncoderSettingsControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QVideoEncoderSettingsControl_") {
		n.SetObjectName("QVideoEncoderSettingsControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QVideoEncoderSettingsControl) QVideoEncoderSettingsControl_PTR() *QVideoEncoderSettingsControl {
	return ptr
}

func (ptr *QVideoEncoderSettingsControl) SetVideoSettings(settings QVideoEncoderSettings_ITF) {
	defer qt.Recovering("QVideoEncoderSettingsControl::setVideoSettings")

	if ptr.Pointer() != nil {
		C.QVideoEncoderSettingsControl_SetVideoSettings(ptr.Pointer(), PointerFromQVideoEncoderSettings(settings))
	}
}

func (ptr *QVideoEncoderSettingsControl) SupportedVideoCodecs() []string {
	defer qt.Recovering("QVideoEncoderSettingsControl::supportedVideoCodecs")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QVideoEncoderSettingsControl_SupportedVideoCodecs(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QVideoEncoderSettingsControl) VideoCodecDescription(codec string) string {
	defer qt.Recovering("QVideoEncoderSettingsControl::videoCodecDescription")

	if ptr.Pointer() != nil {
		return C.GoString(C.QVideoEncoderSettingsControl_VideoCodecDescription(ptr.Pointer(), C.CString(codec)))
	}
	return ""
}

func (ptr *QVideoEncoderSettingsControl) DestroyQVideoEncoderSettingsControl() {
	defer qt.Recovering("QVideoEncoderSettingsControl::~QVideoEncoderSettingsControl")

	if ptr.Pointer() != nil {
		C.QVideoEncoderSettingsControl_DestroyQVideoEncoderSettingsControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QVideoEncoderSettingsControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QVideoEncoderSettingsControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QVideoEncoderSettingsControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QVideoEncoderSettingsControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQVideoEncoderSettingsControlTimerEvent
func callbackQVideoEncoderSettingsControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoEncoderSettingsControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQVideoEncoderSettingsControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QVideoEncoderSettingsControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QVideoEncoderSettingsControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QVideoEncoderSettingsControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QVideoEncoderSettingsControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QVideoEncoderSettingsControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QVideoEncoderSettingsControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QVideoEncoderSettingsControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QVideoEncoderSettingsControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QVideoEncoderSettingsControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QVideoEncoderSettingsControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQVideoEncoderSettingsControlChildEvent
func callbackQVideoEncoderSettingsControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoEncoderSettingsControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQVideoEncoderSettingsControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QVideoEncoderSettingsControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QVideoEncoderSettingsControl::childEvent")

	if ptr.Pointer() != nil {
		C.QVideoEncoderSettingsControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QVideoEncoderSettingsControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QVideoEncoderSettingsControl::childEvent")

	if ptr.Pointer() != nil {
		C.QVideoEncoderSettingsControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QVideoEncoderSettingsControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QVideoEncoderSettingsControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QVideoEncoderSettingsControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QVideoEncoderSettingsControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQVideoEncoderSettingsControlCustomEvent
func callbackQVideoEncoderSettingsControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoEncoderSettingsControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQVideoEncoderSettingsControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QVideoEncoderSettingsControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QVideoEncoderSettingsControl::customEvent")

	if ptr.Pointer() != nil {
		C.QVideoEncoderSettingsControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QVideoEncoderSettingsControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QVideoEncoderSettingsControl::customEvent")

	if ptr.Pointer() != nil {
		C.QVideoEncoderSettingsControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
