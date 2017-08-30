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
	"runtime"
	"strings"
	"unsafe"
)

func cGoUnpackString(s C.struct_QtPositioning_PackedString) string {
	if len := int(s.len); len == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
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

func NewQGeoAddressFromPointer(ptr unsafe.Pointer) *QGeoAddress {
	var n = new(QGeoAddress)
	n.SetPointer(ptr)
	return n
}
func NewQGeoAddress() *QGeoAddress {
	var tmpValue = NewQGeoAddressFromPointer(C.QGeoAddress_NewQGeoAddress())
	runtime.SetFinalizer(tmpValue, (*QGeoAddress).DestroyQGeoAddress)
	return tmpValue
}

func NewQGeoAddress2(other QGeoAddress_ITF) *QGeoAddress {
	var tmpValue = NewQGeoAddressFromPointer(C.QGeoAddress_NewQGeoAddress2(PointerFromQGeoAddress(other)))
	runtime.SetFinalizer(tmpValue, (*QGeoAddress).DestroyQGeoAddress)
	return tmpValue
}

func (ptr *QGeoAddress) Clear() {
	if ptr.Pointer() != nil {
		C.QGeoAddress_Clear(ptr.Pointer())
	}
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

func (ptr *QGeoAddress) DestroyQGeoAddress() {
	if ptr.Pointer() != nil {
		C.QGeoAddress_DestroyQGeoAddress(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QGeoAddress) City() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoAddress_City(ptr.Pointer()))
	}
	return ""
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

func (ptr *QGeoAddress) PostalCode() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoAddress_PostalCode(ptr.Pointer()))
	}
	return ""
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

func (ptr *QGeoAddress) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QGeoAddress_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGeoAddress) IsTextGenerated() bool {
	if ptr.Pointer() != nil {
		return C.QGeoAddress_IsTextGenerated(ptr.Pointer()) != 0
	}
	return false
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

func NewQGeoAreaMonitorInfoFromPointer(ptr unsafe.Pointer) *QGeoAreaMonitorInfo {
	var n = new(QGeoAreaMonitorInfo)
	n.SetPointer(ptr)
	return n
}
func NewQGeoAreaMonitorInfo2(other QGeoAreaMonitorInfo_ITF) *QGeoAreaMonitorInfo {
	var tmpValue = NewQGeoAreaMonitorInfoFromPointer(C.QGeoAreaMonitorInfo_NewQGeoAreaMonitorInfo2(PointerFromQGeoAreaMonitorInfo(other)))
	runtime.SetFinalizer(tmpValue, (*QGeoAreaMonitorInfo).DestroyQGeoAreaMonitorInfo)
	return tmpValue
}

func NewQGeoAreaMonitorInfo(name string) *QGeoAreaMonitorInfo {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	var tmpValue = NewQGeoAreaMonitorInfoFromPointer(C.QGeoAreaMonitorInfo_NewQGeoAreaMonitorInfo(C.struct_QtPositioning_PackedString{data: nameC, len: C.longlong(len(name))}))
	runtime.SetFinalizer(tmpValue, (*QGeoAreaMonitorInfo).DestroyQGeoAreaMonitorInfo)
	return tmpValue
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
			var tmpList = NewQGeoAreaMonitorInfoFromPointer(NewQGeoAreaMonitorInfoFromPointer(nil).__setNotificationParameters_parameters_newList())
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
		C.QGeoAreaMonitorInfo_DestroyQGeoAreaMonitorInfo(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QGeoAreaMonitorInfo) Expiration() *core.QDateTime {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQDateTimeFromPointer(C.QGeoAreaMonitorInfo_Expiration(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QDateTime).DestroyQDateTime)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoAreaMonitorInfo) Area() *QGeoShape {
	if ptr.Pointer() != nil {
		var tmpValue = NewQGeoShapeFromPointer(C.QGeoAreaMonitorInfo_Area(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QGeoShape).DestroyQGeoShape)
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

func (ptr *QGeoAreaMonitorInfo) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoAreaMonitorInfo_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoAreaMonitorInfo) NotificationParameters() map[string]*core.QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtPositioning_PackedList) map[string]*core.QVariant {
			var out = make(map[string]*core.QVariant, int(l.len))
			for _, i := range NewQGeoAreaMonitorInfoFromPointer(l.data).__notificationParameters_keyList() {
				out[i] = NewQGeoAreaMonitorInfoFromPointer(l.data).__notificationParameters_atList(i)
			}
			return out
		}(C.QGeoAreaMonitorInfo_NotificationParameters(ptr.Pointer()))
	}
	return make(map[string]*core.QVariant, 0)
}

func (ptr *QGeoAreaMonitorInfo) IsPersistent() bool {
	if ptr.Pointer() != nil {
		return C.QGeoAreaMonitorInfo_IsPersistent(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGeoAreaMonitorInfo) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QGeoAreaMonitorInfo_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGeoAreaMonitorInfo) __setNotificationParameters_parameters_atList(i string) *core.QVariant {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		var tmpValue = core.NewQVariantFromPointer(C.QGeoAreaMonitorInfo___setNotificationParameters_parameters_atList(ptr.Pointer(), C.struct_QtPositioning_PackedString{data: iC, len: C.longlong(len(i))}))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
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
	return unsafe.Pointer(C.QGeoAreaMonitorInfo___setNotificationParameters_parameters_newList(ptr.Pointer()))
}

func (ptr *QGeoAreaMonitorInfo) __setNotificationParameters_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtPositioning_PackedList) []string {
			var out = make([]string, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQGeoAreaMonitorInfoFromPointer(l.data).____setNotificationParameters_keyList_atList(i)
			}
			return out
		}(C.QGeoAreaMonitorInfo___setNotificationParameters_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QGeoAreaMonitorInfo) __notificationParameters_atList(i string) *core.QVariant {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		var tmpValue = core.NewQVariantFromPointer(C.QGeoAreaMonitorInfo___notificationParameters_atList(ptr.Pointer(), C.struct_QtPositioning_PackedString{data: iC, len: C.longlong(len(i))}))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
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
	return unsafe.Pointer(C.QGeoAreaMonitorInfo___notificationParameters_newList(ptr.Pointer()))
}

func (ptr *QGeoAreaMonitorInfo) __notificationParameters_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtPositioning_PackedList) []string {
			var out = make([]string, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQGeoAreaMonitorInfoFromPointer(l.data).____notificationParameters_keyList_atList(i)
			}
			return out
		}(C.QGeoAreaMonitorInfo___notificationParameters_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QGeoAreaMonitorInfo) ____setNotificationParameters_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoAreaMonitorInfo_____setNotificationParameters_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QGeoAreaMonitorInfo) ____setNotificationParameters_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QGeoAreaMonitorInfo_____setNotificationParameters_keyList_setList(ptr.Pointer(), C.struct_QtPositioning_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QGeoAreaMonitorInfo) ____setNotificationParameters_keyList_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QGeoAreaMonitorInfo_____setNotificationParameters_keyList_newList(ptr.Pointer()))
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
	return unsafe.Pointer(C.QGeoAreaMonitorInfo_____notificationParameters_keyList_newList(ptr.Pointer()))
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

