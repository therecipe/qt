// +build !minimal

package serialport

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "serialport.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"reflect"
	"strings"
	"unsafe"
)

func cGoFreePacked(ptr unsafe.Pointer) { core.NewQByteArrayFromPointer(ptr).DestroyQByteArray() }
func cGoUnpackString(s C.struct_QtSerialPort_PackedString) string {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtSerialPort_PackedString) []byte {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		gs := C.GoString(s.data)
		return *(*[]byte)(unsafe.Pointer(&gs))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}
func unpackStringList(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}
	return strings.Split(s, "¡¦!")
}

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

func NewQSerialPortFromPointer(ptr unsafe.Pointer) (n *QSerialPort) {
	n = new(QSerialPort)
	n.SetPointer(ptr)
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
	tmpValue := NewQSerialPortFromPointer(C.QSerialPort_NewQSerialPort(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQSerialPort2(name string, parent core.QObject_ITF) *QSerialPort {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := NewQSerialPortFromPointer(C.QSerialPort_NewQSerialPort2(C.struct_QtSerialPort_PackedString{data: nameC, len: C.longlong(len(name))}, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQSerialPort3(serialPortInfo QSerialPortInfo_ITF, parent core.QObject_ITF) *QSerialPort {
	tmpValue := NewQSerialPortFromPointer(C.QSerialPort_NewQSerialPort3(PointerFromQSerialPortInfo(serialPortInfo), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQSerialPort_AtEnd
func callbackQSerialPort_AtEnd(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "atEnd"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSerialPortFromPointer(ptr).AtEndDefault())))
}

func (ptr *QSerialPort) AtEndDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSerialPort_AtEndDefault(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSerialPort) BaudRate(directions QSerialPort__Direction) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSerialPort_BaudRate(ptr.Pointer(), C.longlong(directions))))
	}
	return 0
}

//export callbackQSerialPort_BaudRateChanged
func callbackQSerialPort_BaudRateChanged(ptr unsafe.Pointer, baudRate C.int, directions C.longlong) {
	if signal := qt.GetSignal(ptr, "baudRateChanged"); signal != nil {
		(*(*func(int, QSerialPort__Direction))(signal))(int(int32(baudRate)), QSerialPort__Direction(directions))
	}

}

func (ptr *QSerialPort) ConnectBaudRateChanged(f func(baudRate int, directions QSerialPort__Direction)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "baudRateChanged") {
			C.QSerialPort_ConnectBaudRateChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "baudRateChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "baudRateChanged"); signal != nil {
			f := func(baudRate int, directions QSerialPort__Direction) {
				(*(*func(int, QSerialPort__Direction))(signal))(baudRate, directions)
				f(baudRate, directions)
			}
			qt.ConnectSignal(ptr.Pointer(), "baudRateChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "baudRateChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSerialPort) DisconnectBaudRateChanged() {
	if ptr.Pointer() != nil {
		C.QSerialPort_DisconnectBaudRateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "baudRateChanged")
	}
}

func (ptr *QSerialPort) BaudRateChanged(baudRate int, directions QSerialPort__Direction) {
	if ptr.Pointer() != nil {
		C.QSerialPort_BaudRateChanged(ptr.Pointer(), C.int(int32(baudRate)), C.longlong(directions))
	}
}

//export callbackQSerialPort_BreakEnabledChanged
func callbackQSerialPort_BreakEnabledChanged(ptr unsafe.Pointer, set C.char) {
	if signal := qt.GetSignal(ptr, "breakEnabledChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(set) != 0)
	}

}

func (ptr *QSerialPort) ConnectBreakEnabledChanged(f func(set bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "breakEnabledChanged") {
			C.QSerialPort_ConnectBreakEnabledChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "breakEnabledChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "breakEnabledChanged"); signal != nil {
			f := func(set bool) {
				(*(*func(bool))(signal))(set)
				f(set)
			}
			qt.ConnectSignal(ptr.Pointer(), "breakEnabledChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "breakEnabledChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSerialPort) DisconnectBreakEnabledChanged() {
	if ptr.Pointer() != nil {
		C.QSerialPort_DisconnectBreakEnabledChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "breakEnabledChanged")
	}
}

func (ptr *QSerialPort) BreakEnabledChanged(set bool) {
	if ptr.Pointer() != nil {
		C.QSerialPort_BreakEnabledChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(set))))
	}
}

//export callbackQSerialPort_BytesAvailable
func callbackQSerialPort_BytesAvailable(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "bytesAvailable"); signal != nil {
		return C.longlong((*(*func() int64)(signal))())
	}

	return C.longlong(NewQSerialPortFromPointer(ptr).BytesAvailableDefault())
}

func (ptr *QSerialPort) BytesAvailableDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QSerialPort_BytesAvailableDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQSerialPort_BytesToWrite
func callbackQSerialPort_BytesToWrite(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "bytesToWrite"); signal != nil {
		return C.longlong((*(*func() int64)(signal))())
	}

	return C.longlong(NewQSerialPortFromPointer(ptr).BytesToWriteDefault())
}

func (ptr *QSerialPort) BytesToWriteDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QSerialPort_BytesToWriteDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQSerialPort_CanReadLine
func callbackQSerialPort_CanReadLine(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canReadLine"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSerialPortFromPointer(ptr).CanReadLineDefault())))
}

func (ptr *QSerialPort) CanReadLineDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSerialPort_CanReadLineDefault(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSerialPort) Clear(directions QSerialPort__Direction) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSerialPort_Clear(ptr.Pointer(), C.longlong(directions))) != 0
	}
	return false
}

func (ptr *QSerialPort) ClearError() {
	if ptr.Pointer() != nil {
		C.QSerialPort_ClearError(ptr.Pointer())
	}
}

