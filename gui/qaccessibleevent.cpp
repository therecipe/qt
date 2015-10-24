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

QtObjectPtr QAccessibleEvent_NewQAccessibleEvent2(QtObjectPtr interfa, int ty){
	return new QAccessibleEvent(static_cast<QAccessibleInterface*>(interfa), static_cast<QAccessible::Event>(ty));
}

QtObjectPtr QAccessibleEvent_NewQAccessibleEvent(QtObjectPtr object, int ty){
	return new QAccessibleEvent(static_cast<QObject*>(object), static_cast<QAccessible::Event>(ty));
}

QtObjectPtr QAccessibleEvent_AccessibleInterface(QtObjectPtr ptr){
	return static_cast<QAccessibleEvent*>(ptr)->accessibleInterface();
}

int QAccessibleEvent_Child(QtObjectPtr ptr){
	return static_cast<QAccessibleEvent*>(ptr)->child();
}

QtObjectPtr QAccessibleEvent_Object(QtObjectPtr ptr){
	return static_cast<QAccessibleEvent*>(ptr)->object();
}

void QAccessibleEvent_SetChild(QtObjectPtr ptr, int child){
	static_cast<QAccessibleEvent*>(ptr)->setChild(child);
}

int QAccessibleEvent_Type(QtObjectPtr ptr){
	return static_cast<QAccessibleEvent*>(ptr)->type();
}

void QAccessibleEvent_DestroyQAccessibleEvent(QtObjectPtr ptr){
	static_cast<QAccessibleEvent*>(ptr)->~QAccessibleEvent();
}

