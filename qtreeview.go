package qt

//#include "qtreeview.h"
import "C"

type qtreeview struct {
	qabstractitemview
}

type QTreeView interface {
	QAbstractItemView
	AllColumnsShowFocus() bool
	AutoExpandDelay() int
	ColumnAt_Int(x int) int
	ColumnViewportPosition_Int(column int) int
	ColumnWidth_Int(column int) int
	ExpandsOnDoubleClick() bool
	Indentation() int
	IsAnimated() bool
	IsColumnHidden_Int(column int) bool
	IsHeaderHidden() bool
	IsSortingEnabled() bool
	ItemsExpandable() bool
	RootIsDecorated() bool
	SetAllColumnsShowFocus_Bool(enable bool)
	SetAnimated_Bool(enable bool)
	SetAutoExpandDelay_Int(delay int)
	SetColumnHidden_Int_Bool(column int, hide bool)
	SetColumnWidth_Int_Int(column int, width int)
	SetExpandsOnDoubleClick_Bool(enable bool)
	SetHeaderHidden_Bool(hide bool)
	SetIndentation_Int(i int)
	SetItemsExpandable_Bool(enable bool)
	SetRootIsDecorated_Bool(show bool)
	SetSortingEnabled_Bool(enable bool)
	SetTreePosition_Int(index int)
	SetUniformRowHeights_Bool(uniform bool)
	SetWordWrap_Bool(on bool)
	SortByColumn_Int_SortOrder(column int, order SortOrder)
	TreePosition() int
	UniformRowHeights() bool
	WordWrap() bool
	ConnectSlotCollapseAll()
	DisconnectSlotCollapseAll()
	SlotCollapseAll()
	ConnectSlotExpandAll()
	DisconnectSlotExpandAll()
	SlotExpandAll()
	ConnectSlotExpandToDepth()
	DisconnectSlotExpandToDepth()
	SlotExpandToDepth_Int(depth int)
	ConnectSlotHideColumn()
	DisconnectSlotHideColumn()
	SlotHideColumn_Int(column int)
	ConnectSlotResizeColumnToContents()
	DisconnectSlotResizeColumnToContents()
	SlotResizeColumnToContents_Int(column int)
	ConnectSlotShowColumn()
	DisconnectSlotShowColumn()
	SlotShowColumn_Int(column int)
}

func (p *qtreeview) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qtreeview) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQTreeView_QWidget(parent QWidget) QTreeView {
	var parentPtr C.QtObjectPtr = nil
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qtreeview = new(qtreeview)
	qtreeview.SetPointer(C.QTreeView_New_QWidget(parentPtr))
	qtreeview.SetObjectName_String("QTreeView_" + randomIdentifier())
	return qtreeview
}

func (p *qtreeview) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QTreeView_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qtreeview) AllColumnsShowFocus() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QTreeView_AllColumnsShowFocus(p.Pointer()) != 0
	}
}

func (p *qtreeview) AutoExpandDelay() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QTreeView_AutoExpandDelay(p.Pointer()))
	}
}

func (p *qtreeview) ColumnAt_Int(x int) int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QTreeView_ColumnAt_Int(p.Pointer(), C.int(x)))
	}
}

func (p *qtreeview) ColumnViewportPosition_Int(column int) int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QTreeView_ColumnViewportPosition_Int(p.Pointer(), C.int(column)))
	}
}

func (p *qtreeview) ColumnWidth_Int(column int) int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QTreeView_ColumnWidth_Int(p.Pointer(), C.int(column)))
	}
}

func (p *qtreeview) ExpandsOnDoubleClick() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QTreeView_ExpandsOnDoubleClick(p.Pointer()) != 0
	}
}

func (p *qtreeview) Indentation() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QTreeView_Indentation(p.Pointer()))
	}
}

func (p *qtreeview) IsAnimated() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QTreeView_IsAnimated(p.Pointer()) != 0
	}
}

func (p *qtreeview) IsColumnHidden_Int(column int) bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QTreeView_IsColumnHidden_Int(p.Pointer(), C.int(column)) != 0
	}
}

func (p *qtreeview) IsHeaderHidden() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QTreeView_IsHeaderHidden(p.Pointer()) != 0
	}
}

func (p *qtreeview) IsSortingEnabled() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QTreeView_IsSortingEnabled(p.Pointer()) != 0
	}
}

func (p *qtreeview) ItemsExpandable() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QTreeView_ItemsExpandable(p.Pointer()) != 0
	}
}

func (p *qtreeview) RootIsDecorated() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QTreeView_RootIsDecorated(p.Pointer()) != 0
	}
}

func (p *qtreeview) SetAllColumnsShowFocus_Bool(enable bool) {
	if p.Pointer() != nil {
		C.QTreeView_SetAllColumnsShowFocus_Bool(p.Pointer(), goBoolToCInt(enable))
	}
}

func (p *qtreeview) SetAnimated_Bool(enable bool) {
	if p.Pointer() != nil {
		C.QTreeView_SetAnimated_Bool(p.Pointer(), goBoolToCInt(enable))
	}
}

func (p *qtreeview) SetAutoExpandDelay_Int(delay int) {
	if p.Pointer() != nil {
		C.QTreeView_SetAutoExpandDelay_Int(p.Pointer(), C.int(delay))
	}
}

func (p *qtreeview) SetColumnHidden_Int_Bool(column int, hide bool) {
	if p.Pointer() != nil {
		C.QTreeView_SetColumnHidden_Int_Bool(p.Pointer(), C.int(column), goBoolToCInt(hide))
	}
}

func (p *qtreeview) SetColumnWidth_Int_Int(column int, width int) {
	if p.Pointer() != nil {
		C.QTreeView_SetColumnWidth_Int_Int(p.Pointer(), C.int(column), C.int(width))
	}
}

