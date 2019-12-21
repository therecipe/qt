// +build !minimal

package positioning

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "positioning.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

func cGoFreePacked(ptr unsafe.Pointer) { core.NewQByteArrayFromPointer(ptr).DestroyQByteArray() }
func cGoUnpackString(s C.struct_QtPositioning_PackedString) string {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtPositioning_PackedString) []byte {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		gs := C.GoString(s.data)
		return *(*[]byte)(unsafe.Pointer(&gs))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}
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

func NewLocationSingletonFromPointer(ptr unsafe.Pointer) (n *LocationSingleton) {
	n = new(LocationSingleton)
	n.SetPointer(ptr)
	return
}

type QGeoAddress struct {
	ptr unsafe.Pointer
}

type QGeoAddress_ITF interface {
	QGeoAddress_PTR() *QGeoAddress
}

func (ptr *QGeoAddress) QGeoAddress_PTR() *QGeoAddress {
	return ptr
}

func (ptr *QGeoAddress) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QGeoAddress) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQGeoAddress(ptr QGeoAddress_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoAddress_PTR().Pointer()
	}
	return nil
}

func NewQGeoAddressFromPointer(ptr unsafe.Pointer) (n *QGeoAddress) {
	n = new(QGeoAddress)
	n.SetPointer(ptr)
	return
}
func NewQGeoAddress() *QGeoAddress {
	tmpValue := NewQGeoAddressFromPointer(C.QGeoAddress_NewQGeoAddress())
	qt.SetFinalizer(tmpValue, (*QGeoAddress).DestroyQGeoAddress)
	return tmpValue
}

func NewQGeoAddress2(other QGeoAddress_ITF) *QGeoAddress {
	tmpValue := NewQGeoAddressFromPointer(C.QGeoAddress_NewQGeoAddress2(PointerFromQGeoAddress(other)))
	qt.SetFinalizer(tmpValue, (*QGeoAddress).DestroyQGeoAddress)
	return tmpValue
}

func (ptr *QGeoAddress) City() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoAddress_City(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoAddress) Clear() {
	if ptr.Pointer() != nil {
		C.QGeoAddress_Clear(ptr.Pointer())
	}
}

func (ptr *QGeoAddress) Country() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoAddress_Country(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoAddress) CountryCode() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoAddress_CountryCode(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoAddress) County() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoAddress_County(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoAddress) District() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoAddress_District(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoAddress) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGeoAddress_IsEmpty(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGeoAddress) IsTextGenerated() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGeoAddress_IsTextGenerated(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGeoAddress) PostalCode() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoAddress_PostalCode(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoAddress) SetCity(city string) {
	if ptr.Pointer() != nil {
		var cityC *C.char
		if city != "" {
			cityC = C.CString(city)
			defer C.free(unsafe.Pointer(cityC))
		}
		C.QGeoAddress_SetCity(ptr.Pointer(), C.struct_QtPositioning_PackedString{data: cityC, len: C.longlong(len(city))})
	}
}

func (ptr *QGeoAddress) SetCountry(country string) {
	if ptr.Pointer() != nil {
		var countryC *C.char
		if country != "" {
			countryC = C.CString(country)
			defer C.free(unsafe.Pointer(countryC))
		}
		C.QGeoAddress_SetCountry(ptr.Pointer(), C.struct_QtPositioning_PackedString{data: countryC, len: C.longlong(len(country))})
	}
}

func (ptr *QGeoAddress) SetCountryCode(countryCode string) {
	if ptr.Pointer() != nil {
		var countryCodeC *C.char
		if countryCode != "" {
			countryCodeC = C.CString(countryCode)
			defer C.free(unsafe.Pointer(countryCodeC))
		}
		C.QGeoAddress_SetCountryCode(ptr.Pointer(), C.struct_QtPositioning_PackedString{data: countryCodeC, len: C.longlong(len(countryCode))})
	}
}

func (ptr *QGeoAddress) SetCounty(county string) {
	if ptr.Pointer() != nil {
		var countyC *C.char
		if county != "" {
			countyC = C.CString(county)
			defer C.free(unsafe.Pointer(countyC))
		}
		C.QGeoAddress_SetCounty(ptr.Pointer(), C.struct_QtPositioning_PackedString{data: countyC, len: C.longlong(len(county))})
	}
}

func (ptr *QGeoAddress) SetDistrict(district string) {
	if ptr.Pointer() != nil {
		var districtC *C.char
		if district != "" {
			districtC = C.CString(district)
			defer C.free(unsafe.Pointer(districtC))
		}
		C.QGeoAddress_SetDistrict(ptr.Pointer(), C.struct_QtPositioning_PackedString{data: districtC, len: C.longlong(len(district))})
	}
}

func (ptr *QGeoAddress) SetPostalCode(postalCode string) {
	if ptr.Pointer() != nil {
		var postalCodeC *C.char
		if postalCode != "" {
			postalCodeC = C.CString(postalCode)
			defer C.free(unsafe.Pointer(postalCodeC))
		}
		C.QGeoAddress_SetPostalCode(ptr.Pointer(), C.struct_QtPositioning_PackedString{data: postalCodeC, len: C.longlong(len(postalCode))})
	}
}

func (ptr *QGeoAddress) SetState(state string) {
	if ptr.Pointer() != nil {
		var stateC *C.char
		if state != "" {
			stateC = C.CString(state)
			defer C.free(unsafe.Pointer(stateC))
		}
		C.QGeoAddress_SetState(ptr.Pointer(), C.struct_QtPositioning_PackedString{data: stateC, len: C.longlong(len(state))})
	}
}

func (ptr *QGeoAddress) SetStreet(street string) {
	if ptr.Pointer() != nil {
		var streetC *C.char
		if street != "" {
			streetC = C.CString(street)
			defer C.free(unsafe.Pointer(streetC))
		}
		C.QGeoAddress_SetStreet(ptr.Pointer(), C.struct_QtPositioning_PackedString{data: streetC, len: C.longlong(len(street))})
	}
}

func (ptr *QGeoAddress) SetText(text string) {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.QGeoAddress_SetText(ptr.Pointer(), C.struct_QtPositioning_PackedString{data: textC, len: C.longlong(len(text))})
	}
}

func (ptr *QGeoAddress) State() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoAddress_State(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoAddress) Street() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoAddress_Street(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoAddress) Text() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoAddress_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoAddress) DestroyQGeoAddress() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGeoAddress_DestroyQGeoAddress(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QGeoAreaMonitorInfo struct {
	ptr unsafe.Pointer
}

type QGeoAreaMonitorInfo_ITF interface {
	QGeoAreaMonitorInfo_PTR() *QGeoAreaMonitorInfo
}

func (ptr *QGeoAreaMonitorInfo) QGeoAreaMonitorInfo_PTR() *QGeoAreaMonitorInfo {
	return ptr
}

func (ptr *QGeoAreaMonitorInfo) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QGeoAreaMonitorInfo) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQGeoAreaMonitorInfo(ptr QGeoAreaMonitorInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoAreaMonitorInfo_PTR().Pointer()
	}
	return nil
}

func NewQGeoAreaMonitorInfoFromPointer(ptr unsafe.Pointer) (n *QGeoAreaMonitorInfo) {
	n = new(QGeoAreaMonitorInfo)
	n.SetPointer(ptr)
	return
}
func NewQGeoAreaMonitorInfo(name string) *QGeoAreaMonitorInfo {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := NewQGeoAreaMonitorInfoFromPointer(C.QGeoAreaMonitorInfo_NewQGeoAreaMonitorInfo(C.struct_QtPositioning_PackedString{data: nameC, len: C.longlong(len(name))}))
	qt.SetFinalizer(tmpValue, (*QGeoAreaMonitorInfo).DestroyQGeoAreaMonitorInfo)
	return tmpValue
}

func NewQGeoAreaMonitorInfo2(other QGeoAreaMonitorInfo_ITF) *QGeoAreaMonitorInfo {
	tmpValue := NewQGeoAreaMonitorInfoFromPointer(C.QGeoAreaMonitorInfo_NewQGeoAreaMonitorInfo2(PointerFromQGeoAreaMonitorInfo(other)))
	qt.SetFinalizer(tmpValue, (*QGeoAreaMonitorInfo).DestroyQGeoAreaMonitorInfo)
	return tmpValue
}

func (ptr *QGeoAreaMonitorInfo) Area() *QGeoShape {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoShapeFromPointer(C.QGeoAreaMonitorInfo_Area(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QGeoShape).DestroyQGeoShape)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoAreaMonitorInfo) Expiration() *core.QDateTime {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQDateTimeFromPointer(C.QGeoAreaMonitorInfo_Expiration(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QDateTime).DestroyQDateTime)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoAreaMonitorInfo) Identifier() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoAreaMonitorInfo_Identifier(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoAreaMonitorInfo) IsPersistent() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGeoAreaMonitorInfo_IsPersistent(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGeoAreaMonitorInfo) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGeoAreaMonitorInfo_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGeoAreaMonitorInfo) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoAreaMonitorInfo_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoAreaMonitorInfo) NotificationParameters() map[string]*core.QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtPositioning_PackedList) map[string]*core.QVariant {
			out := make(map[string]*core.QVariant, int(l.len))
			tmpList := NewQGeoAreaMonitorInfoFromPointer(l.data)
			for i, v := range tmpList.__notificationParameters_keyList() {
				out[v] = tmpList.__notificationParameters_atList(v, i)
			}
			return out
		}(C.QGeoAreaMonitorInfo_NotificationParameters(ptr.Pointer()))
	}
	return make(map[string]*core.QVariant, 0)
}

func (ptr *QGeoAreaMonitorInfo) SetArea(newShape QGeoShape_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorInfo_SetArea(ptr.Pointer(), PointerFromQGeoShape(newShape))
	}
}

func (ptr *QGeoAreaMonitorInfo) SetExpiration(expiry core.QDateTime_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorInfo_SetExpiration(ptr.Pointer(), core.PointerFromQDateTime(expiry))
	}
}

func (ptr *QGeoAreaMonitorInfo) SetName(name string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QGeoAreaMonitorInfo_SetName(ptr.Pointer(), C.struct_QtPositioning_PackedString{data: nameC, len: C.longlong(len(name))})
	}
}

func (ptr *QGeoAreaMonitorInfo) SetNotificationParameters(parameters map[string]*core.QVariant) {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorInfo_SetNotificationParameters(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQGeoAreaMonitorInfoFromPointer(NewQGeoAreaMonitorInfoFromPointer(nil).__setNotificationParameters_parameters_newList())
			for k, v := range parameters {
				tmpList.__setNotificationParameters_parameters_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QGeoAreaMonitorInfo) SetPersistent(isPersistent bool) {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorInfo_SetPersistent(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(isPersistent))))
	}
}

func (ptr *QGeoAreaMonitorInfo) DestroyQGeoAreaMonitorInfo() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGeoAreaMonitorInfo_DestroyQGeoAreaMonitorInfo(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGeoAreaMonitorInfo) __notificationParameters_atList(v string, i int) *core.QVariant {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		tmpValue := core.NewQVariantFromPointer(C.QGeoAreaMonitorInfo___notificationParameters_atList(ptr.Pointer(), C.struct_QtPositioning_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoAreaMonitorInfo) __notificationParameters_setList(key string, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QGeoAreaMonitorInfo___notificationParameters_setList(ptr.Pointer(), C.struct_QtPositioning_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQVariant(i))
	}
}

func (ptr *QGeoAreaMonitorInfo) __notificationParameters_newList() unsafe.Pointer {
	return C.QGeoAreaMonitorInfo___notificationParameters_newList(ptr.Pointer())
}

func (ptr *QGeoAreaMonitorInfo) __notificationParameters_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtPositioning_PackedList) []string {
			out := make([]string, int(l.len))
			tmpList := NewQGeoAreaMonitorInfoFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____notificationParameters_keyList_atList(i)
			}
			return out
		}(C.QGeoAreaMonitorInfo___notificationParameters_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QGeoAreaMonitorInfo) __setNotificationParameters_parameters_atList(v string, i int) *core.QVariant {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		tmpValue := core.NewQVariantFromPointer(C.QGeoAreaMonitorInfo___setNotificationParameters_parameters_atList(ptr.Pointer(), C.struct_QtPositioning_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoAreaMonitorInfo) __setNotificationParameters_parameters_setList(key string, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QGeoAreaMonitorInfo___setNotificationParameters_parameters_setList(ptr.Pointer(), C.struct_QtPositioning_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQVariant(i))
	}
}

func (ptr *QGeoAreaMonitorInfo) __setNotificationParameters_parameters_newList() unsafe.Pointer {
	return C.QGeoAreaMonitorInfo___setNotificationParameters_parameters_newList(ptr.Pointer())
}

func (ptr *QGeoAreaMonitorInfo) __setNotificationParameters_parameters_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtPositioning_PackedList) []string {
			out := make([]string, int(l.len))
			tmpList := NewQGeoAreaMonitorInfoFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____setNotificationParameters_parameters_keyList_atList(i)
			}
			return out
		}(C.QGeoAreaMonitorInfo___setNotificationParameters_parameters_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QGeoAreaMonitorInfo) ____notificationParameters_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoAreaMonitorInfo_____notificationParameters_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QGeoAreaMonitorInfo) ____notificationParameters_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QGeoAreaMonitorInfo_____notificationParameters_keyList_setList(ptr.Pointer(), C.struct_QtPositioning_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QGeoAreaMonitorInfo) ____notificationParameters_keyList_newList() unsafe.Pointer {
	return C.QGeoAreaMonitorInfo_____notificationParameters_keyList_newList(ptr.Pointer())
}

func (ptr *QGeoAreaMonitorInfo) ____setNotificationParameters_parameters_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoAreaMonitorInfo_____setNotificationParameters_parameters_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QGeoAreaMonitorInfo) ____setNotificationParameters_parameters_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QGeoAreaMonitorInfo_____setNotificationParameters_parameters_keyList_setList(ptr.Pointer(), C.struct_QtPositioning_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QGeoAreaMonitorInfo) ____setNotificationParameters_parameters_keyList_newList() unsafe.Pointer {
	return C.QGeoAreaMonitorInfo_____setNotificationParameters_parameters_keyList_newList(ptr.Pointer())
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

func NewQGeoAreaMonitorSourceFromPointer(ptr unsafe.Pointer) (n *QGeoAreaMonitorSource) {
	n = new(QGeoAreaMonitorSource)
	n.SetPointer(ptr)
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
	tmpValue := NewQGeoAreaMonitorSourceFromPointer(C.QGeoAreaMonitorSource_NewQGeoAreaMonitorSource(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQGeoAreaMonitorSource_ActiveMonitors
func callbackQGeoAreaMonitorSource_ActiveMonitors(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "activeMonitors"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewQGeoAreaMonitorSourceFromPointer(NewQGeoAreaMonitorSourceFromPointer(nil).__activeMonitors_newList())
			for _, v := range (*(*func() []*QGeoAreaMonitorInfo)(signal))() {
				tmpList.__activeMonitors_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewQGeoAreaMonitorSourceFromPointer(NewQGeoAreaMonitorSourceFromPointer(nil).__activeMonitors_newList())
		for _, v := range make([]*QGeoAreaMonitorInfo, 0) {
			tmpList.__activeMonitors_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *QGeoAreaMonitorSource) ConnectActiveMonitors(f func() []*QGeoAreaMonitorInfo) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "activeMonitors"); signal != nil {
			f := func() []*QGeoAreaMonitorInfo {
				(*(*func() []*QGeoAreaMonitorInfo)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "activeMonitors", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "activeMonitors", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoAreaMonitorSource) DisconnectActiveMonitors() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "activeMonitors")
	}
}

func (ptr *QGeoAreaMonitorSource) ActiveMonitors() []*QGeoAreaMonitorInfo {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtPositioning_PackedList) []*QGeoAreaMonitorInfo {
			out := make([]*QGeoAreaMonitorInfo, int(l.len))
			tmpList := NewQGeoAreaMonitorSourceFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__activeMonitors_atList(i)
			}
			return out
		}(C.QGeoAreaMonitorSource_ActiveMonitors(ptr.Pointer()))
	}
	return make([]*QGeoAreaMonitorInfo, 0)
}

//export callbackQGeoAreaMonitorSource_ActiveMonitors2
func callbackQGeoAreaMonitorSource_ActiveMonitors2(ptr unsafe.Pointer, lookupArea unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "activeMonitors2"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewQGeoAreaMonitorSourceFromPointer(NewQGeoAreaMonitorSourceFromPointer(nil).__activeMonitors_newList2())
			for _, v := range (*(*func(*QGeoShape) []*QGeoAreaMonitorInfo)(signal))(NewQGeoShapeFromPointer(lookupArea)) {
				tmpList.__activeMonitors_setList2(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewQGeoAreaMonitorSourceFromPointer(NewQGeoAreaMonitorSourceFromPointer(nil).__activeMonitors_newList2())
		for _, v := range make([]*QGeoAreaMonitorInfo, 0) {
			tmpList.__activeMonitors_setList2(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *QGeoAreaMonitorSource) ConnectActiveMonitors2(f func(lookupArea *QGeoShape) []*QGeoAreaMonitorInfo) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "activeMonitors2"); signal != nil {
			f := func(lookupArea *QGeoShape) []*QGeoAreaMonitorInfo {
				(*(*func(*QGeoShape) []*QGeoAreaMonitorInfo)(signal))(lookupArea)
				return f(lookupArea)
			}
			qt.ConnectSignal(ptr.Pointer(), "activeMonitors2", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "activeMonitors2", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoAreaMonitorSource) DisconnectActiveMonitors2() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "activeMonitors2")
	}
}

func (ptr *QGeoAreaMonitorSource) ActiveMonitors2(lookupArea QGeoShape_ITF) []*QGeoAreaMonitorInfo {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtPositioning_PackedList) []*QGeoAreaMonitorInfo {
			out := make([]*QGeoAreaMonitorInfo, int(l.len))
			tmpList := NewQGeoAreaMonitorSourceFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__activeMonitors_atList2(i)
			}
			return out
		}(C.QGeoAreaMonitorSource_ActiveMonitors2(ptr.Pointer(), PointerFromQGeoShape(lookupArea)))
	}
	return make([]*QGeoAreaMonitorInfo, 0)
}

//export callbackQGeoAreaMonitorSource_AreaEntered
func callbackQGeoAreaMonitorSource_AreaEntered(ptr unsafe.Pointer, monitor unsafe.Pointer, update unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "areaEntered"); signal != nil {
		(*(*func(*QGeoAreaMonitorInfo, *QGeoPositionInfo))(signal))(NewQGeoAreaMonitorInfoFromPointer(monitor), NewQGeoPositionInfoFromPointer(update))
	}

}

func (ptr *QGeoAreaMonitorSource) ConnectAreaEntered(f func(monitor *QGeoAreaMonitorInfo, update *QGeoPositionInfo)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "areaEntered") {
			C.QGeoAreaMonitorSource_ConnectAreaEntered(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "areaEntered")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "areaEntered"); signal != nil {
			f := func(monitor *QGeoAreaMonitorInfo, update *QGeoPositionInfo) {
				(*(*func(*QGeoAreaMonitorInfo, *QGeoPositionInfo))(signal))(monitor, update)
				f(monitor, update)
			}
			qt.ConnectSignal(ptr.Pointer(), "areaEntered", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "areaEntered", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoAreaMonitorSource) DisconnectAreaEntered() {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_DisconnectAreaEntered(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "areaEntered")
	}
}

func (ptr *QGeoAreaMonitorSource) AreaEntered(monitor QGeoAreaMonitorInfo_ITF, update QGeoPositionInfo_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_AreaEntered(ptr.Pointer(), PointerFromQGeoAreaMonitorInfo(monitor), PointerFromQGeoPositionInfo(update))
	}
}

