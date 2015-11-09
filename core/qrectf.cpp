#include "qrectf.h"
#include <QString>
#include <QUrl>
#include <QRect>
#include <QSizeF>
#include <QPoint>
#include <QVariant>
#include <QModelIndex>
#include <QPointF>
#include <QSize>
#include <QRectF>
#include "_cgo_export.h"

class MyQRectF: public QRectF {
public:
};

int QRectF_Contains(void* ptr, void* point){
	return static_cast<QRectF*>(ptr)->contains(*static_cast<QPointF*>(point));
}

int QRectF_Contains3(void* ptr, void* rectangle){
	return static_cast<QRectF*>(ptr)->contains(*static_cast<QRectF*>(rectangle));
}

int QRectF_Intersects(void* ptr, void* rectangle){
	return static_cast<QRectF*>(ptr)->intersects(*static_cast<QRectF*>(rectangle));
}

void* QRectF_NewQRectF(){
	return new QRectF();
}

void* QRectF_NewQRectF3(void* topLeft, void* bottomRight){
	return new QRectF(*static_cast<QPointF*>(topLeft), *static_cast<QPointF*>(bottomRight));
}

void* QRectF_NewQRectF2(void* topLeft, void* size){
	return new QRectF(*static_cast<QPointF*>(topLeft), *static_cast<QSizeF*>(size));
}

void* QRectF_NewQRectF5(void* rectangle){
	return new QRectF(*static_cast<QRect*>(rectangle));
}

void* QRectF_NewQRectF4(double x, double y, double width, double height){
	return new QRectF(static_cast<qreal>(x), static_cast<qreal>(y), static_cast<qreal>(width), static_cast<qreal>(height));
}

void QRectF_Adjust(void* ptr, double dx1, double dy1, double dx2, double dy2){
	static_cast<QRectF*>(ptr)->adjust(static_cast<qreal>(dx1), static_cast<qreal>(dy1), static_cast<qreal>(dx2), static_cast<qreal>(dy2));
}

double QRectF_Bottom(void* ptr){
	return static_cast<double>(static_cast<QRectF*>(ptr)->bottom());
}

int QRectF_Contains2(void* ptr, double x, double y){
	return static_cast<QRectF*>(ptr)->contains(static_cast<qreal>(x), static_cast<qreal>(y));
}

double QRectF_Height(void* ptr){
	return static_cast<double>(static_cast<QRectF*>(ptr)->height());
}

int QRectF_IsEmpty(void* ptr){
	return static_cast<QRectF*>(ptr)->isEmpty();
}

int QRectF_IsNull(void* ptr){
	return static_cast<QRectF*>(ptr)->isNull();
}

int QRectF_IsValid(void* ptr){
	return static_cast<QRectF*>(ptr)->isValid();
}

double QRectF_Left(void* ptr){
	return static_cast<double>(static_cast<QRectF*>(ptr)->left());
}

void QRectF_MoveBottom(void* ptr, double y){
	static_cast<QRectF*>(ptr)->moveBottom(static_cast<qreal>(y));
}

void QRectF_MoveBottomLeft(void* ptr, void* position){
	static_cast<QRectF*>(ptr)->moveBottomLeft(*static_cast<QPointF*>(position));
}

void QRectF_MoveBottomRight(void* ptr, void* position){
	static_cast<QRectF*>(ptr)->moveBottomRight(*static_cast<QPointF*>(position));
}

void QRectF_MoveCenter(void* ptr, void* position){
	static_cast<QRectF*>(ptr)->moveCenter(*static_cast<QPointF*>(position));
}

void QRectF_MoveLeft(void* ptr, double x){
	static_cast<QRectF*>(ptr)->moveLeft(static_cast<qreal>(x));
}

void QRectF_MoveRight(void* ptr, double x){
	static_cast<QRectF*>(ptr)->moveRight(static_cast<qreal>(x));
}

void QRectF_MoveTo2(void* ptr, void* position){
	static_cast<QRectF*>(ptr)->moveTo(*static_cast<QPointF*>(position));
}

