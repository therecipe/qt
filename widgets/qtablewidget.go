package widgets

//#include "qtablewidget.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QTableWidget struct {
	QTableView
}

type QTableWidgetITF interface {
	QTableViewITF
	QTableWidgetPTR() *QTableWidget
}

func PointerFromQTableWidget(ptr QTableWidgetITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTableWidgetPTR().Pointer()
	}
	return nil
}

func QTableWidgetFromPointer(ptr unsafe.Pointer) *QTableWidget {
	var n = new(QTableWidget)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QTableWidget_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTableWidget) QTableWidgetPTR() *QTableWidget {
	return ptr
}

func (ptr *QTableWidget) ItemAt(point core.QPointITF) *QTableWidgetItem {
	if ptr.Pointer() != nil {
		return QTableWidgetItemFromPointer(unsafe.Pointer(C.QTableWidget_ItemAt(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(point)))))
	}
	return nil
}

func (ptr *QTableWidget) ConnectCellActivated(f func(row int, column int)) {
	if ptr.Pointer() != nil {
		C.QTableWidget_ConnectCellActivated(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "cellActivated", f)
	}
}

func (ptr *QTableWidget) DisconnectCellActivated() {
	if ptr.Pointer() != nil {
		C.QTableWidget_DisconnectCellActivated(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "cellActivated")
	}
}

//export callbackQTableWidgetCellActivated
func callbackQTableWidgetCellActivated(ptrName *C.char, row C.int, column C.int) {
	qt.GetSignal(C.GoString(ptrName), "cellActivated").(func(int, int))(int(row), int(column))
}

func (ptr *QTableWidget) ConnectCellChanged(f func(row int, column int)) {
	if ptr.Pointer() != nil {
		C.QTableWidget_ConnectCellChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "cellChanged", f)
	}
}

func (ptr *QTableWidget) DisconnectCellChanged() {
	if ptr.Pointer() != nil {
		C.QTableWidget_DisconnectCellChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "cellChanged")
	}
}

//export callbackQTableWidgetCellChanged
func callbackQTableWidgetCellChanged(ptrName *C.char, row C.int, column C.int) {
	qt.GetSignal(C.GoString(ptrName), "cellChanged").(func(int, int))(int(row), int(column))
}

func (ptr *QTableWidget) ConnectCellClicked(f func(row int, column int)) {
	if ptr.Pointer() != nil {
		C.QTableWidget_ConnectCellClicked(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "cellClicked", f)
	}
}

func (ptr *QTableWidget) DisconnectCellClicked() {
	if ptr.Pointer() != nil {
		C.QTableWidget_DisconnectCellClicked(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "cellClicked")
	}
}

//export callbackQTableWidgetCellClicked
func callbackQTableWidgetCellClicked(ptrName *C.char, row C.int, column C.int) {
	qt.GetSignal(C.GoString(ptrName), "cellClicked").(func(int, int))(int(row), int(column))
}

func (ptr *QTableWidget) ConnectCellDoubleClicked(f func(row int, column int)) {
	if ptr.Pointer() != nil {
		C.QTableWidget_ConnectCellDoubleClicked(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "cellDoubleClicked", f)
	}
}

func (ptr *QTableWidget) DisconnectCellDoubleClicked() {
	if ptr.Pointer() != nil {
		C.QTableWidget_DisconnectCellDoubleClicked(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "cellDoubleClicked")
	}
}

//export callbackQTableWidgetCellDoubleClicked
func callbackQTableWidgetCellDoubleClicked(ptrName *C.char, row C.int, column C.int) {
	qt.GetSignal(C.GoString(ptrName), "cellDoubleClicked").(func(int, int))(int(row), int(column))
}

func (ptr *QTableWidget) ConnectCellEntered(f func(row int, column int)) {
	if ptr.Pointer() != nil {
		C.QTableWidget_ConnectCellEntered(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "cellEntered", f)
	}
}

func (ptr *QTableWidget) DisconnectCellEntered() {
	if ptr.Pointer() != nil {
		C.QTableWidget_DisconnectCellEntered(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "cellEntered")
	}
}

//export callbackQTableWidgetCellEntered
func callbackQTableWidgetCellEntered(ptrName *C.char, row C.int, column C.int) {
	qt.GetSignal(C.GoString(ptrName), "cellEntered").(func(int, int))(int(row), int(column))
}

