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
func callbackQIdentityProxyModelSetSourceModel(ptr unsafe.Pointer, ptrName *C.char, newSourceModel unsafe.Pointer) {
	defer qt.Recovering("callback QIdentityProxyModel::setSourceModel")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSourceModel"); signal != nil {
		signal.(func(*QAbstractItemModel))(NewQAbstractItemModelFromPointer(newSourceModel))
	} else {
		NewQIdentityProxyModelFromPointer(ptr).SetSourceModelDefault(NewQAbstractItemModelFromPointer(newSourceModel))
	}
}

func (ptr *QIdentityProxyModel) SetSourceModel(newSourceModel QAbstractItemModel_ITF) {
	defer qt.Recovering("QIdentityProxyModel::setSourceModel")

	if ptr.Pointer() != nil {
		C.QIdentityProxyModel_SetSourceModel(ptr.Pointer(), PointerFromQAbstractItemModel(newSourceModel))
	}
}

func (ptr *QIdentityProxyModel) SetSourceModelDefault(newSourceModel QAbstractItemModel_ITF) {
	defer qt.Recovering("QIdentityProxyModel::setSourceModel")

	if ptr.Pointer() != nil {
		C.QIdentityProxyModel_SetSourceModelDefault(ptr.Pointer(), PointerFromQAbstractItemModel(newSourceModel))
	}
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

func (ptr *QIdentityProxyModel) ConnectFetchMore(f func(parent *QModelIndex)) {
	defer qt.Recovering("connect QIdentityProxyModel::fetchMore")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "fetchMore", f)
	}
}

func (ptr *QIdentityProxyModel) DisconnectFetchMore() {
	defer qt.Recovering("disconnect QIdentityProxyModel::fetchMore")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "fetchMore")
	}
}

//export callbackQIdentityProxyModelFetchMore
func callbackQIdentityProxyModelFetchMore(ptr unsafe.Pointer, ptrName *C.char, parent unsafe.Pointer) {
	defer qt.Recovering("callback QIdentityProxyModel::fetchMore")

	if signal := qt.GetSignal(C.GoString(ptrName), "fetchMore"); signal != nil {
		signal.(func(*QModelIndex))(NewQModelIndexFromPointer(parent))
	} else {
		NewQIdentityProxyModelFromPointer(ptr).FetchMoreDefault(NewQModelIndexFromPointer(parent))
	}
}

func (ptr *QIdentityProxyModel) FetchMore(parent QModelIndex_ITF) {
	defer qt.Recovering("QIdentityProxyModel::fetchMore")

	if ptr.Pointer() != nil {
		C.QIdentityProxyModel_FetchMore(ptr.Pointer(), PointerFromQModelIndex(parent))
	}
}

func (ptr *QIdentityProxyModel) FetchMoreDefault(parent QModelIndex_ITF) {
	defer qt.Recovering("QIdentityProxyModel::fetchMore")

	if ptr.Pointer() != nil {
		C.QIdentityProxyModel_FetchMoreDefault(ptr.Pointer(), PointerFromQModelIndex(parent))
	}
}

func (ptr *QIdentityProxyModel) ConnectRevert(f func()) {
	defer qt.Recovering("connect QIdentityProxyModel::revert")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "revert", f)
	}
}

func (ptr *QIdentityProxyModel) DisconnectRevert() {
	defer qt.Recovering("disconnect QIdentityProxyModel::revert")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "revert")
	}
}

//export callbackQIdentityProxyModelRevert
func callbackQIdentityProxyModelRevert(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QIdentityProxyModel::revert")

	if signal := qt.GetSignal(C.GoString(ptrName), "revert"); signal != nil {
		signal.(func())()
	} else {
		NewQIdentityProxyModelFromPointer(ptr).RevertDefault()
	}
}

func (ptr *QIdentityProxyModel) Revert() {
	defer qt.Recovering("QIdentityProxyModel::revert")

	if ptr.Pointer() != nil {
		C.QIdentityProxyModel_Revert(ptr.Pointer())
	}
}

func (ptr *QIdentityProxyModel) RevertDefault() {
	defer qt.Recovering("QIdentityProxyModel::revert")

	if ptr.Pointer() != nil {
		C.QIdentityProxyModel_RevertDefault(ptr.Pointer())
	}
}

func (ptr *QIdentityProxyModel) ConnectSort(f func(column int, order Qt__SortOrder)) {
	defer qt.Recovering("connect QIdentityProxyModel::sort")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "sort", f)
	}
}

func (ptr *QIdentityProxyModel) DisconnectSort() {
	defer qt.Recovering("disconnect QIdentityProxyModel::sort")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "sort")
	}
}

