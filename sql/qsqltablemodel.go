package sql

//#include "sql.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QSqlTableModel struct {
	QSqlQueryModel
}

type QSqlTableModel_ITF interface {
	QSqlQueryModel_ITF
	QSqlTableModel_PTR() *QSqlTableModel
}

func PointerFromQSqlTableModel(ptr QSqlTableModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlTableModel_PTR().Pointer()
	}
	return nil
}

func NewQSqlTableModelFromPointer(ptr unsafe.Pointer) *QSqlTableModel {
	var n = new(QSqlTableModel)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSqlTableModel_") {
		n.SetObjectName("QSqlTableModel_" + qt.Identifier())
	}
	return n
}

func (ptr *QSqlTableModel) QSqlTableModel_PTR() *QSqlTableModel {
	return ptr
}

//QSqlTableModel::EditStrategy
type QSqlTableModel__EditStrategy int64

const (
	QSqlTableModel__OnFieldChange  = QSqlTableModel__EditStrategy(0)
	QSqlTableModel__OnRowChange    = QSqlTableModel__EditStrategy(1)
	QSqlTableModel__OnManualSubmit = QSqlTableModel__EditStrategy(2)
)

func (ptr *QSqlTableModel) ConnectBeforeDelete(f func(row int)) {
	defer qt.Recovering("connect QSqlTableModel::beforeDelete")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_ConnectBeforeDelete(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "beforeDelete", f)
	}
}

func (ptr *QSqlTableModel) DisconnectBeforeDelete() {
	defer qt.Recovering("disconnect QSqlTableModel::beforeDelete")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_DisconnectBeforeDelete(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "beforeDelete")
	}
}

//export callbackQSqlTableModelBeforeDelete
func callbackQSqlTableModelBeforeDelete(ptrName *C.char, row C.int) {
	defer qt.Recovering("callback QSqlTableModel::beforeDelete")

	var signal = qt.GetSignal(C.GoString(ptrName), "beforeDelete")
	if signal != nil {
		signal.(func(int))(int(row))
	}

}

func (ptr *QSqlTableModel) ConnectClear(f func()) {
	defer qt.Recovering("connect QSqlTableModel::clear")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "clear", f)
	}
}

func (ptr *QSqlTableModel) DisconnectClear() {
	defer qt.Recovering("disconnect QSqlTableModel::clear")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "clear")
	}
}

//export callbackQSqlTableModelClear
func callbackQSqlTableModelClear(ptrName *C.char) bool {
	defer qt.Recovering("callback QSqlTableModel::clear")

	var signal = qt.GetSignal(C.GoString(ptrName), "clear")
	if signal != nil {
		defer signal.(func())()
		return true
	}
	return false

}

func (ptr *QSqlTableModel) Data(index core.QModelIndex_ITF, role int) *core.QVariant {
	defer qt.Recovering("QSqlTableModel::data")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSqlTableModel_Data(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(role)))
	}
	return nil
}

