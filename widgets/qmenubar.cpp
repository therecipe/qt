#include "qmenubar.h"
#include <QWidget>
#include <QPoint>
#include <QMenu>
#include <QUrl>
#include <QIcon>
#include <QModelIndex>
#include <QAction>
#include <QObject>
#include <QMetaObject>
#include <QString>
#include <QVariant>
#include <QMenuBar>
#include "_cgo_export.h"

class MyQMenuBar: public QMenuBar {
public:
void Signal_Hovered(QAction * action){callbackQMenuBarHovered(this->objectName().toUtf8().data(), action);};
void Signal_Triggered(QAction * action){callbackQMenuBarTriggered(this->objectName().toUtf8().data(), action);};
};

int QMenuBar_IsDefaultUp(void* ptr){
	return static_cast<QMenuBar*>(ptr)->isDefaultUp();
}

int QMenuBar_IsNativeMenuBar(void* ptr){
	return static_cast<QMenuBar*>(ptr)->isNativeMenuBar();
}

void QMenuBar_SetCornerWidget(void* ptr, void* widget, int corner){
	static_cast<QMenuBar*>(ptr)->setCornerWidget(static_cast<QWidget*>(widget), static_cast<Qt::Corner>(corner));
}

void QMenuBar_SetDefaultUp(void* ptr, int v){
	static_cast<QMenuBar*>(ptr)->setDefaultUp(v != 0);
}

void QMenuBar_SetNativeMenuBar(void* ptr, int nativeMenuBar){
	static_cast<QMenuBar*>(ptr)->setNativeMenuBar(nativeMenuBar != 0);
}

void* QMenuBar_NewQMenuBar(void* parent){
	return new QMenuBar(static_cast<QWidget*>(parent));
}

void* QMenuBar_ActionAt(void* ptr, void* pt){
	return static_cast<QMenuBar*>(ptr)->actionAt(*static_cast<QPoint*>(pt));
}

void* QMenuBar_ActiveAction(void* ptr){
	return static_cast<QMenuBar*>(ptr)->activeAction();
}

void* QMenuBar_AddAction(void* ptr, char* text){
	return static_cast<QMenuBar*>(ptr)->addAction(QString(text));
}

void* QMenuBar_AddAction2(void* ptr, char* text, void* receiver, char* member){
	return static_cast<QMenuBar*>(ptr)->addAction(QString(text), static_cast<QObject*>(receiver), const_cast<const char*>(member));
}

void* QMenuBar_AddMenu(void* ptr, void* menu){
	return static_cast<QMenuBar*>(ptr)->addMenu(static_cast<QMenu*>(menu));
}

void* QMenuBar_AddMenu3(void* ptr, void* icon, char* title){
	return static_cast<QMenuBar*>(ptr)->addMenu(*static_cast<QIcon*>(icon), QString(title));
}

void* QMenuBar_AddMenu2(void* ptr, char* title){
	return static_cast<QMenuBar*>(ptr)->addMenu(QString(title));
}

void* QMenuBar_AddSeparator(void* ptr){
	return static_cast<QMenuBar*>(ptr)->addSeparator();
}

void QMenuBar_Clear(void* ptr){
	static_cast<QMenuBar*>(ptr)->clear();
}

void* QMenuBar_CornerWidget(void* ptr, int corner){
	return static_cast<QMenuBar*>(ptr)->cornerWidget(static_cast<Qt::Corner>(corner));
}

int QMenuBar_HeightForWidth(void* ptr, int v){
	return static_cast<QMenuBar*>(ptr)->heightForWidth(v);
}

void QMenuBar_ConnectHovered(void* ptr){
	QObject::connect(static_cast<QMenuBar*>(ptr), static_cast<void (QMenuBar::*)(QAction *)>(&QMenuBar::hovered), static_cast<MyQMenuBar*>(ptr), static_cast<void (MyQMenuBar::*)(QAction *)>(&MyQMenuBar::Signal_Hovered));;
}

void QMenuBar_DisconnectHovered(void* ptr){
	QObject::disconnect(static_cast<QMenuBar*>(ptr), static_cast<void (QMenuBar::*)(QAction *)>(&QMenuBar::hovered), static_cast<MyQMenuBar*>(ptr), static_cast<void (MyQMenuBar::*)(QAction *)>(&MyQMenuBar::Signal_Hovered));;
}

void* QMenuBar_InsertMenu(void* ptr, void* before, void* menu){
	return static_cast<QMenuBar*>(ptr)->insertMenu(static_cast<QAction*>(before), static_cast<QMenu*>(menu));
}

void* QMenuBar_InsertSeparator(void* ptr, void* before){
	return static_cast<QMenuBar*>(ptr)->insertSeparator(static_cast<QAction*>(before));
}

void QMenuBar_SetActiveAction(void* ptr, void* act){
	static_cast<QMenuBar*>(ptr)->setActiveAction(static_cast<QAction*>(act));
}

void QMenuBar_SetVisible(void* ptr, int visible){
	QMetaObject::invokeMethod(static_cast<QMenuBar*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void QMenuBar_ConnectTriggered(void* ptr){
	QObject::connect(static_cast<QMenuBar*>(ptr), static_cast<void (QMenuBar::*)(QAction *)>(&QMenuBar::triggered), static_cast<MyQMenuBar*>(ptr), static_cast<void (MyQMenuBar::*)(QAction *)>(&MyQMenuBar::Signal_Triggered));;
}

void QMenuBar_DisconnectTriggered(void* ptr){
	QObject::disconnect(static_cast<QMenuBar*>(ptr), static_cast<void (QMenuBar::*)(QAction *)>(&QMenuBar::triggered), static_cast<MyQMenuBar*>(ptr), static_cast<void (MyQMenuBar::*)(QAction *)>(&MyQMenuBar::Signal_Triggered));;
}

void QMenuBar_DestroyQMenuBar(void* ptr){
	static_cast<QMenuBar*>(ptr)->~QMenuBar();
}

