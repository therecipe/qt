package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QListView struct {
	QAbstractItemView
}

type QListView_ITF interface {
	QAbstractItemView_ITF
	QListView_PTR() *QListView
}

func PointerFromQListView(ptr QListView_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QListView_PTR().Pointer()
	}
	return nil
}

func NewQListViewFromPointer(ptr unsafe.Pointer) *QListView {
	var n = new(QListView)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QListView_") {
		n.SetObjectName("QListView_" + qt.Identifier())
	}
	return n
}

func (ptr *QListView) QListView_PTR() *QListView {
	return ptr
}

//QListView::Flow
type QListView__Flow int64

const (
	QListView__LeftToRight = QListView__Flow(0)
	QListView__TopToBottom = QListView__Flow(1)
)

//QListView::LayoutMode
type QListView__LayoutMode int64

const (
	QListView__SinglePass = QListView__LayoutMode(0)
	QListView__Batched    = QListView__LayoutMode(1)
)

//QListView::Movement
type QListView__Movement int64

const (
	QListView__Static = QListView__Movement(0)
	QListView__Free   = QListView__Movement(1)
	QListView__Snap   = QListView__Movement(2)
)

//QListView::ResizeMode
type QListView__ResizeMode int64

const (
	QListView__Fixed  = QListView__ResizeMode(0)
	QListView__Adjust = QListView__ResizeMode(1)
)

//QListView::ViewMode
type QListView__ViewMode int64

const (
	QListView__ListMode = QListView__ViewMode(0)
	QListView__IconMode = QListView__ViewMode(1)
)

