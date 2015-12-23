package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QStatusBar struct {
	QWidget
}

type QStatusBar_ITF interface {
	QWidget_ITF
	QStatusBar_PTR() *QStatusBar
}

func PointerFromQStatusBar(ptr QStatusBar_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStatusBar_PTR().Pointer()
	}
	return nil
}

func NewQStatusBarFromPointer(ptr unsafe.Pointer) *QStatusBar {
	var n = new(QStatusBar)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QStatusBar_") {
		n.SetObjectName("QStatusBar_" + qt.Identifier())
	}
	return n
}

func (ptr *QStatusBar) QStatusBar_PTR() *QStatusBar {
	return ptr
}

func (ptr *QStatusBar) IsSizeGripEnabled() bool {
	defer qt.Recovering("QStatusBar::isSizeGripEnabled")

	if ptr.Pointer() != nil {
		return C.QStatusBar_IsSizeGripEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStatusBar) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QStatusBar::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QStatusBar::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQStatusBarPaintEvent
func callbackQStatusBarPaintEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QStatusBar::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QStatusBar) SetSizeGripEnabled(v bool) {
	defer qt.Recovering("QStatusBar::setSizeGripEnabled")

	if ptr.Pointer() != nil {
		C.QStatusBar_SetSizeGripEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func NewQStatusBar(parent QWidget_ITF) *QStatusBar {
	defer qt.Recovering("QStatusBar::QStatusBar")

	return NewQStatusBarFromPointer(C.QStatusBar_NewQStatusBar(PointerFromQWidget(parent)))
}

func (ptr *QStatusBar) AddPermanentWidget(widget QWidget_ITF, stretch int) {
	defer qt.Recovering("QStatusBar::addPermanentWidget")

	if ptr.Pointer() != nil {
		C.QStatusBar_AddPermanentWidget(ptr.Pointer(), PointerFromQWidget(widget), C.int(stretch))
	}
}

func (ptr *QStatusBar) AddWidget(widget QWidget_ITF, stretch int) {
	defer qt.Recovering("QStatusBar::addWidget")

	if ptr.Pointer() != nil {
		C.QStatusBar_AddWidget(ptr.Pointer(), PointerFromQWidget(widget), C.int(stretch))
	}
}

func (ptr *QStatusBar) ClearMessage() {
	defer qt.Recovering("QStatusBar::clearMessage")

	if ptr.Pointer() != nil {
		C.QStatusBar_ClearMessage(ptr.Pointer())
	}
}

func (ptr *QStatusBar) CurrentMessage() string {
	defer qt.Recovering("QStatusBar::currentMessage")

	if ptr.Pointer() != nil {
		return C.GoString(C.QStatusBar_CurrentMessage(ptr.Pointer()))
	}
	return ""
}

func (ptr *QStatusBar) InsertPermanentWidget(index int, widget QWidget_ITF, stretch int) int {
	defer qt.Recovering("QStatusBar::insertPermanentWidget")

	if ptr.Pointer() != nil {
		return int(C.QStatusBar_InsertPermanentWidget(ptr.Pointer(), C.int(index), PointerFromQWidget(widget), C.int(stretch)))
	}
	return 0
}

func (ptr *QStatusBar) InsertWidget(index int, widget QWidget_ITF, stretch int) int {
	defer qt.Recovering("QStatusBar::insertWidget")

	if ptr.Pointer() != nil {
		return int(C.QStatusBar_InsertWidget(ptr.Pointer(), C.int(index), PointerFromQWidget(widget), C.int(stretch)))
	}
	return 0
}

func (ptr *QStatusBar) ConnectMessageChanged(f func(message string)) {
	defer qt.Recovering("connect QStatusBar::messageChanged")

	if ptr.Pointer() != nil {
		C.QStatusBar_ConnectMessageChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "messageChanged", f)
	}
}

func (ptr *QStatusBar) DisconnectMessageChanged() {
	defer qt.Recovering("disconnect QStatusBar::messageChanged")

	if ptr.Pointer() != nil {
		C.QStatusBar_DisconnectMessageChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "messageChanged")
	}
}

//export callbackQStatusBarMessageChanged
func callbackQStatusBarMessageChanged(ptrName *C.char, message *C.char) {
	defer qt.Recovering("callback QStatusBar::messageChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "messageChanged"); signal != nil {
		signal.(func(string))(C.GoString(message))
	}

}

