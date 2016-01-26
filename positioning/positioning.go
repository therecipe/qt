package positioning

//#include "positioning.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QGeoAddress struct {
	ptr unsafe.Pointer
}

type QGeoAddress_ITF interface {
	QGeoAddress_PTR() *QGeoAddress
}

func (p *QGeoAddress) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QGeoAddress) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQGeoAddress(ptr QGeoAddress_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoAddress_PTR().Pointer()
	}
	return nil
}

func NewQGeoAddressFromPointer(ptr unsafe.Pointer) *QGeoAddress {
	var n = new(QGeoAddress)
	n.SetPointer(ptr)
	return n
}

func newQGeoAddressFromPointer(ptr unsafe.Pointer) *QGeoAddress {
	var n = NewQGeoAddressFromPointer(ptr)
	return n
}

func (ptr *QGeoAddress) QGeoAddress_PTR() *QGeoAddress {
	return ptr
}

func NewQGeoAddress() *QGeoAddress {
	defer qt.Recovering("QGeoAddress::QGeoAddress")

	return newQGeoAddressFromPointer(C.QGeoAddress_NewQGeoAddress())
}

func NewQGeoAddress2(other QGeoAddress_ITF) *QGeoAddress {
	defer qt.Recovering("QGeoAddress::QGeoAddress")

	return newQGeoAddressFromPointer(C.QGeoAddress_NewQGeoAddress2(PointerFromQGeoAddress(other)))
}

