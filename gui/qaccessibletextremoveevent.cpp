#include "qaccessibletextremoveevent.h"
#include <QAccessible>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QAccessibleInterface>
#include <QAccessibleTextRemoveEvent>
#include "_cgo_export.h"

class MyQAccessibleTextRemoveEvent: public QAccessibleTextRemoveEvent {
public:
};

void* QAccessibleTextRemoveEvent_NewQAccessibleTextRemoveEvent2(void* iface, int position, char* text){
	return new QAccessibleTextRemoveEvent(static_cast<QAccessibleInterface*>(iface), position, QString(text));
}

void* QAccessibleTextRemoveEvent_NewQAccessibleTextRemoveEvent(void* object, int position, char* text){
	return new QAccessibleTextRemoveEvent(static_cast<QObject*>(object), position, QString(text));
}

int QAccessibleTextRemoveEvent_ChangePosition(void* ptr){
	return static_cast<QAccessibleTextRemoveEvent*>(ptr)->changePosition();
}

char* QAccessibleTextRemoveEvent_TextRemoved(void* ptr){
	return static_cast<QAccessibleTextRemoveEvent*>(ptr)->textRemoved().toUtf8().data();
}

