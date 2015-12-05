package help

//#include "help.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
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
		n.SetObjectName("QHelpContentModel_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QHelpContentModel) QHelpContentModel_PTR() *QHelpContentModel {
	return ptr
}

func (ptr *QHelpContentModel) ColumnCount(parent core.QModelIndex_ITF) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHelpContentModel::columnCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QHelpContentModel_ColumnCount(ptr.Pointer(), core.PointerFromQModelIndex(parent)))
	}
	return 0
}

func (ptr *QHelpContentModel) ContentItemAt(index core.QModelIndex_ITF) *QHelpContentItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHelpContentModel::contentItemAt")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQHelpContentItemFromPointer(C.QHelpContentModel_ContentItemAt(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return nil
}

func (ptr *QHelpContentModel) ConnectContentsCreated(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHelpContentModel::contentsCreated")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHelpContentModel_ConnectContentsCreated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "contentsCreated", f)
	}
}

func (ptr *QHelpContentModel) DisconnectContentsCreated() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHelpContentModel::contentsCreated")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHelpContentModel_DisconnectContentsCreated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "contentsCreated")
	}
}

//export callbackQHelpContentModelContentsCreated
func callbackQHelpContentModelContentsCreated(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHelpContentModel::contentsCreated")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "contentsCreated").(func())()
}

func (ptr *QHelpContentModel) ConnectContentsCreationStarted(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHelpContentModel::contentsCreationStarted")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHelpContentModel_ConnectContentsCreationStarted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "contentsCreationStarted", f)
	}
}

func (ptr *QHelpContentModel) DisconnectContentsCreationStarted() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHelpContentModel::contentsCreationStarted")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHelpContentModel_DisconnectContentsCreationStarted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "contentsCreationStarted")
	}
}

//export callbackQHelpContentModelContentsCreationStarted
func callbackQHelpContentModelContentsCreationStarted(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHelpContentModel::contentsCreationStarted")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "contentsCreationStarted").(func())()
}

func (ptr *QHelpContentModel) CreateContents(customFilterName string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHelpContentModel::createContents")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHelpContentModel_CreateContents(ptr.Pointer(), C.CString(customFilterName))
	}
}

func (ptr *QHelpContentModel) Data(index core.QModelIndex_ITF, role int) *core.QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHelpContentModel::data")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QHelpContentModel_Data(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(role)))
	}
	return nil
}

func (ptr *QHelpContentModel) Index(row int, column int, parent core.QModelIndex_ITF) *core.QModelIndex {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHelpContentModel::index")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QHelpContentModel_Index(ptr.Pointer(), C.int(row), C.int(column), core.PointerFromQModelIndex(parent)))
	}
	return nil
}

func (ptr *QHelpContentModel) IsCreatingContents() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHelpContentModel::isCreatingContents")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QHelpContentModel_IsCreatingContents(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHelpContentModel) Parent(index core.QModelIndex_ITF) *core.QModelIndex {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHelpContentModel::parent")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QHelpContentModel_Parent(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return nil
}

func (ptr *QHelpContentModel) RowCount(parent core.QModelIndex_ITF) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHelpContentModel::rowCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QHelpContentModel_RowCount(ptr.Pointer(), core.PointerFromQModelIndex(parent)))
	}
	return 0
}

func (ptr *QHelpContentModel) DestroyQHelpContentModel() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHelpContentModel::~QHelpContentModel")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHelpContentModel_DestroyQHelpContentModel(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
