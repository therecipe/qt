package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QToolButton struct {
	QAbstractButton
}

type QToolButton_ITF interface {
	QAbstractButton_ITF
	QToolButton_PTR() *QToolButton
}

func PointerFromQToolButton(ptr QToolButton_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QToolButton_PTR().Pointer()
	}
	return nil
}

func NewQToolButtonFromPointer(ptr unsafe.Pointer) *QToolButton {
	var n = new(QToolButton)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QToolButton_") {
		n.SetObjectName("QToolButton_" + qt.Identifier())
	}
	return n
}

func (ptr *QToolButton) QToolButton_PTR() *QToolButton {
	return ptr
}

//QToolButton::ToolButtonPopupMode
type QToolButton__ToolButtonPopupMode int64

const (
	QToolButton__DelayedPopup    = QToolButton__ToolButtonPopupMode(0)
	QToolButton__MenuButtonPopup = QToolButton__ToolButtonPopupMode(1)
	QToolButton__InstantPopup    = QToolButton__ToolButtonPopupMode(2)
)

func (ptr *QToolButton) ArrowType() core.Qt__ArrowType {
	defer qt.Recovering("QToolButton::arrowType")

	if ptr.Pointer() != nil {
		return core.Qt__ArrowType(C.QToolButton_ArrowType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QToolButton) AutoRaise() bool {
	defer qt.Recovering("QToolButton::autoRaise")

	if ptr.Pointer() != nil {
		return C.QToolButton_AutoRaise(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QToolButton) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QToolButton::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QToolButton) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QToolButton::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQToolButtonPaintEvent
func callbackQToolButtonPaintEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolButton::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QToolButton) PopupMode() QToolButton__ToolButtonPopupMode {
	defer qt.Recovering("QToolButton::popupMode")

	if ptr.Pointer() != nil {
		return QToolButton__ToolButtonPopupMode(C.QToolButton_PopupMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QToolButton) SetArrowType(ty core.Qt__ArrowType) {
	defer qt.Recovering("QToolButton::setArrowType")

	if ptr.Pointer() != nil {
		C.QToolButton_SetArrowType(ptr.Pointer(), C.int(ty))
	}
}

func (ptr *QToolButton) SetAutoRaise(enable bool) {
	defer qt.Recovering("QToolButton::setAutoRaise")

	if ptr.Pointer() != nil {
		C.QToolButton_SetAutoRaise(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QToolButton) SetPopupMode(mode QToolButton__ToolButtonPopupMode) {
	defer qt.Recovering("QToolButton::setPopupMode")

	if ptr.Pointer() != nil {
		C.QToolButton_SetPopupMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QToolButton) SetToolButtonStyle(style core.Qt__ToolButtonStyle) {
	defer qt.Recovering("QToolButton::setToolButtonStyle")

	if ptr.Pointer() != nil {
		C.QToolButton_SetToolButtonStyle(ptr.Pointer(), C.int(style))
	}
}

func (ptr *QToolButton) ToolButtonStyle() core.Qt__ToolButtonStyle {
	defer qt.Recovering("QToolButton::toolButtonStyle")

	if ptr.Pointer() != nil {
		return core.Qt__ToolButtonStyle(C.QToolButton_ToolButtonStyle(ptr.Pointer()))
	}
	return 0
}

func NewQToolButton(parent QWidget_ITF) *QToolButton {
	defer qt.Recovering("QToolButton::QToolButton")

	return NewQToolButtonFromPointer(C.QToolButton_NewQToolButton(PointerFromQWidget(parent)))
}

func (ptr *QToolButton) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QToolButton::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QToolButton) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QToolButton::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQToolButtonActionEvent
func callbackQToolButtonActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolButton::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QToolButton) ConnectChangeEvent(f func(e *core.QEvent)) {
	defer qt.Recovering("connect QToolButton::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QToolButton) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QToolButton::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQToolButtonChangeEvent
func callbackQToolButtonChangeEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolButton::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QToolButton) ConnectEnterEvent(f func(e *core.QEvent)) {
	defer qt.Recovering("connect QToolButton::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QToolButton) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QToolButton::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQToolButtonEnterEvent
func callbackQToolButtonEnterEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolButton::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QToolButton) ConnectLeaveEvent(f func(e *core.QEvent)) {
	defer qt.Recovering("connect QToolButton::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QToolButton) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QToolButton::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQToolButtonLeaveEvent
func callbackQToolButtonLeaveEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolButton::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QToolButton) Menu() *QMenu {
	defer qt.Recovering("QToolButton::menu")

	if ptr.Pointer() != nil {
		return NewQMenuFromPointer(C.QToolButton_Menu(ptr.Pointer()))
	}
	return nil
}

func (ptr *QToolButton) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QToolButton::minimumSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QToolButton_MinimumSizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QToolButton) ConnectMousePressEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QToolButton::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QToolButton) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QToolButton::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQToolButtonMousePressEvent
func callbackQToolButtonMousePressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolButton::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QToolButton) ConnectMouseReleaseEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QToolButton::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QToolButton) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QToolButton::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQToolButtonMouseReleaseEvent
func callbackQToolButtonMouseReleaseEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolButton::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QToolButton) ConnectNextCheckState(f func()) {
	defer qt.Recovering("connect QToolButton::nextCheckState")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "nextCheckState", f)
	}
}

func (ptr *QToolButton) DisconnectNextCheckState() {
	defer qt.Recovering("disconnect QToolButton::nextCheckState")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "nextCheckState")
	}
}

