package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"strings"
	"unsafe"
)

type QSortFilterProxyModel struct {
	QAbstractProxyModel
}

type QSortFilterProxyModel_ITF interface {
	QAbstractProxyModel_ITF
	QSortFilterProxyModel_PTR() *QSortFilterProxyModel
}

func PointerFromQSortFilterProxyModel(ptr QSortFilterProxyModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSortFilterProxyModel_PTR().Pointer()
	}
	return nil
}

func NewQSortFilterProxyModelFromPointer(ptr unsafe.Pointer) *QSortFilterProxyModel {
	var n = new(QSortFilterProxyModel)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSortFilterProxyModel_") {
		n.SetObjectName("QSortFilterProxyModel_" + qt.Identifier())
	}
	return n
}

func (ptr *QSortFilterProxyModel) QSortFilterProxyModel_PTR() *QSortFilterProxyModel {
	return ptr
}

func (ptr *QSortFilterProxyModel) DynamicSortFilter() bool {
	defer qt.Recovering("QSortFilterProxyModel::dynamicSortFilter")

	if ptr.Pointer() != nil {
		return C.QSortFilterProxyModel_DynamicSortFilter(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSortFilterProxyModel) FilterCaseSensitivity() Qt__CaseSensitivity {
	defer qt.Recovering("QSortFilterProxyModel::filterCaseSensitivity")

	if ptr.Pointer() != nil {
		return Qt__CaseSensitivity(C.QSortFilterProxyModel_FilterCaseSensitivity(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSortFilterProxyModel) FilterKeyColumn() int {
	defer qt.Recovering("QSortFilterProxyModel::filterKeyColumn")

	if ptr.Pointer() != nil {
		return int(C.QSortFilterProxyModel_FilterKeyColumn(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSortFilterProxyModel) FilterRegExp() *QRegExp {
	defer qt.Recovering("QSortFilterProxyModel::filterRegExp")

	if ptr.Pointer() != nil {
		return NewQRegExpFromPointer(C.QSortFilterProxyModel_FilterRegExp(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSortFilterProxyModel) FilterRole() int {
	defer qt.Recovering("QSortFilterProxyModel::filterRole")

	if ptr.Pointer() != nil {
		return int(C.QSortFilterProxyModel_FilterRole(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSortFilterProxyModel) IsSortLocaleAware() bool {
	defer qt.Recovering("QSortFilterProxyModel::isSortLocaleAware")

	if ptr.Pointer() != nil {
		return C.QSortFilterProxyModel_IsSortLocaleAware(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSortFilterProxyModel) SetDynamicSortFilter(enable bool) {
	defer qt.Recovering("QSortFilterProxyModel::setDynamicSortFilter")

	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_SetDynamicSortFilter(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QSortFilterProxyModel) SetFilterCaseSensitivity(cs Qt__CaseSensitivity) {
	defer qt.Recovering("QSortFilterProxyModel::setFilterCaseSensitivity")

	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_SetFilterCaseSensitivity(ptr.Pointer(), C.int(cs))
	}
}

func (ptr *QSortFilterProxyModel) SetFilterKeyColumn(column int) {
	defer qt.Recovering("QSortFilterProxyModel::setFilterKeyColumn")

	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_SetFilterKeyColumn(ptr.Pointer(), C.int(column))
	}
}

func (ptr *QSortFilterProxyModel) SetFilterRegExp(regExp QRegExp_ITF) {
	defer qt.Recovering("QSortFilterProxyModel::setFilterRegExp")

	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_SetFilterRegExp(ptr.Pointer(), PointerFromQRegExp(regExp))
	}
}

func (ptr *QSortFilterProxyModel) SetFilterRole(role int) {
	defer qt.Recovering("QSortFilterProxyModel::setFilterRole")

	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_SetFilterRole(ptr.Pointer(), C.int(role))
	}
}

func (ptr *QSortFilterProxyModel) SetSortCaseSensitivity(cs Qt__CaseSensitivity) {
	defer qt.Recovering("QSortFilterProxyModel::setSortCaseSensitivity")

	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_SetSortCaseSensitivity(ptr.Pointer(), C.int(cs))
	}
}

func (ptr *QSortFilterProxyModel) SetSortLocaleAware(on bool) {
	defer qt.Recovering("QSortFilterProxyModel::setSortLocaleAware")

	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_SetSortLocaleAware(ptr.Pointer(), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QSortFilterProxyModel) SetSortRole(role int) {
	defer qt.Recovering("QSortFilterProxyModel::setSortRole")

	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_SetSortRole(ptr.Pointer(), C.int(role))
	}
}

func (ptr *QSortFilterProxyModel) SortCaseSensitivity() Qt__CaseSensitivity {
	defer qt.Recovering("QSortFilterProxyModel::sortCaseSensitivity")

	if ptr.Pointer() != nil {
		return Qt__CaseSensitivity(C.QSortFilterProxyModel_SortCaseSensitivity(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSortFilterProxyModel) SortRole() int {
	defer qt.Recovering("QSortFilterProxyModel::sortRole")

	if ptr.Pointer() != nil {
		return int(C.QSortFilterProxyModel_SortRole(ptr.Pointer()))
	}
	return 0
}

func NewQSortFilterProxyModel(parent QObject_ITF) *QSortFilterProxyModel {
	defer qt.Recovering("QSortFilterProxyModel::QSortFilterProxyModel")

	return NewQSortFilterProxyModelFromPointer(C.QSortFilterProxyModel_NewQSortFilterProxyModel(PointerFromQObject(parent)))
}

func (ptr *QSortFilterProxyModel) Buddy(index QModelIndex_ITF) *QModelIndex {
	defer qt.Recovering("QSortFilterProxyModel::buddy")

	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QSortFilterProxyModel_Buddy(ptr.Pointer(), PointerFromQModelIndex(index)))
	}
	return nil
}

func (ptr *QSortFilterProxyModel) CanFetchMore(parent QModelIndex_ITF) bool {
	defer qt.Recovering("QSortFilterProxyModel::canFetchMore")

	if ptr.Pointer() != nil {
		return C.QSortFilterProxyModel_CanFetchMore(ptr.Pointer(), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSortFilterProxyModel) ColumnCount(parent QModelIndex_ITF) int {
	defer qt.Recovering("QSortFilterProxyModel::columnCount")

	if ptr.Pointer() != nil {
		return int(C.QSortFilterProxyModel_ColumnCount(ptr.Pointer(), PointerFromQModelIndex(parent)))
	}
	return 0
}

func (ptr *QSortFilterProxyModel) Data(index QModelIndex_ITF, role int) *QVariant {
	defer qt.Recovering("QSortFilterProxyModel::data")

	if ptr.Pointer() != nil {
		return NewQVariantFromPointer(C.QSortFilterProxyModel_Data(ptr.Pointer(), PointerFromQModelIndex(index), C.int(role)))
	}
	return nil
}

func (ptr *QSortFilterProxyModel) DropMimeData(data QMimeData_ITF, action Qt__DropAction, row int, column int, parent QModelIndex_ITF) bool {
	defer qt.Recovering("QSortFilterProxyModel::dropMimeData")

	if ptr.Pointer() != nil {
		return C.QSortFilterProxyModel_DropMimeData(ptr.Pointer(), PointerFromQMimeData(data), C.int(action), C.int(row), C.int(column), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSortFilterProxyModel) ConnectFetchMore(f func(parent *QModelIndex)) {
	defer qt.Recovering("connect QSortFilterProxyModel::fetchMore")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "fetchMore", f)
	}
}

func (ptr *QSortFilterProxyModel) DisconnectFetchMore() {
	defer qt.Recovering("disconnect QSortFilterProxyModel::fetchMore")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "fetchMore")
	}
}

//export callbackQSortFilterProxyModelFetchMore
func callbackQSortFilterProxyModelFetchMore(ptr unsafe.Pointer, ptrName *C.char, parent unsafe.Pointer) {
	defer qt.Recovering("callback QSortFilterProxyModel::fetchMore")

	if signal := qt.GetSignal(C.GoString(ptrName), "fetchMore"); signal != nil {
		signal.(func(*QModelIndex))(NewQModelIndexFromPointer(parent))
	} else {
		NewQSortFilterProxyModelFromPointer(ptr).FetchMoreDefault(NewQModelIndexFromPointer(parent))
	}
}

func (ptr *QSortFilterProxyModel) FetchMore(parent QModelIndex_ITF) {
	defer qt.Recovering("QSortFilterProxyModel::fetchMore")

	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_FetchMore(ptr.Pointer(), PointerFromQModelIndex(parent))
	}
}

func (ptr *QSortFilterProxyModel) FetchMoreDefault(parent QModelIndex_ITF) {
	defer qt.Recovering("QSortFilterProxyModel::fetchMore")

	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_FetchMoreDefault(ptr.Pointer(), PointerFromQModelIndex(parent))
	}
}

func (ptr *QSortFilterProxyModel) FilterAcceptsColumn(source_column int, source_parent QModelIndex_ITF) bool {
	defer qt.Recovering("QSortFilterProxyModel::filterAcceptsColumn")

	if ptr.Pointer() != nil {
		return C.QSortFilterProxyModel_FilterAcceptsColumn(ptr.Pointer(), C.int(source_column), PointerFromQModelIndex(source_parent)) != 0
	}
	return false
}

func (ptr *QSortFilterProxyModel) FilterAcceptsRow(source_row int, source_parent QModelIndex_ITF) bool {
	defer qt.Recovering("QSortFilterProxyModel::filterAcceptsRow")

	if ptr.Pointer() != nil {
		return C.QSortFilterProxyModel_FilterAcceptsRow(ptr.Pointer(), C.int(source_row), PointerFromQModelIndex(source_parent)) != 0
	}
	return false
}

func (ptr *QSortFilterProxyModel) Flags(index QModelIndex_ITF) Qt__ItemFlag {
	defer qt.Recovering("QSortFilterProxyModel::flags")

	if ptr.Pointer() != nil {
		return Qt__ItemFlag(C.QSortFilterProxyModel_Flags(ptr.Pointer(), PointerFromQModelIndex(index)))
	}
	return 0
}

func (ptr *QSortFilterProxyModel) HasChildren(parent QModelIndex_ITF) bool {
	defer qt.Recovering("QSortFilterProxyModel::hasChildren")

	if ptr.Pointer() != nil {
		return C.QSortFilterProxyModel_HasChildren(ptr.Pointer(), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSortFilterProxyModel) HeaderData(section int, orientation Qt__Orientation, role int) *QVariant {
	defer qt.Recovering("QSortFilterProxyModel::headerData")

	if ptr.Pointer() != nil {
		return NewQVariantFromPointer(C.QSortFilterProxyModel_HeaderData(ptr.Pointer(), C.int(section), C.int(orientation), C.int(role)))
	}
	return nil
}

func (ptr *QSortFilterProxyModel) Index(row int, column int, parent QModelIndex_ITF) *QModelIndex {
	defer qt.Recovering("QSortFilterProxyModel::index")

	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QSortFilterProxyModel_Index(ptr.Pointer(), C.int(row), C.int(column), PointerFromQModelIndex(parent)))
	}
	return nil
}

func (ptr *QSortFilterProxyModel) InsertColumns(column int, count int, parent QModelIndex_ITF) bool {
	defer qt.Recovering("QSortFilterProxyModel::insertColumns")

	if ptr.Pointer() != nil {
		return C.QSortFilterProxyModel_InsertColumns(ptr.Pointer(), C.int(column), C.int(count), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSortFilterProxyModel) InsertRows(row int, count int, parent QModelIndex_ITF) bool {
	defer qt.Recovering("QSortFilterProxyModel::insertRows")

	if ptr.Pointer() != nil {
		return C.QSortFilterProxyModel_InsertRows(ptr.Pointer(), C.int(row), C.int(count), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSortFilterProxyModel) Invalidate() {
	defer qt.Recovering("QSortFilterProxyModel::invalidate")

	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_Invalidate(ptr.Pointer())
	}
}

func (ptr *QSortFilterProxyModel) LessThan(source_left QModelIndex_ITF, source_right QModelIndex_ITF) bool {
	defer qt.Recovering("QSortFilterProxyModel::lessThan")

	if ptr.Pointer() != nil {
		return C.QSortFilterProxyModel_LessThan(ptr.Pointer(), PointerFromQModelIndex(source_left), PointerFromQModelIndex(source_right)) != 0
	}
	return false
}

func (ptr *QSortFilterProxyModel) MapFromSource(sourceIndex QModelIndex_ITF) *QModelIndex {
	defer qt.Recovering("QSortFilterProxyModel::mapFromSource")

	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QSortFilterProxyModel_MapFromSource(ptr.Pointer(), PointerFromQModelIndex(sourceIndex)))
	}
	return nil
}

func (ptr *QSortFilterProxyModel) MapToSource(proxyIndex QModelIndex_ITF) *QModelIndex {
	defer qt.Recovering("QSortFilterProxyModel::mapToSource")

	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QSortFilterProxyModel_MapToSource(ptr.Pointer(), PointerFromQModelIndex(proxyIndex)))
	}
	return nil
}

func (ptr *QSortFilterProxyModel) MimeTypes() []string {
	defer qt.Recovering("QSortFilterProxyModel::mimeTypes")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSortFilterProxyModel_MimeTypes(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QSortFilterProxyModel) Parent(child QModelIndex_ITF) *QModelIndex {
	defer qt.Recovering("QSortFilterProxyModel::parent")

	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QSortFilterProxyModel_Parent(ptr.Pointer(), PointerFromQModelIndex(child)))
	}
	return nil
}

func (ptr *QSortFilterProxyModel) RemoveColumns(column int, count int, parent QModelIndex_ITF) bool {
	defer qt.Recovering("QSortFilterProxyModel::removeColumns")

	if ptr.Pointer() != nil {
		return C.QSortFilterProxyModel_RemoveColumns(ptr.Pointer(), C.int(column), C.int(count), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSortFilterProxyModel) RemoveRows(row int, count int, parent QModelIndex_ITF) bool {
	defer qt.Recovering("QSortFilterProxyModel::removeRows")

	if ptr.Pointer() != nil {
		return C.QSortFilterProxyModel_RemoveRows(ptr.Pointer(), C.int(row), C.int(count), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSortFilterProxyModel) RowCount(parent QModelIndex_ITF) int {
	defer qt.Recovering("QSortFilterProxyModel::rowCount")

	if ptr.Pointer() != nil {
		return int(C.QSortFilterProxyModel_RowCount(ptr.Pointer(), PointerFromQModelIndex(parent)))
	}
	return 0
}

func (ptr *QSortFilterProxyModel) SetData(index QModelIndex_ITF, value QVariant_ITF, role int) bool {
	defer qt.Recovering("QSortFilterProxyModel::setData")

	if ptr.Pointer() != nil {
		return C.QSortFilterProxyModel_SetData(ptr.Pointer(), PointerFromQModelIndex(index), PointerFromQVariant(value), C.int(role)) != 0
	}
	return false
}

func (ptr *QSortFilterProxyModel) SetFilterFixedString(pattern string) {
	defer qt.Recovering("QSortFilterProxyModel::setFilterFixedString")

	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_SetFilterFixedString(ptr.Pointer(), C.CString(pattern))
	}
}

func (ptr *QSortFilterProxyModel) SetFilterRegExp2(pattern string) {
	defer qt.Recovering("QSortFilterProxyModel::setFilterRegExp")

	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_SetFilterRegExp2(ptr.Pointer(), C.CString(pattern))
	}
}

func (ptr *QSortFilterProxyModel) SetFilterWildcard(pattern string) {
	defer qt.Recovering("QSortFilterProxyModel::setFilterWildcard")

	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_SetFilterWildcard(ptr.Pointer(), C.CString(pattern))
	}
}

func (ptr *QSortFilterProxyModel) SetHeaderData(section int, orientation Qt__Orientation, value QVariant_ITF, role int) bool {
	defer qt.Recovering("QSortFilterProxyModel::setHeaderData")

	if ptr.Pointer() != nil {
		return C.QSortFilterProxyModel_SetHeaderData(ptr.Pointer(), C.int(section), C.int(orientation), PointerFromQVariant(value), C.int(role)) != 0
	}
	return false
}

func (ptr *QSortFilterProxyModel) ConnectSetSourceModel(f func(sourceModel *QAbstractItemModel)) {
	defer qt.Recovering("connect QSortFilterProxyModel::setSourceModel")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setSourceModel", f)
	}
}

func (ptr *QSortFilterProxyModel) DisconnectSetSourceModel() {
	defer qt.Recovering("disconnect QSortFilterProxyModel::setSourceModel")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setSourceModel")
	}
}

//export callbackQSortFilterProxyModelSetSourceModel
func callbackQSortFilterProxyModelSetSourceModel(ptr unsafe.Pointer, ptrName *C.char, sourceModel unsafe.Pointer) {
	defer qt.Recovering("callback QSortFilterProxyModel::setSourceModel")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSourceModel"); signal != nil {
		signal.(func(*QAbstractItemModel))(NewQAbstractItemModelFromPointer(sourceModel))
	} else {
		NewQSortFilterProxyModelFromPointer(ptr).SetSourceModelDefault(NewQAbstractItemModelFromPointer(sourceModel))
	}
}

func (ptr *QSortFilterProxyModel) SetSourceModel(sourceModel QAbstractItemModel_ITF) {
	defer qt.Recovering("QSortFilterProxyModel::setSourceModel")

	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_SetSourceModel(ptr.Pointer(), PointerFromQAbstractItemModel(sourceModel))
	}
}

func (ptr *QSortFilterProxyModel) SetSourceModelDefault(sourceModel QAbstractItemModel_ITF) {
	defer qt.Recovering("QSortFilterProxyModel::setSourceModel")

	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_SetSourceModelDefault(ptr.Pointer(), PointerFromQAbstractItemModel(sourceModel))
	}
}

func (ptr *QSortFilterProxyModel) Sibling(row int, column int, idx QModelIndex_ITF) *QModelIndex {
	defer qt.Recovering("QSortFilterProxyModel::sibling")

	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QSortFilterProxyModel_Sibling(ptr.Pointer(), C.int(row), C.int(column), PointerFromQModelIndex(idx)))
	}
	return nil
}

func (ptr *QSortFilterProxyModel) ConnectSort(f func(column int, order Qt__SortOrder)) {
	defer qt.Recovering("connect QSortFilterProxyModel::sort")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "sort", f)
	}
}

func (ptr *QSortFilterProxyModel) DisconnectSort() {
	defer qt.Recovering("disconnect QSortFilterProxyModel::sort")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "sort")
	}
}

