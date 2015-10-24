#include "qpointf.h"
#include <QUrl>
#include <QModelIndex>
#include <QPoint>
#include <QString>
#include <QVariant>
#include <QPointF>
#include "_cgo_export.h"

class MyQPointF: public QPointF {
public:
};

QtObjectPtr QPointF_NewQPointF(){
	return new QPointF();
}

QtObjectPtr QPointF_NewQPointF2(QtObjectPtr point){
	return new QPointF(*static_cast<QPoint*>(point));
}

int QPointF_IsNull(QtObjectPtr ptr){
	return static_cast<QPointF*>(ptr)->isNull();
}

