package main

//#include <stdlib.h>
//#include "moc.h"
import "C"
import (
	"encoding/hex"
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"runtime"
	"unsafe"
)

type DropArea_ITF interface {
	widgets.QLabel_ITF
	DropArea_PTR() *DropArea
}

func (p *DropArea) DropArea_PTR() *DropArea {
	return p
}

func (p *DropArea) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QLabel_PTR().Pointer()
	}
	return nil
}

func (p *DropArea) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QLabel_PTR().SetPointer(ptr)
	}
}

func PointerFromDropArea(ptr DropArea_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.DropArea_PTR().Pointer()
	}
	return nil
}

func NewDropAreaFromPointer(ptr unsafe.Pointer) *DropArea {
	var n = new(DropArea)
	n.SetPointer(ptr)
	return n
}

func newDropAreaFromPointer(ptr unsafe.Pointer) *DropArea {
	var n = NewDropAreaFromPointer(ptr)
	for len(n.ObjectName()) < len("DropArea_") {
		n.SetObjectName("DropArea_" + qt.Identifier())
	}
	return n
}

//export callbackDropArea_Changed
func callbackDropArea_Changed(ptr unsafe.Pointer, ptrName *C.char, mimeData unsafe.Pointer) {
	defer qt.Recovering("callback DropArea::changed")

	if signal := qt.GetSignal(C.GoString(ptrName), "changed"); signal != nil {
		signal.(func(*core.QMimeData))(core.NewQMimeDataFromPointer(mimeData))
	}

}

func (ptr *DropArea) ConnectChanged(f func(mimeData *core.QMimeData)) {
	defer qt.Recovering("connect DropArea::changed")

	if ptr.Pointer() != nil {
		C.DropArea_ConnectChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "changed", f)
	}
}

func (ptr *DropArea) DisconnectChanged() {
	defer qt.Recovering("disconnect DropArea::changed")

	if ptr.Pointer() != nil {
		C.DropArea_DisconnectChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "changed")
	}
}

func (ptr *DropArea) Changed(mimeData core.QMimeData_ITF) {
	defer qt.Recovering("DropArea::changed")

	if ptr.Pointer() != nil {
		C.DropArea_Changed(ptr.Pointer(), core.PointerFromQMimeData(mimeData))
	}
}

func NewDropArea(parent widgets.QWidget_ITF, f core.Qt__WindowType) *DropArea {
	defer qt.Recovering("DropArea::DropArea")

	return newDropAreaFromPointer(C.DropArea_NewDropArea(widgets.PointerFromQWidget(parent), C.int(f)))
}

func (ptr *DropArea) DestroyDropArea() {
	defer qt.Recovering("DropArea::~DropArea")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.DropArea_DestroyDropArea(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackDropArea_SetPixmap
func callbackDropArea_SetPixmap(ptr unsafe.Pointer, ptrName *C.char, vqp unsafe.Pointer) {
	defer qt.Recovering("callback DropArea::setPixmap")

	if signal := qt.GetSignal(C.GoString(ptrName), "setPixmap"); signal != nil {
		signal.(func(*gui.QPixmap))(gui.NewQPixmapFromPointer(vqp))
	} else {
		NewDropAreaFromPointer(ptr).SetPixmapDefault(gui.NewQPixmapFromPointer(vqp))
	}
}

func (ptr *DropArea) ConnectSetPixmap(f func(vqp *gui.QPixmap)) {
	defer qt.Recovering("connect DropArea::setPixmap")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setPixmap", f)
	}
}

func (ptr *DropArea) DisconnectSetPixmap() {
	defer qt.Recovering("disconnect DropArea::setPixmap")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setPixmap")
	}
}

func (ptr *DropArea) SetPixmap(vqp gui.QPixmap_ITF) {
	defer qt.Recovering("DropArea::setPixmap")

	if ptr.Pointer() != nil {
		C.DropArea_SetPixmap(ptr.Pointer(), gui.PointerFromQPixmap(vqp))
	}
}

func (ptr *DropArea) SetPixmapDefault(vqp gui.QPixmap_ITF) {
	defer qt.Recovering("DropArea::setPixmap")

	if ptr.Pointer() != nil {
		C.DropArea_SetPixmapDefault(ptr.Pointer(), gui.PointerFromQPixmap(vqp))
	}
}

//export callbackDropArea_SetText
func callbackDropArea_SetText(ptr unsafe.Pointer, ptrName *C.char, vqs *C.char) {
	defer qt.Recovering("callback DropArea::setText")

	if signal := qt.GetSignal(C.GoString(ptrName), "setText"); signal != nil {
		signal.(func(string))(C.GoString(vqs))
	} else {
		NewDropAreaFromPointer(ptr).SetTextDefault(C.GoString(vqs))
	}
}

func (ptr *DropArea) ConnectSetText(f func(vqs string)) {
	defer qt.Recovering("connect DropArea::setText")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setText", f)
	}
}

func (ptr *DropArea) DisconnectSetText() {
	defer qt.Recovering("disconnect DropArea::setText")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setText")
	}
}

func (ptr *DropArea) SetText(vqs string) {
	defer qt.Recovering("DropArea::setText")

	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.DropArea_SetText(ptr.Pointer(), vqsC)
	}
}

func (ptr *DropArea) SetTextDefault(vqs string) {
	defer qt.Recovering("DropArea::setText")

	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.DropArea_SetTextDefault(ptr.Pointer(), vqsC)
	}
}

//export callbackDropArea_ChangeEvent
func callbackDropArea_ChangeEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback DropArea::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
	} else {
		NewDropAreaFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(ev))
	}
}

func (ptr *DropArea) ConnectChangeEvent(f func(ev *core.QEvent)) {
	defer qt.Recovering("connect DropArea::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *DropArea) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect DropArea::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

func (ptr *DropArea) ChangeEvent(ev core.QEvent_ITF) {
	defer qt.Recovering("DropArea::changeEvent")

	if ptr.Pointer() != nil {
		C.DropArea_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *DropArea) ChangeEventDefault(ev core.QEvent_ITF) {
	defer qt.Recovering("DropArea::changeEvent")

	if ptr.Pointer() != nil {
		C.DropArea_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

//export callbackDropArea_Clear
func callbackDropArea_Clear(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback DropArea::clear")

	if signal := qt.GetSignal(C.GoString(ptrName), "clear"); signal != nil {
		signal.(func())()
	} else {
		NewDropAreaFromPointer(ptr).ClearDefault()
	}
}

func (ptr *DropArea) ConnectClear(f func()) {
	defer qt.Recovering("connect DropArea::clear")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "clear", f)
	}
}

func (ptr *DropArea) DisconnectClear() {
	defer qt.Recovering("disconnect DropArea::clear")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "clear")
	}
}

func (ptr *DropArea) Clear() {
	defer qt.Recovering("DropArea::clear")

	if ptr.Pointer() != nil {
		C.DropArea_Clear(ptr.Pointer())
	}
}

func (ptr *DropArea) ClearDefault() {
	defer qt.Recovering("DropArea::clear")

	if ptr.Pointer() != nil {
		C.DropArea_ClearDefault(ptr.Pointer())
	}
}

//export callbackDropArea_ContextMenuEvent
func callbackDropArea_ContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback DropArea::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(ev))
	} else {
		NewDropAreaFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(ev))
	}
}

func (ptr *DropArea) ConnectContextMenuEvent(f func(ev *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect DropArea::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *DropArea) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect DropArea::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

func (ptr *DropArea) ContextMenuEvent(ev gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("DropArea::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.DropArea_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(ev))
	}
}

func (ptr *DropArea) ContextMenuEventDefault(ev gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("DropArea::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.DropArea_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(ev))
	}
}

//export callbackDropArea_FocusInEvent
func callbackDropArea_FocusInEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback DropArea::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(ev))
	} else {
		NewDropAreaFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(ev))
	}
}

func (ptr *DropArea) ConnectFocusInEvent(f func(ev *gui.QFocusEvent)) {
	defer qt.Recovering("connect DropArea::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *DropArea) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect DropArea::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

func (ptr *DropArea) FocusInEvent(ev gui.QFocusEvent_ITF) {
	defer qt.Recovering("DropArea::focusInEvent")

	if ptr.Pointer() != nil {
		C.DropArea_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(ev))
	}
}

func (ptr *DropArea) FocusInEventDefault(ev gui.QFocusEvent_ITF) {
	defer qt.Recovering("DropArea::focusInEvent")

	if ptr.Pointer() != nil {
		C.DropArea_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(ev))
	}
}

//export callbackDropArea_FocusNextPrevChild
func callbackDropArea_FocusNextPrevChild(ptr unsafe.Pointer, ptrName *C.char, next C.int) C.int {
	defer qt.Recovering("callback DropArea::focusNextPrevChild")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusNextPrevChild"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(bool) bool)(int(next) != 0)))
	}

	return C.int(qt.GoBoolToInt(NewDropAreaFromPointer(ptr).FocusNextPrevChildDefault(int(next) != 0)))
}

func (ptr *DropArea) ConnectFocusNextPrevChild(f func(next bool) bool) {
	defer qt.Recovering("connect DropArea::focusNextPrevChild")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusNextPrevChild", f)
	}
}

func (ptr *DropArea) DisconnectFocusNextPrevChild() {
	defer qt.Recovering("disconnect DropArea::focusNextPrevChild")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusNextPrevChild")
	}
}

func (ptr *DropArea) FocusNextPrevChild(next bool) bool {
	defer qt.Recovering("DropArea::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.DropArea_FocusNextPrevChild(ptr.Pointer(), C.int(qt.GoBoolToInt(next))) != 0
	}
	return false
}

func (ptr *DropArea) FocusNextPrevChildDefault(next bool) bool {
	defer qt.Recovering("DropArea::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.DropArea_FocusNextPrevChildDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(next))) != 0
	}
	return false
}

//export callbackDropArea_FocusOutEvent
func callbackDropArea_FocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback DropArea::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(ev))
	} else {
		NewDropAreaFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(ev))
	}
}

func (ptr *DropArea) ConnectFocusOutEvent(f func(ev *gui.QFocusEvent)) {
	defer qt.Recovering("connect DropArea::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *DropArea) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect DropArea::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

func (ptr *DropArea) FocusOutEvent(ev gui.QFocusEvent_ITF) {
	defer qt.Recovering("DropArea::focusOutEvent")

	if ptr.Pointer() != nil {
		C.DropArea_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(ev))
	}
}

func (ptr *DropArea) FocusOutEventDefault(ev gui.QFocusEvent_ITF) {
	defer qt.Recovering("DropArea::focusOutEvent")

	if ptr.Pointer() != nil {
		C.DropArea_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(ev))
	}
}

//export callbackDropArea_HeightForWidth
func callbackDropArea_HeightForWidth(ptr unsafe.Pointer, ptrName *C.char, w C.int) C.int {
	defer qt.Recovering("callback DropArea::heightForWidth")

	if signal := qt.GetSignal(C.GoString(ptrName), "heightForWidth"); signal != nil {
		return C.int(signal.(func(int) int)(int(w)))
	}

	return C.int(NewDropAreaFromPointer(ptr).HeightForWidthDefault(int(w)))
}

func (ptr *DropArea) ConnectHeightForWidth(f func(w int) int) {
	defer qt.Recovering("connect DropArea::heightForWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "heightForWidth", f)
	}
}

func (ptr *DropArea) DisconnectHeightForWidth() {
	defer qt.Recovering("disconnect DropArea::heightForWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "heightForWidth")
	}
}

func (ptr *DropArea) HeightForWidth(w int) int {
	defer qt.Recovering("DropArea::heightForWidth")

	if ptr.Pointer() != nil {
		return int(C.DropArea_HeightForWidth(ptr.Pointer(), C.int(w)))
	}
	return 0
}

func (ptr *DropArea) HeightForWidthDefault(w int) int {
	defer qt.Recovering("DropArea::heightForWidth")

	if ptr.Pointer() != nil {
		return int(C.DropArea_HeightForWidthDefault(ptr.Pointer(), C.int(w)))
	}
	return 0
}

//export callbackDropArea_KeyPressEvent
func callbackDropArea_KeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback DropArea::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(ev))
	} else {
		NewDropAreaFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(ev))
	}
}

func (ptr *DropArea) ConnectKeyPressEvent(f func(ev *gui.QKeyEvent)) {
	defer qt.Recovering("connect DropArea::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *DropArea) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect DropArea::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

func (ptr *DropArea) KeyPressEvent(ev gui.QKeyEvent_ITF) {
	defer qt.Recovering("DropArea::keyPressEvent")

	if ptr.Pointer() != nil {
		C.DropArea_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(ev))
	}
}

func (ptr *DropArea) KeyPressEventDefault(ev gui.QKeyEvent_ITF) {
	defer qt.Recovering("DropArea::keyPressEvent")

	if ptr.Pointer() != nil {
		C.DropArea_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(ev))
	}
}

