#include "qaccessibletextinsertevent.h"
#include <QObject>
#include <QAccessibleInterface>
#include <QAccessible>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAccessibleTextInsertEvent>
#include "_cgo_export.h"

class MyQAccessibleTextInsertEvent: public QAccessibleTextInsertEvent {
public:
};

QtObjectPtr QAccessibleTextInsertEvent_NewQAccessibleTextInsertEvent2(QtObjectPtr iface, int position, char* text){
	return new QAccessibleTextInsertEvent(static_cast<QAccessibleInterface*>(iface), position, QString(text));
}

QtObjectPtr QAccessibleTextInsertEvent_NewQAccessibleTextInsertEvent(QtObjectPtr object, int position, char* text){
	return new QAccessibleTextInsertEvent(static_cast<QObject*>(object), position, QString(text));
}

int QAccessibleTextInsertEvent_ChangePosition(QtObjectPtr ptr){
	return static_cast<QAccessibleTextInsertEvent*>(ptr)->changePosition();
}

char* QAccessibleTextInsertEvent_TextInserted(QtObjectPtr ptr){
	return static_cast<QAccessibleTextInsertEvent*>(ptr)->textInserted().toUtf8().data();
}

