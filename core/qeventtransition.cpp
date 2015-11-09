#include "qeventtransition.h"
#include <QUrl>
#include <QModelIndex>
#include <QState>
#include <QObject>
#include <QEvent>
#include <QString>
#include <QVariant>
#include <QEventTransition>
#include "_cgo_export.h"

class MyQEventTransition: public QEventTransition {
public:
};

void* QEventTransition_NewQEventTransition2(void* object, int ty, void* sourceState){
	return new QEventTransition(static_cast<QObject*>(object), static_cast<QEvent::Type>(ty), static_cast<QState*>(sourceState));
}

void* QEventTransition_NewQEventTransition(void* sourceState){
	return new QEventTransition(static_cast<QState*>(sourceState));
}

void* QEventTransition_EventSource(void* ptr){
	return static_cast<QEventTransition*>(ptr)->eventSource();
}

int QEventTransition_EventType(void* ptr){
	return static_cast<QEventTransition*>(ptr)->eventType();
}

void QEventTransition_SetEventSource(void* ptr, void* object){
	static_cast<QEventTransition*>(ptr)->setEventSource(static_cast<QObject*>(object));
}

void QEventTransition_SetEventType(void* ptr, int ty){
	static_cast<QEventTransition*>(ptr)->setEventType(static_cast<QEvent::Type>(ty));
}

void QEventTransition_DestroyQEventTransition(void* ptr){
	static_cast<QEventTransition*>(ptr)->~QEventTransition();
}

