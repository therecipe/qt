#include "qregion.h"
#include <QBitmap>
#include <QRect>
#include <QPolygon>
#include <QPoint>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QRegion>
#include "_cgo_export.h"

class MyQRegion: public QRegion {
public:
};

void* QRegion_NewQRegion(){
	return new QRegion();
}

void* QRegion_NewQRegion5(void* bm){
	return new QRegion(*static_cast<QBitmap*>(bm));
}

void* QRegion_NewQRegion3(void* a, int fillRule){
	return new QRegion(*static_cast<QPolygon*>(a), static_cast<Qt::FillRule>(fillRule));
}

void* QRegion_NewQRegion6(void* r, int t){
	return new QRegion(*static_cast<QRect*>(r), static_cast<QRegion::RegionType>(t));
}

void* QRegion_NewQRegion4(void* r){
	return new QRegion(*static_cast<QRegion*>(r));
}

int QRegion_Contains(void* ptr, void* p){
	return static_cast<QRegion*>(ptr)->contains(*static_cast<QPoint*>(p));
}

int QRegion_Contains2(void* ptr, void* r){
	return static_cast<QRegion*>(ptr)->contains(*static_cast<QRect*>(r));
}

void* QRegion_Intersected2(void* ptr, void* rect){
	return new QRegion(static_cast<QRegion*>(ptr)->intersected(*static_cast<QRect*>(rect)));
}

void* QRegion_Intersected(void* ptr, void* r){
	return new QRegion(static_cast<QRegion*>(ptr)->intersected(*static_cast<QRegion*>(r)));
}

int QRegion_Intersects2(void* ptr, void* rect){
	return static_cast<QRegion*>(ptr)->intersects(*static_cast<QRect*>(rect));
}

int QRegion_IsEmpty(void* ptr){
	return static_cast<QRegion*>(ptr)->isEmpty();
}

int QRegion_IsNull(void* ptr){
	return static_cast<QRegion*>(ptr)->isNull();
}

int QRegion_RectCount(void* ptr){
	return static_cast<QRegion*>(ptr)->rectCount();
}

void QRegion_SetRects(void* ptr, void* rects, int number){
	static_cast<QRegion*>(ptr)->setRects(static_cast<QRect*>(rects), number);
}

void* QRegion_Subtracted(void* ptr, void* r){
	return new QRegion(static_cast<QRegion*>(ptr)->subtracted(*static_cast<QRegion*>(r)));
}

void QRegion_Translate(void* ptr, int dx, int dy){
	static_cast<QRegion*>(ptr)->translate(dx, dy);
}

void* QRegion_United2(void* ptr, void* rect){
	return new QRegion(static_cast<QRegion*>(ptr)->united(*static_cast<QRect*>(rect)));
}

void* QRegion_United(void* ptr, void* r){
	return new QRegion(static_cast<QRegion*>(ptr)->united(*static_cast<QRegion*>(r)));
}

void* QRegion_Xored(void* ptr, void* r){
	return new QRegion(static_cast<QRegion*>(ptr)->xored(*static_cast<QRegion*>(r)));
}

void* QRegion_NewQRegion2(int x, int y, int w, int h, int t){
	return new QRegion(x, y, w, h, static_cast<QRegion::RegionType>(t));
}

int QRegion_Intersects(void* ptr, void* region){
	return static_cast<QRegion*>(ptr)->intersects(*static_cast<QRegion*>(region));
}

void QRegion_Swap(void* ptr, void* other){
	static_cast<QRegion*>(ptr)->swap(*static_cast<QRegion*>(other));
}

void QRegion_Translate2(void* ptr, void* point){
	static_cast<QRegion*>(ptr)->translate(*static_cast<QPoint*>(point));
}

void* QRegion_Translated2(void* ptr, void* p){
	return new QRegion(static_cast<QRegion*>(ptr)->translated(*static_cast<QPoint*>(p)));
}

void* QRegion_Translated(void* ptr, int dx, int dy){
	return new QRegion(static_cast<QRegion*>(ptr)->translated(dx, dy));
}

