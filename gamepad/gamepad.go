// +build !minimal

package gamepad

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "gamepad.h"
import "C"
import (
	"fmt"
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"runtime"
	"unsafe"
)

func cGoUnpackString(s C.struct_QtGamepad_PackedString) string {
	if len := int(s.len); len == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}

type QGamepad struct {
	core.QObject
}

type QGamepad_ITF interface {
	core.QObject_ITF
	QGamepad_PTR() *QGamepad
}

func (ptr *QGamepad) QGamepad_PTR() *QGamepad {
	return ptr
}

func (ptr *QGamepad) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QGamepad) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQGamepad(ptr QGamepad_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGamepad_PTR().Pointer()
	}
	return nil
}

func NewQGamepadFromPointer(ptr unsafe.Pointer) *QGamepad {
	var n = new(QGamepad)
	n.SetPointer(ptr)
	return n
}
func NewQGamepad(deviceId int, parent core.QObject_ITF) *QGamepad {
	var tmpValue = NewQGamepadFromPointer(C.QGamepad_NewQGamepad(C.int(int32(deviceId)), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQGamepad_AxisLeftXChanged
func callbackQGamepad_AxisLeftXChanged(ptr unsafe.Pointer, value C.double) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "axisLeftXChanged"); signal != nil {
		signal.(func(float64))(float64(value))
	}

}

func (ptr *QGamepad) ConnectAxisLeftXChanged(f func(value float64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "axisLeftXChanged") {
			C.QGamepad_ConnectAxisLeftXChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "axisLeftXChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "axisLeftXChanged", func(value float64) {
				signal.(func(float64))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "axisLeftXChanged", f)
		}
	}
}

func (ptr *QGamepad) DisconnectAxisLeftXChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectAxisLeftXChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "axisLeftXChanged")
	}
}

func (ptr *QGamepad) AxisLeftXChanged(value float64) {
	if ptr.Pointer() != nil {
		C.QGamepad_AxisLeftXChanged(ptr.Pointer(), C.double(value))
	}
}

//export callbackQGamepad_AxisLeftYChanged
func callbackQGamepad_AxisLeftYChanged(ptr unsafe.Pointer, value C.double) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "axisLeftYChanged"); signal != nil {
		signal.(func(float64))(float64(value))
	}

}

func (ptr *QGamepad) ConnectAxisLeftYChanged(f func(value float64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "axisLeftYChanged") {
			C.QGamepad_ConnectAxisLeftYChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "axisLeftYChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "axisLeftYChanged", func(value float64) {
				signal.(func(float64))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "axisLeftYChanged", f)
		}
	}
}

func (ptr *QGamepad) DisconnectAxisLeftYChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectAxisLeftYChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "axisLeftYChanged")
	}
}

func (ptr *QGamepad) AxisLeftYChanged(value float64) {
	if ptr.Pointer() != nil {
		C.QGamepad_AxisLeftYChanged(ptr.Pointer(), C.double(value))
	}
}

//export callbackQGamepad_AxisRightXChanged
func callbackQGamepad_AxisRightXChanged(ptr unsafe.Pointer, value C.double) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "axisRightXChanged"); signal != nil {
		signal.(func(float64))(float64(value))
	}

}

func (ptr *QGamepad) ConnectAxisRightXChanged(f func(value float64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "axisRightXChanged") {
			C.QGamepad_ConnectAxisRightXChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "axisRightXChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "axisRightXChanged", func(value float64) {
				signal.(func(float64))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "axisRightXChanged", f)
		}
	}
}

func (ptr *QGamepad) DisconnectAxisRightXChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectAxisRightXChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "axisRightXChanged")
	}
}

func (ptr *QGamepad) AxisRightXChanged(value float64) {
	if ptr.Pointer() != nil {
		C.QGamepad_AxisRightXChanged(ptr.Pointer(), C.double(value))
	}
}

//export callbackQGamepad_AxisRightYChanged
func callbackQGamepad_AxisRightYChanged(ptr unsafe.Pointer, value C.double) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "axisRightYChanged"); signal != nil {
		signal.(func(float64))(float64(value))
	}

}

func (ptr *QGamepad) ConnectAxisRightYChanged(f func(value float64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "axisRightYChanged") {
			C.QGamepad_ConnectAxisRightYChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "axisRightYChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "axisRightYChanged", func(value float64) {
				signal.(func(float64))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "axisRightYChanged", f)
		}
	}
}

func (ptr *QGamepad) DisconnectAxisRightYChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectAxisRightYChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "axisRightYChanged")
	}
}

func (ptr *QGamepad) AxisRightYChanged(value float64) {
	if ptr.Pointer() != nil {
		C.QGamepad_AxisRightYChanged(ptr.Pointer(), C.double(value))
	}
}

