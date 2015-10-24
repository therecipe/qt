package core

//#include "qabstracttablemodel.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QAbstractTableModel struct {
	QAbstractItemModel
}

type QAbstractTableModelITF interface {
	QAbstractItemModelITF
	QAbstractTableModelPTR() *QAbstractTableModel
}

func PointerFromQAbstractTableModel(ptr QAbstractTableModelITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractTableModelPTR().Pointer()
	}
	return nil
}

func QAbstractTableModelFromPointer(ptr unsafe.Pointer) *QAbstractTableModel {
	var n = new(QAbstractTableModel)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QAbstractTableModel_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAbstractTableModel) QAbstractTableModelPTR() *QAbstractTableModel {
	return ptr
}

func (ptr *QAbstractTableModel) Index(row int, column int, parent QModelIndexITF) *QModelIndex {
	if ptr.Pointer() != nil {
		return QModelIndexFromPointer(unsafe.Pointer(C.QAbstractTableModel_Index(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(column), C.QtObjectPtr(PointerFromQModelIndex(parent)))))
	}
	return nil
}

func (ptr *QAbstractTableModel) DropMimeData(data QMimeDataITF, action Qt__DropAction, row int, column int, parent QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractTableModel_DropMimeData(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQMimeData(data)), C.int(action), C.int(row), C.int(column), C.QtObjectPtr(PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QAbstractTableModel) Flags(index QModelIndexITF) Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return Qt__ItemFlag(C.QAbstractTableModel_Flags(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(index))))
	}
	return 0
}

func (ptr *QAbstractTableModel) Sibling(row int, column int, idx QModelIndexITF) *QModelIndex {
	if ptr.Pointer() != nil {
		return QModelIndexFromPointer(unsafe.Pointer(C.QAbstractTableModel_Sibling(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(column), C.QtObjectPtr(PointerFromQModelIndex(idx)))))
	}
	return nil
}

func (ptr *QAbstractTableModel) DestroyQAbstractTableModel() {
	if ptr.Pointer() != nil {
		C.QAbstractTableModel_DestroyQAbstractTableModel(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