//export callbackQSerialPort_Close
func callbackQSerialPort_Close(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQSerialPortFromPointer(ptr).CloseDefault()
	}
}

func (ptr *QSerialPort) CloseDefault() {
	if ptr.Pointer() != nil {
		C.QSerialPort_CloseDefault(ptr.Pointer())
	}
}

func (ptr *QSerialPort) DataBits() QSerialPort__DataBits {
	if ptr.Pointer() != nil {
		return QSerialPort__DataBits(C.QSerialPort_DataBits(ptr.Pointer()))
	}
	return 0
}

//export callbackQSerialPort_DataBitsChanged
func callbackQSerialPort_DataBitsChanged(ptr unsafe.Pointer, dataBits C.longlong) {
	if signal := qt.GetSignal(ptr, "dataBitsChanged"); signal != nil {
		(*(*func(QSerialPort__DataBits))(signal))(QSerialPort__DataBits(dataBits))
	}

}

func (ptr *QSerialPort) ConnectDataBitsChanged(f func(dataBits QSerialPort__DataBits)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "dataBitsChanged") {
			C.QSerialPort_ConnectDataBitsChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "dataBitsChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "dataBitsChanged"); signal != nil {
			f := func(dataBits QSerialPort__DataBits) {
				(*(*func(QSerialPort__DataBits))(signal))(dataBits)
				f(dataBits)
			}
			qt.ConnectSignal(ptr.Pointer(), "dataBitsChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "dataBitsChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSerialPort) DisconnectDataBitsChanged() {
	if ptr.Pointer() != nil {
		C.QSerialPort_DisconnectDataBitsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "dataBitsChanged")
	}
}

func (ptr *QSerialPort) DataBitsChanged(dataBits QSerialPort__DataBits) {
	if ptr.Pointer() != nil {
		C.QSerialPort_DataBitsChanged(ptr.Pointer(), C.longlong(dataBits))
	}
}

//export callbackQSerialPort_DataTerminalReadyChanged
func callbackQSerialPort_DataTerminalReadyChanged(ptr unsafe.Pointer, set C.char) {
	if signal := qt.GetSignal(ptr, "dataTerminalReadyChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(set) != 0)
	}

}

func (ptr *QSerialPort) ConnectDataTerminalReadyChanged(f func(set bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "dataTerminalReadyChanged") {
			C.QSerialPort_ConnectDataTerminalReadyChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "dataTerminalReadyChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "dataTerminalReadyChanged"); signal != nil {
			f := func(set bool) {
				(*(*func(bool))(signal))(set)
				f(set)
			}
			qt.ConnectSignal(ptr.Pointer(), "dataTerminalReadyChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "dataTerminalReadyChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSerialPort) DisconnectDataTerminalReadyChanged() {
	if ptr.Pointer() != nil {
		C.QSerialPort_DisconnectDataTerminalReadyChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "dataTerminalReadyChanged")
	}
}

func (ptr *QSerialPort) DataTerminalReadyChanged(set bool) {
	if ptr.Pointer() != nil {
		C.QSerialPort_DataTerminalReadyChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(set))))
	}
}

func (ptr *QSerialPort) Error() QSerialPort__SerialPortError {
	if ptr.Pointer() != nil {
		return QSerialPort__SerialPortError(C.QSerialPort_Error(ptr.Pointer()))
	}
	return 0
}

//export callbackQSerialPort_ErrorOccurred
func callbackQSerialPort_ErrorOccurred(ptr unsafe.Pointer, error C.longlong) {
	if signal := qt.GetSignal(ptr, "errorOccurred"); signal != nil {
		(*(*func(QSerialPort__SerialPortError))(signal))(QSerialPort__SerialPortError(error))
	}

}

func (ptr *QSerialPort) ConnectErrorOccurred(f func(error QSerialPort__SerialPortError)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "errorOccurred") {
			C.QSerialPort_ConnectErrorOccurred(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "errorOccurred")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "errorOccurred"); signal != nil {
			f := func(error QSerialPort__SerialPortError) {
				(*(*func(QSerialPort__SerialPortError))(signal))(error)
				f(error)
			}
			qt.ConnectSignal(ptr.Pointer(), "errorOccurred", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "errorOccurred", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSerialPort) DisconnectErrorOccurred() {
	if ptr.Pointer() != nil {
		C.QSerialPort_DisconnectErrorOccurred(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "errorOccurred")
	}
}

func (ptr *QSerialPort) ErrorOccurred(error QSerialPort__SerialPortError) {
	if ptr.Pointer() != nil {
		C.QSerialPort_ErrorOccurred(ptr.Pointer(), C.longlong(error))
	}
}

func (ptr *QSerialPort) FlowControl() QSerialPort__FlowControl {
	if ptr.Pointer() != nil {
		return QSerialPort__FlowControl(C.QSerialPort_FlowControl(ptr.Pointer()))
	}
	return 0
}

//export callbackQSerialPort_FlowControlChanged
func callbackQSerialPort_FlowControlChanged(ptr unsafe.Pointer, flow C.longlong) {
	if signal := qt.GetSignal(ptr, "flowControlChanged"); signal != nil {
		(*(*func(QSerialPort__FlowControl))(signal))(QSerialPort__FlowControl(flow))
	}

}

func (ptr *QSerialPort) ConnectFlowControlChanged(f func(flow QSerialPort__FlowControl)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "flowControlChanged") {
			C.QSerialPort_ConnectFlowControlChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "flowControlChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "flowControlChanged"); signal != nil {
			f := func(flow QSerialPort__FlowControl) {
				(*(*func(QSerialPort__FlowControl))(signal))(flow)
				f(flow)
			}
			qt.ConnectSignal(ptr.Pointer(), "flowControlChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "flowControlChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSerialPort) DisconnectFlowControlChanged() {
	if ptr.Pointer() != nil {
		C.QSerialPort_DisconnectFlowControlChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "flowControlChanged")
	}
}

func (ptr *QSerialPort) FlowControlChanged(flow QSerialPort__FlowControl) {
	if ptr.Pointer() != nil {
		C.QSerialPort_FlowControlChanged(ptr.Pointer(), C.longlong(flow))
	}
}

func (ptr *QSerialPort) Flush() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSerialPort_Flush(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSerialPort) IsBreakEnabled() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSerialPort_IsBreakEnabled(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSerialPort) IsDataTerminalReady() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSerialPort_IsDataTerminalReady(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSerialPort) IsRequestToSend() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSerialPort_IsRequestToSend(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQSerialPort_IsSequential
func callbackQSerialPort_IsSequential(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isSequential"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSerialPortFromPointer(ptr).IsSequentialDefault())))
}

func (ptr *QSerialPort) IsSequentialDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSerialPort_IsSequentialDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQSerialPort_Open
func callbackQSerialPort_Open(ptr unsafe.Pointer, mode C.longlong) C.char {
	if signal := qt.GetSignal(ptr, "open"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(core.QIODevice__OpenModeFlag) bool)(signal))(core.QIODevice__OpenModeFlag(mode)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSerialPortFromPointer(ptr).OpenDefault(core.QIODevice__OpenModeFlag(mode)))))
}

