// +build !minimal

package gamepad

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/internal"
	"strings"
	"unsafe"
)

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

func (n *QGamepad) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QGamepad) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQGamepadFromPointer(ptr unsafe.Pointer) (n *QGamepad) {
	n = new(QGamepad)
	n.InitFromInternal(uintptr(ptr), "gamepad.QGamepad")
	return
}

func (ptr *QGamepad) DestroyQGamepad() {
}

func NewQGamepad(deviceId int, parent core.QObject_ITF) *QGamepad {

	return internal.CallLocalFunction([]interface{}{"", "", "gamepad.NewQGamepad", "", deviceId, parent}).(*QGamepad)
}

func (ptr *QGamepad) AxisLeftX() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AxisLeftX"}).(float64)
}

func (ptr *QGamepad) ConnectAxisLeftXChanged(f func(value float64)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAxisLeftXChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepad) DisconnectAxisLeftXChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAxisLeftXChanged"})
}

func (ptr *QGamepad) AxisLeftXChanged(value float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AxisLeftXChanged", value})
}

func (ptr *QGamepad) AxisLeftY() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AxisLeftY"}).(float64)
}

func (ptr *QGamepad) ConnectAxisLeftYChanged(f func(value float64)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAxisLeftYChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepad) DisconnectAxisLeftYChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAxisLeftYChanged"})
}

func (ptr *QGamepad) AxisLeftYChanged(value float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AxisLeftYChanged", value})
}

func (ptr *QGamepad) AxisRightX() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AxisRightX"}).(float64)
}

func (ptr *QGamepad) ConnectAxisRightXChanged(f func(value float64)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAxisRightXChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepad) DisconnectAxisRightXChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAxisRightXChanged"})
}

func (ptr *QGamepad) AxisRightXChanged(value float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AxisRightXChanged", value})
}

func (ptr *QGamepad) AxisRightY() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AxisRightY"}).(float64)
}

func (ptr *QGamepad) ConnectAxisRightYChanged(f func(value float64)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAxisRightYChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepad) DisconnectAxisRightYChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAxisRightYChanged"})
}

func (ptr *QGamepad) AxisRightYChanged(value float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AxisRightYChanged", value})
}

func (ptr *QGamepad) ButtonA() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonA"}).(bool)
}

func (ptr *QGamepad) ConnectButtonAChanged(f func(value bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectButtonAChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepad) DisconnectButtonAChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectButtonAChanged"})
}

func (ptr *QGamepad) ButtonAChanged(value bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonAChanged", value})
}

func (ptr *QGamepad) ButtonB() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonB"}).(bool)
}

func (ptr *QGamepad) ConnectButtonBChanged(f func(value bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectButtonBChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepad) DisconnectButtonBChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectButtonBChanged"})
}

func (ptr *QGamepad) ButtonBChanged(value bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonBChanged", value})
}

func (ptr *QGamepad) ButtonCenter() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonCenter"}).(bool)
}

func (ptr *QGamepad) ConnectButtonCenterChanged(f func(value bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectButtonCenterChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepad) DisconnectButtonCenterChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectButtonCenterChanged"})
}

func (ptr *QGamepad) ButtonCenterChanged(value bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonCenterChanged", value})
}

func (ptr *QGamepad) ButtonDown() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonDown"}).(bool)
}

func (ptr *QGamepad) ConnectButtonDownChanged(f func(value bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectButtonDownChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepad) DisconnectButtonDownChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectButtonDownChanged"})
}

func (ptr *QGamepad) ButtonDownChanged(value bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonDownChanged", value})
}

func (ptr *QGamepad) ButtonGuide() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonGuide"}).(bool)
}

func (ptr *QGamepad) ConnectButtonGuideChanged(f func(value bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectButtonGuideChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepad) DisconnectButtonGuideChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectButtonGuideChanged"})
}

func (ptr *QGamepad) ButtonGuideChanged(value bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonGuideChanged", value})
}

func (ptr *QGamepad) ButtonL1() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonL1"}).(bool)
}

func (ptr *QGamepad) ConnectButtonL1Changed(f func(value bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectButtonL1Changed", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepad) DisconnectButtonL1Changed() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectButtonL1Changed"})
}

func (ptr *QGamepad) ButtonL1Changed(value bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonL1Changed", value})
}

