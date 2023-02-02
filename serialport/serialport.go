// +build !minimal

package serialport

import (
	"github.com/bluszcz/cutego"
	"github.com/bluszcz/cutego/core"
	"reflect"
	"strings"
	"unsafe"
)

type QSerialPort struct {
	core.QIODevice
}

type QSerialPort_ITF interface {
	core.QIODevice_ITF
	QSerialPort_PTR() *QSerialPort
}

func (ptr *QSerialPort) QSerialPort_PTR() *QSerialPort {
	return ptr
}

func (ptr *QSerialPort) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QIODevice_PTR().Pointer()
	}
	return nil
}

func (ptr *QSerialPort) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QIODevice_PTR().SetPointer(p)
	}
}

func PointerFromQSerialPort(ptr QSerialPort_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSerialPort_PTR().Pointer()
	}
	return nil
}

func (n *QSerialPort) InitFromInternal(ptr uintptr, name string) {
	n.QIODevice_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSerialPort) ClassNameInternalF() string {
	return n.QIODevice_PTR().ClassNameInternalF()
}

func NewQSerialPortFromPointer(ptr unsafe.Pointer) (n *QSerialPort) {
	n = new(QSerialPort)
	n.InitFromInternal(uintptr(ptr), "serialport.QSerialPort")
	return
}

//go:generate stringer -type=QSerialPort__Direction
//QSerialPort::Direction
type QSerialPort__Direction int64

const (
	QSerialPort__Input         QSerialPort__Direction = QSerialPort__Direction(1)
	QSerialPort__Output        QSerialPort__Direction = QSerialPort__Direction(2)
	QSerialPort__AllDirections QSerialPort__Direction = QSerialPort__Direction(QSerialPort__Input | QSerialPort__Output)
)

//go:generate stringer -type=QSerialPort__BaudRate
//QSerialPort::BaudRate
type QSerialPort__BaudRate int64

const (
	QSerialPort__Baud1200    QSerialPort__BaudRate = QSerialPort__BaudRate(1200)
	QSerialPort__Baud2400    QSerialPort__BaudRate = QSerialPort__BaudRate(2400)
	QSerialPort__Baud4800    QSerialPort__BaudRate = QSerialPort__BaudRate(4800)
	QSerialPort__Baud9600    QSerialPort__BaudRate = QSerialPort__BaudRate(9600)
	QSerialPort__Baud19200   QSerialPort__BaudRate = QSerialPort__BaudRate(19200)
	QSerialPort__Baud38400   QSerialPort__BaudRate = QSerialPort__BaudRate(38400)
	QSerialPort__Baud57600   QSerialPort__BaudRate = QSerialPort__BaudRate(57600)
	QSerialPort__Baud115200  QSerialPort__BaudRate = QSerialPort__BaudRate(115200)
	QSerialPort__UnknownBaud QSerialPort__BaudRate = QSerialPort__BaudRate(-1)
)

//go:generate stringer -type=QSerialPort__DataBits
//QSerialPort::DataBits
type QSerialPort__DataBits int64

const (
	QSerialPort__Data5           QSerialPort__DataBits = QSerialPort__DataBits(5)
	QSerialPort__Data6           QSerialPort__DataBits = QSerialPort__DataBits(6)
	QSerialPort__Data7           QSerialPort__DataBits = QSerialPort__DataBits(7)
	QSerialPort__Data8           QSerialPort__DataBits = QSerialPort__DataBits(8)
	QSerialPort__UnknownDataBits QSerialPort__DataBits = QSerialPort__DataBits(-1)
)

//go:generate stringer -type=QSerialPort__Parity
//QSerialPort::Parity
type QSerialPort__Parity int64

const (
	QSerialPort__NoParity      QSerialPort__Parity = QSerialPort__Parity(0)
	QSerialPort__EvenParity    QSerialPort__Parity = QSerialPort__Parity(2)
	QSerialPort__OddParity     QSerialPort__Parity = QSerialPort__Parity(3)
	QSerialPort__SpaceParity   QSerialPort__Parity = QSerialPort__Parity(4)
	QSerialPort__MarkParity    QSerialPort__Parity = QSerialPort__Parity(5)
	QSerialPort__UnknownParity QSerialPort__Parity = QSerialPort__Parity(-1)
)

