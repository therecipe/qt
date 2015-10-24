package multimedia

//#include "qmediatimeinterval.h"
import "C"
import (
	"unsafe"
)

type QMediaTimeInterval struct {
	ptr unsafe.Pointer
}

type QMediaTimeIntervalITF interface {
	QMediaTimeIntervalPTR() *QMediaTimeInterval
}

func (p *QMediaTimeInterval) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMediaTimeInterval) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMediaTimeInterval(ptr QMediaTimeIntervalITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaTimeIntervalPTR().Pointer()
	}
	return nil
}

func QMediaTimeIntervalFromPointer(ptr unsafe.Pointer) *QMediaTimeInterval {
	var n = new(QMediaTimeInterval)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMediaTimeInterval) QMediaTimeIntervalPTR() *QMediaTimeInterval {
	return ptr
}

func NewQMediaTimeInterval() *QMediaTimeInterval {
	return QMediaTimeIntervalFromPointer(unsafe.Pointer(C.QMediaTimeInterval_NewQMediaTimeInterval()))
}

func NewQMediaTimeInterval3(other QMediaTimeIntervalITF) *QMediaTimeInterval {
	return QMediaTimeIntervalFromPointer(unsafe.Pointer(C.QMediaTimeInterval_NewQMediaTimeInterval3(C.QtObjectPtr(PointerFromQMediaTimeInterval(other)))))
}

func (ptr *QMediaTimeInterval) IsNormal() bool {
	if ptr.Pointer() != nil {
		return C.QMediaTimeInterval_IsNormal(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}
