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
func callbackQRasterWindowPaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRasterWindow::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*QPaintEvent))(NewQPaintEventFromPointer(event))
	} else {
		NewQRasterWindowFromPointer(ptr).PaintEventDefault(NewQPaintEventFromPointer(event))
	}
}

func (ptr *QRasterWindow) PaintEvent(event QPaintEvent_ITF) {
	defer qt.Recovering("QRasterWindow::paintEvent")

	if ptr.Pointer() != nil {
		C.QRasterWindow_PaintEvent(ptr.Pointer(), PointerFromQPaintEvent(event))
	}
}

func (ptr *QRasterWindow) PaintEventDefault(event QPaintEvent_ITF) {
	defer qt.Recovering("QRasterWindow::paintEvent")

	if ptr.Pointer() != nil {
		C.QRasterWindow_PaintEventDefault(ptr.Pointer(), PointerFromQPaintEvent(event))
	}
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
func callbackQRasterWindowExposeEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QRasterWindow::exposeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "exposeEvent"); signal != nil {
		signal.(func(*QExposeEvent))(NewQExposeEventFromPointer(ev))
	} else {
		NewQRasterWindowFromPointer(ptr).ExposeEventDefault(NewQExposeEventFromPointer(ev))
	}
}

func (ptr *QRasterWindow) ExposeEvent(ev QExposeEvent_ITF) {
	defer qt.Recovering("QRasterWindow::exposeEvent")

	if ptr.Pointer() != nil {
		C.QRasterWindow_ExposeEvent(ptr.Pointer(), PointerFromQExposeEvent(ev))
	}
}

func (ptr *QRasterWindow) ExposeEventDefault(ev QExposeEvent_ITF) {
	defer qt.Recovering("QRasterWindow::exposeEvent")

	if ptr.Pointer() != nil {
		C.QRasterWindow_ExposeEventDefault(ptr.Pointer(), PointerFromQExposeEvent(ev))
	}
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
func callbackQRasterWindowFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QRasterWindow::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*QFocusEvent))(NewQFocusEventFromPointer(ev))
	} else {
		NewQRasterWindowFromPointer(ptr).FocusInEventDefault(NewQFocusEventFromPointer(ev))
	}
}

func (ptr *QRasterWindow) FocusInEvent(ev QFocusEvent_ITF) {
	defer qt.Recovering("QRasterWindow::focusInEvent")

	if ptr.Pointer() != nil {
		C.QRasterWindow_FocusInEvent(ptr.Pointer(), PointerFromQFocusEvent(ev))
	}
}

func (ptr *QRasterWindow) FocusInEventDefault(ev QFocusEvent_ITF) {
	defer qt.Recovering("QRasterWindow::focusInEvent")

	if ptr.Pointer() != nil {
		C.QRasterWindow_FocusInEventDefault(ptr.Pointer(), PointerFromQFocusEvent(ev))
	}
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
func callbackQRasterWindowFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QRasterWindow::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*QFocusEvent))(NewQFocusEventFromPointer(ev))
	} else {
		NewQRasterWindowFromPointer(ptr).FocusOutEventDefault(NewQFocusEventFromPointer(ev))
	}
}

func (ptr *QRasterWindow) FocusOutEvent(ev QFocusEvent_ITF) {
	defer qt.Recovering("QRasterWindow::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QRasterWindow_FocusOutEvent(ptr.Pointer(), PointerFromQFocusEvent(ev))
	}
}

func (ptr *QRasterWindow) FocusOutEventDefault(ev QFocusEvent_ITF) {
	defer qt.Recovering("QRasterWindow::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QRasterWindow_FocusOutEventDefault(ptr.Pointer(), PointerFromQFocusEvent(ev))
	}
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
func callbackQRasterWindowHideEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QRasterWindow::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*QHideEvent))(NewQHideEventFromPointer(ev))
	} else {
		NewQRasterWindowFromPointer(ptr).HideEventDefault(NewQHideEventFromPointer(ev))
	}
}

