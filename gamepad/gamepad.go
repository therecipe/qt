// +build !minimal

package gamepad

//#include <stdint.h>
//#include <stdlib.h>
//#include "gamepad.h"
import "C"
import (
	"fmt"
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QGamepad struct {
	core.QObject
}

type QGamepad_ITF interface {
	core.QObject_ITF
	QGamepad_PTR() *QGamepad
}

func (p *QGamepad) QGamepad_PTR() *QGamepad {
	return p
}

func (p *QGamepad) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QGamepad) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
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
func (ptr *QGamepad) AxisLeftX() float64 {
	defer qt.Recovering("QGamepad::axisLeftX")

	if ptr.Pointer() != nil {
		return float64(C.QGamepad_AxisLeftX(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGamepad) AxisLeftY() float64 {
	defer qt.Recovering("QGamepad::axisLeftY")

	if ptr.Pointer() != nil {
		return float64(C.QGamepad_AxisLeftY(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGamepad) AxisRightX() float64 {
	defer qt.Recovering("QGamepad::axisRightX")

	if ptr.Pointer() != nil {
		return float64(C.QGamepad_AxisRightX(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGamepad) AxisRightY() float64 {
	defer qt.Recovering("QGamepad::axisRightY")

	if ptr.Pointer() != nil {
		return float64(C.QGamepad_AxisRightY(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGamepad) ButtonA() bool {
	defer qt.Recovering("QGamepad::buttonA")

	if ptr.Pointer() != nil {
		return C.QGamepad_ButtonA(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGamepad) ButtonB() bool {
	defer qt.Recovering("QGamepad::buttonB")

	if ptr.Pointer() != nil {
		return C.QGamepad_ButtonB(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGamepad) ButtonCenter() bool {
	defer qt.Recovering("QGamepad::buttonCenter")

	if ptr.Pointer() != nil {
		return C.QGamepad_ButtonCenter(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGamepad) ButtonDown() bool {
	defer qt.Recovering("QGamepad::buttonDown")

	if ptr.Pointer() != nil {
		return C.QGamepad_ButtonDown(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGamepad) ButtonGuide() bool {
	defer qt.Recovering("QGamepad::buttonGuide")

	if ptr.Pointer() != nil {
		return C.QGamepad_ButtonGuide(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGamepad) ButtonL1() bool {
	defer qt.Recovering("QGamepad::buttonL1")

	if ptr.Pointer() != nil {
		return C.QGamepad_ButtonL1(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGamepad) ButtonL2() float64 {
	defer qt.Recovering("QGamepad::buttonL2")

	if ptr.Pointer() != nil {
		return float64(C.QGamepad_ButtonL2(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGamepad) ButtonL3() bool {
	defer qt.Recovering("QGamepad::buttonL3")

	if ptr.Pointer() != nil {
		return C.QGamepad_ButtonL3(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGamepad) ButtonLeft() bool {
	defer qt.Recovering("QGamepad::buttonLeft")

	if ptr.Pointer() != nil {
		return C.QGamepad_ButtonLeft(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGamepad) ButtonR1() bool {
	defer qt.Recovering("QGamepad::buttonR1")

	if ptr.Pointer() != nil {
		return C.QGamepad_ButtonR1(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGamepad) ButtonR2() float64 {
	defer qt.Recovering("QGamepad::buttonR2")

	if ptr.Pointer() != nil {
		return float64(C.QGamepad_ButtonR2(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGamepad) ButtonR3() bool {
	defer qt.Recovering("QGamepad::buttonR3")

	if ptr.Pointer() != nil {
		return C.QGamepad_ButtonR3(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGamepad) ButtonRight() bool {
	defer qt.Recovering("QGamepad::buttonRight")

	if ptr.Pointer() != nil {
		return C.QGamepad_ButtonRight(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGamepad) ButtonSelect() bool {
	defer qt.Recovering("QGamepad::buttonSelect")

	if ptr.Pointer() != nil {
		return C.QGamepad_ButtonSelect(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGamepad) ButtonStart() bool {
	defer qt.Recovering("QGamepad::buttonStart")

	if ptr.Pointer() != nil {
		return C.QGamepad_ButtonStart(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGamepad) ButtonUp() bool {
	defer qt.Recovering("QGamepad::buttonUp")

	if ptr.Pointer() != nil {
		return C.QGamepad_ButtonUp(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGamepad) ButtonX() bool {
	defer qt.Recovering("QGamepad::buttonX")

	if ptr.Pointer() != nil {
		return C.QGamepad_ButtonX(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGamepad) ButtonY() bool {
	defer qt.Recovering("QGamepad::buttonY")

	if ptr.Pointer() != nil {
		return C.QGamepad_ButtonY(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGamepad) DeviceId() int {
	defer qt.Recovering("QGamepad::deviceId")

	if ptr.Pointer() != nil {
		return int(int32(C.QGamepad_DeviceId(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGamepad) IsConnected() bool {
	defer qt.Recovering("QGamepad::isConnected")

	if ptr.Pointer() != nil {
		return C.QGamepad_IsConnected(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGamepad) Name() string {
	defer qt.Recovering("QGamepad::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGamepad_Name(ptr.Pointer()))
	}
	return ""
}

//export callbackQGamepad_SetDeviceId
func callbackQGamepad_SetDeviceId(ptr unsafe.Pointer, number C.int) {
	defer qt.Recovering("callback QGamepad::setDeviceId")

	if signal := qt.GetSignal(fmt.Sprintf("QGamepad(%v)", ptr), "setDeviceId"); signal != nil {
		signal.(func(int))(int(int32(number)))
	}

}

func (ptr *QGamepad) ConnectSetDeviceId(f func(number int)) {
	defer qt.Recovering("connect QGamepad::setDeviceId")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "setDeviceId", f)
	}
}

func (ptr *QGamepad) DisconnectSetDeviceId(number int) {
	defer qt.Recovering("disconnect QGamepad::setDeviceId")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "setDeviceId")
	}
}

func (ptr *QGamepad) SetDeviceId(number int) {
	defer qt.Recovering("QGamepad::setDeviceId")

	if ptr.Pointer() != nil {
		C.QGamepad_SetDeviceId(ptr.Pointer(), C.int(int32(number)))
	}
}

func NewQGamepad(deviceId int, parent core.QObject_ITF) *QGamepad {
	defer qt.Recovering("QGamepad::QGamepad")

	return NewQGamepadFromPointer(C.QGamepad_NewQGamepad(C.int(int32(deviceId)), core.PointerFromQObject(parent)))
}

//export callbackQGamepad_AxisLeftXChanged
func callbackQGamepad_AxisLeftXChanged(ptr unsafe.Pointer, value C.double) {
	defer qt.Recovering("callback QGamepad::axisLeftXChanged")

	if signal := qt.GetSignal(fmt.Sprintf("QGamepad(%v)", ptr), "axisLeftXChanged"); signal != nil {
		signal.(func(float64))(float64(value))
	}

}

func (ptr *QGamepad) ConnectAxisLeftXChanged(f func(value float64)) {
	defer qt.Recovering("connect QGamepad::axisLeftXChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_ConnectAxisLeftXChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "axisLeftXChanged", f)
	}
}

func (ptr *QGamepad) DisconnectAxisLeftXChanged() {
	defer qt.Recovering("disconnect QGamepad::axisLeftXChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectAxisLeftXChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "axisLeftXChanged")
	}
}

func (ptr *QGamepad) AxisLeftXChanged(value float64) {
	defer qt.Recovering("QGamepad::axisLeftXChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_AxisLeftXChanged(ptr.Pointer(), C.double(value))
	}
}

//export callbackQGamepad_AxisLeftYChanged
func callbackQGamepad_AxisLeftYChanged(ptr unsafe.Pointer, value C.double) {
	defer qt.Recovering("callback QGamepad::axisLeftYChanged")

	if signal := qt.GetSignal(fmt.Sprintf("QGamepad(%v)", ptr), "axisLeftYChanged"); signal != nil {
		signal.(func(float64))(float64(value))
	}

}

func (ptr *QGamepad) ConnectAxisLeftYChanged(f func(value float64)) {
	defer qt.Recovering("connect QGamepad::axisLeftYChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_ConnectAxisLeftYChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "axisLeftYChanged", f)
	}
}

func (ptr *QGamepad) DisconnectAxisLeftYChanged() {
	defer qt.Recovering("disconnect QGamepad::axisLeftYChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectAxisLeftYChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "axisLeftYChanged")
	}
}

func (ptr *QGamepad) AxisLeftYChanged(value float64) {
	defer qt.Recovering("QGamepad::axisLeftYChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_AxisLeftYChanged(ptr.Pointer(), C.double(value))
	}
}

//export callbackQGamepad_AxisRightXChanged
func callbackQGamepad_AxisRightXChanged(ptr unsafe.Pointer, value C.double) {
	defer qt.Recovering("callback QGamepad::axisRightXChanged")

	if signal := qt.GetSignal(fmt.Sprintf("QGamepad(%v)", ptr), "axisRightXChanged"); signal != nil {
		signal.(func(float64))(float64(value))
	}

}

func (ptr *QGamepad) ConnectAxisRightXChanged(f func(value float64)) {
	defer qt.Recovering("connect QGamepad::axisRightXChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_ConnectAxisRightXChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "axisRightXChanged", f)
	}
}

func (ptr *QGamepad) DisconnectAxisRightXChanged() {
	defer qt.Recovering("disconnect QGamepad::axisRightXChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectAxisRightXChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "axisRightXChanged")
	}
}

func (ptr *QGamepad) AxisRightXChanged(value float64) {
	defer qt.Recovering("QGamepad::axisRightXChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_AxisRightXChanged(ptr.Pointer(), C.double(value))
	}
}

//export callbackQGamepad_AxisRightYChanged
func callbackQGamepad_AxisRightYChanged(ptr unsafe.Pointer, value C.double) {
	defer qt.Recovering("callback QGamepad::axisRightYChanged")

	if signal := qt.GetSignal(fmt.Sprintf("QGamepad(%v)", ptr), "axisRightYChanged"); signal != nil {
		signal.(func(float64))(float64(value))
	}

}

func (ptr *QGamepad) ConnectAxisRightYChanged(f func(value float64)) {
	defer qt.Recovering("connect QGamepad::axisRightYChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_ConnectAxisRightYChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "axisRightYChanged", f)
	}
}

func (ptr *QGamepad) DisconnectAxisRightYChanged() {
	defer qt.Recovering("disconnect QGamepad::axisRightYChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectAxisRightYChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "axisRightYChanged")
	}
}

func (ptr *QGamepad) AxisRightYChanged(value float64) {
	defer qt.Recovering("QGamepad::axisRightYChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_AxisRightYChanged(ptr.Pointer(), C.double(value))
	}
}

//export callbackQGamepad_ButtonAChanged
func callbackQGamepad_ButtonAChanged(ptr unsafe.Pointer, value C.char) {
	defer qt.Recovering("callback QGamepad::buttonAChanged")

	if signal := qt.GetSignal(fmt.Sprintf("QGamepad(%v)", ptr), "buttonAChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonAChanged(f func(value bool)) {
	defer qt.Recovering("connect QGamepad::buttonAChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_ConnectButtonAChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "buttonAChanged", f)
	}
}

func (ptr *QGamepad) DisconnectButtonAChanged() {
	defer qt.Recovering("disconnect QGamepad::buttonAChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonAChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "buttonAChanged")
	}
}

func (ptr *QGamepad) ButtonAChanged(value bool) {
	defer qt.Recovering("QGamepad::buttonAChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_ButtonAChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_ButtonBChanged
func callbackQGamepad_ButtonBChanged(ptr unsafe.Pointer, value C.char) {
	defer qt.Recovering("callback QGamepad::buttonBChanged")

	if signal := qt.GetSignal(fmt.Sprintf("QGamepad(%v)", ptr), "buttonBChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonBChanged(f func(value bool)) {
	defer qt.Recovering("connect QGamepad::buttonBChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_ConnectButtonBChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "buttonBChanged", f)
	}
}

func (ptr *QGamepad) DisconnectButtonBChanged() {
	defer qt.Recovering("disconnect QGamepad::buttonBChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonBChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "buttonBChanged")
	}
}

func (ptr *QGamepad) ButtonBChanged(value bool) {
	defer qt.Recovering("QGamepad::buttonBChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_ButtonBChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_ButtonCenterChanged
func callbackQGamepad_ButtonCenterChanged(ptr unsafe.Pointer, value C.char) {
	defer qt.Recovering("callback QGamepad::buttonCenterChanged")

	if signal := qt.GetSignal(fmt.Sprintf("QGamepad(%v)", ptr), "buttonCenterChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonCenterChanged(f func(value bool)) {
	defer qt.Recovering("connect QGamepad::buttonCenterChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_ConnectButtonCenterChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "buttonCenterChanged", f)
	}
}

func (ptr *QGamepad) DisconnectButtonCenterChanged() {
	defer qt.Recovering("disconnect QGamepad::buttonCenterChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonCenterChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "buttonCenterChanged")
	}
}

func (ptr *QGamepad) ButtonCenterChanged(value bool) {
	defer qt.Recovering("QGamepad::buttonCenterChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_ButtonCenterChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_ButtonDownChanged
func callbackQGamepad_ButtonDownChanged(ptr unsafe.Pointer, value C.char) {
	defer qt.Recovering("callback QGamepad::buttonDownChanged")

	if signal := qt.GetSignal(fmt.Sprintf("QGamepad(%v)", ptr), "buttonDownChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonDownChanged(f func(value bool)) {
	defer qt.Recovering("connect QGamepad::buttonDownChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_ConnectButtonDownChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "buttonDownChanged", f)
	}
}

func (ptr *QGamepad) DisconnectButtonDownChanged() {
	defer qt.Recovering("disconnect QGamepad::buttonDownChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonDownChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "buttonDownChanged")
	}
}

func (ptr *QGamepad) ButtonDownChanged(value bool) {
	defer qt.Recovering("QGamepad::buttonDownChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_ButtonDownChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_ButtonGuideChanged
func callbackQGamepad_ButtonGuideChanged(ptr unsafe.Pointer, value C.char) {
	defer qt.Recovering("callback QGamepad::buttonGuideChanged")

	if signal := qt.GetSignal(fmt.Sprintf("QGamepad(%v)", ptr), "buttonGuideChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonGuideChanged(f func(value bool)) {
	defer qt.Recovering("connect QGamepad::buttonGuideChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_ConnectButtonGuideChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "buttonGuideChanged", f)
	}
}

func (ptr *QGamepad) DisconnectButtonGuideChanged() {
	defer qt.Recovering("disconnect QGamepad::buttonGuideChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonGuideChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "buttonGuideChanged")
	}
}

func (ptr *QGamepad) ButtonGuideChanged(value bool) {
	defer qt.Recovering("QGamepad::buttonGuideChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_ButtonGuideChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_ButtonL1Changed
func callbackQGamepad_ButtonL1Changed(ptr unsafe.Pointer, value C.char) {
	defer qt.Recovering("callback QGamepad::buttonL1Changed")

	if signal := qt.GetSignal(fmt.Sprintf("QGamepad(%v)", ptr), "buttonL1Changed"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonL1Changed(f func(value bool)) {
	defer qt.Recovering("connect QGamepad::buttonL1Changed")

	if ptr.Pointer() != nil {
		C.QGamepad_ConnectButtonL1Changed(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "buttonL1Changed", f)
	}
}

func (ptr *QGamepad) DisconnectButtonL1Changed() {
	defer qt.Recovering("disconnect QGamepad::buttonL1Changed")

	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonL1Changed(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "buttonL1Changed")
	}
}

func (ptr *QGamepad) ButtonL1Changed(value bool) {
	defer qt.Recovering("QGamepad::buttonL1Changed")

	if ptr.Pointer() != nil {
		C.QGamepad_ButtonL1Changed(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_ButtonL2Changed
func callbackQGamepad_ButtonL2Changed(ptr unsafe.Pointer, value C.double) {
	defer qt.Recovering("callback QGamepad::buttonL2Changed")

	if signal := qt.GetSignal(fmt.Sprintf("QGamepad(%v)", ptr), "buttonL2Changed"); signal != nil {
		signal.(func(float64))(float64(value))
	}

}

func (ptr *QGamepad) ConnectButtonL2Changed(f func(value float64)) {
	defer qt.Recovering("connect QGamepad::buttonL2Changed")

	if ptr.Pointer() != nil {
		C.QGamepad_ConnectButtonL2Changed(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "buttonL2Changed", f)
	}
}

func (ptr *QGamepad) DisconnectButtonL2Changed() {
	defer qt.Recovering("disconnect QGamepad::buttonL2Changed")

	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonL2Changed(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "buttonL2Changed")
	}
}

func (ptr *QGamepad) ButtonL2Changed(value float64) {
	defer qt.Recovering("QGamepad::buttonL2Changed")

	if ptr.Pointer() != nil {
		C.QGamepad_ButtonL2Changed(ptr.Pointer(), C.double(value))
	}
}

//export callbackQGamepad_ButtonL3Changed
func callbackQGamepad_ButtonL3Changed(ptr unsafe.Pointer, value C.char) {
	defer qt.Recovering("callback QGamepad::buttonL3Changed")

	if signal := qt.GetSignal(fmt.Sprintf("QGamepad(%v)", ptr), "buttonL3Changed"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonL3Changed(f func(value bool)) {
	defer qt.Recovering("connect QGamepad::buttonL3Changed")

	if ptr.Pointer() != nil {
		C.QGamepad_ConnectButtonL3Changed(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "buttonL3Changed", f)
	}
}

func (ptr *QGamepad) DisconnectButtonL3Changed() {
	defer qt.Recovering("disconnect QGamepad::buttonL3Changed")

	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonL3Changed(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "buttonL3Changed")
	}
}

func (ptr *QGamepad) ButtonL3Changed(value bool) {
	defer qt.Recovering("QGamepad::buttonL3Changed")

	if ptr.Pointer() != nil {
		C.QGamepad_ButtonL3Changed(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_ButtonLeftChanged
func callbackQGamepad_ButtonLeftChanged(ptr unsafe.Pointer, value C.char) {
	defer qt.Recovering("callback QGamepad::buttonLeftChanged")

	if signal := qt.GetSignal(fmt.Sprintf("QGamepad(%v)", ptr), "buttonLeftChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonLeftChanged(f func(value bool)) {
	defer qt.Recovering("connect QGamepad::buttonLeftChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_ConnectButtonLeftChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "buttonLeftChanged", f)
	}
}

func (ptr *QGamepad) DisconnectButtonLeftChanged() {
	defer qt.Recovering("disconnect QGamepad::buttonLeftChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonLeftChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "buttonLeftChanged")
	}
}

func (ptr *QGamepad) ButtonLeftChanged(value bool) {
	defer qt.Recovering("QGamepad::buttonLeftChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_ButtonLeftChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_ButtonR1Changed
func callbackQGamepad_ButtonR1Changed(ptr unsafe.Pointer, value C.char) {
	defer qt.Recovering("callback QGamepad::buttonR1Changed")

	if signal := qt.GetSignal(fmt.Sprintf("QGamepad(%v)", ptr), "buttonR1Changed"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonR1Changed(f func(value bool)) {
	defer qt.Recovering("connect QGamepad::buttonR1Changed")

	if ptr.Pointer() != nil {
		C.QGamepad_ConnectButtonR1Changed(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "buttonR1Changed", f)
	}
}

func (ptr *QGamepad) DisconnectButtonR1Changed() {
	defer qt.Recovering("disconnect QGamepad::buttonR1Changed")

	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonR1Changed(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "buttonR1Changed")
	}
}

func (ptr *QGamepad) ButtonR1Changed(value bool) {
	defer qt.Recovering("QGamepad::buttonR1Changed")

	if ptr.Pointer() != nil {
		C.QGamepad_ButtonR1Changed(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_ButtonR2Changed
func callbackQGamepad_ButtonR2Changed(ptr unsafe.Pointer, value C.double) {
	defer qt.Recovering("callback QGamepad::buttonR2Changed")

	if signal := qt.GetSignal(fmt.Sprintf("QGamepad(%v)", ptr), "buttonR2Changed"); signal != nil {
		signal.(func(float64))(float64(value))
	}

}

func (ptr *QGamepad) ConnectButtonR2Changed(f func(value float64)) {
	defer qt.Recovering("connect QGamepad::buttonR2Changed")

	if ptr.Pointer() != nil {
		C.QGamepad_ConnectButtonR2Changed(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "buttonR2Changed", f)
	}
}

func (ptr *QGamepad) DisconnectButtonR2Changed() {
	defer qt.Recovering("disconnect QGamepad::buttonR2Changed")

	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonR2Changed(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "buttonR2Changed")
	}
}

func (ptr *QGamepad) ButtonR2Changed(value float64) {
	defer qt.Recovering("QGamepad::buttonR2Changed")

	if ptr.Pointer() != nil {
		C.QGamepad_ButtonR2Changed(ptr.Pointer(), C.double(value))
	}
}

//export callbackQGamepad_ButtonR3Changed
func callbackQGamepad_ButtonR3Changed(ptr unsafe.Pointer, value C.char) {
	defer qt.Recovering("callback QGamepad::buttonR3Changed")

	if signal := qt.GetSignal(fmt.Sprintf("QGamepad(%v)", ptr), "buttonR3Changed"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonR3Changed(f func(value bool)) {
	defer qt.Recovering("connect QGamepad::buttonR3Changed")

	if ptr.Pointer() != nil {
		C.QGamepad_ConnectButtonR3Changed(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "buttonR3Changed", f)
	}
}

func (ptr *QGamepad) DisconnectButtonR3Changed() {
	defer qt.Recovering("disconnect QGamepad::buttonR3Changed")

	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonR3Changed(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "buttonR3Changed")
	}
}

func (ptr *QGamepad) ButtonR3Changed(value bool) {
	defer qt.Recovering("QGamepad::buttonR3Changed")

	if ptr.Pointer() != nil {
		C.QGamepad_ButtonR3Changed(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_ButtonRightChanged
func callbackQGamepad_ButtonRightChanged(ptr unsafe.Pointer, value C.char) {
	defer qt.Recovering("callback QGamepad::buttonRightChanged")

	if signal := qt.GetSignal(fmt.Sprintf("QGamepad(%v)", ptr), "buttonRightChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonRightChanged(f func(value bool)) {
	defer qt.Recovering("connect QGamepad::buttonRightChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_ConnectButtonRightChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "buttonRightChanged", f)
	}
}

func (ptr *QGamepad) DisconnectButtonRightChanged() {
	defer qt.Recovering("disconnect QGamepad::buttonRightChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonRightChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "buttonRightChanged")
	}
}

func (ptr *QGamepad) ButtonRightChanged(value bool) {
	defer qt.Recovering("QGamepad::buttonRightChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_ButtonRightChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_ButtonSelectChanged
func callbackQGamepad_ButtonSelectChanged(ptr unsafe.Pointer, value C.char) {
	defer qt.Recovering("callback QGamepad::buttonSelectChanged")

	if signal := qt.GetSignal(fmt.Sprintf("QGamepad(%v)", ptr), "buttonSelectChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonSelectChanged(f func(value bool)) {
	defer qt.Recovering("connect QGamepad::buttonSelectChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_ConnectButtonSelectChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "buttonSelectChanged", f)
	}
}

func (ptr *QGamepad) DisconnectButtonSelectChanged() {
	defer qt.Recovering("disconnect QGamepad::buttonSelectChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonSelectChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "buttonSelectChanged")
	}
}

func (ptr *QGamepad) ButtonSelectChanged(value bool) {
	defer qt.Recovering("QGamepad::buttonSelectChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_ButtonSelectChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_ButtonStartChanged
func callbackQGamepad_ButtonStartChanged(ptr unsafe.Pointer, value C.char) {
	defer qt.Recovering("callback QGamepad::buttonStartChanged")

	if signal := qt.GetSignal(fmt.Sprintf("QGamepad(%v)", ptr), "buttonStartChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonStartChanged(f func(value bool)) {
	defer qt.Recovering("connect QGamepad::buttonStartChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_ConnectButtonStartChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "buttonStartChanged", f)
	}
}

func (ptr *QGamepad) DisconnectButtonStartChanged() {
	defer qt.Recovering("disconnect QGamepad::buttonStartChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonStartChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "buttonStartChanged")
	}
}

func (ptr *QGamepad) ButtonStartChanged(value bool) {
	defer qt.Recovering("QGamepad::buttonStartChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_ButtonStartChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_ButtonUpChanged
func callbackQGamepad_ButtonUpChanged(ptr unsafe.Pointer, value C.char) {
	defer qt.Recovering("callback QGamepad::buttonUpChanged")

	if signal := qt.GetSignal(fmt.Sprintf("QGamepad(%v)", ptr), "buttonUpChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonUpChanged(f func(value bool)) {
	defer qt.Recovering("connect QGamepad::buttonUpChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_ConnectButtonUpChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "buttonUpChanged", f)
	}
}

func (ptr *QGamepad) DisconnectButtonUpChanged() {
	defer qt.Recovering("disconnect QGamepad::buttonUpChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonUpChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "buttonUpChanged")
	}
}

func (ptr *QGamepad) ButtonUpChanged(value bool) {
	defer qt.Recovering("QGamepad::buttonUpChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_ButtonUpChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_ButtonXChanged
func callbackQGamepad_ButtonXChanged(ptr unsafe.Pointer, value C.char) {
	defer qt.Recovering("callback QGamepad::buttonXChanged")

	if signal := qt.GetSignal(fmt.Sprintf("QGamepad(%v)", ptr), "buttonXChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonXChanged(f func(value bool)) {
	defer qt.Recovering("connect QGamepad::buttonXChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_ConnectButtonXChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "buttonXChanged", f)
	}
}

func (ptr *QGamepad) DisconnectButtonXChanged() {
	defer qt.Recovering("disconnect QGamepad::buttonXChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonXChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "buttonXChanged")
	}
}

func (ptr *QGamepad) ButtonXChanged(value bool) {
	defer qt.Recovering("QGamepad::buttonXChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_ButtonXChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_ButtonYChanged
func callbackQGamepad_ButtonYChanged(ptr unsafe.Pointer, value C.char) {
	defer qt.Recovering("callback QGamepad::buttonYChanged")

	if signal := qt.GetSignal(fmt.Sprintf("QGamepad(%v)", ptr), "buttonYChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonYChanged(f func(value bool)) {
	defer qt.Recovering("connect QGamepad::buttonYChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_ConnectButtonYChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "buttonYChanged", f)
	}
}

func (ptr *QGamepad) DisconnectButtonYChanged() {
	defer qt.Recovering("disconnect QGamepad::buttonYChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonYChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "buttonYChanged")
	}
}

func (ptr *QGamepad) ButtonYChanged(value bool) {
	defer qt.Recovering("QGamepad::buttonYChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_ButtonYChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_ConnectedChanged
func callbackQGamepad_ConnectedChanged(ptr unsafe.Pointer, value C.char) {
	defer qt.Recovering("callback QGamepad::connectedChanged")

	if signal := qt.GetSignal(fmt.Sprintf("QGamepad(%v)", ptr), "connectedChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectConnectedChanged(f func(value bool)) {
	defer qt.Recovering("connect QGamepad::connectedChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_ConnectConnectedChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "connectedChanged", f)
	}
}

func (ptr *QGamepad) DisconnectConnectedChanged() {
	defer qt.Recovering("disconnect QGamepad::connectedChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectConnectedChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "connectedChanged")
	}
}

func (ptr *QGamepad) ConnectedChanged(value bool) {
	defer qt.Recovering("QGamepad::connectedChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_ConnectedChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_DeviceIdChanged
func callbackQGamepad_DeviceIdChanged(ptr unsafe.Pointer, value C.int) {
	defer qt.Recovering("callback QGamepad::deviceIdChanged")

	if signal := qt.GetSignal(fmt.Sprintf("QGamepad(%v)", ptr), "deviceIdChanged"); signal != nil {
		signal.(func(int))(int(int32(value)))
	}

}

func (ptr *QGamepad) ConnectDeviceIdChanged(f func(value int)) {
	defer qt.Recovering("connect QGamepad::deviceIdChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_ConnectDeviceIdChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "deviceIdChanged", f)
	}
}

func (ptr *QGamepad) DisconnectDeviceIdChanged() {
	defer qt.Recovering("disconnect QGamepad::deviceIdChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectDeviceIdChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "deviceIdChanged")
	}
}

func (ptr *QGamepad) DeviceIdChanged(value int) {
	defer qt.Recovering("QGamepad::deviceIdChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_DeviceIdChanged(ptr.Pointer(), C.int(int32(value)))
	}
}

//export callbackQGamepad_NameChanged
func callbackQGamepad_NameChanged(ptr unsafe.Pointer, value *C.char) {
	defer qt.Recovering("callback QGamepad::nameChanged")

	if signal := qt.GetSignal(fmt.Sprintf("QGamepad(%v)", ptr), "nameChanged"); signal != nil {
		signal.(func(string))(C.GoString(value))
	}

}

func (ptr *QGamepad) ConnectNameChanged(f func(value string)) {
	defer qt.Recovering("connect QGamepad::nameChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_ConnectNameChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "nameChanged", f)
	}
}

func (ptr *QGamepad) DisconnectNameChanged() {
	defer qt.Recovering("disconnect QGamepad::nameChanged")

	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectNameChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "nameChanged")
	}
}

func (ptr *QGamepad) NameChanged(value string) {
	defer qt.Recovering("QGamepad::nameChanged")

	if ptr.Pointer() != nil {
		var valueC = C.CString(value)
		defer C.free(unsafe.Pointer(valueC))
		C.QGamepad_NameChanged(ptr.Pointer(), valueC)
	}
}

func (ptr *QGamepad) DestroyQGamepad() {
	defer qt.Recovering("QGamepad::~QGamepad")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()))
		C.QGamepad_DestroyQGamepad(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQGamepad_TimerEvent
func callbackQGamepad_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QGamepad::timerEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QGamepad(%v)", ptr), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGamepadFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGamepad) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QGamepad::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "timerEvent", f)
	}
}

func (ptr *QGamepad) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QGamepad::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "timerEvent")
	}
}

func (ptr *QGamepad) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGamepad::timerEvent")

	if ptr.Pointer() != nil {
		C.QGamepad_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGamepad) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGamepad::timerEvent")

	if ptr.Pointer() != nil {
		C.QGamepad_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQGamepad_ChildEvent
func callbackQGamepad_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QGamepad::childEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QGamepad(%v)", ptr), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQGamepadFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QGamepad) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QGamepad::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "childEvent", f)
	}
}

func (ptr *QGamepad) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QGamepad::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "childEvent")
	}
}

func (ptr *QGamepad) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QGamepad::childEvent")

	if ptr.Pointer() != nil {
		C.QGamepad_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGamepad) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QGamepad::childEvent")

	if ptr.Pointer() != nil {
		C.QGamepad_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQGamepad_ConnectNotify
func callbackQGamepad_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QGamepad::connectNotify")

	if signal := qt.GetSignal(fmt.Sprintf("QGamepad(%v)", ptr), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGamepadFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGamepad) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QGamepad::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "connectNotify", f)
	}
}

func (ptr *QGamepad) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QGamepad::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "connectNotify")
	}
}

func (ptr *QGamepad) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QGamepad::connectNotify")

	if ptr.Pointer() != nil {
		C.QGamepad_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QGamepad) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QGamepad::connectNotify")

	if ptr.Pointer() != nil {
		C.QGamepad_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGamepad_CustomEvent
func callbackQGamepad_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QGamepad::customEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QGamepad(%v)", ptr), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGamepadFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGamepad) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGamepad::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "customEvent", f)
	}
}

func (ptr *QGamepad) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QGamepad::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "customEvent")
	}
}

func (ptr *QGamepad) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QGamepad::customEvent")

	if ptr.Pointer() != nil {
		C.QGamepad_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGamepad) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QGamepad::customEvent")

	if ptr.Pointer() != nil {
		C.QGamepad_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQGamepad_DeleteLater
func callbackQGamepad_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QGamepad::deleteLater")

	if signal := qt.GetSignal(fmt.Sprintf("QGamepad(%v)", ptr), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQGamepadFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QGamepad) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QGamepad::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "deleteLater", f)
	}
}

func (ptr *QGamepad) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QGamepad::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "deleteLater")
	}
}

func (ptr *QGamepad) DeleteLater() {
	defer qt.Recovering("QGamepad::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()))
		C.QGamepad_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGamepad) DeleteLaterDefault() {
	defer qt.Recovering("QGamepad::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()))
		C.QGamepad_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQGamepad_DisconnectNotify
func callbackQGamepad_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QGamepad::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprintf("QGamepad(%v)", ptr), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGamepadFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGamepad) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QGamepad::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "disconnectNotify", f)
	}
}

func (ptr *QGamepad) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QGamepad::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "disconnectNotify")
	}
}

func (ptr *QGamepad) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QGamepad::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QGamepad) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QGamepad::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGamepad_Event
func callbackQGamepad_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QGamepad::event")

	if signal := qt.GetSignal(fmt.Sprintf("QGamepad(%v)", ptr), "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGamepadFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QGamepad) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QGamepad::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "event", f)
	}
}

func (ptr *QGamepad) DisconnectEvent() {
	defer qt.Recovering("disconnect QGamepad::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "event")
	}
}

func (ptr *QGamepad) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QGamepad::event")

	if ptr.Pointer() != nil {
		return C.QGamepad_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QGamepad) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QGamepad::event")

	if ptr.Pointer() != nil {
		return C.QGamepad_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQGamepad_EventFilter
func callbackQGamepad_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QGamepad::eventFilter")

	if signal := qt.GetSignal(fmt.Sprintf("QGamepad(%v)", ptr), "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGamepadFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QGamepad) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QGamepad::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "eventFilter", f)
	}
}

func (ptr *QGamepad) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QGamepad::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "eventFilter")
	}
}

func (ptr *QGamepad) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QGamepad::eventFilter")

	if ptr.Pointer() != nil {
		return C.QGamepad_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QGamepad) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QGamepad::eventFilter")

	if ptr.Pointer() != nil {
		return C.QGamepad_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQGamepad_MetaObject
func callbackQGamepad_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QGamepad::metaObject")

	if signal := qt.GetSignal(fmt.Sprintf("QGamepad(%v)", ptr), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQGamepadFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QGamepad) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QGamepad::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "metaObject", f)
	}
}

func (ptr *QGamepad) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QGamepad::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QGamepad(%v)", ptr.Pointer()), "metaObject")
	}
}

func (ptr *QGamepad) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QGamepad::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QGamepad_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGamepad) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QGamepad::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QGamepad_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}
