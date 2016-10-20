// +build !minimal

package positioning

//#include <stdint.h>
//#include <stdlib.h>
//#include "positioning.h"
import "C"
import (
	"fmt"
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"runtime"
	"strings"
	"unsafe"
)

type QGeoAddress struct {
	ptr unsafe.Pointer
}

type QGeoAddress_ITF interface {
	QGeoAddress_PTR() *QGeoAddress
}

func (p *QGeoAddress) QGeoAddress_PTR() *QGeoAddress {
	return p
}

func (p *QGeoAddress) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QGeoAddress) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
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

func (ptr *QGeoAddress) City() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoAddress_City(ptr.Pointer()))
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
		return C.GoString(C.QGeoAddress_Country(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoAddress) CountryCode() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoAddress_CountryCode(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoAddress) County() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoAddress_County(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoAddress) District() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoAddress_District(ptr.Pointer()))
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

func (ptr *QGeoAddress) PostalCode() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoAddress_PostalCode(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoAddress) SetCity(city string) {
	if ptr.Pointer() != nil {
		var cityC = C.CString(city)
		defer C.free(unsafe.Pointer(cityC))
		C.QGeoAddress_SetCity(ptr.Pointer(), cityC)
	}
}

func (ptr *QGeoAddress) SetCountry(country string) {
	if ptr.Pointer() != nil {
		var countryC = C.CString(country)
		defer C.free(unsafe.Pointer(countryC))
		C.QGeoAddress_SetCountry(ptr.Pointer(), countryC)
	}
}

func (ptr *QGeoAddress) SetCountryCode(countryCode string) {
	if ptr.Pointer() != nil {
		var countryCodeC = C.CString(countryCode)
		defer C.free(unsafe.Pointer(countryCodeC))
		C.QGeoAddress_SetCountryCode(ptr.Pointer(), countryCodeC)
	}
}

func (ptr *QGeoAddress) SetCounty(county string) {
	if ptr.Pointer() != nil {
		var countyC = C.CString(county)
		defer C.free(unsafe.Pointer(countyC))
		C.QGeoAddress_SetCounty(ptr.Pointer(), countyC)
	}
}

func (ptr *QGeoAddress) SetDistrict(district string) {
	if ptr.Pointer() != nil {
		var districtC = C.CString(district)
		defer C.free(unsafe.Pointer(districtC))
		C.QGeoAddress_SetDistrict(ptr.Pointer(), districtC)
	}
}

func (ptr *QGeoAddress) SetPostalCode(postalCode string) {
	if ptr.Pointer() != nil {
		var postalCodeC = C.CString(postalCode)
		defer C.free(unsafe.Pointer(postalCodeC))
		C.QGeoAddress_SetPostalCode(ptr.Pointer(), postalCodeC)
	}
}

func (ptr *QGeoAddress) SetState(state string) {
	if ptr.Pointer() != nil {
		var stateC = C.CString(state)
		defer C.free(unsafe.Pointer(stateC))
		C.QGeoAddress_SetState(ptr.Pointer(), stateC)
	}
}

func (ptr *QGeoAddress) SetStreet(street string) {
	if ptr.Pointer() != nil {
		var streetC = C.CString(street)
		defer C.free(unsafe.Pointer(streetC))
		C.QGeoAddress_SetStreet(ptr.Pointer(), streetC)
	}
}

func (ptr *QGeoAddress) SetText(text string) {
	if ptr.Pointer() != nil {
		var textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
		C.QGeoAddress_SetText(ptr.Pointer(), textC)
	}
}

func (ptr *QGeoAddress) State() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoAddress_State(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoAddress) Street() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoAddress_Street(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoAddress) Text() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoAddress_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoAddress) DestroyQGeoAddress() {
	if ptr.Pointer() != nil {
		C.QGeoAddress_DestroyQGeoAddress(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QGeoAreaMonitorInfo struct {
	ptr unsafe.Pointer
}

type QGeoAreaMonitorInfo_ITF interface {
	QGeoAreaMonitorInfo_PTR() *QGeoAreaMonitorInfo
}

func (p *QGeoAreaMonitorInfo) QGeoAreaMonitorInfo_PTR() *QGeoAreaMonitorInfo {
	return p
}

func (p *QGeoAreaMonitorInfo) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QGeoAreaMonitorInfo) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
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
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = NewQGeoAreaMonitorInfoFromPointer(C.QGeoAreaMonitorInfo_NewQGeoAreaMonitorInfo(nameC))
	runtime.SetFinalizer(tmpValue, (*QGeoAreaMonitorInfo).DestroyQGeoAreaMonitorInfo)
	return tmpValue
}

func (ptr *QGeoAreaMonitorInfo) Area() *QGeoShape {
	if ptr.Pointer() != nil {
		var tmpValue = NewQGeoShapeFromPointer(C.QGeoAreaMonitorInfo_Area(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QGeoShape).DestroyQGeoShape)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoAreaMonitorInfo) Expiration() *core.QDateTime {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQDateTimeFromPointer(C.QGeoAreaMonitorInfo_Expiration(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QDateTime).DestroyQDateTime)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoAreaMonitorInfo) Identifier() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoAreaMonitorInfo_Identifier(ptr.Pointer()))
	}
	return ""
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

func (ptr *QGeoAreaMonitorInfo) Name() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoAreaMonitorInfo_Name(ptr.Pointer()))
	}
	return ""
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
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QGeoAreaMonitorInfo_SetName(ptr.Pointer(), nameC)
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
	}
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

type QGeoAreaMonitorSource struct {
	core.QObject
}

type QGeoAreaMonitorSource_ITF interface {
	core.QObject_ITF
	QGeoAreaMonitorSource_PTR() *QGeoAreaMonitorSource
}

func (p *QGeoAreaMonitorSource) QGeoAreaMonitorSource_PTR() *QGeoAreaMonitorSource {
	return p
}

func (p *QGeoAreaMonitorSource) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QGeoAreaMonitorSource) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
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

//export callbackQGeoAreaMonitorSource_AreaEntered
func callbackQGeoAreaMonitorSource_AreaEntered(ptr unsafe.Pointer, monitor unsafe.Pointer, update unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoAreaMonitorSource::areaEntered"); signal != nil {
		signal.(func(*QGeoAreaMonitorInfo, *QGeoPositionInfo))(NewQGeoAreaMonitorInfoFromPointer(monitor), NewQGeoPositionInfoFromPointer(update))
	}

}

func (ptr *QGeoAreaMonitorSource) ConnectAreaEntered(f func(monitor *QGeoAreaMonitorInfo, update *QGeoPositionInfo)) {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_ConnectAreaEntered(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoAreaMonitorSource::areaEntered", f)
	}
}

func (ptr *QGeoAreaMonitorSource) DisconnectAreaEntered() {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_DisconnectAreaEntered(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoAreaMonitorSource::areaEntered")
	}
}

func (ptr *QGeoAreaMonitorSource) AreaEntered(monitor QGeoAreaMonitorInfo_ITF, update QGeoPositionInfo_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_AreaEntered(ptr.Pointer(), PointerFromQGeoAreaMonitorInfo(monitor), PointerFromQGeoPositionInfo(update))
	}
}

//export callbackQGeoAreaMonitorSource_AreaExited
func callbackQGeoAreaMonitorSource_AreaExited(ptr unsafe.Pointer, monitor unsafe.Pointer, update unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoAreaMonitorSource::areaExited"); signal != nil {
		signal.(func(*QGeoAreaMonitorInfo, *QGeoPositionInfo))(NewQGeoAreaMonitorInfoFromPointer(monitor), NewQGeoPositionInfoFromPointer(update))
	}

}

func (ptr *QGeoAreaMonitorSource) ConnectAreaExited(f func(monitor *QGeoAreaMonitorInfo, update *QGeoPositionInfo)) {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_ConnectAreaExited(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoAreaMonitorSource::areaExited", f)
	}
}

func (ptr *QGeoAreaMonitorSource) DisconnectAreaExited() {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_DisconnectAreaExited(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoAreaMonitorSource::areaExited")
	}
}

func (ptr *QGeoAreaMonitorSource) AreaExited(monitor QGeoAreaMonitorInfo_ITF, update QGeoPositionInfo_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_AreaExited(ptr.Pointer(), PointerFromQGeoAreaMonitorInfo(monitor), PointerFromQGeoPositionInfo(update))
	}
}

func QGeoAreaMonitorSource_AvailableSources() []string {
	return strings.Split(C.GoString(C.QGeoAreaMonitorSource_QGeoAreaMonitorSource_AvailableSources()), "|")
}

func (ptr *QGeoAreaMonitorSource) AvailableSources() []string {
	return strings.Split(C.GoString(C.QGeoAreaMonitorSource_QGeoAreaMonitorSource_AvailableSources()), "|")
}

