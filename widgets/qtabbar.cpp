#include "qtabbar.h"
#include <QUrl>
#include <QModelIndex>
#include <QColor>
#include <QMetaObject>
#include <QObject>
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QIcon>
#include <QSize>
#include <QPoint>
#include <QTabBar>
#include "_cgo_export.h"

class MyQTabBar: public QTabBar {
public:
void Signal_CurrentChanged(int index){callbackQTabBarCurrentChanged(this->objectName().toUtf8().data(), index);};
void Signal_TabBarClicked(int index){callbackQTabBarTabBarClicked(this->objectName().toUtf8().data(), index);};
void Signal_TabBarDoubleClicked(int index){callbackQTabBarTabBarDoubleClicked(this->objectName().toUtf8().data(), index);};
void Signal_TabCloseRequested(int index){callbackQTabBarTabCloseRequested(this->objectName().toUtf8().data(), index);};
void Signal_TabMoved(int from, int to){callbackQTabBarTabMoved(this->objectName().toUtf8().data(), from, to);};
};

int QTabBar_AutoHide(QtObjectPtr ptr){
	return static_cast<QTabBar*>(ptr)->autoHide();
}

int QTabBar_ChangeCurrentOnDrag(QtObjectPtr ptr){
	return static_cast<QTabBar*>(ptr)->changeCurrentOnDrag();
}

int QTabBar_Count(QtObjectPtr ptr){
	return static_cast<QTabBar*>(ptr)->count();
}

int QTabBar_CurrentIndex(QtObjectPtr ptr){
	return static_cast<QTabBar*>(ptr)->currentIndex();
}

int QTabBar_DocumentMode(QtObjectPtr ptr){
	return static_cast<QTabBar*>(ptr)->documentMode();
}

int QTabBar_DrawBase(QtObjectPtr ptr){
	return static_cast<QTabBar*>(ptr)->drawBase();
}

int QTabBar_ElideMode(QtObjectPtr ptr){
	return static_cast<QTabBar*>(ptr)->elideMode();
}

int QTabBar_Expanding(QtObjectPtr ptr){
	return static_cast<QTabBar*>(ptr)->expanding();
}

int QTabBar_IsMovable(QtObjectPtr ptr){
	return static_cast<QTabBar*>(ptr)->isMovable();
}

int QTabBar_SelectionBehaviorOnRemove(QtObjectPtr ptr){
	return static_cast<QTabBar*>(ptr)->selectionBehaviorOnRemove();
}

void QTabBar_SetAutoHide(QtObjectPtr ptr, int hide){
	static_cast<QTabBar*>(ptr)->setAutoHide(hide != 0);
}

void QTabBar_SetChangeCurrentOnDrag(QtObjectPtr ptr, int change){
	static_cast<QTabBar*>(ptr)->setChangeCurrentOnDrag(change != 0);
}

void QTabBar_SetCurrentIndex(QtObjectPtr ptr, int index){
	QMetaObject::invokeMethod(static_cast<QTabBar*>(ptr), "setCurrentIndex", Q_ARG(int, index));
}

void QTabBar_SetDocumentMode(QtObjectPtr ptr, int set){
	static_cast<QTabBar*>(ptr)->setDocumentMode(set != 0);
}

void QTabBar_SetDrawBase(QtObjectPtr ptr, int drawTheBase){
	static_cast<QTabBar*>(ptr)->setDrawBase(drawTheBase != 0);
}

void QTabBar_SetElideMode(QtObjectPtr ptr, int v){
	static_cast<QTabBar*>(ptr)->setElideMode(static_cast<Qt::TextElideMode>(v));
}

void QTabBar_SetExpanding(QtObjectPtr ptr, int enabled){
	static_cast<QTabBar*>(ptr)->setExpanding(enabled != 0);
}

void QTabBar_SetIconSize(QtObjectPtr ptr, QtObjectPtr size){
	static_cast<QTabBar*>(ptr)->setIconSize(*static_cast<QSize*>(size));
}

void QTabBar_SetMovable(QtObjectPtr ptr, int movable){
	static_cast<QTabBar*>(ptr)->setMovable(movable != 0);
}

void QTabBar_SetSelectionBehaviorOnRemove(QtObjectPtr ptr, int behavior){
	static_cast<QTabBar*>(ptr)->setSelectionBehaviorOnRemove(static_cast<QTabBar::SelectionBehavior>(behavior));
}

