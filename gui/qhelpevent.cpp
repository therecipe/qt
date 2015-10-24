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

QtObjectPtr QHelpEvent_NewQHelpEvent(int ty, QtObjectPtr pos, QtObjectPtr globalPos){
	return new QHelpEvent(static_cast<QEvent::Type>(ty), *static_cast<QPoint*>(pos), *static_cast<QPoint*>(globalPos));
}

int QHelpEvent_GlobalX(QtObjectPtr ptr){
	return static_cast<QHelpEvent*>(ptr)->globalX();
}

int QHelpEvent_GlobalY(QtObjectPtr ptr){
	return static_cast<QHelpEvent*>(ptr)->globalY();
}

int QHelpEvent_X(QtObjectPtr ptr){
	return static_cast<QHelpEvent*>(ptr)->x();
}

int QHelpEvent_Y(QtObjectPtr ptr){
	return static_cast<QHelpEvent*>(ptr)->y();
}

