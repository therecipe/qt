// +build !minimal

package bluetooth

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/internal"
	"github.com/therecipe/qt/network"
	"strings"
	"unsafe"
)

func unpackStringList(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}
	return strings.Split(s, "¡¦!")
}

type QBluetooth struct {
	internal.Internal
}

type QBluetooth_ITF interface {
	QBluetooth_PTR() *QBluetooth
}

func (ptr *QBluetooth) QBluetooth_PTR() *QBluetooth {
	return ptr
}

func (ptr *QBluetooth) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QBluetooth) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQBluetooth(ptr QBluetooth_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetooth_PTR().Pointer()
	}
	return nil
}

func (n *QBluetooth) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQBluetoothFromPointer(ptr unsafe.Pointer) (n *QBluetooth) {
	n = new(QBluetooth)
	n.InitFromInternal(uintptr(ptr), "bluetooth.QBluetooth")
	return
}

func (ptr *QBluetooth) DestroyQBluetooth() {
}

//go:generate stringer -type=QBluetooth__AttAccessConstraint
//QBluetooth::AttAccessConstraint
type QBluetooth__AttAccessConstraint int64

const (
	QBluetooth__AttAuthorizationRequired  QBluetooth__AttAccessConstraint = QBluetooth__AttAccessConstraint(0x1)
	QBluetooth__AttAuthenticationRequired QBluetooth__AttAccessConstraint = QBluetooth__AttAccessConstraint(0x2)
	QBluetooth__AttEncryptionRequired     QBluetooth__AttAccessConstraint = QBluetooth__AttAccessConstraint(0x4)
)

//go:generate stringer -type=QBluetooth__Security
//QBluetooth::Security
type QBluetooth__Security int64

const (
	QBluetooth__NoSecurity     QBluetooth__Security = QBluetooth__Security(0x00)
	QBluetooth__Authorization  QBluetooth__Security = QBluetooth__Security(0x01)
	QBluetooth__Authentication QBluetooth__Security = QBluetooth__Security(0x02)
	QBluetooth__Encryption     QBluetooth__Security = QBluetooth__Security(0x04)
	QBluetooth__Secure         QBluetooth__Security = QBluetooth__Security(0x08)
)

type QBluetoothAddress struct {
	internal.Internal
}

type QBluetoothAddress_ITF interface {
	QBluetoothAddress_PTR() *QBluetoothAddress
}

func (ptr *QBluetoothAddress) QBluetoothAddress_PTR() *QBluetoothAddress {
	return ptr
}

func (ptr *QBluetoothAddress) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QBluetoothAddress) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQBluetoothAddress(ptr QBluetoothAddress_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothAddress_PTR().Pointer()
	}
	return nil
}

func (n *QBluetoothAddress) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQBluetoothAddressFromPointer(ptr unsafe.Pointer) (n *QBluetoothAddress) {
	n = new(QBluetoothAddress)
	n.InitFromInternal(uintptr(ptr), "bluetooth.QBluetoothAddress")
	return
}
func NewQBluetoothAddress() *QBluetoothAddress {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQBluetoothAddress", ""}).(*QBluetoothAddress)
}

func NewQBluetoothAddress2(address uint64) *QBluetoothAddress {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQBluetoothAddress2", "", address}).(*QBluetoothAddress)
}

func NewQBluetoothAddress3(address string) *QBluetoothAddress {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQBluetoothAddress3", "", address}).(*QBluetoothAddress)
}

func NewQBluetoothAddress4(other QBluetoothAddress_ITF) *QBluetoothAddress {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQBluetoothAddress4", "", other}).(*QBluetoothAddress)
}

func (ptr *QBluetoothAddress) Clear() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Clear"})
}

func (ptr *QBluetoothAddress) IsNull() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsNull"}).(bool)
}

func (ptr *QBluetoothAddress) ToString() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToString"}).(string)
}

func (ptr *QBluetoothAddress) ToUInt64() uint64 {

	return uint64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToUInt64"}).(float64))
}

func (ptr *QBluetoothAddress) DestroyQBluetoothAddress() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQBluetoothAddress"})
}

type QBluetoothDeviceDiscoveryAgent struct {
	core.QObject
}

type QBluetoothDeviceDiscoveryAgent_ITF interface {
	core.QObject_ITF
	QBluetoothDeviceDiscoveryAgent_PTR() *QBluetoothDeviceDiscoveryAgent
}

func (ptr *QBluetoothDeviceDiscoveryAgent) QBluetoothDeviceDiscoveryAgent_PTR() *QBluetoothDeviceDiscoveryAgent {
	return ptr
}

func (ptr *QBluetoothDeviceDiscoveryAgent) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QBluetoothDeviceDiscoveryAgent) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQBluetoothDeviceDiscoveryAgent(ptr QBluetoothDeviceDiscoveryAgent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothDeviceDiscoveryAgent_PTR().Pointer()
	}
	return nil
}

func (n *QBluetoothDeviceDiscoveryAgent) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QBluetoothDeviceDiscoveryAgent) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr unsafe.Pointer) (n *QBluetoothDeviceDiscoveryAgent) {
	n = new(QBluetoothDeviceDiscoveryAgent)
	n.InitFromInternal(uintptr(ptr), "bluetooth.QBluetoothDeviceDiscoveryAgent")
	return
}

//go:generate stringer -type=QBluetoothDeviceDiscoveryAgent__Error
//QBluetoothDeviceDiscoveryAgent::Error
type QBluetoothDeviceDiscoveryAgent__Error int64

const (
	QBluetoothDeviceDiscoveryAgent__NoError                      QBluetoothDeviceDiscoveryAgent__Error = QBluetoothDeviceDiscoveryAgent__Error(0)
	QBluetoothDeviceDiscoveryAgent__InputOutputError             QBluetoothDeviceDiscoveryAgent__Error = QBluetoothDeviceDiscoveryAgent__Error(1)
	QBluetoothDeviceDiscoveryAgent__PoweredOffError              QBluetoothDeviceDiscoveryAgent__Error = QBluetoothDeviceDiscoveryAgent__Error(2)
	QBluetoothDeviceDiscoveryAgent__InvalidBluetoothAdapterError QBluetoothDeviceDiscoveryAgent__Error = QBluetoothDeviceDiscoveryAgent__Error(3)
	QBluetoothDeviceDiscoveryAgent__UnsupportedPlatformError     QBluetoothDeviceDiscoveryAgent__Error = QBluetoothDeviceDiscoveryAgent__Error(4)
	QBluetoothDeviceDiscoveryAgent__UnsupportedDiscoveryMethod   QBluetoothDeviceDiscoveryAgent__Error = QBluetoothDeviceDiscoveryAgent__Error(5)
	QBluetoothDeviceDiscoveryAgent__UnknownError                 QBluetoothDeviceDiscoveryAgent__Error = QBluetoothDeviceDiscoveryAgent__Error(100)
)

//go:generate stringer -type=QBluetoothDeviceDiscoveryAgent__InquiryType
//QBluetoothDeviceDiscoveryAgent::InquiryType
type QBluetoothDeviceDiscoveryAgent__InquiryType int64

const (
	QBluetoothDeviceDiscoveryAgent__GeneralUnlimitedInquiry QBluetoothDeviceDiscoveryAgent__InquiryType = QBluetoothDeviceDiscoveryAgent__InquiryType(0)
	QBluetoothDeviceDiscoveryAgent__LimitedInquiry          QBluetoothDeviceDiscoveryAgent__InquiryType = QBluetoothDeviceDiscoveryAgent__InquiryType(1)
)

//go:generate stringer -type=QBluetoothDeviceDiscoveryAgent__DiscoveryMethod
//QBluetoothDeviceDiscoveryAgent::DiscoveryMethod
type QBluetoothDeviceDiscoveryAgent__DiscoveryMethod int64

const (
	QBluetoothDeviceDiscoveryAgent__NoMethod        QBluetoothDeviceDiscoveryAgent__DiscoveryMethod = QBluetoothDeviceDiscoveryAgent__DiscoveryMethod(0x0)
	QBluetoothDeviceDiscoveryAgent__ClassicMethod   QBluetoothDeviceDiscoveryAgent__DiscoveryMethod = QBluetoothDeviceDiscoveryAgent__DiscoveryMethod(0x01)
	QBluetoothDeviceDiscoveryAgent__LowEnergyMethod QBluetoothDeviceDiscoveryAgent__DiscoveryMethod = QBluetoothDeviceDiscoveryAgent__DiscoveryMethod(0x02)
)

func NewQBluetoothDeviceDiscoveryAgent(parent core.QObject_ITF) *QBluetoothDeviceDiscoveryAgent {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQBluetoothDeviceDiscoveryAgent", "", parent}).(*QBluetoothDeviceDiscoveryAgent)
}

func NewQBluetoothDeviceDiscoveryAgent2(deviceAdapter QBluetoothAddress_ITF, parent core.QObject_ITF) *QBluetoothDeviceDiscoveryAgent {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQBluetoothDeviceDiscoveryAgent2", "", deviceAdapter, parent}).(*QBluetoothDeviceDiscoveryAgent)
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectCanceled(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCanceled", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectCanceled() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCanceled"})
}

func (ptr *QBluetoothDeviceDiscoveryAgent) Canceled() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Canceled"})
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectDeviceDiscovered(f func(info *QBluetoothDeviceInfo)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDeviceDiscovered", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectDeviceDiscovered() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDeviceDiscovered"})
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DeviceDiscovered(info QBluetoothDeviceInfo_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeviceDiscovered", info})
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectDeviceUpdated(f func(info *QBluetoothDeviceInfo, updatedFields QBluetoothDeviceInfo__Field)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDeviceUpdated", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectDeviceUpdated() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDeviceUpdated"})
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DeviceUpdated(info QBluetoothDeviceInfo_ITF, updatedFields QBluetoothDeviceInfo__Field) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeviceUpdated", info, updatedFields})
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DiscoveredDevices() []*QBluetoothDeviceInfo {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DiscoveredDevices"}).([]*QBluetoothDeviceInfo)
}

func (ptr *QBluetoothDeviceDiscoveryAgent) Error() QBluetoothDeviceDiscoveryAgent__Error {

	return QBluetoothDeviceDiscoveryAgent__Error(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error"}).(float64))
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectError2(f func(error QBluetoothDeviceDiscoveryAgent__Error)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectError2", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectError2() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectError2"})
}

func (ptr *QBluetoothDeviceDiscoveryAgent) Error2(error QBluetoothDeviceDiscoveryAgent__Error) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error2", error})
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ErrorString() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ErrorString"}).(string)
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectFinished(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFinished", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectFinished() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFinished"})
}

func (ptr *QBluetoothDeviceDiscoveryAgent) Finished() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Finished"})
}

func (ptr *QBluetoothDeviceDiscoveryAgent) InquiryType() QBluetoothDeviceDiscoveryAgent__InquiryType {

	return QBluetoothDeviceDiscoveryAgent__InquiryType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InquiryType"}).(float64))
}

func (ptr *QBluetoothDeviceDiscoveryAgent) IsActive() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsActive"}).(bool)
}

func (ptr *QBluetoothDeviceDiscoveryAgent) LowEnergyDiscoveryTimeout() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LowEnergyDiscoveryTimeout"}).(float64))
}

func (ptr *QBluetoothDeviceDiscoveryAgent) SetInquiryType(ty QBluetoothDeviceDiscoveryAgent__InquiryType) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetInquiryType", ty})
}

func (ptr *QBluetoothDeviceDiscoveryAgent) SetLowEnergyDiscoveryTimeout(timeout int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLowEnergyDiscoveryTimeout", timeout})
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectStart(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStart", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectStart() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStart"})
}

func (ptr *QBluetoothDeviceDiscoveryAgent) Start() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Start"})
}

func (ptr *QBluetoothDeviceDiscoveryAgent) StartDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StartDefault"})
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectStart2(f func(methods QBluetoothDeviceDiscoveryAgent__DiscoveryMethod)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStart2", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectStart2() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStart2"})
}

func (ptr *QBluetoothDeviceDiscoveryAgent) Start2(methods QBluetoothDeviceDiscoveryAgent__DiscoveryMethod) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Start2", methods})
}

func (ptr *QBluetoothDeviceDiscoveryAgent) Start2Default(methods QBluetoothDeviceDiscoveryAgent__DiscoveryMethod) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Start2Default", methods})
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectStop(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStop", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectStop() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStop"})
}

func (ptr *QBluetoothDeviceDiscoveryAgent) Stop() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Stop"})
}

func (ptr *QBluetoothDeviceDiscoveryAgent) StopDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StopDefault"})
}

func QBluetoothDeviceDiscoveryAgent_SupportedDiscoveryMethods() QBluetoothDeviceDiscoveryAgent__DiscoveryMethod {

	return QBluetoothDeviceDiscoveryAgent__DiscoveryMethod(internal.CallLocalFunction([]interface{}{"", "", "bluetooth.QBluetoothDeviceDiscoveryAgent_SupportedDiscoveryMethods", ""}).(float64))
}

func (ptr *QBluetoothDeviceDiscoveryAgent) SupportedDiscoveryMethods() QBluetoothDeviceDiscoveryAgent__DiscoveryMethod {

	return QBluetoothDeviceDiscoveryAgent__DiscoveryMethod(internal.CallLocalFunction([]interface{}{"", "", "bluetooth.QBluetoothDeviceDiscoveryAgent_SupportedDiscoveryMethods", ""}).(float64))
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectDestroyQBluetoothDeviceDiscoveryAgent(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQBluetoothDeviceDiscoveryAgent", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectDestroyQBluetoothDeviceDiscoveryAgent() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQBluetoothDeviceDiscoveryAgent"})
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DestroyQBluetoothDeviceDiscoveryAgent() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQBluetoothDeviceDiscoveryAgent"})
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DestroyQBluetoothDeviceDiscoveryAgentDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQBluetoothDeviceDiscoveryAgentDefault"})
}

func (ptr *QBluetoothDeviceDiscoveryAgent) __discoveredDevices_atList(i int) *QBluetoothDeviceInfo {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__discoveredDevices_atList", i}).(*QBluetoothDeviceInfo)
}

func (ptr *QBluetoothDeviceDiscoveryAgent) __discoveredDevices_setList(i QBluetoothDeviceInfo_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__discoveredDevices_setList", i})
}

func (ptr *QBluetoothDeviceDiscoveryAgent) __discoveredDevices_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__discoveredDevices_newList"}).(unsafe.Pointer)
}

func (ptr *QBluetoothDeviceDiscoveryAgent) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QBluetoothDeviceDiscoveryAgent) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QBluetoothDeviceDiscoveryAgent) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QBluetoothDeviceDiscoveryAgent) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QBluetoothDeviceDiscoveryAgent) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QBluetoothDeviceDiscoveryAgent) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QBluetoothDeviceDiscoveryAgent) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QBluetoothDeviceDiscoveryAgent) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QBluetoothDeviceDiscoveryAgent) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QBluetoothDeviceDiscoveryAgent) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QBluetoothDeviceDiscoveryAgent) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QBluetoothDeviceDiscoveryAgent) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QBluetoothDeviceDiscoveryAgent) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QBluetoothDeviceDiscoveryAgent) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QBluetoothDeviceDiscoveryAgent) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QBluetoothDeviceDiscoveryAgent) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QBluetoothDeviceDiscoveryAgent) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QBluetoothDeviceInfo struct {
	internal.Internal
}

type QBluetoothDeviceInfo_ITF interface {
	QBluetoothDeviceInfo_PTR() *QBluetoothDeviceInfo
}

func (ptr *QBluetoothDeviceInfo) QBluetoothDeviceInfo_PTR() *QBluetoothDeviceInfo {
	return ptr
}

func (ptr *QBluetoothDeviceInfo) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QBluetoothDeviceInfo) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQBluetoothDeviceInfo(ptr QBluetoothDeviceInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothDeviceInfo_PTR().Pointer()
	}
	return nil
}

func (n *QBluetoothDeviceInfo) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQBluetoothDeviceInfoFromPointer(ptr unsafe.Pointer) (n *QBluetoothDeviceInfo) {
	n = new(QBluetoothDeviceInfo)
	n.InitFromInternal(uintptr(ptr), "bluetooth.QBluetoothDeviceInfo")
	return
}

//go:generate stringer -type=QBluetoothDeviceInfo__MajorDeviceClass
//QBluetoothDeviceInfo::MajorDeviceClass
type QBluetoothDeviceInfo__MajorDeviceClass int64

const (
	QBluetoothDeviceInfo__MiscellaneousDevice QBluetoothDeviceInfo__MajorDeviceClass = QBluetoothDeviceInfo__MajorDeviceClass(0)
	QBluetoothDeviceInfo__ComputerDevice      QBluetoothDeviceInfo__MajorDeviceClass = QBluetoothDeviceInfo__MajorDeviceClass(1)
	QBluetoothDeviceInfo__PhoneDevice         QBluetoothDeviceInfo__MajorDeviceClass = QBluetoothDeviceInfo__MajorDeviceClass(2)
	QBluetoothDeviceInfo__LANAccessDevice     QBluetoothDeviceInfo__MajorDeviceClass = QBluetoothDeviceInfo__MajorDeviceClass(3)
	QBluetoothDeviceInfo__NetworkDevice       QBluetoothDeviceInfo__MajorDeviceClass = QBluetoothDeviceInfo__MajorDeviceClass(3)
	QBluetoothDeviceInfo__AudioVideoDevice    QBluetoothDeviceInfo__MajorDeviceClass = QBluetoothDeviceInfo__MajorDeviceClass(4)
	QBluetoothDeviceInfo__PeripheralDevice    QBluetoothDeviceInfo__MajorDeviceClass = QBluetoothDeviceInfo__MajorDeviceClass(5)
	QBluetoothDeviceInfo__ImagingDevice       QBluetoothDeviceInfo__MajorDeviceClass = QBluetoothDeviceInfo__MajorDeviceClass(6)
	QBluetoothDeviceInfo__WearableDevice      QBluetoothDeviceInfo__MajorDeviceClass = QBluetoothDeviceInfo__MajorDeviceClass(7)
	QBluetoothDeviceInfo__ToyDevice           QBluetoothDeviceInfo__MajorDeviceClass = QBluetoothDeviceInfo__MajorDeviceClass(8)
	QBluetoothDeviceInfo__HealthDevice        QBluetoothDeviceInfo__MajorDeviceClass = QBluetoothDeviceInfo__MajorDeviceClass(9)
	QBluetoothDeviceInfo__UncategorizedDevice QBluetoothDeviceInfo__MajorDeviceClass = QBluetoothDeviceInfo__MajorDeviceClass(31)
)

//go:generate stringer -type=QBluetoothDeviceInfo__MinorMiscellaneousClass
//QBluetoothDeviceInfo::MinorMiscellaneousClass
type QBluetoothDeviceInfo__MinorMiscellaneousClass int64

const (
	QBluetoothDeviceInfo__UncategorizedMiscellaneous QBluetoothDeviceInfo__MinorMiscellaneousClass = QBluetoothDeviceInfo__MinorMiscellaneousClass(0)
)

//go:generate stringer -type=QBluetoothDeviceInfo__MinorComputerClass
//QBluetoothDeviceInfo::MinorComputerClass
type QBluetoothDeviceInfo__MinorComputerClass int64

const (
	QBluetoothDeviceInfo__UncategorizedComputer     QBluetoothDeviceInfo__MinorComputerClass = QBluetoothDeviceInfo__MinorComputerClass(0)
	QBluetoothDeviceInfo__DesktopComputer           QBluetoothDeviceInfo__MinorComputerClass = QBluetoothDeviceInfo__MinorComputerClass(1)
	QBluetoothDeviceInfo__ServerComputer            QBluetoothDeviceInfo__MinorComputerClass = QBluetoothDeviceInfo__MinorComputerClass(2)
	QBluetoothDeviceInfo__LaptopComputer            QBluetoothDeviceInfo__MinorComputerClass = QBluetoothDeviceInfo__MinorComputerClass(3)
	QBluetoothDeviceInfo__HandheldClamShellComputer QBluetoothDeviceInfo__MinorComputerClass = QBluetoothDeviceInfo__MinorComputerClass(4)
	QBluetoothDeviceInfo__HandheldComputer          QBluetoothDeviceInfo__MinorComputerClass = QBluetoothDeviceInfo__MinorComputerClass(5)
	QBluetoothDeviceInfo__WearableComputer          QBluetoothDeviceInfo__MinorComputerClass = QBluetoothDeviceInfo__MinorComputerClass(6)
)

//go:generate stringer -type=QBluetoothDeviceInfo__MinorPhoneClass
//QBluetoothDeviceInfo::MinorPhoneClass
type QBluetoothDeviceInfo__MinorPhoneClass int64

const (
	QBluetoothDeviceInfo__UncategorizedPhone            QBluetoothDeviceInfo__MinorPhoneClass = QBluetoothDeviceInfo__MinorPhoneClass(0)
	QBluetoothDeviceInfo__CellularPhone                 QBluetoothDeviceInfo__MinorPhoneClass = QBluetoothDeviceInfo__MinorPhoneClass(1)
	QBluetoothDeviceInfo__CordlessPhone                 QBluetoothDeviceInfo__MinorPhoneClass = QBluetoothDeviceInfo__MinorPhoneClass(2)
	QBluetoothDeviceInfo__SmartPhone                    QBluetoothDeviceInfo__MinorPhoneClass = QBluetoothDeviceInfo__MinorPhoneClass(3)
	QBluetoothDeviceInfo__WiredModemOrVoiceGatewayPhone QBluetoothDeviceInfo__MinorPhoneClass = QBluetoothDeviceInfo__MinorPhoneClass(4)
	QBluetoothDeviceInfo__CommonIsdnAccessPhone         QBluetoothDeviceInfo__MinorPhoneClass = QBluetoothDeviceInfo__MinorPhoneClass(5)
)

//go:generate stringer -type=QBluetoothDeviceInfo__MinorNetworkClass
//QBluetoothDeviceInfo::MinorNetworkClass
type QBluetoothDeviceInfo__MinorNetworkClass int64

const (
	QBluetoothDeviceInfo__NetworkFullService     QBluetoothDeviceInfo__MinorNetworkClass = QBluetoothDeviceInfo__MinorNetworkClass(0x00)
	QBluetoothDeviceInfo__NetworkLoadFactorOne   QBluetoothDeviceInfo__MinorNetworkClass = QBluetoothDeviceInfo__MinorNetworkClass(0x08)
	QBluetoothDeviceInfo__NetworkLoadFactorTwo   QBluetoothDeviceInfo__MinorNetworkClass = QBluetoothDeviceInfo__MinorNetworkClass(0x10)
	QBluetoothDeviceInfo__NetworkLoadFactorThree QBluetoothDeviceInfo__MinorNetworkClass = QBluetoothDeviceInfo__MinorNetworkClass(0x18)
	QBluetoothDeviceInfo__NetworkLoadFactorFour  QBluetoothDeviceInfo__MinorNetworkClass = QBluetoothDeviceInfo__MinorNetworkClass(0x20)
	QBluetoothDeviceInfo__NetworkLoadFactorFive  QBluetoothDeviceInfo__MinorNetworkClass = QBluetoothDeviceInfo__MinorNetworkClass(0x28)
	QBluetoothDeviceInfo__NetworkLoadFactorSix   QBluetoothDeviceInfo__MinorNetworkClass = QBluetoothDeviceInfo__MinorNetworkClass(0x30)
	QBluetoothDeviceInfo__NetworkNoService       QBluetoothDeviceInfo__MinorNetworkClass = QBluetoothDeviceInfo__MinorNetworkClass(0x38)
)

