package sensors

//#include "qaccelerometerfilter.h"
import "C"
import (
	"unsafe"
)

type QAccelerometerFilter struct {
	QSensorFilter
}

type QAccelerometerFilter_ITF interface {
	QSensorFilter_ITF
	QAccelerometerFilter_PTR() *QAccelerometerFilter
}

func PointerFromQAccelerometerFilter(ptr QAccelerometerFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccelerometerFilter_PTR().Pointer()
	}
	return nil
}

func NewQAccelerometerFilterFromPointer(ptr unsafe.Pointer) *QAccelerometerFilter {
	var n = new(QAccelerometerFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAccelerometerFilter) QAccelerometerFilter_PTR() *QAccelerometerFilter {
	return ptr
}

func (ptr *QAccelerometerFilter) Filter(reading QAccelerometerReading_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAccelerometerFilter_Filter(ptr.Pointer(), PointerFromQAccelerometerReading(reading)) != 0
	}
	return false
}
