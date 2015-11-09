#include "qmdisubwindow.h"
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QObject>
#include <QMenu>
#include <QMetaObject>
#include <QString>
#include <QVariant>
#include <QMdiSubWindow>
#include "_cgo_export.h"

class MyQMdiSubWindow: public QMdiSubWindow {
public:
void Signal_AboutToActivate(){callbackQMdiSubWindowAboutToActivate(this->objectName().toUtf8().data());};
void Signal_WindowStateChanged(Qt::WindowStates oldState, Qt::WindowStates newState){callbackQMdiSubWindowWindowStateChanged(this->objectName().toUtf8().data(), oldState, newState);};
};

int QMdiSubWindow_KeyboardPageStep(void* ptr){
	return static_cast<QMdiSubWindow*>(ptr)->keyboardPageStep();
}

int QMdiSubWindow_KeyboardSingleStep(void* ptr){
	return static_cast<QMdiSubWindow*>(ptr)->keyboardSingleStep();
}

void QMdiSubWindow_SetKeyboardPageStep(void* ptr, int step){
	static_cast<QMdiSubWindow*>(ptr)->setKeyboardPageStep(step);
}

void QMdiSubWindow_SetKeyboardSingleStep(void* ptr, int step){
	static_cast<QMdiSubWindow*>(ptr)->setKeyboardSingleStep(step);
}

void* QMdiSubWindow_NewQMdiSubWindow(void* parent, int flags){
	return new QMdiSubWindow(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(flags));
}

void QMdiSubWindow_ConnectAboutToActivate(void* ptr){
	QObject::connect(static_cast<QMdiSubWindow*>(ptr), static_cast<void (QMdiSubWindow::*)()>(&QMdiSubWindow::aboutToActivate), static_cast<MyQMdiSubWindow*>(ptr), static_cast<void (MyQMdiSubWindow::*)()>(&MyQMdiSubWindow::Signal_AboutToActivate));;
}

void QMdiSubWindow_DisconnectAboutToActivate(void* ptr){
	QObject::disconnect(static_cast<QMdiSubWindow*>(ptr), static_cast<void (QMdiSubWindow::*)()>(&QMdiSubWindow::aboutToActivate), static_cast<MyQMdiSubWindow*>(ptr), static_cast<void (MyQMdiSubWindow::*)()>(&MyQMdiSubWindow::Signal_AboutToActivate));;
}

int QMdiSubWindow_IsShaded(void* ptr){
	return static_cast<QMdiSubWindow*>(ptr)->isShaded();
}

void* QMdiSubWindow_MdiArea(void* ptr){
	return static_cast<QMdiSubWindow*>(ptr)->mdiArea();
}

void QMdiSubWindow_SetOption(void* ptr, int option, int on){
	static_cast<QMdiSubWindow*>(ptr)->setOption(static_cast<QMdiSubWindow::SubWindowOption>(option), on != 0);
}

void QMdiSubWindow_SetSystemMenu(void* ptr, void* systemMenu){
	static_cast<QMdiSubWindow*>(ptr)->setSystemMenu(static_cast<QMenu*>(systemMenu));
}

void QMdiSubWindow_SetWidget(void* ptr, void* widget){
	static_cast<QMdiSubWindow*>(ptr)->setWidget(static_cast<QWidget*>(widget));
}

void QMdiSubWindow_ShowShaded(void* ptr){
	QMetaObject::invokeMethod(static_cast<QMdiSubWindow*>(ptr), "showShaded");
}

void QMdiSubWindow_ShowSystemMenu(void* ptr){
	QMetaObject::invokeMethod(static_cast<QMdiSubWindow*>(ptr), "showSystemMenu");
}

void* QMdiSubWindow_SystemMenu(void* ptr){
	return static_cast<QMdiSubWindow*>(ptr)->systemMenu();
}

int QMdiSubWindow_TestOption(void* ptr, int option){
	return static_cast<QMdiSubWindow*>(ptr)->testOption(static_cast<QMdiSubWindow::SubWindowOption>(option));
}

void* QMdiSubWindow_Widget(void* ptr){
	return static_cast<QMdiSubWindow*>(ptr)->widget();
}

void QMdiSubWindow_ConnectWindowStateChanged(void* ptr){
	QObject::connect(static_cast<QMdiSubWindow*>(ptr), static_cast<void (QMdiSubWindow::*)(Qt::WindowStates, Qt::WindowStates)>(&QMdiSubWindow::windowStateChanged), static_cast<MyQMdiSubWindow*>(ptr), static_cast<void (MyQMdiSubWindow::*)(Qt::WindowStates, Qt::WindowStates)>(&MyQMdiSubWindow::Signal_WindowStateChanged));;
}

void QMdiSubWindow_DisconnectWindowStateChanged(void* ptr){
	QObject::disconnect(static_cast<QMdiSubWindow*>(ptr), static_cast<void (QMdiSubWindow::*)(Qt::WindowStates, Qt::WindowStates)>(&QMdiSubWindow::windowStateChanged), static_cast<MyQMdiSubWindow*>(ptr), static_cast<void (MyQMdiSubWindow::*)(Qt::WindowStates, Qt::WindowStates)>(&MyQMdiSubWindow::Signal_WindowStateChanged));;
}

void QMdiSubWindow_DestroyQMdiSubWindow(void* ptr){
	static_cast<QMdiSubWindow*>(ptr)->~QMdiSubWindow();
}