//export callbackDropArea_MinimumSizeHint
func callbackDropArea_MinimumSizeHint(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback DropArea::minimumSizeHint")

	if signal := qt.GetSignal(C.GoString(ptrName), "minimumSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewDropAreaFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *DropArea) ConnectMinimumSizeHint(f func() *core.QSize) {
	defer qt.Recovering("connect DropArea::minimumSizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "minimumSizeHint", f)
	}
}

func (ptr *DropArea) DisconnectMinimumSizeHint() {
	defer qt.Recovering("disconnect DropArea::minimumSizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "minimumSizeHint")
	}
}

func (ptr *DropArea) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("DropArea::minimumSizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.DropArea_MinimumSizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *DropArea) MinimumSizeHintDefault() *core.QSize {
	defer qt.Recovering("DropArea::minimumSizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.DropArea_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackDropArea_MouseMoveEvent
func callbackDropArea_MouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback DropArea::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(ev))
	} else {
		NewDropAreaFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(ev))
	}
}

func (ptr *DropArea) ConnectMouseMoveEvent(f func(ev *gui.QMouseEvent)) {
	defer qt.Recovering("connect DropArea::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *DropArea) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect DropArea::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

func (ptr *DropArea) MouseMoveEvent(ev gui.QMouseEvent_ITF) {
	defer qt.Recovering("DropArea::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.DropArea_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(ev))
	}
}

func (ptr *DropArea) MouseMoveEventDefault(ev gui.QMouseEvent_ITF) {
	defer qt.Recovering("DropArea::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.DropArea_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(ev))
	}
}

//export callbackDropArea_MousePressEvent
func callbackDropArea_MousePressEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback DropArea::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(ev))
	} else {
		NewDropAreaFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(ev))
	}
}

func (ptr *DropArea) ConnectMousePressEvent(f func(ev *gui.QMouseEvent)) {
	defer qt.Recovering("connect DropArea::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *DropArea) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect DropArea::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

func (ptr *DropArea) MousePressEvent(ev gui.QMouseEvent_ITF) {
	defer qt.Recovering("DropArea::mousePressEvent")

	if ptr.Pointer() != nil {
		C.DropArea_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(ev))
	}
}

func (ptr *DropArea) MousePressEventDefault(ev gui.QMouseEvent_ITF) {
	defer qt.Recovering("DropArea::mousePressEvent")

	if ptr.Pointer() != nil {
		C.DropArea_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(ev))
	}
}

//export callbackDropArea_MouseReleaseEvent
func callbackDropArea_MouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback DropArea::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(ev))
	} else {
		NewDropAreaFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(ev))
	}
}

func (ptr *DropArea) ConnectMouseReleaseEvent(f func(ev *gui.QMouseEvent)) {
	defer qt.Recovering("connect DropArea::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *DropArea) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect DropArea::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

func (ptr *DropArea) MouseReleaseEvent(ev gui.QMouseEvent_ITF) {
	defer qt.Recovering("DropArea::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.DropArea_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(ev))
	}
}

func (ptr *DropArea) MouseReleaseEventDefault(ev gui.QMouseEvent_ITF) {
	defer qt.Recovering("DropArea::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.DropArea_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(ev))
	}
}

//export callbackDropArea_PaintEvent
func callbackDropArea_PaintEvent(ptr unsafe.Pointer, ptrName *C.char, vqp unsafe.Pointer) {
	defer qt.Recovering("callback DropArea::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(vqp))
	} else {
		NewDropAreaFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(vqp))
	}
}

func (ptr *DropArea) ConnectPaintEvent(f func(vqp *gui.QPaintEvent)) {
	defer qt.Recovering("connect DropArea::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *DropArea) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect DropArea::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

func (ptr *DropArea) PaintEvent(vqp gui.QPaintEvent_ITF) {
	defer qt.Recovering("DropArea::paintEvent")

	if ptr.Pointer() != nil {
		C.DropArea_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(vqp))
	}
}

func (ptr *DropArea) PaintEventDefault(vqp gui.QPaintEvent_ITF) {
	defer qt.Recovering("DropArea::paintEvent")

	if ptr.Pointer() != nil {
		C.DropArea_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(vqp))
	}
}

//export callbackDropArea_SetMovie
func callbackDropArea_SetMovie(ptr unsafe.Pointer, ptrName *C.char, movie unsafe.Pointer) {
	defer qt.Recovering("callback DropArea::setMovie")

	if signal := qt.GetSignal(C.GoString(ptrName), "setMovie"); signal != nil {
		signal.(func(*gui.QMovie))(gui.NewQMovieFromPointer(movie))
	} else {
		NewDropAreaFromPointer(ptr).SetMovieDefault(gui.NewQMovieFromPointer(movie))
	}
}

func (ptr *DropArea) ConnectSetMovie(f func(movie *gui.QMovie)) {
	defer qt.Recovering("connect DropArea::setMovie")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setMovie", f)
	}
}

func (ptr *DropArea) DisconnectSetMovie() {
	defer qt.Recovering("disconnect DropArea::setMovie")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setMovie")
	}
}

func (ptr *DropArea) SetMovie(movie gui.QMovie_ITF) {
	defer qt.Recovering("DropArea::setMovie")

	if ptr.Pointer() != nil {
		C.DropArea_SetMovie(ptr.Pointer(), gui.PointerFromQMovie(movie))
	}
}

func (ptr *DropArea) SetMovieDefault(movie gui.QMovie_ITF) {
	defer qt.Recovering("DropArea::setMovie")

	if ptr.Pointer() != nil {
		C.DropArea_SetMovieDefault(ptr.Pointer(), gui.PointerFromQMovie(movie))
	}
}

//export callbackDropArea_SetNum2

func (ptr *DropArea) DisconnectSetNum2() {
	defer qt.Recovering("disconnect DropArea::setNum")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setNum2")
	}
}

//export callbackDropArea_SetNum
func callbackDropArea_SetNum(ptr unsafe.Pointer, ptrName *C.char, num C.int) {
	defer qt.Recovering("callback DropArea::setNum")

	if signal := qt.GetSignal(C.GoString(ptrName), "setNum"); signal != nil {
		signal.(func(int))(int(num))
	} else {
		NewDropAreaFromPointer(ptr).SetNumDefault(int(num))
	}
}

func (ptr *DropArea) ConnectSetNum(f func(num int)) {
	defer qt.Recovering("connect DropArea::setNum")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setNum", f)
	}
}

func (ptr *DropArea) DisconnectSetNum() {
	defer qt.Recovering("disconnect DropArea::setNum")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setNum")
	}
}

func (ptr *DropArea) SetNum(num int) {
	defer qt.Recovering("DropArea::setNum")

	if ptr.Pointer() != nil {
		C.DropArea_SetNum(ptr.Pointer(), C.int(num))
	}
}

func (ptr *DropArea) SetNumDefault(num int) {
	defer qt.Recovering("DropArea::setNum")

	if ptr.Pointer() != nil {
		C.DropArea_SetNumDefault(ptr.Pointer(), C.int(num))
	}
}

//export callbackDropArea_SetPicture
func callbackDropArea_SetPicture(ptr unsafe.Pointer, ptrName *C.char, picture unsafe.Pointer) {
	defer qt.Recovering("callback DropArea::setPicture")

	if signal := qt.GetSignal(C.GoString(ptrName), "setPicture"); signal != nil {
		signal.(func(*gui.QPicture))(gui.NewQPictureFromPointer(picture))
	} else {
		NewDropAreaFromPointer(ptr).SetPictureDefault(gui.NewQPictureFromPointer(picture))
	}
}

func (ptr *DropArea) ConnectSetPicture(f func(picture *gui.QPicture)) {
	defer qt.Recovering("connect DropArea::setPicture")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setPicture", f)
	}
}

func (ptr *DropArea) DisconnectSetPicture() {
	defer qt.Recovering("disconnect DropArea::setPicture")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setPicture")
	}
}

func (ptr *DropArea) SetPicture(picture gui.QPicture_ITF) {
	defer qt.Recovering("DropArea::setPicture")

	if ptr.Pointer() != nil {
		C.DropArea_SetPicture(ptr.Pointer(), gui.PointerFromQPicture(picture))
	}
}

func (ptr *DropArea) SetPictureDefault(picture gui.QPicture_ITF) {
	defer qt.Recovering("DropArea::setPicture")

	if ptr.Pointer() != nil {
		C.DropArea_SetPictureDefault(ptr.Pointer(), gui.PointerFromQPicture(picture))
	}
}

//export callbackDropArea_SizeHint
func callbackDropArea_SizeHint(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback DropArea::sizeHint")

	if signal := qt.GetSignal(C.GoString(ptrName), "sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewDropAreaFromPointer(ptr).SizeHintDefault())
}

func (ptr *DropArea) ConnectSizeHint(f func() *core.QSize) {
	defer qt.Recovering("connect DropArea::sizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "sizeHint", f)
	}
}

func (ptr *DropArea) DisconnectSizeHint() {
	defer qt.Recovering("disconnect DropArea::sizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "sizeHint")
	}
}

func (ptr *DropArea) SizeHint() *core.QSize {
	defer qt.Recovering("DropArea::sizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.DropArea_SizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *DropArea) SizeHintDefault() *core.QSize {
	defer qt.Recovering("DropArea::sizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.DropArea_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackDropArea_ActionEvent
func callbackDropArea_ActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback DropArea::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewDropAreaFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *DropArea) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect DropArea::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *DropArea) DisconnectActionEvent() {
	defer qt.Recovering("disconnect DropArea::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

func (ptr *DropArea) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("DropArea::actionEvent")

	if ptr.Pointer() != nil {
		C.DropArea_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *DropArea) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("DropArea::actionEvent")

	if ptr.Pointer() != nil {
		C.DropArea_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackDropArea_DragEnterEvent
func callbackDropArea_DragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback DropArea::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewDropAreaFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *DropArea) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect DropArea::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *DropArea) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect DropArea::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

func (ptr *DropArea) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("DropArea::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.DropArea_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *DropArea) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("DropArea::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.DropArea_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackDropArea_DragLeaveEvent
func callbackDropArea_DragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback DropArea::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewDropAreaFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *DropArea) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect DropArea::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *DropArea) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect DropArea::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

func (ptr *DropArea) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("DropArea::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.DropArea_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *DropArea) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("DropArea::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.DropArea_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackDropArea_DragMoveEvent
func callbackDropArea_DragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback DropArea::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewDropAreaFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *DropArea) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect DropArea::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *DropArea) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect DropArea::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

func (ptr *DropArea) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("DropArea::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.DropArea_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *DropArea) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("DropArea::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.DropArea_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackDropArea_DropEvent
func callbackDropArea_DropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback DropArea::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewDropAreaFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *DropArea) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect DropArea::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *DropArea) DisconnectDropEvent() {
	defer qt.Recovering("disconnect DropArea::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

func (ptr *DropArea) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("DropArea::dropEvent")

	if ptr.Pointer() != nil {
		C.DropArea_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *DropArea) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("DropArea::dropEvent")

	if ptr.Pointer() != nil {
		C.DropArea_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

//export callbackDropArea_EnterEvent
func callbackDropArea_EnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback DropArea::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewDropAreaFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *DropArea) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect DropArea::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *DropArea) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect DropArea::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

func (ptr *DropArea) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("DropArea::enterEvent")

	if ptr.Pointer() != nil {
		C.DropArea_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *DropArea) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("DropArea::enterEvent")

	if ptr.Pointer() != nil {
		C.DropArea_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackDropArea_HideEvent
func callbackDropArea_HideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback DropArea::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewDropAreaFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *DropArea) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect DropArea::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *DropArea) DisconnectHideEvent() {
	defer qt.Recovering("disconnect DropArea::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

func (ptr *DropArea) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("DropArea::hideEvent")

	if ptr.Pointer() != nil {
		C.DropArea_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *DropArea) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("DropArea::hideEvent")

	if ptr.Pointer() != nil {
		C.DropArea_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackDropArea_LeaveEvent
func callbackDropArea_LeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback DropArea::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewDropAreaFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *DropArea) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect DropArea::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *DropArea) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect DropArea::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

func (ptr *DropArea) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("DropArea::leaveEvent")

	if ptr.Pointer() != nil {
		C.DropArea_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *DropArea) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("DropArea::leaveEvent")

	if ptr.Pointer() != nil {
		C.DropArea_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackDropArea_MoveEvent
func callbackDropArea_MoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback DropArea::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewDropAreaFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *DropArea) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect DropArea::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *DropArea) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect DropArea::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

