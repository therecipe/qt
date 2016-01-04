package sql

//#include "sql.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QSqlRelationalTableModel struct {
	QSqlTableModel
}

type QSqlRelationalTableModel_ITF interface {
	QSqlTableModel_ITF
	QSqlRelationalTableModel_PTR() *QSqlRelationalTableModel
}

func PointerFromQSqlRelationalTableModel(ptr QSqlRelationalTableModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlRelationalTableModel_PTR().Pointer()
	}
	return nil
}

func NewQSqlRelationalTableModelFromPointer(ptr unsafe.Pointer) *QSqlRelationalTableModel {
	var n = new(QSqlRelationalTableModel)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSqlRelationalTableModel_") {
		n.SetObjectName("QSqlRelationalTableModel_" + qt.Identifier())
	}
	return n
}

func (ptr *QSqlRelationalTableModel) QSqlRelationalTableModel_PTR() *QSqlRelationalTableModel {
	return ptr
}

//QSqlRelationalTableModel::JoinMode
type QSqlRelationalTableModel__JoinMode int64

const (
	QSqlRelationalTableModel__InnerJoin = QSqlRelationalTableModel__JoinMode(0)
	QSqlRelationalTableModel__LeftJoin  = QSqlRelationalTableModel__JoinMode(1)
)

func (ptr *QSqlRelationalTableModel) ConnectClear(f func()) {
	defer qt.Recovering("connect QSqlRelationalTableModel::clear")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "clear", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectClear() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::clear")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "clear")
	}
}

//export callbackQSqlRelationalTableModelClear
func callbackQSqlRelationalTableModelClear(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSqlRelationalTableModel::clear")

	if signal := qt.GetSignal(C.GoString(ptrName), "clear"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlRelationalTableModelFromPointer(ptr).ClearDefault()
	}
}

func (ptr *QSqlRelationalTableModel) Clear() {
	defer qt.Recovering("QSqlRelationalTableModel::clear")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_Clear(ptr.Pointer())
	}
}

func (ptr *QSqlRelationalTableModel) ClearDefault() {
	defer qt.Recovering("QSqlRelationalTableModel::clear")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_ClearDefault(ptr.Pointer())
	}
}

func (ptr *QSqlRelationalTableModel) Data(index core.QModelIndex_ITF, role int) *core.QVariant {
	defer qt.Recovering("QSqlRelationalTableModel::data")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSqlRelationalTableModel_Data(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(role)))
	}
	return nil
}

func (ptr *QSqlRelationalTableModel) InsertRowIntoTable(values QSqlRecord_ITF) bool {
	defer qt.Recovering("QSqlRelationalTableModel::insertRowIntoTable")

	if ptr.Pointer() != nil {
		return C.QSqlRelationalTableModel_InsertRowIntoTable(ptr.Pointer(), PointerFromQSqlRecord(values)) != 0
	}
	return false
}

func (ptr *QSqlRelationalTableModel) OrderByClause() string {
	defer qt.Recovering("QSqlRelationalTableModel::orderByClause")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlRelationalTableModel_OrderByClause(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlRelationalTableModel) RelationModel(column int) *QSqlTableModel {
	defer qt.Recovering("QSqlRelationalTableModel::relationModel")

	if ptr.Pointer() != nil {
		return NewQSqlTableModelFromPointer(C.QSqlRelationalTableModel_RelationModel(ptr.Pointer(), C.int(column)))
	}
	return nil
}

func (ptr *QSqlRelationalTableModel) RemoveColumns(column int, count int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlRelationalTableModel::removeColumns")

	if ptr.Pointer() != nil {
		return C.QSqlRelationalTableModel_RemoveColumns(ptr.Pointer(), C.int(column), C.int(count), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSqlRelationalTableModel) ConnectRevertRow(f func(row int)) {
	defer qt.Recovering("connect QSqlRelationalTableModel::revertRow")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "revertRow", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectRevertRow() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::revertRow")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "revertRow")
	}
}

//export callbackQSqlRelationalTableModelRevertRow
func callbackQSqlRelationalTableModelRevertRow(ptr unsafe.Pointer, ptrName *C.char, row C.int) bool {
	defer qt.Recovering("callback QSqlRelationalTableModel::revertRow")

	if signal := qt.GetSignal(C.GoString(ptrName), "revertRow"); signal != nil {
		signal.(func(int))(int(row))
		return true
	}
	return false

}

func (ptr *QSqlRelationalTableModel) RevertRow(row int) {
	defer qt.Recovering("QSqlRelationalTableModel::revertRow")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_RevertRow(ptr.Pointer(), C.int(row))
	}
}

func (ptr *QSqlRelationalTableModel) RevertRowDefault(row int) {
	defer qt.Recovering("QSqlRelationalTableModel::revertRow")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_RevertRowDefault(ptr.Pointer(), C.int(row))
	}
}

