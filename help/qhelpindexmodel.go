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
	if n.ObjectName() == "" {
		n.SetObjectName("QHelpIndexModel_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QHelpIndexModel) QHelpIndexModel_PTR() *QHelpIndexModel {
	return ptr
}

func (ptr *QHelpIndexModel) CreateIndex(customFilterName string) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_CreateIndex(ptr.Pointer(), C.CString(customFilterName))
	}
}

func (ptr *QHelpIndexModel) Filter(filter string, wildcard string) *core.QModelIndex {
	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QHelpIndexModel_Filter(ptr.Pointer(), C.CString(filter), C.CString(wildcard)))
	}
	return nil
}

func (ptr *QHelpIndexModel) ConnectIndexCreated(f func()) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_ConnectIndexCreated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "indexCreated", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectIndexCreated() {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_DisconnectIndexCreated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "indexCreated")
	}
}

//export callbackQHelpIndexModelIndexCreated
func callbackQHelpIndexModelIndexCreated(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "indexCreated").(func())()
}

func (ptr *QHelpIndexModel) ConnectIndexCreationStarted(f func()) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_ConnectIndexCreationStarted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "indexCreationStarted", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectIndexCreationStarted() {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_DisconnectIndexCreationStarted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "indexCreationStarted")
	}
}

//export callbackQHelpIndexModelIndexCreationStarted
func callbackQHelpIndexModelIndexCreationStarted(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "indexCreationStarted").(func())()
}

func (ptr *QHelpIndexModel) IsCreatingIndex() bool {
	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_IsCreatingIndex(ptr.Pointer()) != 0
	}
	return false
}
