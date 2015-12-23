package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QMdiSubWindow struct {
	QWidget
}

type QMdiSubWindow_ITF interface {
	QWidget_ITF
	QMdiSubWindow_PTR() *QMdiSubWindow
}

func PointerFromQMdiSubWindow(ptr QMdiSubWindow_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMdiSubWindow_PTR().Pointer()
	}
	return nil
}

func NewQMdiSubWindowFromPointer(ptr unsafe.Pointer) *QMdiSubWindow {
	var n = new(QMdiSubWindow)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMdiSubWindow_") {
		n.SetObjectName("QMdiSubWindow_" + qt.Identifier())
	}
	return n
}

func (ptr *QMdiSubWindow) QMdiSubWindow_PTR() *QMdiSubWindow {
	return ptr
}

//QMdiSubWindow::SubWindowOption
type QMdiSubWindow__SubWindowOption int64

const (
	QMdiSubWindow__AllowOutsideAreaHorizontally = QMdiSubWindow__SubWindowOption(0x1)
	QMdiSubWindow__AllowOutsideAreaVertically   = QMdiSubWindow__SubWindowOption(0x2)
	QMdiSubWindow__RubberBandResize             = QMdiSubWindow__SubWindowOption(0x4)
	QMdiSubWindow__RubberBandMove               = QMdiSubWindow__SubWindowOption(0x8)
)

