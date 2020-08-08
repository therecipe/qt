// +build !minimal

package gamepad

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "gamepad.h"
import "C"
import (
	"github.com/StarAurryon/qt"
	"github.com/StarAurryon/qt/core"
	"strings"
	"unsafe"
)

func cGoFreePacked(ptr unsafe.Pointer) { core.NewQByteArrayFromPointer(ptr).DestroyQByteArray() }
func cGoUnpackString(s C.struct_QtGamepad_PackedString) string {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtGamepad_PackedString) []byte {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		gs := C.GoString(s.data)
		return []byte(gs)
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}
func unpackStringList(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}
	return strings.Split(s, "¡¦!")
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

func NewQGamepadFromPointer(ptr unsafe.Pointer) (n *QGamepad) {
	n = new(QGamepad)
	n.SetPointer(ptr)
	return
}
func NewQGamepad(deviceId int, parent core.QObject_ITF) *QGamepad {
	tmpValue := NewQGamepadFromPointer(C.QGamepad_NewQGamepad(C.int(int32(deviceId)), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QGamepad) AxisLeftX() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGamepad_AxisLeftX(ptr.Pointer()))
	}
	return 0
}

//export callbackQGamepad_AxisLeftXChanged
func callbackQGamepad_AxisLeftXChanged(ptr unsafe.Pointer, value C.double) {
	if signal := qt.GetSignal(ptr, "axisLeftXChanged"); signal != nil {
		(*(*func(float64))(signal))(float64(value))
	}

}

func (ptr *QGamepad) ConnectAxisLeftXChanged(f func(value float64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "axisLeftXChanged") {
			C.QGamepad_ConnectAxisLeftXChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "axisLeftXChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "axisLeftXChanged"); signal != nil {
			f := func(value float64) {
				(*(*func(float64))(signal))(value)
				f(value)
			}
			qt.ConnectSignal(ptr.Pointer(), "axisLeftXChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "axisLeftXChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepad) DisconnectAxisLeftXChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectAxisLeftXChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "axisLeftXChanged")
	}
}

func (ptr *QGamepad) AxisLeftXChanged(value float64) {
	if ptr.Pointer() != nil {
		C.QGamepad_AxisLeftXChanged(ptr.Pointer(), C.double(value))
	}
}

func (ptr *QGamepad) AxisLeftY() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGamepad_AxisLeftY(ptr.Pointer()))
	}
	return 0
}

//export callbackQGamepad_AxisLeftYChanged
func callbackQGamepad_AxisLeftYChanged(ptr unsafe.Pointer, value C.double) {
	if signal := qt.GetSignal(ptr, "axisLeftYChanged"); signal != nil {
		(*(*func(float64))(signal))(float64(value))
	}

}

func (ptr *QGamepad) ConnectAxisLeftYChanged(f func(value float64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "axisLeftYChanged") {
			C.QGamepad_ConnectAxisLeftYChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "axisLeftYChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "axisLeftYChanged"); signal != nil {
			f := func(value float64) {
				(*(*func(float64))(signal))(value)
				f(value)
			}
			qt.ConnectSignal(ptr.Pointer(), "axisLeftYChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "axisLeftYChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepad) DisconnectAxisLeftYChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectAxisLeftYChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "axisLeftYChanged")
	}
}

func (ptr *QGamepad) AxisLeftYChanged(value float64) {
	if ptr.Pointer() != nil {
		C.QGamepad_AxisLeftYChanged(ptr.Pointer(), C.double(value))
	}
}

func (ptr *QGamepad) AxisRightX() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGamepad_AxisRightX(ptr.Pointer()))
	}
	return 0
}

//export callbackQGamepad_AxisRightXChanged
func callbackQGamepad_AxisRightXChanged(ptr unsafe.Pointer, value C.double) {
	if signal := qt.GetSignal(ptr, "axisRightXChanged"); signal != nil {
		(*(*func(float64))(signal))(float64(value))
	}

}

func (ptr *QGamepad) ConnectAxisRightXChanged(f func(value float64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "axisRightXChanged") {
			C.QGamepad_ConnectAxisRightXChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "axisRightXChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "axisRightXChanged"); signal != nil {
			f := func(value float64) {
				(*(*func(float64))(signal))(value)
				f(value)
			}
			qt.ConnectSignal(ptr.Pointer(), "axisRightXChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "axisRightXChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepad) DisconnectAxisRightXChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectAxisRightXChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "axisRightXChanged")
	}
}

func (ptr *QGamepad) AxisRightXChanged(value float64) {
	if ptr.Pointer() != nil {
		C.QGamepad_AxisRightXChanged(ptr.Pointer(), C.double(value))
	}
}

func (ptr *QGamepad) AxisRightY() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGamepad_AxisRightY(ptr.Pointer()))
	}
	return 0
}

//export callbackQGamepad_AxisRightYChanged
func callbackQGamepad_AxisRightYChanged(ptr unsafe.Pointer, value C.double) {
	if signal := qt.GetSignal(ptr, "axisRightYChanged"); signal != nil {
		(*(*func(float64))(signal))(float64(value))
	}

}

func (ptr *QGamepad) ConnectAxisRightYChanged(f func(value float64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "axisRightYChanged") {
			C.QGamepad_ConnectAxisRightYChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "axisRightYChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "axisRightYChanged"); signal != nil {
			f := func(value float64) {
				(*(*func(float64))(signal))(value)
				f(value)
			}
			qt.ConnectSignal(ptr.Pointer(), "axisRightYChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "axisRightYChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepad) DisconnectAxisRightYChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectAxisRightYChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "axisRightYChanged")
	}
}

func (ptr *QGamepad) AxisRightYChanged(value float64) {
	if ptr.Pointer() != nil {
		C.QGamepad_AxisRightYChanged(ptr.Pointer(), C.double(value))
	}
}

