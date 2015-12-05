package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
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
	for len(n.ObjectName()) < len("QTreeView_") {
		n.SetObjectName("QTreeView_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTreeView) QTreeView_PTR() *QTreeView {
	return ptr
}

func (ptr *QTreeView) AllColumnsShowFocus() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::allColumnsShowFocus")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTreeView_AllColumnsShowFocus(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTreeView) AutoExpandDelay() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::autoExpandDelay")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTreeView_AutoExpandDelay(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTreeView) Collapse(index core.QModelIndex_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::collapse")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeView_Collapse(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QTreeView) Expand(index core.QModelIndex_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::expand")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeView_Expand(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QTreeView) ExpandsOnDoubleClick() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::expandsOnDoubleClick")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTreeView_ExpandsOnDoubleClick(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTreeView) Indentation() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::indentation")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTreeView_Indentation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTreeView) IsAnimated() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::isAnimated")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTreeView_IsAnimated(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTreeView) IsExpanded(index core.QModelIndex_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::isExpanded")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTreeView_IsExpanded(ptr.Pointer(), core.PointerFromQModelIndex(index)) != 0
	}
	return false
}

func (ptr *QTreeView) IsHeaderHidden() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::isHeaderHidden")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTreeView_IsHeaderHidden(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTreeView) IsSortingEnabled() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::isSortingEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTreeView_IsSortingEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTreeView) ItemsExpandable() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::itemsExpandable")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTreeView_ItemsExpandable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTreeView) ResetIndentation() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::resetIndentation")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeView_ResetIndentation(ptr.Pointer())
	}
}

func (ptr *QTreeView) RootIsDecorated() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::rootIsDecorated")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTreeView_RootIsDecorated(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTreeView) SetAllColumnsShowFocus(enable bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::setAllColumnsShowFocus")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeView_SetAllColumnsShowFocus(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QTreeView) SetAnimated(enable bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::setAnimated")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeView_SetAnimated(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QTreeView) SetAutoExpandDelay(delay int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::setAutoExpandDelay")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeView_SetAutoExpandDelay(ptr.Pointer(), C.int(delay))
	}
}

func (ptr *QTreeView) SetExpandsOnDoubleClick(enable bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::setExpandsOnDoubleClick")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeView_SetExpandsOnDoubleClick(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QTreeView) SetHeaderHidden(hide bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::setHeaderHidden")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeView_SetHeaderHidden(ptr.Pointer(), C.int(qt.GoBoolToInt(hide)))
	}
}

func (ptr *QTreeView) SetIndentation(i int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::setIndentation")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeView_SetIndentation(ptr.Pointer(), C.int(i))
	}
}

func (ptr *QTreeView) SetItemsExpandable(enable bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::setItemsExpandable")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeView_SetItemsExpandable(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QTreeView) SetRootIsDecorated(show bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::setRootIsDecorated")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeView_SetRootIsDecorated(ptr.Pointer(), C.int(qt.GoBoolToInt(show)))
	}
}

func (ptr *QTreeView) SetSortingEnabled(enable bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::setSortingEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeView_SetSortingEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QTreeView) SetUniformRowHeights(uniform bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::setUniformRowHeights")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeView_SetUniformRowHeights(ptr.Pointer(), C.int(qt.GoBoolToInt(uniform)))
	}
}

func (ptr *QTreeView) SetWordWrap(on bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::setWordWrap")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeView_SetWordWrap(ptr.Pointer(), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QTreeView) UniformRowHeights() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::uniformRowHeights")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTreeView_UniformRowHeights(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTreeView) WordWrap() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::wordWrap")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTreeView_WordWrap(ptr.Pointer()) != 0
	}
	return false
}

func NewQTreeView(parent QWidget_ITF) *QTreeView {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::QTreeView")
		}
	}()

	return NewQTreeViewFromPointer(C.QTreeView_NewQTreeView(PointerFromQWidget(parent)))
}

func (ptr *QTreeView) CollapseAll() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::collapseAll")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeView_CollapseAll(ptr.Pointer())
	}
}

func (ptr *QTreeView) ConnectCollapsed(f func(index *core.QModelIndex)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::collapsed")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeView_ConnectCollapsed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "collapsed", f)
	}
}

func (ptr *QTreeView) DisconnectCollapsed() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::collapsed")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeView_DisconnectCollapsed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "collapsed")
	}
}

//export callbackQTreeViewCollapsed
func callbackQTreeViewCollapsed(ptrName *C.char, index unsafe.Pointer) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::collapsed")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "collapsed").(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
}

func (ptr *QTreeView) ColumnAt(x int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::columnAt")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTreeView_ColumnAt(ptr.Pointer(), C.int(x)))
	}
	return 0
}

func (ptr *QTreeView) ColumnViewportPosition(column int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::columnViewportPosition")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTreeView_ColumnViewportPosition(ptr.Pointer(), C.int(column)))
	}
	return 0
}

func (ptr *QTreeView) ColumnWidth(column int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::columnWidth")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTreeView_ColumnWidth(ptr.Pointer(), C.int(column)))
	}
	return 0
}

func (ptr *QTreeView) ExpandAll() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::expandAll")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeView_ExpandAll(ptr.Pointer())
	}
}

