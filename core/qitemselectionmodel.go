package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QItemSelectionModel struct {
	QObject
}

type QItemSelectionModel_ITF interface {
	QObject_ITF
	QItemSelectionModel_PTR() *QItemSelectionModel
}

func PointerFromQItemSelectionModel(ptr QItemSelectionModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QItemSelectionModel_PTR().Pointer()
	}
	return nil
}

func NewQItemSelectionModelFromPointer(ptr unsafe.Pointer) *QItemSelectionModel {
	var n = new(QItemSelectionModel)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QItemSelectionModel_") {
		n.SetObjectName("QItemSelectionModel_" + qt.Identifier())
	}
	return n
}

func (ptr *QItemSelectionModel) QItemSelectionModel_PTR() *QItemSelectionModel {
	return ptr
}

//QItemSelectionModel::SelectionFlag
type QItemSelectionModel__SelectionFlag int64

const (
	QItemSelectionModel__NoUpdate       = QItemSelectionModel__SelectionFlag(0x0000)
	QItemSelectionModel__Clear          = QItemSelectionModel__SelectionFlag(0x0001)
	QItemSelectionModel__Select         = QItemSelectionModel__SelectionFlag(0x0002)
	QItemSelectionModel__Deselect       = QItemSelectionModel__SelectionFlag(0x0004)
	QItemSelectionModel__Toggle         = QItemSelectionModel__SelectionFlag(0x0008)
	QItemSelectionModel__Current        = QItemSelectionModel__SelectionFlag(0x0010)
	QItemSelectionModel__Rows           = QItemSelectionModel__SelectionFlag(0x0020)
	QItemSelectionModel__Columns        = QItemSelectionModel__SelectionFlag(0x0040)
	QItemSelectionModel__SelectCurrent  = QItemSelectionModel__SelectionFlag(QItemSelectionModel__Select | QItemSelectionModel__Current)
	QItemSelectionModel__ToggleCurrent  = QItemSelectionModel__SelectionFlag(QItemSelectionModel__Toggle | QItemSelectionModel__Current)
	QItemSelectionModel__ClearAndSelect = QItemSelectionModel__SelectionFlag(QItemSelectionModel__Clear | QItemSelectionModel__Select)
)

func NewQItemSelectionModel(model QAbstractItemModel_ITF) *QItemSelectionModel {
	defer qt.Recovering("QItemSelectionModel::QItemSelectionModel")

	return NewQItemSelectionModelFromPointer(C.QItemSelectionModel_NewQItemSelectionModel(PointerFromQAbstractItemModel(model)))
}

func NewQItemSelectionModel2(model QAbstractItemModel_ITF, parent QObject_ITF) *QItemSelectionModel {
	defer qt.Recovering("QItemSelectionModel::QItemSelectionModel")

	return NewQItemSelectionModelFromPointer(C.QItemSelectionModel_NewQItemSelectionModel2(PointerFromQAbstractItemModel(model), PointerFromQObject(parent)))
}

func (ptr *QItemSelectionModel) ConnectClear(f func()) {
	defer qt.Recovering("connect QItemSelectionModel::clear")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "clear", f)
	}
}

func (ptr *QItemSelectionModel) DisconnectClear() {
	defer qt.Recovering("disconnect QItemSelectionModel::clear")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "clear")
	}
}

//export callbackQItemSelectionModelClear
func callbackQItemSelectionModelClear(ptrName *C.char) bool {
	defer qt.Recovering("callback QItemSelectionModel::clear")

	if signal := qt.GetSignal(C.GoString(ptrName), "clear"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QItemSelectionModel) ConnectClearCurrentIndex(f func()) {
	defer qt.Recovering("connect QItemSelectionModel::clearCurrentIndex")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "clearCurrentIndex", f)
	}
}

func (ptr *QItemSelectionModel) DisconnectClearCurrentIndex() {
	defer qt.Recovering("disconnect QItemSelectionModel::clearCurrentIndex")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "clearCurrentIndex")
	}
}

