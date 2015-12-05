package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
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
		n.SetObjectName("QTableView_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTableView) QTableView_PTR() *QTableView {
	return ptr
}

func (ptr *QTableView) GridStyle() core.Qt__PenStyle {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableView::gridStyle")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__PenStyle(C.QTableView_GridStyle(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTableView) IsCornerButtonEnabled() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableView::isCornerButtonEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTableView_IsCornerButtonEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTableView) IsSortingEnabled() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableView::isSortingEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTableView_IsSortingEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTableView) SetCornerButtonEnabled(enable bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableView::setCornerButtonEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableView_SetCornerButtonEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QTableView) SetGridStyle(style core.Qt__PenStyle) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableView::setGridStyle")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableView_SetGridStyle(ptr.Pointer(), C.int(style))
	}
}

func (ptr *QTableView) SetShowGrid(show bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableView::setShowGrid")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableView_SetShowGrid(ptr.Pointer(), C.int(qt.GoBoolToInt(show)))
	}
}

func (ptr *QTableView) SetSpan(row int, column int, rowSpanCount int, columnSpanCount int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableView::setSpan")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableView_SetSpan(ptr.Pointer(), C.int(row), C.int(column), C.int(rowSpanCount), C.int(columnSpanCount))
	}
}

func (ptr *QTableView) SetWordWrap(on bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableView::setWordWrap")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableView_SetWordWrap(ptr.Pointer(), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QTableView) ShowGrid() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableView::showGrid")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTableView_ShowGrid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTableView) WordWrap() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableView::wordWrap")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTableView_WordWrap(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTableView) ClearSpans() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableView::clearSpans")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableView_ClearSpans(ptr.Pointer())
	}
}

func (ptr *QTableView) ColumnAt(x int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableView::columnAt")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTableView_ColumnAt(ptr.Pointer(), C.int(x)))
	}
	return 0
}

func (ptr *QTableView) ColumnSpan(row int, column int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableView::columnSpan")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTableView_ColumnSpan(ptr.Pointer(), C.int(row), C.int(column)))
	}
	return 0
}

func (ptr *QTableView) ColumnViewportPosition(column int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableView::columnViewportPosition")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTableView_ColumnViewportPosition(ptr.Pointer(), C.int(column)))
	}
	return 0
}

func (ptr *QTableView) ColumnWidth(column int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableView::columnWidth")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTableView_ColumnWidth(ptr.Pointer(), C.int(column)))
	}
	return 0
}

func (ptr *QTableView) HideColumn(column int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableView::hideColumn")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableView_HideColumn(ptr.Pointer(), C.int(column))
	}
}

func (ptr *QTableView) HideRow(row int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableView::hideRow")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableView_HideRow(ptr.Pointer(), C.int(row))
	}
}

func (ptr *QTableView) HorizontalHeader() *QHeaderView {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableView::horizontalHeader")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQHeaderViewFromPointer(C.QTableView_HorizontalHeader(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTableView) IndexAt(pos core.QPoint_ITF) *core.QModelIndex {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableView::indexAt")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QTableView_IndexAt(ptr.Pointer(), core.PointerFromQPoint(pos)))
	}
	return nil
}

func (ptr *QTableView) IsColumnHidden(column int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableView::isColumnHidden")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTableView_IsColumnHidden(ptr.Pointer(), C.int(column)) != 0
	}
	return false
}

func (ptr *QTableView) IsRowHidden(row int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableView::isRowHidden")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTableView_IsRowHidden(ptr.Pointer(), C.int(row)) != 0
	}
	return false
}

func (ptr *QTableView) ResizeColumnToContents(column int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableView::resizeColumnToContents")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableView_ResizeColumnToContents(ptr.Pointer(), C.int(column))
	}
}

func (ptr *QTableView) ResizeColumnsToContents() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableView::resizeColumnsToContents")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableView_ResizeColumnsToContents(ptr.Pointer())
	}
}

