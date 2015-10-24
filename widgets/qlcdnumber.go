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

type QLCDNumberITF interface {
	QFrameITF
	QLCDNumberPTR() *QLCDNumber
}

func PointerFromQLCDNumber(ptr QLCDNumberITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLCDNumberPTR().Pointer()
	}
	return nil
}

func QLCDNumberFromPointer(ptr unsafe.Pointer) *QLCDNumber {
	var n = new(QLCDNumber)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QLCDNumber_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QLCDNumber) QLCDNumberPTR() *QLCDNumber {
	return ptr
}

//QLCDNumber::Mode
type QLCDNumber__Mode int

var (
	QLCDNumber__Hex = QLCDNumber__Mode(0)
	QLCDNumber__Dec = QLCDNumber__Mode(1)
	QLCDNumber__Oct = QLCDNumber__Mode(2)
	QLCDNumber__Bin = QLCDNumber__Mode(3)
)

//QLCDNumber::SegmentStyle
type QLCDNumber__SegmentStyle int

var (
	QLCDNumber__Outline = QLCDNumber__SegmentStyle(0)
	QLCDNumber__Filled  = QLCDNumber__SegmentStyle(1)
	QLCDNumber__Flat    = QLCDNumber__SegmentStyle(2)
)

func (ptr *QLCDNumber) IntValue() int {
	if ptr.Pointer() != nil {
		return int(C.QLCDNumber_IntValue(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLCDNumber) Mode() QLCDNumber__Mode {
	if ptr.Pointer() != nil {
		return QLCDNumber__Mode(C.QLCDNumber_Mode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLCDNumber) SegmentStyle() QLCDNumber__SegmentStyle {
	if ptr.Pointer() != nil {
		return QLCDNumber__SegmentStyle(C.QLCDNumber_SegmentStyle(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLCDNumber) SetMode(v QLCDNumber__Mode) {
	if ptr.Pointer() != nil {
		C.QLCDNumber_SetMode(C.QtObjectPtr(ptr.Pointer()), C.int(v))
	}
}

func (ptr *QLCDNumber) SetSegmentStyle(v QLCDNumber__SegmentStyle) {
	if ptr.Pointer() != nil {
		C.QLCDNumber_SetSegmentStyle(C.QtObjectPtr(ptr.Pointer()), C.int(v))
	}
}

func (ptr *QLCDNumber) SetSmallDecimalPoint(v bool) {
	if ptr.Pointer() != nil {
		C.QLCDNumber_SetSmallDecimalPoint(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QLCDNumber) SmallDecimalPoint() bool {
	if ptr.Pointer() != nil {
		return C.QLCDNumber_SmallDecimalPoint(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func NewQLCDNumber(parent QWidgetITF) *QLCDNumber {
	return QLCDNumberFromPointer(unsafe.Pointer(C.QLCDNumber_NewQLCDNumber(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QLCDNumber) CheckOverflow2(num int) bool {
	if ptr.Pointer() != nil {
		return C.QLCDNumber_CheckOverflow2(C.QtObjectPtr(ptr.Pointer()), C.int(num)) != 0
	}
	return false
}

func (ptr *QLCDNumber) DigitCount() int {
	if ptr.Pointer() != nil {
		return int(C.QLCDNumber_DigitCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLCDNumber) Display(s string) {
	if ptr.Pointer() != nil {
		C.QLCDNumber_Display(C.QtObjectPtr(ptr.Pointer()), C.CString(s))
	}
}

func (ptr *QLCDNumber) Display3(num int) {
	if ptr.Pointer() != nil {
		C.QLCDNumber_Display3(C.QtObjectPtr(ptr.Pointer()), C.int(num))
	}
}

func (ptr *QLCDNumber) ConnectOverflow(f func()) {
	if ptr.Pointer() != nil {
		C.QLCDNumber_ConnectOverflow(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "overflow", f)
	}
}

func (ptr *QLCDNumber) DisconnectOverflow() {
	if ptr.Pointer() != nil {
		C.QLCDNumber_DisconnectOverflow(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "overflow")
	}
}

//export callbackQLCDNumberOverflow
func callbackQLCDNumberOverflow(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "overflow").(func())()
}

func (ptr *QLCDNumber) SetBinMode() {
	if ptr.Pointer() != nil {
		C.QLCDNumber_SetBinMode(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QLCDNumber) SetDecMode() {
	if ptr.Pointer() != nil {
		C.QLCDNumber_SetDecMode(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QLCDNumber) SetDigitCount(numDigits int) {
	if ptr.Pointer() != nil {
		C.QLCDNumber_SetDigitCount(C.QtObjectPtr(ptr.Pointer()), C.int(numDigits))
	}
}

func (ptr *QLCDNumber) SetHexMode() {
	if ptr.Pointer() != nil {
		C.QLCDNumber_SetHexMode(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QLCDNumber) SetOctMode() {
	if ptr.Pointer() != nil {
		C.QLCDNumber_SetOctMode(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QLCDNumber) DestroyQLCDNumber() {
	if ptr.Pointer() != nil {
		C.QLCDNumber_DestroyQLCDNumber(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
