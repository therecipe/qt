// +build !minimal

package positioning

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/internal"
	"strings"
	"unsafe"
)

func unpackStringList(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}
	return strings.Split(s, "¡¦!")
}

type LocationSingleton struct {
	core.QObject
}

type LocationSingleton_ITF interface {
	core.QObject_ITF
	LocationSingleton_PTR() *LocationSingleton
}

func (ptr *LocationSingleton) LocationSingleton_PTR() *LocationSingleton {
	return ptr
}

func (ptr *LocationSingleton) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *LocationSingleton) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromLocationSingleton(ptr LocationSingleton_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.LocationSingleton_PTR().Pointer()
	}
	return nil
}

func (n *LocationSingleton) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *LocationSingleton) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewLocationSingletonFromPointer(ptr unsafe.Pointer) (n *LocationSingleton) {
	n = new(LocationSingleton)
	n.InitFromInternal(uintptr(ptr), "positioning.LocationSingleton")
	return
}

func (ptr *LocationSingleton) DestroyLocationSingleton() {
}

type QGeoAddress struct {
	internal.Internal
}

type QGeoAddress_ITF interface {
	QGeoAddress_PTR() *QGeoAddress
}

func (ptr *QGeoAddress) QGeoAddress_PTR() *QGeoAddress {
	return ptr
}

func (ptr *QGeoAddress) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QGeoAddress) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQGeoAddress(ptr QGeoAddress_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoAddress_PTR().Pointer()
	}
	return nil
}

func (n *QGeoAddress) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQGeoAddressFromPointer(ptr unsafe.Pointer) (n *QGeoAddress) {
	n = new(QGeoAddress)
	n.InitFromInternal(uintptr(ptr), "positioning.QGeoAddress")
	return
}
func NewQGeoAddress() *QGeoAddress {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.NewQGeoAddress", ""}).(*QGeoAddress)
}

func NewQGeoAddress2(other QGeoAddress_ITF) *QGeoAddress {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.NewQGeoAddress2", "", other}).(*QGeoAddress)
}

func (ptr *QGeoAddress) City() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "City"}).(string)
}

func (ptr *QGeoAddress) Clear() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Clear"})
}

func (ptr *QGeoAddress) Country() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Country"}).(string)
}

func (ptr *QGeoAddress) CountryCode() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CountryCode"}).(string)
}

func (ptr *QGeoAddress) County() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "County"}).(string)
}

func (ptr *QGeoAddress) District() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "District"}).(string)
}

func (ptr *QGeoAddress) IsEmpty() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsEmpty"}).(bool)
}

func (ptr *QGeoAddress) IsTextGenerated() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsTextGenerated"}).(bool)
}

func (ptr *QGeoAddress) PostalCode() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PostalCode"}).(string)
}

func (ptr *QGeoAddress) SetCity(city string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetCity", city})
}

func (ptr *QGeoAddress) SetCountry(country string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetCountry", country})
}

func (ptr *QGeoAddress) SetCountryCode(countryCode string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetCountryCode", countryCode})
}

func (ptr *QGeoAddress) SetCounty(county string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetCounty", county})
}

func (ptr *QGeoAddress) SetDistrict(district string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDistrict", district})
}

func (ptr *QGeoAddress) SetPostalCode(postalCode string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPostalCode", postalCode})
}

func (ptr *QGeoAddress) SetState(state string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetState", state})
}

func (ptr *QGeoAddress) SetStreet(street string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetStreet", street})
}

func (ptr *QGeoAddress) SetText(text string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetText", text})
}

func (ptr *QGeoAddress) State() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "State"}).(string)
}

func (ptr *QGeoAddress) Street() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Street"}).(string)
}

func (ptr *QGeoAddress) Text() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Text"}).(string)
}

func (ptr *QGeoAddress) DestroyQGeoAddress() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQGeoAddress"})
}

type QGeoAreaMonitorInfo struct {
	internal.Internal
}

type QGeoAreaMonitorInfo_ITF interface {
	QGeoAreaMonitorInfo_PTR() *QGeoAreaMonitorInfo
}

func (ptr *QGeoAreaMonitorInfo) QGeoAreaMonitorInfo_PTR() *QGeoAreaMonitorInfo {
	return ptr
}

func (ptr *QGeoAreaMonitorInfo) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QGeoAreaMonitorInfo) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQGeoAreaMonitorInfo(ptr QGeoAreaMonitorInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoAreaMonitorInfo_PTR().Pointer()
	}
	return nil
}

func (n *QGeoAreaMonitorInfo) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQGeoAreaMonitorInfoFromPointer(ptr unsafe.Pointer) (n *QGeoAreaMonitorInfo) {
	n = new(QGeoAreaMonitorInfo)
	n.InitFromInternal(uintptr(ptr), "positioning.QGeoAreaMonitorInfo")
	return
}
func NewQGeoAreaMonitorInfo(name string) *QGeoAreaMonitorInfo {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.NewQGeoAreaMonitorInfo", "", name}).(*QGeoAreaMonitorInfo)
}

func NewQGeoAreaMonitorInfo2(other QGeoAreaMonitorInfo_ITF) *QGeoAreaMonitorInfo {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.NewQGeoAreaMonitorInfo2", "", other}).(*QGeoAreaMonitorInfo)
}

func (ptr *QGeoAreaMonitorInfo) Area() *QGeoShape {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Area"}).(*QGeoShape)
}

func (ptr *QGeoAreaMonitorInfo) Expiration() *core.QDateTime {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Expiration"}).(*core.QDateTime)
}

func (ptr *QGeoAreaMonitorInfo) Identifier() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Identifier"}).(string)
}

func (ptr *QGeoAreaMonitorInfo) IsPersistent() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsPersistent"}).(bool)
}

func (ptr *QGeoAreaMonitorInfo) IsValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsValid"}).(bool)
}

func (ptr *QGeoAreaMonitorInfo) Name() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Name"}).(string)
}

func (ptr *QGeoAreaMonitorInfo) NotificationParameters() map[string]*core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NotificationParameters"}).(map[string]*core.QVariant)
}

func (ptr *QGeoAreaMonitorInfo) SetArea(newShape QGeoShape_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetArea", newShape})
}

func (ptr *QGeoAreaMonitorInfo) SetExpiration(expiry core.QDateTime_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetExpiration", expiry})
}

func (ptr *QGeoAreaMonitorInfo) SetName(name string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetName", name})
}

func (ptr *QGeoAreaMonitorInfo) SetNotificationParameters(parameters map[string]*core.QVariant) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetNotificationParameters", parameters})
}

func (ptr *QGeoAreaMonitorInfo) SetPersistent(isPersistent bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPersistent", isPersistent})
}

func (ptr *QGeoAreaMonitorInfo) DestroyQGeoAreaMonitorInfo() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQGeoAreaMonitorInfo"})
}

func (ptr *QGeoAreaMonitorInfo) __notificationParameters_atList(v string, i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__notificationParameters_atList", v, i}).(*core.QVariant)
}

func (ptr *QGeoAreaMonitorInfo) __notificationParameters_setList(key string, i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__notificationParameters_setList", key, i})
}

func (ptr *QGeoAreaMonitorInfo) __notificationParameters_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__notificationParameters_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoAreaMonitorInfo) __notificationParameters_keyList() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__notificationParameters_keyList"}).([]string)
}

func (ptr *QGeoAreaMonitorInfo) __setNotificationParameters_parameters_atList(v string, i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setNotificationParameters_parameters_atList", v, i}).(*core.QVariant)
}

func (ptr *QGeoAreaMonitorInfo) __setNotificationParameters_parameters_setList(key string, i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setNotificationParameters_parameters_setList", key, i})
}

func (ptr *QGeoAreaMonitorInfo) __setNotificationParameters_parameters_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setNotificationParameters_parameters_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoAreaMonitorInfo) __setNotificationParameters_parameters_keyList() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setNotificationParameters_parameters_keyList"}).([]string)
}

func (ptr *QGeoAreaMonitorInfo) ____notificationParameters_keyList_atList(i int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____notificationParameters_keyList_atList", i}).(string)
}

func (ptr *QGeoAreaMonitorInfo) ____notificationParameters_keyList_setList(i string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____notificationParameters_keyList_setList", i})
}

func (ptr *QGeoAreaMonitorInfo) ____notificationParameters_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____notificationParameters_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoAreaMonitorInfo) ____setNotificationParameters_parameters_keyList_atList(i int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setNotificationParameters_parameters_keyList_atList", i}).(string)
}

func (ptr *QGeoAreaMonitorInfo) ____setNotificationParameters_parameters_keyList_setList(i string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setNotificationParameters_parameters_keyList_setList", i})
}

func (ptr *QGeoAreaMonitorInfo) ____setNotificationParameters_parameters_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setNotificationParameters_parameters_keyList_newList"}).(unsafe.Pointer)
}

type QGeoAreaMonitorSource struct {
	core.QObject
}

type QGeoAreaMonitorSource_ITF interface {
	core.QObject_ITF
	QGeoAreaMonitorSource_PTR() *QGeoAreaMonitorSource
}

func (ptr *QGeoAreaMonitorSource) QGeoAreaMonitorSource_PTR() *QGeoAreaMonitorSource {
	return ptr
}

func (ptr *QGeoAreaMonitorSource) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QGeoAreaMonitorSource) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQGeoAreaMonitorSource(ptr QGeoAreaMonitorSource_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoAreaMonitorSource_PTR().Pointer()
	}
	return nil
}

func (n *QGeoAreaMonitorSource) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QGeoAreaMonitorSource) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQGeoAreaMonitorSourceFromPointer(ptr unsafe.Pointer) (n *QGeoAreaMonitorSource) {
	n = new(QGeoAreaMonitorSource)
	n.InitFromInternal(uintptr(ptr), "positioning.QGeoAreaMonitorSource")
	return
}

//go:generate stringer -type=QGeoAreaMonitorSource__Error
//QGeoAreaMonitorSource::Error
type QGeoAreaMonitorSource__Error int64

const (
	QGeoAreaMonitorSource__AccessError              QGeoAreaMonitorSource__Error = QGeoAreaMonitorSource__Error(0)
	QGeoAreaMonitorSource__InsufficientPositionInfo QGeoAreaMonitorSource__Error = QGeoAreaMonitorSource__Error(1)
	QGeoAreaMonitorSource__UnknownSourceError       QGeoAreaMonitorSource__Error = QGeoAreaMonitorSource__Error(2)
	QGeoAreaMonitorSource__NoError                  QGeoAreaMonitorSource__Error = QGeoAreaMonitorSource__Error(3)
)

//go:generate stringer -type=QGeoAreaMonitorSource__AreaMonitorFeature
//QGeoAreaMonitorSource::AreaMonitorFeature
type QGeoAreaMonitorSource__AreaMonitorFeature int64

const (
	QGeoAreaMonitorSource__PersistentAreaMonitorFeature QGeoAreaMonitorSource__AreaMonitorFeature = QGeoAreaMonitorSource__AreaMonitorFeature(0x00000001)
	QGeoAreaMonitorSource__AnyAreaMonitorFeature        QGeoAreaMonitorSource__AreaMonitorFeature = QGeoAreaMonitorSource__AreaMonitorFeature(0xffffffff)
)

func NewQGeoAreaMonitorSource(parent core.QObject_ITF) *QGeoAreaMonitorSource {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.NewQGeoAreaMonitorSource", "", parent}).(*QGeoAreaMonitorSource)
}

