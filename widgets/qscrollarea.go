package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QScrollArea struct {
	QAbstractScrollArea
}

type QScrollArea_ITF interface {
	QAbstractScrollArea_ITF
	QScrollArea_PTR() *QScrollArea
}

func PointerFromQScrollArea(ptr QScrollArea_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScrollArea_PTR().Pointer()
	}
	return nil
}

func NewQScrollAreaFromPointer(ptr unsafe.Pointer) *QScrollArea {
	var n = new(QScrollArea)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QScrollArea_") {
		n.SetObjectName("QScrollArea_" + qt.Identifier())
	}
	return n
}

func (ptr *QScrollArea) QScrollArea_PTR() *QScrollArea {
	return ptr
}

func (ptr *QScrollArea) Alignment() core.Qt__AlignmentFlag {
	defer qt.Recovering("QScrollArea::alignment")

	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QScrollArea_Alignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScrollArea) SetAlignment(v core.Qt__AlignmentFlag) {
	defer qt.Recovering("QScrollArea::setAlignment")

	if ptr.Pointer() != nil {
		C.QScrollArea_SetAlignment(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QScrollArea) SetWidget(widget QWidget_ITF) {
	defer qt.Recovering("QScrollArea::setWidget")

	if ptr.Pointer() != nil {
		C.QScrollArea_SetWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QScrollArea) SetWidgetResizable(resizable bool) {
	defer qt.Recovering("QScrollArea::setWidgetResizable")

	if ptr.Pointer() != nil {
		C.QScrollArea_SetWidgetResizable(ptr.Pointer(), C.int(qt.GoBoolToInt(resizable)))
	}
}

func (ptr *QScrollArea) WidgetResizable() bool {
	defer qt.Recovering("QScrollArea::widgetResizable")

	if ptr.Pointer() != nil {
		return C.QScrollArea_WidgetResizable(ptr.Pointer()) != 0
	}
	return false
}

func NewQScrollArea(parent QWidget_ITF) *QScrollArea {
	defer qt.Recovering("QScrollArea::QScrollArea")

	return NewQScrollAreaFromPointer(C.QScrollArea_NewQScrollArea(PointerFromQWidget(parent)))
}

func (ptr *QScrollArea) EnsureVisible(x int, y int, xmargin int, ymargin int) {
	defer qt.Recovering("QScrollArea::ensureVisible")

	if ptr.Pointer() != nil {
		C.QScrollArea_EnsureVisible(ptr.Pointer(), C.int(x), C.int(y), C.int(xmargin), C.int(ymargin))
	}
}

func (ptr *QScrollArea) EnsureWidgetVisible(childWidget QWidget_ITF, xmargin int, ymargin int) {
	defer qt.Recovering("QScrollArea::ensureWidgetVisible")

	if ptr.Pointer() != nil {
		C.QScrollArea_EnsureWidgetVisible(ptr.Pointer(), PointerFromQWidget(childWidget), C.int(xmargin), C.int(ymargin))
	}
}

func (ptr *QScrollArea) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QScrollArea::event")

	if ptr.Pointer() != nil {
		return C.QScrollArea_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QScrollArea) EventFilter(o core.QObject_ITF, e core.QEvent_ITF) bool {
	defer qt.Recovering("QScrollArea::eventFilter")

	if ptr.Pointer() != nil {
		return C.QScrollArea_EventFilter(ptr.Pointer(), core.PointerFromQObject(o), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QScrollArea) FocusNextPrevChild(next bool) bool {
	defer qt.Recovering("QScrollArea::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QScrollArea_FocusNextPrevChild(ptr.Pointer(), C.int(qt.GoBoolToInt(next))) != 0
	}
	return false
}

func (ptr *QScrollArea) ConnectResizeEvent(f func(v *gui.QResizeEvent)) {
	defer qt.Recovering("connect QScrollArea::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QScrollArea) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QScrollArea::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQScrollAreaResizeEvent
func callbackQScrollAreaResizeEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QScrollArea::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(v))
	} else {
		NewQScrollAreaFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(v))
	}
}

func (ptr *QScrollArea) ResizeEvent(v gui.QResizeEvent_ITF) {
	defer qt.Recovering("QScrollArea::resizeEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(v))
	}
}

func (ptr *QScrollArea) ResizeEventDefault(v gui.QResizeEvent_ITF) {
	defer qt.Recovering("QScrollArea::resizeEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(v))
	}
}

func (ptr *QScrollArea) ConnectScrollContentsBy(f func(dx int, dy int)) {
	defer qt.Recovering("connect QScrollArea::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "scrollContentsBy", f)
	}
}

func (ptr *QScrollArea) DisconnectScrollContentsBy() {
	defer qt.Recovering("disconnect QScrollArea::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "scrollContentsBy")
	}
}

//export callbackQScrollAreaScrollContentsBy
func callbackQScrollAreaScrollContentsBy(ptr unsafe.Pointer, ptrName *C.char, dx C.int, dy C.int) {
	defer qt.Recovering("callback QScrollArea::scrollContentsBy")

	if signal := qt.GetSignal(C.GoString(ptrName), "scrollContentsBy"); signal != nil {
		signal.(func(int, int))(int(dx), int(dy))
	} else {
		NewQScrollAreaFromPointer(ptr).ScrollContentsByDefault(int(dx), int(dy))
	}
}

func (ptr *QScrollArea) ScrollContentsBy(dx int, dy int) {
	defer qt.Recovering("QScrollArea::scrollContentsBy")

	if ptr.Pointer() != nil {
		C.QScrollArea_ScrollContentsBy(ptr.Pointer(), C.int(dx), C.int(dy))
	}
}

func (ptr *QScrollArea) ScrollContentsByDefault(dx int, dy int) {
	defer qt.Recovering("QScrollArea::scrollContentsBy")

	if ptr.Pointer() != nil {
		C.QScrollArea_ScrollContentsByDefault(ptr.Pointer(), C.int(dx), C.int(dy))
	}
}

func (ptr *QScrollArea) SizeHint() *core.QSize {
	defer qt.Recovering("QScrollArea::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QScrollArea_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScrollArea) TakeWidget() *QWidget {
	defer qt.Recovering("QScrollArea::takeWidget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QScrollArea_TakeWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScrollArea) ViewportSizeHint() *core.QSize {
	defer qt.Recovering("QScrollArea::viewportSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QScrollArea_ViewportSizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScrollArea) Widget() *QWidget {
	defer qt.Recovering("QScrollArea::widget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QScrollArea_Widget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScrollArea) DestroyQScrollArea() {
	defer qt.Recovering("QScrollArea::~QScrollArea")

	if ptr.Pointer() != nil {
		C.QScrollArea_DestroyQScrollArea(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QScrollArea) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QScrollArea::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QScrollArea) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QScrollArea::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQScrollAreaDragEnterEvent
func callbackQScrollAreaDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScrollArea::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQScrollAreaFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QScrollArea) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QScrollArea::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QScrollArea) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QScrollArea::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QScrollArea) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QScrollArea::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QScrollArea) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QScrollArea::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQScrollAreaDragLeaveEvent
func callbackQScrollAreaDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScrollArea::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQScrollAreaFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QScrollArea) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QScrollArea::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QScrollArea) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QScrollArea::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QScrollArea) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QScrollArea::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QScrollArea) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QScrollArea::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQScrollAreaDragMoveEvent
func callbackQScrollAreaDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScrollArea::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQScrollAreaFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QScrollArea) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QScrollArea::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QScrollArea) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QScrollArea::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QScrollArea) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QScrollArea::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QScrollArea) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QScrollArea::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQScrollAreaDropEvent
func callbackQScrollAreaDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScrollArea::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQScrollAreaFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QScrollArea) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QScrollArea::dropEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QScrollArea) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QScrollArea::dropEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QScrollArea) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QScrollArea::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QScrollArea) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QScrollArea::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQScrollAreaPaintEvent
func callbackQScrollAreaPaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScrollArea::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQScrollAreaFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QScrollArea) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QScrollArea::paintEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QScrollArea) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QScrollArea::paintEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QScrollArea) ConnectContextMenuEvent(f func(e *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QScrollArea::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QScrollArea) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QScrollArea::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQScrollAreaContextMenuEvent
func callbackQScrollAreaContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QScrollArea::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
	} else {
		NewQScrollAreaFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(e))
	}
}

