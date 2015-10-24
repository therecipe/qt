#include "qsgsimpletexturenode.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSGTexture>
#include <QRect>
#include <QRectF>
#include <QSGSimpleTextureNode>
#include "_cgo_export.h"

class MyQSGSimpleTextureNode: public QSGSimpleTextureNode {
public:
};

QtObjectPtr QSGSimpleTextureNode_NewQSGSimpleTextureNode(){
	return new QSGSimpleTextureNode();
}

int QSGSimpleTextureNode_Filtering(QtObjectPtr ptr){
	return static_cast<QSGSimpleTextureNode*>(ptr)->filtering();
}

int QSGSimpleTextureNode_OwnsTexture(QtObjectPtr ptr){
	return static_cast<QSGSimpleTextureNode*>(ptr)->ownsTexture();
}

void QSGSimpleTextureNode_SetFiltering(QtObjectPtr ptr, int filtering){
	static_cast<QSGSimpleTextureNode*>(ptr)->setFiltering(static_cast<QSGTexture::Filtering>(filtering));
}

void QSGSimpleTextureNode_SetOwnsTexture(QtObjectPtr ptr, int owns){
	static_cast<QSGSimpleTextureNode*>(ptr)->setOwnsTexture(owns != 0);
}

void QSGSimpleTextureNode_SetRect(QtObjectPtr ptr, QtObjectPtr r){
	static_cast<QSGSimpleTextureNode*>(ptr)->setRect(*static_cast<QRectF*>(r));
}

void QSGSimpleTextureNode_SetSourceRect(QtObjectPtr ptr, QtObjectPtr r){
	static_cast<QSGSimpleTextureNode*>(ptr)->setSourceRect(*static_cast<QRectF*>(r));
}

void QSGSimpleTextureNode_SetTexture(QtObjectPtr ptr, QtObjectPtr texture){
	static_cast<QSGSimpleTextureNode*>(ptr)->setTexture(static_cast<QSGTexture*>(texture));
}

void QSGSimpleTextureNode_SetTextureCoordinatesTransform(QtObjectPtr ptr, int mode){
	static_cast<QSGSimpleTextureNode*>(ptr)->setTextureCoordinatesTransform(static_cast<QSGSimpleTextureNode::TextureCoordinatesTransformFlag>(mode));
}

QtObjectPtr QSGSimpleTextureNode_Texture(QtObjectPtr ptr){
	return static_cast<QSGSimpleTextureNode*>(ptr)->texture();
}

int QSGSimpleTextureNode_TextureCoordinatesTransform(QtObjectPtr ptr){
	return static_cast<QSGSimpleTextureNode*>(ptr)->textureCoordinatesTransform();
}

void QSGSimpleTextureNode_DestroyQSGSimpleTextureNode(QtObjectPtr ptr){
	static_cast<QSGSimpleTextureNode*>(ptr)->~QSGSimpleTextureNode();
}

