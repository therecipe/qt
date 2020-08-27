// +build !minimal

package serialbus

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/internal"
	"github.com/therecipe/qt/network"
	"unsafe"
)

type QCanBus struct {
	core.QObject
}

type QCanBus_ITF interface {
	core.QObject_ITF
	QCanBus_PTR() *QCanBus
}

func (ptr *QCanBus) QCanBus_PTR() *QCanBus {
	return ptr
}

func (ptr *QCanBus) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QCanBus) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQCanBus(ptr QCanBus_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCanBus_PTR().Pointer()
	}
	return nil
}

func (n *QCanBus) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QCanBus) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQCanBusFromPointer(ptr unsafe.Pointer) (n *QCanBus) {
	n = new(QCanBus)
	n.InitFromInternal(uintptr(ptr), "serialbus.QCanBus")
	return
}

func (ptr *QCanBus) DestroyQCanBus() {
}

func (ptr *QCanBus) AvailableDevices(plugin string, errorMessage string) []*QCanBusDeviceInfo {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AvailableDevices", plugin, errorMessage}).([]*QCanBusDeviceInfo)
}

func (ptr *QCanBus) CreateDevice(plugin string, interfaceName string, errorMessage string) *QCanBusDevice {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateDevice", plugin, interfaceName, errorMessage}).(*QCanBusDevice)
}

func QCanBus_Instance() *QCanBus {

	return internal.CallLocalFunction([]interface{}{"", "", "serialbus.QCanBus_Instance", ""}).(*QCanBus)
}

func (ptr *QCanBus) Instance() *QCanBus {

	return internal.CallLocalFunction([]interface{}{"", "", "serialbus.QCanBus_Instance", ""}).(*QCanBus)
}

func (ptr *QCanBus) Plugins() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Plugins"}).([]string)
}

func (ptr *QCanBus) __availableDevices_atList(i int) *QCanBusDeviceInfo {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__availableDevices_atList", i}).(*QCanBusDeviceInfo)
}

func (ptr *QCanBus) __availableDevices_setList(i QCanBusDeviceInfo_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__availableDevices_setList", i})
}

func (ptr *QCanBus) __availableDevices_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__availableDevices_newList"}).(unsafe.Pointer)
}

func (ptr *QCanBus) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QCanBus) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QCanBus) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QCanBus) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QCanBus) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QCanBus) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QCanBus) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QCanBus) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QCanBus) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QCanBus) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QCanBus) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QCanBus) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QCanBus) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QCanBus) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QCanBus) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QCanBus) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QCanBus) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QCanBus) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QCanBus) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QCanBus) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QCanBus) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QCanBusDevice struct {
	core.QObject
}

type QCanBusDevice_ITF interface {
	core.QObject_ITF
	QCanBusDevice_PTR() *QCanBusDevice
}

func (ptr *QCanBusDevice) QCanBusDevice_PTR() *QCanBusDevice {
	return ptr
}

func (ptr *QCanBusDevice) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QCanBusDevice) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQCanBusDevice(ptr QCanBusDevice_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCanBusDevice_PTR().Pointer()
	}
	return nil
}

func (n *QCanBusDevice) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QCanBusDevice) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQCanBusDeviceFromPointer(ptr unsafe.Pointer) (n *QCanBusDevice) {
	n = new(QCanBusDevice)
	n.InitFromInternal(uintptr(ptr), "serialbus.QCanBusDevice")
	return
}

func (ptr *QCanBusDevice) DestroyQCanBusDevice() {
}

//go:generate stringer -type=QCanBusDevice__CanBusError
//QCanBusDevice::CanBusError
type QCanBusDevice__CanBusError int64

const (
	QCanBusDevice__NoError            QCanBusDevice__CanBusError = QCanBusDevice__CanBusError(0)
	QCanBusDevice__ReadError          QCanBusDevice__CanBusError = QCanBusDevice__CanBusError(1)
	QCanBusDevice__WriteError         QCanBusDevice__CanBusError = QCanBusDevice__CanBusError(2)
	QCanBusDevice__ConnectionError    QCanBusDevice__CanBusError = QCanBusDevice__CanBusError(3)
	QCanBusDevice__ConfigurationError QCanBusDevice__CanBusError = QCanBusDevice__CanBusError(4)
	QCanBusDevice__UnknownError       QCanBusDevice__CanBusError = QCanBusDevice__CanBusError(5)
)

//go:generate stringer -type=QCanBusDevice__CanBusDeviceState
//QCanBusDevice::CanBusDeviceState
type QCanBusDevice__CanBusDeviceState int64

const (
	QCanBusDevice__UnconnectedState QCanBusDevice__CanBusDeviceState = QCanBusDevice__CanBusDeviceState(0)
	QCanBusDevice__ConnectingState  QCanBusDevice__CanBusDeviceState = QCanBusDevice__CanBusDeviceState(1)
	QCanBusDevice__ConnectedState   QCanBusDevice__CanBusDeviceState = QCanBusDevice__CanBusDeviceState(2)
	QCanBusDevice__ClosingState     QCanBusDevice__CanBusDeviceState = QCanBusDevice__CanBusDeviceState(3)
)

//go:generate stringer -type=QCanBusDevice__ConfigurationKey
//QCanBusDevice::ConfigurationKey
type QCanBusDevice__ConfigurationKey int64

const (
	QCanBusDevice__RawFilterKey   QCanBusDevice__ConfigurationKey = QCanBusDevice__ConfigurationKey(0)
	QCanBusDevice__ErrorFilterKey QCanBusDevice__ConfigurationKey = QCanBusDevice__ConfigurationKey(1)
	QCanBusDevice__LoopbackKey    QCanBusDevice__ConfigurationKey = QCanBusDevice__ConfigurationKey(2)
	QCanBusDevice__ReceiveOwnKey  QCanBusDevice__ConfigurationKey = QCanBusDevice__ConfigurationKey(3)
	QCanBusDevice__BitRateKey     QCanBusDevice__ConfigurationKey = QCanBusDevice__ConfigurationKey(4)
	QCanBusDevice__CanFdKey       QCanBusDevice__ConfigurationKey = QCanBusDevice__ConfigurationKey(5)
	QCanBusDevice__DataBitRateKey QCanBusDevice__ConfigurationKey = QCanBusDevice__ConfigurationKey(6)
	QCanBusDevice__UserKey        QCanBusDevice__ConfigurationKey = QCanBusDevice__ConfigurationKey(30)
)

//go:generate stringer -type=QCanBusDevice__Direction
//QCanBusDevice::Direction
type QCanBusDevice__Direction int64

const (
	QCanBusDevice__Input         QCanBusDevice__Direction = QCanBusDevice__Direction(1)
	QCanBusDevice__Output        QCanBusDevice__Direction = QCanBusDevice__Direction(2)
	QCanBusDevice__AllDirections QCanBusDevice__Direction = QCanBusDevice__Direction(QCanBusDevice__Input | QCanBusDevice__Output)
)

func NewQCanBusDevice2(parent core.QObject_ITF) *QCanBusDevice {

	return internal.CallLocalFunction([]interface{}{"", "", "serialbus.NewQCanBusDevice2", "", parent}).(*QCanBusDevice)
}

func (ptr *QCanBusDevice) Clear(direction QCanBusDevice__Direction) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Clear", direction})
}

func (ptr *QCanBusDevice) ConnectClose(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectClose", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCanBusDevice) DisconnectClose() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectClose"})
}

func (ptr *QCanBusDevice) Close() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Close"})
}

func (ptr *QCanBusDevice) ConfigurationKeys() []int {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConfigurationKeys"}).([]int)
}

func (ptr *QCanBusDevice) ConfigurationParameter(key int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConfigurationParameter", key}).(*core.QVariant)
}

func (ptr *QCanBusDevice) ConnectDevice() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDevice"}).(bool)
}

func (ptr *QCanBusDevice) DisconnectDevice() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDevice"})
}

func (ptr *QCanBusDevice) EnqueueOutgoingFrame(newFrame QCanBusFrame_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EnqueueOutgoingFrame", newFrame})
}

func (ptr *QCanBusDevice) EnqueueReceivedFrames(newFrames []*QCanBusFrame) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EnqueueReceivedFrames", newFrames})
}

func (ptr *QCanBusDevice) Error() QCanBusDevice__CanBusError {

	return QCanBusDevice__CanBusError(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error"}).(float64))
}

func (ptr *QCanBusDevice) ConnectErrorOccurred(f func(vqc QCanBusDevice__CanBusError)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectErrorOccurred", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCanBusDevice) DisconnectErrorOccurred() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectErrorOccurred"})
}

func (ptr *QCanBusDevice) ErrorOccurred(vqc QCanBusDevice__CanBusError) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ErrorOccurred", vqc})
}

func (ptr *QCanBusDevice) ErrorString() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ErrorString"}).(string)
}

func (ptr *QCanBusDevice) FramesAvailable() int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FramesAvailable"}).(float64))
}

func (ptr *QCanBusDevice) ConnectFramesReceived(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFramesReceived", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCanBusDevice) DisconnectFramesReceived() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFramesReceived"})
}

func (ptr *QCanBusDevice) FramesReceived() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FramesReceived"})
}

func (ptr *QCanBusDevice) FramesToWrite() int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FramesToWrite"}).(float64))
}

func (ptr *QCanBusDevice) ConnectFramesWritten(f func(framesCount int64)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFramesWritten", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCanBusDevice) DisconnectFramesWritten() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFramesWritten"})
}

func (ptr *QCanBusDevice) FramesWritten(framesCount int64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FramesWritten", framesCount})
}

func (ptr *QCanBusDevice) HasOutgoingFrames() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasOutgoingFrames"}).(bool)
}

func (ptr *QCanBusDevice) ConnectInterpretErrorFrame(f func(frame *QCanBusFrame) string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectInterpretErrorFrame", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCanBusDevice) DisconnectInterpretErrorFrame() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectInterpretErrorFrame"})
}

func (ptr *QCanBusDevice) InterpretErrorFrame(frame QCanBusFrame_ITF) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InterpretErrorFrame", frame}).(string)
}

func (ptr *QCanBusDevice) ConnectOpen(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectOpen", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCanBusDevice) DisconnectOpen() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectOpen"})
}

func (ptr *QCanBusDevice) Open() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Open"}).(bool)
}

func (ptr *QCanBusDevice) ConnectSetConfigurationParameter(f func(key int, value *core.QVariant)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetConfigurationParameter", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCanBusDevice) DisconnectSetConfigurationParameter() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetConfigurationParameter"})
}

func (ptr *QCanBusDevice) SetConfigurationParameter(key int, value core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetConfigurationParameter", key, value})
}

func (ptr *QCanBusDevice) SetConfigurationParameterDefault(key int, value core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetConfigurationParameterDefault", key, value})
}

func (ptr *QCanBusDevice) SetError(errorText string, errorId QCanBusDevice__CanBusError) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetError", errorText, errorId})
}

func (ptr *QCanBusDevice) SetState(newState QCanBusDevice__CanBusDeviceState) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetState", newState})
}

func (ptr *QCanBusDevice) State() QCanBusDevice__CanBusDeviceState {

	return QCanBusDevice__CanBusDeviceState(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "State"}).(float64))
}

func (ptr *QCanBusDevice) ConnectStateChanged(f func(state QCanBusDevice__CanBusDeviceState)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStateChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCanBusDevice) DisconnectStateChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStateChanged"})
}

func (ptr *QCanBusDevice) StateChanged(state QCanBusDevice__CanBusDeviceState) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StateChanged", state})
}

func (ptr *QCanBusDevice) ConnectWaitForFramesReceived(f func(msecs int) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectWaitForFramesReceived", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCanBusDevice) DisconnectWaitForFramesReceived() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectWaitForFramesReceived"})
}

