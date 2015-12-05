package bluetooth

//#include "bluetooth.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"unsafe"
)

type QBluetoothDeviceInfo struct {
	ptr unsafe.Pointer
}

type QBluetoothDeviceInfo_ITF interface {
	QBluetoothDeviceInfo_PTR() *QBluetoothDeviceInfo
}

func (p *QBluetoothDeviceInfo) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QBluetoothDeviceInfo) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQBluetoothDeviceInfo(ptr QBluetoothDeviceInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothDeviceInfo_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothDeviceInfoFromPointer(ptr unsafe.Pointer) *QBluetoothDeviceInfo {
	var n = new(QBluetoothDeviceInfo)
	n.SetPointer(ptr)
	return n
}

func (ptr *QBluetoothDeviceInfo) QBluetoothDeviceInfo_PTR() *QBluetoothDeviceInfo {
	return ptr
}

//QBluetoothDeviceInfo::CoreConfiguration
type QBluetoothDeviceInfo__CoreConfiguration int64

const (
	QBluetoothDeviceInfo__UnknownCoreConfiguration              = QBluetoothDeviceInfo__CoreConfiguration(0x0)
	QBluetoothDeviceInfo__LowEnergyCoreConfiguration            = QBluetoothDeviceInfo__CoreConfiguration(0x01)
	QBluetoothDeviceInfo__BaseRateCoreConfiguration             = QBluetoothDeviceInfo__CoreConfiguration(0x02)
	QBluetoothDeviceInfo__BaseRateAndLowEnergyCoreConfiguration = QBluetoothDeviceInfo__CoreConfiguration(0x03)
)

//QBluetoothDeviceInfo::DataCompleteness
type QBluetoothDeviceInfo__DataCompleteness int64

const (
	QBluetoothDeviceInfo__DataComplete    = QBluetoothDeviceInfo__DataCompleteness(0)
	QBluetoothDeviceInfo__DataIncomplete  = QBluetoothDeviceInfo__DataCompleteness(1)
	QBluetoothDeviceInfo__DataUnavailable = QBluetoothDeviceInfo__DataCompleteness(2)
)

//QBluetoothDeviceInfo::MajorDeviceClass
type QBluetoothDeviceInfo__MajorDeviceClass int64

const (
	QBluetoothDeviceInfo__MiscellaneousDevice = QBluetoothDeviceInfo__MajorDeviceClass(0)
	QBluetoothDeviceInfo__ComputerDevice      = QBluetoothDeviceInfo__MajorDeviceClass(1)
	QBluetoothDeviceInfo__PhoneDevice         = QBluetoothDeviceInfo__MajorDeviceClass(2)
	QBluetoothDeviceInfo__LANAccessDevice     = QBluetoothDeviceInfo__MajorDeviceClass(3)
	QBluetoothDeviceInfo__AudioVideoDevice    = QBluetoothDeviceInfo__MajorDeviceClass(4)
	QBluetoothDeviceInfo__PeripheralDevice    = QBluetoothDeviceInfo__MajorDeviceClass(5)
	QBluetoothDeviceInfo__ImagingDevice       = QBluetoothDeviceInfo__MajorDeviceClass(6)
	QBluetoothDeviceInfo__WearableDevice      = QBluetoothDeviceInfo__MajorDeviceClass(7)
	QBluetoothDeviceInfo__ToyDevice           = QBluetoothDeviceInfo__MajorDeviceClass(8)
	QBluetoothDeviceInfo__HealthDevice        = QBluetoothDeviceInfo__MajorDeviceClass(9)
	QBluetoothDeviceInfo__UncategorizedDevice = QBluetoothDeviceInfo__MajorDeviceClass(31)
)

//QBluetoothDeviceInfo::MinorAudioVideoClass
type QBluetoothDeviceInfo__MinorAudioVideoClass int64

const (
	QBluetoothDeviceInfo__UncategorizedAudioVideoDevice = QBluetoothDeviceInfo__MinorAudioVideoClass(0)
	QBluetoothDeviceInfo__WearableHeadsetDevice         = QBluetoothDeviceInfo__MinorAudioVideoClass(1)
	QBluetoothDeviceInfo__HandsFreeDevice               = QBluetoothDeviceInfo__MinorAudioVideoClass(2)
	QBluetoothDeviceInfo__Microphone                    = QBluetoothDeviceInfo__MinorAudioVideoClass(4)
	QBluetoothDeviceInfo__Loudspeaker                   = QBluetoothDeviceInfo__MinorAudioVideoClass(5)
	QBluetoothDeviceInfo__Headphones                    = QBluetoothDeviceInfo__MinorAudioVideoClass(6)
	QBluetoothDeviceInfo__PortableAudioDevice           = QBluetoothDeviceInfo__MinorAudioVideoClass(7)
	QBluetoothDeviceInfo__CarAudio                      = QBluetoothDeviceInfo__MinorAudioVideoClass(8)
	QBluetoothDeviceInfo__SetTopBox                     = QBluetoothDeviceInfo__MinorAudioVideoClass(9)
	QBluetoothDeviceInfo__HiFiAudioDevice               = QBluetoothDeviceInfo__MinorAudioVideoClass(10)
	QBluetoothDeviceInfo__Vcr                           = QBluetoothDeviceInfo__MinorAudioVideoClass(11)
	QBluetoothDeviceInfo__VideoCamera                   = QBluetoothDeviceInfo__MinorAudioVideoClass(12)
	QBluetoothDeviceInfo__Camcorder                     = QBluetoothDeviceInfo__MinorAudioVideoClass(13)
	QBluetoothDeviceInfo__VideoMonitor                  = QBluetoothDeviceInfo__MinorAudioVideoClass(14)
	QBluetoothDeviceInfo__VideoDisplayAndLoudspeaker    = QBluetoothDeviceInfo__MinorAudioVideoClass(15)
	QBluetoothDeviceInfo__VideoConferencing             = QBluetoothDeviceInfo__MinorAudioVideoClass(16)
	QBluetoothDeviceInfo__GamingDevice                  = QBluetoothDeviceInfo__MinorAudioVideoClass(18)
)

//QBluetoothDeviceInfo::MinorComputerClass
type QBluetoothDeviceInfo__MinorComputerClass int64

const (
	QBluetoothDeviceInfo__UncategorizedComputer     = QBluetoothDeviceInfo__MinorComputerClass(0)
	QBluetoothDeviceInfo__DesktopComputer           = QBluetoothDeviceInfo__MinorComputerClass(1)
	QBluetoothDeviceInfo__ServerComputer            = QBluetoothDeviceInfo__MinorComputerClass(2)
	QBluetoothDeviceInfo__LaptopComputer            = QBluetoothDeviceInfo__MinorComputerClass(3)
	QBluetoothDeviceInfo__HandheldClamShellComputer = QBluetoothDeviceInfo__MinorComputerClass(4)
	QBluetoothDeviceInfo__HandheldComputer          = QBluetoothDeviceInfo__MinorComputerClass(5)
	QBluetoothDeviceInfo__WearableComputer          = QBluetoothDeviceInfo__MinorComputerClass(6)
)

//QBluetoothDeviceInfo::MinorHealthClass
type QBluetoothDeviceInfo__MinorHealthClass int64

const (
	QBluetoothDeviceInfo__UncategorizedHealthDevice  = QBluetoothDeviceInfo__MinorHealthClass(0)
	QBluetoothDeviceInfo__HealthBloodPressureMonitor = QBluetoothDeviceInfo__MinorHealthClass(0x1)
	QBluetoothDeviceInfo__HealthThermometer          = QBluetoothDeviceInfo__MinorHealthClass(0x2)
	QBluetoothDeviceInfo__HealthWeightScale          = QBluetoothDeviceInfo__MinorHealthClass(0x3)
	QBluetoothDeviceInfo__HealthGlucoseMeter         = QBluetoothDeviceInfo__MinorHealthClass(0x4)
	QBluetoothDeviceInfo__HealthPulseOximeter        = QBluetoothDeviceInfo__MinorHealthClass(0x5)
	QBluetoothDeviceInfo__HealthDataDisplay          = QBluetoothDeviceInfo__MinorHealthClass(0x7)
	QBluetoothDeviceInfo__HealthStepCounter          = QBluetoothDeviceInfo__MinorHealthClass(0x8)
)

//QBluetoothDeviceInfo::MinorImagingClass
type QBluetoothDeviceInfo__MinorImagingClass int64

const (
	QBluetoothDeviceInfo__UncategorizedImagingDevice = QBluetoothDeviceInfo__MinorImagingClass(0)
	QBluetoothDeviceInfo__ImageDisplay               = QBluetoothDeviceInfo__MinorImagingClass(0x04)
	QBluetoothDeviceInfo__ImageCamera                = QBluetoothDeviceInfo__MinorImagingClass(0x08)
	QBluetoothDeviceInfo__ImageScanner               = QBluetoothDeviceInfo__MinorImagingClass(0x10)
	QBluetoothDeviceInfo__ImagePrinter               = QBluetoothDeviceInfo__MinorImagingClass(0x20)
)

//QBluetoothDeviceInfo::MinorMiscellaneousClass
type QBluetoothDeviceInfo__MinorMiscellaneousClass int64

const (
	QBluetoothDeviceInfo__UncategorizedMiscellaneous = QBluetoothDeviceInfo__MinorMiscellaneousClass(0)
)

//QBluetoothDeviceInfo::MinorNetworkClass
type QBluetoothDeviceInfo__MinorNetworkClass int64

const (
	QBluetoothDeviceInfo__NetworkFullService     = QBluetoothDeviceInfo__MinorNetworkClass(0x00)
	QBluetoothDeviceInfo__NetworkLoadFactorOne   = QBluetoothDeviceInfo__MinorNetworkClass(0x08)
	QBluetoothDeviceInfo__NetworkLoadFactorTwo   = QBluetoothDeviceInfo__MinorNetworkClass(0x10)
	QBluetoothDeviceInfo__NetworkLoadFactorThree = QBluetoothDeviceInfo__MinorNetworkClass(0x18)
	QBluetoothDeviceInfo__NetworkLoadFactorFour  = QBluetoothDeviceInfo__MinorNetworkClass(0x20)
	QBluetoothDeviceInfo__NetworkLoadFactorFive  = QBluetoothDeviceInfo__MinorNetworkClass(0x28)
	QBluetoothDeviceInfo__NetworkLoadFactorSix   = QBluetoothDeviceInfo__MinorNetworkClass(0x30)
	QBluetoothDeviceInfo__NetworkNoService       = QBluetoothDeviceInfo__MinorNetworkClass(0x38)
)

//QBluetoothDeviceInfo::MinorPeripheralClass
type QBluetoothDeviceInfo__MinorPeripheralClass int64

const (
	QBluetoothDeviceInfo__UncategorizedPeripheral              = QBluetoothDeviceInfo__MinorPeripheralClass(0)
	QBluetoothDeviceInfo__KeyboardPeripheral                   = QBluetoothDeviceInfo__MinorPeripheralClass(0x10)
	QBluetoothDeviceInfo__PointingDevicePeripheral             = QBluetoothDeviceInfo__MinorPeripheralClass(0x20)
	QBluetoothDeviceInfo__KeyboardWithPointingDevicePeripheral = QBluetoothDeviceInfo__MinorPeripheralClass(0x30)
	QBluetoothDeviceInfo__JoystickPeripheral                   = QBluetoothDeviceInfo__MinorPeripheralClass(0x01)
	QBluetoothDeviceInfo__GamepadPeripheral                    = QBluetoothDeviceInfo__MinorPeripheralClass(0x02)
	QBluetoothDeviceInfo__RemoteControlPeripheral              = QBluetoothDeviceInfo__MinorPeripheralClass(0x03)
	QBluetoothDeviceInfo__SensingDevicePeripheral              = QBluetoothDeviceInfo__MinorPeripheralClass(0x04)
	QBluetoothDeviceInfo__DigitizerTabletPeripheral            = QBluetoothDeviceInfo__MinorPeripheralClass(0x05)
	QBluetoothDeviceInfo__CardReaderPeripheral                 = QBluetoothDeviceInfo__MinorPeripheralClass(0x06)
)

//QBluetoothDeviceInfo::MinorPhoneClass
type QBluetoothDeviceInfo__MinorPhoneClass int64

const (
	QBluetoothDeviceInfo__UncategorizedPhone            = QBluetoothDeviceInfo__MinorPhoneClass(0)
	QBluetoothDeviceInfo__CellularPhone                 = QBluetoothDeviceInfo__MinorPhoneClass(1)
	QBluetoothDeviceInfo__CordlessPhone                 = QBluetoothDeviceInfo__MinorPhoneClass(2)
	QBluetoothDeviceInfo__SmartPhone                    = QBluetoothDeviceInfo__MinorPhoneClass(3)
	QBluetoothDeviceInfo__WiredModemOrVoiceGatewayPhone = QBluetoothDeviceInfo__MinorPhoneClass(4)
	QBluetoothDeviceInfo__CommonIsdnAccessPhone         = QBluetoothDeviceInfo__MinorPhoneClass(5)
)

//QBluetoothDeviceInfo::MinorToyClass
type QBluetoothDeviceInfo__MinorToyClass int64

const (
	QBluetoothDeviceInfo__UncategorizedToy = QBluetoothDeviceInfo__MinorToyClass(0)
	QBluetoothDeviceInfo__ToyRobot         = QBluetoothDeviceInfo__MinorToyClass(1)
	QBluetoothDeviceInfo__ToyVehicle       = QBluetoothDeviceInfo__MinorToyClass(2)
	QBluetoothDeviceInfo__ToyDoll          = QBluetoothDeviceInfo__MinorToyClass(3)
	QBluetoothDeviceInfo__ToyController    = QBluetoothDeviceInfo__MinorToyClass(4)
	QBluetoothDeviceInfo__ToyGame          = QBluetoothDeviceInfo__MinorToyClass(5)
)

//QBluetoothDeviceInfo::MinorWearableClass
type QBluetoothDeviceInfo__MinorWearableClass int64

const (
	QBluetoothDeviceInfo__UncategorizedWearableDevice = QBluetoothDeviceInfo__MinorWearableClass(0)
	QBluetoothDeviceInfo__WearableWristWatch          = QBluetoothDeviceInfo__MinorWearableClass(1)
	QBluetoothDeviceInfo__WearablePager               = QBluetoothDeviceInfo__MinorWearableClass(2)
	QBluetoothDeviceInfo__WearableJacket              = QBluetoothDeviceInfo__MinorWearableClass(3)
	QBluetoothDeviceInfo__WearableHelmet              = QBluetoothDeviceInfo__MinorWearableClass(4)
	QBluetoothDeviceInfo__WearableGlasses             = QBluetoothDeviceInfo__MinorWearableClass(5)
)

//QBluetoothDeviceInfo::ServiceClass
type QBluetoothDeviceInfo__ServiceClass int64

const (
	QBluetoothDeviceInfo__NoService             = QBluetoothDeviceInfo__ServiceClass(0x0000)
	QBluetoothDeviceInfo__PositioningService    = QBluetoothDeviceInfo__ServiceClass(0x0001)
	QBluetoothDeviceInfo__NetworkingService     = QBluetoothDeviceInfo__ServiceClass(0x0002)
	QBluetoothDeviceInfo__RenderingService      = QBluetoothDeviceInfo__ServiceClass(0x0004)
	QBluetoothDeviceInfo__CapturingService      = QBluetoothDeviceInfo__ServiceClass(0x0008)
	QBluetoothDeviceInfo__ObjectTransferService = QBluetoothDeviceInfo__ServiceClass(0x0010)
	QBluetoothDeviceInfo__AudioService          = QBluetoothDeviceInfo__ServiceClass(0x0020)
	QBluetoothDeviceInfo__TelephonyService      = QBluetoothDeviceInfo__ServiceClass(0x0040)
	QBluetoothDeviceInfo__InformationService    = QBluetoothDeviceInfo__ServiceClass(0x0080)
	QBluetoothDeviceInfo__AllServices           = QBluetoothDeviceInfo__ServiceClass(0x07ff)
)

func NewQBluetoothDeviceInfo() *QBluetoothDeviceInfo {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothDeviceInfo::QBluetoothDeviceInfo")
		}
	}()

	return NewQBluetoothDeviceInfoFromPointer(C.QBluetoothDeviceInfo_NewQBluetoothDeviceInfo())
}

