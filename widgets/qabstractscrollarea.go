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
func callbackQAbstractScrollAreaDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractScrollArea::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

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
func callbackQAbstractScrollAreaDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractScrollArea::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

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
func callbackQAbstractScrollAreaDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractScrollArea::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

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
func callbackQAbstractScrollAreaDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractScrollArea::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
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
func callbackQAbstractScrollAreaPaintEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractScrollArea::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
		return true
	}
	return false

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
func callbackQAbstractScrollAreaResizeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractScrollArea::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
		return true
	}
	return false

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
func callbackQAbstractScrollAreaContextMenuEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractScrollArea::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
		return true
	}
	return false

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
func callbackQAbstractScrollAreaKeyPressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractScrollArea::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
		return true
	}
	return false

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
func callbackQAbstractScrollAreaMouseDoubleClickEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractScrollArea::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

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
func callbackQAbstractScrollAreaMouseMoveEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractScrollArea::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

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
func callbackQAbstractScrollAreaMousePressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractScrollArea::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

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
func callbackQAbstractScrollAreaMouseReleaseEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractScrollArea::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

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
func callbackQAbstractScrollAreaScrollContentsBy(ptrName *C.char, dx C.int, dy C.int) bool {
	defer qt.Recovering("callback QAbstractScrollArea::scrollContentsBy")

	if signal := qt.GetSignal(C.GoString(ptrName), "scrollContentsBy"); signal != nil {
		signal.(func(int, int))(int(dx), int(dy))
		return true
	}
	return false

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
func callbackQAbstractScrollAreaSetupViewport(ptrName *C.char, viewport unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractScrollArea::setupViewport")

	if signal := qt.GetSignal(C.GoString(ptrName), "setupViewport"); signal != nil {
		signal.(func(*QWidget))(NewQWidgetFromPointer(viewport))
		return true
	}
	return false

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
func callbackQAbstractScrollAreaWheelEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractScrollArea::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
		return true
	}
	return false

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
func callbackQAbstractScrollAreaChangeEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractScrollArea::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
		return true
	}
	return false

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
func callbackQAbstractScrollAreaActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractScrollArea::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

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
func callbackQAbstractScrollAreaEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractScrollArea::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

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
func callbackQAbstractScrollAreaFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractScrollArea::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

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
func callbackQAbstractScrollAreaHideEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractScrollArea::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
		return true
	}
	return false

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
func callbackQAbstractScrollAreaLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractScrollArea::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

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
func callbackQAbstractScrollAreaMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractScrollArea::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

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
func callbackQAbstractScrollAreaSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QAbstractScrollArea::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

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
func callbackQAbstractScrollAreaShowEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractScrollArea::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
		return true
	}
	return false

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
func callbackQAbstractScrollAreaCloseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractScrollArea::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQAbstractScrollAreaInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractScrollArea::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

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
func callbackQAbstractScrollAreaInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractScrollArea::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

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
func callbackQAbstractScrollAreaKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractScrollArea::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

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
func callbackQAbstractScrollAreaTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractScrollArea::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

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
func callbackQAbstractScrollAreaTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractScrollArea::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

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
func callbackQAbstractScrollAreaChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractScrollArea::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

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
func callbackQAbstractScrollAreaCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractScrollArea::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
