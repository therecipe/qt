#include "qchildevent.h"
#include <QUrl>
#include <QModelIndex>
#include <QEvent>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QChildEvent>
#include "_cgo_export.h"

class MyQChildEvent: public QChildEvent {
public:
};

QtObjectPtr QChildEvent_NewQChildEvent(int ty, QtObjectPtr child){
	return new QChildEvent(static_cast<QEvent::Type>(ty), static_cast<QObject*>(child));
}

int QChildEvent_Added(QtObjectPtr ptr){
	return static_cast<QChildEvent*>(ptr)->added();
}

QtObjectPtr QChildEvent_Child(QtObjectPtr ptr){
	return static_cast<QChildEvent*>(ptr)->child();
}

int QChildEvent_Polished(QtObjectPtr ptr){
	return static_cast<QChildEvent*>(ptr)->polished();
}

int QChildEvent_Removed(QtObjectPtr ptr){
	return static_cast<QChildEvent*>(ptr)->removed();
}

