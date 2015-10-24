#include "qquicktexturefactory.h"
#include <QModelIndex>
#include <QQuickWindow>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QQuickTextureFactory>
#include "_cgo_export.h"

class MyQQuickTextureFactory: public QQuickTextureFactory {
public:
};

QtObjectPtr QQuickTextureFactory_CreateTexture(QtObjectPtr ptr, QtObjectPtr window){
	return static_cast<QQuickTextureFactory*>(ptr)->createTexture(static_cast<QQuickWindow*>(window));
}

int QQuickTextureFactory_TextureByteCount(QtObjectPtr ptr){
	return static_cast<QQuickTextureFactory*>(ptr)->textureByteCount();
}

void QQuickTextureFactory_DestroyQQuickTextureFactory(QtObjectPtr ptr){
	static_cast<QQuickTextureFactory*>(ptr)->~QQuickTextureFactory();
}

