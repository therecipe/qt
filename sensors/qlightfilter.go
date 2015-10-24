package sensors

//#include "qlightfilter.h"
import "C"
import (
	"unsafe"
)

type QLightFilter struct {
	QSensorFilter
}

type QLightFilterITF interface {
	QSensorFilterITF
	QLightFilterPTR() *QLightFilter
}

func PointerFromQLightFilter(ptr QLightFilterITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLightFilterPTR().Pointer()
	}
	return nil
}

func QLightFilterFromPointer(ptr unsafe.Pointer) *QLightFilter {
	var n = new(QLightFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QLightFilter) QLightFilterPTR() *QLightFilter {
	return ptr
}

func (ptr *QLightFilter) Filter(reading QLightReadingITF) bool {
	if ptr.Pointer() != nil {
		return C.QLightFilter_Filter(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQLightReading(reading))) != 0
	}
	return false
}
