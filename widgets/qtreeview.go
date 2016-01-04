package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
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
		n.SetObjectName("QTreeView_" + qt.Identifier())
	}
	return n
}

func (ptr *QTreeView) QTreeView_PTR() *QTreeView {
	return ptr
}

func (ptr *QTreeView) AllColumnsShowFocus() bool {
	defer qt.Recovering("QTreeView::allColumnsShowFocus")

	if ptr.Pointer() != nil {
		return C.QTreeView_AllColumnsShowFocus(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTreeView) AutoExpandDelay() int {
	defer qt.Recovering("QTreeView::autoExpandDelay")

	if ptr.Pointer() != nil {
		return int(C.QTreeView_AutoExpandDelay(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTreeView) Collapse(index core.QModelIndex_ITF) {
	defer qt.Recovering("QTreeView::collapse")

	if ptr.Pointer() != nil {
		C.QTreeView_Collapse(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QTreeView) Expand(index core.QModelIndex_ITF) {
	defer qt.Recovering("QTreeView::expand")

	if ptr.Pointer() != nil {
		C.QTreeView_Expand(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QTreeView) ExpandsOnDoubleClick() bool {
	defer qt.Recovering("QTreeView::expandsOnDoubleClick")

	if ptr.Pointer() != nil {
		return C.QTreeView_ExpandsOnDoubleClick(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTreeView) Indentation() int {
	defer qt.Recovering("QTreeView::indentation")

	if ptr.Pointer() != nil {
		return int(C.QTreeView_Indentation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTreeView) IsAnimated() bool {
	defer qt.Recovering("QTreeView::isAnimated")

	if ptr.Pointer() != nil {
		return C.QTreeView_IsAnimated(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTreeView) IsExpanded(index core.QModelIndex_ITF) bool {
	defer qt.Recovering("QTreeView::isExpanded")

	if ptr.Pointer() != nil {
		return C.QTreeView_IsExpanded(ptr.Pointer(), core.PointerFromQModelIndex(index)) != 0
	}
	return false
}

func (ptr *QTreeView) IsHeaderHidden() bool {
	defer qt.Recovering("QTreeView::isHeaderHidden")

	if ptr.Pointer() != nil {
		return C.QTreeView_IsHeaderHidden(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTreeView) IsSortingEnabled() bool {
	defer qt.Recovering("QTreeView::isSortingEnabled")

	if ptr.Pointer() != nil {
		return C.QTreeView_IsSortingEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTreeView) ItemsExpandable() bool {
	defer qt.Recovering("QTreeView::itemsExpandable")

	if ptr.Pointer() != nil {
		return C.QTreeView_ItemsExpandable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTreeView) ResetIndentation() {
	defer qt.Recovering("QTreeView::resetIndentation")

	if ptr.Pointer() != nil {
		C.QTreeView_ResetIndentation(ptr.Pointer())
	}
}

func (ptr *QTreeView) RootIsDecorated() bool {
	defer qt.Recovering("QTreeView::rootIsDecorated")

	if ptr.Pointer() != nil {
		return C.QTreeView_RootIsDecorated(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTreeView) SetAllColumnsShowFocus(enable bool) {
	defer qt.Recovering("QTreeView::setAllColumnsShowFocus")

	if ptr.Pointer() != nil {
		C.QTreeView_SetAllColumnsShowFocus(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QTreeView) SetAnimated(enable bool) {
	defer qt.Recovering("QTreeView::setAnimated")

	if ptr.Pointer() != nil {
		C.QTreeView_SetAnimated(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QTreeView) SetAutoExpandDelay(delay int) {
	defer qt.Recovering("QTreeView::setAutoExpandDelay")

	if ptr.Pointer() != nil {
		C.QTreeView_SetAutoExpandDelay(ptr.Pointer(), C.int(delay))
	}
}

func (ptr *QTreeView) SetExpandsOnDoubleClick(enable bool) {
	defer qt.Recovering("QTreeView::setExpandsOnDoubleClick")

	if ptr.Pointer() != nil {
		C.QTreeView_SetExpandsOnDoubleClick(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QTreeView) SetHeaderHidden(hide bool) {
	defer qt.Recovering("QTreeView::setHeaderHidden")

	if ptr.Pointer() != nil {
		C.QTreeView_SetHeaderHidden(ptr.Pointer(), C.int(qt.GoBoolToInt(hide)))
	}
}

func (ptr *QTreeView) SetIndentation(i int) {
	defer qt.Recovering("QTreeView::setIndentation")

	if ptr.Pointer() != nil {
		C.QTreeView_SetIndentation(ptr.Pointer(), C.int(i))
	}
}

func (ptr *QTreeView) SetItemsExpandable(enable bool) {
	defer qt.Recovering("QTreeView::setItemsExpandable")

	if ptr.Pointer() != nil {
		C.QTreeView_SetItemsExpandable(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QTreeView) SetRootIsDecorated(show bool) {
	defer qt.Recovering("QTreeView::setRootIsDecorated")

	if ptr.Pointer() != nil {
		C.QTreeView_SetRootIsDecorated(ptr.Pointer(), C.int(qt.GoBoolToInt(show)))
	}
}

func (ptr *QTreeView) SetSortingEnabled(enable bool) {
	defer qt.Recovering("QTreeView::setSortingEnabled")

	if ptr.Pointer() != nil {
		C.QTreeView_SetSortingEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QTreeView) SetUniformRowHeights(uniform bool) {
	defer qt.Recovering("QTreeView::setUniformRowHeights")

	if ptr.Pointer() != nil {
		C.QTreeView_SetUniformRowHeights(ptr.Pointer(), C.int(qt.GoBoolToInt(uniform)))
	}
}

func (ptr *QTreeView) SetWordWrap(on bool) {
	defer qt.Recovering("QTreeView::setWordWrap")

	if ptr.Pointer() != nil {
		C.QTreeView_SetWordWrap(ptr.Pointer(), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QTreeView) UniformRowHeights() bool {
	defer qt.Recovering("QTreeView::uniformRowHeights")

	if ptr.Pointer() != nil {
		return C.QTreeView_UniformRowHeights(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTreeView) WordWrap() bool {
	defer qt.Recovering("QTreeView::wordWrap")

	if ptr.Pointer() != nil {
		return C.QTreeView_WordWrap(ptr.Pointer()) != 0
	}
	return false
}

func NewQTreeView(parent QWidget_ITF) *QTreeView {
	defer qt.Recovering("QTreeView::QTreeView")

	return NewQTreeViewFromPointer(C.QTreeView_NewQTreeView(PointerFromQWidget(parent)))
}

func (ptr *QTreeView) CollapseAll() {
	defer qt.Recovering("QTreeView::collapseAll")

	if ptr.Pointer() != nil {
		C.QTreeView_CollapseAll(ptr.Pointer())
	}
}

func (ptr *QTreeView) ConnectCollapsed(f func(index *core.QModelIndex)) {
	defer qt.Recovering("connect QTreeView::collapsed")

	if ptr.Pointer() != nil {
		C.QTreeView_ConnectCollapsed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "collapsed", f)
	}
}

func (ptr *QTreeView) DisconnectCollapsed() {
	defer qt.Recovering("disconnect QTreeView::collapsed")

	if ptr.Pointer() != nil {
		C.QTreeView_DisconnectCollapsed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "collapsed")
	}
}

//export callbackQTreeViewCollapsed
func callbackQTreeViewCollapsed(ptr unsafe.Pointer, ptrName *C.char, index unsafe.Pointer) {
	defer qt.Recovering("callback QTreeView::collapsed")

	if signal := qt.GetSignal(C.GoString(ptrName), "collapsed"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	}

}

func (ptr *QTreeView) Collapsed(index core.QModelIndex_ITF) {
	defer qt.Recovering("QTreeView::collapsed")

	if ptr.Pointer() != nil {
		C.QTreeView_Collapsed(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QTreeView) ColumnAt(x int) int {
	defer qt.Recovering("QTreeView::columnAt")

	if ptr.Pointer() != nil {
		return int(C.QTreeView_ColumnAt(ptr.Pointer(), C.int(x)))
	}
	return 0
}

func (ptr *QTreeView) ColumnViewportPosition(column int) int {
	defer qt.Recovering("QTreeView::columnViewportPosition")

	if ptr.Pointer() != nil {
		return int(C.QTreeView_ColumnViewportPosition(ptr.Pointer(), C.int(column)))
	}
	return 0
}

func (ptr *QTreeView) ColumnWidth(column int) int {
	defer qt.Recovering("QTreeView::columnWidth")

	if ptr.Pointer() != nil {
		return int(C.QTreeView_ColumnWidth(ptr.Pointer(), C.int(column)))
	}
	return 0
}

func (ptr *QTreeView) ConnectCurrentChanged(f func(current *core.QModelIndex, previous *core.QModelIndex)) {
	defer qt.Recovering("connect QTreeView::currentChanged")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "currentChanged", f)
	}
}

func (ptr *QTreeView) DisconnectCurrentChanged() {
	defer qt.Recovering("disconnect QTreeView::currentChanged")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "currentChanged")
	}
}

//export callbackQTreeViewCurrentChanged
func callbackQTreeViewCurrentChanged(ptr unsafe.Pointer, ptrName *C.char, current unsafe.Pointer, previous unsafe.Pointer) {
	defer qt.Recovering("callback QTreeView::currentChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "currentChanged"); signal != nil {
		signal.(func(*core.QModelIndex, *core.QModelIndex))(core.NewQModelIndexFromPointer(current), core.NewQModelIndexFromPointer(previous))
	} else {
		NewQTreeViewFromPointer(ptr).CurrentChangedDefault(core.NewQModelIndexFromPointer(current), core.NewQModelIndexFromPointer(previous))
	}
}

func (ptr *QTreeView) CurrentChanged(current core.QModelIndex_ITF, previous core.QModelIndex_ITF) {
	defer qt.Recovering("QTreeView::currentChanged")

	if ptr.Pointer() != nil {
		C.QTreeView_CurrentChanged(ptr.Pointer(), core.PointerFromQModelIndex(current), core.PointerFromQModelIndex(previous))
	}
}

func (ptr *QTreeView) CurrentChangedDefault(current core.QModelIndex_ITF, previous core.QModelIndex_ITF) {
	defer qt.Recovering("QTreeView::currentChanged")

	if ptr.Pointer() != nil {
		C.QTreeView_CurrentChangedDefault(ptr.Pointer(), core.PointerFromQModelIndex(current), core.PointerFromQModelIndex(previous))
	}
}

func (ptr *QTreeView) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QTreeView::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QTreeView) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QTreeView::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQTreeViewDragMoveEvent
func callbackQTreeViewDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTreeView::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQTreeViewFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QTreeView) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QTreeView::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QTreeView) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QTreeView::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QTreeView) ConnectDrawBranches(f func(painter *gui.QPainter, rect *core.QRect, index *core.QModelIndex)) {
	defer qt.Recovering("connect QTreeView::drawBranches")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "drawBranches", f)
	}
}

func (ptr *QTreeView) DisconnectDrawBranches() {
	defer qt.Recovering("disconnect QTreeView::drawBranches")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "drawBranches")
	}
}

//export callbackQTreeViewDrawBranches
func callbackQTreeViewDrawBranches(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer, rect unsafe.Pointer, index unsafe.Pointer) {
	defer qt.Recovering("callback QTreeView::drawBranches")

	if signal := qt.GetSignal(C.GoString(ptrName), "drawBranches"); signal != nil {
		signal.(func(*gui.QPainter, *core.QRect, *core.QModelIndex))(gui.NewQPainterFromPointer(painter), core.NewQRectFromPointer(rect), core.NewQModelIndexFromPointer(index))
	} else {
		NewQTreeViewFromPointer(ptr).DrawBranchesDefault(gui.NewQPainterFromPointer(painter), core.NewQRectFromPointer(rect), core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QTreeView) DrawBranches(painter gui.QPainter_ITF, rect core.QRect_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("QTreeView::drawBranches")

	if ptr.Pointer() != nil {
		C.QTreeView_DrawBranches(ptr.Pointer(), gui.PointerFromQPainter(painter), core.PointerFromQRect(rect), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QTreeView) DrawBranchesDefault(painter gui.QPainter_ITF, rect core.QRect_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("QTreeView::drawBranches")

	if ptr.Pointer() != nil {
		C.QTreeView_DrawBranchesDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), core.PointerFromQRect(rect), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QTreeView) ExpandAll() {
	defer qt.Recovering("QTreeView::expandAll")

	if ptr.Pointer() != nil {
		C.QTreeView_ExpandAll(ptr.Pointer())
	}
}

func (ptr *QTreeView) ExpandToDepth(depth int) {
	defer qt.Recovering("QTreeView::expandToDepth")

	if ptr.Pointer() != nil {
		C.QTreeView_ExpandToDepth(ptr.Pointer(), C.int(depth))
	}
}

func (ptr *QTreeView) ConnectExpanded(f func(index *core.QModelIndex)) {
	defer qt.Recovering("connect QTreeView::expanded")

	if ptr.Pointer() != nil {
		C.QTreeView_ConnectExpanded(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "expanded", f)
	}
}

func (ptr *QTreeView) DisconnectExpanded() {
	defer qt.Recovering("disconnect QTreeView::expanded")

	if ptr.Pointer() != nil {
		C.QTreeView_DisconnectExpanded(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "expanded")
	}
}

//export callbackQTreeViewExpanded
func callbackQTreeViewExpanded(ptr unsafe.Pointer, ptrName *C.char, index unsafe.Pointer) {
	defer qt.Recovering("callback QTreeView::expanded")

	if signal := qt.GetSignal(C.GoString(ptrName), "expanded"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	}

}

func (ptr *QTreeView) Expanded(index core.QModelIndex_ITF) {
	defer qt.Recovering("QTreeView::expanded")

	if ptr.Pointer() != nil {
		C.QTreeView_Expanded(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QTreeView) Header() *QHeaderView {
	defer qt.Recovering("QTreeView::header")

	if ptr.Pointer() != nil {
		return NewQHeaderViewFromPointer(C.QTreeView_Header(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTreeView) HideColumn(column int) {
	defer qt.Recovering("QTreeView::hideColumn")

	if ptr.Pointer() != nil {
		C.QTreeView_HideColumn(ptr.Pointer(), C.int(column))
	}
}

func (ptr *QTreeView) HorizontalOffset() int {
	defer qt.Recovering("QTreeView::horizontalOffset")

	if ptr.Pointer() != nil {
		return int(C.QTreeView_HorizontalOffset(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTreeView) IndexAbove(index core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QTreeView::indexAbove")

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QTreeView_IndexAbove(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return nil
}

func (ptr *QTreeView) IndexAt(point core.QPoint_ITF) *core.QModelIndex {
	defer qt.Recovering("QTreeView::indexAt")

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QTreeView_IndexAt(ptr.Pointer(), core.PointerFromQPoint(point)))
	}
	return nil
}

func (ptr *QTreeView) IndexBelow(index core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QTreeView::indexBelow")

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QTreeView_IndexBelow(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return nil
}

func (ptr *QTreeView) IsColumnHidden(column int) bool {
	defer qt.Recovering("QTreeView::isColumnHidden")

	if ptr.Pointer() != nil {
		return C.QTreeView_IsColumnHidden(ptr.Pointer(), C.int(column)) != 0
	}
	return false
}

func (ptr *QTreeView) IsFirstColumnSpanned(row int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QTreeView::isFirstColumnSpanned")

	if ptr.Pointer() != nil {
		return C.QTreeView_IsFirstColumnSpanned(ptr.Pointer(), C.int(row), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QTreeView) IsIndexHidden(index core.QModelIndex_ITF) bool {
	defer qt.Recovering("QTreeView::isIndexHidden")

	if ptr.Pointer() != nil {
		return C.QTreeView_IsIndexHidden(ptr.Pointer(), core.PointerFromQModelIndex(index)) != 0
	}
	return false
}

func (ptr *QTreeView) IsRowHidden(row int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QTreeView::isRowHidden")

	if ptr.Pointer() != nil {
		return C.QTreeView_IsRowHidden(ptr.Pointer(), C.int(row), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QTreeView) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QTreeView::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QTreeView) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QTreeView::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQTreeViewKeyPressEvent
func callbackQTreeViewKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTreeView::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQTreeViewFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QTreeView) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QTreeView::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QTreeView) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QTreeView::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QTreeView) ConnectKeyboardSearch(f func(search string)) {
	defer qt.Recovering("connect QTreeView::keyboardSearch")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyboardSearch", f)
	}
}

func (ptr *QTreeView) DisconnectKeyboardSearch() {
	defer qt.Recovering("disconnect QTreeView::keyboardSearch")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyboardSearch")
	}
}

//export callbackQTreeViewKeyboardSearch
func callbackQTreeViewKeyboardSearch(ptr unsafe.Pointer, ptrName *C.char, search *C.char) {
	defer qt.Recovering("callback QTreeView::keyboardSearch")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyboardSearch"); signal != nil {
		signal.(func(string))(C.GoString(search))
	} else {
		NewQTreeViewFromPointer(ptr).KeyboardSearchDefault(C.GoString(search))
	}
}

func (ptr *QTreeView) KeyboardSearch(search string) {
	defer qt.Recovering("QTreeView::keyboardSearch")

	if ptr.Pointer() != nil {
		C.QTreeView_KeyboardSearch(ptr.Pointer(), C.CString(search))
	}
}

func (ptr *QTreeView) KeyboardSearchDefault(search string) {
	defer qt.Recovering("QTreeView::keyboardSearch")

	if ptr.Pointer() != nil {
		C.QTreeView_KeyboardSearchDefault(ptr.Pointer(), C.CString(search))
	}
}

func (ptr *QTreeView) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QTreeView::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QTreeView) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QTreeView::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQTreeViewMouseDoubleClickEvent
func callbackQTreeViewMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTreeView::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQTreeViewFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QTreeView) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QTreeView::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QTreeView) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QTreeView::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QTreeView) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QTreeView::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QTreeView) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QTreeView::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQTreeViewMouseMoveEvent
func callbackQTreeViewMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTreeView::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQTreeViewFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QTreeView) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QTreeView::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QTreeView) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QTreeView::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QTreeView) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QTreeView::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QTreeView) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QTreeView::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQTreeViewMousePressEvent
func callbackQTreeViewMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTreeView::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQTreeViewFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QTreeView) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QTreeView::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QTreeView) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QTreeView::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QTreeView) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QTreeView::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QTreeView) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QTreeView::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQTreeViewMouseReleaseEvent
func callbackQTreeViewMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTreeView::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQTreeViewFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QTreeView) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QTreeView::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QTreeView) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QTreeView::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QTreeView) MoveCursor(cursorAction QAbstractItemView__CursorAction, modifiers core.Qt__KeyboardModifier) *core.QModelIndex {
	defer qt.Recovering("QTreeView::moveCursor")

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QTreeView_MoveCursor(ptr.Pointer(), C.int(cursorAction), C.int(modifiers)))
	}
	return nil
}

func (ptr *QTreeView) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QTreeView::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QTreeView) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QTreeView::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQTreeViewPaintEvent
func callbackQTreeViewPaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTreeView::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQTreeViewFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QTreeView) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QTreeView::paintEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QTreeView) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QTreeView::paintEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QTreeView) ConnectReset(f func()) {
	defer qt.Recovering("connect QTreeView::reset")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "reset", f)
	}
}

func (ptr *QTreeView) DisconnectReset() {
	defer qt.Recovering("disconnect QTreeView::reset")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "reset")
	}
}

//export callbackQTreeViewReset
func callbackQTreeViewReset(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QTreeView::reset")

	if signal := qt.GetSignal(C.GoString(ptrName), "reset"); signal != nil {
		signal.(func())()
	} else {
		NewQTreeViewFromPointer(ptr).ResetDefault()
	}
}

func (ptr *QTreeView) Reset() {
	defer qt.Recovering("QTreeView::reset")

	if ptr.Pointer() != nil {
		C.QTreeView_Reset(ptr.Pointer())
	}
}

func (ptr *QTreeView) ResetDefault() {
	defer qt.Recovering("QTreeView::reset")

	if ptr.Pointer() != nil {
		C.QTreeView_ResetDefault(ptr.Pointer())
	}
}

func (ptr *QTreeView) ResizeColumnToContents(column int) {
	defer qt.Recovering("QTreeView::resizeColumnToContents")

	if ptr.Pointer() != nil {
		C.QTreeView_ResizeColumnToContents(ptr.Pointer(), C.int(column))
	}
}

func (ptr *QTreeView) ConnectRowsAboutToBeRemoved(f func(parent *core.QModelIndex, start int, end int)) {
	defer qt.Recovering("connect QTreeView::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "rowsAboutToBeRemoved", f)
	}
}

func (ptr *QTreeView) DisconnectRowsAboutToBeRemoved() {
	defer qt.Recovering("disconnect QTreeView::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "rowsAboutToBeRemoved")
	}
}

//export callbackQTreeViewRowsAboutToBeRemoved
func callbackQTreeViewRowsAboutToBeRemoved(ptr unsafe.Pointer, ptrName *C.char, parent unsafe.Pointer, start C.int, end C.int) {
	defer qt.Recovering("callback QTreeView::rowsAboutToBeRemoved")

	if signal := qt.GetSignal(C.GoString(ptrName), "rowsAboutToBeRemoved"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(start), int(end))
	} else {
		NewQTreeViewFromPointer(ptr).RowsAboutToBeRemovedDefault(core.NewQModelIndexFromPointer(parent), int(start), int(end))
	}
}

func (ptr *QTreeView) RowsAboutToBeRemoved(parent core.QModelIndex_ITF, start int, end int) {
	defer qt.Recovering("QTreeView::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {
		C.QTreeView_RowsAboutToBeRemoved(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(start), C.int(end))
	}
}

func (ptr *QTreeView) RowsAboutToBeRemovedDefault(parent core.QModelIndex_ITF, start int, end int) {
	defer qt.Recovering("QTreeView::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {
		C.QTreeView_RowsAboutToBeRemovedDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(start), C.int(end))
	}
}

func (ptr *QTreeView) ConnectRowsInserted(f func(parent *core.QModelIndex, start int, end int)) {
	defer qt.Recovering("connect QTreeView::rowsInserted")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "rowsInserted", f)
	}
}

func (ptr *QTreeView) DisconnectRowsInserted() {
	defer qt.Recovering("disconnect QTreeView::rowsInserted")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "rowsInserted")
	}
}