func (ptr *QSqlTableModel) EditStrategy() QSqlTableModel__EditStrategy {
	defer qt.Recovering("QSqlTableModel::editStrategy")

	if ptr.Pointer() != nil {
		return QSqlTableModel__EditStrategy(C.QSqlTableModel_EditStrategy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSqlTableModel) FieldIndex(fieldName string) int {
	defer qt.Recovering("QSqlTableModel::fieldIndex")

	if ptr.Pointer() != nil {
		return int(C.QSqlTableModel_FieldIndex(ptr.Pointer(), C.CString(fieldName)))
	}
	return 0
}

func (ptr *QSqlTableModel) Filter() string {
	defer qt.Recovering("QSqlTableModel::filter")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlTableModel_Filter(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlTableModel) Flags(index core.QModelIndex_ITF) core.Qt__ItemFlag {
	defer qt.Recovering("QSqlTableModel::flags")

	if ptr.Pointer() != nil {
		return core.Qt__ItemFlag(C.QSqlTableModel_Flags(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return 0
}

func (ptr *QSqlTableModel) HeaderData(section int, orientation core.Qt__Orientation, role int) *core.QVariant {
	defer qt.Recovering("QSqlTableModel::headerData")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSqlTableModel_HeaderData(ptr.Pointer(), C.int(section), C.int(orientation), C.int(role)))
	}
	return nil
}

func (ptr *QSqlTableModel) InsertRecord(row int, record QSqlRecord_ITF) bool {
	defer qt.Recovering("QSqlTableModel::insertRecord")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_InsertRecord(ptr.Pointer(), C.int(row), PointerFromQSqlRecord(record)) != 0
	}
	return false
}

func (ptr *QSqlTableModel) InsertRows(row int, count int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlTableModel::insertRows")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_InsertRows(ptr.Pointer(), C.int(row), C.int(count), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSqlTableModel) IsDirty2() bool {
	defer qt.Recovering("QSqlTableModel::isDirty")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_IsDirty2(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlTableModel) IsDirty(index core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlTableModel::isDirty")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_IsDirty(ptr.Pointer(), core.PointerFromQModelIndex(index)) != 0
	}
	return false
}

func (ptr *QSqlTableModel) RemoveColumns(column int, count int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlTableModel::removeColumns")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_RemoveColumns(ptr.Pointer(), C.int(column), C.int(count), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSqlTableModel) RemoveRows(row int, count int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlTableModel::removeRows")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_RemoveRows(ptr.Pointer(), C.int(row), C.int(count), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSqlTableModel) ConnectRevert(f func()) {
	defer qt.Recovering("connect QSqlTableModel::revert")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "revert", f)
	}
}

func (ptr *QSqlTableModel) DisconnectRevert() {
	defer qt.Recovering("disconnect QSqlTableModel::revert")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "revert")
	}
}

//export callbackQSqlTableModelRevert
func callbackQSqlTableModelRevert(ptrName *C.char) bool {
	defer qt.Recovering("callback QSqlTableModel::revert")

	var signal = qt.GetSignal(C.GoString(ptrName), "revert")
	if signal != nil {
		defer signal.(func())()
		return true
	}
	return false

}

func (ptr *QSqlTableModel) RevertAll() {
	defer qt.Recovering("QSqlTableModel::revertAll")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_RevertAll(ptr.Pointer())
	}
}

func (ptr *QSqlTableModel) ConnectRevertRow(f func(row int)) {
	defer qt.Recovering("connect QSqlTableModel::revertRow")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "revertRow", f)
	}
}

func (ptr *QSqlTableModel) DisconnectRevertRow() {
	defer qt.Recovering("disconnect QSqlTableModel::revertRow")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "revertRow")
	}
}

//export callbackQSqlTableModelRevertRow
func callbackQSqlTableModelRevertRow(ptrName *C.char, row C.int) bool {
	defer qt.Recovering("callback QSqlTableModel::revertRow")

	var signal = qt.GetSignal(C.GoString(ptrName), "revertRow")
	if signal != nil {
		defer signal.(func(int))(int(row))
		return true
	}
	return false

}

func (ptr *QSqlTableModel) RowCount(parent core.QModelIndex_ITF) int {
	defer qt.Recovering("QSqlTableModel::rowCount")

	if ptr.Pointer() != nil {
		return int(C.QSqlTableModel_RowCount(ptr.Pointer(), core.PointerFromQModelIndex(parent)))
	}
	return 0
}

func (ptr *QSqlTableModel) Select() bool {
	defer qt.Recovering("QSqlTableModel::select")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_Select(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlTableModel) SelectRow(row int) bool {
	defer qt.Recovering("QSqlTableModel::selectRow")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_SelectRow(ptr.Pointer(), C.int(row)) != 0
	}
	return false
}

func (ptr *QSqlTableModel) SetData(index core.QModelIndex_ITF, value core.QVariant_ITF, role int) bool {
	defer qt.Recovering("QSqlTableModel::setData")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_SetData(ptr.Pointer(), core.PointerFromQModelIndex(index), core.PointerFromQVariant(value), C.int(role)) != 0
	}
	return false
}

func (ptr *QSqlTableModel) ConnectSetEditStrategy(f func(strategy QSqlTableModel__EditStrategy)) {
	defer qt.Recovering("connect QSqlTableModel::setEditStrategy")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setEditStrategy", f)
	}
}

func (ptr *QSqlTableModel) DisconnectSetEditStrategy() {
	defer qt.Recovering("disconnect QSqlTableModel::setEditStrategy")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setEditStrategy")
	}
}