func (ptr *QGeoAddress) City() string {
	defer qt.Recovering("QGeoAddress::city")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoAddress_City(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoAddress) Clear() {
	defer qt.Recovering("QGeoAddress::clear")

	if ptr.Pointer() != nil {
		C.QGeoAddress_Clear(ptr.Pointer())
	}
}

func (ptr *QGeoAddress) Country() string {
	defer qt.Recovering("QGeoAddress::country")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoAddress_Country(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoAddress) CountryCode() string {
	defer qt.Recovering("QGeoAddress::countryCode")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoAddress_CountryCode(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoAddress) County() string {
	defer qt.Recovering("QGeoAddress::county")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoAddress_County(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoAddress) District() string {
	defer qt.Recovering("QGeoAddress::district")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoAddress_District(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoAddress) IsEmpty() bool {
	defer qt.Recovering("QGeoAddress::isEmpty")

	if ptr.Pointer() != nil {
		return C.QGeoAddress_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGeoAddress) IsTextGenerated() bool {
	defer qt.Recovering("QGeoAddress::isTextGenerated")

	if ptr.Pointer() != nil {
		return C.QGeoAddress_IsTextGenerated(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGeoAddress) PostalCode() string {
	defer qt.Recovering("QGeoAddress::postalCode")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoAddress_PostalCode(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoAddress) SetCity(city string) {
	defer qt.Recovering("QGeoAddress::setCity")

	if ptr.Pointer() != nil {
		C.QGeoAddress_SetCity(ptr.Pointer(), C.CString(city))
	}
}

func (ptr *QGeoAddress) SetCountry(country string) {
	defer qt.Recovering("QGeoAddress::setCountry")

	if ptr.Pointer() != nil {
		C.QGeoAddress_SetCountry(ptr.Pointer(), C.CString(country))
	}
}

func (ptr *QGeoAddress) SetCountryCode(countryCode string) {
	defer qt.Recovering("QGeoAddress::setCountryCode")

	if ptr.Pointer() != nil {
		C.QGeoAddress_SetCountryCode(ptr.Pointer(), C.CString(countryCode))
	}
}

func (ptr *QGeoAddress) SetCounty(county string) {
	defer qt.Recovering("QGeoAddress::setCounty")

	if ptr.Pointer() != nil {
		C.QGeoAddress_SetCounty(ptr.Pointer(), C.CString(county))
	}
}

func (ptr *QGeoAddress) SetDistrict(district string) {
	defer qt.Recovering("QGeoAddress::setDistrict")

	if ptr.Pointer() != nil {
		C.QGeoAddress_SetDistrict(ptr.Pointer(), C.CString(district))
	}
}

func (ptr *QGeoAddress) SetPostalCode(postalCode string) {
	defer qt.Recovering("QGeoAddress::setPostalCode")

	if ptr.Pointer() != nil {
		C.QGeoAddress_SetPostalCode(ptr.Pointer(), C.CString(postalCode))
	}
}

func (ptr *QGeoAddress) SetState(state string) {
	defer qt.Recovering("QGeoAddress::setState")

	if ptr.Pointer() != nil {
		C.QGeoAddress_SetState(ptr.Pointer(), C.CString(state))
	}
}

func (ptr *QGeoAddress) SetStreet(street string) {
	defer qt.Recovering("QGeoAddress::setStreet")

	if ptr.Pointer() != nil {
		C.QGeoAddress_SetStreet(ptr.Pointer(), C.CString(street))
	}
}

func (ptr *QGeoAddress) SetText(text string) {
	defer qt.Recovering("QGeoAddress::setText")

	if ptr.Pointer() != nil {
		C.QGeoAddress_SetText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QGeoAddress) State() string {
	defer qt.Recovering("QGeoAddress::state")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoAddress_State(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoAddress) Street() string {
	defer qt.Recovering("QGeoAddress::street")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoAddress_Street(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoAddress) Text() string {
	defer qt.Recovering("QGeoAddress::text")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoAddress_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoAddress) DestroyQGeoAddress() {
	defer qt.Recovering("QGeoAddress::~QGeoAddress")

	if ptr.Pointer() != nil {
		C.QGeoAddress_DestroyQGeoAddress(ptr.Pointer())
	}
}

type QGeoAreaMonitorInfo struct {
	ptr unsafe.Pointer
}

type QGeoAreaMonitorInfo_ITF interface {
	QGeoAreaMonitorInfo_PTR() *QGeoAreaMonitorInfo
}

func (p *QGeoAreaMonitorInfo) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QGeoAreaMonitorInfo) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQGeoAreaMonitorInfo(ptr QGeoAreaMonitorInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoAreaMonitorInfo_PTR().Pointer()
	}
	return nil
}

func NewQGeoAreaMonitorInfoFromPointer(ptr unsafe.Pointer) *QGeoAreaMonitorInfo {
	var n = new(QGeoAreaMonitorInfo)
	n.SetPointer(ptr)
	return n
}

func newQGeoAreaMonitorInfoFromPointer(ptr unsafe.Pointer) *QGeoAreaMonitorInfo {
	var n = NewQGeoAreaMonitorInfoFromPointer(ptr)
	return n
}

func (ptr *QGeoAreaMonitorInfo) QGeoAreaMonitorInfo_PTR() *QGeoAreaMonitorInfo {
	return ptr
}

func NewQGeoAreaMonitorInfo2(other QGeoAreaMonitorInfo_ITF) *QGeoAreaMonitorInfo {
	defer qt.Recovering("QGeoAreaMonitorInfo::QGeoAreaMonitorInfo")

	return newQGeoAreaMonitorInfoFromPointer(C.QGeoAreaMonitorInfo_NewQGeoAreaMonitorInfo2(PointerFromQGeoAreaMonitorInfo(other)))
}

func NewQGeoAreaMonitorInfo(name string) *QGeoAreaMonitorInfo {
	defer qt.Recovering("QGeoAreaMonitorInfo::QGeoAreaMonitorInfo")

	return newQGeoAreaMonitorInfoFromPointer(C.QGeoAreaMonitorInfo_NewQGeoAreaMonitorInfo(C.CString(name)))
}

func (ptr *QGeoAreaMonitorInfo) Expiration() *core.QDateTime {
	defer qt.Recovering("QGeoAreaMonitorInfo::expiration")

	if ptr.Pointer() != nil {
		return core.NewQDateTimeFromPointer(C.QGeoAreaMonitorInfo_Expiration(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGeoAreaMonitorInfo) Identifier() string {
	defer qt.Recovering("QGeoAreaMonitorInfo::identifier")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoAreaMonitorInfo_Identifier(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoAreaMonitorInfo) IsPersistent() bool {
	defer qt.Recovering("QGeoAreaMonitorInfo::isPersistent")

	if ptr.Pointer() != nil {
		return C.QGeoAreaMonitorInfo_IsPersistent(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGeoAreaMonitorInfo) IsValid() bool {
	defer qt.Recovering("QGeoAreaMonitorInfo::isValid")

	if ptr.Pointer() != nil {
		return C.QGeoAreaMonitorInfo_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGeoAreaMonitorInfo) Name() string {
	defer qt.Recovering("QGeoAreaMonitorInfo::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoAreaMonitorInfo_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoAreaMonitorInfo) SetArea(newShape QGeoShape_ITF) {
	defer qt.Recovering("QGeoAreaMonitorInfo::setArea")

	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorInfo_SetArea(ptr.Pointer(), PointerFromQGeoShape(newShape))
	}
}

func (ptr *QGeoAreaMonitorInfo) SetExpiration(expiry core.QDateTime_ITF) {
	defer qt.Recovering("QGeoAreaMonitorInfo::setExpiration")

	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorInfo_SetExpiration(ptr.Pointer(), core.PointerFromQDateTime(expiry))
	}
}

func (ptr *QGeoAreaMonitorInfo) SetName(name string) {
	defer qt.Recovering("QGeoAreaMonitorInfo::setName")

	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorInfo_SetName(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QGeoAreaMonitorInfo) SetPersistent(isPersistent bool) {
	defer qt.Recovering("QGeoAreaMonitorInfo::setPersistent")

	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorInfo_SetPersistent(ptr.Pointer(), C.int(qt.GoBoolToInt(isPersistent)))
	}
}

func (ptr *QGeoAreaMonitorInfo) DestroyQGeoAreaMonitorInfo() {
	defer qt.Recovering("QGeoAreaMonitorInfo::~QGeoAreaMonitorInfo")

	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorInfo_DestroyQGeoAreaMonitorInfo(ptr.Pointer())
	}
}

type QGeoAreaMonitorSource struct {
	core.QObject
}

type QGeoAreaMonitorSource_ITF interface {
	core.QObject_ITF
	QGeoAreaMonitorSource_PTR() *QGeoAreaMonitorSource
}

func PointerFromQGeoAreaMonitorSource(ptr QGeoAreaMonitorSource_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoAreaMonitorSource_PTR().Pointer()
	}
	return nil
}

func NewQGeoAreaMonitorSourceFromPointer(ptr unsafe.Pointer) *QGeoAreaMonitorSource {
	var n = new(QGeoAreaMonitorSource)
	n.SetPointer(ptr)
	return n
}

func newQGeoAreaMonitorSourceFromPointer(ptr unsafe.Pointer) *QGeoAreaMonitorSource {
	var n = NewQGeoAreaMonitorSourceFromPointer(ptr)
	for len(n.ObjectName()) < len("QGeoAreaMonitorSource_") {
		n.SetObjectName("QGeoAreaMonitorSource_" + qt.Identifier())
	}
	return n
}

func (ptr *QGeoAreaMonitorSource) QGeoAreaMonitorSource_PTR() *QGeoAreaMonitorSource {
	return ptr
}

//QGeoAreaMonitorSource::AreaMonitorFeature
type QGeoAreaMonitorSource__AreaMonitorFeature int64

const (
	QGeoAreaMonitorSource__PersistentAreaMonitorFeature = QGeoAreaMonitorSource__AreaMonitorFeature(0x00000001)
	QGeoAreaMonitorSource__AnyAreaMonitorFeature        = QGeoAreaMonitorSource__AreaMonitorFeature(0xffffffff)
)

//QGeoAreaMonitorSource::Error
type QGeoAreaMonitorSource__Error int64

const (
	QGeoAreaMonitorSource__AccessError              = QGeoAreaMonitorSource__Error(0)
	QGeoAreaMonitorSource__InsufficientPositionInfo = QGeoAreaMonitorSource__Error(1)
	QGeoAreaMonitorSource__UnknownSourceError       = QGeoAreaMonitorSource__Error(2)
	QGeoAreaMonitorSource__NoError                  = QGeoAreaMonitorSource__Error(3)
)

func QGeoAreaMonitorSource_AvailableSources() []string {
	defer qt.Recovering("QGeoAreaMonitorSource::availableSources")

	return strings.Split(C.GoString(C.QGeoAreaMonitorSource_QGeoAreaMonitorSource_AvailableSources()), "|")
}

func QGeoAreaMonitorSource_CreateDefaultSource(parent core.QObject_ITF) *QGeoAreaMonitorSource {
	defer qt.Recovering("QGeoAreaMonitorSource::createDefaultSource")

	return NewQGeoAreaMonitorSourceFromPointer(C.QGeoAreaMonitorSource_QGeoAreaMonitorSource_CreateDefaultSource(core.PointerFromQObject(parent)))
}

func QGeoAreaMonitorSource_CreateSource(sourceName string, parent core.QObject_ITF) *QGeoAreaMonitorSource {
	defer qt.Recovering("QGeoAreaMonitorSource::createSource")

	return NewQGeoAreaMonitorSourceFromPointer(C.QGeoAreaMonitorSource_QGeoAreaMonitorSource_CreateSource(C.CString(sourceName), core.PointerFromQObject(parent)))
}

func (ptr *QGeoAreaMonitorSource) ConnectError2(f func(areaMonitoringError QGeoAreaMonitorSource__Error)) {
	defer qt.Recovering("connect QGeoAreaMonitorSource::error")

	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error2", f)
	}
}

func (ptr *QGeoAreaMonitorSource) DisconnectError2() {
	defer qt.Recovering("disconnect QGeoAreaMonitorSource::error")

	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error2")
	}
}

//export callbackQGeoAreaMonitorSourceError2
func callbackQGeoAreaMonitorSourceError2(ptr unsafe.Pointer, ptrName *C.char, areaMonitoringError C.int) {
	defer qt.Recovering("callback QGeoAreaMonitorSource::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(QGeoAreaMonitorSource__Error))(QGeoAreaMonitorSource__Error(areaMonitoringError))
	}

}

func (ptr *QGeoAreaMonitorSource) Error2(areaMonitoringError QGeoAreaMonitorSource__Error) {
	defer qt.Recovering("QGeoAreaMonitorSource::error")

	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_Error2(ptr.Pointer(), C.int(areaMonitoringError))
	}
}

func (ptr *QGeoAreaMonitorSource) Error() QGeoAreaMonitorSource__Error {
	defer qt.Recovering("QGeoAreaMonitorSource::error")

	if ptr.Pointer() != nil {
		return QGeoAreaMonitorSource__Error(C.QGeoAreaMonitorSource_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoAreaMonitorSource) PositionInfoSource() *QGeoPositionInfoSource {
	defer qt.Recovering("QGeoAreaMonitorSource::positionInfoSource")

	if ptr.Pointer() != nil {
		return NewQGeoPositionInfoSourceFromPointer(C.QGeoAreaMonitorSource_PositionInfoSource(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGeoAreaMonitorSource) RequestUpdate(monitor QGeoAreaMonitorInfo_ITF, signal string) bool {
	defer qt.Recovering("QGeoAreaMonitorSource::requestUpdate")

	if ptr.Pointer() != nil {
		return C.QGeoAreaMonitorSource_RequestUpdate(ptr.Pointer(), PointerFromQGeoAreaMonitorInfo(monitor), C.CString(signal)) != 0
	}
	return false
}

func (ptr *QGeoAreaMonitorSource) ConnectSetPositionInfoSource(f func(newSource *QGeoPositionInfoSource)) {
	defer qt.Recovering("connect QGeoAreaMonitorSource::setPositionInfoSource")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setPositionInfoSource", f)
	}
}

func (ptr *QGeoAreaMonitorSource) DisconnectSetPositionInfoSource() {
	defer qt.Recovering("disconnect QGeoAreaMonitorSource::setPositionInfoSource")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setPositionInfoSource")
	}
}

//export callbackQGeoAreaMonitorSourceSetPositionInfoSource
func callbackQGeoAreaMonitorSourceSetPositionInfoSource(ptr unsafe.Pointer, ptrName *C.char, newSource unsafe.Pointer) {
	defer qt.Recovering("callback QGeoAreaMonitorSource::setPositionInfoSource")

	if signal := qt.GetSignal(C.GoString(ptrName), "setPositionInfoSource"); signal != nil {
		signal.(func(*QGeoPositionInfoSource))(NewQGeoPositionInfoSourceFromPointer(newSource))
	} else {
		NewQGeoAreaMonitorSourceFromPointer(ptr).SetPositionInfoSourceDefault(NewQGeoPositionInfoSourceFromPointer(newSource))
	}
}

func (ptr *QGeoAreaMonitorSource) SetPositionInfoSource(newSource QGeoPositionInfoSource_ITF) {
	defer qt.Recovering("QGeoAreaMonitorSource::setPositionInfoSource")

	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_SetPositionInfoSource(ptr.Pointer(), PointerFromQGeoPositionInfoSource(newSource))
	}
}

func (ptr *QGeoAreaMonitorSource) SetPositionInfoSourceDefault(newSource QGeoPositionInfoSource_ITF) {
	defer qt.Recovering("QGeoAreaMonitorSource::setPositionInfoSource")

	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_SetPositionInfoSourceDefault(ptr.Pointer(), PointerFromQGeoPositionInfoSource(newSource))
	}
}

func (ptr *QGeoAreaMonitorSource) SourceName() string {
	defer qt.Recovering("QGeoAreaMonitorSource::sourceName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoAreaMonitorSource_SourceName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoAreaMonitorSource) StartMonitoring(monitor QGeoAreaMonitorInfo_ITF) bool {
	defer qt.Recovering("QGeoAreaMonitorSource::startMonitoring")

	if ptr.Pointer() != nil {
		return C.QGeoAreaMonitorSource_StartMonitoring(ptr.Pointer(), PointerFromQGeoAreaMonitorInfo(monitor)) != 0
	}
	return false
}

func (ptr *QGeoAreaMonitorSource) StopMonitoring(monitor QGeoAreaMonitorInfo_ITF) bool {
	defer qt.Recovering("QGeoAreaMonitorSource::stopMonitoring")

	if ptr.Pointer() != nil {
		return C.QGeoAreaMonitorSource_StopMonitoring(ptr.Pointer(), PointerFromQGeoAreaMonitorInfo(monitor)) != 0
	}
	return false
}

func (ptr *QGeoAreaMonitorSource) SupportedAreaMonitorFeatures() QGeoAreaMonitorSource__AreaMonitorFeature {
	defer qt.Recovering("QGeoAreaMonitorSource::supportedAreaMonitorFeatures")

	if ptr.Pointer() != nil {
		return QGeoAreaMonitorSource__AreaMonitorFeature(C.QGeoAreaMonitorSource_SupportedAreaMonitorFeatures(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoAreaMonitorSource) DestroyQGeoAreaMonitorSource() {
	defer qt.Recovering("QGeoAreaMonitorSource::~QGeoAreaMonitorSource")

	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_DestroyQGeoAreaMonitorSource(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGeoAreaMonitorSource) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QGeoAreaMonitorSource::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QGeoAreaMonitorSource) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QGeoAreaMonitorSource::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQGeoAreaMonitorSourceTimerEvent
func callbackQGeoAreaMonitorSourceTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGeoAreaMonitorSource::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGeoAreaMonitorSourceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGeoAreaMonitorSource) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGeoAreaMonitorSource::timerEvent")

	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGeoAreaMonitorSource) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGeoAreaMonitorSource::timerEvent")

	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGeoAreaMonitorSource) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QGeoAreaMonitorSource::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QGeoAreaMonitorSource) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QGeoAreaMonitorSource::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQGeoAreaMonitorSourceChildEvent
func callbackQGeoAreaMonitorSourceChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGeoAreaMonitorSource::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQGeoAreaMonitorSourceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QGeoAreaMonitorSource) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QGeoAreaMonitorSource::childEvent")

	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGeoAreaMonitorSource) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QGeoAreaMonitorSource::childEvent")

	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGeoAreaMonitorSource) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGeoAreaMonitorSource::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QGeoAreaMonitorSource) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QGeoAreaMonitorSource::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQGeoAreaMonitorSourceCustomEvent
func callbackQGeoAreaMonitorSourceCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGeoAreaMonitorSource::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGeoAreaMonitorSourceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGeoAreaMonitorSource) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QGeoAreaMonitorSource::customEvent")

	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGeoAreaMonitorSource) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QGeoAreaMonitorSource::customEvent")

	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QGeoCircle struct {
	QGeoShape
}

type QGeoCircle_ITF interface {
	QGeoShape_ITF
	QGeoCircle_PTR() *QGeoCircle
}

func PointerFromQGeoCircle(ptr QGeoCircle_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoCircle_PTR().Pointer()
	}
	return nil
}

func NewQGeoCircleFromPointer(ptr unsafe.Pointer) *QGeoCircle {
	var n = new(QGeoCircle)
	n.SetPointer(ptr)
	return n
}

func newQGeoCircleFromPointer(ptr unsafe.Pointer) *QGeoCircle {
	var n = NewQGeoCircleFromPointer(ptr)
	return n
}

func (ptr *QGeoCircle) QGeoCircle_PTR() *QGeoCircle {
	return ptr
}

func NewQGeoCircle() *QGeoCircle {
	defer qt.Recovering("QGeoCircle::QGeoCircle")

	return newQGeoCircleFromPointer(C.QGeoCircle_NewQGeoCircle())
}

func NewQGeoCircle3(other QGeoCircle_ITF) *QGeoCircle {
	defer qt.Recovering("QGeoCircle::QGeoCircle")

	return newQGeoCircleFromPointer(C.QGeoCircle_NewQGeoCircle3(PointerFromQGeoCircle(other)))
}

func NewQGeoCircle2(center QGeoCoordinate_ITF, radius float64) *QGeoCircle {
	defer qt.Recovering("QGeoCircle::QGeoCircle")

	return newQGeoCircleFromPointer(C.QGeoCircle_NewQGeoCircle2(PointerFromQGeoCoordinate(center), C.double(radius)))
}

func NewQGeoCircle4(other QGeoShape_ITF) *QGeoCircle {
	defer qt.Recovering("QGeoCircle::QGeoCircle")

	return newQGeoCircleFromPointer(C.QGeoCircle_NewQGeoCircle4(PointerFromQGeoShape(other)))
}

func (ptr *QGeoCircle) Radius() float64 {
	defer qt.Recovering("QGeoCircle::radius")

	if ptr.Pointer() != nil {
		return float64(C.QGeoCircle_Radius(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoCircle) SetCenter(center QGeoCoordinate_ITF) {
	defer qt.Recovering("QGeoCircle::setCenter")

	if ptr.Pointer() != nil {
		C.QGeoCircle_SetCenter(ptr.Pointer(), PointerFromQGeoCoordinate(center))
	}
}

func (ptr *QGeoCircle) SetRadius(radius float64) {
	defer qt.Recovering("QGeoCircle::setRadius")

	if ptr.Pointer() != nil {
		C.QGeoCircle_SetRadius(ptr.Pointer(), C.double(radius))
	}
}

func (ptr *QGeoCircle) ToString() string {
	defer qt.Recovering("QGeoCircle::toString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoCircle_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoCircle) DestroyQGeoCircle() {
	defer qt.Recovering("QGeoCircle::~QGeoCircle")

	if ptr.Pointer() != nil {
		C.QGeoCircle_DestroyQGeoCircle(ptr.Pointer())
	}
}

type QGeoCoordinate struct {
	ptr unsafe.Pointer
}

type QGeoCoordinate_ITF interface {
	QGeoCoordinate_PTR() *QGeoCoordinate
}

func (p *QGeoCoordinate) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QGeoCoordinate) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQGeoCoordinate(ptr QGeoCoordinate_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoCoordinate_PTR().Pointer()
	}
	return nil
}

func NewQGeoCoordinateFromPointer(ptr unsafe.Pointer) *QGeoCoordinate {
	var n = new(QGeoCoordinate)
	n.SetPointer(ptr)
	return n
}

func newQGeoCoordinateFromPointer(ptr unsafe.Pointer) *QGeoCoordinate {
	var n = NewQGeoCoordinateFromPointer(ptr)
	return n
}

func (ptr *QGeoCoordinate) QGeoCoordinate_PTR() *QGeoCoordinate {
	return ptr
}

//QGeoCoordinate::CoordinateFormat
type QGeoCoordinate__CoordinateFormat int64

const (
	QGeoCoordinate__Degrees                             = QGeoCoordinate__CoordinateFormat(0)
	QGeoCoordinate__DegreesWithHemisphere               = QGeoCoordinate__CoordinateFormat(1)
	QGeoCoordinate__DegreesMinutes                      = QGeoCoordinate__CoordinateFormat(2)
	QGeoCoordinate__DegreesMinutesWithHemisphere        = QGeoCoordinate__CoordinateFormat(3)
	QGeoCoordinate__DegreesMinutesSeconds               = QGeoCoordinate__CoordinateFormat(4)
	QGeoCoordinate__DegreesMinutesSecondsWithHemisphere = QGeoCoordinate__CoordinateFormat(5)
)

//QGeoCoordinate::CoordinateType
type QGeoCoordinate__CoordinateType int64

const (
	QGeoCoordinate__InvalidCoordinate = QGeoCoordinate__CoordinateType(0)
	QGeoCoordinate__Coordinate2D      = QGeoCoordinate__CoordinateType(1)
	QGeoCoordinate__Coordinate3D      = QGeoCoordinate__CoordinateType(2)
)

func NewQGeoCoordinate() *QGeoCoordinate {
	defer qt.Recovering("QGeoCoordinate::QGeoCoordinate")

	return newQGeoCoordinateFromPointer(C.QGeoCoordinate_NewQGeoCoordinate())
}

func NewQGeoCoordinate4(other QGeoCoordinate_ITF) *QGeoCoordinate {
	defer qt.Recovering("QGeoCoordinate::QGeoCoordinate")

	return newQGeoCoordinateFromPointer(C.QGeoCoordinate_NewQGeoCoordinate4(PointerFromQGeoCoordinate(other)))
}

func (ptr *QGeoCoordinate) AzimuthTo(other QGeoCoordinate_ITF) float64 {
	defer qt.Recovering("QGeoCoordinate::azimuthTo")

	if ptr.Pointer() != nil {
		return float64(C.QGeoCoordinate_AzimuthTo(ptr.Pointer(), PointerFromQGeoCoordinate(other)))
	}
	return 0
}

func (ptr *QGeoCoordinate) DistanceTo(other QGeoCoordinate_ITF) float64 {
	defer qt.Recovering("QGeoCoordinate::distanceTo")

	if ptr.Pointer() != nil {
		return float64(C.QGeoCoordinate_DistanceTo(ptr.Pointer(), PointerFromQGeoCoordinate(other)))
	}
	return 0
}

func (ptr *QGeoCoordinate) IsValid() bool {
	defer qt.Recovering("QGeoCoordinate::isValid")

	if ptr.Pointer() != nil {
		return C.QGeoCoordinate_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGeoCoordinate) ToString(format QGeoCoordinate__CoordinateFormat) string {
	defer qt.Recovering("QGeoCoordinate::toString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoCoordinate_ToString(ptr.Pointer(), C.int(format)))
	}
	return ""
}

func (ptr *QGeoCoordinate) Type() QGeoCoordinate__CoordinateType {
	defer qt.Recovering("QGeoCoordinate::type")

	if ptr.Pointer() != nil {
		return QGeoCoordinate__CoordinateType(C.QGeoCoordinate_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoCoordinate) DestroyQGeoCoordinate() {
	defer qt.Recovering("QGeoCoordinate::~QGeoCoordinate")

	if ptr.Pointer() != nil {
		C.QGeoCoordinate_DestroyQGeoCoordinate(ptr.Pointer())
	}
}

type QGeoLocation struct {
	ptr unsafe.Pointer
}

type QGeoLocation_ITF interface {
	QGeoLocation_PTR() *QGeoLocation
}

func (p *QGeoLocation) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QGeoLocation) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQGeoLocation(ptr QGeoLocation_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoLocation_PTR().Pointer()
	}
	return nil
}

func NewQGeoLocationFromPointer(ptr unsafe.Pointer) *QGeoLocation {
	var n = new(QGeoLocation)
	n.SetPointer(ptr)
	return n
}

func newQGeoLocationFromPointer(ptr unsafe.Pointer) *QGeoLocation {
	var n = NewQGeoLocationFromPointer(ptr)
	return n
}

func (ptr *QGeoLocation) QGeoLocation_PTR() *QGeoLocation {
	return ptr
}

type QGeoPositionInfo struct {
	ptr unsafe.Pointer
}

type QGeoPositionInfo_ITF interface {
	QGeoPositionInfo_PTR() *QGeoPositionInfo
}

func (p *QGeoPositionInfo) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QGeoPositionInfo) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQGeoPositionInfo(ptr QGeoPositionInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoPositionInfo_PTR().Pointer()
	}
	return nil
}

func NewQGeoPositionInfoFromPointer(ptr unsafe.Pointer) *QGeoPositionInfo {
	var n = new(QGeoPositionInfo)
	n.SetPointer(ptr)
	return n
}

func newQGeoPositionInfoFromPointer(ptr unsafe.Pointer) *QGeoPositionInfo {
	var n = NewQGeoPositionInfoFromPointer(ptr)
	return n
}

func (ptr *QGeoPositionInfo) QGeoPositionInfo_PTR() *QGeoPositionInfo {
	return ptr
}

//QGeoPositionInfo::Attribute
type QGeoPositionInfo__Attribute int64

const (
	QGeoPositionInfo__Direction          = QGeoPositionInfo__Attribute(0)
	QGeoPositionInfo__GroundSpeed        = QGeoPositionInfo__Attribute(1)
	QGeoPositionInfo__VerticalSpeed      = QGeoPositionInfo__Attribute(2)
	QGeoPositionInfo__MagneticVariation  = QGeoPositionInfo__Attribute(3)
	QGeoPositionInfo__HorizontalAccuracy = QGeoPositionInfo__Attribute(4)
	QGeoPositionInfo__VerticalAccuracy   = QGeoPositionInfo__Attribute(5)
)

func NewQGeoPositionInfo() *QGeoPositionInfo {
	defer qt.Recovering("QGeoPositionInfo::QGeoPositionInfo")

	return newQGeoPositionInfoFromPointer(C.QGeoPositionInfo_NewQGeoPositionInfo())
}

func NewQGeoPositionInfo2(coordinate QGeoCoordinate_ITF, timestamp core.QDateTime_ITF) *QGeoPositionInfo {
	defer qt.Recovering("QGeoPositionInfo::QGeoPositionInfo")

	return newQGeoPositionInfoFromPointer(C.QGeoPositionInfo_NewQGeoPositionInfo2(PointerFromQGeoCoordinate(coordinate), core.PointerFromQDateTime(timestamp)))
}

func NewQGeoPositionInfo3(other QGeoPositionInfo_ITF) *QGeoPositionInfo {
	defer qt.Recovering("QGeoPositionInfo::QGeoPositionInfo")

	return newQGeoPositionInfoFromPointer(C.QGeoPositionInfo_NewQGeoPositionInfo3(PointerFromQGeoPositionInfo(other)))
}

func (ptr *QGeoPositionInfo) Attribute(attribute QGeoPositionInfo__Attribute) float64 {
	defer qt.Recovering("QGeoPositionInfo::attribute")

	if ptr.Pointer() != nil {
		return float64(C.QGeoPositionInfo_Attribute(ptr.Pointer(), C.int(attribute)))
	}
	return 0
}

func (ptr *QGeoPositionInfo) HasAttribute(attribute QGeoPositionInfo__Attribute) bool {
	defer qt.Recovering("QGeoPositionInfo::hasAttribute")

	if ptr.Pointer() != nil {
		return C.QGeoPositionInfo_HasAttribute(ptr.Pointer(), C.int(attribute)) != 0
	}
	return false
}

func (ptr *QGeoPositionInfo) IsValid() bool {
	defer qt.Recovering("QGeoPositionInfo::isValid")

	if ptr.Pointer() != nil {
		return C.QGeoPositionInfo_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGeoPositionInfo) RemoveAttribute(attribute QGeoPositionInfo__Attribute) {
	defer qt.Recovering("QGeoPositionInfo::removeAttribute")

	if ptr.Pointer() != nil {
		C.QGeoPositionInfo_RemoveAttribute(ptr.Pointer(), C.int(attribute))
	}
}

func (ptr *QGeoPositionInfo) SetAttribute(attribute QGeoPositionInfo__Attribute, value float64) {
	defer qt.Recovering("QGeoPositionInfo::setAttribute")

	if ptr.Pointer() != nil {
		C.QGeoPositionInfo_SetAttribute(ptr.Pointer(), C.int(attribute), C.double(value))
	}
}

func (ptr *QGeoPositionInfo) SetCoordinate(coordinate QGeoCoordinate_ITF) {
	defer qt.Recovering("QGeoPositionInfo::setCoordinate")

	if ptr.Pointer() != nil {
		C.QGeoPositionInfo_SetCoordinate(ptr.Pointer(), PointerFromQGeoCoordinate(coordinate))
	}
}

func (ptr *QGeoPositionInfo) SetTimestamp(timestamp core.QDateTime_ITF) {
	defer qt.Recovering("QGeoPositionInfo::setTimestamp")

	if ptr.Pointer() != nil {
		C.QGeoPositionInfo_SetTimestamp(ptr.Pointer(), core.PointerFromQDateTime(timestamp))
	}
}

func (ptr *QGeoPositionInfo) Timestamp() *core.QDateTime {
	defer qt.Recovering("QGeoPositionInfo::timestamp")

	if ptr.Pointer() != nil {
		return core.NewQDateTimeFromPointer(C.QGeoPositionInfo_Timestamp(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGeoPositionInfo) DestroyQGeoPositionInfo() {
	defer qt.Recovering("QGeoPositionInfo::~QGeoPositionInfo")

	if ptr.Pointer() != nil {
		C.QGeoPositionInfo_DestroyQGeoPositionInfo(ptr.Pointer())
	}
}

type QGeoPositionInfoSource struct {
	core.QObject
}

type QGeoPositionInfoSource_ITF interface {
	core.QObject_ITF
	QGeoPositionInfoSource_PTR() *QGeoPositionInfoSource
}

func PointerFromQGeoPositionInfoSource(ptr QGeoPositionInfoSource_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoPositionInfoSource_PTR().Pointer()
	}
	return nil
}

func NewQGeoPositionInfoSourceFromPointer(ptr unsafe.Pointer) *QGeoPositionInfoSource {
	var n = new(QGeoPositionInfoSource)
	n.SetPointer(ptr)
	return n
}

func newQGeoPositionInfoSourceFromPointer(ptr unsafe.Pointer) *QGeoPositionInfoSource {
	var n = NewQGeoPositionInfoSourceFromPointer(ptr)
	for len(n.ObjectName()) < len("QGeoPositionInfoSource_") {
		n.SetObjectName("QGeoPositionInfoSource_" + qt.Identifier())
	}
	return n
}

func (ptr *QGeoPositionInfoSource) QGeoPositionInfoSource_PTR() *QGeoPositionInfoSource {
	return ptr
}

//QGeoPositionInfoSource::Error
type QGeoPositionInfoSource__Error int64

const (
	QGeoPositionInfoSource__AccessError        = QGeoPositionInfoSource__Error(0)
	QGeoPositionInfoSource__ClosedError        = QGeoPositionInfoSource__Error(1)
	QGeoPositionInfoSource__UnknownSourceError = QGeoPositionInfoSource__Error(2)
	QGeoPositionInfoSource__NoError            = QGeoPositionInfoSource__Error(3)
)

//QGeoPositionInfoSource::PositioningMethod
type QGeoPositionInfoSource__PositioningMethod int64

const (
	QGeoPositionInfoSource__NoPositioningMethods           = QGeoPositionInfoSource__PositioningMethod(0x00000000)
	QGeoPositionInfoSource__SatellitePositioningMethods    = QGeoPositionInfoSource__PositioningMethod(0x000000ff)
	QGeoPositionInfoSource__NonSatellitePositioningMethods = QGeoPositionInfoSource__PositioningMethod(0xffffff00)
	QGeoPositionInfoSource__AllPositioningMethods          = QGeoPositionInfoSource__PositioningMethod(0xffffffff)
)

func (ptr *QGeoPositionInfoSource) ConnectSetUpdateInterval(f func(msec int)) {
	defer qt.Recovering("connect QGeoPositionInfoSource::setUpdateInterval")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setUpdateInterval", f)
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectSetUpdateInterval() {
	defer qt.Recovering("disconnect QGeoPositionInfoSource::setUpdateInterval")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setUpdateInterval")
	}
}

//export callbackQGeoPositionInfoSourceSetUpdateInterval
func callbackQGeoPositionInfoSourceSetUpdateInterval(ptr unsafe.Pointer, ptrName *C.char, msec C.int) {
	defer qt.Recovering("callback QGeoPositionInfoSource::setUpdateInterval")

	if signal := qt.GetSignal(C.GoString(ptrName), "setUpdateInterval"); signal != nil {
		signal.(func(int))(int(msec))
	} else {
		NewQGeoPositionInfoSourceFromPointer(ptr).SetUpdateIntervalDefault(int(msec))
	}
}

func (ptr *QGeoPositionInfoSource) SetUpdateInterval(msec int) {
	defer qt.Recovering("QGeoPositionInfoSource::setUpdateInterval")

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_SetUpdateInterval(ptr.Pointer(), C.int(msec))
	}
}

func (ptr *QGeoPositionInfoSource) SetUpdateIntervalDefault(msec int) {
	defer qt.Recovering("QGeoPositionInfoSource::setUpdateInterval")

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_SetUpdateIntervalDefault(ptr.Pointer(), C.int(msec))
	}
}

func (ptr *QGeoPositionInfoSource) SourceName() string {
	defer qt.Recovering("QGeoPositionInfoSource::sourceName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoPositionInfoSource_SourceName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoPositionInfoSource) UpdateInterval() int {
	defer qt.Recovering("QGeoPositionInfoSource::updateInterval")

	if ptr.Pointer() != nil {
		return int(C.QGeoPositionInfoSource_UpdateInterval(ptr.Pointer()))
	}
	return 0
}

func QGeoPositionInfoSource_AvailableSources() []string {
	defer qt.Recovering("QGeoPositionInfoSource::availableSources")

	return strings.Split(C.GoString(C.QGeoPositionInfoSource_QGeoPositionInfoSource_AvailableSources()), "|")
}

func QGeoPositionInfoSource_CreateDefaultSource(parent core.QObject_ITF) *QGeoPositionInfoSource {
	defer qt.Recovering("QGeoPositionInfoSource::createDefaultSource")

	return NewQGeoPositionInfoSourceFromPointer(C.QGeoPositionInfoSource_QGeoPositionInfoSource_CreateDefaultSource(core.PointerFromQObject(parent)))
}

func QGeoPositionInfoSource_CreateSource(sourceName string, parent core.QObject_ITF) *QGeoPositionInfoSource {
	defer qt.Recovering("QGeoPositionInfoSource::createSource")

	return NewQGeoPositionInfoSourceFromPointer(C.QGeoPositionInfoSource_QGeoPositionInfoSource_CreateSource(C.CString(sourceName), core.PointerFromQObject(parent)))
}

func (ptr *QGeoPositionInfoSource) ConnectError2(f func(positioningError QGeoPositionInfoSource__Error)) {
	defer qt.Recovering("connect QGeoPositionInfoSource::error")

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error2", f)
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectError2() {
	defer qt.Recovering("disconnect QGeoPositionInfoSource::error")

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error2")
	}
}

//export callbackQGeoPositionInfoSourceError2
func callbackQGeoPositionInfoSourceError2(ptr unsafe.Pointer, ptrName *C.char, positioningError C.int) {
	defer qt.Recovering("callback QGeoPositionInfoSource::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(QGeoPositionInfoSource__Error))(QGeoPositionInfoSource__Error(positioningError))
	}

}

func (ptr *QGeoPositionInfoSource) Error2(positioningError QGeoPositionInfoSource__Error) {
	defer qt.Recovering("QGeoPositionInfoSource::error")

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_Error2(ptr.Pointer(), C.int(positioningError))
	}
}

func (ptr *QGeoPositionInfoSource) Error() QGeoPositionInfoSource__Error {
	defer qt.Recovering("QGeoPositionInfoSource::error")

	if ptr.Pointer() != nil {
		return QGeoPositionInfoSource__Error(C.QGeoPositionInfoSource_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoPositionInfoSource) MinimumUpdateInterval() int {
	defer qt.Recovering("QGeoPositionInfoSource::minimumUpdateInterval")

	if ptr.Pointer() != nil {
		return int(C.QGeoPositionInfoSource_MinimumUpdateInterval(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoPositionInfoSource) PreferredPositioningMethods() QGeoPositionInfoSource__PositioningMethod {
	defer qt.Recovering("QGeoPositionInfoSource::preferredPositioningMethods")

	if ptr.Pointer() != nil {
		return QGeoPositionInfoSource__PositioningMethod(C.QGeoPositionInfoSource_PreferredPositioningMethods(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoPositionInfoSource) RequestUpdate(timeout int) {
	defer qt.Recovering("QGeoPositionInfoSource::requestUpdate")

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_RequestUpdate(ptr.Pointer(), C.int(timeout))
	}
}

func (ptr *QGeoPositionInfoSource) ConnectSetPreferredPositioningMethods(f func(methods QGeoPositionInfoSource__PositioningMethod)) {
	defer qt.Recovering("connect QGeoPositionInfoSource::setPreferredPositioningMethods")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setPreferredPositioningMethods", f)
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectSetPreferredPositioningMethods() {
	defer qt.Recovering("disconnect QGeoPositionInfoSource::setPreferredPositioningMethods")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setPreferredPositioningMethods")
	}
}

//export callbackQGeoPositionInfoSourceSetPreferredPositioningMethods
func callbackQGeoPositionInfoSourceSetPreferredPositioningMethods(ptr unsafe.Pointer, ptrName *C.char, methods C.int) {
	defer qt.Recovering("callback QGeoPositionInfoSource::setPreferredPositioningMethods")

	if signal := qt.GetSignal(C.GoString(ptrName), "setPreferredPositioningMethods"); signal != nil {
		signal.(func(QGeoPositionInfoSource__PositioningMethod))(QGeoPositionInfoSource__PositioningMethod(methods))
	} else {
		NewQGeoPositionInfoSourceFromPointer(ptr).SetPreferredPositioningMethodsDefault(QGeoPositionInfoSource__PositioningMethod(methods))
	}
}

func (ptr *QGeoPositionInfoSource) SetPreferredPositioningMethods(methods QGeoPositionInfoSource__PositioningMethod) {
	defer qt.Recovering("QGeoPositionInfoSource::setPreferredPositioningMethods")

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_SetPreferredPositioningMethods(ptr.Pointer(), C.int(methods))
	}
}

func (ptr *QGeoPositionInfoSource) SetPreferredPositioningMethodsDefault(methods QGeoPositionInfoSource__PositioningMethod) {
	defer qt.Recovering("QGeoPositionInfoSource::setPreferredPositioningMethods")

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_SetPreferredPositioningMethodsDefault(ptr.Pointer(), C.int(methods))
	}
}

func (ptr *QGeoPositionInfoSource) StartUpdates() {
	defer qt.Recovering("QGeoPositionInfoSource::startUpdates")

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_StartUpdates(ptr.Pointer())
	}
}

func (ptr *QGeoPositionInfoSource) StopUpdates() {
	defer qt.Recovering("QGeoPositionInfoSource::stopUpdates")

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_StopUpdates(ptr.Pointer())
	}
}

func (ptr *QGeoPositionInfoSource) SupportedPositioningMethods() QGeoPositionInfoSource__PositioningMethod {
	defer qt.Recovering("QGeoPositionInfoSource::supportedPositioningMethods")

	if ptr.Pointer() != nil {
		return QGeoPositionInfoSource__PositioningMethod(C.QGeoPositionInfoSource_SupportedPositioningMethods(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoPositionInfoSource) ConnectUpdateTimeout(f func()) {
	defer qt.Recovering("connect QGeoPositionInfoSource::updateTimeout")

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_ConnectUpdateTimeout(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "updateTimeout", f)
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectUpdateTimeout() {
	defer qt.Recovering("disconnect QGeoPositionInfoSource::updateTimeout")

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_DisconnectUpdateTimeout(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "updateTimeout")
	}
}

//export callbackQGeoPositionInfoSourceUpdateTimeout
func callbackQGeoPositionInfoSourceUpdateTimeout(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QGeoPositionInfoSource::updateTimeout")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateTimeout"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QGeoPositionInfoSource) UpdateTimeout() {
	defer qt.Recovering("QGeoPositionInfoSource::updateTimeout")

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_UpdateTimeout(ptr.Pointer())
	}
}

func (ptr *QGeoPositionInfoSource) DestroyQGeoPositionInfoSource() {
	defer qt.Recovering("QGeoPositionInfoSource::~QGeoPositionInfoSource")

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_DestroyQGeoPositionInfoSource(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGeoPositionInfoSource) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QGeoPositionInfoSource::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QGeoPositionInfoSource::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQGeoPositionInfoSourceTimerEvent
func callbackQGeoPositionInfoSourceTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGeoPositionInfoSource::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGeoPositionInfoSourceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGeoPositionInfoSource) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGeoPositionInfoSource::timerEvent")

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGeoPositionInfoSource) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGeoPositionInfoSource::timerEvent")

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGeoPositionInfoSource) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QGeoPositionInfoSource::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QGeoPositionInfoSource::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQGeoPositionInfoSourceChildEvent
func callbackQGeoPositionInfoSourceChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGeoPositionInfoSource::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQGeoPositionInfoSourceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QGeoPositionInfoSource) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QGeoPositionInfoSource::childEvent")

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGeoPositionInfoSource) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QGeoPositionInfoSource::childEvent")

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGeoPositionInfoSource) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGeoPositionInfoSource::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QGeoPositionInfoSource::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQGeoPositionInfoSourceCustomEvent
func callbackQGeoPositionInfoSourceCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGeoPositionInfoSource::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGeoPositionInfoSourceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGeoPositionInfoSource) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QGeoPositionInfoSource::customEvent")

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGeoPositionInfoSource) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QGeoPositionInfoSource::customEvent")

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QGeoPositionInfoSourceFactory struct {
	ptr unsafe.Pointer
}

type QGeoPositionInfoSourceFactory_ITF interface {
	QGeoPositionInfoSourceFactory_PTR() *QGeoPositionInfoSourceFactory
}

func (p *QGeoPositionInfoSourceFactory) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QGeoPositionInfoSourceFactory) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQGeoPositionInfoSourceFactory(ptr QGeoPositionInfoSourceFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoPositionInfoSourceFactory_PTR().Pointer()
	}
	return nil
}

func NewQGeoPositionInfoSourceFactoryFromPointer(ptr unsafe.Pointer) *QGeoPositionInfoSourceFactory {
	var n = new(QGeoPositionInfoSourceFactory)
	n.SetPointer(ptr)
	return n
}

func newQGeoPositionInfoSourceFactoryFromPointer(ptr unsafe.Pointer) *QGeoPositionInfoSourceFactory {
	var n = NewQGeoPositionInfoSourceFactoryFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QGeoPositionInfoSourceFactory_") {
		n.SetObjectNameAbs("QGeoPositionInfoSourceFactory_" + qt.Identifier())
	}
	return n
}

func (ptr *QGeoPositionInfoSourceFactory) QGeoPositionInfoSourceFactory_PTR() *QGeoPositionInfoSourceFactory {
	return ptr
}

func (ptr *QGeoPositionInfoSourceFactory) AreaMonitor(parent core.QObject_ITF) *QGeoAreaMonitorSource {
	defer qt.Recovering("QGeoPositionInfoSourceFactory::areaMonitor")

	if ptr.Pointer() != nil {
		return NewQGeoAreaMonitorSourceFromPointer(C.QGeoPositionInfoSourceFactory_AreaMonitor(ptr.Pointer(), core.PointerFromQObject(parent)))
	}
	return nil
}

func (ptr *QGeoPositionInfoSourceFactory) PositionInfoSource(parent core.QObject_ITF) *QGeoPositionInfoSource {
	defer qt.Recovering("QGeoPositionInfoSourceFactory::positionInfoSource")

	if ptr.Pointer() != nil {
		return NewQGeoPositionInfoSourceFromPointer(C.QGeoPositionInfoSourceFactory_PositionInfoSource(ptr.Pointer(), core.PointerFromQObject(parent)))
	}
	return nil
}

func (ptr *QGeoPositionInfoSourceFactory) SatelliteInfoSource(parent core.QObject_ITF) *QGeoSatelliteInfoSource {
	defer qt.Recovering("QGeoPositionInfoSourceFactory::satelliteInfoSource")

	if ptr.Pointer() != nil {
		return NewQGeoSatelliteInfoSourceFromPointer(C.QGeoPositionInfoSourceFactory_SatelliteInfoSource(ptr.Pointer(), core.PointerFromQObject(parent)))
	}
	return nil
}

func (ptr *QGeoPositionInfoSourceFactory) DestroyQGeoPositionInfoSourceFactory() {
	defer qt.Recovering("QGeoPositionInfoSourceFactory::~QGeoPositionInfoSourceFactory")

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSourceFactory_DestroyQGeoPositionInfoSourceFactory(ptr.Pointer())
	}
}

func (ptr *QGeoPositionInfoSourceFactory) ObjectNameAbs() string {
	defer qt.Recovering("QGeoPositionInfoSourceFactory::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoPositionInfoSourceFactory_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoPositionInfoSourceFactory) SetObjectNameAbs(name string) {
	defer qt.Recovering("QGeoPositionInfoSourceFactory::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSourceFactory_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QGeoRectangle struct {
	QGeoShape
}

type QGeoRectangle_ITF interface {
	QGeoShape_ITF
	QGeoRectangle_PTR() *QGeoRectangle
}

func PointerFromQGeoRectangle(ptr QGeoRectangle_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoRectangle_PTR().Pointer()
	}
	return nil
}

func NewQGeoRectangleFromPointer(ptr unsafe.Pointer) *QGeoRectangle {
	var n = new(QGeoRectangle)
	n.SetPointer(ptr)
	return n
}

func newQGeoRectangleFromPointer(ptr unsafe.Pointer) *QGeoRectangle {
	var n = NewQGeoRectangleFromPointer(ptr)
	return n
}

func (ptr *QGeoRectangle) QGeoRectangle_PTR() *QGeoRectangle {
	return ptr
}

func NewQGeoRectangle() *QGeoRectangle {
	defer qt.Recovering("QGeoRectangle::QGeoRectangle")

	return newQGeoRectangleFromPointer(C.QGeoRectangle_NewQGeoRectangle())
}

func NewQGeoRectangle3(topLeft QGeoCoordinate_ITF, bottomRight QGeoCoordinate_ITF) *QGeoRectangle {
	defer qt.Recovering("QGeoRectangle::QGeoRectangle")

	return newQGeoRectangleFromPointer(C.QGeoRectangle_NewQGeoRectangle3(PointerFromQGeoCoordinate(topLeft), PointerFromQGeoCoordinate(bottomRight)))
}

func NewQGeoRectangle5(other QGeoRectangle_ITF) *QGeoRectangle {
	defer qt.Recovering("QGeoRectangle::QGeoRectangle")

	return newQGeoRectangleFromPointer(C.QGeoRectangle_NewQGeoRectangle5(PointerFromQGeoRectangle(other)))
}

func NewQGeoRectangle6(other QGeoShape_ITF) *QGeoRectangle {
	defer qt.Recovering("QGeoRectangle::QGeoRectangle")

	return newQGeoRectangleFromPointer(C.QGeoRectangle_NewQGeoRectangle6(PointerFromQGeoShape(other)))
}

func (ptr *QGeoRectangle) Contains(rectangle QGeoRectangle_ITF) bool {
	defer qt.Recovering("QGeoRectangle::contains")

	if ptr.Pointer() != nil {
		return C.QGeoRectangle_Contains(ptr.Pointer(), PointerFromQGeoRectangle(rectangle)) != 0
	}
	return false
}

func (ptr *QGeoRectangle) Intersects(rectangle QGeoRectangle_ITF) bool {
	defer qt.Recovering("QGeoRectangle::intersects")

	if ptr.Pointer() != nil {
		return C.QGeoRectangle_Intersects(ptr.Pointer(), PointerFromQGeoRectangle(rectangle)) != 0
	}
	return false
}

func (ptr *QGeoRectangle) SetBottomLeft(bottomLeft QGeoCoordinate_ITF) {
	defer qt.Recovering("QGeoRectangle::setBottomLeft")

	if ptr.Pointer() != nil {
		C.QGeoRectangle_SetBottomLeft(ptr.Pointer(), PointerFromQGeoCoordinate(bottomLeft))
	}
}

func (ptr *QGeoRectangle) SetBottomRight(bottomRight QGeoCoordinate_ITF) {
	defer qt.Recovering("QGeoRectangle::setBottomRight")

	if ptr.Pointer() != nil {
		C.QGeoRectangle_SetBottomRight(ptr.Pointer(), PointerFromQGeoCoordinate(bottomRight))
	}
}

func (ptr *QGeoRectangle) SetCenter(center QGeoCoordinate_ITF) {
	defer qt.Recovering("QGeoRectangle::setCenter")

	if ptr.Pointer() != nil {
		C.QGeoRectangle_SetCenter(ptr.Pointer(), PointerFromQGeoCoordinate(center))
	}
}

func (ptr *QGeoRectangle) SetTopLeft(topLeft QGeoCoordinate_ITF) {
	defer qt.Recovering("QGeoRectangle::setTopLeft")

	if ptr.Pointer() != nil {
		C.QGeoRectangle_SetTopLeft(ptr.Pointer(), PointerFromQGeoCoordinate(topLeft))
	}
}

func (ptr *QGeoRectangle) SetTopRight(topRight QGeoCoordinate_ITF) {
	defer qt.Recovering("QGeoRectangle::setTopRight")

	if ptr.Pointer() != nil {
		C.QGeoRectangle_SetTopRight(ptr.Pointer(), PointerFromQGeoCoordinate(topRight))
	}
}

func (ptr *QGeoRectangle) ToString() string {
	defer qt.Recovering("QGeoRectangle::toString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoRectangle_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoRectangle) DestroyQGeoRectangle() {
	defer qt.Recovering("QGeoRectangle::~QGeoRectangle")

	if ptr.Pointer() != nil {
		C.QGeoRectangle_DestroyQGeoRectangle(ptr.Pointer())
	}
}

type QGeoSatelliteInfo struct {
	ptr unsafe.Pointer
}

type QGeoSatelliteInfo_ITF interface {
	QGeoSatelliteInfo_PTR() *QGeoSatelliteInfo
}

func (p *QGeoSatelliteInfo) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QGeoSatelliteInfo) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQGeoSatelliteInfo(ptr QGeoSatelliteInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoSatelliteInfo_PTR().Pointer()
	}
	return nil
}

func NewQGeoSatelliteInfoFromPointer(ptr unsafe.Pointer) *QGeoSatelliteInfo {
	var n = new(QGeoSatelliteInfo)
	n.SetPointer(ptr)
	return n
}

func newQGeoSatelliteInfoFromPointer(ptr unsafe.Pointer) *QGeoSatelliteInfo {
	var n = NewQGeoSatelliteInfoFromPointer(ptr)
	return n
}

func (ptr *QGeoSatelliteInfo) QGeoSatelliteInfo_PTR() *QGeoSatelliteInfo {
	return ptr
}

//QGeoSatelliteInfo::Attribute
type QGeoSatelliteInfo__Attribute int64

const (
	QGeoSatelliteInfo__Elevation = QGeoSatelliteInfo__Attribute(0)
	QGeoSatelliteInfo__Azimuth   = QGeoSatelliteInfo__Attribute(1)
)

//QGeoSatelliteInfo::SatelliteSystem
type QGeoSatelliteInfo__SatelliteSystem int64

const (
	QGeoSatelliteInfo__Undefined = QGeoSatelliteInfo__SatelliteSystem(0x00)
	QGeoSatelliteInfo__GPS       = QGeoSatelliteInfo__SatelliteSystem(0x01)
	QGeoSatelliteInfo__GLONASS   = QGeoSatelliteInfo__SatelliteSystem(0x02)
)

func NewQGeoSatelliteInfo() *QGeoSatelliteInfo {
	defer qt.Recovering("QGeoSatelliteInfo::QGeoSatelliteInfo")

	return newQGeoSatelliteInfoFromPointer(C.QGeoSatelliteInfo_NewQGeoSatelliteInfo())
}

func NewQGeoSatelliteInfo2(other QGeoSatelliteInfo_ITF) *QGeoSatelliteInfo {
	defer qt.Recovering("QGeoSatelliteInfo::QGeoSatelliteInfo")

	return newQGeoSatelliteInfoFromPointer(C.QGeoSatelliteInfo_NewQGeoSatelliteInfo2(PointerFromQGeoSatelliteInfo(other)))
}

func (ptr *QGeoSatelliteInfo) Attribute(attribute QGeoSatelliteInfo__Attribute) float64 {
	defer qt.Recovering("QGeoSatelliteInfo::attribute")

	if ptr.Pointer() != nil {
		return float64(C.QGeoSatelliteInfo_Attribute(ptr.Pointer(), C.int(attribute)))
	}
	return 0
}

func (ptr *QGeoSatelliteInfo) HasAttribute(attribute QGeoSatelliteInfo__Attribute) bool {
	defer qt.Recovering("QGeoSatelliteInfo::hasAttribute")

	if ptr.Pointer() != nil {
		return C.QGeoSatelliteInfo_HasAttribute(ptr.Pointer(), C.int(attribute)) != 0
	}
	return false
}

func (ptr *QGeoSatelliteInfo) RemoveAttribute(attribute QGeoSatelliteInfo__Attribute) {
	defer qt.Recovering("QGeoSatelliteInfo::removeAttribute")

	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfo_RemoveAttribute(ptr.Pointer(), C.int(attribute))
	}
}

