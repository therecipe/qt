package widgets

//#include "qlistwidget.h"
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
	if n.ObjectName() == "" {
		n.SetObjectName("QListWidget_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QListWidget) QListWidget_PTR() *QListWidget {
	return ptr
}

func (ptr *QListWidget) Count() int {
	if ptr.Pointer() != nil {
		return int(C.QListWidget_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QListWidget) CurrentRow() int {
	if ptr.Pointer() != nil {
		return int(C.QListWidget_CurrentRow(ptr.Pointer()))
	}
	return 0
}

func (ptr *QListWidget) IsSortingEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QListWidget_IsSortingEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QListWidget) SetCurrentRow(row int) {
	if ptr.Pointer() != nil {
		C.QListWidget_SetCurrentRow(ptr.Pointer(), C.int(row))
	}
}

func (ptr *QListWidget) SetSortingEnabled(enable bool) {
	if ptr.Pointer() != nil {
		C.QListWidget_SetSortingEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func NewQListWidget(parent QWidget_ITF) *QListWidget {
	return NewQListWidgetFromPointer(C.QListWidget_NewQListWidget(PointerFromQWidget(parent)))
}

func (ptr *QListWidget) AddItem2(item QListWidgetItem_ITF) {
	if ptr.Pointer() != nil {
		C.QListWidget_AddItem2(ptr.Pointer(), PointerFromQListWidgetItem(item))
	}
}

func (ptr *QListWidget) AddItem(label string) {
	if ptr.Pointer() != nil {
		C.QListWidget_AddItem(ptr.Pointer(), C.CString(label))
	}
}

func (ptr *QListWidget) AddItems(labels []string) {
	if ptr.Pointer() != nil {
		C.QListWidget_AddItems(ptr.Pointer(), C.CString(strings.Join(labels, "|")))
	}
}

func (ptr *QListWidget) Clear() {
	if ptr.Pointer() != nil {
		C.QListWidget_Clear(ptr.Pointer())
	}
}

func (ptr *QListWidget) ClosePersistentEditor(item QListWidgetItem_ITF) {
	if ptr.Pointer() != nil {
		C.QListWidget_ClosePersistentEditor(ptr.Pointer(), PointerFromQListWidgetItem(item))
	}
}

func (ptr *QListWidget) CurrentItem() *QListWidgetItem {
	if ptr.Pointer() != nil {
		return NewQListWidgetItemFromPointer(C.QListWidget_CurrentItem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QListWidget) ConnectCurrentItemChanged(f func(current *QListWidgetItem, previous *QListWidgetItem)) {
	if ptr.Pointer() != nil {
		C.QListWidget_ConnectCurrentItemChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentItemChanged", f)
	}
}

func (ptr *QListWidget) DisconnectCurrentItemChanged() {
	if ptr.Pointer() != nil {
		C.QListWidget_DisconnectCurrentItemChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentItemChanged")
	}
}

//export callbackQListWidgetCurrentItemChanged
func callbackQListWidgetCurrentItemChanged(ptrName *C.char, current unsafe.Pointer, previous unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "currentItemChanged").(func(*QListWidgetItem, *QListWidgetItem))(NewQListWidgetItemFromPointer(current), NewQListWidgetItemFromPointer(previous))
}

func (ptr *QListWidget) ConnectCurrentRowChanged(f func(currentRow int)) {
	if ptr.Pointer() != nil {
		C.QListWidget_ConnectCurrentRowChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentRowChanged", f)
	}
}

func (ptr *QListWidget) DisconnectCurrentRowChanged() {
	if ptr.Pointer() != nil {
		C.QListWidget_DisconnectCurrentRowChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentRowChanged")
	}
}

//export callbackQListWidgetCurrentRowChanged
func callbackQListWidgetCurrentRowChanged(ptrName *C.char, currentRow C.int) {
	qt.GetSignal(C.GoString(ptrName), "currentRowChanged").(func(int))(int(currentRow))
}

func (ptr *QListWidget) ConnectCurrentTextChanged(f func(currentText string)) {
	if ptr.Pointer() != nil {
		C.QListWidget_ConnectCurrentTextChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentTextChanged", f)
	}
}

func (ptr *QListWidget) DisconnectCurrentTextChanged() {
	if ptr.Pointer() != nil {
		C.QListWidget_DisconnectCurrentTextChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentTextChanged")
	}
}

//export callbackQListWidgetCurrentTextChanged
func callbackQListWidgetCurrentTextChanged(ptrName *C.char, currentText *C.char) {
	qt.GetSignal(C.GoString(ptrName), "currentTextChanged").(func(string))(C.GoString(currentText))
}

func (ptr *QListWidget) DropEvent(event gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QListWidget_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QListWidget) EditItem(item QListWidgetItem_ITF) {
	if ptr.Pointer() != nil {
		C.QListWidget_EditItem(ptr.Pointer(), PointerFromQListWidgetItem(item))
	}
}

func (ptr *QListWidget) InsertItem(row int, item QListWidgetItem_ITF) {
	if ptr.Pointer() != nil {
		C.QListWidget_InsertItem(ptr.Pointer(), C.int(row), PointerFromQListWidgetItem(item))
	}
}

func (ptr *QListWidget) InsertItem2(row int, label string) {
	if ptr.Pointer() != nil {
		C.QListWidget_InsertItem2(ptr.Pointer(), C.int(row), C.CString(label))
	}
}

func (ptr *QListWidget) InsertItems(row int, labels []string) {
	if ptr.Pointer() != nil {
		C.QListWidget_InsertItems(ptr.Pointer(), C.int(row), C.CString(strings.Join(labels, "|")))
	}
}

func (ptr *QListWidget) Item(row int) *QListWidgetItem {
	if ptr.Pointer() != nil {
		return NewQListWidgetItemFromPointer(C.QListWidget_Item(ptr.Pointer(), C.int(row)))
	}
	return nil
}

func (ptr *QListWidget) ConnectItemActivated(f func(item *QListWidgetItem)) {
	if ptr.Pointer() != nil {
		C.QListWidget_ConnectItemActivated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "itemActivated", f)
	}
}

func (ptr *QListWidget) DisconnectItemActivated() {
	if ptr.Pointer() != nil {
		C.QListWidget_DisconnectItemActivated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "itemActivated")
	}
}

//export callbackQListWidgetItemActivated
func callbackQListWidgetItemActivated(ptrName *C.char, item unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "itemActivated").(func(*QListWidgetItem))(NewQListWidgetItemFromPointer(item))
}

func (ptr *QListWidget) ItemAt(p core.QPoint_ITF) *QListWidgetItem {
	if ptr.Pointer() != nil {
		return NewQListWidgetItemFromPointer(C.QListWidget_ItemAt(ptr.Pointer(), core.PointerFromQPoint(p)))
	}
	return nil
}

func (ptr *QListWidget) ItemAt2(x int, y int) *QListWidgetItem {
	if ptr.Pointer() != nil {
		return NewQListWidgetItemFromPointer(C.QListWidget_ItemAt2(ptr.Pointer(), C.int(x), C.int(y)))
	}
	return nil
}

func (ptr *QListWidget) ConnectItemChanged(f func(item *QListWidgetItem)) {
	if ptr.Pointer() != nil {
		C.QListWidget_ConnectItemChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "itemChanged", f)
	}
}

func (ptr *QListWidget) DisconnectItemChanged() {
	if ptr.Pointer() != nil {
		C.QListWidget_DisconnectItemChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "itemChanged")
	}
}