//export callbackQSortFilterProxyModelSort
func callbackQSortFilterProxyModelSort(ptr unsafe.Pointer, ptrName *C.char, column C.int, order C.int) {
	defer qt.Recovering("callback QSortFilterProxyModel::sort")

	if signal := qt.GetSignal(C.GoString(ptrName), "sort"); signal != nil {
		signal.(func(int, Qt__SortOrder))(int(column), Qt__SortOrder(order))
	} else {
		NewQSortFilterProxyModelFromPointer(ptr).SortDefault(int(column), Qt__SortOrder(order))
	}
}

func (ptr *QSortFilterProxyModel) Sort(column int, order Qt__SortOrder) {
	defer qt.Recovering("QSortFilterProxyModel::sort")

	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_Sort(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QSortFilterProxyModel) SortDefault(column int, order Qt__SortOrder) {
	defer qt.Recovering("QSortFilterProxyModel::sort")

	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_SortDefault(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QSortFilterProxyModel) SortColumn() int {
	defer qt.Recovering("QSortFilterProxyModel::sortColumn")

	if ptr.Pointer() != nil {
		return int(C.QSortFilterProxyModel_SortColumn(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSortFilterProxyModel) SortOrder() Qt__SortOrder {
	defer qt.Recovering("QSortFilterProxyModel::sortOrder")

	if ptr.Pointer() != nil {
		return Qt__SortOrder(C.QSortFilterProxyModel_SortOrder(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSortFilterProxyModel) Span(index QModelIndex_ITF) *QSize {
	defer qt.Recovering("QSortFilterProxyModel::span")

	if ptr.Pointer() != nil {
		return NewQSizeFromPointer(C.QSortFilterProxyModel_Span(ptr.Pointer(), PointerFromQModelIndex(index)))
	}
	return nil
}

func (ptr *QSortFilterProxyModel) SupportedDropActions() Qt__DropAction {
	defer qt.Recovering("QSortFilterProxyModel::supportedDropActions")

	if ptr.Pointer() != nil {
		return Qt__DropAction(C.QSortFilterProxyModel_SupportedDropActions(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSortFilterProxyModel) DestroyQSortFilterProxyModel() {
	defer qt.Recovering("QSortFilterProxyModel::~QSortFilterProxyModel")

	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_DestroyQSortFilterProxyModel(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSortFilterProxyModel) ConnectRevert(f func()) {
	defer qt.Recovering("connect QSortFilterProxyModel::revert")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "revert", f)
	}
}

func (ptr *QSortFilterProxyModel) DisconnectRevert() {
	defer qt.Recovering("disconnect QSortFilterProxyModel::revert")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "revert")
	}
}

//export callbackQSortFilterProxyModelRevert
func callbackQSortFilterProxyModelRevert(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSortFilterProxyModel::revert")

	if signal := qt.GetSignal(C.GoString(ptrName), "revert"); signal != nil {
		signal.(func())()
	} else {
		NewQSortFilterProxyModelFromPointer(ptr).RevertDefault()
	}
}

func (ptr *QSortFilterProxyModel) Revert() {
	defer qt.Recovering("QSortFilterProxyModel::revert")

	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_Revert(ptr.Pointer())
	}
}

func (ptr *QSortFilterProxyModel) RevertDefault() {
	defer qt.Recovering("QSortFilterProxyModel::revert")

	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_RevertDefault(ptr.Pointer())
	}
}

func (ptr *QSortFilterProxyModel) ConnectTimerEvent(f func(event *QTimerEvent)) {
	defer qt.Recovering("connect QSortFilterProxyModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSortFilterProxyModel) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSortFilterProxyModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSortFilterProxyModelTimerEvent
func callbackQSortFilterProxyModelTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSortFilterProxyModel::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*QTimerEvent))(NewQTimerEventFromPointer(event))
	} else {
		NewQSortFilterProxyModelFromPointer(ptr).TimerEventDefault(NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSortFilterProxyModel) TimerEvent(event QTimerEvent_ITF) {
	defer qt.Recovering("QSortFilterProxyModel::timerEvent")

	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_TimerEvent(ptr.Pointer(), PointerFromQTimerEvent(event))
	}
}

func (ptr *QSortFilterProxyModel) TimerEventDefault(event QTimerEvent_ITF) {
	defer qt.Recovering("QSortFilterProxyModel::timerEvent")

	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_TimerEventDefault(ptr.Pointer(), PointerFromQTimerEvent(event))
	}
}

func (ptr *QSortFilterProxyModel) ConnectChildEvent(f func(event *QChildEvent)) {
	defer qt.Recovering("connect QSortFilterProxyModel::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSortFilterProxyModel) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSortFilterProxyModel::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSortFilterProxyModelChildEvent
func callbackQSortFilterProxyModelChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSortFilterProxyModel::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*QChildEvent))(NewQChildEventFromPointer(event))
	} else {
		NewQSortFilterProxyModelFromPointer(ptr).ChildEventDefault(NewQChildEventFromPointer(event))
	}
}

func (ptr *QSortFilterProxyModel) ChildEvent(event QChildEvent_ITF) {
	defer qt.Recovering("QSortFilterProxyModel::childEvent")

	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_ChildEvent(ptr.Pointer(), PointerFromQChildEvent(event))
	}
}

func (ptr *QSortFilterProxyModel) ChildEventDefault(event QChildEvent_ITF) {
	defer qt.Recovering("QSortFilterProxyModel::childEvent")

	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_ChildEventDefault(ptr.Pointer(), PointerFromQChildEvent(event))
	}
}

func (ptr *QSortFilterProxyModel) ConnectCustomEvent(f func(event *QEvent)) {
	defer qt.Recovering("connect QSortFilterProxyModel::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSortFilterProxyModel) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSortFilterProxyModel::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSortFilterProxyModelCustomEvent
func callbackQSortFilterProxyModelCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSortFilterProxyModel::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
	} else {
		NewQSortFilterProxyModelFromPointer(ptr).CustomEventDefault(NewQEventFromPointer(event))
	}
}

func (ptr *QSortFilterProxyModel) CustomEvent(event QEvent_ITF) {
	defer qt.Recovering("QSortFilterProxyModel::customEvent")

	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_CustomEvent(ptr.Pointer(), PointerFromQEvent(event))
	}
}

func (ptr *QSortFilterProxyModel) CustomEventDefault(event QEvent_ITF) {
	defer qt.Recovering("QSortFilterProxyModel::customEvent")

	if ptr.Pointer() != nil {
		C.QSortFilterProxyModel_CustomEventDefault(ptr.Pointer(), PointerFromQEvent(event))
	}
}
