package bluetooth

//#include "qlowenergyservice.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QLowEnergyService struct {
	core.QObject
}

type QLowEnergyServiceITF interface {
	core.QObjectITF
	QLowEnergyServicePTR() *QLowEnergyService
}

func PointerFromQLowEnergyService(ptr QLowEnergyServiceITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLowEnergyServicePTR().Pointer()
	}
	return nil
}

func QLowEnergyServiceFromPointer(ptr unsafe.Pointer) *QLowEnergyService {
	var n = new(QLowEnergyService)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QLowEnergyService_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QLowEnergyService) QLowEnergyServicePTR() *QLowEnergyService {
	return ptr
}

//QLowEnergyService::ServiceError
type QLowEnergyService__ServiceError int

var (
	QLowEnergyService__NoError                  = QLowEnergyService__ServiceError(0)
	QLowEnergyService__OperationError           = QLowEnergyService__ServiceError(1)
	QLowEnergyService__CharacteristicWriteError = QLowEnergyService__ServiceError(2)
	QLowEnergyService__DescriptorWriteError     = QLowEnergyService__ServiceError(3)
	QLowEnergyService__UnknownError             = QLowEnergyService__ServiceError(4)
	QLowEnergyService__CharacteristicReadError  = QLowEnergyService__ServiceError(5)
	QLowEnergyService__DescriptorReadError      = QLowEnergyService__ServiceError(6)
)

//QLowEnergyService::ServiceState
type QLowEnergyService__ServiceState int

var (
	QLowEnergyService__InvalidService      = QLowEnergyService__ServiceState(0)
	QLowEnergyService__DiscoveryRequired   = QLowEnergyService__ServiceState(1)
	QLowEnergyService__DiscoveringServices = QLowEnergyService__ServiceState(2)
	QLowEnergyService__ServiceDiscovered   = QLowEnergyService__ServiceState(3)
)

//QLowEnergyService::ServiceType
type QLowEnergyService__ServiceType int

var (
	QLowEnergyService__PrimaryService  = QLowEnergyService__ServiceType(0x0001)
	QLowEnergyService__IncludedService = QLowEnergyService__ServiceType(0x0002)
)

//QLowEnergyService::WriteMode
type QLowEnergyService__WriteMode int

var (
	QLowEnergyService__WriteWithResponse    = QLowEnergyService__WriteMode(0)
	QLowEnergyService__WriteWithoutResponse = QLowEnergyService__WriteMode(1)
)

func (ptr *QLowEnergyService) ConnectStateChanged(f func(newState QLowEnergyService__ServiceState)) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_ConnectStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QLowEnergyService) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_DisconnectStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQLowEnergyServiceStateChanged
func callbackQLowEnergyServiceStateChanged(ptrName *C.char, newState C.int) {
	qt.GetSignal(C.GoString(ptrName), "stateChanged").(func(QLowEnergyService__ServiceState))(QLowEnergyService__ServiceState(newState))
}

func (ptr *QLowEnergyService) Contains(characteristic QLowEnergyCharacteristicITF) bool {
	if ptr.Pointer() != nil {
		return C.QLowEnergyService_Contains(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQLowEnergyCharacteristic(characteristic))) != 0
	}
	return false
}

func (ptr *QLowEnergyService) Contains2(descriptor QLowEnergyDescriptorITF) bool {
	if ptr.Pointer() != nil {
		return C.QLowEnergyService_Contains2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQLowEnergyDescriptor(descriptor))) != 0
	}
	return false
}

func (ptr *QLowEnergyService) DiscoverDetails() {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_DiscoverDetails(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QLowEnergyService) Error() QLowEnergyService__ServiceError {
	if ptr.Pointer() != nil {
		return QLowEnergyService__ServiceError(C.QLowEnergyService_Error(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLowEnergyService) ReadCharacteristic(characteristic QLowEnergyCharacteristicITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_ReadCharacteristic(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQLowEnergyCharacteristic(characteristic)))
	}
}

func (ptr *QLowEnergyService) ReadDescriptor(descriptor QLowEnergyDescriptorITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_ReadDescriptor(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQLowEnergyDescriptor(descriptor)))
	}
}

func (ptr *QLowEnergyService) ServiceName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QLowEnergyService_ServiceName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QLowEnergyService) Type() QLowEnergyService__ServiceType {
	if ptr.Pointer() != nil {
		return QLowEnergyService__ServiceType(C.QLowEnergyService_Type(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLowEnergyService) WriteCharacteristic(characteristic QLowEnergyCharacteristicITF, newValue core.QByteArrayITF, mode QLowEnergyService__WriteMode) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_WriteCharacteristic(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQLowEnergyCharacteristic(characteristic)), C.QtObjectPtr(core.PointerFromQByteArray(newValue)), C.int(mode))
	}
}

func (ptr *QLowEnergyService) WriteDescriptor(descriptor QLowEnergyDescriptorITF, newValue core.QByteArrayITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_WriteDescriptor(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQLowEnergyDescriptor(descriptor)), C.QtObjectPtr(core.PointerFromQByteArray(newValue)))
	}
}

func (ptr *QLowEnergyService) DestroyQLowEnergyService() {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_DestroyQLowEnergyService(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
