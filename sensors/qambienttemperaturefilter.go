package sensors

//#include "qambienttemperaturefilter.h"
import "C"
import (
	"unsafe"
)

type QAmbientTemperatureFilter struct {
	QSensorFilter
}

type QAmbientTemperatureFilter_ITF interface {
	QSensorFilter_ITF
	QAmbientTemperatureFilter_PTR() *QAmbientTemperatureFilter
}

func PointerFromQAmbientTemperatureFilter(ptr QAmbientTemperatureFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAmbientTemperatureFilter_PTR().Pointer()
	}
	return nil
}

func NewQAmbientTemperatureFilterFromPointer(ptr unsafe.Pointer) *QAmbientTemperatureFilter {
	var n = new(QAmbientTemperatureFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAmbientTemperatureFilter) QAmbientTemperatureFilter_PTR() *QAmbientTemperatureFilter {
	return ptr
}

func (ptr *QAmbientTemperatureFilter) Filter(reading QAmbientTemperatureReading_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAmbientTemperatureFilter_Filter(ptr.Pointer(), PointerFromQAmbientTemperatureReading(reading)) != 0
	}
	return false
}
