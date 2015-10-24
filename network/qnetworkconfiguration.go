package network

//#include "qnetworkconfiguration.h"
import "C"
import (
	"unsafe"
)

type QNetworkConfiguration struct {
	ptr unsafe.Pointer
}

type QNetworkConfigurationITF interface {
	QNetworkConfigurationPTR() *QNetworkConfiguration
}

func (p *QNetworkConfiguration) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QNetworkConfiguration) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQNetworkConfiguration(ptr QNetworkConfigurationITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkConfigurationPTR().Pointer()
	}
	return nil
}

func QNetworkConfigurationFromPointer(ptr unsafe.Pointer) *QNetworkConfiguration {
	var n = new(QNetworkConfiguration)
	n.SetPointer(ptr)
	return n
}

func (ptr *QNetworkConfiguration) QNetworkConfigurationPTR() *QNetworkConfiguration {
	return ptr
}

//QNetworkConfiguration::BearerType
type QNetworkConfiguration__BearerType int

var (
	QNetworkConfiguration__BearerUnknown   = QNetworkConfiguration__BearerType(0)
	QNetworkConfiguration__BearerEthernet  = QNetworkConfiguration__BearerType(1)
	QNetworkConfiguration__BearerWLAN      = QNetworkConfiguration__BearerType(2)
	QNetworkConfiguration__Bearer2G        = QNetworkConfiguration__BearerType(3)
	QNetworkConfiguration__BearerCDMA2000  = QNetworkConfiguration__BearerType(4)
	QNetworkConfiguration__BearerWCDMA     = QNetworkConfiguration__BearerType(5)
	QNetworkConfiguration__BearerHSPA      = QNetworkConfiguration__BearerType(6)
	QNetworkConfiguration__BearerBluetooth = QNetworkConfiguration__BearerType(7)
	QNetworkConfiguration__BearerWiMAX     = QNetworkConfiguration__BearerType(8)
	QNetworkConfiguration__BearerEVDO      = QNetworkConfiguration__BearerType(9)
	QNetworkConfiguration__BearerLTE       = QNetworkConfiguration__BearerType(10)
	QNetworkConfiguration__Bearer3G        = QNetworkConfiguration__BearerType(11)
	QNetworkConfiguration__Bearer4G        = QNetworkConfiguration__BearerType(12)
)

//QNetworkConfiguration::Purpose
type QNetworkConfiguration__Purpose int

var (
	QNetworkConfiguration__UnknownPurpose         = QNetworkConfiguration__Purpose(0)
	QNetworkConfiguration__PublicPurpose          = QNetworkConfiguration__Purpose(1)
	QNetworkConfiguration__PrivatePurpose         = QNetworkConfiguration__Purpose(2)
	QNetworkConfiguration__ServiceSpecificPurpose = QNetworkConfiguration__Purpose(3)
)

//QNetworkConfiguration::StateFlag
type QNetworkConfiguration__StateFlag int

var (
	QNetworkConfiguration__Undefined  = QNetworkConfiguration__StateFlag(0x0000001)
	QNetworkConfiguration__Defined    = QNetworkConfiguration__StateFlag(0x0000002)
	QNetworkConfiguration__Discovered = QNetworkConfiguration__StateFlag(0x0000006)
	QNetworkConfiguration__Active     = QNetworkConfiguration__StateFlag(0x000000e)
)

//QNetworkConfiguration::Type
type QNetworkConfiguration__Type int

var (
	QNetworkConfiguration__InternetAccessPoint = QNetworkConfiguration__Type(0)
	QNetworkConfiguration__ServiceNetwork      = QNetworkConfiguration__Type(1)
	QNetworkConfiguration__UserChoice          = QNetworkConfiguration__Type(2)
	QNetworkConfiguration__Invalid             = QNetworkConfiguration__Type(3)
)

func NewQNetworkConfiguration() *QNetworkConfiguration {
	return QNetworkConfigurationFromPointer(unsafe.Pointer(C.QNetworkConfiguration_NewQNetworkConfiguration()))
}

func NewQNetworkConfiguration2(other QNetworkConfigurationITF) *QNetworkConfiguration {
	return QNetworkConfigurationFromPointer(unsafe.Pointer(C.QNetworkConfiguration_NewQNetworkConfiguration2(C.QtObjectPtr(PointerFromQNetworkConfiguration(other)))))
}

func (ptr *QNetworkConfiguration) BearerType() QNetworkConfiguration__BearerType {
	if ptr.Pointer() != nil {
		return QNetworkConfiguration__BearerType(C.QNetworkConfiguration_BearerType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNetworkConfiguration) BearerTypeFamily() QNetworkConfiguration__BearerType {
	if ptr.Pointer() != nil {
		return QNetworkConfiguration__BearerType(C.QNetworkConfiguration_BearerTypeFamily(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNetworkConfiguration) BearerTypeName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkConfiguration_BearerTypeName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QNetworkConfiguration) Identifier() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkConfiguration_Identifier(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QNetworkConfiguration) IsRoamingAvailable() bool {
	if ptr.Pointer() != nil {
		return C.QNetworkConfiguration_IsRoamingAvailable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNetworkConfiguration) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QNetworkConfiguration_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNetworkConfiguration) Name() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkConfiguration_Name(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QNetworkConfiguration) Purpose() QNetworkConfiguration__Purpose {
	if ptr.Pointer() != nil {
		return QNetworkConfiguration__Purpose(C.QNetworkConfiguration_Purpose(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNetworkConfiguration) Swap(other QNetworkConfigurationITF) {
	if ptr.Pointer() != nil {
		C.QNetworkConfiguration_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQNetworkConfiguration(other)))
	}
}

func (ptr *QNetworkConfiguration) Type() QNetworkConfiguration__Type {
	if ptr.Pointer() != nil {
		return QNetworkConfiguration__Type(C.QNetworkConfiguration_Type(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNetworkConfiguration) DestroyQNetworkConfiguration() {
	if ptr.Pointer() != nil {
		C.QNetworkConfiguration_DestroyQNetworkConfiguration(C.QtObjectPtr(ptr.Pointer()))
	}
}
