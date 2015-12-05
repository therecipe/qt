package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"strings"
	"unsafe"
)

type QTreeWidget struct {
	QTreeView
}

type QTreeWidget_ITF interface {
	QTreeView_ITF
	QTreeWidget_PTR() *QTreeWidget
}

func PointerFromQTreeWidget(ptr QTreeWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTreeWidget_PTR().Pointer()
	}
	return nil
}

func NewQTreeWidgetFromPointer(ptr unsafe.Pointer) *QTreeWidget {
	var n = new(QTreeWidget)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QTreeWidget_") {
		n.SetObjectName("QTreeWidget_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTreeWidget) QTreeWidget_PTR() *QTreeWidget {
	return ptr
}

func (ptr *QTreeWidget) ColumnCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::columnCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTreeWidget_ColumnCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTreeWidget) SetColumnCount(columns int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::setColumnCount")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidget_SetColumnCount(ptr.Pointer(), C.int(columns))
	}
}

func (ptr *QTreeWidget) TopLevelItemCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::topLevelItemCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTreeWidget_TopLevelItemCount(ptr.Pointer()))
	}
	return 0
}

func NewQTreeWidget(parent QWidget_ITF) *QTreeWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::QTreeWidget")
		}
	}()

	return NewQTreeWidgetFromPointer(C.QTreeWidget_NewQTreeWidget(PointerFromQWidget(parent)))
}

func (ptr *QTreeWidget) AddTopLevelItem(item QTreeWidgetItem_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::addTopLevelItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidget_AddTopLevelItem(ptr.Pointer(), PointerFromQTreeWidgetItem(item))
	}
}

func (ptr *QTreeWidget) Clear() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::clear")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidget_Clear(ptr.Pointer())
	}
}

func (ptr *QTreeWidget) ClosePersistentEditor(item QTreeWidgetItem_ITF, column int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::closePersistentEditor")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidget_ClosePersistentEditor(ptr.Pointer(), PointerFromQTreeWidgetItem(item), C.int(column))
	}
}

func (ptr *QTreeWidget) CollapseItem(item QTreeWidgetItem_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::collapseItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidget_CollapseItem(ptr.Pointer(), PointerFromQTreeWidgetItem(item))
	}
}

func (ptr *QTreeWidget) CurrentColumn() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::currentColumn")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTreeWidget_CurrentColumn(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTreeWidget) CurrentItem() *QTreeWidgetItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::currentItem")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQTreeWidgetItemFromPointer(C.QTreeWidget_CurrentItem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTreeWidget) ConnectCurrentItemChanged(f func(current *QTreeWidgetItem, previous *QTreeWidgetItem)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::currentItemChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidget_ConnectCurrentItemChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentItemChanged", f)
	}
}

func (ptr *QTreeWidget) DisconnectCurrentItemChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::currentItemChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidget_DisconnectCurrentItemChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentItemChanged")
	}
}

//export callbackQTreeWidgetCurrentItemChanged
func callbackQTreeWidgetCurrentItemChanged(ptrName *C.char, current unsafe.Pointer, previous unsafe.Pointer) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::currentItemChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "currentItemChanged").(func(*QTreeWidgetItem, *QTreeWidgetItem))(NewQTreeWidgetItemFromPointer(current), NewQTreeWidgetItemFromPointer(previous))
}

func (ptr *QTreeWidget) EditItem(item QTreeWidgetItem_ITF, column int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::editItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidget_EditItem(ptr.Pointer(), PointerFromQTreeWidgetItem(item), C.int(column))
	}
}

func (ptr *QTreeWidget) ExpandItem(item QTreeWidgetItem_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::expandItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidget_ExpandItem(ptr.Pointer(), PointerFromQTreeWidgetItem(item))
	}
}

func (ptr *QTreeWidget) HeaderItem() *QTreeWidgetItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::headerItem")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQTreeWidgetItemFromPointer(C.QTreeWidget_HeaderItem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTreeWidget) IndexOfTopLevelItem(item QTreeWidgetItem_ITF) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::indexOfTopLevelItem")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTreeWidget_IndexOfTopLevelItem(ptr.Pointer(), PointerFromQTreeWidgetItem(item)))
	}
	return 0
}

