package core

//#include "qsequentialanimationgroup.h"
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
	if n.ObjectName() == "" {
		n.SetObjectName("QSequentialAnimationGroup_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSequentialAnimationGroup) QSequentialAnimationGroup_PTR() *QSequentialAnimationGroup {
	return ptr
}

func (ptr *QSequentialAnimationGroup) CurrentAnimation() *QAbstractAnimation {
	if ptr.Pointer() != nil {
		return NewQAbstractAnimationFromPointer(C.QSequentialAnimationGroup_CurrentAnimation(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSequentialAnimationGroup) AddPause(msecs int) *QPauseAnimation {
	if ptr.Pointer() != nil {
		return NewQPauseAnimationFromPointer(C.QSequentialAnimationGroup_AddPause(ptr.Pointer(), C.int(msecs)))
	}
	return nil
}

func (ptr *QSequentialAnimationGroup) ConnectCurrentAnimationChanged(f func(current *QAbstractAnimation)) {
	if ptr.Pointer() != nil {
		C.QSequentialAnimationGroup_ConnectCurrentAnimationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentAnimationChanged", f)
	}
}

func (ptr *QSequentialAnimationGroup) DisconnectCurrentAnimationChanged() {
	if ptr.Pointer() != nil {
		C.QSequentialAnimationGroup_DisconnectCurrentAnimationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentAnimationChanged")
	}
}

//export callbackQSequentialAnimationGroupCurrentAnimationChanged
func callbackQSequentialAnimationGroupCurrentAnimationChanged(ptrName *C.char, current unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "currentAnimationChanged").(func(*QAbstractAnimation))(NewQAbstractAnimationFromPointer(current))
}

func (ptr *QSequentialAnimationGroup) Duration() int {
	if ptr.Pointer() != nil {
		return int(C.QSequentialAnimationGroup_Duration(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSequentialAnimationGroup) InsertPause(index int, msecs int) *QPauseAnimation {
	if ptr.Pointer() != nil {
		return NewQPauseAnimationFromPointer(C.QSequentialAnimationGroup_InsertPause(ptr.Pointer(), C.int(index), C.int(msecs)))
	}
	return nil
}

func (ptr *QSequentialAnimationGroup) DestroyQSequentialAnimationGroup() {
	if ptr.Pointer() != nil {
		C.QSequentialAnimationGroup_DestroyQSequentialAnimationGroup(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