//export callbackQTreeViewRowsInserted
func callbackQTreeViewRowsInserted(ptr unsafe.Pointer, ptrName *C.char, parent unsafe.Pointer, start C.int, end C.int) {
	defer qt.Recovering("callback QTreeView::rowsInserted")

	if signal := qt.GetSignal(C.GoString(ptrName), "rowsInserted"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(start), int(end))
	} else {
		NewQTreeViewFromPointer(ptr).RowsInsertedDefault(core.NewQModelIndexFromPointer(parent), int(start), int(end))
	}
}

func (ptr *QTreeView) RowsInserted(parent core.QModelIndex_ITF, start int, end int) {
	defer qt.Recovering("QTreeView::rowsInserted")

	if ptr.Pointer() != nil {
		C.QTreeView_RowsInserted(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(start), C.int(end))
	}
}

func (ptr *QTreeView) RowsInsertedDefault(parent core.QModelIndex_ITF, start int, end int) {
	defer qt.Recovering("QTreeView::rowsInserted")

	if ptr.Pointer() != nil {
		C.QTreeView_RowsInsertedDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(start), C.int(end))
	}
}

func (ptr *QTreeView) ConnectScrollContentsBy(f func(dx int, dy int)) {
	defer qt.Recovering("connect QTreeView::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "scrollContentsBy", f)
	}
}

