#include "qgeoshape.h"
#include <QUrl>
#include <QModelIndex>
#include <QGeoCoordinate>
#include <QString>
#include <QVariant>
#include <QGeoShape>
#include "_cgo_export.h"

class MyQGeoShape: public QGeoShape {
public:
};

void* QGeoShape_NewQGeoShape(){
	return new QGeoShape();
}

void* QGeoShape_NewQGeoShape2(void* other){
	return new QGeoShape(*static_cast<QGeoShape*>(other));
}

int QGeoShape_Contains(void* ptr, void* coordinate){
	return static_cast<QGeoShape*>(ptr)->contains(*static_cast<QGeoCoordinate*>(coordinate));
}

void QGeoShape_ExtendShape(void* ptr, void* coordinate){
	static_cast<QGeoShape*>(ptr)->extendShape(*static_cast<QGeoCoordinate*>(coordinate));
}

int QGeoShape_IsEmpty(void* ptr){
	return static_cast<QGeoShape*>(ptr)->isEmpty();
}

int QGeoShape_IsValid(void* ptr){
	return static_cast<QGeoShape*>(ptr)->isValid();
}

char* QGeoShape_ToString(void* ptr){
	return static_cast<QGeoShape*>(ptr)->toString().toUtf8().data();
}

int QGeoShape_Type(void* ptr){
	return static_cast<QGeoShape*>(ptr)->type();
}

void QGeoShape_DestroyQGeoShape(void* ptr){
	static_cast<QGeoShape*>(ptr)->~QGeoShape();
}

