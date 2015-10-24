package sensors

//#include "qtapfilter.h"
import "C"
import (
	"unsafe"
)

type QTapFilter struct {
	QSensorFilter
}

type QTapFilterITF interface {
	QSensorFilterITF
	QTapFilterPTR() *QTapFilter
}

func PointerFromQTapFilter(ptr QTapFilterITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTapFilterPTR().Pointer()
	}
	return nil
}

func QTapFilterFromPointer(ptr unsafe.Pointer) *QTapFilter {
	var n = new(QTapFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTapFilter) QTapFilterPTR() *QTapFilter {
	return ptr
}

func (ptr *QTapFilter) Filter(reading QTapReadingITF) bool {
	if ptr.Pointer() != nil {
		return C.QTapFilter_Filter(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTapReading(reading))) != 0
	}
	return false
}