func (ptr *QGamepad) ButtonL2() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonL2"}).(float64)
}

func (ptr *QGamepad) ConnectButtonL2Changed(f func(value float64)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectButtonL2Changed", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepad) DisconnectButtonL2Changed() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectButtonL2Changed"})
}

func (ptr *QGamepad) ButtonL2Changed(value float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonL2Changed", value})
}

func (ptr *QGamepad) ButtonL3() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonL3"}).(bool)
}

func (ptr *QGamepad) ConnectButtonL3Changed(f func(value bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectButtonL3Changed", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepad) DisconnectButtonL3Changed() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectButtonL3Changed"})
}

func (ptr *QGamepad) ButtonL3Changed(value bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonL3Changed", value})
}

func (ptr *QGamepad) ButtonLeft() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonLeft"}).(bool)
}

func (ptr *QGamepad) ConnectButtonLeftChanged(f func(value bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectButtonLeftChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepad) DisconnectButtonLeftChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectButtonLeftChanged"})
}

func (ptr *QGamepad) ButtonLeftChanged(value bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonLeftChanged", value})
}

func (ptr *QGamepad) ButtonR1() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonR1"}).(bool)
}

func (ptr *QGamepad) ConnectButtonR1Changed(f func(value bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectButtonR1Changed", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepad) DisconnectButtonR1Changed() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectButtonR1Changed"})
}

func (ptr *QGamepad) ButtonR1Changed(value bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonR1Changed", value})
}

func (ptr *QGamepad) ButtonR2() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonR2"}).(float64)
}

func (ptr *QGamepad) ConnectButtonR2Changed(f func(value float64)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectButtonR2Changed", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepad) DisconnectButtonR2Changed() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectButtonR2Changed"})
}

func (ptr *QGamepad) ButtonR2Changed(value float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonR2Changed", value})
}

func (ptr *QGamepad) ButtonR3() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonR3"}).(bool)
}

func (ptr *QGamepad) ConnectButtonR3Changed(f func(value bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectButtonR3Changed", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepad) DisconnectButtonR3Changed() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectButtonR3Changed"})
}

func (ptr *QGamepad) ButtonR3Changed(value bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonR3Changed", value})
}

func (ptr *QGamepad) ButtonRight() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonRight"}).(bool)
}

func (ptr *QGamepad) ConnectButtonRightChanged(f func(value bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectButtonRightChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepad) DisconnectButtonRightChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectButtonRightChanged"})
}

func (ptr *QGamepad) ButtonRightChanged(value bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonRightChanged", value})
}

func (ptr *QGamepad) ButtonSelect() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonSelect"}).(bool)
}

func (ptr *QGamepad) ConnectButtonSelectChanged(f func(value bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectButtonSelectChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepad) DisconnectButtonSelectChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectButtonSelectChanged"})
}

func (ptr *QGamepad) ButtonSelectChanged(value bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonSelectChanged", value})
}

func (ptr *QGamepad) ButtonStart() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonStart"}).(bool)
}

func (ptr *QGamepad) ConnectButtonStartChanged(f func(value bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectButtonStartChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepad) DisconnectButtonStartChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectButtonStartChanged"})
}

func (ptr *QGamepad) ButtonStartChanged(value bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonStartChanged", value})
}

func (ptr *QGamepad) ButtonUp() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonUp"}).(bool)
}

func (ptr *QGamepad) ConnectButtonUpChanged(f func(value bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectButtonUpChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepad) DisconnectButtonUpChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectButtonUpChanged"})
}

func (ptr *QGamepad) ButtonUpChanged(value bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonUpChanged", value})
}

func (ptr *QGamepad) ButtonX() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonX"}).(bool)
}

func (ptr *QGamepad) ConnectButtonXChanged(f func(value bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectButtonXChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepad) DisconnectButtonXChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectButtonXChanged"})
}

func (ptr *QGamepad) ButtonXChanged(value bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonXChanged", value})
}

func (ptr *QGamepad) ButtonY() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonY"}).(bool)
}

func (ptr *QGamepad) ConnectButtonYChanged(f func(value bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectButtonYChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepad) DisconnectButtonYChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectButtonYChanged"})
}

func (ptr *QGamepad) ButtonYChanged(value bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonYChanged", value})
}

