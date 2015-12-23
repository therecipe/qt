package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QCommandLinkButton struct {
	QPushButton
}

type QCommandLinkButton_ITF interface {
	QPushButton_ITF
	QCommandLinkButton_PTR() *QCommandLinkButton
}

func PointerFromQCommandLinkButton(ptr QCommandLinkButton_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCommandLinkButton_PTR().Pointer()
	}
	return nil
}

func NewQCommandLinkButtonFromPointer(ptr unsafe.Pointer) *QCommandLinkButton {
	var n = new(QCommandLinkButton)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QCommandLinkButton_") {
		n.SetObjectName("QCommandLinkButton_" + qt.Identifier())
	}
	return n
}

func (ptr *QCommandLinkButton) QCommandLinkButton_PTR() *QCommandLinkButton {
	return ptr
}

func (ptr *QCommandLinkButton) Description() string {
	defer qt.Recovering("QCommandLinkButton::description")

	if ptr.Pointer() != nil {
		return C.GoString(C.QCommandLinkButton_Description(ptr.Pointer()))
	}
	return ""
}

func (ptr *QCommandLinkButton) SetDescription(description string) {
	defer qt.Recovering("QCommandLinkButton::setDescription")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_SetDescription(ptr.Pointer(), C.CString(description))
	}
}

func NewQCommandLinkButton(parent QWidget_ITF) *QCommandLinkButton {
	defer qt.Recovering("QCommandLinkButton::QCommandLinkButton")

	return NewQCommandLinkButtonFromPointer(C.QCommandLinkButton_NewQCommandLinkButton(PointerFromQWidget(parent)))
}

func NewQCommandLinkButton2(text string, parent QWidget_ITF) *QCommandLinkButton {
	defer qt.Recovering("QCommandLinkButton::QCommandLinkButton")

	return NewQCommandLinkButtonFromPointer(C.QCommandLinkButton_NewQCommandLinkButton2(C.CString(text), PointerFromQWidget(parent)))
}

func NewQCommandLinkButton3(text string, description string, parent QWidget_ITF) *QCommandLinkButton {
	defer qt.Recovering("QCommandLinkButton::QCommandLinkButton")

	return NewQCommandLinkButtonFromPointer(C.QCommandLinkButton_NewQCommandLinkButton3(C.CString(text), C.CString(description), PointerFromQWidget(parent)))
}

func (ptr *QCommandLinkButton) ConnectPaintEvent(f func(v *gui.QPaintEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQCommandLinkButtonPaintEvent
func callbackQCommandLinkButtonPaintEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QCommandLinkButton::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(v))
		return true
	}
	return false

}

func (ptr *QCommandLinkButton) DestroyQCommandLinkButton() {
	defer qt.Recovering("QCommandLinkButton::~QCommandLinkButton")

	if ptr.Pointer() != nil {
		C.QCommandLinkButton_DestroyQCommandLinkButton(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCommandLinkButton) ConnectFocusInEvent(f func(e *gui.QFocusEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQCommandLinkButtonFocusInEvent
func callbackQCommandLinkButtonFocusInEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QCommandLinkButton::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QCommandLinkButton) ConnectFocusOutEvent(f func(e *gui.QFocusEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQCommandLinkButtonFocusOutEvent
func callbackQCommandLinkButtonFocusOutEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QCommandLinkButton::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QCommandLinkButton) ConnectKeyPressEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQCommandLinkButtonKeyPressEvent
func callbackQCommandLinkButtonKeyPressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QCommandLinkButton::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QCommandLinkButton) ConnectChangeEvent(f func(e *core.QEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQCommandLinkButtonChangeEvent
func callbackQCommandLinkButtonChangeEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QCommandLinkButton::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QCommandLinkButton) ConnectCheckStateSet(f func()) {
	defer qt.Recovering("connect QCommandLinkButton::checkStateSet")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "checkStateSet", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectCheckStateSet() {
	defer qt.Recovering("disconnect QCommandLinkButton::checkStateSet")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "checkStateSet")
	}
}

//export callbackQCommandLinkButtonCheckStateSet
func callbackQCommandLinkButtonCheckStateSet(ptrName *C.char) bool {
	defer qt.Recovering("callback QCommandLinkButton::checkStateSet")

	if signal := qt.GetSignal(C.GoString(ptrName), "checkStateSet"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QCommandLinkButton) ConnectKeyReleaseEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQCommandLinkButtonKeyReleaseEvent
func callbackQCommandLinkButtonKeyReleaseEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QCommandLinkButton::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QCommandLinkButton) ConnectMouseMoveEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQCommandLinkButtonMouseMoveEvent
func callbackQCommandLinkButtonMouseMoveEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QCommandLinkButton::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QCommandLinkButton) ConnectMousePressEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQCommandLinkButtonMousePressEvent
func callbackQCommandLinkButtonMousePressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QCommandLinkButton::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QCommandLinkButton) ConnectMouseReleaseEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQCommandLinkButtonMouseReleaseEvent
func callbackQCommandLinkButtonMouseReleaseEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QCommandLinkButton::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QCommandLinkButton) ConnectNextCheckState(f func()) {
	defer qt.Recovering("connect QCommandLinkButton::nextCheckState")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "nextCheckState", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectNextCheckState() {
	defer qt.Recovering("disconnect QCommandLinkButton::nextCheckState")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "nextCheckState")
	}
}