func (ptr *QListView) BatchSize() int {
	defer qt.Recovering("QListView::batchSize")

	if ptr.Pointer() != nil {
		return int(C.QListView_BatchSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QListView) Flow() QListView__Flow {
	defer qt.Recovering("QListView::flow")

	if ptr.Pointer() != nil {
		return QListView__Flow(C.QListView_Flow(ptr.Pointer()))
	}
	return 0
}

func (ptr *QListView) GridSize() *core.QSize {
	defer qt.Recovering("QListView::gridSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QListView_GridSize(ptr.Pointer()))
	}
	return nil
}

func (ptr *QListView) IsSelectionRectVisible() bool {
	defer qt.Recovering("QListView::isSelectionRectVisible")

	if ptr.Pointer() != nil {
		return C.QListView_IsSelectionRectVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QListView) IsWrapping() bool {
	defer qt.Recovering("QListView::isWrapping")

	if ptr.Pointer() != nil {
		return C.QListView_IsWrapping(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QListView) LayoutMode() QListView__LayoutMode {
	defer qt.Recovering("QListView::layoutMode")

	if ptr.Pointer() != nil {
		return QListView__LayoutMode(C.QListView_LayoutMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QListView) ModelColumn() int {
	defer qt.Recovering("QListView::modelColumn")

	if ptr.Pointer() != nil {
		return int(C.QListView_ModelColumn(ptr.Pointer()))
	}
	return 0
}

func (ptr *QListView) Movement() QListView__Movement {
	defer qt.Recovering("QListView::movement")

	if ptr.Pointer() != nil {
		return QListView__Movement(C.QListView_Movement(ptr.Pointer()))
	}
	return 0
}

func (ptr *QListView) ResizeMode() QListView__ResizeMode {
	defer qt.Recovering("QListView::resizeMode")

	if ptr.Pointer() != nil {
		return QListView__ResizeMode(C.QListView_ResizeMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QListView) SetBatchSize(batchSize int) {
	defer qt.Recovering("QListView::setBatchSize")

	if ptr.Pointer() != nil {
		C.QListView_SetBatchSize(ptr.Pointer(), C.int(batchSize))
	}
}

func (ptr *QListView) SetFlow(flow QListView__Flow) {
	defer qt.Recovering("QListView::setFlow")

	if ptr.Pointer() != nil {
		C.QListView_SetFlow(ptr.Pointer(), C.int(flow))
	}
}

func (ptr *QListView) SetGridSize(size core.QSize_ITF) {
	defer qt.Recovering("QListView::setGridSize")

	if ptr.Pointer() != nil {
		C.QListView_SetGridSize(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QListView) SetLayoutMode(mode QListView__LayoutMode) {
	defer qt.Recovering("QListView::setLayoutMode")

	if ptr.Pointer() != nil {
		C.QListView_SetLayoutMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QListView) SetModelColumn(column int) {
	defer qt.Recovering("QListView::setModelColumn")

	if ptr.Pointer() != nil {
		C.QListView_SetModelColumn(ptr.Pointer(), C.int(column))
	}
}

func (ptr *QListView) SetMovement(movement QListView__Movement) {
	defer qt.Recovering("QListView::setMovement")

	if ptr.Pointer() != nil {
		C.QListView_SetMovement(ptr.Pointer(), C.int(movement))
	}
}

func (ptr *QListView) SetResizeMode(mode QListView__ResizeMode) {
	defer qt.Recovering("QListView::setResizeMode")

	if ptr.Pointer() != nil {
		C.QListView_SetResizeMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QListView) SetSelectionRectVisible(show bool) {
	defer qt.Recovering("QListView::setSelectionRectVisible")

	if ptr.Pointer() != nil {
		C.QListView_SetSelectionRectVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(show)))
	}
}

func (ptr *QListView) SetSpacing(space int) {
	defer qt.Recovering("QListView::setSpacing")

	if ptr.Pointer() != nil {
		C.QListView_SetSpacing(ptr.Pointer(), C.int(space))
	}
}

func (ptr *QListView) SetUniformItemSizes(enable bool) {
	defer qt.Recovering("QListView::setUniformItemSizes")

	if ptr.Pointer() != nil {
		C.QListView_SetUniformItemSizes(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QListView) SetViewMode(mode QListView__ViewMode) {
	defer qt.Recovering("QListView::setViewMode")

	if ptr.Pointer() != nil {
		C.QListView_SetViewMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QListView) SetWordWrap(on bool) {
	defer qt.Recovering("QListView::setWordWrap")

	if ptr.Pointer() != nil {
		C.QListView_SetWordWrap(ptr.Pointer(), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QListView) SetWrapping(enable bool) {
	defer qt.Recovering("QListView::setWrapping")

	if ptr.Pointer() != nil {
		C.QListView_SetWrapping(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QListView) Spacing() int {
	defer qt.Recovering("QListView::spacing")

	if ptr.Pointer() != nil {
		return int(C.QListView_Spacing(ptr.Pointer()))
	}
	return 0
}

func (ptr *QListView) UniformItemSizes() bool {
	defer qt.Recovering("QListView::uniformItemSizes")

	if ptr.Pointer() != nil {
		return C.QListView_UniformItemSizes(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QListView) ViewMode() QListView__ViewMode {
	defer qt.Recovering("QListView::viewMode")

	if ptr.Pointer() != nil {
		return QListView__ViewMode(C.QListView_ViewMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QListView) WordWrap() bool {
	defer qt.Recovering("QListView::wordWrap")

	if ptr.Pointer() != nil {
		return C.QListView_WordWrap(ptr.Pointer()) != 0
	}
	return false
}

func NewQListView(parent QWidget_ITF) *QListView {
	defer qt.Recovering("QListView::QListView")

	return NewQListViewFromPointer(C.QListView_NewQListView(PointerFromQWidget(parent)))
}

func (ptr *QListView) ClearPropertyFlags() {
	defer qt.Recovering("QListView::clearPropertyFlags")

	if ptr.Pointer() != nil {
		C.QListView_ClearPropertyFlags(ptr.Pointer())
	}
}

func (ptr *QListView) ConnectCurrentChanged(f func(current *core.QModelIndex, previous *core.QModelIndex)) {
	defer qt.Recovering("connect QListView::currentChanged")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "currentChanged", f)
	}
}

func (ptr *QListView) DisconnectCurrentChanged() {
	defer qt.Recovering("disconnect QListView::currentChanged")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "currentChanged")
	}
}

//export callbackQListViewCurrentChanged
func callbackQListViewCurrentChanged(ptr unsafe.Pointer, ptrName *C.char, current unsafe.Pointer, previous unsafe.Pointer) bool {
	defer qt.Recovering("callback QListView::currentChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "currentChanged"); signal != nil {
		signal.(func(*core.QModelIndex, *core.QModelIndex))(core.NewQModelIndexFromPointer(current), core.NewQModelIndexFromPointer(previous))
		return true
	}
	return false

}

func (ptr *QListView) CurrentChanged(current core.QModelIndex_ITF, previous core.QModelIndex_ITF) {
	defer qt.Recovering("QListView::currentChanged")

	if ptr.Pointer() != nil {
		C.QListView_CurrentChanged(ptr.Pointer(), core.PointerFromQModelIndex(current), core.PointerFromQModelIndex(previous))
	}
}

func (ptr *QListView) CurrentChangedDefault(current core.QModelIndex_ITF, previous core.QModelIndex_ITF) {
	defer qt.Recovering("QListView::currentChanged")

	if ptr.Pointer() != nil {
		C.QListView_CurrentChangedDefault(ptr.Pointer(), core.PointerFromQModelIndex(current), core.PointerFromQModelIndex(previous))
	}
}

func (ptr *QListView) ConnectDragLeaveEvent(f func(e *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QListView::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QListView) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QListView::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQListViewDragLeaveEvent
func callbackQListViewDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QListView::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(e))
	} else {
		NewQListViewFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(e))
	}
}

func (ptr *QListView) DragLeaveEvent(e gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QListView::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QListView_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(e))
	}
}

func (ptr *QListView) DragLeaveEventDefault(e gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QListView::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QListView_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(e))
	}
}

func (ptr *QListView) ConnectDragMoveEvent(f func(e *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QListView::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QListView) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QListView::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQListViewDragMoveEvent
func callbackQListViewDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QListView::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(e))
	} else {
		NewQListViewFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(e))
	}
}

func (ptr *QListView) DragMoveEvent(e gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QListView::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QListView_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(e))
	}
}

func (ptr *QListView) DragMoveEventDefault(e gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QListView::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QListView_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(e))
	}
}

func (ptr *QListView) ConnectDropEvent(f func(e *gui.QDropEvent)) {
	defer qt.Recovering("connect QListView::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QListView) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QListView::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQListViewDropEvent
func callbackQListViewDropEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QListView::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(e))
	} else {
		NewQListViewFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(e))
	}
}

func (ptr *QListView) DropEvent(e gui.QDropEvent_ITF) {
	defer qt.Recovering("QListView::dropEvent")

	if ptr.Pointer() != nil {
		C.QListView_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(e))
	}
}

func (ptr *QListView) DropEventDefault(e gui.QDropEvent_ITF) {
	defer qt.Recovering("QListView::dropEvent")

	if ptr.Pointer() != nil {
		C.QListView_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(e))
	}
}

func (ptr *QListView) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QListView::event")

	if ptr.Pointer() != nil {
		return C.QListView_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QListView) HorizontalOffset() int {
	defer qt.Recovering("QListView::horizontalOffset")

	if ptr.Pointer() != nil {
		return int(C.QListView_HorizontalOffset(ptr.Pointer()))
	}
	return 0
}

func (ptr *QListView) IndexAt(p core.QPoint_ITF) *core.QModelIndex {
	defer qt.Recovering("QListView::indexAt")

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QListView_IndexAt(ptr.Pointer(), core.PointerFromQPoint(p)))
	}
	return nil
}