func (ptr *QGamepad) ConnectConnectedChanged(f func(value bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectConnectedChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepad) DisconnectConnectedChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectConnectedChanged"})
}

func (ptr *QGamepad) ConnectedChanged(value bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectedChanged", value})
}

func (ptr *QGamepad) DeviceId() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeviceId"}).(float64))
}

func (ptr *QGamepad) ConnectDeviceIdChanged(f func(value int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDeviceIdChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepad) DisconnectDeviceIdChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDeviceIdChanged"})
}

func (ptr *QGamepad) DeviceIdChanged(value int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeviceIdChanged", value})
}

func (ptr *QGamepad) IsConnected() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsConnected"}).(bool)
}

func (ptr *QGamepad) Name() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Name"}).(string)
}

func (ptr *QGamepad) ConnectNameChanged(f func(value string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNameChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepad) DisconnectNameChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNameChanged"})
}

func (ptr *QGamepad) NameChanged(value string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NameChanged", value})
}

func (ptr *QGamepad) ConnectSetDeviceId(f func(number int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetDeviceId", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepad) DisconnectSetDeviceId() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetDeviceId"})
}

func (ptr *QGamepad) SetDeviceId(number int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDeviceId", number})
}

func (ptr *QGamepad) SetDeviceIdDefault(number int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDeviceIdDefault", number})
}

func (ptr *QGamepad) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QGamepad) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QGamepad) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QGamepad) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QGamepad) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QGamepad) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QGamepad) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QGamepad) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QGamepad) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QGamepad) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QGamepad) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QGamepad) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QGamepad) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QGamepad) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QGamepad) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QGamepad) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QGamepad) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QGamepad) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QGamepad) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QGamepad) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QGamepad) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
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

func (n *QGamepadKeyNavigation) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QGamepadKeyNavigation) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQGamepadKeyNavigationFromPointer(ptr unsafe.Pointer) (n *QGamepadKeyNavigation) {
	n = new(QGamepadKeyNavigation)
	n.InitFromInternal(uintptr(ptr), "gamepad.QGamepadKeyNavigation")
	return
}

func (ptr *QGamepadKeyNavigation) DestroyQGamepadKeyNavigation() {
}

func NewQGamepadKeyNavigation(parent core.QObject_ITF) *QGamepadKeyNavigation {

	return internal.CallLocalFunction([]interface{}{"", "", "gamepad.NewQGamepadKeyNavigation", "", parent}).(*QGamepadKeyNavigation)
}

func (ptr *QGamepadKeyNavigation) Active() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Active"}).(bool)
}

func (ptr *QGamepadKeyNavigation) ConnectActiveChanged(f func(isActive bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectActiveChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepadKeyNavigation) DisconnectActiveChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectActiveChanged"})
}

func (ptr *QGamepadKeyNavigation) ActiveChanged(isActive bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ActiveChanged", isActive})
}

func (ptr *QGamepadKeyNavigation) ButtonAKey() core.Qt__Key {

	return core.Qt__Key(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonAKey"}).(float64))
}

func (ptr *QGamepadKeyNavigation) ConnectButtonAKeyChanged(f func(key core.Qt__Key)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectButtonAKeyChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepadKeyNavigation) DisconnectButtonAKeyChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectButtonAKeyChanged"})
}

func (ptr *QGamepadKeyNavigation) ButtonAKeyChanged(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonAKeyChanged", key})
}

func (ptr *QGamepadKeyNavigation) ButtonBKey() core.Qt__Key {

	return core.Qt__Key(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonBKey"}).(float64))
}

func (ptr *QGamepadKeyNavigation) ConnectButtonBKeyChanged(f func(key core.Qt__Key)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectButtonBKeyChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepadKeyNavigation) DisconnectButtonBKeyChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectButtonBKeyChanged"})
}

func (ptr *QGamepadKeyNavigation) ButtonBKeyChanged(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonBKeyChanged", key})
}

func (ptr *QGamepadKeyNavigation) ButtonGuideKey() core.Qt__Key {

	return core.Qt__Key(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonGuideKey"}).(float64))
}

func (ptr *QGamepadKeyNavigation) ConnectButtonGuideKeyChanged(f func(key core.Qt__Key)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectButtonGuideKeyChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepadKeyNavigation) DisconnectButtonGuideKeyChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectButtonGuideKeyChanged"})
}

