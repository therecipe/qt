#include "qgraphicsscenemouseevent.h"
#include <QModelIndex>
#include <QGraphicsScene>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QGraphicsSceneMouseEvent>
#include "_cgo_export.h"

class MyQGraphicsSceneMouseEvent: public QGraphicsSceneMouseEvent {
public:
};

int QGraphicsSceneMouseEvent_Button(void* ptr){
	return static_cast<QGraphicsSceneMouseEvent*>(ptr)->button();
}

int QGraphicsSceneMouseEvent_Buttons(void* ptr){
	return static_cast<QGraphicsSceneMouseEvent*>(ptr)->buttons();
}

int QGraphicsSceneMouseEvent_Flags(void* ptr){
	return static_cast<QGraphicsSceneMouseEvent*>(ptr)->flags();
}

int QGraphicsSceneMouseEvent_Modifiers(void* ptr){
	return static_cast<QGraphicsSceneMouseEvent*>(ptr)->modifiers();
}

int QGraphicsSceneMouseEvent_Source(void* ptr){
	return static_cast<QGraphicsSceneMouseEvent*>(ptr)->source();
}

void QGraphicsSceneMouseEvent_DestroyQGraphicsSceneMouseEvent(void* ptr){
	static_cast<QGraphicsSceneMouseEvent*>(ptr)->~QGraphicsSceneMouseEvent();
}

