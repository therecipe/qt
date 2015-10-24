package sensors

//#include "qambientlightfilter.h"
import "C"
import (
	"unsafe"
)

type QAmbientLightFilter struct {
	QSensorFilter
}

type QAmbientLightFilterITF interface {
	QSensorFilterITF
	QAmbientLightFilterPTR() *QAmbientLightFilter
}

func PointerFromQAmbientLightFilter(ptr QAmbientLightFilterITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAmbientLightFilterPTR().Pointer()
	}
	return nil
}

func QAmbientLightFilterFromPointer(ptr unsafe.Pointer) *QAmbientLightFilter {
	var n = new(QAmbientLightFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAmbientLightFilter) QAmbientLightFilterPTR() *QAmbientLightFilter {
	return ptr
}

func (ptr *QAmbientLightFilter) Filter(reading QAmbientLightReadingITF) bool {
	if ptr.Pointer() != nil {
		return C.QAmbientLightFilter_Filter(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAmbientLightReading(reading))) != 0
	}
	return false
}