func (ptr *QCanBusDevice) WaitForFramesReceived(msecs int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WaitForFramesReceived", msecs}).(bool)
}

func (ptr *QCanBusDevice) WaitForFramesReceivedDefault(msecs int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WaitForFramesReceivedDefault", msecs}).(bool)
}

func (ptr *QCanBusDevice) ConnectWaitForFramesWritten(f func(msecs int) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectWaitForFramesWritten", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCanBusDevice) DisconnectWaitForFramesWritten() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectWaitForFramesWritten"})
}

func (ptr *QCanBusDevice) WaitForFramesWritten(msecs int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WaitForFramesWritten", msecs}).(bool)
}

func (ptr *QCanBusDevice) WaitForFramesWrittenDefault(msecs int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WaitForFramesWrittenDefault", msecs}).(bool)
}

func (ptr *QCanBusDevice) ConnectWriteFrame(f func(frame *QCanBusFrame) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectWriteFrame", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCanBusDevice) DisconnectWriteFrame() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectWriteFrame"})
}

func (ptr *QCanBusDevice) WriteFrame(frame QCanBusFrame_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WriteFrame", frame}).(bool)
}

func (ptr *QCanBusDevice) __configurationKeys_atList(i int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__configurationKeys_atList", i}).(float64))
}

func (ptr *QCanBusDevice) __configurationKeys_setList(i int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__configurationKeys_setList", i})
}

func (ptr *QCanBusDevice) __configurationKeys_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__configurationKeys_newList"}).(unsafe.Pointer)
}

func (ptr *QCanBusDevice) __enqueueReceivedFrames_newFrames_setList(i QCanBusFrame_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__enqueueReceivedFrames_newFrames_setList", i})
}

func (ptr *QCanBusDevice) __enqueueReceivedFrames_newFrames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__enqueueReceivedFrames_newFrames_newList"}).(unsafe.Pointer)
}

func (ptr *QCanBusDevice) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QCanBusDevice) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QCanBusDevice) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QCanBusDevice) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QCanBusDevice) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QCanBusDevice) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QCanBusDevice) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QCanBusDevice) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QCanBusDevice) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QCanBusDevice) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QCanBusDevice) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QCanBusDevice) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QCanBusDevice) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QCanBusDevice) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QCanBusDevice) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QCanBusDevice) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QCanBusDevice) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QCanBusDevice) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QCanBusDevice) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QCanBusDevice) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QCanBusDevice) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QCanBusDeviceInfo struct {
	internal.Internal
}

type QCanBusDeviceInfo_ITF interface {
	QCanBusDeviceInfo_PTR() *QCanBusDeviceInfo
}

func (ptr *QCanBusDeviceInfo) QCanBusDeviceInfo_PTR() *QCanBusDeviceInfo {
	return ptr
}

func (ptr *QCanBusDeviceInfo) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QCanBusDeviceInfo) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQCanBusDeviceInfo(ptr QCanBusDeviceInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCanBusDeviceInfo_PTR().Pointer()
	}
	return nil
}

func (n *QCanBusDeviceInfo) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQCanBusDeviceInfoFromPointer(ptr unsafe.Pointer) (n *QCanBusDeviceInfo) {
	n = new(QCanBusDeviceInfo)
	n.InitFromInternal(uintptr(ptr), "serialbus.QCanBusDeviceInfo")
	return
}
func (ptr *QCanBusDeviceInfo) Channel() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Channel"}).(float64))
}

func (ptr *QCanBusDeviceInfo) Description() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Description"}).(string)
}

func (ptr *QCanBusDeviceInfo) HasFlexibleDataRate() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasFlexibleDataRate"}).(bool)
}

func (ptr *QCanBusDeviceInfo) IsVirtual() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsVirtual"}).(bool)
}

func (ptr *QCanBusDeviceInfo) Name() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Name"}).(string)
}

func (ptr *QCanBusDeviceInfo) SerialNumber() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SerialNumber"}).(string)
}

func (ptr *QCanBusDeviceInfo) Swap(other QCanBusDeviceInfo_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", other})
}

func (ptr *QCanBusDeviceInfo) DestroyQCanBusDeviceInfo() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQCanBusDeviceInfo"})
}

type QCanBusFactory struct {
	internal.Internal
}

type QCanBusFactory_ITF interface {
	QCanBusFactory_PTR() *QCanBusFactory
}

func (ptr *QCanBusFactory) QCanBusFactory_PTR() *QCanBusFactory {
	return ptr
}

func (ptr *QCanBusFactory) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QCanBusFactory) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQCanBusFactory(ptr QCanBusFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCanBusFactory_PTR().Pointer()
	}
	return nil
}

func (n *QCanBusFactory) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQCanBusFactoryFromPointer(ptr unsafe.Pointer) (n *QCanBusFactory) {
	n = new(QCanBusFactory)
	n.InitFromInternal(uintptr(ptr), "serialbus.QCanBusFactory")
	return
}

func (ptr *QCanBusFactory) DestroyQCanBusFactory() {
}

func (ptr *QCanBusFactory) ConnectCreateDevice(f func(interfaceName string, errorMessage string) *QCanBusDevice) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCreateDevice", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCanBusFactory) DisconnectCreateDevice() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCreateDevice"})
}

func (ptr *QCanBusFactory) CreateDevice(interfaceName string, errorMessage string) *QCanBusDevice {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateDevice", interfaceName, errorMessage}).(*QCanBusDevice)
}

type QCanBusFactoryV2 struct {
	QCanBusFactory
}

type QCanBusFactoryV2_ITF interface {
	QCanBusFactory_ITF
	QCanBusFactoryV2_PTR() *QCanBusFactoryV2
}

func (ptr *QCanBusFactoryV2) QCanBusFactoryV2_PTR() *QCanBusFactoryV2 {
	return ptr
}

func (ptr *QCanBusFactoryV2) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QCanBusFactory_PTR().Pointer()
	}
	return nil
}

func (ptr *QCanBusFactoryV2) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QCanBusFactory_PTR().SetPointer(p)
	}
}

func PointerFromQCanBusFactoryV2(ptr QCanBusFactoryV2_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCanBusFactoryV2_PTR().Pointer()
	}
	return nil
}

func (n *QCanBusFactoryV2) InitFromInternal(ptr uintptr, name string) {
	n.QCanBusFactory_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QCanBusFactoryV2) ClassNameInternalF() string {
	return n.QCanBusFactory_PTR().ClassNameInternalF()
}

func NewQCanBusFactoryV2FromPointer(ptr unsafe.Pointer) (n *QCanBusFactoryV2) {
	n = new(QCanBusFactoryV2)
	n.InitFromInternal(uintptr(ptr), "serialbus.QCanBusFactoryV2")
	return
}

func (ptr *QCanBusFactoryV2) DestroyQCanBusFactoryV2() {
}

func (ptr *QCanBusFactoryV2) ConnectAvailableDevices(f func(errorMessage string) []*QCanBusDeviceInfo) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAvailableDevices", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCanBusFactoryV2) DisconnectAvailableDevices() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAvailableDevices"})
}

func (ptr *QCanBusFactoryV2) AvailableDevices(errorMessage string) []*QCanBusDeviceInfo {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AvailableDevices", errorMessage}).([]*QCanBusDeviceInfo)
}

func (ptr *QCanBusFactoryV2) ConnectCreateDevice(f func(interfaceName string, errorMessage string) *QCanBusDevice) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCreateDevice", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QCanBusFactoryV2) DisconnectCreateDevice() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCreateDevice"})
}

func (ptr *QCanBusFactoryV2) CreateDevice(interfaceName string, errorMessage string) *QCanBusDevice {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateDevice", interfaceName, errorMessage}).(*QCanBusDevice)
}

func (ptr *QCanBusFactoryV2) __availableDevices_atList(i int) *QCanBusDeviceInfo {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__availableDevices_atList", i}).(*QCanBusDeviceInfo)
}

func (ptr *QCanBusFactoryV2) __availableDevices_setList(i QCanBusDeviceInfo_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__availableDevices_setList", i})
}

func (ptr *QCanBusFactoryV2) __availableDevices_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__availableDevices_newList"}).(unsafe.Pointer)
}

type QCanBusFrame struct {
	internal.Internal
}

type QCanBusFrame_ITF interface {
	QCanBusFrame_PTR() *QCanBusFrame
}

func (ptr *QCanBusFrame) QCanBusFrame_PTR() *QCanBusFrame {
	return ptr
}

func (ptr *QCanBusFrame) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QCanBusFrame) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQCanBusFrame(ptr QCanBusFrame_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCanBusFrame_PTR().Pointer()
	}
	return nil
}

func (n *QCanBusFrame) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQCanBusFrameFromPointer(ptr unsafe.Pointer) (n *QCanBusFrame) {
	n = new(QCanBusFrame)
	n.InitFromInternal(uintptr(ptr), "serialbus.QCanBusFrame")
	return
}

func (ptr *QCanBusFrame) DestroyQCanBusFrame() {
}

//go:generate stringer -type=QCanBusFrame__FrameType
//QCanBusFrame::FrameType
type QCanBusFrame__FrameType int64

const (
	QCanBusFrame__UnknownFrame       QCanBusFrame__FrameType = QCanBusFrame__FrameType(0x0)
	QCanBusFrame__DataFrame          QCanBusFrame__FrameType = QCanBusFrame__FrameType(0x1)
	QCanBusFrame__ErrorFrame         QCanBusFrame__FrameType = QCanBusFrame__FrameType(0x2)
	QCanBusFrame__RemoteRequestFrame QCanBusFrame__FrameType = QCanBusFrame__FrameType(0x3)
	QCanBusFrame__InvalidFrame       QCanBusFrame__FrameType = QCanBusFrame__FrameType(0x4)
)

//go:generate stringer -type=QCanBusFrame__FrameError
//QCanBusFrame::FrameError
type QCanBusFrame__FrameError int64

var (
	QCanBusFrame__NoError                    QCanBusFrame__FrameError = QCanBusFrame__FrameError(0)
	QCanBusFrame__TransmissionTimeoutError   QCanBusFrame__FrameError = QCanBusFrame__FrameError(0)
	QCanBusFrame__LostArbitrationError       QCanBusFrame__FrameError = QCanBusFrame__FrameError(0)
	QCanBusFrame__ControllerError            QCanBusFrame__FrameError = QCanBusFrame__FrameError(0)
	QCanBusFrame__ProtocolViolationError     QCanBusFrame__FrameError = QCanBusFrame__FrameError(0)
	QCanBusFrame__TransceiverError           QCanBusFrame__FrameError = QCanBusFrame__FrameError(0)
	QCanBusFrame__MissingAcknowledgmentError QCanBusFrame__FrameError = QCanBusFrame__FrameError(0)
	QCanBusFrame__BusOffError                QCanBusFrame__FrameError = QCanBusFrame__FrameError(0)
	QCanBusFrame__BusError                   QCanBusFrame__FrameError = QCanBusFrame__FrameError(0)
	QCanBusFrame__ControllerRestartError     QCanBusFrame__FrameError = QCanBusFrame__FrameError(0)
	QCanBusFrame__UnknownError               QCanBusFrame__FrameError = QCanBusFrame__FrameError(0)
	QCanBusFrame__AnyError                   QCanBusFrame__FrameError = QCanBusFrame__FrameError(0)
)

func NewQCanBusFrame(ty QCanBusFrame__FrameType) *QCanBusFrame {

	return internal.CallLocalFunction([]interface{}{"", "", "serialbus.NewQCanBusFrame", "", ty}).(*QCanBusFrame)
}

func NewQCanBusFrame2(identifier uint, data core.QByteArray_ITF) *QCanBusFrame {

	return internal.CallLocalFunction([]interface{}{"", "", "serialbus.NewQCanBusFrame2", "", identifier, data}).(*QCanBusFrame)
}

