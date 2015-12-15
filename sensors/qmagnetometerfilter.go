package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QMagnetometerFilter struct {
	QSensorFilter
}

type QMagnetometerFilter_ITF interface {
	QSensorFilter_ITF
	QMagnetometerFilter_PTR() *QMagnetometerFilter
}

func PointerFromQMagnetometerFilter(ptr QMagnetometerFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMagnetometerFilter_PTR().Pointer()
	}
	return nil
}

func NewQMagnetometerFilterFromPointer(ptr unsafe.Pointer) *QMagnetometerFilter {
	var n = new(QMagnetometerFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMagnetometerFilter) QMagnetometerFilter_PTR() *QMagnetometerFilter {
	return ptr
}

func (ptr *QMagnetometerFilter) Filter(reading QMagnetometerReading_ITF) bool {
	defer qt.Recovering("QMagnetometerFilter::filter")

	if ptr.Pointer() != nil {
		return C.QMagnetometerFilter_Filter(ptr.Pointer(), PointerFromQMagnetometerReading(reading)) != 0
	}
	return false
}
