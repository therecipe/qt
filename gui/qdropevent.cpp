#include "qdropevent.h"
#include <QMimeData>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPoint>
#include <QPointF>
#include <QEvent>
#include <QDropEvent>
#include "_cgo_export.h"

class MyQDropEvent: public QDropEvent {
public:
};

void QDropEvent_SetDropAction(void* ptr, int action){
	static_cast<QDropEvent*>(ptr)->setDropAction(static_cast<Qt::DropAction>(action));
}

void* QDropEvent_NewQDropEvent(void* pos, int actions, void* data, int buttons, int modifiers, int ty){
	return new QDropEvent(*static_cast<QPointF*>(pos), static_cast<Qt::DropAction>(actions), static_cast<QMimeData*>(data), static_cast<Qt::MouseButton>(buttons), static_cast<Qt::KeyboardModifier>(modifiers), static_cast<QEvent::Type>(ty));
}

void QDropEvent_AcceptProposedAction(void* ptr){
	static_cast<QDropEvent*>(ptr)->acceptProposedAction();
}

int QDropEvent_DropAction(void* ptr){
	return static_cast<QDropEvent*>(ptr)->dropAction();
}

int QDropEvent_KeyboardModifiers(void* ptr){
	return static_cast<QDropEvent*>(ptr)->keyboardModifiers();
}

void* QDropEvent_MimeData(void* ptr){
	return const_cast<QMimeData*>(static_cast<QDropEvent*>(ptr)->mimeData());
}

int QDropEvent_MouseButtons(void* ptr){
	return static_cast<QDropEvent*>(ptr)->mouseButtons();
}

int QDropEvent_PossibleActions(void* ptr){
	return static_cast<QDropEvent*>(ptr)->possibleActions();
}

int QDropEvent_ProposedAction(void* ptr){
	return static_cast<QDropEvent*>(ptr)->proposedAction();
}

void* QDropEvent_Source(void* ptr){
	return static_cast<QDropEvent*>(ptr)->source();
}

