#include "qmenu.h"
#include <QVariant>
#include <QModelIndex>
#include <QWidget>
#include <QIcon>
#include <QPoint>
#include <QString>
#include <QUrl>
#include <QObject>
#include <QKeySequence>
#include <QAction>
#include <QMenu>
#include "_cgo_export.h"

class MyQMenu: public QMenu {
public:
void Signal_AboutToHide(){callbackQMenuAboutToHide(this->objectName().toUtf8().data());};
void Signal_AboutToShow(){callbackQMenuAboutToShow(this->objectName().toUtf8().data());};
void Signal_Hovered(QAction * action){callbackQMenuHovered(this->objectName().toUtf8().data(), action);};
void Signal_Triggered(QAction * action){callbackQMenuTriggered(this->objectName().toUtf8().data(), action);};
};

int QMenu_IsTearOffEnabled(void* ptr){
	return static_cast<QMenu*>(ptr)->isTearOffEnabled();
}

int QMenu_SeparatorsCollapsible(void* ptr){
	return static_cast<QMenu*>(ptr)->separatorsCollapsible();
}

void QMenu_SetIcon(void* ptr, void* icon){
	static_cast<QMenu*>(ptr)->setIcon(*static_cast<QIcon*>(icon));
}

void QMenu_SetSeparatorsCollapsible(void* ptr, int collapse){
	static_cast<QMenu*>(ptr)->setSeparatorsCollapsible(collapse != 0);
}

void QMenu_SetTearOffEnabled(void* ptr, int v){
	static_cast<QMenu*>(ptr)->setTearOffEnabled(v != 0);
}

void QMenu_SetTitle(void* ptr, char* title){
	static_cast<QMenu*>(ptr)->setTitle(QString(title));
}

void QMenu_SetToolTipsVisible(void* ptr, int visible){
	static_cast<QMenu*>(ptr)->setToolTipsVisible(visible != 0);
}

char* QMenu_Title(void* ptr){
	return static_cast<QMenu*>(ptr)->title().toUtf8().data();
}

int QMenu_ToolTipsVisible(void* ptr){
	return static_cast<QMenu*>(ptr)->toolTipsVisible();
}

void* QMenu_NewQMenu(void* parent){
	return new QMenu(static_cast<QWidget*>(parent));
}

void* QMenu_NewQMenu2(char* title, void* parent){
	return new QMenu(QString(title), static_cast<QWidget*>(parent));
}

void QMenu_ConnectAboutToHide(void* ptr){
	QObject::connect(static_cast<QMenu*>(ptr), static_cast<void (QMenu::*)()>(&QMenu::aboutToHide), static_cast<MyQMenu*>(ptr), static_cast<void (MyQMenu::*)()>(&MyQMenu::Signal_AboutToHide));;
}

void QMenu_DisconnectAboutToHide(void* ptr){
	QObject::disconnect(static_cast<QMenu*>(ptr), static_cast<void (QMenu::*)()>(&QMenu::aboutToHide), static_cast<MyQMenu*>(ptr), static_cast<void (MyQMenu::*)()>(&MyQMenu::Signal_AboutToHide));;
}

void QMenu_ConnectAboutToShow(void* ptr){
	QObject::connect(static_cast<QMenu*>(ptr), static_cast<void (QMenu::*)()>(&QMenu::aboutToShow), static_cast<MyQMenu*>(ptr), static_cast<void (MyQMenu::*)()>(&MyQMenu::Signal_AboutToShow));;
}

void QMenu_DisconnectAboutToShow(void* ptr){
	QObject::disconnect(static_cast<QMenu*>(ptr), static_cast<void (QMenu::*)()>(&QMenu::aboutToShow), static_cast<MyQMenu*>(ptr), static_cast<void (MyQMenu::*)()>(&MyQMenu::Signal_AboutToShow));;
}

void* QMenu_ActionAt(void* ptr, void* pt){
	return static_cast<QMenu*>(ptr)->actionAt(*static_cast<QPoint*>(pt));
}

void* QMenu_ActiveAction(void* ptr){
	return static_cast<QMenu*>(ptr)->activeAction();
}

void* QMenu_AddAction2(void* ptr, void* icon, char* text){
	return static_cast<QMenu*>(ptr)->addAction(*static_cast<QIcon*>(icon), QString(text));
}

void* QMenu_AddAction4(void* ptr, void* icon, char* text, void* receiver, char* member, void* shortcut){
	return static_cast<QMenu*>(ptr)->addAction(*static_cast<QIcon*>(icon), QString(text), static_cast<QObject*>(receiver), const_cast<const char*>(member), *static_cast<QKeySequence*>(shortcut));
}

void* QMenu_AddAction(void* ptr, char* text){
	return static_cast<QMenu*>(ptr)->addAction(QString(text));
}

