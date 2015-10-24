#include "qenterevent.h"
#include <QModelIndex>
#include <QPointF>
#include <QPoint>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QEnterEvent>
#include "_cgo_export.h"

class MyQEnterEvent: public QEnterEvent {
public:
};

QtObjectPtr QEnterEvent_NewQEnterEvent(QtObjectPtr localPos, QtObjectPtr windowPos, QtObjectPtr screenPos){
	return new QEnterEvent(*static_cast<QPointF*>(localPos), *static_cast<QPointF*>(windowPos), *static_cast<QPointF*>(screenPos));
}

int QEnterEvent_GlobalX(QtObjectPtr ptr){
	return static_cast<QEnterEvent*>(ptr)->globalX();
}

int QEnterEvent_GlobalY(QtObjectPtr ptr){
	return static_cast<QEnterEvent*>(ptr)->globalY();
}

int QEnterEvent_X(QtObjectPtr ptr){
	return static_cast<QEnterEvent*>(ptr)->x();
}

int QEnterEvent_Y(QtObjectPtr ptr){
	return static_cast<QEnterEvent*>(ptr)->y();
}

