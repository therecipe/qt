package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
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
		n.SetObjectName("QTreeWidget_" + qt.Identifier())
	}
	return n
}

func (ptr *QTreeWidget) QTreeWidget_PTR() *QTreeWidget {
	return ptr
}

func (ptr *QTreeWidget) ColumnCount() int {
	defer qt.Recovering("QTreeWidget::columnCount")

	if ptr.Pointer() != nil {
		return int(C.QTreeWidget_ColumnCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTreeWidget) SetColumnCount(columns int) {
	defer qt.Recovering("QTreeWidget::setColumnCount")

	if ptr.Pointer() != nil {
		C.QTreeWidget_SetColumnCount(ptr.Pointer(), C.int(columns))
	}
}

func (ptr *QTreeWidget) TopLevelItemCount() int {
	defer qt.Recovering("QTreeWidget::topLevelItemCount")

	if ptr.Pointer() != nil {
		return int(C.QTreeWidget_TopLevelItemCount(ptr.Pointer()))
	}
	return 0
}

func NewQTreeWidget(parent QWidget_ITF) *QTreeWidget {
	defer qt.Recovering("QTreeWidget::QTreeWidget")

	return NewQTreeWidgetFromPointer(C.QTreeWidget_NewQTreeWidget(PointerFromQWidget(parent)))
}

func (ptr *QTreeWidget) AddTopLevelItem(item QTreeWidgetItem_ITF) {
	defer qt.Recovering("QTreeWidget::addTopLevelItem")

	if ptr.Pointer() != nil {
		C.QTreeWidget_AddTopLevelItem(ptr.Pointer(), PointerFromQTreeWidgetItem(item))
	}
}

func (ptr *QTreeWidget) Clear() {
	defer qt.Recovering("QTreeWidget::clear")

	if ptr.Pointer() != nil {
		C.QTreeWidget_Clear(ptr.Pointer())
	}
}

func (ptr *QTreeWidget) ClosePersistentEditor(item QTreeWidgetItem_ITF, column int) {
	defer qt.Recovering("QTreeWidget::closePersistentEditor")

	if ptr.Pointer() != nil {
		C.QTreeWidget_ClosePersistentEditor(ptr.Pointer(), PointerFromQTreeWidgetItem(item), C.int(column))
	}
}

func (ptr *QTreeWidget) CollapseItem(item QTreeWidgetItem_ITF) {
	defer qt.Recovering("QTreeWidget::collapseItem")

	if ptr.Pointer() != nil {
		C.QTreeWidget_CollapseItem(ptr.Pointer(), PointerFromQTreeWidgetItem(item))
	}
}

func (ptr *QTreeWidget) CurrentColumn() int {
	defer qt.Recovering("QTreeWidget::currentColumn")

	if ptr.Pointer() != nil {
		return int(C.QTreeWidget_CurrentColumn(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTreeWidget) CurrentItem() *QTreeWidgetItem {
	defer qt.Recovering("QTreeWidget::currentItem")

	if ptr.Pointer() != nil {
		return NewQTreeWidgetItemFromPointer(C.QTreeWidget_CurrentItem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTreeWidget) ConnectCurrentItemChanged(f func(current *QTreeWidgetItem, previous *QTreeWidgetItem)) {
	defer qt.Recovering("connect QTreeWidget::currentItemChanged")

	if ptr.Pointer() != nil {
		C.QTreeWidget_ConnectCurrentItemChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentItemChanged", f)
	}
}

func (ptr *QTreeWidget) DisconnectCurrentItemChanged() {
	defer qt.Recovering("disconnect QTreeWidget::currentItemChanged")

	if ptr.Pointer() != nil {
		C.QTreeWidget_DisconnectCurrentItemChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentItemChanged")
	}
}

//export callbackQTreeWidgetCurrentItemChanged
func callbackQTreeWidgetCurrentItemChanged(ptrName *C.char, current unsafe.Pointer, previous unsafe.Pointer) {
	defer qt.Recovering("callback QTreeWidget::currentItemChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "currentItemChanged"); signal != nil {
		signal.(func(*QTreeWidgetItem, *QTreeWidgetItem))(NewQTreeWidgetItemFromPointer(current), NewQTreeWidgetItemFromPointer(previous))
	}

}

func (ptr *QTreeWidget) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QTreeWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QTreeWidget) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QTreeWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQTreeWidgetDropEvent
func callbackQTreeWidgetDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeWidget::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTreeWidget) EditItem(item QTreeWidgetItem_ITF, column int) {
	defer qt.Recovering("QTreeWidget::editItem")

	if ptr.Pointer() != nil {
		C.QTreeWidget_EditItem(ptr.Pointer(), PointerFromQTreeWidgetItem(item), C.int(column))
	}
}

func (ptr *QTreeWidget) ExpandItem(item QTreeWidgetItem_ITF) {
	defer qt.Recovering("QTreeWidget::expandItem")

	if ptr.Pointer() != nil {
		C.QTreeWidget_ExpandItem(ptr.Pointer(), PointerFromQTreeWidgetItem(item))
	}
}

func (ptr *QTreeWidget) HeaderItem() *QTreeWidgetItem {
	defer qt.Recovering("QTreeWidget::headerItem")

	if ptr.Pointer() != nil {
		return NewQTreeWidgetItemFromPointer(C.QTreeWidget_HeaderItem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTreeWidget) IndexOfTopLevelItem(item QTreeWidgetItem_ITF) int {
	defer qt.Recovering("QTreeWidget::indexOfTopLevelItem")

	if ptr.Pointer() != nil {
		return int(C.QTreeWidget_IndexOfTopLevelItem(ptr.Pointer(), PointerFromQTreeWidgetItem(item)))
	}
	return 0
}

func (ptr *QTreeWidget) InsertTopLevelItem(index int, item QTreeWidgetItem_ITF) {
	defer qt.Recovering("QTreeWidget::insertTopLevelItem")

	if ptr.Pointer() != nil {
		C.QTreeWidget_InsertTopLevelItem(ptr.Pointer(), C.int(index), PointerFromQTreeWidgetItem(item))
	}
}

func (ptr *QTreeWidget) InvisibleRootItem() *QTreeWidgetItem {
	defer qt.Recovering("QTreeWidget::invisibleRootItem")

	if ptr.Pointer() != nil {
		return NewQTreeWidgetItemFromPointer(C.QTreeWidget_InvisibleRootItem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTreeWidget) IsFirstItemColumnSpanned(item QTreeWidgetItem_ITF) bool {
	defer qt.Recovering("QTreeWidget::isFirstItemColumnSpanned")

	if ptr.Pointer() != nil {
		return C.QTreeWidget_IsFirstItemColumnSpanned(ptr.Pointer(), PointerFromQTreeWidgetItem(item)) != 0
	}
	return false
}

func (ptr *QTreeWidget) ItemAbove(item QTreeWidgetItem_ITF) *QTreeWidgetItem {
	defer qt.Recovering("QTreeWidget::itemAbove")

	if ptr.Pointer() != nil {
		return NewQTreeWidgetItemFromPointer(C.QTreeWidget_ItemAbove(ptr.Pointer(), PointerFromQTreeWidgetItem(item)))
	}
	return nil
}

func (ptr *QTreeWidget) ConnectItemActivated(f func(item *QTreeWidgetItem, column int)) {
	defer qt.Recovering("connect QTreeWidget::itemActivated")

	if ptr.Pointer() != nil {
		C.QTreeWidget_ConnectItemActivated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "itemActivated", f)
	}
}

func (ptr *QTreeWidget) DisconnectItemActivated() {
	defer qt.Recovering("disconnect QTreeWidget::itemActivated")

	if ptr.Pointer() != nil {
		C.QTreeWidget_DisconnectItemActivated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "itemActivated")
	}
}

//export callbackQTreeWidgetItemActivated
func callbackQTreeWidgetItemActivated(ptrName *C.char, item unsafe.Pointer, column C.int) {
	defer qt.Recovering("callback QTreeWidget::itemActivated")

	if signal := qt.GetSignal(C.GoString(ptrName), "itemActivated"); signal != nil {
		signal.(func(*QTreeWidgetItem, int))(NewQTreeWidgetItemFromPointer(item), int(column))
	}

}

func (ptr *QTreeWidget) ItemAt(p core.QPoint_ITF) *QTreeWidgetItem {
	defer qt.Recovering("QTreeWidget::itemAt")

	if ptr.Pointer() != nil {
		return NewQTreeWidgetItemFromPointer(C.QTreeWidget_ItemAt(ptr.Pointer(), core.PointerFromQPoint(p)))
	}
	return nil
}

func (ptr *QTreeWidget) ItemAt2(x int, y int) *QTreeWidgetItem {
	defer qt.Recovering("QTreeWidget::itemAt")

	if ptr.Pointer() != nil {
		return NewQTreeWidgetItemFromPointer(C.QTreeWidget_ItemAt2(ptr.Pointer(), C.int(x), C.int(y)))
	}
	return nil
}

func (ptr *QTreeWidget) ItemBelow(item QTreeWidgetItem_ITF) *QTreeWidgetItem {
	defer qt.Recovering("QTreeWidget::itemBelow")

	if ptr.Pointer() != nil {
		return NewQTreeWidgetItemFromPointer(C.QTreeWidget_ItemBelow(ptr.Pointer(), PointerFromQTreeWidgetItem(item)))
	}
	return nil
}

func (ptr *QTreeWidget) ConnectItemChanged(f func(item *QTreeWidgetItem, column int)) {
	defer qt.Recovering("connect QTreeWidget::itemChanged")

	if ptr.Pointer() != nil {
		C.QTreeWidget_ConnectItemChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "itemChanged", f)
	}
}

func (ptr *QTreeWidget) DisconnectItemChanged() {
	defer qt.Recovering("disconnect QTreeWidget::itemChanged")

	if ptr.Pointer() != nil {
		C.QTreeWidget_DisconnectItemChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "itemChanged")
	}
}

//export callbackQTreeWidgetItemChanged
func callbackQTreeWidgetItemChanged(ptrName *C.char, item unsafe.Pointer, column C.int) {
	defer qt.Recovering("callback QTreeWidget::itemChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "itemChanged"); signal != nil {
		signal.(func(*QTreeWidgetItem, int))(NewQTreeWidgetItemFromPointer(item), int(column))
	}

}

func (ptr *QTreeWidget) ConnectItemClicked(f func(item *QTreeWidgetItem, column int)) {
	defer qt.Recovering("connect QTreeWidget::itemClicked")

	if ptr.Pointer() != nil {
		C.QTreeWidget_ConnectItemClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "itemClicked", f)
	}
}

func (ptr *QTreeWidget) DisconnectItemClicked() {
	defer qt.Recovering("disconnect QTreeWidget::itemClicked")

	if ptr.Pointer() != nil {
		C.QTreeWidget_DisconnectItemClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "itemClicked")
	}
}

//export callbackQTreeWidgetItemClicked
func callbackQTreeWidgetItemClicked(ptrName *C.char, item unsafe.Pointer, column C.int) {
	defer qt.Recovering("callback QTreeWidget::itemClicked")

	if signal := qt.GetSignal(C.GoString(ptrName), "itemClicked"); signal != nil {
		signal.(func(*QTreeWidgetItem, int))(NewQTreeWidgetItemFromPointer(item), int(column))
	}

}

func (ptr *QTreeWidget) ConnectItemCollapsed(f func(item *QTreeWidgetItem)) {
	defer qt.Recovering("connect QTreeWidget::itemCollapsed")

	if ptr.Pointer() != nil {
		C.QTreeWidget_ConnectItemCollapsed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "itemCollapsed", f)
	}
}

func (ptr *QTreeWidget) DisconnectItemCollapsed() {
	defer qt.Recovering("disconnect QTreeWidget::itemCollapsed")

	if ptr.Pointer() != nil {
		C.QTreeWidget_DisconnectItemCollapsed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "itemCollapsed")
	}
}

//export callbackQTreeWidgetItemCollapsed
func callbackQTreeWidgetItemCollapsed(ptrName *C.char, item unsafe.Pointer) {
	defer qt.Recovering("callback QTreeWidget::itemCollapsed")

	if signal := qt.GetSignal(C.GoString(ptrName), "itemCollapsed"); signal != nil {
		signal.(func(*QTreeWidgetItem))(NewQTreeWidgetItemFromPointer(item))
	}

}

func (ptr *QTreeWidget) ConnectItemDoubleClicked(f func(item *QTreeWidgetItem, column int)) {
	defer qt.Recovering("connect QTreeWidget::itemDoubleClicked")

	if ptr.Pointer() != nil {
		C.QTreeWidget_ConnectItemDoubleClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "itemDoubleClicked", f)
	}
}

