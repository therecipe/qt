#include "qaccessibletextcursorevent.h"
#include <QObject>
#include <QAccessibleInterface>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAccessible>
#include <QAccessibleTextCursorEvent>
#include "_cgo_export.h"

class MyQAccessibleTextCursorEvent: public QAccessibleTextCursorEvent {
public:
};

void* QAccessibleTextCursorEvent_NewQAccessibleTextCursorEvent2(void* iface, int cursorPos){
	return new QAccessibleTextCursorEvent(static_cast<QAccessibleInterface*>(iface), cursorPos);
}

void* QAccessibleTextCursorEvent_NewQAccessibleTextCursorEvent(void* object, int cursorPos){
	return new QAccessibleTextCursorEvent(static_cast<QObject*>(object), cursorPos);
}

int QAccessibleTextCursorEvent_CursorPosition(void* ptr){
	return static_cast<QAccessibleTextCursorEvent*>(ptr)->cursorPosition();
}

void QAccessibleTextCursorEvent_SetCursorPosition(void* ptr, int position){
	static_cast<QAccessibleTextCursorEvent*>(ptr)->setCursorPosition(position);
}

