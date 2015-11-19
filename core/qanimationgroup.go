package core

//#include "qanimationgroup.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QAnimationGroup struct {
	QAbstractAnimation
}

type QAnimationGroup_ITF interface {
	QAbstractAnimation_ITF
	QAnimationGroup_PTR() *QAnimationGroup
}

func PointerFromQAnimationGroup(ptr QAnimationGroup_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAnimationGroup_PTR().Pointer()
	}
	return nil
}

func NewQAnimationGroupFromPointer(ptr unsafe.Pointer) *QAnimationGroup {
	var n = new(QAnimationGroup)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QAnimationGroup_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAnimationGroup) QAnimationGroup_PTR() *QAnimationGroup {
	return ptr
}
