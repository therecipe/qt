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

type QHelpContentWidget struct {
	widgets.QTreeView
}

type QHelpContentWidget_ITF interface {
	widgets.QTreeView_ITF
	QHelpContentWidget_PTR() *QHelpContentWidget
}

func PointerFromQHelpContentWidget(ptr QHelpContentWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpContentWidget_PTR().Pointer()
	}
	return nil
}

func NewQHelpContentWidgetFromPointer(ptr unsafe.Pointer) *QHelpContentWidget {
	var n = new(QHelpContentWidget)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QHelpContentWidget_") {
		n.SetObjectName("QHelpContentWidget_" + qt.Identifier())
	}
	return n
}

func (ptr *QHelpContentWidget) QHelpContentWidget_PTR() *QHelpContentWidget {
	return ptr
}

func (ptr *QHelpContentWidget) IndexOf(link core.QUrl_ITF) *core.QModelIndex {
	defer qt.Recovering("QHelpContentWidget::indexOf")

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QHelpContentWidget_IndexOf(ptr.Pointer(), core.PointerFromQUrl(link)))
	}
	return nil
}

func (ptr *QHelpContentWidget) ConnectLinkActivated(f func(link *core.QUrl)) {
	defer qt.Recovering("connect QHelpContentWidget::linkActivated")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ConnectLinkActivated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "linkActivated", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectLinkActivated() {
	defer qt.Recovering("disconnect QHelpContentWidget::linkActivated")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DisconnectLinkActivated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "linkActivated")
	}
}

//export callbackQHelpContentWidgetLinkActivated
func callbackQHelpContentWidgetLinkActivated(ptr unsafe.Pointer, ptrName *C.char, link unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::linkActivated")

	if signal := qt.GetSignal(C.GoString(ptrName), "linkActivated"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(link))
	}

}

func (ptr *QHelpContentWidget) LinkActivated(link core.QUrl_ITF) {
	defer qt.Recovering("QHelpContentWidget::linkActivated")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_LinkActivated(ptr.Pointer(), core.PointerFromQUrl(link))
	}
}

func (ptr *QHelpContentWidget) ConnectCurrentChanged(f func(current *core.QModelIndex, previous *core.QModelIndex)) {
	defer qt.Recovering("connect QHelpContentWidget::currentChanged")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "currentChanged", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectCurrentChanged() {
	defer qt.Recovering("disconnect QHelpContentWidget::currentChanged")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "currentChanged")
	}
}

//export callbackQHelpContentWidgetCurrentChanged
func callbackQHelpContentWidgetCurrentChanged(ptr unsafe.Pointer, ptrName *C.char, current unsafe.Pointer, previous unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::currentChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "currentChanged"); signal != nil {
		signal.(func(*core.QModelIndex, *core.QModelIndex))(core.NewQModelIndexFromPointer(current), core.NewQModelIndexFromPointer(previous))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).CurrentChangedDefault(core.NewQModelIndexFromPointer(current), core.NewQModelIndexFromPointer(previous))
	}
}

func (ptr *QHelpContentWidget) CurrentChanged(current core.QModelIndex_ITF, previous core.QModelIndex_ITF) {
	defer qt.Recovering("QHelpContentWidget::currentChanged")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CurrentChanged(ptr.Pointer(), core.PointerFromQModelIndex(current), core.PointerFromQModelIndex(previous))
	}
}

func (ptr *QHelpContentWidget) CurrentChangedDefault(current core.QModelIndex_ITF, previous core.QModelIndex_ITF) {
	defer qt.Recovering("QHelpContentWidget::currentChanged")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CurrentChangedDefault(ptr.Pointer(), core.PointerFromQModelIndex(current), core.PointerFromQModelIndex(previous))
	}
}

func (ptr *QHelpContentWidget) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQHelpContentWidgetDragMoveEvent
func callbackQHelpContentWidgetDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QHelpContentWidget) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectDrawBranches(f func(painter *gui.QPainter, rect *core.QRect, index *core.QModelIndex)) {
	defer qt.Recovering("connect QHelpContentWidget::drawBranches")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "drawBranches", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectDrawBranches() {
	defer qt.Recovering("disconnect QHelpContentWidget::drawBranches")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "drawBranches")
	}
}

