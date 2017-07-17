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
	"runtime"
	"strings"
	"unsafe"
)

func cGoUnpackString(s C.struct_QtSerialPort_PackedString) string {
	if len := int(s.len); len == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
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

func NewQSerialPortFromPointer(ptr unsafe.Pointer) *QSerialPort {
	var n = new(QSerialPort)
	n.SetPointer(ptr)
	return n
}

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

//go:generate stringer -type=QSerialPort__Direction
//QSerialPort::Direction
type QSerialPort__Direction int64

const (
	QSerialPort__Input         QSerialPort__Direction = QSerialPort__Direction(1)
	QSerialPort__Output        QSerialPort__Direction = QSerialPort__Direction(2)
	QSerialPort__AllDirections QSerialPort__Direction = QSerialPort__Direction(QSerialPort__Input | QSerialPort__Output)
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

//go:generate stringer -type=QSerialPort__StopBits
//QSerialPort::StopBits
type QSerialPort__StopBits int64

const (
	QSerialPort__OneStop         QSerialPort__StopBits = QSerialPort__StopBits(1)
	QSerialPort__OneAndHalfStop  QSerialPort__StopBits = QSerialPort__StopBits(3)
	QSerialPort__TwoStop         QSerialPort__StopBits = QSerialPort__StopBits(2)
	QSerialPort__UnknownStopBits QSerialPort__StopBits = QSerialPort__StopBits(-1)
)

func (ptr *QSerialPort) PinoutSignals() QSerialPort__PinoutSignal {
	if ptr.Pointer() != nil {
		return QSerialPort__PinoutSignal(C.QSerialPort_PinoutSignals(ptr.Pointer()))
	}
	return 0
}

func NewQSerialPort(parent core.QObject_ITF) *QSerialPort {
	var tmpValue = NewQSerialPortFromPointer(C.QSerialPort_NewQSerialPort(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQSerialPort3(serialPortInfo QSerialPortInfo_ITF, parent core.QObject_ITF) *QSerialPort {
	var tmpValue = NewQSerialPortFromPointer(C.QSerialPort_NewQSerialPort3(PointerFromQSerialPortInfo(serialPortInfo), core.PointerFromQObject(parent)))
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
	var tmpValue = NewQSerialPortFromPointer(C.QSerialPort_NewQSerialPort2(C.struct_QtSerialPort_PackedString{data: nameC, len: C.longlong(len(name))}, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QSerialPort) Clear(directions QSerialPort__Direction) bool {
	if ptr.Pointer() != nil {
		return C.QSerialPort_Clear(ptr.Pointer(), C.longlong(directions)) != 0
	}
	return false
}

func (ptr *QSerialPort) Flush() bool {
	if ptr.Pointer() != nil {
		return C.QSerialPort_Flush(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSerialPort) IsDataTerminalReady() bool {
	if ptr.Pointer() != nil {
		return C.QSerialPort_IsDataTerminalReady(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSerialPort) IsRequestToSend() bool {
	if ptr.Pointer() != nil {
		return C.QSerialPort_IsRequestToSend(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSerialPort_Open
func callbackQSerialPort_Open(ptr unsafe.Pointer, mode C.longlong) C.char {
	if signal := qt.GetSignal(ptr, "open"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(core.QIODevice__OpenModeFlag) bool)(core.QIODevice__OpenModeFlag(mode)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSerialPortFromPointer(ptr).OpenDefault(core.QIODevice__OpenModeFlag(mode)))))
}

func (ptr *QSerialPort) OpenDefault(mode core.QIODevice__OpenModeFlag) bool {
	if ptr.Pointer() != nil {
		return C.QSerialPort_OpenDefault(ptr.Pointer(), C.longlong(mode)) != 0
	}
	return false
}

func (ptr *QSerialPort) SetBaudRate(baudRate int, directions QSerialPort__Direction) bool {
	if ptr.Pointer() != nil {
		return C.QSerialPort_SetBaudRate(ptr.Pointer(), C.int(int32(baudRate)), C.longlong(directions)) != 0
	}
	return false
}

func (ptr *QSerialPort) SetBreakEnabled(set bool) bool {
	if ptr.Pointer() != nil {
		return C.QSerialPort_SetBreakEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(set)))) != 0
	}
	return false
}

func (ptr *QSerialPort) SetDataBits(dataBits QSerialPort__DataBits) bool {
	if ptr.Pointer() != nil {
		return C.QSerialPort_SetDataBits(ptr.Pointer(), C.longlong(dataBits)) != 0
	}
	return false
}

func (ptr *QSerialPort) SetDataTerminalReady(set bool) bool {
	if ptr.Pointer() != nil {
		return C.QSerialPort_SetDataTerminalReady(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(set)))) != 0
	}
	return false
}

func (ptr *QSerialPort) SetFlowControl(flowControl QSerialPort__FlowControl) bool {
	if ptr.Pointer() != nil {
		return C.QSerialPort_SetFlowControl(ptr.Pointer(), C.longlong(flowControl)) != 0
	}
	return false
}

func (ptr *QSerialPort) SetParity(parity QSerialPort__Parity) bool {
	if ptr.Pointer() != nil {
		return C.QSerialPort_SetParity(ptr.Pointer(), C.longlong(parity)) != 0
	}
	return false
}

func (ptr *QSerialPort) SetRequestToSend(set bool) bool {
	if ptr.Pointer() != nil {
		return C.QSerialPort_SetRequestToSend(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(set)))) != 0
	}
	return false
}

func (ptr *QSerialPort) SetStopBits(stopBits QSerialPort__StopBits) bool {
	if ptr.Pointer() != nil {
		return C.QSerialPort_SetStopBits(ptr.Pointer(), C.longlong(stopBits)) != 0
	}
	return false
}

//export callbackQSerialPort_WaitForBytesWritten
func callbackQSerialPort_WaitForBytesWritten(ptr unsafe.Pointer, msecs C.int) C.char {
	if signal := qt.GetSignal(ptr, "waitForBytesWritten"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSerialPortFromPointer(ptr).WaitForBytesWrittenDefault(int(int32(msecs))))))
}

func (ptr *QSerialPort) WaitForBytesWrittenDefault(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QSerialPort_WaitForBytesWrittenDefault(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

//export callbackQSerialPort_WaitForReadyRead
func callbackQSerialPort_WaitForReadyRead(ptr unsafe.Pointer, msecs C.int) C.char {
	if signal := qt.GetSignal(ptr, "waitForReadyRead"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSerialPortFromPointer(ptr).WaitForReadyReadDefault(int(int32(msecs))))))
}

func (ptr *QSerialPort) WaitForReadyReadDefault(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QSerialPort_WaitForReadyReadDefault(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

//export callbackQSerialPort_ReadData
func callbackQSerialPort_ReadData(ptr unsafe.Pointer, data C.struct_QtSerialPort_PackedString, maxSize C.longlong) C.longlong {
	if signal := qt.GetSignal(ptr, "readData"); signal != nil {
		var retS = cGoUnpackString(data)
		var ret = C.longlong(signal.(func(*string, int64) int64)(&retS, int64(maxSize)))
		if ret > 0 {
			C.memcpy(unsafe.Pointer(data.data), unsafe.Pointer(C.CString(retS)), C.size_t(ret))
		}
		return ret
	}
	var retS = cGoUnpackString(data)
	var ret = C.longlong(NewQSerialPortFromPointer(ptr).ReadDataDefault(&retS, int64(maxSize)))
	if ret > 0 {
		C.memcpy(unsafe.Pointer(data.data), unsafe.Pointer(C.CString(retS)), C.size_t(ret))
	}
	return ret
}

func (ptr *QSerialPort) ConnectReadData(f func(data *string, maxSize int64) int64) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "readData"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "readData", func(data *string, maxSize int64) int64 {
				signal.(func(*string, int64) int64)(data, maxSize)
				return f(data, maxSize)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "readData", f)
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
		var dataC = C.CString(strings.Repeat("0", int(maxSize)))
		defer C.free(unsafe.Pointer(dataC))
		var ret = int64(C.QSerialPort_ReadData(ptr.Pointer(), dataC, C.longlong(maxSize)))
		if ret > 0 {
			*data = C.GoStringN(dataC, C.int(ret))
		}
		return ret
	}
	return 0
}

func (ptr *QSerialPort) ReadDataDefault(data *string, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		var dataC = C.CString(strings.Repeat("0", int(maxSize)))
		defer C.free(unsafe.Pointer(dataC))
		var ret = int64(C.QSerialPort_ReadDataDefault(ptr.Pointer(), dataC, C.longlong(maxSize)))
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
		return C.longlong(signal.(func(string, int64) int64)(cGoUnpackString(data), int64(maxSize)))
	}

	return C.longlong(NewQSerialPortFromPointer(ptr).ReadLineDataDefault(cGoUnpackString(data), int64(maxSize)))
}

func (ptr *QSerialPort) ReadLineDataDefault(data string, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if data != "" {
			dataC = C.CString(data)
			defer C.free(unsafe.Pointer(dataC))
		}
		return int64(C.QSerialPort_ReadLineDataDefault(ptr.Pointer(), dataC, C.longlong(maxSize)))
	}
	return 0
}

//export callbackQSerialPort_WriteData
func callbackQSerialPort_WriteData(ptr unsafe.Pointer, data C.struct_QtSerialPort_PackedString, maxSize C.longlong) C.longlong {
	if signal := qt.GetSignal(ptr, "writeData"); signal != nil {
		return C.longlong(signal.(func(string, int64) int64)(cGoUnpackString(data), int64(maxSize)))
	}

	return C.longlong(NewQSerialPortFromPointer(ptr).WriteDataDefault(cGoUnpackString(data), int64(maxSize)))
}

func (ptr *QSerialPort) ConnectWriteData(f func(data string, maxSize int64) int64) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "writeData"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "writeData", func(data string, maxSize int64) int64 {
				signal.(func(string, int64) int64)(data, maxSize)
				return f(data, maxSize)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "writeData", f)
		}
	}
}

