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

func (ptr *QGamepad) ButtonL2() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGamepad_ButtonL2(ptr.Pointer()))
	}
	return 0
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

func (ptr *QGamepad) ButtonR2() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGamepad_ButtonR2(ptr.Pointer()))
	}
	return 0
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

func (ptr *QGamepad) DeviceId() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QGamepad_DeviceId(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGamepad) IsConnected() bool {
	if ptr.Pointer() != nil {
		return C.QGamepad_IsConnected(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGamepad) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGamepad_Name(ptr.Pointer()))
	}
	return ""
}

//export callbackQGamepad_SetDeviceId
func callbackQGamepad_SetDeviceId(ptr unsafe.Pointer, number C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGamepad::setDeviceId"); signal != nil {
		signal.(func(int))(int(int32(number)))
	}

}

func (ptr *QGamepad) ConnectSetDeviceId(f func(number int)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::setDeviceId", f)
	}
}

func (ptr *QGamepad) DisconnectSetDeviceId(number int) {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::setDeviceId")
	}
}

func (ptr *QGamepad) SetDeviceId(number int) {
	if ptr.Pointer() != nil {
		C.QGamepad_SetDeviceId(ptr.Pointer(), C.int(int32(number)))
	}
}

func NewQGamepad(deviceId int, parent core.QObject_ITF) *QGamepad {
	var tmpValue = NewQGamepadFromPointer(C.QGamepad_NewQGamepad(C.int(int32(deviceId)), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQGamepad_AxisLeftXChanged
func callbackQGamepad_AxisLeftXChanged(ptr unsafe.Pointer, value C.double) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGamepad::axisLeftXChanged"); signal != nil {
		signal.(func(float64))(float64(value))
	}

}

func (ptr *QGamepad) ConnectAxisLeftXChanged(f func(value float64)) {
	if ptr.Pointer() != nil {
		C.QGamepad_ConnectAxisLeftXChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::axisLeftXChanged", f)
	}
}

func (ptr *QGamepad) DisconnectAxisLeftXChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectAxisLeftXChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::axisLeftXChanged")
	}
}

func (ptr *QGamepad) AxisLeftXChanged(value float64) {
	if ptr.Pointer() != nil {
		C.QGamepad_AxisLeftXChanged(ptr.Pointer(), C.double(value))
	}
}

//export callbackQGamepad_AxisLeftYChanged
func callbackQGamepad_AxisLeftYChanged(ptr unsafe.Pointer, value C.double) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGamepad::axisLeftYChanged"); signal != nil {
		signal.(func(float64))(float64(value))
	}

}

func (ptr *QGamepad) ConnectAxisLeftYChanged(f func(value float64)) {
	if ptr.Pointer() != nil {
		C.QGamepad_ConnectAxisLeftYChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::axisLeftYChanged", f)
	}
}

func (ptr *QGamepad) DisconnectAxisLeftYChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectAxisLeftYChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::axisLeftYChanged")
	}
}

func (ptr *QGamepad) AxisLeftYChanged(value float64) {
	if ptr.Pointer() != nil {
		C.QGamepad_AxisLeftYChanged(ptr.Pointer(), C.double(value))
	}
}

//export callbackQGamepad_AxisRightXChanged
func callbackQGamepad_AxisRightXChanged(ptr unsafe.Pointer, value C.double) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGamepad::axisRightXChanged"); signal != nil {
		signal.(func(float64))(float64(value))
	}

}

func (ptr *QGamepad) ConnectAxisRightXChanged(f func(value float64)) {
	if ptr.Pointer() != nil {
		C.QGamepad_ConnectAxisRightXChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::axisRightXChanged", f)
	}
}

func (ptr *QGamepad) DisconnectAxisRightXChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectAxisRightXChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::axisRightXChanged")
	}
}

func (ptr *QGamepad) AxisRightXChanged(value float64) {
	if ptr.Pointer() != nil {
		C.QGamepad_AxisRightXChanged(ptr.Pointer(), C.double(value))
	}
}

