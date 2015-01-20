package qt

//#include "qtableview.h"
import "C"

type qtableview struct {
	qabstractitemview
}

type QTableView interface {
	QAbstractItemView
	ClearSpans()
	ColumnAt(x int) int
	ColumnSpan(row int, column int) int
	ColumnViewportPosition(column int) int
	ColumnWidth(column int) int
	GridStyle() PenStyle
	IsColumnHidden(column int) bool
	IsCornerButtonEnabled() bool
	IsRowHidden(row int) bool
	IsSortingEnabled() bool
	RowAt(y int) int
	RowHeight(row int) int
	RowSpan(row int, column int) int
	RowViewportPosition(row int) int
	SetColumnHidden(column int, hide bool)
	SetColumnWidth(column int, width int)
	SetCornerButtonEnabled(enable bool)
	SetGridStyle(style PenStyle)
	SetRowHeight(row int, height int)
	SetRowHidden(row int, hide bool)
	SetSortingEnabled(enable bool)
	SetSpan(row int, column int, rowSpanCount int, columnSpanCount int)
	SetWordWrap(on bool)
	ShowGrid() bool
	SortByColumn(column int, order SortOrder)
	WordWrap() bool
	ConnectSlotHideColumn()
	DisconnectSlotHideColumn()
	SlotHideColumn(column int)
	ConnectSlotHideRow()
	DisconnectSlotHideRow()
	SlotHideRow(row int)
	ConnectSlotResizeColumnToContents()
	DisconnectSlotResizeColumnToContents()
	SlotResizeColumnToContents(column int)
	ConnectSlotResizeColumnsToContents()
	DisconnectSlotResizeColumnsToContents()
	SlotResizeColumnsToContents()
	ConnectSlotResizeRowToContents()
	DisconnectSlotResizeRowToContents()
	SlotResizeRowToContents(row int)
	ConnectSlotResizeRowsToContents()
	DisconnectSlotResizeRowsToContents()
	SlotResizeRowsToContents()
	ConnectSlotSelectColumn()
	DisconnectSlotSelectColumn()
	SlotSelectColumn(column int)
	ConnectSlotSelectRow()
	DisconnectSlotSelectRow()
	SlotSelectRow(row int)
	ConnectSlotSetShowGrid()
	DisconnectSlotSetShowGrid()
	SlotSetShowGrid(show bool)
	ConnectSlotShowColumn()
	DisconnectSlotShowColumn()
	SlotShowColumn(column int)
	ConnectSlotShowRow()
	DisconnectSlotShowRow()
	SlotShowRow(row int)
}

func (p *qtableview) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qtableview) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQTableView(parent QWidget) QTableView {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qtableview = new(qtableview)
	qtableview.SetPointer(C.QTableView_New_QWidget(parentPtr))
	qtableview.SetObjectName("QTableView_" + randomIdentifier())
	return qtableview
}

func (p *qtableview) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QTableView_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qtableview) ClearSpans() {
	if p.Pointer() != nil {
		C.QTableView_ClearSpans(p.Pointer())
	}
}

func (p *qtableview) ColumnAt(x int) int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QTableView_ColumnAt_Int(p.Pointer(), C.int(x)))
}

func (p *qtableview) ColumnSpan(row int, column int) int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QTableView_ColumnSpan_Int_Int(p.Pointer(), C.int(row), C.int(column)))
}

func (p *qtableview) ColumnViewportPosition(column int) int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QTableView_ColumnViewportPosition_Int(p.Pointer(), C.int(column)))
}

func (p *qtableview) ColumnWidth(column int) int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QTableView_ColumnWidth_Int(p.Pointer(), C.int(column)))
}

func (p *qtableview) GridStyle() PenStyle {
	if p.Pointer() == nil {
		return 0
	}
	return PenStyle(C.QTableView_GridStyle(p.Pointer()))
}

func (p *qtableview) IsColumnHidden(column int) bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QTableView_IsColumnHidden_Int(p.Pointer(), C.int(column)) != 0
}

func (p *qtableview) IsCornerButtonEnabled() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QTableView_IsCornerButtonEnabled(p.Pointer()) != 0
}

func (p *qtableview) IsRowHidden(row int) bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QTableView_IsRowHidden_Int(p.Pointer(), C.int(row)) != 0
}