func (ptr *QSqlRelationalTableModel) Select() bool {
	defer qt.Recovering("QSqlRelationalTableModel::select")

	if ptr.Pointer() != nil {
		return C.QSqlRelationalTableModel_Select(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlRelationalTableModel) SelectStatement() string {
	defer qt.Recovering("QSqlRelationalTableModel::selectStatement")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlRelationalTableModel_SelectStatement(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlRelationalTableModel) SetData(index core.QModelIndex_ITF, value core.QVariant_ITF, role int) bool {
	defer qt.Recovering("QSqlRelationalTableModel::setData")

	if ptr.Pointer() != nil {
		return C.QSqlRelationalTableModel_SetData(ptr.Pointer(), core.PointerFromQModelIndex(index), core.PointerFromQVariant(value), C.int(role)) != 0
	}
	return false
}

func (ptr *QSqlRelationalTableModel) SetJoinMode(joinMode QSqlRelationalTableModel__JoinMode) {
	defer qt.Recovering("QSqlRelationalTableModel::setJoinMode")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_SetJoinMode(ptr.Pointer(), C.int(joinMode))
	}
}

func (ptr *QSqlRelationalTableModel) ConnectSetTable(f func(table string)) {
	defer qt.Recovering("connect QSqlRelationalTableModel::setTable")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setTable", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectSetTable() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::setTable")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setTable")
	}
}

//export callbackQSqlRelationalTableModelSetTable
func callbackQSqlRelationalTableModelSetTable(ptr unsafe.Pointer, ptrName *C.char, table *C.char) {
	defer qt.Recovering("callback QSqlRelationalTableModel::setTable")

	if signal := qt.GetSignal(C.GoString(ptrName), "setTable"); signal != nil {
		signal.(func(string))(C.GoString(table))
	} else {
		NewQSqlRelationalTableModelFromPointer(ptr).SetTableDefault(C.GoString(table))
	}
}

func (ptr *QSqlRelationalTableModel) SetTable(table string) {
	defer qt.Recovering("QSqlRelationalTableModel::setTable")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_SetTable(ptr.Pointer(), C.CString(table))
	}
}

func (ptr *QSqlRelationalTableModel) SetTableDefault(table string) {
	defer qt.Recovering("QSqlRelationalTableModel::setTable")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_SetTableDefault(ptr.Pointer(), C.CString(table))
	}
}

func (ptr *QSqlRelationalTableModel) UpdateRowInTable(row int, values QSqlRecord_ITF) bool {
	defer qt.Recovering("QSqlRelationalTableModel::updateRowInTable")

	if ptr.Pointer() != nil {
		return C.QSqlRelationalTableModel_UpdateRowInTable(ptr.Pointer(), C.int(row), PointerFromQSqlRecord(values)) != 0
	}
	return false
}

func (ptr *QSqlRelationalTableModel) DestroyQSqlRelationalTableModel() {
	defer qt.Recovering("QSqlRelationalTableModel::~QSqlRelationalTableModel")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_DestroyQSqlRelationalTableModel(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSqlRelationalTableModel) ConnectRevert(f func()) {
	defer qt.Recovering("connect QSqlRelationalTableModel::revert")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "revert", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectRevert() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::revert")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "revert")
	}
}

//export callbackQSqlRelationalTableModelRevert
func callbackQSqlRelationalTableModelRevert(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QSqlRelationalTableModel::revert")

	if signal := qt.GetSignal(C.GoString(ptrName), "revert"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QSqlRelationalTableModel) Revert() {
	defer qt.Recovering("QSqlRelationalTableModel::revert")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_Revert(ptr.Pointer())
	}
}

func (ptr *QSqlRelationalTableModel) RevertDefault() {
	defer qt.Recovering("QSqlRelationalTableModel::revert")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_RevertDefault(ptr.Pointer())
	}
}

