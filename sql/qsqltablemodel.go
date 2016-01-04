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
func callbackQSqlTableModelBeforeDelete(ptr unsafe.Pointer, ptrName *C.char, row C.int) {
	defer qt.Recovering("callback QSqlTableModel::beforeDelete")

	if signal := qt.GetSignal(C.GoString(ptrName), "beforeDelete"); signal != nil {
		signal.(func(int))(int(row))
	}

}

func (ptr *QSqlTableModel) BeforeDelete(row int) {
	defer qt.Recovering("QSqlTableModel::beforeDelete")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_BeforeDelete(ptr.Pointer(), C.int(row))
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
func callbackQSqlTableModelClear(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSqlTableModel::clear")

	if signal := qt.GetSignal(C.GoString(ptrName), "clear"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlTableModelFromPointer(ptr).ClearDefault()
	}
}

func (ptr *QSqlTableModel) Clear() {
	defer qt.Recovering("QSqlTableModel::clear")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_Clear(ptr.Pointer())
	}
}

func (ptr *QSqlTableModel) ClearDefault() {
	defer qt.Recovering("QSqlTableModel::clear")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_ClearDefault(ptr.Pointer())
	}
}

func (ptr *QSqlTableModel) Data(index core.QModelIndex_ITF, role int) *core.QVariant {
	defer qt.Recovering("QSqlTableModel::data")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSqlTableModel_Data(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(role)))
	}
	return nil
}

func (ptr *QSqlTableModel) DeleteRowFromTable(row int) bool {
	defer qt.Recovering("QSqlTableModel::deleteRowFromTable")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_DeleteRowFromTable(ptr.Pointer(), C.int(row)) != 0
	}
	return false
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

func (ptr *QSqlTableModel) IndexInQuery(item core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QSqlTableModel::indexInQuery")

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QSqlTableModel_IndexInQuery(ptr.Pointer(), core.PointerFromQModelIndex(item)))
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

func (ptr *QSqlTableModel) InsertRowIntoTable(values QSqlRecord_ITF) bool {
	defer qt.Recovering("QSqlTableModel::insertRowIntoTable")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_InsertRowIntoTable(ptr.Pointer(), PointerFromQSqlRecord(values)) != 0
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

func (ptr *QSqlTableModel) OrderByClause() string {
	defer qt.Recovering("QSqlTableModel::orderByClause")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlTableModel_OrderByClause(ptr.Pointer()))
	}
	return ""
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
func callbackQSqlTableModelRevert(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QSqlTableModel::revert")

	if signal := qt.GetSignal(C.GoString(ptrName), "revert"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QSqlTableModel) Revert() {
	defer qt.Recovering("QSqlTableModel::revert")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_Revert(ptr.Pointer())
	}
}

func (ptr *QSqlTableModel) RevertDefault() {
	defer qt.Recovering("QSqlTableModel::revert")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_RevertDefault(ptr.Pointer())
	}
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
func callbackQSqlTableModelRevertRow(ptr unsafe.Pointer, ptrName *C.char, row C.int) {
	defer qt.Recovering("callback QSqlTableModel::revertRow")

	if signal := qt.GetSignal(C.GoString(ptrName), "revertRow"); signal != nil {
		signal.(func(int))(int(row))
	} else {
		NewQSqlTableModelFromPointer(ptr).RevertRowDefault(int(row))
	}
}

func (ptr *QSqlTableModel) RevertRow(row int) {
	defer qt.Recovering("QSqlTableModel::revertRow")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_RevertRow(ptr.Pointer(), C.int(row))
	}
}

func (ptr *QSqlTableModel) RevertRowDefault(row int) {
	defer qt.Recovering("QSqlTableModel::revertRow")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_RevertRowDefault(ptr.Pointer(), C.int(row))
	}
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

func (ptr *QSqlTableModel) SelectStatement() string {
	defer qt.Recovering("QSqlTableModel::selectStatement")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlTableModel_SelectStatement(ptr.Pointer()))
	}
	return ""
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
func callbackQSqlTableModelSetEditStrategy(ptr unsafe.Pointer, ptrName *C.char, strategy C.int) {
	defer qt.Recovering("callback QSqlTableModel::setEditStrategy")

	if signal := qt.GetSignal(C.GoString(ptrName), "setEditStrategy"); signal != nil {
		signal.(func(QSqlTableModel__EditStrategy))(QSqlTableModel__EditStrategy(strategy))
	} else {
		NewQSqlTableModelFromPointer(ptr).SetEditStrategyDefault(QSqlTableModel__EditStrategy(strategy))
	}
}