func (ptr *QGeoAreaMonitorSource) ConnectActiveMonitors(f func() []*QGeoAreaMonitorInfo) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectActiveMonitors", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoAreaMonitorSource) DisconnectActiveMonitors() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectActiveMonitors"})
}

func (ptr *QGeoAreaMonitorSource) ActiveMonitors() []*QGeoAreaMonitorInfo {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ActiveMonitors"}).([]*QGeoAreaMonitorInfo)
}

func (ptr *QGeoAreaMonitorSource) ConnectActiveMonitors2(f func(lookupArea *QGeoShape) []*QGeoAreaMonitorInfo) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectActiveMonitors2", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoAreaMonitorSource) DisconnectActiveMonitors2() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectActiveMonitors2"})
}

func (ptr *QGeoAreaMonitorSource) ActiveMonitors2(lookupArea QGeoShape_ITF) []*QGeoAreaMonitorInfo {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ActiveMonitors2", lookupArea}).([]*QGeoAreaMonitorInfo)
}

func (ptr *QGeoAreaMonitorSource) ConnectAreaEntered(f func(monitor *QGeoAreaMonitorInfo, update *QGeoPositionInfo)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAreaEntered", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoAreaMonitorSource) DisconnectAreaEntered() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAreaEntered"})
}

func (ptr *QGeoAreaMonitorSource) AreaEntered(monitor QGeoAreaMonitorInfo_ITF, update QGeoPositionInfo_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AreaEntered", monitor, update})
}

func (ptr *QGeoAreaMonitorSource) ConnectAreaExited(f func(monitor *QGeoAreaMonitorInfo, update *QGeoPositionInfo)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAreaExited", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoAreaMonitorSource) DisconnectAreaExited() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAreaExited"})
}

func (ptr *QGeoAreaMonitorSource) AreaExited(monitor QGeoAreaMonitorInfo_ITF, update QGeoPositionInfo_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AreaExited", monitor, update})
}

func QGeoAreaMonitorSource_AvailableSources() []string {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.QGeoAreaMonitorSource_AvailableSources", ""}).([]string)
}

func (ptr *QGeoAreaMonitorSource) AvailableSources() []string {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.QGeoAreaMonitorSource_AvailableSources", ""}).([]string)
}

func QGeoAreaMonitorSource_CreateDefaultSource(parent core.QObject_ITF) *QGeoAreaMonitorSource {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.QGeoAreaMonitorSource_CreateDefaultSource", "", parent}).(*QGeoAreaMonitorSource)
}

func (ptr *QGeoAreaMonitorSource) CreateDefaultSource(parent core.QObject_ITF) *QGeoAreaMonitorSource {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.QGeoAreaMonitorSource_CreateDefaultSource", "", parent}).(*QGeoAreaMonitorSource)
}

func QGeoAreaMonitorSource_CreateSource(sourceName string, parent core.QObject_ITF) *QGeoAreaMonitorSource {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.QGeoAreaMonitorSource_CreateSource", "", sourceName, parent}).(*QGeoAreaMonitorSource)
}

func (ptr *QGeoAreaMonitorSource) CreateSource(sourceName string, parent core.QObject_ITF) *QGeoAreaMonitorSource {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.QGeoAreaMonitorSource_CreateSource", "", sourceName, parent}).(*QGeoAreaMonitorSource)
}

func (ptr *QGeoAreaMonitorSource) ConnectError(f func() QGeoAreaMonitorSource__Error) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectError", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoAreaMonitorSource) DisconnectError() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectError"})
}

func (ptr *QGeoAreaMonitorSource) Error() QGeoAreaMonitorSource__Error {

	return QGeoAreaMonitorSource__Error(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error"}).(float64))
}

func (ptr *QGeoAreaMonitorSource) ConnectError2(f func(areaMonitoringError QGeoAreaMonitorSource__Error)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectError2", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoAreaMonitorSource) DisconnectError2() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectError2"})
}

func (ptr *QGeoAreaMonitorSource) Error2(areaMonitoringError QGeoAreaMonitorSource__Error) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error2", areaMonitoringError})
}

func (ptr *QGeoAreaMonitorSource) ConnectMonitorExpired(f func(monitor *QGeoAreaMonitorInfo)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMonitorExpired", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoAreaMonitorSource) DisconnectMonitorExpired() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMonitorExpired"})
}

func (ptr *QGeoAreaMonitorSource) MonitorExpired(monitor QGeoAreaMonitorInfo_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MonitorExpired", monitor})
}

func (ptr *QGeoAreaMonitorSource) ConnectPositionInfoSource(f func() *QGeoPositionInfoSource) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPositionInfoSource", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoAreaMonitorSource) DisconnectPositionInfoSource() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPositionInfoSource"})
}

func (ptr *QGeoAreaMonitorSource) PositionInfoSource() *QGeoPositionInfoSource {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PositionInfoSource"}).(*QGeoPositionInfoSource)
}

func (ptr *QGeoAreaMonitorSource) PositionInfoSourceDefault() *QGeoPositionInfoSource {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PositionInfoSourceDefault"}).(*QGeoPositionInfoSource)
}

func (ptr *QGeoAreaMonitorSource) ConnectRequestUpdate(f func(monitor *QGeoAreaMonitorInfo, sign string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRequestUpdate", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoAreaMonitorSource) DisconnectRequestUpdate() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRequestUpdate"})
}

func (ptr *QGeoAreaMonitorSource) RequestUpdate(monitor QGeoAreaMonitorInfo_ITF, sign string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RequestUpdate", monitor, sign}).(bool)
}

func (ptr *QGeoAreaMonitorSource) ConnectSetPositionInfoSource(f func(newSource *QGeoPositionInfoSource)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetPositionInfoSource", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoAreaMonitorSource) DisconnectSetPositionInfoSource() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetPositionInfoSource"})
}

func (ptr *QGeoAreaMonitorSource) SetPositionInfoSource(newSource QGeoPositionInfoSource_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPositionInfoSource", newSource})
}

func (ptr *QGeoAreaMonitorSource) SetPositionInfoSourceDefault(newSource QGeoPositionInfoSource_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPositionInfoSourceDefault", newSource})
}

func (ptr *QGeoAreaMonitorSource) SourceName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SourceName"}).(string)
}

func (ptr *QGeoAreaMonitorSource) ConnectStartMonitoring(f func(monitor *QGeoAreaMonitorInfo) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStartMonitoring", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoAreaMonitorSource) DisconnectStartMonitoring() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStartMonitoring"})
}

func (ptr *QGeoAreaMonitorSource) StartMonitoring(monitor QGeoAreaMonitorInfo_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StartMonitoring", monitor}).(bool)
}

func (ptr *QGeoAreaMonitorSource) ConnectStopMonitoring(f func(monitor *QGeoAreaMonitorInfo) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStopMonitoring", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoAreaMonitorSource) DisconnectStopMonitoring() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStopMonitoring"})
}

func (ptr *QGeoAreaMonitorSource) StopMonitoring(monitor QGeoAreaMonitorInfo_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StopMonitoring", monitor}).(bool)
}

func (ptr *QGeoAreaMonitorSource) ConnectSupportedAreaMonitorFeatures(f func() QGeoAreaMonitorSource__AreaMonitorFeature) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSupportedAreaMonitorFeatures", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoAreaMonitorSource) DisconnectSupportedAreaMonitorFeatures() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSupportedAreaMonitorFeatures"})
}

func (ptr *QGeoAreaMonitorSource) SupportedAreaMonitorFeatures() QGeoAreaMonitorSource__AreaMonitorFeature {

	return QGeoAreaMonitorSource__AreaMonitorFeature(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SupportedAreaMonitorFeatures"}).(float64))
}

func (ptr *QGeoAreaMonitorSource) ConnectDestroyQGeoAreaMonitorSource(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQGeoAreaMonitorSource", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoAreaMonitorSource) DisconnectDestroyQGeoAreaMonitorSource() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQGeoAreaMonitorSource"})
}

func (ptr *QGeoAreaMonitorSource) DestroyQGeoAreaMonitorSource() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQGeoAreaMonitorSource"})
}

func (ptr *QGeoAreaMonitorSource) DestroyQGeoAreaMonitorSourceDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQGeoAreaMonitorSourceDefault"})
}

func (ptr *QGeoAreaMonitorSource) __activeMonitors_atList(i int) *QGeoAreaMonitorInfo {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__activeMonitors_atList", i}).(*QGeoAreaMonitorInfo)
}

func (ptr *QGeoAreaMonitorSource) __activeMonitors_setList(i QGeoAreaMonitorInfo_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__activeMonitors_setList", i})
}

func (ptr *QGeoAreaMonitorSource) __activeMonitors_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__activeMonitors_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoAreaMonitorSource) __activeMonitors_atList2(i int) *QGeoAreaMonitorInfo {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__activeMonitors_atList2", i}).(*QGeoAreaMonitorInfo)
}

func (ptr *QGeoAreaMonitorSource) __activeMonitors_setList2(i QGeoAreaMonitorInfo_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__activeMonitors_setList2", i})
}

func (ptr *QGeoAreaMonitorSource) __activeMonitors_newList2() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__activeMonitors_newList2"}).(unsafe.Pointer)
}

func (ptr *QGeoAreaMonitorSource) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QGeoAreaMonitorSource) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QGeoAreaMonitorSource) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoAreaMonitorSource) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QGeoAreaMonitorSource) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QGeoAreaMonitorSource) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoAreaMonitorSource) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QGeoAreaMonitorSource) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QGeoAreaMonitorSource) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoAreaMonitorSource) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QGeoAreaMonitorSource) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QGeoAreaMonitorSource) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QGeoAreaMonitorSource) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QGeoAreaMonitorSource) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QGeoAreaMonitorSource) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QGeoAreaMonitorSource) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QGeoAreaMonitorSource) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QGeoAreaMonitorSource) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QGeoAreaMonitorSource) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QGeoAreaMonitorSource) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QGeoAreaMonitorSource) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QGeoCircle struct {
	QGeoShape
}

type QGeoCircle_ITF interface {
	QGeoShape_ITF
	QGeoCircle_PTR() *QGeoCircle
}

func (ptr *QGeoCircle) QGeoCircle_PTR() *QGeoCircle {
	return ptr
}

func (ptr *QGeoCircle) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoShape_PTR().Pointer()
	}
	return nil
}

func (ptr *QGeoCircle) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QGeoShape_PTR().SetPointer(p)
	}
}

func PointerFromQGeoCircle(ptr QGeoCircle_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoCircle_PTR().Pointer()
	}
	return nil
}

func (n *QGeoCircle) InitFromInternal(ptr uintptr, name string) {
	n.QGeoShape_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QGeoCircle) ClassNameInternalF() string {
	return n.QGeoShape_PTR().ClassNameInternalF()
}

func NewQGeoCircleFromPointer(ptr unsafe.Pointer) (n *QGeoCircle) {
	n = new(QGeoCircle)
	n.InitFromInternal(uintptr(ptr), "positioning.QGeoCircle")
	return
}
func NewQGeoCircle() *QGeoCircle {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.NewQGeoCircle", ""}).(*QGeoCircle)
}

func NewQGeoCircle2(center QGeoCoordinate_ITF, radius float64) *QGeoCircle {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.NewQGeoCircle2", "", center, radius}).(*QGeoCircle)
}

func NewQGeoCircle3(other QGeoCircle_ITF) *QGeoCircle {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.NewQGeoCircle3", "", other}).(*QGeoCircle)
}

func NewQGeoCircle4(other QGeoShape_ITF) *QGeoCircle {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.NewQGeoCircle4", "", other}).(*QGeoCircle)
}

