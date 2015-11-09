package core

//#include "qabstractlistmodel.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QAbstractListModel struct {
	QAbstractItemModel
}

type QAbstractListModel_ITF interface {
	QAbstractItemModel_ITF
	QAbstractListModel_PTR() *QAbstractListModel
}

func PointerFromQAbstractListModel(ptr QAbstractListModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractListModel_PTR().Pointer()
	}
	return nil
}

func NewQAbstractListModelFromPointer(ptr unsafe.Pointer) *QAbstractListModel {
	var n = new(QAbstractListModel)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QAbstractListModel_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAbstractListModel) QAbstractListModel_PTR() *QAbstractListModel {
	return ptr
}

func (ptr *QAbstractListModel) Index(row int, column int, parent QModelIndex_ITF) *QModelIndex {
	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QAbstractListModel_Index(ptr.Pointer(), C.int(row), C.int(column), PointerFromQModelIndex(parent)))
	}
	return nil
}

func (ptr *QAbstractListModel) DropMimeData(data QMimeData_ITF, action Qt__DropAction, row int, column int, parent QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractListModel_DropMimeData(ptr.Pointer(), PointerFromQMimeData(data), C.int(action), C.int(row), C.int(column), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QAbstractListModel) Flags(index QModelIndex_ITF) Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return Qt__ItemFlag(C.QAbstractListModel_Flags(ptr.Pointer(), PointerFromQModelIndex(index)))
	}
	return 0
}

func (ptr *QAbstractListModel) Sibling(row int, column int, idx QModelIndex_ITF) *QModelIndex {
	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QAbstractListModel_Sibling(ptr.Pointer(), C.int(row), C.int(column), PointerFromQModelIndex(idx)))
	}
	return nil
}

func (ptr *QAbstractListModel) DestroyQAbstractListModel() {
	if ptr.Pointer() != nil {
		C.QAbstractListModel_DestroyQAbstractListModel(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
