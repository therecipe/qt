package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QUndoView struct {
	QListView
}

type QUndoView_ITF interface {
	QListView_ITF
	QUndoView_PTR() *QUndoView
}

func PointerFromQUndoView(ptr QUndoView_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QUndoView_PTR().Pointer()
	}
	return nil
}

func NewQUndoViewFromPointer(ptr unsafe.Pointer) *QUndoView {
	var n = new(QUndoView)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QUndoView_") {
		n.SetObjectName("QUndoView_" + qt.Identifier())
	}
	return n
}

func (ptr *QUndoView) QUndoView_PTR() *QUndoView {
	return ptr
}

func (ptr *QUndoView) CleanIcon() *gui.QIcon {
	defer qt.Recovering("QUndoView::cleanIcon")

	if ptr.Pointer() != nil {
		return gui.NewQIconFromPointer(C.QUndoView_CleanIcon(ptr.Pointer()))
	}
	return nil
}

func (ptr *QUndoView) EmptyLabel() string {
	defer qt.Recovering("QUndoView::emptyLabel")

	if ptr.Pointer() != nil {
		return C.GoString(C.QUndoView_EmptyLabel(ptr.Pointer()))
	}
	return ""
}

func (ptr *QUndoView) SetCleanIcon(icon gui.QIcon_ITF) {
	defer qt.Recovering("QUndoView::setCleanIcon")

	if ptr.Pointer() != nil {
		C.QUndoView_SetCleanIcon(ptr.Pointer(), gui.PointerFromQIcon(icon))
	}
}

func (ptr *QUndoView) SetEmptyLabel(label string) {
	defer qt.Recovering("QUndoView::setEmptyLabel")

	if ptr.Pointer() != nil {
		C.QUndoView_SetEmptyLabel(ptr.Pointer(), C.CString(label))
	}
}

func NewQUndoView3(group QUndoGroup_ITF, parent QWidget_ITF) *QUndoView {
	defer qt.Recovering("QUndoView::QUndoView")

	return NewQUndoViewFromPointer(C.QUndoView_NewQUndoView3(PointerFromQUndoGroup(group), PointerFromQWidget(parent)))
}

func NewQUndoView2(stack QUndoStack_ITF, parent QWidget_ITF) *QUndoView {
	defer qt.Recovering("QUndoView::QUndoView")

	return NewQUndoViewFromPointer(C.QUndoView_NewQUndoView2(PointerFromQUndoStack(stack), PointerFromQWidget(parent)))
}

func NewQUndoView(parent QWidget_ITF) *QUndoView {
	defer qt.Recovering("QUndoView::QUndoView")

	return NewQUndoViewFromPointer(C.QUndoView_NewQUndoView(PointerFromQWidget(parent)))
}

func (ptr *QUndoView) Group() *QUndoGroup {
	defer qt.Recovering("QUndoView::group")

	if ptr.Pointer() != nil {
		return NewQUndoGroupFromPointer(C.QUndoView_Group(ptr.Pointer()))
	}
	return nil
}

func (ptr *QUndoView) SetGroup(group QUndoGroup_ITF) {
	defer qt.Recovering("QUndoView::setGroup")

	if ptr.Pointer() != nil {
		C.QUndoView_SetGroup(ptr.Pointer(), PointerFromQUndoGroup(group))
	}
}

func (ptr *QUndoView) SetStack(stack QUndoStack_ITF) {
	defer qt.Recovering("QUndoView::setStack")

	if ptr.Pointer() != nil {
		C.QUndoView_SetStack(ptr.Pointer(), PointerFromQUndoStack(stack))
	}
}

func (ptr *QUndoView) Stack() *QUndoStack {
	defer qt.Recovering("QUndoView::stack")

	if ptr.Pointer() != nil {
		return NewQUndoStackFromPointer(C.QUndoView_Stack(ptr.Pointer()))
	}
	return nil
}

func (ptr *QUndoView) DestroyQUndoView() {
	defer qt.Recovering("QUndoView::~QUndoView")

	if ptr.Pointer() != nil {
		C.QUndoView_DestroyQUndoView(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QUndoView) ConnectCurrentChanged(f func(current *core.QModelIndex, previous *core.QModelIndex)) {
	defer qt.Recovering("connect QUndoView::currentChanged")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "currentChanged", f)
	}
}

func (ptr *QUndoView) DisconnectCurrentChanged() {
	defer qt.Recovering("disconnect QUndoView::currentChanged")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "currentChanged")
	}
}

//export callbackQUndoViewCurrentChanged
func callbackQUndoViewCurrentChanged(ptr unsafe.Pointer, ptrName *C.char, current unsafe.Pointer, previous unsafe.Pointer) {
	defer qt.Recovering("callback QUndoView::currentChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "currentChanged"); signal != nil {
		signal.(func(*core.QModelIndex, *core.QModelIndex))(core.NewQModelIndexFromPointer(current), core.NewQModelIndexFromPointer(previous))
	} else {
		NewQUndoViewFromPointer(ptr).CurrentChangedDefault(core.NewQModelIndexFromPointer(current), core.NewQModelIndexFromPointer(previous))
	}
}

