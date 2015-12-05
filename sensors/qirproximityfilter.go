package sensors

//#include "sensors.h"
import "C"
import (
	"log"
	"unsafe"
)

type QIRProximityFilter struct {
	QSensorFilter
}

type QIRProximityFilter_ITF interface {
	QSensorFilter_ITF
	QIRProximityFilter_PTR() *QIRProximityFilter
}

func PointerFromQIRProximityFilter(ptr QIRProximityFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QIRProximityFilter_PTR().Pointer()
	}
	return nil
}

func NewQIRProximityFilterFromPointer(ptr unsafe.Pointer) *QIRProximityFilter {
	var n = new(QIRProximityFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QIRProximityFilter) QIRProximityFilter_PTR() *QIRProximityFilter {
	return ptr
}

func (ptr *QIRProximityFilter) Filter(reading QIRProximityReading_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QIRProximityFilter::filter")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QIRProximityFilter_Filter(ptr.Pointer(), PointerFromQIRProximityReading(reading)) != 0
	}
	return false
}
