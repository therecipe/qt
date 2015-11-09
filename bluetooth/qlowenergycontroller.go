package bluetooth

//#include "qlowenergycontroller.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QLowEnergyController struct {
	core.QObject
}

type QLowEnergyController_ITF interface {
	core.QObject_ITF
	QLowEnergyController_PTR() *QLowEnergyController
}

func PointerFromQLowEnergyController(ptr QLowEnergyController_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLowEnergyController_PTR().Pointer()
	}
	return nil
}

func NewQLowEnergyControllerFromPointer(ptr unsafe.Pointer) *QLowEnergyController {
	var n = new(QLowEnergyController)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QLowEnergyController_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QLowEnergyController) QLowEnergyController_PTR() *QLowEnergyController {
	return ptr
}

//QLowEnergyController::ControllerState
type QLowEnergyController__ControllerState int64

const (
	QLowEnergyController__UnconnectedState = QLowEnergyController__ControllerState(0)
	QLowEnergyController__ConnectingState  = QLowEnergyController__ControllerState(1)
	QLowEnergyController__ConnectedState   = QLowEnergyController__ControllerState(2)
	QLowEnergyController__DiscoveringState = QLowEnergyController__ControllerState(3)
	QLowEnergyController__DiscoveredState  = QLowEnergyController__ControllerState(4)
	QLowEnergyController__ClosingState     = QLowEnergyController__ControllerState(5)
)

//QLowEnergyController::Error
type QLowEnergyController__Error int64

const (
	QLowEnergyController__NoError                      = QLowEnergyController__Error(0)
	QLowEnergyController__UnknownError                 = QLowEnergyController__Error(1)
	QLowEnergyController__UnknownRemoteDeviceError     = QLowEnergyController__Error(2)
	QLowEnergyController__NetworkError                 = QLowEnergyController__Error(3)
	QLowEnergyController__InvalidBluetoothAdapterError = QLowEnergyController__Error(4)
	QLowEnergyController__ConnectionError              = QLowEnergyController__Error(5)
)

//QLowEnergyController::RemoteAddressType
type QLowEnergyController__RemoteAddressType int64

const (
	QLowEnergyController__PublicAddress = QLowEnergyController__RemoteAddressType(0)
	QLowEnergyController__RandomAddress = QLowEnergyController__RemoteAddressType(1)
)

func (ptr *QLowEnergyController) ConnectConnected(f func()) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_ConnectConnected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "connected", f)
	}
}

func (ptr *QLowEnergyController) DisconnectConnected() {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_DisconnectConnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "connected")
	}
}

//export callbackQLowEnergyControllerConnected
func callbackQLowEnergyControllerConnected(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "connected").(func())()
}

func (ptr *QLowEnergyController) ConnectDisconnected(f func()) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_ConnectDisconnected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "disconnected", f)
	}
}

func (ptr *QLowEnergyController) DisconnectDisconnected() {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_DisconnectDisconnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "disconnected")
	}
}

//export callbackQLowEnergyControllerDisconnected
func callbackQLowEnergyControllerDisconnected(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "disconnected").(func())()
}

func (ptr *QLowEnergyController) ConnectDiscoveryFinished(f func()) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_ConnectDiscoveryFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "discoveryFinished", f)
	}
}

func (ptr *QLowEnergyController) DisconnectDiscoveryFinished() {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_DisconnectDiscoveryFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "discoveryFinished")
	}
}

//export callbackQLowEnergyControllerDiscoveryFinished
func callbackQLowEnergyControllerDiscoveryFinished(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "discoveryFinished").(func())()
}

func (ptr *QLowEnergyController) ConnectStateChanged(f func(state QLowEnergyController__ControllerState)) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QLowEnergyController) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQLowEnergyControllerStateChanged
func callbackQLowEnergyControllerStateChanged(ptrName *C.char, state C.int) {
	qt.GetSignal(C.GoString(ptrName), "stateChanged").(func(QLowEnergyController__ControllerState))(QLowEnergyController__ControllerState(state))
}

func NewQLowEnergyController(remoteDeviceInfo QBluetoothDeviceInfo_ITF, parent core.QObject_ITF) *QLowEnergyController {
	return NewQLowEnergyControllerFromPointer(C.QLowEnergyController_NewQLowEnergyController(PointerFromQBluetoothDeviceInfo(remoteDeviceInfo), core.PointerFromQObject(parent)))
}

func (ptr *QLowEnergyController) ConnectToDevice() {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_ConnectToDevice(ptr.Pointer())
	}
}

func (ptr *QLowEnergyController) CreateServiceObject(serviceUuid QBluetoothUuid_ITF, parent core.QObject_ITF) *QLowEnergyService {
	if ptr.Pointer() != nil {
		return NewQLowEnergyServiceFromPointer(C.QLowEnergyController_CreateServiceObject(ptr.Pointer(), PointerFromQBluetoothUuid(serviceUuid), core.PointerFromQObject(parent)))
	}
	return nil
}

func (ptr *QLowEnergyController) DisconnectFromDevice() {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_DisconnectFromDevice(ptr.Pointer())
	}
}

func (ptr *QLowEnergyController) DiscoverServices() {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_DiscoverServices(ptr.Pointer())
	}
}

func (ptr *QLowEnergyController) Error() QLowEnergyController__Error {
	if ptr.Pointer() != nil {
		return QLowEnergyController__Error(C.QLowEnergyController_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyController) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QLowEnergyController_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLowEnergyController) RemoteAddressType() QLowEnergyController__RemoteAddressType {
	if ptr.Pointer() != nil {
		return QLowEnergyController__RemoteAddressType(C.QLowEnergyController_RemoteAddressType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyController) RemoteName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QLowEnergyController_RemoteName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLowEnergyController) SetRemoteAddressType(ty QLowEnergyController__RemoteAddressType) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_SetRemoteAddressType(ptr.Pointer(), C.int(ty))
	}
}

func (ptr *QLowEnergyController) DestroyQLowEnergyController() {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_DestroyQLowEnergyController(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
