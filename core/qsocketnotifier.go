package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QSocketNotifier struct {
	QObject
}

type QSocketNotifier_ITF interface {
	QObject_ITF
	QSocketNotifier_PTR() *QSocketNotifier
}

func PointerFromQSocketNotifier(ptr QSocketNotifier_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSocketNotifier_PTR().Pointer()
	}
	return nil
}

func NewQSocketNotifierFromPointer(ptr unsafe.Pointer) *QSocketNotifier {
	var n = new(QSocketNotifier)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSocketNotifier_") {
		n.SetObjectName("QSocketNotifier_" + qt.Identifier())
	}
	return n
}

func (ptr *QSocketNotifier) QSocketNotifier_PTR() *QSocketNotifier {
	return ptr
}

//QSocketNotifier::Type
type QSocketNotifier__Type int64

const (
	QSocketNotifier__Read      = QSocketNotifier__Type(0)
	QSocketNotifier__Write     = QSocketNotifier__Type(1)
	QSocketNotifier__Exception = QSocketNotifier__Type(2)
)

func (ptr *QSocketNotifier) ConnectActivated(f func(socket int)) {
	defer qt.Recovering("connect QSocketNotifier::activated")

	if ptr.Pointer() != nil {
		C.QSocketNotifier_ConnectActivated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "activated", f)
	}
}

func (ptr *QSocketNotifier) DisconnectActivated() {
	defer qt.Recovering("disconnect QSocketNotifier::activated")

	if ptr.Pointer() != nil {
		C.QSocketNotifier_DisconnectActivated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "activated")
	}
}

//export callbackQSocketNotifierActivated
func callbackQSocketNotifierActivated(ptr unsafe.Pointer, ptrName *C.char, socket C.int) {
	defer qt.Recovering("callback QSocketNotifier::activated")

	if signal := qt.GetSignal(C.GoString(ptrName), "activated"); signal != nil {
		signal.(func(int))(int(socket))
	}

}

func (ptr *QSocketNotifier) Event(e QEvent_ITF) bool {
	defer qt.Recovering("QSocketNotifier::event")

	if ptr.Pointer() != nil {
		return C.QSocketNotifier_Event(ptr.Pointer(), PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QSocketNotifier) IsEnabled() bool {
	defer qt.Recovering("QSocketNotifier::isEnabled")

	if ptr.Pointer() != nil {
		return C.QSocketNotifier_IsEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSocketNotifier) SetEnabled(enable bool) {
	defer qt.Recovering("QSocketNotifier::setEnabled")

	if ptr.Pointer() != nil {
		C.QSocketNotifier_SetEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QSocketNotifier) Type() QSocketNotifier__Type {
	defer qt.Recovering("QSocketNotifier::type")

	if ptr.Pointer() != nil {
		return QSocketNotifier__Type(C.QSocketNotifier_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSocketNotifier) DestroyQSocketNotifier() {
	defer qt.Recovering("QSocketNotifier::~QSocketNotifier")

	if ptr.Pointer() != nil {
		C.QSocketNotifier_DestroyQSocketNotifier(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSocketNotifier) ConnectTimerEvent(f func(event *QTimerEvent)) {
	defer qt.Recovering("connect QSocketNotifier::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSocketNotifier) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSocketNotifier::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSocketNotifierTimerEvent
func callbackQSocketNotifierTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSocketNotifier::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*QTimerEvent))(NewQTimerEventFromPointer(event))
	} else {
		NewQSocketNotifierFromPointer(ptr).TimerEventDefault(NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSocketNotifier) TimerEvent(event QTimerEvent_ITF) {
	defer qt.Recovering("QSocketNotifier::timerEvent")

	if ptr.Pointer() != nil {
		C.QSocketNotifier_TimerEvent(ptr.Pointer(), PointerFromQTimerEvent(event))
	}
}

func (ptr *QSocketNotifier) TimerEventDefault(event QTimerEvent_ITF) {
	defer qt.Recovering("QSocketNotifier::timerEvent")

	if ptr.Pointer() != nil {
		C.QSocketNotifier_TimerEventDefault(ptr.Pointer(), PointerFromQTimerEvent(event))
	}
}

func (ptr *QSocketNotifier) ConnectChildEvent(f func(event *QChildEvent)) {
	defer qt.Recovering("connect QSocketNotifier::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSocketNotifier) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSocketNotifier::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSocketNotifierChildEvent
func callbackQSocketNotifierChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSocketNotifier::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*QChildEvent))(NewQChildEventFromPointer(event))
	} else {
		NewQSocketNotifierFromPointer(ptr).ChildEventDefault(NewQChildEventFromPointer(event))
	}
}

func (ptr *QSocketNotifier) ChildEvent(event QChildEvent_ITF) {
	defer qt.Recovering("QSocketNotifier::childEvent")

	if ptr.Pointer() != nil {
		C.QSocketNotifier_ChildEvent(ptr.Pointer(), PointerFromQChildEvent(event))
	}
}

func (ptr *QSocketNotifier) ChildEventDefault(event QChildEvent_ITF) {
	defer qt.Recovering("QSocketNotifier::childEvent")

	if ptr.Pointer() != nil {
		C.QSocketNotifier_ChildEventDefault(ptr.Pointer(), PointerFromQChildEvent(event))
	}
}

func (ptr *QSocketNotifier) ConnectCustomEvent(f func(event *QEvent)) {
	defer qt.Recovering("connect QSocketNotifier::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSocketNotifier) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSocketNotifier::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSocketNotifierCustomEvent
func callbackQSocketNotifierCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSocketNotifier::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
	} else {
		NewQSocketNotifierFromPointer(ptr).CustomEventDefault(NewQEventFromPointer(event))
	}
}

func (ptr *QSocketNotifier) CustomEvent(event QEvent_ITF) {
	defer qt.Recovering("QSocketNotifier::customEvent")

	if ptr.Pointer() != nil {
		C.QSocketNotifier_CustomEvent(ptr.Pointer(), PointerFromQEvent(event))
	}
}

func (ptr *QSocketNotifier) CustomEventDefault(event QEvent_ITF) {
	defer qt.Recovering("QSocketNotifier::customEvent")

	if ptr.Pointer() != nil {
		C.QSocketNotifier_CustomEventDefault(ptr.Pointer(), PointerFromQEvent(event))
	}
}
