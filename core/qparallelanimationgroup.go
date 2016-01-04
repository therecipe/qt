package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QParallelAnimationGroup struct {
	QAnimationGroup
}

type QParallelAnimationGroup_ITF interface {
	QAnimationGroup_ITF
	QParallelAnimationGroup_PTR() *QParallelAnimationGroup
}

func PointerFromQParallelAnimationGroup(ptr QParallelAnimationGroup_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QParallelAnimationGroup_PTR().Pointer()
	}
	return nil
}

func NewQParallelAnimationGroupFromPointer(ptr unsafe.Pointer) *QParallelAnimationGroup {
	var n = new(QParallelAnimationGroup)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QParallelAnimationGroup_") {
		n.SetObjectName("QParallelAnimationGroup_" + qt.Identifier())
	}
	return n
}

func (ptr *QParallelAnimationGroup) QParallelAnimationGroup_PTR() *QParallelAnimationGroup {
	return ptr
}

func (ptr *QParallelAnimationGroup) Duration() int {
	defer qt.Recovering("QParallelAnimationGroup::duration")

	if ptr.Pointer() != nil {
		return int(C.QParallelAnimationGroup_Duration(ptr.Pointer()))
	}
	return 0
}

func (ptr *QParallelAnimationGroup) Event(event QEvent_ITF) bool {
	defer qt.Recovering("QParallelAnimationGroup::event")

	if ptr.Pointer() != nil {
		return C.QParallelAnimationGroup_Event(ptr.Pointer(), PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QParallelAnimationGroup) ConnectUpdateCurrentTime(f func(currentTime int)) {
	defer qt.Recovering("connect QParallelAnimationGroup::updateCurrentTime")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updateCurrentTime", f)
	}
}

func (ptr *QParallelAnimationGroup) DisconnectUpdateCurrentTime() {
	defer qt.Recovering("disconnect QParallelAnimationGroup::updateCurrentTime")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updateCurrentTime")
	}
}

//export callbackQParallelAnimationGroupUpdateCurrentTime
func callbackQParallelAnimationGroupUpdateCurrentTime(ptr unsafe.Pointer, ptrName *C.char, currentTime C.int) {
	defer qt.Recovering("callback QParallelAnimationGroup::updateCurrentTime")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateCurrentTime"); signal != nil {
		signal.(func(int))(int(currentTime))
	}

}

func (ptr *QParallelAnimationGroup) UpdateCurrentTime(currentTime int) {
	defer qt.Recovering("QParallelAnimationGroup::updateCurrentTime")

	if ptr.Pointer() != nil {
		C.QParallelAnimationGroup_UpdateCurrentTime(ptr.Pointer(), C.int(currentTime))
	}
}

func (ptr *QParallelAnimationGroup) UpdateCurrentTimeDefault(currentTime int) {
	defer qt.Recovering("QParallelAnimationGroup::updateCurrentTime")

	if ptr.Pointer() != nil {
		C.QParallelAnimationGroup_UpdateCurrentTimeDefault(ptr.Pointer(), C.int(currentTime))
	}
}

func (ptr *QParallelAnimationGroup) ConnectUpdateDirection(f func(direction QAbstractAnimation__Direction)) {
	defer qt.Recovering("connect QParallelAnimationGroup::updateDirection")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updateDirection", f)
	}
}

func (ptr *QParallelAnimationGroup) DisconnectUpdateDirection() {
	defer qt.Recovering("disconnect QParallelAnimationGroup::updateDirection")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updateDirection")
	}
}

//export callbackQParallelAnimationGroupUpdateDirection
func callbackQParallelAnimationGroupUpdateDirection(ptr unsafe.Pointer, ptrName *C.char, direction C.int) {
	defer qt.Recovering("callback QParallelAnimationGroup::updateDirection")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateDirection"); signal != nil {
		signal.(func(QAbstractAnimation__Direction))(QAbstractAnimation__Direction(direction))
	} else {
		NewQParallelAnimationGroupFromPointer(ptr).UpdateDirectionDefault(QAbstractAnimation__Direction(direction))
	}
}

func (ptr *QParallelAnimationGroup) UpdateDirection(direction QAbstractAnimation__Direction) {
	defer qt.Recovering("QParallelAnimationGroup::updateDirection")

	if ptr.Pointer() != nil {
		C.QParallelAnimationGroup_UpdateDirection(ptr.Pointer(), C.int(direction))
	}
}

func (ptr *QParallelAnimationGroup) UpdateDirectionDefault(direction QAbstractAnimation__Direction) {
	defer qt.Recovering("QParallelAnimationGroup::updateDirection")

	if ptr.Pointer() != nil {
		C.QParallelAnimationGroup_UpdateDirectionDefault(ptr.Pointer(), C.int(direction))
	}
}

func (ptr *QParallelAnimationGroup) ConnectUpdateState(f func(newState QAbstractAnimation__State, oldState QAbstractAnimation__State)) {
	defer qt.Recovering("connect QParallelAnimationGroup::updateState")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updateState", f)
	}
}

func (ptr *QParallelAnimationGroup) DisconnectUpdateState() {
	defer qt.Recovering("disconnect QParallelAnimationGroup::updateState")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updateState")
	}
}