func (ptr *QTreeWidget) DisconnectItemDoubleClicked() {
	defer qt.Recovering("disconnect QTreeWidget::itemDoubleClicked")

	if ptr.Pointer() != nil {
		C.QTreeWidget_DisconnectItemDoubleClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "itemDoubleClicked")
	}
}

//export callbackQTreeWidgetItemDoubleClicked
func callbackQTreeWidgetItemDoubleClicked(ptrName *C.char, item unsafe.Pointer, column C.int) {
	defer qt.Recovering("callback QTreeWidget::itemDoubleClicked")

	if signal := qt.GetSignal(C.GoString(ptrName), "itemDoubleClicked"); signal != nil {
		signal.(func(*QTreeWidgetItem, int))(NewQTreeWidgetItemFromPointer(item), int(column))
	}

}

func (ptr *QTreeWidget) ConnectItemEntered(f func(item *QTreeWidgetItem, column int)) {
	defer qt.Recovering("connect QTreeWidget::itemEntered")

	if ptr.Pointer() != nil {
		C.QTreeWidget_ConnectItemEntered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "itemEntered", f)
	}
}

func (ptr *QTreeWidget) DisconnectItemEntered() {
	defer qt.Recovering("disconnect QTreeWidget::itemEntered")

	if ptr.Pointer() != nil {
		C.QTreeWidget_DisconnectItemEntered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "itemEntered")
	}
}

