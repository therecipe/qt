#include "qeasingcurve.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPointF>
#include <QPoint>
#include <QEasingCurve>
#include "_cgo_export.h"

class MyQEasingCurve: public QEasingCurve {
public:
};

QtObjectPtr QEasingCurve_NewQEasingCurve3(QtObjectPtr other){
	return new QEasingCurve(*static_cast<QEasingCurve*>(other));
}

QtObjectPtr QEasingCurve_NewQEasingCurve(int ty){
	return new QEasingCurve(static_cast<QEasingCurve::Type>(ty));
}

QtObjectPtr QEasingCurve_NewQEasingCurve2(QtObjectPtr other){
	return new QEasingCurve(*static_cast<QEasingCurve*>(other));
}

void QEasingCurve_AddCubicBezierSegment(QtObjectPtr ptr, QtObjectPtr c1, QtObjectPtr c2, QtObjectPtr endPoint){
	static_cast<QEasingCurve*>(ptr)->addCubicBezierSegment(*static_cast<QPointF*>(c1), *static_cast<QPointF*>(c2), *static_cast<QPointF*>(endPoint));
}

void QEasingCurve_SetType(QtObjectPtr ptr, int ty){
	static_cast<QEasingCurve*>(ptr)->setType(static_cast<QEasingCurve::Type>(ty));
}

void QEasingCurve_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QEasingCurve*>(ptr)->swap(*static_cast<QEasingCurve*>(other));
}

int QEasingCurve_Type(QtObjectPtr ptr){
	return static_cast<QEasingCurve*>(ptr)->type();
}

void QEasingCurve_DestroyQEasingCurve(QtObjectPtr ptr){
	static_cast<QEasingCurve*>(ptr)->~QEasingCurve();
}

