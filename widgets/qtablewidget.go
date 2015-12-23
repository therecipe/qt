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

	if signal := qt.GetSignal(C.GoString(ptrName), "cellActivated"); signal != nil {
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

	if signal := qt.GetSignal(C.GoString(ptrName), "cellChanged"); signal != nil {
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

	if signal := qt.GetSignal(C.GoString(ptrName), "cellClicked"); signal != nil {
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

	if signal := qt.GetSignal(C.GoString(ptrName), "cellDoubleClicked"); signal != nil {
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

	if signal := qt.GetSignal(C.GoString(ptrName), "cellEntered"); signal != nil {
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

	if signal := qt.GetSignal(C.GoString(ptrName), "cellPressed"); signal != nil {
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

	if signal := qt.GetSignal(C.GoString(ptrName), "currentCellChanged"); signal != nil {
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

	if signal := qt.GetSignal(C.GoString(ptrName), "currentItemChanged"); signal != nil {
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

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
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

	if signal := qt.GetSignal(C.GoString(ptrName), "itemActivated"); signal != nil {
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

	if signal := qt.GetSignal(C.GoString(ptrName), "itemChanged"); signal != nil {
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

	if signal := qt.GetSignal(C.GoString(ptrName), "itemClicked"); signal != nil {
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

	if signal := qt.GetSignal(C.GoString(ptrName), "itemDoubleClicked"); signal != nil {
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

	if signal := qt.GetSignal(C.GoString(ptrName), "itemEntered"); signal != nil {
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

	if signal := qt.GetSignal(C.GoString(ptrName), "itemPressed"); signal != nil {
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

	if signal := qt.GetSignal(C.GoString(ptrName), "itemSelectionChanged"); signal != nil {
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

func (ptr *QTableWidget) ConnectSetSelection(f func(rect *core.QRect, flags core.QItemSelectionModel__SelectionFlag)) {
	defer qt.Recovering("connect QTableWidget::setSelection")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setSelection", f)
	}
}

func (ptr *QTableWidget) DisconnectSetSelection() {
	defer qt.Recovering("disconnect QTableWidget::setSelection")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setSelection")
	}
}

//export callbackQTableWidgetSetSelection
func callbackQTableWidgetSetSelection(ptrName *C.char, rect unsafe.Pointer, flags C.int) bool {
	defer qt.Recovering("callback QTableWidget::setSelection")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSelection"); signal != nil {
		signal.(func(*core.QRect, core.QItemSelectionModel__SelectionFlag))(core.NewQRectFromPointer(rect), core.QItemSelectionModel__SelectionFlag(flags))
		return true
	}
	return false

}

func (ptr *QTableWidget) ConnectCurrentChanged(f func(current *core.QModelIndex, previous *core.QModelIndex)) {
	defer qt.Recovering("connect QTableWidget::currentChanged")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "currentChanged", f)
	}
}

func (ptr *QTableWidget) DisconnectCurrentChanged() {
	defer qt.Recovering("disconnect QTableWidget::currentChanged")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "currentChanged")
	}
}

//export callbackQTableWidgetCurrentChanged
func callbackQTableWidgetCurrentChanged(ptrName *C.char, current unsafe.Pointer, previous unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableWidget::currentChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "currentChanged"); signal != nil {
		signal.(func(*core.QModelIndex, *core.QModelIndex))(core.NewQModelIndexFromPointer(current), core.NewQModelIndexFromPointer(previous))
		return true
	}
	return false

}

func (ptr *QTableWidget) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QTableWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QTableWidget) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QTableWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQTableWidgetPaintEvent
func callbackQTableWidgetPaintEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableWidget::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableWidget) ConnectSetRootIndex(f func(index *core.QModelIndex)) {
	defer qt.Recovering("connect QTableWidget::setRootIndex")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setRootIndex", f)
	}
}

func (ptr *QTableWidget) DisconnectSetRootIndex() {
	defer qt.Recovering("disconnect QTableWidget::setRootIndex")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setRootIndex")
	}
}

//export callbackQTableWidgetSetRootIndex
func callbackQTableWidgetSetRootIndex(ptrName *C.char, index unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableWidget::setRootIndex")

	if signal := qt.GetSignal(C.GoString(ptrName), "setRootIndex"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
		return true
	}
	return false

}

func (ptr *QTableWidget) ConnectSetSelectionModel(f func(selectionModel *core.QItemSelectionModel)) {
	defer qt.Recovering("connect QTableWidget::setSelectionModel")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setSelectionModel", f)
	}
}

func (ptr *QTableWidget) DisconnectSetSelectionModel() {
	defer qt.Recovering("disconnect QTableWidget::setSelectionModel")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setSelectionModel")
	}
}

//export callbackQTableWidgetSetSelectionModel
func callbackQTableWidgetSetSelectionModel(ptrName *C.char, selectionModel unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableWidget::setSelectionModel")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSelectionModel"); signal != nil {
		signal.(func(*core.QItemSelectionModel))(core.NewQItemSelectionModelFromPointer(selectionModel))
		return true
	}
	return false

}

