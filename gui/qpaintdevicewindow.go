package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QPaintDeviceWindow struct {
	QWindow
	QPaintDevice
}

type QPaintDeviceWindow_ITF interface {
	QWindow_ITF
	QPaintDevice_ITF
	QPaintDeviceWindow_PTR() *QPaintDeviceWindow
}

func (p *QPaintDeviceWindow) Pointer() unsafe.Pointer {
	return p.QWindow_PTR().Pointer()
}

func (p *QPaintDeviceWindow) SetPointer(ptr unsafe.Pointer) {
	p.QWindow_PTR().SetPointer(ptr)
	p.QPaintDevice_PTR().SetPointer(ptr)
}

func PointerFromQPaintDeviceWindow(ptr QPaintDeviceWindow_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPaintDeviceWindow_PTR().Pointer()
	}
	return nil
}

func NewQPaintDeviceWindowFromPointer(ptr unsafe.Pointer) *QPaintDeviceWindow {
	var n = new(QPaintDeviceWindow)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QPaintDeviceWindow_") {
		n.SetObjectName("QPaintDeviceWindow_" + qt.Identifier())
	}
	return n
}

func (ptr *QPaintDeviceWindow) QPaintDeviceWindow_PTR() *QPaintDeviceWindow {
	return ptr
}

func (ptr *QPaintDeviceWindow) ConnectPaintEvent(f func(event *QPaintEvent)) {
	defer qt.Recovering("connect QPaintDeviceWindow::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QPaintDeviceWindow) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QPaintDeviceWindow::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQPaintDeviceWindowPaintEvent
func callbackQPaintDeviceWindowPaintEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QPaintDeviceWindow::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*QPaintEvent))(NewQPaintEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QPaintDeviceWindow) Update3() {
	defer qt.Recovering("QPaintDeviceWindow::update")

	if ptr.Pointer() != nil {
		C.QPaintDeviceWindow_Update3(ptr.Pointer())
	}
}

func (ptr *QPaintDeviceWindow) Update(rect core.QRect_ITF) {
	defer qt.Recovering("QPaintDeviceWindow::update")

	if ptr.Pointer() != nil {
		C.QPaintDeviceWindow_Update(ptr.Pointer(), core.PointerFromQRect(rect))
	}
}

func (ptr *QPaintDeviceWindow) Update2(region QRegion_ITF) {
	defer qt.Recovering("QPaintDeviceWindow::update")

	if ptr.Pointer() != nil {
		C.QPaintDeviceWindow_Update2(ptr.Pointer(), PointerFromQRegion(region))
	}
}

func (ptr *QPaintDeviceWindow) ConnectExposeEvent(f func(ev *QExposeEvent)) {
	defer qt.Recovering("connect QPaintDeviceWindow::exposeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "exposeEvent", f)
	}
}

func (ptr *QPaintDeviceWindow) DisconnectExposeEvent() {
	defer qt.Recovering("disconnect QPaintDeviceWindow::exposeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "exposeEvent")
	}
}