func (ptr *QSerialPort) OpenDefault(mode core.QIODevice__OpenModeFlag) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSerialPort_OpenDefault(ptr.Pointer(), C.longlong(mode))) != 0
	}
	return false
}

func (ptr *QSerialPort) Parity() QSerialPort__Parity {
	if ptr.Pointer() != nil {
		return QSerialPort__Parity(C.QSerialPort_Parity(ptr.Pointer()))
	}
	return 0
}

//export callbackQSerialPort_ParityChanged
func callbackQSerialPort_ParityChanged(ptr unsafe.Pointer, parity C.longlong) {
	if signal := qt.GetSignal(ptr, "parityChanged"); signal != nil {
		(*(*func(QSerialPort__Parity))(signal))(QSerialPort__Parity(parity))
	}

}

func (ptr *QSerialPort) ConnectParityChanged(f func(parity QSerialPort__Parity)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "parityChanged") {
			C.QSerialPort_ConnectParityChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "parityChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "parityChanged"); signal != nil {
			f := func(parity QSerialPort__Parity) {
				(*(*func(QSerialPort__Parity))(signal))(parity)
				f(parity)
			}
			qt.ConnectSignal(ptr.Pointer(), "parityChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "parityChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSerialPort) DisconnectParityChanged() {
	if ptr.Pointer() != nil {
		C.QSerialPort_DisconnectParityChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "parityChanged")
	}
}

func (ptr *QSerialPort) ParityChanged(parity QSerialPort__Parity) {
	if ptr.Pointer() != nil {
		C.QSerialPort_ParityChanged(ptr.Pointer(), C.longlong(parity))
	}
}

func (ptr *QSerialPort) PinoutSignals() QSerialPort__PinoutSignal {
	if ptr.Pointer() != nil {
		return QSerialPort__PinoutSignal(C.QSerialPort_PinoutSignals(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSerialPort) PortName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSerialPort_PortName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSerialPort) ReadBufferSize() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QSerialPort_ReadBufferSize(ptr.Pointer()))
	}
	return 0
}

//export callbackQSerialPort_ReadData
func callbackQSerialPort_ReadData(ptr unsafe.Pointer, data C.struct_QtSerialPort_PackedString, maxSize C.longlong) C.longlong {
	if signal := qt.GetSignal(ptr, "readData"); signal != nil {
		retS := cGoUnpackString(data)
		ret := C.longlong((*(*func(*string, int64) int64)(signal))(&retS, int64(maxSize)))
		if ret > 0 {
			C.memcpy(unsafe.Pointer(data.data), unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&retS)).Data), C.size_t(ret))
		}
		return ret
	}
	retS := cGoUnpackString(data)
	ret := C.longlong(NewQSerialPortFromPointer(ptr).ReadDataDefault(&retS, int64(maxSize)))
	if ret > 0 {
		C.memcpy(unsafe.Pointer(data.data), unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&retS)).Data), C.size_t(ret))
	}
	return ret
}

func (ptr *QSerialPort) ConnectReadData(f func(data *string, maxSize int64) int64) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "readData"); signal != nil {
			f := func(data *string, maxSize int64) int64 {
				(*(*func(*string, int64) int64)(signal))(data, maxSize)
				return f(data, maxSize)
			}
			qt.ConnectSignal(ptr.Pointer(), "readData", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "readData", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSerialPort) DisconnectReadData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "readData")
	}
}

func (ptr *QSerialPort) ReadData(data *string, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		dataC := C.CString(strings.Repeat("0", int(maxSize)))
		defer C.free(unsafe.Pointer(dataC))
		ret := int64(C.QSerialPort_ReadData(ptr.Pointer(), dataC, C.longlong(maxSize)))
		if ret > 0 {
			*data = C.GoStringN(dataC, C.int(ret))
		}
		return ret
	}
	return 0
}

func (ptr *QSerialPort) ReadDataDefault(data *string, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		dataC := C.CString(strings.Repeat("0", int(maxSize)))
		defer C.free(unsafe.Pointer(dataC))
		ret := int64(C.QSerialPort_ReadDataDefault(ptr.Pointer(), dataC, C.longlong(maxSize)))
		if ret > 0 {
			*data = C.GoStringN(dataC, C.int(ret))
		}
		return ret
	}
	return 0
}