func (ptr *DropArea) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("DropArea::moveEvent")

	if ptr.Pointer() != nil {
		C.DropArea_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *DropArea) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("DropArea::moveEvent")

	if ptr.Pointer() != nil {
		C.DropArea_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackDropArea_SetEnabled
func callbackDropArea_SetEnabled(ptr unsafe.Pointer, ptrName *C.char, vbo C.int) {
	defer qt.Recovering("callback DropArea::setEnabled")

	if signal := qt.GetSignal(C.GoString(ptrName), "setEnabled"); signal != nil {
		signal.(func(bool))(int(vbo) != 0)
	} else {
		NewDropAreaFromPointer(ptr).SetEnabledDefault(int(vbo) != 0)
	}
}

func (ptr *DropArea) ConnectSetEnabled(f func(vbo bool)) {
	defer qt.Recovering("connect DropArea::setEnabled")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setEnabled", f)
	}
}

func (ptr *DropArea) DisconnectSetEnabled() {
	defer qt.Recovering("disconnect DropArea::setEnabled")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setEnabled")
	}
}

func (ptr *DropArea) SetEnabled(vbo bool) {
	defer qt.Recovering("DropArea::setEnabled")

	if ptr.Pointer() != nil {
		C.DropArea_SetEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(vbo)))
	}
}

func (ptr *DropArea) SetEnabledDefault(vbo bool) {
	defer qt.Recovering("DropArea::setEnabled")

	if ptr.Pointer() != nil {
		C.DropArea_SetEnabledDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(vbo)))
	}
}

//export callbackDropArea_SetStyleSheet
func callbackDropArea_SetStyleSheet(ptr unsafe.Pointer, ptrName *C.char, styleSheet *C.char) {
	defer qt.Recovering("callback DropArea::setStyleSheet")

	if signal := qt.GetSignal(C.GoString(ptrName), "setStyleSheet"); signal != nil {
		signal.(func(string))(C.GoString(styleSheet))
	} else {
		NewDropAreaFromPointer(ptr).SetStyleSheetDefault(C.GoString(styleSheet))
	}
}

func (ptr *DropArea) ConnectSetStyleSheet(f func(styleSheet string)) {
	defer qt.Recovering("connect DropArea::setStyleSheet")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setStyleSheet", f)
	}
}

func (ptr *DropArea) DisconnectSetStyleSheet() {
	defer qt.Recovering("disconnect DropArea::setStyleSheet")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setStyleSheet")
	}
}

func (ptr *DropArea) SetStyleSheet(styleSheet string) {
	defer qt.Recovering("DropArea::setStyleSheet")

	if ptr.Pointer() != nil {
		var styleSheetC = C.CString(styleSheet)
		defer C.free(unsafe.Pointer(styleSheetC))
		C.DropArea_SetStyleSheet(ptr.Pointer(), styleSheetC)
	}
}

func (ptr *DropArea) SetStyleSheetDefault(styleSheet string) {
	defer qt.Recovering("DropArea::setStyleSheet")

	if ptr.Pointer() != nil {
		var styleSheetC = C.CString(styleSheet)
		defer C.free(unsafe.Pointer(styleSheetC))
		C.DropArea_SetStyleSheetDefault(ptr.Pointer(), styleSheetC)
	}
}

//export callbackDropArea_SetVisible
func callbackDropArea_SetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) {
	defer qt.Recovering("callback DropArea::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
	} else {
		NewDropAreaFromPointer(ptr).SetVisibleDefault(int(visible) != 0)
	}
}

func (ptr *DropArea) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect DropArea::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *DropArea) DisconnectSetVisible() {
	defer qt.Recovering("disconnect DropArea::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

func (ptr *DropArea) SetVisible(visible bool) {
	defer qt.Recovering("DropArea::setVisible")

	if ptr.Pointer() != nil {
		C.DropArea_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *DropArea) SetVisibleDefault(visible bool) {
	defer qt.Recovering("DropArea::setVisible")

	if ptr.Pointer() != nil {
		C.DropArea_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

//export callbackDropArea_SetWindowModified
func callbackDropArea_SetWindowModified(ptr unsafe.Pointer, ptrName *C.char, vbo C.int) {
	defer qt.Recovering("callback DropArea::setWindowModified")

	if signal := qt.GetSignal(C.GoString(ptrName), "setWindowModified"); signal != nil {
		signal.(func(bool))(int(vbo) != 0)
	} else {
		NewDropAreaFromPointer(ptr).SetWindowModifiedDefault(int(vbo) != 0)
	}
}

func (ptr *DropArea) ConnectSetWindowModified(f func(vbo bool)) {
	defer qt.Recovering("connect DropArea::setWindowModified")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setWindowModified", f)
	}
}

func (ptr *DropArea) DisconnectSetWindowModified() {
	defer qt.Recovering("disconnect DropArea::setWindowModified")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setWindowModified")
	}
}

func (ptr *DropArea) SetWindowModified(vbo bool) {
	defer qt.Recovering("DropArea::setWindowModified")

	if ptr.Pointer() != nil {
		C.DropArea_SetWindowModified(ptr.Pointer(), C.int(qt.GoBoolToInt(vbo)))
	}
}

func (ptr *DropArea) SetWindowModifiedDefault(vbo bool) {
	defer qt.Recovering("DropArea::setWindowModified")

	if ptr.Pointer() != nil {
		C.DropArea_SetWindowModifiedDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(vbo)))
	}
}

//export callbackDropArea_SetWindowTitle
func callbackDropArea_SetWindowTitle(ptr unsafe.Pointer, ptrName *C.char, vqs *C.char) {
	defer qt.Recovering("callback DropArea::setWindowTitle")

	if signal := qt.GetSignal(C.GoString(ptrName), "setWindowTitle"); signal != nil {
		signal.(func(string))(C.GoString(vqs))
	} else {
		NewDropAreaFromPointer(ptr).SetWindowTitleDefault(C.GoString(vqs))
	}
}

func (ptr *DropArea) ConnectSetWindowTitle(f func(vqs string)) {
	defer qt.Recovering("connect DropArea::setWindowTitle")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setWindowTitle", f)
	}
}

func (ptr *DropArea) DisconnectSetWindowTitle() {
	defer qt.Recovering("disconnect DropArea::setWindowTitle")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setWindowTitle")
	}
}

func (ptr *DropArea) SetWindowTitle(vqs string) {
	defer qt.Recovering("DropArea::setWindowTitle")

	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.DropArea_SetWindowTitle(ptr.Pointer(), vqsC)
	}
}

func (ptr *DropArea) SetWindowTitleDefault(vqs string) {
	defer qt.Recovering("DropArea::setWindowTitle")

	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.DropArea_SetWindowTitleDefault(ptr.Pointer(), vqsC)
	}
}

//export callbackDropArea_ShowEvent
func callbackDropArea_ShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback DropArea::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewDropAreaFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *DropArea) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect DropArea::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *DropArea) DisconnectShowEvent() {
	defer qt.Recovering("disconnect DropArea::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

func (ptr *DropArea) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("DropArea::showEvent")

	if ptr.Pointer() != nil {
		C.DropArea_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *DropArea) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("DropArea::showEvent")

	if ptr.Pointer() != nil {
		C.DropArea_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackDropArea_Close
func callbackDropArea_Close(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback DropArea::close")

	if signal := qt.GetSignal(C.GoString(ptrName), "close"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewDropAreaFromPointer(ptr).CloseDefault()))
}

func (ptr *DropArea) ConnectClose(f func() bool) {
	defer qt.Recovering("connect DropArea::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "close", f)
	}
}

func (ptr *DropArea) DisconnectClose() {
	defer qt.Recovering("disconnect DropArea::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "close")
	}
}

func (ptr *DropArea) Close() bool {
	defer qt.Recovering("DropArea::close")

	if ptr.Pointer() != nil {
		return C.DropArea_Close(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *DropArea) CloseDefault() bool {
	defer qt.Recovering("DropArea::close")

	if ptr.Pointer() != nil {
		return C.DropArea_CloseDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackDropArea_CloseEvent
func callbackDropArea_CloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback DropArea::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewDropAreaFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *DropArea) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect DropArea::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *DropArea) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect DropArea::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

func (ptr *DropArea) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("DropArea::closeEvent")

	if ptr.Pointer() != nil {
		C.DropArea_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *DropArea) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("DropArea::closeEvent")

	if ptr.Pointer() != nil {
		C.DropArea_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

//export callbackDropArea_HasHeightForWidth
func callbackDropArea_HasHeightForWidth(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback DropArea::hasHeightForWidth")

	if signal := qt.GetSignal(C.GoString(ptrName), "hasHeightForWidth"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewDropAreaFromPointer(ptr).HasHeightForWidthDefault()))
}

func (ptr *DropArea) ConnectHasHeightForWidth(f func() bool) {
	defer qt.Recovering("connect DropArea::hasHeightForWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hasHeightForWidth", f)
	}
}

func (ptr *DropArea) DisconnectHasHeightForWidth() {
	defer qt.Recovering("disconnect DropArea::hasHeightForWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hasHeightForWidth")
	}
}

func (ptr *DropArea) HasHeightForWidth() bool {
	defer qt.Recovering("DropArea::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.DropArea_HasHeightForWidth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *DropArea) HasHeightForWidthDefault() bool {
	defer qt.Recovering("DropArea::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.DropArea_HasHeightForWidthDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackDropArea_Hide
func callbackDropArea_Hide(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback DropArea::hide")

	if signal := qt.GetSignal(C.GoString(ptrName), "hide"); signal != nil {
		signal.(func())()
	} else {
		NewDropAreaFromPointer(ptr).HideDefault()
	}
}

func (ptr *DropArea) ConnectHide(f func()) {
	defer qt.Recovering("connect DropArea::hide")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hide", f)
	}
}

func (ptr *DropArea) DisconnectHide() {
	defer qt.Recovering("disconnect DropArea::hide")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hide")
	}
}

func (ptr *DropArea) Hide() {
	defer qt.Recovering("DropArea::hide")

	if ptr.Pointer() != nil {
		C.DropArea_Hide(ptr.Pointer())
	}
}

func (ptr *DropArea) HideDefault() {
	defer qt.Recovering("DropArea::hide")

	if ptr.Pointer() != nil {
		C.DropArea_HideDefault(ptr.Pointer())
	}
}

//export callbackDropArea_InputMethodEvent
func callbackDropArea_InputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback DropArea::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewDropAreaFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *DropArea) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect DropArea::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *DropArea) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect DropArea::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

func (ptr *DropArea) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("DropArea::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.DropArea_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *DropArea) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("DropArea::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.DropArea_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackDropArea_InputMethodQuery
func callbackDropArea_InputMethodQuery(ptr unsafe.Pointer, ptrName *C.char, query C.int) unsafe.Pointer {
	defer qt.Recovering("callback DropArea::inputMethodQuery")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewDropAreaFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *DropArea) ConnectInputMethodQuery(f func(query core.Qt__InputMethodQuery) *core.QVariant) {
	defer qt.Recovering("connect DropArea::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodQuery", f)
	}
}

func (ptr *DropArea) DisconnectInputMethodQuery() {
	defer qt.Recovering("disconnect DropArea::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodQuery")
	}
}

func (ptr *DropArea) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("DropArea::inputMethodQuery")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.DropArea_InputMethodQuery(ptr.Pointer(), C.int(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *DropArea) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("DropArea::inputMethodQuery")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.DropArea_InputMethodQueryDefault(ptr.Pointer(), C.int(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackDropArea_KeyReleaseEvent
func callbackDropArea_KeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback DropArea::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewDropAreaFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *DropArea) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect DropArea::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *DropArea) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect DropArea::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

func (ptr *DropArea) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("DropArea::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.DropArea_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *DropArea) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("DropArea::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.DropArea_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackDropArea_Lower
func callbackDropArea_Lower(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback DropArea::lower")

	if signal := qt.GetSignal(C.GoString(ptrName), "lower"); signal != nil {
		signal.(func())()
	} else {
		NewDropAreaFromPointer(ptr).LowerDefault()
	}
}

func (ptr *DropArea) ConnectLower(f func()) {
	defer qt.Recovering("connect DropArea::lower")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "lower", f)
	}
}

func (ptr *DropArea) DisconnectLower() {
	defer qt.Recovering("disconnect DropArea::lower")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "lower")
	}
}

func (ptr *DropArea) Lower() {
	defer qt.Recovering("DropArea::lower")

	if ptr.Pointer() != nil {
		C.DropArea_Lower(ptr.Pointer())
	}
}