//export callbackQTreeWidgetItemEntered
func callbackQTreeWidgetItemEntered(ptrName *C.char, item unsafe.Pointer, column C.int) {
	defer qt.Recovering("callback QTreeWidget::itemEntered")

	if signal := qt.GetSignal(C.GoString(ptrName), "itemEntered"); signal != nil {
		signal.(func(*QTreeWidgetItem, int))(NewQTreeWidgetItemFromPointer(item), int(column))
	}

}

func (ptr *QTreeWidget) ConnectItemExpanded(f func(item *QTreeWidgetItem)) {
	defer qt.Recovering("connect QTreeWidget::itemExpanded")

	if ptr.Pointer() != nil {
		C.QTreeWidget_ConnectItemExpanded(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "itemExpanded", f)
	}
}

func (ptr *QTreeWidget) DisconnectItemExpanded() {
	defer qt.Recovering("disconnect QTreeWidget::itemExpanded")

	if ptr.Pointer() != nil {
		C.QTreeWidget_DisconnectItemExpanded(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "itemExpanded")
	}
}

//export callbackQTreeWidgetItemExpanded
func callbackQTreeWidgetItemExpanded(ptrName *C.char, item unsafe.Pointer) {
	defer qt.Recovering("callback QTreeWidget::itemExpanded")

	if signal := qt.GetSignal(C.GoString(ptrName), "itemExpanded"); signal != nil {
		signal.(func(*QTreeWidgetItem))(NewQTreeWidgetItemFromPointer(item))
	}

}

func (ptr *QTreeWidget) ConnectItemPressed(f func(item *QTreeWidgetItem, column int)) {
	defer qt.Recovering("connect QTreeWidget::itemPressed")

	if ptr.Pointer() != nil {
		C.QTreeWidget_ConnectItemPressed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "itemPressed", f)
	}
}

