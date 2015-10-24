package location

//#include "qgeorouterequest.h"
import "C"
import (
	"github.com/therecipe/qt/positioning"
	"unsafe"
)

type QGeoRouteRequest struct {
	ptr unsafe.Pointer
}

type QGeoRouteRequestITF interface {
	QGeoRouteRequestPTR() *QGeoRouteRequest
}

func (p *QGeoRouteRequest) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QGeoRouteRequest) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQGeoRouteRequest(ptr QGeoRouteRequestITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoRouteRequestPTR().Pointer()
	}
	return nil
}

func QGeoRouteRequestFromPointer(ptr unsafe.Pointer) *QGeoRouteRequest {
	var n = new(QGeoRouteRequest)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGeoRouteRequest) QGeoRouteRequestPTR() *QGeoRouteRequest {
	return ptr
}

//QGeoRouteRequest::FeatureType
type QGeoRouteRequest__FeatureType int

var (
	QGeoRouteRequest__NoFeature            = QGeoRouteRequest__FeatureType(0x00000000)
	QGeoRouteRequest__TollFeature          = QGeoRouteRequest__FeatureType(0x00000001)
	QGeoRouteRequest__HighwayFeature       = QGeoRouteRequest__FeatureType(0x00000002)
	QGeoRouteRequest__PublicTransitFeature = QGeoRouteRequest__FeatureType(0x00000004)
	QGeoRouteRequest__FerryFeature         = QGeoRouteRequest__FeatureType(0x00000008)
	QGeoRouteRequest__TunnelFeature        = QGeoRouteRequest__FeatureType(0x00000010)
	QGeoRouteRequest__DirtRoadFeature      = QGeoRouteRequest__FeatureType(0x00000020)
	QGeoRouteRequest__ParksFeature         = QGeoRouteRequest__FeatureType(0x00000040)
	QGeoRouteRequest__MotorPoolLaneFeature = QGeoRouteRequest__FeatureType(0x00000080)
)

//QGeoRouteRequest::FeatureWeight
type QGeoRouteRequest__FeatureWeight int

var (
	QGeoRouteRequest__NeutralFeatureWeight  = QGeoRouteRequest__FeatureWeight(0x00000000)
	QGeoRouteRequest__PreferFeatureWeight   = QGeoRouteRequest__FeatureWeight(0x00000001)
	QGeoRouteRequest__RequireFeatureWeight  = QGeoRouteRequest__FeatureWeight(0x00000002)
	QGeoRouteRequest__AvoidFeatureWeight    = QGeoRouteRequest__FeatureWeight(0x00000004)
	QGeoRouteRequest__DisallowFeatureWeight = QGeoRouteRequest__FeatureWeight(0x00000008)
)

//QGeoRouteRequest::ManeuverDetail
type QGeoRouteRequest__ManeuverDetail int

var (
	QGeoRouteRequest__NoManeuvers    = QGeoRouteRequest__ManeuverDetail(0x0000)
	QGeoRouteRequest__BasicManeuvers = QGeoRouteRequest__ManeuverDetail(0x0001)
)

//QGeoRouteRequest::RouteOptimization
type QGeoRouteRequest__RouteOptimization int

var (
	QGeoRouteRequest__ShortestRoute     = QGeoRouteRequest__RouteOptimization(0x0001)
	QGeoRouteRequest__FastestRoute      = QGeoRouteRequest__RouteOptimization(0x0002)
	QGeoRouteRequest__MostEconomicRoute = QGeoRouteRequest__RouteOptimization(0x0004)
	QGeoRouteRequest__MostScenicRoute   = QGeoRouteRequest__RouteOptimization(0x0008)
)

//QGeoRouteRequest::SegmentDetail
type QGeoRouteRequest__SegmentDetail int

var (
	QGeoRouteRequest__NoSegmentData    = QGeoRouteRequest__SegmentDetail(0x0000)
	QGeoRouteRequest__BasicSegmentData = QGeoRouteRequest__SegmentDetail(0x0001)
)

//QGeoRouteRequest::TravelMode
type QGeoRouteRequest__TravelMode int

var (
	QGeoRouteRequest__CarTravel           = QGeoRouteRequest__TravelMode(0x0001)
	QGeoRouteRequest__PedestrianTravel    = QGeoRouteRequest__TravelMode(0x0002)
	QGeoRouteRequest__BicycleTravel       = QGeoRouteRequest__TravelMode(0x0004)
	QGeoRouteRequest__PublicTransitTravel = QGeoRouteRequest__TravelMode(0x0008)
	QGeoRouteRequest__TruckTravel         = QGeoRouteRequest__TravelMode(0x0010)
)

