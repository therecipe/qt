package help

//#include "help.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"unsafe"
)

type QHelpIndexWidget struct {
	widgets.QListView
}

type QHelpIndexWidget_ITF interface {
	widgets.QListView_ITF
	QHelpIndexWidget_PTR() *QHelpIndexWidget
}

func PointerFromQHelpIndexWidget(ptr QHelpIndexWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpIndexWidget_PTR().Pointer()
	}
	return nil
}

func NewQHelpIndexWidgetFromPointer(ptr unsafe.Pointer) *QHelpIndexWidget {
	var n = new(QHelpIndexWidget)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QHelpIndexWidget_") {
		n.SetObjectName("QHelpIndexWidget_" + qt.Identifier())
	}
	return n
}

func (ptr *QHelpIndexWidget) QHelpIndexWidget_PTR() *QHelpIndexWidget {
	return ptr
}

func (ptr *QHelpIndexWidget) ActivateCurrentItem() {
	defer qt.Recovering("QHelpIndexWidget::activateCurrentItem")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ActivateCurrentItem(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) FilterIndices(filter string, wildcard string) {
	defer qt.Recovering("QHelpIndexWidget::filterIndices")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_FilterIndices(ptr.Pointer(), C.CString(filter), C.CString(wildcard))
	}
}

func (ptr *QHelpIndexWidget) ConnectLinkActivated(f func(link *core.QUrl, keyword string)) {
	defer qt.Recovering("connect QHelpIndexWidget::linkActivated")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ConnectLinkActivated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "linkActivated", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectLinkActivated() {
	defer qt.Recovering("disconnect QHelpIndexWidget::linkActivated")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DisconnectLinkActivated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "linkActivated")
	}
}

//export callbackQHelpIndexWidgetLinkActivated
func callbackQHelpIndexWidgetLinkActivated(ptr unsafe.Pointer, ptrName *C.char, link unsafe.Pointer, keyword *C.char) {
	defer qt.Recovering("callback QHelpIndexWidget::linkActivated")

	if signal := qt.GetSignal(C.GoString(ptrName), "linkActivated"); signal != nil {
		signal.(func(*core.QUrl, string))(core.NewQUrlFromPointer(link), C.GoString(keyword))
	}

}

func (ptr *QHelpIndexWidget) LinkActivated(link core.QUrl_ITF, keyword string) {
	defer qt.Recovering("QHelpIndexWidget::linkActivated")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_LinkActivated(ptr.Pointer(), core.PointerFromQUrl(link), C.CString(keyword))
	}
}

func (ptr *QHelpIndexWidget) ConnectCurrentChanged(f func(current *core.QModelIndex, previous *core.QModelIndex)) {
	defer qt.Recovering("connect QHelpIndexWidget::currentChanged")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "currentChanged", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectCurrentChanged() {
	defer qt.Recovering("disconnect QHelpIndexWidget::currentChanged")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "currentChanged")
	}
}

//export callbackQHelpIndexWidgetCurrentChanged
func callbackQHelpIndexWidgetCurrentChanged(ptr unsafe.Pointer, ptrName *C.char, current unsafe.Pointer, previous unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::currentChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "currentChanged"); signal != nil {
		signal.(func(*core.QModelIndex, *core.QModelIndex))(core.NewQModelIndexFromPointer(current), core.NewQModelIndexFromPointer(previous))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).CurrentChangedDefault(core.NewQModelIndexFromPointer(current), core.NewQModelIndexFromPointer(previous))
	}
}

func (ptr *QHelpIndexWidget) CurrentChanged(current core.QModelIndex_ITF, previous core.QModelIndex_ITF) {
	defer qt.Recovering("QHelpIndexWidget::currentChanged")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_CurrentChanged(ptr.Pointer(), core.PointerFromQModelIndex(current), core.PointerFromQModelIndex(previous))
	}
}

func (ptr *QHelpIndexWidget) CurrentChangedDefault(current core.QModelIndex_ITF, previous core.QModelIndex_ITF) {
	defer qt.Recovering("QHelpIndexWidget::currentChanged")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_CurrentChangedDefault(ptr.Pointer(), core.PointerFromQModelIndex(current), core.PointerFromQModelIndex(previous))
	}
}

