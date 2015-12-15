package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QAbstractTableModel struct {
	QAbstractItemModel
}

type QAbstractTableModel_ITF interface {
	QAbstractItemModel_ITF
	QAbstractTableModel_PTR() *QAbstractTableModel
}

func PointerFromQAbstractTableModel(ptr QAbstractTableModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractTableModel_PTR().Pointer()
	}
	return nil
}

func NewQAbstractTableModelFromPointer(ptr unsafe.Pointer) *QAbstractTableModel {
	var n = new(QAbstractTableModel)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAbstractTableModel_") {
		n.SetObjectName("QAbstractTableModel_" + qt.Identifier())
	}
	return n
}

func (ptr *QAbstractTableModel) QAbstractTableModel_PTR() *QAbstractTableModel {
	return ptr
}

func (ptr *QAbstractTableModel) Index(row int, column int, parent QModelIndex_ITF) *QModelIndex {
	defer qt.Recovering("QAbstractTableModel::index")

	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QAbstractTableModel_Index(ptr.Pointer(), C.int(row), C.int(column), PointerFromQModelIndex(parent)))
	}
	return nil
}

func (ptr *QAbstractTableModel) DropMimeData(data QMimeData_ITF, action Qt__DropAction, row int, column int, parent QModelIndex_ITF) bool {
	defer qt.Recovering("QAbstractTableModel::dropMimeData")

	if ptr.Pointer() != nil {
		return C.QAbstractTableModel_DropMimeData(ptr.Pointer(), PointerFromQMimeData(data), C.int(action), C.int(row), C.int(column), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QAbstractTableModel) Flags(index QModelIndex_ITF) Qt__ItemFlag {
	defer qt.Recovering("QAbstractTableModel::flags")

	if ptr.Pointer() != nil {
		return Qt__ItemFlag(C.QAbstractTableModel_Flags(ptr.Pointer(), PointerFromQModelIndex(index)))
	}
	return 0
}

func (ptr *QAbstractTableModel) Sibling(row int, column int, idx QModelIndex_ITF) *QModelIndex {
	defer qt.Recovering("QAbstractTableModel::sibling")

	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QAbstractTableModel_Sibling(ptr.Pointer(), C.int(row), C.int(column), PointerFromQModelIndex(idx)))
	}
	return nil
}

func (ptr *QAbstractTableModel) DestroyQAbstractTableModel() {
	defer qt.Recovering("QAbstractTableModel::~QAbstractTableModel")

	if ptr.Pointer() != nil {
		C.QAbstractTableModel_DestroyQAbstractTableModel(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