func (ptr *QGeoCircle) ExtendCircle(coordinate QGeoCoordinate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ExtendCircle", coordinate})
}

func (ptr *QGeoCircle) Radius() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Radius"}).(float64)
}

func (ptr *QGeoCircle) SetCenter(center QGeoCoordinate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetCenter", center})
}

func (ptr *QGeoCircle) SetRadius(radius float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRadius", radius})
}

func (ptr *QGeoCircle) Translate(degreesLatitude float64, degreesLongitude float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Translate", degreesLatitude, degreesLongitude})
}

func (ptr *QGeoCircle) Translated(degreesLatitude float64, degreesLongitude float64) *QGeoCircle {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Translated", degreesLatitude, degreesLongitude}).(*QGeoCircle)
}

func (ptr *QGeoCircle) DestroyQGeoCircle() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQGeoCircle"})
}

type QGeoCoordinate struct {
	internal.Internal
}

type QGeoCoordinate_ITF interface {
	QGeoCoordinate_PTR() *QGeoCoordinate
}

func (ptr *QGeoCoordinate) QGeoCoordinate_PTR() *QGeoCoordinate {
	return ptr
}

func (ptr *QGeoCoordinate) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QGeoCoordinate) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQGeoCoordinate(ptr QGeoCoordinate_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoCoordinate_PTR().Pointer()
	}
	return nil
}

func (n *QGeoCoordinate) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQGeoCoordinateFromPointer(ptr unsafe.Pointer) (n *QGeoCoordinate) {
	n = new(QGeoCoordinate)
	n.InitFromInternal(uintptr(ptr), "positioning.QGeoCoordinate")
	return
}

//go:generate stringer -type=QGeoCoordinate__CoordinateType
//QGeoCoordinate::CoordinateType
type QGeoCoordinate__CoordinateType int64

const (
	QGeoCoordinate__InvalidCoordinate QGeoCoordinate__CoordinateType = QGeoCoordinate__CoordinateType(0)
	QGeoCoordinate__Coordinate2D      QGeoCoordinate__CoordinateType = QGeoCoordinate__CoordinateType(1)
	QGeoCoordinate__Coordinate3D      QGeoCoordinate__CoordinateType = QGeoCoordinate__CoordinateType(2)
)

//go:generate stringer -type=QGeoCoordinate__CoordinateFormat
//QGeoCoordinate::CoordinateFormat
type QGeoCoordinate__CoordinateFormat int64

const (
	QGeoCoordinate__Degrees                             QGeoCoordinate__CoordinateFormat = QGeoCoordinate__CoordinateFormat(0)
	QGeoCoordinate__DegreesWithHemisphere               QGeoCoordinate__CoordinateFormat = QGeoCoordinate__CoordinateFormat(1)
	QGeoCoordinate__DegreesMinutes                      QGeoCoordinate__CoordinateFormat = QGeoCoordinate__CoordinateFormat(2)
	QGeoCoordinate__DegreesMinutesWithHemisphere        QGeoCoordinate__CoordinateFormat = QGeoCoordinate__CoordinateFormat(3)
	QGeoCoordinate__DegreesMinutesSeconds               QGeoCoordinate__CoordinateFormat = QGeoCoordinate__CoordinateFormat(4)
	QGeoCoordinate__DegreesMinutesSecondsWithHemisphere QGeoCoordinate__CoordinateFormat = QGeoCoordinate__CoordinateFormat(5)
)

func NewQGeoCoordinate() *QGeoCoordinate {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.NewQGeoCoordinate", ""}).(*QGeoCoordinate)
}

func NewQGeoCoordinate2(latitude float64, longitude float64) *QGeoCoordinate {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.NewQGeoCoordinate2", "", latitude, longitude}).(*QGeoCoordinate)
}

func NewQGeoCoordinate3(latitude float64, longitude float64, altitude float64) *QGeoCoordinate {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.NewQGeoCoordinate3", "", latitude, longitude, altitude}).(*QGeoCoordinate)
}

func NewQGeoCoordinate4(other QGeoCoordinate_ITF) *QGeoCoordinate {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.NewQGeoCoordinate4", "", other}).(*QGeoCoordinate)
}

func (ptr *QGeoCoordinate) Altitude() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Altitude"}).(float64)
}

func (ptr *QGeoCoordinate) AtDistanceAndAzimuth(distance float64, azimuth float64, distanceUp float64) *QGeoCoordinate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AtDistanceAndAzimuth", distance, azimuth, distanceUp}).(*QGeoCoordinate)
}

func (ptr *QGeoCoordinate) AzimuthTo(other QGeoCoordinate_ITF) float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AzimuthTo", other}).(float64)
}

func (ptr *QGeoCoordinate) DistanceTo(other QGeoCoordinate_ITF) float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DistanceTo", other}).(float64)
}

func (ptr *QGeoCoordinate) IsValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsValid"}).(bool)
}

func (ptr *QGeoCoordinate) Latitude() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Latitude"}).(float64)
}

func (ptr *QGeoCoordinate) Longitude() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Longitude"}).(float64)
}

func (ptr *QGeoCoordinate) SetAltitude(altitude float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAltitude", altitude})
}

func (ptr *QGeoCoordinate) SetLatitude(latitude float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLatitude", latitude})
}

func (ptr *QGeoCoordinate) SetLongitude(longitude float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLongitude", longitude})
}

func (ptr *QGeoCoordinate) ToString(format QGeoCoordinate__CoordinateFormat) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToString", format}).(string)
}

func (ptr *QGeoCoordinate) Type() QGeoCoordinate__CoordinateType {

	return QGeoCoordinate__CoordinateType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(float64))
}

func (ptr *QGeoCoordinate) DestroyQGeoCoordinate() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQGeoCoordinate"})
}

type QGeoLocation struct {
	internal.Internal
}

type QGeoLocation_ITF interface {
	QGeoLocation_PTR() *QGeoLocation
}

func (ptr *QGeoLocation) QGeoLocation_PTR() *QGeoLocation {
	return ptr
}

func (ptr *QGeoLocation) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QGeoLocation) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQGeoLocation(ptr QGeoLocation_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoLocation_PTR().Pointer()
	}
	return nil
}

func (n *QGeoLocation) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQGeoLocationFromPointer(ptr unsafe.Pointer) (n *QGeoLocation) {
	n = new(QGeoLocation)
	n.InitFromInternal(uintptr(ptr), "positioning.QGeoLocation")
	return
}

type QGeoPath struct {
	QGeoShape
}

type QGeoPath_ITF interface {
	QGeoShape_ITF
	QGeoPath_PTR() *QGeoPath
}

func (ptr *QGeoPath) QGeoPath_PTR() *QGeoPath {
	return ptr
}

func (ptr *QGeoPath) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoShape_PTR().Pointer()
	}
	return nil
}

func (ptr *QGeoPath) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QGeoShape_PTR().SetPointer(p)
	}
}

func PointerFromQGeoPath(ptr QGeoPath_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoPath_PTR().Pointer()
	}
	return nil
}

func (n *QGeoPath) InitFromInternal(ptr uintptr, name string) {
	n.QGeoShape_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QGeoPath) ClassNameInternalF() string {
	return n.QGeoShape_PTR().ClassNameInternalF()
}

func NewQGeoPathFromPointer(ptr unsafe.Pointer) (n *QGeoPath) {
	n = new(QGeoPath)
	n.InitFromInternal(uintptr(ptr), "positioning.QGeoPath")
	return
}
func NewQGeoPath() *QGeoPath {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.NewQGeoPath", ""}).(*QGeoPath)
}

func NewQGeoPath2(path []*QGeoCoordinate, width float64) *QGeoPath {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.NewQGeoPath2", "", path, width}).(*QGeoPath)
}

func NewQGeoPath3(other QGeoPath_ITF) *QGeoPath {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.NewQGeoPath3", "", other}).(*QGeoPath)
}

func NewQGeoPath4(other QGeoShape_ITF) *QGeoPath {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.NewQGeoPath4", "", other}).(*QGeoPath)
}

func (ptr *QGeoPath) AddCoordinate(coordinate QGeoCoordinate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddCoordinate", coordinate})
}

func (ptr *QGeoPath) ClearPath() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ClearPath"})
}

func (ptr *QGeoPath) ContainsCoordinate(coordinate QGeoCoordinate_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ContainsCoordinate", coordinate}).(bool)
}

func (ptr *QGeoPath) CoordinateAt(index int) *QGeoCoordinate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CoordinateAt", index}).(*QGeoCoordinate)
}

func (ptr *QGeoPath) InsertCoordinate(index int, coordinate QGeoCoordinate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InsertCoordinate", index, coordinate})
}

func (ptr *QGeoPath) Length(indexFrom int, indexTo int) float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Length", indexFrom, indexTo}).(float64)
}

func (ptr *QGeoPath) Path() []*QGeoCoordinate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Path"}).([]*QGeoCoordinate)
}

func (ptr *QGeoPath) RemoveCoordinate(coordinate QGeoCoordinate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveCoordinate", coordinate})
}

func (ptr *QGeoPath) RemoveCoordinate2(index int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveCoordinate2", index})
}

func (ptr *QGeoPath) ReplaceCoordinate(index int, coordinate QGeoCoordinate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReplaceCoordinate", index, coordinate})
}

func (ptr *QGeoPath) SetPath(path []*QGeoCoordinate) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPath", path})
}

func (ptr *QGeoPath) SetWidth(width float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWidth", width})
}

func (ptr *QGeoPath) Size() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Size"}).(float64))
}

func (ptr *QGeoPath) Translate(degreesLatitude float64, degreesLongitude float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Translate", degreesLatitude, degreesLongitude})
}

func (ptr *QGeoPath) Translated(degreesLatitude float64, degreesLongitude float64) *QGeoPath {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Translated", degreesLatitude, degreesLongitude}).(*QGeoPath)
}

func (ptr *QGeoPath) Width() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Width"}).(float64)
}

func (ptr *QGeoPath) DestroyQGeoPath() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQGeoPath"})
}

func (ptr *QGeoPath) __QGeoPath_path_atList2(i int) *QGeoCoordinate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__QGeoPath_path_atList2", i}).(*QGeoCoordinate)
}

func (ptr *QGeoPath) __QGeoPath_path_setList2(i QGeoCoordinate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__QGeoPath_path_setList2", i})
}

func (ptr *QGeoPath) __QGeoPath_path_newList2() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__QGeoPath_path_newList2"}).(unsafe.Pointer)
}

func (ptr *QGeoPath) __path_atList(i int) *QGeoCoordinate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__path_atList", i}).(*QGeoCoordinate)
}

func (ptr *QGeoPath) __path_setList(i QGeoCoordinate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__path_setList", i})
}

func (ptr *QGeoPath) __path_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__path_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoPath) __setPath_path_atList(i int) *QGeoCoordinate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setPath_path_atList", i}).(*QGeoCoordinate)
}

func (ptr *QGeoPath) __setPath_path_setList(i QGeoCoordinate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setPath_path_setList", i})
}

func (ptr *QGeoPath) __setPath_path_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setPath_path_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoPath) __setVariantPath_path_atList(i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setVariantPath_path_atList", i}).(*core.QVariant)
}

func (ptr *QGeoPath) __setVariantPath_path_setList(i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setVariantPath_path_setList", i})
}

func (ptr *QGeoPath) __setVariantPath_path_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setVariantPath_path_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoPath) __variantPath_atList(i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__variantPath_atList", i}).(*core.QVariant)
}

func (ptr *QGeoPath) __variantPath_setList(i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__variantPath_setList", i})
}

