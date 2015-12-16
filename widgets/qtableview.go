package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QTableView struct {
	QAbstractItemView
}

type QTableView_ITF interface {
	QAbstractItemView_ITF
	QTableView_PTR() *QTableView
}

func PointerFromQTableView(ptr QTableView_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTableView_PTR().Pointer()
	}
	return nil
}

func NewQTableViewFromPointer(ptr unsafe.Pointer) *QTableView {
	var n = new(QTableView)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QTableView_") {
		n.SetObjectName("QTableView_" + qt.Identifier())
	}
	return n
}

func (ptr *QTableView) QTableView_PTR() *QTableView {
	return ptr
}

func (ptr *QTableView) GridStyle() core.Qt__PenStyle {
	defer qt.Recovering("QTableView::gridStyle")

	if ptr.Pointer() != nil {
		return core.Qt__PenStyle(C.QTableView_GridStyle(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTableView) IsCornerButtonEnabled() bool {
	defer qt.Recovering("QTableView::isCornerButtonEnabled")

	if ptr.Pointer() != nil {
		return C.QTableView_IsCornerButtonEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTableView) IsSortingEnabled() bool {
	defer qt.Recovering("QTableView::isSortingEnabled")

	if ptr.Pointer() != nil {
		return C.QTableView_IsSortingEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTableView) SetCornerButtonEnabled(enable bool) {
	defer qt.Recovering("QTableView::setCornerButtonEnabled")

	if ptr.Pointer() != nil {
		C.QTableView_SetCornerButtonEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QTableView) SetGridStyle(style core.Qt__PenStyle) {
	defer qt.Recovering("QTableView::setGridStyle")

	if ptr.Pointer() != nil {
		C.QTableView_SetGridStyle(ptr.Pointer(), C.int(style))
	}
}

func (ptr *QTableView) ConnectSetSelection(f func(rect *core.QRect, flags core.QItemSelectionModel__SelectionFlag)) {
	defer qt.Recovering("connect QTableView::setSelection")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setSelection", f)
	}
}

func (ptr *QTableView) DisconnectSetSelection() {
	defer qt.Recovering("disconnect QTableView::setSelection")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setSelection")
	}
}

//export callbackQTableViewSetSelection
func callbackQTableViewSetSelection(ptrName *C.char, rect unsafe.Pointer, flags C.int) bool {
	defer qt.Recovering("callback QTableView::setSelection")

	var signal = qt.GetSignal(C.GoString(ptrName), "setSelection")
	if signal != nil {
		defer signal.(func(*core.QRect, core.QItemSelectionModel__SelectionFlag))(core.NewQRectFromPointer(rect), core.QItemSelectionModel__SelectionFlag(flags))
		return true
	}
	return false

}

func (ptr *QTableView) SetShowGrid(show bool) {
	defer qt.Recovering("QTableView::setShowGrid")

	if ptr.Pointer() != nil {
		C.QTableView_SetShowGrid(ptr.Pointer(), C.int(qt.GoBoolToInt(show)))
	}
}

func (ptr *QTableView) SetSpan(row int, column int, rowSpanCount int, columnSpanCount int) {
	defer qt.Recovering("QTableView::setSpan")

	if ptr.Pointer() != nil {
		C.QTableView_SetSpan(ptr.Pointer(), C.int(row), C.int(column), C.int(rowSpanCount), C.int(columnSpanCount))
	}
}

func (ptr *QTableView) SetWordWrap(on bool) {
	defer qt.Recovering("QTableView::setWordWrap")

	if ptr.Pointer() != nil {
		C.QTableView_SetWordWrap(ptr.Pointer(), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QTableView) ShowGrid() bool {
	defer qt.Recovering("QTableView::showGrid")

	if ptr.Pointer() != nil {
		return C.QTableView_ShowGrid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTableView) WordWrap() bool {
	defer qt.Recovering("QTableView::wordWrap")

	if ptr.Pointer() != nil {
		return C.QTableView_WordWrap(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTableView) ClearSpans() {
	defer qt.Recovering("QTableView::clearSpans")

	if ptr.Pointer() != nil {
		C.QTableView_ClearSpans(ptr.Pointer())
	}
}

func (ptr *QTableView) ColumnAt(x int) int {
	defer qt.Recovering("QTableView::columnAt")

	if ptr.Pointer() != nil {
		return int(C.QTableView_ColumnAt(ptr.Pointer(), C.int(x)))
	}
	return 0
}

func (ptr *QTableView) ColumnSpan(row int, column int) int {
	defer qt.Recovering("QTableView::columnSpan")

	if ptr.Pointer() != nil {
		return int(C.QTableView_ColumnSpan(ptr.Pointer(), C.int(row), C.int(column)))
	}
	return 0
}

func (ptr *QTableView) ColumnViewportPosition(column int) int {
	defer qt.Recovering("QTableView::columnViewportPosition")

	if ptr.Pointer() != nil {
		return int(C.QTableView_ColumnViewportPosition(ptr.Pointer(), C.int(column)))
	}
	return 0
}

func (ptr *QTableView) ColumnWidth(column int) int {
	defer qt.Recovering("QTableView::columnWidth")

	if ptr.Pointer() != nil {
		return int(C.QTableView_ColumnWidth(ptr.Pointer(), C.int(column)))
	}
	return 0
}

func (ptr *QTableView) ConnectCurrentChanged(f func(current *core.QModelIndex, previous *core.QModelIndex)) {
	defer qt.Recovering("connect QTableView::currentChanged")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "currentChanged", f)
	}
}

func (ptr *QTableView) DisconnectCurrentChanged() {
	defer qt.Recovering("disconnect QTableView::currentChanged")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "currentChanged")
	}
}

//export callbackQTableViewCurrentChanged
func callbackQTableViewCurrentChanged(ptrName *C.char, current unsafe.Pointer, previous unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableView::currentChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "currentChanged")
	if signal != nil {
		defer signal.(func(*core.QModelIndex, *core.QModelIndex))(core.NewQModelIndexFromPointer(current), core.NewQModelIndexFromPointer(previous))
		return true
	}
	return false

}

func (ptr *QTableView) HideColumn(column int) {
	defer qt.Recovering("QTableView::hideColumn")

	if ptr.Pointer() != nil {
		C.QTableView_HideColumn(ptr.Pointer(), C.int(column))
	}
}

func (ptr *QTableView) HideRow(row int) {
	defer qt.Recovering("QTableView::hideRow")

	if ptr.Pointer() != nil {
		C.QTableView_HideRow(ptr.Pointer(), C.int(row))
	}
}

func (ptr *QTableView) HorizontalHeader() *QHeaderView {
	defer qt.Recovering("QTableView::horizontalHeader")

	if ptr.Pointer() != nil {
		return NewQHeaderViewFromPointer(C.QTableView_HorizontalHeader(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTableView) IndexAt(pos core.QPoint_ITF) *core.QModelIndex {
	defer qt.Recovering("QTableView::indexAt")

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QTableView_IndexAt(ptr.Pointer(), core.PointerFromQPoint(pos)))
	}
	return nil
}

func (ptr *QTableView) IsColumnHidden(column int) bool {
	defer qt.Recovering("QTableView::isColumnHidden")

	if ptr.Pointer() != nil {
		return C.QTableView_IsColumnHidden(ptr.Pointer(), C.int(column)) != 0
	}
	return false
}

func (ptr *QTableView) IsRowHidden(row int) bool {
	defer qt.Recovering("QTableView::isRowHidden")

	if ptr.Pointer() != nil {
		return C.QTableView_IsRowHidden(ptr.Pointer(), C.int(row)) != 0
	}
	return false
}

func (ptr *QTableView) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QTableView::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QTableView) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QTableView::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQTableViewPaintEvent
func callbackQTableViewPaintEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableView::paintEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "paintEvent")
	if signal != nil {
		defer signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableView) ResizeColumnToContents(column int) {
	defer qt.Recovering("QTableView::resizeColumnToContents")

	if ptr.Pointer() != nil {
		C.QTableView_ResizeColumnToContents(ptr.Pointer(), C.int(column))
	}
}

func (ptr *QTableView) ResizeColumnsToContents() {
	defer qt.Recovering("QTableView::resizeColumnsToContents")

	if ptr.Pointer() != nil {
		C.QTableView_ResizeColumnsToContents(ptr.Pointer())
	}
}

func (ptr *QTableView) ResizeRowToContents(row int) {
	defer qt.Recovering("QTableView::resizeRowToContents")

	if ptr.Pointer() != nil {
		C.QTableView_ResizeRowToContents(ptr.Pointer(), C.int(row))
	}
}

func (ptr *QTableView) ResizeRowsToContents() {
	defer qt.Recovering("QTableView::resizeRowsToContents")

	if ptr.Pointer() != nil {
		C.QTableView_ResizeRowsToContents(ptr.Pointer())
	}
}

func (ptr *QTableView) RowAt(y int) int {
	defer qt.Recovering("QTableView::rowAt")

	if ptr.Pointer() != nil {
		return int(C.QTableView_RowAt(ptr.Pointer(), C.int(y)))
	}
	return 0
}

func (ptr *QTableView) RowHeight(row int) int {
	defer qt.Recovering("QTableView::rowHeight")

	if ptr.Pointer() != nil {
		return int(C.QTableView_RowHeight(ptr.Pointer(), C.int(row)))
	}
	return 0
}

func (ptr *QTableView) RowSpan(row int, column int) int {
	defer qt.Recovering("QTableView::rowSpan")

	if ptr.Pointer() != nil {
		return int(C.QTableView_RowSpan(ptr.Pointer(), C.int(row), C.int(column)))
	}
	return 0
}

func (ptr *QTableView) RowViewportPosition(row int) int {
	defer qt.Recovering("QTableView::rowViewportPosition")

	if ptr.Pointer() != nil {
		return int(C.QTableView_RowViewportPosition(ptr.Pointer(), C.int(row)))
	}
	return 0
}

func (ptr *QTableView) SelectColumn(column int) {
	defer qt.Recovering("QTableView::selectColumn")

	if ptr.Pointer() != nil {
		C.QTableView_SelectColumn(ptr.Pointer(), C.int(column))
	}
}

func (ptr *QTableView) SelectRow(row int) {
	defer qt.Recovering("QTableView::selectRow")

	if ptr.Pointer() != nil {
		C.QTableView_SelectRow(ptr.Pointer(), C.int(row))
	}
}

func (ptr *QTableView) SetColumnHidden(column int, hide bool) {
	defer qt.Recovering("QTableView::setColumnHidden")

	if ptr.Pointer() != nil {
		C.QTableView_SetColumnHidden(ptr.Pointer(), C.int(column), C.int(qt.GoBoolToInt(hide)))
	}
}

func (ptr *QTableView) SetColumnWidth(column int, width int) {
	defer qt.Recovering("QTableView::setColumnWidth")

	if ptr.Pointer() != nil {
		C.QTableView_SetColumnWidth(ptr.Pointer(), C.int(column), C.int(width))
	}
}

func (ptr *QTableView) SetHorizontalHeader(header QHeaderView_ITF) {
	defer qt.Recovering("QTableView::setHorizontalHeader")

	if ptr.Pointer() != nil {
		C.QTableView_SetHorizontalHeader(ptr.Pointer(), PointerFromQHeaderView(header))
	}
}

func (ptr *QTableView) ConnectSetModel(f func(model *core.QAbstractItemModel)) {
	defer qt.Recovering("connect QTableView::setModel")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setModel", f)
	}
}

func (ptr *QTableView) DisconnectSetModel() {
	defer qt.Recovering("disconnect QTableView::setModel")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setModel")
	}
}

//export callbackQTableViewSetModel
func callbackQTableViewSetModel(ptrName *C.char, model unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableView::setModel")

	var signal = qt.GetSignal(C.GoString(ptrName), "setModel")
	if signal != nil {
		defer signal.(func(*core.QAbstractItemModel))(core.NewQAbstractItemModelFromPointer(model))
		return true
	}
	return false

}

func (ptr *QTableView) ConnectSetRootIndex(f func(index *core.QModelIndex)) {
	defer qt.Recovering("connect QTableView::setRootIndex")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setRootIndex", f)
	}
}

