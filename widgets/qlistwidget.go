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

type QListWidgetITF interface {
	QListViewITF
	QListWidgetPTR() *QListWidget
}

func PointerFromQListWidget(ptr QListWidgetITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QListWidgetPTR().Pointer()
	}
	return nil
}

func QListWidgetFromPointer(ptr unsafe.Pointer) *QListWidget {
	var n = new(QListWidget)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QListWidget_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QListWidget) QListWidgetPTR() *QListWidget {
	return ptr
}

func (ptr *QListWidget) Count() int {
	if ptr.Pointer() != nil {
		return int(C.QListWidget_Count(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QListWidget) CurrentRow() int {
	if ptr.Pointer() != nil {
		return int(C.QListWidget_CurrentRow(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QListWidget) IsSortingEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QListWidget_IsSortingEnabled(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QListWidget) SetCurrentRow(row int) {
	if ptr.Pointer() != nil {
		C.QListWidget_SetCurrentRow(C.QtObjectPtr(ptr.Pointer()), C.int(row))
	}
}

func (ptr *QListWidget) SetSortingEnabled(enable bool) {
	if ptr.Pointer() != nil {
		C.QListWidget_SetSortingEnabled(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func NewQListWidget(parent QWidgetITF) *QListWidget {
	return QListWidgetFromPointer(unsafe.Pointer(C.QListWidget_NewQListWidget(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QListWidget) AddItem2(item QListWidgetItemITF) {
	if ptr.Pointer() != nil {
		C.QListWidget_AddItem2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQListWidgetItem(item)))
	}
}

func (ptr *QListWidget) AddItem(label string) {
	if ptr.Pointer() != nil {
		C.QListWidget_AddItem(C.QtObjectPtr(ptr.Pointer()), C.CString(label))
	}
}

func (ptr *QListWidget) AddItems(labels []string) {
	if ptr.Pointer() != nil {
		C.QListWidget_AddItems(C.QtObjectPtr(ptr.Pointer()), C.CString(strings.Join(labels, "|")))
	}
}

func (ptr *QListWidget) Clear() {
	if ptr.Pointer() != nil {
		C.QListWidget_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QListWidget) ClosePersistentEditor(item QListWidgetItemITF) {
	if ptr.Pointer() != nil {
		C.QListWidget_ClosePersistentEditor(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQListWidgetItem(item)))
	}
}

func (ptr *QListWidget) CurrentItem() *QListWidgetItem {
	if ptr.Pointer() != nil {
		return QListWidgetItemFromPointer(unsafe.Pointer(C.QListWidget_CurrentItem(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QListWidget) ConnectCurrentItemChanged(f func(current QListWidgetItemITF, previous QListWidgetItemITF)) {
	if ptr.Pointer() != nil {
		C.QListWidget_ConnectCurrentItemChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "currentItemChanged", f)
	}
}

func (ptr *QListWidget) DisconnectCurrentItemChanged() {
	if ptr.Pointer() != nil {
		C.QListWidget_DisconnectCurrentItemChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "currentItemChanged")
	}
}

//export callbackQListWidgetCurrentItemChanged
func callbackQListWidgetCurrentItemChanged(ptrName *C.char, current unsafe.Pointer, previous unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "currentItemChanged").(func(*QListWidgetItem, *QListWidgetItem))(QListWidgetItemFromPointer(current), QListWidgetItemFromPointer(previous))
}

func (ptr *QListWidget) ConnectCurrentRowChanged(f func(currentRow int)) {
	if ptr.Pointer() != nil {
		C.QListWidget_ConnectCurrentRowChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "currentRowChanged", f)
	}
}

func (ptr *QListWidget) DisconnectCurrentRowChanged() {
	if ptr.Pointer() != nil {
		C.QListWidget_DisconnectCurrentRowChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "currentRowChanged")
	}
}

//export callbackQListWidgetCurrentRowChanged
func callbackQListWidgetCurrentRowChanged(ptrName *C.char, currentRow C.int) {
	qt.GetSignal(C.GoString(ptrName), "currentRowChanged").(func(int))(int(currentRow))
}

func (ptr *QListWidget) ConnectCurrentTextChanged(f func(currentText string)) {
	if ptr.Pointer() != nil {
		C.QListWidget_ConnectCurrentTextChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "currentTextChanged", f)
	}
}

func (ptr *QListWidget) DisconnectCurrentTextChanged() {
	if ptr.Pointer() != nil {
		C.QListWidget_DisconnectCurrentTextChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "currentTextChanged")
	}
}

//export callbackQListWidgetCurrentTextChanged
func callbackQListWidgetCurrentTextChanged(ptrName *C.char, currentText *C.char) {
	qt.GetSignal(C.GoString(ptrName), "currentTextChanged").(func(string))(C.GoString(currentText))
}

func (ptr *QListWidget) DropEvent(event gui.QDropEventITF) {
	if ptr.Pointer() != nil {
		C.QListWidget_DropEvent(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQDropEvent(event)))
	}
}

func (ptr *QListWidget) EditItem(item QListWidgetItemITF) {
	if ptr.Pointer() != nil {
		C.QListWidget_EditItem(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQListWidgetItem(item)))
	}
}

func (ptr *QListWidget) InsertItem(row int, item QListWidgetItemITF) {
	if ptr.Pointer() != nil {
		C.QListWidget_InsertItem(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.QtObjectPtr(PointerFromQListWidgetItem(item)))
	}
}

func (ptr *QListWidget) InsertItem2(row int, label string) {
	if ptr.Pointer() != nil {
		C.QListWidget_InsertItem2(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.CString(label))
	}
}

func (ptr *QListWidget) InsertItems(row int, labels []string) {
	if ptr.Pointer() != nil {
		C.QListWidget_InsertItems(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.CString(strings.Join(labels, "|")))
	}
}

func (ptr *QListWidget) Item(row int) *QListWidgetItem {
	if ptr.Pointer() != nil {
		return QListWidgetItemFromPointer(unsafe.Pointer(C.QListWidget_Item(C.QtObjectPtr(ptr.Pointer()), C.int(row))))
	}
	return nil
}

func (ptr *QListWidget) ConnectItemActivated(f func(item QListWidgetItemITF)) {
	if ptr.Pointer() != nil {
		C.QListWidget_ConnectItemActivated(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "itemActivated", f)
	}
}

func (ptr *QListWidget) DisconnectItemActivated() {
	if ptr.Pointer() != nil {
		C.QListWidget_DisconnectItemActivated(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "itemActivated")
	}
}

//export callbackQListWidgetItemActivated
func callbackQListWidgetItemActivated(ptrName *C.char, item unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "itemActivated").(func(*QListWidgetItem))(QListWidgetItemFromPointer(item))
}

func (ptr *QListWidget) ItemAt(p core.QPointITF) *QListWidgetItem {
	if ptr.Pointer() != nil {
		return QListWidgetItemFromPointer(unsafe.Pointer(C.QListWidget_ItemAt(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(p)))))
	}
	return nil
}

func (ptr *QListWidget) ItemAt2(x int, y int) *QListWidgetItem {
	if ptr.Pointer() != nil {
		return QListWidgetItemFromPointer(unsafe.Pointer(C.QListWidget_ItemAt2(C.QtObjectPtr(ptr.Pointer()), C.int(x), C.int(y))))
	}
	return nil
}

func (ptr *QListWidget) ConnectItemChanged(f func(item QListWidgetItemITF)) {
	if ptr.Pointer() != nil {
		C.QListWidget_ConnectItemChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "itemChanged", f)
	}
}

func (ptr *QListWidget) DisconnectItemChanged() {
	if ptr.Pointer() != nil {
		C.QListWidget_DisconnectItemChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "itemChanged")
	}
}

//export callbackQListWidgetItemChanged
func callbackQListWidgetItemChanged(ptrName *C.char, item unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "itemChanged").(func(*QListWidgetItem))(QListWidgetItemFromPointer(item))
}

