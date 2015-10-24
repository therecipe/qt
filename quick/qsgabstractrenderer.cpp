#include "qsgabstractrenderer.h"
#include <QRectF>
#include <QVariant>
#include <QModelIndex>
#include <QSize>
#include <QObject>
#include <QMatrix4x4>
#include <QRect>
#include <QString>
#include <QUrl>
#include <QColor>
#include <QSGAbstractRenderer>
#include "_cgo_export.h"

class MyQSGAbstractRenderer: public QSGAbstractRenderer {
public:
void Signal_SceneGraphChanged(){callbackQSGAbstractRendererSceneGraphChanged(this->objectName().toUtf8().data());};
};

int QSGAbstractRenderer_ClearMode(QtObjectPtr ptr){
	return static_cast<QSGAbstractRenderer*>(ptr)->clearMode();
}

void QSGAbstractRenderer_ConnectSceneGraphChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QSGAbstractRenderer*>(ptr), static_cast<void (QSGAbstractRenderer::*)()>(&QSGAbstractRenderer::sceneGraphChanged), static_cast<MyQSGAbstractRenderer*>(ptr), static_cast<void (MyQSGAbstractRenderer::*)()>(&MyQSGAbstractRenderer::Signal_SceneGraphChanged));;
}

void QSGAbstractRenderer_DisconnectSceneGraphChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QSGAbstractRenderer*>(ptr), static_cast<void (QSGAbstractRenderer::*)()>(&QSGAbstractRenderer::sceneGraphChanged), static_cast<MyQSGAbstractRenderer*>(ptr), static_cast<void (MyQSGAbstractRenderer::*)()>(&MyQSGAbstractRenderer::Signal_SceneGraphChanged));;
}

void QSGAbstractRenderer_SetClearColor(QtObjectPtr ptr, QtObjectPtr color){
	static_cast<QSGAbstractRenderer*>(ptr)->setClearColor(*static_cast<QColor*>(color));
}

void QSGAbstractRenderer_SetClearMode(QtObjectPtr ptr, int mode){
	static_cast<QSGAbstractRenderer*>(ptr)->setClearMode(static_cast<QSGAbstractRenderer::ClearModeBit>(mode));
}

void QSGAbstractRenderer_SetDeviceRect(QtObjectPtr ptr, QtObjectPtr rect){
	static_cast<QSGAbstractRenderer*>(ptr)->setDeviceRect(*static_cast<QRect*>(rect));
}

void QSGAbstractRenderer_SetDeviceRect2(QtObjectPtr ptr, QtObjectPtr size){
	static_cast<QSGAbstractRenderer*>(ptr)->setDeviceRect(*static_cast<QSize*>(size));
}

void QSGAbstractRenderer_SetProjectionMatrix(QtObjectPtr ptr, QtObjectPtr matrix){
	static_cast<QSGAbstractRenderer*>(ptr)->setProjectionMatrix(*static_cast<QMatrix4x4*>(matrix));
}

void QSGAbstractRenderer_SetProjectionMatrixToRect(QtObjectPtr ptr, QtObjectPtr rect){
	static_cast<QSGAbstractRenderer*>(ptr)->setProjectionMatrixToRect(*static_cast<QRectF*>(rect));
}

void QSGAbstractRenderer_SetViewportRect(QtObjectPtr ptr, QtObjectPtr rect){
	static_cast<QSGAbstractRenderer*>(ptr)->setViewportRect(*static_cast<QRect*>(rect));
}

void QSGAbstractRenderer_SetViewportRect2(QtObjectPtr ptr, QtObjectPtr size){
	static_cast<QSGAbstractRenderer*>(ptr)->setViewportRect(*static_cast<QSize*>(size));
}

