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

int QGraphicsSceneWheelEvent_Buttons(QtObjectPtr ptr){
	return static_cast<QGraphicsSceneWheelEvent*>(ptr)->buttons();
}

int QGraphicsSceneWheelEvent_Delta(QtObjectPtr ptr){
	return static_cast<QGraphicsSceneWheelEvent*>(ptr)->delta();
}

int QGraphicsSceneWheelEvent_Modifiers(QtObjectPtr ptr){
	return static_cast<QGraphicsSceneWheelEvent*>(ptr)->modifiers();
}

int QGraphicsSceneWheelEvent_Orientation(QtObjectPtr ptr){
	return static_cast<QGraphicsSceneWheelEvent*>(ptr)->orientation();
}

void QGraphicsSceneWheelEvent_DestroyQGraphicsSceneWheelEvent(QtObjectPtr ptr){
	static_cast<QGraphicsSceneWheelEvent*>(ptr)->~QGraphicsSceneWheelEvent();
}

