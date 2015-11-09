package sensors

//#include "qsensorfilter.h"
import "C"
import (
	"unsafe"
)

type QSensorFilter struct {
	ptr unsafe.Pointer
}

type QSensorFilter_ITF interface {
	QSensorFilter_PTR() *QSensorFilter
}

func (p *QSensorFilter) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSensorFilter) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSensorFilter(ptr QSensorFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func NewQSensorFilterFromPointer(ptr unsafe.Pointer) *QSensorFilter {
	var n = new(QSensorFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSensorFilter) QSensorFilter_PTR() *QSensorFilter {
	return ptr
}

func (ptr *QSensorFilter) Filter(reading QSensorReading_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QSensorFilter_Filter(ptr.Pointer(), PointerFromQSensorReading(reading)) != 0
	}
	return false
}
