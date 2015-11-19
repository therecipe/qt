package core

//#include "qiodevice.h"
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
	if n.ObjectName() == "" {
		n.SetObjectName("QIODevice_" + qt.RandomIdentifier())
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
	if ptr.Pointer() != nil {
		return C.QIODevice_GetChar(ptr.Pointer(), C.CString(c)) != 0
	}
	return false
}

func (ptr *QIODevice) PutChar(c string) bool {
	if ptr.Pointer() != nil {
		return C.QIODevice_PutChar(ptr.Pointer(), C.CString(c)) != 0
	}
	return false
}

func (ptr *QIODevice) ConnectAboutToClose(f func()) {
	if ptr.Pointer() != nil {
		C.QIODevice_ConnectAboutToClose(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "aboutToClose", f)
	}
}

func (ptr *QIODevice) DisconnectAboutToClose() {
	if ptr.Pointer() != nil {
		C.QIODevice_DisconnectAboutToClose(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "aboutToClose")
	}
}

//export callbackQIODeviceAboutToClose
func callbackQIODeviceAboutToClose(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "aboutToClose").(func())()
}

func (ptr *QIODevice) AtEnd() bool {
	if ptr.Pointer() != nil {
		return C.QIODevice_AtEnd(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QIODevice) CanReadLine() bool {
	if ptr.Pointer() != nil {
		return C.QIODevice_CanReadLine(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QIODevice) Close() {
	if ptr.Pointer() != nil {
		C.QIODevice_Close(ptr.Pointer())
	}
}

func (ptr *QIODevice) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QIODevice_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QIODevice) IsOpen() bool {
	if ptr.Pointer() != nil {
		return C.QIODevice_IsOpen(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QIODevice) IsReadable() bool {
	if ptr.Pointer() != nil {
		return C.QIODevice_IsReadable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QIODevice) IsSequential() bool {
	if ptr.Pointer() != nil {
		return C.QIODevice_IsSequential(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QIODevice) IsTextModeEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QIODevice_IsTextModeEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QIODevice) IsWritable() bool {
	if ptr.Pointer() != nil {
		return C.QIODevice_IsWritable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QIODevice) Open(mode QIODevice__OpenModeFlag) bool {
	if ptr.Pointer() != nil {
		return C.QIODevice_Open(ptr.Pointer(), C.int(mode)) != 0
	}
	return false
}

func (ptr *QIODevice) OpenMode() QIODevice__OpenModeFlag {
	if ptr.Pointer() != nil {
		return QIODevice__OpenModeFlag(C.QIODevice_OpenMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QIODevice) ReadAll() *QByteArray {
	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QIODevice_ReadAll(ptr.Pointer()))
	}
	return nil
}

func (ptr *QIODevice) ConnectReadChannelFinished(f func()) {
	if ptr.Pointer() != nil {
		C.QIODevice_ConnectReadChannelFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "readChannelFinished", f)
	}
}

func (ptr *QIODevice) DisconnectReadChannelFinished() {
	if ptr.Pointer() != nil {
		C.QIODevice_DisconnectReadChannelFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "readChannelFinished")
	}
}

//export callbackQIODeviceReadChannelFinished
func callbackQIODeviceReadChannelFinished(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "readChannelFinished").(func())()
}

func (ptr *QIODevice) ConnectReadyRead(f func()) {
	if ptr.Pointer() != nil {
		C.QIODevice_ConnectReadyRead(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "readyRead", f)
	}
}

func (ptr *QIODevice) DisconnectReadyRead() {
	if ptr.Pointer() != nil {
		C.QIODevice_DisconnectReadyRead(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "readyRead")
	}
}

//export callbackQIODeviceReadyRead
func callbackQIODeviceReadyRead(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "readyRead").(func())()
}

func (ptr *QIODevice) Reset() bool {
	if ptr.Pointer() != nil {
		return C.QIODevice_Reset(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QIODevice) SetTextModeEnabled(enabled bool) {
	if ptr.Pointer() != nil {
		C.QIODevice_SetTextModeEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QIODevice) UngetChar(c string) {
	if ptr.Pointer() != nil {
		C.QIODevice_UngetChar(ptr.Pointer(), C.CString(c))
	}
}

func (ptr *QIODevice) WaitForBytesWritten(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QIODevice_WaitForBytesWritten(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QIODevice) WaitForReadyRead(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QIODevice_WaitForReadyRead(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QIODevice) DestroyQIODevice() {
	if ptr.Pointer() != nil {
		C.QIODevice_DestroyQIODevice(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