//export callbackQIdentityProxyModelSort
func callbackQIdentityProxyModelSort(ptr unsafe.Pointer, ptrName *C.char, column C.int, order C.int) {
	defer qt.Recovering("callback QIdentityProxyModel::sort")

	if signal := qt.GetSignal(C.GoString(ptrName), "sort"); signal != nil {
		signal.(func(int, Qt__SortOrder))(int(column), Qt__SortOrder(order))
	} else {
		NewQIdentityProxyModelFromPointer(ptr).SortDefault(int(column), Qt__SortOrder(order))
	}
}

func (ptr *QIdentityProxyModel) Sort(column int, order Qt__SortOrder) {
	defer qt.Recovering("QIdentityProxyModel::sort")

	if ptr.Pointer() != nil {
		C.QIdentityProxyModel_Sort(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QIdentityProxyModel) SortDefault(column int, order Qt__SortOrder) {
	defer qt.Recovering("QIdentityProxyModel::sort")

	if ptr.Pointer() != nil {
		C.QIdentityProxyModel_SortDefault(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QIdentityProxyModel) ConnectTimerEvent(f func(event *QTimerEvent)) {
	defer qt.Recovering("connect QIdentityProxyModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QIdentityProxyModel) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QIdentityProxyModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQIdentityProxyModelTimerEvent
func callbackQIdentityProxyModelTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QIdentityProxyModel::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*QTimerEvent))(NewQTimerEventFromPointer(event))
	} else {
		NewQIdentityProxyModelFromPointer(ptr).TimerEventDefault(NewQTimerEventFromPointer(event))
	}
}

func (ptr *QIdentityProxyModel) TimerEvent(event QTimerEvent_ITF) {
	defer qt.Recovering("QIdentityProxyModel::timerEvent")

	if ptr.Pointer() != nil {
		C.QIdentityProxyModel_TimerEvent(ptr.Pointer(), PointerFromQTimerEvent(event))
	}
}

func (ptr *QIdentityProxyModel) TimerEventDefault(event QTimerEvent_ITF) {
	defer qt.Recovering("QIdentityProxyModel::timerEvent")

	if ptr.Pointer() != nil {
		C.QIdentityProxyModel_TimerEventDefault(ptr.Pointer(), PointerFromQTimerEvent(event))
	}
}

func (ptr *QIdentityProxyModel) ConnectChildEvent(f func(event *QChildEvent)) {
	defer qt.Recovering("connect QIdentityProxyModel::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QIdentityProxyModel) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QIdentityProxyModel::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQIdentityProxyModelChildEvent
func callbackQIdentityProxyModelChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QIdentityProxyModel::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*QChildEvent))(NewQChildEventFromPointer(event))
	} else {
		NewQIdentityProxyModelFromPointer(ptr).ChildEventDefault(NewQChildEventFromPointer(event))
	}
}

func (ptr *QIdentityProxyModel) ChildEvent(event QChildEvent_ITF) {
	defer qt.Recovering("QIdentityProxyModel::childEvent")

	if ptr.Pointer() != nil {
		C.QIdentityProxyModel_ChildEvent(ptr.Pointer(), PointerFromQChildEvent(event))
	}
}

func (ptr *QIdentityProxyModel) ChildEventDefault(event QChildEvent_ITF) {
	defer qt.Recovering("QIdentityProxyModel::childEvent")

	if ptr.Pointer() != nil {
		C.QIdentityProxyModel_ChildEventDefault(ptr.Pointer(), PointerFromQChildEvent(event))
	}
}

func (ptr *QIdentityProxyModel) ConnectCustomEvent(f func(event *QEvent)) {
	defer qt.Recovering("connect QIdentityProxyModel::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QIdentityProxyModel) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QIdentityProxyModel::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQIdentityProxyModelCustomEvent
func callbackQIdentityProxyModelCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QIdentityProxyModel::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
	} else {
		NewQIdentityProxyModelFromPointer(ptr).CustomEventDefault(NewQEventFromPointer(event))
	}
}

func (ptr *QIdentityProxyModel) CustomEvent(event QEvent_ITF) {
	defer qt.Recovering("QIdentityProxyModel::customEvent")

	if ptr.Pointer() != nil {
		C.QIdentityProxyModel_CustomEvent(ptr.Pointer(), PointerFromQEvent(event))
	}
}

func (ptr *QIdentityProxyModel) CustomEventDefault(event QEvent_ITF) {
	defer qt.Recovering("QIdentityProxyModel::customEvent")

	if ptr.Pointer() != nil {
		C.QIdentityProxyModel_CustomEventDefault(ptr.Pointer(), PointerFromQEvent(event))
	}
}