func (ptr *QListView) IsIndexHidden(index core.QModelIndex_ITF) bool {
	defer qt.Recovering("QListView::isIndexHidden")

	if ptr.Pointer() != nil {
		return C.QListView_IsIndexHidden(ptr.Pointer(), core.PointerFromQModelIndex(index)) != 0
	}
	return false
}

func (ptr *QListView) IsRowHidden(row int) bool {
	defer qt.Recovering("QListView::isRowHidden")

	if ptr.Pointer() != nil {
		return C.QListView_IsRowHidden(ptr.Pointer(), C.int(row)) != 0
	}
	return false
}

func (ptr *QListView) ConnectMouseMoveEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QListView::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QListView) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QListView::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQListViewMouseMoveEvent
func callbackQListViewMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QListView::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQListViewFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QListView) MouseMoveEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QListView::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QListView_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QListView) MouseMoveEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QListView::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QListView_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QListView) ConnectMouseReleaseEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QListView::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QListView) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QListView::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQListViewMouseReleaseEvent
func callbackQListViewMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QListView::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQListViewFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QListView) MouseReleaseEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QListView::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QListView_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QListView) MouseReleaseEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QListView::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QListView_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QListView) MoveCursor(cursorAction QAbstractItemView__CursorAction, modifiers core.Qt__KeyboardModifier) *core.QModelIndex {
	defer qt.Recovering("QListView::moveCursor")

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QListView_MoveCursor(ptr.Pointer(), C.int(cursorAction), C.int(modifiers)))
	}
	return nil
}

func (ptr *QListView) ConnectPaintEvent(f func(e *gui.QPaintEvent)) {
	defer qt.Recovering("connect QListView::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QListView) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QListView::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQListViewPaintEvent
func callbackQListViewPaintEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QListView::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(e))
	} else {
		NewQListViewFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(e))
	}
}

func (ptr *QListView) PaintEvent(e gui.QPaintEvent_ITF) {
	defer qt.Recovering("QListView::paintEvent")

	if ptr.Pointer() != nil {
		C.QListView_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(e))
	}
}

func (ptr *QListView) PaintEventDefault(e gui.QPaintEvent_ITF) {
	defer qt.Recovering("QListView::paintEvent")

	if ptr.Pointer() != nil {
		C.QListView_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(e))
	}
}

func (ptr *QListView) ConnectResizeEvent(f func(e *gui.QResizeEvent)) {
	defer qt.Recovering("connect QListView::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QListView) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QListView::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQListViewResizeEvent
func callbackQListViewResizeEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QListView::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(e))
	} else {
		NewQListViewFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(e))
	}
}

func (ptr *QListView) ResizeEvent(e gui.QResizeEvent_ITF) {
	defer qt.Recovering("QListView::resizeEvent")

	if ptr.Pointer() != nil {
		C.QListView_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(e))
	}
}

func (ptr *QListView) ResizeEventDefault(e gui.QResizeEvent_ITF) {
	defer qt.Recovering("QListView::resizeEvent")

	if ptr.Pointer() != nil {
		C.QListView_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(e))
	}
}

func (ptr *QListView) ConnectRowsAboutToBeRemoved(f func(parent *core.QModelIndex, start int, end int)) {
	defer qt.Recovering("connect QListView::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "rowsAboutToBeRemoved", f)
	}
}

func (ptr *QListView) DisconnectRowsAboutToBeRemoved() {
	defer qt.Recovering("disconnect QListView::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "rowsAboutToBeRemoved")
	}
}

//export callbackQListViewRowsAboutToBeRemoved
func callbackQListViewRowsAboutToBeRemoved(ptr unsafe.Pointer, ptrName *C.char, parent unsafe.Pointer, start C.int, end C.int) bool {
	defer qt.Recovering("callback QListView::rowsAboutToBeRemoved")

	if signal := qt.GetSignal(C.GoString(ptrName), "rowsAboutToBeRemoved"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(start), int(end))
		return true
	}
	return false

}

func (ptr *QListView) RowsAboutToBeRemoved(parent core.QModelIndex_ITF, start int, end int) {
	defer qt.Recovering("QListView::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {
		C.QListView_RowsAboutToBeRemoved(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(start), C.int(end))
	}
}

func (ptr *QListView) RowsAboutToBeRemovedDefault(parent core.QModelIndex_ITF, start int, end int) {
	defer qt.Recovering("QListView::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {
		C.QListView_RowsAboutToBeRemovedDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(start), C.int(end))
	}
}

func (ptr *QListView) ConnectRowsInserted(f func(parent *core.QModelIndex, start int, end int)) {
	defer qt.Recovering("connect QListView::rowsInserted")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "rowsInserted", f)
	}
}

func (ptr *QListView) DisconnectRowsInserted() {
	defer qt.Recovering("disconnect QListView::rowsInserted")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "rowsInserted")
	}
}

//export callbackQListViewRowsInserted
func callbackQListViewRowsInserted(ptr unsafe.Pointer, ptrName *C.char, parent unsafe.Pointer, start C.int, end C.int) bool {
	defer qt.Recovering("callback QListView::rowsInserted")

	if signal := qt.GetSignal(C.GoString(ptrName), "rowsInserted"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(start), int(end))
		return true
	}
	return false

}

func (ptr *QListView) RowsInserted(parent core.QModelIndex_ITF, start int, end int) {
	defer qt.Recovering("QListView::rowsInserted")

	if ptr.Pointer() != nil {
		C.QListView_RowsInserted(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(start), C.int(end))
	}
}

func (ptr *QListView) RowsInsertedDefault(parent core.QModelIndex_ITF, start int, end int) {
	defer qt.Recovering("QListView::rowsInserted")

	if ptr.Pointer() != nil {
		C.QListView_RowsInsertedDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(start), C.int(end))
	}
}

func (ptr *QListView) ConnectScrollTo(f func(index *core.QModelIndex, hint QAbstractItemView__ScrollHint)) {
	defer qt.Recovering("connect QListView::scrollTo")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "scrollTo", f)
	}
}

