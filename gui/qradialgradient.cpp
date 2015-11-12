#include "qradialgradient.h"
#include <QModelIndex>
#include <QPoint>
#include <QPointF>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QRadialGradient>
#include "_cgo_export.h"

class MyQRadialGradient: public QRadialGradient {
public:
};

void* QRadialGradient_NewQRadialGradient(){
	return new QRadialGradient();
}

void* QRadialGradient_NewQRadialGradient6(void* center, double centerRadius, void* focalPoint, double focalRadius){
	return new QRadialGradient(*static_cast<QPointF*>(center), static_cast<qreal>(centerRadius), *static_cast<QPointF*>(focalPoint), static_cast<qreal>(focalRadius));
}

void* QRadialGradient_NewQRadialGradient4(void* center, double radius){
	return new QRadialGradient(*static_cast<QPointF*>(center), static_cast<qreal>(radius));
}

void* QRadialGradient_NewQRadialGradient2(void* center, double radius, void* focalPoint){
	return new QRadialGradient(*static_cast<QPointF*>(center), static_cast<qreal>(radius), *static_cast<QPointF*>(focalPoint));
}

void* QRadialGradient_NewQRadialGradient7(double cx, double cy, double centerRadius, double fx, double fy, double focalRadius){
	return new QRadialGradient(static_cast<qreal>(cx), static_cast<qreal>(cy), static_cast<qreal>(centerRadius), static_cast<qreal>(fx), static_cast<qreal>(fy), static_cast<qreal>(focalRadius));
}

void* QRadialGradient_NewQRadialGradient5(double cx, double cy, double radius){
	return new QRadialGradient(static_cast<qreal>(cx), static_cast<qreal>(cy), static_cast<qreal>(radius));
}

void* QRadialGradient_NewQRadialGradient3(double cx, double cy, double radius, double fx, double fy){
	return new QRadialGradient(static_cast<qreal>(cx), static_cast<qreal>(cy), static_cast<qreal>(radius), static_cast<qreal>(fx), static_cast<qreal>(fy));
}

double QRadialGradient_CenterRadius(void* ptr){
	return static_cast<double>(static_cast<QRadialGradient*>(ptr)->centerRadius());
}

double QRadialGradient_FocalRadius(void* ptr){
	return static_cast<double>(static_cast<QRadialGradient*>(ptr)->focalRadius());
}

double QRadialGradient_Radius(void* ptr){
	return static_cast<double>(static_cast<QRadialGradient*>(ptr)->radius());
}

void QRadialGradient_SetCenter(void* ptr, void* center){
	static_cast<QRadialGradient*>(ptr)->setCenter(*static_cast<QPointF*>(center));
}

void QRadialGradient_SetCenter2(void* ptr, double x, double y){
	static_cast<QRadialGradient*>(ptr)->setCenter(static_cast<qreal>(x), static_cast<qreal>(y));
}

void QRadialGradient_SetCenterRadius(void* ptr, double radius){
	static_cast<QRadialGradient*>(ptr)->setCenterRadius(static_cast<qreal>(radius));
}

void QRadialGradient_SetFocalPoint(void* ptr, void* focalPoint){
	static_cast<QRadialGradient*>(ptr)->setFocalPoint(*static_cast<QPointF*>(focalPoint));
}

void QRadialGradient_SetFocalPoint2(void* ptr, double x, double y){
	static_cast<QRadialGradient*>(ptr)->setFocalPoint(static_cast<qreal>(x), static_cast<qreal>(y));
}

void QRadialGradient_SetFocalRadius(void* ptr, double radius){
	static_cast<QRadialGradient*>(ptr)->setFocalRadius(static_cast<qreal>(radius));
}

void QRadialGradient_SetRadius(void* ptr, double radius){
	static_cast<QRadialGradient*>(ptr)->setRadius(static_cast<qreal>(radius));
}