func (ptr *QCanBusFrame) Error() QCanBusFrame__FrameError {

	return QCanBusFrame__FrameError(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error"}).(float64))
}

func (ptr *QCanBusFrame) FrameId() uint {

	return uint(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FrameId"}).(float64))
}

func (ptr *QCanBusFrame) FrameType() QCanBusFrame__FrameType {

	return QCanBusFrame__FrameType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FrameType"}).(float64))
}

func (ptr *QCanBusFrame) HasBitrateSwitch() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasBitrateSwitch"}).(bool)
}

func (ptr *QCanBusFrame) HasErrorStateIndicator() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasErrorStateIndicator"}).(bool)
}

func (ptr *QCanBusFrame) HasExtendedFrameFormat() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasExtendedFrameFormat"}).(bool)
}

func (ptr *QCanBusFrame) HasFlexibleDataRateFormat() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasFlexibleDataRateFormat"}).(bool)
}

func (ptr *QCanBusFrame) HasLocalEcho() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasLocalEcho"}).(bool)
}

func (ptr *QCanBusFrame) IsValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsValid"}).(bool)
}

func (ptr *QCanBusFrame) Payload() *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Payload"}).(*core.QByteArray)
}

func (ptr *QCanBusFrame) SetBitrateSwitch(bitrateSwitch bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBitrateSwitch", bitrateSwitch})
}

func (ptr *QCanBusFrame) SetError(error QCanBusFrame__FrameError) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetError", error})
}

func (ptr *QCanBusFrame) SetErrorStateIndicator(errorStateIndicator bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetErrorStateIndicator", errorStateIndicator})
}

func (ptr *QCanBusFrame) SetExtendedFrameFormat(isExtended bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetExtendedFrameFormat", isExtended})
}

func (ptr *QCanBusFrame) SetFlexibleDataRateFormat(isFlexibleData bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFlexibleDataRateFormat", isFlexibleData})
}

func (ptr *QCanBusFrame) SetFrameId(newFrameId uint) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFrameId", newFrameId})
}

func (ptr *QCanBusFrame) SetFrameType(newType QCanBusFrame__FrameType) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFrameType", newType})
}

func (ptr *QCanBusFrame) SetLocalEcho(echo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLocalEcho", echo})
}

func (ptr *QCanBusFrame) SetPayload(data core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPayload", data})
}

func (ptr *QCanBusFrame) ToString() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToString"}).(string)
}

type QModbusClient struct {
	QModbusDevice
}

type QModbusClient_ITF interface {
	QModbusDevice_ITF
	QModbusClient_PTR() *QModbusClient
}

func (ptr *QModbusClient) QModbusClient_PTR() *QModbusClient {
	return ptr
}

func (ptr *QModbusClient) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QModbusDevice_PTR().Pointer()
	}
	return nil
}

func (ptr *QModbusClient) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QModbusDevice_PTR().SetPointer(p)
	}
}

func PointerFromQModbusClient(ptr QModbusClient_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QModbusClient_PTR().Pointer()
	}
	return nil
}

func (n *QModbusClient) InitFromInternal(ptr uintptr, name string) {
	n.QModbusDevice_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QModbusClient) ClassNameInternalF() string {
	return n.QModbusDevice_PTR().ClassNameInternalF()
}

func NewQModbusClientFromPointer(ptr unsafe.Pointer) (n *QModbusClient) {
	n = new(QModbusClient)
	n.InitFromInternal(uintptr(ptr), "serialbus.QModbusClient")
	return
}

func (ptr *QModbusClient) DestroyQModbusClient() {
}

func NewQModbusClient(parent core.QObject_ITF) *QModbusClient {

	return internal.CallLocalFunction([]interface{}{"", "", "serialbus.NewQModbusClient", "", parent}).(*QModbusClient)
}

func (ptr *QModbusClient) NumberOfRetries() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NumberOfRetries"}).(float64))
}

func (ptr *QModbusClient) ConnectProcessPrivateResponse(f func(response *QModbusResponse, data *QModbusDataUnit) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectProcessPrivateResponse", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QModbusClient) DisconnectProcessPrivateResponse() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectProcessPrivateResponse"})
}

func (ptr *QModbusClient) ProcessPrivateResponse(response QModbusResponse_ITF, data QModbusDataUnit_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ProcessPrivateResponse", response, data}).(bool)
}

func (ptr *QModbusClient) ProcessPrivateResponseDefault(response QModbusResponse_ITF, data QModbusDataUnit_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ProcessPrivateResponseDefault", response, data}).(bool)
}

func (ptr *QModbusClient) ConnectProcessResponse(f func(response *QModbusResponse, data *QModbusDataUnit) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectProcessResponse", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QModbusClient) DisconnectProcessResponse() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectProcessResponse"})
}

func (ptr *QModbusClient) ProcessResponse(response QModbusResponse_ITF, data QModbusDataUnit_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ProcessResponse", response, data}).(bool)
}

func (ptr *QModbusClient) ProcessResponseDefault(response QModbusResponse_ITF, data QModbusDataUnit_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ProcessResponseDefault", response, data}).(bool)
}

func (ptr *QModbusClient) SendRawRequest(request QModbusRequest_ITF, serverAddress int) *QModbusReply {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SendRawRequest", request, serverAddress}).(*QModbusReply)
}

func (ptr *QModbusClient) SendReadRequest(read QModbusDataUnit_ITF, serverAddress int) *QModbusReply {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SendReadRequest", read, serverAddress}).(*QModbusReply)
}

func (ptr *QModbusClient) SendReadWriteRequest(read QModbusDataUnit_ITF, write QModbusDataUnit_ITF, serverAddress int) *QModbusReply {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SendReadWriteRequest", read, write, serverAddress}).(*QModbusReply)
}

func (ptr *QModbusClient) SendWriteRequest(write QModbusDataUnit_ITF, serverAddress int) *QModbusReply {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SendWriteRequest", write, serverAddress}).(*QModbusReply)
}

func (ptr *QModbusClient) SetNumberOfRetries(number int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetNumberOfRetries", number})
}

func (ptr *QModbusClient) SetTimeout(newTimeout int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTimeout", newTimeout})
}

func (ptr *QModbusClient) Timeout() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Timeout"}).(float64))
}

func (ptr *QModbusClient) ConnectTimeoutChanged(f func(newTimeout int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTimeoutChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QModbusClient) DisconnectTimeoutChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTimeoutChanged"})
}

func (ptr *QModbusClient) TimeoutChanged(newTimeout int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimeoutChanged", newTimeout})
}

func (ptr *QModbusClient) Close() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Close"})
}

func (ptr *QModbusClient) CloseDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseDefault"})
}

func (ptr *QModbusClient) Open() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Open"}).(bool)
}

func (ptr *QModbusClient) OpenDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OpenDefault"}).(bool)
}

type QModbusDataUnit struct {
	internal.Internal
}

type QModbusDataUnit_ITF interface {
	QModbusDataUnit_PTR() *QModbusDataUnit
}

func (ptr *QModbusDataUnit) QModbusDataUnit_PTR() *QModbusDataUnit {
	return ptr
}

func (ptr *QModbusDataUnit) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QModbusDataUnit) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQModbusDataUnit(ptr QModbusDataUnit_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QModbusDataUnit_PTR().Pointer()
	}
	return nil
}

func (n *QModbusDataUnit) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQModbusDataUnitFromPointer(ptr unsafe.Pointer) (n *QModbusDataUnit) {
	n = new(QModbusDataUnit)
	n.InitFromInternal(uintptr(ptr), "serialbus.QModbusDataUnit")
	return
}

func (ptr *QModbusDataUnit) DestroyQModbusDataUnit() {
}

//go:generate stringer -type=QModbusDataUnit__RegisterType
//QModbusDataUnit::RegisterType
type QModbusDataUnit__RegisterType int64

const (
	QModbusDataUnit__Invalid          QModbusDataUnit__RegisterType = QModbusDataUnit__RegisterType(0)
	QModbusDataUnit__DiscreteInputs   QModbusDataUnit__RegisterType = QModbusDataUnit__RegisterType(1)
	QModbusDataUnit__Coils            QModbusDataUnit__RegisterType = QModbusDataUnit__RegisterType(2)
	QModbusDataUnit__InputRegisters   QModbusDataUnit__RegisterType = QModbusDataUnit__RegisterType(3)
	QModbusDataUnit__HoldingRegisters QModbusDataUnit__RegisterType = QModbusDataUnit__RegisterType(4)
)

func NewQModbusDataUnit() *QModbusDataUnit {

	return internal.CallLocalFunction([]interface{}{"", "", "serialbus.NewQModbusDataUnit", ""}).(*QModbusDataUnit)
}

func NewQModbusDataUnit2(ty QModbusDataUnit__RegisterType) *QModbusDataUnit {

	return internal.CallLocalFunction([]interface{}{"", "", "serialbus.NewQModbusDataUnit2", "", ty}).(*QModbusDataUnit)
}

func NewQModbusDataUnit3(ty QModbusDataUnit__RegisterType, address int, size uint16) *QModbusDataUnit {

	return internal.CallLocalFunction([]interface{}{"", "", "serialbus.NewQModbusDataUnit3", "", ty, address, size}).(*QModbusDataUnit)
}

func NewQModbusDataUnit4(ty QModbusDataUnit__RegisterType, address int, data []uint16) *QModbusDataUnit {

	return internal.CallLocalFunction([]interface{}{"", "", "serialbus.NewQModbusDataUnit4", "", ty, address, data}).(*QModbusDataUnit)
}

func (ptr *QModbusDataUnit) IsValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsValid"}).(bool)
}

func (ptr *QModbusDataUnit) RegisterType() QModbusDataUnit__RegisterType {

	return QModbusDataUnit__RegisterType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RegisterType"}).(float64))
}

func (ptr *QModbusDataUnit) SetRegisterType(ty QModbusDataUnit__RegisterType) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRegisterType", ty})
}

func (ptr *QModbusDataUnit) SetStartAddress(address int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetStartAddress", address})
}

func (ptr *QModbusDataUnit) SetValue(index int, value uint16) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetValue", index, value})
}

func (ptr *QModbusDataUnit) SetValueCount(newCount uint) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetValueCount", newCount})
}

func (ptr *QModbusDataUnit) SetValues(values []uint16) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetValues", values})
}

func (ptr *QModbusDataUnit) StartAddress() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StartAddress"}).(float64))
}

func (ptr *QModbusDataUnit) Value(index int) uint16 {

	return uint16(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Value", index}).(float64))
}

func (ptr *QModbusDataUnit) ValueCount() uint {

	return uint(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ValueCount"}).(float64))
}

func (ptr *QModbusDataUnit) Values() []uint16 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Values"}).([]uint16)
}

func (ptr *QModbusDataUnit) __QModbusDataUnit_data_atList4(i int) uint16 {

	return uint16(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__QModbusDataUnit_data_atList4", i}).(float64))
}

func (ptr *QModbusDataUnit) __QModbusDataUnit_data_setList4(i uint16) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__QModbusDataUnit_data_setList4", i})
}

func (ptr *QModbusDataUnit) __QModbusDataUnit_data_newList4() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__QModbusDataUnit_data_newList4"}).(unsafe.Pointer)
}

func (ptr *QModbusDataUnit) __setValues_values_atList(i int) uint16 {

	return uint16(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setValues_values_atList", i}).(float64))
}

func (ptr *QModbusDataUnit) __setValues_values_setList(i uint16) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setValues_values_setList", i})
}

func (ptr *QModbusDataUnit) __setValues_values_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setValues_values_newList"}).(unsafe.Pointer)
}

func (ptr *QModbusDataUnit) __values_atList(i int) uint16 {

	return uint16(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__values_atList", i}).(float64))
}

func (ptr *QModbusDataUnit) __values_setList(i uint16) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__values_setList", i})
}

