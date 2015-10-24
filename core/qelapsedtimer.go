package core

//#include "qelapsedtimer.h"
import "C"
import (
	"unsafe"
)

type QElapsedTimer struct {
	ptr unsafe.Pointer
}

type QElapsedTimerITF interface {
	QElapsedTimerPTR() *QElapsedTimer
}

func (p *QElapsedTimer) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QElapsedTimer) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQElapsedTimer(ptr QElapsedTimerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QElapsedTimerPTR().Pointer()
	}
	return nil
}

func QElapsedTimerFromPointer(ptr unsafe.Pointer) *QElapsedTimer {
	var n = new(QElapsedTimer)
	n.SetPointer(ptr)
	return n
}

func (ptr *QElapsedTimer) QElapsedTimerPTR() *QElapsedTimer {
	return ptr
}

//QElapsedTimer::ClockType
type QElapsedTimer__ClockType int

var (
	QElapsedTimer__SystemTime         = QElapsedTimer__ClockType(0)
	QElapsedTimer__MonotonicClock     = QElapsedTimer__ClockType(1)
	QElapsedTimer__TickCounter        = QElapsedTimer__ClockType(2)
	QElapsedTimer__MachAbsoluteTime   = QElapsedTimer__ClockType(3)
	QElapsedTimer__PerformanceCounter = QElapsedTimer__ClockType(4)
)

func NewQElapsedTimer() *QElapsedTimer {
	return QElapsedTimerFromPointer(unsafe.Pointer(C.QElapsedTimer_NewQElapsedTimer()))
}

func (ptr *QElapsedTimer) Invalidate() {
	if ptr.Pointer() != nil {
		C.QElapsedTimer_Invalidate(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QElapsedTimer) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QElapsedTimer_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func QElapsedTimer_ClockType() QElapsedTimer__ClockType {
	return QElapsedTimer__ClockType(C.QElapsedTimer_QElapsedTimer_ClockType())
}

func QElapsedTimer_IsMonotonic() bool {
	return C.QElapsedTimer_QElapsedTimer_IsMonotonic() != 0
}

func (ptr *QElapsedTimer) Start() {
	if ptr.Pointer() != nil {
		C.QElapsedTimer_Start(C.QtObjectPtr(ptr.Pointer()))
	}
}