func (ptr *QUndoView) CurrentChanged(current core.QModelIndex_ITF, previous core.QModelIndex_ITF) {
	defer qt.Recovering("QUndoView::currentChanged")

	if ptr.Pointer() != nil {
		C.QUndoView_CurrentChanged(ptr.Pointer(), core.PointerFromQModelIndex(current), core.PointerFromQModelIndex(previous))
	}
}

func (ptr *QUndoView) CurrentChangedDefault(current core.QModelIndex_ITF, previous core.QModelIndex_ITF) {
	defer qt.Recovering("QUndoView::currentChanged")

	if ptr.Pointer() != nil {
		C.QUndoView_CurrentChangedDefault(ptr.Pointer(), core.PointerFromQModelIndex(current), core.PointerFromQModelIndex(previous))
	}
}

func (ptr *QUndoView) ConnectDragLeaveEvent(f func(e *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QUndoView::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QUndoView) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QUndoView::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQUndoViewDragLeaveEvent
func callbackQUndoViewDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QUndoView::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(e))
	} else {
		NewQUndoViewFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(e))
	}
}

func (ptr *QUndoView) DragLeaveEvent(e gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QUndoView::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(e))
	}
}

func (ptr *QUndoView) DragLeaveEventDefault(e gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QUndoView::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(e))
	}
}

func (ptr *QUndoView) ConnectDragMoveEvent(f func(e *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QUndoView::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QUndoView) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QUndoView::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQUndoViewDragMoveEvent
func callbackQUndoViewDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QUndoView::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(e))
	} else {
		NewQUndoViewFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(e))
	}
}

func (ptr *QUndoView) DragMoveEvent(e gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QUndoView::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(e))
	}
}

func (ptr *QUndoView) DragMoveEventDefault(e gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QUndoView::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(e))
	}
}

func (ptr *QUndoView) ConnectDropEvent(f func(e *gui.QDropEvent)) {
	defer qt.Recovering("connect QUndoView::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QUndoView) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QUndoView::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQUndoViewDropEvent
func callbackQUndoViewDropEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QUndoView::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(e))
	} else {
		NewQUndoViewFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(e))
	}
}

func (ptr *QUndoView) DropEvent(e gui.QDropEvent_ITF) {
	defer qt.Recovering("QUndoView::dropEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(e))
	}
}

func (ptr *QUndoView) DropEventDefault(e gui.QDropEvent_ITF) {
	defer qt.Recovering("QUndoView::dropEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(e))
	}
}

func (ptr *QUndoView) ConnectMouseMoveEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QUndoView::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QUndoView) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QUndoView::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQUndoViewMouseMoveEvent
func callbackQUndoViewMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QUndoView::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQUndoViewFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QUndoView) MouseMoveEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QUndoView::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QUndoView) MouseMoveEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QUndoView::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QUndoView) ConnectMouseReleaseEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QUndoView::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QUndoView) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QUndoView::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQUndoViewMouseReleaseEvent
func callbackQUndoViewMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QUndoView::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQUndoViewFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QUndoView) MouseReleaseEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QUndoView::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QUndoView) MouseReleaseEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QUndoView::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QUndoView) ConnectPaintEvent(f func(e *gui.QPaintEvent)) {
	defer qt.Recovering("connect QUndoView::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QUndoView) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QUndoView::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQUndoViewPaintEvent
func callbackQUndoViewPaintEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QUndoView::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(e))
	} else {
		NewQUndoViewFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(e))
	}
}

func (ptr *QUndoView) PaintEvent(e gui.QPaintEvent_ITF) {
	defer qt.Recovering("QUndoView::paintEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(e))
	}
}

func (ptr *QUndoView) PaintEventDefault(e gui.QPaintEvent_ITF) {
	defer qt.Recovering("QUndoView::paintEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(e))
	}
}

func (ptr *QUndoView) ConnectResizeEvent(f func(e *gui.QResizeEvent)) {
	defer qt.Recovering("connect QUndoView::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QUndoView) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QUndoView::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQUndoViewResizeEvent
func callbackQUndoViewResizeEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QUndoView::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(e))
	} else {
		NewQUndoViewFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(e))
	}
}

func (ptr *QUndoView) ResizeEvent(e gui.QResizeEvent_ITF) {
	defer qt.Recovering("QUndoView::resizeEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(e))
	}
}

func (ptr *QUndoView) ResizeEventDefault(e gui.QResizeEvent_ITF) {
	defer qt.Recovering("QUndoView::resizeEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(e))
	}
}

func (ptr *QUndoView) ConnectRowsAboutToBeRemoved(f func(parent *core.QModelIndex, start int, end int)) {
	defer qt.Recovering("connect QUndoView::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "rowsAboutToBeRemoved", f)
	}
}

func (ptr *QUndoView) DisconnectRowsAboutToBeRemoved() {
	defer qt.Recovering("disconnect QUndoView::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "rowsAboutToBeRemoved")
	}
}

//export callbackQUndoViewRowsAboutToBeRemoved
func callbackQUndoViewRowsAboutToBeRemoved(ptr unsafe.Pointer, ptrName *C.char, parent unsafe.Pointer, start C.int, end C.int) {
	defer qt.Recovering("callback QUndoView::rowsAboutToBeRemoved")

	if signal := qt.GetSignal(C.GoString(ptrName), "rowsAboutToBeRemoved"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(start), int(end))
	} else {
		NewQUndoViewFromPointer(ptr).RowsAboutToBeRemovedDefault(core.NewQModelIndexFromPointer(parent), int(start), int(end))
	}
}

