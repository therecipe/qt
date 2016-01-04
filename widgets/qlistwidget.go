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

type QListWidget struct {
	QListView
}

type QListWidget_ITF interface {
	QListView_ITF
	QListWidget_PTR() *QListWidget
}

func PointerFromQListWidget(ptr QListWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QListWidget_PTR().Pointer()
	}
	return nil
}

func NewQListWidgetFromPointer(ptr unsafe.Pointer) *QListWidget {
	var n = new(QListWidget)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QListWidget_") {
		n.SetObjectName("QListWidget_" + qt.Identifier())
	}
	return n
}

func (ptr *QListWidget) QListWidget_PTR() *QListWidget {
	return ptr
}

func (ptr *QListWidget) Count() int {
	defer qt.Recovering("QListWidget::count")

	if ptr.Pointer() != nil {
		return int(C.QListWidget_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QListWidget) CurrentRow() int {
	defer qt.Recovering("QListWidget::currentRow")

	if ptr.Pointer() != nil {
		return int(C.QListWidget_CurrentRow(ptr.Pointer()))
	}
	return 0
}

func (ptr *QListWidget) IsSortingEnabled() bool {
	defer qt.Recovering("QListWidget::isSortingEnabled")

	if ptr.Pointer() != nil {
		return C.QListWidget_IsSortingEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QListWidget) SetCurrentRow(row int) {
	defer qt.Recovering("QListWidget::setCurrentRow")

	if ptr.Pointer() != nil {
		C.QListWidget_SetCurrentRow(ptr.Pointer(), C.int(row))
	}
}

func (ptr *QListWidget) SetSortingEnabled(enable bool) {
	defer qt.Recovering("QListWidget::setSortingEnabled")

	if ptr.Pointer() != nil {
		C.QListWidget_SetSortingEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func NewQListWidget(parent QWidget_ITF) *QListWidget {
	defer qt.Recovering("QListWidget::QListWidget")

	return NewQListWidgetFromPointer(C.QListWidget_NewQListWidget(PointerFromQWidget(parent)))
}

func (ptr *QListWidget) AddItem2(item QListWidgetItem_ITF) {
	defer qt.Recovering("QListWidget::addItem")

	if ptr.Pointer() != nil {
		C.QListWidget_AddItem2(ptr.Pointer(), PointerFromQListWidgetItem(item))
	}
}

func (ptr *QListWidget) AddItem(label string) {
	defer qt.Recovering("QListWidget::addItem")

	if ptr.Pointer() != nil {
		C.QListWidget_AddItem(ptr.Pointer(), C.CString(label))
	}
}

func (ptr *QListWidget) AddItems(labels []string) {
	defer qt.Recovering("QListWidget::addItems")

	if ptr.Pointer() != nil {
		C.QListWidget_AddItems(ptr.Pointer(), C.CString(strings.Join(labels, ",,,")))
	}
}

func (ptr *QListWidget) Clear() {
	defer qt.Recovering("QListWidget::clear")

	if ptr.Pointer() != nil {
		C.QListWidget_Clear(ptr.Pointer())
	}
}

func (ptr *QListWidget) ClosePersistentEditor(item QListWidgetItem_ITF) {
	defer qt.Recovering("QListWidget::closePersistentEditor")

	if ptr.Pointer() != nil {
		C.QListWidget_ClosePersistentEditor(ptr.Pointer(), PointerFromQListWidgetItem(item))
	}
}

func (ptr *QListWidget) CurrentItem() *QListWidgetItem {
	defer qt.Recovering("QListWidget::currentItem")

	if ptr.Pointer() != nil {
		return NewQListWidgetItemFromPointer(C.QListWidget_CurrentItem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QListWidget) ConnectCurrentItemChanged(f func(current *QListWidgetItem, previous *QListWidgetItem)) {
	defer qt.Recovering("connect QListWidget::currentItemChanged")

	if ptr.Pointer() != nil {
		C.QListWidget_ConnectCurrentItemChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentItemChanged", f)
	}
}

func (ptr *QListWidget) DisconnectCurrentItemChanged() {
	defer qt.Recovering("disconnect QListWidget::currentItemChanged")

	if ptr.Pointer() != nil {
		C.QListWidget_DisconnectCurrentItemChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentItemChanged")
	}
}

//export callbackQListWidgetCurrentItemChanged
func callbackQListWidgetCurrentItemChanged(ptr unsafe.Pointer, ptrName *C.char, current unsafe.Pointer, previous unsafe.Pointer) {
	defer qt.Recovering("callback QListWidget::currentItemChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "currentItemChanged"); signal != nil {
		signal.(func(*QListWidgetItem, *QListWidgetItem))(NewQListWidgetItemFromPointer(current), NewQListWidgetItemFromPointer(previous))
	}

}

func (ptr *QListWidget) CurrentItemChanged(current QListWidgetItem_ITF, previous QListWidgetItem_ITF) {
	defer qt.Recovering("QListWidget::currentItemChanged")

	if ptr.Pointer() != nil {
		C.QListWidget_CurrentItemChanged(ptr.Pointer(), PointerFromQListWidgetItem(current), PointerFromQListWidgetItem(previous))
	}
}

func (ptr *QListWidget) ConnectCurrentRowChanged(f func(currentRow int)) {
	defer qt.Recovering("connect QListWidget::currentRowChanged")

	if ptr.Pointer() != nil {
		C.QListWidget_ConnectCurrentRowChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentRowChanged", f)
	}
}

func (ptr *QListWidget) DisconnectCurrentRowChanged() {
	defer qt.Recovering("disconnect QListWidget::currentRowChanged")

	if ptr.Pointer() != nil {
		C.QListWidget_DisconnectCurrentRowChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentRowChanged")
	}
}

//export callbackQListWidgetCurrentRowChanged
func callbackQListWidgetCurrentRowChanged(ptr unsafe.Pointer, ptrName *C.char, currentRow C.int) {
	defer qt.Recovering("callback QListWidget::currentRowChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "currentRowChanged"); signal != nil {
		signal.(func(int))(int(currentRow))
	}

}

func (ptr *QListWidget) CurrentRowChanged(currentRow int) {
	defer qt.Recovering("QListWidget::currentRowChanged")

	if ptr.Pointer() != nil {
		C.QListWidget_CurrentRowChanged(ptr.Pointer(), C.int(currentRow))
	}
}

func (ptr *QListWidget) ConnectCurrentTextChanged(f func(currentText string)) {
	defer qt.Recovering("connect QListWidget::currentTextChanged")

	if ptr.Pointer() != nil {
		C.QListWidget_ConnectCurrentTextChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentTextChanged", f)
	}
}

func (ptr *QListWidget) DisconnectCurrentTextChanged() {
	defer qt.Recovering("disconnect QListWidget::currentTextChanged")

	if ptr.Pointer() != nil {
		C.QListWidget_DisconnectCurrentTextChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentTextChanged")
	}
}

//export callbackQListWidgetCurrentTextChanged
func callbackQListWidgetCurrentTextChanged(ptr unsafe.Pointer, ptrName *C.char, currentText *C.char) {
	defer qt.Recovering("callback QListWidget::currentTextChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "currentTextChanged"); signal != nil {
		signal.(func(string))(C.GoString(currentText))
	}

}

func (ptr *QListWidget) CurrentTextChanged(currentText string) {
	defer qt.Recovering("QListWidget::currentTextChanged")

	if ptr.Pointer() != nil {
		C.QListWidget_CurrentTextChanged(ptr.Pointer(), C.CString(currentText))
	}
}

func (ptr *QListWidget) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QListWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QListWidget) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QListWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQListWidgetDropEvent
func callbackQListWidgetDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QListWidget::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQListWidgetFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QListWidget) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QListWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QListWidget) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QListWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QListWidget) DropMimeData(index int, data core.QMimeData_ITF, action core.Qt__DropAction) bool {
	defer qt.Recovering("QListWidget::dropMimeData")

	if ptr.Pointer() != nil {
		return C.QListWidget_DropMimeData(ptr.Pointer(), C.int(index), core.PointerFromQMimeData(data), C.int(action)) != 0
	}
	return false
}

func (ptr *QListWidget) EditItem(item QListWidgetItem_ITF) {
	defer qt.Recovering("QListWidget::editItem")

	if ptr.Pointer() != nil {
		C.QListWidget_EditItem(ptr.Pointer(), PointerFromQListWidgetItem(item))
	}
}

func (ptr *QListWidget) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QListWidget::event")

	if ptr.Pointer() != nil {
		return C.QListWidget_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QListWidget) InsertItem(row int, item QListWidgetItem_ITF) {
	defer qt.Recovering("QListWidget::insertItem")

	if ptr.Pointer() != nil {
		C.QListWidget_InsertItem(ptr.Pointer(), C.int(row), PointerFromQListWidgetItem(item))
	}
}

func (ptr *QListWidget) InsertItem2(row int, label string) {
	defer qt.Recovering("QListWidget::insertItem")

	if ptr.Pointer() != nil {
		C.QListWidget_InsertItem2(ptr.Pointer(), C.int(row), C.CString(label))
	}
}

func (ptr *QListWidget) InsertItems(row int, labels []string) {
	defer qt.Recovering("QListWidget::insertItems")

	if ptr.Pointer() != nil {
		C.QListWidget_InsertItems(ptr.Pointer(), C.int(row), C.CString(strings.Join(labels, ",,,")))
	}
}

func (ptr *QListWidget) Item(row int) *QListWidgetItem {
	defer qt.Recovering("QListWidget::item")

	if ptr.Pointer() != nil {
		return NewQListWidgetItemFromPointer(C.QListWidget_Item(ptr.Pointer(), C.int(row)))
	}
	return nil
}

func (ptr *QListWidget) ConnectItemActivated(f func(item *QListWidgetItem)) {
	defer qt.Recovering("connect QListWidget::itemActivated")

	if ptr.Pointer() != nil {
		C.QListWidget_ConnectItemActivated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "itemActivated", f)
	}
}

func (ptr *QListWidget) DisconnectItemActivated() {
	defer qt.Recovering("disconnect QListWidget::itemActivated")

	if ptr.Pointer() != nil {
		C.QListWidget_DisconnectItemActivated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "itemActivated")
	}
}

//export callbackQListWidgetItemActivated
func callbackQListWidgetItemActivated(ptr unsafe.Pointer, ptrName *C.char, item unsafe.Pointer) {
	defer qt.Recovering("callback QListWidget::itemActivated")

	if signal := qt.GetSignal(C.GoString(ptrName), "itemActivated"); signal != nil {
		signal.(func(*QListWidgetItem))(NewQListWidgetItemFromPointer(item))
	}

}

func (ptr *QListWidget) ItemActivated(item QListWidgetItem_ITF) {
	defer qt.Recovering("QListWidget::itemActivated")

	if ptr.Pointer() != nil {
		C.QListWidget_ItemActivated(ptr.Pointer(), PointerFromQListWidgetItem(item))
	}
}

func (ptr *QListWidget) ItemAt(p core.QPoint_ITF) *QListWidgetItem {
	defer qt.Recovering("QListWidget::itemAt")

	if ptr.Pointer() != nil {
		return NewQListWidgetItemFromPointer(C.QListWidget_ItemAt(ptr.Pointer(), core.PointerFromQPoint(p)))
	}
	return nil
}

func (ptr *QListWidget) ItemAt2(x int, y int) *QListWidgetItem {
	defer qt.Recovering("QListWidget::itemAt")

	if ptr.Pointer() != nil {
		return NewQListWidgetItemFromPointer(C.QListWidget_ItemAt2(ptr.Pointer(), C.int(x), C.int(y)))
	}
	return nil
}

func (ptr *QListWidget) ConnectItemChanged(f func(item *QListWidgetItem)) {
	defer qt.Recovering("connect QListWidget::itemChanged")

	if ptr.Pointer() != nil {
		C.QListWidget_ConnectItemChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "itemChanged", f)
	}
}

