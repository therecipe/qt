#include "qgeoshape.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGeoCoordinate>
#include <QGeoShape>
#include "_cgo_export.h"

class MyQGeoShape: public QGeoShape {
public:
};

QtObjectPtr QGeoShape_NewQGeoShape(){
	return new QGeoShape();
}

QtObjectPtr QGeoShape_NewQGeoShape2(QtObjectPtr other){
	return new QGeoShape(*static_cast<QGeoShape*>(other));
}

int QGeoShape_Contains(QtObjectPtr ptr, QtObjectPtr coordinate){
	return static_cast<QGeoShape*>(ptr)->contains(*static_cast<QGeoCoordinate*>(coordinate));
}

void QGeoShape_ExtendShape(QtObjectPtr ptr, QtObjectPtr coordinate){
	static_cast<QGeoShape*>(ptr)->extendShape(*static_cast<QGeoCoordinate*>(coordinate));
}

int QGeoShape_IsEmpty(QtObjectPtr ptr){
	return static_cast<QGeoShape*>(ptr)->isEmpty();
}

int QGeoShape_IsValid(QtObjectPtr ptr){
	return static_cast<QGeoShape*>(ptr)->isValid();
}

char* QGeoShape_ToString(QtObjectPtr ptr){
	return static_cast<QGeoShape*>(ptr)->toString().toUtf8().data();
}

int QGeoShape_Type(QtObjectPtr ptr){
	return static_cast<QGeoShape*>(ptr)->type();
}

void QGeoShape_DestroyQGeoShape(QtObjectPtr ptr){
	static_cast<QGeoShape*>(ptr)->~QGeoShape();
}

