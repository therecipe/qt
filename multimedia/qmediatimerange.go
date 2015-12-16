package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QMediaTimeRange struct {
	ptr unsafe.Pointer
}

type QMediaTimeRange_ITF interface {
	QMediaTimeRange_PTR() *QMediaTimeRange
}

func (p *QMediaTimeRange) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMediaTimeRange) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMediaTimeRange(ptr QMediaTimeRange_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaTimeRange_PTR().Pointer()
	}
	return nil
}

func NewQMediaTimeRangeFromPointer(ptr unsafe.Pointer) *QMediaTimeRange {
	var n = new(QMediaTimeRange)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMediaTimeRange) QMediaTimeRange_PTR() *QMediaTimeRange {
	return ptr
}

func NewQMediaTimeRange() *QMediaTimeRange {
	defer qt.Recovering("QMediaTimeRange::QMediaTimeRange")

	return NewQMediaTimeRangeFromPointer(C.QMediaTimeRange_NewQMediaTimeRange())
}

func NewQMediaTimeRange3(interval QMediaTimeInterval_ITF) *QMediaTimeRange {
	defer qt.Recovering("QMediaTimeRange::QMediaTimeRange")

	return NewQMediaTimeRangeFromPointer(C.QMediaTimeRange_NewQMediaTimeRange3(PointerFromQMediaTimeInterval(interval)))
}

func NewQMediaTimeRange4(ran QMediaTimeRange_ITF) *QMediaTimeRange {
	defer qt.Recovering("QMediaTimeRange::QMediaTimeRange")

	return NewQMediaTimeRangeFromPointer(C.QMediaTimeRange_NewQMediaTimeRange4(PointerFromQMediaTimeRange(ran)))
}

func NewQMediaTimeRange2(start int64, end int64) *QMediaTimeRange {
	defer qt.Recovering("QMediaTimeRange::QMediaTimeRange")

	return NewQMediaTimeRangeFromPointer(C.QMediaTimeRange_NewQMediaTimeRange2(C.longlong(start), C.longlong(end)))
}

func (ptr *QMediaTimeRange) AddInterval(interval QMediaTimeInterval_ITF) {
	defer qt.Recovering("QMediaTimeRange::addInterval")

	if ptr.Pointer() != nil {
		C.QMediaTimeRange_AddInterval(ptr.Pointer(), PointerFromQMediaTimeInterval(interval))
	}
}

func (ptr *QMediaTimeRange) AddInterval2(start int64, end int64) {
	defer qt.Recovering("QMediaTimeRange::addInterval")

	if ptr.Pointer() != nil {
		C.QMediaTimeRange_AddInterval2(ptr.Pointer(), C.longlong(start), C.longlong(end))
	}
}

func (ptr *QMediaTimeRange) AddTimeRange(ran QMediaTimeRange_ITF) {
	defer qt.Recovering("QMediaTimeRange::addTimeRange")

	if ptr.Pointer() != nil {
		C.QMediaTimeRange_AddTimeRange(ptr.Pointer(), PointerFromQMediaTimeRange(ran))
	}
}

func (ptr *QMediaTimeRange) Clear() {
	defer qt.Recovering("QMediaTimeRange::clear")

	if ptr.Pointer() != nil {
		C.QMediaTimeRange_Clear(ptr.Pointer())
	}
}

func (ptr *QMediaTimeRange) Contains(time int64) bool {
	defer qt.Recovering("QMediaTimeRange::contains")

	if ptr.Pointer() != nil {
		return C.QMediaTimeRange_Contains(ptr.Pointer(), C.longlong(time)) != 0
	}
	return false
}

func (ptr *QMediaTimeRange) EarliestTime() int64 {
	defer qt.Recovering("QMediaTimeRange::earliestTime")

	if ptr.Pointer() != nil {
		return int64(C.QMediaTimeRange_EarliestTime(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaTimeRange) IsContinuous() bool {
	defer qt.Recovering("QMediaTimeRange::isContinuous")

	if ptr.Pointer() != nil {
		return C.QMediaTimeRange_IsContinuous(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaTimeRange) IsEmpty() bool {
	defer qt.Recovering("QMediaTimeRange::isEmpty")

	if ptr.Pointer() != nil {
		return C.QMediaTimeRange_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaTimeRange) LatestTime() int64 {
	defer qt.Recovering("QMediaTimeRange::latestTime")

	if ptr.Pointer() != nil {
		return int64(C.QMediaTimeRange_LatestTime(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaTimeRange) RemoveInterval(interval QMediaTimeInterval_ITF) {
	defer qt.Recovering("QMediaTimeRange::removeInterval")

	if ptr.Pointer() != nil {
		C.QMediaTimeRange_RemoveInterval(ptr.Pointer(), PointerFromQMediaTimeInterval(interval))
	}
}

func (ptr *QMediaTimeRange) RemoveInterval2(start int64, end int64) {
	defer qt.Recovering("QMediaTimeRange::removeInterval")

	if ptr.Pointer() != nil {
		C.QMediaTimeRange_RemoveInterval2(ptr.Pointer(), C.longlong(start), C.longlong(end))
	}
}

func (ptr *QMediaTimeRange) RemoveTimeRange(ran QMediaTimeRange_ITF) {
	defer qt.Recovering("QMediaTimeRange::removeTimeRange")

	if ptr.Pointer() != nil {
		C.QMediaTimeRange_RemoveTimeRange(ptr.Pointer(), PointerFromQMediaTimeRange(ran))
	}
}

func (ptr *QMediaTimeRange) DestroyQMediaTimeRange() {
	defer qt.Recovering("QMediaTimeRange::~QMediaTimeRange")

	if ptr.Pointer() != nil {
		C.QMediaTimeRange_DestroyQMediaTimeRange(ptr.Pointer())
	}
}