func (ptr *QTableView) DisconnectSetRootIndex() {
	defer qt.Recovering("disconnect QTableView::setRootIndex")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setRootIndex")
	}
}

//export callbackQTableViewSetRootIndex
func callbackQTableViewSetRootIndex(ptrName *C.char, index unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableView::setRootIndex")

	var signal = qt.GetSignal(C.GoString(ptrName), "setRootIndex")
	if signal != nil {
		defer signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
		return true
	}
	return false

}

func (ptr *QTableView) SetRowHeight(row int, height int) {
	defer qt.Recovering("QTableView::setRowHeight")

	if ptr.Pointer() != nil {
		C.QTableView_SetRowHeight(ptr.Pointer(), C.int(row), C.int(height))
	}
}

func (ptr *QTableView) SetRowHidden(row int, hide bool) {
	defer qt.Recovering("QTableView::setRowHidden")

	if ptr.Pointer() != nil {
		C.QTableView_SetRowHidden(ptr.Pointer(), C.int(row), C.int(qt.GoBoolToInt(hide)))
	}
}

func (ptr *QTableView) ConnectSetSelectionModel(f func(selectionModel *core.QItemSelectionModel)) {
	defer qt.Recovering("connect QTableView::setSelectionModel")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setSelectionModel", f)
	}
}

