package sensors

//#include "qproximityfilter.h"
import "C"
import (
	"unsafe"
)

type QProximityFilter struct {
	QSensorFilter
}

type QProximityFilter_ITF interface {
	QSensorFilter_ITF
	QProximityFilter_PTR() *QProximityFilter
}

func PointerFromQProximityFilter(ptr QProximityFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QProximityFilter_PTR().Pointer()
	}
	return nil
}

func NewQProximityFilterFromPointer(ptr unsafe.Pointer) *QProximityFilter {
	var n = new(QProximityFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QProximityFilter) QProximityFilter_PTR() *QProximityFilter {
	return ptr
}

func (ptr *QProximityFilter) Filter(reading QProximityReading_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QProximityFilter_Filter(ptr.Pointer(), PointerFromQProximityReading(reading)) != 0
	}
	return false
}
