#include "qgeorectangle.h"
#include <QUrl>
#include <QModelIndex>
#include <QGeoShape>
#include <QGeoCoordinate>
#include <QString>
#include <QVariant>
#include <QGeoRectangle>
#include "_cgo_export.h"

class MyQGeoRectangle: public QGeoRectangle {
public:
};

void* QGeoRectangle_NewQGeoRectangle(){
	return new QGeoRectangle();
}

void* QGeoRectangle_NewQGeoRectangle3(void* topLeft, void* bottomRight){
	return new QGeoRectangle(*static_cast<QGeoCoordinate*>(topLeft), *static_cast<QGeoCoordinate*>(bottomRight));
}

void* QGeoRectangle_NewQGeoRectangle5(void* other){
	return new QGeoRectangle(*static_cast<QGeoRectangle*>(other));
}

void* QGeoRectangle_NewQGeoRectangle6(void* other){
	return new QGeoRectangle(*static_cast<QGeoShape*>(other));
}

int QGeoRectangle_Contains(void* ptr, void* rectangle){
	return static_cast<QGeoRectangle*>(ptr)->contains(*static_cast<QGeoRectangle*>(rectangle));
}

int QGeoRectangle_Intersects(void* ptr, void* rectangle){
	return static_cast<QGeoRectangle*>(ptr)->intersects(*static_cast<QGeoRectangle*>(rectangle));
}

void QGeoRectangle_SetBottomLeft(void* ptr, void* bottomLeft){
	static_cast<QGeoRectangle*>(ptr)->setBottomLeft(*static_cast<QGeoCoordinate*>(bottomLeft));
}

void QGeoRectangle_SetBottomRight(void* ptr, void* bottomRight){
	static_cast<QGeoRectangle*>(ptr)->setBottomRight(*static_cast<QGeoCoordinate*>(bottomRight));
}

void QGeoRectangle_SetCenter(void* ptr, void* center){
	static_cast<QGeoRectangle*>(ptr)->setCenter(*static_cast<QGeoCoordinate*>(center));
}

void QGeoRectangle_SetTopLeft(void* ptr, void* topLeft){
	static_cast<QGeoRectangle*>(ptr)->setTopLeft(*static_cast<QGeoCoordinate*>(topLeft));
}

void QGeoRectangle_SetTopRight(void* ptr, void* topRight){
	static_cast<QGeoRectangle*>(ptr)->setTopRight(*static_cast<QGeoCoordinate*>(topRight));
}

char* QGeoRectangle_ToString(void* ptr){
	return static_cast<QGeoRectangle*>(ptr)->toString().toUtf8().data();
}

void QGeoRectangle_DestroyQGeoRectangle(void* ptr){
	static_cast<QGeoRectangle*>(ptr)->~QGeoRectangle();
}

