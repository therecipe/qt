package bluetooth

//#include "bluetooth.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QLowEnergyService struct {
	core.QObject
}

type QLowEnergyService_ITF interface {
	core.QObject_ITF
	QLowEnergyService_PTR() *QLowEnergyService
}

func PointerFromQLowEnergyService(ptr QLowEnergyService_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLowEnergyService_PTR().Pointer()
	}
	return nil
}

func NewQLowEnergyServiceFromPointer(ptr unsafe.Pointer) *QLowEnergyService {
	var n = new(QLowEnergyService)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QLowEnergyService_") {
		n.SetObjectName("QLowEnergyService_" + qt.Identifier())
	}
	return n
}

func (ptr *QLowEnergyService) QLowEnergyService_PTR() *QLowEnergyService {
	return ptr
}

//QLowEnergyService::ServiceError
type QLowEnergyService__ServiceError int64

const (
	QLowEnergyService__NoError                  = QLowEnergyService__ServiceError(0)
	QLowEnergyService__OperationError           = QLowEnergyService__ServiceError(1)
	QLowEnergyService__CharacteristicWriteError = QLowEnergyService__ServiceError(2)
	QLowEnergyService__DescriptorWriteError     = QLowEnergyService__ServiceError(3)
	QLowEnergyService__UnknownError             = QLowEnergyService__ServiceError(4)
	QLowEnergyService__CharacteristicReadError  = QLowEnergyService__ServiceError(5)
	QLowEnergyService__DescriptorReadError      = QLowEnergyService__ServiceError(6)
)

//QLowEnergyService::ServiceState
type QLowEnergyService__ServiceState int64

const (
	QLowEnergyService__InvalidService      = QLowEnergyService__ServiceState(0)
	QLowEnergyService__DiscoveryRequired   = QLowEnergyService__ServiceState(1)
	QLowEnergyService__DiscoveringServices = QLowEnergyService__ServiceState(2)
	QLowEnergyService__ServiceDiscovered   = QLowEnergyService__ServiceState(3)
)

//QLowEnergyService::ServiceType
type QLowEnergyService__ServiceType int64

const (
	QLowEnergyService__PrimaryService  = QLowEnergyService__ServiceType(0x0001)
	QLowEnergyService__IncludedService = QLowEnergyService__ServiceType(0x0002)
)

//QLowEnergyService::WriteMode
type QLowEnergyService__WriteMode int64

const (
	QLowEnergyService__WriteWithResponse    = QLowEnergyService__WriteMode(0)
	QLowEnergyService__WriteWithoutResponse = QLowEnergyService__WriteMode(1)
)

func (ptr *QLowEnergyService) ConnectError2(f func(newError QLowEnergyService__ServiceError)) {
	defer qt.Recovering("connect QLowEnergyService::error")

	if ptr.Pointer() != nil {
		C.QLowEnergyService_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error2", f)
	}
}

func (ptr *QLowEnergyService) DisconnectError2() {
	defer qt.Recovering("disconnect QLowEnergyService::error")

	if ptr.Pointer() != nil {
		C.QLowEnergyService_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error2")
	}
}

//export callbackQLowEnergyServiceError2
func callbackQLowEnergyServiceError2(ptrName *C.char, newError C.int) {
	defer qt.Recovering("callback QLowEnergyService::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(QLowEnergyService__ServiceError))(QLowEnergyService__ServiceError(newError))
	}

}

