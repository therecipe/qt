package help

//#include "help.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QHelpContentModel struct {
	core.QAbstractItemModel
}

type QHelpContentModel_ITF interface {
	core.QAbstractItemModel_ITF
	QHelpContentModel_PTR() *QHelpContentModel
}

func PointerFromQHelpContentModel(ptr QHelpContentModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpContentModel_PTR().Pointer()
	}
	return nil
}

func NewQHelpContentModelFromPointer(ptr unsafe.Pointer) *QHelpContentModel {
	var n = new(QHelpContentModel)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QHelpContentModel_") {
		n.SetObjectName("QHelpContentModel_" + qt.Identifier())
	}
	return n
}

func (ptr *QHelpContentModel) QHelpContentModel_PTR() *QHelpContentModel {
	return ptr
}

func (ptr *QHelpContentModel) ColumnCount(parent core.QModelIndex_ITF) int {
	defer qt.Recovering("QHelpContentModel::columnCount")

	if ptr.Pointer() != nil {
		return int(C.QHelpContentModel_ColumnCount(ptr.Pointer(), core.PointerFromQModelIndex(parent)))
	}
	return 0
}

func (ptr *QHelpContentModel) ContentItemAt(index core.QModelIndex_ITF) *QHelpContentItem {
	defer qt.Recovering("QHelpContentModel::contentItemAt")

	if ptr.Pointer() != nil {
		return NewQHelpContentItemFromPointer(C.QHelpContentModel_ContentItemAt(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return nil
}

func (ptr *QHelpContentModel) ConnectContentsCreated(f func()) {
	defer qt.Recovering("connect QHelpContentModel::contentsCreated")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_ConnectContentsCreated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "contentsCreated", f)
	}
}

func (ptr *QHelpContentModel) DisconnectContentsCreated() {
	defer qt.Recovering("disconnect QHelpContentModel::contentsCreated")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_DisconnectContentsCreated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "contentsCreated")
	}
}

//export callbackQHelpContentModelContentsCreated
func callbackQHelpContentModelContentsCreated(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QHelpContentModel::contentsCreated")

	if signal := qt.GetSignal(C.GoString(ptrName), "contentsCreated"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpContentModel) ContentsCreated() {
	defer qt.Recovering("QHelpContentModel::contentsCreated")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_ContentsCreated(ptr.Pointer())
	}
}

func (ptr *QHelpContentModel) ConnectContentsCreationStarted(f func()) {
	defer qt.Recovering("connect QHelpContentModel::contentsCreationStarted")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_ConnectContentsCreationStarted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "contentsCreationStarted", f)
	}
}

func (ptr *QHelpContentModel) DisconnectContentsCreationStarted() {
	defer qt.Recovering("disconnect QHelpContentModel::contentsCreationStarted")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_DisconnectContentsCreationStarted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "contentsCreationStarted")
	}
}

//export callbackQHelpContentModelContentsCreationStarted
func callbackQHelpContentModelContentsCreationStarted(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QHelpContentModel::contentsCreationStarted")

	if signal := qt.GetSignal(C.GoString(ptrName), "contentsCreationStarted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpContentModel) ContentsCreationStarted() {
	defer qt.Recovering("QHelpContentModel::contentsCreationStarted")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_ContentsCreationStarted(ptr.Pointer())
	}
}

func (ptr *QHelpContentModel) CreateContents(customFilterName string) {
	defer qt.Recovering("QHelpContentModel::createContents")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_CreateContents(ptr.Pointer(), C.CString(customFilterName))
	}
}

func (ptr *QHelpContentModel) Data(index core.QModelIndex_ITF, role int) *core.QVariant {
	defer qt.Recovering("QHelpContentModel::data")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QHelpContentModel_Data(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(role)))
	}
	return nil
}

func (ptr *QHelpContentModel) Index(row int, column int, parent core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QHelpContentModel::index")

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QHelpContentModel_Index(ptr.Pointer(), C.int(row), C.int(column), core.PointerFromQModelIndex(parent)))
	}
	return nil
}