func (ptr *QTreeView) DisconnectScrollContentsBy() {
	defer qt.Recovering("disconnect QTreeView::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "scrollContentsBy")
	}
}

//export callbackQTreeViewScrollContentsBy
func callbackQTreeViewScrollContentsBy(ptr unsafe.Pointer, ptrName *C.char, dx C.int, dy C.int) {
	defer qt.Recovering("callback QTreeView::scrollContentsBy")

	if signal := qt.GetSignal(C.GoString(ptrName), "scrollContentsBy"); signal != nil {
		signal.(func(int, int))(int(dx), int(dy))
	} else {
		NewQTreeViewFromPointer(ptr).ScrollContentsByDefault(int(dx), int(dy))
	}
}

func (ptr *QTreeView) ScrollContentsBy(dx int, dy int) {
	defer qt.Recovering("QTreeView::scrollContentsBy")

	if ptr.Pointer() != nil {
		C.QTreeView_ScrollContentsBy(ptr.Pointer(), C.int(dx), C.int(dy))
	}
}

func (ptr *QTreeView) ScrollContentsByDefault(dx int, dy int) {
	defer qt.Recovering("QTreeView::scrollContentsBy")

	if ptr.Pointer() != nil {
		C.QTreeView_ScrollContentsByDefault(ptr.Pointer(), C.int(dx), C.int(dy))
	}
}