//export callbackQGeoAreaMonitorSource_AreaExited
func callbackQGeoAreaMonitorSource_AreaExited(ptr unsafe.Pointer, monitor unsafe.Pointer, update unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "areaExited"); signal != nil {
		(*(*func(*QGeoAreaMonitorInfo, *QGeoPositionInfo))(signal))(NewQGeoAreaMonitorInfoFromPointer(monitor), NewQGeoPositionInfoFromPointer(update))
	}

}

func (ptr *QGeoAreaMonitorSource) ConnectAreaExited(f func(monitor *QGeoAreaMonitorInfo, update *QGeoPositionInfo)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "areaExited") {
			C.QGeoAreaMonitorSource_ConnectAreaExited(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "areaExited")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "areaExited"); signal != nil {
			f := func(monitor *QGeoAreaMonitorInfo, update *QGeoPositionInfo) {
				(*(*func(*QGeoAreaMonitorInfo, *QGeoPositionInfo))(signal))(monitor, update)
				f(monitor, update)
			}
			qt.ConnectSignal(ptr.Pointer(), "areaExited", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "areaExited", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoAreaMonitorSource) DisconnectAreaExited() {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_DisconnectAreaExited(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "areaExited")
	}
}

func (ptr *QGeoAreaMonitorSource) AreaExited(monitor QGeoAreaMonitorInfo_ITF, update QGeoPositionInfo_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_AreaExited(ptr.Pointer(), PointerFromQGeoAreaMonitorInfo(monitor), PointerFromQGeoPositionInfo(update))
	}
}

func QGeoAreaMonitorSource_AvailableSources() []string {
	return unpackStringList(cGoUnpackString(C.QGeoAreaMonitorSource_QGeoAreaMonitorSource_AvailableSources()))
}

func (ptr *QGeoAreaMonitorSource) AvailableSources() []string {
	return unpackStringList(cGoUnpackString(C.QGeoAreaMonitorSource_QGeoAreaMonitorSource_AvailableSources()))
}

func QGeoAreaMonitorSource_CreateDefaultSource(parent core.QObject_ITF) *QGeoAreaMonitorSource {
	tmpValue := NewQGeoAreaMonitorSourceFromPointer(C.QGeoAreaMonitorSource_QGeoAreaMonitorSource_CreateDefaultSource(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QGeoAreaMonitorSource) CreateDefaultSource(parent core.QObject_ITF) *QGeoAreaMonitorSource {
	tmpValue := NewQGeoAreaMonitorSourceFromPointer(C.QGeoAreaMonitorSource_QGeoAreaMonitorSource_CreateDefaultSource(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QGeoAreaMonitorSource_CreateSource(sourceName string, parent core.QObject_ITF) *QGeoAreaMonitorSource {
	var sourceNameC *C.char
	if sourceName != "" {
		sourceNameC = C.CString(sourceName)
		defer C.free(unsafe.Pointer(sourceNameC))
	}
	tmpValue := NewQGeoAreaMonitorSourceFromPointer(C.QGeoAreaMonitorSource_QGeoAreaMonitorSource_CreateSource(C.struct_QtPositioning_PackedString{data: sourceNameC, len: C.longlong(len(sourceName))}, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QGeoAreaMonitorSource) CreateSource(sourceName string, parent core.QObject_ITF) *QGeoAreaMonitorSource {
	var sourceNameC *C.char
	if sourceName != "" {
		sourceNameC = C.CString(sourceName)
		defer C.free(unsafe.Pointer(sourceNameC))
	}
	tmpValue := NewQGeoAreaMonitorSourceFromPointer(C.QGeoAreaMonitorSource_QGeoAreaMonitorSource_CreateSource(C.struct_QtPositioning_PackedString{data: sourceNameC, len: C.longlong(len(sourceName))}, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQGeoAreaMonitorSource_Error
func callbackQGeoAreaMonitorSource_Error(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "error"); signal != nil {
		return C.longlong((*(*func() QGeoAreaMonitorSource__Error)(signal))())
	}

	return C.longlong(0)
}

func (ptr *QGeoAreaMonitorSource) ConnectError(f func() QGeoAreaMonitorSource__Error) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "error"); signal != nil {
			f := func() QGeoAreaMonitorSource__Error {
				(*(*func() QGeoAreaMonitorSource__Error)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "error", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "error", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoAreaMonitorSource) DisconnectError() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "error")
	}
}

func (ptr *QGeoAreaMonitorSource) Error() QGeoAreaMonitorSource__Error {
	if ptr.Pointer() != nil {
		return QGeoAreaMonitorSource__Error(C.QGeoAreaMonitorSource_Error(ptr.Pointer()))
	}
	return 0
}

//export callbackQGeoAreaMonitorSource_Error2
func callbackQGeoAreaMonitorSource_Error2(ptr unsafe.Pointer, areaMonitoringError C.longlong) {
	if signal := qt.GetSignal(ptr, "error2"); signal != nil {
		(*(*func(QGeoAreaMonitorSource__Error))(signal))(QGeoAreaMonitorSource__Error(areaMonitoringError))
	}

}

func (ptr *QGeoAreaMonitorSource) ConnectError2(f func(areaMonitoringError QGeoAreaMonitorSource__Error)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "error2") {
			C.QGeoAreaMonitorSource_ConnectError2(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "error")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "error2"); signal != nil {
			f := func(areaMonitoringError QGeoAreaMonitorSource__Error) {
				(*(*func(QGeoAreaMonitorSource__Error))(signal))(areaMonitoringError)
				f(areaMonitoringError)
			}
			qt.ConnectSignal(ptr.Pointer(), "error2", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "error2", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoAreaMonitorSource) DisconnectError2() {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "error2")
	}
}

func (ptr *QGeoAreaMonitorSource) Error2(areaMonitoringError QGeoAreaMonitorSource__Error) {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_Error2(ptr.Pointer(), C.longlong(areaMonitoringError))
	}
}

//export callbackQGeoAreaMonitorSource_MonitorExpired
func callbackQGeoAreaMonitorSource_MonitorExpired(ptr unsafe.Pointer, monitor unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "monitorExpired"); signal != nil {
		(*(*func(*QGeoAreaMonitorInfo))(signal))(NewQGeoAreaMonitorInfoFromPointer(monitor))
	}

}

func (ptr *QGeoAreaMonitorSource) ConnectMonitorExpired(f func(monitor *QGeoAreaMonitorInfo)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "monitorExpired") {
			C.QGeoAreaMonitorSource_ConnectMonitorExpired(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "monitorExpired")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "monitorExpired"); signal != nil {
			f := func(monitor *QGeoAreaMonitorInfo) {
				(*(*func(*QGeoAreaMonitorInfo))(signal))(monitor)
				f(monitor)
			}
			qt.ConnectSignal(ptr.Pointer(), "monitorExpired", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "monitorExpired", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoAreaMonitorSource) DisconnectMonitorExpired() {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_DisconnectMonitorExpired(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "monitorExpired")
	}
}

func (ptr *QGeoAreaMonitorSource) MonitorExpired(monitor QGeoAreaMonitorInfo_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_MonitorExpired(ptr.Pointer(), PointerFromQGeoAreaMonitorInfo(monitor))
	}
}

//export callbackQGeoAreaMonitorSource_PositionInfoSource
func callbackQGeoAreaMonitorSource_PositionInfoSource(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "positionInfoSource"); signal != nil {
		return PointerFromQGeoPositionInfoSource((*(*func() *QGeoPositionInfoSource)(signal))())
	}

	return PointerFromQGeoPositionInfoSource(NewQGeoAreaMonitorSourceFromPointer(ptr).PositionInfoSourceDefault())
}

func (ptr *QGeoAreaMonitorSource) ConnectPositionInfoSource(f func() *QGeoPositionInfoSource) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "positionInfoSource"); signal != nil {
			f := func() *QGeoPositionInfoSource {
				(*(*func() *QGeoPositionInfoSource)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "positionInfoSource", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "positionInfoSource", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoAreaMonitorSource) DisconnectPositionInfoSource() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "positionInfoSource")
	}
}

func (ptr *QGeoAreaMonitorSource) PositionInfoSource() *QGeoPositionInfoSource {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoPositionInfoSourceFromPointer(C.QGeoAreaMonitorSource_PositionInfoSource(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGeoAreaMonitorSource) PositionInfoSourceDefault() *QGeoPositionInfoSource {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoPositionInfoSourceFromPointer(C.QGeoAreaMonitorSource_PositionInfoSourceDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQGeoAreaMonitorSource_RequestUpdate
func callbackQGeoAreaMonitorSource_RequestUpdate(ptr unsafe.Pointer, monitor unsafe.Pointer, sign C.struct_QtPositioning_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "requestUpdate"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*QGeoAreaMonitorInfo, string) bool)(signal))(NewQGeoAreaMonitorInfoFromPointer(monitor), cGoUnpackString(sign)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QGeoAreaMonitorSource) ConnectRequestUpdate(f func(monitor *QGeoAreaMonitorInfo, sign string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "requestUpdate"); signal != nil {
			f := func(monitor *QGeoAreaMonitorInfo, sign string) bool {
				(*(*func(*QGeoAreaMonitorInfo, string) bool)(signal))(monitor, sign)
				return f(monitor, sign)
			}
			qt.ConnectSignal(ptr.Pointer(), "requestUpdate", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "requestUpdate", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoAreaMonitorSource) DisconnectRequestUpdate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "requestUpdate")
	}
}

func (ptr *QGeoAreaMonitorSource) RequestUpdate(monitor QGeoAreaMonitorInfo_ITF, sign string) bool {
	if ptr.Pointer() != nil {
		var signC *C.char
		if sign != "" {
			signC = C.CString(sign)
			defer C.free(unsafe.Pointer(signC))
		}
		return int8(C.QGeoAreaMonitorSource_RequestUpdate(ptr.Pointer(), PointerFromQGeoAreaMonitorInfo(monitor), signC)) != 0
	}
	return false
}

//export callbackQGeoAreaMonitorSource_SetPositionInfoSource
func callbackQGeoAreaMonitorSource_SetPositionInfoSource(ptr unsafe.Pointer, newSource unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setPositionInfoSource"); signal != nil {
		(*(*func(*QGeoPositionInfoSource))(signal))(NewQGeoPositionInfoSourceFromPointer(newSource))
	} else {
		NewQGeoAreaMonitorSourceFromPointer(ptr).SetPositionInfoSourceDefault(NewQGeoPositionInfoSourceFromPointer(newSource))
	}
}

func (ptr *QGeoAreaMonitorSource) ConnectSetPositionInfoSource(f func(newSource *QGeoPositionInfoSource)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setPositionInfoSource"); signal != nil {
			f := func(newSource *QGeoPositionInfoSource) {
				(*(*func(*QGeoPositionInfoSource))(signal))(newSource)
				f(newSource)
			}
			qt.ConnectSignal(ptr.Pointer(), "setPositionInfoSource", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setPositionInfoSource", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoAreaMonitorSource) DisconnectSetPositionInfoSource() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setPositionInfoSource")
	}
}

func (ptr *QGeoAreaMonitorSource) SetPositionInfoSource(newSource QGeoPositionInfoSource_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_SetPositionInfoSource(ptr.Pointer(), PointerFromQGeoPositionInfoSource(newSource))
	}
}

func (ptr *QGeoAreaMonitorSource) SetPositionInfoSourceDefault(newSource QGeoPositionInfoSource_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_SetPositionInfoSourceDefault(ptr.Pointer(), PointerFromQGeoPositionInfoSource(newSource))
	}
}

func (ptr *QGeoAreaMonitorSource) SourceName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoAreaMonitorSource_SourceName(ptr.Pointer()))
	}
	return ""
}

//export callbackQGeoAreaMonitorSource_StartMonitoring
func callbackQGeoAreaMonitorSource_StartMonitoring(ptr unsafe.Pointer, monitor unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "startMonitoring"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*QGeoAreaMonitorInfo) bool)(signal))(NewQGeoAreaMonitorInfoFromPointer(monitor)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QGeoAreaMonitorSource) ConnectStartMonitoring(f func(monitor *QGeoAreaMonitorInfo) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "startMonitoring"); signal != nil {
			f := func(monitor *QGeoAreaMonitorInfo) bool {
				(*(*func(*QGeoAreaMonitorInfo) bool)(signal))(monitor)
				return f(monitor)
			}
			qt.ConnectSignal(ptr.Pointer(), "startMonitoring", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "startMonitoring", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoAreaMonitorSource) DisconnectStartMonitoring() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "startMonitoring")
	}
}

func (ptr *QGeoAreaMonitorSource) StartMonitoring(monitor QGeoAreaMonitorInfo_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGeoAreaMonitorSource_StartMonitoring(ptr.Pointer(), PointerFromQGeoAreaMonitorInfo(monitor))) != 0
	}
	return false
}

//export callbackQGeoAreaMonitorSource_StopMonitoring
func callbackQGeoAreaMonitorSource_StopMonitoring(ptr unsafe.Pointer, monitor unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "stopMonitoring"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*QGeoAreaMonitorInfo) bool)(signal))(NewQGeoAreaMonitorInfoFromPointer(monitor)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QGeoAreaMonitorSource) ConnectStopMonitoring(f func(monitor *QGeoAreaMonitorInfo) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "stopMonitoring"); signal != nil {
			f := func(monitor *QGeoAreaMonitorInfo) bool {
				(*(*func(*QGeoAreaMonitorInfo) bool)(signal))(monitor)
				return f(monitor)
			}
			qt.ConnectSignal(ptr.Pointer(), "stopMonitoring", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stopMonitoring", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoAreaMonitorSource) DisconnectStopMonitoring() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "stopMonitoring")
	}
}

func (ptr *QGeoAreaMonitorSource) StopMonitoring(monitor QGeoAreaMonitorInfo_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGeoAreaMonitorSource_StopMonitoring(ptr.Pointer(), PointerFromQGeoAreaMonitorInfo(monitor))) != 0
	}
	return false
}

//export callbackQGeoAreaMonitorSource_SupportedAreaMonitorFeatures
func callbackQGeoAreaMonitorSource_SupportedAreaMonitorFeatures(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedAreaMonitorFeatures"); signal != nil {
		return C.longlong((*(*func() QGeoAreaMonitorSource__AreaMonitorFeature)(signal))())
	}

	return C.longlong(0)
}

func (ptr *QGeoAreaMonitorSource) ConnectSupportedAreaMonitorFeatures(f func() QGeoAreaMonitorSource__AreaMonitorFeature) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "supportedAreaMonitorFeatures"); signal != nil {
			f := func() QGeoAreaMonitorSource__AreaMonitorFeature {
				(*(*func() QGeoAreaMonitorSource__AreaMonitorFeature)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "supportedAreaMonitorFeatures", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "supportedAreaMonitorFeatures", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoAreaMonitorSource) DisconnectSupportedAreaMonitorFeatures() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "supportedAreaMonitorFeatures")
	}
}

func (ptr *QGeoAreaMonitorSource) SupportedAreaMonitorFeatures() QGeoAreaMonitorSource__AreaMonitorFeature {
	if ptr.Pointer() != nil {
		return QGeoAreaMonitorSource__AreaMonitorFeature(C.QGeoAreaMonitorSource_SupportedAreaMonitorFeatures(ptr.Pointer()))
	}
	return 0
}

//export callbackQGeoAreaMonitorSource_DestroyQGeoAreaMonitorSource
func callbackQGeoAreaMonitorSource_DestroyQGeoAreaMonitorSource(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QGeoAreaMonitorSource"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQGeoAreaMonitorSourceFromPointer(ptr).DestroyQGeoAreaMonitorSourceDefault()
	}
}

func (ptr *QGeoAreaMonitorSource) ConnectDestroyQGeoAreaMonitorSource(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QGeoAreaMonitorSource"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QGeoAreaMonitorSource", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QGeoAreaMonitorSource", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoAreaMonitorSource) DisconnectDestroyQGeoAreaMonitorSource() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QGeoAreaMonitorSource")
	}
}

func (ptr *QGeoAreaMonitorSource) DestroyQGeoAreaMonitorSource() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGeoAreaMonitorSource_DestroyQGeoAreaMonitorSource(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGeoAreaMonitorSource) DestroyQGeoAreaMonitorSourceDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGeoAreaMonitorSource_DestroyQGeoAreaMonitorSourceDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGeoAreaMonitorSource) __activeMonitors_atList(i int) *QGeoAreaMonitorInfo {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoAreaMonitorInfoFromPointer(C.QGeoAreaMonitorSource___activeMonitors_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QGeoAreaMonitorInfo).DestroyQGeoAreaMonitorInfo)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoAreaMonitorSource) __activeMonitors_setList(i QGeoAreaMonitorInfo_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource___activeMonitors_setList(ptr.Pointer(), PointerFromQGeoAreaMonitorInfo(i))
	}
}

func (ptr *QGeoAreaMonitorSource) __activeMonitors_newList() unsafe.Pointer {
	return C.QGeoAreaMonitorSource___activeMonitors_newList(ptr.Pointer())
}

func (ptr *QGeoAreaMonitorSource) __activeMonitors_atList2(i int) *QGeoAreaMonitorInfo {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoAreaMonitorInfoFromPointer(C.QGeoAreaMonitorSource___activeMonitors_atList2(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QGeoAreaMonitorInfo).DestroyQGeoAreaMonitorInfo)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoAreaMonitorSource) __activeMonitors_setList2(i QGeoAreaMonitorInfo_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource___activeMonitors_setList2(ptr.Pointer(), PointerFromQGeoAreaMonitorInfo(i))
	}
}

func (ptr *QGeoAreaMonitorSource) __activeMonitors_newList2() unsafe.Pointer {
	return C.QGeoAreaMonitorSource___activeMonitors_newList2(ptr.Pointer())
}

func (ptr *QGeoAreaMonitorSource) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGeoAreaMonitorSource___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGeoAreaMonitorSource) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGeoAreaMonitorSource) __children_newList() unsafe.Pointer {
	return C.QGeoAreaMonitorSource___children_newList(ptr.Pointer())
}

func (ptr *QGeoAreaMonitorSource) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QGeoAreaMonitorSource___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoAreaMonitorSource) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QGeoAreaMonitorSource) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QGeoAreaMonitorSource___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QGeoAreaMonitorSource) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGeoAreaMonitorSource___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGeoAreaMonitorSource) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGeoAreaMonitorSource) __findChildren_newList() unsafe.Pointer {
	return C.QGeoAreaMonitorSource___findChildren_newList(ptr.Pointer())
}

func (ptr *QGeoAreaMonitorSource) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGeoAreaMonitorSource___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGeoAreaMonitorSource) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGeoAreaMonitorSource) __findChildren_newList3() unsafe.Pointer {
	return C.QGeoAreaMonitorSource___findChildren_newList3(ptr.Pointer())
}

//export callbackQGeoAreaMonitorSource_ChildEvent
func callbackQGeoAreaMonitorSource_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQGeoAreaMonitorSourceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QGeoAreaMonitorSource) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQGeoAreaMonitorSource_ConnectNotify
func callbackQGeoAreaMonitorSource_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGeoAreaMonitorSourceFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGeoAreaMonitorSource) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGeoAreaMonitorSource_CustomEvent
func callbackQGeoAreaMonitorSource_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQGeoAreaMonitorSourceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGeoAreaMonitorSource) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQGeoAreaMonitorSource_DeleteLater
func callbackQGeoAreaMonitorSource_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQGeoAreaMonitorSourceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QGeoAreaMonitorSource) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGeoAreaMonitorSource_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQGeoAreaMonitorSource_Destroyed
func callbackQGeoAreaMonitorSource_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQGeoAreaMonitorSource_DisconnectNotify
func callbackQGeoAreaMonitorSource_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGeoAreaMonitorSourceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGeoAreaMonitorSource) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGeoAreaMonitorSource_Event
func callbackQGeoAreaMonitorSource_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGeoAreaMonitorSourceFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QGeoAreaMonitorSource) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGeoAreaMonitorSource_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQGeoAreaMonitorSource_EventFilter
func callbackQGeoAreaMonitorSource_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGeoAreaMonitorSourceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QGeoAreaMonitorSource) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGeoAreaMonitorSource_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQGeoAreaMonitorSource_MetaObject
func callbackQGeoAreaMonitorSource_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQGeoAreaMonitorSourceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QGeoAreaMonitorSource) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QGeoAreaMonitorSource_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQGeoAreaMonitorSource_ObjectNameChanged
func callbackQGeoAreaMonitorSource_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtPositioning_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQGeoAreaMonitorSource_TimerEvent
func callbackQGeoAreaMonitorSource_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGeoAreaMonitorSourceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGeoAreaMonitorSource) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
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

