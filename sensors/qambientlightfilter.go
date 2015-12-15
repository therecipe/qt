package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QAmbientLightFilter struct {
	QSensorFilter
}

type QAmbientLightFilter_ITF interface {
	QSensorFilter_ITF
	QAmbientLightFilter_PTR() *QAmbientLightFilter
}

func PointerFromQAmbientLightFilter(ptr QAmbientLightFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAmbientLightFilter_PTR().Pointer()
	}
	return nil
}

func NewQAmbientLightFilterFromPointer(ptr unsafe.Pointer) *QAmbientLightFilter {
	var n = new(QAmbientLightFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAmbientLightFilter) QAmbientLightFilter_PTR() *QAmbientLightFilter {
	return ptr
}

func (ptr *QAmbientLightFilter) Filter(reading QAmbientLightReading_ITF) bool {
	defer qt.Recovering("QAmbientLightFilter::filter")

	if ptr.Pointer() != nil {
		return C.QAmbientLightFilter_Filter(ptr.Pointer(), PointerFromQAmbientLightReading(reading)) != 0
	}
	return false
}
