package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QStandardItemModel struct {
	core.QAbstractItemModel
}

type QStandardItemModel_ITF interface {
	core.QAbstractItemModel_ITF
	QStandardItemModel_PTR() *QStandardItemModel
}

func PointerFromQStandardItemModel(ptr QStandardItemModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStandardItemModel_PTR().Pointer()
	}
	return nil
}

func NewQStandardItemModelFromPointer(ptr unsafe.Pointer) *QStandardItemModel {
	var n = new(QStandardItemModel)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QStandardItemModel_") {
		n.SetObjectName("QStandardItemModel_" + qt.Identifier())
	}
	return n
}

func (ptr *QStandardItemModel) QStandardItemModel_PTR() *QStandardItemModel {
	return ptr
}

func (ptr *QStandardItemModel) SetSortRole(role int) {
	defer qt.Recovering("QStandardItemModel::setSortRole")

	if ptr.Pointer() != nil {
		C.QStandardItemModel_SetSortRole(ptr.Pointer(), C.int(role))
	}
}

func (ptr *QStandardItemModel) SortRole() int {
	defer qt.Recovering("QStandardItemModel::sortRole")

	if ptr.Pointer() != nil {
		return int(C.QStandardItemModel_SortRole(ptr.Pointer()))
	}
	return 0
}

func NewQStandardItemModel(parent core.QObject_ITF) *QStandardItemModel {
	defer qt.Recovering("QStandardItemModel::QStandardItemModel")

	return NewQStandardItemModelFromPointer(C.QStandardItemModel_NewQStandardItemModel(core.PointerFromQObject(parent)))
}

func NewQStandardItemModel2(rows int, columns int, parent core.QObject_ITF) *QStandardItemModel {
	defer qt.Recovering("QStandardItemModel::QStandardItemModel")

	return NewQStandardItemModelFromPointer(C.QStandardItemModel_NewQStandardItemModel2(C.int(rows), C.int(columns), core.PointerFromQObject(parent)))
}

func (ptr *QStandardItemModel) AppendRow2(item QStandardItem_ITF) {
	defer qt.Recovering("QStandardItemModel::appendRow")

	if ptr.Pointer() != nil {
		C.QStandardItemModel_AppendRow2(ptr.Pointer(), PointerFromQStandardItem(item))
	}
}

func (ptr *QStandardItemModel) Clear() {
	defer qt.Recovering("QStandardItemModel::clear")

	if ptr.Pointer() != nil {
		C.QStandardItemModel_Clear(ptr.Pointer())
	}
}

func (ptr *QStandardItemModel) ColumnCount(parent core.QModelIndex_ITF) int {
	defer qt.Recovering("QStandardItemModel::columnCount")

	if ptr.Pointer() != nil {
		return int(C.QStandardItemModel_ColumnCount(ptr.Pointer(), core.PointerFromQModelIndex(parent)))
	}
	return 0
}

func (ptr *QStandardItemModel) Data(index core.QModelIndex_ITF, role int) *core.QVariant {
	defer qt.Recovering("QStandardItemModel::data")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QStandardItemModel_Data(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(role)))
	}
	return nil
}

func (ptr *QStandardItemModel) DropMimeData(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QStandardItemModel::dropMimeData")

	if ptr.Pointer() != nil {
		return C.QStandardItemModel_DropMimeData(ptr.Pointer(), core.PointerFromQMimeData(data), C.int(action), C.int(row), C.int(column), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QStandardItemModel) Flags(index core.QModelIndex_ITF) core.Qt__ItemFlag {
	defer qt.Recovering("QStandardItemModel::flags")

	if ptr.Pointer() != nil {
		return core.Qt__ItemFlag(C.QStandardItemModel_Flags(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return 0
}

func (ptr *QStandardItemModel) HasChildren(parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QStandardItemModel::hasChildren")

	if ptr.Pointer() != nil {
		return C.QStandardItemModel_HasChildren(ptr.Pointer(), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QStandardItemModel) HeaderData(section int, orientation core.Qt__Orientation, role int) *core.QVariant {
	defer qt.Recovering("QStandardItemModel::headerData")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QStandardItemModel_HeaderData(ptr.Pointer(), C.int(section), C.int(orientation), C.int(role)))
	}
	return nil
}

func (ptr *QStandardItemModel) HorizontalHeaderItem(column int) *QStandardItem {
	defer qt.Recovering("QStandardItemModel::horizontalHeaderItem")

	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItemModel_HorizontalHeaderItem(ptr.Pointer(), C.int(column)))
	}
	return nil
}

func (ptr *QStandardItemModel) Index(row int, column int, parent core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QStandardItemModel::index")

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QStandardItemModel_Index(ptr.Pointer(), C.int(row), C.int(column), core.PointerFromQModelIndex(parent)))
	}
	return nil
}

func (ptr *QStandardItemModel) IndexFromItem(item QStandardItem_ITF) *core.QModelIndex {
	defer qt.Recovering("QStandardItemModel::indexFromItem")

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QStandardItemModel_IndexFromItem(ptr.Pointer(), PointerFromQStandardItem(item)))
	}
	return nil
}