func (ptr *QUndoView) RowsAboutToBeRemoved(parent core.QModelIndex_ITF, start int, end int) {
	defer qt.Recovering("QUndoView::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {
		C.QUndoView_RowsAboutToBeRemoved(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(start), C.int(end))
	}
}

func (ptr *QUndoView) RowsAboutToBeRemovedDefault(parent core.QModelIndex_ITF, start int, end int) {
	defer qt.Recovering("QUndoView::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {
		C.QUndoView_RowsAboutToBeRemovedDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(start), C.int(end))
	}
}

func (ptr *QUndoView) ConnectRowsInserted(f func(parent *core.QModelIndex, start int, end int)) {
	defer qt.Recovering("connect QUndoView::rowsInserted")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "rowsInserted", f)
	}
}

func (ptr *QUndoView) DisconnectRowsInserted() {
	defer qt.Recovering("disconnect QUndoView::rowsInserted")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "rowsInserted")
	}
}

//export callbackQUndoViewRowsInserted
func callbackQUndoViewRowsInserted(ptr unsafe.Pointer, ptrName *C.char, parent unsafe.Pointer, start C.int, end C.int) {
	defer qt.Recovering("callback QUndoView::rowsInserted")

	if signal := qt.GetSignal(C.GoString(ptrName), "rowsInserted"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(start), int(end))
	} else {
		NewQUndoViewFromPointer(ptr).RowsInsertedDefault(core.NewQModelIndexFromPointer(parent), int(start), int(end))
	}
}

func (ptr *QUndoView) RowsInserted(parent core.QModelIndex_ITF, start int, end int) {
	defer qt.Recovering("QUndoView::rowsInserted")

	if ptr.Pointer() != nil {
		C.QUndoView_RowsInserted(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(start), C.int(end))
	}
}

func (ptr *QUndoView) RowsInsertedDefault(parent core.QModelIndex_ITF, start int, end int) {
	defer qt.Recovering("QUndoView::rowsInserted")

	if ptr.Pointer() != nil {
		C.QUndoView_RowsInsertedDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(start), C.int(end))
	}
}

func (ptr *QUndoView) ConnectScrollTo(f func(index *core.QModelIndex, hint QAbstractItemView__ScrollHint)) {
	defer qt.Recovering("connect QUndoView::scrollTo")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "scrollTo", f)
	}
}

func (ptr *QUndoView) DisconnectScrollTo() {
	defer qt.Recovering("disconnect QUndoView::scrollTo")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "scrollTo")
	}
}

//export callbackQUndoViewScrollTo
func callbackQUndoViewScrollTo(ptr unsafe.Pointer, ptrName *C.char, index unsafe.Pointer, hint C.int) {
	defer qt.Recovering("callback QUndoView::scrollTo")

	if signal := qt.GetSignal(C.GoString(ptrName), "scrollTo"); signal != nil {
		signal.(func(*core.QModelIndex, QAbstractItemView__ScrollHint))(core.NewQModelIndexFromPointer(index), QAbstractItemView__ScrollHint(hint))
	} else {
		NewQUndoViewFromPointer(ptr).ScrollToDefault(core.NewQModelIndexFromPointer(index), QAbstractItemView__ScrollHint(hint))
	}
}

func (ptr *QUndoView) ScrollTo(index core.QModelIndex_ITF, hint QAbstractItemView__ScrollHint) {
	defer qt.Recovering("QUndoView::scrollTo")

	if ptr.Pointer() != nil {
		C.QUndoView_ScrollTo(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(hint))
	}
}

func (ptr *QUndoView) ScrollToDefault(index core.QModelIndex_ITF, hint QAbstractItemView__ScrollHint) {
	defer qt.Recovering("QUndoView::scrollTo")

	if ptr.Pointer() != nil {
		C.QUndoView_ScrollToDefault(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(hint))
	}
}

func (ptr *QUndoView) ConnectSetSelection(f func(rect *core.QRect, command core.QItemSelectionModel__SelectionFlag)) {
	defer qt.Recovering("connect QUndoView::setSelection")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setSelection", f)
	}
}

func (ptr *QUndoView) DisconnectSetSelection() {
	defer qt.Recovering("disconnect QUndoView::setSelection")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setSelection")
	}
}

//export callbackQUndoViewSetSelection
func callbackQUndoViewSetSelection(ptr unsafe.Pointer, ptrName *C.char, rect unsafe.Pointer, command C.int) {
	defer qt.Recovering("callback QUndoView::setSelection")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSelection"); signal != nil {
		signal.(func(*core.QRect, core.QItemSelectionModel__SelectionFlag))(core.NewQRectFromPointer(rect), core.QItemSelectionModel__SelectionFlag(command))
	} else {
		NewQUndoViewFromPointer(ptr).SetSelectionDefault(core.NewQRectFromPointer(rect), core.QItemSelectionModel__SelectionFlag(command))
	}
}

