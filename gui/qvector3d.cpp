#include "qvector3d.h"
#include <QString>
#include <QModelIndex>
#include <QVector4D>
#include <QPointF>
#include <QVariant>
#include <QUrl>
#include <QPoint>
#include <QVector2D>
#include <QVector>
#include <QVector3D>
#include "_cgo_export.h"

class MyQVector3D: public QVector3D {
public:
};

void* QVector3D_NewQVector3D(){
	return new QVector3D();
}

void* QVector3D_NewQVector3D4(void* point){
	return new QVector3D(*static_cast<QPoint*>(point));
}

void* QVector3D_NewQVector3D5(void* point){
	return new QVector3D(*static_cast<QPointF*>(point));
}

void* QVector3D_NewQVector3D6(void* vector){
	return new QVector3D(*static_cast<QVector2D*>(vector));
}

void* QVector3D_NewQVector3D8(void* vector){
	return new QVector3D(*static_cast<QVector4D*>(vector));
}

int QVector3D_IsNull(void* ptr){
	return static_cast<QVector3D*>(ptr)->isNull();
}

void QVector3D_Normalize(void* ptr){
	static_cast<QVector3D*>(ptr)->normalize();
}