func (ptr *QGeoSatelliteInfo) SatelliteIdentifier() int {
	defer qt.Recovering("QGeoSatelliteInfo::satelliteIdentifier")

	if ptr.Pointer() != nil {
		return int(C.QGeoSatelliteInfo_SatelliteIdentifier(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoSatelliteInfo) SatelliteSystem() QGeoSatelliteInfo__SatelliteSystem {
	defer qt.Recovering("QGeoSatelliteInfo::satelliteSystem")

	if ptr.Pointer() != nil {
		return QGeoSatelliteInfo__SatelliteSystem(C.QGeoSatelliteInfo_SatelliteSystem(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoSatelliteInfo) SetAttribute(attribute QGeoSatelliteInfo__Attribute, value float64) {
	defer qt.Recovering("QGeoSatelliteInfo::setAttribute")

	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfo_SetAttribute(ptr.Pointer(), C.int(attribute), C.double(value))
	}
}

func (ptr *QGeoSatelliteInfo) SetSatelliteIdentifier(satId int) {
	defer qt.Recovering("QGeoSatelliteInfo::setSatelliteIdentifier")

	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfo_SetSatelliteIdentifier(ptr.Pointer(), C.int(satId))
	}
}

func (ptr *QGeoSatelliteInfo) SetSatelliteSystem(system QGeoSatelliteInfo__SatelliteSystem) {
	defer qt.Recovering("QGeoSatelliteInfo::setSatelliteSystem")

	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfo_SetSatelliteSystem(ptr.Pointer(), C.int(system))
	}
}

func (ptr *QGeoSatelliteInfo) SetSignalStrength(signalStrength int) {
	defer qt.Recovering("QGeoSatelliteInfo::setSignalStrength")

	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfo_SetSignalStrength(ptr.Pointer(), C.int(signalStrength))
	}
}

func (ptr *QGeoSatelliteInfo) SignalStrength() int {
	defer qt.Recovering("QGeoSatelliteInfo::signalStrength")

	if ptr.Pointer() != nil {
		return int(C.QGeoSatelliteInfo_SignalStrength(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoSatelliteInfo) DestroyQGeoSatelliteInfo() {
	defer qt.Recovering("QGeoSatelliteInfo::~QGeoSatelliteInfo")

	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfo_DestroyQGeoSatelliteInfo(ptr.Pointer())
	}
}

type QGeoSatelliteInfoSource struct {
	core.QObject
}

type QGeoSatelliteInfoSource_ITF interface {
	core.QObject_ITF
	QGeoSatelliteInfoSource_PTR() *QGeoSatelliteInfoSource
}

func PointerFromQGeoSatelliteInfoSource(ptr QGeoSatelliteInfoSource_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoSatelliteInfoSource_PTR().Pointer()
	}
	return nil
}

func NewQGeoSatelliteInfoSourceFromPointer(ptr unsafe.Pointer) *QGeoSatelliteInfoSource {
	var n = new(QGeoSatelliteInfoSource)
	n.SetPointer(ptr)
	return n
}

func newQGeoSatelliteInfoSourceFromPointer(ptr unsafe.Pointer) *QGeoSatelliteInfoSource {
	var n = NewQGeoSatelliteInfoSourceFromPointer(ptr)
	for len(n.ObjectName()) < len("QGeoSatelliteInfoSource_") {
		n.SetObjectName("QGeoSatelliteInfoSource_" + qt.Identifier())
	}
	return n
}

func (ptr *QGeoSatelliteInfoSource) QGeoSatelliteInfoSource_PTR() *QGeoSatelliteInfoSource {
	return ptr
}

//QGeoSatelliteInfoSource::Error
type QGeoSatelliteInfoSource__Error int64

const (
	QGeoSatelliteInfoSource__AccessError        = QGeoSatelliteInfoSource__Error(0)
	QGeoSatelliteInfoSource__ClosedError        = QGeoSatelliteInfoSource__Error(1)
	QGeoSatelliteInfoSource__NoError            = QGeoSatelliteInfoSource__Error(2)
	QGeoSatelliteInfoSource__UnknownSourceError = QGeoSatelliteInfoSource__Error(-1)
)

func (ptr *QGeoSatelliteInfoSource) ConnectSetUpdateInterval(f func(msec int)) {
	defer qt.Recovering("connect QGeoSatelliteInfoSource::setUpdateInterval")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setUpdateInterval", f)
	}
}

func (ptr *QGeoSatelliteInfoSource) DisconnectSetUpdateInterval() {
	defer qt.Recovering("disconnect QGeoSatelliteInfoSource::setUpdateInterval")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setUpdateInterval")
	}
}

//export callbackQGeoSatelliteInfoSourceSetUpdateInterval
func callbackQGeoSatelliteInfoSourceSetUpdateInterval(ptr unsafe.Pointer, ptrName *C.char, msec C.int) {
	defer qt.Recovering("callback QGeoSatelliteInfoSource::setUpdateInterval")

	if signal := qt.GetSignal(C.GoString(ptrName), "setUpdateInterval"); signal != nil {
		signal.(func(int))(int(msec))
	} else {
		NewQGeoSatelliteInfoSourceFromPointer(ptr).SetUpdateIntervalDefault(int(msec))
	}
}

func (ptr *QGeoSatelliteInfoSource) SetUpdateInterval(msec int) {
	defer qt.Recovering("QGeoSatelliteInfoSource::setUpdateInterval")

	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_SetUpdateInterval(ptr.Pointer(), C.int(msec))
	}
}