//go:generate stringer -type=QBluetoothDeviceInfo__MinorAudioVideoClass
//QBluetoothDeviceInfo::MinorAudioVideoClass
type QBluetoothDeviceInfo__MinorAudioVideoClass int64

const (
	QBluetoothDeviceInfo__UncategorizedAudioVideoDevice QBluetoothDeviceInfo__MinorAudioVideoClass = QBluetoothDeviceInfo__MinorAudioVideoClass(0)
	QBluetoothDeviceInfo__WearableHeadsetDevice         QBluetoothDeviceInfo__MinorAudioVideoClass = QBluetoothDeviceInfo__MinorAudioVideoClass(1)
	QBluetoothDeviceInfo__HandsFreeDevice               QBluetoothDeviceInfo__MinorAudioVideoClass = QBluetoothDeviceInfo__MinorAudioVideoClass(2)
	QBluetoothDeviceInfo__Microphone                    QBluetoothDeviceInfo__MinorAudioVideoClass = QBluetoothDeviceInfo__MinorAudioVideoClass(4)
	QBluetoothDeviceInfo__Loudspeaker                   QBluetoothDeviceInfo__MinorAudioVideoClass = QBluetoothDeviceInfo__MinorAudioVideoClass(5)
	QBluetoothDeviceInfo__Headphones                    QBluetoothDeviceInfo__MinorAudioVideoClass = QBluetoothDeviceInfo__MinorAudioVideoClass(6)
	QBluetoothDeviceInfo__PortableAudioDevice           QBluetoothDeviceInfo__MinorAudioVideoClass = QBluetoothDeviceInfo__MinorAudioVideoClass(7)
	QBluetoothDeviceInfo__CarAudio                      QBluetoothDeviceInfo__MinorAudioVideoClass = QBluetoothDeviceInfo__MinorAudioVideoClass(8)
	QBluetoothDeviceInfo__SetTopBox                     QBluetoothDeviceInfo__MinorAudioVideoClass = QBluetoothDeviceInfo__MinorAudioVideoClass(9)
	QBluetoothDeviceInfo__HiFiAudioDevice               QBluetoothDeviceInfo__MinorAudioVideoClass = QBluetoothDeviceInfo__MinorAudioVideoClass(10)
	QBluetoothDeviceInfo__Vcr                           QBluetoothDeviceInfo__MinorAudioVideoClass = QBluetoothDeviceInfo__MinorAudioVideoClass(11)
	QBluetoothDeviceInfo__VideoCamera                   QBluetoothDeviceInfo__MinorAudioVideoClass = QBluetoothDeviceInfo__MinorAudioVideoClass(12)
	QBluetoothDeviceInfo__Camcorder                     QBluetoothDeviceInfo__MinorAudioVideoClass = QBluetoothDeviceInfo__MinorAudioVideoClass(13)
	QBluetoothDeviceInfo__VideoMonitor                  QBluetoothDeviceInfo__MinorAudioVideoClass = QBluetoothDeviceInfo__MinorAudioVideoClass(14)
	QBluetoothDeviceInfo__VideoDisplayAndLoudspeaker    QBluetoothDeviceInfo__MinorAudioVideoClass = QBluetoothDeviceInfo__MinorAudioVideoClass(15)
	QBluetoothDeviceInfo__VideoConferencing             QBluetoothDeviceInfo__MinorAudioVideoClass = QBluetoothDeviceInfo__MinorAudioVideoClass(16)
	QBluetoothDeviceInfo__GamingDevice                  QBluetoothDeviceInfo__MinorAudioVideoClass = QBluetoothDeviceInfo__MinorAudioVideoClass(18)
)

//go:generate stringer -type=QBluetoothDeviceInfo__MinorPeripheralClass
//QBluetoothDeviceInfo::MinorPeripheralClass
type QBluetoothDeviceInfo__MinorPeripheralClass int64

const (
	QBluetoothDeviceInfo__UncategorizedPeripheral              QBluetoothDeviceInfo__MinorPeripheralClass = QBluetoothDeviceInfo__MinorPeripheralClass(0)
	QBluetoothDeviceInfo__KeyboardPeripheral                   QBluetoothDeviceInfo__MinorPeripheralClass = QBluetoothDeviceInfo__MinorPeripheralClass(0x10)
	QBluetoothDeviceInfo__PointingDevicePeripheral             QBluetoothDeviceInfo__MinorPeripheralClass = QBluetoothDeviceInfo__MinorPeripheralClass(0x20)
	QBluetoothDeviceInfo__KeyboardWithPointingDevicePeripheral QBluetoothDeviceInfo__MinorPeripheralClass = QBluetoothDeviceInfo__MinorPeripheralClass(0x30)
	QBluetoothDeviceInfo__JoystickPeripheral                   QBluetoothDeviceInfo__MinorPeripheralClass = QBluetoothDeviceInfo__MinorPeripheralClass(0x01)
	QBluetoothDeviceInfo__GamepadPeripheral                    QBluetoothDeviceInfo__MinorPeripheralClass = QBluetoothDeviceInfo__MinorPeripheralClass(0x02)
	QBluetoothDeviceInfo__RemoteControlPeripheral              QBluetoothDeviceInfo__MinorPeripheralClass = QBluetoothDeviceInfo__MinorPeripheralClass(0x03)
	QBluetoothDeviceInfo__SensingDevicePeripheral              QBluetoothDeviceInfo__MinorPeripheralClass = QBluetoothDeviceInfo__MinorPeripheralClass(0x04)
	QBluetoothDeviceInfo__DigitizerTabletPeripheral            QBluetoothDeviceInfo__MinorPeripheralClass = QBluetoothDeviceInfo__MinorPeripheralClass(0x05)
	QBluetoothDeviceInfo__CardReaderPeripheral                 QBluetoothDeviceInfo__MinorPeripheralClass = QBluetoothDeviceInfo__MinorPeripheralClass(0x06)
)

//go:generate stringer -type=QBluetoothDeviceInfo__MinorImagingClass
//QBluetoothDeviceInfo::MinorImagingClass
type QBluetoothDeviceInfo__MinorImagingClass int64

const (
	QBluetoothDeviceInfo__UncategorizedImagingDevice QBluetoothDeviceInfo__MinorImagingClass = QBluetoothDeviceInfo__MinorImagingClass(0)
	QBluetoothDeviceInfo__ImageDisplay               QBluetoothDeviceInfo__MinorImagingClass = QBluetoothDeviceInfo__MinorImagingClass(0x04)
	QBluetoothDeviceInfo__ImageCamera                QBluetoothDeviceInfo__MinorImagingClass = QBluetoothDeviceInfo__MinorImagingClass(0x08)
	QBluetoothDeviceInfo__ImageScanner               QBluetoothDeviceInfo__MinorImagingClass = QBluetoothDeviceInfo__MinorImagingClass(0x10)
	QBluetoothDeviceInfo__ImagePrinter               QBluetoothDeviceInfo__MinorImagingClass = QBluetoothDeviceInfo__MinorImagingClass(0x20)
)

//go:generate stringer -type=QBluetoothDeviceInfo__MinorWearableClass
//QBluetoothDeviceInfo::MinorWearableClass
type QBluetoothDeviceInfo__MinorWearableClass int64

const (
	QBluetoothDeviceInfo__UncategorizedWearableDevice QBluetoothDeviceInfo__MinorWearableClass = QBluetoothDeviceInfo__MinorWearableClass(0)
	QBluetoothDeviceInfo__WearableWristWatch          QBluetoothDeviceInfo__MinorWearableClass = QBluetoothDeviceInfo__MinorWearableClass(1)
	QBluetoothDeviceInfo__WearablePager               QBluetoothDeviceInfo__MinorWearableClass = QBluetoothDeviceInfo__MinorWearableClass(2)
	QBluetoothDeviceInfo__WearableJacket              QBluetoothDeviceInfo__MinorWearableClass = QBluetoothDeviceInfo__MinorWearableClass(3)
	QBluetoothDeviceInfo__WearableHelmet              QBluetoothDeviceInfo__MinorWearableClass = QBluetoothDeviceInfo__MinorWearableClass(4)
	QBluetoothDeviceInfo__WearableGlasses             QBluetoothDeviceInfo__MinorWearableClass = QBluetoothDeviceInfo__MinorWearableClass(5)
)

//go:generate stringer -type=QBluetoothDeviceInfo__MinorToyClass
//QBluetoothDeviceInfo::MinorToyClass
type QBluetoothDeviceInfo__MinorToyClass int64

const (
	QBluetoothDeviceInfo__UncategorizedToy QBluetoothDeviceInfo__MinorToyClass = QBluetoothDeviceInfo__MinorToyClass(0)
	QBluetoothDeviceInfo__ToyRobot         QBluetoothDeviceInfo__MinorToyClass = QBluetoothDeviceInfo__MinorToyClass(1)
	QBluetoothDeviceInfo__ToyVehicle       QBluetoothDeviceInfo__MinorToyClass = QBluetoothDeviceInfo__MinorToyClass(2)
	QBluetoothDeviceInfo__ToyDoll          QBluetoothDeviceInfo__MinorToyClass = QBluetoothDeviceInfo__MinorToyClass(3)
	QBluetoothDeviceInfo__ToyController    QBluetoothDeviceInfo__MinorToyClass = QBluetoothDeviceInfo__MinorToyClass(4)
	QBluetoothDeviceInfo__ToyGame          QBluetoothDeviceInfo__MinorToyClass = QBluetoothDeviceInfo__MinorToyClass(5)
)

//go:generate stringer -type=QBluetoothDeviceInfo__MinorHealthClass
//QBluetoothDeviceInfo::MinorHealthClass
type QBluetoothDeviceInfo__MinorHealthClass int64

const (
	QBluetoothDeviceInfo__UncategorizedHealthDevice  QBluetoothDeviceInfo__MinorHealthClass = QBluetoothDeviceInfo__MinorHealthClass(0)
	QBluetoothDeviceInfo__HealthBloodPressureMonitor QBluetoothDeviceInfo__MinorHealthClass = QBluetoothDeviceInfo__MinorHealthClass(0x1)
	QBluetoothDeviceInfo__HealthThermometer          QBluetoothDeviceInfo__MinorHealthClass = QBluetoothDeviceInfo__MinorHealthClass(0x2)
	QBluetoothDeviceInfo__HealthWeightScale          QBluetoothDeviceInfo__MinorHealthClass = QBluetoothDeviceInfo__MinorHealthClass(0x3)
	QBluetoothDeviceInfo__HealthGlucoseMeter         QBluetoothDeviceInfo__MinorHealthClass = QBluetoothDeviceInfo__MinorHealthClass(0x4)
	QBluetoothDeviceInfo__HealthPulseOximeter        QBluetoothDeviceInfo__MinorHealthClass = QBluetoothDeviceInfo__MinorHealthClass(0x5)
	QBluetoothDeviceInfo__HealthDataDisplay          QBluetoothDeviceInfo__MinorHealthClass = QBluetoothDeviceInfo__MinorHealthClass(0x7)
	QBluetoothDeviceInfo__HealthStepCounter          QBluetoothDeviceInfo__MinorHealthClass = QBluetoothDeviceInfo__MinorHealthClass(0x8)
)

//go:generate stringer -type=QBluetoothDeviceInfo__ServiceClass
//QBluetoothDeviceInfo::ServiceClass
type QBluetoothDeviceInfo__ServiceClass int64

const (
	QBluetoothDeviceInfo__NoService             QBluetoothDeviceInfo__ServiceClass = QBluetoothDeviceInfo__ServiceClass(0x0000)
	QBluetoothDeviceInfo__PositioningService    QBluetoothDeviceInfo__ServiceClass = QBluetoothDeviceInfo__ServiceClass(0x0001)
	QBluetoothDeviceInfo__NetworkingService     QBluetoothDeviceInfo__ServiceClass = QBluetoothDeviceInfo__ServiceClass(0x0002)
	QBluetoothDeviceInfo__RenderingService      QBluetoothDeviceInfo__ServiceClass = QBluetoothDeviceInfo__ServiceClass(0x0004)
	QBluetoothDeviceInfo__CapturingService      QBluetoothDeviceInfo__ServiceClass = QBluetoothDeviceInfo__ServiceClass(0x0008)
	QBluetoothDeviceInfo__ObjectTransferService QBluetoothDeviceInfo__ServiceClass = QBluetoothDeviceInfo__ServiceClass(0x0010)
	QBluetoothDeviceInfo__AudioService          QBluetoothDeviceInfo__ServiceClass = QBluetoothDeviceInfo__ServiceClass(0x0020)
	QBluetoothDeviceInfo__TelephonyService      QBluetoothDeviceInfo__ServiceClass = QBluetoothDeviceInfo__ServiceClass(0x0040)
	QBluetoothDeviceInfo__InformationService    QBluetoothDeviceInfo__ServiceClass = QBluetoothDeviceInfo__ServiceClass(0x0080)
	QBluetoothDeviceInfo__AllServices           QBluetoothDeviceInfo__ServiceClass = QBluetoothDeviceInfo__ServiceClass(0x07ff)
)

//go:generate stringer -type=QBluetoothDeviceInfo__DataCompleteness
//QBluetoothDeviceInfo::DataCompleteness
type QBluetoothDeviceInfo__DataCompleteness int64

const (
	QBluetoothDeviceInfo__DataComplete    QBluetoothDeviceInfo__DataCompleteness = QBluetoothDeviceInfo__DataCompleteness(0)
	QBluetoothDeviceInfo__DataIncomplete  QBluetoothDeviceInfo__DataCompleteness = QBluetoothDeviceInfo__DataCompleteness(1)
	QBluetoothDeviceInfo__DataUnavailable QBluetoothDeviceInfo__DataCompleteness = QBluetoothDeviceInfo__DataCompleteness(2)
)

//go:generate stringer -type=QBluetoothDeviceInfo__Field
//QBluetoothDeviceInfo::Field
type QBluetoothDeviceInfo__Field int64

const (
	QBluetoothDeviceInfo__None             QBluetoothDeviceInfo__Field = QBluetoothDeviceInfo__Field(0x0000)
	QBluetoothDeviceInfo__RSSI             QBluetoothDeviceInfo__Field = QBluetoothDeviceInfo__Field(0x0001)
	QBluetoothDeviceInfo__ManufacturerData QBluetoothDeviceInfo__Field = QBluetoothDeviceInfo__Field(0x0002)
	QBluetoothDeviceInfo__All              QBluetoothDeviceInfo__Field = QBluetoothDeviceInfo__Field(0x7fff)
)

//go:generate stringer -type=QBluetoothDeviceInfo__CoreConfiguration
//QBluetoothDeviceInfo::CoreConfiguration
type QBluetoothDeviceInfo__CoreConfiguration int64

const (
	QBluetoothDeviceInfo__UnknownCoreConfiguration              QBluetoothDeviceInfo__CoreConfiguration = QBluetoothDeviceInfo__CoreConfiguration(0x0)
	QBluetoothDeviceInfo__LowEnergyCoreConfiguration            QBluetoothDeviceInfo__CoreConfiguration = QBluetoothDeviceInfo__CoreConfiguration(0x01)
	QBluetoothDeviceInfo__BaseRateCoreConfiguration             QBluetoothDeviceInfo__CoreConfiguration = QBluetoothDeviceInfo__CoreConfiguration(0x02)
	QBluetoothDeviceInfo__BaseRateAndLowEnergyCoreConfiguration QBluetoothDeviceInfo__CoreConfiguration = QBluetoothDeviceInfo__CoreConfiguration(0x03)
)

func NewQBluetoothDeviceInfo() *QBluetoothDeviceInfo {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQBluetoothDeviceInfo", ""}).(*QBluetoothDeviceInfo)
}

func NewQBluetoothDeviceInfo2(address QBluetoothAddress_ITF, name string, classOfDevice uint) *QBluetoothDeviceInfo {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQBluetoothDeviceInfo2", "", address, name, classOfDevice}).(*QBluetoothDeviceInfo)
}

func NewQBluetoothDeviceInfo3(uuid QBluetoothUuid_ITF, name string, classOfDevice uint) *QBluetoothDeviceInfo {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQBluetoothDeviceInfo3", "", uuid, name, classOfDevice}).(*QBluetoothDeviceInfo)
}

func NewQBluetoothDeviceInfo4(other QBluetoothDeviceInfo_ITF) *QBluetoothDeviceInfo {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQBluetoothDeviceInfo4", "", other}).(*QBluetoothDeviceInfo)
}

func (ptr *QBluetoothDeviceInfo) Address() *QBluetoothAddress {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Address"}).(*QBluetoothAddress)
}

func (ptr *QBluetoothDeviceInfo) CoreConfigurations() QBluetoothDeviceInfo__CoreConfiguration {

	return QBluetoothDeviceInfo__CoreConfiguration(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CoreConfigurations"}).(float64))
}

func (ptr *QBluetoothDeviceInfo) DeviceUuid() *QBluetoothUuid {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeviceUuid"}).(*QBluetoothUuid)
}

func (ptr *QBluetoothDeviceInfo) IsCached() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsCached"}).(bool)
}

func (ptr *QBluetoothDeviceInfo) IsValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsValid"}).(bool)
}

func (ptr *QBluetoothDeviceInfo) MajorDeviceClass() QBluetoothDeviceInfo__MajorDeviceClass {

	return QBluetoothDeviceInfo__MajorDeviceClass(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MajorDeviceClass"}).(float64))
}

func (ptr *QBluetoothDeviceInfo) ManufacturerIds() []uint16 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ManufacturerIds"}).([]uint16)
}

func (ptr *QBluetoothDeviceInfo) MinorDeviceClass() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinorDeviceClass"}).(string)
}

func (ptr *QBluetoothDeviceInfo) Name() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Name"}).(string)
}

func (ptr *QBluetoothDeviceInfo) Rssi() int16 {

	return int16(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Rssi"}).(float64))
}

func (ptr *QBluetoothDeviceInfo) ServiceClasses() QBluetoothDeviceInfo__ServiceClass {

	return QBluetoothDeviceInfo__ServiceClass(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ServiceClasses"}).(float64))
}

func (ptr *QBluetoothDeviceInfo) SetCached(cached bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetCached", cached})
}

func (ptr *QBluetoothDeviceInfo) SetCoreConfigurations(coreConfigs QBluetoothDeviceInfo__CoreConfiguration) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetCoreConfigurations", coreConfigs})
}

func (ptr *QBluetoothDeviceInfo) SetDeviceUuid(uuid QBluetoothUuid_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDeviceUuid", uuid})
}

func (ptr *QBluetoothDeviceInfo) SetManufacturerData(manufacturerId uint16, data core.QByteArray_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetManufacturerData", manufacturerId, data}).(bool)
}

func (ptr *QBluetoothDeviceInfo) SetRssi(sign int16) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRssi", sign})
}

func (ptr *QBluetoothDeviceInfo) SetServiceUuids2(uuids []*QBluetoothUuid) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetServiceUuids2", uuids})
}

func (ptr *QBluetoothDeviceInfo) DestroyQBluetoothDeviceInfo() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQBluetoothDeviceInfo"})
}

func (ptr *QBluetoothDeviceInfo) __manufacturerIds_atList(i int) uint16 {

	return uint16(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__manufacturerIds_atList", i}).(float64))
}

func (ptr *QBluetoothDeviceInfo) __manufacturerIds_setList(i uint16) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__manufacturerIds_setList", i})
}

func (ptr *QBluetoothDeviceInfo) __manufacturerIds_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__manufacturerIds_newList"}).(unsafe.Pointer)
}

func (ptr *QBluetoothDeviceInfo) __serviceUuids_atList(i int) *QBluetoothUuid {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__serviceUuids_atList", i}).(*QBluetoothUuid)
}

func (ptr *QBluetoothDeviceInfo) __serviceUuids_setList(i QBluetoothUuid_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__serviceUuids_setList", i})
}

func (ptr *QBluetoothDeviceInfo) __serviceUuids_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__serviceUuids_newList"}).(unsafe.Pointer)
}

func (ptr *QBluetoothDeviceInfo) __setServiceUuids_uuids_atList(i int) *QBluetoothUuid {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setServiceUuids_uuids_atList", i}).(*QBluetoothUuid)
}

func (ptr *QBluetoothDeviceInfo) __setServiceUuids_uuids_setList(i QBluetoothUuid_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setServiceUuids_uuids_setList", i})
}

func (ptr *QBluetoothDeviceInfo) __setServiceUuids_uuids_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setServiceUuids_uuids_newList"}).(unsafe.Pointer)
}

func (ptr *QBluetoothDeviceInfo) __setServiceUuids_uuids_atList2(i int) *QBluetoothUuid {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setServiceUuids_uuids_atList2", i}).(*QBluetoothUuid)
}

func (ptr *QBluetoothDeviceInfo) __setServiceUuids_uuids_setList2(i QBluetoothUuid_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setServiceUuids_uuids_setList2", i})
}

func (ptr *QBluetoothDeviceInfo) __setServiceUuids_uuids_newList2() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setServiceUuids_uuids_newList2"}).(unsafe.Pointer)
}

func (ptr *QBluetoothDeviceInfo) ____manufacturerData_keyList_atList2(i int) uint16 {

	return uint16(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____manufacturerData_keyList_atList2", i}).(float64))
}

func (ptr *QBluetoothDeviceInfo) ____manufacturerData_keyList_setList2(i uint16) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____manufacturerData_keyList_setList2", i})
}

func (ptr *QBluetoothDeviceInfo) ____manufacturerData_keyList_newList2() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____manufacturerData_keyList_newList2"}).(unsafe.Pointer)
}

type QBluetoothHostInfo struct {
	internal.Internal
}

type QBluetoothHostInfo_ITF interface {
	QBluetoothHostInfo_PTR() *QBluetoothHostInfo
}

func (ptr *QBluetoothHostInfo) QBluetoothHostInfo_PTR() *QBluetoothHostInfo {
	return ptr
}

func (ptr *QBluetoothHostInfo) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QBluetoothHostInfo) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQBluetoothHostInfo(ptr QBluetoothHostInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothHostInfo_PTR().Pointer()
	}
	return nil
}

func (n *QBluetoothHostInfo) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQBluetoothHostInfoFromPointer(ptr unsafe.Pointer) (n *QBluetoothHostInfo) {
	n = new(QBluetoothHostInfo)
	n.InitFromInternal(uintptr(ptr), "bluetooth.QBluetoothHostInfo")
	return
}
func NewQBluetoothHostInfo() *QBluetoothHostInfo {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQBluetoothHostInfo", ""}).(*QBluetoothHostInfo)
}

func NewQBluetoothHostInfo2(other QBluetoothHostInfo_ITF) *QBluetoothHostInfo {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQBluetoothHostInfo2", "", other}).(*QBluetoothHostInfo)
}

func (ptr *QBluetoothHostInfo) Address() *QBluetoothAddress {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Address"}).(*QBluetoothAddress)
}

func (ptr *QBluetoothHostInfo) Name() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Name"}).(string)
}

func (ptr *QBluetoothHostInfo) SetAddress(address QBluetoothAddress_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAddress", address})
}

func (ptr *QBluetoothHostInfo) SetName(name string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetName", name})
}

func (ptr *QBluetoothHostInfo) DestroyQBluetoothHostInfo() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQBluetoothHostInfo"})
}

type QBluetoothLocalDevice struct {
	core.QObject
}

type QBluetoothLocalDevice_ITF interface {
	core.QObject_ITF
	QBluetoothLocalDevice_PTR() *QBluetoothLocalDevice
}

