#include "qgeocoordinate.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGeoCoordinate>
#include "_cgo_export.h"

class MyQGeoCoordinate: public QGeoCoordinate {
public:
};

void* QGeoCoordinate_NewQGeoCoordinate(){
	return new QGeoCoordinate();
}

void* QGeoCoordinate_NewQGeoCoordinate4(void* other){
	return new QGeoCoordinate(*static_cast<QGeoCoordinate*>(other));
}

double QGeoCoordinate_AzimuthTo(void* ptr, void* other){
	return static_cast<double>(static_cast<QGeoCoordinate*>(ptr)->azimuthTo(*static_cast<QGeoCoordinate*>(other)));
}

double QGeoCoordinate_DistanceTo(void* ptr, void* other){
	return static_cast<double>(static_cast<QGeoCoordinate*>(ptr)->distanceTo(*static_cast<QGeoCoordinate*>(other)));
}

int QGeoCoordinate_IsValid(void* ptr){
	return static_cast<QGeoCoordinate*>(ptr)->isValid();
}

char* QGeoCoordinate_ToString(void* ptr, int format){
	return static_cast<QGeoCoordinate*>(ptr)->toString(static_cast<QGeoCoordinate::CoordinateFormat>(format)).toUtf8().data();
}

int QGeoCoordinate_Type(void* ptr){
	return static_cast<QGeoCoordinate*>(ptr)->type();
}

void QGeoCoordinate_DestroyQGeoCoordinate(void* ptr){
	static_cast<QGeoCoordinate*>(ptr)->~QGeoCoordinate();
}

