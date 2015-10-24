#include "qdropevent.h"
#include <QMimeData>
#include <QEvent>
#include <QPointF>
#include <QPoint>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDropEvent>
#include "_cgo_export.h"

class MyQDropEvent: public QDropEvent {
public:
};

void QDropEvent_SetDropAction(QtObjectPtr ptr, int action){
	static_cast<QDropEvent*>(ptr)->setDropAction(static_cast<Qt::DropAction>(action));
}

QtObjectPtr QDropEvent_NewQDropEvent(QtObjectPtr pos, int actions, QtObjectPtr data, int buttons, int modifiers, int ty){
	return new QDropEvent(*static_cast<QPointF*>(pos), static_cast<Qt::DropAction>(actions), static_cast<QMimeData*>(data), static_cast<Qt::MouseButton>(buttons), static_cast<Qt::KeyboardModifier>(modifiers), static_cast<QEvent::Type>(ty));
}

void QDropEvent_AcceptProposedAction(QtObjectPtr ptr){
	static_cast<QDropEvent*>(ptr)->acceptProposedAction();
}

int QDropEvent_DropAction(QtObjectPtr ptr){
	return static_cast<QDropEvent*>(ptr)->dropAction();
}

int QDropEvent_KeyboardModifiers(QtObjectPtr ptr){
	return static_cast<QDropEvent*>(ptr)->keyboardModifiers();
}

QtObjectPtr QDropEvent_MimeData(QtObjectPtr ptr){
	return const_cast<QMimeData*>(static_cast<QDropEvent*>(ptr)->mimeData());
}

int QDropEvent_MouseButtons(QtObjectPtr ptr){
	return static_cast<QDropEvent*>(ptr)->mouseButtons();
}

int QDropEvent_PossibleActions(QtObjectPtr ptr){
	return static_cast<QDropEvent*>(ptr)->possibleActions();
}

int QDropEvent_ProposedAction(QtObjectPtr ptr){
	return static_cast<QDropEvent*>(ptr)->proposedAction();
}

QtObjectPtr QDropEvent_Source(QtObjectPtr ptr){
	return static_cast<QDropEvent*>(ptr)->source();
}

