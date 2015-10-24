package core

//#include "qabstractanimation.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QAbstractAnimation struct {
	QObject
}

type QAbstractAnimationITF interface {
	QObjectITF
	QAbstractAnimationPTR() *QAbstractAnimation
}

func PointerFromQAbstractAnimation(ptr QAbstractAnimationITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractAnimationPTR().Pointer()
	}
	return nil
}

func QAbstractAnimationFromPointer(ptr unsafe.Pointer) *QAbstractAnimation {
	var n = new(QAbstractAnimation)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QAbstractAnimation_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAbstractAnimation) QAbstractAnimationPTR() *QAbstractAnimation {
	return ptr
}

//QAbstractAnimation::DeletionPolicy
type QAbstractAnimation__DeletionPolicy int

var (
	QAbstractAnimation__KeepWhenStopped   = QAbstractAnimation__DeletionPolicy(0)
	QAbstractAnimation__DeleteWhenStopped = QAbstractAnimation__DeletionPolicy(1)
)

//QAbstractAnimation::Direction
type QAbstractAnimation__Direction int

var (
	QAbstractAnimation__Forward  = QAbstractAnimation__Direction(0)
	QAbstractAnimation__Backward = QAbstractAnimation__Direction(1)
)

//QAbstractAnimation::State
type QAbstractAnimation__State int

var (
	QAbstractAnimation__Stopped = QAbstractAnimation__State(0)
	QAbstractAnimation__Paused  = QAbstractAnimation__State(1)
	QAbstractAnimation__Running = QAbstractAnimation__State(2)
)

func (ptr *QAbstractAnimation) CurrentLoop() int {
	if ptr.Pointer() != nil {
		return int(C.QAbstractAnimation_CurrentLoop(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractAnimation) CurrentTime() int {
	if ptr.Pointer() != nil {
		return int(C.QAbstractAnimation_CurrentTime(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractAnimation) Direction() QAbstractAnimation__Direction {
	if ptr.Pointer() != nil {
		return QAbstractAnimation__Direction(C.QAbstractAnimation_Direction(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractAnimation) LoopCount() int {
	if ptr.Pointer() != nil {
		return int(C.QAbstractAnimation_LoopCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractAnimation) SetCurrentTime(msecs int) {
	if ptr.Pointer() != nil {
		C.QAbstractAnimation_SetCurrentTime(C.QtObjectPtr(ptr.Pointer()), C.int(msecs))
	}
}

func (ptr *QAbstractAnimation) SetDirection(direction QAbstractAnimation__Direction) {
	if ptr.Pointer() != nil {
		C.QAbstractAnimation_SetDirection(C.QtObjectPtr(ptr.Pointer()), C.int(direction))
	}
}

func (ptr *QAbstractAnimation) SetLoopCount(loopCount int) {
	if ptr.Pointer() != nil {
		C.QAbstractAnimation_SetLoopCount(C.QtObjectPtr(ptr.Pointer()), C.int(loopCount))
	}
}

func (ptr *QAbstractAnimation) ConnectCurrentLoopChanged(f func(currentLoop int)) {
	if ptr.Pointer() != nil {
		C.QAbstractAnimation_ConnectCurrentLoopChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "currentLoopChanged", f)
	}
}

func (ptr *QAbstractAnimation) DisconnectCurrentLoopChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAnimation_DisconnectCurrentLoopChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "currentLoopChanged")
	}
}

//export callbackQAbstractAnimationCurrentLoopChanged
func callbackQAbstractAnimationCurrentLoopChanged(ptrName *C.char, currentLoop C.int) {
	qt.GetSignal(C.GoString(ptrName), "currentLoopChanged").(func(int))(int(currentLoop))
}

func (ptr *QAbstractAnimation) CurrentLoopTime() int {
	if ptr.Pointer() != nil {
		return int(C.QAbstractAnimation_CurrentLoopTime(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractAnimation) ConnectDirectionChanged(f func(newDirection QAbstractAnimation__Direction)) {
	if ptr.Pointer() != nil {
		C.QAbstractAnimation_ConnectDirectionChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "directionChanged", f)
	}
}

func (ptr *QAbstractAnimation) DisconnectDirectionChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAnimation_DisconnectDirectionChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "directionChanged")
	}
}

//export callbackQAbstractAnimationDirectionChanged
func callbackQAbstractAnimationDirectionChanged(ptrName *C.char, newDirection C.int) {
	qt.GetSignal(C.GoString(ptrName), "directionChanged").(func(QAbstractAnimation__Direction))(QAbstractAnimation__Direction(newDirection))
}

func (ptr *QAbstractAnimation) Duration() int {
	if ptr.Pointer() != nil {
		return int(C.QAbstractAnimation_Duration(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractAnimation) ConnectFinished(f func()) {
	if ptr.Pointer() != nil {
		C.QAbstractAnimation_ConnectFinished(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QAbstractAnimation) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QAbstractAnimation_DisconnectFinished(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQAbstractAnimationFinished
func callbackQAbstractAnimationFinished(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "finished").(func())()
}

func (ptr *QAbstractAnimation) Group() *QAnimationGroup {
	if ptr.Pointer() != nil {
		return QAnimationGroupFromPointer(unsafe.Pointer(C.QAbstractAnimation_Group(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QAbstractAnimation) Pause() {
	if ptr.Pointer() != nil {
		C.QAbstractAnimation_Pause(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QAbstractAnimation) Resume() {
	if ptr.Pointer() != nil {
		C.QAbstractAnimation_Resume(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QAbstractAnimation) SetPaused(paused bool) {
	if ptr.Pointer() != nil {
		C.QAbstractAnimation_SetPaused(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(paused)))
	}
}

func (ptr *QAbstractAnimation) Start(policy QAbstractAnimation__DeletionPolicy) {
	if ptr.Pointer() != nil {
		C.QAbstractAnimation_Start(C.QtObjectPtr(ptr.Pointer()), C.int(policy))
	}
}

func (ptr *QAbstractAnimation) ConnectStateChanged(f func(newState QAbstractAnimation__State, oldState QAbstractAnimation__State)) {
	if ptr.Pointer() != nil {
		C.QAbstractAnimation_ConnectStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QAbstractAnimation) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractAnimation_DisconnectStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQAbstractAnimationStateChanged
func callbackQAbstractAnimationStateChanged(ptrName *C.char, newState C.int, oldState C.int) {
	qt.GetSignal(C.GoString(ptrName), "stateChanged").(func(QAbstractAnimation__State, QAbstractAnimation__State))(QAbstractAnimation__State(newState), QAbstractAnimation__State(oldState))
}

func (ptr *QAbstractAnimation) Stop() {
	if ptr.Pointer() != nil {
		C.QAbstractAnimation_Stop(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QAbstractAnimation) TotalDuration() int {
	if ptr.Pointer() != nil {
		return int(C.QAbstractAnimation_TotalDuration(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractAnimation) DestroyQAbstractAnimation() {
	if ptr.Pointer() != nil {
		C.QAbstractAnimation_DestroyQAbstractAnimation(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