func (ptr *QSerialPort) DisconnectWriteData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "writeData")
	}
}

func (ptr *QSerialPort) WriteData(data string, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if data != "" {
			dataC = C.CString(data)
			defer C.free(unsafe.Pointer(dataC))
		}
		return int64(C.QSerialPort_WriteData(ptr.Pointer(), dataC, C.longlong(maxSize)))
	}
	return 0
}

func (ptr *QSerialPort) WriteDataDefault(data string, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if data != "" {
			dataC = C.CString(data)
			defer C.free(unsafe.Pointer(dataC))
		}
		return int64(C.QSerialPort_WriteDataDefault(ptr.Pointer(), dataC, C.longlong(maxSize)))
	}
	return 0
}

//export callbackQSerialPort_BaudRateChanged
func callbackQSerialPort_BaudRateChanged(ptr unsafe.Pointer, baudRate C.int, directions C.longlong) {
	if signal := qt.GetSignal(ptr, "baudRateChanged"); signal != nil {
		signal.(func(int, QSerialPort__Direction))(int(int32(baudRate)), QSerialPort__Direction(directions))
	}

}

func (ptr *QSerialPort) ConnectBaudRateChanged(f func(baudRate int, directions QSerialPort__Direction)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "baudRateChanged") {
			C.QSerialPort_ConnectBaudRateChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "baudRateChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "baudRateChanged", func(baudRate int, directions QSerialPort__Direction) {
				signal.(func(int, QSerialPort__Direction))(baudRate, directions)
				f(baudRate, directions)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "baudRateChanged", f)
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
		signal.(func(bool))(int8(set) != 0)
	}

}

