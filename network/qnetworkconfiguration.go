package network

//#include "network.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QNetworkConfiguration struct {
	ptr unsafe.Pointer
}

type QNetworkConfiguration_ITF interface {
	QNetworkConfiguration_PTR() *QNetworkConfiguration
}

func (p *QNetworkConfiguration) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QNetworkConfiguration) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQNetworkConfiguration(ptr QNetworkConfiguration_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkConfiguration_PTR().Pointer()
	}
	return nil
}

func NewQNetworkConfigurationFromPointer(ptr unsafe.Pointer) *QNetworkConfiguration {
	var n = new(QNetworkConfiguration)
	n.SetPointer(ptr)
	return n
}

func (ptr *QNetworkConfiguration) QNetworkConfiguration_PTR() *QNetworkConfiguration {
	return ptr
}

//QNetworkConfiguration::BearerType
type QNetworkConfiguration__BearerType int64

const (
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
type QNetworkConfiguration__Purpose int64

const (
	QNetworkConfiguration__UnknownPurpose         = QNetworkConfiguration__Purpose(0)
	QNetworkConfiguration__PublicPurpose          = QNetworkConfiguration__Purpose(1)
	QNetworkConfiguration__PrivatePurpose         = QNetworkConfiguration__Purpose(2)
	QNetworkConfiguration__ServiceSpecificPurpose = QNetworkConfiguration__Purpose(3)
)

//QNetworkConfiguration::StateFlag
type QNetworkConfiguration__StateFlag int64

const (
	QNetworkConfiguration__Undefined  = QNetworkConfiguration__StateFlag(0x0000001)
	QNetworkConfiguration__Defined    = QNetworkConfiguration__StateFlag(0x0000002)
	QNetworkConfiguration__Discovered = QNetworkConfiguration__StateFlag(0x0000006)
	QNetworkConfiguration__Active     = QNetworkConfiguration__StateFlag(0x000000e)
)

//QNetworkConfiguration::Type
type QNetworkConfiguration__Type int64

const (
	QNetworkConfiguration__InternetAccessPoint = QNetworkConfiguration__Type(0)
	QNetworkConfiguration__ServiceNetwork      = QNetworkConfiguration__Type(1)
	QNetworkConfiguration__UserChoice          = QNetworkConfiguration__Type(2)
	QNetworkConfiguration__Invalid             = QNetworkConfiguration__Type(3)
)

func NewQNetworkConfiguration() *QNetworkConfiguration {
	defer qt.Recovering("QNetworkConfiguration::QNetworkConfiguration")

	return NewQNetworkConfigurationFromPointer(C.QNetworkConfiguration_NewQNetworkConfiguration())
}

func NewQNetworkConfiguration2(other QNetworkConfiguration_ITF) *QNetworkConfiguration {
	defer qt.Recovering("QNetworkConfiguration::QNetworkConfiguration")

	return NewQNetworkConfigurationFromPointer(C.QNetworkConfiguration_NewQNetworkConfiguration2(PointerFromQNetworkConfiguration(other)))
}

func (ptr *QNetworkConfiguration) BearerType() QNetworkConfiguration__BearerType {
	defer qt.Recovering("QNetworkConfiguration::bearerType")

	if ptr.Pointer() != nil {
		return QNetworkConfiguration__BearerType(C.QNetworkConfiguration_BearerType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkConfiguration) BearerTypeFamily() QNetworkConfiguration__BearerType {
	defer qt.Recovering("QNetworkConfiguration::bearerTypeFamily")

	if ptr.Pointer() != nil {
		return QNetworkConfiguration__BearerType(C.QNetworkConfiguration_BearerTypeFamily(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkConfiguration) BearerTypeName() string {
	defer qt.Recovering("QNetworkConfiguration::bearerTypeName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkConfiguration_BearerTypeName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkConfiguration) Identifier() string {
	defer qt.Recovering("QNetworkConfiguration::identifier")

	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkConfiguration_Identifier(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkConfiguration) IsRoamingAvailable() bool {
	defer qt.Recovering("QNetworkConfiguration::isRoamingAvailable")

	if ptr.Pointer() != nil {
		return C.QNetworkConfiguration_IsRoamingAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkConfiguration) IsValid() bool {
	defer qt.Recovering("QNetworkConfiguration::isValid")

	if ptr.Pointer() != nil {
		return C.QNetworkConfiguration_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkConfiguration) Name() string {
	defer qt.Recovering("QNetworkConfiguration::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkConfiguration_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkConfiguration) Purpose() QNetworkConfiguration__Purpose {
	defer qt.Recovering("QNetworkConfiguration::purpose")

	if ptr.Pointer() != nil {
		return QNetworkConfiguration__Purpose(C.QNetworkConfiguration_Purpose(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkConfiguration) Swap(other QNetworkConfiguration_ITF) {
	defer qt.Recovering("QNetworkConfiguration::swap")

	if ptr.Pointer() != nil {
		C.QNetworkConfiguration_Swap(ptr.Pointer(), PointerFromQNetworkConfiguration(other))
	}
}

func (ptr *QNetworkConfiguration) Type() QNetworkConfiguration__Type {
	defer qt.Recovering("QNetworkConfiguration::type")

	if ptr.Pointer() != nil {
		return QNetworkConfiguration__Type(C.QNetworkConfiguration_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkConfiguration) DestroyQNetworkConfiguration() {
	defer qt.Recovering("QNetworkConfiguration::~QNetworkConfiguration")

	if ptr.Pointer() != nil {
		C.QNetworkConfiguration_DestroyQNetworkConfiguration(ptr.Pointer())
	}
}
