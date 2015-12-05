package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QItemSelectionModel::QItemSelectionModel")
		}
	}()

	return NewQItemSelectionModelFromPointer(C.QItemSelectionModel_NewQItemSelectionModel(PointerFromQAbstractItemModel(model)))
}

func NewQItemSelectionModel2(model QAbstractItemModel_ITF, parent QObject_ITF) *QItemSelectionModel {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QItemSelectionModel::QItemSelectionModel")
		}
	}()

	return NewQItemSelectionModelFromPointer(C.QItemSelectionModel_NewQItemSelectionModel2(PointerFromQAbstractItemModel(model), PointerFromQObject(parent)))
}

func (ptr *QItemSelectionModel) Clear() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QItemSelectionModel::clear")
		}
	}()

	if ptr.Pointer() != nil {
		C.QItemSelectionModel_Clear(ptr.Pointer())
	}
}

func (ptr *QItemSelectionModel) ClearCurrentIndex() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QItemSelectionModel::clearCurrentIndex")
		}
	}()

	if ptr.Pointer() != nil {
		C.QItemSelectionModel_ClearCurrentIndex(ptr.Pointer())
	}
}

func (ptr *QItemSelectionModel) ClearSelection() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QItemSelectionModel::clearSelection")
		}
	}()

	if ptr.Pointer() != nil {
		C.QItemSelectionModel_ClearSelection(ptr.Pointer())
	}
}

func (ptr *QItemSelectionModel) ColumnIntersectsSelection(column int, parent QModelIndex_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QItemSelectionModel::columnIntersectsSelection")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QItemSelectionModel_ColumnIntersectsSelection(ptr.Pointer(), C.int(column), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QItemSelectionModel) ConnectCurrentChanged(f func(current *QModelIndex, previous *QModelIndex)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QItemSelectionModel::currentChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QItemSelectionModel_ConnectCurrentChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentChanged", f)
	}
}

func (ptr *QItemSelectionModel) DisconnectCurrentChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QItemSelectionModel::currentChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QItemSelectionModel_DisconnectCurrentChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentChanged")
	}
}

//export callbackQItemSelectionModelCurrentChanged
func callbackQItemSelectionModelCurrentChanged(ptrName *C.char, current unsafe.Pointer, previous unsafe.Pointer) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QItemSelectionModel::currentChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "currentChanged").(func(*QModelIndex, *QModelIndex))(NewQModelIndexFromPointer(current), NewQModelIndexFromPointer(previous))
}

func (ptr *QItemSelectionModel) ConnectCurrentColumnChanged(f func(current *QModelIndex, previous *QModelIndex)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QItemSelectionModel::currentColumnChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QItemSelectionModel_ConnectCurrentColumnChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentColumnChanged", f)
	}
}

func (ptr *QItemSelectionModel) DisconnectCurrentColumnChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QItemSelectionModel::currentColumnChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QItemSelectionModel_DisconnectCurrentColumnChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentColumnChanged")
	}
}

//export callbackQItemSelectionModelCurrentColumnChanged
func callbackQItemSelectionModelCurrentColumnChanged(ptrName *C.char, current unsafe.Pointer, previous unsafe.Pointer) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QItemSelectionModel::currentColumnChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "currentColumnChanged").(func(*QModelIndex, *QModelIndex))(NewQModelIndexFromPointer(current), NewQModelIndexFromPointer(previous))
}

func (ptr *QItemSelectionModel) CurrentIndex() *QModelIndex {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QItemSelectionModel::currentIndex")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QItemSelectionModel_CurrentIndex(ptr.Pointer()))
	}
	return nil
}

func (ptr *QItemSelectionModel) ConnectCurrentRowChanged(f func(current *QModelIndex, previous *QModelIndex)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QItemSelectionModel::currentRowChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QItemSelectionModel_ConnectCurrentRowChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentRowChanged", f)
	}
}

func (ptr *QItemSelectionModel) DisconnectCurrentRowChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QItemSelectionModel::currentRowChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QItemSelectionModel_DisconnectCurrentRowChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentRowChanged")
	}
}