func (ptr *QTableWidget) ConnectCellPressed(f func(row int, column int)) {
	if ptr.Pointer() != nil {
		C.QTableWidget_ConnectCellPressed(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "cellPressed", f)
	}
}

func (ptr *QTableWidget) DisconnectCellPressed() {
	if ptr.Pointer() != nil {
		C.QTableWidget_DisconnectCellPressed(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "cellPressed")
	}
}

//export callbackQTableWidgetCellPressed
func callbackQTableWidgetCellPressed(ptrName *C.char, row C.int, column C.int) {
	qt.GetSignal(C.GoString(ptrName), "cellPressed").(func(int, int))(int(row), int(column))
}

func (ptr *QTableWidget) CellWidget(row int, column int) *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QTableWidget_CellWidget(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(column))))
	}
	return nil
}

func (ptr *QTableWidget) Clear() {
	if ptr.Pointer() != nil {
		C.QTableWidget_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTableWidget) ClearContents() {
	if ptr.Pointer() != nil {
		C.QTableWidget_ClearContents(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTableWidget) ClosePersistentEditor(item QTableWidgetItemITF) {
	if ptr.Pointer() != nil {
		C.QTableWidget_ClosePersistentEditor(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTableWidgetItem(item)))
	}
}

func (ptr *QTableWidget) Column(item QTableWidgetItemITF) int {
	if ptr.Pointer() != nil {
		return int(C.QTableWidget_Column(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTableWidgetItem(item))))
	}
	return 0
}

func (ptr *QTableWidget) ColumnCount() int {
	if ptr.Pointer() != nil {
		return int(C.QTableWidget_ColumnCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTableWidget) ConnectCurrentCellChanged(f func(currentRow int, currentColumn int, previousRow int, previousColumn int)) {
	if ptr.Pointer() != nil {
		C.QTableWidget_ConnectCurrentCellChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "currentCellChanged", f)
	}
}

func (ptr *QTableWidget) DisconnectCurrentCellChanged() {
	if ptr.Pointer() != nil {
		C.QTableWidget_DisconnectCurrentCellChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "currentCellChanged")
	}
}

//export callbackQTableWidgetCurrentCellChanged
func callbackQTableWidgetCurrentCellChanged(ptrName *C.char, currentRow C.int, currentColumn C.int, previousRow C.int, previousColumn C.int) {
	qt.GetSignal(C.GoString(ptrName), "currentCellChanged").(func(int, int, int, int))(int(currentRow), int(currentColumn), int(previousRow), int(previousColumn))
}

func (ptr *QTableWidget) CurrentColumn() int {
	if ptr.Pointer() != nil {
		return int(C.QTableWidget_CurrentColumn(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTableWidget) CurrentItem() *QTableWidgetItem {
	if ptr.Pointer() != nil {
		return QTableWidgetItemFromPointer(unsafe.Pointer(C.QTableWidget_CurrentItem(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QTableWidget) ConnectCurrentItemChanged(f func(current QTableWidgetItemITF, previous QTableWidgetItemITF)) {
	if ptr.Pointer() != nil {
		C.QTableWidget_ConnectCurrentItemChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "currentItemChanged", f)
	}
}

func (ptr *QTableWidget) DisconnectCurrentItemChanged() {
	if ptr.Pointer() != nil {
		C.QTableWidget_DisconnectCurrentItemChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "currentItemChanged")
	}
}

//export callbackQTableWidgetCurrentItemChanged
func callbackQTableWidgetCurrentItemChanged(ptrName *C.char, current unsafe.Pointer, previous unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "currentItemChanged").(func(*QTableWidgetItem, *QTableWidgetItem))(QTableWidgetItemFromPointer(current), QTableWidgetItemFromPointer(previous))
}

func (ptr *QTableWidget) CurrentRow() int {
	if ptr.Pointer() != nil {
		return int(C.QTableWidget_CurrentRow(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTableWidget) EditItem(item QTableWidgetItemITF) {
	if ptr.Pointer() != nil {
		C.QTableWidget_EditItem(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTableWidgetItem(item)))
	}
}

func (ptr *QTableWidget) HorizontalHeaderItem(column int) *QTableWidgetItem {
	if ptr.Pointer() != nil {
		return QTableWidgetItemFromPointer(unsafe.Pointer(C.QTableWidget_HorizontalHeaderItem(C.QtObjectPtr(ptr.Pointer()), C.int(column))))
	}
	return nil
}

func (ptr *QTableWidget) InsertColumn(column int) {
	if ptr.Pointer() != nil {
		C.QTableWidget_InsertColumn(C.QtObjectPtr(ptr.Pointer()), C.int(column))
	}
}

func (ptr *QTableWidget) InsertRow(row int) {
	if ptr.Pointer() != nil {
		C.QTableWidget_InsertRow(C.QtObjectPtr(ptr.Pointer()), C.int(row))
	}
}

func (ptr *QTableWidget) Item(row int, column int) *QTableWidgetItem {
	if ptr.Pointer() != nil {
		return QTableWidgetItemFromPointer(unsafe.Pointer(C.QTableWidget_Item(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(column))))
	}
	return nil
}

func (ptr *QTableWidget) ConnectItemActivated(f func(item QTableWidgetItemITF)) {
	if ptr.Pointer() != nil {
		C.QTableWidget_ConnectItemActivated(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "itemActivated", f)
	}
}

func (ptr *QTableWidget) DisconnectItemActivated() {
	if ptr.Pointer() != nil {
		C.QTableWidget_DisconnectItemActivated(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "itemActivated")
	}
}

//export callbackQTableWidgetItemActivated
func callbackQTableWidgetItemActivated(ptrName *C.char, item unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "itemActivated").(func(*QTableWidgetItem))(QTableWidgetItemFromPointer(item))
}

func (ptr *QTableWidget) ItemAt2(ax int, ay int) *QTableWidgetItem {
	if ptr.Pointer() != nil {
		return QTableWidgetItemFromPointer(unsafe.Pointer(C.QTableWidget_ItemAt2(C.QtObjectPtr(ptr.Pointer()), C.int(ax), C.int(ay))))
	}
	return nil
}

func (ptr *QTableWidget) ConnectItemChanged(f func(item QTableWidgetItemITF)) {
	if ptr.Pointer() != nil {
		C.QTableWidget_ConnectItemChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "itemChanged", f)
	}
}

func (ptr *QTableWidget) DisconnectItemChanged() {
	if ptr.Pointer() != nil {
		C.QTableWidget_DisconnectItemChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "itemChanged")
	}
}

//export callbackQTableWidgetItemChanged
func callbackQTableWidgetItemChanged(ptrName *C.char, item unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "itemChanged").(func(*QTableWidgetItem))(QTableWidgetItemFromPointer(item))
}

func (ptr *QTableWidget) ConnectItemClicked(f func(item QTableWidgetItemITF)) {
	if ptr.Pointer() != nil {
		C.QTableWidget_ConnectItemClicked(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "itemClicked", f)
	}
}

func (ptr *QTableWidget) DisconnectItemClicked() {
	if ptr.Pointer() != nil {
		C.QTableWidget_DisconnectItemClicked(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "itemClicked")
	}
}

//export callbackQTableWidgetItemClicked
func callbackQTableWidgetItemClicked(ptrName *C.char, item unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "itemClicked").(func(*QTableWidgetItem))(QTableWidgetItemFromPointer(item))
}

func (ptr *QTableWidget) ConnectItemDoubleClicked(f func(item QTableWidgetItemITF)) {
	if ptr.Pointer() != nil {
		C.QTableWidget_ConnectItemDoubleClicked(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "itemDoubleClicked", f)
	}
}

func (ptr *QTableWidget) DisconnectItemDoubleClicked() {
	if ptr.Pointer() != nil {
		C.QTableWidget_DisconnectItemDoubleClicked(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "itemDoubleClicked")
	}
}

//export callbackQTableWidgetItemDoubleClicked
func callbackQTableWidgetItemDoubleClicked(ptrName *C.char, item unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "itemDoubleClicked").(func(*QTableWidgetItem))(QTableWidgetItemFromPointer(item))
}

func (ptr *QTableWidget) ConnectItemEntered(f func(item QTableWidgetItemITF)) {
	if ptr.Pointer() != nil {
		C.QTableWidget_ConnectItemEntered(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "itemEntered", f)
	}
}

func (ptr *QTableWidget) DisconnectItemEntered() {
	if ptr.Pointer() != nil {
		C.QTableWidget_DisconnectItemEntered(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "itemEntered")
	}
}

//export callbackQTableWidgetItemEntered
func callbackQTableWidgetItemEntered(ptrName *C.char, item unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "itemEntered").(func(*QTableWidgetItem))(QTableWidgetItemFromPointer(item))
}

