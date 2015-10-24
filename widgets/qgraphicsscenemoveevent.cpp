#include "qgraphicsscenemoveevent.h"
#include <QGraphicsScene>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGraphicsSceneMoveEvent>
#include "_cgo_export.h"

class MyQGraphicsSceneMoveEvent: public QGraphicsSceneMoveEvent {
public:
};

QtObjectPtr QGraphicsSceneMoveEvent_NewQGraphicsSceneMoveEvent(){
	return new QGraphicsSceneMoveEvent();
}

void QGraphicsSceneMoveEvent_DestroyQGraphicsSceneMoveEvent(QtObjectPtr ptr){
	static_cast<QGraphicsSceneMoveEvent*>(ptr)->~QGraphicsSceneMoveEvent();
}

