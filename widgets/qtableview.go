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

	if signal := qt.GetSignal(C.GoString(ptrName), "setSelection"); signal != nil {
		signal.(func(*core.QRect, core.QItemSelectionModel__SelectionFlag))(core.NewQRectFromPointer(rect), core.QItemSelectionModel__SelectionFlag(flags))
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

	if signal := qt.GetSignal(C.GoString(ptrName), "currentChanged"); signal != nil {
		signal.(func(*core.QModelIndex, *core.QModelIndex))(core.NewQModelIndexFromPointer(current), core.NewQModelIndexFromPointer(previous))
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

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
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

	if signal := qt.GetSignal(C.GoString(ptrName), "setRootIndex"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
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

	if signal := qt.GetSignal(C.GoString(ptrName), "setSelectionModel"); signal != nil {
		signal.(func(*core.QItemSelectionModel))(core.NewQItemSelectionModelFromPointer(selectionModel))
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

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
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

	if signal := qt.GetSignal(C.GoString(ptrName), "updateGeometries"); signal != nil {
		signal.(func())()
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

func (ptr *QTableView) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QTableView::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QTableView) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QTableView::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQTableViewDragLeaveEvent
func callbackQTableViewDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableView::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableView) ConnectCloseEditor(f func(editor *QWidget, hint QAbstractItemDelegate__EndEditHint)) {
	defer qt.Recovering("connect QTableView::closeEditor")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEditor", f)
	}
}

func (ptr *QTableView) DisconnectCloseEditor() {
	defer qt.Recovering("disconnect QTableView::closeEditor")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEditor")
	}
}

//export callbackQTableViewCloseEditor
func callbackQTableViewCloseEditor(ptrName *C.char, editor unsafe.Pointer, hint C.int) bool {
	defer qt.Recovering("callback QTableView::closeEditor")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEditor"); signal != nil {
		signal.(func(*QWidget, QAbstractItemDelegate__EndEditHint))(NewQWidgetFromPointer(editor), QAbstractItemDelegate__EndEditHint(hint))
		return true
	}
	return false

}

func (ptr *QTableView) ConnectCommitData(f func(editor *QWidget)) {
	defer qt.Recovering("connect QTableView::commitData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "commitData", f)
	}
}

func (ptr *QTableView) DisconnectCommitData() {
	defer qt.Recovering("disconnect QTableView::commitData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "commitData")
	}
}

//export callbackQTableViewCommitData
func callbackQTableViewCommitData(ptrName *C.char, editor unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableView::commitData")

	if signal := qt.GetSignal(C.GoString(ptrName), "commitData"); signal != nil {
		signal.(func(*QWidget))(NewQWidgetFromPointer(editor))
		return true
	}
	return false

}

func (ptr *QTableView) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QTableView::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QTableView) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QTableView::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQTableViewDragEnterEvent
func callbackQTableViewDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableView::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableView) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QTableView::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QTableView) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QTableView::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQTableViewDragMoveEvent
func callbackQTableViewDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableView::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableView) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QTableView::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QTableView) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QTableView::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQTableViewDropEvent
func callbackQTableViewDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableView::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableView) ConnectEditorDestroyed(f func(editor *core.QObject)) {
	defer qt.Recovering("connect QTableView::editorDestroyed")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "editorDestroyed", f)
	}
}

func (ptr *QTableView) DisconnectEditorDestroyed() {
	defer qt.Recovering("disconnect QTableView::editorDestroyed")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "editorDestroyed")
	}
}

//export callbackQTableViewEditorDestroyed
func callbackQTableViewEditorDestroyed(ptrName *C.char, editor unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableView::editorDestroyed")

	if signal := qt.GetSignal(C.GoString(ptrName), "editorDestroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(editor))
		return true
	}
	return false

}

func (ptr *QTableView) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QTableView::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QTableView) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QTableView::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQTableViewFocusInEvent
func callbackQTableViewFocusInEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableView::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableView) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QTableView::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QTableView) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QTableView::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQTableViewFocusOutEvent
func callbackQTableViewFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableView::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableView) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QTableView::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QTableView) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QTableView::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQTableViewInputMethodEvent
func callbackQTableViewInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableView::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableView) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QTableView::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QTableView) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QTableView::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQTableViewKeyPressEvent
func callbackQTableViewKeyPressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableView::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableView) ConnectKeyboardSearch(f func(search string)) {
	defer qt.Recovering("connect QTableView::keyboardSearch")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyboardSearch", f)
	}
}

