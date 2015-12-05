package core

//#include "core.h"
import "C"
import (
	"log"
	"unsafe"
)

type QTimeZone struct {
	ptr unsafe.Pointer
}

type QTimeZone_ITF interface {
	QTimeZone_PTR() *QTimeZone
}

func (p *QTimeZone) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTimeZone) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTimeZone(ptr QTimeZone_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTimeZone_PTR().Pointer()
	}
	return nil
}

func NewQTimeZoneFromPointer(ptr unsafe.Pointer) *QTimeZone {
	var n = new(QTimeZone)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTimeZone) QTimeZone_PTR() *QTimeZone {
	return ptr
}

//QTimeZone::NameType
type QTimeZone__NameType int64

const (
	QTimeZone__DefaultName = QTimeZone__NameType(0)
	QTimeZone__LongName    = QTimeZone__NameType(1)
	QTimeZone__ShortName   = QTimeZone__NameType(2)
	QTimeZone__OffsetName  = QTimeZone__NameType(3)
)

//QTimeZone::TimeType
type QTimeZone__TimeType int64

const (
	QTimeZone__StandardTime = QTimeZone__TimeType(0)
	QTimeZone__DaylightTime = QTimeZone__TimeType(1)
	QTimeZone__GenericTime  = QTimeZone__TimeType(2)
)

func NewQTimeZone() *QTimeZone {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTimeZone::QTimeZone")
		}
	}()

	return NewQTimeZoneFromPointer(C.QTimeZone_NewQTimeZone())
}

func NewQTimeZone2(ianaId QByteArray_ITF) *QTimeZone {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTimeZone::QTimeZone")
		}
	}()

	return NewQTimeZoneFromPointer(C.QTimeZone_NewQTimeZone2(PointerFromQByteArray(ianaId)))
}

func NewQTimeZone4(ianaId QByteArray_ITF, offsetSeconds int, name string, abbreviation string, country QLocale__Country, comment string) *QTimeZone {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTimeZone::QTimeZone")
		}
	}()

	return NewQTimeZoneFromPointer(C.QTimeZone_NewQTimeZone4(PointerFromQByteArray(ianaId), C.int(offsetSeconds), C.CString(name), C.CString(abbreviation), C.int(country), C.CString(comment)))
}

func NewQTimeZone5(other QTimeZone_ITF) *QTimeZone {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTimeZone::QTimeZone")
		}
	}()

	return NewQTimeZoneFromPointer(C.QTimeZone_NewQTimeZone5(PointerFromQTimeZone(other)))
}

func NewQTimeZone3(offsetSeconds int) *QTimeZone {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTimeZone::QTimeZone")
		}
	}()

	return NewQTimeZoneFromPointer(C.QTimeZone_NewQTimeZone3(C.int(offsetSeconds)))
}

func (ptr *QTimeZone) Abbreviation(atDateTime QDateTime_ITF) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTimeZone::abbreviation")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QTimeZone_Abbreviation(ptr.Pointer(), PointerFromQDateTime(atDateTime)))
	}
	return ""
}

func (ptr *QTimeZone) Comment() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTimeZone::comment")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QTimeZone_Comment(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTimeZone) Country() QLocale__Country {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTimeZone::country")
		}
	}()

	if ptr.Pointer() != nil {
		return QLocale__Country(C.QTimeZone_Country(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTimeZone) DaylightTimeOffset(atDateTime QDateTime_ITF) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTimeZone::daylightTimeOffset")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTimeZone_DaylightTimeOffset(ptr.Pointer(), PointerFromQDateTime(atDateTime)))
	}
	return 0
}

func (ptr *QTimeZone) DisplayName2(timeType QTimeZone__TimeType, nameType QTimeZone__NameType, locale QLocale_ITF) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTimeZone::displayName")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QTimeZone_DisplayName2(ptr.Pointer(), C.int(timeType), C.int(nameType), PointerFromQLocale(locale)))
	}
	return ""
}