func (ptr *QGeoPath) __variantPath_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__variantPath_newList"}).(unsafe.Pointer)
}

type QGeoPolygon struct {
	QGeoShape
}

type QGeoPolygon_ITF interface {
	QGeoShape_ITF
	QGeoPolygon_PTR() *QGeoPolygon
}

func (ptr *QGeoPolygon) QGeoPolygon_PTR() *QGeoPolygon {
	return ptr
}

func (ptr *QGeoPolygon) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoShape_PTR().Pointer()
	}
	return nil
}

func (ptr *QGeoPolygon) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QGeoShape_PTR().SetPointer(p)
	}
}

func PointerFromQGeoPolygon(ptr QGeoPolygon_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoPolygon_PTR().Pointer()
	}
	return nil
}

func (n *QGeoPolygon) InitFromInternal(ptr uintptr, name string) {
	n.QGeoShape_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QGeoPolygon) ClassNameInternalF() string {
	return n.QGeoShape_PTR().ClassNameInternalF()
}

func NewQGeoPolygonFromPointer(ptr unsafe.Pointer) (n *QGeoPolygon) {
	n = new(QGeoPolygon)
	n.InitFromInternal(uintptr(ptr), "positioning.QGeoPolygon")
	return
}
func NewQGeoPolygon() *QGeoPolygon {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.NewQGeoPolygon", ""}).(*QGeoPolygon)
}

func NewQGeoPolygon2(path []*QGeoCoordinate) *QGeoPolygon {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.NewQGeoPolygon2", "", path}).(*QGeoPolygon)
}

func NewQGeoPolygon3(other QGeoPolygon_ITF) *QGeoPolygon {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.NewQGeoPolygon3", "", other}).(*QGeoPolygon)
}

func NewQGeoPolygon4(other QGeoShape_ITF) *QGeoPolygon {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.NewQGeoPolygon4", "", other}).(*QGeoPolygon)
}

func (ptr *QGeoPolygon) AddCoordinate(coordinate QGeoCoordinate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddCoordinate", coordinate})
}

func (ptr *QGeoPolygon) AddHole(holePath core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddHole", holePath})
}

func (ptr *QGeoPolygon) AddHole2(holePath []*QGeoCoordinate) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddHole2", holePath})
}

func (ptr *QGeoPolygon) ContainsCoordinate(coordinate QGeoCoordinate_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ContainsCoordinate", coordinate}).(bool)
}

func (ptr *QGeoPolygon) CoordinateAt(index int) *QGeoCoordinate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CoordinateAt", index}).(*QGeoCoordinate)
}

func (ptr *QGeoPolygon) Hole(index int) []*core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Hole", index}).([]*core.QVariant)
}

func (ptr *QGeoPolygon) HolePath(index int) []*QGeoCoordinate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HolePath", index}).([]*QGeoCoordinate)
}

func (ptr *QGeoPolygon) HolesCount() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HolesCount"}).(float64))
}

func (ptr *QGeoPolygon) InsertCoordinate(index int, coordinate QGeoCoordinate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InsertCoordinate", index, coordinate})
}

func (ptr *QGeoPolygon) Length(indexFrom int, indexTo int) float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Length", indexFrom, indexTo}).(float64)
}

func (ptr *QGeoPolygon) Path() []*QGeoCoordinate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Path"}).([]*QGeoCoordinate)
}

func (ptr *QGeoPolygon) Perimeter() []*core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Perimeter"}).([]*core.QVariant)
}

func (ptr *QGeoPolygon) RemoveCoordinate(coordinate QGeoCoordinate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveCoordinate", coordinate})
}

func (ptr *QGeoPolygon) RemoveCoordinate2(index int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveCoordinate2", index})
}

func (ptr *QGeoPolygon) RemoveHole(index int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveHole", index})
}

func (ptr *QGeoPolygon) ReplaceCoordinate(index int, coordinate QGeoCoordinate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReplaceCoordinate", index, coordinate})
}

func (ptr *QGeoPolygon) SetPath(path []*QGeoCoordinate) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPath", path})
}

func (ptr *QGeoPolygon) SetPerimeter(path []*core.QVariant) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPerimeter", path})
}

func (ptr *QGeoPolygon) Size() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Size"}).(float64))
}

func (ptr *QGeoPolygon) Translate(degreesLatitude float64, degreesLongitude float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Translate", degreesLatitude, degreesLongitude})
}

func (ptr *QGeoPolygon) Translated(degreesLatitude float64, degreesLongitude float64) *QGeoPolygon {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Translated", degreesLatitude, degreesLongitude}).(*QGeoPolygon)
}

func (ptr *QGeoPolygon) DestroyQGeoPolygon() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQGeoPolygon"})
}

func (ptr *QGeoPolygon) __QGeoPolygon_path_atList2(i int) *QGeoCoordinate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__QGeoPolygon_path_atList2", i}).(*QGeoCoordinate)
}

func (ptr *QGeoPolygon) __QGeoPolygon_path_setList2(i QGeoCoordinate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__QGeoPolygon_path_setList2", i})
}

func (ptr *QGeoPolygon) __QGeoPolygon_path_newList2() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__QGeoPolygon_path_newList2"}).(unsafe.Pointer)
}

func (ptr *QGeoPolygon) __addHole_holePath_atList2(i int) *QGeoCoordinate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addHole_holePath_atList2", i}).(*QGeoCoordinate)
}

func (ptr *QGeoPolygon) __addHole_holePath_setList2(i QGeoCoordinate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addHole_holePath_setList2", i})
}

func (ptr *QGeoPolygon) __addHole_holePath_newList2() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addHole_holePath_newList2"}).(unsafe.Pointer)
}

func (ptr *QGeoPolygon) __hole_atList(i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__hole_atList", i}).(*core.QVariant)
}

func (ptr *QGeoPolygon) __hole_setList(i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__hole_setList", i})
}

func (ptr *QGeoPolygon) __hole_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__hole_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoPolygon) __holePath_atList(i int) *QGeoCoordinate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__holePath_atList", i}).(*QGeoCoordinate)
}

func (ptr *QGeoPolygon) __holePath_setList(i QGeoCoordinate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__holePath_setList", i})
}

func (ptr *QGeoPolygon) __holePath_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__holePath_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoPolygon) __path_atList(i int) *QGeoCoordinate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__path_atList", i}).(*QGeoCoordinate)
}

func (ptr *QGeoPolygon) __path_setList(i QGeoCoordinate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__path_setList", i})
}

func (ptr *QGeoPolygon) __path_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__path_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoPolygon) __perimeter_atList(i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__perimeter_atList", i}).(*core.QVariant)
}

func (ptr *QGeoPolygon) __perimeter_setList(i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__perimeter_setList", i})
}

func (ptr *QGeoPolygon) __perimeter_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__perimeter_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoPolygon) __setPath_path_atList(i int) *QGeoCoordinate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setPath_path_atList", i}).(*QGeoCoordinate)
}

func (ptr *QGeoPolygon) __setPath_path_setList(i QGeoCoordinate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setPath_path_setList", i})
}

func (ptr *QGeoPolygon) __setPath_path_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setPath_path_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoPolygon) __setPerimeter_path_atList(i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setPerimeter_path_atList", i}).(*core.QVariant)
}

func (ptr *QGeoPolygon) __setPerimeter_path_setList(i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setPerimeter_path_setList", i})
}

func (ptr *QGeoPolygon) __setPerimeter_path_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setPerimeter_path_newList"}).(unsafe.Pointer)
}

type QGeoPositionInfo struct {
	internal.Internal
}

type QGeoPositionInfo_ITF interface {
	QGeoPositionInfo_PTR() *QGeoPositionInfo
}

func (ptr *QGeoPositionInfo) QGeoPositionInfo_PTR() *QGeoPositionInfo {
	return ptr
}

func (ptr *QGeoPositionInfo) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QGeoPositionInfo) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQGeoPositionInfo(ptr QGeoPositionInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoPositionInfo_PTR().Pointer()
	}
	return nil
}

func (n *QGeoPositionInfo) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQGeoPositionInfoFromPointer(ptr unsafe.Pointer) (n *QGeoPositionInfo) {
	n = new(QGeoPositionInfo)
	n.InitFromInternal(uintptr(ptr), "positioning.QGeoPositionInfo")
	return
}

//go:generate stringer -type=QGeoPositionInfo__Attribute
//QGeoPositionInfo::Attribute
type QGeoPositionInfo__Attribute int64

const (
	QGeoPositionInfo__Direction          QGeoPositionInfo__Attribute = QGeoPositionInfo__Attribute(0)
	QGeoPositionInfo__GroundSpeed        QGeoPositionInfo__Attribute = QGeoPositionInfo__Attribute(1)
	QGeoPositionInfo__VerticalSpeed      QGeoPositionInfo__Attribute = QGeoPositionInfo__Attribute(2)
	QGeoPositionInfo__MagneticVariation  QGeoPositionInfo__Attribute = QGeoPositionInfo__Attribute(3)
	QGeoPositionInfo__HorizontalAccuracy QGeoPositionInfo__Attribute = QGeoPositionInfo__Attribute(4)
	QGeoPositionInfo__VerticalAccuracy   QGeoPositionInfo__Attribute = QGeoPositionInfo__Attribute(5)
)

func NewQGeoPositionInfo() *QGeoPositionInfo {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.NewQGeoPositionInfo", ""}).(*QGeoPositionInfo)
}

func NewQGeoPositionInfo2(coordinate QGeoCoordinate_ITF, timestamp core.QDateTime_ITF) *QGeoPositionInfo {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.NewQGeoPositionInfo2", "", coordinate, timestamp}).(*QGeoPositionInfo)
}

func NewQGeoPositionInfo3(other QGeoPositionInfo_ITF) *QGeoPositionInfo {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.NewQGeoPositionInfo3", "", other}).(*QGeoPositionInfo)
}

func (ptr *QGeoPositionInfo) Attribute(attribute QGeoPositionInfo__Attribute) float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Attribute", attribute}).(float64)
}

func (ptr *QGeoPositionInfo) Coordinate() *QGeoCoordinate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Coordinate"}).(*QGeoCoordinate)
}

func (ptr *QGeoPositionInfo) HasAttribute(attribute QGeoPositionInfo__Attribute) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasAttribute", attribute}).(bool)
}

func (ptr *QGeoPositionInfo) IsValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsValid"}).(bool)
}

func (ptr *QGeoPositionInfo) RemoveAttribute(attribute QGeoPositionInfo__Attribute) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveAttribute", attribute})
}

func (ptr *QGeoPositionInfo) SetAttribute(attribute QGeoPositionInfo__Attribute, value float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAttribute", attribute, value})
}

func (ptr *QGeoPositionInfo) SetCoordinate(coordinate QGeoCoordinate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetCoordinate", coordinate})
}

func (ptr *QGeoPositionInfo) SetTimestamp(timestamp core.QDateTime_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTimestamp", timestamp})
}

func (ptr *QGeoPositionInfo) Timestamp() *core.QDateTime {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Timestamp"}).(*core.QDateTime)
}

func (ptr *QGeoPositionInfo) DestroyQGeoPositionInfo() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQGeoPositionInfo"})
}

type QGeoPositionInfoSource struct {
	core.QObject
}

type QGeoPositionInfoSource_ITF interface {
	core.QObject_ITF
	QGeoPositionInfoSource_PTR() *QGeoPositionInfoSource
}

func (ptr *QGeoPositionInfoSource) QGeoPositionInfoSource_PTR() *QGeoPositionInfoSource {
	return ptr
}

