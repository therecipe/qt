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
func callbackQParallelAnimationGroupUpdateCurrentTime(ptrName *C.char, currentTime C.int) bool {
	defer qt.Recovering("callback QParallelAnimationGroup::updateCurrentTime")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateCurrentTime"); signal != nil {
		signal.(func(int))(int(currentTime))
		return true
	}
	return false

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
func callbackQParallelAnimationGroupUpdateDirection(ptrName *C.char, direction C.int) bool {
	defer qt.Recovering("callback QParallelAnimationGroup::updateDirection")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateDirection"); signal != nil {
		signal.(func(QAbstractAnimation__Direction))(QAbstractAnimation__Direction(direction))
		return true
	}
	return false

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
func callbackQParallelAnimationGroupUpdateState(ptrName *C.char, newState C.int, oldState C.int) bool {
	defer qt.Recovering("callback QParallelAnimationGroup::updateState")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateState"); signal != nil {
		signal.(func(QAbstractAnimation__State, QAbstractAnimation__State))(QAbstractAnimation__State(newState), QAbstractAnimation__State(oldState))
		return true
	}
	return false

}

func (ptr *QParallelAnimationGroup) DestroyQParallelAnimationGroup() {
	defer qt.Recovering("QParallelAnimationGroup::~QParallelAnimationGroup")

	if ptr.Pointer() != nil {
		C.QParallelAnimationGroup_DestroyQParallelAnimationGroup(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