func NewQGeoRouteRequest2(origin positioning.QGeoCoordinateITF, destination positioning.QGeoCoordinateITF) *QGeoRouteRequest {
	return QGeoRouteRequestFromPointer(unsafe.Pointer(C.QGeoRouteRequest_NewQGeoRouteRequest2(C.QtObjectPtr(positioning.PointerFromQGeoCoordinate(origin)), C.QtObjectPtr(positioning.PointerFromQGeoCoordinate(destination)))))
}

func NewQGeoRouteRequest3(other QGeoRouteRequestITF) *QGeoRouteRequest {
	return QGeoRouteRequestFromPointer(unsafe.Pointer(C.QGeoRouteRequest_NewQGeoRouteRequest3(C.QtObjectPtr(PointerFromQGeoRouteRequest(other)))))
}

func (ptr *QGeoRouteRequest) FeatureWeight(featureType QGeoRouteRequest__FeatureType) QGeoRouteRequest__FeatureWeight {
	if ptr.Pointer() != nil {
		return QGeoRouteRequest__FeatureWeight(C.QGeoRouteRequest_FeatureWeight(C.QtObjectPtr(ptr.Pointer()), C.int(featureType)))
	}
	return 0
}

func (ptr *QGeoRouteRequest) ManeuverDetail() QGeoRouteRequest__ManeuverDetail {
	if ptr.Pointer() != nil {
		return QGeoRouteRequest__ManeuverDetail(C.QGeoRouteRequest_ManeuverDetail(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoRouteRequest) NumberAlternativeRoutes() int {
	if ptr.Pointer() != nil {
		return int(C.QGeoRouteRequest_NumberAlternativeRoutes(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoRouteRequest) RouteOptimization() QGeoRouteRequest__RouteOptimization {
	if ptr.Pointer() != nil {
		return QGeoRouteRequest__RouteOptimization(C.QGeoRouteRequest_RouteOptimization(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoRouteRequest) SegmentDetail() QGeoRouteRequest__SegmentDetail {
	if ptr.Pointer() != nil {
		return QGeoRouteRequest__SegmentDetail(C.QGeoRouteRequest_SegmentDetail(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoRouteRequest) SetFeatureWeight(featureType QGeoRouteRequest__FeatureType, featureWeight QGeoRouteRequest__FeatureWeight) {
	if ptr.Pointer() != nil {
		C.QGeoRouteRequest_SetFeatureWeight(C.QtObjectPtr(ptr.Pointer()), C.int(featureType), C.int(featureWeight))
	}
}

func (ptr *QGeoRouteRequest) SetManeuverDetail(maneuverDetail QGeoRouteRequest__ManeuverDetail) {
	if ptr.Pointer() != nil {
		C.QGeoRouteRequest_SetManeuverDetail(C.QtObjectPtr(ptr.Pointer()), C.int(maneuverDetail))
	}
}

func (ptr *QGeoRouteRequest) SetNumberAlternativeRoutes(alternatives int) {
	if ptr.Pointer() != nil {
		C.QGeoRouteRequest_SetNumberAlternativeRoutes(C.QtObjectPtr(ptr.Pointer()), C.int(alternatives))
	}
}

func (ptr *QGeoRouteRequest) SetRouteOptimization(optimization QGeoRouteRequest__RouteOptimization) {
	if ptr.Pointer() != nil {
		C.QGeoRouteRequest_SetRouteOptimization(C.QtObjectPtr(ptr.Pointer()), C.int(optimization))
	}
}

func (ptr *QGeoRouteRequest) SetSegmentDetail(segmentDetail QGeoRouteRequest__SegmentDetail) {
	if ptr.Pointer() != nil {
		C.QGeoRouteRequest_SetSegmentDetail(C.QtObjectPtr(ptr.Pointer()), C.int(segmentDetail))
	}
}

func (ptr *QGeoRouteRequest) SetTravelModes(travelModes QGeoRouteRequest__TravelMode) {
	if ptr.Pointer() != nil {
		C.QGeoRouteRequest_SetTravelModes(C.QtObjectPtr(ptr.Pointer()), C.int(travelModes))
	}
}

func (ptr *QGeoRouteRequest) TravelModes() QGeoRouteRequest__TravelMode {
	if ptr.Pointer() != nil {
		return QGeoRouteRequest__TravelMode(C.QGeoRouteRequest_TravelModes(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoRouteRequest) DestroyQGeoRouteRequest() {
	if ptr.Pointer() != nil {
		C.QGeoRouteRequest_DestroyQGeoRouteRequest(C.QtObjectPtr(ptr.Pointer()))
	}
}
