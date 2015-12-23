package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QRasterWindow struct {
	QPaintDeviceWindow
}

type QRasterWindow_ITF interface {
	QPaintDeviceWindow_ITF
	QRasterWindow_PTR() *QRasterWindow
}

func PointerFromQRasterWindow(ptr QRasterWindow_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRasterWindow_PTR().Pointer()
	}
	return nil
}

func NewQRasterWindowFromPointer(ptr unsafe.Pointer) *QRasterWindow {
	var n = new(QRasterWindow)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QRasterWindow_") {
		n.SetObjectName("QRasterWindow_" + qt.Identifier())
	}
	return n
}

func (ptr *QRasterWindow) QRasterWindow_PTR() *QRasterWindow {
	return ptr
}

func NewQRasterWindow(parent QWindow_ITF) *QRasterWindow {
	defer qt.Recovering("QRasterWindow::QRasterWindow")

	return NewQRasterWindowFromPointer(C.QRasterWindow_NewQRasterWindow(PointerFromQWindow(parent)))
}

func (ptr *QRasterWindow) ConnectPaintEvent(f func(event *QPaintEvent)) {
	defer qt.Recovering("connect QRasterWindow::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QRasterWindow) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QRasterWindow::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQRasterWindowPaintEvent
func callbackQRasterWindowPaintEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRasterWindow::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*QPaintEvent))(NewQPaintEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QRasterWindow) ConnectExposeEvent(f func(ev *QExposeEvent)) {
	defer qt.Recovering("connect QRasterWindow::exposeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "exposeEvent", f)
	}
}

func (ptr *QRasterWindow) DisconnectExposeEvent() {
	defer qt.Recovering("disconnect QRasterWindow::exposeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "exposeEvent")
	}
}

