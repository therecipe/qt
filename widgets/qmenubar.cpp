#include "qmenubar.h"
#include <QMenu>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QAction>
#include <QWidget>
#include <QModelIndex>
#include <QMetaObject>
#include <QObject>
#include <QPoint>
#include <QIcon>
#include <QMenuBar>
#include "_cgo_export.h"

class MyQMenuBar: public QMenuBar {
public:
void Signal_Hovered(QAction * action){callbackQMenuBarHovered(this->objectName().toUtf8().data(), action);};
void Signal_Triggered(QAction * action){callbackQMenuBarTriggered(this->objectName().toUtf8().data(), action);};
};

int QMenuBar_IsDefaultUp(QtObjectPtr ptr){
	return static_cast<QMenuBar*>(ptr)->isDefaultUp();
}

int QMenuBar_IsNativeMenuBar(QtObjectPtr ptr){
	return static_cast<QMenuBar*>(ptr)->isNativeMenuBar();
}

void QMenuBar_SetCornerWidget(QtObjectPtr ptr, QtObjectPtr widget, int corner){
	static_cast<QMenuBar*>(ptr)->setCornerWidget(static_cast<QWidget*>(widget), static_cast<Qt::Corner>(corner));
}

void QMenuBar_SetDefaultUp(QtObjectPtr ptr, int v){
	static_cast<QMenuBar*>(ptr)->setDefaultUp(v != 0);
}

void QMenuBar_SetNativeMenuBar(QtObjectPtr ptr, int nativeMenuBar){
	static_cast<QMenuBar*>(ptr)->setNativeMenuBar(nativeMenuBar != 0);
}

QtObjectPtr QMenuBar_NewQMenuBar(QtObjectPtr parent){
	return new QMenuBar(static_cast<QWidget*>(parent));
}

QtObjectPtr QMenuBar_ActionAt(QtObjectPtr ptr, QtObjectPtr pt){
	return static_cast<QMenuBar*>(ptr)->actionAt(*static_cast<QPoint*>(pt));
}

QtObjectPtr QMenuBar_ActiveAction(QtObjectPtr ptr){
	return static_cast<QMenuBar*>(ptr)->activeAction();
}

QtObjectPtr QMenuBar_AddAction(QtObjectPtr ptr, char* text){
	return static_cast<QMenuBar*>(ptr)->addAction(QString(text));
}

QtObjectPtr QMenuBar_AddAction2(QtObjectPtr ptr, char* text, QtObjectPtr receiver, char* member){
	return static_cast<QMenuBar*>(ptr)->addAction(QString(text), static_cast<QObject*>(receiver), const_cast<const char*>(member));
}

QtObjectPtr QMenuBar_AddMenu(QtObjectPtr ptr, QtObjectPtr menu){
	return static_cast<QMenuBar*>(ptr)->addMenu(static_cast<QMenu*>(menu));
}

QtObjectPtr QMenuBar_AddMenu3(QtObjectPtr ptr, QtObjectPtr icon, char* title){
	return static_cast<QMenuBar*>(ptr)->addMenu(*static_cast<QIcon*>(icon), QString(title));
}

QtObjectPtr QMenuBar_AddMenu2(QtObjectPtr ptr, char* title){
	return static_cast<QMenuBar*>(ptr)->addMenu(QString(title));
}

QtObjectPtr QMenuBar_AddSeparator(QtObjectPtr ptr){
	return static_cast<QMenuBar*>(ptr)->addSeparator();
}

void QMenuBar_Clear(QtObjectPtr ptr){
	static_cast<QMenuBar*>(ptr)->clear();
}

QtObjectPtr QMenuBar_CornerWidget(QtObjectPtr ptr, int corner){
	return static_cast<QMenuBar*>(ptr)->cornerWidget(static_cast<Qt::Corner>(corner));
}

int QMenuBar_HeightForWidth(QtObjectPtr ptr, int v){
	return static_cast<QMenuBar*>(ptr)->heightForWidth(v);
}

void QMenuBar_ConnectHovered(QtObjectPtr ptr){
	QObject::connect(static_cast<QMenuBar*>(ptr), static_cast<void (QMenuBar::*)(QAction *)>(&QMenuBar::hovered), static_cast<MyQMenuBar*>(ptr), static_cast<void (MyQMenuBar::*)(QAction *)>(&MyQMenuBar::Signal_Hovered));;
}

void QMenuBar_DisconnectHovered(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMenuBar*>(ptr), static_cast<void (QMenuBar::*)(QAction *)>(&QMenuBar::hovered), static_cast<MyQMenuBar*>(ptr), static_cast<void (MyQMenuBar::*)(QAction *)>(&MyQMenuBar::Signal_Hovered));;
}

QtObjectPtr QMenuBar_InsertMenu(QtObjectPtr ptr, QtObjectPtr before, QtObjectPtr menu){
	return static_cast<QMenuBar*>(ptr)->insertMenu(static_cast<QAction*>(before), static_cast<QMenu*>(menu));
}

QtObjectPtr QMenuBar_InsertSeparator(QtObjectPtr ptr, QtObjectPtr before){
	return static_cast<QMenuBar*>(ptr)->insertSeparator(static_cast<QAction*>(before));
}

void QMenuBar_SetActiveAction(QtObjectPtr ptr, QtObjectPtr act){
	static_cast<QMenuBar*>(ptr)->setActiveAction(static_cast<QAction*>(act));
}

void QMenuBar_SetVisible(QtObjectPtr ptr, int visible){
	QMetaObject::invokeMethod(static_cast<QMenuBar*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void QMenuBar_ConnectTriggered(QtObjectPtr ptr){
	QObject::connect(static_cast<QMenuBar*>(ptr), static_cast<void (QMenuBar::*)(QAction *)>(&QMenuBar::triggered), static_cast<MyQMenuBar*>(ptr), static_cast<void (MyQMenuBar::*)(QAction *)>(&MyQMenuBar::Signal_Triggered));;
}

void QMenuBar_DisconnectTriggered(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMenuBar*>(ptr), static_cast<void (QMenuBar::*)(QAction *)>(&QMenuBar::triggered), static_cast<MyQMenuBar*>(ptr), static_cast<void (MyQMenuBar::*)(QAction *)>(&MyQMenuBar::Signal_Triggered));;
}

void QMenuBar_DestroyQMenuBar(QtObjectPtr ptr){
	static_cast<QMenuBar*>(ptr)->~QMenuBar();
}

