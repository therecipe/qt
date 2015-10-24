package sensors

//#include "qorientationfilter.h"
import "C"
import (
	"unsafe"
)

type QOrientationFilter struct {
	QSensorFilter
}

type QOrientationFilterITF interface {
	QSensorFilterITF
	QOrientationFilterPTR() *QOrientationFilter
}

func PointerFromQOrientationFilter(ptr QOrientationFilterITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOrientationFilterPTR().Pointer()
	}
	return nil
}

func QOrientationFilterFromPointer(ptr unsafe.Pointer) *QOrientationFilter {
	var n = new(QOrientationFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOrientationFilter) QOrientationFilterPTR() *QOrientationFilter {
	return ptr
}

func (ptr *QOrientationFilter) Filter(reading QOrientationReadingITF) bool {
	if ptr.Pointer() != nil {
		return C.QOrientationFilter_Filter(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQOrientationReading(reading))) != 0
	}
	return false
}