//export callbackQListWidgetItemChanged
func callbackQListWidgetItemChanged(ptrName *C.char, item unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "itemChanged").(func(*QListWidgetItem))(NewQListWidgetItemFromPointer(item))
}

func (ptr *QListWidget) ConnectItemClicked(f func(item *QListWidgetItem)) {
	if ptr.Pointer() != nil {
		C.QListWidget_ConnectItemClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "itemClicked", f)
	}
}

func (ptr *QListWidget) DisconnectItemClicked() {
	if ptr.Pointer() != nil {
		C.QListWidget_DisconnectItemClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "itemClicked")
	}
}

//export callbackQListWidgetItemClicked
func callbackQListWidgetItemClicked(ptrName *C.char, item unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "itemClicked").(func(*QListWidgetItem))(NewQListWidgetItemFromPointer(item))
}

func (ptr *QListWidget) ConnectItemDoubleClicked(f func(item *QListWidgetItem)) {
	if ptr.Pointer() != nil {
		C.QListWidget_ConnectItemDoubleClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "itemDoubleClicked", f)
	}
}

func (ptr *QListWidget) DisconnectItemDoubleClicked() {
	if ptr.Pointer() != nil {
		C.QListWidget_DisconnectItemDoubleClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "itemDoubleClicked")
	}
}

//export callbackQListWidgetItemDoubleClicked
func callbackQListWidgetItemDoubleClicked(ptrName *C.char, item unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "itemDoubleClicked").(func(*QListWidgetItem))(NewQListWidgetItemFromPointer(item))
}

func (ptr *QListWidget) ConnectItemEntered(f func(item *QListWidgetItem)) {
	if ptr.Pointer() != nil {
		C.QListWidget_ConnectItemEntered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "itemEntered", f)
	}
}

