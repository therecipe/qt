package core

//#include "qidentityproxymodel.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QIdentityProxyModel struct {
	QAbstractProxyModel
}

type QIdentityProxyModel_ITF interface {
	QAbstractProxyModel_ITF
	QIdentityProxyModel_PTR() *QIdentityProxyModel
}

func PointerFromQIdentityProxyModel(ptr QIdentityProxyModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QIdentityProxyModel_PTR().Pointer()
	}
	return nil
}

func NewQIdentityProxyModelFromPointer(ptr unsafe.Pointer) *QIdentityProxyModel {
	var n = new(QIdentityProxyModel)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QIdentityProxyModel_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QIdentityProxyModel) QIdentityProxyModel_PTR() *QIdentityProxyModel {
	return ptr
}

func NewQIdentityProxyModel(parent QObject_ITF) *QIdentityProxyModel {
	return NewQIdentityProxyModelFromPointer(C.QIdentityProxyModel_NewQIdentityProxyModel(PointerFromQObject(parent)))
}

func (ptr *QIdentityProxyModel) ColumnCount(parent QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(C.QIdentityProxyModel_ColumnCount(ptr.Pointer(), PointerFromQModelIndex(parent)))
	}
	return 0
}

func (ptr *QIdentityProxyModel) DropMimeData(data QMimeData_ITF, action Qt__DropAction, row int, column int, parent QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QIdentityProxyModel_DropMimeData(ptr.Pointer(), PointerFromQMimeData(data), C.int(action), C.int(row), C.int(column), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QIdentityProxyModel) HeaderData(section int, orientation Qt__Orientation, role int) *QVariant {
	if ptr.Pointer() != nil {
		return NewQVariantFromPointer(C.QIdentityProxyModel_HeaderData(ptr.Pointer(), C.int(section), C.int(orientation), C.int(role)))
	}
	return nil
}

func (ptr *QIdentityProxyModel) Index(row int, column int, parent QModelIndex_ITF) *QModelIndex {
	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QIdentityProxyModel_Index(ptr.Pointer(), C.int(row), C.int(column), PointerFromQModelIndex(parent)))
	}
	return nil
}

func (ptr *QIdentityProxyModel) InsertColumns(column int, count int, parent QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QIdentityProxyModel_InsertColumns(ptr.Pointer(), C.int(column), C.int(count), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QIdentityProxyModel) InsertRows(row int, count int, parent QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QIdentityProxyModel_InsertRows(ptr.Pointer(), C.int(row), C.int(count), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QIdentityProxyModel) MapFromSource(sourceIndex QModelIndex_ITF) *QModelIndex {
	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QIdentityProxyModel_MapFromSource(ptr.Pointer(), PointerFromQModelIndex(sourceIndex)))
	}
	return nil
}

func (ptr *QIdentityProxyModel) MapToSource(proxyIndex QModelIndex_ITF) *QModelIndex {
	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QIdentityProxyModel_MapToSource(ptr.Pointer(), PointerFromQModelIndex(proxyIndex)))
	}
	return nil
}

func (ptr *QIdentityProxyModel) Parent(child QModelIndex_ITF) *QModelIndex {
	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QIdentityProxyModel_Parent(ptr.Pointer(), PointerFromQModelIndex(child)))
	}
	return nil
}

func (ptr *QIdentityProxyModel) RemoveColumns(column int, count int, parent QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QIdentityProxyModel_RemoveColumns(ptr.Pointer(), C.int(column), C.int(count), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QIdentityProxyModel) RemoveRows(row int, count int, parent QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QIdentityProxyModel_RemoveRows(ptr.Pointer(), C.int(row), C.int(count), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QIdentityProxyModel) RowCount(parent QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(C.QIdentityProxyModel_RowCount(ptr.Pointer(), PointerFromQModelIndex(parent)))
	}
	return 0
}

func (ptr *QIdentityProxyModel) SetSourceModel(newSourceModel QAbstractItemModel_ITF) {
	if ptr.Pointer() != nil {
		C.QIdentityProxyModel_SetSourceModel(ptr.Pointer(), PointerFromQAbstractItemModel(newSourceModel))
	}
}

func (ptr *QIdentityProxyModel) Sibling(row int, column int, idx QModelIndex_ITF) *QModelIndex {
	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QIdentityProxyModel_Sibling(ptr.Pointer(), C.int(row), C.int(column), PointerFromQModelIndex(idx)))
	}
	return nil
}

func (ptr *QIdentityProxyModel) DestroyQIdentityProxyModel() {
	if ptr.Pointer() != nil {
		C.QIdentityProxyModel_DestroyQIdentityProxyModel(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