//go:generate stringer -type=QSerialPort__StopBits
//QSerialPort::StopBits
type QSerialPort__StopBits int64

const (
	QSerialPort__OneStop         QSerialPort__StopBits = QSerialPort__StopBits(1)
	QSerialPort__OneAndHalfStop  QSerialPort__StopBits = QSerialPort__StopBits(3)
	QSerialPort__TwoStop         QSerialPort__StopBits = QSerialPort__StopBits(2)
	QSerialPort__UnknownStopBits QSerialPort__StopBits = QSerialPort__StopBits(-1)
)

//go:generate stringer -type=QSerialPort__FlowControl
//QSerialPort::FlowControl
type QSerialPort__FlowControl int64

const (
	QSerialPort__NoFlowControl      QSerialPort__FlowControl = QSerialPort__FlowControl(0)
	QSerialPort__HardwareControl    QSerialPort__FlowControl = QSerialPort__FlowControl(1)
	QSerialPort__SoftwareControl    QSerialPort__FlowControl = QSerialPort__FlowControl(2)
	QSerialPort__UnknownFlowControl QSerialPort__FlowControl = QSerialPort__FlowControl(-1)
)

//go:generate stringer -type=QSerialPort__PinoutSignal
//QSerialPort::PinoutSignal
type QSerialPort__PinoutSignal int64

const (
	QSerialPort__NoSignal                       QSerialPort__PinoutSignal = QSerialPort__PinoutSignal(0x00)
	QSerialPort__TransmittedDataSignal          QSerialPort__PinoutSignal = QSerialPort__PinoutSignal(0x01)
	QSerialPort__ReceivedDataSignal             QSerialPort__PinoutSignal = QSerialPort__PinoutSignal(0x02)
	QSerialPort__DataTerminalReadySignal        QSerialPort__PinoutSignal = QSerialPort__PinoutSignal(0x04)
	QSerialPort__DataCarrierDetectSignal        QSerialPort__PinoutSignal = QSerialPort__PinoutSignal(0x08)
	QSerialPort__DataSetReadySignal             QSerialPort__PinoutSignal = QSerialPort__PinoutSignal(0x10)
	QSerialPort__RingIndicatorSignal            QSerialPort__PinoutSignal = QSerialPort__PinoutSignal(0x20)
	QSerialPort__RequestToSendSignal            QSerialPort__PinoutSignal = QSerialPort__PinoutSignal(0x40)
	QSerialPort__ClearToSendSignal              QSerialPort__PinoutSignal = QSerialPort__PinoutSignal(0x80)
	QSerialPort__SecondaryTransmittedDataSignal QSerialPort__PinoutSignal = QSerialPort__PinoutSignal(0x100)
	QSerialPort__SecondaryReceivedDataSignal    QSerialPort__PinoutSignal = QSerialPort__PinoutSignal(0x200)
)

//go:generate stringer -type=QSerialPort__SerialPortError
//QSerialPort::SerialPortError
type QSerialPort__SerialPortError int64

const (
	QSerialPort__NoError                   QSerialPort__SerialPortError = QSerialPort__SerialPortError(0)
	QSerialPort__DeviceNotFoundError       QSerialPort__SerialPortError = QSerialPort__SerialPortError(1)
	QSerialPort__PermissionError           QSerialPort__SerialPortError = QSerialPort__SerialPortError(2)
	QSerialPort__OpenError                 QSerialPort__SerialPortError = QSerialPort__SerialPortError(3)
	QSerialPort__ParityError               QSerialPort__SerialPortError = QSerialPort__SerialPortError(4)
	QSerialPort__FramingError              QSerialPort__SerialPortError = QSerialPort__SerialPortError(5)
	QSerialPort__BreakConditionError       QSerialPort__SerialPortError = QSerialPort__SerialPortError(6)
	QSerialPort__WriteError                QSerialPort__SerialPortError = QSerialPort__SerialPortError(7)
	QSerialPort__ReadError                 QSerialPort__SerialPortError = QSerialPort__SerialPortError(8)
	QSerialPort__ResourceError             QSerialPort__SerialPortError = QSerialPort__SerialPortError(9)
	QSerialPort__UnsupportedOperationError QSerialPort__SerialPortError = QSerialPort__SerialPortError(10)
	QSerialPort__UnknownError              QSerialPort__SerialPortError = QSerialPort__SerialPortError(11)
	QSerialPort__TimeoutError              QSerialPort__SerialPortError = QSerialPort__SerialPortError(12)
	QSerialPort__NotOpenError              QSerialPort__SerialPortError = QSerialPort__SerialPortError(13)
)

