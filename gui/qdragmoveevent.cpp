#include "qdragmoveevent.h"
#include <QString>
#include <QUrl>
#include <QModelIndex>
#include <QMimeData>
#include <QDrag>
#include <QEvent>
#include <QVariant>
#include <QRect>
#include <QPoint>
#include <QDragMoveEvent>
#include "_cgo_export.h"

class MyQDragMoveEvent: public QDragMoveEvent {
public:
};

QtObjectPtr QDragMoveEvent_NewQDragMoveEvent(QtObjectPtr pos, int actions, QtObjectPtr data, int buttons, int modifiers, int ty){
	return new QDragMoveEvent(*static_cast<QPoint*>(pos), static_cast<Qt::DropAction>(actions), static_cast<QMimeData*>(data), static_cast<Qt::MouseButton>(buttons), static_cast<Qt::KeyboardModifier>(modifiers), static_cast<QEvent::Type>(ty));
}

void QDragMoveEvent_Accept2(QtObjectPtr ptr){
	static_cast<QDragMoveEvent*>(ptr)->accept();
}

void QDragMoveEvent_Accept(QtObjectPtr ptr, QtObjectPtr rectangle){
	static_cast<QDragMoveEvent*>(ptr)->accept(*static_cast<QRect*>(rectangle));
}

void QDragMoveEvent_Ignore2(QtObjectPtr ptr){
	static_cast<QDragMoveEvent*>(ptr)->ignore();
}

void QDragMoveEvent_Ignore(QtObjectPtr ptr, QtObjectPtr rectangle){
	static_cast<QDragMoveEvent*>(ptr)->ignore(*static_cast<QRect*>(rectangle));
}

void QDragMoveEvent_DestroyQDragMoveEvent(QtObjectPtr ptr){
	static_cast<QDragMoveEvent*>(ptr)->~QDragMoveEvent();
}

