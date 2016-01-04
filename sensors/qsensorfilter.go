package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QSensorFilter struct {
	ptr unsafe.Pointer
}

type QSensorFilter_ITF interface {
	QSensorFilter_PTR() *QSensorFilter
}

func (p *QSensorFilter) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSensorFilter) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSensorFilter(ptr QSensorFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func NewQSensorFilterFromPointer(ptr unsafe.Pointer) *QSensorFilter {
	var n = new(QSensorFilter)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QSensorFilter_") {
		n.SetObjectNameAbs("QSensorFilter_" + qt.Identifier())
	}
	return n
}

func (ptr *QSensorFilter) QSensorFilter_PTR() *QSensorFilter {
	return ptr
}

func (ptr *QSensorFilter) Filter(reading QSensorReading_ITF) bool {
	defer qt.Recovering("QSensorFilter::filter")

	if ptr.Pointer() != nil {
		return C.QSensorFilter_Filter(ptr.Pointer(), PointerFromQSensorReading(reading)) != 0
	}
	return false
}

func (ptr *QSensorFilter) DestroyQSensorFilter() {
	defer qt.Recovering("QSensorFilter::~QSensorFilter")

	if ptr.Pointer() != nil {
		C.QSensorFilter_DestroyQSensorFilter(ptr.Pointer())
	}
}

func (ptr *QSensorFilter) ObjectNameAbs() string {
	defer qt.Recovering("QSensorFilter::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSensorFilter_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSensorFilter) SetObjectNameAbs(name string) {
	defer qt.Recovering("QSensorFilter::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QSensorFilter_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
