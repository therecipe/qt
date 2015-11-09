#include "qtabbar.h"
#include <QIcon>
#include <QModelIndex>
#include <QPoint>
#include <QColor>
#include <QSize>
#include <QObject>
#include <QMetaObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QWidget>
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

int QTabBar_AutoHide(void* ptr){
	return static_cast<QTabBar*>(ptr)->autoHide();
}

int QTabBar_ChangeCurrentOnDrag(void* ptr){
	return static_cast<QTabBar*>(ptr)->changeCurrentOnDrag();
}

int QTabBar_Count(void* ptr){
	return static_cast<QTabBar*>(ptr)->count();
}

int QTabBar_CurrentIndex(void* ptr){
	return static_cast<QTabBar*>(ptr)->currentIndex();
}

int QTabBar_DocumentMode(void* ptr){
	return static_cast<QTabBar*>(ptr)->documentMode();
}

int QTabBar_DrawBase(void* ptr){
	return static_cast<QTabBar*>(ptr)->drawBase();
}

int QTabBar_ElideMode(void* ptr){
	return static_cast<QTabBar*>(ptr)->elideMode();
}

int QTabBar_Expanding(void* ptr){
	return static_cast<QTabBar*>(ptr)->expanding();
}

int QTabBar_IsMovable(void* ptr){
	return static_cast<QTabBar*>(ptr)->isMovable();
}

int QTabBar_SelectionBehaviorOnRemove(void* ptr){
	return static_cast<QTabBar*>(ptr)->selectionBehaviorOnRemove();
}

void QTabBar_SetAutoHide(void* ptr, int hide){
	static_cast<QTabBar*>(ptr)->setAutoHide(hide != 0);
}

void QTabBar_SetChangeCurrentOnDrag(void* ptr, int change){
	static_cast<QTabBar*>(ptr)->setChangeCurrentOnDrag(change != 0);
}

void QTabBar_SetCurrentIndex(void* ptr, int index){
	QMetaObject::invokeMethod(static_cast<QTabBar*>(ptr), "setCurrentIndex", Q_ARG(int, index));
}

void QTabBar_SetDocumentMode(void* ptr, int set){
	static_cast<QTabBar*>(ptr)->setDocumentMode(set != 0);
}

void QTabBar_SetDrawBase(void* ptr, int drawTheBase){
	static_cast<QTabBar*>(ptr)->setDrawBase(drawTheBase != 0);
}

void QTabBar_SetElideMode(void* ptr, int v){
	static_cast<QTabBar*>(ptr)->setElideMode(static_cast<Qt::TextElideMode>(v));
}

void QTabBar_SetExpanding(void* ptr, int enabled){
	static_cast<QTabBar*>(ptr)->setExpanding(enabled != 0);
}

void QTabBar_SetIconSize(void* ptr, void* size){
	static_cast<QTabBar*>(ptr)->setIconSize(*static_cast<QSize*>(size));
}

void QTabBar_SetMovable(void* ptr, int movable){
	static_cast<QTabBar*>(ptr)->setMovable(movable != 0);
}

void QTabBar_SetSelectionBehaviorOnRemove(void* ptr, int behavior){
	static_cast<QTabBar*>(ptr)->setSelectionBehaviorOnRemove(static_cast<QTabBar::SelectionBehavior>(behavior));
}

void QTabBar_SetShape(void* ptr, int shape){
	static_cast<QTabBar*>(ptr)->setShape(static_cast<QTabBar::Shape>(shape));
}

void QTabBar_SetTabsClosable(void* ptr, int closable){
	static_cast<QTabBar*>(ptr)->setTabsClosable(closable != 0);
}

void QTabBar_SetUsesScrollButtons(void* ptr, int useButtons){
	static_cast<QTabBar*>(ptr)->setUsesScrollButtons(useButtons != 0);
}

int QTabBar_Shape(void* ptr){
	return static_cast<QTabBar*>(ptr)->shape();
}

int QTabBar_TabsClosable(void* ptr){
	return static_cast<QTabBar*>(ptr)->tabsClosable();
}

