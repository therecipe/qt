#include "qaccessibleevent.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAccessible>
#include <QObject>
#include <QAccessibleInterface>
#include <QAccessibleEvent>
#include "_cgo_export.h"

class MyQAccessibleEvent: public QAccessibleEvent {
public:
};

void* QAccessibleEvent_NewQAccessibleEvent2(void* interfa, int ty){
	return new QAccessibleEvent(static_cast<QAccessibleInterface*>(interfa), static_cast<QAccessible::Event>(ty));
}

void* QAccessibleEvent_NewQAccessibleEvent(void* object, int ty){
	return new QAccessibleEvent(static_cast<QObject*>(object), static_cast<QAccessible::Event>(ty));
}

void* QAccessibleEvent_AccessibleInterface(void* ptr){
	return static_cast<QAccessibleEvent*>(ptr)->accessibleInterface();
}

int QAccessibleEvent_Child(void* ptr){
	return static_cast<QAccessibleEvent*>(ptr)->child();
}

void* QAccessibleEvent_Object(void* ptr){
	return static_cast<QAccessibleEvent*>(ptr)->object();
}

void QAccessibleEvent_SetChild(void* ptr, int child){
	static_cast<QAccessibleEvent*>(ptr)->setChild(child);
}

int QAccessibleEvent_Type(void* ptr){
	return static_cast<QAccessibleEvent*>(ptr)->type();
}

void QAccessibleEvent_DestroyQAccessibleEvent(void* ptr){
	static_cast<QAccessibleEvent*>(ptr)->~QAccessibleEvent();
}