func (p *qtreeview) SetExpandsOnDoubleClick_Bool(enable bool) {
	if p.Pointer() != nil {
		C.QTreeView_SetExpandsOnDoubleClick_Bool(p.Pointer(), goBoolToCInt(enable))
	}
}

func (p *qtreeview) SetHeaderHidden_Bool(hide bool) {
	if p.Pointer() != nil {
		C.QTreeView_SetHeaderHidden_Bool(p.Pointer(), goBoolToCInt(hide))
	}
}

func (p *qtreeview) SetIndentation_Int(i int) {
	if p.Pointer() != nil {
		C.QTreeView_SetIndentation_Int(p.Pointer(), C.int(i))
	}
}

func (p *qtreeview) SetItemsExpandable_Bool(enable bool) {
	if p.Pointer() != nil {
		C.QTreeView_SetItemsExpandable_Bool(p.Pointer(), goBoolToCInt(enable))
	}
}

func (p *qtreeview) SetRootIsDecorated_Bool(show bool) {
	if p.Pointer() != nil {
		C.QTreeView_SetRootIsDecorated_Bool(p.Pointer(), goBoolToCInt(show))
	}
}

func (p *qtreeview) SetSortingEnabled_Bool(enable bool) {
	if p.Pointer() != nil {
		C.QTreeView_SetSortingEnabled_Bool(p.Pointer(), goBoolToCInt(enable))
	}
}

func (p *qtreeview) SetTreePosition_Int(index int) {
	if p.Pointer() != nil {
		C.QTreeView_SetTreePosition_Int(p.Pointer(), C.int(index))
	}
}

func (p *qtreeview) SetUniformRowHeights_Bool(uniform bool) {
	if p.Pointer() != nil {
		C.QTreeView_SetUniformRowHeights_Bool(p.Pointer(), goBoolToCInt(uniform))
	}
}

func (p *qtreeview) SetWordWrap_Bool(on bool) {
	if p.Pointer() != nil {
		C.QTreeView_SetWordWrap_Bool(p.Pointer(), goBoolToCInt(on))
	}
}

func (p *qtreeview) SortByColumn_Int_SortOrder(column int, order SortOrder) {
	if p.Pointer() != nil {
		C.QTreeView_SortByColumn_Int_SortOrder(p.Pointer(), C.int(column), C.int(order))
	}
}

func (p *qtreeview) TreePosition() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QTreeView_TreePosition(p.Pointer()))
	}
}

func (p *qtreeview) UniformRowHeights() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QTreeView_UniformRowHeights(p.Pointer()) != 0
	}
}

func (p *qtreeview) WordWrap() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QTreeView_WordWrap(p.Pointer()) != 0
	}
}

func (p *qtreeview) ConnectSlotCollapseAll() {
	C.QTreeView_ConnectSlotCollapseAll(p.Pointer())
}

func (p *qtreeview) DisconnectSlotCollapseAll() {
	C.QTreeView_DisconnectSlotCollapseAll(p.Pointer())
}

func (p *qtreeview) SlotCollapseAll() {
	if p.Pointer() != nil {
		C.QTreeView_CollapseAll(p.Pointer())
	}
}

func (p *qtreeview) ConnectSlotExpandAll() {
	C.QTreeView_ConnectSlotExpandAll(p.Pointer())
}

func (p *qtreeview) DisconnectSlotExpandAll() {
	C.QTreeView_DisconnectSlotExpandAll(p.Pointer())
}

func (p *qtreeview) SlotExpandAll() {
	if p.Pointer() != nil {
		C.QTreeView_ExpandAll(p.Pointer())
	}
}

func (p *qtreeview) ConnectSlotExpandToDepth() {
	C.QTreeView_ConnectSlotExpandToDepth(p.Pointer())
}

func (p *qtreeview) DisconnectSlotExpandToDepth() {
	C.QTreeView_DisconnectSlotExpandToDepth(p.Pointer())
}

func (p *qtreeview) SlotExpandToDepth_Int(depth int) {
	if p.Pointer() != nil {
		C.QTreeView_ExpandToDepth_Int(p.Pointer(), C.int(depth))
	}
}

func (p *qtreeview) ConnectSlotHideColumn() {
	C.QTreeView_ConnectSlotHideColumn(p.Pointer())
}

func (p *qtreeview) DisconnectSlotHideColumn() {
	C.QTreeView_DisconnectSlotHideColumn(p.Pointer())
}

func (p *qtreeview) SlotHideColumn_Int(column int) {
	if p.Pointer() != nil {
		C.QTreeView_HideColumn_Int(p.Pointer(), C.int(column))
	}
}

func (p *qtreeview) ConnectSlotResizeColumnToContents() {
	C.QTreeView_ConnectSlotResizeColumnToContents(p.Pointer())
}

func (p *qtreeview) DisconnectSlotResizeColumnToContents() {
	C.QTreeView_DisconnectSlotResizeColumnToContents(p.Pointer())
}

func (p *qtreeview) SlotResizeColumnToContents_Int(column int) {
	if p.Pointer() != nil {
		C.QTreeView_ResizeColumnToContents_Int(p.Pointer(), C.int(column))
	}
}

func (p *qtreeview) ConnectSlotShowColumn() {
	C.QTreeView_ConnectSlotShowColumn(p.Pointer())
}

func (p *qtreeview) DisconnectSlotShowColumn() {
	C.QTreeView_DisconnectSlotShowColumn(p.Pointer())
}

func (p *qtreeview) SlotShowColumn_Int(column int) {
	if p.Pointer() != nil {
		C.QTreeView_ShowColumn_Int(p.Pointer(), C.int(column))
	}
}