func NewQSerialPort(parent core.QObject_ITF) *QSerialPort {

	return internal.CallLocalFunction([]interface{}{"", "", "serialport.NewQSerialPort", "", parent}).(*QSerialPort)
}

func NewQSerialPort2(name string, parent core.QObject_ITF) *QSerialPort {

	return internal.CallLocalFunction([]interface{}{"", "", "serialport.NewQSerialPort2", "", name, parent}).(*QSerialPort)
}

func NewQSerialPort3(serialPortInfo QSerialPortInfo_ITF, parent core.QObject_ITF) *QSerialPort {

	return internal.CallLocalFunction([]interface{}{"", "", "serialport.NewQSerialPort3", "", serialPortInfo, parent}).(*QSerialPort)
}

func (ptr *QSerialPort) AtEndDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AtEndDefault"}).(bool)
}

func (ptr *QSerialPort) BaudRate(directions QSerialPort__Direction) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BaudRate", directions}).(float64))
}

func (ptr *QSerialPort) ConnectBaudRateChanged(f func(baudRate int, directions QSerialPort__Direction)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBaudRateChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSerialPort) DisconnectBaudRateChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBaudRateChanged"})
}

func (ptr *QSerialPort) BaudRateChanged(baudRate int, directions QSerialPort__Direction) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BaudRateChanged", baudRate, directions})
}

func (ptr *QSerialPort) ConnectBreakEnabledChanged(f func(set bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBreakEnabledChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSerialPort) DisconnectBreakEnabledChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBreakEnabledChanged"})
}

func (ptr *QSerialPort) BreakEnabledChanged(set bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BreakEnabledChanged", set})
}

func (ptr *QSerialPort) BytesAvailableDefault() int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BytesAvailableDefault"}).(float64))
}

func (ptr *QSerialPort) BytesToWriteDefault() int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BytesToWriteDefault"}).(float64))
}

func (ptr *QSerialPort) CanReadLineDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CanReadLineDefault"}).(bool)
}

func (ptr *QSerialPort) Clear(directions QSerialPort__Direction) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Clear", directions}).(bool)
}

func (ptr *QSerialPort) ClearError() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ClearError"})
}

func (ptr *QSerialPort) CloseDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseDefault"})
}

func (ptr *QSerialPort) DataBits() QSerialPort__DataBits {

	return QSerialPort__DataBits(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DataBits"}).(float64))
}

func (ptr *QSerialPort) ConnectDataBitsChanged(f func(dataBits QSerialPort__DataBits)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDataBitsChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSerialPort) DisconnectDataBitsChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDataBitsChanged"})
}

func (ptr *QSerialPort) DataBitsChanged(dataBits QSerialPort__DataBits) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DataBitsChanged", dataBits})
}

func (ptr *QSerialPort) ConnectDataTerminalReadyChanged(f func(set bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDataTerminalReadyChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSerialPort) DisconnectDataTerminalReadyChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDataTerminalReadyChanged"})
}

func (ptr *QSerialPort) DataTerminalReadyChanged(set bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DataTerminalReadyChanged", set})
}

func (ptr *QSerialPort) Error() QSerialPort__SerialPortError {

	return QSerialPort__SerialPortError(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error"}).(float64))
}

func (ptr *QSerialPort) ConnectErrorOccurred(f func(error QSerialPort__SerialPortError)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectErrorOccurred", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSerialPort) DisconnectErrorOccurred() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectErrorOccurred"})
}