func (ptr *QSerialPort) ConnectBreakEnabledChanged(f func(set bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "breakEnabledChanged") {
			C.QSerialPort_ConnectBreakEnabledChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "breakEnabledChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "breakEnabledChanged", func(set bool) {
				signal.(func(bool))(set)
				f(set)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "breakEnabledChanged", f)
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

func (ptr *QSerialPort) ClearError() {
	if ptr.Pointer() != nil {
		C.QSerialPort_ClearError(ptr.Pointer())
	}
}

//export callbackQSerialPort_Close
func callbackQSerialPort_Close(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		signal.(func())()
	} else {
		NewQSerialPortFromPointer(ptr).CloseDefault()
	}
}

func (ptr *QSerialPort) CloseDefault() {
	if ptr.Pointer() != nil {
		C.QSerialPort_CloseDefault(ptr.Pointer())
	}
}

//export callbackQSerialPort_DataBitsChanged
func callbackQSerialPort_DataBitsChanged(ptr unsafe.Pointer, dataBits C.longlong) {
	if signal := qt.GetSignal(ptr, "dataBitsChanged"); signal != nil {
		signal.(func(QSerialPort__DataBits))(QSerialPort__DataBits(dataBits))
	}

}

func (ptr *QSerialPort) ConnectDataBitsChanged(f func(dataBits QSerialPort__DataBits)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "dataBitsChanged") {
			C.QSerialPort_ConnectDataBitsChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "dataBitsChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "dataBitsChanged", func(dataBits QSerialPort__DataBits) {
				signal.(func(QSerialPort__DataBits))(dataBits)
				f(dataBits)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "dataBitsChanged", f)
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
		signal.(func(bool))(int8(set) != 0)
	}

}

