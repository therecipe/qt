#include "qquickwindow.h"
#include <QWindow>
#include <QOpenGLFramebufferObject>
#include <QQuickItem>
#include <QOpenGLContext>
#include <QRunnable>
#include <QUrl>
#include <QMetaObject>
#include <QEvent>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QModelIndex>
#include <QImage>
#include <QColor>
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
void Signal_ColorChanged(const QColor & v){callbackQQuickWindowColorChanged(this->objectName().toUtf8().data(), new QColor(v));};
void Signal_FrameSwapped(){callbackQQuickWindowFrameSwapped(this->objectName().toUtf8().data());};
void Signal_OpenglContextCreated(QOpenGLContext * context){callbackQQuickWindowOpenglContextCreated(this->objectName().toUtf8().data(), context);};
void Signal_SceneGraphAboutToStop(){callbackQQuickWindowSceneGraphAboutToStop(this->objectName().toUtf8().data());};
void Signal_SceneGraphError(QQuickWindow::SceneGraphError error, const QString & message){callbackQQuickWindowSceneGraphError(this->objectName().toUtf8().data(), error, message.toUtf8().data());};
void Signal_SceneGraphInitialized(){callbackQQuickWindowSceneGraphInitialized(this->objectName().toUtf8().data());};
void Signal_SceneGraphInvalidated(){callbackQQuickWindowSceneGraphInvalidated(this->objectName().toUtf8().data());};
};

void* QQuickWindow_ActiveFocusItem(void* ptr){
	return static_cast<QQuickWindow*>(ptr)->activeFocusItem();
}

void* QQuickWindow_Color(void* ptr){
	return new QColor(static_cast<QQuickWindow*>(ptr)->color());
}

void* QQuickWindow_ContentItem(void* ptr){
	return static_cast<QQuickWindow*>(ptr)->contentItem();
}

void QQuickWindow_SetColor(void* ptr, void* color){
	static_cast<QQuickWindow*>(ptr)->setColor(*static_cast<QColor*>(color));
}

void* QQuickWindow_NewQQuickWindow(void* parent){
	return new QQuickWindow(static_cast<QWindow*>(parent));
}

void* QQuickWindow_AccessibleRoot(void* ptr){
	return static_cast<QQuickWindow*>(ptr)->accessibleRoot();
}

void QQuickWindow_ConnectActiveFocusItemChanged(void* ptr){
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::activeFocusItemChanged), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_ActiveFocusItemChanged));;
}

void QQuickWindow_DisconnectActiveFocusItemChanged(void* ptr){
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::activeFocusItemChanged), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_ActiveFocusItemChanged));;
}

void QQuickWindow_ConnectAfterAnimating(void* ptr){
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::afterAnimating), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_AfterAnimating));;
}

void QQuickWindow_DisconnectAfterAnimating(void* ptr){
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::afterAnimating), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_AfterAnimating));;
}

void QQuickWindow_ConnectAfterRendering(void* ptr){
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::afterRendering), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_AfterRendering));;
}

void QQuickWindow_DisconnectAfterRendering(void* ptr){
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::afterRendering), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_AfterRendering));;
}

void QQuickWindow_ConnectAfterSynchronizing(void* ptr){
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::afterSynchronizing), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_AfterSynchronizing));;
}

void QQuickWindow_DisconnectAfterSynchronizing(void* ptr){
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::afterSynchronizing), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_AfterSynchronizing));;
}

void QQuickWindow_ConnectBeforeRendering(void* ptr){
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::beforeRendering), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_BeforeRendering));;
}

void QQuickWindow_DisconnectBeforeRendering(void* ptr){
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::beforeRendering), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_BeforeRendering));;
}

void QQuickWindow_ConnectBeforeSynchronizing(void* ptr){
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::beforeSynchronizing), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_BeforeSynchronizing));;
}

void QQuickWindow_DisconnectBeforeSynchronizing(void* ptr){
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::beforeSynchronizing), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_BeforeSynchronizing));;
}

int QQuickWindow_ClearBeforeRendering(void* ptr){
	return static_cast<QQuickWindow*>(ptr)->clearBeforeRendering();
}

void QQuickWindow_ConnectColorChanged(void* ptr){
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)(const QColor &)>(&QQuickWindow::colorChanged), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)(const QColor &)>(&MyQQuickWindow::Signal_ColorChanged));;
}

void QQuickWindow_DisconnectColorChanged(void* ptr){
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)(const QColor &)>(&QQuickWindow::colorChanged), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)(const QColor &)>(&MyQQuickWindow::Signal_ColorChanged));;
}

void* QQuickWindow_CreateTextureFromImage2(void* ptr, void* image){
	return static_cast<QQuickWindow*>(ptr)->createTextureFromImage(*static_cast<QImage*>(image));
}

void* QQuickWindow_CreateTextureFromImage(void* ptr, void* image, int options){
	return static_cast<QQuickWindow*>(ptr)->createTextureFromImage(*static_cast<QImage*>(image), static_cast<QQuickWindow::CreateTextureOption>(options));
}

double QQuickWindow_EffectiveDevicePixelRatio(void* ptr){
	return static_cast<double>(static_cast<QQuickWindow*>(ptr)->effectiveDevicePixelRatio());
}

void QQuickWindow_ConnectFrameSwapped(void* ptr){
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::frameSwapped), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_FrameSwapped));;
}

void QQuickWindow_DisconnectFrameSwapped(void* ptr){
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::frameSwapped), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_FrameSwapped));;
}

