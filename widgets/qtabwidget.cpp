#include "qtabwidget.h"
#include <QObject>
#include <QWidget>
#include <QVariant>
#include <QUrl>
#include <QSize>
#include <QIcon>
#include <QString>
#include <QModelIndex>
#include <QMetaObject>
#include <QTabWidget>
#include "_cgo_export.h"

class MyQTabWidget: public QTabWidget {
public:
void Signal_CurrentChanged(int index){callbackQTabWidgetCurrentChanged(this->objectName().toUtf8().data(), index);};
void Signal_TabBarClicked(int index){callbackQTabWidgetTabBarClicked(this->objectName().toUtf8().data(), index);};
void Signal_TabBarDoubleClicked(int index){callbackQTabWidgetTabBarDoubleClicked(this->objectName().toUtf8().data(), index);};
void Signal_TabCloseRequested(int index){callbackQTabWidgetTabCloseRequested(this->objectName().toUtf8().data(), index);};
};

int QTabWidget_AddTab2(QtObjectPtr ptr, QtObjectPtr page, QtObjectPtr icon, char* label){
	return static_cast<QTabWidget*>(ptr)->addTab(static_cast<QWidget*>(page), *static_cast<QIcon*>(icon), QString(label));
}

int QTabWidget_AddTab(QtObjectPtr ptr, QtObjectPtr page, char* label){
	return static_cast<QTabWidget*>(ptr)->addTab(static_cast<QWidget*>(page), QString(label));
}

int QTabWidget_Count(QtObjectPtr ptr){
	return static_cast<QTabWidget*>(ptr)->count();
}

int QTabWidget_CurrentIndex(QtObjectPtr ptr){
	return static_cast<QTabWidget*>(ptr)->currentIndex();
}

int QTabWidget_DocumentMode(QtObjectPtr ptr){
	return static_cast<QTabWidget*>(ptr)->documentMode();
}

int QTabWidget_ElideMode(QtObjectPtr ptr){
	return static_cast<QTabWidget*>(ptr)->elideMode();
}

int QTabWidget_InsertTab2(QtObjectPtr ptr, int index, QtObjectPtr page, QtObjectPtr icon, char* label){
	return static_cast<QTabWidget*>(ptr)->insertTab(index, static_cast<QWidget*>(page), *static_cast<QIcon*>(icon), QString(label));
}

int QTabWidget_InsertTab(QtObjectPtr ptr, int index, QtObjectPtr page, char* label){
	return static_cast<QTabWidget*>(ptr)->insertTab(index, static_cast<QWidget*>(page), QString(label));
}

int QTabWidget_IsMovable(QtObjectPtr ptr){
	return static_cast<QTabWidget*>(ptr)->isMovable();
}

void QTabWidget_SetCornerWidget(QtObjectPtr ptr, QtObjectPtr widget, int corner){
	static_cast<QTabWidget*>(ptr)->setCornerWidget(static_cast<QWidget*>(widget), static_cast<Qt::Corner>(corner));
}

void QTabWidget_SetCurrentIndex(QtObjectPtr ptr, int index){
	QMetaObject::invokeMethod(static_cast<QTabWidget*>(ptr), "setCurrentIndex", Q_ARG(int, index));
}

void QTabWidget_SetDocumentMode(QtObjectPtr ptr, int set){
	static_cast<QTabWidget*>(ptr)->setDocumentMode(set != 0);
}

void QTabWidget_SetElideMode(QtObjectPtr ptr, int v){
	static_cast<QTabWidget*>(ptr)->setElideMode(static_cast<Qt::TextElideMode>(v));
}

void QTabWidget_SetIconSize(QtObjectPtr ptr, QtObjectPtr size){
	static_cast<QTabWidget*>(ptr)->setIconSize(*static_cast<QSize*>(size));
}

void QTabWidget_SetMovable(QtObjectPtr ptr, int movable){
	static_cast<QTabWidget*>(ptr)->setMovable(movable != 0);
}

void QTabWidget_SetTabBarAutoHide(QtObjectPtr ptr, int enabled){
	static_cast<QTabWidget*>(ptr)->setTabBarAutoHide(enabled != 0);
}

