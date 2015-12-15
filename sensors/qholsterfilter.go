package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QHolsterFilter struct {
	QSensorFilter
}

type QHolsterFilter_ITF interface {
	QSensorFilter_ITF
	QHolsterFilter_PTR() *QHolsterFilter
}

func PointerFromQHolsterFilter(ptr QHolsterFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHolsterFilter_PTR().Pointer()
	}
	return nil
}

func NewQHolsterFilterFromPointer(ptr unsafe.Pointer) *QHolsterFilter {
	var n = new(QHolsterFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QHolsterFilter) QHolsterFilter_PTR() *QHolsterFilter {
	return ptr
}

func (ptr *QHolsterFilter) Filter(reading QHolsterReading_ITF) bool {
	defer qt.Recovering("QHolsterFilter::filter")

	if ptr.Pointer() != nil {
		return C.QHolsterFilter_Filter(ptr.Pointer(), PointerFromQHolsterReading(reading)) != 0
	}
	return false
}
