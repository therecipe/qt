package multimedia

//#include "qmediatimerange.h"
import "C"
import (
	"unsafe"
)

type QMediaTimeRange struct {
	ptr unsafe.Pointer
}

type QMediaTimeRangeITF interface {
	QMediaTimeRangePTR() *QMediaTimeRange
}

func (p *QMediaTimeRange) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMediaTimeRange) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMediaTimeRange(ptr QMediaTimeRangeITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaTimeRangePTR().Pointer()
	}
	return nil
}

func QMediaTimeRangeFromPointer(ptr unsafe.Pointer) *QMediaTimeRange {
	var n = new(QMediaTimeRange)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMediaTimeRange) QMediaTimeRangePTR() *QMediaTimeRange {
	return ptr
}

func NewQMediaTimeRange() *QMediaTimeRange {
	return QMediaTimeRangeFromPointer(unsafe.Pointer(C.QMediaTimeRange_NewQMediaTimeRange()))
}

func NewQMediaTimeRange3(interval QMediaTimeIntervalITF) *QMediaTimeRange {
	return QMediaTimeRangeFromPointer(unsafe.Pointer(C.QMediaTimeRange_NewQMediaTimeRange3(C.QtObjectPtr(PointerFromQMediaTimeInterval(interval)))))
}

func NewQMediaTimeRange4(ran QMediaTimeRangeITF) *QMediaTimeRange {
	return QMediaTimeRangeFromPointer(unsafe.Pointer(C.QMediaTimeRange_NewQMediaTimeRange4(C.QtObjectPtr(PointerFromQMediaTimeRange(ran)))))
}

func (ptr *QMediaTimeRange) AddInterval(interval QMediaTimeIntervalITF) {
	if ptr.Pointer() != nil {
		C.QMediaTimeRange_AddInterval(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQMediaTimeInterval(interval)))
	}
}

func (ptr *QMediaTimeRange) AddTimeRange(ran QMediaTimeRangeITF) {
	if ptr.Pointer() != nil {
		C.QMediaTimeRange_AddTimeRange(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQMediaTimeRange(ran)))
	}
}

func (ptr *QMediaTimeRange) Clear() {
	if ptr.Pointer() != nil {
		C.QMediaTimeRange_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QMediaTimeRange) IsContinuous() bool {
	if ptr.Pointer() != nil {
		return C.QMediaTimeRange_IsContinuous(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMediaTimeRange) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QMediaTimeRange_IsEmpty(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMediaTimeRange) RemoveInterval(interval QMediaTimeIntervalITF) {
	if ptr.Pointer() != nil {
		C.QMediaTimeRange_RemoveInterval(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQMediaTimeInterval(interval)))
	}
}

func (ptr *QMediaTimeRange) RemoveTimeRange(ran QMediaTimeRangeITF) {
	if ptr.Pointer() != nil {
		C.QMediaTimeRange_RemoveTimeRange(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQMediaTimeRange(ran)))
	}
}

func (ptr *QMediaTimeRange) DestroyQMediaTimeRange() {
	if ptr.Pointer() != nil {
		C.QMediaTimeRange_DestroyQMediaTimeRange(C.QtObjectPtr(ptr.Pointer()))
	}
}
