#include "qenterevent.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPoint>
#include <QPointF>
#include <QEnterEvent>
#include "_cgo_export.h"

class MyQEnterEvent: public QEnterEvent {
public:
};

void* QEnterEvent_NewQEnterEvent(void* localPos, void* windowPos, void* screenPos){
	return new QEnterEvent(*static_cast<QPointF*>(localPos), *static_cast<QPointF*>(windowPos), *static_cast<QPointF*>(screenPos));
}

int QEnterEvent_GlobalX(void* ptr){
	return static_cast<QEnterEvent*>(ptr)->globalX();
}

int QEnterEvent_GlobalY(void* ptr){
	return static_cast<QEnterEvent*>(ptr)->globalY();
}

int QEnterEvent_X(void* ptr){
	return static_cast<QEnterEvent*>(ptr)->x();
}

int QEnterEvent_Y(void* ptr){
	return static_cast<QEnterEvent*>(ptr)->y();
}