func NewQGeoCircleFromPointer(ptr unsafe.Pointer) (n *QGeoCircle) {
	n = new(QGeoCircle)
	n.SetPointer(ptr)
	return
}
func NewQGeoCircle() *QGeoCircle {
	tmpValue := NewQGeoCircleFromPointer(C.QGeoCircle_NewQGeoCircle())
	qt.SetFinalizer(tmpValue, (*QGeoCircle).DestroyQGeoCircle)
	return tmpValue
}

func NewQGeoCircle2(center QGeoCoordinate_ITF, radius float64) *QGeoCircle {
	tmpValue := NewQGeoCircleFromPointer(C.QGeoCircle_NewQGeoCircle2(PointerFromQGeoCoordinate(center), C.double(radius)))
	qt.SetFinalizer(tmpValue, (*QGeoCircle).DestroyQGeoCircle)
	return tmpValue
}

func NewQGeoCircle3(other QGeoCircle_ITF) *QGeoCircle {
	tmpValue := NewQGeoCircleFromPointer(C.QGeoCircle_NewQGeoCircle3(PointerFromQGeoCircle(other)))
	qt.SetFinalizer(tmpValue, (*QGeoCircle).DestroyQGeoCircle)
	return tmpValue
}

func NewQGeoCircle4(other QGeoShape_ITF) *QGeoCircle {
	tmpValue := NewQGeoCircleFromPointer(C.QGeoCircle_NewQGeoCircle4(PointerFromQGeoShape(other)))
	qt.SetFinalizer(tmpValue, (*QGeoCircle).DestroyQGeoCircle)
	return tmpValue
}

func (ptr *QGeoCircle) ExtendCircle(coordinate QGeoCoordinate_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoCircle_ExtendCircle(ptr.Pointer(), PointerFromQGeoCoordinate(coordinate))
	}
}

func (ptr *QGeoCircle) Radius() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGeoCircle_Radius(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoCircle) SetCenter(center QGeoCoordinate_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoCircle_SetCenter(ptr.Pointer(), PointerFromQGeoCoordinate(center))
	}
}

func (ptr *QGeoCircle) SetRadius(radius float64) {
	if ptr.Pointer() != nil {
		C.QGeoCircle_SetRadius(ptr.Pointer(), C.double(radius))
	}
}

func (ptr *QGeoCircle) Translate(degreesLatitude float64, degreesLongitude float64) {
	if ptr.Pointer() != nil {
		C.QGeoCircle_Translate(ptr.Pointer(), C.double(degreesLatitude), C.double(degreesLongitude))
	}
}

func (ptr *QGeoCircle) Translated(degreesLatitude float64, degreesLongitude float64) *QGeoCircle {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoCircleFromPointer(C.QGeoCircle_Translated(ptr.Pointer(), C.double(degreesLatitude), C.double(degreesLongitude)))
		qt.SetFinalizer(tmpValue, (*QGeoCircle).DestroyQGeoCircle)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoCircle) DestroyQGeoCircle() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGeoCircle_DestroyQGeoCircle(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QGeoCoordinate struct {
	ptr unsafe.Pointer
}

type QGeoCoordinate_ITF interface {
	QGeoCoordinate_PTR() *QGeoCoordinate
}

func (ptr *QGeoCoordinate) QGeoCoordinate_PTR() *QGeoCoordinate {
	return ptr
}

func (ptr *QGeoCoordinate) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QGeoCoordinate) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQGeoCoordinate(ptr QGeoCoordinate_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoCoordinate_PTR().Pointer()
	}
	return nil
}

func NewQGeoCoordinateFromPointer(ptr unsafe.Pointer) (n *QGeoCoordinate) {
	n = new(QGeoCoordinate)
	n.SetPointer(ptr)
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
	tmpValue := NewQGeoCoordinateFromPointer(C.QGeoCoordinate_NewQGeoCoordinate())
	qt.SetFinalizer(tmpValue, (*QGeoCoordinate).DestroyQGeoCoordinate)
	return tmpValue
}

func NewQGeoCoordinate2(latitude float64, longitude float64) *QGeoCoordinate {
	tmpValue := NewQGeoCoordinateFromPointer(C.QGeoCoordinate_NewQGeoCoordinate2(C.double(latitude), C.double(longitude)))
	qt.SetFinalizer(tmpValue, (*QGeoCoordinate).DestroyQGeoCoordinate)
	return tmpValue
}

func NewQGeoCoordinate3(latitude float64, longitude float64, altitude float64) *QGeoCoordinate {
	tmpValue := NewQGeoCoordinateFromPointer(C.QGeoCoordinate_NewQGeoCoordinate3(C.double(latitude), C.double(longitude), C.double(altitude)))
	qt.SetFinalizer(tmpValue, (*QGeoCoordinate).DestroyQGeoCoordinate)
	return tmpValue
}

func NewQGeoCoordinate4(other QGeoCoordinate_ITF) *QGeoCoordinate {
	tmpValue := NewQGeoCoordinateFromPointer(C.QGeoCoordinate_NewQGeoCoordinate4(PointerFromQGeoCoordinate(other)))
	qt.SetFinalizer(tmpValue, (*QGeoCoordinate).DestroyQGeoCoordinate)
	return tmpValue
}

func (ptr *QGeoCoordinate) Altitude() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGeoCoordinate_Altitude(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoCoordinate) AtDistanceAndAzimuth(distance float64, azimuth float64, distanceUp float64) *QGeoCoordinate {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoCoordinateFromPointer(C.QGeoCoordinate_AtDistanceAndAzimuth(ptr.Pointer(), C.double(distance), C.double(azimuth), C.double(distanceUp)))
		qt.SetFinalizer(tmpValue, (*QGeoCoordinate).DestroyQGeoCoordinate)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoCoordinate) AzimuthTo(other QGeoCoordinate_ITF) float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGeoCoordinate_AzimuthTo(ptr.Pointer(), PointerFromQGeoCoordinate(other)))
	}
	return 0
}

func (ptr *QGeoCoordinate) DistanceTo(other QGeoCoordinate_ITF) float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGeoCoordinate_DistanceTo(ptr.Pointer(), PointerFromQGeoCoordinate(other)))
	}
	return 0
}

func (ptr *QGeoCoordinate) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGeoCoordinate_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGeoCoordinate) Latitude() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGeoCoordinate_Latitude(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoCoordinate) Longitude() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGeoCoordinate_Longitude(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoCoordinate) SetAltitude(altitude float64) {
	if ptr.Pointer() != nil {
		C.QGeoCoordinate_SetAltitude(ptr.Pointer(), C.double(altitude))
	}
}

func (ptr *QGeoCoordinate) SetLatitude(latitude float64) {
	if ptr.Pointer() != nil {
		C.QGeoCoordinate_SetLatitude(ptr.Pointer(), C.double(latitude))
	}
}

func (ptr *QGeoCoordinate) SetLongitude(longitude float64) {
	if ptr.Pointer() != nil {
		C.QGeoCoordinate_SetLongitude(ptr.Pointer(), C.double(longitude))
	}
}

func (ptr *QGeoCoordinate) ToString(format QGeoCoordinate__CoordinateFormat) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoCoordinate_ToString(ptr.Pointer(), C.longlong(format)))
	}
	return ""
}

func (ptr *QGeoCoordinate) Type() QGeoCoordinate__CoordinateType {
	if ptr.Pointer() != nil {
		return QGeoCoordinate__CoordinateType(C.QGeoCoordinate_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoCoordinate) DestroyQGeoCoordinate() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGeoCoordinate_DestroyQGeoCoordinate(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QGeoLocation struct {
	ptr unsafe.Pointer
}

type QGeoLocation_ITF interface {
	QGeoLocation_PTR() *QGeoLocation
}

func (ptr *QGeoLocation) QGeoLocation_PTR() *QGeoLocation {
	return ptr
}

func (ptr *QGeoLocation) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QGeoLocation) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQGeoLocation(ptr QGeoLocation_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoLocation_PTR().Pointer()
	}
	return nil
}

func NewQGeoLocationFromPointer(ptr unsafe.Pointer) (n *QGeoLocation) {
	n = new(QGeoLocation)
	n.SetPointer(ptr)
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

func NewQGeoPathFromPointer(ptr unsafe.Pointer) (n *QGeoPath) {
	n = new(QGeoPath)
	n.SetPointer(ptr)
	return
}
func NewQGeoPath() *QGeoPath {
	tmpValue := NewQGeoPathFromPointer(C.QGeoPath_NewQGeoPath())
	qt.SetFinalizer(tmpValue, (*QGeoPath).DestroyQGeoPath)
	return tmpValue
}

func NewQGeoPath2(path []*QGeoCoordinate, width float64) *QGeoPath {
	tmpValue := NewQGeoPathFromPointer(C.QGeoPath_NewQGeoPath2(func() unsafe.Pointer {
		tmpList := NewQGeoPathFromPointer(NewQGeoPathFromPointer(nil).__QGeoPath_path_newList2())
		for _, v := range path {
			tmpList.__QGeoPath_path_setList2(v)
		}
		return tmpList.Pointer()
	}(), C.double(width)))
	qt.SetFinalizer(tmpValue, (*QGeoPath).DestroyQGeoPath)
	return tmpValue
}

func NewQGeoPath3(other QGeoPath_ITF) *QGeoPath {
	tmpValue := NewQGeoPathFromPointer(C.QGeoPath_NewQGeoPath3(PointerFromQGeoPath(other)))
	qt.SetFinalizer(tmpValue, (*QGeoPath).DestroyQGeoPath)
	return tmpValue
}

func NewQGeoPath4(other QGeoShape_ITF) *QGeoPath {
	tmpValue := NewQGeoPathFromPointer(C.QGeoPath_NewQGeoPath4(PointerFromQGeoShape(other)))
	qt.SetFinalizer(tmpValue, (*QGeoPath).DestroyQGeoPath)
	return tmpValue
}

func (ptr *QGeoPath) AddCoordinate(coordinate QGeoCoordinate_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPath_AddCoordinate(ptr.Pointer(), PointerFromQGeoCoordinate(coordinate))
	}
}

func (ptr *QGeoPath) ClearPath() {
	if ptr.Pointer() != nil {
		C.QGeoPath_ClearPath(ptr.Pointer())
	}
}

func (ptr *QGeoPath) ContainsCoordinate(coordinate QGeoCoordinate_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGeoPath_ContainsCoordinate(ptr.Pointer(), PointerFromQGeoCoordinate(coordinate))) != 0
	}
	return false
}

func (ptr *QGeoPath) CoordinateAt(index int) *QGeoCoordinate {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoCoordinateFromPointer(C.QGeoPath_CoordinateAt(ptr.Pointer(), C.int(int32(index))))
		qt.SetFinalizer(tmpValue, (*QGeoCoordinate).DestroyQGeoCoordinate)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoPath) InsertCoordinate(index int, coordinate QGeoCoordinate_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPath_InsertCoordinate(ptr.Pointer(), C.int(int32(index)), PointerFromQGeoCoordinate(coordinate))
	}
}

func (ptr *QGeoPath) Length(indexFrom int, indexTo int) float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGeoPath_Length(ptr.Pointer(), C.int(int32(indexFrom)), C.int(int32(indexTo))))
	}
	return 0
}

func (ptr *QGeoPath) Path() []*QGeoCoordinate {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtPositioning_PackedList) []*QGeoCoordinate {
			out := make([]*QGeoCoordinate, int(l.len))
			tmpList := NewQGeoPathFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__path_atList(i)
			}
			return out
		}(C.QGeoPath_Path(ptr.Pointer()))
	}
	return make([]*QGeoCoordinate, 0)
}

func (ptr *QGeoPath) RemoveCoordinate(coordinate QGeoCoordinate_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPath_RemoveCoordinate(ptr.Pointer(), PointerFromQGeoCoordinate(coordinate))
	}
}

func (ptr *QGeoPath) RemoveCoordinate2(index int) {
	if ptr.Pointer() != nil {
		C.QGeoPath_RemoveCoordinate2(ptr.Pointer(), C.int(int32(index)))
	}
}

func (ptr *QGeoPath) ReplaceCoordinate(index int, coordinate QGeoCoordinate_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPath_ReplaceCoordinate(ptr.Pointer(), C.int(int32(index)), PointerFromQGeoCoordinate(coordinate))
	}
}

func (ptr *QGeoPath) SetPath(path []*QGeoCoordinate) {
	if ptr.Pointer() != nil {
		C.QGeoPath_SetPath(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQGeoPathFromPointer(NewQGeoPathFromPointer(nil).__setPath_path_newList())
			for _, v := range path {
				tmpList.__setPath_path_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QGeoPath) SetWidth(width float64) {
	if ptr.Pointer() != nil {
		C.QGeoPath_SetWidth(ptr.Pointer(), C.double(width))
	}
}

func (ptr *QGeoPath) Size() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QGeoPath_Size(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoPath) Translate(degreesLatitude float64, degreesLongitude float64) {
	if ptr.Pointer() != nil {
		C.QGeoPath_Translate(ptr.Pointer(), C.double(degreesLatitude), C.double(degreesLongitude))
	}
}

func (ptr *QGeoPath) Translated(degreesLatitude float64, degreesLongitude float64) *QGeoPath {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoPathFromPointer(C.QGeoPath_Translated(ptr.Pointer(), C.double(degreesLatitude), C.double(degreesLongitude)))
		qt.SetFinalizer(tmpValue, (*QGeoPath).DestroyQGeoPath)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoPath) Width() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGeoPath_Width(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoPath) DestroyQGeoPath() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGeoPath_DestroyQGeoPath(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGeoPath) __QGeoPath_path_atList2(i int) *QGeoCoordinate {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoCoordinateFromPointer(C.QGeoPath___QGeoPath_path_atList2(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QGeoCoordinate).DestroyQGeoCoordinate)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoPath) __QGeoPath_path_setList2(i QGeoCoordinate_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPath___QGeoPath_path_setList2(ptr.Pointer(), PointerFromQGeoCoordinate(i))
	}
}

func (ptr *QGeoPath) __QGeoPath_path_newList2() unsafe.Pointer {
	return C.QGeoPath___QGeoPath_path_newList2(ptr.Pointer())
}

func (ptr *QGeoPath) __path_atList(i int) *QGeoCoordinate {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoCoordinateFromPointer(C.QGeoPath___path_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QGeoCoordinate).DestroyQGeoCoordinate)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoPath) __path_setList(i QGeoCoordinate_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPath___path_setList(ptr.Pointer(), PointerFromQGeoCoordinate(i))
	}
}

func (ptr *QGeoPath) __path_newList() unsafe.Pointer {
	return C.QGeoPath___path_newList(ptr.Pointer())
}

func (ptr *QGeoPath) __setPath_path_atList(i int) *QGeoCoordinate {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoCoordinateFromPointer(C.QGeoPath___setPath_path_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QGeoCoordinate).DestroyQGeoCoordinate)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoPath) __setPath_path_setList(i QGeoCoordinate_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPath___setPath_path_setList(ptr.Pointer(), PointerFromQGeoCoordinate(i))
	}
}

func (ptr *QGeoPath) __setPath_path_newList() unsafe.Pointer {
	return C.QGeoPath___setPath_path_newList(ptr.Pointer())
}

func (ptr *QGeoPath) __setVariantPath_path_atList(i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QGeoPath___setVariantPath_path_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoPath) __setVariantPath_path_setList(i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPath___setVariantPath_path_setList(ptr.Pointer(), core.PointerFromQVariant(i))
	}
}

func (ptr *QGeoPath) __setVariantPath_path_newList() unsafe.Pointer {
	return C.QGeoPath___setVariantPath_path_newList(ptr.Pointer())
}

func (ptr *QGeoPath) __variantPath_atList(i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QGeoPath___variantPath_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoPath) __variantPath_setList(i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPath___variantPath_setList(ptr.Pointer(), core.PointerFromQVariant(i))
	}
}

func (ptr *QGeoPath) __variantPath_newList() unsafe.Pointer {
	return C.QGeoPath___variantPath_newList(ptr.Pointer())
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

func NewQGeoPolygonFromPointer(ptr unsafe.Pointer) (n *QGeoPolygon) {
	n = new(QGeoPolygon)
	n.SetPointer(ptr)
	return
}
func NewQGeoPolygon() *QGeoPolygon {
	tmpValue := NewQGeoPolygonFromPointer(C.QGeoPolygon_NewQGeoPolygon())
	qt.SetFinalizer(tmpValue, (*QGeoPolygon).DestroyQGeoPolygon)
	return tmpValue
}

func NewQGeoPolygon2(path []*QGeoCoordinate) *QGeoPolygon {
	tmpValue := NewQGeoPolygonFromPointer(C.QGeoPolygon_NewQGeoPolygon2(func() unsafe.Pointer {
		tmpList := NewQGeoPolygonFromPointer(NewQGeoPolygonFromPointer(nil).__QGeoPolygon_path_newList2())
		for _, v := range path {
			tmpList.__QGeoPolygon_path_setList2(v)
		}
		return tmpList.Pointer()
	}()))
	qt.SetFinalizer(tmpValue, (*QGeoPolygon).DestroyQGeoPolygon)
	return tmpValue
}

func NewQGeoPolygon3(other QGeoPolygon_ITF) *QGeoPolygon {
	tmpValue := NewQGeoPolygonFromPointer(C.QGeoPolygon_NewQGeoPolygon3(PointerFromQGeoPolygon(other)))
	qt.SetFinalizer(tmpValue, (*QGeoPolygon).DestroyQGeoPolygon)
	return tmpValue
}

func NewQGeoPolygon4(other QGeoShape_ITF) *QGeoPolygon {
	tmpValue := NewQGeoPolygonFromPointer(C.QGeoPolygon_NewQGeoPolygon4(PointerFromQGeoShape(other)))
	qt.SetFinalizer(tmpValue, (*QGeoPolygon).DestroyQGeoPolygon)
	return tmpValue
}

func (ptr *QGeoPolygon) AddCoordinate(coordinate QGeoCoordinate_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPolygon_AddCoordinate(ptr.Pointer(), PointerFromQGeoCoordinate(coordinate))
	}
}

func (ptr *QGeoPolygon) AddHole(holePath core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPolygon_AddHole(ptr.Pointer(), core.PointerFromQVariant(holePath))
	}
}

func (ptr *QGeoPolygon) AddHole2(holePath []*QGeoCoordinate) {
	if ptr.Pointer() != nil {
		C.QGeoPolygon_AddHole2(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQGeoPolygonFromPointer(NewQGeoPolygonFromPointer(nil).__addHole_holePath_newList2())
			for _, v := range holePath {
				tmpList.__addHole_holePath_setList2(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QGeoPolygon) ContainsCoordinate(coordinate QGeoCoordinate_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGeoPolygon_ContainsCoordinate(ptr.Pointer(), PointerFromQGeoCoordinate(coordinate))) != 0
	}
	return false
}

func (ptr *QGeoPolygon) CoordinateAt(index int) *QGeoCoordinate {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoCoordinateFromPointer(C.QGeoPolygon_CoordinateAt(ptr.Pointer(), C.int(int32(index))))
		qt.SetFinalizer(tmpValue, (*QGeoCoordinate).DestroyQGeoCoordinate)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoPolygon) Hole(index int) []*core.QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtPositioning_PackedList) []*core.QVariant {
			out := make([]*core.QVariant, int(l.len))
			tmpList := NewQGeoPolygonFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__hole_atList(i)
			}
			return out
		}(C.QGeoPolygon_Hole(ptr.Pointer(), C.int(int32(index))))
	}
	return make([]*core.QVariant, 0)
}