//export callbackQHelpContentWidgetDrawBranches
func callbackQHelpContentWidgetDrawBranches(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer, rect unsafe.Pointer, index unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::drawBranches")

	if signal := qt.GetSignal(C.GoString(ptrName), "drawBranches"); signal != nil {
		signal.(func(*gui.QPainter, *core.QRect, *core.QModelIndex))(gui.NewQPainterFromPointer(painter), core.NewQRectFromPointer(rect), core.NewQModelIndexFromPointer(index))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).DrawBranchesDefault(gui.NewQPainterFromPointer(painter), core.NewQRectFromPointer(rect), core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QHelpContentWidget) DrawBranches(painter gui.QPainter_ITF, rect core.QRect_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("QHelpContentWidget::drawBranches")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DrawBranches(ptr.Pointer(), gui.PointerFromQPainter(painter), core.PointerFromQRect(rect), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QHelpContentWidget) DrawBranchesDefault(painter gui.QPainter_ITF, rect core.QRect_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("QHelpContentWidget::drawBranches")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DrawBranchesDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), core.PointerFromQRect(rect), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QHelpContentWidget) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQHelpContentWidgetKeyPressEvent
func callbackQHelpContentWidgetKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QHelpContentWidget) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectKeyboardSearch(f func(search string)) {
	defer qt.Recovering("connect QHelpContentWidget::keyboardSearch")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyboardSearch", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectKeyboardSearch() {
	defer qt.Recovering("disconnect QHelpContentWidget::keyboardSearch")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyboardSearch")
	}
}

//export callbackQHelpContentWidgetKeyboardSearch
func callbackQHelpContentWidgetKeyboardSearch(ptr unsafe.Pointer, ptrName *C.char, search *C.char) {
	defer qt.Recovering("callback QHelpContentWidget::keyboardSearch")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyboardSearch"); signal != nil {
		signal.(func(string))(C.GoString(search))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).KeyboardSearchDefault(C.GoString(search))
	}
}

func (ptr *QHelpContentWidget) KeyboardSearch(search string) {
	defer qt.Recovering("QHelpContentWidget::keyboardSearch")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_KeyboardSearch(ptr.Pointer(), C.CString(search))
	}
}

func (ptr *QHelpContentWidget) KeyboardSearchDefault(search string) {
	defer qt.Recovering("QHelpContentWidget::keyboardSearch")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_KeyboardSearchDefault(ptr.Pointer(), C.CString(search))
	}
}

func (ptr *QHelpContentWidget) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQHelpContentWidgetMouseDoubleClickEvent
func callbackQHelpContentWidgetMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpContentWidget) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQHelpContentWidgetMouseMoveEvent
func callbackQHelpContentWidgetMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpContentWidget) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQHelpContentWidgetMousePressEvent
func callbackQHelpContentWidgetMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpContentWidget) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQHelpContentWidgetMouseReleaseEvent
func callbackQHelpContentWidgetMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpContentWidget) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQHelpContentWidgetPaintEvent
func callbackQHelpContentWidgetPaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::paintEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QHelpContentWidget) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::paintEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectReset(f func()) {
	defer qt.Recovering("connect QHelpContentWidget::reset")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "reset", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectReset() {
	defer qt.Recovering("disconnect QHelpContentWidget::reset")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "reset")
	}
}

//export callbackQHelpContentWidgetReset
func callbackQHelpContentWidgetReset(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QHelpContentWidget::reset")

	if signal := qt.GetSignal(C.GoString(ptrName), "reset"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ResetDefault()
	}
}

