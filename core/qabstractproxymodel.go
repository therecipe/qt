package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"strings"
	"unsafe"
)

type QAbstractProxyModel struct {
	QAbstractItemModel
}

type QAbstractProxyModel_ITF interface {
	QAbstractItemModel_ITF
	QAbstractProxyModel_PTR() *QAbstractProxyModel
}

func PointerFromQAbstractProxyModel(ptr QAbstractProxyModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractProxyModel_PTR().Pointer()
	}
	return nil
}

func NewQAbstractProxyModelFromPointer(ptr unsafe.Pointer) *QAbstractProxyModel {
	var n = new(QAbstractProxyModel)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAbstractProxyModel_") {
		n.SetObjectName("QAbstractProxyModel_" + qt.Identifier())
	}
	return n
}

func (ptr *QAbstractProxyModel) QAbstractProxyModel_PTR() *QAbstractProxyModel {
	return ptr
}

func (ptr *QAbstractProxyModel) Buddy(index QModelIndex_ITF) *QModelIndex {
	defer qt.Recovering("QAbstractProxyModel::buddy")

	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QAbstractProxyModel_Buddy(ptr.Pointer(), PointerFromQModelIndex(index)))
	}
	return nil
}

func (ptr *QAbstractProxyModel) CanDropMimeData(data QMimeData_ITF, action Qt__DropAction, row int, column int, parent QModelIndex_ITF) bool {
	defer qt.Recovering("QAbstractProxyModel::canDropMimeData")

	if ptr.Pointer() != nil {
		return C.QAbstractProxyModel_CanDropMimeData(ptr.Pointer(), PointerFromQMimeData(data), C.int(action), C.int(row), C.int(column), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QAbstractProxyModel) CanFetchMore(parent QModelIndex_ITF) bool {
	defer qt.Recovering("QAbstractProxyModel::canFetchMore")

	if ptr.Pointer() != nil {
		return C.QAbstractProxyModel_CanFetchMore(ptr.Pointer(), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QAbstractProxyModel) Data(proxyIndex QModelIndex_ITF, role int) *QVariant {
	defer qt.Recovering("QAbstractProxyModel::data")

	if ptr.Pointer() != nil {
		return NewQVariantFromPointer(C.QAbstractProxyModel_Data(ptr.Pointer(), PointerFromQModelIndex(proxyIndex), C.int(role)))
	}
	return nil
}

func (ptr *QAbstractProxyModel) DropMimeData(data QMimeData_ITF, action Qt__DropAction, row int, column int, parent QModelIndex_ITF) bool {
	defer qt.Recovering("QAbstractProxyModel::dropMimeData")

	if ptr.Pointer() != nil {
		return C.QAbstractProxyModel_DropMimeData(ptr.Pointer(), PointerFromQMimeData(data), C.int(action), C.int(row), C.int(column), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QAbstractProxyModel) ConnectFetchMore(f func(parent *QModelIndex)) {
	defer qt.Recovering("connect QAbstractProxyModel::fetchMore")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "fetchMore", f)
	}
}

func (ptr *QAbstractProxyModel) DisconnectFetchMore() {
	defer qt.Recovering("disconnect QAbstractProxyModel::fetchMore")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "fetchMore")
	}
}

//export callbackQAbstractProxyModelFetchMore
func callbackQAbstractProxyModelFetchMore(ptr unsafe.Pointer, ptrName *C.char, parent unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractProxyModel::fetchMore")

	if signal := qt.GetSignal(C.GoString(ptrName), "fetchMore"); signal != nil {
		signal.(func(*QModelIndex))(NewQModelIndexFromPointer(parent))
	} else {
		NewQAbstractProxyModelFromPointer(ptr).FetchMoreDefault(NewQModelIndexFromPointer(parent))
	}
}

func (ptr *QAbstractProxyModel) FetchMore(parent QModelIndex_ITF) {
	defer qt.Recovering("QAbstractProxyModel::fetchMore")

	if ptr.Pointer() != nil {
		C.QAbstractProxyModel_FetchMore(ptr.Pointer(), PointerFromQModelIndex(parent))
	}
}

func (ptr *QAbstractProxyModel) FetchMoreDefault(parent QModelIndex_ITF) {
	defer qt.Recovering("QAbstractProxyModel::fetchMore")

	if ptr.Pointer() != nil {
		C.QAbstractProxyModel_FetchMoreDefault(ptr.Pointer(), PointerFromQModelIndex(parent))
	}
}

func (ptr *QAbstractProxyModel) Flags(index QModelIndex_ITF) Qt__ItemFlag {
	defer qt.Recovering("QAbstractProxyModel::flags")

	if ptr.Pointer() != nil {
		return Qt__ItemFlag(C.QAbstractProxyModel_Flags(ptr.Pointer(), PointerFromQModelIndex(index)))
	}
	return 0
}

func (ptr *QAbstractProxyModel) HasChildren(parent QModelIndex_ITF) bool {
	defer qt.Recovering("QAbstractProxyModel::hasChildren")

	if ptr.Pointer() != nil {
		return C.QAbstractProxyModel_HasChildren(ptr.Pointer(), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QAbstractProxyModel) HeaderData(section int, orientation Qt__Orientation, role int) *QVariant {
	defer qt.Recovering("QAbstractProxyModel::headerData")

	if ptr.Pointer() != nil {
		return NewQVariantFromPointer(C.QAbstractProxyModel_HeaderData(ptr.Pointer(), C.int(section), C.int(orientation), C.int(role)))
	}
	return nil
}

func (ptr *QAbstractProxyModel) MapFromSource(sourceIndex QModelIndex_ITF) *QModelIndex {
	defer qt.Recovering("QAbstractProxyModel::mapFromSource")

	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QAbstractProxyModel_MapFromSource(ptr.Pointer(), PointerFromQModelIndex(sourceIndex)))
	}
	return nil
}

func (ptr *QAbstractProxyModel) MapToSource(proxyIndex QModelIndex_ITF) *QModelIndex {
	defer qt.Recovering("QAbstractProxyModel::mapToSource")

	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QAbstractProxyModel_MapToSource(ptr.Pointer(), PointerFromQModelIndex(proxyIndex)))
	}
	return nil
}

