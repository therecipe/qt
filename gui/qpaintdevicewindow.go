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
func callbackQPaintDeviceWindowPaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QPaintDeviceWindow::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*QPaintEvent))(NewQPaintEventFromPointer(event))
	} else {
		NewQPaintDeviceWindowFromPointer(ptr).PaintEventDefault(NewQPaintEventFromPointer(event))
	}
}

func (ptr *QPaintDeviceWindow) PaintEvent(event QPaintEvent_ITF) {
	defer qt.Recovering("QPaintDeviceWindow::paintEvent")

	if ptr.Pointer() != nil {
		C.QPaintDeviceWindow_PaintEvent(ptr.Pointer(), PointerFromQPaintEvent(event))
	}
}

func (ptr *QPaintDeviceWindow) PaintEventDefault(event QPaintEvent_ITF) {
	defer qt.Recovering("QPaintDeviceWindow::paintEvent")

	if ptr.Pointer() != nil {
		C.QPaintDeviceWindow_PaintEventDefault(ptr.Pointer(), PointerFromQPaintEvent(event))
	}
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
func callbackQPaintDeviceWindowExposeEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QPaintDeviceWindow::exposeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "exposeEvent"); signal != nil {
		signal.(func(*QExposeEvent))(NewQExposeEventFromPointer(ev))
	} else {
		NewQPaintDeviceWindowFromPointer(ptr).ExposeEventDefault(NewQExposeEventFromPointer(ev))
	}
}

func (ptr *QPaintDeviceWindow) ExposeEvent(ev QExposeEvent_ITF) {
	defer qt.Recovering("QPaintDeviceWindow::exposeEvent")

	if ptr.Pointer() != nil {
		C.QPaintDeviceWindow_ExposeEvent(ptr.Pointer(), PointerFromQExposeEvent(ev))
	}
}

func (ptr *QPaintDeviceWindow) ExposeEventDefault(ev QExposeEvent_ITF) {
	defer qt.Recovering("QPaintDeviceWindow::exposeEvent")

	if ptr.Pointer() != nil {
		C.QPaintDeviceWindow_ExposeEventDefault(ptr.Pointer(), PointerFromQExposeEvent(ev))
	}
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
func callbackQPaintDeviceWindowFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QPaintDeviceWindow::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*QFocusEvent))(NewQFocusEventFromPointer(ev))
	} else {
		NewQPaintDeviceWindowFromPointer(ptr).FocusInEventDefault(NewQFocusEventFromPointer(ev))
	}
}

func (ptr *QPaintDeviceWindow) FocusInEvent(ev QFocusEvent_ITF) {
	defer qt.Recovering("QPaintDeviceWindow::focusInEvent")

	if ptr.Pointer() != nil {
		C.QPaintDeviceWindow_FocusInEvent(ptr.Pointer(), PointerFromQFocusEvent(ev))
	}
}

func (ptr *QPaintDeviceWindow) FocusInEventDefault(ev QFocusEvent_ITF) {
	defer qt.Recovering("QPaintDeviceWindow::focusInEvent")

	if ptr.Pointer() != nil {
		C.QPaintDeviceWindow_FocusInEventDefault(ptr.Pointer(), PointerFromQFocusEvent(ev))
	}
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
func callbackQPaintDeviceWindowFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QPaintDeviceWindow::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*QFocusEvent))(NewQFocusEventFromPointer(ev))
	} else {
		NewQPaintDeviceWindowFromPointer(ptr).FocusOutEventDefault(NewQFocusEventFromPointer(ev))
	}
}

func (ptr *QPaintDeviceWindow) FocusOutEvent(ev QFocusEvent_ITF) {
	defer qt.Recovering("QPaintDeviceWindow::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QPaintDeviceWindow_FocusOutEvent(ptr.Pointer(), PointerFromQFocusEvent(ev))
	}
}

func (ptr *QPaintDeviceWindow) FocusOutEventDefault(ev QFocusEvent_ITF) {
	defer qt.Recovering("QPaintDeviceWindow::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QPaintDeviceWindow_FocusOutEventDefault(ptr.Pointer(), PointerFromQFocusEvent(ev))
	}
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
func callbackQPaintDeviceWindowHideEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QPaintDeviceWindow::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*QHideEvent))(NewQHideEventFromPointer(ev))
	} else {
		NewQPaintDeviceWindowFromPointer(ptr).HideEventDefault(NewQHideEventFromPointer(ev))
	}
}