func (ptr *QTableWidget) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QTableWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QTableWidget) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QTableWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQTableWidgetTimerEvent
func callbackQTableWidgetTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableWidget::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableWidget) ConnectUpdateGeometries(f func()) {
	defer qt.Recovering("connect QTableWidget::updateGeometries")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updateGeometries", f)
	}
}

func (ptr *QTableWidget) DisconnectUpdateGeometries() {
	defer qt.Recovering("disconnect QTableWidget::updateGeometries")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updateGeometries")
	}
}

//export callbackQTableWidgetUpdateGeometries
func callbackQTableWidgetUpdateGeometries(ptrName *C.char) bool {
	defer qt.Recovering("callback QTableWidget::updateGeometries")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateGeometries"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QTableWidget) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QTableWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QTableWidget) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QTableWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQTableWidgetDragLeaveEvent
func callbackQTableWidgetDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableWidget::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableWidget) ConnectCloseEditor(f func(editor *QWidget, hint QAbstractItemDelegate__EndEditHint)) {
	defer qt.Recovering("connect QTableWidget::closeEditor")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEditor", f)
	}
}

func (ptr *QTableWidget) DisconnectCloseEditor() {
	defer qt.Recovering("disconnect QTableWidget::closeEditor")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEditor")
	}
}

//export callbackQTableWidgetCloseEditor
func callbackQTableWidgetCloseEditor(ptrName *C.char, editor unsafe.Pointer, hint C.int) bool {
	defer qt.Recovering("callback QTableWidget::closeEditor")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEditor"); signal != nil {
		signal.(func(*QWidget, QAbstractItemDelegate__EndEditHint))(NewQWidgetFromPointer(editor), QAbstractItemDelegate__EndEditHint(hint))
		return true
	}
	return false

}

func (ptr *QTableWidget) ConnectCommitData(f func(editor *QWidget)) {
	defer qt.Recovering("connect QTableWidget::commitData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "commitData", f)
	}
}

func (ptr *QTableWidget) DisconnectCommitData() {
	defer qt.Recovering("disconnect QTableWidget::commitData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "commitData")
	}
}

//export callbackQTableWidgetCommitData
func callbackQTableWidgetCommitData(ptrName *C.char, editor unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableWidget::commitData")

	if signal := qt.GetSignal(C.GoString(ptrName), "commitData"); signal != nil {
		signal.(func(*QWidget))(NewQWidgetFromPointer(editor))
		return true
	}
	return false

}

func (ptr *QTableWidget) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QTableWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QTableWidget) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QTableWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQTableWidgetDragEnterEvent
func callbackQTableWidgetDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableWidget::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableWidget) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QTableWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QTableWidget) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QTableWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQTableWidgetDragMoveEvent
func callbackQTableWidgetDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableWidget::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableWidget) ConnectEditorDestroyed(f func(editor *core.QObject)) {
	defer qt.Recovering("connect QTableWidget::editorDestroyed")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "editorDestroyed", f)
	}
}

func (ptr *QTableWidget) DisconnectEditorDestroyed() {
	defer qt.Recovering("disconnect QTableWidget::editorDestroyed")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "editorDestroyed")
	}
}

//export callbackQTableWidgetEditorDestroyed
func callbackQTableWidgetEditorDestroyed(ptrName *C.char, editor unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableWidget::editorDestroyed")

	if signal := qt.GetSignal(C.GoString(ptrName), "editorDestroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(editor))
		return true
	}
	return false

}

func (ptr *QTableWidget) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QTableWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QTableWidget) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QTableWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQTableWidgetFocusInEvent
func callbackQTableWidgetFocusInEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableWidget::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableWidget) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QTableWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QTableWidget) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QTableWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQTableWidgetFocusOutEvent
func callbackQTableWidgetFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableWidget::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableWidget) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QTableWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QTableWidget) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QTableWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQTableWidgetInputMethodEvent
func callbackQTableWidgetInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableWidget::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableWidget) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QTableWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QTableWidget) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QTableWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQTableWidgetKeyPressEvent
func callbackQTableWidgetKeyPressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableWidget::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableWidget) ConnectKeyboardSearch(f func(search string)) {
	defer qt.Recovering("connect QTableWidget::keyboardSearch")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyboardSearch", f)
	}
}

