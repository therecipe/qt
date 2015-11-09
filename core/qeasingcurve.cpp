#include "qeasingcurve.h"
#include <QPointF>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPoint>
#include <QEasingCurve>
#include "_cgo_export.h"

class MyQEasingCurve: public QEasingCurve {
public:
};

void* QEasingCurve_NewQEasingCurve3(void* other){
	return new QEasingCurve(*static_cast<QEasingCurve*>(other));
}

void* QEasingCurve_NewQEasingCurve(int ty){
	return new QEasingCurve(static_cast<QEasingCurve::Type>(ty));
}

void* QEasingCurve_NewQEasingCurve2(void* other){
	return new QEasingCurve(*static_cast<QEasingCurve*>(other));
}

void QEasingCurve_AddCubicBezierSegment(void* ptr, void* c1, void* c2, void* endPoint){
	static_cast<QEasingCurve*>(ptr)->addCubicBezierSegment(*static_cast<QPointF*>(c1), *static_cast<QPointF*>(c2), *static_cast<QPointF*>(endPoint));
}

void QEasingCurve_AddTCBSegment(void* ptr, void* nextPoint, double t, double c, double b){
	static_cast<QEasingCurve*>(ptr)->addTCBSegment(*static_cast<QPointF*>(nextPoint), static_cast<qreal>(t), static_cast<qreal>(c), static_cast<qreal>(b));
}

double QEasingCurve_Amplitude(void* ptr){
	return static_cast<double>(static_cast<QEasingCurve*>(ptr)->amplitude());
}

double QEasingCurve_Overshoot(void* ptr){
	return static_cast<double>(static_cast<QEasingCurve*>(ptr)->overshoot());
}

double QEasingCurve_Period(void* ptr){
	return static_cast<double>(static_cast<QEasingCurve*>(ptr)->period());
}

void QEasingCurve_SetAmplitude(void* ptr, double amplitude){
	static_cast<QEasingCurve*>(ptr)->setAmplitude(static_cast<qreal>(amplitude));
}

void QEasingCurve_SetOvershoot(void* ptr, double overshoot){
	static_cast<QEasingCurve*>(ptr)->setOvershoot(static_cast<qreal>(overshoot));
}

void QEasingCurve_SetPeriod(void* ptr, double period){
	static_cast<QEasingCurve*>(ptr)->setPeriod(static_cast<qreal>(period));
}

void QEasingCurve_SetType(void* ptr, int ty){
	static_cast<QEasingCurve*>(ptr)->setType(static_cast<QEasingCurve::Type>(ty));
}

void QEasingCurve_Swap(void* ptr, void* other){
	static_cast<QEasingCurve*>(ptr)->swap(*static_cast<QEasingCurve*>(other));
}

int QEasingCurve_Type(void* ptr){
	return static_cast<QEasingCurve*>(ptr)->type();
}

double QEasingCurve_ValueForProgress(void* ptr, double progress){
	return static_cast<double>(static_cast<QEasingCurve*>(ptr)->valueForProgress(static_cast<qreal>(progress)));
}

void QEasingCurve_DestroyQEasingCurve(void* ptr){
	static_cast<QEasingCurve*>(ptr)->~QEasingCurve();
}