func (ptr *QGeoSatelliteInfoSource) SetUpdateIntervalDefault(msec int) {
	defer qt.Recovering("QGeoSatelliteInfoSource::setUpdateInterval")

	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_SetUpdateIntervalDefault(ptr.Pointer(), C.int(msec))
	}
}

func (ptr *QGeoSatelliteInfoSource) UpdateInterval() int {
	defer qt.Recovering("QGeoSatelliteInfoSource::updateInterval")

	if ptr.Pointer() != nil {
		return int(C.QGeoSatelliteInfoSource_UpdateInterval(ptr.Pointer()))
	}
	return 0
}

func QGeoSatelliteInfoSource_AvailableSources() []string {
	defer qt.Recovering("QGeoSatelliteInfoSource::availableSources")

	return strings.Split(C.GoString(C.QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_AvailableSources()), "|")
}

func QGeoSatelliteInfoSource_CreateDefaultSource(parent core.QObject_ITF) *QGeoSatelliteInfoSource {
	defer qt.Recovering("QGeoSatelliteInfoSource::createDefaultSource")

	return NewQGeoSatelliteInfoSourceFromPointer(C.QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_CreateDefaultSource(core.PointerFromQObject(parent)))
}

func QGeoSatelliteInfoSource_CreateSource(sourceName string, parent core.QObject_ITF) *QGeoSatelliteInfoSource {
	defer qt.Recovering("QGeoSatelliteInfoSource::createSource")

	return NewQGeoSatelliteInfoSourceFromPointer(C.QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_CreateSource(C.CString(sourceName), core.PointerFromQObject(parent)))
}