//export callbackQSqlTableModelSetEditStrategy
func callbackQSqlTableModelSetEditStrategy(ptrName *C.char, strategy C.int) bool {
	defer qt.Recovering("callback QSqlTableModel::setEditStrategy")

	var signal = qt.GetSignal(C.GoString(ptrName), "setEditStrategy")
	if signal != nil {
		defer signal.(func(QSqlTableModel__EditStrategy))(QSqlTableModel__EditStrategy(strategy))
		return true
	}
	return false

}

func (ptr *QSqlTableModel) ConnectSetFilter(f func(filter string)) {
	defer qt.Recovering("connect QSqlTableModel::setFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setFilter", f)
	}
}

func (ptr *QSqlTableModel) DisconnectSetFilter() {
	defer qt.Recovering("disconnect QSqlTableModel::setFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setFilter")
	}
}

//export callbackQSqlTableModelSetFilter
func callbackQSqlTableModelSetFilter(ptrName *C.char, filter *C.char) bool {
	defer qt.Recovering("callback QSqlTableModel::setFilter")

	var signal = qt.GetSignal(C.GoString(ptrName), "setFilter")
	if signal != nil {
		defer signal.(func(string))(C.GoString(filter))
		return true
	}
	return false

}

func (ptr *QSqlTableModel) SetRecord(row int, values QSqlRecord_ITF) bool {
	defer qt.Recovering("QSqlTableModel::setRecord")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_SetRecord(ptr.Pointer(), C.int(row), PointerFromQSqlRecord(values)) != 0
	}
	return false
}

func (ptr *QSqlTableModel) ConnectSetSort(f func(column int, order core.Qt__SortOrder)) {
	defer qt.Recovering("connect QSqlTableModel::setSort")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setSort", f)
	}
}

func (ptr *QSqlTableModel) DisconnectSetSort() {
	defer qt.Recovering("disconnect QSqlTableModel::setSort")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setSort")
	}
}

//export callbackQSqlTableModelSetSort
func callbackQSqlTableModelSetSort(ptrName *C.char, column C.int, order C.int) bool {
	defer qt.Recovering("callback QSqlTableModel::setSort")

	var signal = qt.GetSignal(C.GoString(ptrName), "setSort")
	if signal != nil {
		defer signal.(func(int, core.Qt__SortOrder))(int(column), core.Qt__SortOrder(order))
		return true
	}
	return false

}

func (ptr *QSqlTableModel) ConnectSetTable(f func(tableName string)) {
	defer qt.Recovering("connect QSqlTableModel::setTable")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setTable", f)
	}
}

func (ptr *QSqlTableModel) DisconnectSetTable() {
	defer qt.Recovering("disconnect QSqlTableModel::setTable")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setTable")
	}
}

//export callbackQSqlTableModelSetTable
func callbackQSqlTableModelSetTable(ptrName *C.char, tableName *C.char) bool {
	defer qt.Recovering("callback QSqlTableModel::setTable")

	var signal = qt.GetSignal(C.GoString(ptrName), "setTable")
	if signal != nil {
		defer signal.(func(string))(C.GoString(tableName))
		return true
	}
	return false

}

func (ptr *QSqlTableModel) ConnectSort(f func(column int, order core.Qt__SortOrder)) {
	defer qt.Recovering("connect QSqlTableModel::sort")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "sort", f)
	}
}

func (ptr *QSqlTableModel) DisconnectSort() {
	defer qt.Recovering("disconnect QSqlTableModel::sort")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "sort")
	}
}

//export callbackQSqlTableModelSort
func callbackQSqlTableModelSort(ptrName *C.char, column C.int, order C.int) bool {
	defer qt.Recovering("callback QSqlTableModel::sort")

	var signal = qt.GetSignal(C.GoString(ptrName), "sort")
	if signal != nil {
		defer signal.(func(int, core.Qt__SortOrder))(int(column), core.Qt__SortOrder(order))
		return true
	}
	return false

}

func (ptr *QSqlTableModel) Submit() bool {
	defer qt.Recovering("QSqlTableModel::submit")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_Submit(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlTableModel) SubmitAll() bool {
	defer qt.Recovering("QSqlTableModel::submitAll")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_SubmitAll(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlTableModel) TableName() string {
	defer qt.Recovering("QSqlTableModel::tableName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlTableModel_TableName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlTableModel) DestroyQSqlTableModel() {
	defer qt.Recovering("QSqlTableModel::~QSqlTableModel")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_DestroyQSqlTableModel(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