func (ptr *QGamepadKeyNavigation) ButtonGuideKeyChanged(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonGuideKeyChanged", key})
}

func (ptr *QGamepadKeyNavigation) ButtonL1Key() core.Qt__Key {

	return core.Qt__Key(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonL1Key"}).(float64))
}

func (ptr *QGamepadKeyNavigation) ConnectButtonL1KeyChanged(f func(key core.Qt__Key)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectButtonL1KeyChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepadKeyNavigation) DisconnectButtonL1KeyChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectButtonL1KeyChanged"})
}

func (ptr *QGamepadKeyNavigation) ButtonL1KeyChanged(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonL1KeyChanged", key})
}

func (ptr *QGamepadKeyNavigation) ButtonL2Key() core.Qt__Key {

	return core.Qt__Key(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonL2Key"}).(float64))
}

func (ptr *QGamepadKeyNavigation) ConnectButtonL2KeyChanged(f func(key core.Qt__Key)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectButtonL2KeyChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepadKeyNavigation) DisconnectButtonL2KeyChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectButtonL2KeyChanged"})
}

func (ptr *QGamepadKeyNavigation) ButtonL2KeyChanged(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonL2KeyChanged", key})
}

func (ptr *QGamepadKeyNavigation) ButtonL3Key() core.Qt__Key {

	return core.Qt__Key(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonL3Key"}).(float64))
}

func (ptr *QGamepadKeyNavigation) ConnectButtonL3KeyChanged(f func(key core.Qt__Key)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectButtonL3KeyChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepadKeyNavigation) DisconnectButtonL3KeyChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectButtonL3KeyChanged"})
}

func (ptr *QGamepadKeyNavigation) ButtonL3KeyChanged(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonL3KeyChanged", key})
}

func (ptr *QGamepadKeyNavigation) ButtonR1Key() core.Qt__Key {

	return core.Qt__Key(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonR1Key"}).(float64))
}

func (ptr *QGamepadKeyNavigation) ConnectButtonR1KeyChanged(f func(key core.Qt__Key)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectButtonR1KeyChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepadKeyNavigation) DisconnectButtonR1KeyChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectButtonR1KeyChanged"})
}

func (ptr *QGamepadKeyNavigation) ButtonR1KeyChanged(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonR1KeyChanged", key})
}

func (ptr *QGamepadKeyNavigation) ButtonR2Key() core.Qt__Key {

	return core.Qt__Key(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonR2Key"}).(float64))
}

func (ptr *QGamepadKeyNavigation) ConnectButtonR2KeyChanged(f func(key core.Qt__Key)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectButtonR2KeyChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepadKeyNavigation) DisconnectButtonR2KeyChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectButtonR2KeyChanged"})
}

func (ptr *QGamepadKeyNavigation) ButtonR2KeyChanged(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonR2KeyChanged", key})
}

func (ptr *QGamepadKeyNavigation) ButtonR3Key() core.Qt__Key {

	return core.Qt__Key(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonR3Key"}).(float64))
}

func (ptr *QGamepadKeyNavigation) ConnectButtonR3KeyChanged(f func(key core.Qt__Key)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectButtonR3KeyChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepadKeyNavigation) DisconnectButtonR3KeyChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectButtonR3KeyChanged"})
}

func (ptr *QGamepadKeyNavigation) ButtonR3KeyChanged(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonR3KeyChanged", key})
}

func (ptr *QGamepadKeyNavigation) ButtonSelectKey() core.Qt__Key {

	return core.Qt__Key(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonSelectKey"}).(float64))
}

func (ptr *QGamepadKeyNavigation) ConnectButtonSelectKeyChanged(f func(key core.Qt__Key)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectButtonSelectKeyChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepadKeyNavigation) DisconnectButtonSelectKeyChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectButtonSelectKeyChanged"})
}

func (ptr *QGamepadKeyNavigation) ButtonSelectKeyChanged(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonSelectKeyChanged", key})
}

func (ptr *QGamepadKeyNavigation) ButtonStartKey() core.Qt__Key {

	return core.Qt__Key(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonStartKey"}).(float64))
}

func (ptr *QGamepadKeyNavigation) ConnectButtonStartKeyChanged(f func(key core.Qt__Key)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectButtonStartKeyChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepadKeyNavigation) DisconnectButtonStartKeyChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectButtonStartKeyChanged"})
}

