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

type QBluetoothLocalDevice_ITF interface {
	core.QObject_ITF
	QBluetoothLocalDevice_PTR() *QBluetoothLocalDevice
}

func PointerFromQBluetoothLocalDevice(ptr QBluetoothLocalDevice_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothLocalDevice_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothLocalDeviceFromPointer(ptr unsafe.Pointer) *QBluetoothLocalDevice {
	var n = new(QBluetoothLocalDevice)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QBluetoothLocalDevice_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QBluetoothLocalDevice) QBluetoothLocalDevice_PTR() *QBluetoothLocalDevice {
	return ptr
}

//QBluetoothLocalDevice::Error
type QBluetoothLocalDevice__Error int64

const (
	QBluetoothLocalDevice__NoError      = QBluetoothLocalDevice__Error(0)
	QBluetoothLocalDevice__PairingError = QBluetoothLocalDevice__Error(1)
	QBluetoothLocalDevice__UnknownError = QBluetoothLocalDevice__Error(100)
)

//QBluetoothLocalDevice::HostMode
type QBluetoothLocalDevice__HostMode int64

const (
	QBluetoothLocalDevice__HostPoweredOff                 = QBluetoothLocalDevice__HostMode(0)
	QBluetoothLocalDevice__HostConnectable                = QBluetoothLocalDevice__HostMode(1)
	QBluetoothLocalDevice__HostDiscoverable               = QBluetoothLocalDevice__HostMode(2)
	QBluetoothLocalDevice__HostDiscoverableLimitedInquiry = QBluetoothLocalDevice__HostMode(3)
)

//QBluetoothLocalDevice::Pairing
type QBluetoothLocalDevice__Pairing int64

const (
	QBluetoothLocalDevice__Unpaired         = QBluetoothLocalDevice__Pairing(0)
	QBluetoothLocalDevice__Paired           = QBluetoothLocalDevice__Pairing(1)
	QBluetoothLocalDevice__AuthorizedPaired = QBluetoothLocalDevice__Pairing(2)
)

func (ptr *QBluetoothLocalDevice) ConnectError(f func(error QBluetoothLocalDevice__Error)) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_ConnectError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error", f)
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectError() {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DisconnectError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error")
	}
}

//export callbackQBluetoothLocalDeviceError
func callbackQBluetoothLocalDeviceError(ptrName *C.char, error C.int) {
	qt.GetSignal(C.GoString(ptrName), "error").(func(QBluetoothLocalDevice__Error))(QBluetoothLocalDevice__Error(error))
}

func (ptr *QBluetoothLocalDevice) ConnectHostModeStateChanged(f func(state QBluetoothLocalDevice__HostMode)) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_ConnectHostModeStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "hostModeStateChanged", f)
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectHostModeStateChanged() {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DisconnectHostModeStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "hostModeStateChanged")
	}
}

//export callbackQBluetoothLocalDeviceHostModeStateChanged
func callbackQBluetoothLocalDeviceHostModeStateChanged(ptrName *C.char, state C.int) {
	qt.GetSignal(C.GoString(ptrName), "hostModeStateChanged").(func(QBluetoothLocalDevice__HostMode))(QBluetoothLocalDevice__HostMode(state))
}

func (ptr *QBluetoothLocalDevice) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothLocalDevice_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothLocalDevice) DestroyQBluetoothLocalDevice() {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DestroyQBluetoothLocalDevice(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func NewQBluetoothLocalDevice(parent core.QObject_ITF) *QBluetoothLocalDevice {
	return NewQBluetoothLocalDeviceFromPointer(C.QBluetoothLocalDevice_NewQBluetoothLocalDevice(core.PointerFromQObject(parent)))
}

func NewQBluetoothLocalDevice2(address QBluetoothAddress_ITF, parent core.QObject_ITF) *QBluetoothLocalDevice {
	return NewQBluetoothLocalDeviceFromPointer(C.QBluetoothLocalDevice_NewQBluetoothLocalDevice2(PointerFromQBluetoothAddress(address), core.PointerFromQObject(parent)))
}

func (ptr *QBluetoothLocalDevice) HostMode() QBluetoothLocalDevice__HostMode {
	if ptr.Pointer() != nil {
		return QBluetoothLocalDevice__HostMode(C.QBluetoothLocalDevice_HostMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothLocalDevice) Name() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothLocalDevice_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothLocalDevice) PairingConfirmation(accept bool) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_PairingConfirmation(ptr.Pointer(), C.int(qt.GoBoolToInt(accept)))
	}
}

func (ptr *QBluetoothLocalDevice) PairingStatus(address QBluetoothAddress_ITF) QBluetoothLocalDevice__Pairing {
	if ptr.Pointer() != nil {
		return QBluetoothLocalDevice__Pairing(C.QBluetoothLocalDevice_PairingStatus(ptr.Pointer(), PointerFromQBluetoothAddress(address)))
	}
	return 0
}

func (ptr *QBluetoothLocalDevice) PowerOn() {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_PowerOn(ptr.Pointer())
	}
}

func (ptr *QBluetoothLocalDevice) RequestPairing(address QBluetoothAddress_ITF, pairing QBluetoothLocalDevice__Pairing) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_RequestPairing(ptr.Pointer(), PointerFromQBluetoothAddress(address), C.int(pairing))
	}
}

func (ptr *QBluetoothLocalDevice) SetHostMode(mode QBluetoothLocalDevice__HostMode) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_SetHostMode(ptr.Pointer(), C.int(mode))
	}
}