func (ptr *QBluetoothLocalDevice) QBluetoothLocalDevice_PTR() *QBluetoothLocalDevice {
	return ptr
}

func (ptr *QBluetoothLocalDevice) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QBluetoothLocalDevice) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQBluetoothLocalDevice(ptr QBluetoothLocalDevice_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothLocalDevice_PTR().Pointer()
	}
	return nil
}

func (n *QBluetoothLocalDevice) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QBluetoothLocalDevice) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQBluetoothLocalDeviceFromPointer(ptr unsafe.Pointer) (n *QBluetoothLocalDevice) {
	n = new(QBluetoothLocalDevice)
	n.InitFromInternal(uintptr(ptr), "bluetooth.QBluetoothLocalDevice")
	return
}

//go:generate stringer -type=QBluetoothLocalDevice__Pairing
//QBluetoothLocalDevice::Pairing
type QBluetoothLocalDevice__Pairing int64

const (
	QBluetoothLocalDevice__Unpaired         QBluetoothLocalDevice__Pairing = QBluetoothLocalDevice__Pairing(0)
	QBluetoothLocalDevice__Paired           QBluetoothLocalDevice__Pairing = QBluetoothLocalDevice__Pairing(1)
	QBluetoothLocalDevice__AuthorizedPaired QBluetoothLocalDevice__Pairing = QBluetoothLocalDevice__Pairing(2)
)

//go:generate stringer -type=QBluetoothLocalDevice__HostMode
//QBluetoothLocalDevice::HostMode
type QBluetoothLocalDevice__HostMode int64

const (
	QBluetoothLocalDevice__HostPoweredOff                 QBluetoothLocalDevice__HostMode = QBluetoothLocalDevice__HostMode(0)
	QBluetoothLocalDevice__HostConnectable                QBluetoothLocalDevice__HostMode = QBluetoothLocalDevice__HostMode(1)
	QBluetoothLocalDevice__HostDiscoverable               QBluetoothLocalDevice__HostMode = QBluetoothLocalDevice__HostMode(2)
	QBluetoothLocalDevice__HostDiscoverableLimitedInquiry QBluetoothLocalDevice__HostMode = QBluetoothLocalDevice__HostMode(3)
)

//go:generate stringer -type=QBluetoothLocalDevice__Error
//QBluetoothLocalDevice::Error
type QBluetoothLocalDevice__Error int64

const (
	QBluetoothLocalDevice__NoError      QBluetoothLocalDevice__Error = QBluetoothLocalDevice__Error(0)
	QBluetoothLocalDevice__PairingError QBluetoothLocalDevice__Error = QBluetoothLocalDevice__Error(1)
	QBluetoothLocalDevice__UnknownError QBluetoothLocalDevice__Error = QBluetoothLocalDevice__Error(100)
)

func NewQBluetoothLocalDevice(parent core.QObject_ITF) *QBluetoothLocalDevice {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQBluetoothLocalDevice", "", parent}).(*QBluetoothLocalDevice)
}

func NewQBluetoothLocalDevice2(address QBluetoothAddress_ITF, parent core.QObject_ITF) *QBluetoothLocalDevice {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQBluetoothLocalDevice2", "", address, parent}).(*QBluetoothLocalDevice)
}

func (ptr *QBluetoothLocalDevice) Address() *QBluetoothAddress {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Address"}).(*QBluetoothAddress)
}

func QBluetoothLocalDevice_AllDevices() []*QBluetoothHostInfo {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.QBluetoothLocalDevice_AllDevices", ""}).([]*QBluetoothHostInfo)
}

func (ptr *QBluetoothLocalDevice) AllDevices() []*QBluetoothHostInfo {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.QBluetoothLocalDevice_AllDevices", ""}).([]*QBluetoothHostInfo)
}

func (ptr *QBluetoothLocalDevice) ConnectedDevices() []*QBluetoothAddress {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectedDevices"}).([]*QBluetoothAddress)
}

func (ptr *QBluetoothLocalDevice) ConnectDeviceConnected(f func(address *QBluetoothAddress)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDeviceConnected", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBluetoothLocalDevice) DisconnectDeviceConnected() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDeviceConnected"})
}

func (ptr *QBluetoothLocalDevice) DeviceConnected(address QBluetoothAddress_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeviceConnected", address})
}

func (ptr *QBluetoothLocalDevice) ConnectDeviceDisconnected(f func(address *QBluetoothAddress)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDeviceDisconnected", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBluetoothLocalDevice) DisconnectDeviceDisconnected() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDeviceDisconnected"})
}

func (ptr *QBluetoothLocalDevice) DeviceDisconnected(address QBluetoothAddress_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeviceDisconnected", address})
}

func (ptr *QBluetoothLocalDevice) ConnectError(f func(error QBluetoothLocalDevice__Error)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectError", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBluetoothLocalDevice) DisconnectError() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectError"})
}

func (ptr *QBluetoothLocalDevice) Error(error QBluetoothLocalDevice__Error) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error", error})
}

func (ptr *QBluetoothLocalDevice) HostMode() QBluetoothLocalDevice__HostMode {

	return QBluetoothLocalDevice__HostMode(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HostMode"}).(float64))
}

func (ptr *QBluetoothLocalDevice) ConnectHostModeStateChanged(f func(state QBluetoothLocalDevice__HostMode)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectHostModeStateChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBluetoothLocalDevice) DisconnectHostModeStateChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectHostModeStateChanged"})
}

func (ptr *QBluetoothLocalDevice) HostModeStateChanged(state QBluetoothLocalDevice__HostMode) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HostModeStateChanged", state})
}

func (ptr *QBluetoothLocalDevice) IsValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsValid"}).(bool)
}

func (ptr *QBluetoothLocalDevice) Name() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Name"}).(string)
}

func (ptr *QBluetoothLocalDevice) ConnectPairingConfirmation(f func(confirmation bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPairingConfirmation", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBluetoothLocalDevice) DisconnectPairingConfirmation() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPairingConfirmation"})
}

func (ptr *QBluetoothLocalDevice) PairingConfirmation(confirmation bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PairingConfirmation", confirmation})
}

func (ptr *QBluetoothLocalDevice) PairingConfirmationDefault(confirmation bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PairingConfirmationDefault", confirmation})
}

func (ptr *QBluetoothLocalDevice) ConnectPairingDisplayConfirmation(f func(address *QBluetoothAddress, pin string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPairingDisplayConfirmation", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBluetoothLocalDevice) DisconnectPairingDisplayConfirmation() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPairingDisplayConfirmation"})
}

func (ptr *QBluetoothLocalDevice) PairingDisplayConfirmation(address QBluetoothAddress_ITF, pin string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PairingDisplayConfirmation", address, pin})
}

func (ptr *QBluetoothLocalDevice) ConnectPairingDisplayPinCode(f func(address *QBluetoothAddress, pin string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPairingDisplayPinCode", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBluetoothLocalDevice) DisconnectPairingDisplayPinCode() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPairingDisplayPinCode"})
}

func (ptr *QBluetoothLocalDevice) PairingDisplayPinCode(address QBluetoothAddress_ITF, pin string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PairingDisplayPinCode", address, pin})
}

func (ptr *QBluetoothLocalDevice) ConnectPairingFinished(f func(address *QBluetoothAddress, pairing QBluetoothLocalDevice__Pairing)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPairingFinished", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBluetoothLocalDevice) DisconnectPairingFinished() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPairingFinished"})
}

func (ptr *QBluetoothLocalDevice) PairingFinished(address QBluetoothAddress_ITF, pairing QBluetoothLocalDevice__Pairing) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PairingFinished", address, pairing})
}

func (ptr *QBluetoothLocalDevice) PairingStatus(address QBluetoothAddress_ITF) QBluetoothLocalDevice__Pairing {

	return QBluetoothLocalDevice__Pairing(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PairingStatus", address}).(float64))
}

func (ptr *QBluetoothLocalDevice) PowerOn() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PowerOn"})
}

func (ptr *QBluetoothLocalDevice) RequestPairing(address QBluetoothAddress_ITF, pairing QBluetoothLocalDevice__Pairing) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RequestPairing", address, pairing})
}

func (ptr *QBluetoothLocalDevice) SetHostMode(mode QBluetoothLocalDevice__HostMode) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHostMode", mode})
}

func (ptr *QBluetoothLocalDevice) ConnectDestroyQBluetoothLocalDevice(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQBluetoothLocalDevice", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBluetoothLocalDevice) DisconnectDestroyQBluetoothLocalDevice() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQBluetoothLocalDevice"})
}

func (ptr *QBluetoothLocalDevice) DestroyQBluetoothLocalDevice() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQBluetoothLocalDevice"})
}

func (ptr *QBluetoothLocalDevice) DestroyQBluetoothLocalDeviceDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQBluetoothLocalDeviceDefault"})
}

func (ptr *QBluetoothLocalDevice) __allDevices_atList(i int) *QBluetoothHostInfo {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__allDevices_atList", i}).(*QBluetoothHostInfo)
}

func (ptr *QBluetoothLocalDevice) __allDevices_setList(i QBluetoothHostInfo_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__allDevices_setList", i})
}

func (ptr *QBluetoothLocalDevice) __allDevices_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__allDevices_newList"}).(unsafe.Pointer)
}

func (ptr *QBluetoothLocalDevice) __connectedDevices_atList(i int) *QBluetoothAddress {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__connectedDevices_atList", i}).(*QBluetoothAddress)
}

func (ptr *QBluetoothLocalDevice) __connectedDevices_setList(i QBluetoothAddress_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__connectedDevices_setList", i})
}

func (ptr *QBluetoothLocalDevice) __connectedDevices_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__connectedDevices_newList"}).(unsafe.Pointer)
}

func (ptr *QBluetoothLocalDevice) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QBluetoothLocalDevice) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QBluetoothLocalDevice) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QBluetoothLocalDevice) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QBluetoothLocalDevice) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QBluetoothLocalDevice) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QBluetoothLocalDevice) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QBluetoothLocalDevice) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QBluetoothLocalDevice) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QBluetoothLocalDevice) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QBluetoothLocalDevice) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QBluetoothLocalDevice) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QBluetoothLocalDevice) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QBluetoothLocalDevice) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QBluetoothLocalDevice) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QBluetoothLocalDevice) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QBluetoothLocalDevice) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QBluetoothLocalDevice) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QBluetoothLocalDevice) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QBluetoothLocalDevice) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QBluetoothLocalDevice) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QBluetoothServer struct {
	core.QObject
}

type QBluetoothServer_ITF interface {
	core.QObject_ITF
	QBluetoothServer_PTR() *QBluetoothServer
}

func (ptr *QBluetoothServer) QBluetoothServer_PTR() *QBluetoothServer {
	return ptr
}

func (ptr *QBluetoothServer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QBluetoothServer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQBluetoothServer(ptr QBluetoothServer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothServer_PTR().Pointer()
	}
	return nil
}

func (n *QBluetoothServer) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QBluetoothServer) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQBluetoothServerFromPointer(ptr unsafe.Pointer) (n *QBluetoothServer) {
	n = new(QBluetoothServer)
	n.InitFromInternal(uintptr(ptr), "bluetooth.QBluetoothServer")
	return
}

//go:generate stringer -type=QBluetoothServer__Error
//QBluetoothServer::Error
type QBluetoothServer__Error int64

const (
	QBluetoothServer__NoError                       QBluetoothServer__Error = QBluetoothServer__Error(0)
	QBluetoothServer__UnknownError                  QBluetoothServer__Error = QBluetoothServer__Error(1)
	QBluetoothServer__PoweredOffError               QBluetoothServer__Error = QBluetoothServer__Error(2)
	QBluetoothServer__InputOutputError              QBluetoothServer__Error = QBluetoothServer__Error(3)
	QBluetoothServer__ServiceAlreadyRegisteredError QBluetoothServer__Error = QBluetoothServer__Error(4)
	QBluetoothServer__UnsupportedProtocolError      QBluetoothServer__Error = QBluetoothServer__Error(5)
)

func NewQBluetoothServer(serverType QBluetoothServiceInfo__Protocol, parent core.QObject_ITF) *QBluetoothServer {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQBluetoothServer", "", serverType, parent}).(*QBluetoothServer)
}

func (ptr *QBluetoothServer) Close() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Close"})
}

func (ptr *QBluetoothServer) Error() QBluetoothServer__Error {

	return QBluetoothServer__Error(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error"}).(float64))
}

func (ptr *QBluetoothServer) ConnectError2(f func(error QBluetoothServer__Error)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectError2", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBluetoothServer) DisconnectError2() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectError2"})
}

func (ptr *QBluetoothServer) Error2(error QBluetoothServer__Error) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error2", error})
}

func (ptr *QBluetoothServer) HasPendingConnections() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasPendingConnections"}).(bool)
}

func (ptr *QBluetoothServer) IsListening() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsListening"}).(bool)
}

func (ptr *QBluetoothServer) Listen(address QBluetoothAddress_ITF, port uint16) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Listen", address, port}).(bool)
}

func (ptr *QBluetoothServer) Listen2(uuid QBluetoothUuid_ITF, serviceName string) *QBluetoothServiceInfo {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Listen2", uuid, serviceName}).(*QBluetoothServiceInfo)
}

func (ptr *QBluetoothServer) MaxPendingConnections() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MaxPendingConnections"}).(float64))
}

func (ptr *QBluetoothServer) ConnectNewConnection(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNewConnection", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBluetoothServer) DisconnectNewConnection() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNewConnection"})
}

func (ptr *QBluetoothServer) NewConnection() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NewConnection"})
}

func (ptr *QBluetoothServer) NextPendingConnection() *QBluetoothSocket {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NextPendingConnection"}).(*QBluetoothSocket)
}

func (ptr *QBluetoothServer) SecurityFlags() QBluetooth__Security {

	return QBluetooth__Security(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SecurityFlags"}).(float64))
}

func (ptr *QBluetoothServer) ServerAddress() *QBluetoothAddress {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ServerAddress"}).(*QBluetoothAddress)
}

func (ptr *QBluetoothServer) ServerPort() uint16 {

	return uint16(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ServerPort"}).(float64))
}

func (ptr *QBluetoothServer) ServerType() QBluetoothServiceInfo__Protocol {

	return QBluetoothServiceInfo__Protocol(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ServerType"}).(float64))
}

func (ptr *QBluetoothServer) SetMaxPendingConnections(numConnections int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMaxPendingConnections", numConnections})
}

func (ptr *QBluetoothServer) SetSecurityFlags(security QBluetooth__Security) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSecurityFlags", security})
}

func (ptr *QBluetoothServer) ConnectDestroyQBluetoothServer(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQBluetoothServer", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBluetoothServer) DisconnectDestroyQBluetoothServer() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQBluetoothServer"})
}

func (ptr *QBluetoothServer) DestroyQBluetoothServer() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQBluetoothServer"})
}

func (ptr *QBluetoothServer) DestroyQBluetoothServerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQBluetoothServerDefault"})
}

func (ptr *QBluetoothServer) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QBluetoothServer) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QBluetoothServer) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QBluetoothServer) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QBluetoothServer) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QBluetoothServer) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QBluetoothServer) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QBluetoothServer) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QBluetoothServer) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QBluetoothServer) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QBluetoothServer) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QBluetoothServer) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QBluetoothServer) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QBluetoothServer) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QBluetoothServer) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QBluetoothServer) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QBluetoothServer) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QBluetoothServer) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QBluetoothServer) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QBluetoothServer) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QBluetoothServer) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QBluetoothServiceDiscoveryAgent struct {
	core.QObject
}

type QBluetoothServiceDiscoveryAgent_ITF interface {
	core.QObject_ITF
	QBluetoothServiceDiscoveryAgent_PTR() *QBluetoothServiceDiscoveryAgent
}

func (ptr *QBluetoothServiceDiscoveryAgent) QBluetoothServiceDiscoveryAgent_PTR() *QBluetoothServiceDiscoveryAgent {
	return ptr
}

func (ptr *QBluetoothServiceDiscoveryAgent) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QBluetoothServiceDiscoveryAgent) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQBluetoothServiceDiscoveryAgent(ptr QBluetoothServiceDiscoveryAgent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothServiceDiscoveryAgent_PTR().Pointer()
	}
	return nil
}

func (n *QBluetoothServiceDiscoveryAgent) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QBluetoothServiceDiscoveryAgent) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQBluetoothServiceDiscoveryAgentFromPointer(ptr unsafe.Pointer) (n *QBluetoothServiceDiscoveryAgent) {
	n = new(QBluetoothServiceDiscoveryAgent)
	n.InitFromInternal(uintptr(ptr), "bluetooth.QBluetoothServiceDiscoveryAgent")
	return
}

//go:generate stringer -type=QBluetoothServiceDiscoveryAgent__Error
//QBluetoothServiceDiscoveryAgent::Error
type QBluetoothServiceDiscoveryAgent__Error int64

const (
	QBluetoothServiceDiscoveryAgent__NoError                      QBluetoothServiceDiscoveryAgent__Error = QBluetoothServiceDiscoveryAgent__Error(QBluetoothDeviceDiscoveryAgent__NoError)
	QBluetoothServiceDiscoveryAgent__InputOutputError             QBluetoothServiceDiscoveryAgent__Error = QBluetoothServiceDiscoveryAgent__Error(QBluetoothDeviceDiscoveryAgent__InputOutputError)
	QBluetoothServiceDiscoveryAgent__PoweredOffError              QBluetoothServiceDiscoveryAgent__Error = QBluetoothServiceDiscoveryAgent__Error(QBluetoothDeviceDiscoveryAgent__PoweredOffError)
	QBluetoothServiceDiscoveryAgent__InvalidBluetoothAdapterError QBluetoothServiceDiscoveryAgent__Error = QBluetoothServiceDiscoveryAgent__Error(QBluetoothDeviceDiscoveryAgent__InvalidBluetoothAdapterError)
	QBluetoothServiceDiscoveryAgent__UnknownError                 QBluetoothServiceDiscoveryAgent__Error = QBluetoothServiceDiscoveryAgent__Error(QBluetoothDeviceDiscoveryAgent__UnknownError)
)

//go:generate stringer -type=QBluetoothServiceDiscoveryAgent__DiscoveryMode
//QBluetoothServiceDiscoveryAgent::DiscoveryMode
type QBluetoothServiceDiscoveryAgent__DiscoveryMode int64

const (
	QBluetoothServiceDiscoveryAgent__MinimalDiscovery QBluetoothServiceDiscoveryAgent__DiscoveryMode = QBluetoothServiceDiscoveryAgent__DiscoveryMode(0)
	QBluetoothServiceDiscoveryAgent__FullDiscovery    QBluetoothServiceDiscoveryAgent__DiscoveryMode = QBluetoothServiceDiscoveryAgent__DiscoveryMode(1)
)

func NewQBluetoothServiceDiscoveryAgent(parent core.QObject_ITF) *QBluetoothServiceDiscoveryAgent {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQBluetoothServiceDiscoveryAgent", "", parent}).(*QBluetoothServiceDiscoveryAgent)
}

func NewQBluetoothServiceDiscoveryAgent2(deviceAdapter QBluetoothAddress_ITF, parent core.QObject_ITF) *QBluetoothServiceDiscoveryAgent {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQBluetoothServiceDiscoveryAgent2", "", deviceAdapter, parent}).(*QBluetoothServiceDiscoveryAgent)
}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectCanceled(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCanceled", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectCanceled() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCanceled"})
}

func (ptr *QBluetoothServiceDiscoveryAgent) Canceled() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Canceled"})
}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectClear(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectClear", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectClear() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectClear"})
}

func (ptr *QBluetoothServiceDiscoveryAgent) Clear() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Clear"})
}

func (ptr *QBluetoothServiceDiscoveryAgent) ClearDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ClearDefault"})
}

func (ptr *QBluetoothServiceDiscoveryAgent) DiscoveredServices() []*QBluetoothServiceInfo {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DiscoveredServices"}).([]*QBluetoothServiceInfo)
}

func (ptr *QBluetoothServiceDiscoveryAgent) Error() QBluetoothServiceDiscoveryAgent__Error {

	return QBluetoothServiceDiscoveryAgent__Error(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error"}).(float64))
}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectError2(f func(error QBluetoothServiceDiscoveryAgent__Error)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectError2", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectError2() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectError2"})
}

func (ptr *QBluetoothServiceDiscoveryAgent) Error2(error QBluetoothServiceDiscoveryAgent__Error) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error2", error})
}

func (ptr *QBluetoothServiceDiscoveryAgent) ErrorString() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ErrorString"}).(string)
}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectFinished(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFinished", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectFinished() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFinished"})
}

func (ptr *QBluetoothServiceDiscoveryAgent) Finished() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Finished"})
}

func (ptr *QBluetoothServiceDiscoveryAgent) IsActive() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsActive"}).(bool)
}

func (ptr *QBluetoothServiceDiscoveryAgent) RemoteAddress() *QBluetoothAddress {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoteAddress"}).(*QBluetoothAddress)
}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectServiceDiscovered(f func(info *QBluetoothServiceInfo)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectServiceDiscovered", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectServiceDiscovered() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectServiceDiscovered"})
}

func (ptr *QBluetoothServiceDiscoveryAgent) ServiceDiscovered(info QBluetoothServiceInfo_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ServiceDiscovered", info})
}

func (ptr *QBluetoothServiceDiscoveryAgent) SetRemoteAddress(address QBluetoothAddress_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRemoteAddress", address}).(bool)
}

func (ptr *QBluetoothServiceDiscoveryAgent) SetUuidFilter(uuids []*QBluetoothUuid) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetUuidFilter", uuids})
}

func (ptr *QBluetoothServiceDiscoveryAgent) SetUuidFilter2(uuid QBluetoothUuid_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetUuidFilter2", uuid})
}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectStart(f func(mode QBluetoothServiceDiscoveryAgent__DiscoveryMode)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStart", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectStart() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStart"})
}

func (ptr *QBluetoothServiceDiscoveryAgent) Start(mode QBluetoothServiceDiscoveryAgent__DiscoveryMode) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Start", mode})
}

func (ptr *QBluetoothServiceDiscoveryAgent) StartDefault(mode QBluetoothServiceDiscoveryAgent__DiscoveryMode) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StartDefault", mode})
}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectStop(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStop", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectStop() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStop"})
}

func (ptr *QBluetoothServiceDiscoveryAgent) Stop() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Stop"})
}

func (ptr *QBluetoothServiceDiscoveryAgent) StopDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StopDefault"})
}

func (ptr *QBluetoothServiceDiscoveryAgent) UuidFilter() []*QBluetoothUuid {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UuidFilter"}).([]*QBluetoothUuid)
}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectDestroyQBluetoothServiceDiscoveryAgent(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQBluetoothServiceDiscoveryAgent", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectDestroyQBluetoothServiceDiscoveryAgent() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQBluetoothServiceDiscoveryAgent"})
}

func (ptr *QBluetoothServiceDiscoveryAgent) DestroyQBluetoothServiceDiscoveryAgent() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQBluetoothServiceDiscoveryAgent"})
}

func (ptr *QBluetoothServiceDiscoveryAgent) DestroyQBluetoothServiceDiscoveryAgentDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQBluetoothServiceDiscoveryAgentDefault"})
}

func (ptr *QBluetoothServiceDiscoveryAgent) __discoveredServices_atList(i int) *QBluetoothServiceInfo {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__discoveredServices_atList", i}).(*QBluetoothServiceInfo)
}

func (ptr *QBluetoothServiceDiscoveryAgent) __discoveredServices_setList(i QBluetoothServiceInfo_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__discoveredServices_setList", i})
}