func (ptr *QHelpIndexWidget) ConnectDragLeaveEvent(f func(e *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQHelpIndexWidgetDragLeaveEvent
func callbackQHelpIndexWidgetDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) DragLeaveEvent(e gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(e))
	}
}

func (ptr *QHelpIndexWidget) DragLeaveEventDefault(e gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(e))
	}
}

func (ptr *QHelpIndexWidget) ConnectDragMoveEvent(f func(e *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQHelpIndexWidgetDragMoveEvent
func callbackQHelpIndexWidgetDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) DragMoveEvent(e gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(e))
	}
}

func (ptr *QHelpIndexWidget) DragMoveEventDefault(e gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(e))
	}
}

func (ptr *QHelpIndexWidget) ConnectDropEvent(f func(e *gui.QDropEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQHelpIndexWidgetDropEvent
func callbackQHelpIndexWidgetDropEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) DropEvent(e gui.QDropEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(e))
	}
}

func (ptr *QHelpIndexWidget) DropEventDefault(e gui.QDropEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(e))
	}
}

func (ptr *QHelpIndexWidget) ConnectMouseMoveEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQHelpIndexWidgetMouseMoveEvent
func callbackQHelpIndexWidgetMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) MouseMoveEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QHelpIndexWidget) MouseMoveEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QHelpIndexWidget) ConnectMouseReleaseEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQHelpIndexWidgetMouseReleaseEvent
func callbackQHelpIndexWidgetMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) MouseReleaseEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QHelpIndexWidget) MouseReleaseEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QHelpIndexWidget) ConnectPaintEvent(f func(e *gui.QPaintEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQHelpIndexWidgetPaintEvent
func callbackQHelpIndexWidgetPaintEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) PaintEvent(e gui.QPaintEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::paintEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(e))
	}
}

func (ptr *QHelpIndexWidget) PaintEventDefault(e gui.QPaintEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::paintEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(e))
	}
}

func (ptr *QHelpIndexWidget) ConnectResizeEvent(f func(e *gui.QResizeEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQHelpIndexWidgetResizeEvent
func callbackQHelpIndexWidgetResizeEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) ResizeEvent(e gui.QResizeEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(e))
	}
}

func (ptr *QHelpIndexWidget) ResizeEventDefault(e gui.QResizeEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(e))
	}
}

func (ptr *QHelpIndexWidget) ConnectRowsAboutToBeRemoved(f func(parent *core.QModelIndex, start int, end int)) {
	defer qt.Recovering("connect QHelpIndexWidget::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "rowsAboutToBeRemoved", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectRowsAboutToBeRemoved() {
	defer qt.Recovering("disconnect QHelpIndexWidget::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "rowsAboutToBeRemoved")
	}
}

//export callbackQHelpIndexWidgetRowsAboutToBeRemoved
func callbackQHelpIndexWidgetRowsAboutToBeRemoved(ptr unsafe.Pointer, ptrName *C.char, parent unsafe.Pointer, start C.int, end C.int) {
	defer qt.Recovering("callback QHelpIndexWidget::rowsAboutToBeRemoved")

	if signal := qt.GetSignal(C.GoString(ptrName), "rowsAboutToBeRemoved"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(start), int(end))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).RowsAboutToBeRemovedDefault(core.NewQModelIndexFromPointer(parent), int(start), int(end))
	}
}

func (ptr *QHelpIndexWidget) RowsAboutToBeRemoved(parent core.QModelIndex_ITF, start int, end int) {
	defer qt.Recovering("QHelpIndexWidget::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_RowsAboutToBeRemoved(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(start), C.int(end))
	}
}

func (ptr *QHelpIndexWidget) RowsAboutToBeRemovedDefault(parent core.QModelIndex_ITF, start int, end int) {
	defer qt.Recovering("QHelpIndexWidget::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_RowsAboutToBeRemovedDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(start), C.int(end))
	}
}

func (ptr *QHelpIndexWidget) ConnectRowsInserted(f func(parent *core.QModelIndex, start int, end int)) {
	defer qt.Recovering("connect QHelpIndexWidget::rowsInserted")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "rowsInserted", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectRowsInserted() {
	defer qt.Recovering("disconnect QHelpIndexWidget::rowsInserted")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "rowsInserted")
	}
}