func QGeoAreaMonitorSource_CreateDefaultSource(parent core.QObject_ITF) *QGeoAreaMonitorSource {
	var tmpValue = NewQGeoAreaMonitorSourceFromPointer(C.QGeoAreaMonitorSource_QGeoAreaMonitorSource_CreateDefaultSource(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QGeoAreaMonitorSource) CreateDefaultSource(parent core.QObject_ITF) *QGeoAreaMonitorSource {
	var tmpValue = NewQGeoAreaMonitorSourceFromPointer(C.QGeoAreaMonitorSource_QGeoAreaMonitorSource_CreateDefaultSource(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QGeoAreaMonitorSource_CreateSource(sourceName string, parent core.QObject_ITF) *QGeoAreaMonitorSource {
	var sourceNameC = C.CString(sourceName)
	defer C.free(unsafe.Pointer(sourceNameC))
	var tmpValue = NewQGeoAreaMonitorSourceFromPointer(C.QGeoAreaMonitorSource_QGeoAreaMonitorSource_CreateSource(sourceNameC, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QGeoAreaMonitorSource) CreateSource(sourceName string, parent core.QObject_ITF) *QGeoAreaMonitorSource {
	var sourceNameC = C.CString(sourceName)
	defer C.free(unsafe.Pointer(sourceNameC))
	var tmpValue = NewQGeoAreaMonitorSourceFromPointer(C.QGeoAreaMonitorSource_QGeoAreaMonitorSource_CreateSource(sourceNameC, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQGeoAreaMonitorSource_Error2
func callbackQGeoAreaMonitorSource_Error2(ptr unsafe.Pointer, areaMonitoringError C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoAreaMonitorSource::error2"); signal != nil {
		signal.(func(QGeoAreaMonitorSource__Error))(QGeoAreaMonitorSource__Error(areaMonitoringError))
	}

}

func (ptr *QGeoAreaMonitorSource) ConnectError2(f func(areaMonitoringError QGeoAreaMonitorSource__Error)) {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoAreaMonitorSource::error2", f)
	}
}

func (ptr *QGeoAreaMonitorSource) DisconnectError2() {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoAreaMonitorSource::error2")
	}
}

func (ptr *QGeoAreaMonitorSource) Error2(areaMonitoringError QGeoAreaMonitorSource__Error) {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_Error2(ptr.Pointer(), C.longlong(areaMonitoringError))
	}
}

//export callbackQGeoAreaMonitorSource_Error
func callbackQGeoAreaMonitorSource_Error(ptr unsafe.Pointer) C.longlong {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoAreaMonitorSource::error"); signal != nil {
		return C.longlong(signal.(func() QGeoAreaMonitorSource__Error)())
	}

	return C.longlong(0)
}

func (ptr *QGeoAreaMonitorSource) ConnectError(f func() QGeoAreaMonitorSource__Error) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoAreaMonitorSource::error", f)
	}
}

func (ptr *QGeoAreaMonitorSource) DisconnectError() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoAreaMonitorSource::error")
	}
}

func (ptr *QGeoAreaMonitorSource) Error() QGeoAreaMonitorSource__Error {
	if ptr.Pointer() != nil {
		return QGeoAreaMonitorSource__Error(C.QGeoAreaMonitorSource_Error(ptr.Pointer()))
	}
	return 0
}

//export callbackQGeoAreaMonitorSource_MonitorExpired
func callbackQGeoAreaMonitorSource_MonitorExpired(ptr unsafe.Pointer, monitor unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoAreaMonitorSource::monitorExpired"); signal != nil {
		signal.(func(*QGeoAreaMonitorInfo))(NewQGeoAreaMonitorInfoFromPointer(monitor))
	}

}

func (ptr *QGeoAreaMonitorSource) ConnectMonitorExpired(f func(monitor *QGeoAreaMonitorInfo)) {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_ConnectMonitorExpired(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoAreaMonitorSource::monitorExpired", f)
	}
}

func (ptr *QGeoAreaMonitorSource) DisconnectMonitorExpired() {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_DisconnectMonitorExpired(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoAreaMonitorSource::monitorExpired")
	}
}

func (ptr *QGeoAreaMonitorSource) MonitorExpired(monitor QGeoAreaMonitorInfo_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_MonitorExpired(ptr.Pointer(), PointerFromQGeoAreaMonitorInfo(monitor))
	}
}

//export callbackQGeoAreaMonitorSource_PositionInfoSource
func callbackQGeoAreaMonitorSource_PositionInfoSource(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoAreaMonitorSource::positionInfoSource"); signal != nil {
		return PointerFromQGeoPositionInfoSource(signal.(func() *QGeoPositionInfoSource)())
	}

	return PointerFromQGeoPositionInfoSource(NewQGeoAreaMonitorSourceFromPointer(ptr).PositionInfoSourceDefault())
}

func (ptr *QGeoAreaMonitorSource) ConnectPositionInfoSource(f func() *QGeoPositionInfoSource) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoAreaMonitorSource::positionInfoSource", f)
	}
}

func (ptr *QGeoAreaMonitorSource) DisconnectPositionInfoSource() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoAreaMonitorSource::positionInfoSource")
	}
}

func (ptr *QGeoAreaMonitorSource) PositionInfoSource() *QGeoPositionInfoSource {
	if ptr.Pointer() != nil {
		var tmpValue = NewQGeoPositionInfoSourceFromPointer(C.QGeoAreaMonitorSource_PositionInfoSource(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGeoAreaMonitorSource) PositionInfoSourceDefault() *QGeoPositionInfoSource {
	if ptr.Pointer() != nil {
		var tmpValue = NewQGeoPositionInfoSourceFromPointer(C.QGeoAreaMonitorSource_PositionInfoSourceDefault(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQGeoAreaMonitorSource_RequestUpdate
func callbackQGeoAreaMonitorSource_RequestUpdate(ptr unsafe.Pointer, monitor unsafe.Pointer, sign *C.char) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoAreaMonitorSource::requestUpdate"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QGeoAreaMonitorInfo, string) bool)(NewQGeoAreaMonitorInfoFromPointer(monitor), C.GoString(sign)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QGeoAreaMonitorSource) ConnectRequestUpdate(f func(monitor *QGeoAreaMonitorInfo, sign string) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoAreaMonitorSource::requestUpdate", f)
	}
}

func (ptr *QGeoAreaMonitorSource) DisconnectRequestUpdate(monitor QGeoAreaMonitorInfo_ITF, sign string) {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoAreaMonitorSource::requestUpdate")
	}
}

func (ptr *QGeoAreaMonitorSource) RequestUpdate(monitor QGeoAreaMonitorInfo_ITF, sign string) bool {
	if ptr.Pointer() != nil {
		var signC = C.CString(sign)
		defer C.free(unsafe.Pointer(signC))
		return C.QGeoAreaMonitorSource_RequestUpdate(ptr.Pointer(), PointerFromQGeoAreaMonitorInfo(monitor), signC) != 0
	}
	return false
}

//export callbackQGeoAreaMonitorSource_SetPositionInfoSource
func callbackQGeoAreaMonitorSource_SetPositionInfoSource(ptr unsafe.Pointer, newSource unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoAreaMonitorSource::setPositionInfoSource"); signal != nil {
		signal.(func(*QGeoPositionInfoSource))(NewQGeoPositionInfoSourceFromPointer(newSource))
	} else {
		NewQGeoAreaMonitorSourceFromPointer(ptr).SetPositionInfoSourceDefault(NewQGeoPositionInfoSourceFromPointer(newSource))
	}
}

func (ptr *QGeoAreaMonitorSource) ConnectSetPositionInfoSource(f func(newSource *QGeoPositionInfoSource)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoAreaMonitorSource::setPositionInfoSource", f)
	}
}

func (ptr *QGeoAreaMonitorSource) DisconnectSetPositionInfoSource() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoAreaMonitorSource::setPositionInfoSource")
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
		return C.GoString(C.QGeoAreaMonitorSource_SourceName(ptr.Pointer()))
	}
	return ""
}

//export callbackQGeoAreaMonitorSource_StartMonitoring
func callbackQGeoAreaMonitorSource_StartMonitoring(ptr unsafe.Pointer, monitor unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoAreaMonitorSource::startMonitoring"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QGeoAreaMonitorInfo) bool)(NewQGeoAreaMonitorInfoFromPointer(monitor)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QGeoAreaMonitorSource) ConnectStartMonitoring(f func(monitor *QGeoAreaMonitorInfo) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoAreaMonitorSource::startMonitoring", f)
	}
}

func (ptr *QGeoAreaMonitorSource) DisconnectStartMonitoring(monitor QGeoAreaMonitorInfo_ITF) {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoAreaMonitorSource::startMonitoring")
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

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoAreaMonitorSource::stopMonitoring"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QGeoAreaMonitorInfo) bool)(NewQGeoAreaMonitorInfoFromPointer(monitor)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QGeoAreaMonitorSource) ConnectStopMonitoring(f func(monitor *QGeoAreaMonitorInfo) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoAreaMonitorSource::stopMonitoring", f)
	}
}

func (ptr *QGeoAreaMonitorSource) DisconnectStopMonitoring(monitor QGeoAreaMonitorInfo_ITF) {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoAreaMonitorSource::stopMonitoring")
	}
}

func (ptr *QGeoAreaMonitorSource) StopMonitoring(monitor QGeoAreaMonitorInfo_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGeoAreaMonitorSource_StopMonitoring(ptr.Pointer(), PointerFromQGeoAreaMonitorInfo(monitor)) != 0
	}
	return false
}

//export callbackQGeoAreaMonitorSource_SupportedAreaMonitorFeatures
func callbackQGeoAreaMonitorSource_SupportedAreaMonitorFeatures(ptr unsafe.Pointer) C.longlong {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoAreaMonitorSource::supportedAreaMonitorFeatures"); signal != nil {
		return C.longlong(signal.(func() QGeoAreaMonitorSource__AreaMonitorFeature)())
	}

	return C.longlong(0)
}

func (ptr *QGeoAreaMonitorSource) ConnectSupportedAreaMonitorFeatures(f func() QGeoAreaMonitorSource__AreaMonitorFeature) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoAreaMonitorSource::supportedAreaMonitorFeatures", f)
	}
}

func (ptr *QGeoAreaMonitorSource) DisconnectSupportedAreaMonitorFeatures() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoAreaMonitorSource::supportedAreaMonitorFeatures")
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

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoAreaMonitorSource::~QGeoAreaMonitorSource"); signal != nil {
		signal.(func())()
	} else {
		NewQGeoAreaMonitorSourceFromPointer(ptr).DestroyQGeoAreaMonitorSourceDefault()
	}
}

func (ptr *QGeoAreaMonitorSource) ConnectDestroyQGeoAreaMonitorSource(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoAreaMonitorSource::~QGeoAreaMonitorSource", f)
	}
}

func (ptr *QGeoAreaMonitorSource) DisconnectDestroyQGeoAreaMonitorSource() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoAreaMonitorSource::~QGeoAreaMonitorSource")
	}
}

func (ptr *QGeoAreaMonitorSource) DestroyQGeoAreaMonitorSource() {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_DestroyQGeoAreaMonitorSource(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QGeoAreaMonitorSource) DestroyQGeoAreaMonitorSourceDefault() {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_DestroyQGeoAreaMonitorSourceDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQGeoAreaMonitorSource_TimerEvent
func callbackQGeoAreaMonitorSource_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoAreaMonitorSource::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGeoAreaMonitorSourceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGeoAreaMonitorSource) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoAreaMonitorSource::timerEvent", f)
	}
}

func (ptr *QGeoAreaMonitorSource) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoAreaMonitorSource::timerEvent")
	}
}

