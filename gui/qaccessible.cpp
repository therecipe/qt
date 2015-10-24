#include "qaccessible.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QAccessibleEvent>
#include <QString>
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

QtObjectPtr QAccessible_QAccessible_QueryAccessibleInterface(QtObjectPtr object){
	return QAccessible::queryAccessibleInterface(static_cast<QObject*>(object));
}

void QAccessible_QAccessible_SetRootObject(QtObjectPtr object){
	QAccessible::setRootObject(static_cast<QObject*>(object));
}

void QAccessible_QAccessible_UpdateAccessibility(QtObjectPtr event){
	QAccessible::updateAccessibility(static_cast<QAccessibleEvent*>(event));
}

