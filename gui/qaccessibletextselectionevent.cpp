#include "qaccessibletextselectionevent.h"
#include <QAccessible>
#include <QObject>
#include <QAccessibleInterface>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAccessibleTextSelectionEvent>
#include "_cgo_export.h"

class MyQAccessibleTextSelectionEvent: public QAccessibleTextSelectionEvent {
public:
};

void* QAccessibleTextSelectionEvent_NewQAccessibleTextSelectionEvent2(void* iface, int start, int end){
	return new QAccessibleTextSelectionEvent(static_cast<QAccessibleInterface*>(iface), start, end);
}

void* QAccessibleTextSelectionEvent_NewQAccessibleTextSelectionEvent(void* object, int start, int end){
	return new QAccessibleTextSelectionEvent(static_cast<QObject*>(object), start, end);
}

int QAccessibleTextSelectionEvent_SelectionEnd(void* ptr){
	return static_cast<QAccessibleTextSelectionEvent*>(ptr)->selectionEnd();
}

int QAccessibleTextSelectionEvent_SelectionStart(void* ptr){
	return static_cast<QAccessibleTextSelectionEvent*>(ptr)->selectionStart();
}

void QAccessibleTextSelectionEvent_SetSelection(void* ptr, int start, int end){
	static_cast<QAccessibleTextSelectionEvent*>(ptr)->setSelection(start, end);
}