//export callbackQHelpIndexWidgetRowsInserted
func callbackQHelpIndexWidgetRowsInserted(ptr unsafe.Pointer, ptrName *C.char, parent unsafe.Pointer, start C.int, end C.int) {
	defer qt.Recovering("callback QHelpIndexWidget::rowsInserted")

	if signal := qt.GetSignal(C.GoString(ptrName), "rowsInserted"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(start), int(end))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).RowsInsertedDefault(core.NewQModelIndexFromPointer(parent), int(start), int(end))
	}
}

func (ptr *QHelpIndexWidget) RowsInserted(parent core.QModelIndex_ITF, start int, end int) {
	defer qt.Recovering("QHelpIndexWidget::rowsInserted")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_RowsInserted(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(start), C.int(end))
	}
}

func (ptr *QHelpIndexWidget) RowsInsertedDefault(parent core.QModelIndex_ITF, start int, end int) {
	defer qt.Recovering("QHelpIndexWidget::rowsInserted")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_RowsInsertedDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(start), C.int(end))
	}
}

func (ptr *QHelpIndexWidget) ConnectScrollTo(f func(index *core.QModelIndex, hint widgets.QAbstractItemView__ScrollHint)) {
	defer qt.Recovering("connect QHelpIndexWidget::scrollTo")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "scrollTo", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectScrollTo() {
	defer qt.Recovering("disconnect QHelpIndexWidget::scrollTo")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "scrollTo")
	}
}

//export callbackQHelpIndexWidgetScrollTo
func callbackQHelpIndexWidgetScrollTo(ptr unsafe.Pointer, ptrName *C.char, index unsafe.Pointer, hint C.int) {
	defer qt.Recovering("callback QHelpIndexWidget::scrollTo")

	if signal := qt.GetSignal(C.GoString(ptrName), "scrollTo"); signal != nil {
		signal.(func(*core.QModelIndex, widgets.QAbstractItemView__ScrollHint))(core.NewQModelIndexFromPointer(index), widgets.QAbstractItemView__ScrollHint(hint))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ScrollToDefault(core.NewQModelIndexFromPointer(index), widgets.QAbstractItemView__ScrollHint(hint))
	}
}

func (ptr *QHelpIndexWidget) ScrollTo(index core.QModelIndex_ITF, hint widgets.QAbstractItemView__ScrollHint) {
	defer qt.Recovering("QHelpIndexWidget::scrollTo")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ScrollTo(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(hint))
	}
}

func (ptr *QHelpIndexWidget) ScrollToDefault(index core.QModelIndex_ITF, hint widgets.QAbstractItemView__ScrollHint) {
	defer qt.Recovering("QHelpIndexWidget::scrollTo")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ScrollToDefault(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(hint))
	}
}

func (ptr *QHelpIndexWidget) ConnectSetSelection(f func(rect *core.QRect, command core.QItemSelectionModel__SelectionFlag)) {
	defer qt.Recovering("connect QHelpIndexWidget::setSelection")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setSelection", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSetSelection() {
	defer qt.Recovering("disconnect QHelpIndexWidget::setSelection")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setSelection")
	}
}

//export callbackQHelpIndexWidgetSetSelection
func callbackQHelpIndexWidgetSetSelection(ptr unsafe.Pointer, ptrName *C.char, rect unsafe.Pointer, command C.int) {
	defer qt.Recovering("callback QHelpIndexWidget::setSelection")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSelection"); signal != nil {
		signal.(func(*core.QRect, core.QItemSelectionModel__SelectionFlag))(core.NewQRectFromPointer(rect), core.QItemSelectionModel__SelectionFlag(command))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetSelectionDefault(core.NewQRectFromPointer(rect), core.QItemSelectionModel__SelectionFlag(command))
	}
}

func (ptr *QHelpIndexWidget) SetSelection(rect core.QRect_ITF, command core.QItemSelectionModel__SelectionFlag) {
	defer qt.Recovering("QHelpIndexWidget::setSelection")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetSelection(ptr.Pointer(), core.PointerFromQRect(rect), C.int(command))
	}
}

