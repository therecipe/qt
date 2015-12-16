package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QMediaTimeInterval struct {
	ptr unsafe.Pointer
}

type QMediaTimeInterval_ITF interface {
	QMediaTimeInterval_PTR() *QMediaTimeInterval
}

func (p *QMediaTimeInterval) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMediaTimeInterval) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMediaTimeInterval(ptr QMediaTimeInterval_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaTimeInterval_PTR().Pointer()
	}
	return nil
}

func NewQMediaTimeIntervalFromPointer(ptr unsafe.Pointer) *QMediaTimeInterval {
	var n = new(QMediaTimeInterval)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMediaTimeInterval) QMediaTimeInterval_PTR() *QMediaTimeInterval {
	return ptr
}

func NewQMediaTimeInterval() *QMediaTimeInterval {
	defer qt.Recovering("QMediaTimeInterval::QMediaTimeInterval")

	return NewQMediaTimeIntervalFromPointer(C.QMediaTimeInterval_NewQMediaTimeInterval())
}

func NewQMediaTimeInterval3(other QMediaTimeInterval_ITF) *QMediaTimeInterval {
	defer qt.Recovering("QMediaTimeInterval::QMediaTimeInterval")

	return NewQMediaTimeIntervalFromPointer(C.QMediaTimeInterval_NewQMediaTimeInterval3(PointerFromQMediaTimeInterval(other)))
}

func NewQMediaTimeInterval2(start int64, end int64) *QMediaTimeInterval {
	defer qt.Recovering("QMediaTimeInterval::QMediaTimeInterval")

	return NewQMediaTimeIntervalFromPointer(C.QMediaTimeInterval_NewQMediaTimeInterval2(C.longlong(start), C.longlong(end)))
}

func (ptr *QMediaTimeInterval) Contains(time int64) bool {
	defer qt.Recovering("QMediaTimeInterval::contains")

	if ptr.Pointer() != nil {
		return C.QMediaTimeInterval_Contains(ptr.Pointer(), C.longlong(time)) != 0
	}
	return false
}

func (ptr *QMediaTimeInterval) End() int64 {
	defer qt.Recovering("QMediaTimeInterval::end")

	if ptr.Pointer() != nil {
		return int64(C.QMediaTimeInterval_End(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaTimeInterval) IsNormal() bool {
	defer qt.Recovering("QMediaTimeInterval::isNormal")

	if ptr.Pointer() != nil {
		return C.QMediaTimeInterval_IsNormal(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaTimeInterval) Start() int64 {
	defer qt.Recovering("QMediaTimeInterval::start")

	if ptr.Pointer() != nil {
		return int64(C.QMediaTimeInterval_Start(ptr.Pointer()))
	}
	return 0
}
