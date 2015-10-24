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

type QIdentityProxyModelITF interface {
	QAbstractProxyModelITF
	QIdentityProxyModelPTR() *QIdentityProxyModel
}

func PointerFromQIdentityProxyModel(ptr QIdentityProxyModelITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QIdentityProxyModelPTR().Pointer()
	}
	return nil
}

func QIdentityProxyModelFromPointer(ptr unsafe.Pointer) *QIdentityProxyModel {
	var n = new(QIdentityProxyModel)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QIdentityProxyModel_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QIdentityProxyModel) QIdentityProxyModelPTR() *QIdentityProxyModel {
	return ptr
}

func NewQIdentityProxyModel(parent QObjectITF) *QIdentityProxyModel {
	return QIdentityProxyModelFromPointer(unsafe.Pointer(C.QIdentityProxyModel_NewQIdentityProxyModel(C.QtObjectPtr(PointerFromQObject(parent)))))
}

func (ptr *QIdentityProxyModel) ColumnCount(parent QModelIndexITF) int {
	if ptr.Pointer() != nil {
		return int(C.QIdentityProxyModel_ColumnCount(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(parent))))
	}
	return 0
}

func (ptr *QIdentityProxyModel) DropMimeData(data QMimeDataITF, action Qt__DropAction, row int, column int, parent QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QIdentityProxyModel_DropMimeData(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQMimeData(data)), C.int(action), C.int(row), C.int(column), C.QtObjectPtr(PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QIdentityProxyModel) HeaderData(section int, orientation Qt__Orientation, role int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QIdentityProxyModel_HeaderData(C.QtObjectPtr(ptr.Pointer()), C.int(section), C.int(orientation), C.int(role)))
	}
	return ""
}

func (ptr *QIdentityProxyModel) Index(row int, column int, parent QModelIndexITF) *QModelIndex {
	if ptr.Pointer() != nil {
		return QModelIndexFromPointer(unsafe.Pointer(C.QIdentityProxyModel_Index(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(column), C.QtObjectPtr(PointerFromQModelIndex(parent)))))
	}
	return nil
}

func (ptr *QIdentityProxyModel) InsertColumns(column int, count int, parent QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QIdentityProxyModel_InsertColumns(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.int(count), C.QtObjectPtr(PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QIdentityProxyModel) InsertRows(row int, count int, parent QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QIdentityProxyModel_InsertRows(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(count), C.QtObjectPtr(PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QIdentityProxyModel) MapFromSource(sourceIndex QModelIndexITF) *QModelIndex {
	if ptr.Pointer() != nil {
		return QModelIndexFromPointer(unsafe.Pointer(C.QIdentityProxyModel_MapFromSource(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(sourceIndex)))))
	}
	return nil
}

func (ptr *QIdentityProxyModel) MapToSource(proxyIndex QModelIndexITF) *QModelIndex {
	if ptr.Pointer() != nil {
		return QModelIndexFromPointer(unsafe.Pointer(C.QIdentityProxyModel_MapToSource(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(proxyIndex)))))
	}
	return nil
}

func (ptr *QIdentityProxyModel) Parent(child QModelIndexITF) *QModelIndex {
	if ptr.Pointer() != nil {
		return QModelIndexFromPointer(unsafe.Pointer(C.QIdentityProxyModel_Parent(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(child)))))
	}
	return nil
}

func (ptr *QIdentityProxyModel) RemoveColumns(column int, count int, parent QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QIdentityProxyModel_RemoveColumns(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.int(count), C.QtObjectPtr(PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QIdentityProxyModel) RemoveRows(row int, count int, parent QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QIdentityProxyModel_RemoveRows(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(count), C.QtObjectPtr(PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QIdentityProxyModel) RowCount(parent QModelIndexITF) int {
	if ptr.Pointer() != nil {
		return int(C.QIdentityProxyModel_RowCount(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(parent))))
	}
	return 0
}

func (ptr *QIdentityProxyModel) SetSourceModel(newSourceModel QAbstractItemModelITF) {
	if ptr.Pointer() != nil {
		C.QIdentityProxyModel_SetSourceModel(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAbstractItemModel(newSourceModel)))
	}
}

func (ptr *QIdentityProxyModel) Sibling(row int, column int, idx QModelIndexITF) *QModelIndex {
	if ptr.Pointer() != nil {
		return QModelIndexFromPointer(unsafe.Pointer(C.QIdentityProxyModel_Sibling(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(column), C.QtObjectPtr(PointerFromQModelIndex(idx)))))
	}
	return nil
}

func (ptr *QIdentityProxyModel) DestroyQIdentityProxyModel() {
	if ptr.Pointer() != nil {
		C.QIdentityProxyModel_DestroyQIdentityProxyModel(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
