package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
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
		n.SetObjectName("QParallelAnimationGroup_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QParallelAnimationGroup) QParallelAnimationGroup_PTR() *QParallelAnimationGroup {
	return ptr
}

func (ptr *QParallelAnimationGroup) Duration() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QParallelAnimationGroup::duration")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QParallelAnimationGroup_Duration(ptr.Pointer()))
	}
	return 0
}

func (ptr *QParallelAnimationGroup) DestroyQParallelAnimationGroup() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QParallelAnimationGroup::~QParallelAnimationGroup")
		}
	}()

	if ptr.Pointer() != nil {
		C.QParallelAnimationGroup_DestroyQParallelAnimationGroup(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
