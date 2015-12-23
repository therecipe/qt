package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QFocusFrame struct {
	QWidget
}

type QFocusFrame_ITF interface {
	QWidget_ITF
	QFocusFrame_PTR() *QFocusFrame
}

func PointerFromQFocusFrame(ptr QFocusFrame_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFocusFrame_PTR().Pointer()
	}
	return nil
}

func NewQFocusFrameFromPointer(ptr unsafe.Pointer) *QFocusFrame {
	var n = new(QFocusFrame)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QFocusFrame_") {
		n.SetObjectName("QFocusFrame_" + qt.Identifier())
	}
	return n
}

func (ptr *QFocusFrame) QFocusFrame_PTR() *QFocusFrame {
	return ptr
}

func NewQFocusFrame(parent QWidget_ITF) *QFocusFrame {
	defer qt.Recovering("QFocusFrame::QFocusFrame")

	return NewQFocusFrameFromPointer(C.QFocusFrame_NewQFocusFrame(PointerFromQWidget(parent)))
}

func (ptr *QFocusFrame) ConnectPaintEvent(f func(v *gui.QPaintEvent)) {
	defer qt.Recovering("connect QFocusFrame::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QFocusFrame::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQFocusFramePaintEvent
func callbackQFocusFramePaintEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QFocusFrame::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(v))
		return true
	}
	return false

}

func (ptr *QFocusFrame) SetWidget(widget QWidget_ITF) {
	defer qt.Recovering("QFocusFrame::setWidget")

	if ptr.Pointer() != nil {
		C.QFocusFrame_SetWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QFocusFrame) Widget() *QWidget {
	defer qt.Recovering("QFocusFrame::widget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QFocusFrame_Widget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QFocusFrame) DestroyQFocusFrame() {
	defer qt.Recovering("QFocusFrame::~QFocusFrame")

	if ptr.Pointer() != nil {
		C.QFocusFrame_DestroyQFocusFrame(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QFocusFrame) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QFocusFrame::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QFocusFrame::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQFocusFrameActionEvent
func callbackQFocusFrameActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFocusFrame::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFocusFrame) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QFocusFrame::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QFocusFrame::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQFocusFrameDragEnterEvent
func callbackQFocusFrameDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFocusFrame::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFocusFrame) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QFocusFrame::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QFocusFrame::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQFocusFrameDragLeaveEvent
func callbackQFocusFrameDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFocusFrame::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFocusFrame) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QFocusFrame::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QFocusFrame::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQFocusFrameDragMoveEvent
func callbackQFocusFrameDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFocusFrame::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFocusFrame) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QFocusFrame::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QFocusFrame::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQFocusFrameDropEvent
func callbackQFocusFrameDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFocusFrame::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFocusFrame) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QFocusFrame::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QFocusFrame::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQFocusFrameEnterEvent
func callbackQFocusFrameEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFocusFrame::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFocusFrame) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QFocusFrame::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QFocusFrame::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQFocusFrameFocusInEvent
func callbackQFocusFrameFocusInEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFocusFrame::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFocusFrame) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QFocusFrame::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QFocusFrame::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQFocusFrameFocusOutEvent
func callbackQFocusFrameFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFocusFrame::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFocusFrame) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QFocusFrame::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QFocusFrame::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQFocusFrameHideEvent
func callbackQFocusFrameHideEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFocusFrame::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFocusFrame) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QFocusFrame::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QFocusFrame::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQFocusFrameLeaveEvent
func callbackQFocusFrameLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFocusFrame::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFocusFrame) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QFocusFrame::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QFocusFrame::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQFocusFrameMoveEvent
func callbackQFocusFrameMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFocusFrame::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFocusFrame) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QFocusFrame::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QFocusFrame) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QFocusFrame::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQFocusFrameSetVisible
func callbackQFocusFrameSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QFocusFrame::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QFocusFrame) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QFocusFrame::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QFocusFrame::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQFocusFrameShowEvent
func callbackQFocusFrameShowEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFocusFrame::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFocusFrame) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QFocusFrame::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QFocusFrame::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQFocusFrameChangeEvent
func callbackQFocusFrameChangeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFocusFrame::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFocusFrame) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QFocusFrame::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QFocusFrame::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQFocusFrameCloseEvent
func callbackQFocusFrameCloseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFocusFrame::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFocusFrame) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QFocusFrame::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QFocusFrame::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQFocusFrameContextMenuEvent
func callbackQFocusFrameContextMenuEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFocusFrame::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFocusFrame) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QFocusFrame::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QFocusFrame) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QFocusFrame::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQFocusFrameInitPainter
func callbackQFocusFrameInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QFocusFrame::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

}

func (ptr *QFocusFrame) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QFocusFrame::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QFocusFrame::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQFocusFrameInputMethodEvent
func callbackQFocusFrameInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFocusFrame::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFocusFrame) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QFocusFrame::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QFocusFrame::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQFocusFrameKeyPressEvent
func callbackQFocusFrameKeyPressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFocusFrame::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFocusFrame) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QFocusFrame::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QFocusFrame::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQFocusFrameKeyReleaseEvent
func callbackQFocusFrameKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFocusFrame::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFocusFrame) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QFocusFrame::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QFocusFrame::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQFocusFrameMouseDoubleClickEvent
func callbackQFocusFrameMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFocusFrame::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFocusFrame) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QFocusFrame::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QFocusFrame::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQFocusFrameMouseMoveEvent
func callbackQFocusFrameMouseMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFocusFrame::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFocusFrame) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QFocusFrame::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QFocusFrame::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQFocusFrameMousePressEvent
func callbackQFocusFrameMousePressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFocusFrame::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFocusFrame) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QFocusFrame::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QFocusFrame::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQFocusFrameMouseReleaseEvent
func callbackQFocusFrameMouseReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFocusFrame::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFocusFrame) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QFocusFrame::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QFocusFrame::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQFocusFrameResizeEvent
func callbackQFocusFrameResizeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFocusFrame::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFocusFrame) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QFocusFrame::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QFocusFrame::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQFocusFrameTabletEvent
func callbackQFocusFrameTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFocusFrame::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFocusFrame) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QFocusFrame::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QFocusFrame::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQFocusFrameWheelEvent
func callbackQFocusFrameWheelEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFocusFrame::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFocusFrame) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QFocusFrame::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QFocusFrame::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQFocusFrameTimerEvent
func callbackQFocusFrameTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFocusFrame::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFocusFrame) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QFocusFrame::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QFocusFrame::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQFocusFrameChildEvent
func callbackQFocusFrameChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFocusFrame::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFocusFrame) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QFocusFrame::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QFocusFrame) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QFocusFrame::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQFocusFrameCustomEvent
func callbackQFocusFrameCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFocusFrame::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
