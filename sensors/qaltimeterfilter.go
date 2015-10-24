package sensors

//#include "qaltimeterfilter.h"
import "C"
import (
	"unsafe"
)

type QAltimeterFilter struct {
	QSensorFilter
}

type QAltimeterFilterITF interface {
	QSensorFilterITF
	QAltimeterFilterPTR() *QAltimeterFilter
}

func PointerFromQAltimeterFilter(ptr QAltimeterFilterITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAltimeterFilterPTR().Pointer()
	}
	return nil
}

func QAltimeterFilterFromPointer(ptr unsafe.Pointer) *QAltimeterFilter {
	var n = new(QAltimeterFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAltimeterFilter) QAltimeterFilterPTR() *QAltimeterFilter {
	return ptr
}

func (ptr *QAltimeterFilter) Filter(reading QAltimeterReadingITF) bool {
	if ptr.Pointer() != nil {
		return C.QAltimeterFilter_Filter(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAltimeterReading(reading))) != 0
	}
	return false
}