func (ptr *QListWidget) DisconnectItemChanged() {
	defer qt.Recovering("disconnect QListWidget::itemChanged")

	if ptr.Pointer() != nil {
		C.QListWidget_DisconnectItemChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "itemChanged")
	}
}

//export callbackQListWidgetItemChanged
func callbackQListWidgetItemChanged(ptr unsafe.Pointer, ptrName *C.char, item unsafe.Pointer) {
	defer qt.Recovering("callback QListWidget::itemChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "itemChanged"); signal != nil {
		signal.(func(*QListWidgetItem))(NewQListWidgetItemFromPointer(item))
	}

}

func (ptr *QListWidget) ItemChanged(item QListWidgetItem_ITF) {
	defer qt.Recovering("QListWidget::itemChanged")

	if ptr.Pointer() != nil {
		C.QListWidget_ItemChanged(ptr.Pointer(), PointerFromQListWidgetItem(item))
	}
}

func (ptr *QListWidget) ConnectItemClicked(f func(item *QListWidgetItem)) {
	defer qt.Recovering("connect QListWidget::itemClicked")

	if ptr.Pointer() != nil {
		C.QListWidget_ConnectItemClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "itemClicked", f)
	}
}

func (ptr *QListWidget) DisconnectItemClicked() {
	defer qt.Recovering("disconnect QListWidget::itemClicked")

	if ptr.Pointer() != nil {
		C.QListWidget_DisconnectItemClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "itemClicked")
	}
}

//export callbackQListWidgetItemClicked
func callbackQListWidgetItemClicked(ptr unsafe.Pointer, ptrName *C.char, item unsafe.Pointer) {
	defer qt.Recovering("callback QListWidget::itemClicked")

	if signal := qt.GetSignal(C.GoString(ptrName), "itemClicked"); signal != nil {
		signal.(func(*QListWidgetItem))(NewQListWidgetItemFromPointer(item))
	}

}

func (ptr *QListWidget) ItemClicked(item QListWidgetItem_ITF) {
	defer qt.Recovering("QListWidget::itemClicked")

	if ptr.Pointer() != nil {
		C.QListWidget_ItemClicked(ptr.Pointer(), PointerFromQListWidgetItem(item))
	}
}

