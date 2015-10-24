#include "qrect.h"
#include <QPoint>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSize>
#include <QRect>
#include "_cgo_export.h"

class MyQRect: public QRect {
public:
};

int QRect_Contains(QtObjectPtr ptr, QtObjectPtr point, int proper){
	return static_cast<QRect*>(ptr)->contains(*static_cast<QPoint*>(point), proper != 0);
}

int QRect_Contains4(QtObjectPtr ptr, QtObjectPtr rectangle, int proper){
	return static_cast<QRect*>(ptr)->contains(*static_cast<QRect*>(rectangle), proper != 0);
}

int QRect_Intersects(QtObjectPtr ptr, QtObjectPtr rectangle){
	return static_cast<QRect*>(ptr)->intersects(*static_cast<QRect*>(rectangle));
}

QtObjectPtr QRect_NewQRect(){
	return new QRect();
}

QtObjectPtr QRect_NewQRect2(QtObjectPtr topLeft, QtObjectPtr bottomRight){
	return new QRect(*static_cast<QPoint*>(topLeft), *static_cast<QPoint*>(bottomRight));
}

QtObjectPtr QRect_NewQRect3(QtObjectPtr topLeft, QtObjectPtr size){
	return new QRect(*static_cast<QPoint*>(topLeft), *static_cast<QSize*>(size));
}

QtObjectPtr QRect_NewQRect4(int x, int y, int width, int height){
	return new QRect(x, y, width, height);
}

void QRect_Adjust(QtObjectPtr ptr, int dx1, int dy1, int dx2, int dy2){
	static_cast<QRect*>(ptr)->adjust(dx1, dy1, dx2, dy2);
}

int QRect_Bottom(QtObjectPtr ptr){
	return static_cast<QRect*>(ptr)->bottom();
}

int QRect_Contains3(QtObjectPtr ptr, int x, int y){
	return static_cast<QRect*>(ptr)->contains(x, y);
}

int QRect_Contains2(QtObjectPtr ptr, int x, int y, int proper){
	return static_cast<QRect*>(ptr)->contains(x, y, proper != 0);
}

void QRect_GetCoords(QtObjectPtr ptr, int x1, int y1, int x2, int y2){
	static_cast<QRect*>(ptr)->getCoords(&x1, &y1, &x2, &y2);
}

void QRect_GetRect(QtObjectPtr ptr, int x, int y, int width, int height){
	static_cast<QRect*>(ptr)->getRect(&x, &y, &width, &height);
}

int QRect_Height(QtObjectPtr ptr){
	return static_cast<QRect*>(ptr)->height();
}

int QRect_IsEmpty(QtObjectPtr ptr){
	return static_cast<QRect*>(ptr)->isEmpty();
}

int QRect_IsNull(QtObjectPtr ptr){
	return static_cast<QRect*>(ptr)->isNull();
}

int QRect_IsValid(QtObjectPtr ptr){
	return static_cast<QRect*>(ptr)->isValid();
}

int QRect_Left(QtObjectPtr ptr){
	return static_cast<QRect*>(ptr)->left();
}

void QRect_MoveBottom(QtObjectPtr ptr, int y){
	static_cast<QRect*>(ptr)->moveBottom(y);
}

void QRect_MoveBottomLeft(QtObjectPtr ptr, QtObjectPtr position){
	static_cast<QRect*>(ptr)->moveBottomLeft(*static_cast<QPoint*>(position));
}

void QRect_MoveBottomRight(QtObjectPtr ptr, QtObjectPtr position){
	static_cast<QRect*>(ptr)->moveBottomRight(*static_cast<QPoint*>(position));
}

void QRect_MoveCenter(QtObjectPtr ptr, QtObjectPtr position){
	static_cast<QRect*>(ptr)->moveCenter(*static_cast<QPoint*>(position));
}

void QRect_MoveLeft(QtObjectPtr ptr, int x){
	static_cast<QRect*>(ptr)->moveLeft(x);
}

void QRect_MoveRight(QtObjectPtr ptr, int x){
	static_cast<QRect*>(ptr)->moveRight(x);
}