//export callbackQGamepad_AxisRightYChanged
func callbackQGamepad_AxisRightYChanged(ptr unsafe.Pointer, value C.double) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGamepad::axisRightYChanged"); signal != nil {
		signal.(func(float64))(float64(value))
	}

}

func (ptr *QGamepad) ConnectAxisRightYChanged(f func(value float64)) {
	if ptr.Pointer() != nil {
		C.QGamepad_ConnectAxisRightYChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::axisRightYChanged", f)
	}
}

func (ptr *QGamepad) DisconnectAxisRightYChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectAxisRightYChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::axisRightYChanged")
	}
}

func (ptr *QGamepad) AxisRightYChanged(value float64) {
	if ptr.Pointer() != nil {
		C.QGamepad_AxisRightYChanged(ptr.Pointer(), C.double(value))
	}
}

//export callbackQGamepad_ButtonAChanged
func callbackQGamepad_ButtonAChanged(ptr unsafe.Pointer, value C.char) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGamepad::buttonAChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonAChanged(f func(value bool)) {
	if ptr.Pointer() != nil {
		C.QGamepad_ConnectButtonAChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::buttonAChanged", f)
	}
}

func (ptr *QGamepad) DisconnectButtonAChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonAChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::buttonAChanged")
	}
}

func (ptr *QGamepad) ButtonAChanged(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonAChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_ButtonBChanged
func callbackQGamepad_ButtonBChanged(ptr unsafe.Pointer, value C.char) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGamepad::buttonBChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonBChanged(f func(value bool)) {
	if ptr.Pointer() != nil {
		C.QGamepad_ConnectButtonBChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::buttonBChanged", f)
	}
}

func (ptr *QGamepad) DisconnectButtonBChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonBChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::buttonBChanged")
	}
}

func (ptr *QGamepad) ButtonBChanged(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonBChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_ButtonCenterChanged
func callbackQGamepad_ButtonCenterChanged(ptr unsafe.Pointer, value C.char) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGamepad::buttonCenterChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonCenterChanged(f func(value bool)) {
	if ptr.Pointer() != nil {
		C.QGamepad_ConnectButtonCenterChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::buttonCenterChanged", f)
	}
}

func (ptr *QGamepad) DisconnectButtonCenterChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonCenterChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::buttonCenterChanged")
	}
}

func (ptr *QGamepad) ButtonCenterChanged(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonCenterChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_ButtonDownChanged
func callbackQGamepad_ButtonDownChanged(ptr unsafe.Pointer, value C.char) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGamepad::buttonDownChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonDownChanged(f func(value bool)) {
	if ptr.Pointer() != nil {
		C.QGamepad_ConnectButtonDownChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::buttonDownChanged", f)
	}
}

func (ptr *QGamepad) DisconnectButtonDownChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonDownChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::buttonDownChanged")
	}
}

func (ptr *QGamepad) ButtonDownChanged(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonDownChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_ButtonGuideChanged
func callbackQGamepad_ButtonGuideChanged(ptr unsafe.Pointer, value C.char) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGamepad::buttonGuideChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonGuideChanged(f func(value bool)) {
	if ptr.Pointer() != nil {
		C.QGamepad_ConnectButtonGuideChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::buttonGuideChanged", f)
	}
}

func (ptr *QGamepad) DisconnectButtonGuideChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonGuideChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::buttonGuideChanged")
	}
}

func (ptr *QGamepad) ButtonGuideChanged(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonGuideChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_ButtonL1Changed
func callbackQGamepad_ButtonL1Changed(ptr unsafe.Pointer, value C.char) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGamepad::buttonL1Changed"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonL1Changed(f func(value bool)) {
	if ptr.Pointer() != nil {
		C.QGamepad_ConnectButtonL1Changed(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::buttonL1Changed", f)
	}
}

func (ptr *QGamepad) DisconnectButtonL1Changed() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonL1Changed(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::buttonL1Changed")
	}
}

