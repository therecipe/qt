package location

//#include "qgeoroute.h"
import "C"
import (
	"github.com/therecipe/qt/positioning"
	"unsafe"
)

type QGeoRoute struct {
	ptr unsafe.Pointer
}

type QGeoRoute_ITF interface {
	QGeoRoute_PTR() *QGeoRoute
}

func (p *QGeoRoute) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QGeoRoute) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQGeoRoute(ptr QGeoRoute_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoRoute_PTR().Pointer()
	}
	return nil
}

func NewQGeoRouteFromPointer(ptr unsafe.Pointer) *QGeoRoute {
	var n = new(QGeoRoute)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGeoRoute) QGeoRoute_PTR() *QGeoRoute {
	return ptr
}

func NewQGeoRoute() *QGeoRoute {
	return NewQGeoRouteFromPointer(C.QGeoRoute_NewQGeoRoute())
}

func NewQGeoRoute2(other QGeoRoute_ITF) *QGeoRoute {
	return NewQGeoRouteFromPointer(C.QGeoRoute_NewQGeoRoute2(PointerFromQGeoRoute(other)))
}

func (ptr *QGeoRoute) Distance() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGeoRoute_Distance(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRoute) RouteId() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoRoute_RouteId(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoRoute) SetBounds(bounds positioning.QGeoRectangle_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRoute_SetBounds(ptr.Pointer(), positioning.PointerFromQGeoRectangle(bounds))
	}
}

func (ptr *QGeoRoute) SetDistance(distance float64) {
	if ptr.Pointer() != nil {
		C.QGeoRoute_SetDistance(ptr.Pointer(), C.double(distance))
	}
}

func (ptr *QGeoRoute) SetFirstRouteSegment(routeSegment QGeoRouteSegment_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRoute_SetFirstRouteSegment(ptr.Pointer(), PointerFromQGeoRouteSegment(routeSegment))
	}
}

func (ptr *QGeoRoute) SetRequest(request QGeoRouteRequest_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoRoute_SetRequest(ptr.Pointer(), PointerFromQGeoRouteRequest(request))
	}
}

func (ptr *QGeoRoute) SetRouteId(id string) {
	if ptr.Pointer() != nil {
		C.QGeoRoute_SetRouteId(ptr.Pointer(), C.CString(id))
	}
}

func (ptr *QGeoRoute) SetTravelMode(mode QGeoRouteRequest__TravelMode) {
	if ptr.Pointer() != nil {
		C.QGeoRoute_SetTravelMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QGeoRoute) SetTravelTime(secs int) {
	if ptr.Pointer() != nil {
		C.QGeoRoute_SetTravelTime(ptr.Pointer(), C.int(secs))
	}
}

func (ptr *QGeoRoute) TravelMode() QGeoRouteRequest__TravelMode {
	if ptr.Pointer() != nil {
		return QGeoRouteRequest__TravelMode(C.QGeoRoute_TravelMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRoute) TravelTime() int {
	if ptr.Pointer() != nil {
		return int(C.QGeoRoute_TravelTime(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRoute) DestroyQGeoRoute() {
	if ptr.Pointer() != nil {
		C.QGeoRoute_DestroyQGeoRoute(ptr.Pointer())
	}
}
