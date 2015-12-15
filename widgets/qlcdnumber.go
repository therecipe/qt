package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
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

	var signal = qt.GetSignal(C.GoString(ptrName), "overflow")
	if signal != nil {
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

	var signal = qt.GetSignal(C.GoString(ptrName), "paintEvent")
	if signal != nil {
		defer signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(v))
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

func (ptr *QLCDNumber) DestroyQLCDNumber() {
	defer qt.Recovering("QLCDNumber::~QLCDNumber")

	if ptr.Pointer() != nil {
		C.QLCDNumber_DestroyQLCDNumber(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
