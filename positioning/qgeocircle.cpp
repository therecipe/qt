#include "qgeocircle.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGeoCoordinate>
#include <QGeoShape>
#include <QGeoCircle>
#include "_cgo_export.h"

class MyQGeoCircle: public QGeoCircle {
public:
};

void* QGeoCircle_NewQGeoCircle(){
	return new QGeoCircle();
}

void* QGeoCircle_NewQGeoCircle3(void* other){
	return new QGeoCircle(*static_cast<QGeoCircle*>(other));
}

void* QGeoCircle_NewQGeoCircle2(void* center, double radius){
	return new QGeoCircle(*static_cast<QGeoCoordinate*>(center), static_cast<qreal>(radius));
}

void* QGeoCircle_NewQGeoCircle4(void* other){
	return new QGeoCircle(*static_cast<QGeoShape*>(other));
}

double QGeoCircle_Radius(void* ptr){
	return static_cast<double>(static_cast<QGeoCircle*>(ptr)->radius());
}

void QGeoCircle_SetCenter(void* ptr, void* center){
	static_cast<QGeoCircle*>(ptr)->setCenter(*static_cast<QGeoCoordinate*>(center));
}

void QGeoCircle_SetRadius(void* ptr, double radius){
	static_cast<QGeoCircle*>(ptr)->setRadius(static_cast<qreal>(radius));
}

char* QGeoCircle_ToString(void* ptr){
	return static_cast<QGeoCircle*>(ptr)->toString().toUtf8().data();
}

void QGeoCircle_DestroyQGeoCircle(void* ptr){
	static_cast<QGeoCircle*>(ptr)->~QGeoCircle();
}

