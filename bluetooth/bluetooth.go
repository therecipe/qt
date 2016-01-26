package bluetooth

//#include "bluetooth.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/network"
	"unsafe"
)

type QBluetoothAddress struct {
	ptr unsafe.Pointer
}

type QBluetoothAddress_ITF interface {
	QBluetoothAddress_PTR() *QBluetoothAddress
}

func (p *QBluetoothAddress) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QBluetoothAddress) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQBluetoothAddress(ptr QBluetoothAddress_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothAddress_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothAddressFromPointer(ptr unsafe.Pointer) *QBluetoothAddress {
	var n = new(QBluetoothAddress)
	n.SetPointer(ptr)
	return n
}

func newQBluetoothAddressFromPointer(ptr unsafe.Pointer) *QBluetoothAddress {
	var n = NewQBluetoothAddressFromPointer(ptr)
	return n
}

func (ptr *QBluetoothAddress) QBluetoothAddress_PTR() *QBluetoothAddress {
	return ptr
}

func NewQBluetoothAddress() *QBluetoothAddress {
	defer qt.Recovering("QBluetoothAddress::QBluetoothAddress")

	return newQBluetoothAddressFromPointer(C.QBluetoothAddress_NewQBluetoothAddress())
}

func NewQBluetoothAddress4(other QBluetoothAddress_ITF) *QBluetoothAddress {
	defer qt.Recovering("QBluetoothAddress::QBluetoothAddress")

	return newQBluetoothAddressFromPointer(C.QBluetoothAddress_NewQBluetoothAddress4(PointerFromQBluetoothAddress(other)))
}

func NewQBluetoothAddress3(address string) *QBluetoothAddress {
	defer qt.Recovering("QBluetoothAddress::QBluetoothAddress")

	return newQBluetoothAddressFromPointer(C.QBluetoothAddress_NewQBluetoothAddress3(C.CString(address)))
}

func (ptr *QBluetoothAddress) Clear() {
	defer qt.Recovering("QBluetoothAddress::clear")

	if ptr.Pointer() != nil {
		C.QBluetoothAddress_Clear(ptr.Pointer())
	}
}

func (ptr *QBluetoothAddress) IsNull() bool {
	defer qt.Recovering("QBluetoothAddress::isNull")

	if ptr.Pointer() != nil {
		return C.QBluetoothAddress_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothAddress) ToString() string {
	defer qt.Recovering("QBluetoothAddress::toString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothAddress_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothAddress) DestroyQBluetoothAddress() {
	defer qt.Recovering("QBluetoothAddress::~QBluetoothAddress")

	if ptr.Pointer() != nil {
		C.QBluetoothAddress_DestroyQBluetoothAddress(ptr.Pointer())
	}
}

type QBluetoothDeviceDiscoveryAgent struct {
	core.QObject
}

type QBluetoothDeviceDiscoveryAgent_ITF interface {
	core.QObject_ITF
	QBluetoothDeviceDiscoveryAgent_PTR() *QBluetoothDeviceDiscoveryAgent
}

func PointerFromQBluetoothDeviceDiscoveryAgent(ptr QBluetoothDeviceDiscoveryAgent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothDeviceDiscoveryAgent_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr unsafe.Pointer) *QBluetoothDeviceDiscoveryAgent {
	var n = new(QBluetoothDeviceDiscoveryAgent)
	n.SetPointer(ptr)
	return n
}

func newQBluetoothDeviceDiscoveryAgentFromPointer(ptr unsafe.Pointer) *QBluetoothDeviceDiscoveryAgent {
	var n = NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr)
	for len(n.ObjectName()) < len("QBluetoothDeviceDiscoveryAgent_") {
		n.SetObjectName("QBluetoothDeviceDiscoveryAgent_" + qt.Identifier())
	}
	return n
}

func (ptr *QBluetoothDeviceDiscoveryAgent) QBluetoothDeviceDiscoveryAgent_PTR() *QBluetoothDeviceDiscoveryAgent {
	return ptr
}

//QBluetoothDeviceDiscoveryAgent::Error
type QBluetoothDeviceDiscoveryAgent__Error int64

const (
	QBluetoothDeviceDiscoveryAgent__NoError                      = QBluetoothDeviceDiscoveryAgent__Error(0)
	QBluetoothDeviceDiscoveryAgent__InputOutputError             = QBluetoothDeviceDiscoveryAgent__Error(1)
	QBluetoothDeviceDiscoveryAgent__PoweredOffError              = QBluetoothDeviceDiscoveryAgent__Error(2)
	QBluetoothDeviceDiscoveryAgent__InvalidBluetoothAdapterError = QBluetoothDeviceDiscoveryAgent__Error(3)
	QBluetoothDeviceDiscoveryAgent__UnsupportedPlatformError     = QBluetoothDeviceDiscoveryAgent__Error(4)
	QBluetoothDeviceDiscoveryAgent__UnknownError                 = QBluetoothDeviceDiscoveryAgent__Error(100)
)

//QBluetoothDeviceDiscoveryAgent::InquiryType
type QBluetoothDeviceDiscoveryAgent__InquiryType int64

const (
	QBluetoothDeviceDiscoveryAgent__GeneralUnlimitedInquiry = QBluetoothDeviceDiscoveryAgent__InquiryType(0)
	QBluetoothDeviceDiscoveryAgent__LimitedInquiry          = QBluetoothDeviceDiscoveryAgent__InquiryType(1)
)

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectCanceled(f func()) {
	defer qt.Recovering("connect QBluetoothDeviceDiscoveryAgent::canceled")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_ConnectCanceled(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "canceled", f)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectCanceled() {
	defer qt.Recovering("disconnect QBluetoothDeviceDiscoveryAgent::canceled")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_DisconnectCanceled(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "canceled")
	}
}

//export callbackQBluetoothDeviceDiscoveryAgentCanceled
func callbackQBluetoothDeviceDiscoveryAgentCanceled(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QBluetoothDeviceDiscoveryAgent::canceled")

	if signal := qt.GetSignal(C.GoString(ptrName), "canceled"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBluetoothDeviceDiscoveryAgent) Canceled() {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::canceled")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_Canceled(ptr.Pointer())
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectError2(f func(error QBluetoothDeviceDiscoveryAgent__Error)) {
	defer qt.Recovering("connect QBluetoothDeviceDiscoveryAgent::error")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error2", f)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectError2() {
	defer qt.Recovering("disconnect QBluetoothDeviceDiscoveryAgent::error")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error2")
	}
}

//export callbackQBluetoothDeviceDiscoveryAgentError2
func callbackQBluetoothDeviceDiscoveryAgentError2(ptr unsafe.Pointer, ptrName *C.char, error C.int) {
	defer qt.Recovering("callback QBluetoothDeviceDiscoveryAgent::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(QBluetoothDeviceDiscoveryAgent__Error))(QBluetoothDeviceDiscoveryAgent__Error(error))
	}

}

func (ptr *QBluetoothDeviceDiscoveryAgent) Error2(error QBluetoothDeviceDiscoveryAgent__Error) {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::error")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_Error2(ptr.Pointer(), C.int(error))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectFinished(f func()) {
	defer qt.Recovering("connect QBluetoothDeviceDiscoveryAgent::finished")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectFinished() {
	defer qt.Recovering("disconnect QBluetoothDeviceDiscoveryAgent::finished")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQBluetoothDeviceDiscoveryAgentFinished
func callbackQBluetoothDeviceDiscoveryAgentFinished(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QBluetoothDeviceDiscoveryAgent::finished")

	if signal := qt.GetSignal(C.GoString(ptrName), "finished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBluetoothDeviceDiscoveryAgent) Finished() {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::finished")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_Finished(ptr.Pointer())
	}
}

func NewQBluetoothDeviceDiscoveryAgent(parent core.QObject_ITF) *QBluetoothDeviceDiscoveryAgent {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::QBluetoothDeviceDiscoveryAgent")

	return newQBluetoothDeviceDiscoveryAgentFromPointer(C.QBluetoothDeviceDiscoveryAgent_NewQBluetoothDeviceDiscoveryAgent(core.PointerFromQObject(parent)))
}

func NewQBluetoothDeviceDiscoveryAgent2(deviceAdapter QBluetoothAddress_ITF, parent core.QObject_ITF) *QBluetoothDeviceDiscoveryAgent {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::QBluetoothDeviceDiscoveryAgent")

	return newQBluetoothDeviceDiscoveryAgentFromPointer(C.QBluetoothDeviceDiscoveryAgent_NewQBluetoothDeviceDiscoveryAgent2(PointerFromQBluetoothAddress(deviceAdapter), core.PointerFromQObject(parent)))
}

func (ptr *QBluetoothDeviceDiscoveryAgent) Error() QBluetoothDeviceDiscoveryAgent__Error {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::error")

	if ptr.Pointer() != nil {
		return QBluetoothDeviceDiscoveryAgent__Error(C.QBluetoothDeviceDiscoveryAgent_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ErrorString() string {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothDeviceDiscoveryAgent_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothDeviceDiscoveryAgent) InquiryType() QBluetoothDeviceDiscoveryAgent__InquiryType {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::inquiryType")

	if ptr.Pointer() != nil {
		return QBluetoothDeviceDiscoveryAgent__InquiryType(C.QBluetoothDeviceDiscoveryAgent_InquiryType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothDeviceDiscoveryAgent) IsActive() bool {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::isActive")

	if ptr.Pointer() != nil {
		return C.QBluetoothDeviceDiscoveryAgent_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothDeviceDiscoveryAgent) SetInquiryType(ty QBluetoothDeviceDiscoveryAgent__InquiryType) {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::setInquiryType")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_SetInquiryType(ptr.Pointer(), C.int(ty))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) Start() {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::start")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_Start(ptr.Pointer())
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) Stop() {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::stop")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_Stop(ptr.Pointer())
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DestroyQBluetoothDeviceDiscoveryAgent() {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::~QBluetoothDeviceDiscoveryAgent")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_DestroyQBluetoothDeviceDiscoveryAgent(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QBluetoothDeviceDiscoveryAgent::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QBluetoothDeviceDiscoveryAgent::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQBluetoothDeviceDiscoveryAgentTimerEvent
func callbackQBluetoothDeviceDiscoveryAgentTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothDeviceDiscoveryAgent::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::timerEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::timerEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QBluetoothDeviceDiscoveryAgent::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QBluetoothDeviceDiscoveryAgent::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQBluetoothDeviceDiscoveryAgentChildEvent
func callbackQBluetoothDeviceDiscoveryAgentChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothDeviceDiscoveryAgent::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::childEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::childEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QBluetoothDeviceDiscoveryAgent::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QBluetoothDeviceDiscoveryAgent::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQBluetoothDeviceDiscoveryAgentCustomEvent
func callbackQBluetoothDeviceDiscoveryAgentCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothDeviceDiscoveryAgent::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::customEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::customEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

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

func newQBluetoothDeviceInfoFromPointer(ptr unsafe.Pointer) *QBluetoothDeviceInfo {
	var n = NewQBluetoothDeviceInfoFromPointer(ptr)
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
	defer qt.Recovering("QBluetoothDeviceInfo::QBluetoothDeviceInfo")

	return newQBluetoothDeviceInfoFromPointer(C.QBluetoothDeviceInfo_NewQBluetoothDeviceInfo())
}

func NewQBluetoothDeviceInfo4(other QBluetoothDeviceInfo_ITF) *QBluetoothDeviceInfo {
	defer qt.Recovering("QBluetoothDeviceInfo::QBluetoothDeviceInfo")

	return newQBluetoothDeviceInfoFromPointer(C.QBluetoothDeviceInfo_NewQBluetoothDeviceInfo4(PointerFromQBluetoothDeviceInfo(other)))
}

func (ptr *QBluetoothDeviceInfo) CoreConfigurations() QBluetoothDeviceInfo__CoreConfiguration {
	defer qt.Recovering("QBluetoothDeviceInfo::coreConfigurations")

	if ptr.Pointer() != nil {
		return QBluetoothDeviceInfo__CoreConfiguration(C.QBluetoothDeviceInfo_CoreConfigurations(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothDeviceInfo) IsCached() bool {
	defer qt.Recovering("QBluetoothDeviceInfo::isCached")

	if ptr.Pointer() != nil {
		return C.QBluetoothDeviceInfo_IsCached(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothDeviceInfo) IsValid() bool {
	defer qt.Recovering("QBluetoothDeviceInfo::isValid")

	if ptr.Pointer() != nil {
		return C.QBluetoothDeviceInfo_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothDeviceInfo) MajorDeviceClass() QBluetoothDeviceInfo__MajorDeviceClass {
	defer qt.Recovering("QBluetoothDeviceInfo::majorDeviceClass")

	if ptr.Pointer() != nil {
		return QBluetoothDeviceInfo__MajorDeviceClass(C.QBluetoothDeviceInfo_MajorDeviceClass(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothDeviceInfo) Name() string {
	defer qt.Recovering("QBluetoothDeviceInfo::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothDeviceInfo_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothDeviceInfo) ServiceClasses() QBluetoothDeviceInfo__ServiceClass {
	defer qt.Recovering("QBluetoothDeviceInfo::serviceClasses")

	if ptr.Pointer() != nil {
		return QBluetoothDeviceInfo__ServiceClass(C.QBluetoothDeviceInfo_ServiceClasses(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothDeviceInfo) ServiceUuidsCompleteness() QBluetoothDeviceInfo__DataCompleteness {
	defer qt.Recovering("QBluetoothDeviceInfo::serviceUuidsCompleteness")

	if ptr.Pointer() != nil {
		return QBluetoothDeviceInfo__DataCompleteness(C.QBluetoothDeviceInfo_ServiceUuidsCompleteness(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothDeviceInfo) SetCached(cached bool) {
	defer qt.Recovering("QBluetoothDeviceInfo::setCached")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceInfo_SetCached(ptr.Pointer(), C.int(qt.GoBoolToInt(cached)))
	}
}

func (ptr *QBluetoothDeviceInfo) SetCoreConfigurations(coreConfigs QBluetoothDeviceInfo__CoreConfiguration) {
	defer qt.Recovering("QBluetoothDeviceInfo::setCoreConfigurations")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceInfo_SetCoreConfigurations(ptr.Pointer(), C.int(coreConfigs))
	}
}

func (ptr *QBluetoothDeviceInfo) SetDeviceUuid(uuid QBluetoothUuid_ITF) {
	defer qt.Recovering("QBluetoothDeviceInfo::setDeviceUuid")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceInfo_SetDeviceUuid(ptr.Pointer(), PointerFromQBluetoothUuid(uuid))
	}
}

func (ptr *QBluetoothDeviceInfo) DestroyQBluetoothDeviceInfo() {
	defer qt.Recovering("QBluetoothDeviceInfo::~QBluetoothDeviceInfo")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceInfo_DestroyQBluetoothDeviceInfo(ptr.Pointer())
	}
}

type QBluetoothHostInfo struct {
	ptr unsafe.Pointer
}

type QBluetoothHostInfo_ITF interface {
	QBluetoothHostInfo_PTR() *QBluetoothHostInfo
}

func (p *QBluetoothHostInfo) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QBluetoothHostInfo) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQBluetoothHostInfo(ptr QBluetoothHostInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothHostInfo_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothHostInfoFromPointer(ptr unsafe.Pointer) *QBluetoothHostInfo {
	var n = new(QBluetoothHostInfo)
	n.SetPointer(ptr)
	return n
}

func newQBluetoothHostInfoFromPointer(ptr unsafe.Pointer) *QBluetoothHostInfo {
	var n = NewQBluetoothHostInfoFromPointer(ptr)
	return n
}

func (ptr *QBluetoothHostInfo) QBluetoothHostInfo_PTR() *QBluetoothHostInfo {
	return ptr
}

func NewQBluetoothHostInfo() *QBluetoothHostInfo {
	defer qt.Recovering("QBluetoothHostInfo::QBluetoothHostInfo")

	return newQBluetoothHostInfoFromPointer(C.QBluetoothHostInfo_NewQBluetoothHostInfo())
}

func NewQBluetoothHostInfo2(other QBluetoothHostInfo_ITF) *QBluetoothHostInfo {
	defer qt.Recovering("QBluetoothHostInfo::QBluetoothHostInfo")

	return newQBluetoothHostInfoFromPointer(C.QBluetoothHostInfo_NewQBluetoothHostInfo2(PointerFromQBluetoothHostInfo(other)))
}

func (ptr *QBluetoothHostInfo) Name() string {
	defer qt.Recovering("QBluetoothHostInfo::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothHostInfo_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothHostInfo) SetAddress(address QBluetoothAddress_ITF) {
	defer qt.Recovering("QBluetoothHostInfo::setAddress")

	if ptr.Pointer() != nil {
		C.QBluetoothHostInfo_SetAddress(ptr.Pointer(), PointerFromQBluetoothAddress(address))
	}
}

func (ptr *QBluetoothHostInfo) SetName(name string) {
	defer qt.Recovering("QBluetoothHostInfo::setName")

	if ptr.Pointer() != nil {
		C.QBluetoothHostInfo_SetName(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QBluetoothHostInfo) DestroyQBluetoothHostInfo() {
	defer qt.Recovering("QBluetoothHostInfo::~QBluetoothHostInfo")

	if ptr.Pointer() != nil {
		C.QBluetoothHostInfo_DestroyQBluetoothHostInfo(ptr.Pointer())
	}
}

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
	return n
}

func newQBluetoothLocalDeviceFromPointer(ptr unsafe.Pointer) *QBluetoothLocalDevice {
	var n = NewQBluetoothLocalDeviceFromPointer(ptr)
	for len(n.ObjectName()) < len("QBluetoothLocalDevice_") {
		n.SetObjectName("QBluetoothLocalDevice_" + qt.Identifier())
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
	defer qt.Recovering("connect QBluetoothLocalDevice::error")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_ConnectError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error", f)
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectError() {
	defer qt.Recovering("disconnect QBluetoothLocalDevice::error")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DisconnectError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error")
	}
}

//export callbackQBluetoothLocalDeviceError
func callbackQBluetoothLocalDeviceError(ptr unsafe.Pointer, ptrName *C.char, error C.int) {
	defer qt.Recovering("callback QBluetoothLocalDevice::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error"); signal != nil {
		signal.(func(QBluetoothLocalDevice__Error))(QBluetoothLocalDevice__Error(error))
	}

}

func (ptr *QBluetoothLocalDevice) Error(error QBluetoothLocalDevice__Error) {
	defer qt.Recovering("QBluetoothLocalDevice::error")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_Error(ptr.Pointer(), C.int(error))
	}
}

func (ptr *QBluetoothLocalDevice) ConnectHostModeStateChanged(f func(state QBluetoothLocalDevice__HostMode)) {
	defer qt.Recovering("connect QBluetoothLocalDevice::hostModeStateChanged")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_ConnectHostModeStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "hostModeStateChanged", f)
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectHostModeStateChanged() {
	defer qt.Recovering("disconnect QBluetoothLocalDevice::hostModeStateChanged")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DisconnectHostModeStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "hostModeStateChanged")
	}
}

//export callbackQBluetoothLocalDeviceHostModeStateChanged
func callbackQBluetoothLocalDeviceHostModeStateChanged(ptr unsafe.Pointer, ptrName *C.char, state C.int) {
	defer qt.Recovering("callback QBluetoothLocalDevice::hostModeStateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "hostModeStateChanged"); signal != nil {
		signal.(func(QBluetoothLocalDevice__HostMode))(QBluetoothLocalDevice__HostMode(state))
	}

}

func (ptr *QBluetoothLocalDevice) HostModeStateChanged(state QBluetoothLocalDevice__HostMode) {
	defer qt.Recovering("QBluetoothLocalDevice::hostModeStateChanged")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_HostModeStateChanged(ptr.Pointer(), C.int(state))
	}
}

func (ptr *QBluetoothLocalDevice) IsValid() bool {
	defer qt.Recovering("QBluetoothLocalDevice::isValid")

	if ptr.Pointer() != nil {
		return C.QBluetoothLocalDevice_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothLocalDevice) DestroyQBluetoothLocalDevice() {
	defer qt.Recovering("QBluetoothLocalDevice::~QBluetoothLocalDevice")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DestroyQBluetoothLocalDevice(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func NewQBluetoothLocalDevice(parent core.QObject_ITF) *QBluetoothLocalDevice {
	defer qt.Recovering("QBluetoothLocalDevice::QBluetoothLocalDevice")

	return newQBluetoothLocalDeviceFromPointer(C.QBluetoothLocalDevice_NewQBluetoothLocalDevice(core.PointerFromQObject(parent)))
}

func NewQBluetoothLocalDevice2(address QBluetoothAddress_ITF, parent core.QObject_ITF) *QBluetoothLocalDevice {
	defer qt.Recovering("QBluetoothLocalDevice::QBluetoothLocalDevice")

	return newQBluetoothLocalDeviceFromPointer(C.QBluetoothLocalDevice_NewQBluetoothLocalDevice2(PointerFromQBluetoothAddress(address), core.PointerFromQObject(parent)))
}

func (ptr *QBluetoothLocalDevice) HostMode() QBluetoothLocalDevice__HostMode {
	defer qt.Recovering("QBluetoothLocalDevice::hostMode")

	if ptr.Pointer() != nil {
		return QBluetoothLocalDevice__HostMode(C.QBluetoothLocalDevice_HostMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothLocalDevice) Name() string {
	defer qt.Recovering("QBluetoothLocalDevice::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothLocalDevice_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothLocalDevice) PairingConfirmation(accept bool) {
	defer qt.Recovering("QBluetoothLocalDevice::pairingConfirmation")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_PairingConfirmation(ptr.Pointer(), C.int(qt.GoBoolToInt(accept)))
	}
}

func (ptr *QBluetoothLocalDevice) PairingStatus(address QBluetoothAddress_ITF) QBluetoothLocalDevice__Pairing {
	defer qt.Recovering("QBluetoothLocalDevice::pairingStatus")

	if ptr.Pointer() != nil {
		return QBluetoothLocalDevice__Pairing(C.QBluetoothLocalDevice_PairingStatus(ptr.Pointer(), PointerFromQBluetoothAddress(address)))
	}
	return 0
}

func (ptr *QBluetoothLocalDevice) PowerOn() {
	defer qt.Recovering("QBluetoothLocalDevice::powerOn")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_PowerOn(ptr.Pointer())
	}
}

func (ptr *QBluetoothLocalDevice) RequestPairing(address QBluetoothAddress_ITF, pairing QBluetoothLocalDevice__Pairing) {
	defer qt.Recovering("QBluetoothLocalDevice::requestPairing")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_RequestPairing(ptr.Pointer(), PointerFromQBluetoothAddress(address), C.int(pairing))
	}
}

func (ptr *QBluetoothLocalDevice) SetHostMode(mode QBluetoothLocalDevice__HostMode) {
	defer qt.Recovering("QBluetoothLocalDevice::setHostMode")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_SetHostMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QBluetoothLocalDevice) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QBluetoothLocalDevice::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QBluetoothLocalDevice::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQBluetoothLocalDeviceTimerEvent
func callbackQBluetoothLocalDeviceTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothLocalDevice::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQBluetoothLocalDeviceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QBluetoothLocalDevice) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QBluetoothLocalDevice::timerEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QBluetoothLocalDevice) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QBluetoothLocalDevice::timerEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QBluetoothLocalDevice) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QBluetoothLocalDevice::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QBluetoothLocalDevice::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQBluetoothLocalDeviceChildEvent
func callbackQBluetoothLocalDeviceChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothLocalDevice::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQBluetoothLocalDeviceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QBluetoothLocalDevice) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QBluetoothLocalDevice::childEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QBluetoothLocalDevice) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QBluetoothLocalDevice::childEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QBluetoothLocalDevice) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QBluetoothLocalDevice::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QBluetoothLocalDevice::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQBluetoothLocalDeviceCustomEvent
func callbackQBluetoothLocalDeviceCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothLocalDevice::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQBluetoothLocalDeviceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QBluetoothLocalDevice) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QBluetoothLocalDevice::customEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QBluetoothLocalDevice) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QBluetoothLocalDevice::customEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QBluetoothServer struct {
	core.QObject
}

type QBluetoothServer_ITF interface {
	core.QObject_ITF
	QBluetoothServer_PTR() *QBluetoothServer
}

func PointerFromQBluetoothServer(ptr QBluetoothServer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothServer_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothServerFromPointer(ptr unsafe.Pointer) *QBluetoothServer {
	var n = new(QBluetoothServer)
	n.SetPointer(ptr)
	return n
}

func newQBluetoothServerFromPointer(ptr unsafe.Pointer) *QBluetoothServer {
	var n = NewQBluetoothServerFromPointer(ptr)
	for len(n.ObjectName()) < len("QBluetoothServer_") {
		n.SetObjectName("QBluetoothServer_" + qt.Identifier())
	}
	return n
}

func (ptr *QBluetoothServer) QBluetoothServer_PTR() *QBluetoothServer {
	return ptr
}

//QBluetoothServer::Error
type QBluetoothServer__Error int64

const (
	QBluetoothServer__NoError                       = QBluetoothServer__Error(0)
	QBluetoothServer__UnknownError                  = QBluetoothServer__Error(1)
	QBluetoothServer__PoweredOffError               = QBluetoothServer__Error(2)
	QBluetoothServer__InputOutputError              = QBluetoothServer__Error(3)
	QBluetoothServer__ServiceAlreadyRegisteredError = QBluetoothServer__Error(4)
	QBluetoothServer__UnsupportedProtocolError      = QBluetoothServer__Error(5)
)

func NewQBluetoothServer(serverType QBluetoothServiceInfo__Protocol, parent core.QObject_ITF) *QBluetoothServer {
	defer qt.Recovering("QBluetoothServer::QBluetoothServer")

	return newQBluetoothServerFromPointer(C.QBluetoothServer_NewQBluetoothServer(C.int(serverType), core.PointerFromQObject(parent)))
}

func (ptr *QBluetoothServer) ConnectError2(f func(error QBluetoothServer__Error)) {
	defer qt.Recovering("connect QBluetoothServer::error")

	if ptr.Pointer() != nil {
		C.QBluetoothServer_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error2", f)
	}
}

func (ptr *QBluetoothServer) DisconnectError2() {
	defer qt.Recovering("disconnect QBluetoothServer::error")

	if ptr.Pointer() != nil {
		C.QBluetoothServer_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error2")
	}
}

//export callbackQBluetoothServerError2
func callbackQBluetoothServerError2(ptr unsafe.Pointer, ptrName *C.char, error C.int) {
	defer qt.Recovering("callback QBluetoothServer::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(QBluetoothServer__Error))(QBluetoothServer__Error(error))
	}

}

func (ptr *QBluetoothServer) Error2(error QBluetoothServer__Error) {
	defer qt.Recovering("QBluetoothServer::error")

	if ptr.Pointer() != nil {
		C.QBluetoothServer_Error2(ptr.Pointer(), C.int(error))
	}
}

func (ptr *QBluetoothServer) ConnectNewConnection(f func()) {
	defer qt.Recovering("connect QBluetoothServer::newConnection")

	if ptr.Pointer() != nil {
		C.QBluetoothServer_ConnectNewConnection(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "newConnection", f)
	}
}

func (ptr *QBluetoothServer) DisconnectNewConnection() {
	defer qt.Recovering("disconnect QBluetoothServer::newConnection")

	if ptr.Pointer() != nil {
		C.QBluetoothServer_DisconnectNewConnection(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "newConnection")
	}
}

//export callbackQBluetoothServerNewConnection
func callbackQBluetoothServerNewConnection(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QBluetoothServer::newConnection")

	if signal := qt.GetSignal(C.GoString(ptrName), "newConnection"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBluetoothServer) NewConnection() {
	defer qt.Recovering("QBluetoothServer::newConnection")

	if ptr.Pointer() != nil {
		C.QBluetoothServer_NewConnection(ptr.Pointer())
	}
}

func (ptr *QBluetoothServer) Error() QBluetoothServer__Error {
	defer qt.Recovering("QBluetoothServer::error")

	if ptr.Pointer() != nil {
		return QBluetoothServer__Error(C.QBluetoothServer_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothServer) IsListening() bool {
	defer qt.Recovering("QBluetoothServer::isListening")

	if ptr.Pointer() != nil {
		return C.QBluetoothServer_IsListening(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothServer) MaxPendingConnections() int {
	defer qt.Recovering("QBluetoothServer::maxPendingConnections")

	if ptr.Pointer() != nil {
		return int(C.QBluetoothServer_MaxPendingConnections(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothServer) ServerType() QBluetoothServiceInfo__Protocol {
	defer qt.Recovering("QBluetoothServer::serverType")

	if ptr.Pointer() != nil {
		return QBluetoothServiceInfo__Protocol(C.QBluetoothServer_ServerType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothServer) DestroyQBluetoothServer() {
	defer qt.Recovering("QBluetoothServer::~QBluetoothServer")

	if ptr.Pointer() != nil {
		C.QBluetoothServer_DestroyQBluetoothServer(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothServer) Close() {
	defer qt.Recovering("QBluetoothServer::close")

	if ptr.Pointer() != nil {
		C.QBluetoothServer_Close(ptr.Pointer())
	}
}

func (ptr *QBluetoothServer) HasPendingConnections() bool {
	defer qt.Recovering("QBluetoothServer::hasPendingConnections")

	if ptr.Pointer() != nil {
		return C.QBluetoothServer_HasPendingConnections(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothServer) NextPendingConnection() *QBluetoothSocket {
	defer qt.Recovering("QBluetoothServer::nextPendingConnection")

	if ptr.Pointer() != nil {
		return NewQBluetoothSocketFromPointer(C.QBluetoothServer_NextPendingConnection(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBluetoothServer) SetMaxPendingConnections(numConnections int) {
	defer qt.Recovering("QBluetoothServer::setMaxPendingConnections")

	if ptr.Pointer() != nil {
		C.QBluetoothServer_SetMaxPendingConnections(ptr.Pointer(), C.int(numConnections))
	}
}

func (ptr *QBluetoothServer) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QBluetoothServer::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QBluetoothServer) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QBluetoothServer::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQBluetoothServerTimerEvent
func callbackQBluetoothServerTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothServer::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQBluetoothServerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QBluetoothServer) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QBluetoothServer::timerEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothServer_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QBluetoothServer) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QBluetoothServer::timerEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothServer_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QBluetoothServer) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QBluetoothServer::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QBluetoothServer) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QBluetoothServer::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQBluetoothServerChildEvent
func callbackQBluetoothServerChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothServer::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQBluetoothServerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QBluetoothServer) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QBluetoothServer::childEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothServer_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QBluetoothServer) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QBluetoothServer::childEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothServer_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QBluetoothServer) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QBluetoothServer::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QBluetoothServer) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QBluetoothServer::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQBluetoothServerCustomEvent
func callbackQBluetoothServerCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothServer::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQBluetoothServerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QBluetoothServer) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QBluetoothServer::customEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothServer_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QBluetoothServer) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QBluetoothServer::customEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothServer_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QBluetoothServiceDiscoveryAgent struct {
	core.QObject
}

type QBluetoothServiceDiscoveryAgent_ITF interface {
	core.QObject_ITF
	QBluetoothServiceDiscoveryAgent_PTR() *QBluetoothServiceDiscoveryAgent
}

func PointerFromQBluetoothServiceDiscoveryAgent(ptr QBluetoothServiceDiscoveryAgent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothServiceDiscoveryAgent_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothServiceDiscoveryAgentFromPointer(ptr unsafe.Pointer) *QBluetoothServiceDiscoveryAgent {
	var n = new(QBluetoothServiceDiscoveryAgent)
	n.SetPointer(ptr)
	return n
}

func newQBluetoothServiceDiscoveryAgentFromPointer(ptr unsafe.Pointer) *QBluetoothServiceDiscoveryAgent {
	var n = NewQBluetoothServiceDiscoveryAgentFromPointer(ptr)
	for len(n.ObjectName()) < len("QBluetoothServiceDiscoveryAgent_") {
		n.SetObjectName("QBluetoothServiceDiscoveryAgent_" + qt.Identifier())
	}
	return n
}

func (ptr *QBluetoothServiceDiscoveryAgent) QBluetoothServiceDiscoveryAgent_PTR() *QBluetoothServiceDiscoveryAgent {
	return ptr
}

//QBluetoothServiceDiscoveryAgent::DiscoveryMode
type QBluetoothServiceDiscoveryAgent__DiscoveryMode int64

const (
	QBluetoothServiceDiscoveryAgent__MinimalDiscovery = QBluetoothServiceDiscoveryAgent__DiscoveryMode(0)
	QBluetoothServiceDiscoveryAgent__FullDiscovery    = QBluetoothServiceDiscoveryAgent__DiscoveryMode(1)
)

//QBluetoothServiceDiscoveryAgent::Error
type QBluetoothServiceDiscoveryAgent__Error int64

const (
	QBluetoothServiceDiscoveryAgent__NoError                      = QBluetoothServiceDiscoveryAgent__Error(QBluetoothDeviceDiscoveryAgent__NoError)
	QBluetoothServiceDiscoveryAgent__InputOutputError             = QBluetoothServiceDiscoveryAgent__Error(QBluetoothDeviceDiscoveryAgent__InputOutputError)
	QBluetoothServiceDiscoveryAgent__PoweredOffError              = QBluetoothServiceDiscoveryAgent__Error(QBluetoothDeviceDiscoveryAgent__PoweredOffError)
	QBluetoothServiceDiscoveryAgent__InvalidBluetoothAdapterError = QBluetoothServiceDiscoveryAgent__Error(QBluetoothDeviceDiscoveryAgent__InvalidBluetoothAdapterError)
	QBluetoothServiceDiscoveryAgent__UnknownError                 = QBluetoothServiceDiscoveryAgent__Error(QBluetoothDeviceDiscoveryAgent__UnknownError)
)

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectCanceled(f func()) {
	defer qt.Recovering("connect QBluetoothServiceDiscoveryAgent::canceled")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_ConnectCanceled(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "canceled", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectCanceled() {
	defer qt.Recovering("disconnect QBluetoothServiceDiscoveryAgent::canceled")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_DisconnectCanceled(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "canceled")
	}
}

//export callbackQBluetoothServiceDiscoveryAgentCanceled
func callbackQBluetoothServiceDiscoveryAgentCanceled(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QBluetoothServiceDiscoveryAgent::canceled")

	if signal := qt.GetSignal(C.GoString(ptrName), "canceled"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBluetoothServiceDiscoveryAgent) Canceled() {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::canceled")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_Canceled(ptr.Pointer())
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectError2(f func(error QBluetoothServiceDiscoveryAgent__Error)) {
	defer qt.Recovering("connect QBluetoothServiceDiscoveryAgent::error")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error2", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectError2() {
	defer qt.Recovering("disconnect QBluetoothServiceDiscoveryAgent::error")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error2")
	}
}

//export callbackQBluetoothServiceDiscoveryAgentError2
func callbackQBluetoothServiceDiscoveryAgentError2(ptr unsafe.Pointer, ptrName *C.char, error C.int) {
	defer qt.Recovering("callback QBluetoothServiceDiscoveryAgent::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(QBluetoothServiceDiscoveryAgent__Error))(QBluetoothServiceDiscoveryAgent__Error(error))
	}

}

func (ptr *QBluetoothServiceDiscoveryAgent) Error2(error QBluetoothServiceDiscoveryAgent__Error) {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::error")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_Error2(ptr.Pointer(), C.int(error))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectFinished(f func()) {
	defer qt.Recovering("connect QBluetoothServiceDiscoveryAgent::finished")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectFinished() {
	defer qt.Recovering("disconnect QBluetoothServiceDiscoveryAgent::finished")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQBluetoothServiceDiscoveryAgentFinished
func callbackQBluetoothServiceDiscoveryAgentFinished(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QBluetoothServiceDiscoveryAgent::finished")

	if signal := qt.GetSignal(C.GoString(ptrName), "finished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBluetoothServiceDiscoveryAgent) Finished() {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::finished")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_Finished(ptr.Pointer())
	}
}

func NewQBluetoothServiceDiscoveryAgent(parent core.QObject_ITF) *QBluetoothServiceDiscoveryAgent {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::QBluetoothServiceDiscoveryAgent")

	return newQBluetoothServiceDiscoveryAgentFromPointer(C.QBluetoothServiceDiscoveryAgent_NewQBluetoothServiceDiscoveryAgent(core.PointerFromQObject(parent)))
}

func NewQBluetoothServiceDiscoveryAgent2(deviceAdapter QBluetoothAddress_ITF, parent core.QObject_ITF) *QBluetoothServiceDiscoveryAgent {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::QBluetoothServiceDiscoveryAgent")

	return newQBluetoothServiceDiscoveryAgentFromPointer(C.QBluetoothServiceDiscoveryAgent_NewQBluetoothServiceDiscoveryAgent2(PointerFromQBluetoothAddress(deviceAdapter), core.PointerFromQObject(parent)))
}

func (ptr *QBluetoothServiceDiscoveryAgent) Clear() {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::clear")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_Clear(ptr.Pointer())
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) Error() QBluetoothServiceDiscoveryAgent__Error {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::error")

	if ptr.Pointer() != nil {
		return QBluetoothServiceDiscoveryAgent__Error(C.QBluetoothServiceDiscoveryAgent_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothServiceDiscoveryAgent) ErrorString() string {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothServiceDiscoveryAgent_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothServiceDiscoveryAgent) IsActive() bool {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::isActive")

	if ptr.Pointer() != nil {
		return C.QBluetoothServiceDiscoveryAgent_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothServiceDiscoveryAgent) SetRemoteAddress(address QBluetoothAddress_ITF) bool {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::setRemoteAddress")

	if ptr.Pointer() != nil {
		return C.QBluetoothServiceDiscoveryAgent_SetRemoteAddress(ptr.Pointer(), PointerFromQBluetoothAddress(address)) != 0
	}
	return false
}

func (ptr *QBluetoothServiceDiscoveryAgent) SetUuidFilter2(uuid QBluetoothUuid_ITF) {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::setUuidFilter")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_SetUuidFilter2(ptr.Pointer(), PointerFromQBluetoothUuid(uuid))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) Start(mode QBluetoothServiceDiscoveryAgent__DiscoveryMode) {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::start")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_Start(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) Stop() {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::stop")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_Stop(ptr.Pointer())
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DestroyQBluetoothServiceDiscoveryAgent() {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::~QBluetoothServiceDiscoveryAgent")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_DestroyQBluetoothServiceDiscoveryAgent(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QBluetoothServiceDiscoveryAgent::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QBluetoothServiceDiscoveryAgent::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQBluetoothServiceDiscoveryAgentTimerEvent
func callbackQBluetoothServiceDiscoveryAgentTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothServiceDiscoveryAgent::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQBluetoothServiceDiscoveryAgentFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::timerEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::timerEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QBluetoothServiceDiscoveryAgent::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QBluetoothServiceDiscoveryAgent::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQBluetoothServiceDiscoveryAgentChildEvent
func callbackQBluetoothServiceDiscoveryAgentChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothServiceDiscoveryAgent::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQBluetoothServiceDiscoveryAgentFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::childEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::childEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QBluetoothServiceDiscoveryAgent::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QBluetoothServiceDiscoveryAgent::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQBluetoothServiceDiscoveryAgentCustomEvent
func callbackQBluetoothServiceDiscoveryAgentCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothServiceDiscoveryAgent::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQBluetoothServiceDiscoveryAgentFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::customEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::customEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QBluetoothServiceInfo struct {
	ptr unsafe.Pointer
}

type QBluetoothServiceInfo_ITF interface {
	QBluetoothServiceInfo_PTR() *QBluetoothServiceInfo
}

func (p *QBluetoothServiceInfo) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QBluetoothServiceInfo) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQBluetoothServiceInfo(ptr QBluetoothServiceInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothServiceInfo_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothServiceInfoFromPointer(ptr unsafe.Pointer) *QBluetoothServiceInfo {
	var n = new(QBluetoothServiceInfo)
	n.SetPointer(ptr)
	return n
}

func newQBluetoothServiceInfoFromPointer(ptr unsafe.Pointer) *QBluetoothServiceInfo {
	var n = NewQBluetoothServiceInfoFromPointer(ptr)
	return n
}

func (ptr *QBluetoothServiceInfo) QBluetoothServiceInfo_PTR() *QBluetoothServiceInfo {
	return ptr
}

//QBluetoothServiceInfo::AttributeId
type QBluetoothServiceInfo__AttributeId int64

var (
	QBluetoothServiceInfo__ServiceRecordHandle              = QBluetoothServiceInfo__AttributeId(0x0000)
	QBluetoothServiceInfo__ServiceClassIds                  = QBluetoothServiceInfo__AttributeId(0x0001)
	QBluetoothServiceInfo__ServiceRecordState               = QBluetoothServiceInfo__AttributeId(0x0002)
	QBluetoothServiceInfo__ServiceId                        = QBluetoothServiceInfo__AttributeId(0x0003)
	QBluetoothServiceInfo__ProtocolDescriptorList           = QBluetoothServiceInfo__AttributeId(0x0004)
	QBluetoothServiceInfo__BrowseGroupList                  = QBluetoothServiceInfo__AttributeId(0x0005)
	QBluetoothServiceInfo__LanguageBaseAttributeIdList      = QBluetoothServiceInfo__AttributeId(0x0006)
	QBluetoothServiceInfo__ServiceInfoTimeToLive            = QBluetoothServiceInfo__AttributeId(0x0007)
	QBluetoothServiceInfo__ServiceAvailability              = QBluetoothServiceInfo__AttributeId(0x0008)
	QBluetoothServiceInfo__BluetoothProfileDescriptorList   = QBluetoothServiceInfo__AttributeId(0x0009)
	QBluetoothServiceInfo__DocumentationUrl                 = QBluetoothServiceInfo__AttributeId(0x000A)
	QBluetoothServiceInfo__ClientExecutableUrl              = QBluetoothServiceInfo__AttributeId(0x000B)
	QBluetoothServiceInfo__IconUrl                          = QBluetoothServiceInfo__AttributeId(0x000C)
	QBluetoothServiceInfo__AdditionalProtocolDescriptorList = QBluetoothServiceInfo__AttributeId(0x000D)
	QBluetoothServiceInfo__PrimaryLanguageBase              = QBluetoothServiceInfo__AttributeId(0x0100)
	QBluetoothServiceInfo__ServiceName                      = QBluetoothServiceInfo__AttributeId(C.QBluetoothServiceInfo_ServiceName_Type())
	QBluetoothServiceInfo__ServiceDescription               = QBluetoothServiceInfo__AttributeId(C.QBluetoothServiceInfo_ServiceDescription_Type())
	QBluetoothServiceInfo__ServiceProvider                  = QBluetoothServiceInfo__AttributeId(C.QBluetoothServiceInfo_ServiceProvider_Type())
)

//QBluetoothServiceInfo::Protocol
type QBluetoothServiceInfo__Protocol int64

const (
	QBluetoothServiceInfo__UnknownProtocol = QBluetoothServiceInfo__Protocol(0)
	QBluetoothServiceInfo__L2capProtocol   = QBluetoothServiceInfo__Protocol(1)
	QBluetoothServiceInfo__RfcommProtocol  = QBluetoothServiceInfo__Protocol(2)
)

func (ptr *QBluetoothServiceInfo) ServiceDescription() string {
	defer qt.Recovering("QBluetoothServiceInfo::serviceDescription")

	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothServiceInfo_ServiceDescription(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothServiceInfo) ServiceName() string {
	defer qt.Recovering("QBluetoothServiceInfo::serviceName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothServiceInfo_ServiceName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothServiceInfo) ServiceProvider() string {
	defer qt.Recovering("QBluetoothServiceInfo::serviceProvider")

	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothServiceInfo_ServiceProvider(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothServiceInfo) SetServiceDescription(description string) {
	defer qt.Recovering("QBluetoothServiceInfo::setServiceDescription")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceInfo_SetServiceDescription(ptr.Pointer(), C.CString(description))
	}
}

func (ptr *QBluetoothServiceInfo) SetServiceName(name string) {
	defer qt.Recovering("QBluetoothServiceInfo::setServiceName")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceInfo_SetServiceName(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QBluetoothServiceInfo) SetServiceProvider(provider string) {
	defer qt.Recovering("QBluetoothServiceInfo::setServiceProvider")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceInfo_SetServiceProvider(ptr.Pointer(), C.CString(provider))
	}
}

func (ptr *QBluetoothServiceInfo) SetServiceUuid(uuid QBluetoothUuid_ITF) {
	defer qt.Recovering("QBluetoothServiceInfo::setServiceUuid")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceInfo_SetServiceUuid(ptr.Pointer(), PointerFromQBluetoothUuid(uuid))
	}
}

func NewQBluetoothServiceInfo() *QBluetoothServiceInfo {
	defer qt.Recovering("QBluetoothServiceInfo::QBluetoothServiceInfo")

	return newQBluetoothServiceInfoFromPointer(C.QBluetoothServiceInfo_NewQBluetoothServiceInfo())
}

func NewQBluetoothServiceInfo2(other QBluetoothServiceInfo_ITF) *QBluetoothServiceInfo {
	defer qt.Recovering("QBluetoothServiceInfo::QBluetoothServiceInfo")

	return newQBluetoothServiceInfoFromPointer(C.QBluetoothServiceInfo_NewQBluetoothServiceInfo2(PointerFromQBluetoothServiceInfo(other)))
}

func (ptr *QBluetoothServiceInfo) IsComplete() bool {
	defer qt.Recovering("QBluetoothServiceInfo::isComplete")

	if ptr.Pointer() != nil {
		return C.QBluetoothServiceInfo_IsComplete(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothServiceInfo) IsRegistered() bool {
	defer qt.Recovering("QBluetoothServiceInfo::isRegistered")

	if ptr.Pointer() != nil {
		return C.QBluetoothServiceInfo_IsRegistered(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothServiceInfo) IsValid() bool {
	defer qt.Recovering("QBluetoothServiceInfo::isValid")

	if ptr.Pointer() != nil {
		return C.QBluetoothServiceInfo_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothServiceInfo) ProtocolServiceMultiplexer() int {
	defer qt.Recovering("QBluetoothServiceInfo::protocolServiceMultiplexer")

	if ptr.Pointer() != nil {
		return int(C.QBluetoothServiceInfo_ProtocolServiceMultiplexer(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothServiceInfo) RegisterService(localAdapter QBluetoothAddress_ITF) bool {
	defer qt.Recovering("QBluetoothServiceInfo::registerService")

	if ptr.Pointer() != nil {
		return C.QBluetoothServiceInfo_RegisterService(ptr.Pointer(), PointerFromQBluetoothAddress(localAdapter)) != 0
	}
	return false
}

func (ptr *QBluetoothServiceInfo) ServerChannel() int {
	defer qt.Recovering("QBluetoothServiceInfo::serverChannel")

	if ptr.Pointer() != nil {
		return int(C.QBluetoothServiceInfo_ServerChannel(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothServiceInfo) SetDevice(device QBluetoothDeviceInfo_ITF) {
	defer qt.Recovering("QBluetoothServiceInfo::setDevice")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceInfo_SetDevice(ptr.Pointer(), PointerFromQBluetoothDeviceInfo(device))
	}
}

func (ptr *QBluetoothServiceInfo) SocketProtocol() QBluetoothServiceInfo__Protocol {
	defer qt.Recovering("QBluetoothServiceInfo::socketProtocol")

	if ptr.Pointer() != nil {
		return QBluetoothServiceInfo__Protocol(C.QBluetoothServiceInfo_SocketProtocol(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothServiceInfo) UnregisterService() bool {
	defer qt.Recovering("QBluetoothServiceInfo::unregisterService")

	if ptr.Pointer() != nil {
		return C.QBluetoothServiceInfo_UnregisterService(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothServiceInfo) DestroyQBluetoothServiceInfo() {
	defer qt.Recovering("QBluetoothServiceInfo::~QBluetoothServiceInfo")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceInfo_DestroyQBluetoothServiceInfo(ptr.Pointer())
	}
}

type QBluetoothSocket struct {
	core.QIODevice
}

type QBluetoothSocket_ITF interface {
	core.QIODevice_ITF
	QBluetoothSocket_PTR() *QBluetoothSocket
}

func PointerFromQBluetoothSocket(ptr QBluetoothSocket_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothSocket_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothSocketFromPointer(ptr unsafe.Pointer) *QBluetoothSocket {
	var n = new(QBluetoothSocket)
	n.SetPointer(ptr)
	return n
}

func newQBluetoothSocketFromPointer(ptr unsafe.Pointer) *QBluetoothSocket {
	var n = NewQBluetoothSocketFromPointer(ptr)
	for len(n.ObjectName()) < len("QBluetoothSocket_") {
		n.SetObjectName("QBluetoothSocket_" + qt.Identifier())
	}
	return n
}

func (ptr *QBluetoothSocket) QBluetoothSocket_PTR() *QBluetoothSocket {
	return ptr
}

//QBluetoothSocket::SocketError
type QBluetoothSocket__SocketError int64

const (
	QBluetoothSocket__NoSocketError            = QBluetoothSocket__SocketError(-2)
	QBluetoothSocket__UnknownSocketError       = QBluetoothSocket__SocketError(network.QAbstractSocket__UnknownSocketError)
	QBluetoothSocket__HostNotFoundError        = QBluetoothSocket__SocketError(network.QAbstractSocket__HostNotFoundError)
	QBluetoothSocket__ServiceNotFoundError     = QBluetoothSocket__SocketError(network.QAbstractSocket__SocketAddressNotAvailableError)
	QBluetoothSocket__NetworkError             = QBluetoothSocket__SocketError(network.QAbstractSocket__NetworkError)
	QBluetoothSocket__UnsupportedProtocolError = QBluetoothSocket__SocketError(8)
	QBluetoothSocket__OperationError           = QBluetoothSocket__SocketError(network.QAbstractSocket__OperationError)
)

//QBluetoothSocket::SocketState
type QBluetoothSocket__SocketState int64

const (
	QBluetoothSocket__UnconnectedState   = QBluetoothSocket__SocketState(network.QAbstractSocket__UnconnectedState)
	QBluetoothSocket__ServiceLookupState = QBluetoothSocket__SocketState(network.QAbstractSocket__HostLookupState)
	QBluetoothSocket__ConnectingState    = QBluetoothSocket__SocketState(network.QAbstractSocket__ConnectingState)
	QBluetoothSocket__ConnectedState     = QBluetoothSocket__SocketState(network.QAbstractSocket__ConnectedState)
	QBluetoothSocket__BoundState         = QBluetoothSocket__SocketState(network.QAbstractSocket__BoundState)
	QBluetoothSocket__ClosingState       = QBluetoothSocket__SocketState(network.QAbstractSocket__ClosingState)
	QBluetoothSocket__ListeningState     = QBluetoothSocket__SocketState(network.QAbstractSocket__ListeningState)
)

func (ptr *QBluetoothSocket) ConnectConnected(f func()) {
	defer qt.Recovering("connect QBluetoothSocket::connected")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_ConnectConnected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "connected", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectConnected() {
	defer qt.Recovering("disconnect QBluetoothSocket::connected")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DisconnectConnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "connected")
	}
}

//export callbackQBluetoothSocketConnected
func callbackQBluetoothSocketConnected(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QBluetoothSocket::connected")

	if signal := qt.GetSignal(C.GoString(ptrName), "connected"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBluetoothSocket) Connected() {
	defer qt.Recovering("QBluetoothSocket::connected")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_Connected(ptr.Pointer())
	}
}

func (ptr *QBluetoothSocket) ConnectDisconnected(f func()) {
	defer qt.Recovering("connect QBluetoothSocket::disconnected")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_ConnectDisconnected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "disconnected", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectDisconnected() {
	defer qt.Recovering("disconnect QBluetoothSocket::disconnected")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DisconnectDisconnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "disconnected")
	}
}

//export callbackQBluetoothSocketDisconnected
func callbackQBluetoothSocketDisconnected(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QBluetoothSocket::disconnected")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnected"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBluetoothSocket) Disconnected() {
	defer qt.Recovering("QBluetoothSocket::disconnected")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_Disconnected(ptr.Pointer())
	}
}

func (ptr *QBluetoothSocket) ConnectError2(f func(error QBluetoothSocket__SocketError)) {
	defer qt.Recovering("connect QBluetoothSocket::error")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error2", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectError2() {
	defer qt.Recovering("disconnect QBluetoothSocket::error")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error2")
	}
}

//export callbackQBluetoothSocketError2
func callbackQBluetoothSocketError2(ptr unsafe.Pointer, ptrName *C.char, error C.int) {
	defer qt.Recovering("callback QBluetoothSocket::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(QBluetoothSocket__SocketError))(QBluetoothSocket__SocketError(error))
	}

}

func (ptr *QBluetoothSocket) Error2(error QBluetoothSocket__SocketError) {
	defer qt.Recovering("QBluetoothSocket::error")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_Error2(ptr.Pointer(), C.int(error))
	}
}

func (ptr *QBluetoothSocket) ConnectStateChanged(f func(state QBluetoothSocket__SocketState)) {
	defer qt.Recovering("connect QBluetoothSocket::stateChanged")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectStateChanged() {
	defer qt.Recovering("disconnect QBluetoothSocket::stateChanged")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQBluetoothSocketStateChanged
func callbackQBluetoothSocketStateChanged(ptr unsafe.Pointer, ptrName *C.char, state C.int) {
	defer qt.Recovering("callback QBluetoothSocket::stateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stateChanged"); signal != nil {
		signal.(func(QBluetoothSocket__SocketState))(QBluetoothSocket__SocketState(state))
	}

}

func (ptr *QBluetoothSocket) StateChanged(state QBluetoothSocket__SocketState) {
	defer qt.Recovering("QBluetoothSocket::stateChanged")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_StateChanged(ptr.Pointer(), C.int(state))
	}
}

func NewQBluetoothSocket(socketType QBluetoothServiceInfo__Protocol, parent core.QObject_ITF) *QBluetoothSocket {
	defer qt.Recovering("QBluetoothSocket::QBluetoothSocket")

	return newQBluetoothSocketFromPointer(C.QBluetoothSocket_NewQBluetoothSocket(C.int(socketType), core.PointerFromQObject(parent)))
}

func NewQBluetoothSocket2(parent core.QObject_ITF) *QBluetoothSocket {
	defer qt.Recovering("QBluetoothSocket::QBluetoothSocket")

	return newQBluetoothSocketFromPointer(C.QBluetoothSocket_NewQBluetoothSocket2(core.PointerFromQObject(parent)))
}

func (ptr *QBluetoothSocket) Abort() {
	defer qt.Recovering("QBluetoothSocket::abort")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_Abort(ptr.Pointer())
	}
}

func (ptr *QBluetoothSocket) BytesAvailable() int64 {
	defer qt.Recovering("QBluetoothSocket::bytesAvailable")

	if ptr.Pointer() != nil {
		return int64(C.QBluetoothSocket_BytesAvailable(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothSocket) BytesToWrite() int64 {
	defer qt.Recovering("QBluetoothSocket::bytesToWrite")

	if ptr.Pointer() != nil {
		return int64(C.QBluetoothSocket_BytesToWrite(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothSocket) CanReadLine() bool {
	defer qt.Recovering("QBluetoothSocket::canReadLine")

	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_CanReadLine(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothSocket) ConnectClose(f func()) {
	defer qt.Recovering("connect QBluetoothSocket::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "close", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectClose() {
	defer qt.Recovering("disconnect QBluetoothSocket::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "close")
	}
}

//export callbackQBluetoothSocketClose
func callbackQBluetoothSocketClose(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QBluetoothSocket::close")

	if signal := qt.GetSignal(C.GoString(ptrName), "close"); signal != nil {
		signal.(func())()
	} else {
		NewQBluetoothSocketFromPointer(ptr).CloseDefault()
	}
}

func (ptr *QBluetoothSocket) Close() {
	defer qt.Recovering("QBluetoothSocket::close")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_Close(ptr.Pointer())
	}
}

func (ptr *QBluetoothSocket) CloseDefault() {
	defer qt.Recovering("QBluetoothSocket::close")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_CloseDefault(ptr.Pointer())
	}
}

func (ptr *QBluetoothSocket) ConnectToService2(address QBluetoothAddress_ITF, uuid QBluetoothUuid_ITF, openMode core.QIODevice__OpenModeFlag) {
	defer qt.Recovering("QBluetoothSocket::connectToService")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_ConnectToService2(ptr.Pointer(), PointerFromQBluetoothAddress(address), PointerFromQBluetoothUuid(uuid), C.int(openMode))
	}
}

func (ptr *QBluetoothSocket) ConnectToService(service QBluetoothServiceInfo_ITF, openMode core.QIODevice__OpenModeFlag) {
	defer qt.Recovering("QBluetoothSocket::connectToService")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_ConnectToService(ptr.Pointer(), PointerFromQBluetoothServiceInfo(service), C.int(openMode))
	}
}

func (ptr *QBluetoothSocket) DisconnectFromService() {
	defer qt.Recovering("QBluetoothSocket::disconnectFromService")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DisconnectFromService(ptr.Pointer())
	}
}

func (ptr *QBluetoothSocket) Error() QBluetoothSocket__SocketError {
	defer qt.Recovering("QBluetoothSocket::error")

	if ptr.Pointer() != nil {
		return QBluetoothSocket__SocketError(C.QBluetoothSocket_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothSocket) ErrorString() string {
	defer qt.Recovering("QBluetoothSocket::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothSocket_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothSocket) IsSequential() bool {
	defer qt.Recovering("QBluetoothSocket::isSequential")

	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_IsSequential(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothSocket) LocalName() string {
	defer qt.Recovering("QBluetoothSocket::localName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothSocket_LocalName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothSocket) PeerName() string {
	defer qt.Recovering("QBluetoothSocket::peerName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothSocket_PeerName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothSocket) ReadData(data string, maxSize int64) int64 {
	defer qt.Recovering("QBluetoothSocket::readData")

	if ptr.Pointer() != nil {
		return int64(C.QBluetoothSocket_ReadData(ptr.Pointer(), C.CString(data), C.longlong(maxSize)))
	}
	return 0
}

func (ptr *QBluetoothSocket) SetSocketDescriptor(socketDescriptor int, socketType QBluetoothServiceInfo__Protocol, socketState QBluetoothSocket__SocketState, openMode core.QIODevice__OpenModeFlag) bool {
	defer qt.Recovering("QBluetoothSocket::setSocketDescriptor")

	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_SetSocketDescriptor(ptr.Pointer(), C.int(socketDescriptor), C.int(socketType), C.int(socketState), C.int(openMode)) != 0
	}
	return false
}

func (ptr *QBluetoothSocket) SocketDescriptor() int {
	defer qt.Recovering("QBluetoothSocket::socketDescriptor")

	if ptr.Pointer() != nil {
		return int(C.QBluetoothSocket_SocketDescriptor(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothSocket) SocketType() QBluetoothServiceInfo__Protocol {
	defer qt.Recovering("QBluetoothSocket::socketType")

	if ptr.Pointer() != nil {
		return QBluetoothServiceInfo__Protocol(C.QBluetoothSocket_SocketType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothSocket) State() QBluetoothSocket__SocketState {
	defer qt.Recovering("QBluetoothSocket::state")

	if ptr.Pointer() != nil {
		return QBluetoothSocket__SocketState(C.QBluetoothSocket_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothSocket) WriteData(data string, maxSize int64) int64 {
	defer qt.Recovering("QBluetoothSocket::writeData")

	if ptr.Pointer() != nil {
		return int64(C.QBluetoothSocket_WriteData(ptr.Pointer(), C.CString(data), C.longlong(maxSize)))
	}
	return 0
}

func (ptr *QBluetoothSocket) DestroyQBluetoothSocket() {
	defer qt.Recovering("QBluetoothSocket::~QBluetoothSocket")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DestroyQBluetoothSocket(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothSocket) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QBluetoothSocket::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QBluetoothSocket::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQBluetoothSocketTimerEvent
func callbackQBluetoothSocketTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothSocket::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQBluetoothSocketFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QBluetoothSocket) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QBluetoothSocket::timerEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QBluetoothSocket) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QBluetoothSocket::timerEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QBluetoothSocket) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QBluetoothSocket::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QBluetoothSocket::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQBluetoothSocketChildEvent
func callbackQBluetoothSocketChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothSocket::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQBluetoothSocketFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QBluetoothSocket) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QBluetoothSocket::childEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QBluetoothSocket) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QBluetoothSocket::childEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QBluetoothSocket) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QBluetoothSocket::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QBluetoothSocket::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQBluetoothSocketCustomEvent
func callbackQBluetoothSocketCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothSocket::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQBluetoothSocketFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QBluetoothSocket) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QBluetoothSocket::customEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QBluetoothSocket) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QBluetoothSocket::customEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QBluetoothTransferManager struct {
	core.QObject
}

type QBluetoothTransferManager_ITF interface {
	core.QObject_ITF
	QBluetoothTransferManager_PTR() *QBluetoothTransferManager
}

func PointerFromQBluetoothTransferManager(ptr QBluetoothTransferManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothTransferManager_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothTransferManagerFromPointer(ptr unsafe.Pointer) *QBluetoothTransferManager {
	var n = new(QBluetoothTransferManager)
	n.SetPointer(ptr)
	return n
}

func newQBluetoothTransferManagerFromPointer(ptr unsafe.Pointer) *QBluetoothTransferManager {
	var n = NewQBluetoothTransferManagerFromPointer(ptr)
	for len(n.ObjectName()) < len("QBluetoothTransferManager_") {
		n.SetObjectName("QBluetoothTransferManager_" + qt.Identifier())
	}
	return n
}

func (ptr *QBluetoothTransferManager) QBluetoothTransferManager_PTR() *QBluetoothTransferManager {
	return ptr
}

func (ptr *QBluetoothTransferManager) Put(request QBluetoothTransferRequest_ITF, data core.QIODevice_ITF) *QBluetoothTransferReply {
	defer qt.Recovering("QBluetoothTransferManager::put")

	if ptr.Pointer() != nil {
		return NewQBluetoothTransferReplyFromPointer(C.QBluetoothTransferManager_Put(ptr.Pointer(), PointerFromQBluetoothTransferRequest(request), core.PointerFromQIODevice(data)))
	}
	return nil
}

func NewQBluetoothTransferManager(parent core.QObject_ITF) *QBluetoothTransferManager {
	defer qt.Recovering("QBluetoothTransferManager::QBluetoothTransferManager")

	return newQBluetoothTransferManagerFromPointer(C.QBluetoothTransferManager_NewQBluetoothTransferManager(core.PointerFromQObject(parent)))
}

func (ptr *QBluetoothTransferManager) ConnectFinished(f func(reply *QBluetoothTransferReply)) {
	defer qt.Recovering("connect QBluetoothTransferManager::finished")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QBluetoothTransferManager) DisconnectFinished() {
	defer qt.Recovering("disconnect QBluetoothTransferManager::finished")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQBluetoothTransferManagerFinished
func callbackQBluetoothTransferManagerFinished(ptr unsafe.Pointer, ptrName *C.char, reply unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothTransferManager::finished")

	if signal := qt.GetSignal(C.GoString(ptrName), "finished"); signal != nil {
		signal.(func(*QBluetoothTransferReply))(NewQBluetoothTransferReplyFromPointer(reply))
	}

}

func (ptr *QBluetoothTransferManager) Finished(reply QBluetoothTransferReply_ITF) {
	defer qt.Recovering("QBluetoothTransferManager::finished")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_Finished(ptr.Pointer(), PointerFromQBluetoothTransferReply(reply))
	}
}

func (ptr *QBluetoothTransferManager) DestroyQBluetoothTransferManager() {
	defer qt.Recovering("QBluetoothTransferManager::~QBluetoothTransferManager")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_DestroyQBluetoothTransferManager(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothTransferManager) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QBluetoothTransferManager::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QBluetoothTransferManager) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QBluetoothTransferManager::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQBluetoothTransferManagerTimerEvent
func callbackQBluetoothTransferManagerTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothTransferManager::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQBluetoothTransferManagerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QBluetoothTransferManager) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QBluetoothTransferManager::timerEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QBluetoothTransferManager) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QBluetoothTransferManager::timerEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QBluetoothTransferManager) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QBluetoothTransferManager::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QBluetoothTransferManager) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QBluetoothTransferManager::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQBluetoothTransferManagerChildEvent
func callbackQBluetoothTransferManagerChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothTransferManager::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQBluetoothTransferManagerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QBluetoothTransferManager) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QBluetoothTransferManager::childEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QBluetoothTransferManager) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QBluetoothTransferManager::childEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QBluetoothTransferManager) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QBluetoothTransferManager::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QBluetoothTransferManager) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QBluetoothTransferManager::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQBluetoothTransferManagerCustomEvent
func callbackQBluetoothTransferManagerCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothTransferManager::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQBluetoothTransferManagerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QBluetoothTransferManager) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QBluetoothTransferManager::customEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QBluetoothTransferManager) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QBluetoothTransferManager::customEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QBluetoothTransferReply struct {
	core.QObject
}

type QBluetoothTransferReply_ITF interface {
	core.QObject_ITF
	QBluetoothTransferReply_PTR() *QBluetoothTransferReply
}

func PointerFromQBluetoothTransferReply(ptr QBluetoothTransferReply_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothTransferReply_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothTransferReplyFromPointer(ptr unsafe.Pointer) *QBluetoothTransferReply {
	var n = new(QBluetoothTransferReply)
	n.SetPointer(ptr)
	return n
}

func newQBluetoothTransferReplyFromPointer(ptr unsafe.Pointer) *QBluetoothTransferReply {
	var n = NewQBluetoothTransferReplyFromPointer(ptr)
	for len(n.ObjectName()) < len("QBluetoothTransferReply_") {
		n.SetObjectName("QBluetoothTransferReply_" + qt.Identifier())
	}
	return n
}

func (ptr *QBluetoothTransferReply) QBluetoothTransferReply_PTR() *QBluetoothTransferReply {
	return ptr
}

//QBluetoothTransferReply::TransferError
type QBluetoothTransferReply__TransferError int64

const (
	QBluetoothTransferReply__NoError                   = QBluetoothTransferReply__TransferError(0)
	QBluetoothTransferReply__UnknownError              = QBluetoothTransferReply__TransferError(1)
	QBluetoothTransferReply__FileNotFoundError         = QBluetoothTransferReply__TransferError(2)
	QBluetoothTransferReply__HostNotFoundError         = QBluetoothTransferReply__TransferError(3)
	QBluetoothTransferReply__UserCanceledTransferError = QBluetoothTransferReply__TransferError(4)
	QBluetoothTransferReply__IODeviceNotReadableError  = QBluetoothTransferReply__TransferError(5)
	QBluetoothTransferReply__ResourceBusyError         = QBluetoothTransferReply__TransferError(6)
	QBluetoothTransferReply__SessionError              = QBluetoothTransferReply__TransferError(7)
)

func (ptr *QBluetoothTransferReply) Abort() {
	defer qt.Recovering("QBluetoothTransferReply::abort")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_Abort(ptr.Pointer())
	}
}

func (ptr *QBluetoothTransferReply) ConnectError2(f func(errorType QBluetoothTransferReply__TransferError)) {
	defer qt.Recovering("connect QBluetoothTransferReply::error")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error2", f)
	}
}

func (ptr *QBluetoothTransferReply) DisconnectError2() {
	defer qt.Recovering("disconnect QBluetoothTransferReply::error")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error2")
	}
}

//export callbackQBluetoothTransferReplyError2
func callbackQBluetoothTransferReplyError2(ptr unsafe.Pointer, ptrName *C.char, errorType C.int) {
	defer qt.Recovering("callback QBluetoothTransferReply::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(QBluetoothTransferReply__TransferError))(QBluetoothTransferReply__TransferError(errorType))
	}

}

func (ptr *QBluetoothTransferReply) Error2(errorType QBluetoothTransferReply__TransferError) {
	defer qt.Recovering("QBluetoothTransferReply::error")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_Error2(ptr.Pointer(), C.int(errorType))
	}
}

func (ptr *QBluetoothTransferReply) Error() QBluetoothTransferReply__TransferError {
	defer qt.Recovering("QBluetoothTransferReply::error")

	if ptr.Pointer() != nil {
		return QBluetoothTransferReply__TransferError(C.QBluetoothTransferReply_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothTransferReply) ErrorString() string {
	defer qt.Recovering("QBluetoothTransferReply::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothTransferReply_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothTransferReply) ConnectFinished(f func(reply *QBluetoothTransferReply)) {
	defer qt.Recovering("connect QBluetoothTransferReply::finished")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QBluetoothTransferReply) DisconnectFinished() {
	defer qt.Recovering("disconnect QBluetoothTransferReply::finished")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQBluetoothTransferReplyFinished
func callbackQBluetoothTransferReplyFinished(ptr unsafe.Pointer, ptrName *C.char, reply unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothTransferReply::finished")

	if signal := qt.GetSignal(C.GoString(ptrName), "finished"); signal != nil {
		signal.(func(*QBluetoothTransferReply))(NewQBluetoothTransferReplyFromPointer(reply))
	}

}

func (ptr *QBluetoothTransferReply) Finished(reply QBluetoothTransferReply_ITF) {
	defer qt.Recovering("QBluetoothTransferReply::finished")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_Finished(ptr.Pointer(), PointerFromQBluetoothTransferReply(reply))
	}
}

func (ptr *QBluetoothTransferReply) IsFinished() bool {
	defer qt.Recovering("QBluetoothTransferReply::isFinished")

	if ptr.Pointer() != nil {
		return C.QBluetoothTransferReply_IsFinished(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothTransferReply) IsRunning() bool {
	defer qt.Recovering("QBluetoothTransferReply::isRunning")

	if ptr.Pointer() != nil {
		return C.QBluetoothTransferReply_IsRunning(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothTransferReply) Manager() *QBluetoothTransferManager {
	defer qt.Recovering("QBluetoothTransferReply::manager")

	if ptr.Pointer() != nil {
		return NewQBluetoothTransferManagerFromPointer(C.QBluetoothTransferReply_Manager(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBluetoothTransferReply) ConnectTransferProgress(f func(bytesTransferred int64, bytesTotal int64)) {
	defer qt.Recovering("connect QBluetoothTransferReply::transferProgress")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_ConnectTransferProgress(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "transferProgress", f)
	}
}

func (ptr *QBluetoothTransferReply) DisconnectTransferProgress() {
	defer qt.Recovering("disconnect QBluetoothTransferReply::transferProgress")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_DisconnectTransferProgress(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "transferProgress")
	}
}

//export callbackQBluetoothTransferReplyTransferProgress
func callbackQBluetoothTransferReplyTransferProgress(ptr unsafe.Pointer, ptrName *C.char, bytesTransferred C.longlong, bytesTotal C.longlong) {
	defer qt.Recovering("callback QBluetoothTransferReply::transferProgress")

	if signal := qt.GetSignal(C.GoString(ptrName), "transferProgress"); signal != nil {
		signal.(func(int64, int64))(int64(bytesTransferred), int64(bytesTotal))
	}

}

func (ptr *QBluetoothTransferReply) TransferProgress(bytesTransferred int64, bytesTotal int64) {
	defer qt.Recovering("QBluetoothTransferReply::transferProgress")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_TransferProgress(ptr.Pointer(), C.longlong(bytesTransferred), C.longlong(bytesTotal))
	}
}

func (ptr *QBluetoothTransferReply) DestroyQBluetoothTransferReply() {
	defer qt.Recovering("QBluetoothTransferReply::~QBluetoothTransferReply")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_DestroyQBluetoothTransferReply(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothTransferReply) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QBluetoothTransferReply::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QBluetoothTransferReply) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QBluetoothTransferReply::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQBluetoothTransferReplyTimerEvent
func callbackQBluetoothTransferReplyTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothTransferReply::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQBluetoothTransferReplyFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QBluetoothTransferReply) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QBluetoothTransferReply::timerEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QBluetoothTransferReply) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QBluetoothTransferReply::timerEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QBluetoothTransferReply) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QBluetoothTransferReply::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QBluetoothTransferReply) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QBluetoothTransferReply::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQBluetoothTransferReplyChildEvent
func callbackQBluetoothTransferReplyChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothTransferReply::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQBluetoothTransferReplyFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QBluetoothTransferReply) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QBluetoothTransferReply::childEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QBluetoothTransferReply) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QBluetoothTransferReply::childEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QBluetoothTransferReply) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QBluetoothTransferReply::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QBluetoothTransferReply) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QBluetoothTransferReply::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQBluetoothTransferReplyCustomEvent
func callbackQBluetoothTransferReplyCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothTransferReply::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQBluetoothTransferReplyFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QBluetoothTransferReply) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QBluetoothTransferReply::customEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QBluetoothTransferReply) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QBluetoothTransferReply::customEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QBluetoothTransferRequest struct {
	ptr unsafe.Pointer
}

type QBluetoothTransferRequest_ITF interface {
	QBluetoothTransferRequest_PTR() *QBluetoothTransferRequest
}

func (p *QBluetoothTransferRequest) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QBluetoothTransferRequest) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQBluetoothTransferRequest(ptr QBluetoothTransferRequest_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothTransferRequest_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothTransferRequestFromPointer(ptr unsafe.Pointer) *QBluetoothTransferRequest {
	var n = new(QBluetoothTransferRequest)
	n.SetPointer(ptr)
	return n
}

func newQBluetoothTransferRequestFromPointer(ptr unsafe.Pointer) *QBluetoothTransferRequest {
	var n = NewQBluetoothTransferRequestFromPointer(ptr)
	return n
}

func (ptr *QBluetoothTransferRequest) QBluetoothTransferRequest_PTR() *QBluetoothTransferRequest {
	return ptr
}

//QBluetoothTransferRequest::Attribute
type QBluetoothTransferRequest__Attribute int64

const (
	QBluetoothTransferRequest__DescriptionAttribute = QBluetoothTransferRequest__Attribute(0)
	QBluetoothTransferRequest__TimeAttribute        = QBluetoothTransferRequest__Attribute(1)
	QBluetoothTransferRequest__TypeAttribute        = QBluetoothTransferRequest__Attribute(2)
	QBluetoothTransferRequest__LengthAttribute      = QBluetoothTransferRequest__Attribute(3)
	QBluetoothTransferRequest__NameAttribute        = QBluetoothTransferRequest__Attribute(4)
)

func NewQBluetoothTransferRequest(address QBluetoothAddress_ITF) *QBluetoothTransferRequest {
	defer qt.Recovering("QBluetoothTransferRequest::QBluetoothTransferRequest")

	return newQBluetoothTransferRequestFromPointer(C.QBluetoothTransferRequest_NewQBluetoothTransferRequest(PointerFromQBluetoothAddress(address)))
}

func NewQBluetoothTransferRequest2(other QBluetoothTransferRequest_ITF) *QBluetoothTransferRequest {
	defer qt.Recovering("QBluetoothTransferRequest::QBluetoothTransferRequest")

	return newQBluetoothTransferRequestFromPointer(C.QBluetoothTransferRequest_NewQBluetoothTransferRequest2(PointerFromQBluetoothTransferRequest(other)))
}

func (ptr *QBluetoothTransferRequest) Attribute(code QBluetoothTransferRequest__Attribute, defaultValue core.QVariant_ITF) *core.QVariant {
	defer qt.Recovering("QBluetoothTransferRequest::attribute")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QBluetoothTransferRequest_Attribute(ptr.Pointer(), C.int(code), core.PointerFromQVariant(defaultValue)))
	}
	return nil
}

func (ptr *QBluetoothTransferRequest) SetAttribute(code QBluetoothTransferRequest__Attribute, value core.QVariant_ITF) {
	defer qt.Recovering("QBluetoothTransferRequest::setAttribute")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferRequest_SetAttribute(ptr.Pointer(), C.int(code), core.PointerFromQVariant(value))
	}
}

func (ptr *QBluetoothTransferRequest) DestroyQBluetoothTransferRequest() {
	defer qt.Recovering("QBluetoothTransferRequest::~QBluetoothTransferRequest")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferRequest_DestroyQBluetoothTransferRequest(ptr.Pointer())
	}
}

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

func newQBluetoothUuidFromPointer(ptr unsafe.Pointer) *QBluetoothUuid {
	var n = NewQBluetoothUuidFromPointer(ptr)
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
	defer qt.Recovering("QBluetoothUuid::QBluetoothUuid")

	return newQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid())
}

func NewQBluetoothUuid4(uuid QBluetoothUuid__CharacteristicType) *QBluetoothUuid {
	defer qt.Recovering("QBluetoothUuid::QBluetoothUuid")

	return newQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid4(C.int(uuid)))
}

func NewQBluetoothUuid5(uuid QBluetoothUuid__DescriptorType) *QBluetoothUuid {
	defer qt.Recovering("QBluetoothUuid::QBluetoothUuid")

	return newQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid5(C.int(uuid)))
}

func NewQBluetoothUuid2(uuid QBluetoothUuid__ProtocolUuid) *QBluetoothUuid {
	defer qt.Recovering("QBluetoothUuid::QBluetoothUuid")

	return newQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid2(C.int(uuid)))
}

func NewQBluetoothUuid3(uuid QBluetoothUuid__ServiceClassUuid) *QBluetoothUuid {
	defer qt.Recovering("QBluetoothUuid::QBluetoothUuid")

	return newQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid3(C.int(uuid)))
}

func NewQBluetoothUuid10(uuid QBluetoothUuid_ITF) *QBluetoothUuid {
	defer qt.Recovering("QBluetoothUuid::QBluetoothUuid")

	return newQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid10(PointerFromQBluetoothUuid(uuid)))
}

func NewQBluetoothUuid9(uuid string) *QBluetoothUuid {
	defer qt.Recovering("QBluetoothUuid::QBluetoothUuid")

	return newQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid9(C.CString(uuid)))
}

func NewQBluetoothUuid11(uuid core.QUuid_ITF) *QBluetoothUuid {
	defer qt.Recovering("QBluetoothUuid::QBluetoothUuid")

	return newQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid11(core.PointerFromQUuid(uuid)))
}

func QBluetoothUuid_CharacteristicToString(uuid QBluetoothUuid__CharacteristicType) string {
	defer qt.Recovering("QBluetoothUuid::characteristicToString")

	return C.GoString(C.QBluetoothUuid_QBluetoothUuid_CharacteristicToString(C.int(uuid)))
}

func QBluetoothUuid_DescriptorToString(uuid QBluetoothUuid__DescriptorType) string {
	defer qt.Recovering("QBluetoothUuid::descriptorToString")

	return C.GoString(C.QBluetoothUuid_QBluetoothUuid_DescriptorToString(C.int(uuid)))
}

func (ptr *QBluetoothUuid) MinimumSize() int {
	defer qt.Recovering("QBluetoothUuid::minimumSize")

	if ptr.Pointer() != nil {
		return int(C.QBluetoothUuid_MinimumSize(ptr.Pointer()))
	}
	return 0
}

func QBluetoothUuid_ProtocolToString(uuid QBluetoothUuid__ProtocolUuid) string {
	defer qt.Recovering("QBluetoothUuid::protocolToString")

	return C.GoString(C.QBluetoothUuid_QBluetoothUuid_ProtocolToString(C.int(uuid)))
}

func QBluetoothUuid_ServiceClassToString(uuid QBluetoothUuid__ServiceClassUuid) string {
	defer qt.Recovering("QBluetoothUuid::serviceClassToString")

	return C.GoString(C.QBluetoothUuid_QBluetoothUuid_ServiceClassToString(C.int(uuid)))
}

func (ptr *QBluetoothUuid) DestroyQBluetoothUuid() {
	defer qt.Recovering("QBluetoothUuid::~QBluetoothUuid")

	if ptr.Pointer() != nil {
		C.QBluetoothUuid_DestroyQBluetoothUuid(ptr.Pointer())
	}
}

type QLowEnergyCharacteristic struct {
	ptr unsafe.Pointer
}

type QLowEnergyCharacteristic_ITF interface {
	QLowEnergyCharacteristic_PTR() *QLowEnergyCharacteristic
}

func (p *QLowEnergyCharacteristic) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QLowEnergyCharacteristic) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQLowEnergyCharacteristic(ptr QLowEnergyCharacteristic_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLowEnergyCharacteristic_PTR().Pointer()
	}
	return nil
}

func NewQLowEnergyCharacteristicFromPointer(ptr unsafe.Pointer) *QLowEnergyCharacteristic {
	var n = new(QLowEnergyCharacteristic)
	n.SetPointer(ptr)
	return n
}

func newQLowEnergyCharacteristicFromPointer(ptr unsafe.Pointer) *QLowEnergyCharacteristic {
	var n = NewQLowEnergyCharacteristicFromPointer(ptr)
	return n
}

func (ptr *QLowEnergyCharacteristic) QLowEnergyCharacteristic_PTR() *QLowEnergyCharacteristic {
	return ptr
}

//QLowEnergyCharacteristic::PropertyType
type QLowEnergyCharacteristic__PropertyType int64

const (
	QLowEnergyCharacteristic__Unknown          = QLowEnergyCharacteristic__PropertyType(0x00)
	QLowEnergyCharacteristic__Broadcasting     = QLowEnergyCharacteristic__PropertyType(0x01)
	QLowEnergyCharacteristic__Read             = QLowEnergyCharacteristic__PropertyType(0x02)
	QLowEnergyCharacteristic__WriteNoResponse  = QLowEnergyCharacteristic__PropertyType(0x04)
	QLowEnergyCharacteristic__Write            = QLowEnergyCharacteristic__PropertyType(0x08)
	QLowEnergyCharacteristic__Notify           = QLowEnergyCharacteristic__PropertyType(0x10)
	QLowEnergyCharacteristic__Indicate         = QLowEnergyCharacteristic__PropertyType(0x20)
	QLowEnergyCharacteristic__WriteSigned      = QLowEnergyCharacteristic__PropertyType(0x40)
	QLowEnergyCharacteristic__ExtendedProperty = QLowEnergyCharacteristic__PropertyType(0x80)
)

func NewQLowEnergyCharacteristic() *QLowEnergyCharacteristic {
	defer qt.Recovering("QLowEnergyCharacteristic::QLowEnergyCharacteristic")

	return newQLowEnergyCharacteristicFromPointer(C.QLowEnergyCharacteristic_NewQLowEnergyCharacteristic())
}

func NewQLowEnergyCharacteristic2(other QLowEnergyCharacteristic_ITF) *QLowEnergyCharacteristic {
	defer qt.Recovering("QLowEnergyCharacteristic::QLowEnergyCharacteristic")

	return newQLowEnergyCharacteristicFromPointer(C.QLowEnergyCharacteristic_NewQLowEnergyCharacteristic2(PointerFromQLowEnergyCharacteristic(other)))
}

func (ptr *QLowEnergyCharacteristic) IsValid() bool {
	defer qt.Recovering("QLowEnergyCharacteristic::isValid")

	if ptr.Pointer() != nil {
		return C.QLowEnergyCharacteristic_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLowEnergyCharacteristic) Name() string {
	defer qt.Recovering("QLowEnergyCharacteristic::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QLowEnergyCharacteristic_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLowEnergyCharacteristic) Properties() QLowEnergyCharacteristic__PropertyType {
	defer qt.Recovering("QLowEnergyCharacteristic::properties")

	if ptr.Pointer() != nil {
		return QLowEnergyCharacteristic__PropertyType(C.QLowEnergyCharacteristic_Properties(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyCharacteristic) Value() *core.QByteArray {
	defer qt.Recovering("QLowEnergyCharacteristic::value")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QLowEnergyCharacteristic_Value(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLowEnergyCharacteristic) DestroyQLowEnergyCharacteristic() {
	defer qt.Recovering("QLowEnergyCharacteristic::~QLowEnergyCharacteristic")

	if ptr.Pointer() != nil {
		C.QLowEnergyCharacteristic_DestroyQLowEnergyCharacteristic(ptr.Pointer())
	}
}

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
	return n
}

func newQLowEnergyControllerFromPointer(ptr unsafe.Pointer) *QLowEnergyController {
	var n = NewQLowEnergyControllerFromPointer(ptr)
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
func callbackQLowEnergyControllerConnected(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QLowEnergyController::connected")

	if signal := qt.GetSignal(C.GoString(ptrName), "connected"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QLowEnergyController) Connected() {
	defer qt.Recovering("QLowEnergyController::connected")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_Connected(ptr.Pointer())
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
func callbackQLowEnergyControllerDisconnected(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QLowEnergyController::disconnected")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnected"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QLowEnergyController) Disconnected() {
	defer qt.Recovering("QLowEnergyController::disconnected")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_Disconnected(ptr.Pointer())
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
func callbackQLowEnergyControllerDiscoveryFinished(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QLowEnergyController::discoveryFinished")

	if signal := qt.GetSignal(C.GoString(ptrName), "discoveryFinished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QLowEnergyController) DiscoveryFinished() {
	defer qt.Recovering("QLowEnergyController::discoveryFinished")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_DiscoveryFinished(ptr.Pointer())
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
func callbackQLowEnergyControllerError2(ptr unsafe.Pointer, ptrName *C.char, newError C.int) {
	defer qt.Recovering("callback QLowEnergyController::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(QLowEnergyController__Error))(QLowEnergyController__Error(newError))
	}

}

func (ptr *QLowEnergyController) Error2(newError QLowEnergyController__Error) {
	defer qt.Recovering("QLowEnergyController::error")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_Error2(ptr.Pointer(), C.int(newError))
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
func callbackQLowEnergyControllerStateChanged(ptr unsafe.Pointer, ptrName *C.char, state C.int) {
	defer qt.Recovering("callback QLowEnergyController::stateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stateChanged"); signal != nil {
		signal.(func(QLowEnergyController__ControllerState))(QLowEnergyController__ControllerState(state))
	}

}

func (ptr *QLowEnergyController) StateChanged(state QLowEnergyController__ControllerState) {
	defer qt.Recovering("QLowEnergyController::stateChanged")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_StateChanged(ptr.Pointer(), C.int(state))
	}
}

func NewQLowEnergyController(remoteDeviceInfo QBluetoothDeviceInfo_ITF, parent core.QObject_ITF) *QLowEnergyController {
	defer qt.Recovering("QLowEnergyController::QLowEnergyController")

	return newQLowEnergyControllerFromPointer(C.QLowEnergyController_NewQLowEnergyController(PointerFromQBluetoothDeviceInfo(remoteDeviceInfo), core.PointerFromQObject(parent)))
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
func callbackQLowEnergyControllerTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLowEnergyController::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQLowEnergyControllerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QLowEnergyController) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QLowEnergyController::timerEvent")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QLowEnergyController) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QLowEnergyController::timerEvent")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
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
func callbackQLowEnergyControllerChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLowEnergyController::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQLowEnergyControllerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QLowEnergyController) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QLowEnergyController::childEvent")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QLowEnergyController) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QLowEnergyController::childEvent")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
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
func callbackQLowEnergyControllerCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLowEnergyController::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQLowEnergyControllerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QLowEnergyController) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QLowEnergyController::customEvent")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QLowEnergyController) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QLowEnergyController::customEvent")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QLowEnergyDescriptor struct {
	ptr unsafe.Pointer
}

type QLowEnergyDescriptor_ITF interface {
	QLowEnergyDescriptor_PTR() *QLowEnergyDescriptor
}

func (p *QLowEnergyDescriptor) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QLowEnergyDescriptor) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQLowEnergyDescriptor(ptr QLowEnergyDescriptor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLowEnergyDescriptor_PTR().Pointer()
	}
	return nil
}

func NewQLowEnergyDescriptorFromPointer(ptr unsafe.Pointer) *QLowEnergyDescriptor {
	var n = new(QLowEnergyDescriptor)
	n.SetPointer(ptr)
	return n
}

func newQLowEnergyDescriptorFromPointer(ptr unsafe.Pointer) *QLowEnergyDescriptor {
	var n = NewQLowEnergyDescriptorFromPointer(ptr)
	return n
}

func (ptr *QLowEnergyDescriptor) QLowEnergyDescriptor_PTR() *QLowEnergyDescriptor {
	return ptr
}

func NewQLowEnergyDescriptor() *QLowEnergyDescriptor {
	defer qt.Recovering("QLowEnergyDescriptor::QLowEnergyDescriptor")

	return newQLowEnergyDescriptorFromPointer(C.QLowEnergyDescriptor_NewQLowEnergyDescriptor())
}

func NewQLowEnergyDescriptor2(other QLowEnergyDescriptor_ITF) *QLowEnergyDescriptor {
	defer qt.Recovering("QLowEnergyDescriptor::QLowEnergyDescriptor")

	return newQLowEnergyDescriptorFromPointer(C.QLowEnergyDescriptor_NewQLowEnergyDescriptor2(PointerFromQLowEnergyDescriptor(other)))
}

func (ptr *QLowEnergyDescriptor) IsValid() bool {
	defer qt.Recovering("QLowEnergyDescriptor::isValid")

	if ptr.Pointer() != nil {
		return C.QLowEnergyDescriptor_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLowEnergyDescriptor) Name() string {
	defer qt.Recovering("QLowEnergyDescriptor::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QLowEnergyDescriptor_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLowEnergyDescriptor) Type() QBluetoothUuid__DescriptorType {
	defer qt.Recovering("QLowEnergyDescriptor::type")

	if ptr.Pointer() != nil {
		return QBluetoothUuid__DescriptorType(C.QLowEnergyDescriptor_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyDescriptor) Value() *core.QByteArray {
	defer qt.Recovering("QLowEnergyDescriptor::value")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QLowEnergyDescriptor_Value(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLowEnergyDescriptor) DestroyQLowEnergyDescriptor() {
	defer qt.Recovering("QLowEnergyDescriptor::~QLowEnergyDescriptor")

	if ptr.Pointer() != nil {
		C.QLowEnergyDescriptor_DestroyQLowEnergyDescriptor(ptr.Pointer())
	}
}

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
	return n
}

func newQLowEnergyServiceFromPointer(ptr unsafe.Pointer) *QLowEnergyService {
	var n = NewQLowEnergyServiceFromPointer(ptr)
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
func callbackQLowEnergyServiceError2(ptr unsafe.Pointer, ptrName *C.char, newError C.int) {
	defer qt.Recovering("callback QLowEnergyService::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(QLowEnergyService__ServiceError))(QLowEnergyService__ServiceError(newError))
	}

}

func (ptr *QLowEnergyService) Error2(newError QLowEnergyService__ServiceError) {
	defer qt.Recovering("QLowEnergyService::error")

	if ptr.Pointer() != nil {
		C.QLowEnergyService_Error2(ptr.Pointer(), C.int(newError))
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
func callbackQLowEnergyServiceStateChanged(ptr unsafe.Pointer, ptrName *C.char, newState C.int) {
	defer qt.Recovering("callback QLowEnergyService::stateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stateChanged"); signal != nil {
		signal.(func(QLowEnergyService__ServiceState))(QLowEnergyService__ServiceState(newState))
	}

}

func (ptr *QLowEnergyService) StateChanged(newState QLowEnergyService__ServiceState) {
	defer qt.Recovering("QLowEnergyService::stateChanged")

	if ptr.Pointer() != nil {
		C.QLowEnergyService_StateChanged(ptr.Pointer(), C.int(newState))
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
func callbackQLowEnergyServiceTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLowEnergyService::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQLowEnergyServiceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QLowEnergyService) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QLowEnergyService::timerEvent")

	if ptr.Pointer() != nil {
		C.QLowEnergyService_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QLowEnergyService) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QLowEnergyService::timerEvent")

	if ptr.Pointer() != nil {
		C.QLowEnergyService_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
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
func callbackQLowEnergyServiceChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLowEnergyService::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQLowEnergyServiceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QLowEnergyService) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QLowEnergyService::childEvent")

	if ptr.Pointer() != nil {
		C.QLowEnergyService_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QLowEnergyService) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QLowEnergyService::childEvent")

	if ptr.Pointer() != nil {
		C.QLowEnergyService_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
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
func callbackQLowEnergyServiceCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLowEnergyService::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQLowEnergyServiceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QLowEnergyService) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QLowEnergyService::customEvent")

	if ptr.Pointer() != nil {
		C.QLowEnergyService_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QLowEnergyService) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QLowEnergyService::customEvent")

	if ptr.Pointer() != nil {
		C.QLowEnergyService_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
