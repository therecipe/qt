#include "qaccessibletextinsertevent.h"
#include <QAccessible>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAccessibleInterface>
#include <QAccessibleTextInsertEvent>
#include "_cgo_export.h"

class MyQAccessibleTextInsertEvent: public QAccessibleTextInsertEvent {
public:
};

void* QAccessibleTextInsertEvent_NewQAccessibleTextInsertEvent2(void* iface, int position, char* text){
	return new QAccessibleTextInsertEvent(static_cast<QAccessibleInterface*>(iface), position, QString(text));
}

void* QAccessibleTextInsertEvent_NewQAccessibleTextInsertEvent(void* object, int position, char* text){
	return new QAccessibleTextInsertEvent(static_cast<QObject*>(object), position, QString(text));
}

int QAccessibleTextInsertEvent_ChangePosition(void* ptr){
	return static_cast<QAccessibleTextInsertEvent*>(ptr)->changePosition();
}

char* QAccessibleTextInsertEvent_TextInserted(void* ptr){
	return static_cast<QAccessibleTextInsertEvent*>(ptr)->textInserted().toUtf8().data();
}

