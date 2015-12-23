package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QErrorMessage struct {
	QDialog
}

type QErrorMessage_ITF interface {
	QDialog_ITF
	QErrorMessage_PTR() *QErrorMessage
}

func PointerFromQErrorMessage(ptr QErrorMessage_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QErrorMessage_PTR().Pointer()
	}
	return nil
}

func NewQErrorMessageFromPointer(ptr unsafe.Pointer) *QErrorMessage {
	var n = new(QErrorMessage)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QErrorMessage_") {
		n.SetObjectName("QErrorMessage_" + qt.Identifier())
	}
	return n
}

func (ptr *QErrorMessage) QErrorMessage_PTR() *QErrorMessage {
	return ptr
}

func NewQErrorMessage(parent QWidget_ITF) *QErrorMessage {
	defer qt.Recovering("QErrorMessage::QErrorMessage")

	return NewQErrorMessageFromPointer(C.QErrorMessage_NewQErrorMessage(PointerFromQWidget(parent)))
}

func (ptr *QErrorMessage) ConnectChangeEvent(f func(e *core.QEvent)) {
	defer qt.Recovering("connect QErrorMessage::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QErrorMessage::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQErrorMessageChangeEvent
func callbackQErrorMessageChangeEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QErrorMessage::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QErrorMessage) ConnectDone(f func(a int)) {
	defer qt.Recovering("connect QErrorMessage::done")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "done", f)
	}
}

func (ptr *QErrorMessage) DisconnectDone() {
	defer qt.Recovering("disconnect QErrorMessage::done")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "done")
	}
}

//export callbackQErrorMessageDone
func callbackQErrorMessageDone(ptrName *C.char, a C.int) bool {
	defer qt.Recovering("callback QErrorMessage::done")

	if signal := qt.GetSignal(C.GoString(ptrName), "done"); signal != nil {
		signal.(func(int))(int(a))
		return true
	}
	return false

}

func QErrorMessage_QtHandler() *QErrorMessage {
	defer qt.Recovering("QErrorMessage::qtHandler")

	return NewQErrorMessageFromPointer(C.QErrorMessage_QErrorMessage_QtHandler())
}

func (ptr *QErrorMessage) ShowMessage(message string) {
	defer qt.Recovering("QErrorMessage::showMessage")

	if ptr.Pointer() != nil {
		C.QErrorMessage_ShowMessage(ptr.Pointer(), C.CString(message))
	}
}

func (ptr *QErrorMessage) ShowMessage2(message string, ty string) {
	defer qt.Recovering("QErrorMessage::showMessage")

	if ptr.Pointer() != nil {
		C.QErrorMessage_ShowMessage2(ptr.Pointer(), C.CString(message), C.CString(ty))
	}
}

func (ptr *QErrorMessage) DestroyQErrorMessage() {
	defer qt.Recovering("QErrorMessage::~QErrorMessage")

	if ptr.Pointer() != nil {
		C.QErrorMessage_DestroyQErrorMessage(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QErrorMessage) ConnectAccept(f func()) {
	defer qt.Recovering("connect QErrorMessage::accept")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "accept", f)
	}
}

func (ptr *QErrorMessage) DisconnectAccept() {
	defer qt.Recovering("disconnect QErrorMessage::accept")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "accept")
	}
}

//export callbackQErrorMessageAccept
func callbackQErrorMessageAccept(ptrName *C.char) bool {
	defer qt.Recovering("callback QErrorMessage::accept")

	if signal := qt.GetSignal(C.GoString(ptrName), "accept"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QErrorMessage) ConnectCloseEvent(f func(e *gui.QCloseEvent)) {
	defer qt.Recovering("connect QErrorMessage::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QErrorMessage::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQErrorMessageCloseEvent
func callbackQErrorMessageCloseEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QErrorMessage::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QErrorMessage) ConnectContextMenuEvent(f func(e *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QErrorMessage::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QErrorMessage::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQErrorMessageContextMenuEvent
func callbackQErrorMessageContextMenuEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QErrorMessage::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QErrorMessage) ConnectKeyPressEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QErrorMessage::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QErrorMessage::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQErrorMessageKeyPressEvent
func callbackQErrorMessageKeyPressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QErrorMessage::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QErrorMessage) ConnectOpen(f func()) {
	defer qt.Recovering("connect QErrorMessage::open")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "open", f)
	}
}

func (ptr *QErrorMessage) DisconnectOpen() {
	defer qt.Recovering("disconnect QErrorMessage::open")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "open")
	}
}

