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
func callbackQHelpContentModelContentsCreated(ptrName *C.char) {
	defer qt.Recovering("callback QHelpContentModel::contentsCreated")

	if signal := qt.GetSignal(C.GoString(ptrName), "contentsCreated"); signal != nil {
		signal.(func())()
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
func callbackQHelpContentModelContentsCreationStarted(ptrName *C.char) {
	defer qt.Recovering("callback QHelpContentModel::contentsCreationStarted")

	if signal := qt.GetSignal(C.GoString(ptrName), "contentsCreationStarted"); signal != nil {
		signal.(func())()
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
