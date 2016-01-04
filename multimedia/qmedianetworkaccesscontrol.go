package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QMediaNetworkAccessControl struct {
	QMediaControl
}

type QMediaNetworkAccessControl_ITF interface {
	QMediaControl_ITF
	QMediaNetworkAccessControl_PTR() *QMediaNetworkAccessControl
}

func PointerFromQMediaNetworkAccessControl(ptr QMediaNetworkAccessControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaNetworkAccessControl_PTR().Pointer()
	}
	return nil
}

func NewQMediaNetworkAccessControlFromPointer(ptr unsafe.Pointer) *QMediaNetworkAccessControl {
	var n = new(QMediaNetworkAccessControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMediaNetworkAccessControl_") {
		n.SetObjectName("QMediaNetworkAccessControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QMediaNetworkAccessControl) QMediaNetworkAccessControl_PTR() *QMediaNetworkAccessControl {
	return ptr
}

func (ptr *QMediaNetworkAccessControl) DestroyQMediaNetworkAccessControl() {
	defer qt.Recovering("QMediaNetworkAccessControl::~QMediaNetworkAccessControl")

	if ptr.Pointer() != nil {
		C.QMediaNetworkAccessControl_DestroyQMediaNetworkAccessControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMediaNetworkAccessControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMediaNetworkAccessControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMediaNetworkAccessControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMediaNetworkAccessControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMediaNetworkAccessControlTimerEvent
func callbackQMediaNetworkAccessControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaNetworkAccessControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMediaNetworkAccessControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMediaNetworkAccessControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMediaNetworkAccessControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QMediaNetworkAccessControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMediaNetworkAccessControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMediaNetworkAccessControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QMediaNetworkAccessControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMediaNetworkAccessControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMediaNetworkAccessControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMediaNetworkAccessControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMediaNetworkAccessControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMediaNetworkAccessControlChildEvent
func callbackQMediaNetworkAccessControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaNetworkAccessControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMediaNetworkAccessControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMediaNetworkAccessControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMediaNetworkAccessControl::childEvent")

	if ptr.Pointer() != nil {
		C.QMediaNetworkAccessControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMediaNetworkAccessControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMediaNetworkAccessControl::childEvent")

	if ptr.Pointer() != nil {
		C.QMediaNetworkAccessControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMediaNetworkAccessControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMediaNetworkAccessControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMediaNetworkAccessControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMediaNetworkAccessControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMediaNetworkAccessControlCustomEvent
func callbackQMediaNetworkAccessControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaNetworkAccessControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMediaNetworkAccessControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMediaNetworkAccessControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMediaNetworkAccessControl::customEvent")

	if ptr.Pointer() != nil {
		C.QMediaNetworkAccessControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMediaNetworkAccessControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMediaNetworkAccessControl::customEvent")

	if ptr.Pointer() != nil {
		C.QMediaNetworkAccessControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