func (ptr *QGamepad) ButtonL1Changed(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonL1Changed(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_ButtonL2Changed
func callbackQGamepad_ButtonL2Changed(ptr unsafe.Pointer, value C.double) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGamepad::buttonL2Changed"); signal != nil {
		signal.(func(float64))(float64(value))
	}

}

func (ptr *QGamepad) ConnectButtonL2Changed(f func(value float64)) {
	if ptr.Pointer() != nil {
		C.QGamepad_ConnectButtonL2Changed(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::buttonL2Changed", f)
	}
}

func (ptr *QGamepad) DisconnectButtonL2Changed() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonL2Changed(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::buttonL2Changed")
	}
}

func (ptr *QGamepad) ButtonL2Changed(value float64) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonL2Changed(ptr.Pointer(), C.double(value))
	}
}

//export callbackQGamepad_ButtonL3Changed
func callbackQGamepad_ButtonL3Changed(ptr unsafe.Pointer, value C.char) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGamepad::buttonL3Changed"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonL3Changed(f func(value bool)) {
	if ptr.Pointer() != nil {
		C.QGamepad_ConnectButtonL3Changed(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::buttonL3Changed", f)
	}
}

func (ptr *QGamepad) DisconnectButtonL3Changed() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonL3Changed(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::buttonL3Changed")
	}
}

func (ptr *QGamepad) ButtonL3Changed(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonL3Changed(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_ButtonLeftChanged
func callbackQGamepad_ButtonLeftChanged(ptr unsafe.Pointer, value C.char) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGamepad::buttonLeftChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonLeftChanged(f func(value bool)) {
	if ptr.Pointer() != nil {
		C.QGamepad_ConnectButtonLeftChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::buttonLeftChanged", f)
	}
}

func (ptr *QGamepad) DisconnectButtonLeftChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonLeftChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::buttonLeftChanged")
	}
}

func (ptr *QGamepad) ButtonLeftChanged(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonLeftChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_ButtonR1Changed
func callbackQGamepad_ButtonR1Changed(ptr unsafe.Pointer, value C.char) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGamepad::buttonR1Changed"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonR1Changed(f func(value bool)) {
	if ptr.Pointer() != nil {
		C.QGamepad_ConnectButtonR1Changed(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::buttonR1Changed", f)
	}
}

func (ptr *QGamepad) DisconnectButtonR1Changed() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonR1Changed(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::buttonR1Changed")
	}
}

func (ptr *QGamepad) ButtonR1Changed(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonR1Changed(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_ButtonR2Changed
func callbackQGamepad_ButtonR2Changed(ptr unsafe.Pointer, value C.double) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGamepad::buttonR2Changed"); signal != nil {
		signal.(func(float64))(float64(value))
	}

}

func (ptr *QGamepad) ConnectButtonR2Changed(f func(value float64)) {
	if ptr.Pointer() != nil {
		C.QGamepad_ConnectButtonR2Changed(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::buttonR2Changed", f)
	}
}

func (ptr *QGamepad) DisconnectButtonR2Changed() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonR2Changed(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::buttonR2Changed")
	}
}

func (ptr *QGamepad) ButtonR2Changed(value float64) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonR2Changed(ptr.Pointer(), C.double(value))
	}
}

//export callbackQGamepad_ButtonR3Changed
func callbackQGamepad_ButtonR3Changed(ptr unsafe.Pointer, value C.char) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGamepad::buttonR3Changed"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonR3Changed(f func(value bool)) {
	if ptr.Pointer() != nil {
		C.QGamepad_ConnectButtonR3Changed(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::buttonR3Changed", f)
	}
}

func (ptr *QGamepad) DisconnectButtonR3Changed() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonR3Changed(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::buttonR3Changed")
	}
}

func (ptr *QGamepad) ButtonR3Changed(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonR3Changed(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_ButtonRightChanged
func callbackQGamepad_ButtonRightChanged(ptr unsafe.Pointer, value C.char) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGamepad::buttonRightChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonRightChanged(f func(value bool)) {
	if ptr.Pointer() != nil {
		C.QGamepad_ConnectButtonRightChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::buttonRightChanged", f)
	}
}

func (ptr *QGamepad) DisconnectButtonRightChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonRightChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::buttonRightChanged")
	}
}