func (ptr *QModbusDataUnit) __values_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__values_newList"}).(unsafe.Pointer)
}

func (ptr *QModbusDataUnit) __m_values_atList(i int) uint16 {

	return uint16(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__m_values_atList", i}).(float64))
}

func (ptr *QModbusDataUnit) __m_values_setList(i uint16) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__m_values_setList", i})
}

func (ptr *QModbusDataUnit) __m_values_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__m_values_newList"}).(unsafe.Pointer)
}

func (ptr *QModbusDataUnit) __setM_values__atList(i int) uint16 {

	return uint16(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setM_values__atList", i}).(float64))
}

func (ptr *QModbusDataUnit) __setM_values__setList(i uint16) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setM_values__setList", i})
}

func (ptr *QModbusDataUnit) __setM_values__newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setM_values__newList"}).(unsafe.Pointer)
}

type QModbusDevice struct {
	core.QObject
}

type QModbusDevice_ITF interface {
	core.QObject_ITF
	QModbusDevice_PTR() *QModbusDevice
}

func (ptr *QModbusDevice) QModbusDevice_PTR() *QModbusDevice {
	return ptr
}

func (ptr *QModbusDevice) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QModbusDevice) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQModbusDevice(ptr QModbusDevice_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QModbusDevice_PTR().Pointer()
	}
	return nil
}

func (n *QModbusDevice) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QModbusDevice) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQModbusDeviceFromPointer(ptr unsafe.Pointer) (n *QModbusDevice) {
	n = new(QModbusDevice)
	n.InitFromInternal(uintptr(ptr), "serialbus.QModbusDevice")
	return
}

//go:generate stringer -type=QModbusDevice__Error
//QModbusDevice::Error
type QModbusDevice__Error int64

const (
	QModbusDevice__NoError            QModbusDevice__Error = QModbusDevice__Error(0)
	QModbusDevice__ReadError          QModbusDevice__Error = QModbusDevice__Error(1)
	QModbusDevice__WriteError         QModbusDevice__Error = QModbusDevice__Error(2)
	QModbusDevice__ConnectionError    QModbusDevice__Error = QModbusDevice__Error(3)
	QModbusDevice__ConfigurationError QModbusDevice__Error = QModbusDevice__Error(4)
	QModbusDevice__TimeoutError       QModbusDevice__Error = QModbusDevice__Error(5)
	QModbusDevice__ProtocolError      QModbusDevice__Error = QModbusDevice__Error(6)
	QModbusDevice__ReplyAbortedError  QModbusDevice__Error = QModbusDevice__Error(7)
	QModbusDevice__UnknownError       QModbusDevice__Error = QModbusDevice__Error(8)
)

//go:generate stringer -type=QModbusDevice__State
//QModbusDevice::State
type QModbusDevice__State int64

const (
	QModbusDevice__UnconnectedState QModbusDevice__State = QModbusDevice__State(0)
	QModbusDevice__ConnectingState  QModbusDevice__State = QModbusDevice__State(1)
	QModbusDevice__ConnectedState   QModbusDevice__State = QModbusDevice__State(2)
	QModbusDevice__ClosingState     QModbusDevice__State = QModbusDevice__State(3)
)

//go:generate stringer -type=QModbusDevice__ConnectionParameter
//QModbusDevice::ConnectionParameter
type QModbusDevice__ConnectionParameter int64

const (
	QModbusDevice__SerialPortNameParameter QModbusDevice__ConnectionParameter = QModbusDevice__ConnectionParameter(0)
	QModbusDevice__SerialParityParameter   QModbusDevice__ConnectionParameter = QModbusDevice__ConnectionParameter(1)
	QModbusDevice__SerialBaudRateParameter QModbusDevice__ConnectionParameter = QModbusDevice__ConnectionParameter(2)
	QModbusDevice__SerialDataBitsParameter QModbusDevice__ConnectionParameter = QModbusDevice__ConnectionParameter(3)
	QModbusDevice__SerialStopBitsParameter QModbusDevice__ConnectionParameter = QModbusDevice__ConnectionParameter(4)
	QModbusDevice__NetworkPortParameter    QModbusDevice__ConnectionParameter = QModbusDevice__ConnectionParameter(5)
	QModbusDevice__NetworkAddressParameter QModbusDevice__ConnectionParameter = QModbusDevice__ConnectionParameter(6)
	QModbusDevice__UserParameter           QModbusDevice__ConnectionParameter = QModbusDevice__ConnectionParameter(0x100)
)

func NewQModbusDevice(parent core.QObject_ITF) *QModbusDevice {

	return internal.CallLocalFunction([]interface{}{"", "", "serialbus.NewQModbusDevice", "", parent}).(*QModbusDevice)
}

func (ptr *QModbusDevice) ConnectClose(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectClose", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QModbusDevice) DisconnectClose() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectClose"})
}

func (ptr *QModbusDevice) Close() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Close"})
}

func (ptr *QModbusDevice) ConnectDevice() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDevice"}).(bool)
}

func (ptr *QModbusDevice) ConnectionParameter(parameter int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectionParameter", parameter}).(*core.QVariant)
}

func (ptr *QModbusDevice) DisconnectDevice() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDevice"})
}

func (ptr *QModbusDevice) Error() QModbusDevice__Error {

	return QModbusDevice__Error(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error"}).(float64))
}

func (ptr *QModbusDevice) ConnectErrorOccurred(f func(error QModbusDevice__Error)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectErrorOccurred", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QModbusDevice) DisconnectErrorOccurred() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectErrorOccurred"})
}

func (ptr *QModbusDevice) ErrorOccurred(error QModbusDevice__Error) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ErrorOccurred", error})
}

func (ptr *QModbusDevice) ErrorString() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ErrorString"}).(string)
}

func (ptr *QModbusDevice) ConnectOpen(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectOpen", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QModbusDevice) DisconnectOpen() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectOpen"})
}

func (ptr *QModbusDevice) Open() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Open"}).(bool)
}

func (ptr *QModbusDevice) SetConnectionParameter(parameter int, value core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetConnectionParameter", parameter, value})
}

func (ptr *QModbusDevice) SetError(errorText string, error QModbusDevice__Error) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetError", errorText, error})
}

func (ptr *QModbusDevice) SetState(newState QModbusDevice__State) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetState", newState})
}

func (ptr *QModbusDevice) State() QModbusDevice__State {

	return QModbusDevice__State(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "State"}).(float64))
}

func (ptr *QModbusDevice) ConnectStateChanged(f func(state QModbusDevice__State)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStateChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QModbusDevice) DisconnectStateChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStateChanged"})
}

func (ptr *QModbusDevice) StateChanged(state QModbusDevice__State) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StateChanged", state})
}

func (ptr *QModbusDevice) ConnectDestroyQModbusDevice(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQModbusDevice", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QModbusDevice) DisconnectDestroyQModbusDevice() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQModbusDevice"})
}

func (ptr *QModbusDevice) DestroyQModbusDevice() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQModbusDevice"})
}

func (ptr *QModbusDevice) DestroyQModbusDeviceDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQModbusDeviceDefault"})
}

func (ptr *QModbusDevice) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QModbusDevice) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QModbusDevice) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QModbusDevice) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QModbusDevice) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QModbusDevice) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QModbusDevice) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QModbusDevice) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QModbusDevice) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QModbusDevice) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QModbusDevice) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QModbusDevice) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QModbusDevice) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QModbusDevice) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QModbusDevice) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QModbusDevice) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QModbusDevice) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QModbusDevice) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QModbusDevice) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QModbusDevice) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QModbusDevice) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QModbusDeviceIdentification struct {
	internal.Internal
}

type QModbusDeviceIdentification_ITF interface {
	QModbusDeviceIdentification_PTR() *QModbusDeviceIdentification
}

func (ptr *QModbusDeviceIdentification) QModbusDeviceIdentification_PTR() *QModbusDeviceIdentification {
	return ptr
}

func (ptr *QModbusDeviceIdentification) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QModbusDeviceIdentification) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQModbusDeviceIdentification(ptr QModbusDeviceIdentification_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QModbusDeviceIdentification_PTR().Pointer()
	}
	return nil
}

func (n *QModbusDeviceIdentification) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQModbusDeviceIdentificationFromPointer(ptr unsafe.Pointer) (n *QModbusDeviceIdentification) {
	n = new(QModbusDeviceIdentification)
	n.InitFromInternal(uintptr(ptr), "serialbus.QModbusDeviceIdentification")
	return
}

func (ptr *QModbusDeviceIdentification) DestroyQModbusDeviceIdentification() {
}

//go:generate stringer -type=QModbusDeviceIdentification__ObjectId
//QModbusDeviceIdentification::ObjectId
type QModbusDeviceIdentification__ObjectId int64

const (
	QModbusDeviceIdentification__VendorNameObjectId          QModbusDeviceIdentification__ObjectId = QModbusDeviceIdentification__ObjectId(0x00)
	QModbusDeviceIdentification__ProductCodeObjectId         QModbusDeviceIdentification__ObjectId = QModbusDeviceIdentification__ObjectId(0x01)
	QModbusDeviceIdentification__MajorMinorRevisionObjectId  QModbusDeviceIdentification__ObjectId = QModbusDeviceIdentification__ObjectId(0x02)
	QModbusDeviceIdentification__VendorUrlObjectId           QModbusDeviceIdentification__ObjectId = QModbusDeviceIdentification__ObjectId(0x03)
	QModbusDeviceIdentification__ProductNameObjectId         QModbusDeviceIdentification__ObjectId = QModbusDeviceIdentification__ObjectId(0x04)
	QModbusDeviceIdentification__ModelNameObjectId           QModbusDeviceIdentification__ObjectId = QModbusDeviceIdentification__ObjectId(0x05)
	QModbusDeviceIdentification__UserApplicationNameObjectId QModbusDeviceIdentification__ObjectId = QModbusDeviceIdentification__ObjectId(0x06)
	QModbusDeviceIdentification__ReservedObjectId            QModbusDeviceIdentification__ObjectId = QModbusDeviceIdentification__ObjectId(0x07)
	QModbusDeviceIdentification__ProductDependentObjectId    QModbusDeviceIdentification__ObjectId = QModbusDeviceIdentification__ObjectId(0x80)
	QModbusDeviceIdentification__UndefinedObjectId           QModbusDeviceIdentification__ObjectId = QModbusDeviceIdentification__ObjectId(0x100)
)

//go:generate stringer -type=QModbusDeviceIdentification__ReadDeviceIdCode
//QModbusDeviceIdentification::ReadDeviceIdCode
type QModbusDeviceIdentification__ReadDeviceIdCode int64

const (
	QModbusDeviceIdentification__BasicReadDeviceIdCode      QModbusDeviceIdentification__ReadDeviceIdCode = QModbusDeviceIdentification__ReadDeviceIdCode(0x01)
	QModbusDeviceIdentification__RegularReadDeviceIdCode    QModbusDeviceIdentification__ReadDeviceIdCode = QModbusDeviceIdentification__ReadDeviceIdCode(0x02)
	QModbusDeviceIdentification__ExtendedReadDeviceIdCode   QModbusDeviceIdentification__ReadDeviceIdCode = QModbusDeviceIdentification__ReadDeviceIdCode(0x03)
	QModbusDeviceIdentification__IndividualReadDeviceIdCode QModbusDeviceIdentification__ReadDeviceIdCode = QModbusDeviceIdentification__ReadDeviceIdCode(0x04)
)

//go:generate stringer -type=QModbusDeviceIdentification__ConformityLevel
//QModbusDeviceIdentification::ConformityLevel
type QModbusDeviceIdentification__ConformityLevel int64

