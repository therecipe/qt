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
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QParallelAnimationGroup_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QParallelAnimationGroup) QParallelAnimationGroup_PTR() *QParallelAnimationGroup {
	return ptr
}

func (ptr *QParallelAnimationGroup) Duration() int {
	if ptr.Pointer() != nil {
		return int(C.QParallelAnimationGroup_Duration(ptr.Pointer()))
	}
	return 0
}

func (ptr *QParallelAnimationGroup) DestroyQParallelAnimationGroup() {
	if ptr.Pointer() != nil {
		C.QParallelAnimationGroup_DestroyQParallelAnimationGroup(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
