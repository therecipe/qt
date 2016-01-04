package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QAbstractAnimation struct {
	QObject
}

type QAbstractAnimation_ITF interface {
	QObject_ITF
	QAbstractAnimation_PTR() *QAbstractAnimation
}

func PointerFromQAbstractAnimation(ptr QAbstractAnimation_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractAnimation_PTR().Pointer()
	}
	return nil
}

func NewQAbstractAnimationFromPointer(ptr unsafe.Pointer) *QAbstractAnimation {
	var n = new(QAbstractAnimation)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAbstractAnimation_") {
		n.SetObjectName("QAbstractAnimation_" + qt.Identifier())
	}
	return n
}

func (ptr *QAbstractAnimation) QAbstractAnimation_PTR() *QAbstractAnimation {
	return ptr
}

//QAbstractAnimation::DeletionPolicy
type QAbstractAnimation__DeletionPolicy int64

const (
	QAbstractAnimation__KeepWhenStopped   = QAbstractAnimation__DeletionPolicy(0)
	QAbstractAnimation__DeleteWhenStopped = QAbstractAnimation__DeletionPolicy(1)
)

//QAbstractAnimation::Direction
type QAbstractAnimation__Direction int64

const (
	QAbstractAnimation__Forward  = QAbstractAnimation__Direction(0)
	QAbstractAnimation__Backward = QAbstractAnimation__Direction(1)
)

//QAbstractAnimation::State
type QAbstractAnimation__State int64

const (
	QAbstractAnimation__Stopped = QAbstractAnimation__State(0)
	QAbstractAnimation__Paused  = QAbstractAnimation__State(1)
	QAbstractAnimation__Running = QAbstractAnimation__State(2)
)

