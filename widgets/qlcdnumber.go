package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QLCDNumber struct {
	QFrame
}

type QLCDNumber_ITF interface {
	QFrame_ITF
	QLCDNumber_PTR() *QLCDNumber
}

func PointerFromQLCDNumber(ptr QLCDNumber_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLCDNumber_PTR().Pointer()
	}
	return nil
}

func NewQLCDNumberFromPointer(ptr unsafe.Pointer) *QLCDNumber {
	var n = new(QLCDNumber)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QLCDNumber_") {
		n.SetObjectName("QLCDNumber_" + qt.Identifier())
	}
	return n
}

func (ptr *QLCDNumber) QLCDNumber_PTR() *QLCDNumber {
	return ptr
}

//QLCDNumber::Mode
type QLCDNumber__Mode int64

const (
	QLCDNumber__Hex = QLCDNumber__Mode(0)
	QLCDNumber__Dec = QLCDNumber__Mode(1)
	QLCDNumber__Oct = QLCDNumber__Mode(2)
	QLCDNumber__Bin = QLCDNumber__Mode(3)
)

//QLCDNumber::SegmentStyle
type QLCDNumber__SegmentStyle int64

var (
	QLCDNumber__Outline = QLCDNumber__SegmentStyle(0)
	QLCDNumber__Filled  = QLCDNumber__SegmentStyle(1)
	QLCDNumber__Flat    = QLCDNumber__SegmentStyle(2)
)

