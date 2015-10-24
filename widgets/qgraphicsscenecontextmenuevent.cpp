#include "qgraphicsscenecontextmenuevent.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGraphicsScene>
#include <QGraphicsSceneContextMenuEvent>
#include "_cgo_export.h"

class MyQGraphicsSceneContextMenuEvent: public QGraphicsSceneContextMenuEvent {
public:
};

int QGraphicsSceneContextMenuEvent_Modifiers(QtObjectPtr ptr){
	return static_cast<QGraphicsSceneContextMenuEvent*>(ptr)->modifiers();
}

int QGraphicsSceneContextMenuEvent_Reason(QtObjectPtr ptr){
	return static_cast<QGraphicsSceneContextMenuEvent*>(ptr)->reason();
}

void QGraphicsSceneContextMenuEvent_DestroyQGraphicsSceneContextMenuEvent(QtObjectPtr ptr){
	static_cast<QGraphicsSceneContextMenuEvent*>(ptr)->~QGraphicsSceneContextMenuEvent();
}