func (ptr *QListWidget) ConnectItemDoubleClicked(f func(item *QListWidgetItem)) {
	defer qt.Recovering("connect QListWidget::itemDoubleClicked")

	if ptr.Pointer() != nil {
		C.QListWidget_ConnectItemDoubleClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "itemDoubleClicked", f)
	}
}

func (ptr *QListWidget) DisconnectItemDoubleClicked() {
	defer qt.Recovering("disconnect QListWidget::itemDoubleClicked")

	if ptr.Pointer() != nil {
		C.QListWidget_DisconnectItemDoubleClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "itemDoubleClicked")
	}
}

//export callbackQListWidgetItemDoubleClicked
func callbackQListWidgetItemDoubleClicked(ptr unsafe.Pointer, ptrName *C.char, item unsafe.Pointer) {
	defer qt.Recovering("callback QListWidget::itemDoubleClicked")

	if signal := qt.GetSignal(C.GoString(ptrName), "itemDoubleClicked"); signal != nil {
		signal.(func(*QListWidgetItem))(NewQListWidgetItemFromPointer(item))
	}

}

func (ptr *QListWidget) ItemDoubleClicked(item QListWidgetItem_ITF) {
	defer qt.Recovering("QListWidget::itemDoubleClicked")

	if ptr.Pointer() != nil {
		C.QListWidget_ItemDoubleClicked(ptr.Pointer(), PointerFromQListWidgetItem(item))
	}
}

func (ptr *QListWidget) ConnectItemEntered(f func(item *QListWidgetItem)) {
	defer qt.Recovering("connect QListWidget::itemEntered")

	if ptr.Pointer() != nil {
		C.QListWidget_ConnectItemEntered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "itemEntered", f)
	}
}

func (ptr *QListWidget) DisconnectItemEntered() {
	defer qt.Recovering("disconnect QListWidget::itemEntered")

	if ptr.Pointer() != nil {
		C.QListWidget_DisconnectItemEntered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "itemEntered")
	}
}

//export callbackQListWidgetItemEntered
func callbackQListWidgetItemEntered(ptr unsafe.Pointer, ptrName *C.char, item unsafe.Pointer) {
	defer qt.Recovering("callback QListWidget::itemEntered")

	if signal := qt.GetSignal(C.GoString(ptrName), "itemEntered"); signal != nil {
		signal.(func(*QListWidgetItem))(NewQListWidgetItemFromPointer(item))
	}

}

func (ptr *QListWidget) ItemEntered(item QListWidgetItem_ITF) {
	defer qt.Recovering("QListWidget::itemEntered")

	if ptr.Pointer() != nil {
		C.QListWidget_ItemEntered(ptr.Pointer(), PointerFromQListWidgetItem(item))
	}
}

func (ptr *QListWidget) ConnectItemPressed(f func(item *QListWidgetItem)) {
	defer qt.Recovering("connect QListWidget::itemPressed")

	if ptr.Pointer() != nil {
		C.QListWidget_ConnectItemPressed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "itemPressed", f)
	}
}

func (ptr *QListWidget) DisconnectItemPressed() {
	defer qt.Recovering("disconnect QListWidget::itemPressed")

	if ptr.Pointer() != nil {
		C.QListWidget_DisconnectItemPressed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "itemPressed")
	}
}

//export callbackQListWidgetItemPressed
func callbackQListWidgetItemPressed(ptr unsafe.Pointer, ptrName *C.char, item unsafe.Pointer) {
	defer qt.Recovering("callback QListWidget::itemPressed")

	if signal := qt.GetSignal(C.GoString(ptrName), "itemPressed"); signal != nil {
		signal.(func(*QListWidgetItem))(NewQListWidgetItemFromPointer(item))
	}

}

func (ptr *QListWidget) ItemPressed(item QListWidgetItem_ITF) {
	defer qt.Recovering("QListWidget::itemPressed")

	if ptr.Pointer() != nil {
		C.QListWidget_ItemPressed(ptr.Pointer(), PointerFromQListWidgetItem(item))
	}
}

func (ptr *QListWidget) ConnectItemSelectionChanged(f func()) {
	defer qt.Recovering("connect QListWidget::itemSelectionChanged")

	if ptr.Pointer() != nil {
		C.QListWidget_ConnectItemSelectionChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "itemSelectionChanged", f)
	}
}

func (ptr *QListWidget) DisconnectItemSelectionChanged() {
	defer qt.Recovering("disconnect QListWidget::itemSelectionChanged")

	if ptr.Pointer() != nil {
		C.QListWidget_DisconnectItemSelectionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "itemSelectionChanged")
	}
}

//export callbackQListWidgetItemSelectionChanged
func callbackQListWidgetItemSelectionChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QListWidget::itemSelectionChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "itemSelectionChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QListWidget) ItemSelectionChanged() {
	defer qt.Recovering("QListWidget::itemSelectionChanged")

	if ptr.Pointer() != nil {
		C.QListWidget_ItemSelectionChanged(ptr.Pointer())
	}
}

func (ptr *QListWidget) ItemWidget(item QListWidgetItem_ITF) *QWidget {
	defer qt.Recovering("QListWidget::itemWidget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QListWidget_ItemWidget(ptr.Pointer(), PointerFromQListWidgetItem(item)))
	}
	return nil
}

func (ptr *QListWidget) MimeTypes() []string {
	defer qt.Recovering("QListWidget::mimeTypes")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QListWidget_MimeTypes(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QListWidget) OpenPersistentEditor(item QListWidgetItem_ITF) {
	defer qt.Recovering("QListWidget::openPersistentEditor")

	if ptr.Pointer() != nil {
		C.QListWidget_OpenPersistentEditor(ptr.Pointer(), PointerFromQListWidgetItem(item))
	}
}

func (ptr *QListWidget) RemoveItemWidget(item QListWidgetItem_ITF) {
	defer qt.Recovering("QListWidget::removeItemWidget")

	if ptr.Pointer() != nil {
		C.QListWidget_RemoveItemWidget(ptr.Pointer(), PointerFromQListWidgetItem(item))
	}
}

func (ptr *QListWidget) Row(item QListWidgetItem_ITF) int {
	defer qt.Recovering("QListWidget::row")

	if ptr.Pointer() != nil {
		return int(C.QListWidget_Row(ptr.Pointer(), PointerFromQListWidgetItem(item)))
	}
	return 0
}

func (ptr *QListWidget) ScrollToItem(item QListWidgetItem_ITF, hint QAbstractItemView__ScrollHint) {
	defer qt.Recovering("QListWidget::scrollToItem")

	if ptr.Pointer() != nil {
		C.QListWidget_ScrollToItem(ptr.Pointer(), PointerFromQListWidgetItem(item), C.int(hint))
	}
}

func (ptr *QListWidget) SetCurrentItem(item QListWidgetItem_ITF) {
	defer qt.Recovering("QListWidget::setCurrentItem")

	if ptr.Pointer() != nil {
		C.QListWidget_SetCurrentItem(ptr.Pointer(), PointerFromQListWidgetItem(item))
	}
}

func (ptr *QListWidget) SetCurrentItem2(item QListWidgetItem_ITF, command core.QItemSelectionModel__SelectionFlag) {
	defer qt.Recovering("QListWidget::setCurrentItem")

	if ptr.Pointer() != nil {
		C.QListWidget_SetCurrentItem2(ptr.Pointer(), PointerFromQListWidgetItem(item), C.int(command))
	}
}

func (ptr *QListWidget) SetCurrentRow2(row int, command core.QItemSelectionModel__SelectionFlag) {
	defer qt.Recovering("QListWidget::setCurrentRow")

	if ptr.Pointer() != nil {
		C.QListWidget_SetCurrentRow2(ptr.Pointer(), C.int(row), C.int(command))
	}
}

func (ptr *QListWidget) SetItemWidget(item QListWidgetItem_ITF, widget QWidget_ITF) {
	defer qt.Recovering("QListWidget::setItemWidget")

	if ptr.Pointer() != nil {
		C.QListWidget_SetItemWidget(ptr.Pointer(), PointerFromQListWidgetItem(item), PointerFromQWidget(widget))
	}
}

func (ptr *QListWidget) SortItems(order core.Qt__SortOrder) {
	defer qt.Recovering("QListWidget::sortItems")

	if ptr.Pointer() != nil {
		C.QListWidget_SortItems(ptr.Pointer(), C.int(order))
	}
}

func (ptr *QListWidget) SupportedDropActions() core.Qt__DropAction {
	defer qt.Recovering("QListWidget::supportedDropActions")

	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QListWidget_SupportedDropActions(ptr.Pointer()))
	}
	return 0
}

