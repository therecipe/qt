package core

//#include "core.h"
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
	for len(n.ObjectName()) < len("QIdentityProxyModel_") {
		n.SetObjectName("QIdentityProxyModel_" + qt.Identifier())
	}
	return n
}

func (ptr *QIdentityProxyModel) QIdentityProxyModel_PTR() *QIdentityProxyModel {
	return ptr
}

func NewQIdentityProxyModel(parent QObject_ITF) *QIdentityProxyModel {
	defer qt.Recovering("QIdentityProxyModel::QIdentityProxyModel")

	return NewQIdentityProxyModelFromPointer(C.QIdentityProxyModel_NewQIdentityProxyModel(PointerFromQObject(parent)))
}

func (ptr *QIdentityProxyModel) ColumnCount(parent QModelIndex_ITF) int {
	defer qt.Recovering("QIdentityProxyModel::columnCount")

	if ptr.Pointer() != nil {
		return int(C.QIdentityProxyModel_ColumnCount(ptr.Pointer(), PointerFromQModelIndex(parent)))
	}
	return 0
}

func (ptr *QIdentityProxyModel) DropMimeData(data QMimeData_ITF, action Qt__DropAction, row int, column int, parent QModelIndex_ITF) bool {
	defer qt.Recovering("QIdentityProxyModel::dropMimeData")

	if ptr.Pointer() != nil {
		return C.QIdentityProxyModel_DropMimeData(ptr.Pointer(), PointerFromQMimeData(data), C.int(action), C.int(row), C.int(column), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QIdentityProxyModel) HeaderData(section int, orientation Qt__Orientation, role int) *QVariant {
	defer qt.Recovering("QIdentityProxyModel::headerData")

	if ptr.Pointer() != nil {
		return NewQVariantFromPointer(C.QIdentityProxyModel_HeaderData(ptr.Pointer(), C.int(section), C.int(orientation), C.int(role)))
	}
	return nil
}

func (ptr *QIdentityProxyModel) Index(row int, column int, parent QModelIndex_ITF) *QModelIndex {
	defer qt.Recovering("QIdentityProxyModel::index")

	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QIdentityProxyModel_Index(ptr.Pointer(), C.int(row), C.int(column), PointerFromQModelIndex(parent)))
	}
	return nil
}

func (ptr *QIdentityProxyModel) InsertColumns(column int, count int, parent QModelIndex_ITF) bool {
	defer qt.Recovering("QIdentityProxyModel::insertColumns")

	if ptr.Pointer() != nil {
		return C.QIdentityProxyModel_InsertColumns(ptr.Pointer(), C.int(column), C.int(count), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QIdentityProxyModel) InsertRows(row int, count int, parent QModelIndex_ITF) bool {
	defer qt.Recovering("QIdentityProxyModel::insertRows")

	if ptr.Pointer() != nil {
		return C.QIdentityProxyModel_InsertRows(ptr.Pointer(), C.int(row), C.int(count), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QIdentityProxyModel) MapFromSource(sourceIndex QModelIndex_ITF) *QModelIndex {
	defer qt.Recovering("QIdentityProxyModel::mapFromSource")

	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QIdentityProxyModel_MapFromSource(ptr.Pointer(), PointerFromQModelIndex(sourceIndex)))
	}
	return nil
}

func (ptr *QIdentityProxyModel) MapToSource(proxyIndex QModelIndex_ITF) *QModelIndex {
	defer qt.Recovering("QIdentityProxyModel::mapToSource")

	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QIdentityProxyModel_MapToSource(ptr.Pointer(), PointerFromQModelIndex(proxyIndex)))
	}
	return nil
}

func (ptr *QIdentityProxyModel) Parent(child QModelIndex_ITF) *QModelIndex {
	defer qt.Recovering("QIdentityProxyModel::parent")

	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QIdentityProxyModel_Parent(ptr.Pointer(), PointerFromQModelIndex(child)))
	}
	return nil
}

func (ptr *QIdentityProxyModel) RemoveColumns(column int, count int, parent QModelIndex_ITF) bool {
	defer qt.Recovering("QIdentityProxyModel::removeColumns")

	if ptr.Pointer() != nil {
		return C.QIdentityProxyModel_RemoveColumns(ptr.Pointer(), C.int(column), C.int(count), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QIdentityProxyModel) RemoveRows(row int, count int, parent QModelIndex_ITF) bool {
	defer qt.Recovering("QIdentityProxyModel::removeRows")

	if ptr.Pointer() != nil {
		return C.QIdentityProxyModel_RemoveRows(ptr.Pointer(), C.int(row), C.int(count), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QIdentityProxyModel) RowCount(parent QModelIndex_ITF) int {
	defer qt.Recovering("QIdentityProxyModel::rowCount")

	if ptr.Pointer() != nil {
		return int(C.QIdentityProxyModel_RowCount(ptr.Pointer(), PointerFromQModelIndex(parent)))
	}
	return 0
}

func (ptr *QIdentityProxyModel) ConnectSetSourceModel(f func(newSourceModel *QAbstractItemModel)) {
	defer qt.Recovering("connect QIdentityProxyModel::setSourceModel")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setSourceModel", f)
	}
}

func (ptr *QIdentityProxyModel) DisconnectSetSourceModel() {
	defer qt.Recovering("disconnect QIdentityProxyModel::setSourceModel")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setSourceModel")
	}
}

//export callbackQIdentityProxyModelSetSourceModel
func callbackQIdentityProxyModelSetSourceModel(ptrName *C.char, newSourceModel unsafe.Pointer) bool {
	defer qt.Recovering("callback QIdentityProxyModel::setSourceModel")

	var signal = qt.GetSignal(C.GoString(ptrName), "setSourceModel")
	if signal != nil {
		defer signal.(func(*QAbstractItemModel))(NewQAbstractItemModelFromPointer(newSourceModel))
		return true
	}
	return false

}

func (ptr *QIdentityProxyModel) Sibling(row int, column int, idx QModelIndex_ITF) *QModelIndex {
	defer qt.Recovering("QIdentityProxyModel::sibling")

	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QIdentityProxyModel_Sibling(ptr.Pointer(), C.int(row), C.int(column), PointerFromQModelIndex(idx)))
	}
	return nil
}

func (ptr *QIdentityProxyModel) DestroyQIdentityProxyModel() {
	defer qt.Recovering("QIdentityProxyModel::~QIdentityProxyModel")

	if ptr.Pointer() != nil {
		C.QIdentityProxyModel_DestroyQIdentityProxyModel(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