func (ptr *QTreeWidget) DisconnectItemPressed() {
	defer qt.Recovering("disconnect QTreeWidget::itemPressed")

	if ptr.Pointer() != nil {
		C.QTreeWidget_DisconnectItemPressed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "itemPressed")
	}
}

//export callbackQTreeWidgetItemPressed
func callbackQTreeWidgetItemPressed(ptrName *C.char, item unsafe.Pointer, column C.int) {
	defer qt.Recovering("callback QTreeWidget::itemPressed")

	if signal := qt.GetSignal(C.GoString(ptrName), "itemPressed"); signal != nil {
		signal.(func(*QTreeWidgetItem, int))(NewQTreeWidgetItemFromPointer(item), int(column))
	}

}

func (ptr *QTreeWidget) ConnectItemSelectionChanged(f func()) {
	defer qt.Recovering("connect QTreeWidget::itemSelectionChanged")

	if ptr.Pointer() != nil {
		C.QTreeWidget_ConnectItemSelectionChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "itemSelectionChanged", f)
	}
}

func (ptr *QTreeWidget) DisconnectItemSelectionChanged() {
	defer qt.Recovering("disconnect QTreeWidget::itemSelectionChanged")

	if ptr.Pointer() != nil {
		C.QTreeWidget_DisconnectItemSelectionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "itemSelectionChanged")
	}
}

//export callbackQTreeWidgetItemSelectionChanged
func callbackQTreeWidgetItemSelectionChanged(ptrName *C.char) {
	defer qt.Recovering("callback QTreeWidget::itemSelectionChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "itemSelectionChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QTreeWidget) ItemWidget(item QTreeWidgetItem_ITF, column int) *QWidget {
	defer qt.Recovering("QTreeWidget::itemWidget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QTreeWidget_ItemWidget(ptr.Pointer(), PointerFromQTreeWidgetItem(item), C.int(column)))
	}
	return nil
}

func (ptr *QTreeWidget) OpenPersistentEditor(item QTreeWidgetItem_ITF, column int) {
	defer qt.Recovering("QTreeWidget::openPersistentEditor")

	if ptr.Pointer() != nil {
		C.QTreeWidget_OpenPersistentEditor(ptr.Pointer(), PointerFromQTreeWidgetItem(item), C.int(column))
	}
}

func (ptr *QTreeWidget) RemoveItemWidget(item QTreeWidgetItem_ITF, column int) {
	defer qt.Recovering("QTreeWidget::removeItemWidget")

	if ptr.Pointer() != nil {
		C.QTreeWidget_RemoveItemWidget(ptr.Pointer(), PointerFromQTreeWidgetItem(item), C.int(column))
	}
}

func (ptr *QTreeWidget) ScrollToItem(item QTreeWidgetItem_ITF, hint QAbstractItemView__ScrollHint) {
	defer qt.Recovering("QTreeWidget::scrollToItem")

	if ptr.Pointer() != nil {
		C.QTreeWidget_ScrollToItem(ptr.Pointer(), PointerFromQTreeWidgetItem(item), C.int(hint))
	}
}

func (ptr *QTreeWidget) SetCurrentItem(item QTreeWidgetItem_ITF) {
	defer qt.Recovering("QTreeWidget::setCurrentItem")

	if ptr.Pointer() != nil {
		C.QTreeWidget_SetCurrentItem(ptr.Pointer(), PointerFromQTreeWidgetItem(item))
	}
}

func (ptr *QTreeWidget) SetCurrentItem2(item QTreeWidgetItem_ITF, column int) {
	defer qt.Recovering("QTreeWidget::setCurrentItem")

	if ptr.Pointer() != nil {
		C.QTreeWidget_SetCurrentItem2(ptr.Pointer(), PointerFromQTreeWidgetItem(item), C.int(column))
	}
}

func (ptr *QTreeWidget) SetCurrentItem3(item QTreeWidgetItem_ITF, column int, command core.QItemSelectionModel__SelectionFlag) {
	defer qt.Recovering("QTreeWidget::setCurrentItem")

	if ptr.Pointer() != nil {
		C.QTreeWidget_SetCurrentItem3(ptr.Pointer(), PointerFromQTreeWidgetItem(item), C.int(column), C.int(command))
	}
}

func (ptr *QTreeWidget) SetFirstItemColumnSpanned(item QTreeWidgetItem_ITF, span bool) {
	defer qt.Recovering("QTreeWidget::setFirstItemColumnSpanned")

	if ptr.Pointer() != nil {
		C.QTreeWidget_SetFirstItemColumnSpanned(ptr.Pointer(), PointerFromQTreeWidgetItem(item), C.int(qt.GoBoolToInt(span)))
	}
}

func (ptr *QTreeWidget) SetHeaderItem(item QTreeWidgetItem_ITF) {
	defer qt.Recovering("QTreeWidget::setHeaderItem")

	if ptr.Pointer() != nil {
		C.QTreeWidget_SetHeaderItem(ptr.Pointer(), PointerFromQTreeWidgetItem(item))
	}
}

func (ptr *QTreeWidget) SetHeaderLabel(label string) {
	defer qt.Recovering("QTreeWidget::setHeaderLabel")

	if ptr.Pointer() != nil {
		C.QTreeWidget_SetHeaderLabel(ptr.Pointer(), C.CString(label))
	}
}

func (ptr *QTreeWidget) SetHeaderLabels(labels []string) {
	defer qt.Recovering("QTreeWidget::setHeaderLabels")

	if ptr.Pointer() != nil {
		C.QTreeWidget_SetHeaderLabels(ptr.Pointer(), C.CString(strings.Join(labels, ",,,")))
	}
}

func (ptr *QTreeWidget) SetItemWidget(item QTreeWidgetItem_ITF, column int, widget QWidget_ITF) {
	defer qt.Recovering("QTreeWidget::setItemWidget")

	if ptr.Pointer() != nil {
		C.QTreeWidget_SetItemWidget(ptr.Pointer(), PointerFromQTreeWidgetItem(item), C.int(column), PointerFromQWidget(widget))
	}
}

func (ptr *QTreeWidget) ConnectSetSelectionModel(f func(selectionModel *core.QItemSelectionModel)) {
	defer qt.Recovering("connect QTreeWidget::setSelectionModel")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setSelectionModel", f)
	}
}