//export callbackQItemSelectionModelClearCurrentIndex
func callbackQItemSelectionModelClearCurrentIndex(ptrName *C.char) bool {
	defer qt.Recovering("callback QItemSelectionModel::clearCurrentIndex")

	if signal := qt.GetSignal(C.GoString(ptrName), "clearCurrentIndex"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QItemSelectionModel) ClearSelection() {
	defer qt.Recovering("QItemSelectionModel::clearSelection")

	if ptr.Pointer() != nil {
		C.QItemSelectionModel_ClearSelection(ptr.Pointer())
	}
}

func (ptr *QItemSelectionModel) ColumnIntersectsSelection(column int, parent QModelIndex_ITF) bool {
	defer qt.Recovering("QItemSelectionModel::columnIntersectsSelection")

	if ptr.Pointer() != nil {
		return C.QItemSelectionModel_ColumnIntersectsSelection(ptr.Pointer(), C.int(column), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QItemSelectionModel) ConnectCurrentChanged(f func(current *QModelIndex, previous *QModelIndex)) {
	defer qt.Recovering("connect QItemSelectionModel::currentChanged")

	if ptr.Pointer() != nil {
		C.QItemSelectionModel_ConnectCurrentChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentChanged", f)
	}
}

func (ptr *QItemSelectionModel) DisconnectCurrentChanged() {
	defer qt.Recovering("disconnect QItemSelectionModel::currentChanged")

	if ptr.Pointer() != nil {
		C.QItemSelectionModel_DisconnectCurrentChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentChanged")
	}
}

//export callbackQItemSelectionModelCurrentChanged
func callbackQItemSelectionModelCurrentChanged(ptrName *C.char, current unsafe.Pointer, previous unsafe.Pointer) {
	defer qt.Recovering("callback QItemSelectionModel::currentChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "currentChanged"); signal != nil {
		signal.(func(*QModelIndex, *QModelIndex))(NewQModelIndexFromPointer(current), NewQModelIndexFromPointer(previous))
	}

}

func (ptr *QItemSelectionModel) ConnectCurrentColumnChanged(f func(current *QModelIndex, previous *QModelIndex)) {
	defer qt.Recovering("connect QItemSelectionModel::currentColumnChanged")

	if ptr.Pointer() != nil {
		C.QItemSelectionModel_ConnectCurrentColumnChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentColumnChanged", f)
	}
}

func (ptr *QItemSelectionModel) DisconnectCurrentColumnChanged() {
	defer qt.Recovering("disconnect QItemSelectionModel::currentColumnChanged")

	if ptr.Pointer() != nil {
		C.QItemSelectionModel_DisconnectCurrentColumnChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentColumnChanged")
	}
}

//export callbackQItemSelectionModelCurrentColumnChanged
func callbackQItemSelectionModelCurrentColumnChanged(ptrName *C.char, current unsafe.Pointer, previous unsafe.Pointer) {
	defer qt.Recovering("callback QItemSelectionModel::currentColumnChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "currentColumnChanged"); signal != nil {
		signal.(func(*QModelIndex, *QModelIndex))(NewQModelIndexFromPointer(current), NewQModelIndexFromPointer(previous))
	}

}

func (ptr *QItemSelectionModel) CurrentIndex() *QModelIndex {
	defer qt.Recovering("QItemSelectionModel::currentIndex")

	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QItemSelectionModel_CurrentIndex(ptr.Pointer()))
	}
	return nil
}

func (ptr *QItemSelectionModel) ConnectCurrentRowChanged(f func(current *QModelIndex, previous *QModelIndex)) {
	defer qt.Recovering("connect QItemSelectionModel::currentRowChanged")

	if ptr.Pointer() != nil {
		C.QItemSelectionModel_ConnectCurrentRowChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentRowChanged", f)
	}
}

func (ptr *QItemSelectionModel) DisconnectCurrentRowChanged() {
	defer qt.Recovering("disconnect QItemSelectionModel::currentRowChanged")

	if ptr.Pointer() != nil {
		C.QItemSelectionModel_DisconnectCurrentRowChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentRowChanged")
	}
}

//export callbackQItemSelectionModelCurrentRowChanged
func callbackQItemSelectionModelCurrentRowChanged(ptrName *C.char, current unsafe.Pointer, previous unsafe.Pointer) {
	defer qt.Recovering("callback QItemSelectionModel::currentRowChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "currentRowChanged"); signal != nil {
		signal.(func(*QModelIndex, *QModelIndex))(NewQModelIndexFromPointer(current), NewQModelIndexFromPointer(previous))
	}

}

func (ptr *QItemSelectionModel) HasSelection() bool {
	defer qt.Recovering("QItemSelectionModel::hasSelection")

	if ptr.Pointer() != nil {
		return C.QItemSelectionModel_HasSelection(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QItemSelectionModel) IsColumnSelected(column int, parent QModelIndex_ITF) bool {
	defer qt.Recovering("QItemSelectionModel::isColumnSelected")

	if ptr.Pointer() != nil {
		return C.QItemSelectionModel_IsColumnSelected(ptr.Pointer(), C.int(column), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QItemSelectionModel) IsRowSelected(row int, parent QModelIndex_ITF) bool {
	defer qt.Recovering("QItemSelectionModel::isRowSelected")

	if ptr.Pointer() != nil {
		return C.QItemSelectionModel_IsRowSelected(ptr.Pointer(), C.int(row), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QItemSelectionModel) IsSelected(index QModelIndex_ITF) bool {
	defer qt.Recovering("QItemSelectionModel::isSelected")

	if ptr.Pointer() != nil {
		return C.QItemSelectionModel_IsSelected(ptr.Pointer(), PointerFromQModelIndex(index)) != 0
	}
	return false
}

func (ptr *QItemSelectionModel) Model2() *QAbstractItemModel {
	defer qt.Recovering("QItemSelectionModel::model")

	if ptr.Pointer() != nil {
		return NewQAbstractItemModelFromPointer(C.QItemSelectionModel_Model2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QItemSelectionModel) Model() *QAbstractItemModel {
	defer qt.Recovering("QItemSelectionModel::model")

	if ptr.Pointer() != nil {
		return NewQAbstractItemModelFromPointer(C.QItemSelectionModel_Model(ptr.Pointer()))
	}
	return nil
}

func (ptr *QItemSelectionModel) ConnectModelChanged(f func(model *QAbstractItemModel)) {
	defer qt.Recovering("connect QItemSelectionModel::modelChanged")

	if ptr.Pointer() != nil {
		C.QItemSelectionModel_ConnectModelChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "modelChanged", f)
	}
}

func (ptr *QItemSelectionModel) DisconnectModelChanged() {
	defer qt.Recovering("disconnect QItemSelectionModel::modelChanged")

	if ptr.Pointer() != nil {
		C.QItemSelectionModel_DisconnectModelChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "modelChanged")
	}
}

//export callbackQItemSelectionModelModelChanged
func callbackQItemSelectionModelModelChanged(ptrName *C.char, model unsafe.Pointer) {
	defer qt.Recovering("callback QItemSelectionModel::modelChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "modelChanged"); signal != nil {
		signal.(func(*QAbstractItemModel))(NewQAbstractItemModelFromPointer(model))
	}

}

func (ptr *QItemSelectionModel) ConnectReset(f func()) {
	defer qt.Recovering("connect QItemSelectionModel::reset")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "reset", f)
	}
}

func (ptr *QItemSelectionModel) DisconnectReset() {
	defer qt.Recovering("disconnect QItemSelectionModel::reset")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "reset")
	}
}

//export callbackQItemSelectionModelReset
func callbackQItemSelectionModelReset(ptrName *C.char) bool {
	defer qt.Recovering("callback QItemSelectionModel::reset")

	if signal := qt.GetSignal(C.GoString(ptrName), "reset"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QItemSelectionModel) RowIntersectsSelection(row int, parent QModelIndex_ITF) bool {
	defer qt.Recovering("QItemSelectionModel::rowIntersectsSelection")

	if ptr.Pointer() != nil {
		return C.QItemSelectionModel_RowIntersectsSelection(ptr.Pointer(), C.int(row), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QItemSelectionModel) ConnectSelect(f func(index *QModelIndex, command QItemSelectionModel__SelectionFlag)) {
	defer qt.Recovering("connect QItemSelectionModel::select")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "select", f)
	}
}

func (ptr *QItemSelectionModel) DisconnectSelect() {
	defer qt.Recovering("disconnect QItemSelectionModel::select")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "select")
	}
}

//export callbackQItemSelectionModelSelect
func callbackQItemSelectionModelSelect(ptrName *C.char, index unsafe.Pointer, command C.int) bool {
	defer qt.Recovering("callback QItemSelectionModel::select")

	if signal := qt.GetSignal(C.GoString(ptrName), "select"); signal != nil {
		signal.(func(*QModelIndex, QItemSelectionModel__SelectionFlag))(NewQModelIndexFromPointer(index), QItemSelectionModel__SelectionFlag(command))
		return true
	}
	return false

}

func (ptr *QItemSelectionModel) ConnectSetCurrentIndex(f func(index *QModelIndex, command QItemSelectionModel__SelectionFlag)) {
	defer qt.Recovering("connect QItemSelectionModel::setCurrentIndex")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setCurrentIndex", f)
	}
}

func (ptr *QItemSelectionModel) DisconnectSetCurrentIndex() {
	defer qt.Recovering("disconnect QItemSelectionModel::setCurrentIndex")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setCurrentIndex")
	}
}

