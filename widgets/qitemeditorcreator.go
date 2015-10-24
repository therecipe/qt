package widgets

//#include "qitemeditorcreator.h"
import "C"
import (
	"unsafe"
)

type QItemEditorCreator struct {
	QItemEditorCreatorBase
}

type QItemEditorCreatorITF interface {
	QItemEditorCreatorBaseITF
	QItemEditorCreatorPTR() *QItemEditorCreator
}

func PointerFromQItemEditorCreator(ptr QItemEditorCreatorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QItemEditorCreatorPTR().Pointer()
	}
	return nil
}

func QItemEditorCreatorFromPointer(ptr unsafe.Pointer) *QItemEditorCreator {
	var n = new(QItemEditorCreator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QItemEditorCreator) QItemEditorCreatorPTR() *QItemEditorCreator {
	return ptr
}
