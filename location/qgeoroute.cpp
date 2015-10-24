#include "qgeoroute.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGeoRouteSegment>
#include <QGeoRouteRequest>
#include <QGeoRectangle>
#include <QString>
#include <QGeoRoute>
#include "_cgo_export.h"

class MyQGeoRoute: public QGeoRoute {
public:
};

QtObjectPtr QGeoRoute_NewQGeoRoute(){
	return new QGeoRoute();
}

QtObjectPtr QGeoRoute_NewQGeoRoute2(QtObjectPtr other){
	return new QGeoRoute(*static_cast<QGeoRoute*>(other));
}

char* QGeoRoute_RouteId(QtObjectPtr ptr){
	return static_cast<QGeoRoute*>(ptr)->routeId().toUtf8().data();
}

void QGeoRoute_SetBounds(QtObjectPtr ptr, QtObjectPtr bounds){
	static_cast<QGeoRoute*>(ptr)->setBounds(*static_cast<QGeoRectangle*>(bounds));
}

void QGeoRoute_SetFirstRouteSegment(QtObjectPtr ptr, QtObjectPtr routeSegment){
	static_cast<QGeoRoute*>(ptr)->setFirstRouteSegment(*static_cast<QGeoRouteSegment*>(routeSegment));
}

void QGeoRoute_SetRequest(QtObjectPtr ptr, QtObjectPtr request){
	static_cast<QGeoRoute*>(ptr)->setRequest(*static_cast<QGeoRouteRequest*>(request));
}

void QGeoRoute_SetRouteId(QtObjectPtr ptr, char* id){
	static_cast<QGeoRoute*>(ptr)->setRouteId(QString(id));
}

void QGeoRoute_SetTravelMode(QtObjectPtr ptr, int mode){
	static_cast<QGeoRoute*>(ptr)->setTravelMode(static_cast<QGeoRouteRequest::TravelMode>(mode));
}

void QGeoRoute_SetTravelTime(QtObjectPtr ptr, int secs){
	static_cast<QGeoRoute*>(ptr)->setTravelTime(secs);
}

int QGeoRoute_TravelMode(QtObjectPtr ptr){
	return static_cast<QGeoRoute*>(ptr)->travelMode();
}

int QGeoRoute_TravelTime(QtObjectPtr ptr){
	return static_cast<QGeoRoute*>(ptr)->travelTime();
}

void QGeoRoute_DestroyQGeoRoute(QtObjectPtr ptr){
	static_cast<QGeoRoute*>(ptr)->~QGeoRoute();
}

