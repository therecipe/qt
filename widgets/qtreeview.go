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

type QTreeViewITF interface {
	QAbstractItemViewITF
	QTreeViewPTR() *QTreeView
}

func PointerFromQTreeView(ptr QTreeViewITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTreeViewPTR().Pointer()
	}
	return nil
}

func QTreeViewFromPointer(ptr unsafe.Pointer) *QTreeView {
	var n = new(QTreeView)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QTreeView_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTreeView) QTreeViewPTR() *QTreeView {
	return ptr
}

func (ptr *QTreeView) AllColumnsShowFocus() bool {
	if ptr.Pointer() != nil {
		return C.QTreeView_AllColumnsShowFocus(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTreeView) AutoExpandDelay() int {
	if ptr.Pointer() != nil {
		return int(C.QTreeView_AutoExpandDelay(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTreeView) Collapse(index core.QModelIndexITF) {
	if ptr.Pointer() != nil {
		C.QTreeView_Collapse(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(index)))
	}
}

func (ptr *QTreeView) Expand(index core.QModelIndexITF) {
	if ptr.Pointer() != nil {
		C.QTreeView_Expand(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(index)))
	}
}

func (ptr *QTreeView) ExpandsOnDoubleClick() bool {
	if ptr.Pointer() != nil {
		return C.QTreeView_ExpandsOnDoubleClick(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTreeView) Indentation() int {
	if ptr.Pointer() != nil {
		return int(C.QTreeView_Indentation(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTreeView) IsAnimated() bool {
	if ptr.Pointer() != nil {
		return C.QTreeView_IsAnimated(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTreeView) IsExpanded(index core.QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QTreeView_IsExpanded(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(index))) != 0
	}
	return false
}

func (ptr *QTreeView) IsHeaderHidden() bool {
	if ptr.Pointer() != nil {
		return C.QTreeView_IsHeaderHidden(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTreeView) IsSortingEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QTreeView_IsSortingEnabled(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTreeView) ItemsExpandable() bool {
	if ptr.Pointer() != nil {
		return C.QTreeView_ItemsExpandable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTreeView) ResetIndentation() {
	if ptr.Pointer() != nil {
		C.QTreeView_ResetIndentation(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTreeView) RootIsDecorated() bool {
	if ptr.Pointer() != nil {
		return C.QTreeView_RootIsDecorated(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTreeView) SetAllColumnsShowFocus(enable bool) {
	if ptr.Pointer() != nil {
		C.QTreeView_SetAllColumnsShowFocus(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QTreeView) SetAnimated(enable bool) {
	if ptr.Pointer() != nil {
		C.QTreeView_SetAnimated(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QTreeView) SetAutoExpandDelay(delay int) {
	if ptr.Pointer() != nil {
		C.QTreeView_SetAutoExpandDelay(C.QtObjectPtr(ptr.Pointer()), C.int(delay))
	}
}

func (ptr *QTreeView) SetExpandsOnDoubleClick(enable bool) {
	if ptr.Pointer() != nil {
		C.QTreeView_SetExpandsOnDoubleClick(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QTreeView) SetHeaderHidden(hide bool) {
	if ptr.Pointer() != nil {
		C.QTreeView_SetHeaderHidden(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(hide)))
	}
}

func (ptr *QTreeView) SetIndentation(i int) {
	if ptr.Pointer() != nil {
		C.QTreeView_SetIndentation(C.QtObjectPtr(ptr.Pointer()), C.int(i))
	}
}

func (ptr *QTreeView) SetItemsExpandable(enable bool) {
	if ptr.Pointer() != nil {
		C.QTreeView_SetItemsExpandable(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QTreeView) SetRootIsDecorated(show bool) {
	if ptr.Pointer() != nil {
		C.QTreeView_SetRootIsDecorated(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(show)))
	}
}

func (ptr *QTreeView) SetSortingEnabled(enable bool) {
	if ptr.Pointer() != nil {
		C.QTreeView_SetSortingEnabled(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QTreeView) SetUniformRowHeights(uniform bool) {
	if ptr.Pointer() != nil {
		C.QTreeView_SetUniformRowHeights(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(uniform)))
	}
}

func (ptr *QTreeView) SetWordWrap(on bool) {
	if ptr.Pointer() != nil {
		C.QTreeView_SetWordWrap(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QTreeView) UniformRowHeights() bool {
	if ptr.Pointer() != nil {
		return C.QTreeView_UniformRowHeights(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTreeView) WordWrap() bool {
	if ptr.Pointer() != nil {
		return C.QTreeView_WordWrap(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func NewQTreeView(parent QWidgetITF) *QTreeView {
	return QTreeViewFromPointer(unsafe.Pointer(C.QTreeView_NewQTreeView(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QTreeView) CollapseAll() {
	if ptr.Pointer() != nil {
		C.QTreeView_CollapseAll(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTreeView) ConnectCollapsed(f func(index core.QModelIndexITF)) {
	if ptr.Pointer() != nil {
		C.QTreeView_ConnectCollapsed(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "collapsed", f)
	}
}

func (ptr *QTreeView) DisconnectCollapsed() {
	if ptr.Pointer() != nil {
		C.QTreeView_DisconnectCollapsed(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "collapsed")
	}
}

//export callbackQTreeViewCollapsed
func callbackQTreeViewCollapsed(ptrName *C.char, index unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "collapsed").(func(*core.QModelIndex))(core.QModelIndexFromPointer(index))
}

func (ptr *QTreeView) ColumnAt(x int) int {
	if ptr.Pointer() != nil {
		return int(C.QTreeView_ColumnAt(C.QtObjectPtr(ptr.Pointer()), C.int(x)))
	}
	return 0
}

func (ptr *QTreeView) ColumnViewportPosition(column int) int {
	if ptr.Pointer() != nil {
		return int(C.QTreeView_ColumnViewportPosition(C.QtObjectPtr(ptr.Pointer()), C.int(column)))
	}
	return 0
}

func (ptr *QTreeView) ColumnWidth(column int) int {
	if ptr.Pointer() != nil {
		return int(C.QTreeView_ColumnWidth(C.QtObjectPtr(ptr.Pointer()), C.int(column)))
	}
	return 0
}

func (ptr *QTreeView) ExpandAll() {
	if ptr.Pointer() != nil {
		C.QTreeView_ExpandAll(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTreeView) ExpandToDepth(depth int) {
	if ptr.Pointer() != nil {
		C.QTreeView_ExpandToDepth(C.QtObjectPtr(ptr.Pointer()), C.int(depth))
	}
}

func (ptr *QTreeView) ConnectExpanded(f func(index core.QModelIndexITF)) {
	if ptr.Pointer() != nil {
		C.QTreeView_ConnectExpanded(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "expanded", f)
	}
}

func (ptr *QTreeView) DisconnectExpanded() {
	if ptr.Pointer() != nil {
		C.QTreeView_DisconnectExpanded(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "expanded")
	}
}

//export callbackQTreeViewExpanded
func callbackQTreeViewExpanded(ptrName *C.char, index unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "expanded").(func(*core.QModelIndex))(core.QModelIndexFromPointer(index))
}

func (ptr *QTreeView) Header() *QHeaderView {
	if ptr.Pointer() != nil {
		return QHeaderViewFromPointer(unsafe.Pointer(C.QTreeView_Header(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QTreeView) HideColumn(column int) {
	if ptr.Pointer() != nil {
		C.QTreeView_HideColumn(C.QtObjectPtr(ptr.Pointer()), C.int(column))
	}
}

func (ptr *QTreeView) IndexAbove(index core.QModelIndexITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		return core.QModelIndexFromPointer(unsafe.Pointer(C.QTreeView_IndexAbove(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(index)))))
	}
	return nil
}

func (ptr *QTreeView) IndexAt(point core.QPointITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		return core.QModelIndexFromPointer(unsafe.Pointer(C.QTreeView_IndexAt(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(point)))))
	}
	return nil
}

func (ptr *QTreeView) IndexBelow(index core.QModelIndexITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		return core.QModelIndexFromPointer(unsafe.Pointer(C.QTreeView_IndexBelow(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(index)))))
	}
	return nil
}

func (ptr *QTreeView) IsColumnHidden(column int) bool {
	if ptr.Pointer() != nil {
		return C.QTreeView_IsColumnHidden(C.QtObjectPtr(ptr.Pointer()), C.int(column)) != 0
	}
	return false
}

func (ptr *QTreeView) IsFirstColumnSpanned(row int, parent core.QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QTreeView_IsFirstColumnSpanned(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.QtObjectPtr(core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QTreeView) IsRowHidden(row int, parent core.QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QTreeView_IsRowHidden(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.QtObjectPtr(core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QTreeView) KeyboardSearch(search string) {
	if ptr.Pointer() != nil {
		C.QTreeView_KeyboardSearch(C.QtObjectPtr(ptr.Pointer()), C.CString(search))
	}
}

func (ptr *QTreeView) Reset() {
	if ptr.Pointer() != nil {
		C.QTreeView_Reset(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTreeView) ResizeColumnToContents(column int) {
	if ptr.Pointer() != nil {
		C.QTreeView_ResizeColumnToContents(C.QtObjectPtr(ptr.Pointer()), C.int(column))
	}
}

func (ptr *QTreeView) ScrollTo(index core.QModelIndexITF, hint QAbstractItemView__ScrollHint) {
	if ptr.Pointer() != nil {
		C.QTreeView_ScrollTo(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(index)), C.int(hint))
	}
}

func (ptr *QTreeView) SelectAll() {
	if ptr.Pointer() != nil {
		C.QTreeView_SelectAll(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTreeView) SetColumnHidden(column int, hide bool) {
	if ptr.Pointer() != nil {
		C.QTreeView_SetColumnHidden(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.int(qt.GoBoolToInt(hide)))
	}
}

func (ptr *QTreeView) SetColumnWidth(column int, width int) {
	if ptr.Pointer() != nil {
		C.QTreeView_SetColumnWidth(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.int(width))
	}
}

func (ptr *QTreeView) SetExpanded(index core.QModelIndexITF, expanded bool) {
	if ptr.Pointer() != nil {
		C.QTreeView_SetExpanded(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(index)), C.int(qt.GoBoolToInt(expanded)))
	}
}

func (ptr *QTreeView) SetFirstColumnSpanned(row int, parent core.QModelIndexITF, span bool) {
	if ptr.Pointer() != nil {
		C.QTreeView_SetFirstColumnSpanned(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.QtObjectPtr(core.PointerFromQModelIndex(parent)), C.int(qt.GoBoolToInt(span)))
	}
}

func (ptr *QTreeView) SetHeader(header QHeaderViewITF) {
	if ptr.Pointer() != nil {
		C.QTreeView_SetHeader(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQHeaderView(header)))
	}
}

func (ptr *QTreeView) SetModel(model core.QAbstractItemModelITF) {
	if ptr.Pointer() != nil {
		C.QTreeView_SetModel(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQAbstractItemModel(model)))
	}
}

func (ptr *QTreeView) SetRootIndex(index core.QModelIndexITF) {
	if ptr.Pointer() != nil {
		C.QTreeView_SetRootIndex(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(index)))
	}
}

func (ptr *QTreeView) SetRowHidden(row int, parent core.QModelIndexITF, hide bool) {
	if ptr.Pointer() != nil {
		C.QTreeView_SetRowHidden(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.QtObjectPtr(core.PointerFromQModelIndex(parent)), C.int(qt.GoBoolToInt(hide)))
	}
}

func (ptr *QTreeView) SetSelectionModel(selectionModel core.QItemSelectionModelITF) {
	if ptr.Pointer() != nil {
		C.QTreeView_SetSelectionModel(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQItemSelectionModel(selectionModel)))
	}
}

func (ptr *QTreeView) SetTreePosition(index int) {
	if ptr.Pointer() != nil {
		C.QTreeView_SetTreePosition(C.QtObjectPtr(ptr.Pointer()), C.int(index))
	}
}

func (ptr *QTreeView) ShowColumn(column int) {
	if ptr.Pointer() != nil {
		C.QTreeView_ShowColumn(C.QtObjectPtr(ptr.Pointer()), C.int(column))
	}
}

func (ptr *QTreeView) SortByColumn(column int, order core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.QTreeView_SortByColumn(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.int(order))
	}
}

func (ptr *QTreeView) TreePosition() int {
	if ptr.Pointer() != nil {
		return int(C.QTreeView_TreePosition(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTreeView) DestroyQTreeView() {
	if ptr.Pointer() != nil {
		C.QTreeView_DestroyQTreeView(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
