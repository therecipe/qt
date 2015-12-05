package sensors

//#include "sensors.h"
import "C"
import (
	"log"
	"unsafe"
)

type QAltimeterFilter struct {
	QSensorFilter
}

type QAltimeterFilter_ITF interface {
	QSensorFilter_ITF
	QAltimeterFilter_PTR() *QAltimeterFilter
}

func PointerFromQAltimeterFilter(ptr QAltimeterFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAltimeterFilter_PTR().Pointer()
	}
	return nil
}

func NewQAltimeterFilterFromPointer(ptr unsafe.Pointer) *QAltimeterFilter {
	var n = new(QAltimeterFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAltimeterFilter) QAltimeterFilter_PTR() *QAltimeterFilter {
	return ptr
}

func (ptr *QAltimeterFilter) Filter(reading QAltimeterReading_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAltimeterFilter::filter")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAltimeterFilter_Filter(ptr.Pointer(), PointerFromQAltimeterReading(reading)) != 0
	}
	return false
}
