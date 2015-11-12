#include "qgraphicssceneresizeevent.h"
#include <QUrl>
#include <QModelIndex>
#include <QGraphicsScene>
#include <QString>
#include <QVariant>
#include <QGraphicsSceneResizeEvent>
#include "_cgo_export.h"

class MyQGraphicsSceneResizeEvent: public QGraphicsSceneResizeEvent {
public:
};

void* QGraphicsSceneResizeEvent_NewQGraphicsSceneResizeEvent(){
	return new QGraphicsSceneResizeEvent();
}

void QGraphicsSceneResizeEvent_DestroyQGraphicsSceneResizeEvent(void* ptr){
	static_cast<QGraphicsSceneResizeEvent*>(ptr)->~QGraphicsSceneResizeEvent();
}

