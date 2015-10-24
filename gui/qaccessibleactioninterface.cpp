#include "qaccessibleactioninterface.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAccessible>
#include <QAccessibleActionInterface>
#include "_cgo_export.h"

class MyQAccessibleActionInterface: public QAccessibleActionInterface {
public:
};

char* QAccessibleActionInterface_LocalizedActionDescription(QtObjectPtr ptr, char* actionName){
	return static_cast<QAccessibleActionInterface*>(ptr)->localizedActionDescription(QString(actionName)).toUtf8().data();
}

char* QAccessibleActionInterface_LocalizedActionName(QtObjectPtr ptr, char* actionName){
	return static_cast<QAccessibleActionInterface*>(ptr)->localizedActionName(QString(actionName)).toUtf8().data();
}

char* QAccessibleActionInterface_ActionNames(QtObjectPtr ptr){
	return static_cast<QAccessibleActionInterface*>(ptr)->actionNames().join("|").toUtf8().data();
}

char* QAccessibleActionInterface_QAccessibleActionInterface_DecreaseAction(){
	return QAccessibleActionInterface::decreaseAction().toUtf8().data();
}

void QAccessibleActionInterface_DoAction(QtObjectPtr ptr, char* actionName){
	static_cast<QAccessibleActionInterface*>(ptr)->doAction(QString(actionName));
}

char* QAccessibleActionInterface_QAccessibleActionInterface_IncreaseAction(){
	return QAccessibleActionInterface::increaseAction().toUtf8().data();
}

char* QAccessibleActionInterface_KeyBindingsForAction(QtObjectPtr ptr, char* actionName){
	return static_cast<QAccessibleActionInterface*>(ptr)->keyBindingsForAction(QString(actionName)).join("|").toUtf8().data();
}

char* QAccessibleActionInterface_QAccessibleActionInterface_NextPageAction(){
	return QAccessibleActionInterface::nextPageAction().toUtf8().data();
}

char* QAccessibleActionInterface_QAccessibleActionInterface_PressAction(){
	return QAccessibleActionInterface::pressAction().toUtf8().data();
}

char* QAccessibleActionInterface_QAccessibleActionInterface_PreviousPageAction(){
	return QAccessibleActionInterface::previousPageAction().toUtf8().data();
}

char* QAccessibleActionInterface_QAccessibleActionInterface_ScrollDownAction(){
	return QAccessibleActionInterface::scrollDownAction().toUtf8().data();
}

char* QAccessibleActionInterface_QAccessibleActionInterface_ScrollLeftAction(){
	return QAccessibleActionInterface::scrollLeftAction().toUtf8().data();
}

char* QAccessibleActionInterface_QAccessibleActionInterface_ScrollRightAction(){
	return QAccessibleActionInterface::scrollRightAction().toUtf8().data();
}

char* QAccessibleActionInterface_QAccessibleActionInterface_ScrollUpAction(){
	return QAccessibleActionInterface::scrollUpAction().toUtf8().data();
}

char* QAccessibleActionInterface_QAccessibleActionInterface_SetFocusAction(){
	return QAccessibleActionInterface::setFocusAction().toUtf8().data();
}

char* QAccessibleActionInterface_QAccessibleActionInterface_ShowMenuAction(){
	return QAccessibleActionInterface::showMenuAction().toUtf8().data();
}

char* QAccessibleActionInterface_QAccessibleActionInterface_ToggleAction(){
	return QAccessibleActionInterface::toggleAction().toUtf8().data();
}

void QAccessibleActionInterface_DestroyQAccessibleActionInterface(QtObjectPtr ptr){
	static_cast<QAccessibleActionInterface*>(ptr)->~QAccessibleActionInterface();
}