func NewQGeoAreaMonitorSourceFromPointer(ptr unsafe.Pointer) *QGeoAreaMonitorSource {
	var n = new(QGeoAreaMonitorSource)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QGeoAreaMonitorSource__AreaMonitorFeature
//QGeoAreaMonitorSource::AreaMonitorFeature
type QGeoAreaMonitorSource__AreaMonitorFeature int64

const (
	QGeoAreaMonitorSource__PersistentAreaMonitorFeature QGeoAreaMonitorSource__AreaMonitorFeature = QGeoAreaMonitorSource__AreaMonitorFeature(0x00000001)
	QGeoAreaMonitorSource__AnyAreaMonitorFeature        QGeoAreaMonitorSource__AreaMonitorFeature = QGeoAreaMonitorSource__AreaMonitorFeature(0xffffffff)
)

//go:generate stringer -type=QGeoAreaMonitorSource__Error
//QGeoAreaMonitorSource::Error
type QGeoAreaMonitorSource__Error int64

const (
	QGeoAreaMonitorSource__AccessError              QGeoAreaMonitorSource__Error = QGeoAreaMonitorSource__Error(0)
	QGeoAreaMonitorSource__InsufficientPositionInfo QGeoAreaMonitorSource__Error = QGeoAreaMonitorSource__Error(1)
	QGeoAreaMonitorSource__UnknownSourceError       QGeoAreaMonitorSource__Error = QGeoAreaMonitorSource__Error(2)
	QGeoAreaMonitorSource__NoError                  QGeoAreaMonitorSource__Error = QGeoAreaMonitorSource__Error(3)
)

func QGeoAreaMonitorSource_CreateDefaultSource(parent core.QObject_ITF) *QGeoAreaMonitorSource {
	var tmpValue = NewQGeoAreaMonitorSourceFromPointer(C.QGeoAreaMonitorSource_QGeoAreaMonitorSource_CreateDefaultSource(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QGeoAreaMonitorSource) CreateDefaultSource(parent core.QObject_ITF) *QGeoAreaMonitorSource {
	var tmpValue = NewQGeoAreaMonitorSourceFromPointer(C.QGeoAreaMonitorSource_QGeoAreaMonitorSource_CreateDefaultSource(core.PointerFromQObject(parent)))
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
	var tmpValue = NewQGeoAreaMonitorSourceFromPointer(C.QGeoAreaMonitorSource_QGeoAreaMonitorSource_CreateSource(C.struct_QtPositioning_PackedString{data: sourceNameC, len: C.longlong(len(sourceName))}, core.PointerFromQObject(parent)))
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
	var tmpValue = NewQGeoAreaMonitorSourceFromPointer(C.QGeoAreaMonitorSource_QGeoAreaMonitorSource_CreateSource(C.struct_QtPositioning_PackedString{data: sourceNameC, len: C.longlong(len(sourceName))}, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQGeoAreaMonitorSource(parent core.QObject_ITF) *QGeoAreaMonitorSource {
	var tmpValue = NewQGeoAreaMonitorSourceFromPointer(C.QGeoAreaMonitorSource_NewQGeoAreaMonitorSource(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QGeoAreaMonitorSource_AvailableSources() []string {
	return strings.Split(cGoUnpackString(C.QGeoAreaMonitorSource_QGeoAreaMonitorSource_AvailableSources()), "|")
}

func (ptr *QGeoAreaMonitorSource) AvailableSources() []string {
	return strings.Split(cGoUnpackString(C.QGeoAreaMonitorSource_QGeoAreaMonitorSource_AvailableSources()), "|")
}

//export callbackQGeoAreaMonitorSource_RequestUpdate
func callbackQGeoAreaMonitorSource_RequestUpdate(ptr unsafe.Pointer, monitor unsafe.Pointer, sign C.struct_QtPositioning_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "requestUpdate"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QGeoAreaMonitorInfo, string) bool)(NewQGeoAreaMonitorInfoFromPointer(monitor), cGoUnpackString(sign)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QGeoAreaMonitorSource) ConnectRequestUpdate(f func(monitor *QGeoAreaMonitorInfo, sign string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "requestUpdate"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "requestUpdate", func(monitor *QGeoAreaMonitorInfo, sign string) bool {
				signal.(func(*QGeoAreaMonitorInfo, string) bool)(monitor, sign)
				return f(monitor, sign)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "requestUpdate", f)
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
		return C.QGeoAreaMonitorSource_RequestUpdate(ptr.Pointer(), PointerFromQGeoAreaMonitorInfo(monitor), signC) != 0
	}
	return false
}

//export callbackQGeoAreaMonitorSource_StartMonitoring
func callbackQGeoAreaMonitorSource_StartMonitoring(ptr unsafe.Pointer, monitor unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "startMonitoring"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QGeoAreaMonitorInfo) bool)(NewQGeoAreaMonitorInfoFromPointer(monitor)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QGeoAreaMonitorSource) ConnectStartMonitoring(f func(monitor *QGeoAreaMonitorInfo) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "startMonitoring"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "startMonitoring", func(monitor *QGeoAreaMonitorInfo) bool {
				signal.(func(*QGeoAreaMonitorInfo) bool)(monitor)
				return f(monitor)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "startMonitoring", f)
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
		return C.QGeoAreaMonitorSource_StartMonitoring(ptr.Pointer(), PointerFromQGeoAreaMonitorInfo(monitor)) != 0
	}
	return false
}

//export callbackQGeoAreaMonitorSource_StopMonitoring
func callbackQGeoAreaMonitorSource_StopMonitoring(ptr unsafe.Pointer, monitor unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "stopMonitoring"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QGeoAreaMonitorInfo) bool)(NewQGeoAreaMonitorInfoFromPointer(monitor)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QGeoAreaMonitorSource) ConnectStopMonitoring(f func(monitor *QGeoAreaMonitorInfo) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "stopMonitoring"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "stopMonitoring", func(monitor *QGeoAreaMonitorInfo) bool {
				signal.(func(*QGeoAreaMonitorInfo) bool)(monitor)
				return f(monitor)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stopMonitoring", f)
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
		return C.QGeoAreaMonitorSource_StopMonitoring(ptr.Pointer(), PointerFromQGeoAreaMonitorInfo(monitor)) != 0
	}
	return false
}

//export callbackQGeoAreaMonitorSource_AreaEntered
func callbackQGeoAreaMonitorSource_AreaEntered(ptr unsafe.Pointer, monitor unsafe.Pointer, update unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "areaEntered"); signal != nil {
		signal.(func(*QGeoAreaMonitorInfo, *QGeoPositionInfo))(NewQGeoAreaMonitorInfoFromPointer(monitor), NewQGeoPositionInfoFromPointer(update))
	}

}

func (ptr *QGeoAreaMonitorSource) ConnectAreaEntered(f func(monitor *QGeoAreaMonitorInfo, update *QGeoPositionInfo)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "areaEntered") {
			C.QGeoAreaMonitorSource_ConnectAreaEntered(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "areaEntered"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "areaEntered", func(monitor *QGeoAreaMonitorInfo, update *QGeoPositionInfo) {
				signal.(func(*QGeoAreaMonitorInfo, *QGeoPositionInfo))(monitor, update)
				f(monitor, update)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "areaEntered", f)
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
		signal.(func(*QGeoAreaMonitorInfo, *QGeoPositionInfo))(NewQGeoAreaMonitorInfoFromPointer(monitor), NewQGeoPositionInfoFromPointer(update))
	}

}

func (ptr *QGeoAreaMonitorSource) ConnectAreaExited(f func(monitor *QGeoAreaMonitorInfo, update *QGeoPositionInfo)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "areaExited") {
			C.QGeoAreaMonitorSource_ConnectAreaExited(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "areaExited"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "areaExited", func(monitor *QGeoAreaMonitorInfo, update *QGeoPositionInfo) {
				signal.(func(*QGeoAreaMonitorInfo, *QGeoPositionInfo))(monitor, update)
				f(monitor, update)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "areaExited", f)
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

//export callbackQGeoAreaMonitorSource_Error2
func callbackQGeoAreaMonitorSource_Error2(ptr unsafe.Pointer, areaMonitoringError C.longlong) {
	if signal := qt.GetSignal(ptr, "error2"); signal != nil {
		signal.(func(QGeoAreaMonitorSource__Error))(QGeoAreaMonitorSource__Error(areaMonitoringError))
	}

}

func (ptr *QGeoAreaMonitorSource) ConnectError2(f func(areaMonitoringError QGeoAreaMonitorSource__Error)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "error2") {
			C.QGeoAreaMonitorSource_ConnectError2(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "error2"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "error2", func(areaMonitoringError QGeoAreaMonitorSource__Error) {
				signal.(func(QGeoAreaMonitorSource__Error))(areaMonitoringError)
				f(areaMonitoringError)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "error2", f)
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
		signal.(func(*QGeoAreaMonitorInfo))(NewQGeoAreaMonitorInfoFromPointer(monitor))
	}

}

func (ptr *QGeoAreaMonitorSource) ConnectMonitorExpired(f func(monitor *QGeoAreaMonitorInfo)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "monitorExpired") {
			C.QGeoAreaMonitorSource_ConnectMonitorExpired(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "monitorExpired"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "monitorExpired", func(monitor *QGeoAreaMonitorInfo) {
				signal.(func(*QGeoAreaMonitorInfo))(monitor)
				f(monitor)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "monitorExpired", f)
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

//export callbackQGeoAreaMonitorSource_SetPositionInfoSource
func callbackQGeoAreaMonitorSource_SetPositionInfoSource(ptr unsafe.Pointer, newSource unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setPositionInfoSource"); signal != nil {
		signal.(func(*QGeoPositionInfoSource))(NewQGeoPositionInfoSourceFromPointer(newSource))
	} else {
		NewQGeoAreaMonitorSourceFromPointer(ptr).SetPositionInfoSourceDefault(NewQGeoPositionInfoSourceFromPointer(newSource))
	}
}

func (ptr *QGeoAreaMonitorSource) ConnectSetPositionInfoSource(f func(newSource *QGeoPositionInfoSource)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setPositionInfoSource"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setPositionInfoSource", func(newSource *QGeoPositionInfoSource) {
				signal.(func(*QGeoPositionInfoSource))(newSource)
				f(newSource)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setPositionInfoSource", f)
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

//export callbackQGeoAreaMonitorSource_DestroyQGeoAreaMonitorSource
func callbackQGeoAreaMonitorSource_DestroyQGeoAreaMonitorSource(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QGeoAreaMonitorSource"); signal != nil {
		signal.(func())()
	} else {
		NewQGeoAreaMonitorSourceFromPointer(ptr).DestroyQGeoAreaMonitorSourceDefault()
	}
}

func (ptr *QGeoAreaMonitorSource) ConnectDestroyQGeoAreaMonitorSource(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QGeoAreaMonitorSource"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QGeoAreaMonitorSource", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QGeoAreaMonitorSource", f)
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
		C.QGeoAreaMonitorSource_DestroyQGeoAreaMonitorSource(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QGeoAreaMonitorSource) DestroyQGeoAreaMonitorSourceDefault() {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_DestroyQGeoAreaMonitorSourceDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQGeoAreaMonitorSource_SupportedAreaMonitorFeatures
func callbackQGeoAreaMonitorSource_SupportedAreaMonitorFeatures(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedAreaMonitorFeatures"); signal != nil {
		return C.longlong(signal.(func() QGeoAreaMonitorSource__AreaMonitorFeature)())
	}

	return C.longlong(0)
}

func (ptr *QGeoAreaMonitorSource) ConnectSupportedAreaMonitorFeatures(f func() QGeoAreaMonitorSource__AreaMonitorFeature) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "supportedAreaMonitorFeatures"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "supportedAreaMonitorFeatures", func() QGeoAreaMonitorSource__AreaMonitorFeature {
				signal.(func() QGeoAreaMonitorSource__AreaMonitorFeature)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "supportedAreaMonitorFeatures", f)
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

//export callbackQGeoAreaMonitorSource_Error
func callbackQGeoAreaMonitorSource_Error(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "error"); signal != nil {
		return C.longlong(signal.(func() QGeoAreaMonitorSource__Error)())
	}

	return C.longlong(0)
}

func (ptr *QGeoAreaMonitorSource) ConnectError(f func() QGeoAreaMonitorSource__Error) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "error"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "error", func() QGeoAreaMonitorSource__Error {
				signal.(func() QGeoAreaMonitorSource__Error)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "error", f)
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

//export callbackQGeoAreaMonitorSource_PositionInfoSource
func callbackQGeoAreaMonitorSource_PositionInfoSource(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "positionInfoSource"); signal != nil {
		return PointerFromQGeoPositionInfoSource(signal.(func() *QGeoPositionInfoSource)())
	}

	return PointerFromQGeoPositionInfoSource(NewQGeoAreaMonitorSourceFromPointer(ptr).PositionInfoSourceDefault())
}

func (ptr *QGeoAreaMonitorSource) ConnectPositionInfoSource(f func() *QGeoPositionInfoSource) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "positionInfoSource"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "positionInfoSource", func() *QGeoPositionInfoSource {
				signal.(func() *QGeoPositionInfoSource)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "positionInfoSource", f)
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
		var tmpValue = NewQGeoPositionInfoSourceFromPointer(C.QGeoAreaMonitorSource_PositionInfoSource(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGeoAreaMonitorSource) PositionInfoSourceDefault() *QGeoPositionInfoSource {
	if ptr.Pointer() != nil {
		var tmpValue = NewQGeoPositionInfoSourceFromPointer(C.QGeoAreaMonitorSource_PositionInfoSourceDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQGeoAreaMonitorSource_ActiveMonitors
func callbackQGeoAreaMonitorSource_ActiveMonitors(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "activeMonitors"); signal != nil {
		return func() unsafe.Pointer {
			var tmpList = NewQGeoAreaMonitorSourceFromPointer(NewQGeoAreaMonitorSourceFromPointer(nil).__activeMonitors_newList())
			for _, v := range signal.(func() []*QGeoAreaMonitorInfo)() {
				tmpList.__activeMonitors_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		var tmpList = NewQGeoAreaMonitorSourceFromPointer(NewQGeoAreaMonitorSourceFromPointer(nil).__activeMonitors_newList())
		for _, v := range make([]*QGeoAreaMonitorInfo, 0) {
			tmpList.__activeMonitors_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *QGeoAreaMonitorSource) ConnectActiveMonitors(f func() []*QGeoAreaMonitorInfo) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "activeMonitors"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "activeMonitors", func() []*QGeoAreaMonitorInfo {
				signal.(func() []*QGeoAreaMonitorInfo)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "activeMonitors", f)
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
			var out = make([]*QGeoAreaMonitorInfo, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQGeoAreaMonitorSourceFromPointer(l.data).__activeMonitors_atList(i)
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
			var tmpList = NewQGeoAreaMonitorSourceFromPointer(NewQGeoAreaMonitorSourceFromPointer(nil).__activeMonitors_newList2())
			for _, v := range signal.(func(*QGeoShape) []*QGeoAreaMonitorInfo)(NewQGeoShapeFromPointer(lookupArea)) {
				tmpList.__activeMonitors_setList2(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		var tmpList = NewQGeoAreaMonitorSourceFromPointer(NewQGeoAreaMonitorSourceFromPointer(nil).__activeMonitors_newList2())
		for _, v := range make([]*QGeoAreaMonitorInfo, 0) {
			tmpList.__activeMonitors_setList2(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *QGeoAreaMonitorSource) ConnectActiveMonitors2(f func(lookupArea *QGeoShape) []*QGeoAreaMonitorInfo) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "activeMonitors2"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "activeMonitors2", func(lookupArea *QGeoShape) []*QGeoAreaMonitorInfo {
				signal.(func(*QGeoShape) []*QGeoAreaMonitorInfo)(lookupArea)
				return f(lookupArea)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "activeMonitors2", f)
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
			var out = make([]*QGeoAreaMonitorInfo, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQGeoAreaMonitorSourceFromPointer(l.data).__activeMonitors_atList2(i)
			}
			return out
		}(C.QGeoAreaMonitorSource_ActiveMonitors2(ptr.Pointer(), PointerFromQGeoShape(lookupArea)))
	}
	return make([]*QGeoAreaMonitorInfo, 0)
}

func (ptr *QGeoAreaMonitorSource) SourceName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoAreaMonitorSource_SourceName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoAreaMonitorSource) __activeMonitors_atList(i int) *QGeoAreaMonitorInfo {
	if ptr.Pointer() != nil {
		var tmpValue = NewQGeoAreaMonitorInfoFromPointer(C.QGeoAreaMonitorSource___activeMonitors_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QGeoAreaMonitorInfo).DestroyQGeoAreaMonitorInfo)
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
	return unsafe.Pointer(C.QGeoAreaMonitorSource___activeMonitors_newList(ptr.Pointer()))
}

func (ptr *QGeoAreaMonitorSource) __activeMonitors_atList2(i int) *QGeoAreaMonitorInfo {
	if ptr.Pointer() != nil {
		var tmpValue = NewQGeoAreaMonitorInfoFromPointer(C.QGeoAreaMonitorSource___activeMonitors_atList2(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QGeoAreaMonitorInfo).DestroyQGeoAreaMonitorInfo)
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
	return unsafe.Pointer(C.QGeoAreaMonitorSource___activeMonitors_newList2(ptr.Pointer()))
}

func (ptr *QGeoAreaMonitorSource) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QGeoAreaMonitorSource___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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
	return unsafe.Pointer(C.QGeoAreaMonitorSource___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QGeoAreaMonitorSource) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QGeoAreaMonitorSource___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGeoAreaMonitorSource) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGeoAreaMonitorSource) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QGeoAreaMonitorSource___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QGeoAreaMonitorSource) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QGeoAreaMonitorSource___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
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
	return unsafe.Pointer(C.QGeoAreaMonitorSource___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QGeoAreaMonitorSource) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QGeoAreaMonitorSource___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
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
	return unsafe.Pointer(C.QGeoAreaMonitorSource___findChildren_newList(ptr.Pointer()))
}

func (ptr *QGeoAreaMonitorSource) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QGeoAreaMonitorSource___children_atList(ptr.Pointer(), C.int(int32(i))))
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
	return unsafe.Pointer(C.QGeoAreaMonitorSource___children_newList(ptr.Pointer()))
}

//export callbackQGeoAreaMonitorSource_Event
func callbackQGeoAreaMonitorSource_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGeoAreaMonitorSourceFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QGeoAreaMonitorSource) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGeoAreaMonitorSource_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQGeoAreaMonitorSource_EventFilter
func callbackQGeoAreaMonitorSource_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGeoAreaMonitorSourceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QGeoAreaMonitorSource) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGeoAreaMonitorSource_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQGeoAreaMonitorSource_ChildEvent
func callbackQGeoAreaMonitorSource_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
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
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
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
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
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
		signal.(func())()
	} else {
		NewQGeoAreaMonitorSourceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QGeoAreaMonitorSource) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQGeoAreaMonitorSource_Destroyed
func callbackQGeoAreaMonitorSource_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQGeoAreaMonitorSource_DisconnectNotify
func callbackQGeoAreaMonitorSource_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGeoAreaMonitorSourceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGeoAreaMonitorSource) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGeoAreaMonitorSource_ObjectNameChanged
func callbackQGeoAreaMonitorSource_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtPositioning_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQGeoAreaMonitorSource_TimerEvent
func callbackQGeoAreaMonitorSource_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGeoAreaMonitorSourceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGeoAreaMonitorSource) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQGeoAreaMonitorSource_MetaObject
func callbackQGeoAreaMonitorSource_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQGeoAreaMonitorSourceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QGeoAreaMonitorSource) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QGeoAreaMonitorSource_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
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

func NewQGeoCircleFromPointer(ptr unsafe.Pointer) *QGeoCircle {
	var n = new(QGeoCircle)
	n.SetPointer(ptr)
	return n
}
func NewQGeoCircle() *QGeoCircle {
	var tmpValue = NewQGeoCircleFromPointer(C.QGeoCircle_NewQGeoCircle())
	runtime.SetFinalizer(tmpValue, (*QGeoCircle).DestroyQGeoCircle)
	return tmpValue
}

func NewQGeoCircle3(other QGeoCircle_ITF) *QGeoCircle {
	var tmpValue = NewQGeoCircleFromPointer(C.QGeoCircle_NewQGeoCircle3(PointerFromQGeoCircle(other)))
	runtime.SetFinalizer(tmpValue, (*QGeoCircle).DestroyQGeoCircle)
	return tmpValue
}

func NewQGeoCircle2(center QGeoCoordinate_ITF, radius float64) *QGeoCircle {
	var tmpValue = NewQGeoCircleFromPointer(C.QGeoCircle_NewQGeoCircle2(PointerFromQGeoCoordinate(center), C.double(radius)))
	runtime.SetFinalizer(tmpValue, (*QGeoCircle).DestroyQGeoCircle)
	return tmpValue
}

func NewQGeoCircle4(other QGeoShape_ITF) *QGeoCircle {
	var tmpValue = NewQGeoCircleFromPointer(C.QGeoCircle_NewQGeoCircle4(PointerFromQGeoShape(other)))
	runtime.SetFinalizer(tmpValue, (*QGeoCircle).DestroyQGeoCircle)
	return tmpValue
}

func (ptr *QGeoCircle) ExtendCircle(coordinate QGeoCoordinate_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoCircle_ExtendCircle(ptr.Pointer(), PointerFromQGeoCoordinate(coordinate))
	}
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

func (ptr *QGeoCircle) DestroyQGeoCircle() {
	if ptr.Pointer() != nil {
		C.QGeoCircle_DestroyQGeoCircle(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QGeoCircle) Translated(degreesLatitude float64, degreesLongitude float64) *QGeoCircle {
	if ptr.Pointer() != nil {
		var tmpValue = NewQGeoCircleFromPointer(C.QGeoCircle_Translated(ptr.Pointer(), C.double(degreesLatitude), C.double(degreesLongitude)))
		runtime.SetFinalizer(tmpValue, (*QGeoCircle).DestroyQGeoCircle)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoCircle) Radius() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGeoCircle_Radius(ptr.Pointer()))
	}
	return 0
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

func NewQGeoCoordinateFromPointer(ptr unsafe.Pointer) *QGeoCoordinate {
	var n = new(QGeoCoordinate)
	n.SetPointer(ptr)
	return n
}

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

//go:generate stringer -type=QGeoCoordinate__CoordinateType
//QGeoCoordinate::CoordinateType
type QGeoCoordinate__CoordinateType int64

const (
	QGeoCoordinate__InvalidCoordinate QGeoCoordinate__CoordinateType = QGeoCoordinate__CoordinateType(0)
	QGeoCoordinate__Coordinate2D      QGeoCoordinate__CoordinateType = QGeoCoordinate__CoordinateType(1)
	QGeoCoordinate__Coordinate3D      QGeoCoordinate__CoordinateType = QGeoCoordinate__CoordinateType(2)
)

func NewQGeoCoordinate() *QGeoCoordinate {
	var tmpValue = NewQGeoCoordinateFromPointer(C.QGeoCoordinate_NewQGeoCoordinate())
	runtime.SetFinalizer(tmpValue, (*QGeoCoordinate).DestroyQGeoCoordinate)
	return tmpValue
}

func NewQGeoCoordinate4(other QGeoCoordinate_ITF) *QGeoCoordinate {
	var tmpValue = NewQGeoCoordinateFromPointer(C.QGeoCoordinate_NewQGeoCoordinate4(PointerFromQGeoCoordinate(other)))
	runtime.SetFinalizer(tmpValue, (*QGeoCoordinate).DestroyQGeoCoordinate)
	return tmpValue
}

func NewQGeoCoordinate2(latitude float64, longitude float64) *QGeoCoordinate {
	var tmpValue = NewQGeoCoordinateFromPointer(C.QGeoCoordinate_NewQGeoCoordinate2(C.double(latitude), C.double(longitude)))
	runtime.SetFinalizer(tmpValue, (*QGeoCoordinate).DestroyQGeoCoordinate)
	return tmpValue
}

func NewQGeoCoordinate3(latitude float64, longitude float64, altitude float64) *QGeoCoordinate {
	var tmpValue = NewQGeoCoordinateFromPointer(C.QGeoCoordinate_NewQGeoCoordinate3(C.double(latitude), C.double(longitude), C.double(altitude)))
	runtime.SetFinalizer(tmpValue, (*QGeoCoordinate).DestroyQGeoCoordinate)
	return tmpValue
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

func (ptr *QGeoCoordinate) DestroyQGeoCoordinate() {
	if ptr.Pointer() != nil {
		C.QGeoCoordinate_DestroyQGeoCoordinate(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QGeoCoordinate) Type() QGeoCoordinate__CoordinateType {
	if ptr.Pointer() != nil {
		return QGeoCoordinate__CoordinateType(C.QGeoCoordinate_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoCoordinate) AtDistanceAndAzimuth(distance float64, azimuth float64, distanceUp float64) *QGeoCoordinate {
	if ptr.Pointer() != nil {
		var tmpValue = NewQGeoCoordinateFromPointer(C.QGeoCoordinate_AtDistanceAndAzimuth(ptr.Pointer(), C.double(distance), C.double(azimuth), C.double(distanceUp)))
		runtime.SetFinalizer(tmpValue, (*QGeoCoordinate).DestroyQGeoCoordinate)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoCoordinate) ToString(format QGeoCoordinate__CoordinateFormat) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoCoordinate_ToString(ptr.Pointer(), C.longlong(format)))
	}
	return ""
}

func (ptr *QGeoCoordinate) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QGeoCoordinate_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGeoCoordinate) Altitude() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGeoCoordinate_Altitude(ptr.Pointer()))
	}
	return 0
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

func NewQGeoLocationFromPointer(ptr unsafe.Pointer) *QGeoLocation {
	var n = new(QGeoLocation)
	n.SetPointer(ptr)
	return n
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

func NewQGeoPathFromPointer(ptr unsafe.Pointer) *QGeoPath {
	var n = new(QGeoPath)
	n.SetPointer(ptr)
	return n
}
func NewQGeoPath() *QGeoPath {
	var tmpValue = NewQGeoPathFromPointer(C.QGeoPath_NewQGeoPath())
	runtime.SetFinalizer(tmpValue, (*QGeoPath).DestroyQGeoPath)
	return tmpValue
}

func NewQGeoPath3(other QGeoPath_ITF) *QGeoPath {
	var tmpValue = NewQGeoPathFromPointer(C.QGeoPath_NewQGeoPath3(PointerFromQGeoPath(other)))
	runtime.SetFinalizer(tmpValue, (*QGeoPath).DestroyQGeoPath)
	return tmpValue
}

func NewQGeoPath4(other QGeoShape_ITF) *QGeoPath {
	var tmpValue = NewQGeoPathFromPointer(C.QGeoPath_NewQGeoPath4(PointerFromQGeoShape(other)))
	runtime.SetFinalizer(tmpValue, (*QGeoPath).DestroyQGeoPath)
	return tmpValue
}

func NewQGeoPath2(path []*QGeoCoordinate, width float64) *QGeoPath {
	var tmpValue = NewQGeoPathFromPointer(C.QGeoPath_NewQGeoPath2(func() unsafe.Pointer {
		var tmpList = NewQGeoPathFromPointer(NewQGeoPathFromPointer(nil).__QGeoPath_path_newList2())
		for _, v := range path {
			tmpList.__QGeoPath_path_setList2(v)
		}
		return tmpList.Pointer()
	}(), C.double(width)))
	runtime.SetFinalizer(tmpValue, (*QGeoPath).DestroyQGeoPath)
	return tmpValue
}

func (ptr *QGeoPath) AddCoordinate(coordinate QGeoCoordinate_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPath_AddCoordinate(ptr.Pointer(), PointerFromQGeoCoordinate(coordinate))
	}
}

func (ptr *QGeoPath) InsertCoordinate(index int, coordinate QGeoCoordinate_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPath_InsertCoordinate(ptr.Pointer(), C.int(int32(index)), PointerFromQGeoCoordinate(coordinate))
	}
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
			var tmpList = NewQGeoPathFromPointer(NewQGeoPathFromPointer(nil).__setPath_path_newList())
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

func (ptr *QGeoPath) Translate(degreesLatitude float64, degreesLongitude float64) {
	if ptr.Pointer() != nil {
		C.QGeoPath_Translate(ptr.Pointer(), C.double(degreesLatitude), C.double(degreesLongitude))
	}
}

func (ptr *QGeoPath) DestroyQGeoPath() {
	if ptr.Pointer() != nil {
		C.QGeoPath_DestroyQGeoPath(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QGeoPath) CoordinateAt(index int) *QGeoCoordinate {
	if ptr.Pointer() != nil {
		var tmpValue = NewQGeoCoordinateFromPointer(C.QGeoPath_CoordinateAt(ptr.Pointer(), C.int(int32(index))))
		runtime.SetFinalizer(tmpValue, (*QGeoCoordinate).DestroyQGeoCoordinate)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoPath) Translated(degreesLatitude float64, degreesLongitude float64) *QGeoPath {
	if ptr.Pointer() != nil {
		var tmpValue = NewQGeoPathFromPointer(C.QGeoPath_Translated(ptr.Pointer(), C.double(degreesLatitude), C.double(degreesLongitude)))
		runtime.SetFinalizer(tmpValue, (*QGeoPath).DestroyQGeoPath)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoPath) ContainsCoordinate(coordinate QGeoCoordinate_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGeoPath_ContainsCoordinate(ptr.Pointer(), PointerFromQGeoCoordinate(coordinate)) != 0
	}
	return false
}

func (ptr *QGeoPath) Path() []*QGeoCoordinate {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtPositioning_PackedList) []*QGeoCoordinate {
			var out = make([]*QGeoCoordinate, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQGeoPathFromPointer(l.data).__path_atList(i)
			}
			return out
		}(C.QGeoPath_Path(ptr.Pointer()))
	}
	return make([]*QGeoCoordinate, 0)
}

func (ptr *QGeoPath) Length(indexFrom int, indexTo int) float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGeoPath_Length(ptr.Pointer(), C.int(int32(indexFrom)), C.int(int32(indexTo))))
	}
	return 0
}

func (ptr *QGeoPath) Width() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGeoPath_Width(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoPath) __QGeoPath_path_atList2(i int) *QGeoCoordinate {
	if ptr.Pointer() != nil {
		var tmpValue = NewQGeoCoordinateFromPointer(C.QGeoPath___QGeoPath_path_atList2(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QGeoCoordinate).DestroyQGeoCoordinate)
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
	return unsafe.Pointer(C.QGeoPath___QGeoPath_path_newList2(ptr.Pointer()))
}

func (ptr *QGeoPath) __setPath_path_atList(i int) *QGeoCoordinate {
	if ptr.Pointer() != nil {
		var tmpValue = NewQGeoCoordinateFromPointer(C.QGeoPath___setPath_path_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QGeoCoordinate).DestroyQGeoCoordinate)
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
	return unsafe.Pointer(C.QGeoPath___setPath_path_newList(ptr.Pointer()))
}

func (ptr *QGeoPath) __path_atList(i int) *QGeoCoordinate {
	if ptr.Pointer() != nil {
		var tmpValue = NewQGeoCoordinateFromPointer(C.QGeoPath___path_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QGeoCoordinate).DestroyQGeoCoordinate)
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
	return unsafe.Pointer(C.QGeoPath___path_newList(ptr.Pointer()))
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

func NewQGeoPositionInfoFromPointer(ptr unsafe.Pointer) *QGeoPositionInfo {
	var n = new(QGeoPositionInfo)
	n.SetPointer(ptr)
	return n
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
	var tmpValue = NewQGeoPositionInfoFromPointer(C.QGeoPositionInfo_NewQGeoPositionInfo())
	runtime.SetFinalizer(tmpValue, (*QGeoPositionInfo).DestroyQGeoPositionInfo)
	return tmpValue
}

func NewQGeoPositionInfo2(coordinate QGeoCoordinate_ITF, timestamp core.QDateTime_ITF) *QGeoPositionInfo {
	var tmpValue = NewQGeoPositionInfoFromPointer(C.QGeoPositionInfo_NewQGeoPositionInfo2(PointerFromQGeoCoordinate(coordinate), core.PointerFromQDateTime(timestamp)))
	runtime.SetFinalizer(tmpValue, (*QGeoPositionInfo).DestroyQGeoPositionInfo)
	return tmpValue
}

func NewQGeoPositionInfo3(other QGeoPositionInfo_ITF) *QGeoPositionInfo {
	var tmpValue = NewQGeoPositionInfoFromPointer(C.QGeoPositionInfo_NewQGeoPositionInfo3(PointerFromQGeoPositionInfo(other)))
	runtime.SetFinalizer(tmpValue, (*QGeoPositionInfo).DestroyQGeoPositionInfo)
	return tmpValue
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

func (ptr *QGeoPositionInfo) DestroyQGeoPositionInfo() {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfo_DestroyQGeoPositionInfo(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QGeoPositionInfo) Timestamp() *core.QDateTime {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQDateTimeFromPointer(C.QGeoPositionInfo_Timestamp(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QDateTime).DestroyQDateTime)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoPositionInfo) Coordinate() *QGeoCoordinate {
	if ptr.Pointer() != nil {
		var tmpValue = NewQGeoCoordinateFromPointer(C.QGeoPositionInfo_Coordinate(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QGeoCoordinate).DestroyQGeoCoordinate)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoPositionInfo) HasAttribute(attribute QGeoPositionInfo__Attribute) bool {
	if ptr.Pointer() != nil {
		return C.QGeoPositionInfo_HasAttribute(ptr.Pointer(), C.longlong(attribute)) != 0
	}
	return false
}

func (ptr *QGeoPositionInfo) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QGeoPositionInfo_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGeoPositionInfo) Attribute(attribute QGeoPositionInfo__Attribute) float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGeoPositionInfo_Attribute(ptr.Pointer(), C.longlong(attribute)))
	}
	return 0
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

func NewQGeoPositionInfoSourceFromPointer(ptr unsafe.Pointer) *QGeoPositionInfoSource {
	var n = new(QGeoPositionInfoSource)
	n.SetPointer(ptr)
	return n
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

//export callbackQGeoPositionInfoSource_PositionUpdated
func callbackQGeoPositionInfoSource_PositionUpdated(ptr unsafe.Pointer, update unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "positionUpdated"); signal != nil {
		signal.(func(*QGeoPositionInfo))(NewQGeoPositionInfoFromPointer(update))
	}

}

func (ptr *QGeoPositionInfoSource) ConnectPositionUpdated(f func(update *QGeoPositionInfo)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "positionUpdated") {
			C.QGeoPositionInfoSource_ConnectPositionUpdated(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "positionUpdated"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "positionUpdated", func(update *QGeoPositionInfo) {
				signal.(func(*QGeoPositionInfo))(update)
				f(update)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "positionUpdated", f)
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

//export callbackQGeoPositionInfoSource_SetUpdateInterval
func callbackQGeoPositionInfoSource_SetUpdateInterval(ptr unsafe.Pointer, msec C.int) {
	if signal := qt.GetSignal(ptr, "setUpdateInterval"); signal != nil {
		signal.(func(int))(int(int32(msec)))
	} else {
		NewQGeoPositionInfoSourceFromPointer(ptr).SetUpdateIntervalDefault(int(int32(msec)))
	}
}

func (ptr *QGeoPositionInfoSource) ConnectSetUpdateInterval(f func(msec int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setUpdateInterval"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setUpdateInterval", func(msec int) {
				signal.(func(int))(msec)
				f(msec)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setUpdateInterval", f)
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

func QGeoPositionInfoSource_CreateDefaultSource(parent core.QObject_ITF) *QGeoPositionInfoSource {
	var tmpValue = NewQGeoPositionInfoSourceFromPointer(C.QGeoPositionInfoSource_QGeoPositionInfoSource_CreateDefaultSource(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QGeoPositionInfoSource) CreateDefaultSource(parent core.QObject_ITF) *QGeoPositionInfoSource {
	var tmpValue = NewQGeoPositionInfoSourceFromPointer(C.QGeoPositionInfoSource_QGeoPositionInfoSource_CreateDefaultSource(core.PointerFromQObject(parent)))
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
	var tmpValue = NewQGeoPositionInfoSourceFromPointer(C.QGeoPositionInfoSource_QGeoPositionInfoSource_CreateSource(C.struct_QtPositioning_PackedString{data: sourceNameC, len: C.longlong(len(sourceName))}, core.PointerFromQObject(parent)))
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
	var tmpValue = NewQGeoPositionInfoSourceFromPointer(C.QGeoPositionInfoSource_QGeoPositionInfoSource_CreateSource(C.struct_QtPositioning_PackedString{data: sourceNameC, len: C.longlong(len(sourceName))}, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQGeoPositionInfoSource(parent core.QObject_ITF) *QGeoPositionInfoSource {
	var tmpValue = NewQGeoPositionInfoSourceFromPointer(C.QGeoPositionInfoSource_NewQGeoPositionInfoSource(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QGeoPositionInfoSource_AvailableSources() []string {
	return strings.Split(cGoUnpackString(C.QGeoPositionInfoSource_QGeoPositionInfoSource_AvailableSources()), "|")
}

func (ptr *QGeoPositionInfoSource) AvailableSources() []string {
	return strings.Split(cGoUnpackString(C.QGeoPositionInfoSource_QGeoPositionInfoSource_AvailableSources()), "|")
}

//export callbackQGeoPositionInfoSource_Error2
func callbackQGeoPositionInfoSource_Error2(ptr unsafe.Pointer, positioningError C.longlong) {
	if signal := qt.GetSignal(ptr, "error2"); signal != nil {
		signal.(func(QGeoPositionInfoSource__Error))(QGeoPositionInfoSource__Error(positioningError))
	}

}

func (ptr *QGeoPositionInfoSource) ConnectError2(f func(positioningError QGeoPositionInfoSource__Error)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "error2") {
			C.QGeoPositionInfoSource_ConnectError2(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "error2"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "error2", func(positioningError QGeoPositionInfoSource__Error) {
				signal.(func(QGeoPositionInfoSource__Error))(positioningError)
				f(positioningError)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "error2", f)
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

//export callbackQGeoPositionInfoSource_RequestUpdate
func callbackQGeoPositionInfoSource_RequestUpdate(ptr unsafe.Pointer, timeout C.int) {
	if signal := qt.GetSignal(ptr, "requestUpdate"); signal != nil {
		signal.(func(int))(int(int32(timeout)))
	}

}

func (ptr *QGeoPositionInfoSource) ConnectRequestUpdate(f func(timeout int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "requestUpdate"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "requestUpdate", func(timeout int) {
				signal.(func(int))(timeout)
				f(timeout)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "requestUpdate", f)
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
		signal.(func(QGeoPositionInfoSource__PositioningMethod))(QGeoPositionInfoSource__PositioningMethod(methods))
	} else {
		NewQGeoPositionInfoSourceFromPointer(ptr).SetPreferredPositioningMethodsDefault(QGeoPositionInfoSource__PositioningMethod(methods))
	}
}

func (ptr *QGeoPositionInfoSource) ConnectSetPreferredPositioningMethods(f func(methods QGeoPositionInfoSource__PositioningMethod)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setPreferredPositioningMethods"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setPreferredPositioningMethods", func(methods QGeoPositionInfoSource__PositioningMethod) {
				signal.(func(QGeoPositionInfoSource__PositioningMethod))(methods)
				f(methods)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setPreferredPositioningMethods", f)
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

//export callbackQGeoPositionInfoSource_StartUpdates
func callbackQGeoPositionInfoSource_StartUpdates(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "startUpdates"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QGeoPositionInfoSource) ConnectStartUpdates(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "startUpdates"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "startUpdates", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "startUpdates", f)
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
		signal.(func())()
	}

}

func (ptr *QGeoPositionInfoSource) ConnectStopUpdates(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "stopUpdates"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "stopUpdates", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stopUpdates", f)
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

//export callbackQGeoPositionInfoSource_UpdateTimeout
func callbackQGeoPositionInfoSource_UpdateTimeout(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateTimeout"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QGeoPositionInfoSource) ConnectUpdateTimeout(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "updateTimeout") {
			C.QGeoPositionInfoSource_ConnectUpdateTimeout(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "updateTimeout"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "updateTimeout", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "updateTimeout", f)
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
		signal.(func())()
	} else {
		NewQGeoPositionInfoSourceFromPointer(ptr).DestroyQGeoPositionInfoSourceDefault()
	}
}

func (ptr *QGeoPositionInfoSource) ConnectDestroyQGeoPositionInfoSource(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QGeoPositionInfoSource"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QGeoPositionInfoSource", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QGeoPositionInfoSource", f)
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
		C.QGeoPositionInfoSource_DestroyQGeoPositionInfoSource(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QGeoPositionInfoSource) DestroyQGeoPositionInfoSourceDefault() {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_DestroyQGeoPositionInfoSourceDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQGeoPositionInfoSource_Error
func callbackQGeoPositionInfoSource_Error(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "error"); signal != nil {
		return C.longlong(signal.(func() QGeoPositionInfoSource__Error)())
	}

	return C.longlong(0)
}

func (ptr *QGeoPositionInfoSource) ConnectError(f func() QGeoPositionInfoSource__Error) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "error"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "error", func() QGeoPositionInfoSource__Error {
				signal.(func() QGeoPositionInfoSource__Error)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "error", f)
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

func (ptr *QGeoPositionInfoSource) PreferredPositioningMethods() QGeoPositionInfoSource__PositioningMethod {
	if ptr.Pointer() != nil {
		return QGeoPositionInfoSource__PositioningMethod(C.QGeoPositionInfoSource_PreferredPositioningMethods(ptr.Pointer()))
	}
	return 0
}

//export callbackQGeoPositionInfoSource_SupportedPositioningMethods
func callbackQGeoPositionInfoSource_SupportedPositioningMethods(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedPositioningMethods"); signal != nil {
		return C.longlong(signal.(func() QGeoPositionInfoSource__PositioningMethod)())
	}

	return C.longlong(0)
}

func (ptr *QGeoPositionInfoSource) ConnectSupportedPositioningMethods(f func() QGeoPositionInfoSource__PositioningMethod) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "supportedPositioningMethods"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "supportedPositioningMethods", func() QGeoPositionInfoSource__PositioningMethod {
				signal.(func() QGeoPositionInfoSource__PositioningMethod)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "supportedPositioningMethods", f)
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

//export callbackQGeoPositionInfoSource_LastKnownPosition
func callbackQGeoPositionInfoSource_LastKnownPosition(ptr unsafe.Pointer, fromSatellitePositioningMethodsOnly C.char) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "lastKnownPosition"); signal != nil {
		return PointerFromQGeoPositionInfo(signal.(func(bool) *QGeoPositionInfo)(int8(fromSatellitePositioningMethodsOnly) != 0))
	}

	return PointerFromQGeoPositionInfo(NewQGeoPositionInfo())
}

func (ptr *QGeoPositionInfoSource) ConnectLastKnownPosition(f func(fromSatellitePositioningMethodsOnly bool) *QGeoPositionInfo) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "lastKnownPosition"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "lastKnownPosition", func(fromSatellitePositioningMethodsOnly bool) *QGeoPositionInfo {
				signal.(func(bool) *QGeoPositionInfo)(fromSatellitePositioningMethodsOnly)
				return f(fromSatellitePositioningMethodsOnly)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "lastKnownPosition", f)
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
		var tmpValue = NewQGeoPositionInfoFromPointer(C.QGeoPositionInfoSource_LastKnownPosition(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(fromSatellitePositioningMethodsOnly)))))
		runtime.SetFinalizer(tmpValue, (*QGeoPositionInfo).DestroyQGeoPositionInfo)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoPositionInfoSource) SourceName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoPositionInfoSource_SourceName(ptr.Pointer()))
	}
	return ""
}

//export callbackQGeoPositionInfoSource_MinimumUpdateInterval
func callbackQGeoPositionInfoSource_MinimumUpdateInterval(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "minimumUpdateInterval"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(0))
}

func (ptr *QGeoPositionInfoSource) ConnectMinimumUpdateInterval(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "minimumUpdateInterval"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "minimumUpdateInterval", func() int {
				signal.(func() int)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "minimumUpdateInterval", f)
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

func (ptr *QGeoPositionInfoSource) UpdateInterval() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QGeoPositionInfoSource_UpdateInterval(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoPositionInfoSource) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QGeoPositionInfoSource___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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
	return unsafe.Pointer(C.QGeoPositionInfoSource___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QGeoPositionInfoSource) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QGeoPositionInfoSource___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGeoPositionInfoSource) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGeoPositionInfoSource) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QGeoPositionInfoSource___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QGeoPositionInfoSource) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QGeoPositionInfoSource___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
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
	return unsafe.Pointer(C.QGeoPositionInfoSource___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QGeoPositionInfoSource) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QGeoPositionInfoSource___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
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
	return unsafe.Pointer(C.QGeoPositionInfoSource___findChildren_newList(ptr.Pointer()))
}

func (ptr *QGeoPositionInfoSource) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QGeoPositionInfoSource___children_atList(ptr.Pointer(), C.int(int32(i))))
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
	return unsafe.Pointer(C.QGeoPositionInfoSource___children_newList(ptr.Pointer()))
}

//export callbackQGeoPositionInfoSource_Event
func callbackQGeoPositionInfoSource_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGeoPositionInfoSourceFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QGeoPositionInfoSource) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGeoPositionInfoSource_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQGeoPositionInfoSource_EventFilter
func callbackQGeoPositionInfoSource_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGeoPositionInfoSourceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QGeoPositionInfoSource) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGeoPositionInfoSource_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQGeoPositionInfoSource_ChildEvent
func callbackQGeoPositionInfoSource_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
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
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
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
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
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
		signal.(func())()
	} else {
		NewQGeoPositionInfoSourceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QGeoPositionInfoSource) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQGeoPositionInfoSource_Destroyed
func callbackQGeoPositionInfoSource_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQGeoPositionInfoSource_DisconnectNotify
func callbackQGeoPositionInfoSource_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGeoPositionInfoSourceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGeoPositionInfoSource_ObjectNameChanged
func callbackQGeoPositionInfoSource_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtPositioning_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQGeoPositionInfoSource_TimerEvent
func callbackQGeoPositionInfoSource_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGeoPositionInfoSourceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGeoPositionInfoSource) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQGeoPositionInfoSource_MetaObject
func callbackQGeoPositionInfoSource_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQGeoPositionInfoSourceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QGeoPositionInfoSource) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QGeoPositionInfoSource_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
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