func (ptr *QHelpIndexWidget) SetSelectionDefault(rect core.QRect_ITF, command core.QItemSelectionModel__SelectionFlag) {
	defer qt.Recovering("QHelpIndexWidget::setSelection")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetSelectionDefault(ptr.Pointer(), core.PointerFromQRect(rect), C.int(command))
	}
}

func (ptr *QHelpIndexWidget) ConnectStartDrag(f func(supportedActions core.Qt__DropAction)) {
	defer qt.Recovering("connect QHelpIndexWidget::startDrag")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "startDrag", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectStartDrag() {
	defer qt.Recovering("disconnect QHelpIndexWidget::startDrag")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "startDrag")
	}
}

//export callbackQHelpIndexWidgetStartDrag
func callbackQHelpIndexWidgetStartDrag(ptr unsafe.Pointer, ptrName *C.char, supportedActions C.int) {
	defer qt.Recovering("callback QHelpIndexWidget::startDrag")

	if signal := qt.GetSignal(C.GoString(ptrName), "startDrag"); signal != nil {
		signal.(func(core.Qt__DropAction))(core.Qt__DropAction(supportedActions))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).StartDragDefault(core.Qt__DropAction(supportedActions))
	}
}

func (ptr *QHelpIndexWidget) StartDrag(supportedActions core.Qt__DropAction) {
	defer qt.Recovering("QHelpIndexWidget::startDrag")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_StartDrag(ptr.Pointer(), C.int(supportedActions))
	}
}

func (ptr *QHelpIndexWidget) StartDragDefault(supportedActions core.Qt__DropAction) {
	defer qt.Recovering("QHelpIndexWidget::startDrag")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_StartDragDefault(ptr.Pointer(), C.int(supportedActions))
	}
}

func (ptr *QHelpIndexWidget) ConnectTimerEvent(f func(e *core.QTimerEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQHelpIndexWidgetTimerEvent
func callbackQHelpIndexWidgetTimerEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) TimerEvent(e core.QTimerEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::timerEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(e))
	}
}

func (ptr *QHelpIndexWidget) TimerEventDefault(e core.QTimerEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::timerEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(e))
	}
}

func (ptr *QHelpIndexWidget) ConnectUpdateGeometries(f func()) {
	defer qt.Recovering("connect QHelpIndexWidget::updateGeometries")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updateGeometries", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectUpdateGeometries() {
	defer qt.Recovering("disconnect QHelpIndexWidget::updateGeometries")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updateGeometries")
	}
}

//export callbackQHelpIndexWidgetUpdateGeometries
func callbackQHelpIndexWidgetUpdateGeometries(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QHelpIndexWidget::updateGeometries")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateGeometries"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).UpdateGeometriesDefault()
	}
}

func (ptr *QHelpIndexWidget) UpdateGeometries() {
	defer qt.Recovering("QHelpIndexWidget::updateGeometries")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_UpdateGeometries(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) UpdateGeometriesDefault() {
	defer qt.Recovering("QHelpIndexWidget::updateGeometries")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_UpdateGeometriesDefault(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) ConnectCloseEditor(f func(editor *widgets.QWidget, hint widgets.QAbstractItemDelegate__EndEditHint)) {
	defer qt.Recovering("connect QHelpIndexWidget::closeEditor")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEditor", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectCloseEditor() {
	defer qt.Recovering("disconnect QHelpIndexWidget::closeEditor")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEditor")
	}
}

//export callbackQHelpIndexWidgetCloseEditor
func callbackQHelpIndexWidgetCloseEditor(ptr unsafe.Pointer, ptrName *C.char, editor unsafe.Pointer, hint C.int) bool {
	defer qt.Recovering("callback QHelpIndexWidget::closeEditor")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEditor"); signal != nil {
		signal.(func(*widgets.QWidget, widgets.QAbstractItemDelegate__EndEditHint))(widgets.NewQWidgetFromPointer(editor), widgets.QAbstractItemDelegate__EndEditHint(hint))
		return true
	}
	return false

}

