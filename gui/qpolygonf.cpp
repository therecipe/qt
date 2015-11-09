#include "qpolygonf.h"
#include <QString>
#include <QPointF>
#include <QPolygon>
#include <QRectF>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QRect>
#include <QPoint>
#include <QPolygonF>
#include "_cgo_export.h"

class MyQPolygonF: public QPolygonF {
public:
};

void* QPolygonF_NewQPolygonF6(void* polygon){
	return new QPolygonF(*static_cast<QPolygon*>(polygon));
}

void* QPolygonF_NewQPolygonF5(void* rectangle){
	return new QPolygonF(*static_cast<QRectF*>(rectangle));
}

int QPolygonF_ContainsPoint(void* ptr, void* point, int fillRule){
	return static_cast<QPolygonF*>(ptr)->containsPoint(*static_cast<QPointF*>(point), static_cast<Qt::FillRule>(fillRule));
}

void* QPolygonF_NewQPolygonF(){
	return new QPolygonF();
}

void* QPolygonF_NewQPolygonF3(void* polygon){
	return new QPolygonF(*static_cast<QPolygonF*>(polygon));
}

void* QPolygonF_NewQPolygonF2(int size){
	return new QPolygonF(size);
}

int QPolygonF_IsClosed(void* ptr){
	return static_cast<QPolygonF*>(ptr)->isClosed();
}

void QPolygonF_Swap(void* ptr, void* other){
	static_cast<QPolygonF*>(ptr)->swap(*static_cast<QPolygonF*>(other));
}

void QPolygonF_Translate(void* ptr, void* offset){
	static_cast<QPolygonF*>(ptr)->translate(*static_cast<QPointF*>(offset));
}

void QPolygonF_Translate2(void* ptr, double dx, double dy){
	static_cast<QPolygonF*>(ptr)->translate(static_cast<qreal>(dx), static_cast<qreal>(dy));
}

void QPolygonF_DestroyQPolygonF(void* ptr){
	static_cast<QPolygonF*>(ptr)->~QPolygonF();
}

