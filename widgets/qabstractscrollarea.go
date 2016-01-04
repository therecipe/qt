package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QAbstractScrollArea struct {
	QFrame
}

type QAbstractScrollArea_ITF interface {
	QFrame_ITF
	QAbstractScrollArea_PTR() *QAbstractScrollArea
}

func PointerFromQAbstractScrollArea(ptr QAbstractScrollArea_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractScrollArea_PTR().Pointer()
	}
	return nil
}

func NewQAbstractScrollAreaFromPointer(ptr unsafe.Pointer) *QAbstractScrollArea {
	var n = new(QAbstractScrollArea)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAbstractScrollArea_") {
		n.SetObjectName("QAbstractScrollArea_" + qt.Identifier())
	}
	return n
}

func (ptr *QAbstractScrollArea) QAbstractScrollArea_PTR() *QAbstractScrollArea {
	return ptr
}

//QAbstractScrollArea::SizeAdjustPolicy
type QAbstractScrollArea__SizeAdjustPolicy int64

const (
	QAbstractScrollArea__AdjustIgnored               = QAbstractScrollArea__SizeAdjustPolicy(0)
	QAbstractScrollArea__AdjustToContentsOnFirstShow = QAbstractScrollArea__SizeAdjustPolicy(1)
	QAbstractScrollArea__AdjustToContents            = QAbstractScrollArea__SizeAdjustPolicy(2)
)

func (ptr *QAbstractScrollArea) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QAbstractScrollArea::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QAbstractScrollArea) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QAbstractScrollArea::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQAbstractScrollAreaDragEnterEvent
func callbackQAbstractScrollAreaDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractScrollArea::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQAbstractScrollAreaFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QAbstractScrollArea) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QAbstractScrollArea) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QAbstractScrollArea) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QAbstractScrollArea::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QAbstractScrollArea) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QAbstractScrollArea::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQAbstractScrollAreaDragLeaveEvent
func callbackQAbstractScrollAreaDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractScrollArea::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQAbstractScrollAreaFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QAbstractScrollArea) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QAbstractScrollArea) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QAbstractScrollArea) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QAbstractScrollArea::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QAbstractScrollArea) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QAbstractScrollArea::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQAbstractScrollAreaDragMoveEvent
func callbackQAbstractScrollAreaDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractScrollArea::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQAbstractScrollAreaFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QAbstractScrollArea) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QAbstractScrollArea) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QAbstractScrollArea) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QAbstractScrollArea::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QAbstractScrollArea) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QAbstractScrollArea::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQAbstractScrollAreaDropEvent
func callbackQAbstractScrollAreaDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractScrollArea::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQAbstractScrollAreaFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QAbstractScrollArea) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::dropEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QAbstractScrollArea) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::dropEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QAbstractScrollArea) Event(event core.QEvent_ITF) bool {
	defer qt.Recovering("QAbstractScrollArea::event")

	if ptr.Pointer() != nil {
		return C.QAbstractScrollArea_Event(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QAbstractScrollArea) HorizontalScrollBarPolicy() core.Qt__ScrollBarPolicy {
	defer qt.Recovering("QAbstractScrollArea::horizontalScrollBarPolicy")

	if ptr.Pointer() != nil {
		return core.Qt__ScrollBarPolicy(C.QAbstractScrollArea_HorizontalScrollBarPolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractScrollArea) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QAbstractScrollArea::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QAbstractScrollArea) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QAbstractScrollArea::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQAbstractScrollAreaPaintEvent
func callbackQAbstractScrollAreaPaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractScrollArea::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQAbstractScrollAreaFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QAbstractScrollArea) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::paintEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QAbstractScrollArea) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::paintEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QAbstractScrollArea) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QAbstractScrollArea::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QAbstractScrollArea) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QAbstractScrollArea::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQAbstractScrollAreaResizeEvent
func callbackQAbstractScrollAreaResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractScrollArea::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQAbstractScrollAreaFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QAbstractScrollArea) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::resizeEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QAbstractScrollArea) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::resizeEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QAbstractScrollArea) SetHorizontalScrollBarPolicy(v core.Qt__ScrollBarPolicy) {
	defer qt.Recovering("QAbstractScrollArea::setHorizontalScrollBarPolicy")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_SetHorizontalScrollBarPolicy(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QAbstractScrollArea) SetSizeAdjustPolicy(policy QAbstractScrollArea__SizeAdjustPolicy) {
	defer qt.Recovering("QAbstractScrollArea::setSizeAdjustPolicy")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_SetSizeAdjustPolicy(ptr.Pointer(), C.int(policy))
	}
}

