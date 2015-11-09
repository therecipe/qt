#include "qsgsimpletexturenode.h"
#include <QSGTexture>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QRectF>
#include <QRect>
#include <QSGSimpleTextureNode>
#include "_cgo_export.h"

class MyQSGSimpleTextureNode: public QSGSimpleTextureNode {
public:
};

void* QSGSimpleTextureNode_NewQSGSimpleTextureNode(){
	return new QSGSimpleTextureNode();
}

int QSGSimpleTextureNode_Filtering(void* ptr){
	return static_cast<QSGSimpleTextureNode*>(ptr)->filtering();
}

int QSGSimpleTextureNode_OwnsTexture(void* ptr){
	return static_cast<QSGSimpleTextureNode*>(ptr)->ownsTexture();
}

void QSGSimpleTextureNode_SetFiltering(void* ptr, int filtering){
	static_cast<QSGSimpleTextureNode*>(ptr)->setFiltering(static_cast<QSGTexture::Filtering>(filtering));
}

void QSGSimpleTextureNode_SetOwnsTexture(void* ptr, int owns){
	static_cast<QSGSimpleTextureNode*>(ptr)->setOwnsTexture(owns != 0);
}

void QSGSimpleTextureNode_SetRect(void* ptr, void* r){
	static_cast<QSGSimpleTextureNode*>(ptr)->setRect(*static_cast<QRectF*>(r));
}

void QSGSimpleTextureNode_SetRect2(void* ptr, double x, double y, double w, double h){
	static_cast<QSGSimpleTextureNode*>(ptr)->setRect(static_cast<qreal>(x), static_cast<qreal>(y), static_cast<qreal>(w), static_cast<qreal>(h));
}

void QSGSimpleTextureNode_SetSourceRect(void* ptr, void* r){
	static_cast<QSGSimpleTextureNode*>(ptr)->setSourceRect(*static_cast<QRectF*>(r));
}

void QSGSimpleTextureNode_SetSourceRect2(void* ptr, double x, double y, double w, double h){
	static_cast<QSGSimpleTextureNode*>(ptr)->setSourceRect(static_cast<qreal>(x), static_cast<qreal>(y), static_cast<qreal>(w), static_cast<qreal>(h));
}

void QSGSimpleTextureNode_SetTexture(void* ptr, void* texture){
	static_cast<QSGSimpleTextureNode*>(ptr)->setTexture(static_cast<QSGTexture*>(texture));
}

void QSGSimpleTextureNode_SetTextureCoordinatesTransform(void* ptr, int mode){
	static_cast<QSGSimpleTextureNode*>(ptr)->setTextureCoordinatesTransform(static_cast<QSGSimpleTextureNode::TextureCoordinatesTransformFlag>(mode));
}

void* QSGSimpleTextureNode_Texture(void* ptr){
	return static_cast<QSGSimpleTextureNode*>(ptr)->texture();
}

int QSGSimpleTextureNode_TextureCoordinatesTransform(void* ptr){
	return static_cast<QSGSimpleTextureNode*>(ptr)->textureCoordinatesTransform();
}

void QSGSimpleTextureNode_DestroyQSGSimpleTextureNode(void* ptr){
	static_cast<QSGSimpleTextureNode*>(ptr)->~QSGSimpleTextureNode();
}

