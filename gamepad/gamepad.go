// +build !minimal

package gamepad

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "gamepad.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"runtime"
	"unsafe"
)

func cGoUnpackString(s C.struct_QtGamepad_PackedString) string {
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtGamepad_PackedString) []byte {
	if int(s.len) == -1 {
		return []byte(C.GoString(s.data))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
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

func QGamepad_Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QGamepad_QGamepad_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QGamepad) Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QGamepad_QGamepad_Tr(sC, cC, C.int(int32(n))))
}

func QGamepad_TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QGamepad_QGamepad_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QGamepad) TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QGamepad_QGamepad_TrUtf8(sC, cC, C.int(int32(n))))
}

//export callbackQGamepad_AxisLeftXChanged
func callbackQGamepad_AxisLeftXChanged(ptr unsafe.Pointer, value C.double) {
	if signal := qt.GetSignal(ptr, "axisLeftXChanged"); signal != nil {
		signal.(func(float64))(float64(value))
	}

}

func (ptr *QGamepad) ConnectAxisLeftXChanged(f func(value float64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "axisLeftXChanged") {
			C.QGamepad_ConnectAxisLeftXChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "axisLeftXChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "axisLeftXChanged", func(value float64) {
				signal.(func(float64))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "axisLeftXChanged", f)
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

//export callbackQGamepad_AxisLeftYChanged
func callbackQGamepad_AxisLeftYChanged(ptr unsafe.Pointer, value C.double) {
	if signal := qt.GetSignal(ptr, "axisLeftYChanged"); signal != nil {
		signal.(func(float64))(float64(value))
	}

}

func (ptr *QGamepad) ConnectAxisLeftYChanged(f func(value float64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "axisLeftYChanged") {
			C.QGamepad_ConnectAxisLeftYChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "axisLeftYChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "axisLeftYChanged", func(value float64) {
				signal.(func(float64))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "axisLeftYChanged", f)
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

//export callbackQGamepad_AxisRightXChanged
func callbackQGamepad_AxisRightXChanged(ptr unsafe.Pointer, value C.double) {
	if signal := qt.GetSignal(ptr, "axisRightXChanged"); signal != nil {
		signal.(func(float64))(float64(value))
	}

}

func (ptr *QGamepad) ConnectAxisRightXChanged(f func(value float64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "axisRightXChanged") {
			C.QGamepad_ConnectAxisRightXChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "axisRightXChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "axisRightXChanged", func(value float64) {
				signal.(func(float64))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "axisRightXChanged", f)
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

//export callbackQGamepad_AxisRightYChanged
func callbackQGamepad_AxisRightYChanged(ptr unsafe.Pointer, value C.double) {
	if signal := qt.GetSignal(ptr, "axisRightYChanged"); signal != nil {
		signal.(func(float64))(float64(value))
	}

}

func (ptr *QGamepad) ConnectAxisRightYChanged(f func(value float64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "axisRightYChanged") {
			C.QGamepad_ConnectAxisRightYChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "axisRightYChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "axisRightYChanged", func(value float64) {
				signal.(func(float64))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "axisRightYChanged", f)
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

//export callbackQGamepad_ButtonAChanged
func callbackQGamepad_ButtonAChanged(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(ptr, "buttonAChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonAChanged(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonAChanged") {
			C.QGamepad_ConnectButtonAChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonAChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "buttonAChanged", func(value bool) {
				signal.(func(bool))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonAChanged", f)
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

//export callbackQGamepad_ButtonBChanged
func callbackQGamepad_ButtonBChanged(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(ptr, "buttonBChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonBChanged(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonBChanged") {
			C.QGamepad_ConnectButtonBChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonBChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "buttonBChanged", func(value bool) {
				signal.(func(bool))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonBChanged", f)
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

//export callbackQGamepad_ButtonCenterChanged
func callbackQGamepad_ButtonCenterChanged(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(ptr, "buttonCenterChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonCenterChanged(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonCenterChanged") {
			C.QGamepad_ConnectButtonCenterChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonCenterChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "buttonCenterChanged", func(value bool) {
				signal.(func(bool))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonCenterChanged", f)
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

//export callbackQGamepad_ButtonDownChanged
func callbackQGamepad_ButtonDownChanged(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(ptr, "buttonDownChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonDownChanged(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonDownChanged") {
			C.QGamepad_ConnectButtonDownChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonDownChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "buttonDownChanged", func(value bool) {
				signal.(func(bool))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonDownChanged", f)
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

//export callbackQGamepad_ButtonGuideChanged
func callbackQGamepad_ButtonGuideChanged(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(ptr, "buttonGuideChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonGuideChanged(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonGuideChanged") {
			C.QGamepad_ConnectButtonGuideChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonGuideChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "buttonGuideChanged", func(value bool) {
				signal.(func(bool))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonGuideChanged", f)
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

//export callbackQGamepad_ButtonL1Changed
func callbackQGamepad_ButtonL1Changed(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(ptr, "buttonL1Changed"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonL1Changed(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonL1Changed") {
			C.QGamepad_ConnectButtonL1Changed(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonL1Changed"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "buttonL1Changed", func(value bool) {
				signal.(func(bool))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonL1Changed", f)
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

//export callbackQGamepad_ButtonL2Changed
func callbackQGamepad_ButtonL2Changed(ptr unsafe.Pointer, value C.double) {
	if signal := qt.GetSignal(ptr, "buttonL2Changed"); signal != nil {
		signal.(func(float64))(float64(value))
	}

}

func (ptr *QGamepad) ConnectButtonL2Changed(f func(value float64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonL2Changed") {
			C.QGamepad_ConnectButtonL2Changed(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonL2Changed"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "buttonL2Changed", func(value float64) {
				signal.(func(float64))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonL2Changed", f)
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

//export callbackQGamepad_ButtonL3Changed
func callbackQGamepad_ButtonL3Changed(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(ptr, "buttonL3Changed"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonL3Changed(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonL3Changed") {
			C.QGamepad_ConnectButtonL3Changed(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonL3Changed"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "buttonL3Changed", func(value bool) {
				signal.(func(bool))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonL3Changed", f)
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

//export callbackQGamepad_ButtonLeftChanged
func callbackQGamepad_ButtonLeftChanged(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(ptr, "buttonLeftChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonLeftChanged(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonLeftChanged") {
			C.QGamepad_ConnectButtonLeftChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonLeftChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "buttonLeftChanged", func(value bool) {
				signal.(func(bool))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonLeftChanged", f)
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

//export callbackQGamepad_ButtonR1Changed
func callbackQGamepad_ButtonR1Changed(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(ptr, "buttonR1Changed"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonR1Changed(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonR1Changed") {
			C.QGamepad_ConnectButtonR1Changed(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonR1Changed"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "buttonR1Changed", func(value bool) {
				signal.(func(bool))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonR1Changed", f)
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

//export callbackQGamepad_ButtonR2Changed
func callbackQGamepad_ButtonR2Changed(ptr unsafe.Pointer, value C.double) {
	if signal := qt.GetSignal(ptr, "buttonR2Changed"); signal != nil {
		signal.(func(float64))(float64(value))
	}

}

func (ptr *QGamepad) ConnectButtonR2Changed(f func(value float64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonR2Changed") {
			C.QGamepad_ConnectButtonR2Changed(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonR2Changed"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "buttonR2Changed", func(value float64) {
				signal.(func(float64))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonR2Changed", f)
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

//export callbackQGamepad_ButtonR3Changed
func callbackQGamepad_ButtonR3Changed(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(ptr, "buttonR3Changed"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonR3Changed(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonR3Changed") {
			C.QGamepad_ConnectButtonR3Changed(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonR3Changed"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "buttonR3Changed", func(value bool) {
				signal.(func(bool))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonR3Changed", f)
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

//export callbackQGamepad_ButtonRightChanged
func callbackQGamepad_ButtonRightChanged(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(ptr, "buttonRightChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonRightChanged(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonRightChanged") {
			C.QGamepad_ConnectButtonRightChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonRightChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "buttonRightChanged", func(value bool) {
				signal.(func(bool))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonRightChanged", f)
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

//export callbackQGamepad_ButtonSelectChanged
func callbackQGamepad_ButtonSelectChanged(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(ptr, "buttonSelectChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonSelectChanged(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonSelectChanged") {
			C.QGamepad_ConnectButtonSelectChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonSelectChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "buttonSelectChanged", func(value bool) {
				signal.(func(bool))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonSelectChanged", f)
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

//export callbackQGamepad_ButtonStartChanged
func callbackQGamepad_ButtonStartChanged(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(ptr, "buttonStartChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonStartChanged(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonStartChanged") {
			C.QGamepad_ConnectButtonStartChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonStartChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "buttonStartChanged", func(value bool) {
				signal.(func(bool))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonStartChanged", f)
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

//export callbackQGamepad_ButtonUpChanged
func callbackQGamepad_ButtonUpChanged(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(ptr, "buttonUpChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonUpChanged(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonUpChanged") {
			C.QGamepad_ConnectButtonUpChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonUpChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "buttonUpChanged", func(value bool) {
				signal.(func(bool))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonUpChanged", f)
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

//export callbackQGamepad_ButtonXChanged
func callbackQGamepad_ButtonXChanged(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(ptr, "buttonXChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonXChanged(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonXChanged") {
			C.QGamepad_ConnectButtonXChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonXChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "buttonXChanged", func(value bool) {
				signal.(func(bool))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonXChanged", f)
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

//export callbackQGamepad_ButtonYChanged
func callbackQGamepad_ButtonYChanged(ptr unsafe.Pointer, value C.char) {
	if signal := qt.GetSignal(ptr, "buttonYChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonYChanged(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonYChanged") {
			C.QGamepad_ConnectButtonYChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonYChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "buttonYChanged", func(value bool) {
				signal.(func(bool))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonYChanged", f)
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
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectConnectedChanged(f func(value bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "connectedChanged") {
			C.QGamepad_ConnectConnectedChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "connectedChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "connectedChanged", func(value bool) {
				signal.(func(bool))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "connectedChanged", f)
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

//export callbackQGamepad_DeviceIdChanged
func callbackQGamepad_DeviceIdChanged(ptr unsafe.Pointer, value C.int) {
	if signal := qt.GetSignal(ptr, "deviceIdChanged"); signal != nil {
		signal.(func(int))(int(int32(value)))
	}

}

func (ptr *QGamepad) ConnectDeviceIdChanged(f func(value int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "deviceIdChanged") {
			C.QGamepad_ConnectDeviceIdChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "deviceIdChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "deviceIdChanged", func(value int) {
				signal.(func(int))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "deviceIdChanged", f)
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

//export callbackQGamepad_NameChanged
func callbackQGamepad_NameChanged(ptr unsafe.Pointer, value C.struct_QtGamepad_PackedString) {
	if signal := qt.GetSignal(ptr, "nameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(value))
	}

}

func (ptr *QGamepad) ConnectNameChanged(f func(value string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "nameChanged") {
			C.QGamepad_ConnectNameChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "nameChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "nameChanged", func(value string) {
				signal.(func(string))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "nameChanged", f)
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
		signal.(func(int))(int(int32(number)))
	} else {
		NewQGamepadFromPointer(ptr).SetDeviceIdDefault(int(int32(number)))
	}
}

func (ptr *QGamepad) ConnectSetDeviceId(f func(number int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setDeviceId"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setDeviceId", func(number int) {
				signal.(func(int))(number)
				f(number)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setDeviceId", f)
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

//export callbackQGamepad_DestroyQGamepad
func callbackQGamepad_DestroyQGamepad(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QGamepad"); signal != nil {
		signal.(func())()
	} else {
		NewQGamepadFromPointer(ptr).DestroyQGamepadDefault()
	}
}

func (ptr *QGamepad) ConnectDestroyQGamepad(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QGamepad"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QGamepad", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QGamepad", f)
		}
	}
}

func (ptr *QGamepad) DisconnectDestroyQGamepad() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QGamepad")
	}
}

func (ptr *QGamepad) DestroyQGamepad() {
	if ptr.Pointer() != nil {
		C.QGamepad_DestroyQGamepad(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QGamepad) DestroyQGamepadDefault() {
	if ptr.Pointer() != nil {
		C.QGamepad_DestroyQGamepadDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
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
		return int8(C.QGamepad_ButtonA(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGamepad) ButtonB() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepad_ButtonB(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGamepad) ButtonCenter() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepad_ButtonCenter(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGamepad) ButtonDown() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepad_ButtonDown(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGamepad) ButtonGuide() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepad_ButtonGuide(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGamepad) ButtonL1() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepad_ButtonL1(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGamepad) ButtonL3() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepad_ButtonL3(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGamepad) ButtonLeft() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepad_ButtonLeft(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGamepad) ButtonR1() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepad_ButtonR1(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGamepad) ButtonR3() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepad_ButtonR3(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGamepad) ButtonRight() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepad_ButtonRight(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGamepad) ButtonSelect() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepad_ButtonSelect(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGamepad) ButtonStart() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepad_ButtonStart(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGamepad) ButtonUp() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepad_ButtonUp(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGamepad) ButtonX() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepad_ButtonX(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGamepad) ButtonY() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepad_ButtonY(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGamepad) IsConnected() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepad_IsConnected(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQGamepad_MetaObject
func callbackQGamepad_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
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
		tmpValue := core.NewQByteArrayFromPointer(C.QGamepad___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
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
	return C.QGamepad___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QGamepad) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGamepad___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
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
	return C.QGamepad___findChildren_newList2(ptr.Pointer())
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

//export callbackQGamepad_Event
func callbackQGamepad_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
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
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGamepadFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QGamepad) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepad_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQGamepad_ChildEvent
func callbackQGamepad_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
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
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
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
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
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
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQGamepadFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QGamepad) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QGamepad_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQGamepad_Destroyed
func callbackQGamepad_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQGamepad_DisconnectNotify
func callbackQGamepad_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
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
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQGamepad_TimerEvent
func callbackQGamepad_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
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

func QGamepadKeyNavigation_Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QGamepadKeyNavigation_QGamepadKeyNavigation_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QGamepadKeyNavigation) Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QGamepadKeyNavigation_QGamepadKeyNavigation_Tr(sC, cC, C.int(int32(n))))
}

func QGamepadKeyNavigation_TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QGamepadKeyNavigation_QGamepadKeyNavigation_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QGamepadKeyNavigation) TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QGamepadKeyNavigation_QGamepadKeyNavigation_TrUtf8(sC, cC, C.int(int32(n))))
}

//export callbackQGamepadKeyNavigation_ActiveChanged
func callbackQGamepadKeyNavigation_ActiveChanged(ptr unsafe.Pointer, isActive C.char) {
	if signal := qt.GetSignal(ptr, "activeChanged"); signal != nil {
		signal.(func(bool))(int8(isActive) != 0)
	}

}

func (ptr *QGamepadKeyNavigation) ConnectActiveChanged(f func(isActive bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "activeChanged") {
			C.QGamepadKeyNavigation_ConnectActiveChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "activeChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "activeChanged", func(isActive bool) {
				signal.(func(bool))(isActive)
				f(isActive)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "activeChanged", f)
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

//export callbackQGamepadKeyNavigation_ButtonAKeyChanged
func callbackQGamepadKeyNavigation_ButtonAKeyChanged(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "buttonAKeyChanged"); signal != nil {
		signal.(func(core.Qt__Key))(core.Qt__Key(key))
	}

}

func (ptr *QGamepadKeyNavigation) ConnectButtonAKeyChanged(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonAKeyChanged") {
			C.QGamepadKeyNavigation_ConnectButtonAKeyChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonAKeyChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "buttonAKeyChanged", func(key core.Qt__Key) {
				signal.(func(core.Qt__Key))(key)
				f(key)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonAKeyChanged", f)
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

//export callbackQGamepadKeyNavigation_ButtonBKeyChanged
func callbackQGamepadKeyNavigation_ButtonBKeyChanged(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "buttonBKeyChanged"); signal != nil {
		signal.(func(core.Qt__Key))(core.Qt__Key(key))
	}

}

func (ptr *QGamepadKeyNavigation) ConnectButtonBKeyChanged(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonBKeyChanged") {
			C.QGamepadKeyNavigation_ConnectButtonBKeyChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonBKeyChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "buttonBKeyChanged", func(key core.Qt__Key) {
				signal.(func(core.Qt__Key))(key)
				f(key)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonBKeyChanged", f)
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

//export callbackQGamepadKeyNavigation_ButtonGuideKeyChanged
func callbackQGamepadKeyNavigation_ButtonGuideKeyChanged(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "buttonGuideKeyChanged"); signal != nil {
		signal.(func(core.Qt__Key))(core.Qt__Key(key))
	}

}

func (ptr *QGamepadKeyNavigation) ConnectButtonGuideKeyChanged(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonGuideKeyChanged") {
			C.QGamepadKeyNavigation_ConnectButtonGuideKeyChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonGuideKeyChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "buttonGuideKeyChanged", func(key core.Qt__Key) {
				signal.(func(core.Qt__Key))(key)
				f(key)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonGuideKeyChanged", f)
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

//export callbackQGamepadKeyNavigation_ButtonL1KeyChanged
func callbackQGamepadKeyNavigation_ButtonL1KeyChanged(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "buttonL1KeyChanged"); signal != nil {
		signal.(func(core.Qt__Key))(core.Qt__Key(key))
	}

}

func (ptr *QGamepadKeyNavigation) ConnectButtonL1KeyChanged(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonL1KeyChanged") {
			C.QGamepadKeyNavigation_ConnectButtonL1KeyChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonL1KeyChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "buttonL1KeyChanged", func(key core.Qt__Key) {
				signal.(func(core.Qt__Key))(key)
				f(key)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonL1KeyChanged", f)
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

//export callbackQGamepadKeyNavigation_ButtonL2KeyChanged
func callbackQGamepadKeyNavigation_ButtonL2KeyChanged(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "buttonL2KeyChanged"); signal != nil {
		signal.(func(core.Qt__Key))(core.Qt__Key(key))
	}

}

func (ptr *QGamepadKeyNavigation) ConnectButtonL2KeyChanged(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonL2KeyChanged") {
			C.QGamepadKeyNavigation_ConnectButtonL2KeyChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonL2KeyChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "buttonL2KeyChanged", func(key core.Qt__Key) {
				signal.(func(core.Qt__Key))(key)
				f(key)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonL2KeyChanged", f)
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

//export callbackQGamepadKeyNavigation_ButtonL3KeyChanged
func callbackQGamepadKeyNavigation_ButtonL3KeyChanged(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "buttonL3KeyChanged"); signal != nil {
		signal.(func(core.Qt__Key))(core.Qt__Key(key))
	}

}

func (ptr *QGamepadKeyNavigation) ConnectButtonL3KeyChanged(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonL3KeyChanged") {
			C.QGamepadKeyNavigation_ConnectButtonL3KeyChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonL3KeyChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "buttonL3KeyChanged", func(key core.Qt__Key) {
				signal.(func(core.Qt__Key))(key)
				f(key)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonL3KeyChanged", f)
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

//export callbackQGamepadKeyNavigation_ButtonR1KeyChanged
func callbackQGamepadKeyNavigation_ButtonR1KeyChanged(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "buttonR1KeyChanged"); signal != nil {
		signal.(func(core.Qt__Key))(core.Qt__Key(key))
	}

}

func (ptr *QGamepadKeyNavigation) ConnectButtonR1KeyChanged(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonR1KeyChanged") {
			C.QGamepadKeyNavigation_ConnectButtonR1KeyChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonR1KeyChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "buttonR1KeyChanged", func(key core.Qt__Key) {
				signal.(func(core.Qt__Key))(key)
				f(key)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonR1KeyChanged", f)
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

//export callbackQGamepadKeyNavigation_ButtonR2KeyChanged
func callbackQGamepadKeyNavigation_ButtonR2KeyChanged(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "buttonR2KeyChanged"); signal != nil {
		signal.(func(core.Qt__Key))(core.Qt__Key(key))
	}

}

func (ptr *QGamepadKeyNavigation) ConnectButtonR2KeyChanged(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonR2KeyChanged") {
			C.QGamepadKeyNavigation_ConnectButtonR2KeyChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonR2KeyChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "buttonR2KeyChanged", func(key core.Qt__Key) {
				signal.(func(core.Qt__Key))(key)
				f(key)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonR2KeyChanged", f)
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

//export callbackQGamepadKeyNavigation_ButtonR3KeyChanged
func callbackQGamepadKeyNavigation_ButtonR3KeyChanged(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "buttonR3KeyChanged"); signal != nil {
		signal.(func(core.Qt__Key))(core.Qt__Key(key))
	}

}

func (ptr *QGamepadKeyNavigation) ConnectButtonR3KeyChanged(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonR3KeyChanged") {
			C.QGamepadKeyNavigation_ConnectButtonR3KeyChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonR3KeyChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "buttonR3KeyChanged", func(key core.Qt__Key) {
				signal.(func(core.Qt__Key))(key)
				f(key)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonR3KeyChanged", f)
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

//export callbackQGamepadKeyNavigation_ButtonSelectKeyChanged
func callbackQGamepadKeyNavigation_ButtonSelectKeyChanged(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "buttonSelectKeyChanged"); signal != nil {
		signal.(func(core.Qt__Key))(core.Qt__Key(key))
	}

}

func (ptr *QGamepadKeyNavigation) ConnectButtonSelectKeyChanged(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonSelectKeyChanged") {
			C.QGamepadKeyNavigation_ConnectButtonSelectKeyChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonSelectKeyChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "buttonSelectKeyChanged", func(key core.Qt__Key) {
				signal.(func(core.Qt__Key))(key)
				f(key)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonSelectKeyChanged", f)
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

//export callbackQGamepadKeyNavigation_ButtonStartKeyChanged
func callbackQGamepadKeyNavigation_ButtonStartKeyChanged(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "buttonStartKeyChanged"); signal != nil {
		signal.(func(core.Qt__Key))(core.Qt__Key(key))
	}

}

func (ptr *QGamepadKeyNavigation) ConnectButtonStartKeyChanged(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonStartKeyChanged") {
			C.QGamepadKeyNavigation_ConnectButtonStartKeyChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonStartKeyChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "buttonStartKeyChanged", func(key core.Qt__Key) {
				signal.(func(core.Qt__Key))(key)
				f(key)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonStartKeyChanged", f)
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

//export callbackQGamepadKeyNavigation_ButtonXKeyChanged
func callbackQGamepadKeyNavigation_ButtonXKeyChanged(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "buttonXKeyChanged"); signal != nil {
		signal.(func(core.Qt__Key))(core.Qt__Key(key))
	}

}

func (ptr *QGamepadKeyNavigation) ConnectButtonXKeyChanged(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonXKeyChanged") {
			C.QGamepadKeyNavigation_ConnectButtonXKeyChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonXKeyChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "buttonXKeyChanged", func(key core.Qt__Key) {
				signal.(func(core.Qt__Key))(key)
				f(key)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonXKeyChanged", f)
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

//export callbackQGamepadKeyNavigation_ButtonYKeyChanged
func callbackQGamepadKeyNavigation_ButtonYKeyChanged(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "buttonYKeyChanged"); signal != nil {
		signal.(func(core.Qt__Key))(core.Qt__Key(key))
	}

}

func (ptr *QGamepadKeyNavigation) ConnectButtonYKeyChanged(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonYKeyChanged") {
			C.QGamepadKeyNavigation_ConnectButtonYKeyChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonYKeyChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "buttonYKeyChanged", func(key core.Qt__Key) {
				signal.(func(core.Qt__Key))(key)
				f(key)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonYKeyChanged", f)
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

//export callbackQGamepadKeyNavigation_DownKeyChanged
func callbackQGamepadKeyNavigation_DownKeyChanged(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "downKeyChanged"); signal != nil {
		signal.(func(core.Qt__Key))(core.Qt__Key(key))
	}

}

func (ptr *QGamepadKeyNavigation) ConnectDownKeyChanged(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "downKeyChanged") {
			C.QGamepadKeyNavigation_ConnectDownKeyChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "downKeyChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "downKeyChanged", func(key core.Qt__Key) {
				signal.(func(core.Qt__Key))(key)
				f(key)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "downKeyChanged", f)
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

//export callbackQGamepadKeyNavigation_GamepadChanged
func callbackQGamepadKeyNavigation_GamepadChanged(ptr unsafe.Pointer, gamepad unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "gamepadChanged"); signal != nil {
		signal.(func(*QGamepad))(NewQGamepadFromPointer(gamepad))
	}

}

func (ptr *QGamepadKeyNavigation) ConnectGamepadChanged(f func(gamepad *QGamepad)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "gamepadChanged") {
			C.QGamepadKeyNavigation_ConnectGamepadChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "gamepadChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "gamepadChanged", func(gamepad *QGamepad) {
				signal.(func(*QGamepad))(gamepad)
				f(gamepad)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "gamepadChanged", f)
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

//export callbackQGamepadKeyNavigation_LeftKeyChanged
func callbackQGamepadKeyNavigation_LeftKeyChanged(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "leftKeyChanged"); signal != nil {
		signal.(func(core.Qt__Key))(core.Qt__Key(key))
	}

}

func (ptr *QGamepadKeyNavigation) ConnectLeftKeyChanged(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "leftKeyChanged") {
			C.QGamepadKeyNavigation_ConnectLeftKeyChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "leftKeyChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "leftKeyChanged", func(key core.Qt__Key) {
				signal.(func(core.Qt__Key))(key)
				f(key)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "leftKeyChanged", f)
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

//export callbackQGamepadKeyNavigation_RightKeyChanged
func callbackQGamepadKeyNavigation_RightKeyChanged(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "rightKeyChanged"); signal != nil {
		signal.(func(core.Qt__Key))(core.Qt__Key(key))
	}

}

func (ptr *QGamepadKeyNavigation) ConnectRightKeyChanged(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "rightKeyChanged") {
			C.QGamepadKeyNavigation_ConnectRightKeyChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "rightKeyChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "rightKeyChanged", func(key core.Qt__Key) {
				signal.(func(core.Qt__Key))(key)
				f(key)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "rightKeyChanged", f)
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
		signal.(func(bool))(int8(isActive) != 0)
	} else {
		NewQGamepadKeyNavigationFromPointer(ptr).SetActiveDefault(int8(isActive) != 0)
	}
}

func (ptr *QGamepadKeyNavigation) ConnectSetActive(f func(isActive bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setActive"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setActive", func(isActive bool) {
				signal.(func(bool))(isActive)
				f(isActive)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setActive", f)
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
		signal.(func(core.Qt__Key))(core.Qt__Key(key))
	} else {
		NewQGamepadKeyNavigationFromPointer(ptr).SetButtonAKeyDefault(core.Qt__Key(key))
	}
}

func (ptr *QGamepadKeyNavigation) ConnectSetButtonAKey(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setButtonAKey"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setButtonAKey", func(key core.Qt__Key) {
				signal.(func(core.Qt__Key))(key)
				f(key)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setButtonAKey", f)
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
		signal.(func(core.Qt__Key))(core.Qt__Key(key))
	} else {
		NewQGamepadKeyNavigationFromPointer(ptr).SetButtonBKeyDefault(core.Qt__Key(key))
	}
}

func (ptr *QGamepadKeyNavigation) ConnectSetButtonBKey(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setButtonBKey"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setButtonBKey", func(key core.Qt__Key) {
				signal.(func(core.Qt__Key))(key)
				f(key)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setButtonBKey", f)
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
		signal.(func(core.Qt__Key))(core.Qt__Key(key))
	} else {
		NewQGamepadKeyNavigationFromPointer(ptr).SetButtonGuideKeyDefault(core.Qt__Key(key))
	}
}

func (ptr *QGamepadKeyNavigation) ConnectSetButtonGuideKey(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setButtonGuideKey"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setButtonGuideKey", func(key core.Qt__Key) {
				signal.(func(core.Qt__Key))(key)
				f(key)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setButtonGuideKey", f)
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
		signal.(func(core.Qt__Key))(core.Qt__Key(key))
	} else {
		NewQGamepadKeyNavigationFromPointer(ptr).SetButtonL1KeyDefault(core.Qt__Key(key))
	}
}

func (ptr *QGamepadKeyNavigation) ConnectSetButtonL1Key(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setButtonL1Key"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setButtonL1Key", func(key core.Qt__Key) {
				signal.(func(core.Qt__Key))(key)
				f(key)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setButtonL1Key", f)
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
		signal.(func(core.Qt__Key))(core.Qt__Key(key))
	} else {
		NewQGamepadKeyNavigationFromPointer(ptr).SetButtonL2KeyDefault(core.Qt__Key(key))
	}
}

func (ptr *QGamepadKeyNavigation) ConnectSetButtonL2Key(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setButtonL2Key"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setButtonL2Key", func(key core.Qt__Key) {
				signal.(func(core.Qt__Key))(key)
				f(key)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setButtonL2Key", f)
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
		signal.(func(core.Qt__Key))(core.Qt__Key(key))
	} else {
		NewQGamepadKeyNavigationFromPointer(ptr).SetButtonL3KeyDefault(core.Qt__Key(key))
	}
}

func (ptr *QGamepadKeyNavigation) ConnectSetButtonL3Key(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setButtonL3Key"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setButtonL3Key", func(key core.Qt__Key) {
				signal.(func(core.Qt__Key))(key)
				f(key)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setButtonL3Key", f)
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
		signal.(func(core.Qt__Key))(core.Qt__Key(key))
	} else {
		NewQGamepadKeyNavigationFromPointer(ptr).SetButtonR1KeyDefault(core.Qt__Key(key))
	}
}

func (ptr *QGamepadKeyNavigation) ConnectSetButtonR1Key(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setButtonR1Key"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setButtonR1Key", func(key core.Qt__Key) {
				signal.(func(core.Qt__Key))(key)
				f(key)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setButtonR1Key", f)
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
		signal.(func(core.Qt__Key))(core.Qt__Key(key))
	} else {
		NewQGamepadKeyNavigationFromPointer(ptr).SetButtonR2KeyDefault(core.Qt__Key(key))
	}
}

func (ptr *QGamepadKeyNavigation) ConnectSetButtonR2Key(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setButtonR2Key"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setButtonR2Key", func(key core.Qt__Key) {
				signal.(func(core.Qt__Key))(key)
				f(key)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setButtonR2Key", f)
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
		signal.(func(core.Qt__Key))(core.Qt__Key(key))
	} else {
		NewQGamepadKeyNavigationFromPointer(ptr).SetButtonR3KeyDefault(core.Qt__Key(key))
	}
}

func (ptr *QGamepadKeyNavigation) ConnectSetButtonR3Key(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setButtonR3Key"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setButtonR3Key", func(key core.Qt__Key) {
				signal.(func(core.Qt__Key))(key)
				f(key)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setButtonR3Key", f)
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
		signal.(func(core.Qt__Key))(core.Qt__Key(key))
	} else {
		NewQGamepadKeyNavigationFromPointer(ptr).SetButtonSelectKeyDefault(core.Qt__Key(key))
	}
}

func (ptr *QGamepadKeyNavigation) ConnectSetButtonSelectKey(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setButtonSelectKey"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setButtonSelectKey", func(key core.Qt__Key) {
				signal.(func(core.Qt__Key))(key)
				f(key)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setButtonSelectKey", f)
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
		signal.(func(core.Qt__Key))(core.Qt__Key(key))
	} else {
		NewQGamepadKeyNavigationFromPointer(ptr).SetButtonStartKeyDefault(core.Qt__Key(key))
	}
}

func (ptr *QGamepadKeyNavigation) ConnectSetButtonStartKey(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setButtonStartKey"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setButtonStartKey", func(key core.Qt__Key) {
				signal.(func(core.Qt__Key))(key)
				f(key)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setButtonStartKey", f)
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
		signal.(func(core.Qt__Key))(core.Qt__Key(key))
	} else {
		NewQGamepadKeyNavigationFromPointer(ptr).SetButtonXKeyDefault(core.Qt__Key(key))
	}
}

func (ptr *QGamepadKeyNavigation) ConnectSetButtonXKey(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setButtonXKey"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setButtonXKey", func(key core.Qt__Key) {
				signal.(func(core.Qt__Key))(key)
				f(key)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setButtonXKey", f)
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
		signal.(func(core.Qt__Key))(core.Qt__Key(key))
	} else {
		NewQGamepadKeyNavigationFromPointer(ptr).SetButtonYKeyDefault(core.Qt__Key(key))
	}
}

func (ptr *QGamepadKeyNavigation) ConnectSetButtonYKey(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setButtonYKey"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setButtonYKey", func(key core.Qt__Key) {
				signal.(func(core.Qt__Key))(key)
				f(key)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setButtonYKey", f)
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
		signal.(func(core.Qt__Key))(core.Qt__Key(key))
	} else {
		NewQGamepadKeyNavigationFromPointer(ptr).SetDownKeyDefault(core.Qt__Key(key))
	}
}

func (ptr *QGamepadKeyNavigation) ConnectSetDownKey(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setDownKey"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setDownKey", func(key core.Qt__Key) {
				signal.(func(core.Qt__Key))(key)
				f(key)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setDownKey", f)
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
		signal.(func(*QGamepad))(NewQGamepadFromPointer(gamepad))
	} else {
		NewQGamepadKeyNavigationFromPointer(ptr).SetGamepadDefault(NewQGamepadFromPointer(gamepad))
	}
}

func (ptr *QGamepadKeyNavigation) ConnectSetGamepad(f func(gamepad *QGamepad)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setGamepad"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setGamepad", func(gamepad *QGamepad) {
				signal.(func(*QGamepad))(gamepad)
				f(gamepad)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setGamepad", f)
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
		signal.(func(core.Qt__Key))(core.Qt__Key(key))
	} else {
		NewQGamepadKeyNavigationFromPointer(ptr).SetLeftKeyDefault(core.Qt__Key(key))
	}
}

func (ptr *QGamepadKeyNavigation) ConnectSetLeftKey(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setLeftKey"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setLeftKey", func(key core.Qt__Key) {
				signal.(func(core.Qt__Key))(key)
				f(key)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setLeftKey", f)
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
		signal.(func(core.Qt__Key))(core.Qt__Key(key))
	} else {
		NewQGamepadKeyNavigationFromPointer(ptr).SetRightKeyDefault(core.Qt__Key(key))
	}
}

func (ptr *QGamepadKeyNavigation) ConnectSetRightKey(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setRightKey"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setRightKey", func(key core.Qt__Key) {
				signal.(func(core.Qt__Key))(key)
				f(key)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setRightKey", f)
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
		signal.(func(core.Qt__Key))(core.Qt__Key(key))
	} else {
		NewQGamepadKeyNavigationFromPointer(ptr).SetUpKeyDefault(core.Qt__Key(key))
	}
}

func (ptr *QGamepadKeyNavigation) ConnectSetUpKey(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setUpKey"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setUpKey", func(key core.Qt__Key) {
				signal.(func(core.Qt__Key))(key)
				f(key)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setUpKey", f)
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

//export callbackQGamepadKeyNavigation_UpKeyChanged
func callbackQGamepadKeyNavigation_UpKeyChanged(ptr unsafe.Pointer, key C.longlong) {
	if signal := qt.GetSignal(ptr, "upKeyChanged"); signal != nil {
		signal.(func(core.Qt__Key))(core.Qt__Key(key))
	}

}

func (ptr *QGamepadKeyNavigation) ConnectUpKeyChanged(f func(key core.Qt__Key)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "upKeyChanged") {
			C.QGamepadKeyNavigation_ConnectUpKeyChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "upKeyChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "upKeyChanged", func(key core.Qt__Key) {
				signal.(func(core.Qt__Key))(key)
				f(key)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "upKeyChanged", f)
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

func (ptr *QGamepadKeyNavigation) ButtonAKey() core.Qt__Key {
	if ptr.Pointer() != nil {
		return core.Qt__Key(C.QGamepadKeyNavigation_ButtonAKey(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGamepadKeyNavigation) ButtonBKey() core.Qt__Key {
	if ptr.Pointer() != nil {
		return core.Qt__Key(C.QGamepadKeyNavigation_ButtonBKey(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGamepadKeyNavigation) ButtonGuideKey() core.Qt__Key {
	if ptr.Pointer() != nil {
		return core.Qt__Key(C.QGamepadKeyNavigation_ButtonGuideKey(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGamepadKeyNavigation) ButtonL1Key() core.Qt__Key {
	if ptr.Pointer() != nil {
		return core.Qt__Key(C.QGamepadKeyNavigation_ButtonL1Key(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGamepadKeyNavigation) ButtonL2Key() core.Qt__Key {
	if ptr.Pointer() != nil {
		return core.Qt__Key(C.QGamepadKeyNavigation_ButtonL2Key(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGamepadKeyNavigation) ButtonL3Key() core.Qt__Key {
	if ptr.Pointer() != nil {
		return core.Qt__Key(C.QGamepadKeyNavigation_ButtonL3Key(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGamepadKeyNavigation) ButtonR1Key() core.Qt__Key {
	if ptr.Pointer() != nil {
		return core.Qt__Key(C.QGamepadKeyNavigation_ButtonR1Key(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGamepadKeyNavigation) ButtonR2Key() core.Qt__Key {
	if ptr.Pointer() != nil {
		return core.Qt__Key(C.QGamepadKeyNavigation_ButtonR2Key(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGamepadKeyNavigation) ButtonR3Key() core.Qt__Key {
	if ptr.Pointer() != nil {
		return core.Qt__Key(C.QGamepadKeyNavigation_ButtonR3Key(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGamepadKeyNavigation) ButtonSelectKey() core.Qt__Key {
	if ptr.Pointer() != nil {
		return core.Qt__Key(C.QGamepadKeyNavigation_ButtonSelectKey(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGamepadKeyNavigation) ButtonStartKey() core.Qt__Key {
	if ptr.Pointer() != nil {
		return core.Qt__Key(C.QGamepadKeyNavigation_ButtonStartKey(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGamepadKeyNavigation) ButtonXKey() core.Qt__Key {
	if ptr.Pointer() != nil {
		return core.Qt__Key(C.QGamepadKeyNavigation_ButtonXKey(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGamepadKeyNavigation) ButtonYKey() core.Qt__Key {
	if ptr.Pointer() != nil {
		return core.Qt__Key(C.QGamepadKeyNavigation_ButtonYKey(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGamepadKeyNavigation) DownKey() core.Qt__Key {
	if ptr.Pointer() != nil {
		return core.Qt__Key(C.QGamepadKeyNavigation_DownKey(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGamepadKeyNavigation) LeftKey() core.Qt__Key {
	if ptr.Pointer() != nil {
		return core.Qt__Key(C.QGamepadKeyNavigation_LeftKey(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGamepadKeyNavigation) RightKey() core.Qt__Key {
	if ptr.Pointer() != nil {
		return core.Qt__Key(C.QGamepadKeyNavigation_RightKey(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGamepadKeyNavigation) UpKey() core.Qt__Key {
	if ptr.Pointer() != nil {
		return core.Qt__Key(C.QGamepadKeyNavigation_UpKey(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGamepadKeyNavigation) Active() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepadKeyNavigation_Active(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQGamepadKeyNavigation_MetaObject
func callbackQGamepadKeyNavigation_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQGamepadKeyNavigationFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QGamepadKeyNavigation) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QGamepadKeyNavigation_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGamepadKeyNavigation) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QGamepadKeyNavigation___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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

func (ptr *QGamepadKeyNavigation) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGamepadKeyNavigation___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGamepadKeyNavigation) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGamepadKeyNavigation) __findChildren_newList2() unsafe.Pointer {
	return C.QGamepadKeyNavigation___findChildren_newList2(ptr.Pointer())
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

//export callbackQGamepadKeyNavigation_Event
func callbackQGamepadKeyNavigation_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
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
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGamepadKeyNavigationFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QGamepadKeyNavigation) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepadKeyNavigation_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQGamepadKeyNavigation_ChildEvent
func callbackQGamepadKeyNavigation_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
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
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
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
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
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
		signal.(func())()
	} else {
		NewQGamepadKeyNavigationFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QGamepadKeyNavigation) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQGamepadKeyNavigation_Destroyed
func callbackQGamepadKeyNavigation_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQGamepadKeyNavigation_DisconnectNotify
func callbackQGamepadKeyNavigation_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGamepadKeyNavigationFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGamepadKeyNavigation) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGamepadKeyNavigation_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGamepadKeyNavigation_ObjectNameChanged
func callbackQGamepadKeyNavigation_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtGamepad_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQGamepadKeyNavigation_TimerEvent
func callbackQGamepadKeyNavigation_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
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

//go:generate stringer -type=QGamepadManager__GamepadAxis
//QGamepadManager::GamepadAxis
type QGamepadManager__GamepadAxis int64

const (
	QGamepadManager__AxisInvalid QGamepadManager__GamepadAxis = QGamepadManager__GamepadAxis(-1)
	QGamepadManager__AxisLeftX   QGamepadManager__GamepadAxis = QGamepadManager__GamepadAxis(0)
	QGamepadManager__AxisLeftY   QGamepadManager__GamepadAxis = QGamepadManager__GamepadAxis(1)
	QGamepadManager__AxisRightX  QGamepadManager__GamepadAxis = QGamepadManager__GamepadAxis(2)
	QGamepadManager__AxisRightY  QGamepadManager__GamepadAxis = QGamepadManager__GamepadAxis(3)
)

//go:generate stringer -type=QGamepadManager__GamepadButton
//QGamepadManager::GamepadButton
type QGamepadManager__GamepadButton int64

const (
	QGamepadManager__ButtonInvalid QGamepadManager__GamepadButton = QGamepadManager__GamepadButton(-1)
	QGamepadManager__ButtonA       QGamepadManager__GamepadButton = QGamepadManager__GamepadButton(0)
	QGamepadManager__ButtonB       QGamepadManager__GamepadButton = QGamepadManager__GamepadButton(1)
	QGamepadManager__ButtonX       QGamepadManager__GamepadButton = QGamepadManager__GamepadButton(2)
	QGamepadManager__ButtonY       QGamepadManager__GamepadButton = QGamepadManager__GamepadButton(3)
	QGamepadManager__ButtonL1      QGamepadManager__GamepadButton = QGamepadManager__GamepadButton(4)
	QGamepadManager__ButtonR1      QGamepadManager__GamepadButton = QGamepadManager__GamepadButton(5)
	QGamepadManager__ButtonL2      QGamepadManager__GamepadButton = QGamepadManager__GamepadButton(6)
	QGamepadManager__ButtonR2      QGamepadManager__GamepadButton = QGamepadManager__GamepadButton(7)
	QGamepadManager__ButtonSelect  QGamepadManager__GamepadButton = QGamepadManager__GamepadButton(8)
	QGamepadManager__ButtonStart   QGamepadManager__GamepadButton = QGamepadManager__GamepadButton(9)
	QGamepadManager__ButtonL3      QGamepadManager__GamepadButton = QGamepadManager__GamepadButton(10)
	QGamepadManager__ButtonR3      QGamepadManager__GamepadButton = QGamepadManager__GamepadButton(11)
	QGamepadManager__ButtonUp      QGamepadManager__GamepadButton = QGamepadManager__GamepadButton(12)
	QGamepadManager__ButtonDown    QGamepadManager__GamepadButton = QGamepadManager__GamepadButton(13)
	QGamepadManager__ButtonRight   QGamepadManager__GamepadButton = QGamepadManager__GamepadButton(14)
	QGamepadManager__ButtonLeft    QGamepadManager__GamepadButton = QGamepadManager__GamepadButton(15)
	QGamepadManager__ButtonCenter  QGamepadManager__GamepadButton = QGamepadManager__GamepadButton(16)
	QGamepadManager__ButtonGuide   QGamepadManager__GamepadButton = QGamepadManager__GamepadButton(17)
)

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

func QGamepadManager_Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QGamepadManager_QGamepadManager_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QGamepadManager) Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QGamepadManager_QGamepadManager_Tr(sC, cC, C.int(int32(n))))
}

func QGamepadManager_TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QGamepadManager_QGamepadManager_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QGamepadManager) TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QGamepadManager_QGamepadManager_TrUtf8(sC, cC, C.int(int32(n))))
}

//export callbackQGamepadManager_ConfigureAxis
func callbackQGamepadManager_ConfigureAxis(ptr unsafe.Pointer, deviceId C.int, axis C.longlong) C.char {
	if signal := qt.GetSignal(ptr, "configureAxis"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, QGamepadManager__GamepadAxis) bool)(int(int32(deviceId)), QGamepadManager__GamepadAxis(axis)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGamepadManagerFromPointer(ptr).ConfigureAxisDefault(int(int32(deviceId)), QGamepadManager__GamepadAxis(axis)))))
}

func (ptr *QGamepadManager) ConnectConfigureAxis(f func(deviceId int, axis QGamepadManager__GamepadAxis) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "configureAxis"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "configureAxis", func(deviceId int, axis QGamepadManager__GamepadAxis) bool {
				signal.(func(int, QGamepadManager__GamepadAxis) bool)(deviceId, axis)
				return f(deviceId, axis)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "configureAxis", f)
		}
	}
}

func (ptr *QGamepadManager) DisconnectConfigureAxis() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "configureAxis")
	}
}

func (ptr *QGamepadManager) ConfigureAxis(deviceId int, axis QGamepadManager__GamepadAxis) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepadManager_ConfigureAxis(ptr.Pointer(), C.int(int32(deviceId)), C.longlong(axis))) != 0
	}
	return false
}

func (ptr *QGamepadManager) ConfigureAxisDefault(deviceId int, axis QGamepadManager__GamepadAxis) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepadManager_ConfigureAxisDefault(ptr.Pointer(), C.int(int32(deviceId)), C.longlong(axis))) != 0
	}
	return false
}

//export callbackQGamepadManager_ConfigureButton
func callbackQGamepadManager_ConfigureButton(ptr unsafe.Pointer, deviceId C.int, button C.longlong) C.char {
	if signal := qt.GetSignal(ptr, "configureButton"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, QGamepadManager__GamepadButton) bool)(int(int32(deviceId)), QGamepadManager__GamepadButton(button)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGamepadManagerFromPointer(ptr).ConfigureButtonDefault(int(int32(deviceId)), QGamepadManager__GamepadButton(button)))))
}

func (ptr *QGamepadManager) ConnectConfigureButton(f func(deviceId int, button QGamepadManager__GamepadButton) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "configureButton"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "configureButton", func(deviceId int, button QGamepadManager__GamepadButton) bool {
				signal.(func(int, QGamepadManager__GamepadButton) bool)(deviceId, button)
				return f(deviceId, button)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "configureButton", f)
		}
	}
}

func (ptr *QGamepadManager) DisconnectConfigureButton() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "configureButton")
	}
}

func (ptr *QGamepadManager) ConfigureButton(deviceId int, button QGamepadManager__GamepadButton) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepadManager_ConfigureButton(ptr.Pointer(), C.int(int32(deviceId)), C.longlong(button))) != 0
	}
	return false
}

func (ptr *QGamepadManager) ConfigureButtonDefault(deviceId int, button QGamepadManager__GamepadButton) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepadManager_ConfigureButtonDefault(ptr.Pointer(), C.int(int32(deviceId)), C.longlong(button))) != 0
	}
	return false
}

//export callbackQGamepadManager_SetCancelConfigureButton
func callbackQGamepadManager_SetCancelConfigureButton(ptr unsafe.Pointer, deviceId C.int, button C.longlong) C.char {
	if signal := qt.GetSignal(ptr, "setCancelConfigureButton"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, QGamepadManager__GamepadButton) bool)(int(int32(deviceId)), QGamepadManager__GamepadButton(button)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGamepadManagerFromPointer(ptr).SetCancelConfigureButtonDefault(int(int32(deviceId)), QGamepadManager__GamepadButton(button)))))
}

func (ptr *QGamepadManager) ConnectSetCancelConfigureButton(f func(deviceId int, button QGamepadManager__GamepadButton) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setCancelConfigureButton"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setCancelConfigureButton", func(deviceId int, button QGamepadManager__GamepadButton) bool {
				signal.(func(int, QGamepadManager__GamepadButton) bool)(deviceId, button)
				return f(deviceId, button)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setCancelConfigureButton", f)
		}
	}
}

func (ptr *QGamepadManager) DisconnectSetCancelConfigureButton() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setCancelConfigureButton")
	}
}

func (ptr *QGamepadManager) SetCancelConfigureButton(deviceId int, button QGamepadManager__GamepadButton) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepadManager_SetCancelConfigureButton(ptr.Pointer(), C.int(int32(deviceId)), C.longlong(button))) != 0
	}
	return false
}

func (ptr *QGamepadManager) SetCancelConfigureButtonDefault(deviceId int, button QGamepadManager__GamepadButton) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepadManager_SetCancelConfigureButtonDefault(ptr.Pointer(), C.int(int32(deviceId)), C.longlong(button))) != 0
	}
	return false
}

//export callbackQGamepadManager_AxisConfigured
func callbackQGamepadManager_AxisConfigured(ptr unsafe.Pointer, deviceId C.int, axis C.longlong) {
	if signal := qt.GetSignal(ptr, "axisConfigured"); signal != nil {
		signal.(func(int, QGamepadManager__GamepadAxis))(int(int32(deviceId)), QGamepadManager__GamepadAxis(axis))
	}

}

func (ptr *QGamepadManager) ConnectAxisConfigured(f func(deviceId int, axis QGamepadManager__GamepadAxis)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "axisConfigured") {
			C.QGamepadManager_ConnectAxisConfigured(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "axisConfigured"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "axisConfigured", func(deviceId int, axis QGamepadManager__GamepadAxis) {
				signal.(func(int, QGamepadManager__GamepadAxis))(deviceId, axis)
				f(deviceId, axis)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "axisConfigured", f)
		}
	}
}

func (ptr *QGamepadManager) DisconnectAxisConfigured() {
	if ptr.Pointer() != nil {
		C.QGamepadManager_DisconnectAxisConfigured(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "axisConfigured")
	}
}

func (ptr *QGamepadManager) AxisConfigured(deviceId int, axis QGamepadManager__GamepadAxis) {
	if ptr.Pointer() != nil {
		C.QGamepadManager_AxisConfigured(ptr.Pointer(), C.int(int32(deviceId)), C.longlong(axis))
	}
}

//export callbackQGamepadManager_ButtonConfigured
func callbackQGamepadManager_ButtonConfigured(ptr unsafe.Pointer, deviceId C.int, button C.longlong) {
	if signal := qt.GetSignal(ptr, "buttonConfigured"); signal != nil {
		signal.(func(int, QGamepadManager__GamepadButton))(int(int32(deviceId)), QGamepadManager__GamepadButton(button))
	}

}

func (ptr *QGamepadManager) ConnectButtonConfigured(f func(deviceId int, button QGamepadManager__GamepadButton)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "buttonConfigured") {
			C.QGamepadManager_ConnectButtonConfigured(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "buttonConfigured"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "buttonConfigured", func(deviceId int, button QGamepadManager__GamepadButton) {
				signal.(func(int, QGamepadManager__GamepadButton))(deviceId, button)
				f(deviceId, button)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "buttonConfigured", f)
		}
	}
}

func (ptr *QGamepadManager) DisconnectButtonConfigured() {
	if ptr.Pointer() != nil {
		C.QGamepadManager_DisconnectButtonConfigured(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "buttonConfigured")
	}
}

func (ptr *QGamepadManager) ButtonConfigured(deviceId int, button QGamepadManager__GamepadButton) {
	if ptr.Pointer() != nil {
		C.QGamepadManager_ButtonConfigured(ptr.Pointer(), C.int(int32(deviceId)), C.longlong(button))
	}
}

//export callbackQGamepadManager_ConfigurationCanceled
func callbackQGamepadManager_ConfigurationCanceled(ptr unsafe.Pointer, deviceId C.int) {
	if signal := qt.GetSignal(ptr, "configurationCanceled"); signal != nil {
		signal.(func(int))(int(int32(deviceId)))
	}

}

func (ptr *QGamepadManager) ConnectConfigurationCanceled(f func(deviceId int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "configurationCanceled") {
			C.QGamepadManager_ConnectConfigurationCanceled(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "configurationCanceled"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "configurationCanceled", func(deviceId int) {
				signal.(func(int))(deviceId)
				f(deviceId)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "configurationCanceled", f)
		}
	}
}

func (ptr *QGamepadManager) DisconnectConfigurationCanceled() {
	if ptr.Pointer() != nil {
		C.QGamepadManager_DisconnectConfigurationCanceled(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "configurationCanceled")
	}
}

func (ptr *QGamepadManager) ConfigurationCanceled(deviceId int) {
	if ptr.Pointer() != nil {
		C.QGamepadManager_ConfigurationCanceled(ptr.Pointer(), C.int(int32(deviceId)))
	}
}

//export callbackQGamepadManager_ConnectedGamepadsChanged
func callbackQGamepadManager_ConnectedGamepadsChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectedGamepadsChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QGamepadManager) ConnectConnectedGamepadsChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "connectedGamepadsChanged") {
			C.QGamepadManager_ConnectConnectedGamepadsChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "connectedGamepadsChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "connectedGamepadsChanged", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "connectedGamepadsChanged", f)
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

//export callbackQGamepadManager_GamepadAxisEvent
func callbackQGamepadManager_GamepadAxisEvent(ptr unsafe.Pointer, deviceId C.int, axis C.longlong, value C.double) {
	if signal := qt.GetSignal(ptr, "gamepadAxisEvent"); signal != nil {
		signal.(func(int, QGamepadManager__GamepadAxis, float64))(int(int32(deviceId)), QGamepadManager__GamepadAxis(axis), float64(value))
	}

}

func (ptr *QGamepadManager) ConnectGamepadAxisEvent(f func(deviceId int, axis QGamepadManager__GamepadAxis, value float64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "gamepadAxisEvent") {
			C.QGamepadManager_ConnectGamepadAxisEvent(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "gamepadAxisEvent"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "gamepadAxisEvent", func(deviceId int, axis QGamepadManager__GamepadAxis, value float64) {
				signal.(func(int, QGamepadManager__GamepadAxis, float64))(deviceId, axis, value)
				f(deviceId, axis, value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "gamepadAxisEvent", f)
		}
	}
}

func (ptr *QGamepadManager) DisconnectGamepadAxisEvent() {
	if ptr.Pointer() != nil {
		C.QGamepadManager_DisconnectGamepadAxisEvent(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "gamepadAxisEvent")
	}
}

func (ptr *QGamepadManager) GamepadAxisEvent(deviceId int, axis QGamepadManager__GamepadAxis, value float64) {
	if ptr.Pointer() != nil {
		C.QGamepadManager_GamepadAxisEvent(ptr.Pointer(), C.int(int32(deviceId)), C.longlong(axis), C.double(value))
	}
}

//export callbackQGamepadManager_GamepadButtonPressEvent
func callbackQGamepadManager_GamepadButtonPressEvent(ptr unsafe.Pointer, deviceId C.int, button C.longlong, value C.double) {
	if signal := qt.GetSignal(ptr, "gamepadButtonPressEvent"); signal != nil {
		signal.(func(int, QGamepadManager__GamepadButton, float64))(int(int32(deviceId)), QGamepadManager__GamepadButton(button), float64(value))
	}

}

func (ptr *QGamepadManager) ConnectGamepadButtonPressEvent(f func(deviceId int, button QGamepadManager__GamepadButton, value float64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "gamepadButtonPressEvent") {
			C.QGamepadManager_ConnectGamepadButtonPressEvent(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "gamepadButtonPressEvent"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "gamepadButtonPressEvent", func(deviceId int, button QGamepadManager__GamepadButton, value float64) {
				signal.(func(int, QGamepadManager__GamepadButton, float64))(deviceId, button, value)
				f(deviceId, button, value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "gamepadButtonPressEvent", f)
		}
	}
}

func (ptr *QGamepadManager) DisconnectGamepadButtonPressEvent() {
	if ptr.Pointer() != nil {
		C.QGamepadManager_DisconnectGamepadButtonPressEvent(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "gamepadButtonPressEvent")
	}
}

func (ptr *QGamepadManager) GamepadButtonPressEvent(deviceId int, button QGamepadManager__GamepadButton, value float64) {
	if ptr.Pointer() != nil {
		C.QGamepadManager_GamepadButtonPressEvent(ptr.Pointer(), C.int(int32(deviceId)), C.longlong(button), C.double(value))
	}
}

//export callbackQGamepadManager_GamepadButtonReleaseEvent
func callbackQGamepadManager_GamepadButtonReleaseEvent(ptr unsafe.Pointer, deviceId C.int, button C.longlong) {
	if signal := qt.GetSignal(ptr, "gamepadButtonReleaseEvent"); signal != nil {
		signal.(func(int, QGamepadManager__GamepadButton))(int(int32(deviceId)), QGamepadManager__GamepadButton(button))
	}

}

func (ptr *QGamepadManager) ConnectGamepadButtonReleaseEvent(f func(deviceId int, button QGamepadManager__GamepadButton)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "gamepadButtonReleaseEvent") {
			C.QGamepadManager_ConnectGamepadButtonReleaseEvent(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "gamepadButtonReleaseEvent"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "gamepadButtonReleaseEvent", func(deviceId int, button QGamepadManager__GamepadButton) {
				signal.(func(int, QGamepadManager__GamepadButton))(deviceId, button)
				f(deviceId, button)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "gamepadButtonReleaseEvent", f)
		}
	}
}

func (ptr *QGamepadManager) DisconnectGamepadButtonReleaseEvent() {
	if ptr.Pointer() != nil {
		C.QGamepadManager_DisconnectGamepadButtonReleaseEvent(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "gamepadButtonReleaseEvent")
	}
}

func (ptr *QGamepadManager) GamepadButtonReleaseEvent(deviceId int, button QGamepadManager__GamepadButton) {
	if ptr.Pointer() != nil {
		C.QGamepadManager_GamepadButtonReleaseEvent(ptr.Pointer(), C.int(int32(deviceId)), C.longlong(button))
	}
}

//export callbackQGamepadManager_GamepadConnected
func callbackQGamepadManager_GamepadConnected(ptr unsafe.Pointer, deviceId C.int) {
	if signal := qt.GetSignal(ptr, "gamepadConnected"); signal != nil {
		signal.(func(int))(int(int32(deviceId)))
	}

}

func (ptr *QGamepadManager) ConnectGamepadConnected(f func(deviceId int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "gamepadConnected") {
			C.QGamepadManager_ConnectGamepadConnected(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "gamepadConnected"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "gamepadConnected", func(deviceId int) {
				signal.(func(int))(deviceId)
				f(deviceId)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "gamepadConnected", f)
		}
	}
}

func (ptr *QGamepadManager) DisconnectGamepadConnected() {
	if ptr.Pointer() != nil {
		C.QGamepadManager_DisconnectGamepadConnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "gamepadConnected")
	}
}

func (ptr *QGamepadManager) GamepadConnected(deviceId int) {
	if ptr.Pointer() != nil {
		C.QGamepadManager_GamepadConnected(ptr.Pointer(), C.int(int32(deviceId)))
	}
}

//export callbackQGamepadManager_GamepadDisconnected
func callbackQGamepadManager_GamepadDisconnected(ptr unsafe.Pointer, deviceId C.int) {
	if signal := qt.GetSignal(ptr, "gamepadDisconnected"); signal != nil {
		signal.(func(int))(int(int32(deviceId)))
	}

}

func (ptr *QGamepadManager) ConnectGamepadDisconnected(f func(deviceId int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "gamepadDisconnected") {
			C.QGamepadManager_ConnectGamepadDisconnected(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "gamepadDisconnected"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "gamepadDisconnected", func(deviceId int) {
				signal.(func(int))(deviceId)
				f(deviceId)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "gamepadDisconnected", f)
		}
	}
}

func (ptr *QGamepadManager) DisconnectGamepadDisconnected() {
	if ptr.Pointer() != nil {
		C.QGamepadManager_DisconnectGamepadDisconnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "gamepadDisconnected")
	}
}

func (ptr *QGamepadManager) GamepadDisconnected(deviceId int) {
	if ptr.Pointer() != nil {
		C.QGamepadManager_GamepadDisconnected(ptr.Pointer(), C.int(int32(deviceId)))
	}
}

//export callbackQGamepadManager_GamepadNameChanged
func callbackQGamepadManager_GamepadNameChanged(ptr unsafe.Pointer, deviceId C.int, name C.struct_QtGamepad_PackedString) {
	if signal := qt.GetSignal(ptr, "gamepadNameChanged"); signal != nil {
		signal.(func(int, string))(int(int32(deviceId)), cGoUnpackString(name))
	}

}

func (ptr *QGamepadManager) ConnectGamepadNameChanged(f func(deviceId int, name string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "gamepadNameChanged") {
			C.QGamepadManager_ConnectGamepadNameChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "gamepadNameChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "gamepadNameChanged", func(deviceId int, name string) {
				signal.(func(int, string))(deviceId, name)
				f(deviceId, name)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "gamepadNameChanged", f)
		}
	}
}

func (ptr *QGamepadManager) DisconnectGamepadNameChanged() {
	if ptr.Pointer() != nil {
		C.QGamepadManager_DisconnectGamepadNameChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "gamepadNameChanged")
	}
}

func (ptr *QGamepadManager) GamepadNameChanged(deviceId int, name string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QGamepadManager_GamepadNameChanged(ptr.Pointer(), C.int(int32(deviceId)), C.struct_QtGamepad_PackedString{data: nameC, len: C.longlong(len(name))})
	}
}

//export callbackQGamepadManager_ResetConfiguration
func callbackQGamepadManager_ResetConfiguration(ptr unsafe.Pointer, deviceId C.int) {
	if signal := qt.GetSignal(ptr, "resetConfiguration"); signal != nil {
		signal.(func(int))(int(int32(deviceId)))
	} else {
		NewQGamepadManagerFromPointer(ptr).ResetConfigurationDefault(int(int32(deviceId)))
	}
}

func (ptr *QGamepadManager) ConnectResetConfiguration(f func(deviceId int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "resetConfiguration"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "resetConfiguration", func(deviceId int) {
				signal.(func(int))(deviceId)
				f(deviceId)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "resetConfiguration", f)
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
		signal.(func(string))(cGoUnpackString(file))
	} else {
		NewQGamepadManagerFromPointer(ptr).SetSettingsFileDefault(cGoUnpackString(file))
	}
}

func (ptr *QGamepadManager) ConnectSetSettingsFile(f func(file string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setSettingsFile"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setSettingsFile", func(file string) {
				signal.(func(string))(file)
				f(file)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setSettingsFile", f)
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

func (ptr *QGamepadManager) GamepadName(deviceId int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGamepadManager_GamepadName(ptr.Pointer(), C.int(int32(deviceId))))
	}
	return ""
}

//export callbackQGamepadManager_IsConfigurationNeeded
func callbackQGamepadManager_IsConfigurationNeeded(ptr unsafe.Pointer, deviceId C.int) C.char {
	if signal := qt.GetSignal(ptr, "isConfigurationNeeded"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(deviceId))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGamepadManagerFromPointer(ptr).IsConfigurationNeededDefault(int(int32(deviceId))))))
}

func (ptr *QGamepadManager) ConnectIsConfigurationNeeded(f func(deviceId int) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "isConfigurationNeeded"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "isConfigurationNeeded", func(deviceId int) bool {
				signal.(func(int) bool)(deviceId)
				return f(deviceId)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "isConfigurationNeeded", f)
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

//export callbackQGamepadManager_MetaObject
func callbackQGamepadManager_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQGamepadManagerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QGamepadManager) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QGamepadManager_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
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

func (ptr *QGamepadManager) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QGamepadManager___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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

func (ptr *QGamepadManager) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGamepadManager___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGamepadManager) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGamepadManager___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGamepadManager) __findChildren_newList2() unsafe.Pointer {
	return C.QGamepadManager___findChildren_newList2(ptr.Pointer())
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

//export callbackQGamepadManager_Event
func callbackQGamepadManager_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
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
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGamepadManagerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QGamepadManager) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGamepadManager_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQGamepadManager_ChildEvent
func callbackQGamepadManager_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
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
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
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
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
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
		signal.(func())()
	} else {
		NewQGamepadManagerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QGamepadManager) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QGamepadManager_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQGamepadManager_Destroyed
func callbackQGamepadManager_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQGamepadManager_DisconnectNotify
func callbackQGamepadManager_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGamepadManagerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGamepadManager) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGamepadManager_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGamepadManager_ObjectNameChanged
func callbackQGamepadManager_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtGamepad_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQGamepadManager_TimerEvent
func callbackQGamepadManager_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGamepadManagerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGamepadManager) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGamepadManager_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}
