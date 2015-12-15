package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QIODevice struct {
	QObject
}

type QIODevice_ITF interface {
	QObject_ITF
	QIODevice_PTR() *QIODevice
}

func PointerFromQIODevice(ptr QIODevice_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QIODevice_PTR().Pointer()
	}
	return nil
}

func NewQIODeviceFromPointer(ptr unsafe.Pointer) *QIODevice {
	var n = new(QIODevice)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QIODevice_") {
		n.SetObjectName("QIODevice_" + qt.Identifier())
	}
	return n
}

func (ptr *QIODevice) QIODevice_PTR() *QIODevice {
	return ptr
}

//QIODevice::OpenModeFlag
type QIODevice__OpenModeFlag int64

const (
	QIODevice__NotOpen    = QIODevice__OpenModeFlag(0x0000)
	QIODevice__ReadOnly   = QIODevice__OpenModeFlag(0x0001)
	QIODevice__WriteOnly  = QIODevice__OpenModeFlag(0x0002)
	QIODevice__ReadWrite  = QIODevice__OpenModeFlag(QIODevice__ReadOnly | QIODevice__WriteOnly)
	QIODevice__Append     = QIODevice__OpenModeFlag(0x0004)
	QIODevice__Truncate   = QIODevice__OpenModeFlag(0x0008)
	QIODevice__Text       = QIODevice__OpenModeFlag(0x0010)
	QIODevice__Unbuffered = QIODevice__OpenModeFlag(0x0020)
)

func (ptr *QIODevice) GetChar(c string) bool {
	defer qt.Recovering("QIODevice::getChar")

	if ptr.Pointer() != nil {
		return C.QIODevice_GetChar(ptr.Pointer(), C.CString(c)) != 0
	}
	return false
}

func (ptr *QIODevice) PutChar(c string) bool {
	defer qt.Recovering("QIODevice::putChar")

	if ptr.Pointer() != nil {
		return C.QIODevice_PutChar(ptr.Pointer(), C.CString(c)) != 0
	}
	return false
}

func (ptr *QIODevice) ConnectAboutToClose(f func()) {
	defer qt.Recovering("connect QIODevice::aboutToClose")

	if ptr.Pointer() != nil {
		C.QIODevice_ConnectAboutToClose(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "aboutToClose", f)
	}
}

func (ptr *QIODevice) DisconnectAboutToClose() {
	defer qt.Recovering("disconnect QIODevice::aboutToClose")

	if ptr.Pointer() != nil {
		C.QIODevice_DisconnectAboutToClose(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "aboutToClose")
	}
}