func (ptr *QPaintDeviceWindow) HideEvent(ev QHideEvent_ITF) {
	defer qt.Recovering("QPaintDeviceWindow::hideEvent")

	if ptr.Pointer() != nil {
		C.QPaintDeviceWindow_HideEvent(ptr.Pointer(), PointerFromQHideEvent(ev))
	}
}

func (ptr *QPaintDeviceWindow) HideEventDefault(ev QHideEvent_ITF) {
	defer qt.Recovering("QPaintDeviceWindow::hideEvent")

	if ptr.Pointer() != nil {
		C.QPaintDeviceWindow_HideEventDefault(ptr.Pointer(), PointerFromQHideEvent(ev))
	}
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
func callbackQPaintDeviceWindowKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QPaintDeviceWindow::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*QKeyEvent))(NewQKeyEventFromPointer(ev))
	} else {
		NewQPaintDeviceWindowFromPointer(ptr).KeyPressEventDefault(NewQKeyEventFromPointer(ev))
	}
}

func (ptr *QPaintDeviceWindow) KeyPressEvent(ev QKeyEvent_ITF) {
	defer qt.Recovering("QPaintDeviceWindow::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QPaintDeviceWindow_KeyPressEvent(ptr.Pointer(), PointerFromQKeyEvent(ev))
	}
}

func (ptr *QPaintDeviceWindow) KeyPressEventDefault(ev QKeyEvent_ITF) {
	defer qt.Recovering("QPaintDeviceWindow::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QPaintDeviceWindow_KeyPressEventDefault(ptr.Pointer(), PointerFromQKeyEvent(ev))
	}
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
func callbackQPaintDeviceWindowKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QPaintDeviceWindow::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*QKeyEvent))(NewQKeyEventFromPointer(ev))
	} else {
		NewQPaintDeviceWindowFromPointer(ptr).KeyReleaseEventDefault(NewQKeyEventFromPointer(ev))
	}
}

func (ptr *QPaintDeviceWindow) KeyReleaseEvent(ev QKeyEvent_ITF) {
	defer qt.Recovering("QPaintDeviceWindow::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QPaintDeviceWindow_KeyReleaseEvent(ptr.Pointer(), PointerFromQKeyEvent(ev))
	}
}

func (ptr *QPaintDeviceWindow) KeyReleaseEventDefault(ev QKeyEvent_ITF) {
	defer qt.Recovering("QPaintDeviceWindow::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QPaintDeviceWindow_KeyReleaseEventDefault(ptr.Pointer(), PointerFromQKeyEvent(ev))
	}
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
func callbackQPaintDeviceWindowMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QPaintDeviceWindow::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*QMouseEvent))(NewQMouseEventFromPointer(ev))
	} else {
		NewQPaintDeviceWindowFromPointer(ptr).MouseDoubleClickEventDefault(NewQMouseEventFromPointer(ev))
	}
}

func (ptr *QPaintDeviceWindow) MouseDoubleClickEvent(ev QMouseEvent_ITF) {
	defer qt.Recovering("QPaintDeviceWindow::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QPaintDeviceWindow_MouseDoubleClickEvent(ptr.Pointer(), PointerFromQMouseEvent(ev))
	}
}

func (ptr *QPaintDeviceWindow) MouseDoubleClickEventDefault(ev QMouseEvent_ITF) {
	defer qt.Recovering("QPaintDeviceWindow::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QPaintDeviceWindow_MouseDoubleClickEventDefault(ptr.Pointer(), PointerFromQMouseEvent(ev))
	}
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
func callbackQPaintDeviceWindowMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QPaintDeviceWindow::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*QMouseEvent))(NewQMouseEventFromPointer(ev))
	} else {
		NewQPaintDeviceWindowFromPointer(ptr).MouseMoveEventDefault(NewQMouseEventFromPointer(ev))
	}
}

func (ptr *QPaintDeviceWindow) MouseMoveEvent(ev QMouseEvent_ITF) {
	defer qt.Recovering("QPaintDeviceWindow::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QPaintDeviceWindow_MouseMoveEvent(ptr.Pointer(), PointerFromQMouseEvent(ev))
	}
}

func (ptr *QPaintDeviceWindow) MouseMoveEventDefault(ev QMouseEvent_ITF) {
	defer qt.Recovering("QPaintDeviceWindow::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QPaintDeviceWindow_MouseMoveEventDefault(ptr.Pointer(), PointerFromQMouseEvent(ev))
	}
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
func callbackQPaintDeviceWindowMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QPaintDeviceWindow::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*QMouseEvent))(NewQMouseEventFromPointer(ev))
	} else {
		NewQPaintDeviceWindowFromPointer(ptr).MousePressEventDefault(NewQMouseEventFromPointer(ev))
	}
}

