package widgets

//#include "qtreewidget.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QTreeWidget struct {
	QTreeView
}

type QTreeWidgetITF interface {
	QTreeViewITF
	QTreeWidgetPTR() *QTreeWidget
}

func PointerFromQTreeWidget(ptr QTreeWidgetITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTreeWidgetPTR().Pointer()
	}
	return nil
}

func QTreeWidgetFromPointer(ptr unsafe.Pointer) *QTreeWidget {
	var n = new(QTreeWidget)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QTreeWidget_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTreeWidget) QTreeWidgetPTR() *QTreeWidget {
	return ptr
}

func (ptr *QTreeWidget) ColumnCount() int {
	if ptr.Pointer() != nil {
		return int(C.QTreeWidget_ColumnCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTreeWidget) SetColumnCount(columns int) {
	if ptr.Pointer() != nil {
		C.QTreeWidget_SetColumnCount(C.QtObjectPtr(ptr.Pointer()), C.int(columns))
	}
}

func (ptr *QTreeWidget) TopLevelItemCount() int {
	if ptr.Pointer() != nil {
		return int(C.QTreeWidget_TopLevelItemCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func NewQTreeWidget(parent QWidgetITF) *QTreeWidget {
	return QTreeWidgetFromPointer(unsafe.Pointer(C.QTreeWidget_NewQTreeWidget(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QTreeWidget) AddTopLevelItem(item QTreeWidgetItemITF) {
	if ptr.Pointer() != nil {
		C.QTreeWidget_AddTopLevelItem(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTreeWidgetItem(item)))
	}
}

func (ptr *QTreeWidget) Clear() {
	if ptr.Pointer() != nil {
		C.QTreeWidget_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTreeWidget) ClosePersistentEditor(item QTreeWidgetItemITF, column int) {
	if ptr.Pointer() != nil {
		C.QTreeWidget_ClosePersistentEditor(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTreeWidgetItem(item)), C.int(column))
	}
}

func (ptr *QTreeWidget) CollapseItem(item QTreeWidgetItemITF) {
	if ptr.Pointer() != nil {
		C.QTreeWidget_CollapseItem(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTreeWidgetItem(item)))
	}
}

func (ptr *QTreeWidget) CurrentColumn() int {
	if ptr.Pointer() != nil {
		return int(C.QTreeWidget_CurrentColumn(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTreeWidget) CurrentItem() *QTreeWidgetItem {
	if ptr.Pointer() != nil {
		return QTreeWidgetItemFromPointer(unsafe.Pointer(C.QTreeWidget_CurrentItem(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QTreeWidget) ConnectCurrentItemChanged(f func(current QTreeWidgetItemITF, previous QTreeWidgetItemITF)) {
	if ptr.Pointer() != nil {
		C.QTreeWidget_ConnectCurrentItemChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "currentItemChanged", f)
	}
}

func (ptr *QTreeWidget) DisconnectCurrentItemChanged() {
	if ptr.Pointer() != nil {
		C.QTreeWidget_DisconnectCurrentItemChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "currentItemChanged")
	}
}

//export callbackQTreeWidgetCurrentItemChanged
func callbackQTreeWidgetCurrentItemChanged(ptrName *C.char, current unsafe.Pointer, previous unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "currentItemChanged").(func(*QTreeWidgetItem, *QTreeWidgetItem))(QTreeWidgetItemFromPointer(current), QTreeWidgetItemFromPointer(previous))
}

func (ptr *QTreeWidget) EditItem(item QTreeWidgetItemITF, column int) {
	if ptr.Pointer() != nil {
		C.QTreeWidget_EditItem(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTreeWidgetItem(item)), C.int(column))
	}
}

func (ptr *QTreeWidget) ExpandItem(item QTreeWidgetItemITF) {
	if ptr.Pointer() != nil {
		C.QTreeWidget_ExpandItem(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTreeWidgetItem(item)))
	}
}

func (ptr *QTreeWidget) HeaderItem() *QTreeWidgetItem {
	if ptr.Pointer() != nil {
		return QTreeWidgetItemFromPointer(unsafe.Pointer(C.QTreeWidget_HeaderItem(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QTreeWidget) IndexOfTopLevelItem(item QTreeWidgetItemITF) int {
	if ptr.Pointer() != nil {
		return int(C.QTreeWidget_IndexOfTopLevelItem(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTreeWidgetItem(item))))
	}
	return 0
}

func (ptr *QTreeWidget) InsertTopLevelItem(index int, item QTreeWidgetItemITF) {
	if ptr.Pointer() != nil {
		C.QTreeWidget_InsertTopLevelItem(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.QtObjectPtr(PointerFromQTreeWidgetItem(item)))
	}
}

func (ptr *QTreeWidget) InvisibleRootItem() *QTreeWidgetItem {
	if ptr.Pointer() != nil {
		return QTreeWidgetItemFromPointer(unsafe.Pointer(C.QTreeWidget_InvisibleRootItem(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QTreeWidget) IsFirstItemColumnSpanned(item QTreeWidgetItemITF) bool {
	if ptr.Pointer() != nil {
		return C.QTreeWidget_IsFirstItemColumnSpanned(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTreeWidgetItem(item))) != 0
	}
	return false
}

func (ptr *QTreeWidget) ItemAbove(item QTreeWidgetItemITF) *QTreeWidgetItem {
	if ptr.Pointer() != nil {
		return QTreeWidgetItemFromPointer(unsafe.Pointer(C.QTreeWidget_ItemAbove(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTreeWidgetItem(item)))))
	}
	return nil
}

func (ptr *QTreeWidget) ConnectItemActivated(f func(item QTreeWidgetItemITF, column int)) {
	if ptr.Pointer() != nil {
		C.QTreeWidget_ConnectItemActivated(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "itemActivated", f)
	}
}

func (ptr *QTreeWidget) DisconnectItemActivated() {
	if ptr.Pointer() != nil {
		C.QTreeWidget_DisconnectItemActivated(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "itemActivated")
	}
}

//export callbackQTreeWidgetItemActivated
func callbackQTreeWidgetItemActivated(ptrName *C.char, item unsafe.Pointer, column C.int) {
	qt.GetSignal(C.GoString(ptrName), "itemActivated").(func(*QTreeWidgetItem, int))(QTreeWidgetItemFromPointer(item), int(column))
}

func (ptr *QTreeWidget) ItemAt(p core.QPointITF) *QTreeWidgetItem {
	if ptr.Pointer() != nil {
		return QTreeWidgetItemFromPointer(unsafe.Pointer(C.QTreeWidget_ItemAt(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(p)))))
	}
	return nil
}

func (ptr *QTreeWidget) ItemAt2(x int, y int) *QTreeWidgetItem {
	if ptr.Pointer() != nil {
		return QTreeWidgetItemFromPointer(unsafe.Pointer(C.QTreeWidget_ItemAt2(C.QtObjectPtr(ptr.Pointer()), C.int(x), C.int(y))))
	}
	return nil
}

func (ptr *QTreeWidget) ItemBelow(item QTreeWidgetItemITF) *QTreeWidgetItem {
	if ptr.Pointer() != nil {
		return QTreeWidgetItemFromPointer(unsafe.Pointer(C.QTreeWidget_ItemBelow(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTreeWidgetItem(item)))))
	}
	return nil
}

func (ptr *QTreeWidget) ConnectItemChanged(f func(item QTreeWidgetItemITF, column int)) {
	if ptr.Pointer() != nil {
		C.QTreeWidget_ConnectItemChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "itemChanged", f)
	}
}

func (ptr *QTreeWidget) DisconnectItemChanged() {
	if ptr.Pointer() != nil {
		C.QTreeWidget_DisconnectItemChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "itemChanged")
	}
}

//export callbackQTreeWidgetItemChanged
func callbackQTreeWidgetItemChanged(ptrName *C.char, item unsafe.Pointer, column C.int) {
	qt.GetSignal(C.GoString(ptrName), "itemChanged").(func(*QTreeWidgetItem, int))(QTreeWidgetItemFromPointer(item), int(column))
}

func (ptr *QTreeWidget) ConnectItemClicked(f func(item QTreeWidgetItemITF, column int)) {
	if ptr.Pointer() != nil {
		C.QTreeWidget_ConnectItemClicked(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "itemClicked", f)
	}
}

func (ptr *QTreeWidget) DisconnectItemClicked() {
	if ptr.Pointer() != nil {
		C.QTreeWidget_DisconnectItemClicked(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "itemClicked")
	}
}

//export callbackQTreeWidgetItemClicked
func callbackQTreeWidgetItemClicked(ptrName *C.char, item unsafe.Pointer, column C.int) {
	qt.GetSignal(C.GoString(ptrName), "itemClicked").(func(*QTreeWidgetItem, int))(QTreeWidgetItemFromPointer(item), int(column))
}

func (ptr *QTreeWidget) ConnectItemCollapsed(f func(item QTreeWidgetItemITF)) {
	if ptr.Pointer() != nil {
		C.QTreeWidget_ConnectItemCollapsed(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "itemCollapsed", f)
	}
}

func (ptr *QTreeWidget) DisconnectItemCollapsed() {
	if ptr.Pointer() != nil {
		C.QTreeWidget_DisconnectItemCollapsed(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "itemCollapsed")
	}
}

//export callbackQTreeWidgetItemCollapsed
func callbackQTreeWidgetItemCollapsed(ptrName *C.char, item unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "itemCollapsed").(func(*QTreeWidgetItem))(QTreeWidgetItemFromPointer(item))
}

func (ptr *QTreeWidget) ConnectItemDoubleClicked(f func(item QTreeWidgetItemITF, column int)) {
	if ptr.Pointer() != nil {
		C.QTreeWidget_ConnectItemDoubleClicked(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "itemDoubleClicked", f)
	}
}

func (ptr *QTreeWidget) DisconnectItemDoubleClicked() {
	if ptr.Pointer() != nil {
		C.QTreeWidget_DisconnectItemDoubleClicked(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "itemDoubleClicked")
	}
}

//export callbackQTreeWidgetItemDoubleClicked
func callbackQTreeWidgetItemDoubleClicked(ptrName *C.char, item unsafe.Pointer, column C.int) {
	qt.GetSignal(C.GoString(ptrName), "itemDoubleClicked").(func(*QTreeWidgetItem, int))(QTreeWidgetItemFromPointer(item), int(column))
}

func (ptr *QTreeWidget) ConnectItemEntered(f func(item QTreeWidgetItemITF, column int)) {
	if ptr.Pointer() != nil {
		C.QTreeWidget_ConnectItemEntered(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "itemEntered", f)
	}
}

func (ptr *QTreeWidget) DisconnectItemEntered() {
	if ptr.Pointer() != nil {
		C.QTreeWidget_DisconnectItemEntered(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "itemEntered")
	}
}

//export callbackQTreeWidgetItemEntered
func callbackQTreeWidgetItemEntered(ptrName *C.char, item unsafe.Pointer, column C.int) {
	qt.GetSignal(C.GoString(ptrName), "itemEntered").(func(*QTreeWidgetItem, int))(QTreeWidgetItemFromPointer(item), int(column))
}

func (ptr *QTreeWidget) ConnectItemExpanded(f func(item QTreeWidgetItemITF)) {
	if ptr.Pointer() != nil {
		C.QTreeWidget_ConnectItemExpanded(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "itemExpanded", f)
	}
}

func (ptr *QTreeWidget) DisconnectItemExpanded() {
	if ptr.Pointer() != nil {
		C.QTreeWidget_DisconnectItemExpanded(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "itemExpanded")
	}
}

//export callbackQTreeWidgetItemExpanded
func callbackQTreeWidgetItemExpanded(ptrName *C.char, item unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "itemExpanded").(func(*QTreeWidgetItem))(QTreeWidgetItemFromPointer(item))
}

func (ptr *QTreeWidget) ConnectItemPressed(f func(item QTreeWidgetItemITF, column int)) {
	if ptr.Pointer() != nil {
		C.QTreeWidget_ConnectItemPressed(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "itemPressed", f)
	}
}

func (ptr *QTreeWidget) DisconnectItemPressed() {
	if ptr.Pointer() != nil {
		C.QTreeWidget_DisconnectItemPressed(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "itemPressed")
	}
}

//export callbackQTreeWidgetItemPressed
func callbackQTreeWidgetItemPressed(ptrName *C.char, item unsafe.Pointer, column C.int) {
	qt.GetSignal(C.GoString(ptrName), "itemPressed").(func(*QTreeWidgetItem, int))(QTreeWidgetItemFromPointer(item), int(column))
}

func (ptr *QTreeWidget) ConnectItemSelectionChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QTreeWidget_ConnectItemSelectionChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "itemSelectionChanged", f)
	}
}

func (ptr *QTreeWidget) DisconnectItemSelectionChanged() {
	if ptr.Pointer() != nil {
		C.QTreeWidget_DisconnectItemSelectionChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "itemSelectionChanged")
	}
}

//export callbackQTreeWidgetItemSelectionChanged
func callbackQTreeWidgetItemSelectionChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "itemSelectionChanged").(func())()
}

func (ptr *QTreeWidget) ItemWidget(item QTreeWidgetItemITF, column int) *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QTreeWidget_ItemWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTreeWidgetItem(item)), C.int(column))))
	}
	return nil
}

