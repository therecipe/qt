package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QSizeGrip struct {
	QWidget
}

type QSizeGrip_ITF interface {
	QWidget_ITF
	QSizeGrip_PTR() *QSizeGrip
}

func PointerFromQSizeGrip(ptr QSizeGrip_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSizeGrip_PTR().Pointer()
	}
	return nil
}

func NewQSizeGripFromPointer(ptr unsafe.Pointer) *QSizeGrip {
	var n = new(QSizeGrip)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSizeGrip_") {
		n.SetObjectName("QSizeGrip_" + qt.Identifier())
	}
	return n
}

func (ptr *QSizeGrip) QSizeGrip_PTR() *QSizeGrip {
	return ptr
}

func (ptr *QSizeGrip) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QSizeGrip::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QSizeGrip::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQSizeGripMouseMoveEvent
func callbackQSizeGripMouseMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSizeGrip::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSizeGrip) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QSizeGrip::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QSizeGrip::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQSizeGripMousePressEvent
func callbackQSizeGripMousePressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSizeGrip::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func NewQSizeGrip(parent QWidget_ITF) *QSizeGrip {
	defer qt.Recovering("QSizeGrip::QSizeGrip")

	return NewQSizeGripFromPointer(C.QSizeGrip_NewQSizeGrip(PointerFromQWidget(parent)))
}

func (ptr *QSizeGrip) ConnectHideEvent(f func(hideEvent *gui.QHideEvent)) {
	defer qt.Recovering("connect QSizeGrip::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QSizeGrip::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQSizeGripHideEvent
func callbackQSizeGripHideEvent(ptrName *C.char, hideEvent unsafe.Pointer) bool {
	defer qt.Recovering("callback QSizeGrip::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(hideEvent))
		return true
	}
	return false

}

func (ptr *QSizeGrip) ConnectMouseReleaseEvent(f func(mouseEvent *gui.QMouseEvent)) {
	defer qt.Recovering("connect QSizeGrip::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QSizeGrip::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQSizeGripMouseReleaseEvent
func callbackQSizeGripMouseReleaseEvent(ptrName *C.char, mouseEvent unsafe.Pointer) bool {
	defer qt.Recovering("callback QSizeGrip::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(mouseEvent))
		return true
	}
	return false

}

func (ptr *QSizeGrip) ConnectMoveEvent(f func(moveEvent *gui.QMoveEvent)) {
	defer qt.Recovering("connect QSizeGrip::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QSizeGrip::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQSizeGripMoveEvent
func callbackQSizeGripMoveEvent(ptrName *C.char, moveEvent unsafe.Pointer) bool {
	defer qt.Recovering("callback QSizeGrip::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(moveEvent))
		return true
	}
	return false

}

func (ptr *QSizeGrip) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QSizeGrip::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QSizeGrip::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQSizeGripPaintEvent
func callbackQSizeGripPaintEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSizeGrip::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSizeGrip) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QSizeGrip::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QSizeGrip) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QSizeGrip::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQSizeGripSetVisible
func callbackQSizeGripSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QSizeGrip::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QSizeGrip) ConnectShowEvent(f func(showEvent *gui.QShowEvent)) {
	defer qt.Recovering("connect QSizeGrip::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QSizeGrip::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQSizeGripShowEvent
func callbackQSizeGripShowEvent(ptrName *C.char, showEvent unsafe.Pointer) bool {
	defer qt.Recovering("callback QSizeGrip::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(showEvent))
		return true
	}
	return false

}

func (ptr *QSizeGrip) SizeHint() *core.QSize {
	defer qt.Recovering("QSizeGrip::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QSizeGrip_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSizeGrip) DestroyQSizeGrip() {
	defer qt.Recovering("QSizeGrip::~QSizeGrip")

	if ptr.Pointer() != nil {
		C.QSizeGrip_DestroyQSizeGrip(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSizeGrip) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QSizeGrip::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QSizeGrip::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQSizeGripActionEvent
func callbackQSizeGripActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSizeGrip::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSizeGrip) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QSizeGrip::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QSizeGrip::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQSizeGripDragEnterEvent
func callbackQSizeGripDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSizeGrip::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSizeGrip) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QSizeGrip::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QSizeGrip::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQSizeGripDragLeaveEvent
func callbackQSizeGripDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSizeGrip::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSizeGrip) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QSizeGrip::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QSizeGrip::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQSizeGripDragMoveEvent
func callbackQSizeGripDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSizeGrip::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSizeGrip) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QSizeGrip::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QSizeGrip::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQSizeGripDropEvent
func callbackQSizeGripDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSizeGrip::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSizeGrip) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSizeGrip::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QSizeGrip::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQSizeGripEnterEvent
func callbackQSizeGripEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSizeGrip::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSizeGrip) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QSizeGrip::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QSizeGrip::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQSizeGripFocusInEvent
func callbackQSizeGripFocusInEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSizeGrip::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSizeGrip) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QSizeGrip::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QSizeGrip::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQSizeGripFocusOutEvent
func callbackQSizeGripFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSizeGrip::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSizeGrip) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSizeGrip::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QSizeGrip::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQSizeGripLeaveEvent
func callbackQSizeGripLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSizeGrip::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSizeGrip) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSizeGrip::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QSizeGrip::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQSizeGripChangeEvent
func callbackQSizeGripChangeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSizeGrip::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSizeGrip) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QSizeGrip::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QSizeGrip::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQSizeGripCloseEvent
func callbackQSizeGripCloseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSizeGrip::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSizeGrip) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QSizeGrip::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QSizeGrip::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQSizeGripContextMenuEvent
func callbackQSizeGripContextMenuEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSizeGrip::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSizeGrip) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QSizeGrip::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QSizeGrip) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QSizeGrip::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQSizeGripInitPainter
func callbackQSizeGripInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QSizeGrip::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

}

func (ptr *QSizeGrip) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QSizeGrip::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QSizeGrip::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQSizeGripInputMethodEvent
func callbackQSizeGripInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSizeGrip::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSizeGrip) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QSizeGrip::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QSizeGrip::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQSizeGripKeyPressEvent
func callbackQSizeGripKeyPressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSizeGrip::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSizeGrip) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QSizeGrip::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QSizeGrip::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQSizeGripKeyReleaseEvent
func callbackQSizeGripKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSizeGrip::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSizeGrip) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QSizeGrip::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QSizeGrip::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQSizeGripMouseDoubleClickEvent
func callbackQSizeGripMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSizeGrip::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSizeGrip) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QSizeGrip::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QSizeGrip::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQSizeGripResizeEvent
func callbackQSizeGripResizeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSizeGrip::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSizeGrip) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QSizeGrip::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QSizeGrip::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQSizeGripTabletEvent
func callbackQSizeGripTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSizeGrip::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSizeGrip) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QSizeGrip::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QSizeGrip::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQSizeGripWheelEvent
func callbackQSizeGripWheelEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSizeGrip::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSizeGrip) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSizeGrip::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSizeGrip::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSizeGripTimerEvent
func callbackQSizeGripTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSizeGrip::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSizeGrip) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSizeGrip::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSizeGrip::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSizeGripChildEvent
func callbackQSizeGripChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSizeGrip::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSizeGrip) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSizeGrip::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSizeGrip) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSizeGrip::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSizeGripCustomEvent
func callbackQSizeGripCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSizeGrip::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