func (ptr *QTableWidget) ConnectItemPressed(f func(item QTableWidgetItemITF)) {
	if ptr.Pointer() != nil {
		C.QTableWidget_ConnectItemPressed(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "itemPressed", f)
	}
}

func (ptr *QTableWidget) DisconnectItemPressed() {
	if ptr.Pointer() != nil {
		C.QTableWidget_DisconnectItemPressed(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "itemPressed")
	}
}

//export callbackQTableWidgetItemPressed
func callbackQTableWidgetItemPressed(ptrName *C.char, item unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "itemPressed").(func(*QTableWidgetItem))(QTableWidgetItemFromPointer(item))
}

func (ptr *QTableWidget) ItemPrototype() *QTableWidgetItem {
	if ptr.Pointer() != nil {
		return QTableWidgetItemFromPointer(unsafe.Pointer(C.QTableWidget_ItemPrototype(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QTableWidget) ConnectItemSelectionChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QTableWidget_ConnectItemSelectionChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "itemSelectionChanged", f)
	}
}

func (ptr *QTableWidget) DisconnectItemSelectionChanged() {
	if ptr.Pointer() != nil {
		C.QTableWidget_DisconnectItemSelectionChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "itemSelectionChanged")
	}
}

//export callbackQTableWidgetItemSelectionChanged
func callbackQTableWidgetItemSelectionChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "itemSelectionChanged").(func())()
}

