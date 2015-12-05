package bluetooth

//#include "bluetooth.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QBluetoothUuid struct {
	core.QUuid
}

type QBluetoothUuid_ITF interface {
	core.QUuid_ITF
	QBluetoothUuid_PTR() *QBluetoothUuid
}

func PointerFromQBluetoothUuid(ptr QBluetoothUuid_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothUuid_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothUuidFromPointer(ptr unsafe.Pointer) *QBluetoothUuid {
	var n = new(QBluetoothUuid)
	n.SetPointer(ptr)
	return n
}

func (ptr *QBluetoothUuid) QBluetoothUuid_PTR() *QBluetoothUuid {
	return ptr
}

//QBluetoothUuid::CharacteristicType
type QBluetoothUuid__CharacteristicType int64

const (
	QBluetoothUuid__DeviceName                                    = QBluetoothUuid__CharacteristicType(0x2a00)
	QBluetoothUuid__Appearance                                    = QBluetoothUuid__CharacteristicType(0x2a01)
	QBluetoothUuid__PeripheralPrivacyFlag                         = QBluetoothUuid__CharacteristicType(0x2a02)
	QBluetoothUuid__ReconnectionAddress                           = QBluetoothUuid__CharacteristicType(0x2a03)
	QBluetoothUuid__PeripheralPreferredConnectionParameters       = QBluetoothUuid__CharacteristicType(0x2a04)
	QBluetoothUuid__ServiceChanged                                = QBluetoothUuid__CharacteristicType(0x2a05)
	QBluetoothUuid__AlertLevel                                    = QBluetoothUuid__CharacteristicType(0x2a06)
	QBluetoothUuid__TxPowerLevel                                  = QBluetoothUuid__CharacteristicType(0x2a07)
	QBluetoothUuid__DateTime                                      = QBluetoothUuid__CharacteristicType(0x2a08)
	QBluetoothUuid__DayOfWeek                                     = QBluetoothUuid__CharacteristicType(0x2a09)
	QBluetoothUuid__DayDateTime                                   = QBluetoothUuid__CharacteristicType(0x2a0a)
	QBluetoothUuid__ExactTime256                                  = QBluetoothUuid__CharacteristicType(0x2a0c)
	QBluetoothUuid__DSTOffset                                     = QBluetoothUuid__CharacteristicType(0x2a0d)
	QBluetoothUuid__TimeZone                                      = QBluetoothUuid__CharacteristicType(0x2a0e)
	QBluetoothUuid__LocalTimeInformation                          = QBluetoothUuid__CharacteristicType(0x2a0f)
	QBluetoothUuid__TimeWithDST                                   = QBluetoothUuid__CharacteristicType(0x2a11)
	QBluetoothUuid__TimeAccuracy                                  = QBluetoothUuid__CharacteristicType(0x2a12)
	QBluetoothUuid__TimeSource                                    = QBluetoothUuid__CharacteristicType(0x2a13)
	QBluetoothUuid__ReferenceTimeInformation                      = QBluetoothUuid__CharacteristicType(0x2a14)
	QBluetoothUuid__TimeUpdateControlPoint                        = QBluetoothUuid__CharacteristicType(0x2a16)
	QBluetoothUuid__TimeUpdateState                               = QBluetoothUuid__CharacteristicType(0x2a17)
	QBluetoothUuid__GlucoseMeasurement                            = QBluetoothUuid__CharacteristicType(0x2a18)
	QBluetoothUuid__BatteryLevel                                  = QBluetoothUuid__CharacteristicType(0x2a19)
	QBluetoothUuid__TemperatureMeasurement                        = QBluetoothUuid__CharacteristicType(0x2a1c)
	QBluetoothUuid__TemperatureType                               = QBluetoothUuid__CharacteristicType(0x2a1d)
	QBluetoothUuid__IntermediateTemperature                       = QBluetoothUuid__CharacteristicType(0x2a1e)
	QBluetoothUuid__MeasurementInterval                           = QBluetoothUuid__CharacteristicType(0x2a21)
	QBluetoothUuid__BootKeyboardInputReport                       = QBluetoothUuid__CharacteristicType(0x2a22)
	QBluetoothUuid__SystemID                                      = QBluetoothUuid__CharacteristicType(0x2a23)
	QBluetoothUuid__ModelNumberString                             = QBluetoothUuid__CharacteristicType(0x2a24)
	QBluetoothUuid__SerialNumberString                            = QBluetoothUuid__CharacteristicType(0x2a25)
	QBluetoothUuid__FirmwareRevisionString                        = QBluetoothUuid__CharacteristicType(0x2a26)
	QBluetoothUuid__HardwareRevisionString                        = QBluetoothUuid__CharacteristicType(0x2a27)
	QBluetoothUuid__SoftwareRevisionString                        = QBluetoothUuid__CharacteristicType(0x2a28)
	QBluetoothUuid__ManufacturerNameString                        = QBluetoothUuid__CharacteristicType(0x2a29)
	QBluetoothUuid__IEEE1107320601RegulatoryCertificationDataList = QBluetoothUuid__CharacteristicType(0x2a2a)
	QBluetoothUuid__CurrentTime                                   = QBluetoothUuid__CharacteristicType(0x2a2b)
	QBluetoothUuid__MagneticDeclination                           = QBluetoothUuid__CharacteristicType(0x2a2c)
	QBluetoothUuid__ScanRefresh                                   = QBluetoothUuid__CharacteristicType(0x2a31)
	QBluetoothUuid__BootKeyboardOutputReport                      = QBluetoothUuid__CharacteristicType(0x2a32)
	QBluetoothUuid__BootMouseInputReport                          = QBluetoothUuid__CharacteristicType(0x2a33)
	QBluetoothUuid__GlucoseMeasurementContext                     = QBluetoothUuid__CharacteristicType(0x2a34)
	QBluetoothUuid__BloodPressureMeasurement                      = QBluetoothUuid__CharacteristicType(0x2a35)
	QBluetoothUuid__IntermediateCuffPressure                      = QBluetoothUuid__CharacteristicType(0x2a36)
	QBluetoothUuid__HeartRateMeasurement                          = QBluetoothUuid__CharacteristicType(0x2a37)
	QBluetoothUuid__BodySensorLocation                            = QBluetoothUuid__CharacteristicType(0x2a38)
	QBluetoothUuid__HeartRateControlPoint                         = QBluetoothUuid__CharacteristicType(0x2a39)
	QBluetoothUuid__AlertStatus                                   = QBluetoothUuid__CharacteristicType(0x2a3f)
	QBluetoothUuid__RingerControlPoint                            = QBluetoothUuid__CharacteristicType(0x2a40)
	QBluetoothUuid__RingerSetting                                 = QBluetoothUuid__CharacteristicType(0x2a41)
	QBluetoothUuid__AlertCategoryIDBitMask                        = QBluetoothUuid__CharacteristicType(0x2a42)
	QBluetoothUuid__AlertCategoryID                               = QBluetoothUuid__CharacteristicType(0x2a43)
	QBluetoothUuid__AlertNotificationControlPoint                 = QBluetoothUuid__CharacteristicType(0x2a44)
	QBluetoothUuid__UnreadAlertStatus                             = QBluetoothUuid__CharacteristicType(0x2a45)
	QBluetoothUuid__NewAlert                                      = QBluetoothUuid__CharacteristicType(0x2a46)
	QBluetoothUuid__SupportedNewAlertCategory                     = QBluetoothUuid__CharacteristicType(0x2a47)
	QBluetoothUuid__SupportedUnreadAlertCategory                  = QBluetoothUuid__CharacteristicType(0x2a48)
	QBluetoothUuid__BloodPressureFeature                          = QBluetoothUuid__CharacteristicType(0x2a49)
	QBluetoothUuid__HIDInformation                                = QBluetoothUuid__CharacteristicType(0x2a4a)
	QBluetoothUuid__ReportMap                                     = QBluetoothUuid__CharacteristicType(0x2a4b)
	QBluetoothUuid__HIDControlPoint                               = QBluetoothUuid__CharacteristicType(0x2a4c)
	QBluetoothUuid__Report                                        = QBluetoothUuid__CharacteristicType(0x2a4d)
	QBluetoothUuid__ProtocolMode                                  = QBluetoothUuid__CharacteristicType(0x2a4e)
	QBluetoothUuid__ScanIntervalWindow                            = QBluetoothUuid__CharacteristicType(0x2a4f)
	QBluetoothUuid__PnPID                                         = QBluetoothUuid__CharacteristicType(0x2a50)
	QBluetoothUuid__GlucoseFeature                                = QBluetoothUuid__CharacteristicType(0x2a51)
	QBluetoothUuid__RecordAccessControlPoint                      = QBluetoothUuid__CharacteristicType(0x2a52)
	QBluetoothUuid__RSCMeasurement                                = QBluetoothUuid__CharacteristicType(0x2a53)
	QBluetoothUuid__RSCFeature                                    = QBluetoothUuid__CharacteristicType(0x2a54)
	QBluetoothUuid__SCControlPoint                                = QBluetoothUuid__CharacteristicType(0x2a55)
	QBluetoothUuid__CSCMeasurement                                = QBluetoothUuid__CharacteristicType(0x2a5b)
	QBluetoothUuid__CSCFeature                                    = QBluetoothUuid__CharacteristicType(0x2a5c)
	QBluetoothUuid__SensorLocation                                = QBluetoothUuid__CharacteristicType(0x2a5d)
	QBluetoothUuid__CyclingPowerMeasurement                       = QBluetoothUuid__CharacteristicType(0x2a63)
	QBluetoothUuid__CyclingPowerVector                            = QBluetoothUuid__CharacteristicType(0x2a64)
	QBluetoothUuid__CyclingPowerFeature                           = QBluetoothUuid__CharacteristicType(0x2a65)
	QBluetoothUuid__CyclingPowerControlPoint                      = QBluetoothUuid__CharacteristicType(0x2a66)
	QBluetoothUuid__LocationAndSpeed                              = QBluetoothUuid__CharacteristicType(0x2a67)
	QBluetoothUuid__Navigation                                    = QBluetoothUuid__CharacteristicType(0x2a68)
	QBluetoothUuid__PositionQuality                               = QBluetoothUuid__CharacteristicType(0x2a69)
	QBluetoothUuid__LNFeature                                     = QBluetoothUuid__CharacteristicType(0x2a6a)
	QBluetoothUuid__LNControlPoint                                = QBluetoothUuid__CharacteristicType(0x2a6b)
	QBluetoothUuid__Elevation                                     = QBluetoothUuid__CharacteristicType(0x2a6c)
	QBluetoothUuid__Pressure                                      = QBluetoothUuid__CharacteristicType(0x2a6d)
	QBluetoothUuid__Temperature                                   = QBluetoothUuid__CharacteristicType(0x2a6e)
	QBluetoothUuid__Humidity                                      = QBluetoothUuid__CharacteristicType(0x2a6f)
	QBluetoothUuid__TrueWindSpeed                                 = QBluetoothUuid__CharacteristicType(0x2a70)
	QBluetoothUuid__TrueWindDirection                             = QBluetoothUuid__CharacteristicType(0x2a71)
	QBluetoothUuid__ApparentWindSpeed                             = QBluetoothUuid__CharacteristicType(0x2a72)
	QBluetoothUuid__ApparentWindDirection                         = QBluetoothUuid__CharacteristicType(0x2a73)
	QBluetoothUuid__GustFactor                                    = QBluetoothUuid__CharacteristicType(0x2a74)
	QBluetoothUuid__PollenConcentration                           = QBluetoothUuid__CharacteristicType(0x2a75)
	QBluetoothUuid__UVIndex                                       = QBluetoothUuid__CharacteristicType(0x2a76)
	QBluetoothUuid__Irradiance                                    = QBluetoothUuid__CharacteristicType(0x2a77)
	QBluetoothUuid__Rainfall                                      = QBluetoothUuid__CharacteristicType(0x2a78)
	QBluetoothUuid__WindChill                                     = QBluetoothUuid__CharacteristicType(0x2a79)
	QBluetoothUuid__HeatIndex                                     = QBluetoothUuid__CharacteristicType(0x2a7a)
	QBluetoothUuid__DewPoint                                      = QBluetoothUuid__CharacteristicType(0x2a7b)
	QBluetoothUuid__DescriptorValueChanged                        = QBluetoothUuid__CharacteristicType(0x2a7d)
	QBluetoothUuid__AerobicHeartRateLowerLimit                    = QBluetoothUuid__CharacteristicType(0x2a7e)
	QBluetoothUuid__AerobicThreshold                              = QBluetoothUuid__CharacteristicType(0x2a7f)
	QBluetoothUuid__Age                                           = QBluetoothUuid__CharacteristicType(0x2a80)
	QBluetoothUuid__AnaerobicHeartRateLowerLimit                  = QBluetoothUuid__CharacteristicType(0x2a81)
	QBluetoothUuid__AnaerobicHeartRateUpperLimit                  = QBluetoothUuid__CharacteristicType(0x2a82)
	QBluetoothUuid__AnaerobicThreshold                            = QBluetoothUuid__CharacteristicType(0x2a83)
	QBluetoothUuid__AerobicHeartRateUpperLimit                    = QBluetoothUuid__CharacteristicType(0x2a84)
	QBluetoothUuid__DateOfBirth                                   = QBluetoothUuid__CharacteristicType(0x2a85)
	QBluetoothUuid__DateOfThresholdAssessment                     = QBluetoothUuid__CharacteristicType(0x2a86)
	QBluetoothUuid__EmailAddress                                  = QBluetoothUuid__CharacteristicType(0x2a87)
	QBluetoothUuid__FatBurnHeartRateLowerLimit                    = QBluetoothUuid__CharacteristicType(0x2a88)
	QBluetoothUuid__FatBurnHeartRateUpperLimit                    = QBluetoothUuid__CharacteristicType(0x2a89)
	QBluetoothUuid__FirstName                                     = QBluetoothUuid__CharacteristicType(0x2a8a)
	QBluetoothUuid__FiveZoneHeartRateLimits                       = QBluetoothUuid__CharacteristicType(0x2a8b)
	QBluetoothUuid__Gender                                        = QBluetoothUuid__CharacteristicType(0x2a8c)
	QBluetoothUuid__HeartRateMax                                  = QBluetoothUuid__CharacteristicType(0x2a8d)
	QBluetoothUuid__Height                                        = QBluetoothUuid__CharacteristicType(0x2a8e)
	QBluetoothUuid__HipCircumference                              = QBluetoothUuid__CharacteristicType(0x2a8f)
	QBluetoothUuid__LastName                                      = QBluetoothUuid__CharacteristicType(0x2a90)
	QBluetoothUuid__MaximumRecommendedHeartRate                   = QBluetoothUuid__CharacteristicType(0x2a91)
	QBluetoothUuid__RestingHeartRate                              = QBluetoothUuid__CharacteristicType(0x2a92)
	QBluetoothUuid__SportTypeForAerobicAnaerobicThresholds        = QBluetoothUuid__CharacteristicType(0x2a93)
	QBluetoothUuid__ThreeZoneHeartRateLimits                      = QBluetoothUuid__CharacteristicType(0x2a94)
	QBluetoothUuid__TwoZoneHeartRateLimits                        = QBluetoothUuid__CharacteristicType(0x2a95)
	QBluetoothUuid__VO2Max                                        = QBluetoothUuid__CharacteristicType(0x2a96)
	QBluetoothUuid__WaistCircumference                            = QBluetoothUuid__CharacteristicType(0x2a97)
	QBluetoothUuid__Weight                                        = QBluetoothUuid__CharacteristicType(0x2a98)
	QBluetoothUuid__DatabaseChangeIncrement                       = QBluetoothUuid__CharacteristicType(0x2a99)
	QBluetoothUuid__UserIndex                                     = QBluetoothUuid__CharacteristicType(0x2a9a)
	QBluetoothUuid__BodyCompositionFeature                        = QBluetoothUuid__CharacteristicType(0x2a9b)
	QBluetoothUuid__BodyCompositionMeasurement                    = QBluetoothUuid__CharacteristicType(0x2a9c)
	QBluetoothUuid__WeightMeasurement                             = QBluetoothUuid__CharacteristicType(0x2a9d)
	QBluetoothUuid__WeightScaleFeature                            = QBluetoothUuid__CharacteristicType(0x2a9e)
	QBluetoothUuid__UserControlPoint                              = QBluetoothUuid__CharacteristicType(0x2a9f)
	QBluetoothUuid__MagneticFluxDensity2D                         = QBluetoothUuid__CharacteristicType(0x2aa0)
	QBluetoothUuid__MagneticFluxDensity3D                         = QBluetoothUuid__CharacteristicType(0x2aa1)
	QBluetoothUuid__Language                                      = QBluetoothUuid__CharacteristicType(0x2aa2)
	QBluetoothUuid__BarometricPressureTrend                       = QBluetoothUuid__CharacteristicType(0x2aa3)
)

//QBluetoothUuid::DescriptorType
type QBluetoothUuid__DescriptorType int64

const (
	QBluetoothUuid__UnknownDescriptorType              = QBluetoothUuid__DescriptorType(0x0)
	QBluetoothUuid__CharacteristicExtendedProperties   = QBluetoothUuid__DescriptorType(0x2900)
	QBluetoothUuid__CharacteristicUserDescription      = QBluetoothUuid__DescriptorType(0x2901)
	QBluetoothUuid__ClientCharacteristicConfiguration  = QBluetoothUuid__DescriptorType(0x2902)
	QBluetoothUuid__ServerCharacteristicConfiguration  = QBluetoothUuid__DescriptorType(0x2903)
	QBluetoothUuid__CharacteristicPresentationFormat   = QBluetoothUuid__DescriptorType(0x2904)
	QBluetoothUuid__CharacteristicAggregateFormat      = QBluetoothUuid__DescriptorType(0x2905)
	QBluetoothUuid__ValidRange                         = QBluetoothUuid__DescriptorType(0x2906)
	QBluetoothUuid__ExternalReportReference            = QBluetoothUuid__DescriptorType(0x2907)
	QBluetoothUuid__ReportReference                    = QBluetoothUuid__DescriptorType(0x2908)
	QBluetoothUuid__EnvironmentalSensingConfiguration  = QBluetoothUuid__DescriptorType(0x290b)
	QBluetoothUuid__EnvironmentalSensingMeasurement    = QBluetoothUuid__DescriptorType(0x290c)
	QBluetoothUuid__EnvironmentalSensingTriggerSetting = QBluetoothUuid__DescriptorType(0x290d)
)

//QBluetoothUuid::ProtocolUuid
type QBluetoothUuid__ProtocolUuid int64

const (
	QBluetoothUuid__Sdp                    = QBluetoothUuid__ProtocolUuid(0x0001)
	QBluetoothUuid__Udp                    = QBluetoothUuid__ProtocolUuid(0x0002)
	QBluetoothUuid__Rfcomm                 = QBluetoothUuid__ProtocolUuid(0x0003)
	QBluetoothUuid__Tcp                    = QBluetoothUuid__ProtocolUuid(0x0004)
	QBluetoothUuid__TcsBin                 = QBluetoothUuid__ProtocolUuid(0x0005)
	QBluetoothUuid__TcsAt                  = QBluetoothUuid__ProtocolUuid(0x0006)
	QBluetoothUuid__Att                    = QBluetoothUuid__ProtocolUuid(0x0007)
	QBluetoothUuid__Obex                   = QBluetoothUuid__ProtocolUuid(0x0008)
	QBluetoothUuid__Ip                     = QBluetoothUuid__ProtocolUuid(0x0009)
	QBluetoothUuid__Ftp                    = QBluetoothUuid__ProtocolUuid(0x000A)
	QBluetoothUuid__Http                   = QBluetoothUuid__ProtocolUuid(0x000C)
	QBluetoothUuid__Wsp                    = QBluetoothUuid__ProtocolUuid(0x000E)
	QBluetoothUuid__Bnep                   = QBluetoothUuid__ProtocolUuid(0x000F)
	QBluetoothUuid__Upnp                   = QBluetoothUuid__ProtocolUuid(0x0010)
	QBluetoothUuid__Hidp                   = QBluetoothUuid__ProtocolUuid(0x0011)
	QBluetoothUuid__HardcopyControlChannel = QBluetoothUuid__ProtocolUuid(0x0012)
	QBluetoothUuid__HardcopyDataChannel    = QBluetoothUuid__ProtocolUuid(0x0014)
	QBluetoothUuid__HardcopyNotification   = QBluetoothUuid__ProtocolUuid(0x0016)
	QBluetoothUuid__Avctp                  = QBluetoothUuid__ProtocolUuid(0x0017)
	QBluetoothUuid__Avdtp                  = QBluetoothUuid__ProtocolUuid(0x0019)
	QBluetoothUuid__Cmtp                   = QBluetoothUuid__ProtocolUuid(0x001B)
	QBluetoothUuid__UdiCPlain              = QBluetoothUuid__ProtocolUuid(0x001D)
	QBluetoothUuid__McapControlChannel     = QBluetoothUuid__ProtocolUuid(0x001E)
	QBluetoothUuid__McapDataChannel        = QBluetoothUuid__ProtocolUuid(0x001F)
	QBluetoothUuid__L2cap                  = QBluetoothUuid__ProtocolUuid(0x0100)
)

//QBluetoothUuid::ServiceClassUuid
type QBluetoothUuid__ServiceClassUuid int64

const (
	QBluetoothUuid__ServiceDiscoveryServer                = QBluetoothUuid__ServiceClassUuid(0x1000)
	QBluetoothUuid__BrowseGroupDescriptor                 = QBluetoothUuid__ServiceClassUuid(0x1001)
	QBluetoothUuid__PublicBrowseGroup                     = QBluetoothUuid__ServiceClassUuid(0x1002)
	QBluetoothUuid__SerialPort                            = QBluetoothUuid__ServiceClassUuid(0x1101)
	QBluetoothUuid__LANAccessUsingPPP                     = QBluetoothUuid__ServiceClassUuid(0x1102)
	QBluetoothUuid__DialupNetworking                      = QBluetoothUuid__ServiceClassUuid(0x1103)
	QBluetoothUuid__IrMCSync                              = QBluetoothUuid__ServiceClassUuid(0x1104)
	QBluetoothUuid__ObexObjectPush                        = QBluetoothUuid__ServiceClassUuid(0x1105)
	QBluetoothUuid__OBEXFileTransfer                      = QBluetoothUuid__ServiceClassUuid(0x1106)
	QBluetoothUuid__IrMCSyncCommand                       = QBluetoothUuid__ServiceClassUuid(0x1107)
	QBluetoothUuid__Headset                               = QBluetoothUuid__ServiceClassUuid(0x1108)
	QBluetoothUuid__AudioSource                           = QBluetoothUuid__ServiceClassUuid(0x110a)
	QBluetoothUuid__AudioSink                             = QBluetoothUuid__ServiceClassUuid(0x110b)
	QBluetoothUuid__AV_RemoteControlTarget                = QBluetoothUuid__ServiceClassUuid(0x110c)
	QBluetoothUuid__AdvancedAudioDistribution             = QBluetoothUuid__ServiceClassUuid(0x110d)
	QBluetoothUuid__AV_RemoteControl                      = QBluetoothUuid__ServiceClassUuid(0x110e)
	QBluetoothUuid__AV_RemoteControlController            = QBluetoothUuid__ServiceClassUuid(0x110f)
	QBluetoothUuid__HeadsetAG                             = QBluetoothUuid__ServiceClassUuid(0x1112)
	QBluetoothUuid__PANU                                  = QBluetoothUuid__ServiceClassUuid(0x1115)
	QBluetoothUuid__NAP                                   = QBluetoothUuid__ServiceClassUuid(0x1116)
	QBluetoothUuid__GN                                    = QBluetoothUuid__ServiceClassUuid(0x1117)
	QBluetoothUuid__DirectPrinting                        = QBluetoothUuid__ServiceClassUuid(0x1118)
	QBluetoothUuid__ReferencePrinting                     = QBluetoothUuid__ServiceClassUuid(0x1119)
	QBluetoothUuid__BasicImage                            = QBluetoothUuid__ServiceClassUuid(0x111a)
	QBluetoothUuid__ImagingResponder                      = QBluetoothUuid__ServiceClassUuid(0x111b)
	QBluetoothUuid__ImagingAutomaticArchive               = QBluetoothUuid__ServiceClassUuid(0x111c)
	QBluetoothUuid__ImagingReferenceObjects               = QBluetoothUuid__ServiceClassUuid(0x111d)
	QBluetoothUuid__Handsfree                             = QBluetoothUuid__ServiceClassUuid(0x111e)
	QBluetoothUuid__HandsfreeAudioGateway                 = QBluetoothUuid__ServiceClassUuid(0x111f)
	QBluetoothUuid__DirectPrintingReferenceObjectsService = QBluetoothUuid__ServiceClassUuid(0x1120)
	QBluetoothUuid__ReflectedUI                           = QBluetoothUuid__ServiceClassUuid(0x1121)
	QBluetoothUuid__BasicPrinting                         = QBluetoothUuid__ServiceClassUuid(0x1122)
	QBluetoothUuid__PrintingStatus                        = QBluetoothUuid__ServiceClassUuid(0x1123)
	QBluetoothUuid__HumanInterfaceDeviceService           = QBluetoothUuid__ServiceClassUuid(0x1124)
	QBluetoothUuid__HardcopyCableReplacement              = QBluetoothUuid__ServiceClassUuid(0x1125)
	QBluetoothUuid__HCRPrint                              = QBluetoothUuid__ServiceClassUuid(0x1126)
	QBluetoothUuid__HCRScan                               = QBluetoothUuid__ServiceClassUuid(0x1127)
	QBluetoothUuid__SIMAccess                             = QBluetoothUuid__ServiceClassUuid(0x112d)
	QBluetoothUuid__PhonebookAccessPCE                    = QBluetoothUuid__ServiceClassUuid(0x112e)
	QBluetoothUuid__PhonebookAccessPSE                    = QBluetoothUuid__ServiceClassUuid(0x112f)
	QBluetoothUuid__PhonebookAccess                       = QBluetoothUuid__ServiceClassUuid(0x1130)
	QBluetoothUuid__HeadsetHS                             = QBluetoothUuid__ServiceClassUuid(0x1131)
	QBluetoothUuid__MessageAccessServer                   = QBluetoothUuid__ServiceClassUuid(0x1132)
	QBluetoothUuid__MessageNotificationServer             = QBluetoothUuid__ServiceClassUuid(0x1133)
	QBluetoothUuid__MessageAccessProfile                  = QBluetoothUuid__ServiceClassUuid(0x1134)
	QBluetoothUuid__GNSS                                  = QBluetoothUuid__ServiceClassUuid(0x1135)
	QBluetoothUuid__GNSSServer                            = QBluetoothUuid__ServiceClassUuid(0x1136)
	QBluetoothUuid__Display3D                             = QBluetoothUuid__ServiceClassUuid(0x1137)
	QBluetoothUuid__Glasses3D                             = QBluetoothUuid__ServiceClassUuid(0x1138)
	QBluetoothUuid__Synchronization3D                     = QBluetoothUuid__ServiceClassUuid(0x1139)
	QBluetoothUuid__MPSProfile                            = QBluetoothUuid__ServiceClassUuid(0x113a)
	QBluetoothUuid__MPSService                            = QBluetoothUuid__ServiceClassUuid(0x113b)
	QBluetoothUuid__PnPInformation                        = QBluetoothUuid__ServiceClassUuid(0x1200)
	QBluetoothUuid__GenericNetworking                     = QBluetoothUuid__ServiceClassUuid(0x1201)
	QBluetoothUuid__GenericFileTransfer                   = QBluetoothUuid__ServiceClassUuid(0x1202)
	QBluetoothUuid__GenericAudio                          = QBluetoothUuid__ServiceClassUuid(0x1203)
	QBluetoothUuid__GenericTelephony                      = QBluetoothUuid__ServiceClassUuid(0x1204)
	QBluetoothUuid__VideoSource                           = QBluetoothUuid__ServiceClassUuid(0x1303)
	QBluetoothUuid__VideoSink                             = QBluetoothUuid__ServiceClassUuid(0x1304)
	QBluetoothUuid__VideoDistribution                     = QBluetoothUuid__ServiceClassUuid(0x1305)
	QBluetoothUuid__HDP                                   = QBluetoothUuid__ServiceClassUuid(0x1400)
	QBluetoothUuid__HDPSource                             = QBluetoothUuid__ServiceClassUuid(0x1401)
	QBluetoothUuid__HDPSink                               = QBluetoothUuid__ServiceClassUuid(0x1402)
	QBluetoothUuid__GenericAccess                         = QBluetoothUuid__ServiceClassUuid(0x1800)
	QBluetoothUuid__GenericAttribute                      = QBluetoothUuid__ServiceClassUuid(0x1801)
	QBluetoothUuid__ImmediateAlert                        = QBluetoothUuid__ServiceClassUuid(0x1802)
	QBluetoothUuid__LinkLoss                              = QBluetoothUuid__ServiceClassUuid(0x1803)
	QBluetoothUuid__TxPower                               = QBluetoothUuid__ServiceClassUuid(0x1804)
	QBluetoothUuid__CurrentTimeService                    = QBluetoothUuid__ServiceClassUuid(0x1805)
	QBluetoothUuid__ReferenceTimeUpdateService            = QBluetoothUuid__ServiceClassUuid(0x1806)
	QBluetoothUuid__NextDSTChangeService                  = QBluetoothUuid__ServiceClassUuid(0x1807)
	QBluetoothUuid__Glucose                               = QBluetoothUuid__ServiceClassUuid(0x1808)
	QBluetoothUuid__HealthThermometer                     = QBluetoothUuid__ServiceClassUuid(0x1809)
	QBluetoothUuid__DeviceInformation                     = QBluetoothUuid__ServiceClassUuid(0x180a)
	QBluetoothUuid__HeartRate                             = QBluetoothUuid__ServiceClassUuid(0x180d)
	QBluetoothUuid__PhoneAlertStatusService               = QBluetoothUuid__ServiceClassUuid(0x180e)
	QBluetoothUuid__BatteryService                        = QBluetoothUuid__ServiceClassUuid(0x180f)
	QBluetoothUuid__BloodPressure                         = QBluetoothUuid__ServiceClassUuid(0x1810)
	QBluetoothUuid__AlertNotificationService              = QBluetoothUuid__ServiceClassUuid(0x1811)
	QBluetoothUuid__HumanInterfaceDevice                  = QBluetoothUuid__ServiceClassUuid(0x1812)
	QBluetoothUuid__ScanParameters                        = QBluetoothUuid__ServiceClassUuid(0x1813)
	QBluetoothUuid__RunningSpeedAndCadence                = QBluetoothUuid__ServiceClassUuid(0x1814)
	QBluetoothUuid__CyclingSpeedAndCadence                = QBluetoothUuid__ServiceClassUuid(0x1816)
	QBluetoothUuid__CyclingPower                          = QBluetoothUuid__ServiceClassUuid(0x1818)
	QBluetoothUuid__LocationAndNavigation                 = QBluetoothUuid__ServiceClassUuid(0x1819)
	QBluetoothUuid__EnvironmentalSensing                  = QBluetoothUuid__ServiceClassUuid(0x181a)
	QBluetoothUuid__BodyComposition                       = QBluetoothUuid__ServiceClassUuid(0x181b)
	QBluetoothUuid__UserData                              = QBluetoothUuid__ServiceClassUuid(0x181c)
	QBluetoothUuid__WeightScale                           = QBluetoothUuid__ServiceClassUuid(0x181d)
	QBluetoothUuid__BondManagement                        = QBluetoothUuid__ServiceClassUuid(0x181e)
	QBluetoothUuid__ContinuousGlucoseMonitoring           = QBluetoothUuid__ServiceClassUuid(0x181f)
)

func NewQBluetoothUuid() *QBluetoothUuid {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothUuid::QBluetoothUuid")
		}
	}()

	return NewQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid())
}