func (ptr *QTreeWidget) OpenPersistentEditor(item QTreeWidgetItemITF, column int) {
	if ptr.Pointer() != nil {
		C.QTreeWidget_OpenPersistentEditor(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTreeWidgetItem(item)), C.int(column))
	}
}

func (ptr *QTreeWidget) RemoveItemWidget(item QTreeWidgetItemITF, column int) {
	if ptr.Pointer() != nil {
		C.QTreeWidget_RemoveItemWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTreeWidgetItem(item)), C.int(column))
	}
}

func (ptr *QTreeWidget) ScrollToItem(item QTreeWidgetItemITF, hint QAbstractItemView__ScrollHint) {
	if ptr.Pointer() != nil {
		C.QTreeWidget_ScrollToItem(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTreeWidgetItem(item)), C.int(hint))
	}
}

func (ptr *QTreeWidget) SetCurrentItem(item QTreeWidgetItemITF) {
	if ptr.Pointer() != nil {
		C.QTreeWidget_SetCurrentItem(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTreeWidgetItem(item)))
	}
}

func (ptr *QTreeWidget) SetCurrentItem2(item QTreeWidgetItemITF, column int) {
	if ptr.Pointer() != nil {
		C.QTreeWidget_SetCurrentItem2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTreeWidgetItem(item)), C.int(column))
	}
}

