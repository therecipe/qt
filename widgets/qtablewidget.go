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

type QTableWidget struct {
	QTableView
}

type QTableWidget_ITF interface {
	QTableView_ITF
	QTableWidget_PTR() *QTableWidget
}

func PointerFromQTableWidget(ptr QTableWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTableWidget_PTR().Pointer()
	}
	return nil
}

func NewQTableWidgetFromPointer(ptr unsafe.Pointer) *QTableWidget {
	var n = new(QTableWidget)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QTableWidget_") {
		n.SetObjectName("QTableWidget_" + qt.Identifier())
	}
	return n
}

func (ptr *QTableWidget) QTableWidget_PTR() *QTableWidget {
	return ptr
}

func (ptr *QTableWidget) ItemAt(point core.QPoint_ITF) *QTableWidgetItem {
	defer qt.Recovering("QTableWidget::itemAt")

	if ptr.Pointer() != nil {
		return NewQTableWidgetItemFromPointer(C.QTableWidget_ItemAt(ptr.Pointer(), core.PointerFromQPoint(point)))
	}
	return nil
}

func (ptr *QTableWidget) ConnectCellActivated(f func(row int, column int)) {
	defer qt.Recovering("connect QTableWidget::cellActivated")

	if ptr.Pointer() != nil {
		C.QTableWidget_ConnectCellActivated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "cellActivated", f)
	}
}

func (ptr *QTableWidget) DisconnectCellActivated() {
	defer qt.Recovering("disconnect QTableWidget::cellActivated")

	if ptr.Pointer() != nil {
		C.QTableWidget_DisconnectCellActivated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "cellActivated")
	}
}

//export callbackQTableWidgetCellActivated
func callbackQTableWidgetCellActivated(ptrName *C.char, row C.int, column C.int) {
	defer qt.Recovering("callback QTableWidget::cellActivated")

	var signal = qt.GetSignal(C.GoString(ptrName), "cellActivated")
	if signal != nil {
		signal.(func(int, int))(int(row), int(column))
	}

}

func (ptr *QTableWidget) ConnectCellChanged(f func(row int, column int)) {
	defer qt.Recovering("connect QTableWidget::cellChanged")

	if ptr.Pointer() != nil {
		C.QTableWidget_ConnectCellChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "cellChanged", f)
	}
}

func (ptr *QTableWidget) DisconnectCellChanged() {
	defer qt.Recovering("disconnect QTableWidget::cellChanged")

	if ptr.Pointer() != nil {
		C.QTableWidget_DisconnectCellChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "cellChanged")
	}
}

//export callbackQTableWidgetCellChanged
func callbackQTableWidgetCellChanged(ptrName *C.char, row C.int, column C.int) {
	defer qt.Recovering("callback QTableWidget::cellChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "cellChanged")
	if signal != nil {
		signal.(func(int, int))(int(row), int(column))
	}

}

func (ptr *QTableWidget) ConnectCellClicked(f func(row int, column int)) {
	defer qt.Recovering("connect QTableWidget::cellClicked")

	if ptr.Pointer() != nil {
		C.QTableWidget_ConnectCellClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "cellClicked", f)
	}
}

func (ptr *QTableWidget) DisconnectCellClicked() {
	defer qt.Recovering("disconnect QTableWidget::cellClicked")

	if ptr.Pointer() != nil {
		C.QTableWidget_DisconnectCellClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "cellClicked")
	}
}

//export callbackQTableWidgetCellClicked
func callbackQTableWidgetCellClicked(ptrName *C.char, row C.int, column C.int) {
	defer qt.Recovering("callback QTableWidget::cellClicked")

	var signal = qt.GetSignal(C.GoString(ptrName), "cellClicked")
	if signal != nil {
		signal.(func(int, int))(int(row), int(column))
	}

}

func (ptr *QTableWidget) ConnectCellDoubleClicked(f func(row int, column int)) {
	defer qt.Recovering("connect QTableWidget::cellDoubleClicked")

	if ptr.Pointer() != nil {
		C.QTableWidget_ConnectCellDoubleClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "cellDoubleClicked", f)
	}
}

func (ptr *QTableWidget) DisconnectCellDoubleClicked() {
	defer qt.Recovering("disconnect QTableWidget::cellDoubleClicked")

	if ptr.Pointer() != nil {
		C.QTableWidget_DisconnectCellDoubleClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "cellDoubleClicked")
	}
}