func (ptr *QStandardItemModel) InsertColumn2(column int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QStandardItemModel::insertColumn")

	if ptr.Pointer() != nil {
		return C.QStandardItemModel_InsertColumn2(ptr.Pointer(), C.int(column), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QStandardItemModel) InsertColumns(column int, count int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QStandardItemModel::insertColumns")

	if ptr.Pointer() != nil {
		return C.QStandardItemModel_InsertColumns(ptr.Pointer(), C.int(column), C.int(count), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QStandardItemModel) InsertRow2(row int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QStandardItemModel::insertRow")

	if ptr.Pointer() != nil {
		return C.QStandardItemModel_InsertRow2(ptr.Pointer(), C.int(row), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QStandardItemModel) InsertRow3(row int, item QStandardItem_ITF) {
	defer qt.Recovering("QStandardItemModel::insertRow")

	if ptr.Pointer() != nil {
		C.QStandardItemModel_InsertRow3(ptr.Pointer(), C.int(row), PointerFromQStandardItem(item))
	}
}

func (ptr *QStandardItemModel) InsertRows(row int, count int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QStandardItemModel::insertRows")

	if ptr.Pointer() != nil {
		return C.QStandardItemModel_InsertRows(ptr.Pointer(), C.int(row), C.int(count), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QStandardItemModel) InvisibleRootItem() *QStandardItem {
	defer qt.Recovering("QStandardItemModel::invisibleRootItem")

	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItemModel_InvisibleRootItem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStandardItemModel) Item(row int, column int) *QStandardItem {
	defer qt.Recovering("QStandardItemModel::item")

	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItemModel_Item(ptr.Pointer(), C.int(row), C.int(column)))
	}
	return nil
}

func (ptr *QStandardItemModel) ConnectItemChanged(f func(item *QStandardItem)) {
	defer qt.Recovering("connect QStandardItemModel::itemChanged")

	if ptr.Pointer() != nil {
		C.QStandardItemModel_ConnectItemChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "itemChanged", f)
	}
}

func (ptr *QStandardItemModel) DisconnectItemChanged() {
	defer qt.Recovering("disconnect QStandardItemModel::itemChanged")

	if ptr.Pointer() != nil {
		C.QStandardItemModel_DisconnectItemChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "itemChanged")
	}
}

//export callbackQStandardItemModelItemChanged
func callbackQStandardItemModelItemChanged(ptrName *C.char, item unsafe.Pointer) {
	defer qt.Recovering("callback QStandardItemModel::itemChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "itemChanged"); signal != nil {
		signal.(func(*QStandardItem))(NewQStandardItemFromPointer(item))
	}

}

func (ptr *QStandardItemModel) ItemFromIndex(index core.QModelIndex_ITF) *QStandardItem {
	defer qt.Recovering("QStandardItemModel::itemFromIndex")

	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItemModel_ItemFromIndex(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return nil
}

func (ptr *QStandardItemModel) ItemPrototype() *QStandardItem {
	defer qt.Recovering("QStandardItemModel::itemPrototype")

	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItemModel_ItemPrototype(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStandardItemModel) MimeTypes() []string {
	defer qt.Recovering("QStandardItemModel::mimeTypes")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QStandardItemModel_MimeTypes(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QStandardItemModel) Parent(child core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QStandardItemModel::parent")

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QStandardItemModel_Parent(ptr.Pointer(), core.PointerFromQModelIndex(child)))
	}
	return nil
}

func (ptr *QStandardItemModel) RemoveColumns(column int, count int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QStandardItemModel::removeColumns")

	if ptr.Pointer() != nil {
		return C.QStandardItemModel_RemoveColumns(ptr.Pointer(), C.int(column), C.int(count), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QStandardItemModel) RemoveRows(row int, count int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QStandardItemModel::removeRows")

	if ptr.Pointer() != nil {
		return C.QStandardItemModel_RemoveRows(ptr.Pointer(), C.int(row), C.int(count), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QStandardItemModel) RowCount(parent core.QModelIndex_ITF) int {
	defer qt.Recovering("QStandardItemModel::rowCount")

	if ptr.Pointer() != nil {
		return int(C.QStandardItemModel_RowCount(ptr.Pointer(), core.PointerFromQModelIndex(parent)))
	}
	return 0
}

func (ptr *QStandardItemModel) SetColumnCount(columns int) {
	defer qt.Recovering("QStandardItemModel::setColumnCount")

	if ptr.Pointer() != nil {
		C.QStandardItemModel_SetColumnCount(ptr.Pointer(), C.int(columns))
	}
}

func (ptr *QStandardItemModel) SetData(index core.QModelIndex_ITF, value core.QVariant_ITF, role int) bool {
	defer qt.Recovering("QStandardItemModel::setData")

	if ptr.Pointer() != nil {
		return C.QStandardItemModel_SetData(ptr.Pointer(), core.PointerFromQModelIndex(index), core.PointerFromQVariant(value), C.int(role)) != 0
	}
	return false
}

func (ptr *QStandardItemModel) SetHeaderData(section int, orientation core.Qt__Orientation, value core.QVariant_ITF, role int) bool {
	defer qt.Recovering("QStandardItemModel::setHeaderData")

	if ptr.Pointer() != nil {
		return C.QStandardItemModel_SetHeaderData(ptr.Pointer(), C.int(section), C.int(orientation), core.PointerFromQVariant(value), C.int(role)) != 0
	}
	return false
}

func (ptr *QStandardItemModel) SetHorizontalHeaderItem(column int, item QStandardItem_ITF) {
	defer qt.Recovering("QStandardItemModel::setHorizontalHeaderItem")

	if ptr.Pointer() != nil {
		C.QStandardItemModel_SetHorizontalHeaderItem(ptr.Pointer(), C.int(column), PointerFromQStandardItem(item))
	}
}

func (ptr *QStandardItemModel) SetHorizontalHeaderLabels(labels []string) {
	defer qt.Recovering("QStandardItemModel::setHorizontalHeaderLabels")

	if ptr.Pointer() != nil {
		C.QStandardItemModel_SetHorizontalHeaderLabels(ptr.Pointer(), C.CString(strings.Join(labels, ",,,")))
	}
}

func (ptr *QStandardItemModel) SetItem2(row int, item QStandardItem_ITF) {
	defer qt.Recovering("QStandardItemModel::setItem")

	if ptr.Pointer() != nil {
		C.QStandardItemModel_SetItem2(ptr.Pointer(), C.int(row), PointerFromQStandardItem(item))
	}
}

func (ptr *QStandardItemModel) SetItem(row int, column int, item QStandardItem_ITF) {
	defer qt.Recovering("QStandardItemModel::setItem")

	if ptr.Pointer() != nil {
		C.QStandardItemModel_SetItem(ptr.Pointer(), C.int(row), C.int(column), PointerFromQStandardItem(item))
	}
}

func (ptr *QStandardItemModel) SetItemPrototype(item QStandardItem_ITF) {
	defer qt.Recovering("QStandardItemModel::setItemPrototype")

	if ptr.Pointer() != nil {
		C.QStandardItemModel_SetItemPrototype(ptr.Pointer(), PointerFromQStandardItem(item))
	}
}

func (ptr *QStandardItemModel) SetRowCount(rows int) {
	defer qt.Recovering("QStandardItemModel::setRowCount")

	if ptr.Pointer() != nil {
		C.QStandardItemModel_SetRowCount(ptr.Pointer(), C.int(rows))
	}
}

func (ptr *QStandardItemModel) SetVerticalHeaderItem(row int, item QStandardItem_ITF) {
	defer qt.Recovering("QStandardItemModel::setVerticalHeaderItem")

	if ptr.Pointer() != nil {
		C.QStandardItemModel_SetVerticalHeaderItem(ptr.Pointer(), C.int(row), PointerFromQStandardItem(item))
	}
}

func (ptr *QStandardItemModel) SetVerticalHeaderLabels(labels []string) {
	defer qt.Recovering("QStandardItemModel::setVerticalHeaderLabels")

	if ptr.Pointer() != nil {
		C.QStandardItemModel_SetVerticalHeaderLabels(ptr.Pointer(), C.CString(strings.Join(labels, ",,,")))
	}
}

func (ptr *QStandardItemModel) Sibling(row int, column int, idx core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QStandardItemModel::sibling")

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QStandardItemModel_Sibling(ptr.Pointer(), C.int(row), C.int(column), core.PointerFromQModelIndex(idx)))
	}
	return nil
}

func (ptr *QStandardItemModel) ConnectSort(f func(column int, order core.Qt__SortOrder)) {
	defer qt.Recovering("connect QStandardItemModel::sort")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "sort", f)
	}
}

func (ptr *QStandardItemModel) DisconnectSort() {
	defer qt.Recovering("disconnect QStandardItemModel::sort")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "sort")
	}
}

//export callbackQStandardItemModelSort
func callbackQStandardItemModelSort(ptrName *C.char, column C.int, order C.int) bool {
	defer qt.Recovering("callback QStandardItemModel::sort")

	if signal := qt.GetSignal(C.GoString(ptrName), "sort"); signal != nil {
		signal.(func(int, core.Qt__SortOrder))(int(column), core.Qt__SortOrder(order))
		return true
	}
	return false

}

func (ptr *QStandardItemModel) SupportedDropActions() core.Qt__DropAction {
	defer qt.Recovering("QStandardItemModel::supportedDropActions")

	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QStandardItemModel_SupportedDropActions(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStandardItemModel) TakeHorizontalHeaderItem(column int) *QStandardItem {
	defer qt.Recovering("QStandardItemModel::takeHorizontalHeaderItem")

	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItemModel_TakeHorizontalHeaderItem(ptr.Pointer(), C.int(column)))
	}
	return nil
}

func (ptr *QStandardItemModel) TakeItem(row int, column int) *QStandardItem {
	defer qt.Recovering("QStandardItemModel::takeItem")

	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItemModel_TakeItem(ptr.Pointer(), C.int(row), C.int(column)))
	}
	return nil
}

