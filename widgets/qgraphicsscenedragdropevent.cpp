#include "qgraphicsscenedragdropevent.h"
#include <QModelIndex>
#include <QMimeData>
#include <QGraphicsScene>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QGraphicsSceneDragDropEvent>
#include "_cgo_export.h"

class MyQGraphicsSceneDragDropEvent: public QGraphicsSceneDragDropEvent {
public:
};

void QGraphicsSceneDragDropEvent_AcceptProposedAction(void* ptr){
	static_cast<QGraphicsSceneDragDropEvent*>(ptr)->acceptProposedAction();
}

int QGraphicsSceneDragDropEvent_Buttons(void* ptr){
	return static_cast<QGraphicsSceneDragDropEvent*>(ptr)->buttons();
}

int QGraphicsSceneDragDropEvent_DropAction(void* ptr){
	return static_cast<QGraphicsSceneDragDropEvent*>(ptr)->dropAction();
}

void* QGraphicsSceneDragDropEvent_MimeData(void* ptr){
	return const_cast<QMimeData*>(static_cast<QGraphicsSceneDragDropEvent*>(ptr)->mimeData());
}

int QGraphicsSceneDragDropEvent_Modifiers(void* ptr){
	return static_cast<QGraphicsSceneDragDropEvent*>(ptr)->modifiers();
}

int QGraphicsSceneDragDropEvent_PossibleActions(void* ptr){
	return static_cast<QGraphicsSceneDragDropEvent*>(ptr)->possibleActions();
}

int QGraphicsSceneDragDropEvent_ProposedAction(void* ptr){
	return static_cast<QGraphicsSceneDragDropEvent*>(ptr)->proposedAction();
}

void QGraphicsSceneDragDropEvent_SetDropAction(void* ptr, int action){
	static_cast<QGraphicsSceneDragDropEvent*>(ptr)->setDropAction(static_cast<Qt::DropAction>(action));
}

void* QGraphicsSceneDragDropEvent_Source(void* ptr){
	return static_cast<QGraphicsSceneDragDropEvent*>(ptr)->source();
}

void QGraphicsSceneDragDropEvent_DestroyQGraphicsSceneDragDropEvent(void* ptr){
	static_cast<QGraphicsSceneDragDropEvent*>(ptr)->~QGraphicsSceneDragDropEvent();
}