func (p *qtableview) IsSortingEnabled() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QTableView_IsSortingEnabled(p.Pointer()) != 0
}

func (p *qtableview) RowAt(y int) int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QTableView_RowAt_Int(p.Pointer(), C.int(y)))
}

func (p *qtableview) RowHeight(row int) int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QTableView_RowHeight_Int(p.Pointer(), C.int(row)))
}

func (p *qtableview) RowSpan(row int, column int) int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QTableView_RowSpan_Int_Int(p.Pointer(), C.int(row), C.int(column)))
}

func (p *qtableview) RowViewportPosition(row int) int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QTableView_RowViewportPosition_Int(p.Pointer(), C.int(row)))
}

func (p *qtableview) SetColumnHidden(column int, hide bool) {
	if p.Pointer() != nil {
		C.QTableView_SetColumnHidden_Int_Bool(p.Pointer(), C.int(column), goBoolToCInt(hide))
	}
}

func (p *qtableview) SetColumnWidth(column int, width int) {
	if p.Pointer() != nil {
		C.QTableView_SetColumnWidth_Int_Int(p.Pointer(), C.int(column), C.int(width))
	}
}

func (p *qtableview) SetCornerButtonEnabled(enable bool) {
	if p.Pointer() != nil {
		C.QTableView_SetCornerButtonEnabled_Bool(p.Pointer(), goBoolToCInt(enable))
	}
}

func (p *qtableview) SetGridStyle(style PenStyle) {
	if p.Pointer() != nil {
		C.QTableView_SetGridStyle_PenStyle(p.Pointer(), C.int(style))
	}
}

func (p *qtableview) SetRowHeight(row int, height int) {
	if p.Pointer() != nil {
		C.QTableView_SetRowHeight_Int_Int(p.Pointer(), C.int(row), C.int(height))
	}
}

func (p *qtableview) SetRowHidden(row int, hide bool) {
	if p.Pointer() != nil {
		C.QTableView_SetRowHidden_Int_Bool(p.Pointer(), C.int(row), goBoolToCInt(hide))
	}
}

func (p *qtableview) SetSortingEnabled(enable bool) {
	if p.Pointer() != nil {
		C.QTableView_SetSortingEnabled_Bool(p.Pointer(), goBoolToCInt(enable))
	}
}

func (p *qtableview) SetSpan(row int, column int, rowSpanCount int, columnSpanCount int) {
	if p.Pointer() != nil {
		C.QTableView_SetSpan_Int_Int_Int_Int(p.Pointer(), C.int(row), C.int(column), C.int(rowSpanCount), C.int(columnSpanCount))
	}
}

func (p *qtableview) SetWordWrap(on bool) {
	if p.Pointer() != nil {
		C.QTableView_SetWordWrap_Bool(p.Pointer(), goBoolToCInt(on))
	}
}

func (p *qtableview) ShowGrid() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QTableView_ShowGrid(p.Pointer()) != 0
}

func (p *qtableview) SortByColumn(column int, order SortOrder) {
	if p.Pointer() != nil {
		C.QTableView_SortByColumn_Int_SortOrder(p.Pointer(), C.int(column), C.int(order))
	}
}

func (p *qtableview) WordWrap() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QTableView_WordWrap(p.Pointer()) != 0
}

func (p *qtableview) ConnectSlotHideColumn() {
	C.QTableView_ConnectSlotHideColumn(p.Pointer())
}

func (p *qtableview) DisconnectSlotHideColumn() {
	C.QTableView_DisconnectSlotHideColumn(p.Pointer())
}

func (p *qtableview) SlotHideColumn(column int) {
	if p.Pointer() != nil {
		C.QTableView_HideColumn_Int(p.Pointer(), C.int(column))
	}
}

func (p *qtableview) ConnectSlotHideRow() {
	C.QTableView_ConnectSlotHideRow(p.Pointer())
}

func (p *qtableview) DisconnectSlotHideRow() {
	C.QTableView_DisconnectSlotHideRow(p.Pointer())
}

func (p *qtableview) SlotHideRow(row int) {
	if p.Pointer() != nil {
		C.QTableView_HideRow_Int(p.Pointer(), C.int(row))
	}
}

func (p *qtableview) ConnectSlotResizeColumnToContents() {
	C.QTableView_ConnectSlotResizeColumnToContents(p.Pointer())
}