func (ptr *DropArea) LowerDefault() {
	defer qt.Recovering("DropArea::lower")

	if ptr.Pointer() != nil {
		C.DropArea_LowerDefault(ptr.Pointer())
	}
}

//export callbackDropArea_MouseDoubleClickEvent
func callbackDropArea_MouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback DropArea::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewDropAreaFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *DropArea) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect DropArea::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *DropArea) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect DropArea::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

func (ptr *DropArea) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("DropArea::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.DropArea_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *DropArea) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("DropArea::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.DropArea_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackDropArea_NativeEvent
func callbackDropArea_NativeEvent(ptr unsafe.Pointer, ptrName *C.char, eventType *C.char, message unsafe.Pointer, result C.long) C.int {
	defer qt.Recovering("callback DropArea::nativeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "nativeEvent"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string, unsafe.Pointer, int) bool)(qt.HexDecodeToString(C.GoString(eventType)), message, int(result))))
	}

	return C.int(qt.GoBoolToInt(NewDropAreaFromPointer(ptr).NativeEventDefault(qt.HexDecodeToString(C.GoString(eventType)), message, int(result))))
}

func (ptr *DropArea) ConnectNativeEvent(f func(eventType string, message unsafe.Pointer, result int) bool) {
	defer qt.Recovering("connect DropArea::nativeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "nativeEvent", f)
	}
}

func (ptr *DropArea) DisconnectNativeEvent() {
	defer qt.Recovering("disconnect DropArea::nativeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "nativeEvent")
	}
}

func (ptr *DropArea) NativeEvent(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("DropArea::nativeEvent")

	if ptr.Pointer() != nil {
		var eventTypeC = C.CString(hex.EncodeToString([]byte(eventType)))
		defer C.free(unsafe.Pointer(eventTypeC))
		return C.DropArea_NativeEvent(ptr.Pointer(), eventTypeC, message, C.long(result)) != 0
	}
	return false
}

func (ptr *DropArea) NativeEventDefault(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("DropArea::nativeEvent")

	if ptr.Pointer() != nil {
		var eventTypeC = C.CString(hex.EncodeToString([]byte(eventType)))
		defer C.free(unsafe.Pointer(eventTypeC))
		return C.DropArea_NativeEventDefault(ptr.Pointer(), eventTypeC, message, C.long(result)) != 0
	}
	return false
}

//export callbackDropArea_Raise
func callbackDropArea_Raise(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback DropArea::raise")

	if signal := qt.GetSignal(C.GoString(ptrName), "raise"); signal != nil {
		signal.(func())()
	} else {
		NewDropAreaFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *DropArea) ConnectRaise(f func()) {
	defer qt.Recovering("connect DropArea::raise")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "raise", f)
	}
}

func (ptr *DropArea) DisconnectRaise() {
	defer qt.Recovering("disconnect DropArea::raise")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "raise")
	}
}

func (ptr *DropArea) Raise() {
	defer qt.Recovering("DropArea::raise")

	if ptr.Pointer() != nil {
		C.DropArea_Raise(ptr.Pointer())
	}
}

func (ptr *DropArea) RaiseDefault() {
	defer qt.Recovering("DropArea::raise")

	if ptr.Pointer() != nil {
		C.DropArea_RaiseDefault(ptr.Pointer())
	}
}

//export callbackDropArea_Repaint
func callbackDropArea_Repaint(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback DropArea::repaint")

	if signal := qt.GetSignal(C.GoString(ptrName), "repaint"); signal != nil {
		signal.(func())()
	} else {
		NewDropAreaFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *DropArea) ConnectRepaint(f func()) {
	defer qt.Recovering("connect DropArea::repaint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "repaint", f)
	}
}

func (ptr *DropArea) DisconnectRepaint() {
	defer qt.Recovering("disconnect DropArea::repaint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "repaint")
	}
}

func (ptr *DropArea) Repaint() {
	defer qt.Recovering("DropArea::repaint")

	if ptr.Pointer() != nil {
		C.DropArea_Repaint(ptr.Pointer())
	}
}

func (ptr *DropArea) RepaintDefault() {
	defer qt.Recovering("DropArea::repaint")

	if ptr.Pointer() != nil {
		C.DropArea_RepaintDefault(ptr.Pointer())
	}
}

//export callbackDropArea_ResizeEvent
func callbackDropArea_ResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback DropArea::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewDropAreaFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *DropArea) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect DropArea::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *DropArea) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect DropArea::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

func (ptr *DropArea) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("DropArea::resizeEvent")

	if ptr.Pointer() != nil {
		C.DropArea_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *DropArea) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("DropArea::resizeEvent")

	if ptr.Pointer() != nil {
		C.DropArea_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

//export callbackDropArea_SetDisabled
func callbackDropArea_SetDisabled(ptr unsafe.Pointer, ptrName *C.char, disable C.int) {
	defer qt.Recovering("callback DropArea::setDisabled")

	if signal := qt.GetSignal(C.GoString(ptrName), "setDisabled"); signal != nil {
		signal.(func(bool))(int(disable) != 0)
	} else {
		NewDropAreaFromPointer(ptr).SetDisabledDefault(int(disable) != 0)
	}
}

func (ptr *DropArea) ConnectSetDisabled(f func(disable bool)) {
	defer qt.Recovering("connect DropArea::setDisabled")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setDisabled", f)
	}
}

func (ptr *DropArea) DisconnectSetDisabled() {
	defer qt.Recovering("disconnect DropArea::setDisabled")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setDisabled")
	}
}

func (ptr *DropArea) SetDisabled(disable bool) {
	defer qt.Recovering("DropArea::setDisabled")

	if ptr.Pointer() != nil {
		C.DropArea_SetDisabled(ptr.Pointer(), C.int(qt.GoBoolToInt(disable)))
	}
}

func (ptr *DropArea) SetDisabledDefault(disable bool) {
	defer qt.Recovering("DropArea::setDisabled")

	if ptr.Pointer() != nil {
		C.DropArea_SetDisabledDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(disable)))
	}
}

//export callbackDropArea_SetFocus2
func callbackDropArea_SetFocus2(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback DropArea::setFocus")

	if signal := qt.GetSignal(C.GoString(ptrName), "setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewDropAreaFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *DropArea) ConnectSetFocus2(f func()) {
	defer qt.Recovering("connect DropArea::setFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setFocus2", f)
	}
}

func (ptr *DropArea) DisconnectSetFocus2() {
	defer qt.Recovering("disconnect DropArea::setFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setFocus2")
	}
}

func (ptr *DropArea) SetFocus2() {
	defer qt.Recovering("DropArea::setFocus")

	if ptr.Pointer() != nil {
		C.DropArea_SetFocus2(ptr.Pointer())
	}
}

func (ptr *DropArea) SetFocus2Default() {
	defer qt.Recovering("DropArea::setFocus")

	if ptr.Pointer() != nil {
		C.DropArea_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackDropArea_SetHidden
func callbackDropArea_SetHidden(ptr unsafe.Pointer, ptrName *C.char, hidden C.int) {
	defer qt.Recovering("callback DropArea::setHidden")

	if signal := qt.GetSignal(C.GoString(ptrName), "setHidden"); signal != nil {
		signal.(func(bool))(int(hidden) != 0)
	} else {
		NewDropAreaFromPointer(ptr).SetHiddenDefault(int(hidden) != 0)
	}
}

func (ptr *DropArea) ConnectSetHidden(f func(hidden bool)) {
	defer qt.Recovering("connect DropArea::setHidden")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setHidden", f)
	}
}

func (ptr *DropArea) DisconnectSetHidden() {
	defer qt.Recovering("disconnect DropArea::setHidden")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setHidden")
	}
}

func (ptr *DropArea) SetHidden(hidden bool) {
	defer qt.Recovering("DropArea::setHidden")

	if ptr.Pointer() != nil {
		C.DropArea_SetHidden(ptr.Pointer(), C.int(qt.GoBoolToInt(hidden)))
	}
}

func (ptr *DropArea) SetHiddenDefault(hidden bool) {
	defer qt.Recovering("DropArea::setHidden")

	if ptr.Pointer() != nil {
		C.DropArea_SetHiddenDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(hidden)))
	}
}

//export callbackDropArea_Show
func callbackDropArea_Show(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback DropArea::show")

	if signal := qt.GetSignal(C.GoString(ptrName), "show"); signal != nil {
		signal.(func())()
	} else {
		NewDropAreaFromPointer(ptr).ShowDefault()
	}
}

func (ptr *DropArea) ConnectShow(f func()) {
	defer qt.Recovering("connect DropArea::show")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "show", f)
	}
}

func (ptr *DropArea) DisconnectShow() {
	defer qt.Recovering("disconnect DropArea::show")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "show")
	}
}

func (ptr *DropArea) Show() {
	defer qt.Recovering("DropArea::show")

	if ptr.Pointer() != nil {
		C.DropArea_Show(ptr.Pointer())
	}
}

func (ptr *DropArea) ShowDefault() {
	defer qt.Recovering("DropArea::show")

	if ptr.Pointer() != nil {
		C.DropArea_ShowDefault(ptr.Pointer())
	}
}

//export callbackDropArea_ShowFullScreen
func callbackDropArea_ShowFullScreen(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback DropArea::showFullScreen")

	if signal := qt.GetSignal(C.GoString(ptrName), "showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewDropAreaFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *DropArea) ConnectShowFullScreen(f func()) {
	defer qt.Recovering("connect DropArea::showFullScreen")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showFullScreen", f)
	}
}

func (ptr *DropArea) DisconnectShowFullScreen() {
	defer qt.Recovering("disconnect DropArea::showFullScreen")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showFullScreen")
	}
}

func (ptr *DropArea) ShowFullScreen() {
	defer qt.Recovering("DropArea::showFullScreen")

	if ptr.Pointer() != nil {
		C.DropArea_ShowFullScreen(ptr.Pointer())
	}
}

func (ptr *DropArea) ShowFullScreenDefault() {
	defer qt.Recovering("DropArea::showFullScreen")

	if ptr.Pointer() != nil {
		C.DropArea_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackDropArea_ShowMaximized
func callbackDropArea_ShowMaximized(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback DropArea::showMaximized")

	if signal := qt.GetSignal(C.GoString(ptrName), "showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewDropAreaFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *DropArea) ConnectShowMaximized(f func()) {
	defer qt.Recovering("connect DropArea::showMaximized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showMaximized", f)
	}
}

func (ptr *DropArea) DisconnectShowMaximized() {
	defer qt.Recovering("disconnect DropArea::showMaximized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showMaximized")
	}
}

func (ptr *DropArea) ShowMaximized() {
	defer qt.Recovering("DropArea::showMaximized")

	if ptr.Pointer() != nil {
		C.DropArea_ShowMaximized(ptr.Pointer())
	}
}

func (ptr *DropArea) ShowMaximizedDefault() {
	defer qt.Recovering("DropArea::showMaximized")

	if ptr.Pointer() != nil {
		C.DropArea_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackDropArea_ShowMinimized
func callbackDropArea_ShowMinimized(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback DropArea::showMinimized")

	if signal := qt.GetSignal(C.GoString(ptrName), "showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewDropAreaFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *DropArea) ConnectShowMinimized(f func()) {
	defer qt.Recovering("connect DropArea::showMinimized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showMinimized", f)
	}
}

func (ptr *DropArea) DisconnectShowMinimized() {
	defer qt.Recovering("disconnect DropArea::showMinimized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showMinimized")
	}
}

func (ptr *DropArea) ShowMinimized() {
	defer qt.Recovering("DropArea::showMinimized")

	if ptr.Pointer() != nil {
		C.DropArea_ShowMinimized(ptr.Pointer())
	}
}

func (ptr *DropArea) ShowMinimizedDefault() {
	defer qt.Recovering("DropArea::showMinimized")

	if ptr.Pointer() != nil {
		C.DropArea_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackDropArea_ShowNormal
func callbackDropArea_ShowNormal(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback DropArea::showNormal")

	if signal := qt.GetSignal(C.GoString(ptrName), "showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewDropAreaFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *DropArea) ConnectShowNormal(f func()) {
	defer qt.Recovering("connect DropArea::showNormal")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showNormal", f)
	}
}

func (ptr *DropArea) DisconnectShowNormal() {
	defer qt.Recovering("disconnect DropArea::showNormal")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showNormal")
	}
}

func (ptr *DropArea) ShowNormal() {
	defer qt.Recovering("DropArea::showNormal")

	if ptr.Pointer() != nil {
		C.DropArea_ShowNormal(ptr.Pointer())
	}
}

func (ptr *DropArea) ShowNormalDefault() {
	defer qt.Recovering("DropArea::showNormal")

	if ptr.Pointer() != nil {
		C.DropArea_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackDropArea_TabletEvent
func callbackDropArea_TabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback DropArea::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewDropAreaFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *DropArea) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect DropArea::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *DropArea) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect DropArea::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

func (ptr *DropArea) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("DropArea::tabletEvent")

	if ptr.Pointer() != nil {
		C.DropArea_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *DropArea) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("DropArea::tabletEvent")

	if ptr.Pointer() != nil {
		C.DropArea_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

//export callbackDropArea_Update
func callbackDropArea_Update(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback DropArea::update")

	if signal := qt.GetSignal(C.GoString(ptrName), "update"); signal != nil {
		signal.(func())()
	} else {
		NewDropAreaFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *DropArea) ConnectUpdate(f func()) {
	defer qt.Recovering("connect DropArea::update")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "update", f)
	}
}

func (ptr *DropArea) DisconnectUpdate() {
	defer qt.Recovering("disconnect DropArea::update")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "update")
	}
}

func (ptr *DropArea) Update() {
	defer qt.Recovering("DropArea::update")

	if ptr.Pointer() != nil {
		C.DropArea_Update(ptr.Pointer())
	}
}

func (ptr *DropArea) UpdateDefault() {
	defer qt.Recovering("DropArea::update")

	if ptr.Pointer() != nil {
		C.DropArea_UpdateDefault(ptr.Pointer())
	}
}

//export callbackDropArea_UpdateMicroFocus
func callbackDropArea_UpdateMicroFocus(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback DropArea::updateMicroFocus")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewDropAreaFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *DropArea) ConnectUpdateMicroFocus(f func()) {
	defer qt.Recovering("connect DropArea::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updateMicroFocus", f)
	}
}