func (ptr *QBluetoothServiceDiscoveryAgent) __discoveredServices_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__discoveredServices_newList"}).(unsafe.Pointer)
}

func (ptr *QBluetoothServiceDiscoveryAgent) __setUuidFilter_uuids_atList(i int) *QBluetoothUuid {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setUuidFilter_uuids_atList", i}).(*QBluetoothUuid)
}

func (ptr *QBluetoothServiceDiscoveryAgent) __setUuidFilter_uuids_setList(i QBluetoothUuid_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setUuidFilter_uuids_setList", i})
}

func (ptr *QBluetoothServiceDiscoveryAgent) __setUuidFilter_uuids_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setUuidFilter_uuids_newList"}).(unsafe.Pointer)
}

func (ptr *QBluetoothServiceDiscoveryAgent) __uuidFilter_atList(i int) *QBluetoothUuid {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__uuidFilter_atList", i}).(*QBluetoothUuid)
}

func (ptr *QBluetoothServiceDiscoveryAgent) __uuidFilter_setList(i QBluetoothUuid_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__uuidFilter_setList", i})
}

func (ptr *QBluetoothServiceDiscoveryAgent) __uuidFilter_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__uuidFilter_newList"}).(unsafe.Pointer)
}

func (ptr *QBluetoothServiceDiscoveryAgent) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QBluetoothServiceDiscoveryAgent) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QBluetoothServiceDiscoveryAgent) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QBluetoothServiceDiscoveryAgent) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QBluetoothServiceDiscoveryAgent) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QBluetoothServiceDiscoveryAgent) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QBluetoothServiceDiscoveryAgent) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QBluetoothServiceDiscoveryAgent) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QBluetoothServiceDiscoveryAgent) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QBluetoothServiceDiscoveryAgent) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QBluetoothServiceDiscoveryAgent) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QBluetoothServiceDiscoveryAgent) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QBluetoothServiceDiscoveryAgent) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QBluetoothServiceDiscoveryAgent) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QBluetoothServiceDiscoveryAgent) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QBluetoothServiceDiscoveryAgent) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QBluetoothServiceDiscoveryAgent) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QBluetoothServiceDiscoveryAgent) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QBluetoothServiceDiscoveryAgent) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QBluetoothServiceInfo struct {
	internal.Internal
}

type QBluetoothServiceInfo_ITF interface {
	QBluetoothServiceInfo_PTR() *QBluetoothServiceInfo
}

func (ptr *QBluetoothServiceInfo) QBluetoothServiceInfo_PTR() *QBluetoothServiceInfo {
	return ptr
}

func (ptr *QBluetoothServiceInfo) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QBluetoothServiceInfo) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQBluetoothServiceInfo(ptr QBluetoothServiceInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothServiceInfo_PTR().Pointer()
	}
	return nil
}

func (n *QBluetoothServiceInfo) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQBluetoothServiceInfoFromPointer(ptr unsafe.Pointer) (n *QBluetoothServiceInfo) {
	n = new(QBluetoothServiceInfo)
	n.InitFromInternal(uintptr(ptr), "bluetooth.QBluetoothServiceInfo")
	return
}

//go:generate stringer -type=QBluetoothServiceInfo__AttributeId
//QBluetoothServiceInfo::AttributeId
type QBluetoothServiceInfo__AttributeId int64

var (
	QBluetoothServiceInfo__ServiceRecordHandle              QBluetoothServiceInfo__AttributeId = QBluetoothServiceInfo__AttributeId(0x0000)
	QBluetoothServiceInfo__ServiceClassIds                  QBluetoothServiceInfo__AttributeId = QBluetoothServiceInfo__AttributeId(0x0001)
	QBluetoothServiceInfo__ServiceRecordState               QBluetoothServiceInfo__AttributeId = QBluetoothServiceInfo__AttributeId(0x0002)
	QBluetoothServiceInfo__ServiceId                        QBluetoothServiceInfo__AttributeId = QBluetoothServiceInfo__AttributeId(0x0003)
	QBluetoothServiceInfo__ProtocolDescriptorList           QBluetoothServiceInfo__AttributeId = QBluetoothServiceInfo__AttributeId(0x0004)
	QBluetoothServiceInfo__BrowseGroupList                  QBluetoothServiceInfo__AttributeId = QBluetoothServiceInfo__AttributeId(0x0005)
	QBluetoothServiceInfo__LanguageBaseAttributeIdList      QBluetoothServiceInfo__AttributeId = QBluetoothServiceInfo__AttributeId(0x0006)
	QBluetoothServiceInfo__ServiceInfoTimeToLive            QBluetoothServiceInfo__AttributeId = QBluetoothServiceInfo__AttributeId(0x0007)
	QBluetoothServiceInfo__ServiceAvailability              QBluetoothServiceInfo__AttributeId = QBluetoothServiceInfo__AttributeId(0x0008)
	QBluetoothServiceInfo__BluetoothProfileDescriptorList   QBluetoothServiceInfo__AttributeId = QBluetoothServiceInfo__AttributeId(0x0009)
	QBluetoothServiceInfo__DocumentationUrl                 QBluetoothServiceInfo__AttributeId = QBluetoothServiceInfo__AttributeId(0x000A)
	QBluetoothServiceInfo__ClientExecutableUrl              QBluetoothServiceInfo__AttributeId = QBluetoothServiceInfo__AttributeId(0x000B)
	QBluetoothServiceInfo__IconUrl                          QBluetoothServiceInfo__AttributeId = QBluetoothServiceInfo__AttributeId(0x000C)
	QBluetoothServiceInfo__AdditionalProtocolDescriptorList QBluetoothServiceInfo__AttributeId = QBluetoothServiceInfo__AttributeId(0x000D)
	QBluetoothServiceInfo__PrimaryLanguageBase              QBluetoothServiceInfo__AttributeId = QBluetoothServiceInfo__AttributeId(0x0100)
	QBluetoothServiceInfo__ServiceName                      QBluetoothServiceInfo__AttributeId = QBluetoothServiceInfo__AttributeId(0)
	QBluetoothServiceInfo__ServiceDescription               QBluetoothServiceInfo__AttributeId = QBluetoothServiceInfo__AttributeId(0)
	QBluetoothServiceInfo__ServiceProvider                  QBluetoothServiceInfo__AttributeId = QBluetoothServiceInfo__AttributeId(0)
)

//go:generate stringer -type=QBluetoothServiceInfo__Protocol
//QBluetoothServiceInfo::Protocol
type QBluetoothServiceInfo__Protocol int64

const (
	QBluetoothServiceInfo__UnknownProtocol QBluetoothServiceInfo__Protocol = QBluetoothServiceInfo__Protocol(0)
	QBluetoothServiceInfo__L2capProtocol   QBluetoothServiceInfo__Protocol = QBluetoothServiceInfo__Protocol(1)
	QBluetoothServiceInfo__RfcommProtocol  QBluetoothServiceInfo__Protocol = QBluetoothServiceInfo__Protocol(2)
)

func NewQBluetoothServiceInfo() *QBluetoothServiceInfo {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQBluetoothServiceInfo", ""}).(*QBluetoothServiceInfo)
}

func NewQBluetoothServiceInfo2(other QBluetoothServiceInfo_ITF) *QBluetoothServiceInfo {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQBluetoothServiceInfo2", "", other}).(*QBluetoothServiceInfo)
}

func (ptr *QBluetoothServiceInfo) Attribute(attributeId uint16) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Attribute", attributeId}).(*core.QVariant)
}

func (ptr *QBluetoothServiceInfo) Contains(attributeId uint16) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Contains", attributeId}).(bool)
}

func (ptr *QBluetoothServiceInfo) Device() *QBluetoothDeviceInfo {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Device"}).(*QBluetoothDeviceInfo)
}

func (ptr *QBluetoothServiceInfo) IsComplete() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsComplete"}).(bool)
}

func (ptr *QBluetoothServiceInfo) IsRegistered() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsRegistered"}).(bool)
}

func (ptr *QBluetoothServiceInfo) IsValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsValid"}).(bool)
}

func (ptr *QBluetoothServiceInfo) ProtocolServiceMultiplexer() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ProtocolServiceMultiplexer"}).(float64))
}

func (ptr *QBluetoothServiceInfo) RegisterService(localAdapter QBluetoothAddress_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RegisterService", localAdapter}).(bool)
}

func (ptr *QBluetoothServiceInfo) RemoveAttribute(attributeId uint16) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveAttribute", attributeId})
}

func (ptr *QBluetoothServiceInfo) ServerChannel() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ServerChannel"}).(float64))
}

func (ptr *QBluetoothServiceInfo) ServiceAvailability() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ServiceAvailability"}).(string)
}

func (ptr *QBluetoothServiceInfo) ServiceClassUuids() []*QBluetoothUuid {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ServiceClassUuids"}).([]*QBluetoothUuid)
}

func (ptr *QBluetoothServiceInfo) ServiceDescription() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ServiceDescription"}).(string)
}

func (ptr *QBluetoothServiceInfo) ServiceName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ServiceName"}).(string)
}

func (ptr *QBluetoothServiceInfo) ServiceProvider() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ServiceProvider"}).(string)
}

func (ptr *QBluetoothServiceInfo) ServiceUuid() *QBluetoothUuid {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ServiceUuid"}).(*QBluetoothUuid)
}

func (ptr *QBluetoothServiceInfo) SetAttribute(attributeId uint16, value core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAttribute", attributeId, value})
}

func (ptr *QBluetoothServiceInfo) SetAttribute2(attributeId uint16, value QBluetoothUuid_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAttribute2", attributeId, value})
}

func (ptr *QBluetoothServiceInfo) SetDevice(device QBluetoothDeviceInfo_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDevice", device})
}

func (ptr *QBluetoothServiceInfo) SetServiceAvailability(availability string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetServiceAvailability", availability})
}

func (ptr *QBluetoothServiceInfo) SetServiceDescription(description string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetServiceDescription", description})
}

func (ptr *QBluetoothServiceInfo) SetServiceName(name string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetServiceName", name})
}

func (ptr *QBluetoothServiceInfo) SetServiceProvider(provider string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetServiceProvider", provider})
}

func (ptr *QBluetoothServiceInfo) SetServiceUuid(uuid QBluetoothUuid_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetServiceUuid", uuid})
}

func (ptr *QBluetoothServiceInfo) SocketProtocol() QBluetoothServiceInfo__Protocol {

	return QBluetoothServiceInfo__Protocol(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SocketProtocol"}).(float64))
}

func (ptr *QBluetoothServiceInfo) UnregisterService() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UnregisterService"}).(bool)
}

func (ptr *QBluetoothServiceInfo) DestroyQBluetoothServiceInfo() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQBluetoothServiceInfo"})
}

func (ptr *QBluetoothServiceInfo) __serviceClassUuids_atList(i int) *QBluetoothUuid {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__serviceClassUuids_atList", i}).(*QBluetoothUuid)
}

func (ptr *QBluetoothServiceInfo) __serviceClassUuids_setList(i QBluetoothUuid_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__serviceClassUuids_setList", i})
}

func (ptr *QBluetoothServiceInfo) __serviceClassUuids_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__serviceClassUuids_newList"}).(unsafe.Pointer)
}

type QBluetoothSocket struct {
	core.QIODevice
}

type QBluetoothSocket_ITF interface {
	core.QIODevice_ITF
	QBluetoothSocket_PTR() *QBluetoothSocket
}

func (ptr *QBluetoothSocket) QBluetoothSocket_PTR() *QBluetoothSocket {
	return ptr
}

func (ptr *QBluetoothSocket) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QIODevice_PTR().Pointer()
	}
	return nil
}

func (ptr *QBluetoothSocket) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QIODevice_PTR().SetPointer(p)
	}
}

func PointerFromQBluetoothSocket(ptr QBluetoothSocket_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothSocket_PTR().Pointer()
	}
	return nil
}

func (n *QBluetoothSocket) InitFromInternal(ptr uintptr, name string) {
	n.QIODevice_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QBluetoothSocket) ClassNameInternalF() string {
	return n.QIODevice_PTR().ClassNameInternalF()
}

func NewQBluetoothSocketFromPointer(ptr unsafe.Pointer) (n *QBluetoothSocket) {
	n = new(QBluetoothSocket)
	n.InitFromInternal(uintptr(ptr), "bluetooth.QBluetoothSocket")
	return
}

//go:generate stringer -type=QBluetoothSocket__SocketState
//QBluetoothSocket::SocketState
type QBluetoothSocket__SocketState int64

const (
	QBluetoothSocket__UnconnectedState   QBluetoothSocket__SocketState = QBluetoothSocket__SocketState(network.QAbstractSocket__UnconnectedState)
	QBluetoothSocket__ServiceLookupState QBluetoothSocket__SocketState = QBluetoothSocket__SocketState(network.QAbstractSocket__HostLookupState)
	QBluetoothSocket__ConnectingState    QBluetoothSocket__SocketState = QBluetoothSocket__SocketState(network.QAbstractSocket__ConnectingState)
	QBluetoothSocket__ConnectedState     QBluetoothSocket__SocketState = QBluetoothSocket__SocketState(network.QAbstractSocket__ConnectedState)
	QBluetoothSocket__BoundState         QBluetoothSocket__SocketState = QBluetoothSocket__SocketState(network.QAbstractSocket__BoundState)
	QBluetoothSocket__ClosingState       QBluetoothSocket__SocketState = QBluetoothSocket__SocketState(network.QAbstractSocket__ClosingState)
	QBluetoothSocket__ListeningState     QBluetoothSocket__SocketState = QBluetoothSocket__SocketState(network.QAbstractSocket__ListeningState)
)

//go:generate stringer -type=QBluetoothSocket__SocketError
//QBluetoothSocket::SocketError
type QBluetoothSocket__SocketError int64

const (
	QBluetoothSocket__NoSocketError            QBluetoothSocket__SocketError = QBluetoothSocket__SocketError(-2)
	QBluetoothSocket__UnknownSocketError       QBluetoothSocket__SocketError = QBluetoothSocket__SocketError(network.QAbstractSocket__UnknownSocketError)
	QBluetoothSocket__RemoteHostClosedError    QBluetoothSocket__SocketError = QBluetoothSocket__SocketError(network.QAbstractSocket__RemoteHostClosedError)
	QBluetoothSocket__HostNotFoundError        QBluetoothSocket__SocketError = QBluetoothSocket__SocketError(network.QAbstractSocket__HostNotFoundError)
	QBluetoothSocket__ServiceNotFoundError     QBluetoothSocket__SocketError = QBluetoothSocket__SocketError(network.QAbstractSocket__SocketAddressNotAvailableError)
	QBluetoothSocket__NetworkError             QBluetoothSocket__SocketError = QBluetoothSocket__SocketError(network.QAbstractSocket__NetworkError)
	QBluetoothSocket__UnsupportedProtocolError QBluetoothSocket__SocketError = QBluetoothSocket__SocketError(8)
	QBluetoothSocket__OperationError           QBluetoothSocket__SocketError = QBluetoothSocket__SocketError(network.QAbstractSocket__OperationError)
)

func NewQBluetoothSocket(socketType QBluetoothServiceInfo__Protocol, parent core.QObject_ITF) *QBluetoothSocket {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQBluetoothSocket", "", socketType, parent}).(*QBluetoothSocket)
}

func NewQBluetoothSocket2(parent core.QObject_ITF) *QBluetoothSocket {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQBluetoothSocket2", "", parent}).(*QBluetoothSocket)
}

func (ptr *QBluetoothSocket) Abort() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Abort"})
}

func (ptr *QBluetoothSocket) BytesAvailableDefault() int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BytesAvailableDefault"}).(float64))
}

func (ptr *QBluetoothSocket) BytesToWriteDefault() int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BytesToWriteDefault"}).(float64))
}

func (ptr *QBluetoothSocket) CanReadLineDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CanReadLineDefault"}).(bool)
}

func (ptr *QBluetoothSocket) CloseDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseDefault"})
}

func (ptr *QBluetoothSocket) ConnectToService(service QBluetoothServiceInfo_ITF, openMode core.QIODevice__OpenModeFlag) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectToService", service, openMode})
}

func (ptr *QBluetoothSocket) ConnectToService2(address QBluetoothAddress_ITF, uuid QBluetoothUuid_ITF, openMode core.QIODevice__OpenModeFlag) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectToService2", address, uuid, openMode})
}

func (ptr *QBluetoothSocket) ConnectToService3(address QBluetoothAddress_ITF, port uint16, openMode core.QIODevice__OpenModeFlag) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectToService3", address, port, openMode})
}

func (ptr *QBluetoothSocket) ConnectConnected(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectConnected", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBluetoothSocket) DisconnectConnected() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectConnected"})
}

func (ptr *QBluetoothSocket) Connected() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Connected"})
}

func (ptr *QBluetoothSocket) DisconnectFromService() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFromService"})
}

func (ptr *QBluetoothSocket) ConnectDisconnected(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDisconnected", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBluetoothSocket) DisconnectDisconnected() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDisconnected"})
}

func (ptr *QBluetoothSocket) Disconnected() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Disconnected"})
}

func (ptr *QBluetoothSocket) DoDeviceDiscovery(service QBluetoothServiceInfo_ITF, openMode core.QIODevice__OpenModeFlag) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DoDeviceDiscovery", service, openMode})
}

func (ptr *QBluetoothSocket) Error() QBluetoothSocket__SocketError {

	return QBluetoothSocket__SocketError(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error"}).(float64))
}

func (ptr *QBluetoothSocket) ConnectError2(f func(error QBluetoothSocket__SocketError)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectError2", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBluetoothSocket) DisconnectError2() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectError2"})
}

func (ptr *QBluetoothSocket) Error2(error QBluetoothSocket__SocketError) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error2", error})
}

func (ptr *QBluetoothSocket) IsSequentialDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsSequentialDefault"}).(bool)
}

func (ptr *QBluetoothSocket) LocalAddress() *QBluetoothAddress {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LocalAddress"}).(*QBluetoothAddress)
}

func (ptr *QBluetoothSocket) LocalName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LocalName"}).(string)
}

func (ptr *QBluetoothSocket) LocalPort() uint16 {

	return uint16(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LocalPort"}).(float64))
}

func (ptr *QBluetoothSocket) PeerAddress() *QBluetoothAddress {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PeerAddress"}).(*QBluetoothAddress)
}

func (ptr *QBluetoothSocket) PeerName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PeerName"}).(string)
}

func (ptr *QBluetoothSocket) PeerPort() uint16 {

	return uint16(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PeerPort"}).(float64))
}

func (ptr *QBluetoothSocket) PreferredSecurityFlags() QBluetooth__Security {

	return QBluetooth__Security(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PreferredSecurityFlags"}).(float64))
}

func (ptr *QBluetoothSocket) ConnectReadData(f func(data *string, maxSize int64) int64) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectReadData", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBluetoothSocket) DisconnectReadData() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectReadData"})
}

func (ptr *QBluetoothSocket) ReadData(data *string, maxSize int64) int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReadData", data, maxSize}).(float64))
}

func (ptr *QBluetoothSocket) ReadDataDefault(data *string, maxSize int64) int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReadDataDefault", data, maxSize}).(float64))
}

func (ptr *QBluetoothSocket) SetPreferredSecurityFlags(flags QBluetooth__Security) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPreferredSecurityFlags", flags})
}

func (ptr *QBluetoothSocket) SetSocketDescriptor(socketDescriptor int, socketType QBluetoothServiceInfo__Protocol, socketState QBluetoothSocket__SocketState, openMode core.QIODevice__OpenModeFlag) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSocketDescriptor", socketDescriptor, socketType, socketState, openMode}).(bool)
}

func (ptr *QBluetoothSocket) SetSocketError(error_ QBluetoothSocket__SocketError) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSocketError", error_})
}

func (ptr *QBluetoothSocket) SetSocketState(state QBluetoothSocket__SocketState) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSocketState", state})
}

func (ptr *QBluetoothSocket) SocketDescriptor() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SocketDescriptor"}).(float64))
}

func (ptr *QBluetoothSocket) SocketType() QBluetoothServiceInfo__Protocol {

	return QBluetoothServiceInfo__Protocol(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SocketType"}).(float64))
}

func (ptr *QBluetoothSocket) State() QBluetoothSocket__SocketState {

	return QBluetoothSocket__SocketState(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "State"}).(float64))
}

func (ptr *QBluetoothSocket) ConnectStateChanged(f func(state QBluetoothSocket__SocketState)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStateChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBluetoothSocket) DisconnectStateChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStateChanged"})
}

func (ptr *QBluetoothSocket) StateChanged(state QBluetoothSocket__SocketState) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StateChanged", state})
}

func (ptr *QBluetoothSocket) ConnectWriteData(f func(data []byte, maxSize int64) int64) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectWriteData", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBluetoothSocket) DisconnectWriteData() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectWriteData"})
}

func (ptr *QBluetoothSocket) WriteData(data []byte, maxSize int64) int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WriteData", data, maxSize}).(float64))
}

func (ptr *QBluetoothSocket) WriteDataDefault(data []byte, maxSize int64) int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WriteDataDefault", data, maxSize}).(float64))
}

func (ptr *QBluetoothSocket) ConnectDestroyQBluetoothSocket(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQBluetoothSocket", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBluetoothSocket) DisconnectDestroyQBluetoothSocket() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQBluetoothSocket"})
}

func (ptr *QBluetoothSocket) DestroyQBluetoothSocket() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQBluetoothSocket"})
}

func (ptr *QBluetoothSocket) DestroyQBluetoothSocketDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQBluetoothSocketDefault"})
}

func (ptr *QBluetoothSocket) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QBluetoothSocket) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QBluetoothSocket) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QBluetoothSocket) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QBluetoothSocket) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QBluetoothSocket) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QBluetoothSocket) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QBluetoothSocket) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QBluetoothSocket) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QBluetoothSocket) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QBluetoothSocket) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QBluetoothSocket) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QBluetoothSocket) AtEndDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AtEndDefault"}).(bool)
}

func (ptr *QBluetoothSocket) OpenDefault(mode core.QIODevice__OpenModeFlag) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OpenDefault", mode}).(bool)
}

func (ptr *QBluetoothSocket) PosDefault() int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PosDefault"}).(float64))
}

func (ptr *QBluetoothSocket) ReadLineDataDefault(data []byte, maxSize int64) int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReadLineDataDefault", data, maxSize}).(float64))
}

func (ptr *QBluetoothSocket) ResetDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResetDefault"}).(bool)
}

func (ptr *QBluetoothSocket) SeekDefault(pos int64) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SeekDefault", pos}).(bool)
}

func (ptr *QBluetoothSocket) SizeDefault() int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SizeDefault"}).(float64))
}

func (ptr *QBluetoothSocket) WaitForBytesWrittenDefault(msecs int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WaitForBytesWrittenDefault", msecs}).(bool)
}

func (ptr *QBluetoothSocket) WaitForReadyReadDefault(msecs int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WaitForReadyReadDefault", msecs}).(bool)
}

func (ptr *QBluetoothSocket) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QBluetoothSocket) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QBluetoothSocket) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QBluetoothSocket) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QBluetoothSocket) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QBluetoothSocket) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QBluetoothSocket) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QBluetoothSocket) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QBluetoothSocket) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QBluetoothTransferManager struct {
	core.QObject
}

type QBluetoothTransferManager_ITF interface {
	core.QObject_ITF
	QBluetoothTransferManager_PTR() *QBluetoothTransferManager
}

func (ptr *QBluetoothTransferManager) QBluetoothTransferManager_PTR() *QBluetoothTransferManager {
	return ptr
}