func (ptr *QTreeWidget) DisconnectSetSelectionModel() {
	defer qt.Recovering("disconnect QTreeWidget::setSelectionModel")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setSelectionModel")
	}
}

//export callbackQTreeWidgetSetSelectionModel
func callbackQTreeWidgetSetSelectionModel(ptrName *C.char, selectionModel unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeWidget::setSelectionModel")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSelectionModel"); signal != nil {
		signal.(func(*core.QItemSelectionModel))(core.NewQItemSelectionModelFromPointer(selectionModel))
		return true
	}
	return false

}

func (ptr *QTreeWidget) SortColumn() int {
	defer qt.Recovering("QTreeWidget::sortColumn")

	if ptr.Pointer() != nil {
		return int(C.QTreeWidget_SortColumn(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTreeWidget) SortItems(column int, order core.Qt__SortOrder) {
	defer qt.Recovering("QTreeWidget::sortItems")

	if ptr.Pointer() != nil {
		C.QTreeWidget_SortItems(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QTreeWidget) TakeTopLevelItem(index int) *QTreeWidgetItem {
	defer qt.Recovering("QTreeWidget::takeTopLevelItem")

	if ptr.Pointer() != nil {
		return NewQTreeWidgetItemFromPointer(C.QTreeWidget_TakeTopLevelItem(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QTreeWidget) TopLevelItem(index int) *QTreeWidgetItem {
	defer qt.Recovering("QTreeWidget::topLevelItem")

	if ptr.Pointer() != nil {
		return NewQTreeWidgetItemFromPointer(C.QTreeWidget_TopLevelItem(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QTreeWidget) VisualItemRect(item QTreeWidgetItem_ITF) *core.QRect {
	defer qt.Recovering("QTreeWidget::visualItemRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QTreeWidget_VisualItemRect(ptr.Pointer(), PointerFromQTreeWidgetItem(item)))
	}
	return nil
}

func (ptr *QTreeWidget) DestroyQTreeWidget() {
	defer qt.Recovering("QTreeWidget::~QTreeWidget")

	if ptr.Pointer() != nil {
		C.QTreeWidget_DestroyQTreeWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QTreeWidget) ConnectCurrentChanged(f func(current *core.QModelIndex, previous *core.QModelIndex)) {
	defer qt.Recovering("connect QTreeWidget::currentChanged")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "currentChanged", f)
	}
}

func (ptr *QTreeWidget) DisconnectCurrentChanged() {
	defer qt.Recovering("disconnect QTreeWidget::currentChanged")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "currentChanged")
	}
}

//export callbackQTreeWidgetCurrentChanged
func callbackQTreeWidgetCurrentChanged(ptrName *C.char, current unsafe.Pointer, previous unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeWidget::currentChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "currentChanged"); signal != nil {
		signal.(func(*core.QModelIndex, *core.QModelIndex))(core.NewQModelIndexFromPointer(current), core.NewQModelIndexFromPointer(previous))
		return true
	}
	return false

}

func (ptr *QTreeWidget) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QTreeWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QTreeWidget) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QTreeWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQTreeWidgetDragMoveEvent
func callbackQTreeWidgetDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeWidget::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTreeWidget) ConnectDrawBranches(f func(painter *gui.QPainter, rect *core.QRect, index *core.QModelIndex)) {
	defer qt.Recovering("connect QTreeWidget::drawBranches")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "drawBranches", f)
	}
}

func (ptr *QTreeWidget) DisconnectDrawBranches() {
	defer qt.Recovering("disconnect QTreeWidget::drawBranches")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "drawBranches")
	}
}

//export callbackQTreeWidgetDrawBranches
func callbackQTreeWidgetDrawBranches(ptrName *C.char, painter unsafe.Pointer, rect unsafe.Pointer, index unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeWidget::drawBranches")

	if signal := qt.GetSignal(C.GoString(ptrName), "drawBranches"); signal != nil {
		signal.(func(*gui.QPainter, *core.QRect, *core.QModelIndex))(gui.NewQPainterFromPointer(painter), core.NewQRectFromPointer(rect), core.NewQModelIndexFromPointer(index))
		return true
	}
	return false

}

func (ptr *QTreeWidget) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QTreeWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QTreeWidget) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QTreeWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQTreeWidgetKeyPressEvent
func callbackQTreeWidgetKeyPressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeWidget::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTreeWidget) ConnectKeyboardSearch(f func(search string)) {
	defer qt.Recovering("connect QTreeWidget::keyboardSearch")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyboardSearch", f)
	}
}

