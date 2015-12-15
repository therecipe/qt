package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QTime struct {
	ptr unsafe.Pointer
}

type QTime_ITF interface {
	QTime_PTR() *QTime
}

func (p *QTime) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTime) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTime(ptr QTime_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTime_PTR().Pointer()
	}
	return nil
}

func NewQTimeFromPointer(ptr unsafe.Pointer) *QTime {
	var n = new(QTime)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTime) QTime_PTR() *QTime {
	return ptr
}

func NewQTime() *QTime {
	defer qt.Recovering("QTime::QTime")

	return NewQTimeFromPointer(C.QTime_NewQTime())
}

func NewQTime3(h int, m int, s int, ms int) *QTime {
	defer qt.Recovering("QTime::QTime")

	return NewQTimeFromPointer(C.QTime_NewQTime3(C.int(h), C.int(m), C.int(s), C.int(ms)))
}

func (ptr *QTime) Elapsed() int {
	defer qt.Recovering("QTime::elapsed")

	if ptr.Pointer() != nil {
		return int(C.QTime_Elapsed(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTime) Hour() int {
	defer qt.Recovering("QTime::hour")

	if ptr.Pointer() != nil {
		return int(C.QTime_Hour(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTime) IsNull() bool {
	defer qt.Recovering("QTime::isNull")

	if ptr.Pointer() != nil {
		return C.QTime_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func QTime_IsValid2(h int, m int, s int, ms int) bool {
	defer qt.Recovering("QTime::isValid")

	return C.QTime_QTime_IsValid2(C.int(h), C.int(m), C.int(s), C.int(ms)) != 0
}

func (ptr *QTime) IsValid() bool {
	defer qt.Recovering("QTime::isValid")

	if ptr.Pointer() != nil {
		return C.QTime_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTime) Minute() int {
	defer qt.Recovering("QTime::minute")

	if ptr.Pointer() != nil {
		return int(C.QTime_Minute(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTime) Msec() int {
	defer qt.Recovering("QTime::msec")

	if ptr.Pointer() != nil {
		return int(C.QTime_Msec(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTime) MsecsSinceStartOfDay() int {
	defer qt.Recovering("QTime::msecsSinceStartOfDay")

	if ptr.Pointer() != nil {
		return int(C.QTime_MsecsSinceStartOfDay(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTime) MsecsTo(t QTime_ITF) int {
	defer qt.Recovering("QTime::msecsTo")

	if ptr.Pointer() != nil {
		return int(C.QTime_MsecsTo(ptr.Pointer(), PointerFromQTime(t)))
	}
	return 0
}

func (ptr *QTime) Restart() int {
	defer qt.Recovering("QTime::restart")

	if ptr.Pointer() != nil {
		return int(C.QTime_Restart(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTime) Second() int {
	defer qt.Recovering("QTime::second")

	if ptr.Pointer() != nil {
		return int(C.QTime_Second(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTime) SecsTo(t QTime_ITF) int {
	defer qt.Recovering("QTime::secsTo")

	if ptr.Pointer() != nil {
		return int(C.QTime_SecsTo(ptr.Pointer(), PointerFromQTime(t)))
	}
	return 0
}

func (ptr *QTime) SetHMS(h int, m int, s int, ms int) bool {
	defer qt.Recovering("QTime::setHMS")

	if ptr.Pointer() != nil {
		return C.QTime_SetHMS(ptr.Pointer(), C.int(h), C.int(m), C.int(s), C.int(ms)) != 0
	}
	return false
}

func (ptr *QTime) Start() {
	defer qt.Recovering("QTime::start")

	if ptr.Pointer() != nil {
		C.QTime_Start(ptr.Pointer())
	}
}

func (ptr *QTime) ToString2(format Qt__DateFormat) string {
	defer qt.Recovering("QTime::toString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QTime_ToString2(ptr.Pointer(), C.int(format)))
	}
	return ""
}

func (ptr *QTime) ToString(format string) string {
	defer qt.Recovering("QTime::toString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QTime_ToString(ptr.Pointer(), C.CString(format)))
	}
	return ""
}
