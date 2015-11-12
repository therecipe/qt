#include "qaccessiblevaluechangeevent.h"
#include <QModelIndex>
#include <QAccessible>
#include <QObject>
#include <QAccessibleInterface>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QAccessibleValueChangeEvent>
#include "_cgo_export.h"

class MyQAccessibleValueChangeEvent: public QAccessibleValueChangeEvent {
public:
};

void* QAccessibleValueChangeEvent_NewQAccessibleValueChangeEvent2(void* iface, void* val){
	return new QAccessibleValueChangeEvent(static_cast<QAccessibleInterface*>(iface), *static_cast<QVariant*>(val));
}

void* QAccessibleValueChangeEvent_NewQAccessibleValueChangeEvent(void* object, void* value){
	return new QAccessibleValueChangeEvent(static_cast<QObject*>(object), *static_cast<QVariant*>(value));
}

void QAccessibleValueChangeEvent_SetValue(void* ptr, void* value){
	static_cast<QAccessibleValueChangeEvent*>(ptr)->setValue(*static_cast<QVariant*>(value));
}

void* QAccessibleValueChangeEvent_Value(void* ptr){
	return new QVariant(static_cast<QAccessibleValueChangeEvent*>(ptr)->value());
}