func (ptr *QStatusBar) RemoveWidget(widget QWidget_ITF) {
	defer qt.Recovering("QStatusBar::removeWidget")

	if ptr.Pointer() != nil {
		C.QStatusBar_RemoveWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QStatusBar) ConnectResizeEvent(f func(e *gui.QResizeEvent)) {
	defer qt.Recovering("connect QStatusBar::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QStatusBar::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQStatusBarResizeEvent
func callbackQStatusBarResizeEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QStatusBar::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QStatusBar) ConnectShowEvent(f func(v *gui.QShowEvent)) {
	defer qt.Recovering("connect QStatusBar::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QStatusBar::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQStatusBarShowEvent
func callbackQStatusBarShowEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QStatusBar::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(v))
		return true
	}
	return false

}

func (ptr *QStatusBar) ShowMessage(message string, timeout int) {
	defer qt.Recovering("QStatusBar::showMessage")

	if ptr.Pointer() != nil {
		C.QStatusBar_ShowMessage(ptr.Pointer(), C.CString(message), C.int(timeout))
	}
}

func (ptr *QStatusBar) DestroyQStatusBar() {
	defer qt.Recovering("QStatusBar::~QStatusBar")

	if ptr.Pointer() != nil {
		C.QStatusBar_DestroyQStatusBar(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QStatusBar) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QStatusBar::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QStatusBar::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQStatusBarActionEvent
func callbackQStatusBarActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QStatusBar::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QStatusBar) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QStatusBar::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QStatusBar::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQStatusBarDragEnterEvent
func callbackQStatusBarDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QStatusBar::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QStatusBar) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QStatusBar::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QStatusBar::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQStatusBarDragLeaveEvent
func callbackQStatusBarDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QStatusBar::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QStatusBar) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QStatusBar::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QStatusBar::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQStatusBarDragMoveEvent
func callbackQStatusBarDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QStatusBar::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QStatusBar) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QStatusBar::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QStatusBar::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQStatusBarDropEvent
func callbackQStatusBarDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QStatusBar::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QStatusBar) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QStatusBar::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QStatusBar::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQStatusBarEnterEvent
func callbackQStatusBarEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QStatusBar::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QStatusBar) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QStatusBar::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QStatusBar::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQStatusBarFocusInEvent
func callbackQStatusBarFocusInEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QStatusBar::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QStatusBar) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QStatusBar::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QStatusBar::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQStatusBarFocusOutEvent
func callbackQStatusBarFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QStatusBar::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QStatusBar) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QStatusBar::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QStatusBar::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQStatusBarHideEvent
func callbackQStatusBarHideEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QStatusBar::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QStatusBar) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QStatusBar::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QStatusBar::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQStatusBarLeaveEvent
func callbackQStatusBarLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QStatusBar::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QStatusBar) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QStatusBar::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QStatusBar::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQStatusBarMoveEvent
func callbackQStatusBarMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QStatusBar::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QStatusBar) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QStatusBar::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QStatusBar) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QStatusBar::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQStatusBarSetVisible
func callbackQStatusBarSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QStatusBar::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QStatusBar) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QStatusBar::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QStatusBar::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQStatusBarChangeEvent
func callbackQStatusBarChangeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QStatusBar::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QStatusBar) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QStatusBar::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QStatusBar::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQStatusBarCloseEvent
func callbackQStatusBarCloseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QStatusBar::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QStatusBar) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QStatusBar::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QStatusBar::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQStatusBarContextMenuEvent
func callbackQStatusBarContextMenuEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QStatusBar::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QStatusBar) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QStatusBar::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QStatusBar) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QStatusBar::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQStatusBarInitPainter
func callbackQStatusBarInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QStatusBar::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

}

func (ptr *QStatusBar) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QStatusBar::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QStatusBar::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQStatusBarInputMethodEvent
func callbackQStatusBarInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QStatusBar::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QStatusBar) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QStatusBar::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QStatusBar::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQStatusBarKeyPressEvent
func callbackQStatusBarKeyPressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QStatusBar::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QStatusBar) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QStatusBar::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QStatusBar::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQStatusBarKeyReleaseEvent
func callbackQStatusBarKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QStatusBar::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QStatusBar) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QStatusBar::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QStatusBar::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQStatusBarMouseDoubleClickEvent
func callbackQStatusBarMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QStatusBar::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QStatusBar) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QStatusBar::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QStatusBar::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQStatusBarMouseMoveEvent
func callbackQStatusBarMouseMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QStatusBar::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QStatusBar) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QStatusBar::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QStatusBar::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQStatusBarMousePressEvent
func callbackQStatusBarMousePressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QStatusBar::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QStatusBar) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QStatusBar::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QStatusBar::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQStatusBarMouseReleaseEvent
func callbackQStatusBarMouseReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QStatusBar::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QStatusBar) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QStatusBar::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QStatusBar::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQStatusBarTabletEvent
func callbackQStatusBarTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QStatusBar::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QStatusBar) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QStatusBar::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QStatusBar::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQStatusBarWheelEvent
func callbackQStatusBarWheelEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QStatusBar::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QStatusBar) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QStatusBar::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QStatusBar::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQStatusBarTimerEvent
func callbackQStatusBarTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QStatusBar::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QStatusBar) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QStatusBar::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QStatusBar::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQStatusBarChildEvent
func callbackQStatusBarChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QStatusBar::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QStatusBar) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QStatusBar::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QStatusBar::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQStatusBarCustomEvent
func callbackQStatusBarCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QStatusBar::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