func (ptr *QSerialPort) ConnectDataTerminalReadyChanged(f func(set bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "dataTerminalReadyChanged") {
			C.QSerialPort_ConnectDataTerminalReadyChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "dataTerminalReadyChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "dataTerminalReadyChanged", func(set bool) {
				signal.(func(bool))(set)
				f(set)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "dataTerminalReadyChanged", f)
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

//export callbackQSerialPort_ErrorOccurred
func callbackQSerialPort_ErrorOccurred(ptr unsafe.Pointer, error C.longlong) {
	if signal := qt.GetSignal(ptr, "errorOccurred"); signal != nil {
		signal.(func(QSerialPort__SerialPortError))(QSerialPort__SerialPortError(error))
	}

}

func (ptr *QSerialPort) ConnectErrorOccurred(f func(error QSerialPort__SerialPortError)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "errorOccurred") {
			C.QSerialPort_ConnectErrorOccurred(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "errorOccurred"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "errorOccurred", func(error QSerialPort__SerialPortError) {
				signal.(func(QSerialPort__SerialPortError))(error)
				f(error)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "errorOccurred", f)
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

//export callbackQSerialPort_FlowControlChanged
func callbackQSerialPort_FlowControlChanged(ptr unsafe.Pointer, flow C.longlong) {
	if signal := qt.GetSignal(ptr, "flowControlChanged"); signal != nil {
		signal.(func(QSerialPort__FlowControl))(QSerialPort__FlowControl(flow))
	}

}

func (ptr *QSerialPort) ConnectFlowControlChanged(f func(flow QSerialPort__FlowControl)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "flowControlChanged") {
			C.QSerialPort_ConnectFlowControlChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "flowControlChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "flowControlChanged", func(flow QSerialPort__FlowControl) {
				signal.(func(QSerialPort__FlowControl))(flow)
				f(flow)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "flowControlChanged", f)
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

//export callbackQSerialPort_ParityChanged
func callbackQSerialPort_ParityChanged(ptr unsafe.Pointer, parity C.longlong) {
	if signal := qt.GetSignal(ptr, "parityChanged"); signal != nil {
		signal.(func(QSerialPort__Parity))(QSerialPort__Parity(parity))
	}

}

func (ptr *QSerialPort) ConnectParityChanged(f func(parity QSerialPort__Parity)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "parityChanged") {
			C.QSerialPort_ConnectParityChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "parityChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "parityChanged", func(parity QSerialPort__Parity) {
				signal.(func(QSerialPort__Parity))(parity)
				f(parity)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "parityChanged", f)
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

//export callbackQSerialPort_RequestToSendChanged
func callbackQSerialPort_RequestToSendChanged(ptr unsafe.Pointer, set C.char) {
	if signal := qt.GetSignal(ptr, "requestToSendChanged"); signal != nil {
		signal.(func(bool))(int8(set) != 0)
	}

}

func (ptr *QSerialPort) ConnectRequestToSendChanged(f func(set bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "requestToSendChanged") {
			C.QSerialPort_ConnectRequestToSendChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "requestToSendChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "requestToSendChanged", func(set bool) {
				signal.(func(bool))(set)
				f(set)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "requestToSendChanged", f)
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

//export callbackQSerialPort_StopBitsChanged
func callbackQSerialPort_StopBitsChanged(ptr unsafe.Pointer, stopBits C.longlong) {
	if signal := qt.GetSignal(ptr, "stopBitsChanged"); signal != nil {
		signal.(func(QSerialPort__StopBits))(QSerialPort__StopBits(stopBits))
	}

}

func (ptr *QSerialPort) ConnectStopBitsChanged(f func(stopBits QSerialPort__StopBits)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "stopBitsChanged") {
			C.QSerialPort_ConnectStopBitsChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "stopBitsChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "stopBitsChanged", func(stopBits QSerialPort__StopBits) {
				signal.(func(QSerialPort__StopBits))(stopBits)
				f(stopBits)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stopBitsChanged", f)
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

//export callbackQSerialPort_DestroyQSerialPort
func callbackQSerialPort_DestroyQSerialPort(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QSerialPort"); signal != nil {
		signal.(func())()
	} else {
		NewQSerialPortFromPointer(ptr).DestroyQSerialPortDefault()
	}
}

func (ptr *QSerialPort) ConnectDestroyQSerialPort(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QSerialPort"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QSerialPort", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QSerialPort", f)
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
		C.QSerialPort_DestroyQSerialPort(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QSerialPort) DestroyQSerialPortDefault() {
	if ptr.Pointer() != nil {
		C.QSerialPort_DestroyQSerialPortDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QSerialPort) DataBits() QSerialPort__DataBits {
	if ptr.Pointer() != nil {
		return QSerialPort__DataBits(C.QSerialPort_DataBits(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSerialPort) FlowControl() QSerialPort__FlowControl {
	if ptr.Pointer() != nil {
		return QSerialPort__FlowControl(C.QSerialPort_FlowControl(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSerialPort) Parity() QSerialPort__Parity {
	if ptr.Pointer() != nil {
		return QSerialPort__Parity(C.QSerialPort_Parity(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSerialPort) PortName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSerialPort_PortName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSerialPort) Error() QSerialPort__SerialPortError {
	if ptr.Pointer() != nil {
		return QSerialPort__SerialPortError(C.QSerialPort_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSerialPort) StopBits() QSerialPort__StopBits {
	if ptr.Pointer() != nil {
		return QSerialPort__StopBits(C.QSerialPort_StopBits(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSerialPort) IsBreakEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QSerialPort_IsBreakEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSerialPort) BaudRate(directions QSerialPort__Direction) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSerialPort_BaudRate(ptr.Pointer(), C.longlong(directions))))
	}
	return 0
}

//export callbackQSerialPort_AtEnd
func callbackQSerialPort_AtEnd(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "atEnd"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSerialPortFromPointer(ptr).AtEndDefault())))
}

func (ptr *QSerialPort) AtEndDefault() bool {
	if ptr.Pointer() != nil {
		return C.QSerialPort_AtEndDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSerialPort_CanReadLine
func callbackQSerialPort_CanReadLine(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canReadLine"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSerialPortFromPointer(ptr).CanReadLineDefault())))
}

func (ptr *QSerialPort) CanReadLineDefault() bool {
	if ptr.Pointer() != nil {
		return C.QSerialPort_CanReadLineDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSerialPort_IsSequential
func callbackQSerialPort_IsSequential(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isSequential"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSerialPortFromPointer(ptr).IsSequentialDefault())))
}

func (ptr *QSerialPort) IsSequentialDefault() bool {
	if ptr.Pointer() != nil {
		return C.QSerialPort_IsSequentialDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSerialPort_BytesAvailable
func callbackQSerialPort_BytesAvailable(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "bytesAvailable"); signal != nil {
		return C.longlong(signal.(func() int64)())
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
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQSerialPortFromPointer(ptr).BytesToWriteDefault())
}

func (ptr *QSerialPort) BytesToWriteDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QSerialPort_BytesToWriteDefault(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSerialPort) ReadBufferSize() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QSerialPort_ReadBufferSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSerialPort) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QSerialPort___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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
	return unsafe.Pointer(C.QSerialPort___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QSerialPort) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QSerialPort___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSerialPort) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSerialPort___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSerialPort) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QSerialPort___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QSerialPort) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QSerialPort___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
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
	return unsafe.Pointer(C.QSerialPort___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QSerialPort) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QSerialPort___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
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
	return unsafe.Pointer(C.QSerialPort___findChildren_newList(ptr.Pointer()))
}

func (ptr *QSerialPort) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QSerialPort___children_atList(ptr.Pointer(), C.int(int32(i))))
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
	return unsafe.Pointer(C.QSerialPort___children_newList(ptr.Pointer()))
}

//export callbackQSerialPort_Reset
func callbackQSerialPort_Reset(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "reset"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSerialPortFromPointer(ptr).ResetDefault())))
}

func (ptr *QSerialPort) ResetDefault() bool {
	if ptr.Pointer() != nil {
		return C.QSerialPort_ResetDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSerialPort_Seek
func callbackQSerialPort_Seek(ptr unsafe.Pointer, pos C.longlong) C.char {
	if signal := qt.GetSignal(ptr, "seek"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int64) bool)(int64(pos)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSerialPortFromPointer(ptr).SeekDefault(int64(pos)))))
}

func (ptr *QSerialPort) SeekDefault(pos int64) bool {
	if ptr.Pointer() != nil {
		return C.QSerialPort_SeekDefault(ptr.Pointer(), C.longlong(pos)) != 0
	}
	return false
}

//export callbackQSerialPort_AboutToClose
func callbackQSerialPort_AboutToClose(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "aboutToClose"); signal != nil {
		signal.(func())()
	}

}

//export callbackQSerialPort_BytesWritten
func callbackQSerialPort_BytesWritten(ptr unsafe.Pointer, bytes C.longlong) {
	if signal := qt.GetSignal(ptr, "bytesWritten"); signal != nil {
		signal.(func(int64))(int64(bytes))
	}

}

//export callbackQSerialPort_ChannelBytesWritten
func callbackQSerialPort_ChannelBytesWritten(ptr unsafe.Pointer, channel C.int, bytes C.longlong) {
	if signal := qt.GetSignal(ptr, "channelBytesWritten"); signal != nil {
		signal.(func(int, int64))(int(int32(channel)), int64(bytes))
	}

}

//export callbackQSerialPort_ChannelReadyRead
func callbackQSerialPort_ChannelReadyRead(ptr unsafe.Pointer, channel C.int) {
	if signal := qt.GetSignal(ptr, "channelReadyRead"); signal != nil {
		signal.(func(int))(int(int32(channel)))
	}

}

//export callbackQSerialPort_ReadChannelFinished
func callbackQSerialPort_ReadChannelFinished(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "readChannelFinished"); signal != nil {
		signal.(func())()
	}

}

//export callbackQSerialPort_ReadyRead
func callbackQSerialPort_ReadyRead(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "readyRead"); signal != nil {
		signal.(func())()
	}

}

//export callbackQSerialPort_Pos
func callbackQSerialPort_Pos(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "pos"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQSerialPortFromPointer(ptr).PosDefault())
}

func (ptr *QSerialPort) PosDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QSerialPort_PosDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQSerialPort_Size
func callbackQSerialPort_Size(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "size"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQSerialPortFromPointer(ptr).SizeDefault())
}

