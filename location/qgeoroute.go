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

type QGeoRouteITF interface {
	QGeoRoutePTR() *QGeoRoute
}

func (p *QGeoRoute) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QGeoRoute) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQGeoRoute(ptr QGeoRouteITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoRoutePTR().Pointer()
	}
	return nil
}

func QGeoRouteFromPointer(ptr unsafe.Pointer) *QGeoRoute {
	var n = new(QGeoRoute)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGeoRoute) QGeoRoutePTR() *QGeoRoute {
	return ptr
}

func NewQGeoRoute() *QGeoRoute {
	return QGeoRouteFromPointer(unsafe.Pointer(C.QGeoRoute_NewQGeoRoute()))
}

func NewQGeoRoute2(other QGeoRouteITF) *QGeoRoute {
	return QGeoRouteFromPointer(unsafe.Pointer(C.QGeoRoute_NewQGeoRoute2(C.QtObjectPtr(PointerFromQGeoRoute(other)))))
}

func (ptr *QGeoRoute) RouteId() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoRoute_RouteId(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QGeoRoute) SetBounds(bounds positioning.QGeoRectangleITF) {
	if ptr.Pointer() != nil {
		C.QGeoRoute_SetBounds(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(positioning.PointerFromQGeoRectangle(bounds)))
	}
}

func (ptr *QGeoRoute) SetFirstRouteSegment(routeSegment QGeoRouteSegmentITF) {
	if ptr.Pointer() != nil {
		C.QGeoRoute_SetFirstRouteSegment(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGeoRouteSegment(routeSegment)))
	}
}

func (ptr *QGeoRoute) SetRequest(request QGeoRouteRequestITF) {
	if ptr.Pointer() != nil {
		C.QGeoRoute_SetRequest(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGeoRouteRequest(request)))
	}
}

func (ptr *QGeoRoute) SetRouteId(id string) {
	if ptr.Pointer() != nil {
		C.QGeoRoute_SetRouteId(C.QtObjectPtr(ptr.Pointer()), C.CString(id))
	}
}

func (ptr *QGeoRoute) SetTravelMode(mode QGeoRouteRequest__TravelMode) {
	if ptr.Pointer() != nil {
		C.QGeoRoute_SetTravelMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QGeoRoute) SetTravelTime(secs int) {
	if ptr.Pointer() != nil {
		C.QGeoRoute_SetTravelTime(C.QtObjectPtr(ptr.Pointer()), C.int(secs))
	}
}

func (ptr *QGeoRoute) TravelMode() QGeoRouteRequest__TravelMode {
	if ptr.Pointer() != nil {
		return QGeoRouteRequest__TravelMode(C.QGeoRoute_TravelMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoRoute) TravelTime() int {
	if ptr.Pointer() != nil {
		return int(C.QGeoRoute_TravelTime(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoRoute) DestroyQGeoRoute() {
	if ptr.Pointer() != nil {
		C.QGeoRoute_DestroyQGeoRoute(C.QtObjectPtr(ptr.Pointer()))
	}
}