//export callbackQGamepad_ButtonAChanged
func callbackQGamepad_ButtonAChanged(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "buttonAChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonAChanged(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "buttonAChanged") {
			C.QGamepad_ConnectButtonAChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "buttonAChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "buttonAChanged", func(value bool) {
				signal.(func(bool))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "buttonAChanged", f)
		}
	}
}

func (ptr *QGamepad) DisconnectButtonAChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonAChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "buttonAChanged")
	}
}

func (ptr *QGamepad) ButtonAChanged(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonAChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_ButtonBChanged
func callbackQGamepad_ButtonBChanged(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "buttonBChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonBChanged(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "buttonBChanged") {
			C.QGamepad_ConnectButtonBChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "buttonBChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "buttonBChanged", func(value bool) {
				signal.(func(bool))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "buttonBChanged", f)
		}
	}
}

func (ptr *QGamepad) DisconnectButtonBChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonBChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "buttonBChanged")
	}
}

func (ptr *QGamepad) ButtonBChanged(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonBChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_ButtonCenterChanged
func callbackQGamepad_ButtonCenterChanged(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "buttonCenterChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonCenterChanged(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "buttonCenterChanged") {
			C.QGamepad_ConnectButtonCenterChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "buttonCenterChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "buttonCenterChanged", func(value bool) {
				signal.(func(bool))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "buttonCenterChanged", f)
		}
	}
}

func (ptr *QGamepad) DisconnectButtonCenterChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonCenterChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "buttonCenterChanged")
	}
}

func (ptr *QGamepad) ButtonCenterChanged(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonCenterChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_ButtonDownChanged
func callbackQGamepad_ButtonDownChanged(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "buttonDownChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonDownChanged(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "buttonDownChanged") {
			C.QGamepad_ConnectButtonDownChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "buttonDownChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "buttonDownChanged", func(value bool) {
				signal.(func(bool))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "buttonDownChanged", f)
		}
	}
}

func (ptr *QGamepad) DisconnectButtonDownChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonDownChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "buttonDownChanged")
	}
}

func (ptr *QGamepad) ButtonDownChanged(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonDownChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_ButtonGuideChanged
func callbackQGamepad_ButtonGuideChanged(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "buttonGuideChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonGuideChanged(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "buttonGuideChanged") {
			C.QGamepad_ConnectButtonGuideChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "buttonGuideChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "buttonGuideChanged", func(value bool) {
				signal.(func(bool))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "buttonGuideChanged", f)
		}
	}
}

func (ptr *QGamepad) DisconnectButtonGuideChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonGuideChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "buttonGuideChanged")
	}
}

func (ptr *QGamepad) ButtonGuideChanged(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonGuideChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_ButtonL1Changed
func callbackQGamepad_ButtonL1Changed(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "buttonL1Changed"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonL1Changed(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "buttonL1Changed") {
			C.QGamepad_ConnectButtonL1Changed(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "buttonL1Changed"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "buttonL1Changed", func(value bool) {
				signal.(func(bool))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "buttonL1Changed", f)
		}
	}
}

func (ptr *QGamepad) DisconnectButtonL1Changed() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonL1Changed(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "buttonL1Changed")
	}
}

func (ptr *QGamepad) ButtonL1Changed(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonL1Changed(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_ButtonL2Changed
func callbackQGamepad_ButtonL2Changed(ptr unsafe.Pointer, value C.double) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "buttonL2Changed"); signal != nil {
		signal.(func(float64))(float64(value))
	}

}

func (ptr *QGamepad) ConnectButtonL2Changed(f func(value float64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "buttonL2Changed") {
			C.QGamepad_ConnectButtonL2Changed(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "buttonL2Changed"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "buttonL2Changed", func(value float64) {
				signal.(func(float64))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "buttonL2Changed", f)
		}
	}
}

func (ptr *QGamepad) DisconnectButtonL2Changed() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonL2Changed(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "buttonL2Changed")
	}
}

func (ptr *QGamepad) ButtonL2Changed(value float64) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonL2Changed(ptr.Pointer(), C.double(value))
	}
}

//export callbackQGamepad_ButtonL3Changed
func callbackQGamepad_ButtonL3Changed(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "buttonL3Changed"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonL3Changed(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "buttonL3Changed") {
			C.QGamepad_ConnectButtonL3Changed(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "buttonL3Changed"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "buttonL3Changed", func(value bool) {
				signal.(func(bool))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "buttonL3Changed", f)
		}
	}
}

func (ptr *QGamepad) DisconnectButtonL3Changed() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonL3Changed(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "buttonL3Changed")
	}
}

