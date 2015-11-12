#include "qsgabstractrenderer.h"
#include <QModelIndex>
#include <QObject>
#include <QSize>
#include <QRectF>
#include <QRect>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QColor>
#include <QMatrix4x4>
#include <QSGAbstractRenderer>
#include "_cgo_export.h"

class MyQSGAbstractRenderer: public QSGAbstractRenderer {
public:
void Signal_SceneGraphChanged(){callbackQSGAbstractRendererSceneGraphChanged(this->objectName().toUtf8().data());};
};

void* QSGAbstractRenderer_ClearColor(void* ptr){
	return new QColor(static_cast<QSGAbstractRenderer*>(ptr)->clearColor());
}

int QSGAbstractRenderer_ClearMode(void* ptr){
	return static_cast<QSGAbstractRenderer*>(ptr)->clearMode();
}

void QSGAbstractRenderer_ConnectSceneGraphChanged(void* ptr){
	QObject::connect(static_cast<QSGAbstractRenderer*>(ptr), static_cast<void (QSGAbstractRenderer::*)()>(&QSGAbstractRenderer::sceneGraphChanged), static_cast<MyQSGAbstractRenderer*>(ptr), static_cast<void (MyQSGAbstractRenderer::*)()>(&MyQSGAbstractRenderer::Signal_SceneGraphChanged));;
}

void QSGAbstractRenderer_DisconnectSceneGraphChanged(void* ptr){
	QObject::disconnect(static_cast<QSGAbstractRenderer*>(ptr), static_cast<void (QSGAbstractRenderer::*)()>(&QSGAbstractRenderer::sceneGraphChanged), static_cast<MyQSGAbstractRenderer*>(ptr), static_cast<void (MyQSGAbstractRenderer::*)()>(&MyQSGAbstractRenderer::Signal_SceneGraphChanged));;
}

void QSGAbstractRenderer_SetClearColor(void* ptr, void* color){
	static_cast<QSGAbstractRenderer*>(ptr)->setClearColor(*static_cast<QColor*>(color));
}

void QSGAbstractRenderer_SetClearMode(void* ptr, int mode){
	static_cast<QSGAbstractRenderer*>(ptr)->setClearMode(static_cast<QSGAbstractRenderer::ClearModeBit>(mode));
}

void QSGAbstractRenderer_SetDeviceRect(void* ptr, void* rect){
	static_cast<QSGAbstractRenderer*>(ptr)->setDeviceRect(*static_cast<QRect*>(rect));
}

void QSGAbstractRenderer_SetDeviceRect2(void* ptr, void* size){
	static_cast<QSGAbstractRenderer*>(ptr)->setDeviceRect(*static_cast<QSize*>(size));
}

void QSGAbstractRenderer_SetProjectionMatrix(void* ptr, void* matrix){
	static_cast<QSGAbstractRenderer*>(ptr)->setProjectionMatrix(*static_cast<QMatrix4x4*>(matrix));
}

void QSGAbstractRenderer_SetProjectionMatrixToRect(void* ptr, void* rect){
	static_cast<QSGAbstractRenderer*>(ptr)->setProjectionMatrixToRect(*static_cast<QRectF*>(rect));
}

void QSGAbstractRenderer_SetViewportRect(void* ptr, void* rect){
	static_cast<QSGAbstractRenderer*>(ptr)->setViewportRect(*static_cast<QRect*>(rect));
}

void QSGAbstractRenderer_SetViewportRect2(void* ptr, void* size){
	static_cast<QSGAbstractRenderer*>(ptr)->setViewportRect(*static_cast<QSize*>(size));
}