func (ptr *QBluetoothTransferManager) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QBluetoothTransferManager) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQBluetoothTransferManager(ptr QBluetoothTransferManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothTransferManager_PTR().Pointer()
	}
	return nil
}

func (n *QBluetoothTransferManager) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QBluetoothTransferManager) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQBluetoothTransferManagerFromPointer(ptr unsafe.Pointer) (n *QBluetoothTransferManager) {
	n = new(QBluetoothTransferManager)
	n.InitFromInternal(uintptr(ptr), "bluetooth.QBluetoothTransferManager")
	return
}
func NewQBluetoothTransferManager(parent core.QObject_ITF) *QBluetoothTransferManager {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQBluetoothTransferManager", "", parent}).(*QBluetoothTransferManager)
}

func (ptr *QBluetoothTransferManager) ConnectFinished(f func(reply *QBluetoothTransferReply)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFinished", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBluetoothTransferManager) DisconnectFinished() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFinished"})
}

func (ptr *QBluetoothTransferManager) Finished(reply QBluetoothTransferReply_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Finished", reply})
}

func (ptr *QBluetoothTransferManager) Put(request QBluetoothTransferRequest_ITF, data core.QIODevice_ITF) *QBluetoothTransferReply {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Put", request, data}).(*QBluetoothTransferReply)
}

func (ptr *QBluetoothTransferManager) ConnectDestroyQBluetoothTransferManager(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQBluetoothTransferManager", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBluetoothTransferManager) DisconnectDestroyQBluetoothTransferManager() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQBluetoothTransferManager"})
}

func (ptr *QBluetoothTransferManager) DestroyQBluetoothTransferManager() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQBluetoothTransferManager"})
}

func (ptr *QBluetoothTransferManager) DestroyQBluetoothTransferManagerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQBluetoothTransferManagerDefault"})
}

func (ptr *QBluetoothTransferManager) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QBluetoothTransferManager) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QBluetoothTransferManager) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QBluetoothTransferManager) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QBluetoothTransferManager) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QBluetoothTransferManager) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QBluetoothTransferManager) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QBluetoothTransferManager) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QBluetoothTransferManager) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QBluetoothTransferManager) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QBluetoothTransferManager) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QBluetoothTransferManager) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QBluetoothTransferManager) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QBluetoothTransferManager) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QBluetoothTransferManager) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QBluetoothTransferManager) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QBluetoothTransferManager) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QBluetoothTransferManager) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QBluetoothTransferManager) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QBluetoothTransferManager) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QBluetoothTransferManager) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QBluetoothTransferReply struct {
	core.QObject
}

type QBluetoothTransferReply_ITF interface {
	core.QObject_ITF
	QBluetoothTransferReply_PTR() *QBluetoothTransferReply
}

func (ptr *QBluetoothTransferReply) QBluetoothTransferReply_PTR() *QBluetoothTransferReply {
	return ptr
}

func (ptr *QBluetoothTransferReply) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QBluetoothTransferReply) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQBluetoothTransferReply(ptr QBluetoothTransferReply_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothTransferReply_PTR().Pointer()
	}
	return nil
}

func (n *QBluetoothTransferReply) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QBluetoothTransferReply) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQBluetoothTransferReplyFromPointer(ptr unsafe.Pointer) (n *QBluetoothTransferReply) {
	n = new(QBluetoothTransferReply)
	n.InitFromInternal(uintptr(ptr), "bluetooth.QBluetoothTransferReply")
	return
}

//go:generate stringer -type=QBluetoothTransferReply__TransferError
//QBluetoothTransferReply::TransferError
type QBluetoothTransferReply__TransferError int64

const (
	QBluetoothTransferReply__NoError                   QBluetoothTransferReply__TransferError = QBluetoothTransferReply__TransferError(0)
	QBluetoothTransferReply__UnknownError              QBluetoothTransferReply__TransferError = QBluetoothTransferReply__TransferError(1)
	QBluetoothTransferReply__FileNotFoundError         QBluetoothTransferReply__TransferError = QBluetoothTransferReply__TransferError(2)
	QBluetoothTransferReply__HostNotFoundError         QBluetoothTransferReply__TransferError = QBluetoothTransferReply__TransferError(3)
	QBluetoothTransferReply__UserCanceledTransferError QBluetoothTransferReply__TransferError = QBluetoothTransferReply__TransferError(4)
	QBluetoothTransferReply__IODeviceNotReadableError  QBluetoothTransferReply__TransferError = QBluetoothTransferReply__TransferError(5)
	QBluetoothTransferReply__ResourceBusyError         QBluetoothTransferReply__TransferError = QBluetoothTransferReply__TransferError(6)
	QBluetoothTransferReply__SessionError              QBluetoothTransferReply__TransferError = QBluetoothTransferReply__TransferError(7)
)

func NewQBluetoothTransferReply(parent core.QObject_ITF) *QBluetoothTransferReply {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQBluetoothTransferReply", "", parent}).(*QBluetoothTransferReply)
}

func (ptr *QBluetoothTransferReply) ConnectAbort(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAbort", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBluetoothTransferReply) DisconnectAbort() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAbort"})
}

func (ptr *QBluetoothTransferReply) Abort() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Abort"})
}

func (ptr *QBluetoothTransferReply) AbortDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AbortDefault"})
}

func (ptr *QBluetoothTransferReply) ConnectError(f func() QBluetoothTransferReply__TransferError) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectError", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBluetoothTransferReply) DisconnectError() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectError"})
}

func (ptr *QBluetoothTransferReply) Error() QBluetoothTransferReply__TransferError {

	return QBluetoothTransferReply__TransferError(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error"}).(float64))
}

func (ptr *QBluetoothTransferReply) ConnectError2(f func(errorType QBluetoothTransferReply__TransferError)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectError2", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBluetoothTransferReply) DisconnectError2() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectError2"})
}

func (ptr *QBluetoothTransferReply) Error2(errorType QBluetoothTransferReply__TransferError) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error2", errorType})
}

func (ptr *QBluetoothTransferReply) ConnectErrorString(f func() string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectErrorString", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBluetoothTransferReply) DisconnectErrorString() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectErrorString"})
}

func (ptr *QBluetoothTransferReply) ErrorString() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ErrorString"}).(string)
}

func (ptr *QBluetoothTransferReply) ConnectFinished(f func(reply *QBluetoothTransferReply)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFinished", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBluetoothTransferReply) DisconnectFinished() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFinished"})
}

func (ptr *QBluetoothTransferReply) Finished(reply QBluetoothTransferReply_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Finished", reply})
}

func (ptr *QBluetoothTransferReply) ConnectIsFinished(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectIsFinished", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBluetoothTransferReply) DisconnectIsFinished() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectIsFinished"})
}

func (ptr *QBluetoothTransferReply) IsFinished() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsFinished"}).(bool)
}

func (ptr *QBluetoothTransferReply) ConnectIsRunning(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectIsRunning", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBluetoothTransferReply) DisconnectIsRunning() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectIsRunning"})
}

func (ptr *QBluetoothTransferReply) IsRunning() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsRunning"}).(bool)
}

func (ptr *QBluetoothTransferReply) Manager() *QBluetoothTransferManager {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Manager"}).(*QBluetoothTransferManager)
}

func (ptr *QBluetoothTransferReply) Request() *QBluetoothTransferRequest {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Request"}).(*QBluetoothTransferRequest)
}

func (ptr *QBluetoothTransferReply) SetManager(manager QBluetoothTransferManager_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetManager", manager})
}

func (ptr *QBluetoothTransferReply) SetRequest(request QBluetoothTransferRequest_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRequest", request})
}

func (ptr *QBluetoothTransferReply) ConnectTransferProgress(f func(bytesTransferred int64, bytesTotal int64)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTransferProgress", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBluetoothTransferReply) DisconnectTransferProgress() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTransferProgress"})
}

func (ptr *QBluetoothTransferReply) TransferProgress(bytesTransferred int64, bytesTotal int64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TransferProgress", bytesTransferred, bytesTotal})
}

func (ptr *QBluetoothTransferReply) ConnectDestroyQBluetoothTransferReply(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQBluetoothTransferReply", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QBluetoothTransferReply) DisconnectDestroyQBluetoothTransferReply() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQBluetoothTransferReply"})
}

func (ptr *QBluetoothTransferReply) DestroyQBluetoothTransferReply() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQBluetoothTransferReply"})
}

func (ptr *QBluetoothTransferReply) DestroyQBluetoothTransferReplyDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQBluetoothTransferReplyDefault"})
}

func (ptr *QBluetoothTransferReply) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QBluetoothTransferReply) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QBluetoothTransferReply) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QBluetoothTransferReply) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QBluetoothTransferReply) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QBluetoothTransferReply) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QBluetoothTransferReply) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QBluetoothTransferReply) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QBluetoothTransferReply) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QBluetoothTransferReply) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QBluetoothTransferReply) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QBluetoothTransferReply) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QBluetoothTransferReply) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QBluetoothTransferReply) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QBluetoothTransferReply) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QBluetoothTransferReply) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QBluetoothTransferReply) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QBluetoothTransferReply) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QBluetoothTransferReply) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QBluetoothTransferReply) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QBluetoothTransferReply) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QBluetoothTransferRequest struct {
	internal.Internal
}

type QBluetoothTransferRequest_ITF interface {
	QBluetoothTransferRequest_PTR() *QBluetoothTransferRequest
}

func (ptr *QBluetoothTransferRequest) QBluetoothTransferRequest_PTR() *QBluetoothTransferRequest {
	return ptr
}

func (ptr *QBluetoothTransferRequest) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QBluetoothTransferRequest) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQBluetoothTransferRequest(ptr QBluetoothTransferRequest_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothTransferRequest_PTR().Pointer()
	}
	return nil
}

func (n *QBluetoothTransferRequest) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQBluetoothTransferRequestFromPointer(ptr unsafe.Pointer) (n *QBluetoothTransferRequest) {
	n = new(QBluetoothTransferRequest)
	n.InitFromInternal(uintptr(ptr), "bluetooth.QBluetoothTransferRequest")
	return
}

//go:generate stringer -type=QBluetoothTransferRequest__Attribute
//QBluetoothTransferRequest::Attribute
type QBluetoothTransferRequest__Attribute int64

const (
	QBluetoothTransferRequest__DescriptionAttribute QBluetoothTransferRequest__Attribute = QBluetoothTransferRequest__Attribute(0)
	QBluetoothTransferRequest__TimeAttribute        QBluetoothTransferRequest__Attribute = QBluetoothTransferRequest__Attribute(1)
	QBluetoothTransferRequest__TypeAttribute        QBluetoothTransferRequest__Attribute = QBluetoothTransferRequest__Attribute(2)
	QBluetoothTransferRequest__LengthAttribute      QBluetoothTransferRequest__Attribute = QBluetoothTransferRequest__Attribute(3)
	QBluetoothTransferRequest__NameAttribute        QBluetoothTransferRequest__Attribute = QBluetoothTransferRequest__Attribute(4)
)

func NewQBluetoothTransferRequest(address QBluetoothAddress_ITF) *QBluetoothTransferRequest {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQBluetoothTransferRequest", "", address}).(*QBluetoothTransferRequest)
}

func NewQBluetoothTransferRequest2(other QBluetoothTransferRequest_ITF) *QBluetoothTransferRequest {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQBluetoothTransferRequest2", "", other}).(*QBluetoothTransferRequest)
}

func (ptr *QBluetoothTransferRequest) Address() *QBluetoothAddress {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Address"}).(*QBluetoothAddress)
}

func (ptr *QBluetoothTransferRequest) Attribute(code QBluetoothTransferRequest__Attribute, defaultValue core.QVariant_ITF) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Attribute", code, defaultValue}).(*core.QVariant)
}

func (ptr *QBluetoothTransferRequest) SetAttribute(code QBluetoothTransferRequest__Attribute, value core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAttribute", code, value})
}

func (ptr *QBluetoothTransferRequest) DestroyQBluetoothTransferRequest() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQBluetoothTransferRequest"})
}

type QBluetoothUuid struct {
	core.QUuid
}

type QBluetoothUuid_ITF interface {
	core.QUuid_ITF
	QBluetoothUuid_PTR() *QBluetoothUuid
}

func (ptr *QBluetoothUuid) QBluetoothUuid_PTR() *QBluetoothUuid {
	return ptr
}

func (ptr *QBluetoothUuid) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QUuid_PTR().Pointer()
	}
	return nil
}

func (ptr *QBluetoothUuid) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QUuid_PTR().SetPointer(p)
	}
}

func PointerFromQBluetoothUuid(ptr QBluetoothUuid_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothUuid_PTR().Pointer()
	}
	return nil
}

func (n *QBluetoothUuid) InitFromInternal(ptr uintptr, name string) {
	n.QUuid_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QBluetoothUuid) ClassNameInternalF() string {
	return n.QUuid_PTR().ClassNameInternalF()
}

func NewQBluetoothUuidFromPointer(ptr unsafe.Pointer) (n *QBluetoothUuid) {
	n = new(QBluetoothUuid)
	n.InitFromInternal(uintptr(ptr), "bluetooth.QBluetoothUuid")
	return
}

//go:generate stringer -type=QBluetoothUuid__ProtocolUuid
//QBluetoothUuid::ProtocolUuid
type QBluetoothUuid__ProtocolUuid int64

const (
	QBluetoothUuid__Sdp                    QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x0001)
	QBluetoothUuid__Udp                    QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x0002)
	QBluetoothUuid__Rfcomm                 QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x0003)
	QBluetoothUuid__Tcp                    QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x0004)
	QBluetoothUuid__TcsBin                 QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x0005)
	QBluetoothUuid__TcsAt                  QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x0006)
	QBluetoothUuid__Att                    QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x0007)
	QBluetoothUuid__Obex                   QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x0008)
	QBluetoothUuid__Ip                     QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x0009)
	QBluetoothUuid__Ftp                    QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x000A)
	QBluetoothUuid__Http                   QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x000C)
	QBluetoothUuid__Wsp                    QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x000E)
	QBluetoothUuid__Bnep                   QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x000F)
	QBluetoothUuid__Upnp                   QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x0010)
	QBluetoothUuid__Hidp                   QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x0011)
	QBluetoothUuid__HardcopyControlChannel QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x0012)
	QBluetoothUuid__HardcopyDataChannel    QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x0014)
	QBluetoothUuid__HardcopyNotification   QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x0016)
	QBluetoothUuid__Avctp                  QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x0017)
	QBluetoothUuid__Avdtp                  QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x0019)
	QBluetoothUuid__Cmtp                   QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x001B)
	QBluetoothUuid__UdiCPlain              QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x001D)
	QBluetoothUuid__McapControlChannel     QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x001E)
	QBluetoothUuid__McapDataChannel        QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x001F)
	QBluetoothUuid__L2cap                  QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x0100)
)

//go:generate stringer -type=QBluetoothUuid__ServiceClassUuid
//QBluetoothUuid::ServiceClassUuid
type QBluetoothUuid__ServiceClassUuid int64

const (
	QBluetoothUuid__ServiceDiscoveryServer                QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1000)
	QBluetoothUuid__BrowseGroupDescriptor                 QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1001)
	QBluetoothUuid__PublicBrowseGroup                     QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1002)
	QBluetoothUuid__SerialPort                            QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1101)
	QBluetoothUuid__LANAccessUsingPPP                     QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1102)
	QBluetoothUuid__DialupNetworking                      QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1103)
	QBluetoothUuid__IrMCSync                              QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1104)
	QBluetoothUuid__ObexObjectPush                        QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1105)
	QBluetoothUuid__OBEXFileTransfer                      QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1106)
	QBluetoothUuid__IrMCSyncCommand                       QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1107)
	QBluetoothUuid__Headset                               QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1108)
	QBluetoothUuid__AudioSource                           QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x110a)
	QBluetoothUuid__AudioSink                             QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x110b)
	QBluetoothUuid__AV_RemoteControlTarget                QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x110c)
	QBluetoothUuid__AdvancedAudioDistribution             QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x110d)
	QBluetoothUuid__AV_RemoteControl                      QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x110e)
	QBluetoothUuid__AV_RemoteControlController            QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x110f)
	QBluetoothUuid__HeadsetAG                             QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1112)
	QBluetoothUuid__PANU                                  QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1115)
	QBluetoothUuid__NAP                                   QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1116)
	QBluetoothUuid__GN                                    QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1117)
	QBluetoothUuid__DirectPrinting                        QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1118)
	QBluetoothUuid__ReferencePrinting                     QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1119)
	QBluetoothUuid__BasicImage                            QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x111a)
	QBluetoothUuid__ImagingResponder                      QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x111b)
	QBluetoothUuid__ImagingAutomaticArchive               QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x111c)
	QBluetoothUuid__ImagingReferenceObjects               QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x111d)
	QBluetoothUuid__Handsfree                             QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x111e)
	QBluetoothUuid__HandsfreeAudioGateway                 QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x111f)
	QBluetoothUuid__DirectPrintingReferenceObjectsService QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1120)
	QBluetoothUuid__ReflectedUI                           QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1121)
	QBluetoothUuid__BasicPrinting                         QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1122)
	QBluetoothUuid__PrintingStatus                        QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1123)
	QBluetoothUuid__HumanInterfaceDeviceService           QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1124)
	QBluetoothUuid__HardcopyCableReplacement              QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1125)
	QBluetoothUuid__HCRPrint                              QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1126)
	QBluetoothUuid__HCRScan                               QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1127)
	QBluetoothUuid__SIMAccess                             QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x112d)
	QBluetoothUuid__PhonebookAccessPCE                    QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x112e)
	QBluetoothUuid__PhonebookAccessPSE                    QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x112f)
	QBluetoothUuid__PhonebookAccess                       QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1130)
	QBluetoothUuid__HeadsetHS                             QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1131)
	QBluetoothUuid__MessageAccessServer                   QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1132)
	QBluetoothUuid__MessageNotificationServer             QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1133)
	QBluetoothUuid__MessageAccessProfile                  QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1134)
	QBluetoothUuid__GNSS                                  QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1135)
	QBluetoothUuid__GNSSServer                            QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1136)
	QBluetoothUuid__Display3D                             QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1137)
	QBluetoothUuid__Glasses3D                             QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1138)
	QBluetoothUuid__Synchronization3D                     QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1139)
	QBluetoothUuid__MPSProfile                            QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x113a)
	QBluetoothUuid__MPSService                            QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x113b)
	QBluetoothUuid__PnPInformation                        QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1200)
	QBluetoothUuid__GenericNetworking                     QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1201)
	QBluetoothUuid__GenericFileTransfer                   QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1202)
	QBluetoothUuid__GenericAudio                          QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1203)
	QBluetoothUuid__GenericTelephony                      QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1204)
	QBluetoothUuid__VideoSource                           QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1303)
	QBluetoothUuid__VideoSink                             QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1304)
	QBluetoothUuid__VideoDistribution                     QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1305)
	QBluetoothUuid__HDP                                   QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1400)
	QBluetoothUuid__HDPSource                             QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1401)
	QBluetoothUuid__HDPSink                               QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1402)
	QBluetoothUuid__GenericAccess                         QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1800)
	QBluetoothUuid__GenericAttribute                      QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1801)
	QBluetoothUuid__ImmediateAlert                        QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1802)
	QBluetoothUuid__LinkLoss                              QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1803)
	QBluetoothUuid__TxPower                               QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1804)
	QBluetoothUuid__CurrentTimeService                    QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1805)
	QBluetoothUuid__ReferenceTimeUpdateService            QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1806)
	QBluetoothUuid__NextDSTChangeService                  QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1807)
	QBluetoothUuid__Glucose                               QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1808)
	QBluetoothUuid__HealthThermometer                     QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1809)
	QBluetoothUuid__DeviceInformation                     QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x180a)
	QBluetoothUuid__HeartRate                             QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x180d)
	QBluetoothUuid__PhoneAlertStatusService               QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x180e)
	QBluetoothUuid__BatteryService                        QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x180f)
	QBluetoothUuid__BloodPressure                         QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1810)
	QBluetoothUuid__AlertNotificationService              QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1811)
	QBluetoothUuid__HumanInterfaceDevice                  QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1812)
	QBluetoothUuid__ScanParameters                        QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1813)
	QBluetoothUuid__RunningSpeedAndCadence                QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1814)
	QBluetoothUuid__CyclingSpeedAndCadence                QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1816)
	QBluetoothUuid__CyclingPower                          QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1818)
	QBluetoothUuid__LocationAndNavigation                 QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1819)
	QBluetoothUuid__EnvironmentalSensing                  QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x181a)
	QBluetoothUuid__BodyComposition                       QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x181b)
	QBluetoothUuid__UserData                              QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x181c)
	QBluetoothUuid__WeightScale                           QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x181d)
	QBluetoothUuid__BondManagement                        QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x181e)
	QBluetoothUuid__ContinuousGlucoseMonitoring           QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x181f)
)

//go:generate stringer -type=QBluetoothUuid__CharacteristicType
//QBluetoothUuid::CharacteristicType
type QBluetoothUuid__CharacteristicType int64

