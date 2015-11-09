#include "qtransform.h"
#include <QModelIndex>
#include <QPolygon>
#include <QRegion>
#include <QPolygonF>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QTransform>
#include "_cgo_export.h"

class MyQTransform: public QTransform {
public:
};

void* QTransform_NewQTransform3(double m11, double m12, double m13, double m21, double m22, double m23, double m31, double m32, double m33){
	return new QTransform(static_cast<qreal>(m11), static_cast<qreal>(m12), static_cast<qreal>(m13), static_cast<qreal>(m21), static_cast<qreal>(m22), static_cast<qreal>(m23), static_cast<qreal>(m31), static_cast<qreal>(m32), static_cast<qreal>(m33));
}

void* QTransform_NewQTransform4(double m11, double m12, double m21, double m22, double dx, double dy){
	return new QTransform(static_cast<qreal>(m11), static_cast<qreal>(m12), static_cast<qreal>(m21), static_cast<qreal>(m22), static_cast<qreal>(dx), static_cast<qreal>(dy));
}

void* QTransform_Map8(void* ptr, void* region){
	return new QRegion(static_cast<QTransform*>(ptr)->map(*static_cast<QRegion*>(region)));
}

int QTransform_QTransform_QuadToSquare(void* quad, void* trans){
	return QTransform::quadToSquare(*static_cast<QPolygonF*>(quad), *static_cast<QTransform*>(trans));
}

void* QTransform_NewQTransform(){
	return new QTransform();
}

double QTransform_Determinant(void* ptr){
	return static_cast<double>(static_cast<QTransform*>(ptr)->determinant());
}

double QTransform_Dx(void* ptr){
	return static_cast<double>(static_cast<QTransform*>(ptr)->dx());
}

double QTransform_Dy(void* ptr){
	return static_cast<double>(static_cast<QTransform*>(ptr)->dy());
}

int QTransform_IsAffine(void* ptr){
	return static_cast<QTransform*>(ptr)->isAffine();
}

int QTransform_IsIdentity(void* ptr){
	return static_cast<QTransform*>(ptr)->isIdentity();
}

int QTransform_IsInvertible(void* ptr){
	return static_cast<QTransform*>(ptr)->isInvertible();
}

int QTransform_IsRotating(void* ptr){
	return static_cast<QTransform*>(ptr)->isRotating();
}

int QTransform_IsScaling(void* ptr){
	return static_cast<QTransform*>(ptr)->isScaling();
}

int QTransform_IsTranslating(void* ptr){
	return static_cast<QTransform*>(ptr)->isTranslating();
}

double QTransform_M11(void* ptr){
	return static_cast<double>(static_cast<QTransform*>(ptr)->m11());
}

double QTransform_M12(void* ptr){
	return static_cast<double>(static_cast<QTransform*>(ptr)->m12());
}

double QTransform_M13(void* ptr){
	return static_cast<double>(static_cast<QTransform*>(ptr)->m13());
}

double QTransform_M21(void* ptr){
	return static_cast<double>(static_cast<QTransform*>(ptr)->m21());
}

double QTransform_M22(void* ptr){
	return static_cast<double>(static_cast<QTransform*>(ptr)->m22());
}

double QTransform_M23(void* ptr){
	return static_cast<double>(static_cast<QTransform*>(ptr)->m23());
}

double QTransform_M31(void* ptr){
	return static_cast<double>(static_cast<QTransform*>(ptr)->m31());
}

double QTransform_M32(void* ptr){
	return static_cast<double>(static_cast<QTransform*>(ptr)->m32());
}

double QTransform_M33(void* ptr){
	return static_cast<double>(static_cast<QTransform*>(ptr)->m33());
}

void QTransform_Map10(void* ptr, int x, int y, int tx, int ty){
	static_cast<QTransform*>(ptr)->map(x, y, &tx, &ty);
}

int QTransform_QTransform_QuadToQuad(void* one, void* two, void* trans){
	return QTransform::quadToQuad(*static_cast<QPolygonF*>(one), *static_cast<QPolygonF*>(two), *static_cast<QTransform*>(trans));
}

void QTransform_Reset(void* ptr){
	static_cast<QTransform*>(ptr)->reset();
}

void QTransform_SetMatrix(void* ptr, double m11, double m12, double m13, double m21, double m22, double m23, double m31, double m32, double m33){
	static_cast<QTransform*>(ptr)->setMatrix(static_cast<qreal>(m11), static_cast<qreal>(m12), static_cast<qreal>(m13), static_cast<qreal>(m21), static_cast<qreal>(m22), static_cast<qreal>(m23), static_cast<qreal>(m31), static_cast<qreal>(m32), static_cast<qreal>(m33));
}

int QTransform_QTransform_SquareToQuad(void* quad, void* trans){
	return QTransform::squareToQuad(*static_cast<QPolygonF*>(quad), *static_cast<QTransform*>(trans));
}

int QTransform_Type(void* ptr){
	return static_cast<QTransform*>(ptr)->type();
}

