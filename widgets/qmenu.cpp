#include "qmenu.h"
#include <QUrl>
#include <QIcon>
#include <QString>
#include <QVariant>
#include <QObject>
#include <QPoint>
#include <QAction>
#include <QWidget>
#include <QModelIndex>
#include <QKeySequence>
#include <QMenu>
#include "_cgo_export.h"

class MyQMenu: public QMenu {
public:
void Signal_AboutToHide(){callbackQMenuAboutToHide(this->objectName().toUtf8().data());};
void Signal_AboutToShow(){callbackQMenuAboutToShow(this->objectName().toUtf8().data());};
void Signal_Hovered(QAction * action){callbackQMenuHovered(this->objectName().toUtf8().data(), action);};
void Signal_Triggered(QAction * action){callbackQMenuTriggered(this->objectName().toUtf8().data(), action);};
};

int QMenu_IsTearOffEnabled(QtObjectPtr ptr){
	return static_cast<QMenu*>(ptr)->isTearOffEnabled();
}

int QMenu_SeparatorsCollapsible(QtObjectPtr ptr){
	return static_cast<QMenu*>(ptr)->separatorsCollapsible();
}

void QMenu_SetIcon(QtObjectPtr ptr, QtObjectPtr icon){
	static_cast<QMenu*>(ptr)->setIcon(*static_cast<QIcon*>(icon));
}

void QMenu_SetSeparatorsCollapsible(QtObjectPtr ptr, int collapse){
	static_cast<QMenu*>(ptr)->setSeparatorsCollapsible(collapse != 0);
}

void QMenu_SetTearOffEnabled(QtObjectPtr ptr, int v){
	static_cast<QMenu*>(ptr)->setTearOffEnabled(v != 0);
}

void QMenu_SetTitle(QtObjectPtr ptr, char* title){
	static_cast<QMenu*>(ptr)->setTitle(QString(title));
}

void QMenu_SetToolTipsVisible(QtObjectPtr ptr, int visible){
	static_cast<QMenu*>(ptr)->setToolTipsVisible(visible != 0);
}

char* QMenu_Title(QtObjectPtr ptr){
	return static_cast<QMenu*>(ptr)->title().toUtf8().data();
}

int QMenu_ToolTipsVisible(QtObjectPtr ptr){
	return static_cast<QMenu*>(ptr)->toolTipsVisible();
}

QtObjectPtr QMenu_NewQMenu(QtObjectPtr parent){
	return new QMenu(static_cast<QWidget*>(parent));
}

QtObjectPtr QMenu_NewQMenu2(char* title, QtObjectPtr parent){
	return new QMenu(QString(title), static_cast<QWidget*>(parent));
}

void QMenu_ConnectAboutToHide(QtObjectPtr ptr){
	QObject::connect(static_cast<QMenu*>(ptr), static_cast<void (QMenu::*)()>(&QMenu::aboutToHide), static_cast<MyQMenu*>(ptr), static_cast<void (MyQMenu::*)()>(&MyQMenu::Signal_AboutToHide));;
}

void QMenu_DisconnectAboutToHide(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMenu*>(ptr), static_cast<void (QMenu::*)()>(&QMenu::aboutToHide), static_cast<MyQMenu*>(ptr), static_cast<void (MyQMenu::*)()>(&MyQMenu::Signal_AboutToHide));;
}

void QMenu_ConnectAboutToShow(QtObjectPtr ptr){
	QObject::connect(static_cast<QMenu*>(ptr), static_cast<void (QMenu::*)()>(&QMenu::aboutToShow), static_cast<MyQMenu*>(ptr), static_cast<void (MyQMenu::*)()>(&MyQMenu::Signal_AboutToShow));;
}

void QMenu_DisconnectAboutToShow(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMenu*>(ptr), static_cast<void (QMenu::*)()>(&QMenu::aboutToShow), static_cast<MyQMenu*>(ptr), static_cast<void (MyQMenu::*)()>(&MyQMenu::Signal_AboutToShow));;
}

QtObjectPtr QMenu_ActionAt(QtObjectPtr ptr, QtObjectPtr pt){
	return static_cast<QMenu*>(ptr)->actionAt(*static_cast<QPoint*>(pt));
}

QtObjectPtr QMenu_ActiveAction(QtObjectPtr ptr){
	return static_cast<QMenu*>(ptr)->activeAction();
}

QtObjectPtr QMenu_AddAction2(QtObjectPtr ptr, QtObjectPtr icon, char* text){
	return static_cast<QMenu*>(ptr)->addAction(*static_cast<QIcon*>(icon), QString(text));
}

QtObjectPtr QMenu_AddAction4(QtObjectPtr ptr, QtObjectPtr icon, char* text, QtObjectPtr receiver, char* member, QtObjectPtr shortcut){
	return static_cast<QMenu*>(ptr)->addAction(*static_cast<QIcon*>(icon), QString(text), static_cast<QObject*>(receiver), const_cast<const char*>(member), *static_cast<QKeySequence*>(shortcut));
}

QtObjectPtr QMenu_AddAction(QtObjectPtr ptr, char* text){
	return static_cast<QMenu*>(ptr)->addAction(QString(text));
}

QtObjectPtr QMenu_AddAction3(QtObjectPtr ptr, char* text, QtObjectPtr receiver, char* member, QtObjectPtr shortcut){
	return static_cast<QMenu*>(ptr)->addAction(QString(text), static_cast<QObject*>(receiver), const_cast<const char*>(member), *static_cast<QKeySequence*>(shortcut));
}