//export callbackQParallelAnimationGroupUpdateState
func callbackQParallelAnimationGroupUpdateState(ptr unsafe.Pointer, ptrName *C.char, newState C.int, oldState C.int) {
	defer qt.Recovering("callback QParallelAnimationGroup::updateState")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateState"); signal != nil {
		signal.(func(QAbstractAnimation__State, QAbstractAnimation__State))(QAbstractAnimation__State(newState), QAbstractAnimation__State(oldState))
	} else {
		NewQParallelAnimationGroupFromPointer(ptr).UpdateStateDefault(QAbstractAnimation__State(newState), QAbstractAnimation__State(oldState))
	}
}

func (ptr *QParallelAnimationGroup) UpdateState(newState QAbstractAnimation__State, oldState QAbstractAnimation__State) {
	defer qt.Recovering("QParallelAnimationGroup::updateState")

	if ptr.Pointer() != nil {
		C.QParallelAnimationGroup_UpdateState(ptr.Pointer(), C.int(newState), C.int(oldState))
	}
}

func (ptr *QParallelAnimationGroup) UpdateStateDefault(newState QAbstractAnimation__State, oldState QAbstractAnimation__State) {
	defer qt.Recovering("QParallelAnimationGroup::updateState")

	if ptr.Pointer() != nil {
		C.QParallelAnimationGroup_UpdateStateDefault(ptr.Pointer(), C.int(newState), C.int(oldState))
	}
}

func (ptr *QParallelAnimationGroup) DestroyQParallelAnimationGroup() {
	defer qt.Recovering("QParallelAnimationGroup::~QParallelAnimationGroup")

	if ptr.Pointer() != nil {
		C.QParallelAnimationGroup_DestroyQParallelAnimationGroup(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QParallelAnimationGroup) ConnectTimerEvent(f func(event *QTimerEvent)) {
	defer qt.Recovering("connect QParallelAnimationGroup::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QParallelAnimationGroup) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QParallelAnimationGroup::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQParallelAnimationGroupTimerEvent
func callbackQParallelAnimationGroupTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QParallelAnimationGroup::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*QTimerEvent))(NewQTimerEventFromPointer(event))
	} else {
		NewQParallelAnimationGroupFromPointer(ptr).TimerEventDefault(NewQTimerEventFromPointer(event))
	}
}

func (ptr *QParallelAnimationGroup) TimerEvent(event QTimerEvent_ITF) {
	defer qt.Recovering("QParallelAnimationGroup::timerEvent")

	if ptr.Pointer() != nil {
		C.QParallelAnimationGroup_TimerEvent(ptr.Pointer(), PointerFromQTimerEvent(event))
	}
}

func (ptr *QParallelAnimationGroup) TimerEventDefault(event QTimerEvent_ITF) {
	defer qt.Recovering("QParallelAnimationGroup::timerEvent")

	if ptr.Pointer() != nil {
		C.QParallelAnimationGroup_TimerEventDefault(ptr.Pointer(), PointerFromQTimerEvent(event))
	}
}

func (ptr *QParallelAnimationGroup) ConnectChildEvent(f func(event *QChildEvent)) {
	defer qt.Recovering("connect QParallelAnimationGroup::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QParallelAnimationGroup) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QParallelAnimationGroup::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQParallelAnimationGroupChildEvent
func callbackQParallelAnimationGroupChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QParallelAnimationGroup::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*QChildEvent))(NewQChildEventFromPointer(event))
	} else {
		NewQParallelAnimationGroupFromPointer(ptr).ChildEventDefault(NewQChildEventFromPointer(event))
	}
}

func (ptr *QParallelAnimationGroup) ChildEvent(event QChildEvent_ITF) {
	defer qt.Recovering("QParallelAnimationGroup::childEvent")

	if ptr.Pointer() != nil {
		C.QParallelAnimationGroup_ChildEvent(ptr.Pointer(), PointerFromQChildEvent(event))
	}
}

func (ptr *QParallelAnimationGroup) ChildEventDefault(event QChildEvent_ITF) {
	defer qt.Recovering("QParallelAnimationGroup::childEvent")

	if ptr.Pointer() != nil {
		C.QParallelAnimationGroup_ChildEventDefault(ptr.Pointer(), PointerFromQChildEvent(event))
	}
}

func (ptr *QParallelAnimationGroup) ConnectCustomEvent(f func(event *QEvent)) {
	defer qt.Recovering("connect QParallelAnimationGroup::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QParallelAnimationGroup) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QParallelAnimationGroup::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQParallelAnimationGroupCustomEvent
func callbackQParallelAnimationGroupCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QParallelAnimationGroup::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
	} else {
		NewQParallelAnimationGroupFromPointer(ptr).CustomEventDefault(NewQEventFromPointer(event))
	}
}

func (ptr *QParallelAnimationGroup) CustomEvent(event QEvent_ITF) {
	defer qt.Recovering("QParallelAnimationGroup::customEvent")

	if ptr.Pointer() != nil {
		C.QParallelAnimationGroup_CustomEvent(ptr.Pointer(), PointerFromQEvent(event))
	}
}

func (ptr *QParallelAnimationGroup) CustomEventDefault(event QEvent_ITF) {
	defer qt.Recovering("QParallelAnimationGroup::customEvent")

	if ptr.Pointer() != nil {
		C.QParallelAnimationGroup_CustomEventDefault(ptr.Pointer(), PointerFromQEvent(event))
	}
}