func (ptr *QTimeZone) DisplayName(atDateTime QDateTime_ITF, nameType QTimeZone__NameType, locale QLocale_ITF) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTimeZone::displayName")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QTimeZone_DisplayName(ptr.Pointer(), PointerFromQDateTime(atDateTime), C.int(nameType), PointerFromQLocale(locale)))
	}
	return ""
}

func (ptr *QTimeZone) HasDaylightTime() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTimeZone::hasDaylightTime")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTimeZone_HasDaylightTime(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTimeZone) HasTransitions() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTimeZone::hasTransitions")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTimeZone_HasTransitions(ptr.Pointer()) != 0
	}
	return false
}

func QTimeZone_IanaIdToWindowsId(ianaId QByteArray_ITF) *QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTimeZone::ianaIdToWindowsId")
		}
	}()

	return NewQByteArrayFromPointer(C.QTimeZone_QTimeZone_IanaIdToWindowsId(PointerFromQByteArray(ianaId)))
}

func (ptr *QTimeZone) Id() *QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTimeZone::id")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QTimeZone_Id(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTimeZone) IsDaylightTime(atDateTime QDateTime_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTimeZone::isDaylightTime")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTimeZone_IsDaylightTime(ptr.Pointer(), PointerFromQDateTime(atDateTime)) != 0
	}
	return false
}

func QTimeZone_IsTimeZoneIdAvailable(ianaId QByteArray_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTimeZone::isTimeZoneIdAvailable")
		}
	}()

	return C.QTimeZone_QTimeZone_IsTimeZoneIdAvailable(PointerFromQByteArray(ianaId)) != 0
}

func (ptr *QTimeZone) IsValid() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTimeZone::isValid")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTimeZone_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTimeZone) OffsetFromUtc(atDateTime QDateTime_ITF) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTimeZone::offsetFromUtc")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTimeZone_OffsetFromUtc(ptr.Pointer(), PointerFromQDateTime(atDateTime)))
	}
	return 0
}

func (ptr *QTimeZone) StandardTimeOffset(atDateTime QDateTime_ITF) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTimeZone::standardTimeOffset")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTimeZone_StandardTimeOffset(ptr.Pointer(), PointerFromQDateTime(atDateTime)))
	}
	return 0
}

func (ptr *QTimeZone) Swap(other QTimeZone_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTimeZone::swap")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTimeZone_Swap(ptr.Pointer(), PointerFromQTimeZone(other))
	}
}

func QTimeZone_SystemTimeZone() *QTimeZone {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTimeZone::systemTimeZone")
		}
	}()

	return NewQTimeZoneFromPointer(C.QTimeZone_QTimeZone_SystemTimeZone())
}

func QTimeZone_SystemTimeZoneId() *QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTimeZone::systemTimeZoneId")
		}
	}()

	return NewQByteArrayFromPointer(C.QTimeZone_QTimeZone_SystemTimeZoneId())
}

func QTimeZone_Utc() *QTimeZone {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTimeZone::utc")
		}
	}()

	return NewQTimeZoneFromPointer(C.QTimeZone_QTimeZone_Utc())
}

func QTimeZone_WindowsIdToDefaultIanaId(windowsId QByteArray_ITF) *QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTimeZone::windowsIdToDefaultIanaId")
		}
	}()

	return NewQByteArrayFromPointer(C.QTimeZone_QTimeZone_WindowsIdToDefaultIanaId(PointerFromQByteArray(windowsId)))
}

func QTimeZone_WindowsIdToDefaultIanaId2(windowsId QByteArray_ITF, country QLocale__Country) *QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTimeZone::windowsIdToDefaultIanaId")
		}
	}()

	return NewQByteArrayFromPointer(C.QTimeZone_QTimeZone_WindowsIdToDefaultIanaId2(PointerFromQByteArray(windowsId), C.int(country)))
}

func (ptr *QTimeZone) DestroyQTimeZone() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTimeZone::~QTimeZone")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTimeZone_DestroyQTimeZone(ptr.Pointer())
	}
}