//export callbackQPaintDeviceWindowExposeEvent
func callbackQPaintDeviceWindowExposeEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QPaintDeviceWindow::exposeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "exposeEvent"); signal != nil {
		signal.(func(*QExposeEvent))(NewQExposeEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QPaintDeviceWindow) ConnectFocusInEvent(f func(ev *QFocusEvent)) {
	defer qt.Recovering("connect QPaintDeviceWindow::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QPaintDeviceWindow) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QPaintDeviceWindow::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQPaintDeviceWindowFocusInEvent
func callbackQPaintDeviceWindowFocusInEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QPaintDeviceWindow::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*QFocusEvent))(NewQFocusEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QPaintDeviceWindow) ConnectFocusOutEvent(f func(ev *QFocusEvent)) {
	defer qt.Recovering("connect QPaintDeviceWindow::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QPaintDeviceWindow) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QPaintDeviceWindow::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQPaintDeviceWindowFocusOutEvent
func callbackQPaintDeviceWindowFocusOutEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QPaintDeviceWindow::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*QFocusEvent))(NewQFocusEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QPaintDeviceWindow) ConnectHideEvent(f func(ev *QHideEvent)) {
	defer qt.Recovering("connect QPaintDeviceWindow::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QPaintDeviceWindow) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QPaintDeviceWindow::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQPaintDeviceWindowHideEvent
func callbackQPaintDeviceWindowHideEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QPaintDeviceWindow::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*QHideEvent))(NewQHideEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QPaintDeviceWindow) ConnectKeyPressEvent(f func(ev *QKeyEvent)) {
	defer qt.Recovering("connect QPaintDeviceWindow::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QPaintDeviceWindow) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QPaintDeviceWindow::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQPaintDeviceWindowKeyPressEvent
func callbackQPaintDeviceWindowKeyPressEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QPaintDeviceWindow::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*QKeyEvent))(NewQKeyEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QPaintDeviceWindow) ConnectKeyReleaseEvent(f func(ev *QKeyEvent)) {
	defer qt.Recovering("connect QPaintDeviceWindow::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QPaintDeviceWindow) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QPaintDeviceWindow::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQPaintDeviceWindowKeyReleaseEvent
func callbackQPaintDeviceWindowKeyReleaseEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QPaintDeviceWindow::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*QKeyEvent))(NewQKeyEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QPaintDeviceWindow) ConnectMouseDoubleClickEvent(f func(ev *QMouseEvent)) {
	defer qt.Recovering("connect QPaintDeviceWindow::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QPaintDeviceWindow) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QPaintDeviceWindow::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQPaintDeviceWindowMouseDoubleClickEvent
func callbackQPaintDeviceWindowMouseDoubleClickEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QPaintDeviceWindow::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*QMouseEvent))(NewQMouseEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QPaintDeviceWindow) ConnectMouseMoveEvent(f func(ev *QMouseEvent)) {
	defer qt.Recovering("connect QPaintDeviceWindow::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QPaintDeviceWindow) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QPaintDeviceWindow::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQPaintDeviceWindowMouseMoveEvent
func callbackQPaintDeviceWindowMouseMoveEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QPaintDeviceWindow::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*QMouseEvent))(NewQMouseEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QPaintDeviceWindow) ConnectMousePressEvent(f func(ev *QMouseEvent)) {
	defer qt.Recovering("connect QPaintDeviceWindow::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QPaintDeviceWindow) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QPaintDeviceWindow::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQPaintDeviceWindowMousePressEvent
func callbackQPaintDeviceWindowMousePressEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QPaintDeviceWindow::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*QMouseEvent))(NewQMouseEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QPaintDeviceWindow) ConnectMouseReleaseEvent(f func(ev *QMouseEvent)) {
	defer qt.Recovering("connect QPaintDeviceWindow::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QPaintDeviceWindow) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QPaintDeviceWindow::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQPaintDeviceWindowMouseReleaseEvent
func callbackQPaintDeviceWindowMouseReleaseEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QPaintDeviceWindow::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*QMouseEvent))(NewQMouseEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QPaintDeviceWindow) ConnectMoveEvent(f func(ev *QMoveEvent)) {
	defer qt.Recovering("connect QPaintDeviceWindow::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QPaintDeviceWindow) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QPaintDeviceWindow::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQPaintDeviceWindowMoveEvent
func callbackQPaintDeviceWindowMoveEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QPaintDeviceWindow::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*QMoveEvent))(NewQMoveEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QPaintDeviceWindow) ConnectResizeEvent(f func(ev *QResizeEvent)) {
	defer qt.Recovering("connect QPaintDeviceWindow::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QPaintDeviceWindow) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QPaintDeviceWindow::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQPaintDeviceWindowResizeEvent
func callbackQPaintDeviceWindowResizeEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QPaintDeviceWindow::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*QResizeEvent))(NewQResizeEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QPaintDeviceWindow) ConnectShowEvent(f func(ev *QShowEvent)) {
	defer qt.Recovering("connect QPaintDeviceWindow::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QPaintDeviceWindow) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QPaintDeviceWindow::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQPaintDeviceWindowShowEvent
func callbackQPaintDeviceWindowShowEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QPaintDeviceWindow::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*QShowEvent))(NewQShowEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QPaintDeviceWindow) ConnectTabletEvent(f func(ev *QTabletEvent)) {
	defer qt.Recovering("connect QPaintDeviceWindow::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QPaintDeviceWindow) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QPaintDeviceWindow::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQPaintDeviceWindowTabletEvent
func callbackQPaintDeviceWindowTabletEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QPaintDeviceWindow::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*QTabletEvent))(NewQTabletEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QPaintDeviceWindow) ConnectTouchEvent(f func(ev *QTouchEvent)) {
	defer qt.Recovering("connect QPaintDeviceWindow::touchEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "touchEvent", f)
	}
}

func (ptr *QPaintDeviceWindow) DisconnectTouchEvent() {
	defer qt.Recovering("disconnect QPaintDeviceWindow::touchEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "touchEvent")
	}
}

//export callbackQPaintDeviceWindowTouchEvent
func callbackQPaintDeviceWindowTouchEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QPaintDeviceWindow::touchEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "touchEvent"); signal != nil {
		signal.(func(*QTouchEvent))(NewQTouchEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QPaintDeviceWindow) ConnectWheelEvent(f func(ev *QWheelEvent)) {
	defer qt.Recovering("connect QPaintDeviceWindow::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QPaintDeviceWindow) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QPaintDeviceWindow::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQPaintDeviceWindowWheelEvent
func callbackQPaintDeviceWindowWheelEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QPaintDeviceWindow::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*QWheelEvent))(NewQWheelEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QPaintDeviceWindow) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QPaintDeviceWindow::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QPaintDeviceWindow) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QPaintDeviceWindow::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQPaintDeviceWindowTimerEvent
func callbackQPaintDeviceWindowTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QPaintDeviceWindow::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QPaintDeviceWindow) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QPaintDeviceWindow::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QPaintDeviceWindow) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QPaintDeviceWindow::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQPaintDeviceWindowChildEvent
func callbackQPaintDeviceWindowChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QPaintDeviceWindow::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QPaintDeviceWindow) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QPaintDeviceWindow::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QPaintDeviceWindow) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QPaintDeviceWindow::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQPaintDeviceWindowCustomEvent
func callbackQPaintDeviceWindowCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QPaintDeviceWindow::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
