#include "qtabwidget.h"
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QIcon>
#include <QSize>
#include <QMetaObject>
#include <QObject>
#include <QUrl>
#include <QModelIndex>
#include <QTabWidget>
#include "_cgo_export.h"

class MyQTabWidget: public QTabWidget {
public:
void Signal_CurrentChanged(int index){callbackQTabWidgetCurrentChanged(this->objectName().toUtf8().data(), index);};
void Signal_TabBarClicked(int index){callbackQTabWidgetTabBarClicked(this->objectName().toUtf8().data(), index);};
void Signal_TabBarDoubleClicked(int index){callbackQTabWidgetTabBarDoubleClicked(this->objectName().toUtf8().data(), index);};
void Signal_TabCloseRequested(int index){callbackQTabWidgetTabCloseRequested(this->objectName().toUtf8().data(), index);};
};

int QTabWidget_AddTab2(void* ptr, void* page, void* icon, char* label){
	return static_cast<QTabWidget*>(ptr)->addTab(static_cast<QWidget*>(page), *static_cast<QIcon*>(icon), QString(label));
}

int QTabWidget_AddTab(void* ptr, void* page, char* label){
	return static_cast<QTabWidget*>(ptr)->addTab(static_cast<QWidget*>(page), QString(label));
}

int QTabWidget_Count(void* ptr){
	return static_cast<QTabWidget*>(ptr)->count();
}

int QTabWidget_CurrentIndex(void* ptr){
	return static_cast<QTabWidget*>(ptr)->currentIndex();
}

int QTabWidget_DocumentMode(void* ptr){
	return static_cast<QTabWidget*>(ptr)->documentMode();
}

int QTabWidget_ElideMode(void* ptr){
	return static_cast<QTabWidget*>(ptr)->elideMode();
}

int QTabWidget_InsertTab2(void* ptr, int index, void* page, void* icon, char* label){
	return static_cast<QTabWidget*>(ptr)->insertTab(index, static_cast<QWidget*>(page), *static_cast<QIcon*>(icon), QString(label));
}

int QTabWidget_InsertTab(void* ptr, int index, void* page, char* label){
	return static_cast<QTabWidget*>(ptr)->insertTab(index, static_cast<QWidget*>(page), QString(label));
}

int QTabWidget_IsMovable(void* ptr){
	return static_cast<QTabWidget*>(ptr)->isMovable();
}

void QTabWidget_SetCornerWidget(void* ptr, void* widget, int corner){
	static_cast<QTabWidget*>(ptr)->setCornerWidget(static_cast<QWidget*>(widget), static_cast<Qt::Corner>(corner));
}

void QTabWidget_SetCurrentIndex(void* ptr, int index){
	QMetaObject::invokeMethod(static_cast<QTabWidget*>(ptr), "setCurrentIndex", Q_ARG(int, index));
}

void QTabWidget_SetDocumentMode(void* ptr, int set){
	static_cast<QTabWidget*>(ptr)->setDocumentMode(set != 0);
}

void QTabWidget_SetElideMode(void* ptr, int v){
	static_cast<QTabWidget*>(ptr)->setElideMode(static_cast<Qt::TextElideMode>(v));
}

void QTabWidget_SetIconSize(void* ptr, void* size){
	static_cast<QTabWidget*>(ptr)->setIconSize(*static_cast<QSize*>(size));
}

void QTabWidget_SetMovable(void* ptr, int movable){
	static_cast<QTabWidget*>(ptr)->setMovable(movable != 0);
}

void QTabWidget_SetTabBarAutoHide(void* ptr, int enabled){
	static_cast<QTabWidget*>(ptr)->setTabBarAutoHide(enabled != 0);
}

void QTabWidget_SetTabPosition(void* ptr, int v){
	static_cast<QTabWidget*>(ptr)->setTabPosition(static_cast<QTabWidget::TabPosition>(v));
}

