package sensors

//#include "qtiltfilter.h"
import "C"
import (
	"unsafe"
)

type QTiltFilter struct {
	QSensorFilter
}

type QTiltFilterITF interface {
	QSensorFilterITF
	QTiltFilterPTR() *QTiltFilter
}

func PointerFromQTiltFilter(ptr QTiltFilterITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTiltFilterPTR().Pointer()
	}
	return nil
}

func QTiltFilterFromPointer(ptr unsafe.Pointer) *QTiltFilter {
	var n = new(QTiltFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTiltFilter) QTiltFilterPTR() *QTiltFilter {
	return ptr
}

func (ptr *QTiltFilter) Filter(reading QTiltReadingITF) bool {
	if ptr.Pointer() != nil {
		return C.QTiltFilter_Filter(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTiltReading(reading))) != 0
	}
	return false
}
