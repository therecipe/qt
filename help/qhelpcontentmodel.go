package help

//#include "qhelpcontentmodel.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QHelpContentModel struct {
	core.QAbstractItemModel
}

type QHelpContentModelITF interface {
	core.QAbstractItemModelITF
	QHelpContentModelPTR() *QHelpContentModel
}

func PointerFromQHelpContentModel(ptr QHelpContentModelITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpContentModelPTR().Pointer()
	}
	return nil
}

func QHelpContentModelFromPointer(ptr unsafe.Pointer) *QHelpContentModel {
	var n = new(QHelpContentModel)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QHelpContentModel_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QHelpContentModel) QHelpContentModelPTR() *QHelpContentModel {
	return ptr
}

func (ptr *QHelpContentModel) ColumnCount(parent core.QModelIndexITF) int {
	if ptr.Pointer() != nil {
		return int(C.QHelpContentModel_ColumnCount(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(parent))))
	}
	return 0
}

func (ptr *QHelpContentModel) ContentItemAt(index core.QModelIndexITF) *QHelpContentItem {
	if ptr.Pointer() != nil {
		return QHelpContentItemFromPointer(unsafe.Pointer(C.QHelpContentModel_ContentItemAt(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(index)))))
	}
	return nil
}

func (ptr *QHelpContentModel) ConnectContentsCreated(f func()) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_ConnectContentsCreated(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "contentsCreated", f)
	}
}

func (ptr *QHelpContentModel) DisconnectContentsCreated() {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_DisconnectContentsCreated(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "contentsCreated")
	}
}

//export callbackQHelpContentModelContentsCreated
func callbackQHelpContentModelContentsCreated(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "contentsCreated").(func())()
}

func (ptr *QHelpContentModel) ConnectContentsCreationStarted(f func()) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_ConnectContentsCreationStarted(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "contentsCreationStarted", f)
	}
}

func (ptr *QHelpContentModel) DisconnectContentsCreationStarted() {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_DisconnectContentsCreationStarted(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "contentsCreationStarted")
	}
}

//export callbackQHelpContentModelContentsCreationStarted
func callbackQHelpContentModelContentsCreationStarted(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "contentsCreationStarted").(func())()
}

func (ptr *QHelpContentModel) CreateContents(customFilterName string) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_CreateContents(C.QtObjectPtr(ptr.Pointer()), C.CString(customFilterName))
	}
}

func (ptr *QHelpContentModel) Data(index core.QModelIndexITF, role int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QHelpContentModel_Data(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(index)), C.int(role)))
	}
	return ""
}

func (ptr *QHelpContentModel) Index(row int, column int, parent core.QModelIndexITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		return core.QModelIndexFromPointer(unsafe.Pointer(C.QHelpContentModel_Index(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(column), C.QtObjectPtr(core.PointerFromQModelIndex(parent)))))
	}
	return nil
}

func (ptr *QHelpContentModel) IsCreatingContents() bool {
	if ptr.Pointer() != nil {
		return C.QHelpContentModel_IsCreatingContents(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QHelpContentModel) Parent(index core.QModelIndexITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		return core.QModelIndexFromPointer(unsafe.Pointer(C.QHelpContentModel_Parent(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(index)))))
	}
	return nil
}

func (ptr *QHelpContentModel) RowCount(parent core.QModelIndexITF) int {
	if ptr.Pointer() != nil {
		return int(C.QHelpContentModel_RowCount(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(parent))))
	}
	return 0
}

func (ptr *QHelpContentModel) DestroyQHelpContentModel() {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_DestroyQHelpContentModel(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