void QTabWidget_SetTabShape(void* ptr, int s){
	static_cast<QTabWidget*>(ptr)->setTabShape(static_cast<QTabWidget::TabShape>(s));
}

void QTabWidget_SetTabsClosable(void* ptr, int closeable){
	static_cast<QTabWidget*>(ptr)->setTabsClosable(closeable != 0);
}

void QTabWidget_SetUsesScrollButtons(void* ptr, int useButtons){
	static_cast<QTabWidget*>(ptr)->setUsesScrollButtons(useButtons != 0);
}

int QTabWidget_TabBarAutoHide(void* ptr){
	return static_cast<QTabWidget*>(ptr)->tabBarAutoHide();
}

int QTabWidget_TabPosition(void* ptr){
	return static_cast<QTabWidget*>(ptr)->tabPosition();
}

int QTabWidget_TabShape(void* ptr){
	return static_cast<QTabWidget*>(ptr)->tabShape();
}

int QTabWidget_TabsClosable(void* ptr){
	return static_cast<QTabWidget*>(ptr)->tabsClosable();
}

int QTabWidget_UsesScrollButtons(void* ptr){
	return static_cast<QTabWidget*>(ptr)->usesScrollButtons();
}

void* QTabWidget_NewQTabWidget(void* parent){
	return new QTabWidget(static_cast<QWidget*>(parent));
}

void QTabWidget_Clear(void* ptr){
	static_cast<QTabWidget*>(ptr)->clear();
}

void* QTabWidget_CornerWidget(void* ptr, int corner){
	return static_cast<QTabWidget*>(ptr)->cornerWidget(static_cast<Qt::Corner>(corner));
}

void QTabWidget_ConnectCurrentChanged(void* ptr){
	QObject::connect(static_cast<QTabWidget*>(ptr), static_cast<void (QTabWidget::*)(int)>(&QTabWidget::currentChanged), static_cast<MyQTabWidget*>(ptr), static_cast<void (MyQTabWidget::*)(int)>(&MyQTabWidget::Signal_CurrentChanged));;
}

void QTabWidget_DisconnectCurrentChanged(void* ptr){
	QObject::disconnect(static_cast<QTabWidget*>(ptr), static_cast<void (QTabWidget::*)(int)>(&QTabWidget::currentChanged), static_cast<MyQTabWidget*>(ptr), static_cast<void (MyQTabWidget::*)(int)>(&MyQTabWidget::Signal_CurrentChanged));;
}

void* QTabWidget_CurrentWidget(void* ptr){
	return static_cast<QTabWidget*>(ptr)->currentWidget();
}

int QTabWidget_HasHeightForWidth(void* ptr){
	return static_cast<QTabWidget*>(ptr)->hasHeightForWidth();
}

int QTabWidget_HeightForWidth(void* ptr, int width){
	return static_cast<QTabWidget*>(ptr)->heightForWidth(width);
}

int QTabWidget_IndexOf(void* ptr, void* w){
	return static_cast<QTabWidget*>(ptr)->indexOf(static_cast<QWidget*>(w));
}

int QTabWidget_IsTabEnabled(void* ptr, int index){
	return static_cast<QTabWidget*>(ptr)->isTabEnabled(index);
}

void QTabWidget_RemoveTab(void* ptr, int index){
	static_cast<QTabWidget*>(ptr)->removeTab(index);
}

void QTabWidget_SetCurrentWidget(void* ptr, void* widget){
	QMetaObject::invokeMethod(static_cast<QTabWidget*>(ptr), "setCurrentWidget", Q_ARG(QWidget*, static_cast<QWidget*>(widget)));
}

void QTabWidget_SetTabEnabled(void* ptr, int index, int enable){
	static_cast<QTabWidget*>(ptr)->setTabEnabled(index, enable != 0);
}

void QTabWidget_SetTabIcon(void* ptr, int index, void* icon){
	static_cast<QTabWidget*>(ptr)->setTabIcon(index, *static_cast<QIcon*>(icon));
}