func (ptr *QListWidget) TakeItem(row int) *QListWidgetItem {
	defer qt.Recovering("QListWidget::takeItem")

	if ptr.Pointer() != nil {
		return NewQListWidgetItemFromPointer(C.QListWidget_TakeItem(ptr.Pointer(), C.int(row)))
	}
	return nil
}

func (ptr *QListWidget) VisualItemRect(item QListWidgetItem_ITF) *core.QRect {
	defer qt.Recovering("QListWidget::visualItemRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QListWidget_VisualItemRect(ptr.Pointer(), PointerFromQListWidgetItem(item)))
	}
	return nil
}

func (ptr *QListWidget) DestroyQListWidget() {
	defer qt.Recovering("QListWidget::~QListWidget")

	if ptr.Pointer() != nil {
		C.QListWidget_DestroyQListWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QListWidget) ConnectCurrentChanged(f func(current *core.QModelIndex, previous *core.QModelIndex)) {
	defer qt.Recovering("connect QListWidget::currentChanged")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "currentChanged", f)
	}
}

func (ptr *QListWidget) DisconnectCurrentChanged() {
	defer qt.Recovering("disconnect QListWidget::currentChanged")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "currentChanged")
	}
}

//export callbackQListWidgetCurrentChanged
func callbackQListWidgetCurrentChanged(ptr unsafe.Pointer, ptrName *C.char, current unsafe.Pointer, previous unsafe.Pointer) bool {
	defer qt.Recovering("callback QListWidget::currentChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "currentChanged"); signal != nil {
		signal.(func(*core.QModelIndex, *core.QModelIndex))(core.NewQModelIndexFromPointer(current), core.NewQModelIndexFromPointer(previous))
		return true
	}
	return false

}

func (ptr *QListWidget) CurrentChanged(current core.QModelIndex_ITF, previous core.QModelIndex_ITF) {
	defer qt.Recovering("QListWidget::currentChanged")

	if ptr.Pointer() != nil {
		C.QListWidget_CurrentChanged(ptr.Pointer(), core.PointerFromQModelIndex(current), core.PointerFromQModelIndex(previous))
	}
}

func (ptr *QListWidget) CurrentChangedDefault(current core.QModelIndex_ITF, previous core.QModelIndex_ITF) {
	defer qt.Recovering("QListWidget::currentChanged")

	if ptr.Pointer() != nil {
		C.QListWidget_CurrentChangedDefault(ptr.Pointer(), core.PointerFromQModelIndex(current), core.PointerFromQModelIndex(previous))
	}
}

func (ptr *QListWidget) ConnectDragLeaveEvent(f func(e *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QListWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QListWidget) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QListWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQListWidgetDragLeaveEvent
func callbackQListWidgetDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QListWidget::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(e))
	} else {
		NewQListWidgetFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(e))
	}
}

func (ptr *QListWidget) DragLeaveEvent(e gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QListWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(e))
	}
}

func (ptr *QListWidget) DragLeaveEventDefault(e gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QListWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(e))
	}
}

func (ptr *QListWidget) ConnectDragMoveEvent(f func(e *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QListWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QListWidget) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QListWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQListWidgetDragMoveEvent
func callbackQListWidgetDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QListWidget::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(e))
	} else {
		NewQListWidgetFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(e))
	}
}

func (ptr *QListWidget) DragMoveEvent(e gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QListWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(e))
	}
}

func (ptr *QListWidget) DragMoveEventDefault(e gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QListWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(e))
	}
}

func (ptr *QListWidget) ConnectMouseMoveEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QListWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QListWidget) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QListWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQListWidgetMouseMoveEvent
func callbackQListWidgetMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QListWidget::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQListWidgetFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QListWidget) MouseMoveEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QListWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QListWidget) MouseMoveEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QListWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QListWidget) ConnectMouseReleaseEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QListWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QListWidget) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QListWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQListWidgetMouseReleaseEvent
func callbackQListWidgetMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QListWidget::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQListWidgetFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QListWidget) MouseReleaseEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QListWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QListWidget) MouseReleaseEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QListWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QListWidget) ConnectPaintEvent(f func(e *gui.QPaintEvent)) {
	defer qt.Recovering("connect QListWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QListWidget) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QListWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQListWidgetPaintEvent
func callbackQListWidgetPaintEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QListWidget::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(e))
	} else {
		NewQListWidgetFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(e))
	}
}

func (ptr *QListWidget) PaintEvent(e gui.QPaintEvent_ITF) {
	defer qt.Recovering("QListWidget::paintEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(e))
	}
}

func (ptr *QListWidget) PaintEventDefault(e gui.QPaintEvent_ITF) {
	defer qt.Recovering("QListWidget::paintEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(e))
	}
}

func (ptr *QListWidget) ConnectResizeEvent(f func(e *gui.QResizeEvent)) {
	defer qt.Recovering("connect QListWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QListWidget) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QListWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQListWidgetResizeEvent
func callbackQListWidgetResizeEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QListWidget::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(e))
	} else {
		NewQListWidgetFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(e))
	}
}

func (ptr *QListWidget) ResizeEvent(e gui.QResizeEvent_ITF) {
	defer qt.Recovering("QListWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(e))
	}
}

func (ptr *QListWidget) ResizeEventDefault(e gui.QResizeEvent_ITF) {
	defer qt.Recovering("QListWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(e))
	}
}

func (ptr *QListWidget) ConnectRowsAboutToBeRemoved(f func(parent *core.QModelIndex, start int, end int)) {
	defer qt.Recovering("connect QListWidget::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "rowsAboutToBeRemoved", f)
	}
}