func (ptr *QAbstractProxyModel) MimeTypes() []string {
	defer qt.Recovering("QAbstractProxyModel::mimeTypes")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QAbstractProxyModel_MimeTypes(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QAbstractProxyModel) ConnectRevert(f func()) {
	defer qt.Recovering("connect QAbstractProxyModel::revert")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "revert", f)
	}
}

func (ptr *QAbstractProxyModel) DisconnectRevert() {
	defer qt.Recovering("disconnect QAbstractProxyModel::revert")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "revert")
	}
}

//export callbackQAbstractProxyModelRevert
func callbackQAbstractProxyModelRevert(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAbstractProxyModel::revert")

	if signal := qt.GetSignal(C.GoString(ptrName), "revert"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractProxyModelFromPointer(ptr).RevertDefault()
	}
}

func (ptr *QAbstractProxyModel) Revert() {
	defer qt.Recovering("QAbstractProxyModel::revert")

	if ptr.Pointer() != nil {
		C.QAbstractProxyModel_Revert(ptr.Pointer())
	}
}

func (ptr *QAbstractProxyModel) RevertDefault() {
	defer qt.Recovering("QAbstractProxyModel::revert")

	if ptr.Pointer() != nil {
		C.QAbstractProxyModel_RevertDefault(ptr.Pointer())
	}
}

func (ptr *QAbstractProxyModel) SetData(index QModelIndex_ITF, value QVariant_ITF, role int) bool {
	defer qt.Recovering("QAbstractProxyModel::setData")

	if ptr.Pointer() != nil {
		return C.QAbstractProxyModel_SetData(ptr.Pointer(), PointerFromQModelIndex(index), PointerFromQVariant(value), C.int(role)) != 0
	}
	return false
}

func (ptr *QAbstractProxyModel) SetHeaderData(section int, orientation Qt__Orientation, value QVariant_ITF, role int) bool {
	defer qt.Recovering("QAbstractProxyModel::setHeaderData")

	if ptr.Pointer() != nil {
		return C.QAbstractProxyModel_SetHeaderData(ptr.Pointer(), C.int(section), C.int(orientation), PointerFromQVariant(value), C.int(role)) != 0
	}
	return false
}

func (ptr *QAbstractProxyModel) ConnectSetSourceModel(f func(sourceModel *QAbstractItemModel)) {
	defer qt.Recovering("connect QAbstractProxyModel::setSourceModel")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setSourceModel", f)
	}
}

func (ptr *QAbstractProxyModel) DisconnectSetSourceModel() {
	defer qt.Recovering("disconnect QAbstractProxyModel::setSourceModel")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setSourceModel")
	}
}

