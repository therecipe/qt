#include "qmdisubwindow.h"
#include <QModelIndex>
#include <QMetaObject>
#include <QObject>
#include <QWidget>
#include <QMenu>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QMdiSubWindow>
#include "_cgo_export.h"

class MyQMdiSubWindow: public QMdiSubWindow {
public:
void Signal_AboutToActivate(){callbackQMdiSubWindowAboutToActivate(this->objectName().toUtf8().data());};
void Signal_WindowStateChanged(Qt::WindowStates oldState, Qt::WindowStates newState){callbackQMdiSubWindowWindowStateChanged(this->objectName().toUtf8().data(), oldState, newState);};
};

int QMdiSubWindow_KeyboardPageStep(QtObjectPtr ptr){
	return static_cast<QMdiSubWindow*>(ptr)->keyboardPageStep();
}

int QMdiSubWindow_KeyboardSingleStep(QtObjectPtr ptr){
	return static_cast<QMdiSubWindow*>(ptr)->keyboardSingleStep();
}

void QMdiSubWindow_SetKeyboardPageStep(QtObjectPtr ptr, int step){
	static_cast<QMdiSubWindow*>(ptr)->setKeyboardPageStep(step);
}

void QMdiSubWindow_SetKeyboardSingleStep(QtObjectPtr ptr, int step){
	static_cast<QMdiSubWindow*>(ptr)->setKeyboardSingleStep(step);
}

QtObjectPtr QMdiSubWindow_NewQMdiSubWindow(QtObjectPtr parent, int flags){
	return new QMdiSubWindow(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(flags));
}

void QMdiSubWindow_ConnectAboutToActivate(QtObjectPtr ptr){
	QObject::connect(static_cast<QMdiSubWindow*>(ptr), static_cast<void (QMdiSubWindow::*)()>(&QMdiSubWindow::aboutToActivate), static_cast<MyQMdiSubWindow*>(ptr), static_cast<void (MyQMdiSubWindow::*)()>(&MyQMdiSubWindow::Signal_AboutToActivate));;
}

void QMdiSubWindow_DisconnectAboutToActivate(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMdiSubWindow*>(ptr), static_cast<void (QMdiSubWindow::*)()>(&QMdiSubWindow::aboutToActivate), static_cast<MyQMdiSubWindow*>(ptr), static_cast<void (MyQMdiSubWindow::*)()>(&MyQMdiSubWindow::Signal_AboutToActivate));;
}

int QMdiSubWindow_IsShaded(QtObjectPtr ptr){
	return static_cast<QMdiSubWindow*>(ptr)->isShaded();
}

QtObjectPtr QMdiSubWindow_MdiArea(QtObjectPtr ptr){
	return static_cast<QMdiSubWindow*>(ptr)->mdiArea();
}

void QMdiSubWindow_SetOption(QtObjectPtr ptr, int option, int on){
	static_cast<QMdiSubWindow*>(ptr)->setOption(static_cast<QMdiSubWindow::SubWindowOption>(option), on != 0);
}

void QMdiSubWindow_SetSystemMenu(QtObjectPtr ptr, QtObjectPtr systemMenu){
	static_cast<QMdiSubWindow*>(ptr)->setSystemMenu(static_cast<QMenu*>(systemMenu));
}

void QMdiSubWindow_SetWidget(QtObjectPtr ptr, QtObjectPtr widget){
	static_cast<QMdiSubWindow*>(ptr)->setWidget(static_cast<QWidget*>(widget));
}

void QMdiSubWindow_ShowShaded(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QMdiSubWindow*>(ptr), "showShaded");
}

void QMdiSubWindow_ShowSystemMenu(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QMdiSubWindow*>(ptr), "showSystemMenu");
}

QtObjectPtr QMdiSubWindow_SystemMenu(QtObjectPtr ptr){
	return static_cast<QMdiSubWindow*>(ptr)->systemMenu();
}

int QMdiSubWindow_TestOption(QtObjectPtr ptr, int option){
	return static_cast<QMdiSubWindow*>(ptr)->testOption(static_cast<QMdiSubWindow::SubWindowOption>(option));
}

QtObjectPtr QMdiSubWindow_Widget(QtObjectPtr ptr){
	return static_cast<QMdiSubWindow*>(ptr)->widget();
}

void QMdiSubWindow_ConnectWindowStateChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QMdiSubWindow*>(ptr), static_cast<void (QMdiSubWindow::*)(Qt::WindowStates, Qt::WindowStates)>(&QMdiSubWindow::windowStateChanged), static_cast<MyQMdiSubWindow*>(ptr), static_cast<void (MyQMdiSubWindow::*)(Qt::WindowStates, Qt::WindowStates)>(&MyQMdiSubWindow::Signal_WindowStateChanged));;
}

void QMdiSubWindow_DisconnectWindowStateChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMdiSubWindow*>(ptr), static_cast<void (QMdiSubWindow::*)(Qt::WindowStates, Qt::WindowStates)>(&QMdiSubWindow::windowStateChanged), static_cast<MyQMdiSubWindow*>(ptr), static_cast<void (MyQMdiSubWindow::*)(Qt::WindowStates, Qt::WindowStates)>(&MyQMdiSubWindow::Signal_WindowStateChanged));;
}

void QMdiSubWindow_DestroyQMdiSubWindow(QtObjectPtr ptr){
	static_cast<QMdiSubWindow*>(ptr)->~QMdiSubWindow();
}

