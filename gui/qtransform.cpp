#include "qtransform.h"
#include <QPolygon>
#include <QPolygonF>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTransform>
#include "_cgo_export.h"

class MyQTransform: public QTransform {
public:
};

int QTransform_QTransform_QuadToSquare(QtObjectPtr quad, QtObjectPtr trans){
	return QTransform::quadToSquare(*static_cast<QPolygonF*>(quad), *static_cast<QTransform*>(trans));
}

QtObjectPtr QTransform_NewQTransform(){
	return new QTransform();
}

int QTransform_IsAffine(QtObjectPtr ptr){
	return static_cast<QTransform*>(ptr)->isAffine();
}

int QTransform_IsIdentity(QtObjectPtr ptr){
	return static_cast<QTransform*>(ptr)->isIdentity();
}

int QTransform_IsInvertible(QtObjectPtr ptr){
	return static_cast<QTransform*>(ptr)->isInvertible();
}

int QTransform_IsRotating(QtObjectPtr ptr){
	return static_cast<QTransform*>(ptr)->isRotating();
}

int QTransform_IsScaling(QtObjectPtr ptr){
	return static_cast<QTransform*>(ptr)->isScaling();
}

int QTransform_IsTranslating(QtObjectPtr ptr){
	return static_cast<QTransform*>(ptr)->isTranslating();
}

void QTransform_Map10(QtObjectPtr ptr, int x, int y, int tx, int ty){
	static_cast<QTransform*>(ptr)->map(x, y, &tx, &ty);
}

int QTransform_QTransform_QuadToQuad(QtObjectPtr one, QtObjectPtr two, QtObjectPtr trans){
	return QTransform::quadToQuad(*static_cast<QPolygonF*>(one), *static_cast<QPolygonF*>(two), *static_cast<QTransform*>(trans));
}

void QTransform_Reset(QtObjectPtr ptr){
	static_cast<QTransform*>(ptr)->reset();
}

int QTransform_QTransform_SquareToQuad(QtObjectPtr quad, QtObjectPtr trans){
	return QTransform::squareToQuad(*static_cast<QPolygonF*>(quad), *static_cast<QTransform*>(trans));
}

int QTransform_Type(QtObjectPtr ptr){
	return static_cast<QTransform*>(ptr)->type();
}

