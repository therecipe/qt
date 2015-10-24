package core

//#include "qabstractproxymodel.h"
import "C"
import (
	"github.com/therecipe/qt"
	"strings"
	"unsafe"
)

type QAbstractProxyModel struct {
	QAbstractItemModel
}

type QAbstractProxyModelITF interface {
	QAbstractItemModelITF
	QAbstractProxyModelPTR() *QAbstractProxyModel
}

func PointerFromQAbstractProxyModel(ptr QAbstractProxyModelITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractProxyModelPTR().Pointer()
	}
	return nil
}

func QAbstractProxyModelFromPointer(ptr unsafe.Pointer) *QAbstractProxyModel {
	var n = new(QAbstractProxyModel)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QAbstractProxyModel_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAbstractProxyModel) QAbstractProxyModelPTR() *QAbstractProxyModel {
	return ptr
}

func (ptr *QAbstractProxyModel) Buddy(index QModelIndexITF) *QModelIndex {
	if ptr.Pointer() != nil {
		return QModelIndexFromPointer(unsafe.Pointer(C.QAbstractProxyModel_Buddy(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(index)))))
	}
	return nil
}

func (ptr *QAbstractProxyModel) CanDropMimeData(data QMimeDataITF, action Qt__DropAction, row int, column int, parent QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractProxyModel_CanDropMimeData(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQMimeData(data)), C.int(action), C.int(row), C.int(column), C.QtObjectPtr(PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QAbstractProxyModel) CanFetchMore(parent QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractProxyModel_CanFetchMore(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QAbstractProxyModel) Data(proxyIndex QModelIndexITF, role int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAbstractProxyModel_Data(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(proxyIndex)), C.int(role)))
	}
	return ""
}

func (ptr *QAbstractProxyModel) DropMimeData(data QMimeDataITF, action Qt__DropAction, row int, column int, parent QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractProxyModel_DropMimeData(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQMimeData(data)), C.int(action), C.int(row), C.int(column), C.QtObjectPtr(PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QAbstractProxyModel) FetchMore(parent QModelIndexITF) {
	if ptr.Pointer() != nil {
		C.QAbstractProxyModel_FetchMore(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(parent)))
	}
}

func (ptr *QAbstractProxyModel) Flags(index QModelIndexITF) Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return Qt__ItemFlag(C.QAbstractProxyModel_Flags(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(index))))
	}
	return 0
}

func (ptr *QAbstractProxyModel) HasChildren(parent QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractProxyModel_HasChildren(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QAbstractProxyModel) HeaderData(section int, orientation Qt__Orientation, role int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAbstractProxyModel_HeaderData(C.QtObjectPtr(ptr.Pointer()), C.int(section), C.int(orientation), C.int(role)))
	}
	return ""
}

func (ptr *QAbstractProxyModel) MapFromSource(sourceIndex QModelIndexITF) *QModelIndex {
	if ptr.Pointer() != nil {
		return QModelIndexFromPointer(unsafe.Pointer(C.QAbstractProxyModel_MapFromSource(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(sourceIndex)))))
	}
	return nil
}

func (ptr *QAbstractProxyModel) MapToSource(proxyIndex QModelIndexITF) *QModelIndex {
	if ptr.Pointer() != nil {
		return QModelIndexFromPointer(unsafe.Pointer(C.QAbstractProxyModel_MapToSource(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(proxyIndex)))))
	}
	return nil
}

func (ptr *QAbstractProxyModel) MimeTypes() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QAbstractProxyModel_MimeTypes(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QAbstractProxyModel) Revert() {
	if ptr.Pointer() != nil {
		C.QAbstractProxyModel_Revert(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QAbstractProxyModel) SetData(index QModelIndexITF, value string, role int) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractProxyModel_SetData(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(index)), C.CString(value), C.int(role)) != 0
	}
	return false
}

func (ptr *QAbstractProxyModel) SetHeaderData(section int, orientation Qt__Orientation, value string, role int) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractProxyModel_SetHeaderData(C.QtObjectPtr(ptr.Pointer()), C.int(section), C.int(orientation), C.CString(value), C.int(role)) != 0
	}
	return false
}

func (ptr *QAbstractProxyModel) SetSourceModel(sourceModel QAbstractItemModelITF) {
	if ptr.Pointer() != nil {
		C.QAbstractProxyModel_SetSourceModel(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAbstractItemModel(sourceModel)))
	}
}

func (ptr *QAbstractProxyModel) Sibling(row int, column int, idx QModelIndexITF) *QModelIndex {
	if ptr.Pointer() != nil {
		return QModelIndexFromPointer(unsafe.Pointer(C.QAbstractProxyModel_Sibling(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(column), C.QtObjectPtr(PointerFromQModelIndex(idx)))))
	}
	return nil
}

func (ptr *QAbstractProxyModel) Sort(column int, order Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.QAbstractProxyModel_Sort(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.int(order))
	}
}

func (ptr *QAbstractProxyModel) SourceModel() *QAbstractItemModel {
	if ptr.Pointer() != nil {
		return QAbstractItemModelFromPointer(unsafe.Pointer(C.QAbstractProxyModel_SourceModel(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QAbstractProxyModel) ConnectSourceModelChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QAbstractProxyModel_ConnectSourceModelChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "sourceModelChanged", f)
	}
}

func (ptr *QAbstractProxyModel) DisconnectSourceModelChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractProxyModel_DisconnectSourceModelChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "sourceModelChanged")
	}
}

//export callbackQAbstractProxyModelSourceModelChanged
func callbackQAbstractProxyModelSourceModelChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "sourceModelChanged").(func())()
}

func (ptr *QAbstractProxyModel) Submit() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractProxyModel_Submit(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAbstractProxyModel) SupportedDragActions() Qt__DropAction {
	if ptr.Pointer() != nil {
		return Qt__DropAction(C.QAbstractProxyModel_SupportedDragActions(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractProxyModel) SupportedDropActions() Qt__DropAction {
	if ptr.Pointer() != nil {
		return Qt__DropAction(C.QAbstractProxyModel_SupportedDropActions(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractProxyModel) DestroyQAbstractProxyModel() {
	if ptr.Pointer() != nil {
		C.QAbstractProxyModel_DestroyQAbstractProxyModel(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