func (ptr *QAbstractAnimation) CurrentLoop() int {
	defer qt.Recovering("QAbstractAnimation::currentLoop")

	if ptr.Pointer() != nil {
		return int(C.QAbstractAnimation_CurrentLoop(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractAnimation) CurrentTime() int {
	defer qt.Recovering("QAbstractAnimation::currentTime")

	if ptr.Pointer() != nil {
		return int(C.QAbstractAnimation_CurrentTime(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractAnimation) Direction() QAbstractAnimation__Direction {
	defer qt.Recovering("QAbstractAnimation::direction")

	if ptr.Pointer() != nil {
		return QAbstractAnimation__Direction(C.QAbstractAnimation_Direction(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractAnimation) LoopCount() int {
	defer qt.Recovering("QAbstractAnimation::loopCount")

	if ptr.Pointer() != nil {
		return int(C.QAbstractAnimation_LoopCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractAnimation) SetCurrentTime(msecs int) {
	defer qt.Recovering("QAbstractAnimation::setCurrentTime")

	if ptr.Pointer() != nil {
		C.QAbstractAnimation_SetCurrentTime(ptr.Pointer(), C.int(msecs))
	}
}

func (ptr *QAbstractAnimation) SetDirection(direction QAbstractAnimation__Direction) {
	defer qt.Recovering("QAbstractAnimation::setDirection")

	if ptr.Pointer() != nil {
		C.QAbstractAnimation_SetDirection(ptr.Pointer(), C.int(direction))
	}
}

func (ptr *QAbstractAnimation) SetLoopCount(loopCount int) {
	defer qt.Recovering("QAbstractAnimation::setLoopCount")

	if ptr.Pointer() != nil {
		C.QAbstractAnimation_SetLoopCount(ptr.Pointer(), C.int(loopCount))
	}
}

func (ptr *QAbstractAnimation) State() QAbstractAnimation__State {
	defer qt.Recovering("QAbstractAnimation::state")

	if ptr.Pointer() != nil {
		return QAbstractAnimation__State(C.QAbstractAnimation_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractAnimation) ConnectCurrentLoopChanged(f func(currentLoop int)) {
	defer qt.Recovering("connect QAbstractAnimation::currentLoopChanged")

	if ptr.Pointer() != nil {
		C.QAbstractAnimation_ConnectCurrentLoopChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentLoopChanged", f)
	}
}

func (ptr *QAbstractAnimation) DisconnectCurrentLoopChanged() {
	defer qt.Recovering("disconnect QAbstractAnimation::currentLoopChanged")

	if ptr.Pointer() != nil {
		C.QAbstractAnimation_DisconnectCurrentLoopChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentLoopChanged")
	}
}

//export callbackQAbstractAnimationCurrentLoopChanged
func callbackQAbstractAnimationCurrentLoopChanged(ptr unsafe.Pointer, ptrName *C.char, currentLoop C.int) {
	defer qt.Recovering("callback QAbstractAnimation::currentLoopChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "currentLoopChanged"); signal != nil {
		signal.(func(int))(int(currentLoop))
	}

}

func (ptr *QAbstractAnimation) CurrentLoopChanged(currentLoop int) {
	defer qt.Recovering("QAbstractAnimation::currentLoopChanged")

	if ptr.Pointer() != nil {
		C.QAbstractAnimation_CurrentLoopChanged(ptr.Pointer(), C.int(currentLoop))
	}
}

func (ptr *QAbstractAnimation) CurrentLoopTime() int {
	defer qt.Recovering("QAbstractAnimation::currentLoopTime")

	if ptr.Pointer() != nil {
		return int(C.QAbstractAnimation_CurrentLoopTime(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractAnimation) ConnectDirectionChanged(f func(newDirection QAbstractAnimation__Direction)) {
	defer qt.Recovering("connect QAbstractAnimation::directionChanged")

	if ptr.Pointer() != nil {
		C.QAbstractAnimation_ConnectDirectionChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "directionChanged", f)
	}
}

func (ptr *QAbstractAnimation) DisconnectDirectionChanged() {
	defer qt.Recovering("disconnect QAbstractAnimation::directionChanged")

	if ptr.Pointer() != nil {
		C.QAbstractAnimation_DisconnectDirectionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "directionChanged")
	}
}

//export callbackQAbstractAnimationDirectionChanged
func callbackQAbstractAnimationDirectionChanged(ptr unsafe.Pointer, ptrName *C.char, newDirection C.int) {
	defer qt.Recovering("callback QAbstractAnimation::directionChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "directionChanged"); signal != nil {
		signal.(func(QAbstractAnimation__Direction))(QAbstractAnimation__Direction(newDirection))
	}

}

func (ptr *QAbstractAnimation) DirectionChanged(newDirection QAbstractAnimation__Direction) {
	defer qt.Recovering("QAbstractAnimation::directionChanged")

	if ptr.Pointer() != nil {
		C.QAbstractAnimation_DirectionChanged(ptr.Pointer(), C.int(newDirection))
	}
}

func (ptr *QAbstractAnimation) Duration() int {
	defer qt.Recovering("QAbstractAnimation::duration")

	if ptr.Pointer() != nil {
		return int(C.QAbstractAnimation_Duration(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractAnimation) Event(event QEvent_ITF) bool {
	defer qt.Recovering("QAbstractAnimation::event")

	if ptr.Pointer() != nil {
		return C.QAbstractAnimation_Event(ptr.Pointer(), PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QAbstractAnimation) ConnectFinished(f func()) {
	defer qt.Recovering("connect QAbstractAnimation::finished")

	if ptr.Pointer() != nil {
		C.QAbstractAnimation_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QAbstractAnimation) DisconnectFinished() {
	defer qt.Recovering("disconnect QAbstractAnimation::finished")

	if ptr.Pointer() != nil {
		C.QAbstractAnimation_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQAbstractAnimationFinished
func callbackQAbstractAnimationFinished(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAbstractAnimation::finished")

	if signal := qt.GetSignal(C.GoString(ptrName), "finished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractAnimation) Finished() {
	defer qt.Recovering("QAbstractAnimation::finished")

	if ptr.Pointer() != nil {
		C.QAbstractAnimation_Finished(ptr.Pointer())
	}
}

func (ptr *QAbstractAnimation) Group() *QAnimationGroup {
	defer qt.Recovering("QAbstractAnimation::group")

	if ptr.Pointer() != nil {
		return NewQAnimationGroupFromPointer(C.QAbstractAnimation_Group(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractAnimation) Pause() {
	defer qt.Recovering("QAbstractAnimation::pause")

	if ptr.Pointer() != nil {
		C.QAbstractAnimation_Pause(ptr.Pointer())
	}
}

func (ptr *QAbstractAnimation) Resume() {
	defer qt.Recovering("QAbstractAnimation::resume")

	if ptr.Pointer() != nil {
		C.QAbstractAnimation_Resume(ptr.Pointer())
	}
}

func (ptr *QAbstractAnimation) SetPaused(paused bool) {
	defer qt.Recovering("QAbstractAnimation::setPaused")

	if ptr.Pointer() != nil {
		C.QAbstractAnimation_SetPaused(ptr.Pointer(), C.int(qt.GoBoolToInt(paused)))
	}
}

func (ptr *QAbstractAnimation) Start(policy QAbstractAnimation__DeletionPolicy) {
	defer qt.Recovering("QAbstractAnimation::start")

	if ptr.Pointer() != nil {
		C.QAbstractAnimation_Start(ptr.Pointer(), C.int(policy))
	}
}

func (ptr *QAbstractAnimation) ConnectStateChanged(f func(newState QAbstractAnimation__State, oldState QAbstractAnimation__State)) {
	defer qt.Recovering("connect QAbstractAnimation::stateChanged")

	if ptr.Pointer() != nil {
		C.QAbstractAnimation_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QAbstractAnimation) DisconnectStateChanged() {
	defer qt.Recovering("disconnect QAbstractAnimation::stateChanged")

	if ptr.Pointer() != nil {
		C.QAbstractAnimation_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQAbstractAnimationStateChanged
func callbackQAbstractAnimationStateChanged(ptr unsafe.Pointer, ptrName *C.char, newState C.int, oldState C.int) {
	defer qt.Recovering("callback QAbstractAnimation::stateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stateChanged"); signal != nil {
		signal.(func(QAbstractAnimation__State, QAbstractAnimation__State))(QAbstractAnimation__State(newState), QAbstractAnimation__State(oldState))
	}

}

func (ptr *QAbstractAnimation) StateChanged(newState QAbstractAnimation__State, oldState QAbstractAnimation__State) {
	defer qt.Recovering("QAbstractAnimation::stateChanged")

	if ptr.Pointer() != nil {
		C.QAbstractAnimation_StateChanged(ptr.Pointer(), C.int(newState), C.int(oldState))
	}
}

func (ptr *QAbstractAnimation) Stop() {
	defer qt.Recovering("QAbstractAnimation::stop")

	if ptr.Pointer() != nil {
		C.QAbstractAnimation_Stop(ptr.Pointer())
	}
}

func (ptr *QAbstractAnimation) TotalDuration() int {
	defer qt.Recovering("QAbstractAnimation::totalDuration")

	if ptr.Pointer() != nil {
		return int(C.QAbstractAnimation_TotalDuration(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractAnimation) ConnectUpdateDirection(f func(direction QAbstractAnimation__Direction)) {
	defer qt.Recovering("connect QAbstractAnimation::updateDirection")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updateDirection", f)
	}
}

func (ptr *QAbstractAnimation) DisconnectUpdateDirection() {
	defer qt.Recovering("disconnect QAbstractAnimation::updateDirection")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updateDirection")
	}
}

//export callbackQAbstractAnimationUpdateDirection
func callbackQAbstractAnimationUpdateDirection(ptr unsafe.Pointer, ptrName *C.char, direction C.int) {
	defer qt.Recovering("callback QAbstractAnimation::updateDirection")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateDirection"); signal != nil {
		signal.(func(QAbstractAnimation__Direction))(QAbstractAnimation__Direction(direction))
	} else {
		NewQAbstractAnimationFromPointer(ptr).UpdateDirectionDefault(QAbstractAnimation__Direction(direction))
	}
}

func (ptr *QAbstractAnimation) UpdateDirection(direction QAbstractAnimation__Direction) {
	defer qt.Recovering("QAbstractAnimation::updateDirection")

	if ptr.Pointer() != nil {
		C.QAbstractAnimation_UpdateDirection(ptr.Pointer(), C.int(direction))
	}
}

func (ptr *QAbstractAnimation) UpdateDirectionDefault(direction QAbstractAnimation__Direction) {
	defer qt.Recovering("QAbstractAnimation::updateDirection")

	if ptr.Pointer() != nil {
		C.QAbstractAnimation_UpdateDirectionDefault(ptr.Pointer(), C.int(direction))
	}
}

func (ptr *QAbstractAnimation) ConnectUpdateState(f func(newState QAbstractAnimation__State, oldState QAbstractAnimation__State)) {
	defer qt.Recovering("connect QAbstractAnimation::updateState")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updateState", f)
	}
}

func (ptr *QAbstractAnimation) DisconnectUpdateState() {
	defer qt.Recovering("disconnect QAbstractAnimation::updateState")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updateState")
	}
}

//export callbackQAbstractAnimationUpdateState
func callbackQAbstractAnimationUpdateState(ptr unsafe.Pointer, ptrName *C.char, newState C.int, oldState C.int) {
	defer qt.Recovering("callback QAbstractAnimation::updateState")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateState"); signal != nil {
		signal.(func(QAbstractAnimation__State, QAbstractAnimation__State))(QAbstractAnimation__State(newState), QAbstractAnimation__State(oldState))
	} else {
		NewQAbstractAnimationFromPointer(ptr).UpdateStateDefault(QAbstractAnimation__State(newState), QAbstractAnimation__State(oldState))
	}
}

func (ptr *QAbstractAnimation) UpdateState(newState QAbstractAnimation__State, oldState QAbstractAnimation__State) {
	defer qt.Recovering("QAbstractAnimation::updateState")

	if ptr.Pointer() != nil {
		C.QAbstractAnimation_UpdateState(ptr.Pointer(), C.int(newState), C.int(oldState))
	}
}

func (ptr *QAbstractAnimation) UpdateStateDefault(newState QAbstractAnimation__State, oldState QAbstractAnimation__State) {
	defer qt.Recovering("QAbstractAnimation::updateState")

	if ptr.Pointer() != nil {
		C.QAbstractAnimation_UpdateStateDefault(ptr.Pointer(), C.int(newState), C.int(oldState))
	}
}

func (ptr *QAbstractAnimation) DestroyQAbstractAnimation() {
	defer qt.Recovering("QAbstractAnimation::~QAbstractAnimation")

	if ptr.Pointer() != nil {
		C.QAbstractAnimation_DestroyQAbstractAnimation(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractAnimation) ConnectTimerEvent(f func(event *QTimerEvent)) {
	defer qt.Recovering("connect QAbstractAnimation::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAbstractAnimation) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAbstractAnimation::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAbstractAnimationTimerEvent
func callbackQAbstractAnimationTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractAnimation::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*QTimerEvent))(NewQTimerEventFromPointer(event))
	} else {
		NewQAbstractAnimationFromPointer(ptr).TimerEventDefault(NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAbstractAnimation) TimerEvent(event QTimerEvent_ITF) {
	defer qt.Recovering("QAbstractAnimation::timerEvent")

	if ptr.Pointer() != nil {
		C.QAbstractAnimation_TimerEvent(ptr.Pointer(), PointerFromQTimerEvent(event))
	}
}

func (ptr *QAbstractAnimation) TimerEventDefault(event QTimerEvent_ITF) {
	defer qt.Recovering("QAbstractAnimation::timerEvent")

	if ptr.Pointer() != nil {
		C.QAbstractAnimation_TimerEventDefault(ptr.Pointer(), PointerFromQTimerEvent(event))
	}
}

func (ptr *QAbstractAnimation) ConnectChildEvent(f func(event *QChildEvent)) {
	defer qt.Recovering("connect QAbstractAnimation::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAbstractAnimation) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAbstractAnimation::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAbstractAnimationChildEvent
func callbackQAbstractAnimationChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractAnimation::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*QChildEvent))(NewQChildEventFromPointer(event))
	} else {
		NewQAbstractAnimationFromPointer(ptr).ChildEventDefault(NewQChildEventFromPointer(event))
	}
}

func (ptr *QAbstractAnimation) ChildEvent(event QChildEvent_ITF) {
	defer qt.Recovering("QAbstractAnimation::childEvent")

	if ptr.Pointer() != nil {
		C.QAbstractAnimation_ChildEvent(ptr.Pointer(), PointerFromQChildEvent(event))
	}
}

func (ptr *QAbstractAnimation) ChildEventDefault(event QChildEvent_ITF) {
	defer qt.Recovering("QAbstractAnimation::childEvent")

	if ptr.Pointer() != nil {
		C.QAbstractAnimation_ChildEventDefault(ptr.Pointer(), PointerFromQChildEvent(event))
	}
}

func (ptr *QAbstractAnimation) ConnectCustomEvent(f func(event *QEvent)) {
	defer qt.Recovering("connect QAbstractAnimation::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAbstractAnimation) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAbstractAnimation::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAbstractAnimationCustomEvent
func callbackQAbstractAnimationCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractAnimation::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
	} else {
		NewQAbstractAnimationFromPointer(ptr).CustomEventDefault(NewQEventFromPointer(event))
	}
}

func (ptr *QAbstractAnimation) CustomEvent(event QEvent_ITF) {
	defer qt.Recovering("QAbstractAnimation::customEvent")

	if ptr.Pointer() != nil {
		C.QAbstractAnimation_CustomEvent(ptr.Pointer(), PointerFromQEvent(event))
	}
}

func (ptr *QAbstractAnimation) CustomEventDefault(event QEvent_ITF) {
	defer qt.Recovering("QAbstractAnimation::customEvent")

	if ptr.Pointer() != nil {
		C.QAbstractAnimation_CustomEventDefault(ptr.Pointer(), PointerFromQEvent(event))
	}
}
