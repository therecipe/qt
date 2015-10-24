package gui

//#include "qstandarditemmodel.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QStandardItemModel struct {
	core.QAbstractItemModel
}

type QStandardItemModelITF interface {
	core.QAbstractItemModelITF
	QStandardItemModelPTR() *QStandardItemModel
}

func PointerFromQStandardItemModel(ptr QStandardItemModelITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStandardItemModelPTR().Pointer()
	}
	return nil
}

func QStandardItemModelFromPointer(ptr unsafe.Pointer) *QStandardItemModel {
	var n = new(QStandardItemModel)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QStandardItemModel_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QStandardItemModel) QStandardItemModelPTR() *QStandardItemModel {
	return ptr
}

func (ptr *QStandardItemModel) SetSortRole(role int) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel_SetSortRole(C.QtObjectPtr(ptr.Pointer()), C.int(role))
	}
}

func (ptr *QStandardItemModel) SortRole() int {
	if ptr.Pointer() != nil {
		return int(C.QStandardItemModel_SortRole(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func NewQStandardItemModel(parent core.QObjectITF) *QStandardItemModel {
	return QStandardItemModelFromPointer(unsafe.Pointer(C.QStandardItemModel_NewQStandardItemModel(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func NewQStandardItemModel2(rows int, columns int, parent core.QObjectITF) *QStandardItemModel {
	return QStandardItemModelFromPointer(unsafe.Pointer(C.QStandardItemModel_NewQStandardItemModel2(C.int(rows), C.int(columns), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QStandardItemModel) AppendRow2(item QStandardItemITF) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel_AppendRow2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQStandardItem(item)))
	}
}

func (ptr *QStandardItemModel) Clear() {
	if ptr.Pointer() != nil {
		C.QStandardItemModel_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QStandardItemModel) ColumnCount(parent core.QModelIndexITF) int {
	if ptr.Pointer() != nil {
		return int(C.QStandardItemModel_ColumnCount(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(parent))))
	}
	return 0
}

func (ptr *QStandardItemModel) Data(index core.QModelIndexITF, role int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QStandardItemModel_Data(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(index)), C.int(role)))
	}
	return ""
}

func (ptr *QStandardItemModel) DropMimeData(data core.QMimeDataITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QStandardItemModel_DropMimeData(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQMimeData(data)), C.int(action), C.int(row), C.int(column), C.QtObjectPtr(core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QStandardItemModel) Flags(index core.QModelIndexITF) core.Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return core.Qt__ItemFlag(C.QStandardItemModel_Flags(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(index))))
	}
	return 0
}

func (ptr *QStandardItemModel) HasChildren(parent core.QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QStandardItemModel_HasChildren(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QStandardItemModel) HeaderData(section int, orientation core.Qt__Orientation, role int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QStandardItemModel_HeaderData(C.QtObjectPtr(ptr.Pointer()), C.int(section), C.int(orientation), C.int(role)))
	}
	return ""
}

func (ptr *QStandardItemModel) HorizontalHeaderItem(column int) *QStandardItem {
	if ptr.Pointer() != nil {
		return QStandardItemFromPointer(unsafe.Pointer(C.QStandardItemModel_HorizontalHeaderItem(C.QtObjectPtr(ptr.Pointer()), C.int(column))))
	}
	return nil
}

func (ptr *QStandardItemModel) Index(row int, column int, parent core.QModelIndexITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		return core.QModelIndexFromPointer(unsafe.Pointer(C.QStandardItemModel_Index(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(column), C.QtObjectPtr(core.PointerFromQModelIndex(parent)))))
	}
	return nil
}

func (ptr *QStandardItemModel) IndexFromItem(item QStandardItemITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		return core.QModelIndexFromPointer(unsafe.Pointer(C.QStandardItemModel_IndexFromItem(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQStandardItem(item)))))
	}
	return nil
}