func (ptr *QTreeWidget) SetCurrentItem3(item QTreeWidgetItemITF, column int, command core.QItemSelectionModel__SelectionFlag) {
	if ptr.Pointer() != nil {
		C.QTreeWidget_SetCurrentItem3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTreeWidgetItem(item)), C.int(column), C.int(command))
	}
}

func (ptr *QTreeWidget) SetFirstItemColumnSpanned(item QTreeWidgetItemITF, span bool) {
	if ptr.Pointer() != nil {
		C.QTreeWidget_SetFirstItemColumnSpanned(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTreeWidgetItem(item)), C.int(qt.GoBoolToInt(span)))
	}
}

func (ptr *QTreeWidget) SetHeaderItem(item QTreeWidgetItemITF) {
	if ptr.Pointer() != nil {
		C.QTreeWidget_SetHeaderItem(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTreeWidgetItem(item)))
	}
}

func (ptr *QTreeWidget) SetHeaderLabel(label string) {
	if ptr.Pointer() != nil {
		C.QTreeWidget_SetHeaderLabel(C.QtObjectPtr(ptr.Pointer()), C.CString(label))
	}
}

func (ptr *QTreeWidget) SetHeaderLabels(labels []string) {
	if ptr.Pointer() != nil {
		C.QTreeWidget_SetHeaderLabels(C.QtObjectPtr(ptr.Pointer()), C.CString(strings.Join(labels, "|")))
	}
}

