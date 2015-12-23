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
func callbackQStringListModelSort(ptrName *C.char, column C.int, order C.int) bool {
	defer qt.Recovering("callback QStringListModel::sort")

	if signal := qt.GetSignal(C.GoString(ptrName), "sort"); signal != nil {
		signal.(func(int, Qt__SortOrder))(int(column), Qt__SortOrder(order))
		return true
	}
	return false

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
func callbackQStringListModelFetchMore(ptrName *C.char, parent unsafe.Pointer) bool {
	defer qt.Recovering("callback QStringListModel::fetchMore")

	if signal := qt.GetSignal(C.GoString(ptrName), "fetchMore"); signal != nil {
		signal.(func(*QModelIndex))(NewQModelIndexFromPointer(parent))
		return true
	}
	return false

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
func callbackQStringListModelRevert(ptrName *C.char) bool {
	defer qt.Recovering("callback QStringListModel::revert")

	if signal := qt.GetSignal(C.GoString(ptrName), "revert"); signal != nil {
		signal.(func())()
		return true
	}
	return false

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
func callbackQStringListModelTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QStringListModel::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*QTimerEvent))(NewQTimerEventFromPointer(event))
		return true
	}
	return false

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
func callbackQStringListModelChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QStringListModel::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*QChildEvent))(NewQChildEventFromPointer(event))
		return true
	}
	return false

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
func callbackQStringListModelCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QStringListModel::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
		return true
	}
	return false

}
