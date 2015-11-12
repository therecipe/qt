#include "qgeorouterequest.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGeoCoordinate>
#include <QGeoRoute>
#include <QGeoRouteRequest>
#include "_cgo_export.h"

class MyQGeoRouteRequest: public QGeoRouteRequest {
public:
};

void* QGeoRouteRequest_NewQGeoRouteRequest2(void* origin, void* destination){
	return new QGeoRouteRequest(*static_cast<QGeoCoordinate*>(origin), *static_cast<QGeoCoordinate*>(destination));
}

void* QGeoRouteRequest_NewQGeoRouteRequest3(void* other){
	return new QGeoRouteRequest(*static_cast<QGeoRouteRequest*>(other));
}

int QGeoRouteRequest_FeatureWeight(void* ptr, int featureType){
	return static_cast<QGeoRouteRequest*>(ptr)->featureWeight(static_cast<QGeoRouteRequest::FeatureType>(featureType));
}

int QGeoRouteRequest_ManeuverDetail(void* ptr){
	return static_cast<QGeoRouteRequest*>(ptr)->maneuverDetail();
}

int QGeoRouteRequest_NumberAlternativeRoutes(void* ptr){
	return static_cast<QGeoRouteRequest*>(ptr)->numberAlternativeRoutes();
}

int QGeoRouteRequest_RouteOptimization(void* ptr){
	return static_cast<QGeoRouteRequest*>(ptr)->routeOptimization();
}

int QGeoRouteRequest_SegmentDetail(void* ptr){
	return static_cast<QGeoRouteRequest*>(ptr)->segmentDetail();
}

void QGeoRouteRequest_SetFeatureWeight(void* ptr, int featureType, int featureWeight){
	static_cast<QGeoRouteRequest*>(ptr)->setFeatureWeight(static_cast<QGeoRouteRequest::FeatureType>(featureType), static_cast<QGeoRouteRequest::FeatureWeight>(featureWeight));
}

void QGeoRouteRequest_SetManeuverDetail(void* ptr, int maneuverDetail){
	static_cast<QGeoRouteRequest*>(ptr)->setManeuverDetail(static_cast<QGeoRouteRequest::ManeuverDetail>(maneuverDetail));
}

void QGeoRouteRequest_SetNumberAlternativeRoutes(void* ptr, int alternatives){
	static_cast<QGeoRouteRequest*>(ptr)->setNumberAlternativeRoutes(alternatives);
}

void QGeoRouteRequest_SetRouteOptimization(void* ptr, int optimization){
	static_cast<QGeoRouteRequest*>(ptr)->setRouteOptimization(static_cast<QGeoRouteRequest::RouteOptimization>(optimization));
}

void QGeoRouteRequest_SetSegmentDetail(void* ptr, int segmentDetail){
	static_cast<QGeoRouteRequest*>(ptr)->setSegmentDetail(static_cast<QGeoRouteRequest::SegmentDetail>(segmentDetail));
}

void QGeoRouteRequest_SetTravelModes(void* ptr, int travelModes){
	static_cast<QGeoRouteRequest*>(ptr)->setTravelModes(static_cast<QGeoRouteRequest::TravelMode>(travelModes));
}

int QGeoRouteRequest_TravelModes(void* ptr){
	return static_cast<QGeoRouteRequest*>(ptr)->travelModes();
}

void QGeoRouteRequest_DestroyQGeoRouteRequest(void* ptr){
	static_cast<QGeoRouteRequest*>(ptr)->~QGeoRouteRequest();
}

