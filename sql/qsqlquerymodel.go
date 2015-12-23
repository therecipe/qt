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
func callbackQSqlQueryModelClear(ptrName *C.char) bool {
	defer qt.Recovering("callback QSqlQueryModel::clear")

	if signal := qt.GetSignal(C.GoString(ptrName), "clear"); signal != nil {
		signal.(func())()
		return true
	}
	return false

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
func callbackQSqlQueryModelFetchMore(ptrName *C.char, parent unsafe.Pointer) bool {
	defer qt.Recovering("callback QSqlQueryModel::fetchMore")

	if signal := qt.GetSignal(C.GoString(ptrName), "fetchMore"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(parent))
		return true
	}
	return false

}

func (ptr *QSqlQueryModel) HeaderData(section int, orientation core.Qt__Orientation, role int) *core.QVariant {
	defer qt.Recovering("QSqlQueryModel::headerData")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSqlQueryModel_HeaderData(ptr.Pointer(), C.int(section), C.int(orientation), C.int(role)))
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
func callbackQSqlQueryModelQueryChange(ptrName *C.char) bool {
	defer qt.Recovering("callback QSqlQueryModel::queryChange")

	if signal := qt.GetSignal(C.GoString(ptrName), "queryChange"); signal != nil {
		signal.(func())()
		return true
	}
	return false

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
func callbackQSqlQueryModelRevert(ptrName *C.char) bool {
	defer qt.Recovering("callback QSqlQueryModel::revert")

	if signal := qt.GetSignal(C.GoString(ptrName), "revert"); signal != nil {
		signal.(func())()
		return true
	}
	return false

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
func callbackQSqlQueryModelSort(ptrName *C.char, column C.int, order C.int) bool {
	defer qt.Recovering("callback QSqlQueryModel::sort")

	if signal := qt.GetSignal(C.GoString(ptrName), "sort"); signal != nil {
		signal.(func(int, core.Qt__SortOrder))(int(column), core.Qt__SortOrder(order))
		return true
	}
	return false

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
func callbackQSqlQueryModelTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSqlQueryModel::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSqlQueryModelChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSqlQueryModel::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSqlQueryModelCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSqlQueryModel::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