func NewQBluetoothUuid4(uuid QBluetoothUuid__CharacteristicType) *QBluetoothUuid {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothUuid::QBluetoothUuid")
		}
	}()

	return NewQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid4(C.int(uuid)))
}

func NewQBluetoothUuid5(uuid QBluetoothUuid__DescriptorType) *QBluetoothUuid {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothUuid::QBluetoothUuid")
		}
	}()

	return NewQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid5(C.int(uuid)))
}

func NewQBluetoothUuid2(uuid QBluetoothUuid__ProtocolUuid) *QBluetoothUuid {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothUuid::QBluetoothUuid")
		}
	}()

	return NewQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid2(C.int(uuid)))
}

func NewQBluetoothUuid3(uuid QBluetoothUuid__ServiceClassUuid) *QBluetoothUuid {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothUuid::QBluetoothUuid")
		}
	}()

	return NewQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid3(C.int(uuid)))
}

func NewQBluetoothUuid10(uuid QBluetoothUuid_ITF) *QBluetoothUuid {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothUuid::QBluetoothUuid")
		}
	}()

	return NewQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid10(PointerFromQBluetoothUuid(uuid)))
}

func NewQBluetoothUuid9(uuid string) *QBluetoothUuid {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothUuid::QBluetoothUuid")
		}
	}()

	return NewQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid9(C.CString(uuid)))
}

