package help

//#include "qhelpindexmodel.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QHelpIndexModel struct {
	core.QStringListModel
}

type QHelpIndexModelITF interface {
	core.QStringListModelITF
	QHelpIndexModelPTR() *QHelpIndexModel
}

func PointerFromQHelpIndexModel(ptr QHelpIndexModelITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpIndexModelPTR().Pointer()
	}
	return nil
}

func QHelpIndexModelFromPointer(ptr unsafe.Pointer) *QHelpIndexModel {
	var n = new(QHelpIndexModel)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QHelpIndexModel_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QHelpIndexModel) QHelpIndexModelPTR() *QHelpIndexModel {
	return ptr
}

func (ptr *QHelpIndexModel) CreateIndex(customFilterName string) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_CreateIndex(C.QtObjectPtr(ptr.Pointer()), C.CString(customFilterName))
	}
}

func (ptr *QHelpIndexModel) Filter(filter string, wildcard string) *core.QModelIndex {
	if ptr.Pointer() != nil {
		return core.QModelIndexFromPointer(unsafe.Pointer(C.QHelpIndexModel_Filter(C.QtObjectPtr(ptr.Pointer()), C.CString(filter), C.CString(wildcard))))
	}
	return nil
}

func (ptr *QHelpIndexModel) ConnectIndexCreated(f func()) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_ConnectIndexCreated(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "indexCreated", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectIndexCreated() {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_DisconnectIndexCreated(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "indexCreated")
	}
}

//export callbackQHelpIndexModelIndexCreated
func callbackQHelpIndexModelIndexCreated(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "indexCreated").(func())()
}

func (ptr *QHelpIndexModel) ConnectIndexCreationStarted(f func()) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_ConnectIndexCreationStarted(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "indexCreationStarted", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectIndexCreationStarted() {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_DisconnectIndexCreationStarted(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "indexCreationStarted")
	}
}

//export callbackQHelpIndexModelIndexCreationStarted
func callbackQHelpIndexModelIndexCreationStarted(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "indexCreationStarted").(func())()
}

func (ptr *QHelpIndexModel) IsCreatingIndex() bool {
	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_IsCreatingIndex(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}
