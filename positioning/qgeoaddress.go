package positioning

//#include "positioning.h"
import "C"
import (
	"github.com/therecipe/qt"
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
	defer qt.Recovering("QGeoAddress::QGeoAddress")

	return NewQGeoAddressFromPointer(C.QGeoAddress_NewQGeoAddress())
}

func NewQGeoAddress2(other QGeoAddress_ITF) *QGeoAddress {
	defer qt.Recovering("QGeoAddress::QGeoAddress")

	return NewQGeoAddressFromPointer(C.QGeoAddress_NewQGeoAddress2(PointerFromQGeoAddress(other)))
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
