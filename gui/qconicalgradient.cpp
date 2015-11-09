#include "qconicalgradient.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPoint>
#include <QPointF>
#include <QConicalGradient>
#include "_cgo_export.h"

class MyQConicalGradient: public QConicalGradient {
public:
};

void* QConicalGradient_NewQConicalGradient(){
	return new QConicalGradient();
}

void* QConicalGradient_NewQConicalGradient2(void* center, double angle){
	return new QConicalGradient(*static_cast<QPointF*>(center), static_cast<qreal>(angle));
}

void* QConicalGradient_NewQConicalGradient3(double cx, double cy, double angle){
	return new QConicalGradient(static_cast<qreal>(cx), static_cast<qreal>(cy), static_cast<qreal>(angle));
}

double QConicalGradient_Angle(void* ptr){
	return static_cast<double>(static_cast<QConicalGradient*>(ptr)->angle());
}

void QConicalGradient_SetAngle(void* ptr, double angle){
	static_cast<QConicalGradient*>(ptr)->setAngle(static_cast<qreal>(angle));
}

void QConicalGradient_SetCenter(void* ptr, void* center){
	static_cast<QConicalGradient*>(ptr)->setCenter(*static_cast<QPointF*>(center));
}

void QConicalGradient_SetCenter2(void* ptr, double x, double y){
	static_cast<QConicalGradient*>(ptr)->setCenter(static_cast<qreal>(x), static_cast<qreal>(y));
}

