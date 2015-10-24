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

void QSGTexture_Bind(QtObjectPtr ptr){
	static_cast<QSGTexture*>(ptr)->bind();
}

int QSGTexture_Filtering(QtObjectPtr ptr){
	return static_cast<QSGTexture*>(ptr)->filtering();
}

int QSGTexture_HasAlphaChannel(QtObjectPtr ptr){
	return static_cast<QSGTexture*>(ptr)->hasAlphaChannel();
}

int QSGTexture_HasMipmaps(QtObjectPtr ptr){
	return static_cast<QSGTexture*>(ptr)->hasMipmaps();
}

int QSGTexture_HorizontalWrapMode(QtObjectPtr ptr){
	return static_cast<QSGTexture*>(ptr)->horizontalWrapMode();
}

int QSGTexture_IsAtlasTexture(QtObjectPtr ptr){
	return static_cast<QSGTexture*>(ptr)->isAtlasTexture();
}

int QSGTexture_MipmapFiltering(QtObjectPtr ptr){
	return static_cast<QSGTexture*>(ptr)->mipmapFiltering();
}

QtObjectPtr QSGTexture_RemovedFromAtlas(QtObjectPtr ptr){
	return static_cast<QSGTexture*>(ptr)->removedFromAtlas();
}

void QSGTexture_SetFiltering(QtObjectPtr ptr, int filter){
	static_cast<QSGTexture*>(ptr)->setFiltering(static_cast<QSGTexture::Filtering>(filter));
}

void QSGTexture_SetHorizontalWrapMode(QtObjectPtr ptr, int hwrap){
	static_cast<QSGTexture*>(ptr)->setHorizontalWrapMode(static_cast<QSGTexture::WrapMode>(hwrap));
}

void QSGTexture_SetMipmapFiltering(QtObjectPtr ptr, int filter){
	static_cast<QSGTexture*>(ptr)->setMipmapFiltering(static_cast<QSGTexture::Filtering>(filter));
}

void QSGTexture_SetVerticalWrapMode(QtObjectPtr ptr, int vwrap){
	static_cast<QSGTexture*>(ptr)->setVerticalWrapMode(static_cast<QSGTexture::WrapMode>(vwrap));
}

int QSGTexture_TextureId(QtObjectPtr ptr){
	return static_cast<QSGTexture*>(ptr)->textureId();
}

void QSGTexture_UpdateBindOptions(QtObjectPtr ptr, int force){
	static_cast<QSGTexture*>(ptr)->updateBindOptions(force != 0);
}

int QSGTexture_VerticalWrapMode(QtObjectPtr ptr){
	return static_cast<QSGTexture*>(ptr)->verticalWrapMode();
}

void QSGTexture_DestroyQSGTexture(QtObjectPtr ptr){
	static_cast<QSGTexture*>(ptr)->~QSGTexture();
}

