#include "qgraphicsscenehoverevent.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGraphicsScene>
#include <QString>
#include <QGraphicsSceneHoverEvent>
#include "_cgo_export.h"

class MyQGraphicsSceneHoverEvent: public QGraphicsSceneHoverEvent {
public:
};

int QGraphicsSceneHoverEvent_Modifiers(void* ptr){
	return static_cast<QGraphicsSceneHoverEvent*>(ptr)->modifiers();
}

void QGraphicsSceneHoverEvent_DestroyQGraphicsSceneHoverEvent(void* ptr){
	static_cast<QGraphicsSceneHoverEvent*>(ptr)->~QGraphicsSceneHoverEvent();
}

