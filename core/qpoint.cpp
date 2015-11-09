#include "qpoint.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPoint>
#include "_cgo_export.h"

class MyQPoint: public QPoint {
public:
};

void* QPoint_NewQPoint(){
	return new QPoint();
}

void* QPoint_NewQPoint2(int xpos, int ypos){
	return new QPoint(xpos, ypos);
}

int QPoint_QPoint_DotProduct(void* p1, void* p2){
	return QPoint::dotProduct(*static_cast<QPoint*>(p1), *static_cast<QPoint*>(p2));
}

int QPoint_IsNull(void* ptr){
	return static_cast<QPoint*>(ptr)->isNull();
}

int QPoint_ManhattanLength(void* ptr){
	return static_cast<QPoint*>(ptr)->manhattanLength();
}

int QPoint_Rx(void* ptr){
	return static_cast<QPoint*>(ptr)->rx();
}

int QPoint_Ry(void* ptr){
	return static_cast<QPoint*>(ptr)->ry();
}

void QPoint_SetX(void* ptr, int x){
	static_cast<QPoint*>(ptr)->setX(x);
}

void QPoint_SetY(void* ptr, int y){
	static_cast<QPoint*>(ptr)->setY(y);
}

int QPoint_X(void* ptr){
	return static_cast<QPoint*>(ptr)->x();
}

int QPoint_Y(void* ptr){
	return static_cast<QPoint*>(ptr)->y();
}

