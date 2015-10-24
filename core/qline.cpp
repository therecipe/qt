#include "qline.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPoint>
#include <QLine>
#include "_cgo_export.h"

class MyQLine: public QLine {
public:
};

QtObjectPtr QLine_NewQLine(){
	return new QLine();
}

QtObjectPtr QLine_NewQLine2(QtObjectPtr p1, QtObjectPtr p2){
	return new QLine(*static_cast<QPoint*>(p1), *static_cast<QPoint*>(p2));
}

QtObjectPtr QLine_NewQLine3(int x1, int y1, int x2, int y2){
	return new QLine(x1, y1, x2, y2);
}

int QLine_Dx(QtObjectPtr ptr){
	return static_cast<QLine*>(ptr)->dx();
}

int QLine_Dy(QtObjectPtr ptr){
	return static_cast<QLine*>(ptr)->dy();
}

int QLine_IsNull(QtObjectPtr ptr){
	return static_cast<QLine*>(ptr)->isNull();
}

void QLine_SetLine(QtObjectPtr ptr, int x1, int y1, int x2, int y2){
	static_cast<QLine*>(ptr)->setLine(x1, y1, x2, y2);
}

void QLine_SetP1(QtObjectPtr ptr, QtObjectPtr p1){
	static_cast<QLine*>(ptr)->setP1(*static_cast<QPoint*>(p1));
}

void QLine_SetP2(QtObjectPtr ptr, QtObjectPtr p2){
	static_cast<QLine*>(ptr)->setP2(*static_cast<QPoint*>(p2));
}

void QLine_SetPoints(QtObjectPtr ptr, QtObjectPtr p1, QtObjectPtr p2){
	static_cast<QLine*>(ptr)->setPoints(*static_cast<QPoint*>(p1), *static_cast<QPoint*>(p2));
}

void QLine_Translate(QtObjectPtr ptr, QtObjectPtr offset){
	static_cast<QLine*>(ptr)->translate(*static_cast<QPoint*>(offset));
}

void QLine_Translate2(QtObjectPtr ptr, int dx, int dy){
	static_cast<QLine*>(ptr)->translate(dx, dy);
}

int QLine_X1(QtObjectPtr ptr){
	return static_cast<QLine*>(ptr)->x1();
}

int QLine_X2(QtObjectPtr ptr){
	return static_cast<QLine*>(ptr)->x2();
}

int QLine_Y1(QtObjectPtr ptr){
	return static_cast<QLine*>(ptr)->y1();
}

int QLine_Y2(QtObjectPtr ptr){
	return static_cast<QLine*>(ptr)->y2();
}