func (ptr *QGeoPolygon) HolePath(index int) []*QGeoCoordinate {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtPositioning_PackedList) []*QGeoCoordinate {
			out := make([]*QGeoCoordinate, int(l.len))
			tmpList := NewQGeoPolygonFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__holePath_atList(i)
			}
			return out
		}(C.QGeoPolygon_HolePath(ptr.Pointer(), C.int(int32(index))))
	}
	return make([]*QGeoCoordinate, 0)
}

func (ptr *QGeoPolygon) HolesCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QGeoPolygon_HolesCount(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoPolygon) InsertCoordinate(index int, coordinate QGeoCoordinate_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPolygon_InsertCoordinate(ptr.Pointer(), C.int(int32(index)), PointerFromQGeoCoordinate(coordinate))
	}
}

func (ptr *QGeoPolygon) Length(indexFrom int, indexTo int) float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGeoPolygon_Length(ptr.Pointer(), C.int(int32(indexFrom)), C.int(int32(indexTo))))
	}
	return 0
}

func (ptr *QGeoPolygon) Path() []*QGeoCoordinate {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtPositioning_PackedList) []*QGeoCoordinate {
			out := make([]*QGeoCoordinate, int(l.len))
			tmpList := NewQGeoPolygonFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__path_atList(i)
			}
			return out
		}(C.QGeoPolygon_Path(ptr.Pointer()))
	}
	return make([]*QGeoCoordinate, 0)
}

func (ptr *QGeoPolygon) Perimeter() []*core.QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtPositioning_PackedList) []*core.QVariant {
			out := make([]*core.QVariant, int(l.len))
			tmpList := NewQGeoPolygonFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__perimeter_atList(i)
			}
			return out
		}(C.QGeoPolygon_Perimeter(ptr.Pointer()))
	}
	return make([]*core.QVariant, 0)
}

func (ptr *QGeoPolygon) RemoveCoordinate(coordinate QGeoCoordinate_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPolygon_RemoveCoordinate(ptr.Pointer(), PointerFromQGeoCoordinate(coordinate))
	}
}

func (ptr *QGeoPolygon) RemoveCoordinate2(index int) {
	if ptr.Pointer() != nil {
		C.QGeoPolygon_RemoveCoordinate2(ptr.Pointer(), C.int(int32(index)))
	}
}

func (ptr *QGeoPolygon) RemoveHole(index int) {
	if ptr.Pointer() != nil {
		C.QGeoPolygon_RemoveHole(ptr.Pointer(), C.int(int32(index)))
	}
}

func (ptr *QGeoPolygon) ReplaceCoordinate(index int, coordinate QGeoCoordinate_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPolygon_ReplaceCoordinate(ptr.Pointer(), C.int(int32(index)), PointerFromQGeoCoordinate(coordinate))
	}
}

func (ptr *QGeoPolygon) SetPath(path []*QGeoCoordinate) {
	if ptr.Pointer() != nil {
		C.QGeoPolygon_SetPath(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQGeoPolygonFromPointer(NewQGeoPolygonFromPointer(nil).__setPath_path_newList())
			for _, v := range path {
				tmpList.__setPath_path_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QGeoPolygon) SetPerimeter(path []*core.QVariant) {
	if ptr.Pointer() != nil {
		C.QGeoPolygon_SetPerimeter(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQGeoPolygonFromPointer(NewQGeoPolygonFromPointer(nil).__setPerimeter_path_newList())
			for _, v := range path {
				tmpList.__setPerimeter_path_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QGeoPolygon) Size() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QGeoPolygon_Size(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoPolygon) Translate(degreesLatitude float64, degreesLongitude float64) {
	if ptr.Pointer() != nil {
		C.QGeoPolygon_Translate(ptr.Pointer(), C.double(degreesLatitude), C.double(degreesLongitude))
	}
}

func (ptr *QGeoPolygon) Translated(degreesLatitude float64, degreesLongitude float64) *QGeoPolygon {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoPolygonFromPointer(C.QGeoPolygon_Translated(ptr.Pointer(), C.double(degreesLatitude), C.double(degreesLongitude)))
		qt.SetFinalizer(tmpValue, (*QGeoPolygon).DestroyQGeoPolygon)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoPolygon) DestroyQGeoPolygon() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGeoPolygon_DestroyQGeoPolygon(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGeoPolygon) __QGeoPolygon_path_atList2(i int) *QGeoCoordinate {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoCoordinateFromPointer(C.QGeoPolygon___QGeoPolygon_path_atList2(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QGeoCoordinate).DestroyQGeoCoordinate)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoPolygon) __QGeoPolygon_path_setList2(i QGeoCoordinate_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPolygon___QGeoPolygon_path_setList2(ptr.Pointer(), PointerFromQGeoCoordinate(i))
	}
}

func (ptr *QGeoPolygon) __QGeoPolygon_path_newList2() unsafe.Pointer {
	return C.QGeoPolygon___QGeoPolygon_path_newList2(ptr.Pointer())
}

func (ptr *QGeoPolygon) __addHole_holePath_atList2(i int) *QGeoCoordinate {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoCoordinateFromPointer(C.QGeoPolygon___addHole_holePath_atList2(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QGeoCoordinate).DestroyQGeoCoordinate)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoPolygon) __addHole_holePath_setList2(i QGeoCoordinate_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPolygon___addHole_holePath_setList2(ptr.Pointer(), PointerFromQGeoCoordinate(i))
	}
}

func (ptr *QGeoPolygon) __addHole_holePath_newList2() unsafe.Pointer {
	return C.QGeoPolygon___addHole_holePath_newList2(ptr.Pointer())
}

func (ptr *QGeoPolygon) __hole_atList(i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QGeoPolygon___hole_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoPolygon) __hole_setList(i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPolygon___hole_setList(ptr.Pointer(), core.PointerFromQVariant(i))
	}
}

func (ptr *QGeoPolygon) __hole_newList() unsafe.Pointer {
	return C.QGeoPolygon___hole_newList(ptr.Pointer())
}

func (ptr *QGeoPolygon) __holePath_atList(i int) *QGeoCoordinate {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoCoordinateFromPointer(C.QGeoPolygon___holePath_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QGeoCoordinate).DestroyQGeoCoordinate)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoPolygon) __holePath_setList(i QGeoCoordinate_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPolygon___holePath_setList(ptr.Pointer(), PointerFromQGeoCoordinate(i))
	}
}

func (ptr *QGeoPolygon) __holePath_newList() unsafe.Pointer {
	return C.QGeoPolygon___holePath_newList(ptr.Pointer())
}

func (ptr *QGeoPolygon) __path_atList(i int) *QGeoCoordinate {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoCoordinateFromPointer(C.QGeoPolygon___path_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QGeoCoordinate).DestroyQGeoCoordinate)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoPolygon) __path_setList(i QGeoCoordinate_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPolygon___path_setList(ptr.Pointer(), PointerFromQGeoCoordinate(i))
	}
}

func (ptr *QGeoPolygon) __path_newList() unsafe.Pointer {
	return C.QGeoPolygon___path_newList(ptr.Pointer())
}

func (ptr *QGeoPolygon) __perimeter_atList(i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QGeoPolygon___perimeter_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoPolygon) __perimeter_setList(i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPolygon___perimeter_setList(ptr.Pointer(), core.PointerFromQVariant(i))
	}
}

func (ptr *QGeoPolygon) __perimeter_newList() unsafe.Pointer {
	return C.QGeoPolygon___perimeter_newList(ptr.Pointer())
}

func (ptr *QGeoPolygon) __setPath_path_atList(i int) *QGeoCoordinate {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoCoordinateFromPointer(C.QGeoPolygon___setPath_path_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QGeoCoordinate).DestroyQGeoCoordinate)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoPolygon) __setPath_path_setList(i QGeoCoordinate_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPolygon___setPath_path_setList(ptr.Pointer(), PointerFromQGeoCoordinate(i))
	}
}

func (ptr *QGeoPolygon) __setPath_path_newList() unsafe.Pointer {
	return C.QGeoPolygon___setPath_path_newList(ptr.Pointer())
}

func (ptr *QGeoPolygon) __setPerimeter_path_atList(i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QGeoPolygon___setPerimeter_path_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoPolygon) __setPerimeter_path_setList(i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPolygon___setPerimeter_path_setList(ptr.Pointer(), core.PointerFromQVariant(i))
	}
}

func (ptr *QGeoPolygon) __setPerimeter_path_newList() unsafe.Pointer {
	return C.QGeoPolygon___setPerimeter_path_newList(ptr.Pointer())
}

type QGeoPositionInfo struct {
	ptr unsafe.Pointer
}

type QGeoPositionInfo_ITF interface {
	QGeoPositionInfo_PTR() *QGeoPositionInfo
}

func (ptr *QGeoPositionInfo) QGeoPositionInfo_PTR() *QGeoPositionInfo {
	return ptr
}

func (ptr *QGeoPositionInfo) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QGeoPositionInfo) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQGeoPositionInfo(ptr QGeoPositionInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoPositionInfo_PTR().Pointer()
	}
	return nil
}

func NewQGeoPositionInfoFromPointer(ptr unsafe.Pointer) (n *QGeoPositionInfo) {
	n = new(QGeoPositionInfo)
	n.SetPointer(ptr)
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
	tmpValue := NewQGeoPositionInfoFromPointer(C.QGeoPositionInfo_NewQGeoPositionInfo())
	qt.SetFinalizer(tmpValue, (*QGeoPositionInfo).DestroyQGeoPositionInfo)
	return tmpValue
}

func NewQGeoPositionInfo2(coordinate QGeoCoordinate_ITF, timestamp core.QDateTime_ITF) *QGeoPositionInfo {
	tmpValue := NewQGeoPositionInfoFromPointer(C.QGeoPositionInfo_NewQGeoPositionInfo2(PointerFromQGeoCoordinate(coordinate), core.PointerFromQDateTime(timestamp)))
	qt.SetFinalizer(tmpValue, (*QGeoPositionInfo).DestroyQGeoPositionInfo)
	return tmpValue
}

func NewQGeoPositionInfo3(other QGeoPositionInfo_ITF) *QGeoPositionInfo {
	tmpValue := NewQGeoPositionInfoFromPointer(C.QGeoPositionInfo_NewQGeoPositionInfo3(PointerFromQGeoPositionInfo(other)))
	qt.SetFinalizer(tmpValue, (*QGeoPositionInfo).DestroyQGeoPositionInfo)
	return tmpValue
}

func (ptr *QGeoPositionInfo) Attribute(attribute QGeoPositionInfo__Attribute) float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGeoPositionInfo_Attribute(ptr.Pointer(), C.longlong(attribute)))
	}
	return 0
}

func (ptr *QGeoPositionInfo) Coordinate() *QGeoCoordinate {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoCoordinateFromPointer(C.QGeoPositionInfo_Coordinate(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QGeoCoordinate).DestroyQGeoCoordinate)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoPositionInfo) HasAttribute(attribute QGeoPositionInfo__Attribute) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGeoPositionInfo_HasAttribute(ptr.Pointer(), C.longlong(attribute))) != 0
	}
	return false
}

func (ptr *QGeoPositionInfo) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGeoPositionInfo_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGeoPositionInfo) RemoveAttribute(attribute QGeoPositionInfo__Attribute) {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfo_RemoveAttribute(ptr.Pointer(), C.longlong(attribute))
	}
}

func (ptr *QGeoPositionInfo) SetAttribute(attribute QGeoPositionInfo__Attribute, value float64) {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfo_SetAttribute(ptr.Pointer(), C.longlong(attribute), C.double(value))
	}
}

func (ptr *QGeoPositionInfo) SetCoordinate(coordinate QGeoCoordinate_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfo_SetCoordinate(ptr.Pointer(), PointerFromQGeoCoordinate(coordinate))
	}
}

func (ptr *QGeoPositionInfo) SetTimestamp(timestamp core.QDateTime_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfo_SetTimestamp(ptr.Pointer(), core.PointerFromQDateTime(timestamp))
	}
}

func (ptr *QGeoPositionInfo) Timestamp() *core.QDateTime {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQDateTimeFromPointer(C.QGeoPositionInfo_Timestamp(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QDateTime).DestroyQDateTime)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoPositionInfo) DestroyQGeoPositionInfo() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGeoPositionInfo_DestroyQGeoPositionInfo(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
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

func NewQGeoPositionInfoSourceFromPointer(ptr unsafe.Pointer) (n *QGeoPositionInfoSource) {
	n = new(QGeoPositionInfoSource)
	n.SetPointer(ptr)
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
	tmpValue := NewQGeoPositionInfoSourceFromPointer(C.QGeoPositionInfoSource_NewQGeoPositionInfoSource(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QGeoPositionInfoSource_AvailableSources() []string {
	return unpackStringList(cGoUnpackString(C.QGeoPositionInfoSource_QGeoPositionInfoSource_AvailableSources()))
}

func (ptr *QGeoPositionInfoSource) AvailableSources() []string {
	return unpackStringList(cGoUnpackString(C.QGeoPositionInfoSource_QGeoPositionInfoSource_AvailableSources()))
}

func QGeoPositionInfoSource_CreateDefaultSource(parent core.QObject_ITF) *QGeoPositionInfoSource {
	tmpValue := NewQGeoPositionInfoSourceFromPointer(C.QGeoPositionInfoSource_QGeoPositionInfoSource_CreateDefaultSource(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QGeoPositionInfoSource) CreateDefaultSource(parent core.QObject_ITF) *QGeoPositionInfoSource {
	tmpValue := NewQGeoPositionInfoSourceFromPointer(C.QGeoPositionInfoSource_QGeoPositionInfoSource_CreateDefaultSource(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QGeoPositionInfoSource_CreateSource(sourceName string, parent core.QObject_ITF) *QGeoPositionInfoSource {
	var sourceNameC *C.char
	if sourceName != "" {
		sourceNameC = C.CString(sourceName)
		defer C.free(unsafe.Pointer(sourceNameC))
	}
	tmpValue := NewQGeoPositionInfoSourceFromPointer(C.QGeoPositionInfoSource_QGeoPositionInfoSource_CreateSource(C.struct_QtPositioning_PackedString{data: sourceNameC, len: C.longlong(len(sourceName))}, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QGeoPositionInfoSource) CreateSource(sourceName string, parent core.QObject_ITF) *QGeoPositionInfoSource {
	var sourceNameC *C.char
	if sourceName != "" {
		sourceNameC = C.CString(sourceName)
		defer C.free(unsafe.Pointer(sourceNameC))
	}
	tmpValue := NewQGeoPositionInfoSourceFromPointer(C.QGeoPositionInfoSource_QGeoPositionInfoSource_CreateSource(C.struct_QtPositioning_PackedString{data: sourceNameC, len: C.longlong(len(sourceName))}, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQGeoPositionInfoSource_Error
func callbackQGeoPositionInfoSource_Error(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "error"); signal != nil {
		return C.longlong((*(*func() QGeoPositionInfoSource__Error)(signal))())
	}

	return C.longlong(0)
}

func (ptr *QGeoPositionInfoSource) ConnectError(f func() QGeoPositionInfoSource__Error) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "error"); signal != nil {
			f := func() QGeoPositionInfoSource__Error {
				(*(*func() QGeoPositionInfoSource__Error)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "error", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "error", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectError() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "error")
	}
}

func (ptr *QGeoPositionInfoSource) Error() QGeoPositionInfoSource__Error {
	if ptr.Pointer() != nil {
		return QGeoPositionInfoSource__Error(C.QGeoPositionInfoSource_Error(ptr.Pointer()))
	}
	return 0
}

//export callbackQGeoPositionInfoSource_Error2
func callbackQGeoPositionInfoSource_Error2(ptr unsafe.Pointer, positioningError C.longlong) {
	if signal := qt.GetSignal(ptr, "error2"); signal != nil {
		(*(*func(QGeoPositionInfoSource__Error))(signal))(QGeoPositionInfoSource__Error(positioningError))
	}

}

func (ptr *QGeoPositionInfoSource) ConnectError2(f func(positioningError QGeoPositionInfoSource__Error)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "error2") {
			C.QGeoPositionInfoSource_ConnectError2(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "error")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "error2"); signal != nil {
			f := func(positioningError QGeoPositionInfoSource__Error) {
				(*(*func(QGeoPositionInfoSource__Error))(signal))(positioningError)
				f(positioningError)
			}
			qt.ConnectSignal(ptr.Pointer(), "error2", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "error2", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectError2() {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "error2")
	}
}

func (ptr *QGeoPositionInfoSource) Error2(positioningError QGeoPositionInfoSource__Error) {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_Error2(ptr.Pointer(), C.longlong(positioningError))
	}
}

//export callbackQGeoPositionInfoSource_LastKnownPosition
func callbackQGeoPositionInfoSource_LastKnownPosition(ptr unsafe.Pointer, fromSatellitePositioningMethodsOnly C.char) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "lastKnownPosition"); signal != nil {
		return PointerFromQGeoPositionInfo((*(*func(bool) *QGeoPositionInfo)(signal))(int8(fromSatellitePositioningMethodsOnly) != 0))
	}

	return PointerFromQGeoPositionInfo(NewQGeoPositionInfo())
}

func (ptr *QGeoPositionInfoSource) ConnectLastKnownPosition(f func(fromSatellitePositioningMethodsOnly bool) *QGeoPositionInfo) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "lastKnownPosition"); signal != nil {
			f := func(fromSatellitePositioningMethodsOnly bool) *QGeoPositionInfo {
				(*(*func(bool) *QGeoPositionInfo)(signal))(fromSatellitePositioningMethodsOnly)
				return f(fromSatellitePositioningMethodsOnly)
			}
			qt.ConnectSignal(ptr.Pointer(), "lastKnownPosition", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "lastKnownPosition", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectLastKnownPosition() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "lastKnownPosition")
	}
}

func (ptr *QGeoPositionInfoSource) LastKnownPosition(fromSatellitePositioningMethodsOnly bool) *QGeoPositionInfo {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoPositionInfoFromPointer(C.QGeoPositionInfoSource_LastKnownPosition(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(fromSatellitePositioningMethodsOnly)))))
		qt.SetFinalizer(tmpValue, (*QGeoPositionInfo).DestroyQGeoPositionInfo)
		return tmpValue
	}
	return nil
}

//export callbackQGeoPositionInfoSource_MinimumUpdateInterval
func callbackQGeoPositionInfoSource_MinimumUpdateInterval(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "minimumUpdateInterval"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(0))
}

func (ptr *QGeoPositionInfoSource) ConnectMinimumUpdateInterval(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "minimumUpdateInterval"); signal != nil {
			f := func() int {
				(*(*func() int)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "minimumUpdateInterval", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "minimumUpdateInterval", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectMinimumUpdateInterval() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "minimumUpdateInterval")
	}
}

func (ptr *QGeoPositionInfoSource) MinimumUpdateInterval() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QGeoPositionInfoSource_MinimumUpdateInterval(ptr.Pointer())))
	}
	return 0
}

//export callbackQGeoPositionInfoSource_PositionUpdated
func callbackQGeoPositionInfoSource_PositionUpdated(ptr unsafe.Pointer, update unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "positionUpdated"); signal != nil {
		(*(*func(*QGeoPositionInfo))(signal))(NewQGeoPositionInfoFromPointer(update))
	}

}

