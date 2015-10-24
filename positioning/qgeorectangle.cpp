#include "qgeorectangle.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGeoShape>
#include <QGeoCoordinate>
#include <QString>
#include <QGeoRectangle>
#include "_cgo_export.h"

class MyQGeoRectangle: public QGeoRectangle {
public:
};

QtObjectPtr QGeoRectangle_NewQGeoRectangle(){
	return new QGeoRectangle();
}

QtObjectPtr QGeoRectangle_NewQGeoRectangle3(QtObjectPtr topLeft, QtObjectPtr bottomRight){
	return new QGeoRectangle(*static_cast<QGeoCoordinate*>(topLeft), *static_cast<QGeoCoordinate*>(bottomRight));
}

QtObjectPtr QGeoRectangle_NewQGeoRectangle5(QtObjectPtr other){
	return new QGeoRectangle(*static_cast<QGeoRectangle*>(other));
}

QtObjectPtr QGeoRectangle_NewQGeoRectangle6(QtObjectPtr other){
	return new QGeoRectangle(*static_cast<QGeoShape*>(other));
}

int QGeoRectangle_Contains(QtObjectPtr ptr, QtObjectPtr rectangle){
	return static_cast<QGeoRectangle*>(ptr)->contains(*static_cast<QGeoRectangle*>(rectangle));
}

int QGeoRectangle_Intersects(QtObjectPtr ptr, QtObjectPtr rectangle){
	return static_cast<QGeoRectangle*>(ptr)->intersects(*static_cast<QGeoRectangle*>(rectangle));
}

void QGeoRectangle_SetBottomLeft(QtObjectPtr ptr, QtObjectPtr bottomLeft){
	static_cast<QGeoRectangle*>(ptr)->setBottomLeft(*static_cast<QGeoCoordinate*>(bottomLeft));
}

void QGeoRectangle_SetBottomRight(QtObjectPtr ptr, QtObjectPtr bottomRight){
	static_cast<QGeoRectangle*>(ptr)->setBottomRight(*static_cast<QGeoCoordinate*>(bottomRight));
}

void QGeoRectangle_SetCenter(QtObjectPtr ptr, QtObjectPtr center){
	static_cast<QGeoRectangle*>(ptr)->setCenter(*static_cast<QGeoCoordinate*>(center));
}

void QGeoRectangle_SetTopLeft(QtObjectPtr ptr, QtObjectPtr topLeft){
	static_cast<QGeoRectangle*>(ptr)->setTopLeft(*static_cast<QGeoCoordinate*>(topLeft));
}

void QGeoRectangle_SetTopRight(QtObjectPtr ptr, QtObjectPtr topRight){
	static_cast<QGeoRectangle*>(ptr)->setTopRight(*static_cast<QGeoCoordinate*>(topRight));
}

char* QGeoRectangle_ToString(QtObjectPtr ptr){
	return static_cast<QGeoRectangle*>(ptr)->toString().toUtf8().data();
}

void QGeoRectangle_DestroyQGeoRectangle(QtObjectPtr ptr){
	static_cast<QGeoRectangle*>(ptr)->~QGeoRectangle();
}

