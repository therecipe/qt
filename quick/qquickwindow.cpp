#include "qquickwindow.h"
#include <QColor>
#include <QOpenGLContext>
#include <QEvent>
#include <QMetaObject>
#include <QRunnable>
#include <QVariant>
#include <QWindow>
#include <QImage>
#include <QOpenGLFramebufferObject>
#include <QObject>
#include <QUrl>
#include <QModelIndex>
#include <QQuickItem>
#include <QString>
#include <QQuickWindow>
#include "_cgo_export.h"

class MyQQuickWindow: public QQuickWindow {
public:
void Signal_ActiveFocusItemChanged(){callbackQQuickWindowActiveFocusItemChanged(this->objectName().toUtf8().data());};
void Signal_AfterAnimating(){callbackQQuickWindowAfterAnimating(this->objectName().toUtf8().data());};
void Signal_AfterRendering(){callbackQQuickWindowAfterRendering(this->objectName().toUtf8().data());};
void Signal_AfterSynchronizing(){callbackQQuickWindowAfterSynchronizing(this->objectName().toUtf8().data());};
void Signal_BeforeRendering(){callbackQQuickWindowBeforeRendering(this->objectName().toUtf8().data());};
void Signal_BeforeSynchronizing(){callbackQQuickWindowBeforeSynchronizing(this->objectName().toUtf8().data());};
void Signal_FrameSwapped(){callbackQQuickWindowFrameSwapped(this->objectName().toUtf8().data());};
void Signal_OpenglContextCreated(QOpenGLContext * context){callbackQQuickWindowOpenglContextCreated(this->objectName().toUtf8().data(), context);};
void Signal_SceneGraphAboutToStop(){callbackQQuickWindowSceneGraphAboutToStop(this->objectName().toUtf8().data());};
void Signal_SceneGraphError(QQuickWindow::SceneGraphError error, const QString & message){callbackQQuickWindowSceneGraphError(this->objectName().toUtf8().data(), error, message.toUtf8().data());};
void Signal_SceneGraphInitialized(){callbackQQuickWindowSceneGraphInitialized(this->objectName().toUtf8().data());};
void Signal_SceneGraphInvalidated(){callbackQQuickWindowSceneGraphInvalidated(this->objectName().toUtf8().data());};
};

QtObjectPtr QQuickWindow_ActiveFocusItem(QtObjectPtr ptr){
	return static_cast<QQuickWindow*>(ptr)->activeFocusItem();
}

QtObjectPtr QQuickWindow_ContentItem(QtObjectPtr ptr){
	return static_cast<QQuickWindow*>(ptr)->contentItem();
}

void QQuickWindow_SetColor(QtObjectPtr ptr, QtObjectPtr color){
	static_cast<QQuickWindow*>(ptr)->setColor(*static_cast<QColor*>(color));
}

QtObjectPtr QQuickWindow_NewQQuickWindow(QtObjectPtr parent){
	return new QQuickWindow(static_cast<QWindow*>(parent));
}

QtObjectPtr QQuickWindow_AccessibleRoot(QtObjectPtr ptr){
	return static_cast<QQuickWindow*>(ptr)->accessibleRoot();
}

void QQuickWindow_ConnectActiveFocusItemChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::activeFocusItemChanged), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_ActiveFocusItemChanged));;
}

void QQuickWindow_DisconnectActiveFocusItemChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::activeFocusItemChanged), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_ActiveFocusItemChanged));;
}

void QQuickWindow_ConnectAfterAnimating(QtObjectPtr ptr){
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::afterAnimating), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_AfterAnimating));;
}

void QQuickWindow_DisconnectAfterAnimating(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::afterAnimating), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_AfterAnimating));;
}

void QQuickWindow_ConnectAfterRendering(QtObjectPtr ptr){
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::afterRendering), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_AfterRendering));;
}

void QQuickWindow_DisconnectAfterRendering(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::afterRendering), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_AfterRendering));;
}

void QQuickWindow_ConnectAfterSynchronizing(QtObjectPtr ptr){
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::afterSynchronizing), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_AfterSynchronizing));;
}

void QQuickWindow_DisconnectAfterSynchronizing(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::afterSynchronizing), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_AfterSynchronizing));;
}

void QQuickWindow_ConnectBeforeRendering(QtObjectPtr ptr){
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::beforeRendering), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_BeforeRendering));;
}

void QQuickWindow_DisconnectBeforeRendering(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::beforeRendering), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_BeforeRendering));;
}

void QQuickWindow_ConnectBeforeSynchronizing(QtObjectPtr ptr){
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::beforeSynchronizing), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_BeforeSynchronizing));;
}

void QQuickWindow_DisconnectBeforeSynchronizing(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::beforeSynchronizing), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_BeforeSynchronizing));;
}

int QQuickWindow_ClearBeforeRendering(QtObjectPtr ptr){
	return static_cast<QQuickWindow*>(ptr)->clearBeforeRendering();
}

QtObjectPtr QQuickWindow_CreateTextureFromImage2(QtObjectPtr ptr, QtObjectPtr image){
	return static_cast<QQuickWindow*>(ptr)->createTextureFromImage(*static_cast<QImage*>(image));
}

QtObjectPtr QQuickWindow_CreateTextureFromImage(QtObjectPtr ptr, QtObjectPtr image, int options){
	return static_cast<QQuickWindow*>(ptr)->createTextureFromImage(*static_cast<QImage*>(image), static_cast<QQuickWindow::CreateTextureOption>(options));
}

void QQuickWindow_ConnectFrameSwapped(QtObjectPtr ptr){
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::frameSwapped), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_FrameSwapped));;
}

void QQuickWindow_DisconnectFrameSwapped(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::frameSwapped), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_FrameSwapped));;
}

