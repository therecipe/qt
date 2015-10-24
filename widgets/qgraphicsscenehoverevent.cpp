#include "qgraphicsscenehoverevent.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGraphicsScene>
#include <QGraphicsSceneHoverEvent>
#include "_cgo_export.h"

class MyQGraphicsSceneHoverEvent: public QGraphicsSceneHoverEvent {
public:
};

int QGraphicsSceneHoverEvent_Modifiers(QtObjectPtr ptr){
	return static_cast<QGraphicsSceneHoverEvent*>(ptr)->modifiers();
}

void QGraphicsSceneHoverEvent_DestroyQGraphicsSceneHoverEvent(QtObjectPtr ptr){
	static_cast<QGraphicsSceneHoverEvent*>(ptr)->~QGraphicsSceneHoverEvent();
}

