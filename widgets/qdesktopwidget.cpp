#include "qdesktopwidget.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QObject>
#include <QPoint>
#include <QDesktopWidget>
#include "_cgo_export.h"

class MyQDesktopWidget: public QDesktopWidget {
public:
void Signal_Resized(int screen){callbackQDesktopWidgetResized(this->objectName().toUtf8().data(), screen);};
void Signal_ScreenCountChanged(int newCount){callbackQDesktopWidgetScreenCountChanged(this->objectName().toUtf8().data(), newCount);};
void Signal_WorkAreaResized(int screen){callbackQDesktopWidgetWorkAreaResized(this->objectName().toUtf8().data(), screen);};
};

int QDesktopWidget_IsVirtualDesktop(QtObjectPtr ptr){
	return static_cast<QDesktopWidget*>(ptr)->isVirtualDesktop();
}

int QDesktopWidget_PrimaryScreen(QtObjectPtr ptr){
	return static_cast<QDesktopWidget*>(ptr)->primaryScreen();
}

QtObjectPtr QDesktopWidget_Screen(QtObjectPtr ptr, int screen){
	return static_cast<QDesktopWidget*>(ptr)->screen(screen);
}

int QDesktopWidget_ScreenNumber2(QtObjectPtr ptr, QtObjectPtr point){
	return static_cast<QDesktopWidget*>(ptr)->screenNumber(*static_cast<QPoint*>(point));
}

int QDesktopWidget_ScreenNumber(QtObjectPtr ptr, QtObjectPtr widget){
	return static_cast<QDesktopWidget*>(ptr)->screenNumber(static_cast<QWidget*>(widget));
}

void QDesktopWidget_ConnectResized(QtObjectPtr ptr){
	QObject::connect(static_cast<QDesktopWidget*>(ptr), static_cast<void (QDesktopWidget::*)(int)>(&QDesktopWidget::resized), static_cast<MyQDesktopWidget*>(ptr), static_cast<void (MyQDesktopWidget::*)(int)>(&MyQDesktopWidget::Signal_Resized));;
}

void QDesktopWidget_DisconnectResized(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QDesktopWidget*>(ptr), static_cast<void (QDesktopWidget::*)(int)>(&QDesktopWidget::resized), static_cast<MyQDesktopWidget*>(ptr), static_cast<void (MyQDesktopWidget::*)(int)>(&MyQDesktopWidget::Signal_Resized));;
}

int QDesktopWidget_ScreenCount(QtObjectPtr ptr){
	return static_cast<QDesktopWidget*>(ptr)->screenCount();
}

void QDesktopWidget_ConnectScreenCountChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QDesktopWidget*>(ptr), static_cast<void (QDesktopWidget::*)(int)>(&QDesktopWidget::screenCountChanged), static_cast<MyQDesktopWidget*>(ptr), static_cast<void (MyQDesktopWidget::*)(int)>(&MyQDesktopWidget::Signal_ScreenCountChanged));;
}

void QDesktopWidget_DisconnectScreenCountChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QDesktopWidget*>(ptr), static_cast<void (QDesktopWidget::*)(int)>(&QDesktopWidget::screenCountChanged), static_cast<MyQDesktopWidget*>(ptr), static_cast<void (MyQDesktopWidget::*)(int)>(&MyQDesktopWidget::Signal_ScreenCountChanged));;
}

void QDesktopWidget_ConnectWorkAreaResized(QtObjectPtr ptr){
	QObject::connect(static_cast<QDesktopWidget*>(ptr), static_cast<void (QDesktopWidget::*)(int)>(&QDesktopWidget::workAreaResized), static_cast<MyQDesktopWidget*>(ptr), static_cast<void (MyQDesktopWidget::*)(int)>(&MyQDesktopWidget::Signal_WorkAreaResized));;
}

void QDesktopWidget_DisconnectWorkAreaResized(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QDesktopWidget*>(ptr), static_cast<void (QDesktopWidget::*)(int)>(&QDesktopWidget::workAreaResized), static_cast<MyQDesktopWidget*>(ptr), static_cast<void (MyQDesktopWidget::*)(int)>(&MyQDesktopWidget::Signal_WorkAreaResized));;
}