func (ptr *QScrollArea) ContextMenuEvent(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QScrollArea::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

func (ptr *QScrollArea) ContextMenuEventDefault(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QScrollArea::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

func (ptr *QScrollArea) ConnectKeyPressEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QScrollArea::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QScrollArea) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QScrollArea::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQScrollAreaKeyPressEvent
func callbackQScrollAreaKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QScrollArea::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQScrollAreaFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QScrollArea) KeyPressEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QScrollArea::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QScrollArea) KeyPressEventDefault(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QScrollArea::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QScrollArea) ConnectMouseDoubleClickEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QScrollArea::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QScrollArea) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QScrollArea::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQScrollAreaMouseDoubleClickEvent
func callbackQScrollAreaMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QScrollArea::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQScrollAreaFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QScrollArea) MouseDoubleClickEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QScrollArea::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QScrollArea) MouseDoubleClickEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QScrollArea::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QScrollArea) ConnectMouseMoveEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QScrollArea::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QScrollArea) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QScrollArea::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQScrollAreaMouseMoveEvent
func callbackQScrollAreaMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QScrollArea::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQScrollAreaFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QScrollArea) MouseMoveEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QScrollArea::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QScrollArea) MouseMoveEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QScrollArea::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QScrollArea) ConnectMousePressEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QScrollArea::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QScrollArea) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QScrollArea::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQScrollAreaMousePressEvent
func callbackQScrollAreaMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QScrollArea::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQScrollAreaFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QScrollArea) MousePressEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QScrollArea::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QScrollArea) MousePressEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QScrollArea::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QScrollArea) ConnectMouseReleaseEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QScrollArea::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QScrollArea) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QScrollArea::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQScrollAreaMouseReleaseEvent
func callbackQScrollAreaMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QScrollArea::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQScrollAreaFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QScrollArea) MouseReleaseEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QScrollArea::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QScrollArea) MouseReleaseEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QScrollArea::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QScrollArea) ConnectSetupViewport(f func(viewport *QWidget)) {
	defer qt.Recovering("connect QScrollArea::setupViewport")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setupViewport", f)
	}
}

