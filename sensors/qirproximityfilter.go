package sensors

//#include "qirproximityfilter.h"
import "C"
import (
	"unsafe"
)

type QIRProximityFilter struct {
	QSensorFilter
}

type QIRProximityFilterITF interface {
	QSensorFilterITF
	QIRProximityFilterPTR() *QIRProximityFilter
}

func PointerFromQIRProximityFilter(ptr QIRProximityFilterITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QIRProximityFilterPTR().Pointer()
	}
	return nil
}

func QIRProximityFilterFromPointer(ptr unsafe.Pointer) *QIRProximityFilter {
	var n = new(QIRProximityFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QIRProximityFilter) QIRProximityFilterPTR() *QIRProximityFilter {
	return ptr
}

func (ptr *QIRProximityFilter) Filter(reading QIRProximityReadingITF) bool {
	if ptr.Pointer() != nil {
		return C.QIRProximityFilter_Filter(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQIRProximityReading(reading))) != 0
	}
	return false
}