func (ptr *QTreeView) ConnectScrollTo(f func(index *core.QModelIndex, hint QAbstractItemView__ScrollHint)) {
	defer qt.Recovering("connect QTreeView::scrollTo")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "scrollTo", f)
	}
}

func (ptr *QTreeView) DisconnectScrollTo() {
	defer qt.Recovering("disconnect QTreeView::scrollTo")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "scrollTo")
	}
}

//export callbackQTreeViewScrollTo
func callbackQTreeViewScrollTo(ptr unsafe.Pointer, ptrName *C.char, index unsafe.Pointer, hint C.int) {
	defer qt.Recovering("callback QTreeView::scrollTo")

	if signal := qt.GetSignal(C.GoString(ptrName), "scrollTo"); signal != nil {
		signal.(func(*core.QModelIndex, QAbstractItemView__ScrollHint))(core.NewQModelIndexFromPointer(index), QAbstractItemView__ScrollHint(hint))
	} else {
		NewQTreeViewFromPointer(ptr).ScrollToDefault(core.NewQModelIndexFromPointer(index), QAbstractItemView__ScrollHint(hint))
	}
}

func (ptr *QTreeView) ScrollTo(index core.QModelIndex_ITF, hint QAbstractItemView__ScrollHint) {
	defer qt.Recovering("QTreeView::scrollTo")

	if ptr.Pointer() != nil {
		C.QTreeView_ScrollTo(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(hint))
	}
}

func (ptr *QTreeView) ScrollToDefault(index core.QModelIndex_ITF, hint QAbstractItemView__ScrollHint) {
	defer qt.Recovering("QTreeView::scrollTo")

	if ptr.Pointer() != nil {
		C.QTreeView_ScrollToDefault(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(hint))
	}
}