func (ptr *DropArea) DisconnectUpdateMicroFocus() {
	defer qt.Recovering("disconnect DropArea::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updateMicroFocus")
	}
}

func (ptr *DropArea) UpdateMicroFocus() {
	defer qt.Recovering("DropArea::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.DropArea_UpdateMicroFocus(ptr.Pointer())
	}
}

func (ptr *DropArea) UpdateMicroFocusDefault() {
	defer qt.Recovering("DropArea::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.DropArea_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackDropArea_WheelEvent
func callbackDropArea_WheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback DropArea::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewDropAreaFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *DropArea) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect DropArea::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *DropArea) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect DropArea::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

func (ptr *DropArea) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("DropArea::wheelEvent")

	if ptr.Pointer() != nil {
		C.DropArea_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *DropArea) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("DropArea::wheelEvent")

	if ptr.Pointer() != nil {
		C.DropArea_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

//export callbackDropArea_TimerEvent
func callbackDropArea_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback DropArea::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewDropAreaFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *DropArea) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect DropArea::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *DropArea) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect DropArea::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *DropArea) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("DropArea::timerEvent")

	if ptr.Pointer() != nil {
		C.DropArea_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *DropArea) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("DropArea::timerEvent")

	if ptr.Pointer() != nil {
		C.DropArea_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackDropArea_ChildEvent
func callbackDropArea_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback DropArea::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewDropAreaFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *DropArea) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect DropArea::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *DropArea) DisconnectChildEvent() {
	defer qt.Recovering("disconnect DropArea::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *DropArea) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("DropArea::childEvent")

	if ptr.Pointer() != nil {
		C.DropArea_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *DropArea) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("DropArea::childEvent")

	if ptr.Pointer() != nil {
		C.DropArea_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackDropArea_ConnectNotify
func callbackDropArea_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback DropArea::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewDropAreaFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *DropArea) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect DropArea::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *DropArea) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect DropArea::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *DropArea) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("DropArea::connectNotify")

	if ptr.Pointer() != nil {
		C.DropArea_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *DropArea) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("DropArea::connectNotify")

	if ptr.Pointer() != nil {
		C.DropArea_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackDropArea_CustomEvent
func callbackDropArea_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback DropArea::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewDropAreaFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *DropArea) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect DropArea::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *DropArea) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect DropArea::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *DropArea) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("DropArea::customEvent")

	if ptr.Pointer() != nil {
		C.DropArea_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *DropArea) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("DropArea::customEvent")

	if ptr.Pointer() != nil {
		C.DropArea_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackDropArea_DeleteLater
func callbackDropArea_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback DropArea::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewDropAreaFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *DropArea) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect DropArea::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *DropArea) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect DropArea::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *DropArea) DeleteLater() {
	defer qt.Recovering("DropArea::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.DropArea_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *DropArea) DeleteLaterDefault() {
	defer qt.Recovering("DropArea::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.DropArea_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackDropArea_DisconnectNotify
func callbackDropArea_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback DropArea::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewDropAreaFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *DropArea) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect DropArea::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *DropArea) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect DropArea::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *DropArea) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("DropArea::disconnectNotify")

	if ptr.Pointer() != nil {
		C.DropArea_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *DropArea) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("DropArea::disconnectNotify")

	if ptr.Pointer() != nil {
		C.DropArea_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackDropArea_EventFilter
func callbackDropArea_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback DropArea::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewDropAreaFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *DropArea) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect DropArea::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *DropArea) DisconnectEventFilter() {
	defer qt.Recovering("disconnect DropArea::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *DropArea) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("DropArea::eventFilter")

	if ptr.Pointer() != nil {
		return C.DropArea_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *DropArea) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("DropArea::eventFilter")

	if ptr.Pointer() != nil {
		return C.DropArea_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackDropArea_MetaObject

type DropSiteWindow_ITF interface {
	widgets.QWidget_ITF
	DropSiteWindow_PTR() *DropSiteWindow
}

func (p *DropSiteWindow) DropSiteWindow_PTR() *DropSiteWindow {
	return p
}

func (p *DropSiteWindow) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QWidget_PTR().Pointer()
	}
	return nil
}

func (p *DropSiteWindow) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QWidget_PTR().SetPointer(ptr)
	}
}

func PointerFromDropSiteWindow(ptr DropSiteWindow_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.DropSiteWindow_PTR().Pointer()
	}
	return nil
}

func NewDropSiteWindowFromPointer(ptr unsafe.Pointer) *DropSiteWindow {
	var n = new(DropSiteWindow)
	n.SetPointer(ptr)
	return n
}

func newDropSiteWindowFromPointer(ptr unsafe.Pointer) *DropSiteWindow {
	var n = NewDropSiteWindowFromPointer(ptr)
	for len(n.ObjectName()) < len("DropSiteWindow_") {
		n.SetObjectName("DropSiteWindow_" + qt.Identifier())
	}
	return n
}

//export callbackDropSiteWindow_UpdateFormatsTable
func callbackDropSiteWindow_UpdateFormatsTable(ptr unsafe.Pointer, ptrName *C.char, mimeData unsafe.Pointer) {
	defer qt.Recovering("callback DropSiteWindow::updateFormatsTable")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateFormatsTable"); signal != nil {
		signal.(func(*core.QMimeData))(core.NewQMimeDataFromPointer(mimeData))
	}

}

func (ptr *DropSiteWindow) ConnectUpdateFormatsTable(f func(mimeData *core.QMimeData)) {
	defer qt.Recovering("connect DropSiteWindow::updateFormatsTable")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updateFormatsTable", f)
	}
}

func (ptr *DropSiteWindow) DisconnectUpdateFormatsTable(mimeData core.QMimeData_ITF) {
	defer qt.Recovering("disconnect DropSiteWindow::updateFormatsTable")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updateFormatsTable")
	}
}

func (ptr *DropSiteWindow) UpdateFormatsTable(mimeData core.QMimeData_ITF) {
	defer qt.Recovering("DropSiteWindow::updateFormatsTable")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_UpdateFormatsTable(ptr.Pointer(), core.PointerFromQMimeData(mimeData))
	}
}

func NewDropSiteWindow(parent widgets.QWidget_ITF, f core.Qt__WindowType) *DropSiteWindow {
	defer qt.Recovering("DropSiteWindow::DropSiteWindow")

	return newDropSiteWindowFromPointer(C.DropSiteWindow_NewDropSiteWindow(widgets.PointerFromQWidget(parent), C.int(f)))
}

func (ptr *DropSiteWindow) DestroyDropSiteWindow() {
	defer qt.Recovering("DropSiteWindow::~DropSiteWindow")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.DropSiteWindow_DestroyDropSiteWindow(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackDropSiteWindow_ActionEvent
func callbackDropSiteWindow_ActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback DropSiteWindow::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewDropSiteWindowFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *DropSiteWindow) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect DropSiteWindow::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *DropSiteWindow) DisconnectActionEvent() {
	defer qt.Recovering("disconnect DropSiteWindow::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

func (ptr *DropSiteWindow) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::actionEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *DropSiteWindow) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::actionEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackDropSiteWindow_DragEnterEvent
func callbackDropSiteWindow_DragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback DropSiteWindow::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewDropSiteWindowFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *DropSiteWindow) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect DropSiteWindow::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *DropSiteWindow) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect DropSiteWindow::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

func (ptr *DropSiteWindow) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *DropSiteWindow) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackDropSiteWindow_DragLeaveEvent
func callbackDropSiteWindow_DragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback DropSiteWindow::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewDropSiteWindowFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *DropSiteWindow) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect DropSiteWindow::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *DropSiteWindow) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect DropSiteWindow::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

func (ptr *DropSiteWindow) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *DropSiteWindow) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackDropSiteWindow_DragMoveEvent
func callbackDropSiteWindow_DragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback DropSiteWindow::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewDropSiteWindowFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *DropSiteWindow) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect DropSiteWindow::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *DropSiteWindow) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect DropSiteWindow::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

func (ptr *DropSiteWindow) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *DropSiteWindow) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackDropSiteWindow_DropEvent
func callbackDropSiteWindow_DropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback DropSiteWindow::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewDropSiteWindowFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *DropSiteWindow) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect DropSiteWindow::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *DropSiteWindow) DisconnectDropEvent() {
	defer qt.Recovering("disconnect DropSiteWindow::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

func (ptr *DropSiteWindow) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::dropEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *DropSiteWindow) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::dropEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

//export callbackDropSiteWindow_EnterEvent
func callbackDropSiteWindow_EnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback DropSiteWindow::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewDropSiteWindowFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *DropSiteWindow) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect DropSiteWindow::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *DropSiteWindow) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect DropSiteWindow::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

func (ptr *DropSiteWindow) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::enterEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *DropSiteWindow) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::enterEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackDropSiteWindow_FocusInEvent
func callbackDropSiteWindow_FocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback DropSiteWindow::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewDropSiteWindowFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *DropSiteWindow) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect DropSiteWindow::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *DropSiteWindow) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect DropSiteWindow::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

func (ptr *DropSiteWindow) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::focusInEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *DropSiteWindow) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::focusInEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackDropSiteWindow_FocusOutEvent
func callbackDropSiteWindow_FocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback DropSiteWindow::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewDropSiteWindowFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *DropSiteWindow) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect DropSiteWindow::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *DropSiteWindow) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect DropSiteWindow::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

func (ptr *DropSiteWindow) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::focusOutEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *DropSiteWindow) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::focusOutEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackDropSiteWindow_HideEvent
func callbackDropSiteWindow_HideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback DropSiteWindow::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewDropSiteWindowFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *DropSiteWindow) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect DropSiteWindow::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *DropSiteWindow) DisconnectHideEvent() {
	defer qt.Recovering("disconnect DropSiteWindow::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

func (ptr *DropSiteWindow) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::hideEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *DropSiteWindow) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::hideEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackDropSiteWindow_LeaveEvent
func callbackDropSiteWindow_LeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback DropSiteWindow::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewDropSiteWindowFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *DropSiteWindow) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect DropSiteWindow::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *DropSiteWindow) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect DropSiteWindow::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

func (ptr *DropSiteWindow) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::leaveEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *DropSiteWindow) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::leaveEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackDropSiteWindow_MinimumSizeHint
func callbackDropSiteWindow_MinimumSizeHint(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback DropSiteWindow::minimumSizeHint")

	if signal := qt.GetSignal(C.GoString(ptrName), "minimumSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewDropSiteWindowFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *DropSiteWindow) ConnectMinimumSizeHint(f func() *core.QSize) {
	defer qt.Recovering("connect DropSiteWindow::minimumSizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "minimumSizeHint", f)
	}
}

func (ptr *DropSiteWindow) DisconnectMinimumSizeHint() {
	defer qt.Recovering("disconnect DropSiteWindow::minimumSizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "minimumSizeHint")
	}
}

func (ptr *DropSiteWindow) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("DropSiteWindow::minimumSizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.DropSiteWindow_MinimumSizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *DropSiteWindow) MinimumSizeHintDefault() *core.QSize {
	defer qt.Recovering("DropSiteWindow::minimumSizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.DropSiteWindow_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackDropSiteWindow_MoveEvent
func callbackDropSiteWindow_MoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback DropSiteWindow::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewDropSiteWindowFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *DropSiteWindow) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect DropSiteWindow::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *DropSiteWindow) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect DropSiteWindow::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

func (ptr *DropSiteWindow) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::moveEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *DropSiteWindow) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::moveEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackDropSiteWindow_PaintEvent
func callbackDropSiteWindow_PaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback DropSiteWindow::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewDropSiteWindowFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *DropSiteWindow) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect DropSiteWindow::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *DropSiteWindow) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect DropSiteWindow::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