//export callbackQSerialPort_ReadLineData
func callbackQSerialPort_ReadLineData(ptr unsafe.Pointer, data C.struct_QtSerialPort_PackedString, maxSize C.longlong) C.longlong {
	if signal := qt.GetSignal(ptr, "readLineData"); signal != nil {
		return C.longlong((*(*func([]byte, int64) int64)(signal))(cGoUnpackBytes(data), int64(maxSize)))
	}

	return C.longlong(NewQSerialPortFromPointer(ptr).ReadLineDataDefault(cGoUnpackBytes(data), int64(maxSize)))
}

func (ptr *QSerialPort) ReadLineDataDefault(data []byte, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if len(data) != 0 {
			dataC = (*C.char)(unsafe.Pointer(&data[0]))
		}
		return int64(C.QSerialPort_ReadLineDataDefault(ptr.Pointer(), dataC, C.longlong(maxSize)))
	}
	return 0
}

//export callbackQSerialPort_RequestToSendChanged
func callbackQSerialPort_RequestToSendChanged(ptr unsafe.Pointer, set C.char) {
	if signal := qt.GetSignal(ptr, "requestToSendChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(set) != 0)
	}

}

func (ptr *QSerialPort) ConnectRequestToSendChanged(f func(set bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "requestToSendChanged") {
			C.QSerialPort_ConnectRequestToSendChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "requestToSendChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "requestToSendChanged"); signal != nil {
			f := func(set bool) {
				(*(*func(bool))(signal))(set)
				f(set)
			}
			qt.ConnectSignal(ptr.Pointer(), "requestToSendChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "requestToSendChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSerialPort) DisconnectRequestToSendChanged() {
	if ptr.Pointer() != nil {
		C.QSerialPort_DisconnectRequestToSendChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "requestToSendChanged")
	}
}

func (ptr *QSerialPort) RequestToSendChanged(set bool) {
	if ptr.Pointer() != nil {
		C.QSerialPort_RequestToSendChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(set))))
	}
}

func (ptr *QSerialPort) SendBreak(duration int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSerialPort_SendBreak(ptr.Pointer(), C.int(int32(duration)))) != 0
	}
	return false
}

func (ptr *QSerialPort) SetBaudRate(baudRate int, directions QSerialPort__Direction) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSerialPort_SetBaudRate(ptr.Pointer(), C.int(int32(baudRate)), C.longlong(directions))) != 0
	}
	return false
}

func (ptr *QSerialPort) SetBreakEnabled(set bool) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSerialPort_SetBreakEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(set))))) != 0
	}
	return false
}

func (ptr *QSerialPort) SetDataBits(dataBits QSerialPort__DataBits) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSerialPort_SetDataBits(ptr.Pointer(), C.longlong(dataBits))) != 0
	}
	return false
}

func (ptr *QSerialPort) SetDataTerminalReady(set bool) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSerialPort_SetDataTerminalReady(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(set))))) != 0
	}
	return false
}

func (ptr *QSerialPort) SetFlowControl(flowControl QSerialPort__FlowControl) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSerialPort_SetFlowControl(ptr.Pointer(), C.longlong(flowControl))) != 0
	}
	return false
}

func (ptr *QSerialPort) SetParity(parity QSerialPort__Parity) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSerialPort_SetParity(ptr.Pointer(), C.longlong(parity))) != 0
	}
	return false
}

func (ptr *QSerialPort) SetPort(serialPortInfo QSerialPortInfo_ITF) {
	if ptr.Pointer() != nil {
		C.QSerialPort_SetPort(ptr.Pointer(), PointerFromQSerialPortInfo(serialPortInfo))
	}
}

func (ptr *QSerialPort) SetPortName(name string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QSerialPort_SetPortName(ptr.Pointer(), C.struct_QtSerialPort_PackedString{data: nameC, len: C.longlong(len(name))})
	}
}

func (ptr *QSerialPort) SetReadBufferSize(size int64) {
	if ptr.Pointer() != nil {
		C.QSerialPort_SetReadBufferSize(ptr.Pointer(), C.longlong(size))
	}
}

func (ptr *QSerialPort) SetRequestToSend(set bool) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSerialPort_SetRequestToSend(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(set))))) != 0
	}
	return false
}

func (ptr *QSerialPort) SetSettingsRestoredOnClose(restore bool) {
	if ptr.Pointer() != nil {
		C.QSerialPort_SetSettingsRestoredOnClose(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(restore))))
	}
}

func (ptr *QSerialPort) SetStopBits(stopBits QSerialPort__StopBits) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSerialPort_SetStopBits(ptr.Pointer(), C.longlong(stopBits))) != 0
	}
	return false
}

func (ptr *QSerialPort) SettingsRestoredOnClose() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSerialPort_SettingsRestoredOnClose(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSerialPort) StopBits() QSerialPort__StopBits {
	if ptr.Pointer() != nil {
		return QSerialPort__StopBits(C.QSerialPort_StopBits(ptr.Pointer()))
	}
	return 0
}

//export callbackQSerialPort_StopBitsChanged
func callbackQSerialPort_StopBitsChanged(ptr unsafe.Pointer, stopBits C.longlong) {
	if signal := qt.GetSignal(ptr, "stopBitsChanged"); signal != nil {
		(*(*func(QSerialPort__StopBits))(signal))(QSerialPort__StopBits(stopBits))
	}

}

func (ptr *QSerialPort) ConnectStopBitsChanged(f func(stopBits QSerialPort__StopBits)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "stopBitsChanged") {
			C.QSerialPort_ConnectStopBitsChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "stopBitsChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "stopBitsChanged"); signal != nil {
			f := func(stopBits QSerialPort__StopBits) {
				(*(*func(QSerialPort__StopBits))(signal))(stopBits)
				f(stopBits)
			}
			qt.ConnectSignal(ptr.Pointer(), "stopBitsChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stopBitsChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSerialPort) DisconnectStopBitsChanged() {
	if ptr.Pointer() != nil {
		C.QSerialPort_DisconnectStopBitsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "stopBitsChanged")
	}
}