func (ptr *QListView) DisconnectScrollTo() {
	defer qt.Recovering("disconnect QListView::scrollTo")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "scrollTo")
	}
}

//export callbackQListViewScrollTo
func callbackQListViewScrollTo(ptr unsafe.Pointer, ptrName *C.char, index unsafe.Pointer, hint C.int) {
	defer qt.Recovering("callback QListView::scrollTo")

	if signal := qt.GetSignal(C.GoString(ptrName), "scrollTo"); signal != nil {
		signal.(func(*core.QModelIndex, QAbstractItemView__ScrollHint))(core.NewQModelIndexFromPointer(index), QAbstractItemView__ScrollHint(hint))
	}

}

func (ptr *QListView) ScrollTo(index core.QModelIndex_ITF, hint QAbstractItemView__ScrollHint) {
	defer qt.Recovering("QListView::scrollTo")

	if ptr.Pointer() != nil {
		C.QListView_ScrollTo(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(hint))
	}
}

func (ptr *QListView) ScrollToDefault(index core.QModelIndex_ITF, hint QAbstractItemView__ScrollHint) {
	defer qt.Recovering("QListView::scrollTo")

	if ptr.Pointer() != nil {
		C.QListView_ScrollToDefault(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(hint))
	}
}

func (ptr *QListView) SetRowHidden(row int, hide bool) {
	defer qt.Recovering("QListView::setRowHidden")

	if ptr.Pointer() != nil {
		C.QListView_SetRowHidden(ptr.Pointer(), C.int(row), C.int(qt.GoBoolToInt(hide)))
	}
}

func (ptr *QListView) ConnectSetSelection(f func(rect *core.QRect, command core.QItemSelectionModel__SelectionFlag)) {
	defer qt.Recovering("connect QListView::setSelection")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setSelection", f)
	}
}

func (ptr *QListView) DisconnectSetSelection() {
	defer qt.Recovering("disconnect QListView::setSelection")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setSelection")
	}
}

//export callbackQListViewSetSelection
func callbackQListViewSetSelection(ptr unsafe.Pointer, ptrName *C.char, rect unsafe.Pointer, command C.int) {
	defer qt.Recovering("callback QListView::setSelection")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSelection"); signal != nil {
		signal.(func(*core.QRect, core.QItemSelectionModel__SelectionFlag))(core.NewQRectFromPointer(rect), core.QItemSelectionModel__SelectionFlag(command))
	}

}

func (ptr *QListView) SetSelection(rect core.QRect_ITF, command core.QItemSelectionModel__SelectionFlag) {
	defer qt.Recovering("QListView::setSelection")

	if ptr.Pointer() != nil {
		C.QListView_SetSelection(ptr.Pointer(), core.PointerFromQRect(rect), C.int(command))
	}
}

func (ptr *QListView) SetSelectionDefault(rect core.QRect_ITF, command core.QItemSelectionModel__SelectionFlag) {
	defer qt.Recovering("QListView::setSelection")

	if ptr.Pointer() != nil {
		C.QListView_SetSelectionDefault(ptr.Pointer(), core.PointerFromQRect(rect), C.int(command))
	}
}

func (ptr *QListView) ConnectStartDrag(f func(supportedActions core.Qt__DropAction)) {
	defer qt.Recovering("connect QListView::startDrag")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "startDrag", f)
	}
}

func (ptr *QListView) DisconnectStartDrag() {
	defer qt.Recovering("disconnect QListView::startDrag")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "startDrag")
	}
}

//export callbackQListViewStartDrag
func callbackQListViewStartDrag(ptr unsafe.Pointer, ptrName *C.char, supportedActions C.int) {
	defer qt.Recovering("callback QListView::startDrag")

	if signal := qt.GetSignal(C.GoString(ptrName), "startDrag"); signal != nil {
		signal.(func(core.Qt__DropAction))(core.Qt__DropAction(supportedActions))
	} else {
		NewQListViewFromPointer(ptr).StartDragDefault(core.Qt__DropAction(supportedActions))
	}
}

func (ptr *QListView) StartDrag(supportedActions core.Qt__DropAction) {
	defer qt.Recovering("QListView::startDrag")

	if ptr.Pointer() != nil {
		C.QListView_StartDrag(ptr.Pointer(), C.int(supportedActions))
	}
}

func (ptr *QListView) StartDragDefault(supportedActions core.Qt__DropAction) {
	defer qt.Recovering("QListView::startDrag")

	if ptr.Pointer() != nil {
		C.QListView_StartDragDefault(ptr.Pointer(), C.int(supportedActions))
	}
}

func (ptr *QListView) ConnectTimerEvent(f func(e *core.QTimerEvent)) {
	defer qt.Recovering("connect QListView::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QListView) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QListView::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQListViewTimerEvent
func callbackQListViewTimerEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QListView::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(e))
	} else {
		NewQListViewFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(e))
	}
}

func (ptr *QListView) TimerEvent(e core.QTimerEvent_ITF) {
	defer qt.Recovering("QListView::timerEvent")

	if ptr.Pointer() != nil {
		C.QListView_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(e))
	}
}

func (ptr *QListView) TimerEventDefault(e core.QTimerEvent_ITF) {
	defer qt.Recovering("QListView::timerEvent")

	if ptr.Pointer() != nil {
		C.QListView_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(e))
	}
}

func (ptr *QListView) ConnectUpdateGeometries(f func()) {
	defer qt.Recovering("connect QListView::updateGeometries")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updateGeometries", f)
	}
}

func (ptr *QListView) DisconnectUpdateGeometries() {
	defer qt.Recovering("disconnect QListView::updateGeometries")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updateGeometries")
	}
}