func (ptr *QGamepad) ButtonRightChanged(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonRightChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_ButtonSelectChanged
func callbackQGamepad_ButtonSelectChanged(ptr unsafe.Pointer, value C.char) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGamepad::buttonSelectChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonSelectChanged(f func(value bool)) {
	if ptr.Pointer() != nil {
		C.QGamepad_ConnectButtonSelectChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::buttonSelectChanged", f)
	}
}

func (ptr *QGamepad) DisconnectButtonSelectChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonSelectChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::buttonSelectChanged")
	}
}

func (ptr *QGamepad) ButtonSelectChanged(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonSelectChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_ButtonStartChanged
func callbackQGamepad_ButtonStartChanged(ptr unsafe.Pointer, value C.char) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGamepad::buttonStartChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonStartChanged(f func(value bool)) {
	if ptr.Pointer() != nil {
		C.QGamepad_ConnectButtonStartChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::buttonStartChanged", f)
	}
}

func (ptr *QGamepad) DisconnectButtonStartChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonStartChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::buttonStartChanged")
	}
}

func (ptr *QGamepad) ButtonStartChanged(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonStartChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_ButtonUpChanged
func callbackQGamepad_ButtonUpChanged(ptr unsafe.Pointer, value C.char) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGamepad::buttonUpChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonUpChanged(f func(value bool)) {
	if ptr.Pointer() != nil {
		C.QGamepad_ConnectButtonUpChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::buttonUpChanged", f)
	}
}

func (ptr *QGamepad) DisconnectButtonUpChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonUpChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::buttonUpChanged")
	}
}

func (ptr *QGamepad) ButtonUpChanged(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonUpChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_ButtonXChanged
func callbackQGamepad_ButtonXChanged(ptr unsafe.Pointer, value C.char) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGamepad::buttonXChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonXChanged(f func(value bool)) {
	if ptr.Pointer() != nil {
		C.QGamepad_ConnectButtonXChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::buttonXChanged", f)
	}
}

func (ptr *QGamepad) DisconnectButtonXChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonXChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::buttonXChanged")
	}
}

func (ptr *QGamepad) ButtonXChanged(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonXChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_ButtonYChanged
func callbackQGamepad_ButtonYChanged(ptr unsafe.Pointer, value C.char) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGamepad::buttonYChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectButtonYChanged(f func(value bool)) {
	if ptr.Pointer() != nil {
		C.QGamepad_ConnectButtonYChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::buttonYChanged", f)
	}
}

func (ptr *QGamepad) DisconnectButtonYChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectButtonYChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::buttonYChanged")
	}
}

func (ptr *QGamepad) ButtonYChanged(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ButtonYChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_ConnectedChanged
func callbackQGamepad_ConnectedChanged(ptr unsafe.Pointer, value C.char) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGamepad::connectedChanged"); signal != nil {
		signal.(func(bool))(int8(value) != 0)
	}

}

func (ptr *QGamepad) ConnectConnectedChanged(f func(value bool)) {
	if ptr.Pointer() != nil {
		C.QGamepad_ConnectConnectedChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::connectedChanged", f)
	}
}

func (ptr *QGamepad) DisconnectConnectedChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectConnectedChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::connectedChanged")
	}
}

func (ptr *QGamepad) ConnectedChanged(value bool) {
	if ptr.Pointer() != nil {
		C.QGamepad_ConnectedChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQGamepad_DeviceIdChanged
func callbackQGamepad_DeviceIdChanged(ptr unsafe.Pointer, value C.int) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGamepad::deviceIdChanged"); signal != nil {
		signal.(func(int))(int(int32(value)))
	}

}

func (ptr *QGamepad) ConnectDeviceIdChanged(f func(value int)) {
	if ptr.Pointer() != nil {
		C.QGamepad_ConnectDeviceIdChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::deviceIdChanged", f)
	}
}

