#include "qgraphicssceneevent.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGraphicsScene>
#include <QGraphicsSceneEvent>
#include "_cgo_export.h"

class MyQGraphicsSceneEvent: public QGraphicsSceneEvent {
public:
};

void* QGraphicsSceneEvent_Widget(void* ptr){
	return static_cast<QGraphicsSceneEvent*>(ptr)->widget();
}

void QGraphicsSceneEvent_DestroyQGraphicsSceneEvent(void* ptr){
	static_cast<QGraphicsSceneEvent*>(ptr)->~QGraphicsSceneEvent();
}