func (ptr *QListWidget) DisconnectRowsAboutToBeRemoved() {
	defer qt.Recovering("disconnect QListWidget::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "rowsAboutToBeRemoved")
	}
}

//export callbackQListWidgetRowsAboutToBeRemoved
func callbackQListWidgetRowsAboutToBeRemoved(ptr unsafe.Pointer, ptrName *C.char, parent unsafe.Pointer, start C.int, end C.int) bool {
	defer qt.Recovering("callback QListWidget::rowsAboutToBeRemoved")

	if signal := qt.GetSignal(C.GoString(ptrName), "rowsAboutToBeRemoved"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(start), int(end))
		return true
	}
	return false

}

func (ptr *QListWidget) RowsAboutToBeRemoved(parent core.QModelIndex_ITF, start int, end int) {
	defer qt.Recovering("QListWidget::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {
		C.QListWidget_RowsAboutToBeRemoved(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(start), C.int(end))
	}
}

func (ptr *QListWidget) RowsAboutToBeRemovedDefault(parent core.QModelIndex_ITF, start int, end int) {
	defer qt.Recovering("QListWidget::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {
		C.QListWidget_RowsAboutToBeRemovedDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(start), C.int(end))
	}
}

func (ptr *QListWidget) ConnectRowsInserted(f func(parent *core.QModelIndex, start int, end int)) {
	defer qt.Recovering("connect QListWidget::rowsInserted")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "rowsInserted", f)
	}
}

func (ptr *QListWidget) DisconnectRowsInserted() {
	defer qt.Recovering("disconnect QListWidget::rowsInserted")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "rowsInserted")
	}
}

//export callbackQListWidgetRowsInserted
func callbackQListWidgetRowsInserted(ptr unsafe.Pointer, ptrName *C.char, parent unsafe.Pointer, start C.int, end C.int) bool {
	defer qt.Recovering("callback QListWidget::rowsInserted")

	if signal := qt.GetSignal(C.GoString(ptrName), "rowsInserted"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(start), int(end))
		return true
	}
	return false

}

func (ptr *QListWidget) RowsInserted(parent core.QModelIndex_ITF, start int, end int) {
	defer qt.Recovering("QListWidget::rowsInserted")

	if ptr.Pointer() != nil {
		C.QListWidget_RowsInserted(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(start), C.int(end))
	}
}

func (ptr *QListWidget) RowsInsertedDefault(parent core.QModelIndex_ITF, start int, end int) {
	defer qt.Recovering("QListWidget::rowsInserted")

	if ptr.Pointer() != nil {
		C.QListWidget_RowsInsertedDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(start), C.int(end))
	}
}

func (ptr *QListWidget) ConnectScrollTo(f func(index *core.QModelIndex, hint QAbstractItemView__ScrollHint)) {
	defer qt.Recovering("connect QListWidget::scrollTo")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "scrollTo", f)
	}
}

func (ptr *QListWidget) DisconnectScrollTo() {
	defer qt.Recovering("disconnect QListWidget::scrollTo")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "scrollTo")
	}
}

//export callbackQListWidgetScrollTo
func callbackQListWidgetScrollTo(ptr unsafe.Pointer, ptrName *C.char, index unsafe.Pointer, hint C.int) {
	defer qt.Recovering("callback QListWidget::scrollTo")

	if signal := qt.GetSignal(C.GoString(ptrName), "scrollTo"); signal != nil {
		signal.(func(*core.QModelIndex, QAbstractItemView__ScrollHint))(core.NewQModelIndexFromPointer(index), QAbstractItemView__ScrollHint(hint))
	}

}

func (ptr *QListWidget) ScrollTo(index core.QModelIndex_ITF, hint QAbstractItemView__ScrollHint) {
	defer qt.Recovering("QListWidget::scrollTo")

	if ptr.Pointer() != nil {
		C.QListWidget_ScrollTo(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(hint))
	}
}

func (ptr *QListWidget) ScrollToDefault(index core.QModelIndex_ITF, hint QAbstractItemView__ScrollHint) {
	defer qt.Recovering("QListWidget::scrollTo")

	if ptr.Pointer() != nil {
		C.QListWidget_ScrollToDefault(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(hint))
	}
}

func (ptr *QListWidget) ConnectSetSelection(f func(rect *core.QRect, command core.QItemSelectionModel__SelectionFlag)) {
	defer qt.Recovering("connect QListWidget::setSelection")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setSelection", f)
	}
}

func (ptr *QListWidget) DisconnectSetSelection() {
	defer qt.Recovering("disconnect QListWidget::setSelection")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setSelection")
	}
}

//export callbackQListWidgetSetSelection
func callbackQListWidgetSetSelection(ptr unsafe.Pointer, ptrName *C.char, rect unsafe.Pointer, command C.int) {
	defer qt.Recovering("callback QListWidget::setSelection")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSelection"); signal != nil {
		signal.(func(*core.QRect, core.QItemSelectionModel__SelectionFlag))(core.NewQRectFromPointer(rect), core.QItemSelectionModel__SelectionFlag(command))
	}

}

func (ptr *QListWidget) SetSelection(rect core.QRect_ITF, command core.QItemSelectionModel__SelectionFlag) {
	defer qt.Recovering("QListWidget::setSelection")

	if ptr.Pointer() != nil {
		C.QListWidget_SetSelection(ptr.Pointer(), core.PointerFromQRect(rect), C.int(command))
	}
}

func (ptr *QListWidget) SetSelectionDefault(rect core.QRect_ITF, command core.QItemSelectionModel__SelectionFlag) {
	defer qt.Recovering("QListWidget::setSelection")

	if ptr.Pointer() != nil {
		C.QListWidget_SetSelectionDefault(ptr.Pointer(), core.PointerFromQRect(rect), C.int(command))
	}
}

func (ptr *QListWidget) ConnectStartDrag(f func(supportedActions core.Qt__DropAction)) {
	defer qt.Recovering("connect QListWidget::startDrag")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "startDrag", f)
	}
}

func (ptr *QListWidget) DisconnectStartDrag() {
	defer qt.Recovering("disconnect QListWidget::startDrag")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "startDrag")
	}
}

//export callbackQListWidgetStartDrag
func callbackQListWidgetStartDrag(ptr unsafe.Pointer, ptrName *C.char, supportedActions C.int) {
	defer qt.Recovering("callback QListWidget::startDrag")

	if signal := qt.GetSignal(C.GoString(ptrName), "startDrag"); signal != nil {
		signal.(func(core.Qt__DropAction))(core.Qt__DropAction(supportedActions))
	} else {
		NewQListWidgetFromPointer(ptr).StartDragDefault(core.Qt__DropAction(supportedActions))
	}
}

func (ptr *QListWidget) StartDrag(supportedActions core.Qt__DropAction) {
	defer qt.Recovering("QListWidget::startDrag")

	if ptr.Pointer() != nil {
		C.QListWidget_StartDrag(ptr.Pointer(), C.int(supportedActions))
	}
}

func (ptr *QListWidget) StartDragDefault(supportedActions core.Qt__DropAction) {
	defer qt.Recovering("QListWidget::startDrag")

	if ptr.Pointer() != nil {
		C.QListWidget_StartDragDefault(ptr.Pointer(), C.int(supportedActions))
	}
}