func (ptr *QGamepad) ButtonA() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepad_ButtonA(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQGamepad_ButtonAChanged
func callbackQGamepad_ButtonAChanged(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(ptr, "buttonAChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonAChanged(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonAChanged") {
			C.QGamepad_ConnectButtonAChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "buttonAChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonAChanged"); signal != nil {
			f := func(value bool) {
				(*(*func(bool))(signal))(value)
				f(value)
			}
			qt.ConnectSignal(ptr.Pointer(), "buttonAChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonAChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepad) DisconnectButtonAChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonAChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "buttonAChanged")
	}
}

func (ptr *QGamepad) ButtonAChanged(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonAChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

func (ptr *QGamepad) ButtonB() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepad_ButtonB(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQGamepad_ButtonBChanged
func callbackQGamepad_ButtonBChanged(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(ptr, "buttonBChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonBChanged(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonBChanged") {
			C.QGamepad_ConnectButtonBChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "buttonBChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonBChanged"); signal != nil {
			f := func(value bool) {
				(*(*func(bool))(signal))(value)
				f(value)
			}
			qt.ConnectSignal(ptr.Pointer(), "buttonBChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonBChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepad) DisconnectButtonBChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonBChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "buttonBChanged")
	}
}

func (ptr *QGamepad) ButtonBChanged(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonBChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

func (ptr *QGamepad) ButtonCenter() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepad_ButtonCenter(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQGamepad_ButtonCenterChanged
func callbackQGamepad_ButtonCenterChanged(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(ptr, "buttonCenterChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonCenterChanged(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonCenterChanged") {
			C.QGamepad_ConnectButtonCenterChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "buttonCenterChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonCenterChanged"); signal != nil {
			f := func(value bool) {
				(*(*func(bool))(signal))(value)
				f(value)
			}
			qt.ConnectSignal(ptr.Pointer(), "buttonCenterChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonCenterChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepad) DisconnectButtonCenterChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonCenterChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "buttonCenterChanged")
	}
}

func (ptr *QGamepad) ButtonCenterChanged(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonCenterChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

func (ptr *QGamepad) ButtonDown() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepad_ButtonDown(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQGamepad_ButtonDownChanged
func callbackQGamepad_ButtonDownChanged(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(ptr, "buttonDownChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonDownChanged(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonDownChanged") {
			C.QGamepad_ConnectButtonDownChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "buttonDownChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonDownChanged"); signal != nil {
			f := func(value bool) {
				(*(*func(bool))(signal))(value)
				f(value)
			}
			qt.ConnectSignal(ptr.Pointer(), "buttonDownChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonDownChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepad) DisconnectButtonDownChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonDownChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "buttonDownChanged")
	}
}

func (ptr *QGamepad) ButtonDownChanged(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonDownChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

func (ptr *QGamepad) ButtonGuide() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepad_ButtonGuide(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQGamepad_ButtonGuideChanged
func callbackQGamepad_ButtonGuideChanged(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(ptr, "buttonGuideChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonGuideChanged(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonGuideChanged") {
			C.QGamepad_ConnectButtonGuideChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "buttonGuideChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonGuideChanged"); signal != nil {
			f := func(value bool) {
				(*(*func(bool))(signal))(value)
				f(value)
			}
			qt.ConnectSignal(ptr.Pointer(), "buttonGuideChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonGuideChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepad) DisconnectButtonGuideChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonGuideChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "buttonGuideChanged")
	}
}

func (ptr *QGamepad) ButtonGuideChanged(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonGuideChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

func (ptr *QGamepad) ButtonL1() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepad_ButtonL1(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQGamepad_ButtonL1Changed
func callbackQGamepad_ButtonL1Changed(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(ptr, "buttonL1Changed"); signal != nil {
		(*(*func(bool))(signal))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonL1Changed(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonL1Changed") {
			C.QGamepad_ConnectButtonL1Changed(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "buttonL1Changed")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonL1Changed"); signal != nil {
			f := func(value bool) {
				(*(*func(bool))(signal))(value)
				f(value)
			}
			qt.ConnectSignal(ptr.Pointer(), "buttonL1Changed", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonL1Changed", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepad) DisconnectButtonL1Changed() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonL1Changed(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "buttonL1Changed")
	}
}

func (ptr *QGamepad) ButtonL1Changed(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonL1Changed(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

func (ptr *QGamepad) ButtonL2() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGamepad_ButtonL2(ptr.Pointer()))
	}
	return 0
}

//export callbackQGamepad_ButtonL2Changed
func callbackQGamepad_ButtonL2Changed(ptr unsafe.Pointer, value C.double) {
	if signal := qt.GetSignal(ptr, "buttonL2Changed"); signal != nil {
		(*(*func(float64))(signal))(float64(value))
	}

}

func (ptr *QGamepad) ConnectButtonL2Changed(f func(value float64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonL2Changed") {
			C.QGamepad_ConnectButtonL2Changed(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "buttonL2Changed")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonL2Changed"); signal != nil {
			f := func(value float64) {
				(*(*func(float64))(signal))(value)
				f(value)
			}
			qt.ConnectSignal(ptr.Pointer(), "buttonL2Changed", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonL2Changed", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepad) DisconnectButtonL2Changed() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonL2Changed(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "buttonL2Changed")
	}
}

func (ptr *QGamepad) ButtonL2Changed(value float64) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonL2Changed(ptr.Pointer(), C.double(value))
	}
}

func (ptr *QGamepad) ButtonL3() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepad_ButtonL3(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQGamepad_ButtonL3Changed
func callbackQGamepad_ButtonL3Changed(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(ptr, "buttonL3Changed"); signal != nil {
		(*(*func(bool))(signal))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonL3Changed(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonL3Changed") {
			C.QGamepad_ConnectButtonL3Changed(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "buttonL3Changed")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonL3Changed"); signal != nil {
			f := func(value bool) {
				(*(*func(bool))(signal))(value)
				f(value)
			}
			qt.ConnectSignal(ptr.Pointer(), "buttonL3Changed", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonL3Changed", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepad) DisconnectButtonL3Changed() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonL3Changed(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "buttonL3Changed")
	}
}

func (ptr *QGamepad) ButtonL3Changed(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonL3Changed(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

func (ptr *QGamepad) ButtonLeft() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepad_ButtonLeft(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQGamepad_ButtonLeftChanged
func callbackQGamepad_ButtonLeftChanged(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(ptr, "buttonLeftChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonLeftChanged(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonLeftChanged") {
			C.QGamepad_ConnectButtonLeftChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "buttonLeftChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonLeftChanged"); signal != nil {
			f := func(value bool) {
				(*(*func(bool))(signal))(value)
				f(value)
			}
			qt.ConnectSignal(ptr.Pointer(), "buttonLeftChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonLeftChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepad) DisconnectButtonLeftChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonLeftChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "buttonLeftChanged")
	}
}

func (ptr *QGamepad) ButtonLeftChanged(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonLeftChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

func (ptr *QGamepad) ButtonR1() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepad_ButtonR1(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQGamepad_ButtonR1Changed
func callbackQGamepad_ButtonR1Changed(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(ptr, "buttonR1Changed"); signal != nil {
		(*(*func(bool))(signal))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonR1Changed(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonR1Changed") {
			C.QGamepad_ConnectButtonR1Changed(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "buttonR1Changed")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonR1Changed"); signal != nil {
			f := func(value bool) {
				(*(*func(bool))(signal))(value)
				f(value)
			}
			qt.ConnectSignal(ptr.Pointer(), "buttonR1Changed", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonR1Changed", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepad) DisconnectButtonR1Changed() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonR1Changed(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "buttonR1Changed")
	}
}

func (ptr *QGamepad) ButtonR1Changed(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonR1Changed(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

func (ptr *QGamepad) ButtonR2() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGamepad_ButtonR2(ptr.Pointer()))
	}
	return 0
}

//export callbackQGamepad_ButtonR2Changed
func callbackQGamepad_ButtonR2Changed(ptr unsafe.Pointer, value C.double) {
	if signal := qt.GetSignal(ptr, "buttonR2Changed"); signal != nil {
		(*(*func(float64))(signal))(float64(value))
	}

}

func (ptr *QGamepad) ConnectButtonR2Changed(f func(value float64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonR2Changed") {
			C.QGamepad_ConnectButtonR2Changed(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "buttonR2Changed")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonR2Changed"); signal != nil {
			f := func(value float64) {
				(*(*func(float64))(signal))(value)
				f(value)
			}
			qt.ConnectSignal(ptr.Pointer(), "buttonR2Changed", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonR2Changed", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepad) DisconnectButtonR2Changed() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonR2Changed(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "buttonR2Changed")
	}
}

func (ptr *QGamepad) ButtonR2Changed(value float64) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonR2Changed(ptr.Pointer(), C.double(value))
	}
}

func (ptr *QGamepad) ButtonR3() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepad_ButtonR3(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQGamepad_ButtonR3Changed
func callbackQGamepad_ButtonR3Changed(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(ptr, "buttonR3Changed"); signal != nil {
		(*(*func(bool))(signal))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonR3Changed(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonR3Changed") {
			C.QGamepad_ConnectButtonR3Changed(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "buttonR3Changed")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonR3Changed"); signal != nil {
			f := func(value bool) {
				(*(*func(bool))(signal))(value)
				f(value)
			}
			qt.ConnectSignal(ptr.Pointer(), "buttonR3Changed", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonR3Changed", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepad) DisconnectButtonR3Changed() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonR3Changed(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "buttonR3Changed")
	}
}

func (ptr *QGamepad) ButtonR3Changed(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonR3Changed(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

func (ptr *QGamepad) ButtonRight() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepad_ButtonRight(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQGamepad_ButtonRightChanged
func callbackQGamepad_ButtonRightChanged(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(ptr, "buttonRightChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonRightChanged(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonRightChanged") {
			C.QGamepad_ConnectButtonRightChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "buttonRightChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonRightChanged"); signal != nil {
			f := func(value bool) {
				(*(*func(bool))(signal))(value)
				f(value)
			}
			qt.ConnectSignal(ptr.Pointer(), "buttonRightChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonRightChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepad) DisconnectButtonRightChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonRightChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "buttonRightChanged")
	}
}

func (ptr *QGamepad) ButtonRightChanged(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonRightChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

func (ptr *QGamepad) ButtonSelect() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepad_ButtonSelect(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQGamepad_ButtonSelectChanged
func callbackQGamepad_ButtonSelectChanged(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(ptr, "buttonSelectChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonSelectChanged(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonSelectChanged") {
			C.QGamepad_ConnectButtonSelectChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "buttonSelectChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonSelectChanged"); signal != nil {
			f := func(value bool) {
				(*(*func(bool))(signal))(value)
				f(value)
			}
			qt.ConnectSignal(ptr.Pointer(), "buttonSelectChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonSelectChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepad) DisconnectButtonSelectChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonSelectChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "buttonSelectChanged")
	}
}

func (ptr *QGamepad) ButtonSelectChanged(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonSelectChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

func (ptr *QGamepad) ButtonStart() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepad_ButtonStart(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQGamepad_ButtonStartChanged
func callbackQGamepad_ButtonStartChanged(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(ptr, "buttonStartChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonStartChanged(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonStartChanged") {
			C.QGamepad_ConnectButtonStartChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "buttonStartChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonStartChanged"); signal != nil {
			f := func(value bool) {
				(*(*func(bool))(signal))(value)
				f(value)
			}
			qt.ConnectSignal(ptr.Pointer(), "buttonStartChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonStartChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepad) DisconnectButtonStartChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonStartChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "buttonStartChanged")
	}
}

func (ptr *QGamepad) ButtonStartChanged(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonStartChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

func (ptr *QGamepad) ButtonUp() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepad_ButtonUp(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQGamepad_ButtonUpChanged
func callbackQGamepad_ButtonUpChanged(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(ptr, "buttonUpChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonUpChanged(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonUpChanged") {
			C.QGamepad_ConnectButtonUpChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "buttonUpChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonUpChanged"); signal != nil {
			f := func(value bool) {
				(*(*func(bool))(signal))(value)
				f(value)
			}
			qt.ConnectSignal(ptr.Pointer(), "buttonUpChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonUpChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepad) DisconnectButtonUpChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonUpChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "buttonUpChanged")
	}
}

func (ptr *QGamepad) ButtonUpChanged(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonUpChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

func (ptr *QGamepad) ButtonX() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepad_ButtonX(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQGamepad_ButtonXChanged
func callbackQGamepad_ButtonXChanged(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(ptr, "buttonXChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonXChanged(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonXChanged") {
			C.QGamepad_ConnectButtonXChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "buttonXChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonXChanged"); signal != nil {
			f := func(value bool) {
				(*(*func(bool))(signal))(value)
				f(value)
			}
			qt.ConnectSignal(ptr.Pointer(), "buttonXChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonXChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepad) DisconnectButtonXChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonXChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "buttonXChanged")
	}
}

func (ptr *QGamepad) ButtonXChanged(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonXChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

func (ptr *QGamepad) ButtonY() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepad_ButtonY(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQGamepad_ButtonYChanged
func callbackQGamepad_ButtonYChanged(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(ptr, "buttonYChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonYChanged(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonYChanged") {
			C.QGamepad_ConnectButtonYChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "buttonYChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonYChanged"); signal != nil {
			f := func(value bool) {
				(*(*func(bool))(signal))(value)
				f(value)
			}
			qt.ConnectSignal(ptr.Pointer(), "buttonYChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonYChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepad) DisconnectButtonYChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonYChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "buttonYChanged")
	}
}

func (ptr *QGamepad) ButtonYChanged(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonYChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_ConnectedChanged
func callbackQGamepad_ConnectedChanged(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(ptr, "connectedChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectConnectedChanged(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "connectedChanged") {
			C.QGamepad_ConnectConnectedChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "connectedChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "connectedChanged"); signal != nil {
			f := func(value bool) {
				(*(*func(bool))(signal))(value)
				f(value)
			}
			qt.ConnectSignal(ptr.Pointer(), "connectedChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "connectedChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepad) DisconnectConnectedChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectConnectedChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "connectedChanged")
	}
}

func (ptr *QGamepad) ConnectedChanged(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ConnectedChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

func (ptr *QGamepad) DeviceId() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QGamepad_DeviceId(ptr.Pointer())))
	}
	return 0
}

//export callbackQGamepad_DeviceIdChanged
func callbackQGamepad_DeviceIdChanged(ptr unsafe.Pointer, value C.int) {
	if signal := qt.GetSignal(ptr, "deviceIdChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(value)))
	}

}

func (ptr *QGamepad) ConnectDeviceIdChanged(f func(value int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "deviceIdChanged") {
			C.QGamepad_ConnectDeviceIdChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "deviceIdChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "deviceIdChanged"); signal != nil {
			f := func(value int) {
				(*(*func(int))(signal))(value)
				f(value)
			}
			qt.ConnectSignal(ptr.Pointer(), "deviceIdChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "deviceIdChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepad) DisconnectDeviceIdChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectDeviceIdChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "deviceIdChanged")
	}
}

func (ptr *QGamepad) DeviceIdChanged(value int) {
	if ptr.Pointer() != nil {
		C.QGamepad_DeviceIdChanged(ptr.Pointer(), C.int(int32(value)))
	}
}

func (ptr *QGamepad) IsConnected() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepad_IsConnected(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGamepad) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGamepad_Name(ptr.Pointer()))
	}
	return ""
}

//export callbackQGamepad_NameChanged
func callbackQGamepad_NameChanged(ptr unsafe.Pointer, value C.struct_QtGamepad_PackedString) {
	if signal := qt.GetSignal(ptr, "nameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(value))
	}

}

func (ptr *QGamepad) ConnectNameChanged(f func(value string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "nameChanged") {
			C.QGamepad_ConnectNameChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "nameChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "nameChanged"); signal != nil {
			f := func(value string) {
				(*(*func(string))(signal))(value)
				f(value)
			}
			qt.ConnectSignal(ptr.Pointer(), "nameChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "nameChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepad) DisconnectNameChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectNameChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "nameChanged")
	}
}

func (ptr *QGamepad) NameChanged(value string) {
	if ptr.Pointer() != nil {
		var valueC *C.char
		if value != "" {
			valueC = C.CString(value)
			defer C.free(unsafe.Pointer(valueC))
		}
		C.QGamepad_NameChanged(ptr.Pointer(), C.struct_QtGamepad_PackedString{data: valueC, len: C.longlong(len(value))})
	}
}

//export callbackQGamepad_SetDeviceId
func callbackQGamepad_SetDeviceId(ptr unsafe.Pointer, number C.int) {
	if signal := qt.GetSignal(ptr, "setDeviceId"); signal != nil {
		(*(*func(int))(signal))(int(int32(number)))
	} else {
		NewQGamepadFromPointer(ptr).SetDeviceIdDefault(int(int32(number)))
	}
}

func (ptr *QGamepad) ConnectSetDeviceId(f func(number int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setDeviceId"); signal != nil {
			f := func(number int) {
				(*(*func(int))(signal))(number)
				f(number)
			}
			qt.ConnectSignal(ptr.Pointer(), "setDeviceId", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setDeviceId", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepad) DisconnectSetDeviceId() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setDeviceId")
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

func (ptr *QGamepad) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGamepad___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
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
	return C.QGamepad___children_newList(ptr.Pointer())
}

func (ptr *QGamepad) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QGamepad___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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
	return C.QGamepad___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QGamepad) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGamepad___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
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
	return C.QGamepad___findChildren_newList(ptr.Pointer())
}

func (ptr *QGamepad) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGamepad___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
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
	return C.QGamepad___findChildren_newList3(ptr.Pointer())
}

//export callbackQGamepad_ChildEvent
func callbackQGamepad_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
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
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
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
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
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
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQGamepadFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QGamepad) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGamepad_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQGamepad_Destroyed
func callbackQGamepad_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQGamepad_DisconnectNotify
func callbackQGamepad_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGamepadFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGamepad) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGamepad_Event
func callbackQGamepad_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGamepadFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QGamepad) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepad_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQGamepad_EventFilter
func callbackQGamepad_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGamepadFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QGamepad) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepad_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQGamepad_MetaObject
func callbackQGamepad_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQGamepadFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QGamepad) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QGamepad_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQGamepad_ObjectNameChanged
func callbackQGamepad_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtGamepad_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQGamepad_TimerEvent
func callbackQGamepad_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGamepadFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGamepad) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGamepad_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QGamepadKeyNavigation struct {
	core.QObject
}

type QGamepadKeyNavigation_ITF interface {
	core.QObject_ITF
	QGamepadKeyNavigation_PTR() *QGamepadKeyNavigation
}

func (ptr *QGamepadKeyNavigation) QGamepadKeyNavigation_PTR() *QGamepadKeyNavigation {
	return ptr
}

func (ptr *QGamepadKeyNavigation) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QGamepadKeyNavigation) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQGamepadKeyNavigation(ptr QGamepadKeyNavigation_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGamepadKeyNavigation_PTR().Pointer()
	}
	return nil
}

func NewQGamepadKeyNavigationFromPointer(ptr unsafe.Pointer) (n *QGamepadKeyNavigation) {
	n = new(QGamepadKeyNavigation)
	n.SetPointer(ptr)
	return
}
func NewQGamepadKeyNavigation(parent core.QObject_ITF) *QGamepadKeyNavigation {
	tmpValue := NewQGamepadKeyNavigationFromPointer(C.QGamepadKeyNavigation_NewQGamepadKeyNavigation(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QGamepadKeyNavigation) Active() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepadKeyNavigation_Active(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQGamepadKeyNavigation_ActiveChanged
func callbackQGamepadKeyNavigation_ActiveChanged(ptr unsafe.Pointer, isActive C.char) {
	if signal := qt.GetSignal(ptr, "activeChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(isActive) != 0)
	}

}

func (ptr *QGamepadKeyNavigation) ConnectActiveChanged(f func(isActive bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "activeChanged") {
			C.QGamepadKeyNavigation_ConnectActiveChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "activeChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "activeChanged"); signal != nil {
			f := func(isActive bool) {
				(*(*func(bool))(signal))(isActive)
				f(isActive)
			}
			qt.ConnectSignal(ptr.Pointer(), "activeChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "activeChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepadKeyNavigation) DisconnectActiveChanged() {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_DisconnectActiveChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "activeChanged")
	}
}

func (ptr *QGamepadKeyNavigation) ActiveChanged(isActive bool) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_ActiveChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(isActive))))
	}
}

func (ptr *QGamepadKeyNavigation) ButtonAKey() core.Qt__Key {
	if ptr.Pointer() != nil {
		return core.Qt__Key(C.QGamepadKeyNavigation_ButtonAKey(ptr.Pointer()))
	}
	return 0
}

//export callbackQGamepadKeyNavigation_ButtonAKeyChanged
func callbackQGamepadKeyNavigation_ButtonAKeyChanged(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "buttonAKeyChanged"); signal != nil {
		(*(*func(core.Qt__Key))(signal))(core.Qt__Key(key))
	}

}

func (ptr *QGamepadKeyNavigation) ConnectButtonAKeyChanged(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonAKeyChanged") {
			C.QGamepadKeyNavigation_ConnectButtonAKeyChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "buttonAKeyChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonAKeyChanged"); signal != nil {
			f := func(key core.Qt__Key) {
				(*(*func(core.Qt__Key))(signal))(key)
				f(key)
			}
			qt.ConnectSignal(ptr.Pointer(), "buttonAKeyChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonAKeyChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepadKeyNavigation) DisconnectButtonAKeyChanged() {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_DisconnectButtonAKeyChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "buttonAKeyChanged")
	}
}

func (ptr *QGamepadKeyNavigation) ButtonAKeyChanged(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_ButtonAKeyChanged(ptr.Pointer(), C.longlong(key))
	}
}

func (ptr *QGamepadKeyNavigation) ButtonBKey() core.Qt__Key {
	if ptr.Pointer() != nil {
		return core.Qt__Key(C.QGamepadKeyNavigation_ButtonBKey(ptr.Pointer()))
	}
	return 0
}

//export callbackQGamepadKeyNavigation_ButtonBKeyChanged
func callbackQGamepadKeyNavigation_ButtonBKeyChanged(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "buttonBKeyChanged"); signal != nil {
		(*(*func(core.Qt__Key))(signal))(core.Qt__Key(key))
	}

}

func (ptr *QGamepadKeyNavigation) ConnectButtonBKeyChanged(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonBKeyChanged") {
			C.QGamepadKeyNavigation_ConnectButtonBKeyChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "buttonBKeyChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonBKeyChanged"); signal != nil {
			f := func(key core.Qt__Key) {
				(*(*func(core.Qt__Key))(signal))(key)
				f(key)
			}
			qt.ConnectSignal(ptr.Pointer(), "buttonBKeyChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonBKeyChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepadKeyNavigation) DisconnectButtonBKeyChanged() {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_DisconnectButtonBKeyChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "buttonBKeyChanged")
	}
}

func (ptr *QGamepadKeyNavigation) ButtonBKeyChanged(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_ButtonBKeyChanged(ptr.Pointer(), C.longlong(key))
	}
}

func (ptr *QGamepadKeyNavigation) ButtonGuideKey() core.Qt__Key {
	if ptr.Pointer() != nil {
		return core.Qt__Key(C.QGamepadKeyNavigation_ButtonGuideKey(ptr.Pointer()))
	}
	return 0
}

//export callbackQGamepadKeyNavigation_ButtonGuideKeyChanged
func callbackQGamepadKeyNavigation_ButtonGuideKeyChanged(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "buttonGuideKeyChanged"); signal != nil {
		(*(*func(core.Qt__Key))(signal))(core.Qt__Key(key))
	}

}

func (ptr *QGamepadKeyNavigation) ConnectButtonGuideKeyChanged(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonGuideKeyChanged") {
			C.QGamepadKeyNavigation_ConnectButtonGuideKeyChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "buttonGuideKeyChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonGuideKeyChanged"); signal != nil {
			f := func(key core.Qt__Key) {
				(*(*func(core.Qt__Key))(signal))(key)
				f(key)
			}
			qt.ConnectSignal(ptr.Pointer(), "buttonGuideKeyChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonGuideKeyChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepadKeyNavigation) DisconnectButtonGuideKeyChanged() {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_DisconnectButtonGuideKeyChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "buttonGuideKeyChanged")
	}
}

func (ptr *QGamepadKeyNavigation) ButtonGuideKeyChanged(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_ButtonGuideKeyChanged(ptr.Pointer(), C.longlong(key))
	}
}

func (ptr *QGamepadKeyNavigation) ButtonL1Key() core.Qt__Key {
	if ptr.Pointer() != nil {
		return core.Qt__Key(C.QGamepadKeyNavigation_ButtonL1Key(ptr.Pointer()))
	}
	return 0
}

//export callbackQGamepadKeyNavigation_ButtonL1KeyChanged
func callbackQGamepadKeyNavigation_ButtonL1KeyChanged(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "buttonL1KeyChanged"); signal != nil {
		(*(*func(core.Qt__Key))(signal))(core.Qt__Key(key))
	}

}

func (ptr *QGamepadKeyNavigation) ConnectButtonL1KeyChanged(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonL1KeyChanged") {
			C.QGamepadKeyNavigation_ConnectButtonL1KeyChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "buttonL1KeyChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonL1KeyChanged"); signal != nil {
			f := func(key core.Qt__Key) {
				(*(*func(core.Qt__Key))(signal))(key)
				f(key)
			}
			qt.ConnectSignal(ptr.Pointer(), "buttonL1KeyChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonL1KeyChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepadKeyNavigation) DisconnectButtonL1KeyChanged() {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_DisconnectButtonL1KeyChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "buttonL1KeyChanged")
	}
}

func (ptr *QGamepadKeyNavigation) ButtonL1KeyChanged(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_ButtonL1KeyChanged(ptr.Pointer(), C.longlong(key))
	}
}

func (ptr *QGamepadKeyNavigation) ButtonL2Key() core.Qt__Key {
	if ptr.Pointer() != nil {
		return core.Qt__Key(C.QGamepadKeyNavigation_ButtonL2Key(ptr.Pointer()))
	}
	return 0
}

//export callbackQGamepadKeyNavigation_ButtonL2KeyChanged
func callbackQGamepadKeyNavigation_ButtonL2KeyChanged(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "buttonL2KeyChanged"); signal != nil {
		(*(*func(core.Qt__Key))(signal))(core.Qt__Key(key))
	}

}

func (ptr *QGamepadKeyNavigation) ConnectButtonL2KeyChanged(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonL2KeyChanged") {
			C.QGamepadKeyNavigation_ConnectButtonL2KeyChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "buttonL2KeyChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonL2KeyChanged"); signal != nil {
			f := func(key core.Qt__Key) {
				(*(*func(core.Qt__Key))(signal))(key)
				f(key)
			}
			qt.ConnectSignal(ptr.Pointer(), "buttonL2KeyChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonL2KeyChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepadKeyNavigation) DisconnectButtonL2KeyChanged() {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_DisconnectButtonL2KeyChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "buttonL2KeyChanged")
	}
}

func (ptr *QGamepadKeyNavigation) ButtonL2KeyChanged(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_ButtonL2KeyChanged(ptr.Pointer(), C.longlong(key))
	}
}

func (ptr *QGamepadKeyNavigation) ButtonL3Key() core.Qt__Key {
	if ptr.Pointer() != nil {
		return core.Qt__Key(C.QGamepadKeyNavigation_ButtonL3Key(ptr.Pointer()))
	}
	return 0
}

//export callbackQGamepadKeyNavigation_ButtonL3KeyChanged
func callbackQGamepadKeyNavigation_ButtonL3KeyChanged(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "buttonL3KeyChanged"); signal != nil {
		(*(*func(core.Qt__Key))(signal))(core.Qt__Key(key))
	}

}

func (ptr *QGamepadKeyNavigation) ConnectButtonL3KeyChanged(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonL3KeyChanged") {
			C.QGamepadKeyNavigation_ConnectButtonL3KeyChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "buttonL3KeyChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonL3KeyChanged"); signal != nil {
			f := func(key core.Qt__Key) {
				(*(*func(core.Qt__Key))(signal))(key)
				f(key)
			}
			qt.ConnectSignal(ptr.Pointer(), "buttonL3KeyChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonL3KeyChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepadKeyNavigation) DisconnectButtonL3KeyChanged() {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_DisconnectButtonL3KeyChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "buttonL3KeyChanged")
	}
}

func (ptr *QGamepadKeyNavigation) ButtonL3KeyChanged(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_ButtonL3KeyChanged(ptr.Pointer(), C.longlong(key))
	}
}

func (ptr *QGamepadKeyNavigation) ButtonR1Key() core.Qt__Key {
	if ptr.Pointer() != nil {
		return core.Qt__Key(C.QGamepadKeyNavigation_ButtonR1Key(ptr.Pointer()))
	}
	return 0
}

//export callbackQGamepadKeyNavigation_ButtonR1KeyChanged
func callbackQGamepadKeyNavigation_ButtonR1KeyChanged(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "buttonR1KeyChanged"); signal != nil {
		(*(*func(core.Qt__Key))(signal))(core.Qt__Key(key))
	}

}

func (ptr *QGamepadKeyNavigation) ConnectButtonR1KeyChanged(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonR1KeyChanged") {
			C.QGamepadKeyNavigation_ConnectButtonR1KeyChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "buttonR1KeyChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonR1KeyChanged"); signal != nil {
			f := func(key core.Qt__Key) {
				(*(*func(core.Qt__Key))(signal))(key)
				f(key)
			}
			qt.ConnectSignal(ptr.Pointer(), "buttonR1KeyChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonR1KeyChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepadKeyNavigation) DisconnectButtonR1KeyChanged() {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_DisconnectButtonR1KeyChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "buttonR1KeyChanged")
	}
}

func (ptr *QGamepadKeyNavigation) ButtonR1KeyChanged(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_ButtonR1KeyChanged(ptr.Pointer(), C.longlong(key))
	}
}

func (ptr *QGamepadKeyNavigation) ButtonR2Key() core.Qt__Key {
	if ptr.Pointer() != nil {
		return core.Qt__Key(C.QGamepadKeyNavigation_ButtonR2Key(ptr.Pointer()))
	}
	return 0
}

//export callbackQGamepadKeyNavigation_ButtonR2KeyChanged
func callbackQGamepadKeyNavigation_ButtonR2KeyChanged(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "buttonR2KeyChanged"); signal != nil {
		(*(*func(core.Qt__Key))(signal))(core.Qt__Key(key))
	}

}

func (ptr *QGamepadKeyNavigation) ConnectButtonR2KeyChanged(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonR2KeyChanged") {
			C.QGamepadKeyNavigation_ConnectButtonR2KeyChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "buttonR2KeyChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonR2KeyChanged"); signal != nil {
			f := func(key core.Qt__Key) {
				(*(*func(core.Qt__Key))(signal))(key)
				f(key)
			}
			qt.ConnectSignal(ptr.Pointer(), "buttonR2KeyChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonR2KeyChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepadKeyNavigation) DisconnectButtonR2KeyChanged() {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_DisconnectButtonR2KeyChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "buttonR2KeyChanged")
	}
}

func (ptr *QGamepadKeyNavigation) ButtonR2KeyChanged(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_ButtonR2KeyChanged(ptr.Pointer(), C.longlong(key))
	}
}

func (ptr *QGamepadKeyNavigation) ButtonR3Key() core.Qt__Key {
	if ptr.Pointer() != nil {
		return core.Qt__Key(C.QGamepadKeyNavigation_ButtonR3Key(ptr.Pointer()))
	}
	return 0
}

//export callbackQGamepadKeyNavigation_ButtonR3KeyChanged
func callbackQGamepadKeyNavigation_ButtonR3KeyChanged(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "buttonR3KeyChanged"); signal != nil {
		(*(*func(core.Qt__Key))(signal))(core.Qt__Key(key))
	}

}

func (ptr *QGamepadKeyNavigation) ConnectButtonR3KeyChanged(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonR3KeyChanged") {
			C.QGamepadKeyNavigation_ConnectButtonR3KeyChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "buttonR3KeyChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonR3KeyChanged"); signal != nil {
			f := func(key core.Qt__Key) {
				(*(*func(core.Qt__Key))(signal))(key)
				f(key)
			}
			qt.ConnectSignal(ptr.Pointer(), "buttonR3KeyChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonR3KeyChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepadKeyNavigation) DisconnectButtonR3KeyChanged() {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_DisconnectButtonR3KeyChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "buttonR3KeyChanged")
	}
}

func (ptr *QGamepadKeyNavigation) ButtonR3KeyChanged(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_ButtonR3KeyChanged(ptr.Pointer(), C.longlong(key))
	}
}

func (ptr *QGamepadKeyNavigation) ButtonSelectKey() core.Qt__Key {
	if ptr.Pointer() != nil {
		return core.Qt__Key(C.QGamepadKeyNavigation_ButtonSelectKey(ptr.Pointer()))
	}
	return 0
}

//export callbackQGamepadKeyNavigation_ButtonSelectKeyChanged
func callbackQGamepadKeyNavigation_ButtonSelectKeyChanged(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "buttonSelectKeyChanged"); signal != nil {
		(*(*func(core.Qt__Key))(signal))(core.Qt__Key(key))
	}

}

func (ptr *QGamepadKeyNavigation) ConnectButtonSelectKeyChanged(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonSelectKeyChanged") {
			C.QGamepadKeyNavigation_ConnectButtonSelectKeyChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "buttonSelectKeyChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonSelectKeyChanged"); signal != nil {
			f := func(key core.Qt__Key) {
				(*(*func(core.Qt__Key))(signal))(key)
				f(key)
			}
			qt.ConnectSignal(ptr.Pointer(), "buttonSelectKeyChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonSelectKeyChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepadKeyNavigation) DisconnectButtonSelectKeyChanged() {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_DisconnectButtonSelectKeyChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "buttonSelectKeyChanged")
	}
}

func (ptr *QGamepadKeyNavigation) ButtonSelectKeyChanged(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_ButtonSelectKeyChanged(ptr.Pointer(), C.longlong(key))
	}
}

func (ptr *QGamepadKeyNavigation) ButtonStartKey() core.Qt__Key {
	if ptr.Pointer() != nil {
		return core.Qt__Key(C.QGamepadKeyNavigation_ButtonStartKey(ptr.Pointer()))
	}
	return 0
}

//export callbackQGamepadKeyNavigation_ButtonStartKeyChanged
func callbackQGamepadKeyNavigation_ButtonStartKeyChanged(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "buttonStartKeyChanged"); signal != nil {
		(*(*func(core.Qt__Key))(signal))(core.Qt__Key(key))
	}

}

func (ptr *QGamepadKeyNavigation) ConnectButtonStartKeyChanged(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonStartKeyChanged") {
			C.QGamepadKeyNavigation_ConnectButtonStartKeyChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "buttonStartKeyChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonStartKeyChanged"); signal != nil {
			f := func(key core.Qt__Key) {
				(*(*func(core.Qt__Key))(signal))(key)
				f(key)
			}
			qt.ConnectSignal(ptr.Pointer(), "buttonStartKeyChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonStartKeyChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepadKeyNavigation) DisconnectButtonStartKeyChanged() {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_DisconnectButtonStartKeyChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "buttonStartKeyChanged")
	}
}

func (ptr *QGamepadKeyNavigation) ButtonStartKeyChanged(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_ButtonStartKeyChanged(ptr.Pointer(), C.longlong(key))
	}
}

func (ptr *QGamepadKeyNavigation) ButtonXKey() core.Qt__Key {
	if ptr.Pointer() != nil {
		return core.Qt__Key(C.QGamepadKeyNavigation_ButtonXKey(ptr.Pointer()))
	}
	return 0
}

//export callbackQGamepadKeyNavigation_ButtonXKeyChanged
func callbackQGamepadKeyNavigation_ButtonXKeyChanged(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "buttonXKeyChanged"); signal != nil {
		(*(*func(core.Qt__Key))(signal))(core.Qt__Key(key))
	}

}

func (ptr *QGamepadKeyNavigation) ConnectButtonXKeyChanged(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonXKeyChanged") {
			C.QGamepadKeyNavigation_ConnectButtonXKeyChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "buttonXKeyChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonXKeyChanged"); signal != nil {
			f := func(key core.Qt__Key) {
				(*(*func(core.Qt__Key))(signal))(key)
				f(key)
			}
			qt.ConnectSignal(ptr.Pointer(), "buttonXKeyChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonXKeyChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepadKeyNavigation) DisconnectButtonXKeyChanged() {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_DisconnectButtonXKeyChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "buttonXKeyChanged")
	}
}

func (ptr *QGamepadKeyNavigation) ButtonXKeyChanged(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_ButtonXKeyChanged(ptr.Pointer(), C.longlong(key))
	}
}

func (ptr *QGamepadKeyNavigation) ButtonYKey() core.Qt__Key {
	if ptr.Pointer() != nil {
		return core.Qt__Key(C.QGamepadKeyNavigation_ButtonYKey(ptr.Pointer()))
	}
	return 0
}

//export callbackQGamepadKeyNavigation_ButtonYKeyChanged
func callbackQGamepadKeyNavigation_ButtonYKeyChanged(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "buttonYKeyChanged"); signal != nil {
		(*(*func(core.Qt__Key))(signal))(core.Qt__Key(key))
	}

}

func (ptr *QGamepadKeyNavigation) ConnectButtonYKeyChanged(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonYKeyChanged") {
			C.QGamepadKeyNavigation_ConnectButtonYKeyChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "buttonYKeyChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonYKeyChanged"); signal != nil {
			f := func(key core.Qt__Key) {
				(*(*func(core.Qt__Key))(signal))(key)
				f(key)
			}
			qt.ConnectSignal(ptr.Pointer(), "buttonYKeyChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonYKeyChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepadKeyNavigation) DisconnectButtonYKeyChanged() {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_DisconnectButtonYKeyChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "buttonYKeyChanged")
	}
}

func (ptr *QGamepadKeyNavigation) ButtonYKeyChanged(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_ButtonYKeyChanged(ptr.Pointer(), C.longlong(key))
	}
}

func (ptr *QGamepadKeyNavigation) DownKey() core.Qt__Key {
	if ptr.Pointer() != nil {
		return core.Qt__Key(C.QGamepadKeyNavigation_DownKey(ptr.Pointer()))
	}
	return 0
}

//export callbackQGamepadKeyNavigation_DownKeyChanged
func callbackQGamepadKeyNavigation_DownKeyChanged(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "downKeyChanged"); signal != nil {
		(*(*func(core.Qt__Key))(signal))(core.Qt__Key(key))
	}

}