func (ptr *QTreeView) ConnectSelectAll(f func()) {
	defer qt.Recovering("connect QTreeView::selectAll")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "selectAll", f)
	}
}

func (ptr *QTreeView) DisconnectSelectAll() {
	defer qt.Recovering("disconnect QTreeView::selectAll")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "selectAll")
	}
}

//export callbackQTreeViewSelectAll
func callbackQTreeViewSelectAll(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QTreeView::selectAll")

	if signal := qt.GetSignal(C.GoString(ptrName), "selectAll"); signal != nil {
		signal.(func())()
	} else {
		NewQTreeViewFromPointer(ptr).SelectAllDefault()
	}
}

func (ptr *QTreeView) SelectAll() {
	defer qt.Recovering("QTreeView::selectAll")

	if ptr.Pointer() != nil {
		C.QTreeView_SelectAll(ptr.Pointer())
	}
}

func (ptr *QTreeView) SelectAllDefault() {
	defer qt.Recovering("QTreeView::selectAll")

	if ptr.Pointer() != nil {
		C.QTreeView_SelectAllDefault(ptr.Pointer())
	}
}

func (ptr *QTreeView) SetColumnHidden(column int, hide bool) {
	defer qt.Recovering("QTreeView::setColumnHidden")

	if ptr.Pointer() != nil {
		C.QTreeView_SetColumnHidden(ptr.Pointer(), C.int(column), C.int(qt.GoBoolToInt(hide)))
	}
}

func (ptr *QTreeView) SetColumnWidth(column int, width int) {
	defer qt.Recovering("QTreeView::setColumnWidth")

	if ptr.Pointer() != nil {
		C.QTreeView_SetColumnWidth(ptr.Pointer(), C.int(column), C.int(width))
	}
}

func (ptr *QTreeView) SetExpanded(index core.QModelIndex_ITF, expanded bool) {
	defer qt.Recovering("QTreeView::setExpanded")

	if ptr.Pointer() != nil {
		C.QTreeView_SetExpanded(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(qt.GoBoolToInt(expanded)))
	}
}

func (ptr *QTreeView) SetFirstColumnSpanned(row int, parent core.QModelIndex_ITF, span bool) {
	defer qt.Recovering("QTreeView::setFirstColumnSpanned")

	if ptr.Pointer() != nil {
		C.QTreeView_SetFirstColumnSpanned(ptr.Pointer(), C.int(row), core.PointerFromQModelIndex(parent), C.int(qt.GoBoolToInt(span)))
	}
}

func (ptr *QTreeView) SetHeader(header QHeaderView_ITF) {
	defer qt.Recovering("QTreeView::setHeader")

	if ptr.Pointer() != nil {
		C.QTreeView_SetHeader(ptr.Pointer(), PointerFromQHeaderView(header))
	}
}

func (ptr *QTreeView) ConnectSetModel(f func(model *core.QAbstractItemModel)) {
	defer qt.Recovering("connect QTreeView::setModel")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setModel", f)
	}
}

func (ptr *QTreeView) DisconnectSetModel() {
	defer qt.Recovering("disconnect QTreeView::setModel")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setModel")
	}
}

//export callbackQTreeViewSetModel
func callbackQTreeViewSetModel(ptr unsafe.Pointer, ptrName *C.char, model unsafe.Pointer) {
	defer qt.Recovering("callback QTreeView::setModel")

	if signal := qt.GetSignal(C.GoString(ptrName), "setModel"); signal != nil {
		signal.(func(*core.QAbstractItemModel))(core.NewQAbstractItemModelFromPointer(model))
	} else {
		NewQTreeViewFromPointer(ptr).SetModelDefault(core.NewQAbstractItemModelFromPointer(model))
	}
}

