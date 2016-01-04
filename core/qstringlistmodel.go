package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"strings"
	"unsafe"
)

type QStringListModel struct {
	QAbstractListModel
}

type QStringListModel_ITF interface {
	QAbstractListModel_ITF
	QStringListModel_PTR() *QStringListModel
}

func PointerFromQStringListModel(ptr QStringListModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStringListModel_PTR().Pointer()
	}
	return nil
}

func NewQStringListModelFromPointer(ptr unsafe.Pointer) *QStringListModel {
	var n = new(QStringListModel)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QStringListModel_") {
		n.SetObjectName("QStringListModel_" + qt.Identifier())
	}
	return n
}

func (ptr *QStringListModel) QStringListModel_PTR() *QStringListModel {
	return ptr
}

func (ptr *QStringListModel) Data(index QModelIndex_ITF, role int) *QVariant {
	defer qt.Recovering("QStringListModel::data")

	if ptr.Pointer() != nil {
		return NewQVariantFromPointer(C.QStringListModel_Data(ptr.Pointer(), PointerFromQModelIndex(index), C.int(role)))
	}
	return nil
}

func (ptr *QStringListModel) Flags(index QModelIndex_ITF) Qt__ItemFlag {
	defer qt.Recovering("QStringListModel::flags")

	if ptr.Pointer() != nil {
		return Qt__ItemFlag(C.QStringListModel_Flags(ptr.Pointer(), PointerFromQModelIndex(index)))
	}
	return 0
}

func (ptr *QStringListModel) InsertRows(row int, count int, parent QModelIndex_ITF) bool {
	defer qt.Recovering("QStringListModel::insertRows")

	if ptr.Pointer() != nil {
		return C.QStringListModel_InsertRows(ptr.Pointer(), C.int(row), C.int(count), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QStringListModel) RemoveRows(row int, count int, parent QModelIndex_ITF) bool {
	defer qt.Recovering("QStringListModel::removeRows")

	if ptr.Pointer() != nil {
		return C.QStringListModel_RemoveRows(ptr.Pointer(), C.int(row), C.int(count), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QStringListModel) RowCount(parent QModelIndex_ITF) int {
	defer qt.Recovering("QStringListModel::rowCount")

	if ptr.Pointer() != nil {
		return int(C.QStringListModel_RowCount(ptr.Pointer(), PointerFromQModelIndex(parent)))
	}
	return 0
}

func (ptr *QStringListModel) SetData(index QModelIndex_ITF, value QVariant_ITF, role int) bool {
	defer qt.Recovering("QStringListModel::setData")

	if ptr.Pointer() != nil {
		return C.QStringListModel_SetData(ptr.Pointer(), PointerFromQModelIndex(index), PointerFromQVariant(value), C.int(role)) != 0
	}
	return false
}

func (ptr *QStringListModel) SetStringList(strin []string) {
	defer qt.Recovering("QStringListModel::setStringList")

	if ptr.Pointer() != nil {
		C.QStringListModel_SetStringList(ptr.Pointer(), C.CString(strings.Join(strin, ",,,")))
	}
}

func (ptr *QStringListModel) Sibling(row int, column int, idx QModelIndex_ITF) *QModelIndex {
	defer qt.Recovering("QStringListModel::sibling")

	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QStringListModel_Sibling(ptr.Pointer(), C.int(row), C.int(column), PointerFromQModelIndex(idx)))
	}
	return nil
}

func (ptr *QStringListModel) ConnectSort(f func(column int, order Qt__SortOrder)) {
	defer qt.Recovering("connect QStringListModel::sort")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "sort", f)
	}
}

func (ptr *QStringListModel) DisconnectSort() {
	defer qt.Recovering("disconnect QStringListModel::sort")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "sort")
	}
}

//export callbackQStringListModelSort
func callbackQStringListModelSort(ptr unsafe.Pointer, ptrName *C.char, column C.int, order C.int) {
	defer qt.Recovering("callback QStringListModel::sort")

	if signal := qt.GetSignal(C.GoString(ptrName), "sort"); signal != nil {
		signal.(func(int, Qt__SortOrder))(int(column), Qt__SortOrder(order))
	} else {
		NewQStringListModelFromPointer(ptr).SortDefault(int(column), Qt__SortOrder(order))
	}
}