void QRect_MoveTo2(QtObjectPtr ptr, QtObjectPtr position){
	static_cast<QRect*>(ptr)->moveTo(*static_cast<QPoint*>(position));
}

void QRect_MoveTo(QtObjectPtr ptr, int x, int y){
	static_cast<QRect*>(ptr)->moveTo(x, y);
}

void QRect_MoveTop(QtObjectPtr ptr, int y){
	static_cast<QRect*>(ptr)->moveTop(y);
}

void QRect_MoveTopLeft(QtObjectPtr ptr, QtObjectPtr position){
	static_cast<QRect*>(ptr)->moveTopLeft(*static_cast<QPoint*>(position));
}

void QRect_MoveTopRight(QtObjectPtr ptr, QtObjectPtr position){
	static_cast<QRect*>(ptr)->moveTopRight(*static_cast<QPoint*>(position));
}

int QRect_Right(QtObjectPtr ptr){
	return static_cast<QRect*>(ptr)->right();
}

void QRect_SetBottom(QtObjectPtr ptr, int y){
	static_cast<QRect*>(ptr)->setBottom(y);
}

void QRect_SetBottomLeft(QtObjectPtr ptr, QtObjectPtr position){
	static_cast<QRect*>(ptr)->setBottomLeft(*static_cast<QPoint*>(position));
}

void QRect_SetBottomRight(QtObjectPtr ptr, QtObjectPtr position){
	static_cast<QRect*>(ptr)->setBottomRight(*static_cast<QPoint*>(position));
}

void QRect_SetCoords(QtObjectPtr ptr, int x1, int y1, int x2, int y2){
	static_cast<QRect*>(ptr)->setCoords(x1, y1, x2, y2);
}

void QRect_SetHeight(QtObjectPtr ptr, int height){
	static_cast<QRect*>(ptr)->setHeight(height);
}

void QRect_SetLeft(QtObjectPtr ptr, int x){
	static_cast<QRect*>(ptr)->setLeft(x);
}

void QRect_SetRect(QtObjectPtr ptr, int x, int y, int width, int height){
	static_cast<QRect*>(ptr)->setRect(x, y, width, height);
}

void QRect_SetRight(QtObjectPtr ptr, int x){
	static_cast<QRect*>(ptr)->setRight(x);
}

void QRect_SetSize(QtObjectPtr ptr, QtObjectPtr size){
	static_cast<QRect*>(ptr)->setSize(*static_cast<QSize*>(size));
}

void QRect_SetTop(QtObjectPtr ptr, int y){
	static_cast<QRect*>(ptr)->setTop(y);
}

void QRect_SetTopLeft(QtObjectPtr ptr, QtObjectPtr position){
	static_cast<QRect*>(ptr)->setTopLeft(*static_cast<QPoint*>(position));
}

void QRect_SetTopRight(QtObjectPtr ptr, QtObjectPtr position){
	static_cast<QRect*>(ptr)->setTopRight(*static_cast<QPoint*>(position));
}

void QRect_SetWidth(QtObjectPtr ptr, int width){
	static_cast<QRect*>(ptr)->setWidth(width);
}

void QRect_SetX(QtObjectPtr ptr, int x){
	static_cast<QRect*>(ptr)->setX(x);
}

void QRect_SetY(QtObjectPtr ptr, int y){
	static_cast<QRect*>(ptr)->setY(y);
}

int QRect_Top(QtObjectPtr ptr){
	return static_cast<QRect*>(ptr)->top();
}

void QRect_Translate2(QtObjectPtr ptr, QtObjectPtr offset){
	static_cast<QRect*>(ptr)->translate(*static_cast<QPoint*>(offset));
}

void QRect_Translate(QtObjectPtr ptr, int dx, int dy){
	static_cast<QRect*>(ptr)->translate(dx, dy);
}

int QRect_Width(QtObjectPtr ptr){
	return static_cast<QRect*>(ptr)->width();
}

int QRect_X(QtObjectPtr ptr){
	return static_cast<QRect*>(ptr)->x();
}

int QRect_Y(QtObjectPtr ptr){
	return static_cast<QRect*>(ptr)->y();
}