//export callbackQToolButtonNextCheckState
func callbackQToolButtonNextCheckState(ptrName *C.char) bool {
	defer qt.Recovering("callback QToolButton::nextCheckState")

	if signal := qt.GetSignal(C.GoString(ptrName), "nextCheckState"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QToolButton) SetMenu(menu QMenu_ITF) {
	defer qt.Recovering("QToolButton::setMenu")

	if ptr.Pointer() != nil {
		C.QToolButton_SetMenu(ptr.Pointer(), PointerFromQMenu(menu))
	}
}

func (ptr *QToolButton) ShowMenu() {
	defer qt.Recovering("QToolButton::showMenu")

	if ptr.Pointer() != nil {
		C.QToolButton_ShowMenu(ptr.Pointer())
	}
}

func (ptr *QToolButton) SizeHint() *core.QSize {
	defer qt.Recovering("QToolButton::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QToolButton_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QToolButton) ConnectTimerEvent(f func(e *core.QTimerEvent)) {
	defer qt.Recovering("connect QToolButton::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QToolButton) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QToolButton::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQToolButtonTimerEvent
func callbackQToolButtonTimerEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolButton::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QToolButton) ConnectTriggered(f func(action *QAction)) {
	defer qt.Recovering("connect QToolButton::triggered")

	if ptr.Pointer() != nil {
		C.QToolButton_ConnectTriggered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "triggered", f)
	}
}

func (ptr *QToolButton) DisconnectTriggered() {
	defer qt.Recovering("disconnect QToolButton::triggered")

	if ptr.Pointer() != nil {
		C.QToolButton_DisconnectTriggered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "triggered")
	}
}

//export callbackQToolButtonTriggered
func callbackQToolButtonTriggered(ptrName *C.char, action unsafe.Pointer) {
	defer qt.Recovering("callback QToolButton::triggered")

	if signal := qt.GetSignal(C.GoString(ptrName), "triggered"); signal != nil {
		signal.(func(*QAction))(NewQActionFromPointer(action))
	}

}

func (ptr *QToolButton) DestroyQToolButton() {
	defer qt.Recovering("QToolButton::~QToolButton")

	if ptr.Pointer() != nil {
		C.QToolButton_DestroyQToolButton(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QToolButton) ConnectCheckStateSet(f func()) {
	defer qt.Recovering("connect QToolButton::checkStateSet")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "checkStateSet", f)
	}
}

func (ptr *QToolButton) DisconnectCheckStateSet() {
	defer qt.Recovering("disconnect QToolButton::checkStateSet")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "checkStateSet")
	}
}

