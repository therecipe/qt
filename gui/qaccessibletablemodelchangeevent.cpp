#include "qaccessibletablemodelchangeevent.h"
#include <QUrl>
#include <QModelIndex>
#include <QAccessible>
#include <QObject>
#include <QAccessibleInterface>
#include <QString>
#include <QVariant>
#include <QAccessibleTableModelChangeEvent>
#include "_cgo_export.h"

class MyQAccessibleTableModelChangeEvent: public QAccessibleTableModelChangeEvent {
public:
};

void* QAccessibleTableModelChangeEvent_NewQAccessibleTableModelChangeEvent2(void* iface, int changeType){
	return new QAccessibleTableModelChangeEvent(static_cast<QAccessibleInterface*>(iface), static_cast<QAccessibleTableModelChangeEvent::ModelChangeType>(changeType));
}

void* QAccessibleTableModelChangeEvent_NewQAccessibleTableModelChangeEvent(void* object, int changeType){
	return new QAccessibleTableModelChangeEvent(static_cast<QObject*>(object), static_cast<QAccessibleTableModelChangeEvent::ModelChangeType>(changeType));
}

int QAccessibleTableModelChangeEvent_FirstColumn(void* ptr){
	return static_cast<QAccessibleTableModelChangeEvent*>(ptr)->firstColumn();
}

int QAccessibleTableModelChangeEvent_FirstRow(void* ptr){
	return static_cast<QAccessibleTableModelChangeEvent*>(ptr)->firstRow();
}

int QAccessibleTableModelChangeEvent_LastColumn(void* ptr){
	return static_cast<QAccessibleTableModelChangeEvent*>(ptr)->lastColumn();
}

int QAccessibleTableModelChangeEvent_LastRow(void* ptr){
	return static_cast<QAccessibleTableModelChangeEvent*>(ptr)->lastRow();
}

int QAccessibleTableModelChangeEvent_ModelChangeType(void* ptr){
	return static_cast<QAccessibleTableModelChangeEvent*>(ptr)->modelChangeType();
}

void QAccessibleTableModelChangeEvent_SetFirstColumn(void* ptr, int column){
	static_cast<QAccessibleTableModelChangeEvent*>(ptr)->setFirstColumn(column);
}

void QAccessibleTableModelChangeEvent_SetFirstRow(void* ptr, int row){
	static_cast<QAccessibleTableModelChangeEvent*>(ptr)->setFirstRow(row);
}

void QAccessibleTableModelChangeEvent_SetLastColumn(void* ptr, int column){
	static_cast<QAccessibleTableModelChangeEvent*>(ptr)->setLastColumn(column);
}

void QAccessibleTableModelChangeEvent_SetLastRow(void* ptr, int row){
	static_cast<QAccessibleTableModelChangeEvent*>(ptr)->setLastRow(row);
}

void QAccessibleTableModelChangeEvent_SetModelChangeType(void* ptr, int changeType){
	static_cast<QAccessibleTableModelChangeEvent*>(ptr)->setModelChangeType(static_cast<QAccessibleTableModelChangeEvent::ModelChangeType>(changeType));
}