func (ptr *QListWidget) DisconnectItemEntered() {
	if ptr.Pointer() != nil {
		C.QListWidget_DisconnectItemEntered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "itemEntered")
	}
}

//export callbackQListWidgetItemEntered
func callbackQListWidgetItemEntered(ptrName *C.char, item unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "itemEntered").(func(*QListWidgetItem))(NewQListWidgetItemFromPointer(item))
}

func (ptr *QListWidget) ConnectItemPressed(f func(item *QListWidgetItem)) {
	if ptr.Pointer() != nil {
		C.QListWidget_ConnectItemPressed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "itemPressed", f)
	}
}

func (ptr *QListWidget) DisconnectItemPressed() {
	if ptr.Pointer() != nil {
		C.QListWidget_DisconnectItemPressed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "itemPressed")
	}
}

//export callbackQListWidgetItemPressed
func callbackQListWidgetItemPressed(ptrName *C.char, item unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "itemPressed").(func(*QListWidgetItem))(NewQListWidgetItemFromPointer(item))
}

func (ptr *QListWidget) ConnectItemSelectionChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QListWidget_ConnectItemSelectionChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "itemSelectionChanged", f)
	}
}

func (ptr *QListWidget) DisconnectItemSelectionChanged() {
	if ptr.Pointer() != nil {
		C.QListWidget_DisconnectItemSelectionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "itemSelectionChanged")
	}
}

//export callbackQListWidgetItemSelectionChanged
func callbackQListWidgetItemSelectionChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "itemSelectionChanged").(func())()
}

func (ptr *QListWidget) ItemWidget(item QListWidgetItem_ITF) *QWidget {
	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QListWidget_ItemWidget(ptr.Pointer(), PointerFromQListWidgetItem(item)))
	}
	return nil
}

func (ptr *QListWidget) OpenPersistentEditor(item QListWidgetItem_ITF) {
	if ptr.Pointer() != nil {
		C.QListWidget_OpenPersistentEditor(ptr.Pointer(), PointerFromQListWidgetItem(item))
	}
}

func (ptr *QListWidget) RemoveItemWidget(item QListWidgetItem_ITF) {
	if ptr.Pointer() != nil {
		C.QListWidget_RemoveItemWidget(ptr.Pointer(), PointerFromQListWidgetItem(item))
	}
}

func (ptr *QListWidget) Row(item QListWidgetItem_ITF) int {
	if ptr.Pointer() != nil {
		return int(C.QListWidget_Row(ptr.Pointer(), PointerFromQListWidgetItem(item)))
	}
	return 0
}

func (ptr *QListWidget) ScrollToItem(item QListWidgetItem_ITF, hint QAbstractItemView__ScrollHint) {
	if ptr.Pointer() != nil {
		C.QListWidget_ScrollToItem(ptr.Pointer(), PointerFromQListWidgetItem(item), C.int(hint))
	}
}

func (ptr *QListWidget) SetCurrentItem(item QListWidgetItem_ITF) {
	if ptr.Pointer() != nil {
		C.QListWidget_SetCurrentItem(ptr.Pointer(), PointerFromQListWidgetItem(item))
	}
}

func (ptr *QListWidget) SetCurrentItem2(item QListWidgetItem_ITF, command core.QItemSelectionModel__SelectionFlag) {
	if ptr.Pointer() != nil {
		C.QListWidget_SetCurrentItem2(ptr.Pointer(), PointerFromQListWidgetItem(item), C.int(command))
	}
}

func (ptr *QListWidget) SetCurrentRow2(row int, command core.QItemSelectionModel__SelectionFlag) {
	if ptr.Pointer() != nil {
		C.QListWidget_SetCurrentRow2(ptr.Pointer(), C.int(row), C.int(command))
	}
}

func (ptr *QListWidget) SetItemWidget(item QListWidgetItem_ITF, widget QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QListWidget_SetItemWidget(ptr.Pointer(), PointerFromQListWidgetItem(item), PointerFromQWidget(widget))
	}
}

func (ptr *QListWidget) SortItems(order core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.QListWidget_SortItems(ptr.Pointer(), C.int(order))
	}
}

func (ptr *QListWidget) TakeItem(row int) *QListWidgetItem {
	if ptr.Pointer() != nil {
		return NewQListWidgetItemFromPointer(C.QListWidget_TakeItem(ptr.Pointer(), C.int(row)))
	}
	return nil
}

func (ptr *QListWidget) DestroyQListWidget() {
	if ptr.Pointer() != nil {
		C.QListWidget_DestroyQListWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
