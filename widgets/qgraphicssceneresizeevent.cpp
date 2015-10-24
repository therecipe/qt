#include "qgraphicssceneresizeevent.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGraphicsScene>
#include <QGraphicsSceneResizeEvent>
#include "_cgo_export.h"

class MyQGraphicsSceneResizeEvent: public QGraphicsSceneResizeEvent {
public:
};

QtObjectPtr QGraphicsSceneResizeEvent_NewQGraphicsSceneResizeEvent(){
	return new QGraphicsSceneResizeEvent();
}

void QGraphicsSceneResizeEvent_DestroyQGraphicsSceneResizeEvent(QtObjectPtr ptr){
	static_cast<QGraphicsSceneResizeEvent*>(ptr)->~QGraphicsSceneResizeEvent();
}

