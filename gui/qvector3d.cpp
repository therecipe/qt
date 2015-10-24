#include "qvector3d.h"
#include <QVariant>
#include <QModelIndex>
#include <QPoint>
#include <QVector2D>
#include <QString>
#include <QUrl>
#include <QVector>
#include <QVector4D>
#include <QPointF>
#include <QVector3D>
#include "_cgo_export.h"

class MyQVector3D: public QVector3D {
public:
};

QtObjectPtr QVector3D_NewQVector3D(){
	return new QVector3D();
}

QtObjectPtr QVector3D_NewQVector3D4(QtObjectPtr point){
	return new QVector3D(*static_cast<QPoint*>(point));
}

QtObjectPtr QVector3D_NewQVector3D5(QtObjectPtr point){
	return new QVector3D(*static_cast<QPointF*>(point));
}

QtObjectPtr QVector3D_NewQVector3D6(QtObjectPtr vector){
	return new QVector3D(*static_cast<QVector2D*>(vector));
}

QtObjectPtr QVector3D_NewQVector3D8(QtObjectPtr vector){
	return new QVector3D(*static_cast<QVector4D*>(vector));
}

int QVector3D_IsNull(QtObjectPtr ptr){
	return static_cast<QVector3D*>(ptr)->isNull();
}

void QVector3D_Normalize(QtObjectPtr ptr){
	static_cast<QVector3D*>(ptr)->normalize();
}

