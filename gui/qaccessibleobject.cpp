#include "qaccessibleobject.h"
#include <QModelIndex>
#include <QAccessible>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QAccessibleObject>
#include "_cgo_export.h"

class MyQAccessibleObject: public QAccessibleObject {
public:
};

void* QAccessibleObject_ChildAt(void* ptr, int x, int y){
	return static_cast<QAccessibleObject*>(ptr)->childAt(x, y);
}

int QAccessibleObject_IsValid(void* ptr){
	return static_cast<QAccessibleObject*>(ptr)->isValid();
}

void* QAccessibleObject_Object(void* ptr){
	return static_cast<QAccessibleObject*>(ptr)->object();
}

void QAccessibleObject_SetText(void* ptr, int t, char* text){
	static_cast<QAccessibleObject*>(ptr)->setText(static_cast<QAccessible::Text>(t), QString(text));
}