func (ptr *QLCDNumber) IntValue() int {
	defer qt.Recovering("QLCDNumber::intValue")

	if ptr.Pointer() != nil {
		return int(C.QLCDNumber_IntValue(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLCDNumber) Mode() QLCDNumber__Mode {
	defer qt.Recovering("QLCDNumber::mode")

	if ptr.Pointer() != nil {
		return QLCDNumber__Mode(C.QLCDNumber_Mode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLCDNumber) SegmentStyle() QLCDNumber__SegmentStyle {
	defer qt.Recovering("QLCDNumber::segmentStyle")

	if ptr.Pointer() != nil {
		return QLCDNumber__SegmentStyle(C.QLCDNumber_SegmentStyle(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLCDNumber) SetMode(v QLCDNumber__Mode) {
	defer qt.Recovering("QLCDNumber::setMode")

	if ptr.Pointer() != nil {
		C.QLCDNumber_SetMode(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QLCDNumber) SetSegmentStyle(v QLCDNumber__SegmentStyle) {
	defer qt.Recovering("QLCDNumber::setSegmentStyle")

	if ptr.Pointer() != nil {
		C.QLCDNumber_SetSegmentStyle(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QLCDNumber) SetSmallDecimalPoint(v bool) {
	defer qt.Recovering("QLCDNumber::setSmallDecimalPoint")

	if ptr.Pointer() != nil {
		C.QLCDNumber_SetSmallDecimalPoint(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QLCDNumber) SmallDecimalPoint() bool {
	defer qt.Recovering("QLCDNumber::smallDecimalPoint")

	if ptr.Pointer() != nil {
		return C.QLCDNumber_SmallDecimalPoint(ptr.Pointer()) != 0
	}
	return false
}

func NewQLCDNumber(parent QWidget_ITF) *QLCDNumber {
	defer qt.Recovering("QLCDNumber::QLCDNumber")

	return NewQLCDNumberFromPointer(C.QLCDNumber_NewQLCDNumber(PointerFromQWidget(parent)))
}

func (ptr *QLCDNumber) CheckOverflow2(num int) bool {
	defer qt.Recovering("QLCDNumber::checkOverflow")

	if ptr.Pointer() != nil {
		return C.QLCDNumber_CheckOverflow2(ptr.Pointer(), C.int(num)) != 0
	}
	return false
}

func (ptr *QLCDNumber) DigitCount() int {
	defer qt.Recovering("QLCDNumber::digitCount")

	if ptr.Pointer() != nil {
		return int(C.QLCDNumber_DigitCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLCDNumber) Display(s string) {
	defer qt.Recovering("QLCDNumber::display")

	if ptr.Pointer() != nil {
		C.QLCDNumber_Display(ptr.Pointer(), C.CString(s))
	}
}

func (ptr *QLCDNumber) Display3(num int) {
	defer qt.Recovering("QLCDNumber::display")

	if ptr.Pointer() != nil {
		C.QLCDNumber_Display3(ptr.Pointer(), C.int(num))
	}
}

func (ptr *QLCDNumber) ConnectOverflow(f func()) {
	defer qt.Recovering("connect QLCDNumber::overflow")

	if ptr.Pointer() != nil {
		C.QLCDNumber_ConnectOverflow(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "overflow", f)
	}
}

func (ptr *QLCDNumber) DisconnectOverflow() {
	defer qt.Recovering("disconnect QLCDNumber::overflow")

	if ptr.Pointer() != nil {
		C.QLCDNumber_DisconnectOverflow(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "overflow")
	}
}

//export callbackQLCDNumberOverflow
func callbackQLCDNumberOverflow(ptrName *C.char) {
	defer qt.Recovering("callback QLCDNumber::overflow")

	if signal := qt.GetSignal(C.GoString(ptrName), "overflow"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QLCDNumber) ConnectPaintEvent(f func(v *gui.QPaintEvent)) {
	defer qt.Recovering("connect QLCDNumber::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QLCDNumber) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QLCDNumber::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQLCDNumberPaintEvent
func callbackQLCDNumberPaintEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QLCDNumber::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(v))
		return true
	}
	return false

}

func (ptr *QLCDNumber) SetBinMode() {
	defer qt.Recovering("QLCDNumber::setBinMode")

	if ptr.Pointer() != nil {
		C.QLCDNumber_SetBinMode(ptr.Pointer())
	}
}

func (ptr *QLCDNumber) SetDecMode() {
	defer qt.Recovering("QLCDNumber::setDecMode")

	if ptr.Pointer() != nil {
		C.QLCDNumber_SetDecMode(ptr.Pointer())
	}
}

func (ptr *QLCDNumber) SetDigitCount(numDigits int) {
	defer qt.Recovering("QLCDNumber::setDigitCount")

	if ptr.Pointer() != nil {
		C.QLCDNumber_SetDigitCount(ptr.Pointer(), C.int(numDigits))
	}
}

func (ptr *QLCDNumber) SetHexMode() {
	defer qt.Recovering("QLCDNumber::setHexMode")

	if ptr.Pointer() != nil {
		C.QLCDNumber_SetHexMode(ptr.Pointer())
	}
}

func (ptr *QLCDNumber) SetOctMode() {
	defer qt.Recovering("QLCDNumber::setOctMode")

	if ptr.Pointer() != nil {
		C.QLCDNumber_SetOctMode(ptr.Pointer())
	}
}

func (ptr *QLCDNumber) SizeHint() *core.QSize {
	defer qt.Recovering("QLCDNumber::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QLCDNumber_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLCDNumber) DestroyQLCDNumber() {
	defer qt.Recovering("QLCDNumber::~QLCDNumber")

	if ptr.Pointer() != nil {
		C.QLCDNumber_DestroyQLCDNumber(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QLCDNumber) ConnectChangeEvent(f func(ev *core.QEvent)) {
	defer qt.Recovering("connect QLCDNumber::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QLCDNumber) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QLCDNumber::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQLCDNumberChangeEvent
func callbackQLCDNumberChangeEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QLCDNumber::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QLCDNumber) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QLCDNumber::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QLCDNumber) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QLCDNumber::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQLCDNumberActionEvent
func callbackQLCDNumberActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLCDNumber::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QLCDNumber) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QLCDNumber::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QLCDNumber) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QLCDNumber::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQLCDNumberDragEnterEvent
func callbackQLCDNumberDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLCDNumber::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QLCDNumber) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QLCDNumber::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QLCDNumber) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QLCDNumber::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQLCDNumberDragLeaveEvent
func callbackQLCDNumberDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLCDNumber::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QLCDNumber) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QLCDNumber::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QLCDNumber) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QLCDNumber::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQLCDNumberDragMoveEvent
func callbackQLCDNumberDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLCDNumber::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QLCDNumber) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QLCDNumber::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QLCDNumber) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QLCDNumber::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQLCDNumberDropEvent
func callbackQLCDNumberDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLCDNumber::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QLCDNumber) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QLCDNumber::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QLCDNumber) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QLCDNumber::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQLCDNumberEnterEvent
func callbackQLCDNumberEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLCDNumber::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QLCDNumber) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QLCDNumber::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QLCDNumber) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QLCDNumber::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQLCDNumberFocusOutEvent
func callbackQLCDNumberFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLCDNumber::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QLCDNumber) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QLCDNumber::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QLCDNumber) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QLCDNumber::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQLCDNumberHideEvent
func callbackQLCDNumberHideEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLCDNumber::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QLCDNumber) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QLCDNumber::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QLCDNumber) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QLCDNumber::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQLCDNumberLeaveEvent
func callbackQLCDNumberLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLCDNumber::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QLCDNumber) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QLCDNumber::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QLCDNumber) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QLCDNumber::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQLCDNumberMoveEvent
func callbackQLCDNumberMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLCDNumber::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QLCDNumber) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QLCDNumber::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QLCDNumber) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QLCDNumber::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQLCDNumberSetVisible
func callbackQLCDNumberSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QLCDNumber::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QLCDNumber) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QLCDNumber::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QLCDNumber) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QLCDNumber::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQLCDNumberShowEvent
func callbackQLCDNumberShowEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLCDNumber::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QLCDNumber) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QLCDNumber::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QLCDNumber) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QLCDNumber::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQLCDNumberCloseEvent
func callbackQLCDNumberCloseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLCDNumber::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QLCDNumber) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QLCDNumber::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QLCDNumber) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QLCDNumber::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQLCDNumberContextMenuEvent
func callbackQLCDNumberContextMenuEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLCDNumber::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QLCDNumber) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QLCDNumber::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QLCDNumber) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QLCDNumber::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQLCDNumberInitPainter
func callbackQLCDNumberInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QLCDNumber::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

}

