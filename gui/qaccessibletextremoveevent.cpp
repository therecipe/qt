#include "qaccessibletextremoveevent.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QAccessibleInterface>
#include <QAccessible>
#include <QAccessibleTextRemoveEvent>
#include "_cgo_export.h"

class MyQAccessibleTextRemoveEvent: public QAccessibleTextRemoveEvent {
public:
};

QtObjectPtr QAccessibleTextRemoveEvent_NewQAccessibleTextRemoveEvent2(QtObjectPtr iface, int position, char* text){
	return new QAccessibleTextRemoveEvent(static_cast<QAccessibleInterface*>(iface), position, QString(text));
}

QtObjectPtr QAccessibleTextRemoveEvent_NewQAccessibleTextRemoveEvent(QtObjectPtr object, int position, char* text){
	return new QAccessibleTextRemoveEvent(static_cast<QObject*>(object), position, QString(text));
}

int QAccessibleTextRemoveEvent_ChangePosition(QtObjectPtr ptr){
	return static_cast<QAccessibleTextRemoveEvent*>(ptr)->changePosition();
}

char* QAccessibleTextRemoveEvent_TextRemoved(QtObjectPtr ptr){
	return static_cast<QAccessibleTextRemoveEvent*>(ptr)->textRemoved().toUtf8().data();
}

