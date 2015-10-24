#include "qpolygon.h"
#include <QPoint>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QRect>
#include <QPolygon>
#include "_cgo_export.h"

class MyQPolygon: public QPolygon {
public:
};

QtObjectPtr QPolygon_NewQPolygon5(QtObjectPtr rectangle, int closed){
	return new QPolygon(*static_cast<QRect*>(rectangle), closed != 0);
}

int QPolygon_ContainsPoint(QtObjectPtr ptr, QtObjectPtr point, int fillRule){
	return static_cast<QPolygon*>(ptr)->containsPoint(*static_cast<QPoint*>(point), static_cast<Qt::FillRule>(fillRule));
}

void QPolygon_PutPoints3(QtObjectPtr ptr, int index, int nPoints, QtObjectPtr fromPolygon, int fromIndex){
	static_cast<QPolygon*>(ptr)->putPoints(index, nPoints, *static_cast<QPolygon*>(fromPolygon), fromIndex);
}

QtObjectPtr QPolygon_NewQPolygon(){
	return new QPolygon();
}

QtObjectPtr QPolygon_NewQPolygon3(QtObjectPtr polygon){
	return new QPolygon(*static_cast<QPolygon*>(polygon));
}

QtObjectPtr QPolygon_NewQPolygon2(int size){
	return new QPolygon(size);
}

void QPolygon_Point(QtObjectPtr ptr, int index, int x, int y){
	static_cast<QPolygon*>(ptr)->point(index, &x, &y);
}

void QPolygon_SetPoint2(QtObjectPtr ptr, int index, QtObjectPtr point){
	static_cast<QPolygon*>(ptr)->setPoint(index, *static_cast<QPoint*>(point));
}

void QPolygon_SetPoint(QtObjectPtr ptr, int index, int x, int y){
	static_cast<QPolygon*>(ptr)->setPoint(index, x, y);
}

void QPolygon_SetPoints(QtObjectPtr ptr, int nPoints, int points){
	static_cast<QPolygon*>(ptr)->setPoints(nPoints, &points);
}

void QPolygon_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QPolygon*>(ptr)->swap(*static_cast<QPolygon*>(other));
}

void QPolygon_Translate2(QtObjectPtr ptr, QtObjectPtr offset){
	static_cast<QPolygon*>(ptr)->translate(*static_cast<QPoint*>(offset));
}

void QPolygon_Translate(QtObjectPtr ptr, int dx, int dy){
	static_cast<QPolygon*>(ptr)->translate(dx, dy);
}

void QPolygon_DestroyQPolygon(QtObjectPtr ptr){
	static_cast<QPolygon*>(ptr)->~QPolygon();
}

