package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QMediaControl struct {
	core.QObject
}

type QMediaControl_ITF interface {
	core.QObject_ITF
	QMediaControl_PTR() *QMediaControl
}

func PointerFromQMediaControl(ptr QMediaControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaControl_PTR().Pointer()
	}
	return nil
}

func NewQMediaControlFromPointer(ptr unsafe.Pointer) *QMediaControl {
	var n = new(QMediaControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMediaControl_") {
		n.SetObjectName("QMediaControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QMediaControl) QMediaControl_PTR() *QMediaControl {
	return ptr
}

func (ptr *QMediaControl) DestroyQMediaControl() {
	defer qt.Recovering("QMediaControl::~QMediaControl")

	if ptr.Pointer() != nil {
		C.QMediaControl_DestroyQMediaControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMediaControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMediaControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMediaControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMediaControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMediaControlTimerEvent
func callbackQMediaControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMediaControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMediaControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMediaControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QMediaControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMediaControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMediaControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QMediaControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMediaControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMediaControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMediaControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMediaControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMediaControlChildEvent
func callbackQMediaControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMediaControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMediaControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMediaControl::childEvent")

	if ptr.Pointer() != nil {
		C.QMediaControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMediaControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMediaControl::childEvent")

	if ptr.Pointer() != nil {
		C.QMediaControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMediaControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMediaControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMediaControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMediaControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMediaControlCustomEvent
func callbackQMediaControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMediaControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMediaControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMediaControl::customEvent")

	if ptr.Pointer() != nil {
		C.QMediaControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMediaControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMediaControl::customEvent")

	if ptr.Pointer() != nil {
		C.QMediaControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
