#include "qaccessibletextupdateevent.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAccessibleInterface>
#include <QAccessible>
#include <QObject>
#include <QString>
#include <QAccessibleTextUpdateEvent>
#include "_cgo_export.h"

class MyQAccessibleTextUpdateEvent: public QAccessibleTextUpdateEvent {
public:
};

void* QAccessibleTextUpdateEvent_NewQAccessibleTextUpdateEvent2(void* iface, int position, char* oldText, char* text){
	return new QAccessibleTextUpdateEvent(static_cast<QAccessibleInterface*>(iface), position, QString(oldText), QString(text));
}

void* QAccessibleTextUpdateEvent_NewQAccessibleTextUpdateEvent(void* object, int position, char* oldText, char* text){
	return new QAccessibleTextUpdateEvent(static_cast<QObject*>(object), position, QString(oldText), QString(text));
}

int QAccessibleTextUpdateEvent_ChangePosition(void* ptr){
	return static_cast<QAccessibleTextUpdateEvent*>(ptr)->changePosition();
}

char* QAccessibleTextUpdateEvent_TextInserted(void* ptr){
	return static_cast<QAccessibleTextUpdateEvent*>(ptr)->textInserted().toUtf8().data();
}

char* QAccessibleTextUpdateEvent_TextRemoved(void* ptr){
	return static_cast<QAccessibleTextUpdateEvent*>(ptr)->textRemoved().toUtf8().data();
}