func (ptr *QListWidget) ConnectItemClicked(f func(item QListWidgetItemITF)) {
	if ptr.Pointer() != nil {
		C.QListWidget_ConnectItemClicked(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "itemClicked", f)
	}
}

func (ptr *QListWidget) DisconnectItemClicked() {
	if ptr.Pointer() != nil {
		C.QListWidget_DisconnectItemClicked(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "itemClicked")
	}
}

//export callbackQListWidgetItemClicked
func callbackQListWidgetItemClicked(ptrName *C.char, item unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "itemClicked").(func(*QListWidgetItem))(QListWidgetItemFromPointer(item))
}

func (ptr *QListWidget) ConnectItemDoubleClicked(f func(item QListWidgetItemITF)) {
	if ptr.Pointer() != nil {
		C.QListWidget_ConnectItemDoubleClicked(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "itemDoubleClicked", f)
	}
}

func (ptr *QListWidget) DisconnectItemDoubleClicked() {
	if ptr.Pointer() != nil {
		C.QListWidget_DisconnectItemDoubleClicked(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "itemDoubleClicked")
	}
}

//export callbackQListWidgetItemDoubleClicked
func callbackQListWidgetItemDoubleClicked(ptrName *C.char, item unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "itemDoubleClicked").(func(*QListWidgetItem))(QListWidgetItemFromPointer(item))
}

