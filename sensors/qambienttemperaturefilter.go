package sensors

//#include "qambienttemperaturefilter.h"
import "C"
import (
	"unsafe"
)

type QAmbientTemperatureFilter struct {
	QSensorFilter
}

type QAmbientTemperatureFilterITF interface {
	QSensorFilterITF
	QAmbientTemperatureFilterPTR() *QAmbientTemperatureFilter
}

func PointerFromQAmbientTemperatureFilter(ptr QAmbientTemperatureFilterITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAmbientTemperatureFilterPTR().Pointer()
	}
	return nil
}

func QAmbientTemperatureFilterFromPointer(ptr unsafe.Pointer) *QAmbientTemperatureFilter {
	var n = new(QAmbientTemperatureFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAmbientTemperatureFilter) QAmbientTemperatureFilterPTR() *QAmbientTemperatureFilter {
	return ptr
}

func (ptr *QAmbientTemperatureFilter) Filter(reading QAmbientTemperatureReadingITF) bool {
	if ptr.Pointer() != nil {
		return C.QAmbientTemperatureFilter_Filter(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAmbientTemperatureReading(reading))) != 0
	}
	return false
}
