package core

//#include "qitemselection.h"
import "C"
import (
	"unsafe"
)

type QItemSelection struct {
	QList
}

type QItemSelectionITF interface {
	QListITF
	QItemSelectionPTR() *QItemSelection
}

func PointerFromQItemSelection(ptr QItemSelectionITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QItemSelectionPTR().Pointer()
	}
	return nil
}

func QItemSelectionFromPointer(ptr unsafe.Pointer) *QItemSelection {
	var n = new(QItemSelection)
	n.SetPointer(ptr)
	return n
}

func (ptr *QItemSelection) QItemSelectionPTR() *QItemSelection {
	return ptr
}

func NewQItemSelection() *QItemSelection {
	return QItemSelectionFromPointer(unsafe.Pointer(C.QItemSelection_NewQItemSelection()))
}

func NewQItemSelection2(topLeft QModelIndexITF, bottomRight QModelIndexITF) *QItemSelection {
	return QItemSelectionFromPointer(unsafe.Pointer(C.QItemSelection_NewQItemSelection2(C.QtObjectPtr(PointerFromQModelIndex(topLeft)), C.QtObjectPtr(PointerFromQModelIndex(bottomRight)))))
}

func (ptr *QItemSelection) Contains(index QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QItemSelection_Contains(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(index))) != 0
	}
	return false
}

func (ptr *QItemSelection) Merge(other QItemSelectionITF, command QItemSelectionModel__SelectionFlag) {
	if ptr.Pointer() != nil {
		C.QItemSelection_Merge(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQItemSelection(other)), C.int(command))
	}
}

func (ptr *QItemSelection) Select(topLeft QModelIndexITF, bottomRight QModelIndexITF) {
	if ptr.Pointer() != nil {
		C.QItemSelection_Select(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(topLeft)), C.QtObjectPtr(PointerFromQModelIndex(bottomRight)))
	}
}

func QItemSelection_Split(ran QItemSelectionRangeITF, other QItemSelectionRangeITF, result QItemSelectionITF) {
	C.QItemSelection_QItemSelection_Split(C.QtObjectPtr(PointerFromQItemSelectionRange(ran)), C.QtObjectPtr(PointerFromQItemSelectionRange(other)), C.QtObjectPtr(PointerFromQItemSelection(result)))
}
