package core

//#include "qtimezone.h"
import "C"
import (
	"unsafe"
)

type QTimeZone struct {
	ptr unsafe.Pointer
}

type QTimeZoneITF interface {
	QTimeZonePTR() *QTimeZone
}

func (p *QTimeZone) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTimeZone) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTimeZone(ptr QTimeZoneITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTimeZonePTR().Pointer()
	}
	return nil
}

func QTimeZoneFromPointer(ptr unsafe.Pointer) *QTimeZone {
	var n = new(QTimeZone)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTimeZone) QTimeZonePTR() *QTimeZone {
	return ptr
}

//QTimeZone::NameType
type QTimeZone__NameType int

var (
	QTimeZone__DefaultName = QTimeZone__NameType(0)
	QTimeZone__LongName    = QTimeZone__NameType(1)
	QTimeZone__ShortName   = QTimeZone__NameType(2)
	QTimeZone__OffsetName  = QTimeZone__NameType(3)
)

//QTimeZone::TimeType
type QTimeZone__TimeType int

var (
	QTimeZone__StandardTime = QTimeZone__TimeType(0)
	QTimeZone__DaylightTime = QTimeZone__TimeType(1)
	QTimeZone__GenericTime  = QTimeZone__TimeType(2)
)

func NewQTimeZone() *QTimeZone {
	return QTimeZoneFromPointer(unsafe.Pointer(C.QTimeZone_NewQTimeZone()))
}

func NewQTimeZone2(ianaId QByteArrayITF) *QTimeZone {
	return QTimeZoneFromPointer(unsafe.Pointer(C.QTimeZone_NewQTimeZone2(C.QtObjectPtr(PointerFromQByteArray(ianaId)))))
}

func NewQTimeZone4(ianaId QByteArrayITF, offsetSeconds int, name string, abbreviation string, country QLocale__Country, comment string) *QTimeZone {
	return QTimeZoneFromPointer(unsafe.Pointer(C.QTimeZone_NewQTimeZone4(C.QtObjectPtr(PointerFromQByteArray(ianaId)), C.int(offsetSeconds), C.CString(name), C.CString(abbreviation), C.int(country), C.CString(comment))))
}

func NewQTimeZone5(other QTimeZoneITF) *QTimeZone {
	return QTimeZoneFromPointer(unsafe.Pointer(C.QTimeZone_NewQTimeZone5(C.QtObjectPtr(PointerFromQTimeZone(other)))))
}

func NewQTimeZone3(offsetSeconds int) *QTimeZone {
	return QTimeZoneFromPointer(unsafe.Pointer(C.QTimeZone_NewQTimeZone3(C.int(offsetSeconds))))
}

func (ptr *QTimeZone) Abbreviation(atDateTime QDateTimeITF) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTimeZone_Abbreviation(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQDateTime(atDateTime))))
	}
	return ""
}

func (ptr *QTimeZone) Comment() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTimeZone_Comment(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QTimeZone) Country() QLocale__Country {
	if ptr.Pointer() != nil {
		return QLocale__Country(C.QTimeZone_Country(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTimeZone) DaylightTimeOffset(atDateTime QDateTimeITF) int {
	if ptr.Pointer() != nil {
		return int(C.QTimeZone_DaylightTimeOffset(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQDateTime(atDateTime))))
	}
	return 0
}

func (ptr *QTimeZone) DisplayName2(timeType QTimeZone__TimeType, nameType QTimeZone__NameType, locale QLocaleITF) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTimeZone_DisplayName2(C.QtObjectPtr(ptr.Pointer()), C.int(timeType), C.int(nameType), C.QtObjectPtr(PointerFromQLocale(locale))))
	}
	return ""
}

func (ptr *QTimeZone) DisplayName(atDateTime QDateTimeITF, nameType QTimeZone__NameType, locale QLocaleITF) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTimeZone_DisplayName(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQDateTime(atDateTime)), C.int(nameType), C.QtObjectPtr(PointerFromQLocale(locale))))
	}
	return ""
}

func (ptr *QTimeZone) HasDaylightTime() bool {
	if ptr.Pointer() != nil {
		return C.QTimeZone_HasDaylightTime(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTimeZone) HasTransitions() bool {
	if ptr.Pointer() != nil {
		return C.QTimeZone_HasTransitions(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTimeZone) IsDaylightTime(atDateTime QDateTimeITF) bool {
	if ptr.Pointer() != nil {
		return C.QTimeZone_IsDaylightTime(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQDateTime(atDateTime))) != 0
	}
	return false
}

func QTimeZone_IsTimeZoneIdAvailable(ianaId QByteArrayITF) bool {
	return C.QTimeZone_QTimeZone_IsTimeZoneIdAvailable(C.QtObjectPtr(PointerFromQByteArray(ianaId))) != 0
}

func (ptr *QTimeZone) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QTimeZone_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTimeZone) OffsetFromUtc(atDateTime QDateTimeITF) int {
	if ptr.Pointer() != nil {
		return int(C.QTimeZone_OffsetFromUtc(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQDateTime(atDateTime))))
	}
	return 0
}

func (ptr *QTimeZone) StandardTimeOffset(atDateTime QDateTimeITF) int {
	if ptr.Pointer() != nil {
		return int(C.QTimeZone_StandardTimeOffset(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQDateTime(atDateTime))))
	}
	return 0
}

func (ptr *QTimeZone) Swap(other QTimeZoneITF) {
	if ptr.Pointer() != nil {
		C.QTimeZone_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTimeZone(other)))
	}
}

func (ptr *QTimeZone) DestroyQTimeZone() {
	if ptr.Pointer() != nil {
		C.QTimeZone_DestroyQTimeZone(C.QtObjectPtr(ptr.Pointer()))
	}
}
