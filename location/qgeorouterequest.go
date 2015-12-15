package location

//#include "location.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/positioning"
	"unsafe"
)

type QGeoRouteRequest struct {
	ptr unsafe.Pointer
}

type QGeoRouteRequest_ITF interface {
	QGeoRouteRequest_PTR() *QGeoRouteRequest
}

func (p *QGeoRouteRequest) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QGeoRouteRequest) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQGeoRouteRequest(ptr QGeoRouteRequest_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoRouteRequest_PTR().Pointer()
	}
	return nil
}

func NewQGeoRouteRequestFromPointer(ptr unsafe.Pointer) *QGeoRouteRequest {
	var n = new(QGeoRouteRequest)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGeoRouteRequest) QGeoRouteRequest_PTR() *QGeoRouteRequest {
	return ptr
}

//QGeoRouteRequest::FeatureType
type QGeoRouteRequest__FeatureType int64

const (
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
type QGeoRouteRequest__FeatureWeight int64

const (
	QGeoRouteRequest__NeutralFeatureWeight  = QGeoRouteRequest__FeatureWeight(0x00000000)
	QGeoRouteRequest__PreferFeatureWeight   = QGeoRouteRequest__FeatureWeight(0x00000001)
	QGeoRouteRequest__RequireFeatureWeight  = QGeoRouteRequest__FeatureWeight(0x00000002)
	QGeoRouteRequest__AvoidFeatureWeight    = QGeoRouteRequest__FeatureWeight(0x00000004)
	QGeoRouteRequest__DisallowFeatureWeight = QGeoRouteRequest__FeatureWeight(0x00000008)
)

//QGeoRouteRequest::ManeuverDetail
type QGeoRouteRequest__ManeuverDetail int64

const (
	QGeoRouteRequest__NoManeuvers    = QGeoRouteRequest__ManeuverDetail(0x0000)
	QGeoRouteRequest__BasicManeuvers = QGeoRouteRequest__ManeuverDetail(0x0001)
)

//QGeoRouteRequest::RouteOptimization
type QGeoRouteRequest__RouteOptimization int64

const (
	QGeoRouteRequest__ShortestRoute     = QGeoRouteRequest__RouteOptimization(0x0001)
	QGeoRouteRequest__FastestRoute      = QGeoRouteRequest__RouteOptimization(0x0002)
	QGeoRouteRequest__MostEconomicRoute = QGeoRouteRequest__RouteOptimization(0x0004)
	QGeoRouteRequest__MostScenicRoute   = QGeoRouteRequest__RouteOptimization(0x0008)
)

//QGeoRouteRequest::SegmentDetail
type QGeoRouteRequest__SegmentDetail int64

const (
	QGeoRouteRequest__NoSegmentData    = QGeoRouteRequest__SegmentDetail(0x0000)
	QGeoRouteRequest__BasicSegmentData = QGeoRouteRequest__SegmentDetail(0x0001)
)

//QGeoRouteRequest::TravelMode
type QGeoRouteRequest__TravelMode int64

const (
	QGeoRouteRequest__CarTravel           = QGeoRouteRequest__TravelMode(0x0001)
	QGeoRouteRequest__PedestrianTravel    = QGeoRouteRequest__TravelMode(0x0002)
	QGeoRouteRequest__BicycleTravel       = QGeoRouteRequest__TravelMode(0x0004)
	QGeoRouteRequest__PublicTransitTravel = QGeoRouteRequest__TravelMode(0x0008)
	QGeoRouteRequest__TruckTravel         = QGeoRouteRequest__TravelMode(0x0010)
)

func NewQGeoRouteRequest2(origin positioning.QGeoCoordinate_ITF, destination positioning.QGeoCoordinate_ITF) *QGeoRouteRequest {
	defer qt.Recovering("QGeoRouteRequest::QGeoRouteRequest")

	return NewQGeoRouteRequestFromPointer(C.QGeoRouteRequest_NewQGeoRouteRequest2(positioning.PointerFromQGeoCoordinate(origin), positioning.PointerFromQGeoCoordinate(destination)))
}

func NewQGeoRouteRequest3(other QGeoRouteRequest_ITF) *QGeoRouteRequest {
	defer qt.Recovering("QGeoRouteRequest::QGeoRouteRequest")

	return NewQGeoRouteRequestFromPointer(C.QGeoRouteRequest_NewQGeoRouteRequest3(PointerFromQGeoRouteRequest(other)))
}

func (ptr *QGeoRouteRequest) FeatureWeight(featureType QGeoRouteRequest__FeatureType) QGeoRouteRequest__FeatureWeight {
	defer qt.Recovering("QGeoRouteRequest::featureWeight")

	if ptr.Pointer() != nil {
		return QGeoRouteRequest__FeatureWeight(C.QGeoRouteRequest_FeatureWeight(ptr.Pointer(), C.int(featureType)))
	}
	return 0
}

func (ptr *QGeoRouteRequest) ManeuverDetail() QGeoRouteRequest__ManeuverDetail {
	defer qt.Recovering("QGeoRouteRequest::maneuverDetail")

	if ptr.Pointer() != nil {
		return QGeoRouteRequest__ManeuverDetail(C.QGeoRouteRequest_ManeuverDetail(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRouteRequest) NumberAlternativeRoutes() int {
	defer qt.Recovering("QGeoRouteRequest::numberAlternativeRoutes")

	if ptr.Pointer() != nil {
		return int(C.QGeoRouteRequest_NumberAlternativeRoutes(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRouteRequest) RouteOptimization() QGeoRouteRequest__RouteOptimization {
	defer qt.Recovering("QGeoRouteRequest::routeOptimization")

	if ptr.Pointer() != nil {
		return QGeoRouteRequest__RouteOptimization(C.QGeoRouteRequest_RouteOptimization(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRouteRequest) SegmentDetail() QGeoRouteRequest__SegmentDetail {
	defer qt.Recovering("QGeoRouteRequest::segmentDetail")

	if ptr.Pointer() != nil {
		return QGeoRouteRequest__SegmentDetail(C.QGeoRouteRequest_SegmentDetail(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRouteRequest) SetFeatureWeight(featureType QGeoRouteRequest__FeatureType, featureWeight QGeoRouteRequest__FeatureWeight) {
	defer qt.Recovering("QGeoRouteRequest::setFeatureWeight")

	if ptr.Pointer() != nil {
		C.QGeoRouteRequest_SetFeatureWeight(ptr.Pointer(), C.int(featureType), C.int(featureWeight))
	}
}

func (ptr *QGeoRouteRequest) SetManeuverDetail(maneuverDetail QGeoRouteRequest__ManeuverDetail) {
	defer qt.Recovering("QGeoRouteRequest::setManeuverDetail")

	if ptr.Pointer() != nil {
		C.QGeoRouteRequest_SetManeuverDetail(ptr.Pointer(), C.int(maneuverDetail))
	}
}

func (ptr *QGeoRouteRequest) SetNumberAlternativeRoutes(alternatives int) {
	defer qt.Recovering("QGeoRouteRequest::setNumberAlternativeRoutes")

	if ptr.Pointer() != nil {
		C.QGeoRouteRequest_SetNumberAlternativeRoutes(ptr.Pointer(), C.int(alternatives))
	}
}

func (ptr *QGeoRouteRequest) SetRouteOptimization(optimization QGeoRouteRequest__RouteOptimization) {
	defer qt.Recovering("QGeoRouteRequest::setRouteOptimization")

	if ptr.Pointer() != nil {
		C.QGeoRouteRequest_SetRouteOptimization(ptr.Pointer(), C.int(optimization))
	}
}

func (ptr *QGeoRouteRequest) SetSegmentDetail(segmentDetail QGeoRouteRequest__SegmentDetail) {
	defer qt.Recovering("QGeoRouteRequest::setSegmentDetail")

	if ptr.Pointer() != nil {
		C.QGeoRouteRequest_SetSegmentDetail(ptr.Pointer(), C.int(segmentDetail))
	}
}

func (ptr *QGeoRouteRequest) SetTravelModes(travelModes QGeoRouteRequest__TravelMode) {
	defer qt.Recovering("QGeoRouteRequest::setTravelModes")

	if ptr.Pointer() != nil {
		C.QGeoRouteRequest_SetTravelModes(ptr.Pointer(), C.int(travelModes))
	}
}

func (ptr *QGeoRouteRequest) TravelModes() QGeoRouteRequest__TravelMode {
	defer qt.Recovering("QGeoRouteRequest::travelModes")

	if ptr.Pointer() != nil {
		return QGeoRouteRequest__TravelMode(C.QGeoRouteRequest_TravelModes(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRouteRequest) DestroyQGeoRouteRequest() {
	defer qt.Recovering("QGeoRouteRequest::~QGeoRouteRequest")

	if ptr.Pointer() != nil {
		C.QGeoRouteRequest_DestroyQGeoRouteRequest(ptr.Pointer())
	}
}