func (ptr *QSerialPort) SizeDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QSerialPort_SizeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQSerialPort_Event
func callbackQSerialPort_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSerialPortFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QSerialPort) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSerialPort_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQSerialPort_EventFilter
func callbackQSerialPort_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSerialPortFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QSerialPort) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSerialPort_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQSerialPort_ChildEvent
func callbackQSerialPort_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
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
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
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
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
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
		signal.(func())()
	} else {
		NewQSerialPortFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QSerialPort) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QSerialPort_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQSerialPort_Destroyed
func callbackQSerialPort_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQSerialPort_DisconnectNotify
func callbackQSerialPort_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSerialPortFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSerialPort) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSerialPort_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSerialPort_ObjectNameChanged
func callbackQSerialPort_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtSerialPort_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQSerialPort_TimerEvent
func callbackQSerialPort_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSerialPortFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSerialPort) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSerialPort_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQSerialPort_MetaObject
func callbackQSerialPort_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQSerialPortFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QSerialPort) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSerialPort_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
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

func NewQSerialPortInfoFromPointer(ptr unsafe.Pointer) *QSerialPortInfo {
	var n = new(QSerialPortInfo)
	n.SetPointer(ptr)
	return n
}
func QSerialPortInfo_AvailablePorts() []*QSerialPortInfo {
	return func(l C.struct_QtSerialPort_PackedList) []*QSerialPortInfo {
		var out = make([]*QSerialPortInfo, int(l.len))
		for i := 0; i < int(l.len); i++ {
			out[i] = NewQSerialPortInfoFromPointer(l.data).__availablePorts_atList(i)
		}
		return out
	}(C.QSerialPortInfo_QSerialPortInfo_AvailablePorts())
}

