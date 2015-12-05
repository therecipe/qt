package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLCDNumber::intValue")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QLCDNumber_IntValue(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLCDNumber) Mode() QLCDNumber__Mode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLCDNumber::mode")
		}
	}()

	if ptr.Pointer() != nil {
		return QLCDNumber__Mode(C.QLCDNumber_Mode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLCDNumber) SegmentStyle() QLCDNumber__SegmentStyle {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLCDNumber::segmentStyle")
		}
	}()

	if ptr.Pointer() != nil {
		return QLCDNumber__SegmentStyle(C.QLCDNumber_SegmentStyle(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLCDNumber) SetMode(v QLCDNumber__Mode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLCDNumber::setMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLCDNumber_SetMode(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QLCDNumber) SetSegmentStyle(v QLCDNumber__SegmentStyle) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLCDNumber::setSegmentStyle")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLCDNumber_SetSegmentStyle(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QLCDNumber) SetSmallDecimalPoint(v bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLCDNumber::setSmallDecimalPoint")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLCDNumber_SetSmallDecimalPoint(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QLCDNumber) SmallDecimalPoint() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLCDNumber::smallDecimalPoint")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QLCDNumber_SmallDecimalPoint(ptr.Pointer()) != 0
	}
	return false
}

func NewQLCDNumber(parent QWidget_ITF) *QLCDNumber {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLCDNumber::QLCDNumber")
		}
	}()

	return NewQLCDNumberFromPointer(C.QLCDNumber_NewQLCDNumber(PointerFromQWidget(parent)))
}

func (ptr *QLCDNumber) CheckOverflow2(num int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLCDNumber::checkOverflow")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QLCDNumber_CheckOverflow2(ptr.Pointer(), C.int(num)) != 0
	}
	return false
}

func (ptr *QLCDNumber) DigitCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLCDNumber::digitCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QLCDNumber_DigitCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLCDNumber) Display(s string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLCDNumber::display")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLCDNumber_Display(ptr.Pointer(), C.CString(s))
	}
}

func (ptr *QLCDNumber) Display3(num int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLCDNumber::display")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLCDNumber_Display3(ptr.Pointer(), C.int(num))
	}
}

func (ptr *QLCDNumber) ConnectOverflow(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLCDNumber::overflow")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLCDNumber_ConnectOverflow(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "overflow", f)
	}
}

func (ptr *QLCDNumber) DisconnectOverflow() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLCDNumber::overflow")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLCDNumber_DisconnectOverflow(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "overflow")
	}
}

//export callbackQLCDNumberOverflow
func callbackQLCDNumberOverflow(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLCDNumber::overflow")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "overflow").(func())()
}

func (ptr *QLCDNumber) SetBinMode() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLCDNumber::setBinMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLCDNumber_SetBinMode(ptr.Pointer())
	}
}

func (ptr *QLCDNumber) SetDecMode() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLCDNumber::setDecMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLCDNumber_SetDecMode(ptr.Pointer())
	}
}

func (ptr *QLCDNumber) SetDigitCount(numDigits int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLCDNumber::setDigitCount")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLCDNumber_SetDigitCount(ptr.Pointer(), C.int(numDigits))
	}
}

func (ptr *QLCDNumber) SetHexMode() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLCDNumber::setHexMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLCDNumber_SetHexMode(ptr.Pointer())
	}
}

func (ptr *QLCDNumber) SetOctMode() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLCDNumber::setOctMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLCDNumber_SetOctMode(ptr.Pointer())
	}
}

func (ptr *QLCDNumber) DestroyQLCDNumber() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLCDNumber::~QLCDNumber")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLCDNumber_DestroyQLCDNumber(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