func (ptr *QHelpContentWidget) Reset() {
	defer qt.Recovering("QHelpContentWidget::reset")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_Reset(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) ResetDefault() {
	defer qt.Recovering("QHelpContentWidget::reset")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ResetDefault(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) ConnectRowsAboutToBeRemoved(f func(parent *core.QModelIndex, start int, end int)) {
	defer qt.Recovering("connect QHelpContentWidget::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "rowsAboutToBeRemoved", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectRowsAboutToBeRemoved() {
	defer qt.Recovering("disconnect QHelpContentWidget::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "rowsAboutToBeRemoved")
	}
}

//export callbackQHelpContentWidgetRowsAboutToBeRemoved
func callbackQHelpContentWidgetRowsAboutToBeRemoved(ptr unsafe.Pointer, ptrName *C.char, parent unsafe.Pointer, start C.int, end C.int) {
	defer qt.Recovering("callback QHelpContentWidget::rowsAboutToBeRemoved")

	if signal := qt.GetSignal(C.GoString(ptrName), "rowsAboutToBeRemoved"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(start), int(end))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).RowsAboutToBeRemovedDefault(core.NewQModelIndexFromPointer(parent), int(start), int(end))
	}
}

func (ptr *QHelpContentWidget) RowsAboutToBeRemoved(parent core.QModelIndex_ITF, start int, end int) {
	defer qt.Recovering("QHelpContentWidget::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_RowsAboutToBeRemoved(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(start), C.int(end))
	}
}

func (ptr *QHelpContentWidget) RowsAboutToBeRemovedDefault(parent core.QModelIndex_ITF, start int, end int) {
	defer qt.Recovering("QHelpContentWidget::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_RowsAboutToBeRemovedDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(start), C.int(end))
	}
}

func (ptr *QHelpContentWidget) ConnectRowsInserted(f func(parent *core.QModelIndex, start int, end int)) {
	defer qt.Recovering("connect QHelpContentWidget::rowsInserted")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "rowsInserted", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectRowsInserted() {
	defer qt.Recovering("disconnect QHelpContentWidget::rowsInserted")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "rowsInserted")
	}
}

//export callbackQHelpContentWidgetRowsInserted
func callbackQHelpContentWidgetRowsInserted(ptr unsafe.Pointer, ptrName *C.char, parent unsafe.Pointer, start C.int, end C.int) {
	defer qt.Recovering("callback QHelpContentWidget::rowsInserted")

	if signal := qt.GetSignal(C.GoString(ptrName), "rowsInserted"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(start), int(end))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).RowsInsertedDefault(core.NewQModelIndexFromPointer(parent), int(start), int(end))
	}
}

func (ptr *QHelpContentWidget) RowsInserted(parent core.QModelIndex_ITF, start int, end int) {
	defer qt.Recovering("QHelpContentWidget::rowsInserted")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_RowsInserted(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(start), C.int(end))
	}
}

func (ptr *QHelpContentWidget) RowsInsertedDefault(parent core.QModelIndex_ITF, start int, end int) {
	defer qt.Recovering("QHelpContentWidget::rowsInserted")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_RowsInsertedDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(start), C.int(end))
	}
}

func (ptr *QHelpContentWidget) ConnectScrollContentsBy(f func(dx int, dy int)) {
	defer qt.Recovering("connect QHelpContentWidget::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "scrollContentsBy", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectScrollContentsBy() {
	defer qt.Recovering("disconnect QHelpContentWidget::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "scrollContentsBy")
	}
}

//export callbackQHelpContentWidgetScrollContentsBy
func callbackQHelpContentWidgetScrollContentsBy(ptr unsafe.Pointer, ptrName *C.char, dx C.int, dy C.int) {
	defer qt.Recovering("callback QHelpContentWidget::scrollContentsBy")

	if signal := qt.GetSignal(C.GoString(ptrName), "scrollContentsBy"); signal != nil {
		signal.(func(int, int))(int(dx), int(dy))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ScrollContentsByDefault(int(dx), int(dy))
	}
}

func (ptr *QHelpContentWidget) ScrollContentsBy(dx int, dy int) {
	defer qt.Recovering("QHelpContentWidget::scrollContentsBy")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ScrollContentsBy(ptr.Pointer(), C.int(dx), C.int(dy))
	}
}

