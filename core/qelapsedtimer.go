package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QElapsedTimer struct {
	ptr unsafe.Pointer
}

type QElapsedTimer_ITF interface {
	QElapsedTimer_PTR() *QElapsedTimer
}

func (p *QElapsedTimer) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QElapsedTimer) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQElapsedTimer(ptr QElapsedTimer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QElapsedTimer_PTR().Pointer()
	}
	return nil
}

func NewQElapsedTimerFromPointer(ptr unsafe.Pointer) *QElapsedTimer {
	var n = new(QElapsedTimer)
	n.SetPointer(ptr)
	return n
}

func (ptr *QElapsedTimer) QElapsedTimer_PTR() *QElapsedTimer {
	return ptr
}

//QElapsedTimer::ClockType
type QElapsedTimer__ClockType int64

const (
	QElapsedTimer__SystemTime         = QElapsedTimer__ClockType(0)
	QElapsedTimer__MonotonicClock     = QElapsedTimer__ClockType(1)
	QElapsedTimer__TickCounter        = QElapsedTimer__ClockType(2)
	QElapsedTimer__MachAbsoluteTime   = QElapsedTimer__ClockType(3)
	QElapsedTimer__PerformanceCounter = QElapsedTimer__ClockType(4)
)

func NewQElapsedTimer() *QElapsedTimer {
	defer qt.Recovering("QElapsedTimer::QElapsedTimer")

	return NewQElapsedTimerFromPointer(C.QElapsedTimer_NewQElapsedTimer())
}

func (ptr *QElapsedTimer) HasExpired(timeout int64) bool {
	defer qt.Recovering("QElapsedTimer::hasExpired")

	if ptr.Pointer() != nil {
		return C.QElapsedTimer_HasExpired(ptr.Pointer(), C.longlong(timeout)) != 0
	}
	return false
}

func (ptr *QElapsedTimer) Invalidate() {
	defer qt.Recovering("QElapsedTimer::invalidate")

	if ptr.Pointer() != nil {
		C.QElapsedTimer_Invalidate(ptr.Pointer())
	}
}

func (ptr *QElapsedTimer) IsValid() bool {
	defer qt.Recovering("QElapsedTimer::isValid")

	if ptr.Pointer() != nil {
		return C.QElapsedTimer_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func QElapsedTimer_ClockType() QElapsedTimer__ClockType {
	defer qt.Recovering("QElapsedTimer::clockType")

	return QElapsedTimer__ClockType(C.QElapsedTimer_QElapsedTimer_ClockType())
}

func (ptr *QElapsedTimer) Elapsed() int64 {
	defer qt.Recovering("QElapsedTimer::elapsed")

	if ptr.Pointer() != nil {
		return int64(C.QElapsedTimer_Elapsed(ptr.Pointer()))
	}
	return 0
}

func QElapsedTimer_IsMonotonic() bool {
	defer qt.Recovering("QElapsedTimer::isMonotonic")

	return C.QElapsedTimer_QElapsedTimer_IsMonotonic() != 0
}

func (ptr *QElapsedTimer) MsecsSinceReference() int64 {
	defer qt.Recovering("QElapsedTimer::msecsSinceReference")

	if ptr.Pointer() != nil {
		return int64(C.QElapsedTimer_MsecsSinceReference(ptr.Pointer()))
	}
	return 0
}

func (ptr *QElapsedTimer) MsecsTo(other QElapsedTimer_ITF) int64 {
	defer qt.Recovering("QElapsedTimer::msecsTo")

	if ptr.Pointer() != nil {
		return int64(C.QElapsedTimer_MsecsTo(ptr.Pointer(), PointerFromQElapsedTimer(other)))
	}
	return 0
}

func (ptr *QElapsedTimer) NsecsElapsed() int64 {
	defer qt.Recovering("QElapsedTimer::nsecsElapsed")

	if ptr.Pointer() != nil {
		return int64(C.QElapsedTimer_NsecsElapsed(ptr.Pointer()))
	}
	return 0
}

func (ptr *QElapsedTimer) Restart() int64 {
	defer qt.Recovering("QElapsedTimer::restart")

	if ptr.Pointer() != nil {
		return int64(C.QElapsedTimer_Restart(ptr.Pointer()))
	}
	return 0
}

func (ptr *QElapsedTimer) SecsTo(other QElapsedTimer_ITF) int64 {
	defer qt.Recovering("QElapsedTimer::secsTo")

	if ptr.Pointer() != nil {
		return int64(C.QElapsedTimer_SecsTo(ptr.Pointer(), PointerFromQElapsedTimer(other)))
	}
	return 0
}

func (ptr *QElapsedTimer) Start() {
	defer qt.Recovering("QElapsedTimer::start")

	if ptr.Pointer() != nil {
		C.QElapsedTimer_Start(ptr.Pointer())
	}
}
