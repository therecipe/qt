package sensors

//#include "qdistancefilter.h"
import "C"
import (
	"unsafe"
)

type QDistanceFilter struct {
	QSensorFilter
}

type QDistanceFilterITF interface {
	QSensorFilterITF
	QDistanceFilterPTR() *QDistanceFilter
}

func PointerFromQDistanceFilter(ptr QDistanceFilterITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDistanceFilterPTR().Pointer()
	}
	return nil
}

func QDistanceFilterFromPointer(ptr unsafe.Pointer) *QDistanceFilter {
	var n = new(QDistanceFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDistanceFilter) QDistanceFilterPTR() *QDistanceFilter {
	return ptr
}

func (ptr *QDistanceFilter) Filter(reading QDistanceReadingITF) bool {
	if ptr.Pointer() != nil {
		return C.QDistanceFilter_Filter(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQDistanceReading(reading))) != 0
	}
	return false
}
