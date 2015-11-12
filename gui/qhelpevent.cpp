#include "qhelpevent.h"
#include <QUrl>
#include <QModelIndex>
#include <QPoint>
#include <QEvent>
#include <QString>
#include <QVariant>
#include <QHelpEvent>
#include "_cgo_export.h"

class MyQHelpEvent: public QHelpEvent {
public:
};

void* QHelpEvent_NewQHelpEvent(int ty, void* pos, void* globalPos){
	return new QHelpEvent(static_cast<QEvent::Type>(ty), *static_cast<QPoint*>(pos), *static_cast<QPoint*>(globalPos));
}

int QHelpEvent_GlobalX(void* ptr){
	return static_cast<QHelpEvent*>(ptr)->globalX();
}

int QHelpEvent_GlobalY(void* ptr){
	return static_cast<QHelpEvent*>(ptr)->globalY();
}

int QHelpEvent_X(void* ptr){
	return static_cast<QHelpEvent*>(ptr)->x();
}

int QHelpEvent_Y(void* ptr){
	return static_cast<QHelpEvent*>(ptr)->y();
}

