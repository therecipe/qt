#include "qlinef.h"
#include <QLine>
#include <QPoint>
#include <QPointF>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QLineF>
#include "_cgo_export.h"

class MyQLineF: public QLineF {
public:
};

double QLineF_AngleTo(void* ptr, void* line){
	return static_cast<double>(static_cast<QLineF*>(ptr)->angleTo(*static_cast<QLineF*>(line)));
}

int QLineF_Intersect(void* ptr, void* line, void* intersectionPoint){
	return static_cast<QLineF*>(ptr)->intersect(*static_cast<QLineF*>(line), static_cast<QPointF*>(intersectionPoint));
}

void* QLineF_NewQLineF(){
	return new QLineF();
}

void* QLineF_NewQLineF4(void* line){
	return new QLineF(*static_cast<QLine*>(line));
}

void* QLineF_NewQLineF2(void* p1, void* p2){
	return new QLineF(*static_cast<QPointF*>(p1), *static_cast<QPointF*>(p2));
}

void* QLineF_NewQLineF3(double x1, double y1, double x2, double y2){
	return new QLineF(static_cast<qreal>(x1), static_cast<qreal>(y1), static_cast<qreal>(x2), static_cast<qreal>(y2));
}

double QLineF_Angle(void* ptr){
	return static_cast<double>(static_cast<QLineF*>(ptr)->angle());
}

double QLineF_Dx(void* ptr){
	return static_cast<double>(static_cast<QLineF*>(ptr)->dx());
}

double QLineF_Dy(void* ptr){
	return static_cast<double>(static_cast<QLineF*>(ptr)->dy());
}

int QLineF_IsNull(void* ptr){
	return static_cast<QLineF*>(ptr)->isNull();
}

double QLineF_Length(void* ptr){
	return static_cast<double>(static_cast<QLineF*>(ptr)->length());
}

void QLineF_SetAngle(void* ptr, double angle){
	static_cast<QLineF*>(ptr)->setAngle(static_cast<qreal>(angle));
}

void QLineF_SetLength(void* ptr, double length){
	static_cast<QLineF*>(ptr)->setLength(static_cast<qreal>(length));
}

void QLineF_SetLine(void* ptr, double x1, double y1, double x2, double y2){
	static_cast<QLineF*>(ptr)->setLine(static_cast<qreal>(x1), static_cast<qreal>(y1), static_cast<qreal>(x2), static_cast<qreal>(y2));
}

void QLineF_SetP1(void* ptr, void* p1){
	static_cast<QLineF*>(ptr)->setP1(*static_cast<QPointF*>(p1));
}

void QLineF_SetP2(void* ptr, void* p2){
	static_cast<QLineF*>(ptr)->setP2(*static_cast<QPointF*>(p2));
}

void QLineF_SetPoints(void* ptr, void* p1, void* p2){
	static_cast<QLineF*>(ptr)->setPoints(*static_cast<QPointF*>(p1), *static_cast<QPointF*>(p2));
}

void QLineF_Translate(void* ptr, void* offset){
	static_cast<QLineF*>(ptr)->translate(*static_cast<QPointF*>(offset));
}

void QLineF_Translate2(void* ptr, double dx, double dy){
	static_cast<QLineF*>(ptr)->translate(static_cast<qreal>(dx), static_cast<qreal>(dy));
}

double QLineF_X1(void* ptr){
	return static_cast<double>(static_cast<QLineF*>(ptr)->x1());
}

double QLineF_X2(void* ptr){
	return static_cast<double>(static_cast<QLineF*>(ptr)->x2());
}

double QLineF_Y1(void* ptr){
	return static_cast<double>(static_cast<QLineF*>(ptr)->y1());
}

double QLineF_Y2(void* ptr){
	return static_cast<double>(static_cast<QLineF*>(ptr)->y2());
}

