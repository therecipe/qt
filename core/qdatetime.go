package core

//#include "qdatetime.h"
import "C"
import (
	"unsafe"
)

type QDateTime struct {
	ptr unsafe.Pointer
}

type QDateTimeITF interface {
	QDateTimePTR() *QDateTime
}

func (p *QDateTime) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDateTime) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDateTime(ptr QDateTimeITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDateTimePTR().Pointer()
	}
	return nil
}

func QDateTimeFromPointer(ptr unsafe.Pointer) *QDateTime {
	var n = new(QDateTime)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDateTime) QDateTimePTR() *QDateTime {
	return ptr
}

func (ptr *QDateTime) ToString2(format Qt__DateFormat) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDateTime_ToString2(C.QtObjectPtr(ptr.Pointer()), C.int(format)))
	}
	return ""
}

func NewQDateTime() *QDateTime {
	return QDateTimeFromPointer(unsafe.Pointer(C.QDateTime_NewQDateTime()))
}

func NewQDateTime2(date QDateITF) *QDateTime {
	return QDateTimeFromPointer(unsafe.Pointer(C.QDateTime_NewQDateTime2(C.QtObjectPtr(PointerFromQDate(date)))))
}

func NewQDateTime3(date QDateITF, time QTimeITF, spec Qt__TimeSpec) *QDateTime {
	return QDateTimeFromPointer(unsafe.Pointer(C.QDateTime_NewQDateTime3(C.QtObjectPtr(PointerFromQDate(date)), C.QtObjectPtr(PointerFromQTime(time)), C.int(spec))))
}

func NewQDateTime4(date QDateITF, time QTimeITF, spec Qt__TimeSpec, offsetSeconds int) *QDateTime {
	return QDateTimeFromPointer(unsafe.Pointer(C.QDateTime_NewQDateTime4(C.QtObjectPtr(PointerFromQDate(date)), C.QtObjectPtr(PointerFromQTime(time)), C.int(spec), C.int(offsetSeconds))))
}

func NewQDateTime5(date QDateITF, time QTimeITF, timeZone QTimeZoneITF) *QDateTime {
	return QDateTimeFromPointer(unsafe.Pointer(C.QDateTime_NewQDateTime5(C.QtObjectPtr(PointerFromQDate(date)), C.QtObjectPtr(PointerFromQTime(time)), C.QtObjectPtr(PointerFromQTimeZone(timeZone)))))
}

func NewQDateTime6(other QDateTimeITF) *QDateTime {
	return QDateTimeFromPointer(unsafe.Pointer(C.QDateTime_NewQDateTime6(C.QtObjectPtr(PointerFromQDateTime(other)))))
}

func (ptr *QDateTime) IsDaylightTime() bool {
	if ptr.Pointer() != nil {
		return C.QDateTime_IsDaylightTime(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDateTime) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QDateTime_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDateTime) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QDateTime_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDateTime) OffsetFromUtc() int {
	if ptr.Pointer() != nil {
		return int(C.QDateTime_OffsetFromUtc(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDateTime) SetDate(date QDateITF) {
	if ptr.Pointer() != nil {
		C.QDateTime_SetDate(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQDate(date)))
	}
}

func (ptr *QDateTime) SetOffsetFromUtc(offsetSeconds int) {
	if ptr.Pointer() != nil {
		C.QDateTime_SetOffsetFromUtc(C.QtObjectPtr(ptr.Pointer()), C.int(offsetSeconds))
	}
}

func (ptr *QDateTime) SetTime(time QTimeITF) {
	if ptr.Pointer() != nil {
		C.QDateTime_SetTime(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTime(time)))
	}
}

func (ptr *QDateTime) SetTimeSpec(spec Qt__TimeSpec) {
	if ptr.Pointer() != nil {
		C.QDateTime_SetTimeSpec(C.QtObjectPtr(ptr.Pointer()), C.int(spec))
	}
}

func (ptr *QDateTime) SetTimeZone(toZone QTimeZoneITF) {
	if ptr.Pointer() != nil {
		C.QDateTime_SetTimeZone(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTimeZone(toZone)))
	}
}

func (ptr *QDateTime) Swap(other QDateTimeITF) {
	if ptr.Pointer() != nil {
		C.QDateTime_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQDateTime(other)))
	}
}

func (ptr *QDateTime) TimeSpec() Qt__TimeSpec {
	if ptr.Pointer() != nil {
		return Qt__TimeSpec(C.QDateTime_TimeSpec(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDateTime) TimeZoneAbbreviation() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDateTime_TimeZoneAbbreviation(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QDateTime) ToString(format string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDateTime_ToString(C.QtObjectPtr(ptr.Pointer()), C.CString(format)))
	}
	return ""
}

func (ptr *QDateTime) DestroyQDateTime() {
	if ptr.Pointer() != nil {
		C.QDateTime_DestroyQDateTime(C.QtObjectPtr(ptr.Pointer()))
	}
}