func (ptr *QStandardItemModel) TakeVerticalHeaderItem(row int) *QStandardItem {
	defer qt.Recovering("QStandardItemModel::takeVerticalHeaderItem")

	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItemModel_TakeVerticalHeaderItem(ptr.Pointer(), C.int(row)))
	}
	return nil
}

func (ptr *QStandardItemModel) VerticalHeaderItem(row int) *QStandardItem {
	defer qt.Recovering("QStandardItemModel::verticalHeaderItem")

	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItemModel_VerticalHeaderItem(ptr.Pointer(), C.int(row)))
	}
	return nil
}

func (ptr *QStandardItemModel) DestroyQStandardItemModel() {
	defer qt.Recovering("QStandardItemModel::~QStandardItemModel")

	if ptr.Pointer() != nil {
		C.QStandardItemModel_DestroyQStandardItemModel(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QStandardItemModel) ConnectFetchMore(f func(parent *core.QModelIndex)) {
	defer qt.Recovering("connect QStandardItemModel::fetchMore")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "fetchMore", f)
	}
}

func (ptr *QStandardItemModel) DisconnectFetchMore() {
	defer qt.Recovering("disconnect QStandardItemModel::fetchMore")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "fetchMore")
	}
}

//export callbackQStandardItemModelFetchMore
func callbackQStandardItemModelFetchMore(ptrName *C.char, parent unsafe.Pointer) bool {
	defer qt.Recovering("callback QStandardItemModel::fetchMore")

	if signal := qt.GetSignal(C.GoString(ptrName), "fetchMore"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(parent))
		return true
	}
	return false

}

func (ptr *QStandardItemModel) ConnectRevert(f func()) {
	defer qt.Recovering("connect QStandardItemModel::revert")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "revert", f)
	}
}

func (ptr *QStandardItemModel) DisconnectRevert() {
	defer qt.Recovering("disconnect QStandardItemModel::revert")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "revert")
	}
}

//export callbackQStandardItemModelRevert
func callbackQStandardItemModelRevert(ptrName *C.char) bool {
	defer qt.Recovering("callback QStandardItemModel::revert")

	if signal := qt.GetSignal(C.GoString(ptrName), "revert"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QStandardItemModel) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QStandardItemModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QStandardItemModel) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QStandardItemModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQStandardItemModelTimerEvent
func callbackQStandardItemModelTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QStandardItemModel::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QStandardItemModel) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QStandardItemModel::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QStandardItemModel) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QStandardItemModel::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQStandardItemModelChildEvent
func callbackQStandardItemModelChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QStandardItemModel::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QStandardItemModel) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QStandardItemModel::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QStandardItemModel) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QStandardItemModel::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQStandardItemModelCustomEvent
func callbackQStandardItemModelCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QStandardItemModel::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