func (ptr *DropSiteWindow) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::paintEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *DropSiteWindow) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::paintEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

//export callbackDropSiteWindow_SetEnabled
func callbackDropSiteWindow_SetEnabled(ptr unsafe.Pointer, ptrName *C.char, vbo C.int) {
	defer qt.Recovering("callback DropSiteWindow::setEnabled")

	if signal := qt.GetSignal(C.GoString(ptrName), "setEnabled"); signal != nil {
		signal.(func(bool))(int(vbo) != 0)
	} else {
		NewDropSiteWindowFromPointer(ptr).SetEnabledDefault(int(vbo) != 0)
	}
}

func (ptr *DropSiteWindow) ConnectSetEnabled(f func(vbo bool)) {
	defer qt.Recovering("connect DropSiteWindow::setEnabled")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setEnabled", f)
	}
}

func (ptr *DropSiteWindow) DisconnectSetEnabled() {
	defer qt.Recovering("disconnect DropSiteWindow::setEnabled")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setEnabled")
	}
}

func (ptr *DropSiteWindow) SetEnabled(vbo bool) {
	defer qt.Recovering("DropSiteWindow::setEnabled")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_SetEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(vbo)))
	}
}

func (ptr *DropSiteWindow) SetEnabledDefault(vbo bool) {
	defer qt.Recovering("DropSiteWindow::setEnabled")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_SetEnabledDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(vbo)))
	}
}

//export callbackDropSiteWindow_SetStyleSheet
func callbackDropSiteWindow_SetStyleSheet(ptr unsafe.Pointer, ptrName *C.char, styleSheet *C.char) {
	defer qt.Recovering("callback DropSiteWindow::setStyleSheet")

	if signal := qt.GetSignal(C.GoString(ptrName), "setStyleSheet"); signal != nil {
		signal.(func(string))(C.GoString(styleSheet))
	} else {
		NewDropSiteWindowFromPointer(ptr).SetStyleSheetDefault(C.GoString(styleSheet))
	}
}

func (ptr *DropSiteWindow) ConnectSetStyleSheet(f func(styleSheet string)) {
	defer qt.Recovering("connect DropSiteWindow::setStyleSheet")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setStyleSheet", f)
	}
}

func (ptr *DropSiteWindow) DisconnectSetStyleSheet() {
	defer qt.Recovering("disconnect DropSiteWindow::setStyleSheet")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setStyleSheet")
	}
}

func (ptr *DropSiteWindow) SetStyleSheet(styleSheet string) {
	defer qt.Recovering("DropSiteWindow::setStyleSheet")

	if ptr.Pointer() != nil {
		var styleSheetC = C.CString(styleSheet)
		defer C.free(unsafe.Pointer(styleSheetC))
		C.DropSiteWindow_SetStyleSheet(ptr.Pointer(), styleSheetC)
	}
}

func (ptr *DropSiteWindow) SetStyleSheetDefault(styleSheet string) {
	defer qt.Recovering("DropSiteWindow::setStyleSheet")

	if ptr.Pointer() != nil {
		var styleSheetC = C.CString(styleSheet)
		defer C.free(unsafe.Pointer(styleSheetC))
		C.DropSiteWindow_SetStyleSheetDefault(ptr.Pointer(), styleSheetC)
	}
}

//export callbackDropSiteWindow_SetVisible
func callbackDropSiteWindow_SetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) {
	defer qt.Recovering("callback DropSiteWindow::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
	} else {
		NewDropSiteWindowFromPointer(ptr).SetVisibleDefault(int(visible) != 0)
	}
}

func (ptr *DropSiteWindow) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect DropSiteWindow::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *DropSiteWindow) DisconnectSetVisible() {
	defer qt.Recovering("disconnect DropSiteWindow::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

func (ptr *DropSiteWindow) SetVisible(visible bool) {
	defer qt.Recovering("DropSiteWindow::setVisible")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *DropSiteWindow) SetVisibleDefault(visible bool) {
	defer qt.Recovering("DropSiteWindow::setVisible")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

//export callbackDropSiteWindow_SetWindowModified
func callbackDropSiteWindow_SetWindowModified(ptr unsafe.Pointer, ptrName *C.char, vbo C.int) {
	defer qt.Recovering("callback DropSiteWindow::setWindowModified")

	if signal := qt.GetSignal(C.GoString(ptrName), "setWindowModified"); signal != nil {
		signal.(func(bool))(int(vbo) != 0)
	} else {
		NewDropSiteWindowFromPointer(ptr).SetWindowModifiedDefault(int(vbo) != 0)
	}
}

func (ptr *DropSiteWindow) ConnectSetWindowModified(f func(vbo bool)) {
	defer qt.Recovering("connect DropSiteWindow::setWindowModified")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setWindowModified", f)
	}
}

func (ptr *DropSiteWindow) DisconnectSetWindowModified() {
	defer qt.Recovering("disconnect DropSiteWindow::setWindowModified")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setWindowModified")
	}
}

func (ptr *DropSiteWindow) SetWindowModified(vbo bool) {
	defer qt.Recovering("DropSiteWindow::setWindowModified")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_SetWindowModified(ptr.Pointer(), C.int(qt.GoBoolToInt(vbo)))
	}
}

func (ptr *DropSiteWindow) SetWindowModifiedDefault(vbo bool) {
	defer qt.Recovering("DropSiteWindow::setWindowModified")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_SetWindowModifiedDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(vbo)))
	}
}

//export callbackDropSiteWindow_SetWindowTitle
func callbackDropSiteWindow_SetWindowTitle(ptr unsafe.Pointer, ptrName *C.char, vqs *C.char) {
	defer qt.Recovering("callback DropSiteWindow::setWindowTitle")

	if signal := qt.GetSignal(C.GoString(ptrName), "setWindowTitle"); signal != nil {
		signal.(func(string))(C.GoString(vqs))
	} else {
		NewDropSiteWindowFromPointer(ptr).SetWindowTitleDefault(C.GoString(vqs))
	}
}

func (ptr *DropSiteWindow) ConnectSetWindowTitle(f func(vqs string)) {
	defer qt.Recovering("connect DropSiteWindow::setWindowTitle")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setWindowTitle", f)
	}
}

func (ptr *DropSiteWindow) DisconnectSetWindowTitle() {
	defer qt.Recovering("disconnect DropSiteWindow::setWindowTitle")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setWindowTitle")
	}
}

func (ptr *DropSiteWindow) SetWindowTitle(vqs string) {
	defer qt.Recovering("DropSiteWindow::setWindowTitle")

	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.DropSiteWindow_SetWindowTitle(ptr.Pointer(), vqsC)
	}
}

func (ptr *DropSiteWindow) SetWindowTitleDefault(vqs string) {
	defer qt.Recovering("DropSiteWindow::setWindowTitle")

	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.DropSiteWindow_SetWindowTitleDefault(ptr.Pointer(), vqsC)
	}
}

//export callbackDropSiteWindow_ShowEvent
func callbackDropSiteWindow_ShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback DropSiteWindow::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewDropSiteWindowFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *DropSiteWindow) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect DropSiteWindow::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *DropSiteWindow) DisconnectShowEvent() {
	defer qt.Recovering("disconnect DropSiteWindow::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

func (ptr *DropSiteWindow) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::showEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *DropSiteWindow) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::showEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackDropSiteWindow_SizeHint
func callbackDropSiteWindow_SizeHint(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback DropSiteWindow::sizeHint")

	if signal := qt.GetSignal(C.GoString(ptrName), "sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewDropSiteWindowFromPointer(ptr).SizeHintDefault())
}

func (ptr *DropSiteWindow) ConnectSizeHint(f func() *core.QSize) {
	defer qt.Recovering("connect DropSiteWindow::sizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "sizeHint", f)
	}
}

func (ptr *DropSiteWindow) DisconnectSizeHint() {
	defer qt.Recovering("disconnect DropSiteWindow::sizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "sizeHint")
	}
}

func (ptr *DropSiteWindow) SizeHint() *core.QSize {
	defer qt.Recovering("DropSiteWindow::sizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.DropSiteWindow_SizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *DropSiteWindow) SizeHintDefault() *core.QSize {
	defer qt.Recovering("DropSiteWindow::sizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.DropSiteWindow_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackDropSiteWindow_ChangeEvent
func callbackDropSiteWindow_ChangeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback DropSiteWindow::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewDropSiteWindowFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *DropSiteWindow) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect DropSiteWindow::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *DropSiteWindow) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect DropSiteWindow::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

func (ptr *DropSiteWindow) ChangeEvent(event core.QEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::changeEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *DropSiteWindow) ChangeEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::changeEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackDropSiteWindow_Close
func callbackDropSiteWindow_Close(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback DropSiteWindow::close")

	if signal := qt.GetSignal(C.GoString(ptrName), "close"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewDropSiteWindowFromPointer(ptr).CloseDefault()))
}

func (ptr *DropSiteWindow) ConnectClose(f func() bool) {
	defer qt.Recovering("connect DropSiteWindow::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "close", f)
	}
}

func (ptr *DropSiteWindow) DisconnectClose() {
	defer qt.Recovering("disconnect DropSiteWindow::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "close")
	}
}

func (ptr *DropSiteWindow) Close() bool {
	defer qt.Recovering("DropSiteWindow::close")

	if ptr.Pointer() != nil {
		return C.DropSiteWindow_Close(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *DropSiteWindow) CloseDefault() bool {
	defer qt.Recovering("DropSiteWindow::close")

	if ptr.Pointer() != nil {
		return C.DropSiteWindow_CloseDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackDropSiteWindow_CloseEvent
func callbackDropSiteWindow_CloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback DropSiteWindow::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewDropSiteWindowFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *DropSiteWindow) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect DropSiteWindow::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *DropSiteWindow) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect DropSiteWindow::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

func (ptr *DropSiteWindow) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::closeEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *DropSiteWindow) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::closeEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

//export callbackDropSiteWindow_ContextMenuEvent
func callbackDropSiteWindow_ContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback DropSiteWindow::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewDropSiteWindowFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *DropSiteWindow) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect DropSiteWindow::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *DropSiteWindow) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect DropSiteWindow::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

func (ptr *DropSiteWindow) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *DropSiteWindow) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

//export callbackDropSiteWindow_FocusNextPrevChild
func callbackDropSiteWindow_FocusNextPrevChild(ptr unsafe.Pointer, ptrName *C.char, next C.int) C.int {
	defer qt.Recovering("callback DropSiteWindow::focusNextPrevChild")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusNextPrevChild"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(bool) bool)(int(next) != 0)))
	}

	return C.int(qt.GoBoolToInt(NewDropSiteWindowFromPointer(ptr).FocusNextPrevChildDefault(int(next) != 0)))
}

func (ptr *DropSiteWindow) ConnectFocusNextPrevChild(f func(next bool) bool) {
	defer qt.Recovering("connect DropSiteWindow::focusNextPrevChild")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusNextPrevChild", f)
	}
}

func (ptr *DropSiteWindow) DisconnectFocusNextPrevChild() {
	defer qt.Recovering("disconnect DropSiteWindow::focusNextPrevChild")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusNextPrevChild")
	}
}

func (ptr *DropSiteWindow) FocusNextPrevChild(next bool) bool {
	defer qt.Recovering("DropSiteWindow::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.DropSiteWindow_FocusNextPrevChild(ptr.Pointer(), C.int(qt.GoBoolToInt(next))) != 0
	}
	return false
}

func (ptr *DropSiteWindow) FocusNextPrevChildDefault(next bool) bool {
	defer qt.Recovering("DropSiteWindow::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.DropSiteWindow_FocusNextPrevChildDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(next))) != 0
	}
	return false
}

//export callbackDropSiteWindow_HasHeightForWidth
func callbackDropSiteWindow_HasHeightForWidth(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback DropSiteWindow::hasHeightForWidth")

	if signal := qt.GetSignal(C.GoString(ptrName), "hasHeightForWidth"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewDropSiteWindowFromPointer(ptr).HasHeightForWidthDefault()))
}