void QTabBar_SetShape(QtObjectPtr ptr, int shape){
	static_cast<QTabBar*>(ptr)->setShape(static_cast<QTabBar::Shape>(shape));
}

void QTabBar_SetTabsClosable(QtObjectPtr ptr, int closable){
	static_cast<QTabBar*>(ptr)->setTabsClosable(closable != 0);
}

void QTabBar_SetUsesScrollButtons(QtObjectPtr ptr, int useButtons){
	static_cast<QTabBar*>(ptr)->setUsesScrollButtons(useButtons != 0);
}

int QTabBar_Shape(QtObjectPtr ptr){
	return static_cast<QTabBar*>(ptr)->shape();
}

int QTabBar_TabsClosable(QtObjectPtr ptr){
	return static_cast<QTabBar*>(ptr)->tabsClosable();
}

int QTabBar_UsesScrollButtons(QtObjectPtr ptr){
	return static_cast<QTabBar*>(ptr)->usesScrollButtons();
}

QtObjectPtr QTabBar_NewQTabBar(QtObjectPtr parent){
	return new QTabBar(static_cast<QWidget*>(parent));
}

int QTabBar_AddTab2(QtObjectPtr ptr, QtObjectPtr icon, char* text){
	return static_cast<QTabBar*>(ptr)->addTab(*static_cast<QIcon*>(icon), QString(text));
}

int QTabBar_AddTab(QtObjectPtr ptr, char* text){
	return static_cast<QTabBar*>(ptr)->addTab(QString(text));
}

void QTabBar_ConnectCurrentChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QTabBar*>(ptr), static_cast<void (QTabBar::*)(int)>(&QTabBar::currentChanged), static_cast<MyQTabBar*>(ptr), static_cast<void (MyQTabBar::*)(int)>(&MyQTabBar::Signal_CurrentChanged));;
}

void QTabBar_DisconnectCurrentChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTabBar*>(ptr), static_cast<void (QTabBar::*)(int)>(&QTabBar::currentChanged), static_cast<MyQTabBar*>(ptr), static_cast<void (MyQTabBar::*)(int)>(&MyQTabBar::Signal_CurrentChanged));;
}

int QTabBar_InsertTab2(QtObjectPtr ptr, int index, QtObjectPtr icon, char* text){
	return static_cast<QTabBar*>(ptr)->insertTab(index, *static_cast<QIcon*>(icon), QString(text));
}

int QTabBar_InsertTab(QtObjectPtr ptr, int index, char* text){
	return static_cast<QTabBar*>(ptr)->insertTab(index, QString(text));
}

int QTabBar_IsTabEnabled(QtObjectPtr ptr, int index){
	return static_cast<QTabBar*>(ptr)->isTabEnabled(index);
}

void QTabBar_MoveTab(QtObjectPtr ptr, int from, int to){
	static_cast<QTabBar*>(ptr)->moveTab(from, to);
}

void QTabBar_RemoveTab(QtObjectPtr ptr, int index){
	static_cast<QTabBar*>(ptr)->removeTab(index);
}

void QTabBar_SetTabButton(QtObjectPtr ptr, int index, int position, QtObjectPtr widget){
	static_cast<QTabBar*>(ptr)->setTabButton(index, static_cast<QTabBar::ButtonPosition>(position), static_cast<QWidget*>(widget));
}

void QTabBar_SetTabData(QtObjectPtr ptr, int index, char* data){
	static_cast<QTabBar*>(ptr)->setTabData(index, QVariant(data));
}

void QTabBar_SetTabEnabled(QtObjectPtr ptr, int index, int enabled){
	static_cast<QTabBar*>(ptr)->setTabEnabled(index, enabled != 0);
}

void QTabBar_SetTabIcon(QtObjectPtr ptr, int index, QtObjectPtr icon){
	static_cast<QTabBar*>(ptr)->setTabIcon(index, *static_cast<QIcon*>(icon));
}

void QTabBar_SetTabText(QtObjectPtr ptr, int index, char* text){
	static_cast<QTabBar*>(ptr)->setTabText(index, QString(text));
}

void QTabBar_SetTabTextColor(QtObjectPtr ptr, int index, QtObjectPtr color){
	static_cast<QTabBar*>(ptr)->setTabTextColor(index, *static_cast<QColor*>(color));
}

void QTabBar_SetTabToolTip(QtObjectPtr ptr, int index, char* tip){
	static_cast<QTabBar*>(ptr)->setTabToolTip(index, QString(tip));
}