func (ptr *QSerialPortInfo) AvailablePorts() []*QSerialPortInfo {
	return func(l C.struct_QtSerialPort_PackedList) []*QSerialPortInfo {
		var out = make([]*QSerialPortInfo, int(l.len))
		for i := 0; i < int(l.len); i++ {
			out[i] = NewQSerialPortInfoFromPointer(l.data).__availablePorts_atList(i)
		}
		return out
	}(C.QSerialPortInfo_QSerialPortInfo_AvailablePorts())
}

func QSerialPortInfo_StandardBaudRates() []int {
	return func(l C.struct_QtSerialPort_PackedList) []int {
		var out = make([]int, int(l.len))
		for i := 0; i < int(l.len); i++ {
			out[i] = NewQSerialPortInfoFromPointer(l.data).__standardBaudRates_atList(i)
		}
		return out
	}(C.QSerialPortInfo_QSerialPortInfo_StandardBaudRates())
}

func (ptr *QSerialPortInfo) StandardBaudRates() []int {
	return func(l C.struct_QtSerialPort_PackedList) []int {
		var out = make([]int, int(l.len))
		for i := 0; i < int(l.len); i++ {
			out[i] = NewQSerialPortInfoFromPointer(l.data).__standardBaudRates_atList(i)
		}
		return out
	}(C.QSerialPortInfo_QSerialPortInfo_StandardBaudRates())
}