func (ptr *QUndoView) SetSelection(rect core.QRect_ITF, command core.QItemSelectionModel__SelectionFlag) {
	defer qt.Recovering("QUndoView::setSelection")

	if ptr.Pointer() != nil {
		C.QUndoView_SetSelection(ptr.Pointer(), core.PointerFromQRect(rect), C.int(command))
	}
}

func (ptr *QUndoView) SetSelectionDefault(rect core.QRect_ITF, command core.QItemSelectionModel__SelectionFlag) {
	defer qt.Recovering("QUndoView::setSelection")

	if ptr.Pointer() != nil {
		C.QUndoView_SetSelectionDefault(ptr.Pointer(), core.PointerFromQRect(rect), C.int(command))
	}
}

func (ptr *QUndoView) ConnectStartDrag(f func(supportedActions core.Qt__DropAction)) {
	defer qt.Recovering("connect QUndoView::startDrag")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "startDrag", f)
	}
}

func (ptr *QUndoView) DisconnectStartDrag() {
	defer qt.Recovering("disconnect QUndoView::startDrag")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "startDrag")
	}
}

//export callbackQUndoViewStartDrag
func callbackQUndoViewStartDrag(ptr unsafe.Pointer, ptrName *C.char, supportedActions C.int) {
	defer qt.Recovering("callback QUndoView::startDrag")

	if signal := qt.GetSignal(C.GoString(ptrName), "startDrag"); signal != nil {
		signal.(func(core.Qt__DropAction))(core.Qt__DropAction(supportedActions))
	} else {
		NewQUndoViewFromPointer(ptr).StartDragDefault(core.Qt__DropAction(supportedActions))
	}
}

func (ptr *QUndoView) StartDrag(supportedActions core.Qt__DropAction) {
	defer qt.Recovering("QUndoView::startDrag")

	if ptr.Pointer() != nil {
		C.QUndoView_StartDrag(ptr.Pointer(), C.int(supportedActions))
	}
}

func (ptr *QUndoView) StartDragDefault(supportedActions core.Qt__DropAction) {
	defer qt.Recovering("QUndoView::startDrag")

	if ptr.Pointer() != nil {
		C.QUndoView_StartDragDefault(ptr.Pointer(), C.int(supportedActions))
	}
}

func (ptr *QUndoView) ConnectTimerEvent(f func(e *core.QTimerEvent)) {
	defer qt.Recovering("connect QUndoView::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QUndoView) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QUndoView::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQUndoViewTimerEvent
func callbackQUndoViewTimerEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QUndoView::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(e))
	} else {
		NewQUndoViewFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(e))
	}
}

func (ptr *QUndoView) TimerEvent(e core.QTimerEvent_ITF) {
	defer qt.Recovering("QUndoView::timerEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(e))
	}
}

func (ptr *QUndoView) TimerEventDefault(e core.QTimerEvent_ITF) {
	defer qt.Recovering("QUndoView::timerEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(e))
	}
}

func (ptr *QUndoView) ConnectUpdateGeometries(f func()) {
	defer qt.Recovering("connect QUndoView::updateGeometries")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updateGeometries", f)
	}
}

func (ptr *QUndoView) DisconnectUpdateGeometries() {
	defer qt.Recovering("disconnect QUndoView::updateGeometries")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updateGeometries")
	}
}

//export callbackQUndoViewUpdateGeometries
func callbackQUndoViewUpdateGeometries(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QUndoView::updateGeometries")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateGeometries"); signal != nil {
		signal.(func())()
	} else {
		NewQUndoViewFromPointer(ptr).UpdateGeometriesDefault()
	}
}

func (ptr *QUndoView) UpdateGeometries() {
	defer qt.Recovering("QUndoView::updateGeometries")

	if ptr.Pointer() != nil {
		C.QUndoView_UpdateGeometries(ptr.Pointer())
	}
}

func (ptr *QUndoView) UpdateGeometriesDefault() {
	defer qt.Recovering("QUndoView::updateGeometries")

	if ptr.Pointer() != nil {
		C.QUndoView_UpdateGeometriesDefault(ptr.Pointer())
	}
}

func (ptr *QUndoView) ConnectCloseEditor(f func(editor *QWidget, hint QAbstractItemDelegate__EndEditHint)) {
	defer qt.Recovering("connect QUndoView::closeEditor")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEditor", f)
	}
}

func (ptr *QUndoView) DisconnectCloseEditor() {
	defer qt.Recovering("disconnect QUndoView::closeEditor")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEditor")
	}
}

//export callbackQUndoViewCloseEditor
func callbackQUndoViewCloseEditor(ptr unsafe.Pointer, ptrName *C.char, editor unsafe.Pointer, hint C.int) bool {
	defer qt.Recovering("callback QUndoView::closeEditor")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEditor"); signal != nil {
		signal.(func(*QWidget, QAbstractItemDelegate__EndEditHint))(NewQWidgetFromPointer(editor), QAbstractItemDelegate__EndEditHint(hint))
		return true
	}
	return false

}

func (ptr *QUndoView) CloseEditor(editor QWidget_ITF, hint QAbstractItemDelegate__EndEditHint) {
	defer qt.Recovering("QUndoView::closeEditor")

	if ptr.Pointer() != nil {
		C.QUndoView_CloseEditor(ptr.Pointer(), PointerFromQWidget(editor), C.int(hint))
	}
}