//export callbackQCommandLinkButtonNextCheckState
func callbackQCommandLinkButtonNextCheckState(ptrName *C.char) bool {
	defer qt.Recovering("callback QCommandLinkButton::nextCheckState")

	if signal := qt.GetSignal(C.GoString(ptrName), "nextCheckState"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QCommandLinkButton) ConnectTimerEvent(f func(e *core.QTimerEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQCommandLinkButtonTimerEvent
func callbackQCommandLinkButtonTimerEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QCommandLinkButton::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QCommandLinkButton) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQCommandLinkButtonActionEvent
func callbackQCommandLinkButtonActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCommandLinkButton::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCommandLinkButton) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQCommandLinkButtonDragEnterEvent
func callbackQCommandLinkButtonDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCommandLinkButton::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCommandLinkButton) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQCommandLinkButtonDragLeaveEvent
func callbackQCommandLinkButtonDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCommandLinkButton::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCommandLinkButton) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQCommandLinkButtonDragMoveEvent
func callbackQCommandLinkButtonDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCommandLinkButton::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCommandLinkButton) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQCommandLinkButtonDropEvent
func callbackQCommandLinkButtonDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCommandLinkButton::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCommandLinkButton) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQCommandLinkButtonEnterEvent
func callbackQCommandLinkButtonEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCommandLinkButton::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCommandLinkButton) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQCommandLinkButtonHideEvent
func callbackQCommandLinkButtonHideEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCommandLinkButton::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCommandLinkButton) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQCommandLinkButtonLeaveEvent
func callbackQCommandLinkButtonLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCommandLinkButton::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCommandLinkButton) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQCommandLinkButtonMoveEvent
func callbackQCommandLinkButtonMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCommandLinkButton::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCommandLinkButton) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QCommandLinkButton::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QCommandLinkButton::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQCommandLinkButtonSetVisible
func callbackQCommandLinkButtonSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QCommandLinkButton::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QCommandLinkButton) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQCommandLinkButtonShowEvent
func callbackQCommandLinkButtonShowEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCommandLinkButton::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCommandLinkButton) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQCommandLinkButtonCloseEvent
func callbackQCommandLinkButtonCloseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCommandLinkButton::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCommandLinkButton) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQCommandLinkButtonContextMenuEvent
func callbackQCommandLinkButtonContextMenuEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCommandLinkButton::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCommandLinkButton) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QCommandLinkButton::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QCommandLinkButton::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQCommandLinkButtonInitPainter
func callbackQCommandLinkButtonInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QCommandLinkButton::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

}

func (ptr *QCommandLinkButton) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQCommandLinkButtonInputMethodEvent
func callbackQCommandLinkButtonInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCommandLinkButton::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCommandLinkButton) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQCommandLinkButtonMouseDoubleClickEvent
func callbackQCommandLinkButtonMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCommandLinkButton::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCommandLinkButton) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQCommandLinkButtonResizeEvent
func callbackQCommandLinkButtonResizeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCommandLinkButton::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCommandLinkButton) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQCommandLinkButtonTabletEvent
func callbackQCommandLinkButtonTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCommandLinkButton::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCommandLinkButton) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQCommandLinkButtonWheelEvent
func callbackQCommandLinkButtonWheelEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCommandLinkButton::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCommandLinkButton) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQCommandLinkButtonChildEvent
func callbackQCommandLinkButtonChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCommandLinkButton::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCommandLinkButton) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCommandLinkButton::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCommandLinkButton) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCommandLinkButton::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQCommandLinkButtonCustomEvent
func callbackQCommandLinkButtonCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCommandLinkButton::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
