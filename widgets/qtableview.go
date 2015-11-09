package widgets

//#include "qtableview.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
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
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QTableView_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTableView) QTableView_PTR() *QTableView {
	return ptr
}

func (ptr *QTableView) GridStyle() core.Qt__PenStyle {
	if ptr.Pointer() != nil {
		return core.Qt__PenStyle(C.QTableView_GridStyle(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTableView) IsCornerButtonEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QTableView_IsCornerButtonEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTableView) IsSortingEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QTableView_IsSortingEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTableView) SetCornerButtonEnabled(enable bool) {
	if ptr.Pointer() != nil {
		C.QTableView_SetCornerButtonEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QTableView) SetGridStyle(style core.Qt__PenStyle) {
	if ptr.Pointer() != nil {
		C.QTableView_SetGridStyle(ptr.Pointer(), C.int(style))
	}
}

func (ptr *QTableView) SetShowGrid(show bool) {
	if ptr.Pointer() != nil {
		C.QTableView_SetShowGrid(ptr.Pointer(), C.int(qt.GoBoolToInt(show)))
	}
}

func (ptr *QTableView) SetSpan(row int, column int, rowSpanCount int, columnSpanCount int) {
	if ptr.Pointer() != nil {
		C.QTableView_SetSpan(ptr.Pointer(), C.int(row), C.int(column), C.int(rowSpanCount), C.int(columnSpanCount))
	}
}

func (ptr *QTableView) SetWordWrap(on bool) {
	if ptr.Pointer() != nil {
		C.QTableView_SetWordWrap(ptr.Pointer(), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QTableView) ShowGrid() bool {
	if ptr.Pointer() != nil {
		return C.QTableView_ShowGrid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTableView) WordWrap() bool {
	if ptr.Pointer() != nil {
		return C.QTableView_WordWrap(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTableView) ClearSpans() {
	if ptr.Pointer() != nil {
		C.QTableView_ClearSpans(ptr.Pointer())
	}
}

func (ptr *QTableView) ColumnAt(x int) int {
	if ptr.Pointer() != nil {
		return int(C.QTableView_ColumnAt(ptr.Pointer(), C.int(x)))
	}
	return 0
}

func (ptr *QTableView) ColumnSpan(row int, column int) int {
	if ptr.Pointer() != nil {
		return int(C.QTableView_ColumnSpan(ptr.Pointer(), C.int(row), C.int(column)))
	}
	return 0
}

func (ptr *QTableView) ColumnViewportPosition(column int) int {
	if ptr.Pointer() != nil {
		return int(C.QTableView_ColumnViewportPosition(ptr.Pointer(), C.int(column)))
	}
	return 0
}

func (ptr *QTableView) ColumnWidth(column int) int {
	if ptr.Pointer() != nil {
		return int(C.QTableView_ColumnWidth(ptr.Pointer(), C.int(column)))
	}
	return 0
}

func (ptr *QTableView) HideColumn(column int) {
	if ptr.Pointer() != nil {
		C.QTableView_HideColumn(ptr.Pointer(), C.int(column))
	}
}

func (ptr *QTableView) HideRow(row int) {
	if ptr.Pointer() != nil {
		C.QTableView_HideRow(ptr.Pointer(), C.int(row))
	}
}

func (ptr *QTableView) HorizontalHeader() *QHeaderView {
	if ptr.Pointer() != nil {
		return NewQHeaderViewFromPointer(C.QTableView_HorizontalHeader(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTableView) IndexAt(pos core.QPoint_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QTableView_IndexAt(ptr.Pointer(), core.PointerFromQPoint(pos)))
	}
	return nil
}

func (ptr *QTableView) IsColumnHidden(column int) bool {
	if ptr.Pointer() != nil {
		return C.QTableView_IsColumnHidden(ptr.Pointer(), C.int(column)) != 0
	}
	return false
}

func (ptr *QTableView) IsRowHidden(row int) bool {
	if ptr.Pointer() != nil {
		return C.QTableView_IsRowHidden(ptr.Pointer(), C.int(row)) != 0
	}
	return false
}

func (ptr *QTableView) ResizeColumnToContents(column int) {
	if ptr.Pointer() != nil {
		C.QTableView_ResizeColumnToContents(ptr.Pointer(), C.int(column))
	}
}

func (ptr *QTableView) ResizeColumnsToContents() {
	if ptr.Pointer() != nil {
		C.QTableView_ResizeColumnsToContents(ptr.Pointer())
	}
}

func (ptr *QTableView) ResizeRowToContents(row int) {
	if ptr.Pointer() != nil {
		C.QTableView_ResizeRowToContents(ptr.Pointer(), C.int(row))
	}
}

func (ptr *QTableView) ResizeRowsToContents() {
	if ptr.Pointer() != nil {
		C.QTableView_ResizeRowsToContents(ptr.Pointer())
	}
}

func (ptr *QTableView) RowAt(y int) int {
	if ptr.Pointer() != nil {
		return int(C.QTableView_RowAt(ptr.Pointer(), C.int(y)))
	}
	return 0
}

func (ptr *QTableView) RowHeight(row int) int {
	if ptr.Pointer() != nil {
		return int(C.QTableView_RowHeight(ptr.Pointer(), C.int(row)))
	}
	return 0
}

func (ptr *QTableView) RowSpan(row int, column int) int {
	if ptr.Pointer() != nil {
		return int(C.QTableView_RowSpan(ptr.Pointer(), C.int(row), C.int(column)))
	}
	return 0
}

func (ptr *QTableView) RowViewportPosition(row int) int {
	if ptr.Pointer() != nil {
		return int(C.QTableView_RowViewportPosition(ptr.Pointer(), C.int(row)))
	}
	return 0
}

func (ptr *QTableView) SelectColumn(column int) {
	if ptr.Pointer() != nil {
		C.QTableView_SelectColumn(ptr.Pointer(), C.int(column))
	}
}

func (ptr *QTableView) SelectRow(row int) {
	if ptr.Pointer() != nil {
		C.QTableView_SelectRow(ptr.Pointer(), C.int(row))
	}
}

func (ptr *QTableView) SetColumnHidden(column int, hide bool) {
	if ptr.Pointer() != nil {
		C.QTableView_SetColumnHidden(ptr.Pointer(), C.int(column), C.int(qt.GoBoolToInt(hide)))
	}
}

func (ptr *QTableView) SetColumnWidth(column int, width int) {
	if ptr.Pointer() != nil {
		C.QTableView_SetColumnWidth(ptr.Pointer(), C.int(column), C.int(width))
	}
}

func (ptr *QTableView) SetHorizontalHeader(header QHeaderView_ITF) {
	if ptr.Pointer() != nil {
		C.QTableView_SetHorizontalHeader(ptr.Pointer(), PointerFromQHeaderView(header))
	}
}

func (ptr *QTableView) SetModel(model core.QAbstractItemModel_ITF) {
	if ptr.Pointer() != nil {
		C.QTableView_SetModel(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QTableView) SetRootIndex(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QTableView_SetRootIndex(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QTableView) SetRowHeight(row int, height int) {
	if ptr.Pointer() != nil {
		C.QTableView_SetRowHeight(ptr.Pointer(), C.int(row), C.int(height))
	}
}

func (ptr *QTableView) SetRowHidden(row int, hide bool) {
	if ptr.Pointer() != nil {
		C.QTableView_SetRowHidden(ptr.Pointer(), C.int(row), C.int(qt.GoBoolToInt(hide)))
	}
}

func (ptr *QTableView) SetSelectionModel(selectionModel core.QItemSelectionModel_ITF) {
	if ptr.Pointer() != nil {
		C.QTableView_SetSelectionModel(ptr.Pointer(), core.PointerFromQItemSelectionModel(selectionModel))
	}
}

func (ptr *QTableView) SetSortingEnabled(enable bool) {
	if ptr.Pointer() != nil {
		C.QTableView_SetSortingEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QTableView) SetVerticalHeader(header QHeaderView_ITF) {
	if ptr.Pointer() != nil {
		C.QTableView_SetVerticalHeader(ptr.Pointer(), PointerFromQHeaderView(header))
	}
}

func (ptr *QTableView) ShowColumn(column int) {
	if ptr.Pointer() != nil {
		C.QTableView_ShowColumn(ptr.Pointer(), C.int(column))
	}
}

func (ptr *QTableView) ShowRow(row int) {
	if ptr.Pointer() != nil {
		C.QTableView_ShowRow(ptr.Pointer(), C.int(row))
	}
}

func (ptr *QTableView) SortByColumn(column int, order core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.QTableView_SortByColumn(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QTableView) VerticalHeader() *QHeaderView {
	if ptr.Pointer() != nil {
		return NewQHeaderViewFromPointer(C.QTableView_VerticalHeader(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTableView) DestroyQTableView() {
	if ptr.Pointer() != nil {
		C.QTableView_DestroyQTableView(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