//export callbackQTableWidgetCellDoubleClicked
func callbackQTableWidgetCellDoubleClicked(ptrName *C.char, row C.int, column C.int) {
	defer qt.Recovering("callback QTableWidget::cellDoubleClicked")

	var signal = qt.GetSignal(C.GoString(ptrName), "cellDoubleClicked")
	if signal != nil {
		signal.(func(int, int))(int(row), int(column))
	}

}

func (ptr *QTableWidget) ConnectCellEntered(f func(row int, column int)) {
	defer qt.Recovering("connect QTableWidget::cellEntered")

	if ptr.Pointer() != nil {
		C.QTableWidget_ConnectCellEntered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "cellEntered", f)
	}
}

func (ptr *QTableWidget) DisconnectCellEntered() {
	defer qt.Recovering("disconnect QTableWidget::cellEntered")

	if ptr.Pointer() != nil {
		C.QTableWidget_DisconnectCellEntered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "cellEntered")
	}
}

//export callbackQTableWidgetCellEntered
func callbackQTableWidgetCellEntered(ptrName *C.char, row C.int, column C.int) {
	defer qt.Recovering("callback QTableWidget::cellEntered")

	var signal = qt.GetSignal(C.GoString(ptrName), "cellEntered")
	if signal != nil {
		signal.(func(int, int))(int(row), int(column))
	}

}

func (ptr *QTableWidget) ConnectCellPressed(f func(row int, column int)) {
	defer qt.Recovering("connect QTableWidget::cellPressed")

	if ptr.Pointer() != nil {
		C.QTableWidget_ConnectCellPressed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "cellPressed", f)
	}
}

func (ptr *QTableWidget) DisconnectCellPressed() {
	defer qt.Recovering("disconnect QTableWidget::cellPressed")

	if ptr.Pointer() != nil {
		C.QTableWidget_DisconnectCellPressed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "cellPressed")
	}
}

//export callbackQTableWidgetCellPressed
func callbackQTableWidgetCellPressed(ptrName *C.char, row C.int, column C.int) {
	defer qt.Recovering("callback QTableWidget::cellPressed")

	var signal = qt.GetSignal(C.GoString(ptrName), "cellPressed")
	if signal != nil {
		signal.(func(int, int))(int(row), int(column))
	}

}

func (ptr *QTableWidget) CellWidget(row int, column int) *QWidget {
	defer qt.Recovering("QTableWidget::cellWidget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QTableWidget_CellWidget(ptr.Pointer(), C.int(row), C.int(column)))
	}
	return nil
}

func (ptr *QTableWidget) Clear() {
	defer qt.Recovering("QTableWidget::clear")

	if ptr.Pointer() != nil {
		C.QTableWidget_Clear(ptr.Pointer())
	}
}

func (ptr *QTableWidget) ClearContents() {
	defer qt.Recovering("QTableWidget::clearContents")

	if ptr.Pointer() != nil {
		C.QTableWidget_ClearContents(ptr.Pointer())
	}
}

func (ptr *QTableWidget) ClosePersistentEditor(item QTableWidgetItem_ITF) {
	defer qt.Recovering("QTableWidget::closePersistentEditor")

	if ptr.Pointer() != nil {
		C.QTableWidget_ClosePersistentEditor(ptr.Pointer(), PointerFromQTableWidgetItem(item))
	}
}

func (ptr *QTableWidget) Column(item QTableWidgetItem_ITF) int {
	defer qt.Recovering("QTableWidget::column")

	if ptr.Pointer() != nil {
		return int(C.QTableWidget_Column(ptr.Pointer(), PointerFromQTableWidgetItem(item)))
	}
	return 0
}