func (ptr *QGamepadKeyNavigation) ButtonStartKeyChanged(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonStartKeyChanged", key})
}

func (ptr *QGamepadKeyNavigation) ButtonXKey() core.Qt__Key {

	return core.Qt__Key(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonXKey"}).(float64))
}

func (ptr *QGamepadKeyNavigation) ConnectButtonXKeyChanged(f func(key core.Qt__Key)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectButtonXKeyChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepadKeyNavigation) DisconnectButtonXKeyChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectButtonXKeyChanged"})
}

func (ptr *QGamepadKeyNavigation) ButtonXKeyChanged(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonXKeyChanged", key})
}

func (ptr *QGamepadKeyNavigation) ButtonYKey() core.Qt__Key {

	return core.Qt__Key(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonYKey"}).(float64))
}

func (ptr *QGamepadKeyNavigation) ConnectButtonYKeyChanged(f func(key core.Qt__Key)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectButtonYKeyChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepadKeyNavigation) DisconnectButtonYKeyChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectButtonYKeyChanged"})
}

func (ptr *QGamepadKeyNavigation) ButtonYKeyChanged(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ButtonYKeyChanged", key})
}

func (ptr *QGamepadKeyNavigation) DownKey() core.Qt__Key {

	return core.Qt__Key(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DownKey"}).(float64))
}

func (ptr *QGamepadKeyNavigation) ConnectDownKeyChanged(f func(key core.Qt__Key)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDownKeyChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepadKeyNavigation) DisconnectDownKeyChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDownKeyChanged"})
}

func (ptr *QGamepadKeyNavigation) DownKeyChanged(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DownKeyChanged", key})
}

func (ptr *QGamepadKeyNavigation) Gamepad() *QGamepad {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Gamepad"}).(*QGamepad)
}

func (ptr *QGamepadKeyNavigation) ConnectGamepadChanged(f func(gamepad *QGamepad)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectGamepadChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepadKeyNavigation) DisconnectGamepadChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectGamepadChanged"})
}

func (ptr *QGamepadKeyNavigation) GamepadChanged(gamepad QGamepad_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "GamepadChanged", gamepad})
}

func (ptr *QGamepadKeyNavigation) LeftKey() core.Qt__Key {

	return core.Qt__Key(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LeftKey"}).(float64))
}

func (ptr *QGamepadKeyNavigation) ConnectLeftKeyChanged(f func(key core.Qt__Key)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLeftKeyChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepadKeyNavigation) DisconnectLeftKeyChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLeftKeyChanged"})
}

func (ptr *QGamepadKeyNavigation) LeftKeyChanged(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LeftKeyChanged", key})
}

func (ptr *QGamepadKeyNavigation) RightKey() core.Qt__Key {

	return core.Qt__Key(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RightKey"}).(float64))
}

func (ptr *QGamepadKeyNavigation) ConnectRightKeyChanged(f func(key core.Qt__Key)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRightKeyChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepadKeyNavigation) DisconnectRightKeyChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRightKeyChanged"})
}

func (ptr *QGamepadKeyNavigation) RightKeyChanged(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RightKeyChanged", key})
}

func (ptr *QGamepadKeyNavigation) ConnectSetActive(f func(isActive bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetActive", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepadKeyNavigation) DisconnectSetActive() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetActive"})
}

func (ptr *QGamepadKeyNavigation) SetActive(isActive bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetActive", isActive})
}

func (ptr *QGamepadKeyNavigation) SetActiveDefault(isActive bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetActiveDefault", isActive})
}

func (ptr *QGamepadKeyNavigation) ConnectSetButtonAKey(f func(key core.Qt__Key)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetButtonAKey", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepadKeyNavigation) DisconnectSetButtonAKey() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetButtonAKey"})
}

func (ptr *QGamepadKeyNavigation) SetButtonAKey(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetButtonAKey", key})
}

func (ptr *QGamepadKeyNavigation) SetButtonAKeyDefault(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetButtonAKeyDefault", key})
}

func (ptr *QGamepadKeyNavigation) ConnectSetButtonBKey(f func(key core.Qt__Key)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetButtonBKey", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepadKeyNavigation) DisconnectSetButtonBKey() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetButtonBKey"})
}

