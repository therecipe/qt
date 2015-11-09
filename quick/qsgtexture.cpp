#include "qsgtexture.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSGTexture>
#include "_cgo_export.h"

class MyQSGTexture: public QSGTexture {
public:
};

void QSGTexture_Bind(void* ptr){
	static_cast<QSGTexture*>(ptr)->bind();
}

int QSGTexture_Filtering(void* ptr){
	return static_cast<QSGTexture*>(ptr)->filtering();
}

int QSGTexture_HasAlphaChannel(void* ptr){
	return static_cast<QSGTexture*>(ptr)->hasAlphaChannel();
}

int QSGTexture_HasMipmaps(void* ptr){
	return static_cast<QSGTexture*>(ptr)->hasMipmaps();
}

int QSGTexture_HorizontalWrapMode(void* ptr){
	return static_cast<QSGTexture*>(ptr)->horizontalWrapMode();
}

int QSGTexture_IsAtlasTexture(void* ptr){
	return static_cast<QSGTexture*>(ptr)->isAtlasTexture();
}

int QSGTexture_MipmapFiltering(void* ptr){
	return static_cast<QSGTexture*>(ptr)->mipmapFiltering();
}

void* QSGTexture_RemovedFromAtlas(void* ptr){
	return static_cast<QSGTexture*>(ptr)->removedFromAtlas();
}

void QSGTexture_SetFiltering(void* ptr, int filter){
	static_cast<QSGTexture*>(ptr)->setFiltering(static_cast<QSGTexture::Filtering>(filter));
}

void QSGTexture_SetHorizontalWrapMode(void* ptr, int hwrap){
	static_cast<QSGTexture*>(ptr)->setHorizontalWrapMode(static_cast<QSGTexture::WrapMode>(hwrap));
}

void QSGTexture_SetMipmapFiltering(void* ptr, int filter){
	static_cast<QSGTexture*>(ptr)->setMipmapFiltering(static_cast<QSGTexture::Filtering>(filter));
}

void QSGTexture_SetVerticalWrapMode(void* ptr, int vwrap){
	static_cast<QSGTexture*>(ptr)->setVerticalWrapMode(static_cast<QSGTexture::WrapMode>(vwrap));
}

int QSGTexture_TextureId(void* ptr){
	return static_cast<QSGTexture*>(ptr)->textureId();
}

void QSGTexture_UpdateBindOptions(void* ptr, int force){
	static_cast<QSGTexture*>(ptr)->updateBindOptions(force != 0);
}

int QSGTexture_VerticalWrapMode(void* ptr){
	return static_cast<QSGTexture*>(ptr)->verticalWrapMode();
}

void QSGTexture_DestroyQSGTexture(void* ptr){
	static_cast<QSGTexture*>(ptr)->~QSGTexture();
}