func (ptr *QUndoView) CloseEditorDefault(editor QWidget_ITF, hint QAbstractItemDelegate__EndEditHint) {
	defer qt.Recovering("QUndoView::closeEditor")

	if ptr.Pointer() != nil {
		C.QUndoView_CloseEditorDefault(ptr.Pointer(), PointerFromQWidget(editor), C.int(hint))
	}
}

func (ptr *QUndoView) ConnectCommitData(f func(editor *QWidget)) {
	defer qt.Recovering("connect QUndoView::commitData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "commitData", f)
	}
}

func (ptr *QUndoView) DisconnectCommitData() {
	defer qt.Recovering("disconnect QUndoView::commitData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "commitData")
	}
}

//export callbackQUndoViewCommitData
func callbackQUndoViewCommitData(ptr unsafe.Pointer, ptrName *C.char, editor unsafe.Pointer) bool {
	defer qt.Recovering("callback QUndoView::commitData")

	if signal := qt.GetSignal(C.GoString(ptrName), "commitData"); signal != nil {
		signal.(func(*QWidget))(NewQWidgetFromPointer(editor))
		return true
	}
	return false

}

func (ptr *QUndoView) CommitData(editor QWidget_ITF) {
	defer qt.Recovering("QUndoView::commitData")

	if ptr.Pointer() != nil {
		C.QUndoView_CommitData(ptr.Pointer(), PointerFromQWidget(editor))
	}
}

func (ptr *QUndoView) CommitDataDefault(editor QWidget_ITF) {
	defer qt.Recovering("QUndoView::commitData")

	if ptr.Pointer() != nil {
		C.QUndoView_CommitDataDefault(ptr.Pointer(), PointerFromQWidget(editor))
	}
}

func (ptr *QUndoView) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QUndoView::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QUndoView) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QUndoView::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQUndoViewDragEnterEvent
func callbackQUndoViewDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QUndoView::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQUndoViewFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QUndoView) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QUndoView::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QUndoView) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QUndoView::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QUndoView) ConnectEditorDestroyed(f func(editor *core.QObject)) {
	defer qt.Recovering("connect QUndoView::editorDestroyed")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "editorDestroyed", f)
	}
}

func (ptr *QUndoView) DisconnectEditorDestroyed() {
	defer qt.Recovering("disconnect QUndoView::editorDestroyed")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "editorDestroyed")
	}
}

//export callbackQUndoViewEditorDestroyed
func callbackQUndoViewEditorDestroyed(ptr unsafe.Pointer, ptrName *C.char, editor unsafe.Pointer) bool {
	defer qt.Recovering("callback QUndoView::editorDestroyed")

	if signal := qt.GetSignal(C.GoString(ptrName), "editorDestroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(editor))
		return true
	}
	return false

}

func (ptr *QUndoView) EditorDestroyed(editor core.QObject_ITF) {
	defer qt.Recovering("QUndoView::editorDestroyed")

	if ptr.Pointer() != nil {
		C.QUndoView_EditorDestroyed(ptr.Pointer(), core.PointerFromQObject(editor))
	}
}

func (ptr *QUndoView) EditorDestroyedDefault(editor core.QObject_ITF) {
	defer qt.Recovering("QUndoView::editorDestroyed")

	if ptr.Pointer() != nil {
		C.QUndoView_EditorDestroyedDefault(ptr.Pointer(), core.PointerFromQObject(editor))
	}
}

func (ptr *QUndoView) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QUndoView::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QUndoView) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QUndoView::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQUndoViewFocusInEvent
func callbackQUndoViewFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QUndoView::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQUndoViewFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QUndoView) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QUndoView::focusInEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QUndoView) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QUndoView::focusInEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QUndoView) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QUndoView::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QUndoView) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QUndoView::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQUndoViewFocusOutEvent
func callbackQUndoViewFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QUndoView::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQUndoViewFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QUndoView) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QUndoView::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QUndoView) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QUndoView::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QUndoView) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QUndoView::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QUndoView) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QUndoView::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQUndoViewInputMethodEvent
func callbackQUndoViewInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QUndoView::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQUndoViewFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QUndoView) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QUndoView::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QUndoView) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QUndoView::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QUndoView) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QUndoView::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QUndoView) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QUndoView::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQUndoViewKeyPressEvent
func callbackQUndoViewKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QUndoView::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQUndoViewFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QUndoView) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QUndoView::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QUndoView) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QUndoView::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QUndoView) ConnectKeyboardSearch(f func(search string)) {
	defer qt.Recovering("connect QUndoView::keyboardSearch")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyboardSearch", f)
	}
}

func (ptr *QUndoView) DisconnectKeyboardSearch() {
	defer qt.Recovering("disconnect QUndoView::keyboardSearch")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyboardSearch")
	}
}

//export callbackQUndoViewKeyboardSearch
func callbackQUndoViewKeyboardSearch(ptr unsafe.Pointer, ptrName *C.char, search *C.char) {
	defer qt.Recovering("callback QUndoView::keyboardSearch")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyboardSearch"); signal != nil {
		signal.(func(string))(C.GoString(search))
	} else {
		NewQUndoViewFromPointer(ptr).KeyboardSearchDefault(C.GoString(search))
	}
}