func (ptr *QListWidget) ConnectTimerEvent(f func(e *core.QTimerEvent)) {
	defer qt.Recovering("connect QListWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QListWidget) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QListWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQListWidgetTimerEvent
func callbackQListWidgetTimerEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QListWidget::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(e))
	} else {
		NewQListWidgetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(e))
	}
}

func (ptr *QListWidget) TimerEvent(e core.QTimerEvent_ITF) {
	defer qt.Recovering("QListWidget::timerEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(e))
	}
}

func (ptr *QListWidget) TimerEventDefault(e core.QTimerEvent_ITF) {
	defer qt.Recovering("QListWidget::timerEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(e))
	}
}

func (ptr *QListWidget) ConnectUpdateGeometries(f func()) {
	defer qt.Recovering("connect QListWidget::updateGeometries")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updateGeometries", f)
	}
}

func (ptr *QListWidget) DisconnectUpdateGeometries() {
	defer qt.Recovering("disconnect QListWidget::updateGeometries")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updateGeometries")
	}
}

//export callbackQListWidgetUpdateGeometries
func callbackQListWidgetUpdateGeometries(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QListWidget::updateGeometries")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateGeometries"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QListWidget) UpdateGeometries() {
	defer qt.Recovering("QListWidget::updateGeometries")

	if ptr.Pointer() != nil {
		C.QListWidget_UpdateGeometries(ptr.Pointer())
	}
}

func (ptr *QListWidget) UpdateGeometriesDefault() {
	defer qt.Recovering("QListWidget::updateGeometries")

	if ptr.Pointer() != nil {
		C.QListWidget_UpdateGeometriesDefault(ptr.Pointer())
	}
}

func (ptr *QListWidget) ConnectCloseEditor(f func(editor *QWidget, hint QAbstractItemDelegate__EndEditHint)) {
	defer qt.Recovering("connect QListWidget::closeEditor")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEditor", f)
	}
}

func (ptr *QListWidget) DisconnectCloseEditor() {
	defer qt.Recovering("disconnect QListWidget::closeEditor")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEditor")
	}
}

//export callbackQListWidgetCloseEditor
func callbackQListWidgetCloseEditor(ptr unsafe.Pointer, ptrName *C.char, editor unsafe.Pointer, hint C.int) bool {
	defer qt.Recovering("callback QListWidget::closeEditor")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEditor"); signal != nil {
		signal.(func(*QWidget, QAbstractItemDelegate__EndEditHint))(NewQWidgetFromPointer(editor), QAbstractItemDelegate__EndEditHint(hint))
		return true
	}
	return false

}

func (ptr *QListWidget) CloseEditor(editor QWidget_ITF, hint QAbstractItemDelegate__EndEditHint) {
	defer qt.Recovering("QListWidget::closeEditor")

	if ptr.Pointer() != nil {
		C.QListWidget_CloseEditor(ptr.Pointer(), PointerFromQWidget(editor), C.int(hint))
	}
}

func (ptr *QListWidget) CloseEditorDefault(editor QWidget_ITF, hint QAbstractItemDelegate__EndEditHint) {
	defer qt.Recovering("QListWidget::closeEditor")

	if ptr.Pointer() != nil {
		C.QListWidget_CloseEditorDefault(ptr.Pointer(), PointerFromQWidget(editor), C.int(hint))
	}
}

func (ptr *QListWidget) ConnectCommitData(f func(editor *QWidget)) {
	defer qt.Recovering("connect QListWidget::commitData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "commitData", f)
	}
}

func (ptr *QListWidget) DisconnectCommitData() {
	defer qt.Recovering("disconnect QListWidget::commitData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "commitData")
	}
}

//export callbackQListWidgetCommitData
func callbackQListWidgetCommitData(ptr unsafe.Pointer, ptrName *C.char, editor unsafe.Pointer) bool {
	defer qt.Recovering("callback QListWidget::commitData")

	if signal := qt.GetSignal(C.GoString(ptrName), "commitData"); signal != nil {
		signal.(func(*QWidget))(NewQWidgetFromPointer(editor))
		return true
	}
	return false

}

func (ptr *QListWidget) CommitData(editor QWidget_ITF) {
	defer qt.Recovering("QListWidget::commitData")

	if ptr.Pointer() != nil {
		C.QListWidget_CommitData(ptr.Pointer(), PointerFromQWidget(editor))
	}
}

func (ptr *QListWidget) CommitDataDefault(editor QWidget_ITF) {
	defer qt.Recovering("QListWidget::commitData")

	if ptr.Pointer() != nil {
		C.QListWidget_CommitDataDefault(ptr.Pointer(), PointerFromQWidget(editor))
	}
}

func (ptr *QListWidget) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QListWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QListWidget) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QListWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQListWidgetDragEnterEvent
func callbackQListWidgetDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QListWidget::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQListWidgetFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QListWidget) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QListWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QListWidget) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QListWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QListWidget) ConnectEditorDestroyed(f func(editor *core.QObject)) {
	defer qt.Recovering("connect QListWidget::editorDestroyed")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "editorDestroyed", f)
	}
}

func (ptr *QListWidget) DisconnectEditorDestroyed() {
	defer qt.Recovering("disconnect QListWidget::editorDestroyed")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "editorDestroyed")
	}
}

//export callbackQListWidgetEditorDestroyed
func callbackQListWidgetEditorDestroyed(ptr unsafe.Pointer, ptrName *C.char, editor unsafe.Pointer) bool {
	defer qt.Recovering("callback QListWidget::editorDestroyed")

	if signal := qt.GetSignal(C.GoString(ptrName), "editorDestroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(editor))
		return true
	}
	return false

}

func (ptr *QListWidget) EditorDestroyed(editor core.QObject_ITF) {
	defer qt.Recovering("QListWidget::editorDestroyed")

	if ptr.Pointer() != nil {
		C.QListWidget_EditorDestroyed(ptr.Pointer(), core.PointerFromQObject(editor))
	}
}

func (ptr *QListWidget) EditorDestroyedDefault(editor core.QObject_ITF) {
	defer qt.Recovering("QListWidget::editorDestroyed")

	if ptr.Pointer() != nil {
		C.QListWidget_EditorDestroyedDefault(ptr.Pointer(), core.PointerFromQObject(editor))
	}
}

func (ptr *QListWidget) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QListWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QListWidget) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QListWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQListWidgetFocusInEvent
func callbackQListWidgetFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QListWidget::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQListWidgetFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QListWidget) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QListWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QListWidget) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QListWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QListWidget) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QListWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QListWidget) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QListWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQListWidgetFocusOutEvent
func callbackQListWidgetFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QListWidget::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQListWidgetFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QListWidget) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QListWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QListWidget) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QListWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QListWidget) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QListWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QListWidget) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QListWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQListWidgetInputMethodEvent
func callbackQListWidgetInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QListWidget::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQListWidgetFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QListWidget) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QListWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QListWidget) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QListWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QListWidget) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QListWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QListWidget) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QListWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQListWidgetKeyPressEvent
func callbackQListWidgetKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QListWidget::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQListWidgetFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QListWidget) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QListWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QListWidget) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QListWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QListWidget) ConnectKeyboardSearch(f func(search string)) {
	defer qt.Recovering("connect QListWidget::keyboardSearch")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyboardSearch", f)
	}
}

