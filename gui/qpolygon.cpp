#include "qpolygon.h"
#include <QRect>
#include <QPoint>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPolygon>
#include "_cgo_export.h"

class MyQPolygon: public QPolygon {
public:
};

void* QPolygon_NewQPolygon5(void* rectangle, int closed){
	return new QPolygon(*static_cast<QRect*>(rectangle), closed != 0);
}

int QPolygon_ContainsPoint(void* ptr, void* point, int fillRule){
	return static_cast<QPolygon*>(ptr)->containsPoint(*static_cast<QPoint*>(point), static_cast<Qt::FillRule>(fillRule));
}

void QPolygon_PutPoints3(void* ptr, int index, int nPoints, void* fromPolygon, int fromIndex){
	static_cast<QPolygon*>(ptr)->putPoints(index, nPoints, *static_cast<QPolygon*>(fromPolygon), fromIndex);
}

void* QPolygon_NewQPolygon(){
	return new QPolygon();
}

void* QPolygon_NewQPolygon3(void* polygon){
	return new QPolygon(*static_cast<QPolygon*>(polygon));
}

void* QPolygon_NewQPolygon2(int size){
	return new QPolygon(size);
}

void QPolygon_Point(void* ptr, int index, int x, int y){
	static_cast<QPolygon*>(ptr)->point(index, &x, &y);
}

void QPolygon_SetPoint2(void* ptr, int index, void* point){
	static_cast<QPolygon*>(ptr)->setPoint(index, *static_cast<QPoint*>(point));
}

void QPolygon_SetPoint(void* ptr, int index, int x, int y){
	static_cast<QPolygon*>(ptr)->setPoint(index, x, y);
}

void QPolygon_SetPoints(void* ptr, int nPoints, int points){
	static_cast<QPolygon*>(ptr)->setPoints(nPoints, &points);
}

void QPolygon_Swap(void* ptr, void* other){
	static_cast<QPolygon*>(ptr)->swap(*static_cast<QPolygon*>(other));
}

void QPolygon_Translate2(void* ptr, void* offset){
	static_cast<QPolygon*>(ptr)->translate(*static_cast<QPoint*>(offset));
}

void QPolygon_Translate(void* ptr, int dx, int dy){
	static_cast<QPolygon*>(ptr)->translate(dx, dy);
}

void QPolygon_DestroyQPolygon(void* ptr){
	static_cast<QPolygon*>(ptr)->~QPolygon();
}

