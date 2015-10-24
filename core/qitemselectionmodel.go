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

type QItemSelectionModelITF interface {
	QObjectITF
	QItemSelectionModelPTR() *QItemSelectionModel
}

func PointerFromQItemSelectionModel(ptr QItemSelectionModelITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QItemSelectionModelPTR().Pointer()
	}
	return nil
}

func QItemSelectionModelFromPointer(ptr unsafe.Pointer) *QItemSelectionModel {
	var n = new(QItemSelectionModel)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QItemSelectionModel_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QItemSelectionModel) QItemSelectionModelPTR() *QItemSelectionModel {
	return ptr
}

//QItemSelectionModel::SelectionFlag
type QItemSelectionModel__SelectionFlag int

var (
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

func NewQItemSelectionModel(model QAbstractItemModelITF) *QItemSelectionModel {
	return QItemSelectionModelFromPointer(unsafe.Pointer(C.QItemSelectionModel_NewQItemSelectionModel(C.QtObjectPtr(PointerFromQAbstractItemModel(model)))))
}

func NewQItemSelectionModel2(model QAbstractItemModelITF, parent QObjectITF) *QItemSelectionModel {
	return QItemSelectionModelFromPointer(unsafe.Pointer(C.QItemSelectionModel_NewQItemSelectionModel2(C.QtObjectPtr(PointerFromQAbstractItemModel(model)), C.QtObjectPtr(PointerFromQObject(parent)))))
}

func (ptr *QItemSelectionModel) Clear() {
	if ptr.Pointer() != nil {
		C.QItemSelectionModel_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QItemSelectionModel) ClearCurrentIndex() {
	if ptr.Pointer() != nil {
		C.QItemSelectionModel_ClearCurrentIndex(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QItemSelectionModel) ClearSelection() {
	if ptr.Pointer() != nil {
		C.QItemSelectionModel_ClearSelection(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QItemSelectionModel) ColumnIntersectsSelection(column int, parent QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QItemSelectionModel_ColumnIntersectsSelection(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.QtObjectPtr(PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QItemSelectionModel) ConnectCurrentChanged(f func(current QModelIndexITF, previous QModelIndexITF)) {
	if ptr.Pointer() != nil {
		C.QItemSelectionModel_ConnectCurrentChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "currentChanged", f)
	}
}

func (ptr *QItemSelectionModel) DisconnectCurrentChanged() {
	if ptr.Pointer() != nil {
		C.QItemSelectionModel_DisconnectCurrentChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "currentChanged")
	}
}

//export callbackQItemSelectionModelCurrentChanged
func callbackQItemSelectionModelCurrentChanged(ptrName *C.char, current unsafe.Pointer, previous unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "currentChanged").(func(*QModelIndex, *QModelIndex))(QModelIndexFromPointer(current), QModelIndexFromPointer(previous))
}

func (ptr *QItemSelectionModel) ConnectCurrentColumnChanged(f func(current QModelIndexITF, previous QModelIndexITF)) {
	if ptr.Pointer() != nil {
		C.QItemSelectionModel_ConnectCurrentColumnChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "currentColumnChanged", f)
	}
}

func (ptr *QItemSelectionModel) DisconnectCurrentColumnChanged() {
	if ptr.Pointer() != nil {
		C.QItemSelectionModel_DisconnectCurrentColumnChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "currentColumnChanged")
	}
}

//export callbackQItemSelectionModelCurrentColumnChanged
func callbackQItemSelectionModelCurrentColumnChanged(ptrName *C.char, current unsafe.Pointer, previous unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "currentColumnChanged").(func(*QModelIndex, *QModelIndex))(QModelIndexFromPointer(current), QModelIndexFromPointer(previous))
}

func (ptr *QItemSelectionModel) CurrentIndex() *QModelIndex {
	if ptr.Pointer() != nil {
		return QModelIndexFromPointer(unsafe.Pointer(C.QItemSelectionModel_CurrentIndex(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QItemSelectionModel) ConnectCurrentRowChanged(f func(current QModelIndexITF, previous QModelIndexITF)) {
	if ptr.Pointer() != nil {
		C.QItemSelectionModel_ConnectCurrentRowChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "currentRowChanged", f)
	}
}

func (ptr *QItemSelectionModel) DisconnectCurrentRowChanged() {
	if ptr.Pointer() != nil {
		C.QItemSelectionModel_DisconnectCurrentRowChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "currentRowChanged")
	}
}

//export callbackQItemSelectionModelCurrentRowChanged
func callbackQItemSelectionModelCurrentRowChanged(ptrName *C.char, current unsafe.Pointer, previous unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "currentRowChanged").(func(*QModelIndex, *QModelIndex))(QModelIndexFromPointer(current), QModelIndexFromPointer(previous))
}

func (ptr *QItemSelectionModel) HasSelection() bool {
	if ptr.Pointer() != nil {
		return C.QItemSelectionModel_HasSelection(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QItemSelectionModel) IsColumnSelected(column int, parent QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QItemSelectionModel_IsColumnSelected(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.QtObjectPtr(PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QItemSelectionModel) IsRowSelected(row int, parent QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QItemSelectionModel_IsRowSelected(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.QtObjectPtr(PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QItemSelectionModel) IsSelected(index QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QItemSelectionModel_IsSelected(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(index))) != 0
	}
	return false
}

func (ptr *QItemSelectionModel) Model2() *QAbstractItemModel {
	if ptr.Pointer() != nil {
		return QAbstractItemModelFromPointer(unsafe.Pointer(C.QItemSelectionModel_Model2(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QItemSelectionModel) Model() *QAbstractItemModel {
	if ptr.Pointer() != nil {
		return QAbstractItemModelFromPointer(unsafe.Pointer(C.QItemSelectionModel_Model(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QItemSelectionModel) ConnectModelChanged(f func(model QAbstractItemModelITF)) {
	if ptr.Pointer() != nil {
		C.QItemSelectionModel_ConnectModelChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "modelChanged", f)
	}
}

func (ptr *QItemSelectionModel) DisconnectModelChanged() {
	if ptr.Pointer() != nil {
		C.QItemSelectionModel_DisconnectModelChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "modelChanged")
	}
}

//export callbackQItemSelectionModelModelChanged
func callbackQItemSelectionModelModelChanged(ptrName *C.char, model unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "modelChanged").(func(*QAbstractItemModel))(QAbstractItemModelFromPointer(model))
}

func (ptr *QItemSelectionModel) Reset() {
	if ptr.Pointer() != nil {
		C.QItemSelectionModel_Reset(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QItemSelectionModel) RowIntersectsSelection(row int, parent QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QItemSelectionModel_RowIntersectsSelection(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.QtObjectPtr(PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QItemSelectionModel) Select2(selection QItemSelectionITF, command QItemSelectionModel__SelectionFlag) {
	if ptr.Pointer() != nil {
		C.QItemSelectionModel_Select2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQItemSelection(selection)), C.int(command))
	}
}

func (ptr *QItemSelectionModel) Select(index QModelIndexITF, command QItemSelectionModel__SelectionFlag) {
	if ptr.Pointer() != nil {
		C.QItemSelectionModel_Select(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(index)), C.int(command))
	}
}

func (ptr *QItemSelectionModel) SetCurrentIndex(index QModelIndexITF, command QItemSelectionModel__SelectionFlag) {
	if ptr.Pointer() != nil {
		C.QItemSelectionModel_SetCurrentIndex(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(index)), C.int(command))
	}
}

func (ptr *QItemSelectionModel) SetModel(model QAbstractItemModelITF) {
	if ptr.Pointer() != nil {
		C.QItemSelectionModel_SetModel(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAbstractItemModel(model)))
	}
}

func (ptr *QItemSelectionModel) DestroyQItemSelectionModel() {
	if ptr.Pointer() != nil {
		C.QItemSelectionModel_DestroyQItemSelectionModel(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