void QTabWidget_SetTabPosition(QtObjectPtr ptr, int v){
	static_cast<QTabWidget*>(ptr)->setTabPosition(static_cast<QTabWidget::TabPosition>(v));
}

void QTabWidget_SetTabShape(QtObjectPtr ptr, int s){
	static_cast<QTabWidget*>(ptr)->setTabShape(static_cast<QTabWidget::TabShape>(s));
}

void QTabWidget_SetTabsClosable(QtObjectPtr ptr, int closeable){
	static_cast<QTabWidget*>(ptr)->setTabsClosable(closeable != 0);
}

void QTabWidget_SetUsesScrollButtons(QtObjectPtr ptr, int useButtons){
	static_cast<QTabWidget*>(ptr)->setUsesScrollButtons(useButtons != 0);
}

int QTabWidget_TabBarAutoHide(QtObjectPtr ptr){
	return static_cast<QTabWidget*>(ptr)->tabBarAutoHide();
}

int QTabWidget_TabPosition(QtObjectPtr ptr){
	return static_cast<QTabWidget*>(ptr)->tabPosition();
}

int QTabWidget_TabShape(QtObjectPtr ptr){
	return static_cast<QTabWidget*>(ptr)->tabShape();
}

int QTabWidget_TabsClosable(QtObjectPtr ptr){
	return static_cast<QTabWidget*>(ptr)->tabsClosable();
}

int QTabWidget_UsesScrollButtons(QtObjectPtr ptr){
	return static_cast<QTabWidget*>(ptr)->usesScrollButtons();
}

QtObjectPtr QTabWidget_NewQTabWidget(QtObjectPtr parent){
	return new QTabWidget(static_cast<QWidget*>(parent));
}

void QTabWidget_Clear(QtObjectPtr ptr){
	static_cast<QTabWidget*>(ptr)->clear();
}

QtObjectPtr QTabWidget_CornerWidget(QtObjectPtr ptr, int corner){
	return static_cast<QTabWidget*>(ptr)->cornerWidget(static_cast<Qt::Corner>(corner));
}

void QTabWidget_ConnectCurrentChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QTabWidget*>(ptr), static_cast<void (QTabWidget::*)(int)>(&QTabWidget::currentChanged), static_cast<MyQTabWidget*>(ptr), static_cast<void (MyQTabWidget::*)(int)>(&MyQTabWidget::Signal_CurrentChanged));;
}

void QTabWidget_DisconnectCurrentChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTabWidget*>(ptr), static_cast<void (QTabWidget::*)(int)>(&QTabWidget::currentChanged), static_cast<MyQTabWidget*>(ptr), static_cast<void (MyQTabWidget::*)(int)>(&MyQTabWidget::Signal_CurrentChanged));;
}

QtObjectPtr QTabWidget_CurrentWidget(QtObjectPtr ptr){
	return static_cast<QTabWidget*>(ptr)->currentWidget();
}

int QTabWidget_HasHeightForWidth(QtObjectPtr ptr){
	return static_cast<QTabWidget*>(ptr)->hasHeightForWidth();
}

int QTabWidget_HeightForWidth(QtObjectPtr ptr, int width){
	return static_cast<QTabWidget*>(ptr)->heightForWidth(width);
}

int QTabWidget_IndexOf(QtObjectPtr ptr, QtObjectPtr w){
	return static_cast<QTabWidget*>(ptr)->indexOf(static_cast<QWidget*>(w));
}

int QTabWidget_IsTabEnabled(QtObjectPtr ptr, int index){
	return static_cast<QTabWidget*>(ptr)->isTabEnabled(index);
}

void QTabWidget_RemoveTab(QtObjectPtr ptr, int index){
	static_cast<QTabWidget*>(ptr)->removeTab(index);
}

void QTabWidget_SetCurrentWidget(QtObjectPtr ptr, QtObjectPtr widget){
	QMetaObject::invokeMethod(static_cast<QTabWidget*>(ptr), "setCurrentWidget", Q_ARG(QWidget*, static_cast<QWidget*>(widget)));
}

void QTabWidget_SetTabEnabled(QtObjectPtr ptr, int index, int enable){
	static_cast<QTabWidget*>(ptr)->setTabEnabled(index, enable != 0);
}