func (ptr *QGeoPositionInfoSource) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QGeoPositionInfoSource) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQGeoPositionInfoSource(ptr QGeoPositionInfoSource_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoPositionInfoSource_PTR().Pointer()
	}
	return nil
}

func (n *QGeoPositionInfoSource) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QGeoPositionInfoSource) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQGeoPositionInfoSourceFromPointer(ptr unsafe.Pointer) (n *QGeoPositionInfoSource) {
	n = new(QGeoPositionInfoSource)
	n.InitFromInternal(uintptr(ptr), "positioning.QGeoPositionInfoSource")
	return
}

//go:generate stringer -type=QGeoPositionInfoSource__Error
//QGeoPositionInfoSource::Error
type QGeoPositionInfoSource__Error int64

const (
	QGeoPositionInfoSource__AccessError        QGeoPositionInfoSource__Error = QGeoPositionInfoSource__Error(0)
	QGeoPositionInfoSource__ClosedError        QGeoPositionInfoSource__Error = QGeoPositionInfoSource__Error(1)
	QGeoPositionInfoSource__UnknownSourceError QGeoPositionInfoSource__Error = QGeoPositionInfoSource__Error(2)
	QGeoPositionInfoSource__NoError            QGeoPositionInfoSource__Error = QGeoPositionInfoSource__Error(3)
)

//go:generate stringer -type=QGeoPositionInfoSource__PositioningMethod
//QGeoPositionInfoSource::PositioningMethod
type QGeoPositionInfoSource__PositioningMethod int64

const (
	QGeoPositionInfoSource__NoPositioningMethods           QGeoPositionInfoSource__PositioningMethod = QGeoPositionInfoSource__PositioningMethod(0x00000000)
	QGeoPositionInfoSource__SatellitePositioningMethods    QGeoPositionInfoSource__PositioningMethod = QGeoPositionInfoSource__PositioningMethod(0x000000ff)
	QGeoPositionInfoSource__NonSatellitePositioningMethods QGeoPositionInfoSource__PositioningMethod = QGeoPositionInfoSource__PositioningMethod(0xffffff00)
	QGeoPositionInfoSource__AllPositioningMethods          QGeoPositionInfoSource__PositioningMethod = QGeoPositionInfoSource__PositioningMethod(0xffffffff)
)

func NewQGeoPositionInfoSource(parent core.QObject_ITF) *QGeoPositionInfoSource {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.NewQGeoPositionInfoSource", "", parent}).(*QGeoPositionInfoSource)
}

func QGeoPositionInfoSource_AvailableSources() []string {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.QGeoPositionInfoSource_AvailableSources", ""}).([]string)
}

func (ptr *QGeoPositionInfoSource) AvailableSources() []string {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.QGeoPositionInfoSource_AvailableSources", ""}).([]string)
}

func QGeoPositionInfoSource_CreateDefaultSource(parent core.QObject_ITF) *QGeoPositionInfoSource {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.QGeoPositionInfoSource_CreateDefaultSource", "", parent}).(*QGeoPositionInfoSource)
}

func (ptr *QGeoPositionInfoSource) CreateDefaultSource(parent core.QObject_ITF) *QGeoPositionInfoSource {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.QGeoPositionInfoSource_CreateDefaultSource", "", parent}).(*QGeoPositionInfoSource)
}

func QGeoPositionInfoSource_CreateSource(sourceName string, parent core.QObject_ITF) *QGeoPositionInfoSource {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.QGeoPositionInfoSource_CreateSource", "", sourceName, parent}).(*QGeoPositionInfoSource)
}

func (ptr *QGeoPositionInfoSource) CreateSource(sourceName string, parent core.QObject_ITF) *QGeoPositionInfoSource {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.QGeoPositionInfoSource_CreateSource", "", sourceName, parent}).(*QGeoPositionInfoSource)
}

func (ptr *QGeoPositionInfoSource) ConnectError(f func() QGeoPositionInfoSource__Error) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectError", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoPositionInfoSource) DisconnectError() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectError"})
}

func (ptr *QGeoPositionInfoSource) Error() QGeoPositionInfoSource__Error {

	return QGeoPositionInfoSource__Error(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error"}).(float64))
}

func (ptr *QGeoPositionInfoSource) ConnectError2(f func(positioningError QGeoPositionInfoSource__Error)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectError2", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoPositionInfoSource) DisconnectError2() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectError2"})
}

func (ptr *QGeoPositionInfoSource) Error2(positioningError QGeoPositionInfoSource__Error) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error2", positioningError})
}

func (ptr *QGeoPositionInfoSource) ConnectLastKnownPosition(f func(fromSatellitePositioningMethodsOnly bool) *QGeoPositionInfo) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLastKnownPosition", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoPositionInfoSource) DisconnectLastKnownPosition() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLastKnownPosition"})
}

func (ptr *QGeoPositionInfoSource) LastKnownPosition(fromSatellitePositioningMethodsOnly bool) *QGeoPositionInfo {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LastKnownPosition", fromSatellitePositioningMethodsOnly}).(*QGeoPositionInfo)
}

func (ptr *QGeoPositionInfoSource) ConnectMinimumUpdateInterval(f func() int) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMinimumUpdateInterval", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoPositionInfoSource) DisconnectMinimumUpdateInterval() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMinimumUpdateInterval"})
}

func (ptr *QGeoPositionInfoSource) MinimumUpdateInterval() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinimumUpdateInterval"}).(float64))
}

func (ptr *QGeoPositionInfoSource) ConnectPositionUpdated(f func(update *QGeoPositionInfo)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPositionUpdated", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoPositionInfoSource) DisconnectPositionUpdated() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPositionUpdated"})
}

func (ptr *QGeoPositionInfoSource) PositionUpdated(update QGeoPositionInfo_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PositionUpdated", update})
}

func (ptr *QGeoPositionInfoSource) PreferredPositioningMethods() QGeoPositionInfoSource__PositioningMethod {

	return QGeoPositionInfoSource__PositioningMethod(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PreferredPositioningMethods"}).(float64))
}

func (ptr *QGeoPositionInfoSource) ConnectRequestUpdate(f func(timeout int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRequestUpdate", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoPositionInfoSource) DisconnectRequestUpdate() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRequestUpdate"})
}

func (ptr *QGeoPositionInfoSource) RequestUpdate(timeout int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RequestUpdate", timeout})
}

func (ptr *QGeoPositionInfoSource) ConnectSetPreferredPositioningMethods(f func(methods QGeoPositionInfoSource__PositioningMethod)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetPreferredPositioningMethods", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoPositionInfoSource) DisconnectSetPreferredPositioningMethods() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetPreferredPositioningMethods"})
}

func (ptr *QGeoPositionInfoSource) SetPreferredPositioningMethods(methods QGeoPositionInfoSource__PositioningMethod) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPreferredPositioningMethods", methods})
}

func (ptr *QGeoPositionInfoSource) SetPreferredPositioningMethodsDefault(methods QGeoPositionInfoSource__PositioningMethod) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPreferredPositioningMethodsDefault", methods})
}

func (ptr *QGeoPositionInfoSource) ConnectSetUpdateInterval(f func(msec int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetUpdateInterval", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoPositionInfoSource) DisconnectSetUpdateInterval() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetUpdateInterval"})
}

func (ptr *QGeoPositionInfoSource) SetUpdateInterval(msec int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetUpdateInterval", msec})
}

func (ptr *QGeoPositionInfoSource) SetUpdateIntervalDefault(msec int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetUpdateIntervalDefault", msec})
}

func (ptr *QGeoPositionInfoSource) SourceName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SourceName"}).(string)
}

func (ptr *QGeoPositionInfoSource) ConnectStartUpdates(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStartUpdates", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoPositionInfoSource) DisconnectStartUpdates() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStartUpdates"})
}

func (ptr *QGeoPositionInfoSource) StartUpdates() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StartUpdates"})
}

func (ptr *QGeoPositionInfoSource) ConnectStopUpdates(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStopUpdates", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoPositionInfoSource) DisconnectStopUpdates() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStopUpdates"})
}

func (ptr *QGeoPositionInfoSource) StopUpdates() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StopUpdates"})
}

func (ptr *QGeoPositionInfoSource) ConnectSupportedPositioningMethods(f func() QGeoPositionInfoSource__PositioningMethod) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSupportedPositioningMethods", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoPositionInfoSource) DisconnectSupportedPositioningMethods() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSupportedPositioningMethods"})
}

func (ptr *QGeoPositionInfoSource) SupportedPositioningMethods() QGeoPositionInfoSource__PositioningMethod {

	return QGeoPositionInfoSource__PositioningMethod(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SupportedPositioningMethods"}).(float64))
}

func (ptr *QGeoPositionInfoSource) ConnectSupportedPositioningMethodsChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSupportedPositioningMethodsChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoPositionInfoSource) DisconnectSupportedPositioningMethodsChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSupportedPositioningMethodsChanged"})
}

func (ptr *QGeoPositionInfoSource) SupportedPositioningMethodsChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SupportedPositioningMethodsChanged"})
}

func (ptr *QGeoPositionInfoSource) UpdateInterval() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateInterval"}).(float64))
}

func (ptr *QGeoPositionInfoSource) ConnectUpdateTimeout(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectUpdateTimeout", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoPositionInfoSource) DisconnectUpdateTimeout() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectUpdateTimeout"})
}

func (ptr *QGeoPositionInfoSource) UpdateTimeout() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateTimeout"})
}

func (ptr *QGeoPositionInfoSource) ConnectDestroyQGeoPositionInfoSource(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQGeoPositionInfoSource", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoPositionInfoSource) DisconnectDestroyQGeoPositionInfoSource() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQGeoPositionInfoSource"})
}

func (ptr *QGeoPositionInfoSource) DestroyQGeoPositionInfoSource() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQGeoPositionInfoSource"})
}

func (ptr *QGeoPositionInfoSource) DestroyQGeoPositionInfoSourceDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQGeoPositionInfoSourceDefault"})
}

func (ptr *QGeoPositionInfoSource) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QGeoPositionInfoSource) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QGeoPositionInfoSource) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoPositionInfoSource) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QGeoPositionInfoSource) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QGeoPositionInfoSource) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoPositionInfoSource) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QGeoPositionInfoSource) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QGeoPositionInfoSource) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoPositionInfoSource) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QGeoPositionInfoSource) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QGeoPositionInfoSource) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QGeoPositionInfoSource) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QGeoPositionInfoSource) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QGeoPositionInfoSource) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QGeoPositionInfoSource) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QGeoPositionInfoSource) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QGeoPositionInfoSource) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QGeoPositionInfoSource) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QGeoPositionInfoSource) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QGeoPositionInfoSource) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QGeoPositionInfoSourceFactory struct {
	internal.Internal
}

type QGeoPositionInfoSourceFactory_ITF interface {
	QGeoPositionInfoSourceFactory_PTR() *QGeoPositionInfoSourceFactory
}

func (ptr *QGeoPositionInfoSourceFactory) QGeoPositionInfoSourceFactory_PTR() *QGeoPositionInfoSourceFactory {
	return ptr
}

func (ptr *QGeoPositionInfoSourceFactory) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QGeoPositionInfoSourceFactory) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQGeoPositionInfoSourceFactory(ptr QGeoPositionInfoSourceFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoPositionInfoSourceFactory_PTR().Pointer()
	}
	return nil
}

