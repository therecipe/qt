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

func (ptr *QLCDNumber) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QLCDNumber::event")

	if ptr.Pointer() != nil {
		return C.QLCDNumber_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
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
func callbackQLCDNumberOverflow(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QLCDNumber::overflow")

	if signal := qt.GetSignal(C.GoString(ptrName), "overflow"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QLCDNumber) Overflow() {
	defer qt.Recovering("QLCDNumber::overflow")

	if ptr.Pointer() != nil {
		C.QLCDNumber_Overflow(ptr.Pointer())
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
func callbackQLCDNumberPaintEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QLCDNumber::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(v))
	} else {
		NewQLCDNumberFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(v))
	}
}

func (ptr *QLCDNumber) PaintEvent(v gui.QPaintEvent_ITF) {
	defer qt.Recovering("QLCDNumber::paintEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(v))
	}
}

func (ptr *QLCDNumber) PaintEventDefault(v gui.QPaintEvent_ITF) {
	defer qt.Recovering("QLCDNumber::paintEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(v))
	}
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
func callbackQLCDNumberChangeEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QLCDNumber::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
	} else {
		NewQLCDNumberFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(ev))
	}
}

func (ptr *QLCDNumber) ChangeEvent(ev core.QEvent_ITF) {
	defer qt.Recovering("QLCDNumber::changeEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QLCDNumber) ChangeEventDefault(ev core.QEvent_ITF) {
	defer qt.Recovering("QLCDNumber::changeEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
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
func callbackQLCDNumberActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLCDNumber::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQLCDNumberFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QLCDNumber) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QLCDNumber::actionEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QLCDNumber) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QLCDNumber::actionEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
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
func callbackQLCDNumberDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLCDNumber::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQLCDNumberFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QLCDNumber) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QLCDNumber::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QLCDNumber) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QLCDNumber::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
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
func callbackQLCDNumberDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLCDNumber::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQLCDNumberFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QLCDNumber) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QLCDNumber::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QLCDNumber) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QLCDNumber::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
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
func callbackQLCDNumberDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLCDNumber::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQLCDNumberFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QLCDNumber) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QLCDNumber::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QLCDNumber) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QLCDNumber::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
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
func callbackQLCDNumberDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLCDNumber::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQLCDNumberFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QLCDNumber) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QLCDNumber::dropEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QLCDNumber) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QLCDNumber::dropEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
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
func callbackQLCDNumberEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLCDNumber::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQLCDNumberFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QLCDNumber) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QLCDNumber::enterEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QLCDNumber) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QLCDNumber::enterEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QLCDNumber) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QLCDNumber::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QLCDNumber) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QLCDNumber::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQLCDNumberFocusInEvent
func callbackQLCDNumberFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLCDNumber::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQLCDNumberFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QLCDNumber) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QLCDNumber::focusInEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QLCDNumber) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QLCDNumber::focusInEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
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
func callbackQLCDNumberFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLCDNumber::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQLCDNumberFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QLCDNumber) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QLCDNumber::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QLCDNumber) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QLCDNumber::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
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
func callbackQLCDNumberHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLCDNumber::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQLCDNumberFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QLCDNumber) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QLCDNumber::hideEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QLCDNumber) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QLCDNumber::hideEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
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
func callbackQLCDNumberLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLCDNumber::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQLCDNumberFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QLCDNumber) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QLCDNumber::leaveEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QLCDNumber) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QLCDNumber::leaveEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
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
func callbackQLCDNumberMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLCDNumber::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQLCDNumberFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QLCDNumber) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QLCDNumber::moveEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QLCDNumber) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QLCDNumber::moveEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
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
func callbackQLCDNumberSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QLCDNumber::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QLCDNumber) SetVisible(visible bool) {
	defer qt.Recovering("QLCDNumber::setVisible")

	if ptr.Pointer() != nil {
		C.QLCDNumber_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QLCDNumber) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QLCDNumber::setVisible")

	if ptr.Pointer() != nil {
		C.QLCDNumber_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
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
func callbackQLCDNumberShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLCDNumber::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQLCDNumberFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QLCDNumber) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QLCDNumber::showEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QLCDNumber) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QLCDNumber::showEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
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
func callbackQLCDNumberCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLCDNumber::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQLCDNumberFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QLCDNumber) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QLCDNumber::closeEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QLCDNumber) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QLCDNumber::closeEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
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
func callbackQLCDNumberContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLCDNumber::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQLCDNumberFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QLCDNumber) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QLCDNumber::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QLCDNumber) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QLCDNumber::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
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
func callbackQLCDNumberInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QLCDNumber::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQLCDNumberFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QLCDNumber) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QLCDNumber::initPainter")

	if ptr.Pointer() != nil {
		C.QLCDNumber_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QLCDNumber) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QLCDNumber::initPainter")

	if ptr.Pointer() != nil {
		C.QLCDNumber_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
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
func callbackQLCDNumberInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLCDNumber::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQLCDNumberFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QLCDNumber) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QLCDNumber::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QLCDNumber) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QLCDNumber::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
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
func callbackQLCDNumberKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLCDNumber::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQLCDNumberFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QLCDNumber) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QLCDNumber::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QLCDNumber) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QLCDNumber::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
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
func callbackQLCDNumberKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLCDNumber::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQLCDNumberFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QLCDNumber) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QLCDNumber::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QLCDNumber) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QLCDNumber::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
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
func callbackQLCDNumberMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLCDNumber::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQLCDNumberFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QLCDNumber) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QLCDNumber::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QLCDNumber) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QLCDNumber::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
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
func callbackQLCDNumberMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLCDNumber::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQLCDNumberFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QLCDNumber) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QLCDNumber::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QLCDNumber) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QLCDNumber::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
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
func callbackQLCDNumberMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLCDNumber::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQLCDNumberFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QLCDNumber) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QLCDNumber::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QLCDNumber) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QLCDNumber::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
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
func callbackQLCDNumberMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLCDNumber::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQLCDNumberFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QLCDNumber) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QLCDNumber::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QLCDNumber) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QLCDNumber::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
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
func callbackQLCDNumberResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLCDNumber::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQLCDNumberFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QLCDNumber) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QLCDNumber::resizeEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QLCDNumber) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QLCDNumber::resizeEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
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
func callbackQLCDNumberTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLCDNumber::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQLCDNumberFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QLCDNumber) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QLCDNumber::tabletEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QLCDNumber) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QLCDNumber::tabletEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
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
func callbackQLCDNumberWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLCDNumber::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQLCDNumberFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QLCDNumber) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QLCDNumber::wheelEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QLCDNumber) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QLCDNumber::wheelEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
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
func callbackQLCDNumberTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLCDNumber::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQLCDNumberFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QLCDNumber) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QLCDNumber::timerEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QLCDNumber) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QLCDNumber::timerEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
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
func callbackQLCDNumberChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLCDNumber::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQLCDNumberFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QLCDNumber) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QLCDNumber::childEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QLCDNumber) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QLCDNumber::childEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
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
func callbackQLCDNumberCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLCDNumber::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQLCDNumberFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QLCDNumber) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QLCDNumber::customEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QLCDNumber) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QLCDNumber::customEvent")

	if ptr.Pointer() != nil {
		C.QLCDNumber_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