int QTabBar_UsesScrollButtons(void* ptr){
	return static_cast<QTabBar*>(ptr)->usesScrollButtons();
}

void* QTabBar_NewQTabBar(void* parent){
	return new QTabBar(static_cast<QWidget*>(parent));
}

int QTabBar_AddTab2(void* ptr, void* icon, char* text){
	return static_cast<QTabBar*>(ptr)->addTab(*static_cast<QIcon*>(icon), QString(text));
}

int QTabBar_AddTab(void* ptr, char* text){
	return static_cast<QTabBar*>(ptr)->addTab(QString(text));
}

void QTabBar_ConnectCurrentChanged(void* ptr){
	QObject::connect(static_cast<QTabBar*>(ptr), static_cast<void (QTabBar::*)(int)>(&QTabBar::currentChanged), static_cast<MyQTabBar*>(ptr), static_cast<void (MyQTabBar::*)(int)>(&MyQTabBar::Signal_CurrentChanged));;
}

void QTabBar_DisconnectCurrentChanged(void* ptr){
	QObject::disconnect(static_cast<QTabBar*>(ptr), static_cast<void (QTabBar::*)(int)>(&QTabBar::currentChanged), static_cast<MyQTabBar*>(ptr), static_cast<void (MyQTabBar::*)(int)>(&MyQTabBar::Signal_CurrentChanged));;
}

int QTabBar_InsertTab2(void* ptr, int index, void* icon, char* text){
	return static_cast<QTabBar*>(ptr)->insertTab(index, *static_cast<QIcon*>(icon), QString(text));
}

int QTabBar_InsertTab(void* ptr, int index, char* text){
	return static_cast<QTabBar*>(ptr)->insertTab(index, QString(text));
}

int QTabBar_IsTabEnabled(void* ptr, int index){
	return static_cast<QTabBar*>(ptr)->isTabEnabled(index);
}

void QTabBar_MoveTab(void* ptr, int from, int to){
	static_cast<QTabBar*>(ptr)->moveTab(from, to);
}

void QTabBar_RemoveTab(void* ptr, int index){
	static_cast<QTabBar*>(ptr)->removeTab(index);
}

void QTabBar_SetTabButton(void* ptr, int index, int position, void* widget){
	static_cast<QTabBar*>(ptr)->setTabButton(index, static_cast<QTabBar::ButtonPosition>(position), static_cast<QWidget*>(widget));
}

void QTabBar_SetTabData(void* ptr, int index, void* data){
	static_cast<QTabBar*>(ptr)->setTabData(index, *static_cast<QVariant*>(data));
}

void QTabBar_SetTabEnabled(void* ptr, int index, int enabled){
	static_cast<QTabBar*>(ptr)->setTabEnabled(index, enabled != 0);
}

void QTabBar_SetTabIcon(void* ptr, int index, void* icon){
	static_cast<QTabBar*>(ptr)->setTabIcon(index, *static_cast<QIcon*>(icon));
}

void QTabBar_SetTabText(void* ptr, int index, char* text){
	static_cast<QTabBar*>(ptr)->setTabText(index, QString(text));
}

void QTabBar_SetTabTextColor(void* ptr, int index, void* color){
	static_cast<QTabBar*>(ptr)->setTabTextColor(index, *static_cast<QColor*>(color));
}

void QTabBar_SetTabToolTip(void* ptr, int index, char* tip){
	static_cast<QTabBar*>(ptr)->setTabToolTip(index, QString(tip));
}

void QTabBar_SetTabWhatsThis(void* ptr, int index, char* text){
	static_cast<QTabBar*>(ptr)->setTabWhatsThis(index, QString(text));
}

int QTabBar_TabAt(void* ptr, void* position){
	return static_cast<QTabBar*>(ptr)->tabAt(*static_cast<QPoint*>(position));
}

void QTabBar_ConnectTabBarClicked(void* ptr){
	QObject::connect(static_cast<QTabBar*>(ptr), static_cast<void (QTabBar::*)(int)>(&QTabBar::tabBarClicked), static_cast<MyQTabBar*>(ptr), static_cast<void (MyQTabBar::*)(int)>(&MyQTabBar::Signal_TabBarClicked));;
}