func (ptr *QTableWidget) DisconnectKeyboardSearch() {
	defer qt.Recovering("disconnect QTableWidget::keyboardSearch")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyboardSearch")
	}
}

//export callbackQTableWidgetKeyboardSearch
func callbackQTableWidgetKeyboardSearch(ptrName *C.char, search *C.char) bool {
	defer qt.Recovering("callback QTableWidget::keyboardSearch")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyboardSearch"); signal != nil {
		signal.(func(string))(C.GoString(search))
		return true
	}
	return false

}

func (ptr *QTableWidget) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QTableWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QTableWidget) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QTableWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQTableWidgetMouseDoubleClickEvent
func callbackQTableWidgetMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableWidget::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableWidget) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QTableWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QTableWidget) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QTableWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQTableWidgetMouseMoveEvent
func callbackQTableWidgetMouseMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableWidget::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableWidget) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QTableWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QTableWidget) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QTableWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQTableWidgetMousePressEvent
func callbackQTableWidgetMousePressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableWidget::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableWidget) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QTableWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QTableWidget) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QTableWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQTableWidgetMouseReleaseEvent
func callbackQTableWidgetMouseReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableWidget::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableWidget) ConnectReset(f func()) {
	defer qt.Recovering("connect QTableWidget::reset")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "reset", f)
	}
}

func (ptr *QTableWidget) DisconnectReset() {
	defer qt.Recovering("disconnect QTableWidget::reset")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "reset")
	}
}

//export callbackQTableWidgetReset
func callbackQTableWidgetReset(ptrName *C.char) bool {
	defer qt.Recovering("callback QTableWidget::reset")

	if signal := qt.GetSignal(C.GoString(ptrName), "reset"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QTableWidget) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QTableWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QTableWidget) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QTableWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQTableWidgetResizeEvent
func callbackQTableWidgetResizeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableWidget::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableWidget) ConnectRowsAboutToBeRemoved(f func(parent *core.QModelIndex, start int, end int)) {
	defer qt.Recovering("connect QTableWidget::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "rowsAboutToBeRemoved", f)
	}
}

func (ptr *QTableWidget) DisconnectRowsAboutToBeRemoved() {
	defer qt.Recovering("disconnect QTableWidget::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "rowsAboutToBeRemoved")
	}
}

//export callbackQTableWidgetRowsAboutToBeRemoved
func callbackQTableWidgetRowsAboutToBeRemoved(ptrName *C.char, parent unsafe.Pointer, start C.int, end C.int) bool {
	defer qt.Recovering("callback QTableWidget::rowsAboutToBeRemoved")

	if signal := qt.GetSignal(C.GoString(ptrName), "rowsAboutToBeRemoved"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(start), int(end))
		return true
	}
	return false

}

func (ptr *QTableWidget) ConnectRowsInserted(f func(parent *core.QModelIndex, start int, end int)) {
	defer qt.Recovering("connect QTableWidget::rowsInserted")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "rowsInserted", f)
	}
}

func (ptr *QTableWidget) DisconnectRowsInserted() {
	defer qt.Recovering("disconnect QTableWidget::rowsInserted")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "rowsInserted")
	}
}

//export callbackQTableWidgetRowsInserted
func callbackQTableWidgetRowsInserted(ptrName *C.char, parent unsafe.Pointer, start C.int, end C.int) bool {
	defer qt.Recovering("callback QTableWidget::rowsInserted")

	if signal := qt.GetSignal(C.GoString(ptrName), "rowsInserted"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(start), int(end))
		return true
	}
	return false

}

func (ptr *QTableWidget) ConnectSelectAll(f func()) {
	defer qt.Recovering("connect QTableWidget::selectAll")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "selectAll", f)
	}
}

func (ptr *QTableWidget) DisconnectSelectAll() {
	defer qt.Recovering("disconnect QTableWidget::selectAll")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "selectAll")
	}
}

//export callbackQTableWidgetSelectAll
func callbackQTableWidgetSelectAll(ptrName *C.char) bool {
	defer qt.Recovering("callback QTableWidget::selectAll")

	if signal := qt.GetSignal(C.GoString(ptrName), "selectAll"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QTableWidget) ConnectStartDrag(f func(supportedActions core.Qt__DropAction)) {
	defer qt.Recovering("connect QTableWidget::startDrag")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "startDrag", f)
	}
}

func (ptr *QTableWidget) DisconnectStartDrag() {
	defer qt.Recovering("disconnect QTableWidget::startDrag")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "startDrag")
	}
}

