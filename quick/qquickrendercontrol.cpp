#include "qquickrendercontrol.h"
#include <QVariant>
#include <QModelIndex>
#include <QThread>
#include <QOpenGLContext>
#include <QQuickWindow>
#include <QString>
#include <QUrl>
#include <QObject>
#include <QPoint>
#include <QQuickRenderControl>
#include "_cgo_export.h"

class MyQQuickRenderControl: public QQuickRenderControl {
public:
void Signal_RenderRequested(){callbackQQuickRenderControlRenderRequested(this->objectName().toUtf8().data());};
void Signal_SceneChanged(){callbackQQuickRenderControlSceneChanged(this->objectName().toUtf8().data());};
};

QtObjectPtr QQuickRenderControl_NewQQuickRenderControl(QtObjectPtr parent){
	return new QQuickRenderControl(static_cast<QObject*>(parent));
}

void QQuickRenderControl_Initialize(QtObjectPtr ptr, QtObjectPtr gl){
	static_cast<QQuickRenderControl*>(ptr)->initialize(static_cast<QOpenGLContext*>(gl));
}

void QQuickRenderControl_Invalidate(QtObjectPtr ptr){
	static_cast<QQuickRenderControl*>(ptr)->invalidate();
}

void QQuickRenderControl_PolishItems(QtObjectPtr ptr){
	static_cast<QQuickRenderControl*>(ptr)->polishItems();
}

void QQuickRenderControl_PrepareThread(QtObjectPtr ptr, QtObjectPtr targetThread){
	static_cast<QQuickRenderControl*>(ptr)->prepareThread(static_cast<QThread*>(targetThread));
}

void QQuickRenderControl_Render(QtObjectPtr ptr){
	static_cast<QQuickRenderControl*>(ptr)->render();
}

void QQuickRenderControl_ConnectRenderRequested(QtObjectPtr ptr){
	QObject::connect(static_cast<QQuickRenderControl*>(ptr), static_cast<void (QQuickRenderControl::*)()>(&QQuickRenderControl::renderRequested), static_cast<MyQQuickRenderControl*>(ptr), static_cast<void (MyQQuickRenderControl::*)()>(&MyQQuickRenderControl::Signal_RenderRequested));;
}

void QQuickRenderControl_DisconnectRenderRequested(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QQuickRenderControl*>(ptr), static_cast<void (QQuickRenderControl::*)()>(&QQuickRenderControl::renderRequested), static_cast<MyQQuickRenderControl*>(ptr), static_cast<void (MyQQuickRenderControl::*)()>(&MyQQuickRenderControl::Signal_RenderRequested));;
}

QtObjectPtr QQuickRenderControl_RenderWindow(QtObjectPtr ptr, QtObjectPtr offset){
	return static_cast<QQuickRenderControl*>(ptr)->renderWindow(static_cast<QPoint*>(offset));
}

QtObjectPtr QQuickRenderControl_QQuickRenderControl_RenderWindowFor(QtObjectPtr win, QtObjectPtr offset){
	return QQuickRenderControl::renderWindowFor(static_cast<QQuickWindow*>(win), static_cast<QPoint*>(offset));
}

void QQuickRenderControl_ConnectSceneChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QQuickRenderControl*>(ptr), static_cast<void (QQuickRenderControl::*)()>(&QQuickRenderControl::sceneChanged), static_cast<MyQQuickRenderControl*>(ptr), static_cast<void (MyQQuickRenderControl::*)()>(&MyQQuickRenderControl::Signal_SceneChanged));;
}

void QQuickRenderControl_DisconnectSceneChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QQuickRenderControl*>(ptr), static_cast<void (QQuickRenderControl::*)()>(&QQuickRenderControl::sceneChanged), static_cast<MyQQuickRenderControl*>(ptr), static_cast<void (MyQQuickRenderControl::*)()>(&MyQQuickRenderControl::Signal_SceneChanged));;
}

int QQuickRenderControl_Sync(QtObjectPtr ptr){
	return static_cast<QQuickRenderControl*>(ptr)->sync();
}

void QQuickRenderControl_DestroyQQuickRenderControl(QtObjectPtr ptr){
	static_cast<QQuickRenderControl*>(ptr)->~QQuickRenderControl();
}

