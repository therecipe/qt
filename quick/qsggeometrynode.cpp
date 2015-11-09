#include "qsggeometrynode.h"
#include <QSGGeometry>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSGMaterial>
#include <QSGGeometryNode>
#include "_cgo_export.h"

class MyQSGGeometryNode: public QSGGeometryNode {
public:
};

void* QSGGeometryNode_NewQSGGeometryNode(){
	return new QSGGeometryNode();
}

void* QSGGeometryNode_Material(void* ptr){
	return static_cast<QSGGeometryNode*>(ptr)->material();
}

void* QSGGeometryNode_OpaqueMaterial(void* ptr){
	return static_cast<QSGGeometryNode*>(ptr)->opaqueMaterial();
}

void QSGGeometryNode_SetMaterial(void* ptr, void* material){
	static_cast<QSGGeometryNode*>(ptr)->setMaterial(static_cast<QSGMaterial*>(material));
}

void QSGGeometryNode_SetOpaqueMaterial(void* ptr, void* material){
	static_cast<QSGGeometryNode*>(ptr)->setOpaqueMaterial(static_cast<QSGMaterial*>(material));
}

void QSGGeometryNode_DestroyQSGGeometryNode(void* ptr){
	static_cast<QSGGeometryNode*>(ptr)->~QSGGeometryNode();
}

