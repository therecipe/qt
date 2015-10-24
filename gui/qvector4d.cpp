#include "qvector4d.h"
#include <QPoint>
#include <QVariant>
#include <QUrl>
#include <QPointF>
#include <QVector>
#include <QVector3D>
#include <QVector2D>
#include <QString>
#include <QModelIndex>
#include <QVector4D>
#include "_cgo_export.h"

class MyQVector4D: public QVector4D {
public:
};

QtObjectPtr QVector4D_NewQVector4D(){
	return new QVector4D();
}

QtObjectPtr QVector4D_NewQVector4D4(QtObjectPtr point){
	return new QVector4D(*static_cast<QPoint*>(point));
}

QtObjectPtr QVector4D_NewQVector4D5(QtObjectPtr point){
	return new QVector4D(*static_cast<QPointF*>(point));
}

QtObjectPtr QVector4D_NewQVector4D6(QtObjectPtr vector){
	return new QVector4D(*static_cast<QVector2D*>(vector));
}

QtObjectPtr QVector4D_NewQVector4D8(QtObjectPtr vector){
	return new QVector4D(*static_cast<QVector3D*>(vector));
}

int QVector4D_IsNull(QtObjectPtr ptr){
	return static_cast<QVector4D*>(ptr)->isNull();
}

void QVector4D_Normalize(QtObjectPtr ptr){
	static_cast<QVector4D*>(ptr)->normalize();
}

