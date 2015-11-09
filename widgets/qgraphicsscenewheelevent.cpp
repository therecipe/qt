#include "qgraphicsscenewheelevent.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGraphicsScene>
#include <QGraphicsSceneWheelEvent>
#include "_cgo_export.h"

class MyQGraphicsSceneWheelEvent: public QGraphicsSceneWheelEvent {
public:
};

int QGraphicsSceneWheelEvent_Buttons(void* ptr){
	return static_cast<QGraphicsSceneWheelEvent*>(ptr)->buttons();
}

int QGraphicsSceneWheelEvent_Delta(void* ptr){
	return static_cast<QGraphicsSceneWheelEvent*>(ptr)->delta();
}

int QGraphicsSceneWheelEvent_Modifiers(void* ptr){
	return static_cast<QGraphicsSceneWheelEvent*>(ptr)->modifiers();
}

int QGraphicsSceneWheelEvent_Orientation(void* ptr){
	return static_cast<QGraphicsSceneWheelEvent*>(ptr)->orientation();
}

void QGraphicsSceneWheelEvent_DestroyQGraphicsSceneWheelEvent(void* ptr){
	static_cast<QGraphicsSceneWheelEvent*>(ptr)->~QGraphicsSceneWheelEvent();
}