//export callbackQListViewUpdateGeometries
func callbackQListViewUpdateGeometries(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QListView::updateGeometries")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateGeometries"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QListView) UpdateGeometries() {
	defer qt.Recovering("QListView::updateGeometries")

	if ptr.Pointer() != nil {
		C.QListView_UpdateGeometries(ptr.Pointer())
	}
}

func (ptr *QListView) UpdateGeometriesDefault() {
	defer qt.Recovering("QListView::updateGeometries")

	if ptr.Pointer() != nil {
		C.QListView_UpdateGeometriesDefault(ptr.Pointer())
	}
}

func (ptr *QListView) VerticalOffset() int {
	defer qt.Recovering("QListView::verticalOffset")

	if ptr.Pointer() != nil {
		return int(C.QListView_VerticalOffset(ptr.Pointer()))
	}
	return 0
}

func (ptr *QListView) ViewportSizeHint() *core.QSize {
	defer qt.Recovering("QListView::viewportSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QListView_ViewportSizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QListView) VisualRect(index core.QModelIndex_ITF) *core.QRect {
	defer qt.Recovering("QListView::visualRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QListView_VisualRect(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return nil
}

func (ptr *QListView) VisualRegionForSelection(selection core.QItemSelection_ITF) *gui.QRegion {
	defer qt.Recovering("QListView::visualRegionForSelection")

	if ptr.Pointer() != nil {
		return gui.NewQRegionFromPointer(C.QListView_VisualRegionForSelection(ptr.Pointer(), core.PointerFromQItemSelection(selection)))
	}
	return nil
}

func (ptr *QListView) DestroyQListView() {
	defer qt.Recovering("QListView::~QListView")

	if ptr.Pointer() != nil {
		C.QListView_DestroyQListView(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QListView) ConnectCloseEditor(f func(editor *QWidget, hint QAbstractItemDelegate__EndEditHint)) {
	defer qt.Recovering("connect QListView::closeEditor")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEditor", f)
	}
}

func (ptr *QListView) DisconnectCloseEditor() {
	defer qt.Recovering("disconnect QListView::closeEditor")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEditor")
	}
}

//export callbackQListViewCloseEditor
func callbackQListViewCloseEditor(ptr unsafe.Pointer, ptrName *C.char, editor unsafe.Pointer, hint C.int) bool {
	defer qt.Recovering("callback QListView::closeEditor")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEditor"); signal != nil {
		signal.(func(*QWidget, QAbstractItemDelegate__EndEditHint))(NewQWidgetFromPointer(editor), QAbstractItemDelegate__EndEditHint(hint))
		return true
	}
	return false

}

func (ptr *QListView) CloseEditor(editor QWidget_ITF, hint QAbstractItemDelegate__EndEditHint) {
	defer qt.Recovering("QListView::closeEditor")

	if ptr.Pointer() != nil {
		C.QListView_CloseEditor(ptr.Pointer(), PointerFromQWidget(editor), C.int(hint))
	}
}

func (ptr *QListView) CloseEditorDefault(editor QWidget_ITF, hint QAbstractItemDelegate__EndEditHint) {
	defer qt.Recovering("QListView::closeEditor")

	if ptr.Pointer() != nil {
		C.QListView_CloseEditorDefault(ptr.Pointer(), PointerFromQWidget(editor), C.int(hint))
	}
}

func (ptr *QListView) ConnectCommitData(f func(editor *QWidget)) {
	defer qt.Recovering("connect QListView::commitData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "commitData", f)
	}
}

func (ptr *QListView) DisconnectCommitData() {
	defer qt.Recovering("disconnect QListView::commitData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "commitData")
	}
}

//export callbackQListViewCommitData
func callbackQListViewCommitData(ptr unsafe.Pointer, ptrName *C.char, editor unsafe.Pointer) bool {
	defer qt.Recovering("callback QListView::commitData")

	if signal := qt.GetSignal(C.GoString(ptrName), "commitData"); signal != nil {
		signal.(func(*QWidget))(NewQWidgetFromPointer(editor))
		return true
	}
	return false

}

func (ptr *QListView) CommitData(editor QWidget_ITF) {
	defer qt.Recovering("QListView::commitData")

	if ptr.Pointer() != nil {
		C.QListView_CommitData(ptr.Pointer(), PointerFromQWidget(editor))
	}
}

func (ptr *QListView) CommitDataDefault(editor QWidget_ITF) {
	defer qt.Recovering("QListView::commitData")

	if ptr.Pointer() != nil {
		C.QListView_CommitDataDefault(ptr.Pointer(), PointerFromQWidget(editor))
	}
}

func (ptr *QListView) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QListView::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QListView) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QListView::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQListViewDragEnterEvent
func callbackQListViewDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QListView::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQListViewFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QListView) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QListView::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QListView_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QListView) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QListView::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QListView_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QListView) ConnectEditorDestroyed(f func(editor *core.QObject)) {
	defer qt.Recovering("connect QListView::editorDestroyed")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "editorDestroyed", f)
	}
}

func (ptr *QListView) DisconnectEditorDestroyed() {
	defer qt.Recovering("disconnect QListView::editorDestroyed")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "editorDestroyed")
	}
}

//export callbackQListViewEditorDestroyed
func callbackQListViewEditorDestroyed(ptr unsafe.Pointer, ptrName *C.char, editor unsafe.Pointer) bool {
	defer qt.Recovering("callback QListView::editorDestroyed")

	if signal := qt.GetSignal(C.GoString(ptrName), "editorDestroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(editor))
		return true
	}
	return false

}

func (ptr *QListView) EditorDestroyed(editor core.QObject_ITF) {
	defer qt.Recovering("QListView::editorDestroyed")

	if ptr.Pointer() != nil {
		C.QListView_EditorDestroyed(ptr.Pointer(), core.PointerFromQObject(editor))
	}
}

