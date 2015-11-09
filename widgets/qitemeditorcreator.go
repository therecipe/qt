package widgets

//#include "qitemeditorcreator.h"
import "C"
import (
	"unsafe"
)

type QItemEditorCreator struct {
	QItemEditorCreatorBase
}

type QItemEditorCreator_ITF interface {
	QItemEditorCreatorBase_ITF
	QItemEditorCreator_PTR() *QItemEditorCreator
}

func PointerFromQItemEditorCreator(ptr QItemEditorCreator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QItemEditorCreator_PTR().Pointer()
	}
	return nil
}

func NewQItemEditorCreatorFromPointer(ptr unsafe.Pointer) *QItemEditorCreator {
	var n = new(QItemEditorCreator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QItemEditorCreator) QItemEditorCreator_PTR() *QItemEditorCreator {
	return ptr
}
