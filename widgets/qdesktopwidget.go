package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QDesktopWidget struct {
	QWidget
}

type QDesktopWidget_ITF interface {
	QWidget_ITF
	QDesktopWidget_PTR() *QDesktopWidget
}

func PointerFromQDesktopWidget(ptr QDesktopWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDesktopWidget_PTR().Pointer()
	}
	return nil
}

func NewQDesktopWidgetFromPointer(ptr unsafe.Pointer) *QDesktopWidget {
	var n = new(QDesktopWidget)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QDesktopWidget_") {
		n.SetObjectName("QDesktopWidget_" + qt.Identifier())
	}
	return n
}

func (ptr *QDesktopWidget) QDesktopWidget_PTR() *QDesktopWidget {
	return ptr
}

func (ptr *QDesktopWidget) AvailableGeometry2(widget QWidget_ITF) *core.QRect {
	defer qt.Recovering("QDesktopWidget::availableGeometry")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QDesktopWidget_AvailableGeometry2(ptr.Pointer(), PointerFromQWidget(widget)))
	}
	return nil
}

func (ptr *QDesktopWidget) AvailableGeometry(screen int) *core.QRect {
	defer qt.Recovering("QDesktopWidget::availableGeometry")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QDesktopWidget_AvailableGeometry(ptr.Pointer(), C.int(screen)))
	}
	return nil
}

func (ptr *QDesktopWidget) IsVirtualDesktop() bool {
	defer qt.Recovering("QDesktopWidget::isVirtualDesktop")

	if ptr.Pointer() != nil {
		return C.QDesktopWidget_IsVirtualDesktop(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDesktopWidget) PrimaryScreen() int {
	defer qt.Recovering("QDesktopWidget::primaryScreen")

	if ptr.Pointer() != nil {
		return int(C.QDesktopWidget_PrimaryScreen(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDesktopWidget) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QDesktopWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQDesktopWidgetResizeEvent
func callbackQDesktopWidgetResizeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDesktopWidget::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDesktopWidget) Screen(screen int) *QWidget {
	defer qt.Recovering("QDesktopWidget::screen")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QDesktopWidget_Screen(ptr.Pointer(), C.int(screen)))
	}
	return nil
}

func (ptr *QDesktopWidget) ScreenGeometry2(widget QWidget_ITF) *core.QRect {
	defer qt.Recovering("QDesktopWidget::screenGeometry")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QDesktopWidget_ScreenGeometry2(ptr.Pointer(), PointerFromQWidget(widget)))
	}
	return nil
}

func (ptr *QDesktopWidget) ScreenGeometry(screen int) *core.QRect {
	defer qt.Recovering("QDesktopWidget::screenGeometry")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QDesktopWidget_ScreenGeometry(ptr.Pointer(), C.int(screen)))
	}
	return nil
}

func (ptr *QDesktopWidget) ScreenNumber2(point core.QPoint_ITF) int {
	defer qt.Recovering("QDesktopWidget::screenNumber")

	if ptr.Pointer() != nil {
		return int(C.QDesktopWidget_ScreenNumber2(ptr.Pointer(), core.PointerFromQPoint(point)))
	}
	return 0
}

func (ptr *QDesktopWidget) ScreenNumber(widget QWidget_ITF) int {
	defer qt.Recovering("QDesktopWidget::screenNumber")

	if ptr.Pointer() != nil {
		return int(C.QDesktopWidget_ScreenNumber(ptr.Pointer(), PointerFromQWidget(widget)))
	}
	return 0
}

func (ptr *QDesktopWidget) AvailableGeometry3(p core.QPoint_ITF) *core.QRect {
	defer qt.Recovering("QDesktopWidget::availableGeometry")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QDesktopWidget_AvailableGeometry3(ptr.Pointer(), core.PointerFromQPoint(p)))
	}
	return nil
}

func (ptr *QDesktopWidget) ConnectResized(f func(screen int)) {
	defer qt.Recovering("connect QDesktopWidget::resized")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_ConnectResized(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "resized", f)
	}
}

func (ptr *QDesktopWidget) DisconnectResized() {
	defer qt.Recovering("disconnect QDesktopWidget::resized")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_DisconnectResized(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "resized")
	}
}

//export callbackQDesktopWidgetResized
func callbackQDesktopWidgetResized(ptrName *C.char, screen C.int) {
	defer qt.Recovering("callback QDesktopWidget::resized")

	if signal := qt.GetSignal(C.GoString(ptrName), "resized"); signal != nil {
		signal.(func(int))(int(screen))
	}

}

func (ptr *QDesktopWidget) ScreenCount() int {
	defer qt.Recovering("QDesktopWidget::screenCount")

	if ptr.Pointer() != nil {
		return int(C.QDesktopWidget_ScreenCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDesktopWidget) ConnectScreenCountChanged(f func(newCount int)) {
	defer qt.Recovering("connect QDesktopWidget::screenCountChanged")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_ConnectScreenCountChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "screenCountChanged", f)
	}
}

func (ptr *QDesktopWidget) DisconnectScreenCountChanged() {
	defer qt.Recovering("disconnect QDesktopWidget::screenCountChanged")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_DisconnectScreenCountChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "screenCountChanged")
	}
}

//export callbackQDesktopWidgetScreenCountChanged
func callbackQDesktopWidgetScreenCountChanged(ptrName *C.char, newCount C.int) {
	defer qt.Recovering("callback QDesktopWidget::screenCountChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "screenCountChanged"); signal != nil {
		signal.(func(int))(int(newCount))
	}

}