func (ptr *QHelpContentWidget) ScrollContentsByDefault(dx int, dy int) {
	defer qt.Recovering("QHelpContentWidget::scrollContentsBy")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ScrollContentsByDefault(ptr.Pointer(), C.int(dx), C.int(dy))
	}
}

func (ptr *QHelpContentWidget) ConnectScrollTo(f func(index *core.QModelIndex, hint widgets.QAbstractItemView__ScrollHint)) {
	defer qt.Recovering("connect QHelpContentWidget::scrollTo")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "scrollTo", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectScrollTo() {
	defer qt.Recovering("disconnect QHelpContentWidget::scrollTo")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "scrollTo")
	}
}

//export callbackQHelpContentWidgetScrollTo
func callbackQHelpContentWidgetScrollTo(ptr unsafe.Pointer, ptrName *C.char, index unsafe.Pointer, hint C.int) {
	defer qt.Recovering("callback QHelpContentWidget::scrollTo")

	if signal := qt.GetSignal(C.GoString(ptrName), "scrollTo"); signal != nil {
		signal.(func(*core.QModelIndex, widgets.QAbstractItemView__ScrollHint))(core.NewQModelIndexFromPointer(index), widgets.QAbstractItemView__ScrollHint(hint))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ScrollToDefault(core.NewQModelIndexFromPointer(index), widgets.QAbstractItemView__ScrollHint(hint))
	}
}

func (ptr *QHelpContentWidget) ScrollTo(index core.QModelIndex_ITF, hint widgets.QAbstractItemView__ScrollHint) {
	defer qt.Recovering("QHelpContentWidget::scrollTo")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ScrollTo(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(hint))
	}
}

func (ptr *QHelpContentWidget) ScrollToDefault(index core.QModelIndex_ITF, hint widgets.QAbstractItemView__ScrollHint) {
	defer qt.Recovering("QHelpContentWidget::scrollTo")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ScrollToDefault(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(hint))
	}
}

func (ptr *QHelpContentWidget) ConnectSelectAll(f func()) {
	defer qt.Recovering("connect QHelpContentWidget::selectAll")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "selectAll", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSelectAll() {
	defer qt.Recovering("disconnect QHelpContentWidget::selectAll")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "selectAll")
	}
}

//export callbackQHelpContentWidgetSelectAll
func callbackQHelpContentWidgetSelectAll(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QHelpContentWidget::selectAll")

	if signal := qt.GetSignal(C.GoString(ptrName), "selectAll"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SelectAllDefault()
	}
}

