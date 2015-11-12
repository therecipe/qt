#include "qgeoroute.h"
#include <QGeoRectangle>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGeoRouteRequest>
#include <QGeoRouteSegment>
#include <QGeoRoute>
#include "_cgo_export.h"

class MyQGeoRoute: public QGeoRoute {
public:
};

void* QGeoRoute_NewQGeoRoute(){
	return new QGeoRoute();
}

void* QGeoRoute_NewQGeoRoute2(void* other){
	return new QGeoRoute(*static_cast<QGeoRoute*>(other));
}

double QGeoRoute_Distance(void* ptr){
	return static_cast<double>(static_cast<QGeoRoute*>(ptr)->distance());
}

char* QGeoRoute_RouteId(void* ptr){
	return static_cast<QGeoRoute*>(ptr)->routeId().toUtf8().data();
}

void QGeoRoute_SetBounds(void* ptr, void* bounds){
	static_cast<QGeoRoute*>(ptr)->setBounds(*static_cast<QGeoRectangle*>(bounds));
}

void QGeoRoute_SetDistance(void* ptr, double distance){
	static_cast<QGeoRoute*>(ptr)->setDistance(static_cast<qreal>(distance));
}

void QGeoRoute_SetFirstRouteSegment(void* ptr, void* routeSegment){
	static_cast<QGeoRoute*>(ptr)->setFirstRouteSegment(*static_cast<QGeoRouteSegment*>(routeSegment));
}

void QGeoRoute_SetRequest(void* ptr, void* request){
	static_cast<QGeoRoute*>(ptr)->setRequest(*static_cast<QGeoRouteRequest*>(request));
}

void QGeoRoute_SetRouteId(void* ptr, char* id){
	static_cast<QGeoRoute*>(ptr)->setRouteId(QString(id));
}

void QGeoRoute_SetTravelMode(void* ptr, int mode){
	static_cast<QGeoRoute*>(ptr)->setTravelMode(static_cast<QGeoRouteRequest::TravelMode>(mode));
}

void QGeoRoute_SetTravelTime(void* ptr, int secs){
	static_cast<QGeoRoute*>(ptr)->setTravelTime(secs);
}

int QGeoRoute_TravelMode(void* ptr){
	return static_cast<QGeoRoute*>(ptr)->travelMode();
}

int QGeoRoute_TravelTime(void* ptr){
	return static_cast<QGeoRoute*>(ptr)->travelTime();
}

void QGeoRoute_DestroyQGeoRoute(void* ptr){
	static_cast<QGeoRoute*>(ptr)->~QGeoRoute();
}