func (ptr *QStandardItemModel) InsertColumn2(column int, parent core.QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QStandardItemModel_InsertColumn2(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.QtObjectPtr(core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QStandardItemModel) InsertColumns(column int, count int, parent core.QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QStandardItemModel_InsertColumns(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.int(count), C.QtObjectPtr(core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QStandardItemModel) InsertRow2(row int, parent core.QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QStandardItemModel_InsertRow2(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.QtObjectPtr(core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QStandardItemModel) InsertRow3(row int, item QStandardItemITF) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel_InsertRow3(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.QtObjectPtr(PointerFromQStandardItem(item)))
	}
}

func (ptr *QStandardItemModel) InsertRows(row int, count int, parent core.QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QStandardItemModel_InsertRows(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(count), C.QtObjectPtr(core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QStandardItemModel) InvisibleRootItem() *QStandardItem {
	if ptr.Pointer() != nil {
		return QStandardItemFromPointer(unsafe.Pointer(C.QStandardItemModel_InvisibleRootItem(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QStandardItemModel) Item(row int, column int) *QStandardItem {
	if ptr.Pointer() != nil {
		return QStandardItemFromPointer(unsafe.Pointer(C.QStandardItemModel_Item(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(column))))
	}
	return nil
}

func (ptr *QStandardItemModel) ConnectItemChanged(f func(item QStandardItemITF)) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel_ConnectItemChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "itemChanged", f)
	}
}

func (ptr *QStandardItemModel) DisconnectItemChanged() {
	if ptr.Pointer() != nil {
		C.QStandardItemModel_DisconnectItemChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "itemChanged")
	}
}

//export callbackQStandardItemModelItemChanged
func callbackQStandardItemModelItemChanged(ptrName *C.char, item unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "itemChanged").(func(*QStandardItem))(QStandardItemFromPointer(item))
}

func (ptr *QStandardItemModel) ItemFromIndex(index core.QModelIndexITF) *QStandardItem {
	if ptr.Pointer() != nil {
		return QStandardItemFromPointer(unsafe.Pointer(C.QStandardItemModel_ItemFromIndex(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(index)))))
	}
	return nil
}

func (ptr *QStandardItemModel) ItemPrototype() *QStandardItem {
	if ptr.Pointer() != nil {
		return QStandardItemFromPointer(unsafe.Pointer(C.QStandardItemModel_ItemPrototype(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QStandardItemModel) MimeTypes() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QStandardItemModel_MimeTypes(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QStandardItemModel) Parent(child core.QModelIndexITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		return core.QModelIndexFromPointer(unsafe.Pointer(C.QStandardItemModel_Parent(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(child)))))
	}
	return nil
}

func (ptr *QStandardItemModel) RemoveColumns(column int, count int, parent core.QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QStandardItemModel_RemoveColumns(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.int(count), C.QtObjectPtr(core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QStandardItemModel) RemoveRows(row int, count int, parent core.QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QStandardItemModel_RemoveRows(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(count), C.QtObjectPtr(core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QStandardItemModel) RowCount(parent core.QModelIndexITF) int {
	if ptr.Pointer() != nil {
		return int(C.QStandardItemModel_RowCount(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(parent))))
	}
	return 0
}

func (ptr *QStandardItemModel) SetColumnCount(columns int) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel_SetColumnCount(C.QtObjectPtr(ptr.Pointer()), C.int(columns))
	}
}

func (ptr *QStandardItemModel) SetData(index core.QModelIndexITF, value string, role int) bool {
	if ptr.Pointer() != nil {
		return C.QStandardItemModel_SetData(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(index)), C.CString(value), C.int(role)) != 0
	}
	return false
}

func (ptr *QStandardItemModel) SetHeaderData(section int, orientation core.Qt__Orientation, value string, role int) bool {
	if ptr.Pointer() != nil {
		return C.QStandardItemModel_SetHeaderData(C.QtObjectPtr(ptr.Pointer()), C.int(section), C.int(orientation), C.CString(value), C.int(role)) != 0
	}
	return false
}

func (ptr *QStandardItemModel) SetHorizontalHeaderItem(column int, item QStandardItemITF) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel_SetHorizontalHeaderItem(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.QtObjectPtr(PointerFromQStandardItem(item)))
	}
}

func (ptr *QStandardItemModel) SetHorizontalHeaderLabels(labels []string) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel_SetHorizontalHeaderLabels(C.QtObjectPtr(ptr.Pointer()), C.CString(strings.Join(labels, "|")))
	}
}

func (ptr *QStandardItemModel) SetItem2(row int, item QStandardItemITF) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel_SetItem2(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.QtObjectPtr(PointerFromQStandardItem(item)))
	}
}

