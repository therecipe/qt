#include "qsgtextureprovider.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSGTexture>
#include <QObject>
#include <QSGTextureProvider>
#include "_cgo_export.h"

class MyQSGTextureProvider: public QSGTextureProvider {
public:
void Signal_TextureChanged(){callbackQSGTextureProviderTextureChanged(this->objectName().toUtf8().data());};
};

QtObjectPtr QSGTextureProvider_Texture(QtObjectPtr ptr){
	return static_cast<QSGTextureProvider*>(ptr)->texture();
}

void QSGTextureProvider_ConnectTextureChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QSGTextureProvider*>(ptr), static_cast<void (QSGTextureProvider::*)()>(&QSGTextureProvider::textureChanged), static_cast<MyQSGTextureProvider*>(ptr), static_cast<void (MyQSGTextureProvider::*)()>(&MyQSGTextureProvider::Signal_TextureChanged));;
}

void QSGTextureProvider_DisconnectTextureChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QSGTextureProvider*>(ptr), static_cast<void (QSGTextureProvider::*)()>(&QSGTextureProvider::textureChanged), static_cast<MyQSGTextureProvider*>(ptr), static_cast<void (MyQSGTextureProvider::*)()>(&MyQSGTextureProvider::Signal_TextureChanged));;
}

