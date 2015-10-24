#include "qlineargradient.h"
#include <QModelIndex>
#include <QPoint>
#include <QLine>
#include <QPointF>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QLinearGradient>
#include "_cgo_export.h"

class MyQLinearGradient: public QLinearGradient {
public:
};

QtObjectPtr QLinearGradient_NewQLinearGradient(){
	return new QLinearGradient();
}

QtObjectPtr QLinearGradient_NewQLinearGradient2(QtObjectPtr start, QtObjectPtr finalStop){
	return new QLinearGradient(*static_cast<QPointF*>(start), *static_cast<QPointF*>(finalStop));
}

void QLinearGradient_SetFinalStop(QtObjectPtr ptr, QtObjectPtr stop){
	static_cast<QLinearGradient*>(ptr)->setFinalStop(*static_cast<QPointF*>(stop));
}

void QLinearGradient_SetStart(QtObjectPtr ptr, QtObjectPtr start){
	static_cast<QLinearGradient*>(ptr)->setStart(*static_cast<QPointF*>(start));
}