const (
	QBluetoothUuid__DeviceName                                    QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a00)
	QBluetoothUuid__Appearance                                    QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a01)
	QBluetoothUuid__PeripheralPrivacyFlag                         QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a02)
	QBluetoothUuid__ReconnectionAddress                           QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a03)
	QBluetoothUuid__PeripheralPreferredConnectionParameters       QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a04)
	QBluetoothUuid__ServiceChanged                                QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a05)
	QBluetoothUuid__AlertLevel                                    QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a06)
	QBluetoothUuid__TxPowerLevel                                  QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a07)
	QBluetoothUuid__DateTime                                      QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a08)
	QBluetoothUuid__DayOfWeek                                     QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a09)
	QBluetoothUuid__DayDateTime                                   QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a0a)
	QBluetoothUuid__ExactTime256                                  QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a0c)
	QBluetoothUuid__DSTOffset                                     QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a0d)
	QBluetoothUuid__TimeZone                                      QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a0e)
	QBluetoothUuid__LocalTimeInformation                          QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a0f)
	QBluetoothUuid__TimeWithDST                                   QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a11)
	QBluetoothUuid__TimeAccuracy                                  QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a12)
	QBluetoothUuid__TimeSource                                    QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a13)
	QBluetoothUuid__ReferenceTimeInformation                      QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a14)
	QBluetoothUuid__TimeUpdateControlPoint                        QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a16)
	QBluetoothUuid__TimeUpdateState                               QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a17)
	QBluetoothUuid__GlucoseMeasurement                            QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a18)
	QBluetoothUuid__BatteryLevel                                  QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a19)
	QBluetoothUuid__TemperatureMeasurement                        QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a1c)
	QBluetoothUuid__TemperatureType                               QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a1d)
	QBluetoothUuid__IntermediateTemperature                       QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a1e)
	QBluetoothUuid__MeasurementInterval                           QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a21)
	QBluetoothUuid__BootKeyboardInputReport                       QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a22)
	QBluetoothUuid__SystemID                                      QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a23)
	QBluetoothUuid__ModelNumberString                             QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a24)
	QBluetoothUuid__SerialNumberString                            QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a25)
	QBluetoothUuid__FirmwareRevisionString                        QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a26)
	QBluetoothUuid__HardwareRevisionString                        QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a27)
	QBluetoothUuid__SoftwareRevisionString                        QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a28)
	QBluetoothUuid__ManufacturerNameString                        QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a29)
	QBluetoothUuid__IEEE1107320601RegulatoryCertificationDataList QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a2a)
	QBluetoothUuid__CurrentTime                                   QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a2b)
	QBluetoothUuid__MagneticDeclination                           QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a2c)
	QBluetoothUuid__ScanRefresh                                   QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a31)
	QBluetoothUuid__BootKeyboardOutputReport                      QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a32)
	QBluetoothUuid__BootMouseInputReport                          QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a33)
	QBluetoothUuid__GlucoseMeasurementContext                     QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a34)
	QBluetoothUuid__BloodPressureMeasurement                      QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a35)
	QBluetoothUuid__IntermediateCuffPressure                      QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a36)
	QBluetoothUuid__HeartRateMeasurement                          QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a37)
	QBluetoothUuid__BodySensorLocation                            QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a38)
	QBluetoothUuid__HeartRateControlPoint                         QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a39)
	QBluetoothUuid__AlertStatus                                   QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a3f)
	QBluetoothUuid__RingerControlPoint                            QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a40)
	QBluetoothUuid__RingerSetting                                 QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a41)
	QBluetoothUuid__AlertCategoryIDBitMask                        QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a42)
	QBluetoothUuid__AlertCategoryID                               QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a43)
	QBluetoothUuid__AlertNotificationControlPoint                 QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a44)
	QBluetoothUuid__UnreadAlertStatus                             QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a45)
	QBluetoothUuid__NewAlert                                      QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a46)
	QBluetoothUuid__SupportedNewAlertCategory                     QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a47)
	QBluetoothUuid__SupportedUnreadAlertCategory                  QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a48)
	QBluetoothUuid__BloodPressureFeature                          QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a49)
	QBluetoothUuid__HIDInformation                                QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a4a)
	QBluetoothUuid__ReportMap                                     QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a4b)
	QBluetoothUuid__HIDControlPoint                               QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a4c)
	QBluetoothUuid__Report                                        QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a4d)
	QBluetoothUuid__ProtocolMode                                  QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a4e)
	QBluetoothUuid__ScanIntervalWindow                            QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a4f)
	QBluetoothUuid__PnPID                                         QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a50)
	QBluetoothUuid__GlucoseFeature                                QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a51)
	QBluetoothUuid__RecordAccessControlPoint                      QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a52)
	QBluetoothUuid__RSCMeasurement                                QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a53)
	QBluetoothUuid__RSCFeature                                    QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a54)
	QBluetoothUuid__SCControlPoint                                QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a55)
	QBluetoothUuid__CSCMeasurement                                QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a5b)
	QBluetoothUuid__CSCFeature                                    QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a5c)
	QBluetoothUuid__SensorLocation                                QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a5d)
	QBluetoothUuid__CyclingPowerMeasurement                       QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a63)
	QBluetoothUuid__CyclingPowerVector                            QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a64)
	QBluetoothUuid__CyclingPowerFeature                           QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a65)
	QBluetoothUuid__CyclingPowerControlPoint                      QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a66)
	QBluetoothUuid__LocationAndSpeed                              QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a67)
	QBluetoothUuid__Navigation                                    QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a68)
	QBluetoothUuid__PositionQuality                               QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a69)
	QBluetoothUuid__LNFeature                                     QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a6a)
	QBluetoothUuid__LNControlPoint                                QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a6b)
	QBluetoothUuid__Elevation                                     QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a6c)
	QBluetoothUuid__Pressure                                      QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a6d)
	QBluetoothUuid__Temperature                                   QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a6e)
	QBluetoothUuid__Humidity                                      QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a6f)
	QBluetoothUuid__TrueWindSpeed                                 QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a70)
	QBluetoothUuid__TrueWindDirection                             QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a71)
	QBluetoothUuid__ApparentWindSpeed                             QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a72)
	QBluetoothUuid__ApparentWindDirection                         QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a73)
	QBluetoothUuid__GustFactor                                    QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a74)
	QBluetoothUuid__PollenConcentration                           QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a75)
	QBluetoothUuid__UVIndex                                       QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a76)
	QBluetoothUuid__Irradiance                                    QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a77)
	QBluetoothUuid__Rainfall                                      QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a78)
	QBluetoothUuid__WindChill                                     QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a79)
	QBluetoothUuid__HeatIndex                                     QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a7a)
	QBluetoothUuid__DewPoint                                      QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a7b)
	QBluetoothUuid__DescriptorValueChanged                        QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a7d)
	QBluetoothUuid__AerobicHeartRateLowerLimit                    QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a7e)
	QBluetoothUuid__AerobicThreshold                              QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a7f)
	QBluetoothUuid__Age                                           QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a80)
	QBluetoothUuid__AnaerobicHeartRateLowerLimit                  QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a81)
	QBluetoothUuid__AnaerobicHeartRateUpperLimit                  QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a82)
	QBluetoothUuid__AnaerobicThreshold                            QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a83)
	QBluetoothUuid__AerobicHeartRateUpperLimit                    QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a84)
	QBluetoothUuid__DateOfBirth                                   QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a85)
	QBluetoothUuid__DateOfThresholdAssessment                     QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a86)
	QBluetoothUuid__EmailAddress                                  QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a87)
	QBluetoothUuid__FatBurnHeartRateLowerLimit                    QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a88)
	QBluetoothUuid__FatBurnHeartRateUpperLimit                    QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a89)
	QBluetoothUuid__FirstName                                     QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a8a)
	QBluetoothUuid__FiveZoneHeartRateLimits                       QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a8b)
	QBluetoothUuid__Gender                                        QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a8c)
	QBluetoothUuid__HeartRateMax                                  QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a8d)
	QBluetoothUuid__Height                                        QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a8e)
	QBluetoothUuid__HipCircumference                              QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a8f)
	QBluetoothUuid__LastName                                      QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a90)
	QBluetoothUuid__MaximumRecommendedHeartRate                   QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a91)
	QBluetoothUuid__RestingHeartRate                              QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a92)
	QBluetoothUuid__SportTypeForAerobicAnaerobicThresholds        QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a93)
	QBluetoothUuid__ThreeZoneHeartRateLimits                      QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a94)
	QBluetoothUuid__TwoZoneHeartRateLimits                        QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a95)
	QBluetoothUuid__VO2Max                                        QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a96)
	QBluetoothUuid__WaistCircumference                            QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a97)
	QBluetoothUuid__Weight                                        QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a98)
	QBluetoothUuid__DatabaseChangeIncrement                       QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a99)
	QBluetoothUuid__UserIndex                                     QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a9a)
	QBluetoothUuid__BodyCompositionFeature                        QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a9b)
	QBluetoothUuid__BodyCompositionMeasurement                    QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a9c)
	QBluetoothUuid__WeightMeasurement                             QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a9d)
	QBluetoothUuid__WeightScaleFeature                            QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a9e)
	QBluetoothUuid__UserControlPoint                              QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a9f)
	QBluetoothUuid__MagneticFluxDensity2D                         QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2aa0)
	QBluetoothUuid__MagneticFluxDensity3D                         QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2aa1)
	QBluetoothUuid__Language                                      QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2aa2)
	QBluetoothUuid__BarometricPressureTrend                       QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2aa3)
)

//go:generate stringer -type=QBluetoothUuid__DescriptorType
//QBluetoothUuid::DescriptorType
type QBluetoothUuid__DescriptorType int64

const (
	QBluetoothUuid__UnknownDescriptorType              QBluetoothUuid__DescriptorType = QBluetoothUuid__DescriptorType(0x0)
	QBluetoothUuid__CharacteristicExtendedProperties   QBluetoothUuid__DescriptorType = QBluetoothUuid__DescriptorType(0x2900)
	QBluetoothUuid__CharacteristicUserDescription      QBluetoothUuid__DescriptorType = QBluetoothUuid__DescriptorType(0x2901)
	QBluetoothUuid__ClientCharacteristicConfiguration  QBluetoothUuid__DescriptorType = QBluetoothUuid__DescriptorType(0x2902)
	QBluetoothUuid__ServerCharacteristicConfiguration  QBluetoothUuid__DescriptorType = QBluetoothUuid__DescriptorType(0x2903)
	QBluetoothUuid__CharacteristicPresentationFormat   QBluetoothUuid__DescriptorType = QBluetoothUuid__DescriptorType(0x2904)
	QBluetoothUuid__CharacteristicAggregateFormat      QBluetoothUuid__DescriptorType = QBluetoothUuid__DescriptorType(0x2905)
	QBluetoothUuid__ValidRange                         QBluetoothUuid__DescriptorType = QBluetoothUuid__DescriptorType(0x2906)
	QBluetoothUuid__ExternalReportReference            QBluetoothUuid__DescriptorType = QBluetoothUuid__DescriptorType(0x2907)
	QBluetoothUuid__ReportReference                    QBluetoothUuid__DescriptorType = QBluetoothUuid__DescriptorType(0x2908)
	QBluetoothUuid__EnvironmentalSensingConfiguration  QBluetoothUuid__DescriptorType = QBluetoothUuid__DescriptorType(0x290b)
	QBluetoothUuid__EnvironmentalSensingMeasurement    QBluetoothUuid__DescriptorType = QBluetoothUuid__DescriptorType(0x290c)
	QBluetoothUuid__EnvironmentalSensingTriggerSetting QBluetoothUuid__DescriptorType = QBluetoothUuid__DescriptorType(0x290d)
)

func NewQBluetoothUuid() *QBluetoothUuid {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQBluetoothUuid", ""}).(*QBluetoothUuid)
}

func NewQBluetoothUuid2(uuid QBluetoothUuid__ProtocolUuid) *QBluetoothUuid {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQBluetoothUuid2", "", uuid}).(*QBluetoothUuid)
}

func NewQBluetoothUuid3(uuid QBluetoothUuid__ServiceClassUuid) *QBluetoothUuid {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQBluetoothUuid3", "", uuid}).(*QBluetoothUuid)
}

func NewQBluetoothUuid4(uuid QBluetoothUuid__CharacteristicType) *QBluetoothUuid {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQBluetoothUuid4", "", uuid}).(*QBluetoothUuid)
}

func NewQBluetoothUuid5(uuid QBluetoothUuid__DescriptorType) *QBluetoothUuid {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQBluetoothUuid5", "", uuid}).(*QBluetoothUuid)
}

func NewQBluetoothUuid6(uuid uint16) *QBluetoothUuid {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQBluetoothUuid6", "", uuid}).(*QBluetoothUuid)
}

func NewQBluetoothUuid7(uuid uint) *QBluetoothUuid {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQBluetoothUuid7", "", uuid}).(*QBluetoothUuid)
}

func NewQBluetoothUuid9(uuid string) *QBluetoothUuid {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQBluetoothUuid9", "", uuid}).(*QBluetoothUuid)
}

func NewQBluetoothUuid10(uuid QBluetoothUuid_ITF) *QBluetoothUuid {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQBluetoothUuid10", "", uuid}).(*QBluetoothUuid)
}

func NewQBluetoothUuid11(uuid core.QUuid_ITF) *QBluetoothUuid {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQBluetoothUuid11", "", uuid}).(*QBluetoothUuid)
}

func QBluetoothUuid_CharacteristicToString(uuid QBluetoothUuid__CharacteristicType) string {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.QBluetoothUuid_CharacteristicToString", "", uuid}).(string)
}

func (ptr *QBluetoothUuid) CharacteristicToString(uuid QBluetoothUuid__CharacteristicType) string {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.QBluetoothUuid_CharacteristicToString", "", uuid}).(string)
}

func QBluetoothUuid_DescriptorToString(uuid QBluetoothUuid__DescriptorType) string {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.QBluetoothUuid_DescriptorToString", "", uuid}).(string)
}

func (ptr *QBluetoothUuid) DescriptorToString(uuid QBluetoothUuid__DescriptorType) string {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.QBluetoothUuid_DescriptorToString", "", uuid}).(string)
}

func (ptr *QBluetoothUuid) MinimumSize() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinimumSize"}).(float64))
}

func QBluetoothUuid_ProtocolToString(uuid QBluetoothUuid__ProtocolUuid) string {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.QBluetoothUuid_ProtocolToString", "", uuid}).(string)
}

func (ptr *QBluetoothUuid) ProtocolToString(uuid QBluetoothUuid__ProtocolUuid) string {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.QBluetoothUuid_ProtocolToString", "", uuid}).(string)
}

func QBluetoothUuid_ServiceClassToString(uuid QBluetoothUuid__ServiceClassUuid) string {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.QBluetoothUuid_ServiceClassToString", "", uuid}).(string)
}

func (ptr *QBluetoothUuid) ServiceClassToString(uuid QBluetoothUuid__ServiceClassUuid) string {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.QBluetoothUuid_ServiceClassToString", "", uuid}).(string)
}

func (ptr *QBluetoothUuid) ToUInt16(ok *bool) uint16 {

	return uint16(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToUInt16", ok}).(float64))
}

func (ptr *QBluetoothUuid) ToUInt32(ok *bool) uint {

	return uint(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToUInt32", ok}).(float64))
}

func (ptr *QBluetoothUuid) DestroyQBluetoothUuid() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQBluetoothUuid"})
}

type QLowEnergyAdvertisingData struct {
	internal.Internal
}

type QLowEnergyAdvertisingData_ITF interface {
	QLowEnergyAdvertisingData_PTR() *QLowEnergyAdvertisingData
}

func (ptr *QLowEnergyAdvertisingData) QLowEnergyAdvertisingData_PTR() *QLowEnergyAdvertisingData {
	return ptr
}

func (ptr *QLowEnergyAdvertisingData) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QLowEnergyAdvertisingData) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQLowEnergyAdvertisingData(ptr QLowEnergyAdvertisingData_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLowEnergyAdvertisingData_PTR().Pointer()
	}
	return nil
}

func (n *QLowEnergyAdvertisingData) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQLowEnergyAdvertisingDataFromPointer(ptr unsafe.Pointer) (n *QLowEnergyAdvertisingData) {
	n = new(QLowEnergyAdvertisingData)
	n.InitFromInternal(uintptr(ptr), "bluetooth.QLowEnergyAdvertisingData")
	return
}

//go:generate stringer -type=QLowEnergyAdvertisingData__Discoverability
//QLowEnergyAdvertisingData::Discoverability
type QLowEnergyAdvertisingData__Discoverability int64

const (
	QLowEnergyAdvertisingData__DiscoverabilityNone    QLowEnergyAdvertisingData__Discoverability = QLowEnergyAdvertisingData__Discoverability(0)
	QLowEnergyAdvertisingData__DiscoverabilityLimited QLowEnergyAdvertisingData__Discoverability = QLowEnergyAdvertisingData__Discoverability(1)
	QLowEnergyAdvertisingData__DiscoverabilityGeneral QLowEnergyAdvertisingData__Discoverability = QLowEnergyAdvertisingData__Discoverability(2)
)

func NewQLowEnergyAdvertisingData() *QLowEnergyAdvertisingData {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQLowEnergyAdvertisingData", ""}).(*QLowEnergyAdvertisingData)
}

func NewQLowEnergyAdvertisingData2(other QLowEnergyAdvertisingData_ITF) *QLowEnergyAdvertisingData {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQLowEnergyAdvertisingData2", "", other}).(*QLowEnergyAdvertisingData)
}

func (ptr *QLowEnergyAdvertisingData) Discoverability() QLowEnergyAdvertisingData__Discoverability {

	return QLowEnergyAdvertisingData__Discoverability(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Discoverability"}).(float64))
}

func (ptr *QLowEnergyAdvertisingData) IncludePowerLevel() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IncludePowerLevel"}).(bool)
}

func QLowEnergyAdvertisingData_InvalidManufacturerId() uint16 {

	return uint16(internal.CallLocalFunction([]interface{}{"", "", "bluetooth.QLowEnergyAdvertisingData_InvalidManufacturerId", ""}).(float64))
}

func (ptr *QLowEnergyAdvertisingData) InvalidManufacturerId() uint16 {

	return uint16(internal.CallLocalFunction([]interface{}{"", "", "bluetooth.QLowEnergyAdvertisingData_InvalidManufacturerId", ""}).(float64))
}

func (ptr *QLowEnergyAdvertisingData) LocalName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LocalName"}).(string)
}

func (ptr *QLowEnergyAdvertisingData) ManufacturerId() uint16 {

	return uint16(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ManufacturerId"}).(float64))
}

func (ptr *QLowEnergyAdvertisingData) RawData() *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RawData"}).(*core.QByteArray)
}

func (ptr *QLowEnergyAdvertisingData) Services() []*QBluetoothUuid {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Services"}).([]*QBluetoothUuid)
}

func (ptr *QLowEnergyAdvertisingData) SetDiscoverability(mode QLowEnergyAdvertisingData__Discoverability) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDiscoverability", mode})
}

func (ptr *QLowEnergyAdvertisingData) SetIncludePowerLevel(doInclude bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetIncludePowerLevel", doInclude})
}

func (ptr *QLowEnergyAdvertisingData) SetLocalName(name string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLocalName", name})
}

func (ptr *QLowEnergyAdvertisingData) SetManufacturerData(id uint16, data core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetManufacturerData", id, data})
}

func (ptr *QLowEnergyAdvertisingData) SetRawData(data core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRawData", data})
}

func (ptr *QLowEnergyAdvertisingData) SetServices(services []*QBluetoothUuid) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetServices", services})
}

func (ptr *QLowEnergyAdvertisingData) Swap(other QLowEnergyAdvertisingData_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", other})
}

func (ptr *QLowEnergyAdvertisingData) DestroyQLowEnergyAdvertisingData() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQLowEnergyAdvertisingData"})
}

func (ptr *QLowEnergyAdvertisingData) __services_atList(i int) *QBluetoothUuid {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__services_atList", i}).(*QBluetoothUuid)
}

func (ptr *QLowEnergyAdvertisingData) __services_setList(i QBluetoothUuid_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__services_setList", i})
}

func (ptr *QLowEnergyAdvertisingData) __services_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__services_newList"}).(unsafe.Pointer)
}

func (ptr *QLowEnergyAdvertisingData) __setServices_services_atList(i int) *QBluetoothUuid {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setServices_services_atList", i}).(*QBluetoothUuid)
}

func (ptr *QLowEnergyAdvertisingData) __setServices_services_setList(i QBluetoothUuid_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setServices_services_setList", i})
}

func (ptr *QLowEnergyAdvertisingData) __setServices_services_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setServices_services_newList"}).(unsafe.Pointer)
}

type QLowEnergyAdvertisingParameters struct {
	internal.Internal
}

type QLowEnergyAdvertisingParameters_ITF interface {
	QLowEnergyAdvertisingParameters_PTR() *QLowEnergyAdvertisingParameters
}

func (ptr *QLowEnergyAdvertisingParameters) QLowEnergyAdvertisingParameters_PTR() *QLowEnergyAdvertisingParameters {
	return ptr
}

func (ptr *QLowEnergyAdvertisingParameters) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QLowEnergyAdvertisingParameters) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQLowEnergyAdvertisingParameters(ptr QLowEnergyAdvertisingParameters_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLowEnergyAdvertisingParameters_PTR().Pointer()
	}
	return nil
}

func (n *QLowEnergyAdvertisingParameters) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQLowEnergyAdvertisingParametersFromPointer(ptr unsafe.Pointer) (n *QLowEnergyAdvertisingParameters) {
	n = new(QLowEnergyAdvertisingParameters)
	n.InitFromInternal(uintptr(ptr), "bluetooth.QLowEnergyAdvertisingParameters")
	return
}

//go:generate stringer -type=QLowEnergyAdvertisingParameters__Mode
//QLowEnergyAdvertisingParameters::Mode
type QLowEnergyAdvertisingParameters__Mode int64

const (
	QLowEnergyAdvertisingParameters__AdvInd        QLowEnergyAdvertisingParameters__Mode = QLowEnergyAdvertisingParameters__Mode(0x0)
	QLowEnergyAdvertisingParameters__AdvScanInd    QLowEnergyAdvertisingParameters__Mode = QLowEnergyAdvertisingParameters__Mode(0x2)
	QLowEnergyAdvertisingParameters__AdvNonConnInd QLowEnergyAdvertisingParameters__Mode = QLowEnergyAdvertisingParameters__Mode(0x3)
)

//go:generate stringer -type=QLowEnergyAdvertisingParameters__FilterPolicy
//QLowEnergyAdvertisingParameters::FilterPolicy
type QLowEnergyAdvertisingParameters__FilterPolicy int64

const (
	QLowEnergyAdvertisingParameters__IgnoreWhiteList                      QLowEnergyAdvertisingParameters__FilterPolicy = QLowEnergyAdvertisingParameters__FilterPolicy(0x00)
	QLowEnergyAdvertisingParameters__UseWhiteListForScanning              QLowEnergyAdvertisingParameters__FilterPolicy = QLowEnergyAdvertisingParameters__FilterPolicy(0x01)
	QLowEnergyAdvertisingParameters__UseWhiteListForConnecting            QLowEnergyAdvertisingParameters__FilterPolicy = QLowEnergyAdvertisingParameters__FilterPolicy(0x02)
	QLowEnergyAdvertisingParameters__UseWhiteListForScanningAndConnecting QLowEnergyAdvertisingParameters__FilterPolicy = QLowEnergyAdvertisingParameters__FilterPolicy(0x03)
)

func NewQLowEnergyAdvertisingParameters() *QLowEnergyAdvertisingParameters {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQLowEnergyAdvertisingParameters", ""}).(*QLowEnergyAdvertisingParameters)
}

func NewQLowEnergyAdvertisingParameters2(other QLowEnergyAdvertisingParameters_ITF) *QLowEnergyAdvertisingParameters {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQLowEnergyAdvertisingParameters2", "", other}).(*QLowEnergyAdvertisingParameters)
}

func (ptr *QLowEnergyAdvertisingParameters) FilterPolicy() QLowEnergyAdvertisingParameters__FilterPolicy {

	return QLowEnergyAdvertisingParameters__FilterPolicy(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FilterPolicy"}).(float64))
}

func (ptr *QLowEnergyAdvertisingParameters) MaximumInterval() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MaximumInterval"}).(float64))
}

func (ptr *QLowEnergyAdvertisingParameters) MinimumInterval() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinimumInterval"}).(float64))
}

func (ptr *QLowEnergyAdvertisingParameters) Mode() QLowEnergyAdvertisingParameters__Mode {

	return QLowEnergyAdvertisingParameters__Mode(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Mode"}).(float64))
}

func (ptr *QLowEnergyAdvertisingParameters) SetInterval(minimum uint16, maximum uint16) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetInterval", minimum, maximum})
}

func (ptr *QLowEnergyAdvertisingParameters) SetMode(mode QLowEnergyAdvertisingParameters__Mode) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMode", mode})
}

func (ptr *QLowEnergyAdvertisingParameters) Swap(other QLowEnergyAdvertisingParameters_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", other})
}

func (ptr *QLowEnergyAdvertisingParameters) DestroyQLowEnergyAdvertisingParameters() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQLowEnergyAdvertisingParameters"})
}

type QLowEnergyCharacteristic struct {
	internal.Internal
}