func (ptr *QTreeWidget) InsertTopLevelItem(index int, item QTreeWidgetItem_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::insertTopLevelItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidget_InsertTopLevelItem(ptr.Pointer(), C.int(index), PointerFromQTreeWidgetItem(item))
	}
}

func (ptr *QTreeWidget) InvisibleRootItem() *QTreeWidgetItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::invisibleRootItem")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQTreeWidgetItemFromPointer(C.QTreeWidget_InvisibleRootItem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTreeWidget) IsFirstItemColumnSpanned(item QTreeWidgetItem_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::isFirstItemColumnSpanned")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTreeWidget_IsFirstItemColumnSpanned(ptr.Pointer(), PointerFromQTreeWidgetItem(item)) != 0
	}
	return false
}

func (ptr *QTreeWidget) ItemAbove(item QTreeWidgetItem_ITF) *QTreeWidgetItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::itemAbove")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQTreeWidgetItemFromPointer(C.QTreeWidget_ItemAbove(ptr.Pointer(), PointerFromQTreeWidgetItem(item)))
	}
	return nil
}

func (ptr *QTreeWidget) ConnectItemActivated(f func(item *QTreeWidgetItem, column int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::itemActivated")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidget_ConnectItemActivated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "itemActivated", f)
	}
}

func (ptr *QTreeWidget) DisconnectItemActivated() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::itemActivated")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidget_DisconnectItemActivated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "itemActivated")
	}
}

//export callbackQTreeWidgetItemActivated
func callbackQTreeWidgetItemActivated(ptrName *C.char, item unsafe.Pointer, column C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::itemActivated")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "itemActivated").(func(*QTreeWidgetItem, int))(NewQTreeWidgetItemFromPointer(item), int(column))
}

func (ptr *QTreeWidget) ItemAt(p core.QPoint_ITF) *QTreeWidgetItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::itemAt")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQTreeWidgetItemFromPointer(C.QTreeWidget_ItemAt(ptr.Pointer(), core.PointerFromQPoint(p)))
	}
	return nil
}

func (ptr *QTreeWidget) ItemAt2(x int, y int) *QTreeWidgetItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::itemAt")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQTreeWidgetItemFromPointer(C.QTreeWidget_ItemAt2(ptr.Pointer(), C.int(x), C.int(y)))
	}
	return nil
}

func (ptr *QTreeWidget) ItemBelow(item QTreeWidgetItem_ITF) *QTreeWidgetItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::itemBelow")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQTreeWidgetItemFromPointer(C.QTreeWidget_ItemBelow(ptr.Pointer(), PointerFromQTreeWidgetItem(item)))
	}
	return nil
}

func (ptr *QTreeWidget) ConnectItemChanged(f func(item *QTreeWidgetItem, column int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::itemChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidget_ConnectItemChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "itemChanged", f)
	}
}

func (ptr *QTreeWidget) DisconnectItemChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::itemChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidget_DisconnectItemChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "itemChanged")
	}
}

//export callbackQTreeWidgetItemChanged
func callbackQTreeWidgetItemChanged(ptrName *C.char, item unsafe.Pointer, column C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::itemChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "itemChanged").(func(*QTreeWidgetItem, int))(NewQTreeWidgetItemFromPointer(item), int(column))
}

func (ptr *QTreeWidget) ConnectItemClicked(f func(item *QTreeWidgetItem, column int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::itemClicked")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidget_ConnectItemClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "itemClicked", f)
	}
}

func (ptr *QTreeWidget) DisconnectItemClicked() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::itemClicked")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidget_DisconnectItemClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "itemClicked")
	}
}

//export callbackQTreeWidgetItemClicked
func callbackQTreeWidgetItemClicked(ptrName *C.char, item unsafe.Pointer, column C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::itemClicked")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "itemClicked").(func(*QTreeWidgetItem, int))(NewQTreeWidgetItemFromPointer(item), int(column))
}

