#include "qaccessibletablemodelchangeevent.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAccessible>
#include <QObject>
#include <QAccessibleInterface>
#include <QAccessibleTableModelChangeEvent>
#include "_cgo_export.h"

class MyQAccessibleTableModelChangeEvent: public QAccessibleTableModelChangeEvent {
public:
};

QtObjectPtr QAccessibleTableModelChangeEvent_NewQAccessibleTableModelChangeEvent2(QtObjectPtr iface, int changeType){
	return new QAccessibleTableModelChangeEvent(static_cast<QAccessibleInterface*>(iface), static_cast<QAccessibleTableModelChangeEvent::ModelChangeType>(changeType));
}

QtObjectPtr QAccessibleTableModelChangeEvent_NewQAccessibleTableModelChangeEvent(QtObjectPtr object, int changeType){
	return new QAccessibleTableModelChangeEvent(static_cast<QObject*>(object), static_cast<QAccessibleTableModelChangeEvent::ModelChangeType>(changeType));
}

int QAccessibleTableModelChangeEvent_FirstColumn(QtObjectPtr ptr){
	return static_cast<QAccessibleTableModelChangeEvent*>(ptr)->firstColumn();
}

int QAccessibleTableModelChangeEvent_FirstRow(QtObjectPtr ptr){
	return static_cast<QAccessibleTableModelChangeEvent*>(ptr)->firstRow();
}

int QAccessibleTableModelChangeEvent_LastColumn(QtObjectPtr ptr){
	return static_cast<QAccessibleTableModelChangeEvent*>(ptr)->lastColumn();
}

int QAccessibleTableModelChangeEvent_LastRow(QtObjectPtr ptr){
	return static_cast<QAccessibleTableModelChangeEvent*>(ptr)->lastRow();
}

int QAccessibleTableModelChangeEvent_ModelChangeType(QtObjectPtr ptr){
	return static_cast<QAccessibleTableModelChangeEvent*>(ptr)->modelChangeType();
}

void QAccessibleTableModelChangeEvent_SetFirstColumn(QtObjectPtr ptr, int column){
	static_cast<QAccessibleTableModelChangeEvent*>(ptr)->setFirstColumn(column);
}

void QAccessibleTableModelChangeEvent_SetFirstRow(QtObjectPtr ptr, int row){
	static_cast<QAccessibleTableModelChangeEvent*>(ptr)->setFirstRow(row);
}

void QAccessibleTableModelChangeEvent_SetLastColumn(QtObjectPtr ptr, int column){
	static_cast<QAccessibleTableModelChangeEvent*>(ptr)->setLastColumn(column);
}

void QAccessibleTableModelChangeEvent_SetLastRow(QtObjectPtr ptr, int row){
	static_cast<QAccessibleTableModelChangeEvent*>(ptr)->setLastRow(row);
}

void QAccessibleTableModelChangeEvent_SetModelChangeType(QtObjectPtr ptr, int changeType){
	static_cast<QAccessibleTableModelChangeEvent*>(ptr)->setModelChangeType(static_cast<QAccessibleTableModelChangeEvent::ModelChangeType>(changeType));
}

