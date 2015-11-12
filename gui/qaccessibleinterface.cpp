#include "qaccessibleinterface.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QColor>
#include <QAccessible>
#include <QAccessibleInterface>
#include "_cgo_export.h"

class MyQAccessibleInterface: public QAccessibleInterface {
public:
};

void* QAccessibleInterface_ActionInterface(void* ptr){
	return static_cast<QAccessibleInterface*>(ptr)->actionInterface();
}

void* QAccessibleInterface_BackgroundColor(void* ptr){
	return new QColor(static_cast<QAccessibleInterface*>(ptr)->backgroundColor());
}

void* QAccessibleInterface_Child(void* ptr, int index){
	return static_cast<QAccessibleInterface*>(ptr)->child(index);
}

void* QAccessibleInterface_ChildAt(void* ptr, int x, int y){
	return static_cast<QAccessibleInterface*>(ptr)->childAt(x, y);
}

int QAccessibleInterface_ChildCount(void* ptr){
	return static_cast<QAccessibleInterface*>(ptr)->childCount();
}

void* QAccessibleInterface_FocusChild(void* ptr){
	return static_cast<QAccessibleInterface*>(ptr)->focusChild();
}

void* QAccessibleInterface_ForegroundColor(void* ptr){
	return new QColor(static_cast<QAccessibleInterface*>(ptr)->foregroundColor());
}

int QAccessibleInterface_IndexOfChild(void* ptr, void* child){
	return static_cast<QAccessibleInterface*>(ptr)->indexOfChild(static_cast<QAccessibleInterface*>(child));
}

void* QAccessibleInterface_Interface_cast(void* ptr, int ty){
	return static_cast<QAccessibleInterface*>(ptr)->interface_cast(static_cast<QAccessible::InterfaceType>(ty));
}

int QAccessibleInterface_IsValid(void* ptr){
	return static_cast<QAccessibleInterface*>(ptr)->isValid();
}

void* QAccessibleInterface_Object(void* ptr){
	return static_cast<QAccessibleInterface*>(ptr)->object();
}

void* QAccessibleInterface_Parent(void* ptr){
	return static_cast<QAccessibleInterface*>(ptr)->parent();
}

int QAccessibleInterface_Role(void* ptr){
	return static_cast<QAccessibleInterface*>(ptr)->role();
}

void QAccessibleInterface_SetText(void* ptr, int t, char* text){
	static_cast<QAccessibleInterface*>(ptr)->setText(static_cast<QAccessible::Text>(t), QString(text));
}

void* QAccessibleInterface_TableCellInterface(void* ptr){
	return static_cast<QAccessibleInterface*>(ptr)->tableCellInterface();
}

void* QAccessibleInterface_TableInterface(void* ptr){
	return static_cast<QAccessibleInterface*>(ptr)->tableInterface();
}

char* QAccessibleInterface_Text(void* ptr, int t){
	return static_cast<QAccessibleInterface*>(ptr)->text(static_cast<QAccessible::Text>(t)).toUtf8().data();
}

void* QAccessibleInterface_TextInterface(void* ptr){
	return static_cast<QAccessibleInterface*>(ptr)->textInterface();
}

void* QAccessibleInterface_ValueInterface(void* ptr){
	return static_cast<QAccessibleInterface*>(ptr)->valueInterface();
}

void* QAccessibleInterface_Window(void* ptr){
	return static_cast<QAccessibleInterface*>(ptr)->window();
}

