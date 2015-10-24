#include "qaccessibleinterface.h"
#include <QModelIndex>
#include <QAccessible>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QAccessibleInterface>
#include "_cgo_export.h"

class MyQAccessibleInterface: public QAccessibleInterface {
public:
};

QtObjectPtr QAccessibleInterface_ActionInterface(QtObjectPtr ptr){
	return static_cast<QAccessibleInterface*>(ptr)->actionInterface();
}

QtObjectPtr QAccessibleInterface_Child(QtObjectPtr ptr, int index){
	return static_cast<QAccessibleInterface*>(ptr)->child(index);
}

QtObjectPtr QAccessibleInterface_ChildAt(QtObjectPtr ptr, int x, int y){
	return static_cast<QAccessibleInterface*>(ptr)->childAt(x, y);
}

int QAccessibleInterface_ChildCount(QtObjectPtr ptr){
	return static_cast<QAccessibleInterface*>(ptr)->childCount();
}

QtObjectPtr QAccessibleInterface_FocusChild(QtObjectPtr ptr){
	return static_cast<QAccessibleInterface*>(ptr)->focusChild();
}

int QAccessibleInterface_IndexOfChild(QtObjectPtr ptr, QtObjectPtr child){
	return static_cast<QAccessibleInterface*>(ptr)->indexOfChild(static_cast<QAccessibleInterface*>(child));
}

void QAccessibleInterface_Interface_cast(QtObjectPtr ptr, int ty){
	static_cast<QAccessibleInterface*>(ptr)->interface_cast(static_cast<QAccessible::InterfaceType>(ty));
}

int QAccessibleInterface_IsValid(QtObjectPtr ptr){
	return static_cast<QAccessibleInterface*>(ptr)->isValid();
}

QtObjectPtr QAccessibleInterface_Object(QtObjectPtr ptr){
	return static_cast<QAccessibleInterface*>(ptr)->object();
}

QtObjectPtr QAccessibleInterface_Parent(QtObjectPtr ptr){
	return static_cast<QAccessibleInterface*>(ptr)->parent();
}

int QAccessibleInterface_Role(QtObjectPtr ptr){
	return static_cast<QAccessibleInterface*>(ptr)->role();
}

void QAccessibleInterface_SetText(QtObjectPtr ptr, int t, char* text){
	static_cast<QAccessibleInterface*>(ptr)->setText(static_cast<QAccessible::Text>(t), QString(text));
}

QtObjectPtr QAccessibleInterface_TableCellInterface(QtObjectPtr ptr){
	return static_cast<QAccessibleInterface*>(ptr)->tableCellInterface();
}

QtObjectPtr QAccessibleInterface_TableInterface(QtObjectPtr ptr){
	return static_cast<QAccessibleInterface*>(ptr)->tableInterface();
}

char* QAccessibleInterface_Text(QtObjectPtr ptr, int t){
	return static_cast<QAccessibleInterface*>(ptr)->text(static_cast<QAccessible::Text>(t)).toUtf8().data();
}

QtObjectPtr QAccessibleInterface_TextInterface(QtObjectPtr ptr){
	return static_cast<QAccessibleInterface*>(ptr)->textInterface();
}

QtObjectPtr QAccessibleInterface_ValueInterface(QtObjectPtr ptr){
	return static_cast<QAccessibleInterface*>(ptr)->valueInterface();
}

QtObjectPtr QAccessibleInterface_Window(QtObjectPtr ptr){
	return static_cast<QAccessibleInterface*>(ptr)->window();
}