func (ptr *QSerialPort) StopBitsChanged(stopBits QSerialPort__StopBits) {
	if ptr.Pointer() != nil {
		C.QSerialPort_StopBitsChanged(ptr.Pointer(), C.longlong(stopBits))
	}
}

//export callbackQSerialPort_WaitForBytesWritten
func callbackQSerialPort_WaitForBytesWritten(ptr unsafe.Pointer, msecs C.int) C.char {
	if signal := qt.GetSignal(ptr, "waitForBytesWritten"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int) bool)(signal))(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSerialPortFromPointer(ptr).WaitForBytesWrittenDefault(int(int32(msecs))))))
}

func (ptr *QSerialPort) WaitForBytesWrittenDefault(msecs int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSerialPort_WaitForBytesWrittenDefault(ptr.Pointer(), C.int(int32(msecs)))) != 0
	}
	return false
}

//export callbackQSerialPort_WaitForReadyRead
func callbackQSerialPort_WaitForReadyRead(ptr unsafe.Pointer, msecs C.int) C.char {
	if signal := qt.GetSignal(ptr, "waitForReadyRead"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int) bool)(signal))(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSerialPortFromPointer(ptr).WaitForReadyReadDefault(int(int32(msecs))))))
}

func (ptr *QSerialPort) WaitForReadyReadDefault(msecs int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSerialPort_WaitForReadyReadDefault(ptr.Pointer(), C.int(int32(msecs)))) != 0
	}
	return false
}

//export callbackQSerialPort_WriteData
func callbackQSerialPort_WriteData(ptr unsafe.Pointer, data C.struct_QtSerialPort_PackedString, maxSize C.longlong) C.longlong {
	if signal := qt.GetSignal(ptr, "writeData"); signal != nil {
		return C.longlong((*(*func([]byte, int64) int64)(signal))(cGoUnpackBytes(data), int64(maxSize)))
	}

	return C.longlong(NewQSerialPortFromPointer(ptr).WriteDataDefault(cGoUnpackBytes(data), int64(maxSize)))
}

func (ptr *QSerialPort) ConnectWriteData(f func(data []byte, maxSize int64) int64) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "writeData"); signal != nil {
			f := func(data []byte, maxSize int64) int64 {
				(*(*func([]byte, int64) int64)(signal))(data, maxSize)
				return f(data, maxSize)
			}
			qt.ConnectSignal(ptr.Pointer(), "writeData", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "writeData", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSerialPort) DisconnectWriteData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "writeData")
	}
}

func (ptr *QSerialPort) WriteData(data []byte, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if len(data) != 0 {
			dataC = (*C.char)(unsafe.Pointer(&data[0]))
		}
		return int64(C.QSerialPort_WriteData(ptr.Pointer(), dataC, C.longlong(maxSize)))
	}
	return 0
}

func (ptr *QSerialPort) WriteDataDefault(data []byte, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if len(data) != 0 {
			dataC = (*C.char)(unsafe.Pointer(&data[0]))
		}
		return int64(C.QSerialPort_WriteDataDefault(ptr.Pointer(), dataC, C.longlong(maxSize)))
	}
	return 0
}

//export callbackQSerialPort_DestroyQSerialPort
func callbackQSerialPort_DestroyQSerialPort(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QSerialPort"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQSerialPortFromPointer(ptr).DestroyQSerialPortDefault()
	}
}

func (ptr *QSerialPort) ConnectDestroyQSerialPort(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QSerialPort"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QSerialPort", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QSerialPort", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSerialPort) DisconnectDestroyQSerialPort() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QSerialPort")
	}
}

func (ptr *QSerialPort) DestroyQSerialPort() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QSerialPort_DestroyQSerialPort(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSerialPort) DestroyQSerialPortDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QSerialPort_DestroyQSerialPortDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSerialPort) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSerialPort___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSerialPort) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSerialPort___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSerialPort) __children_newList() unsafe.Pointer {
	return C.QSerialPort___children_newList(ptr.Pointer())
}

func (ptr *QSerialPort) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QSerialPort___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSerialPort) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QSerialPort___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QSerialPort) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QSerialPort___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QSerialPort) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSerialPort___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSerialPort) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSerialPort___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSerialPort) __findChildren_newList() unsafe.Pointer {
	return C.QSerialPort___findChildren_newList(ptr.Pointer())
}

func (ptr *QSerialPort) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSerialPort___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSerialPort) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSerialPort___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSerialPort) __findChildren_newList3() unsafe.Pointer {
	return C.QSerialPort___findChildren_newList3(ptr.Pointer())
}

//export callbackQSerialPort_AboutToClose
func callbackQSerialPort_AboutToClose(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "aboutToClose"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackQSerialPort_BytesWritten
func callbackQSerialPort_BytesWritten(ptr unsafe.Pointer, bytes C.longlong) {
	if signal := qt.GetSignal(ptr, "bytesWritten"); signal != nil {
		(*(*func(int64))(signal))(int64(bytes))
	}

}

//export callbackQSerialPort_ChannelBytesWritten
func callbackQSerialPort_ChannelBytesWritten(ptr unsafe.Pointer, channel C.int, bytes C.longlong) {
	if signal := qt.GetSignal(ptr, "channelBytesWritten"); signal != nil {
		(*(*func(int, int64))(signal))(int(int32(channel)), int64(bytes))
	}

}

//export callbackQSerialPort_ChannelReadyRead
func callbackQSerialPort_ChannelReadyRead(ptr unsafe.Pointer, channel C.int) {
	if signal := qt.GetSignal(ptr, "channelReadyRead"); signal != nil {
		(*(*func(int))(signal))(int(int32(channel)))
	}

}

//export callbackQSerialPort_Pos
func callbackQSerialPort_Pos(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "pos"); signal != nil {
		return C.longlong((*(*func() int64)(signal))())
	}

	return C.longlong(NewQSerialPortFromPointer(ptr).PosDefault())
}

