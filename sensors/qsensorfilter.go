package sensors

//#include "qsensorfilter.h"
import "C"
import (
	"unsafe"
)

type QSensorFilter struct {
	ptr unsafe.Pointer
}

type QSensorFilterITF interface {
	QSensorFilterPTR() *QSensorFilter
}

func (p *QSensorFilter) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSensorFilter) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSensorFilter(ptr QSensorFilterITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorFilterPTR().Pointer()
	}
	return nil
}

func QSensorFilterFromPointer(ptr unsafe.Pointer) *QSensorFilter {
	var n = new(QSensorFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSensorFilter) QSensorFilterPTR() *QSensorFilter {
	return ptr
}

func (ptr *QSensorFilter) Filter(reading QSensorReadingITF) bool {
	if ptr.Pointer() != nil {
		return C.QSensorFilter_Filter(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSensorReading(reading))) != 0
	}
	return false
}