func NewQGeoPositionInfoSourceFactoryFromPointer(ptr unsafe.Pointer) *QGeoPositionInfoSourceFactory {
	var n = new(QGeoPositionInfoSourceFactory)
	n.SetPointer(ptr)
	return n
}

//export callbackQGeoPositionInfoSourceFactory_AreaMonitor
func callbackQGeoPositionInfoSourceFactory_AreaMonitor(ptr unsafe.Pointer, parent unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "areaMonitor"); signal != nil {
		return PointerFromQGeoAreaMonitorSource(signal.(func(*core.QObject) *QGeoAreaMonitorSource)(core.NewQObjectFromPointer(parent)))
	}

	return PointerFromQGeoAreaMonitorSource(NewQGeoAreaMonitorSource(nil))
}

func (ptr *QGeoPositionInfoSourceFactory) ConnectAreaMonitor(f func(parent *core.QObject) *QGeoAreaMonitorSource) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "areaMonitor"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "areaMonitor", func(parent *core.QObject) *QGeoAreaMonitorSource {
				signal.(func(*core.QObject) *QGeoAreaMonitorSource)(parent)
				return f(parent)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "areaMonitor", f)
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
		var tmpValue = NewQGeoAreaMonitorSourceFromPointer(C.QGeoPositionInfoSourceFactory_AreaMonitor(ptr.Pointer(), core.PointerFromQObject(parent)))
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
		return PointerFromQGeoPositionInfoSource(signal.(func(*core.QObject) *QGeoPositionInfoSource)(core.NewQObjectFromPointer(parent)))
	}

	return PointerFromQGeoPositionInfoSource(NewQGeoPositionInfoSource(nil))
}