func NewQBluetoothUuid11(uuid core.QUuid_ITF) *QBluetoothUuid {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothUuid::QBluetoothUuid")
		}
	}()

	return NewQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid11(core.PointerFromQUuid(uuid)))
}

func QBluetoothUuid_CharacteristicToString(uuid QBluetoothUuid__CharacteristicType) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothUuid::characteristicToString")
		}
	}()

	return C.GoString(C.QBluetoothUuid_QBluetoothUuid_CharacteristicToString(C.int(uuid)))
}

func QBluetoothUuid_DescriptorToString(uuid QBluetoothUuid__DescriptorType) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothUuid::descriptorToString")
		}
	}()

	return C.GoString(C.QBluetoothUuid_QBluetoothUuid_DescriptorToString(C.int(uuid)))
}

func (ptr *QBluetoothUuid) MinimumSize() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothUuid::minimumSize")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QBluetoothUuid_MinimumSize(ptr.Pointer()))
	}
	return 0
}

func QBluetoothUuid_ProtocolToString(uuid QBluetoothUuid__ProtocolUuid) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothUuid::protocolToString")
		}
	}()

	return C.GoString(C.QBluetoothUuid_QBluetoothUuid_ProtocolToString(C.int(uuid)))
}

func QBluetoothUuid_ServiceClassToString(uuid QBluetoothUuid__ServiceClassUuid) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothUuid::serviceClassToString")
		}
	}()

	return C.GoString(C.QBluetoothUuid_QBluetoothUuid_ServiceClassToString(C.int(uuid)))
}

func (ptr *QBluetoothUuid) DestroyQBluetoothUuid() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothUuid::~QBluetoothUuid")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBluetoothUuid_DestroyQBluetoothUuid(ptr.Pointer())
	}
}
