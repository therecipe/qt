package sensors

//#include "qrotationfilter.h"
import "C"
import (
	"unsafe"
)

type QRotationFilter struct {
	QSensorFilter
}

type QRotationFilterITF interface {
	QSensorFilterITF
	QRotationFilterPTR() *QRotationFilter
}

func PointerFromQRotationFilter(ptr QRotationFilterITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRotationFilterPTR().Pointer()
	}
	return nil
}

func QRotationFilterFromPointer(ptr unsafe.Pointer) *QRotationFilter {
	var n = new(QRotationFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QRotationFilter) QRotationFilterPTR() *QRotationFilter {
	return ptr
}

func (ptr *QRotationFilter) Filter(reading QRotationReadingITF) bool {
	if ptr.Pointer() != nil {
		return C.QRotationFilter_Filter(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQRotationReading(reading))) != 0
	}
	return false
}
