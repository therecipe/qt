#include "qquickframebufferobject.h"
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QQuickFramebufferObject>
#include "_cgo_export.h"

class MyQQuickFramebufferObject: public QQuickFramebufferObject {
public:
void Signal_TextureFollowsItemSizeChanged(bool v){callbackQQuickFramebufferObjectTextureFollowsItemSizeChanged(this->objectName().toUtf8().data(), v);};
};

void QQuickFramebufferObject_SetTextureFollowsItemSize(QtObjectPtr ptr, int follows){
	static_cast<QQuickFramebufferObject*>(ptr)->setTextureFollowsItemSize(follows != 0);
}

int QQuickFramebufferObject_TextureFollowsItemSize(QtObjectPtr ptr){
	return static_cast<QQuickFramebufferObject*>(ptr)->textureFollowsItemSize();
}

int QQuickFramebufferObject_IsTextureProvider(QtObjectPtr ptr){
	return static_cast<QQuickFramebufferObject*>(ptr)->isTextureProvider();
}

void QQuickFramebufferObject_ReleaseResources(QtObjectPtr ptr){
	static_cast<QQuickFramebufferObject*>(ptr)->releaseResources();
}

void QQuickFramebufferObject_ConnectTextureFollowsItemSizeChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QQuickFramebufferObject*>(ptr), static_cast<void (QQuickFramebufferObject::*)(bool)>(&QQuickFramebufferObject::textureFollowsItemSizeChanged), static_cast<MyQQuickFramebufferObject*>(ptr), static_cast<void (MyQQuickFramebufferObject::*)(bool)>(&MyQQuickFramebufferObject::Signal_TextureFollowsItemSizeChanged));;
}

void QQuickFramebufferObject_DisconnectTextureFollowsItemSizeChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QQuickFramebufferObject*>(ptr), static_cast<void (QQuickFramebufferObject::*)(bool)>(&QQuickFramebufferObject::textureFollowsItemSizeChanged), static_cast<MyQQuickFramebufferObject*>(ptr), static_cast<void (MyQQuickFramebufferObject::*)(bool)>(&MyQQuickFramebufferObject::Signal_TextureFollowsItemSizeChanged));;
}

QtObjectPtr QQuickFramebufferObject_TextureProvider(QtObjectPtr ptr){
	return static_cast<QQuickFramebufferObject*>(ptr)->textureProvider();
}

