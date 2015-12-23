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
func callbackQListViewCurrentChanged(ptrName *C.char, current unsafe.Pointer, previous unsafe.Pointer) bool {
	defer qt.Recovering("callback QListView::currentChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "currentChanged"); signal != nil {
		signal.(func(*core.QModelIndex, *core.QModelIndex))(core.NewQModelIndexFromPointer(current), core.NewQModelIndexFromPointer(previous))
		return true
	}
	return false

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
func callbackQListViewDragLeaveEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QListView::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(e))
		return true
	}
	return false

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
func callbackQListViewDragMoveEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QListView::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(e))
		return true
	}
	return false

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
func callbackQListViewDropEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QListView::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QListView) IndexAt(p core.QPoint_ITF) *core.QModelIndex {
	defer qt.Recovering("QListView::indexAt")

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QListView_IndexAt(ptr.Pointer(), core.PointerFromQPoint(p)))
	}
	return nil
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
func callbackQListViewMouseMoveEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QListView::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

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
func callbackQListViewMouseReleaseEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QListView::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

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
func callbackQListViewPaintEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QListView::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(e))
		return true
	}
	return false

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
func callbackQListViewResizeEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QListView::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(e))
		return true
	}
	return false

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
func callbackQListViewRowsAboutToBeRemoved(ptrName *C.char, parent unsafe.Pointer, start C.int, end C.int) bool {
	defer qt.Recovering("callback QListView::rowsAboutToBeRemoved")

	if signal := qt.GetSignal(C.GoString(ptrName), "rowsAboutToBeRemoved"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(start), int(end))
		return true
	}
	return false

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
func callbackQListViewRowsInserted(ptrName *C.char, parent unsafe.Pointer, start C.int, end C.int) bool {
	defer qt.Recovering("callback QListView::rowsInserted")

	if signal := qt.GetSignal(C.GoString(ptrName), "rowsInserted"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(start), int(end))
		return true
	}
	return false

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
func callbackQListViewScrollTo(ptrName *C.char, index unsafe.Pointer, hint C.int) bool {
	defer qt.Recovering("callback QListView::scrollTo")

	if signal := qt.GetSignal(C.GoString(ptrName), "scrollTo"); signal != nil {
		signal.(func(*core.QModelIndex, QAbstractItemView__ScrollHint))(core.NewQModelIndexFromPointer(index), QAbstractItemView__ScrollHint(hint))
		return true
	}
	return false

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
func callbackQListViewSetSelection(ptrName *C.char, rect unsafe.Pointer, command C.int) bool {
	defer qt.Recovering("callback QListView::setSelection")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSelection"); signal != nil {
		signal.(func(*core.QRect, core.QItemSelectionModel__SelectionFlag))(core.NewQRectFromPointer(rect), core.QItemSelectionModel__SelectionFlag(command))
		return true
	}
	return false

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
func callbackQListViewStartDrag(ptrName *C.char, supportedActions C.int) bool {
	defer qt.Recovering("callback QListView::startDrag")

	if signal := qt.GetSignal(C.GoString(ptrName), "startDrag"); signal != nil {
		signal.(func(core.Qt__DropAction))(core.Qt__DropAction(supportedActions))
		return true
	}
	return false

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
func callbackQListViewTimerEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QListView::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(e))
		return true
	}
	return false

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
func callbackQListViewUpdateGeometries(ptrName *C.char) bool {
	defer qt.Recovering("callback QListView::updateGeometries")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateGeometries"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QListView) VisualRect(index core.QModelIndex_ITF) *core.QRect {
	defer qt.Recovering("QListView::visualRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QListView_VisualRect(ptr.Pointer(), core.PointerFromQModelIndex(index)))
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
func callbackQListViewCloseEditor(ptrName *C.char, editor unsafe.Pointer, hint C.int) bool {
	defer qt.Recovering("callback QListView::closeEditor")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEditor"); signal != nil {
		signal.(func(*QWidget, QAbstractItemDelegate__EndEditHint))(NewQWidgetFromPointer(editor), QAbstractItemDelegate__EndEditHint(hint))
		return true
	}
	return false

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
func callbackQListViewCommitData(ptrName *C.char, editor unsafe.Pointer) bool {
	defer qt.Recovering("callback QListView::commitData")

	if signal := qt.GetSignal(C.GoString(ptrName), "commitData"); signal != nil {
		signal.(func(*QWidget))(NewQWidgetFromPointer(editor))
		return true
	}
	return false

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
func callbackQListViewDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QListView::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

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
func callbackQListViewEditorDestroyed(ptrName *C.char, editor unsafe.Pointer) bool {
	defer qt.Recovering("callback QListView::editorDestroyed")

	if signal := qt.GetSignal(C.GoString(ptrName), "editorDestroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(editor))
		return true
	}
	return false

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
func callbackQListViewFocusInEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QListView::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

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
func callbackQListViewFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QListView::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

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
func callbackQListViewInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QListView::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

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
func callbackQListViewKeyPressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QListView::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

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
func callbackQListViewKeyboardSearch(ptrName *C.char, search *C.char) bool {
	defer qt.Recovering("callback QListView::keyboardSearch")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyboardSearch"); signal != nil {
		signal.(func(string))(C.GoString(search))
		return true
	}
	return false

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
func callbackQListViewMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QListView::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQListViewMousePressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QListView::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQListViewReset(ptrName *C.char) bool {
	defer qt.Recovering("callback QListView::reset")

	if signal := qt.GetSignal(C.GoString(ptrName), "reset"); signal != nil {
		signal.(func())()
		return true
	}
	return false

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
func callbackQListViewSelectAll(ptrName *C.char) bool {
	defer qt.Recovering("callback QListView::selectAll")

	if signal := qt.GetSignal(C.GoString(ptrName), "selectAll"); signal != nil {
		signal.(func())()
		return true
	}
	return false

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
func callbackQListViewSetRootIndex(ptrName *C.char, index unsafe.Pointer) bool {
	defer qt.Recovering("callback QListView::setRootIndex")

	if signal := qt.GetSignal(C.GoString(ptrName), "setRootIndex"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
		return true
	}
	return false

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
func callbackQListViewSetSelectionModel(ptrName *C.char, selectionModel unsafe.Pointer) bool {
	defer qt.Recovering("callback QListView::setSelectionModel")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSelectionModel"); signal != nil {
		signal.(func(*core.QItemSelectionModel))(core.NewQItemSelectionModelFromPointer(selectionModel))
		return true
	}
	return false

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
func callbackQListViewContextMenuEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QListView::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
		return true
	}
	return false

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
func callbackQListViewScrollContentsBy(ptrName *C.char, dx C.int, dy C.int) bool {
	defer qt.Recovering("callback QListView::scrollContentsBy")

	if signal := qt.GetSignal(C.GoString(ptrName), "scrollContentsBy"); signal != nil {
		signal.(func(int, int))(int(dx), int(dy))
		return true
	}
	return false

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
func callbackQListViewSetupViewport(ptrName *C.char, viewport unsafe.Pointer) bool {
	defer qt.Recovering("callback QListView::setupViewport")

	if signal := qt.GetSignal(C.GoString(ptrName), "setupViewport"); signal != nil {
		signal.(func(*QWidget))(NewQWidgetFromPointer(viewport))
		return true
	}
	return false

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
func callbackQListViewWheelEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QListView::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
		return true
	}
	return false

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
func callbackQListViewChangeEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QListView::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
		return true
	}
	return false

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
func callbackQListViewActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QListView::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

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
func callbackQListViewEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QListView::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

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
func callbackQListViewHideEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QListView::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
		return true
	}
	return false

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
func callbackQListViewLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QListView::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

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
func callbackQListViewMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QListView::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

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
func callbackQListViewSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QListView::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

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
func callbackQListViewShowEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QListView::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
		return true
	}
	return false

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
func callbackQListViewCloseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QListView::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQListViewInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QListView::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

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
func callbackQListViewKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QListView::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

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
func callbackQListViewTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QListView::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

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
func callbackQListViewChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QListView::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

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
func callbackQListViewCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QListView::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