func (ptr *QGamepadKeyNavigation) SetButtonBKey(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetButtonBKey", key})
}

func (ptr *QGamepadKeyNavigation) SetButtonBKeyDefault(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetButtonBKeyDefault", key})
}

func (ptr *QGamepadKeyNavigation) ConnectSetButtonGuideKey(f func(key core.Qt__Key)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetButtonGuideKey", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepadKeyNavigation) DisconnectSetButtonGuideKey() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetButtonGuideKey"})
}

func (ptr *QGamepadKeyNavigation) SetButtonGuideKey(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetButtonGuideKey", key})
}

func (ptr *QGamepadKeyNavigation) SetButtonGuideKeyDefault(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetButtonGuideKeyDefault", key})
}

func (ptr *QGamepadKeyNavigation) ConnectSetButtonL1Key(f func(key core.Qt__Key)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetButtonL1Key", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepadKeyNavigation) DisconnectSetButtonL1Key() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetButtonL1Key"})
}

func (ptr *QGamepadKeyNavigation) SetButtonL1Key(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetButtonL1Key", key})
}

func (ptr *QGamepadKeyNavigation) SetButtonL1KeyDefault(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetButtonL1KeyDefault", key})
}

func (ptr *QGamepadKeyNavigation) ConnectSetButtonL2Key(f func(key core.Qt__Key)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetButtonL2Key", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepadKeyNavigation) DisconnectSetButtonL2Key() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetButtonL2Key"})
}

func (ptr *QGamepadKeyNavigation) SetButtonL2Key(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetButtonL2Key", key})
}

func (ptr *QGamepadKeyNavigation) SetButtonL2KeyDefault(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetButtonL2KeyDefault", key})
}

func (ptr *QGamepadKeyNavigation) ConnectSetButtonL3Key(f func(key core.Qt__Key)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetButtonL3Key", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepadKeyNavigation) DisconnectSetButtonL3Key() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetButtonL3Key"})
}

func (ptr *QGamepadKeyNavigation) SetButtonL3Key(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetButtonL3Key", key})
}

func (ptr *QGamepadKeyNavigation) SetButtonL3KeyDefault(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetButtonL3KeyDefault", key})
}

func (ptr *QGamepadKeyNavigation) ConnectSetButtonR1Key(f func(key core.Qt__Key)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetButtonR1Key", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepadKeyNavigation) DisconnectSetButtonR1Key() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetButtonR1Key"})
}

func (ptr *QGamepadKeyNavigation) SetButtonR1Key(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetButtonR1Key", key})
}

func (ptr *QGamepadKeyNavigation) SetButtonR1KeyDefault(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetButtonR1KeyDefault", key})
}

func (ptr *QGamepadKeyNavigation) ConnectSetButtonR2Key(f func(key core.Qt__Key)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetButtonR2Key", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepadKeyNavigation) DisconnectSetButtonR2Key() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetButtonR2Key"})
}

func (ptr *QGamepadKeyNavigation) SetButtonR2Key(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetButtonR2Key", key})
}

func (ptr *QGamepadKeyNavigation) SetButtonR2KeyDefault(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetButtonR2KeyDefault", key})
}

func (ptr *QGamepadKeyNavigation) ConnectSetButtonR3Key(f func(key core.Qt__Key)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetButtonR3Key", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepadKeyNavigation) DisconnectSetButtonR3Key() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetButtonR3Key"})
}

func (ptr *QGamepadKeyNavigation) SetButtonR3Key(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetButtonR3Key", key})
}

func (ptr *QGamepadKeyNavigation) SetButtonR3KeyDefault(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetButtonR3KeyDefault", key})
}

func (ptr *QGamepadKeyNavigation) ConnectSetButtonSelectKey(f func(key core.Qt__Key)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetButtonSelectKey", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepadKeyNavigation) DisconnectSetButtonSelectKey() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetButtonSelectKey"})
}

func (ptr *QGamepadKeyNavigation) SetButtonSelectKey(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetButtonSelectKey", key})
}

func (ptr *QGamepadKeyNavigation) SetButtonSelectKeyDefault(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetButtonSelectKeyDefault", key})
}

func (ptr *QGamepadKeyNavigation) ConnectSetButtonStartKey(f func(key core.Qt__Key)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetButtonStartKey", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepadKeyNavigation) DisconnectSetButtonStartKey() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetButtonStartKey"})
}

