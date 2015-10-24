#include "qgeorouterequest.h"
#include <QUrl>
#include <QModelIndex>
#include <QGeoCoordinate>
#include <QGeoRoute>
#include <QString>
#include <QVariant>
#include <QGeoRouteRequest>
#include "_cgo_export.h"

class MyQGeoRouteRequest: public QGeoRouteRequest {
public:
};

QtObjectPtr QGeoRouteRequest_NewQGeoRouteRequest2(QtObjectPtr origin, QtObjectPtr destination){
	return new QGeoRouteRequest(*static_cast<QGeoCoordinate*>(origin), *static_cast<QGeoCoordinate*>(destination));
}

QtObjectPtr QGeoRouteRequest_NewQGeoRouteRequest3(QtObjectPtr other){
	return new QGeoRouteRequest(*static_cast<QGeoRouteRequest*>(other));
}

int QGeoRouteRequest_FeatureWeight(QtObjectPtr ptr, int featureType){
	return static_cast<QGeoRouteRequest*>(ptr)->featureWeight(static_cast<QGeoRouteRequest::FeatureType>(featureType));
}

int QGeoRouteRequest_ManeuverDetail(QtObjectPtr ptr){
	return static_cast<QGeoRouteRequest*>(ptr)->maneuverDetail();
}

int QGeoRouteRequest_NumberAlternativeRoutes(QtObjectPtr ptr){
	return static_cast<QGeoRouteRequest*>(ptr)->numberAlternativeRoutes();
}

int QGeoRouteRequest_RouteOptimization(QtObjectPtr ptr){
	return static_cast<QGeoRouteRequest*>(ptr)->routeOptimization();
}

int QGeoRouteRequest_SegmentDetail(QtObjectPtr ptr){
	return static_cast<QGeoRouteRequest*>(ptr)->segmentDetail();
}

void QGeoRouteRequest_SetFeatureWeight(QtObjectPtr ptr, int featureType, int featureWeight){
	static_cast<QGeoRouteRequest*>(ptr)->setFeatureWeight(static_cast<QGeoRouteRequest::FeatureType>(featureType), static_cast<QGeoRouteRequest::FeatureWeight>(featureWeight));
}

void QGeoRouteRequest_SetManeuverDetail(QtObjectPtr ptr, int maneuverDetail){
	static_cast<QGeoRouteRequest*>(ptr)->setManeuverDetail(static_cast<QGeoRouteRequest::ManeuverDetail>(maneuverDetail));
}

void QGeoRouteRequest_SetNumberAlternativeRoutes(QtObjectPtr ptr, int alternatives){
	static_cast<QGeoRouteRequest*>(ptr)->setNumberAlternativeRoutes(alternatives);
}

void QGeoRouteRequest_SetRouteOptimization(QtObjectPtr ptr, int optimization){
	static_cast<QGeoRouteRequest*>(ptr)->setRouteOptimization(static_cast<QGeoRouteRequest::RouteOptimization>(optimization));
}

void QGeoRouteRequest_SetSegmentDetail(QtObjectPtr ptr, int segmentDetail){
	static_cast<QGeoRouteRequest*>(ptr)->setSegmentDetail(static_cast<QGeoRouteRequest::SegmentDetail>(segmentDetail));
}

void QGeoRouteRequest_SetTravelModes(QtObjectPtr ptr, int travelModes){
	static_cast<QGeoRouteRequest*>(ptr)->setTravelModes(static_cast<QGeoRouteRequest::TravelMode>(travelModes));
}

int QGeoRouteRequest_TravelModes(QtObjectPtr ptr){
	return static_cast<QGeoRouteRequest*>(ptr)->travelModes();
}

void QGeoRouteRequest_DestroyQGeoRouteRequest(QtObjectPtr ptr){
	static_cast<QGeoRouteRequest*>(ptr)->~QGeoRouteRequest();
}

