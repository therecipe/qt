package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QAbstractListModel struct {
	QAbstractItemModel
}

type QAbstractListModel_ITF interface {
	QAbstractItemModel_ITF
	QAbstractListModel_PTR() *QAbstractListModel
}

func PointerFromQAbstractListModel(ptr QAbstractListModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractListModel_PTR().Pointer()
	}
	return nil
}

func NewQAbstractListModelFromPointer(ptr unsafe.Pointer) *QAbstractListModel {
	var n = new(QAbstractListModel)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAbstractListModel_") {
		n.SetObjectName("QAbstractListModel_" + qt.Identifier())
	}
	return n
}

func (ptr *QAbstractListModel) QAbstractListModel_PTR() *QAbstractListModel {
	return ptr
}

func (ptr *QAbstractListModel) Index(row int, column int, parent QModelIndex_ITF) *QModelIndex {
	defer qt.Recovering("QAbstractListModel::index")

	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QAbstractListModel_Index(ptr.Pointer(), C.int(row), C.int(column), PointerFromQModelIndex(parent)))
	}
	return nil
}

func (ptr *QAbstractListModel) DropMimeData(data QMimeData_ITF, action Qt__DropAction, row int, column int, parent QModelIndex_ITF) bool {
	defer qt.Recovering("QAbstractListModel::dropMimeData")

	if ptr.Pointer() != nil {
		return C.QAbstractListModel_DropMimeData(ptr.Pointer(), PointerFromQMimeData(data), C.int(action), C.int(row), C.int(column), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QAbstractListModel) Flags(index QModelIndex_ITF) Qt__ItemFlag {
	defer qt.Recovering("QAbstractListModel::flags")

	if ptr.Pointer() != nil {
		return Qt__ItemFlag(C.QAbstractListModel_Flags(ptr.Pointer(), PointerFromQModelIndex(index)))
	}
	return 0
}

func (ptr *QAbstractListModel) Sibling(row int, column int, idx QModelIndex_ITF) *QModelIndex {
	defer qt.Recovering("QAbstractListModel::sibling")

	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QAbstractListModel_Sibling(ptr.Pointer(), C.int(row), C.int(column), PointerFromQModelIndex(idx)))
	}
	return nil
}

func (ptr *QAbstractListModel) DestroyQAbstractListModel() {
	defer qt.Recovering("QAbstractListModel::~QAbstractListModel")

	if ptr.Pointer() != nil {
		C.QAbstractListModel_DestroyQAbstractListModel(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractListModel) ConnectFetchMore(f func(parent *QModelIndex)) {
	defer qt.Recovering("connect QAbstractListModel::fetchMore")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "fetchMore", f)
	}
}

func (ptr *QAbstractListModel) DisconnectFetchMore() {
	defer qt.Recovering("disconnect QAbstractListModel::fetchMore")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "fetchMore")
	}
}

//export callbackQAbstractListModelFetchMore
func callbackQAbstractListModelFetchMore(ptrName *C.char, parent unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractListModel::fetchMore")

	if signal := qt.GetSignal(C.GoString(ptrName), "fetchMore"); signal != nil {
		signal.(func(*QModelIndex))(NewQModelIndexFromPointer(parent))
		return true
	}
	return false

}

func (ptr *QAbstractListModel) ConnectRevert(f func()) {
	defer qt.Recovering("connect QAbstractListModel::revert")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "revert", f)
	}
}

func (ptr *QAbstractListModel) DisconnectRevert() {
	defer qt.Recovering("disconnect QAbstractListModel::revert")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "revert")
	}
}

//export callbackQAbstractListModelRevert
func callbackQAbstractListModelRevert(ptrName *C.char) bool {
	defer qt.Recovering("callback QAbstractListModel::revert")

	if signal := qt.GetSignal(C.GoString(ptrName), "revert"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QAbstractListModel) ConnectSort(f func(column int, order Qt__SortOrder)) {
	defer qt.Recovering("connect QAbstractListModel::sort")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "sort", f)
	}
}

func (ptr *QAbstractListModel) DisconnectSort() {
	defer qt.Recovering("disconnect QAbstractListModel::sort")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "sort")
	}
}

//export callbackQAbstractListModelSort
func callbackQAbstractListModelSort(ptrName *C.char, column C.int, order C.int) bool {
	defer qt.Recovering("callback QAbstractListModel::sort")

	if signal := qt.GetSignal(C.GoString(ptrName), "sort"); signal != nil {
		signal.(func(int, Qt__SortOrder))(int(column), Qt__SortOrder(order))
		return true
	}
	return false

}

func (ptr *QAbstractListModel) ConnectTimerEvent(f func(event *QTimerEvent)) {
	defer qt.Recovering("connect QAbstractListModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAbstractListModel) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAbstractListModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAbstractListModelTimerEvent
func callbackQAbstractListModelTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractListModel::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*QTimerEvent))(NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractListModel) ConnectChildEvent(f func(event *QChildEvent)) {
	defer qt.Recovering("connect QAbstractListModel::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAbstractListModel) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAbstractListModel::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAbstractListModelChildEvent
func callbackQAbstractListModelChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractListModel::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*QChildEvent))(NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractListModel) ConnectCustomEvent(f func(event *QEvent)) {
	defer qt.Recovering("connect QAbstractListModel::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAbstractListModel) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAbstractListModel::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAbstractListModelCustomEvent
func callbackQAbstractListModelCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractListModel::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
		return true
	}
	return false

}
