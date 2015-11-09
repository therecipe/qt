package positioning

//#include "qgeoaddress.h"
import "C"
import (
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

func (ptr *QGeoAddress) QGeoAddress_PTR() *QGeoAddress {
	return ptr
}

func NewQGeoAddress() *QGeoAddress {
	return NewQGeoAddressFromPointer(C.QGeoAddress_NewQGeoAddress())
}

func NewQGeoAddress2(other QGeoAddress_ITF) *QGeoAddress {
	return NewQGeoAddressFromPointer(C.QGeoAddress_NewQGeoAddress2(PointerFromQGeoAddress(other)))
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
		C.QGeoAddress_SetCity(ptr.Pointer(), C.CString(city))
	}
}

func (ptr *QGeoAddress) SetCountry(country string) {
	if ptr.Pointer() != nil {
		C.QGeoAddress_SetCountry(ptr.Pointer(), C.CString(country))
	}
}

func (ptr *QGeoAddress) SetCountryCode(countryCode string) {
	if ptr.Pointer() != nil {
		C.QGeoAddress_SetCountryCode(ptr.Pointer(), C.CString(countryCode))
	}
}

func (ptr *QGeoAddress) SetCounty(county string) {
	if ptr.Pointer() != nil {
		C.QGeoAddress_SetCounty(ptr.Pointer(), C.CString(county))
	}
}

func (ptr *QGeoAddress) SetDistrict(district string) {
	if ptr.Pointer() != nil {
		C.QGeoAddress_SetDistrict(ptr.Pointer(), C.CString(district))
	}
}

func (ptr *QGeoAddress) SetPostalCode(postalCode string) {
	if ptr.Pointer() != nil {
		C.QGeoAddress_SetPostalCode(ptr.Pointer(), C.CString(postalCode))
	}
}

func (ptr *QGeoAddress) SetState(state string) {
	if ptr.Pointer() != nil {
		C.QGeoAddress_SetState(ptr.Pointer(), C.CString(state))
	}
}

func (ptr *QGeoAddress) SetStreet(street string) {
	if ptr.Pointer() != nil {
		C.QGeoAddress_SetStreet(ptr.Pointer(), C.CString(street))
	}
}

func (ptr *QGeoAddress) SetText(text string) {
	if ptr.Pointer() != nil {
		C.QGeoAddress_SetText(ptr.Pointer(), C.CString(text))
	}
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
	}
}
