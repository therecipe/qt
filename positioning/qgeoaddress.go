package positioning

//#include "qgeoaddress.h"
import "C"
import (
	"unsafe"
)

type QGeoAddress struct {
	ptr unsafe.Pointer
}

type QGeoAddressITF interface {
	QGeoAddressPTR() *QGeoAddress
}

func (p *QGeoAddress) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QGeoAddress) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQGeoAddress(ptr QGeoAddressITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoAddressPTR().Pointer()
	}
	return nil
}

func QGeoAddressFromPointer(ptr unsafe.Pointer) *QGeoAddress {
	var n = new(QGeoAddress)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGeoAddress) QGeoAddressPTR() *QGeoAddress {
	return ptr
}

func NewQGeoAddress() *QGeoAddress {
	return QGeoAddressFromPointer(unsafe.Pointer(C.QGeoAddress_NewQGeoAddress()))
}

func NewQGeoAddress2(other QGeoAddressITF) *QGeoAddress {
	return QGeoAddressFromPointer(unsafe.Pointer(C.QGeoAddress_NewQGeoAddress2(C.QtObjectPtr(PointerFromQGeoAddress(other)))))
}

func (ptr *QGeoAddress) City() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoAddress_City(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QGeoAddress) Clear() {
	if ptr.Pointer() != nil {
		C.QGeoAddress_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QGeoAddress) Country() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoAddress_Country(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QGeoAddress) CountryCode() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoAddress_CountryCode(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QGeoAddress) County() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoAddress_County(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QGeoAddress) District() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoAddress_District(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QGeoAddress) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QGeoAddress_IsEmpty(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGeoAddress) IsTextGenerated() bool {
	if ptr.Pointer() != nil {
		return C.QGeoAddress_IsTextGenerated(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGeoAddress) PostalCode() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoAddress_PostalCode(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QGeoAddress) SetCity(city string) {
	if ptr.Pointer() != nil {
		C.QGeoAddress_SetCity(C.QtObjectPtr(ptr.Pointer()), C.CString(city))
	}
}

func (ptr *QGeoAddress) SetCountry(country string) {
	if ptr.Pointer() != nil {
		C.QGeoAddress_SetCountry(C.QtObjectPtr(ptr.Pointer()), C.CString(country))
	}
}

func (ptr *QGeoAddress) SetCountryCode(countryCode string) {
	if ptr.Pointer() != nil {
		C.QGeoAddress_SetCountryCode(C.QtObjectPtr(ptr.Pointer()), C.CString(countryCode))
	}
}

func (ptr *QGeoAddress) SetCounty(county string) {
	if ptr.Pointer() != nil {
		C.QGeoAddress_SetCounty(C.QtObjectPtr(ptr.Pointer()), C.CString(county))
	}
}

func (ptr *QGeoAddress) SetDistrict(district string) {
	if ptr.Pointer() != nil {
		C.QGeoAddress_SetDistrict(C.QtObjectPtr(ptr.Pointer()), C.CString(district))
	}
}

func (ptr *QGeoAddress) SetPostalCode(postalCode string) {
	if ptr.Pointer() != nil {
		C.QGeoAddress_SetPostalCode(C.QtObjectPtr(ptr.Pointer()), C.CString(postalCode))
	}
}

func (ptr *QGeoAddress) SetState(state string) {
	if ptr.Pointer() != nil {
		C.QGeoAddress_SetState(C.QtObjectPtr(ptr.Pointer()), C.CString(state))
	}
}

func (ptr *QGeoAddress) SetStreet(street string) {
	if ptr.Pointer() != nil {
		C.QGeoAddress_SetStreet(C.QtObjectPtr(ptr.Pointer()), C.CString(street))
	}
}

func (ptr *QGeoAddress) SetText(text string) {
	if ptr.Pointer() != nil {
		C.QGeoAddress_SetText(C.QtObjectPtr(ptr.Pointer()), C.CString(text))
	}
}

func (ptr *QGeoAddress) Street() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoAddress_Street(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QGeoAddress) Text() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoAddress_Text(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QGeoAddress) DestroyQGeoAddress() {
	if ptr.Pointer() != nil {
		C.QGeoAddress_DestroyQGeoAddress(C.QtObjectPtr(ptr.Pointer()))
	}
}