const (
	QModbusDeviceIdentification__BasicConformityLevel              QModbusDeviceIdentification__ConformityLevel = QModbusDeviceIdentification__ConformityLevel(0x01)
	QModbusDeviceIdentification__RegularConformityLevel            QModbusDeviceIdentification__ConformityLevel = QModbusDeviceIdentification__ConformityLevel(0x02)
	QModbusDeviceIdentification__ExtendedConformityLevel           QModbusDeviceIdentification__ConformityLevel = QModbusDeviceIdentification__ConformityLevel(0x03)
	QModbusDeviceIdentification__BasicIndividualConformityLevel    QModbusDeviceIdentification__ConformityLevel = QModbusDeviceIdentification__ConformityLevel(0x81)
	QModbusDeviceIdentification__RegularIndividualConformityLevel  QModbusDeviceIdentification__ConformityLevel = QModbusDeviceIdentification__ConformityLevel(0x82)
	QModbusDeviceIdentification__ExtendedIndividualConformityLevel QModbusDeviceIdentification__ConformityLevel = QModbusDeviceIdentification__ConformityLevel(0x83)
)

func NewQModbusDeviceIdentification() *QModbusDeviceIdentification {

	return internal.CallLocalFunction([]interface{}{"", "", "serialbus.NewQModbusDeviceIdentification", ""}).(*QModbusDeviceIdentification)
}

func (ptr *QModbusDeviceIdentification) ConformityLevel() QModbusDeviceIdentification__ConformityLevel {

	return QModbusDeviceIdentification__ConformityLevel(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConformityLevel"}).(float64))
}

func (ptr *QModbusDeviceIdentification) Contains(objectId uint) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Contains", objectId}).(bool)
}

func QModbusDeviceIdentification_FromByteArray(ba core.QByteArray_ITF) *QModbusDeviceIdentification {

	return internal.CallLocalFunction([]interface{}{"", "", "serialbus.QModbusDeviceIdentification_FromByteArray", "", ba}).(*QModbusDeviceIdentification)
}

func (ptr *QModbusDeviceIdentification) FromByteArray(ba core.QByteArray_ITF) *QModbusDeviceIdentification {

	return internal.CallLocalFunction([]interface{}{"", "", "serialbus.QModbusDeviceIdentification_FromByteArray", "", ba}).(*QModbusDeviceIdentification)
}

func (ptr *QModbusDeviceIdentification) Insert(objectId uint, value core.QByteArray_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Insert", objectId, value}).(bool)
}

func (ptr *QModbusDeviceIdentification) IsValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsValid"}).(bool)
}

func (ptr *QModbusDeviceIdentification) ObjectIds() []int {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ObjectIds"}).([]int)
}

func (ptr *QModbusDeviceIdentification) Remove(objectId uint) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Remove", objectId})
}

func (ptr *QModbusDeviceIdentification) SetConformityLevel(level QModbusDeviceIdentification__ConformityLevel) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetConformityLevel", level})
}

func (ptr *QModbusDeviceIdentification) Value(objectId uint) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Value", objectId}).(*core.QByteArray)
}

func (ptr *QModbusDeviceIdentification) __objectIds_atList(i int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__objectIds_atList", i}).(float64))
}

func (ptr *QModbusDeviceIdentification) __objectIds_setList(i int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__objectIds_setList", i})
}

func (ptr *QModbusDeviceIdentification) __objectIds_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__objectIds_newList"}).(unsafe.Pointer)
}

func (ptr *QModbusDeviceIdentification) __m_objects_atList(v int, i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__m_objects_atList", v, i}).(*core.QByteArray)
}

func (ptr *QModbusDeviceIdentification) __m_objects_setList(key int, i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__m_objects_setList", key, i})
}

func (ptr *QModbusDeviceIdentification) __m_objects_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__m_objects_newList"}).(unsafe.Pointer)
}

func (ptr *QModbusDeviceIdentification) __m_objects_keyList() []int {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__m_objects_keyList"}).([]int)
}

func (ptr *QModbusDeviceIdentification) __setM_objects__atList(v int, i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setM_objects__atList", v, i}).(*core.QByteArray)
}

func (ptr *QModbusDeviceIdentification) __setM_objects__setList(key int, i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setM_objects__setList", key, i})
}

func (ptr *QModbusDeviceIdentification) __setM_objects__newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setM_objects__newList"}).(unsafe.Pointer)
}

func (ptr *QModbusDeviceIdentification) __setM_objects__keyList() []int {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setM_objects__keyList"}).([]int)
}

func (ptr *QModbusDeviceIdentification) ____m_objects_keyList_atList(i int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____m_objects_keyList_atList", i}).(float64))
}

func (ptr *QModbusDeviceIdentification) ____m_objects_keyList_setList(i int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____m_objects_keyList_setList", i})
}

func (ptr *QModbusDeviceIdentification) ____m_objects_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____m_objects_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QModbusDeviceIdentification) ____setM_objects__keyList_atList(i int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setM_objects__keyList_atList", i}).(float64))
}

func (ptr *QModbusDeviceIdentification) ____setM_objects__keyList_setList(i int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setM_objects__keyList_setList", i})
}

func (ptr *QModbusDeviceIdentification) ____setM_objects__keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setM_objects__keyList_newList"}).(unsafe.Pointer)
}

type QModbusExceptionResponse struct {
	QModbusResponse
}

type QModbusExceptionResponse_ITF interface {
	QModbusResponse_ITF
	QModbusExceptionResponse_PTR() *QModbusExceptionResponse
}

func (ptr *QModbusExceptionResponse) QModbusExceptionResponse_PTR() *QModbusExceptionResponse {
	return ptr
}

func (ptr *QModbusExceptionResponse) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QModbusResponse_PTR().Pointer()
	}
	return nil
}

func (ptr *QModbusExceptionResponse) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QModbusResponse_PTR().SetPointer(p)
	}
}

func PointerFromQModbusExceptionResponse(ptr QModbusExceptionResponse_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QModbusExceptionResponse_PTR().Pointer()
	}
	return nil
}

func (n *QModbusExceptionResponse) InitFromInternal(ptr uintptr, name string) {
	n.QModbusResponse_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QModbusExceptionResponse) ClassNameInternalF() string {
	return n.QModbusResponse_PTR().ClassNameInternalF()
}

func NewQModbusExceptionResponseFromPointer(ptr unsafe.Pointer) (n *QModbusExceptionResponse) {
	n = new(QModbusExceptionResponse)
	n.InitFromInternal(uintptr(ptr), "serialbus.QModbusExceptionResponse")
	return
}

func (ptr *QModbusExceptionResponse) DestroyQModbusExceptionResponse() {
}

func NewQModbusExceptionResponse() *QModbusExceptionResponse {

	return internal.CallLocalFunction([]interface{}{"", "", "serialbus.NewQModbusExceptionResponse", ""}).(*QModbusExceptionResponse)
}

func NewQModbusExceptionResponse2(pdu QModbusPdu_ITF) *QModbusExceptionResponse {

	return internal.CallLocalFunction([]interface{}{"", "", "serialbus.NewQModbusExceptionResponse2", "", pdu}).(*QModbusExceptionResponse)
}

func NewQModbusExceptionResponse3(code QModbusPdu__FunctionCode, ec QModbusPdu__ExceptionCode) *QModbusExceptionResponse {

	return internal.CallLocalFunction([]interface{}{"", "", "serialbus.NewQModbusExceptionResponse3", "", code, ec}).(*QModbusExceptionResponse)
}

func (ptr *QModbusExceptionResponse) SetExceptionCode(ec QModbusPdu__ExceptionCode) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetExceptionCode", ec})
}

type QModbusPdu struct {
	internal.Internal
}

type QModbusPdu_ITF interface {
	QModbusPdu_PTR() *QModbusPdu
}

func (ptr *QModbusPdu) QModbusPdu_PTR() *QModbusPdu {
	return ptr
}

func (ptr *QModbusPdu) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QModbusPdu) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQModbusPdu(ptr QModbusPdu_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QModbusPdu_PTR().Pointer()
	}
	return nil
}

func (n *QModbusPdu) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQModbusPduFromPointer(ptr unsafe.Pointer) (n *QModbusPdu) {
	n = new(QModbusPdu)
	n.InitFromInternal(uintptr(ptr), "serialbus.QModbusPdu")
	return
}

//go:generate stringer -type=QModbusPdu__ExceptionCode
//QModbusPdu::ExceptionCode
type QModbusPdu__ExceptionCode int64

const (
	QModbusPdu__IllegalFunction                    QModbusPdu__ExceptionCode = QModbusPdu__ExceptionCode(0x01)
	QModbusPdu__IllegalDataAddress                 QModbusPdu__ExceptionCode = QModbusPdu__ExceptionCode(0x02)
	QModbusPdu__IllegalDataValue                   QModbusPdu__ExceptionCode = QModbusPdu__ExceptionCode(0x03)
	QModbusPdu__ServerDeviceFailure                QModbusPdu__ExceptionCode = QModbusPdu__ExceptionCode(0x04)
	QModbusPdu__Acknowledge                        QModbusPdu__ExceptionCode = QModbusPdu__ExceptionCode(0x05)
	QModbusPdu__ServerDeviceBusy                   QModbusPdu__ExceptionCode = QModbusPdu__ExceptionCode(0x06)
	QModbusPdu__NegativeAcknowledge                QModbusPdu__ExceptionCode = QModbusPdu__ExceptionCode(0x07)
	QModbusPdu__MemoryParityError                  QModbusPdu__ExceptionCode = QModbusPdu__ExceptionCode(0x08)
	QModbusPdu__GatewayPathUnavailable             QModbusPdu__ExceptionCode = QModbusPdu__ExceptionCode(0x0A)
	QModbusPdu__GatewayTargetDeviceFailedToRespond QModbusPdu__ExceptionCode = QModbusPdu__ExceptionCode(0x0B)
	QModbusPdu__ExtendedException                  QModbusPdu__ExceptionCode = QModbusPdu__ExceptionCode(0xFF)
)

//go:generate stringer -type=QModbusPdu__FunctionCode
//QModbusPdu::FunctionCode
type QModbusPdu__FunctionCode int64

const (
	QModbusPdu__Invalid                        QModbusPdu__FunctionCode = QModbusPdu__FunctionCode(0x00)
	QModbusPdu__ReadCoils                      QModbusPdu__FunctionCode = QModbusPdu__FunctionCode(0x01)
	QModbusPdu__ReadDiscreteInputs             QModbusPdu__FunctionCode = QModbusPdu__FunctionCode(0x02)
	QModbusPdu__ReadHoldingRegisters           QModbusPdu__FunctionCode = QModbusPdu__FunctionCode(0x03)
	QModbusPdu__ReadInputRegisters             QModbusPdu__FunctionCode = QModbusPdu__FunctionCode(0x04)
	QModbusPdu__WriteSingleCoil                QModbusPdu__FunctionCode = QModbusPdu__FunctionCode(0x05)
	QModbusPdu__WriteSingleRegister            QModbusPdu__FunctionCode = QModbusPdu__FunctionCode(0x06)
	QModbusPdu__ReadExceptionStatus            QModbusPdu__FunctionCode = QModbusPdu__FunctionCode(0x07)
	QModbusPdu__Diagnostics                    QModbusPdu__FunctionCode = QModbusPdu__FunctionCode(0x08)
	QModbusPdu__GetCommEventCounter            QModbusPdu__FunctionCode = QModbusPdu__FunctionCode(0x0B)
	QModbusPdu__GetCommEventLog                QModbusPdu__FunctionCode = QModbusPdu__FunctionCode(0x0C)
	QModbusPdu__WriteMultipleCoils             QModbusPdu__FunctionCode = QModbusPdu__FunctionCode(0x0F)
	QModbusPdu__WriteMultipleRegisters         QModbusPdu__FunctionCode = QModbusPdu__FunctionCode(0x10)
	QModbusPdu__ReportServerId                 QModbusPdu__FunctionCode = QModbusPdu__FunctionCode(0x11)
	QModbusPdu__ReadFileRecord                 QModbusPdu__FunctionCode = QModbusPdu__FunctionCode(0x14)
	QModbusPdu__WriteFileRecord                QModbusPdu__FunctionCode = QModbusPdu__FunctionCode(0x15)
	QModbusPdu__MaskWriteRegister              QModbusPdu__FunctionCode = QModbusPdu__FunctionCode(0x16)
	QModbusPdu__ReadWriteMultipleRegisters     QModbusPdu__FunctionCode = QModbusPdu__FunctionCode(0x17)
	QModbusPdu__ReadFifoQueue                  QModbusPdu__FunctionCode = QModbusPdu__FunctionCode(0x18)
	QModbusPdu__EncapsulatedInterfaceTransport QModbusPdu__FunctionCode = QModbusPdu__FunctionCode(0x2B)
	QModbusPdu__UndefinedFunctionCode          QModbusPdu__FunctionCode = QModbusPdu__FunctionCode(0x100)
)