void QRectF_MoveTo(void* ptr, double x, double y){
	static_cast<QRectF*>(ptr)->moveTo(static_cast<qreal>(x), static_cast<qreal>(y));
}

void QRectF_MoveTop(void* ptr, double y){
	static_cast<QRectF*>(ptr)->moveTop(static_cast<qreal>(y));
}

void QRectF_MoveTopLeft(void* ptr, void* position){
	static_cast<QRectF*>(ptr)->moveTopLeft(*static_cast<QPointF*>(position));
}

void QRectF_MoveTopRight(void* ptr, void* position){
	static_cast<QRectF*>(ptr)->moveTopRight(*static_cast<QPointF*>(position));
}

double QRectF_Right(void* ptr){
	return static_cast<double>(static_cast<QRectF*>(ptr)->right());
}

void QRectF_SetBottom(void* ptr, double y){
	static_cast<QRectF*>(ptr)->setBottom(static_cast<qreal>(y));
}

void QRectF_SetBottomLeft(void* ptr, void* position){
	static_cast<QRectF*>(ptr)->setBottomLeft(*static_cast<QPointF*>(position));
}

void QRectF_SetBottomRight(void* ptr, void* position){
	static_cast<QRectF*>(ptr)->setBottomRight(*static_cast<QPointF*>(position));
}

void QRectF_SetCoords(void* ptr, double x1, double y1, double x2, double y2){
	static_cast<QRectF*>(ptr)->setCoords(static_cast<qreal>(x1), static_cast<qreal>(y1), static_cast<qreal>(x2), static_cast<qreal>(y2));
}

void QRectF_SetHeight(void* ptr, double height){
	static_cast<QRectF*>(ptr)->setHeight(static_cast<qreal>(height));
}

void QRectF_SetLeft(void* ptr, double x){
	static_cast<QRectF*>(ptr)->setLeft(static_cast<qreal>(x));
}

void QRectF_SetRect(void* ptr, double x, double y, double width, double height){
	static_cast<QRectF*>(ptr)->setRect(static_cast<qreal>(x), static_cast<qreal>(y), static_cast<qreal>(width), static_cast<qreal>(height));
}

void QRectF_SetRight(void* ptr, double x){
	static_cast<QRectF*>(ptr)->setRight(static_cast<qreal>(x));
}

void QRectF_SetSize(void* ptr, void* size){
	static_cast<QRectF*>(ptr)->setSize(*static_cast<QSizeF*>(size));
}

void QRectF_SetTop(void* ptr, double y){
	static_cast<QRectF*>(ptr)->setTop(static_cast<qreal>(y));
}

void QRectF_SetTopLeft(void* ptr, void* position){
	static_cast<QRectF*>(ptr)->setTopLeft(*static_cast<QPointF*>(position));
}

void QRectF_SetTopRight(void* ptr, void* position){
	static_cast<QRectF*>(ptr)->setTopRight(*static_cast<QPointF*>(position));
}

void QRectF_SetWidth(void* ptr, double width){
	static_cast<QRectF*>(ptr)->setWidth(static_cast<qreal>(width));
}

void QRectF_SetX(void* ptr, double x){
	static_cast<QRectF*>(ptr)->setX(static_cast<qreal>(x));
}

void QRectF_SetY(void* ptr, double y){
	static_cast<QRectF*>(ptr)->setY(static_cast<qreal>(y));
}

double QRectF_Top(void* ptr){
	return static_cast<double>(static_cast<QRectF*>(ptr)->top());
}

void QRectF_Translate2(void* ptr, void* offset){
	static_cast<QRectF*>(ptr)->translate(*static_cast<QPointF*>(offset));
}

void QRectF_Translate(void* ptr, double dx, double dy){
	static_cast<QRectF*>(ptr)->translate(static_cast<qreal>(dx), static_cast<qreal>(dy));
}

double QRectF_Width(void* ptr){
	return static_cast<double>(static_cast<QRectF*>(ptr)->width());
}

double QRectF_X(void* ptr){
	return static_cast<double>(static_cast<QRectF*>(ptr)->x());
}

double QRectF_Y(void* ptr){
	return static_cast<double>(static_cast<QRectF*>(ptr)->y());
}

