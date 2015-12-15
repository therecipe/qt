package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"strings"
	"unsafe"
)

type QStringListModel struct {
	QAbstractListModel
}

type QStringListModel_ITF interface {
	QAbstractListModel_ITF
	QStringListModel_PTR() *QStringListModel
}

func PointerFromQStringListModel(ptr QStringListModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStringListModel_PTR().Pointer()
	}
	return nil
}

func NewQStringListModelFromPointer(ptr unsafe.Pointer) *QStringListModel {
	var n = new(QStringListModel)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QStringListModel_") {
		n.SetObjectName("QStringListModel_" + qt.Identifier())
	}
	return n
}

func (ptr *QStringListModel) QStringListModel_PTR() *QStringListModel {
	return ptr
}

func (ptr *QStringListModel) Data(index QModelIndex_ITF, role int) *QVariant {
	defer qt.Recovering("QStringListModel::data")

	if ptr.Pointer() != nil {
		return NewQVariantFromPointer(C.QStringListModel_Data(ptr.Pointer(), PointerFromQModelIndex(index), C.int(role)))
	}
	return nil
}

func (ptr *QStringListModel) Flags(index QModelIndex_ITF) Qt__ItemFlag {
	defer qt.Recovering("QStringListModel::flags")

	if ptr.Pointer() != nil {
		return Qt__ItemFlag(C.QStringListModel_Flags(ptr.Pointer(), PointerFromQModelIndex(index)))
	}
	return 0
}

func (ptr *QStringListModel) InsertRows(row int, count int, parent QModelIndex_ITF) bool {
	defer qt.Recovering("QStringListModel::insertRows")

	if ptr.Pointer() != nil {
		return C.QStringListModel_InsertRows(ptr.Pointer(), C.int(row), C.int(count), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QStringListModel) RemoveRows(row int, count int, parent QModelIndex_ITF) bool {
	defer qt.Recovering("QStringListModel::removeRows")

	if ptr.Pointer() != nil {
		return C.QStringListModel_RemoveRows(ptr.Pointer(), C.int(row), C.int(count), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QStringListModel) RowCount(parent QModelIndex_ITF) int {
	defer qt.Recovering("QStringListModel::rowCount")

	if ptr.Pointer() != nil {
		return int(C.QStringListModel_RowCount(ptr.Pointer(), PointerFromQModelIndex(parent)))
	}
	return 0
}

func (ptr *QStringListModel) SetData(index QModelIndex_ITF, value QVariant_ITF, role int) bool {
	defer qt.Recovering("QStringListModel::setData")

	if ptr.Pointer() != nil {
		return C.QStringListModel_SetData(ptr.Pointer(), PointerFromQModelIndex(index), PointerFromQVariant(value), C.int(role)) != 0
	}
	return false
}

func (ptr *QStringListModel) SetStringList(strin []string) {
	defer qt.Recovering("QStringListModel::setStringList")

	if ptr.Pointer() != nil {
		C.QStringListModel_SetStringList(ptr.Pointer(), C.CString(strings.Join(strin, ",,,")))
	}
}

func (ptr *QStringListModel) Sibling(row int, column int, idx QModelIndex_ITF) *QModelIndex {
	defer qt.Recovering("QStringListModel::sibling")

	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QStringListModel_Sibling(ptr.Pointer(), C.int(row), C.int(column), PointerFromQModelIndex(idx)))
	}
	return nil
}

func (ptr *QStringListModel) ConnectSort(f func(column int, order Qt__SortOrder)) {
	defer qt.Recovering("connect QStringListModel::sort")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "sort", f)
	}
}

func (ptr *QStringListModel) DisconnectSort() {
	defer qt.Recovering("disconnect QStringListModel::sort")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "sort")
	}
}

//export callbackQStringListModelSort
func callbackQStringListModelSort(ptrName *C.char, column C.int, order C.int) bool {
	defer qt.Recovering("callback QStringListModel::sort")

	var signal = qt.GetSignal(C.GoString(ptrName), "sort")
	if signal != nil {
		defer signal.(func(int, Qt__SortOrder))(int(column), Qt__SortOrder(order))
		return true
	}
	return false

}

func (ptr *QStringListModel) StringList() []string {
	defer qt.Recovering("QStringListModel::stringList")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QStringListModel_StringList(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QStringListModel) SupportedDropActions() Qt__DropAction {
	defer qt.Recovering("QStringListModel::supportedDropActions")

	if ptr.Pointer() != nil {
		return Qt__DropAction(C.QStringListModel_SupportedDropActions(ptr.Pointer()))
	}
	return 0
}
