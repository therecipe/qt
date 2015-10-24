#include "qscrollevent.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPointF>
#include <QPoint>
#include <QScrollEvent>
#include "_cgo_export.h"

class MyQScrollEvent: public QScrollEvent {
public:
};

QtObjectPtr QScrollEvent_NewQScrollEvent(QtObjectPtr contentPos, QtObjectPtr overshootDistance, int scrollState){
	return new QScrollEvent(*static_cast<QPointF*>(contentPos), *static_cast<QPointF*>(overshootDistance), static_cast<QScrollEvent::ScrollState>(scrollState));
}

int QScrollEvent_ScrollState(QtObjectPtr ptr){
	return static_cast<QScrollEvent*>(ptr)->scrollState();
}

void QScrollEvent_DestroyQScrollEvent(QtObjectPtr ptr){
	static_cast<QScrollEvent*>(ptr)->~QScrollEvent();
}

