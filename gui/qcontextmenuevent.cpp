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

QtObjectPtr QContextMenuEvent_NewQContextMenuEvent3(int reason, QtObjectPtr pos){
	return new QContextMenuEvent(static_cast<QContextMenuEvent::Reason>(reason), *static_cast<QPoint*>(pos));
}

QtObjectPtr QContextMenuEvent_NewQContextMenuEvent2(int reason, QtObjectPtr pos, QtObjectPtr globalPos){
	return new QContextMenuEvent(static_cast<QContextMenuEvent::Reason>(reason), *static_cast<QPoint*>(pos), *static_cast<QPoint*>(globalPos));
}

QtObjectPtr QContextMenuEvent_NewQContextMenuEvent(int reason, QtObjectPtr pos, QtObjectPtr globalPos, int modifiers){
	return new QContextMenuEvent(static_cast<QContextMenuEvent::Reason>(reason), *static_cast<QPoint*>(pos), *static_cast<QPoint*>(globalPos), static_cast<Qt::KeyboardModifier>(modifiers));
}

int QContextMenuEvent_GlobalX(QtObjectPtr ptr){
	return static_cast<QContextMenuEvent*>(ptr)->globalX();
}

int QContextMenuEvent_GlobalY(QtObjectPtr ptr){
	return static_cast<QContextMenuEvent*>(ptr)->globalY();
}

int QContextMenuEvent_Reason(QtObjectPtr ptr){
	return static_cast<QContextMenuEvent*>(ptr)->reason();
}

int QContextMenuEvent_X(QtObjectPtr ptr){
	return static_cast<QContextMenuEvent*>(ptr)->x();
}

int QContextMenuEvent_Y(QtObjectPtr ptr){
	return static_cast<QContextMenuEvent*>(ptr)->y();
}

