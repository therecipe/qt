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

type QSequentialAnimationGroupITF interface {
	QAnimationGroupITF
	QSequentialAnimationGroupPTR() *QSequentialAnimationGroup
}

func PointerFromQSequentialAnimationGroup(ptr QSequentialAnimationGroupITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSequentialAnimationGroupPTR().Pointer()
	}
	return nil
}

func QSequentialAnimationGroupFromPointer(ptr unsafe.Pointer) *QSequentialAnimationGroup {
	var n = new(QSequentialAnimationGroup)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QSequentialAnimationGroup_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSequentialAnimationGroup) QSequentialAnimationGroupPTR() *QSequentialAnimationGroup {
	return ptr
}

func (ptr *QSequentialAnimationGroup) CurrentAnimation() *QAbstractAnimation {
	if ptr.Pointer() != nil {
		return QAbstractAnimationFromPointer(unsafe.Pointer(C.QSequentialAnimationGroup_CurrentAnimation(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QSequentialAnimationGroup) AddPause(msecs int) *QPauseAnimation {
	if ptr.Pointer() != nil {
		return QPauseAnimationFromPointer(unsafe.Pointer(C.QSequentialAnimationGroup_AddPause(C.QtObjectPtr(ptr.Pointer()), C.int(msecs))))
	}
	return nil
}

func (ptr *QSequentialAnimationGroup) ConnectCurrentAnimationChanged(f func(current QAbstractAnimationITF)) {
	if ptr.Pointer() != nil {
		C.QSequentialAnimationGroup_ConnectCurrentAnimationChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "currentAnimationChanged", f)
	}
}

func (ptr *QSequentialAnimationGroup) DisconnectCurrentAnimationChanged() {
	if ptr.Pointer() != nil {
		C.QSequentialAnimationGroup_DisconnectCurrentAnimationChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "currentAnimationChanged")
	}
}

//export callbackQSequentialAnimationGroupCurrentAnimationChanged
func callbackQSequentialAnimationGroupCurrentAnimationChanged(ptrName *C.char, current unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "currentAnimationChanged").(func(*QAbstractAnimation))(QAbstractAnimationFromPointer(current))
}

func (ptr *QSequentialAnimationGroup) Duration() int {
	if ptr.Pointer() != nil {
		return int(C.QSequentialAnimationGroup_Duration(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSequentialAnimationGroup) InsertPause(index int, msecs int) *QPauseAnimation {
	if ptr.Pointer() != nil {
		return QPauseAnimationFromPointer(unsafe.Pointer(C.QSequentialAnimationGroup_InsertPause(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.int(msecs))))
	}
	return nil
}

func (ptr *QSequentialAnimationGroup) DestroyQSequentialAnimationGroup() {
	if ptr.Pointer() != nil {
		C.QSequentialAnimationGroup_DestroyQSequentialAnimationGroup(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