func (ptr *QUndoView) KeyboardSearch(search string) {
	defer qt.Recovering("QUndoView::keyboardSearch")

	if ptr.Pointer() != nil {
		C.QUndoView_KeyboardSearch(ptr.Pointer(), C.CString(search))
	}
}

func (ptr *QUndoView) KeyboardSearchDefault(search string) {
	defer qt.Recovering("QUndoView::keyboardSearch")

	if ptr.Pointer() != nil {
		C.QUndoView_KeyboardSearchDefault(ptr.Pointer(), C.CString(search))
	}
}

func (ptr *QUndoView) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QUndoView::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QUndoView) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QUndoView::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQUndoViewMouseDoubleClickEvent
func callbackQUndoViewMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QUndoView::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQUndoViewFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QUndoView) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QUndoView::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QUndoView) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QUndoView::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QUndoView) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QUndoView::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QUndoView) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QUndoView::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQUndoViewMousePressEvent
func callbackQUndoViewMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QUndoView::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQUndoViewFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QUndoView) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QUndoView::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QUndoView) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QUndoView::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QUndoView) ConnectReset(f func()) {
	defer qt.Recovering("connect QUndoView::reset")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "reset", f)
	}
}

func (ptr *QUndoView) DisconnectReset() {
	defer qt.Recovering("disconnect QUndoView::reset")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "reset")
	}
}

//export callbackQUndoViewReset
func callbackQUndoViewReset(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QUndoView::reset")

	if signal := qt.GetSignal(C.GoString(ptrName), "reset"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QUndoView) Reset() {
	defer qt.Recovering("QUndoView::reset")

	if ptr.Pointer() != nil {
		C.QUndoView_Reset(ptr.Pointer())
	}
}

func (ptr *QUndoView) ResetDefault() {
	defer qt.Recovering("QUndoView::reset")

	if ptr.Pointer() != nil {
		C.QUndoView_ResetDefault(ptr.Pointer())
	}
}

func (ptr *QUndoView) ConnectSelectAll(f func()) {
	defer qt.Recovering("connect QUndoView::selectAll")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "selectAll", f)
	}
}

func (ptr *QUndoView) DisconnectSelectAll() {
	defer qt.Recovering("disconnect QUndoView::selectAll")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "selectAll")
	}
}

//export callbackQUndoViewSelectAll
func callbackQUndoViewSelectAll(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QUndoView::selectAll")

	if signal := qt.GetSignal(C.GoString(ptrName), "selectAll"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QUndoView) SelectAll() {
	defer qt.Recovering("QUndoView::selectAll")

	if ptr.Pointer() != nil {
		C.QUndoView_SelectAll(ptr.Pointer())
	}
}

func (ptr *QUndoView) SelectAllDefault() {
	defer qt.Recovering("QUndoView::selectAll")

	if ptr.Pointer() != nil {
		C.QUndoView_SelectAllDefault(ptr.Pointer())
	}
}

func (ptr *QUndoView) ConnectSetModel(f func(model *core.QAbstractItemModel)) {
	defer qt.Recovering("connect QUndoView::setModel")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setModel", f)
	}
}

func (ptr *QUndoView) DisconnectSetModel() {
	defer qt.Recovering("disconnect QUndoView::setModel")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setModel")
	}
}

//export callbackQUndoViewSetModel
func callbackQUndoViewSetModel(ptr unsafe.Pointer, ptrName *C.char, model unsafe.Pointer) {
	defer qt.Recovering("callback QUndoView::setModel")

	if signal := qt.GetSignal(C.GoString(ptrName), "setModel"); signal != nil {
		signal.(func(*core.QAbstractItemModel))(core.NewQAbstractItemModelFromPointer(model))
	} else {
		NewQUndoViewFromPointer(ptr).SetModelDefault(core.NewQAbstractItemModelFromPointer(model))
	}
}