void QTabBar_DisconnectTabBarClicked(void* ptr){
	QObject::disconnect(static_cast<QTabBar*>(ptr), static_cast<void (QTabBar::*)(int)>(&QTabBar::tabBarClicked), static_cast<MyQTabBar*>(ptr), static_cast<void (MyQTabBar::*)(int)>(&MyQTabBar::Signal_TabBarClicked));;
}

void QTabBar_ConnectTabBarDoubleClicked(void* ptr){
	QObject::connect(static_cast<QTabBar*>(ptr), static_cast<void (QTabBar::*)(int)>(&QTabBar::tabBarDoubleClicked), static_cast<MyQTabBar*>(ptr), static_cast<void (MyQTabBar::*)(int)>(&MyQTabBar::Signal_TabBarDoubleClicked));;
}

void QTabBar_DisconnectTabBarDoubleClicked(void* ptr){
	QObject::disconnect(static_cast<QTabBar*>(ptr), static_cast<void (QTabBar::*)(int)>(&QTabBar::tabBarDoubleClicked), static_cast<MyQTabBar*>(ptr), static_cast<void (MyQTabBar::*)(int)>(&MyQTabBar::Signal_TabBarDoubleClicked));;
}

void* QTabBar_TabButton(void* ptr, int index, int position){
	return static_cast<QTabBar*>(ptr)->tabButton(index, static_cast<QTabBar::ButtonPosition>(position));
}

void QTabBar_ConnectTabCloseRequested(void* ptr){
	QObject::connect(static_cast<QTabBar*>(ptr), static_cast<void (QTabBar::*)(int)>(&QTabBar::tabCloseRequested), static_cast<MyQTabBar*>(ptr), static_cast<void (MyQTabBar::*)(int)>(&MyQTabBar::Signal_TabCloseRequested));;
}

void QTabBar_DisconnectTabCloseRequested(void* ptr){
	QObject::disconnect(static_cast<QTabBar*>(ptr), static_cast<void (QTabBar::*)(int)>(&QTabBar::tabCloseRequested), static_cast<MyQTabBar*>(ptr), static_cast<void (MyQTabBar::*)(int)>(&MyQTabBar::Signal_TabCloseRequested));;
}

void* QTabBar_TabData(void* ptr, int index){
	return new QVariant(static_cast<QTabBar*>(ptr)->tabData(index));
}

void QTabBar_ConnectTabMoved(void* ptr){
	QObject::connect(static_cast<QTabBar*>(ptr), static_cast<void (QTabBar::*)(int, int)>(&QTabBar::tabMoved), static_cast<MyQTabBar*>(ptr), static_cast<void (MyQTabBar::*)(int, int)>(&MyQTabBar::Signal_TabMoved));;
}

void QTabBar_DisconnectTabMoved(void* ptr){
	QObject::disconnect(static_cast<QTabBar*>(ptr), static_cast<void (QTabBar::*)(int, int)>(&QTabBar::tabMoved), static_cast<MyQTabBar*>(ptr), static_cast<void (MyQTabBar::*)(int, int)>(&MyQTabBar::Signal_TabMoved));;
}

char* QTabBar_TabText(void* ptr, int index){
	return static_cast<QTabBar*>(ptr)->tabText(index).toUtf8().data();
}

void* QTabBar_TabTextColor(void* ptr, int index){
	return new QColor(static_cast<QTabBar*>(ptr)->tabTextColor(index));
}

char* QTabBar_TabToolTip(void* ptr, int index){
	return static_cast<QTabBar*>(ptr)->tabToolTip(index).toUtf8().data();
}

char* QTabBar_TabWhatsThis(void* ptr, int index){
	return static_cast<QTabBar*>(ptr)->tabWhatsThis(index).toUtf8().data();
}

void QTabBar_DestroyQTabBar(void* ptr){
	static_cast<QTabBar*>(ptr)->~QTabBar();
}

