package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QGesture struct {
	core.QObject
}

type QGesture_ITF interface {
	core.QObject_ITF
	QGesture_PTR() *QGesture
}

func PointerFromQGesture(ptr QGesture_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGesture_PTR().Pointer()
	}
	return nil
}

func NewQGestureFromPointer(ptr unsafe.Pointer) *QGesture {
	var n = new(QGesture)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QGesture_") {
		n.SetObjectName("QGesture_" + qt.Identifier())
	}
	return n
}

func (ptr *QGesture) QGesture_PTR() *QGesture {
	return ptr
}

//QGesture::GestureCancelPolicy
type QGesture__GestureCancelPolicy int64

const (
	QGesture__CancelNone         = QGesture__GestureCancelPolicy(0)
	QGesture__CancelAllInContext = QGesture__GestureCancelPolicy(1)
)

func (ptr *QGesture) GestureCancelPolicy() QGesture__GestureCancelPolicy {
	defer qt.Recovering("QGesture::gestureCancelPolicy")

	if ptr.Pointer() != nil {
		return QGesture__GestureCancelPolicy(C.QGesture_GestureCancelPolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGesture) GestureType() core.Qt__GestureType {
	defer qt.Recovering("QGesture::gestureType")

	if ptr.Pointer() != nil {
		return core.Qt__GestureType(C.QGesture_GestureType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGesture) HasHotSpot() bool {
	defer qt.Recovering("QGesture::hasHotSpot")

	if ptr.Pointer() != nil {
		return C.QGesture_HasHotSpot(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGesture) SetGestureCancelPolicy(policy QGesture__GestureCancelPolicy) {
	defer qt.Recovering("QGesture::setGestureCancelPolicy")

	if ptr.Pointer() != nil {
		C.QGesture_SetGestureCancelPolicy(ptr.Pointer(), C.int(policy))
	}
}

func (ptr *QGesture) SetHotSpot(value core.QPointF_ITF) {
	defer qt.Recovering("QGesture::setHotSpot")

	if ptr.Pointer() != nil {
		C.QGesture_SetHotSpot(ptr.Pointer(), core.PointerFromQPointF(value))
	}
}

func (ptr *QGesture) State() core.Qt__GestureState {
	defer qt.Recovering("QGesture::state")

	if ptr.Pointer() != nil {
		return core.Qt__GestureState(C.QGesture_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGesture) UnsetHotSpot() {
	defer qt.Recovering("QGesture::unsetHotSpot")

	if ptr.Pointer() != nil {
		C.QGesture_UnsetHotSpot(ptr.Pointer())
	}
}

func NewQGesture(parent core.QObject_ITF) *QGesture {
	defer qt.Recovering("QGesture::QGesture")

	return NewQGestureFromPointer(C.QGesture_NewQGesture(core.PointerFromQObject(parent)))
}

func (ptr *QGesture) DestroyQGesture() {
	defer qt.Recovering("QGesture::~QGesture")

	if ptr.Pointer() != nil {
		C.QGesture_DestroyQGesture(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGesture) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QGesture::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QGesture) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QGesture::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQGestureTimerEvent
func callbackQGestureTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGesture::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGestureFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGesture) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGesture::timerEvent")

	if ptr.Pointer() != nil {
		C.QGesture_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGesture) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGesture::timerEvent")

	if ptr.Pointer() != nil {
		C.QGesture_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGesture) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QGesture::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QGesture) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QGesture::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQGestureChildEvent
func callbackQGestureChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGesture::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQGestureFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QGesture) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QGesture::childEvent")

	if ptr.Pointer() != nil {
		C.QGesture_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGesture) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QGesture::childEvent")

	if ptr.Pointer() != nil {
		C.QGesture_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGesture) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGesture::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QGesture) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QGesture::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQGestureCustomEvent
func callbackQGestureCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGesture::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGestureFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGesture) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QGesture::customEvent")

	if ptr.Pointer() != nil {
		C.QGesture_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGesture) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QGesture::customEvent")

	if ptr.Pointer() != nil {
		C.QGesture_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