func (ptr *QScrollArea) DisconnectSetupViewport() {
	defer qt.Recovering("disconnect QScrollArea::setupViewport")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setupViewport")
	}
}

//export callbackQScrollAreaSetupViewport
func callbackQScrollAreaSetupViewport(ptr unsafe.Pointer, ptrName *C.char, viewport unsafe.Pointer) {
	defer qt.Recovering("callback QScrollArea::setupViewport")

	if signal := qt.GetSignal(C.GoString(ptrName), "setupViewport"); signal != nil {
		signal.(func(*QWidget))(NewQWidgetFromPointer(viewport))
	} else {
		NewQScrollAreaFromPointer(ptr).SetupViewportDefault(NewQWidgetFromPointer(viewport))
	}
}

func (ptr *QScrollArea) SetupViewport(viewport QWidget_ITF) {
	defer qt.Recovering("QScrollArea::setupViewport")

	if ptr.Pointer() != nil {
		C.QScrollArea_SetupViewport(ptr.Pointer(), PointerFromQWidget(viewport))
	}
}

func (ptr *QScrollArea) SetupViewportDefault(viewport QWidget_ITF) {
	defer qt.Recovering("QScrollArea::setupViewport")

	if ptr.Pointer() != nil {
		C.QScrollArea_SetupViewportDefault(ptr.Pointer(), PointerFromQWidget(viewport))
	}
}

func (ptr *QScrollArea) ConnectWheelEvent(f func(e *gui.QWheelEvent)) {
	defer qt.Recovering("connect QScrollArea::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QScrollArea) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QScrollArea::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQScrollAreaWheelEvent
func callbackQScrollAreaWheelEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QScrollArea::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
	} else {
		NewQScrollAreaFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(e))
	}
}

func (ptr *QScrollArea) WheelEvent(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QScrollArea::wheelEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QScrollArea) WheelEventDefault(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QScrollArea::wheelEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QScrollArea) ConnectChangeEvent(f func(ev *core.QEvent)) {
	defer qt.Recovering("connect QScrollArea::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QScrollArea) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QScrollArea::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQScrollAreaChangeEvent
func callbackQScrollAreaChangeEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QScrollArea::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
	} else {
		NewQScrollAreaFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(ev))
	}
}