func (ptr *QDesktopWidget) ScreenGeometry3(p core.QPoint_ITF) *core.QRect {
	defer qt.Recovering("QDesktopWidget::screenGeometry")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QDesktopWidget_ScreenGeometry3(ptr.Pointer(), core.PointerFromQPoint(p)))
	}
	return nil
}

func (ptr *QDesktopWidget) ConnectWorkAreaResized(f func(screen int)) {
	defer qt.Recovering("connect QDesktopWidget::workAreaResized")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_ConnectWorkAreaResized(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "workAreaResized", f)
	}
}

func (ptr *QDesktopWidget) DisconnectWorkAreaResized() {
	defer qt.Recovering("disconnect QDesktopWidget::workAreaResized")

	if ptr.Pointer() != nil {
		C.QDesktopWidget_DisconnectWorkAreaResized(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "workAreaResized")
	}
}

//export callbackQDesktopWidgetWorkAreaResized
func callbackQDesktopWidgetWorkAreaResized(ptrName *C.char, screen C.int) {
	defer qt.Recovering("callback QDesktopWidget::workAreaResized")

	if signal := qt.GetSignal(C.GoString(ptrName), "workAreaResized"); signal != nil {
		signal.(func(int))(int(screen))
	}

}

func (ptr *QDesktopWidget) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QDesktopWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQDesktopWidgetActionEvent
func callbackQDesktopWidgetActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDesktopWidget::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDesktopWidget) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QDesktopWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQDesktopWidgetDragEnterEvent
func callbackQDesktopWidgetDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDesktopWidget::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDesktopWidget) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QDesktopWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQDesktopWidgetDragLeaveEvent
func callbackQDesktopWidgetDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDesktopWidget::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDesktopWidget) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QDesktopWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQDesktopWidgetDragMoveEvent
func callbackQDesktopWidgetDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDesktopWidget::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDesktopWidget) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QDesktopWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQDesktopWidgetDropEvent
func callbackQDesktopWidgetDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDesktopWidget::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDesktopWidget) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDesktopWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQDesktopWidgetEnterEvent
func callbackQDesktopWidgetEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDesktopWidget::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDesktopWidget) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QDesktopWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQDesktopWidgetFocusOutEvent
func callbackQDesktopWidgetFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDesktopWidget::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDesktopWidget) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QDesktopWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQDesktopWidgetHideEvent
func callbackQDesktopWidgetHideEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDesktopWidget::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDesktopWidget) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDesktopWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQDesktopWidgetLeaveEvent
func callbackQDesktopWidgetLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDesktopWidget::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDesktopWidget) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QDesktopWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQDesktopWidgetMoveEvent
func callbackQDesktopWidgetMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDesktopWidget::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDesktopWidget) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QDesktopWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQDesktopWidgetPaintEvent
func callbackQDesktopWidgetPaintEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDesktopWidget::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDesktopWidget) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QDesktopWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QDesktopWidget) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QDesktopWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQDesktopWidgetSetVisible
func callbackQDesktopWidgetSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QDesktopWidget::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QDesktopWidget) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QDesktopWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQDesktopWidgetShowEvent
func callbackQDesktopWidgetShowEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDesktopWidget::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDesktopWidget) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QDesktopWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQDesktopWidgetCloseEvent
func callbackQDesktopWidgetCloseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDesktopWidget::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDesktopWidget) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QDesktopWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQDesktopWidgetContextMenuEvent
func callbackQDesktopWidgetContextMenuEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDesktopWidget::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDesktopWidget) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QDesktopWidget::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QDesktopWidget) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QDesktopWidget::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQDesktopWidgetInitPainter
func callbackQDesktopWidgetInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QDesktopWidget::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

}

func (ptr *QDesktopWidget) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QDesktopWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQDesktopWidgetInputMethodEvent
func callbackQDesktopWidgetInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDesktopWidget::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDesktopWidget) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QDesktopWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQDesktopWidgetKeyPressEvent
func callbackQDesktopWidgetKeyPressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDesktopWidget::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDesktopWidget) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QDesktopWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQDesktopWidgetKeyReleaseEvent
func callbackQDesktopWidgetKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDesktopWidget::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDesktopWidget) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDesktopWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQDesktopWidgetMouseDoubleClickEvent
func callbackQDesktopWidgetMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDesktopWidget::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDesktopWidget) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDesktopWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQDesktopWidgetMouseMoveEvent
func callbackQDesktopWidgetMouseMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDesktopWidget::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDesktopWidget) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDesktopWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQDesktopWidgetMousePressEvent
func callbackQDesktopWidgetMousePressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDesktopWidget::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDesktopWidget) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDesktopWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQDesktopWidgetMouseReleaseEvent
func callbackQDesktopWidgetMouseReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDesktopWidget::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDesktopWidget) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QDesktopWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQDesktopWidgetTabletEvent
func callbackQDesktopWidgetTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDesktopWidget::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDesktopWidget) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QDesktopWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQDesktopWidgetWheelEvent
func callbackQDesktopWidgetWheelEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDesktopWidget::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDesktopWidget) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDesktopWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQDesktopWidgetTimerEvent
func callbackQDesktopWidgetTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDesktopWidget::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDesktopWidget) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDesktopWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQDesktopWidgetChildEvent
func callbackQDesktopWidgetChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDesktopWidget::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QDesktopWidget) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDesktopWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QDesktopWidget) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDesktopWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQDesktopWidgetCustomEvent
func callbackQDesktopWidgetCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDesktopWidget::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
