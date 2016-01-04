package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QSequentialAnimationGroup struct {
	QAnimationGroup
}

type QSequentialAnimationGroup_ITF interface {
	QAnimationGroup_ITF
	QSequentialAnimationGroup_PTR() *QSequentialAnimationGroup
}

func PointerFromQSequentialAnimationGroup(ptr QSequentialAnimationGroup_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSequentialAnimationGroup_PTR().Pointer()
	}
	return nil
}

func NewQSequentialAnimationGroupFromPointer(ptr unsafe.Pointer) *QSequentialAnimationGroup {
	var n = new(QSequentialAnimationGroup)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSequentialAnimationGroup_") {
		n.SetObjectName("QSequentialAnimationGroup_" + qt.Identifier())
	}
	return n
}

func (ptr *QSequentialAnimationGroup) QSequentialAnimationGroup_PTR() *QSequentialAnimationGroup {
	return ptr
}

func (ptr *QSequentialAnimationGroup) CurrentAnimation() *QAbstractAnimation {
	defer qt.Recovering("QSequentialAnimationGroup::currentAnimation")

	if ptr.Pointer() != nil {
		return NewQAbstractAnimationFromPointer(C.QSequentialAnimationGroup_CurrentAnimation(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSequentialAnimationGroup) AddPause(msecs int) *QPauseAnimation {
	defer qt.Recovering("QSequentialAnimationGroup::addPause")

	if ptr.Pointer() != nil {
		return NewQPauseAnimationFromPointer(C.QSequentialAnimationGroup_AddPause(ptr.Pointer(), C.int(msecs)))
	}
	return nil
}

func (ptr *QSequentialAnimationGroup) ConnectCurrentAnimationChanged(f func(current *QAbstractAnimation)) {
	defer qt.Recovering("connect QSequentialAnimationGroup::currentAnimationChanged")

	if ptr.Pointer() != nil {
		C.QSequentialAnimationGroup_ConnectCurrentAnimationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentAnimationChanged", f)
	}
}

func (ptr *QSequentialAnimationGroup) DisconnectCurrentAnimationChanged() {
	defer qt.Recovering("disconnect QSequentialAnimationGroup::currentAnimationChanged")

	if ptr.Pointer() != nil {
		C.QSequentialAnimationGroup_DisconnectCurrentAnimationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentAnimationChanged")
	}
}

//export callbackQSequentialAnimationGroupCurrentAnimationChanged
func callbackQSequentialAnimationGroupCurrentAnimationChanged(ptr unsafe.Pointer, ptrName *C.char, current unsafe.Pointer) {
	defer qt.Recovering("callback QSequentialAnimationGroup::currentAnimationChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "currentAnimationChanged"); signal != nil {
		signal.(func(*QAbstractAnimation))(NewQAbstractAnimationFromPointer(current))
	}

}

func (ptr *QSequentialAnimationGroup) CurrentAnimationChanged(current QAbstractAnimation_ITF) {
	defer qt.Recovering("QSequentialAnimationGroup::currentAnimationChanged")

	if ptr.Pointer() != nil {
		C.QSequentialAnimationGroup_CurrentAnimationChanged(ptr.Pointer(), PointerFromQAbstractAnimation(current))
	}
}

func (ptr *QSequentialAnimationGroup) Duration() int {
	defer qt.Recovering("QSequentialAnimationGroup::duration")

	if ptr.Pointer() != nil {
		return int(C.QSequentialAnimationGroup_Duration(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSequentialAnimationGroup) Event(event QEvent_ITF) bool {
	defer qt.Recovering("QSequentialAnimationGroup::event")

	if ptr.Pointer() != nil {
		return C.QSequentialAnimationGroup_Event(ptr.Pointer(), PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QSequentialAnimationGroup) InsertPause(index int, msecs int) *QPauseAnimation {
	defer qt.Recovering("QSequentialAnimationGroup::insertPause")

	if ptr.Pointer() != nil {
		return NewQPauseAnimationFromPointer(C.QSequentialAnimationGroup_InsertPause(ptr.Pointer(), C.int(index), C.int(msecs)))
	}
	return nil
}

func (ptr *QSequentialAnimationGroup) ConnectUpdateCurrentTime(f func(currentTime int)) {
	defer qt.Recovering("connect QSequentialAnimationGroup::updateCurrentTime")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updateCurrentTime", f)
	}
}

func (ptr *QSequentialAnimationGroup) DisconnectUpdateCurrentTime() {
	defer qt.Recovering("disconnect QSequentialAnimationGroup::updateCurrentTime")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updateCurrentTime")
	}
}

//export callbackQSequentialAnimationGroupUpdateCurrentTime
func callbackQSequentialAnimationGroupUpdateCurrentTime(ptr unsafe.Pointer, ptrName *C.char, currentTime C.int) {
	defer qt.Recovering("callback QSequentialAnimationGroup::updateCurrentTime")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateCurrentTime"); signal != nil {
		signal.(func(int))(int(currentTime))
	} else {
		NewQSequentialAnimationGroupFromPointer(ptr).UpdateCurrentTimeDefault(int(currentTime))
	}
}

func (ptr *QSequentialAnimationGroup) UpdateCurrentTime(currentTime int) {
	defer qt.Recovering("QSequentialAnimationGroup::updateCurrentTime")

	if ptr.Pointer() != nil {
		C.QSequentialAnimationGroup_UpdateCurrentTime(ptr.Pointer(), C.int(currentTime))
	}
}

func (ptr *QSequentialAnimationGroup) UpdateCurrentTimeDefault(currentTime int) {
	defer qt.Recovering("QSequentialAnimationGroup::updateCurrentTime")

	if ptr.Pointer() != nil {
		C.QSequentialAnimationGroup_UpdateCurrentTimeDefault(ptr.Pointer(), C.int(currentTime))
	}
}

func (ptr *QSequentialAnimationGroup) ConnectUpdateDirection(f func(direction QAbstractAnimation__Direction)) {
	defer qt.Recovering("connect QSequentialAnimationGroup::updateDirection")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updateDirection", f)
	}
}

func (ptr *QSequentialAnimationGroup) DisconnectUpdateDirection() {
	defer qt.Recovering("disconnect QSequentialAnimationGroup::updateDirection")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updateDirection")
	}
}

//export callbackQSequentialAnimationGroupUpdateDirection
func callbackQSequentialAnimationGroupUpdateDirection(ptr unsafe.Pointer, ptrName *C.char, direction C.int) {
	defer qt.Recovering("callback QSequentialAnimationGroup::updateDirection")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateDirection"); signal != nil {
		signal.(func(QAbstractAnimation__Direction))(QAbstractAnimation__Direction(direction))
	} else {
		NewQSequentialAnimationGroupFromPointer(ptr).UpdateDirectionDefault(QAbstractAnimation__Direction(direction))
	}
}

func (ptr *QSequentialAnimationGroup) UpdateDirection(direction QAbstractAnimation__Direction) {
	defer qt.Recovering("QSequentialAnimationGroup::updateDirection")

	if ptr.Pointer() != nil {
		C.QSequentialAnimationGroup_UpdateDirection(ptr.Pointer(), C.int(direction))
	}
}

func (ptr *QSequentialAnimationGroup) UpdateDirectionDefault(direction QAbstractAnimation__Direction) {
	defer qt.Recovering("QSequentialAnimationGroup::updateDirection")

	if ptr.Pointer() != nil {
		C.QSequentialAnimationGroup_UpdateDirectionDefault(ptr.Pointer(), C.int(direction))
	}
}

func (ptr *QSequentialAnimationGroup) ConnectUpdateState(f func(newState QAbstractAnimation__State, oldState QAbstractAnimation__State)) {
	defer qt.Recovering("connect QSequentialAnimationGroup::updateState")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updateState", f)
	}
}

func (ptr *QSequentialAnimationGroup) DisconnectUpdateState() {
	defer qt.Recovering("disconnect QSequentialAnimationGroup::updateState")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updateState")
	}
}

