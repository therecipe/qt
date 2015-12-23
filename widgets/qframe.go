package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QFrame struct {
	QWidget
}

type QFrame_ITF interface {
	QWidget_ITF
	QFrame_PTR() *QFrame
}

func PointerFromQFrame(ptr QFrame_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFrame_PTR().Pointer()
	}
	return nil
}

func NewQFrameFromPointer(ptr unsafe.Pointer) *QFrame {
	var n = new(QFrame)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QFrame_") {
		n.SetObjectName("QFrame_" + qt.Identifier())
	}
	return n
}

func (ptr *QFrame) QFrame_PTR() *QFrame {
	return ptr
}

//QFrame::Shadow
type QFrame__Shadow int64

const (
	QFrame__Plain  = QFrame__Shadow(0x0010)
	QFrame__Raised = QFrame__Shadow(0x0020)
	QFrame__Sunken = QFrame__Shadow(0x0030)
)

//QFrame::Shape
type QFrame__Shape int64

const (
	QFrame__NoFrame     = QFrame__Shape(0)
	QFrame__Box         = QFrame__Shape(0x0001)
	QFrame__Panel       = QFrame__Shape(0x0002)
	QFrame__WinPanel    = QFrame__Shape(0x0003)
	QFrame__HLine       = QFrame__Shape(0x0004)
	QFrame__VLine       = QFrame__Shape(0x0005)
	QFrame__StyledPanel = QFrame__Shape(0x0006)
)

//QFrame::StyleMask
type QFrame__StyleMask int64

var (
	QFrame__Shadow_Mask = QFrame__StyleMask(0x00f0)
	QFrame__Shape_Mask  = QFrame__StyleMask(0x000f)
)

func (ptr *QFrame) FrameRect() *core.QRect {
	defer qt.Recovering("QFrame::frameRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QFrame_FrameRect(ptr.Pointer()))
	}
	return nil
}

