#include "qaccessiblewidget.h"
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAccessible>
#include <QAccessibleInterface>
#include <QAccessibleWidget>
#include "_cgo_export.h"

class MyQAccessibleWidget: public QAccessibleWidget {
public:
};

QtObjectPtr QAccessibleWidget_NewQAccessibleWidget(QtObjectPtr w, int role, char* name){
	return new QAccessibleWidget(static_cast<QWidget*>(w), static_cast<QAccessible::Role>(role), QString(name));
}

char* QAccessibleWidget_ActionNames(QtObjectPtr ptr){
	return static_cast<QAccessibleWidget*>(ptr)->actionNames().join("|").toUtf8().data();
}

QtObjectPtr QAccessibleWidget_Child(QtObjectPtr ptr, int index){
	return static_cast<QAccessibleWidget*>(ptr)->child(index);
}

int QAccessibleWidget_ChildCount(QtObjectPtr ptr){
	return static_cast<QAccessibleWidget*>(ptr)->childCount();
}

void QAccessibleWidget_DoAction(QtObjectPtr ptr, char* actionName){
	static_cast<QAccessibleWidget*>(ptr)->doAction(QString(actionName));
}

QtObjectPtr QAccessibleWidget_FocusChild(QtObjectPtr ptr){
	return static_cast<QAccessibleWidget*>(ptr)->focusChild();
}

int QAccessibleWidget_IndexOfChild(QtObjectPtr ptr, QtObjectPtr child){
	return static_cast<QAccessibleWidget*>(ptr)->indexOfChild(static_cast<QAccessibleInterface*>(child));
}

void QAccessibleWidget_Interface_cast(QtObjectPtr ptr, int t){
	static_cast<QAccessibleWidget*>(ptr)->interface_cast(static_cast<QAccessible::InterfaceType>(t));
}

int QAccessibleWidget_IsValid(QtObjectPtr ptr){
	return static_cast<QAccessibleWidget*>(ptr)->isValid();
}

char* QAccessibleWidget_KeyBindingsForAction(QtObjectPtr ptr, char* actionName){
	return static_cast<QAccessibleWidget*>(ptr)->keyBindingsForAction(QString(actionName)).join("|").toUtf8().data();
}

QtObjectPtr QAccessibleWidget_Parent(QtObjectPtr ptr){
	return static_cast<QAccessibleWidget*>(ptr)->parent();
}

int QAccessibleWidget_Role(QtObjectPtr ptr){
	return static_cast<QAccessibleWidget*>(ptr)->role();
}

char* QAccessibleWidget_Text(QtObjectPtr ptr, int t){
	return static_cast<QAccessibleWidget*>(ptr)->text(static_cast<QAccessible::Text>(t)).toUtf8().data();
}

QtObjectPtr QAccessibleWidget_Window(QtObjectPtr ptr){
	return static_cast<QAccessibleWidget*>(ptr)->window();
}