func (ptr *QGamepadKeyNavigation) ConnectDownKeyChanged(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "downKeyChanged") {
			C.QGamepadKeyNavigation_ConnectDownKeyChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "downKeyChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "downKeyChanged"); signal != nil {
			f := func(key core.Qt__Key) {
				(*(*func(core.Qt__Key))(signal))(key)
				f(key)
			}
			qt.ConnectSignal(ptr.Pointer(), "downKeyChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "downKeyChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepadKeyNavigation) DisconnectDownKeyChanged() {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_DisconnectDownKeyChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "downKeyChanged")
	}
}

func (ptr *QGamepadKeyNavigation) DownKeyChanged(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_DownKeyChanged(ptr.Pointer(), C.longlong(key))
	}
}

func (ptr *QGamepadKeyNavigation) Gamepad() *QGamepad {
	if ptr.Pointer() != nil {
		tmpValue := NewQGamepadFromPointer(C.QGamepadKeyNavigation_Gamepad(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQGamepadKeyNavigation_GamepadChanged
func callbackQGamepadKeyNavigation_GamepadChanged(ptr unsafe.Pointer, gamepad unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "gamepadChanged"); signal != nil {
		(*(*func(*QGamepad))(signal))(NewQGamepadFromPointer(gamepad))
	}

}

func (ptr *QGamepadKeyNavigation) ConnectGamepadChanged(f func(gamepad *QGamepad)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "gamepadChanged") {
			C.QGamepadKeyNavigation_ConnectGamepadChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "gamepadChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "gamepadChanged"); signal != nil {
			f := func(gamepad *QGamepad) {
				(*(*func(*QGamepad))(signal))(gamepad)
				f(gamepad)
			}
			qt.ConnectSignal(ptr.Pointer(), "gamepadChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "gamepadChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepadKeyNavigation) DisconnectGamepadChanged() {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_DisconnectGamepadChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "gamepadChanged")
	}
}

func (ptr *QGamepadKeyNavigation) GamepadChanged(gamepad QGamepad_ITF) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_GamepadChanged(ptr.Pointer(), PointerFromQGamepad(gamepad))
	}
}

