#include "qgraphicsscenecontextmenuevent.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGraphicsScene>
#include <QString>
#include <QGraphicsSceneContextMenuEvent>
#include "_cgo_export.h"

class MyQGraphicsSceneContextMenuEvent: public QGraphicsSceneContextMenuEvent {
public:
};

int QGraphicsSceneContextMenuEvent_Modifiers(void* ptr){
	return static_cast<QGraphicsSceneContextMenuEvent*>(ptr)->modifiers();
}

int QGraphicsSceneContextMenuEvent_Reason(void* ptr){
	return static_cast<QGraphicsSceneContextMenuEvent*>(ptr)->reason();
}

void QGraphicsSceneContextMenuEvent_DestroyQGraphicsSceneContextMenuEvent(void* ptr){
	static_cast<QGraphicsSceneContextMenuEvent*>(ptr)->~QGraphicsSceneContextMenuEvent();
}

