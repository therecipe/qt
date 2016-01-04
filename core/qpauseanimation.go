package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QPauseAnimation struct {
	QAbstractAnimation
}

type QPauseAnimation_ITF interface {
	QAbstractAnimation_ITF
	QPauseAnimation_PTR() *QPauseAnimation
}

func PointerFromQPauseAnimation(ptr QPauseAnimation_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPauseAnimation_PTR().Pointer()
	}
	return nil
}

func NewQPauseAnimationFromPointer(ptr unsafe.Pointer) *QPauseAnimation {
	var n = new(QPauseAnimation)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QPauseAnimation_") {
		n.SetObjectName("QPauseAnimation_" + qt.Identifier())
	}
	return n
}

func (ptr *QPauseAnimation) QPauseAnimation_PTR() *QPauseAnimation {
	return ptr
}

func (ptr *QPauseAnimation) Duration() int {
	defer qt.Recovering("QPauseAnimation::duration")

	if ptr.Pointer() != nil {
		return int(C.QPauseAnimation_Duration(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPauseAnimation) SetDuration(msecs int) {
	defer qt.Recovering("QPauseAnimation::setDuration")

	if ptr.Pointer() != nil {
		C.QPauseAnimation_SetDuration(ptr.Pointer(), C.int(msecs))
	}
}

func NewQPauseAnimation(parent QObject_ITF) *QPauseAnimation {
	defer qt.Recovering("QPauseAnimation::QPauseAnimation")

	return NewQPauseAnimationFromPointer(C.QPauseAnimation_NewQPauseAnimation(PointerFromQObject(parent)))
}

func NewQPauseAnimation2(msecs int, parent QObject_ITF) *QPauseAnimation {
	defer qt.Recovering("QPauseAnimation::QPauseAnimation")

	return NewQPauseAnimationFromPointer(C.QPauseAnimation_NewQPauseAnimation2(C.int(msecs), PointerFromQObject(parent)))
}

func (ptr *QPauseAnimation) Event(e QEvent_ITF) bool {
	defer qt.Recovering("QPauseAnimation::event")

	if ptr.Pointer() != nil {
		return C.QPauseAnimation_Event(ptr.Pointer(), PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QPauseAnimation) ConnectUpdateCurrentTime(f func(v int)) {
	defer qt.Recovering("connect QPauseAnimation::updateCurrentTime")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updateCurrentTime", f)
	}
}

func (ptr *QPauseAnimation) DisconnectUpdateCurrentTime() {
	defer qt.Recovering("disconnect QPauseAnimation::updateCurrentTime")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updateCurrentTime")
	}
}

//export callbackQPauseAnimationUpdateCurrentTime
func callbackQPauseAnimationUpdateCurrentTime(ptr unsafe.Pointer, ptrName *C.char, v C.int) {
	defer qt.Recovering("callback QPauseAnimation::updateCurrentTime")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateCurrentTime"); signal != nil {
		signal.(func(int))(int(v))
	}

}

func (ptr *QPauseAnimation) UpdateCurrentTime(v int) {
	defer qt.Recovering("QPauseAnimation::updateCurrentTime")

	if ptr.Pointer() != nil {
		C.QPauseAnimation_UpdateCurrentTime(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QPauseAnimation) UpdateCurrentTimeDefault(v int) {
	defer qt.Recovering("QPauseAnimation::updateCurrentTime")

	if ptr.Pointer() != nil {
		C.QPauseAnimation_UpdateCurrentTimeDefault(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QPauseAnimation) DestroyQPauseAnimation() {
	defer qt.Recovering("QPauseAnimation::~QPauseAnimation")

	if ptr.Pointer() != nil {
		C.QPauseAnimation_DestroyQPauseAnimation(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QPauseAnimation) ConnectUpdateDirection(f func(direction QAbstractAnimation__Direction)) {
	defer qt.Recovering("connect QPauseAnimation::updateDirection")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updateDirection", f)
	}
}

func (ptr *QPauseAnimation) DisconnectUpdateDirection() {
	defer qt.Recovering("disconnect QPauseAnimation::updateDirection")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updateDirection")
	}
}

//export callbackQPauseAnimationUpdateDirection
func callbackQPauseAnimationUpdateDirection(ptr unsafe.Pointer, ptrName *C.char, direction C.int) {
	defer qt.Recovering("callback QPauseAnimation::updateDirection")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateDirection"); signal != nil {
		signal.(func(QAbstractAnimation__Direction))(QAbstractAnimation__Direction(direction))
	} else {
		NewQPauseAnimationFromPointer(ptr).UpdateDirectionDefault(QAbstractAnimation__Direction(direction))
	}
}

func (ptr *QPauseAnimation) UpdateDirection(direction QAbstractAnimation__Direction) {
	defer qt.Recovering("QPauseAnimation::updateDirection")

	if ptr.Pointer() != nil {
		C.QPauseAnimation_UpdateDirection(ptr.Pointer(), C.int(direction))
	}
}

func (ptr *QPauseAnimation) UpdateDirectionDefault(direction QAbstractAnimation__Direction) {
	defer qt.Recovering("QPauseAnimation::updateDirection")

	if ptr.Pointer() != nil {
		C.QPauseAnimation_UpdateDirectionDefault(ptr.Pointer(), C.int(direction))
	}
}

func (ptr *QPauseAnimation) ConnectUpdateState(f func(newState QAbstractAnimation__State, oldState QAbstractAnimation__State)) {
	defer qt.Recovering("connect QPauseAnimation::updateState")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updateState", f)
	}
}

func (ptr *QPauseAnimation) DisconnectUpdateState() {
	defer qt.Recovering("disconnect QPauseAnimation::updateState")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updateState")
	}
}

//export callbackQPauseAnimationUpdateState
func callbackQPauseAnimationUpdateState(ptr unsafe.Pointer, ptrName *C.char, newState C.int, oldState C.int) {
	defer qt.Recovering("callback QPauseAnimation::updateState")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateState"); signal != nil {
		signal.(func(QAbstractAnimation__State, QAbstractAnimation__State))(QAbstractAnimation__State(newState), QAbstractAnimation__State(oldState))
	} else {
		NewQPauseAnimationFromPointer(ptr).UpdateStateDefault(QAbstractAnimation__State(newState), QAbstractAnimation__State(oldState))
	}
}

func (ptr *QPauseAnimation) UpdateState(newState QAbstractAnimation__State, oldState QAbstractAnimation__State) {
	defer qt.Recovering("QPauseAnimation::updateState")

	if ptr.Pointer() != nil {
		C.QPauseAnimation_UpdateState(ptr.Pointer(), C.int(newState), C.int(oldState))
	}
}

func (ptr *QPauseAnimation) UpdateStateDefault(newState QAbstractAnimation__State, oldState QAbstractAnimation__State) {
	defer qt.Recovering("QPauseAnimation::updateState")

	if ptr.Pointer() != nil {
		C.QPauseAnimation_UpdateStateDefault(ptr.Pointer(), C.int(newState), C.int(oldState))
	}
}

func (ptr *QPauseAnimation) ConnectTimerEvent(f func(event *QTimerEvent)) {
	defer qt.Recovering("connect QPauseAnimation::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QPauseAnimation) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QPauseAnimation::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQPauseAnimationTimerEvent
func callbackQPauseAnimationTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QPauseAnimation::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*QTimerEvent))(NewQTimerEventFromPointer(event))
	} else {
		NewQPauseAnimationFromPointer(ptr).TimerEventDefault(NewQTimerEventFromPointer(event))
	}
}

