package location

//#include "qgeoroutesegment.h"
import "C"
import (
	"unsafe"
)

type QGeoRouteSegment struct {
	ptr unsafe.Pointer
}

type QGeoRouteSegmentITF interface {
	QGeoRouteSegmentPTR() *QGeoRouteSegment
}

func (p *QGeoRouteSegment) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QGeoRouteSegment) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQGeoRouteSegment(ptr QGeoRouteSegmentITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoRouteSegmentPTR().Pointer()
	}
	return nil
}

func QGeoRouteSegmentFromPointer(ptr unsafe.Pointer) *QGeoRouteSegment {
	var n = new(QGeoRouteSegment)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGeoRouteSegment) QGeoRouteSegmentPTR() *QGeoRouteSegment {
	return ptr
}

func NewQGeoRouteSegment() *QGeoRouteSegment {
	return QGeoRouteSegmentFromPointer(unsafe.Pointer(C.QGeoRouteSegment_NewQGeoRouteSegment()))
}

func NewQGeoRouteSegment2(other QGeoRouteSegmentITF) *QGeoRouteSegment {
	return QGeoRouteSegmentFromPointer(unsafe.Pointer(C.QGeoRouteSegment_NewQGeoRouteSegment2(C.QtObjectPtr(PointerFromQGeoRouteSegment(other)))))
}

func (ptr *QGeoRouteSegment) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QGeoRouteSegment_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGeoRouteSegment) SetManeuver(maneuver QGeoManeuverITF) {
	if ptr.Pointer() != nil {
		C.QGeoRouteSegment_SetManeuver(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGeoManeuver(maneuver)))
	}
}

func (ptr *QGeoRouteSegment) SetNextRouteSegment(routeSegment QGeoRouteSegmentITF) {
	if ptr.Pointer() != nil {
		C.QGeoRouteSegment_SetNextRouteSegment(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGeoRouteSegment(routeSegment)))
	}
}

func (ptr *QGeoRouteSegment) SetTravelTime(secs int) {
	if ptr.Pointer() != nil {
		C.QGeoRouteSegment_SetTravelTime(C.QtObjectPtr(ptr.Pointer()), C.int(secs))
	}
}

func (ptr *QGeoRouteSegment) TravelTime() int {
	if ptr.Pointer() != nil {
		return int(C.QGeoRouteSegment_TravelTime(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoRouteSegment) DestroyQGeoRouteSegment() {
	if ptr.Pointer() != nil {
		C.QGeoRouteSegment_DestroyQGeoRouteSegment(C.QtObjectPtr(ptr.Pointer()))
	}
}
