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

int QSGOpaqueTextureMaterial_Filtering(QtObjectPtr ptr){
	return static_cast<QSGOpaqueTextureMaterial*>(ptr)->filtering();
}

int QSGOpaqueTextureMaterial_HorizontalWrapMode(QtObjectPtr ptr){
	return static_cast<QSGOpaqueTextureMaterial*>(ptr)->horizontalWrapMode();
}

int QSGOpaqueTextureMaterial_MipmapFiltering(QtObjectPtr ptr){
	return static_cast<QSGOpaqueTextureMaterial*>(ptr)->mipmapFiltering();
}

void QSGOpaqueTextureMaterial_SetFiltering(QtObjectPtr ptr, int filtering){
	static_cast<QSGOpaqueTextureMaterial*>(ptr)->setFiltering(static_cast<QSGTexture::Filtering>(filtering));
}

void QSGOpaqueTextureMaterial_SetHorizontalWrapMode(QtObjectPtr ptr, int mode){
	static_cast<QSGOpaqueTextureMaterial*>(ptr)->setHorizontalWrapMode(static_cast<QSGTexture::WrapMode>(mode));
}

void QSGOpaqueTextureMaterial_SetMipmapFiltering(QtObjectPtr ptr, int filtering){
	static_cast<QSGOpaqueTextureMaterial*>(ptr)->setMipmapFiltering(static_cast<QSGTexture::Filtering>(filtering));
}

void QSGOpaqueTextureMaterial_SetTexture(QtObjectPtr ptr, QtObjectPtr texture){
	static_cast<QSGOpaqueTextureMaterial*>(ptr)->setTexture(static_cast<QSGTexture*>(texture));
}

void QSGOpaqueTextureMaterial_SetVerticalWrapMode(QtObjectPtr ptr, int mode){
	static_cast<QSGOpaqueTextureMaterial*>(ptr)->setVerticalWrapMode(static_cast<QSGTexture::WrapMode>(mode));
}

QtObjectPtr QSGOpaqueTextureMaterial_Texture(QtObjectPtr ptr){
	return static_cast<QSGOpaqueTextureMaterial*>(ptr)->texture();
}

int QSGOpaqueTextureMaterial_VerticalWrapMode(QtObjectPtr ptr){
	return static_cast<QSGOpaqueTextureMaterial*>(ptr)->verticalWrapMode();
}