func (ptr *QFrame) FrameShadow() QFrame__Shadow {
	defer qt.Recovering("QFrame::frameShadow")

	if ptr.Pointer() != nil {
		return QFrame__Shadow(C.QFrame_FrameShadow(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFrame) FrameShape() QFrame__Shape {
	defer qt.Recovering("QFrame::frameShape")

	if ptr.Pointer() != nil {
		return QFrame__Shape(C.QFrame_FrameShape(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFrame) FrameWidth() int {
	defer qt.Recovering("QFrame::frameWidth")

	if ptr.Pointer() != nil {
		return int(C.QFrame_FrameWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFrame) LineWidth() int {
	defer qt.Recovering("QFrame::lineWidth")

	if ptr.Pointer() != nil {
		return int(C.QFrame_LineWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFrame) MidLineWidth() int {
	defer qt.Recovering("QFrame::midLineWidth")

	if ptr.Pointer() != nil {
		return int(C.QFrame_MidLineWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFrame) SetFrameRect(v core.QRect_ITF) {
	defer qt.Recovering("QFrame::setFrameRect")

	if ptr.Pointer() != nil {
		C.QFrame_SetFrameRect(ptr.Pointer(), core.PointerFromQRect(v))
	}
}

func (ptr *QFrame) SetFrameShadow(v QFrame__Shadow) {
	defer qt.Recovering("QFrame::setFrameShadow")

	if ptr.Pointer() != nil {
		C.QFrame_SetFrameShadow(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QFrame) SetFrameShape(v QFrame__Shape) {
	defer qt.Recovering("QFrame::setFrameShape")

	if ptr.Pointer() != nil {
		C.QFrame_SetFrameShape(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QFrame) SetLineWidth(v int) {
	defer qt.Recovering("QFrame::setLineWidth")

	if ptr.Pointer() != nil {
		C.QFrame_SetLineWidth(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QFrame) SetMidLineWidth(v int) {
	defer qt.Recovering("QFrame::setMidLineWidth")

	if ptr.Pointer() != nil {
		C.QFrame_SetMidLineWidth(ptr.Pointer(), C.int(v))
	}
}

func NewQFrame(parent QWidget_ITF, f core.Qt__WindowType) *QFrame {
	defer qt.Recovering("QFrame::QFrame")

	return NewQFrameFromPointer(C.QFrame_NewQFrame(PointerFromQWidget(parent), C.int(f)))
}

func (ptr *QFrame) ConnectChangeEvent(f func(ev *core.QEvent)) {
	defer qt.Recovering("connect QFrame::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QFrame) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QFrame::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQFrameChangeEvent
func callbackQFrameChangeEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QFrame::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QFrame) FrameStyle() int {
	defer qt.Recovering("QFrame::frameStyle")

	if ptr.Pointer() != nil {
		return int(C.QFrame_FrameStyle(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFrame) ConnectPaintEvent(f func(v *gui.QPaintEvent)) {
	defer qt.Recovering("connect QFrame::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QFrame) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QFrame::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQFramePaintEvent
func callbackQFramePaintEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QFrame::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(v))
		return true
	}
	return false

}

func (ptr *QFrame) SetFrameStyle(style int) {
	defer qt.Recovering("QFrame::setFrameStyle")

	if ptr.Pointer() != nil {
		C.QFrame_SetFrameStyle(ptr.Pointer(), C.int(style))
	}
}

func (ptr *QFrame) SizeHint() *core.QSize {
	defer qt.Recovering("QFrame::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QFrame_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QFrame) DestroyQFrame() {
	defer qt.Recovering("QFrame::~QFrame")

	if ptr.Pointer() != nil {
		C.QFrame_DestroyQFrame(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QFrame) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QFrame::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QFrame) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QFrame::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQFrameActionEvent
func callbackQFrameActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFrame::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFrame) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QFrame::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QFrame) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QFrame::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQFrameDragEnterEvent
func callbackQFrameDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFrame::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFrame) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QFrame::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QFrame) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QFrame::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQFrameDragLeaveEvent
func callbackQFrameDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFrame::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFrame) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QFrame::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QFrame) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QFrame::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQFrameDragMoveEvent
func callbackQFrameDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFrame::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFrame) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QFrame::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QFrame) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QFrame::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQFrameDropEvent
func callbackQFrameDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFrame::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFrame) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QFrame::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QFrame) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QFrame::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQFrameEnterEvent
func callbackQFrameEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFrame::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFrame) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QFrame::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QFrame) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QFrame::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQFrameFocusInEvent
func callbackQFrameFocusInEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFrame::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFrame) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QFrame::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QFrame) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QFrame::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQFrameFocusOutEvent
func callbackQFrameFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFrame::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFrame) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QFrame::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QFrame) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QFrame::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQFrameHideEvent
func callbackQFrameHideEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFrame::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFrame) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QFrame::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QFrame) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QFrame::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQFrameLeaveEvent
func callbackQFrameLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFrame::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFrame) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QFrame::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QFrame) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QFrame::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQFrameMoveEvent
func callbackQFrameMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFrame::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFrame) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QFrame::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QFrame) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QFrame::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQFrameSetVisible
func callbackQFrameSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QFrame::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QFrame) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QFrame::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QFrame) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QFrame::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQFrameShowEvent
func callbackQFrameShowEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFrame::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFrame) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QFrame::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QFrame) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QFrame::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQFrameCloseEvent
func callbackQFrameCloseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFrame::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFrame) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QFrame::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QFrame) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QFrame::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQFrameContextMenuEvent
func callbackQFrameContextMenuEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFrame::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFrame) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QFrame::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QFrame) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QFrame::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQFrameInitPainter
func callbackQFrameInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QFrame::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

}

func (ptr *QFrame) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QFrame::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QFrame) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QFrame::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQFrameInputMethodEvent
func callbackQFrameInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFrame::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFrame) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QFrame::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QFrame) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QFrame::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQFrameKeyPressEvent
func callbackQFrameKeyPressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFrame::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFrame) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QFrame::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QFrame) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QFrame::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQFrameKeyReleaseEvent
func callbackQFrameKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFrame::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFrame) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QFrame::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QFrame) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QFrame::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQFrameMouseDoubleClickEvent
func callbackQFrameMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFrame::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFrame) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QFrame::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QFrame) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QFrame::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQFrameMouseMoveEvent
func callbackQFrameMouseMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFrame::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFrame) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QFrame::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QFrame) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QFrame::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQFrameMousePressEvent
func callbackQFrameMousePressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFrame::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFrame) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QFrame::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QFrame) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QFrame::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQFrameMouseReleaseEvent
func callbackQFrameMouseReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFrame::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFrame) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QFrame::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QFrame) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QFrame::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQFrameResizeEvent
func callbackQFrameResizeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFrame::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFrame) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QFrame::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QFrame) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QFrame::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQFrameTabletEvent
func callbackQFrameTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFrame::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFrame) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QFrame::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QFrame) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QFrame::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQFrameWheelEvent
func callbackQFrameWheelEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFrame::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFrame) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QFrame::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QFrame) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QFrame::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQFrameTimerEvent
func callbackQFrameTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFrame::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFrame) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QFrame::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QFrame) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QFrame::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQFrameChildEvent
func callbackQFrameChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFrame::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QFrame) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QFrame::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QFrame) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QFrame::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQFrameCustomEvent
func callbackQFrameCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QFrame::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