func (ptr *QSqlTableModel) SetEditStrategy(strategy QSqlTableModel__EditStrategy) {
	defer qt.Recovering("QSqlTableModel::setEditStrategy")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_SetEditStrategy(ptr.Pointer(), C.int(strategy))
	}
}

func (ptr *QSqlTableModel) SetEditStrategyDefault(strategy QSqlTableModel__EditStrategy) {
	defer qt.Recovering("QSqlTableModel::setEditStrategy")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_SetEditStrategyDefault(ptr.Pointer(), C.int(strategy))
	}
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
func callbackQSqlTableModelSetFilter(ptr unsafe.Pointer, ptrName *C.char, filter *C.char) {
	defer qt.Recovering("callback QSqlTableModel::setFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "setFilter"); signal != nil {
		signal.(func(string))(C.GoString(filter))
	} else {
		NewQSqlTableModelFromPointer(ptr).SetFilterDefault(C.GoString(filter))
	}
}

func (ptr *QSqlTableModel) SetFilter(filter string) {
	defer qt.Recovering("QSqlTableModel::setFilter")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_SetFilter(ptr.Pointer(), C.CString(filter))
	}
}

func (ptr *QSqlTableModel) SetFilterDefault(filter string) {
	defer qt.Recovering("QSqlTableModel::setFilter")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_SetFilterDefault(ptr.Pointer(), C.CString(filter))
	}
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
func callbackQSqlTableModelSetSort(ptr unsafe.Pointer, ptrName *C.char, column C.int, order C.int) {
	defer qt.Recovering("callback QSqlTableModel::setSort")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSort"); signal != nil {
		signal.(func(int, core.Qt__SortOrder))(int(column), core.Qt__SortOrder(order))
	} else {
		NewQSqlTableModelFromPointer(ptr).SetSortDefault(int(column), core.Qt__SortOrder(order))
	}
}

func (ptr *QSqlTableModel) SetSort(column int, order core.Qt__SortOrder) {
	defer qt.Recovering("QSqlTableModel::setSort")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_SetSort(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QSqlTableModel) SetSortDefault(column int, order core.Qt__SortOrder) {
	defer qt.Recovering("QSqlTableModel::setSort")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_SetSortDefault(ptr.Pointer(), C.int(column), C.int(order))
	}
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
func callbackQSqlTableModelSetTable(ptr unsafe.Pointer, ptrName *C.char, tableName *C.char) {
	defer qt.Recovering("callback QSqlTableModel::setTable")

	if signal := qt.GetSignal(C.GoString(ptrName), "setTable"); signal != nil {
		signal.(func(string))(C.GoString(tableName))
	} else {
		NewQSqlTableModelFromPointer(ptr).SetTableDefault(C.GoString(tableName))
	}
}

func (ptr *QSqlTableModel) SetTable(tableName string) {
	defer qt.Recovering("QSqlTableModel::setTable")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_SetTable(ptr.Pointer(), C.CString(tableName))
	}
}

func (ptr *QSqlTableModel) SetTableDefault(tableName string) {
	defer qt.Recovering("QSqlTableModel::setTable")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_SetTableDefault(ptr.Pointer(), C.CString(tableName))
	}
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
func callbackQSqlTableModelSort(ptr unsafe.Pointer, ptrName *C.char, column C.int, order C.int) {
	defer qt.Recovering("callback QSqlTableModel::sort")

	if signal := qt.GetSignal(C.GoString(ptrName), "sort"); signal != nil {
		signal.(func(int, core.Qt__SortOrder))(int(column), core.Qt__SortOrder(order))
	} else {
		NewQSqlTableModelFromPointer(ptr).SortDefault(int(column), core.Qt__SortOrder(order))
	}
}

