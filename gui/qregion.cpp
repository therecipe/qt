#include "qregion.h"
#include <QBitmap>
#include <QPoint>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPolygon>
#include <QRect>
#include <QRegion>
#include "_cgo_export.h"

class MyQRegion: public QRegion {
public:
};

QtObjectPtr QRegion_NewQRegion(){
	return new QRegion();
}

QtObjectPtr QRegion_NewQRegion5(QtObjectPtr bm){
	return new QRegion(*static_cast<QBitmap*>(bm));
}

QtObjectPtr QRegion_NewQRegion3(QtObjectPtr a, int fillRule){
	return new QRegion(*static_cast<QPolygon*>(a), static_cast<Qt::FillRule>(fillRule));
}

QtObjectPtr QRegion_NewQRegion6(QtObjectPtr r, int t){
	return new QRegion(*static_cast<QRect*>(r), static_cast<QRegion::RegionType>(t));
}

QtObjectPtr QRegion_NewQRegion4(QtObjectPtr r){
	return new QRegion(*static_cast<QRegion*>(r));
}

int QRegion_Contains(QtObjectPtr ptr, QtObjectPtr p){
	return static_cast<QRegion*>(ptr)->contains(*static_cast<QPoint*>(p));
}

int QRegion_Contains2(QtObjectPtr ptr, QtObjectPtr r){
	return static_cast<QRegion*>(ptr)->contains(*static_cast<QRect*>(r));
}

int QRegion_Intersects2(QtObjectPtr ptr, QtObjectPtr rect){
	return static_cast<QRegion*>(ptr)->intersects(*static_cast<QRect*>(rect));
}

int QRegion_IsEmpty(QtObjectPtr ptr){
	return static_cast<QRegion*>(ptr)->isEmpty();
}

int QRegion_IsNull(QtObjectPtr ptr){
	return static_cast<QRegion*>(ptr)->isNull();
}

int QRegion_RectCount(QtObjectPtr ptr){
	return static_cast<QRegion*>(ptr)->rectCount();
}

void QRegion_SetRects(QtObjectPtr ptr, QtObjectPtr rects, int number){
	static_cast<QRegion*>(ptr)->setRects(static_cast<QRect*>(rects), number);
}

void QRegion_Translate(QtObjectPtr ptr, int dx, int dy){
	static_cast<QRegion*>(ptr)->translate(dx, dy);
}

QtObjectPtr QRegion_NewQRegion2(int x, int y, int w, int h, int t){
	return new QRegion(x, y, w, h, static_cast<QRegion::RegionType>(t));
}

int QRegion_Intersects(QtObjectPtr ptr, QtObjectPtr region){
	return static_cast<QRegion*>(ptr)->intersects(*static_cast<QRegion*>(region));
}

void QRegion_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QRegion*>(ptr)->swap(*static_cast<QRegion*>(other));
}

void QRegion_Translate2(QtObjectPtr ptr, QtObjectPtr point){
	static_cast<QRegion*>(ptr)->translate(*static_cast<QPoint*>(point));
}

