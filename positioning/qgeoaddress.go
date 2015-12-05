package positioning

//#include "positioning.h"
import "C"
import (
	"log"
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoAddress::QGeoAddress")
		}
	}()

	return NewQGeoAddressFromPointer(C.QGeoAddress_NewQGeoAddress())
}

func NewQGeoAddress2(other QGeoAddress_ITF) *QGeoAddress {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoAddress::QGeoAddress")
		}
	}()

	return NewQGeoAddressFromPointer(C.QGeoAddress_NewQGeoAddress2(PointerFromQGeoAddress(other)))
}

func (ptr *QGeoAddress) City() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoAddress::city")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoAddress_City(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoAddress) Clear() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoAddress::clear")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoAddress_Clear(ptr.Pointer())
	}
}

func (ptr *QGeoAddress) Country() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoAddress::country")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoAddress_Country(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoAddress) CountryCode() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoAddress::countryCode")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoAddress_CountryCode(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoAddress) County() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoAddress::county")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoAddress_County(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoAddress) District() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoAddress::district")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoAddress_District(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoAddress) IsEmpty() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoAddress::isEmpty")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QGeoAddress_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGeoAddress) IsTextGenerated() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoAddress::isTextGenerated")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QGeoAddress_IsTextGenerated(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGeoAddress) PostalCode() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoAddress::postalCode")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoAddress_PostalCode(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoAddress) SetCity(city string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoAddress::setCity")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoAddress_SetCity(ptr.Pointer(), C.CString(city))
	}
}

func (ptr *QGeoAddress) SetCountry(country string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoAddress::setCountry")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoAddress_SetCountry(ptr.Pointer(), C.CString(country))
	}
}

func (ptr *QGeoAddress) SetCountryCode(countryCode string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoAddress::setCountryCode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoAddress_SetCountryCode(ptr.Pointer(), C.CString(countryCode))
	}
}

func (ptr *QGeoAddress) SetCounty(county string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoAddress::setCounty")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoAddress_SetCounty(ptr.Pointer(), C.CString(county))
	}
}

func (ptr *QGeoAddress) SetDistrict(district string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoAddress::setDistrict")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoAddress_SetDistrict(ptr.Pointer(), C.CString(district))
	}
}

func (ptr *QGeoAddress) SetPostalCode(postalCode string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoAddress::setPostalCode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoAddress_SetPostalCode(ptr.Pointer(), C.CString(postalCode))
	}
}

func (ptr *QGeoAddress) SetState(state string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoAddress::setState")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoAddress_SetState(ptr.Pointer(), C.CString(state))
	}
}

func (ptr *QGeoAddress) SetStreet(street string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoAddress::setStreet")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoAddress_SetStreet(ptr.Pointer(), C.CString(street))
	}
}

func (ptr *QGeoAddress) SetText(text string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoAddress::setText")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoAddress_SetText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QGeoAddress) Street() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoAddress::street")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoAddress_Street(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoAddress) Text() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoAddress::text")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoAddress_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoAddress) DestroyQGeoAddress() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoAddress::~QGeoAddress")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoAddress_DestroyQGeoAddress(ptr.Pointer())
	}
}