func (ptr *QHelpContentWidget) SelectAll() {
	defer qt.Recovering("QHelpContentWidget::selectAll")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SelectAll(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) SelectAllDefault() {
	defer qt.Recovering("QHelpContentWidget::selectAll")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SelectAllDefault(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) ConnectSetModel(f func(model *core.QAbstractItemModel)) {
	defer qt.Recovering("connect QHelpContentWidget::setModel")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setModel", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSetModel() {
	defer qt.Recovering("disconnect QHelpContentWidget::setModel")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setModel")
	}
}

//export callbackQHelpContentWidgetSetModel
func callbackQHelpContentWidgetSetModel(ptr unsafe.Pointer, ptrName *C.char, model unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::setModel")

	if signal := qt.GetSignal(C.GoString(ptrName), "setModel"); signal != nil {
		signal.(func(*core.QAbstractItemModel))(core.NewQAbstractItemModelFromPointer(model))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetModelDefault(core.NewQAbstractItemModelFromPointer(model))
	}
}

func (ptr *QHelpContentWidget) SetModel(model core.QAbstractItemModel_ITF) {
	defer qt.Recovering("QHelpContentWidget::setModel")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetModel(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QHelpContentWidget) SetModelDefault(model core.QAbstractItemModel_ITF) {
	defer qt.Recovering("QHelpContentWidget::setModel")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetModelDefault(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QHelpContentWidget) ConnectSetRootIndex(f func(index *core.QModelIndex)) {
	defer qt.Recovering("connect QHelpContentWidget::setRootIndex")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setRootIndex", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSetRootIndex() {
	defer qt.Recovering("disconnect QHelpContentWidget::setRootIndex")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setRootIndex")
	}
}

//export callbackQHelpContentWidgetSetRootIndex
func callbackQHelpContentWidgetSetRootIndex(ptr unsafe.Pointer, ptrName *C.char, index unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::setRootIndex")

	if signal := qt.GetSignal(C.GoString(ptrName), "setRootIndex"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetRootIndexDefault(core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QHelpContentWidget) SetRootIndex(index core.QModelIndex_ITF) {
	defer qt.Recovering("QHelpContentWidget::setRootIndex")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetRootIndex(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QHelpContentWidget) SetRootIndexDefault(index core.QModelIndex_ITF) {
	defer qt.Recovering("QHelpContentWidget::setRootIndex")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetRootIndexDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QHelpContentWidget) ConnectSetSelection(f func(rect *core.QRect, command core.QItemSelectionModel__SelectionFlag)) {
	defer qt.Recovering("connect QHelpContentWidget::setSelection")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setSelection", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSetSelection() {
	defer qt.Recovering("disconnect QHelpContentWidget::setSelection")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setSelection")
	}
}

//export callbackQHelpContentWidgetSetSelection
func callbackQHelpContentWidgetSetSelection(ptr unsafe.Pointer, ptrName *C.char, rect unsafe.Pointer, command C.int) {
	defer qt.Recovering("callback QHelpContentWidget::setSelection")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSelection"); signal != nil {
		signal.(func(*core.QRect, core.QItemSelectionModel__SelectionFlag))(core.NewQRectFromPointer(rect), core.QItemSelectionModel__SelectionFlag(command))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetSelectionDefault(core.NewQRectFromPointer(rect), core.QItemSelectionModel__SelectionFlag(command))
	}
}

func (ptr *QHelpContentWidget) SetSelection(rect core.QRect_ITF, command core.QItemSelectionModel__SelectionFlag) {
	defer qt.Recovering("QHelpContentWidget::setSelection")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetSelection(ptr.Pointer(), core.PointerFromQRect(rect), C.int(command))
	}
}

func (ptr *QHelpContentWidget) SetSelectionDefault(rect core.QRect_ITF, command core.QItemSelectionModel__SelectionFlag) {
	defer qt.Recovering("QHelpContentWidget::setSelection")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetSelectionDefault(ptr.Pointer(), core.PointerFromQRect(rect), C.int(command))
	}
}

func (ptr *QHelpContentWidget) ConnectSetSelectionModel(f func(selectionModel *core.QItemSelectionModel)) {
	defer qt.Recovering("connect QHelpContentWidget::setSelectionModel")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setSelectionModel", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSetSelectionModel() {
	defer qt.Recovering("disconnect QHelpContentWidget::setSelectionModel")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setSelectionModel")
	}
}

//export callbackQHelpContentWidgetSetSelectionModel
func callbackQHelpContentWidgetSetSelectionModel(ptr unsafe.Pointer, ptrName *C.char, selectionModel unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::setSelectionModel")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSelectionModel"); signal != nil {
		signal.(func(*core.QItemSelectionModel))(core.NewQItemSelectionModelFromPointer(selectionModel))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetSelectionModelDefault(core.NewQItemSelectionModelFromPointer(selectionModel))
	}
}

func (ptr *QHelpContentWidget) SetSelectionModel(selectionModel core.QItemSelectionModel_ITF) {
	defer qt.Recovering("QHelpContentWidget::setSelectionModel")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetSelectionModel(ptr.Pointer(), core.PointerFromQItemSelectionModel(selectionModel))
	}
}

func (ptr *QHelpContentWidget) SetSelectionModelDefault(selectionModel core.QItemSelectionModel_ITF) {
	defer qt.Recovering("QHelpContentWidget::setSelectionModel")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetSelectionModelDefault(ptr.Pointer(), core.PointerFromQItemSelectionModel(selectionModel))
	}
}