func (ptr *QUndoView) SetModel(model core.QAbstractItemModel_ITF) {
	defer qt.Recovering("QUndoView::setModel")

	if ptr.Pointer() != nil {
		C.QUndoView_SetModel(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QUndoView) SetModelDefault(model core.QAbstractItemModel_ITF) {
	defer qt.Recovering("QUndoView::setModel")

	if ptr.Pointer() != nil {
		C.QUndoView_SetModelDefault(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QUndoView) ConnectSetRootIndex(f func(index *core.QModelIndex)) {
	defer qt.Recovering("connect QUndoView::setRootIndex")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setRootIndex", f)
	}
}

func (ptr *QUndoView) DisconnectSetRootIndex() {
	defer qt.Recovering("disconnect QUndoView::setRootIndex")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setRootIndex")
	}
}

//export callbackQUndoViewSetRootIndex
func callbackQUndoViewSetRootIndex(ptr unsafe.Pointer, ptrName *C.char, index unsafe.Pointer) bool {
	defer qt.Recovering("callback QUndoView::setRootIndex")

	if signal := qt.GetSignal(C.GoString(ptrName), "setRootIndex"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
		return true
	}
	return false

}

func (ptr *QUndoView) SetRootIndex(index core.QModelIndex_ITF) {
	defer qt.Recovering("QUndoView::setRootIndex")

	if ptr.Pointer() != nil {
		C.QUndoView_SetRootIndex(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QUndoView) SetRootIndexDefault(index core.QModelIndex_ITF) {
	defer qt.Recovering("QUndoView::setRootIndex")

	if ptr.Pointer() != nil {
		C.QUndoView_SetRootIndexDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QUndoView) ConnectSetSelectionModel(f func(selectionModel *core.QItemSelectionModel)) {
	defer qt.Recovering("connect QUndoView::setSelectionModel")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setSelectionModel", f)
	}
}

func (ptr *QUndoView) DisconnectSetSelectionModel() {
	defer qt.Recovering("disconnect QUndoView::setSelectionModel")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setSelectionModel")
	}
}

//export callbackQUndoViewSetSelectionModel
func callbackQUndoViewSetSelectionModel(ptr unsafe.Pointer, ptrName *C.char, selectionModel unsafe.Pointer) {
	defer qt.Recovering("callback QUndoView::setSelectionModel")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSelectionModel"); signal != nil {
		signal.(func(*core.QItemSelectionModel))(core.NewQItemSelectionModelFromPointer(selectionModel))
	} else {
		NewQUndoViewFromPointer(ptr).SetSelectionModelDefault(core.NewQItemSelectionModelFromPointer(selectionModel))
	}
}

func (ptr *QUndoView) SetSelectionModel(selectionModel core.QItemSelectionModel_ITF) {
	defer qt.Recovering("QUndoView::setSelectionModel")

	if ptr.Pointer() != nil {
		C.QUndoView_SetSelectionModel(ptr.Pointer(), core.PointerFromQItemSelectionModel(selectionModel))
	}
}

func (ptr *QUndoView) SetSelectionModelDefault(selectionModel core.QItemSelectionModel_ITF) {
	defer qt.Recovering("QUndoView::setSelectionModel")

	if ptr.Pointer() != nil {
		C.QUndoView_SetSelectionModelDefault(ptr.Pointer(), core.PointerFromQItemSelectionModel(selectionModel))
	}
}

func (ptr *QUndoView) ConnectContextMenuEvent(f func(e *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QUndoView::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QUndoView) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QUndoView::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQUndoViewContextMenuEvent
func callbackQUndoViewContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QUndoView::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
	} else {
		NewQUndoViewFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(e))
	}
}

func (ptr *QUndoView) ContextMenuEvent(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QUndoView::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

func (ptr *QUndoView) ContextMenuEventDefault(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QUndoView::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

func (ptr *QUndoView) ConnectScrollContentsBy(f func(dx int, dy int)) {
	defer qt.Recovering("connect QUndoView::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "scrollContentsBy", f)
	}
}

func (ptr *QUndoView) DisconnectScrollContentsBy() {
	defer qt.Recovering("disconnect QUndoView::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "scrollContentsBy")
	}
}

//export callbackQUndoViewScrollContentsBy
func callbackQUndoViewScrollContentsBy(ptr unsafe.Pointer, ptrName *C.char, dx C.int, dy C.int) {
	defer qt.Recovering("callback QUndoView::scrollContentsBy")

	if signal := qt.GetSignal(C.GoString(ptrName), "scrollContentsBy"); signal != nil {
		signal.(func(int, int))(int(dx), int(dy))
	} else {
		NewQUndoViewFromPointer(ptr).ScrollContentsByDefault(int(dx), int(dy))
	}
}

func (ptr *QUndoView) ScrollContentsBy(dx int, dy int) {
	defer qt.Recovering("QUndoView::scrollContentsBy")

	if ptr.Pointer() != nil {
		C.QUndoView_ScrollContentsBy(ptr.Pointer(), C.int(dx), C.int(dy))
	}
}

func (ptr *QUndoView) ScrollContentsByDefault(dx int, dy int) {
	defer qt.Recovering("QUndoView::scrollContentsBy")

	if ptr.Pointer() != nil {
		C.QUndoView_ScrollContentsByDefault(ptr.Pointer(), C.int(dx), C.int(dy))
	}
}

func (ptr *QUndoView) ConnectSetupViewport(f func(viewport *QWidget)) {
	defer qt.Recovering("connect QUndoView::setupViewport")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setupViewport", f)
	}
}

func (ptr *QUndoView) DisconnectSetupViewport() {
	defer qt.Recovering("disconnect QUndoView::setupViewport")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setupViewport")
	}
}

//export callbackQUndoViewSetupViewport
func callbackQUndoViewSetupViewport(ptr unsafe.Pointer, ptrName *C.char, viewport unsafe.Pointer) {
	defer qt.Recovering("callback QUndoView::setupViewport")

	if signal := qt.GetSignal(C.GoString(ptrName), "setupViewport"); signal != nil {
		signal.(func(*QWidget))(NewQWidgetFromPointer(viewport))
	} else {
		NewQUndoViewFromPointer(ptr).SetupViewportDefault(NewQWidgetFromPointer(viewport))
	}
}

func (ptr *QUndoView) SetupViewport(viewport QWidget_ITF) {
	defer qt.Recovering("QUndoView::setupViewport")

	if ptr.Pointer() != nil {
		C.QUndoView_SetupViewport(ptr.Pointer(), PointerFromQWidget(viewport))
	}
}