func (ptr *QPaintDeviceWindow) MousePressEvent(ev QMouseEvent_ITF) {
	defer qt.Recovering("QPaintDeviceWindow::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QPaintDeviceWindow_MousePressEvent(ptr.Pointer(), PointerFromQMouseEvent(ev))
	}
}

func (ptr *QPaintDeviceWindow) MousePressEventDefault(ev QMouseEvent_ITF) {
	defer qt.Recovering("QPaintDeviceWindow::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QPaintDeviceWindow_MousePressEventDefault(ptr.Pointer(), PointerFromQMouseEvent(ev))
	}
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
func callbackQPaintDeviceWindowMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QPaintDeviceWindow::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*QMouseEvent))(NewQMouseEventFromPointer(ev))
	} else {
		NewQPaintDeviceWindowFromPointer(ptr).MouseReleaseEventDefault(NewQMouseEventFromPointer(ev))
	}
}

func (ptr *QPaintDeviceWindow) MouseReleaseEvent(ev QMouseEvent_ITF) {
	defer qt.Recovering("QPaintDeviceWindow::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QPaintDeviceWindow_MouseReleaseEvent(ptr.Pointer(), PointerFromQMouseEvent(ev))
	}
}

func (ptr *QPaintDeviceWindow) MouseReleaseEventDefault(ev QMouseEvent_ITF) {
	defer qt.Recovering("QPaintDeviceWindow::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QPaintDeviceWindow_MouseReleaseEventDefault(ptr.Pointer(), PointerFromQMouseEvent(ev))
	}
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
func callbackQPaintDeviceWindowMoveEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QPaintDeviceWindow::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*QMoveEvent))(NewQMoveEventFromPointer(ev))
	} else {
		NewQPaintDeviceWindowFromPointer(ptr).MoveEventDefault(NewQMoveEventFromPointer(ev))
	}
}

func (ptr *QPaintDeviceWindow) MoveEvent(ev QMoveEvent_ITF) {
	defer qt.Recovering("QPaintDeviceWindow::moveEvent")

	if ptr.Pointer() != nil {
		C.QPaintDeviceWindow_MoveEvent(ptr.Pointer(), PointerFromQMoveEvent(ev))
	}
}

func (ptr *QPaintDeviceWindow) MoveEventDefault(ev QMoveEvent_ITF) {
	defer qt.Recovering("QPaintDeviceWindow::moveEvent")

	if ptr.Pointer() != nil {
		C.QPaintDeviceWindow_MoveEventDefault(ptr.Pointer(), PointerFromQMoveEvent(ev))
	}
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
func callbackQPaintDeviceWindowResizeEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QPaintDeviceWindow::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*QResizeEvent))(NewQResizeEventFromPointer(ev))
	} else {
		NewQPaintDeviceWindowFromPointer(ptr).ResizeEventDefault(NewQResizeEventFromPointer(ev))
	}
}

func (ptr *QPaintDeviceWindow) ResizeEvent(ev QResizeEvent_ITF) {
	defer qt.Recovering("QPaintDeviceWindow::resizeEvent")

	if ptr.Pointer() != nil {
		C.QPaintDeviceWindow_ResizeEvent(ptr.Pointer(), PointerFromQResizeEvent(ev))
	}
}

func (ptr *QPaintDeviceWindow) ResizeEventDefault(ev QResizeEvent_ITF) {
	defer qt.Recovering("QPaintDeviceWindow::resizeEvent")

	if ptr.Pointer() != nil {
		C.QPaintDeviceWindow_ResizeEventDefault(ptr.Pointer(), PointerFromQResizeEvent(ev))
	}
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
func callbackQPaintDeviceWindowShowEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QPaintDeviceWindow::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*QShowEvent))(NewQShowEventFromPointer(ev))
	} else {
		NewQPaintDeviceWindowFromPointer(ptr).ShowEventDefault(NewQShowEventFromPointer(ev))
	}
}

func (ptr *QPaintDeviceWindow) ShowEvent(ev QShowEvent_ITF) {
	defer qt.Recovering("QPaintDeviceWindow::showEvent")

	if ptr.Pointer() != nil {
		C.QPaintDeviceWindow_ShowEvent(ptr.Pointer(), PointerFromQShowEvent(ev))
	}
}