//export callbackQRasterWindowExposeEvent
func callbackQRasterWindowExposeEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QRasterWindow::exposeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "exposeEvent"); signal != nil {
		signal.(func(*QExposeEvent))(NewQExposeEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QRasterWindow) ConnectFocusInEvent(f func(ev *QFocusEvent)) {
	defer qt.Recovering("connect QRasterWindow::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QRasterWindow) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QRasterWindow::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQRasterWindowFocusInEvent
func callbackQRasterWindowFocusInEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QRasterWindow::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*QFocusEvent))(NewQFocusEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QRasterWindow) ConnectFocusOutEvent(f func(ev *QFocusEvent)) {
	defer qt.Recovering("connect QRasterWindow::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QRasterWindow) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QRasterWindow::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQRasterWindowFocusOutEvent
func callbackQRasterWindowFocusOutEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QRasterWindow::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*QFocusEvent))(NewQFocusEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QRasterWindow) ConnectHideEvent(f func(ev *QHideEvent)) {
	defer qt.Recovering("connect QRasterWindow::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QRasterWindow) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QRasterWindow::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQRasterWindowHideEvent
func callbackQRasterWindowHideEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QRasterWindow::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*QHideEvent))(NewQHideEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QRasterWindow) ConnectKeyPressEvent(f func(ev *QKeyEvent)) {
	defer qt.Recovering("connect QRasterWindow::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QRasterWindow) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QRasterWindow::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQRasterWindowKeyPressEvent
func callbackQRasterWindowKeyPressEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QRasterWindow::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*QKeyEvent))(NewQKeyEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QRasterWindow) ConnectKeyReleaseEvent(f func(ev *QKeyEvent)) {
	defer qt.Recovering("connect QRasterWindow::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QRasterWindow) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QRasterWindow::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQRasterWindowKeyReleaseEvent
func callbackQRasterWindowKeyReleaseEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QRasterWindow::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*QKeyEvent))(NewQKeyEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QRasterWindow) ConnectMouseDoubleClickEvent(f func(ev *QMouseEvent)) {
	defer qt.Recovering("connect QRasterWindow::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QRasterWindow) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QRasterWindow::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQRasterWindowMouseDoubleClickEvent
func callbackQRasterWindowMouseDoubleClickEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QRasterWindow::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*QMouseEvent))(NewQMouseEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QRasterWindow) ConnectMouseMoveEvent(f func(ev *QMouseEvent)) {
	defer qt.Recovering("connect QRasterWindow::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QRasterWindow) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QRasterWindow::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQRasterWindowMouseMoveEvent
func callbackQRasterWindowMouseMoveEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QRasterWindow::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*QMouseEvent))(NewQMouseEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QRasterWindow) ConnectMousePressEvent(f func(ev *QMouseEvent)) {
	defer qt.Recovering("connect QRasterWindow::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QRasterWindow) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QRasterWindow::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQRasterWindowMousePressEvent
func callbackQRasterWindowMousePressEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QRasterWindow::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*QMouseEvent))(NewQMouseEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QRasterWindow) ConnectMouseReleaseEvent(f func(ev *QMouseEvent)) {
	defer qt.Recovering("connect QRasterWindow::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QRasterWindow) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QRasterWindow::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQRasterWindowMouseReleaseEvent
func callbackQRasterWindowMouseReleaseEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QRasterWindow::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*QMouseEvent))(NewQMouseEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QRasterWindow) ConnectMoveEvent(f func(ev *QMoveEvent)) {
	defer qt.Recovering("connect QRasterWindow::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QRasterWindow) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QRasterWindow::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQRasterWindowMoveEvent
func callbackQRasterWindowMoveEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QRasterWindow::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*QMoveEvent))(NewQMoveEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QRasterWindow) ConnectResizeEvent(f func(ev *QResizeEvent)) {
	defer qt.Recovering("connect QRasterWindow::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QRasterWindow) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QRasterWindow::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQRasterWindowResizeEvent
func callbackQRasterWindowResizeEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QRasterWindow::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*QResizeEvent))(NewQResizeEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QRasterWindow) ConnectShowEvent(f func(ev *QShowEvent)) {
	defer qt.Recovering("connect QRasterWindow::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QRasterWindow) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QRasterWindow::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQRasterWindowShowEvent
func callbackQRasterWindowShowEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QRasterWindow::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*QShowEvent))(NewQShowEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QRasterWindow) ConnectTabletEvent(f func(ev *QTabletEvent)) {
	defer qt.Recovering("connect QRasterWindow::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QRasterWindow) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QRasterWindow::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQRasterWindowTabletEvent
func callbackQRasterWindowTabletEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QRasterWindow::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*QTabletEvent))(NewQTabletEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QRasterWindow) ConnectTouchEvent(f func(ev *QTouchEvent)) {
	defer qt.Recovering("connect QRasterWindow::touchEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "touchEvent", f)
	}
}

func (ptr *QRasterWindow) DisconnectTouchEvent() {
	defer qt.Recovering("disconnect QRasterWindow::touchEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "touchEvent")
	}
}

//export callbackQRasterWindowTouchEvent
func callbackQRasterWindowTouchEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QRasterWindow::touchEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "touchEvent"); signal != nil {
		signal.(func(*QTouchEvent))(NewQTouchEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QRasterWindow) ConnectWheelEvent(f func(ev *QWheelEvent)) {
	defer qt.Recovering("connect QRasterWindow::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QRasterWindow) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QRasterWindow::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQRasterWindowWheelEvent
func callbackQRasterWindowWheelEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QRasterWindow::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*QWheelEvent))(NewQWheelEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QRasterWindow) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QRasterWindow::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QRasterWindow) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QRasterWindow::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQRasterWindowTimerEvent
func callbackQRasterWindowTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRasterWindow::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QRasterWindow) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QRasterWindow::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QRasterWindow) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QRasterWindow::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQRasterWindowChildEvent
func callbackQRasterWindowChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRasterWindow::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QRasterWindow) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QRasterWindow::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QRasterWindow) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QRasterWindow::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQRasterWindowCustomEvent
func callbackQRasterWindowCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRasterWindow::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