void QTabBar_SetTabWhatsThis(QtObjectPtr ptr, int index, char* text){
	static_cast<QTabBar*>(ptr)->setTabWhatsThis(index, QString(text));
}

int QTabBar_TabAt(QtObjectPtr ptr, QtObjectPtr position){
	return static_cast<QTabBar*>(ptr)->tabAt(*static_cast<QPoint*>(position));
}

void QTabBar_ConnectTabBarClicked(QtObjectPtr ptr){
	QObject::connect(static_cast<QTabBar*>(ptr), static_cast<void (QTabBar::*)(int)>(&QTabBar::tabBarClicked), static_cast<MyQTabBar*>(ptr), static_cast<void (MyQTabBar::*)(int)>(&MyQTabBar::Signal_TabBarClicked));;
}

void QTabBar_DisconnectTabBarClicked(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTabBar*>(ptr), static_cast<void (QTabBar::*)(int)>(&QTabBar::tabBarClicked), static_cast<MyQTabBar*>(ptr), static_cast<void (MyQTabBar::*)(int)>(&MyQTabBar::Signal_TabBarClicked));;
}

void QTabBar_ConnectTabBarDoubleClicked(QtObjectPtr ptr){
	QObject::connect(static_cast<QTabBar*>(ptr), static_cast<void (QTabBar::*)(int)>(&QTabBar::tabBarDoubleClicked), static_cast<MyQTabBar*>(ptr), static_cast<void (MyQTabBar::*)(int)>(&MyQTabBar::Signal_TabBarDoubleClicked));;
}

void QTabBar_DisconnectTabBarDoubleClicked(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTabBar*>(ptr), static_cast<void (QTabBar::*)(int)>(&QTabBar::tabBarDoubleClicked), static_cast<MyQTabBar*>(ptr), static_cast<void (MyQTabBar::*)(int)>(&MyQTabBar::Signal_TabBarDoubleClicked));;
}

QtObjectPtr QTabBar_TabButton(QtObjectPtr ptr, int index, int position){
	return static_cast<QTabBar*>(ptr)->tabButton(index, static_cast<QTabBar::ButtonPosition>(position));
}

void QTabBar_ConnectTabCloseRequested(QtObjectPtr ptr){
	QObject::connect(static_cast<QTabBar*>(ptr), static_cast<void (QTabBar::*)(int)>(&QTabBar::tabCloseRequested), static_cast<MyQTabBar*>(ptr), static_cast<void (MyQTabBar::*)(int)>(&MyQTabBar::Signal_TabCloseRequested));;
}

void QTabBar_DisconnectTabCloseRequested(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTabBar*>(ptr), static_cast<void (QTabBar::*)(int)>(&QTabBar::tabCloseRequested), static_cast<MyQTabBar*>(ptr), static_cast<void (MyQTabBar::*)(int)>(&MyQTabBar::Signal_TabCloseRequested));;
}

char* QTabBar_TabData(QtObjectPtr ptr, int index){
	return static_cast<QTabBar*>(ptr)->tabData(index).toString().toUtf8().data();
}

void QTabBar_ConnectTabMoved(QtObjectPtr ptr){
	QObject::connect(static_cast<QTabBar*>(ptr), static_cast<void (QTabBar::*)(int, int)>(&QTabBar::tabMoved), static_cast<MyQTabBar*>(ptr), static_cast<void (MyQTabBar::*)(int, int)>(&MyQTabBar::Signal_TabMoved));;
}

void QTabBar_DisconnectTabMoved(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTabBar*>(ptr), static_cast<void (QTabBar::*)(int, int)>(&QTabBar::tabMoved), static_cast<MyQTabBar*>(ptr), static_cast<void (MyQTabBar::*)(int, int)>(&MyQTabBar::Signal_TabMoved));;
}

char* QTabBar_TabText(QtObjectPtr ptr, int index){
	return static_cast<QTabBar*>(ptr)->tabText(index).toUtf8().data();
}

char* QTabBar_TabToolTip(QtObjectPtr ptr, int index){
	return static_cast<QTabBar*>(ptr)->tabToolTip(index).toUtf8().data();
}

char* QTabBar_TabWhatsThis(QtObjectPtr ptr, int index){
	return static_cast<QTabBar*>(ptr)->tabWhatsThis(index).toUtf8().data();
}

void QTabBar_DestroyQTabBar(QtObjectPtr ptr){
	static_cast<QTabBar*>(ptr)->~QTabBar();
}

