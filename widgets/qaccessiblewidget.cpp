#include "qaccessiblewidget.h"
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QAccessibleInterface>
#include <QColor>
#include <QAccessible>
#include <QString>
#include <QVariant>
#include <QAccessibleWidget>
#include "_cgo_export.h"

class MyQAccessibleWidget: public QAccessibleWidget {
public:
};

void* QAccessibleWidget_NewQAccessibleWidget(void* w, int role, char* name){
	return new QAccessibleWidget(static_cast<QWidget*>(w), static_cast<QAccessible::Role>(role), QString(name));
}

char* QAccessibleWidget_ActionNames(void* ptr){
	return static_cast<QAccessibleWidget*>(ptr)->actionNames().join("|").toUtf8().data();
}

void* QAccessibleWidget_BackgroundColor(void* ptr){
	return new QColor(static_cast<QAccessibleWidget*>(ptr)->backgroundColor());
}

void* QAccessibleWidget_Child(void* ptr, int index){
	return static_cast<QAccessibleWidget*>(ptr)->child(index);
}

int QAccessibleWidget_ChildCount(void* ptr){
	return static_cast<QAccessibleWidget*>(ptr)->childCount();
}

void QAccessibleWidget_DoAction(void* ptr, char* actionName){
	static_cast<QAccessibleWidget*>(ptr)->doAction(QString(actionName));
}

void* QAccessibleWidget_FocusChild(void* ptr){
	return static_cast<QAccessibleWidget*>(ptr)->focusChild();
}

void* QAccessibleWidget_ForegroundColor(void* ptr){
	return new QColor(static_cast<QAccessibleWidget*>(ptr)->foregroundColor());
}

int QAccessibleWidget_IndexOfChild(void* ptr, void* child){
	return static_cast<QAccessibleWidget*>(ptr)->indexOfChild(static_cast<QAccessibleInterface*>(child));
}

void* QAccessibleWidget_Interface_cast(void* ptr, int t){
	return static_cast<QAccessibleWidget*>(ptr)->interface_cast(static_cast<QAccessible::InterfaceType>(t));
}

int QAccessibleWidget_IsValid(void* ptr){
	return static_cast<QAccessibleWidget*>(ptr)->isValid();
}

char* QAccessibleWidget_KeyBindingsForAction(void* ptr, char* actionName){
	return static_cast<QAccessibleWidget*>(ptr)->keyBindingsForAction(QString(actionName)).join("|").toUtf8().data();
}

void* QAccessibleWidget_Parent(void* ptr){
	return static_cast<QAccessibleWidget*>(ptr)->parent();
}

int QAccessibleWidget_Role(void* ptr){
	return static_cast<QAccessibleWidget*>(ptr)->role();
}

char* QAccessibleWidget_Text(void* ptr, int t){
	return static_cast<QAccessibleWidget*>(ptr)->text(static_cast<QAccessible::Text>(t)).toUtf8().data();
}

void* QAccessibleWidget_Window(void* ptr){
	return static_cast<QAccessibleWidget*>(ptr)->window();
}