func (ptr *QGamepad) ButtonL3Changed(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonL3Changed(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_ButtonLeftChanged
func callbackQGamepad_ButtonLeftChanged(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "buttonLeftChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonLeftChanged(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "buttonLeftChanged") {
			C.QGamepad_ConnectButtonLeftChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "buttonLeftChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "buttonLeftChanged", func(value bool) {
				signal.(func(bool))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "buttonLeftChanged", f)
		}
	}
}

func (ptr *QGamepad) DisconnectButtonLeftChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonLeftChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "buttonLeftChanged")
	}
}

func (ptr *QGamepad) ButtonLeftChanged(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonLeftChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_ButtonR1Changed
func callbackQGamepad_ButtonR1Changed(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "buttonR1Changed"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonR1Changed(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "buttonR1Changed") {
			C.QGamepad_ConnectButtonR1Changed(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "buttonR1Changed"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "buttonR1Changed", func(value bool) {
				signal.(func(bool))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "buttonR1Changed", f)
		}
	}
}

func (ptr *QGamepad) DisconnectButtonR1Changed() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonR1Changed(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "buttonR1Changed")
	}
}

func (ptr *QGamepad) ButtonR1Changed(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonR1Changed(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_ButtonR2Changed
func callbackQGamepad_ButtonR2Changed(ptr unsafe.Pointer, value C.double) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "buttonR2Changed"); signal != nil {
		signal.(func(float64))(float64(value))
	}

}

func (ptr *QGamepad) ConnectButtonR2Changed(f func(value float64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "buttonR2Changed") {
			C.QGamepad_ConnectButtonR2Changed(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "buttonR2Changed"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "buttonR2Changed", func(value float64) {
				signal.(func(float64))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "buttonR2Changed", f)
		}
	}
}

func (ptr *QGamepad) DisconnectButtonR2Changed() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonR2Changed(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "buttonR2Changed")
	}
}

func (ptr *QGamepad) ButtonR2Changed(value float64) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonR2Changed(ptr.Pointer(), C.double(value))
	}
}

//export callbackQGamepad_ButtonR3Changed
func callbackQGamepad_ButtonR3Changed(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "buttonR3Changed"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonR3Changed(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "buttonR3Changed") {
			C.QGamepad_ConnectButtonR3Changed(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "buttonR3Changed"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "buttonR3Changed", func(value bool) {
				signal.(func(bool))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "buttonR3Changed", f)
		}
	}
}

func (ptr *QGamepad) DisconnectButtonR3Changed() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonR3Changed(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "buttonR3Changed")
	}
}

func (ptr *QGamepad) ButtonR3Changed(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonR3Changed(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_ButtonRightChanged
func callbackQGamepad_ButtonRightChanged(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "buttonRightChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonRightChanged(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "buttonRightChanged") {
			C.QGamepad_ConnectButtonRightChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "buttonRightChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "buttonRightChanged", func(value bool) {
				signal.(func(bool))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "buttonRightChanged", f)
		}
	}
}

func (ptr *QGamepad) DisconnectButtonRightChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonRightChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "buttonRightChanged")
	}
}

func (ptr *QGamepad) ButtonRightChanged(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonRightChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_ButtonSelectChanged
func callbackQGamepad_ButtonSelectChanged(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "buttonSelectChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonSelectChanged(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "buttonSelectChanged") {
			C.QGamepad_ConnectButtonSelectChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "buttonSelectChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "buttonSelectChanged", func(value bool) {
				signal.(func(bool))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "buttonSelectChanged", f)
		}
	}
}

func (ptr *QGamepad) DisconnectButtonSelectChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonSelectChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "buttonSelectChanged")
	}
}

func (ptr *QGamepad) ButtonSelectChanged(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonSelectChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_ButtonStartChanged
func callbackQGamepad_ButtonStartChanged(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "buttonStartChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonStartChanged(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "buttonStartChanged") {
			C.QGamepad_ConnectButtonStartChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "buttonStartChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "buttonStartChanged", func(value bool) {
				signal.(func(bool))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "buttonStartChanged", f)
		}
	}
}

func (ptr *QGamepad) DisconnectButtonStartChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonStartChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "buttonStartChanged")
	}
}

func (ptr *QGamepad) ButtonStartChanged(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonStartChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_ButtonUpChanged
func callbackQGamepad_ButtonUpChanged(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "buttonUpChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonUpChanged(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "buttonUpChanged") {
			C.QGamepad_ConnectButtonUpChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "buttonUpChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "buttonUpChanged", func(value bool) {
				signal.(func(bool))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "buttonUpChanged", f)
		}
	}
}

