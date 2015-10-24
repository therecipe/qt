#include "qstackedlayout.h"
#include <QString>
#include <QUrl>
#include <QWidget>
#include <QMetaObject>
#include <QStack>
#include <QRect>
#include <QVariant>
#include <QModelIndex>
#include <QLayout>
#include <QObject>
#include <QLayoutItem>
#include <QStackedLayout>
#include "_cgo_export.h"

class MyQStackedLayout: public QStackedLayout {
public:
void Signal_CurrentChanged(int index){callbackQStackedLayoutCurrentChanged(this->objectName().toUtf8().data(), index);};
void Signal_WidgetRemoved(int index){callbackQStackedLayoutWidgetRemoved(this->objectName().toUtf8().data(), index);};
};

int QStackedLayout_Count(QtObjectPtr ptr){
	return static_cast<QStackedLayout*>(ptr)->count();
}

int QStackedLayout_CurrentIndex(QtObjectPtr ptr){
	return static_cast<QStackedLayout*>(ptr)->currentIndex();
}

void QStackedLayout_SetCurrentIndex(QtObjectPtr ptr, int index){
	QMetaObject::invokeMethod(static_cast<QStackedLayout*>(ptr), "setCurrentIndex", Q_ARG(int, index));
}

void QStackedLayout_SetCurrentWidget(QtObjectPtr ptr, QtObjectPtr widget){
	QMetaObject::invokeMethod(static_cast<QStackedLayout*>(ptr), "setCurrentWidget", Q_ARG(QWidget*, static_cast<QWidget*>(widget)));
}

void QStackedLayout_SetStackingMode(QtObjectPtr ptr, int stackingMode){
	static_cast<QStackedLayout*>(ptr)->setStackingMode(static_cast<QStackedLayout::StackingMode>(stackingMode));
}

int QStackedLayout_StackingMode(QtObjectPtr ptr){
	return static_cast<QStackedLayout*>(ptr)->stackingMode();
}

QtObjectPtr QStackedLayout_NewQStackedLayout(){
	return new QStackedLayout();
}

QtObjectPtr QStackedLayout_NewQStackedLayout3(QtObjectPtr parentLayout){
	return new QStackedLayout(static_cast<QLayout*>(parentLayout));
}

QtObjectPtr QStackedLayout_NewQStackedLayout2(QtObjectPtr parent){
	return new QStackedLayout(static_cast<QWidget*>(parent));
}

void QStackedLayout_AddItem(QtObjectPtr ptr, QtObjectPtr item){
	static_cast<QStackedLayout*>(ptr)->addItem(static_cast<QLayoutItem*>(item));
}

int QStackedLayout_AddWidget(QtObjectPtr ptr, QtObjectPtr widget){
	return static_cast<QStackedLayout*>(ptr)->addWidget(static_cast<QWidget*>(widget));
}

void QStackedLayout_ConnectCurrentChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QStackedLayout*>(ptr), static_cast<void (QStackedLayout::*)(int)>(&QStackedLayout::currentChanged), static_cast<MyQStackedLayout*>(ptr), static_cast<void (MyQStackedLayout::*)(int)>(&MyQStackedLayout::Signal_CurrentChanged));;
}

void QStackedLayout_DisconnectCurrentChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QStackedLayout*>(ptr), static_cast<void (QStackedLayout::*)(int)>(&QStackedLayout::currentChanged), static_cast<MyQStackedLayout*>(ptr), static_cast<void (MyQStackedLayout::*)(int)>(&MyQStackedLayout::Signal_CurrentChanged));;
}

QtObjectPtr QStackedLayout_CurrentWidget(QtObjectPtr ptr){
	return static_cast<QStackedLayout*>(ptr)->currentWidget();
}

int QStackedLayout_HasHeightForWidth(QtObjectPtr ptr){
	return static_cast<QStackedLayout*>(ptr)->hasHeightForWidth();
}

int QStackedLayout_HeightForWidth(QtObjectPtr ptr, int width){
	return static_cast<QStackedLayout*>(ptr)->heightForWidth(width);
}

int QStackedLayout_InsertWidget(QtObjectPtr ptr, int index, QtObjectPtr widget){
	return static_cast<QStackedLayout*>(ptr)->insertWidget(index, static_cast<QWidget*>(widget));
}

QtObjectPtr QStackedLayout_ItemAt(QtObjectPtr ptr, int index){
	return static_cast<QStackedLayout*>(ptr)->itemAt(index);
}

void QStackedLayout_SetGeometry(QtObjectPtr ptr, QtObjectPtr rect){
	static_cast<QStackedLayout*>(ptr)->setGeometry(*static_cast<QRect*>(rect));
}

QtObjectPtr QStackedLayout_TakeAt(QtObjectPtr ptr, int index){
	return static_cast<QStackedLayout*>(ptr)->takeAt(index);
}

QtObjectPtr QStackedLayout_Widget(QtObjectPtr ptr, int index){
	return static_cast<QStackedLayout*>(ptr)->widget(index);
}

void QStackedLayout_ConnectWidgetRemoved(QtObjectPtr ptr){
	QObject::connect(static_cast<QStackedLayout*>(ptr), static_cast<void (QStackedLayout::*)(int)>(&QStackedLayout::widgetRemoved), static_cast<MyQStackedLayout*>(ptr), static_cast<void (MyQStackedLayout::*)(int)>(&MyQStackedLayout::Signal_WidgetRemoved));;
}

void QStackedLayout_DisconnectWidgetRemoved(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QStackedLayout*>(ptr), static_cast<void (QStackedLayout::*)(int)>(&QStackedLayout::widgetRemoved), static_cast<MyQStackedLayout*>(ptr), static_cast<void (MyQStackedLayout::*)(int)>(&MyQStackedLayout::Signal_WidgetRemoved));;
}

void QStackedLayout_DestroyQStackedLayout(QtObjectPtr ptr){
	static_cast<QStackedLayout*>(ptr)->~QStackedLayout();
}