func (ptr *QListView) EditorDestroyedDefault(editor core.QObject_ITF) {
	defer qt.Recovering("QListView::editorDestroyed")

	if ptr.Pointer() != nil {
		C.QListView_EditorDestroyedDefault(ptr.Pointer(), core.PointerFromQObject(editor))
	}
}

func (ptr *QListView) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QListView::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QListView) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QListView::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQListViewFocusInEvent
func callbackQListViewFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QListView::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQListViewFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QListView) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QListView::focusInEvent")

	if ptr.Pointer() != nil {
		C.QListView_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QListView) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QListView::focusInEvent")

	if ptr.Pointer() != nil {
		C.QListView_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QListView) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QListView::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QListView) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QListView::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQListViewFocusOutEvent
func callbackQListViewFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QListView::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQListViewFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QListView) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QListView::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QListView_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QListView) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QListView::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QListView_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QListView) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QListView::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QListView) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QListView::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQListViewInputMethodEvent
func callbackQListViewInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QListView::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQListViewFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QListView) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QListView::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QListView_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QListView) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QListView::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QListView_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QListView) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QListView::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QListView) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QListView::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQListViewKeyPressEvent
func callbackQListViewKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QListView::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQListViewFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QListView) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QListView::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QListView_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QListView) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QListView::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QListView_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QListView) ConnectKeyboardSearch(f func(search string)) {
	defer qt.Recovering("connect QListView::keyboardSearch")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyboardSearch", f)
	}
}

func (ptr *QListView) DisconnectKeyboardSearch() {
	defer qt.Recovering("disconnect QListView::keyboardSearch")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyboardSearch")
	}
}

//export callbackQListViewKeyboardSearch
func callbackQListViewKeyboardSearch(ptr unsafe.Pointer, ptrName *C.char, search *C.char) {
	defer qt.Recovering("callback QListView::keyboardSearch")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyboardSearch"); signal != nil {
		signal.(func(string))(C.GoString(search))
	} else {
		NewQListViewFromPointer(ptr).KeyboardSearchDefault(C.GoString(search))
	}
}

func (ptr *QListView) KeyboardSearch(search string) {
	defer qt.Recovering("QListView::keyboardSearch")

	if ptr.Pointer() != nil {
		C.QListView_KeyboardSearch(ptr.Pointer(), C.CString(search))
	}
}

func (ptr *QListView) KeyboardSearchDefault(search string) {
	defer qt.Recovering("QListView::keyboardSearch")

	if ptr.Pointer() != nil {
		C.QListView_KeyboardSearchDefault(ptr.Pointer(), C.CString(search))
	}
}

func (ptr *QListView) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QListView::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QListView) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QListView::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQListViewMouseDoubleClickEvent
func callbackQListViewMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QListView::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQListViewFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QListView) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QListView::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QListView_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QListView) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QListView::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QListView_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QListView) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QListView::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QListView) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QListView::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQListViewMousePressEvent
func callbackQListViewMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QListView::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQListViewFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QListView) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QListView::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QListView_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QListView) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QListView::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QListView_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QListView) ConnectReset(f func()) {
	defer qt.Recovering("connect QListView::reset")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "reset", f)
	}
}

func (ptr *QListView) DisconnectReset() {
	defer qt.Recovering("disconnect QListView::reset")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "reset")
	}
}

//export callbackQListViewReset
func callbackQListViewReset(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QListView::reset")

	if signal := qt.GetSignal(C.GoString(ptrName), "reset"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QListView) Reset() {
	defer qt.Recovering("QListView::reset")

	if ptr.Pointer() != nil {
		C.QListView_Reset(ptr.Pointer())
	}
}

func (ptr *QListView) ResetDefault() {
	defer qt.Recovering("QListView::reset")

	if ptr.Pointer() != nil {
		C.QListView_ResetDefault(ptr.Pointer())
	}
}

func (ptr *QListView) ConnectSelectAll(f func()) {
	defer qt.Recovering("connect QListView::selectAll")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "selectAll", f)
	}
}

func (ptr *QListView) DisconnectSelectAll() {
	defer qt.Recovering("disconnect QListView::selectAll")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "selectAll")
	}
}

//export callbackQListViewSelectAll
func callbackQListViewSelectAll(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QListView::selectAll")

	if signal := qt.GetSignal(C.GoString(ptrName), "selectAll"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QListView) SelectAll() {
	defer qt.Recovering("QListView::selectAll")

	if ptr.Pointer() != nil {
		C.QListView_SelectAll(ptr.Pointer())
	}
}

func (ptr *QListView) SelectAllDefault() {
	defer qt.Recovering("QListView::selectAll")

	if ptr.Pointer() != nil {
		C.QListView_SelectAllDefault(ptr.Pointer())
	}
}

func (ptr *QListView) ConnectSetModel(f func(model *core.QAbstractItemModel)) {
	defer qt.Recovering("connect QListView::setModel")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setModel", f)
	}
}

func (ptr *QListView) DisconnectSetModel() {
	defer qt.Recovering("disconnect QListView::setModel")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setModel")
	}
}

//export callbackQListViewSetModel
func callbackQListViewSetModel(ptr unsafe.Pointer, ptrName *C.char, model unsafe.Pointer) {
	defer qt.Recovering("callback QListView::setModel")

	if signal := qt.GetSignal(C.GoString(ptrName), "setModel"); signal != nil {
		signal.(func(*core.QAbstractItemModel))(core.NewQAbstractItemModelFromPointer(model))
	} else {
		NewQListViewFromPointer(ptr).SetModelDefault(core.NewQAbstractItemModelFromPointer(model))
	}
}

