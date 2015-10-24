package core

//#include "qtime.h"
import "C"
import (
	"unsafe"
)

type QTime struct {
	ptr unsafe.Pointer
}

type QTimeITF interface {
	QTimePTR() *QTime
}

func (p *QTime) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTime) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTime(ptr QTimeITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTimePTR().Pointer()
	}
	return nil
}

func QTimeFromPointer(ptr unsafe.Pointer) *QTime {
	var n = new(QTime)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTime) QTimePTR() *QTime {
	return ptr
}

func NewQTime() *QTime {
	return QTimeFromPointer(unsafe.Pointer(C.QTime_NewQTime()))
}

func NewQTime3(h int, m int, s int, ms int) *QTime {
	return QTimeFromPointer(unsafe.Pointer(C.QTime_NewQTime3(C.int(h), C.int(m), C.int(s), C.int(ms))))
}

func (ptr *QTime) Elapsed() int {
	if ptr.Pointer() != nil {
		return int(C.QTime_Elapsed(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTime) Hour() int {
	if ptr.Pointer() != nil {
		return int(C.QTime_Hour(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTime) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QTime_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func QTime_IsValid2(h int, m int, s int, ms int) bool {
	return C.QTime_QTime_IsValid2(C.int(h), C.int(m), C.int(s), C.int(ms)) != 0
}

func (ptr *QTime) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QTime_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTime) Minute() int {
	if ptr.Pointer() != nil {
		return int(C.QTime_Minute(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTime) Msec() int {
	if ptr.Pointer() != nil {
		return int(C.QTime_Msec(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTime) MsecsSinceStartOfDay() int {
	if ptr.Pointer() != nil {
		return int(C.QTime_MsecsSinceStartOfDay(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTime) MsecsTo(t QTimeITF) int {
	if ptr.Pointer() != nil {
		return int(C.QTime_MsecsTo(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTime(t))))
	}
	return 0
}

func (ptr *QTime) Restart() int {
	if ptr.Pointer() != nil {
		return int(C.QTime_Restart(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTime) Second() int {
	if ptr.Pointer() != nil {
		return int(C.QTime_Second(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTime) SecsTo(t QTimeITF) int {
	if ptr.Pointer() != nil {
		return int(C.QTime_SecsTo(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTime(t))))
	}
	return 0
}

func (ptr *QTime) SetHMS(h int, m int, s int, ms int) bool {
	if ptr.Pointer() != nil {
		return C.QTime_SetHMS(C.QtObjectPtr(ptr.Pointer()), C.int(h), C.int(m), C.int(s), C.int(ms)) != 0
	}
	return false
}

func (ptr *QTime) Start() {
	if ptr.Pointer() != nil {
		C.QTime_Start(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTime) ToString2(format Qt__DateFormat) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTime_ToString2(C.QtObjectPtr(ptr.Pointer()), C.int(format)))
	}
	return ""
}

func (ptr *QTime) ToString(format string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTime_ToString(C.QtObjectPtr(ptr.Pointer()), C.CString(format)))
	}
	return ""
}