func NewQBluetoothDeviceInfo4(other QBluetoothDeviceInfo_ITF) *QBluetoothDeviceInfo {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothDeviceInfo::QBluetoothDeviceInfo")
		}
	}()

	return NewQBluetoothDeviceInfoFromPointer(C.QBluetoothDeviceInfo_NewQBluetoothDeviceInfo4(PointerFromQBluetoothDeviceInfo(other)))
}

func (ptr *QBluetoothDeviceInfo) CoreConfigurations() QBluetoothDeviceInfo__CoreConfiguration {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothDeviceInfo::coreConfigurations")
		}
	}()

	if ptr.Pointer() != nil {
		return QBluetoothDeviceInfo__CoreConfiguration(C.QBluetoothDeviceInfo_CoreConfigurations(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothDeviceInfo) IsCached() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothDeviceInfo::isCached")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QBluetoothDeviceInfo_IsCached(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothDeviceInfo) IsValid() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothDeviceInfo::isValid")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QBluetoothDeviceInfo_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothDeviceInfo) MajorDeviceClass() QBluetoothDeviceInfo__MajorDeviceClass {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothDeviceInfo::majorDeviceClass")
		}
	}()

	if ptr.Pointer() != nil {
		return QBluetoothDeviceInfo__MajorDeviceClass(C.QBluetoothDeviceInfo_MajorDeviceClass(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothDeviceInfo) Name() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothDeviceInfo::name")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothDeviceInfo_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothDeviceInfo) ServiceClasses() QBluetoothDeviceInfo__ServiceClass {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothDeviceInfo::serviceClasses")
		}
	}()

	if ptr.Pointer() != nil {
		return QBluetoothDeviceInfo__ServiceClass(C.QBluetoothDeviceInfo_ServiceClasses(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothDeviceInfo) ServiceUuidsCompleteness() QBluetoothDeviceInfo__DataCompleteness {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothDeviceInfo::serviceUuidsCompleteness")
		}
	}()

	if ptr.Pointer() != nil {
		return QBluetoothDeviceInfo__DataCompleteness(C.QBluetoothDeviceInfo_ServiceUuidsCompleteness(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothDeviceInfo) SetCached(cached bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothDeviceInfo::setCached")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceInfo_SetCached(ptr.Pointer(), C.int(qt.GoBoolToInt(cached)))
	}
}

func (ptr *QBluetoothDeviceInfo) SetCoreConfigurations(coreConfigs QBluetoothDeviceInfo__CoreConfiguration) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothDeviceInfo::setCoreConfigurations")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceInfo_SetCoreConfigurations(ptr.Pointer(), C.int(coreConfigs))
	}
}

func (ptr *QBluetoothDeviceInfo) SetDeviceUuid(uuid QBluetoothUuid_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothDeviceInfo::setDeviceUuid")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceInfo_SetDeviceUuid(ptr.Pointer(), PointerFromQBluetoothUuid(uuid))
	}
}

func (ptr *QBluetoothDeviceInfo) DestroyQBluetoothDeviceInfo() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothDeviceInfo::~QBluetoothDeviceInfo")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceInfo_DestroyQBluetoothDeviceInfo(ptr.Pointer())
	}
}