func (ptr *QGamepadKeyNavigation) SetButtonStartKey(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetButtonStartKey", key})
}

func (ptr *QGamepadKeyNavigation) SetButtonStartKeyDefault(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetButtonStartKeyDefault", key})
}

func (ptr *QGamepadKeyNavigation) ConnectSetButtonXKey(f func(key core.Qt__Key)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetButtonXKey", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepadKeyNavigation) DisconnectSetButtonXKey() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetButtonXKey"})
}

func (ptr *QGamepadKeyNavigation) SetButtonXKey(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetButtonXKey", key})
}

func (ptr *QGamepadKeyNavigation) SetButtonXKeyDefault(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetButtonXKeyDefault", key})
}

func (ptr *QGamepadKeyNavigation) ConnectSetButtonYKey(f func(key core.Qt__Key)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetButtonYKey", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepadKeyNavigation) DisconnectSetButtonYKey() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetButtonYKey"})
}

func (ptr *QGamepadKeyNavigation) SetButtonYKey(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetButtonYKey", key})
}

func (ptr *QGamepadKeyNavigation) SetButtonYKeyDefault(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetButtonYKeyDefault", key})
}

func (ptr *QGamepadKeyNavigation) ConnectSetDownKey(f func(key core.Qt__Key)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetDownKey", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepadKeyNavigation) DisconnectSetDownKey() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetDownKey"})
}

func (ptr *QGamepadKeyNavigation) SetDownKey(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDownKey", key})
}

func (ptr *QGamepadKeyNavigation) SetDownKeyDefault(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDownKeyDefault", key})
}

func (ptr *QGamepadKeyNavigation) ConnectSetGamepad(f func(gamepad *QGamepad)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetGamepad", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepadKeyNavigation) DisconnectSetGamepad() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetGamepad"})
}

func (ptr *QGamepadKeyNavigation) SetGamepad(gamepad QGamepad_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetGamepad", gamepad})
}

func (ptr *QGamepadKeyNavigation) SetGamepadDefault(gamepad QGamepad_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetGamepadDefault", gamepad})
}

func (ptr *QGamepadKeyNavigation) ConnectSetLeftKey(f func(key core.Qt__Key)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetLeftKey", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepadKeyNavigation) DisconnectSetLeftKey() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetLeftKey"})
}

func (ptr *QGamepadKeyNavigation) SetLeftKey(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLeftKey", key})
}

func (ptr *QGamepadKeyNavigation) SetLeftKeyDefault(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLeftKeyDefault", key})
}

func (ptr *QGamepadKeyNavigation) ConnectSetRightKey(f func(key core.Qt__Key)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetRightKey", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepadKeyNavigation) DisconnectSetRightKey() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetRightKey"})
}

func (ptr *QGamepadKeyNavigation) SetRightKey(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRightKey", key})
}

func (ptr *QGamepadKeyNavigation) SetRightKeyDefault(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRightKeyDefault", key})
}

func (ptr *QGamepadKeyNavigation) ConnectSetUpKey(f func(key core.Qt__Key)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetUpKey", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepadKeyNavigation) DisconnectSetUpKey() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetUpKey"})
}

func (ptr *QGamepadKeyNavigation) SetUpKey(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetUpKey", key})
}

func (ptr *QGamepadKeyNavigation) SetUpKeyDefault(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetUpKeyDefault", key})
}

func (ptr *QGamepadKeyNavigation) UpKey() core.Qt__Key {

	return core.Qt__Key(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpKey"}).(float64))
}

func (ptr *QGamepadKeyNavigation) ConnectUpKeyChanged(f func(key core.Qt__Key)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectUpKeyChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepadKeyNavigation) DisconnectUpKeyChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectUpKeyChanged"})
}

func (ptr *QGamepadKeyNavigation) UpKeyChanged(key core.Qt__Key) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpKeyChanged", key})
}

func (ptr *QGamepadKeyNavigation) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QGamepadKeyNavigation) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QGamepadKeyNavigation) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QGamepadKeyNavigation) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QGamepadKeyNavigation) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QGamepadKeyNavigation) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QGamepadKeyNavigation) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QGamepadKeyNavigation) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QGamepadKeyNavigation) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QGamepadKeyNavigation) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QGamepadKeyNavigation) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QGamepadKeyNavigation) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QGamepadKeyNavigation) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QGamepadKeyNavigation) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QGamepadKeyNavigation) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QGamepadKeyNavigation) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QGamepadKeyNavigation) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QGamepadKeyNavigation) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QGamepadKeyNavigation) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QGamepadKeyNavigation) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QGamepadKeyNavigation) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
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