func (ptr *QGeoAreaMonitorSource) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGeoAreaMonitorSource) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQGeoAreaMonitorSource_ChildEvent
func callbackQGeoAreaMonitorSource_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoAreaMonitorSource::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQGeoAreaMonitorSourceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QGeoAreaMonitorSource) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoAreaMonitorSource::childEvent", f)
	}
}

func (ptr *QGeoAreaMonitorSource) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoAreaMonitorSource::childEvent")
	}
}

func (ptr *QGeoAreaMonitorSource) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGeoAreaMonitorSource) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQGeoAreaMonitorSource_ConnectNotify
func callbackQGeoAreaMonitorSource_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoAreaMonitorSource::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGeoAreaMonitorSourceFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGeoAreaMonitorSource) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoAreaMonitorSource::connectNotify", f)
	}
}

func (ptr *QGeoAreaMonitorSource) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoAreaMonitorSource::connectNotify")
	}
}

func (ptr *QGeoAreaMonitorSource) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QGeoAreaMonitorSource) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGeoAreaMonitorSource_CustomEvent
func callbackQGeoAreaMonitorSource_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoAreaMonitorSource::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGeoAreaMonitorSourceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGeoAreaMonitorSource) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoAreaMonitorSource::customEvent", f)
	}
}

func (ptr *QGeoAreaMonitorSource) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoAreaMonitorSource::customEvent")
	}
}

func (ptr *QGeoAreaMonitorSource) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGeoAreaMonitorSource) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQGeoAreaMonitorSource_DeleteLater
func callbackQGeoAreaMonitorSource_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoAreaMonitorSource::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQGeoAreaMonitorSourceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QGeoAreaMonitorSource) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoAreaMonitorSource::deleteLater", f)
	}
}

func (ptr *QGeoAreaMonitorSource) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoAreaMonitorSource::deleteLater")
	}
}

func (ptr *QGeoAreaMonitorSource) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QGeoAreaMonitorSource) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQGeoAreaMonitorSource_DisconnectNotify
func callbackQGeoAreaMonitorSource_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoAreaMonitorSource::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGeoAreaMonitorSourceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGeoAreaMonitorSource) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoAreaMonitorSource::disconnectNotify", f)
	}
}

func (ptr *QGeoAreaMonitorSource) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoAreaMonitorSource::disconnectNotify")
	}
}

func (ptr *QGeoAreaMonitorSource) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QGeoAreaMonitorSource) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorSource_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGeoAreaMonitorSource_Event
func callbackQGeoAreaMonitorSource_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoAreaMonitorSource::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGeoAreaMonitorSourceFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QGeoAreaMonitorSource) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoAreaMonitorSource::event", f)
	}
}

func (ptr *QGeoAreaMonitorSource) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoAreaMonitorSource::event")
	}
}

func (ptr *QGeoAreaMonitorSource) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGeoAreaMonitorSource_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QGeoAreaMonitorSource) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGeoAreaMonitorSource_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQGeoAreaMonitorSource_EventFilter
func callbackQGeoAreaMonitorSource_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoAreaMonitorSource::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGeoAreaMonitorSourceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QGeoAreaMonitorSource) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoAreaMonitorSource::eventFilter", f)
	}
}

func (ptr *QGeoAreaMonitorSource) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoAreaMonitorSource::eventFilter")
	}
}

func (ptr *QGeoAreaMonitorSource) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGeoAreaMonitorSource_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QGeoAreaMonitorSource) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGeoAreaMonitorSource_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQGeoAreaMonitorSource_MetaObject
func callbackQGeoAreaMonitorSource_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoAreaMonitorSource::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQGeoAreaMonitorSourceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QGeoAreaMonitorSource) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoAreaMonitorSource::metaObject", f)
	}
}

func (ptr *QGeoAreaMonitorSource) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoAreaMonitorSource::metaObject")
	}
}

func (ptr *QGeoAreaMonitorSource) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QGeoAreaMonitorSource_MetaObject(ptr.Pointer()))
	}
	return nil
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

func (p *QGeoCircle) QGeoCircle_PTR() *QGeoCircle {
	return p
}

func (p *QGeoCircle) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QGeoShape_PTR().Pointer()
	}
	return nil
}

func (p *QGeoCircle) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QGeoShape_PTR().SetPointer(ptr)
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

func (ptr *QGeoCircle) Center() *QGeoCoordinate {
	if ptr.Pointer() != nil {
		var tmpValue = NewQGeoCoordinateFromPointer(C.QGeoCircle_Center(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QGeoCoordinate).DestroyQGeoCoordinate)
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

func (ptr *QGeoCircle) ToString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoCircle_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoCircle) Translate(degreesLatitude float64, degreesLongitude float64) {
	if ptr.Pointer() != nil {
		C.QGeoCircle_Translate(ptr.Pointer(), C.double(degreesLatitude), C.double(degreesLongitude))
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

func (ptr *QGeoCircle) DestroyQGeoCircle() {
	if ptr.Pointer() != nil {
		C.QGeoCircle_DestroyQGeoCircle(ptr.Pointer())
		ptr.SetPointer(nil)
	}
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

type QGeoCoordinate struct {
	ptr unsafe.Pointer
}

type QGeoCoordinate_ITF interface {
	QGeoCoordinate_PTR() *QGeoCoordinate
}

func (p *QGeoCoordinate) QGeoCoordinate_PTR() *QGeoCoordinate {
	return p
}

func (p *QGeoCoordinate) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QGeoCoordinate) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
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

func (ptr *QGeoCoordinate) Altitude() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGeoCoordinate_Altitude(ptr.Pointer()))
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
		return C.QGeoCoordinate_IsValid(ptr.Pointer()) != 0
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
		return C.GoString(C.QGeoCoordinate_ToString(ptr.Pointer(), C.longlong(format)))
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
		C.QGeoCoordinate_DestroyQGeoCoordinate(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QGeoLocation struct {
	ptr unsafe.Pointer
}

type QGeoLocation_ITF interface {
	QGeoLocation_PTR() *QGeoLocation
}

func (p *QGeoLocation) QGeoLocation_PTR() *QGeoLocation {
	return p
}

func (p *QGeoLocation) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QGeoLocation) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
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

type QGeoPositionInfo struct {
	ptr unsafe.Pointer
}

type QGeoPositionInfo_ITF interface {
	QGeoPositionInfo_PTR() *QGeoPositionInfo
}

func (p *QGeoPositionInfo) QGeoPositionInfo_PTR() *QGeoPositionInfo {
	return p
}

func (p *QGeoPositionInfo) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QGeoPositionInfo) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
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

func (ptr *QGeoPositionInfo) Attribute(attribute QGeoPositionInfo__Attribute) float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGeoPositionInfo_Attribute(ptr.Pointer(), C.longlong(attribute)))
	}
	return 0
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
		var tmpValue = core.NewQDateTimeFromPointer(C.QGeoPositionInfo_Timestamp(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QDateTime).DestroyQDateTime)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoPositionInfo) DestroyQGeoPositionInfo() {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfo_DestroyQGeoPositionInfo(ptr.Pointer())
		ptr.SetPointer(nil)
	}
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

type QGeoPositionInfoSource struct {
	core.QObject
}

type QGeoPositionInfoSource_ITF interface {
	core.QObject_ITF
	QGeoPositionInfoSource_PTR() *QGeoPositionInfoSource
}

func (p *QGeoPositionInfoSource) QGeoPositionInfoSource_PTR() *QGeoPositionInfoSource {
	return p
}

func (p *QGeoPositionInfoSource) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QGeoPositionInfoSource) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
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

//export callbackQGeoPositionInfoSource_SetUpdateInterval
func callbackQGeoPositionInfoSource_SetUpdateInterval(ptr unsafe.Pointer, msec C.int) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoPositionInfoSource::setUpdateInterval"); signal != nil {
		signal.(func(int))(int(int32(msec)))
	} else {
		NewQGeoPositionInfoSourceFromPointer(ptr).SetUpdateIntervalDefault(int(int32(msec)))
	}
}

func (ptr *QGeoPositionInfoSource) ConnectSetUpdateInterval(f func(msec int)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSource::setUpdateInterval", f)
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectSetUpdateInterval() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSource::setUpdateInterval")
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
		return C.GoString(C.QGeoPositionInfoSource_SourceName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoPositionInfoSource) UpdateInterval() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QGeoPositionInfoSource_UpdateInterval(ptr.Pointer())))
	}
	return 0
}

