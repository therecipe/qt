#include "qrect.h"
#include <QUrl>
#include <QModelIndex>
#include <QPoint>
#include <QSize>
#include <QString>
#include <QVariant>
#include <QRect>
#include "_cgo_export.h"

class MyQRect: public QRect {
public:
};

int QRect_Contains(void* ptr, void* point, int proper){
	return static_cast<QRect*>(ptr)->contains(*static_cast<QPoint*>(point), proper != 0);
}

int QRect_Contains4(void* ptr, void* rectangle, int proper){
	return static_cast<QRect*>(ptr)->contains(*static_cast<QRect*>(rectangle), proper != 0);
}

int QRect_Intersects(void* ptr, void* rectangle){
	return static_cast<QRect*>(ptr)->intersects(*static_cast<QRect*>(rectangle));
}

void* QRect_NewQRect(){
	return new QRect();
}

void* QRect_NewQRect2(void* topLeft, void* bottomRight){
	return new QRect(*static_cast<QPoint*>(topLeft), *static_cast<QPoint*>(bottomRight));
}

void* QRect_NewQRect3(void* topLeft, void* size){
	return new QRect(*static_cast<QPoint*>(topLeft), *static_cast<QSize*>(size));
}

void* QRect_NewQRect4(int x, int y, int width, int height){
	return new QRect(x, y, width, height);
}

void QRect_Adjust(void* ptr, int dx1, int dy1, int dx2, int dy2){
	static_cast<QRect*>(ptr)->adjust(dx1, dy1, dx2, dy2);
}

int QRect_Bottom(void* ptr){
	return static_cast<QRect*>(ptr)->bottom();
}

int QRect_Contains3(void* ptr, int x, int y){
	return static_cast<QRect*>(ptr)->contains(x, y);
}

int QRect_Contains2(void* ptr, int x, int y, int proper){
	return static_cast<QRect*>(ptr)->contains(x, y, proper != 0);
}

void QRect_GetCoords(void* ptr, int x1, int y1, int x2, int y2){
	static_cast<QRect*>(ptr)->getCoords(&x1, &y1, &x2, &y2);
}

void QRect_GetRect(void* ptr, int x, int y, int width, int height){
	static_cast<QRect*>(ptr)->getRect(&x, &y, &width, &height);
}

int QRect_Height(void* ptr){
	return static_cast<QRect*>(ptr)->height();
}

int QRect_IsEmpty(void* ptr){
	return static_cast<QRect*>(ptr)->isEmpty();
}

int QRect_IsNull(void* ptr){
	return static_cast<QRect*>(ptr)->isNull();
}

int QRect_IsValid(void* ptr){
	return static_cast<QRect*>(ptr)->isValid();
}

int QRect_Left(void* ptr){
	return static_cast<QRect*>(ptr)->left();
}

void QRect_MoveBottom(void* ptr, int y){
	static_cast<QRect*>(ptr)->moveBottom(y);
}

void QRect_MoveBottomLeft(void* ptr, void* position){
	static_cast<QRect*>(ptr)->moveBottomLeft(*static_cast<QPoint*>(position));
}

void QRect_MoveBottomRight(void* ptr, void* position){
	static_cast<QRect*>(ptr)->moveBottomRight(*static_cast<QPoint*>(position));
}

void QRect_MoveCenter(void* ptr, void* position){
	static_cast<QRect*>(ptr)->moveCenter(*static_cast<QPoint*>(position));
}

void QRect_MoveLeft(void* ptr, int x){
	static_cast<QRect*>(ptr)->moveLeft(x);
}

void QRect_MoveRight(void* ptr, int x){
	static_cast<QRect*>(ptr)->moveRight(x);
}

void QRect_MoveTo2(void* ptr, void* position){
	static_cast<QRect*>(ptr)->moveTo(*static_cast<QPoint*>(position));
}

void QRect_MoveTo(void* ptr, int x, int y){
	static_cast<QRect*>(ptr)->moveTo(x, y);
}

void QRect_MoveTop(void* ptr, int y){
	static_cast<QRect*>(ptr)->moveTop(y);
}

void QRect_MoveTopLeft(void* ptr, void* position){
	static_cast<QRect*>(ptr)->moveTopLeft(*static_cast<QPoint*>(position));
}

void QRect_MoveTopRight(void* ptr, void* position){
	static_cast<QRect*>(ptr)->moveTopRight(*static_cast<QPoint*>(position));
}

int QRect_Right(void* ptr){
	return static_cast<QRect*>(ptr)->right();
}

void QRect_SetBottom(void* ptr, int y){
	static_cast<QRect*>(ptr)->setBottom(y);
}

void QRect_SetBottomLeft(void* ptr, void* position){
	static_cast<QRect*>(ptr)->setBottomLeft(*static_cast<QPoint*>(position));
}

void QRect_SetBottomRight(void* ptr, void* position){
	static_cast<QRect*>(ptr)->setBottomRight(*static_cast<QPoint*>(position));
}

void QRect_SetCoords(void* ptr, int x1, int y1, int x2, int y2){
	static_cast<QRect*>(ptr)->setCoords(x1, y1, x2, y2);
}

void QRect_SetHeight(void* ptr, int height){
	static_cast<QRect*>(ptr)->setHeight(height);
}

void QRect_SetLeft(void* ptr, int x){
	static_cast<QRect*>(ptr)->setLeft(x);
}

void QRect_SetRect(void* ptr, int x, int y, int width, int height){
	static_cast<QRect*>(ptr)->setRect(x, y, width, height);
}

void QRect_SetRight(void* ptr, int x){
	static_cast<QRect*>(ptr)->setRight(x);
}

void QRect_SetSize(void* ptr, void* size){
	static_cast<QRect*>(ptr)->setSize(*static_cast<QSize*>(size));
}

void QRect_SetTop(void* ptr, int y){
	static_cast<QRect*>(ptr)->setTop(y);
}

void QRect_SetTopLeft(void* ptr, void* position){
	static_cast<QRect*>(ptr)->setTopLeft(*static_cast<QPoint*>(position));
}

void QRect_SetTopRight(void* ptr, void* position){
	static_cast<QRect*>(ptr)->setTopRight(*static_cast<QPoint*>(position));
}

void QRect_SetWidth(void* ptr, int width){
	static_cast<QRect*>(ptr)->setWidth(width);
}

void QRect_SetX(void* ptr, int x){
	static_cast<QRect*>(ptr)->setX(x);
}

void QRect_SetY(void* ptr, int y){
	static_cast<QRect*>(ptr)->setY(y);
}

int QRect_Top(void* ptr){
	return static_cast<QRect*>(ptr)->top();
}

void QRect_Translate2(void* ptr, void* offset){
	static_cast<QRect*>(ptr)->translate(*static_cast<QPoint*>(offset));
}

void QRect_Translate(void* ptr, int dx, int dy){
	static_cast<QRect*>(ptr)->translate(dx, dy);
}

int QRect_Width(void* ptr){
	return static_cast<QRect*>(ptr)->width();
}

int QRect_X(void* ptr){
	return static_cast<QRect*>(ptr)->x();
}

int QRect_Y(void* ptr){
	return static_cast<QRect*>(ptr)->y();
}