func (ptr *QSqlRelationalTableModel) ConnectSetEditStrategy(f func(strategy QSqlTableModel__EditStrategy)) {
	defer qt.Recovering("connect QSqlRelationalTableModel::setEditStrategy")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setEditStrategy", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectSetEditStrategy() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::setEditStrategy")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setEditStrategy")
	}
}

//export callbackQSqlRelationalTableModelSetEditStrategy
func callbackQSqlRelationalTableModelSetEditStrategy(ptr unsafe.Pointer, ptrName *C.char, strategy C.int) {
	defer qt.Recovering("callback QSqlRelationalTableModel::setEditStrategy")

	if signal := qt.GetSignal(C.GoString(ptrName), "setEditStrategy"); signal != nil {
		signal.(func(QSqlTableModel__EditStrategy))(QSqlTableModel__EditStrategy(strategy))
	} else {
		NewQSqlRelationalTableModelFromPointer(ptr).SetEditStrategyDefault(QSqlTableModel__EditStrategy(strategy))
	}
}

func (ptr *QSqlRelationalTableModel) SetEditStrategy(strategy QSqlTableModel__EditStrategy) {
	defer qt.Recovering("QSqlRelationalTableModel::setEditStrategy")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_SetEditStrategy(ptr.Pointer(), C.int(strategy))
	}
}

func (ptr *QSqlRelationalTableModel) SetEditStrategyDefault(strategy QSqlTableModel__EditStrategy) {
	defer qt.Recovering("QSqlRelationalTableModel::setEditStrategy")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_SetEditStrategyDefault(ptr.Pointer(), C.int(strategy))
	}
}

func (ptr *QSqlRelationalTableModel) ConnectSetFilter(f func(filter string)) {
	defer qt.Recovering("connect QSqlRelationalTableModel::setFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setFilter", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectSetFilter() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::setFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setFilter")
	}
}

//export callbackQSqlRelationalTableModelSetFilter
func callbackQSqlRelationalTableModelSetFilter(ptr unsafe.Pointer, ptrName *C.char, filter *C.char) {
	defer qt.Recovering("callback QSqlRelationalTableModel::setFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "setFilter"); signal != nil {
		signal.(func(string))(C.GoString(filter))
	} else {
		NewQSqlRelationalTableModelFromPointer(ptr).SetFilterDefault(C.GoString(filter))
	}
}

func (ptr *QSqlRelationalTableModel) SetFilter(filter string) {
	defer qt.Recovering("QSqlRelationalTableModel::setFilter")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_SetFilter(ptr.Pointer(), C.CString(filter))
	}
}

func (ptr *QSqlRelationalTableModel) SetFilterDefault(filter string) {
	defer qt.Recovering("QSqlRelationalTableModel::setFilter")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_SetFilterDefault(ptr.Pointer(), C.CString(filter))
	}
}

func (ptr *QSqlRelationalTableModel) ConnectSetSort(f func(column int, order core.Qt__SortOrder)) {
	defer qt.Recovering("connect QSqlRelationalTableModel::setSort")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setSort", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectSetSort() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::setSort")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setSort")
	}
}

//export callbackQSqlRelationalTableModelSetSort
func callbackQSqlRelationalTableModelSetSort(ptr unsafe.Pointer, ptrName *C.char, column C.int, order C.int) {
	defer qt.Recovering("callback QSqlRelationalTableModel::setSort")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSort"); signal != nil {
		signal.(func(int, core.Qt__SortOrder))(int(column), core.Qt__SortOrder(order))
	} else {
		NewQSqlRelationalTableModelFromPointer(ptr).SetSortDefault(int(column), core.Qt__SortOrder(order))
	}
}

