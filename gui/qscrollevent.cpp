#include "qscrollevent.h"
#include <QUrl>
#include <QModelIndex>
#include <QPointF>
#include <QPoint>
#include <QString>
#include <QVariant>
#include <QScrollEvent>
#include "_cgo_export.h"

class MyQScrollEvent: public QScrollEvent {
public:
};

void* QScrollEvent_NewQScrollEvent(void* contentPos, void* overshootDistance, int scrollState){
	return new QScrollEvent(*static_cast<QPointF*>(contentPos), *static_cast<QPointF*>(overshootDistance), static_cast<QScrollEvent::ScrollState>(scrollState));
}

int QScrollEvent_ScrollState(void* ptr){
	return static_cast<QScrollEvent*>(ptr)->scrollState();
}

void QScrollEvent_DestroyQScrollEvent(void* ptr){
	static_cast<QScrollEvent*>(ptr)->~QScrollEvent();
}