func (ptr *QRasterWindow) HideEvent(ev QHideEvent_ITF) {
	defer qt.Recovering("QRasterWindow::hideEvent")

	if ptr.Pointer() != nil {
		C.QRasterWindow_HideEvent(ptr.Pointer(), PointerFromQHideEvent(ev))
	}
}

func (ptr *QRasterWindow) HideEventDefault(ev QHideEvent_ITF) {
	defer qt.Recovering("QRasterWindow::hideEvent")

	if ptr.Pointer() != nil {
		C.QRasterWindow_HideEventDefault(ptr.Pointer(), PointerFromQHideEvent(ev))
	}
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
func callbackQRasterWindowKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QRasterWindow::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*QKeyEvent))(NewQKeyEventFromPointer(ev))
	} else {
		NewQRasterWindowFromPointer(ptr).KeyPressEventDefault(NewQKeyEventFromPointer(ev))
	}
}

func (ptr *QRasterWindow) KeyPressEvent(ev QKeyEvent_ITF) {
	defer qt.Recovering("QRasterWindow::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QRasterWindow_KeyPressEvent(ptr.Pointer(), PointerFromQKeyEvent(ev))
	}
}

func (ptr *QRasterWindow) KeyPressEventDefault(ev QKeyEvent_ITF) {
	defer qt.Recovering("QRasterWindow::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QRasterWindow_KeyPressEventDefault(ptr.Pointer(), PointerFromQKeyEvent(ev))
	}
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
func callbackQRasterWindowKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QRasterWindow::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*QKeyEvent))(NewQKeyEventFromPointer(ev))
	} else {
		NewQRasterWindowFromPointer(ptr).KeyReleaseEventDefault(NewQKeyEventFromPointer(ev))
	}
}

func (ptr *QRasterWindow) KeyReleaseEvent(ev QKeyEvent_ITF) {
	defer qt.Recovering("QRasterWindow::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QRasterWindow_KeyReleaseEvent(ptr.Pointer(), PointerFromQKeyEvent(ev))
	}
}

func (ptr *QRasterWindow) KeyReleaseEventDefault(ev QKeyEvent_ITF) {
	defer qt.Recovering("QRasterWindow::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QRasterWindow_KeyReleaseEventDefault(ptr.Pointer(), PointerFromQKeyEvent(ev))
	}
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
func callbackQRasterWindowMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QRasterWindow::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*QMouseEvent))(NewQMouseEventFromPointer(ev))
	} else {
		NewQRasterWindowFromPointer(ptr).MouseDoubleClickEventDefault(NewQMouseEventFromPointer(ev))
	}
}

func (ptr *QRasterWindow) MouseDoubleClickEvent(ev QMouseEvent_ITF) {
	defer qt.Recovering("QRasterWindow::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QRasterWindow_MouseDoubleClickEvent(ptr.Pointer(), PointerFromQMouseEvent(ev))
	}
}

func (ptr *QRasterWindow) MouseDoubleClickEventDefault(ev QMouseEvent_ITF) {
	defer qt.Recovering("QRasterWindow::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QRasterWindow_MouseDoubleClickEventDefault(ptr.Pointer(), PointerFromQMouseEvent(ev))
	}
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
func callbackQRasterWindowMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QRasterWindow::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*QMouseEvent))(NewQMouseEventFromPointer(ev))
	} else {
		NewQRasterWindowFromPointer(ptr).MouseMoveEventDefault(NewQMouseEventFromPointer(ev))
	}
}

func (ptr *QRasterWindow) MouseMoveEvent(ev QMouseEvent_ITF) {
	defer qt.Recovering("QRasterWindow::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QRasterWindow_MouseMoveEvent(ptr.Pointer(), PointerFromQMouseEvent(ev))
	}
}

func (ptr *QRasterWindow) MouseMoveEventDefault(ev QMouseEvent_ITF) {
	defer qt.Recovering("QRasterWindow::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QRasterWindow_MouseMoveEventDefault(ptr.Pointer(), PointerFromQMouseEvent(ev))
	}
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
func callbackQRasterWindowMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QRasterWindow::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*QMouseEvent))(NewQMouseEventFromPointer(ev))
	} else {
		NewQRasterWindowFromPointer(ptr).MousePressEventDefault(NewQMouseEventFromPointer(ev))
	}
}

