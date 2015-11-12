#include "qquicktexturefactory.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QQuickWindow>
#include <QQuickTextureFactory>
#include "_cgo_export.h"

class MyQQuickTextureFactory: public QQuickTextureFactory {
public:
};

void* QQuickTextureFactory_CreateTexture(void* ptr, void* window){
	return static_cast<QQuickTextureFactory*>(ptr)->createTexture(static_cast<QQuickWindow*>(window));
}

int QQuickTextureFactory_TextureByteCount(void* ptr){
	return static_cast<QQuickTextureFactory*>(ptr)->textureByteCount();
}

void QQuickTextureFactory_DestroyQQuickTextureFactory(void* ptr){
	static_cast<QQuickTextureFactory*>(ptr)->~QQuickTextureFactory();
}

