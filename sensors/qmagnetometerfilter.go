package sensors

//#include "qmagnetometerfilter.h"
import "C"
import (
	"unsafe"
)

type QMagnetometerFilter struct {
	QSensorFilter
}

type QMagnetometerFilterITF interface {
	QSensorFilterITF
	QMagnetometerFilterPTR() *QMagnetometerFilter
}

func PointerFromQMagnetometerFilter(ptr QMagnetometerFilterITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMagnetometerFilterPTR().Pointer()
	}
	return nil
}

func QMagnetometerFilterFromPointer(ptr unsafe.Pointer) *QMagnetometerFilter {
	var n = new(QMagnetometerFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMagnetometerFilter) QMagnetometerFilterPTR() *QMagnetometerFilter {
	return ptr
}

func (ptr *QMagnetometerFilter) Filter(reading QMagnetometerReadingITF) bool {
	if ptr.Pointer() != nil {
		return C.QMagnetometerFilter_Filter(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQMagnetometerReading(reading))) != 0
	}
	return false
}