func (ptr *QListView) SetModel(model core.QAbstractItemModel_ITF) {
	defer qt.Recovering("QListView::setModel")

	if ptr.Pointer() != nil {
		C.QListView_SetModel(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QListView) SetModelDefault(model core.QAbstractItemModel_ITF) {
	defer qt.Recovering("QListView::setModel")

	if ptr.Pointer() != nil {
		C.QListView_SetModelDefault(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QListView) ConnectSetRootIndex(f func(index *core.QModelIndex)) {
	defer qt.Recovering("connect QListView::setRootIndex")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setRootIndex", f)
	}
}

func (ptr *QListView) DisconnectSetRootIndex() {
	defer qt.Recovering("disconnect QListView::setRootIndex")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setRootIndex")
	}
}

//export callbackQListViewSetRootIndex
func callbackQListViewSetRootIndex(ptr unsafe.Pointer, ptrName *C.char, index unsafe.Pointer) bool {
	defer qt.Recovering("callback QListView::setRootIndex")

	if signal := qt.GetSignal(C.GoString(ptrName), "setRootIndex"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
		return true
	}
	return false

}

func (ptr *QListView) SetRootIndex(index core.QModelIndex_ITF) {
	defer qt.Recovering("QListView::setRootIndex")

	if ptr.Pointer() != nil {
		C.QListView_SetRootIndex(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QListView) SetRootIndexDefault(index core.QModelIndex_ITF) {
	defer qt.Recovering("QListView::setRootIndex")

	if ptr.Pointer() != nil {
		C.QListView_SetRootIndexDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QListView) ConnectSetSelectionModel(f func(selectionModel *core.QItemSelectionModel)) {
	defer qt.Recovering("connect QListView::setSelectionModel")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setSelectionModel", f)
	}
}

func (ptr *QListView) DisconnectSetSelectionModel() {
	defer qt.Recovering("disconnect QListView::setSelectionModel")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setSelectionModel")
	}
}

//export callbackQListViewSetSelectionModel
func callbackQListViewSetSelectionModel(ptr unsafe.Pointer, ptrName *C.char, selectionModel unsafe.Pointer) {
	defer qt.Recovering("callback QListView::setSelectionModel")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSelectionModel"); signal != nil {
		signal.(func(*core.QItemSelectionModel))(core.NewQItemSelectionModelFromPointer(selectionModel))
	} else {
		NewQListViewFromPointer(ptr).SetSelectionModelDefault(core.NewQItemSelectionModelFromPointer(selectionModel))
	}
}

func (ptr *QListView) SetSelectionModel(selectionModel core.QItemSelectionModel_ITF) {
	defer qt.Recovering("QListView::setSelectionModel")

	if ptr.Pointer() != nil {
		C.QListView_SetSelectionModel(ptr.Pointer(), core.PointerFromQItemSelectionModel(selectionModel))
	}
}

func (ptr *QListView) SetSelectionModelDefault(selectionModel core.QItemSelectionModel_ITF) {
	defer qt.Recovering("QListView::setSelectionModel")

	if ptr.Pointer() != nil {
		C.QListView_SetSelectionModelDefault(ptr.Pointer(), core.PointerFromQItemSelectionModel(selectionModel))
	}
}

func (ptr *QListView) ConnectContextMenuEvent(f func(e *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QListView::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QListView) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QListView::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQListViewContextMenuEvent
func callbackQListViewContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QListView::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
	} else {
		NewQListViewFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(e))
	}
}

func (ptr *QListView) ContextMenuEvent(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QListView::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QListView_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

func (ptr *QListView) ContextMenuEventDefault(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QListView::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QListView_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

func (ptr *QListView) ConnectScrollContentsBy(f func(dx int, dy int)) {
	defer qt.Recovering("connect QListView::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "scrollContentsBy", f)
	}
}

func (ptr *QListView) DisconnectScrollContentsBy() {
	defer qt.Recovering("disconnect QListView::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "scrollContentsBy")
	}
}

//export callbackQListViewScrollContentsBy
func callbackQListViewScrollContentsBy(ptr unsafe.Pointer, ptrName *C.char, dx C.int, dy C.int) {
	defer qt.Recovering("callback QListView::scrollContentsBy")

	if signal := qt.GetSignal(C.GoString(ptrName), "scrollContentsBy"); signal != nil {
		signal.(func(int, int))(int(dx), int(dy))
	} else {
		NewQListViewFromPointer(ptr).ScrollContentsByDefault(int(dx), int(dy))
	}
}

func (ptr *QListView) ScrollContentsBy(dx int, dy int) {
	defer qt.Recovering("QListView::scrollContentsBy")

	if ptr.Pointer() != nil {
		C.QListView_ScrollContentsBy(ptr.Pointer(), C.int(dx), C.int(dy))
	}
}

func (ptr *QListView) ScrollContentsByDefault(dx int, dy int) {
	defer qt.Recovering("QListView::scrollContentsBy")

	if ptr.Pointer() != nil {
		C.QListView_ScrollContentsByDefault(ptr.Pointer(), C.int(dx), C.int(dy))
	}
}

func (ptr *QListView) ConnectSetupViewport(f func(viewport *QWidget)) {
	defer qt.Recovering("connect QListView::setupViewport")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setupViewport", f)
	}
}

func (ptr *QListView) DisconnectSetupViewport() {
	defer qt.Recovering("disconnect QListView::setupViewport")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setupViewport")
	}
}

//export callbackQListViewSetupViewport
func callbackQListViewSetupViewport(ptr unsafe.Pointer, ptrName *C.char, viewport unsafe.Pointer) {
	defer qt.Recovering("callback QListView::setupViewport")

	if signal := qt.GetSignal(C.GoString(ptrName), "setupViewport"); signal != nil {
		signal.(func(*QWidget))(NewQWidgetFromPointer(viewport))
	} else {
		NewQListViewFromPointer(ptr).SetupViewportDefault(NewQWidgetFromPointer(viewport))
	}
}

func (ptr *QListView) SetupViewport(viewport QWidget_ITF) {
	defer qt.Recovering("QListView::setupViewport")

	if ptr.Pointer() != nil {
		C.QListView_SetupViewport(ptr.Pointer(), PointerFromQWidget(viewport))
	}
}