func (ptr *QScrollArea) ChangeEvent(ev core.QEvent_ITF) {
	defer qt.Recovering("QScrollArea::changeEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QScrollArea) ChangeEventDefault(ev core.QEvent_ITF) {
	defer qt.Recovering("QScrollArea::changeEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QScrollArea) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QScrollArea::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QScrollArea) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QScrollArea::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQScrollAreaActionEvent
func callbackQScrollAreaActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScrollArea::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQScrollAreaFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QScrollArea) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QScrollArea::actionEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QScrollArea) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QScrollArea::actionEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QScrollArea) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QScrollArea::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QScrollArea) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QScrollArea::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQScrollAreaEnterEvent
func callbackQScrollAreaEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScrollArea::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQScrollAreaFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QScrollArea) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QScrollArea::enterEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QScrollArea) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QScrollArea::enterEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QScrollArea) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QScrollArea::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QScrollArea) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QScrollArea::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQScrollAreaFocusInEvent
func callbackQScrollAreaFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScrollArea::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQScrollAreaFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QScrollArea) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QScrollArea::focusInEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QScrollArea) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QScrollArea::focusInEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QScrollArea) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QScrollArea::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QScrollArea) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QScrollArea::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQScrollAreaFocusOutEvent
func callbackQScrollAreaFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScrollArea::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQScrollAreaFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QScrollArea) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QScrollArea::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QScrollArea) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QScrollArea::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QScrollArea) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QScrollArea::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QScrollArea) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QScrollArea::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQScrollAreaHideEvent
func callbackQScrollAreaHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScrollArea::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQScrollAreaFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QScrollArea) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QScrollArea::hideEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QScrollArea) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QScrollArea::hideEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QScrollArea) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QScrollArea::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QScrollArea) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QScrollArea::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQScrollAreaLeaveEvent
func callbackQScrollAreaLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScrollArea::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQScrollAreaFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QScrollArea) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QScrollArea::leaveEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QScrollArea) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QScrollArea::leaveEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QScrollArea) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QScrollArea::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QScrollArea) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QScrollArea::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQScrollAreaMoveEvent
func callbackQScrollAreaMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScrollArea::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQScrollAreaFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QScrollArea) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QScrollArea::moveEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QScrollArea) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QScrollArea::moveEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QScrollArea) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QScrollArea::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QScrollArea) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QScrollArea::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQScrollAreaSetVisible
func callbackQScrollAreaSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QScrollArea::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QScrollArea) SetVisible(visible bool) {
	defer qt.Recovering("QScrollArea::setVisible")

	if ptr.Pointer() != nil {
		C.QScrollArea_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QScrollArea) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QScrollArea::setVisible")

	if ptr.Pointer() != nil {
		C.QScrollArea_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QScrollArea) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QScrollArea::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QScrollArea) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QScrollArea::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQScrollAreaShowEvent
func callbackQScrollAreaShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScrollArea::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQScrollAreaFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QScrollArea) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QScrollArea::showEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QScrollArea) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QScrollArea::showEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QScrollArea) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QScrollArea::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QScrollArea) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QScrollArea::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQScrollAreaCloseEvent
func callbackQScrollAreaCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScrollArea::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQScrollAreaFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QScrollArea) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QScrollArea::closeEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QScrollArea) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QScrollArea::closeEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QScrollArea) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QScrollArea::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QScrollArea) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QScrollArea::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQScrollAreaInitPainter
func callbackQScrollAreaInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QScrollArea::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQScrollAreaFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QScrollArea) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QScrollArea::initPainter")

	if ptr.Pointer() != nil {
		C.QScrollArea_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QScrollArea) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QScrollArea::initPainter")

	if ptr.Pointer() != nil {
		C.QScrollArea_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QScrollArea) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QScrollArea::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QScrollArea) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QScrollArea::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQScrollAreaInputMethodEvent
func callbackQScrollAreaInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScrollArea::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQScrollAreaFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QScrollArea) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QScrollArea::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QScrollArea) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QScrollArea::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QScrollArea) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QScrollArea::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QScrollArea) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QScrollArea::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQScrollAreaKeyReleaseEvent
func callbackQScrollAreaKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScrollArea::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQScrollAreaFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QScrollArea) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QScrollArea::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QScrollArea) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QScrollArea::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QScrollArea) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QScrollArea::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QScrollArea) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QScrollArea::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQScrollAreaTabletEvent
func callbackQScrollAreaTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScrollArea::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQScrollAreaFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QScrollArea) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QScrollArea::tabletEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QScrollArea) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QScrollArea::tabletEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QScrollArea) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QScrollArea::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QScrollArea) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QScrollArea::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQScrollAreaTimerEvent
func callbackQScrollAreaTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScrollArea::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQScrollAreaFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QScrollArea) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QScrollArea::timerEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QScrollArea) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QScrollArea::timerEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QScrollArea) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QScrollArea::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QScrollArea) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QScrollArea::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQScrollAreaChildEvent
func callbackQScrollAreaChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScrollArea::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQScrollAreaFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QScrollArea) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QScrollArea::childEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QScrollArea) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QScrollArea::childEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QScrollArea) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QScrollArea::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QScrollArea) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QScrollArea::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQScrollAreaCustomEvent
func callbackQScrollAreaCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScrollArea::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQScrollAreaFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QScrollArea) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QScrollArea::customEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QScrollArea) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QScrollArea::customEvent")

	if ptr.Pointer() != nil {
		C.QScrollArea_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