//export callbackQTableWidgetStartDrag
func callbackQTableWidgetStartDrag(ptrName *C.char, supportedActions C.int) bool {
	defer qt.Recovering("callback QTableWidget::startDrag")

	if signal := qt.GetSignal(C.GoString(ptrName), "startDrag"); signal != nil {
		signal.(func(core.Qt__DropAction))(core.Qt__DropAction(supportedActions))
		return true
	}
	return false

}

func (ptr *QTableWidget) ConnectContextMenuEvent(f func(e *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QTableWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QTableWidget) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QTableWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQTableWidgetContextMenuEvent
func callbackQTableWidgetContextMenuEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableWidget::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QTableWidget) ConnectScrollContentsBy(f func(dx int, dy int)) {
	defer qt.Recovering("connect QTableWidget::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "scrollContentsBy", f)
	}
}

func (ptr *QTableWidget) DisconnectScrollContentsBy() {
	defer qt.Recovering("disconnect QTableWidget::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "scrollContentsBy")
	}
}

//export callbackQTableWidgetScrollContentsBy
func callbackQTableWidgetScrollContentsBy(ptrName *C.char, dx C.int, dy C.int) bool {
	defer qt.Recovering("callback QTableWidget::scrollContentsBy")

	if signal := qt.GetSignal(C.GoString(ptrName), "scrollContentsBy"); signal != nil {
		signal.(func(int, int))(int(dx), int(dy))
		return true
	}
	return false

}

func (ptr *QTableWidget) ConnectSetupViewport(f func(viewport *QWidget)) {
	defer qt.Recovering("connect QTableWidget::setupViewport")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setupViewport", f)
	}
}

func (ptr *QTableWidget) DisconnectSetupViewport() {
	defer qt.Recovering("disconnect QTableWidget::setupViewport")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setupViewport")
	}
}

//export callbackQTableWidgetSetupViewport
func callbackQTableWidgetSetupViewport(ptrName *C.char, viewport unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableWidget::setupViewport")

	if signal := qt.GetSignal(C.GoString(ptrName), "setupViewport"); signal != nil {
		signal.(func(*QWidget))(NewQWidgetFromPointer(viewport))
		return true
	}
	return false

}

func (ptr *QTableWidget) ConnectWheelEvent(f func(e *gui.QWheelEvent)) {
	defer qt.Recovering("connect QTableWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QTableWidget) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QTableWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQTableWidgetWheelEvent
func callbackQTableWidgetWheelEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableWidget::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QTableWidget) ConnectChangeEvent(f func(ev *core.QEvent)) {
	defer qt.Recovering("connect QTableWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QTableWidget) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QTableWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQTableWidgetChangeEvent
func callbackQTableWidgetChangeEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableWidget::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QTableWidget) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QTableWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QTableWidget) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QTableWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQTableWidgetActionEvent
func callbackQTableWidgetActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableWidget::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableWidget) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QTableWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QTableWidget) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QTableWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQTableWidgetEnterEvent
func callbackQTableWidgetEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableWidget::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableWidget) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QTableWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QTableWidget) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QTableWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQTableWidgetHideEvent
func callbackQTableWidgetHideEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableWidget::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableWidget) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QTableWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QTableWidget) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QTableWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQTableWidgetLeaveEvent
func callbackQTableWidgetLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableWidget::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableWidget) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QTableWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QTableWidget) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QTableWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQTableWidgetMoveEvent
func callbackQTableWidgetMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableWidget::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableWidget) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QTableWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QTableWidget) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QTableWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQTableWidgetSetVisible
func callbackQTableWidgetSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QTableWidget::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QTableWidget) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QTableWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QTableWidget) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QTableWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQTableWidgetShowEvent
func callbackQTableWidgetShowEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableWidget::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableWidget) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QTableWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QTableWidget) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QTableWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQTableWidgetCloseEvent
func callbackQTableWidgetCloseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableWidget::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableWidget) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QTableWidget::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QTableWidget) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QTableWidget::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQTableWidgetInitPainter
func callbackQTableWidgetInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableWidget::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

}

func (ptr *QTableWidget) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QTableWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QTableWidget) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QTableWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQTableWidgetKeyReleaseEvent
func callbackQTableWidgetKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableWidget::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableWidget) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QTableWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QTableWidget) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QTableWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQTableWidgetTabletEvent
func callbackQTableWidgetTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableWidget::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableWidget) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QTableWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QTableWidget) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QTableWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQTableWidgetChildEvent
func callbackQTableWidgetChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableWidget::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableWidget) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QTableWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QTableWidget) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QTableWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQTableWidgetCustomEvent
func callbackQTableWidgetCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableWidget::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
