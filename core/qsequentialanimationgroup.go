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
func callbackQSequentialAnimationGroupCurrentAnimationChanged(ptrName *C.char, current unsafe.Pointer) {
	defer qt.Recovering("callback QSequentialAnimationGroup::currentAnimationChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "currentAnimationChanged"); signal != nil {
		signal.(func(*QAbstractAnimation))(NewQAbstractAnimationFromPointer(current))
	}

}

func (ptr *QSequentialAnimationGroup) Duration() int {
	defer qt.Recovering("QSequentialAnimationGroup::duration")

	if ptr.Pointer() != nil {
		return int(C.QSequentialAnimationGroup_Duration(ptr.Pointer()))
	}
	return 0
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
func callbackQSequentialAnimationGroupUpdateCurrentTime(ptrName *C.char, currentTime C.int) bool {
	defer qt.Recovering("callback QSequentialAnimationGroup::updateCurrentTime")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateCurrentTime"); signal != nil {
		signal.(func(int))(int(currentTime))
		return true
	}
	return false

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
func callbackQSequentialAnimationGroupUpdateDirection(ptrName *C.char, direction C.int) bool {
	defer qt.Recovering("callback QSequentialAnimationGroup::updateDirection")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateDirection"); signal != nil {
		signal.(func(QAbstractAnimation__Direction))(QAbstractAnimation__Direction(direction))
		return true
	}
	return false

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
func callbackQSequentialAnimationGroupUpdateState(ptrName *C.char, newState C.int, oldState C.int) bool {
	defer qt.Recovering("callback QSequentialAnimationGroup::updateState")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateState"); signal != nil {
		signal.(func(QAbstractAnimation__State, QAbstractAnimation__State))(QAbstractAnimation__State(newState), QAbstractAnimation__State(oldState))
		return true
	}
	return false

}

func (ptr *QSequentialAnimationGroup) DestroyQSequentialAnimationGroup() {
	defer qt.Recovering("QSequentialAnimationGroup::~QSequentialAnimationGroup")

	if ptr.Pointer() != nil {
		C.QSequentialAnimationGroup_DestroyQSequentialAnimationGroup(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
