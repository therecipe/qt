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
func callbackQTreeViewCollapsed(ptrName *C.char, index unsafe.Pointer) {
	defer qt.Recovering("callback QTreeView::collapsed")

	if signal := qt.GetSignal(C.GoString(ptrName), "collapsed"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
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
func callbackQTreeViewCurrentChanged(ptrName *C.char, current unsafe.Pointer, previous unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeView::currentChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "currentChanged"); signal != nil {
		signal.(func(*core.QModelIndex, *core.QModelIndex))(core.NewQModelIndexFromPointer(current), core.NewQModelIndexFromPointer(previous))
		return true
	}
	return false

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
func callbackQTreeViewDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeView::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

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
func callbackQTreeViewDrawBranches(ptrName *C.char, painter unsafe.Pointer, rect unsafe.Pointer, index unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeView::drawBranches")

	if signal := qt.GetSignal(C.GoString(ptrName), "drawBranches"); signal != nil {
		signal.(func(*gui.QPainter, *core.QRect, *core.QModelIndex))(gui.NewQPainterFromPointer(painter), core.NewQRectFromPointer(rect), core.NewQModelIndexFromPointer(index))
		return true
	}
	return false

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
func callbackQTreeViewExpanded(ptrName *C.char, index unsafe.Pointer) {
	defer qt.Recovering("callback QTreeView::expanded")

	if signal := qt.GetSignal(C.GoString(ptrName), "expanded"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
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
func callbackQTreeViewKeyPressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeView::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

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
func callbackQTreeViewKeyboardSearch(ptrName *C.char, search *C.char) bool {
	defer qt.Recovering("callback QTreeView::keyboardSearch")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyboardSearch"); signal != nil {
		signal.(func(string))(C.GoString(search))
		return true
	}
	return false

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
func callbackQTreeViewMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeView::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQTreeViewMouseMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeView::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQTreeViewMousePressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeView::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQTreeViewMouseReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeView::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQTreeViewPaintEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeView::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
		return true
	}
	return false

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
func callbackQTreeViewReset(ptrName *C.char) bool {
	defer qt.Recovering("callback QTreeView::reset")

	if signal := qt.GetSignal(C.GoString(ptrName), "reset"); signal != nil {
		signal.(func())()
		return true
	}
	return false

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
func callbackQTreeViewRowsAboutToBeRemoved(ptrName *C.char, parent unsafe.Pointer, start C.int, end C.int) bool {
	defer qt.Recovering("callback QTreeView::rowsAboutToBeRemoved")

	if signal := qt.GetSignal(C.GoString(ptrName), "rowsAboutToBeRemoved"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(start), int(end))
		return true
	}
	return false

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
func callbackQTreeViewRowsInserted(ptrName *C.char, parent unsafe.Pointer, start C.int, end C.int) bool {
	defer qt.Recovering("callback QTreeView::rowsInserted")

	if signal := qt.GetSignal(C.GoString(ptrName), "rowsInserted"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(start), int(end))
		return true
	}
	return false

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
func callbackQTreeViewScrollContentsBy(ptrName *C.char, dx C.int, dy C.int) bool {
	defer qt.Recovering("callback QTreeView::scrollContentsBy")

	if signal := qt.GetSignal(C.GoString(ptrName), "scrollContentsBy"); signal != nil {
		signal.(func(int, int))(int(dx), int(dy))
		return true
	}
	return false

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
func callbackQTreeViewScrollTo(ptrName *C.char, index unsafe.Pointer, hint C.int) bool {
	defer qt.Recovering("callback QTreeView::scrollTo")

	if signal := qt.GetSignal(C.GoString(ptrName), "scrollTo"); signal != nil {
		signal.(func(*core.QModelIndex, QAbstractItemView__ScrollHint))(core.NewQModelIndexFromPointer(index), QAbstractItemView__ScrollHint(hint))
		return true
	}
	return false

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
func callbackQTreeViewSelectAll(ptrName *C.char) bool {
	defer qt.Recovering("callback QTreeView::selectAll")

	if signal := qt.GetSignal(C.GoString(ptrName), "selectAll"); signal != nil {
		signal.(func())()
		return true
	}
	return false

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
func callbackQTreeViewSetModel(ptrName *C.char, model unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeView::setModel")

	if signal := qt.GetSignal(C.GoString(ptrName), "setModel"); signal != nil {
		signal.(func(*core.QAbstractItemModel))(core.NewQAbstractItemModelFromPointer(model))
		return true
	}
	return false

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
func callbackQTreeViewSetRootIndex(ptrName *C.char, index unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeView::setRootIndex")

	if signal := qt.GetSignal(C.GoString(ptrName), "setRootIndex"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
		return true
	}
	return false

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
func callbackQTreeViewSetSelection(ptrName *C.char, rect unsafe.Pointer, command C.int) bool {
	defer qt.Recovering("callback QTreeView::setSelection")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSelection"); signal != nil {
		signal.(func(*core.QRect, core.QItemSelectionModel__SelectionFlag))(core.NewQRectFromPointer(rect), core.QItemSelectionModel__SelectionFlag(command))
		return true
	}
	return false

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
func callbackQTreeViewSetSelectionModel(ptrName *C.char, selectionModel unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeView::setSelectionModel")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSelectionModel"); signal != nil {
		signal.(func(*core.QItemSelectionModel))(core.NewQItemSelectionModelFromPointer(selectionModel))
		return true
	}
	return false

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
func callbackQTreeViewTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeView::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

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
func callbackQTreeViewUpdateGeometries(ptrName *C.char) bool {
	defer qt.Recovering("callback QTreeView::updateGeometries")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateGeometries"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QTreeView) VisualRect(index core.QModelIndex_ITF) *core.QRect {
	defer qt.Recovering("QTreeView::visualRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QTreeView_VisualRect(ptr.Pointer(), core.PointerFromQModelIndex(index)))
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
func callbackQTreeViewDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeView::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

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
func callbackQTreeViewCloseEditor(ptrName *C.char, editor unsafe.Pointer, hint C.int) bool {
	defer qt.Recovering("callback QTreeView::closeEditor")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEditor"); signal != nil {
		signal.(func(*QWidget, QAbstractItemDelegate__EndEditHint))(NewQWidgetFromPointer(editor), QAbstractItemDelegate__EndEditHint(hint))
		return true
	}
	return false

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
func callbackQTreeViewCommitData(ptrName *C.char, editor unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeView::commitData")

	if signal := qt.GetSignal(C.GoString(ptrName), "commitData"); signal != nil {
		signal.(func(*QWidget))(NewQWidgetFromPointer(editor))
		return true
	}
	return false

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
func callbackQTreeViewDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeView::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

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
func callbackQTreeViewDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeView::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

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
func callbackQTreeViewEditorDestroyed(ptrName *C.char, editor unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeView::editorDestroyed")

	if signal := qt.GetSignal(C.GoString(ptrName), "editorDestroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(editor))
		return true
	}
	return false

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
func callbackQTreeViewFocusInEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeView::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

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
func callbackQTreeViewFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeView::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

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
func callbackQTreeViewInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeView::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

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
func callbackQTreeViewResizeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeView::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
		return true
	}
	return false

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
func callbackQTreeViewStartDrag(ptrName *C.char, supportedActions C.int) bool {
	defer qt.Recovering("callback QTreeView::startDrag")

	if signal := qt.GetSignal(C.GoString(ptrName), "startDrag"); signal != nil {
		signal.(func(core.Qt__DropAction))(core.Qt__DropAction(supportedActions))
		return true
	}
	return false

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
func callbackQTreeViewContextMenuEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeView::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
		return true
	}
	return false

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
func callbackQTreeViewSetupViewport(ptrName *C.char, viewport unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeView::setupViewport")

	if signal := qt.GetSignal(C.GoString(ptrName), "setupViewport"); signal != nil {
		signal.(func(*QWidget))(NewQWidgetFromPointer(viewport))
		return true
	}
	return false

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
func callbackQTreeViewWheelEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeView::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
		return true
	}
	return false

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
func callbackQTreeViewChangeEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeView::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
		return true
	}
	return false

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
func callbackQTreeViewActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeView::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

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
func callbackQTreeViewEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeView::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

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
func callbackQTreeViewHideEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeView::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
		return true
	}
	return false

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
func callbackQTreeViewLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeView::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

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
func callbackQTreeViewMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeView::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

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
func callbackQTreeViewSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QTreeView::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

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
func callbackQTreeViewShowEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeView::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
		return true
	}
	return false

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
func callbackQTreeViewCloseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeView::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQTreeViewInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeView::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

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
func callbackQTreeViewKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeView::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

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
func callbackQTreeViewTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeView::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

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
func callbackQTreeViewChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeView::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

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
func callbackQTreeViewCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTreeView::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