//export callbackQAbstractProxyModelSetSourceModel
func callbackQAbstractProxyModelSetSourceModel(ptr unsafe.Pointer, ptrName *C.char, sourceModel unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractProxyModel::setSourceModel")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSourceModel"); signal != nil {
		signal.(func(*QAbstractItemModel))(NewQAbstractItemModelFromPointer(sourceModel))
	} else {
		NewQAbstractProxyModelFromPointer(ptr).SetSourceModelDefault(NewQAbstractItemModelFromPointer(sourceModel))
	}
}

func (ptr *QAbstractProxyModel) SetSourceModel(sourceModel QAbstractItemModel_ITF) {
	defer qt.Recovering("QAbstractProxyModel::setSourceModel")

	if ptr.Pointer() != nil {
		C.QAbstractProxyModel_SetSourceModel(ptr.Pointer(), PointerFromQAbstractItemModel(sourceModel))
	}
}

func (ptr *QAbstractProxyModel) SetSourceModelDefault(sourceModel QAbstractItemModel_ITF) {
	defer qt.Recovering("QAbstractProxyModel::setSourceModel")

	if ptr.Pointer() != nil {
		C.QAbstractProxyModel_SetSourceModelDefault(ptr.Pointer(), PointerFromQAbstractItemModel(sourceModel))
	}
}

func (ptr *QAbstractProxyModel) Sibling(row int, column int, idx QModelIndex_ITF) *QModelIndex {
	defer qt.Recovering("QAbstractProxyModel::sibling")

	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QAbstractProxyModel_Sibling(ptr.Pointer(), C.int(row), C.int(column), PointerFromQModelIndex(idx)))
	}
	return nil
}

func (ptr *QAbstractProxyModel) ConnectSort(f func(column int, order Qt__SortOrder)) {
	defer qt.Recovering("connect QAbstractProxyModel::sort")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "sort", f)
	}
}

func (ptr *QAbstractProxyModel) DisconnectSort() {
	defer qt.Recovering("disconnect QAbstractProxyModel::sort")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "sort")
	}
}

//export callbackQAbstractProxyModelSort
func callbackQAbstractProxyModelSort(ptr unsafe.Pointer, ptrName *C.char, column C.int, order C.int) {
	defer qt.Recovering("callback QAbstractProxyModel::sort")

	if signal := qt.GetSignal(C.GoString(ptrName), "sort"); signal != nil {
		signal.(func(int, Qt__SortOrder))(int(column), Qt__SortOrder(order))
	} else {
		NewQAbstractProxyModelFromPointer(ptr).SortDefault(int(column), Qt__SortOrder(order))
	}
}

