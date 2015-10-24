package core

//#include "qparallelanimationgroup.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QParallelAnimationGroup struct {
	QAnimationGroup
}

type QParallelAnimationGroupITF interface {
	QAnimationGroupITF
	QParallelAnimationGroupPTR() *QParallelAnimationGroup
}

func PointerFromQParallelAnimationGroup(ptr QParallelAnimationGroupITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QParallelAnimationGroupPTR().Pointer()
	}
	return nil
}

func QParallelAnimationGroupFromPointer(ptr unsafe.Pointer) *QParallelAnimationGroup {
	var n = new(QParallelAnimationGroup)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QParallelAnimationGroup_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QParallelAnimationGroup) QParallelAnimationGroupPTR() *QParallelAnimationGroup {
	return ptr
}

func (ptr *QParallelAnimationGroup) Duration() int {
	if ptr.Pointer() != nil {
		return int(C.QParallelAnimationGroup_Duration(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QParallelAnimationGroup) DestroyQParallelAnimationGroup() {
	if ptr.Pointer() != nil {
		C.QParallelAnimationGroup_DestroyQParallelAnimationGroup(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