func (ptr *QGamepad) DisconnectDeviceIdChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectDeviceIdChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::deviceIdChanged")
	}
}

func (ptr *QGamepad) DeviceIdChanged(value int) {
	if ptr.Pointer() != nil {
		C.QGamepad_DeviceIdChanged(ptr.Pointer(), C.int(int32(value)))
	}
}

//export callbackQGamepad_NameChanged
func callbackQGamepad_NameChanged(ptr unsafe.Pointer, value C.struct_QtGamepad_PackedString) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGamepad::nameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(value))
	}

}

func (ptr *QGamepad) ConnectNameChanged(f func(value string)) {
	if ptr.Pointer() != nil {
		C.QGamepad_ConnectNameChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::nameChanged", f)
	}
}

func (ptr *QGamepad) DisconnectNameChanged() {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectNameChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::nameChanged")
	}
}

func (ptr *QGamepad) NameChanged(value string) {
	if ptr.Pointer() != nil {
		var valueC = C.CString(value)
		defer C.free(unsafe.Pointer(valueC))
		C.QGamepad_NameChanged(ptr.Pointer(), valueC)
	}
}

func (ptr *QGamepad) DestroyQGamepad() {
	if ptr.Pointer() != nil {
		C.QGamepad_DestroyQGamepad(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQGamepad_TimerEvent
func callbackQGamepad_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGamepad::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGamepadFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGamepad) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::timerEvent", f)
	}
}

func (ptr *QGamepad) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::timerEvent")
	}
}

func (ptr *QGamepad) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGamepad_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGamepad) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGamepad_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQGamepad_ChildEvent
func callbackQGamepad_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGamepad::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQGamepadFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QGamepad) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::childEvent", f)
	}
}

func (ptr *QGamepad) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::childEvent")
	}
}

func (ptr *QGamepad) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGamepad_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGamepad) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGamepad_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQGamepad_ConnectNotify
func callbackQGamepad_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGamepad::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGamepadFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGamepad) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::connectNotify", f)
	}
}

func (ptr *QGamepad) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::connectNotify")
	}
}

func (ptr *QGamepad) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGamepad_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QGamepad) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGamepad_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGamepad_CustomEvent
func callbackQGamepad_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGamepad::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGamepadFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGamepad) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::customEvent", f)
	}
}

func (ptr *QGamepad) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::customEvent")
	}
}

func (ptr *QGamepad) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGamepad_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGamepad) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGamepad_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQGamepad_DeleteLater
func callbackQGamepad_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGamepad::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQGamepadFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QGamepad) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::deleteLater", f)
	}
}

func (ptr *QGamepad) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::deleteLater")
	}
}

func (ptr *QGamepad) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QGamepad_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QGamepad) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QGamepad_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQGamepad_DisconnectNotify
func callbackQGamepad_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGamepad::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGamepadFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGamepad) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::disconnectNotify", f)
	}
}

func (ptr *QGamepad) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::disconnectNotify")
	}
}

func (ptr *QGamepad) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QGamepad) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGamepad_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGamepad_Event
func callbackQGamepad_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGamepad::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGamepadFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QGamepad) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::event", f)
	}
}

func (ptr *QGamepad) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::event")
	}
}

func (ptr *QGamepad) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGamepad_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QGamepad) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGamepad_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQGamepad_EventFilter
func callbackQGamepad_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGamepad::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGamepadFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QGamepad) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::eventFilter", f)
	}
}

func (ptr *QGamepad) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::eventFilter")
	}
}

func (ptr *QGamepad) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGamepad_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QGamepad) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGamepad_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQGamepad_MetaObject
func callbackQGamepad_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGamepad::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQGamepadFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QGamepad) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::metaObject", f)
	}
}

func (ptr *QGamepad) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGamepad::metaObject")
	}
}

func (ptr *QGamepad) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QGamepad_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGamepad) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QGamepad_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}
