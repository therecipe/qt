#include "qeventtransition.h"
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QEvent>
#include <QState>
#include <QEventTransition>
#include "_cgo_export.h"

class MyQEventTransition: public QEventTransition {
public:
};

QtObjectPtr QEventTransition_NewQEventTransition2(QtObjectPtr object, int ty, QtObjectPtr sourceState){
	return new QEventTransition(static_cast<QObject*>(object), static_cast<QEvent::Type>(ty), static_cast<QState*>(sourceState));
}

QtObjectPtr QEventTransition_NewQEventTransition(QtObjectPtr sourceState){
	return new QEventTransition(static_cast<QState*>(sourceState));
}

QtObjectPtr QEventTransition_EventSource(QtObjectPtr ptr){
	return static_cast<QEventTransition*>(ptr)->eventSource();
}

int QEventTransition_EventType(QtObjectPtr ptr){
	return static_cast<QEventTransition*>(ptr)->eventType();
}

void QEventTransition_SetEventSource(QtObjectPtr ptr, QtObjectPtr object){
	static_cast<QEventTransition*>(ptr)->setEventSource(static_cast<QObject*>(object));
}

void QEventTransition_SetEventType(QtObjectPtr ptr, int ty){
	static_cast<QEventTransition*>(ptr)->setEventType(static_cast<QEvent::Type>(ty));
}

void QEventTransition_DestroyQEventTransition(QtObjectPtr ptr){
	static_cast<QEventTransition*>(ptr)->~QEventTransition();
}