func (ptr *QGeoPositionInfoSourceFactory) ConnectPositionInfoSource(f func(parent *core.QObject) *QGeoPositionInfoSource) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "positionInfoSource"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "positionInfoSource", func(parent *core.QObject) *QGeoPositionInfoSource {
				signal.(func(*core.QObject) *QGeoPositionInfoSource)(parent)
				return f(parent)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "positionInfoSource", f)
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
		var tmpValue = NewQGeoPositionInfoSourceFromPointer(C.QGeoPositionInfoSourceFactory_PositionInfoSource(ptr.Pointer(), core.PointerFromQObject(parent)))
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
		return PointerFromQGeoSatelliteInfoSource(signal.(func(*core.QObject) *QGeoSatelliteInfoSource)(core.NewQObjectFromPointer(parent)))
	}

	return PointerFromQGeoSatelliteInfoSource(NewQGeoSatelliteInfoSource(nil))
}

func (ptr *QGeoPositionInfoSourceFactory) ConnectSatelliteInfoSource(f func(parent *core.QObject) *QGeoSatelliteInfoSource) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "satelliteInfoSource"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "satelliteInfoSource", func(parent *core.QObject) *QGeoSatelliteInfoSource {
				signal.(func(*core.QObject) *QGeoSatelliteInfoSource)(parent)
				return f(parent)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "satelliteInfoSource", f)
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
		var tmpValue = NewQGeoSatelliteInfoSourceFromPointer(C.QGeoPositionInfoSourceFactory_SatelliteInfoSource(ptr.Pointer(), core.PointerFromQObject(parent)))
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
		signal.(func())()
	} else {
		NewQGeoPositionInfoSourceFactoryFromPointer(ptr).DestroyQGeoPositionInfoSourceFactoryDefault()
	}
}