QtObjectPtr QMenu_AddMenu(QtObjectPtr ptr, QtObjectPtr menu){
	return static_cast<QMenu*>(ptr)->addMenu(static_cast<QMenu*>(menu));
}

QtObjectPtr QMenu_AddMenu3(QtObjectPtr ptr, QtObjectPtr icon, char* title){
	return static_cast<QMenu*>(ptr)->addMenu(*static_cast<QIcon*>(icon), QString(title));
}

QtObjectPtr QMenu_AddMenu2(QtObjectPtr ptr, char* title){
	return static_cast<QMenu*>(ptr)->addMenu(QString(title));
}

QtObjectPtr QMenu_AddSection2(QtObjectPtr ptr, QtObjectPtr icon, char* text){
	return static_cast<QMenu*>(ptr)->addSection(*static_cast<QIcon*>(icon), QString(text));
}

QtObjectPtr QMenu_AddSection(QtObjectPtr ptr, char* text){
	return static_cast<QMenu*>(ptr)->addSection(QString(text));
}

QtObjectPtr QMenu_AddSeparator(QtObjectPtr ptr){
	return static_cast<QMenu*>(ptr)->addSeparator();
}

void QMenu_Clear(QtObjectPtr ptr){
	static_cast<QMenu*>(ptr)->clear();
}

QtObjectPtr QMenu_Exec(QtObjectPtr ptr){
	return static_cast<QMenu*>(ptr)->exec();
}

QtObjectPtr QMenu_Exec2(QtObjectPtr ptr, QtObjectPtr p, QtObjectPtr action){
	return static_cast<QMenu*>(ptr)->exec(*static_cast<QPoint*>(p), static_cast<QAction*>(action));
}

void QMenu_HideTearOffMenu(QtObjectPtr ptr){
	static_cast<QMenu*>(ptr)->hideTearOffMenu();
}

void QMenu_ConnectHovered(QtObjectPtr ptr){
	QObject::connect(static_cast<QMenu*>(ptr), static_cast<void (QMenu::*)(QAction *)>(&QMenu::hovered), static_cast<MyQMenu*>(ptr), static_cast<void (MyQMenu::*)(QAction *)>(&MyQMenu::Signal_Hovered));;
}

void QMenu_DisconnectHovered(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMenu*>(ptr), static_cast<void (QMenu::*)(QAction *)>(&QMenu::hovered), static_cast<MyQMenu*>(ptr), static_cast<void (MyQMenu::*)(QAction *)>(&MyQMenu::Signal_Hovered));;
}

QtObjectPtr QMenu_InsertMenu(QtObjectPtr ptr, QtObjectPtr before, QtObjectPtr menu){
	return static_cast<QMenu*>(ptr)->insertMenu(static_cast<QAction*>(before), static_cast<QMenu*>(menu));
}

QtObjectPtr QMenu_InsertSection2(QtObjectPtr ptr, QtObjectPtr before, QtObjectPtr icon, char* text){
	return static_cast<QMenu*>(ptr)->insertSection(static_cast<QAction*>(before), *static_cast<QIcon*>(icon), QString(text));
}

QtObjectPtr QMenu_InsertSection(QtObjectPtr ptr, QtObjectPtr before, char* text){
	return static_cast<QMenu*>(ptr)->insertSection(static_cast<QAction*>(before), QString(text));
}

QtObjectPtr QMenu_InsertSeparator(QtObjectPtr ptr, QtObjectPtr before){
	return static_cast<QMenu*>(ptr)->insertSeparator(static_cast<QAction*>(before));
}

int QMenu_IsEmpty(QtObjectPtr ptr){
	return static_cast<QMenu*>(ptr)->isEmpty();
}

int QMenu_IsTearOffMenuVisible(QtObjectPtr ptr){
	return static_cast<QMenu*>(ptr)->isTearOffMenuVisible();
}

QtObjectPtr QMenu_MenuAction(QtObjectPtr ptr){
	return static_cast<QMenu*>(ptr)->menuAction();
}

void QMenu_Popup(QtObjectPtr ptr, QtObjectPtr p, QtObjectPtr atAction){
	static_cast<QMenu*>(ptr)->popup(*static_cast<QPoint*>(p), static_cast<QAction*>(atAction));
}

void QMenu_SetActiveAction(QtObjectPtr ptr, QtObjectPtr act){
	static_cast<QMenu*>(ptr)->setActiveAction(static_cast<QAction*>(act));
}

void QMenu_ConnectTriggered(QtObjectPtr ptr){
	QObject::connect(static_cast<QMenu*>(ptr), static_cast<void (QMenu::*)(QAction *)>(&QMenu::triggered), static_cast<MyQMenu*>(ptr), static_cast<void (MyQMenu::*)(QAction *)>(&MyQMenu::Signal_Triggered));;
}

void QMenu_DisconnectTriggered(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMenu*>(ptr), static_cast<void (QMenu::*)(QAction *)>(&QMenu::triggered), static_cast<MyQMenu*>(ptr), static_cast<void (MyQMenu::*)(QAction *)>(&MyQMenu::Signal_Triggered));;
}

void QMenu_DestroyQMenu(QtObjectPtr ptr){
	static_cast<QMenu*>(ptr)->~QMenu();
}

