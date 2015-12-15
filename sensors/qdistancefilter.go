package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QDistanceFilter struct {
	QSensorFilter
}

type QDistanceFilter_ITF interface {
	QSensorFilter_ITF
	QDistanceFilter_PTR() *QDistanceFilter
}

func PointerFromQDistanceFilter(ptr QDistanceFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDistanceFilter_PTR().Pointer()
	}
	return nil
}

func NewQDistanceFilterFromPointer(ptr unsafe.Pointer) *QDistanceFilter {
	var n = new(QDistanceFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDistanceFilter) QDistanceFilter_PTR() *QDistanceFilter {
	return ptr
}

func (ptr *QDistanceFilter) Filter(reading QDistanceReading_ITF) bool {
	defer qt.Recovering("QDistanceFilter::filter")

	if ptr.Pointer() != nil {
		return C.QDistanceFilter_Filter(ptr.Pointer(), PointerFromQDistanceReading(reading)) != 0
	}
	return false
}