func (ptr *QSerialPort) ErrorOccurred(error QSerialPort__SerialPortError) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ErrorOccurred", error})
}

func (ptr *QSerialPort) FlowControl() QSerialPort__FlowControl {

	return QSerialPort__FlowControl(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FlowControl"}).(float64))
}

func (ptr *QSerialPort) ConnectFlowControlChanged(f func(flow QSerialPort__FlowControl)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFlowControlChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSerialPort) DisconnectFlowControlChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFlowControlChanged"})
}

func (ptr *QSerialPort) FlowControlChanged(flow QSerialPort__FlowControl) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FlowControlChanged", flow})
}

func (ptr *QSerialPort) Flush() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Flush"}).(bool)
}

func (ptr *QSerialPort) IsBreakEnabled() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsBreakEnabled"}).(bool)
}

func (ptr *QSerialPort) IsDataTerminalReady() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsDataTerminalReady"}).(bool)
}

func (ptr *QSerialPort) IsRequestToSend() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsRequestToSend"}).(bool)
}

func (ptr *QSerialPort) IsSequentialDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsSequentialDefault"}).(bool)
}

func (ptr *QSerialPort) OpenDefault(mode core.QIODevice__OpenModeFlag) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OpenDefault", mode}).(bool)
}

func (ptr *QSerialPort) Parity() QSerialPort__Parity {

	return QSerialPort__Parity(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Parity"}).(float64))
}

func (ptr *QSerialPort) ConnectParityChanged(f func(parity QSerialPort__Parity)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectParityChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSerialPort) DisconnectParityChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectParityChanged"})
}

func (ptr *QSerialPort) ParityChanged(parity QSerialPort__Parity) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ParityChanged", parity})
}

func (ptr *QSerialPort) PinoutSignals() QSerialPort__PinoutSignal {

	return QSerialPort__PinoutSignal(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PinoutSignals"}).(float64))
}

func (ptr *QSerialPort) PortName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PortName"}).(string)
}

func (ptr *QSerialPort) ReadBufferSize() int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReadBufferSize"}).(float64))
}

func (ptr *QSerialPort) ConnectReadData(f func(data *string, maxSize int64) int64) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectReadData", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSerialPort) DisconnectReadData() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectReadData"})
}

func (ptr *QSerialPort) ReadData(data *string, maxSize int64) int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReadData", data, maxSize}).(float64))
}

func (ptr *QSerialPort) ReadDataDefault(data *string, maxSize int64) int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReadDataDefault", data, maxSize}).(float64))
}

func (ptr *QSerialPort) ReadLineDataDefault(data []byte, maxSize int64) int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReadLineDataDefault", data, maxSize}).(float64))
}

func (ptr *QSerialPort) ConnectRequestToSendChanged(f func(set bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRequestToSendChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSerialPort) DisconnectRequestToSendChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRequestToSendChanged"})
}

func (ptr *QSerialPort) RequestToSendChanged(set bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RequestToSendChanged", set})
}

func (ptr *QSerialPort) SendBreak(duration int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SendBreak", duration}).(bool)
}

func (ptr *QSerialPort) SetBaudRate(baudRate int, directions QSerialPort__Direction) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBaudRate", baudRate, directions}).(bool)
}

func (ptr *QSerialPort) SetBreakEnabled(set bool) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBreakEnabled", set}).(bool)
}

func (ptr *QSerialPort) SetDataBits(dataBits QSerialPort__DataBits) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDataBits", dataBits}).(bool)
}

func (ptr *QSerialPort) SetDataTerminalReady(set bool) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDataTerminalReady", set}).(bool)
}

func (ptr *QSerialPort) SetFlowControl(flowControl QSerialPort__FlowControl) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFlowControl", flowControl}).(bool)
}

func (ptr *QSerialPort) SetParity(parity QSerialPort__Parity) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetParity", parity}).(bool)
}

func (ptr *QSerialPort) SetPort(serialPortInfo QSerialPortInfo_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPort", serialPortInfo})
}

