#include "qsgopaquetexturematerial.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSGTexture>
#include <QSGOpaqueTextureMaterial>
#include "_cgo_export.h"

class MyQSGOpaqueTextureMaterial: public QSGOpaqueTextureMaterial {
public:
};

int QSGOpaqueTextureMaterial_Filtering(void* ptr){
	return static_cast<QSGOpaqueTextureMaterial*>(ptr)->filtering();
}

int QSGOpaqueTextureMaterial_HorizontalWrapMode(void* ptr){
	return static_cast<QSGOpaqueTextureMaterial*>(ptr)->horizontalWrapMode();
}

int QSGOpaqueTextureMaterial_MipmapFiltering(void* ptr){
	return static_cast<QSGOpaqueTextureMaterial*>(ptr)->mipmapFiltering();
}

void QSGOpaqueTextureMaterial_SetFiltering(void* ptr, int filtering){
	static_cast<QSGOpaqueTextureMaterial*>(ptr)->setFiltering(static_cast<QSGTexture::Filtering>(filtering));
}

void QSGOpaqueTextureMaterial_SetHorizontalWrapMode(void* ptr, int mode){
	static_cast<QSGOpaqueTextureMaterial*>(ptr)->setHorizontalWrapMode(static_cast<QSGTexture::WrapMode>(mode));
}

void QSGOpaqueTextureMaterial_SetMipmapFiltering(void* ptr, int filtering){
	static_cast<QSGOpaqueTextureMaterial*>(ptr)->setMipmapFiltering(static_cast<QSGTexture::Filtering>(filtering));
}

void QSGOpaqueTextureMaterial_SetTexture(void* ptr, void* texture){
	static_cast<QSGOpaqueTextureMaterial*>(ptr)->setTexture(static_cast<QSGTexture*>(texture));
}

void QSGOpaqueTextureMaterial_SetVerticalWrapMode(void* ptr, int mode){
	static_cast<QSGOpaqueTextureMaterial*>(ptr)->setVerticalWrapMode(static_cast<QSGTexture::WrapMode>(mode));
}

void* QSGOpaqueTextureMaterial_Texture(void* ptr){
	return static_cast<QSGOpaqueTextureMaterial*>(ptr)->texture();
}

int QSGOpaqueTextureMaterial_VerticalWrapMode(void* ptr){
	return static_cast<QSGOpaqueTextureMaterial*>(ptr)->verticalWrapMode();
}

