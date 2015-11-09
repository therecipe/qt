#include "qline.h"
#include <QModelIndex>
#include <QPoint>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QLine>
#include "_cgo_export.h"

class MyQLine: public QLine {
public:
};

void* QLine_NewQLine(){
	return new QLine();
}

void* QLine_NewQLine2(void* p1, void* p2){
	return new QLine(*static_cast<QPoint*>(p1), *static_cast<QPoint*>(p2));
}

void* QLine_NewQLine3(int x1, int y1, int x2, int y2){
	return new QLine(x1, y1, x2, y2);
}

int QLine_Dx(void* ptr){
	return static_cast<QLine*>(ptr)->dx();
}

int QLine_Dy(void* ptr){
	return static_cast<QLine*>(ptr)->dy();
}

int QLine_IsNull(void* ptr){
	return static_cast<QLine*>(ptr)->isNull();
}

void QLine_SetLine(void* ptr, int x1, int y1, int x2, int y2){
	static_cast<QLine*>(ptr)->setLine(x1, y1, x2, y2);
}

void QLine_SetP1(void* ptr, void* p1){
	static_cast<QLine*>(ptr)->setP1(*static_cast<QPoint*>(p1));
}

void QLine_SetP2(void* ptr, void* p2){
	static_cast<QLine*>(ptr)->setP2(*static_cast<QPoint*>(p2));
}

void QLine_SetPoints(void* ptr, void* p1, void* p2){
	static_cast<QLine*>(ptr)->setPoints(*static_cast<QPoint*>(p1), *static_cast<QPoint*>(p2));
}

void QLine_Translate(void* ptr, void* offset){
	static_cast<QLine*>(ptr)->translate(*static_cast<QPoint*>(offset));
}

void QLine_Translate2(void* ptr, int dx, int dy){
	static_cast<QLine*>(ptr)->translate(dx, dy);
}

int QLine_X1(void* ptr){
	return static_cast<QLine*>(ptr)->x1();
}

int QLine_X2(void* ptr){
	return static_cast<QLine*>(ptr)->x2();
}

int QLine_Y1(void* ptr){
	return static_cast<QLine*>(ptr)->y1();
}

int QLine_Y2(void* ptr){
	return static_cast<QLine*>(ptr)->y2();
}

