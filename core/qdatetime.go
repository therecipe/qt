package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QDateTime struct {
	ptr unsafe.Pointer
}

type QDateTime_ITF interface {
	QDateTime_PTR() *QDateTime
}

func (p *QDateTime) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDateTime) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDateTime(ptr QDateTime_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDateTime_PTR().Pointer()
	}
	return nil
}

func NewQDateTimeFromPointer(ptr unsafe.Pointer) *QDateTime {
	var n = new(QDateTime)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDateTime) QDateTime_PTR() *QDateTime {
	return ptr
}

func QDateTime_CurrentDateTime() *QDateTime {
	defer qt.Recovering("QDateTime::currentDateTime")

	return NewQDateTimeFromPointer(C.QDateTime_QDateTime_CurrentDateTime())
}

func QDateTime_CurrentDateTimeUtc() *QDateTime {
	defer qt.Recovering("QDateTime::currentDateTimeUtc")

	return NewQDateTimeFromPointer(C.QDateTime_QDateTime_CurrentDateTimeUtc())
}

func QDateTime_FromString(stri string, format Qt__DateFormat) *QDateTime {
	defer qt.Recovering("QDateTime::fromString")

	return NewQDateTimeFromPointer(C.QDateTime_QDateTime_FromString(C.CString(stri), C.int(format)))
}

func QDateTime_FromString2(stri string, format string) *QDateTime {
	defer qt.Recovering("QDateTime::fromString")

	return NewQDateTimeFromPointer(C.QDateTime_QDateTime_FromString2(C.CString(stri), C.CString(format)))
}

func (ptr *QDateTime) ToOffsetFromUtc(offsetSeconds int) *QDateTime {
	defer qt.Recovering("QDateTime::toOffsetFromUtc")

	if ptr.Pointer() != nil {
		return NewQDateTimeFromPointer(C.QDateTime_ToOffsetFromUtc(ptr.Pointer(), C.int(offsetSeconds)))
	}
	return nil
}

func (ptr *QDateTime) ToString2(format Qt__DateFormat) string {
	defer qt.Recovering("QDateTime::toString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDateTime_ToString2(ptr.Pointer(), C.int(format)))
	}
	return ""
}

func (ptr *QDateTime) ToTimeSpec(spec Qt__TimeSpec) *QDateTime {
	defer qt.Recovering("QDateTime::toTimeSpec")

	if ptr.Pointer() != nil {
		return NewQDateTimeFromPointer(C.QDateTime_ToTimeSpec(ptr.Pointer(), C.int(spec)))
	}
	return nil
}

func NewQDateTime() *QDateTime {
	defer qt.Recovering("QDateTime::QDateTime")

	return NewQDateTimeFromPointer(C.QDateTime_NewQDateTime())
}

func NewQDateTime2(date QDate_ITF) *QDateTime {
	defer qt.Recovering("QDateTime::QDateTime")

	return NewQDateTimeFromPointer(C.QDateTime_NewQDateTime2(PointerFromQDate(date)))
}

func NewQDateTime3(date QDate_ITF, time QTime_ITF, spec Qt__TimeSpec) *QDateTime {
	defer qt.Recovering("QDateTime::QDateTime")

	return NewQDateTimeFromPointer(C.QDateTime_NewQDateTime3(PointerFromQDate(date), PointerFromQTime(time), C.int(spec)))
}

func NewQDateTime4(date QDate_ITF, time QTime_ITF, spec Qt__TimeSpec, offsetSeconds int) *QDateTime {
	defer qt.Recovering("QDateTime::QDateTime")

	return NewQDateTimeFromPointer(C.QDateTime_NewQDateTime4(PointerFromQDate(date), PointerFromQTime(time), C.int(spec), C.int(offsetSeconds)))
}

func NewQDateTime5(date QDate_ITF, time QTime_ITF, timeZone QTimeZone_ITF) *QDateTime {
	defer qt.Recovering("QDateTime::QDateTime")

	return NewQDateTimeFromPointer(C.QDateTime_NewQDateTime5(PointerFromQDate(date), PointerFromQTime(time), PointerFromQTimeZone(timeZone)))
}

func NewQDateTime6(other QDateTime_ITF) *QDateTime {
	defer qt.Recovering("QDateTime::QDateTime")

	return NewQDateTimeFromPointer(C.QDateTime_NewQDateTime6(PointerFromQDateTime(other)))
}

func (ptr *QDateTime) AddMonths(nmonths int) *QDateTime {
	defer qt.Recovering("QDateTime::addMonths")

	if ptr.Pointer() != nil {
		return NewQDateTimeFromPointer(C.QDateTime_AddMonths(ptr.Pointer(), C.int(nmonths)))
	}
	return nil
}

