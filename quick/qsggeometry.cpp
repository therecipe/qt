#include "qsggeometry.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QRectF>
#include <QRect>
#include <QSGGeometry>
#include "_cgo_export.h"

class MyQSGGeometry: public QSGGeometry {
public:
};

void QSGGeometry_Allocate(void* ptr, int vertexCount, int indexCount){
	static_cast<QSGGeometry*>(ptr)->allocate(vertexCount, indexCount);
}

int QSGGeometry_AttributeCount(void* ptr){
	return static_cast<QSGGeometry*>(ptr)->attributeCount();
}

int QSGGeometry_IndexCount(void* ptr){
	return static_cast<QSGGeometry*>(ptr)->indexCount();
}

void* QSGGeometry_IndexData(void* ptr){
	return static_cast<QSGGeometry*>(ptr)->indexData();
}

void* QSGGeometry_IndexData2(void* ptr){
	return const_cast<void*>(static_cast<QSGGeometry*>(ptr)->indexData());
}

int QSGGeometry_IndexDataPattern(void* ptr){
	return static_cast<QSGGeometry*>(ptr)->indexDataPattern();
}

int QSGGeometry_IndexType(void* ptr){
	return static_cast<QSGGeometry*>(ptr)->indexType();
}

void QSGGeometry_MarkIndexDataDirty(void* ptr){
	static_cast<QSGGeometry*>(ptr)->markIndexDataDirty();
}

void QSGGeometry_MarkVertexDataDirty(void* ptr){
	static_cast<QSGGeometry*>(ptr)->markVertexDataDirty();
}

void QSGGeometry_SetIndexDataPattern(void* ptr, int p){
	static_cast<QSGGeometry*>(ptr)->setIndexDataPattern(static_cast<QSGGeometry::DataPattern>(p));
}

void QSGGeometry_SetVertexDataPattern(void* ptr, int p){
	static_cast<QSGGeometry*>(ptr)->setVertexDataPattern(static_cast<QSGGeometry::DataPattern>(p));
}

int QSGGeometry_SizeOfIndex(void* ptr){
	return static_cast<QSGGeometry*>(ptr)->sizeOfIndex();
}

int QSGGeometry_SizeOfVertex(void* ptr){
	return static_cast<QSGGeometry*>(ptr)->sizeOfVertex();
}

void QSGGeometry_QSGGeometry_UpdateRectGeometry(void* g, void* rect){
	QSGGeometry::updateRectGeometry(static_cast<QSGGeometry*>(g), *static_cast<QRectF*>(rect));
}

void QSGGeometry_QSGGeometry_UpdateTexturedRectGeometry(void* g, void* rect, void* textureRect){
	QSGGeometry::updateTexturedRectGeometry(static_cast<QSGGeometry*>(g), *static_cast<QRectF*>(rect), *static_cast<QRectF*>(textureRect));
}

int QSGGeometry_VertexCount(void* ptr){
	return static_cast<QSGGeometry*>(ptr)->vertexCount();
}

void* QSGGeometry_VertexData(void* ptr){
	return static_cast<QSGGeometry*>(ptr)->vertexData();
}

void* QSGGeometry_VertexData2(void* ptr){
	return const_cast<void*>(static_cast<QSGGeometry*>(ptr)->vertexData());
}

int QSGGeometry_VertexDataPattern(void* ptr){
	return static_cast<QSGGeometry*>(ptr)->vertexDataPattern();
}

void QSGGeometry_DestroyQSGGeometry(void* ptr){
	static_cast<QSGGeometry*>(ptr)->~QSGGeometry();
}