func (ptr *QGamepadKeyNavigation) LeftKey() core.Qt__Key {
	if ptr.Pointer() != nil {
		return core.Qt__Key(C.QGamepadKeyNavigation_LeftKey(ptr.Pointer()))
	}
	return 0
}

//export callbackQGamepadKeyNavigation_LeftKeyChanged
func callbackQGamepadKeyNavigation_LeftKeyChanged(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "leftKeyChanged"); signal != nil {
		(*(*func(core.Qt__Key))(signal))(core.Qt__Key(key))
	}

}

func (ptr *QGamepadKeyNavigation) ConnectLeftKeyChanged(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "leftKeyChanged") {
			C.QGamepadKeyNavigation_ConnectLeftKeyChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "leftKeyChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "leftKeyChanged"); signal != nil {
			f := func(key core.Qt__Key) {
				(*(*func(core.Qt__Key))(signal))(key)
				f(key)
			}
			qt.ConnectSignal(ptr.Pointer(), "leftKeyChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "leftKeyChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepadKeyNavigation) DisconnectLeftKeyChanged() {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_DisconnectLeftKeyChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "leftKeyChanged")
	}
}

func (ptr *QGamepadKeyNavigation) LeftKeyChanged(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_LeftKeyChanged(ptr.Pointer(), C.longlong(key))
	}
}