func (ptr *QGeoSatelliteInfoSource) ConnectError2(f func(satelliteError QGeoSatelliteInfoSource__Error)) {
	defer qt.Recovering("connect QGeoSatelliteInfoSource::error")

	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error2", f)
	}
}

func (ptr *QGeoSatelliteInfoSource) DisconnectError2() {
	defer qt.Recovering("disconnect QGeoSatelliteInfoSource::error")

	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error2")
	}
}

//export callbackQGeoSatelliteInfoSourceError2
func callbackQGeoSatelliteInfoSourceError2(ptr unsafe.Pointer, ptrName *C.char, satelliteError C.int) {
	defer qt.Recovering("callback QGeoSatelliteInfoSource::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(QGeoSatelliteInfoSource__Error))(QGeoSatelliteInfoSource__Error(satelliteError))
	}

}

func (ptr *QGeoSatelliteInfoSource) Error2(satelliteError QGeoSatelliteInfoSource__Error) {
	defer qt.Recovering("QGeoSatelliteInfoSource::error")

	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_Error2(ptr.Pointer(), C.int(satelliteError))
	}
}

func (ptr *QGeoSatelliteInfoSource) Error() QGeoSatelliteInfoSource__Error {
	defer qt.Recovering("QGeoSatelliteInfoSource::error")

	if ptr.Pointer() != nil {
		return QGeoSatelliteInfoSource__Error(C.QGeoSatelliteInfoSource_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoSatelliteInfoSource) MinimumUpdateInterval() int {
	defer qt.Recovering("QGeoSatelliteInfoSource::minimumUpdateInterval")

	if ptr.Pointer() != nil {
		return int(C.QGeoSatelliteInfoSource_MinimumUpdateInterval(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoSatelliteInfoSource) ConnectRequestTimeout(f func()) {
	defer qt.Recovering("connect QGeoSatelliteInfoSource::requestTimeout")

	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_ConnectRequestTimeout(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "requestTimeout", f)
	}
}

func (ptr *QGeoSatelliteInfoSource) DisconnectRequestTimeout() {
	defer qt.Recovering("disconnect QGeoSatelliteInfoSource::requestTimeout")

	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_DisconnectRequestTimeout(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "requestTimeout")
	}
}

//export callbackQGeoSatelliteInfoSourceRequestTimeout
func callbackQGeoSatelliteInfoSourceRequestTimeout(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QGeoSatelliteInfoSource::requestTimeout")

	if signal := qt.GetSignal(C.GoString(ptrName), "requestTimeout"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QGeoSatelliteInfoSource) RequestTimeout() {
	defer qt.Recovering("QGeoSatelliteInfoSource::requestTimeout")

	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_RequestTimeout(ptr.Pointer())
	}
}

func (ptr *QGeoSatelliteInfoSource) RequestUpdate(timeout int) {
	defer qt.Recovering("QGeoSatelliteInfoSource::requestUpdate")

	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_RequestUpdate(ptr.Pointer(), C.int(timeout))
	}
}

func (ptr *QGeoSatelliteInfoSource) SourceName() string {
	defer qt.Recovering("QGeoSatelliteInfoSource::sourceName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoSatelliteInfoSource_SourceName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoSatelliteInfoSource) StartUpdates() {
	defer qt.Recovering("QGeoSatelliteInfoSource::startUpdates")

	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_StartUpdates(ptr.Pointer())
	}
}

func (ptr *QGeoSatelliteInfoSource) StopUpdates() {
	defer qt.Recovering("QGeoSatelliteInfoSource::stopUpdates")

	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_StopUpdates(ptr.Pointer())
	}
}

func (ptr *QGeoSatelliteInfoSource) DestroyQGeoSatelliteInfoSource() {
	defer qt.Recovering("QGeoSatelliteInfoSource::~QGeoSatelliteInfoSource")

	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_DestroyQGeoSatelliteInfoSource(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGeoSatelliteInfoSource) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QGeoSatelliteInfoSource::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QGeoSatelliteInfoSource) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QGeoSatelliteInfoSource::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQGeoSatelliteInfoSourceTimerEvent
func callbackQGeoSatelliteInfoSourceTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGeoSatelliteInfoSource::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGeoSatelliteInfoSourceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGeoSatelliteInfoSource) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGeoSatelliteInfoSource::timerEvent")

	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGeoSatelliteInfoSource) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGeoSatelliteInfoSource::timerEvent")

	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGeoSatelliteInfoSource) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QGeoSatelliteInfoSource::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QGeoSatelliteInfoSource) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QGeoSatelliteInfoSource::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQGeoSatelliteInfoSourceChildEvent
func callbackQGeoSatelliteInfoSourceChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGeoSatelliteInfoSource::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQGeoSatelliteInfoSourceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QGeoSatelliteInfoSource) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QGeoSatelliteInfoSource::childEvent")

	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGeoSatelliteInfoSource) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QGeoSatelliteInfoSource::childEvent")

	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGeoSatelliteInfoSource) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGeoSatelliteInfoSource::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QGeoSatelliteInfoSource) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QGeoSatelliteInfoSource::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQGeoSatelliteInfoSourceCustomEvent
func callbackQGeoSatelliteInfoSourceCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGeoSatelliteInfoSource::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGeoSatelliteInfoSourceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGeoSatelliteInfoSource) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QGeoSatelliteInfoSource::customEvent")

	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGeoSatelliteInfoSource) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QGeoSatelliteInfoSource::customEvent")

	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QGeoShape struct {
	ptr unsafe.Pointer
}

type QGeoShape_ITF interface {
	QGeoShape_PTR() *QGeoShape
}

func (p *QGeoShape) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QGeoShape) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQGeoShape(ptr QGeoShape_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoShape_PTR().Pointer()
	}
	return nil
}

func NewQGeoShapeFromPointer(ptr unsafe.Pointer) *QGeoShape {
	var n = new(QGeoShape)
	n.SetPointer(ptr)
	return n
}