func (ptr *QTreeWidget) SetItemWidget(item QTreeWidgetItemITF, column int, widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QTreeWidget_SetItemWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTreeWidgetItem(item)), C.int(column), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QTreeWidget) SetSelectionModel(selectionModel core.QItemSelectionModelITF) {
	if ptr.Pointer() != nil {
		C.QTreeWidget_SetSelectionModel(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQItemSelectionModel(selectionModel)))
	}
}

func (ptr *QTreeWidget) SortColumn() int {
	if ptr.Pointer() != nil {
		return int(C.QTreeWidget_SortColumn(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTreeWidget) SortItems(column int, order core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.QTreeWidget_SortItems(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.int(order))
	}
}

func (ptr *QTreeWidget) TakeTopLevelItem(index int) *QTreeWidgetItem {
	if ptr.Pointer() != nil {
		return QTreeWidgetItemFromPointer(unsafe.Pointer(C.QTreeWidget_TakeTopLevelItem(C.QtObjectPtr(ptr.Pointer()), C.int(index))))
	}
	return nil
}

func (ptr *QTreeWidget) TopLevelItem(index int) *QTreeWidgetItem {
	if ptr.Pointer() != nil {
		return QTreeWidgetItemFromPointer(unsafe.Pointer(C.QTreeWidget_TopLevelItem(C.QtObjectPtr(ptr.Pointer()), C.int(index))))
	}
	return nil
}

func (ptr *QTreeWidget) DestroyQTreeWidget() {
	if ptr.Pointer() != nil {
		C.QTreeWidget_DestroyQTreeWidget(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