//export callbackQItemSelectionModelCurrentRowChanged
func callbackQItemSelectionModelCurrentRowChanged(ptrName *C.char, current unsafe.Pointer, previous unsafe.Pointer) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QItemSelectionModel::currentRowChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "currentRowChanged").(func(*QModelIndex, *QModelIndex))(NewQModelIndexFromPointer(current), NewQModelIndexFromPointer(previous))
}

func (ptr *QItemSelectionModel) HasSelection() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QItemSelectionModel::hasSelection")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QItemSelectionModel_HasSelection(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QItemSelectionModel) IsColumnSelected(column int, parent QModelIndex_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QItemSelectionModel::isColumnSelected")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QItemSelectionModel_IsColumnSelected(ptr.Pointer(), C.int(column), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QItemSelectionModel) IsRowSelected(row int, parent QModelIndex_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QItemSelectionModel::isRowSelected")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QItemSelectionModel_IsRowSelected(ptr.Pointer(), C.int(row), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QItemSelectionModel) IsSelected(index QModelIndex_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QItemSelectionModel::isSelected")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QItemSelectionModel_IsSelected(ptr.Pointer(), PointerFromQModelIndex(index)) != 0
	}
	return false
}

func (ptr *QItemSelectionModel) Model2() *QAbstractItemModel {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QItemSelectionModel::model")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQAbstractItemModelFromPointer(C.QItemSelectionModel_Model2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QItemSelectionModel) Model() *QAbstractItemModel {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QItemSelectionModel::model")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQAbstractItemModelFromPointer(C.QItemSelectionModel_Model(ptr.Pointer()))
	}
	return nil
}

func (ptr *QItemSelectionModel) ConnectModelChanged(f func(model *QAbstractItemModel)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QItemSelectionModel::modelChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QItemSelectionModel_ConnectModelChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "modelChanged", f)
	}
}

func (ptr *QItemSelectionModel) DisconnectModelChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QItemSelectionModel::modelChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QItemSelectionModel_DisconnectModelChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "modelChanged")
	}
}

//export callbackQItemSelectionModelModelChanged
func callbackQItemSelectionModelModelChanged(ptrName *C.char, model unsafe.Pointer) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QItemSelectionModel::modelChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "modelChanged").(func(*QAbstractItemModel))(NewQAbstractItemModelFromPointer(model))
}

func (ptr *QItemSelectionModel) Reset() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QItemSelectionModel::reset")
		}
	}()

	if ptr.Pointer() != nil {
		C.QItemSelectionModel_Reset(ptr.Pointer())
	}
}

func (ptr *QItemSelectionModel) RowIntersectsSelection(row int, parent QModelIndex_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QItemSelectionModel::rowIntersectsSelection")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QItemSelectionModel_RowIntersectsSelection(ptr.Pointer(), C.int(row), PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QItemSelectionModel) Select2(selection QItemSelection_ITF, command QItemSelectionModel__SelectionFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QItemSelectionModel::select")
		}
	}()

	if ptr.Pointer() != nil {
		C.QItemSelectionModel_Select2(ptr.Pointer(), PointerFromQItemSelection(selection), C.int(command))
	}
}

func (ptr *QItemSelectionModel) Select(index QModelIndex_ITF, command QItemSelectionModel__SelectionFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QItemSelectionModel::select")
		}
	}()

	if ptr.Pointer() != nil {
		C.QItemSelectionModel_Select(ptr.Pointer(), PointerFromQModelIndex(index), C.int(command))
	}
}

func (ptr *QItemSelectionModel) SetCurrentIndex(index QModelIndex_ITF, command QItemSelectionModel__SelectionFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QItemSelectionModel::setCurrentIndex")
		}
	}()

	if ptr.Pointer() != nil {
		C.QItemSelectionModel_SetCurrentIndex(ptr.Pointer(), PointerFromQModelIndex(index), C.int(command))
	}
}

func (ptr *QItemSelectionModel) SetModel(model QAbstractItemModel_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QItemSelectionModel::setModel")
		}
	}()

	if ptr.Pointer() != nil {
		C.QItemSelectionModel_SetModel(ptr.Pointer(), PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QItemSelectionModel) DestroyQItemSelectionModel() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QItemSelectionModel::~QItemSelectionModel")
		}
	}()

	if ptr.Pointer() != nil {
		C.QItemSelectionModel_DestroyQItemSelectionModel(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