func (ptr *QHelpContentWidget) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQHelpContentWidgetTimerEvent
func callbackQHelpContentWidgetTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::timerEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHelpContentWidget) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::timerEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectUpdateGeometries(f func()) {
	defer qt.Recovering("connect QHelpContentWidget::updateGeometries")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updateGeometries", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectUpdateGeometries() {
	defer qt.Recovering("disconnect QHelpContentWidget::updateGeometries")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updateGeometries")
	}
}

//export callbackQHelpContentWidgetUpdateGeometries
func callbackQHelpContentWidgetUpdateGeometries(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QHelpContentWidget::updateGeometries")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateGeometries"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).UpdateGeometriesDefault()
	}
}

func (ptr *QHelpContentWidget) UpdateGeometries() {
	defer qt.Recovering("QHelpContentWidget::updateGeometries")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_UpdateGeometries(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) UpdateGeometriesDefault() {
	defer qt.Recovering("QHelpContentWidget::updateGeometries")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_UpdateGeometriesDefault(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQHelpContentWidgetDragLeaveEvent
func callbackQHelpContentWidgetDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QHelpContentWidget) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectCloseEditor(f func(editor *widgets.QWidget, hint widgets.QAbstractItemDelegate__EndEditHint)) {
	defer qt.Recovering("connect QHelpContentWidget::closeEditor")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEditor", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectCloseEditor() {
	defer qt.Recovering("disconnect QHelpContentWidget::closeEditor")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEditor")
	}
}

//export callbackQHelpContentWidgetCloseEditor
func callbackQHelpContentWidgetCloseEditor(ptr unsafe.Pointer, ptrName *C.char, editor unsafe.Pointer, hint C.int) bool {
	defer qt.Recovering("callback QHelpContentWidget::closeEditor")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEditor"); signal != nil {
		signal.(func(*widgets.QWidget, widgets.QAbstractItemDelegate__EndEditHint))(widgets.NewQWidgetFromPointer(editor), widgets.QAbstractItemDelegate__EndEditHint(hint))
		return true
	}
	return false

}

func (ptr *QHelpContentWidget) CloseEditor(editor widgets.QWidget_ITF, hint widgets.QAbstractItemDelegate__EndEditHint) {
	defer qt.Recovering("QHelpContentWidget::closeEditor")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CloseEditor(ptr.Pointer(), widgets.PointerFromQWidget(editor), C.int(hint))
	}
}

func (ptr *QHelpContentWidget) CloseEditorDefault(editor widgets.QWidget_ITF, hint widgets.QAbstractItemDelegate__EndEditHint) {
	defer qt.Recovering("QHelpContentWidget::closeEditor")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CloseEditorDefault(ptr.Pointer(), widgets.PointerFromQWidget(editor), C.int(hint))
	}
}

func (ptr *QHelpContentWidget) ConnectCommitData(f func(editor *widgets.QWidget)) {
	defer qt.Recovering("connect QHelpContentWidget::commitData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "commitData", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectCommitData() {
	defer qt.Recovering("disconnect QHelpContentWidget::commitData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "commitData")
	}
}

//export callbackQHelpContentWidgetCommitData
func callbackQHelpContentWidgetCommitData(ptr unsafe.Pointer, ptrName *C.char, editor unsafe.Pointer) bool {
	defer qt.Recovering("callback QHelpContentWidget::commitData")

	if signal := qt.GetSignal(C.GoString(ptrName), "commitData"); signal != nil {
		signal.(func(*widgets.QWidget))(widgets.NewQWidgetFromPointer(editor))
		return true
	}
	return false

}

func (ptr *QHelpContentWidget) CommitData(editor widgets.QWidget_ITF) {
	defer qt.Recovering("QHelpContentWidget::commitData")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CommitData(ptr.Pointer(), widgets.PointerFromQWidget(editor))
	}
}