//export callbackQToolButtonCheckStateSet
func callbackQToolButtonCheckStateSet(ptrName *C.char) bool {
	defer qt.Recovering("callback QToolButton::checkStateSet")

	if signal := qt.GetSignal(C.GoString(ptrName), "checkStateSet"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QToolButton) ConnectFocusInEvent(f func(e *gui.QFocusEvent)) {
	defer qt.Recovering("connect QToolButton::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QToolButton) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QToolButton::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQToolButtonFocusInEvent
func callbackQToolButtonFocusInEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolButton::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QToolButton) ConnectFocusOutEvent(f func(e *gui.QFocusEvent)) {
	defer qt.Recovering("connect QToolButton::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QToolButton) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QToolButton::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQToolButtonFocusOutEvent
func callbackQToolButtonFocusOutEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolButton::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QToolButton) ConnectKeyPressEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QToolButton::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QToolButton) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QToolButton::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQToolButtonKeyPressEvent
func callbackQToolButtonKeyPressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolButton::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QToolButton) ConnectKeyReleaseEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QToolButton::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QToolButton) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QToolButton::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQToolButtonKeyReleaseEvent
func callbackQToolButtonKeyReleaseEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolButton::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QToolButton) ConnectMouseMoveEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QToolButton::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QToolButton) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QToolButton::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQToolButtonMouseMoveEvent
func callbackQToolButtonMouseMoveEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolButton::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QToolButton) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QToolButton::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QToolButton) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QToolButton::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQToolButtonDragEnterEvent
func callbackQToolButtonDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolButton::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QToolButton) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QToolButton::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QToolButton) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QToolButton::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQToolButtonDragLeaveEvent
func callbackQToolButtonDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolButton::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QToolButton) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QToolButton::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QToolButton) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QToolButton::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQToolButtonDragMoveEvent
func callbackQToolButtonDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolButton::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QToolButton) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QToolButton::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QToolButton) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QToolButton::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQToolButtonDropEvent
func callbackQToolButtonDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolButton::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QToolButton) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QToolButton::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QToolButton) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QToolButton::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQToolButtonHideEvent
func callbackQToolButtonHideEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolButton::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QToolButton) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QToolButton::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QToolButton) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QToolButton::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQToolButtonMoveEvent
func callbackQToolButtonMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolButton::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QToolButton) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QToolButton::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QToolButton) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QToolButton::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQToolButtonSetVisible
func callbackQToolButtonSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QToolButton::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QToolButton) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QToolButton::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QToolButton) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QToolButton::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQToolButtonShowEvent
func callbackQToolButtonShowEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolButton::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QToolButton) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QToolButton::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QToolButton) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QToolButton::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQToolButtonCloseEvent
func callbackQToolButtonCloseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolButton::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QToolButton) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QToolButton::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QToolButton) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QToolButton::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQToolButtonContextMenuEvent
func callbackQToolButtonContextMenuEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolButton::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QToolButton) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QToolButton::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QToolButton) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QToolButton::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQToolButtonInitPainter
func callbackQToolButtonInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolButton::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

}

func (ptr *QToolButton) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QToolButton::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QToolButton) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QToolButton::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQToolButtonInputMethodEvent
func callbackQToolButtonInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolButton::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QToolButton) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QToolButton::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QToolButton) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QToolButton::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQToolButtonMouseDoubleClickEvent
func callbackQToolButtonMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolButton::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QToolButton) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QToolButton::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QToolButton) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QToolButton::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQToolButtonResizeEvent
func callbackQToolButtonResizeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolButton::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QToolButton) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QToolButton::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QToolButton) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QToolButton::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQToolButtonTabletEvent
func callbackQToolButtonTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolButton::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QToolButton) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QToolButton::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QToolButton) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QToolButton::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQToolButtonWheelEvent
func callbackQToolButtonWheelEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolButton::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QToolButton) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QToolButton::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QToolButton) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QToolButton::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQToolButtonChildEvent
func callbackQToolButtonChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolButton::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QToolButton) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QToolButton::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QToolButton) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QToolButton::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQToolButtonCustomEvent
func callbackQToolButtonCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolButton::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