func (ptr *DropSiteWindow) ConnectHasHeightForWidth(f func() bool) {
	defer qt.Recovering("connect DropSiteWindow::hasHeightForWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hasHeightForWidth", f)
	}
}

func (ptr *DropSiteWindow) DisconnectHasHeightForWidth() {
	defer qt.Recovering("disconnect DropSiteWindow::hasHeightForWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hasHeightForWidth")
	}
}

func (ptr *DropSiteWindow) HasHeightForWidth() bool {
	defer qt.Recovering("DropSiteWindow::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.DropSiteWindow_HasHeightForWidth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *DropSiteWindow) HasHeightForWidthDefault() bool {
	defer qt.Recovering("DropSiteWindow::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.DropSiteWindow_HasHeightForWidthDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackDropSiteWindow_HeightForWidth
func callbackDropSiteWindow_HeightForWidth(ptr unsafe.Pointer, ptrName *C.char, w C.int) C.int {
	defer qt.Recovering("callback DropSiteWindow::heightForWidth")

	if signal := qt.GetSignal(C.GoString(ptrName), "heightForWidth"); signal != nil {
		return C.int(signal.(func(int) int)(int(w)))
	}

	return C.int(NewDropSiteWindowFromPointer(ptr).HeightForWidthDefault(int(w)))
}

func (ptr *DropSiteWindow) ConnectHeightForWidth(f func(w int) int) {
	defer qt.Recovering("connect DropSiteWindow::heightForWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "heightForWidth", f)
	}
}

func (ptr *DropSiteWindow) DisconnectHeightForWidth() {
	defer qt.Recovering("disconnect DropSiteWindow::heightForWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "heightForWidth")
	}
}

func (ptr *DropSiteWindow) HeightForWidth(w int) int {
	defer qt.Recovering("DropSiteWindow::heightForWidth")

	if ptr.Pointer() != nil {
		return int(C.DropSiteWindow_HeightForWidth(ptr.Pointer(), C.int(w)))
	}
	return 0
}

func (ptr *DropSiteWindow) HeightForWidthDefault(w int) int {
	defer qt.Recovering("DropSiteWindow::heightForWidth")

	if ptr.Pointer() != nil {
		return int(C.DropSiteWindow_HeightForWidthDefault(ptr.Pointer(), C.int(w)))
	}
	return 0
}

//export callbackDropSiteWindow_Hide
func callbackDropSiteWindow_Hide(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback DropSiteWindow::hide")

	if signal := qt.GetSignal(C.GoString(ptrName), "hide"); signal != nil {
		signal.(func())()
	} else {
		NewDropSiteWindowFromPointer(ptr).HideDefault()
	}
}

func (ptr *DropSiteWindow) ConnectHide(f func()) {
	defer qt.Recovering("connect DropSiteWindow::hide")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hide", f)
	}
}

func (ptr *DropSiteWindow) DisconnectHide() {
	defer qt.Recovering("disconnect DropSiteWindow::hide")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hide")
	}
}

func (ptr *DropSiteWindow) Hide() {
	defer qt.Recovering("DropSiteWindow::hide")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_Hide(ptr.Pointer())
	}
}

func (ptr *DropSiteWindow) HideDefault() {
	defer qt.Recovering("DropSiteWindow::hide")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_HideDefault(ptr.Pointer())
	}
}

//export callbackDropSiteWindow_InputMethodEvent
func callbackDropSiteWindow_InputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback DropSiteWindow::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewDropSiteWindowFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *DropSiteWindow) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect DropSiteWindow::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *DropSiteWindow) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect DropSiteWindow::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

func (ptr *DropSiteWindow) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *DropSiteWindow) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackDropSiteWindow_InputMethodQuery
func callbackDropSiteWindow_InputMethodQuery(ptr unsafe.Pointer, ptrName *C.char, query C.int) unsafe.Pointer {
	defer qt.Recovering("callback DropSiteWindow::inputMethodQuery")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewDropSiteWindowFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *DropSiteWindow) ConnectInputMethodQuery(f func(query core.Qt__InputMethodQuery) *core.QVariant) {
	defer qt.Recovering("connect DropSiteWindow::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodQuery", f)
	}
}

func (ptr *DropSiteWindow) DisconnectInputMethodQuery() {
	defer qt.Recovering("disconnect DropSiteWindow::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodQuery")
	}
}

func (ptr *DropSiteWindow) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("DropSiteWindow::inputMethodQuery")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.DropSiteWindow_InputMethodQuery(ptr.Pointer(), C.int(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *DropSiteWindow) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("DropSiteWindow::inputMethodQuery")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.DropSiteWindow_InputMethodQueryDefault(ptr.Pointer(), C.int(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackDropSiteWindow_KeyPressEvent
func callbackDropSiteWindow_KeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback DropSiteWindow::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewDropSiteWindowFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *DropSiteWindow) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect DropSiteWindow::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *DropSiteWindow) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect DropSiteWindow::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

func (ptr *DropSiteWindow) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::keyPressEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *DropSiteWindow) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::keyPressEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackDropSiteWindow_KeyReleaseEvent
func callbackDropSiteWindow_KeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback DropSiteWindow::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewDropSiteWindowFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *DropSiteWindow) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect DropSiteWindow::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *DropSiteWindow) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect DropSiteWindow::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

func (ptr *DropSiteWindow) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *DropSiteWindow) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackDropSiteWindow_Lower
func callbackDropSiteWindow_Lower(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback DropSiteWindow::lower")

	if signal := qt.GetSignal(C.GoString(ptrName), "lower"); signal != nil {
		signal.(func())()
	} else {
		NewDropSiteWindowFromPointer(ptr).LowerDefault()
	}
}

func (ptr *DropSiteWindow) ConnectLower(f func()) {
	defer qt.Recovering("connect DropSiteWindow::lower")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "lower", f)
	}
}

func (ptr *DropSiteWindow) DisconnectLower() {
	defer qt.Recovering("disconnect DropSiteWindow::lower")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "lower")
	}
}

func (ptr *DropSiteWindow) Lower() {
	defer qt.Recovering("DropSiteWindow::lower")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_Lower(ptr.Pointer())
	}
}

func (ptr *DropSiteWindow) LowerDefault() {
	defer qt.Recovering("DropSiteWindow::lower")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_LowerDefault(ptr.Pointer())
	}
}

//export callbackDropSiteWindow_MouseDoubleClickEvent
func callbackDropSiteWindow_MouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback DropSiteWindow::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewDropSiteWindowFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *DropSiteWindow) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect DropSiteWindow::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *DropSiteWindow) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect DropSiteWindow::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

func (ptr *DropSiteWindow) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *DropSiteWindow) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackDropSiteWindow_MouseMoveEvent
func callbackDropSiteWindow_MouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback DropSiteWindow::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewDropSiteWindowFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *DropSiteWindow) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect DropSiteWindow::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *DropSiteWindow) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect DropSiteWindow::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

func (ptr *DropSiteWindow) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *DropSiteWindow) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackDropSiteWindow_MousePressEvent
func callbackDropSiteWindow_MousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback DropSiteWindow::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewDropSiteWindowFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *DropSiteWindow) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect DropSiteWindow::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *DropSiteWindow) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect DropSiteWindow::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

func (ptr *DropSiteWindow) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::mousePressEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *DropSiteWindow) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::mousePressEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackDropSiteWindow_MouseReleaseEvent
func callbackDropSiteWindow_MouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback DropSiteWindow::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewDropSiteWindowFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *DropSiteWindow) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect DropSiteWindow::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *DropSiteWindow) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect DropSiteWindow::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

func (ptr *DropSiteWindow) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *DropSiteWindow) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackDropSiteWindow_NativeEvent
func callbackDropSiteWindow_NativeEvent(ptr unsafe.Pointer, ptrName *C.char, eventType *C.char, message unsafe.Pointer, result C.long) C.int {
	defer qt.Recovering("callback DropSiteWindow::nativeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "nativeEvent"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string, unsafe.Pointer, int) bool)(qt.HexDecodeToString(C.GoString(eventType)), message, int(result))))
	}

	return C.int(qt.GoBoolToInt(NewDropSiteWindowFromPointer(ptr).NativeEventDefault(qt.HexDecodeToString(C.GoString(eventType)), message, int(result))))
}

func (ptr *DropSiteWindow) ConnectNativeEvent(f func(eventType string, message unsafe.Pointer, result int) bool) {
	defer qt.Recovering("connect DropSiteWindow::nativeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "nativeEvent", f)
	}
}

func (ptr *DropSiteWindow) DisconnectNativeEvent() {
	defer qt.Recovering("disconnect DropSiteWindow::nativeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "nativeEvent")
	}
}

func (ptr *DropSiteWindow) NativeEvent(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("DropSiteWindow::nativeEvent")

	if ptr.Pointer() != nil {
		var eventTypeC = C.CString(hex.EncodeToString([]byte(eventType)))
		defer C.free(unsafe.Pointer(eventTypeC))
		return C.DropSiteWindow_NativeEvent(ptr.Pointer(), eventTypeC, message, C.long(result)) != 0
	}
	return false
}

func (ptr *DropSiteWindow) NativeEventDefault(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("DropSiteWindow::nativeEvent")

	if ptr.Pointer() != nil {
		var eventTypeC = C.CString(hex.EncodeToString([]byte(eventType)))
		defer C.free(unsafe.Pointer(eventTypeC))
		return C.DropSiteWindow_NativeEventDefault(ptr.Pointer(), eventTypeC, message, C.long(result)) != 0
	}
	return false
}

//export callbackDropSiteWindow_Raise
func callbackDropSiteWindow_Raise(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback DropSiteWindow::raise")

	if signal := qt.GetSignal(C.GoString(ptrName), "raise"); signal != nil {
		signal.(func())()
	} else {
		NewDropSiteWindowFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *DropSiteWindow) ConnectRaise(f func()) {
	defer qt.Recovering("connect DropSiteWindow::raise")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "raise", f)
	}
}

func (ptr *DropSiteWindow) DisconnectRaise() {
	defer qt.Recovering("disconnect DropSiteWindow::raise")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "raise")
	}
}

func (ptr *DropSiteWindow) Raise() {
	defer qt.Recovering("DropSiteWindow::raise")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_Raise(ptr.Pointer())
	}
}

func (ptr *DropSiteWindow) RaiseDefault() {
	defer qt.Recovering("DropSiteWindow::raise")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_RaiseDefault(ptr.Pointer())
	}
}

//export callbackDropSiteWindow_Repaint
func callbackDropSiteWindow_Repaint(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback DropSiteWindow::repaint")

	if signal := qt.GetSignal(C.GoString(ptrName), "repaint"); signal != nil {
		signal.(func())()
	} else {
		NewDropSiteWindowFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *DropSiteWindow) ConnectRepaint(f func()) {
	defer qt.Recovering("connect DropSiteWindow::repaint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "repaint", f)
	}
}

func (ptr *DropSiteWindow) DisconnectRepaint() {
	defer qt.Recovering("disconnect DropSiteWindow::repaint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "repaint")
	}
}

func (ptr *DropSiteWindow) Repaint() {
	defer qt.Recovering("DropSiteWindow::repaint")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_Repaint(ptr.Pointer())
	}
}

func (ptr *DropSiteWindow) RepaintDefault() {
	defer qt.Recovering("DropSiteWindow::repaint")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_RepaintDefault(ptr.Pointer())
	}
}

//export callbackDropSiteWindow_ResizeEvent
func callbackDropSiteWindow_ResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback DropSiteWindow::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewDropSiteWindowFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *DropSiteWindow) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect DropSiteWindow::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *DropSiteWindow) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect DropSiteWindow::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

func (ptr *DropSiteWindow) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::resizeEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *DropSiteWindow) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::resizeEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

//export callbackDropSiteWindow_SetDisabled
func callbackDropSiteWindow_SetDisabled(ptr unsafe.Pointer, ptrName *C.char, disable C.int) {
	defer qt.Recovering("callback DropSiteWindow::setDisabled")

	if signal := qt.GetSignal(C.GoString(ptrName), "setDisabled"); signal != nil {
		signal.(func(bool))(int(disable) != 0)
	} else {
		NewDropSiteWindowFromPointer(ptr).SetDisabledDefault(int(disable) != 0)
	}
}

func (ptr *DropSiteWindow) ConnectSetDisabled(f func(disable bool)) {
	defer qt.Recovering("connect DropSiteWindow::setDisabled")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setDisabled", f)
	}
}

func (ptr *DropSiteWindow) DisconnectSetDisabled() {
	defer qt.Recovering("disconnect DropSiteWindow::setDisabled")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setDisabled")
	}
}

