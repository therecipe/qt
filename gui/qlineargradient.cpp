#include "qlineargradient.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QLine>
#include <QPoint>
#include <QPointF>
#include <QString>
#include <QLinearGradient>
#include "_cgo_export.h"

class MyQLinearGradient: public QLinearGradient {
public:
};

void* QLinearGradient_NewQLinearGradient3(double x1, double y1, double x2, double y2){
	return new QLinearGradient(static_cast<qreal>(x1), static_cast<qreal>(y1), static_cast<qreal>(x2), static_cast<qreal>(y2));
}

void* QLinearGradient_NewQLinearGradient(){
	return new QLinearGradient();
}

void* QLinearGradient_NewQLinearGradient2(void* start, void* finalStop){
	return new QLinearGradient(*static_cast<QPointF*>(start), *static_cast<QPointF*>(finalStop));
}

void QLinearGradient_SetFinalStop(void* ptr, void* stop){
	static_cast<QLinearGradient*>(ptr)->setFinalStop(*static_cast<QPointF*>(stop));
}

void QLinearGradient_SetFinalStop2(void* ptr, double x, double y){
	static_cast<QLinearGradient*>(ptr)->setFinalStop(static_cast<qreal>(x), static_cast<qreal>(y));
}

void QLinearGradient_SetStart(void* ptr, void* start){
	static_cast<QLinearGradient*>(ptr)->setStart(*static_cast<QPointF*>(start));
}

void QLinearGradient_SetStart2(void* ptr, double x, double y){
	static_cast<QLinearGradient*>(ptr)->setStart(static_cast<qreal>(x), static_cast<qreal>(y));
}