type QLowEnergyCharacteristic_ITF interface {
	QLowEnergyCharacteristic_PTR() *QLowEnergyCharacteristic
}

func (ptr *QLowEnergyCharacteristic) QLowEnergyCharacteristic_PTR() *QLowEnergyCharacteristic {
	return ptr
}

func (ptr *QLowEnergyCharacteristic) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QLowEnergyCharacteristic) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQLowEnergyCharacteristic(ptr QLowEnergyCharacteristic_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLowEnergyCharacteristic_PTR().Pointer()
	}
	return nil
}

func (n *QLowEnergyCharacteristic) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQLowEnergyCharacteristicFromPointer(ptr unsafe.Pointer) (n *QLowEnergyCharacteristic) {
	n = new(QLowEnergyCharacteristic)
	n.InitFromInternal(uintptr(ptr), "bluetooth.QLowEnergyCharacteristic")
	return
}

//go:generate stringer -type=QLowEnergyCharacteristic__PropertyType
//QLowEnergyCharacteristic::PropertyType
type QLowEnergyCharacteristic__PropertyType int64

const (
	QLowEnergyCharacteristic__Unknown          QLowEnergyCharacteristic__PropertyType = QLowEnergyCharacteristic__PropertyType(0x00)
	QLowEnergyCharacteristic__Broadcasting     QLowEnergyCharacteristic__PropertyType = QLowEnergyCharacteristic__PropertyType(0x01)
	QLowEnergyCharacteristic__Read             QLowEnergyCharacteristic__PropertyType = QLowEnergyCharacteristic__PropertyType(0x02)
	QLowEnergyCharacteristic__WriteNoResponse  QLowEnergyCharacteristic__PropertyType = QLowEnergyCharacteristic__PropertyType(0x04)
	QLowEnergyCharacteristic__Write            QLowEnergyCharacteristic__PropertyType = QLowEnergyCharacteristic__PropertyType(0x08)
	QLowEnergyCharacteristic__Notify           QLowEnergyCharacteristic__PropertyType = QLowEnergyCharacteristic__PropertyType(0x10)
	QLowEnergyCharacteristic__Indicate         QLowEnergyCharacteristic__PropertyType = QLowEnergyCharacteristic__PropertyType(0x20)
	QLowEnergyCharacteristic__WriteSigned      QLowEnergyCharacteristic__PropertyType = QLowEnergyCharacteristic__PropertyType(0x40)
	QLowEnergyCharacteristic__ExtendedProperty QLowEnergyCharacteristic__PropertyType = QLowEnergyCharacteristic__PropertyType(0x80)
)

func NewQLowEnergyCharacteristic() *QLowEnergyCharacteristic {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQLowEnergyCharacteristic", ""}).(*QLowEnergyCharacteristic)
}

func NewQLowEnergyCharacteristic2(other QLowEnergyCharacteristic_ITF) *QLowEnergyCharacteristic {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQLowEnergyCharacteristic2", "", other}).(*QLowEnergyCharacteristic)
}

func (ptr *QLowEnergyCharacteristic) Descriptor(uuid QBluetoothUuid_ITF) *QLowEnergyDescriptor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Descriptor", uuid}).(*QLowEnergyDescriptor)
}

func (ptr *QLowEnergyCharacteristic) Descriptors() []*QLowEnergyDescriptor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Descriptors"}).([]*QLowEnergyDescriptor)
}

func (ptr *QLowEnergyCharacteristic) IsValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsValid"}).(bool)
}

func (ptr *QLowEnergyCharacteristic) Name() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Name"}).(string)
}

func (ptr *QLowEnergyCharacteristic) Properties() QLowEnergyCharacteristic__PropertyType {

	return QLowEnergyCharacteristic__PropertyType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Properties"}).(float64))
}

func (ptr *QLowEnergyCharacteristic) Uuid() *QBluetoothUuid {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Uuid"}).(*QBluetoothUuid)
}

func (ptr *QLowEnergyCharacteristic) Value() *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Value"}).(*core.QByteArray)
}

func (ptr *QLowEnergyCharacteristic) DestroyQLowEnergyCharacteristic() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQLowEnergyCharacteristic"})
}

func (ptr *QLowEnergyCharacteristic) __descriptors_atList(i int) *QLowEnergyDescriptor {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__descriptors_atList", i}).(*QLowEnergyDescriptor)
}

func (ptr *QLowEnergyCharacteristic) __descriptors_setList(i QLowEnergyDescriptor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__descriptors_setList", i})
}

func (ptr *QLowEnergyCharacteristic) __descriptors_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__descriptors_newList"}).(unsafe.Pointer)
}

type QLowEnergyCharacteristicData struct {
	internal.Internal
}

type QLowEnergyCharacteristicData_ITF interface {
	QLowEnergyCharacteristicData_PTR() *QLowEnergyCharacteristicData
}

func (ptr *QLowEnergyCharacteristicData) QLowEnergyCharacteristicData_PTR() *QLowEnergyCharacteristicData {
	return ptr
}

func (ptr *QLowEnergyCharacteristicData) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QLowEnergyCharacteristicData) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQLowEnergyCharacteristicData(ptr QLowEnergyCharacteristicData_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLowEnergyCharacteristicData_PTR().Pointer()
	}
	return nil
}

func (n *QLowEnergyCharacteristicData) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQLowEnergyCharacteristicDataFromPointer(ptr unsafe.Pointer) (n *QLowEnergyCharacteristicData) {
	n = new(QLowEnergyCharacteristicData)
	n.InitFromInternal(uintptr(ptr), "bluetooth.QLowEnergyCharacteristicData")
	return
}
func NewQLowEnergyCharacteristicData() *QLowEnergyCharacteristicData {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQLowEnergyCharacteristicData", ""}).(*QLowEnergyCharacteristicData)
}

func NewQLowEnergyCharacteristicData2(other QLowEnergyCharacteristicData_ITF) *QLowEnergyCharacteristicData {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQLowEnergyCharacteristicData2", "", other}).(*QLowEnergyCharacteristicData)
}

func (ptr *QLowEnergyCharacteristicData) AddDescriptor(descriptor QLowEnergyDescriptorData_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddDescriptor", descriptor})
}

func (ptr *QLowEnergyCharacteristicData) Descriptors() []*QLowEnergyDescriptorData {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Descriptors"}).([]*QLowEnergyDescriptorData)
}

func (ptr *QLowEnergyCharacteristicData) IsValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsValid"}).(bool)
}

func (ptr *QLowEnergyCharacteristicData) MaximumValueLength() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MaximumValueLength"}).(float64))
}

func (ptr *QLowEnergyCharacteristicData) MinimumValueLength() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinimumValueLength"}).(float64))
}

func (ptr *QLowEnergyCharacteristicData) Properties() QLowEnergyCharacteristic__PropertyType {

	return QLowEnergyCharacteristic__PropertyType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Properties"}).(float64))
}

func (ptr *QLowEnergyCharacteristicData) ReadConstraints() QBluetooth__AttAccessConstraint {

	return QBluetooth__AttAccessConstraint(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReadConstraints"}).(float64))
}

func (ptr *QLowEnergyCharacteristicData) SetDescriptors(descriptors []*QLowEnergyDescriptorData) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDescriptors", descriptors})
}

func (ptr *QLowEnergyCharacteristicData) SetProperties(properties QLowEnergyCharacteristic__PropertyType) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetProperties", properties})
}

func (ptr *QLowEnergyCharacteristicData) SetReadConstraints(constraints QBluetooth__AttAccessConstraint) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetReadConstraints", constraints})
}

func (ptr *QLowEnergyCharacteristicData) SetUuid(uuid QBluetoothUuid_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetUuid", uuid})
}

func (ptr *QLowEnergyCharacteristicData) SetValue(value core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetValue", value})
}

func (ptr *QLowEnergyCharacteristicData) SetValueLength(minimum int, maximum int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetValueLength", minimum, maximum})
}

func (ptr *QLowEnergyCharacteristicData) SetWriteConstraints(constraints QBluetooth__AttAccessConstraint) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWriteConstraints", constraints})
}

func (ptr *QLowEnergyCharacteristicData) Swap(other QLowEnergyCharacteristicData_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", other})
}

func (ptr *QLowEnergyCharacteristicData) Uuid() *QBluetoothUuid {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Uuid"}).(*QBluetoothUuid)
}

func (ptr *QLowEnergyCharacteristicData) Value() *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Value"}).(*core.QByteArray)
}

func (ptr *QLowEnergyCharacteristicData) WriteConstraints() QBluetooth__AttAccessConstraint {

	return QBluetooth__AttAccessConstraint(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WriteConstraints"}).(float64))
}

func (ptr *QLowEnergyCharacteristicData) DestroyQLowEnergyCharacteristicData() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQLowEnergyCharacteristicData"})
}

func (ptr *QLowEnergyCharacteristicData) __descriptors_atList(i int) *QLowEnergyDescriptorData {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__descriptors_atList", i}).(*QLowEnergyDescriptorData)
}

func (ptr *QLowEnergyCharacteristicData) __descriptors_setList(i QLowEnergyDescriptorData_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__descriptors_setList", i})
}

func (ptr *QLowEnergyCharacteristicData) __descriptors_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__descriptors_newList"}).(unsafe.Pointer)
}

func (ptr *QLowEnergyCharacteristicData) __setDescriptors_descriptors_atList(i int) *QLowEnergyDescriptorData {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setDescriptors_descriptors_atList", i}).(*QLowEnergyDescriptorData)
}

func (ptr *QLowEnergyCharacteristicData) __setDescriptors_descriptors_setList(i QLowEnergyDescriptorData_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setDescriptors_descriptors_setList", i})
}

func (ptr *QLowEnergyCharacteristicData) __setDescriptors_descriptors_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setDescriptors_descriptors_newList"}).(unsafe.Pointer)
}

type QLowEnergyConnectionParameters struct {
	internal.Internal
}

type QLowEnergyConnectionParameters_ITF interface {
	QLowEnergyConnectionParameters_PTR() *QLowEnergyConnectionParameters
}

func (ptr *QLowEnergyConnectionParameters) QLowEnergyConnectionParameters_PTR() *QLowEnergyConnectionParameters {
	return ptr
}

func (ptr *QLowEnergyConnectionParameters) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QLowEnergyConnectionParameters) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQLowEnergyConnectionParameters(ptr QLowEnergyConnectionParameters_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLowEnergyConnectionParameters_PTR().Pointer()
	}
	return nil
}

func (n *QLowEnergyConnectionParameters) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQLowEnergyConnectionParametersFromPointer(ptr unsafe.Pointer) (n *QLowEnergyConnectionParameters) {
	n = new(QLowEnergyConnectionParameters)
	n.InitFromInternal(uintptr(ptr), "bluetooth.QLowEnergyConnectionParameters")
	return
}
func NewQLowEnergyConnectionParameters() *QLowEnergyConnectionParameters {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQLowEnergyConnectionParameters", ""}).(*QLowEnergyConnectionParameters)
}

func NewQLowEnergyConnectionParameters2(other QLowEnergyConnectionParameters_ITF) *QLowEnergyConnectionParameters {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQLowEnergyConnectionParameters2", "", other}).(*QLowEnergyConnectionParameters)
}

func (ptr *QLowEnergyConnectionParameters) Latency() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Latency"}).(float64))
}

func (ptr *QLowEnergyConnectionParameters) MaximumInterval() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MaximumInterval"}).(float64)
}

func (ptr *QLowEnergyConnectionParameters) MinimumInterval() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinimumInterval"}).(float64)
}

func (ptr *QLowEnergyConnectionParameters) SetIntervalRange(minimum float64, maximum float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetIntervalRange", minimum, maximum})
}

func (ptr *QLowEnergyConnectionParameters) SetLatency(latency int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLatency", latency})
}

func (ptr *QLowEnergyConnectionParameters) SetSupervisionTimeout(timeout int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSupervisionTimeout", timeout})
}

func (ptr *QLowEnergyConnectionParameters) SupervisionTimeout() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SupervisionTimeout"}).(float64))
}

func (ptr *QLowEnergyConnectionParameters) Swap(other QLowEnergyConnectionParameters_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", other})
}

func (ptr *QLowEnergyConnectionParameters) DestroyQLowEnergyConnectionParameters() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQLowEnergyConnectionParameters"})
}

type QLowEnergyController struct {
	core.QObject
}

type QLowEnergyController_ITF interface {
	core.QObject_ITF
	QLowEnergyController_PTR() *QLowEnergyController
}

func (ptr *QLowEnergyController) QLowEnergyController_PTR() *QLowEnergyController {
	return ptr
}

func (ptr *QLowEnergyController) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QLowEnergyController) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQLowEnergyController(ptr QLowEnergyController_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLowEnergyController_PTR().Pointer()
	}
	return nil
}

func (n *QLowEnergyController) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QLowEnergyController) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQLowEnergyControllerFromPointer(ptr unsafe.Pointer) (n *QLowEnergyController) {
	n = new(QLowEnergyController)
	n.InitFromInternal(uintptr(ptr), "bluetooth.QLowEnergyController")
	return
}

//go:generate stringer -type=QLowEnergyController__Error
//QLowEnergyController::Error
type QLowEnergyController__Error int64

const (
	QLowEnergyController__NoError                      QLowEnergyController__Error = QLowEnergyController__Error(0)
	QLowEnergyController__UnknownError                 QLowEnergyController__Error = QLowEnergyController__Error(1)
	QLowEnergyController__UnknownRemoteDeviceError     QLowEnergyController__Error = QLowEnergyController__Error(2)
	QLowEnergyController__NetworkError                 QLowEnergyController__Error = QLowEnergyController__Error(3)
	QLowEnergyController__InvalidBluetoothAdapterError QLowEnergyController__Error = QLowEnergyController__Error(4)
	QLowEnergyController__ConnectionError              QLowEnergyController__Error = QLowEnergyController__Error(5)
	QLowEnergyController__AdvertisingError             QLowEnergyController__Error = QLowEnergyController__Error(6)
	QLowEnergyController__RemoteHostClosedError        QLowEnergyController__Error = QLowEnergyController__Error(7)
)

//go:generate stringer -type=QLowEnergyController__ControllerState
//QLowEnergyController::ControllerState
type QLowEnergyController__ControllerState int64

const (
	QLowEnergyController__UnconnectedState QLowEnergyController__ControllerState = QLowEnergyController__ControllerState(0)
	QLowEnergyController__ConnectingState  QLowEnergyController__ControllerState = QLowEnergyController__ControllerState(1)
	QLowEnergyController__ConnectedState   QLowEnergyController__ControllerState = QLowEnergyController__ControllerState(2)
	QLowEnergyController__DiscoveringState QLowEnergyController__ControllerState = QLowEnergyController__ControllerState(3)
	QLowEnergyController__DiscoveredState  QLowEnergyController__ControllerState = QLowEnergyController__ControllerState(4)
	QLowEnergyController__ClosingState     QLowEnergyController__ControllerState = QLowEnergyController__ControllerState(5)
	QLowEnergyController__AdvertisingState QLowEnergyController__ControllerState = QLowEnergyController__ControllerState(6)
)

//go:generate stringer -type=QLowEnergyController__RemoteAddressType
//QLowEnergyController::RemoteAddressType
type QLowEnergyController__RemoteAddressType int64

const (
	QLowEnergyController__PublicAddress QLowEnergyController__RemoteAddressType = QLowEnergyController__RemoteAddressType(0)
	QLowEnergyController__RandomAddress QLowEnergyController__RemoteAddressType = QLowEnergyController__RemoteAddressType(1)
)

//go:generate stringer -type=QLowEnergyController__Role
//QLowEnergyController::Role
type QLowEnergyController__Role int64

const (
	QLowEnergyController__CentralRole    QLowEnergyController__Role = QLowEnergyController__Role(0)
	QLowEnergyController__PeripheralRole QLowEnergyController__Role = QLowEnergyController__Role(1)
)

func (ptr *QLowEnergyController) AddService(service QLowEnergyServiceData_ITF, parent core.QObject_ITF) *QLowEnergyService {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddService", service, parent}).(*QLowEnergyService)
}

func (ptr *QLowEnergyController) ConnectToDevice() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectToDevice"})
}

func (ptr *QLowEnergyController) ConnectConnected(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectConnected", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLowEnergyController) DisconnectConnected() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectConnected"})
}

func (ptr *QLowEnergyController) Connected() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Connected"})
}

func (ptr *QLowEnergyController) ConnectConnectionUpdated(f func(newParameters *QLowEnergyConnectionParameters)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectConnectionUpdated", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLowEnergyController) DisconnectConnectionUpdated() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectConnectionUpdated"})
}

func (ptr *QLowEnergyController) ConnectionUpdated(newParameters QLowEnergyConnectionParameters_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectionUpdated", newParameters})
}

func QLowEnergyController_CreateCentral(remoteDevice QBluetoothDeviceInfo_ITF, parent core.QObject_ITF) *QLowEnergyController {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.QLowEnergyController_CreateCentral", "", remoteDevice, parent}).(*QLowEnergyController)
}

func (ptr *QLowEnergyController) CreateCentral(remoteDevice QBluetoothDeviceInfo_ITF, parent core.QObject_ITF) *QLowEnergyController {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.QLowEnergyController_CreateCentral", "", remoteDevice, parent}).(*QLowEnergyController)
}

func QLowEnergyController_CreatePeripheral(parent core.QObject_ITF) *QLowEnergyController {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.QLowEnergyController_CreatePeripheral", "", parent}).(*QLowEnergyController)
}

func (ptr *QLowEnergyController) CreatePeripheral(parent core.QObject_ITF) *QLowEnergyController {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.QLowEnergyController_CreatePeripheral", "", parent}).(*QLowEnergyController)
}

func (ptr *QLowEnergyController) CreateServiceObject(serviceUuid QBluetoothUuid_ITF, parent core.QObject_ITF) *QLowEnergyService {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateServiceObject", serviceUuid, parent}).(*QLowEnergyService)
}

func (ptr *QLowEnergyController) DisconnectFromDevice() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFromDevice"})
}

func (ptr *QLowEnergyController) ConnectDisconnected(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDisconnected", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLowEnergyController) DisconnectDisconnected() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDisconnected"})
}

func (ptr *QLowEnergyController) Disconnected() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Disconnected"})
}

func (ptr *QLowEnergyController) DiscoverServices() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DiscoverServices"})
}

func (ptr *QLowEnergyController) ConnectDiscoveryFinished(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDiscoveryFinished", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLowEnergyController) DisconnectDiscoveryFinished() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDiscoveryFinished"})
}

func (ptr *QLowEnergyController) DiscoveryFinished() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DiscoveryFinished"})
}

func (ptr *QLowEnergyController) Error() QLowEnergyController__Error {

	return QLowEnergyController__Error(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error"}).(float64))
}

func (ptr *QLowEnergyController) ConnectError2(f func(newError QLowEnergyController__Error)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectError2", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLowEnergyController) DisconnectError2() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectError2"})
}

func (ptr *QLowEnergyController) Error2(newError QLowEnergyController__Error) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error2", newError})
}

func (ptr *QLowEnergyController) ErrorString() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ErrorString"}).(string)
}

func (ptr *QLowEnergyController) LocalAddress() *QBluetoothAddress {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LocalAddress"}).(*QBluetoothAddress)
}

func (ptr *QLowEnergyController) RemoteAddress() *QBluetoothAddress {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoteAddress"}).(*QBluetoothAddress)
}

func (ptr *QLowEnergyController) RemoteAddressType() QLowEnergyController__RemoteAddressType {

	return QLowEnergyController__RemoteAddressType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoteAddressType"}).(float64))
}

func (ptr *QLowEnergyController) RemoteDeviceUuid() *QBluetoothUuid {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoteDeviceUuid"}).(*QBluetoothUuid)
}

func (ptr *QLowEnergyController) RemoteName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoteName"}).(string)
}

func (ptr *QLowEnergyController) RequestConnectionUpdate(parameters QLowEnergyConnectionParameters_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RequestConnectionUpdate", parameters})
}

func (ptr *QLowEnergyController) Role() QLowEnergyController__Role {

	return QLowEnergyController__Role(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Role"}).(float64))
}

func (ptr *QLowEnergyController) ConnectServiceDiscovered(f func(newService *QBluetoothUuid)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectServiceDiscovered", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLowEnergyController) DisconnectServiceDiscovered() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectServiceDiscovered"})
}

func (ptr *QLowEnergyController) ServiceDiscovered(newService QBluetoothUuid_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ServiceDiscovered", newService})
}

func (ptr *QLowEnergyController) Services() []*QBluetoothUuid {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Services"}).([]*QBluetoothUuid)
}

func (ptr *QLowEnergyController) SetRemoteAddressType(ty QLowEnergyController__RemoteAddressType) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRemoteAddressType", ty})
}

func (ptr *QLowEnergyController) StartAdvertising(parameters QLowEnergyAdvertisingParameters_ITF, advertisingData QLowEnergyAdvertisingData_ITF, scanResponseData QLowEnergyAdvertisingData_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StartAdvertising", parameters, advertisingData, scanResponseData})
}

func (ptr *QLowEnergyController) State() QLowEnergyController__ControllerState {

	return QLowEnergyController__ControllerState(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "State"}).(float64))
}

func (ptr *QLowEnergyController) ConnectStateChanged(f func(state QLowEnergyController__ControllerState)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStateChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLowEnergyController) DisconnectStateChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStateChanged"})
}

func (ptr *QLowEnergyController) StateChanged(state QLowEnergyController__ControllerState) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StateChanged", state})
}

func (ptr *QLowEnergyController) StopAdvertising() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StopAdvertising"})
}

func (ptr *QLowEnergyController) ConnectDestroyQLowEnergyController(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQLowEnergyController", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLowEnergyController) DisconnectDestroyQLowEnergyController() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQLowEnergyController"})
}

func (ptr *QLowEnergyController) DestroyQLowEnergyController() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQLowEnergyController"})
}

func (ptr *QLowEnergyController) DestroyQLowEnergyControllerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQLowEnergyControllerDefault"})
}

func (ptr *QLowEnergyController) __services_atList(i int) *QBluetoothUuid {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__services_atList", i}).(*QBluetoothUuid)
}

func (ptr *QLowEnergyController) __services_setList(i QBluetoothUuid_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__services_setList", i})
}

func (ptr *QLowEnergyController) __services_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__services_newList"}).(unsafe.Pointer)
}

func (ptr *QLowEnergyController) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QLowEnergyController) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QLowEnergyController) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QLowEnergyController) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QLowEnergyController) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QLowEnergyController) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QLowEnergyController) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QLowEnergyController) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QLowEnergyController) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QLowEnergyController) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QLowEnergyController) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QLowEnergyController) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QLowEnergyController) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QLowEnergyController) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QLowEnergyController) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QLowEnergyController) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QLowEnergyController) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QLowEnergyController) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QLowEnergyController) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QLowEnergyController) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QLowEnergyController) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QLowEnergyDescriptor struct {
	internal.Internal
}

type QLowEnergyDescriptor_ITF interface {
	QLowEnergyDescriptor_PTR() *QLowEnergyDescriptor
}

func (ptr *QLowEnergyDescriptor) QLowEnergyDescriptor_PTR() *QLowEnergyDescriptor {
	return ptr
}

func (ptr *QLowEnergyDescriptor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QLowEnergyDescriptor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQLowEnergyDescriptor(ptr QLowEnergyDescriptor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLowEnergyDescriptor_PTR().Pointer()
	}
	return nil
}