func (ptr *QTreeWidget) ConnectItemCollapsed(f func(item *QTreeWidgetItem)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::itemCollapsed")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidget_ConnectItemCollapsed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "itemCollapsed", f)
	}
}

func (ptr *QTreeWidget) DisconnectItemCollapsed() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::itemCollapsed")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidget_DisconnectItemCollapsed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "itemCollapsed")
	}
}

//export callbackQTreeWidgetItemCollapsed
func callbackQTreeWidgetItemCollapsed(ptrName *C.char, item unsafe.Pointer) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::itemCollapsed")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "itemCollapsed").(func(*QTreeWidgetItem))(NewQTreeWidgetItemFromPointer(item))
}

func (ptr *QTreeWidget) ConnectItemDoubleClicked(f func(item *QTreeWidgetItem, column int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::itemDoubleClicked")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidget_ConnectItemDoubleClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "itemDoubleClicked", f)
	}
}

func (ptr *QTreeWidget) DisconnectItemDoubleClicked() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::itemDoubleClicked")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidget_DisconnectItemDoubleClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "itemDoubleClicked")
	}
}

//export callbackQTreeWidgetItemDoubleClicked
func callbackQTreeWidgetItemDoubleClicked(ptrName *C.char, item unsafe.Pointer, column C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::itemDoubleClicked")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "itemDoubleClicked").(func(*QTreeWidgetItem, int))(NewQTreeWidgetItemFromPointer(item), int(column))
}

func (ptr *QTreeWidget) ConnectItemEntered(f func(item *QTreeWidgetItem, column int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::itemEntered")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidget_ConnectItemEntered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "itemEntered", f)
	}
}

func (ptr *QTreeWidget) DisconnectItemEntered() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::itemEntered")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidget_DisconnectItemEntered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "itemEntered")
	}
}

//export callbackQTreeWidgetItemEntered
func callbackQTreeWidgetItemEntered(ptrName *C.char, item unsafe.Pointer, column C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::itemEntered")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "itemEntered").(func(*QTreeWidgetItem, int))(NewQTreeWidgetItemFromPointer(item), int(column))
}

func (ptr *QTreeWidget) ConnectItemExpanded(f func(item *QTreeWidgetItem)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::itemExpanded")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidget_ConnectItemExpanded(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "itemExpanded", f)
	}
}

func (ptr *QTreeWidget) DisconnectItemExpanded() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::itemExpanded")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidget_DisconnectItemExpanded(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "itemExpanded")
	}
}

//export callbackQTreeWidgetItemExpanded
func callbackQTreeWidgetItemExpanded(ptrName *C.char, item unsafe.Pointer) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::itemExpanded")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "itemExpanded").(func(*QTreeWidgetItem))(NewQTreeWidgetItemFromPointer(item))
}

func (ptr *QTreeWidget) ConnectItemPressed(f func(item *QTreeWidgetItem, column int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::itemPressed")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidget_ConnectItemPressed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "itemPressed", f)
	}
}

func (ptr *QTreeWidget) DisconnectItemPressed() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::itemPressed")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidget_DisconnectItemPressed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "itemPressed")
	}
}

//export callbackQTreeWidgetItemPressed
func callbackQTreeWidgetItemPressed(ptrName *C.char, item unsafe.Pointer, column C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::itemPressed")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "itemPressed").(func(*QTreeWidgetItem, int))(NewQTreeWidgetItemFromPointer(item), int(column))
}

func (ptr *QTreeWidget) ConnectItemSelectionChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::itemSelectionChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidget_ConnectItemSelectionChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "itemSelectionChanged", f)
	}
}

func (ptr *QTreeWidget) DisconnectItemSelectionChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::itemSelectionChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidget_DisconnectItemSelectionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "itemSelectionChanged")
	}
}

//export callbackQTreeWidgetItemSelectionChanged
func callbackQTreeWidgetItemSelectionChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::itemSelectionChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "itemSelectionChanged").(func())()
}