func (ptr *QRasterWindow) MousePressEvent(ev QMouseEvent_ITF) {
	defer qt.Recovering("QRasterWindow::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QRasterWindow_MousePressEvent(ptr.Pointer(), PointerFromQMouseEvent(ev))
	}
}

func (ptr *QRasterWindow) MousePressEventDefault(ev QMouseEvent_ITF) {
	defer qt.Recovering("QRasterWindow::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QRasterWindow_MousePressEventDefault(ptr.Pointer(), PointerFromQMouseEvent(ev))
	}
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
func callbackQRasterWindowMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QRasterWindow::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*QMouseEvent))(NewQMouseEventFromPointer(ev))
	} else {
		NewQRasterWindowFromPointer(ptr).MouseReleaseEventDefault(NewQMouseEventFromPointer(ev))
	}
}

func (ptr *QRasterWindow) MouseReleaseEvent(ev QMouseEvent_ITF) {
	defer qt.Recovering("QRasterWindow::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QRasterWindow_MouseReleaseEvent(ptr.Pointer(), PointerFromQMouseEvent(ev))
	}
}

func (ptr *QRasterWindow) MouseReleaseEventDefault(ev QMouseEvent_ITF) {
	defer qt.Recovering("QRasterWindow::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QRasterWindow_MouseReleaseEventDefault(ptr.Pointer(), PointerFromQMouseEvent(ev))
	}
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
func callbackQRasterWindowMoveEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QRasterWindow::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*QMoveEvent))(NewQMoveEventFromPointer(ev))
	} else {
		NewQRasterWindowFromPointer(ptr).MoveEventDefault(NewQMoveEventFromPointer(ev))
	}
}

func (ptr *QRasterWindow) MoveEvent(ev QMoveEvent_ITF) {
	defer qt.Recovering("QRasterWindow::moveEvent")

	if ptr.Pointer() != nil {
		C.QRasterWindow_MoveEvent(ptr.Pointer(), PointerFromQMoveEvent(ev))
	}
}

func (ptr *QRasterWindow) MoveEventDefault(ev QMoveEvent_ITF) {
	defer qt.Recovering("QRasterWindow::moveEvent")

	if ptr.Pointer() != nil {
		C.QRasterWindow_MoveEventDefault(ptr.Pointer(), PointerFromQMoveEvent(ev))
	}
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
func callbackQRasterWindowResizeEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QRasterWindow::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*QResizeEvent))(NewQResizeEventFromPointer(ev))
	} else {
		NewQRasterWindowFromPointer(ptr).ResizeEventDefault(NewQResizeEventFromPointer(ev))
	}
}

func (ptr *QRasterWindow) ResizeEvent(ev QResizeEvent_ITF) {
	defer qt.Recovering("QRasterWindow::resizeEvent")

	if ptr.Pointer() != nil {
		C.QRasterWindow_ResizeEvent(ptr.Pointer(), PointerFromQResizeEvent(ev))
	}
}

func (ptr *QRasterWindow) ResizeEventDefault(ev QResizeEvent_ITF) {
	defer qt.Recovering("QRasterWindow::resizeEvent")

	if ptr.Pointer() != nil {
		C.QRasterWindow_ResizeEventDefault(ptr.Pointer(), PointerFromQResizeEvent(ev))
	}
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
func callbackQRasterWindowShowEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QRasterWindow::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*QShowEvent))(NewQShowEventFromPointer(ev))
	} else {
		NewQRasterWindowFromPointer(ptr).ShowEventDefault(NewQShowEventFromPointer(ev))
	}
}

func (ptr *QRasterWindow) ShowEvent(ev QShowEvent_ITF) {
	defer qt.Recovering("QRasterWindow::showEvent")

	if ptr.Pointer() != nil {
		C.QRasterWindow_ShowEvent(ptr.Pointer(), PointerFromQShowEvent(ev))
	}
}