func (ptr *QSerialPort) PosDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QSerialPort_PosDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQSerialPort_ReadChannelFinished
func callbackQSerialPort_ReadChannelFinished(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "readChannelFinished"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackQSerialPort_ReadyRead
func callbackQSerialPort_ReadyRead(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "readyRead"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackQSerialPort_Reset
func callbackQSerialPort_Reset(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "reset"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSerialPortFromPointer(ptr).ResetDefault())))
}

func (ptr *QSerialPort) ResetDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSerialPort_ResetDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQSerialPort_Seek
func callbackQSerialPort_Seek(ptr unsafe.Pointer, pos C.longlong) C.char {
	if signal := qt.GetSignal(ptr, "seek"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int64) bool)(signal))(int64(pos)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSerialPortFromPointer(ptr).SeekDefault(int64(pos)))))
}

func (ptr *QSerialPort) SeekDefault(pos int64) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSerialPort_SeekDefault(ptr.Pointer(), C.longlong(pos))) != 0
	}
	return false
}

//export callbackQSerialPort_Size
func callbackQSerialPort_Size(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "size"); signal != nil {
		return C.longlong((*(*func() int64)(signal))())
	}

	return C.longlong(NewQSerialPortFromPointer(ptr).SizeDefault())
}

func (ptr *QSerialPort) SizeDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QSerialPort_SizeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQSerialPort_ChildEvent
func callbackQSerialPort_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSerialPortFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSerialPort) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSerialPort_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQSerialPort_ConnectNotify
func callbackQSerialPort_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSerialPortFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSerialPort) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSerialPort_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSerialPort_CustomEvent
func callbackQSerialPort_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQSerialPortFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSerialPort) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSerialPort_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQSerialPort_DeleteLater
func callbackQSerialPort_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQSerialPortFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QSerialPort) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QSerialPort_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQSerialPort_Destroyed
func callbackQSerialPort_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQSerialPort_DisconnectNotify
func callbackQSerialPort_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSerialPortFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSerialPort) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSerialPort_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSerialPort_Event
func callbackQSerialPort_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSerialPortFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QSerialPort) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSerialPort_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQSerialPort_EventFilter
func callbackQSerialPort_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSerialPortFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QSerialPort) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSerialPort_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQSerialPort_MetaObject
func callbackQSerialPort_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQSerialPortFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QSerialPort) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSerialPort_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQSerialPort_ObjectNameChanged
func callbackQSerialPort_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtSerialPort_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQSerialPort_TimerEvent
func callbackQSerialPort_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSerialPortFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSerialPort) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSerialPort_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QSerialPortInfo struct {
	ptr unsafe.Pointer
}

type QSerialPortInfo_ITF interface {
	QSerialPortInfo_PTR() *QSerialPortInfo
}

func (ptr *QSerialPortInfo) QSerialPortInfo_PTR() *QSerialPortInfo {
	return ptr
}

func (ptr *QSerialPortInfo) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QSerialPortInfo) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQSerialPortInfo(ptr QSerialPortInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSerialPortInfo_PTR().Pointer()
	}
	return nil
}

func NewQSerialPortInfoFromPointer(ptr unsafe.Pointer) (n *QSerialPortInfo) {
	n = new(QSerialPortInfo)
	n.SetPointer(ptr)
	return
}
func NewQSerialPortInfo() *QSerialPortInfo {
	tmpValue := NewQSerialPortInfoFromPointer(C.QSerialPortInfo_NewQSerialPortInfo())
	qt.SetFinalizer(tmpValue, (*QSerialPortInfo).DestroyQSerialPortInfo)
	return tmpValue
}

func NewQSerialPortInfo2(port QSerialPort_ITF) *QSerialPortInfo {
	tmpValue := NewQSerialPortInfoFromPointer(C.QSerialPortInfo_NewQSerialPortInfo2(PointerFromQSerialPort(port)))
	qt.SetFinalizer(tmpValue, (*QSerialPortInfo).DestroyQSerialPortInfo)
	return tmpValue
}

func NewQSerialPortInfo3(name string) *QSerialPortInfo {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := NewQSerialPortInfoFromPointer(C.QSerialPortInfo_NewQSerialPortInfo3(C.struct_QtSerialPort_PackedString{data: nameC, len: C.longlong(len(name))}))
	qt.SetFinalizer(tmpValue, (*QSerialPortInfo).DestroyQSerialPortInfo)
	return tmpValue
}

func NewQSerialPortInfo4(other QSerialPortInfo_ITF) *QSerialPortInfo {
	tmpValue := NewQSerialPortInfoFromPointer(C.QSerialPortInfo_NewQSerialPortInfo4(PointerFromQSerialPortInfo(other)))
	qt.SetFinalizer(tmpValue, (*QSerialPortInfo).DestroyQSerialPortInfo)
	return tmpValue
}

func QSerialPortInfo_AvailablePorts() []*QSerialPortInfo {
	return func(l C.struct_QtSerialPort_PackedList) []*QSerialPortInfo {
		out := make([]*QSerialPortInfo, int(l.len))
		tmpList := NewQSerialPortInfoFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__availablePorts_atList(i)
		}
		return out
	}(C.QSerialPortInfo_QSerialPortInfo_AvailablePorts())
}

func (ptr *QSerialPortInfo) AvailablePorts() []*QSerialPortInfo {
	return func(l C.struct_QtSerialPort_PackedList) []*QSerialPortInfo {
		out := make([]*QSerialPortInfo, int(l.len))
		tmpList := NewQSerialPortInfoFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__availablePorts_atList(i)
		}
		return out
	}(C.QSerialPortInfo_QSerialPortInfo_AvailablePorts())
}