func (ptr *QListWidget) DisconnectKeyboardSearch() {
	defer qt.Recovering("disconnect QListWidget::keyboardSearch")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyboardSearch")
	}
}

//export callbackQListWidgetKeyboardSearch
func callbackQListWidgetKeyboardSearch(ptr unsafe.Pointer, ptrName *C.char, search *C.char) {
	defer qt.Recovering("callback QListWidget::keyboardSearch")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyboardSearch"); signal != nil {
		signal.(func(string))(C.GoString(search))
	} else {
		NewQListWidgetFromPointer(ptr).KeyboardSearchDefault(C.GoString(search))
	}
}

func (ptr *QListWidget) KeyboardSearch(search string) {
	defer qt.Recovering("QListWidget::keyboardSearch")

	if ptr.Pointer() != nil {
		C.QListWidget_KeyboardSearch(ptr.Pointer(), C.CString(search))
	}
}

func (ptr *QListWidget) KeyboardSearchDefault(search string) {
	defer qt.Recovering("QListWidget::keyboardSearch")

	if ptr.Pointer() != nil {
		C.QListWidget_KeyboardSearchDefault(ptr.Pointer(), C.CString(search))
	}
}

func (ptr *QListWidget) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QListWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QListWidget) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QListWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQListWidgetMouseDoubleClickEvent
func callbackQListWidgetMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QListWidget::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQListWidgetFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QListWidget) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QListWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QListWidget) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QListWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QListWidget) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QListWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QListWidget) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QListWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQListWidgetMousePressEvent
func callbackQListWidgetMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QListWidget::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQListWidgetFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QListWidget) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QListWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QListWidget) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QListWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QListWidget) ConnectReset(f func()) {
	defer qt.Recovering("connect QListWidget::reset")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "reset", f)
	}
}

func (ptr *QListWidget) DisconnectReset() {
	defer qt.Recovering("disconnect QListWidget::reset")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "reset")
	}
}

//export callbackQListWidgetReset
func callbackQListWidgetReset(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QListWidget::reset")

	if signal := qt.GetSignal(C.GoString(ptrName), "reset"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QListWidget) Reset() {
	defer qt.Recovering("QListWidget::reset")

	if ptr.Pointer() != nil {
		C.QListWidget_Reset(ptr.Pointer())
	}
}

func (ptr *QListWidget) ResetDefault() {
	defer qt.Recovering("QListWidget::reset")

	if ptr.Pointer() != nil {
		C.QListWidget_ResetDefault(ptr.Pointer())
	}
}

func (ptr *QListWidget) ConnectSelectAll(f func()) {
	defer qt.Recovering("connect QListWidget::selectAll")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "selectAll", f)
	}
}

func (ptr *QListWidget) DisconnectSelectAll() {
	defer qt.Recovering("disconnect QListWidget::selectAll")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "selectAll")
	}
}

//export callbackQListWidgetSelectAll
func callbackQListWidgetSelectAll(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QListWidget::selectAll")

	if signal := qt.GetSignal(C.GoString(ptrName), "selectAll"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QListWidget) SelectAll() {
	defer qt.Recovering("QListWidget::selectAll")

	if ptr.Pointer() != nil {
		C.QListWidget_SelectAll(ptr.Pointer())
	}
}

func (ptr *QListWidget) SelectAllDefault() {
	defer qt.Recovering("QListWidget::selectAll")

	if ptr.Pointer() != nil {
		C.QListWidget_SelectAllDefault(ptr.Pointer())
	}
}

func (ptr *QListWidget) ConnectSetRootIndex(f func(index *core.QModelIndex)) {
	defer qt.Recovering("connect QListWidget::setRootIndex")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setRootIndex", f)
	}
}

func (ptr *QListWidget) DisconnectSetRootIndex() {
	defer qt.Recovering("disconnect QListWidget::setRootIndex")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setRootIndex")
	}
}

//export callbackQListWidgetSetRootIndex
func callbackQListWidgetSetRootIndex(ptr unsafe.Pointer, ptrName *C.char, index unsafe.Pointer) bool {
	defer qt.Recovering("callback QListWidget::setRootIndex")

	if signal := qt.GetSignal(C.GoString(ptrName), "setRootIndex"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
		return true
	}
	return false

}

func (ptr *QListWidget) SetRootIndex(index core.QModelIndex_ITF) {
	defer qt.Recovering("QListWidget::setRootIndex")

	if ptr.Pointer() != nil {
		C.QListWidget_SetRootIndex(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QListWidget) SetRootIndexDefault(index core.QModelIndex_ITF) {
	defer qt.Recovering("QListWidget::setRootIndex")

	if ptr.Pointer() != nil {
		C.QListWidget_SetRootIndexDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QListWidget) ConnectSetSelectionModel(f func(selectionModel *core.QItemSelectionModel)) {
	defer qt.Recovering("connect QListWidget::setSelectionModel")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setSelectionModel", f)
	}
}

func (ptr *QListWidget) DisconnectSetSelectionModel() {
	defer qt.Recovering("disconnect QListWidget::setSelectionModel")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setSelectionModel")
	}
}

//export callbackQListWidgetSetSelectionModel
func callbackQListWidgetSetSelectionModel(ptr unsafe.Pointer, ptrName *C.char, selectionModel unsafe.Pointer) {
	defer qt.Recovering("callback QListWidget::setSelectionModel")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSelectionModel"); signal != nil {
		signal.(func(*core.QItemSelectionModel))(core.NewQItemSelectionModelFromPointer(selectionModel))
	} else {
		NewQListWidgetFromPointer(ptr).SetSelectionModelDefault(core.NewQItemSelectionModelFromPointer(selectionModel))
	}
}

func (ptr *QListWidget) SetSelectionModel(selectionModel core.QItemSelectionModel_ITF) {
	defer qt.Recovering("QListWidget::setSelectionModel")

	if ptr.Pointer() != nil {
		C.QListWidget_SetSelectionModel(ptr.Pointer(), core.PointerFromQItemSelectionModel(selectionModel))
	}
}

func (ptr *QListWidget) SetSelectionModelDefault(selectionModel core.QItemSelectionModel_ITF) {
	defer qt.Recovering("QListWidget::setSelectionModel")

	if ptr.Pointer() != nil {
		C.QListWidget_SetSelectionModelDefault(ptr.Pointer(), core.PointerFromQItemSelectionModel(selectionModel))
	}
}

func (ptr *QListWidget) ConnectContextMenuEvent(f func(e *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QListWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QListWidget) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QListWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQListWidgetContextMenuEvent
func callbackQListWidgetContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QListWidget::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
	} else {
		NewQListWidgetFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(e))
	}
}

