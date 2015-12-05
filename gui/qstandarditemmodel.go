package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"strings"
	"unsafe"
)

type QStandardItemModel struct {
	core.QAbstractItemModel
}

type QStandardItemModel_ITF interface {
	core.QAbstractItemModel_ITF
	QStandardItemModel_PTR() *QStandardItemModel
}

func PointerFromQStandardItemModel(ptr QStandardItemModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStandardItemModel_PTR().Pointer()
	}
	return nil
}

func NewQStandardItemModelFromPointer(ptr unsafe.Pointer) *QStandardItemModel {
	var n = new(QStandardItemModel)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QStandardItemModel_") {
		n.SetObjectName("QStandardItemModel_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QStandardItemModel) QStandardItemModel_PTR() *QStandardItemModel {
	return ptr
}

func (ptr *QStandardItemModel) SetSortRole(role int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::setSortRole")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItemModel_SetSortRole(ptr.Pointer(), C.int(role))
	}
}

func (ptr *QStandardItemModel) SortRole() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::sortRole")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStandardItemModel_SortRole(ptr.Pointer()))
	}
	return 0
}

func NewQStandardItemModel(parent core.QObject_ITF) *QStandardItemModel {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::QStandardItemModel")
		}
	}()

	return NewQStandardItemModelFromPointer(C.QStandardItemModel_NewQStandardItemModel(core.PointerFromQObject(parent)))
}

func NewQStandardItemModel2(rows int, columns int, parent core.QObject_ITF) *QStandardItemModel {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::QStandardItemModel")
		}
	}()

	return NewQStandardItemModelFromPointer(C.QStandardItemModel_NewQStandardItemModel2(C.int(rows), C.int(columns), core.PointerFromQObject(parent)))
}

func (ptr *QStandardItemModel) AppendRow2(item QStandardItem_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::appendRow")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItemModel_AppendRow2(ptr.Pointer(), PointerFromQStandardItem(item))
	}
}

func (ptr *QStandardItemModel) Clear() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::clear")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItemModel_Clear(ptr.Pointer())
	}
}

func (ptr *QStandardItemModel) ColumnCount(parent core.QModelIndex_ITF) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::columnCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStandardItemModel_ColumnCount(ptr.Pointer(), core.PointerFromQModelIndex(parent)))
	}
	return 0
}

func (ptr *QStandardItemModel) Data(index core.QModelIndex_ITF, role int) *core.QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::data")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QStandardItemModel_Data(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(role)))
	}
	return nil
}

func (ptr *QStandardItemModel) DropMimeData(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::dropMimeData")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QStandardItemModel_DropMimeData(ptr.Pointer(), core.PointerFromQMimeData(data), C.int(action), C.int(row), C.int(column), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QStandardItemModel) Flags(index core.QModelIndex_ITF) core.Qt__ItemFlag {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::flags")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__ItemFlag(C.QStandardItemModel_Flags(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return 0
}

func (ptr *QStandardItemModel) HasChildren(parent core.QModelIndex_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::hasChildren")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QStandardItemModel_HasChildren(ptr.Pointer(), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QStandardItemModel) HeaderData(section int, orientation core.Qt__Orientation, role int) *core.QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::headerData")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QStandardItemModel_HeaderData(ptr.Pointer(), C.int(section), C.int(orientation), C.int(role)))
	}
	return nil
}

func (ptr *QStandardItemModel) HorizontalHeaderItem(column int) *QStandardItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::horizontalHeaderItem")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItemModel_HorizontalHeaderItem(ptr.Pointer(), C.int(column)))
	}
	return nil
}

func (ptr *QStandardItemModel) Index(row int, column int, parent core.QModelIndex_ITF) *core.QModelIndex {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::index")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QStandardItemModel_Index(ptr.Pointer(), C.int(row), C.int(column), core.PointerFromQModelIndex(parent)))
	}
	return nil
}

