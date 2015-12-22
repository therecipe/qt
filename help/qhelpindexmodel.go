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
func callbackQHelpIndexModelIndexCreated(ptrName *C.char) {
	defer qt.Recovering("callback QHelpIndexModel::indexCreated")

	if signal := qt.GetSignal(C.GoString(ptrName), "indexCreated"); signal != nil {
		signal.(func())()
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
func callbackQHelpIndexModelIndexCreationStarted(ptrName *C.char) {
	defer qt.Recovering("callback QHelpIndexModel::indexCreationStarted")

	if signal := qt.GetSignal(C.GoString(ptrName), "indexCreationStarted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpIndexModel) IsCreatingIndex() bool {
	defer qt.Recovering("QHelpIndexModel::isCreatingIndex")

	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_IsCreatingIndex(ptr.Pointer()) != 0
	}
	return false
}