func (ptr *QHelpContentWidget) CommitDataDefault(editor widgets.QWidget_ITF) {
	defer qt.Recovering("QHelpContentWidget::commitData")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CommitDataDefault(ptr.Pointer(), widgets.PointerFromQWidget(editor))
	}
}

func (ptr *QHelpContentWidget) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQHelpContentWidgetDragEnterEvent
func callbackQHelpContentWidgetDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QHelpContentWidget) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQHelpContentWidgetDropEvent
func callbackQHelpContentWidgetDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QHelpContentWidget) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectEditorDestroyed(f func(editor *core.QObject)) {
	defer qt.Recovering("connect QHelpContentWidget::editorDestroyed")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "editorDestroyed", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectEditorDestroyed() {
	defer qt.Recovering("disconnect QHelpContentWidget::editorDestroyed")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "editorDestroyed")
	}
}

//export callbackQHelpContentWidgetEditorDestroyed
func callbackQHelpContentWidgetEditorDestroyed(ptr unsafe.Pointer, ptrName *C.char, editor unsafe.Pointer) bool {
	defer qt.Recovering("callback QHelpContentWidget::editorDestroyed")

	if signal := qt.GetSignal(C.GoString(ptrName), "editorDestroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(editor))
		return true
	}
	return false

}

func (ptr *QHelpContentWidget) EditorDestroyed(editor core.QObject_ITF) {
	defer qt.Recovering("QHelpContentWidget::editorDestroyed")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_EditorDestroyed(ptr.Pointer(), core.PointerFromQObject(editor))
	}
}

func (ptr *QHelpContentWidget) EditorDestroyedDefault(editor core.QObject_ITF) {
	defer qt.Recovering("QHelpContentWidget::editorDestroyed")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_EditorDestroyedDefault(ptr.Pointer(), core.PointerFromQObject(editor))
	}
}

func (ptr *QHelpContentWidget) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQHelpContentWidgetFocusInEvent
func callbackQHelpContentWidgetFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QHelpContentWidget) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQHelpContentWidgetFocusOutEvent
func callbackQHelpContentWidgetFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QHelpContentWidget) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQHelpContentWidgetInputMethodEvent
func callbackQHelpContentWidgetInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QHelpContentWidget) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQHelpContentWidgetResizeEvent
func callbackQHelpContentWidgetResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QHelpContentWidget) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectStartDrag(f func(supportedActions core.Qt__DropAction)) {
	defer qt.Recovering("connect QHelpContentWidget::startDrag")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "startDrag", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectStartDrag() {
	defer qt.Recovering("disconnect QHelpContentWidget::startDrag")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "startDrag")
	}
}

//export callbackQHelpContentWidgetStartDrag
func callbackQHelpContentWidgetStartDrag(ptr unsafe.Pointer, ptrName *C.char, supportedActions C.int) {
	defer qt.Recovering("callback QHelpContentWidget::startDrag")

	if signal := qt.GetSignal(C.GoString(ptrName), "startDrag"); signal != nil {
		signal.(func(core.Qt__DropAction))(core.Qt__DropAction(supportedActions))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).StartDragDefault(core.Qt__DropAction(supportedActions))
	}
}

func (ptr *QHelpContentWidget) StartDrag(supportedActions core.Qt__DropAction) {
	defer qt.Recovering("QHelpContentWidget::startDrag")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_StartDrag(ptr.Pointer(), C.int(supportedActions))
	}
}

func (ptr *QHelpContentWidget) StartDragDefault(supportedActions core.Qt__DropAction) {
	defer qt.Recovering("QHelpContentWidget::startDrag")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_StartDragDefault(ptr.Pointer(), C.int(supportedActions))
	}
}

func (ptr *QHelpContentWidget) ConnectContextMenuEvent(f func(e *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQHelpContentWidgetContextMenuEvent
func callbackQHelpContentWidgetContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(e))
	}
}

