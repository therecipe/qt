package widgets

//#include "qtreeview.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QTreeView struct {
	QAbstractItemView
}

type QTreeView_ITF interface {
	QAbstractItemView_ITF
	QTreeView_PTR() *QTreeView
}

func PointerFromQTreeView(ptr QTreeView_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTreeView_PTR().Pointer()
	}
	return nil
}

func NewQTreeViewFromPointer(ptr unsafe.Pointer) *QTreeView {
	var n = new(QTreeView)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QTreeView_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTreeView) QTreeView_PTR() *QTreeView {
	return ptr
}

func (ptr *QTreeView) AllColumnsShowFocus() bool {
	if ptr.Pointer() != nil {
		return C.QTreeView_AllColumnsShowFocus(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTreeView) AutoExpandDelay() int {
	if ptr.Pointer() != nil {
		return int(C.QTreeView_AutoExpandDelay(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTreeView) Collapse(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QTreeView_Collapse(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QTreeView) Expand(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QTreeView_Expand(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QTreeView) ExpandsOnDoubleClick() bool {
	if ptr.Pointer() != nil {
		return C.QTreeView_ExpandsOnDoubleClick(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTreeView) Indentation() int {
	if ptr.Pointer() != nil {
		return int(C.QTreeView_Indentation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTreeView) IsAnimated() bool {
	if ptr.Pointer() != nil {
		return C.QTreeView_IsAnimated(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTreeView) IsExpanded(index core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QTreeView_IsExpanded(ptr.Pointer(), core.PointerFromQModelIndex(index)) != 0
	}
	return false
}

func (ptr *QTreeView) IsHeaderHidden() bool {
	if ptr.Pointer() != nil {
		return C.QTreeView_IsHeaderHidden(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTreeView) IsSortingEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QTreeView_IsSortingEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTreeView) ItemsExpandable() bool {
	if ptr.Pointer() != nil {
		return C.QTreeView_ItemsExpandable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTreeView) ResetIndentation() {
	if ptr.Pointer() != nil {
		C.QTreeView_ResetIndentation(ptr.Pointer())
	}
}

func (ptr *QTreeView) RootIsDecorated() bool {
	if ptr.Pointer() != nil {
		return C.QTreeView_RootIsDecorated(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTreeView) SetAllColumnsShowFocus(enable bool) {
	if ptr.Pointer() != nil {
		C.QTreeView_SetAllColumnsShowFocus(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QTreeView) SetAnimated(enable bool) {
	if ptr.Pointer() != nil {
		C.QTreeView_SetAnimated(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QTreeView) SetAutoExpandDelay(delay int) {
	if ptr.Pointer() != nil {
		C.QTreeView_SetAutoExpandDelay(ptr.Pointer(), C.int(delay))
	}
}

func (ptr *QTreeView) SetExpandsOnDoubleClick(enable bool) {
	if ptr.Pointer() != nil {
		C.QTreeView_SetExpandsOnDoubleClick(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QTreeView) SetHeaderHidden(hide bool) {
	if ptr.Pointer() != nil {
		C.QTreeView_SetHeaderHidden(ptr.Pointer(), C.int(qt.GoBoolToInt(hide)))
	}
}

func (ptr *QTreeView) SetIndentation(i int) {
	if ptr.Pointer() != nil {
		C.QTreeView_SetIndentation(ptr.Pointer(), C.int(i))
	}
}

func (ptr *QTreeView) SetItemsExpandable(enable bool) {
	if ptr.Pointer() != nil {
		C.QTreeView_SetItemsExpandable(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QTreeView) SetRootIsDecorated(show bool) {
	if ptr.Pointer() != nil {
		C.QTreeView_SetRootIsDecorated(ptr.Pointer(), C.int(qt.GoBoolToInt(show)))
	}
}

func (ptr *QTreeView) SetSortingEnabled(enable bool) {
	if ptr.Pointer() != nil {
		C.QTreeView_SetSortingEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QTreeView) SetUniformRowHeights(uniform bool) {
	if ptr.Pointer() != nil {
		C.QTreeView_SetUniformRowHeights(ptr.Pointer(), C.int(qt.GoBoolToInt(uniform)))
	}
}

func (ptr *QTreeView) SetWordWrap(on bool) {
	if ptr.Pointer() != nil {
		C.QTreeView_SetWordWrap(ptr.Pointer(), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QTreeView) UniformRowHeights() bool {
	if ptr.Pointer() != nil {
		return C.QTreeView_UniformRowHeights(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTreeView) WordWrap() bool {
	if ptr.Pointer() != nil {
		return C.QTreeView_WordWrap(ptr.Pointer()) != 0
	}
	return false
}

func NewQTreeView(parent QWidget_ITF) *QTreeView {
	return NewQTreeViewFromPointer(C.QTreeView_NewQTreeView(PointerFromQWidget(parent)))
}

func (ptr *QTreeView) CollapseAll() {
	if ptr.Pointer() != nil {
		C.QTreeView_CollapseAll(ptr.Pointer())
	}
}

func (ptr *QTreeView) ConnectCollapsed(f func(index *core.QModelIndex)) {
	if ptr.Pointer() != nil {
		C.QTreeView_ConnectCollapsed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "collapsed", f)
	}
}

func (ptr *QTreeView) DisconnectCollapsed() {
	if ptr.Pointer() != nil {
		C.QTreeView_DisconnectCollapsed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "collapsed")
	}
}

//export callbackQTreeViewCollapsed
func callbackQTreeViewCollapsed(ptrName *C.char, index unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "collapsed").(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
}

func (ptr *QTreeView) ColumnAt(x int) int {
	if ptr.Pointer() != nil {
		return int(C.QTreeView_ColumnAt(ptr.Pointer(), C.int(x)))
	}
	return 0
}

func (ptr *QTreeView) ColumnViewportPosition(column int) int {
	if ptr.Pointer() != nil {
		return int(C.QTreeView_ColumnViewportPosition(ptr.Pointer(), C.int(column)))
	}
	return 0
}

func (ptr *QTreeView) ColumnWidth(column int) int {
	if ptr.Pointer() != nil {
		return int(C.QTreeView_ColumnWidth(ptr.Pointer(), C.int(column)))
	}
	return 0
}

func (ptr *QTreeView) ExpandAll() {
	if ptr.Pointer() != nil {
		C.QTreeView_ExpandAll(ptr.Pointer())
	}
}

func (ptr *QTreeView) ExpandToDepth(depth int) {
	if ptr.Pointer() != nil {
		C.QTreeView_ExpandToDepth(ptr.Pointer(), C.int(depth))
	}
}

func (ptr *QTreeView) ConnectExpanded(f func(index *core.QModelIndex)) {
	if ptr.Pointer() != nil {
		C.QTreeView_ConnectExpanded(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "expanded", f)
	}
}

func (ptr *QTreeView) DisconnectExpanded() {
	if ptr.Pointer() != nil {
		C.QTreeView_DisconnectExpanded(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "expanded")
	}
}

//export callbackQTreeViewExpanded
func callbackQTreeViewExpanded(ptrName *C.char, index unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "expanded").(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
}

func (ptr *QTreeView) Header() *QHeaderView {
	if ptr.Pointer() != nil {
		return NewQHeaderViewFromPointer(C.QTreeView_Header(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTreeView) HideColumn(column int) {
	if ptr.Pointer() != nil {
		C.QTreeView_HideColumn(ptr.Pointer(), C.int(column))
	}
}

func (ptr *QTreeView) IndexAbove(index core.QModelIndex_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QTreeView_IndexAbove(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return nil
}

func (ptr *QTreeView) IndexAt(point core.QPoint_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QTreeView_IndexAt(ptr.Pointer(), core.PointerFromQPoint(point)))
	}
	return nil
}

func (ptr *QTreeView) IndexBelow(index core.QModelIndex_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QTreeView_IndexBelow(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return nil
}

func (ptr *QTreeView) IsColumnHidden(column int) bool {
	if ptr.Pointer() != nil {
		return C.QTreeView_IsColumnHidden(ptr.Pointer(), C.int(column)) != 0
	}
	return false
}

func (ptr *QTreeView) IsFirstColumnSpanned(row int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QTreeView_IsFirstColumnSpanned(ptr.Pointer(), C.int(row), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QTreeView) IsRowHidden(row int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QTreeView_IsRowHidden(ptr.Pointer(), C.int(row), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QTreeView) KeyboardSearch(search string) {
	if ptr.Pointer() != nil {
		C.QTreeView_KeyboardSearch(ptr.Pointer(), C.CString(search))
	}
}

func (ptr *QTreeView) Reset() {
	if ptr.Pointer() != nil {
		C.QTreeView_Reset(ptr.Pointer())
	}
}

func (ptr *QTreeView) ResizeColumnToContents(column int) {
	if ptr.Pointer() != nil {
		C.QTreeView_ResizeColumnToContents(ptr.Pointer(), C.int(column))
	}
}

func (ptr *QTreeView) ScrollTo(index core.QModelIndex_ITF, hint QAbstractItemView__ScrollHint) {
	if ptr.Pointer() != nil {
		C.QTreeView_ScrollTo(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(hint))
	}
}

func (ptr *QTreeView) SelectAll() {
	if ptr.Pointer() != nil {
		C.QTreeView_SelectAll(ptr.Pointer())
	}
}

func (ptr *QTreeView) SetColumnHidden(column int, hide bool) {
	if ptr.Pointer() != nil {
		C.QTreeView_SetColumnHidden(ptr.Pointer(), C.int(column), C.int(qt.GoBoolToInt(hide)))
	}
}

func (ptr *QTreeView) SetColumnWidth(column int, width int) {
	if ptr.Pointer() != nil {
		C.QTreeView_SetColumnWidth(ptr.Pointer(), C.int(column), C.int(width))
	}
}

func (ptr *QTreeView) SetExpanded(index core.QModelIndex_ITF, expanded bool) {
	if ptr.Pointer() != nil {
		C.QTreeView_SetExpanded(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(qt.GoBoolToInt(expanded)))
	}
}

func (ptr *QTreeView) SetFirstColumnSpanned(row int, parent core.QModelIndex_ITF, span bool) {
	if ptr.Pointer() != nil {
		C.QTreeView_SetFirstColumnSpanned(ptr.Pointer(), C.int(row), core.PointerFromQModelIndex(parent), C.int(qt.GoBoolToInt(span)))
	}
}

func (ptr *QTreeView) SetHeader(header QHeaderView_ITF) {
	if ptr.Pointer() != nil {
		C.QTreeView_SetHeader(ptr.Pointer(), PointerFromQHeaderView(header))
	}
}

func (ptr *QTreeView) SetModel(model core.QAbstractItemModel_ITF) {
	if ptr.Pointer() != nil {
		C.QTreeView_SetModel(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QTreeView) SetRootIndex(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QTreeView_SetRootIndex(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QTreeView) SetRowHidden(row int, parent core.QModelIndex_ITF, hide bool) {
	if ptr.Pointer() != nil {
		C.QTreeView_SetRowHidden(ptr.Pointer(), C.int(row), core.PointerFromQModelIndex(parent), C.int(qt.GoBoolToInt(hide)))
	}
}

func (ptr *QTreeView) SetSelectionModel(selectionModel core.QItemSelectionModel_ITF) {
	if ptr.Pointer() != nil {
		C.QTreeView_SetSelectionModel(ptr.Pointer(), core.PointerFromQItemSelectionModel(selectionModel))
	}
}

func (ptr *QTreeView) SetTreePosition(index int) {
	if ptr.Pointer() != nil {
		C.QTreeView_SetTreePosition(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QTreeView) ShowColumn(column int) {
	if ptr.Pointer() != nil {
		C.QTreeView_ShowColumn(ptr.Pointer(), C.int(column))
	}
}

func (ptr *QTreeView) SortByColumn(column int, order core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.QTreeView_SortByColumn(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QTreeView) TreePosition() int {
	if ptr.Pointer() != nil {
		return int(C.QTreeView_TreePosition(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTreeView) DestroyQTreeView() {
	if ptr.Pointer() != nil {
		C.QTreeView_DestroyQTreeView(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