func (ptr *QGamepad) DisconnectButtonUpChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonUpChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "buttonUpChanged")
	}
}

func (ptr *QGamepad) ButtonUpChanged(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonUpChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_ButtonXChanged
func callbackQGamepad_ButtonXChanged(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "buttonXChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonXChanged(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "buttonXChanged") {
			C.QGamepad_ConnectButtonXChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "buttonXChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "buttonXChanged", func(value bool) {
				signal.(func(bool))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "buttonXChanged", f)
		}
	}
}

func (ptr *QGamepad) DisconnectButtonXChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonXChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "buttonXChanged")
	}
}

func (ptr *QGamepad) ButtonXChanged(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonXChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_ButtonYChanged
func callbackQGamepad_ButtonYChanged(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "buttonYChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonYChanged(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "buttonYChanged") {
			C.QGamepad_ConnectButtonYChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "buttonYChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "buttonYChanged", func(value bool) {
				signal.(func(bool))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "buttonYChanged", f)
		}
	}
}

func (ptr *QGamepad) DisconnectButtonYChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonYChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "buttonYChanged")
	}
}

func (ptr *QGamepad) ButtonYChanged(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonYChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_ConnectedChanged
func callbackQGamepad_ConnectedChanged(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "connectedChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectConnectedChanged(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "connectedChanged") {
			C.QGamepad_ConnectConnectedChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "connectedChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "connectedChanged", func(value bool) {
				signal.(func(bool))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "connectedChanged", f)
		}
	}
}

func (ptr *QGamepad) DisconnectConnectedChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectConnectedChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "connectedChanged")
	}
}

func (ptr *QGamepad) ConnectedChanged(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ConnectedChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_DeviceIdChanged
func callbackQGamepad_DeviceIdChanged(ptr unsafe.Pointer, value C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "deviceIdChanged"); signal != nil {
		signal.(func(int))(int(int32(value)))
	}

}

func (ptr *QGamepad) ConnectDeviceIdChanged(f func(value int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "deviceIdChanged") {
			C.QGamepad_ConnectDeviceIdChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "deviceIdChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "deviceIdChanged", func(value int) {
				signal.(func(int))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "deviceIdChanged", f)
		}
	}
}

func (ptr *QGamepad) DisconnectDeviceIdChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectDeviceIdChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "deviceIdChanged")
	}
}

func (ptr *QGamepad) DeviceIdChanged(value int) {
	if ptr.Pointer() != nil {
		C.QGamepad_DeviceIdChanged(ptr.Pointer(), C.int(int32(value)))
	}
}

//export callbackQGamepad_NameChanged
func callbackQGamepad_NameChanged(ptr unsafe.Pointer, value C.struct_QtGamepad_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "nameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(value))
	}

}

func (ptr *QGamepad) ConnectNameChanged(f func(value string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(fmt.Sprint(ptr.Pointer()), "nameChanged") {
			C.QGamepad_ConnectNameChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "nameChanged"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "nameChanged", func(value string) {
				signal.(func(string))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "nameChanged", f)
		}
	}
}

func (ptr *QGamepad) DisconnectNameChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectNameChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "nameChanged")
	}
}

func (ptr *QGamepad) NameChanged(value string) {
	if ptr.Pointer() != nil {
		var valueC *C.char
		if value != "" {
			valueC = C.CString(value)
			defer C.free(unsafe.Pointer(valueC))
		}
		C.QGamepad_NameChanged(ptr.Pointer(), valueC)
	}
}

//export callbackQGamepad_SetDeviceId
func callbackQGamepad_SetDeviceId(ptr unsafe.Pointer, number C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "setDeviceId"); signal != nil {
		signal.(func(int))(int(int32(number)))
	} else {
		NewQGamepadFromPointer(ptr).SetDeviceIdDefault(int(int32(number)))
	}
}

func (ptr *QGamepad) ConnectSetDeviceId(f func(number int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(fmt.Sprint(ptr.Pointer()), "setDeviceId"); signal != nil {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "setDeviceId", func(number int) {
				signal.(func(int))(number)
				f(number)
			})
		} else {
			qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "setDeviceId", f)
		}
	}
}

func (ptr *QGamepad) DisconnectSetDeviceId() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "setDeviceId")
	}
}

func (ptr *QGamepad) SetDeviceId(number int) {
	if ptr.Pointer() != nil {
		C.QGamepad_SetDeviceId(ptr.Pointer(), C.int(int32(number)))
	}
}

func (ptr *QGamepad) SetDeviceIdDefault(number int) {
	if ptr.Pointer() != nil {
		C.QGamepad_SetDeviceIdDefault(ptr.Pointer(), C.int(int32(number)))
	}
}

