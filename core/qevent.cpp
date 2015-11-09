#include "qevent.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QEvent>
#include "_cgo_export.h"

class MyQEvent: public QEvent {
public:
};

void* QEvent_NewQEvent(int ty){
	return new QEvent(static_cast<QEvent::Type>(ty));
}

void QEvent_Accept(void* ptr){
	static_cast<QEvent*>(ptr)->accept();
}

void QEvent_Ignore(void* ptr){
	static_cast<QEvent*>(ptr)->ignore();
}

int QEvent_IsAccepted(void* ptr){
	return static_cast<QEvent*>(ptr)->isAccepted();
}

int QEvent_QEvent_RegisterEventType(int hint){
	return QEvent::registerEventType(hint);
}

void QEvent_SetAccepted(void* ptr, int accepted){
	static_cast<QEvent*>(ptr)->setAccepted(accepted != 0);
}

int QEvent_Spontaneous(void* ptr){
	return static_cast<QEvent*>(ptr)->spontaneous();
}

int QEvent_Type(void* ptr){
	return static_cast<QEvent*>(ptr)->type();
}

void QEvent_DestroyQEvent(void* ptr){
	static_cast<QEvent*>(ptr)->~QEvent();
}

