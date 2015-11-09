#include "qvector4d.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPointF>
#include <QVector2D>
#include <QString>
#include <QPoint>
#include <QVector3D>
#include <QVector>
#include <QVector4D>
#include "_cgo_export.h"

class MyQVector4D: public QVector4D {
public:
};

void* QVector4D_NewQVector4D(){
	return new QVector4D();
}

void* QVector4D_NewQVector4D4(void* point){
	return new QVector4D(*static_cast<QPoint*>(point));
}

void* QVector4D_NewQVector4D5(void* point){
	return new QVector4D(*static_cast<QPointF*>(point));
}

void* QVector4D_NewQVector4D6(void* vector){
	return new QVector4D(*static_cast<QVector2D*>(vector));
}

void* QVector4D_NewQVector4D8(void* vector){
	return new QVector4D(*static_cast<QVector3D*>(vector));
}

int QVector4D_IsNull(void* ptr){
	return static_cast<QVector4D*>(ptr)->isNull();
}

void QVector4D_Normalize(void* ptr){
	static_cast<QVector4D*>(ptr)->normalize();
}