func (n *QGeoPositionInfoSourceFactory) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQGeoPositionInfoSourceFactoryFromPointer(ptr unsafe.Pointer) (n *QGeoPositionInfoSourceFactory) {
	n = new(QGeoPositionInfoSourceFactory)
	n.InitFromInternal(uintptr(ptr), "positioning.QGeoPositionInfoSourceFactory")
	return
}
func (ptr *QGeoPositionInfoSourceFactory) ConnectAreaMonitor(f func(parent *core.QObject) *QGeoAreaMonitorSource) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAreaMonitor", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoPositionInfoSourceFactory) DisconnectAreaMonitor() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAreaMonitor"})
}

func (ptr *QGeoPositionInfoSourceFactory) AreaMonitor(parent core.QObject_ITF) *QGeoAreaMonitorSource {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AreaMonitor", parent}).(*QGeoAreaMonitorSource)
}

func (ptr *QGeoPositionInfoSourceFactory) ConnectPositionInfoSource(f func(parent *core.QObject) *QGeoPositionInfoSource) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPositionInfoSource", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoPositionInfoSourceFactory) DisconnectPositionInfoSource() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPositionInfoSource"})
}

func (ptr *QGeoPositionInfoSourceFactory) PositionInfoSource(parent core.QObject_ITF) *QGeoPositionInfoSource {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PositionInfoSource", parent}).(*QGeoPositionInfoSource)
}

func (ptr *QGeoPositionInfoSourceFactory) ConnectSatelliteInfoSource(f func(parent *core.QObject) *QGeoSatelliteInfoSource) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSatelliteInfoSource", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoPositionInfoSourceFactory) DisconnectSatelliteInfoSource() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSatelliteInfoSource"})
}

func (ptr *QGeoPositionInfoSourceFactory) SatelliteInfoSource(parent core.QObject_ITF) *QGeoSatelliteInfoSource {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SatelliteInfoSource", parent}).(*QGeoSatelliteInfoSource)
}

func (ptr *QGeoPositionInfoSourceFactory) ConnectDestroyQGeoPositionInfoSourceFactory(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQGeoPositionInfoSourceFactory", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoPositionInfoSourceFactory) DisconnectDestroyQGeoPositionInfoSourceFactory() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQGeoPositionInfoSourceFactory"})
}

func (ptr *QGeoPositionInfoSourceFactory) DestroyQGeoPositionInfoSourceFactory() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQGeoPositionInfoSourceFactory"})
}

func (ptr *QGeoPositionInfoSourceFactory) DestroyQGeoPositionInfoSourceFactoryDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQGeoPositionInfoSourceFactoryDefault"})
}

type QGeoRectangle struct {
	QGeoShape
}

type QGeoRectangle_ITF interface {
	QGeoShape_ITF
	QGeoRectangle_PTR() *QGeoRectangle
}

func (ptr *QGeoRectangle) QGeoRectangle_PTR() *QGeoRectangle {
	return ptr
}

func (ptr *QGeoRectangle) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoShape_PTR().Pointer()
	}
	return nil
}

func (ptr *QGeoRectangle) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QGeoShape_PTR().SetPointer(p)
	}
}

func PointerFromQGeoRectangle(ptr QGeoRectangle_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoRectangle_PTR().Pointer()
	}
	return nil
}

func (n *QGeoRectangle) InitFromInternal(ptr uintptr, name string) {
	n.QGeoShape_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QGeoRectangle) ClassNameInternalF() string {
	return n.QGeoShape_PTR().ClassNameInternalF()
}

func NewQGeoRectangleFromPointer(ptr unsafe.Pointer) (n *QGeoRectangle) {
	n = new(QGeoRectangle)
	n.InitFromInternal(uintptr(ptr), "positioning.QGeoRectangle")
	return
}
func NewQGeoRectangle() *QGeoRectangle {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.NewQGeoRectangle", ""}).(*QGeoRectangle)
}

func NewQGeoRectangle2(center QGeoCoordinate_ITF, degreesWidth float64, degreesHeight float64) *QGeoRectangle {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.NewQGeoRectangle2", "", center, degreesWidth, degreesHeight}).(*QGeoRectangle)
}

func NewQGeoRectangle3(topLeft QGeoCoordinate_ITF, bottomRight QGeoCoordinate_ITF) *QGeoRectangle {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.NewQGeoRectangle3", "", topLeft, bottomRight}).(*QGeoRectangle)
}

func NewQGeoRectangle4(coordinates []*QGeoCoordinate) *QGeoRectangle {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.NewQGeoRectangle4", "", coordinates}).(*QGeoRectangle)
}

func NewQGeoRectangle5(other QGeoRectangle_ITF) *QGeoRectangle {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.NewQGeoRectangle5", "", other}).(*QGeoRectangle)
}

func NewQGeoRectangle6(other QGeoShape_ITF) *QGeoRectangle {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.NewQGeoRectangle6", "", other}).(*QGeoRectangle)
}

func (ptr *QGeoRectangle) BottomLeft() *QGeoCoordinate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BottomLeft"}).(*QGeoCoordinate)
}

func (ptr *QGeoRectangle) BottomRight() *QGeoCoordinate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BottomRight"}).(*QGeoCoordinate)
}

func (ptr *QGeoRectangle) Contains(rectangle QGeoRectangle_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Contains", rectangle}).(bool)
}

func (ptr *QGeoRectangle) ExtendRectangle(coordinate QGeoCoordinate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ExtendRectangle", coordinate})
}

func (ptr *QGeoRectangle) Height() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Height"}).(float64)
}

func (ptr *QGeoRectangle) Intersects(rectangle QGeoRectangle_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Intersects", rectangle}).(bool)
}

func (ptr *QGeoRectangle) SetBottomLeft(bottomLeft QGeoCoordinate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBottomLeft", bottomLeft})
}

func (ptr *QGeoRectangle) SetBottomRight(bottomRight QGeoCoordinate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBottomRight", bottomRight})
}

func (ptr *QGeoRectangle) SetCenter(center QGeoCoordinate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetCenter", center})
}

func (ptr *QGeoRectangle) SetHeight(degreesHeight float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHeight", degreesHeight})
}

func (ptr *QGeoRectangle) SetTopLeft(topLeft QGeoCoordinate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTopLeft", topLeft})
}

func (ptr *QGeoRectangle) SetTopRight(topRight QGeoCoordinate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTopRight", topRight})
}

func (ptr *QGeoRectangle) SetWidth(degreesWidth float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWidth", degreesWidth})
}

func (ptr *QGeoRectangle) TopLeft() *QGeoCoordinate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TopLeft"}).(*QGeoCoordinate)
}

func (ptr *QGeoRectangle) TopRight() *QGeoCoordinate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TopRight"}).(*QGeoCoordinate)
}

func (ptr *QGeoRectangle) Translate(degreesLatitude float64, degreesLongitude float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Translate", degreesLatitude, degreesLongitude})
}

func (ptr *QGeoRectangle) Translated(degreesLatitude float64, degreesLongitude float64) *QGeoRectangle {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Translated", degreesLatitude, degreesLongitude}).(*QGeoRectangle)
}

func (ptr *QGeoRectangle) United(rectangle QGeoRectangle_ITF) *QGeoRectangle {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "United", rectangle}).(*QGeoRectangle)
}

func (ptr *QGeoRectangle) Width() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Width"}).(float64)
}

func (ptr *QGeoRectangle) DestroyQGeoRectangle() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQGeoRectangle"})
}

func (ptr *QGeoRectangle) __QGeoRectangle_coordinates_atList4(i int) *QGeoCoordinate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__QGeoRectangle_coordinates_atList4", i}).(*QGeoCoordinate)
}

func (ptr *QGeoRectangle) __QGeoRectangle_coordinates_setList4(i QGeoCoordinate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__QGeoRectangle_coordinates_setList4", i})
}

func (ptr *QGeoRectangle) __QGeoRectangle_coordinates_newList4() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__QGeoRectangle_coordinates_newList4"}).(unsafe.Pointer)
}

type QGeoSatelliteInfo struct {
	internal.Internal
}

type QGeoSatelliteInfo_ITF interface {
	QGeoSatelliteInfo_PTR() *QGeoSatelliteInfo
}

func (ptr *QGeoSatelliteInfo) QGeoSatelliteInfo_PTR() *QGeoSatelliteInfo {
	return ptr
}

func (ptr *QGeoSatelliteInfo) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QGeoSatelliteInfo) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQGeoSatelliteInfo(ptr QGeoSatelliteInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoSatelliteInfo_PTR().Pointer()
	}
	return nil
}

func (n *QGeoSatelliteInfo) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQGeoSatelliteInfoFromPointer(ptr unsafe.Pointer) (n *QGeoSatelliteInfo) {
	n = new(QGeoSatelliteInfo)
	n.InitFromInternal(uintptr(ptr), "positioning.QGeoSatelliteInfo")
	return
}

//go:generate stringer -type=QGeoSatelliteInfo__Attribute
//QGeoSatelliteInfo::Attribute
type QGeoSatelliteInfo__Attribute int64

const (
	QGeoSatelliteInfo__Elevation QGeoSatelliteInfo__Attribute = QGeoSatelliteInfo__Attribute(0)
	QGeoSatelliteInfo__Azimuth   QGeoSatelliteInfo__Attribute = QGeoSatelliteInfo__Attribute(1)
)

//go:generate stringer -type=QGeoSatelliteInfo__SatelliteSystem
//QGeoSatelliteInfo::SatelliteSystem
type QGeoSatelliteInfo__SatelliteSystem int64

const (
	QGeoSatelliteInfo__Undefined QGeoSatelliteInfo__SatelliteSystem = QGeoSatelliteInfo__SatelliteSystem(0x00)
	QGeoSatelliteInfo__GPS       QGeoSatelliteInfo__SatelliteSystem = QGeoSatelliteInfo__SatelliteSystem(0x01)
	QGeoSatelliteInfo__GLONASS   QGeoSatelliteInfo__SatelliteSystem = QGeoSatelliteInfo__SatelliteSystem(0x02)
)

func NewQGeoSatelliteInfo() *QGeoSatelliteInfo {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.NewQGeoSatelliteInfo", ""}).(*QGeoSatelliteInfo)
}

func NewQGeoSatelliteInfo2(other QGeoSatelliteInfo_ITF) *QGeoSatelliteInfo {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.NewQGeoSatelliteInfo2", "", other}).(*QGeoSatelliteInfo)
}

func (ptr *QGeoSatelliteInfo) Attribute(attribute QGeoSatelliteInfo__Attribute) float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Attribute", attribute}).(float64)
}

func (ptr *QGeoSatelliteInfo) HasAttribute(attribute QGeoSatelliteInfo__Attribute) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasAttribute", attribute}).(bool)
}

func (ptr *QGeoSatelliteInfo) RemoveAttribute(attribute QGeoSatelliteInfo__Attribute) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveAttribute", attribute})
}

func (ptr *QGeoSatelliteInfo) SatelliteIdentifier() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SatelliteIdentifier"}).(float64))
}

func (ptr *QGeoSatelliteInfo) SatelliteSystem() QGeoSatelliteInfo__SatelliteSystem {

	return QGeoSatelliteInfo__SatelliteSystem(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SatelliteSystem"}).(float64))
}

func (ptr *QGeoSatelliteInfo) SetAttribute(attribute QGeoSatelliteInfo__Attribute, value float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAttribute", attribute, value})
}

func (ptr *QGeoSatelliteInfo) SetSatelliteIdentifier(satId int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSatelliteIdentifier", satId})
}

func (ptr *QGeoSatelliteInfo) SetSatelliteSystem(system QGeoSatelliteInfo__SatelliteSystem) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSatelliteSystem", system})
}

