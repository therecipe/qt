package qt

//#include "qtableview.h"
import "C"

type qtableview struct {
	qabstractitemview
}

type QTableView interface {
	QAbstractItemView
	ClearSpans()
	ColumnAt_Int(x int) int
	ColumnSpan_Int_Int(row int, column int) int
	ColumnViewportPosition_Int(column int) int
	ColumnWidth_Int(column int) int
	GridStyle() PenStyle
	IsColumnHidden_Int(column int) bool
	IsCornerButtonEnabled() bool
	IsRowHidden_Int(row int) bool
	IsSortingEnabled() bool
	RowAt_Int(y int) int
	RowHeight_Int(row int) int
	RowSpan_Int_Int(row int, column int) int
	RowViewportPosition_Int(row int) int
	SetColumnHidden_Int_Bool(column int, hide bool)
	SetColumnWidth_Int_Int(column int, width int)
	SetCornerButtonEnabled_Bool(enable bool)
	SetGridStyle_PenStyle(style PenStyle)
	SetRowHeight_Int_Int(row int, height int)
	SetRowHidden_Int_Bool(row int, hide bool)
	SetSortingEnabled_Bool(enable bool)
	SetSpan_Int_Int_Int_Int(row int, column int, rowSpanCount int, columnSpanCount int)
	SetWordWrap_Bool(on bool)
	ShowGrid() bool
	SortByColumn_Int_SortOrder(column int, order SortOrder)
	WordWrap() bool
	ConnectSlotHideColumn()
	DisconnectSlotHideColumn()
	SlotHideColumn_Int(column int)
	ConnectSlotHideRow()
	DisconnectSlotHideRow()
	SlotHideRow_Int(row int)
	ConnectSlotResizeColumnToContents()
	DisconnectSlotResizeColumnToContents()
	SlotResizeColumnToContents_Int(column int)
	ConnectSlotResizeColumnsToContents()
	DisconnectSlotResizeColumnsToContents()
	SlotResizeColumnsToContents()
	ConnectSlotResizeRowToContents()
	DisconnectSlotResizeRowToContents()
	SlotResizeRowToContents_Int(row int)
	ConnectSlotResizeRowsToContents()
	DisconnectSlotResizeRowsToContents()
	SlotResizeRowsToContents()
	ConnectSlotSelectColumn()
	DisconnectSlotSelectColumn()
	SlotSelectColumn_Int(column int)
	ConnectSlotSelectRow()
	DisconnectSlotSelectRow()
	SlotSelectRow_Int(row int)
	ConnectSlotSetShowGrid()
	DisconnectSlotSetShowGrid()
	SlotSetShowGrid_Bool(show bool)
	ConnectSlotShowColumn()
	DisconnectSlotShowColumn()
	SlotShowColumn_Int(column int)
	ConnectSlotShowRow()
	DisconnectSlotShowRow()
	SlotShowRow_Int(row int)
}

func (p *qtableview) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qtableview) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQTableView_QWidget(parent QWidget) QTableView {
	var parentPtr C.QtObjectPtr = nil
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qtableview = new(qtableview)
	qtableview.SetPointer(C.QTableView_New_QWidget(parentPtr))
	qtableview.SetObjectName_String("QTableView_" + randomIdentifier())
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

func (p *qtableview) ColumnAt_Int(x int) int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QTableView_ColumnAt_Int(p.Pointer(), C.int(x)))
	}
}

func (p *qtableview) ColumnSpan_Int_Int(row int, column int) int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QTableView_ColumnSpan_Int_Int(p.Pointer(), C.int(row), C.int(column)))
	}
}

func (p *qtableview) ColumnViewportPosition_Int(column int) int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QTableView_ColumnViewportPosition_Int(p.Pointer(), C.int(column)))
	}
}

func (p *qtableview) ColumnWidth_Int(column int) int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QTableView_ColumnWidth_Int(p.Pointer(), C.int(column)))
	}
}

func (p *qtableview) GridStyle() PenStyle {
	if p.Pointer() == nil {
		return 0
	} else {
		return PenStyle(C.QTableView_GridStyle(p.Pointer()))
	}
}

func (p *qtableview) IsColumnHidden_Int(column int) bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QTableView_IsColumnHidden_Int(p.Pointer(), C.int(column)) != 0
	}
}

func (p *qtableview) IsCornerButtonEnabled() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QTableView_IsCornerButtonEnabled(p.Pointer()) != 0
	}
}

func (p *qtableview) IsRowHidden_Int(row int) bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QTableView_IsRowHidden_Int(p.Pointer(), C.int(row)) != 0
	}
}