func (ptr *QHelpIndexWidget) CloseEditor(editor widgets.QWidget_ITF, hint widgets.QAbstractItemDelegate__EndEditHint) {
	defer qt.Recovering("QHelpIndexWidget::closeEditor")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_CloseEditor(ptr.Pointer(), widgets.PointerFromQWidget(editor), C.int(hint))
	}
}

func (ptr *QHelpIndexWidget) CloseEditorDefault(editor widgets.QWidget_ITF, hint widgets.QAbstractItemDelegate__EndEditHint) {
	defer qt.Recovering("QHelpIndexWidget::closeEditor")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_CloseEditorDefault(ptr.Pointer(), widgets.PointerFromQWidget(editor), C.int(hint))
	}
}

func (ptr *QHelpIndexWidget) ConnectCommitData(f func(editor *widgets.QWidget)) {
	defer qt.Recovering("connect QHelpIndexWidget::commitData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "commitData", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectCommitData() {
	defer qt.Recovering("disconnect QHelpIndexWidget::commitData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "commitData")
	}
}

//export callbackQHelpIndexWidgetCommitData
func callbackQHelpIndexWidgetCommitData(ptr unsafe.Pointer, ptrName *C.char, editor unsafe.Pointer) bool {
	defer qt.Recovering("callback QHelpIndexWidget::commitData")

	if signal := qt.GetSignal(C.GoString(ptrName), "commitData"); signal != nil {
		signal.(func(*widgets.QWidget))(widgets.NewQWidgetFromPointer(editor))
		return true
	}
	return false

}

func (ptr *QHelpIndexWidget) CommitData(editor widgets.QWidget_ITF) {
	defer qt.Recovering("QHelpIndexWidget::commitData")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_CommitData(ptr.Pointer(), widgets.PointerFromQWidget(editor))
	}
}

func (ptr *QHelpIndexWidget) CommitDataDefault(editor widgets.QWidget_ITF) {
	defer qt.Recovering("QHelpIndexWidget::commitData")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_CommitDataDefault(ptr.Pointer(), widgets.PointerFromQWidget(editor))
	}
}

func (ptr *QHelpIndexWidget) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQHelpIndexWidgetDragEnterEvent
func callbackQHelpIndexWidgetDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QHelpIndexWidget) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectEditorDestroyed(f func(editor *core.QObject)) {
	defer qt.Recovering("connect QHelpIndexWidget::editorDestroyed")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "editorDestroyed", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectEditorDestroyed() {
	defer qt.Recovering("disconnect QHelpIndexWidget::editorDestroyed")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "editorDestroyed")
	}
}

//export callbackQHelpIndexWidgetEditorDestroyed
func callbackQHelpIndexWidgetEditorDestroyed(ptr unsafe.Pointer, ptrName *C.char, editor unsafe.Pointer) bool {
	defer qt.Recovering("callback QHelpIndexWidget::editorDestroyed")

	if signal := qt.GetSignal(C.GoString(ptrName), "editorDestroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(editor))
		return true
	}
	return false

}

func (ptr *QHelpIndexWidget) EditorDestroyed(editor core.QObject_ITF) {
	defer qt.Recovering("QHelpIndexWidget::editorDestroyed")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_EditorDestroyed(ptr.Pointer(), core.PointerFromQObject(editor))
	}
}

func (ptr *QHelpIndexWidget) EditorDestroyedDefault(editor core.QObject_ITF) {
	defer qt.Recovering("QHelpIndexWidget::editorDestroyed")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_EditorDestroyedDefault(ptr.Pointer(), core.PointerFromQObject(editor))
	}
}

func (ptr *QHelpIndexWidget) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQHelpIndexWidgetFocusInEvent
func callbackQHelpIndexWidgetFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QHelpIndexWidget) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQHelpIndexWidgetFocusOutEvent
func callbackQHelpIndexWidgetFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QHelpIndexWidget) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQHelpIndexWidgetInputMethodEvent
func callbackQHelpIndexWidgetInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QHelpIndexWidget) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQHelpIndexWidgetKeyPressEvent
func callbackQHelpIndexWidgetKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QHelpIndexWidget) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectKeyboardSearch(f func(search string)) {
	defer qt.Recovering("connect QHelpIndexWidget::keyboardSearch")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyboardSearch", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectKeyboardSearch() {
	defer qt.Recovering("disconnect QHelpIndexWidget::keyboardSearch")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyboardSearch")
	}
}