func (ptr *QAbstractScrollArea) SetVerticalScrollBarPolicy(v core.Qt__ScrollBarPolicy) {
	defer qt.Recovering("QAbstractScrollArea::setVerticalScrollBarPolicy")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_SetVerticalScrollBarPolicy(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QAbstractScrollArea) SizeAdjustPolicy() QAbstractScrollArea__SizeAdjustPolicy {
	defer qt.Recovering("QAbstractScrollArea::sizeAdjustPolicy")

	if ptr.Pointer() != nil {
		return QAbstractScrollArea__SizeAdjustPolicy(C.QAbstractScrollArea_SizeAdjustPolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractScrollArea) VerticalScrollBarPolicy() core.Qt__ScrollBarPolicy {
	defer qt.Recovering("QAbstractScrollArea::verticalScrollBarPolicy")

	if ptr.Pointer() != nil {
		return core.Qt__ScrollBarPolicy(C.QAbstractScrollArea_VerticalScrollBarPolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractScrollArea) ViewportEvent(event core.QEvent_ITF) bool {
	defer qt.Recovering("QAbstractScrollArea::viewportEvent")

	if ptr.Pointer() != nil {
		return C.QAbstractScrollArea_ViewportEvent(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func NewQAbstractScrollArea(parent QWidget_ITF) *QAbstractScrollArea {
	defer qt.Recovering("QAbstractScrollArea::QAbstractScrollArea")

	return NewQAbstractScrollAreaFromPointer(C.QAbstractScrollArea_NewQAbstractScrollArea(PointerFromQWidget(parent)))
}

func (ptr *QAbstractScrollArea) AddScrollBarWidget(widget QWidget_ITF, alignment core.Qt__AlignmentFlag) {
	defer qt.Recovering("QAbstractScrollArea::addScrollBarWidget")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_AddScrollBarWidget(ptr.Pointer(), PointerFromQWidget(widget), C.int(alignment))
	}
}

func (ptr *QAbstractScrollArea) ConnectContextMenuEvent(f func(e *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QAbstractScrollArea::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QAbstractScrollArea) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QAbstractScrollArea::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQAbstractScrollAreaContextMenuEvent
func callbackQAbstractScrollAreaContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractScrollArea::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
	} else {
		NewQAbstractScrollAreaFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(e))
	}
}

func (ptr *QAbstractScrollArea) ContextMenuEvent(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

func (ptr *QAbstractScrollArea) ContextMenuEventDefault(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

func (ptr *QAbstractScrollArea) CornerWidget() *QWidget {
	defer qt.Recovering("QAbstractScrollArea::cornerWidget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QAbstractScrollArea_CornerWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractScrollArea) HorizontalScrollBar() *QScrollBar {
	defer qt.Recovering("QAbstractScrollArea::horizontalScrollBar")

	if ptr.Pointer() != nil {
		return NewQScrollBarFromPointer(C.QAbstractScrollArea_HorizontalScrollBar(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractScrollArea) ConnectKeyPressEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QAbstractScrollArea::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QAbstractScrollArea) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QAbstractScrollArea::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQAbstractScrollAreaKeyPressEvent
func callbackQAbstractScrollAreaKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractScrollArea::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQAbstractScrollAreaFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QAbstractScrollArea) KeyPressEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QAbstractScrollArea) KeyPressEventDefault(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QAbstractScrollArea) MaximumViewportSize() *core.QSize {
	defer qt.Recovering("QAbstractScrollArea::maximumViewportSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QAbstractScrollArea_MaximumViewportSize(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractScrollArea) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QAbstractScrollArea::minimumSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QAbstractScrollArea_MinimumSizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractScrollArea) ConnectMouseDoubleClickEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QAbstractScrollArea::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QAbstractScrollArea) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QAbstractScrollArea::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQAbstractScrollAreaMouseDoubleClickEvent
func callbackQAbstractScrollAreaMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractScrollArea::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQAbstractScrollAreaFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QAbstractScrollArea) MouseDoubleClickEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QAbstractScrollArea) MouseDoubleClickEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QAbstractScrollArea) ConnectMouseMoveEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QAbstractScrollArea::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QAbstractScrollArea) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QAbstractScrollArea::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQAbstractScrollAreaMouseMoveEvent
func callbackQAbstractScrollAreaMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractScrollArea::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQAbstractScrollAreaFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QAbstractScrollArea) MouseMoveEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QAbstractScrollArea) MouseMoveEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QAbstractScrollArea) ConnectMousePressEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QAbstractScrollArea::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QAbstractScrollArea) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QAbstractScrollArea::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQAbstractScrollAreaMousePressEvent
func callbackQAbstractScrollAreaMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractScrollArea::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQAbstractScrollAreaFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QAbstractScrollArea) MousePressEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QAbstractScrollArea) MousePressEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QAbstractScrollArea) ConnectMouseReleaseEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QAbstractScrollArea::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QAbstractScrollArea) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QAbstractScrollArea::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQAbstractScrollAreaMouseReleaseEvent
func callbackQAbstractScrollAreaMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractScrollArea::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQAbstractScrollAreaFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QAbstractScrollArea) MouseReleaseEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QAbstractScrollArea) MouseReleaseEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QAbstractScrollArea) ConnectScrollContentsBy(f func(dx int, dy int)) {
	defer qt.Recovering("connect QAbstractScrollArea::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "scrollContentsBy", f)
	}
}