//export callbackQIODeviceAboutToClose
func callbackQIODeviceAboutToClose(ptrName *C.char) {
	defer qt.Recovering("callback QIODevice::aboutToClose")

	var signal = qt.GetSignal(C.GoString(ptrName), "aboutToClose")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QIODevice) AtEnd() bool {
	defer qt.Recovering("QIODevice::atEnd")

	if ptr.Pointer() != nil {
		return C.QIODevice_AtEnd(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QIODevice) CanReadLine() bool {
	defer qt.Recovering("QIODevice::canReadLine")

	if ptr.Pointer() != nil {
		return C.QIODevice_CanReadLine(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QIODevice) ConnectClose(f func()) {
	defer qt.Recovering("connect QIODevice::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "close", f)
	}
}

func (ptr *QIODevice) DisconnectClose() {
	defer qt.Recovering("disconnect QIODevice::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "close")
	}
}

//export callbackQIODeviceClose
func callbackQIODeviceClose(ptrName *C.char) bool {
	defer qt.Recovering("callback QIODevice::close")

	var signal = qt.GetSignal(C.GoString(ptrName), "close")
	if signal != nil {
		defer signal.(func())()
		return true
	}
	return false

}

func (ptr *QIODevice) ErrorString() string {
	defer qt.Recovering("QIODevice::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QIODevice_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QIODevice) IsOpen() bool {
	defer qt.Recovering("QIODevice::isOpen")

	if ptr.Pointer() != nil {
		return C.QIODevice_IsOpen(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QIODevice) IsReadable() bool {
	defer qt.Recovering("QIODevice::isReadable")

	if ptr.Pointer() != nil {
		return C.QIODevice_IsReadable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QIODevice) IsSequential() bool {
	defer qt.Recovering("QIODevice::isSequential")

	if ptr.Pointer() != nil {
		return C.QIODevice_IsSequential(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QIODevice) IsTextModeEnabled() bool {
	defer qt.Recovering("QIODevice::isTextModeEnabled")

	if ptr.Pointer() != nil {
		return C.QIODevice_IsTextModeEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QIODevice) IsWritable() bool {
	defer qt.Recovering("QIODevice::isWritable")

	if ptr.Pointer() != nil {
		return C.QIODevice_IsWritable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QIODevice) Open(mode QIODevice__OpenModeFlag) bool {
	defer qt.Recovering("QIODevice::open")

	if ptr.Pointer() != nil {
		return C.QIODevice_Open(ptr.Pointer(), C.int(mode)) != 0
	}
	return false
}

func (ptr *QIODevice) OpenMode() QIODevice__OpenModeFlag {
	defer qt.Recovering("QIODevice::openMode")

	if ptr.Pointer() != nil {
		return QIODevice__OpenModeFlag(C.QIODevice_OpenMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QIODevice) ReadAll() *QByteArray {
	defer qt.Recovering("QIODevice::readAll")

	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QIODevice_ReadAll(ptr.Pointer()))
	}
	return nil
}

func (ptr *QIODevice) ConnectReadChannelFinished(f func()) {
	defer qt.Recovering("connect QIODevice::readChannelFinished")

	if ptr.Pointer() != nil {
		C.QIODevice_ConnectReadChannelFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "readChannelFinished", f)
	}
}

func (ptr *QIODevice) DisconnectReadChannelFinished() {
	defer qt.Recovering("disconnect QIODevice::readChannelFinished")

	if ptr.Pointer() != nil {
		C.QIODevice_DisconnectReadChannelFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "readChannelFinished")
	}
}

//export callbackQIODeviceReadChannelFinished
func callbackQIODeviceReadChannelFinished(ptrName *C.char) {
	defer qt.Recovering("callback QIODevice::readChannelFinished")

	var signal = qt.GetSignal(C.GoString(ptrName), "readChannelFinished")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QIODevice) ConnectReadyRead(f func()) {
	defer qt.Recovering("connect QIODevice::readyRead")

	if ptr.Pointer() != nil {
		C.QIODevice_ConnectReadyRead(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "readyRead", f)
	}
}

func (ptr *QIODevice) DisconnectReadyRead() {
	defer qt.Recovering("disconnect QIODevice::readyRead")

	if ptr.Pointer() != nil {
		C.QIODevice_DisconnectReadyRead(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "readyRead")
	}
}

//export callbackQIODeviceReadyRead
func callbackQIODeviceReadyRead(ptrName *C.char) {
	defer qt.Recovering("callback QIODevice::readyRead")

	var signal = qt.GetSignal(C.GoString(ptrName), "readyRead")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QIODevice) Reset() bool {
	defer qt.Recovering("QIODevice::reset")

	if ptr.Pointer() != nil {
		return C.QIODevice_Reset(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QIODevice) SetTextModeEnabled(enabled bool) {
	defer qt.Recovering("QIODevice::setTextModeEnabled")

	if ptr.Pointer() != nil {
		C.QIODevice_SetTextModeEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QIODevice) UngetChar(c string) {
	defer qt.Recovering("QIODevice::ungetChar")

	if ptr.Pointer() != nil {
		C.QIODevice_UngetChar(ptr.Pointer(), C.CString(c))
	}
}

func (ptr *QIODevice) WaitForBytesWritten(msecs int) bool {
	defer qt.Recovering("QIODevice::waitForBytesWritten")

	if ptr.Pointer() != nil {
		return C.QIODevice_WaitForBytesWritten(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QIODevice) WaitForReadyRead(msecs int) bool {
	defer qt.Recovering("QIODevice::waitForReadyRead")

	if ptr.Pointer() != nil {
		return C.QIODevice_WaitForReadyRead(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QIODevice) DestroyQIODevice() {
	defer qt.Recovering("QIODevice::~QIODevice")

	if ptr.Pointer() != nil {
		C.QIODevice_DestroyQIODevice(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