func (ptr *QGeoSatelliteInfo) SetSignalStrength(signalStrength int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSignalStrength", signalStrength})
}

func (ptr *QGeoSatelliteInfo) SignalStrength() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SignalStrength"}).(float64))
}

func (ptr *QGeoSatelliteInfo) DestroyQGeoSatelliteInfo() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQGeoSatelliteInfo"})
}

type QGeoSatelliteInfoSource struct {
	core.QObject
}

type QGeoSatelliteInfoSource_ITF interface {
	core.QObject_ITF
	QGeoSatelliteInfoSource_PTR() *QGeoSatelliteInfoSource
}

func (ptr *QGeoSatelliteInfoSource) QGeoSatelliteInfoSource_PTR() *QGeoSatelliteInfoSource {
	return ptr
}

func (ptr *QGeoSatelliteInfoSource) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QGeoSatelliteInfoSource) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQGeoSatelliteInfoSource(ptr QGeoSatelliteInfoSource_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoSatelliteInfoSource_PTR().Pointer()
	}
	return nil
}

func (n *QGeoSatelliteInfoSource) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QGeoSatelliteInfoSource) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQGeoSatelliteInfoSourceFromPointer(ptr unsafe.Pointer) (n *QGeoSatelliteInfoSource) {
	n = new(QGeoSatelliteInfoSource)
	n.InitFromInternal(uintptr(ptr), "positioning.QGeoSatelliteInfoSource")
	return
}

//go:generate stringer -type=QGeoSatelliteInfoSource__Error
//QGeoSatelliteInfoSource::Error
type QGeoSatelliteInfoSource__Error int64

const (
	QGeoSatelliteInfoSource__AccessError        QGeoSatelliteInfoSource__Error = QGeoSatelliteInfoSource__Error(0)
	QGeoSatelliteInfoSource__ClosedError        QGeoSatelliteInfoSource__Error = QGeoSatelliteInfoSource__Error(1)
	QGeoSatelliteInfoSource__NoError            QGeoSatelliteInfoSource__Error = QGeoSatelliteInfoSource__Error(2)
	QGeoSatelliteInfoSource__UnknownSourceError QGeoSatelliteInfoSource__Error = QGeoSatelliteInfoSource__Error(-1)
)

func NewQGeoSatelliteInfoSource(parent core.QObject_ITF) *QGeoSatelliteInfoSource {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.NewQGeoSatelliteInfoSource", "", parent}).(*QGeoSatelliteInfoSource)
}

func QGeoSatelliteInfoSource_AvailableSources() []string {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.QGeoSatelliteInfoSource_AvailableSources", ""}).([]string)
}

func (ptr *QGeoSatelliteInfoSource) AvailableSources() []string {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.QGeoSatelliteInfoSource_AvailableSources", ""}).([]string)
}

func QGeoSatelliteInfoSource_CreateDefaultSource(parent core.QObject_ITF) *QGeoSatelliteInfoSource {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.QGeoSatelliteInfoSource_CreateDefaultSource", "", parent}).(*QGeoSatelliteInfoSource)
}

func (ptr *QGeoSatelliteInfoSource) CreateDefaultSource(parent core.QObject_ITF) *QGeoSatelliteInfoSource {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.QGeoSatelliteInfoSource_CreateDefaultSource", "", parent}).(*QGeoSatelliteInfoSource)
}

func QGeoSatelliteInfoSource_CreateSource(sourceName string, parent core.QObject_ITF) *QGeoSatelliteInfoSource {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.QGeoSatelliteInfoSource_CreateSource", "", sourceName, parent}).(*QGeoSatelliteInfoSource)
}

func (ptr *QGeoSatelliteInfoSource) CreateSource(sourceName string, parent core.QObject_ITF) *QGeoSatelliteInfoSource {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.QGeoSatelliteInfoSource_CreateSource", "", sourceName, parent}).(*QGeoSatelliteInfoSource)
}

func (ptr *QGeoSatelliteInfoSource) ConnectError(f func() QGeoSatelliteInfoSource__Error) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectError", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoSatelliteInfoSource) DisconnectError() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectError"})
}

func (ptr *QGeoSatelliteInfoSource) Error() QGeoSatelliteInfoSource__Error {

	return QGeoSatelliteInfoSource__Error(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error"}).(float64))
}

func (ptr *QGeoSatelliteInfoSource) ConnectError2(f func(satelliteError QGeoSatelliteInfoSource__Error)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectError2", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoSatelliteInfoSource) DisconnectError2() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectError2"})
}

func (ptr *QGeoSatelliteInfoSource) Error2(satelliteError QGeoSatelliteInfoSource__Error) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error2", satelliteError})
}

func (ptr *QGeoSatelliteInfoSource) ConnectMinimumUpdateInterval(f func() int) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMinimumUpdateInterval", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoSatelliteInfoSource) DisconnectMinimumUpdateInterval() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMinimumUpdateInterval"})
}

func (ptr *QGeoSatelliteInfoSource) MinimumUpdateInterval() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinimumUpdateInterval"}).(float64))
}

func (ptr *QGeoSatelliteInfoSource) ConnectRequestTimeout(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRequestTimeout", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoSatelliteInfoSource) DisconnectRequestTimeout() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRequestTimeout"})
}

func (ptr *QGeoSatelliteInfoSource) RequestTimeout() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RequestTimeout"})
}

func (ptr *QGeoSatelliteInfoSource) ConnectRequestUpdate(f func(timeout int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRequestUpdate", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoSatelliteInfoSource) DisconnectRequestUpdate() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRequestUpdate"})
}

func (ptr *QGeoSatelliteInfoSource) RequestUpdate(timeout int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RequestUpdate", timeout})
}

func (ptr *QGeoSatelliteInfoSource) ConnectSatellitesInUseUpdated(f func(satellites []*QGeoSatelliteInfo)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSatellitesInUseUpdated", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoSatelliteInfoSource) DisconnectSatellitesInUseUpdated() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSatellitesInUseUpdated"})
}

func (ptr *QGeoSatelliteInfoSource) SatellitesInUseUpdated(satellites []*QGeoSatelliteInfo) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SatellitesInUseUpdated", satellites})
}

func (ptr *QGeoSatelliteInfoSource) ConnectSatellitesInViewUpdated(f func(satellites []*QGeoSatelliteInfo)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSatellitesInViewUpdated", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoSatelliteInfoSource) DisconnectSatellitesInViewUpdated() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSatellitesInViewUpdated"})
}

func (ptr *QGeoSatelliteInfoSource) SatellitesInViewUpdated(satellites []*QGeoSatelliteInfo) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SatellitesInViewUpdated", satellites})
}

func (ptr *QGeoSatelliteInfoSource) ConnectSetUpdateInterval(f func(msec int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetUpdateInterval", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoSatelliteInfoSource) DisconnectSetUpdateInterval() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetUpdateInterval"})
}

func (ptr *QGeoSatelliteInfoSource) SetUpdateInterval(msec int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetUpdateInterval", msec})
}

func (ptr *QGeoSatelliteInfoSource) SetUpdateIntervalDefault(msec int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetUpdateIntervalDefault", msec})
}

func (ptr *QGeoSatelliteInfoSource) SourceName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SourceName"}).(string)
}

func (ptr *QGeoSatelliteInfoSource) ConnectStartUpdates(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStartUpdates", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoSatelliteInfoSource) DisconnectStartUpdates() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStartUpdates"})
}

func (ptr *QGeoSatelliteInfoSource) StartUpdates() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StartUpdates"})
}

func (ptr *QGeoSatelliteInfoSource) ConnectStopUpdates(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStopUpdates", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoSatelliteInfoSource) DisconnectStopUpdates() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStopUpdates"})
}

func (ptr *QGeoSatelliteInfoSource) StopUpdates() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StopUpdates"})
}

func (ptr *QGeoSatelliteInfoSource) UpdateInterval() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateInterval"}).(float64))
}

func (ptr *QGeoSatelliteInfoSource) ConnectDestroyQGeoSatelliteInfoSource(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQGeoSatelliteInfoSource", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QGeoSatelliteInfoSource) DisconnectDestroyQGeoSatelliteInfoSource() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQGeoSatelliteInfoSource"})
}

func (ptr *QGeoSatelliteInfoSource) DestroyQGeoSatelliteInfoSource() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQGeoSatelliteInfoSource"})
}

func (ptr *QGeoSatelliteInfoSource) DestroyQGeoSatelliteInfoSourceDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQGeoSatelliteInfoSourceDefault"})
}

func (ptr *QGeoSatelliteInfoSource) __satellitesInUseUpdated_satellites_atList(i int) *QGeoSatelliteInfo {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__satellitesInUseUpdated_satellites_atList", i}).(*QGeoSatelliteInfo)
}

func (ptr *QGeoSatelliteInfoSource) __satellitesInUseUpdated_satellites_setList(i QGeoSatelliteInfo_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__satellitesInUseUpdated_satellites_setList", i})
}

func (ptr *QGeoSatelliteInfoSource) __satellitesInUseUpdated_satellites_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__satellitesInUseUpdated_satellites_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoSatelliteInfoSource) __satellitesInViewUpdated_satellites_atList(i int) *QGeoSatelliteInfo {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__satellitesInViewUpdated_satellites_atList", i}).(*QGeoSatelliteInfo)
}

func (ptr *QGeoSatelliteInfoSource) __satellitesInViewUpdated_satellites_setList(i QGeoSatelliteInfo_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__satellitesInViewUpdated_satellites_setList", i})
}

func (ptr *QGeoSatelliteInfoSource) __satellitesInViewUpdated_satellites_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__satellitesInViewUpdated_satellites_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoSatelliteInfoSource) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QGeoSatelliteInfoSource) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QGeoSatelliteInfoSource) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoSatelliteInfoSource) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QGeoSatelliteInfoSource) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QGeoSatelliteInfoSource) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoSatelliteInfoSource) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QGeoSatelliteInfoSource) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QGeoSatelliteInfoSource) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QGeoSatelliteInfoSource) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QGeoSatelliteInfoSource) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QGeoSatelliteInfoSource) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QGeoSatelliteInfoSource) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QGeoSatelliteInfoSource) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QGeoSatelliteInfoSource) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QGeoSatelliteInfoSource) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QGeoSatelliteInfoSource) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QGeoSatelliteInfoSource) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QGeoSatelliteInfoSource) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QGeoSatelliteInfoSource) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QGeoSatelliteInfoSource) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QGeoShape struct {
	internal.Internal
}

type QGeoShape_ITF interface {
	QGeoShape_PTR() *QGeoShape
}

func (ptr *QGeoShape) QGeoShape_PTR() *QGeoShape {
	return ptr
}

func (ptr *QGeoShape) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QGeoShape) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQGeoShape(ptr QGeoShape_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoShape_PTR().Pointer()
	}
	return nil
}

func (n *QGeoShape) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQGeoShapeFromPointer(ptr unsafe.Pointer) (n *QGeoShape) {
	n = new(QGeoShape)
	n.InitFromInternal(uintptr(ptr), "positioning.QGeoShape")
	return
}

//go:generate stringer -type=QGeoShape__ShapeType
//QGeoShape::ShapeType
type QGeoShape__ShapeType int64

const (
	QGeoShape__UnknownType   QGeoShape__ShapeType = QGeoShape__ShapeType(0)
	QGeoShape__RectangleType QGeoShape__ShapeType = QGeoShape__ShapeType(1)
	QGeoShape__CircleType    QGeoShape__ShapeType = QGeoShape__ShapeType(2)
	QGeoShape__PathType      QGeoShape__ShapeType = QGeoShape__ShapeType(3)
	QGeoShape__PolygonType   QGeoShape__ShapeType = QGeoShape__ShapeType(4)
)

