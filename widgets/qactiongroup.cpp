#include "qactiongroup.h"
#include <QModelIndex>
#include <QMetaObject>
#include <QObject>
#include <QIcon>
#include <QAction>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QActionGroup>
#include "_cgo_export.h"

class MyQActionGroup: public QActionGroup {
public:
void Signal_Hovered(QAction * action){callbackQActionGroupHovered(this->objectName().toUtf8().data(), action);};
void Signal_Triggered(QAction * action){callbackQActionGroupTriggered(this->objectName().toUtf8().data(), action);};
};

void* QActionGroup_AddAction(void* ptr, void* action){
	return static_cast<QActionGroup*>(ptr)->addAction(static_cast<QAction*>(action));
}

int QActionGroup_IsEnabled(void* ptr){
	return static_cast<QActionGroup*>(ptr)->isEnabled();
}

int QActionGroup_IsExclusive(void* ptr){
	return static_cast<QActionGroup*>(ptr)->isExclusive();
}

int QActionGroup_IsVisible(void* ptr){
	return static_cast<QActionGroup*>(ptr)->isVisible();
}

void QActionGroup_SetEnabled(void* ptr, int v){
	QMetaObject::invokeMethod(static_cast<QActionGroup*>(ptr), "setEnabled", Q_ARG(bool, v != 0));
}

void QActionGroup_SetExclusive(void* ptr, int v){
	QMetaObject::invokeMethod(static_cast<QActionGroup*>(ptr), "setExclusive", Q_ARG(bool, v != 0));
}

void QActionGroup_SetVisible(void* ptr, int v){
	QMetaObject::invokeMethod(static_cast<QActionGroup*>(ptr), "setVisible", Q_ARG(bool, v != 0));
}

void* QActionGroup_NewQActionGroup(void* parent){
	return new QActionGroup(static_cast<QObject*>(parent));
}

void* QActionGroup_AddAction3(void* ptr, void* icon, char* text){
	return static_cast<QActionGroup*>(ptr)->addAction(*static_cast<QIcon*>(icon), QString(text));
}

void* QActionGroup_AddAction2(void* ptr, char* text){
	return static_cast<QActionGroup*>(ptr)->addAction(QString(text));
}

void* QActionGroup_CheckedAction(void* ptr){
	return static_cast<QActionGroup*>(ptr)->checkedAction();
}

void QActionGroup_ConnectHovered(void* ptr){
	QObject::connect(static_cast<QActionGroup*>(ptr), static_cast<void (QActionGroup::*)(QAction *)>(&QActionGroup::hovered), static_cast<MyQActionGroup*>(ptr), static_cast<void (MyQActionGroup::*)(QAction *)>(&MyQActionGroup::Signal_Hovered));;
}

void QActionGroup_DisconnectHovered(void* ptr){
	QObject::disconnect(static_cast<QActionGroup*>(ptr), static_cast<void (QActionGroup::*)(QAction *)>(&QActionGroup::hovered), static_cast<MyQActionGroup*>(ptr), static_cast<void (MyQActionGroup::*)(QAction *)>(&MyQActionGroup::Signal_Hovered));;
}

void QActionGroup_RemoveAction(void* ptr, void* action){
	static_cast<QActionGroup*>(ptr)->removeAction(static_cast<QAction*>(action));
}

void QActionGroup_SetDisabled(void* ptr, int b){
	QMetaObject::invokeMethod(static_cast<QActionGroup*>(ptr), "setDisabled", Q_ARG(bool, b != 0));
}

void QActionGroup_ConnectTriggered(void* ptr){
	QObject::connect(static_cast<QActionGroup*>(ptr), static_cast<void (QActionGroup::*)(QAction *)>(&QActionGroup::triggered), static_cast<MyQActionGroup*>(ptr), static_cast<void (MyQActionGroup::*)(QAction *)>(&MyQActionGroup::Signal_Triggered));;
}

void QActionGroup_DisconnectTriggered(void* ptr){
	QObject::disconnect(static_cast<QActionGroup*>(ptr), static_cast<void (QActionGroup::*)(QAction *)>(&QActionGroup::triggered), static_cast<MyQActionGroup*>(ptr), static_cast<void (MyQActionGroup::*)(QAction *)>(&MyQActionGroup::Signal_Triggered));;
}

void QActionGroup_DestroyQActionGroup(void* ptr){
	static_cast<QActionGroup*>(ptr)->~QActionGroup();
}