func (n *QGamepadManager) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QGamepadManager) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQGamepadManagerFromPointer(ptr unsafe.Pointer) (n *QGamepadManager) {
	n = new(QGamepadManager)
	n.InitFromInternal(uintptr(ptr), "gamepad.QGamepadManager")
	return
}

func (ptr *QGamepadManager) DestroyQGamepadManager() {
}

func (ptr *QGamepadManager) ConnectedGamepads() []int {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectedGamepads"}).([]int)
}

func (ptr *QGamepadManager) ConnectConnectedGamepadsChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectConnectedGamepadsChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepadManager) DisconnectConnectedGamepadsChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectConnectedGamepadsChanged"})
}

func (ptr *QGamepadManager) ConnectedGamepadsChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectedGamepadsChanged"})
}

func (ptr *QGamepadManager) GamepadName(deviceId int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "GamepadName", deviceId}).(string)
}

func QGamepadManager_Instance() *QGamepadManager {

	return internal.CallLocalFunction([]interface{}{"", "", "gamepad.QGamepadManager_Instance", ""}).(*QGamepadManager)
}

func (ptr *QGamepadManager) Instance() *QGamepadManager {

	return internal.CallLocalFunction([]interface{}{"", "", "gamepad.QGamepadManager_Instance", ""}).(*QGamepadManager)
}

func (ptr *QGamepadManager) ConnectIsConfigurationNeeded(f func(deviceId int) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectIsConfigurationNeeded", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepadManager) DisconnectIsConfigurationNeeded() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectIsConfigurationNeeded"})
}

func (ptr *QGamepadManager) IsConfigurationNeeded(deviceId int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsConfigurationNeeded", deviceId}).(bool)
}

func (ptr *QGamepadManager) IsConfigurationNeededDefault(deviceId int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsConfigurationNeededDefault", deviceId}).(bool)
}

func (ptr *QGamepadManager) IsGamepadConnected(deviceId int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsGamepadConnected", deviceId}).(bool)
}

func (ptr *QGamepadManager) ConnectResetConfiguration(f func(deviceId int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectResetConfiguration", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepadManager) DisconnectResetConfiguration() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectResetConfiguration"})
}

func (ptr *QGamepadManager) ResetConfiguration(deviceId int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResetConfiguration", deviceId})
}

func (ptr *QGamepadManager) ResetConfigurationDefault(deviceId int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResetConfigurationDefault", deviceId})
}

func (ptr *QGamepadManager) ConnectSetSettingsFile(f func(file string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetSettingsFile", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGamepadManager) DisconnectSetSettingsFile() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetSettingsFile"})
}

func (ptr *QGamepadManager) SetSettingsFile(file string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSettingsFile", file})
}

func (ptr *QGamepadManager) SetSettingsFileDefault(file string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSettingsFileDefault", file})
}

func (ptr *QGamepadManager) __connectedGamepads_atList(i int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__connectedGamepads_atList", i}).(float64))
}

func (ptr *QGamepadManager) __connectedGamepads_setList(i int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__connectedGamepads_setList", i})
}

func (ptr *QGamepadManager) __connectedGamepads_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__connectedGamepads_newList"}).(unsafe.Pointer)
}

func (ptr *QGamepadManager) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QGamepadManager) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QGamepadManager) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QGamepadManager) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QGamepadManager) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QGamepadManager) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QGamepadManager) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QGamepadManager) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QGamepadManager) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QGamepadManager) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QGamepadManager) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QGamepadManager) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QGamepadManager) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QGamepadManager) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QGamepadManager) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QGamepadManager) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QGamepadManager) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QGamepadManager) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QGamepadManager) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QGamepadManager) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QGamepadManager) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

func init() {
	internal.ConstructorTable["gamepad.QGamepad"] = NewQGamepadFromPointer
	internal.ConstructorTable["gamepad.QGamepadKeyNavigation"] = NewQGamepadKeyNavigationFromPointer
	internal.ConstructorTable["gamepad.QGamepadManager"] = NewQGamepadManagerFromPointer
}