func (ptr *QGamepadKeyNavigation) RightKey() core.Qt__Key {
	if ptr.Pointer() != nil {
		return core.Qt__Key(C.QGamepadKeyNavigation_RightKey(ptr.Pointer()))
	}
	return 0
}

//export callbackQGamepadKeyNavigation_RightKeyChanged
func callbackQGamepadKeyNavigation_RightKeyChanged(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "rightKeyChanged"); signal != nil {
		(*(*func(core.Qt__Key))(signal))(core.Qt__Key(key))
	}

}

func (ptr *QGamepadKeyNavigation) ConnectRightKeyChanged(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "rightKeyChanged") {
			C.QGamepadKeyNavigation_ConnectRightKeyChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "rightKeyChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "rightKeyChanged"); signal != nil {
			f := func(key core.Qt__Key) {
				(*(*func(core.Qt__Key))(signal))(key)
				f(key)
			}
			qt.ConnectSignal(ptr.Pointer(), "rightKeyChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "rightKeyChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepadKeyNavigation) DisconnectRightKeyChanged() {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_DisconnectRightKeyChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "rightKeyChanged")
	}
}

func (ptr *QGamepadKeyNavigation) RightKeyChanged(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_RightKeyChanged(ptr.Pointer(), C.longlong(key))
	}
}

//export callbackQGamepadKeyNavigation_SetActive
func callbackQGamepadKeyNavigation_SetActive(ptr unsafe.Pointer, isActive C.char) {
	if signal := qt.GetSignal(ptr, "setActive"); signal != nil {
		(*(*func(bool))(signal))(int8(isActive) != 0)
	} else {
		NewQGamepadKeyNavigationFromPointer(ptr).SetActiveDefault(int8(isActive) != 0)
	}
}

func (ptr *QGamepadKeyNavigation) ConnectSetActive(f func(isActive bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setActive"); signal != nil {
			f := func(isActive bool) {
				(*(*func(bool))(signal))(isActive)
				f(isActive)
			}
			qt.ConnectSignal(ptr.Pointer(), "setActive", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setActive", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepadKeyNavigation) DisconnectSetActive() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setActive")
	}
}

func (ptr *QGamepadKeyNavigation) SetActive(isActive bool) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_SetActive(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(isActive))))
	}
}

func (ptr *QGamepadKeyNavigation) SetActiveDefault(isActive bool) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_SetActiveDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(isActive))))
	}
}

//export callbackQGamepadKeyNavigation_SetButtonAKey
func callbackQGamepadKeyNavigation_SetButtonAKey(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "setButtonAKey"); signal != nil {
		(*(*func(core.Qt__Key))(signal))(core.Qt__Key(key))
	} else {
		NewQGamepadKeyNavigationFromPointer(ptr).SetButtonAKeyDefault(core.Qt__Key(key))
	}
}

func (ptr *QGamepadKeyNavigation) ConnectSetButtonAKey(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setButtonAKey"); signal != nil {
			f := func(key core.Qt__Key) {
				(*(*func(core.Qt__Key))(signal))(key)
				f(key)
			}
			qt.ConnectSignal(ptr.Pointer(), "setButtonAKey", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setButtonAKey", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepadKeyNavigation) DisconnectSetButtonAKey() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setButtonAKey")
	}
}

func (ptr *QGamepadKeyNavigation) SetButtonAKey(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_SetButtonAKey(ptr.Pointer(), C.longlong(key))
	}
}

func (ptr *QGamepadKeyNavigation) SetButtonAKeyDefault(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_SetButtonAKeyDefault(ptr.Pointer(), C.longlong(key))
	}
}

//export callbackQGamepadKeyNavigation_SetButtonBKey
func callbackQGamepadKeyNavigation_SetButtonBKey(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "setButtonBKey"); signal != nil {
		(*(*func(core.Qt__Key))(signal))(core.Qt__Key(key))
	} else {
		NewQGamepadKeyNavigationFromPointer(ptr).SetButtonBKeyDefault(core.Qt__Key(key))
	}
}

func (ptr *QGamepadKeyNavigation) ConnectSetButtonBKey(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setButtonBKey"); signal != nil {
			f := func(key core.Qt__Key) {
				(*(*func(core.Qt__Key))(signal))(key)
				f(key)
			}
			qt.ConnectSignal(ptr.Pointer(), "setButtonBKey", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setButtonBKey", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepadKeyNavigation) DisconnectSetButtonBKey() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setButtonBKey")
	}
}

func (ptr *QGamepadKeyNavigation) SetButtonBKey(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_SetButtonBKey(ptr.Pointer(), C.longlong(key))
	}
}

func (ptr *QGamepadKeyNavigation) SetButtonBKeyDefault(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_SetButtonBKeyDefault(ptr.Pointer(), C.longlong(key))
	}
}

//export callbackQGamepadKeyNavigation_SetButtonGuideKey
func callbackQGamepadKeyNavigation_SetButtonGuideKey(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "setButtonGuideKey"); signal != nil {
		(*(*func(core.Qt__Key))(signal))(core.Qt__Key(key))
	} else {
		NewQGamepadKeyNavigationFromPointer(ptr).SetButtonGuideKeyDefault(core.Qt__Key(key))
	}
}

func (ptr *QGamepadKeyNavigation) ConnectSetButtonGuideKey(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setButtonGuideKey"); signal != nil {
			f := func(key core.Qt__Key) {
				(*(*func(core.Qt__Key))(signal))(key)
				f(key)
			}
			qt.ConnectSignal(ptr.Pointer(), "setButtonGuideKey", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setButtonGuideKey", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepadKeyNavigation) DisconnectSetButtonGuideKey() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setButtonGuideKey")
	}
}

