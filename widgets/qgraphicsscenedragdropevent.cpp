#include "qgraphicsscenedragdropevent.h"
#include <QGraphicsScene>
#include <QMimeData>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGraphicsSceneDragDropEvent>
#include "_cgo_export.h"

class MyQGraphicsSceneDragDropEvent: public QGraphicsSceneDragDropEvent {
public:
};

void QGraphicsSceneDragDropEvent_AcceptProposedAction(QtObjectPtr ptr){
	static_cast<QGraphicsSceneDragDropEvent*>(ptr)->acceptProposedAction();
}

int QGraphicsSceneDragDropEvent_Buttons(QtObjectPtr ptr){
	return static_cast<QGraphicsSceneDragDropEvent*>(ptr)->buttons();
}

int QGraphicsSceneDragDropEvent_DropAction(QtObjectPtr ptr){
	return static_cast<QGraphicsSceneDragDropEvent*>(ptr)->dropAction();
}

QtObjectPtr QGraphicsSceneDragDropEvent_MimeData(QtObjectPtr ptr){
	return const_cast<QMimeData*>(static_cast<QGraphicsSceneDragDropEvent*>(ptr)->mimeData());
}

int QGraphicsSceneDragDropEvent_Modifiers(QtObjectPtr ptr){
	return static_cast<QGraphicsSceneDragDropEvent*>(ptr)->modifiers();
}

int QGraphicsSceneDragDropEvent_PossibleActions(QtObjectPtr ptr){
	return static_cast<QGraphicsSceneDragDropEvent*>(ptr)->possibleActions();
}

int QGraphicsSceneDragDropEvent_ProposedAction(QtObjectPtr ptr){
	return static_cast<QGraphicsSceneDragDropEvent*>(ptr)->proposedAction();
}

void QGraphicsSceneDragDropEvent_SetDropAction(QtObjectPtr ptr, int action){
	static_cast<QGraphicsSceneDragDropEvent*>(ptr)->setDropAction(static_cast<Qt::DropAction>(action));
}

QtObjectPtr QGraphicsSceneDragDropEvent_Source(QtObjectPtr ptr){
	return static_cast<QGraphicsSceneDragDropEvent*>(ptr)->source();
}

void QGraphicsSceneDragDropEvent_DestroyQGraphicsSceneDragDropEvent(QtObjectPtr ptr){
	static_cast<QGraphicsSceneDragDropEvent*>(ptr)->~QGraphicsSceneDragDropEvent();
}

