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

type QAbstractListModelITF interface {
	QAbstractItemModelITF
	QAbstractListModelPTR() *QAbstractListModel
}

func PointerFromQAbstractListModel(ptr QAbstractListModelITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractListModelPTR().Pointer()
	}
	return nil
}

func QAbstractListModelFromPointer(ptr unsafe.Pointer) *QAbstractListModel {
	var n = new(QAbstractListModel)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QAbstractListModel_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAbstractListModel) QAbstractListModelPTR() *QAbstractListModel {
	return ptr
}

func (ptr *QAbstractListModel) Index(row int, column int, parent QModelIndexITF) *QModelIndex {
	if ptr.Pointer() != nil {
		return QModelIndexFromPointer(unsafe.Pointer(C.QAbstractListModel_Index(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(column), C.QtObjectPtr(PointerFromQModelIndex(parent)))))
	}
	return nil
}

func (ptr *QAbstractListModel) DropMimeData(data QMimeDataITF, action Qt__DropAction, row int, column int, parent QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractListModel_DropMimeData(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQMimeData(data)), C.int(action), C.int(row), C.int(column), C.QtObjectPtr(PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QAbstractListModel) Flags(index QModelIndexITF) Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return Qt__ItemFlag(C.QAbstractListModel_Flags(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(index))))
	}
	return 0
}

func (ptr *QAbstractListModel) Sibling(row int, column int, idx QModelIndexITF) *QModelIndex {
	if ptr.Pointer() != nil {
		return QModelIndexFromPointer(unsafe.Pointer(C.QAbstractListModel_Sibling(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(column), C.QtObjectPtr(PointerFromQModelIndex(idx)))))
	}
	return nil
}

func (ptr *QAbstractListModel) DestroyQAbstractListModel() {
	if ptr.Pointer() != nil {
		C.QAbstractListModel_DestroyQAbstractListModel(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