func (ptr *QDateTime) AddYears(nyears int) *QDateTime {
	defer qt.Recovering("QDateTime::addYears")

	if ptr.Pointer() != nil {
		return NewQDateTimeFromPointer(C.QDateTime_AddYears(ptr.Pointer(), C.int(nyears)))
	}
	return nil
}

func (ptr *QDateTime) IsDaylightTime() bool {
	defer qt.Recovering("QDateTime::isDaylightTime")

	if ptr.Pointer() != nil {
		return C.QDateTime_IsDaylightTime(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDateTime) IsNull() bool {
	defer qt.Recovering("QDateTime::isNull")

	if ptr.Pointer() != nil {
		return C.QDateTime_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDateTime) IsValid() bool {
	defer qt.Recovering("QDateTime::isValid")

	if ptr.Pointer() != nil {
		return C.QDateTime_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDateTime) OffsetFromUtc() int {
	defer qt.Recovering("QDateTime::offsetFromUtc")

	if ptr.Pointer() != nil {
		return int(C.QDateTime_OffsetFromUtc(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDateTime) SetDate(date QDate_ITF) {
	defer qt.Recovering("QDateTime::setDate")

	if ptr.Pointer() != nil {
		C.QDateTime_SetDate(ptr.Pointer(), PointerFromQDate(date))
	}
}

func (ptr *QDateTime) SetOffsetFromUtc(offsetSeconds int) {
	defer qt.Recovering("QDateTime::setOffsetFromUtc")

	if ptr.Pointer() != nil {
		C.QDateTime_SetOffsetFromUtc(ptr.Pointer(), C.int(offsetSeconds))
	}
}

func (ptr *QDateTime) SetTime(time QTime_ITF) {
	defer qt.Recovering("QDateTime::setTime")

	if ptr.Pointer() != nil {
		C.QDateTime_SetTime(ptr.Pointer(), PointerFromQTime(time))
	}
}

func (ptr *QDateTime) SetTimeSpec(spec Qt__TimeSpec) {
	defer qt.Recovering("QDateTime::setTimeSpec")

	if ptr.Pointer() != nil {
		C.QDateTime_SetTimeSpec(ptr.Pointer(), C.int(spec))
	}
}

func (ptr *QDateTime) SetTimeZone(toZone QTimeZone_ITF) {
	defer qt.Recovering("QDateTime::setTimeZone")

	if ptr.Pointer() != nil {
		C.QDateTime_SetTimeZone(ptr.Pointer(), PointerFromQTimeZone(toZone))
	}
}

func (ptr *QDateTime) Swap(other QDateTime_ITF) {
	defer qt.Recovering("QDateTime::swap")

	if ptr.Pointer() != nil {
		C.QDateTime_Swap(ptr.Pointer(), PointerFromQDateTime(other))
	}
}

func (ptr *QDateTime) TimeSpec() Qt__TimeSpec {
	defer qt.Recovering("QDateTime::timeSpec")

	if ptr.Pointer() != nil {
		return Qt__TimeSpec(C.QDateTime_TimeSpec(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDateTime) TimeZone() *QTimeZone {
	defer qt.Recovering("QDateTime::timeZone")

	if ptr.Pointer() != nil {
		return NewQTimeZoneFromPointer(C.QDateTime_TimeZone(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDateTime) TimeZoneAbbreviation() string {
	defer qt.Recovering("QDateTime::timeZoneAbbreviation")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDateTime_TimeZoneAbbreviation(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDateTime) ToLocalTime() *QDateTime {
	defer qt.Recovering("QDateTime::toLocalTime")

	if ptr.Pointer() != nil {
		return NewQDateTimeFromPointer(C.QDateTime_ToLocalTime(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDateTime) ToString(format string) string {
	defer qt.Recovering("QDateTime::toString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDateTime_ToString(ptr.Pointer(), C.CString(format)))
	}
	return ""
}

func (ptr *QDateTime) ToTimeZone(timeZone QTimeZone_ITF) *QDateTime {
	defer qt.Recovering("QDateTime::toTimeZone")

	if ptr.Pointer() != nil {
		return NewQDateTimeFromPointer(C.QDateTime_ToTimeZone(ptr.Pointer(), PointerFromQTimeZone(timeZone)))
	}
	return nil
}

func (ptr *QDateTime) ToUTC() *QDateTime {
	defer qt.Recovering("QDateTime::toUTC")

	if ptr.Pointer() != nil {
		return NewQDateTimeFromPointer(C.QDateTime_ToUTC(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDateTime) DestroyQDateTime() {
	defer qt.Recovering("QDateTime::~QDateTime")

	if ptr.Pointer() != nil {
		C.QDateTime_DestroyQDateTime(ptr.Pointer())
	}
}