func (ptr *QLowEnergyService) ConnectStateChanged(f func(newState QLowEnergyService__ServiceState)) {
	defer qt.Recovering("connect QLowEnergyService::stateChanged")

	if ptr.Pointer() != nil {
		C.QLowEnergyService_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QLowEnergyService) DisconnectStateChanged() {
	defer qt.Recovering("disconnect QLowEnergyService::stateChanged")

	if ptr.Pointer() != nil {
		C.QLowEnergyService_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQLowEnergyServiceStateChanged
func callbackQLowEnergyServiceStateChanged(ptrName *C.char, newState C.int) {
	defer qt.Recovering("callback QLowEnergyService::stateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stateChanged"); signal != nil {
		signal.(func(QLowEnergyService__ServiceState))(QLowEnergyService__ServiceState(newState))
	}

}

func (ptr *QLowEnergyService) Contains(characteristic QLowEnergyCharacteristic_ITF) bool {
	defer qt.Recovering("QLowEnergyService::contains")

	if ptr.Pointer() != nil {
		return C.QLowEnergyService_Contains(ptr.Pointer(), PointerFromQLowEnergyCharacteristic(characteristic)) != 0
	}
	return false
}

func (ptr *QLowEnergyService) Contains2(descriptor QLowEnergyDescriptor_ITF) bool {
	defer qt.Recovering("QLowEnergyService::contains")

	if ptr.Pointer() != nil {
		return C.QLowEnergyService_Contains2(ptr.Pointer(), PointerFromQLowEnergyDescriptor(descriptor)) != 0
	}
	return false
}

func (ptr *QLowEnergyService) DiscoverDetails() {
	defer qt.Recovering("QLowEnergyService::discoverDetails")

	if ptr.Pointer() != nil {
		C.QLowEnergyService_DiscoverDetails(ptr.Pointer())
	}
}

func (ptr *QLowEnergyService) Error() QLowEnergyService__ServiceError {
	defer qt.Recovering("QLowEnergyService::error")

	if ptr.Pointer() != nil {
		return QLowEnergyService__ServiceError(C.QLowEnergyService_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyService) ReadCharacteristic(characteristic QLowEnergyCharacteristic_ITF) {
	defer qt.Recovering("QLowEnergyService::readCharacteristic")

	if ptr.Pointer() != nil {
		C.QLowEnergyService_ReadCharacteristic(ptr.Pointer(), PointerFromQLowEnergyCharacteristic(characteristic))
	}
}

func (ptr *QLowEnergyService) ReadDescriptor(descriptor QLowEnergyDescriptor_ITF) {
	defer qt.Recovering("QLowEnergyService::readDescriptor")

	if ptr.Pointer() != nil {
		C.QLowEnergyService_ReadDescriptor(ptr.Pointer(), PointerFromQLowEnergyDescriptor(descriptor))
	}
}

func (ptr *QLowEnergyService) ServiceName() string {
	defer qt.Recovering("QLowEnergyService::serviceName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QLowEnergyService_ServiceName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLowEnergyService) State() QLowEnergyService__ServiceState {
	defer qt.Recovering("QLowEnergyService::state")

	if ptr.Pointer() != nil {
		return QLowEnergyService__ServiceState(C.QLowEnergyService_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyService) Type() QLowEnergyService__ServiceType {
	defer qt.Recovering("QLowEnergyService::type")

	if ptr.Pointer() != nil {
		return QLowEnergyService__ServiceType(C.QLowEnergyService_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyService) WriteCharacteristic(characteristic QLowEnergyCharacteristic_ITF, newValue core.QByteArray_ITF, mode QLowEnergyService__WriteMode) {
	defer qt.Recovering("QLowEnergyService::writeCharacteristic")

	if ptr.Pointer() != nil {
		C.QLowEnergyService_WriteCharacteristic(ptr.Pointer(), PointerFromQLowEnergyCharacteristic(characteristic), core.PointerFromQByteArray(newValue), C.int(mode))
	}
}

func (ptr *QLowEnergyService) WriteDescriptor(descriptor QLowEnergyDescriptor_ITF, newValue core.QByteArray_ITF) {
	defer qt.Recovering("QLowEnergyService::writeDescriptor")

	if ptr.Pointer() != nil {
		C.QLowEnergyService_WriteDescriptor(ptr.Pointer(), PointerFromQLowEnergyDescriptor(descriptor), core.PointerFromQByteArray(newValue))
	}
}

func (ptr *QLowEnergyService) DestroyQLowEnergyService() {
	defer qt.Recovering("QLowEnergyService::~QLowEnergyService")

	if ptr.Pointer() != nil {
		C.QLowEnergyService_DestroyQLowEnergyService(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QLowEnergyService) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QLowEnergyService::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QLowEnergyService) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QLowEnergyService::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQLowEnergyServiceTimerEvent
func callbackQLowEnergyServiceTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLowEnergyService::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QLowEnergyService) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QLowEnergyService::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QLowEnergyService) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QLowEnergyService::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQLowEnergyServiceChildEvent
func callbackQLowEnergyServiceChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLowEnergyService::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QLowEnergyService) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QLowEnergyService::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QLowEnergyService) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QLowEnergyService::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQLowEnergyServiceCustomEvent
func callbackQLowEnergyServiceCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QLowEnergyService::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
