package sql

//#include "sql.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QSqlQueryModel struct {
	core.QAbstractTableModel
}

type QSqlQueryModel_ITF interface {
	core.QAbstractTableModel_ITF
	QSqlQueryModel_PTR() *QSqlQueryModel
}

func PointerFromQSqlQueryModel(ptr QSqlQueryModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlQueryModel_PTR().Pointer()
	}
	return nil
}

func NewQSqlQueryModelFromPointer(ptr unsafe.Pointer) *QSqlQueryModel {
	var n = new(QSqlQueryModel)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSqlQueryModel_") {
		n.SetObjectName("QSqlQueryModel_" + qt.Identifier())
	}
	return n
}

func (ptr *QSqlQueryModel) QSqlQueryModel_PTR() *QSqlQueryModel {
	return ptr
}

func (ptr *QSqlQueryModel) RowCount(parent core.QModelIndex_ITF) int {
	defer qt.Recovering("QSqlQueryModel::rowCount")

	if ptr.Pointer() != nil {
		return int(C.QSqlQueryModel_RowCount(ptr.Pointer(), core.PointerFromQModelIndex(parent)))
	}
	return 0
}

func (ptr *QSqlQueryModel) Data(item core.QModelIndex_ITF, role int) *core.QVariant {
	defer qt.Recovering("QSqlQueryModel::data")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSqlQueryModel_Data(ptr.Pointer(), core.PointerFromQModelIndex(item), C.int(role)))
	}
	return nil
}

func (ptr *QSqlQueryModel) CanFetchMore(parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlQueryModel::canFetchMore")

	if ptr.Pointer() != nil {
		return C.QSqlQueryModel_CanFetchMore(ptr.Pointer(), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSqlQueryModel) ConnectClear(f func()) {
	defer qt.Recovering("connect QSqlQueryModel::clear")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "clear", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectClear() {
	defer qt.Recovering("disconnect QSqlQueryModel::clear")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "clear")
	}
}

//export callbackQSqlQueryModelClear
func callbackQSqlQueryModelClear(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSqlQueryModel::clear")

	if signal := qt.GetSignal(C.GoString(ptrName), "clear"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlQueryModelFromPointer(ptr).ClearDefault()
	}
}

func (ptr *QSqlQueryModel) Clear() {
	defer qt.Recovering("QSqlQueryModel::clear")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_Clear(ptr.Pointer())
	}
}

func (ptr *QSqlQueryModel) ClearDefault() {
	defer qt.Recovering("QSqlQueryModel::clear")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_ClearDefault(ptr.Pointer())
	}
}

func (ptr *QSqlQueryModel) ColumnCount(index core.QModelIndex_ITF) int {
	defer qt.Recovering("QSqlQueryModel::columnCount")

	if ptr.Pointer() != nil {
		return int(C.QSqlQueryModel_ColumnCount(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return 0
}

func (ptr *QSqlQueryModel) ConnectFetchMore(f func(parent *core.QModelIndex)) {
	defer qt.Recovering("connect QSqlQueryModel::fetchMore")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "fetchMore", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectFetchMore() {
	defer qt.Recovering("disconnect QSqlQueryModel::fetchMore")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "fetchMore")
	}
}

//export callbackQSqlQueryModelFetchMore
func callbackQSqlQueryModelFetchMore(ptr unsafe.Pointer, ptrName *C.char, parent unsafe.Pointer) {
	defer qt.Recovering("callback QSqlQueryModel::fetchMore")

	if signal := qt.GetSignal(C.GoString(ptrName), "fetchMore"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(parent))
	} else {
		NewQSqlQueryModelFromPointer(ptr).FetchMoreDefault(core.NewQModelIndexFromPointer(parent))
	}
}

func (ptr *QSqlQueryModel) FetchMore(parent core.QModelIndex_ITF) {
	defer qt.Recovering("QSqlQueryModel::fetchMore")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_FetchMore(ptr.Pointer(), core.PointerFromQModelIndex(parent))
	}
}

func (ptr *QSqlQueryModel) FetchMoreDefault(parent core.QModelIndex_ITF) {
	defer qt.Recovering("QSqlQueryModel::fetchMore")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_FetchMoreDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent))
	}
}

func (ptr *QSqlQueryModel) HeaderData(section int, orientation core.Qt__Orientation, role int) *core.QVariant {
	defer qt.Recovering("QSqlQueryModel::headerData")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSqlQueryModel_HeaderData(ptr.Pointer(), C.int(section), C.int(orientation), C.int(role)))
	}
	return nil
}

func (ptr *QSqlQueryModel) IndexInQuery(item core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QSqlQueryModel::indexInQuery")

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QSqlQueryModel_IndexInQuery(ptr.Pointer(), core.PointerFromQModelIndex(item)))
	}
	return nil
}

