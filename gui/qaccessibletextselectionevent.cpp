#include "qaccessibletextselectionevent.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAccessible>
#include <QObject>
#include <QAccessibleInterface>
#include <QAccessibleTextSelectionEvent>
#include "_cgo_export.h"

class MyQAccessibleTextSelectionEvent: public QAccessibleTextSelectionEvent {
public:
};

QtObjectPtr QAccessibleTextSelectionEvent_NewQAccessibleTextSelectionEvent2(QtObjectPtr iface, int start, int end){
	return new QAccessibleTextSelectionEvent(static_cast<QAccessibleInterface*>(iface), start, end);
}

QtObjectPtr QAccessibleTextSelectionEvent_NewQAccessibleTextSelectionEvent(QtObjectPtr object, int start, int end){
	return new QAccessibleTextSelectionEvent(static_cast<QObject*>(object), start, end);
}

int QAccessibleTextSelectionEvent_SelectionEnd(QtObjectPtr ptr){
	return static_cast<QAccessibleTextSelectionEvent*>(ptr)->selectionEnd();
}

int QAccessibleTextSelectionEvent_SelectionStart(QtObjectPtr ptr){
	return static_cast<QAccessibleTextSelectionEvent*>(ptr)->selectionStart();
}

void QAccessibleTextSelectionEvent_SetSelection(QtObjectPtr ptr, int start, int end){
	static_cast<QAccessibleTextSelectionEvent*>(ptr)->setSelection(start, end);
}

