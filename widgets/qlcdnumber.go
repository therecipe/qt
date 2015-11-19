package widgets

//#include "qlcdnumber.h"
import "C"
import (
	"github.com/therecipe/qt"
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
	if n.ObjectName() == "" {
		n.SetObjectName("QLCDNumber_" + qt.RandomIdentifier())
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
	if ptr.Pointer() != nil {
		return int(C.QLCDNumber_IntValue(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLCDNumber) Mode() QLCDNumber__Mode {
	if ptr.Pointer() != nil {
		return QLCDNumber__Mode(C.QLCDNumber_Mode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLCDNumber) SegmentStyle() QLCDNumber__SegmentStyle {
	if ptr.Pointer() != nil {
		return QLCDNumber__SegmentStyle(C.QLCDNumber_SegmentStyle(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLCDNumber) SetMode(v QLCDNumber__Mode) {
	if ptr.Pointer() != nil {
		C.QLCDNumber_SetMode(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QLCDNumber) SetSegmentStyle(v QLCDNumber__SegmentStyle) {
	if ptr.Pointer() != nil {
		C.QLCDNumber_SetSegmentStyle(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QLCDNumber) SetSmallDecimalPoint(v bool) {
	if ptr.Pointer() != nil {
		C.QLCDNumber_SetSmallDecimalPoint(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QLCDNumber) SmallDecimalPoint() bool {
	if ptr.Pointer() != nil {
		return C.QLCDNumber_SmallDecimalPoint(ptr.Pointer()) != 0
	}
	return false
}

func NewQLCDNumber(parent QWidget_ITF) *QLCDNumber {
	return NewQLCDNumberFromPointer(C.QLCDNumber_NewQLCDNumber(PointerFromQWidget(parent)))
}

func (ptr *QLCDNumber) CheckOverflow2(num int) bool {
	if ptr.Pointer() != nil {
		return C.QLCDNumber_CheckOverflow2(ptr.Pointer(), C.int(num)) != 0
	}
	return false
}

func (ptr *QLCDNumber) DigitCount() int {
	if ptr.Pointer() != nil {
		return int(C.QLCDNumber_DigitCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLCDNumber) Display(s string) {
	if ptr.Pointer() != nil {
		C.QLCDNumber_Display(ptr.Pointer(), C.CString(s))
	}
}

func (ptr *QLCDNumber) Display3(num int) {
	if ptr.Pointer() != nil {
		C.QLCDNumber_Display3(ptr.Pointer(), C.int(num))
	}
}

func (ptr *QLCDNumber) ConnectOverflow(f func()) {
	if ptr.Pointer() != nil {
		C.QLCDNumber_ConnectOverflow(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "overflow", f)
	}
}

func (ptr *QLCDNumber) DisconnectOverflow() {
	if ptr.Pointer() != nil {
		C.QLCDNumber_DisconnectOverflow(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "overflow")
	}
}

//export callbackQLCDNumberOverflow
func callbackQLCDNumberOverflow(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "overflow").(func())()
}

func (ptr *QLCDNumber) SetBinMode() {
	if ptr.Pointer() != nil {
		C.QLCDNumber_SetBinMode(ptr.Pointer())
	}
}

func (ptr *QLCDNumber) SetDecMode() {
	if ptr.Pointer() != nil {
		C.QLCDNumber_SetDecMode(ptr.Pointer())
	}
}

func (ptr *QLCDNumber) SetDigitCount(numDigits int) {
	if ptr.Pointer() != nil {
		C.QLCDNumber_SetDigitCount(ptr.Pointer(), C.int(numDigits))
	}
}

func (ptr *QLCDNumber) SetHexMode() {
	if ptr.Pointer() != nil {
		C.QLCDNumber_SetHexMode(ptr.Pointer())
	}
}

func (ptr *QLCDNumber) SetOctMode() {
	if ptr.Pointer() != nil {
		C.QLCDNumber_SetOctMode(ptr.Pointer())
	}
}

func (ptr *QLCDNumber) DestroyQLCDNumber() {
	if ptr.Pointer() != nil {
		C.QLCDNumber_DestroyQLCDNumber(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