func (ptr *QStringListModel) Sort(column int, order Qt__SortOrder) {
	defer qt.Recovering("QStringListModel::sort")

	if ptr.Pointer() != nil {
		C.QStringListModel_Sort(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QStringListModel) SortDefault(column int, order Qt__SortOrder) {
	defer qt.Recovering("QStringListModel::sort")

	if ptr.Pointer() != nil {
		C.QStringListModel_SortDefault(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QStringListModel) StringList() []string {
	defer qt.Recovering("QStringListModel::stringList")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QStringListModel_StringList(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QStringListModel) SupportedDropActions() Qt__DropAction {
	defer qt.Recovering("QStringListModel::supportedDropActions")

	if ptr.Pointer() != nil {
		return Qt__DropAction(C.QStringListModel_SupportedDropActions(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStringListModel) ConnectFetchMore(f func(parent *QModelIndex)) {
	defer qt.Recovering("connect QStringListModel::fetchMore")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "fetchMore", f)
	}
}

func (ptr *QStringListModel) DisconnectFetchMore() {
	defer qt.Recovering("disconnect QStringListModel::fetchMore")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "fetchMore")
	}
}

//export callbackQStringListModelFetchMore
func callbackQStringListModelFetchMore(ptr unsafe.Pointer, ptrName *C.char, parent unsafe.Pointer) {
	defer qt.Recovering("callback QStringListModel::fetchMore")

	if signal := qt.GetSignal(C.GoString(ptrName), "fetchMore"); signal != nil {
		signal.(func(*QModelIndex))(NewQModelIndexFromPointer(parent))
	} else {
		NewQStringListModelFromPointer(ptr).FetchMoreDefault(NewQModelIndexFromPointer(parent))
	}
}

func (ptr *QStringListModel) FetchMore(parent QModelIndex_ITF) {
	defer qt.Recovering("QStringListModel::fetchMore")

	if ptr.Pointer() != nil {
		C.QStringListModel_FetchMore(ptr.Pointer(), PointerFromQModelIndex(parent))
	}
}

func (ptr *QStringListModel) FetchMoreDefault(parent QModelIndex_ITF) {
	defer qt.Recovering("QStringListModel::fetchMore")

	if ptr.Pointer() != nil {
		C.QStringListModel_FetchMoreDefault(ptr.Pointer(), PointerFromQModelIndex(parent))
	}
}

func (ptr *QStringListModel) ConnectRevert(f func()) {
	defer qt.Recovering("connect QStringListModel::revert")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "revert", f)
	}
}

func (ptr *QStringListModel) DisconnectRevert() {
	defer qt.Recovering("disconnect QStringListModel::revert")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "revert")
	}
}

//export callbackQStringListModelRevert
func callbackQStringListModelRevert(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QStringListModel::revert")

	if signal := qt.GetSignal(C.GoString(ptrName), "revert"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QStringListModel) Revert() {
	defer qt.Recovering("QStringListModel::revert")

	if ptr.Pointer() != nil {
		C.QStringListModel_Revert(ptr.Pointer())
	}
}

func (ptr *QStringListModel) RevertDefault() {
	defer qt.Recovering("QStringListModel::revert")

	if ptr.Pointer() != nil {
		C.QStringListModel_RevertDefault(ptr.Pointer())
	}
}

func (ptr *QStringListModel) ConnectTimerEvent(f func(event *QTimerEvent)) {
	defer qt.Recovering("connect QStringListModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QStringListModel) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QStringListModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQStringListModelTimerEvent
func callbackQStringListModelTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStringListModel::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*QTimerEvent))(NewQTimerEventFromPointer(event))
	} else {
		NewQStringListModelFromPointer(ptr).TimerEventDefault(NewQTimerEventFromPointer(event))
	}
}

func (ptr *QStringListModel) TimerEvent(event QTimerEvent_ITF) {
	defer qt.Recovering("QStringListModel::timerEvent")

	if ptr.Pointer() != nil {
		C.QStringListModel_TimerEvent(ptr.Pointer(), PointerFromQTimerEvent(event))
	}
}

func (ptr *QStringListModel) TimerEventDefault(event QTimerEvent_ITF) {
	defer qt.Recovering("QStringListModel::timerEvent")

	if ptr.Pointer() != nil {
		C.QStringListModel_TimerEventDefault(ptr.Pointer(), PointerFromQTimerEvent(event))
	}
}

func (ptr *QStringListModel) ConnectChildEvent(f func(event *QChildEvent)) {
	defer qt.Recovering("connect QStringListModel::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QStringListModel) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QStringListModel::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQStringListModelChildEvent
func callbackQStringListModelChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStringListModel::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*QChildEvent))(NewQChildEventFromPointer(event))
	} else {
		NewQStringListModelFromPointer(ptr).ChildEventDefault(NewQChildEventFromPointer(event))
	}
}

func (ptr *QStringListModel) ChildEvent(event QChildEvent_ITF) {
	defer qt.Recovering("QStringListModel::childEvent")

	if ptr.Pointer() != nil {
		C.QStringListModel_ChildEvent(ptr.Pointer(), PointerFromQChildEvent(event))
	}
}

func (ptr *QStringListModel) ChildEventDefault(event QChildEvent_ITF) {
	defer qt.Recovering("QStringListModel::childEvent")

	if ptr.Pointer() != nil {
		C.QStringListModel_ChildEventDefault(ptr.Pointer(), PointerFromQChildEvent(event))
	}
}

func (ptr *QStringListModel) ConnectCustomEvent(f func(event *QEvent)) {
	defer qt.Recovering("connect QStringListModel::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QStringListModel) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QStringListModel::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQStringListModelCustomEvent
func callbackQStringListModelCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QStringListModel::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
	} else {
		NewQStringListModelFromPointer(ptr).CustomEventDefault(NewQEventFromPointer(event))
	}
}

func (ptr *QStringListModel) CustomEvent(event QEvent_ITF) {
	defer qt.Recovering("QStringListModel::customEvent")

	if ptr.Pointer() != nil {
		C.QStringListModel_CustomEvent(ptr.Pointer(), PointerFromQEvent(event))
	}
}

func (ptr *QStringListModel) CustomEventDefault(event QEvent_ITF) {
	defer qt.Recovering("QStringListModel::customEvent")

	if ptr.Pointer() != nil {
		C.QStringListModel_CustomEventDefault(ptr.Pointer(), PointerFromQEvent(event))
	}
}