int QQuickWindow_QQuickWindow_HasDefaultAlphaBuffer(){
	return QQuickWindow::hasDefaultAlphaBuffer();
}

void* QQuickWindow_IncubationController(void* ptr){
	return static_cast<QQuickWindow*>(ptr)->incubationController();
}

int QQuickWindow_IsPersistentOpenGLContext(void* ptr){
	return static_cast<QQuickWindow*>(ptr)->isPersistentOpenGLContext();
}

int QQuickWindow_IsPersistentSceneGraph(void* ptr){
	return static_cast<QQuickWindow*>(ptr)->isPersistentSceneGraph();
}

int QQuickWindow_IsSceneGraphInitialized(void* ptr){
	return static_cast<QQuickWindow*>(ptr)->isSceneGraphInitialized();
}

void* QQuickWindow_MouseGrabberItem(void* ptr){
	return static_cast<QQuickWindow*>(ptr)->mouseGrabberItem();
}

void* QQuickWindow_OpenglContext(void* ptr){
	return static_cast<QQuickWindow*>(ptr)->openglContext();
}

void QQuickWindow_ConnectOpenglContextCreated(void* ptr){
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)(QOpenGLContext *)>(&QQuickWindow::openglContextCreated), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)(QOpenGLContext *)>(&MyQQuickWindow::Signal_OpenglContextCreated));;
}

void QQuickWindow_DisconnectOpenglContextCreated(void* ptr){
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)(QOpenGLContext *)>(&QQuickWindow::openglContextCreated), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)(QOpenGLContext *)>(&MyQQuickWindow::Signal_OpenglContextCreated));;
}

void QQuickWindow_ReleaseResources(void* ptr){
	QMetaObject::invokeMethod(static_cast<QQuickWindow*>(ptr), "releaseResources");
}

void* QQuickWindow_RenderTarget(void* ptr){
	return static_cast<QQuickWindow*>(ptr)->renderTarget();
}

void QQuickWindow_ResetOpenGLState(void* ptr){
	static_cast<QQuickWindow*>(ptr)->resetOpenGLState();
}

void QQuickWindow_ConnectSceneGraphAboutToStop(void* ptr){
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::sceneGraphAboutToStop), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_SceneGraphAboutToStop));;
}

void QQuickWindow_DisconnectSceneGraphAboutToStop(void* ptr){
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::sceneGraphAboutToStop), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_SceneGraphAboutToStop));;
}

void QQuickWindow_ConnectSceneGraphError(void* ptr){
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)(QQuickWindow::SceneGraphError, const QString &)>(&QQuickWindow::sceneGraphError), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)(QQuickWindow::SceneGraphError, const QString &)>(&MyQQuickWindow::Signal_SceneGraphError));;
}

void QQuickWindow_DisconnectSceneGraphError(void* ptr){
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)(QQuickWindow::SceneGraphError, const QString &)>(&QQuickWindow::sceneGraphError), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)(QQuickWindow::SceneGraphError, const QString &)>(&MyQQuickWindow::Signal_SceneGraphError));;
}

void QQuickWindow_ConnectSceneGraphInitialized(void* ptr){
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::sceneGraphInitialized), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_SceneGraphInitialized));;
}

void QQuickWindow_DisconnectSceneGraphInitialized(void* ptr){
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::sceneGraphInitialized), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_SceneGraphInitialized));;
}

void QQuickWindow_ConnectSceneGraphInvalidated(void* ptr){
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::sceneGraphInvalidated), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_SceneGraphInvalidated));;
}

void QQuickWindow_DisconnectSceneGraphInvalidated(void* ptr){
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::sceneGraphInvalidated), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_SceneGraphInvalidated));;
}

void QQuickWindow_ScheduleRenderJob(void* ptr, void* job, int stage){
	static_cast<QQuickWindow*>(ptr)->scheduleRenderJob(static_cast<QRunnable*>(job), static_cast<QQuickWindow::RenderStage>(stage));
}

int QQuickWindow_SendEvent(void* ptr, void* item, void* e){
	return static_cast<QQuickWindow*>(ptr)->sendEvent(static_cast<QQuickItem*>(item), static_cast<QEvent*>(e));
}

void QQuickWindow_SetClearBeforeRendering(void* ptr, int enabled){
	static_cast<QQuickWindow*>(ptr)->setClearBeforeRendering(enabled != 0);
}

void QQuickWindow_QQuickWindow_SetDefaultAlphaBuffer(int useAlpha){
	QQuickWindow::setDefaultAlphaBuffer(useAlpha != 0);
}

void QQuickWindow_SetPersistentOpenGLContext(void* ptr, int persistent){
	static_cast<QQuickWindow*>(ptr)->setPersistentOpenGLContext(persistent != 0);
}

void QQuickWindow_SetPersistentSceneGraph(void* ptr, int persistent){
	static_cast<QQuickWindow*>(ptr)->setPersistentSceneGraph(persistent != 0);
}

void QQuickWindow_SetRenderTarget(void* ptr, void* fbo){
	static_cast<QQuickWindow*>(ptr)->setRenderTarget(static_cast<QOpenGLFramebufferObject*>(fbo));
}

void QQuickWindow_Update(void* ptr){
	QMetaObject::invokeMethod(static_cast<QQuickWindow*>(ptr), "update");
}

void QQuickWindow_DestroyQQuickWindow(void* ptr){
	static_cast<QQuickWindow*>(ptr)->~QQuickWindow();
}