func NewQGeoPositionInfoSource(parent core.QObject_ITF) *QGeoPositionInfoSource {
	var tmpValue = NewQGeoPositionInfoSourceFromPointer(C.QGeoPositionInfoSource_NewQGeoPositionInfoSource(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QGeoPositionInfoSource_AvailableSources() []string {
	return strings.Split(C.GoString(C.QGeoPositionInfoSource_QGeoPositionInfoSource_AvailableSources()), "|")
}

func (ptr *QGeoPositionInfoSource) AvailableSources() []string {
	return strings.Split(C.GoString(C.QGeoPositionInfoSource_QGeoPositionInfoSource_AvailableSources()), "|")
}

func QGeoPositionInfoSource_CreateDefaultSource(parent core.QObject_ITF) *QGeoPositionInfoSource {
	var tmpValue = NewQGeoPositionInfoSourceFromPointer(C.QGeoPositionInfoSource_QGeoPositionInfoSource_CreateDefaultSource(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QGeoPositionInfoSource) CreateDefaultSource(parent core.QObject_ITF) *QGeoPositionInfoSource {
	var tmpValue = NewQGeoPositionInfoSourceFromPointer(C.QGeoPositionInfoSource_QGeoPositionInfoSource_CreateDefaultSource(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QGeoPositionInfoSource_CreateSource(sourceName string, parent core.QObject_ITF) *QGeoPositionInfoSource {
	var sourceNameC = C.CString(sourceName)
	defer C.free(unsafe.Pointer(sourceNameC))
	var tmpValue = NewQGeoPositionInfoSourceFromPointer(C.QGeoPositionInfoSource_QGeoPositionInfoSource_CreateSource(sourceNameC, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QGeoPositionInfoSource) CreateSource(sourceName string, parent core.QObject_ITF) *QGeoPositionInfoSource {
	var sourceNameC = C.CString(sourceName)
	defer C.free(unsafe.Pointer(sourceNameC))
	var tmpValue = NewQGeoPositionInfoSourceFromPointer(C.QGeoPositionInfoSource_QGeoPositionInfoSource_CreateSource(sourceNameC, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQGeoPositionInfoSource_Error2
func callbackQGeoPositionInfoSource_Error2(ptr unsafe.Pointer, positioningError C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoPositionInfoSource::error2"); signal != nil {
		signal.(func(QGeoPositionInfoSource__Error))(QGeoPositionInfoSource__Error(positioningError))
	}

}

func (ptr *QGeoPositionInfoSource) ConnectError2(f func(positioningError QGeoPositionInfoSource__Error)) {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSource::error2", f)
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectError2() {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSource::error2")
	}
}

func (ptr *QGeoPositionInfoSource) Error2(positioningError QGeoPositionInfoSource__Error) {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_Error2(ptr.Pointer(), C.longlong(positioningError))
	}
}

//export callbackQGeoPositionInfoSource_Error
func callbackQGeoPositionInfoSource_Error(ptr unsafe.Pointer) C.longlong {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoPositionInfoSource::error"); signal != nil {
		return C.longlong(signal.(func() QGeoPositionInfoSource__Error)())
	}

	return C.longlong(0)
}

func (ptr *QGeoPositionInfoSource) ConnectError(f func() QGeoPositionInfoSource__Error) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSource::error", f)
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectError() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSource::error")
	}
}

func (ptr *QGeoPositionInfoSource) Error() QGeoPositionInfoSource__Error {
	if ptr.Pointer() != nil {
		return QGeoPositionInfoSource__Error(C.QGeoPositionInfoSource_Error(ptr.Pointer()))
	}
	return 0
}

//export callbackQGeoPositionInfoSource_LastKnownPosition
func callbackQGeoPositionInfoSource_LastKnownPosition(ptr unsafe.Pointer, fromSatellitePositioningMethodsOnly C.char) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoPositionInfoSource::lastKnownPosition"); signal != nil {
		return PointerFromQGeoPositionInfo(signal.(func(bool) *QGeoPositionInfo)(int8(fromSatellitePositioningMethodsOnly) != 0))
	}

	return PointerFromQGeoPositionInfo(nil)
}

func (ptr *QGeoPositionInfoSource) ConnectLastKnownPosition(f func(fromSatellitePositioningMethodsOnly bool) *QGeoPositionInfo) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSource::lastKnownPosition", f)
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectLastKnownPosition(fromSatellitePositioningMethodsOnly bool) {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSource::lastKnownPosition")
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

//export callbackQGeoPositionInfoSource_MinimumUpdateInterval
func callbackQGeoPositionInfoSource_MinimumUpdateInterval(ptr unsafe.Pointer) C.int {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoPositionInfoSource::minimumUpdateInterval"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(0))
}

func (ptr *QGeoPositionInfoSource) ConnectMinimumUpdateInterval(f func() int) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSource::minimumUpdateInterval", f)
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectMinimumUpdateInterval() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSource::minimumUpdateInterval")
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

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoPositionInfoSource::positionUpdated"); signal != nil {
		signal.(func(*QGeoPositionInfo))(NewQGeoPositionInfoFromPointer(update))
	}

}

func (ptr *QGeoPositionInfoSource) ConnectPositionUpdated(f func(update *QGeoPositionInfo)) {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_ConnectPositionUpdated(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSource::positionUpdated", f)
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectPositionUpdated() {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_DisconnectPositionUpdated(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSource::positionUpdated")
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
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoPositionInfoSource::requestUpdate"); signal != nil {
		signal.(func(int))(int(int32(timeout)))
	}

}

func (ptr *QGeoPositionInfoSource) ConnectRequestUpdate(f func(timeout int)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSource::requestUpdate", f)
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectRequestUpdate(timeout int) {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSource::requestUpdate")
	}
}

func (ptr *QGeoPositionInfoSource) RequestUpdate(timeout int) {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_RequestUpdate(ptr.Pointer(), C.int(int32(timeout)))
	}
}

//export callbackQGeoPositionInfoSource_SetPreferredPositioningMethods
func callbackQGeoPositionInfoSource_SetPreferredPositioningMethods(ptr unsafe.Pointer, methods C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoPositionInfoSource::setPreferredPositioningMethods"); signal != nil {
		signal.(func(QGeoPositionInfoSource__PositioningMethod))(QGeoPositionInfoSource__PositioningMethod(methods))
	} else {
		NewQGeoPositionInfoSourceFromPointer(ptr).SetPreferredPositioningMethodsDefault(QGeoPositionInfoSource__PositioningMethod(methods))
	}
}

func (ptr *QGeoPositionInfoSource) ConnectSetPreferredPositioningMethods(f func(methods QGeoPositionInfoSource__PositioningMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSource::setPreferredPositioningMethods", f)
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectSetPreferredPositioningMethods() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSource::setPreferredPositioningMethods")
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
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoPositionInfoSource::startUpdates"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QGeoPositionInfoSource) ConnectStartUpdates(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSource::startUpdates", f)
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectStartUpdates() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSource::startUpdates")
	}
}

func (ptr *QGeoPositionInfoSource) StartUpdates() {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_StartUpdates(ptr.Pointer())
	}
}

//export callbackQGeoPositionInfoSource_StopUpdates
func callbackQGeoPositionInfoSource_StopUpdates(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoPositionInfoSource::stopUpdates"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QGeoPositionInfoSource) ConnectStopUpdates(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSource::stopUpdates", f)
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectStopUpdates() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSource::stopUpdates")
	}
}

func (ptr *QGeoPositionInfoSource) StopUpdates() {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_StopUpdates(ptr.Pointer())
	}
}

//export callbackQGeoPositionInfoSource_SupportedPositioningMethods
func callbackQGeoPositionInfoSource_SupportedPositioningMethods(ptr unsafe.Pointer) C.longlong {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoPositionInfoSource::supportedPositioningMethods"); signal != nil {
		return C.longlong(signal.(func() QGeoPositionInfoSource__PositioningMethod)())
	}

	return C.longlong(0)
}

func (ptr *QGeoPositionInfoSource) ConnectSupportedPositioningMethods(f func() QGeoPositionInfoSource__PositioningMethod) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSource::supportedPositioningMethods", f)
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectSupportedPositioningMethods() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSource::supportedPositioningMethods")
	}
}

func (ptr *QGeoPositionInfoSource) SupportedPositioningMethods() QGeoPositionInfoSource__PositioningMethod {
	if ptr.Pointer() != nil {
		return QGeoPositionInfoSource__PositioningMethod(C.QGeoPositionInfoSource_SupportedPositioningMethods(ptr.Pointer()))
	}
	return 0
}

//export callbackQGeoPositionInfoSource_UpdateTimeout
func callbackQGeoPositionInfoSource_UpdateTimeout(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoPositionInfoSource::updateTimeout"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QGeoPositionInfoSource) ConnectUpdateTimeout(f func()) {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_ConnectUpdateTimeout(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSource::updateTimeout", f)
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectUpdateTimeout() {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_DisconnectUpdateTimeout(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSource::updateTimeout")
	}
}

func (ptr *QGeoPositionInfoSource) UpdateTimeout() {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_UpdateTimeout(ptr.Pointer())
	}
}

//export callbackQGeoPositionInfoSource_DestroyQGeoPositionInfoSource
func callbackQGeoPositionInfoSource_DestroyQGeoPositionInfoSource(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoPositionInfoSource::~QGeoPositionInfoSource"); signal != nil {
		signal.(func())()
	} else {
		NewQGeoPositionInfoSourceFromPointer(ptr).DestroyQGeoPositionInfoSourceDefault()
	}
}

func (ptr *QGeoPositionInfoSource) ConnectDestroyQGeoPositionInfoSource(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSource::~QGeoPositionInfoSource", f)
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectDestroyQGeoPositionInfoSource() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSource::~QGeoPositionInfoSource")
	}
}

func (ptr *QGeoPositionInfoSource) DestroyQGeoPositionInfoSource() {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_DestroyQGeoPositionInfoSource(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QGeoPositionInfoSource) DestroyQGeoPositionInfoSourceDefault() {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_DestroyQGeoPositionInfoSourceDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQGeoPositionInfoSource_TimerEvent
func callbackQGeoPositionInfoSource_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoPositionInfoSource::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGeoPositionInfoSourceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGeoPositionInfoSource) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSource::timerEvent", f)
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSource::timerEvent")
	}
}

func (ptr *QGeoPositionInfoSource) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGeoPositionInfoSource) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQGeoPositionInfoSource_ChildEvent
func callbackQGeoPositionInfoSource_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoPositionInfoSource::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQGeoPositionInfoSourceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QGeoPositionInfoSource) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSource::childEvent", f)
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSource::childEvent")
	}
}

func (ptr *QGeoPositionInfoSource) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGeoPositionInfoSource) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQGeoPositionInfoSource_ConnectNotify
func callbackQGeoPositionInfoSource_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoPositionInfoSource::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGeoPositionInfoSourceFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGeoPositionInfoSource) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSource::connectNotify", f)
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSource::connectNotify")
	}
}

func (ptr *QGeoPositionInfoSource) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QGeoPositionInfoSource) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGeoPositionInfoSource_CustomEvent
func callbackQGeoPositionInfoSource_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoPositionInfoSource::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGeoPositionInfoSourceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGeoPositionInfoSource) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSource::customEvent", f)
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSource::customEvent")
	}
}

func (ptr *QGeoPositionInfoSource) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGeoPositionInfoSource) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQGeoPositionInfoSource_DeleteLater
func callbackQGeoPositionInfoSource_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoPositionInfoSource::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQGeoPositionInfoSourceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QGeoPositionInfoSource) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSource::deleteLater", f)
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSource::deleteLater")
	}
}

func (ptr *QGeoPositionInfoSource) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QGeoPositionInfoSource) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQGeoPositionInfoSource_DisconnectNotify
func callbackQGeoPositionInfoSource_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoPositionInfoSource::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGeoPositionInfoSourceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGeoPositionInfoSource) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSource::disconnectNotify", f)
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSource::disconnectNotify")
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSource_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGeoPositionInfoSource_Event
func callbackQGeoPositionInfoSource_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoPositionInfoSource::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGeoPositionInfoSourceFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QGeoPositionInfoSource) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSource::event", f)
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSource::event")
	}
}