func (ptr *QTreeView) ExpandToDepth(depth int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::expandToDepth")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeView_ExpandToDepth(ptr.Pointer(), C.int(depth))
	}
}

func (ptr *QTreeView) ConnectExpanded(f func(index *core.QModelIndex)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::expanded")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeView_ConnectExpanded(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "expanded", f)
	}
}

func (ptr *QTreeView) DisconnectExpanded() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::expanded")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeView_DisconnectExpanded(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "expanded")
	}
}

//export callbackQTreeViewExpanded
func callbackQTreeViewExpanded(ptrName *C.char, index unsafe.Pointer) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::expanded")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "expanded").(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
}

func (ptr *QTreeView) Header() *QHeaderView {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::header")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQHeaderViewFromPointer(C.QTreeView_Header(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTreeView) HideColumn(column int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::hideColumn")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeView_HideColumn(ptr.Pointer(), C.int(column))
	}
}

func (ptr *QTreeView) IndexAbove(index core.QModelIndex_ITF) *core.QModelIndex {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::indexAbove")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QTreeView_IndexAbove(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return nil
}

func (ptr *QTreeView) IndexAt(point core.QPoint_ITF) *core.QModelIndex {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::indexAt")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QTreeView_IndexAt(ptr.Pointer(), core.PointerFromQPoint(point)))
	}
	return nil
}

func (ptr *QTreeView) IndexBelow(index core.QModelIndex_ITF) *core.QModelIndex {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::indexBelow")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QTreeView_IndexBelow(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return nil
}

func (ptr *QTreeView) IsColumnHidden(column int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::isColumnHidden")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTreeView_IsColumnHidden(ptr.Pointer(), C.int(column)) != 0
	}
	return false
}

func (ptr *QTreeView) IsFirstColumnSpanned(row int, parent core.QModelIndex_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::isFirstColumnSpanned")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTreeView_IsFirstColumnSpanned(ptr.Pointer(), C.int(row), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QTreeView) IsRowHidden(row int, parent core.QModelIndex_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::isRowHidden")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTreeView_IsRowHidden(ptr.Pointer(), C.int(row), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QTreeView) KeyboardSearch(search string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::keyboardSearch")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeView_KeyboardSearch(ptr.Pointer(), C.CString(search))
	}
}

func (ptr *QTreeView) Reset() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::reset")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeView_Reset(ptr.Pointer())
	}
}

func (ptr *QTreeView) ResizeColumnToContents(column int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::resizeColumnToContents")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeView_ResizeColumnToContents(ptr.Pointer(), C.int(column))
	}
}

func (ptr *QTreeView) ScrollTo(index core.QModelIndex_ITF, hint QAbstractItemView__ScrollHint) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::scrollTo")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeView_ScrollTo(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(hint))
	}
}

func (ptr *QTreeView) SelectAll() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::selectAll")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeView_SelectAll(ptr.Pointer())
	}
}

func (ptr *QTreeView) SetColumnHidden(column int, hide bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::setColumnHidden")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeView_SetColumnHidden(ptr.Pointer(), C.int(column), C.int(qt.GoBoolToInt(hide)))
	}
}

func (ptr *QTreeView) SetColumnWidth(column int, width int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::setColumnWidth")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeView_SetColumnWidth(ptr.Pointer(), C.int(column), C.int(width))
	}
}

func (ptr *QTreeView) SetExpanded(index core.QModelIndex_ITF, expanded bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::setExpanded")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeView_SetExpanded(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(qt.GoBoolToInt(expanded)))
	}
}

func (ptr *QTreeView) SetFirstColumnSpanned(row int, parent core.QModelIndex_ITF, span bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::setFirstColumnSpanned")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeView_SetFirstColumnSpanned(ptr.Pointer(), C.int(row), core.PointerFromQModelIndex(parent), C.int(qt.GoBoolToInt(span)))
	}
}

func (ptr *QTreeView) SetHeader(header QHeaderView_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::setHeader")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeView_SetHeader(ptr.Pointer(), PointerFromQHeaderView(header))
	}
}

func (ptr *QTreeView) SetModel(model core.QAbstractItemModel_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::setModel")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeView_SetModel(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QTreeView) SetRootIndex(index core.QModelIndex_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::setRootIndex")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeView_SetRootIndex(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QTreeView) SetRowHidden(row int, parent core.QModelIndex_ITF, hide bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::setRowHidden")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeView_SetRowHidden(ptr.Pointer(), C.int(row), core.PointerFromQModelIndex(parent), C.int(qt.GoBoolToInt(hide)))
	}
}

func (ptr *QTreeView) SetSelectionModel(selectionModel core.QItemSelectionModel_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::setSelectionModel")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeView_SetSelectionModel(ptr.Pointer(), core.PointerFromQItemSelectionModel(selectionModel))
	}
}

func (ptr *QTreeView) SetTreePosition(index int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::setTreePosition")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeView_SetTreePosition(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QTreeView) ShowColumn(column int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::showColumn")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeView_ShowColumn(ptr.Pointer(), C.int(column))
	}
}

func (ptr *QTreeView) SortByColumn(column int, order core.Qt__SortOrder) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::sortByColumn")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeView_SortByColumn(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QTreeView) TreePosition() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::treePosition")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTreeView_TreePosition(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTreeView) DestroyQTreeView() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeView::~QTreeView")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeView_DestroyQTreeView(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
