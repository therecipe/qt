#include "qgeoroutesegment.h"
#include <QGeoManeuver>
#include <QGeoRoute>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGeoRouteSegment>
#include "_cgo_export.h"

class MyQGeoRouteSegment: public QGeoRouteSegment {
public:
};

QtObjectPtr QGeoRouteSegment_NewQGeoRouteSegment(){
	return new QGeoRouteSegment();
}

QtObjectPtr QGeoRouteSegment_NewQGeoRouteSegment2(QtObjectPtr other){
	return new QGeoRouteSegment(*static_cast<QGeoRouteSegment*>(other));
}

int QGeoRouteSegment_IsValid(QtObjectPtr ptr){
	return static_cast<QGeoRouteSegment*>(ptr)->isValid();
}

void QGeoRouteSegment_SetManeuver(QtObjectPtr ptr, QtObjectPtr maneuver){
	static_cast<QGeoRouteSegment*>(ptr)->setManeuver(*static_cast<QGeoManeuver*>(maneuver));
}

void QGeoRouteSegment_SetNextRouteSegment(QtObjectPtr ptr, QtObjectPtr routeSegment){
	static_cast<QGeoRouteSegment*>(ptr)->setNextRouteSegment(*static_cast<QGeoRouteSegment*>(routeSegment));
}

void QGeoRouteSegment_SetTravelTime(QtObjectPtr ptr, int secs){
	static_cast<QGeoRouteSegment*>(ptr)->setTravelTime(secs);
}

int QGeoRouteSegment_TravelTime(QtObjectPtr ptr){
	return static_cast<QGeoRouteSegment*>(ptr)->travelTime();
}

void QGeoRouteSegment_DestroyQGeoRouteSegment(QtObjectPtr ptr){
	static_cast<QGeoRouteSegment*>(ptr)->~QGeoRouteSegment();
}

