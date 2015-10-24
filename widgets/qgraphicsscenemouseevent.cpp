#include "qgraphicsscenemouseevent.h"
#include <QUrl>
#include <QModelIndex>
#include <QGraphicsScene>
#include <QString>
#include <QVariant>
#include <QGraphicsSceneMouseEvent>
#include "_cgo_export.h"

class MyQGraphicsSceneMouseEvent: public QGraphicsSceneMouseEvent {
public:
};

int QGraphicsSceneMouseEvent_Button(QtObjectPtr ptr){
	return static_cast<QGraphicsSceneMouseEvent*>(ptr)->button();
}

int QGraphicsSceneMouseEvent_Buttons(QtObjectPtr ptr){
	return static_cast<QGraphicsSceneMouseEvent*>(ptr)->buttons();
}

int QGraphicsSceneMouseEvent_Flags(QtObjectPtr ptr){
	return static_cast<QGraphicsSceneMouseEvent*>(ptr)->flags();
}

int QGraphicsSceneMouseEvent_Modifiers(QtObjectPtr ptr){
	return static_cast<QGraphicsSceneMouseEvent*>(ptr)->modifiers();
}

int QGraphicsSceneMouseEvent_Source(QtObjectPtr ptr){
	return static_cast<QGraphicsSceneMouseEvent*>(ptr)->source();
}

void QGraphicsSceneMouseEvent_DestroyQGraphicsSceneMouseEvent(QtObjectPtr ptr){
	static_cast<QGraphicsSceneMouseEvent*>(ptr)->~QGraphicsSceneMouseEvent();
}