//export callbackQSequentialAnimationGroupUpdateState
func callbackQSequentialAnimationGroupUpdateState(ptr unsafe.Pointer, ptrName *C.char, newState C.int, oldState C.int) {
	defer qt.Recovering("callback QSequentialAnimationGroup::updateState")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateState"); signal != nil {
		signal.(func(QAbstractAnimation__State, QAbstractAnimation__State))(QAbstractAnimation__State(newState), QAbstractAnimation__State(oldState))
	} else {
		NewQSequentialAnimationGroupFromPointer(ptr).UpdateStateDefault(QAbstractAnimation__State(newState), QAbstractAnimation__State(oldState))
	}
}

func (ptr *QSequentialAnimationGroup) UpdateState(newState QAbstractAnimation__State, oldState QAbstractAnimation__State) {
	defer qt.Recovering("QSequentialAnimationGroup::updateState")

	if ptr.Pointer() != nil {
		C.QSequentialAnimationGroup_UpdateState(ptr.Pointer(), C.int(newState), C.int(oldState))
	}
}

func (ptr *QSequentialAnimationGroup) UpdateStateDefault(newState QAbstractAnimation__State, oldState QAbstractAnimation__State) {
	defer qt.Recovering("QSequentialAnimationGroup::updateState")

	if ptr.Pointer() != nil {
		C.QSequentialAnimationGroup_UpdateStateDefault(ptr.Pointer(), C.int(newState), C.int(oldState))
	}
}

