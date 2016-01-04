package help

//#include "help.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QHelpIndexModel struct {
	core.QStringListModel
}

type QHelpIndexModel_ITF interface {
	core.QStringListModel_ITF
	QHelpIndexModel_PTR() *QHelpIndexModel
}

func PointerFromQHelpIndexModel(ptr QHelpIndexModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpIndexModel_PTR().Pointer()
	}
	return nil
}

func NewQHelpIndexModelFromPointer(ptr unsafe.Pointer) *QHelpIndexModel {
	var n = new(QHelpIndexModel)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QHelpIndexModel_") {
		n.SetObjectName("QHelpIndexModel_" + qt.Identifier())
	}
	return n
}

func (ptr *QHelpIndexModel) QHelpIndexModel_PTR() *QHelpIndexModel {
	return ptr
}

func (ptr *QHelpIndexModel) CreateIndex(customFilterName string) {
	defer qt.Recovering("QHelpIndexModel::createIndex")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_CreateIndex(ptr.Pointer(), C.CString(customFilterName))
	}
}

func (ptr *QHelpIndexModel) Filter(filter string, wildcard string) *core.QModelIndex {
	defer qt.Recovering("QHelpIndexModel::filter")

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QHelpIndexModel_Filter(ptr.Pointer(), C.CString(filter), C.CString(wildcard)))
	}
	return nil
}

func (ptr *QHelpIndexModel) ConnectIndexCreated(f func()) {
	defer qt.Recovering("connect QHelpIndexModel::indexCreated")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_ConnectIndexCreated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "indexCreated", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectIndexCreated() {
	defer qt.Recovering("disconnect QHelpIndexModel::indexCreated")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_DisconnectIndexCreated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "indexCreated")
	}
}

//export callbackQHelpIndexModelIndexCreated
func callbackQHelpIndexModelIndexCreated(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QHelpIndexModel::indexCreated")

	if signal := qt.GetSignal(C.GoString(ptrName), "indexCreated"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpIndexModel) IndexCreated() {
	defer qt.Recovering("QHelpIndexModel::indexCreated")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_IndexCreated(ptr.Pointer())
	}
}

func (ptr *QHelpIndexModel) ConnectIndexCreationStarted(f func()) {
	defer qt.Recovering("connect QHelpIndexModel::indexCreationStarted")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_ConnectIndexCreationStarted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "indexCreationStarted", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectIndexCreationStarted() {
	defer qt.Recovering("disconnect QHelpIndexModel::indexCreationStarted")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_DisconnectIndexCreationStarted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "indexCreationStarted")
	}
}

//export callbackQHelpIndexModelIndexCreationStarted
func callbackQHelpIndexModelIndexCreationStarted(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QHelpIndexModel::indexCreationStarted")

	if signal := qt.GetSignal(C.GoString(ptrName), "indexCreationStarted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpIndexModel) IndexCreationStarted() {
	defer qt.Recovering("QHelpIndexModel::indexCreationStarted")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_IndexCreationStarted(ptr.Pointer())
	}
}

func (ptr *QHelpIndexModel) IsCreatingIndex() bool {
	defer qt.Recovering("QHelpIndexModel::isCreatingIndex")

	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_IsCreatingIndex(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHelpIndexModel) ConnectSort(f func(column int, order core.Qt__SortOrder)) {
	defer qt.Recovering("connect QHelpIndexModel::sort")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "sort", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectSort() {
	defer qt.Recovering("disconnect QHelpIndexModel::sort")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "sort")
	}
}

//export callbackQHelpIndexModelSort
func callbackQHelpIndexModelSort(ptr unsafe.Pointer, ptrName *C.char, column C.int, order C.int) {
	defer qt.Recovering("callback QHelpIndexModel::sort")

	if signal := qt.GetSignal(C.GoString(ptrName), "sort"); signal != nil {
		signal.(func(int, core.Qt__SortOrder))(int(column), core.Qt__SortOrder(order))
	} else {
		NewQHelpIndexModelFromPointer(ptr).SortDefault(int(column), core.Qt__SortOrder(order))
	}
}

func (ptr *QHelpIndexModel) Sort(column int, order core.Qt__SortOrder) {
	defer qt.Recovering("QHelpIndexModel::sort")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_Sort(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QHelpIndexModel) SortDefault(column int, order core.Qt__SortOrder) {
	defer qt.Recovering("QHelpIndexModel::sort")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_SortDefault(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QHelpIndexModel) ConnectFetchMore(f func(parent *core.QModelIndex)) {
	defer qt.Recovering("connect QHelpIndexModel::fetchMore")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "fetchMore", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectFetchMore() {
	defer qt.Recovering("disconnect QHelpIndexModel::fetchMore")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "fetchMore")
	}
}

//export callbackQHelpIndexModelFetchMore
func callbackQHelpIndexModelFetchMore(ptr unsafe.Pointer, ptrName *C.char, parent unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexModel::fetchMore")

	if signal := qt.GetSignal(C.GoString(ptrName), "fetchMore"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(parent))
	} else {
		NewQHelpIndexModelFromPointer(ptr).FetchMoreDefault(core.NewQModelIndexFromPointer(parent))
	}
}

func (ptr *QHelpIndexModel) FetchMore(parent core.QModelIndex_ITF) {
	defer qt.Recovering("QHelpIndexModel::fetchMore")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_FetchMore(ptr.Pointer(), core.PointerFromQModelIndex(parent))
	}
}

func (ptr *QHelpIndexModel) FetchMoreDefault(parent core.QModelIndex_ITF) {
	defer qt.Recovering("QHelpIndexModel::fetchMore")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_FetchMoreDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent))
	}
}

func (ptr *QHelpIndexModel) ConnectRevert(f func()) {
	defer qt.Recovering("connect QHelpIndexModel::revert")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "revert", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectRevert() {
	defer qt.Recovering("disconnect QHelpIndexModel::revert")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "revert")
	}
}

//export callbackQHelpIndexModelRevert
func callbackQHelpIndexModelRevert(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QHelpIndexModel::revert")

	if signal := qt.GetSignal(C.GoString(ptrName), "revert"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QHelpIndexModel) Revert() {
	defer qt.Recovering("QHelpIndexModel::revert")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_Revert(ptr.Pointer())
	}
}

func (ptr *QHelpIndexModel) RevertDefault() {
	defer qt.Recovering("QHelpIndexModel::revert")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_RevertDefault(ptr.Pointer())
	}
}

func (ptr *QHelpIndexModel) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QHelpIndexModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QHelpIndexModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQHelpIndexModelTimerEvent
func callbackQHelpIndexModelTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexModel::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHelpIndexModelFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHelpIndexModel) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHelpIndexModel::timerEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHelpIndexModel) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHelpIndexModel::timerEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHelpIndexModel) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QHelpIndexModel::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QHelpIndexModel::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQHelpIndexModelChildEvent
func callbackQHelpIndexModelChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexModel::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHelpIndexModelFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHelpIndexModel) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHelpIndexModel::childEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHelpIndexModel) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHelpIndexModel::childEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHelpIndexModel) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpIndexModel::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QHelpIndexModel::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQHelpIndexModelCustomEvent
func callbackQHelpIndexModelCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexModel::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpIndexModelFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpIndexModel) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpIndexModel::customEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpIndexModel) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpIndexModel::customEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
