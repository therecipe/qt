#include "qpoint.h"
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QPoint>
#include "_cgo_export.h"

class MyQPoint: public QPoint {
public:
};

QtObjectPtr QPoint_NewQPoint(){
	return new QPoint();
}

QtObjectPtr QPoint_NewQPoint2(int xpos, int ypos){
	return new QPoint(xpos, ypos);
}

int QPoint_QPoint_DotProduct(QtObjectPtr p1, QtObjectPtr p2){
	return QPoint::dotProduct(*static_cast<QPoint*>(p1), *static_cast<QPoint*>(p2));
}

int QPoint_IsNull(QtObjectPtr ptr){
	return static_cast<QPoint*>(ptr)->isNull();
}

int QPoint_ManhattanLength(QtObjectPtr ptr){
	return static_cast<QPoint*>(ptr)->manhattanLength();
}

int QPoint_Rx(QtObjectPtr ptr){
	return static_cast<QPoint*>(ptr)->rx();
}

int QPoint_Ry(QtObjectPtr ptr){
	return static_cast<QPoint*>(ptr)->ry();
}

void QPoint_SetX(QtObjectPtr ptr, int x){
	static_cast<QPoint*>(ptr)->setX(x);
}

void QPoint_SetY(QtObjectPtr ptr, int y){
	static_cast<QPoint*>(ptr)->setY(y);
}

int QPoint_X(QtObjectPtr ptr){
	return static_cast<QPoint*>(ptr)->x();
}

int QPoint_Y(QtObjectPtr ptr){
	return static_cast<QPoint*>(ptr)->y();
}