func (ptr *QGeoPositionInfoSourceFactory) ConnectDestroyQGeoPositionInfoSourceFactory(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QGeoPositionInfoSourceFactory"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QGeoPositionInfoSourceFactory", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QGeoPositionInfoSourceFactory", f)
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

func NewQGeoRectangleFromPointer(ptr unsafe.Pointer) *QGeoRectangle {
	var n = new(QGeoRectangle)
	n.SetPointer(ptr)
	return n
}
func NewQGeoRectangle() *QGeoRectangle {
	var tmpValue = NewQGeoRectangleFromPointer(C.QGeoRectangle_NewQGeoRectangle())
	runtime.SetFinalizer(tmpValue, (*QGeoRectangle).DestroyQGeoRectangle)
	return tmpValue
}

func NewQGeoRectangle2(center QGeoCoordinate_ITF, degreesWidth float64, degreesHeight float64) *QGeoRectangle {
	var tmpValue = NewQGeoRectangleFromPointer(C.QGeoRectangle_NewQGeoRectangle2(PointerFromQGeoCoordinate(center), C.double(degreesWidth), C.double(degreesHeight)))
	runtime.SetFinalizer(tmpValue, (*QGeoRectangle).DestroyQGeoRectangle)
	return tmpValue
}

func NewQGeoRectangle3(topLeft QGeoCoordinate_ITF, bottomRight QGeoCoordinate_ITF) *QGeoRectangle {
	var tmpValue = NewQGeoRectangleFromPointer(C.QGeoRectangle_NewQGeoRectangle3(PointerFromQGeoCoordinate(topLeft), PointerFromQGeoCoordinate(bottomRight)))
	runtime.SetFinalizer(tmpValue, (*QGeoRectangle).DestroyQGeoRectangle)
	return tmpValue
}

func NewQGeoRectangle5(other QGeoRectangle_ITF) *QGeoRectangle {
	var tmpValue = NewQGeoRectangleFromPointer(C.QGeoRectangle_NewQGeoRectangle5(PointerFromQGeoRectangle(other)))
	runtime.SetFinalizer(tmpValue, (*QGeoRectangle).DestroyQGeoRectangle)
	return tmpValue
}

func NewQGeoRectangle6(other QGeoShape_ITF) *QGeoRectangle {
	var tmpValue = NewQGeoRectangleFromPointer(C.QGeoRectangle_NewQGeoRectangle6(PointerFromQGeoShape(other)))
	runtime.SetFinalizer(tmpValue, (*QGeoRectangle).DestroyQGeoRectangle)
	return tmpValue
}

func NewQGeoRectangle4(coordinates []*QGeoCoordinate) *QGeoRectangle {
	var tmpValue = NewQGeoRectangleFromPointer(C.QGeoRectangle_NewQGeoRectangle4(func() unsafe.Pointer {
		var tmpList = NewQGeoRectangleFromPointer(NewQGeoRectangleFromPointer(nil).__QGeoRectangle_coordinates_newList4())
		for _, v := range coordinates {
			tmpList.__QGeoRectangle_coordinates_setList4(v)
		}
		return tmpList.Pointer()
	}()))
	runtime.SetFinalizer(tmpValue, (*QGeoRectangle).DestroyQGeoRectangle)
	return tmpValue
}

func (ptr *QGeoRectangle) ExtendRectangle(coordinate QGeoCoordinate_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRectangle_ExtendRectangle(ptr.Pointer(), PointerFromQGeoCoordinate(coordinate))
	}
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

func (ptr *QGeoRectangle) Translate(degreesLatitude float64, degreesLongitude float64) {
	if ptr.Pointer() != nil {
		C.QGeoRectangle_Translate(ptr.Pointer(), C.double(degreesLatitude), C.double(degreesLongitude))
	}
}

func (ptr *QGeoRectangle) DestroyQGeoRectangle() {
	if ptr.Pointer() != nil {
		C.QGeoRectangle_DestroyQGeoRectangle(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QGeoRectangle) BottomLeft() *QGeoCoordinate {
	if ptr.Pointer() != nil {
		var tmpValue = NewQGeoCoordinateFromPointer(C.QGeoRectangle_BottomLeft(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QGeoCoordinate).DestroyQGeoCoordinate)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRectangle) BottomRight() *QGeoCoordinate {
	if ptr.Pointer() != nil {
		var tmpValue = NewQGeoCoordinateFromPointer(C.QGeoRectangle_BottomRight(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QGeoCoordinate).DestroyQGeoCoordinate)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRectangle) TopLeft() *QGeoCoordinate {
	if ptr.Pointer() != nil {
		var tmpValue = NewQGeoCoordinateFromPointer(C.QGeoRectangle_TopLeft(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QGeoCoordinate).DestroyQGeoCoordinate)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRectangle) TopRight() *QGeoCoordinate {
	if ptr.Pointer() != nil {
		var tmpValue = NewQGeoCoordinateFromPointer(C.QGeoRectangle_TopRight(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QGeoCoordinate).DestroyQGeoCoordinate)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRectangle) Translated(degreesLatitude float64, degreesLongitude float64) *QGeoRectangle {
	if ptr.Pointer() != nil {
		var tmpValue = NewQGeoRectangleFromPointer(C.QGeoRectangle_Translated(ptr.Pointer(), C.double(degreesLatitude), C.double(degreesLongitude)))
		runtime.SetFinalizer(tmpValue, (*QGeoRectangle).DestroyQGeoRectangle)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRectangle) United(rectangle QGeoRectangle_ITF) *QGeoRectangle {
	if ptr.Pointer() != nil {
		var tmpValue = NewQGeoRectangleFromPointer(C.QGeoRectangle_United(ptr.Pointer(), PointerFromQGeoRectangle(rectangle)))
		runtime.SetFinalizer(tmpValue, (*QGeoRectangle).DestroyQGeoRectangle)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoRectangle) Contains(rectangle QGeoRectangle_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGeoRectangle_Contains(ptr.Pointer(), PointerFromQGeoRectangle(rectangle)) != 0
	}
	return false
}

func (ptr *QGeoRectangle) Intersects(rectangle QGeoRectangle_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGeoRectangle_Intersects(ptr.Pointer(), PointerFromQGeoRectangle(rectangle)) != 0
	}
	return false
}

func (ptr *QGeoRectangle) Height() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGeoRectangle_Height(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRectangle) Width() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGeoRectangle_Width(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRectangle) __QGeoRectangle_coordinates_atList4(i int) *QGeoCoordinate {
	if ptr.Pointer() != nil {
		var tmpValue = NewQGeoCoordinateFromPointer(C.QGeoRectangle___QGeoRectangle_coordinates_atList4(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QGeoCoordinate).DestroyQGeoCoordinate)
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
	return unsafe.Pointer(C.QGeoRectangle___QGeoRectangle_coordinates_newList4(ptr.Pointer()))
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

func NewQGeoSatelliteInfoFromPointer(ptr unsafe.Pointer) *QGeoSatelliteInfo {
	var n = new(QGeoSatelliteInfo)
	n.SetPointer(ptr)
	return n
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
	var tmpValue = NewQGeoSatelliteInfoFromPointer(C.QGeoSatelliteInfo_NewQGeoSatelliteInfo())
	runtime.SetFinalizer(tmpValue, (*QGeoSatelliteInfo).DestroyQGeoSatelliteInfo)
	return tmpValue
}

func NewQGeoSatelliteInfo2(other QGeoSatelliteInfo_ITF) *QGeoSatelliteInfo {
	var tmpValue = NewQGeoSatelliteInfoFromPointer(C.QGeoSatelliteInfo_NewQGeoSatelliteInfo2(PointerFromQGeoSatelliteInfo(other)))
	runtime.SetFinalizer(tmpValue, (*QGeoSatelliteInfo).DestroyQGeoSatelliteInfo)
	return tmpValue
}

func (ptr *QGeoSatelliteInfo) RemoveAttribute(attribute QGeoSatelliteInfo__Attribute) {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfo_RemoveAttribute(ptr.Pointer(), C.longlong(attribute))
	}
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

func (ptr *QGeoSatelliteInfo) DestroyQGeoSatelliteInfo() {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfo_DestroyQGeoSatelliteInfo(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QGeoSatelliteInfo) SatelliteSystem() QGeoSatelliteInfo__SatelliteSystem {
	if ptr.Pointer() != nil {
		return QGeoSatelliteInfo__SatelliteSystem(C.QGeoSatelliteInfo_SatelliteSystem(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoSatelliteInfo) HasAttribute(attribute QGeoSatelliteInfo__Attribute) bool {
	if ptr.Pointer() != nil {
		return C.QGeoSatelliteInfo_HasAttribute(ptr.Pointer(), C.longlong(attribute)) != 0
	}
	return false
}

func (ptr *QGeoSatelliteInfo) SatelliteIdentifier() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QGeoSatelliteInfo_SatelliteIdentifier(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoSatelliteInfo) SignalStrength() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QGeoSatelliteInfo_SignalStrength(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoSatelliteInfo) Attribute(attribute QGeoSatelliteInfo__Attribute) float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGeoSatelliteInfo_Attribute(ptr.Pointer(), C.longlong(attribute)))
	}
	return 0
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

func NewQGeoSatelliteInfoSourceFromPointer(ptr unsafe.Pointer) *QGeoSatelliteInfoSource {
	var n = new(QGeoSatelliteInfoSource)
	n.SetPointer(ptr)
	return n
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

func QGeoSatelliteInfoSource_CreateDefaultSource(parent core.QObject_ITF) *QGeoSatelliteInfoSource {
	var tmpValue = NewQGeoSatelliteInfoSourceFromPointer(C.QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_CreateDefaultSource(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QGeoSatelliteInfoSource) CreateDefaultSource(parent core.QObject_ITF) *QGeoSatelliteInfoSource {
	var tmpValue = NewQGeoSatelliteInfoSourceFromPointer(C.QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_CreateDefaultSource(core.PointerFromQObject(parent)))
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
	var tmpValue = NewQGeoSatelliteInfoSourceFromPointer(C.QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_CreateSource(C.struct_QtPositioning_PackedString{data: sourceNameC, len: C.longlong(len(sourceName))}, core.PointerFromQObject(parent)))
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
	var tmpValue = NewQGeoSatelliteInfoSourceFromPointer(C.QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_CreateSource(C.struct_QtPositioning_PackedString{data: sourceNameC, len: C.longlong(len(sourceName))}, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQGeoSatelliteInfoSource(parent core.QObject_ITF) *QGeoSatelliteInfoSource {
	var tmpValue = NewQGeoSatelliteInfoSourceFromPointer(C.QGeoSatelliteInfoSource_NewQGeoSatelliteInfoSource(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QGeoSatelliteInfoSource_AvailableSources() []string {
	return strings.Split(cGoUnpackString(C.QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_AvailableSources()), "|")
}

func (ptr *QGeoSatelliteInfoSource) AvailableSources() []string {
	return strings.Split(cGoUnpackString(C.QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_AvailableSources()), "|")
}

//export callbackQGeoSatelliteInfoSource_Error2
func callbackQGeoSatelliteInfoSource_Error2(ptr unsafe.Pointer, satelliteError C.longlong) {
	if signal := qt.GetSignal(ptr, "error2"); signal != nil {
		signal.(func(QGeoSatelliteInfoSource__Error))(QGeoSatelliteInfoSource__Error(satelliteError))
	}

}

func (ptr *QGeoSatelliteInfoSource) ConnectError2(f func(satelliteError QGeoSatelliteInfoSource__Error)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "error2") {
			C.QGeoSatelliteInfoSource_ConnectError2(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "error2"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "error2", func(satelliteError QGeoSatelliteInfoSource__Error) {
				signal.(func(QGeoSatelliteInfoSource__Error))(satelliteError)
				f(satelliteError)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "error2", f)
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

//export callbackQGeoSatelliteInfoSource_RequestTimeout
func callbackQGeoSatelliteInfoSource_RequestTimeout(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "requestTimeout"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QGeoSatelliteInfoSource) ConnectRequestTimeout(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "requestTimeout") {
			C.QGeoSatelliteInfoSource_ConnectRequestTimeout(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "requestTimeout"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "requestTimeout", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "requestTimeout", f)
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
		signal.(func(int))(int(int32(timeout)))
	}

}

func (ptr *QGeoSatelliteInfoSource) ConnectRequestUpdate(f func(timeout int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "requestUpdate"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "requestUpdate", func(timeout int) {
				signal.(func(int))(timeout)
				f(timeout)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "requestUpdate", f)
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
		signal.(func([]*QGeoSatelliteInfo))(func(l C.struct_QtPositioning_PackedList) []*QGeoSatelliteInfo {
			var out = make([]*QGeoSatelliteInfo, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQGeoSatelliteInfoSourceFromPointer(l.data).__satellitesInUseUpdated_satellites_atList(i)
			}
			return out
		}(satellites))
	}

}

func (ptr *QGeoSatelliteInfoSource) ConnectSatellitesInUseUpdated(f func(satellites []*QGeoSatelliteInfo)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "satellitesInUseUpdated") {
			C.QGeoSatelliteInfoSource_ConnectSatellitesInUseUpdated(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "satellitesInUseUpdated"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "satellitesInUseUpdated", func(satellites []*QGeoSatelliteInfo) {
				signal.(func([]*QGeoSatelliteInfo))(satellites)
				f(satellites)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "satellitesInUseUpdated", f)
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
			var tmpList = NewQGeoSatelliteInfoSourceFromPointer(NewQGeoSatelliteInfoSourceFromPointer(nil).__satellitesInUseUpdated_satellites_newList())
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
		signal.(func([]*QGeoSatelliteInfo))(func(l C.struct_QtPositioning_PackedList) []*QGeoSatelliteInfo {
			var out = make([]*QGeoSatelliteInfo, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQGeoSatelliteInfoSourceFromPointer(l.data).__satellitesInViewUpdated_satellites_atList(i)
			}
			return out
		}(satellites))
	}

}

func (ptr *QGeoSatelliteInfoSource) ConnectSatellitesInViewUpdated(f func(satellites []*QGeoSatelliteInfo)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "satellitesInViewUpdated") {
			C.QGeoSatelliteInfoSource_ConnectSatellitesInViewUpdated(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "satellitesInViewUpdated"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "satellitesInViewUpdated", func(satellites []*QGeoSatelliteInfo) {
				signal.(func([]*QGeoSatelliteInfo))(satellites)
				f(satellites)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "satellitesInViewUpdated", f)
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
			var tmpList = NewQGeoSatelliteInfoSourceFromPointer(NewQGeoSatelliteInfoSourceFromPointer(nil).__satellitesInViewUpdated_satellites_newList())
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
		signal.(func(int))(int(int32(msec)))
	} else {
		NewQGeoSatelliteInfoSourceFromPointer(ptr).SetUpdateIntervalDefault(int(int32(msec)))
	}
}

func (ptr *QGeoSatelliteInfoSource) ConnectSetUpdateInterval(f func(msec int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setUpdateInterval"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setUpdateInterval", func(msec int) {
				signal.(func(int))(msec)
				f(msec)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setUpdateInterval", f)
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

//export callbackQGeoSatelliteInfoSource_StartUpdates
func callbackQGeoSatelliteInfoSource_StartUpdates(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "startUpdates"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QGeoSatelliteInfoSource) ConnectStartUpdates(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "startUpdates"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "startUpdates", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "startUpdates", f)
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
		signal.(func())()
	}

}

func (ptr *QGeoSatelliteInfoSource) ConnectStopUpdates(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "stopUpdates"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "stopUpdates", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stopUpdates", f)
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

//export callbackQGeoSatelliteInfoSource_DestroyQGeoSatelliteInfoSource
func callbackQGeoSatelliteInfoSource_DestroyQGeoSatelliteInfoSource(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QGeoSatelliteInfoSource"); signal != nil {
		signal.(func())()
	} else {
		NewQGeoSatelliteInfoSourceFromPointer(ptr).DestroyQGeoSatelliteInfoSourceDefault()
	}
}

func (ptr *QGeoSatelliteInfoSource) ConnectDestroyQGeoSatelliteInfoSource(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QGeoSatelliteInfoSource"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QGeoSatelliteInfoSource", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QGeoSatelliteInfoSource", f)
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
		C.QGeoSatelliteInfoSource_DestroyQGeoSatelliteInfoSource(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QGeoSatelliteInfoSource) DestroyQGeoSatelliteInfoSourceDefault() {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_DestroyQGeoSatelliteInfoSourceDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQGeoSatelliteInfoSource_Error
func callbackQGeoSatelliteInfoSource_Error(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "error"); signal != nil {
		return C.longlong(signal.(func() QGeoSatelliteInfoSource__Error)())
	}

	return C.longlong(0)
}

func (ptr *QGeoSatelliteInfoSource) ConnectError(f func() QGeoSatelliteInfoSource__Error) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "error"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "error", func() QGeoSatelliteInfoSource__Error {
				signal.(func() QGeoSatelliteInfoSource__Error)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "error", f)
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

func (ptr *QGeoSatelliteInfoSource) SourceName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGeoSatelliteInfoSource_SourceName(ptr.Pointer()))
	}
	return ""
}

//export callbackQGeoSatelliteInfoSource_MinimumUpdateInterval
func callbackQGeoSatelliteInfoSource_MinimumUpdateInterval(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "minimumUpdateInterval"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(0))
}

func (ptr *QGeoSatelliteInfoSource) ConnectMinimumUpdateInterval(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "minimumUpdateInterval"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "minimumUpdateInterval", func() int {
				signal.(func() int)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "minimumUpdateInterval", f)
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

func (ptr *QGeoSatelliteInfoSource) UpdateInterval() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QGeoSatelliteInfoSource_UpdateInterval(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoSatelliteInfoSource) __satellitesInUseUpdated_satellites_atList(i int) *QGeoSatelliteInfo {
	if ptr.Pointer() != nil {
		var tmpValue = NewQGeoSatelliteInfoFromPointer(C.QGeoSatelliteInfoSource___satellitesInUseUpdated_satellites_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QGeoSatelliteInfo).DestroyQGeoSatelliteInfo)
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
	return unsafe.Pointer(C.QGeoSatelliteInfoSource___satellitesInUseUpdated_satellites_newList(ptr.Pointer()))
}

func (ptr *QGeoSatelliteInfoSource) __satellitesInViewUpdated_satellites_atList(i int) *QGeoSatelliteInfo {
	if ptr.Pointer() != nil {
		var tmpValue = NewQGeoSatelliteInfoFromPointer(C.QGeoSatelliteInfoSource___satellitesInViewUpdated_satellites_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QGeoSatelliteInfo).DestroyQGeoSatelliteInfo)
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
	return unsafe.Pointer(C.QGeoSatelliteInfoSource___satellitesInViewUpdated_satellites_newList(ptr.Pointer()))
}

func (ptr *QGeoSatelliteInfoSource) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QGeoSatelliteInfoSource___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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
	return unsafe.Pointer(C.QGeoSatelliteInfoSource___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QGeoSatelliteInfoSource) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QGeoSatelliteInfoSource___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGeoSatelliteInfoSource) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGeoSatelliteInfoSource) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QGeoSatelliteInfoSource___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QGeoSatelliteInfoSource) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QGeoSatelliteInfoSource___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
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
	return unsafe.Pointer(C.QGeoSatelliteInfoSource___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QGeoSatelliteInfoSource) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QGeoSatelliteInfoSource___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
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
	return unsafe.Pointer(C.QGeoSatelliteInfoSource___findChildren_newList(ptr.Pointer()))
}

func (ptr *QGeoSatelliteInfoSource) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QGeoSatelliteInfoSource___children_atList(ptr.Pointer(), C.int(int32(i))))
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
	return unsafe.Pointer(C.QGeoSatelliteInfoSource___children_newList(ptr.Pointer()))
}

//export callbackQGeoSatelliteInfoSource_Event
func callbackQGeoSatelliteInfoSource_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGeoSatelliteInfoSourceFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QGeoSatelliteInfoSource) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGeoSatelliteInfoSource_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQGeoSatelliteInfoSource_EventFilter
func callbackQGeoSatelliteInfoSource_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGeoSatelliteInfoSourceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QGeoSatelliteInfoSource) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGeoSatelliteInfoSource_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQGeoSatelliteInfoSource_ChildEvent
func callbackQGeoSatelliteInfoSource_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
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
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
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
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
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
		signal.(func())()
	} else {
		NewQGeoSatelliteInfoSourceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QGeoSatelliteInfoSource) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQGeoSatelliteInfoSource_Destroyed
func callbackQGeoSatelliteInfoSource_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQGeoSatelliteInfoSource_DisconnectNotify
func callbackQGeoSatelliteInfoSource_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGeoSatelliteInfoSourceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGeoSatelliteInfoSource) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGeoSatelliteInfoSource_ObjectNameChanged
func callbackQGeoSatelliteInfoSource_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtPositioning_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQGeoSatelliteInfoSource_TimerEvent
func callbackQGeoSatelliteInfoSource_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGeoSatelliteInfoSourceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGeoSatelliteInfoSource) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQGeoSatelliteInfoSource_MetaObject
func callbackQGeoSatelliteInfoSource_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQGeoSatelliteInfoSourceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QGeoSatelliteInfoSource) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QGeoSatelliteInfoSource_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
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

func NewQGeoShapeFromPointer(ptr unsafe.Pointer) *QGeoShape {
	var n = new(QGeoShape)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QGeoShape__ShapeType
//QGeoShape::ShapeType
type QGeoShape__ShapeType int64

const (
	QGeoShape__UnknownType   QGeoShape__ShapeType = QGeoShape__ShapeType(0)
	QGeoShape__RectangleType QGeoShape__ShapeType = QGeoShape__ShapeType(1)
	QGeoShape__CircleType    QGeoShape__ShapeType = QGeoShape__ShapeType(2)
	QGeoShape__PathType      QGeoShape__ShapeType = QGeoShape__ShapeType(3)
)

func NewQGeoShape() *QGeoShape {
	var tmpValue = NewQGeoShapeFromPointer(C.QGeoShape_NewQGeoShape())
	runtime.SetFinalizer(tmpValue, (*QGeoShape).DestroyQGeoShape)
	return tmpValue
}

func NewQGeoShape2(other QGeoShape_ITF) *QGeoShape {
	var tmpValue = NewQGeoShapeFromPointer(C.QGeoShape_NewQGeoShape2(PointerFromQGeoShape(other)))
	runtime.SetFinalizer(tmpValue, (*QGeoShape).DestroyQGeoShape)
	return tmpValue
}

func (ptr *QGeoShape) DestroyQGeoShape() {
	if ptr.Pointer() != nil {
		C.QGeoShape_DestroyQGeoShape(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QGeoShape) Center() *QGeoCoordinate {
	if ptr.Pointer() != nil {
		var tmpValue = NewQGeoCoordinateFromPointer(C.QGeoShape_Center(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QGeoCoordinate).DestroyQGeoCoordinate)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoShape) BoundingGeoRectangle() *QGeoRectangle {
	if ptr.Pointer() != nil {
		var tmpValue = NewQGeoRectangleFromPointer(C.QGeoShape_BoundingGeoRectangle(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QGeoRectangle).DestroyQGeoRectangle)
		return tmpValue
	}
	return nil
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

func (ptr *QGeoShape) Contains(coordinate QGeoCoordinate_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGeoShape_Contains(ptr.Pointer(), PointerFromQGeoCoordinate(coordinate)) != 0
	}
	return false
}

func (ptr *QGeoShape) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QGeoShape_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGeoShape) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QGeoShape_IsValid(ptr.Pointer()) != 0
	}
	return false
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

func NewQNmeaPositionInfoSourceFromPointer(ptr unsafe.Pointer) *QNmeaPositionInfoSource {
	var n = new(QNmeaPositionInfoSource)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QNmeaPositionInfoSource__UpdateMode
//QNmeaPositionInfoSource::UpdateMode
type QNmeaPositionInfoSource__UpdateMode int64

const (
	QNmeaPositionInfoSource__RealTimeMode   QNmeaPositionInfoSource__UpdateMode = QNmeaPositionInfoSource__UpdateMode(1)
	QNmeaPositionInfoSource__SimulationMode QNmeaPositionInfoSource__UpdateMode = QNmeaPositionInfoSource__UpdateMode(2)
)

func NewQNmeaPositionInfoSource(updateMode QNmeaPositionInfoSource__UpdateMode, parent core.QObject_ITF) *QNmeaPositionInfoSource {
	var tmpValue = NewQNmeaPositionInfoSourceFromPointer(C.QNmeaPositionInfoSource_NewQNmeaPositionInfoSource(C.longlong(updateMode), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQNmeaPositionInfoSource_ParsePosInfoFromNmeaData
func callbackQNmeaPositionInfoSource_ParsePosInfoFromNmeaData(ptr unsafe.Pointer, data C.struct_QtPositioning_PackedString, size C.int, posInfo unsafe.Pointer, hasFix C.char) C.char {
	if signal := qt.GetSignal(ptr, "parsePosInfoFromNmeaData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, int, *QGeoPositionInfo, bool) bool)(cGoUnpackString(data), int(int32(size)), NewQGeoPositionInfoFromPointer(posInfo), int8(hasFix) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNmeaPositionInfoSourceFromPointer(ptr).ParsePosInfoFromNmeaDataDefault(cGoUnpackString(data), int(int32(size)), NewQGeoPositionInfoFromPointer(posInfo), int8(hasFix) != 0))))
}

func (ptr *QNmeaPositionInfoSource) ConnectParsePosInfoFromNmeaData(f func(data string, size int, posInfo *QGeoPositionInfo, hasFix bool) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "parsePosInfoFromNmeaData"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "parsePosInfoFromNmeaData", func(data string, size int, posInfo *QGeoPositionInfo, hasFix bool) bool {
				signal.(func(string, int, *QGeoPositionInfo, bool) bool)(data, size, posInfo, hasFix)
				return f(data, size, posInfo, hasFix)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "parsePosInfoFromNmeaData", f)
		}
	}
}

func (ptr *QNmeaPositionInfoSource) DisconnectParsePosInfoFromNmeaData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "parsePosInfoFromNmeaData")
	}
}

func (ptr *QNmeaPositionInfoSource) ParsePosInfoFromNmeaData(data string, size int, posInfo QGeoPositionInfo_ITF, hasFix bool) bool {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if data != "" {
			dataC = C.CString(data)
			defer C.free(unsafe.Pointer(dataC))
		}
		return C.QNmeaPositionInfoSource_ParsePosInfoFromNmeaData(ptr.Pointer(), dataC, C.int(int32(size)), PointerFromQGeoPositionInfo(posInfo), C.char(int8(qt.GoBoolToInt(hasFix)))) != 0
	}
	return false
}

func (ptr *QNmeaPositionInfoSource) ParsePosInfoFromNmeaDataDefault(data string, size int, posInfo QGeoPositionInfo_ITF, hasFix bool) bool {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if data != "" {
			dataC = C.CString(data)
			defer C.free(unsafe.Pointer(dataC))
		}
		return C.QNmeaPositionInfoSource_ParsePosInfoFromNmeaDataDefault(ptr.Pointer(), dataC, C.int(int32(size)), PointerFromQGeoPositionInfo(posInfo), C.char(int8(qt.GoBoolToInt(hasFix)))) != 0
	}
	return false
}

//export callbackQNmeaPositionInfoSource_RequestUpdate
func callbackQNmeaPositionInfoSource_RequestUpdate(ptr unsafe.Pointer, msec C.int) {
	if signal := qt.GetSignal(ptr, "requestUpdate"); signal != nil {
		signal.(func(int))(int(int32(msec)))
	} else {
		NewQNmeaPositionInfoSourceFromPointer(ptr).RequestUpdateDefault(int(int32(msec)))
	}
}

func (ptr *QNmeaPositionInfoSource) ConnectRequestUpdate(f func(msec int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "requestUpdate"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "requestUpdate", func(msec int) {
				signal.(func(int))(msec)
				f(msec)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "requestUpdate", f)
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
		signal.(func())()
	} else {
		NewQNmeaPositionInfoSourceFromPointer(ptr).StartUpdatesDefault()
	}
}

func (ptr *QNmeaPositionInfoSource) ConnectStartUpdates(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "startUpdates"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "startUpdates", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "startUpdates", f)
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
		signal.(func())()
	} else {
		NewQNmeaPositionInfoSourceFromPointer(ptr).StopUpdatesDefault()
	}
}

func (ptr *QNmeaPositionInfoSource) ConnectStopUpdates(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "stopUpdates"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "stopUpdates", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stopUpdates", f)
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

func (ptr *QNmeaPositionInfoSource) DestroyQNmeaPositionInfoSource() {
	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_DestroyQNmeaPositionInfoSource(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQNmeaPositionInfoSource_Error
func callbackQNmeaPositionInfoSource_Error(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "error"); signal != nil {
		return C.longlong(signal.(func() QGeoPositionInfoSource__Error)())
	}

	return C.longlong(NewQNmeaPositionInfoSourceFromPointer(ptr).ErrorDefault())
}

func (ptr *QNmeaPositionInfoSource) ConnectError(f func() QGeoPositionInfoSource__Error) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "error"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "error", func() QGeoPositionInfoSource__Error {
				signal.(func() QGeoPositionInfoSource__Error)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "error", f)
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

//export callbackQNmeaPositionInfoSource_SupportedPositioningMethods
func callbackQNmeaPositionInfoSource_SupportedPositioningMethods(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedPositioningMethods"); signal != nil {
		return C.longlong(signal.(func() QGeoPositionInfoSource__PositioningMethod)())
	}

	return C.longlong(NewQNmeaPositionInfoSourceFromPointer(ptr).SupportedPositioningMethodsDefault())
}

func (ptr *QNmeaPositionInfoSource) ConnectSupportedPositioningMethods(f func() QGeoPositionInfoSource__PositioningMethod) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "supportedPositioningMethods"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "supportedPositioningMethods", func() QGeoPositionInfoSource__PositioningMethod {
				signal.(func() QGeoPositionInfoSource__PositioningMethod)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "supportedPositioningMethods", f)
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

//export callbackQNmeaPositionInfoSource_LastKnownPosition
func callbackQNmeaPositionInfoSource_LastKnownPosition(ptr unsafe.Pointer, fromSatellitePositioningMethodsOnly C.char) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "lastKnownPosition"); signal != nil {
		return PointerFromQGeoPositionInfo(signal.(func(bool) *QGeoPositionInfo)(int8(fromSatellitePositioningMethodsOnly) != 0))
	}

	return PointerFromQGeoPositionInfo(NewQNmeaPositionInfoSourceFromPointer(ptr).LastKnownPositionDefault(int8(fromSatellitePositioningMethodsOnly) != 0))
}

func (ptr *QNmeaPositionInfoSource) ConnectLastKnownPosition(f func(fromSatellitePositioningMethodsOnly bool) *QGeoPositionInfo) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "lastKnownPosition"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "lastKnownPosition", func(fromSatellitePositioningMethodsOnly bool) *QGeoPositionInfo {
				signal.(func(bool) *QGeoPositionInfo)(fromSatellitePositioningMethodsOnly)
				return f(fromSatellitePositioningMethodsOnly)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "lastKnownPosition", f)
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
		var tmpValue = NewQGeoPositionInfoFromPointer(C.QNmeaPositionInfoSource_LastKnownPosition(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(fromSatellitePositioningMethodsOnly)))))
		runtime.SetFinalizer(tmpValue, (*QGeoPositionInfo).DestroyQGeoPositionInfo)
		return tmpValue
	}
	return nil
}

func (ptr *QNmeaPositionInfoSource) LastKnownPositionDefault(fromSatellitePositioningMethodsOnly bool) *QGeoPositionInfo {
	if ptr.Pointer() != nil {
		var tmpValue = NewQGeoPositionInfoFromPointer(C.QNmeaPositionInfoSource_LastKnownPositionDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(fromSatellitePositioningMethodsOnly)))))
		runtime.SetFinalizer(tmpValue, (*QGeoPositionInfo).DestroyQGeoPositionInfo)
		return tmpValue
	}
	return nil
}

func (ptr *QNmeaPositionInfoSource) Device() *core.QIODevice {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQIODeviceFromPointer(C.QNmeaPositionInfoSource_Device(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
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

//export callbackQNmeaPositionInfoSource_MinimumUpdateInterval
func callbackQNmeaPositionInfoSource_MinimumUpdateInterval(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "minimumUpdateInterval"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(NewQNmeaPositionInfoSourceFromPointer(ptr).MinimumUpdateIntervalDefault()))
}

func (ptr *QNmeaPositionInfoSource) ConnectMinimumUpdateInterval(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "minimumUpdateInterval"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "minimumUpdateInterval", func() int {
				signal.(func() int)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "minimumUpdateInterval", f)
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