func (ptr *DropSiteWindow) SetDisabled(disable bool) {
	defer qt.Recovering("DropSiteWindow::setDisabled")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_SetDisabled(ptr.Pointer(), C.int(qt.GoBoolToInt(disable)))
	}
}

func (ptr *DropSiteWindow) SetDisabledDefault(disable bool) {
	defer qt.Recovering("DropSiteWindow::setDisabled")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_SetDisabledDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(disable)))
	}
}

//export callbackDropSiteWindow_SetFocus2
func callbackDropSiteWindow_SetFocus2(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback DropSiteWindow::setFocus")

	if signal := qt.GetSignal(C.GoString(ptrName), "setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewDropSiteWindowFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *DropSiteWindow) ConnectSetFocus2(f func()) {
	defer qt.Recovering("connect DropSiteWindow::setFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setFocus2", f)
	}
}

func (ptr *DropSiteWindow) DisconnectSetFocus2() {
	defer qt.Recovering("disconnect DropSiteWindow::setFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setFocus2")
	}
}

func (ptr *DropSiteWindow) SetFocus2() {
	defer qt.Recovering("DropSiteWindow::setFocus")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_SetFocus2(ptr.Pointer())
	}
}

func (ptr *DropSiteWindow) SetFocus2Default() {
	defer qt.Recovering("DropSiteWindow::setFocus")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackDropSiteWindow_SetHidden
func callbackDropSiteWindow_SetHidden(ptr unsafe.Pointer, ptrName *C.char, hidden C.int) {
	defer qt.Recovering("callback DropSiteWindow::setHidden")

	if signal := qt.GetSignal(C.GoString(ptrName), "setHidden"); signal != nil {
		signal.(func(bool))(int(hidden) != 0)
	} else {
		NewDropSiteWindowFromPointer(ptr).SetHiddenDefault(int(hidden) != 0)
	}
}

func (ptr *DropSiteWindow) ConnectSetHidden(f func(hidden bool)) {
	defer qt.Recovering("connect DropSiteWindow::setHidden")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setHidden", f)
	}
}

func (ptr *DropSiteWindow) DisconnectSetHidden() {
	defer qt.Recovering("disconnect DropSiteWindow::setHidden")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setHidden")
	}
}

func (ptr *DropSiteWindow) SetHidden(hidden bool) {
	defer qt.Recovering("DropSiteWindow::setHidden")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_SetHidden(ptr.Pointer(), C.int(qt.GoBoolToInt(hidden)))
	}
}

func (ptr *DropSiteWindow) SetHiddenDefault(hidden bool) {
	defer qt.Recovering("DropSiteWindow::setHidden")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_SetHiddenDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(hidden)))
	}
}

//export callbackDropSiteWindow_Show
func callbackDropSiteWindow_Show(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback DropSiteWindow::show")

	if signal := qt.GetSignal(C.GoString(ptrName), "show"); signal != nil {
		signal.(func())()
	} else {
		NewDropSiteWindowFromPointer(ptr).ShowDefault()
	}
}

func (ptr *DropSiteWindow) ConnectShow(f func()) {
	defer qt.Recovering("connect DropSiteWindow::show")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "show", f)
	}
}

func (ptr *DropSiteWindow) DisconnectShow() {
	defer qt.Recovering("disconnect DropSiteWindow::show")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "show")
	}
}

func (ptr *DropSiteWindow) Show() {
	defer qt.Recovering("DropSiteWindow::show")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_Show(ptr.Pointer())
	}
}

func (ptr *DropSiteWindow) ShowDefault() {
	defer qt.Recovering("DropSiteWindow::show")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_ShowDefault(ptr.Pointer())
	}
}

//export callbackDropSiteWindow_ShowFullScreen
func callbackDropSiteWindow_ShowFullScreen(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback DropSiteWindow::showFullScreen")

	if signal := qt.GetSignal(C.GoString(ptrName), "showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewDropSiteWindowFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *DropSiteWindow) ConnectShowFullScreen(f func()) {
	defer qt.Recovering("connect DropSiteWindow::showFullScreen")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showFullScreen", f)
	}
}

func (ptr *DropSiteWindow) DisconnectShowFullScreen() {
	defer qt.Recovering("disconnect DropSiteWindow::showFullScreen")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showFullScreen")
	}
}

func (ptr *DropSiteWindow) ShowFullScreen() {
	defer qt.Recovering("DropSiteWindow::showFullScreen")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_ShowFullScreen(ptr.Pointer())
	}
}

func (ptr *DropSiteWindow) ShowFullScreenDefault() {
	defer qt.Recovering("DropSiteWindow::showFullScreen")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackDropSiteWindow_ShowMaximized
func callbackDropSiteWindow_ShowMaximized(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback DropSiteWindow::showMaximized")

	if signal := qt.GetSignal(C.GoString(ptrName), "showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewDropSiteWindowFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *DropSiteWindow) ConnectShowMaximized(f func()) {
	defer qt.Recovering("connect DropSiteWindow::showMaximized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showMaximized", f)
	}
}

func (ptr *DropSiteWindow) DisconnectShowMaximized() {
	defer qt.Recovering("disconnect DropSiteWindow::showMaximized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showMaximized")
	}
}

func (ptr *DropSiteWindow) ShowMaximized() {
	defer qt.Recovering("DropSiteWindow::showMaximized")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_ShowMaximized(ptr.Pointer())
	}
}

func (ptr *DropSiteWindow) ShowMaximizedDefault() {
	defer qt.Recovering("DropSiteWindow::showMaximized")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackDropSiteWindow_ShowMinimized
func callbackDropSiteWindow_ShowMinimized(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback DropSiteWindow::showMinimized")

	if signal := qt.GetSignal(C.GoString(ptrName), "showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewDropSiteWindowFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *DropSiteWindow) ConnectShowMinimized(f func()) {
	defer qt.Recovering("connect DropSiteWindow::showMinimized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showMinimized", f)
	}
}

func (ptr *DropSiteWindow) DisconnectShowMinimized() {
	defer qt.Recovering("disconnect DropSiteWindow::showMinimized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showMinimized")
	}
}

func (ptr *DropSiteWindow) ShowMinimized() {
	defer qt.Recovering("DropSiteWindow::showMinimized")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_ShowMinimized(ptr.Pointer())
	}
}

func (ptr *DropSiteWindow) ShowMinimizedDefault() {
	defer qt.Recovering("DropSiteWindow::showMinimized")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackDropSiteWindow_ShowNormal
func callbackDropSiteWindow_ShowNormal(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback DropSiteWindow::showNormal")

	if signal := qt.GetSignal(C.GoString(ptrName), "showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewDropSiteWindowFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *DropSiteWindow) ConnectShowNormal(f func()) {
	defer qt.Recovering("connect DropSiteWindow::showNormal")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showNormal", f)
	}
}

func (ptr *DropSiteWindow) DisconnectShowNormal() {
	defer qt.Recovering("disconnect DropSiteWindow::showNormal")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showNormal")
	}
}

func (ptr *DropSiteWindow) ShowNormal() {
	defer qt.Recovering("DropSiteWindow::showNormal")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_ShowNormal(ptr.Pointer())
	}
}

func (ptr *DropSiteWindow) ShowNormalDefault() {
	defer qt.Recovering("DropSiteWindow::showNormal")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackDropSiteWindow_TabletEvent
func callbackDropSiteWindow_TabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback DropSiteWindow::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewDropSiteWindowFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *DropSiteWindow) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect DropSiteWindow::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *DropSiteWindow) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect DropSiteWindow::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

func (ptr *DropSiteWindow) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::tabletEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *DropSiteWindow) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::tabletEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

//export callbackDropSiteWindow_Update
func callbackDropSiteWindow_Update(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback DropSiteWindow::update")

	if signal := qt.GetSignal(C.GoString(ptrName), "update"); signal != nil {
		signal.(func())()
	} else {
		NewDropSiteWindowFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *DropSiteWindow) ConnectUpdate(f func()) {
	defer qt.Recovering("connect DropSiteWindow::update")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "update", f)
	}
}

func (ptr *DropSiteWindow) DisconnectUpdate() {
	defer qt.Recovering("disconnect DropSiteWindow::update")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "update")
	}
}

func (ptr *DropSiteWindow) Update() {
	defer qt.Recovering("DropSiteWindow::update")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_Update(ptr.Pointer())
	}
}

func (ptr *DropSiteWindow) UpdateDefault() {
	defer qt.Recovering("DropSiteWindow::update")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_UpdateDefault(ptr.Pointer())
	}
}

//export callbackDropSiteWindow_UpdateMicroFocus
func callbackDropSiteWindow_UpdateMicroFocus(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback DropSiteWindow::updateMicroFocus")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewDropSiteWindowFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *DropSiteWindow) ConnectUpdateMicroFocus(f func()) {
	defer qt.Recovering("connect DropSiteWindow::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updateMicroFocus", f)
	}
}

func (ptr *DropSiteWindow) DisconnectUpdateMicroFocus() {
	defer qt.Recovering("disconnect DropSiteWindow::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updateMicroFocus")
	}
}

func (ptr *DropSiteWindow) UpdateMicroFocus() {
	defer qt.Recovering("DropSiteWindow::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_UpdateMicroFocus(ptr.Pointer())
	}
}

func (ptr *DropSiteWindow) UpdateMicroFocusDefault() {
	defer qt.Recovering("DropSiteWindow::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackDropSiteWindow_WheelEvent
func callbackDropSiteWindow_WheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback DropSiteWindow::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewDropSiteWindowFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *DropSiteWindow) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect DropSiteWindow::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *DropSiteWindow) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect DropSiteWindow::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

func (ptr *DropSiteWindow) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::wheelEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *DropSiteWindow) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::wheelEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

//export callbackDropSiteWindow_TimerEvent
func callbackDropSiteWindow_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback DropSiteWindow::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewDropSiteWindowFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *DropSiteWindow) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect DropSiteWindow::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *DropSiteWindow) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect DropSiteWindow::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *DropSiteWindow) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::timerEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *DropSiteWindow) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::timerEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackDropSiteWindow_ChildEvent
func callbackDropSiteWindow_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback DropSiteWindow::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewDropSiteWindowFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *DropSiteWindow) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect DropSiteWindow::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *DropSiteWindow) DisconnectChildEvent() {
	defer qt.Recovering("disconnect DropSiteWindow::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *DropSiteWindow) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::childEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *DropSiteWindow) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::childEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackDropSiteWindow_ConnectNotify
func callbackDropSiteWindow_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback DropSiteWindow::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewDropSiteWindowFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *DropSiteWindow) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect DropSiteWindow::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *DropSiteWindow) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect DropSiteWindow::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *DropSiteWindow) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("DropSiteWindow::connectNotify")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *DropSiteWindow) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("DropSiteWindow::connectNotify")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackDropSiteWindow_CustomEvent
func callbackDropSiteWindow_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback DropSiteWindow::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewDropSiteWindowFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *DropSiteWindow) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect DropSiteWindow::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *DropSiteWindow) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect DropSiteWindow::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *DropSiteWindow) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::customEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *DropSiteWindow) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("DropSiteWindow::customEvent")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackDropSiteWindow_DeleteLater
func callbackDropSiteWindow_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback DropSiteWindow::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewDropSiteWindowFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *DropSiteWindow) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect DropSiteWindow::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *DropSiteWindow) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect DropSiteWindow::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *DropSiteWindow) DeleteLater() {
	defer qt.Recovering("DropSiteWindow::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.DropSiteWindow_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *DropSiteWindow) DeleteLaterDefault() {
	defer qt.Recovering("DropSiteWindow::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.DropSiteWindow_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackDropSiteWindow_DisconnectNotify
func callbackDropSiteWindow_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback DropSiteWindow::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewDropSiteWindowFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *DropSiteWindow) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect DropSiteWindow::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *DropSiteWindow) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect DropSiteWindow::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *DropSiteWindow) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("DropSiteWindow::disconnectNotify")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *DropSiteWindow) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("DropSiteWindow::disconnectNotify")

	if ptr.Pointer() != nil {
		C.DropSiteWindow_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackDropSiteWindow_EventFilter
func callbackDropSiteWindow_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback DropSiteWindow::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewDropSiteWindowFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *DropSiteWindow) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect DropSiteWindow::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *DropSiteWindow) DisconnectEventFilter() {
	defer qt.Recovering("disconnect DropSiteWindow::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *DropSiteWindow) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("DropSiteWindow::eventFilter")

	if ptr.Pointer() != nil {
		return C.DropSiteWindow_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *DropSiteWindow) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("DropSiteWindow::eventFilter")

	if ptr.Pointer() != nil {
		return C.DropSiteWindow_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackDropSiteWindow_MetaObject
