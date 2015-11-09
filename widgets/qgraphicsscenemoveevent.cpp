#include "qgraphicsscenemoveevent.h"
#include <QUrl>
#include <QModelIndex>
#include <QGraphicsScene>
#include <QString>
#include <QVariant>
#include <QGraphicsSceneMoveEvent>
#include "_cgo_export.h"

class MyQGraphicsSceneMoveEvent: public QGraphicsSceneMoveEvent {
public:
};

void* QGraphicsSceneMoveEvent_NewQGraphicsSceneMoveEvent(){
	return new QGraphicsSceneMoveEvent();
}

void QGraphicsSceneMoveEvent_DestroyQGraphicsSceneMoveEvent(void* ptr){
	static_cast<QGraphicsSceneMoveEvent*>(ptr)->~QGraphicsSceneMoveEvent();
}

