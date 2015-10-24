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

type QTableViewITF interface {
	QAbstractItemViewITF
	QTableViewPTR() *QTableView
}

func PointerFromQTableView(ptr QTableViewITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTableViewPTR().Pointer()
	}
	return nil
}

func QTableViewFromPointer(ptr unsafe.Pointer) *QTableView {
	var n = new(QTableView)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QTableView_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTableView) QTableViewPTR() *QTableView {
	return ptr
}

func (ptr *QTableView) GridStyle() core.Qt__PenStyle {
	if ptr.Pointer() != nil {
		return core.Qt__PenStyle(C.QTableView_GridStyle(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTableView) IsCornerButtonEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QTableView_IsCornerButtonEnabled(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTableView) IsSortingEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QTableView_IsSortingEnabled(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTableView) SetCornerButtonEnabled(enable bool) {
	if ptr.Pointer() != nil {
		C.QTableView_SetCornerButtonEnabled(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QTableView) SetGridStyle(style core.Qt__PenStyle) {
	if ptr.Pointer() != nil {
		C.QTableView_SetGridStyle(C.QtObjectPtr(ptr.Pointer()), C.int(style))
	}
}

func (ptr *QTableView) SetShowGrid(show bool) {
	if ptr.Pointer() != nil {
		C.QTableView_SetShowGrid(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(show)))
	}
}

func (ptr *QTableView) SetSpan(row int, column int, rowSpanCount int, columnSpanCount int) {
	if ptr.Pointer() != nil {
		C.QTableView_SetSpan(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(column), C.int(rowSpanCount), C.int(columnSpanCount))
	}
}

func (ptr *QTableView) SetWordWrap(on bool) {
	if ptr.Pointer() != nil {
		C.QTableView_SetWordWrap(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QTableView) ShowGrid() bool {
	if ptr.Pointer() != nil {
		return C.QTableView_ShowGrid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTableView) WordWrap() bool {
	if ptr.Pointer() != nil {
		return C.QTableView_WordWrap(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTableView) ClearSpans() {
	if ptr.Pointer() != nil {
		C.QTableView_ClearSpans(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTableView) ColumnAt(x int) int {
	if ptr.Pointer() != nil {
		return int(C.QTableView_ColumnAt(C.QtObjectPtr(ptr.Pointer()), C.int(x)))
	}
	return 0
}

func (ptr *QTableView) ColumnSpan(row int, column int) int {
	if ptr.Pointer() != nil {
		return int(C.QTableView_ColumnSpan(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(column)))
	}
	return 0
}

func (ptr *QTableView) ColumnViewportPosition(column int) int {
	if ptr.Pointer() != nil {
		return int(C.QTableView_ColumnViewportPosition(C.QtObjectPtr(ptr.Pointer()), C.int(column)))
	}
	return 0
}

func (ptr *QTableView) ColumnWidth(column int) int {
	if ptr.Pointer() != nil {
		return int(C.QTableView_ColumnWidth(C.QtObjectPtr(ptr.Pointer()), C.int(column)))
	}
	return 0
}

func (ptr *QTableView) HideColumn(column int) {
	if ptr.Pointer() != nil {
		C.QTableView_HideColumn(C.QtObjectPtr(ptr.Pointer()), C.int(column))
	}
}

func (ptr *QTableView) HideRow(row int) {
	if ptr.Pointer() != nil {
		C.QTableView_HideRow(C.QtObjectPtr(ptr.Pointer()), C.int(row))
	}
}

func (ptr *QTableView) HorizontalHeader() *QHeaderView {
	if ptr.Pointer() != nil {
		return QHeaderViewFromPointer(unsafe.Pointer(C.QTableView_HorizontalHeader(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QTableView) IndexAt(pos core.QPointITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		return core.QModelIndexFromPointer(unsafe.Pointer(C.QTableView_IndexAt(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(pos)))))
	}
	return nil
}

func (ptr *QTableView) IsColumnHidden(column int) bool {
	if ptr.Pointer() != nil {
		return C.QTableView_IsColumnHidden(C.QtObjectPtr(ptr.Pointer()), C.int(column)) != 0
	}
	return false
}

func (ptr *QTableView) IsRowHidden(row int) bool {
	if ptr.Pointer() != nil {
		return C.QTableView_IsRowHidden(C.QtObjectPtr(ptr.Pointer()), C.int(row)) != 0
	}
	return false
}

func (ptr *QTableView) ResizeColumnToContents(column int) {
	if ptr.Pointer() != nil {
		C.QTableView_ResizeColumnToContents(C.QtObjectPtr(ptr.Pointer()), C.int(column))
	}
}

func (ptr *QTableView) ResizeColumnsToContents() {
	if ptr.Pointer() != nil {
		C.QTableView_ResizeColumnsToContents(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTableView) ResizeRowToContents(row int) {
	if ptr.Pointer() != nil {
		C.QTableView_ResizeRowToContents(C.QtObjectPtr(ptr.Pointer()), C.int(row))
	}
}

func (ptr *QTableView) ResizeRowsToContents() {
	if ptr.Pointer() != nil {
		C.QTableView_ResizeRowsToContents(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTableView) RowAt(y int) int {
	if ptr.Pointer() != nil {
		return int(C.QTableView_RowAt(C.QtObjectPtr(ptr.Pointer()), C.int(y)))
	}
	return 0
}

func (ptr *QTableView) RowHeight(row int) int {
	if ptr.Pointer() != nil {
		return int(C.QTableView_RowHeight(C.QtObjectPtr(ptr.Pointer()), C.int(row)))
	}
	return 0
}

func (ptr *QTableView) RowSpan(row int, column int) int {
	if ptr.Pointer() != nil {
		return int(C.QTableView_RowSpan(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(column)))
	}
	return 0
}

func (ptr *QTableView) RowViewportPosition(row int) int {
	if ptr.Pointer() != nil {
		return int(C.QTableView_RowViewportPosition(C.QtObjectPtr(ptr.Pointer()), C.int(row)))
	}
	return 0
}

func (ptr *QTableView) SelectColumn(column int) {
	if ptr.Pointer() != nil {
		C.QTableView_SelectColumn(C.QtObjectPtr(ptr.Pointer()), C.int(column))
	}
}

func (ptr *QTableView) SelectRow(row int) {
	if ptr.Pointer() != nil {
		C.QTableView_SelectRow(C.QtObjectPtr(ptr.Pointer()), C.int(row))
	}
}

func (ptr *QTableView) SetColumnHidden(column int, hide bool) {
	if ptr.Pointer() != nil {
		C.QTableView_SetColumnHidden(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.int(qt.GoBoolToInt(hide)))
	}
}

func (ptr *QTableView) SetColumnWidth(column int, width int) {
	if ptr.Pointer() != nil {
		C.QTableView_SetColumnWidth(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.int(width))
	}
}

func (ptr *QTableView) SetHorizontalHeader(header QHeaderViewITF) {
	if ptr.Pointer() != nil {
		C.QTableView_SetHorizontalHeader(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQHeaderView(header)))
	}
}

func (ptr *QTableView) SetModel(model core.QAbstractItemModelITF) {
	if ptr.Pointer() != nil {
		C.QTableView_SetModel(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQAbstractItemModel(model)))
	}
}

func (ptr *QTableView) SetRootIndex(index core.QModelIndexITF) {
	if ptr.Pointer() != nil {
		C.QTableView_SetRootIndex(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(index)))
	}
}

func (ptr *QTableView) SetRowHeight(row int, height int) {
	if ptr.Pointer() != nil {
		C.QTableView_SetRowHeight(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(height))
	}
}

func (ptr *QTableView) SetRowHidden(row int, hide bool) {
	if ptr.Pointer() != nil {
		C.QTableView_SetRowHidden(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(qt.GoBoolToInt(hide)))
	}
}

func (ptr *QTableView) SetSelectionModel(selectionModel core.QItemSelectionModelITF) {
	if ptr.Pointer() != nil {
		C.QTableView_SetSelectionModel(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQItemSelectionModel(selectionModel)))
	}
}

func (ptr *QTableView) SetSortingEnabled(enable bool) {
	if ptr.Pointer() != nil {
		C.QTableView_SetSortingEnabled(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QTableView) SetVerticalHeader(header QHeaderViewITF) {
	if ptr.Pointer() != nil {
		C.QTableView_SetVerticalHeader(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQHeaderView(header)))
	}
}

func (ptr *QTableView) ShowColumn(column int) {
	if ptr.Pointer() != nil {
		C.QTableView_ShowColumn(C.QtObjectPtr(ptr.Pointer()), C.int(column))
	}
}

func (ptr *QTableView) ShowRow(row int) {
	if ptr.Pointer() != nil {
		C.QTableView_ShowRow(C.QtObjectPtr(ptr.Pointer()), C.int(row))
	}
}

func (ptr *QTableView) SortByColumn(column int, order core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.QTableView_SortByColumn(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.int(order))
	}
}

func (ptr *QTableView) VerticalHeader() *QHeaderView {
	if ptr.Pointer() != nil {
		return QHeaderViewFromPointer(unsafe.Pointer(C.QTableView_VerticalHeader(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QTableView) DestroyQTableView() {
	if ptr.Pointer() != nil {
		C.QTableView_DestroyQTableView(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