func (ptr *QGamepadKeyNavigation) SetButtonGuideKey(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_SetButtonGuideKey(ptr.Pointer(), C.longlong(key))
	}
}

func (ptr *QGamepadKeyNavigation) SetButtonGuideKeyDefault(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_SetButtonGuideKeyDefault(ptr.Pointer(), C.longlong(key))
	}
}

//export callbackQGamepadKeyNavigation_SetButtonL1Key
func callbackQGamepadKeyNavigation_SetButtonL1Key(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "setButtonL1Key"); signal != nil {
		(*(*func(core.Qt__Key))(signal))(core.Qt__Key(key))
	} else {
		NewQGamepadKeyNavigationFromPointer(ptr).SetButtonL1KeyDefault(core.Qt__Key(key))
	}
}

func (ptr *QGamepadKeyNavigation) ConnectSetButtonL1Key(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setButtonL1Key"); signal != nil {
			f := func(key core.Qt__Key) {
				(*(*func(core.Qt__Key))(signal))(key)
				f(key)
			}
			qt.ConnectSignal(ptr.Pointer(), "setButtonL1Key", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setButtonL1Key", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepadKeyNavigation) DisconnectSetButtonL1Key() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setButtonL1Key")
	}
}

func (ptr *QGamepadKeyNavigation) SetButtonL1Key(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_SetButtonL1Key(ptr.Pointer(), C.longlong(key))
	}
}

func (ptr *QGamepadKeyNavigation) SetButtonL1KeyDefault(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_SetButtonL1KeyDefault(ptr.Pointer(), C.longlong(key))
	}
}

//export callbackQGamepadKeyNavigation_SetButtonL2Key
func callbackQGamepadKeyNavigation_SetButtonL2Key(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "setButtonL2Key"); signal != nil {
		(*(*func(core.Qt__Key))(signal))(core.Qt__Key(key))
	} else {
		NewQGamepadKeyNavigationFromPointer(ptr).SetButtonL2KeyDefault(core.Qt__Key(key))
	}
}

func (ptr *QGamepadKeyNavigation) ConnectSetButtonL2Key(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setButtonL2Key"); signal != nil {
			f := func(key core.Qt__Key) {
				(*(*func(core.Qt__Key))(signal))(key)
				f(key)
			}
			qt.ConnectSignal(ptr.Pointer(), "setButtonL2Key", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setButtonL2Key", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepadKeyNavigation) DisconnectSetButtonL2Key() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setButtonL2Key")
	}
}

func (ptr *QGamepadKeyNavigation) SetButtonL2Key(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_SetButtonL2Key(ptr.Pointer(), C.longlong(key))
	}
}

func (ptr *QGamepadKeyNavigation) SetButtonL2KeyDefault(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_SetButtonL2KeyDefault(ptr.Pointer(), C.longlong(key))
	}
}

//export callbackQGamepadKeyNavigation_SetButtonL3Key
func callbackQGamepadKeyNavigation_SetButtonL3Key(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "setButtonL3Key"); signal != nil {
		(*(*func(core.Qt__Key))(signal))(core.Qt__Key(key))
	} else {
		NewQGamepadKeyNavigationFromPointer(ptr).SetButtonL3KeyDefault(core.Qt__Key(key))
	}
}

func (ptr *QGamepadKeyNavigation) ConnectSetButtonL3Key(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setButtonL3Key"); signal != nil {
			f := func(key core.Qt__Key) {
				(*(*func(core.Qt__Key))(signal))(key)
				f(key)
			}
			qt.ConnectSignal(ptr.Pointer(), "setButtonL3Key", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setButtonL3Key", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepadKeyNavigation) DisconnectSetButtonL3Key() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setButtonL3Key")
	}
}

func (ptr *QGamepadKeyNavigation) SetButtonL3Key(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_SetButtonL3Key(ptr.Pointer(), C.longlong(key))
	}
}

func (ptr *QGamepadKeyNavigation) SetButtonL3KeyDefault(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_SetButtonL3KeyDefault(ptr.Pointer(), C.longlong(key))
	}
}

//export callbackQGamepadKeyNavigation_SetButtonR1Key
func callbackQGamepadKeyNavigation_SetButtonR1Key(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "setButtonR1Key"); signal != nil {
		(*(*func(core.Qt__Key))(signal))(core.Qt__Key(key))
	} else {
		NewQGamepadKeyNavigationFromPointer(ptr).SetButtonR1KeyDefault(core.Qt__Key(key))
	}
}

func (ptr *QGamepadKeyNavigation) ConnectSetButtonR1Key(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setButtonR1Key"); signal != nil {
			f := func(key core.Qt__Key) {
				(*(*func(core.Qt__Key))(signal))(key)
				f(key)
			}
			qt.ConnectSignal(ptr.Pointer(), "setButtonR1Key", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setButtonR1Key", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepadKeyNavigation) DisconnectSetButtonR1Key() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setButtonR1Key")
	}
}

func (ptr *QGamepadKeyNavigation) SetButtonR1Key(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_SetButtonR1Key(ptr.Pointer(), C.longlong(key))
	}
}

func (ptr *QGamepadKeyNavigation) SetButtonR1KeyDefault(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_SetButtonR1KeyDefault(ptr.Pointer(), C.longlong(key))
	}
}

//export callbackQGamepadKeyNavigation_SetButtonR2Key
func callbackQGamepadKeyNavigation_SetButtonR2Key(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "setButtonR2Key"); signal != nil {
		(*(*func(core.Qt__Key))(signal))(core.Qt__Key(key))
	} else {
		NewQGamepadKeyNavigationFromPointer(ptr).SetButtonR2KeyDefault(core.Qt__Key(key))
	}
}

func (ptr *QGamepadKeyNavigation) ConnectSetButtonR2Key(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setButtonR2Key"); signal != nil {
			f := func(key core.Qt__Key) {
				(*(*func(core.Qt__Key))(signal))(key)
				f(key)
			}
			qt.ConnectSignal(ptr.Pointer(), "setButtonR2Key", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setButtonR2Key", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepadKeyNavigation) DisconnectSetButtonR2Key() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setButtonR2Key")
	}
}

func (ptr *QGamepadKeyNavigation) SetButtonR2Key(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_SetButtonR2Key(ptr.Pointer(), C.longlong(key))
	}
}

func (ptr *QGamepadKeyNavigation) SetButtonR2KeyDefault(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_SetButtonR2KeyDefault(ptr.Pointer(), C.longlong(key))
	}
}

//export callbackQGamepadKeyNavigation_SetButtonR3Key
func callbackQGamepadKeyNavigation_SetButtonR3Key(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "setButtonR3Key"); signal != nil {
		(*(*func(core.Qt__Key))(signal))(core.Qt__Key(key))
	} else {
		NewQGamepadKeyNavigationFromPointer(ptr).SetButtonR3KeyDefault(core.Qt__Key(key))
	}
}

func (ptr *QGamepadKeyNavigation) ConnectSetButtonR3Key(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setButtonR3Key"); signal != nil {
			f := func(key core.Qt__Key) {
				(*(*func(core.Qt__Key))(signal))(key)
				f(key)
			}
			qt.ConnectSignal(ptr.Pointer(), "setButtonR3Key", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setButtonR3Key", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepadKeyNavigation) DisconnectSetButtonR3Key() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setButtonR3Key")
	}
}

func (ptr *QGamepadKeyNavigation) SetButtonR3Key(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_SetButtonR3Key(ptr.Pointer(), C.longlong(key))
	}
}

func (ptr *QGamepadKeyNavigation) SetButtonR3KeyDefault(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_SetButtonR3KeyDefault(ptr.Pointer(), C.longlong(key))
	}
}

//export callbackQGamepadKeyNavigation_SetButtonSelectKey
func callbackQGamepadKeyNavigation_SetButtonSelectKey(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "setButtonSelectKey"); signal != nil {
		(*(*func(core.Qt__Key))(signal))(core.Qt__Key(key))
	} else {
		NewQGamepadKeyNavigationFromPointer(ptr).SetButtonSelectKeyDefault(core.Qt__Key(key))
	}
}

func (ptr *QGamepadKeyNavigation) ConnectSetButtonSelectKey(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setButtonSelectKey"); signal != nil {
			f := func(key core.Qt__Key) {
				(*(*func(core.Qt__Key))(signal))(key)
				f(key)
			}
			qt.ConnectSignal(ptr.Pointer(), "setButtonSelectKey", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setButtonSelectKey", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepadKeyNavigation) DisconnectSetButtonSelectKey() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setButtonSelectKey")
	}
}

func (ptr *QGamepadKeyNavigation) SetButtonSelectKey(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_SetButtonSelectKey(ptr.Pointer(), C.longlong(key))
	}
}

func (ptr *QGamepadKeyNavigation) SetButtonSelectKeyDefault(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_SetButtonSelectKeyDefault(ptr.Pointer(), C.longlong(key))
	}
}

//export callbackQGamepadKeyNavigation_SetButtonStartKey
func callbackQGamepadKeyNavigation_SetButtonStartKey(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "setButtonStartKey"); signal != nil {
		(*(*func(core.Qt__Key))(signal))(core.Qt__Key(key))
	} else {
		NewQGamepadKeyNavigationFromPointer(ptr).SetButtonStartKeyDefault(core.Qt__Key(key))
	}
}

func (ptr *QGamepadKeyNavigation) ConnectSetButtonStartKey(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setButtonStartKey"); signal != nil {
			f := func(key core.Qt__Key) {
				(*(*func(core.Qt__Key))(signal))(key)
				f(key)
			}
			qt.ConnectSignal(ptr.Pointer(), "setButtonStartKey", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setButtonStartKey", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepadKeyNavigation) DisconnectSetButtonStartKey() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setButtonStartKey")
	}
}

func (ptr *QGamepadKeyNavigation) SetButtonStartKey(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_SetButtonStartKey(ptr.Pointer(), C.longlong(key))
	}
}

func (ptr *QGamepadKeyNavigation) SetButtonStartKeyDefault(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_SetButtonStartKeyDefault(ptr.Pointer(), C.longlong(key))
	}
}

//export callbackQGamepadKeyNavigation_SetButtonXKey
func callbackQGamepadKeyNavigation_SetButtonXKey(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "setButtonXKey"); signal != nil {
		(*(*func(core.Qt__Key))(signal))(core.Qt__Key(key))
	} else {
		NewQGamepadKeyNavigationFromPointer(ptr).SetButtonXKeyDefault(core.Qt__Key(key))
	}
}