func NewQModbusPdu() *QModbusPdu {

	return internal.CallLocalFunction([]interface{}{"", "", "serialbus.NewQModbusPdu", ""}).(*QModbusPdu)
}

func NewQModbusPdu2(code QModbusPdu__FunctionCode, data core.QByteArray_ITF) *QModbusPdu {

	return internal.CallLocalFunction([]interface{}{"", "", "serialbus.NewQModbusPdu2", "", code, data}).(*QModbusPdu)
}

func NewQModbusPdu3(other QModbusPdu_ITF) *QModbusPdu {

	return internal.CallLocalFunction([]interface{}{"", "", "serialbus.NewQModbusPdu3", "", other}).(*QModbusPdu)
}

func (ptr *QModbusPdu) Data() *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Data"}).(*core.QByteArray)
}

func (ptr *QModbusPdu) DataSize() int16 {

	return int16(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DataSize"}).(float64))
}

func (ptr *QModbusPdu) ExceptionCode() QModbusPdu__ExceptionCode {

	return QModbusPdu__ExceptionCode(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ExceptionCode"}).(float64))
}

func (ptr *QModbusPdu) FunctionCode() QModbusPdu__FunctionCode {

	return QModbusPdu__FunctionCode(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FunctionCode"}).(float64))
}

func (ptr *QModbusPdu) IsException() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsException"}).(bool)
}

func (ptr *QModbusPdu) IsValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsValid"}).(bool)
}

func (ptr *QModbusPdu) SetData(data core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetData", data})
}

func (ptr *QModbusPdu) ConnectSetFunctionCode(f func(code QModbusPdu__FunctionCode)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetFunctionCode", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QModbusPdu) DisconnectSetFunctionCode() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetFunctionCode"})
}

func (ptr *QModbusPdu) SetFunctionCode(code QModbusPdu__FunctionCode) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFunctionCode", code})
}

func (ptr *QModbusPdu) SetFunctionCodeDefault(code QModbusPdu__FunctionCode) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFunctionCodeDefault", code})
}

func (ptr *QModbusPdu) Size() int16 {

	return int16(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Size"}).(float64))
}

func (ptr *QModbusPdu) ConnectDestroyQModbusPdu(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQModbusPdu", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QModbusPdu) DisconnectDestroyQModbusPdu() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQModbusPdu"})
}

func (ptr *QModbusPdu) DestroyQModbusPdu() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQModbusPdu"})
}

func (ptr *QModbusPdu) DestroyQModbusPduDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQModbusPduDefault"})
}

func QModbusPdu_ExceptionByte() string {

	return internal.CallLocalFunction([]interface{}{"", "", "serialbus.QModbusPdu_ExceptionByte", ""}).(string)
}

func (ptr *QModbusPdu) ExceptionByte() string {

	return internal.CallLocalFunction([]interface{}{"", "", "serialbus.QModbusPdu_ExceptionByte", ""}).(string)
}

func (ptr *QModbusPdu) __encode_vector_atList2(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__encode_vector_atList2", i}).(*core.QObject)
}

func (ptr *QModbusPdu) __encode_vector_setList2(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__encode_vector_setList2", i})
}

func (ptr *QModbusPdu) __encode_vector_newList2() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__encode_vector_newList2"}).(unsafe.Pointer)
}

type QModbusReply struct {
	core.QObject
}

type QModbusReply_ITF interface {
	core.QObject_ITF
	QModbusReply_PTR() *QModbusReply
}

func (ptr *QModbusReply) QModbusReply_PTR() *QModbusReply {
	return ptr
}

func (ptr *QModbusReply) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QModbusReply) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQModbusReply(ptr QModbusReply_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QModbusReply_PTR().Pointer()
	}
	return nil
}

func (n *QModbusReply) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QModbusReply) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQModbusReplyFromPointer(ptr unsafe.Pointer) (n *QModbusReply) {
	n = new(QModbusReply)
	n.InitFromInternal(uintptr(ptr), "serialbus.QModbusReply")
	return
}

func (ptr *QModbusReply) DestroyQModbusReply() {
}

//go:generate stringer -type=QModbusReply__ReplyType
//QModbusReply::ReplyType
type QModbusReply__ReplyType int64

const (
	QModbusReply__Raw       QModbusReply__ReplyType = QModbusReply__ReplyType(0)
	QModbusReply__Common    QModbusReply__ReplyType = QModbusReply__ReplyType(1)
	QModbusReply__Broadcast QModbusReply__ReplyType = QModbusReply__ReplyType(2)
)

func NewQModbusReply(ty QModbusReply__ReplyType, serverAddress int, parent core.QObject_ITF) *QModbusReply {

	return internal.CallLocalFunction([]interface{}{"", "", "serialbus.NewQModbusReply", "", ty, serverAddress, parent}).(*QModbusReply)
}

func (ptr *QModbusReply) Error() QModbusDevice__Error {

	return QModbusDevice__Error(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error"}).(float64))
}

func (ptr *QModbusReply) ConnectErrorOccurred(f func(error QModbusDevice__Error)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectErrorOccurred", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QModbusReply) DisconnectErrorOccurred() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectErrorOccurred"})
}

func (ptr *QModbusReply) ErrorOccurred(error QModbusDevice__Error) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ErrorOccurred", error})
}

func (ptr *QModbusReply) ErrorString() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ErrorString"}).(string)
}

func (ptr *QModbusReply) ConnectFinished(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFinished", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QModbusReply) DisconnectFinished() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFinished"})
}

func (ptr *QModbusReply) Finished() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Finished"})
}

func (ptr *QModbusReply) IsFinished() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsFinished"}).(bool)
}

func (ptr *QModbusReply) RawResult() *QModbusResponse {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RawResult"}).(*QModbusResponse)
}

func (ptr *QModbusReply) Result() *QModbusDataUnit {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Result"}).(*QModbusDataUnit)
}

func (ptr *QModbusReply) ServerAddress() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ServerAddress"}).(float64))
}

func (ptr *QModbusReply) Type() QModbusReply__ReplyType {

	return QModbusReply__ReplyType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(float64))
}

func (ptr *QModbusReply) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QModbusReply) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QModbusReply) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QModbusReply) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QModbusReply) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QModbusReply) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QModbusReply) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QModbusReply) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QModbusReply) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QModbusReply) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QModbusReply) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QModbusReply) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QModbusReply) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QModbusReply) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QModbusReply) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QModbusReply) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QModbusReply) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QModbusReply) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QModbusReply) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QModbusReply) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QModbusReply) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QModbusRequest struct {
	QModbusPdu
}

type QModbusRequest_ITF interface {
	QModbusPdu_ITF
	QModbusRequest_PTR() *QModbusRequest
}

func (ptr *QModbusRequest) QModbusRequest_PTR() *QModbusRequest {
	return ptr
}

func (ptr *QModbusRequest) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QModbusPdu_PTR().Pointer()
	}
	return nil
}

func (ptr *QModbusRequest) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QModbusPdu_PTR().SetPointer(p)
	}
}

func PointerFromQModbusRequest(ptr QModbusRequest_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QModbusRequest_PTR().Pointer()
	}
	return nil
}

func (n *QModbusRequest) InitFromInternal(ptr uintptr, name string) {
	n.QModbusPdu_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QModbusRequest) ClassNameInternalF() string {
	return n.QModbusPdu_PTR().ClassNameInternalF()
}

func NewQModbusRequestFromPointer(ptr unsafe.Pointer) (n *QModbusRequest) {
	n = new(QModbusRequest)
	n.InitFromInternal(uintptr(ptr), "serialbus.QModbusRequest")
	return
}

func (ptr *QModbusRequest) DestroyQModbusRequest() {
}

func NewQModbusRequest() *QModbusRequest {

	return internal.CallLocalFunction([]interface{}{"", "", "serialbus.NewQModbusRequest", ""}).(*QModbusRequest)
}

func NewQModbusRequest2(pdu QModbusPdu_ITF) *QModbusRequest {

	return internal.CallLocalFunction([]interface{}{"", "", "serialbus.NewQModbusRequest2", "", pdu}).(*QModbusRequest)
}

func NewQModbusRequest3(code QModbusPdu__FunctionCode, data core.QByteArray_ITF) *QModbusRequest {

	return internal.CallLocalFunction([]interface{}{"", "", "serialbus.NewQModbusRequest3", "", code, data}).(*QModbusRequest)
}

func QModbusRequest_CalculateDataSize(request QModbusRequest_ITF) int {

	return int(internal.CallLocalFunction([]interface{}{"", "", "serialbus.QModbusRequest_CalculateDataSize", "", request}).(float64))
}

func (ptr *QModbusRequest) CalculateDataSize(request QModbusRequest_ITF) int {

	return int(internal.CallLocalFunction([]interface{}{"", "", "serialbus.QModbusRequest_CalculateDataSize", "", request}).(float64))
}

func QModbusRequest_MinimumDataSize(request QModbusRequest_ITF) int {

	return int(internal.CallLocalFunction([]interface{}{"", "", "serialbus.QModbusRequest_MinimumDataSize", "", request}).(float64))
}

func (ptr *QModbusRequest) MinimumDataSize(request QModbusRequest_ITF) int {

	return int(internal.CallLocalFunction([]interface{}{"", "", "serialbus.QModbusRequest_MinimumDataSize", "", request}).(float64))
}

type QModbusResponse struct {
	QModbusPdu
}

type QModbusResponse_ITF interface {
	QModbusPdu_ITF
	QModbusResponse_PTR() *QModbusResponse
}

func (ptr *QModbusResponse) QModbusResponse_PTR() *QModbusResponse {
	return ptr
}

func (ptr *QModbusResponse) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QModbusPdu_PTR().Pointer()
	}
	return nil
}

func (ptr *QModbusResponse) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QModbusPdu_PTR().SetPointer(p)
	}
}

func PointerFromQModbusResponse(ptr QModbusResponse_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QModbusResponse_PTR().Pointer()
	}
	return nil
}

func (n *QModbusResponse) InitFromInternal(ptr uintptr, name string) {
	n.QModbusPdu_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QModbusResponse) ClassNameInternalF() string {
	return n.QModbusPdu_PTR().ClassNameInternalF()
}

func NewQModbusResponseFromPointer(ptr unsafe.Pointer) (n *QModbusResponse) {
	n = new(QModbusResponse)
	n.InitFromInternal(uintptr(ptr), "serialbus.QModbusResponse")
	return
}

func (ptr *QModbusResponse) DestroyQModbusResponse() {
}

func NewQModbusResponse() *QModbusResponse {

	return internal.CallLocalFunction([]interface{}{"", "", "serialbus.NewQModbusResponse", ""}).(*QModbusResponse)
}

func NewQModbusResponse2(pdu QModbusPdu_ITF) *QModbusResponse {

	return internal.CallLocalFunction([]interface{}{"", "", "serialbus.NewQModbusResponse2", "", pdu}).(*QModbusResponse)
}

