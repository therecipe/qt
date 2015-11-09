#include "qpointf.h"
#include <QModelIndex>
#include <QPoint>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QPointF>
#include "_cgo_export.h"

class MyQPointF: public QPointF {
public:
};

void* QPointF_NewQPointF(){
	return new QPointF();
}

void* QPointF_NewQPointF2(void* point){
	return new QPointF(*static_cast<QPoint*>(point));
}

void* QPointF_NewQPointF3(double xpos, double ypos){
	return new QPointF(static_cast<qreal>(xpos), static_cast<qreal>(ypos));
}

double QPointF_QPointF_DotProduct(void* p1, void* p2){
	return static_cast<double>(QPointF::dotProduct(*static_cast<QPointF*>(p1), *static_cast<QPointF*>(p2)));
}

int QPointF_IsNull(void* ptr){
	return static_cast<QPointF*>(ptr)->isNull();
}

double QPointF_ManhattanLength(void* ptr){
	return static_cast<double>(static_cast<QPointF*>(ptr)->manhattanLength());
}

double QPointF_Rx(void* ptr){
	return static_cast<double>(static_cast<QPointF*>(ptr)->rx());
}

double QPointF_Ry(void* ptr){
	return static_cast<double>(static_cast<QPointF*>(ptr)->ry());
}

void QPointF_SetX(void* ptr, double x){
	static_cast<QPointF*>(ptr)->setX(static_cast<qreal>(x));
}

void QPointF_SetY(void* ptr, double y){
	static_cast<QPointF*>(ptr)->setY(static_cast<qreal>(y));
}

double QPointF_X(void* ptr){
	return static_cast<double>(static_cast<QPointF*>(ptr)->x());
}

double QPointF_Y(void* ptr){
	return static_cast<double>(static_cast<QPointF*>(ptr)->y());
}