func (ptr *QAbstractScrollArea) DisconnectScrollContentsBy() {
	defer qt.Recovering("disconnect QAbstractScrollArea::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "scrollContentsBy")
	}
}

//export callbackQAbstractScrollAreaScrollContentsBy
func callbackQAbstractScrollAreaScrollContentsBy(ptr unsafe.Pointer, ptrName *C.char, dx C.int, dy C.int) {
	defer qt.Recovering("callback QAbstractScrollArea::scrollContentsBy")

	if signal := qt.GetSignal(C.GoString(ptrName), "scrollContentsBy"); signal != nil {
		signal.(func(int, int))(int(dx), int(dy))
	} else {
		NewQAbstractScrollAreaFromPointer(ptr).ScrollContentsByDefault(int(dx), int(dy))
	}
}

func (ptr *QAbstractScrollArea) ScrollContentsBy(dx int, dy int) {
	defer qt.Recovering("QAbstractScrollArea::scrollContentsBy")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_ScrollContentsBy(ptr.Pointer(), C.int(dx), C.int(dy))
	}
}

func (ptr *QAbstractScrollArea) ScrollContentsByDefault(dx int, dy int) {
	defer qt.Recovering("QAbstractScrollArea::scrollContentsBy")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_ScrollContentsByDefault(ptr.Pointer(), C.int(dx), C.int(dy))
	}
}

func (ptr *QAbstractScrollArea) SetCornerWidget(widget QWidget_ITF) {
	defer qt.Recovering("QAbstractScrollArea::setCornerWidget")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_SetCornerWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QAbstractScrollArea) SetHorizontalScrollBar(scrollBar QScrollBar_ITF) {
	defer qt.Recovering("QAbstractScrollArea::setHorizontalScrollBar")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_SetHorizontalScrollBar(ptr.Pointer(), PointerFromQScrollBar(scrollBar))
	}
}

func (ptr *QAbstractScrollArea) SetVerticalScrollBar(scrollBar QScrollBar_ITF) {
	defer qt.Recovering("QAbstractScrollArea::setVerticalScrollBar")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_SetVerticalScrollBar(ptr.Pointer(), PointerFromQScrollBar(scrollBar))
	}
}

func (ptr *QAbstractScrollArea) SetViewport(widget QWidget_ITF) {
	defer qt.Recovering("QAbstractScrollArea::setViewport")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_SetViewport(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QAbstractScrollArea) ConnectSetupViewport(f func(viewport *QWidget)) {
	defer qt.Recovering("connect QAbstractScrollArea::setupViewport")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setupViewport", f)
	}
}

func (ptr *QAbstractScrollArea) DisconnectSetupViewport() {
	defer qt.Recovering("disconnect QAbstractScrollArea::setupViewport")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setupViewport")
	}
}

//export callbackQAbstractScrollAreaSetupViewport
func callbackQAbstractScrollAreaSetupViewport(ptr unsafe.Pointer, ptrName *C.char, viewport unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractScrollArea::setupViewport")

	if signal := qt.GetSignal(C.GoString(ptrName), "setupViewport"); signal != nil {
		signal.(func(*QWidget))(NewQWidgetFromPointer(viewport))
	} else {
		NewQAbstractScrollAreaFromPointer(ptr).SetupViewportDefault(NewQWidgetFromPointer(viewport))
	}
}