func (ptr *QTableView) DisconnectKeyboardSearch() {
	defer qt.Recovering("disconnect QTableView::keyboardSearch")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyboardSearch")
	}
}

//export callbackQTableViewKeyboardSearch
func callbackQTableViewKeyboardSearch(ptrName *C.char, search *C.char) bool {
	defer qt.Recovering("callback QTableView::keyboardSearch")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyboardSearch"); signal != nil {
		signal.(func(string))(C.GoString(search))
		return true
	}
	return false

}

func (ptr *QTableView) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QTableView::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QTableView) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QTableView::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQTableViewMouseDoubleClickEvent
func callbackQTableViewMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableView::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableView) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QTableView::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QTableView) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QTableView::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQTableViewMouseMoveEvent
func callbackQTableViewMouseMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableView::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableView) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QTableView::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QTableView) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QTableView::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQTableViewMousePressEvent
func callbackQTableViewMousePressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableView::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableView) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QTableView::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QTableView) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QTableView::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQTableViewMouseReleaseEvent
func callbackQTableViewMouseReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableView::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableView) ConnectReset(f func()) {
	defer qt.Recovering("connect QTableView::reset")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "reset", f)
	}
}

func (ptr *QTableView) DisconnectReset() {
	defer qt.Recovering("disconnect QTableView::reset")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "reset")
	}
}

//export callbackQTableViewReset
func callbackQTableViewReset(ptrName *C.char) bool {
	defer qt.Recovering("callback QTableView::reset")

	if signal := qt.GetSignal(C.GoString(ptrName), "reset"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QTableView) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QTableView::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QTableView) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QTableView::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQTableViewResizeEvent
func callbackQTableViewResizeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableView::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableView) ConnectRowsAboutToBeRemoved(f func(parent *core.QModelIndex, start int, end int)) {
	defer qt.Recovering("connect QTableView::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "rowsAboutToBeRemoved", f)
	}
}

func (ptr *QTableView) DisconnectRowsAboutToBeRemoved() {
	defer qt.Recovering("disconnect QTableView::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "rowsAboutToBeRemoved")
	}
}

//export callbackQTableViewRowsAboutToBeRemoved
func callbackQTableViewRowsAboutToBeRemoved(ptrName *C.char, parent unsafe.Pointer, start C.int, end C.int) bool {
	defer qt.Recovering("callback QTableView::rowsAboutToBeRemoved")

	if signal := qt.GetSignal(C.GoString(ptrName), "rowsAboutToBeRemoved"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(start), int(end))
		return true
	}
	return false

}

func (ptr *QTableView) ConnectRowsInserted(f func(parent *core.QModelIndex, start int, end int)) {
	defer qt.Recovering("connect QTableView::rowsInserted")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "rowsInserted", f)
	}
}

func (ptr *QTableView) DisconnectRowsInserted() {
	defer qt.Recovering("disconnect QTableView::rowsInserted")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "rowsInserted")
	}
}

//export callbackQTableViewRowsInserted
func callbackQTableViewRowsInserted(ptrName *C.char, parent unsafe.Pointer, start C.int, end C.int) bool {
	defer qt.Recovering("callback QTableView::rowsInserted")

	if signal := qt.GetSignal(C.GoString(ptrName), "rowsInserted"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(start), int(end))
		return true
	}
	return false

}

func (ptr *QTableView) ConnectSelectAll(f func()) {
	defer qt.Recovering("connect QTableView::selectAll")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "selectAll", f)
	}
}

func (ptr *QTableView) DisconnectSelectAll() {
	defer qt.Recovering("disconnect QTableView::selectAll")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "selectAll")
	}
}

//export callbackQTableViewSelectAll
func callbackQTableViewSelectAll(ptrName *C.char) bool {
	defer qt.Recovering("callback QTableView::selectAll")

	if signal := qt.GetSignal(C.GoString(ptrName), "selectAll"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QTableView) ConnectStartDrag(f func(supportedActions core.Qt__DropAction)) {
	defer qt.Recovering("connect QTableView::startDrag")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "startDrag", f)
	}
}

func (ptr *QTableView) DisconnectStartDrag() {
	defer qt.Recovering("disconnect QTableView::startDrag")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "startDrag")
	}
}