func (ptr *QSqlQueryModel) InsertColumns(column int, count int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlQueryModel::insertColumns")

	if ptr.Pointer() != nil {
		return C.QSqlQueryModel_InsertColumns(ptr.Pointer(), C.int(column), C.int(count), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSqlQueryModel) ConnectQueryChange(f func()) {
	defer qt.Recovering("connect QSqlQueryModel::queryChange")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "queryChange", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectQueryChange() {
	defer qt.Recovering("disconnect QSqlQueryModel::queryChange")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "queryChange")
	}
}

//export callbackQSqlQueryModelQueryChange
func callbackQSqlQueryModelQueryChange(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSqlQueryModel::queryChange")

	if signal := qt.GetSignal(C.GoString(ptrName), "queryChange"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlQueryModelFromPointer(ptr).QueryChangeDefault()
	}
}

func (ptr *QSqlQueryModel) QueryChange() {
	defer qt.Recovering("QSqlQueryModel::queryChange")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_QueryChange(ptr.Pointer())
	}
}

func (ptr *QSqlQueryModel) QueryChangeDefault() {
	defer qt.Recovering("QSqlQueryModel::queryChange")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_QueryChangeDefault(ptr.Pointer())
	}
}

func (ptr *QSqlQueryModel) RemoveColumns(column int, count int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlQueryModel::removeColumns")

	if ptr.Pointer() != nil {
		return C.QSqlQueryModel_RemoveColumns(ptr.Pointer(), C.int(column), C.int(count), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSqlQueryModel) SetHeaderData(section int, orientation core.Qt__Orientation, value core.QVariant_ITF, role int) bool {
	defer qt.Recovering("QSqlQueryModel::setHeaderData")

	if ptr.Pointer() != nil {
		return C.QSqlQueryModel_SetHeaderData(ptr.Pointer(), C.int(section), C.int(orientation), core.PointerFromQVariant(value), C.int(role)) != 0
	}
	return false
}

func (ptr *QSqlQueryModel) SetQuery(query QSqlQuery_ITF) {
	defer qt.Recovering("QSqlQueryModel::setQuery")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_SetQuery(ptr.Pointer(), PointerFromQSqlQuery(query))
	}
}

func (ptr *QSqlQueryModel) SetQuery2(query string, db QSqlDatabase_ITF) {
	defer qt.Recovering("QSqlQueryModel::setQuery")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_SetQuery2(ptr.Pointer(), C.CString(query), PointerFromQSqlDatabase(db))
	}
}

func (ptr *QSqlQueryModel) DestroyQSqlQueryModel() {
	defer qt.Recovering("QSqlQueryModel::~QSqlQueryModel")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_DestroyQSqlQueryModel(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSqlQueryModel) ConnectRevert(f func()) {
	defer qt.Recovering("connect QSqlQueryModel::revert")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "revert", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectRevert() {
	defer qt.Recovering("disconnect QSqlQueryModel::revert")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "revert")
	}
}

//export callbackQSqlQueryModelRevert
func callbackQSqlQueryModelRevert(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QSqlQueryModel::revert")

	if signal := qt.GetSignal(C.GoString(ptrName), "revert"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QSqlQueryModel) Revert() {
	defer qt.Recovering("QSqlQueryModel::revert")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_Revert(ptr.Pointer())
	}
}

func (ptr *QSqlQueryModel) RevertDefault() {
	defer qt.Recovering("QSqlQueryModel::revert")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_RevertDefault(ptr.Pointer())
	}
}

func (ptr *QSqlQueryModel) ConnectSort(f func(column int, order core.Qt__SortOrder)) {
	defer qt.Recovering("connect QSqlQueryModel::sort")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "sort", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectSort() {
	defer qt.Recovering("disconnect QSqlQueryModel::sort")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "sort")
	}
}

//export callbackQSqlQueryModelSort
func callbackQSqlQueryModelSort(ptr unsafe.Pointer, ptrName *C.char, column C.int, order C.int) {
	defer qt.Recovering("callback QSqlQueryModel::sort")

	if signal := qt.GetSignal(C.GoString(ptrName), "sort"); signal != nil {
		signal.(func(int, core.Qt__SortOrder))(int(column), core.Qt__SortOrder(order))
	} else {
		NewQSqlQueryModelFromPointer(ptr).SortDefault(int(column), core.Qt__SortOrder(order))
	}
}

func (ptr *QSqlQueryModel) Sort(column int, order core.Qt__SortOrder) {
	defer qt.Recovering("QSqlQueryModel::sort")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_Sort(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QSqlQueryModel) SortDefault(column int, order core.Qt__SortOrder) {
	defer qt.Recovering("QSqlQueryModel::sort")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_SortDefault(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QSqlQueryModel) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSqlQueryModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSqlQueryModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSqlQueryModelTimerEvent
func callbackQSqlQueryModelTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSqlQueryModel::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSqlQueryModelFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSqlQueryModel) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSqlQueryModel::timerEvent")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSqlQueryModel) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSqlQueryModel::timerEvent")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSqlQueryModel) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSqlQueryModel::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSqlQueryModel::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSqlQueryModelChildEvent
func callbackQSqlQueryModelChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSqlQueryModel::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSqlQueryModelFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSqlQueryModel) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSqlQueryModel::childEvent")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSqlQueryModel) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSqlQueryModel::childEvent")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSqlQueryModel) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSqlQueryModel::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSqlQueryModel::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSqlQueryModelCustomEvent
func callbackQSqlQueryModelCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSqlQueryModel::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSqlQueryModelFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSqlQueryModel) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSqlQueryModel::customEvent")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSqlQueryModel) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSqlQueryModel::customEvent")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