func NewQSerialPortInfo() *QSerialPortInfo {
	var tmpValue = NewQSerialPortInfoFromPointer(C.QSerialPortInfo_NewQSerialPortInfo())
	runtime.SetFinalizer(tmpValue, (*QSerialPortInfo).DestroyQSerialPortInfo)
	return tmpValue
}

func NewQSerialPortInfo2(port QSerialPort_ITF) *QSerialPortInfo {
	var tmpValue = NewQSerialPortInfoFromPointer(C.QSerialPortInfo_NewQSerialPortInfo2(PointerFromQSerialPort(port)))
	runtime.SetFinalizer(tmpValue, (*QSerialPortInfo).DestroyQSerialPortInfo)
	return tmpValue
}

func NewQSerialPortInfo4(other QSerialPortInfo_ITF) *QSerialPortInfo {
	var tmpValue = NewQSerialPortInfoFromPointer(C.QSerialPortInfo_NewQSerialPortInfo4(PointerFromQSerialPortInfo(other)))
	runtime.SetFinalizer(tmpValue, (*QSerialPortInfo).DestroyQSerialPortInfo)
	return tmpValue
}

func NewQSerialPortInfo3(name string) *QSerialPortInfo {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	var tmpValue = NewQSerialPortInfoFromPointer(C.QSerialPortInfo_NewQSerialPortInfo3(C.struct_QtSerialPort_PackedString{data: nameC, len: C.longlong(len(name))}))
	runtime.SetFinalizer(tmpValue, (*QSerialPortInfo).DestroyQSerialPortInfo)
	return tmpValue
}