func (ptr *QGamepadKeyNavigation) ConnectSetButtonXKey(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setButtonXKey"); signal != nil {
			f := func(key core.Qt__Key) {
				(*(*func(core.Qt__Key))(signal))(key)
				f(key)
			}
			qt.ConnectSignal(ptr.Pointer(), "setButtonXKey", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setButtonXKey", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepadKeyNavigation) DisconnectSetButtonXKey() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setButtonXKey")
	}
}

func (ptr *QGamepadKeyNavigation) SetButtonXKey(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_SetButtonXKey(ptr.Pointer(), C.longlong(key))
	}
}

func (ptr *QGamepadKeyNavigation) SetButtonXKeyDefault(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_SetButtonXKeyDefault(ptr.Pointer(), C.longlong(key))
	}
}

//export callbackQGamepadKeyNavigation_SetButtonYKey
func callbackQGamepadKeyNavigation_SetButtonYKey(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "setButtonYKey"); signal != nil {
		(*(*func(core.Qt__Key))(signal))(core.Qt__Key(key))
	} else {
		NewQGamepadKeyNavigationFromPointer(ptr).SetButtonYKeyDefault(core.Qt__Key(key))
	}
}

func (ptr *QGamepadKeyNavigation) ConnectSetButtonYKey(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setButtonYKey"); signal != nil {
			f := func(key core.Qt__Key) {
				(*(*func(core.Qt__Key))(signal))(key)
				f(key)
			}
			qt.ConnectSignal(ptr.Pointer(), "setButtonYKey", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setButtonYKey", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepadKeyNavigation) DisconnectSetButtonYKey() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setButtonYKey")
	}
}

func (ptr *QGamepadKeyNavigation) SetButtonYKey(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_SetButtonYKey(ptr.Pointer(), C.longlong(key))
	}
}

func (ptr *QGamepadKeyNavigation) SetButtonYKeyDefault(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_SetButtonYKeyDefault(ptr.Pointer(), C.longlong(key))
	}
}

//export callbackQGamepadKeyNavigation_SetDownKey
func callbackQGamepadKeyNavigation_SetDownKey(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "setDownKey"); signal != nil {
		(*(*func(core.Qt__Key))(signal))(core.Qt__Key(key))
	} else {
		NewQGamepadKeyNavigationFromPointer(ptr).SetDownKeyDefault(core.Qt__Key(key))
	}
}

func (ptr *QGamepadKeyNavigation) ConnectSetDownKey(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setDownKey"); signal != nil {
			f := func(key core.Qt__Key) {
				(*(*func(core.Qt__Key))(signal))(key)
				f(key)
			}
			qt.ConnectSignal(ptr.Pointer(), "setDownKey", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setDownKey", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepadKeyNavigation) DisconnectSetDownKey() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setDownKey")
	}
}

func (ptr *QGamepadKeyNavigation) SetDownKey(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_SetDownKey(ptr.Pointer(), C.longlong(key))
	}
}

func (ptr *QGamepadKeyNavigation) SetDownKeyDefault(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_SetDownKeyDefault(ptr.Pointer(), C.longlong(key))
	}
}

//export callbackQGamepadKeyNavigation_SetGamepad
func callbackQGamepadKeyNavigation_SetGamepad(ptr unsafe.Pointer, gamepad unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setGamepad"); signal != nil {
		(*(*func(*QGamepad))(signal))(NewQGamepadFromPointer(gamepad))
	} else {
		NewQGamepadKeyNavigationFromPointer(ptr).SetGamepadDefault(NewQGamepadFromPointer(gamepad))
	}
}

func (ptr *QGamepadKeyNavigation) ConnectSetGamepad(f func(gamepad *QGamepad)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setGamepad"); signal != nil {
			f := func(gamepad *QGamepad) {
				(*(*func(*QGamepad))(signal))(gamepad)
				f(gamepad)
			}
			qt.ConnectSignal(ptr.Pointer(), "setGamepad", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setGamepad", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepadKeyNavigation) DisconnectSetGamepad() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setGamepad")
	}
}

func (ptr *QGamepadKeyNavigation) SetGamepad(gamepad QGamepad_ITF) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_SetGamepad(ptr.Pointer(), PointerFromQGamepad(gamepad))
	}
}

func (ptr *QGamepadKeyNavigation) SetGamepadDefault(gamepad QGamepad_ITF) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_SetGamepadDefault(ptr.Pointer(), PointerFromQGamepad(gamepad))
	}
}

//export callbackQGamepadKeyNavigation_SetLeftKey
func callbackQGamepadKeyNavigation_SetLeftKey(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "setLeftKey"); signal != nil {
		(*(*func(core.Qt__Key))(signal))(core.Qt__Key(key))
	} else {
		NewQGamepadKeyNavigationFromPointer(ptr).SetLeftKeyDefault(core.Qt__Key(key))
	}
}

func (ptr *QGamepadKeyNavigation) ConnectSetLeftKey(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setLeftKey"); signal != nil {
			f := func(key core.Qt__Key) {
				(*(*func(core.Qt__Key))(signal))(key)
				f(key)
			}
			qt.ConnectSignal(ptr.Pointer(), "setLeftKey", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setLeftKey", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepadKeyNavigation) DisconnectSetLeftKey() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setLeftKey")
	}
}

func (ptr *QGamepadKeyNavigation) SetLeftKey(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_SetLeftKey(ptr.Pointer(), C.longlong(key))
	}
}

func (ptr *QGamepadKeyNavigation) SetLeftKeyDefault(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_SetLeftKeyDefault(ptr.Pointer(), C.longlong(key))
	}
}

//export callbackQGamepadKeyNavigation_SetRightKey
func callbackQGamepadKeyNavigation_SetRightKey(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "setRightKey"); signal != nil {
		(*(*func(core.Qt__Key))(signal))(core.Qt__Key(key))
	} else {
		NewQGamepadKeyNavigationFromPointer(ptr).SetRightKeyDefault(core.Qt__Key(key))
	}
}

func (ptr *QGamepadKeyNavigation) ConnectSetRightKey(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setRightKey"); signal != nil {
			f := func(key core.Qt__Key) {
				(*(*func(core.Qt__Key))(signal))(key)
				f(key)
			}
			qt.ConnectSignal(ptr.Pointer(), "setRightKey", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setRightKey", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepadKeyNavigation) DisconnectSetRightKey() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setRightKey")
	}
}

func (ptr *QGamepadKeyNavigation) SetRightKey(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_SetRightKey(ptr.Pointer(), C.longlong(key))
	}
}

func (ptr *QGamepadKeyNavigation) SetRightKeyDefault(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_SetRightKeyDefault(ptr.Pointer(), C.longlong(key))
	}
}

//export callbackQGamepadKeyNavigation_SetUpKey
func callbackQGamepadKeyNavigation_SetUpKey(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "setUpKey"); signal != nil {
		(*(*func(core.Qt__Key))(signal))(core.Qt__Key(key))
	} else {
		NewQGamepadKeyNavigationFromPointer(ptr).SetUpKeyDefault(core.Qt__Key(key))
	}
}

func (ptr *QGamepadKeyNavigation) ConnectSetUpKey(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setUpKey"); signal != nil {
			f := func(key core.Qt__Key) {
				(*(*func(core.Qt__Key))(signal))(key)
				f(key)
			}
			qt.ConnectSignal(ptr.Pointer(), "setUpKey", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setUpKey", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepadKeyNavigation) DisconnectSetUpKey() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setUpKey")
	}
}

func (ptr *QGamepadKeyNavigation) SetUpKey(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_SetUpKey(ptr.Pointer(), C.longlong(key))
	}
}

func (ptr *QGamepadKeyNavigation) SetUpKeyDefault(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_SetUpKeyDefault(ptr.Pointer(), C.longlong(key))
	}
}

func (ptr *QGamepadKeyNavigation) UpKey() core.Qt__Key {
	if ptr.Pointer() != nil {
		return core.Qt__Key(C.QGamepadKeyNavigation_UpKey(ptr.Pointer()))
	}
	return 0
}

//export callbackQGamepadKeyNavigation_UpKeyChanged
func callbackQGamepadKeyNavigation_UpKeyChanged(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "upKeyChanged"); signal != nil {
		(*(*func(core.Qt__Key))(signal))(core.Qt__Key(key))
	}

}

func (ptr *QGamepadKeyNavigation) ConnectUpKeyChanged(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "upKeyChanged") {
			C.QGamepadKeyNavigation_ConnectUpKeyChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "upKeyChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "upKeyChanged"); signal != nil {
			f := func(key core.Qt__Key) {
				(*(*func(core.Qt__Key))(signal))(key)
				f(key)
			}
			qt.ConnectSignal(ptr.Pointer(), "upKeyChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "upKeyChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepadKeyNavigation) DisconnectUpKeyChanged() {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_DisconnectUpKeyChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "upKeyChanged")
	}
}

func (ptr *QGamepadKeyNavigation) UpKeyChanged(key core.Qt__Key) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_UpKeyChanged(ptr.Pointer(), C.longlong(key))
	}
}

func (ptr *QGamepadKeyNavigation) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGamepadKeyNavigation___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGamepadKeyNavigation) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGamepadKeyNavigation) __children_newList() unsafe.Pointer {
	return C.QGamepadKeyNavigation___children_newList(ptr.Pointer())
}

func (ptr *QGamepadKeyNavigation) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QGamepadKeyNavigation___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QGamepadKeyNavigation) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QGamepadKeyNavigation) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QGamepadKeyNavigation___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QGamepadKeyNavigation) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGamepadKeyNavigation___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGamepadKeyNavigation) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGamepadKeyNavigation) __findChildren_newList() unsafe.Pointer {
	return C.QGamepadKeyNavigation___findChildren_newList(ptr.Pointer())
}

func (ptr *QGamepadKeyNavigation) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGamepadKeyNavigation___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGamepadKeyNavigation) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGamepadKeyNavigation) __findChildren_newList3() unsafe.Pointer {
	return C.QGamepadKeyNavigation___findChildren_newList3(ptr.Pointer())
}

//export callbackQGamepadKeyNavigation_ChildEvent
func callbackQGamepadKeyNavigation_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQGamepadKeyNavigationFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QGamepadKeyNavigation) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQGamepadKeyNavigation_ConnectNotify
func callbackQGamepadKeyNavigation_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGamepadKeyNavigationFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGamepadKeyNavigation) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGamepadKeyNavigation_CustomEvent
func callbackQGamepadKeyNavigation_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQGamepadKeyNavigationFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGamepadKeyNavigation) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQGamepadKeyNavigation_DeleteLater
func callbackQGamepadKeyNavigation_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQGamepadKeyNavigationFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QGamepadKeyNavigation) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGamepadKeyNavigation_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQGamepadKeyNavigation_Destroyed
func callbackQGamepadKeyNavigation_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQGamepadKeyNavigation_DisconnectNotify
func callbackQGamepadKeyNavigation_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGamepadKeyNavigationFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGamepadKeyNavigation) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGamepadKeyNavigation_Event
func callbackQGamepadKeyNavigation_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGamepadKeyNavigationFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QGamepadKeyNavigation) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepadKeyNavigation_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQGamepadKeyNavigation_EventFilter
func callbackQGamepadKeyNavigation_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGamepadKeyNavigationFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QGamepadKeyNavigation) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepadKeyNavigation_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQGamepadKeyNavigation_MetaObject
func callbackQGamepadKeyNavigation_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQGamepadKeyNavigationFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QGamepadKeyNavigation) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QGamepadKeyNavigation_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQGamepadKeyNavigation_ObjectNameChanged
func callbackQGamepadKeyNavigation_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtGamepad_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQGamepadKeyNavigation_TimerEvent
func callbackQGamepadKeyNavigation_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGamepadKeyNavigationFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGamepadKeyNavigation) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QGamepadManager struct {
	core.QObject
}

