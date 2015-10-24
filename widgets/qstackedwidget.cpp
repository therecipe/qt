#include "qstackedwidget.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QMetaObject>
#include <QObject>
#include <QStack>
#include <QStackedWidget>
#include "_cgo_export.h"

class MyQStackedWidget: public QStackedWidget {
public:
void Signal_CurrentChanged(int index){callbackQStackedWidgetCurrentChanged(this->objectName().toUtf8().data(), index);};
void Signal_WidgetRemoved(int index){callbackQStackedWidgetWidgetRemoved(this->objectName().toUtf8().data(), index);};
};

int QStackedWidget_Count(QtObjectPtr ptr){
	return static_cast<QStackedWidget*>(ptr)->count();
}

int QStackedWidget_CurrentIndex(QtObjectPtr ptr){
	return static_cast<QStackedWidget*>(ptr)->currentIndex();
}

void QStackedWidget_SetCurrentIndex(QtObjectPtr ptr, int index){
	QMetaObject::invokeMethod(static_cast<QStackedWidget*>(ptr), "setCurrentIndex", Q_ARG(int, index));
}

void QStackedWidget_SetCurrentWidget(QtObjectPtr ptr, QtObjectPtr widget){
	QMetaObject::invokeMethod(static_cast<QStackedWidget*>(ptr), "setCurrentWidget", Q_ARG(QWidget*, static_cast<QWidget*>(widget)));
}

QtObjectPtr QStackedWidget_NewQStackedWidget(QtObjectPtr parent){
	return new QStackedWidget(static_cast<QWidget*>(parent));
}

int QStackedWidget_AddWidget(QtObjectPtr ptr, QtObjectPtr widget){
	return static_cast<QStackedWidget*>(ptr)->addWidget(static_cast<QWidget*>(widget));
}

void QStackedWidget_ConnectCurrentChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QStackedWidget*>(ptr), static_cast<void (QStackedWidget::*)(int)>(&QStackedWidget::currentChanged), static_cast<MyQStackedWidget*>(ptr), static_cast<void (MyQStackedWidget::*)(int)>(&MyQStackedWidget::Signal_CurrentChanged));;
}

void QStackedWidget_DisconnectCurrentChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QStackedWidget*>(ptr), static_cast<void (QStackedWidget::*)(int)>(&QStackedWidget::currentChanged), static_cast<MyQStackedWidget*>(ptr), static_cast<void (MyQStackedWidget::*)(int)>(&MyQStackedWidget::Signal_CurrentChanged));;
}

QtObjectPtr QStackedWidget_CurrentWidget(QtObjectPtr ptr){
	return static_cast<QStackedWidget*>(ptr)->currentWidget();
}

int QStackedWidget_IndexOf(QtObjectPtr ptr, QtObjectPtr widget){
	return static_cast<QStackedWidget*>(ptr)->indexOf(static_cast<QWidget*>(widget));
}

int QStackedWidget_InsertWidget(QtObjectPtr ptr, int index, QtObjectPtr widget){
	return static_cast<QStackedWidget*>(ptr)->insertWidget(index, static_cast<QWidget*>(widget));
}

void QStackedWidget_RemoveWidget(QtObjectPtr ptr, QtObjectPtr widget){
	static_cast<QStackedWidget*>(ptr)->removeWidget(static_cast<QWidget*>(widget));
}

QtObjectPtr QStackedWidget_Widget(QtObjectPtr ptr, int index){
	return static_cast<QStackedWidget*>(ptr)->widget(index);
}

void QStackedWidget_ConnectWidgetRemoved(QtObjectPtr ptr){
	QObject::connect(static_cast<QStackedWidget*>(ptr), static_cast<void (QStackedWidget::*)(int)>(&QStackedWidget::widgetRemoved), static_cast<MyQStackedWidget*>(ptr), static_cast<void (MyQStackedWidget::*)(int)>(&MyQStackedWidget::Signal_WidgetRemoved));;
}

void QStackedWidget_DisconnectWidgetRemoved(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QStackedWidget*>(ptr), static_cast<void (QStackedWidget::*)(int)>(&QStackedWidget::widgetRemoved), static_cast<MyQStackedWidget*>(ptr), static_cast<void (MyQStackedWidget::*)(int)>(&MyQStackedWidget::Signal_WidgetRemoved));;
}

void QStackedWidget_DestroyQStackedWidget(QtObjectPtr ptr){
	static_cast<QStackedWidget*>(ptr)->~QStackedWidget();
}