func (ptr *QTableWidget) ColumnCount() int {
	defer qt.Recovering("QTableWidget::columnCount")

	if ptr.Pointer() != nil {
		return int(C.QTableWidget_ColumnCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTableWidget) ConnectCurrentCellChanged(f func(currentRow int, currentColumn int, previousRow int, previousColumn int)) {
	defer qt.Recovering("connect QTableWidget::currentCellChanged")

	if ptr.Pointer() != nil {
		C.QTableWidget_ConnectCurrentCellChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentCellChanged", f)
	}
}

func (ptr *QTableWidget) DisconnectCurrentCellChanged() {
	defer qt.Recovering("disconnect QTableWidget::currentCellChanged")

	if ptr.Pointer() != nil {
		C.QTableWidget_DisconnectCurrentCellChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentCellChanged")
	}
}

//export callbackQTableWidgetCurrentCellChanged
func callbackQTableWidgetCurrentCellChanged(ptrName *C.char, currentRow C.int, currentColumn C.int, previousRow C.int, previousColumn C.int) {
	defer qt.Recovering("callback QTableWidget::currentCellChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "currentCellChanged")
	if signal != nil {
		signal.(func(int, int, int, int))(int(currentRow), int(currentColumn), int(previousRow), int(previousColumn))
	}

}

func (ptr *QTableWidget) CurrentColumn() int {
	defer qt.Recovering("QTableWidget::currentColumn")

	if ptr.Pointer() != nil {
		return int(C.QTableWidget_CurrentColumn(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTableWidget) CurrentItem() *QTableWidgetItem {
	defer qt.Recovering("QTableWidget::currentItem")

	if ptr.Pointer() != nil {
		return NewQTableWidgetItemFromPointer(C.QTableWidget_CurrentItem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTableWidget) ConnectCurrentItemChanged(f func(current *QTableWidgetItem, previous *QTableWidgetItem)) {
	defer qt.Recovering("connect QTableWidget::currentItemChanged")

	if ptr.Pointer() != nil {
		C.QTableWidget_ConnectCurrentItemChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentItemChanged", f)
	}
}

func (ptr *QTableWidget) DisconnectCurrentItemChanged() {
	defer qt.Recovering("disconnect QTableWidget::currentItemChanged")

	if ptr.Pointer() != nil {
		C.QTableWidget_DisconnectCurrentItemChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentItemChanged")
	}
}

//export callbackQTableWidgetCurrentItemChanged
func callbackQTableWidgetCurrentItemChanged(ptrName *C.char, current unsafe.Pointer, previous unsafe.Pointer) {
	defer qt.Recovering("callback QTableWidget::currentItemChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "currentItemChanged")
	if signal != nil {
		signal.(func(*QTableWidgetItem, *QTableWidgetItem))(NewQTableWidgetItemFromPointer(current), NewQTableWidgetItemFromPointer(previous))
	}

}

func (ptr *QTableWidget) CurrentRow() int {
	defer qt.Recovering("QTableWidget::currentRow")

	if ptr.Pointer() != nil {
		return int(C.QTableWidget_CurrentRow(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTableWidget) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QTableWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QTableWidget) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QTableWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQTableWidgetDropEvent
func callbackQTableWidgetDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableWidget::dropEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "dropEvent")
	if signal != nil {
		defer signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableWidget) EditItem(item QTableWidgetItem_ITF) {
	defer qt.Recovering("QTableWidget::editItem")

	if ptr.Pointer() != nil {
		C.QTableWidget_EditItem(ptr.Pointer(), PointerFromQTableWidgetItem(item))
	}
}

func (ptr *QTableWidget) HorizontalHeaderItem(column int) *QTableWidgetItem {
	defer qt.Recovering("QTableWidget::horizontalHeaderItem")

	if ptr.Pointer() != nil {
		return NewQTableWidgetItemFromPointer(C.QTableWidget_HorizontalHeaderItem(ptr.Pointer(), C.int(column)))
	}
	return nil
}

func (ptr *QTableWidget) InsertColumn(column int) {
	defer qt.Recovering("QTableWidget::insertColumn")

	if ptr.Pointer() != nil {
		C.QTableWidget_InsertColumn(ptr.Pointer(), C.int(column))
	}
}

func (ptr *QTableWidget) InsertRow(row int) {
	defer qt.Recovering("QTableWidget::insertRow")

	if ptr.Pointer() != nil {
		C.QTableWidget_InsertRow(ptr.Pointer(), C.int(row))
	}
}

func (ptr *QTableWidget) Item(row int, column int) *QTableWidgetItem {
	defer qt.Recovering("QTableWidget::item")

	if ptr.Pointer() != nil {
		return NewQTableWidgetItemFromPointer(C.QTableWidget_Item(ptr.Pointer(), C.int(row), C.int(column)))
	}
	return nil
}

func (ptr *QTableWidget) ConnectItemActivated(f func(item *QTableWidgetItem)) {
	defer qt.Recovering("connect QTableWidget::itemActivated")

	if ptr.Pointer() != nil {
		C.QTableWidget_ConnectItemActivated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "itemActivated", f)
	}
}

func (ptr *QTableWidget) DisconnectItemActivated() {
	defer qt.Recovering("disconnect QTableWidget::itemActivated")

	if ptr.Pointer() != nil {
		C.QTableWidget_DisconnectItemActivated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "itemActivated")
	}
}

//export callbackQTableWidgetItemActivated
func callbackQTableWidgetItemActivated(ptrName *C.char, item unsafe.Pointer) {
	defer qt.Recovering("callback QTableWidget::itemActivated")

	var signal = qt.GetSignal(C.GoString(ptrName), "itemActivated")
	if signal != nil {
		signal.(func(*QTableWidgetItem))(NewQTableWidgetItemFromPointer(item))
	}

}

func (ptr *QTableWidget) ItemAt2(ax int, ay int) *QTableWidgetItem {
	defer qt.Recovering("QTableWidget::itemAt")

	if ptr.Pointer() != nil {
		return NewQTableWidgetItemFromPointer(C.QTableWidget_ItemAt2(ptr.Pointer(), C.int(ax), C.int(ay)))
	}
	return nil
}

func (ptr *QTableWidget) ConnectItemChanged(f func(item *QTableWidgetItem)) {
	defer qt.Recovering("connect QTableWidget::itemChanged")

	if ptr.Pointer() != nil {
		C.QTableWidget_ConnectItemChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "itemChanged", f)
	}
}

func (ptr *QTableWidget) DisconnectItemChanged() {
	defer qt.Recovering("disconnect QTableWidget::itemChanged")

	if ptr.Pointer() != nil {
		C.QTableWidget_DisconnectItemChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "itemChanged")
	}
}

//export callbackQTableWidgetItemChanged
func callbackQTableWidgetItemChanged(ptrName *C.char, item unsafe.Pointer) {
	defer qt.Recovering("callback QTableWidget::itemChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "itemChanged")
	if signal != nil {
		signal.(func(*QTableWidgetItem))(NewQTableWidgetItemFromPointer(item))
	}

}

func (ptr *QTableWidget) ConnectItemClicked(f func(item *QTableWidgetItem)) {
	defer qt.Recovering("connect QTableWidget::itemClicked")

	if ptr.Pointer() != nil {
		C.QTableWidget_ConnectItemClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "itemClicked", f)
	}
}

func (ptr *QTableWidget) DisconnectItemClicked() {
	defer qt.Recovering("disconnect QTableWidget::itemClicked")

	if ptr.Pointer() != nil {
		C.QTableWidget_DisconnectItemClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "itemClicked")
	}
}

//export callbackQTableWidgetItemClicked
func callbackQTableWidgetItemClicked(ptrName *C.char, item unsafe.Pointer) {
	defer qt.Recovering("callback QTableWidget::itemClicked")

	var signal = qt.GetSignal(C.GoString(ptrName), "itemClicked")
	if signal != nil {
		signal.(func(*QTableWidgetItem))(NewQTableWidgetItemFromPointer(item))
	}

}

func (ptr *QTableWidget) ConnectItemDoubleClicked(f func(item *QTableWidgetItem)) {
	defer qt.Recovering("connect QTableWidget::itemDoubleClicked")

	if ptr.Pointer() != nil {
		C.QTableWidget_ConnectItemDoubleClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "itemDoubleClicked", f)
	}
}

func (ptr *QTableWidget) DisconnectItemDoubleClicked() {
	defer qt.Recovering("disconnect QTableWidget::itemDoubleClicked")

	if ptr.Pointer() != nil {
		C.QTableWidget_DisconnectItemDoubleClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "itemDoubleClicked")
	}
}

//export callbackQTableWidgetItemDoubleClicked
func callbackQTableWidgetItemDoubleClicked(ptrName *C.char, item unsafe.Pointer) {
	defer qt.Recovering("callback QTableWidget::itemDoubleClicked")

	var signal = qt.GetSignal(C.GoString(ptrName), "itemDoubleClicked")
	if signal != nil {
		signal.(func(*QTableWidgetItem))(NewQTableWidgetItemFromPointer(item))
	}

}

func (ptr *QTableWidget) ConnectItemEntered(f func(item *QTableWidgetItem)) {
	defer qt.Recovering("connect QTableWidget::itemEntered")

	if ptr.Pointer() != nil {
		C.QTableWidget_ConnectItemEntered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "itemEntered", f)
	}
}

func (ptr *QTableWidget) DisconnectItemEntered() {
	defer qt.Recovering("disconnect QTableWidget::itemEntered")

	if ptr.Pointer() != nil {
		C.QTableWidget_DisconnectItemEntered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "itemEntered")
	}
}

//export callbackQTableWidgetItemEntered
func callbackQTableWidgetItemEntered(ptrName *C.char, item unsafe.Pointer) {
	defer qt.Recovering("callback QTableWidget::itemEntered")

	var signal = qt.GetSignal(C.GoString(ptrName), "itemEntered")
	if signal != nil {
		signal.(func(*QTableWidgetItem))(NewQTableWidgetItemFromPointer(item))
	}

}

func (ptr *QTableWidget) ConnectItemPressed(f func(item *QTableWidgetItem)) {
	defer qt.Recovering("connect QTableWidget::itemPressed")

	if ptr.Pointer() != nil {
		C.QTableWidget_ConnectItemPressed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "itemPressed", f)
	}
}

