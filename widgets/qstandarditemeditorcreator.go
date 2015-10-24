package widgets

//#include "qstandarditemeditorcreator.h"
import "C"
import (
	"unsafe"
)

type QStandardItemEditorCreator struct {
	QItemEditorCreatorBase
}

type QStandardItemEditorCreatorITF interface {
	QItemEditorCreatorBaseITF
	QStandardItemEditorCreatorPTR() *QStandardItemEditorCreator
}

func PointerFromQStandardItemEditorCreator(ptr QStandardItemEditorCreatorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStandardItemEditorCreatorPTR().Pointer()
	}
	return nil
}

func QStandardItemEditorCreatorFromPointer(ptr unsafe.Pointer) *QStandardItemEditorCreator {
	var n = new(QStandardItemEditorCreator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStandardItemEditorCreator) QStandardItemEditorCreatorPTR() *QStandardItemEditorCreator {
	return ptr
}