func NewQModbusResponse3(code QModbusPdu__FunctionCode, data core.QByteArray_ITF) *QModbusResponse {

	return internal.CallLocalFunction([]interface{}{"", "", "serialbus.NewQModbusResponse3", "", code, data}).(*QModbusResponse)
}

func QModbusResponse_CalculateDataSize(response QModbusResponse_ITF) int {

	return int(internal.CallLocalFunction([]interface{}{"", "", "serialbus.QModbusResponse_CalculateDataSize", "", response}).(float64))
}

func (ptr *QModbusResponse) CalculateDataSize(response QModbusResponse_ITF) int {

	return int(internal.CallLocalFunction([]interface{}{"", "", "serialbus.QModbusResponse_CalculateDataSize", "", response}).(float64))
}

func QModbusResponse_MinimumDataSize(response QModbusResponse_ITF) int {

	return int(internal.CallLocalFunction([]interface{}{"", "", "serialbus.QModbusResponse_MinimumDataSize", "", response}).(float64))
}

func (ptr *QModbusResponse) MinimumDataSize(response QModbusResponse_ITF) int {

	return int(internal.CallLocalFunction([]interface{}{"", "", "serialbus.QModbusResponse_MinimumDataSize", "", response}).(float64))
}

type QModbusRtuSerialMaster struct {
	QModbusClient
}

type QModbusRtuSerialMaster_ITF interface {
	QModbusClient_ITF
	QModbusRtuSerialMaster_PTR() *QModbusRtuSerialMaster
}

func (ptr *QModbusRtuSerialMaster) QModbusRtuSerialMaster_PTR() *QModbusRtuSerialMaster {
	return ptr
}

func (ptr *QModbusRtuSerialMaster) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QModbusClient_PTR().Pointer()
	}
	return nil
}

func (ptr *QModbusRtuSerialMaster) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QModbusClient_PTR().SetPointer(p)
	}
}

func PointerFromQModbusRtuSerialMaster(ptr QModbusRtuSerialMaster_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QModbusRtuSerialMaster_PTR().Pointer()
	}
	return nil
}

func (n *QModbusRtuSerialMaster) InitFromInternal(ptr uintptr, name string) {
	n.QModbusClient_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QModbusRtuSerialMaster) ClassNameInternalF() string {
	return n.QModbusClient_PTR().ClassNameInternalF()
}

func NewQModbusRtuSerialMasterFromPointer(ptr unsafe.Pointer) (n *QModbusRtuSerialMaster) {
	n = new(QModbusRtuSerialMaster)
	n.InitFromInternal(uintptr(ptr), "serialbus.QModbusRtuSerialMaster")
	return
}

func (ptr *QModbusRtuSerialMaster) DestroyQModbusRtuSerialMaster() {
}

func NewQModbusRtuSerialMaster(parent core.QObject_ITF) *QModbusRtuSerialMaster {

	return internal.CallLocalFunction([]interface{}{"", "", "serialbus.NewQModbusRtuSerialMaster", "", parent}).(*QModbusRtuSerialMaster)
}

func (ptr *QModbusRtuSerialMaster) ConnectClose(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectClose", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QModbusRtuSerialMaster) DisconnectClose() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectClose"})
}

func (ptr *QModbusRtuSerialMaster) Close() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Close"})
}

func (ptr *QModbusRtuSerialMaster) CloseDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseDefault"})
}

func (ptr *QModbusRtuSerialMaster) InterFrameDelay() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InterFrameDelay"}).(float64))
}

func (ptr *QModbusRtuSerialMaster) ConnectOpen(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectOpen", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QModbusRtuSerialMaster) DisconnectOpen() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectOpen"})
}

func (ptr *QModbusRtuSerialMaster) Open() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Open"}).(bool)
}

func (ptr *QModbusRtuSerialMaster) OpenDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OpenDefault"}).(bool)
}

func (ptr *QModbusRtuSerialMaster) SetInterFrameDelay(microseconds int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetInterFrameDelay", microseconds})
}

func (ptr *QModbusRtuSerialMaster) SetTurnaroundDelay(turnaroundDelay int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTurnaroundDelay", turnaroundDelay})
}

func (ptr *QModbusRtuSerialMaster) TurnaroundDelay() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TurnaroundDelay"}).(float64))
}

type QModbusRtuSerialSlave struct {
	QModbusServer
}

type QModbusRtuSerialSlave_ITF interface {
	QModbusServer_ITF
	QModbusRtuSerialSlave_PTR() *QModbusRtuSerialSlave
}

func (ptr *QModbusRtuSerialSlave) QModbusRtuSerialSlave_PTR() *QModbusRtuSerialSlave {
	return ptr
}

func (ptr *QModbusRtuSerialSlave) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QModbusServer_PTR().Pointer()
	}
	return nil
}

func (ptr *QModbusRtuSerialSlave) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QModbusServer_PTR().SetPointer(p)
	}
}

func PointerFromQModbusRtuSerialSlave(ptr QModbusRtuSerialSlave_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QModbusRtuSerialSlave_PTR().Pointer()
	}
	return nil
}

func (n *QModbusRtuSerialSlave) InitFromInternal(ptr uintptr, name string) {
	n.QModbusServer_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QModbusRtuSerialSlave) ClassNameInternalF() string {
	return n.QModbusServer_PTR().ClassNameInternalF()
}

func NewQModbusRtuSerialSlaveFromPointer(ptr unsafe.Pointer) (n *QModbusRtuSerialSlave) {
	n = new(QModbusRtuSerialSlave)
	n.InitFromInternal(uintptr(ptr), "serialbus.QModbusRtuSerialSlave")
	return
}
func NewQModbusRtuSerialSlave(parent core.QObject_ITF) *QModbusRtuSerialSlave {

	return internal.CallLocalFunction([]interface{}{"", "", "serialbus.NewQModbusRtuSerialSlave", "", parent}).(*QModbusRtuSerialSlave)
}

func (ptr *QModbusRtuSerialSlave) ConnectClose(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectClose", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QModbusRtuSerialSlave) DisconnectClose() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectClose"})
}

func (ptr *QModbusRtuSerialSlave) Close() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Close"})
}

func (ptr *QModbusRtuSerialSlave) CloseDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseDefault"})
}

func (ptr *QModbusRtuSerialSlave) ConnectOpen(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectOpen", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QModbusRtuSerialSlave) DisconnectOpen() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectOpen"})
}

func (ptr *QModbusRtuSerialSlave) Open() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Open"}).(bool)
}

func (ptr *QModbusRtuSerialSlave) OpenDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OpenDefault"}).(bool)
}

func (ptr *QModbusRtuSerialSlave) ConnectDestroyQModbusRtuSerialSlave(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQModbusRtuSerialSlave", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QModbusRtuSerialSlave) DisconnectDestroyQModbusRtuSerialSlave() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQModbusRtuSerialSlave"})
}

func (ptr *QModbusRtuSerialSlave) DestroyQModbusRtuSerialSlave() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQModbusRtuSerialSlave"})
}

func (ptr *QModbusRtuSerialSlave) DestroyQModbusRtuSerialSlaveDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQModbusRtuSerialSlaveDefault"})
}

type QModbusServer struct {
	QModbusDevice
}

type QModbusServer_ITF interface {
	QModbusDevice_ITF
	QModbusServer_PTR() *QModbusServer
}

func (ptr *QModbusServer) QModbusServer_PTR() *QModbusServer {
	return ptr
}

func (ptr *QModbusServer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QModbusDevice_PTR().Pointer()
	}
	return nil
}

func (ptr *QModbusServer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QModbusDevice_PTR().SetPointer(p)
	}
}

func PointerFromQModbusServer(ptr QModbusServer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QModbusServer_PTR().Pointer()
	}
	return nil
}

func (n *QModbusServer) InitFromInternal(ptr uintptr, name string) {
	n.QModbusDevice_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QModbusServer) ClassNameInternalF() string {
	return n.QModbusDevice_PTR().ClassNameInternalF()
}

func NewQModbusServerFromPointer(ptr unsafe.Pointer) (n *QModbusServer) {
	n = new(QModbusServer)
	n.InitFromInternal(uintptr(ptr), "serialbus.QModbusServer")
	return
}

func (ptr *QModbusServer) DestroyQModbusServer() {
}

//go:generate stringer -type=QModbusServer__Option
//QModbusServer::Option
type QModbusServer__Option int64

const (
	QModbusServer__DiagnosticRegister    QModbusServer__Option = QModbusServer__Option(0)
	QModbusServer__ExceptionStatusOffset QModbusServer__Option = QModbusServer__Option(1)
	QModbusServer__DeviceBusy            QModbusServer__Option = QModbusServer__Option(2)
	QModbusServer__AsciiInputDelimiter   QModbusServer__Option = QModbusServer__Option(3)
	QModbusServer__ListenOnlyMode        QModbusServer__Option = QModbusServer__Option(4)
	QModbusServer__ServerIdentifier      QModbusServer__Option = QModbusServer__Option(5)
	QModbusServer__RunIndicatorStatus    QModbusServer__Option = QModbusServer__Option(6)
	QModbusServer__AdditionalData        QModbusServer__Option = QModbusServer__Option(7)
	QModbusServer__DeviceIdentification  QModbusServer__Option = QModbusServer__Option(8)
	QModbusServer__UserOption            QModbusServer__Option = QModbusServer__Option(0x100)
)

func NewQModbusServer(parent core.QObject_ITF) *QModbusServer {

	return internal.CallLocalFunction([]interface{}{"", "", "serialbus.NewQModbusServer", "", parent}).(*QModbusServer)
}

func (ptr *QModbusServer) Data(newData QModbusDataUnit_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Data", newData}).(bool)
}

func (ptr *QModbusServer) Data2(table QModbusDataUnit__RegisterType, address uint16, data uint16) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Data2", table, address, data}).(bool)
}

func (ptr *QModbusServer) ConnectDataWritten(f func(table QModbusDataUnit__RegisterType, address int, size int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDataWritten", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QModbusServer) DisconnectDataWritten() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDataWritten"})
}

func (ptr *QModbusServer) DataWritten(table QModbusDataUnit__RegisterType, address int, size int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DataWritten", table, address, size})
}

func (ptr *QModbusServer) ConnectProcessPrivateRequest(f func(request *QModbusPdu) *QModbusResponse) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectProcessPrivateRequest", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QModbusServer) DisconnectProcessPrivateRequest() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectProcessPrivateRequest"})
}

func (ptr *QModbusServer) ProcessPrivateRequest(request QModbusPdu_ITF) *QModbusResponse {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ProcessPrivateRequest", request}).(*QModbusResponse)
}

func (ptr *QModbusServer) ProcessPrivateRequestDefault(request QModbusPdu_ITF) *QModbusResponse {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ProcessPrivateRequestDefault", request}).(*QModbusResponse)
}

func (ptr *QModbusServer) ConnectProcessRequest(f func(request *QModbusPdu) *QModbusResponse) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectProcessRequest", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QModbusServer) DisconnectProcessRequest() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectProcessRequest"})
}

func (ptr *QModbusServer) ProcessRequest(request QModbusPdu_ITF) *QModbusResponse {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ProcessRequest", request}).(*QModbusResponse)
}

func (ptr *QModbusServer) ProcessRequestDefault(request QModbusPdu_ITF) *QModbusResponse {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ProcessRequestDefault", request}).(*QModbusResponse)
}

func (ptr *QModbusServer) ConnectProcessesBroadcast(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectProcessesBroadcast", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QModbusServer) DisconnectProcessesBroadcast() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectProcessesBroadcast"})
}

func (ptr *QModbusServer) ProcessesBroadcast() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ProcessesBroadcast"}).(bool)
}

func (ptr *QModbusServer) ProcessesBroadcastDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ProcessesBroadcastDefault"}).(bool)
}