func (ptr *QHelpContentModel) IsCreatingContents() bool {
	defer qt.Recovering("QHelpContentModel::isCreatingContents")

	if ptr.Pointer() != nil {
		return C.QHelpContentModel_IsCreatingContents(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHelpContentModel) Parent(index core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QHelpContentModel::parent")

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QHelpContentModel_Parent(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return nil
}

func (ptr *QHelpContentModel) RowCount(parent core.QModelIndex_ITF) int {
	defer qt.Recovering("QHelpContentModel::rowCount")

	if ptr.Pointer() != nil {
		return int(C.QHelpContentModel_RowCount(ptr.Pointer(), core.PointerFromQModelIndex(parent)))
	}
	return 0
}

func (ptr *QHelpContentModel) DestroyQHelpContentModel() {
	defer qt.Recovering("QHelpContentModel::~QHelpContentModel")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_DestroyQHelpContentModel(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QHelpContentModel) ConnectFetchMore(f func(parent *core.QModelIndex)) {
	defer qt.Recovering("connect QHelpContentModel::fetchMore")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "fetchMore", f)
	}
}

func (ptr *QHelpContentModel) DisconnectFetchMore() {
	defer qt.Recovering("disconnect QHelpContentModel::fetchMore")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "fetchMore")
	}
}

//export callbackQHelpContentModelFetchMore
func callbackQHelpContentModelFetchMore(ptr unsafe.Pointer, ptrName *C.char, parent unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentModel::fetchMore")

	if signal := qt.GetSignal(C.GoString(ptrName), "fetchMore"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(parent))
	} else {
		NewQHelpContentModelFromPointer(ptr).FetchMoreDefault(core.NewQModelIndexFromPointer(parent))
	}
}

func (ptr *QHelpContentModel) FetchMore(parent core.QModelIndex_ITF) {
	defer qt.Recovering("QHelpContentModel::fetchMore")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_FetchMore(ptr.Pointer(), core.PointerFromQModelIndex(parent))
	}
}

func (ptr *QHelpContentModel) FetchMoreDefault(parent core.QModelIndex_ITF) {
	defer qt.Recovering("QHelpContentModel::fetchMore")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_FetchMoreDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent))
	}
}

func (ptr *QHelpContentModel) ConnectRevert(f func()) {
	defer qt.Recovering("connect QHelpContentModel::revert")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "revert", f)
	}
}

func (ptr *QHelpContentModel) DisconnectRevert() {
	defer qt.Recovering("disconnect QHelpContentModel::revert")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "revert")
	}
}

//export callbackQHelpContentModelRevert
func callbackQHelpContentModelRevert(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QHelpContentModel::revert")

	if signal := qt.GetSignal(C.GoString(ptrName), "revert"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QHelpContentModel) Revert() {
	defer qt.Recovering("QHelpContentModel::revert")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_Revert(ptr.Pointer())
	}
}

func (ptr *QHelpContentModel) RevertDefault() {
	defer qt.Recovering("QHelpContentModel::revert")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_RevertDefault(ptr.Pointer())
	}
}

func (ptr *QHelpContentModel) ConnectSort(f func(column int, order core.Qt__SortOrder)) {
	defer qt.Recovering("connect QHelpContentModel::sort")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "sort", f)
	}
}

func (ptr *QHelpContentModel) DisconnectSort() {
	defer qt.Recovering("disconnect QHelpContentModel::sort")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "sort")
	}
}

//export callbackQHelpContentModelSort
func callbackQHelpContentModelSort(ptr unsafe.Pointer, ptrName *C.char, column C.int, order C.int) {
	defer qt.Recovering("callback QHelpContentModel::sort")

	if signal := qt.GetSignal(C.GoString(ptrName), "sort"); signal != nil {
		signal.(func(int, core.Qt__SortOrder))(int(column), core.Qt__SortOrder(order))
	} else {
		NewQHelpContentModelFromPointer(ptr).SortDefault(int(column), core.Qt__SortOrder(order))
	}
}

func (ptr *QHelpContentModel) Sort(column int, order core.Qt__SortOrder) {
	defer qt.Recovering("QHelpContentModel::sort")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_Sort(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QHelpContentModel) SortDefault(column int, order core.Qt__SortOrder) {
	defer qt.Recovering("QHelpContentModel::sort")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_SortDefault(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QHelpContentModel) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QHelpContentModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QHelpContentModel) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QHelpContentModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQHelpContentModelTimerEvent
func callbackQHelpContentModelTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentModel::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHelpContentModelFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHelpContentModel) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHelpContentModel::timerEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHelpContentModel) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHelpContentModel::timerEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHelpContentModel) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QHelpContentModel::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QHelpContentModel) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QHelpContentModel::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQHelpContentModelChildEvent
func callbackQHelpContentModelChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentModel::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHelpContentModelFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHelpContentModel) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHelpContentModel::childEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHelpContentModel) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHelpContentModel::childEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHelpContentModel) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpContentModel::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QHelpContentModel) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QHelpContentModel::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQHelpContentModelCustomEvent
func callbackQHelpContentModelCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentModel::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpContentModelFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpContentModel) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpContentModel::customEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpContentModel) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpContentModel::customEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