func (ptr *QAbstractProxyModel) Sort(column int, order Qt__SortOrder) {
	defer qt.Recovering("QAbstractProxyModel::sort")

	if ptr.Pointer() != nil {
		C.QAbstractProxyModel_Sort(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QAbstractProxyModel) SortDefault(column int, order Qt__SortOrder) {
	defer qt.Recovering("QAbstractProxyModel::sort")

	if ptr.Pointer() != nil {
		C.QAbstractProxyModel_SortDefault(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QAbstractProxyModel) SourceModel() *QAbstractItemModel {
	defer qt.Recovering("QAbstractProxyModel::sourceModel")

	if ptr.Pointer() != nil {
		return NewQAbstractItemModelFromPointer(C.QAbstractProxyModel_SourceModel(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractProxyModel) ConnectSourceModelChanged(f func()) {
	defer qt.Recovering("connect QAbstractProxyModel::sourceModelChanged")

	if ptr.Pointer() != nil {
		C.QAbstractProxyModel_ConnectSourceModelChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sourceModelChanged", f)
	}
}

func (ptr *QAbstractProxyModel) DisconnectSourceModelChanged() {
	defer qt.Recovering("disconnect QAbstractProxyModel::sourceModelChanged")

	if ptr.Pointer() != nil {
		C.QAbstractProxyModel_DisconnectSourceModelChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sourceModelChanged")
	}
}

//export callbackQAbstractProxyModelSourceModelChanged
func callbackQAbstractProxyModelSourceModelChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAbstractProxyModel::sourceModelChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "sourceModelChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractProxyModel) Span(index QModelIndex_ITF) *QSize {
	defer qt.Recovering("QAbstractProxyModel::span")

	if ptr.Pointer() != nil {
		return NewQSizeFromPointer(C.QAbstractProxyModel_Span(ptr.Pointer(), PointerFromQModelIndex(index)))
	}
	return nil
}

func (ptr *QAbstractProxyModel) Submit() bool {
	defer qt.Recovering("QAbstractProxyModel::submit")

	if ptr.Pointer() != nil {
		return C.QAbstractProxyModel_Submit(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractProxyModel) SupportedDragActions() Qt__DropAction {
	defer qt.Recovering("QAbstractProxyModel::supportedDragActions")

	if ptr.Pointer() != nil {
		return Qt__DropAction(C.QAbstractProxyModel_SupportedDragActions(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractProxyModel) SupportedDropActions() Qt__DropAction {
	defer qt.Recovering("QAbstractProxyModel::supportedDropActions")

	if ptr.Pointer() != nil {
		return Qt__DropAction(C.QAbstractProxyModel_SupportedDropActions(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractProxyModel) DestroyQAbstractProxyModel() {
	defer qt.Recovering("QAbstractProxyModel::~QAbstractProxyModel")

	if ptr.Pointer() != nil {
		C.QAbstractProxyModel_DestroyQAbstractProxyModel(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractProxyModel) ConnectTimerEvent(f func(event *QTimerEvent)) {
	defer qt.Recovering("connect QAbstractProxyModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAbstractProxyModel) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAbstractProxyModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAbstractProxyModelTimerEvent
func callbackQAbstractProxyModelTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractProxyModel::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*QTimerEvent))(NewQTimerEventFromPointer(event))
	} else {
		NewQAbstractProxyModelFromPointer(ptr).TimerEventDefault(NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAbstractProxyModel) TimerEvent(event QTimerEvent_ITF) {
	defer qt.Recovering("QAbstractProxyModel::timerEvent")

	if ptr.Pointer() != nil {
		C.QAbstractProxyModel_TimerEvent(ptr.Pointer(), PointerFromQTimerEvent(event))
	}
}

func (ptr *QAbstractProxyModel) TimerEventDefault(event QTimerEvent_ITF) {
	defer qt.Recovering("QAbstractProxyModel::timerEvent")

	if ptr.Pointer() != nil {
		C.QAbstractProxyModel_TimerEventDefault(ptr.Pointer(), PointerFromQTimerEvent(event))
	}
}

func (ptr *QAbstractProxyModel) ConnectChildEvent(f func(event *QChildEvent)) {
	defer qt.Recovering("connect QAbstractProxyModel::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAbstractProxyModel) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAbstractProxyModel::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAbstractProxyModelChildEvent
func callbackQAbstractProxyModelChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractProxyModel::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*QChildEvent))(NewQChildEventFromPointer(event))
	} else {
		NewQAbstractProxyModelFromPointer(ptr).ChildEventDefault(NewQChildEventFromPointer(event))
	}
}

func (ptr *QAbstractProxyModel) ChildEvent(event QChildEvent_ITF) {
	defer qt.Recovering("QAbstractProxyModel::childEvent")

	if ptr.Pointer() != nil {
		C.QAbstractProxyModel_ChildEvent(ptr.Pointer(), PointerFromQChildEvent(event))
	}
}

func (ptr *QAbstractProxyModel) ChildEventDefault(event QChildEvent_ITF) {
	defer qt.Recovering("QAbstractProxyModel::childEvent")

	if ptr.Pointer() != nil {
		C.QAbstractProxyModel_ChildEventDefault(ptr.Pointer(), PointerFromQChildEvent(event))
	}
}

func (ptr *QAbstractProxyModel) ConnectCustomEvent(f func(event *QEvent)) {
	defer qt.Recovering("connect QAbstractProxyModel::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAbstractProxyModel) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAbstractProxyModel::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAbstractProxyModelCustomEvent
func callbackQAbstractProxyModelCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractProxyModel::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
	} else {
		NewQAbstractProxyModelFromPointer(ptr).CustomEventDefault(NewQEventFromPointer(event))
	}
}

func (ptr *QAbstractProxyModel) CustomEvent(event QEvent_ITF) {
	defer qt.Recovering("QAbstractProxyModel::customEvent")

	if ptr.Pointer() != nil {
		C.QAbstractProxyModel_CustomEvent(ptr.Pointer(), PointerFromQEvent(event))
	}
}

func (ptr *QAbstractProxyModel) CustomEventDefault(event QEvent_ITF) {
	defer qt.Recovering("QAbstractProxyModel::customEvent")

	if ptr.Pointer() != nil {
		C.QAbstractProxyModel_CustomEventDefault(ptr.Pointer(), PointerFromQEvent(event))
	}
}
