#include "qquaternion.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QVector>
#include <QVector4D>
#include <QVector3D>
#include <QQuaternion>
#include "_cgo_export.h"

class MyQQuaternion: public QQuaternion {
public:
};

void* QQuaternion_NewQQuaternion(){
	return new QQuaternion();
}

void* QQuaternion_NewQQuaternion5(void* vector){
	return new QQuaternion(*static_cast<QVector4D*>(vector));
}

void QQuaternion_GetAxes(void* ptr, void* xAxis, void* yAxis, void* zAxis){
	static_cast<QQuaternion*>(ptr)->getAxes(static_cast<QVector3D*>(xAxis), static_cast<QVector3D*>(yAxis), static_cast<QVector3D*>(zAxis));
}

int QQuaternion_IsIdentity(void* ptr){
	return static_cast<QQuaternion*>(ptr)->isIdentity();
}

int QQuaternion_IsNull(void* ptr){
	return static_cast<QQuaternion*>(ptr)->isNull();
}

void QQuaternion_Normalize(void* ptr){
	static_cast<QQuaternion*>(ptr)->normalize();
}

void QQuaternion_SetVector(void* ptr, void* vector){
	static_cast<QQuaternion*>(ptr)->setVector(*static_cast<QVector3D*>(vector));
}