func (ptr *QSerialPort) SetPortName(name string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPortName", name})
}

func (ptr *QSerialPort) SetReadBufferSize(size int64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetReadBufferSize", size})
}

func (ptr *QSerialPort) SetRequestToSend(set bool) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRequestToSend", set}).(bool)
}

func (ptr *QSerialPort) SetSettingsRestoredOnClose(restore bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSettingsRestoredOnClose", restore})
}

func (ptr *QSerialPort) SetStopBits(stopBits QSerialPort__StopBits) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetStopBits", stopBits}).(bool)
}

func (ptr *QSerialPort) SettingsRestoredOnClose() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SettingsRestoredOnClose"}).(bool)
}

func (ptr *QSerialPort) StopBits() QSerialPort__StopBits {

	return QSerialPort__StopBits(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StopBits"}).(float64))
}

func (ptr *QSerialPort) ConnectStopBitsChanged(f func(stopBits QSerialPort__StopBits)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStopBitsChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSerialPort) DisconnectStopBitsChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStopBitsChanged"})
}

func (ptr *QSerialPort) StopBitsChanged(stopBits QSerialPort__StopBits) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StopBitsChanged", stopBits})
}

func (ptr *QSerialPort) WaitForBytesWrittenDefault(msecs int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WaitForBytesWrittenDefault", msecs}).(bool)
}

func (ptr *QSerialPort) WaitForReadyReadDefault(msecs int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WaitForReadyReadDefault", msecs}).(bool)
}

func (ptr *QSerialPort) ConnectWriteData(f func(data []byte, maxSize int64) int64) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectWriteData", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSerialPort) DisconnectWriteData() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectWriteData"})
}

func (ptr *QSerialPort) WriteData(data []byte, maxSize int64) int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WriteData", data, maxSize}).(float64))
}

func (ptr *QSerialPort) WriteDataDefault(data []byte, maxSize int64) int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WriteDataDefault", data, maxSize}).(float64))
}

func (ptr *QSerialPort) ConnectDestroyQSerialPort(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQSerialPort", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSerialPort) DisconnectDestroyQSerialPort() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQSerialPort"})
}

func (ptr *QSerialPort) DestroyQSerialPort() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSerialPort"})
}

func (ptr *QSerialPort) DestroyQSerialPortDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSerialPortDefault"})
}

func (ptr *QSerialPort) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QSerialPort) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QSerialPort) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QSerialPort) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QSerialPort) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QSerialPort) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QSerialPort) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QSerialPort) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QSerialPort) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QSerialPort) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QSerialPort) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QSerialPort) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QSerialPort) PosDefault() int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PosDefault"}).(float64))
}

func (ptr *QSerialPort) ResetDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResetDefault"}).(bool)
}

func (ptr *QSerialPort) SeekDefault(pos int64) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SeekDefault", pos}).(bool)
}

func (ptr *QSerialPort) SizeDefault() int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SizeDefault"}).(float64))
}

func (ptr *QSerialPort) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QSerialPort) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QSerialPort) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QSerialPort) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QSerialPort) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QSerialPort) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QSerialPort) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QSerialPort) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QSerialPort) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QSerialPortInfo struct {
	internal.Internal
}

type QSerialPortInfo_ITF interface {
	QSerialPortInfo_PTR() *QSerialPortInfo
}

func (ptr *QSerialPortInfo) QSerialPortInfo_PTR() *QSerialPortInfo {
	return ptr
}

func (ptr *QSerialPortInfo) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QSerialPortInfo) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQSerialPortInfo(ptr QSerialPortInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSerialPortInfo_PTR().Pointer()
	}
	return nil
}

func (n *QSerialPortInfo) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQSerialPortInfoFromPointer(ptr unsafe.Pointer) (n *QSerialPortInfo) {
	n = new(QSerialPortInfo)
	n.InitFromInternal(uintptr(ptr), "serialport.QSerialPortInfo")
	return
}
func NewQSerialPortInfo() *QSerialPortInfo {

	return internal.CallLocalFunction([]interface{}{"", "", "serialport.NewQSerialPortInfo", ""}).(*QSerialPortInfo)
}