func (ptr *QTreeView) SetModel(model core.QAbstractItemModel_ITF) {
	defer qt.Recovering("QTreeView::setModel")

	if ptr.Pointer() != nil {
		C.QTreeView_SetModel(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QTreeView) SetModelDefault(model core.QAbstractItemModel_ITF) {
	defer qt.Recovering("QTreeView::setModel")

	if ptr.Pointer() != nil {
		C.QTreeView_SetModelDefault(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QTreeView) ConnectSetRootIndex(f func(index *core.QModelIndex)) {
	defer qt.Recovering("connect QTreeView::setRootIndex")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setRootIndex", f)
	}
}

func (ptr *QTreeView) DisconnectSetRootIndex() {
	defer qt.Recovering("disconnect QTreeView::setRootIndex")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setRootIndex")
	}
}

//export callbackQTreeViewSetRootIndex
func callbackQTreeViewSetRootIndex(ptr unsafe.Pointer, ptrName *C.char, index unsafe.Pointer) {
	defer qt.Recovering("callback QTreeView::setRootIndex")

	if signal := qt.GetSignal(C.GoString(ptrName), "setRootIndex"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	} else {
		NewQTreeViewFromPointer(ptr).SetRootIndexDefault(core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QTreeView) SetRootIndex(index core.QModelIndex_ITF) {
	defer qt.Recovering("QTreeView::setRootIndex")

	if ptr.Pointer() != nil {
		C.QTreeView_SetRootIndex(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QTreeView) SetRootIndexDefault(index core.QModelIndex_ITF) {
	defer qt.Recovering("QTreeView::setRootIndex")

	if ptr.Pointer() != nil {
		C.QTreeView_SetRootIndexDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QTreeView) SetRowHidden(row int, parent core.QModelIndex_ITF, hide bool) {
	defer qt.Recovering("QTreeView::setRowHidden")

	if ptr.Pointer() != nil {
		C.QTreeView_SetRowHidden(ptr.Pointer(), C.int(row), core.PointerFromQModelIndex(parent), C.int(qt.GoBoolToInt(hide)))
	}
}

func (ptr *QTreeView) ConnectSetSelection(f func(rect *core.QRect, command core.QItemSelectionModel__SelectionFlag)) {
	defer qt.Recovering("connect QTreeView::setSelection")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setSelection", f)
	}
}

func (ptr *QTreeView) DisconnectSetSelection() {
	defer qt.Recovering("disconnect QTreeView::setSelection")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setSelection")
	}
}

//export callbackQTreeViewSetSelection
func callbackQTreeViewSetSelection(ptr unsafe.Pointer, ptrName *C.char, rect unsafe.Pointer, command C.int) {
	defer qt.Recovering("callback QTreeView::setSelection")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSelection"); signal != nil {
		signal.(func(*core.QRect, core.QItemSelectionModel__SelectionFlag))(core.NewQRectFromPointer(rect), core.QItemSelectionModel__SelectionFlag(command))
	} else {
		NewQTreeViewFromPointer(ptr).SetSelectionDefault(core.NewQRectFromPointer(rect), core.QItemSelectionModel__SelectionFlag(command))
	}
}

func (ptr *QTreeView) SetSelection(rect core.QRect_ITF, command core.QItemSelectionModel__SelectionFlag) {
	defer qt.Recovering("QTreeView::setSelection")

	if ptr.Pointer() != nil {
		C.QTreeView_SetSelection(ptr.Pointer(), core.PointerFromQRect(rect), C.int(command))
	}
}

func (ptr *QTreeView) SetSelectionDefault(rect core.QRect_ITF, command core.QItemSelectionModel__SelectionFlag) {
	defer qt.Recovering("QTreeView::setSelection")

	if ptr.Pointer() != nil {
		C.QTreeView_SetSelectionDefault(ptr.Pointer(), core.PointerFromQRect(rect), C.int(command))
	}
}

func (ptr *QTreeView) ConnectSetSelectionModel(f func(selectionModel *core.QItemSelectionModel)) {
	defer qt.Recovering("connect QTreeView::setSelectionModel")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setSelectionModel", f)
	}
}

func (ptr *QTreeView) DisconnectSetSelectionModel() {
	defer qt.Recovering("disconnect QTreeView::setSelectionModel")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setSelectionModel")
	}
}

//export callbackQTreeViewSetSelectionModel
func callbackQTreeViewSetSelectionModel(ptr unsafe.Pointer, ptrName *C.char, selectionModel unsafe.Pointer) {
	defer qt.Recovering("callback QTreeView::setSelectionModel")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSelectionModel"); signal != nil {
		signal.(func(*core.QItemSelectionModel))(core.NewQItemSelectionModelFromPointer(selectionModel))
	} else {
		NewQTreeViewFromPointer(ptr).SetSelectionModelDefault(core.NewQItemSelectionModelFromPointer(selectionModel))
	}
}

func (ptr *QTreeView) SetSelectionModel(selectionModel core.QItemSelectionModel_ITF) {
	defer qt.Recovering("QTreeView::setSelectionModel")

	if ptr.Pointer() != nil {
		C.QTreeView_SetSelectionModel(ptr.Pointer(), core.PointerFromQItemSelectionModel(selectionModel))
	}
}

func (ptr *QTreeView) SetSelectionModelDefault(selectionModel core.QItemSelectionModel_ITF) {
	defer qt.Recovering("QTreeView::setSelectionModel")

	if ptr.Pointer() != nil {
		C.QTreeView_SetSelectionModelDefault(ptr.Pointer(), core.PointerFromQItemSelectionModel(selectionModel))
	}
}

func (ptr *QTreeView) SetTreePosition(index int) {
	defer qt.Recovering("QTreeView::setTreePosition")

	if ptr.Pointer() != nil {
		C.QTreeView_SetTreePosition(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QTreeView) ShowColumn(column int) {
	defer qt.Recovering("QTreeView::showColumn")

	if ptr.Pointer() != nil {
		C.QTreeView_ShowColumn(ptr.Pointer(), C.int(column))
	}
}

func (ptr *QTreeView) SizeHintForColumn(column int) int {
	defer qt.Recovering("QTreeView::sizeHintForColumn")

	if ptr.Pointer() != nil {
		return int(C.QTreeView_SizeHintForColumn(ptr.Pointer(), C.int(column)))
	}
	return 0
}

func (ptr *QTreeView) SortByColumn(column int, order core.Qt__SortOrder) {
	defer qt.Recovering("QTreeView::sortByColumn")

	if ptr.Pointer() != nil {
		C.QTreeView_SortByColumn(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QTreeView) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QTreeView::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QTreeView) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QTreeView::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQTreeViewTimerEvent
func callbackQTreeViewTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTreeView::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQTreeViewFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QTreeView) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QTreeView::timerEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QTreeView) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QTreeView::timerEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QTreeView) TreePosition() int {
	defer qt.Recovering("QTreeView::treePosition")

	if ptr.Pointer() != nil {
		return int(C.QTreeView_TreePosition(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTreeView) ConnectUpdateGeometries(f func()) {
	defer qt.Recovering("connect QTreeView::updateGeometries")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updateGeometries", f)
	}
}

func (ptr *QTreeView) DisconnectUpdateGeometries() {
	defer qt.Recovering("disconnect QTreeView::updateGeometries")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updateGeometries")
	}
}

//export callbackQTreeViewUpdateGeometries
func callbackQTreeViewUpdateGeometries(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QTreeView::updateGeometries")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateGeometries"); signal != nil {
		signal.(func())()
	} else {
		NewQTreeViewFromPointer(ptr).UpdateGeometriesDefault()
	}
}

func (ptr *QTreeView) UpdateGeometries() {
	defer qt.Recovering("QTreeView::updateGeometries")

	if ptr.Pointer() != nil {
		C.QTreeView_UpdateGeometries(ptr.Pointer())
	}
}

func (ptr *QTreeView) UpdateGeometriesDefault() {
	defer qt.Recovering("QTreeView::updateGeometries")

	if ptr.Pointer() != nil {
		C.QTreeView_UpdateGeometriesDefault(ptr.Pointer())
	}
}

func (ptr *QTreeView) VerticalOffset() int {
	defer qt.Recovering("QTreeView::verticalOffset")

	if ptr.Pointer() != nil {
		return int(C.QTreeView_VerticalOffset(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTreeView) ViewportEvent(event core.QEvent_ITF) bool {
	defer qt.Recovering("QTreeView::viewportEvent")

	if ptr.Pointer() != nil {
		return C.QTreeView_ViewportEvent(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QTreeView) ViewportSizeHint() *core.QSize {
	defer qt.Recovering("QTreeView::viewportSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QTreeView_ViewportSizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTreeView) VisualRect(index core.QModelIndex_ITF) *core.QRect {
	defer qt.Recovering("QTreeView::visualRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QTreeView_VisualRect(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return nil
}

func (ptr *QTreeView) VisualRegionForSelection(selection core.QItemSelection_ITF) *gui.QRegion {
	defer qt.Recovering("QTreeView::visualRegionForSelection")

	if ptr.Pointer() != nil {
		return gui.NewQRegionFromPointer(C.QTreeView_VisualRegionForSelection(ptr.Pointer(), core.PointerFromQItemSelection(selection)))
	}
	return nil
}

func (ptr *QTreeView) DestroyQTreeView() {
	defer qt.Recovering("QTreeView::~QTreeView")

	if ptr.Pointer() != nil {
		C.QTreeView_DestroyQTreeView(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QTreeView) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QTreeView::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QTreeView) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QTreeView::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQTreeViewDragLeaveEvent
func callbackQTreeViewDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTreeView::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQTreeViewFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QTreeView) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QTreeView::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QTreeView) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QTreeView::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QTreeView) ConnectCloseEditor(f func(editor *QWidget, hint QAbstractItemDelegate__EndEditHint)) {
	defer qt.Recovering("connect QTreeView::closeEditor")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEditor", f)
	}
}

func (ptr *QTreeView) DisconnectCloseEditor() {
	defer qt.Recovering("disconnect QTreeView::closeEditor")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEditor")
	}
}

//export callbackQTreeViewCloseEditor
func callbackQTreeViewCloseEditor(ptr unsafe.Pointer, ptrName *C.char, editor unsafe.Pointer, hint C.int) bool {
	defer qt.Recovering("callback QTreeView::closeEditor")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEditor"); signal != nil {
		signal.(func(*QWidget, QAbstractItemDelegate__EndEditHint))(NewQWidgetFromPointer(editor), QAbstractItemDelegate__EndEditHint(hint))
		return true
	}
	return false

}

func (ptr *QTreeView) CloseEditor(editor QWidget_ITF, hint QAbstractItemDelegate__EndEditHint) {
	defer qt.Recovering("QTreeView::closeEditor")

	if ptr.Pointer() != nil {
		C.QTreeView_CloseEditor(ptr.Pointer(), PointerFromQWidget(editor), C.int(hint))
	}
}

func (ptr *QTreeView) CloseEditorDefault(editor QWidget_ITF, hint QAbstractItemDelegate__EndEditHint) {
	defer qt.Recovering("QTreeView::closeEditor")

	if ptr.Pointer() != nil {
		C.QTreeView_CloseEditorDefault(ptr.Pointer(), PointerFromQWidget(editor), C.int(hint))
	}
}

func (ptr *QTreeView) ConnectCommitData(f func(editor *QWidget)) {
	defer qt.Recovering("connect QTreeView::commitData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "commitData", f)
	}
}

func (ptr *QTreeView) DisconnectCommitData() {
	defer qt.Recovering("disconnect QTreeView::commitData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "commitData")
	}
}

//export callbackQTreeViewCommitData
func callbackQTreeViewCommitData(ptr unsafe.Pointer, ptrName *C.char, editor unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeView::commitData")

	if signal := qt.GetSignal(C.GoString(ptrName), "commitData"); signal != nil {
		signal.(func(*QWidget))(NewQWidgetFromPointer(editor))
		return true
	}
	return false

}

func (ptr *QTreeView) CommitData(editor QWidget_ITF) {
	defer qt.Recovering("QTreeView::commitData")

	if ptr.Pointer() != nil {
		C.QTreeView_CommitData(ptr.Pointer(), PointerFromQWidget(editor))
	}
}

func (ptr *QTreeView) CommitDataDefault(editor QWidget_ITF) {
	defer qt.Recovering("QTreeView::commitData")

	if ptr.Pointer() != nil {
		C.QTreeView_CommitDataDefault(ptr.Pointer(), PointerFromQWidget(editor))
	}
}

func (ptr *QTreeView) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QTreeView::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QTreeView) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QTreeView::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQTreeViewDragEnterEvent
func callbackQTreeViewDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTreeView::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQTreeViewFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QTreeView) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QTreeView::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QTreeView) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QTreeView::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QTreeView) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QTreeView::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QTreeView) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QTreeView::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQTreeViewDropEvent
func callbackQTreeViewDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTreeView::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQTreeViewFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QTreeView) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QTreeView::dropEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QTreeView) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QTreeView::dropEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QTreeView) ConnectEditorDestroyed(f func(editor *core.QObject)) {
	defer qt.Recovering("connect QTreeView::editorDestroyed")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "editorDestroyed", f)
	}
}