func (ptr *QLCDNumber) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QLCDNumber::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QLCDNumber) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QLCDNumber::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQLCDNumberInputMethodEvent
func callbackQLCDNumberInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLCDNumber::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QLCDNumber) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QLCDNumber::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QLCDNumber) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QLCDNumber::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQLCDNumberKeyPressEvent
func callbackQLCDNumberKeyPressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLCDNumber::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QLCDNumber) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QLCDNumber::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QLCDNumber) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QLCDNumber::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQLCDNumberKeyReleaseEvent
func callbackQLCDNumberKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLCDNumber::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QLCDNumber) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QLCDNumber::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QLCDNumber) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QLCDNumber::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQLCDNumberMouseDoubleClickEvent
func callbackQLCDNumberMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLCDNumber::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QLCDNumber) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QLCDNumber::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QLCDNumber) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QLCDNumber::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQLCDNumberMouseMoveEvent
func callbackQLCDNumberMouseMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLCDNumber::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QLCDNumber) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QLCDNumber::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QLCDNumber) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QLCDNumber::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQLCDNumberMousePressEvent
func callbackQLCDNumberMousePressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLCDNumber::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QLCDNumber) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QLCDNumber::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QLCDNumber) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QLCDNumber::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQLCDNumberMouseReleaseEvent
func callbackQLCDNumberMouseReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLCDNumber::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QLCDNumber) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QLCDNumber::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QLCDNumber) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QLCDNumber::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQLCDNumberResizeEvent
func callbackQLCDNumberResizeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLCDNumber::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QLCDNumber) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QLCDNumber::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QLCDNumber) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QLCDNumber::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQLCDNumberTabletEvent
func callbackQLCDNumberTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLCDNumber::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QLCDNumber) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QLCDNumber::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QLCDNumber) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QLCDNumber::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQLCDNumberWheelEvent
func callbackQLCDNumberWheelEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLCDNumber::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QLCDNumber) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QLCDNumber::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QLCDNumber) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QLCDNumber::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQLCDNumberTimerEvent
func callbackQLCDNumberTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLCDNumber::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QLCDNumber) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QLCDNumber::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QLCDNumber) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QLCDNumber::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQLCDNumberChildEvent
func callbackQLCDNumberChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLCDNumber::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QLCDNumber) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QLCDNumber::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QLCDNumber) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QLCDNumber::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQLCDNumberCustomEvent
func callbackQLCDNumberCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLCDNumber::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
