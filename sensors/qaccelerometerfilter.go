package sensors

//#include "qaccelerometerfilter.h"
import "C"
import (
	"unsafe"
)

type QAccelerometerFilter struct {
	QSensorFilter
}

type QAccelerometerFilterITF interface {
	QSensorFilterITF
	QAccelerometerFilterPTR() *QAccelerometerFilter
}

func PointerFromQAccelerometerFilter(ptr QAccelerometerFilterITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccelerometerFilterPTR().Pointer()
	}
	return nil
}

func QAccelerometerFilterFromPointer(ptr unsafe.Pointer) *QAccelerometerFilter {
	var n = new(QAccelerometerFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAccelerometerFilter) QAccelerometerFilterPTR() *QAccelerometerFilter {
	return ptr
}

func (ptr *QAccelerometerFilter) Filter(reading QAccelerometerReadingITF) bool {
	if ptr.Pointer() != nil {
		return C.QAccelerometerFilter_Filter(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAccelerometerReading(reading))) != 0
	}
	return false
}
