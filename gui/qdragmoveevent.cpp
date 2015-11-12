#include "qdragmoveevent.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QDrag>
#include <QRect>
#include <QMimeData>
#include <QModelIndex>
#include <QPoint>
#include <QEvent>
#include <QDragMoveEvent>
#include "_cgo_export.h"

class MyQDragMoveEvent: public QDragMoveEvent {
public:
};

void* QDragMoveEvent_NewQDragMoveEvent(void* pos, int actions, void* data, int buttons, int modifiers, int ty){
	return new QDragMoveEvent(*static_cast<QPoint*>(pos), static_cast<Qt::DropAction>(actions), static_cast<QMimeData*>(data), static_cast<Qt::MouseButton>(buttons), static_cast<Qt::KeyboardModifier>(modifiers), static_cast<QEvent::Type>(ty));
}

void QDragMoveEvent_Accept2(void* ptr){
	static_cast<QDragMoveEvent*>(ptr)->accept();
}

void QDragMoveEvent_Accept(void* ptr, void* rectangle){
	static_cast<QDragMoveEvent*>(ptr)->accept(*static_cast<QRect*>(rectangle));
}

void QDragMoveEvent_Ignore2(void* ptr){
	static_cast<QDragMoveEvent*>(ptr)->ignore();
}

void QDragMoveEvent_Ignore(void* ptr, void* rectangle){
	static_cast<QDragMoveEvent*>(ptr)->ignore(*static_cast<QRect*>(rectangle));
}

void QDragMoveEvent_DestroyQDragMoveEvent(void* ptr){
	static_cast<QDragMoveEvent*>(ptr)->~QDragMoveEvent();
}