void QTabWidget_SetTabIcon(QtObjectPtr ptr, int index, QtObjectPtr icon){
	static_cast<QTabWidget*>(ptr)->setTabIcon(index, *static_cast<QIcon*>(icon));
}

void QTabWidget_SetTabText(QtObjectPtr ptr, int index, char* label){
	static_cast<QTabWidget*>(ptr)->setTabText(index, QString(label));
}

void QTabWidget_SetTabToolTip(QtObjectPtr ptr, int index, char* tip){
	static_cast<QTabWidget*>(ptr)->setTabToolTip(index, QString(tip));
}

void QTabWidget_SetTabWhatsThis(QtObjectPtr ptr, int index, char* text){
	static_cast<QTabWidget*>(ptr)->setTabWhatsThis(index, QString(text));
}

QtObjectPtr QTabWidget_TabBar(QtObjectPtr ptr){
	return static_cast<QTabWidget*>(ptr)->tabBar();
}

void QTabWidget_ConnectTabBarClicked(QtObjectPtr ptr){
	QObject::connect(static_cast<QTabWidget*>(ptr), static_cast<void (QTabWidget::*)(int)>(&QTabWidget::tabBarClicked), static_cast<MyQTabWidget*>(ptr), static_cast<void (MyQTabWidget::*)(int)>(&MyQTabWidget::Signal_TabBarClicked));;
}

void QTabWidget_DisconnectTabBarClicked(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTabWidget*>(ptr), static_cast<void (QTabWidget::*)(int)>(&QTabWidget::tabBarClicked), static_cast<MyQTabWidget*>(ptr), static_cast<void (MyQTabWidget::*)(int)>(&MyQTabWidget::Signal_TabBarClicked));;
}

void QTabWidget_ConnectTabBarDoubleClicked(QtObjectPtr ptr){
	QObject::connect(static_cast<QTabWidget*>(ptr), static_cast<void (QTabWidget::*)(int)>(&QTabWidget::tabBarDoubleClicked), static_cast<MyQTabWidget*>(ptr), static_cast<void (MyQTabWidget::*)(int)>(&MyQTabWidget::Signal_TabBarDoubleClicked));;
}

void QTabWidget_DisconnectTabBarDoubleClicked(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTabWidget*>(ptr), static_cast<void (QTabWidget::*)(int)>(&QTabWidget::tabBarDoubleClicked), static_cast<MyQTabWidget*>(ptr), static_cast<void (MyQTabWidget::*)(int)>(&MyQTabWidget::Signal_TabBarDoubleClicked));;
}

void QTabWidget_ConnectTabCloseRequested(QtObjectPtr ptr){
	QObject::connect(static_cast<QTabWidget*>(ptr), static_cast<void (QTabWidget::*)(int)>(&QTabWidget::tabCloseRequested), static_cast<MyQTabWidget*>(ptr), static_cast<void (MyQTabWidget::*)(int)>(&MyQTabWidget::Signal_TabCloseRequested));;
}

void QTabWidget_DisconnectTabCloseRequested(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTabWidget*>(ptr), static_cast<void (QTabWidget::*)(int)>(&QTabWidget::tabCloseRequested), static_cast<MyQTabWidget*>(ptr), static_cast<void (MyQTabWidget::*)(int)>(&MyQTabWidget::Signal_TabCloseRequested));;
}

char* QTabWidget_TabText(QtObjectPtr ptr, int index){
	return static_cast<QTabWidget*>(ptr)->tabText(index).toUtf8().data();
}

char* QTabWidget_TabToolTip(QtObjectPtr ptr, int index){
	return static_cast<QTabWidget*>(ptr)->tabToolTip(index).toUtf8().data();
}

char* QTabWidget_TabWhatsThis(QtObjectPtr ptr, int index){
	return static_cast<QTabWidget*>(ptr)->tabWhatsThis(index).toUtf8().data();
}

QtObjectPtr QTabWidget_Widget(QtObjectPtr ptr, int index){
	return static_cast<QTabWidget*>(ptr)->widget(index);
}

void QTabWidget_DestroyQTabWidget(QtObjectPtr ptr){
	static_cast<QTabWidget*>(ptr)->~QTabWidget();
}

