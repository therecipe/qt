package bluetooth

//#include "qbluetoothlocaldevice.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QBluetoothLocalDevice struct {
	core.QObject
}

type QBluetoothLocalDeviceITF interface {
	core.QObjectITF
	QBluetoothLocalDevicePTR() *QBluetoothLocalDevice
}

func PointerFromQBluetoothLocalDevice(ptr QBluetoothLocalDeviceITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothLocalDevicePTR().Pointer()
	}
	return nil
}

func QBluetoothLocalDeviceFromPointer(ptr unsafe.Pointer) *QBluetoothLocalDevice {
	var n = new(QBluetoothLocalDevice)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QBluetoothLocalDevice_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QBluetoothLocalDevice) QBluetoothLocalDevicePTR() *QBluetoothLocalDevice {
	return ptr
}

//QBluetoothLocalDevice::Error
type QBluetoothLocalDevice__Error int

var (
	QBluetoothLocalDevice__NoError      = QBluetoothLocalDevice__Error(0)
	QBluetoothLocalDevice__PairingError = QBluetoothLocalDevice__Error(1)
	QBluetoothLocalDevice__UnknownError = QBluetoothLocalDevice__Error(100)
)

//QBluetoothLocalDevice::HostMode
type QBluetoothLocalDevice__HostMode int

var (
	QBluetoothLocalDevice__HostPoweredOff                 = QBluetoothLocalDevice__HostMode(0)
	QBluetoothLocalDevice__HostConnectable                = QBluetoothLocalDevice__HostMode(1)
	QBluetoothLocalDevice__HostDiscoverable               = QBluetoothLocalDevice__HostMode(2)
	QBluetoothLocalDevice__HostDiscoverableLimitedInquiry = QBluetoothLocalDevice__HostMode(3)
)

//QBluetoothLocalDevice::Pairing
type QBluetoothLocalDevice__Pairing int

var (
	QBluetoothLocalDevice__Unpaired         = QBluetoothLocalDevice__Pairing(0)
	QBluetoothLocalDevice__Paired           = QBluetoothLocalDevice__Pairing(1)
	QBluetoothLocalDevice__AuthorizedPaired = QBluetoothLocalDevice__Pairing(2)
)

func (ptr *QBluetoothLocalDevice) ConnectError(f func(error QBluetoothLocalDevice__Error)) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_ConnectError(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "error", f)
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectError() {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DisconnectError(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "error")
	}
}

//export callbackQBluetoothLocalDeviceError
func callbackQBluetoothLocalDeviceError(ptrName *C.char, error C.int) {
	qt.GetSignal(C.GoString(ptrName), "error").(func(QBluetoothLocalDevice__Error))(QBluetoothLocalDevice__Error(error))
}

func (ptr *QBluetoothLocalDevice) ConnectHostModeStateChanged(f func(state QBluetoothLocalDevice__HostMode)) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_ConnectHostModeStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "hostModeStateChanged", f)
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectHostModeStateChanged() {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DisconnectHostModeStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "hostModeStateChanged")
	}
}

//export callbackQBluetoothLocalDeviceHostModeStateChanged
func callbackQBluetoothLocalDeviceHostModeStateChanged(ptrName *C.char, state C.int) {
	qt.GetSignal(C.GoString(ptrName), "hostModeStateChanged").(func(QBluetoothLocalDevice__HostMode))(QBluetoothLocalDevice__HostMode(state))
}

func (ptr *QBluetoothLocalDevice) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothLocalDevice_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QBluetoothLocalDevice) DestroyQBluetoothLocalDevice() {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DestroyQBluetoothLocalDevice(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func NewQBluetoothLocalDevice(parent core.QObjectITF) *QBluetoothLocalDevice {
	return QBluetoothLocalDeviceFromPointer(unsafe.Pointer(C.QBluetoothLocalDevice_NewQBluetoothLocalDevice(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func NewQBluetoothLocalDevice2(address QBluetoothAddressITF, parent core.QObjectITF) *QBluetoothLocalDevice {
	return QBluetoothLocalDeviceFromPointer(unsafe.Pointer(C.QBluetoothLocalDevice_NewQBluetoothLocalDevice2(C.QtObjectPtr(PointerFromQBluetoothAddress(address)), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QBluetoothLocalDevice) HostMode() QBluetoothLocalDevice__HostMode {
	if ptr.Pointer() != nil {
		return QBluetoothLocalDevice__HostMode(C.QBluetoothLocalDevice_HostMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QBluetoothLocalDevice) Name() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothLocalDevice_Name(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QBluetoothLocalDevice) PairingConfirmation(accept bool) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_PairingConfirmation(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(accept)))
	}
}

func (ptr *QBluetoothLocalDevice) PairingStatus(address QBluetoothAddressITF) QBluetoothLocalDevice__Pairing {
	if ptr.Pointer() != nil {
		return QBluetoothLocalDevice__Pairing(C.QBluetoothLocalDevice_PairingStatus(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQBluetoothAddress(address))))
	}
	return 0
}

func (ptr *QBluetoothLocalDevice) PowerOn() {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_PowerOn(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QBluetoothLocalDevice) RequestPairing(address QBluetoothAddressITF, pairing QBluetoothLocalDevice__Pairing) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_RequestPairing(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQBluetoothAddress(address)), C.int(pairing))
	}
}

func (ptr *QBluetoothLocalDevice) SetHostMode(mode QBluetoothLocalDevice__HostMode) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_SetHostMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}
