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
func callbackQScrollAreaResizeEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollArea::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(v))
		return true
	}
	return false

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
func callbackQScrollAreaScrollContentsBy(ptrName *C.char, dx C.int, dy C.int) bool {
	defer qt.Recovering("callback QScrollArea::scrollContentsBy")

	if signal := qt.GetSignal(C.GoString(ptrName), "scrollContentsBy"); signal != nil {
		signal.(func(int, int))(int(dx), int(dy))
		return true
	}
	return false

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
func callbackQScrollAreaDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollArea::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

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
func callbackQScrollAreaDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollArea::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

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
func callbackQScrollAreaDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollArea::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

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
func callbackQScrollAreaDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollArea::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

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
func callbackQScrollAreaPaintEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollArea::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
		return true
	}
	return false

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
func callbackQScrollAreaContextMenuEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollArea::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
		return true
	}
	return false

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
func callbackQScrollAreaKeyPressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollArea::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
		return true
	}
	return false

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
func callbackQScrollAreaMouseDoubleClickEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollArea::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

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
func callbackQScrollAreaMouseMoveEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollArea::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

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
func callbackQScrollAreaMousePressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollArea::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

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
func callbackQScrollAreaMouseReleaseEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollArea::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

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
func callbackQScrollAreaSetupViewport(ptrName *C.char, viewport unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollArea::setupViewport")

	if signal := qt.GetSignal(C.GoString(ptrName), "setupViewport"); signal != nil {
		signal.(func(*QWidget))(NewQWidgetFromPointer(viewport))
		return true
	}
	return false

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
func callbackQScrollAreaWheelEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollArea::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
		return true
	}
	return false

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
func callbackQScrollAreaChangeEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollArea::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
		return true
	}
	return false

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
func callbackQScrollAreaActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollArea::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

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
func callbackQScrollAreaEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollArea::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

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
func callbackQScrollAreaFocusInEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollArea::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

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
func callbackQScrollAreaFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollArea::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

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
func callbackQScrollAreaHideEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollArea::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
		return true
	}
	return false

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
func callbackQScrollAreaLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollArea::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

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
func callbackQScrollAreaMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollArea::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

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
func callbackQScrollAreaSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QScrollArea::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

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
func callbackQScrollAreaShowEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollArea::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
		return true
	}
	return false

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
func callbackQScrollAreaCloseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollArea::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQScrollAreaInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollArea::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

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
func callbackQScrollAreaInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollArea::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

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
func callbackQScrollAreaKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollArea::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

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
func callbackQScrollAreaTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollArea::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

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
func callbackQScrollAreaTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollArea::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

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
func callbackQScrollAreaChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollArea::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

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
func callbackQScrollAreaCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollArea::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
