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

QtObjectPtr QGraphicsSceneEvent_Widget(QtObjectPtr ptr){
	return static_cast<QGraphicsSceneEvent*>(ptr)->widget();
}

void QGraphicsSceneEvent_DestroyQGraphicsSceneEvent(QtObjectPtr ptr){
	static_cast<QGraphicsSceneEvent*>(ptr)->~QGraphicsSceneEvent();
}

