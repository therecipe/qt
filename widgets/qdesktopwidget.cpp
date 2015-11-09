#include "qdesktopwidget.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QPoint>
#include <QObject>
#include <QDesktopWidget>
#include "_cgo_export.h"

class MyQDesktopWidget: public QDesktopWidget {
public:
void Signal_Resized(int screen){callbackQDesktopWidgetResized(this->objectName().toUtf8().data(), screen);};
void Signal_ScreenCountChanged(int newCount){callbackQDesktopWidgetScreenCountChanged(this->objectName().toUtf8().data(), newCount);};
void Signal_WorkAreaResized(int screen){callbackQDesktopWidgetWorkAreaResized(this->objectName().toUtf8().data(), screen);};
};

int QDesktopWidget_IsVirtualDesktop(void* ptr){
	return static_cast<QDesktopWidget*>(ptr)->isVirtualDesktop();
}

int QDesktopWidget_PrimaryScreen(void* ptr){
	return static_cast<QDesktopWidget*>(ptr)->primaryScreen();
}

void* QDesktopWidget_Screen(void* ptr, int screen){
	return static_cast<QDesktopWidget*>(ptr)->screen(screen);
}

int QDesktopWidget_ScreenNumber2(void* ptr, void* point){
	return static_cast<QDesktopWidget*>(ptr)->screenNumber(*static_cast<QPoint*>(point));
}

int QDesktopWidget_ScreenNumber(void* ptr, void* widget){
	return static_cast<QDesktopWidget*>(ptr)->screenNumber(static_cast<QWidget*>(widget));
}

void QDesktopWidget_ConnectResized(void* ptr){
	QObject::connect(static_cast<QDesktopWidget*>(ptr), static_cast<void (QDesktopWidget::*)(int)>(&QDesktopWidget::resized), static_cast<MyQDesktopWidget*>(ptr), static_cast<void (MyQDesktopWidget::*)(int)>(&MyQDesktopWidget::Signal_Resized));;
}

void QDesktopWidget_DisconnectResized(void* ptr){
	QObject::disconnect(static_cast<QDesktopWidget*>(ptr), static_cast<void (QDesktopWidget::*)(int)>(&QDesktopWidget::resized), static_cast<MyQDesktopWidget*>(ptr), static_cast<void (MyQDesktopWidget::*)(int)>(&MyQDesktopWidget::Signal_Resized));;
}

int QDesktopWidget_ScreenCount(void* ptr){
	return static_cast<QDesktopWidget*>(ptr)->screenCount();
}

void QDesktopWidget_ConnectScreenCountChanged(void* ptr){
	QObject::connect(static_cast<QDesktopWidget*>(ptr), static_cast<void (QDesktopWidget::*)(int)>(&QDesktopWidget::screenCountChanged), static_cast<MyQDesktopWidget*>(ptr), static_cast<void (MyQDesktopWidget::*)(int)>(&MyQDesktopWidget::Signal_ScreenCountChanged));;
}

void QDesktopWidget_DisconnectScreenCountChanged(void* ptr){
	QObject::disconnect(static_cast<QDesktopWidget*>(ptr), static_cast<void (QDesktopWidget::*)(int)>(&QDesktopWidget::screenCountChanged), static_cast<MyQDesktopWidget*>(ptr), static_cast<void (MyQDesktopWidget::*)(int)>(&MyQDesktopWidget::Signal_ScreenCountChanged));;
}

void QDesktopWidget_ConnectWorkAreaResized(void* ptr){
	QObject::connect(static_cast<QDesktopWidget*>(ptr), static_cast<void (QDesktopWidget::*)(int)>(&QDesktopWidget::workAreaResized), static_cast<MyQDesktopWidget*>(ptr), static_cast<void (MyQDesktopWidget::*)(int)>(&MyQDesktopWidget::Signal_WorkAreaResized));;
}

void QDesktopWidget_DisconnectWorkAreaResized(void* ptr){
	QObject::disconnect(static_cast<QDesktopWidget*>(ptr), static_cast<void (QDesktopWidget::*)(int)>(&QDesktopWidget::workAreaResized), static_cast<MyQDesktopWidget*>(ptr), static_cast<void (MyQDesktopWidget::*)(int)>(&MyQDesktopWidget::Signal_WorkAreaResized));;
}