func (ptr *QTreeWidget) DisconnectKeyboardSearch() {
	defer qt.Recovering("disconnect QTreeWidget::keyboardSearch")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyboardSearch")
	}
}

//export callbackQTreeWidgetKeyboardSearch
func callbackQTreeWidgetKeyboardSearch(ptrName *C.char, search *C.char) bool {
	defer qt.Recovering("callback QTreeWidget::keyboardSearch")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyboardSearch"); signal != nil {
		signal.(func(string))(C.GoString(search))
		return true
	}
	return false

}

func (ptr *QTreeWidget) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QTreeWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QTreeWidget) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QTreeWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQTreeWidgetMouseDoubleClickEvent
func callbackQTreeWidgetMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeWidget::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTreeWidget) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QTreeWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QTreeWidget) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QTreeWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQTreeWidgetMouseMoveEvent
func callbackQTreeWidgetMouseMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeWidget::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTreeWidget) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QTreeWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QTreeWidget) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QTreeWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQTreeWidgetMousePressEvent
func callbackQTreeWidgetMousePressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeWidget::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTreeWidget) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QTreeWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QTreeWidget) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QTreeWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQTreeWidgetMouseReleaseEvent
func callbackQTreeWidgetMouseReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeWidget::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTreeWidget) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QTreeWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QTreeWidget) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QTreeWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQTreeWidgetPaintEvent
func callbackQTreeWidgetPaintEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeWidget::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTreeWidget) ConnectReset(f func()) {
	defer qt.Recovering("connect QTreeWidget::reset")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "reset", f)
	}
}

func (ptr *QTreeWidget) DisconnectReset() {
	defer qt.Recovering("disconnect QTreeWidget::reset")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "reset")
	}
}

//export callbackQTreeWidgetReset
func callbackQTreeWidgetReset(ptrName *C.char) bool {
	defer qt.Recovering("callback QTreeWidget::reset")

	if signal := qt.GetSignal(C.GoString(ptrName), "reset"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QTreeWidget) ConnectRowsAboutToBeRemoved(f func(parent *core.QModelIndex, start int, end int)) {
	defer qt.Recovering("connect QTreeWidget::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "rowsAboutToBeRemoved", f)
	}
}

func (ptr *QTreeWidget) DisconnectRowsAboutToBeRemoved() {
	defer qt.Recovering("disconnect QTreeWidget::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "rowsAboutToBeRemoved")
	}
}

//export callbackQTreeWidgetRowsAboutToBeRemoved
func callbackQTreeWidgetRowsAboutToBeRemoved(ptrName *C.char, parent unsafe.Pointer, start C.int, end C.int) bool {
	defer qt.Recovering("callback QTreeWidget::rowsAboutToBeRemoved")

	if signal := qt.GetSignal(C.GoString(ptrName), "rowsAboutToBeRemoved"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(start), int(end))
		return true
	}
	return false

}

func (ptr *QTreeWidget) ConnectRowsInserted(f func(parent *core.QModelIndex, start int, end int)) {
	defer qt.Recovering("connect QTreeWidget::rowsInserted")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "rowsInserted", f)
	}
}

func (ptr *QTreeWidget) DisconnectRowsInserted() {
	defer qt.Recovering("disconnect QTreeWidget::rowsInserted")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "rowsInserted")
	}
}

//export callbackQTreeWidgetRowsInserted
func callbackQTreeWidgetRowsInserted(ptrName *C.char, parent unsafe.Pointer, start C.int, end C.int) bool {
	defer qt.Recovering("callback QTreeWidget::rowsInserted")

	if signal := qt.GetSignal(C.GoString(ptrName), "rowsInserted"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(start), int(end))
		return true
	}
	return false

}

func (ptr *QTreeWidget) ConnectScrollContentsBy(f func(dx int, dy int)) {
	defer qt.Recovering("connect QTreeWidget::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "scrollContentsBy", f)
	}
}

func (ptr *QTreeWidget) DisconnectScrollContentsBy() {
	defer qt.Recovering("disconnect QTreeWidget::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "scrollContentsBy")
	}
}

//export callbackQTreeWidgetScrollContentsBy
func callbackQTreeWidgetScrollContentsBy(ptrName *C.char, dx C.int, dy C.int) bool {
	defer qt.Recovering("callback QTreeWidget::scrollContentsBy")

	if signal := qt.GetSignal(C.GoString(ptrName), "scrollContentsBy"); signal != nil {
		signal.(func(int, int))(int(dx), int(dy))
		return true
	}
	return false

}

func (ptr *QTreeWidget) ConnectScrollTo(f func(index *core.QModelIndex, hint QAbstractItemView__ScrollHint)) {
	defer qt.Recovering("connect QTreeWidget::scrollTo")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "scrollTo", f)
	}
}

func (ptr *QTreeWidget) DisconnectScrollTo() {
	defer qt.Recovering("disconnect QTreeWidget::scrollTo")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "scrollTo")
	}
}

//export callbackQTreeWidgetScrollTo
func callbackQTreeWidgetScrollTo(ptrName *C.char, index unsafe.Pointer, hint C.int) bool {
	defer qt.Recovering("callback QTreeWidget::scrollTo")

	if signal := qt.GetSignal(C.GoString(ptrName), "scrollTo"); signal != nil {
		signal.(func(*core.QModelIndex, QAbstractItemView__ScrollHint))(core.NewQModelIndexFromPointer(index), QAbstractItemView__ScrollHint(hint))
		return true
	}
	return false

}