void QTabWidget_SetTabText(void* ptr, int index, char* label){
	static_cast<QTabWidget*>(ptr)->setTabText(index, QString(label));
}

void QTabWidget_SetTabToolTip(void* ptr, int index, char* tip){
	static_cast<QTabWidget*>(ptr)->setTabToolTip(index, QString(tip));
}

void QTabWidget_SetTabWhatsThis(void* ptr, int index, char* text){
	static_cast<QTabWidget*>(ptr)->setTabWhatsThis(index, QString(text));
}

void* QTabWidget_TabBar(void* ptr){
	return static_cast<QTabWidget*>(ptr)->tabBar();
}

void QTabWidget_ConnectTabBarClicked(void* ptr){
	QObject::connect(static_cast<QTabWidget*>(ptr), static_cast<void (QTabWidget::*)(int)>(&QTabWidget::tabBarClicked), static_cast<MyQTabWidget*>(ptr), static_cast<void (MyQTabWidget::*)(int)>(&MyQTabWidget::Signal_TabBarClicked));;
}

void QTabWidget_DisconnectTabBarClicked(void* ptr){
	QObject::disconnect(static_cast<QTabWidget*>(ptr), static_cast<void (QTabWidget::*)(int)>(&QTabWidget::tabBarClicked), static_cast<MyQTabWidget*>(ptr), static_cast<void (MyQTabWidget::*)(int)>(&MyQTabWidget::Signal_TabBarClicked));;
}

void QTabWidget_ConnectTabBarDoubleClicked(void* ptr){
	QObject::connect(static_cast<QTabWidget*>(ptr), static_cast<void (QTabWidget::*)(int)>(&QTabWidget::tabBarDoubleClicked), static_cast<MyQTabWidget*>(ptr), static_cast<void (MyQTabWidget::*)(int)>(&MyQTabWidget::Signal_TabBarDoubleClicked));;
}

void QTabWidget_DisconnectTabBarDoubleClicked(void* ptr){
	QObject::disconnect(static_cast<QTabWidget*>(ptr), static_cast<void (QTabWidget::*)(int)>(&QTabWidget::tabBarDoubleClicked), static_cast<MyQTabWidget*>(ptr), static_cast<void (MyQTabWidget::*)(int)>(&MyQTabWidget::Signal_TabBarDoubleClicked));;
}

void QTabWidget_ConnectTabCloseRequested(void* ptr){
	QObject::connect(static_cast<QTabWidget*>(ptr), static_cast<void (QTabWidget::*)(int)>(&QTabWidget::tabCloseRequested), static_cast<MyQTabWidget*>(ptr), static_cast<void (MyQTabWidget::*)(int)>(&MyQTabWidget::Signal_TabCloseRequested));;
}

void QTabWidget_DisconnectTabCloseRequested(void* ptr){
	QObject::disconnect(static_cast<QTabWidget*>(ptr), static_cast<void (QTabWidget::*)(int)>(&QTabWidget::tabCloseRequested), static_cast<MyQTabWidget*>(ptr), static_cast<void (MyQTabWidget::*)(int)>(&MyQTabWidget::Signal_TabCloseRequested));;
}

char* QTabWidget_TabText(void* ptr, int index){
	return static_cast<QTabWidget*>(ptr)->tabText(index).toUtf8().data();
}

char* QTabWidget_TabToolTip(void* ptr, int index){
	return static_cast<QTabWidget*>(ptr)->tabToolTip(index).toUtf8().data();
}

char* QTabWidget_TabWhatsThis(void* ptr, int index){
	return static_cast<QTabWidget*>(ptr)->tabWhatsThis(index).toUtf8().data();
}

void* QTabWidget_Widget(void* ptr, int index){
	return static_cast<QTabWidget*>(ptr)->widget(index);
}

void QTabWidget_DestroyQTabWidget(void* ptr){
	static_cast<QTabWidget*>(ptr)->~QTabWidget();
}