//export callbackQHelpIndexWidgetKeyboardSearch
func callbackQHelpIndexWidgetKeyboardSearch(ptr unsafe.Pointer, ptrName *C.char, search *C.char) {
	defer qt.Recovering("callback QHelpIndexWidget::keyboardSearch")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyboardSearch"); signal != nil {
		signal.(func(string))(C.GoString(search))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).KeyboardSearchDefault(C.GoString(search))
	}
}

func (ptr *QHelpIndexWidget) KeyboardSearch(search string) {
	defer qt.Recovering("QHelpIndexWidget::keyboardSearch")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_KeyboardSearch(ptr.Pointer(), C.CString(search))
	}
}

func (ptr *QHelpIndexWidget) KeyboardSearchDefault(search string) {
	defer qt.Recovering("QHelpIndexWidget::keyboardSearch")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_KeyboardSearchDefault(ptr.Pointer(), C.CString(search))
	}
}

func (ptr *QHelpIndexWidget) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQHelpIndexWidgetMouseDoubleClickEvent
func callbackQHelpIndexWidgetMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpIndexWidget) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQHelpIndexWidgetMousePressEvent
func callbackQHelpIndexWidgetMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpIndexWidget) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectReset(f func()) {
	defer qt.Recovering("connect QHelpIndexWidget::reset")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "reset", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectReset() {
	defer qt.Recovering("disconnect QHelpIndexWidget::reset")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "reset")
	}
}

