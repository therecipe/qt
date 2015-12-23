package bluetooth

//#include "bluetooth.h"
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
	for len(n.ObjectName()) < len("QLowEnergyController_") {
		n.SetObjectName("QLowEnergyController_" + qt.Identifier())
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
	defer qt.Recovering("connect QLowEnergyController::connected")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_ConnectConnected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "connected", f)
	}
}

func (ptr *QLowEnergyController) DisconnectConnected() {
	defer qt.Recovering("disconnect QLowEnergyController::connected")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_DisconnectConnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "connected")
	}
}

//export callbackQLowEnergyControllerConnected
func callbackQLowEnergyControllerConnected(ptrName *C.char) {
	defer qt.Recovering("callback QLowEnergyController::connected")

	if signal := qt.GetSignal(C.GoString(ptrName), "connected"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QLowEnergyController) ConnectDisconnected(f func()) {
	defer qt.Recovering("connect QLowEnergyController::disconnected")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_ConnectDisconnected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "disconnected", f)
	}
}

func (ptr *QLowEnergyController) DisconnectDisconnected() {
	defer qt.Recovering("disconnect QLowEnergyController::disconnected")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_DisconnectDisconnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "disconnected")
	}
}

//export callbackQLowEnergyControllerDisconnected
func callbackQLowEnergyControllerDisconnected(ptrName *C.char) {
	defer qt.Recovering("callback QLowEnergyController::disconnected")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnected"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QLowEnergyController) ConnectDiscoveryFinished(f func()) {
	defer qt.Recovering("connect QLowEnergyController::discoveryFinished")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_ConnectDiscoveryFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "discoveryFinished", f)
	}
}

func (ptr *QLowEnergyController) DisconnectDiscoveryFinished() {
	defer qt.Recovering("disconnect QLowEnergyController::discoveryFinished")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_DisconnectDiscoveryFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "discoveryFinished")
	}
}

//export callbackQLowEnergyControllerDiscoveryFinished
func callbackQLowEnergyControllerDiscoveryFinished(ptrName *C.char) {
	defer qt.Recovering("callback QLowEnergyController::discoveryFinished")

	if signal := qt.GetSignal(C.GoString(ptrName), "discoveryFinished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QLowEnergyController) ConnectError2(f func(newError QLowEnergyController__Error)) {
	defer qt.Recovering("connect QLowEnergyController::error")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error2", f)
	}
}

func (ptr *QLowEnergyController) DisconnectError2() {
	defer qt.Recovering("disconnect QLowEnergyController::error")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error2")
	}
}

//export callbackQLowEnergyControllerError2
func callbackQLowEnergyControllerError2(ptrName *C.char, newError C.int) {
	defer qt.Recovering("callback QLowEnergyController::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(QLowEnergyController__Error))(QLowEnergyController__Error(newError))
	}

}

func (ptr *QLowEnergyController) ConnectStateChanged(f func(state QLowEnergyController__ControllerState)) {
	defer qt.Recovering("connect QLowEnergyController::stateChanged")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QLowEnergyController) DisconnectStateChanged() {
	defer qt.Recovering("disconnect QLowEnergyController::stateChanged")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQLowEnergyControllerStateChanged
func callbackQLowEnergyControllerStateChanged(ptrName *C.char, state C.int) {
	defer qt.Recovering("callback QLowEnergyController::stateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stateChanged"); signal != nil {
		signal.(func(QLowEnergyController__ControllerState))(QLowEnergyController__ControllerState(state))
	}

}

func NewQLowEnergyController(remoteDeviceInfo QBluetoothDeviceInfo_ITF, parent core.QObject_ITF) *QLowEnergyController {
	defer qt.Recovering("QLowEnergyController::QLowEnergyController")

	return NewQLowEnergyControllerFromPointer(C.QLowEnergyController_NewQLowEnergyController(PointerFromQBluetoothDeviceInfo(remoteDeviceInfo), core.PointerFromQObject(parent)))
}

func (ptr *QLowEnergyController) ConnectToDevice() {
	defer qt.Recovering("QLowEnergyController::connectToDevice")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_ConnectToDevice(ptr.Pointer())
	}
}

func (ptr *QLowEnergyController) CreateServiceObject(serviceUuid QBluetoothUuid_ITF, parent core.QObject_ITF) *QLowEnergyService {
	defer qt.Recovering("QLowEnergyController::createServiceObject")

	if ptr.Pointer() != nil {
		return NewQLowEnergyServiceFromPointer(C.QLowEnergyController_CreateServiceObject(ptr.Pointer(), PointerFromQBluetoothUuid(serviceUuid), core.PointerFromQObject(parent)))
	}
	return nil
}

func (ptr *QLowEnergyController) DisconnectFromDevice() {
	defer qt.Recovering("QLowEnergyController::disconnectFromDevice")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_DisconnectFromDevice(ptr.Pointer())
	}
}

func (ptr *QLowEnergyController) DiscoverServices() {
	defer qt.Recovering("QLowEnergyController::discoverServices")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_DiscoverServices(ptr.Pointer())
	}
}

func (ptr *QLowEnergyController) Error() QLowEnergyController__Error {
	defer qt.Recovering("QLowEnergyController::error")

	if ptr.Pointer() != nil {
		return QLowEnergyController__Error(C.QLowEnergyController_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyController) ErrorString() string {
	defer qt.Recovering("QLowEnergyController::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QLowEnergyController_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLowEnergyController) RemoteAddressType() QLowEnergyController__RemoteAddressType {
	defer qt.Recovering("QLowEnergyController::remoteAddressType")

	if ptr.Pointer() != nil {
		return QLowEnergyController__RemoteAddressType(C.QLowEnergyController_RemoteAddressType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyController) RemoteName() string {
	defer qt.Recovering("QLowEnergyController::remoteName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QLowEnergyController_RemoteName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLowEnergyController) SetRemoteAddressType(ty QLowEnergyController__RemoteAddressType) {
	defer qt.Recovering("QLowEnergyController::setRemoteAddressType")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_SetRemoteAddressType(ptr.Pointer(), C.int(ty))
	}
}

func (ptr *QLowEnergyController) State() QLowEnergyController__ControllerState {
	defer qt.Recovering("QLowEnergyController::state")

	if ptr.Pointer() != nil {
		return QLowEnergyController__ControllerState(C.QLowEnergyController_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyController) DestroyQLowEnergyController() {
	defer qt.Recovering("QLowEnergyController::~QLowEnergyController")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_DestroyQLowEnergyController(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QLowEnergyController) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QLowEnergyController::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QLowEnergyController) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QLowEnergyController::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQLowEnergyControllerTimerEvent
func callbackQLowEnergyControllerTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLowEnergyController::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QLowEnergyController) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QLowEnergyController::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QLowEnergyController) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QLowEnergyController::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQLowEnergyControllerChildEvent
func callbackQLowEnergyControllerChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLowEnergyController::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QLowEnergyController) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QLowEnergyController::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QLowEnergyController) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QLowEnergyController::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQLowEnergyControllerCustomEvent
func callbackQLowEnergyControllerCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLowEnergyController::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
