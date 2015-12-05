package sensors

//#include "sensors.h"
import "C"
import (
	"log"
	"unsafe"
)

type QTiltFilter struct {
	QSensorFilter
}

type QTiltFilter_ITF interface {
	QSensorFilter_ITF
	QTiltFilter_PTR() *QTiltFilter
}

func PointerFromQTiltFilter(ptr QTiltFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTiltFilter_PTR().Pointer()
	}
	return nil
}

func NewQTiltFilterFromPointer(ptr unsafe.Pointer) *QTiltFilter {
	var n = new(QTiltFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTiltFilter) QTiltFilter_PTR() *QTiltFilter {
	return ptr
}

func (ptr *QTiltFilter) Filter(reading QTiltReading_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTiltFilter::filter")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTiltFilter_Filter(ptr.Pointer(), PointerFromQTiltReading(reading)) != 0
	}
	return false
}