func (ptr *QTableView) DisconnectSetSelectionModel() {
	defer qt.Recovering("disconnect QTableView::setSelectionModel")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setSelectionModel")
	}
}

//export callbackQTableViewSetSelectionModel
func callbackQTableViewSetSelectionModel(ptrName *C.char, selectionModel unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableView::setSelectionModel")

	var signal = qt.GetSignal(C.GoString(ptrName), "setSelectionModel")
	if signal != nil {
		defer signal.(func(*core.QItemSelectionModel))(core.NewQItemSelectionModelFromPointer(selectionModel))
		return true
	}
	return false

}

func (ptr *QTableView) SetSortingEnabled(enable bool) {
	defer qt.Recovering("QTableView::setSortingEnabled")

	if ptr.Pointer() != nil {
		C.QTableView_SetSortingEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QTableView) SetVerticalHeader(header QHeaderView_ITF) {
	defer qt.Recovering("QTableView::setVerticalHeader")

	if ptr.Pointer() != nil {
		C.QTableView_SetVerticalHeader(ptr.Pointer(), PointerFromQHeaderView(header))
	}
}

func (ptr *QTableView) ShowColumn(column int) {
	defer qt.Recovering("QTableView::showColumn")

	if ptr.Pointer() != nil {
		C.QTableView_ShowColumn(ptr.Pointer(), C.int(column))
	}
}

func (ptr *QTableView) ShowRow(row int) {
	defer qt.Recovering("QTableView::showRow")

	if ptr.Pointer() != nil {
		C.QTableView_ShowRow(ptr.Pointer(), C.int(row))
	}
}

func (ptr *QTableView) SortByColumn(column int, order core.Qt__SortOrder) {
	defer qt.Recovering("QTableView::sortByColumn")

	if ptr.Pointer() != nil {
		C.QTableView_SortByColumn(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QTableView) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QTableView::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QTableView) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QTableView::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQTableViewTimerEvent
func callbackQTableViewTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableView::timerEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "timerEvent")
	if signal != nil {
		defer signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableView) ConnectUpdateGeometries(f func()) {
	defer qt.Recovering("connect QTableView::updateGeometries")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updateGeometries", f)
	}
}

func (ptr *QTableView) DisconnectUpdateGeometries() {
	defer qt.Recovering("disconnect QTableView::updateGeometries")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updateGeometries")
	}
}

//export callbackQTableViewUpdateGeometries
func callbackQTableViewUpdateGeometries(ptrName *C.char) bool {
	defer qt.Recovering("callback QTableView::updateGeometries")

	var signal = qt.GetSignal(C.GoString(ptrName), "updateGeometries")
	if signal != nil {
		defer signal.(func())()
		return true
	}
	return false

}

func (ptr *QTableView) VerticalHeader() *QHeaderView {
	defer qt.Recovering("QTableView::verticalHeader")

	if ptr.Pointer() != nil {
		return NewQHeaderViewFromPointer(C.QTableView_VerticalHeader(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTableView) DestroyQTableView() {
	defer qt.Recovering("QTableView::~QTableView")

	if ptr.Pointer() != nil {
		C.QTableView_DestroyQTableView(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
