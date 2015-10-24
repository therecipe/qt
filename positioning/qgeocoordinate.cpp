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

QtObjectPtr QGeoCoordinate_NewQGeoCoordinate(){
	return new QGeoCoordinate();
}

QtObjectPtr QGeoCoordinate_NewQGeoCoordinate4(QtObjectPtr other){
	return new QGeoCoordinate(*static_cast<QGeoCoordinate*>(other));
}

int QGeoCoordinate_IsValid(QtObjectPtr ptr){
	return static_cast<QGeoCoordinate*>(ptr)->isValid();
}

char* QGeoCoordinate_ToString(QtObjectPtr ptr, int format){
	return static_cast<QGeoCoordinate*>(ptr)->toString(static_cast<QGeoCoordinate::CoordinateFormat>(format)).toUtf8().data();
}

int QGeoCoordinate_Type(QtObjectPtr ptr){
	return static_cast<QGeoCoordinate*>(ptr)->type();
}

void QGeoCoordinate_DestroyQGeoCoordinate(QtObjectPtr ptr){
	static_cast<QGeoCoordinate*>(ptr)->~QGeoCoordinate();
}