func (ptr *QStandardItemModel) IndexFromItem(item QStandardItem_ITF) *core.QModelIndex {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::indexFromItem")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QStandardItemModel_IndexFromItem(ptr.Pointer(), PointerFromQStandardItem(item)))
	}
	return nil
}

func (ptr *QStandardItemModel) InsertColumn2(column int, parent core.QModelIndex_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::insertColumn")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QStandardItemModel_InsertColumn2(ptr.Pointer(), C.int(column), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QStandardItemModel) InsertColumns(column int, count int, parent core.QModelIndex_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::insertColumns")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QStandardItemModel_InsertColumns(ptr.Pointer(), C.int(column), C.int(count), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QStandardItemModel) InsertRow2(row int, parent core.QModelIndex_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::insertRow")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QStandardItemModel_InsertRow2(ptr.Pointer(), C.int(row), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QStandardItemModel) InsertRow3(row int, item QStandardItem_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::insertRow")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItemModel_InsertRow3(ptr.Pointer(), C.int(row), PointerFromQStandardItem(item))
	}
}

func (ptr *QStandardItemModel) InsertRows(row int, count int, parent core.QModelIndex_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::insertRows")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QStandardItemModel_InsertRows(ptr.Pointer(), C.int(row), C.int(count), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QStandardItemModel) InvisibleRootItem() *QStandardItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::invisibleRootItem")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItemModel_InvisibleRootItem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStandardItemModel) Item(row int, column int) *QStandardItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::item")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItemModel_Item(ptr.Pointer(), C.int(row), C.int(column)))
	}
	return nil
}

func (ptr *QStandardItemModel) ConnectItemChanged(f func(item *QStandardItem)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::itemChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItemModel_ConnectItemChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "itemChanged", f)
	}
}

func (ptr *QStandardItemModel) DisconnectItemChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::itemChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItemModel_DisconnectItemChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "itemChanged")
	}
}

//export callbackQStandardItemModelItemChanged
func callbackQStandardItemModelItemChanged(ptrName *C.char, item unsafe.Pointer) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::itemChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "itemChanged").(func(*QStandardItem))(NewQStandardItemFromPointer(item))
}

func (ptr *QStandardItemModel) ItemFromIndex(index core.QModelIndex_ITF) *QStandardItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::itemFromIndex")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItemModel_ItemFromIndex(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return nil
}

func (ptr *QStandardItemModel) ItemPrototype() *QStandardItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::itemPrototype")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItemModel_ItemPrototype(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStandardItemModel) MimeTypes() []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::mimeTypes")
		}
	}()

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QStandardItemModel_MimeTypes(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QStandardItemModel) Parent(child core.QModelIndex_ITF) *core.QModelIndex {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::parent")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QStandardItemModel_Parent(ptr.Pointer(), core.PointerFromQModelIndex(child)))
	}
	return nil
}

func (ptr *QStandardItemModel) RemoveColumns(column int, count int, parent core.QModelIndex_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::removeColumns")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QStandardItemModel_RemoveColumns(ptr.Pointer(), C.int(column), C.int(count), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QStandardItemModel) RemoveRows(row int, count int, parent core.QModelIndex_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::removeRows")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QStandardItemModel_RemoveRows(ptr.Pointer(), C.int(row), C.int(count), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QStandardItemModel) RowCount(parent core.QModelIndex_ITF) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::rowCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStandardItemModel_RowCount(ptr.Pointer(), core.PointerFromQModelIndex(parent)))
	}
	return 0
}

func (ptr *QStandardItemModel) SetColumnCount(columns int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::setColumnCount")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItemModel_SetColumnCount(ptr.Pointer(), C.int(columns))
	}
}

func (ptr *QStandardItemModel) SetData(index core.QModelIndex_ITF, value core.QVariant_ITF, role int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::setData")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QStandardItemModel_SetData(ptr.Pointer(), core.PointerFromQModelIndex(index), core.PointerFromQVariant(value), C.int(role)) != 0
	}
	return false
}