//export callbackQErrorMessageOpen
func callbackQErrorMessageOpen(ptrName *C.char) bool {
	defer qt.Recovering("callback QErrorMessage::open")

	if signal := qt.GetSignal(C.GoString(ptrName), "open"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QErrorMessage) ConnectReject(f func()) {
	defer qt.Recovering("connect QErrorMessage::reject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "reject", f)
	}
}

func (ptr *QErrorMessage) DisconnectReject() {
	defer qt.Recovering("disconnect QErrorMessage::reject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "reject")
	}
}

//export callbackQErrorMessageReject
func callbackQErrorMessageReject(ptrName *C.char) bool {
	defer qt.Recovering("callback QErrorMessage::reject")

	if signal := qt.GetSignal(C.GoString(ptrName), "reject"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QErrorMessage) ConnectResizeEvent(f func(v *gui.QResizeEvent)) {
	defer qt.Recovering("connect QErrorMessage::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QErrorMessage::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQErrorMessageResizeEvent
func callbackQErrorMessageResizeEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QErrorMessage::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(v))
		return true
	}
	return false

}

func (ptr *QErrorMessage) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QErrorMessage::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QErrorMessage) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QErrorMessage::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQErrorMessageSetVisible
func callbackQErrorMessageSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QErrorMessage::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QErrorMessage) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QErrorMessage::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QErrorMessage::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQErrorMessageShowEvent
func callbackQErrorMessageShowEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QErrorMessage::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QErrorMessage) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QErrorMessage::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QErrorMessage::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQErrorMessageActionEvent
func callbackQErrorMessageActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QErrorMessage::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QErrorMessage) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QErrorMessage::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QErrorMessage::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQErrorMessageDragEnterEvent
func callbackQErrorMessageDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QErrorMessage::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QErrorMessage) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QErrorMessage::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QErrorMessage::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQErrorMessageDragLeaveEvent
func callbackQErrorMessageDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QErrorMessage::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QErrorMessage) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QErrorMessage::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QErrorMessage::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQErrorMessageDragMoveEvent
func callbackQErrorMessageDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QErrorMessage::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QErrorMessage) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QErrorMessage::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QErrorMessage::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQErrorMessageDropEvent
func callbackQErrorMessageDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QErrorMessage::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QErrorMessage) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QErrorMessage::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QErrorMessage::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQErrorMessageEnterEvent
func callbackQErrorMessageEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QErrorMessage::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QErrorMessage) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QErrorMessage::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QErrorMessage::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQErrorMessageFocusOutEvent
func callbackQErrorMessageFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QErrorMessage::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QErrorMessage) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QErrorMessage::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QErrorMessage::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQErrorMessageHideEvent
func callbackQErrorMessageHideEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QErrorMessage::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QErrorMessage) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QErrorMessage::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QErrorMessage::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQErrorMessageLeaveEvent
func callbackQErrorMessageLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QErrorMessage::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QErrorMessage) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QErrorMessage::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QErrorMessage::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQErrorMessageMoveEvent
func callbackQErrorMessageMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QErrorMessage::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QErrorMessage) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QErrorMessage::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QErrorMessage::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQErrorMessagePaintEvent
func callbackQErrorMessagePaintEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QErrorMessage::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QErrorMessage) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QErrorMessage::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QErrorMessage) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QErrorMessage::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQErrorMessageInitPainter
func callbackQErrorMessageInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QErrorMessage::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

}

func (ptr *QErrorMessage) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QErrorMessage::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QErrorMessage::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQErrorMessageInputMethodEvent
func callbackQErrorMessageInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QErrorMessage::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QErrorMessage) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QErrorMessage::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QErrorMessage::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQErrorMessageKeyReleaseEvent
func callbackQErrorMessageKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QErrorMessage::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QErrorMessage) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QErrorMessage::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QErrorMessage::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQErrorMessageMouseDoubleClickEvent
func callbackQErrorMessageMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QErrorMessage::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QErrorMessage) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QErrorMessage::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QErrorMessage::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQErrorMessageMouseMoveEvent
func callbackQErrorMessageMouseMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QErrorMessage::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QErrorMessage) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QErrorMessage::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QErrorMessage::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQErrorMessageMousePressEvent
func callbackQErrorMessageMousePressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QErrorMessage::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QErrorMessage) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QErrorMessage::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QErrorMessage::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQErrorMessageMouseReleaseEvent
func callbackQErrorMessageMouseReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QErrorMessage::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QErrorMessage) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QErrorMessage::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QErrorMessage::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQErrorMessageTabletEvent
func callbackQErrorMessageTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QErrorMessage::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QErrorMessage) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QErrorMessage::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QErrorMessage::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQErrorMessageWheelEvent
func callbackQErrorMessageWheelEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QErrorMessage::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QErrorMessage) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QErrorMessage::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QErrorMessage::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQErrorMessageTimerEvent
func callbackQErrorMessageTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QErrorMessage::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QErrorMessage) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QErrorMessage::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QErrorMessage::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQErrorMessageChildEvent
func callbackQErrorMessageChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QErrorMessage::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QErrorMessage) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QErrorMessage::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QErrorMessage) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QErrorMessage::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQErrorMessageCustomEvent
func callbackQErrorMessageCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QErrorMessage::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