func (ptr *QUndoView) SetupViewportDefault(viewport QWidget_ITF) {
	defer qt.Recovering("QUndoView::setupViewport")

	if ptr.Pointer() != nil {
		C.QUndoView_SetupViewportDefault(ptr.Pointer(), PointerFromQWidget(viewport))
	}
}

func (ptr *QUndoView) ConnectWheelEvent(f func(e *gui.QWheelEvent)) {
	defer qt.Recovering("connect QUndoView::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QUndoView) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QUndoView::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQUndoViewWheelEvent
func callbackQUndoViewWheelEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QUndoView::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
	} else {
		NewQUndoViewFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(e))
	}
}

func (ptr *QUndoView) WheelEvent(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QUndoView::wheelEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QUndoView) WheelEventDefault(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QUndoView::wheelEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QUndoView) ConnectChangeEvent(f func(ev *core.QEvent)) {
	defer qt.Recovering("connect QUndoView::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QUndoView) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QUndoView::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQUndoViewChangeEvent
func callbackQUndoViewChangeEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QUndoView::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
	} else {
		NewQUndoViewFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(ev))
	}
}

func (ptr *QUndoView) ChangeEvent(ev core.QEvent_ITF) {
	defer qt.Recovering("QUndoView::changeEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QUndoView) ChangeEventDefault(ev core.QEvent_ITF) {
	defer qt.Recovering("QUndoView::changeEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QUndoView) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QUndoView::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QUndoView) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QUndoView::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQUndoViewActionEvent
func callbackQUndoViewActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QUndoView::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQUndoViewFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QUndoView) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QUndoView::actionEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QUndoView) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QUndoView::actionEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QUndoView) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QUndoView::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QUndoView) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QUndoView::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQUndoViewEnterEvent
func callbackQUndoViewEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QUndoView::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQUndoViewFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QUndoView) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QUndoView::enterEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QUndoView) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QUndoView::enterEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QUndoView) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QUndoView::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QUndoView) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QUndoView::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQUndoViewHideEvent
func callbackQUndoViewHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QUndoView::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQUndoViewFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QUndoView) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QUndoView::hideEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QUndoView) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QUndoView::hideEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QUndoView) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QUndoView::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QUndoView) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QUndoView::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQUndoViewLeaveEvent
func callbackQUndoViewLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QUndoView::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQUndoViewFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QUndoView) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QUndoView::leaveEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QUndoView) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QUndoView::leaveEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QUndoView) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QUndoView::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QUndoView) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QUndoView::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQUndoViewMoveEvent
func callbackQUndoViewMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QUndoView::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQUndoViewFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QUndoView) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QUndoView::moveEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QUndoView) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QUndoView::moveEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QUndoView) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QUndoView::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QUndoView) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QUndoView::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQUndoViewSetVisible
func callbackQUndoViewSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QUndoView::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QUndoView) SetVisible(visible bool) {
	defer qt.Recovering("QUndoView::setVisible")

	if ptr.Pointer() != nil {
		C.QUndoView_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QUndoView) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QUndoView::setVisible")

	if ptr.Pointer() != nil {
		C.QUndoView_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QUndoView) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QUndoView::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QUndoView) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QUndoView::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQUndoViewShowEvent
func callbackQUndoViewShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QUndoView::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQUndoViewFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QUndoView) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QUndoView::showEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QUndoView) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QUndoView::showEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QUndoView) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QUndoView::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QUndoView) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QUndoView::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQUndoViewCloseEvent
func callbackQUndoViewCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QUndoView::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQUndoViewFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QUndoView) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QUndoView::closeEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QUndoView) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QUndoView::closeEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QUndoView) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QUndoView::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QUndoView) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QUndoView::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQUndoViewInitPainter
func callbackQUndoViewInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QUndoView::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQUndoViewFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QUndoView) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QUndoView::initPainter")

	if ptr.Pointer() != nil {
		C.QUndoView_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QUndoView) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QUndoView::initPainter")

	if ptr.Pointer() != nil {
		C.QUndoView_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QUndoView) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QUndoView::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QUndoView) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QUndoView::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQUndoViewKeyReleaseEvent
func callbackQUndoViewKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QUndoView::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQUndoViewFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QUndoView) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QUndoView::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QUndoView) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QUndoView::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QUndoView) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QUndoView::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QUndoView) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QUndoView::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQUndoViewTabletEvent
func callbackQUndoViewTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QUndoView::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQUndoViewFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QUndoView) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QUndoView::tabletEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QUndoView) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QUndoView::tabletEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QUndoView) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QUndoView::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QUndoView) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QUndoView::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQUndoViewChildEvent
func callbackQUndoViewChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QUndoView::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQUndoViewFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QUndoView) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QUndoView::childEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QUndoView) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QUndoView::childEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QUndoView) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QUndoView::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QUndoView) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QUndoView::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQUndoViewCustomEvent
func callbackQUndoViewCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QUndoView::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQUndoViewFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QUndoView) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QUndoView::customEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QUndoView) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QUndoView::customEvent")

	if ptr.Pointer() != nil {
		C.QUndoView_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