func (p *qtableview) IsSortingEnabled() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QTableView_IsSortingEnabled(p.Pointer()) != 0
	}
}

func (p *qtableview) RowAt_Int(y int) int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QTableView_RowAt_Int(p.Pointer(), C.int(y)))
	}
}

func (p *qtableview) RowHeight_Int(row int) int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QTableView_RowHeight_Int(p.Pointer(), C.int(row)))
	}
}

func (p *qtableview) RowSpan_Int_Int(row int, column int) int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QTableView_RowSpan_Int_Int(p.Pointer(), C.int(row), C.int(column)))
	}
}

func (p *qtableview) RowViewportPosition_Int(row int) int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QTableView_RowViewportPosition_Int(p.Pointer(), C.int(row)))
	}
}

func (p *qtableview) SetColumnHidden_Int_Bool(column int, hide bool) {
	if p.Pointer() != nil {
		C.QTableView_SetColumnHidden_Int_Bool(p.Pointer(), C.int(column), goBoolToCInt(hide))
	}
}

func (p *qtableview) SetColumnWidth_Int_Int(column int, width int) {
	if p.Pointer() != nil {
		C.QTableView_SetColumnWidth_Int_Int(p.Pointer(), C.int(column), C.int(width))
	}
}

func (p *qtableview) SetCornerButtonEnabled_Bool(enable bool) {
	if p.Pointer() != nil {
		C.QTableView_SetCornerButtonEnabled_Bool(p.Pointer(), goBoolToCInt(enable))
	}
}

func (p *qtableview) SetGridStyle_PenStyle(style PenStyle) {
	if p.Pointer() != nil {
		C.QTableView_SetGridStyle_PenStyle(p.Pointer(), C.int(style))
	}
}

func (p *qtableview) SetRowHeight_Int_Int(row int, height int) {
	if p.Pointer() != nil {
		C.QTableView_SetRowHeight_Int_Int(p.Pointer(), C.int(row), C.int(height))
	}
}

func (p *qtableview) SetRowHidden_Int_Bool(row int, hide bool) {
	if p.Pointer() != nil {
		C.QTableView_SetRowHidden_Int_Bool(p.Pointer(), C.int(row), goBoolToCInt(hide))
	}
}

func (p *qtableview) SetSortingEnabled_Bool(enable bool) {
	if p.Pointer() != nil {
		C.QTableView_SetSortingEnabled_Bool(p.Pointer(), goBoolToCInt(enable))
	}
}

func (p *qtableview) SetSpan_Int_Int_Int_Int(row int, column int, rowSpanCount int, columnSpanCount int) {
	if p.Pointer() != nil {
		C.QTableView_SetSpan_Int_Int_Int_Int(p.Pointer(), C.int(row), C.int(column), C.int(rowSpanCount), C.int(columnSpanCount))
	}
}

func (p *qtableview) SetWordWrap_Bool(on bool) {
	if p.Pointer() != nil {
		C.QTableView_SetWordWrap_Bool(p.Pointer(), goBoolToCInt(on))
	}
}

func (p *qtableview) ShowGrid() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QTableView_ShowGrid(p.Pointer()) != 0
	}
}

func (p *qtableview) SortByColumn_Int_SortOrder(column int, order SortOrder) {
	if p.Pointer() != nil {
		C.QTableView_SortByColumn_Int_SortOrder(p.Pointer(), C.int(column), C.int(order))
	}
}

func (p *qtableview) WordWrap() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QTableView_WordWrap(p.Pointer()) != 0
	}
}

func (p *qtableview) ConnectSlotHideColumn() {
	C.QTableView_ConnectSlotHideColumn(p.Pointer())
}

func (p *qtableview) DisconnectSlotHideColumn() {
	C.QTableView_DisconnectSlotHideColumn(p.Pointer())
}

func (p *qtableview) SlotHideColumn_Int(column int) {
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

func (p *qtableview) SlotHideRow_Int(row int) {
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

func (p *qtableview) SlotResizeColumnToContents_Int(column int) {
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

func (p *qtableview) SlotResizeRowToContents_Int(row int) {
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

func (p *qtableview) SlotSelectColumn_Int(column int) {
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

func (p *qtableview) SlotSelectRow_Int(row int) {
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

func (p *qtableview) SlotSetShowGrid_Bool(show bool) {
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

func (p *qtableview) SlotShowColumn_Int(column int) {
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

func (p *qtableview) SlotShowRow_Int(row int) {
	if p.Pointer() != nil {
		C.QTableView_ShowRow_Int(p.Pointer(), C.int(row))
	}
}