func (ptr *QAbstractScrollArea) SetupViewport(viewport QWidget_ITF) {
	defer qt.Recovering("QAbstractScrollArea::setupViewport")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_SetupViewport(ptr.Pointer(), PointerFromQWidget(viewport))
	}
}

func (ptr *QAbstractScrollArea) SetupViewportDefault(viewport QWidget_ITF) {
	defer qt.Recovering("QAbstractScrollArea::setupViewport")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_SetupViewportDefault(ptr.Pointer(), PointerFromQWidget(viewport))
	}
}

func (ptr *QAbstractScrollArea) SizeHint() *core.QSize {
	defer qt.Recovering("QAbstractScrollArea::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QAbstractScrollArea_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractScrollArea) VerticalScrollBar() *QScrollBar {
	defer qt.Recovering("QAbstractScrollArea::verticalScrollBar")

	if ptr.Pointer() != nil {
		return NewQScrollBarFromPointer(C.QAbstractScrollArea_VerticalScrollBar(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractScrollArea) Viewport() *QWidget {
	defer qt.Recovering("QAbstractScrollArea::viewport")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QAbstractScrollArea_Viewport(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractScrollArea) ViewportSizeHint() *core.QSize {
	defer qt.Recovering("QAbstractScrollArea::viewportSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QAbstractScrollArea_ViewportSizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractScrollArea) ConnectWheelEvent(f func(e *gui.QWheelEvent)) {
	defer qt.Recovering("connect QAbstractScrollArea::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QAbstractScrollArea) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QAbstractScrollArea::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQAbstractScrollAreaWheelEvent
func callbackQAbstractScrollAreaWheelEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractScrollArea::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
	} else {
		NewQAbstractScrollAreaFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(e))
	}
}

func (ptr *QAbstractScrollArea) WheelEvent(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::wheelEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QAbstractScrollArea) WheelEventDefault(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::wheelEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QAbstractScrollArea) DestroyQAbstractScrollArea() {
	defer qt.Recovering("QAbstractScrollArea::~QAbstractScrollArea")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_DestroyQAbstractScrollArea(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractScrollArea) ConnectChangeEvent(f func(ev *core.QEvent)) {
	defer qt.Recovering("connect QAbstractScrollArea::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QAbstractScrollArea) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QAbstractScrollArea::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQAbstractScrollAreaChangeEvent
func callbackQAbstractScrollAreaChangeEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractScrollArea::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
	} else {
		NewQAbstractScrollAreaFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(ev))
	}
}

func (ptr *QAbstractScrollArea) ChangeEvent(ev core.QEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::changeEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QAbstractScrollArea) ChangeEventDefault(ev core.QEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::changeEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QAbstractScrollArea) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QAbstractScrollArea::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QAbstractScrollArea) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QAbstractScrollArea::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQAbstractScrollAreaActionEvent
func callbackQAbstractScrollAreaActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractScrollArea::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQAbstractScrollAreaFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QAbstractScrollArea) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::actionEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QAbstractScrollArea) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::actionEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QAbstractScrollArea) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAbstractScrollArea::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QAbstractScrollArea) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QAbstractScrollArea::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQAbstractScrollAreaEnterEvent
func callbackQAbstractScrollAreaEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractScrollArea::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAbstractScrollAreaFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAbstractScrollArea) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::enterEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAbstractScrollArea) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::enterEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAbstractScrollArea) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QAbstractScrollArea::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QAbstractScrollArea) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QAbstractScrollArea::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQAbstractScrollAreaFocusInEvent
func callbackQAbstractScrollAreaFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractScrollArea::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQAbstractScrollAreaFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QAbstractScrollArea) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::focusInEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QAbstractScrollArea) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::focusInEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QAbstractScrollArea) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QAbstractScrollArea::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QAbstractScrollArea) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QAbstractScrollArea::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQAbstractScrollAreaFocusOutEvent
func callbackQAbstractScrollAreaFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractScrollArea::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQAbstractScrollAreaFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QAbstractScrollArea) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QAbstractScrollArea) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QAbstractScrollArea) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QAbstractScrollArea::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QAbstractScrollArea) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QAbstractScrollArea::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQAbstractScrollAreaHideEvent
func callbackQAbstractScrollAreaHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractScrollArea::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQAbstractScrollAreaFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QAbstractScrollArea) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::hideEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QAbstractScrollArea) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::hideEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QAbstractScrollArea) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAbstractScrollArea::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QAbstractScrollArea) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QAbstractScrollArea::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQAbstractScrollAreaLeaveEvent
func callbackQAbstractScrollAreaLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractScrollArea::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAbstractScrollAreaFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAbstractScrollArea) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::leaveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAbstractScrollArea) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::leaveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAbstractScrollArea) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QAbstractScrollArea::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QAbstractScrollArea) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QAbstractScrollArea::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQAbstractScrollAreaMoveEvent
func callbackQAbstractScrollAreaMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractScrollArea::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQAbstractScrollAreaFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QAbstractScrollArea) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::moveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QAbstractScrollArea) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::moveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QAbstractScrollArea) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QAbstractScrollArea::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QAbstractScrollArea) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QAbstractScrollArea::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQAbstractScrollAreaSetVisible
func callbackQAbstractScrollAreaSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QAbstractScrollArea::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QAbstractScrollArea) SetVisible(visible bool) {
	defer qt.Recovering("QAbstractScrollArea::setVisible")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QAbstractScrollArea) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QAbstractScrollArea::setVisible")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QAbstractScrollArea) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QAbstractScrollArea::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QAbstractScrollArea) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QAbstractScrollArea::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQAbstractScrollAreaShowEvent
func callbackQAbstractScrollAreaShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractScrollArea::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQAbstractScrollAreaFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QAbstractScrollArea) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::showEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QAbstractScrollArea) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::showEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QAbstractScrollArea) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QAbstractScrollArea::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QAbstractScrollArea) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QAbstractScrollArea::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQAbstractScrollAreaCloseEvent
func callbackQAbstractScrollAreaCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractScrollArea::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQAbstractScrollAreaFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QAbstractScrollArea) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::closeEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QAbstractScrollArea) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::closeEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QAbstractScrollArea) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QAbstractScrollArea::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QAbstractScrollArea) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QAbstractScrollArea::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQAbstractScrollAreaInitPainter
func callbackQAbstractScrollAreaInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractScrollArea::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQAbstractScrollAreaFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QAbstractScrollArea) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QAbstractScrollArea::initPainter")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QAbstractScrollArea) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QAbstractScrollArea::initPainter")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QAbstractScrollArea) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QAbstractScrollArea::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QAbstractScrollArea) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QAbstractScrollArea::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQAbstractScrollAreaInputMethodEvent
func callbackQAbstractScrollAreaInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractScrollArea::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQAbstractScrollAreaFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QAbstractScrollArea) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QAbstractScrollArea) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QAbstractScrollArea) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QAbstractScrollArea::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QAbstractScrollArea) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QAbstractScrollArea::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQAbstractScrollAreaKeyReleaseEvent
func callbackQAbstractScrollAreaKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractScrollArea::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQAbstractScrollAreaFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QAbstractScrollArea) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QAbstractScrollArea) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QAbstractScrollArea) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QAbstractScrollArea::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QAbstractScrollArea) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QAbstractScrollArea::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQAbstractScrollAreaTabletEvent
func callbackQAbstractScrollAreaTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractScrollArea::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQAbstractScrollAreaFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QAbstractScrollArea) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::tabletEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QAbstractScrollArea) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::tabletEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QAbstractScrollArea) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAbstractScrollArea::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAbstractScrollArea) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAbstractScrollArea::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAbstractScrollAreaTimerEvent
func callbackQAbstractScrollAreaTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractScrollArea::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAbstractScrollAreaFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAbstractScrollArea) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::timerEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAbstractScrollArea) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::timerEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAbstractScrollArea) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAbstractScrollArea::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAbstractScrollArea) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAbstractScrollArea::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAbstractScrollAreaChildEvent
func callbackQAbstractScrollAreaChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractScrollArea::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAbstractScrollAreaFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAbstractScrollArea) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::childEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAbstractScrollArea) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::childEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAbstractScrollArea) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAbstractScrollArea::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAbstractScrollArea) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAbstractScrollArea::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAbstractScrollAreaCustomEvent
func callbackQAbstractScrollAreaCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractScrollArea::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAbstractScrollAreaFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAbstractScrollArea) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::customEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAbstractScrollArea) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractScrollArea::customEvent")

	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