//export callbackQHelpIndexWidgetReset
func callbackQHelpIndexWidgetReset(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QHelpIndexWidget::reset")

	if signal := qt.GetSignal(C.GoString(ptrName), "reset"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QHelpIndexWidget) Reset() {
	defer qt.Recovering("QHelpIndexWidget::reset")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_Reset(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) ResetDefault() {
	defer qt.Recovering("QHelpIndexWidget::reset")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ResetDefault(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) ConnectSelectAll(f func()) {
	defer qt.Recovering("connect QHelpIndexWidget::selectAll")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "selectAll", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSelectAll() {
	defer qt.Recovering("disconnect QHelpIndexWidget::selectAll")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "selectAll")
	}
}

//export callbackQHelpIndexWidgetSelectAll
func callbackQHelpIndexWidgetSelectAll(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QHelpIndexWidget::selectAll")

	if signal := qt.GetSignal(C.GoString(ptrName), "selectAll"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QHelpIndexWidget) SelectAll() {
	defer qt.Recovering("QHelpIndexWidget::selectAll")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SelectAll(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) SelectAllDefault() {
	defer qt.Recovering("QHelpIndexWidget::selectAll")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SelectAllDefault(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) ConnectSetModel(f func(model *core.QAbstractItemModel)) {
	defer qt.Recovering("connect QHelpIndexWidget::setModel")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setModel", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSetModel() {
	defer qt.Recovering("disconnect QHelpIndexWidget::setModel")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setModel")
	}
}

//export callbackQHelpIndexWidgetSetModel
func callbackQHelpIndexWidgetSetModel(ptr unsafe.Pointer, ptrName *C.char, model unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::setModel")

	if signal := qt.GetSignal(C.GoString(ptrName), "setModel"); signal != nil {
		signal.(func(*core.QAbstractItemModel))(core.NewQAbstractItemModelFromPointer(model))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetModelDefault(core.NewQAbstractItemModelFromPointer(model))
	}
}

func (ptr *QHelpIndexWidget) SetModel(model core.QAbstractItemModel_ITF) {
	defer qt.Recovering("QHelpIndexWidget::setModel")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetModel(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QHelpIndexWidget) SetModelDefault(model core.QAbstractItemModel_ITF) {
	defer qt.Recovering("QHelpIndexWidget::setModel")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetModelDefault(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QHelpIndexWidget) ConnectSetRootIndex(f func(index *core.QModelIndex)) {
	defer qt.Recovering("connect QHelpIndexWidget::setRootIndex")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setRootIndex", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSetRootIndex() {
	defer qt.Recovering("disconnect QHelpIndexWidget::setRootIndex")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setRootIndex")
	}
}

//export callbackQHelpIndexWidgetSetRootIndex
func callbackQHelpIndexWidgetSetRootIndex(ptr unsafe.Pointer, ptrName *C.char, index unsafe.Pointer) bool {
	defer qt.Recovering("callback QHelpIndexWidget::setRootIndex")

	if signal := qt.GetSignal(C.GoString(ptrName), "setRootIndex"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
		return true
	}
	return false

}

func (ptr *QHelpIndexWidget) SetRootIndex(index core.QModelIndex_ITF) {
	defer qt.Recovering("QHelpIndexWidget::setRootIndex")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetRootIndex(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QHelpIndexWidget) SetRootIndexDefault(index core.QModelIndex_ITF) {
	defer qt.Recovering("QHelpIndexWidget::setRootIndex")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetRootIndexDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QHelpIndexWidget) ConnectSetSelectionModel(f func(selectionModel *core.QItemSelectionModel)) {
	defer qt.Recovering("connect QHelpIndexWidget::setSelectionModel")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setSelectionModel", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSetSelectionModel() {
	defer qt.Recovering("disconnect QHelpIndexWidget::setSelectionModel")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setSelectionModel")
	}
}

//export callbackQHelpIndexWidgetSetSelectionModel
func callbackQHelpIndexWidgetSetSelectionModel(ptr unsafe.Pointer, ptrName *C.char, selectionModel unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::setSelectionModel")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSelectionModel"); signal != nil {
		signal.(func(*core.QItemSelectionModel))(core.NewQItemSelectionModelFromPointer(selectionModel))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetSelectionModelDefault(core.NewQItemSelectionModelFromPointer(selectionModel))
	}
}

func (ptr *QHelpIndexWidget) SetSelectionModel(selectionModel core.QItemSelectionModel_ITF) {
	defer qt.Recovering("QHelpIndexWidget::setSelectionModel")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetSelectionModel(ptr.Pointer(), core.PointerFromQItemSelectionModel(selectionModel))
	}
}

func (ptr *QHelpIndexWidget) SetSelectionModelDefault(selectionModel core.QItemSelectionModel_ITF) {
	defer qt.Recovering("QHelpIndexWidget::setSelectionModel")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetSelectionModelDefault(ptr.Pointer(), core.PointerFromQItemSelectionModel(selectionModel))
	}
}

func (ptr *QHelpIndexWidget) ConnectContextMenuEvent(f func(e *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQHelpIndexWidgetContextMenuEvent
func callbackQHelpIndexWidgetContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) ContextMenuEvent(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

func (ptr *QHelpIndexWidget) ContextMenuEventDefault(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

func (ptr *QHelpIndexWidget) ConnectScrollContentsBy(f func(dx int, dy int)) {
	defer qt.Recovering("connect QHelpIndexWidget::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "scrollContentsBy", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectScrollContentsBy() {
	defer qt.Recovering("disconnect QHelpIndexWidget::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "scrollContentsBy")
	}
}

//export callbackQHelpIndexWidgetScrollContentsBy
func callbackQHelpIndexWidgetScrollContentsBy(ptr unsafe.Pointer, ptrName *C.char, dx C.int, dy C.int) {
	defer qt.Recovering("callback QHelpIndexWidget::scrollContentsBy")

	if signal := qt.GetSignal(C.GoString(ptrName), "scrollContentsBy"); signal != nil {
		signal.(func(int, int))(int(dx), int(dy))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ScrollContentsByDefault(int(dx), int(dy))
	}
}

func (ptr *QHelpIndexWidget) ScrollContentsBy(dx int, dy int) {
	defer qt.Recovering("QHelpIndexWidget::scrollContentsBy")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ScrollContentsBy(ptr.Pointer(), C.int(dx), C.int(dy))
	}
}

func (ptr *QHelpIndexWidget) ScrollContentsByDefault(dx int, dy int) {
	defer qt.Recovering("QHelpIndexWidget::scrollContentsBy")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ScrollContentsByDefault(ptr.Pointer(), C.int(dx), C.int(dy))
	}
}

func (ptr *QHelpIndexWidget) ConnectSetupViewport(f func(viewport *widgets.QWidget)) {
	defer qt.Recovering("connect QHelpIndexWidget::setupViewport")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setupViewport", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSetupViewport() {
	defer qt.Recovering("disconnect QHelpIndexWidget::setupViewport")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setupViewport")
	}
}

//export callbackQHelpIndexWidgetSetupViewport
func callbackQHelpIndexWidgetSetupViewport(ptr unsafe.Pointer, ptrName *C.char, viewport unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::setupViewport")

	if signal := qt.GetSignal(C.GoString(ptrName), "setupViewport"); signal != nil {
		signal.(func(*widgets.QWidget))(widgets.NewQWidgetFromPointer(viewport))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetupViewportDefault(widgets.NewQWidgetFromPointer(viewport))
	}
}

func (ptr *QHelpIndexWidget) SetupViewport(viewport widgets.QWidget_ITF) {
	defer qt.Recovering("QHelpIndexWidget::setupViewport")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetupViewport(ptr.Pointer(), widgets.PointerFromQWidget(viewport))
	}
}

func (ptr *QHelpIndexWidget) SetupViewportDefault(viewport widgets.QWidget_ITF) {
	defer qt.Recovering("QHelpIndexWidget::setupViewport")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetupViewportDefault(ptr.Pointer(), widgets.PointerFromQWidget(viewport))
	}
}

func (ptr *QHelpIndexWidget) ConnectWheelEvent(f func(e *gui.QWheelEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQHelpIndexWidgetWheelEvent
func callbackQHelpIndexWidgetWheelEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) WheelEvent(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QHelpIndexWidget) WheelEventDefault(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QHelpIndexWidget) ConnectChangeEvent(f func(ev *core.QEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQHelpIndexWidgetChangeEvent
func callbackQHelpIndexWidgetChangeEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(ev))
	}
}

func (ptr *QHelpIndexWidget) ChangeEvent(ev core.QEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::changeEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QHelpIndexWidget) ChangeEventDefault(ev core.QEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::changeEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QHelpIndexWidget) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQHelpIndexWidgetActionEvent
func callbackQHelpIndexWidgetActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::actionEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QHelpIndexWidget) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::actionEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQHelpIndexWidgetEnterEvent
func callbackQHelpIndexWidgetEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::enterEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpIndexWidget) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::enterEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQHelpIndexWidgetHideEvent
func callbackQHelpIndexWidgetHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QHelpIndexWidget) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQHelpIndexWidgetLeaveEvent
func callbackQHelpIndexWidgetLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::leaveEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpIndexWidget) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::leaveEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQHelpIndexWidgetMoveEvent
func callbackQHelpIndexWidgetMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QHelpIndexWidget) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QHelpIndexWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QHelpIndexWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQHelpIndexWidgetSetVisible
func callbackQHelpIndexWidgetSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QHelpIndexWidget::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QHelpIndexWidget) SetVisible(visible bool) {
	defer qt.Recovering("QHelpIndexWidget::setVisible")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QHelpIndexWidget) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QHelpIndexWidget::setVisible")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QHelpIndexWidget) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQHelpIndexWidgetShowEvent
func callbackQHelpIndexWidgetShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QHelpIndexWidget) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQHelpIndexWidgetCloseEvent
func callbackQHelpIndexWidgetCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QHelpIndexWidget) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QHelpIndexWidget::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QHelpIndexWidget::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQHelpIndexWidgetInitPainter
func callbackQHelpIndexWidgetInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QHelpIndexWidget) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QHelpIndexWidget::initPainter")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QHelpIndexWidget) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QHelpIndexWidget::initPainter")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QHelpIndexWidget) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQHelpIndexWidgetKeyReleaseEvent
func callbackQHelpIndexWidgetKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QHelpIndexWidget) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQHelpIndexWidgetTabletEvent
func callbackQHelpIndexWidgetTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::tabletEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QHelpIndexWidget) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::tabletEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQHelpIndexWidgetChildEvent
func callbackQHelpIndexWidgetChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHelpIndexWidget) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQHelpIndexWidgetCustomEvent
func callbackQHelpIndexWidgetCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpIndexWidget) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