func (ptr *QListWidget) ConnectItemEntered(f func(item QListWidgetItemITF)) {
	if ptr.Pointer() != nil {
		C.QListWidget_ConnectItemEntered(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "itemEntered", f)
	}
}

func (ptr *QListWidget) DisconnectItemEntered() {
	if ptr.Pointer() != nil {
		C.QListWidget_DisconnectItemEntered(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "itemEntered")
	}
}

//export callbackQListWidgetItemEntered
func callbackQListWidgetItemEntered(ptrName *C.char, item unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "itemEntered").(func(*QListWidgetItem))(QListWidgetItemFromPointer(item))
}

func (ptr *QListWidget) ConnectItemPressed(f func(item QListWidgetItemITF)) {
	if ptr.Pointer() != nil {
		C.QListWidget_ConnectItemPressed(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "itemPressed", f)
	}
}

func (ptr *QListWidget) DisconnectItemPressed() {
	if ptr.Pointer() != nil {
		C.QListWidget_DisconnectItemPressed(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "itemPressed")
	}
}

//export callbackQListWidgetItemPressed
func callbackQListWidgetItemPressed(ptrName *C.char, item unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "itemPressed").(func(*QListWidgetItem))(QListWidgetItemFromPointer(item))
}

func (ptr *QListWidget) ConnectItemSelectionChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QListWidget_ConnectItemSelectionChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "itemSelectionChanged", f)
	}
}

func (ptr *QListWidget) DisconnectItemSelectionChanged() {
	if ptr.Pointer() != nil {
		C.QListWidget_DisconnectItemSelectionChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "itemSelectionChanged")
	}
}

//export callbackQListWidgetItemSelectionChanged
func callbackQListWidgetItemSelectionChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "itemSelectionChanged").(func())()
}

func (ptr *QListWidget) ItemWidget(item QListWidgetItemITF) *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QListWidget_ItemWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQListWidgetItem(item)))))
	}
	return nil
}

func (ptr *QListWidget) OpenPersistentEditor(item QListWidgetItemITF) {
	if ptr.Pointer() != nil {
		C.QListWidget_OpenPersistentEditor(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQListWidgetItem(item)))
	}
}

func (ptr *QListWidget) RemoveItemWidget(item QListWidgetItemITF) {
	if ptr.Pointer() != nil {
		C.QListWidget_RemoveItemWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQListWidgetItem(item)))
	}
}

func (ptr *QListWidget) Row(item QListWidgetItemITF) int {
	if ptr.Pointer() != nil {
		return int(C.QListWidget_Row(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQListWidgetItem(item))))
	}
	return 0
}

func (ptr *QListWidget) ScrollToItem(item QListWidgetItemITF, hint QAbstractItemView__ScrollHint) {
	if ptr.Pointer() != nil {
		C.QListWidget_ScrollToItem(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQListWidgetItem(item)), C.int(hint))
	}
}

func (ptr *QListWidget) SetCurrentItem(item QListWidgetItemITF) {
	if ptr.Pointer() != nil {
		C.QListWidget_SetCurrentItem(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQListWidgetItem(item)))
	}
}

func (ptr *QListWidget) SetCurrentItem2(item QListWidgetItemITF, command core.QItemSelectionModel__SelectionFlag) {
	if ptr.Pointer() != nil {
		C.QListWidget_SetCurrentItem2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQListWidgetItem(item)), C.int(command))
	}
}

func (ptr *QListWidget) SetCurrentRow2(row int, command core.QItemSelectionModel__SelectionFlag) {
	if ptr.Pointer() != nil {
		C.QListWidget_SetCurrentRow2(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(command))
	}
}

func (ptr *QListWidget) SetItemWidget(item QListWidgetItemITF, widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QListWidget_SetItemWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQListWidgetItem(item)), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QListWidget) SortItems(order core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.QListWidget_SortItems(C.QtObjectPtr(ptr.Pointer()), C.int(order))
	}
}

func (ptr *QListWidget) TakeItem(row int) *QListWidgetItem {
	if ptr.Pointer() != nil {
		return QListWidgetItemFromPointer(unsafe.Pointer(C.QListWidget_TakeItem(C.QtObjectPtr(ptr.Pointer()), C.int(row))))
	}
	return nil
}

func (ptr *QListWidget) DestroyQListWidget() {
	if ptr.Pointer() != nil {
		C.QListWidget_DestroyQListWidget(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
