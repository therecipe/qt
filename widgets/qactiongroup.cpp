#include "qactiongroup.h"
#include <QIcon>
#include <QAction>
#include <QMetaObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QActionGroup>
#include "_cgo_export.h"

class MyQActionGroup: public QActionGroup {
public:
void Signal_Hovered(QAction * action){callbackQActionGroupHovered(this->objectName().toUtf8().data(), action);};
void Signal_Triggered(QAction * action){callbackQActionGroupTriggered(this->objectName().toUtf8().data(), action);};
};

QtObjectPtr QActionGroup_AddAction(QtObjectPtr ptr, QtObjectPtr action){
	return static_cast<QActionGroup*>(ptr)->addAction(static_cast<QAction*>(action));
}

int QActionGroup_IsEnabled(QtObjectPtr ptr){
	return static_cast<QActionGroup*>(ptr)->isEnabled();
}

int QActionGroup_IsExclusive(QtObjectPtr ptr){
	return static_cast<QActionGroup*>(ptr)->isExclusive();
}

int QActionGroup_IsVisible(QtObjectPtr ptr){
	return static_cast<QActionGroup*>(ptr)->isVisible();
}

void QActionGroup_SetEnabled(QtObjectPtr ptr, int v){
	QMetaObject::invokeMethod(static_cast<QActionGroup*>(ptr), "setEnabled", Q_ARG(bool, v != 0));
}

void QActionGroup_SetExclusive(QtObjectPtr ptr, int v){
	QMetaObject::invokeMethod(static_cast<QActionGroup*>(ptr), "setExclusive", Q_ARG(bool, v != 0));
}

void QActionGroup_SetVisible(QtObjectPtr ptr, int v){
	QMetaObject::invokeMethod(static_cast<QActionGroup*>(ptr), "setVisible", Q_ARG(bool, v != 0));
}

QtObjectPtr QActionGroup_NewQActionGroup(QtObjectPtr parent){
	return new QActionGroup(static_cast<QObject*>(parent));
}

QtObjectPtr QActionGroup_AddAction3(QtObjectPtr ptr, QtObjectPtr icon, char* text){
	return static_cast<QActionGroup*>(ptr)->addAction(*static_cast<QIcon*>(icon), QString(text));
}

QtObjectPtr QActionGroup_AddAction2(QtObjectPtr ptr, char* text){
	return static_cast<QActionGroup*>(ptr)->addAction(QString(text));
}

QtObjectPtr QActionGroup_CheckedAction(QtObjectPtr ptr){
	return static_cast<QActionGroup*>(ptr)->checkedAction();
}

void QActionGroup_ConnectHovered(QtObjectPtr ptr){
	QObject::connect(static_cast<QActionGroup*>(ptr), static_cast<void (QActionGroup::*)(QAction *)>(&QActionGroup::hovered), static_cast<MyQActionGroup*>(ptr), static_cast<void (MyQActionGroup::*)(QAction *)>(&MyQActionGroup::Signal_Hovered));;
}

void QActionGroup_DisconnectHovered(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QActionGroup*>(ptr), static_cast<void (QActionGroup::*)(QAction *)>(&QActionGroup::hovered), static_cast<MyQActionGroup*>(ptr), static_cast<void (MyQActionGroup::*)(QAction *)>(&MyQActionGroup::Signal_Hovered));;
}

void QActionGroup_RemoveAction(QtObjectPtr ptr, QtObjectPtr action){
	static_cast<QActionGroup*>(ptr)->removeAction(static_cast<QAction*>(action));
}

void QActionGroup_SetDisabled(QtObjectPtr ptr, int b){
	QMetaObject::invokeMethod(static_cast<QActionGroup*>(ptr), "setDisabled", Q_ARG(bool, b != 0));
}

void QActionGroup_ConnectTriggered(QtObjectPtr ptr){
	QObject::connect(static_cast<QActionGroup*>(ptr), static_cast<void (QActionGroup::*)(QAction *)>(&QActionGroup::triggered), static_cast<MyQActionGroup*>(ptr), static_cast<void (MyQActionGroup::*)(QAction *)>(&MyQActionGroup::Signal_Triggered));;
}

void QActionGroup_DisconnectTriggered(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QActionGroup*>(ptr), static_cast<void (QActionGroup::*)(QAction *)>(&QActionGroup::triggered), static_cast<MyQActionGroup*>(ptr), static_cast<void (MyQActionGroup::*)(QAction *)>(&MyQActionGroup::Signal_Triggered));;
}

void QActionGroup_DestroyQActionGroup(QtObjectPtr ptr){
	static_cast<QActionGroup*>(ptr)->~QActionGroup();
}