func (ptr *QGeoPositionInfoSource) ConnectPositionUpdated(f func(update *QGeoPositionInfo)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "positionUpdated") {
			C.QGeoPositionInfoSource_ConnectPositionUpdated(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "positionUpdated")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "positionUpdated"); signal != nil {
			f := func(update *QGeoPositionInfo) {
				(*(*func(*QGeoPositionInfo))(signal))(update)
				f(update)
			}
			qt.ConnectSignal(ptr.Pointer(), "positionUpdated", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "positionUpdated", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectPositionUpdated() {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_DisconnectPositionUpdated(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "positionUpdated")
	}
}

func (ptr *QGeoPositionInfoSource) PositionUpdated(update QGeoPositionInfo_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_PositionUpdated(ptr.Pointer(), PointerFromQGeoPositionInfo(update))
	}
}

func (ptr *QGeoPositionInfoSource) PreferredPositioningMethods() QGeoPositionInfoSource__PositioningMethod {
	if ptr.Pointer() != nil {
		return QGeoPositionInfoSource__PositioningMethod(C.QGeoPositionInfoSource_PreferredPositioningMethods(ptr.Pointer()))
	}
	return 0
}

//export callbackQGeoPositionInfoSource_RequestUpdate
func callbackQGeoPositionInfoSource_RequestUpdate(ptr unsafe.Pointer, timeout C.int) {
	if signal := qt.GetSignal(ptr, "requestUpdate"); signal != nil {
		(*(*func(int))(signal))(int(int32(timeout)))
	}

}

func (ptr *QGeoPositionInfoSource) ConnectRequestUpdate(f func(timeout int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "requestUpdate"); signal != nil {
			f := func(timeout int) {
				(*(*func(int))(signal))(timeout)
				f(timeout)
			}
			qt.ConnectSignal(ptr.Pointer(), "requestUpdate", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "requestUpdate", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectRequestUpdate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "requestUpdate")
	}
}

func (ptr *QGeoPositionInfoSource) RequestUpdate(timeout int) {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_RequestUpdate(ptr.Pointer(), C.int(int32(timeout)))
	}
}

//export callbackQGeoPositionInfoSource_SetPreferredPositioningMethods
func callbackQGeoPositionInfoSource_SetPreferredPositioningMethods(ptr unsafe.Pointer, methods C.longlong) {
	if signal := qt.GetSignal(ptr, "setPreferredPositioningMethods"); signal != nil {
		(*(*func(QGeoPositionInfoSource__PositioningMethod))(signal))(QGeoPositionInfoSource__PositioningMethod(methods))
	} else {
		NewQGeoPositionInfoSourceFromPointer(ptr).SetPreferredPositioningMethodsDefault(QGeoPositionInfoSource__PositioningMethod(methods))
	}
}

func (ptr *QGeoPositionInfoSource) ConnectSetPreferredPositioningMethods(f func(methods QGeoPositionInfoSource__PositioningMethod)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setPreferredPositioningMethods"); signal != nil {
			f := func(methods QGeoPositionInfoSource__PositioningMethod) {
				(*(*func(QGeoPositionInfoSource__PositioningMethod))(signal))(methods)
				f(methods)
			}
			qt.ConnectSignal(ptr.Pointer(), "setPreferredPositioningMethods", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setPreferredPositioningMethods", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectSetPreferredPositioningMethods() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setPreferredPositioningMethods")
	}
}

func (ptr *QGeoPositionInfoSource) SetPreferredPositioningMethods(methods QGeoPositionInfoSource__PositioningMethod) {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_SetPreferredPositioningMethods(ptr.Pointer(), C.longlong(methods))
	}
}

func (ptr *QGeoPositionInfoSource) SetPreferredPositioningMethodsDefault(methods QGeoPositionInfoSource__PositioningMethod) {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_SetPreferredPositioningMethodsDefault(ptr.Pointer(), C.longlong(methods))
	}
}

//export callbackQGeoPositionInfoSource_SetUpdateInterval
func callbackQGeoPositionInfoSource_SetUpdateInterval(ptr unsafe.Pointer, msec C.int) {
	if signal := qt.GetSignal(ptr, "setUpdateInterval"); signal != nil {
		(*(*func(int))(signal))(int(int32(msec)))
	} else {
		NewQGeoPositionInfoSourceFromPointer(ptr).SetUpdateIntervalDefault(int(int32(msec)))
	}
}

func (ptr *QGeoPositionInfoSource) ConnectSetUpdateInterval(f func(msec int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setUpdateInterval"); signal != nil {
			f := func(msec int) {
				(*(*func(int))(signal))(msec)
				f(msec)
			}
			qt.ConnectSignal(ptr.Pointer(), "setUpdateInterval", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setUpdateInterval", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectSetUpdateInterval() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setUpdateInterval")
	}
}

func (ptr *QGeoPositionInfoSource) SetUpdateInterval(msec int) {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_SetUpdateInterval(ptr.Pointer(), C.int(int32(msec)))
	}
}

func (ptr *QGeoPositionInfoSource) SetUpdateIntervalDefault(msec int) {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_SetUpdateIntervalDefault(ptr.Pointer(), C.int(int32(msec)))
	}
}

func (ptr *QGeoPositionInfoSource) SourceName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoPositionInfoSource_SourceName(ptr.Pointer()))
	}
	return ""
}

//export callbackQGeoPositionInfoSource_StartUpdates
func callbackQGeoPositionInfoSource_StartUpdates(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "startUpdates"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QGeoPositionInfoSource) ConnectStartUpdates(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "startUpdates"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "startUpdates", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "startUpdates", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectStartUpdates() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "startUpdates")
	}
}

func (ptr *QGeoPositionInfoSource) StartUpdates() {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_StartUpdates(ptr.Pointer())
	}
}

//export callbackQGeoPositionInfoSource_StopUpdates
func callbackQGeoPositionInfoSource_StopUpdates(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "stopUpdates"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QGeoPositionInfoSource) ConnectStopUpdates(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "stopUpdates"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "stopUpdates", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stopUpdates", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectStopUpdates() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "stopUpdates")
	}
}

func (ptr *QGeoPositionInfoSource) StopUpdates() {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_StopUpdates(ptr.Pointer())
	}
}

//export callbackQGeoPositionInfoSource_SupportedPositioningMethods
func callbackQGeoPositionInfoSource_SupportedPositioningMethods(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedPositioningMethods"); signal != nil {
		return C.longlong((*(*func() QGeoPositionInfoSource__PositioningMethod)(signal))())
	}

	return C.longlong(0)
}

func (ptr *QGeoPositionInfoSource) ConnectSupportedPositioningMethods(f func() QGeoPositionInfoSource__PositioningMethod) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "supportedPositioningMethods"); signal != nil {
			f := func() QGeoPositionInfoSource__PositioningMethod {
				(*(*func() QGeoPositionInfoSource__PositioningMethod)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "supportedPositioningMethods", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "supportedPositioningMethods", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectSupportedPositioningMethods() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "supportedPositioningMethods")
	}
}

func (ptr *QGeoPositionInfoSource) SupportedPositioningMethods() QGeoPositionInfoSource__PositioningMethod {
	if ptr.Pointer() != nil {
		return QGeoPositionInfoSource__PositioningMethod(C.QGeoPositionInfoSource_SupportedPositioningMethods(ptr.Pointer()))
	}
	return 0
}

//export callbackQGeoPositionInfoSource_SupportedPositioningMethodsChanged
func callbackQGeoPositionInfoSource_SupportedPositioningMethodsChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "supportedPositioningMethodsChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QGeoPositionInfoSource) ConnectSupportedPositioningMethodsChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "supportedPositioningMethodsChanged") {
			C.QGeoPositionInfoSource_ConnectSupportedPositioningMethodsChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "supportedPositioningMethodsChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "supportedPositioningMethodsChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "supportedPositioningMethodsChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "supportedPositioningMethodsChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectSupportedPositioningMethodsChanged() {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_DisconnectSupportedPositioningMethodsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "supportedPositioningMethodsChanged")
	}
}

func (ptr *QGeoPositionInfoSource) SupportedPositioningMethodsChanged() {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_SupportedPositioningMethodsChanged(ptr.Pointer())
	}
}

func (ptr *QGeoPositionInfoSource) UpdateInterval() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QGeoPositionInfoSource_UpdateInterval(ptr.Pointer())))
	}
	return 0
}

//export callbackQGeoPositionInfoSource_UpdateTimeout
func callbackQGeoPositionInfoSource_UpdateTimeout(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateTimeout"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QGeoPositionInfoSource) ConnectUpdateTimeout(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "updateTimeout") {
			C.QGeoPositionInfoSource_ConnectUpdateTimeout(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "updateTimeout")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "updateTimeout"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "updateTimeout", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "updateTimeout", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectUpdateTimeout() {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_DisconnectUpdateTimeout(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "updateTimeout")
	}
}

func (ptr *QGeoPositionInfoSource) UpdateTimeout() {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_UpdateTimeout(ptr.Pointer())
	}
}

//export callbackQGeoPositionInfoSource_DestroyQGeoPositionInfoSource
func callbackQGeoPositionInfoSource_DestroyQGeoPositionInfoSource(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QGeoPositionInfoSource"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQGeoPositionInfoSourceFromPointer(ptr).DestroyQGeoPositionInfoSourceDefault()
	}
}

func (ptr *QGeoPositionInfoSource) ConnectDestroyQGeoPositionInfoSource(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QGeoPositionInfoSource"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QGeoPositionInfoSource", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QGeoPositionInfoSource", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectDestroyQGeoPositionInfoSource() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QGeoPositionInfoSource")
	}
}

func (ptr *QGeoPositionInfoSource) DestroyQGeoPositionInfoSource() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGeoPositionInfoSource_DestroyQGeoPositionInfoSource(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGeoPositionInfoSource) DestroyQGeoPositionInfoSourceDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGeoPositionInfoSource_DestroyQGeoPositionInfoSourceDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGeoPositionInfoSource) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGeoPositionInfoSource___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGeoPositionInfoSource) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGeoPositionInfoSource) __children_newList() unsafe.Pointer {
	return C.QGeoPositionInfoSource___children_newList(ptr.Pointer())
}

func (ptr *QGeoPositionInfoSource) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QGeoPositionInfoSource___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoPositionInfoSource) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QGeoPositionInfoSource) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QGeoPositionInfoSource___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QGeoPositionInfoSource) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGeoPositionInfoSource___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGeoPositionInfoSource) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGeoPositionInfoSource) __findChildren_newList() unsafe.Pointer {
	return C.QGeoPositionInfoSource___findChildren_newList(ptr.Pointer())
}

func (ptr *QGeoPositionInfoSource) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGeoPositionInfoSource___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGeoPositionInfoSource) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGeoPositionInfoSource) __findChildren_newList3() unsafe.Pointer {
	return C.QGeoPositionInfoSource___findChildren_newList3(ptr.Pointer())
}

//export callbackQGeoPositionInfoSource_ChildEvent
func callbackQGeoPositionInfoSource_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQGeoPositionInfoSourceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QGeoPositionInfoSource) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQGeoPositionInfoSource_ConnectNotify
func callbackQGeoPositionInfoSource_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGeoPositionInfoSourceFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGeoPositionInfoSource) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGeoPositionInfoSource_CustomEvent
func callbackQGeoPositionInfoSource_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQGeoPositionInfoSourceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGeoPositionInfoSource) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQGeoPositionInfoSource_DeleteLater
func callbackQGeoPositionInfoSource_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQGeoPositionInfoSourceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QGeoPositionInfoSource) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGeoPositionInfoSource_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQGeoPositionInfoSource_Destroyed
func callbackQGeoPositionInfoSource_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQGeoPositionInfoSource_DisconnectNotify
func callbackQGeoPositionInfoSource_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGeoPositionInfoSourceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGeoPositionInfoSource_Event
func callbackQGeoPositionInfoSource_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGeoPositionInfoSourceFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QGeoPositionInfoSource) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGeoPositionInfoSource_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQGeoPositionInfoSource_EventFilter
func callbackQGeoPositionInfoSource_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGeoPositionInfoSourceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QGeoPositionInfoSource) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGeoPositionInfoSource_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQGeoPositionInfoSource_MetaObject
func callbackQGeoPositionInfoSource_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQGeoPositionInfoSourceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QGeoPositionInfoSource) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QGeoPositionInfoSource_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQGeoPositionInfoSource_ObjectNameChanged
func callbackQGeoPositionInfoSource_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtPositioning_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQGeoPositionInfoSource_TimerEvent
func callbackQGeoPositionInfoSource_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGeoPositionInfoSourceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGeoPositionInfoSource) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QGeoPositionInfoSourceFactory struct {
	ptr unsafe.Pointer
}

type QGeoPositionInfoSourceFactory_ITF interface {
	QGeoPositionInfoSourceFactory_PTR() *QGeoPositionInfoSourceFactory
}

func (ptr *QGeoPositionInfoSourceFactory) QGeoPositionInfoSourceFactory_PTR() *QGeoPositionInfoSourceFactory {
	return ptr
}

func (ptr *QGeoPositionInfoSourceFactory) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QGeoPositionInfoSourceFactory) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQGeoPositionInfoSourceFactory(ptr QGeoPositionInfoSourceFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoPositionInfoSourceFactory_PTR().Pointer()
	}
	return nil
}

func NewQGeoPositionInfoSourceFactoryFromPointer(ptr unsafe.Pointer) (n *QGeoPositionInfoSourceFactory) {
	n = new(QGeoPositionInfoSourceFactory)
	n.SetPointer(ptr)
	return
}

//export callbackQGeoPositionInfoSourceFactory_AreaMonitor
func callbackQGeoPositionInfoSourceFactory_AreaMonitor(ptr unsafe.Pointer, parent unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "areaMonitor"); signal != nil {
		return PointerFromQGeoAreaMonitorSource((*(*func(*core.QObject) *QGeoAreaMonitorSource)(signal))(core.NewQObjectFromPointer(parent)))
	}

	return PointerFromQGeoAreaMonitorSource(NewQGeoAreaMonitorSource(nil))
}

func (ptr *QGeoPositionInfoSourceFactory) ConnectAreaMonitor(f func(parent *core.QObject) *QGeoAreaMonitorSource) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "areaMonitor"); signal != nil {
			f := func(parent *core.QObject) *QGeoAreaMonitorSource {
				(*(*func(*core.QObject) *QGeoAreaMonitorSource)(signal))(parent)
				return f(parent)
			}
			qt.ConnectSignal(ptr.Pointer(), "areaMonitor", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "areaMonitor", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoPositionInfoSourceFactory) DisconnectAreaMonitor() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "areaMonitor")
	}
}

func (ptr *QGeoPositionInfoSourceFactory) AreaMonitor(parent core.QObject_ITF) *QGeoAreaMonitorSource {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoAreaMonitorSourceFromPointer(C.QGeoPositionInfoSourceFactory_AreaMonitor(ptr.Pointer(), core.PointerFromQObject(parent)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQGeoPositionInfoSourceFactory_PositionInfoSource
func callbackQGeoPositionInfoSourceFactory_PositionInfoSource(ptr unsafe.Pointer, parent unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "positionInfoSource"); signal != nil {
		return PointerFromQGeoPositionInfoSource((*(*func(*core.QObject) *QGeoPositionInfoSource)(signal))(core.NewQObjectFromPointer(parent)))
	}

	return PointerFromQGeoPositionInfoSource(NewQGeoPositionInfoSource(nil))
}

func (ptr *QGeoPositionInfoSourceFactory) ConnectPositionInfoSource(f func(parent *core.QObject) *QGeoPositionInfoSource) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "positionInfoSource"); signal != nil {
			f := func(parent *core.QObject) *QGeoPositionInfoSource {
				(*(*func(*core.QObject) *QGeoPositionInfoSource)(signal))(parent)
				return f(parent)
			}
			qt.ConnectSignal(ptr.Pointer(), "positionInfoSource", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "positionInfoSource", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoPositionInfoSourceFactory) DisconnectPositionInfoSource() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "positionInfoSource")
	}
}

func (ptr *QGeoPositionInfoSourceFactory) PositionInfoSource(parent core.QObject_ITF) *QGeoPositionInfoSource {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoPositionInfoSourceFromPointer(C.QGeoPositionInfoSourceFactory_PositionInfoSource(ptr.Pointer(), core.PointerFromQObject(parent)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQGeoPositionInfoSourceFactory_SatelliteInfoSource
func callbackQGeoPositionInfoSourceFactory_SatelliteInfoSource(ptr unsafe.Pointer, parent unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "satelliteInfoSource"); signal != nil {
		return PointerFromQGeoSatelliteInfoSource((*(*func(*core.QObject) *QGeoSatelliteInfoSource)(signal))(core.NewQObjectFromPointer(parent)))
	}

	return PointerFromQGeoSatelliteInfoSource(NewQGeoSatelliteInfoSource(nil))
}

func (ptr *QGeoPositionInfoSourceFactory) ConnectSatelliteInfoSource(f func(parent *core.QObject) *QGeoSatelliteInfoSource) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "satelliteInfoSource"); signal != nil {
			f := func(parent *core.QObject) *QGeoSatelliteInfoSource {
				(*(*func(*core.QObject) *QGeoSatelliteInfoSource)(signal))(parent)
				return f(parent)
			}
			qt.ConnectSignal(ptr.Pointer(), "satelliteInfoSource", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "satelliteInfoSource", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoPositionInfoSourceFactory) DisconnectSatelliteInfoSource() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "satelliteInfoSource")
	}
}

func (ptr *QGeoPositionInfoSourceFactory) SatelliteInfoSource(parent core.QObject_ITF) *QGeoSatelliteInfoSource {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoSatelliteInfoSourceFromPointer(C.QGeoPositionInfoSourceFactory_SatelliteInfoSource(ptr.Pointer(), core.PointerFromQObject(parent)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQGeoPositionInfoSourceFactory_DestroyQGeoPositionInfoSourceFactory
func callbackQGeoPositionInfoSourceFactory_DestroyQGeoPositionInfoSourceFactory(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QGeoPositionInfoSourceFactory"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQGeoPositionInfoSourceFactoryFromPointer(ptr).DestroyQGeoPositionInfoSourceFactoryDefault()
	}
}

func (ptr *QGeoPositionInfoSourceFactory) ConnectDestroyQGeoPositionInfoSourceFactory(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QGeoPositionInfoSourceFactory"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QGeoPositionInfoSourceFactory", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QGeoPositionInfoSourceFactory", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoPositionInfoSourceFactory) DisconnectDestroyQGeoPositionInfoSourceFactory() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QGeoPositionInfoSourceFactory")
	}
}

