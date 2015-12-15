package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QTapFilter struct {
	QSensorFilter
}

type QTapFilter_ITF interface {
	QSensorFilter_ITF
	QTapFilter_PTR() *QTapFilter
}

func PointerFromQTapFilter(ptr QTapFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTapFilter_PTR().Pointer()
	}
	return nil
}

func NewQTapFilterFromPointer(ptr unsafe.Pointer) *QTapFilter {
	var n = new(QTapFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTapFilter) QTapFilter_PTR() *QTapFilter {
	return ptr
}

func (ptr *QTapFilter) Filter(reading QTapReading_ITF) bool {
	defer qt.Recovering("QTapFilter::filter")

	if ptr.Pointer() != nil {
		return C.QTapFilter_Filter(ptr.Pointer(), PointerFromQTapReading(reading)) != 0
	}
	return false
}
