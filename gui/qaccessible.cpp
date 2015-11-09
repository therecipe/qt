#include "qaccessible.h"
#include <QModelIndex>
#include <QObject>
#include <QAccessibleEvent>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QAccessible>
#include "_cgo_export.h"

class MyQAccessible: public QAccessible {
public:
};

int QAccessible_InvalidEvent_Type(){
	return QAccessible::InvalidEvent;
}

int QAccessible_QAccessible_IsActive(){
	return QAccessible::isActive();
}

void* QAccessible_QAccessible_QueryAccessibleInterface(void* object){
	return QAccessible::queryAccessibleInterface(static_cast<QObject*>(object));
}

void QAccessible_QAccessible_SetRootObject(void* object){
	QAccessible::setRootObject(static_cast<QObject*>(object));
}

void QAccessible_QAccessible_UpdateAccessibility(void* event){
	QAccessible::updateAccessibility(static_cast<QAccessibleEvent*>(event));
}