func (ptr *QGeoPositionInfoSourceFactory) DestroyQGeoPositionInfoSourceFactory() {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSourceFactory_DestroyQGeoPositionInfoSourceFactory(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGeoPositionInfoSourceFactory) DestroyQGeoPositionInfoSourceFactoryDefault() {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSourceFactory_DestroyQGeoPositionInfoSourceFactoryDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
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

func NewQGeoRectangleFromPointer(ptr unsafe.Pointer) (n *QGeoRectangle) {
	n = new(QGeoRectangle)
	n.SetPointer(ptr)
	return
}
func NewQGeoRectangle() *QGeoRectangle {
	tmpValue := NewQGeoRectangleFromPointer(C.QGeoRectangle_NewQGeoRectangle())
	qt.SetFinalizer(tmpValue, (*QGeoRectangle).DestroyQGeoRectangle)
	return tmpValue
}

func NewQGeoRectangle2(center QGeoCoordinate_ITF, degreesWidth float64, degreesHeight float64) *QGeoRectangle {
	tmpValue := NewQGeoRectangleFromPointer(C.QGeoRectangle_NewQGeoRectangle2(PointerFromQGeoCoordinate(center), C.double(degreesWidth), C.double(degreesHeight)))
	qt.SetFinalizer(tmpValue, (*QGeoRectangle).DestroyQGeoRectangle)
	return tmpValue
}

func NewQGeoRectangle3(topLeft QGeoCoordinate_ITF, bottomRight QGeoCoordinate_ITF) *QGeoRectangle {
	tmpValue := NewQGeoRectangleFromPointer(C.QGeoRectangle_NewQGeoRectangle3(PointerFromQGeoCoordinate(topLeft), PointerFromQGeoCoordinate(bottomRight)))
	qt.SetFinalizer(tmpValue, (*QGeoRectangle).DestroyQGeoRectangle)
	return tmpValue
}

func NewQGeoRectangle4(coordinates []*QGeoCoordinate) *QGeoRectangle {
	tmpValue := NewQGeoRectangleFromPointer(C.QGeoRectangle_NewQGeoRectangle4(func() unsafe.Pointer {
		tmpList := NewQGeoRectangleFromPointer(NewQGeoRectangleFromPointer(nil).__QGeoRectangle_coordinates_newList4())
		for _, v := range coordinates {
			tmpList.__QGeoRectangle_coordinates_setList4(v)
		}
		return tmpList.Pointer()
	}()))
	qt.SetFinalizer(tmpValue, (*QGeoRectangle).DestroyQGeoRectangle)
	return tmpValue
}

func NewQGeoRectangle5(other QGeoRectangle_ITF) *QGeoRectangle {
	tmpValue := NewQGeoRectangleFromPointer(C.QGeoRectangle_NewQGeoRectangle5(PointerFromQGeoRectangle(other)))
	qt.SetFinalizer(tmpValue, (*QGeoRectangle).DestroyQGeoRectangle)
	return tmpValue
}

func NewQGeoRectangle6(other QGeoShape_ITF) *QGeoRectangle {
	tmpValue := NewQGeoRectangleFromPointer(C.QGeoRectangle_NewQGeoRectangle6(PointerFromQGeoShape(other)))
	qt.SetFinalizer(tmpValue, (*QGeoRectangle).DestroyQGeoRectangle)
	return tmpValue
}

func (ptr *QGeoRectangle) BottomLeft() *QGeoCoordinate {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoCoordinateFromPointer(C.QGeoRectangle_BottomLeft(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QGeoCoordinate).DestroyQGeoCoordinate)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRectangle) BottomRight() *QGeoCoordinate {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoCoordinateFromPointer(C.QGeoRectangle_BottomRight(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QGeoCoordinate).DestroyQGeoCoordinate)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRectangle) Contains(rectangle QGeoRectangle_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGeoRectangle_Contains(ptr.Pointer(), PointerFromQGeoRectangle(rectangle))) != 0
	}
	return false
}

func (ptr *QGeoRectangle) ExtendRectangle(coordinate QGeoCoordinate_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRectangle_ExtendRectangle(ptr.Pointer(), PointerFromQGeoCoordinate(coordinate))
	}
}

func (ptr *QGeoRectangle) Height() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGeoRectangle_Height(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRectangle) Intersects(rectangle QGeoRectangle_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGeoRectangle_Intersects(ptr.Pointer(), PointerFromQGeoRectangle(rectangle))) != 0
	}
	return false
}

func (ptr *QGeoRectangle) SetBottomLeft(bottomLeft QGeoCoordinate_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRectangle_SetBottomLeft(ptr.Pointer(), PointerFromQGeoCoordinate(bottomLeft))
	}
}

func (ptr *QGeoRectangle) SetBottomRight(bottomRight QGeoCoordinate_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRectangle_SetBottomRight(ptr.Pointer(), PointerFromQGeoCoordinate(bottomRight))
	}
}

func (ptr *QGeoRectangle) SetCenter(center QGeoCoordinate_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRectangle_SetCenter(ptr.Pointer(), PointerFromQGeoCoordinate(center))
	}
}

func (ptr *QGeoRectangle) SetHeight(degreesHeight float64) {
	if ptr.Pointer() != nil {
		C.QGeoRectangle_SetHeight(ptr.Pointer(), C.double(degreesHeight))
	}
}

func (ptr *QGeoRectangle) SetTopLeft(topLeft QGeoCoordinate_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRectangle_SetTopLeft(ptr.Pointer(), PointerFromQGeoCoordinate(topLeft))
	}
}

func (ptr *QGeoRectangle) SetTopRight(topRight QGeoCoordinate_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRectangle_SetTopRight(ptr.Pointer(), PointerFromQGeoCoordinate(topRight))
	}
}

func (ptr *QGeoRectangle) SetWidth(degreesWidth float64) {
	if ptr.Pointer() != nil {
		C.QGeoRectangle_SetWidth(ptr.Pointer(), C.double(degreesWidth))
	}
}

func (ptr *QGeoRectangle) TopLeft() *QGeoCoordinate {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoCoordinateFromPointer(C.QGeoRectangle_TopLeft(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QGeoCoordinate).DestroyQGeoCoordinate)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRectangle) TopRight() *QGeoCoordinate {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoCoordinateFromPointer(C.QGeoRectangle_TopRight(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QGeoCoordinate).DestroyQGeoCoordinate)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRectangle) Translate(degreesLatitude float64, degreesLongitude float64) {
	if ptr.Pointer() != nil {
		C.QGeoRectangle_Translate(ptr.Pointer(), C.double(degreesLatitude), C.double(degreesLongitude))
	}
}

func (ptr *QGeoRectangle) Translated(degreesLatitude float64, degreesLongitude float64) *QGeoRectangle {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoRectangleFromPointer(C.QGeoRectangle_Translated(ptr.Pointer(), C.double(degreesLatitude), C.double(degreesLongitude)))
		qt.SetFinalizer(tmpValue, (*QGeoRectangle).DestroyQGeoRectangle)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRectangle) United(rectangle QGeoRectangle_ITF) *QGeoRectangle {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoRectangleFromPointer(C.QGeoRectangle_United(ptr.Pointer(), PointerFromQGeoRectangle(rectangle)))
		qt.SetFinalizer(tmpValue, (*QGeoRectangle).DestroyQGeoRectangle)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRectangle) Width() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGeoRectangle_Width(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRectangle) DestroyQGeoRectangle() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGeoRectangle_DestroyQGeoRectangle(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGeoRectangle) __QGeoRectangle_coordinates_atList4(i int) *QGeoCoordinate {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoCoordinateFromPointer(C.QGeoRectangle___QGeoRectangle_coordinates_atList4(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QGeoCoordinate).DestroyQGeoCoordinate)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRectangle) __QGeoRectangle_coordinates_setList4(i QGeoCoordinate_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRectangle___QGeoRectangle_coordinates_setList4(ptr.Pointer(), PointerFromQGeoCoordinate(i))
	}
}

func (ptr *QGeoRectangle) __QGeoRectangle_coordinates_newList4() unsafe.Pointer {
	return C.QGeoRectangle___QGeoRectangle_coordinates_newList4(ptr.Pointer())
}

type QGeoSatelliteInfo struct {
	ptr unsafe.Pointer
}

type QGeoSatelliteInfo_ITF interface {
	QGeoSatelliteInfo_PTR() *QGeoSatelliteInfo
}

func (ptr *QGeoSatelliteInfo) QGeoSatelliteInfo_PTR() *QGeoSatelliteInfo {
	return ptr
}

func (ptr *QGeoSatelliteInfo) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QGeoSatelliteInfo) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQGeoSatelliteInfo(ptr QGeoSatelliteInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoSatelliteInfo_PTR().Pointer()
	}
	return nil
}

func NewQGeoSatelliteInfoFromPointer(ptr unsafe.Pointer) (n *QGeoSatelliteInfo) {
	n = new(QGeoSatelliteInfo)
	n.SetPointer(ptr)
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
	tmpValue := NewQGeoSatelliteInfoFromPointer(C.QGeoSatelliteInfo_NewQGeoSatelliteInfo())
	qt.SetFinalizer(tmpValue, (*QGeoSatelliteInfo).DestroyQGeoSatelliteInfo)
	return tmpValue
}

func NewQGeoSatelliteInfo2(other QGeoSatelliteInfo_ITF) *QGeoSatelliteInfo {
	tmpValue := NewQGeoSatelliteInfoFromPointer(C.QGeoSatelliteInfo_NewQGeoSatelliteInfo2(PointerFromQGeoSatelliteInfo(other)))
	qt.SetFinalizer(tmpValue, (*QGeoSatelliteInfo).DestroyQGeoSatelliteInfo)
	return tmpValue
}

func (ptr *QGeoSatelliteInfo) Attribute(attribute QGeoSatelliteInfo__Attribute) float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGeoSatelliteInfo_Attribute(ptr.Pointer(), C.longlong(attribute)))
	}
	return 0
}

func (ptr *QGeoSatelliteInfo) HasAttribute(attribute QGeoSatelliteInfo__Attribute) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGeoSatelliteInfo_HasAttribute(ptr.Pointer(), C.longlong(attribute))) != 0
	}
	return false
}

func (ptr *QGeoSatelliteInfo) RemoveAttribute(attribute QGeoSatelliteInfo__Attribute) {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfo_RemoveAttribute(ptr.Pointer(), C.longlong(attribute))
	}
}

func (ptr *QGeoSatelliteInfo) SatelliteIdentifier() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QGeoSatelliteInfo_SatelliteIdentifier(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoSatelliteInfo) SatelliteSystem() QGeoSatelliteInfo__SatelliteSystem {
	if ptr.Pointer() != nil {
		return QGeoSatelliteInfo__SatelliteSystem(C.QGeoSatelliteInfo_SatelliteSystem(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoSatelliteInfo) SetAttribute(attribute QGeoSatelliteInfo__Attribute, value float64) {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfo_SetAttribute(ptr.Pointer(), C.longlong(attribute), C.double(value))
	}
}

func (ptr *QGeoSatelliteInfo) SetSatelliteIdentifier(satId int) {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfo_SetSatelliteIdentifier(ptr.Pointer(), C.int(int32(satId)))
	}
}

func (ptr *QGeoSatelliteInfo) SetSatelliteSystem(system QGeoSatelliteInfo__SatelliteSystem) {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfo_SetSatelliteSystem(ptr.Pointer(), C.longlong(system))
	}
}

func (ptr *QGeoSatelliteInfo) SetSignalStrength(signalStrength int) {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfo_SetSignalStrength(ptr.Pointer(), C.int(int32(signalStrength)))
	}
}

func (ptr *QGeoSatelliteInfo) SignalStrength() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QGeoSatelliteInfo_SignalStrength(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoSatelliteInfo) DestroyQGeoSatelliteInfo() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGeoSatelliteInfo_DestroyQGeoSatelliteInfo(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
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

func NewQGeoSatelliteInfoSourceFromPointer(ptr unsafe.Pointer) (n *QGeoSatelliteInfoSource) {
	n = new(QGeoSatelliteInfoSource)
	n.SetPointer(ptr)
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
	tmpValue := NewQGeoSatelliteInfoSourceFromPointer(C.QGeoSatelliteInfoSource_NewQGeoSatelliteInfoSource(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QGeoSatelliteInfoSource_AvailableSources() []string {
	return unpackStringList(cGoUnpackString(C.QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_AvailableSources()))
}

func (ptr *QGeoSatelliteInfoSource) AvailableSources() []string {
	return unpackStringList(cGoUnpackString(C.QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_AvailableSources()))
}

func QGeoSatelliteInfoSource_CreateDefaultSource(parent core.QObject_ITF) *QGeoSatelliteInfoSource {
	tmpValue := NewQGeoSatelliteInfoSourceFromPointer(C.QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_CreateDefaultSource(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QGeoSatelliteInfoSource) CreateDefaultSource(parent core.QObject_ITF) *QGeoSatelliteInfoSource {
	tmpValue := NewQGeoSatelliteInfoSourceFromPointer(C.QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_CreateDefaultSource(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QGeoSatelliteInfoSource_CreateSource(sourceName string, parent core.QObject_ITF) *QGeoSatelliteInfoSource {
	var sourceNameC *C.char
	if sourceName != "" {
		sourceNameC = C.CString(sourceName)
		defer C.free(unsafe.Pointer(sourceNameC))
	}
	tmpValue := NewQGeoSatelliteInfoSourceFromPointer(C.QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_CreateSource(C.struct_QtPositioning_PackedString{data: sourceNameC, len: C.longlong(len(sourceName))}, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QGeoSatelliteInfoSource) CreateSource(sourceName string, parent core.QObject_ITF) *QGeoSatelliteInfoSource {
	var sourceNameC *C.char
	if sourceName != "" {
		sourceNameC = C.CString(sourceName)
		defer C.free(unsafe.Pointer(sourceNameC))
	}
	tmpValue := NewQGeoSatelliteInfoSourceFromPointer(C.QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_CreateSource(C.struct_QtPositioning_PackedString{data: sourceNameC, len: C.longlong(len(sourceName))}, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQGeoSatelliteInfoSource_Error
func callbackQGeoSatelliteInfoSource_Error(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "error"); signal != nil {
		return C.longlong((*(*func() QGeoSatelliteInfoSource__Error)(signal))())
	}

	return C.longlong(0)
}

func (ptr *QGeoSatelliteInfoSource) ConnectError(f func() QGeoSatelliteInfoSource__Error) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "error"); signal != nil {
			f := func() QGeoSatelliteInfoSource__Error {
				(*(*func() QGeoSatelliteInfoSource__Error)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "error", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "error", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoSatelliteInfoSource) DisconnectError() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "error")
	}
}

func (ptr *QGeoSatelliteInfoSource) Error() QGeoSatelliteInfoSource__Error {
	if ptr.Pointer() != nil {
		return QGeoSatelliteInfoSource__Error(C.QGeoSatelliteInfoSource_Error(ptr.Pointer()))
	}
	return 0
}

//export callbackQGeoSatelliteInfoSource_Error2
func callbackQGeoSatelliteInfoSource_Error2(ptr unsafe.Pointer, satelliteError C.longlong) {
	if signal := qt.GetSignal(ptr, "error2"); signal != nil {
		(*(*func(QGeoSatelliteInfoSource__Error))(signal))(QGeoSatelliteInfoSource__Error(satelliteError))
	}

}

func (ptr *QGeoSatelliteInfoSource) ConnectError2(f func(satelliteError QGeoSatelliteInfoSource__Error)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "error2") {
			C.QGeoSatelliteInfoSource_ConnectError2(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "error")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "error2"); signal != nil {
			f := func(satelliteError QGeoSatelliteInfoSource__Error) {
				(*(*func(QGeoSatelliteInfoSource__Error))(signal))(satelliteError)
				f(satelliteError)
			}
			qt.ConnectSignal(ptr.Pointer(), "error2", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "error2", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoSatelliteInfoSource) DisconnectError2() {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "error2")
	}
}

func (ptr *QGeoSatelliteInfoSource) Error2(satelliteError QGeoSatelliteInfoSource__Error) {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_Error2(ptr.Pointer(), C.longlong(satelliteError))
	}
}

//export callbackQGeoSatelliteInfoSource_MinimumUpdateInterval
func callbackQGeoSatelliteInfoSource_MinimumUpdateInterval(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "minimumUpdateInterval"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(0))
}

func (ptr *QGeoSatelliteInfoSource) ConnectMinimumUpdateInterval(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "minimumUpdateInterval"); signal != nil {
			f := func() int {
				(*(*func() int)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "minimumUpdateInterval", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "minimumUpdateInterval", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoSatelliteInfoSource) DisconnectMinimumUpdateInterval() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "minimumUpdateInterval")
	}
}

func (ptr *QGeoSatelliteInfoSource) MinimumUpdateInterval() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QGeoSatelliteInfoSource_MinimumUpdateInterval(ptr.Pointer())))
	}
	return 0
}

//export callbackQGeoSatelliteInfoSource_RequestTimeout
func callbackQGeoSatelliteInfoSource_RequestTimeout(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "requestTimeout"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QGeoSatelliteInfoSource) ConnectRequestTimeout(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "requestTimeout") {
			C.QGeoSatelliteInfoSource_ConnectRequestTimeout(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "requestTimeout")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "requestTimeout"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "requestTimeout", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "requestTimeout", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoSatelliteInfoSource) DisconnectRequestTimeout() {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_DisconnectRequestTimeout(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "requestTimeout")
	}
}

func (ptr *QGeoSatelliteInfoSource) RequestTimeout() {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_RequestTimeout(ptr.Pointer())
	}
}

//export callbackQGeoSatelliteInfoSource_RequestUpdate
func callbackQGeoSatelliteInfoSource_RequestUpdate(ptr unsafe.Pointer, timeout C.int) {
	if signal := qt.GetSignal(ptr, "requestUpdate"); signal != nil {
		(*(*func(int))(signal))(int(int32(timeout)))
	}

}

func (ptr *QGeoSatelliteInfoSource) ConnectRequestUpdate(f func(timeout int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "requestUpdate"); signal != nil {
			f := func(timeout int) {
				(*(*func(int))(signal))(timeout)
				f(timeout)
			}
			qt.ConnectSignal(ptr.Pointer(), "requestUpdate", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "requestUpdate", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoSatelliteInfoSource) DisconnectRequestUpdate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "requestUpdate")
	}
}

func (ptr *QGeoSatelliteInfoSource) RequestUpdate(timeout int) {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_RequestUpdate(ptr.Pointer(), C.int(int32(timeout)))
	}
}

//export callbackQGeoSatelliteInfoSource_SatellitesInUseUpdated
func callbackQGeoSatelliteInfoSource_SatellitesInUseUpdated(ptr unsafe.Pointer, satellites C.struct_QtPositioning_PackedList) {
	if signal := qt.GetSignal(ptr, "satellitesInUseUpdated"); signal != nil {
		(*(*func([]*QGeoSatelliteInfo))(signal))(func(l C.struct_QtPositioning_PackedList) []*QGeoSatelliteInfo {
			out := make([]*QGeoSatelliteInfo, int(l.len))
			tmpList := NewQGeoSatelliteInfoSourceFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__satellitesInUseUpdated_satellites_atList(i)
			}
			return out
		}(satellites))
	}

}

func (ptr *QGeoSatelliteInfoSource) ConnectSatellitesInUseUpdated(f func(satellites []*QGeoSatelliteInfo)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "satellitesInUseUpdated") {
			C.QGeoSatelliteInfoSource_ConnectSatellitesInUseUpdated(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "satellitesInUseUpdated")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "satellitesInUseUpdated"); signal != nil {
			f := func(satellites []*QGeoSatelliteInfo) {
				(*(*func([]*QGeoSatelliteInfo))(signal))(satellites)
				f(satellites)
			}
			qt.ConnectSignal(ptr.Pointer(), "satellitesInUseUpdated", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "satellitesInUseUpdated", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoSatelliteInfoSource) DisconnectSatellitesInUseUpdated() {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_DisconnectSatellitesInUseUpdated(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "satellitesInUseUpdated")
	}
}

