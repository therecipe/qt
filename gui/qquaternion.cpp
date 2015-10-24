#include "qquaternion.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QVector4D>
#include <QVector>
#include <QVector3D>
#include <QString>
#include <QQuaternion>
#include "_cgo_export.h"

class MyQQuaternion: public QQuaternion {
public:
};

QtObjectPtr QQuaternion_NewQQuaternion(){
	return new QQuaternion();
}

QtObjectPtr QQuaternion_NewQQuaternion5(QtObjectPtr vector){
	return new QQuaternion(*static_cast<QVector4D*>(vector));
}

void QQuaternion_GetAxes(QtObjectPtr ptr, QtObjectPtr xAxis, QtObjectPtr yAxis, QtObjectPtr zAxis){
	static_cast<QQuaternion*>(ptr)->getAxes(static_cast<QVector3D*>(xAxis), static_cast<QVector3D*>(yAxis), static_cast<QVector3D*>(zAxis));
}

int QQuaternion_IsIdentity(QtObjectPtr ptr){
	return static_cast<QQuaternion*>(ptr)->isIdentity();
}

int QQuaternion_IsNull(QtObjectPtr ptr){
	return static_cast<QQuaternion*>(ptr)->isNull();
}

void QQuaternion_Normalize(QtObjectPtr ptr){
	static_cast<QQuaternion*>(ptr)->normalize();
}

void QQuaternion_SetVector(QtObjectPtr ptr, QtObjectPtr vector){
	static_cast<QQuaternion*>(ptr)->setVector(*static_cast<QVector3D*>(vector));
}