func newQGeoShapeFromPointer(ptr unsafe.Pointer) *QGeoShape {
	var n = NewQGeoShapeFromPointer(ptr)
	return n
}

func (ptr *QGeoShape) QGeoShape_PTR() *QGeoShape {
	return ptr
}

//QGeoShape::ShapeType
type QGeoShape__ShapeType int64

const (
	QGeoShape__UnknownType   = QGeoShape__ShapeType(0)
	QGeoShape__RectangleType = QGeoShape__ShapeType(1)
	QGeoShape__CircleType    = QGeoShape__ShapeType(2)
)

func NewQGeoShape() *QGeoShape {
	defer qt.Recovering("QGeoShape::QGeoShape")

	return newQGeoShapeFromPointer(C.QGeoShape_NewQGeoShape())
}

func NewQGeoShape2(other QGeoShape_ITF) *QGeoShape {
	defer qt.Recovering("QGeoShape::QGeoShape")

	return newQGeoShapeFromPointer(C.QGeoShape_NewQGeoShape2(PointerFromQGeoShape(other)))
}

func (ptr *QGeoShape) Contains(coordinate QGeoCoordinate_ITF) bool {
	defer qt.Recovering("QGeoShape::contains")

	if ptr.Pointer() != nil {
		return C.QGeoShape_Contains(ptr.Pointer(), PointerFromQGeoCoordinate(coordinate)) != 0
	}
	return false
}

func (ptr *QGeoShape) ExtendShape(coordinate QGeoCoordinate_ITF) {
	defer qt.Recovering("QGeoShape::extendShape")

	if ptr.Pointer() != nil {
		C.QGeoShape_ExtendShape(ptr.Pointer(), PointerFromQGeoCoordinate(coordinate))
	}
}