func (ptr *QMdiSubWindow) KeyboardPageStep() int {
	defer qt.Recovering("QMdiSubWindow::keyboardPageStep")

	if ptr.Pointer() != nil {
		return int(C.QMdiSubWindow_KeyboardPageStep(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMdiSubWindow) KeyboardSingleStep() int {
	defer qt.Recovering("QMdiSubWindow::keyboardSingleStep")

	if ptr.Pointer() != nil {
		return int(C.QMdiSubWindow_KeyboardSingleStep(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMdiSubWindow) SetKeyboardPageStep(step int) {
	defer qt.Recovering("QMdiSubWindow::setKeyboardPageStep")

	if ptr.Pointer() != nil {
		C.QMdiSubWindow_SetKeyboardPageStep(ptr.Pointer(), C.int(step))
	}
}

func (ptr *QMdiSubWindow) SetKeyboardSingleStep(step int) {
	defer qt.Recovering("QMdiSubWindow::setKeyboardSingleStep")

	if ptr.Pointer() != nil {
		C.QMdiSubWindow_SetKeyboardSingleStep(ptr.Pointer(), C.int(step))
	}
}

func NewQMdiSubWindow(parent QWidget_ITF, flags core.Qt__WindowType) *QMdiSubWindow {
	defer qt.Recovering("QMdiSubWindow::QMdiSubWindow")

	return NewQMdiSubWindowFromPointer(C.QMdiSubWindow_NewQMdiSubWindow(PointerFromQWidget(parent), C.int(flags)))
}

func (ptr *QMdiSubWindow) ConnectAboutToActivate(f func()) {
	defer qt.Recovering("connect QMdiSubWindow::aboutToActivate")

	if ptr.Pointer() != nil {
		C.QMdiSubWindow_ConnectAboutToActivate(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "aboutToActivate", f)
	}
}

func (ptr *QMdiSubWindow) DisconnectAboutToActivate() {
	defer qt.Recovering("disconnect QMdiSubWindow::aboutToActivate")

	if ptr.Pointer() != nil {
		C.QMdiSubWindow_DisconnectAboutToActivate(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "aboutToActivate")
	}
}

//export callbackQMdiSubWindowAboutToActivate
func callbackQMdiSubWindowAboutToActivate(ptrName *C.char) {
	defer qt.Recovering("callback QMdiSubWindow::aboutToActivate")

	if signal := qt.GetSignal(C.GoString(ptrName), "aboutToActivate"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QMdiSubWindow) ConnectChangeEvent(f func(changeEvent *core.QEvent)) {
	defer qt.Recovering("connect QMdiSubWindow::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QMdiSubWindow) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QMdiSubWindow::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQMdiSubWindowChangeEvent
func callbackQMdiSubWindowChangeEvent(ptrName *C.char, changeEvent unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiSubWindow::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(changeEvent))
		return true
	}
	return false

}

func (ptr *QMdiSubWindow) ConnectChildEvent(f func(childEvent *core.QChildEvent)) {
	defer qt.Recovering("connect QMdiSubWindow::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMdiSubWindow) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMdiSubWindow::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMdiSubWindowChildEvent
func callbackQMdiSubWindowChildEvent(ptrName *C.char, childEvent unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiSubWindow::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(childEvent))
		return true
	}
	return false

}

func (ptr *QMdiSubWindow) ConnectCloseEvent(f func(closeEvent *gui.QCloseEvent)) {
	defer qt.Recovering("connect QMdiSubWindow::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QMdiSubWindow) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QMdiSubWindow::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQMdiSubWindowCloseEvent
func callbackQMdiSubWindowCloseEvent(ptrName *C.char, closeEvent unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiSubWindow::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(closeEvent))
		return true
	}
	return false

}

func (ptr *QMdiSubWindow) ConnectContextMenuEvent(f func(contextMenuEvent *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QMdiSubWindow::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QMdiSubWindow) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QMdiSubWindow::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQMdiSubWindowContextMenuEvent
func callbackQMdiSubWindowContextMenuEvent(ptrName *C.char, contextMenuEvent unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiSubWindow::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(contextMenuEvent))
		return true
	}
	return false

}

func (ptr *QMdiSubWindow) ConnectFocusInEvent(f func(focusInEvent *gui.QFocusEvent)) {
	defer qt.Recovering("connect QMdiSubWindow::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QMdiSubWindow) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QMdiSubWindow::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQMdiSubWindowFocusInEvent
func callbackQMdiSubWindowFocusInEvent(ptrName *C.char, focusInEvent unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiSubWindow::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(focusInEvent))
		return true
	}
	return false

}

func (ptr *QMdiSubWindow) ConnectFocusOutEvent(f func(focusOutEvent *gui.QFocusEvent)) {
	defer qt.Recovering("connect QMdiSubWindow::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QMdiSubWindow) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QMdiSubWindow::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQMdiSubWindowFocusOutEvent
func callbackQMdiSubWindowFocusOutEvent(ptrName *C.char, focusOutEvent unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiSubWindow::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(focusOutEvent))
		return true
	}
	return false

}

func (ptr *QMdiSubWindow) ConnectHideEvent(f func(hideEvent *gui.QHideEvent)) {
	defer qt.Recovering("connect QMdiSubWindow::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QMdiSubWindow) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QMdiSubWindow::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQMdiSubWindowHideEvent
func callbackQMdiSubWindowHideEvent(ptrName *C.char, hideEvent unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiSubWindow::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(hideEvent))
		return true
	}
	return false

}

func (ptr *QMdiSubWindow) IsShaded() bool {
	defer qt.Recovering("QMdiSubWindow::isShaded")

	if ptr.Pointer() != nil {
		return C.QMdiSubWindow_IsShaded(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMdiSubWindow) ConnectKeyPressEvent(f func(keyEvent *gui.QKeyEvent)) {
	defer qt.Recovering("connect QMdiSubWindow::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QMdiSubWindow) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QMdiSubWindow::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQMdiSubWindowKeyPressEvent
func callbackQMdiSubWindowKeyPressEvent(ptrName *C.char, keyEvent unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiSubWindow::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(keyEvent))
		return true
	}
	return false

}

func (ptr *QMdiSubWindow) ConnectLeaveEvent(f func(leaveEvent *core.QEvent)) {
	defer qt.Recovering("connect QMdiSubWindow::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QMdiSubWindow) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QMdiSubWindow::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQMdiSubWindowLeaveEvent
func callbackQMdiSubWindowLeaveEvent(ptrName *C.char, leaveEvent unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiSubWindow::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(leaveEvent))
		return true
	}
	return false

}

func (ptr *QMdiSubWindow) MdiArea() *QMdiArea {
	defer qt.Recovering("QMdiSubWindow::mdiArea")

	if ptr.Pointer() != nil {
		return NewQMdiAreaFromPointer(C.QMdiSubWindow_MdiArea(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMdiSubWindow) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QMdiSubWindow::minimumSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QMdiSubWindow_MinimumSizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMdiSubWindow) ConnectMouseDoubleClickEvent(f func(mouseEvent *gui.QMouseEvent)) {
	defer qt.Recovering("connect QMdiSubWindow::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QMdiSubWindow) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QMdiSubWindow::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQMdiSubWindowMouseDoubleClickEvent
func callbackQMdiSubWindowMouseDoubleClickEvent(ptrName *C.char, mouseEvent unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiSubWindow::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(mouseEvent))
		return true
	}
	return false

}

func (ptr *QMdiSubWindow) ConnectMouseMoveEvent(f func(mouseEvent *gui.QMouseEvent)) {
	defer qt.Recovering("connect QMdiSubWindow::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QMdiSubWindow) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QMdiSubWindow::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQMdiSubWindowMouseMoveEvent
func callbackQMdiSubWindowMouseMoveEvent(ptrName *C.char, mouseEvent unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiSubWindow::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(mouseEvent))
		return true
	}
	return false

}

func (ptr *QMdiSubWindow) ConnectMousePressEvent(f func(mouseEvent *gui.QMouseEvent)) {
	defer qt.Recovering("connect QMdiSubWindow::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QMdiSubWindow) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QMdiSubWindow::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQMdiSubWindowMousePressEvent
func callbackQMdiSubWindowMousePressEvent(ptrName *C.char, mouseEvent unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiSubWindow::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(mouseEvent))
		return true
	}
	return false

}

func (ptr *QMdiSubWindow) ConnectMouseReleaseEvent(f func(mouseEvent *gui.QMouseEvent)) {
	defer qt.Recovering("connect QMdiSubWindow::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QMdiSubWindow) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QMdiSubWindow::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQMdiSubWindowMouseReleaseEvent
func callbackQMdiSubWindowMouseReleaseEvent(ptrName *C.char, mouseEvent unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiSubWindow::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(mouseEvent))
		return true
	}
	return false

}

func (ptr *QMdiSubWindow) ConnectMoveEvent(f func(moveEvent *gui.QMoveEvent)) {
	defer qt.Recovering("connect QMdiSubWindow::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QMdiSubWindow) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QMdiSubWindow::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQMdiSubWindowMoveEvent
func callbackQMdiSubWindowMoveEvent(ptrName *C.char, moveEvent unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiSubWindow::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(moveEvent))
		return true
	}
	return false

}

func (ptr *QMdiSubWindow) ConnectPaintEvent(f func(paintEvent *gui.QPaintEvent)) {
	defer qt.Recovering("connect QMdiSubWindow::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QMdiSubWindow) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QMdiSubWindow::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQMdiSubWindowPaintEvent
func callbackQMdiSubWindowPaintEvent(ptrName *C.char, paintEvent unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiSubWindow::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(paintEvent))
		return true
	}
	return false

}

func (ptr *QMdiSubWindow) ConnectResizeEvent(f func(resizeEvent *gui.QResizeEvent)) {
	defer qt.Recovering("connect QMdiSubWindow::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QMdiSubWindow) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QMdiSubWindow::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQMdiSubWindowResizeEvent
func callbackQMdiSubWindowResizeEvent(ptrName *C.char, resizeEvent unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiSubWindow::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(resizeEvent))
		return true
	}
	return false

}

func (ptr *QMdiSubWindow) SetOption(option QMdiSubWindow__SubWindowOption, on bool) {
	defer qt.Recovering("QMdiSubWindow::setOption")

	if ptr.Pointer() != nil {
		C.QMdiSubWindow_SetOption(ptr.Pointer(), C.int(option), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QMdiSubWindow) SetSystemMenu(systemMenu QMenu_ITF) {
	defer qt.Recovering("QMdiSubWindow::setSystemMenu")

	if ptr.Pointer() != nil {
		C.QMdiSubWindow_SetSystemMenu(ptr.Pointer(), PointerFromQMenu(systemMenu))
	}
}

func (ptr *QMdiSubWindow) SetWidget(widget QWidget_ITF) {
	defer qt.Recovering("QMdiSubWindow::setWidget")

	if ptr.Pointer() != nil {
		C.QMdiSubWindow_SetWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QMdiSubWindow) ConnectShowEvent(f func(showEvent *gui.QShowEvent)) {
	defer qt.Recovering("connect QMdiSubWindow::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QMdiSubWindow) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QMdiSubWindow::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQMdiSubWindowShowEvent
func callbackQMdiSubWindowShowEvent(ptrName *C.char, showEvent unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiSubWindow::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(showEvent))
		return true
	}
	return false

}

func (ptr *QMdiSubWindow) ShowShaded() {
	defer qt.Recovering("QMdiSubWindow::showShaded")

	if ptr.Pointer() != nil {
		C.QMdiSubWindow_ShowShaded(ptr.Pointer())
	}
}

func (ptr *QMdiSubWindow) ShowSystemMenu() {
	defer qt.Recovering("QMdiSubWindow::showSystemMenu")

	if ptr.Pointer() != nil {
		C.QMdiSubWindow_ShowSystemMenu(ptr.Pointer())
	}
}

func (ptr *QMdiSubWindow) SizeHint() *core.QSize {
	defer qt.Recovering("QMdiSubWindow::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QMdiSubWindow_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMdiSubWindow) SystemMenu() *QMenu {
	defer qt.Recovering("QMdiSubWindow::systemMenu")

	if ptr.Pointer() != nil {
		return NewQMenuFromPointer(C.QMdiSubWindow_SystemMenu(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMdiSubWindow) TestOption(option QMdiSubWindow__SubWindowOption) bool {
	defer qt.Recovering("QMdiSubWindow::testOption")

	if ptr.Pointer() != nil {
		return C.QMdiSubWindow_TestOption(ptr.Pointer(), C.int(option)) != 0
	}
	return false
}

func (ptr *QMdiSubWindow) ConnectTimerEvent(f func(timerEvent *core.QTimerEvent)) {
	defer qt.Recovering("connect QMdiSubWindow::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMdiSubWindow) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMdiSubWindow::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMdiSubWindowTimerEvent
func callbackQMdiSubWindowTimerEvent(ptrName *C.char, timerEvent unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiSubWindow::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(timerEvent))
		return true
	}
	return false

}

func (ptr *QMdiSubWindow) Widget() *QWidget {
	defer qt.Recovering("QMdiSubWindow::widget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QMdiSubWindow_Widget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMdiSubWindow) ConnectWindowStateChanged(f func(oldState core.Qt__WindowState, newState core.Qt__WindowState)) {
	defer qt.Recovering("connect QMdiSubWindow::windowStateChanged")

	if ptr.Pointer() != nil {
		C.QMdiSubWindow_ConnectWindowStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "windowStateChanged", f)
	}
}

func (ptr *QMdiSubWindow) DisconnectWindowStateChanged() {
	defer qt.Recovering("disconnect QMdiSubWindow::windowStateChanged")

	if ptr.Pointer() != nil {
		C.QMdiSubWindow_DisconnectWindowStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "windowStateChanged")
	}
}

//export callbackQMdiSubWindowWindowStateChanged
func callbackQMdiSubWindowWindowStateChanged(ptrName *C.char, oldState C.int, newState C.int) {
	defer qt.Recovering("callback QMdiSubWindow::windowStateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "windowStateChanged"); signal != nil {
		signal.(func(core.Qt__WindowState, core.Qt__WindowState))(core.Qt__WindowState(oldState), core.Qt__WindowState(newState))
	}

}

func (ptr *QMdiSubWindow) DestroyQMdiSubWindow() {
	defer qt.Recovering("QMdiSubWindow::~QMdiSubWindow")

	if ptr.Pointer() != nil {
		C.QMdiSubWindow_DestroyQMdiSubWindow(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMdiSubWindow) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QMdiSubWindow::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QMdiSubWindow) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QMdiSubWindow::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQMdiSubWindowActionEvent
func callbackQMdiSubWindowActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiSubWindow::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMdiSubWindow) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QMdiSubWindow::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QMdiSubWindow) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QMdiSubWindow::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQMdiSubWindowDragEnterEvent
func callbackQMdiSubWindowDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiSubWindow::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMdiSubWindow) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QMdiSubWindow::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QMdiSubWindow) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QMdiSubWindow::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQMdiSubWindowDragLeaveEvent
func callbackQMdiSubWindowDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiSubWindow::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMdiSubWindow) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QMdiSubWindow::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QMdiSubWindow) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QMdiSubWindow::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQMdiSubWindowDragMoveEvent
func callbackQMdiSubWindowDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiSubWindow::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMdiSubWindow) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QMdiSubWindow::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QMdiSubWindow) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QMdiSubWindow::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQMdiSubWindowDropEvent
func callbackQMdiSubWindowDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiSubWindow::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMdiSubWindow) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMdiSubWindow::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QMdiSubWindow) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QMdiSubWindow::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQMdiSubWindowEnterEvent
func callbackQMdiSubWindowEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiSubWindow::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMdiSubWindow) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QMdiSubWindow::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QMdiSubWindow) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QMdiSubWindow::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQMdiSubWindowSetVisible
func callbackQMdiSubWindowSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QMdiSubWindow::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QMdiSubWindow) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QMdiSubWindow::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QMdiSubWindow) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QMdiSubWindow::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQMdiSubWindowInitPainter
func callbackQMdiSubWindowInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiSubWindow::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

}

func (ptr *QMdiSubWindow) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QMdiSubWindow::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QMdiSubWindow) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QMdiSubWindow::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQMdiSubWindowInputMethodEvent
func callbackQMdiSubWindowInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiSubWindow::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMdiSubWindow) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QMdiSubWindow::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QMdiSubWindow) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QMdiSubWindow::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQMdiSubWindowKeyReleaseEvent
func callbackQMdiSubWindowKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiSubWindow::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMdiSubWindow) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QMdiSubWindow::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QMdiSubWindow) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QMdiSubWindow::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQMdiSubWindowTabletEvent
func callbackQMdiSubWindowTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiSubWindow::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMdiSubWindow) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QMdiSubWindow::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QMdiSubWindow) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QMdiSubWindow::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQMdiSubWindowWheelEvent
func callbackQMdiSubWindowWheelEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiSubWindow::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMdiSubWindow) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMdiSubWindow::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMdiSubWindow) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMdiSubWindow::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMdiSubWindowCustomEvent
func callbackQMdiSubWindowCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMdiSubWindow::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
