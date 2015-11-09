#include "qstackedwidget.h"
#include <QMetaObject>
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QStack>
#include <QStackedWidget>
#include "_cgo_export.h"

class MyQStackedWidget: public QStackedWidget {
public:
void Signal_CurrentChanged(int index){callbackQStackedWidgetCurrentChanged(this->objectName().toUtf8().data(), index);};
void Signal_WidgetRemoved(int index){callbackQStackedWidgetWidgetRemoved(this->objectName().toUtf8().data(), index);};
};

int QStackedWidget_Count(void* ptr){
	return static_cast<QStackedWidget*>(ptr)->count();
}

int QStackedWidget_CurrentIndex(void* ptr){
	return static_cast<QStackedWidget*>(ptr)->currentIndex();
}

void QStackedWidget_SetCurrentIndex(void* ptr, int index){
	QMetaObject::invokeMethod(static_cast<QStackedWidget*>(ptr), "setCurrentIndex", Q_ARG(int, index));
}

void QStackedWidget_SetCurrentWidget(void* ptr, void* widget){
	QMetaObject::invokeMethod(static_cast<QStackedWidget*>(ptr), "setCurrentWidget", Q_ARG(QWidget*, static_cast<QWidget*>(widget)));
}

void* QStackedWidget_NewQStackedWidget(void* parent){
	return new QStackedWidget(static_cast<QWidget*>(parent));
}

int QStackedWidget_AddWidget(void* ptr, void* widget){
	return static_cast<QStackedWidget*>(ptr)->addWidget(static_cast<QWidget*>(widget));
}

void QStackedWidget_ConnectCurrentChanged(void* ptr){
	QObject::connect(static_cast<QStackedWidget*>(ptr), static_cast<void (QStackedWidget::*)(int)>(&QStackedWidget::currentChanged), static_cast<MyQStackedWidget*>(ptr), static_cast<void (MyQStackedWidget::*)(int)>(&MyQStackedWidget::Signal_CurrentChanged));;
}

void QStackedWidget_DisconnectCurrentChanged(void* ptr){
	QObject::disconnect(static_cast<QStackedWidget*>(ptr), static_cast<void (QStackedWidget::*)(int)>(&QStackedWidget::currentChanged), static_cast<MyQStackedWidget*>(ptr), static_cast<void (MyQStackedWidget::*)(int)>(&MyQStackedWidget::Signal_CurrentChanged));;
}

void* QStackedWidget_CurrentWidget(void* ptr){
	return static_cast<QStackedWidget*>(ptr)->currentWidget();
}

int QStackedWidget_IndexOf(void* ptr, void* widget){
	return static_cast<QStackedWidget*>(ptr)->indexOf(static_cast<QWidget*>(widget));
}

int QStackedWidget_InsertWidget(void* ptr, int index, void* widget){
	return static_cast<QStackedWidget*>(ptr)->insertWidget(index, static_cast<QWidget*>(widget));
}

void QStackedWidget_RemoveWidget(void* ptr, void* widget){
	static_cast<QStackedWidget*>(ptr)->removeWidget(static_cast<QWidget*>(widget));
}

void* QStackedWidget_Widget(void* ptr, int index){
	return static_cast<QStackedWidget*>(ptr)->widget(index);
}

void QStackedWidget_ConnectWidgetRemoved(void* ptr){
	QObject::connect(static_cast<QStackedWidget*>(ptr), static_cast<void (QStackedWidget::*)(int)>(&QStackedWidget::widgetRemoved), static_cast<MyQStackedWidget*>(ptr), static_cast<void (MyQStackedWidget::*)(int)>(&MyQStackedWidget::Signal_WidgetRemoved));;
}

void QStackedWidget_DisconnectWidgetRemoved(void* ptr){
	QObject::disconnect(static_cast<QStackedWidget*>(ptr), static_cast<void (QStackedWidget::*)(int)>(&QStackedWidget::widgetRemoved), static_cast<MyQStackedWidget*>(ptr), static_cast<void (MyQStackedWidget::*)(int)>(&MyQStackedWidget::Signal_WidgetRemoved));;
}

void QStackedWidget_DestroyQStackedWidget(void* ptr){
	static_cast<QStackedWidget*>(ptr)->~QStackedWidget();
}