void* QMenu_AddAction3(void* ptr, char* text, void* receiver, char* member, void* shortcut){
	return static_cast<QMenu*>(ptr)->addAction(QString(text), static_cast<QObject*>(receiver), const_cast<const char*>(member), *static_cast<QKeySequence*>(shortcut));
}

void* QMenu_AddMenu(void* ptr, void* menu){
	return static_cast<QMenu*>(ptr)->addMenu(static_cast<QMenu*>(menu));
}

void* QMenu_AddMenu3(void* ptr, void* icon, char* title){
	return static_cast<QMenu*>(ptr)->addMenu(*static_cast<QIcon*>(icon), QString(title));
}

void* QMenu_AddMenu2(void* ptr, char* title){
	return static_cast<QMenu*>(ptr)->addMenu(QString(title));
}

void* QMenu_AddSection2(void* ptr, void* icon, char* text){
	return static_cast<QMenu*>(ptr)->addSection(*static_cast<QIcon*>(icon), QString(text));
}

void* QMenu_AddSection(void* ptr, char* text){
	return static_cast<QMenu*>(ptr)->addSection(QString(text));
}

void* QMenu_AddSeparator(void* ptr){
	return static_cast<QMenu*>(ptr)->addSeparator();
}

void QMenu_Clear(void* ptr){
	static_cast<QMenu*>(ptr)->clear();
}

void* QMenu_Exec(void* ptr){
	return static_cast<QMenu*>(ptr)->exec();
}

void* QMenu_Exec2(void* ptr, void* p, void* action){
	return static_cast<QMenu*>(ptr)->exec(*static_cast<QPoint*>(p), static_cast<QAction*>(action));
}

void QMenu_HideTearOffMenu(void* ptr){
	static_cast<QMenu*>(ptr)->hideTearOffMenu();
}

void QMenu_ConnectHovered(void* ptr){
	QObject::connect(static_cast<QMenu*>(ptr), static_cast<void (QMenu::*)(QAction *)>(&QMenu::hovered), static_cast<MyQMenu*>(ptr), static_cast<void (MyQMenu::*)(QAction *)>(&MyQMenu::Signal_Hovered));;
}

void QMenu_DisconnectHovered(void* ptr){
	QObject::disconnect(static_cast<QMenu*>(ptr), static_cast<void (QMenu::*)(QAction *)>(&QMenu::hovered), static_cast<MyQMenu*>(ptr), static_cast<void (MyQMenu::*)(QAction *)>(&MyQMenu::Signal_Hovered));;
}

void* QMenu_InsertMenu(void* ptr, void* before, void* menu){
	return static_cast<QMenu*>(ptr)->insertMenu(static_cast<QAction*>(before), static_cast<QMenu*>(menu));
}

void* QMenu_InsertSection2(void* ptr, void* before, void* icon, char* text){
	return static_cast<QMenu*>(ptr)->insertSection(static_cast<QAction*>(before), *static_cast<QIcon*>(icon), QString(text));
}

void* QMenu_InsertSection(void* ptr, void* before, char* text){
	return static_cast<QMenu*>(ptr)->insertSection(static_cast<QAction*>(before), QString(text));
}

void* QMenu_InsertSeparator(void* ptr, void* before){
	return static_cast<QMenu*>(ptr)->insertSeparator(static_cast<QAction*>(before));
}

int QMenu_IsEmpty(void* ptr){
	return static_cast<QMenu*>(ptr)->isEmpty();
}

int QMenu_IsTearOffMenuVisible(void* ptr){
	return static_cast<QMenu*>(ptr)->isTearOffMenuVisible();
}

void* QMenu_MenuAction(void* ptr){
	return static_cast<QMenu*>(ptr)->menuAction();
}

void QMenu_Popup(void* ptr, void* p, void* atAction){
	static_cast<QMenu*>(ptr)->popup(*static_cast<QPoint*>(p), static_cast<QAction*>(atAction));
}

void QMenu_SetActiveAction(void* ptr, void* act){
	static_cast<QMenu*>(ptr)->setActiveAction(static_cast<QAction*>(act));
}

void QMenu_ConnectTriggered(void* ptr){
	QObject::connect(static_cast<QMenu*>(ptr), static_cast<void (QMenu::*)(QAction *)>(&QMenu::triggered), static_cast<MyQMenu*>(ptr), static_cast<void (MyQMenu::*)(QAction *)>(&MyQMenu::Signal_Triggered));;
}

void QMenu_DisconnectTriggered(void* ptr){
	QObject::disconnect(static_cast<QMenu*>(ptr), static_cast<void (QMenu::*)(QAction *)>(&QMenu::triggered), static_cast<MyQMenu*>(ptr), static_cast<void (MyQMenu::*)(QAction *)>(&MyQMenu::Signal_Triggered));;
}

void QMenu_DestroyQMenu(void* ptr){
	static_cast<QMenu*>(ptr)->~QMenu();
}

