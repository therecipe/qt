#include "qrectf.h"
#include <QRect>
#include <QPointF>
#include <QSizeF>
#include <QSize>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPoint>
#include <QRectF>
#include "_cgo_export.h"

class MyQRectF: public QRectF {
public:
};

int QRectF_Contains(QtObjectPtr ptr, QtObjectPtr point){
	return static_cast<QRectF*>(ptr)->contains(*static_cast<QPointF*>(point));
}

int QRectF_Contains3(QtObjectPtr ptr, QtObjectPtr rectangle){
	return static_cast<QRectF*>(ptr)->contains(*static_cast<QRectF*>(rectangle));
}

int QRectF_Intersects(QtObjectPtr ptr, QtObjectPtr rectangle){
	return static_cast<QRectF*>(ptr)->intersects(*static_cast<QRectF*>(rectangle));
}

QtObjectPtr QRectF_NewQRectF(){
	return new QRectF();
}

QtObjectPtr QRectF_NewQRectF3(QtObjectPtr topLeft, QtObjectPtr bottomRight){
	return new QRectF(*static_cast<QPointF*>(topLeft), *static_cast<QPointF*>(bottomRight));
}

QtObjectPtr QRectF_NewQRectF2(QtObjectPtr topLeft, QtObjectPtr size){
	return new QRectF(*static_cast<QPointF*>(topLeft), *static_cast<QSizeF*>(size));
}

QtObjectPtr QRectF_NewQRectF5(QtObjectPtr rectangle){
	return new QRectF(*static_cast<QRect*>(rectangle));
}

int QRectF_IsEmpty(QtObjectPtr ptr){
	return static_cast<QRectF*>(ptr)->isEmpty();
}

int QRectF_IsNull(QtObjectPtr ptr){
	return static_cast<QRectF*>(ptr)->isNull();
}

int QRectF_IsValid(QtObjectPtr ptr){
	return static_cast<QRectF*>(ptr)->isValid();
}

void QRectF_MoveBottomLeft(QtObjectPtr ptr, QtObjectPtr position){
	static_cast<QRectF*>(ptr)->moveBottomLeft(*static_cast<QPointF*>(position));
}

void QRectF_MoveBottomRight(QtObjectPtr ptr, QtObjectPtr position){
	static_cast<QRectF*>(ptr)->moveBottomRight(*static_cast<QPointF*>(position));
}

void QRectF_MoveCenter(QtObjectPtr ptr, QtObjectPtr position){
	static_cast<QRectF*>(ptr)->moveCenter(*static_cast<QPointF*>(position));
}

void QRectF_MoveTo2(QtObjectPtr ptr, QtObjectPtr position){
	static_cast<QRectF*>(ptr)->moveTo(*static_cast<QPointF*>(position));
}

void QRectF_MoveTopLeft(QtObjectPtr ptr, QtObjectPtr position){
	static_cast<QRectF*>(ptr)->moveTopLeft(*static_cast<QPointF*>(position));
}

void QRectF_MoveTopRight(QtObjectPtr ptr, QtObjectPtr position){
	static_cast<QRectF*>(ptr)->moveTopRight(*static_cast<QPointF*>(position));
}

void QRectF_SetBottomLeft(QtObjectPtr ptr, QtObjectPtr position){
	static_cast<QRectF*>(ptr)->setBottomLeft(*static_cast<QPointF*>(position));
}

void QRectF_SetBottomRight(QtObjectPtr ptr, QtObjectPtr position){
	static_cast<QRectF*>(ptr)->setBottomRight(*static_cast<QPointF*>(position));
}

void QRectF_SetSize(QtObjectPtr ptr, QtObjectPtr size){
	static_cast<QRectF*>(ptr)->setSize(*static_cast<QSizeF*>(size));
}

void QRectF_SetTopLeft(QtObjectPtr ptr, QtObjectPtr position){
	static_cast<QRectF*>(ptr)->setTopLeft(*static_cast<QPointF*>(position));
}

void QRectF_SetTopRight(QtObjectPtr ptr, QtObjectPtr position){
	static_cast<QRectF*>(ptr)->setTopRight(*static_cast<QPointF*>(position));
}

void QRectF_Translate2(QtObjectPtr ptr, QtObjectPtr offset){
	static_cast<QRectF*>(ptr)->translate(*static_cast<QPointF*>(offset));
}

