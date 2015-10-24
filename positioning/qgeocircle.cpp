#include "qgeocircle.h"
#include <QModelIndex>
#include <QGeoShape>
#include <QGeoCoordinate>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QGeoCircle>
#include "_cgo_export.h"

class MyQGeoCircle: public QGeoCircle {
public:
};

QtObjectPtr QGeoCircle_NewQGeoCircle(){
	return new QGeoCircle();
}

QtObjectPtr QGeoCircle_NewQGeoCircle3(QtObjectPtr other){
	return new QGeoCircle(*static_cast<QGeoCircle*>(other));
}

QtObjectPtr QGeoCircle_NewQGeoCircle4(QtObjectPtr other){
	return new QGeoCircle(*static_cast<QGeoShape*>(other));
}

void QGeoCircle_SetCenter(QtObjectPtr ptr, QtObjectPtr center){
	static_cast<QGeoCircle*>(ptr)->setCenter(*static_cast<QGeoCoordinate*>(center));
}

char* QGeoCircle_ToString(QtObjectPtr ptr){
	return static_cast<QGeoCircle*>(ptr)->toString().toUtf8().data();
}

void QGeoCircle_DestroyQGeoCircle(QtObjectPtr ptr){
	static_cast<QGeoCircle*>(ptr)->~QGeoCircle();
}