func (ptr *QGeoShape) IsEmpty() bool {
	defer qt.Recovering("QGeoShape::isEmpty")

	if ptr.Pointer() != nil {
		return C.QGeoShape_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGeoShape) IsValid() bool {
	defer qt.Recovering("QGeoShape::isValid")

	if ptr.Pointer() != nil {
		return C.QGeoShape_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGeoShape) ToString() string {
	defer qt.Recovering("QGeoShape::toString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoShape_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoShape) Type() QGeoShape__ShapeType {
	defer qt.Recovering("QGeoShape::type")

	if ptr.Pointer() != nil {
		return QGeoShape__ShapeType(C.QGeoShape_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoShape) DestroyQGeoShape() {
	defer qt.Recovering("QGeoShape::~QGeoShape")

	if ptr.Pointer() != nil {
		C.QGeoShape_DestroyQGeoShape(ptr.Pointer())
	}
}

type QNmeaPositionInfoSource struct {
	QGeoPositionInfoSource
}

type QNmeaPositionInfoSource_ITF interface {
	QGeoPositionInfoSource_ITF
	QNmeaPositionInfoSource_PTR() *QNmeaPositionInfoSource
}

func PointerFromQNmeaPositionInfoSource(ptr QNmeaPositionInfoSource_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNmeaPositionInfoSource_PTR().Pointer()
	}
	return nil
}

func NewQNmeaPositionInfoSourceFromPointer(ptr unsafe.Pointer) *QNmeaPositionInfoSource {
	var n = new(QNmeaPositionInfoSource)
	n.SetPointer(ptr)
	return n
}

func newQNmeaPositionInfoSourceFromPointer(ptr unsafe.Pointer) *QNmeaPositionInfoSource {
	var n = NewQNmeaPositionInfoSourceFromPointer(ptr)
	for len(n.ObjectName()) < len("QNmeaPositionInfoSource_") {
		n.SetObjectName("QNmeaPositionInfoSource_" + qt.Identifier())
	}
	return n
}

func (ptr *QNmeaPositionInfoSource) QNmeaPositionInfoSource_PTR() *QNmeaPositionInfoSource {
	return ptr
}

//QNmeaPositionInfoSource::UpdateMode
type QNmeaPositionInfoSource__UpdateMode int64

const (
	QNmeaPositionInfoSource__RealTimeMode   = QNmeaPositionInfoSource__UpdateMode(1)
	QNmeaPositionInfoSource__SimulationMode = QNmeaPositionInfoSource__UpdateMode(2)
)

func NewQNmeaPositionInfoSource(updateMode QNmeaPositionInfoSource__UpdateMode, parent core.QObject_ITF) *QNmeaPositionInfoSource {
	defer qt.Recovering("QNmeaPositionInfoSource::QNmeaPositionInfoSource")

	return newQNmeaPositionInfoSourceFromPointer(C.QNmeaPositionInfoSource_NewQNmeaPositionInfoSource(C.int(updateMode), core.PointerFromQObject(parent)))
}

func (ptr *QNmeaPositionInfoSource) Device() *core.QIODevice {
	defer qt.Recovering("QNmeaPositionInfoSource::device")

	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QNmeaPositionInfoSource_Device(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNmeaPositionInfoSource) Error() QGeoPositionInfoSource__Error {
	defer qt.Recovering("QNmeaPositionInfoSource::error")

	if ptr.Pointer() != nil {
		return QGeoPositionInfoSource__Error(C.QNmeaPositionInfoSource_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNmeaPositionInfoSource) MinimumUpdateInterval() int {
	defer qt.Recovering("QNmeaPositionInfoSource::minimumUpdateInterval")

	if ptr.Pointer() != nil {
		return int(C.QNmeaPositionInfoSource_MinimumUpdateInterval(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNmeaPositionInfoSource) ParsePosInfoFromNmeaData(data string, size int, posInfo QGeoPositionInfo_ITF, hasFix bool) bool {
	defer qt.Recovering("QNmeaPositionInfoSource::parsePosInfoFromNmeaData")

	if ptr.Pointer() != nil {
		return C.QNmeaPositionInfoSource_ParsePosInfoFromNmeaData(ptr.Pointer(), C.CString(data), C.int(size), PointerFromQGeoPositionInfo(posInfo), C.int(qt.GoBoolToInt(hasFix))) != 0
	}
	return false
}

func (ptr *QNmeaPositionInfoSource) ConnectRequestUpdate(f func(msec int)) {
	defer qt.Recovering("connect QNmeaPositionInfoSource::requestUpdate")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "requestUpdate", f)
	}
}

func (ptr *QNmeaPositionInfoSource) DisconnectRequestUpdate() {
	defer qt.Recovering("disconnect QNmeaPositionInfoSource::requestUpdate")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "requestUpdate")
	}
}

//export callbackQNmeaPositionInfoSourceRequestUpdate
func callbackQNmeaPositionInfoSourceRequestUpdate(ptr unsafe.Pointer, ptrName *C.char, msec C.int) bool {
	defer qt.Recovering("callback QNmeaPositionInfoSource::requestUpdate")

	if signal := qt.GetSignal(C.GoString(ptrName), "requestUpdate"); signal != nil {
		signal.(func(int))(int(msec))
		return true
	}
	return false

}

func (ptr *QNmeaPositionInfoSource) RequestUpdate(msec int) {
	defer qt.Recovering("QNmeaPositionInfoSource::requestUpdate")

	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_RequestUpdate(ptr.Pointer(), C.int(msec))
	}
}

func (ptr *QNmeaPositionInfoSource) RequestUpdateDefault(msec int) {
	defer qt.Recovering("QNmeaPositionInfoSource::requestUpdate")

	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_RequestUpdateDefault(ptr.Pointer(), C.int(msec))
	}
}

func (ptr *QNmeaPositionInfoSource) SetDevice(device core.QIODevice_ITF) {
	defer qt.Recovering("QNmeaPositionInfoSource::setDevice")

	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_SetDevice(ptr.Pointer(), core.PointerFromQIODevice(device))
	}
}

func (ptr *QNmeaPositionInfoSource) ConnectSetUpdateInterval(f func(msec int)) {
	defer qt.Recovering("connect QNmeaPositionInfoSource::setUpdateInterval")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setUpdateInterval", f)
	}
}

func (ptr *QNmeaPositionInfoSource) DisconnectSetUpdateInterval() {
	defer qt.Recovering("disconnect QNmeaPositionInfoSource::setUpdateInterval")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setUpdateInterval")
	}
}

//export callbackQNmeaPositionInfoSourceSetUpdateInterval
func callbackQNmeaPositionInfoSourceSetUpdateInterval(ptr unsafe.Pointer, ptrName *C.char, msec C.int) {
	defer qt.Recovering("callback QNmeaPositionInfoSource::setUpdateInterval")

	if signal := qt.GetSignal(C.GoString(ptrName), "setUpdateInterval"); signal != nil {
		signal.(func(int))(int(msec))
	} else {
		NewQNmeaPositionInfoSourceFromPointer(ptr).SetUpdateIntervalDefault(int(msec))
	}
}

func (ptr *QNmeaPositionInfoSource) SetUpdateInterval(msec int) {
	defer qt.Recovering("QNmeaPositionInfoSource::setUpdateInterval")

	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_SetUpdateInterval(ptr.Pointer(), C.int(msec))
	}
}

func (ptr *QNmeaPositionInfoSource) SetUpdateIntervalDefault(msec int) {
	defer qt.Recovering("QNmeaPositionInfoSource::setUpdateInterval")

	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_SetUpdateIntervalDefault(ptr.Pointer(), C.int(msec))
	}
}

func (ptr *QNmeaPositionInfoSource) ConnectStartUpdates(f func()) {
	defer qt.Recovering("connect QNmeaPositionInfoSource::startUpdates")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "startUpdates", f)
	}
}

func (ptr *QNmeaPositionInfoSource) DisconnectStartUpdates() {
	defer qt.Recovering("disconnect QNmeaPositionInfoSource::startUpdates")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "startUpdates")
	}
}

//export callbackQNmeaPositionInfoSourceStartUpdates
func callbackQNmeaPositionInfoSourceStartUpdates(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QNmeaPositionInfoSource::startUpdates")

	if signal := qt.GetSignal(C.GoString(ptrName), "startUpdates"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QNmeaPositionInfoSource) StartUpdates() {
	defer qt.Recovering("QNmeaPositionInfoSource::startUpdates")

	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_StartUpdates(ptr.Pointer())
	}
}

func (ptr *QNmeaPositionInfoSource) StartUpdatesDefault() {
	defer qt.Recovering("QNmeaPositionInfoSource::startUpdates")

	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_StartUpdatesDefault(ptr.Pointer())
	}
}

func (ptr *QNmeaPositionInfoSource) ConnectStopUpdates(f func()) {
	defer qt.Recovering("connect QNmeaPositionInfoSource::stopUpdates")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "stopUpdates", f)
	}
}

func (ptr *QNmeaPositionInfoSource) DisconnectStopUpdates() {
	defer qt.Recovering("disconnect QNmeaPositionInfoSource::stopUpdates")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "stopUpdates")
	}
}

//export callbackQNmeaPositionInfoSourceStopUpdates
func callbackQNmeaPositionInfoSourceStopUpdates(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QNmeaPositionInfoSource::stopUpdates")

	if signal := qt.GetSignal(C.GoString(ptrName), "stopUpdates"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QNmeaPositionInfoSource) StopUpdates() {
	defer qt.Recovering("QNmeaPositionInfoSource::stopUpdates")

	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_StopUpdates(ptr.Pointer())
	}
}

func (ptr *QNmeaPositionInfoSource) StopUpdatesDefault() {
	defer qt.Recovering("QNmeaPositionInfoSource::stopUpdates")

	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_StopUpdatesDefault(ptr.Pointer())
	}
}

func (ptr *QNmeaPositionInfoSource) SupportedPositioningMethods() QGeoPositionInfoSource__PositioningMethod {
	defer qt.Recovering("QNmeaPositionInfoSource::supportedPositioningMethods")

	if ptr.Pointer() != nil {
		return QGeoPositionInfoSource__PositioningMethod(C.QNmeaPositionInfoSource_SupportedPositioningMethods(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNmeaPositionInfoSource) UpdateMode() QNmeaPositionInfoSource__UpdateMode {
	defer qt.Recovering("QNmeaPositionInfoSource::updateMode")

	if ptr.Pointer() != nil {
		return QNmeaPositionInfoSource__UpdateMode(C.QNmeaPositionInfoSource_UpdateMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNmeaPositionInfoSource) DestroyQNmeaPositionInfoSource() {
	defer qt.Recovering("QNmeaPositionInfoSource::~QNmeaPositionInfoSource")

	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_DestroyQNmeaPositionInfoSource(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNmeaPositionInfoSource) ConnectSetPreferredPositioningMethods(f func(methods QGeoPositionInfoSource__PositioningMethod)) {
	defer qt.Recovering("connect QNmeaPositionInfoSource::setPreferredPositioningMethods")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setPreferredPositioningMethods", f)
	}
}

func (ptr *QNmeaPositionInfoSource) DisconnectSetPreferredPositioningMethods() {
	defer qt.Recovering("disconnect QNmeaPositionInfoSource::setPreferredPositioningMethods")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setPreferredPositioningMethods")
	}
}

//export callbackQNmeaPositionInfoSourceSetPreferredPositioningMethods
func callbackQNmeaPositionInfoSourceSetPreferredPositioningMethods(ptr unsafe.Pointer, ptrName *C.char, methods C.int) {
	defer qt.Recovering("callback QNmeaPositionInfoSource::setPreferredPositioningMethods")

	if signal := qt.GetSignal(C.GoString(ptrName), "setPreferredPositioningMethods"); signal != nil {
		signal.(func(QGeoPositionInfoSource__PositioningMethod))(QGeoPositionInfoSource__PositioningMethod(methods))
	} else {
		NewQNmeaPositionInfoSourceFromPointer(ptr).SetPreferredPositioningMethodsDefault(QGeoPositionInfoSource__PositioningMethod(methods))
	}
}

func (ptr *QNmeaPositionInfoSource) SetPreferredPositioningMethods(methods QGeoPositionInfoSource__PositioningMethod) {
	defer qt.Recovering("QNmeaPositionInfoSource::setPreferredPositioningMethods")

	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_SetPreferredPositioningMethods(ptr.Pointer(), C.int(methods))
	}
}

func (ptr *QNmeaPositionInfoSource) SetPreferredPositioningMethodsDefault(methods QGeoPositionInfoSource__PositioningMethod) {
	defer qt.Recovering("QNmeaPositionInfoSource::setPreferredPositioningMethods")

	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_SetPreferredPositioningMethodsDefault(ptr.Pointer(), C.int(methods))
	}
}

func (ptr *QNmeaPositionInfoSource) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QNmeaPositionInfoSource::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QNmeaPositionInfoSource) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QNmeaPositionInfoSource::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQNmeaPositionInfoSourceTimerEvent
func callbackQNmeaPositionInfoSourceTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNmeaPositionInfoSource::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQNmeaPositionInfoSourceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QNmeaPositionInfoSource) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNmeaPositionInfoSource::timerEvent")

	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNmeaPositionInfoSource) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNmeaPositionInfoSource::timerEvent")

	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNmeaPositionInfoSource) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QNmeaPositionInfoSource::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QNmeaPositionInfoSource) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QNmeaPositionInfoSource::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQNmeaPositionInfoSourceChildEvent
func callbackQNmeaPositionInfoSourceChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNmeaPositionInfoSource::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQNmeaPositionInfoSourceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QNmeaPositionInfoSource) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNmeaPositionInfoSource::childEvent")

	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNmeaPositionInfoSource) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNmeaPositionInfoSource::childEvent")

	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNmeaPositionInfoSource) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QNmeaPositionInfoSource::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QNmeaPositionInfoSource) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QNmeaPositionInfoSource::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQNmeaPositionInfoSourceCustomEvent
func callbackQNmeaPositionInfoSourceCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNmeaPositionInfoSource::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQNmeaPositionInfoSourceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QNmeaPositionInfoSource) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QNmeaPositionInfoSource::customEvent")

	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QNmeaPositionInfoSource) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QNmeaPositionInfoSource::customEvent")

	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