func (ptr *QTableWidget) OpenPersistentEditor(item QTableWidgetItemITF) {
	if ptr.Pointer() != nil {
		C.QTableWidget_OpenPersistentEditor(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTableWidgetItem(item)))
	}
}

func (ptr *QTableWidget) RemoveCellWidget(row int, column int) {
	if ptr.Pointer() != nil {
		C.QTableWidget_RemoveCellWidget(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(column))
	}
}

func (ptr *QTableWidget) RemoveColumn(column int) {
	if ptr.Pointer() != nil {
		C.QTableWidget_RemoveColumn(C.QtObjectPtr(ptr.Pointer()), C.int(column))
	}
}

func (ptr *QTableWidget) RemoveRow(row int) {
	if ptr.Pointer() != nil {
		C.QTableWidget_RemoveRow(C.QtObjectPtr(ptr.Pointer()), C.int(row))
	}
}

func (ptr *QTableWidget) Row(item QTableWidgetItemITF) int {
	if ptr.Pointer() != nil {
		return int(C.QTableWidget_Row(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTableWidgetItem(item))))
	}
	return 0
}

func (ptr *QTableWidget) RowCount() int {
	if ptr.Pointer() != nil {
		return int(C.QTableWidget_RowCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTableWidget) ScrollToItem(item QTableWidgetItemITF, hint QAbstractItemView__ScrollHint) {
	if ptr.Pointer() != nil {
		C.QTableWidget_ScrollToItem(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTableWidgetItem(item)), C.int(hint))
	}
}

func (ptr *QTableWidget) SetCellWidget(row int, column int, widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QTableWidget_SetCellWidget(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(column), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QTableWidget) SetColumnCount(columns int) {
	if ptr.Pointer() != nil {
		C.QTableWidget_SetColumnCount(C.QtObjectPtr(ptr.Pointer()), C.int(columns))
	}
}

func (ptr *QTableWidget) SetCurrentCell(row int, column int) {
	if ptr.Pointer() != nil {
		C.QTableWidget_SetCurrentCell(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(column))
	}
}

func (ptr *QTableWidget) SetCurrentCell2(row int, column int, command core.QItemSelectionModel__SelectionFlag) {
	if ptr.Pointer() != nil {
		C.QTableWidget_SetCurrentCell2(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(column), C.int(command))
	}
}

func (ptr *QTableWidget) SetCurrentItem(item QTableWidgetItemITF) {
	if ptr.Pointer() != nil {
		C.QTableWidget_SetCurrentItem(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTableWidgetItem(item)))
	}
}

func (ptr *QTableWidget) SetCurrentItem2(item QTableWidgetItemITF, command core.QItemSelectionModel__SelectionFlag) {
	if ptr.Pointer() != nil {
		C.QTableWidget_SetCurrentItem2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTableWidgetItem(item)), C.int(command))
	}
}

