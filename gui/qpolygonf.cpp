#include "qpolygonf.h"
#include <QRectF>
#include <QPoint>
#include <QString>
#include <QUrl>
#include <QModelIndex>
#include <QRect>
#include <QVariant>
#include <QPolygon>
#include <QPointF>
#include <QPolygonF>
#include "_cgo_export.h"

class MyQPolygonF: public QPolygonF {
public:
};

QtObjectPtr QPolygonF_NewQPolygonF6(QtObjectPtr polygon){
	return new QPolygonF(*static_cast<QPolygon*>(polygon));
}

QtObjectPtr QPolygonF_NewQPolygonF5(QtObjectPtr rectangle){
	return new QPolygonF(*static_cast<QRectF*>(rectangle));
}

int QPolygonF_ContainsPoint(QtObjectPtr ptr, QtObjectPtr point, int fillRule){
	return static_cast<QPolygonF*>(ptr)->containsPoint(*static_cast<QPointF*>(point), static_cast<Qt::FillRule>(fillRule));
}

QtObjectPtr QPolygonF_NewQPolygonF(){
	return new QPolygonF();
}

QtObjectPtr QPolygonF_NewQPolygonF3(QtObjectPtr polygon){
	return new QPolygonF(*static_cast<QPolygonF*>(polygon));
}

QtObjectPtr QPolygonF_NewQPolygonF2(int size){
	return new QPolygonF(size);
}

int QPolygonF_IsClosed(QtObjectPtr ptr){
	return static_cast<QPolygonF*>(ptr)->isClosed();
}

void QPolygonF_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QPolygonF*>(ptr)->swap(*static_cast<QPolygonF*>(other));
}

void QPolygonF_Translate(QtObjectPtr ptr, QtObjectPtr offset){
	static_cast<QPolygonF*>(ptr)->translate(*static_cast<QPointF*>(offset));
}

void QPolygonF_DestroyQPolygonF(QtObjectPtr ptr){
	static_cast<QPolygonF*>(ptr)->~QPolygonF();
}