func (ptr *QGeoSatelliteInfoSource) SatellitesInUseUpdated(satellites []*QGeoSatelliteInfo) {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_SatellitesInUseUpdated(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQGeoSatelliteInfoSourceFromPointer(NewQGeoSatelliteInfoSourceFromPointer(nil).__satellitesInUseUpdated_satellites_newList())
			for _, v := range satellites {
				tmpList.__satellitesInUseUpdated_satellites_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackQGeoSatelliteInfoSource_SatellitesInViewUpdated
func callbackQGeoSatelliteInfoSource_SatellitesInViewUpdated(ptr unsafe.Pointer, satellites C.struct_QtPositioning_PackedList) {
	if signal := qt.GetSignal(ptr, "satellitesInViewUpdated"); signal != nil {
		(*(*func([]*QGeoSatelliteInfo))(signal))(func(l C.struct_QtPositioning_PackedList) []*QGeoSatelliteInfo {
			out := make([]*QGeoSatelliteInfo, int(l.len))
			tmpList := NewQGeoSatelliteInfoSourceFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__satellitesInViewUpdated_satellites_atList(i)
			}
			return out
		}(satellites))
	}

}

func (ptr *QGeoSatelliteInfoSource) ConnectSatellitesInViewUpdated(f func(satellites []*QGeoSatelliteInfo)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "satellitesInViewUpdated") {
			C.QGeoSatelliteInfoSource_ConnectSatellitesInViewUpdated(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "satellitesInViewUpdated")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "satellitesInViewUpdated"); signal != nil {
			f := func(satellites []*QGeoSatelliteInfo) {
				(*(*func([]*QGeoSatelliteInfo))(signal))(satellites)
				f(satellites)
			}
			qt.ConnectSignal(ptr.Pointer(), "satellitesInViewUpdated", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "satellitesInViewUpdated", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoSatelliteInfoSource) DisconnectSatellitesInViewUpdated() {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_DisconnectSatellitesInViewUpdated(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "satellitesInViewUpdated")
	}
}

func (ptr *QGeoSatelliteInfoSource) SatellitesInViewUpdated(satellites []*QGeoSatelliteInfo) {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_SatellitesInViewUpdated(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQGeoSatelliteInfoSourceFromPointer(NewQGeoSatelliteInfoSourceFromPointer(nil).__satellitesInViewUpdated_satellites_newList())
			for _, v := range satellites {
				tmpList.__satellitesInViewUpdated_satellites_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackQGeoSatelliteInfoSource_SetUpdateInterval
func callbackQGeoSatelliteInfoSource_SetUpdateInterval(ptr unsafe.Pointer, msec C.int) {
	if signal := qt.GetSignal(ptr, "setUpdateInterval"); signal != nil {
		(*(*func(int))(signal))(int(int32(msec)))
	} else {
		NewQGeoSatelliteInfoSourceFromPointer(ptr).SetUpdateIntervalDefault(int(int32(msec)))
	}
}

func (ptr *QGeoSatelliteInfoSource) ConnectSetUpdateInterval(f func(msec int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setUpdateInterval"); signal != nil {
			f := func(msec int) {
				(*(*func(int))(signal))(msec)
				f(msec)
			}
			qt.ConnectSignal(ptr.Pointer(), "setUpdateInterval", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setUpdateInterval", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoSatelliteInfoSource) DisconnectSetUpdateInterval() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setUpdateInterval")
	}
}

func (ptr *QGeoSatelliteInfoSource) SetUpdateInterval(msec int) {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_SetUpdateInterval(ptr.Pointer(), C.int(int32(msec)))
	}
}

func (ptr *QGeoSatelliteInfoSource) SetUpdateIntervalDefault(msec int) {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_SetUpdateIntervalDefault(ptr.Pointer(), C.int(int32(msec)))
	}
}

func (ptr *QGeoSatelliteInfoSource) SourceName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoSatelliteInfoSource_SourceName(ptr.Pointer()))
	}
	return ""
}

//export callbackQGeoSatelliteInfoSource_StartUpdates
func callbackQGeoSatelliteInfoSource_StartUpdates(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "startUpdates"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QGeoSatelliteInfoSource) ConnectStartUpdates(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "startUpdates"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "startUpdates", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "startUpdates", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoSatelliteInfoSource) DisconnectStartUpdates() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "startUpdates")
	}
}

func (ptr *QGeoSatelliteInfoSource) StartUpdates() {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_StartUpdates(ptr.Pointer())
	}
}

//export callbackQGeoSatelliteInfoSource_StopUpdates
func callbackQGeoSatelliteInfoSource_StopUpdates(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "stopUpdates"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QGeoSatelliteInfoSource) ConnectStopUpdates(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "stopUpdates"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "stopUpdates", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stopUpdates", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoSatelliteInfoSource) DisconnectStopUpdates() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "stopUpdates")
	}
}

func (ptr *QGeoSatelliteInfoSource) StopUpdates() {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_StopUpdates(ptr.Pointer())
	}
}

func (ptr *QGeoSatelliteInfoSource) UpdateInterval() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QGeoSatelliteInfoSource_UpdateInterval(ptr.Pointer())))
	}
	return 0
}

//export callbackQGeoSatelliteInfoSource_DestroyQGeoSatelliteInfoSource
func callbackQGeoSatelliteInfoSource_DestroyQGeoSatelliteInfoSource(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QGeoSatelliteInfoSource"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQGeoSatelliteInfoSourceFromPointer(ptr).DestroyQGeoSatelliteInfoSourceDefault()
	}
}

func (ptr *QGeoSatelliteInfoSource) ConnectDestroyQGeoSatelliteInfoSource(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QGeoSatelliteInfoSource"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QGeoSatelliteInfoSource", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QGeoSatelliteInfoSource", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGeoSatelliteInfoSource) DisconnectDestroyQGeoSatelliteInfoSource() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QGeoSatelliteInfoSource")
	}
}

func (ptr *QGeoSatelliteInfoSource) DestroyQGeoSatelliteInfoSource() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGeoSatelliteInfoSource_DestroyQGeoSatelliteInfoSource(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGeoSatelliteInfoSource) DestroyQGeoSatelliteInfoSourceDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGeoSatelliteInfoSource_DestroyQGeoSatelliteInfoSourceDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGeoSatelliteInfoSource) __satellitesInUseUpdated_satellites_atList(i int) *QGeoSatelliteInfo {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoSatelliteInfoFromPointer(C.QGeoSatelliteInfoSource___satellitesInUseUpdated_satellites_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QGeoSatelliteInfo).DestroyQGeoSatelliteInfo)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoSatelliteInfoSource) __satellitesInUseUpdated_satellites_setList(i QGeoSatelliteInfo_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource___satellitesInUseUpdated_satellites_setList(ptr.Pointer(), PointerFromQGeoSatelliteInfo(i))
	}
}

func (ptr *QGeoSatelliteInfoSource) __satellitesInUseUpdated_satellites_newList() unsafe.Pointer {
	return C.QGeoSatelliteInfoSource___satellitesInUseUpdated_satellites_newList(ptr.Pointer())
}

func (ptr *QGeoSatelliteInfoSource) __satellitesInViewUpdated_satellites_atList(i int) *QGeoSatelliteInfo {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoSatelliteInfoFromPointer(C.QGeoSatelliteInfoSource___satellitesInViewUpdated_satellites_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QGeoSatelliteInfo).DestroyQGeoSatelliteInfo)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoSatelliteInfoSource) __satellitesInViewUpdated_satellites_setList(i QGeoSatelliteInfo_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource___satellitesInViewUpdated_satellites_setList(ptr.Pointer(), PointerFromQGeoSatelliteInfo(i))
	}
}

func (ptr *QGeoSatelliteInfoSource) __satellitesInViewUpdated_satellites_newList() unsafe.Pointer {
	return C.QGeoSatelliteInfoSource___satellitesInViewUpdated_satellites_newList(ptr.Pointer())
}

func (ptr *QGeoSatelliteInfoSource) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGeoSatelliteInfoSource___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGeoSatelliteInfoSource) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGeoSatelliteInfoSource) __children_newList() unsafe.Pointer {
	return C.QGeoSatelliteInfoSource___children_newList(ptr.Pointer())
}

func (ptr *QGeoSatelliteInfoSource) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QGeoSatelliteInfoSource___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoSatelliteInfoSource) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QGeoSatelliteInfoSource) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QGeoSatelliteInfoSource___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QGeoSatelliteInfoSource) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGeoSatelliteInfoSource___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGeoSatelliteInfoSource) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGeoSatelliteInfoSource) __findChildren_newList() unsafe.Pointer {
	return C.QGeoSatelliteInfoSource___findChildren_newList(ptr.Pointer())
}

func (ptr *QGeoSatelliteInfoSource) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGeoSatelliteInfoSource___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGeoSatelliteInfoSource) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGeoSatelliteInfoSource) __findChildren_newList3() unsafe.Pointer {
	return C.QGeoSatelliteInfoSource___findChildren_newList3(ptr.Pointer())
}

//export callbackQGeoSatelliteInfoSource_ChildEvent
func callbackQGeoSatelliteInfoSource_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQGeoSatelliteInfoSourceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QGeoSatelliteInfoSource) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQGeoSatelliteInfoSource_ConnectNotify
func callbackQGeoSatelliteInfoSource_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGeoSatelliteInfoSourceFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGeoSatelliteInfoSource) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGeoSatelliteInfoSource_CustomEvent
func callbackQGeoSatelliteInfoSource_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQGeoSatelliteInfoSourceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGeoSatelliteInfoSource) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQGeoSatelliteInfoSource_DeleteLater
func callbackQGeoSatelliteInfoSource_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQGeoSatelliteInfoSourceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QGeoSatelliteInfoSource) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGeoSatelliteInfoSource_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQGeoSatelliteInfoSource_Destroyed
func callbackQGeoSatelliteInfoSource_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQGeoSatelliteInfoSource_DisconnectNotify
func callbackQGeoSatelliteInfoSource_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGeoSatelliteInfoSourceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGeoSatelliteInfoSource) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGeoSatelliteInfoSource_Event
func callbackQGeoSatelliteInfoSource_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGeoSatelliteInfoSourceFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QGeoSatelliteInfoSource) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGeoSatelliteInfoSource_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQGeoSatelliteInfoSource_EventFilter
func callbackQGeoSatelliteInfoSource_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGeoSatelliteInfoSourceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QGeoSatelliteInfoSource) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGeoSatelliteInfoSource_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQGeoSatelliteInfoSource_MetaObject
func callbackQGeoSatelliteInfoSource_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQGeoSatelliteInfoSourceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QGeoSatelliteInfoSource) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QGeoSatelliteInfoSource_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQGeoSatelliteInfoSource_ObjectNameChanged
func callbackQGeoSatelliteInfoSource_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtPositioning_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQGeoSatelliteInfoSource_TimerEvent
func callbackQGeoSatelliteInfoSource_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGeoSatelliteInfoSourceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGeoSatelliteInfoSource) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QGeoShape struct {
	ptr unsafe.Pointer
}

type QGeoShape_ITF interface {
	QGeoShape_PTR() *QGeoShape
}

func (ptr *QGeoShape) QGeoShape_PTR() *QGeoShape {
	return ptr
}

func (ptr *QGeoShape) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QGeoShape) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQGeoShape(ptr QGeoShape_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoShape_PTR().Pointer()
	}
	return nil
}

func NewQGeoShapeFromPointer(ptr unsafe.Pointer) (n *QGeoShape) {
	n = new(QGeoShape)
	n.SetPointer(ptr)
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
	tmpValue := NewQGeoShapeFromPointer(C.QGeoShape_NewQGeoShape())
	qt.SetFinalizer(tmpValue, (*QGeoShape).DestroyQGeoShape)
	return tmpValue
}

func NewQGeoShape2(other QGeoShape_ITF) *QGeoShape {
	tmpValue := NewQGeoShapeFromPointer(C.QGeoShape_NewQGeoShape2(PointerFromQGeoShape(other)))
	qt.SetFinalizer(tmpValue, (*QGeoShape).DestroyQGeoShape)
	return tmpValue
}

func (ptr *QGeoShape) BoundingGeoRectangle() *QGeoRectangle {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoRectangleFromPointer(C.QGeoShape_BoundingGeoRectangle(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QGeoRectangle).DestroyQGeoRectangle)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoShape) Center() *QGeoCoordinate {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoCoordinateFromPointer(C.QGeoShape_Center(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QGeoCoordinate).DestroyQGeoCoordinate)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoShape) Contains(coordinate QGeoCoordinate_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGeoShape_Contains(ptr.Pointer(), PointerFromQGeoCoordinate(coordinate))) != 0
	}
	return false
}

func (ptr *QGeoShape) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGeoShape_IsEmpty(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGeoShape) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGeoShape_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGeoShape) ToString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoShape_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoShape) Type() QGeoShape__ShapeType {
	if ptr.Pointer() != nil {
		return QGeoShape__ShapeType(C.QGeoShape_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoShape) DestroyQGeoShape() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGeoShape_DestroyQGeoShape(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
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

func NewQNmeaPositionInfoSourceFromPointer(ptr unsafe.Pointer) (n *QNmeaPositionInfoSource) {
	n = new(QNmeaPositionInfoSource)
	n.SetPointer(ptr)
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
	tmpValue := NewQNmeaPositionInfoSourceFromPointer(C.QNmeaPositionInfoSource_NewQNmeaPositionInfoSource(C.longlong(updateMode), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QNmeaPositionInfoSource) Device() *core.QIODevice {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQIODeviceFromPointer(C.QNmeaPositionInfoSource_Device(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQNmeaPositionInfoSource_Error
func callbackQNmeaPositionInfoSource_Error(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "error"); signal != nil {
		return C.longlong((*(*func() QGeoPositionInfoSource__Error)(signal))())
	}

	return C.longlong(NewQNmeaPositionInfoSourceFromPointer(ptr).ErrorDefault())
}

func (ptr *QNmeaPositionInfoSource) ConnectError(f func() QGeoPositionInfoSource__Error) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "error"); signal != nil {
			f := func() QGeoPositionInfoSource__Error {
				(*(*func() QGeoPositionInfoSource__Error)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "error", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "error", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNmeaPositionInfoSource) DisconnectError() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "error")
	}
}

func (ptr *QNmeaPositionInfoSource) Error() QGeoPositionInfoSource__Error {
	if ptr.Pointer() != nil {
		return QGeoPositionInfoSource__Error(C.QNmeaPositionInfoSource_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNmeaPositionInfoSource) ErrorDefault() QGeoPositionInfoSource__Error {
	if ptr.Pointer() != nil {
		return QGeoPositionInfoSource__Error(C.QNmeaPositionInfoSource_ErrorDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQNmeaPositionInfoSource_LastKnownPosition
func callbackQNmeaPositionInfoSource_LastKnownPosition(ptr unsafe.Pointer, fromSatellitePositioningMethodsOnly C.char) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "lastKnownPosition"); signal != nil {
		return PointerFromQGeoPositionInfo((*(*func(bool) *QGeoPositionInfo)(signal))(int8(fromSatellitePositioningMethodsOnly) != 0))
	}

	return PointerFromQGeoPositionInfo(NewQNmeaPositionInfoSourceFromPointer(ptr).LastKnownPositionDefault(int8(fromSatellitePositioningMethodsOnly) != 0))
}

func (ptr *QNmeaPositionInfoSource) ConnectLastKnownPosition(f func(fromSatellitePositioningMethodsOnly bool) *QGeoPositionInfo) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "lastKnownPosition"); signal != nil {
			f := func(fromSatellitePositioningMethodsOnly bool) *QGeoPositionInfo {
				(*(*func(bool) *QGeoPositionInfo)(signal))(fromSatellitePositioningMethodsOnly)
				return f(fromSatellitePositioningMethodsOnly)
			}
			qt.ConnectSignal(ptr.Pointer(), "lastKnownPosition", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "lastKnownPosition", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNmeaPositionInfoSource) DisconnectLastKnownPosition() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "lastKnownPosition")
	}
}

func (ptr *QNmeaPositionInfoSource) LastKnownPosition(fromSatellitePositioningMethodsOnly bool) *QGeoPositionInfo {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoPositionInfoFromPointer(C.QNmeaPositionInfoSource_LastKnownPosition(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(fromSatellitePositioningMethodsOnly)))))
		qt.SetFinalizer(tmpValue, (*QGeoPositionInfo).DestroyQGeoPositionInfo)
		return tmpValue
	}
	return nil
}

func (ptr *QNmeaPositionInfoSource) LastKnownPositionDefault(fromSatellitePositioningMethodsOnly bool) *QGeoPositionInfo {
	if ptr.Pointer() != nil {
		tmpValue := NewQGeoPositionInfoFromPointer(C.QNmeaPositionInfoSource_LastKnownPositionDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(fromSatellitePositioningMethodsOnly)))))
		qt.SetFinalizer(tmpValue, (*QGeoPositionInfo).DestroyQGeoPositionInfo)
		return tmpValue
	}
	return nil
}

//export callbackQNmeaPositionInfoSource_MinimumUpdateInterval
func callbackQNmeaPositionInfoSource_MinimumUpdateInterval(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "minimumUpdateInterval"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(NewQNmeaPositionInfoSourceFromPointer(ptr).MinimumUpdateIntervalDefault()))
}

func (ptr *QNmeaPositionInfoSource) ConnectMinimumUpdateInterval(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "minimumUpdateInterval"); signal != nil {
			f := func() int {
				(*(*func() int)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "minimumUpdateInterval", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "minimumUpdateInterval", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNmeaPositionInfoSource) DisconnectMinimumUpdateInterval() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "minimumUpdateInterval")
	}
}

func (ptr *QNmeaPositionInfoSource) MinimumUpdateInterval() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QNmeaPositionInfoSource_MinimumUpdateInterval(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNmeaPositionInfoSource) MinimumUpdateIntervalDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QNmeaPositionInfoSource_MinimumUpdateIntervalDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackQNmeaPositionInfoSource_ParsePosInfoFromNmeaData
func callbackQNmeaPositionInfoSource_ParsePosInfoFromNmeaData(ptr unsafe.Pointer, data C.struct_QtPositioning_PackedString, size C.int, posInfo unsafe.Pointer, hasFix *C.char) C.char {
	hasFixR := int8(*hasFix) != 0
	defer func() { *hasFix = C.char(int8(qt.GoBoolToInt(hasFixR))) }()
	if signal := qt.GetSignal(ptr, "parsePosInfoFromNmeaData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func([]byte, int, *QGeoPositionInfo, *bool) bool)(signal))(cGoUnpackBytes(data), int(int32(size)), NewQGeoPositionInfoFromPointer(posInfo), &hasFixR))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNmeaPositionInfoSourceFromPointer(ptr).ParsePosInfoFromNmeaDataDefault(cGoUnpackBytes(data), int(int32(size)), NewQGeoPositionInfoFromPointer(posInfo), &hasFixR))))
}

func (ptr *QNmeaPositionInfoSource) ConnectParsePosInfoFromNmeaData(f func(data []byte, size int, posInfo *QGeoPositionInfo, hasFix *bool) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "parsePosInfoFromNmeaData"); signal != nil {
			f := func(data []byte, size int, posInfo *QGeoPositionInfo, hasFix *bool) bool {
				(*(*func([]byte, int, *QGeoPositionInfo, *bool) bool)(signal))(data, size, posInfo, hasFix)
				return f(data, size, posInfo, hasFix)
			}
			qt.ConnectSignal(ptr.Pointer(), "parsePosInfoFromNmeaData", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "parsePosInfoFromNmeaData", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNmeaPositionInfoSource) DisconnectParsePosInfoFromNmeaData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "parsePosInfoFromNmeaData")
	}
}

func (ptr *QNmeaPositionInfoSource) ParsePosInfoFromNmeaData(data []byte, size int, posInfo QGeoPositionInfo_ITF, hasFix *bool) bool {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if len(data) != 0 {
			dataC = (*C.char)(unsafe.Pointer(&data[0]))
		}
		var hasFixC C.char
		if hasFix != nil {
			hasFixC = C.char(int8(qt.GoBoolToInt(*hasFix)))
			defer func() { *hasFix = int8(hasFixC) != 0 }()
		}
		return int8(C.QNmeaPositionInfoSource_ParsePosInfoFromNmeaData(ptr.Pointer(), dataC, C.int(int32(size)), PointerFromQGeoPositionInfo(posInfo), &hasFixC)) != 0
	}
	return false
}

func (ptr *QNmeaPositionInfoSource) ParsePosInfoFromNmeaDataDefault(data []byte, size int, posInfo QGeoPositionInfo_ITF, hasFix *bool) bool {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if len(data) != 0 {
			dataC = (*C.char)(unsafe.Pointer(&data[0]))
		}
		var hasFixC C.char
		if hasFix != nil {
			hasFixC = C.char(int8(qt.GoBoolToInt(*hasFix)))
			defer func() { *hasFix = int8(hasFixC) != 0 }()
		}
		return int8(C.QNmeaPositionInfoSource_ParsePosInfoFromNmeaDataDefault(ptr.Pointer(), dataC, C.int(int32(size)), PointerFromQGeoPositionInfo(posInfo), &hasFixC)) != 0
	}
	return false
}