func (ptr *QGeoPositionInfoSource) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGeoPositionInfoSource_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QGeoPositionInfoSource) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGeoPositionInfoSource_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQGeoPositionInfoSource_EventFilter
func callbackQGeoPositionInfoSource_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoPositionInfoSource::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGeoPositionInfoSourceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QGeoPositionInfoSource) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSource::eventFilter", f)
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSource::eventFilter")
	}
}

func (ptr *QGeoPositionInfoSource) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGeoPositionInfoSource_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QGeoPositionInfoSource) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGeoPositionInfoSource_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQGeoPositionInfoSource_MetaObject
func callbackQGeoPositionInfoSource_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoPositionInfoSource::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQGeoPositionInfoSourceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QGeoPositionInfoSource) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSource::metaObject", f)
	}
}

func (ptr *QGeoPositionInfoSource) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSource::metaObject")
	}
}

func (ptr *QGeoPositionInfoSource) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QGeoPositionInfoSource_MetaObject(ptr.Pointer()))
	}
	return nil
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

func (p *QGeoPositionInfoSourceFactory) QGeoPositionInfoSourceFactory_PTR() *QGeoPositionInfoSourceFactory {
	return p
}

func (p *QGeoPositionInfoSourceFactory) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QGeoPositionInfoSourceFactory) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
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

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoPositionInfoSourceFactory::areaMonitor"); signal != nil {
		return PointerFromQGeoAreaMonitorSource(signal.(func(*core.QObject) *QGeoAreaMonitorSource)(core.NewQObjectFromPointer(parent)))
	}

	return PointerFromQGeoAreaMonitorSource(nil)
}

func (ptr *QGeoPositionInfoSourceFactory) ConnectAreaMonitor(f func(parent *core.QObject) *QGeoAreaMonitorSource) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSourceFactory::areaMonitor", f)
	}
}

func (ptr *QGeoPositionInfoSourceFactory) DisconnectAreaMonitor(parent core.QObject_ITF) {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSourceFactory::areaMonitor")
	}
}

func (ptr *QGeoPositionInfoSourceFactory) AreaMonitor(parent core.QObject_ITF) *QGeoAreaMonitorSource {
	if ptr.Pointer() != nil {
		var tmpValue = NewQGeoAreaMonitorSourceFromPointer(C.QGeoPositionInfoSourceFactory_AreaMonitor(ptr.Pointer(), core.PointerFromQObject(parent)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQGeoPositionInfoSourceFactory_PositionInfoSource
func callbackQGeoPositionInfoSourceFactory_PositionInfoSource(ptr unsafe.Pointer, parent unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoPositionInfoSourceFactory::positionInfoSource"); signal != nil {
		return PointerFromQGeoPositionInfoSource(signal.(func(*core.QObject) *QGeoPositionInfoSource)(core.NewQObjectFromPointer(parent)))
	}

	return PointerFromQGeoPositionInfoSource(nil)
}

func (ptr *QGeoPositionInfoSourceFactory) ConnectPositionInfoSource(f func(parent *core.QObject) *QGeoPositionInfoSource) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSourceFactory::positionInfoSource", f)
	}
}

func (ptr *QGeoPositionInfoSourceFactory) DisconnectPositionInfoSource(parent core.QObject_ITF) {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSourceFactory::positionInfoSource")
	}
}

func (ptr *QGeoPositionInfoSourceFactory) PositionInfoSource(parent core.QObject_ITF) *QGeoPositionInfoSource {
	if ptr.Pointer() != nil {
		var tmpValue = NewQGeoPositionInfoSourceFromPointer(C.QGeoPositionInfoSourceFactory_PositionInfoSource(ptr.Pointer(), core.PointerFromQObject(parent)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQGeoPositionInfoSourceFactory_SatelliteInfoSource
func callbackQGeoPositionInfoSourceFactory_SatelliteInfoSource(ptr unsafe.Pointer, parent unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoPositionInfoSourceFactory::satelliteInfoSource"); signal != nil {
		return PointerFromQGeoSatelliteInfoSource(signal.(func(*core.QObject) *QGeoSatelliteInfoSource)(core.NewQObjectFromPointer(parent)))
	}

	return PointerFromQGeoSatelliteInfoSource(nil)
}

func (ptr *QGeoPositionInfoSourceFactory) ConnectSatelliteInfoSource(f func(parent *core.QObject) *QGeoSatelliteInfoSource) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSourceFactory::satelliteInfoSource", f)
	}
}

func (ptr *QGeoPositionInfoSourceFactory) DisconnectSatelliteInfoSource(parent core.QObject_ITF) {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSourceFactory::satelliteInfoSource")
	}
}

func (ptr *QGeoPositionInfoSourceFactory) SatelliteInfoSource(parent core.QObject_ITF) *QGeoSatelliteInfoSource {
	if ptr.Pointer() != nil {
		var tmpValue = NewQGeoSatelliteInfoSourceFromPointer(C.QGeoPositionInfoSourceFactory_SatelliteInfoSource(ptr.Pointer(), core.PointerFromQObject(parent)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQGeoPositionInfoSourceFactory_DestroyQGeoPositionInfoSourceFactory
func callbackQGeoPositionInfoSourceFactory_DestroyQGeoPositionInfoSourceFactory(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoPositionInfoSourceFactory::~QGeoPositionInfoSourceFactory"); signal != nil {
		signal.(func())()
	} else {
		NewQGeoPositionInfoSourceFactoryFromPointer(ptr).DestroyQGeoPositionInfoSourceFactoryDefault()
	}
}

func (ptr *QGeoPositionInfoSourceFactory) ConnectDestroyQGeoPositionInfoSourceFactory(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSourceFactory::~QGeoPositionInfoSourceFactory", f)
	}
}

func (ptr *QGeoPositionInfoSourceFactory) DisconnectDestroyQGeoPositionInfoSourceFactory() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoPositionInfoSourceFactory::~QGeoPositionInfoSourceFactory")
	}
}

func (ptr *QGeoPositionInfoSourceFactory) DestroyQGeoPositionInfoSourceFactory() {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSourceFactory_DestroyQGeoPositionInfoSourceFactory(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QGeoPositionInfoSourceFactory) DestroyQGeoPositionInfoSourceFactoryDefault() {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSourceFactory_DestroyQGeoPositionInfoSourceFactoryDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
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

func (p *QGeoRectangle) QGeoRectangle_PTR() *QGeoRectangle {
	return p
}

func (p *QGeoRectangle) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QGeoShape_PTR().Pointer()
	}
	return nil
}

func (p *QGeoRectangle) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QGeoShape_PTR().SetPointer(ptr)
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

func (ptr *QGeoRectangle) Center() *QGeoCoordinate {
	if ptr.Pointer() != nil {
		var tmpValue = NewQGeoCoordinateFromPointer(C.QGeoRectangle_Center(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QGeoCoordinate).DestroyQGeoCoordinate)
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

func (ptr *QGeoRectangle) Height() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGeoRectangle_Height(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRectangle) Intersects(rectangle QGeoRectangle_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGeoRectangle_Intersects(ptr.Pointer(), PointerFromQGeoRectangle(rectangle)) != 0
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

func (ptr *QGeoRectangle) ToString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoRectangle_ToString(ptr.Pointer()))
	}
	return ""
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

func (ptr *QGeoRectangle) Translate(degreesLatitude float64, degreesLongitude float64) {
	if ptr.Pointer() != nil {
		C.QGeoRectangle_Translate(ptr.Pointer(), C.double(degreesLatitude), C.double(degreesLongitude))
	}
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

func (ptr *QGeoRectangle) Width() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGeoRectangle_Width(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRectangle) DestroyQGeoRectangle() {
	if ptr.Pointer() != nil {
		C.QGeoRectangle_DestroyQGeoRectangle(ptr.Pointer())
		ptr.SetPointer(nil)
	}
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

type QGeoSatelliteInfo struct {
	ptr unsafe.Pointer
}

type QGeoSatelliteInfo_ITF interface {
	QGeoSatelliteInfo_PTR() *QGeoSatelliteInfo
}

func (p *QGeoSatelliteInfo) QGeoSatelliteInfo_PTR() *QGeoSatelliteInfo {
	return p
}

func (p *QGeoSatelliteInfo) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QGeoSatelliteInfo) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
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

func (ptr *QGeoSatelliteInfo) Attribute(attribute QGeoSatelliteInfo__Attribute) float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGeoSatelliteInfo_Attribute(ptr.Pointer(), C.longlong(attribute)))
	}
	return 0
}

func (ptr *QGeoSatelliteInfo) HasAttribute(attribute QGeoSatelliteInfo__Attribute) bool {
	if ptr.Pointer() != nil {
		return C.QGeoSatelliteInfo_HasAttribute(ptr.Pointer(), C.longlong(attribute)) != 0
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
		C.QGeoSatelliteInfo_DestroyQGeoSatelliteInfo(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//QGeoSatelliteInfoSource::Error
type QGeoSatelliteInfoSource__Error int64

const (
	QGeoSatelliteInfoSource__AccessError        = QGeoSatelliteInfoSource__Error(0)
	QGeoSatelliteInfoSource__ClosedError        = QGeoSatelliteInfoSource__Error(1)
	QGeoSatelliteInfoSource__NoError            = QGeoSatelliteInfoSource__Error(2)
	QGeoSatelliteInfoSource__UnknownSourceError = QGeoSatelliteInfoSource__Error(-1)
)

type QGeoSatelliteInfoSource struct {
	core.QObject
}

type QGeoSatelliteInfoSource_ITF interface {
	core.QObject_ITF
	QGeoSatelliteInfoSource_PTR() *QGeoSatelliteInfoSource
}

func (p *QGeoSatelliteInfoSource) QGeoSatelliteInfoSource_PTR() *QGeoSatelliteInfoSource {
	return p
}

func (p *QGeoSatelliteInfoSource) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QGeoSatelliteInfoSource) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
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

//export callbackQGeoSatelliteInfoSource_SetUpdateInterval
func callbackQGeoSatelliteInfoSource_SetUpdateInterval(ptr unsafe.Pointer, msec C.int) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoSatelliteInfoSource::setUpdateInterval"); signal != nil {
		signal.(func(int))(int(int32(msec)))
	} else {
		NewQGeoSatelliteInfoSourceFromPointer(ptr).SetUpdateIntervalDefault(int(int32(msec)))
	}
}