func (ptr *QTreeView) DisconnectEditorDestroyed() {
	defer qt.Recovering("disconnect QTreeView::editorDestroyed")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "editorDestroyed")
	}
}

//export callbackQTreeViewEditorDestroyed
func callbackQTreeViewEditorDestroyed(ptr unsafe.Pointer, ptrName *C.char, editor unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeView::editorDestroyed")

	if signal := qt.GetSignal(C.GoString(ptrName), "editorDestroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(editor))
		return true
	}
	return false

}

func (ptr *QTreeView) EditorDestroyed(editor core.QObject_ITF) {
	defer qt.Recovering("QTreeView::editorDestroyed")

	if ptr.Pointer() != nil {
		C.QTreeView_EditorDestroyed(ptr.Pointer(), core.PointerFromQObject(editor))
	}
}

func (ptr *QTreeView) EditorDestroyedDefault(editor core.QObject_ITF) {
	defer qt.Recovering("QTreeView::editorDestroyed")

	if ptr.Pointer() != nil {
		C.QTreeView_EditorDestroyedDefault(ptr.Pointer(), core.PointerFromQObject(editor))
	}
}

func (ptr *QTreeView) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QTreeView::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QTreeView) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QTreeView::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQTreeViewFocusInEvent
func callbackQTreeViewFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTreeView::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQTreeViewFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QTreeView) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QTreeView::focusInEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QTreeView) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QTreeView::focusInEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QTreeView) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QTreeView::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QTreeView) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QTreeView::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQTreeViewFocusOutEvent
func callbackQTreeViewFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTreeView::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQTreeViewFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QTreeView) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QTreeView::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QTreeView) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QTreeView::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QTreeView) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QTreeView::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QTreeView) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QTreeView::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQTreeViewInputMethodEvent
func callbackQTreeViewInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTreeView::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQTreeViewFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QTreeView) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QTreeView::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QTreeView) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QTreeView::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QTreeView) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QTreeView::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QTreeView) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QTreeView::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQTreeViewResizeEvent
func callbackQTreeViewResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTreeView::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQTreeViewFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QTreeView) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QTreeView::resizeEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QTreeView) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QTreeView::resizeEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QTreeView) ConnectStartDrag(f func(supportedActions core.Qt__DropAction)) {
	defer qt.Recovering("connect QTreeView::startDrag")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "startDrag", f)
	}
}

func (ptr *QTreeView) DisconnectStartDrag() {
	defer qt.Recovering("disconnect QTreeView::startDrag")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "startDrag")
	}
}

//export callbackQTreeViewStartDrag
func callbackQTreeViewStartDrag(ptr unsafe.Pointer, ptrName *C.char, supportedActions C.int) {
	defer qt.Recovering("callback QTreeView::startDrag")

	if signal := qt.GetSignal(C.GoString(ptrName), "startDrag"); signal != nil {
		signal.(func(core.Qt__DropAction))(core.Qt__DropAction(supportedActions))
	} else {
		NewQTreeViewFromPointer(ptr).StartDragDefault(core.Qt__DropAction(supportedActions))
	}
}

func (ptr *QTreeView) StartDrag(supportedActions core.Qt__DropAction) {
	defer qt.Recovering("QTreeView::startDrag")

	if ptr.Pointer() != nil {
		C.QTreeView_StartDrag(ptr.Pointer(), C.int(supportedActions))
	}
}

func (ptr *QTreeView) StartDragDefault(supportedActions core.Qt__DropAction) {
	defer qt.Recovering("QTreeView::startDrag")

	if ptr.Pointer() != nil {
		C.QTreeView_StartDragDefault(ptr.Pointer(), C.int(supportedActions))
	}
}