func (ptr *QTableView) ResizeRowToContents(row int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableView::resizeRowToContents")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableView_ResizeRowToContents(ptr.Pointer(), C.int(row))
	}
}

func (ptr *QTableView) ResizeRowsToContents() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableView::resizeRowsToContents")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableView_ResizeRowsToContents(ptr.Pointer())
	}
}

func (ptr *QTableView) RowAt(y int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableView::rowAt")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTableView_RowAt(ptr.Pointer(), C.int(y)))
	}
	return 0
}

func (ptr *QTableView) RowHeight(row int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableView::rowHeight")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTableView_RowHeight(ptr.Pointer(), C.int(row)))
	}
	return 0
}

func (ptr *QTableView) RowSpan(row int, column int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableView::rowSpan")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTableView_RowSpan(ptr.Pointer(), C.int(row), C.int(column)))
	}
	return 0
}

func (ptr *QTableView) RowViewportPosition(row int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableView::rowViewportPosition")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTableView_RowViewportPosition(ptr.Pointer(), C.int(row)))
	}
	return 0
}

func (ptr *QTableView) SelectColumn(column int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableView::selectColumn")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableView_SelectColumn(ptr.Pointer(), C.int(column))
	}
}

func (ptr *QTableView) SelectRow(row int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableView::selectRow")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableView_SelectRow(ptr.Pointer(), C.int(row))
	}
}

func (ptr *QTableView) SetColumnHidden(column int, hide bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableView::setColumnHidden")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableView_SetColumnHidden(ptr.Pointer(), C.int(column), C.int(qt.GoBoolToInt(hide)))
	}
}

func (ptr *QTableView) SetColumnWidth(column int, width int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableView::setColumnWidth")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableView_SetColumnWidth(ptr.Pointer(), C.int(column), C.int(width))
	}
}

func (ptr *QTableView) SetHorizontalHeader(header QHeaderView_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableView::setHorizontalHeader")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableView_SetHorizontalHeader(ptr.Pointer(), PointerFromQHeaderView(header))
	}
}

func (ptr *QTableView) SetModel(model core.QAbstractItemModel_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableView::setModel")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableView_SetModel(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QTableView) SetRootIndex(index core.QModelIndex_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableView::setRootIndex")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableView_SetRootIndex(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QTableView) SetRowHeight(row int, height int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableView::setRowHeight")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableView_SetRowHeight(ptr.Pointer(), C.int(row), C.int(height))
	}
}

func (ptr *QTableView) SetRowHidden(row int, hide bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableView::setRowHidden")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableView_SetRowHidden(ptr.Pointer(), C.int(row), C.int(qt.GoBoolToInt(hide)))
	}
}

func (ptr *QTableView) SetSelectionModel(selectionModel core.QItemSelectionModel_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableView::setSelectionModel")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableView_SetSelectionModel(ptr.Pointer(), core.PointerFromQItemSelectionModel(selectionModel))
	}
}

func (ptr *QTableView) SetSortingEnabled(enable bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableView::setSortingEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableView_SetSortingEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QTableView) SetVerticalHeader(header QHeaderView_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableView::setVerticalHeader")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableView_SetVerticalHeader(ptr.Pointer(), PointerFromQHeaderView(header))
	}
}

func (ptr *QTableView) ShowColumn(column int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableView::showColumn")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableView_ShowColumn(ptr.Pointer(), C.int(column))
	}
}

func (ptr *QTableView) ShowRow(row int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableView::showRow")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableView_ShowRow(ptr.Pointer(), C.int(row))
	}
}

func (ptr *QTableView) SortByColumn(column int, order core.Qt__SortOrder) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableView::sortByColumn")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableView_SortByColumn(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QTableView) VerticalHeader() *QHeaderView {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableView::verticalHeader")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQHeaderViewFromPointer(C.QTableView_VerticalHeader(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTableView) DestroyQTableView() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableView::~QTableView")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableView_DestroyQTableView(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