func (ptr *QSerialPortInfo) Description() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSerialPortInfo_Description(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSerialPortInfo) HasProductIdentifier() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSerialPortInfo_HasProductIdentifier(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSerialPortInfo) HasVendorIdentifier() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSerialPortInfo_HasVendorIdentifier(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSerialPortInfo) IsBusy() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSerialPortInfo_IsBusy(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSerialPortInfo) IsNull() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSerialPortInfo_IsNull(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSerialPortInfo) Manufacturer() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSerialPortInfo_Manufacturer(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSerialPortInfo) PortName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSerialPortInfo_PortName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSerialPortInfo) ProductIdentifier() uint16 {
	if ptr.Pointer() != nil {
		return uint16(C.QSerialPortInfo_ProductIdentifier(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSerialPortInfo) SerialNumber() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSerialPortInfo_SerialNumber(ptr.Pointer()))
	}
	return ""
}

func QSerialPortInfo_StandardBaudRates() []int {
	return func(l C.struct_QtSerialPort_PackedList) []int {
		out := make([]int, int(l.len))
		tmpList := NewQSerialPortInfoFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__standardBaudRates_atList(i)
		}
		return out
	}(C.QSerialPortInfo_QSerialPortInfo_StandardBaudRates())
}

func (ptr *QSerialPortInfo) StandardBaudRates() []int {
	return func(l C.struct_QtSerialPort_PackedList) []int {
		out := make([]int, int(l.len))
		tmpList := NewQSerialPortInfoFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__standardBaudRates_atList(i)
		}
		return out
	}(C.QSerialPortInfo_QSerialPortInfo_StandardBaudRates())
}

func (ptr *QSerialPortInfo) Swap(other QSerialPortInfo_ITF) {
	if ptr.Pointer() != nil {
		C.QSerialPortInfo_Swap(ptr.Pointer(), PointerFromQSerialPortInfo(other))
	}
}

