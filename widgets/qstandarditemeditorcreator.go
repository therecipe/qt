package widgets

//#include "qstandarditemeditorcreator.h"
import "C"
import (
	"unsafe"
)

type QStandardItemEditorCreator struct {
	QItemEditorCreatorBase
}

type QStandardItemEditorCreator_ITF interface {
	QItemEditorCreatorBase_ITF
	QStandardItemEditorCreator_PTR() *QStandardItemEditorCreator
}

func PointerFromQStandardItemEditorCreator(ptr QStandardItemEditorCreator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStandardItemEditorCreator_PTR().Pointer()
	}
	return nil
}

func NewQStandardItemEditorCreatorFromPointer(ptr unsafe.Pointer) *QStandardItemEditorCreator {
	var n = new(QStandardItemEditorCreator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStandardItemEditorCreator) QStandardItemEditorCreator_PTR() *QStandardItemEditorCreator {
	return ptr
}
