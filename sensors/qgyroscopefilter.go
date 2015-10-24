package sensors

//#include "qgyroscopefilter.h"
import "C"
import (
	"unsafe"
)

type QGyroscopeFilter struct {
	QSensorFilter
}

type QGyroscopeFilterITF interface {
	QSensorFilterITF
	QGyroscopeFilterPTR() *QGyroscopeFilter
}

func PointerFromQGyroscopeFilter(ptr QGyroscopeFilterITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGyroscopeFilterPTR().Pointer()
	}
	return nil
}

func QGyroscopeFilterFromPointer(ptr unsafe.Pointer) *QGyroscopeFilter {
	var n = new(QGyroscopeFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGyroscopeFilter) QGyroscopeFilterPTR() *QGyroscopeFilter {
	return ptr
}

func (ptr *QGyroscopeFilter) Filter(reading QGyroscopeReadingITF) bool {
	if ptr.Pointer() != nil {
		return C.QGyroscopeFilter_Filter(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGyroscopeReading(reading))) != 0
	}
	return false
}