func (n *QLowEnergyDescriptor) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQLowEnergyDescriptorFromPointer(ptr unsafe.Pointer) (n *QLowEnergyDescriptor) {
	n = new(QLowEnergyDescriptor)
	n.InitFromInternal(uintptr(ptr), "bluetooth.QLowEnergyDescriptor")
	return
}
func NewQLowEnergyDescriptor() *QLowEnergyDescriptor {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQLowEnergyDescriptor", ""}).(*QLowEnergyDescriptor)
}

func NewQLowEnergyDescriptor2(other QLowEnergyDescriptor_ITF) *QLowEnergyDescriptor {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQLowEnergyDescriptor2", "", other}).(*QLowEnergyDescriptor)
}

func (ptr *QLowEnergyDescriptor) IsValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsValid"}).(bool)
}

func (ptr *QLowEnergyDescriptor) Name() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Name"}).(string)
}

func (ptr *QLowEnergyDescriptor) Type() QBluetoothUuid__DescriptorType {

	return QBluetoothUuid__DescriptorType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(float64))
}

func (ptr *QLowEnergyDescriptor) Uuid() *QBluetoothUuid {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Uuid"}).(*QBluetoothUuid)
}

func (ptr *QLowEnergyDescriptor) Value() *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Value"}).(*core.QByteArray)
}

func (ptr *QLowEnergyDescriptor) DestroyQLowEnergyDescriptor() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQLowEnergyDescriptor"})
}

type QLowEnergyDescriptorData struct {
	internal.Internal
}

type QLowEnergyDescriptorData_ITF interface {
	QLowEnergyDescriptorData_PTR() *QLowEnergyDescriptorData
}

func (ptr *QLowEnergyDescriptorData) QLowEnergyDescriptorData_PTR() *QLowEnergyDescriptorData {
	return ptr
}

func (ptr *QLowEnergyDescriptorData) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QLowEnergyDescriptorData) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQLowEnergyDescriptorData(ptr QLowEnergyDescriptorData_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLowEnergyDescriptorData_PTR().Pointer()
	}
	return nil
}

func (n *QLowEnergyDescriptorData) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQLowEnergyDescriptorDataFromPointer(ptr unsafe.Pointer) (n *QLowEnergyDescriptorData) {
	n = new(QLowEnergyDescriptorData)
	n.InitFromInternal(uintptr(ptr), "bluetooth.QLowEnergyDescriptorData")
	return
}
func NewQLowEnergyDescriptorData() *QLowEnergyDescriptorData {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQLowEnergyDescriptorData", ""}).(*QLowEnergyDescriptorData)
}

func NewQLowEnergyDescriptorData2(uuid QBluetoothUuid_ITF, value core.QByteArray_ITF) *QLowEnergyDescriptorData {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQLowEnergyDescriptorData2", "", uuid, value}).(*QLowEnergyDescriptorData)
}

func NewQLowEnergyDescriptorData3(other QLowEnergyDescriptorData_ITF) *QLowEnergyDescriptorData {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQLowEnergyDescriptorData3", "", other}).(*QLowEnergyDescriptorData)
}

func (ptr *QLowEnergyDescriptorData) IsReadable() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsReadable"}).(bool)
}

func (ptr *QLowEnergyDescriptorData) IsValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsValid"}).(bool)
}

func (ptr *QLowEnergyDescriptorData) IsWritable() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsWritable"}).(bool)
}

func (ptr *QLowEnergyDescriptorData) ReadConstraints() QBluetooth__AttAccessConstraint {

	return QBluetooth__AttAccessConstraint(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReadConstraints"}).(float64))
}

func (ptr *QLowEnergyDescriptorData) SetReadPermissions(readable bool, constraints QBluetooth__AttAccessConstraint) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetReadPermissions", readable, constraints})
}

func (ptr *QLowEnergyDescriptorData) SetUuid(uuid QBluetoothUuid_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetUuid", uuid})
}

func (ptr *QLowEnergyDescriptorData) SetValue(value core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetValue", value})
}

func (ptr *QLowEnergyDescriptorData) SetWritePermissions(writable bool, constraints QBluetooth__AttAccessConstraint) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWritePermissions", writable, constraints})
}

func (ptr *QLowEnergyDescriptorData) Swap(other QLowEnergyDescriptorData_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", other})
}

func (ptr *QLowEnergyDescriptorData) Uuid() *QBluetoothUuid {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Uuid"}).(*QBluetoothUuid)
}

func (ptr *QLowEnergyDescriptorData) Value() *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Value"}).(*core.QByteArray)
}

func (ptr *QLowEnergyDescriptorData) WriteConstraints() QBluetooth__AttAccessConstraint {

	return QBluetooth__AttAccessConstraint(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WriteConstraints"}).(float64))
}

func (ptr *QLowEnergyDescriptorData) DestroyQLowEnergyDescriptorData() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQLowEnergyDescriptorData"})
}

type QLowEnergyService struct {
	core.QObject
}

type QLowEnergyService_ITF interface {
	core.QObject_ITF
	QLowEnergyService_PTR() *QLowEnergyService
}

func (ptr *QLowEnergyService) QLowEnergyService_PTR() *QLowEnergyService {
	return ptr
}

func (ptr *QLowEnergyService) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QLowEnergyService) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQLowEnergyService(ptr QLowEnergyService_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLowEnergyService_PTR().Pointer()
	}
	return nil
}

func (n *QLowEnergyService) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QLowEnergyService) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQLowEnergyServiceFromPointer(ptr unsafe.Pointer) (n *QLowEnergyService) {
	n = new(QLowEnergyService)
	n.InitFromInternal(uintptr(ptr), "bluetooth.QLowEnergyService")
	return
}

//go:generate stringer -type=QLowEnergyService__ServiceType
//QLowEnergyService::ServiceType
type QLowEnergyService__ServiceType int64

const (
	QLowEnergyService__PrimaryService  QLowEnergyService__ServiceType = QLowEnergyService__ServiceType(0x0001)
	QLowEnergyService__IncludedService QLowEnergyService__ServiceType = QLowEnergyService__ServiceType(0x0002)
)

//go:generate stringer -type=QLowEnergyService__ServiceError
//QLowEnergyService::ServiceError
type QLowEnergyService__ServiceError int64

const (
	QLowEnergyService__NoError                  QLowEnergyService__ServiceError = QLowEnergyService__ServiceError(0)
	QLowEnergyService__OperationError           QLowEnergyService__ServiceError = QLowEnergyService__ServiceError(1)
	QLowEnergyService__CharacteristicWriteError QLowEnergyService__ServiceError = QLowEnergyService__ServiceError(2)
	QLowEnergyService__DescriptorWriteError     QLowEnergyService__ServiceError = QLowEnergyService__ServiceError(3)
	QLowEnergyService__UnknownError             QLowEnergyService__ServiceError = QLowEnergyService__ServiceError(4)
	QLowEnergyService__CharacteristicReadError  QLowEnergyService__ServiceError = QLowEnergyService__ServiceError(5)
	QLowEnergyService__DescriptorReadError      QLowEnergyService__ServiceError = QLowEnergyService__ServiceError(6)
)

//go:generate stringer -type=QLowEnergyService__ServiceState
//QLowEnergyService::ServiceState
type QLowEnergyService__ServiceState int64

const (
	QLowEnergyService__InvalidService      QLowEnergyService__ServiceState = QLowEnergyService__ServiceState(0)
	QLowEnergyService__DiscoveryRequired   QLowEnergyService__ServiceState = QLowEnergyService__ServiceState(1)
	QLowEnergyService__DiscoveringServices QLowEnergyService__ServiceState = QLowEnergyService__ServiceState(2)
	QLowEnergyService__ServiceDiscovered   QLowEnergyService__ServiceState = QLowEnergyService__ServiceState(3)
	QLowEnergyService__LocalService        QLowEnergyService__ServiceState = QLowEnergyService__ServiceState(4)
)

//go:generate stringer -type=QLowEnergyService__WriteMode
//QLowEnergyService::WriteMode
type QLowEnergyService__WriteMode int64

const (
	QLowEnergyService__WriteWithResponse    QLowEnergyService__WriteMode = QLowEnergyService__WriteMode(0)
	QLowEnergyService__WriteWithoutResponse QLowEnergyService__WriteMode = QLowEnergyService__WriteMode(1)
	QLowEnergyService__WriteSigned          QLowEnergyService__WriteMode = QLowEnergyService__WriteMode(2)
)

func (ptr *QLowEnergyService) Characteristic(uuid QBluetoothUuid_ITF) *QLowEnergyCharacteristic {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Characteristic", uuid}).(*QLowEnergyCharacteristic)
}

func (ptr *QLowEnergyService) ConnectCharacteristicChanged(f func(characteristic *QLowEnergyCharacteristic, newValue *core.QByteArray)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCharacteristicChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLowEnergyService) DisconnectCharacteristicChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCharacteristicChanged"})
}

func (ptr *QLowEnergyService) CharacteristicChanged(characteristic QLowEnergyCharacteristic_ITF, newValue core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CharacteristicChanged", characteristic, newValue})
}

func (ptr *QLowEnergyService) ConnectCharacteristicRead(f func(characteristic *QLowEnergyCharacteristic, value *core.QByteArray)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCharacteristicRead", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLowEnergyService) DisconnectCharacteristicRead() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCharacteristicRead"})
}

func (ptr *QLowEnergyService) CharacteristicRead(characteristic QLowEnergyCharacteristic_ITF, value core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CharacteristicRead", characteristic, value})
}

func (ptr *QLowEnergyService) ConnectCharacteristicWritten(f func(characteristic *QLowEnergyCharacteristic, newValue *core.QByteArray)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCharacteristicWritten", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLowEnergyService) DisconnectCharacteristicWritten() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCharacteristicWritten"})
}

func (ptr *QLowEnergyService) CharacteristicWritten(characteristic QLowEnergyCharacteristic_ITF, newValue core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CharacteristicWritten", characteristic, newValue})
}

func (ptr *QLowEnergyService) Characteristics() []*QLowEnergyCharacteristic {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Characteristics"}).([]*QLowEnergyCharacteristic)
}

func (ptr *QLowEnergyService) Contains(characteristic QLowEnergyCharacteristic_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Contains", characteristic}).(bool)
}

func (ptr *QLowEnergyService) Contains2(descriptor QLowEnergyDescriptor_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Contains2", descriptor}).(bool)
}

func (ptr *QLowEnergyService) ConnectDescriptorRead(f func(descriptor *QLowEnergyDescriptor, value *core.QByteArray)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDescriptorRead", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLowEnergyService) DisconnectDescriptorRead() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDescriptorRead"})
}

func (ptr *QLowEnergyService) DescriptorRead(descriptor QLowEnergyDescriptor_ITF, value core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DescriptorRead", descriptor, value})
}

func (ptr *QLowEnergyService) ConnectDescriptorWritten(f func(descriptor *QLowEnergyDescriptor, newValue *core.QByteArray)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDescriptorWritten", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLowEnergyService) DisconnectDescriptorWritten() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDescriptorWritten"})
}

func (ptr *QLowEnergyService) DescriptorWritten(descriptor QLowEnergyDescriptor_ITF, newValue core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DescriptorWritten", descriptor, newValue})
}

func (ptr *QLowEnergyService) DiscoverDetails() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DiscoverDetails"})
}

func (ptr *QLowEnergyService) Error() QLowEnergyService__ServiceError {

	return QLowEnergyService__ServiceError(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error"}).(float64))
}

func (ptr *QLowEnergyService) ConnectError2(f func(newError QLowEnergyService__ServiceError)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectError2", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLowEnergyService) DisconnectError2() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectError2"})
}

func (ptr *QLowEnergyService) Error2(newError QLowEnergyService__ServiceError) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error2", newError})
}

func (ptr *QLowEnergyService) IncludedServices() []*QBluetoothUuid {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IncludedServices"}).([]*QBluetoothUuid)
}

func (ptr *QLowEnergyService) ReadCharacteristic(characteristic QLowEnergyCharacteristic_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReadCharacteristic", characteristic})
}

func (ptr *QLowEnergyService) ReadDescriptor(descriptor QLowEnergyDescriptor_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReadDescriptor", descriptor})
}

func (ptr *QLowEnergyService) ServiceName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ServiceName"}).(string)
}

func (ptr *QLowEnergyService) ServiceUuid() *QBluetoothUuid {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ServiceUuid"}).(*QBluetoothUuid)
}

func (ptr *QLowEnergyService) State() QLowEnergyService__ServiceState {

	return QLowEnergyService__ServiceState(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "State"}).(float64))
}

func (ptr *QLowEnergyService) ConnectStateChanged(f func(newState QLowEnergyService__ServiceState)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStateChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLowEnergyService) DisconnectStateChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStateChanged"})
}

func (ptr *QLowEnergyService) StateChanged(newState QLowEnergyService__ServiceState) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StateChanged", newState})
}

func (ptr *QLowEnergyService) Type() QLowEnergyService__ServiceType {

	return QLowEnergyService__ServiceType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(float64))
}

func (ptr *QLowEnergyService) WriteCharacteristic(characteristic QLowEnergyCharacteristic_ITF, newValue core.QByteArray_ITF, mode QLowEnergyService__WriteMode) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WriteCharacteristic", characteristic, newValue, mode})
}

func (ptr *QLowEnergyService) WriteDescriptor(descriptor QLowEnergyDescriptor_ITF, newValue core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WriteDescriptor", descriptor, newValue})
}

func (ptr *QLowEnergyService) ConnectDestroyQLowEnergyService(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQLowEnergyService", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLowEnergyService) DisconnectDestroyQLowEnergyService() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQLowEnergyService"})
}

func (ptr *QLowEnergyService) DestroyQLowEnergyService() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQLowEnergyService"})
}

func (ptr *QLowEnergyService) DestroyQLowEnergyServiceDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQLowEnergyServiceDefault"})
}

func (ptr *QLowEnergyService) __characteristics_atList(i int) *QLowEnergyCharacteristic {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__characteristics_atList", i}).(*QLowEnergyCharacteristic)
}

func (ptr *QLowEnergyService) __characteristics_setList(i QLowEnergyCharacteristic_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__characteristics_setList", i})
}

func (ptr *QLowEnergyService) __characteristics_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__characteristics_newList"}).(unsafe.Pointer)
}

func (ptr *QLowEnergyService) __includedServices_atList(i int) *QBluetoothUuid {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__includedServices_atList", i}).(*QBluetoothUuid)
}

func (ptr *QLowEnergyService) __includedServices_setList(i QBluetoothUuid_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__includedServices_setList", i})
}

func (ptr *QLowEnergyService) __includedServices_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__includedServices_newList"}).(unsafe.Pointer)
}

func (ptr *QLowEnergyService) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QLowEnergyService) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QLowEnergyService) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QLowEnergyService) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QLowEnergyService) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QLowEnergyService) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QLowEnergyService) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QLowEnergyService) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QLowEnergyService) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QLowEnergyService) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QLowEnergyService) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QLowEnergyService) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QLowEnergyService) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QLowEnergyService) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QLowEnergyService) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QLowEnergyService) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QLowEnergyService) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QLowEnergyService) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QLowEnergyService) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QLowEnergyService) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QLowEnergyService) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QLowEnergyServiceData struct {
	internal.Internal
}

type QLowEnergyServiceData_ITF interface {
	QLowEnergyServiceData_PTR() *QLowEnergyServiceData
}

func (ptr *QLowEnergyServiceData) QLowEnergyServiceData_PTR() *QLowEnergyServiceData {
	return ptr
}

func (ptr *QLowEnergyServiceData) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QLowEnergyServiceData) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQLowEnergyServiceData(ptr QLowEnergyServiceData_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLowEnergyServiceData_PTR().Pointer()
	}
	return nil
}

func (n *QLowEnergyServiceData) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQLowEnergyServiceDataFromPointer(ptr unsafe.Pointer) (n *QLowEnergyServiceData) {
	n = new(QLowEnergyServiceData)
	n.InitFromInternal(uintptr(ptr), "bluetooth.QLowEnergyServiceData")
	return
}

//go:generate stringer -type=QLowEnergyServiceData__ServiceType
//QLowEnergyServiceData::ServiceType
type QLowEnergyServiceData__ServiceType int64

const (
	QLowEnergyServiceData__ServiceTypePrimary   QLowEnergyServiceData__ServiceType = QLowEnergyServiceData__ServiceType(0x2800)
	QLowEnergyServiceData__ServiceTypeSecondary QLowEnergyServiceData__ServiceType = QLowEnergyServiceData__ServiceType(0x2801)
)

func NewQLowEnergyServiceData() *QLowEnergyServiceData {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQLowEnergyServiceData", ""}).(*QLowEnergyServiceData)
}

func NewQLowEnergyServiceData2(other QLowEnergyServiceData_ITF) *QLowEnergyServiceData {

	return internal.CallLocalFunction([]interface{}{"", "", "bluetooth.NewQLowEnergyServiceData2", "", other}).(*QLowEnergyServiceData)
}

func (ptr *QLowEnergyServiceData) AddCharacteristic(characteristic QLowEnergyCharacteristicData_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddCharacteristic", characteristic})
}

func (ptr *QLowEnergyServiceData) AddIncludedService(service QLowEnergyService_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddIncludedService", service})
}

func (ptr *QLowEnergyServiceData) Characteristics() []*QLowEnergyCharacteristicData {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Characteristics"}).([]*QLowEnergyCharacteristicData)
}

func (ptr *QLowEnergyServiceData) IncludedServices() []*QLowEnergyService {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IncludedServices"}).([]*QLowEnergyService)
}

func (ptr *QLowEnergyServiceData) IsValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsValid"}).(bool)
}

func (ptr *QLowEnergyServiceData) SetCharacteristics(characteristics []*QLowEnergyCharacteristicData) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetCharacteristics", characteristics})
}

func (ptr *QLowEnergyServiceData) SetIncludedServices(services []*QLowEnergyService) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetIncludedServices", services})
}

func (ptr *QLowEnergyServiceData) SetType(ty QLowEnergyServiceData__ServiceType) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetType", ty})
}

func (ptr *QLowEnergyServiceData) SetUuid(uuid QBluetoothUuid_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetUuid", uuid})
}

func (ptr *QLowEnergyServiceData) Swap(other QLowEnergyServiceData_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", other})
}

func (ptr *QLowEnergyServiceData) Type() QLowEnergyServiceData__ServiceType {

	return QLowEnergyServiceData__ServiceType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(float64))
}

func (ptr *QLowEnergyServiceData) Uuid() *QBluetoothUuid {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Uuid"}).(*QBluetoothUuid)
}

func (ptr *QLowEnergyServiceData) DestroyQLowEnergyServiceData() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQLowEnergyServiceData"})
}

func (ptr *QLowEnergyServiceData) __characteristics_atList(i int) *QLowEnergyCharacteristicData {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__characteristics_atList", i}).(*QLowEnergyCharacteristicData)
}

func (ptr *QLowEnergyServiceData) __characteristics_setList(i QLowEnergyCharacteristicData_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__characteristics_setList", i})
}

func (ptr *QLowEnergyServiceData) __characteristics_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__characteristics_newList"}).(unsafe.Pointer)
}

func (ptr *QLowEnergyServiceData) __includedServices_atList(i int) *QLowEnergyService {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__includedServices_atList", i}).(*QLowEnergyService)
}

func (ptr *QLowEnergyServiceData) __includedServices_setList(i QLowEnergyService_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__includedServices_setList", i})
}

func (ptr *QLowEnergyServiceData) __includedServices_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__includedServices_newList"}).(unsafe.Pointer)
}

func (ptr *QLowEnergyServiceData) __setCharacteristics_characteristics_atList(i int) *QLowEnergyCharacteristicData {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setCharacteristics_characteristics_atList", i}).(*QLowEnergyCharacteristicData)
}

func (ptr *QLowEnergyServiceData) __setCharacteristics_characteristics_setList(i QLowEnergyCharacteristicData_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setCharacteristics_characteristics_setList", i})
}

func (ptr *QLowEnergyServiceData) __setCharacteristics_characteristics_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setCharacteristics_characteristics_newList"}).(unsafe.Pointer)
}

func (ptr *QLowEnergyServiceData) __setIncludedServices_services_atList(i int) *QLowEnergyService {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setIncludedServices_services_atList", i}).(*QLowEnergyService)
}

func (ptr *QLowEnergyServiceData) __setIncludedServices_services_setList(i QLowEnergyService_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setIncludedServices_services_setList", i})
}

func (ptr *QLowEnergyServiceData) __setIncludedServices_services_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setIncludedServices_services_newList"}).(unsafe.Pointer)
}

func init() {
	internal.ConstructorTable["bluetooth.QBluetooth"] = NewQBluetoothFromPointer
	internal.ConstructorTable["bluetooth.QBluetoothAddress"] = NewQBluetoothAddressFromPointer
	internal.ConstructorTable["bluetooth.QBluetoothDeviceDiscoveryAgent"] = NewQBluetoothDeviceDiscoveryAgentFromPointer
	internal.ConstructorTable["bluetooth.QBluetoothDeviceInfo"] = NewQBluetoothDeviceInfoFromPointer
	internal.ConstructorTable["bluetooth.QBluetoothHostInfo"] = NewQBluetoothHostInfoFromPointer
	internal.ConstructorTable["bluetooth.QBluetoothLocalDevice"] = NewQBluetoothLocalDeviceFromPointer
	internal.ConstructorTable["bluetooth.QBluetoothServer"] = NewQBluetoothServerFromPointer
	internal.ConstructorTable["bluetooth.QBluetoothServiceDiscoveryAgent"] = NewQBluetoothServiceDiscoveryAgentFromPointer
	internal.ConstructorTable["bluetooth.QBluetoothServiceInfo"] = NewQBluetoothServiceInfoFromPointer
	internal.ConstructorTable["bluetooth.QBluetoothSocket"] = NewQBluetoothSocketFromPointer
	internal.ConstructorTable["bluetooth.QBluetoothTransferManager"] = NewQBluetoothTransferManagerFromPointer
	internal.ConstructorTable["bluetooth.QBluetoothTransferReply"] = NewQBluetoothTransferReplyFromPointer
	internal.ConstructorTable["bluetooth.QBluetoothTransferRequest"] = NewQBluetoothTransferRequestFromPointer
	internal.ConstructorTable["bluetooth.QBluetoothUuid"] = NewQBluetoothUuidFromPointer
	internal.ConstructorTable["bluetooth.QLowEnergyAdvertisingData"] = NewQLowEnergyAdvertisingDataFromPointer
	internal.ConstructorTable["bluetooth.QLowEnergyAdvertisingParameters"] = NewQLowEnergyAdvertisingParametersFromPointer
	internal.ConstructorTable["bluetooth.QLowEnergyCharacteristic"] = NewQLowEnergyCharacteristicFromPointer
	internal.ConstructorTable["bluetooth.QLowEnergyCharacteristicData"] = NewQLowEnergyCharacteristicDataFromPointer
	internal.ConstructorTable["bluetooth.QLowEnergyConnectionParameters"] = NewQLowEnergyConnectionParametersFromPointer
	internal.ConstructorTable["bluetooth.QLowEnergyController"] = NewQLowEnergyControllerFromPointer
	internal.ConstructorTable["bluetooth.QLowEnergyDescriptor"] = NewQLowEnergyDescriptorFromPointer
	internal.ConstructorTable["bluetooth.QLowEnergyDescriptorData"] = NewQLowEnergyDescriptorDataFromPointer
	internal.ConstructorTable["bluetooth.QLowEnergyService"] = NewQLowEnergyServiceFromPointer
	internal.ConstructorTable["bluetooth.QLowEnergyServiceData"] = NewQLowEnergyServiceDataFromPointer
}