type QGamepadManager_ITF interface {
	core.QObject_ITF
	QGamepadManager_PTR() *QGamepadManager
}

func (ptr *QGamepadManager) QGamepadManager_PTR() *QGamepadManager {
	return ptr
}

func (ptr *QGamepadManager) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QGamepadManager) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQGamepadManager(ptr QGamepadManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGamepadManager_PTR().Pointer()
	}
	return nil
}

func NewQGamepadManagerFromPointer(ptr unsafe.Pointer) (n *QGamepadManager) {
	n = new(QGamepadManager)
	n.SetPointer(ptr)
	return
}
func (ptr *QGamepadManager) ConnectedGamepads() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtGamepad_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewQGamepadManagerFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__connectedGamepads_atList(i)
			}
			return out
		}(C.QGamepadManager_ConnectedGamepads(ptr.Pointer()))
	}
	return make([]int, 0)
}

//export callbackQGamepadManager_ConnectedGamepadsChanged
func callbackQGamepadManager_ConnectedGamepadsChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectedGamepadsChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QGamepadManager) ConnectConnectedGamepadsChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "connectedGamepadsChanged") {
			C.QGamepadManager_ConnectConnectedGamepadsChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "connectedGamepadsChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "connectedGamepadsChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "connectedGamepadsChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "connectedGamepadsChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepadManager) DisconnectConnectedGamepadsChanged() {
	if ptr.Pointer() != nil {
		C.QGamepadManager_DisconnectConnectedGamepadsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "connectedGamepadsChanged")
	}
}

func (ptr *QGamepadManager) ConnectedGamepadsChanged() {
	if ptr.Pointer() != nil {
		C.QGamepadManager_ConnectedGamepadsChanged(ptr.Pointer())
	}
}

func (ptr *QGamepadManager) GamepadName(deviceId int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGamepadManager_GamepadName(ptr.Pointer(), C.int(int32(deviceId))))
	}
	return ""
}

func QGamepadManager_Instance() *QGamepadManager {
	tmpValue := NewQGamepadManagerFromPointer(C.QGamepadManager_QGamepadManager_Instance())
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QGamepadManager) Instance() *QGamepadManager {
	tmpValue := NewQGamepadManagerFromPointer(C.QGamepadManager_QGamepadManager_Instance())
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQGamepadManager_IsConfigurationNeeded
func callbackQGamepadManager_IsConfigurationNeeded(ptr unsafe.Pointer, deviceId C.int) C.char {
	if signal := qt.GetSignal(ptr, "isConfigurationNeeded"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int) bool)(signal))(int(int32(deviceId))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGamepadManagerFromPointer(ptr).IsConfigurationNeededDefault(int(int32(deviceId))))))
}

func (ptr *QGamepadManager) ConnectIsConfigurationNeeded(f func(deviceId int) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "isConfigurationNeeded"); signal != nil {
			f := func(deviceId int) bool {
				(*(*func(int) bool)(signal))(deviceId)
				return f(deviceId)
			}
			qt.ConnectSignal(ptr.Pointer(), "isConfigurationNeeded", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "isConfigurationNeeded", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepadManager) DisconnectIsConfigurationNeeded() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "isConfigurationNeeded")
	}
}

func (ptr *QGamepadManager) IsConfigurationNeeded(deviceId int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepadManager_IsConfigurationNeeded(ptr.Pointer(), C.int(int32(deviceId)))) != 0
	}
	return false
}

func (ptr *QGamepadManager) IsConfigurationNeededDefault(deviceId int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepadManager_IsConfigurationNeededDefault(ptr.Pointer(), C.int(int32(deviceId)))) != 0
	}
	return false
}

func (ptr *QGamepadManager) IsGamepadConnected(deviceId int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepadManager_IsGamepadConnected(ptr.Pointer(), C.int(int32(deviceId)))) != 0
	}
	return false
}

//export callbackQGamepadManager_ResetConfiguration
func callbackQGamepadManager_ResetConfiguration(ptr unsafe.Pointer, deviceId C.int) {
	if signal := qt.GetSignal(ptr, "resetConfiguration"); signal != nil {
		(*(*func(int))(signal))(int(int32(deviceId)))
	} else {
		NewQGamepadManagerFromPointer(ptr).ResetConfigurationDefault(int(int32(deviceId)))
	}
}

func (ptr *QGamepadManager) ConnectResetConfiguration(f func(deviceId int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "resetConfiguration"); signal != nil {
			f := func(deviceId int) {
				(*(*func(int))(signal))(deviceId)
				f(deviceId)
			}
			qt.ConnectSignal(ptr.Pointer(), "resetConfiguration", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "resetConfiguration", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepadManager) DisconnectResetConfiguration() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "resetConfiguration")
	}
}

func (ptr *QGamepadManager) ResetConfiguration(deviceId int) {
	if ptr.Pointer() != nil {
		C.QGamepadManager_ResetConfiguration(ptr.Pointer(), C.int(int32(deviceId)))
	}
}

func (ptr *QGamepadManager) ResetConfigurationDefault(deviceId int) {
	if ptr.Pointer() != nil {
		C.QGamepadManager_ResetConfigurationDefault(ptr.Pointer(), C.int(int32(deviceId)))
	}
}

//export callbackQGamepadManager_SetSettingsFile
func callbackQGamepadManager_SetSettingsFile(ptr unsafe.Pointer, file C.struct_QtGamepad_PackedString) {
	if signal := qt.GetSignal(ptr, "setSettingsFile"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(file))
	} else {
		NewQGamepadManagerFromPointer(ptr).SetSettingsFileDefault(cGoUnpackString(file))
	}
}

func (ptr *QGamepadManager) ConnectSetSettingsFile(f func(file string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setSettingsFile"); signal != nil {
			f := func(file string) {
				(*(*func(string))(signal))(file)
				f(file)
			}
			qt.ConnectSignal(ptr.Pointer(), "setSettingsFile", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setSettingsFile", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGamepadManager) DisconnectSetSettingsFile() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setSettingsFile")
	}
}

func (ptr *QGamepadManager) SetSettingsFile(file string) {
	if ptr.Pointer() != nil {
		var fileC *C.char
		if file != "" {
			fileC = C.CString(file)
			defer C.free(unsafe.Pointer(fileC))
		}
		C.QGamepadManager_SetSettingsFile(ptr.Pointer(), C.struct_QtGamepad_PackedString{data: fileC, len: C.longlong(len(file))})
	}
}

func (ptr *QGamepadManager) SetSettingsFileDefault(file string) {
	if ptr.Pointer() != nil {
		var fileC *C.char
		if file != "" {
			fileC = C.CString(file)
			defer C.free(unsafe.Pointer(fileC))
		}
		C.QGamepadManager_SetSettingsFileDefault(ptr.Pointer(), C.struct_QtGamepad_PackedString{data: fileC, len: C.longlong(len(file))})
	}
}

func (ptr *QGamepadManager) __connectedGamepads_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QGamepadManager___connectedGamepads_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QGamepadManager) __connectedGamepads_setList(i int) {
	if ptr.Pointer() != nil {
		C.QGamepadManager___connectedGamepads_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QGamepadManager) __connectedGamepads_newList() unsafe.Pointer {
	return C.QGamepadManager___connectedGamepads_newList(ptr.Pointer())
}

func (ptr *QGamepadManager) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGamepadManager___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGamepadManager) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGamepadManager___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGamepadManager) __children_newList() unsafe.Pointer {
	return C.QGamepadManager___children_newList(ptr.Pointer())
}

func (ptr *QGamepadManager) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QGamepadManager___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QGamepadManager) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QGamepadManager___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QGamepadManager) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QGamepadManager___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QGamepadManager) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGamepadManager___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGamepadManager) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGamepadManager___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGamepadManager) __findChildren_newList() unsafe.Pointer {
	return C.QGamepadManager___findChildren_newList(ptr.Pointer())
}

func (ptr *QGamepadManager) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGamepadManager___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGamepadManager) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGamepadManager___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGamepadManager) __findChildren_newList3() unsafe.Pointer {
	return C.QGamepadManager___findChildren_newList3(ptr.Pointer())
}

//export callbackQGamepadManager_ChildEvent
func callbackQGamepadManager_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQGamepadManagerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QGamepadManager) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGamepadManager_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQGamepadManager_ConnectNotify
func callbackQGamepadManager_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGamepadManagerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGamepadManager) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGamepadManager_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGamepadManager_CustomEvent
func callbackQGamepadManager_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQGamepadManagerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGamepadManager) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGamepadManager_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQGamepadManager_DeleteLater
func callbackQGamepadManager_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQGamepadManagerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QGamepadManager) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGamepadManager_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQGamepadManager_Destroyed
func callbackQGamepadManager_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQGamepadManager_DisconnectNotify
func callbackQGamepadManager_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGamepadManagerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGamepadManager) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGamepadManager_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGamepadManager_Event
func callbackQGamepadManager_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGamepadManagerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QGamepadManager) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepadManager_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQGamepadManager_EventFilter
func callbackQGamepadManager_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGamepadManagerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QGamepadManager) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepadManager_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQGamepadManager_MetaObject
func callbackQGamepadManager_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQGamepadManagerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QGamepadManager) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QGamepadManager_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQGamepadManager_ObjectNameChanged
func callbackQGamepadManager_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtGamepad_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQGamepadManager_TimerEvent
func callbackQGamepadManager_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGamepadManagerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGamepadManager) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGamepadManager_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func init() {
	qt.ItfMap["gamepad.QGamepad_ITF"] = QGamepad{}
	qt.FuncMap["gamepad.NewQGamepad"] = NewQGamepad
	qt.ItfMap["gamepad.QGamepadKeyNavigation_ITF"] = QGamepadKeyNavigation{}
	qt.FuncMap["gamepad.NewQGamepadKeyNavigation"] = NewQGamepadKeyNavigation
	qt.ItfMap["gamepad.QGamepadManager_ITF"] = QGamepadManager{}
	qt.FuncMap["gamepad.QGamepadManager_Instance"] = QGamepadManager_Instance
}
