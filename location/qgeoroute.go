package location

//#include "location.h"
import "C"
import (
	"github.com/therecipe/qt"
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
	defer qt.Recovering("QGeoRoute::QGeoRoute")

	return NewQGeoRouteFromPointer(C.QGeoRoute_NewQGeoRoute())
}

func NewQGeoRoute2(other QGeoRoute_ITF) *QGeoRoute {
	defer qt.Recovering("QGeoRoute::QGeoRoute")

	return NewQGeoRouteFromPointer(C.QGeoRoute_NewQGeoRoute2(PointerFromQGeoRoute(other)))
}

func (ptr *QGeoRoute) Distance() float64 {
	defer qt.Recovering("QGeoRoute::distance")

	if ptr.Pointer() != nil {
		return float64(C.QGeoRoute_Distance(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRoute) RouteId() string {
	defer qt.Recovering("QGeoRoute::routeId")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoRoute_RouteId(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoRoute) SetBounds(bounds positioning.QGeoRectangle_ITF) {
	defer qt.Recovering("QGeoRoute::setBounds")

	if ptr.Pointer() != nil {
		C.QGeoRoute_SetBounds(ptr.Pointer(), positioning.PointerFromQGeoRectangle(bounds))
	}
}

func (ptr *QGeoRoute) SetDistance(distance float64) {
	defer qt.Recovering("QGeoRoute::setDistance")

	if ptr.Pointer() != nil {
		C.QGeoRoute_SetDistance(ptr.Pointer(), C.double(distance))
	}
}

func (ptr *QGeoRoute) SetFirstRouteSegment(routeSegment QGeoRouteSegment_ITF) {
	defer qt.Recovering("QGeoRoute::setFirstRouteSegment")

	if ptr.Pointer() != nil {
		C.QGeoRoute_SetFirstRouteSegment(ptr.Pointer(), PointerFromQGeoRouteSegment(routeSegment))
	}
}

func (ptr *QGeoRoute) SetRequest(request QGeoRouteRequest_ITF) {
	defer qt.Recovering("QGeoRoute::setRequest")

	if ptr.Pointer() != nil {
		C.QGeoRoute_SetRequest(ptr.Pointer(), PointerFromQGeoRouteRequest(request))
	}
}

func (ptr *QGeoRoute) SetRouteId(id string) {
	defer qt.Recovering("QGeoRoute::setRouteId")

	if ptr.Pointer() != nil {
		C.QGeoRoute_SetRouteId(ptr.Pointer(), C.CString(id))
	}
}

func (ptr *QGeoRoute) SetTravelMode(mode QGeoRouteRequest__TravelMode) {
	defer qt.Recovering("QGeoRoute::setTravelMode")

	if ptr.Pointer() != nil {
		C.QGeoRoute_SetTravelMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QGeoRoute) SetTravelTime(secs int) {
	defer qt.Recovering("QGeoRoute::setTravelTime")

	if ptr.Pointer() != nil {
		C.QGeoRoute_SetTravelTime(ptr.Pointer(), C.int(secs))
	}
}

func (ptr *QGeoRoute) TravelMode() QGeoRouteRequest__TravelMode {
	defer qt.Recovering("QGeoRoute::travelMode")

	if ptr.Pointer() != nil {
		return QGeoRouteRequest__TravelMode(C.QGeoRoute_TravelMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRoute) TravelTime() int {
	defer qt.Recovering("QGeoRoute::travelTime")

	if ptr.Pointer() != nil {
		return int(C.QGeoRoute_TravelTime(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRoute) DestroyQGeoRoute() {
	defer qt.Recovering("QGeoRoute::~QGeoRoute")

	if ptr.Pointer() != nil {
		C.QGeoRoute_DestroyQGeoRoute(ptr.Pointer())
	}
}