func (ptr *QGamepad) DestroyQGamepad() {
	if ptr.Pointer() != nil {
		C.QGamepad_DestroyQGamepad(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QGamepad) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGamepad_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGamepad) ButtonA() bool {
	if ptr.Pointer() != nil {
		return C.QGamepad_ButtonA(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGamepad) ButtonB() bool {
	if ptr.Pointer() != nil {
		return C.QGamepad_ButtonB(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGamepad) ButtonCenter() bool {
	if ptr.Pointer() != nil {
		return C.QGamepad_ButtonCenter(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGamepad) ButtonDown() bool {
	if ptr.Pointer() != nil {
		return C.QGamepad_ButtonDown(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGamepad) ButtonGuide() bool {
	if ptr.Pointer() != nil {
		return C.QGamepad_ButtonGuide(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGamepad) ButtonL1() bool {
	if ptr.Pointer() != nil {
		return C.QGamepad_ButtonL1(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGamepad) ButtonL3() bool {
	if ptr.Pointer() != nil {
		return C.QGamepad_ButtonL3(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGamepad) ButtonLeft() bool {
	if ptr.Pointer() != nil {
		return C.QGamepad_ButtonLeft(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGamepad) ButtonR1() bool {
	if ptr.Pointer() != nil {
		return C.QGamepad_ButtonR1(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGamepad) ButtonR3() bool {
	if ptr.Pointer() != nil {
		return C.QGamepad_ButtonR3(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGamepad) ButtonRight() bool {
	if ptr.Pointer() != nil {
		return C.QGamepad_ButtonRight(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGamepad) ButtonSelect() bool {
	if ptr.Pointer() != nil {
		return C.QGamepad_ButtonSelect(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGamepad) ButtonStart() bool {
	if ptr.Pointer() != nil {
		return C.QGamepad_ButtonStart(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGamepad) ButtonUp() bool {
	if ptr.Pointer() != nil {
		return C.QGamepad_ButtonUp(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGamepad) ButtonX() bool {
	if ptr.Pointer() != nil {
		return C.QGamepad_ButtonX(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGamepad) ButtonY() bool {
	if ptr.Pointer() != nil {
		return C.QGamepad_ButtonY(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGamepad) IsConnected() bool {
	if ptr.Pointer() != nil {
		return C.QGamepad_IsConnected(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGamepad) AxisLeftX() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGamepad_AxisLeftX(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGamepad) AxisLeftY() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGamepad_AxisLeftY(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGamepad) AxisRightX() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGamepad_AxisRightX(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGamepad) AxisRightY() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGamepad_AxisRightY(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGamepad) ButtonL2() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGamepad_ButtonL2(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGamepad) ButtonR2() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGamepad_ButtonR2(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGamepad) DeviceId() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QGamepad_DeviceId(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGamepad) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QGamepad___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QGamepad) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QGamepad___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QGamepad) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QGamepad___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QGamepad) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QGamepad___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGamepad) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGamepad___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGamepad) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QGamepad___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QGamepad) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QGamepad___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGamepad) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGamepad___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGamepad) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.QGamepad___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QGamepad) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QGamepad___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGamepad) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGamepad___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGamepad) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QGamepad___findChildren_newList(ptr.Pointer()))
}

func (ptr *QGamepad) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QGamepad___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGamepad) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGamepad___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGamepad) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QGamepad___children_newList(ptr.Pointer()))
}

//export callbackQGamepad_Event
func callbackQGamepad_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGamepadFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QGamepad) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGamepad_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQGamepad_EventFilter
func callbackQGamepad_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGamepadFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QGamepad) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGamepad_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQGamepad_ChildEvent
func callbackQGamepad_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQGamepadFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QGamepad) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGamepad_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQGamepad_ConnectNotify
func callbackQGamepad_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGamepadFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGamepad) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGamepad_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGamepad_CustomEvent
func callbackQGamepad_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGamepadFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGamepad) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGamepad_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQGamepad_DeleteLater
func callbackQGamepad_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQGamepadFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QGamepad) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QGamepad_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQGamepad_Destroyed
func callbackQGamepad_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQGamepad_DisconnectNotify
func callbackQGamepad_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGamepadFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGamepad) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGamepad_ObjectNameChanged
func callbackQGamepad_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtGamepad_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQGamepad_TimerEvent
func callbackQGamepad_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGamepadFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGamepad) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGamepad_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQGamepad_MetaObject
func callbackQGamepad_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQGamepadFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QGamepad) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QGamepad_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}