func (ptr *QSqlRelationalTableModel) SetSort(column int, order core.Qt__SortOrder) {
	defer qt.Recovering("QSqlRelationalTableModel::setSort")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_SetSort(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QSqlRelationalTableModel) SetSortDefault(column int, order core.Qt__SortOrder) {
	defer qt.Recovering("QSqlRelationalTableModel::setSort")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_SetSortDefault(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QSqlRelationalTableModel) ConnectSort(f func(column int, order core.Qt__SortOrder)) {
	defer qt.Recovering("connect QSqlRelationalTableModel::sort")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "sort", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectSort() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::sort")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "sort")
	}
}

//export callbackQSqlRelationalTableModelSort
func callbackQSqlRelationalTableModelSort(ptr unsafe.Pointer, ptrName *C.char, column C.int, order C.int) {
	defer qt.Recovering("callback QSqlRelationalTableModel::sort")

	if signal := qt.GetSignal(C.GoString(ptrName), "sort"); signal != nil {
		signal.(func(int, core.Qt__SortOrder))(int(column), core.Qt__SortOrder(order))
	} else {
		NewQSqlRelationalTableModelFromPointer(ptr).SortDefault(int(column), core.Qt__SortOrder(order))
	}
}

func (ptr *QSqlRelationalTableModel) Sort(column int, order core.Qt__SortOrder) {
	defer qt.Recovering("QSqlRelationalTableModel::sort")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_Sort(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QSqlRelationalTableModel) SortDefault(column int, order core.Qt__SortOrder) {
	defer qt.Recovering("QSqlRelationalTableModel::sort")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_SortDefault(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QSqlRelationalTableModel) ConnectFetchMore(f func(parent *core.QModelIndex)) {
	defer qt.Recovering("connect QSqlRelationalTableModel::fetchMore")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "fetchMore", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectFetchMore() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::fetchMore")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "fetchMore")
	}
}

//export callbackQSqlRelationalTableModelFetchMore
func callbackQSqlRelationalTableModelFetchMore(ptr unsafe.Pointer, ptrName *C.char, parent unsafe.Pointer) {
	defer qt.Recovering("callback QSqlRelationalTableModel::fetchMore")

	if signal := qt.GetSignal(C.GoString(ptrName), "fetchMore"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(parent))
	} else {
		NewQSqlRelationalTableModelFromPointer(ptr).FetchMoreDefault(core.NewQModelIndexFromPointer(parent))
	}
}

func (ptr *QSqlRelationalTableModel) FetchMore(parent core.QModelIndex_ITF) {
	defer qt.Recovering("QSqlRelationalTableModel::fetchMore")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_FetchMore(ptr.Pointer(), core.PointerFromQModelIndex(parent))
	}
}

func (ptr *QSqlRelationalTableModel) FetchMoreDefault(parent core.QModelIndex_ITF) {
	defer qt.Recovering("QSqlRelationalTableModel::fetchMore")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_FetchMoreDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent))
	}
}

func (ptr *QSqlRelationalTableModel) ConnectQueryChange(f func()) {
	defer qt.Recovering("connect QSqlRelationalTableModel::queryChange")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "queryChange", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectQueryChange() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::queryChange")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "queryChange")
	}
}

//export callbackQSqlRelationalTableModelQueryChange
func callbackQSqlRelationalTableModelQueryChange(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSqlRelationalTableModel::queryChange")

	if signal := qt.GetSignal(C.GoString(ptrName), "queryChange"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlRelationalTableModelFromPointer(ptr).QueryChangeDefault()
	}
}

func (ptr *QSqlRelationalTableModel) QueryChange() {
	defer qt.Recovering("QSqlRelationalTableModel::queryChange")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_QueryChange(ptr.Pointer())
	}
}

func (ptr *QSqlRelationalTableModel) QueryChangeDefault() {
	defer qt.Recovering("QSqlRelationalTableModel::queryChange")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_QueryChangeDefault(ptr.Pointer())
	}
}

func (ptr *QSqlRelationalTableModel) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSqlRelationalTableModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSqlRelationalTableModelTimerEvent
func callbackQSqlRelationalTableModelTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSqlRelationalTableModel::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSqlRelationalTableModelFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSqlRelationalTableModel) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSqlRelationalTableModel::timerEvent")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSqlRelationalTableModel) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSqlRelationalTableModel::timerEvent")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSqlRelationalTableModel) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSqlRelationalTableModel::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSqlRelationalTableModelChildEvent
func callbackQSqlRelationalTableModelChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSqlRelationalTableModel::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSqlRelationalTableModelFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSqlRelationalTableModel) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSqlRelationalTableModel::childEvent")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSqlRelationalTableModel) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSqlRelationalTableModel::childEvent")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSqlRelationalTableModel) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSqlRelationalTableModel::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSqlRelationalTableModelCustomEvent
func callbackQSqlRelationalTableModelCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSqlRelationalTableModel::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSqlRelationalTableModelFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSqlRelationalTableModel) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSqlRelationalTableModel::customEvent")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSqlRelationalTableModel) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSqlRelationalTableModel::customEvent")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