func (ptr *QStandardItemModel) SetItem(row int, column int, item QStandardItemITF) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel_SetItem(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(column), C.QtObjectPtr(PointerFromQStandardItem(item)))
	}
}

func (ptr *QStandardItemModel) SetItemPrototype(item QStandardItemITF) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel_SetItemPrototype(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQStandardItem(item)))
	}
}

func (ptr *QStandardItemModel) SetRowCount(rows int) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel_SetRowCount(C.QtObjectPtr(ptr.Pointer()), C.int(rows))
	}
}

func (ptr *QStandardItemModel) SetVerticalHeaderItem(row int, item QStandardItemITF) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel_SetVerticalHeaderItem(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.QtObjectPtr(PointerFromQStandardItem(item)))
	}
}

func (ptr *QStandardItemModel) SetVerticalHeaderLabels(labels []string) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel_SetVerticalHeaderLabels(C.QtObjectPtr(ptr.Pointer()), C.CString(strings.Join(labels, "|")))
	}
}

func (ptr *QStandardItemModel) Sibling(row int, column int, idx core.QModelIndexITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		return core.QModelIndexFromPointer(unsafe.Pointer(C.QStandardItemModel_Sibling(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(column), C.QtObjectPtr(core.PointerFromQModelIndex(idx)))))
	}
	return nil
}

func (ptr *QStandardItemModel) Sort(column int, order core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel_Sort(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.int(order))
	}
}

func (ptr *QStandardItemModel) SupportedDropActions() core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QStandardItemModel_SupportedDropActions(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QStandardItemModel) TakeHorizontalHeaderItem(column int) *QStandardItem {
	if ptr.Pointer() != nil {
		return QStandardItemFromPointer(unsafe.Pointer(C.QStandardItemModel_TakeHorizontalHeaderItem(C.QtObjectPtr(ptr.Pointer()), C.int(column))))
	}
	return nil
}

func (ptr *QStandardItemModel) TakeItem(row int, column int) *QStandardItem {
	if ptr.Pointer() != nil {
		return QStandardItemFromPointer(unsafe.Pointer(C.QStandardItemModel_TakeItem(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(column))))
	}
	return nil
}

func (ptr *QStandardItemModel) TakeVerticalHeaderItem(row int) *QStandardItem {
	if ptr.Pointer() != nil {
		return QStandardItemFromPointer(unsafe.Pointer(C.QStandardItemModel_TakeVerticalHeaderItem(C.QtObjectPtr(ptr.Pointer()), C.int(row))))
	}
	return nil
}

func (ptr *QStandardItemModel) VerticalHeaderItem(row int) *QStandardItem {
	if ptr.Pointer() != nil {
		return QStandardItemFromPointer(unsafe.Pointer(C.QStandardItemModel_VerticalHeaderItem(C.QtObjectPtr(ptr.Pointer()), C.int(row))))
	}
	return nil
}

func (ptr *QStandardItemModel) DestroyQStandardItemModel() {
	if ptr.Pointer() != nil {
		C.QStandardItemModel_DestroyQStandardItemModel(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