//export callbackQItemSelectionModelSetCurrentIndex
func callbackQItemSelectionModelSetCurrentIndex(ptrName *C.char, index unsafe.Pointer, command C.int) bool {
	defer qt.Recovering("callback QItemSelectionModel::setCurrentIndex")

	if signal := qt.GetSignal(C.GoString(ptrName), "setCurrentIndex"); signal != nil {
		signal.(func(*QModelIndex, QItemSelectionModel__SelectionFlag))(NewQModelIndexFromPointer(index), QItemSelectionModel__SelectionFlag(command))
		return true
	}
	return false

}

func (ptr *QItemSelectionModel) SetModel(model QAbstractItemModel_ITF) {
	defer qt.Recovering("QItemSelectionModel::setModel")

	if ptr.Pointer() != nil {
		C.QItemSelectionModel_SetModel(ptr.Pointer(), PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QItemSelectionModel) DestroyQItemSelectionModel() {
	defer qt.Recovering("QItemSelectionModel::~QItemSelectionModel")

	if ptr.Pointer() != nil {
		C.QItemSelectionModel_DestroyQItemSelectionModel(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QItemSelectionModel) ConnectTimerEvent(f func(event *QTimerEvent)) {
	defer qt.Recovering("connect QItemSelectionModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QItemSelectionModel) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QItemSelectionModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQItemSelectionModelTimerEvent
func callbackQItemSelectionModelTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QItemSelectionModel::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*QTimerEvent))(NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QItemSelectionModel) ConnectChildEvent(f func(event *QChildEvent)) {
	defer qt.Recovering("connect QItemSelectionModel::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QItemSelectionModel) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QItemSelectionModel::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQItemSelectionModelChildEvent
func callbackQItemSelectionModelChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QItemSelectionModel::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*QChildEvent))(NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QItemSelectionModel) ConnectCustomEvent(f func(event *QEvent)) {
	defer qt.Recovering("connect QItemSelectionModel::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QItemSelectionModel) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QItemSelectionModel::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQItemSelectionModelCustomEvent
func callbackQItemSelectionModelCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QItemSelectionModel::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
		return true
	}
	return false

}