func (ptr *QListView) SetupViewportDefault(viewport QWidget_ITF) {
	defer qt.Recovering("QListView::setupViewport")

	if ptr.Pointer() != nil {
		C.QListView_SetupViewportDefault(ptr.Pointer(), PointerFromQWidget(viewport))
	}
}

func (ptr *QListView) ConnectWheelEvent(f func(e *gui.QWheelEvent)) {
	defer qt.Recovering("connect QListView::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QListView) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QListView::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQListViewWheelEvent
func callbackQListViewWheelEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QListView::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
	} else {
		NewQListViewFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(e))
	}
}

func (ptr *QListView) WheelEvent(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QListView::wheelEvent")

	if ptr.Pointer() != nil {
		C.QListView_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QListView) WheelEventDefault(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QListView::wheelEvent")

	if ptr.Pointer() != nil {
		C.QListView_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QListView) ConnectChangeEvent(f func(ev *core.QEvent)) {
	defer qt.Recovering("connect QListView::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QListView) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QListView::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQListViewChangeEvent
func callbackQListViewChangeEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QListView::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
	} else {
		NewQListViewFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(ev))
	}
}

func (ptr *QListView) ChangeEvent(ev core.QEvent_ITF) {
	defer qt.Recovering("QListView::changeEvent")

	if ptr.Pointer() != nil {
		C.QListView_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QListView) ChangeEventDefault(ev core.QEvent_ITF) {
	defer qt.Recovering("QListView::changeEvent")

	if ptr.Pointer() != nil {
		C.QListView_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QListView) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QListView::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QListView) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QListView::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQListViewActionEvent
func callbackQListViewActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QListView::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQListViewFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QListView) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QListView::actionEvent")

	if ptr.Pointer() != nil {
		C.QListView_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QListView) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QListView::actionEvent")

	if ptr.Pointer() != nil {
		C.QListView_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QListView) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QListView::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QListView) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QListView::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQListViewEnterEvent
func callbackQListViewEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QListView::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQListViewFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QListView) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QListView::enterEvent")

	if ptr.Pointer() != nil {
		C.QListView_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QListView) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QListView::enterEvent")

	if ptr.Pointer() != nil {
		C.QListView_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QListView) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QListView::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QListView) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QListView::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQListViewHideEvent
func callbackQListViewHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QListView::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQListViewFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QListView) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QListView::hideEvent")

	if ptr.Pointer() != nil {
		C.QListView_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QListView) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QListView::hideEvent")

	if ptr.Pointer() != nil {
		C.QListView_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QListView) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QListView::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QListView) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QListView::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQListViewLeaveEvent
func callbackQListViewLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QListView::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQListViewFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QListView) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QListView::leaveEvent")

	if ptr.Pointer() != nil {
		C.QListView_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QListView) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QListView::leaveEvent")

	if ptr.Pointer() != nil {
		C.QListView_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QListView) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QListView::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QListView) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QListView::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQListViewMoveEvent
func callbackQListViewMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QListView::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQListViewFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QListView) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QListView::moveEvent")

	if ptr.Pointer() != nil {
		C.QListView_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QListView) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QListView::moveEvent")

	if ptr.Pointer() != nil {
		C.QListView_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QListView) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QListView::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QListView) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QListView::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQListViewSetVisible
func callbackQListViewSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QListView::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QListView) SetVisible(visible bool) {
	defer qt.Recovering("QListView::setVisible")

	if ptr.Pointer() != nil {
		C.QListView_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QListView) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QListView::setVisible")

	if ptr.Pointer() != nil {
		C.QListView_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QListView) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QListView::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QListView) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QListView::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQListViewShowEvent
func callbackQListViewShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QListView::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQListViewFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QListView) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QListView::showEvent")

	if ptr.Pointer() != nil {
		C.QListView_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QListView) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QListView::showEvent")

	if ptr.Pointer() != nil {
		C.QListView_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QListView) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QListView::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QListView) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QListView::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQListViewCloseEvent
func callbackQListViewCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QListView::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQListViewFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QListView) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QListView::closeEvent")

	if ptr.Pointer() != nil {
		C.QListView_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QListView) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QListView::closeEvent")

	if ptr.Pointer() != nil {
		C.QListView_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QListView) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QListView::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QListView) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QListView::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQListViewInitPainter
func callbackQListViewInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QListView::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQListViewFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QListView) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QListView::initPainter")

	if ptr.Pointer() != nil {
		C.QListView_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QListView) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QListView::initPainter")

	if ptr.Pointer() != nil {
		C.QListView_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QListView) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QListView::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QListView) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QListView::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQListViewKeyReleaseEvent
func callbackQListViewKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QListView::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQListViewFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QListView) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QListView::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QListView_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QListView) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QListView::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QListView_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QListView) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QListView::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QListView) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QListView::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQListViewTabletEvent
func callbackQListViewTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QListView::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQListViewFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QListView) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QListView::tabletEvent")

	if ptr.Pointer() != nil {
		C.QListView_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QListView) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QListView::tabletEvent")

	if ptr.Pointer() != nil {
		C.QListView_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QListView) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QListView::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QListView) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QListView::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQListViewChildEvent
func callbackQListViewChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QListView::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQListViewFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QListView) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QListView::childEvent")

	if ptr.Pointer() != nil {
		C.QListView_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QListView) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QListView::childEvent")

	if ptr.Pointer() != nil {
		C.QListView_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QListView) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QListView::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QListView) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QListView::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQListViewCustomEvent
func callbackQListViewCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QListView::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQListViewFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QListView) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QListView::customEvent")

	if ptr.Pointer() != nil {
		C.QListView_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QListView) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QListView::customEvent")

	if ptr.Pointer() != nil {
		C.QListView_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