func NewQSerialPortInfo2(port QSerialPort_ITF) *QSerialPortInfo {

	return internal.CallLocalFunction([]interface{}{"", "", "serialport.NewQSerialPortInfo2", "", port}).(*QSerialPortInfo)
}

func NewQSerialPortInfo3(name string) *QSerialPortInfo {

	return internal.CallLocalFunction([]interface{}{"", "", "serialport.NewQSerialPortInfo3", "", name}).(*QSerialPortInfo)
}

func NewQSerialPortInfo4(other QSerialPortInfo_ITF) *QSerialPortInfo {

	return internal.CallLocalFunction([]interface{}{"", "", "serialport.NewQSerialPortInfo4", "", other}).(*QSerialPortInfo)
}

func QSerialPortInfo_AvailablePorts() []*QSerialPortInfo {

	return internal.CallLocalFunction([]interface{}{"", "", "serialport.QSerialPortInfo_AvailablePorts", ""}).([]*QSerialPortInfo)
}

func (ptr *QSerialPortInfo) AvailablePorts() []*QSerialPortInfo {

	return internal.CallLocalFunction([]interface{}{"", "", "serialport.QSerialPortInfo_AvailablePorts", ""}).([]*QSerialPortInfo)
}

func (ptr *QSerialPortInfo) Description() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Description"}).(string)
}

func (ptr *QSerialPortInfo) HasProductIdentifier() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasProductIdentifier"}).(bool)
}

func (ptr *QSerialPortInfo) HasVendorIdentifier() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasVendorIdentifier"}).(bool)
}

func (ptr *QSerialPortInfo) IsBusy() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsBusy"}).(bool)
}

func (ptr *QSerialPortInfo) IsNull() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsNull"}).(bool)
}

func (ptr *QSerialPortInfo) Manufacturer() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Manufacturer"}).(string)
}

func (ptr *QSerialPortInfo) PortName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PortName"}).(string)
}

func (ptr *QSerialPortInfo) ProductIdentifier() uint16 {

	return uint16(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ProductIdentifier"}).(float64))
}

func (ptr *QSerialPortInfo) SerialNumber() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SerialNumber"}).(string)
}

func QSerialPortInfo_StandardBaudRates() []int {

	return internal.CallLocalFunction([]interface{}{"", "", "serialport.QSerialPortInfo_StandardBaudRates", ""}).([]int)
}

func (ptr *QSerialPortInfo) StandardBaudRates() []int {

	return internal.CallLocalFunction([]interface{}{"", "", "serialport.QSerialPortInfo_StandardBaudRates", ""}).([]int)
}

func (ptr *QSerialPortInfo) Swap(other QSerialPortInfo_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", other})
}

func (ptr *QSerialPortInfo) SystemLocation() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SystemLocation"}).(string)
}

func (ptr *QSerialPortInfo) VendorIdentifier() uint16 {

	return uint16(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "VendorIdentifier"}).(float64))
}

func (ptr *QSerialPortInfo) DestroyQSerialPortInfo() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSerialPortInfo"})
}

func (ptr *QSerialPortInfo) __availablePorts_atList(i int) *QSerialPortInfo {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__availablePorts_atList", i}).(*QSerialPortInfo)
}

func (ptr *QSerialPortInfo) __availablePorts_setList(i QSerialPortInfo_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__availablePorts_setList", i})
}

func (ptr *QSerialPortInfo) __availablePorts_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__availablePorts_newList"}).(unsafe.Pointer)
}

func (ptr *QSerialPortInfo) __standardBaudRates_atList(i int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__standardBaudRates_atList", i}).(float64))
}

func (ptr *QSerialPortInfo) __standardBaudRates_setList(i int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__standardBaudRates_setList", i})
}

func (ptr *QSerialPortInfo) __standardBaudRates_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__standardBaudRates_newList"}).(unsafe.Pointer)
}

func init() {
	internal.ConstructorTable["serialport.QSerialPort"] = NewQSerialPortFromPointer
	internal.ConstructorTable["serialport.QSerialPortInfo"] = NewQSerialPortInfoFromPointer
}
