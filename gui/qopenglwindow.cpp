#include "qopenglwindow.h"
#include <QObject>
#include <QOpenGLContext>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWindow>
#include <QOpenGLWindow>
#include "_cgo_export.h"

class MyQOpenGLWindow: public QOpenGLWindow {
public:
void Signal_FrameSwapped(){callbackQOpenGLWindowFrameSwapped(this->objectName().toUtf8().data());};
};

QtObjectPtr QOpenGLWindow_NewQOpenGLWindow2(QtObjectPtr shareContext, int updateBehavior, QtObjectPtr parent){
	return new QOpenGLWindow(static_cast<QOpenGLContext*>(shareContext), static_cast<QOpenGLWindow::UpdateBehavior>(updateBehavior), static_cast<QWindow*>(parent));
}

QtObjectPtr QOpenGLWindow_NewQOpenGLWindow(int updateBehavior, QtObjectPtr parent){
	return new QOpenGLWindow(static_cast<QOpenGLWindow::UpdateBehavior>(updateBehavior), static_cast<QWindow*>(parent));
}

QtObjectPtr QOpenGLWindow_Context(QtObjectPtr ptr){
	return static_cast<QOpenGLWindow*>(ptr)->context();
}

void QOpenGLWindow_DoneCurrent(QtObjectPtr ptr){
	static_cast<QOpenGLWindow*>(ptr)->doneCurrent();
}

void QOpenGLWindow_ConnectFrameSwapped(QtObjectPtr ptr){
	QObject::connect(static_cast<QOpenGLWindow*>(ptr), static_cast<void (QOpenGLWindow::*)()>(&QOpenGLWindow::frameSwapped), static_cast<MyQOpenGLWindow*>(ptr), static_cast<void (MyQOpenGLWindow::*)()>(&MyQOpenGLWindow::Signal_FrameSwapped));;
}

void QOpenGLWindow_DisconnectFrameSwapped(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QOpenGLWindow*>(ptr), static_cast<void (QOpenGLWindow::*)()>(&QOpenGLWindow::frameSwapped), static_cast<MyQOpenGLWindow*>(ptr), static_cast<void (MyQOpenGLWindow::*)()>(&MyQOpenGLWindow::Signal_FrameSwapped));;
}

int QOpenGLWindow_IsValid(QtObjectPtr ptr){
	return static_cast<QOpenGLWindow*>(ptr)->isValid();
}

void QOpenGLWindow_MakeCurrent(QtObjectPtr ptr){
	static_cast<QOpenGLWindow*>(ptr)->makeCurrent();
}

QtObjectPtr QOpenGLWindow_ShareContext(QtObjectPtr ptr){
	return static_cast<QOpenGLWindow*>(ptr)->shareContext();
}

int QOpenGLWindow_UpdateBehavior(QtObjectPtr ptr){
	return static_cast<QOpenGLWindow*>(ptr)->updateBehavior();
}

void QOpenGLWindow_DestroyQOpenGLWindow(QtObjectPtr ptr){
	static_cast<QOpenGLWindow*>(ptr)->~QOpenGLWindow();
}

