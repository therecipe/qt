#include "qchildevent.h"
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QEvent>
#include <QChildEvent>
#include "_cgo_export.h"

class MyQChildEvent: public QChildEvent {
public:
};

void* QChildEvent_NewQChildEvent(int ty, void* child){
	return new QChildEvent(static_cast<QEvent::Type>(ty), static_cast<QObject*>(child));
}

int QChildEvent_Added(void* ptr){
	return static_cast<QChildEvent*>(ptr)->added();
}

void* QChildEvent_Child(void* ptr){
	return static_cast<QChildEvent*>(ptr)->child();
}

int QChildEvent_Polished(void* ptr){
	return static_cast<QChildEvent*>(ptr)->polished();
}

int QChildEvent_Removed(void* ptr){
	return static_cast<QChildEvent*>(ptr)->removed();
}