func (ptr *QRasterWindow) ShowEventDefault(ev QShowEvent_ITF) {
	defer qt.Recovering("QRasterWindow::showEvent")

	if ptr.Pointer() != nil {
		C.QRasterWindow_ShowEventDefault(ptr.Pointer(), PointerFromQShowEvent(ev))
	}
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
func callbackQRasterWindowTabletEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QRasterWindow::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*QTabletEvent))(NewQTabletEventFromPointer(ev))
	} else {
		NewQRasterWindowFromPointer(ptr).TabletEventDefault(NewQTabletEventFromPointer(ev))
	}
}

func (ptr *QRasterWindow) TabletEvent(ev QTabletEvent_ITF) {
	defer qt.Recovering("QRasterWindow::tabletEvent")

	if ptr.Pointer() != nil {
		C.QRasterWindow_TabletEvent(ptr.Pointer(), PointerFromQTabletEvent(ev))
	}
}

func (ptr *QRasterWindow) TabletEventDefault(ev QTabletEvent_ITF) {
	defer qt.Recovering("QRasterWindow::tabletEvent")

	if ptr.Pointer() != nil {
		C.QRasterWindow_TabletEventDefault(ptr.Pointer(), PointerFromQTabletEvent(ev))
	}
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
func callbackQRasterWindowTouchEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QRasterWindow::touchEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "touchEvent"); signal != nil {
		signal.(func(*QTouchEvent))(NewQTouchEventFromPointer(ev))
	} else {
		NewQRasterWindowFromPointer(ptr).TouchEventDefault(NewQTouchEventFromPointer(ev))
	}
}

func (ptr *QRasterWindow) TouchEvent(ev QTouchEvent_ITF) {
	defer qt.Recovering("QRasterWindow::touchEvent")

	if ptr.Pointer() != nil {
		C.QRasterWindow_TouchEvent(ptr.Pointer(), PointerFromQTouchEvent(ev))
	}
}

func (ptr *QRasterWindow) TouchEventDefault(ev QTouchEvent_ITF) {
	defer qt.Recovering("QRasterWindow::touchEvent")

	if ptr.Pointer() != nil {
		C.QRasterWindow_TouchEventDefault(ptr.Pointer(), PointerFromQTouchEvent(ev))
	}
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
func callbackQRasterWindowWheelEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QRasterWindow::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*QWheelEvent))(NewQWheelEventFromPointer(ev))
	} else {
		NewQRasterWindowFromPointer(ptr).WheelEventDefault(NewQWheelEventFromPointer(ev))
	}
}

func (ptr *QRasterWindow) WheelEvent(ev QWheelEvent_ITF) {
	defer qt.Recovering("QRasterWindow::wheelEvent")

	if ptr.Pointer() != nil {
		C.QRasterWindow_WheelEvent(ptr.Pointer(), PointerFromQWheelEvent(ev))
	}
}

func (ptr *QRasterWindow) WheelEventDefault(ev QWheelEvent_ITF) {
	defer qt.Recovering("QRasterWindow::wheelEvent")

	if ptr.Pointer() != nil {
		C.QRasterWindow_WheelEventDefault(ptr.Pointer(), PointerFromQWheelEvent(ev))
	}
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
func callbackQRasterWindowTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRasterWindow::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQRasterWindowFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QRasterWindow) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QRasterWindow::timerEvent")

	if ptr.Pointer() != nil {
		C.QRasterWindow_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QRasterWindow) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QRasterWindow::timerEvent")

	if ptr.Pointer() != nil {
		C.QRasterWindow_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
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
func callbackQRasterWindowChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRasterWindow::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQRasterWindowFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QRasterWindow) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QRasterWindow::childEvent")

	if ptr.Pointer() != nil {
		C.QRasterWindow_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QRasterWindow) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QRasterWindow::childEvent")

	if ptr.Pointer() != nil {
		C.QRasterWindow_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
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
func callbackQRasterWindowCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QRasterWindow::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQRasterWindowFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QRasterWindow) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QRasterWindow::customEvent")

	if ptr.Pointer() != nil {
		C.QRasterWindow_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QRasterWindow) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QRasterWindow::customEvent")

	if ptr.Pointer() != nil {
		C.QRasterWindow_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
