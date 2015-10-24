package sensors

//#include "qcompassfilter.h"
import "C"
import (
	"unsafe"
)

type QCompassFilter struct {
	QSensorFilter
}

type QCompassFilterITF interface {
	QSensorFilterITF
	QCompassFilterPTR() *QCompassFilter
}

func PointerFromQCompassFilter(ptr QCompassFilterITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCompassFilterPTR().Pointer()
	}
	return nil
}

func QCompassFilterFromPointer(ptr unsafe.Pointer) *QCompassFilter {
	var n = new(QCompassFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QCompassFilter) QCompassFilterPTR() *QCompassFilter {
	return ptr
}

func (ptr *QCompassFilter) Filter(reading QCompassReadingITF) bool {
	if ptr.Pointer() != nil {
		return C.QCompassFilter_Filter(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQCompassReading(reading))) != 0
	}
	return false
}