func (ptr *QTableWidget) SetHorizontalHeaderItem(column int, item QTableWidgetItemITF) {
	if ptr.Pointer() != nil {
		C.QTableWidget_SetHorizontalHeaderItem(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.QtObjectPtr(PointerFromQTableWidgetItem(item)))
	}
}

func (ptr *QTableWidget) SetHorizontalHeaderLabels(labels []string) {
	if ptr.Pointer() != nil {
		C.QTableWidget_SetHorizontalHeaderLabels(C.QtObjectPtr(ptr.Pointer()), C.CString(strings.Join(labels, "|")))
	}
}

func (ptr *QTableWidget) SetItem(row int, column int, item QTableWidgetItemITF) {
	if ptr.Pointer() != nil {
		C.QTableWidget_SetItem(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(column), C.QtObjectPtr(PointerFromQTableWidgetItem(item)))
	}
}

func (ptr *QTableWidget) SetItemPrototype(item QTableWidgetItemITF) {
	if ptr.Pointer() != nil {
		C.QTableWidget_SetItemPrototype(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTableWidgetItem(item)))
	}
}

func (ptr *QTableWidget) SetRangeSelected(ran QTableWidgetSelectionRangeITF, sele bool) {
	if ptr.Pointer() != nil {
		C.QTableWidget_SetRangeSelected(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTableWidgetSelectionRange(ran)), C.int(qt.GoBoolToInt(sele)))
	}
}

func (ptr *QTableWidget) SetRowCount(rows int) {
	if ptr.Pointer() != nil {
		C.QTableWidget_SetRowCount(C.QtObjectPtr(ptr.Pointer()), C.int(rows))
	}
}

func (ptr *QTableWidget) SetVerticalHeaderItem(row int, item QTableWidgetItemITF) {
	if ptr.Pointer() != nil {
		C.QTableWidget_SetVerticalHeaderItem(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.QtObjectPtr(PointerFromQTableWidgetItem(item)))
	}
}

func (ptr *QTableWidget) SetVerticalHeaderLabels(labels []string) {
	if ptr.Pointer() != nil {
		C.QTableWidget_SetVerticalHeaderLabels(C.QtObjectPtr(ptr.Pointer()), C.CString(strings.Join(labels, "|")))
	}
}

func (ptr *QTableWidget) SortItems(column int, order core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.QTableWidget_SortItems(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.int(order))
	}
}

func (ptr *QTableWidget) TakeHorizontalHeaderItem(column int) *QTableWidgetItem {
	if ptr.Pointer() != nil {
		return QTableWidgetItemFromPointer(unsafe.Pointer(C.QTableWidget_TakeHorizontalHeaderItem(C.QtObjectPtr(ptr.Pointer()), C.int(column))))
	}
	return nil
}

func (ptr *QTableWidget) TakeItem(row int, column int) *QTableWidgetItem {
	if ptr.Pointer() != nil {
		return QTableWidgetItemFromPointer(unsafe.Pointer(C.QTableWidget_TakeItem(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(column))))
	}
	return nil
}

func (ptr *QTableWidget) TakeVerticalHeaderItem(row int) *QTableWidgetItem {
	if ptr.Pointer() != nil {
		return QTableWidgetItemFromPointer(unsafe.Pointer(C.QTableWidget_TakeVerticalHeaderItem(C.QtObjectPtr(ptr.Pointer()), C.int(row))))
	}
	return nil
}

func (ptr *QTableWidget) VerticalHeaderItem(row int) *QTableWidgetItem {
	if ptr.Pointer() != nil {
		return QTableWidgetItemFromPointer(unsafe.Pointer(C.QTableWidget_VerticalHeaderItem(C.QtObjectPtr(ptr.Pointer()), C.int(row))))
	}
	return nil
}

func (ptr *QTableWidget) VisualColumn(logicalColumn int) int {
	if ptr.Pointer() != nil {
		return int(C.QTableWidget_VisualColumn(C.QtObjectPtr(ptr.Pointer()), C.int(logicalColumn)))
	}
	return 0
}

func (ptr *QTableWidget) VisualRow(logicalRow int) int {
	if ptr.Pointer() != nil {
		return int(C.QTableWidget_VisualRow(C.QtObjectPtr(ptr.Pointer()), C.int(logicalRow)))
	}
	return 0
}

func (ptr *QTableWidget) DestroyQTableWidget() {
	if ptr.Pointer() != nil {
		C.QTableWidget_DestroyQTableWidget(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
