package sensors

//#include "qholsterfilter.h"
import "C"
import (
	"unsafe"
)

type QHolsterFilter struct {
	QSensorFilter
}

type QHolsterFilterITF interface {
	QSensorFilterITF
	QHolsterFilterPTR() *QHolsterFilter
}

func PointerFromQHolsterFilter(ptr QHolsterFilterITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHolsterFilterPTR().Pointer()
	}
	return nil
}

func QHolsterFilterFromPointer(ptr unsafe.Pointer) *QHolsterFilter {
	var n = new(QHolsterFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QHolsterFilter) QHolsterFilterPTR() *QHolsterFilter {
	return ptr
}

func (ptr *QHolsterFilter) Filter(reading QHolsterReadingITF) bool {
	if ptr.Pointer() != nil {
		return C.QHolsterFilter_Filter(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQHolsterReading(reading))) != 0
	}
	return false
}
