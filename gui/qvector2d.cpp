#include "qvector2d.h"
#include <QUrl>
#include <QModelIndex>
#include <QVector4D>
#include <QPointF>
#include <QString>
#include <QVariant>
#include <QVector>
#include <QPoint>
#include <QVector3D>
#include <QVector2D>
#include "_cgo_export.h"

class MyQVector2D: public QVector2D {
public:
};

void* QVector2D_NewQVector2D(){
	return new QVector2D();
}

void* QVector2D_NewQVector2D4(void* point){
	return new QVector2D(*static_cast<QPoint*>(point));
}

void* QVector2D_NewQVector2D5(void* point){
	return new QVector2D(*static_cast<QPointF*>(point));
}

void* QVector2D_NewQVector2D6(void* vector){
	return new QVector2D(*static_cast<QVector3D*>(vector));
}

void* QVector2D_NewQVector2D7(void* vector){
	return new QVector2D(*static_cast<QVector4D*>(vector));
}

int QVector2D_IsNull(void* ptr){
	return static_cast<QVector2D*>(ptr)->isNull();
}

void QVector2D_Normalize(void* ptr){
	static_cast<QVector2D*>(ptr)->normalize();
}

