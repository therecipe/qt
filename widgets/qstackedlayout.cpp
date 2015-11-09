#include "qstackedlayout.h"
#include <QUrl>
#include <QStack>
#include <QRect>
#include <QLayoutItem>
#include <QString>
#include <QVariant>
#include <QModelIndex>
#include <QLayout>
#include <QMetaObject>
#include <QWidget>
#include <QObject>
#include <QStackedLayout>
#include "_cgo_export.h"

class MyQStackedLayout: public QStackedLayout {
public:
void Signal_CurrentChanged(int index){callbackQStackedLayoutCurrentChanged(this->objectName().toUtf8().data(), index);};
void Signal_WidgetRemoved(int index){callbackQStackedLayoutWidgetRemoved(this->objectName().toUtf8().data(), index);};
};

int QStackedLayout_Count(void* ptr){
	return static_cast<QStackedLayout*>(ptr)->count();
}

int QStackedLayout_CurrentIndex(void* ptr){
	return static_cast<QStackedLayout*>(ptr)->currentIndex();
}

void QStackedLayout_SetCurrentIndex(void* ptr, int index){
	QMetaObject::invokeMethod(static_cast<QStackedLayout*>(ptr), "setCurrentIndex", Q_ARG(int, index));
}

void QStackedLayout_SetCurrentWidget(void* ptr, void* widget){
	QMetaObject::invokeMethod(static_cast<QStackedLayout*>(ptr), "setCurrentWidget", Q_ARG(QWidget*, static_cast<QWidget*>(widget)));
}

void QStackedLayout_SetStackingMode(void* ptr, int stackingMode){
	static_cast<QStackedLayout*>(ptr)->setStackingMode(static_cast<QStackedLayout::StackingMode>(stackingMode));
}

int QStackedLayout_StackingMode(void* ptr){
	return static_cast<QStackedLayout*>(ptr)->stackingMode();
}

void* QStackedLayout_NewQStackedLayout(){
	return new QStackedLayout();
}

void* QStackedLayout_NewQStackedLayout3(void* parentLayout){
	return new QStackedLayout(static_cast<QLayout*>(parentLayout));
}

void* QStackedLayout_NewQStackedLayout2(void* parent){
	return new QStackedLayout(static_cast<QWidget*>(parent));
}

void QStackedLayout_AddItem(void* ptr, void* item){
	static_cast<QStackedLayout*>(ptr)->addItem(static_cast<QLayoutItem*>(item));
}

int QStackedLayout_AddWidget(void* ptr, void* widget){
	return static_cast<QStackedLayout*>(ptr)->addWidget(static_cast<QWidget*>(widget));
}

void QStackedLayout_ConnectCurrentChanged(void* ptr){
	QObject::connect(static_cast<QStackedLayout*>(ptr), static_cast<void (QStackedLayout::*)(int)>(&QStackedLayout::currentChanged), static_cast<MyQStackedLayout*>(ptr), static_cast<void (MyQStackedLayout::*)(int)>(&MyQStackedLayout::Signal_CurrentChanged));;
}

void QStackedLayout_DisconnectCurrentChanged(void* ptr){
	QObject::disconnect(static_cast<QStackedLayout*>(ptr), static_cast<void (QStackedLayout::*)(int)>(&QStackedLayout::currentChanged), static_cast<MyQStackedLayout*>(ptr), static_cast<void (MyQStackedLayout::*)(int)>(&MyQStackedLayout::Signal_CurrentChanged));;
}

void* QStackedLayout_CurrentWidget(void* ptr){
	return static_cast<QStackedLayout*>(ptr)->currentWidget();
}

int QStackedLayout_HasHeightForWidth(void* ptr){
	return static_cast<QStackedLayout*>(ptr)->hasHeightForWidth();
}

int QStackedLayout_HeightForWidth(void* ptr, int width){
	return static_cast<QStackedLayout*>(ptr)->heightForWidth(width);
}

int QStackedLayout_InsertWidget(void* ptr, int index, void* widget){
	return static_cast<QStackedLayout*>(ptr)->insertWidget(index, static_cast<QWidget*>(widget));
}

void* QStackedLayout_ItemAt(void* ptr, int index){
	return static_cast<QStackedLayout*>(ptr)->itemAt(index);
}

void QStackedLayout_SetGeometry(void* ptr, void* rect){
	static_cast<QStackedLayout*>(ptr)->setGeometry(*static_cast<QRect*>(rect));
}

void* QStackedLayout_TakeAt(void* ptr, int index){
	return static_cast<QStackedLayout*>(ptr)->takeAt(index);
}

void* QStackedLayout_Widget(void* ptr, int index){
	return static_cast<QStackedLayout*>(ptr)->widget(index);
}

void QStackedLayout_ConnectWidgetRemoved(void* ptr){
	QObject::connect(static_cast<QStackedLayout*>(ptr), static_cast<void (QStackedLayout::*)(int)>(&QStackedLayout::widgetRemoved), static_cast<MyQStackedLayout*>(ptr), static_cast<void (MyQStackedLayout::*)(int)>(&MyQStackedLayout::Signal_WidgetRemoved));;
}

void QStackedLayout_DisconnectWidgetRemoved(void* ptr){
	QObject::disconnect(static_cast<QStackedLayout*>(ptr), static_cast<void (QStackedLayout::*)(int)>(&QStackedLayout::widgetRemoved), static_cast<MyQStackedLayout*>(ptr), static_cast<void (MyQStackedLayout::*)(int)>(&MyQStackedLayout::Signal_WidgetRemoved));;
}

void QStackedLayout_DestroyQStackedLayout(void* ptr){
	static_cast<QStackedLayout*>(ptr)->~QStackedLayout();
}

