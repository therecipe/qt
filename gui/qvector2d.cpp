#include "qvector2d.h"
#include <QPointF>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QPoint>
#include <QVector3D>
#include <QModelIndex>
#include <QVector>
#include <QVector4D>
#include <QVector2D>
#include "_cgo_export.h"

class MyQVector2D: public QVector2D {
public:
};

QtObjectPtr QVector2D_NewQVector2D(){
	return new QVector2D();
}

QtObjectPtr QVector2D_NewQVector2D4(QtObjectPtr point){
	return new QVector2D(*static_cast<QPoint*>(point));
}

QtObjectPtr QVector2D_NewQVector2D5(QtObjectPtr point){
	return new QVector2D(*static_cast<QPointF*>(point));
}

QtObjectPtr QVector2D_NewQVector2D6(QtObjectPtr vector){
	return new QVector2D(*static_cast<QVector3D*>(vector));
}

QtObjectPtr QVector2D_NewQVector2D7(QtObjectPtr vector){
	return new QVector2D(*static_cast<QVector4D*>(vector));
}

int QVector2D_IsNull(QtObjectPtr ptr){
	return static_cast<QVector2D*>(ptr)->isNull();
}

void QVector2D_Normalize(QtObjectPtr ptr){
	static_cast<QVector2D*>(ptr)->normalize();
}