func (ptr *QSerialPortInfo) Swap(other QSerialPortInfo_ITF) {
	if ptr.Pointer() != nil {
		C.QSerialPortInfo_Swap(ptr.Pointer(), PointerFromQSerialPortInfo(other))
	}
}

func (ptr *QSerialPortInfo) DestroyQSerialPortInfo() {
	if ptr.Pointer() != nil {
		C.QSerialPortInfo_DestroyQSerialPortInfo(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QSerialPortInfo) Description() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSerialPortInfo_Description(ptr.Pointer()))
	}
	return ""
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

func (ptr *QSerialPortInfo) SerialNumber() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSerialPortInfo_SerialNumber(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSerialPortInfo) SystemLocation() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSerialPortInfo_SystemLocation(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSerialPortInfo) HasProductIdentifier() bool {
	if ptr.Pointer() != nil {
		return C.QSerialPortInfo_HasProductIdentifier(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSerialPortInfo) HasVendorIdentifier() bool {
	if ptr.Pointer() != nil {
		return C.QSerialPortInfo_HasVendorIdentifier(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSerialPortInfo) IsBusy() bool {
	if ptr.Pointer() != nil {
		return C.QSerialPortInfo_IsBusy(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSerialPortInfo) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QSerialPortInfo_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSerialPortInfo) ProductIdentifier() uint16 {
	if ptr.Pointer() != nil {
		return uint16(C.QSerialPortInfo_ProductIdentifier(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSerialPortInfo) VendorIdentifier() uint16 {
	if ptr.Pointer() != nil {
		return uint16(C.QSerialPortInfo_VendorIdentifier(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSerialPortInfo) __availablePorts_atList(i int) *QSerialPortInfo {
	if ptr.Pointer() != nil {
		var tmpValue = NewQSerialPortInfoFromPointer(C.QSerialPortInfo___availablePorts_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QSerialPortInfo).DestroyQSerialPortInfo)
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
	return unsafe.Pointer(C.QSerialPortInfo___availablePorts_newList(ptr.Pointer()))
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
	return unsafe.Pointer(C.QSerialPortInfo___standardBaudRates_newList(ptr.Pointer()))
}