func (ptr *QStandardItemModel) SetHeaderData(section int, orientation core.Qt__Orientation, value core.QVariant_ITF, role int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::setHeaderData")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QStandardItemModel_SetHeaderData(ptr.Pointer(), C.int(section), C.int(orientation), core.PointerFromQVariant(value), C.int(role)) != 0
	}
	return false
}

func (ptr *QStandardItemModel) SetHorizontalHeaderItem(column int, item QStandardItem_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::setHorizontalHeaderItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItemModel_SetHorizontalHeaderItem(ptr.Pointer(), C.int(column), PointerFromQStandardItem(item))
	}
}

func (ptr *QStandardItemModel) SetHorizontalHeaderLabels(labels []string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::setHorizontalHeaderLabels")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItemModel_SetHorizontalHeaderLabels(ptr.Pointer(), C.CString(strings.Join(labels, ",,,")))
	}
}

func (ptr *QStandardItemModel) SetItem2(row int, item QStandardItem_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::setItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItemModel_SetItem2(ptr.Pointer(), C.int(row), PointerFromQStandardItem(item))
	}
}

func (ptr *QStandardItemModel) SetItem(row int, column int, item QStandardItem_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::setItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItemModel_SetItem(ptr.Pointer(), C.int(row), C.int(column), PointerFromQStandardItem(item))
	}
}

func (ptr *QStandardItemModel) SetItemPrototype(item QStandardItem_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::setItemPrototype")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItemModel_SetItemPrototype(ptr.Pointer(), PointerFromQStandardItem(item))
	}
}

func (ptr *QStandardItemModel) SetRowCount(rows int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::setRowCount")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItemModel_SetRowCount(ptr.Pointer(), C.int(rows))
	}
}

func (ptr *QStandardItemModel) SetVerticalHeaderItem(row int, item QStandardItem_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::setVerticalHeaderItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItemModel_SetVerticalHeaderItem(ptr.Pointer(), C.int(row), PointerFromQStandardItem(item))
	}
}

func (ptr *QStandardItemModel) SetVerticalHeaderLabels(labels []string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::setVerticalHeaderLabels")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItemModel_SetVerticalHeaderLabels(ptr.Pointer(), C.CString(strings.Join(labels, ",,,")))
	}
}

func (ptr *QStandardItemModel) Sibling(row int, column int, idx core.QModelIndex_ITF) *core.QModelIndex {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::sibling")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QStandardItemModel_Sibling(ptr.Pointer(), C.int(row), C.int(column), core.PointerFromQModelIndex(idx)))
	}
	return nil
}

func (ptr *QStandardItemModel) Sort(column int, order core.Qt__SortOrder) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::sort")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItemModel_Sort(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QStandardItemModel) SupportedDropActions() core.Qt__DropAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::supportedDropActions")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QStandardItemModel_SupportedDropActions(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStandardItemModel) TakeHorizontalHeaderItem(column int) *QStandardItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::takeHorizontalHeaderItem")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItemModel_TakeHorizontalHeaderItem(ptr.Pointer(), C.int(column)))
	}
	return nil
}

func (ptr *QStandardItemModel) TakeItem(row int, column int) *QStandardItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::takeItem")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItemModel_TakeItem(ptr.Pointer(), C.int(row), C.int(column)))
	}
	return nil
}

func (ptr *QStandardItemModel) TakeVerticalHeaderItem(row int) *QStandardItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::takeVerticalHeaderItem")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItemModel_TakeVerticalHeaderItem(ptr.Pointer(), C.int(row)))
	}
	return nil
}

func (ptr *QStandardItemModel) VerticalHeaderItem(row int) *QStandardItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::verticalHeaderItem")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItemModel_VerticalHeaderItem(ptr.Pointer(), C.int(row)))
	}
	return nil
}

func (ptr *QStandardItemModel) DestroyQStandardItemModel() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStandardItemModel::~QStandardItemModel")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStandardItemModel_DestroyQStandardItemModel(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
