#include "qgeoroutesegment.h"
#include <QUrl>
#include <QModelIndex>
#include <QGeoManeuver>
#include <QGeoRoute>
#include <QString>
#include <QVariant>
#include <QGeoRouteSegment>
#include "_cgo_export.h"

class MyQGeoRouteSegment: public QGeoRouteSegment {
public:
};

void* QGeoRouteSegment_NewQGeoRouteSegment(){
	return new QGeoRouteSegment();
}

void* QGeoRouteSegment_NewQGeoRouteSegment2(void* other){
	return new QGeoRouteSegment(*static_cast<QGeoRouteSegment*>(other));
}

double QGeoRouteSegment_Distance(void* ptr){
	return static_cast<double>(static_cast<QGeoRouteSegment*>(ptr)->distance());
}

int QGeoRouteSegment_IsValid(void* ptr){
	return static_cast<QGeoRouteSegment*>(ptr)->isValid();
}

void QGeoRouteSegment_SetDistance(void* ptr, double distance){
	static_cast<QGeoRouteSegment*>(ptr)->setDistance(static_cast<qreal>(distance));
}

void QGeoRouteSegment_SetManeuver(void* ptr, void* maneuver){
	static_cast<QGeoRouteSegment*>(ptr)->setManeuver(*static_cast<QGeoManeuver*>(maneuver));
}

void QGeoRouteSegment_SetNextRouteSegment(void* ptr, void* routeSegment){
	static_cast<QGeoRouteSegment*>(ptr)->setNextRouteSegment(*static_cast<QGeoRouteSegment*>(routeSegment));
}

void QGeoRouteSegment_SetTravelTime(void* ptr, int secs){
	static_cast<QGeoRouteSegment*>(ptr)->setTravelTime(secs);
}

int QGeoRouteSegment_TravelTime(void* ptr){
	return static_cast<QGeoRouteSegment*>(ptr)->travelTime();
}

void QGeoRouteSegment_DestroyQGeoRouteSegment(void* ptr){
	static_cast<QGeoRouteSegment*>(ptr)->~QGeoRouteSegment();
}

