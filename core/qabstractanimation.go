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
func callbackQAbstractAnimationCurrentLoopChanged(ptrName *C.char, currentLoop C.int) {
	defer qt.Recovering("callback QAbstractAnimation::currentLoopChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "currentLoopChanged")
	if signal != nil {
		signal.(func(int))(int(currentLoop))
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
func callbackQAbstractAnimationDirectionChanged(ptrName *C.char, newDirection C.int) {
	defer qt.Recovering("callback QAbstractAnimation::directionChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "directionChanged")
	if signal != nil {
		signal.(func(QAbstractAnimation__Direction))(QAbstractAnimation__Direction(newDirection))
	}

}

func (ptr *QAbstractAnimation) Duration() int {
	defer qt.Recovering("QAbstractAnimation::duration")

	if ptr.Pointer() != nil {
		return int(C.QAbstractAnimation_Duration(ptr.Pointer()))
	}
	return 0
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
func callbackQAbstractAnimationFinished(ptrName *C.char) {
	defer qt.Recovering("callback QAbstractAnimation::finished")

	var signal = qt.GetSignal(C.GoString(ptrName), "finished")
	if signal != nil {
		signal.(func())()
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
func callbackQAbstractAnimationStateChanged(ptrName *C.char, newState C.int, oldState C.int) {
	defer qt.Recovering("callback QAbstractAnimation::stateChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "stateChanged")
	if signal != nil {
		signal.(func(QAbstractAnimation__State, QAbstractAnimation__State))(QAbstractAnimation__State(newState), QAbstractAnimation__State(oldState))
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
func callbackQAbstractAnimationUpdateDirection(ptrName *C.char, direction C.int) bool {
	defer qt.Recovering("callback QAbstractAnimation::updateDirection")

	var signal = qt.GetSignal(C.GoString(ptrName), "updateDirection")
	if signal != nil {
		defer signal.(func(QAbstractAnimation__Direction))(QAbstractAnimation__Direction(direction))
		return true
	}
	return false

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
func callbackQAbstractAnimationUpdateState(ptrName *C.char, newState C.int, oldState C.int) bool {
	defer qt.Recovering("callback QAbstractAnimation::updateState")

	var signal = qt.GetSignal(C.GoString(ptrName), "updateState")
	if signal != nil {
		defer signal.(func(QAbstractAnimation__State, QAbstractAnimation__State))(QAbstractAnimation__State(newState), QAbstractAnimation__State(oldState))
		return true
	}
	return false

}

func (ptr *QAbstractAnimation) DestroyQAbstractAnimation() {
	defer qt.Recovering("QAbstractAnimation::~QAbstractAnimation")

	if ptr.Pointer() != nil {
		C.QAbstractAnimation_DestroyQAbstractAnimation(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