func (ptr *QGeoSatelliteInfoSource) ConnectSetUpdateInterval(f func(msec int)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoSatelliteInfoSource::setUpdateInterval", f)
	}
}

func (ptr *QGeoSatelliteInfoSource) DisconnectSetUpdateInterval() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoSatelliteInfoSource::setUpdateInterval")
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

func (ptr *QGeoSatelliteInfoSource) UpdateInterval() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QGeoSatelliteInfoSource_UpdateInterval(ptr.Pointer())))
	}
	return 0
}

func NewQGeoSatelliteInfoSource(parent core.QObject_ITF) *QGeoSatelliteInfoSource {
	var tmpValue = NewQGeoSatelliteInfoSourceFromPointer(C.QGeoSatelliteInfoSource_NewQGeoSatelliteInfoSource(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QGeoSatelliteInfoSource_AvailableSources() []string {
	return strings.Split(C.GoString(C.QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_AvailableSources()), "|")
}

func (ptr *QGeoSatelliteInfoSource) AvailableSources() []string {
	return strings.Split(C.GoString(C.QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_AvailableSources()), "|")
}

func QGeoSatelliteInfoSource_CreateDefaultSource(parent core.QObject_ITF) *QGeoSatelliteInfoSource {
	var tmpValue = NewQGeoSatelliteInfoSourceFromPointer(C.QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_CreateDefaultSource(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QGeoSatelliteInfoSource) CreateDefaultSource(parent core.QObject_ITF) *QGeoSatelliteInfoSource {
	var tmpValue = NewQGeoSatelliteInfoSourceFromPointer(C.QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_CreateDefaultSource(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QGeoSatelliteInfoSource_CreateSource(sourceName string, parent core.QObject_ITF) *QGeoSatelliteInfoSource {
	var sourceNameC = C.CString(sourceName)
	defer C.free(unsafe.Pointer(sourceNameC))
	var tmpValue = NewQGeoSatelliteInfoSourceFromPointer(C.QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_CreateSource(sourceNameC, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QGeoSatelliteInfoSource) CreateSource(sourceName string, parent core.QObject_ITF) *QGeoSatelliteInfoSource {
	var sourceNameC = C.CString(sourceName)
	defer C.free(unsafe.Pointer(sourceNameC))
	var tmpValue = NewQGeoSatelliteInfoSourceFromPointer(C.QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_CreateSource(sourceNameC, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQGeoSatelliteInfoSource_Error2
func callbackQGeoSatelliteInfoSource_Error2(ptr unsafe.Pointer, satelliteError C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoSatelliteInfoSource::error2"); signal != nil {
		signal.(func(QGeoSatelliteInfoSource__Error))(QGeoSatelliteInfoSource__Error(satelliteError))
	}

}

func (ptr *QGeoSatelliteInfoSource) ConnectError2(f func(satelliteError QGeoSatelliteInfoSource__Error)) {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoSatelliteInfoSource::error2", f)
	}
}

func (ptr *QGeoSatelliteInfoSource) DisconnectError2() {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoSatelliteInfoSource::error2")
	}
}

func (ptr *QGeoSatelliteInfoSource) Error2(satelliteError QGeoSatelliteInfoSource__Error) {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_Error2(ptr.Pointer(), C.longlong(satelliteError))
	}
}

//export callbackQGeoSatelliteInfoSource_Error
func callbackQGeoSatelliteInfoSource_Error(ptr unsafe.Pointer) C.longlong {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoSatelliteInfoSource::error"); signal != nil {
		return C.longlong(signal.(func() QGeoSatelliteInfoSource__Error)())
	}

	return C.longlong(0)
}

func (ptr *QGeoSatelliteInfoSource) ConnectError(f func() QGeoSatelliteInfoSource__Error) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoSatelliteInfoSource::error", f)
	}
}

func (ptr *QGeoSatelliteInfoSource) DisconnectError() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoSatelliteInfoSource::error")
	}
}

func (ptr *QGeoSatelliteInfoSource) Error() QGeoSatelliteInfoSource__Error {
	if ptr.Pointer() != nil {
		return QGeoSatelliteInfoSource__Error(C.QGeoSatelliteInfoSource_Error(ptr.Pointer()))
	}
	return 0
}

//export callbackQGeoSatelliteInfoSource_MinimumUpdateInterval
func callbackQGeoSatelliteInfoSource_MinimumUpdateInterval(ptr unsafe.Pointer) C.int {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoSatelliteInfoSource::minimumUpdateInterval"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(0))
}

func (ptr *QGeoSatelliteInfoSource) ConnectMinimumUpdateInterval(f func() int) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoSatelliteInfoSource::minimumUpdateInterval", f)
	}
}

func (ptr *QGeoSatelliteInfoSource) DisconnectMinimumUpdateInterval() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoSatelliteInfoSource::minimumUpdateInterval")
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

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoSatelliteInfoSource::requestTimeout"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QGeoSatelliteInfoSource) ConnectRequestTimeout(f func()) {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_ConnectRequestTimeout(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoSatelliteInfoSource::requestTimeout", f)
	}
}

func (ptr *QGeoSatelliteInfoSource) DisconnectRequestTimeout() {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_DisconnectRequestTimeout(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoSatelliteInfoSource::requestTimeout")
	}
}

func (ptr *QGeoSatelliteInfoSource) RequestTimeout() {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_RequestTimeout(ptr.Pointer())
	}
}

//export callbackQGeoSatelliteInfoSource_RequestUpdate
func callbackQGeoSatelliteInfoSource_RequestUpdate(ptr unsafe.Pointer, timeout C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoSatelliteInfoSource::requestUpdate"); signal != nil {
		signal.(func(int))(int(int32(timeout)))
	}

}

func (ptr *QGeoSatelliteInfoSource) ConnectRequestUpdate(f func(timeout int)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoSatelliteInfoSource::requestUpdate", f)
	}
}

func (ptr *QGeoSatelliteInfoSource) DisconnectRequestUpdate(timeout int) {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoSatelliteInfoSource::requestUpdate")
	}
}

func (ptr *QGeoSatelliteInfoSource) RequestUpdate(timeout int) {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_RequestUpdate(ptr.Pointer(), C.int(int32(timeout)))
	}
}

func (ptr *QGeoSatelliteInfoSource) SourceName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoSatelliteInfoSource_SourceName(ptr.Pointer()))
	}
	return ""
}

//export callbackQGeoSatelliteInfoSource_StartUpdates
func callbackQGeoSatelliteInfoSource_StartUpdates(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoSatelliteInfoSource::startUpdates"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QGeoSatelliteInfoSource) ConnectStartUpdates(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoSatelliteInfoSource::startUpdates", f)
	}
}

func (ptr *QGeoSatelliteInfoSource) DisconnectStartUpdates() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoSatelliteInfoSource::startUpdates")
	}
}

func (ptr *QGeoSatelliteInfoSource) StartUpdates() {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_StartUpdates(ptr.Pointer())
	}
}

//export callbackQGeoSatelliteInfoSource_StopUpdates
func callbackQGeoSatelliteInfoSource_StopUpdates(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoSatelliteInfoSource::stopUpdates"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QGeoSatelliteInfoSource) ConnectStopUpdates(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoSatelliteInfoSource::stopUpdates", f)
	}
}

func (ptr *QGeoSatelliteInfoSource) DisconnectStopUpdates() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoSatelliteInfoSource::stopUpdates")
	}
}

func (ptr *QGeoSatelliteInfoSource) StopUpdates() {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_StopUpdates(ptr.Pointer())
	}
}

//export callbackQGeoSatelliteInfoSource_DestroyQGeoSatelliteInfoSource
func callbackQGeoSatelliteInfoSource_DestroyQGeoSatelliteInfoSource(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoSatelliteInfoSource::~QGeoSatelliteInfoSource"); signal != nil {
		signal.(func())()
	} else {
		NewQGeoSatelliteInfoSourceFromPointer(ptr).DestroyQGeoSatelliteInfoSourceDefault()
	}
}

func (ptr *QGeoSatelliteInfoSource) ConnectDestroyQGeoSatelliteInfoSource(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoSatelliteInfoSource::~QGeoSatelliteInfoSource", f)
	}
}

func (ptr *QGeoSatelliteInfoSource) DisconnectDestroyQGeoSatelliteInfoSource() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoSatelliteInfoSource::~QGeoSatelliteInfoSource")
	}
}

func (ptr *QGeoSatelliteInfoSource) DestroyQGeoSatelliteInfoSource() {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_DestroyQGeoSatelliteInfoSource(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QGeoSatelliteInfoSource) DestroyQGeoSatelliteInfoSourceDefault() {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_DestroyQGeoSatelliteInfoSourceDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQGeoSatelliteInfoSource_TimerEvent
func callbackQGeoSatelliteInfoSource_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoSatelliteInfoSource::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGeoSatelliteInfoSourceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGeoSatelliteInfoSource) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoSatelliteInfoSource::timerEvent", f)
	}
}

func (ptr *QGeoSatelliteInfoSource) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoSatelliteInfoSource::timerEvent")
	}
}

func (ptr *QGeoSatelliteInfoSource) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGeoSatelliteInfoSource) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQGeoSatelliteInfoSource_ChildEvent
func callbackQGeoSatelliteInfoSource_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoSatelliteInfoSource::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQGeoSatelliteInfoSourceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QGeoSatelliteInfoSource) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoSatelliteInfoSource::childEvent", f)
	}
}

func (ptr *QGeoSatelliteInfoSource) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoSatelliteInfoSource::childEvent")
	}
}

func (ptr *QGeoSatelliteInfoSource) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGeoSatelliteInfoSource) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQGeoSatelliteInfoSource_ConnectNotify
func callbackQGeoSatelliteInfoSource_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoSatelliteInfoSource::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGeoSatelliteInfoSourceFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGeoSatelliteInfoSource) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoSatelliteInfoSource::connectNotify", f)
	}
}

func (ptr *QGeoSatelliteInfoSource) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoSatelliteInfoSource::connectNotify")
	}
}