func (ptr *QSequentialAnimationGroup) DestroyQSequentialAnimationGroup() {
	defer qt.Recovering("QSequentialAnimationGroup::~QSequentialAnimationGroup")

	if ptr.Pointer() != nil {
		C.QSequentialAnimationGroup_DestroyQSequentialAnimationGroup(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSequentialAnimationGroup) ConnectTimerEvent(f func(event *QTimerEvent)) {
	defer qt.Recovering("connect QSequentialAnimationGroup::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSequentialAnimationGroup) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSequentialAnimationGroup::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSequentialAnimationGroupTimerEvent
func callbackQSequentialAnimationGroupTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSequentialAnimationGroup::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*QTimerEvent))(NewQTimerEventFromPointer(event))
	} else {
		NewQSequentialAnimationGroupFromPointer(ptr).TimerEventDefault(NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSequentialAnimationGroup) TimerEvent(event QTimerEvent_ITF) {
	defer qt.Recovering("QSequentialAnimationGroup::timerEvent")

	if ptr.Pointer() != nil {
		C.QSequentialAnimationGroup_TimerEvent(ptr.Pointer(), PointerFromQTimerEvent(event))
	}
}

func (ptr *QSequentialAnimationGroup) TimerEventDefault(event QTimerEvent_ITF) {
	defer qt.Recovering("QSequentialAnimationGroup::timerEvent")

	if ptr.Pointer() != nil {
		C.QSequentialAnimationGroup_TimerEventDefault(ptr.Pointer(), PointerFromQTimerEvent(event))
	}
}

func (ptr *QSequentialAnimationGroup) ConnectChildEvent(f func(event *QChildEvent)) {
	defer qt.Recovering("connect QSequentialAnimationGroup::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSequentialAnimationGroup) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSequentialAnimationGroup::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSequentialAnimationGroupChildEvent
func callbackQSequentialAnimationGroupChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSequentialAnimationGroup::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*QChildEvent))(NewQChildEventFromPointer(event))
	} else {
		NewQSequentialAnimationGroupFromPointer(ptr).ChildEventDefault(NewQChildEventFromPointer(event))
	}
}

func (ptr *QSequentialAnimationGroup) ChildEvent(event QChildEvent_ITF) {
	defer qt.Recovering("QSequentialAnimationGroup::childEvent")

	if ptr.Pointer() != nil {
		C.QSequentialAnimationGroup_ChildEvent(ptr.Pointer(), PointerFromQChildEvent(event))
	}
}

func (ptr *QSequentialAnimationGroup) ChildEventDefault(event QChildEvent_ITF) {
	defer qt.Recovering("QSequentialAnimationGroup::childEvent")

	if ptr.Pointer() != nil {
		C.QSequentialAnimationGroup_ChildEventDefault(ptr.Pointer(), PointerFromQChildEvent(event))
	}
}

func (ptr *QSequentialAnimationGroup) ConnectCustomEvent(f func(event *QEvent)) {
	defer qt.Recovering("connect QSequentialAnimationGroup::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSequentialAnimationGroup) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSequentialAnimationGroup::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSequentialAnimationGroupCustomEvent
func callbackQSequentialAnimationGroupCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSequentialAnimationGroup::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
	} else {
		NewQSequentialAnimationGroupFromPointer(ptr).CustomEventDefault(NewQEventFromPointer(event))
	}
}

func (ptr *QSequentialAnimationGroup) CustomEvent(event QEvent_ITF) {
	defer qt.Recovering("QSequentialAnimationGroup::customEvent")

	if ptr.Pointer() != nil {
		C.QSequentialAnimationGroup_CustomEvent(ptr.Pointer(), PointerFromQEvent(event))
	}
}

func (ptr *QSequentialAnimationGroup) CustomEventDefault(event QEvent_ITF) {
	defer qt.Recovering("QSequentialAnimationGroup::customEvent")

	if ptr.Pointer() != nil {
		C.QSequentialAnimationGroup_CustomEventDefault(ptr.Pointer(), PointerFromQEvent(event))
	}
}