func NewQGeoShape() *QGeoShape {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.NewQGeoShape", ""}).(*QGeoShape)
}

func NewQGeoShape2(other QGeoShape_ITF) *QGeoShape {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.NewQGeoShape2", "", other}).(*QGeoShape)
}

func (ptr *QGeoShape) BoundingGeoRectangle() *QGeoRectangle {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BoundingGeoRectangle"}).(*QGeoRectangle)
}

func (ptr *QGeoShape) Center() *QGeoCoordinate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Center"}).(*QGeoCoordinate)
}

func (ptr *QGeoShape) Contains(coordinate QGeoCoordinate_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Contains", coordinate}).(bool)
}

func (ptr *QGeoShape) IsEmpty() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsEmpty"}).(bool)
}

func (ptr *QGeoShape) IsValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsValid"}).(bool)
}

func (ptr *QGeoShape) ToString() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToString"}).(string)
}

func (ptr *QGeoShape) Type() QGeoShape__ShapeType {

	return QGeoShape__ShapeType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(float64))
}

func (ptr *QGeoShape) DestroyQGeoShape() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQGeoShape"})
}

type QNmeaPositionInfoSource struct {
	QGeoPositionInfoSource
}

type QNmeaPositionInfoSource_ITF interface {
	QGeoPositionInfoSource_ITF
	QNmeaPositionInfoSource_PTR() *QNmeaPositionInfoSource
}

func (ptr *QNmeaPositionInfoSource) QNmeaPositionInfoSource_PTR() *QNmeaPositionInfoSource {
	return ptr
}

func (ptr *QNmeaPositionInfoSource) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoPositionInfoSource_PTR().Pointer()
	}
	return nil
}

func (ptr *QNmeaPositionInfoSource) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QGeoPositionInfoSource_PTR().SetPointer(p)
	}
}

func PointerFromQNmeaPositionInfoSource(ptr QNmeaPositionInfoSource_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNmeaPositionInfoSource_PTR().Pointer()
	}
	return nil
}

func (n *QNmeaPositionInfoSource) InitFromInternal(ptr uintptr, name string) {
	n.QGeoPositionInfoSource_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QNmeaPositionInfoSource) ClassNameInternalF() string {
	return n.QGeoPositionInfoSource_PTR().ClassNameInternalF()
}

func NewQNmeaPositionInfoSourceFromPointer(ptr unsafe.Pointer) (n *QNmeaPositionInfoSource) {
	n = new(QNmeaPositionInfoSource)
	n.InitFromInternal(uintptr(ptr), "positioning.QNmeaPositionInfoSource")
	return
}

//go:generate stringer -type=QNmeaPositionInfoSource__UpdateMode
//QNmeaPositionInfoSource::UpdateMode
type QNmeaPositionInfoSource__UpdateMode int64

const (
	QNmeaPositionInfoSource__RealTimeMode   QNmeaPositionInfoSource__UpdateMode = QNmeaPositionInfoSource__UpdateMode(1)
	QNmeaPositionInfoSource__SimulationMode QNmeaPositionInfoSource__UpdateMode = QNmeaPositionInfoSource__UpdateMode(2)
)

func NewQNmeaPositionInfoSource(updateMode QNmeaPositionInfoSource__UpdateMode, parent core.QObject_ITF) *QNmeaPositionInfoSource {

	return internal.CallLocalFunction([]interface{}{"", "", "positioning.NewQNmeaPositionInfoSource", "", updateMode, parent}).(*QNmeaPositionInfoSource)
}

func (ptr *QNmeaPositionInfoSource) Device() *core.QIODevice {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Device"}).(*core.QIODevice)
}

func (ptr *QNmeaPositionInfoSource) ConnectError(f func() QGeoPositionInfoSource__Error) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectError", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNmeaPositionInfoSource) DisconnectError() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectError"})
}

func (ptr *QNmeaPositionInfoSource) Error() QGeoPositionInfoSource__Error {

	return QGeoPositionInfoSource__Error(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error"}).(float64))
}

func (ptr *QNmeaPositionInfoSource) ErrorDefault() QGeoPositionInfoSource__Error {

	return QGeoPositionInfoSource__Error(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ErrorDefault"}).(float64))
}

func (ptr *QNmeaPositionInfoSource) ConnectLastKnownPosition(f func(fromSatellitePositioningMethodsOnly bool) *QGeoPositionInfo) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLastKnownPosition", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNmeaPositionInfoSource) DisconnectLastKnownPosition() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLastKnownPosition"})
}

func (ptr *QNmeaPositionInfoSource) LastKnownPosition(fromSatellitePositioningMethodsOnly bool) *QGeoPositionInfo {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LastKnownPosition", fromSatellitePositioningMethodsOnly}).(*QGeoPositionInfo)
}

func (ptr *QNmeaPositionInfoSource) LastKnownPositionDefault(fromSatellitePositioningMethodsOnly bool) *QGeoPositionInfo {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LastKnownPositionDefault", fromSatellitePositioningMethodsOnly}).(*QGeoPositionInfo)
}

func (ptr *QNmeaPositionInfoSource) ConnectMinimumUpdateInterval(f func() int) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMinimumUpdateInterval", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNmeaPositionInfoSource) DisconnectMinimumUpdateInterval() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMinimumUpdateInterval"})
}

func (ptr *QNmeaPositionInfoSource) MinimumUpdateInterval() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinimumUpdateInterval"}).(float64))
}

func (ptr *QNmeaPositionInfoSource) MinimumUpdateIntervalDefault() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinimumUpdateIntervalDefault"}).(float64))
}

func (ptr *QNmeaPositionInfoSource) ConnectParsePosInfoFromNmeaData(f func(data []byte, size int, posInfo *QGeoPositionInfo, hasFix *bool) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectParsePosInfoFromNmeaData", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNmeaPositionInfoSource) DisconnectParsePosInfoFromNmeaData() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectParsePosInfoFromNmeaData"})
}

func (ptr *QNmeaPositionInfoSource) ParsePosInfoFromNmeaData(data []byte, size int, posInfo QGeoPositionInfo_ITF, hasFix *bool) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ParsePosInfoFromNmeaData", data, size, posInfo, hasFix}).(bool)
}

func (ptr *QNmeaPositionInfoSource) ParsePosInfoFromNmeaDataDefault(data []byte, size int, posInfo QGeoPositionInfo_ITF, hasFix *bool) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ParsePosInfoFromNmeaDataDefault", data, size, posInfo, hasFix}).(bool)
}

func (ptr *QNmeaPositionInfoSource) ConnectRequestUpdate(f func(msec int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRequestUpdate", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNmeaPositionInfoSource) DisconnectRequestUpdate() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRequestUpdate"})
}

func (ptr *QNmeaPositionInfoSource) RequestUpdate(msec int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RequestUpdate", msec})
}

func (ptr *QNmeaPositionInfoSource) RequestUpdateDefault(msec int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RequestUpdateDefault", msec})
}

func (ptr *QNmeaPositionInfoSource) SetDevice(device core.QIODevice_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDevice", device})
}

func (ptr *QNmeaPositionInfoSource) SetUserEquivalentRangeError(uere float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetUserEquivalentRangeError", uere})
}

func (ptr *QNmeaPositionInfoSource) ConnectStartUpdates(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStartUpdates", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNmeaPositionInfoSource) DisconnectStartUpdates() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStartUpdates"})
}

func (ptr *QNmeaPositionInfoSource) StartUpdates() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StartUpdates"})
}

func (ptr *QNmeaPositionInfoSource) StartUpdatesDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StartUpdatesDefault"})
}

func (ptr *QNmeaPositionInfoSource) ConnectStopUpdates(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStopUpdates", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNmeaPositionInfoSource) DisconnectStopUpdates() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStopUpdates"})
}

func (ptr *QNmeaPositionInfoSource) StopUpdates() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StopUpdates"})
}

func (ptr *QNmeaPositionInfoSource) StopUpdatesDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StopUpdatesDefault"})
}

func (ptr *QNmeaPositionInfoSource) ConnectSupportedPositioningMethods(f func() QGeoPositionInfoSource__PositioningMethod) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSupportedPositioningMethods", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNmeaPositionInfoSource) DisconnectSupportedPositioningMethods() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSupportedPositioningMethods"})
}

func (ptr *QNmeaPositionInfoSource) SupportedPositioningMethods() QGeoPositionInfoSource__PositioningMethod {

	return QGeoPositionInfoSource__PositioningMethod(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SupportedPositioningMethods"}).(float64))
}

func (ptr *QNmeaPositionInfoSource) SupportedPositioningMethodsDefault() QGeoPositionInfoSource__PositioningMethod {

	return QGeoPositionInfoSource__PositioningMethod(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SupportedPositioningMethodsDefault"}).(float64))
}

func (ptr *QNmeaPositionInfoSource) UpdateMode() QNmeaPositionInfoSource__UpdateMode {

	return QNmeaPositionInfoSource__UpdateMode(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateMode"}).(float64))
}

func (ptr *QNmeaPositionInfoSource) UserEquivalentRangeError() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UserEquivalentRangeError"}).(float64)
}

func (ptr *QNmeaPositionInfoSource) ConnectDestroyQNmeaPositionInfoSource(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQNmeaPositionInfoSource", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNmeaPositionInfoSource) DisconnectDestroyQNmeaPositionInfoSource() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQNmeaPositionInfoSource"})
}

func (ptr *QNmeaPositionInfoSource) DestroyQNmeaPositionInfoSource() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQNmeaPositionInfoSource"})
}

func (ptr *QNmeaPositionInfoSource) DestroyQNmeaPositionInfoSourceDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQNmeaPositionInfoSourceDefault"})
}

func init() {
	internal.ConstructorTable["positioning.QGeoAddress"] = NewQGeoAddressFromPointer
	internal.ConstructorTable["positioning.QGeoAreaMonitorInfo"] = NewQGeoAreaMonitorInfoFromPointer
	internal.ConstructorTable["positioning.QGeoAreaMonitorSource"] = NewQGeoAreaMonitorSourceFromPointer
	internal.ConstructorTable["positioning.QGeoCircle"] = NewQGeoCircleFromPointer
	internal.ConstructorTable["positioning.QGeoCoordinate"] = NewQGeoCoordinateFromPointer
	internal.ConstructorTable["positioning.QGeoPath"] = NewQGeoPathFromPointer
	internal.ConstructorTable["positioning.QGeoPolygon"] = NewQGeoPolygonFromPointer
	internal.ConstructorTable["positioning.QGeoPositionInfo"] = NewQGeoPositionInfoFromPointer
	internal.ConstructorTable["positioning.QGeoPositionInfoSource"] = NewQGeoPositionInfoSourceFromPointer
	internal.ConstructorTable["positioning.QGeoPositionInfoSourceFactory"] = NewQGeoPositionInfoSourceFactoryFromPointer
	internal.ConstructorTable["positioning.QGeoRectangle"] = NewQGeoRectangleFromPointer
	internal.ConstructorTable["positioning.QGeoSatelliteInfo"] = NewQGeoSatelliteInfoFromPointer
	internal.ConstructorTable["positioning.QGeoSatelliteInfoSource"] = NewQGeoSatelliteInfoSourceFromPointer
	internal.ConstructorTable["positioning.QGeoShape"] = NewQGeoShapeFromPointer
	internal.ConstructorTable["positioning.QNmeaPositionInfoSource"] = NewQNmeaPositionInfoSourceFromPointer
}