func (ptr *QTableWidget) DisconnectItemPressed() {
	defer qt.Recovering("disconnect QTableWidget::itemPressed")

	if ptr.Pointer() != nil {
		C.QTableWidget_DisconnectItemPressed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "itemPressed")
	}
}

//export callbackQTableWidgetItemPressed
func callbackQTableWidgetItemPressed(ptrName *C.char, item unsafe.Pointer) {
	defer qt.Recovering("callback QTableWidget::itemPressed")

	var signal = qt.GetSignal(C.GoString(ptrName), "itemPressed")
	if signal != nil {
		signal.(func(*QTableWidgetItem))(NewQTableWidgetItemFromPointer(item))
	}

}

func (ptr *QTableWidget) ItemPrototype() *QTableWidgetItem {
	defer qt.Recovering("QTableWidget::itemPrototype")

	if ptr.Pointer() != nil {
		return NewQTableWidgetItemFromPointer(C.QTableWidget_ItemPrototype(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTableWidget) ConnectItemSelectionChanged(f func()) {
	defer qt.Recovering("connect QTableWidget::itemSelectionChanged")

	if ptr.Pointer() != nil {
		C.QTableWidget_ConnectItemSelectionChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "itemSelectionChanged", f)
	}
}

func (ptr *QTableWidget) DisconnectItemSelectionChanged() {
	defer qt.Recovering("disconnect QTableWidget::itemSelectionChanged")

	if ptr.Pointer() != nil {
		C.QTableWidget_DisconnectItemSelectionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "itemSelectionChanged")
	}
}

//export callbackQTableWidgetItemSelectionChanged
func callbackQTableWidgetItemSelectionChanged(ptrName *C.char) {
	defer qt.Recovering("callback QTableWidget::itemSelectionChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "itemSelectionChanged")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QTableWidget) OpenPersistentEditor(item QTableWidgetItem_ITF) {
	defer qt.Recovering("QTableWidget::openPersistentEditor")

	if ptr.Pointer() != nil {
		C.QTableWidget_OpenPersistentEditor(ptr.Pointer(), PointerFromQTableWidgetItem(item))
	}
}

func (ptr *QTableWidget) RemoveCellWidget(row int, column int) {
	defer qt.Recovering("QTableWidget::removeCellWidget")

	if ptr.Pointer() != nil {
		C.QTableWidget_RemoveCellWidget(ptr.Pointer(), C.int(row), C.int(column))
	}
}

func (ptr *QTableWidget) RemoveColumn(column int) {
	defer qt.Recovering("QTableWidget::removeColumn")

	if ptr.Pointer() != nil {
		C.QTableWidget_RemoveColumn(ptr.Pointer(), C.int(column))
	}
}

func (ptr *QTableWidget) RemoveRow(row int) {
	defer qt.Recovering("QTableWidget::removeRow")

	if ptr.Pointer() != nil {
		C.QTableWidget_RemoveRow(ptr.Pointer(), C.int(row))
	}
}

func (ptr *QTableWidget) Row(item QTableWidgetItem_ITF) int {
	defer qt.Recovering("QTableWidget::row")

	if ptr.Pointer() != nil {
		return int(C.QTableWidget_Row(ptr.Pointer(), PointerFromQTableWidgetItem(item)))
	}
	return 0
}

func (ptr *QTableWidget) RowCount() int {
	defer qt.Recovering("QTableWidget::rowCount")

	if ptr.Pointer() != nil {
		return int(C.QTableWidget_RowCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTableWidget) ScrollToItem(item QTableWidgetItem_ITF, hint QAbstractItemView__ScrollHint) {
	defer qt.Recovering("QTableWidget::scrollToItem")

	if ptr.Pointer() != nil {
		C.QTableWidget_ScrollToItem(ptr.Pointer(), PointerFromQTableWidgetItem(item), C.int(hint))
	}
}

func (ptr *QTableWidget) SetCellWidget(row int, column int, widget QWidget_ITF) {
	defer qt.Recovering("QTableWidget::setCellWidget")

	if ptr.Pointer() != nil {
		C.QTableWidget_SetCellWidget(ptr.Pointer(), C.int(row), C.int(column), PointerFromQWidget(widget))
	}
}

func (ptr *QTableWidget) SetColumnCount(columns int) {
	defer qt.Recovering("QTableWidget::setColumnCount")

	if ptr.Pointer() != nil {
		C.QTableWidget_SetColumnCount(ptr.Pointer(), C.int(columns))
	}
}

func (ptr *QTableWidget) SetCurrentCell(row int, column int) {
	defer qt.Recovering("QTableWidget::setCurrentCell")

	if ptr.Pointer() != nil {
		C.QTableWidget_SetCurrentCell(ptr.Pointer(), C.int(row), C.int(column))
	}
}

func (ptr *QTableWidget) SetCurrentCell2(row int, column int, command core.QItemSelectionModel__SelectionFlag) {
	defer qt.Recovering("QTableWidget::setCurrentCell")

	if ptr.Pointer() != nil {
		C.QTableWidget_SetCurrentCell2(ptr.Pointer(), C.int(row), C.int(column), C.int(command))
	}
}

func (ptr *QTableWidget) SetCurrentItem(item QTableWidgetItem_ITF) {
	defer qt.Recovering("QTableWidget::setCurrentItem")

	if ptr.Pointer() != nil {
		C.QTableWidget_SetCurrentItem(ptr.Pointer(), PointerFromQTableWidgetItem(item))
	}
}

func (ptr *QTableWidget) SetCurrentItem2(item QTableWidgetItem_ITF, command core.QItemSelectionModel__SelectionFlag) {
	defer qt.Recovering("QTableWidget::setCurrentItem")

	if ptr.Pointer() != nil {
		C.QTableWidget_SetCurrentItem2(ptr.Pointer(), PointerFromQTableWidgetItem(item), C.int(command))
	}
}

func (ptr *QTableWidget) SetHorizontalHeaderItem(column int, item QTableWidgetItem_ITF) {
	defer qt.Recovering("QTableWidget::setHorizontalHeaderItem")

	if ptr.Pointer() != nil {
		C.QTableWidget_SetHorizontalHeaderItem(ptr.Pointer(), C.int(column), PointerFromQTableWidgetItem(item))
	}
}

func (ptr *QTableWidget) SetHorizontalHeaderLabels(labels []string) {
	defer qt.Recovering("QTableWidget::setHorizontalHeaderLabels")

	if ptr.Pointer() != nil {
		C.QTableWidget_SetHorizontalHeaderLabels(ptr.Pointer(), C.CString(strings.Join(labels, ",,,")))
	}
}

func (ptr *QTableWidget) SetItem(row int, column int, item QTableWidgetItem_ITF) {
	defer qt.Recovering("QTableWidget::setItem")

	if ptr.Pointer() != nil {
		C.QTableWidget_SetItem(ptr.Pointer(), C.int(row), C.int(column), PointerFromQTableWidgetItem(item))
	}
}

func (ptr *QTableWidget) SetItemPrototype(item QTableWidgetItem_ITF) {
	defer qt.Recovering("QTableWidget::setItemPrototype")

	if ptr.Pointer() != nil {
		C.QTableWidget_SetItemPrototype(ptr.Pointer(), PointerFromQTableWidgetItem(item))
	}
}

func (ptr *QTableWidget) SetRangeSelected(ran QTableWidgetSelectionRange_ITF, sele bool) {
	defer qt.Recovering("QTableWidget::setRangeSelected")

	if ptr.Pointer() != nil {
		C.QTableWidget_SetRangeSelected(ptr.Pointer(), PointerFromQTableWidgetSelectionRange(ran), C.int(qt.GoBoolToInt(sele)))
	}
}

func (ptr *QTableWidget) SetRowCount(rows int) {
	defer qt.Recovering("QTableWidget::setRowCount")

	if ptr.Pointer() != nil {
		C.QTableWidget_SetRowCount(ptr.Pointer(), C.int(rows))
	}
}

func (ptr *QTableWidget) SetVerticalHeaderItem(row int, item QTableWidgetItem_ITF) {
	defer qt.Recovering("QTableWidget::setVerticalHeaderItem")

	if ptr.Pointer() != nil {
		C.QTableWidget_SetVerticalHeaderItem(ptr.Pointer(), C.int(row), PointerFromQTableWidgetItem(item))
	}
}

func (ptr *QTableWidget) SetVerticalHeaderLabels(labels []string) {
	defer qt.Recovering("QTableWidget::setVerticalHeaderLabels")

	if ptr.Pointer() != nil {
		C.QTableWidget_SetVerticalHeaderLabels(ptr.Pointer(), C.CString(strings.Join(labels, ",,,")))
	}
}

func (ptr *QTableWidget) SortItems(column int, order core.Qt__SortOrder) {
	defer qt.Recovering("QTableWidget::sortItems")

	if ptr.Pointer() != nil {
		C.QTableWidget_SortItems(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QTableWidget) TakeHorizontalHeaderItem(column int) *QTableWidgetItem {
	defer qt.Recovering("QTableWidget::takeHorizontalHeaderItem")

	if ptr.Pointer() != nil {
		return NewQTableWidgetItemFromPointer(C.QTableWidget_TakeHorizontalHeaderItem(ptr.Pointer(), C.int(column)))
	}
	return nil
}

func (ptr *QTableWidget) TakeItem(row int, column int) *QTableWidgetItem {
	defer qt.Recovering("QTableWidget::takeItem")

	if ptr.Pointer() != nil {
		return NewQTableWidgetItemFromPointer(C.QTableWidget_TakeItem(ptr.Pointer(), C.int(row), C.int(column)))
	}
	return nil
}

func (ptr *QTableWidget) TakeVerticalHeaderItem(row int) *QTableWidgetItem {
	defer qt.Recovering("QTableWidget::takeVerticalHeaderItem")

	if ptr.Pointer() != nil {
		return NewQTableWidgetItemFromPointer(C.QTableWidget_TakeVerticalHeaderItem(ptr.Pointer(), C.int(row)))
	}
	return nil
}

func (ptr *QTableWidget) VerticalHeaderItem(row int) *QTableWidgetItem {
	defer qt.Recovering("QTableWidget::verticalHeaderItem")

	if ptr.Pointer() != nil {
		return NewQTableWidgetItemFromPointer(C.QTableWidget_VerticalHeaderItem(ptr.Pointer(), C.int(row)))
	}
	return nil
}

func (ptr *QTableWidget) VisualColumn(logicalColumn int) int {
	defer qt.Recovering("QTableWidget::visualColumn")

	if ptr.Pointer() != nil {
		return int(C.QTableWidget_VisualColumn(ptr.Pointer(), C.int(logicalColumn)))
	}
	return 0
}

func (ptr *QTableWidget) VisualItemRect(item QTableWidgetItem_ITF) *core.QRect {
	defer qt.Recovering("QTableWidget::visualItemRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QTableWidget_VisualItemRect(ptr.Pointer(), PointerFromQTableWidgetItem(item)))
	}
	return nil
}

func (ptr *QTableWidget) VisualRow(logicalRow int) int {
	defer qt.Recovering("QTableWidget::visualRow")

	if ptr.Pointer() != nil {
		return int(C.QTableWidget_VisualRow(ptr.Pointer(), C.int(logicalRow)))
	}
	return 0
}

func (ptr *QTableWidget) DestroyQTableWidget() {
	defer qt.Recovering("QTableWidget::~QTableWidget")

	if ptr.Pointer() != nil {
		C.QTableWidget_DestroyQTableWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
