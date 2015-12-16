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

	var signal = qt.GetSignal(C.GoString(ptrName), "currentChanged")
	if signal != nil {
		defer signal.(func(*core.QModelIndex, *core.QModelIndex))(core.NewQModelIndexFromPointer(current), core.NewQModelIndexFromPointer(previous))
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

	var signal = qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent")
	if signal != nil {
		defer signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(e))
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

	var signal = qt.GetSignal(C.GoString(ptrName), "dragMoveEvent")
	if signal != nil {
		defer signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(e))
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

	var signal = qt.GetSignal(C.GoString(ptrName), "dropEvent")
	if signal != nil {
		defer signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(e))
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

	var signal = qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent")
	if signal != nil {
		defer signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
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

	var signal = qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent")
	if signal != nil {
		defer signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
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

	var signal = qt.GetSignal(C.GoString(ptrName), "paintEvent")
	if signal != nil {
		defer signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(e))
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

	var signal = qt.GetSignal(C.GoString(ptrName), "resizeEvent")
	if signal != nil {
		defer signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(e))
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

	var signal = qt.GetSignal(C.GoString(ptrName), "rowsAboutToBeRemoved")
	if signal != nil {
		defer signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(start), int(end))
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

	var signal = qt.GetSignal(C.GoString(ptrName), "rowsInserted")
	if signal != nil {
		defer signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(start), int(end))
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

	var signal = qt.GetSignal(C.GoString(ptrName), "scrollTo")
	if signal != nil {
		defer signal.(func(*core.QModelIndex, QAbstractItemView__ScrollHint))(core.NewQModelIndexFromPointer(index), QAbstractItemView__ScrollHint(hint))
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

	var signal = qt.GetSignal(C.GoString(ptrName), "setSelection")
	if signal != nil {
		defer signal.(func(*core.QRect, core.QItemSelectionModel__SelectionFlag))(core.NewQRectFromPointer(rect), core.QItemSelectionModel__SelectionFlag(command))
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

	var signal = qt.GetSignal(C.GoString(ptrName), "startDrag")
	if signal != nil {
		defer signal.(func(core.Qt__DropAction))(core.Qt__DropAction(supportedActions))
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

	var signal = qt.GetSignal(C.GoString(ptrName), "timerEvent")
	if signal != nil {
		defer signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(e))
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

	var signal = qt.GetSignal(C.GoString(ptrName), "updateGeometries")
	if signal != nil {
		defer signal.(func())()
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