func (p *qtableview) DisconnectSlotResizeColumnToContents() {
	C.QTableView_DisconnectSlotResizeColumnToContents(p.Pointer())
}

func (p *qtableview) SlotResizeColumnToContents(column int) {
	if p.Pointer() != nil {
		C.QTableView_ResizeColumnToContents_Int(p.Pointer(), C.int(column))
	}
}

func (p *qtableview) ConnectSlotResizeColumnsToContents() {
	C.QTableView_ConnectSlotResizeColumnsToContents(p.Pointer())
}

func (p *qtableview) DisconnectSlotResizeColumnsToContents() {
	C.QTableView_DisconnectSlotResizeColumnsToContents(p.Pointer())
}

func (p *qtableview) SlotResizeColumnsToContents() {
	if p.Pointer() != nil {
		C.QTableView_ResizeColumnsToContents(p.Pointer())
	}
}

func (p *qtableview) ConnectSlotResizeRowToContents() {
	C.QTableView_ConnectSlotResizeRowToContents(p.Pointer())
}

func (p *qtableview) DisconnectSlotResizeRowToContents() {
	C.QTableView_DisconnectSlotResizeRowToContents(p.Pointer())
}

func (p *qtableview) SlotResizeRowToContents(row int) {
	if p.Pointer() != nil {
		C.QTableView_ResizeRowToContents_Int(p.Pointer(), C.int(row))
	}
}

func (p *qtableview) ConnectSlotResizeRowsToContents() {
	C.QTableView_ConnectSlotResizeRowsToContents(p.Pointer())
}

func (p *qtableview) DisconnectSlotResizeRowsToContents() {
	C.QTableView_DisconnectSlotResizeRowsToContents(p.Pointer())
}

func (p *qtableview) SlotResizeRowsToContents() {
	if p.Pointer() != nil {
		C.QTableView_ResizeRowsToContents(p.Pointer())
	}
}

func (p *qtableview) ConnectSlotSelectColumn() {
	C.QTableView_ConnectSlotSelectColumn(p.Pointer())
}

func (p *qtableview) DisconnectSlotSelectColumn() {
	C.QTableView_DisconnectSlotSelectColumn(p.Pointer())
}

func (p *qtableview) SlotSelectColumn(column int) {
	if p.Pointer() != nil {
		C.QTableView_SelectColumn_Int(p.Pointer(), C.int(column))
	}
}

func (p *qtableview) ConnectSlotSelectRow() {
	C.QTableView_ConnectSlotSelectRow(p.Pointer())
}

func (p *qtableview) DisconnectSlotSelectRow() {
	C.QTableView_DisconnectSlotSelectRow(p.Pointer())
}

func (p *qtableview) SlotSelectRow(row int) {
	if p.Pointer() != nil {
		C.QTableView_SelectRow_Int(p.Pointer(), C.int(row))
	}
}

func (p *qtableview) ConnectSlotSetShowGrid() {
	C.QTableView_ConnectSlotSetShowGrid(p.Pointer())
}

func (p *qtableview) DisconnectSlotSetShowGrid() {
	C.QTableView_DisconnectSlotSetShowGrid(p.Pointer())
}

func (p *qtableview) SlotSetShowGrid(show bool) {
	if p.Pointer() != nil {
		C.QTableView_SetShowGrid_Bool(p.Pointer(), goBoolToCInt(show))
	}
}

func (p *qtableview) ConnectSlotShowColumn() {
	C.QTableView_ConnectSlotShowColumn(p.Pointer())
}

func (p *qtableview) DisconnectSlotShowColumn() {
	C.QTableView_DisconnectSlotShowColumn(p.Pointer())
}

func (p *qtableview) SlotShowColumn(column int) {
	if p.Pointer() != nil {
		C.QTableView_ShowColumn_Int(p.Pointer(), C.int(column))
	}
}

func (p *qtableview) ConnectSlotShowRow() {
	C.QTableView_ConnectSlotShowRow(p.Pointer())
}

func (p *qtableview) DisconnectSlotShowRow() {
	C.QTableView_DisconnectSlotShowRow(p.Pointer())
}

func (p *qtableview) SlotShowRow(row int) {
	if p.Pointer() != nil {
		C.QTableView_ShowRow_Int(p.Pointer(), C.int(row))
	}
}