func (ptr *QTreeWidget) ItemWidget(item QTreeWidgetItem_ITF, column int) *QWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::itemWidget")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QTreeWidget_ItemWidget(ptr.Pointer(), PointerFromQTreeWidgetItem(item), C.int(column)))
	}
	return nil
}

func (ptr *QTreeWidget) OpenPersistentEditor(item QTreeWidgetItem_ITF, column int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::openPersistentEditor")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidget_OpenPersistentEditor(ptr.Pointer(), PointerFromQTreeWidgetItem(item), C.int(column))
	}
}

func (ptr *QTreeWidget) RemoveItemWidget(item QTreeWidgetItem_ITF, column int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::removeItemWidget")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidget_RemoveItemWidget(ptr.Pointer(), PointerFromQTreeWidgetItem(item), C.int(column))
	}
}

func (ptr *QTreeWidget) ScrollToItem(item QTreeWidgetItem_ITF, hint QAbstractItemView__ScrollHint) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::scrollToItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidget_ScrollToItem(ptr.Pointer(), PointerFromQTreeWidgetItem(item), C.int(hint))
	}
}

func (ptr *QTreeWidget) SetCurrentItem(item QTreeWidgetItem_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::setCurrentItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidget_SetCurrentItem(ptr.Pointer(), PointerFromQTreeWidgetItem(item))
	}
}

func (ptr *QTreeWidget) SetCurrentItem2(item QTreeWidgetItem_ITF, column int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::setCurrentItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidget_SetCurrentItem2(ptr.Pointer(), PointerFromQTreeWidgetItem(item), C.int(column))
	}
}

func (ptr *QTreeWidget) SetCurrentItem3(item QTreeWidgetItem_ITF, column int, command core.QItemSelectionModel__SelectionFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::setCurrentItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidget_SetCurrentItem3(ptr.Pointer(), PointerFromQTreeWidgetItem(item), C.int(column), C.int(command))
	}
}

func (ptr *QTreeWidget) SetFirstItemColumnSpanned(item QTreeWidgetItem_ITF, span bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::setFirstItemColumnSpanned")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidget_SetFirstItemColumnSpanned(ptr.Pointer(), PointerFromQTreeWidgetItem(item), C.int(qt.GoBoolToInt(span)))
	}
}

func (ptr *QTreeWidget) SetHeaderItem(item QTreeWidgetItem_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::setHeaderItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidget_SetHeaderItem(ptr.Pointer(), PointerFromQTreeWidgetItem(item))
	}
}

func (ptr *QTreeWidget) SetHeaderLabel(label string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::setHeaderLabel")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidget_SetHeaderLabel(ptr.Pointer(), C.CString(label))
	}
}

func (ptr *QTreeWidget) SetHeaderLabels(labels []string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::setHeaderLabels")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidget_SetHeaderLabels(ptr.Pointer(), C.CString(strings.Join(labels, ",,,")))
	}
}

func (ptr *QTreeWidget) SetItemWidget(item QTreeWidgetItem_ITF, column int, widget QWidget_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::setItemWidget")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidget_SetItemWidget(ptr.Pointer(), PointerFromQTreeWidgetItem(item), C.int(column), PointerFromQWidget(widget))
	}
}

func (ptr *QTreeWidget) SetSelectionModel(selectionModel core.QItemSelectionModel_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::setSelectionModel")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidget_SetSelectionModel(ptr.Pointer(), core.PointerFromQItemSelectionModel(selectionModel))
	}
}

func (ptr *QTreeWidget) SortColumn() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::sortColumn")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTreeWidget_SortColumn(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTreeWidget) SortItems(column int, order core.Qt__SortOrder) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::sortItems")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidget_SortItems(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QTreeWidget) TakeTopLevelItem(index int) *QTreeWidgetItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::takeTopLevelItem")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQTreeWidgetItemFromPointer(C.QTreeWidget_TakeTopLevelItem(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QTreeWidget) TopLevelItem(index int) *QTreeWidgetItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::topLevelItem")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQTreeWidgetItemFromPointer(C.QTreeWidget_TopLevelItem(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QTreeWidget) DestroyQTreeWidget() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidget::~QTreeWidget")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidget_DestroyQTreeWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