func (ptr *QModbusServer) ConnectReadData(f func(newData *QModbusDataUnit) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectReadData", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QModbusServer) DisconnectReadData() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectReadData"})
}

func (ptr *QModbusServer) ReadData(newData QModbusDataUnit_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReadData", newData}).(bool)
}

func (ptr *QModbusServer) ReadDataDefault(newData QModbusDataUnit_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReadDataDefault", newData}).(bool)
}

func (ptr *QModbusServer) ServerAddress() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ServerAddress"}).(float64))
}

func (ptr *QModbusServer) SetData(newData QModbusDataUnit_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetData", newData}).(bool)
}

func (ptr *QModbusServer) SetData2(table QModbusDataUnit__RegisterType, address uint16, data uint16) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetData2", table, address, data}).(bool)
}

func (ptr *QModbusServer) SetServerAddress(serverAddress int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetServerAddress", serverAddress})
}

func (ptr *QModbusServer) ConnectSetValue(f func(option int, newValue *core.QVariant) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetValue", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QModbusServer) DisconnectSetValue() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetValue"})
}

func (ptr *QModbusServer) SetValue(option int, newValue core.QVariant_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetValue", option, newValue}).(bool)
}

func (ptr *QModbusServer) SetValueDefault(option int, newValue core.QVariant_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetValueDefault", option, newValue}).(bool)
}

func (ptr *QModbusServer) ConnectValue(f func(option int) *core.QVariant) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectValue", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QModbusServer) DisconnectValue() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectValue"})
}

func (ptr *QModbusServer) Value(option int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Value", option}).(*core.QVariant)
}

func (ptr *QModbusServer) ValueDefault(option int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ValueDefault", option}).(*core.QVariant)
}

func (ptr *QModbusServer) ConnectWriteData(f func(newData *QModbusDataUnit) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectWriteData", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QModbusServer) DisconnectWriteData() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectWriteData"})
}

func (ptr *QModbusServer) WriteData(newData QModbusDataUnit_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WriteData", newData}).(bool)
}

func (ptr *QModbusServer) WriteDataDefault(newData QModbusDataUnit_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WriteDataDefault", newData}).(bool)
}

func (ptr *QModbusServer) Close() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Close"})
}

func (ptr *QModbusServer) CloseDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseDefault"})
}

func (ptr *QModbusServer) Open() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Open"}).(bool)
}

func (ptr *QModbusServer) OpenDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OpenDefault"}).(bool)
}

type QModbusTcpClient struct {
	QModbusClient
}

type QModbusTcpClient_ITF interface {
	QModbusClient_ITF
	QModbusTcpClient_PTR() *QModbusTcpClient
}

func (ptr *QModbusTcpClient) QModbusTcpClient_PTR() *QModbusTcpClient {
	return ptr
}

func (ptr *QModbusTcpClient) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QModbusClient_PTR().Pointer()
	}
	return nil
}

func (ptr *QModbusTcpClient) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QModbusClient_PTR().SetPointer(p)
	}
}

func PointerFromQModbusTcpClient(ptr QModbusTcpClient_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QModbusTcpClient_PTR().Pointer()
	}
	return nil
}

func (n *QModbusTcpClient) InitFromInternal(ptr uintptr, name string) {
	n.QModbusClient_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QModbusTcpClient) ClassNameInternalF() string {
	return n.QModbusClient_PTR().ClassNameInternalF()
}

func NewQModbusTcpClientFromPointer(ptr unsafe.Pointer) (n *QModbusTcpClient) {
	n = new(QModbusTcpClient)
	n.InitFromInternal(uintptr(ptr), "serialbus.QModbusTcpClient")
	return
}
func NewQModbusTcpClient(parent core.QObject_ITF) *QModbusTcpClient {

	return internal.CallLocalFunction([]interface{}{"", "", "serialbus.NewQModbusTcpClient", "", parent}).(*QModbusTcpClient)
}

func (ptr *QModbusTcpClient) ConnectClose(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectClose", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QModbusTcpClient) DisconnectClose() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectClose"})
}

func (ptr *QModbusTcpClient) Close() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Close"})
}

func (ptr *QModbusTcpClient) CloseDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseDefault"})
}

func (ptr *QModbusTcpClient) ConnectOpen(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectOpen", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QModbusTcpClient) DisconnectOpen() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectOpen"})
}

func (ptr *QModbusTcpClient) Open() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Open"}).(bool)
}

func (ptr *QModbusTcpClient) OpenDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OpenDefault"}).(bool)
}

func (ptr *QModbusTcpClient) ConnectDestroyQModbusTcpClient(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQModbusTcpClient", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QModbusTcpClient) DisconnectDestroyQModbusTcpClient() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQModbusTcpClient"})
}

func (ptr *QModbusTcpClient) DestroyQModbusTcpClient() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQModbusTcpClient"})
}

func (ptr *QModbusTcpClient) DestroyQModbusTcpClientDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQModbusTcpClientDefault"})
}

type QModbusTcpConnectionObserver struct {
	internal.Internal
}

type QModbusTcpConnectionObserver_ITF interface {
	QModbusTcpConnectionObserver_PTR() *QModbusTcpConnectionObserver
}

func (ptr *QModbusTcpConnectionObserver) QModbusTcpConnectionObserver_PTR() *QModbusTcpConnectionObserver {
	return ptr
}

func (ptr *QModbusTcpConnectionObserver) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QModbusTcpConnectionObserver) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQModbusTcpConnectionObserver(ptr QModbusTcpConnectionObserver_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QModbusTcpConnectionObserver_PTR().Pointer()
	}
	return nil
}

func (n *QModbusTcpConnectionObserver) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQModbusTcpConnectionObserverFromPointer(ptr unsafe.Pointer) (n *QModbusTcpConnectionObserver) {
	n = new(QModbusTcpConnectionObserver)
	n.InitFromInternal(uintptr(ptr), "serialbus.QModbusTcpConnectionObserver")
	return
}

func (ptr *QModbusTcpConnectionObserver) DestroyQModbusTcpConnectionObserver() {
}

func (ptr *QModbusTcpConnectionObserver) ConnectAcceptNewConnection(f func(newClient *network.QTcpSocket) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAcceptNewConnection", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QModbusTcpConnectionObserver) DisconnectAcceptNewConnection() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAcceptNewConnection"})
}

func (ptr *QModbusTcpConnectionObserver) AcceptNewConnection(newClient network.QTcpSocket_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AcceptNewConnection", newClient}).(bool)
}

type QModbusTcpServer struct {
	QModbusServer
}

type QModbusTcpServer_ITF interface {
	QModbusServer_ITF
	QModbusTcpServer_PTR() *QModbusTcpServer
}

func (ptr *QModbusTcpServer) QModbusTcpServer_PTR() *QModbusTcpServer {
	return ptr
}

func (ptr *QModbusTcpServer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QModbusServer_PTR().Pointer()
	}
	return nil
}

func (ptr *QModbusTcpServer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QModbusServer_PTR().SetPointer(p)
	}
}

func PointerFromQModbusTcpServer(ptr QModbusTcpServer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QModbusTcpServer_PTR().Pointer()
	}
	return nil
}

func (n *QModbusTcpServer) InitFromInternal(ptr uintptr, name string) {
	n.QModbusServer_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QModbusTcpServer) ClassNameInternalF() string {
	return n.QModbusServer_PTR().ClassNameInternalF()
}

func NewQModbusTcpServerFromPointer(ptr unsafe.Pointer) (n *QModbusTcpServer) {
	n = new(QModbusTcpServer)
	n.InitFromInternal(uintptr(ptr), "serialbus.QModbusTcpServer")
	return
}
func NewQModbusTcpServer(parent core.QObject_ITF) *QModbusTcpServer {

	return internal.CallLocalFunction([]interface{}{"", "", "serialbus.NewQModbusTcpServer", "", parent}).(*QModbusTcpServer)
}

func (ptr *QModbusTcpServer) ConnectClose(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectClose", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QModbusTcpServer) DisconnectClose() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectClose"})
}

func (ptr *QModbusTcpServer) Close() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Close"})
}

func (ptr *QModbusTcpServer) CloseDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseDefault"})
}

func (ptr *QModbusTcpServer) InstallConnectionObserver(observer QModbusTcpConnectionObserver_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InstallConnectionObserver", observer})
}

func (ptr *QModbusTcpServer) ConnectModbusClientDisconnected(f func(modbusClient *network.QTcpSocket)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectModbusClientDisconnected", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QModbusTcpServer) DisconnectModbusClientDisconnected() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectModbusClientDisconnected"})
}

func (ptr *QModbusTcpServer) ModbusClientDisconnected(modbusClient network.QTcpSocket_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ModbusClientDisconnected", modbusClient})
}

func (ptr *QModbusTcpServer) ConnectOpen(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectOpen", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QModbusTcpServer) DisconnectOpen() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectOpen"})
}

func (ptr *QModbusTcpServer) Open() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Open"}).(bool)
}

func (ptr *QModbusTcpServer) OpenDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OpenDefault"}).(bool)
}

func (ptr *QModbusTcpServer) ConnectDestroyQModbusTcpServer(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQModbusTcpServer", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QModbusTcpServer) DisconnectDestroyQModbusTcpServer() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQModbusTcpServer"})
}

func (ptr *QModbusTcpServer) DestroyQModbusTcpServer() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQModbusTcpServer"})
}

func (ptr *QModbusTcpServer) DestroyQModbusTcpServerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQModbusTcpServerDefault"})
}

func init() {
	internal.ConstructorTable["serialbus.QCanBus"] = NewQCanBusFromPointer
	internal.ConstructorTable["serialbus.QCanBusDevice"] = NewQCanBusDeviceFromPointer
	internal.ConstructorTable["serialbus.QCanBusDeviceInfo"] = NewQCanBusDeviceInfoFromPointer
	internal.ConstructorTable["serialbus.QCanBusFactory"] = NewQCanBusFactoryFromPointer
	internal.ConstructorTable["serialbus.QCanBusFactoryV2"] = NewQCanBusFactoryV2FromPointer
	internal.ConstructorTable["serialbus.QCanBusFrame"] = NewQCanBusFrameFromPointer
	internal.ConstructorTable["serialbus.QModbusClient"] = NewQModbusClientFromPointer
	internal.ConstructorTable["serialbus.QModbusDataUnit"] = NewQModbusDataUnitFromPointer
	internal.ConstructorTable["serialbus.QModbusDevice"] = NewQModbusDeviceFromPointer
	internal.ConstructorTable["serialbus.QModbusDeviceIdentification"] = NewQModbusDeviceIdentificationFromPointer
	internal.ConstructorTable["serialbus.QModbusExceptionResponse"] = NewQModbusExceptionResponseFromPointer
	internal.ConstructorTable["serialbus.QModbusPdu"] = NewQModbusPduFromPointer
	internal.ConstructorTable["serialbus.QModbusReply"] = NewQModbusReplyFromPointer
	internal.ConstructorTable["serialbus.QModbusRequest"] = NewQModbusRequestFromPointer
	internal.ConstructorTable["serialbus.QModbusResponse"] = NewQModbusResponseFromPointer
	internal.ConstructorTable["serialbus.QModbusRtuSerialMaster"] = NewQModbusRtuSerialMasterFromPointer
	internal.ConstructorTable["serialbus.QModbusRtuSerialSlave"] = NewQModbusRtuSerialSlaveFromPointer
	internal.ConstructorTable["serialbus.QModbusServer"] = NewQModbusServerFromPointer
	internal.ConstructorTable["serialbus.QModbusTcpClient"] = NewQModbusTcpClientFromPointer
	internal.ConstructorTable["serialbus.QModbusTcpConnectionObserver"] = NewQModbusTcpConnectionObserverFromPointer
	internal.ConstructorTable["serialbus.QModbusTcpServer"] = NewQModbusTcpServerFromPointer
}