func (ptr *QSerialPortInfo) SystemLocation() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSerialPortInfo_SystemLocation(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSerialPortInfo) VendorIdentifier() uint16 {
	if ptr.Pointer() != nil {
		return uint16(C.QSerialPortInfo_VendorIdentifier(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSerialPortInfo) DestroyQSerialPortInfo() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QSerialPortInfo_DestroyQSerialPortInfo(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSerialPortInfo) __availablePorts_atList(i int) *QSerialPortInfo {
	if ptr.Pointer() != nil {
		tmpValue := NewQSerialPortInfoFromPointer(C.QSerialPortInfo___availablePorts_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QSerialPortInfo).DestroyQSerialPortInfo)
		return tmpValue
	}
	return nil
}

func (ptr *QSerialPortInfo) __availablePorts_setList(i QSerialPortInfo_ITF) {
	if ptr.Pointer() != nil {
		C.QSerialPortInfo___availablePorts_setList(ptr.Pointer(), PointerFromQSerialPortInfo(i))
	}
}

func (ptr *QSerialPortInfo) __availablePorts_newList() unsafe.Pointer {
	return C.QSerialPortInfo___availablePorts_newList(ptr.Pointer())
}

func (ptr *QSerialPortInfo) __standardBaudRates_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSerialPortInfo___standardBaudRates_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QSerialPortInfo) __standardBaudRates_setList(i int) {
	if ptr.Pointer() != nil {
		C.QSerialPortInfo___standardBaudRates_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QSerialPortInfo) __standardBaudRates_newList() unsafe.Pointer {
	return C.QSerialPortInfo___standardBaudRates_newList(ptr.Pointer())
}

func init() {
	qt.ItfMap["serialport.QSerialPort_ITF"] = QSerialPort{}
	qt.FuncMap["serialport.NewQSerialPort"] = NewQSerialPort
	qt.FuncMap["serialport.NewQSerialPort2"] = NewQSerialPort2
	qt.FuncMap["serialport.NewQSerialPort3"] = NewQSerialPort3
	qt.EnumMap["serialport.QSerialPort__Input"] = int64(QSerialPort__Input)
	qt.EnumMap["serialport.QSerialPort__Output"] = int64(QSerialPort__Output)
	qt.EnumMap["serialport.QSerialPort__AllDirections"] = int64(QSerialPort__AllDirections)
	qt.EnumMap["serialport.QSerialPort__Baud1200"] = int64(QSerialPort__Baud1200)
	qt.EnumMap["serialport.QSerialPort__Baud2400"] = int64(QSerialPort__Baud2400)
	qt.EnumMap["serialport.QSerialPort__Baud4800"] = int64(QSerialPort__Baud4800)
	qt.EnumMap["serialport.QSerialPort__Baud9600"] = int64(QSerialPort__Baud9600)
	qt.EnumMap["serialport.QSerialPort__Baud19200"] = int64(QSerialPort__Baud19200)
	qt.EnumMap["serialport.QSerialPort__Baud38400"] = int64(QSerialPort__Baud38400)
	qt.EnumMap["serialport.QSerialPort__Baud57600"] = int64(QSerialPort__Baud57600)
	qt.EnumMap["serialport.QSerialPort__Baud115200"] = int64(QSerialPort__Baud115200)
	qt.EnumMap["serialport.QSerialPort__UnknownBaud"] = int64(QSerialPort__UnknownBaud)
	qt.EnumMap["serialport.QSerialPort__Data5"] = int64(QSerialPort__Data5)
	qt.EnumMap["serialport.QSerialPort__Data6"] = int64(QSerialPort__Data6)
	qt.EnumMap["serialport.QSerialPort__Data7"] = int64(QSerialPort__Data7)
	qt.EnumMap["serialport.QSerialPort__Data8"] = int64(QSerialPort__Data8)
	qt.EnumMap["serialport.QSerialPort__UnknownDataBits"] = int64(QSerialPort__UnknownDataBits)
	qt.EnumMap["serialport.QSerialPort__NoParity"] = int64(QSerialPort__NoParity)
	qt.EnumMap["serialport.QSerialPort__EvenParity"] = int64(QSerialPort__EvenParity)
	qt.EnumMap["serialport.QSerialPort__OddParity"] = int64(QSerialPort__OddParity)
	qt.EnumMap["serialport.QSerialPort__SpaceParity"] = int64(QSerialPort__SpaceParity)
	qt.EnumMap["serialport.QSerialPort__MarkParity"] = int64(QSerialPort__MarkParity)
	qt.EnumMap["serialport.QSerialPort__UnknownParity"] = int64(QSerialPort__UnknownParity)
	qt.EnumMap["serialport.QSerialPort__OneStop"] = int64(QSerialPort__OneStop)
	qt.EnumMap["serialport.QSerialPort__OneAndHalfStop"] = int64(QSerialPort__OneAndHalfStop)
	qt.EnumMap["serialport.QSerialPort__TwoStop"] = int64(QSerialPort__TwoStop)
	qt.EnumMap["serialport.QSerialPort__UnknownStopBits"] = int64(QSerialPort__UnknownStopBits)
	qt.EnumMap["serialport.QSerialPort__NoFlowControl"] = int64(QSerialPort__NoFlowControl)
	qt.EnumMap["serialport.QSerialPort__HardwareControl"] = int64(QSerialPort__HardwareControl)
	qt.EnumMap["serialport.QSerialPort__SoftwareControl"] = int64(QSerialPort__SoftwareControl)
	qt.EnumMap["serialport.QSerialPort__UnknownFlowControl"] = int64(QSerialPort__UnknownFlowControl)
	qt.EnumMap["serialport.QSerialPort__NoSignal"] = int64(QSerialPort__NoSignal)
	qt.EnumMap["serialport.QSerialPort__TransmittedDataSignal"] = int64(QSerialPort__TransmittedDataSignal)
	qt.EnumMap["serialport.QSerialPort__ReceivedDataSignal"] = int64(QSerialPort__ReceivedDataSignal)
	qt.EnumMap["serialport.QSerialPort__DataTerminalReadySignal"] = int64(QSerialPort__DataTerminalReadySignal)
	qt.EnumMap["serialport.QSerialPort__DataCarrierDetectSignal"] = int64(QSerialPort__DataCarrierDetectSignal)
	qt.EnumMap["serialport.QSerialPort__DataSetReadySignal"] = int64(QSerialPort__DataSetReadySignal)
	qt.EnumMap["serialport.QSerialPort__RingIndicatorSignal"] = int64(QSerialPort__RingIndicatorSignal)
	qt.EnumMap["serialport.QSerialPort__RequestToSendSignal"] = int64(QSerialPort__RequestToSendSignal)
	qt.EnumMap["serialport.QSerialPort__ClearToSendSignal"] = int64(QSerialPort__ClearToSendSignal)
	qt.EnumMap["serialport.QSerialPort__SecondaryTransmittedDataSignal"] = int64(QSerialPort__SecondaryTransmittedDataSignal)
	qt.EnumMap["serialport.QSerialPort__SecondaryReceivedDataSignal"] = int64(QSerialPort__SecondaryReceivedDataSignal)
	qt.EnumMap["serialport.QSerialPort__NoError"] = int64(QSerialPort__NoError)
	qt.EnumMap["serialport.QSerialPort__DeviceNotFoundError"] = int64(QSerialPort__DeviceNotFoundError)
	qt.EnumMap["serialport.QSerialPort__PermissionError"] = int64(QSerialPort__PermissionError)
	qt.EnumMap["serialport.QSerialPort__OpenError"] = int64(QSerialPort__OpenError)
	qt.EnumMap["serialport.QSerialPort__ParityError"] = int64(QSerialPort__ParityError)
	qt.EnumMap["serialport.QSerialPort__FramingError"] = int64(QSerialPort__FramingError)
	qt.EnumMap["serialport.QSerialPort__BreakConditionError"] = int64(QSerialPort__BreakConditionError)
	qt.EnumMap["serialport.QSerialPort__WriteError"] = int64(QSerialPort__WriteError)
	qt.EnumMap["serialport.QSerialPort__ReadError"] = int64(QSerialPort__ReadError)
	qt.EnumMap["serialport.QSerialPort__ResourceError"] = int64(QSerialPort__ResourceError)
	qt.EnumMap["serialport.QSerialPort__UnsupportedOperationError"] = int64(QSerialPort__UnsupportedOperationError)
	qt.EnumMap["serialport.QSerialPort__UnknownError"] = int64(QSerialPort__UnknownError)
	qt.EnumMap["serialport.QSerialPort__TimeoutError"] = int64(QSerialPort__TimeoutError)
	qt.EnumMap["serialport.QSerialPort__NotOpenError"] = int64(QSerialPort__NotOpenError)
	qt.ItfMap["serialport.QSerialPortInfo_ITF"] = QSerialPortInfo{}
	qt.FuncMap["serialport.NewQSerialPortInfo"] = NewQSerialPortInfo
	qt.FuncMap["serialport.NewQSerialPortInfo2"] = NewQSerialPortInfo2
	qt.FuncMap["serialport.NewQSerialPortInfo3"] = NewQSerialPortInfo3
	qt.FuncMap["serialport.NewQSerialPortInfo4"] = NewQSerialPortInfo4
	qt.FuncMap["serialport.QSerialPortInfo_AvailablePorts"] = QSerialPortInfo_AvailablePorts
	qt.FuncMap["serialport.QSerialPortInfo_StandardBaudRates"] = QSerialPortInfo_StandardBaudRates
}
