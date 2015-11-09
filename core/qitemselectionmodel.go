package core

//#include "qitemselectionmodel.h"
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
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QItemSelectionModel_" + qt.RandomIdentifier())
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
	return NewQItemSelectionModelFromPointer(C.QItemSelectionModel_NewQItemSelectionModel(PointerFromQAbstractItemModel(model)))
}

func NewQItemSelectionModel2(model QAbstractItemModel_ITF, parent QObject_ITF) *QItemSelectionModel {
	return NewQItemSelectionModelFromPointer(C.QItemSelectionModel_NewQItemSelectionModel2(PointerFromQAbstractItemModel(model), PointerFromQObject(parent)))
}

func (ptr *QItemSelectionModel) Clear() {
	if ptr.Pointer() != nil {
		C.QItemSelectionModel_Clear(ptr.Pointer())
	}
}

func (ptr *QItemSelectionModel) ClearCurrentIndex() {
	if ptr.Pointer() != nil {
		C.QItemSelectionModel_ClearCurrentIndex(ptr.Pointer())
	}
}

func (ptr *QItemSelectionModel) ClearSelection() {
	if ptr.Pointer() != nil {
		C.QItemSelectionModel_ClearSelection(ptr.Pointer())
	}
}

func (ptr *QItemSelectionModel) ColumnIntersectsSelection(column int, parent QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QItemSelectionModel_ColumnIntersectsSelection(ptr.Pointer(), C.int(column), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QItemSelectionModel) ConnectCurrentChanged(f func(current *QModelIndex, previous *QModelIndex)) {
	if ptr.Pointer() != nil {
		C.QItemSelectionModel_ConnectCurrentChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentChanged", f)
	}
}

func (ptr *QItemSelectionModel) DisconnectCurrentChanged() {
	if ptr.Pointer() != nil {
		C.QItemSelectionModel_DisconnectCurrentChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentChanged")
	}
}

//export callbackQItemSelectionModelCurrentChanged
func callbackQItemSelectionModelCurrentChanged(ptrName *C.char, current unsafe.Pointer, previous unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "currentChanged").(func(*QModelIndex, *QModelIndex))(NewQModelIndexFromPointer(current), NewQModelIndexFromPointer(previous))
}

func (ptr *QItemSelectionModel) ConnectCurrentColumnChanged(f func(current *QModelIndex, previous *QModelIndex)) {
	if ptr.Pointer() != nil {
		C.QItemSelectionModel_ConnectCurrentColumnChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentColumnChanged", f)
	}
}

func (ptr *QItemSelectionModel) DisconnectCurrentColumnChanged() {
	if ptr.Pointer() != nil {
		C.QItemSelectionModel_DisconnectCurrentColumnChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentColumnChanged")
	}
}

//export callbackQItemSelectionModelCurrentColumnChanged
func callbackQItemSelectionModelCurrentColumnChanged(ptrName *C.char, current unsafe.Pointer, previous unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "currentColumnChanged").(func(*QModelIndex, *QModelIndex))(NewQModelIndexFromPointer(current), NewQModelIndexFromPointer(previous))
}

func (ptr *QItemSelectionModel) CurrentIndex() *QModelIndex {
	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QItemSelectionModel_CurrentIndex(ptr.Pointer()))
	}
	return nil
}

func (ptr *QItemSelectionModel) ConnectCurrentRowChanged(f func(current *QModelIndex, previous *QModelIndex)) {
	if ptr.Pointer() != nil {
		C.QItemSelectionModel_ConnectCurrentRowChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentRowChanged", f)
	}
}

func (ptr *QItemSelectionModel) DisconnectCurrentRowChanged() {
	if ptr.Pointer() != nil {
		C.QItemSelectionModel_DisconnectCurrentRowChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentRowChanged")
	}
}

//export callbackQItemSelectionModelCurrentRowChanged
func callbackQItemSelectionModelCurrentRowChanged(ptrName *C.char, current unsafe.Pointer, previous unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "currentRowChanged").(func(*QModelIndex, *QModelIndex))(NewQModelIndexFromPointer(current), NewQModelIndexFromPointer(previous))
}

func (ptr *QItemSelectionModel) HasSelection() bool {
	if ptr.Pointer() != nil {
		return C.QItemSelectionModel_HasSelection(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QItemSelectionModel) IsColumnSelected(column int, parent QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QItemSelectionModel_IsColumnSelected(ptr.Pointer(), C.int(column), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QItemSelectionModel) IsRowSelected(row int, parent QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QItemSelectionModel_IsRowSelected(ptr.Pointer(), C.int(row), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QItemSelectionModel) IsSelected(index QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QItemSelectionModel_IsSelected(ptr.Pointer(), PointerFromQModelIndex(index)) != 0
	}
	return false
}

func (ptr *QItemSelectionModel) Model2() *QAbstractItemModel {
	if ptr.Pointer() != nil {
		return NewQAbstractItemModelFromPointer(C.QItemSelectionModel_Model2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QItemSelectionModel) Model() *QAbstractItemModel {
	if ptr.Pointer() != nil {
		return NewQAbstractItemModelFromPointer(C.QItemSelectionModel_Model(ptr.Pointer()))
	}
	return nil
}

func (ptr *QItemSelectionModel) ConnectModelChanged(f func(model *QAbstractItemModel)) {
	if ptr.Pointer() != nil {
		C.QItemSelectionModel_ConnectModelChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "modelChanged", f)
	}
}

func (ptr *QItemSelectionModel) DisconnectModelChanged() {
	if ptr.Pointer() != nil {
		C.QItemSelectionModel_DisconnectModelChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "modelChanged")
	}
}

//export callbackQItemSelectionModelModelChanged
func callbackQItemSelectionModelModelChanged(ptrName *C.char, model unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "modelChanged").(func(*QAbstractItemModel))(NewQAbstractItemModelFromPointer(model))
}

func (ptr *QItemSelectionModel) Reset() {
	if ptr.Pointer() != nil {
		C.QItemSelectionModel_Reset(ptr.Pointer())
	}
}

func (ptr *QItemSelectionModel) RowIntersectsSelection(row int, parent QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QItemSelectionModel_RowIntersectsSelection(ptr.Pointer(), C.int(row), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QItemSelectionModel) Select2(selection QItemSelection_ITF, command QItemSelectionModel__SelectionFlag) {
	if ptr.Pointer() != nil {
		C.QItemSelectionModel_Select2(ptr.Pointer(), PointerFromQItemSelection(selection), C.int(command))
	}
}

func (ptr *QItemSelectionModel) Select(index QModelIndex_ITF, command QItemSelectionModel__SelectionFlag) {
	if ptr.Pointer() != nil {
		C.QItemSelectionModel_Select(ptr.Pointer(), PointerFromQModelIndex(index), C.int(command))
	}
}

func (ptr *QItemSelectionModel) SetCurrentIndex(index QModelIndex_ITF, command QItemSelectionModel__SelectionFlag) {
	if ptr.Pointer() != nil {
		C.QItemSelectionModel_SetCurrentIndex(ptr.Pointer(), PointerFromQModelIndex(index), C.int(command))
	}
}

func (ptr *QItemSelectionModel) SetModel(model QAbstractItemModel_ITF) {
	if ptr.Pointer() != nil {
		C.QItemSelectionModel_SetModel(ptr.Pointer(), PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QItemSelectionModel) DestroyQItemSelectionModel() {
	if ptr.Pointer() != nil {
		C.QItemSelectionModel_DestroyQItemSelectionModel(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