func (ptr *QGeoSatelliteInfoSource) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QGeoSatelliteInfoSource) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGeoSatelliteInfoSource_CustomEvent
func callbackQGeoSatelliteInfoSource_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoSatelliteInfoSource::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGeoSatelliteInfoSourceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGeoSatelliteInfoSource) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoSatelliteInfoSource::customEvent", f)
	}
}

func (ptr *QGeoSatelliteInfoSource) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoSatelliteInfoSource::customEvent")
	}
}

func (ptr *QGeoSatelliteInfoSource) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGeoSatelliteInfoSource) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQGeoSatelliteInfoSource_DeleteLater
func callbackQGeoSatelliteInfoSource_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoSatelliteInfoSource::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQGeoSatelliteInfoSourceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QGeoSatelliteInfoSource) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoSatelliteInfoSource::deleteLater", f)
	}
}

func (ptr *QGeoSatelliteInfoSource) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoSatelliteInfoSource::deleteLater")
	}
}

func (ptr *QGeoSatelliteInfoSource) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QGeoSatelliteInfoSource) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQGeoSatelliteInfoSource_DisconnectNotify
func callbackQGeoSatelliteInfoSource_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoSatelliteInfoSource::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGeoSatelliteInfoSourceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGeoSatelliteInfoSource) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoSatelliteInfoSource::disconnectNotify", f)
	}
}

func (ptr *QGeoSatelliteInfoSource) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoSatelliteInfoSource::disconnectNotify")
	}
}

func (ptr *QGeoSatelliteInfoSource) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QGeoSatelliteInfoSource) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfoSource_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGeoSatelliteInfoSource_Event
func callbackQGeoSatelliteInfoSource_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoSatelliteInfoSource::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGeoSatelliteInfoSourceFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QGeoSatelliteInfoSource) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoSatelliteInfoSource::event", f)
	}
}

func (ptr *QGeoSatelliteInfoSource) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoSatelliteInfoSource::event")
	}
}

func (ptr *QGeoSatelliteInfoSource) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGeoSatelliteInfoSource_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QGeoSatelliteInfoSource) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGeoSatelliteInfoSource_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQGeoSatelliteInfoSource_EventFilter
func callbackQGeoSatelliteInfoSource_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoSatelliteInfoSource::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGeoSatelliteInfoSourceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QGeoSatelliteInfoSource) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoSatelliteInfoSource::eventFilter", f)
	}
}

func (ptr *QGeoSatelliteInfoSource) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoSatelliteInfoSource::eventFilter")
	}
}

func (ptr *QGeoSatelliteInfoSource) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGeoSatelliteInfoSource_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QGeoSatelliteInfoSource) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGeoSatelliteInfoSource_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQGeoSatelliteInfoSource_MetaObject
func callbackQGeoSatelliteInfoSource_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGeoSatelliteInfoSource::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQGeoSatelliteInfoSourceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QGeoSatelliteInfoSource) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoSatelliteInfoSource::metaObject", f)
	}
}

func (ptr *QGeoSatelliteInfoSource) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGeoSatelliteInfoSource::metaObject")
	}
}

func (ptr *QGeoSatelliteInfoSource) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QGeoSatelliteInfoSource_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGeoSatelliteInfoSource) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QGeoSatelliteInfoSource_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QGeoShape::ShapeType
type QGeoShape__ShapeType int64

const (
	QGeoShape__UnknownType   = QGeoShape__ShapeType(0)
	QGeoShape__RectangleType = QGeoShape__ShapeType(1)
	QGeoShape__CircleType    = QGeoShape__ShapeType(2)
)

type QGeoShape struct {
	ptr unsafe.Pointer
}

type QGeoShape_ITF interface {
	QGeoShape_PTR() *QGeoShape
}

func (p *QGeoShape) QGeoShape_PTR() *QGeoShape {
	return p
}

func (p *QGeoShape) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QGeoShape) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
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

func (ptr *QGeoShape) Center() *QGeoCoordinate {
	if ptr.Pointer() != nil {
		var tmpValue = NewQGeoCoordinateFromPointer(C.QGeoShape_Center(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QGeoCoordinate).DestroyQGeoCoordinate)
		return tmpValue
	}
	return nil
}

func (ptr *QGeoShape) Contains(coordinate QGeoCoordinate_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGeoShape_Contains(ptr.Pointer(), PointerFromQGeoCoordinate(coordinate)) != 0
	}
	return false
}

func (ptr *QGeoShape) ExtendShape(coordinate QGeoCoordinate_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoShape_ExtendShape(ptr.Pointer(), PointerFromQGeoCoordinate(coordinate))
	}
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

func (ptr *QGeoShape) ToString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoShape_ToString(ptr.Pointer()))
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
		C.QGeoShape_DestroyQGeoShape(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//QNmeaPositionInfoSource::UpdateMode
type QNmeaPositionInfoSource__UpdateMode int64

const (
	QNmeaPositionInfoSource__RealTimeMode   = QNmeaPositionInfoSource__UpdateMode(1)
	QNmeaPositionInfoSource__SimulationMode = QNmeaPositionInfoSource__UpdateMode(2)
)

type QNmeaPositionInfoSource struct {
	QGeoPositionInfoSource
}

type QNmeaPositionInfoSource_ITF interface {
	QGeoPositionInfoSource_ITF
	QNmeaPositionInfoSource_PTR() *QNmeaPositionInfoSource
}

func (p *QNmeaPositionInfoSource) QNmeaPositionInfoSource_PTR() *QNmeaPositionInfoSource {
	return p
}

func (p *QNmeaPositionInfoSource) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QGeoPositionInfoSource_PTR().Pointer()
	}
	return nil
}

func (p *QNmeaPositionInfoSource) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QGeoPositionInfoSource_PTR().SetPointer(ptr)
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
func NewQNmeaPositionInfoSource(updateMode QNmeaPositionInfoSource__UpdateMode, parent core.QObject_ITF) *QNmeaPositionInfoSource {
	var tmpValue = NewQNmeaPositionInfoSourceFromPointer(C.QNmeaPositionInfoSource_NewQNmeaPositionInfoSource(C.longlong(updateMode), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QNmeaPositionInfoSource) Device() *core.QIODevice {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQIODeviceFromPointer(C.QNmeaPositionInfoSource_Device(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQNmeaPositionInfoSource_Error
func callbackQNmeaPositionInfoSource_Error(ptr unsafe.Pointer) C.longlong {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNmeaPositionInfoSource::error"); signal != nil {
		return C.longlong(signal.(func() QGeoPositionInfoSource__Error)())
	}

	return C.longlong(NewQNmeaPositionInfoSourceFromPointer(ptr).ErrorDefault())
}

func (ptr *QNmeaPositionInfoSource) ConnectError(f func() QGeoPositionInfoSource__Error) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNmeaPositionInfoSource::error", f)
	}
}

func (ptr *QNmeaPositionInfoSource) DisconnectError() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNmeaPositionInfoSource::error")
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

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNmeaPositionInfoSource::lastKnownPosition"); signal != nil {
		return PointerFromQGeoPositionInfo(signal.(func(bool) *QGeoPositionInfo)(int8(fromSatellitePositioningMethodsOnly) != 0))
	}

	return PointerFromQGeoPositionInfo(NewQNmeaPositionInfoSourceFromPointer(ptr).LastKnownPositionDefault(int8(fromSatellitePositioningMethodsOnly) != 0))
}

func (ptr *QNmeaPositionInfoSource) ConnectLastKnownPosition(f func(fromSatellitePositioningMethodsOnly bool) *QGeoPositionInfo) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNmeaPositionInfoSource::lastKnownPosition", f)
	}
}

func (ptr *QNmeaPositionInfoSource) DisconnectLastKnownPosition() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNmeaPositionInfoSource::lastKnownPosition")
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

//export callbackQNmeaPositionInfoSource_MinimumUpdateInterval
func callbackQNmeaPositionInfoSource_MinimumUpdateInterval(ptr unsafe.Pointer) C.int {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNmeaPositionInfoSource::minimumUpdateInterval"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(NewQNmeaPositionInfoSourceFromPointer(ptr).MinimumUpdateIntervalDefault()))
}

func (ptr *QNmeaPositionInfoSource) ConnectMinimumUpdateInterval(f func() int) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNmeaPositionInfoSource::minimumUpdateInterval", f)
	}
}

func (ptr *QNmeaPositionInfoSource) DisconnectMinimumUpdateInterval() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNmeaPositionInfoSource::minimumUpdateInterval")
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
func callbackQNmeaPositionInfoSource_ParsePosInfoFromNmeaData(ptr unsafe.Pointer, data *C.char, size C.int, posInfo unsafe.Pointer, hasFix C.char) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNmeaPositionInfoSource::parsePosInfoFromNmeaData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, int, *QGeoPositionInfo, bool) bool)(C.GoString(data), int(int32(size)), NewQGeoPositionInfoFromPointer(posInfo), int8(hasFix) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNmeaPositionInfoSourceFromPointer(ptr).ParsePosInfoFromNmeaDataDefault(C.GoString(data), int(int32(size)), NewQGeoPositionInfoFromPointer(posInfo), int8(hasFix) != 0))))
}

func (ptr *QNmeaPositionInfoSource) ConnectParsePosInfoFromNmeaData(f func(data string, size int, posInfo *QGeoPositionInfo, hasFix bool) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNmeaPositionInfoSource::parsePosInfoFromNmeaData", f)
	}
}

func (ptr *QNmeaPositionInfoSource) DisconnectParsePosInfoFromNmeaData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNmeaPositionInfoSource::parsePosInfoFromNmeaData")
	}
}

func (ptr *QNmeaPositionInfoSource) ParsePosInfoFromNmeaData(data string, size int, posInfo QGeoPositionInfo_ITF, hasFix bool) bool {
	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		return C.QNmeaPositionInfoSource_ParsePosInfoFromNmeaData(ptr.Pointer(), dataC, C.int(int32(size)), PointerFromQGeoPositionInfo(posInfo), C.char(int8(qt.GoBoolToInt(hasFix)))) != 0
	}
	return false
}

func (ptr *QNmeaPositionInfoSource) ParsePosInfoFromNmeaDataDefault(data string, size int, posInfo QGeoPositionInfo_ITF, hasFix bool) bool {
	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		return C.QNmeaPositionInfoSource_ParsePosInfoFromNmeaDataDefault(ptr.Pointer(), dataC, C.int(int32(size)), PointerFromQGeoPositionInfo(posInfo), C.char(int8(qt.GoBoolToInt(hasFix)))) != 0
	}
	return false
}