func (ptr *QTreeWidget) ConnectSelectAll(f func()) {
	defer qt.Recovering("connect QTreeWidget::selectAll")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "selectAll", f)
	}
}

func (ptr *QTreeWidget) DisconnectSelectAll() {
	defer qt.Recovering("disconnect QTreeWidget::selectAll")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "selectAll")
	}
}

//export callbackQTreeWidgetSelectAll
func callbackQTreeWidgetSelectAll(ptrName *C.char) bool {
	defer qt.Recovering("callback QTreeWidget::selectAll")

	if signal := qt.GetSignal(C.GoString(ptrName), "selectAll"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QTreeWidget) ConnectSetRootIndex(f func(index *core.QModelIndex)) {
	defer qt.Recovering("connect QTreeWidget::setRootIndex")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setRootIndex", f)
	}
}

func (ptr *QTreeWidget) DisconnectSetRootIndex() {
	defer qt.Recovering("disconnect QTreeWidget::setRootIndex")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setRootIndex")
	}
}

//export callbackQTreeWidgetSetRootIndex
func callbackQTreeWidgetSetRootIndex(ptrName *C.char, index unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeWidget::setRootIndex")

	if signal := qt.GetSignal(C.GoString(ptrName), "setRootIndex"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
		return true
	}
	return false

}

func (ptr *QTreeWidget) ConnectSetSelection(f func(rect *core.QRect, command core.QItemSelectionModel__SelectionFlag)) {
	defer qt.Recovering("connect QTreeWidget::setSelection")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setSelection", f)
	}
}

func (ptr *QTreeWidget) DisconnectSetSelection() {
	defer qt.Recovering("disconnect QTreeWidget::setSelection")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setSelection")
	}
}

//export callbackQTreeWidgetSetSelection
func callbackQTreeWidgetSetSelection(ptrName *C.char, rect unsafe.Pointer, command C.int) bool {
	defer qt.Recovering("callback QTreeWidget::setSelection")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSelection"); signal != nil {
		signal.(func(*core.QRect, core.QItemSelectionModel__SelectionFlag))(core.NewQRectFromPointer(rect), core.QItemSelectionModel__SelectionFlag(command))
		return true
	}
	return false

}

func (ptr *QTreeWidget) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QTreeWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QTreeWidget) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QTreeWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQTreeWidgetTimerEvent
func callbackQTreeWidgetTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeWidget::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTreeWidget) ConnectUpdateGeometries(f func()) {
	defer qt.Recovering("connect QTreeWidget::updateGeometries")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updateGeometries", f)
	}
}

func (ptr *QTreeWidget) DisconnectUpdateGeometries() {
	defer qt.Recovering("disconnect QTreeWidget::updateGeometries")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updateGeometries")
	}
}

//export callbackQTreeWidgetUpdateGeometries
func callbackQTreeWidgetUpdateGeometries(ptrName *C.char) bool {
	defer qt.Recovering("callback QTreeWidget::updateGeometries")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateGeometries"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QTreeWidget) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QTreeWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QTreeWidget) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QTreeWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQTreeWidgetDragLeaveEvent
func callbackQTreeWidgetDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeWidget::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTreeWidget) ConnectCloseEditor(f func(editor *QWidget, hint QAbstractItemDelegate__EndEditHint)) {
	defer qt.Recovering("connect QTreeWidget::closeEditor")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEditor", f)
	}
}

func (ptr *QTreeWidget) DisconnectCloseEditor() {
	defer qt.Recovering("disconnect QTreeWidget::closeEditor")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEditor")
	}
}

//export callbackQTreeWidgetCloseEditor
func callbackQTreeWidgetCloseEditor(ptrName *C.char, editor unsafe.Pointer, hint C.int) bool {
	defer qt.Recovering("callback QTreeWidget::closeEditor")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEditor"); signal != nil {
		signal.(func(*QWidget, QAbstractItemDelegate__EndEditHint))(NewQWidgetFromPointer(editor), QAbstractItemDelegate__EndEditHint(hint))
		return true
	}
	return false

}

func (ptr *QTreeWidget) ConnectCommitData(f func(editor *QWidget)) {
	defer qt.Recovering("connect QTreeWidget::commitData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "commitData", f)
	}
}

func (ptr *QTreeWidget) DisconnectCommitData() {
	defer qt.Recovering("disconnect QTreeWidget::commitData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "commitData")
	}
}

//export callbackQTreeWidgetCommitData
func callbackQTreeWidgetCommitData(ptrName *C.char, editor unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeWidget::commitData")

	if signal := qt.GetSignal(C.GoString(ptrName), "commitData"); signal != nil {
		signal.(func(*QWidget))(NewQWidgetFromPointer(editor))
		return true
	}
	return false

}

func (ptr *QTreeWidget) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QTreeWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QTreeWidget) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QTreeWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQTreeWidgetDragEnterEvent
func callbackQTreeWidgetDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeWidget::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTreeWidget) ConnectEditorDestroyed(f func(editor *core.QObject)) {
	defer qt.Recovering("connect QTreeWidget::editorDestroyed")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "editorDestroyed", f)
	}
}

func (ptr *QTreeWidget) DisconnectEditorDestroyed() {
	defer qt.Recovering("disconnect QTreeWidget::editorDestroyed")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "editorDestroyed")
	}
}

