package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QMediaAudioProbeControl struct {
	QMediaControl
}

type QMediaAudioProbeControl_ITF interface {
	QMediaControl_ITF
	QMediaAudioProbeControl_PTR() *QMediaAudioProbeControl
}

func PointerFromQMediaAudioProbeControl(ptr QMediaAudioProbeControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaAudioProbeControl_PTR().Pointer()
	}
	return nil
}

func NewQMediaAudioProbeControlFromPointer(ptr unsafe.Pointer) *QMediaAudioProbeControl {
	var n = new(QMediaAudioProbeControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMediaAudioProbeControl_") {
		n.SetObjectName("QMediaAudioProbeControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QMediaAudioProbeControl) QMediaAudioProbeControl_PTR() *QMediaAudioProbeControl {
	return ptr
}

func (ptr *QMediaAudioProbeControl) ConnectFlush(f func()) {
	defer qt.Recovering("connect QMediaAudioProbeControl::flush")

	if ptr.Pointer() != nil {
		C.QMediaAudioProbeControl_ConnectFlush(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "flush", f)
	}
}

func (ptr *QMediaAudioProbeControl) DisconnectFlush() {
	defer qt.Recovering("disconnect QMediaAudioProbeControl::flush")

	if ptr.Pointer() != nil {
		C.QMediaAudioProbeControl_DisconnectFlush(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "flush")
	}
}

//export callbackQMediaAudioProbeControlFlush
func callbackQMediaAudioProbeControlFlush(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QMediaAudioProbeControl::flush")

	if signal := qt.GetSignal(C.GoString(ptrName), "flush"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QMediaAudioProbeControl) Flush() {
	defer qt.Recovering("QMediaAudioProbeControl::flush")

	if ptr.Pointer() != nil {
		C.QMediaAudioProbeControl_Flush(ptr.Pointer())
	}
}

func (ptr *QMediaAudioProbeControl) DestroyQMediaAudioProbeControl() {
	defer qt.Recovering("QMediaAudioProbeControl::~QMediaAudioProbeControl")

	if ptr.Pointer() != nil {
		C.QMediaAudioProbeControl_DestroyQMediaAudioProbeControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMediaAudioProbeControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMediaAudioProbeControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMediaAudioProbeControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMediaAudioProbeControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMediaAudioProbeControlTimerEvent
func callbackQMediaAudioProbeControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaAudioProbeControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMediaAudioProbeControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMediaAudioProbeControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMediaAudioProbeControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QMediaAudioProbeControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMediaAudioProbeControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMediaAudioProbeControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QMediaAudioProbeControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMediaAudioProbeControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMediaAudioProbeControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMediaAudioProbeControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMediaAudioProbeControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMediaAudioProbeControlChildEvent
func callbackQMediaAudioProbeControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaAudioProbeControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMediaAudioProbeControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMediaAudioProbeControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMediaAudioProbeControl::childEvent")

	if ptr.Pointer() != nil {
		C.QMediaAudioProbeControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMediaAudioProbeControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMediaAudioProbeControl::childEvent")

	if ptr.Pointer() != nil {
		C.QMediaAudioProbeControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMediaAudioProbeControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMediaAudioProbeControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMediaAudioProbeControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMediaAudioProbeControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMediaAudioProbeControlCustomEvent
func callbackQMediaAudioProbeControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaAudioProbeControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMediaAudioProbeControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMediaAudioProbeControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMediaAudioProbeControl::customEvent")

	if ptr.Pointer() != nil {
		C.QMediaAudioProbeControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMediaAudioProbeControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMediaAudioProbeControl::customEvent")

	if ptr.Pointer() != nil {
		C.QMediaAudioProbeControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
