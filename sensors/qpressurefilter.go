package sensors

//#include "qpressurefilter.h"
import "C"
import (
	"unsafe"
)

type QPressureFilter struct {
	QSensorFilter
}

type QPressureFilterITF interface {
	QSensorFilterITF
	QPressureFilterPTR() *QPressureFilter
}

func PointerFromQPressureFilter(ptr QPressureFilterITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPressureFilterPTR().Pointer()
	}
	return nil
}

func QPressureFilterFromPointer(ptr unsafe.Pointer) *QPressureFilter {
	var n = new(QPressureFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPressureFilter) QPressureFilterPTR() *QPressureFilter {
	return ptr
}

func (ptr *QPressureFilter) Filter(reading QPressureReadingITF) bool {
	if ptr.Pointer() != nil {
		return C.QPressureFilter_Filter(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPressureReading(reading))) != 0
	}
	return false
}