func (ptr *QSqlTableModel) Sort(column int, order core.Qt__SortOrder) {
	defer qt.Recovering("QSqlTableModel::sort")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_Sort(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QSqlTableModel) SortDefault(column int, order core.Qt__SortOrder) {
	defer qt.Recovering("QSqlTableModel::sort")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_SortDefault(ptr.Pointer(), C.int(column), C.int(order))
	}
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

func (ptr *QSqlTableModel) UpdateRowInTable(row int, values QSqlRecord_ITF) bool {
	defer qt.Recovering("QSqlTableModel::updateRowInTable")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_UpdateRowInTable(ptr.Pointer(), C.int(row), PointerFromQSqlRecord(values)) != 0
	}
	return false
}

func (ptr *QSqlTableModel) DestroyQSqlTableModel() {
	defer qt.Recovering("QSqlTableModel::~QSqlTableModel")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_DestroyQSqlTableModel(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSqlTableModel) ConnectFetchMore(f func(parent *core.QModelIndex)) {
	defer qt.Recovering("connect QSqlTableModel::fetchMore")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "fetchMore", f)
	}
}

func (ptr *QSqlTableModel) DisconnectFetchMore() {
	defer qt.Recovering("disconnect QSqlTableModel::fetchMore")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "fetchMore")
	}
}

//export callbackQSqlTableModelFetchMore
func callbackQSqlTableModelFetchMore(ptr unsafe.Pointer, ptrName *C.char, parent unsafe.Pointer) {
	defer qt.Recovering("callback QSqlTableModel::fetchMore")

	if signal := qt.GetSignal(C.GoString(ptrName), "fetchMore"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(parent))
	} else {
		NewQSqlTableModelFromPointer(ptr).FetchMoreDefault(core.NewQModelIndexFromPointer(parent))
	}
}

func (ptr *QSqlTableModel) FetchMore(parent core.QModelIndex_ITF) {
	defer qt.Recovering("QSqlTableModel::fetchMore")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_FetchMore(ptr.Pointer(), core.PointerFromQModelIndex(parent))
	}
}

func (ptr *QSqlTableModel) FetchMoreDefault(parent core.QModelIndex_ITF) {
	defer qt.Recovering("QSqlTableModel::fetchMore")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_FetchMoreDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent))
	}
}

func (ptr *QSqlTableModel) ConnectQueryChange(f func()) {
	defer qt.Recovering("connect QSqlTableModel::queryChange")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "queryChange", f)
	}
}

func (ptr *QSqlTableModel) DisconnectQueryChange() {
	defer qt.Recovering("disconnect QSqlTableModel::queryChange")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "queryChange")
	}
}

//export callbackQSqlTableModelQueryChange
func callbackQSqlTableModelQueryChange(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSqlTableModel::queryChange")

	if signal := qt.GetSignal(C.GoString(ptrName), "queryChange"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlTableModelFromPointer(ptr).QueryChangeDefault()
	}
}

func (ptr *QSqlTableModel) QueryChange() {
	defer qt.Recovering("QSqlTableModel::queryChange")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_QueryChange(ptr.Pointer())
	}
}

func (ptr *QSqlTableModel) QueryChangeDefault() {
	defer qt.Recovering("QSqlTableModel::queryChange")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_QueryChangeDefault(ptr.Pointer())
	}
}

func (ptr *QSqlTableModel) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSqlTableModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSqlTableModel) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSqlTableModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSqlTableModelTimerEvent
func callbackQSqlTableModelTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSqlTableModel::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSqlTableModelFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSqlTableModel) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSqlTableModel::timerEvent")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSqlTableModel) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSqlTableModel::timerEvent")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSqlTableModel) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSqlTableModel::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSqlTableModel) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSqlTableModel::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSqlTableModelChildEvent
func callbackQSqlTableModelChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSqlTableModel::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSqlTableModelFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSqlTableModel) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSqlTableModel::childEvent")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSqlTableModel) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSqlTableModel::childEvent")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSqlTableModel) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSqlTableModel::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSqlTableModel) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSqlTableModel::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSqlTableModelCustomEvent
func callbackQSqlTableModelCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSqlTableModel::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSqlTableModelFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSqlTableModel) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSqlTableModel::customEvent")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSqlTableModel) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSqlTableModel::customEvent")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