func (ptr *QHelpContentWidget) ContextMenuEvent(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

func (ptr *QHelpContentWidget) ContextMenuEventDefault(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

func (ptr *QHelpContentWidget) ConnectSetupViewport(f func(viewport *widgets.QWidget)) {
	defer qt.Recovering("connect QHelpContentWidget::setupViewport")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setupViewport", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSetupViewport() {
	defer qt.Recovering("disconnect QHelpContentWidget::setupViewport")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setupViewport")
	}
}

//export callbackQHelpContentWidgetSetupViewport
func callbackQHelpContentWidgetSetupViewport(ptr unsafe.Pointer, ptrName *C.char, viewport unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::setupViewport")

	if signal := qt.GetSignal(C.GoString(ptrName), "setupViewport"); signal != nil {
		signal.(func(*widgets.QWidget))(widgets.NewQWidgetFromPointer(viewport))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetupViewportDefault(widgets.NewQWidgetFromPointer(viewport))
	}
}

func (ptr *QHelpContentWidget) SetupViewport(viewport widgets.QWidget_ITF) {
	defer qt.Recovering("QHelpContentWidget::setupViewport")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetupViewport(ptr.Pointer(), widgets.PointerFromQWidget(viewport))
	}
}

func (ptr *QHelpContentWidget) SetupViewportDefault(viewport widgets.QWidget_ITF) {
	defer qt.Recovering("QHelpContentWidget::setupViewport")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetupViewportDefault(ptr.Pointer(), widgets.PointerFromQWidget(viewport))
	}
}

func (ptr *QHelpContentWidget) ConnectWheelEvent(f func(e *gui.QWheelEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQHelpContentWidgetWheelEvent
func callbackQHelpContentWidgetWheelEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(e))
	}
}

func (ptr *QHelpContentWidget) WheelEvent(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QHelpContentWidget) WheelEventDefault(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QHelpContentWidget) ConnectChangeEvent(f func(ev *core.QEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQHelpContentWidgetChangeEvent
func callbackQHelpContentWidgetChangeEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(ev))
	}
}

func (ptr *QHelpContentWidget) ChangeEvent(ev core.QEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::changeEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QHelpContentWidget) ChangeEventDefault(ev core.QEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::changeEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QHelpContentWidget) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQHelpContentWidgetActionEvent
func callbackQHelpContentWidgetActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::actionEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QHelpContentWidget) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::actionEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQHelpContentWidgetEnterEvent
func callbackQHelpContentWidgetEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::enterEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpContentWidget) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::enterEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQHelpContentWidgetHideEvent
func callbackQHelpContentWidgetHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QHelpContentWidget) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQHelpContentWidgetLeaveEvent
func callbackQHelpContentWidgetLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::leaveEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpContentWidget) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::leaveEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQHelpContentWidgetMoveEvent
func callbackQHelpContentWidgetMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QHelpContentWidget) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QHelpContentWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QHelpContentWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQHelpContentWidgetSetVisible
func callbackQHelpContentWidgetSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QHelpContentWidget::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QHelpContentWidget) SetVisible(visible bool) {
	defer qt.Recovering("QHelpContentWidget::setVisible")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QHelpContentWidget) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QHelpContentWidget::setVisible")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QHelpContentWidget) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQHelpContentWidgetShowEvent
func callbackQHelpContentWidgetShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QHelpContentWidget) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQHelpContentWidgetCloseEvent
func callbackQHelpContentWidgetCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QHelpContentWidget) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QHelpContentWidget::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QHelpContentWidget::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQHelpContentWidgetInitPainter
func callbackQHelpContentWidgetInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QHelpContentWidget) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QHelpContentWidget::initPainter")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QHelpContentWidget) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QHelpContentWidget::initPainter")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QHelpContentWidget) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQHelpContentWidgetKeyReleaseEvent
func callbackQHelpContentWidgetKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QHelpContentWidget) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQHelpContentWidgetTabletEvent
func callbackQHelpContentWidgetTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::tabletEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QHelpContentWidget) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::tabletEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQHelpContentWidgetChildEvent
func callbackQHelpContentWidgetChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHelpContentWidget) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQHelpContentWidgetCustomEvent
func callbackQHelpContentWidgetCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpContentWidget) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
