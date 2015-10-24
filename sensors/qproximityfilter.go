package sensors

//#include "qproximityfilter.h"
import "C"
import (
	"unsafe"
)

type QProximityFilter struct {
	QSensorFilter
}

type QProximityFilterITF interface {
	QSensorFilterITF
	QProximityFilterPTR() *QProximityFilter
}

func PointerFromQProximityFilter(ptr QProximityFilterITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QProximityFilterPTR().Pointer()
	}
	return nil
}

func QProximityFilterFromPointer(ptr unsafe.Pointer) *QProximityFilter {
	var n = new(QProximityFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QProximityFilter) QProximityFilterPTR() *QProximityFilter {
	return ptr
}

func (ptr *QProximityFilter) Filter(reading QProximityReadingITF) bool {
	if ptr.Pointer() != nil {
		return C.QProximityFilter_Filter(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQProximityReading(reading))) != 0
	}
	return false
}
