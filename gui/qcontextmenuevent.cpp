#include "qcontextmenuevent.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPoint>
#include <QContextMenuEvent>
#include "_cgo_export.h"

class MyQContextMenuEvent: public QContextMenuEvent {
public:
};

void* QContextMenuEvent_NewQContextMenuEvent3(int reason, void* pos){
	return new QContextMenuEvent(static_cast<QContextMenuEvent::Reason>(reason), *static_cast<QPoint*>(pos));
}

void* QContextMenuEvent_NewQContextMenuEvent2(int reason, void* pos, void* globalPos){
	return new QContextMenuEvent(static_cast<QContextMenuEvent::Reason>(reason), *static_cast<QPoint*>(pos), *static_cast<QPoint*>(globalPos));
}

void* QContextMenuEvent_NewQContextMenuEvent(int reason, void* pos, void* globalPos, int modifiers){
	return new QContextMenuEvent(static_cast<QContextMenuEvent::Reason>(reason), *static_cast<QPoint*>(pos), *static_cast<QPoint*>(globalPos), static_cast<Qt::KeyboardModifier>(modifiers));
}

int QContextMenuEvent_GlobalX(void* ptr){
	return static_cast<QContextMenuEvent*>(ptr)->globalX();
}

int QContextMenuEvent_GlobalY(void* ptr){
	return static_cast<QContextMenuEvent*>(ptr)->globalY();
}

int QContextMenuEvent_Reason(void* ptr){
	return static_cast<QContextMenuEvent*>(ptr)->reason();
}

int QContextMenuEvent_X(void* ptr){
	return static_cast<QContextMenuEvent*>(ptr)->x();
}

int QContextMenuEvent_Y(void* ptr){
	return static_cast<QContextMenuEvent*>(ptr)->y();
}

