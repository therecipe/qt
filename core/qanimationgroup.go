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

type QAnimationGroupITF interface {
	QAbstractAnimationITF
	QAnimationGroupPTR() *QAnimationGroup
}

func PointerFromQAnimationGroup(ptr QAnimationGroupITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAnimationGroupPTR().Pointer()
	}
	return nil
}

func QAnimationGroupFromPointer(ptr unsafe.Pointer) *QAnimationGroup {
	var n = new(QAnimationGroup)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QAnimationGroup_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAnimationGroup) QAnimationGroupPTR() *QAnimationGroup {
	return ptr
}