//export callbackQTreeWidgetEditorDestroyed
func callbackQTreeWidgetEditorDestroyed(ptrName *C.char, editor unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeWidget::editorDestroyed")

	if signal := qt.GetSignal(C.GoString(ptrName), "editorDestroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(editor))
		return true
	}
	return false

}

func (ptr *QTreeWidget) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QTreeWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QTreeWidget) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QTreeWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQTreeWidgetFocusInEvent
func callbackQTreeWidgetFocusInEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeWidget::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTreeWidget) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QTreeWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QTreeWidget) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QTreeWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQTreeWidgetFocusOutEvent
func callbackQTreeWidgetFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeWidget::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTreeWidget) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QTreeWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QTreeWidget) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QTreeWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQTreeWidgetInputMethodEvent
func callbackQTreeWidgetInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeWidget::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTreeWidget) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QTreeWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QTreeWidget) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QTreeWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQTreeWidgetResizeEvent
func callbackQTreeWidgetResizeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeWidget::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTreeWidget) ConnectStartDrag(f func(supportedActions core.Qt__DropAction)) {
	defer qt.Recovering("connect QTreeWidget::startDrag")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "startDrag", f)
	}
}

func (ptr *QTreeWidget) DisconnectStartDrag() {
	defer qt.Recovering("disconnect QTreeWidget::startDrag")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "startDrag")
	}
}

//export callbackQTreeWidgetStartDrag
func callbackQTreeWidgetStartDrag(ptrName *C.char, supportedActions C.int) bool {
	defer qt.Recovering("callback QTreeWidget::startDrag")

	if signal := qt.GetSignal(C.GoString(ptrName), "startDrag"); signal != nil {
		signal.(func(core.Qt__DropAction))(core.Qt__DropAction(supportedActions))
		return true
	}
	return false

}

func (ptr *QTreeWidget) ConnectContextMenuEvent(f func(e *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QTreeWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QTreeWidget) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QTreeWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQTreeWidgetContextMenuEvent
func callbackQTreeWidgetContextMenuEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeWidget::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QTreeWidget) ConnectSetupViewport(f func(viewport *QWidget)) {
	defer qt.Recovering("connect QTreeWidget::setupViewport")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setupViewport", f)
	}
}

func (ptr *QTreeWidget) DisconnectSetupViewport() {
	defer qt.Recovering("disconnect QTreeWidget::setupViewport")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setupViewport")
	}
}

//export callbackQTreeWidgetSetupViewport
func callbackQTreeWidgetSetupViewport(ptrName *C.char, viewport unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeWidget::setupViewport")

	if signal := qt.GetSignal(C.GoString(ptrName), "setupViewport"); signal != nil {
		signal.(func(*QWidget))(NewQWidgetFromPointer(viewport))
		return true
	}
	return false

}

func (ptr *QTreeWidget) ConnectWheelEvent(f func(e *gui.QWheelEvent)) {
	defer qt.Recovering("connect QTreeWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QTreeWidget) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QTreeWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQTreeWidgetWheelEvent
func callbackQTreeWidgetWheelEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeWidget::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QTreeWidget) ConnectChangeEvent(f func(ev *core.QEvent)) {
	defer qt.Recovering("connect QTreeWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QTreeWidget) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QTreeWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQTreeWidgetChangeEvent
func callbackQTreeWidgetChangeEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeWidget::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QTreeWidget) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QTreeWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QTreeWidget) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QTreeWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQTreeWidgetActionEvent
func callbackQTreeWidgetActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeWidget::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTreeWidget) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QTreeWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QTreeWidget) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QTreeWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQTreeWidgetEnterEvent
func callbackQTreeWidgetEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeWidget::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTreeWidget) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QTreeWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QTreeWidget) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QTreeWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQTreeWidgetHideEvent
func callbackQTreeWidgetHideEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeWidget::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTreeWidget) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QTreeWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QTreeWidget) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QTreeWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQTreeWidgetLeaveEvent
func callbackQTreeWidgetLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeWidget::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTreeWidget) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QTreeWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QTreeWidget) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QTreeWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQTreeWidgetMoveEvent
func callbackQTreeWidgetMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeWidget::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTreeWidget) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QTreeWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QTreeWidget) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QTreeWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQTreeWidgetSetVisible
func callbackQTreeWidgetSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QTreeWidget::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QTreeWidget) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QTreeWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QTreeWidget) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QTreeWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQTreeWidgetShowEvent
func callbackQTreeWidgetShowEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeWidget::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTreeWidget) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QTreeWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QTreeWidget) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QTreeWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQTreeWidgetCloseEvent
func callbackQTreeWidgetCloseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeWidget::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTreeWidget) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QTreeWidget::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QTreeWidget) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QTreeWidget::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQTreeWidgetInitPainter
func callbackQTreeWidgetInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeWidget::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

}

func (ptr *QTreeWidget) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QTreeWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QTreeWidget) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QTreeWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQTreeWidgetKeyReleaseEvent
func callbackQTreeWidgetKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeWidget::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTreeWidget) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QTreeWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QTreeWidget) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QTreeWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQTreeWidgetTabletEvent
func callbackQTreeWidgetTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeWidget::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTreeWidget) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QTreeWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QTreeWidget) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QTreeWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQTreeWidgetChildEvent
func callbackQTreeWidgetChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeWidget::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTreeWidget) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QTreeWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QTreeWidget) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QTreeWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQTreeWidgetCustomEvent
func callbackQTreeWidgetCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeWidget::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
