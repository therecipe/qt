#include "qaccessiblevaluechangeevent.h"
#include <QAccessibleInterface>
#include <QAccessible>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAccessibleValueChangeEvent>
#include "_cgo_export.h"

class MyQAccessibleValueChangeEvent: public QAccessibleValueChangeEvent {
public:
};

QtObjectPtr QAccessibleValueChangeEvent_NewQAccessibleValueChangeEvent2(QtObjectPtr iface, char* val){
	return new QAccessibleValueChangeEvent(static_cast<QAccessibleInterface*>(iface), QVariant(val));
}

QtObjectPtr QAccessibleValueChangeEvent_NewQAccessibleValueChangeEvent(QtObjectPtr object, char* value){
	return new QAccessibleValueChangeEvent(static_cast<QObject*>(object), QVariant(value));
}

void QAccessibleValueChangeEvent_SetValue(QtObjectPtr ptr, char* value){
	static_cast<QAccessibleValueChangeEvent*>(ptr)->setValue(QVariant(value));
}

char* QAccessibleValueChangeEvent_Value(QtObjectPtr ptr){
	return static_cast<QAccessibleValueChangeEvent*>(ptr)->value().toString().toUtf8().data();
}