func (ptr *QPaintDeviceWindow) ShowEventDefault(ev QShowEvent_ITF) {
	defer qt.Recovering("QPaintDeviceWindow::showEvent")

	if ptr.Pointer() != nil {
		C.QPaintDeviceWindow_ShowEventDefault(ptr.Pointer(), PointerFromQShowEvent(ev))
	}
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
func callbackQPaintDeviceWindowTabletEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QPaintDeviceWindow::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*QTabletEvent))(NewQTabletEventFromPointer(ev))
	} else {
		NewQPaintDeviceWindowFromPointer(ptr).TabletEventDefault(NewQTabletEventFromPointer(ev))
	}
}

func (ptr *QPaintDeviceWindow) TabletEvent(ev QTabletEvent_ITF) {
	defer qt.Recovering("QPaintDeviceWindow::tabletEvent")

	if ptr.Pointer() != nil {
		C.QPaintDeviceWindow_TabletEvent(ptr.Pointer(), PointerFromQTabletEvent(ev))
	}
}

func (ptr *QPaintDeviceWindow) TabletEventDefault(ev QTabletEvent_ITF) {
	defer qt.Recovering("QPaintDeviceWindow::tabletEvent")

	if ptr.Pointer() != nil {
		C.QPaintDeviceWindow_TabletEventDefault(ptr.Pointer(), PointerFromQTabletEvent(ev))
	}
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
func callbackQPaintDeviceWindowTouchEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QPaintDeviceWindow::touchEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "touchEvent"); signal != nil {
		signal.(func(*QTouchEvent))(NewQTouchEventFromPointer(ev))
	} else {
		NewQPaintDeviceWindowFromPointer(ptr).TouchEventDefault(NewQTouchEventFromPointer(ev))
	}
}

func (ptr *QPaintDeviceWindow) TouchEvent(ev QTouchEvent_ITF) {
	defer qt.Recovering("QPaintDeviceWindow::touchEvent")

	if ptr.Pointer() != nil {
		C.QPaintDeviceWindow_TouchEvent(ptr.Pointer(), PointerFromQTouchEvent(ev))
	}
}

func (ptr *QPaintDeviceWindow) TouchEventDefault(ev QTouchEvent_ITF) {
	defer qt.Recovering("QPaintDeviceWindow::touchEvent")

	if ptr.Pointer() != nil {
		C.QPaintDeviceWindow_TouchEventDefault(ptr.Pointer(), PointerFromQTouchEvent(ev))
	}
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
func callbackQPaintDeviceWindowWheelEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QPaintDeviceWindow::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*QWheelEvent))(NewQWheelEventFromPointer(ev))
	} else {
		NewQPaintDeviceWindowFromPointer(ptr).WheelEventDefault(NewQWheelEventFromPointer(ev))
	}
}

func (ptr *QPaintDeviceWindow) WheelEvent(ev QWheelEvent_ITF) {
	defer qt.Recovering("QPaintDeviceWindow::wheelEvent")

	if ptr.Pointer() != nil {
		C.QPaintDeviceWindow_WheelEvent(ptr.Pointer(), PointerFromQWheelEvent(ev))
	}
}

func (ptr *QPaintDeviceWindow) WheelEventDefault(ev QWheelEvent_ITF) {
	defer qt.Recovering("QPaintDeviceWindow::wheelEvent")

	if ptr.Pointer() != nil {
		C.QPaintDeviceWindow_WheelEventDefault(ptr.Pointer(), PointerFromQWheelEvent(ev))
	}
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
func callbackQPaintDeviceWindowTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QPaintDeviceWindow::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQPaintDeviceWindowFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QPaintDeviceWindow) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QPaintDeviceWindow::timerEvent")

	if ptr.Pointer() != nil {
		C.QPaintDeviceWindow_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QPaintDeviceWindow) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QPaintDeviceWindow::timerEvent")

	if ptr.Pointer() != nil {
		C.QPaintDeviceWindow_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
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
func callbackQPaintDeviceWindowChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QPaintDeviceWindow::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQPaintDeviceWindowFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QPaintDeviceWindow) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QPaintDeviceWindow::childEvent")

	if ptr.Pointer() != nil {
		C.QPaintDeviceWindow_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QPaintDeviceWindow) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QPaintDeviceWindow::childEvent")

	if ptr.Pointer() != nil {
		C.QPaintDeviceWindow_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
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
func callbackQPaintDeviceWindowCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QPaintDeviceWindow::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQPaintDeviceWindowFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QPaintDeviceWindow) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QPaintDeviceWindow::customEvent")

	if ptr.Pointer() != nil {
		C.QPaintDeviceWindow_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QPaintDeviceWindow) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QPaintDeviceWindow::customEvent")

	if ptr.Pointer() != nil {
		C.QPaintDeviceWindow_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
