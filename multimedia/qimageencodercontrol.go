package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QImageEncoderControl struct {
	QMediaControl
}

type QImageEncoderControl_ITF interface {
	QMediaControl_ITF
	QImageEncoderControl_PTR() *QImageEncoderControl
}

func PointerFromQImageEncoderControl(ptr QImageEncoderControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QImageEncoderControl_PTR().Pointer()
	}
	return nil
}

func NewQImageEncoderControlFromPointer(ptr unsafe.Pointer) *QImageEncoderControl {
	var n = new(QImageEncoderControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QImageEncoderControl_") {
		n.SetObjectName("QImageEncoderControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QImageEncoderControl) QImageEncoderControl_PTR() *QImageEncoderControl {
	return ptr
}

func (ptr *QImageEncoderControl) ImageCodecDescription(codec string) string {
	defer qt.Recovering("QImageEncoderControl::imageCodecDescription")

	if ptr.Pointer() != nil {
		return C.GoString(C.QImageEncoderControl_ImageCodecDescription(ptr.Pointer(), C.CString(codec)))
	}
	return ""
}

func (ptr *QImageEncoderControl) SetImageSettings(settings QImageEncoderSettings_ITF) {
	defer qt.Recovering("QImageEncoderControl::setImageSettings")

	if ptr.Pointer() != nil {
		C.QImageEncoderControl_SetImageSettings(ptr.Pointer(), PointerFromQImageEncoderSettings(settings))
	}
}

func (ptr *QImageEncoderControl) SupportedImageCodecs() []string {
	defer qt.Recovering("QImageEncoderControl::supportedImageCodecs")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QImageEncoderControl_SupportedImageCodecs(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QImageEncoderControl) DestroyQImageEncoderControl() {
	defer qt.Recovering("QImageEncoderControl::~QImageEncoderControl")

	if ptr.Pointer() != nil {
		C.QImageEncoderControl_DestroyQImageEncoderControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QImageEncoderControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QImageEncoderControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QImageEncoderControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QImageEncoderControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQImageEncoderControlTimerEvent
func callbackQImageEncoderControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QImageEncoderControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQImageEncoderControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QImageEncoderControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QImageEncoderControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QImageEncoderControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QImageEncoderControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QImageEncoderControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QImageEncoderControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QImageEncoderControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QImageEncoderControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QImageEncoderControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QImageEncoderControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQImageEncoderControlChildEvent
func callbackQImageEncoderControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QImageEncoderControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQImageEncoderControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QImageEncoderControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QImageEncoderControl::childEvent")

	if ptr.Pointer() != nil {
		C.QImageEncoderControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QImageEncoderControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QImageEncoderControl::childEvent")

	if ptr.Pointer() != nil {
		C.QImageEncoderControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QImageEncoderControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QImageEncoderControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QImageEncoderControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QImageEncoderControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQImageEncoderControlCustomEvent
func callbackQImageEncoderControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QImageEncoderControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQImageEncoderControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QImageEncoderControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QImageEncoderControl::customEvent")

	if ptr.Pointer() != nil {
		C.QImageEncoderControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QImageEncoderControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QImageEncoderControl::customEvent")

	if ptr.Pointer() != nil {
		C.QImageEncoderControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