func (ptr *QListWidget) ContextMenuEvent(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QListWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

func (ptr *QListWidget) ContextMenuEventDefault(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QListWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

func (ptr *QListWidget) ConnectScrollContentsBy(f func(dx int, dy int)) {
	defer qt.Recovering("connect QListWidget::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "scrollContentsBy", f)
	}
}

func (ptr *QListWidget) DisconnectScrollContentsBy() {
	defer qt.Recovering("disconnect QListWidget::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "scrollContentsBy")
	}
}

//export callbackQListWidgetScrollContentsBy
func callbackQListWidgetScrollContentsBy(ptr unsafe.Pointer, ptrName *C.char, dx C.int, dy C.int) {
	defer qt.Recovering("callback QListWidget::scrollContentsBy")

	if signal := qt.GetSignal(C.GoString(ptrName), "scrollContentsBy"); signal != nil {
		signal.(func(int, int))(int(dx), int(dy))
	} else {
		NewQListWidgetFromPointer(ptr).ScrollContentsByDefault(int(dx), int(dy))
	}
}

func (ptr *QListWidget) ScrollContentsBy(dx int, dy int) {
	defer qt.Recovering("QListWidget::scrollContentsBy")

	if ptr.Pointer() != nil {
		C.QListWidget_ScrollContentsBy(ptr.Pointer(), C.int(dx), C.int(dy))
	}
}

func (ptr *QListWidget) ScrollContentsByDefault(dx int, dy int) {
	defer qt.Recovering("QListWidget::scrollContentsBy")

	if ptr.Pointer() != nil {
		C.QListWidget_ScrollContentsByDefault(ptr.Pointer(), C.int(dx), C.int(dy))
	}
}

func (ptr *QListWidget) ConnectSetupViewport(f func(viewport *QWidget)) {
	defer qt.Recovering("connect QListWidget::setupViewport")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setupViewport", f)
	}
}

func (ptr *QListWidget) DisconnectSetupViewport() {
	defer qt.Recovering("disconnect QListWidget::setupViewport")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setupViewport")
	}
}

//export callbackQListWidgetSetupViewport
func callbackQListWidgetSetupViewport(ptr unsafe.Pointer, ptrName *C.char, viewport unsafe.Pointer) {
	defer qt.Recovering("callback QListWidget::setupViewport")

	if signal := qt.GetSignal(C.GoString(ptrName), "setupViewport"); signal != nil {
		signal.(func(*QWidget))(NewQWidgetFromPointer(viewport))
	} else {
		NewQListWidgetFromPointer(ptr).SetupViewportDefault(NewQWidgetFromPointer(viewport))
	}
}

func (ptr *QListWidget) SetupViewport(viewport QWidget_ITF) {
	defer qt.Recovering("QListWidget::setupViewport")

	if ptr.Pointer() != nil {
		C.QListWidget_SetupViewport(ptr.Pointer(), PointerFromQWidget(viewport))
	}
}

func (ptr *QListWidget) SetupViewportDefault(viewport QWidget_ITF) {
	defer qt.Recovering("QListWidget::setupViewport")

	if ptr.Pointer() != nil {
		C.QListWidget_SetupViewportDefault(ptr.Pointer(), PointerFromQWidget(viewport))
	}
}

func (ptr *QListWidget) ConnectWheelEvent(f func(e *gui.QWheelEvent)) {
	defer qt.Recovering("connect QListWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QListWidget) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QListWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQListWidgetWheelEvent
func callbackQListWidgetWheelEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QListWidget::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
	} else {
		NewQListWidgetFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(e))
	}
}

func (ptr *QListWidget) WheelEvent(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QListWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QListWidget) WheelEventDefault(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QListWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QListWidget) ConnectChangeEvent(f func(ev *core.QEvent)) {
	defer qt.Recovering("connect QListWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QListWidget) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QListWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQListWidgetChangeEvent
func callbackQListWidgetChangeEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QListWidget::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
	} else {
		NewQListWidgetFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(ev))
	}
}

func (ptr *QListWidget) ChangeEvent(ev core.QEvent_ITF) {
	defer qt.Recovering("QListWidget::changeEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QListWidget) ChangeEventDefault(ev core.QEvent_ITF) {
	defer qt.Recovering("QListWidget::changeEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QListWidget) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QListWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QListWidget) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QListWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQListWidgetActionEvent
func callbackQListWidgetActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QListWidget::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQListWidgetFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QListWidget) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QListWidget::actionEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QListWidget) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QListWidget::actionEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QListWidget) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QListWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QListWidget) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QListWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQListWidgetEnterEvent
func callbackQListWidgetEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QListWidget::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQListWidgetFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QListWidget) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QListWidget::enterEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QListWidget) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QListWidget::enterEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QListWidget) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QListWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QListWidget) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QListWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQListWidgetHideEvent
func callbackQListWidgetHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QListWidget::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQListWidgetFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QListWidget) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QListWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QListWidget) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QListWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QListWidget) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QListWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QListWidget) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QListWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQListWidgetLeaveEvent
func callbackQListWidgetLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QListWidget::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQListWidgetFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QListWidget) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QListWidget::leaveEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QListWidget) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QListWidget::leaveEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QListWidget) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QListWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QListWidget) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QListWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQListWidgetMoveEvent
func callbackQListWidgetMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QListWidget::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQListWidgetFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QListWidget) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QListWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QListWidget) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QListWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QListWidget) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QListWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QListWidget) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QListWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQListWidgetSetVisible
func callbackQListWidgetSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QListWidget::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QListWidget) SetVisible(visible bool) {
	defer qt.Recovering("QListWidget::setVisible")

	if ptr.Pointer() != nil {
		C.QListWidget_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QListWidget) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QListWidget::setVisible")

	if ptr.Pointer() != nil {
		C.QListWidget_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QListWidget) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QListWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QListWidget) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QListWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQListWidgetShowEvent
func callbackQListWidgetShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QListWidget::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQListWidgetFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QListWidget) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QListWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QListWidget) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QListWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QListWidget) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QListWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QListWidget) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QListWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQListWidgetCloseEvent
func callbackQListWidgetCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QListWidget::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQListWidgetFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QListWidget) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QListWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QListWidget) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QListWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QListWidget) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QListWidget::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QListWidget) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QListWidget::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQListWidgetInitPainter
func callbackQListWidgetInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QListWidget::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQListWidgetFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QListWidget) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QListWidget::initPainter")

	if ptr.Pointer() != nil {
		C.QListWidget_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QListWidget) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QListWidget::initPainter")

	if ptr.Pointer() != nil {
		C.QListWidget_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QListWidget) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QListWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QListWidget) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QListWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQListWidgetKeyReleaseEvent
func callbackQListWidgetKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QListWidget::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQListWidgetFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QListWidget) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QListWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QListWidget) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QListWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QListWidget) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QListWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QListWidget) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QListWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQListWidgetTabletEvent
func callbackQListWidgetTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QListWidget::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQListWidgetFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QListWidget) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QListWidget::tabletEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QListWidget) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QListWidget::tabletEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QListWidget) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QListWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QListWidget) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QListWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQListWidgetChildEvent
func callbackQListWidgetChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QListWidget::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQListWidgetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QListWidget) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QListWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QListWidget) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QListWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QListWidget) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QListWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QListWidget) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QListWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQListWidgetCustomEvent
func callbackQListWidgetCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QListWidget::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQListWidgetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QListWidget) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QListWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QListWidget) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QListWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QListWidget_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
