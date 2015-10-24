#include "qaccessibletextupdateevent.h"
#include <QAccessibleInterface>
#include <QAccessible>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAccessibleTextUpdateEvent>
#include "_cgo_export.h"

class MyQAccessibleTextUpdateEvent: public QAccessibleTextUpdateEvent {
public:
};

QtObjectPtr QAccessibleTextUpdateEvent_NewQAccessibleTextUpdateEvent2(QtObjectPtr iface, int position, char* oldText, char* text){
	return new QAccessibleTextUpdateEvent(static_cast<QAccessibleInterface*>(iface), position, QString(oldText), QString(text));
}

QtObjectPtr QAccessibleTextUpdateEvent_NewQAccessibleTextUpdateEvent(QtObjectPtr object, int position, char* oldText, char* text){
	return new QAccessibleTextUpdateEvent(static_cast<QObject*>(object), position, QString(oldText), QString(text));
}

int QAccessibleTextUpdateEvent_ChangePosition(QtObjectPtr ptr){
	return static_cast<QAccessibleTextUpdateEvent*>(ptr)->changePosition();
}

char* QAccessibleTextUpdateEvent_TextInserted(QtObjectPtr ptr){
	return static_cast<QAccessibleTextUpdateEvent*>(ptr)->textInserted().toUtf8().data();
}

char* QAccessibleTextUpdateEvent_TextRemoved(QtObjectPtr ptr){
	return static_cast<QAccessibleTextUpdateEvent*>(ptr)->textRemoved().toUtf8().data();
}

