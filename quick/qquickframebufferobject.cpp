#include "qquickframebufferobject.h"
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QQuickFramebufferObject>
#include "_cgo_export.h"

class MyQQuickFramebufferObject: public QQuickFramebufferObject {
public:
void Signal_TextureFollowsItemSizeChanged(bool v){callbackQQuickFramebufferObjectTextureFollowsItemSizeChanged(this->objectName().toUtf8().data(), v);};
};

void QQuickFramebufferObject_SetTextureFollowsItemSize(void* ptr, int follows){
	static_cast<QQuickFramebufferObject*>(ptr)->setTextureFollowsItemSize(follows != 0);
}

int QQuickFramebufferObject_TextureFollowsItemSize(void* ptr){
	return static_cast<QQuickFramebufferObject*>(ptr)->textureFollowsItemSize();
}

int QQuickFramebufferObject_IsTextureProvider(void* ptr){
	return static_cast<QQuickFramebufferObject*>(ptr)->isTextureProvider();
}

void QQuickFramebufferObject_ReleaseResources(void* ptr){
	static_cast<QQuickFramebufferObject*>(ptr)->releaseResources();
}

void QQuickFramebufferObject_ConnectTextureFollowsItemSizeChanged(void* ptr){
	QObject::connect(static_cast<QQuickFramebufferObject*>(ptr), static_cast<void (QQuickFramebufferObject::*)(bool)>(&QQuickFramebufferObject::textureFollowsItemSizeChanged), static_cast<MyQQuickFramebufferObject*>(ptr), static_cast<void (MyQQuickFramebufferObject::*)(bool)>(&MyQQuickFramebufferObject::Signal_TextureFollowsItemSizeChanged));;
}

void QQuickFramebufferObject_DisconnectTextureFollowsItemSizeChanged(void* ptr){
	QObject::disconnect(static_cast<QQuickFramebufferObject*>(ptr), static_cast<void (QQuickFramebufferObject::*)(bool)>(&QQuickFramebufferObject::textureFollowsItemSizeChanged), static_cast<MyQQuickFramebufferObject*>(ptr), static_cast<void (MyQQuickFramebufferObject::*)(bool)>(&MyQQuickFramebufferObject::Signal_TextureFollowsItemSizeChanged));;
}

void* QQuickFramebufferObject_TextureProvider(void* ptr){
	return static_cast<QQuickFramebufferObject*>(ptr)->textureProvider();
}

