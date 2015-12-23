package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QAbstractTableModel struct {
	QAbstractItemModel
}

type QAbstractTableModel_ITF interface {
	QAbstractItemModel_ITF
	QAbstractTableModel_PTR() *QAbstractTableModel
}

func PointerFromQAbstractTableModel(ptr QAbstractTableModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractTableModel_PTR().Pointer()
	}
	return nil
}

func NewQAbstractTableModelFromPointer(ptr unsafe.Pointer) *QAbstractTableModel {
	var n = new(QAbstractTableModel)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAbstractTableModel_") {
		n.SetObjectName("QAbstractTableModel_" + qt.Identifier())
	}
	return n
}

func (ptr *QAbstractTableModel) QAbstractTableModel_PTR() *QAbstractTableModel {
	return ptr
}

func (ptr *QAbstractTableModel) Index(row int, column int, parent QModelIndex_ITF) *QModelIndex {
	defer qt.Recovering("QAbstractTableModel::index")

	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QAbstractTableModel_Index(ptr.Pointer(), C.int(row), C.int(column), PointerFromQModelIndex(parent)))
	}
	return nil
}

func (ptr *QAbstractTableModel) DropMimeData(data QMimeData_ITF, action Qt__DropAction, row int, column int, parent QModelIndex_ITF) bool {
	defer qt.Recovering("QAbstractTableModel::dropMimeData")

	if ptr.Pointer() != nil {
		return C.QAbstractTableModel_DropMimeData(ptr.Pointer(), PointerFromQMimeData(data), C.int(action), C.int(row), C.int(column), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QAbstractTableModel) Flags(index QModelIndex_ITF) Qt__ItemFlag {
	defer qt.Recovering("QAbstractTableModel::flags")

	if ptr.Pointer() != nil {
		return Qt__ItemFlag(C.QAbstractTableModel_Flags(ptr.Pointer(), PointerFromQModelIndex(index)))
	}
	return 0
}

func (ptr *QAbstractTableModel) Sibling(row int, column int, idx QModelIndex_ITF) *QModelIndex {
	defer qt.Recovering("QAbstractTableModel::sibling")

	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QAbstractTableModel_Sibling(ptr.Pointer(), C.int(row), C.int(column), PointerFromQModelIndex(idx)))
	}
	return nil
}

func (ptr *QAbstractTableModel) DestroyQAbstractTableModel() {
	defer qt.Recovering("QAbstractTableModel::~QAbstractTableModel")

	if ptr.Pointer() != nil {
		C.QAbstractTableModel_DestroyQAbstractTableModel(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractTableModel) ConnectFetchMore(f func(parent *QModelIndex)) {
	defer qt.Recovering("connect QAbstractTableModel::fetchMore")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "fetchMore", f)
	}
}

func (ptr *QAbstractTableModel) DisconnectFetchMore() {
	defer qt.Recovering("disconnect QAbstractTableModel::fetchMore")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "fetchMore")
	}
}

//export callbackQAbstractTableModelFetchMore
func callbackQAbstractTableModelFetchMore(ptrName *C.char, parent unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractTableModel::fetchMore")

	if signal := qt.GetSignal(C.GoString(ptrName), "fetchMore"); signal != nil {
		signal.(func(*QModelIndex))(NewQModelIndexFromPointer(parent))
		return true
	}
	return false

}

func (ptr *QAbstractTableModel) ConnectRevert(f func()) {
	defer qt.Recovering("connect QAbstractTableModel::revert")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "revert", f)
	}
}

func (ptr *QAbstractTableModel) DisconnectRevert() {
	defer qt.Recovering("disconnect QAbstractTableModel::revert")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "revert")
	}
}

//export callbackQAbstractTableModelRevert
func callbackQAbstractTableModelRevert(ptrName *C.char) bool {
	defer qt.Recovering("callback QAbstractTableModel::revert")

	if signal := qt.GetSignal(C.GoString(ptrName), "revert"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QAbstractTableModel) ConnectSort(f func(column int, order Qt__SortOrder)) {
	defer qt.Recovering("connect QAbstractTableModel::sort")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "sort", f)
	}
}

func (ptr *QAbstractTableModel) DisconnectSort() {
	defer qt.Recovering("disconnect QAbstractTableModel::sort")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "sort")
	}
}

//export callbackQAbstractTableModelSort
func callbackQAbstractTableModelSort(ptrName *C.char, column C.int, order C.int) bool {
	defer qt.Recovering("callback QAbstractTableModel::sort")

	if signal := qt.GetSignal(C.GoString(ptrName), "sort"); signal != nil {
		signal.(func(int, Qt__SortOrder))(int(column), Qt__SortOrder(order))
		return true
	}
	return false

}

func (ptr *QAbstractTableModel) ConnectTimerEvent(f func(event *QTimerEvent)) {
	defer qt.Recovering("connect QAbstractTableModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAbstractTableModel) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAbstractTableModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAbstractTableModelTimerEvent
func callbackQAbstractTableModelTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractTableModel::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*QTimerEvent))(NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractTableModel) ConnectChildEvent(f func(event *QChildEvent)) {
	defer qt.Recovering("connect QAbstractTableModel::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAbstractTableModel) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAbstractTableModel::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAbstractTableModelChildEvent
func callbackQAbstractTableModelChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractTableModel::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*QChildEvent))(NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractTableModel) ConnectCustomEvent(f func(event *QEvent)) {
	defer qt.Recovering("connect QAbstractTableModel::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAbstractTableModel) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAbstractTableModel::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAbstractTableModelCustomEvent
func callbackQAbstractTableModelCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractTableModel::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
		return true
	}
	return false

}