int QQuickWindow_QQuickWindow_HasDefaultAlphaBuffer(){
	return QQuickWindow::hasDefaultAlphaBuffer();
}

QtObjectPtr QQuickWindow_IncubationController(QtObjectPtr ptr){
	return static_cast<QQuickWindow*>(ptr)->incubationController();
}

int QQuickWindow_IsPersistentOpenGLContext(QtObjectPtr ptr){
	return static_cast<QQuickWindow*>(ptr)->isPersistentOpenGLContext();
}

int QQuickWindow_IsPersistentSceneGraph(QtObjectPtr ptr){
	return static_cast<QQuickWindow*>(ptr)->isPersistentSceneGraph();
}

int QQuickWindow_IsSceneGraphInitialized(QtObjectPtr ptr){
	return static_cast<QQuickWindow*>(ptr)->isSceneGraphInitialized();
}

QtObjectPtr QQuickWindow_MouseGrabberItem(QtObjectPtr ptr){
	return static_cast<QQuickWindow*>(ptr)->mouseGrabberItem();
}

QtObjectPtr QQuickWindow_OpenglContext(QtObjectPtr ptr){
	return static_cast<QQuickWindow*>(ptr)->openglContext();
}

void QQuickWindow_ConnectOpenglContextCreated(QtObjectPtr ptr){
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)(QOpenGLContext *)>(&QQuickWindow::openglContextCreated), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)(QOpenGLContext *)>(&MyQQuickWindow::Signal_OpenglContextCreated));;
}

void QQuickWindow_DisconnectOpenglContextCreated(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)(QOpenGLContext *)>(&QQuickWindow::openglContextCreated), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)(QOpenGLContext *)>(&MyQQuickWindow::Signal_OpenglContextCreated));;
}

void QQuickWindow_ReleaseResources(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QQuickWindow*>(ptr), "releaseResources");
}

QtObjectPtr QQuickWindow_RenderTarget(QtObjectPtr ptr){
	return static_cast<QQuickWindow*>(ptr)->renderTarget();
}

void QQuickWindow_ResetOpenGLState(QtObjectPtr ptr){
	static_cast<QQuickWindow*>(ptr)->resetOpenGLState();
}

void QQuickWindow_ConnectSceneGraphAboutToStop(QtObjectPtr ptr){
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::sceneGraphAboutToStop), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_SceneGraphAboutToStop));;
}

void QQuickWindow_DisconnectSceneGraphAboutToStop(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::sceneGraphAboutToStop), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_SceneGraphAboutToStop));;
}

void QQuickWindow_ConnectSceneGraphError(QtObjectPtr ptr){
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)(QQuickWindow::SceneGraphError, const QString &)>(&QQuickWindow::sceneGraphError), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)(QQuickWindow::SceneGraphError, const QString &)>(&MyQQuickWindow::Signal_SceneGraphError));;
}

void QQuickWindow_DisconnectSceneGraphError(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)(QQuickWindow::SceneGraphError, const QString &)>(&QQuickWindow::sceneGraphError), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)(QQuickWindow::SceneGraphError, const QString &)>(&MyQQuickWindow::Signal_SceneGraphError));;
}

void QQuickWindow_ConnectSceneGraphInitialized(QtObjectPtr ptr){
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::sceneGraphInitialized), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_SceneGraphInitialized));;
}

void QQuickWindow_DisconnectSceneGraphInitialized(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::sceneGraphInitialized), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_SceneGraphInitialized));;
}

void QQuickWindow_ConnectSceneGraphInvalidated(QtObjectPtr ptr){
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::sceneGraphInvalidated), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_SceneGraphInvalidated));;
}

void QQuickWindow_DisconnectSceneGraphInvalidated(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::sceneGraphInvalidated), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_SceneGraphInvalidated));;
}

void QQuickWindow_ScheduleRenderJob(QtObjectPtr ptr, QtObjectPtr job, int stage){
	static_cast<QQuickWindow*>(ptr)->scheduleRenderJob(static_cast<QRunnable*>(job), static_cast<QQuickWindow::RenderStage>(stage));
}

int QQuickWindow_SendEvent(QtObjectPtr ptr, QtObjectPtr item, QtObjectPtr e){
	return static_cast<QQuickWindow*>(ptr)->sendEvent(static_cast<QQuickItem*>(item), static_cast<QEvent*>(e));
}

void QQuickWindow_SetClearBeforeRendering(QtObjectPtr ptr, int enabled){
	static_cast<QQuickWindow*>(ptr)->setClearBeforeRendering(enabled != 0);
}

void QQuickWindow_QQuickWindow_SetDefaultAlphaBuffer(int useAlpha){
	QQuickWindow::setDefaultAlphaBuffer(useAlpha != 0);
}

void QQuickWindow_SetPersistentOpenGLContext(QtObjectPtr ptr, int persistent){
	static_cast<QQuickWindow*>(ptr)->setPersistentOpenGLContext(persistent != 0);
}

void QQuickWindow_SetPersistentSceneGraph(QtObjectPtr ptr, int persistent){
	static_cast<QQuickWindow*>(ptr)->setPersistentSceneGraph(persistent != 0);
}

void QQuickWindow_SetRenderTarget(QtObjectPtr ptr, QtObjectPtr fbo){
	static_cast<QQuickWindow*>(ptr)->setRenderTarget(static_cast<QOpenGLFramebufferObject*>(fbo));
}

void QQuickWindow_Update(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QQuickWindow*>(ptr), "update");
}

void QQuickWindow_DestroyQQuickWindow(QtObjectPtr ptr){
	static_cast<QQuickWindow*>(ptr)->~QQuickWindow();
}