//export callbackQTableViewStartDrag
func callbackQTableViewStartDrag(ptrName *C.char, supportedActions C.int) bool {
	defer qt.Recovering("callback QTableView::startDrag")

	if signal := qt.GetSignal(C.GoString(ptrName), "startDrag"); signal != nil {
		signal.(func(core.Qt__DropAction))(core.Qt__DropAction(supportedActions))
		return true
	}
	return false

}

func (ptr *QTableView) ConnectContextMenuEvent(f func(e *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QTableView::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QTableView) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QTableView::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQTableViewContextMenuEvent
func callbackQTableViewContextMenuEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableView::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QTableView) ConnectScrollContentsBy(f func(dx int, dy int)) {
	defer qt.Recovering("connect QTableView::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "scrollContentsBy", f)
	}
}

func (ptr *QTableView) DisconnectScrollContentsBy() {
	defer qt.Recovering("disconnect QTableView::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "scrollContentsBy")
	}
}

//export callbackQTableViewScrollContentsBy
func callbackQTableViewScrollContentsBy(ptrName *C.char, dx C.int, dy C.int) bool {
	defer qt.Recovering("callback QTableView::scrollContentsBy")

	if signal := qt.GetSignal(C.GoString(ptrName), "scrollContentsBy"); signal != nil {
		signal.(func(int, int))(int(dx), int(dy))
		return true
	}
	return false

}

func (ptr *QTableView) ConnectSetupViewport(f func(viewport *QWidget)) {
	defer qt.Recovering("connect QTableView::setupViewport")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setupViewport", f)
	}
}

func (ptr *QTableView) DisconnectSetupViewport() {
	defer qt.Recovering("disconnect QTableView::setupViewport")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setupViewport")
	}
}

//export callbackQTableViewSetupViewport
func callbackQTableViewSetupViewport(ptrName *C.char, viewport unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableView::setupViewport")

	if signal := qt.GetSignal(C.GoString(ptrName), "setupViewport"); signal != nil {
		signal.(func(*QWidget))(NewQWidgetFromPointer(viewport))
		return true
	}
	return false

}

func (ptr *QTableView) ConnectWheelEvent(f func(e *gui.QWheelEvent)) {
	defer qt.Recovering("connect QTableView::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QTableView) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QTableView::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQTableViewWheelEvent
func callbackQTableViewWheelEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableView::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QTableView) ConnectChangeEvent(f func(ev *core.QEvent)) {
	defer qt.Recovering("connect QTableView::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QTableView) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QTableView::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQTableViewChangeEvent
func callbackQTableViewChangeEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableView::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QTableView) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QTableView::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QTableView) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QTableView::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQTableViewActionEvent
func callbackQTableViewActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableView::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableView) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QTableView::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QTableView) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QTableView::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQTableViewEnterEvent
func callbackQTableViewEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableView::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableView) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QTableView::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QTableView) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QTableView::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQTableViewHideEvent
func callbackQTableViewHideEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableView::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableView) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QTableView::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QTableView) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QTableView::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQTableViewLeaveEvent
func callbackQTableViewLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableView::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableView) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QTableView::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QTableView) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QTableView::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQTableViewMoveEvent
func callbackQTableViewMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableView::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableView) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QTableView::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QTableView) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QTableView::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQTableViewSetVisible
func callbackQTableViewSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QTableView::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QTableView) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QTableView::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QTableView) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QTableView::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQTableViewShowEvent
func callbackQTableViewShowEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableView::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableView) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QTableView::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QTableView) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QTableView::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQTableViewCloseEvent
func callbackQTableViewCloseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableView::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableView) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QTableView::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QTableView) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QTableView::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQTableViewInitPainter
func callbackQTableViewInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableView::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

}

func (ptr *QTableView) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QTableView::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QTableView) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QTableView::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQTableViewKeyReleaseEvent
func callbackQTableViewKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableView::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableView) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QTableView::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QTableView) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QTableView::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQTableViewTabletEvent
func callbackQTableViewTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableView::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableView) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QTableView::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QTableView) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QTableView::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQTableViewChildEvent
func callbackQTableViewChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableView::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTableView) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QTableView::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QTableView) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QTableView::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQTableViewCustomEvent
func callbackQTableViewCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTableView::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