//export callbackQNmeaPositionInfoSource_RequestUpdate
func callbackQNmeaPositionInfoSource_RequestUpdate(ptr unsafe.Pointer, msec C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNmeaPositionInfoSource::requestUpdate"); signal != nil {
		signal.(func(int))(int(int32(msec)))
	} else {
		NewQNmeaPositionInfoSourceFromPointer(ptr).RequestUpdateDefault(int(int32(msec)))
	}
}

func (ptr *QNmeaPositionInfoSource) ConnectRequestUpdate(f func(msec int)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNmeaPositionInfoSource::requestUpdate", f)
	}
}

func (ptr *QNmeaPositionInfoSource) DisconnectRequestUpdate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNmeaPositionInfoSource::requestUpdate")
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

//export callbackQNmeaPositionInfoSource_SetUpdateInterval
func callbackQNmeaPositionInfoSource_SetUpdateInterval(ptr unsafe.Pointer, msec C.int) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNmeaPositionInfoSource::setUpdateInterval"); signal != nil {
		signal.(func(int))(int(int32(msec)))
	} else {
		NewQNmeaPositionInfoSourceFromPointer(ptr).SetUpdateIntervalDefault(int(int32(msec)))
	}
}

func (ptr *QNmeaPositionInfoSource) ConnectSetUpdateInterval(f func(msec int)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNmeaPositionInfoSource::setUpdateInterval", f)
	}
}

func (ptr *QNmeaPositionInfoSource) DisconnectSetUpdateInterval() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNmeaPositionInfoSource::setUpdateInterval")
	}
}

func (ptr *QNmeaPositionInfoSource) SetUpdateInterval(msec int) {
	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_SetUpdateInterval(ptr.Pointer(), C.int(int32(msec)))
	}
}

func (ptr *QNmeaPositionInfoSource) SetUpdateIntervalDefault(msec int) {
	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_SetUpdateIntervalDefault(ptr.Pointer(), C.int(int32(msec)))
	}
}

func (ptr *QNmeaPositionInfoSource) SetUserEquivalentRangeError(uere float64) {
	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_SetUserEquivalentRangeError(ptr.Pointer(), C.double(uere))
	}
}

//export callbackQNmeaPositionInfoSource_StartUpdates
func callbackQNmeaPositionInfoSource_StartUpdates(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNmeaPositionInfoSource::startUpdates"); signal != nil {
		signal.(func())()
	} else {
		NewQNmeaPositionInfoSourceFromPointer(ptr).StartUpdatesDefault()
	}
}

func (ptr *QNmeaPositionInfoSource) ConnectStartUpdates(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNmeaPositionInfoSource::startUpdates", f)
	}
}

func (ptr *QNmeaPositionInfoSource) DisconnectStartUpdates() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNmeaPositionInfoSource::startUpdates")
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
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNmeaPositionInfoSource::stopUpdates"); signal != nil {
		signal.(func())()
	} else {
		NewQNmeaPositionInfoSourceFromPointer(ptr).StopUpdatesDefault()
	}
}

func (ptr *QNmeaPositionInfoSource) ConnectStopUpdates(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNmeaPositionInfoSource::stopUpdates", f)
	}
}

func (ptr *QNmeaPositionInfoSource) DisconnectStopUpdates() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNmeaPositionInfoSource::stopUpdates")
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

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNmeaPositionInfoSource::supportedPositioningMethods"); signal != nil {
		return C.longlong(signal.(func() QGeoPositionInfoSource__PositioningMethod)())
	}

	return C.longlong(NewQNmeaPositionInfoSourceFromPointer(ptr).SupportedPositioningMethodsDefault())
}

func (ptr *QNmeaPositionInfoSource) ConnectSupportedPositioningMethods(f func() QGeoPositionInfoSource__PositioningMethod) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNmeaPositionInfoSource::supportedPositioningMethods", f)
	}
}

func (ptr *QNmeaPositionInfoSource) DisconnectSupportedPositioningMethods() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNmeaPositionInfoSource::supportedPositioningMethods")
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

func (ptr *QNmeaPositionInfoSource) DestroyQNmeaPositionInfoSource() {
	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_DestroyQNmeaPositionInfoSource(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQNmeaPositionInfoSource_SetPreferredPositioningMethods
func callbackQNmeaPositionInfoSource_SetPreferredPositioningMethods(ptr unsafe.Pointer, methods C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNmeaPositionInfoSource::setPreferredPositioningMethods"); signal != nil {
		signal.(func(QGeoPositionInfoSource__PositioningMethod))(QGeoPositionInfoSource__PositioningMethod(methods))
	} else {
		NewQNmeaPositionInfoSourceFromPointer(ptr).SetPreferredPositioningMethodsDefault(QGeoPositionInfoSource__PositioningMethod(methods))
	}
}

func (ptr *QNmeaPositionInfoSource) ConnectSetPreferredPositioningMethods(f func(methods QGeoPositionInfoSource__PositioningMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNmeaPositionInfoSource::setPreferredPositioningMethods", f)
	}
}

func (ptr *QNmeaPositionInfoSource) DisconnectSetPreferredPositioningMethods() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNmeaPositionInfoSource::setPreferredPositioningMethods")
	}
}

func (ptr *QNmeaPositionInfoSource) SetPreferredPositioningMethods(methods QGeoPositionInfoSource__PositioningMethod) {
	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_SetPreferredPositioningMethods(ptr.Pointer(), C.longlong(methods))
	}
}

func (ptr *QNmeaPositionInfoSource) SetPreferredPositioningMethodsDefault(methods QGeoPositionInfoSource__PositioningMethod) {
	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_SetPreferredPositioningMethodsDefault(ptr.Pointer(), C.longlong(methods))
	}
}

//export callbackQNmeaPositionInfoSource_TimerEvent
func callbackQNmeaPositionInfoSource_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNmeaPositionInfoSource::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQNmeaPositionInfoSourceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QNmeaPositionInfoSource) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNmeaPositionInfoSource::timerEvent", f)
	}
}

func (ptr *QNmeaPositionInfoSource) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNmeaPositionInfoSource::timerEvent")
	}
}

func (ptr *QNmeaPositionInfoSource) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNmeaPositionInfoSource) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQNmeaPositionInfoSource_ChildEvent
func callbackQNmeaPositionInfoSource_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNmeaPositionInfoSource::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQNmeaPositionInfoSourceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QNmeaPositionInfoSource) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNmeaPositionInfoSource::childEvent", f)
	}
}

func (ptr *QNmeaPositionInfoSource) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNmeaPositionInfoSource::childEvent")
	}
}

func (ptr *QNmeaPositionInfoSource) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNmeaPositionInfoSource) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQNmeaPositionInfoSource_ConnectNotify
func callbackQNmeaPositionInfoSource_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNmeaPositionInfoSource::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNmeaPositionInfoSourceFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNmeaPositionInfoSource) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNmeaPositionInfoSource::connectNotify", f)
	}
}

func (ptr *QNmeaPositionInfoSource) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNmeaPositionInfoSource::connectNotify")
	}
}

func (ptr *QNmeaPositionInfoSource) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QNmeaPositionInfoSource) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNmeaPositionInfoSource_CustomEvent
func callbackQNmeaPositionInfoSource_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNmeaPositionInfoSource::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQNmeaPositionInfoSourceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QNmeaPositionInfoSource) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNmeaPositionInfoSource::customEvent", f)
	}
}

func (ptr *QNmeaPositionInfoSource) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNmeaPositionInfoSource::customEvent")
	}
}

func (ptr *QNmeaPositionInfoSource) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QNmeaPositionInfoSource) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQNmeaPositionInfoSource_DeleteLater
func callbackQNmeaPositionInfoSource_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNmeaPositionInfoSource::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQNmeaPositionInfoSourceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QNmeaPositionInfoSource) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNmeaPositionInfoSource::deleteLater", f)
	}
}

func (ptr *QNmeaPositionInfoSource) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNmeaPositionInfoSource::deleteLater")
	}
}

func (ptr *QNmeaPositionInfoSource) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QNmeaPositionInfoSource) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQNmeaPositionInfoSource_DisconnectNotify
func callbackQNmeaPositionInfoSource_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNmeaPositionInfoSource::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNmeaPositionInfoSourceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNmeaPositionInfoSource) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNmeaPositionInfoSource::disconnectNotify", f)
	}
}

func (ptr *QNmeaPositionInfoSource) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNmeaPositionInfoSource::disconnectNotify")
	}
}

func (ptr *QNmeaPositionInfoSource) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QNmeaPositionInfoSource) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QNmeaPositionInfoSource_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNmeaPositionInfoSource_Event
func callbackQNmeaPositionInfoSource_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNmeaPositionInfoSource::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNmeaPositionInfoSourceFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QNmeaPositionInfoSource) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNmeaPositionInfoSource::event", f)
	}
}

func (ptr *QNmeaPositionInfoSource) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNmeaPositionInfoSource::event")
	}
}

func (ptr *QNmeaPositionInfoSource) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNmeaPositionInfoSource_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QNmeaPositionInfoSource) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNmeaPositionInfoSource_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQNmeaPositionInfoSource_EventFilter
func callbackQNmeaPositionInfoSource_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNmeaPositionInfoSource::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNmeaPositionInfoSourceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QNmeaPositionInfoSource) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNmeaPositionInfoSource::eventFilter", f)
	}
}

func (ptr *QNmeaPositionInfoSource) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNmeaPositionInfoSource::eventFilter")
	}
}

func (ptr *QNmeaPositionInfoSource) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNmeaPositionInfoSource_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QNmeaPositionInfoSource) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNmeaPositionInfoSource_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQNmeaPositionInfoSource_MetaObject
func callbackQNmeaPositionInfoSource_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNmeaPositionInfoSource::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQNmeaPositionInfoSourceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QNmeaPositionInfoSource) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNmeaPositionInfoSource::metaObject", f)
	}
}

func (ptr *QNmeaPositionInfoSource) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNmeaPositionInfoSource::metaObject")
	}
}

func (ptr *QNmeaPositionInfoSource) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QNmeaPositionInfoSource_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNmeaPositionInfoSource) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QNmeaPositionInfoSource_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}