func (ptr *QPauseAnimation) TimerEvent(event QTimerEvent_ITF) {
	defer qt.Recovering("QPauseAnimation::timerEvent")

	if ptr.Pointer() != nil {
		C.QPauseAnimation_TimerEvent(ptr.Pointer(), PointerFromQTimerEvent(event))
	}
}

func (ptr *QPauseAnimation) TimerEventDefault(event QTimerEvent_ITF) {
	defer qt.Recovering("QPauseAnimation::timerEvent")

	if ptr.Pointer() != nil {
		C.QPauseAnimation_TimerEventDefault(ptr.Pointer(), PointerFromQTimerEvent(event))
	}
}

func (ptr *QPauseAnimation) ConnectChildEvent(f func(event *QChildEvent)) {
	defer qt.Recovering("connect QPauseAnimation::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QPauseAnimation) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QPauseAnimation::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQPauseAnimationChildEvent
func callbackQPauseAnimationChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QPauseAnimation::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*QChildEvent))(NewQChildEventFromPointer(event))
	} else {
		NewQPauseAnimationFromPointer(ptr).ChildEventDefault(NewQChildEventFromPointer(event))
	}
}

func (ptr *QPauseAnimation) ChildEvent(event QChildEvent_ITF) {
	defer qt.Recovering("QPauseAnimation::childEvent")

	if ptr.Pointer() != nil {
		C.QPauseAnimation_ChildEvent(ptr.Pointer(), PointerFromQChildEvent(event))
	}
}

func (ptr *QPauseAnimation) ChildEventDefault(event QChildEvent_ITF) {
	defer qt.Recovering("QPauseAnimation::childEvent")

	if ptr.Pointer() != nil {
		C.QPauseAnimation_ChildEventDefault(ptr.Pointer(), PointerFromQChildEvent(event))
	}
}

func (ptr *QPauseAnimation) ConnectCustomEvent(f func(event *QEvent)) {
	defer qt.Recovering("connect QPauseAnimation::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QPauseAnimation) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QPauseAnimation::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQPauseAnimationCustomEvent
func callbackQPauseAnimationCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QPauseAnimation::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
	} else {
		NewQPauseAnimationFromPointer(ptr).CustomEventDefault(NewQEventFromPointer(event))
	}
}

func (ptr *QPauseAnimation) CustomEvent(event QEvent_ITF) {
	defer qt.Recovering("QPauseAnimation::customEvent")

	if ptr.Pointer() != nil {
		C.QPauseAnimation_CustomEvent(ptr.Pointer(), PointerFromQEvent(event))
	}
}

func (ptr *QPauseAnimation) CustomEventDefault(event QEvent_ITF) {
	defer qt.Recovering("QPauseAnimation::customEvent")

	if ptr.Pointer() != nil {
		C.QPauseAnimation_CustomEventDefault(ptr.Pointer(), PointerFromQEvent(event))
	}
}