func (ptr *QTreeView) ConnectContextMenuEvent(f func(e *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QTreeView::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QTreeView) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QTreeView::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQTreeViewContextMenuEvent
func callbackQTreeViewContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QTreeView::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
	} else {
		NewQTreeViewFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(e))
	}
}

func (ptr *QTreeView) ContextMenuEvent(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QTreeView::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

func (ptr *QTreeView) ContextMenuEventDefault(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QTreeView::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

func (ptr *QTreeView) ConnectSetupViewport(f func(viewport *QWidget)) {
	defer qt.Recovering("connect QTreeView::setupViewport")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setupViewport", f)
	}
}

func (ptr *QTreeView) DisconnectSetupViewport() {
	defer qt.Recovering("disconnect QTreeView::setupViewport")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setupViewport")
	}
}

//export callbackQTreeViewSetupViewport
func callbackQTreeViewSetupViewport(ptr unsafe.Pointer, ptrName *C.char, viewport unsafe.Pointer) {
	defer qt.Recovering("callback QTreeView::setupViewport")

	if signal := qt.GetSignal(C.GoString(ptrName), "setupViewport"); signal != nil {
		signal.(func(*QWidget))(NewQWidgetFromPointer(viewport))
	} else {
		NewQTreeViewFromPointer(ptr).SetupViewportDefault(NewQWidgetFromPointer(viewport))
	}
}

func (ptr *QTreeView) SetupViewport(viewport QWidget_ITF) {
	defer qt.Recovering("QTreeView::setupViewport")

	if ptr.Pointer() != nil {
		C.QTreeView_SetupViewport(ptr.Pointer(), PointerFromQWidget(viewport))
	}
}

func (ptr *QTreeView) SetupViewportDefault(viewport QWidget_ITF) {
	defer qt.Recovering("QTreeView::setupViewport")

	if ptr.Pointer() != nil {
		C.QTreeView_SetupViewportDefault(ptr.Pointer(), PointerFromQWidget(viewport))
	}
}

func (ptr *QTreeView) ConnectWheelEvent(f func(e *gui.QWheelEvent)) {
	defer qt.Recovering("connect QTreeView::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QTreeView) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QTreeView::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQTreeViewWheelEvent
func callbackQTreeViewWheelEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QTreeView::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
	} else {
		NewQTreeViewFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(e))
	}
}

func (ptr *QTreeView) WheelEvent(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QTreeView::wheelEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QTreeView) WheelEventDefault(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QTreeView::wheelEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QTreeView) ConnectChangeEvent(f func(ev *core.QEvent)) {
	defer qt.Recovering("connect QTreeView::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QTreeView) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QTreeView::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQTreeViewChangeEvent
func callbackQTreeViewChangeEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QTreeView::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
	} else {
		NewQTreeViewFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(ev))
	}
}

func (ptr *QTreeView) ChangeEvent(ev core.QEvent_ITF) {
	defer qt.Recovering("QTreeView::changeEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QTreeView) ChangeEventDefault(ev core.QEvent_ITF) {
	defer qt.Recovering("QTreeView::changeEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QTreeView) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QTreeView::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QTreeView) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QTreeView::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQTreeViewActionEvent
func callbackQTreeViewActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTreeView::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQTreeViewFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QTreeView) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QTreeView::actionEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QTreeView) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QTreeView::actionEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QTreeView) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QTreeView::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QTreeView) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QTreeView::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQTreeViewEnterEvent
func callbackQTreeViewEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTreeView::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQTreeViewFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QTreeView) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QTreeView::enterEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QTreeView) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QTreeView::enterEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QTreeView) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QTreeView::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QTreeView) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QTreeView::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQTreeViewHideEvent
func callbackQTreeViewHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTreeView::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQTreeViewFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QTreeView) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QTreeView::hideEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QTreeView) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QTreeView::hideEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QTreeView) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QTreeView::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QTreeView) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QTreeView::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQTreeViewLeaveEvent
func callbackQTreeViewLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTreeView::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQTreeViewFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QTreeView) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QTreeView::leaveEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QTreeView) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QTreeView::leaveEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QTreeView) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QTreeView::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QTreeView) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QTreeView::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQTreeViewMoveEvent
func callbackQTreeViewMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTreeView::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQTreeViewFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QTreeView) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QTreeView::moveEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QTreeView) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QTreeView::moveEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QTreeView) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QTreeView::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QTreeView) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QTreeView::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQTreeViewSetVisible
func callbackQTreeViewSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QTreeView::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QTreeView) SetVisible(visible bool) {
	defer qt.Recovering("QTreeView::setVisible")

	if ptr.Pointer() != nil {
		C.QTreeView_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QTreeView) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QTreeView::setVisible")

	if ptr.Pointer() != nil {
		C.QTreeView_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QTreeView) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QTreeView::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QTreeView) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QTreeView::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQTreeViewShowEvent
func callbackQTreeViewShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTreeView::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQTreeViewFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QTreeView) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QTreeView::showEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QTreeView) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QTreeView::showEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QTreeView) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QTreeView::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QTreeView) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QTreeView::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQTreeViewCloseEvent
func callbackQTreeViewCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTreeView::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQTreeViewFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QTreeView) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QTreeView::closeEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QTreeView) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QTreeView::closeEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QTreeView) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QTreeView::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QTreeView) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QTreeView::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQTreeViewInitPainter
func callbackQTreeViewInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QTreeView::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQTreeViewFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QTreeView) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QTreeView::initPainter")

	if ptr.Pointer() != nil {
		C.QTreeView_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QTreeView) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QTreeView::initPainter")

	if ptr.Pointer() != nil {
		C.QTreeView_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QTreeView) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QTreeView::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QTreeView) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QTreeView::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQTreeViewKeyReleaseEvent
func callbackQTreeViewKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTreeView::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQTreeViewFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QTreeView) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QTreeView::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QTreeView) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QTreeView::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QTreeView) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QTreeView::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QTreeView) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QTreeView::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQTreeViewTabletEvent
func callbackQTreeViewTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTreeView::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQTreeViewFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QTreeView) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QTreeView::tabletEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QTreeView) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QTreeView::tabletEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QTreeView) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QTreeView::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QTreeView) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QTreeView::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQTreeViewChildEvent
func callbackQTreeViewChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTreeView::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQTreeViewFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QTreeView) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QTreeView::childEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QTreeView) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QTreeView::childEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QTreeView) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QTreeView::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QTreeView) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QTreeView::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQTreeViewCustomEvent
func callbackQTreeViewCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTreeView::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQTreeViewFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QTreeView) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QTreeView::customEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QTreeView) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QTreeView::customEvent")

	if ptr.Pointer() != nil {
		C.QTreeView_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
