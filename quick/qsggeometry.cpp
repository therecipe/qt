#include "qsggeometry.h"
#include <QUrl>
#include <QModelIndex>
#include <QRect>
#include <QRectF>
#include <QString>
#include <QVariant>
#include <QSGGeometry>
#include "_cgo_export.h"

class MyQSGGeometry: public QSGGeometry {
public:
};

void QSGGeometry_Allocate(QtObjectPtr ptr, int vertexCount, int indexCount){
	static_cast<QSGGeometry*>(ptr)->allocate(vertexCount, indexCount);
}

int QSGGeometry_AttributeCount(QtObjectPtr ptr){
	return static_cast<QSGGeometry*>(ptr)->attributeCount();
}

int QSGGeometry_IndexCount(QtObjectPtr ptr){
	return static_cast<QSGGeometry*>(ptr)->indexCount();
}

void QSGGeometry_IndexData(QtObjectPtr ptr){
	static_cast<QSGGeometry*>(ptr)->indexData();
}

void QSGGeometry_IndexData2(QtObjectPtr ptr){
	static_cast<QSGGeometry*>(ptr)->indexData();
}

int QSGGeometry_IndexDataPattern(QtObjectPtr ptr){
	return static_cast<QSGGeometry*>(ptr)->indexDataPattern();
}

int QSGGeometry_IndexType(QtObjectPtr ptr){
	return static_cast<QSGGeometry*>(ptr)->indexType();
}

void QSGGeometry_MarkIndexDataDirty(QtObjectPtr ptr){
	static_cast<QSGGeometry*>(ptr)->markIndexDataDirty();
}

void QSGGeometry_MarkVertexDataDirty(QtObjectPtr ptr){
	static_cast<QSGGeometry*>(ptr)->markVertexDataDirty();
}

void QSGGeometry_SetIndexDataPattern(QtObjectPtr ptr, int p){
	static_cast<QSGGeometry*>(ptr)->setIndexDataPattern(static_cast<QSGGeometry::DataPattern>(p));
}

void QSGGeometry_SetVertexDataPattern(QtObjectPtr ptr, int p){
	static_cast<QSGGeometry*>(ptr)->setVertexDataPattern(static_cast<QSGGeometry::DataPattern>(p));
}

int QSGGeometry_SizeOfIndex(QtObjectPtr ptr){
	return static_cast<QSGGeometry*>(ptr)->sizeOfIndex();
}

int QSGGeometry_SizeOfVertex(QtObjectPtr ptr){
	return static_cast<QSGGeometry*>(ptr)->sizeOfVertex();
}

void QSGGeometry_QSGGeometry_UpdateRectGeometry(QtObjectPtr g, QtObjectPtr rect){
	QSGGeometry::updateRectGeometry(static_cast<QSGGeometry*>(g), *static_cast<QRectF*>(rect));
}

void QSGGeometry_QSGGeometry_UpdateTexturedRectGeometry(QtObjectPtr g, QtObjectPtr rect, QtObjectPtr textureRect){
	QSGGeometry::updateTexturedRectGeometry(static_cast<QSGGeometry*>(g), *static_cast<QRectF*>(rect), *static_cast<QRectF*>(textureRect));
}

int QSGGeometry_VertexCount(QtObjectPtr ptr){
	return static_cast<QSGGeometry*>(ptr)->vertexCount();
}

void QSGGeometry_VertexData(QtObjectPtr ptr){
	static_cast<QSGGeometry*>(ptr)->vertexData();
}

void QSGGeometry_VertexData2(QtObjectPtr ptr){
	static_cast<QSGGeometry*>(ptr)->vertexData();
}

int QSGGeometry_VertexDataPattern(QtObjectPtr ptr){
	return static_cast<QSGGeometry*>(ptr)->vertexDataPattern();
}

void QSGGeometry_DestroyQSGGeometry(QtObjectPtr ptr){
	static_cast<QSGGeometry*>(ptr)->~QSGGeometry();
}

