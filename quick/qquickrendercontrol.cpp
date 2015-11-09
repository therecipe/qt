#include "qquickrendercontrol.h"
#include <QObject>
#include <QThread>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QOpenGLContext>
#include <QModelIndex>
#include <QPoint>
#include <QQuickWindow>
#include <QQuickRenderControl>
#include "_cgo_export.h"

class MyQQuickRenderControl: public QQuickRenderControl {
public:
void Signal_RenderRequested(){callbackQQuickRenderControlRenderRequested(this->objectName().toUtf8().data());};
void Signal_SceneChanged(){callbackQQuickRenderControlSceneChanged(this->objectName().toUtf8().data());};
};

void* QQuickRenderControl_NewQQuickRenderControl(void* parent){
	return new QQuickRenderControl(static_cast<QObject*>(parent));
}

void QQuickRenderControl_Initialize(void* ptr, void* gl){
	static_cast<QQuickRenderControl*>(ptr)->initialize(static_cast<QOpenGLContext*>(gl));
}

void QQuickRenderControl_Invalidate(void* ptr){
	static_cast<QQuickRenderControl*>(ptr)->invalidate();
}

void QQuickRenderControl_PolishItems(void* ptr){
	static_cast<QQuickRenderControl*>(ptr)->polishItems();
}

void QQuickRenderControl_PrepareThread(void* ptr, void* targetThread){
	static_cast<QQuickRenderControl*>(ptr)->prepareThread(static_cast<QThread*>(targetThread));
}

void QQuickRenderControl_Render(void* ptr){
	static_cast<QQuickRenderControl*>(ptr)->render();
}

void QQuickRenderControl_ConnectRenderRequested(void* ptr){
	QObject::connect(static_cast<QQuickRenderControl*>(ptr), static_cast<void (QQuickRenderControl::*)()>(&QQuickRenderControl::renderRequested), static_cast<MyQQuickRenderControl*>(ptr), static_cast<void (MyQQuickRenderControl::*)()>(&MyQQuickRenderControl::Signal_RenderRequested));;
}

void QQuickRenderControl_DisconnectRenderRequested(void* ptr){
	QObject::disconnect(static_cast<QQuickRenderControl*>(ptr), static_cast<void (QQuickRenderControl::*)()>(&QQuickRenderControl::renderRequested), static_cast<MyQQuickRenderControl*>(ptr), static_cast<void (MyQQuickRenderControl::*)()>(&MyQQuickRenderControl::Signal_RenderRequested));;
}

void* QQuickRenderControl_RenderWindow(void* ptr, void* offset){
	return static_cast<QQuickRenderControl*>(ptr)->renderWindow(static_cast<QPoint*>(offset));
}

void* QQuickRenderControl_QQuickRenderControl_RenderWindowFor(void* win, void* offset){
	return QQuickRenderControl::renderWindowFor(static_cast<QQuickWindow*>(win), static_cast<QPoint*>(offset));
}

void QQuickRenderControl_ConnectSceneChanged(void* ptr){
	QObject::connect(static_cast<QQuickRenderControl*>(ptr), static_cast<void (QQuickRenderControl::*)()>(&QQuickRenderControl::sceneChanged), static_cast<MyQQuickRenderControl*>(ptr), static_cast<void (MyQQuickRenderControl::*)()>(&MyQQuickRenderControl::Signal_SceneChanged));;
}

void QQuickRenderControl_DisconnectSceneChanged(void* ptr){
	QObject::disconnect(static_cast<QQuickRenderControl*>(ptr), static_cast<void (QQuickRenderControl::*)()>(&QQuickRenderControl::sceneChanged), static_cast<MyQQuickRenderControl*>(ptr), static_cast<void (MyQQuickRenderControl::*)()>(&MyQQuickRenderControl::Signal_SceneChanged));;
}

int QQuickRenderControl_Sync(void* ptr){
	return static_cast<QQuickRenderControl*>(ptr)->sync();
}

void QQuickRenderControl_DestroyQQuickRenderControl(void* ptr){
	static_cast<QQuickRenderControl*>(ptr)->~QQuickRenderControl();
}