//export callbackQNmeaPositionInfoSource_RequestUpdate
func callbackQNmeaPositionInfoSource_RequestUpdate(ptr unsafe.Pointer, msec C.int) {
	if signal := qt.GetSignal(ptr, "requestUpdate"); signal != nil {
		(*(*func(int))(signal))(int(int32(msec)))
	} else {
		NewQNmeaPositionInfoSourceFromPointer(ptr).RequestUpdateDefault(int(int32(msec)))
	}
}

func (ptr *QNmeaPositionInfoSource) ConnectRequestUpdate(f func(msec int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "requestUpdate"); signal != nil {
			f := func(msec int) {
				(*(*func(int))(signal))(msec)
				f(msec)
			}
			qt.ConnectSignal(ptr.Pointer(), "requestUpdate", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "requestUpdate", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNmeaPositionInfoSource) DisconnectRequestUpdate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "requestUpdate")
	}
}

func (ptr *QNmeaPositionInfoSource) RequestUpdate(msec int) {
	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_RequestUpdate(ptr.Pointer(), C.int(int32(msec)))
	}
}

func (ptr *QNmeaPositionInfoSource) RequestUpdateDefault(msec int) {
	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_RequestUpdateDefault(ptr.Pointer(), C.int(int32(msec)))
	}
}

func (ptr *QNmeaPositionInfoSource) SetDevice(device core.QIODevice_ITF) {
	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_SetDevice(ptr.Pointer(), core.PointerFromQIODevice(device))
	}
}

func (ptr *QNmeaPositionInfoSource) SetUserEquivalentRangeError(uere float64) {
	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_SetUserEquivalentRangeError(ptr.Pointer(), C.double(uere))
	}
}

//export callbackQNmeaPositionInfoSource_StartUpdates
func callbackQNmeaPositionInfoSource_StartUpdates(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "startUpdates"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQNmeaPositionInfoSourceFromPointer(ptr).StartUpdatesDefault()
	}
}

func (ptr *QNmeaPositionInfoSource) ConnectStartUpdates(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "startUpdates"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "startUpdates", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "startUpdates", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNmeaPositionInfoSource) DisconnectStartUpdates() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "startUpdates")
	}
}

func (ptr *QNmeaPositionInfoSource) StartUpdates() {
	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_StartUpdates(ptr.Pointer())
	}
}

func (ptr *QNmeaPositionInfoSource) StartUpdatesDefault() {
	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_StartUpdatesDefault(ptr.Pointer())
	}
}

//export callbackQNmeaPositionInfoSource_StopUpdates
func callbackQNmeaPositionInfoSource_StopUpdates(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "stopUpdates"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQNmeaPositionInfoSourceFromPointer(ptr).StopUpdatesDefault()
	}
}

func (ptr *QNmeaPositionInfoSource) ConnectStopUpdates(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "stopUpdates"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "stopUpdates", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stopUpdates", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNmeaPositionInfoSource) DisconnectStopUpdates() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "stopUpdates")
	}
}

func (ptr *QNmeaPositionInfoSource) StopUpdates() {
	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_StopUpdates(ptr.Pointer())
	}
}

func (ptr *QNmeaPositionInfoSource) StopUpdatesDefault() {
	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_StopUpdatesDefault(ptr.Pointer())
	}
}

//export callbackQNmeaPositionInfoSource_SupportedPositioningMethods
func callbackQNmeaPositionInfoSource_SupportedPositioningMethods(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedPositioningMethods"); signal != nil {
		return C.longlong((*(*func() QGeoPositionInfoSource__PositioningMethod)(signal))())
	}

	return C.longlong(NewQNmeaPositionInfoSourceFromPointer(ptr).SupportedPositioningMethodsDefault())
}

func (ptr *QNmeaPositionInfoSource) ConnectSupportedPositioningMethods(f func() QGeoPositionInfoSource__PositioningMethod) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "supportedPositioningMethods"); signal != nil {
			f := func() QGeoPositionInfoSource__PositioningMethod {
				(*(*func() QGeoPositionInfoSource__PositioningMethod)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "supportedPositioningMethods", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "supportedPositioningMethods", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNmeaPositionInfoSource) DisconnectSupportedPositioningMethods() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "supportedPositioningMethods")
	}
}

func (ptr *QNmeaPositionInfoSource) SupportedPositioningMethods() QGeoPositionInfoSource__PositioningMethod {
	if ptr.Pointer() != nil {
		return QGeoPositionInfoSource__PositioningMethod(C.QNmeaPositionInfoSource_SupportedPositioningMethods(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNmeaPositionInfoSource) SupportedPositioningMethodsDefault() QGeoPositionInfoSource__PositioningMethod {
	if ptr.Pointer() != nil {
		return QGeoPositionInfoSource__PositioningMethod(C.QNmeaPositionInfoSource_SupportedPositioningMethodsDefault(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNmeaPositionInfoSource) UpdateMode() QNmeaPositionInfoSource__UpdateMode {
	if ptr.Pointer() != nil {
		return QNmeaPositionInfoSource__UpdateMode(C.QNmeaPositionInfoSource_UpdateMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNmeaPositionInfoSource) UserEquivalentRangeError() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QNmeaPositionInfoSource_UserEquivalentRangeError(ptr.Pointer()))
	}
	return 0
}

//export callbackQNmeaPositionInfoSource_DestroyQNmeaPositionInfoSource
func callbackQNmeaPositionInfoSource_DestroyQNmeaPositionInfoSource(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QNmeaPositionInfoSource"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQNmeaPositionInfoSourceFromPointer(ptr).DestroyQNmeaPositionInfoSourceDefault()
	}
}

func (ptr *QNmeaPositionInfoSource) ConnectDestroyQNmeaPositionInfoSource(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QNmeaPositionInfoSource"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QNmeaPositionInfoSource", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QNmeaPositionInfoSource", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QNmeaPositionInfoSource) DisconnectDestroyQNmeaPositionInfoSource() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QNmeaPositionInfoSource")
	}
}

func (ptr *QNmeaPositionInfoSource) DestroyQNmeaPositionInfoSource() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QNmeaPositionInfoSource_DestroyQNmeaPositionInfoSource(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNmeaPositionInfoSource) DestroyQNmeaPositionInfoSourceDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QNmeaPositionInfoSource_DestroyQNmeaPositionInfoSourceDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQNmeaPositionInfoSource_Error2
func callbackQNmeaPositionInfoSource_Error2(ptr unsafe.Pointer, positioningError C.longlong) {
	if signal := qt.GetSignal(ptr, "error2"); signal != nil {
		(*(*func(QGeoPositionInfoSource__Error))(signal))(QGeoPositionInfoSource__Error(positioningError))
	}

}

func init() {
	qt.ItfMap["positioning.QGeoAddress_ITF"] = QGeoAddress{}
	qt.FuncMap["positioning.NewQGeoAddress"] = NewQGeoAddress
	qt.FuncMap["positioning.NewQGeoAddress2"] = NewQGeoAddress2
	qt.ItfMap["positioning.QGeoAreaMonitorInfo_ITF"] = QGeoAreaMonitorInfo{}
	qt.FuncMap["positioning.NewQGeoAreaMonitorInfo"] = NewQGeoAreaMonitorInfo
	qt.FuncMap["positioning.NewQGeoAreaMonitorInfo2"] = NewQGeoAreaMonitorInfo2
	qt.ItfMap["positioning.QGeoAreaMonitorSource_ITF"] = QGeoAreaMonitorSource{}
	qt.FuncMap["positioning.NewQGeoAreaMonitorSource"] = NewQGeoAreaMonitorSource
	qt.FuncMap["positioning.QGeoAreaMonitorSource_AvailableSources"] = QGeoAreaMonitorSource_AvailableSources
	qt.FuncMap["positioning.QGeoAreaMonitorSource_CreateDefaultSource"] = QGeoAreaMonitorSource_CreateDefaultSource
	qt.FuncMap["positioning.QGeoAreaMonitorSource_CreateSource"] = QGeoAreaMonitorSource_CreateSource
	qt.EnumMap["positioning.QGeoAreaMonitorSource__AccessError"] = int64(QGeoAreaMonitorSource__AccessError)
	qt.EnumMap["positioning.QGeoAreaMonitorSource__InsufficientPositionInfo"] = int64(QGeoAreaMonitorSource__InsufficientPositionInfo)
	qt.EnumMap["positioning.QGeoAreaMonitorSource__UnknownSourceError"] = int64(QGeoAreaMonitorSource__UnknownSourceError)
	qt.EnumMap["positioning.QGeoAreaMonitorSource__NoError"] = int64(QGeoAreaMonitorSource__NoError)
	qt.EnumMap["positioning.QGeoAreaMonitorSource__PersistentAreaMonitorFeature"] = int64(QGeoAreaMonitorSource__PersistentAreaMonitorFeature)
	qt.EnumMap["positioning.QGeoAreaMonitorSource__AnyAreaMonitorFeature"] = int64(QGeoAreaMonitorSource__AnyAreaMonitorFeature)
	qt.ItfMap["positioning.QGeoCircle_ITF"] = QGeoCircle{}
	qt.FuncMap["positioning.NewQGeoCircle"] = NewQGeoCircle
	qt.FuncMap["positioning.NewQGeoCircle2"] = NewQGeoCircle2
	qt.FuncMap["positioning.NewQGeoCircle3"] = NewQGeoCircle3
	qt.FuncMap["positioning.NewQGeoCircle4"] = NewQGeoCircle4
	qt.ItfMap["positioning.QGeoCoordinate_ITF"] = QGeoCoordinate{}
	qt.FuncMap["positioning.NewQGeoCoordinate"] = NewQGeoCoordinate
	qt.FuncMap["positioning.NewQGeoCoordinate2"] = NewQGeoCoordinate2
	qt.FuncMap["positioning.NewQGeoCoordinate3"] = NewQGeoCoordinate3
	qt.FuncMap["positioning.NewQGeoCoordinate4"] = NewQGeoCoordinate4
	qt.EnumMap["positioning.QGeoCoordinate__InvalidCoordinate"] = int64(QGeoCoordinate__InvalidCoordinate)
	qt.EnumMap["positioning.QGeoCoordinate__Coordinate2D"] = int64(QGeoCoordinate__Coordinate2D)
	qt.EnumMap["positioning.QGeoCoordinate__Coordinate3D"] = int64(QGeoCoordinate__Coordinate3D)
	qt.EnumMap["positioning.QGeoCoordinate__Degrees"] = int64(QGeoCoordinate__Degrees)
	qt.EnumMap["positioning.QGeoCoordinate__DegreesWithHemisphere"] = int64(QGeoCoordinate__DegreesWithHemisphere)
	qt.EnumMap["positioning.QGeoCoordinate__DegreesMinutes"] = int64(QGeoCoordinate__DegreesMinutes)
	qt.EnumMap["positioning.QGeoCoordinate__DegreesMinutesWithHemisphere"] = int64(QGeoCoordinate__DegreesMinutesWithHemisphere)
	qt.EnumMap["positioning.QGeoCoordinate__DegreesMinutesSeconds"] = int64(QGeoCoordinate__DegreesMinutesSeconds)
	qt.EnumMap["positioning.QGeoCoordinate__DegreesMinutesSecondsWithHemisphere"] = int64(QGeoCoordinate__DegreesMinutesSecondsWithHemisphere)
	qt.ItfMap["positioning.QGeoPath_ITF"] = QGeoPath{}
	qt.FuncMap["positioning.NewQGeoPath"] = NewQGeoPath
	qt.FuncMap["positioning.NewQGeoPath2"] = NewQGeoPath2
	qt.FuncMap["positioning.NewQGeoPath3"] = NewQGeoPath3
	qt.FuncMap["positioning.NewQGeoPath4"] = NewQGeoPath4
	qt.ItfMap["positioning.QGeoPolygon_ITF"] = QGeoPolygon{}
	qt.FuncMap["positioning.NewQGeoPolygon"] = NewQGeoPolygon
	qt.FuncMap["positioning.NewQGeoPolygon2"] = NewQGeoPolygon2
	qt.FuncMap["positioning.NewQGeoPolygon3"] = NewQGeoPolygon3
	qt.FuncMap["positioning.NewQGeoPolygon4"] = NewQGeoPolygon4
	qt.ItfMap["positioning.QGeoPositionInfo_ITF"] = QGeoPositionInfo{}
	qt.FuncMap["positioning.NewQGeoPositionInfo"] = NewQGeoPositionInfo
	qt.FuncMap["positioning.NewQGeoPositionInfo2"] = NewQGeoPositionInfo2
	qt.FuncMap["positioning.NewQGeoPositionInfo3"] = NewQGeoPositionInfo3
	qt.EnumMap["positioning.QGeoPositionInfo__Direction"] = int64(QGeoPositionInfo__Direction)
	qt.EnumMap["positioning.QGeoPositionInfo__GroundSpeed"] = int64(QGeoPositionInfo__GroundSpeed)
	qt.EnumMap["positioning.QGeoPositionInfo__VerticalSpeed"] = int64(QGeoPositionInfo__VerticalSpeed)
	qt.EnumMap["positioning.QGeoPositionInfo__MagneticVariation"] = int64(QGeoPositionInfo__MagneticVariation)
	qt.EnumMap["positioning.QGeoPositionInfo__HorizontalAccuracy"] = int64(QGeoPositionInfo__HorizontalAccuracy)
	qt.EnumMap["positioning.QGeoPositionInfo__VerticalAccuracy"] = int64(QGeoPositionInfo__VerticalAccuracy)
	qt.ItfMap["positioning.QGeoPositionInfoSource_ITF"] = QGeoPositionInfoSource{}
	qt.FuncMap["positioning.NewQGeoPositionInfoSource"] = NewQGeoPositionInfoSource
	qt.FuncMap["positioning.QGeoPositionInfoSource_AvailableSources"] = QGeoPositionInfoSource_AvailableSources
	qt.FuncMap["positioning.QGeoPositionInfoSource_CreateDefaultSource"] = QGeoPositionInfoSource_CreateDefaultSource
	qt.FuncMap["positioning.QGeoPositionInfoSource_CreateSource"] = QGeoPositionInfoSource_CreateSource
	qt.EnumMap["positioning.QGeoPositionInfoSource__AccessError"] = int64(QGeoPositionInfoSource__AccessError)
	qt.EnumMap["positioning.QGeoPositionInfoSource__ClosedError"] = int64(QGeoPositionInfoSource__ClosedError)
	qt.EnumMap["positioning.QGeoPositionInfoSource__UnknownSourceError"] = int64(QGeoPositionInfoSource__UnknownSourceError)
	qt.EnumMap["positioning.QGeoPositionInfoSource__NoError"] = int64(QGeoPositionInfoSource__NoError)
	qt.EnumMap["positioning.QGeoPositionInfoSource__NoPositioningMethods"] = int64(QGeoPositionInfoSource__NoPositioningMethods)
	qt.EnumMap["positioning.QGeoPositionInfoSource__SatellitePositioningMethods"] = int64(QGeoPositionInfoSource__SatellitePositioningMethods)
	qt.EnumMap["positioning.QGeoPositionInfoSource__NonSatellitePositioningMethods"] = int64(QGeoPositionInfoSource__NonSatellitePositioningMethods)
	qt.EnumMap["positioning.QGeoPositionInfoSource__AllPositioningMethods"] = int64(QGeoPositionInfoSource__AllPositioningMethods)
	qt.ItfMap["positioning.QGeoPositionInfoSourceFactory_ITF"] = QGeoPositionInfoSourceFactory{}
	qt.ItfMap["positioning.QGeoRectangle_ITF"] = QGeoRectangle{}
	qt.FuncMap["positioning.NewQGeoRectangle"] = NewQGeoRectangle
	qt.FuncMap["positioning.NewQGeoRectangle2"] = NewQGeoRectangle2
	qt.FuncMap["positioning.NewQGeoRectangle3"] = NewQGeoRectangle3
	qt.FuncMap["positioning.NewQGeoRectangle4"] = NewQGeoRectangle4
	qt.FuncMap["positioning.NewQGeoRectangle5"] = NewQGeoRectangle5
	qt.FuncMap["positioning.NewQGeoRectangle6"] = NewQGeoRectangle6
	qt.ItfMap["positioning.QGeoSatelliteInfo_ITF"] = QGeoSatelliteInfo{}
	qt.FuncMap["positioning.NewQGeoSatelliteInfo"] = NewQGeoSatelliteInfo
	qt.FuncMap["positioning.NewQGeoSatelliteInfo2"] = NewQGeoSatelliteInfo2
	qt.EnumMap["positioning.QGeoSatelliteInfo__Elevation"] = int64(QGeoSatelliteInfo__Elevation)
	qt.EnumMap["positioning.QGeoSatelliteInfo__Azimuth"] = int64(QGeoSatelliteInfo__Azimuth)
	qt.EnumMap["positioning.QGeoSatelliteInfo__Undefined"] = int64(QGeoSatelliteInfo__Undefined)
	qt.EnumMap["positioning.QGeoSatelliteInfo__GPS"] = int64(QGeoSatelliteInfo__GPS)
	qt.EnumMap["positioning.QGeoSatelliteInfo__GLONASS"] = int64(QGeoSatelliteInfo__GLONASS)
	qt.ItfMap["positioning.QGeoSatelliteInfoSource_ITF"] = QGeoSatelliteInfoSource{}
	qt.FuncMap["positioning.NewQGeoSatelliteInfoSource"] = NewQGeoSatelliteInfoSource
	qt.FuncMap["positioning.QGeoSatelliteInfoSource_AvailableSources"] = QGeoSatelliteInfoSource_AvailableSources
	qt.FuncMap["positioning.QGeoSatelliteInfoSource_CreateDefaultSource"] = QGeoSatelliteInfoSource_CreateDefaultSource
	qt.FuncMap["positioning.QGeoSatelliteInfoSource_CreateSource"] = QGeoSatelliteInfoSource_CreateSource
	qt.EnumMap["positioning.QGeoSatelliteInfoSource__AccessError"] = int64(QGeoSatelliteInfoSource__AccessError)
	qt.EnumMap["positioning.QGeoSatelliteInfoSource__ClosedError"] = int64(QGeoSatelliteInfoSource__ClosedError)
	qt.EnumMap["positioning.QGeoSatelliteInfoSource__NoError"] = int64(QGeoSatelliteInfoSource__NoError)
	qt.EnumMap["positioning.QGeoSatelliteInfoSource__UnknownSourceError"] = int64(QGeoSatelliteInfoSource__UnknownSourceError)
	qt.ItfMap["positioning.QGeoShape_ITF"] = QGeoShape{}
	qt.FuncMap["positioning.NewQGeoShape"] = NewQGeoShape
	qt.FuncMap["positioning.NewQGeoShape2"] = NewQGeoShape2
	qt.EnumMap["positioning.QGeoShape__UnknownType"] = int64(QGeoShape__UnknownType)
	qt.EnumMap["positioning.QGeoShape__RectangleType"] = int64(QGeoShape__RectangleType)
	qt.EnumMap["positioning.QGeoShape__CircleType"] = int64(QGeoShape__CircleType)
	qt.EnumMap["positioning.QGeoShape__PathType"] = int64(QGeoShape__PathType)
	qt.EnumMap["positioning.QGeoShape__PolygonType"] = int64(QGeoShape__PolygonType)
	qt.ItfMap["positioning.QNmeaPositionInfoSource_ITF"] = QNmeaPositionInfoSource{}
	qt.FuncMap["positioning.NewQNmeaPositionInfoSource"] = NewQNmeaPositionInfoSource
	qt.EnumMap["positioning.QNmeaPositionInfoSource__RealTimeMode"] = int64(QNmeaPositionInfoSource__RealTimeMode)
	qt.EnumMap["positioning.QNmeaPositionInfoSource__SimulationMode"] = int64(QNmeaPositionInfoSource__SimulationMode)
}
