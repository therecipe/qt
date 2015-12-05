package sensors

//#include "sensors.h"
import "C"
import (
	"log"
	"unsafe"
)

type QPressureFilter struct {
	QSensorFilter
}

type QPressureFilter_ITF interface {
	QSensorFilter_ITF
	QPressureFilter_PTR() *QPressureFilter
}

func PointerFromQPressureFilter(ptr QPressureFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPressureFilter_PTR().Pointer()
	}
	return nil
}

func NewQPressureFilterFromPointer(ptr unsafe.Pointer) *QPressureFilter {
	var n = new(QPressureFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPressureFilter) QPressureFilter_PTR() *QPressureFilter {
	return ptr
}

func (ptr *QPressureFilter) Filter(reading QPressureReading_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPressureFilter::filter")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QPressureFilter_Filter(ptr.Pointer(), PointerFromQPressureReading(reading)) != 0
	}
	return false
}
