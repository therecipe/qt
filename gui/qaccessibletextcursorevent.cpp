#include "qaccessibletextcursorevent.h"
#include <QObject>
#include <QAccessibleInterface>
#include <QAccessible>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAccessibleTextCursorEvent>
#include "_cgo_export.h"

class MyQAccessibleTextCursorEvent: public QAccessibleTextCursorEvent {
public:
};

QtObjectPtr QAccessibleTextCursorEvent_NewQAccessibleTextCursorEvent2(QtObjectPtr iface, int cursorPos){
	return new QAccessibleTextCursorEvent(static_cast<QAccessibleInterface*>(iface), cursorPos);
}

QtObjectPtr QAccessibleTextCursorEvent_NewQAccessibleTextCursorEvent(QtObjectPtr object, int cursorPos){
	return new QAccessibleTextCursorEvent(static_cast<QObject*>(object), cursorPos);
}

int QAccessibleTextCursorEvent_CursorPosition(QtObjectPtr ptr){
	return static_cast<QAccessibleTextCursorEvent*>(ptr)->cursorPosition();
}

void QAccessibleTextCursorEvent_SetCursorPosition(QtObjectPtr ptr, int position){
	static_cast<QAccessibleTextCursorEvent*>(ptr)->setCursorPosition(position);
}

